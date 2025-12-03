// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColor(v string) *UpdateFlowTagRequest
	GetColor() *string
	SetFlowTagGroupId(v int64) *UpdateFlowTagRequest
	GetFlowTagGroupId() *int64
	SetName(v string) *UpdateFlowTagRequest
	GetName() *string
}

type UpdateFlowTagRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// #1F9AEF
	Color *string `json:"color,omitempty" xml:"color,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11
	FlowTagGroupId *int64 `json:"flowTagGroupId,omitempty" xml:"flowTagGroupId,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateFlowTagRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowTagRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowTagRequest) GetColor() *string {
	return s.Color
}

func (s *UpdateFlowTagRequest) GetFlowTagGroupId() *int64 {
	return s.FlowTagGroupId
}

func (s *UpdateFlowTagRequest) GetName() *string {
	return s.Name
}

func (s *UpdateFlowTagRequest) SetColor(v string) *UpdateFlowTagRequest {
	s.Color = &v
	return s
}

func (s *UpdateFlowTagRequest) SetFlowTagGroupId(v int64) *UpdateFlowTagRequest {
	s.FlowTagGroupId = &v
	return s
}

func (s *UpdateFlowTagRequest) SetName(v string) *UpdateFlowTagRequest {
	s.Name = &v
	return s
}

func (s *UpdateFlowTagRequest) Validate() error {
	return dara.Validate(s)
}
