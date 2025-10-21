// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomConnectorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Connector) *ListCustomConnectorsResponseBody
	GetData() []*Connector
	SetErrorCode(v string) *ListCustomConnectorsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListCustomConnectorsResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ListCustomConnectorsResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListCustomConnectorsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCustomConnectorsResponseBody
	GetSuccess() *bool
}

type ListCustomConnectorsResponseBody struct {
	// If the value of success was true, the list of custom connectors in the namespace was returned. If the value of success was false, a null value was returned.
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

func (s ListCustomConnectorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomConnectorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomConnectorsResponseBody) GetData() []*Connector {
	return s.Data
}

func (s *ListCustomConnectorsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListCustomConnectorsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListCustomConnectorsResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListCustomConnectorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomConnectorsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCustomConnectorsResponseBody) SetData(v []*Connector) *ListCustomConnectorsResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomConnectorsResponseBody) SetErrorCode(v string) *ListCustomConnectorsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListCustomConnectorsResponseBody) SetErrorMessage(v string) *ListCustomConnectorsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListCustomConnectorsResponseBody) SetHttpCode(v int32) *ListCustomConnectorsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListCustomConnectorsResponseBody) SetRequestId(v string) *ListCustomConnectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomConnectorsResponseBody) SetSuccess(v bool) *ListCustomConnectorsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCustomConnectorsResponseBody) Validate() error {
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
