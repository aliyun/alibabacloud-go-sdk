// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConsumptionTimestampResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ModifyConsumptionTimestampResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyConsumptionTimestampResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ModifyConsumptionTimestampResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyConsumptionTimestampResponseBody
	GetSuccess() *string
}

type ModifyConsumptionTimestampResponseBody struct {
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

func (s ModifyConsumptionTimestampResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyConsumptionTimestampResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyConsumptionTimestampResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyConsumptionTimestampResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyConsumptionTimestampResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyConsumptionTimestampResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyConsumptionTimestampResponseBody) SetErrCode(v string) *ModifyConsumptionTimestampResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyConsumptionTimestampResponseBody) SetErrMessage(v string) *ModifyConsumptionTimestampResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyConsumptionTimestampResponseBody) SetRequestId(v string) *ModifyConsumptionTimestampResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyConsumptionTimestampResponseBody) SetSuccess(v string) *ModifyConsumptionTimestampResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyConsumptionTimestampResponseBody) Validate() error {
	return dara.Validate(s)
}
