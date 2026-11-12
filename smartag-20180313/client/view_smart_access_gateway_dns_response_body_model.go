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
	// The status of the query task.
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
	// The time is a UNIX timestamp that indicates the number of seconds elapsed since 00:00:00 UTC on January 1, 1970.
	//
	// example:
	//
	// 1586857309000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code. A value of 200 indicates that the query task is successful.
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
	// The status of the asynchronous task. Valid values:
	//
	// - **Initialized**: The query task is initialized.
	//
	// - **Offline**: The Smart Access Gateway device is offline and the query task is not delivered. The task will be delivered after the device goes online.
	//
	// - **Succeed**: The query task is delivered successfully.
	//
	// - **Processing**: The query task is being delivered.
	//
	// - **VersionNotSupport**: The current version of the Smart Access Gateway device does not support this operation.
	//
	// - **BuildRequestError**: The management and control system does not support this operation.
	//
	// - **HardwareError**: The query task fails to be delivered due to device errors.
	//
	// - **TaskNotExist**: The query task does not exist.
	//
	// - **OfflineNotConfiged**: The Smart Access Gateway device is offline and the query task is not delivered. The task will not be delivered after the device goes online.
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
