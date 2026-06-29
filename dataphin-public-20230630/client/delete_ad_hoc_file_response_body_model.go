// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdHocFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAdHocFileResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteAdHocFileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteAdHocFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAdHocFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAdHocFileResponseBody
	GetSuccess() *bool
}

type DeleteAdHocFileResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned by the backend.
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
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAdHocFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdHocFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAdHocFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAdHocFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteAdHocFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAdHocFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAdHocFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAdHocFileResponseBody) SetCode(v string) *DeleteAdHocFileResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAdHocFileResponseBody) SetHttpStatusCode(v int32) *DeleteAdHocFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAdHocFileResponseBody) SetMessage(v string) *DeleteAdHocFileResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAdHocFileResponseBody) SetRequestId(v string) *DeleteAdHocFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAdHocFileResponseBody) SetSuccess(v bool) *DeleteAdHocFileResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAdHocFileResponseBody) Validate() error {
	return dara.Validate(s)
}
