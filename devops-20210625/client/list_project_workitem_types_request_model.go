// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectWorkitemTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListProjectWorkitemTypesRequest
	GetCategory() *string
	SetSpaceType(v string) *ListProjectWorkitemTypesRequest
	GetSpaceType() *string
}

type ListProjectWorkitemTypesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Req
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Project
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s ListProjectWorkitemTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectWorkitemTypesRequest) GoString() string {
	return s.String()
}

func (s *ListProjectWorkitemTypesRequest) GetCategory() *string {
	return s.Category
}

func (s *ListProjectWorkitemTypesRequest) GetSpaceType() *string {
	return s.SpaceType
}

func (s *ListProjectWorkitemTypesRequest) SetCategory(v string) *ListProjectWorkitemTypesRequest {
	s.Category = &v
	return s
}

func (s *ListProjectWorkitemTypesRequest) SetSpaceType(v string) *ListProjectWorkitemTypesRequest {
	s.SpaceType = &v
	return s
}

func (s *ListProjectWorkitemTypesRequest) Validate() error {
	return dara.Validate(s)
}
