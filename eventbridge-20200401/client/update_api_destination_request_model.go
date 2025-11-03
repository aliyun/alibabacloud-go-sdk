// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiDestinationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiDestinationName(v string) *UpdateApiDestinationRequest
	GetApiDestinationName() *string
	SetConnectionName(v string) *UpdateApiDestinationRequest
	GetConnectionName() *string
	SetDescription(v string) *UpdateApiDestinationRequest
	GetDescription() *string
	SetHttpApiParameters(v *UpdateApiDestinationRequestHttpApiParameters) *UpdateApiDestinationRequest
	GetHttpApiParameters() *UpdateApiDestinationRequestHttpApiParameters
}

type UpdateApiDestinationRequest struct {
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
	HttpApiParameters *UpdateApiDestinationRequestHttpApiParameters `json:"HttpApiParameters,omitempty" xml:"HttpApiParameters,omitempty" type:"Struct"`
}

func (s UpdateApiDestinationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiDestinationRequest) GoString() string {
	return s.String()
}

func (s *UpdateApiDestinationRequest) GetApiDestinationName() *string {
	return s.ApiDestinationName
}

func (s *UpdateApiDestinationRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *UpdateApiDestinationRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApiDestinationRequest) GetHttpApiParameters() *UpdateApiDestinationRequestHttpApiParameters {
	return s.HttpApiParameters
}

func (s *UpdateApiDestinationRequest) SetApiDestinationName(v string) *UpdateApiDestinationRequest {
	s.ApiDestinationName = &v
	return s
}

func (s *UpdateApiDestinationRequest) SetConnectionName(v string) *UpdateApiDestinationRequest {
	s.ConnectionName = &v
	return s
}

func (s *UpdateApiDestinationRequest) SetDescription(v string) *UpdateApiDestinationRequest {
	s.Description = &v
	return s
}

func (s *UpdateApiDestinationRequest) SetHttpApiParameters(v *UpdateApiDestinationRequestHttpApiParameters) *UpdateApiDestinationRequest {
	s.HttpApiParameters = v
	return s
}

func (s *UpdateApiDestinationRequest) Validate() error {
	if s.HttpApiParameters != nil {
		if err := s.HttpApiParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApiDestinationRequestHttpApiParameters struct {
	// The endpoint of the API destination. The endpoint can be up to 127 characters in length.
	//
	// example:
	//
	// http://127.0.0.1:8001/api
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The HTTP request method. Valid values:
	//
	// - GET
	//
	// - POST
	//
	// - HEAD
	//
	// - DELETE
	//
	// - PUT
	//
	// - PATCH
	//
	// example:
	//
	// POST
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s UpdateApiDestinationRequestHttpApiParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiDestinationRequestHttpApiParameters) GoString() string {
	return s.String()
}

func (s *UpdateApiDestinationRequestHttpApiParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateApiDestinationRequestHttpApiParameters) GetMethod() *string {
	return s.Method
}

func (s *UpdateApiDestinationRequestHttpApiParameters) SetEndpoint(v string) *UpdateApiDestinationRequestHttpApiParameters {
	s.Endpoint = &v
	return s
}

func (s *UpdateApiDestinationRequestHttpApiParameters) SetMethod(v string) *UpdateApiDestinationRequestHttpApiParameters {
	s.Method = &v
	return s
}

func (s *UpdateApiDestinationRequestHttpApiParameters) Validate() error {
	return dara.Validate(s)
}
