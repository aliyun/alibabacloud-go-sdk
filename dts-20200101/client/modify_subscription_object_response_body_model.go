// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySubscriptionObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ModifySubscriptionObjectResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifySubscriptionObjectResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ModifySubscriptionObjectResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifySubscriptionObjectResponseBody
	GetSuccess() *string
}

type ModifySubscriptionObjectResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABBACEFC-CBA9-4F80-A337-42F202F5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySubscriptionObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySubscriptionObjectResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionObjectResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifySubscriptionObjectResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifySubscriptionObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySubscriptionObjectResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifySubscriptionObjectResponseBody) SetErrCode(v string) *ModifySubscriptionObjectResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifySubscriptionObjectResponseBody) SetErrMessage(v string) *ModifySubscriptionObjectResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifySubscriptionObjectResponseBody) SetRequestId(v string) *ModifySubscriptionObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySubscriptionObjectResponseBody) SetSuccess(v string) *ModifySubscriptionObjectResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySubscriptionObjectResponseBody) Validate() error {
	return dara.Validate(s)
}
