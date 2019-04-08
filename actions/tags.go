package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/ryonzhang369/blog_app/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Tag)
// DB Table: Plural (tags)
// Resource: Plural (Tags)
// Path: Plural (/tags)
// View Template Folder: Plural (/templates/tags/)

// TagsResource is the resource for the Tag model
type TagsResource struct {
	buffalo.Resource
}

// List gets all Tags. This function is mapped to the path
// GET /tags
func (v TagsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	tags := &models.Tags{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Tags from the DB
	if err := q.All(tags); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.JSON(tags))
}

// Show gets the data for one Tag. This function is mapped to
// the path GET /tags/{tag_id}
func (v TagsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Tag
	tag := &models.Tag{}

	// To find the Tag the parameter tag_id is used.
	if err := tx.Find(tag, c.Param("tag_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.JSON(tag))
}

// New renders the form for creating a new Tag.
// This function is mapped to the path GET /tags/new
func (v TagsResource) New(c buffalo.Context) error {
	return c.Render(200, r.JSON(&models.Tag{}))
}

// Create adds a Tag to the DB. This function is mapped to the
// path POST /tags
func (v TagsResource) Create(c buffalo.Context) error {
	// Allocate an empty Tag
	tag := &models.Tag{}

	// Bind tag to the html form elements
	if err := c.Bind(tag); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(tag)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.JSON(tag))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "tag.created.success"))
	// and redirect to the tags index page
	return c.Render(201, r.JSON(tag))
}

// Edit renders a edit form for a Tag. This function is
// mapped to the path GET /tags/{tag_id}/edit
func (v TagsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Tag
	tag := &models.Tag{}

	if err := tx.Find(tag, c.Param("tag_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.JSON(tag))
}

// Update changes a Tag in the DB. This function is mapped to
// the path PUT /tags/{tag_id}
func (v TagsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Tag
	tag := &models.Tag{}

	if err := tx.Find(tag, c.Param("tag_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Tag to the html form elements
	if err := c.Bind(tag); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(tag)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.JSON(tag))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "tag.updated.success"))
	// and redirect to the tags index page
	return c.Render(200, r.JSON(tag))
}

// Destroy deletes a Tag from the DB. This function is mapped
// to the path DELETE /tags/{tag_id}
func (v TagsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Tag
	tag := &models.Tag{}

	// To find the Tag the parameter tag_id is used.
	if err := tx.Find(tag, c.Param("tag_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(tag); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", T.Translate(c, "tag.destroyed.success"))
	// Redirect to the tags index page
	return c.Render(200, r.JSON(tag))
}