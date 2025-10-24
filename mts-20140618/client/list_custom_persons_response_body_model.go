// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomPersonsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategories(v *ListCustomPersonsResponseBodyCategories) *ListCustomPersonsResponseBody
	GetCategories() *ListCustomPersonsResponseBodyCategories
	SetRequestId(v string) *ListCustomPersonsResponseBody
	GetRequestId() *string
}

type ListCustomPersonsResponseBody struct {
	// The array of the figure libraries.
	Categories *ListCustomPersonsResponseBodyCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// FD4DED6B-0C26-5A8B-A6BE-4FA542AE4D57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCustomPersonsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPersonsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomPersonsResponseBody) GetCategories() *ListCustomPersonsResponseBodyCategories {
	return s.Categories
}

func (s *ListCustomPersonsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomPersonsResponseBody) SetCategories(v *ListCustomPersonsResponseBodyCategories) *ListCustomPersonsResponseBody {
	s.Categories = v
	return s
}

func (s *ListCustomPersonsResponseBody) SetRequestId(v string) *ListCustomPersonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomPersonsResponseBody) Validate() error {
	if s.Categories != nil {
		if err := s.Categories.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomPersonsResponseBodyCategories struct {
	Category []*ListCustomPersonsResponseBodyCategoriesCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Repeated"`
}

func (s ListCustomPersonsResponseBodyCategories) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPersonsResponseBodyCategories) GoString() string {
	return s.String()
}

func (s *ListCustomPersonsResponseBodyCategories) GetCategory() []*ListCustomPersonsResponseBodyCategoriesCategory {
	return s.Category
}

func (s *ListCustomPersonsResponseBodyCategories) SetCategory(v []*ListCustomPersonsResponseBodyCategoriesCategory) *ListCustomPersonsResponseBodyCategories {
	s.Category = v
	return s
}

func (s *ListCustomPersonsResponseBodyCategories) Validate() error {
	if s.Category != nil {
		for _, item := range s.Category {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomPersonsResponseBodyCategoriesCategory struct {
	// The description of the figure library.
	//
	// example:
	//
	// CategoryDescription-****
	CategoryDescription *string `json:"CategoryDescription,omitempty" xml:"CategoryDescription,omitempty"`
	// The ID of the figure library.
	//
	// example:
	//
	// CategoryId-****
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The name of the figure library.
	//
	// example:
	//
	// CategoryName-****
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The array of the figures.
	Persons *ListCustomPersonsResponseBodyCategoriesCategoryPersons `json:"Persons,omitempty" xml:"Persons,omitempty" type:"Struct"`
}

func (s ListCustomPersonsResponseBodyCategoriesCategory) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPersonsResponseBodyCategoriesCategory) GoString() string {
	return s.String()
}

func (s *ListCustomPersonsResponseBodyCategoriesCategory) GetCategoryDescription() *string {
	return s.CategoryDescription
}

func (s *ListCustomPersonsResponseBodyCategoriesCategory) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ListCustomPersonsResponseBodyCategoriesCategory) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ListCustomPersonsResponseBodyCategoriesCategory) GetPersons() *ListCustomPersonsResponseBodyCategoriesCategoryPersons {
	return s.Persons
}

func (s *ListCustomPersonsResponseBodyCategoriesCategory) SetCategoryDescription(v string) *ListCustomPersonsResponseBodyCategoriesCategory {
	s.CategoryDescription = &v
	return s
}

func (s *ListCustomPersonsResponseBodyCategoriesCategory) SetCategoryId(v string) *ListCustomPersonsResponseBodyCategoriesCategory {
	s.CategoryId = &v
	return s
}

func (s *ListCustomPersonsResponseBodyCategoriesCategory) SetCategoryName(v string) *ListCustomPersonsResponseBodyCategoriesCategory {
	s.CategoryName = &v
	return s
}

func (s *ListCustomPersonsResponseBodyCategoriesCategory) SetPersons(v *ListCustomPersonsResponseBodyCategoriesCategoryPersons) *ListCustomPersonsResponseBodyCategoriesCategory {
	s.Persons = v
	return s
}

func (s *ListCustomPersonsResponseBodyCategoriesCategory) Validate() error {
	if s.Persons != nil {
		if err := s.Persons.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomPersonsResponseBodyCategoriesCategoryPersons struct {
	Person []*ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson `json:"Person,omitempty" xml:"Person,omitempty" type:"Repeated"`
}

func (s ListCustomPersonsResponseBodyCategoriesCategoryPersons) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPersonsResponseBodyCategoriesCategoryPersons) GoString() string {
	return s.String()
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersons) GetPerson() []*ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson {
	return s.Person
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersons) SetPerson(v []*ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson) *ListCustomPersonsResponseBodyCategoriesCategoryPersons {
	s.Person = v
	return s
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersons) Validate() error {
	if s.Person != nil {
		for _, item := range s.Person {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson struct {
	// The array of the faces.
	Faces *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Struct"`
	// The description of the figure.
	//
	// example:
	//
	// PersonDescription-****
	PersonDescription *string `json:"PersonDescription,omitempty" xml:"PersonDescription,omitempty"`
	// The ID of the figure.
	//
	// example:
	//
	// PersonId-****
	PersonId *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	// The name of the figure.
	//
	// example:
	//
	// PersonName-****
	PersonName *string `json:"PersonName,omitempty" xml:"PersonName,omitempty"`
}

func (s ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson) GoString() string {
	return s.String()
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson) GetFaces() *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFaces {
	return s.Faces
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson) GetPersonDescription() *string {
	return s.PersonDescription
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson) GetPersonId() *string {
	return s.PersonId
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson) GetPersonName() *string {
	return s.PersonName
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson) SetFaces(v *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFaces) *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson {
	s.Faces = v
	return s
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson) SetPersonDescription(v string) *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson {
	s.PersonDescription = &v
	return s
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson) SetPersonId(v string) *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson {
	s.PersonId = &v
	return s
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson) SetPersonName(v string) *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson {
	s.PersonName = &v
	return s
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPerson) Validate() error {
	if s.Faces != nil {
		if err := s.Faces.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFaces struct {
	Face []*ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFacesFace `json:"Face,omitempty" xml:"Face,omitempty" type:"Repeated"`
}

func (s ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFaces) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFaces) GoString() string {
	return s.String()
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFaces) GetFace() []*ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFacesFace {
	return s.Face
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFaces) SetFace(v []*ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFacesFace) *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFaces {
	s.Face = v
	return s
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFaces) Validate() error {
	if s.Face != nil {
		for _, item := range s.Face {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFacesFace struct {
	// The ID of the face.
	//
	// example:
	//
	// 15****
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	// The URL of the facial image that was registered for the figure.
	//
	// example:
	//
	// http://example-****.jpeg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFacesFace) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFacesFace) GoString() string {
	return s.String()
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFacesFace) GetFaceId() *string {
	return s.FaceId
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFacesFace) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFacesFace) SetFaceId(v string) *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFacesFace {
	s.FaceId = &v
	return s
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFacesFace) SetImageUrl(v string) *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFacesFace {
	s.ImageUrl = &v
	return s
}

func (s *ListCustomPersonsResponseBodyCategoriesCategoryPersonsPersonFacesFace) Validate() error {
	return dara.Validate(s)
}
