// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkItemWorkFlowStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSpaceIdentifier(v string) *ListWorkItemWorkFlowStatusRequest
	GetSpaceIdentifier() *string
	SetSpaceType(v string) *ListWorkItemWorkFlowStatusRequest
	GetSpaceType() *string
	SetWorkitemCategoryIdentifier(v string) *ListWorkItemWorkFlowStatusRequest
	GetWorkitemCategoryIdentifier() *string
	SetWorkitemTypeIdentifier(v string) *ListWorkItemWorkFlowStatusRequest
	GetWorkitemTypeIdentifier() *string
}

type ListWorkItemWorkFlowStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 例：5e70xxxxxxcd000xxxxe96
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
	// Req
	WorkitemCategoryIdentifier *string `json:"workitemCategoryIdentifier,omitempty" xml:"workitemCategoryIdentifier,omitempty"`
	// example:
	//
	// 例：5e7xxxxb3cd3711dd6xxx2c
	WorkitemTypeIdentifier *string `json:"workitemTypeIdentifier,omitempty" xml:"workitemTypeIdentifier,omitempty"`
}

func (s ListWorkItemWorkFlowStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemWorkFlowStatusRequest) GoString() string {
	return s.String()
}

func (s *ListWorkItemWorkFlowStatusRequest) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *ListWorkItemWorkFlowStatusRequest) GetSpaceType() *string {
	return s.SpaceType
}

func (s *ListWorkItemWorkFlowStatusRequest) GetWorkitemCategoryIdentifier() *string {
	return s.WorkitemCategoryIdentifier
}

func (s *ListWorkItemWorkFlowStatusRequest) GetWorkitemTypeIdentifier() *string {
	return s.WorkitemTypeIdentifier
}

func (s *ListWorkItemWorkFlowStatusRequest) SetSpaceIdentifier(v string) *ListWorkItemWorkFlowStatusRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusRequest) SetSpaceType(v string) *ListWorkItemWorkFlowStatusRequest {
	s.SpaceType = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusRequest) SetWorkitemCategoryIdentifier(v string) *ListWorkItemWorkFlowStatusRequest {
	s.WorkitemCategoryIdentifier = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusRequest) SetWorkitemTypeIdentifier(v string) *ListWorkItemWorkFlowStatusRequest {
	s.WorkitemTypeIdentifier = &v
	return s
}

func (s *ListWorkItemWorkFlowStatusRequest) Validate() error {
	return dara.Validate(s)
}
