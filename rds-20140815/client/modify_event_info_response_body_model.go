// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEventInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ModifyEventInfoResponseBody
	GetErrorCode() *string
	SetErrorEventId(v string) *ModifyEventInfoResponseBody
	GetErrorEventId() *string
	SetRequestId(v string) *ModifyEventInfoResponseBody
	GetRequestId() *string
	SetSuccessCount(v int32) *ModifyEventInfoResponseBody
	GetSuccessCount() *int32
	SetSuccessEventId(v string) *ModifyEventInfoResponseBody
	GetSuccessEventId() *string
}

type ModifyEventInfoResponseBody struct {
	// The error code.
	//
	// example:
	//
	// mst.errorcode.success.errormessage
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error ID.
	//
	// example:
	//
	// 12343
	ErrorEventId *string `json:"ErrorEventId,omitempty" xml:"ErrorEventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 224DB9F7-3100-4899-AB9C-C938BCCB43E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of successful records.
	//
	// example:
	//
	// 1
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The ID of the successful event.
	//
	// example:
	//
	// 234221
	SuccessEventId *string `json:"SuccessEventId,omitempty" xml:"SuccessEventId,omitempty"`
}

func (s ModifyEventInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEventInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEventInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ModifyEventInfoResponseBody) GetErrorEventId() *string {
	return s.ErrorEventId
}

func (s *ModifyEventInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEventInfoResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *ModifyEventInfoResponseBody) GetSuccessEventId() *string {
	return s.SuccessEventId
}

func (s *ModifyEventInfoResponseBody) SetErrorCode(v string) *ModifyEventInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ModifyEventInfoResponseBody) SetErrorEventId(v string) *ModifyEventInfoResponseBody {
	s.ErrorEventId = &v
	return s
}

func (s *ModifyEventInfoResponseBody) SetRequestId(v string) *ModifyEventInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEventInfoResponseBody) SetSuccessCount(v int32) *ModifyEventInfoResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *ModifyEventInfoResponseBody) SetSuccessEventId(v string) *ModifyEventInfoResponseBody {
	s.SuccessEventId = &v
	return s
}

func (s *ModifyEventInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
