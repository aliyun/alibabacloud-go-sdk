// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayDnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMasterDns(v string) *ViewSmartAccessGatewayDnsResponseBody
	GetMasterDns() *string
	SetRequestId(v string) *ViewSmartAccessGatewayDnsResponseBody
	GetRequestId() *string
	SetSlaveDns(v string) *ViewSmartAccessGatewayDnsResponseBody
	GetSlaveDns() *string
	SetTaskStates(v []*ViewSmartAccessGatewayDnsResponseBodyTaskStates) *ViewSmartAccessGatewayDnsResponseBody
	GetTaskStates() []*ViewSmartAccessGatewayDnsResponseBodyTaskStates
}

type ViewSmartAccessGatewayDnsResponseBody struct {
	// The IP address of the primary DNS server.
	//
	// example:
	//
	// 114.114.XXX.XXX
	MasterDns *string `json:"MasterDns,omitempty" xml:"MasterDns,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 571AC2E7-8119-58E9-8BFA-1D580CBD1E56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IP address of the secondary DNS server.
	//
	// example:
	//
	// 8.8.XX.XX
	SlaveDns *string `json:"SlaveDns,omitempty" xml:"SlaveDns,omitempty"`
	// The status of the task.
	TaskStates []*ViewSmartAccessGatewayDnsResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s ViewSmartAccessGatewayDnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayDnsResponseBody) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayDnsResponseBody) GetMasterDns() *string {
	return s.MasterDns
}

func (s *ViewSmartAccessGatewayDnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ViewSmartAccessGatewayDnsResponseBody) GetSlaveDns() *string {
	return s.SlaveDns
}

func (s *ViewSmartAccessGatewayDnsResponseBody) GetTaskStates() []*ViewSmartAccessGatewayDnsResponseBodyTaskStates {
	return s.TaskStates
}

func (s *ViewSmartAccessGatewayDnsResponseBody) SetMasterDns(v string) *ViewSmartAccessGatewayDnsResponseBody {
	s.MasterDns = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsResponseBody) SetRequestId(v string) *ViewSmartAccessGatewayDnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsResponseBody) SetSlaveDns(v string) *ViewSmartAccessGatewayDnsResponseBody {
	s.SlaveDns = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsResponseBody) SetTaskStates(v []*ViewSmartAccessGatewayDnsResponseBodyTaskStates) *ViewSmartAccessGatewayDnsResponseBody {
	s.TaskStates = v
	return s
}

func (s *ViewSmartAccessGatewayDnsResponseBody) Validate() error {
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

type ViewSmartAccessGatewayDnsResponseBodyTaskStates struct {
	// The time when the query task was created.
	//
	// The value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1586857309000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code. A value of 200 indicates that the task is successful.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. A value of Successful indicates that the task is successful.
	//
	// example:
	//
	// Successful
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The status of the asynchronous task. Valid values:
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

func (s ViewSmartAccessGatewayDnsResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayDnsResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayDnsResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ViewSmartAccessGatewayDnsResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ViewSmartAccessGatewayDnsResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ViewSmartAccessGatewayDnsResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *ViewSmartAccessGatewayDnsResponseBodyTaskStates) SetCreateTime(v string) *ViewSmartAccessGatewayDnsResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsResponseBodyTaskStates) SetErrorCode(v string) *ViewSmartAccessGatewayDnsResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsResponseBodyTaskStates) SetErrorMessage(v string) *ViewSmartAccessGatewayDnsResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsResponseBodyTaskStates) SetState(v string) *ViewSmartAccessGatewayDnsResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
