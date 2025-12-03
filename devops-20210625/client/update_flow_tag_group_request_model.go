// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowTagGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateFlowTagGroupRequest
	GetName() *string
}

type UpdateFlowTagGroupRequest struct {
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateFlowTagGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowTagGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowTagGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateFlowTagGroupRequest) SetName(v string) *UpdateFlowTagGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateFlowTagGroupRequest) Validate() error {
	return dara.Validate(s)
}
