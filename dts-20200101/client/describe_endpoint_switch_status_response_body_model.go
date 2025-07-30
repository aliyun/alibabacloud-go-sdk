// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndpointSwitchStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeEndpointSwitchStatusResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeEndpointSwitchStatusResponseBody
	GetErrMessage() *string
	SetErrorMessage(v string) *DescribeEndpointSwitchStatusResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DescribeEndpointSwitchStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeEndpointSwitchStatusResponseBody
	GetStatus() *string
	SetSuccess(v string) *DescribeEndpointSwitchStatusResponseBody
	GetSuccess() *string
}

type DescribeEndpointSwitchStatusResponseBody struct {
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
	// This parameter is no longer available.
	//
	// example:
	//
	// 400
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0DD6B201-604B-4CAB-B6A8-4B2953B5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Initializing**: The task is being initialized.
	//
	// 	- **Switching**: The task is running.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is completed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEndpointSwitchStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointSwitchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndpointSwitchStatusResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeEndpointSwitchStatusResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeEndpointSwitchStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeEndpointSwitchStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEndpointSwitchStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeEndpointSwitchStatusResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetErrCode(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetErrMessage(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetErrorMessage(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetRequestId(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetStatus(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetSuccess(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
