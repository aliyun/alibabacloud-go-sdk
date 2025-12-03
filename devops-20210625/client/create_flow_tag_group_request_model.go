// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowTagGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateFlowTagGroupRequest
	GetName() *string
}

type CreateFlowTagGroupRequest struct {
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateFlowTagGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowTagGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowTagGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateFlowTagGroupRequest) SetName(v string) *CreateFlowTagGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowTagGroupRequest) Validate() error {
	return dara.Validate(s)
}
