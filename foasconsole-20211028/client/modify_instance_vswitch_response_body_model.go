// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVswitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceVswitchResponseBody
	GetRequestId() *string
	SetResult(v string) *ModifyInstanceVswitchResponseBody
	GetResult() *string
	SetSuccess(v bool) *ModifyInstanceVswitchResponseBody
	GetSuccess() *bool
}

type ModifyInstanceVswitchResponseBody struct {
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// F2C5B6A8-DD04-51F5-AAD5-BA2FE6FD****
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyInstanceVswitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVswitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVswitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceVswitchResponseBody) GetResult() *string {
	return s.Result
}

func (s *ModifyInstanceVswitchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyInstanceVswitchResponseBody) SetRequestId(v string) *ModifyInstanceVswitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceVswitchResponseBody) SetResult(v string) *ModifyInstanceVswitchResponseBody {
	s.Result = &v
	return s
}

func (s *ModifyInstanceVswitchResponseBody) SetSuccess(v bool) *ModifyInstanceVswitchResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyInstanceVswitchResponseBody) Validate() error {
	return dara.Validate(s)
}
