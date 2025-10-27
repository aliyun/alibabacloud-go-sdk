// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteFlowRequest
	GetName() *string
}

type DeleteFlowRequest struct {
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowRequest) GetName() *string {
	return s.Name
}

func (s *DeleteFlowRequest) SetName(v string) *DeleteFlowRequest {
	s.Name = &v
	return s
}

func (s *DeleteFlowRequest) Validate() error {
	return dara.Validate(s)
}
