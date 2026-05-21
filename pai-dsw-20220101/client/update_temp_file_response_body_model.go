// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTempFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTempFileResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateTempFileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateTempFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTempFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTempFileResponseBody
	GetSuccess() *bool
	SetTempFileId(v string) *UpdateTempFileResponseBody
	GetTempFileId() *string
}

type UpdateTempFileResponseBody struct {
	// example:
	//
	// "200"
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// tempfile-05cexxxxxxxxx
	TempFileId *string `json:"TempFileId,omitempty" xml:"TempFileId,omitempty"`
}

func (s UpdateTempFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTempFileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTempFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTempFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateTempFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTempFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTempFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTempFileResponseBody) GetTempFileId() *string {
	return s.TempFileId
}

func (s *UpdateTempFileResponseBody) SetCode(v string) *UpdateTempFileResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTempFileResponseBody) SetHttpStatusCode(v int32) *UpdateTempFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTempFileResponseBody) SetMessage(v string) *UpdateTempFileResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTempFileResponseBody) SetRequestId(v string) *UpdateTempFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTempFileResponseBody) SetSuccess(v bool) *UpdateTempFileResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTempFileResponseBody) SetTempFileId(v string) *UpdateTempFileResponseBody {
	s.TempFileId = &v
	return s
}

func (s *UpdateTempFileResponseBody) Validate() error {
	return dara.Validate(s)
}
