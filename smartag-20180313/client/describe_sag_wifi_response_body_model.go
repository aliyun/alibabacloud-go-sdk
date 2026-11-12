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
	// - **NONE**: Authentication is disabled.
	//
	// - **WPA-PSK**: WPA-PSK authentication is enabled.
	//
	// - **WPA2-PSK**: WPA2-PSK authentication is enabled.
	//
	// example:
	//
	// WPA2-PSK
	AuthenticationType *string `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// The spectrum bandwidth. Valid values:
	//
	// - **Auto**.
	//
	// - **20MHz**.
	//
	// - **40MHz**.
	//
	// example:
	//
	// Automatic
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The channel.
	//
	// example:
	//
	// 0
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The encryption algorithm. Valid values:
	//
	// - **AUTO**: Automatic encryption.
	//
	// - **TKIP**: Temporal Key Integrity Protocol.
	//
	// - **AES**: Advanced Encryption Standard.
	//
	// example:
	//
	// AES
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Indicates whether the wireless security feature is enabled. Valid values:
	//
	// - **True**: The wireless security feature is enabled.
	//
	// - **False**: The wireless security feature is disabled.
	//
	// example:
	//
	// True
	IsAuth *string `json:"IsAuth,omitempty" xml:"IsAuth,omitempty"`
	// Indicates whether the Wi-Fi broadcast feature is enabled. Valid values:
	//
	// - **True**: The Wi-Fi broadcast feature is enabled.
	//
	// - **False**: The Wi-Fi broadcast feature is disabled.
	//
	// example:
	//
	// True
	IsBroadcast *string `json:"IsBroadcast,omitempty" xml:"IsBroadcast,omitempty"`
	// Indicates whether the wireless feature is enabled. Valid values:
	//
	// - **True**: Enabled.
	//
	// - **False**: Disabled.
	//
	// example:
	//
	// True
	IsEnable *string `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E93884AC-6C21-4FEA-8E3A-7377D33B194F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Wi-Fi name.
	//
	// example:
	//
	// aliyun_sag_123456***
	Ssid *string `json:"Ssid,omitempty" xml:"Ssid,omitempty"`
	// The status of the query task.
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
	// - **Succeed**: The query task is delivered.
	//
	// - **Processing**: The query task is being delivered.
	//
	// - **VersionNotSupport**: The current version of the Smart Access Gateway device does not support this feature.
	//
	// - **BuildRequestError**: The management system does not support the operation.
	//
	// - **HardwareError**: The query task fails to be delivered due to device errors.
	//
	// - **TaskNotExist**: The query task does not exist.
	//
	// - **OfflineNotConfiged**: The Smart Access Gateway device is offline and the query task is not delivered. The task will not be delivered even after the device goes online.
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
