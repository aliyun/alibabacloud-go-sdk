// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineBizEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OnlineBizEntityResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *OnlineBizEntityResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *OnlineBizEntityResponseBody
	GetMessage() *string
	SetRequestId(v string) *OnlineBizEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OnlineBizEntityResponseBody
	GetSuccess() *bool
}

type OnlineBizEntityResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OnlineBizEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnlineBizEntityResponseBody) GoString() string {
	return s.String()
}

func (s *OnlineBizEntityResponseBody) GetCode() *string {
	return s.Code
}

func (s *OnlineBizEntityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *OnlineBizEntityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OnlineBizEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnlineBizEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OnlineBizEntityResponseBody) SetCode(v string) *OnlineBizEntityResponseBody {
	s.Code = &v
	return s
}

func (s *OnlineBizEntityResponseBody) SetHttpStatusCode(v int32) *OnlineBizEntityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *OnlineBizEntityResponseBody) SetMessage(v string) *OnlineBizEntityResponseBody {
	s.Message = &v
	return s
}

func (s *OnlineBizEntityResponseBody) SetRequestId(v string) *OnlineBizEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnlineBizEntityResponseBody) SetSuccess(v bool) *OnlineBizEntityResponseBody {
	s.Success = &v
	return s
}

func (s *OnlineBizEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
