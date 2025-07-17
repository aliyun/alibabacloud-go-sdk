// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClaimAlarmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ClaimAlarmResponseBody
	GetCode() *int64
	SetMessage(v string) *ClaimAlarmResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClaimAlarmResponseBody
	GetRequestId() *string
	SetResult(v bool) *ClaimAlarmResponseBody
	GetResult() *bool
	SetSuccess(v bool) *ClaimAlarmResponseBody
	GetSuccess() *bool
}

type ClaimAlarmResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6A9AEA84-7186-4D8D-B498-4585C6A2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// `true`
	//
	// `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ClaimAlarmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClaimAlarmResponseBody) GoString() string {
	return s.String()
}

func (s *ClaimAlarmResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ClaimAlarmResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClaimAlarmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClaimAlarmResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ClaimAlarmResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ClaimAlarmResponseBody) SetCode(v int64) *ClaimAlarmResponseBody {
	s.Code = &v
	return s
}

func (s *ClaimAlarmResponseBody) SetMessage(v string) *ClaimAlarmResponseBody {
	s.Message = &v
	return s
}

func (s *ClaimAlarmResponseBody) SetRequestId(v string) *ClaimAlarmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClaimAlarmResponseBody) SetResult(v bool) *ClaimAlarmResponseBody {
	s.Result = &v
	return s
}

func (s *ClaimAlarmResponseBody) SetSuccess(v bool) *ClaimAlarmResponseBody {
	s.Success = &v
	return s
}

func (s *ClaimAlarmResponseBody) Validate() error {
	return dara.Validate(s)
}
