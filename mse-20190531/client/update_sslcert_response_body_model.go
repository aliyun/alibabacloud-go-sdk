// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSSLCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateSSLCertResponseBody
	GetCode() *int32
	SetData(v bool) *UpdateSSLCertResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpdateSSLCertResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateSSLCertResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSSLCertResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSSLCertResponseBody
	GetSuccess() *bool
}

type UpdateSSLCertResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the update is successful.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CA689BED-08F3-54C2-A206-A0924E2ACA0A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSSLCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSSLCertResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSSLCertResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateSSLCertResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateSSLCertResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateSSLCertResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSSLCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSSLCertResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSSLCertResponseBody) SetCode(v int32) *UpdateSSLCertResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSSLCertResponseBody) SetData(v bool) *UpdateSSLCertResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateSSLCertResponseBody) SetHttpStatusCode(v int32) *UpdateSSLCertResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateSSLCertResponseBody) SetMessage(v string) *UpdateSSLCertResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSSLCertResponseBody) SetRequestId(v string) *UpdateSSLCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSSLCertResponseBody) SetSuccess(v bool) *UpdateSSLCertResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSSLCertResponseBody) Validate() error {
	return dara.Validate(s)
}
