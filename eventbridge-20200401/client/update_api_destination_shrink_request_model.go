// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiDestinationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiDestinationName(v string) *UpdateApiDestinationShrinkRequest
	GetApiDestinationName() *string
	SetConnectionName(v string) *UpdateApiDestinationShrinkRequest
	GetConnectionName() *string
	SetDescription(v string) *UpdateApiDestinationShrinkRequest
	GetDescription() *string
	SetHttpApiParametersShrink(v string) *UpdateApiDestinationShrinkRequest
	GetHttpApiParametersShrink() *string
}

type UpdateApiDestinationShrinkRequest struct {
	// The name of the API destination. The name must be 2 to 127 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// api-destination-name
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	// The name of the connection. The name must be 2 to 127 characters in length.
	//
	// Note: Before you configure this parameter, you must call the CreateConnection operation to create a connection. Then, set this parameter to the name of the connection that you created.
	//
	// example:
	//
	// connection-name
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the API destination. The description can be up to 255 characters in length.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters that are configured for the API destination.
	HttpApiParametersShrink *string `json:"HttpApiParameters,omitempty" xml:"HttpApiParameters,omitempty"`
}

func (s UpdateApiDestinationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiDestinationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateApiDestinationShrinkRequest) GetApiDestinationName() *string {
	return s.ApiDestinationName
}

func (s *UpdateApiDestinationShrinkRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *UpdateApiDestinationShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApiDestinationShrinkRequest) GetHttpApiParametersShrink() *string {
	return s.HttpApiParametersShrink
}

func (s *UpdateApiDestinationShrinkRequest) SetApiDestinationName(v string) *UpdateApiDestinationShrinkRequest {
	s.ApiDestinationName = &v
	return s
}

func (s *UpdateApiDestinationShrinkRequest) SetConnectionName(v string) *UpdateApiDestinationShrinkRequest {
	s.ConnectionName = &v
	return s
}

func (s *UpdateApiDestinationShrinkRequest) SetDescription(v string) *UpdateApiDestinationShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateApiDestinationShrinkRequest) SetHttpApiParametersShrink(v string) *UpdateApiDestinationShrinkRequest {
	s.HttpApiParametersShrink = &v
	return s
}

func (s *UpdateApiDestinationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
