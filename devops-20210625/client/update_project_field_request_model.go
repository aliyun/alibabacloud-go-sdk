// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStatusIdentifier(v string) *UpdateProjectFieldRequest
	GetStatusIdentifier() *string
	SetUpdateBasicFieldRequestList(v []*UpdateProjectFieldRequestUpdateBasicFieldRequestList) *UpdateProjectFieldRequest
	GetUpdateBasicFieldRequestList() []*UpdateProjectFieldRequestUpdateBasicFieldRequestList
	SetUpdateForOpenApiList(v []*UpdateProjectFieldRequestUpdateForOpenApiList) *UpdateProjectFieldRequest
	GetUpdateForOpenApiList() []*UpdateProjectFieldRequestUpdateForOpenApiList
}

type UpdateProjectFieldRequest struct {
	// example:
	//
	// fdsaadsfasxxxxdddd
	StatusIdentifier            *string                                                 `json:"statusIdentifier,omitempty" xml:"statusIdentifier,omitempty"`
	UpdateBasicFieldRequestList []*UpdateProjectFieldRequestUpdateBasicFieldRequestList `json:"updateBasicFieldRequestList,omitempty" xml:"updateBasicFieldRequestList,omitempty" type:"Repeated"`
	UpdateForOpenApiList        []*UpdateProjectFieldRequestUpdateForOpenApiList        `json:"updateForOpenApiList,omitempty" xml:"updateForOpenApiList,omitempty" type:"Repeated"`
}

func (s UpdateProjectFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectFieldRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectFieldRequest) GetStatusIdentifier() *string {
	return s.StatusIdentifier
}

func (s *UpdateProjectFieldRequest) GetUpdateBasicFieldRequestList() []*UpdateProjectFieldRequestUpdateBasicFieldRequestList {
	return s.UpdateBasicFieldRequestList
}

func (s *UpdateProjectFieldRequest) GetUpdateForOpenApiList() []*UpdateProjectFieldRequestUpdateForOpenApiList {
	return s.UpdateForOpenApiList
}

func (s *UpdateProjectFieldRequest) SetStatusIdentifier(v string) *UpdateProjectFieldRequest {
	s.StatusIdentifier = &v
	return s
}

func (s *UpdateProjectFieldRequest) SetUpdateBasicFieldRequestList(v []*UpdateProjectFieldRequestUpdateBasicFieldRequestList) *UpdateProjectFieldRequest {
	s.UpdateBasicFieldRequestList = v
	return s
}

func (s *UpdateProjectFieldRequest) SetUpdateForOpenApiList(v []*UpdateProjectFieldRequestUpdateForOpenApiList) *UpdateProjectFieldRequest {
	s.UpdateForOpenApiList = v
	return s
}

func (s *UpdateProjectFieldRequest) Validate() error {
	if s.UpdateBasicFieldRequestList != nil {
		for _, item := range s.UpdateBasicFieldRequestList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UpdateForOpenApiList != nil {
		for _, item := range s.UpdateForOpenApiList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateProjectFieldRequestUpdateBasicFieldRequestList struct {
	// example:
	//
	// name
	PropertyKey   *string `json:"propertyKey,omitempty" xml:"propertyKey,omitempty"`
	PropertyValue *string `json:"propertyValue,omitempty" xml:"propertyValue,omitempty"`
}

func (s UpdateProjectFieldRequestUpdateBasicFieldRequestList) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectFieldRequestUpdateBasicFieldRequestList) GoString() string {
	return s.String()
}

func (s *UpdateProjectFieldRequestUpdateBasicFieldRequestList) GetPropertyKey() *string {
	return s.PropertyKey
}

func (s *UpdateProjectFieldRequestUpdateBasicFieldRequestList) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *UpdateProjectFieldRequestUpdateBasicFieldRequestList) SetPropertyKey(v string) *UpdateProjectFieldRequestUpdateBasicFieldRequestList {
	s.PropertyKey = &v
	return s
}

func (s *UpdateProjectFieldRequestUpdateBasicFieldRequestList) SetPropertyValue(v string) *UpdateProjectFieldRequestUpdateBasicFieldRequestList {
	s.PropertyValue = &v
	return s
}

func (s *UpdateProjectFieldRequestUpdateBasicFieldRequestList) Validate() error {
	return dara.Validate(s)
}

type UpdateProjectFieldRequestUpdateForOpenApiList struct {
	// example:
	//
	// c4fd21xxxxxxxx9oj8jk
	FieldIdentifier *string `json:"fieldIdentifier,omitempty" xml:"fieldIdentifier,omitempty"`
	Value           *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateProjectFieldRequestUpdateForOpenApiList) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectFieldRequestUpdateForOpenApiList) GoString() string {
	return s.String()
}

func (s *UpdateProjectFieldRequestUpdateForOpenApiList) GetFieldIdentifier() *string {
	return s.FieldIdentifier
}

func (s *UpdateProjectFieldRequestUpdateForOpenApiList) GetValue() *string {
	return s.Value
}

func (s *UpdateProjectFieldRequestUpdateForOpenApiList) SetFieldIdentifier(v string) *UpdateProjectFieldRequestUpdateForOpenApiList {
	s.FieldIdentifier = &v
	return s
}

func (s *UpdateProjectFieldRequestUpdateForOpenApiList) SetValue(v string) *UpdateProjectFieldRequestUpdateForOpenApiList {
	s.Value = &v
	return s
}

func (s *UpdateProjectFieldRequestUpdateForOpenApiList) Validate() error {
	return dara.Validate(s)
}
