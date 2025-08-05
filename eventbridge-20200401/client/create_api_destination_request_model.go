// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiDestinationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiDestinationName(v string) *CreateApiDestinationRequest
	GetApiDestinationName() *string
	SetConnectionName(v string) *CreateApiDestinationRequest
	GetConnectionName() *string
	SetDescription(v string) *CreateApiDestinationRequest
	GetDescription() *string
	SetHttpApiParameters(v *CreateApiDestinationRequestHttpApiParameters) *CreateApiDestinationRequest
	GetHttpApiParameters() *CreateApiDestinationRequestHttpApiParameters
}

type CreateApiDestinationRequest struct {
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
	HttpApiParameters *CreateApiDestinationRequestHttpApiParameters `json:"HttpApiParameters,omitempty" xml:"HttpApiParameters,omitempty" type:"Struct"`
}

func (s CreateApiDestinationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApiDestinationRequest) GoString() string {
	return s.String()
}

func (s *CreateApiDestinationRequest) GetApiDestinationName() *string {
	return s.ApiDestinationName
}

func (s *CreateApiDestinationRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *CreateApiDestinationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApiDestinationRequest) GetHttpApiParameters() *CreateApiDestinationRequestHttpApiParameters {
	return s.HttpApiParameters
}

func (s *CreateApiDestinationRequest) SetApiDestinationName(v string) *CreateApiDestinationRequest {
	s.ApiDestinationName = &v
	return s
}

func (s *CreateApiDestinationRequest) SetConnectionName(v string) *CreateApiDestinationRequest {
	s.ConnectionName = &v
	return s
}

func (s *CreateApiDestinationRequest) SetDescription(v string) *CreateApiDestinationRequest {
	s.Description = &v
	return s
}

func (s *CreateApiDestinationRequest) SetHttpApiParameters(v *CreateApiDestinationRequestHttpApiParameters) *CreateApiDestinationRequest {
	s.HttpApiParameters = v
	return s
}

func (s *CreateApiDestinationRequest) Validate() error {
	return dara.Validate(s)
}

type CreateApiDestinationRequestHttpApiParameters struct {
	// The endpoint of the API destination. The endpoint can be up to 127 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://127.0.0.1:8001/api
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The HTTP request method. Valid values:
	//
	// 	- GET
	//
	// 	- POST
	//
	// 	- HEAD
	//
	// 	- DELETE
	//
	// 	- PUT
	//
	// 	- PATCH
	//
	// This parameter is required.
	//
	// example:
	//
	// POST
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s CreateApiDestinationRequestHttpApiParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateApiDestinationRequestHttpApiParameters) GoString() string {
	return s.String()
}

func (s *CreateApiDestinationRequestHttpApiParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateApiDestinationRequestHttpApiParameters) GetMethod() *string {
	return s.Method
}

func (s *CreateApiDestinationRequestHttpApiParameters) SetEndpoint(v string) *CreateApiDestinationRequestHttpApiParameters {
	s.Endpoint = &v
	return s
}

func (s *CreateApiDestinationRequestHttpApiParameters) SetMethod(v string) *CreateApiDestinationRequestHttpApiParameters {
	s.Method = &v
	return s
}

func (s *CreateApiDestinationRequestHttpApiParameters) Validate() error {
	return dara.Validate(s)
}
