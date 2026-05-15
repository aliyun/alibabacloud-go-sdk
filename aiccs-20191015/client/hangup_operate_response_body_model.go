// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHangupOperateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HangupOperateResponseBody
	GetCode() *string
	SetMesage(v string) *HangupOperateResponseBody
	GetMesage() *string
	SetRequestId(v string) *HangupOperateResponseBody
	GetRequestId() *string
	SetResult(v bool) *HangupOperateResponseBody
	GetResult() *bool
}

type HangupOperateResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	Mesage *string `json:"Mesage,omitempty" xml:"Mesage,omitempty"`
	// example:
	//
	// EFD543DD-E087-54A2-AC0B-54E0656511D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s HangupOperateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HangupOperateResponseBody) GoString() string {
	return s.String()
}

func (s *HangupOperateResponseBody) GetCode() *string {
	return s.Code
}

func (s *HangupOperateResponseBody) GetMesage() *string {
	return s.Mesage
}

func (s *HangupOperateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HangupOperateResponseBody) GetResult() *bool {
	return s.Result
}

func (s *HangupOperateResponseBody) SetCode(v string) *HangupOperateResponseBody {
	s.Code = &v
	return s
}

func (s *HangupOperateResponseBody) SetMesage(v string) *HangupOperateResponseBody {
	s.Mesage = &v
	return s
}

func (s *HangupOperateResponseBody) SetRequestId(v string) *HangupOperateResponseBody {
	s.RequestId = &v
	return s
}

func (s *HangupOperateResponseBody) SetResult(v bool) *HangupOperateResponseBody {
	s.Result = &v
	return s
}

func (s *HangupOperateResponseBody) Validate() error {
	return dara.Validate(s)
}
