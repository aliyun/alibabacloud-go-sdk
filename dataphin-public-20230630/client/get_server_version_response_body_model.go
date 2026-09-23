// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServerVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetServerVersionResponseBody
	GetCode() *string
	SetData(v string) *GetServerVersionResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *GetServerVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetServerVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetServerVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetServerVersionResponseBody
	GetSuccess() *bool
}

type GetServerVersionResponseBody struct {
	// The error code. A value of OK indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The Dataphin server version number.
	//
	// example:
	//
	// 6.4.0.994114
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetServerVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServerVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetServerVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetServerVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *GetServerVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetServerVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetServerVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServerVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetServerVersionResponseBody) SetCode(v string) *GetServerVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetServerVersionResponseBody) SetData(v string) *GetServerVersionResponseBody {
	s.Data = &v
	return s
}

func (s *GetServerVersionResponseBody) SetHttpStatusCode(v int32) *GetServerVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetServerVersionResponseBody) SetMessage(v string) *GetServerVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetServerVersionResponseBody) SetRequestId(v string) *GetServerVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServerVersionResponseBody) SetSuccess(v bool) *GetServerVersionResponseBody {
	s.Success = &v
	return s
}

func (s *GetServerVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
