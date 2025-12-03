// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreatePipelineGroupRequest
	GetName() *string
}

type CreatePipelineGroupRequest struct {
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreatePipelineGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineGroupRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreatePipelineGroupRequest) SetName(v string) *CreatePipelineGroupRequest {
	s.Name = &v
	return s
}

func (s *CreatePipelineGroupRequest) Validate() error {
	return dara.Validate(s)
}
