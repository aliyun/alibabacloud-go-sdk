// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiDestinationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetApiDestinationResponseBody
	GetCode() *string
	SetData(v *GetApiDestinationResponseBodyData) *GetApiDestinationResponseBody
	GetData() *GetApiDestinationResponseBodyData
	SetMessage(v string) *GetApiDestinationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetApiDestinationResponseBody
	GetRequestId() *string
}

type GetApiDestinationResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetApiDestinationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message. If the request is successful, success is returned. If the request failed, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B896B484-F16D-59DE-9E23-DD0E5C361108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApiDestinationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApiDestinationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApiDestinationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetApiDestinationResponseBody) GetData() *GetApiDestinationResponseBodyData {
	return s.Data
}

func (s *GetApiDestinationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApiDestinationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApiDestinationResponseBody) SetCode(v string) *GetApiDestinationResponseBody {
	s.Code = &v
	return s
}

func (s *GetApiDestinationResponseBody) SetData(v *GetApiDestinationResponseBodyData) *GetApiDestinationResponseBody {
	s.Data = v
	return s
}

func (s *GetApiDestinationResponseBody) SetMessage(v string) *GetApiDestinationResponseBody {
	s.Message = &v
	return s
}

func (s *GetApiDestinationResponseBody) SetRequestId(v string) *GetApiDestinationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApiDestinationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApiDestinationResponseBodyData struct {
	// The name of the API destination.
	//
	// example:
	//
	// demo
	ApiDestinationName *string `json:"ApiDestinationName,omitempty" xml:"ApiDestinationName,omitempty"`
	// The connection name.
	//
	// example:
	//
	// test-basic
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the API destination.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the API destination was created.
	//
	// example:
	//
	// 1649055710565
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The request parameters that are configured for the API destination.
	HttpApiParameters *GetApiDestinationResponseBodyDataHttpApiParameters `json:"HttpApiParameters,omitempty" xml:"HttpApiParameters,omitempty" type:"Struct"`
}

func (s GetApiDestinationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetApiDestinationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApiDestinationResponseBodyData) GetApiDestinationName() *string {
	return s.ApiDestinationName
}

func (s *GetApiDestinationResponseBodyData) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *GetApiDestinationResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetApiDestinationResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetApiDestinationResponseBodyData) GetHttpApiParameters() *GetApiDestinationResponseBodyDataHttpApiParameters {
	return s.HttpApiParameters
}

func (s *GetApiDestinationResponseBodyData) SetApiDestinationName(v string) *GetApiDestinationResponseBodyData {
	s.ApiDestinationName = &v
	return s
}

func (s *GetApiDestinationResponseBodyData) SetConnectionName(v string) *GetApiDestinationResponseBodyData {
	s.ConnectionName = &v
	return s
}

func (s *GetApiDestinationResponseBodyData) SetDescription(v string) *GetApiDestinationResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetApiDestinationResponseBodyData) SetGmtCreate(v int64) *GetApiDestinationResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetApiDestinationResponseBodyData) SetHttpApiParameters(v *GetApiDestinationResponseBodyDataHttpApiParameters) *GetApiDestinationResponseBodyData {
	s.HttpApiParameters = v
	return s
}

func (s *GetApiDestinationResponseBodyData) Validate() error {
	if s.HttpApiParameters != nil {
		if err := s.HttpApiParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApiDestinationResponseBodyDataHttpApiParameters struct {
	// The endpoint of the API destination.
	//
	// example:
	//
	// http://127.0.0.1:8001/api
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The HTTP request method. Valid values:
	//
	// 	- POST
	//
	// 	- GET
	//
	// 	- DELETE
	//
	// 	- PUT
	//
	// 	- HEAD
	//
	// 	- TRACE
	//
	// 	- PATCH
	//
	// example:
	//
	// POST
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s GetApiDestinationResponseBodyDataHttpApiParameters) String() string {
	return dara.Prettify(s)
}

func (s GetApiDestinationResponseBodyDataHttpApiParameters) GoString() string {
	return s.String()
}

func (s *GetApiDestinationResponseBodyDataHttpApiParameters) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetApiDestinationResponseBodyDataHttpApiParameters) GetMethod() *string {
	return s.Method
}

func (s *GetApiDestinationResponseBodyDataHttpApiParameters) SetEndpoint(v string) *GetApiDestinationResponseBodyDataHttpApiParameters {
	s.Endpoint = &v
	return s
}

func (s *GetApiDestinationResponseBodyDataHttpApiParameters) SetMethod(v string) *GetApiDestinationResponseBodyDataHttpApiParameters {
	s.Method = &v
	return s
}

func (s *GetApiDestinationResponseBodyDataHttpApiParameters) Validate() error {
	return dara.Validate(s)
}
