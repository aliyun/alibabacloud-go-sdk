// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTempFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTempFileResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateTempFileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateTempFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTempFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTempFileResponseBody
	GetSuccess() *bool
	SetTempFileId(v string) *CreateTempFileResponseBody
	GetTempFileId() *string
}

type CreateTempFileResponseBody struct {
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

func (s CreateTempFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTempFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTempFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTempFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateTempFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTempFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTempFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTempFileResponseBody) GetTempFileId() *string {
	return s.TempFileId
}

func (s *CreateTempFileResponseBody) SetCode(v string) *CreateTempFileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTempFileResponseBody) SetHttpStatusCode(v int32) *CreateTempFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTempFileResponseBody) SetMessage(v string) *CreateTempFileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTempFileResponseBody) SetRequestId(v string) *CreateTempFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTempFileResponseBody) SetSuccess(v bool) *CreateTempFileResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTempFileResponseBody) SetTempFileId(v string) *CreateTempFileResponseBody {
	s.TempFileId = &v
	return s
}

func (s *CreateTempFileResponseBody) Validate() error {
	return dara.Validate(s)
}
