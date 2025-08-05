// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpgradeClientResponseBody
	GetCode() *string
	SetMessage(v string) *UpgradeClientResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpgradeClientResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeClientResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *UpgradeClientResponseBody
	GetTaskId() *string
}

type UpgradeClientResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 22D97921-16BD-547C-B175-1DC25B1DCD73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the asynchronous job. You can call the DescribeTask operation to query the execution result of an asynchronous job.
	//
	// example:
	//
	// t-000i97jujk0z58a2ignf
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClientResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeClientResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpgradeClientResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeClientResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeClientResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpgradeClientResponseBody) SetCode(v string) *UpgradeClientResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeClientResponseBody) SetMessage(v string) *UpgradeClientResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeClientResponseBody) SetRequestId(v string) *UpgradeClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeClientResponseBody) SetSuccess(v bool) *UpgradeClientResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeClientResponseBody) SetTaskId(v string) *UpgradeClientResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpgradeClientResponseBody) Validate() error {
	return dara.Validate(s)
}
