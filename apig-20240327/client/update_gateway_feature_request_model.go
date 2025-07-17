// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayFeatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetValue(v string) *UpdateGatewayFeatureRequest
	GetValue() *string
}

type UpdateGatewayFeatureRequest struct {
	// Parameter value.
	//
	// example:
	//
	// "true"
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateGatewayFeatureRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayFeatureRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayFeatureRequest) GetValue() *string {
	return s.Value
}

func (s *UpdateGatewayFeatureRequest) SetValue(v string) *UpdateGatewayFeatureRequest {
	s.Value = &v
	return s
}

func (s *UpdateGatewayFeatureRequest) Validate() error {
	return dara.Validate(s)
}
