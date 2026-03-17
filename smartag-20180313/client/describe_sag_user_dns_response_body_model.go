// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagUserDnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMasterDns(v string) *DescribeSagUserDnsResponseBody
	GetMasterDns() *string
	SetRequestId(v string) *DescribeSagUserDnsResponseBody
	GetRequestId() *string
	SetSlaveDns(v string) *DescribeSagUserDnsResponseBody
	GetSlaveDns() *string
	SetTaskStates(v []*DescribeSagUserDnsResponseBodyTaskStates) *DescribeSagUserDnsResponseBody
	GetTaskStates() []*DescribeSagUserDnsResponseBodyTaskStates
}

type DescribeSagUserDnsResponseBody struct {
	// The IP address of the primary DNS server.
	//
	// example:
	//
	// 192.XX.XX.1
	MasterDns *string `json:"MasterDns,omitempty" xml:"MasterDns,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0937DEA0-AB4B-42F4-9314-07B97D30282B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IP address of the secondary DNS server.
	//
	// example:
	//
	// 192.XX.XX.2
	SlaveDns *string `json:"SlaveDns,omitempty" xml:"SlaveDns,omitempty"`
	// The state of the query task.
	TaskStates []*DescribeSagUserDnsResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s DescribeSagUserDnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagUserDnsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagUserDnsResponseBody) GetMasterDns() *string {
	return s.MasterDns
}

func (s *DescribeSagUserDnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagUserDnsResponseBody) GetSlaveDns() *string {
	return s.SlaveDns
}

func (s *DescribeSagUserDnsResponseBody) GetTaskStates() []*DescribeSagUserDnsResponseBodyTaskStates {
	return s.TaskStates
}

func (s *DescribeSagUserDnsResponseBody) SetMasterDns(v string) *DescribeSagUserDnsResponseBody {
	s.MasterDns = &v
	return s
}

func (s *DescribeSagUserDnsResponseBody) SetRequestId(v string) *DescribeSagUserDnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagUserDnsResponseBody) SetSlaveDns(v string) *DescribeSagUserDnsResponseBody {
	s.SlaveDns = &v
	return s
}

func (s *DescribeSagUserDnsResponseBody) SetTaskStates(v []*DescribeSagUserDnsResponseBodyTaskStates) *DescribeSagUserDnsResponseBody {
	s.TaskStates = v
	return s
}

func (s *DescribeSagUserDnsResponseBody) Validate() error {
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

type DescribeSagUserDnsResponseBodyTaskStates struct {
	// The time when the query task was created.
	//
	// example:
	//
	// 1586852928000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code returned. A value of 200 indicates that the query task is successful.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. A value of Successful indicates that the query task is successful.
	//
	// example:
	//
	// Successful
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The state of the query task. Valid values:
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

func (s DescribeSagUserDnsResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagUserDnsResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *DescribeSagUserDnsResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSagUserDnsResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSagUserDnsResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSagUserDnsResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *DescribeSagUserDnsResponseBodyTaskStates) SetCreateTime(v string) *DescribeSagUserDnsResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *DescribeSagUserDnsResponseBodyTaskStates) SetErrorCode(v string) *DescribeSagUserDnsResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSagUserDnsResponseBodyTaskStates) SetErrorMessage(v string) *DescribeSagUserDnsResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSagUserDnsResponseBodyTaskStates) SetState(v string) *DescribeSagUserDnsResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *DescribeSagUserDnsResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
