// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayOspfRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAreaId(v int32) *ViewSmartAccessGatewayOspfRouteResponseBody
	GetAreaId() *int32
	SetAuthenticationType(v string) *ViewSmartAccessGatewayOspfRouteResponseBody
	GetAuthenticationType() *string
	SetDeadTime(v int32) *ViewSmartAccessGatewayOspfRouteResponseBody
	GetDeadTime() *int32
	SetHelloTime(v int32) *ViewSmartAccessGatewayOspfRouteResponseBody
	GetHelloTime() *int32
	SetMd5Key(v string) *ViewSmartAccessGatewayOspfRouteResponseBody
	GetMd5Key() *string
	SetMd5KeyId(v int32) *ViewSmartAccessGatewayOspfRouteResponseBody
	GetMd5KeyId() *int32
	SetRequestId(v string) *ViewSmartAccessGatewayOspfRouteResponseBody
	GetRequestId() *string
	SetRouterId(v string) *ViewSmartAccessGatewayOspfRouteResponseBody
	GetRouterId() *string
	SetTaskStates(v []*ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates) *ViewSmartAccessGatewayOspfRouteResponseBody
	GetTaskStates() []*ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates
}

type ViewSmartAccessGatewayOspfRouteResponseBody struct {
	// The ID of the OSPF area.
	//
	// example:
	//
	// 10
	AreaId *int32 `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The authentication type. Valid values:
	//
	// 	- **NONE**: Authentication is disabled.
	//
	// 	- **CLEARTEXT**: Plaintext authentication is used.
	//
	// 	- **MD5**: MD5 authentication is used.
	//
	// example:
	//
	// NONE
	AuthenticationType *string `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// The timeout period of connections between OSPF peers. Unit: seconds.
	//
	// example:
	//
	// 10
	DeadTime *int32 `json:"DeadTime,omitempty" xml:"DeadTime,omitempty"`
	// The time interval at which Hello packets are sent. Unit: seconds.
	//
	// example:
	//
	// 1
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
	// The request ID.
	//
	// example:
	//
	// AA1AE941-84A9-5F83-A955-C8DAF31C2CB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the router that has OSPF enabled.
	//
	// example:
	//
	// 1.XX.XX.1
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// The status of the task.
	TaskStates []*ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s ViewSmartAccessGatewayOspfRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayOspfRouteResponseBody) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) GetAreaId() *int32 {
	return s.AreaId
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) GetAuthenticationType() *string {
	return s.AuthenticationType
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) GetDeadTime() *int32 {
	return s.DeadTime
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) GetHelloTime() *int32 {
	return s.HelloTime
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) GetMd5Key() *string {
	return s.Md5Key
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) GetMd5KeyId() *int32 {
	return s.Md5KeyId
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) GetRouterId() *string {
	return s.RouterId
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) GetTaskStates() []*ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates {
	return s.TaskStates
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) SetAreaId(v int32) *ViewSmartAccessGatewayOspfRouteResponseBody {
	s.AreaId = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) SetAuthenticationType(v string) *ViewSmartAccessGatewayOspfRouteResponseBody {
	s.AuthenticationType = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) SetDeadTime(v int32) *ViewSmartAccessGatewayOspfRouteResponseBody {
	s.DeadTime = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) SetHelloTime(v int32) *ViewSmartAccessGatewayOspfRouteResponseBody {
	s.HelloTime = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) SetMd5Key(v string) *ViewSmartAccessGatewayOspfRouteResponseBody {
	s.Md5Key = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) SetMd5KeyId(v int32) *ViewSmartAccessGatewayOspfRouteResponseBody {
	s.Md5KeyId = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) SetRequestId(v string) *ViewSmartAccessGatewayOspfRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) SetRouterId(v string) *ViewSmartAccessGatewayOspfRouteResponseBody {
	s.RouterId = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) SetTaskStates(v []*ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates) *ViewSmartAccessGatewayOspfRouteResponseBody {
	s.TaskStates = v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBody) Validate() error {
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

type ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates struct {
	// The timestamp when the task was created. Unit: milliseconds.
	//
	// The value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1586843621000
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

func (s ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates) SetCreateTime(v string) *ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates) SetErrorCode(v string) *ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates) SetErrorMessage(v string) *ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates) SetState(v string) *ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
