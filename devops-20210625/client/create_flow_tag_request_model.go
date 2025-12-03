// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColor(v string) *CreateFlowTagRequest
	GetColor() *string
	SetFlowTagGroupId(v int64) *CreateFlowTagRequest
	GetFlowTagGroupId() *int64
	SetName(v string) *CreateFlowTagRequest
	GetName() *string
}

type CreateFlowTagRequest struct {
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
	// 111
	FlowTagGroupId *int64 `json:"flowTagGroupId,omitempty" xml:"flowTagGroupId,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateFlowTagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowTagRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowTagRequest) GetColor() *string {
	return s.Color
}

func (s *CreateFlowTagRequest) GetFlowTagGroupId() *int64 {
	return s.FlowTagGroupId
}

func (s *CreateFlowTagRequest) GetName() *string {
	return s.Name
}

func (s *CreateFlowTagRequest) SetColor(v string) *CreateFlowTagRequest {
	s.Color = &v
	return s
}

func (s *CreateFlowTagRequest) SetFlowTagGroupId(v int64) *CreateFlowTagRequest {
	s.FlowTagGroupId = &v
	return s
}

func (s *CreateFlowTagRequest) SetName(v string) *CreateFlowTagRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowTagRequest) Validate() error {
	return dara.Validate(s)
}
