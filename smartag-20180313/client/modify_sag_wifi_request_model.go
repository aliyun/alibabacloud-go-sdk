// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagWifiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationType(v string) *ModifySagWifiRequest
	GetAuthenticationType() *string
	SetBandwidth(v string) *ModifySagWifiRequest
	GetBandwidth() *string
	SetChannel(v string) *ModifySagWifiRequest
	GetChannel() *string
	SetEncryptAlgorithm(v string) *ModifySagWifiRequest
	GetEncryptAlgorithm() *string
	SetIsAuth(v string) *ModifySagWifiRequest
	GetIsAuth() *string
	SetIsBroadcast(v string) *ModifySagWifiRequest
	GetIsBroadcast() *string
	SetIsEnable(v string) *ModifySagWifiRequest
	GetIsEnable() *string
	SetOwnerAccount(v string) *ModifySagWifiRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySagWifiRequest
	GetOwnerId() *int64
	SetPassword(v string) *ModifySagWifiRequest
	GetPassword() *string
	SetRegionId(v string) *ModifySagWifiRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySagWifiRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySagWifiRequest
	GetResourceOwnerId() *int64
	SetSSID(v string) *ModifySagWifiRequest
	GetSSID() *string
	SetSmartAGId(v string) *ModifySagWifiRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *ModifySagWifiRequest
	GetSmartAGSn() *string
}

type ModifySagWifiRequest struct {
	// The authentication type. Valid values:
	//
	// 	- **NONE**
	//
	// 	- **WPA-PSK**
	//
	// 	- **WPA2-PSK**
	//
	// example:
	//
	// WPA2-PSK
	AuthenticationType *string `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// The bandwidth of the channel. Valid values:
	//
	// 	- **Automatic**
	//
	// 	- **20 MHz**
	//
	// 	- **40 MHz**
	//
	// example:
	//
	// 20 MHz
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The Wi-Fi channel.
	//
	// Valid values: **0 to 11**.
	//
	// example:
	//
	// 0
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The encryption algorithm. Valid values:
	//
	// 	- **AUTO**: automatically selects the encryption algorithm.
	//
	// 	- **TKIP**: uses the Temporal Key Integrity Protocol (TKIP).
	//
	// 	- **AES**: uses the Advanced Encryption Standard authorized by Wi-Fi®.
	//
	// example:
	//
	// AUTO
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether wireless security is enabled. Valid values:
	//
	// 	- **true**: enables wireless security.
	//
	// 	- **False**: disables wireless security.
	//
	// example:
	//
	// True
	IsAuth *string `json:"IsAuth,omitempty" xml:"IsAuth,omitempty"`
	// Specifies whether broadcast over Wi-Fi is enabled. Valid values:
	//
	// 	- **true**: enables broadcast.
	//
	// 	- **False**: disables broadcast.
	//
	// >  Only after you enable broadcast, terminals that support wireless connections can search the Wi-Fi network by its SSID and receive Wi-Fi signals.
	//
	// example:
	//
	// True
	IsBroadcast *string `json:"IsBroadcast,omitempty" xml:"IsBroadcast,omitempty"`
	// Specifies whether Wi-Fi is enabled. Valid values:
	//
	// 	- **true**: enables Wi-Fi.
	//
	// 	- **False**: disables Wi-Fi.
	//
	// This parameter is required.
	//
	// example:
	//
	// True
	IsEnable     *string `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password used to connect to the Wi-Fi network.
	//
	// The password must be 8 to 32 characters in length, and can contain digits and letters.
	//
	// example:
	//
	// 12345678
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service set identifier (SSID) of the Wi-Fi network.
	//
	// The name must be 1 to 31 characters in length, and can contain digits and letters.
	//
	// example:
	//
	// aliyun_sag_123456****
	SSID *string `json:"SSID,omitempty" xml:"SSID,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-whfn****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The serial number of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
}

func (s ModifySagWifiRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySagWifiRequest) GoString() string {
	return s.String()
}

func (s *ModifySagWifiRequest) GetAuthenticationType() *string {
	return s.AuthenticationType
}

func (s *ModifySagWifiRequest) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *ModifySagWifiRequest) GetChannel() *string {
	return s.Channel
}

func (s *ModifySagWifiRequest) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *ModifySagWifiRequest) GetIsAuth() *string {
	return s.IsAuth
}

func (s *ModifySagWifiRequest) GetIsBroadcast() *string {
	return s.IsBroadcast
}

func (s *ModifySagWifiRequest) GetIsEnable() *string {
	return s.IsEnable
}

func (s *ModifySagWifiRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySagWifiRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySagWifiRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifySagWifiRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySagWifiRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySagWifiRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySagWifiRequest) GetSSID() *string {
	return s.SSID
}

func (s *ModifySagWifiRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySagWifiRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *ModifySagWifiRequest) SetAuthenticationType(v string) *ModifySagWifiRequest {
	s.AuthenticationType = &v
	return s
}

func (s *ModifySagWifiRequest) SetBandwidth(v string) *ModifySagWifiRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifySagWifiRequest) SetChannel(v string) *ModifySagWifiRequest {
	s.Channel = &v
	return s
}

func (s *ModifySagWifiRequest) SetEncryptAlgorithm(v string) *ModifySagWifiRequest {
	s.EncryptAlgorithm = &v
	return s
}

func (s *ModifySagWifiRequest) SetIsAuth(v string) *ModifySagWifiRequest {
	s.IsAuth = &v
	return s
}

func (s *ModifySagWifiRequest) SetIsBroadcast(v string) *ModifySagWifiRequest {
	s.IsBroadcast = &v
	return s
}

func (s *ModifySagWifiRequest) SetIsEnable(v string) *ModifySagWifiRequest {
	s.IsEnable = &v
	return s
}

func (s *ModifySagWifiRequest) SetOwnerAccount(v string) *ModifySagWifiRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySagWifiRequest) SetOwnerId(v int64) *ModifySagWifiRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySagWifiRequest) SetPassword(v string) *ModifySagWifiRequest {
	s.Password = &v
	return s
}

func (s *ModifySagWifiRequest) SetRegionId(v string) *ModifySagWifiRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySagWifiRequest) SetResourceOwnerAccount(v string) *ModifySagWifiRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySagWifiRequest) SetResourceOwnerId(v int64) *ModifySagWifiRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySagWifiRequest) SetSSID(v string) *ModifySagWifiRequest {
	s.SSID = &v
	return s
}

func (s *ModifySagWifiRequest) SetSmartAGId(v string) *ModifySagWifiRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySagWifiRequest) SetSmartAGSn(v string) *ModifySagWifiRequest {
	s.SmartAGSn = &v
	return s
}

func (s *ModifySagWifiRequest) Validate() error {
	return dara.Validate(s)
}
