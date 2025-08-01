// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSSLCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddSSLCertResponseBody
	GetCode() *int32
	SetData(v bool) *AddSSLCertResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *AddSSLCertResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddSSLCertResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddSSLCertResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddSSLCertResponseBody
	GetSuccess() *bool
}

type AddSSLCertResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the association is successful.
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
	// The error message.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E1CC74F0-5BDE-5220-A031-5CA622D80723
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

func (s AddSSLCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSSLCertResponseBody) GoString() string {
	return s.String()
}

func (s *AddSSLCertResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddSSLCertResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddSSLCertResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddSSLCertResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddSSLCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSSLCertResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddSSLCertResponseBody) SetCode(v int32) *AddSSLCertResponseBody {
	s.Code = &v
	return s
}

func (s *AddSSLCertResponseBody) SetData(v bool) *AddSSLCertResponseBody {
	s.Data = &v
	return s
}

func (s *AddSSLCertResponseBody) SetHttpStatusCode(v int32) *AddSSLCertResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddSSLCertResponseBody) SetMessage(v string) *AddSSLCertResponseBody {
	s.Message = &v
	return s
}

func (s *AddSSLCertResponseBody) SetRequestId(v string) *AddSSLCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSSLCertResponseBody) SetSuccess(v bool) *AddSSLCertResponseBody {
	s.Success = &v
	return s
}

func (s *AddSSLCertResponseBody) Validate() error {
	return dara.Validate(s)
}
