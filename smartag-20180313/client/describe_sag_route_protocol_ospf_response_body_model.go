// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagRouteProtocolOspfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAreaId(v string) *DescribeSagRouteProtocolOspfResponseBody
	GetAreaId() *string
	SetAreaType(v string) *DescribeSagRouteProtocolOspfResponseBody
	GetAreaType() *string
	SetAuthenticationType(v string) *DescribeSagRouteProtocolOspfResponseBody
	GetAuthenticationType() *string
	SetDeadTime(v int32) *DescribeSagRouteProtocolOspfResponseBody
	GetDeadTime() *int32
	SetHelloTime(v int32) *DescribeSagRouteProtocolOspfResponseBody
	GetHelloTime() *int32
	SetMd5Key(v string) *DescribeSagRouteProtocolOspfResponseBody
	GetMd5Key() *string
	SetMd5KeyId(v int32) *DescribeSagRouteProtocolOspfResponseBody
	GetMd5KeyId() *int32
	SetRequestId(v string) *DescribeSagRouteProtocolOspfResponseBody
	GetRequestId() *string
	SetRouterId(v string) *DescribeSagRouteProtocolOspfResponseBody
	GetRouterId() *string
	SetTaskStates(v []*DescribeSagRouteProtocolOspfResponseBodyTaskStates) *DescribeSagRouteProtocolOspfResponseBody
	GetTaskStates() []*DescribeSagRouteProtocolOspfResponseBodyTaskStates
}

type DescribeSagRouteProtocolOspfResponseBody struct {
	// The ID of the OSPF area.
	//
	// example:
	//
	// 10
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The type of the OSPF area.
	//
	// >  Only the NSSA area type is supported.
	//
	// example:
	//
	// NSSA
	AreaType *string `json:"AreaType,omitempty" xml:"AreaType,omitempty"`
	// The authentication type. Valid values:
	//
	// 	- **NONE**: Authentication is disabled.
	//
	// 	- **CLEARTEXT**: plaintext authentication
	//
	// 	- **MD5**: MD5 authentication
	//
	// example:
	//
	// MD5
	AuthenticationType *string `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// The timeout period of connections between OSPF peers.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 40
	DeadTime *int32 `json:"DeadTime,omitempty" xml:"DeadTime,omitempty"`
	// The time interval at which Hello packets are sent.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 10
	HelloTime *int32 `json:"HelloTime,omitempty" xml:"HelloTime,omitempty"`
	// The MD5 key value.
	//
	// example:
	//
	// 123****
	Md5Key *string `json:"Md5Key,omitempty" xml:"Md5Key,omitempty"`
	// The ID of the MD5 key.
	//
	// example:
	//
	// 1
	Md5KeyId *int32 `json:"Md5KeyId,omitempty" xml:"Md5KeyId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 60F9B653-25B7-4511-A3C7-BCBAF462393E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the router that has OSPF enabled.
	//
	// example:
	//
	// 1.XX.XX.1
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// The status of the query task.
	TaskStates []*DescribeSagRouteProtocolOspfResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s DescribeSagRouteProtocolOspfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRouteProtocolOspfResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagRouteProtocolOspfResponseBody) GetAreaId() *string {
	return s.AreaId
}

func (s *DescribeSagRouteProtocolOspfResponseBody) GetAreaType() *string {
	return s.AreaType
}

func (s *DescribeSagRouteProtocolOspfResponseBody) GetAuthenticationType() *string {
	return s.AuthenticationType
}

func (s *DescribeSagRouteProtocolOspfResponseBody) GetDeadTime() *int32 {
	return s.DeadTime
}

func (s *DescribeSagRouteProtocolOspfResponseBody) GetHelloTime() *int32 {
	return s.HelloTime
}

func (s *DescribeSagRouteProtocolOspfResponseBody) GetMd5Key() *string {
	return s.Md5Key
}

func (s *DescribeSagRouteProtocolOspfResponseBody) GetMd5KeyId() *int32 {
	return s.Md5KeyId
}

func (s *DescribeSagRouteProtocolOspfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagRouteProtocolOspfResponseBody) GetRouterId() *string {
	return s.RouterId
}

func (s *DescribeSagRouteProtocolOspfResponseBody) GetTaskStates() []*DescribeSagRouteProtocolOspfResponseBodyTaskStates {
	return s.TaskStates
}

func (s *DescribeSagRouteProtocolOspfResponseBody) SetAreaId(v string) *DescribeSagRouteProtocolOspfResponseBody {
	s.AreaId = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponseBody) SetAreaType(v string) *DescribeSagRouteProtocolOspfResponseBody {
	s.AreaType = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponseBody) SetAuthenticationType(v string) *DescribeSagRouteProtocolOspfResponseBody {
	s.AuthenticationType = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponseBody) SetDeadTime(v int32) *DescribeSagRouteProtocolOspfResponseBody {
	s.DeadTime = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponseBody) SetHelloTime(v int32) *DescribeSagRouteProtocolOspfResponseBody {
	s.HelloTime = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponseBody) SetMd5Key(v string) *DescribeSagRouteProtocolOspfResponseBody {
	s.Md5Key = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponseBody) SetMd5KeyId(v int32) *DescribeSagRouteProtocolOspfResponseBody {
	s.Md5KeyId = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponseBody) SetRequestId(v string) *DescribeSagRouteProtocolOspfResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponseBody) SetRouterId(v string) *DescribeSagRouteProtocolOspfResponseBody {
	s.RouterId = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponseBody) SetTaskStates(v []*DescribeSagRouteProtocolOspfResponseBodyTaskStates) *DescribeSagRouteProtocolOspfResponseBody {
	s.TaskStates = v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponseBody) Validate() error {
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

type DescribeSagRouteProtocolOspfResponseBodyTaskStates struct {
	// The time when the query task was created.
	//
	// example:
	//
	// 1586843621000
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

func (s DescribeSagRouteProtocolOspfResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRouteProtocolOspfResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *DescribeSagRouteProtocolOspfResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSagRouteProtocolOspfResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSagRouteProtocolOspfResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSagRouteProtocolOspfResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *DescribeSagRouteProtocolOspfResponseBodyTaskStates) SetCreateTime(v string) *DescribeSagRouteProtocolOspfResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponseBodyTaskStates) SetErrorCode(v string) *DescribeSagRouteProtocolOspfResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponseBodyTaskStates) SetErrorMessage(v string) *DescribeSagRouteProtocolOspfResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponseBodyTaskStates) SetState(v string) *DescribeSagRouteProtocolOspfResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
