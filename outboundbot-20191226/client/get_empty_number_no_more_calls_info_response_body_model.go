// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmptyNumberNoMoreCallsInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEmptyNumberNoMoreCallsInfoResponseBody
	GetCode() *string
	SetEmptyNumberNoMoreCalls(v bool) *GetEmptyNumberNoMoreCallsInfoResponseBody
	GetEmptyNumberNoMoreCalls() *bool
	SetHttpStatusCode(v int32) *GetEmptyNumberNoMoreCallsInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetEmptyNumberNoMoreCallsInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEmptyNumberNoMoreCallsInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetEmptyNumberNoMoreCallsInfoResponseBody
	GetSuccess() *bool
}

type GetEmptyNumberNoMoreCallsInfoResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	EmptyNumberNoMoreCalls *bool `json:"EmptyNumberNoMoreCalls,omitempty" xml:"EmptyNumberNoMoreCalls,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEmptyNumberNoMoreCallsInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEmptyNumberNoMoreCallsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmptyNumberNoMoreCallsInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEmptyNumberNoMoreCallsInfoResponseBody) GetEmptyNumberNoMoreCalls() *bool {
	return s.EmptyNumberNoMoreCalls
}

func (s *GetEmptyNumberNoMoreCallsInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetEmptyNumberNoMoreCallsInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEmptyNumberNoMoreCallsInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEmptyNumberNoMoreCallsInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEmptyNumberNoMoreCallsInfoResponseBody) SetCode(v string) *GetEmptyNumberNoMoreCallsInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetEmptyNumberNoMoreCallsInfoResponseBody) SetEmptyNumberNoMoreCalls(v bool) *GetEmptyNumberNoMoreCallsInfoResponseBody {
	s.EmptyNumberNoMoreCalls = &v
	return s
}

func (s *GetEmptyNumberNoMoreCallsInfoResponseBody) SetHttpStatusCode(v int32) *GetEmptyNumberNoMoreCallsInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetEmptyNumberNoMoreCallsInfoResponseBody) SetMessage(v string) *GetEmptyNumberNoMoreCallsInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetEmptyNumberNoMoreCallsInfoResponseBody) SetRequestId(v string) *GetEmptyNumberNoMoreCallsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEmptyNumberNoMoreCallsInfoResponseBody) SetSuccess(v bool) *GetEmptyNumberNoMoreCallsInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetEmptyNumberNoMoreCallsInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
