// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagWifiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationType(v string) *DescribeSagWifiResponseBody
	GetAuthenticationType() *string
	SetBandwidth(v string) *DescribeSagWifiResponseBody
	GetBandwidth() *string
	SetChannel(v string) *DescribeSagWifiResponseBody
	GetChannel() *string
	SetEncryptAlgorithm(v string) *DescribeSagWifiResponseBody
	GetEncryptAlgorithm() *string
	SetIsAuth(v string) *DescribeSagWifiResponseBody
	GetIsAuth() *string
	SetIsBroadcast(v string) *DescribeSagWifiResponseBody
	GetIsBroadcast() *string
	SetIsEnable(v string) *DescribeSagWifiResponseBody
	GetIsEnable() *string
	SetRequestId(v string) *DescribeSagWifiResponseBody
	GetRequestId() *string
	SetSsid(v string) *DescribeSagWifiResponseBody
	GetSsid() *string
	SetTaskStates(v []*DescribeSagWifiResponseBodyTaskStates) *DescribeSagWifiResponseBody
	GetTaskStates() []*DescribeSagWifiResponseBodyTaskStates
}

type DescribeSagWifiResponseBody struct {
	// The authentication type. Valid values:
	//
	// 	- **NONE**: authentication is disabled.
	//
	// 	- **WPA-PSK**: WPA-PSK authentication is enabled.
	//
	// 	- **WPA2-PSK**: WPA2-PSK authentication is enabled.
	//
	// example:
	//
	// WPA2-PSK
	AuthenticationType *string `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// The bandwidth of the Wi-Fi channel. Valid values:
	//
	// 	- **Automatic**
	//
	// 	- **20 HMz**
	//
	// 	- **40 MHz**
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The Wi-Fi channel.
	//
	// example:
	//
	// 0
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The encryption algorithm.
	//
	// 	- **AUTO**: automatically selects the encryption algorithm.
	//
	// 	- **TKIP**: uses the Temporal Key Integrity Protocol (TKIP).
	//
	// 	- **AES**: uses the Advanced Encryption Standard authorized by Wi-Fi®.
	//
	// example:
	//
	// AES
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Indicates whether wireless security is enabled.
	//
	// 	- **True**: wireless security is enabled.
	//
	// 	- **False**: wireless security is disabled.
	//
	// example:
	//
	// True
	IsAuth *string `json:"IsAuth,omitempty" xml:"IsAuth,omitempty"`
	// Indicates whether broadcast over Wi-Fi is enabled. Valid values:
	//
	// 	- **True**: broadcast is enabled.
	//
	// 	- **False**: broadcast is disabled.
	//
	// example:
	//
	// True
	IsBroadcast *string `json:"IsBroadcast,omitempty" xml:"IsBroadcast,omitempty"`
	// Indicates whether wireless connections are enabled. Valid values:
	//
	// 	- **True**: wireless connections are enabled.
	//
	// 	- **False**: wireless connections are disabled.
	//
	// example:
	//
	// True
	IsEnable *string `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E93884AC-6C21-4FEA-8E3A-7377D33B194F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service set identifier (SSID) of the Wi-Fi network.
	//
	// example:
	//
	// aliyun_sag_123456***
	Ssid *string `json:"Ssid,omitempty" xml:"Ssid,omitempty"`
	// The information about the query task.
	TaskStates []*DescribeSagWifiResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s DescribeSagWifiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagWifiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagWifiResponseBody) GetAuthenticationType() *string {
	return s.AuthenticationType
}

func (s *DescribeSagWifiResponseBody) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *DescribeSagWifiResponseBody) GetChannel() *string {
	return s.Channel
}

func (s *DescribeSagWifiResponseBody) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *DescribeSagWifiResponseBody) GetIsAuth() *string {
	return s.IsAuth
}

func (s *DescribeSagWifiResponseBody) GetIsBroadcast() *string {
	return s.IsBroadcast
}

func (s *DescribeSagWifiResponseBody) GetIsEnable() *string {
	return s.IsEnable
}

func (s *DescribeSagWifiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagWifiResponseBody) GetSsid() *string {
	return s.Ssid
}

func (s *DescribeSagWifiResponseBody) GetTaskStates() []*DescribeSagWifiResponseBodyTaskStates {
	return s.TaskStates
}

func (s *DescribeSagWifiResponseBody) SetAuthenticationType(v string) *DescribeSagWifiResponseBody {
	s.AuthenticationType = &v
	return s
}

func (s *DescribeSagWifiResponseBody) SetBandwidth(v string) *DescribeSagWifiResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeSagWifiResponseBody) SetChannel(v string) *DescribeSagWifiResponseBody {
	s.Channel = &v
	return s
}

func (s *DescribeSagWifiResponseBody) SetEncryptAlgorithm(v string) *DescribeSagWifiResponseBody {
	s.EncryptAlgorithm = &v
	return s
}

func (s *DescribeSagWifiResponseBody) SetIsAuth(v string) *DescribeSagWifiResponseBody {
	s.IsAuth = &v
	return s
}

func (s *DescribeSagWifiResponseBody) SetIsBroadcast(v string) *DescribeSagWifiResponseBody {
	s.IsBroadcast = &v
	return s
}

func (s *DescribeSagWifiResponseBody) SetIsEnable(v string) *DescribeSagWifiResponseBody {
	s.IsEnable = &v
	return s
}

func (s *DescribeSagWifiResponseBody) SetRequestId(v string) *DescribeSagWifiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagWifiResponseBody) SetSsid(v string) *DescribeSagWifiResponseBody {
	s.Ssid = &v
	return s
}

func (s *DescribeSagWifiResponseBody) SetTaskStates(v []*DescribeSagWifiResponseBodyTaskStates) *DescribeSagWifiResponseBody {
	s.TaskStates = v
	return s
}

func (s *DescribeSagWifiResponseBody) Validate() error {
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

type DescribeSagWifiResponseBodyTaskStates struct {
	// The time when the query task was created.
	//
	// example:
	//
	// 1586843621000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code returned for a query task. The 200 error code indicates that the query task is successful.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned for a query task. The Successful error message indicates that the query task is successful.
	//
	// example:
	//
	// Successful
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The status of the query task. Valid values:
	//
	// 	- **Initialized**: The query task has been initialized.
	//
	// 	- **Offline**: The query task is not dispatched because the SAG device is disconnected from Alibaba Cloud. The task will be dispatched after the SAG device is connected to Alibaba Cloud.
	//
	// 	- **Succeed**: The query task has been dispatched.
	//
	// 	- **Processing**: The query task is being dispatched.
	//
	// 	- **VersionNotSupport**: The current version of the SAG device does not support query tasks.
	//
	// 	- **BuildRequestError**: The SAG control system does not support query tasks.
	//
	// 	- **HardwareError**: The query task failed to be dispatched due to device errors.
	//
	// 	- **TaskNotExist**: The query task does not exist.
	//
	// 	- **OfflineNotConfiged**: The query task is not dispatched because the SAG device is disconnected from Alibaba Cloud. The task will not be dispatched after the device is connected to Alibaba Cloud.
	//
	// example:
	//
	// Succeed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeSagWifiResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagWifiResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *DescribeSagWifiResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSagWifiResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSagWifiResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSagWifiResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *DescribeSagWifiResponseBodyTaskStates) SetCreateTime(v string) *DescribeSagWifiResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *DescribeSagWifiResponseBodyTaskStates) SetErrorCode(v string) *DescribeSagWifiResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSagWifiResponseBodyTaskStates) SetErrorMessage(v string) *DescribeSagWifiResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSagWifiResponseBodyTaskStates) SetState(v string) *DescribeSagWifiResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *DescribeSagWifiResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
