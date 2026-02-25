// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSSLConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DisableSSLConnectionResponseBody
	GetData() *bool
	SetErrCode(v string) *DisableSSLConnectionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DisableSSLConnectionResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DisableSSLConnectionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DisableSSLConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableSSLConnectionResponseBody
	GetSuccess() *bool
}

type DisableSSLConnectionResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableSSLConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSSLConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSSLConnectionResponseBody) GetData() *bool {
	return s.Data
}

func (s *DisableSSLConnectionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DisableSSLConnectionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DisableSSLConnectionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DisableSSLConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSSLConnectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableSSLConnectionResponseBody) SetData(v bool) *DisableSSLConnectionResponseBody {
	s.Data = &v
	return s
}

func (s *DisableSSLConnectionResponseBody) SetErrCode(v string) *DisableSSLConnectionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DisableSSLConnectionResponseBody) SetErrMessage(v string) *DisableSSLConnectionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DisableSSLConnectionResponseBody) SetHttpStatusCode(v int32) *DisableSSLConnectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableSSLConnectionResponseBody) SetRequestId(v string) *DisableSSLConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSSLConnectionResponseBody) SetSuccess(v bool) *DisableSSLConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *DisableSSLConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
