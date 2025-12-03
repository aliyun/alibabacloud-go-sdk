// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListProjectTemplatesRequest
	GetCategory() *string
}

type ListProjectTemplatesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Project
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
}

func (s ListProjectTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListProjectTemplatesRequest) GetCategory() *string {
	return s.Category
}

func (s *ListProjectTemplatesRequest) SetCategory(v string) *ListProjectTemplatesRequest {
	s.Category = &v
	return s
}

func (s *ListProjectTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
