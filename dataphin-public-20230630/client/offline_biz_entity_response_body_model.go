// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineBizEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OfflineBizEntityResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *OfflineBizEntityResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *OfflineBizEntityResponseBody
	GetMessage() *string
	SetRequestId(v string) *OfflineBizEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OfflineBizEntityResponseBody
	GetSuccess() *bool
}

type OfflineBizEntityResponseBody struct {
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

func (s OfflineBizEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineBizEntityResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineBizEntityResponseBody) GetCode() *string {
	return s.Code
}

func (s *OfflineBizEntityResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *OfflineBizEntityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OfflineBizEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflineBizEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OfflineBizEntityResponseBody) SetCode(v string) *OfflineBizEntityResponseBody {
	s.Code = &v
	return s
}

func (s *OfflineBizEntityResponseBody) SetHttpStatusCode(v int32) *OfflineBizEntityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *OfflineBizEntityResponseBody) SetMessage(v string) *OfflineBizEntityResponseBody {
	s.Message = &v
	return s
}

func (s *OfflineBizEntityResponseBody) SetRequestId(v string) *OfflineBizEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineBizEntityResponseBody) SetSuccess(v bool) *OfflineBizEntityResponseBody {
	s.Success = &v
	return s
}

func (s *OfflineBizEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
