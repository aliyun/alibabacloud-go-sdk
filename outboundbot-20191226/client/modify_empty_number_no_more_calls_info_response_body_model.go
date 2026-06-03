// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEmptyNumberNoMoreCallsInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyEmptyNumberNoMoreCallsInfoResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyEmptyNumberNoMoreCallsInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyEmptyNumberNoMoreCallsInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyEmptyNumberNoMoreCallsInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyEmptyNumberNoMoreCallsInfoResponseBody
	GetSuccess() *bool
}

type ModifyEmptyNumberNoMoreCallsInfoResponseBody struct {
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

func (s ModifyEmptyNumberNoMoreCallsInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmptyNumberNoMoreCallsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponseBody) SetCode(v string) *ModifyEmptyNumberNoMoreCallsInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponseBody) SetHttpStatusCode(v int32) *ModifyEmptyNumberNoMoreCallsInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponseBody) SetMessage(v string) *ModifyEmptyNumberNoMoreCallsInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponseBody) SetRequestId(v string) *ModifyEmptyNumberNoMoreCallsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponseBody) SetSuccess(v bool) *ModifyEmptyNumberNoMoreCallsInfoResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyEmptyNumberNoMoreCallsInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
