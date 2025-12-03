// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkItemAllFieldsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSpaceIdentifier(v string) *ListWorkItemAllFieldsRequest
	GetSpaceIdentifier() *string
	SetSpaceType(v string) *ListWorkItemAllFieldsRequest
	GetSpaceType() *string
	SetWorkitemTypeIdentifier(v string) *ListWorkItemAllFieldsRequest
	GetWorkitemTypeIdentifier() *string
}

type ListWorkItemAllFieldsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx23
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Project
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dfexxxxxf4fee18xxxxx36
	WorkitemTypeIdentifier *string `json:"workitemTypeIdentifier,omitempty" xml:"workitemTypeIdentifier,omitempty"`
}

func (s ListWorkItemAllFieldsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemAllFieldsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkItemAllFieldsRequest) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *ListWorkItemAllFieldsRequest) GetSpaceType() *string {
	return s.SpaceType
}

func (s *ListWorkItemAllFieldsRequest) GetWorkitemTypeIdentifier() *string {
	return s.WorkitemTypeIdentifier
}

func (s *ListWorkItemAllFieldsRequest) SetSpaceIdentifier(v string) *ListWorkItemAllFieldsRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *ListWorkItemAllFieldsRequest) SetSpaceType(v string) *ListWorkItemAllFieldsRequest {
	s.SpaceType = &v
	return s
}

func (s *ListWorkItemAllFieldsRequest) SetWorkitemTypeIdentifier(v string) *ListWorkItemAllFieldsRequest {
	s.WorkitemTypeIdentifier = &v
	return s
}

func (s *ListWorkItemAllFieldsRequest) Validate() error {
	return dara.Validate(s)
}
