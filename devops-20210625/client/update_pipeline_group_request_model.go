// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdatePipelineGroupRequest
	GetName() *string
}

type UpdatePipelineGroupRequest struct {
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdatePipelineGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePipelineGroupRequest) SetName(v string) *UpdatePipelineGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdatePipelineGroupRequest) Validate() error {
	return dara.Validate(s)
}
