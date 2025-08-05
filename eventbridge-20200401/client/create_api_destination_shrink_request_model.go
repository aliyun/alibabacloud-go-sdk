// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiDestinationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiDestinationName(v string) *CreateApiDestinationShrinkRequest
	GetApiDestinationName() *string
	SetConnectionName(v string) *CreateApiDestinationShrinkRequest
	GetConnectionName() *string
	SetDescription(v string) *CreateApiDestinationShrinkRequest
	GetDescription() *string
	SetHttpApiParametersShrink(v string) *CreateApiDestinationShrinkRequest
	GetHttpApiParametersShrink() *string
}

type CreateApiDestinationShrinkRequest struct {
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
	// >
	//
	// >  Before you configure this parameter, you must call the CreateConnection operation to create a connection. Then, set this parameter to the name of the connection that you created.
	//
	// This parameter is required.
	//
	// example:
	//
	// connection-name
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the API destination. The description can be up to 255 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters that are configured for the API destination.
	//
	// This parameter is required.
	HttpApiParametersShrink *string `json:"HttpApiParameters,omitempty" xml:"HttpApiParameters,omitempty"`
}

func (s CreateApiDestinationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApiDestinationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApiDestinationShrinkRequest) GetApiDestinationName() *string {
	return s.ApiDestinationName
}

func (s *CreateApiDestinationShrinkRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *CreateApiDestinationShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApiDestinationShrinkRequest) GetHttpApiParametersShrink() *string {
	return s.HttpApiParametersShrink
}

func (s *CreateApiDestinationShrinkRequest) SetApiDestinationName(v string) *CreateApiDestinationShrinkRequest {
	s.ApiDestinationName = &v
	return s
}

func (s *CreateApiDestinationShrinkRequest) SetConnectionName(v string) *CreateApiDestinationShrinkRequest {
	s.ConnectionName = &v
	return s
}

func (s *CreateApiDestinationShrinkRequest) SetDescription(v string) *CreateApiDestinationShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateApiDestinationShrinkRequest) SetHttpApiParametersShrink(v string) *CreateApiDestinationShrinkRequest {
	s.HttpApiParametersShrink = &v
	return s
}

func (s *CreateApiDestinationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
