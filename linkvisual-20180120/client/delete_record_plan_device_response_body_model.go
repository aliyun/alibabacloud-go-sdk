// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecordPlanDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteRecordPlanDeviceResponseBody
	GetCode() *string
	SetErrorMessage(v string) *DeleteRecordPlanDeviceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteRecordPlanDeviceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteRecordPlanDeviceResponseBody
	GetSuccess() *bool
}

type DeleteRecordPlanDeviceResponseBody struct {
	// The error code returned if the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
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
	// The ID of the request.
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

func (s DeleteRecordPlanDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecordPlanDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecordPlanDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRecordPlanDeviceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteRecordPlanDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRecordPlanDeviceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRecordPlanDeviceResponseBody) SetCode(v string) *DeleteRecordPlanDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRecordPlanDeviceResponseBody) SetErrorMessage(v string) *DeleteRecordPlanDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRecordPlanDeviceResponseBody) SetRequestId(v string) *DeleteRecordPlanDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecordPlanDeviceResponseBody) SetSuccess(v bool) *DeleteRecordPlanDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRecordPlanDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
