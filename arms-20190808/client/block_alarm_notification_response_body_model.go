// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBlockAlarmNotificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *BlockAlarmNotificationResponseBody
	GetCode() *int64
	SetMessage(v string) *BlockAlarmNotificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *BlockAlarmNotificationResponseBody
	GetRequestId() *string
	SetResult(v bool) *BlockAlarmNotificationResponseBody
	GetResult() *bool
	SetSuccess(v bool) *BlockAlarmNotificationResponseBody
	GetSuccess() *bool
}

type BlockAlarmNotificationResponseBody struct {
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
	// 626037F5-FDEB-45B0-804C-B3C92797****
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
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BlockAlarmNotificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BlockAlarmNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *BlockAlarmNotificationResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *BlockAlarmNotificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BlockAlarmNotificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BlockAlarmNotificationResponseBody) GetResult() *bool {
	return s.Result
}

func (s *BlockAlarmNotificationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BlockAlarmNotificationResponseBody) SetCode(v int64) *BlockAlarmNotificationResponseBody {
	s.Code = &v
	return s
}

func (s *BlockAlarmNotificationResponseBody) SetMessage(v string) *BlockAlarmNotificationResponseBody {
	s.Message = &v
	return s
}

func (s *BlockAlarmNotificationResponseBody) SetRequestId(v string) *BlockAlarmNotificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *BlockAlarmNotificationResponseBody) SetResult(v bool) *BlockAlarmNotificationResponseBody {
	s.Result = &v
	return s
}

func (s *BlockAlarmNotificationResponseBody) SetSuccess(v bool) *BlockAlarmNotificationResponseBody {
	s.Success = &v
	return s
}

func (s *BlockAlarmNotificationResponseBody) Validate() error {
	return dara.Validate(s)
}
