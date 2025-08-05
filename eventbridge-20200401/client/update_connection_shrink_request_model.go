// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthParametersShrink(v string) *UpdateConnectionShrinkRequest
	GetAuthParametersShrink() *string
	SetConnectionName(v string) *UpdateConnectionShrinkRequest
	GetConnectionName() *string
	SetDescription(v string) *UpdateConnectionShrinkRequest
	GetDescription() *string
	SetNetworkParametersShrink(v string) *UpdateConnectionShrinkRequest
	GetNetworkParametersShrink() *string
}

type UpdateConnectionShrinkRequest struct {
	// The parameters that are configured for authentication.
	AuthParametersShrink *string `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty"`
	// The name of the connection that you want to update. The name must be 2 to 127 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// connection-name
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the connection. The description can be up to 255 characters in length.
	//
	// example:
	//
	// The description of the connection.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters that are configured for the network.
	//
	// This parameter is required.
	NetworkParametersShrink *string `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty"`
}

func (s UpdateConnectionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnectionShrinkRequest) GetAuthParametersShrink() *string {
	return s.AuthParametersShrink
}

func (s *UpdateConnectionShrinkRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *UpdateConnectionShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateConnectionShrinkRequest) GetNetworkParametersShrink() *string {
	return s.NetworkParametersShrink
}

func (s *UpdateConnectionShrinkRequest) SetAuthParametersShrink(v string) *UpdateConnectionShrinkRequest {
	s.AuthParametersShrink = &v
	return s
}

func (s *UpdateConnectionShrinkRequest) SetConnectionName(v string) *UpdateConnectionShrinkRequest {
	s.ConnectionName = &v
	return s
}

func (s *UpdateConnectionShrinkRequest) SetDescription(v string) *UpdateConnectionShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateConnectionShrinkRequest) SetNetworkParametersShrink(v string) *UpdateConnectionShrinkRequest {
	s.NetworkParametersShrink = &v
	return s
}

func (s *UpdateConnectionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
