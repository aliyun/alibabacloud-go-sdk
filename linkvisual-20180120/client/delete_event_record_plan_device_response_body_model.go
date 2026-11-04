// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventRecordPlanDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteEventRecordPlanDeviceResponseBody
	GetCode() *string
	SetErrorMessage(v string) *DeleteEventRecordPlanDeviceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteEventRecordPlanDeviceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteEventRecordPlanDeviceResponseBody
	GetSuccess() *bool
}

type DeleteEventRecordPlanDeviceResponseBody struct {
	// The response code. Valid values:
	//
	// - **200**: The call is successful.
	//
	// - Other values: An error occurred. For more information about error codes, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// tenant auth error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06DC77A0-4622-42DB-9EE0-25FIOHS82JK1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEventRecordPlanDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventRecordPlanDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventRecordPlanDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteEventRecordPlanDeviceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteEventRecordPlanDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEventRecordPlanDeviceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteEventRecordPlanDeviceResponseBody) SetCode(v string) *DeleteEventRecordPlanDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceResponseBody) SetErrorMessage(v string) *DeleteEventRecordPlanDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceResponseBody) SetRequestId(v string) *DeleteEventRecordPlanDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceResponseBody) SetSuccess(v bool) *DeleteEventRecordPlanDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEventRecordPlanDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
