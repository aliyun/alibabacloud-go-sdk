// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateGatewayNameRequest
	GetName() *string
}

type UpdateGatewayNameRequest struct {
	// Gateway name.
	//
	// example:
	//
	// dev-itemcenter-router
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateGatewayNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayNameRequest) GetName() *string {
	return s.Name
}

func (s *UpdateGatewayNameRequest) SetName(v string) *UpdateGatewayNameRequest {
	s.Name = &v
	return s
}

func (s *UpdateGatewayNameRequest) Validate() error {
	return dara.Validate(s)
}
