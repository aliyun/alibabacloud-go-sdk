// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagWanSnatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSagWanSnatResponseBody
	GetRequestId() *string
	SetSnat(v string) *DescribeSagWanSnatResponseBody
	GetSnat() *string
	SetTaskStates(v []*DescribeSagWanSnatResponseBodyTaskStates) *DescribeSagWanSnatResponseBody
	GetTaskStates() []*DescribeSagWanSnatResponseBodyTaskStates
}

type DescribeSagWanSnatResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CE6642D4-21EB-4168-9BF9-F217953F9892
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether SNAT is enabled.
	//
	// 	- **ENABLE**: SNAT is enabled.
	//
	// 	- **DISABLE**: SNAT is disabled.
	//
	// example:
	//
	// ENABLE
	Snat *string `json:"Snat,omitempty" xml:"Snat,omitempty"`
	// The details about the status of the query task.
	TaskStates []*DescribeSagWanSnatResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s DescribeSagWanSnatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagWanSnatResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagWanSnatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagWanSnatResponseBody) GetSnat() *string {
	return s.Snat
}

func (s *DescribeSagWanSnatResponseBody) GetTaskStates() []*DescribeSagWanSnatResponseBodyTaskStates {
	return s.TaskStates
}

func (s *DescribeSagWanSnatResponseBody) SetRequestId(v string) *DescribeSagWanSnatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagWanSnatResponseBody) SetSnat(v string) *DescribeSagWanSnatResponseBody {
	s.Snat = &v
	return s
}

func (s *DescribeSagWanSnatResponseBody) SetTaskStates(v []*DescribeSagWanSnatResponseBodyTaskStates) *DescribeSagWanSnatResponseBody {
	s.TaskStates = v
	return s
}

func (s *DescribeSagWanSnatResponseBody) Validate() error {
	if s.TaskStates != nil {
		for _, item := range s.TaskStates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSagWanSnatResponseBodyTaskStates struct {
	// The timestamp when the task was created.
	//
	// The value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1586847787000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code returned. A value of 200 indicates that the query task is successful.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned. A value of Successful indicates that the query task is successful.
	//
	// example:
	//
	// Successful
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The status of the query task. Valid values:
	//
	// 	- **Initialized**: The query task is initialized.
	//
	// 	- **Offline**: The SAG device is disconnected from Alibaba Cloud and Alibaba Cloud has not assigned the query task to the SAG device. After the SAG device is connected to Alibaba Cloud, Alibaba Cloud assigns the query task to the SAG device.
	//
	// 	- **Succeed**: Alibaba Cloud has assigned the query task to the SAG device.
	//
	// 	- **Processing**: Alibaba Cloud is assigning the query task to the SAG device.
	//
	// 	- **VersionNotSupport**: The query task is not supported by the current version of the SAG device.
	//
	// 	- **BuildRequestError**: The query task is not supported by the controller of the SAG device.
	//
	// 	- **HardwareError**: Alibaba Cloud failed to assign the query task to the SAG device because the SAG device is faulty.
	//
	// 	- **TaskNotExist**: The query task does not exist.
	//
	// 	- **OfflineNotConfiged**: The SAG device is disconnected from Alibaba Cloud and Alibaba Cloud has not assigned the query task to the SAG device. Alibaba Cloud does not assign the query task to the SAG device even after the SAG device is connected to Alibaba Cloud.
	//
	// example:
	//
	// Succeed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeSagWanSnatResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagWanSnatResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *DescribeSagWanSnatResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSagWanSnatResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSagWanSnatResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSagWanSnatResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *DescribeSagWanSnatResponseBodyTaskStates) SetCreateTime(v string) *DescribeSagWanSnatResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *DescribeSagWanSnatResponseBodyTaskStates) SetErrorCode(v string) *DescribeSagWanSnatResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSagWanSnatResponseBodyTaskStates) SetErrorMessage(v string) *DescribeSagWanSnatResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSagWanSnatResponseBodyTaskStates) SetState(v string) *DescribeSagWanSnatResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *DescribeSagWanSnatResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
