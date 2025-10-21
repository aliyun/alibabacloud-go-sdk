// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterCustomConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Connector) *RegisterCustomConnectorResponseBody
	GetData() []*Connector
	SetErrorCode(v string) *RegisterCustomConnectorResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RegisterCustomConnectorResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *RegisterCustomConnectorResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *RegisterCustomConnectorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RegisterCustomConnectorResponseBody
	GetSuccess() *bool
}

type RegisterCustomConnectorResponseBody struct {
	// If the value of success was true, a list of deployments in which custom connectors were deleted was returned. If the value of success was false, a null value was returned.
	Data []*Connector `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RegisterCustomConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterCustomConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterCustomConnectorResponseBody) GetData() []*Connector {
	return s.Data
}

func (s *RegisterCustomConnectorResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RegisterCustomConnectorResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RegisterCustomConnectorResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *RegisterCustomConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterCustomConnectorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RegisterCustomConnectorResponseBody) SetData(v []*Connector) *RegisterCustomConnectorResponseBody {
	s.Data = v
	return s
}

func (s *RegisterCustomConnectorResponseBody) SetErrorCode(v string) *RegisterCustomConnectorResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RegisterCustomConnectorResponseBody) SetErrorMessage(v string) *RegisterCustomConnectorResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RegisterCustomConnectorResponseBody) SetHttpCode(v int32) *RegisterCustomConnectorResponseBody {
	s.HttpCode = &v
	return s
}

func (s *RegisterCustomConnectorResponseBody) SetRequestId(v string) *RegisterCustomConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterCustomConnectorResponseBody) SetSuccess(v bool) *RegisterCustomConnectorResponseBody {
	s.Success = &v
	return s
}

func (s *RegisterCustomConnectorResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
