// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsUserInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePdnsUserInfoResponseBody
	GetRequestId() *string
	SetUserInfo(v *DescribePdnsUserInfoResponseBodyUserInfo) *DescribePdnsUserInfoResponseBody
	GetUserInfo() *DescribePdnsUserInfoResponseBodyUserInfo
}

type DescribePdnsUserInfoResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// FD552816-FCC8-4832-B4A2-2DA0C2BA1688
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the user.
	UserInfo *DescribePdnsUserInfoResponseBodyUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s DescribePdnsUserInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePdnsUserInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePdnsUserInfoResponseBody) GetUserInfo() *DescribePdnsUserInfoResponseBodyUserInfo {
	return s.UserInfo
}

func (s *DescribePdnsUserInfoResponseBody) SetRequestId(v string) *DescribePdnsUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePdnsUserInfoResponseBody) SetUserInfo(v *DescribePdnsUserInfoResponseBodyUserInfo) *DescribePdnsUserInfoResponseBody {
	s.UserInfo = v
	return s
}

func (s *DescribePdnsUserInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePdnsUserInfoResponseBodyUserInfo struct {
	// The enabled access security types.
	//
	// example:
	//
	// SECURE
	AvailableAccessSecurityType *string `json:"AvailableAccessSecurityType,omitempty" xml:"AvailableAccessSecurityType,omitempty"`
	// The enabled public recursive DNS service.
	//
	// example:
	//
	// HTTP,HTTPS
	AvailableService *string `json:"AvailableService,omitempty" xml:"AvailableService,omitempty"`
	// The configuration ID of the users in public recursive DNS.
	//
	// example:
	//
	// 10001
	PdnsId *int64 `json:"PdnsId,omitempty" xml:"PdnsId,omitempty"`
	// The SecretKey configured for a UDP-based CIDR block.
	//
	// example:
	//
	// 1c092d715b7a48de
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The type of the public recursive DNS service.
	//
	// example:
	//
	// normal
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The status of the public recursive DNS service.
	//
	// example:
	//
	// AVAILABLE
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status of the traffic analysis switch for the user in public recursive DNS service.
	//
	// example:
	//
	// CLOSED
	StatisticSwitchStatus *string `json:"StatisticSwitchStatus,omitempty" xml:"StatisticSwitchStatus,omitempty"`
	// The disabled public recursive DNS service.
	StoppedService *string `json:"StoppedService,omitempty" xml:"StoppedService,omitempty"`
}

func (s DescribePdnsUserInfoResponseBodyUserInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsUserInfoResponseBodyUserInfo) GoString() string {
	return s.String()
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) GetAvailableAccessSecurityType() *string {
	return s.AvailableAccessSecurityType
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) GetAvailableService() *string {
	return s.AvailableService
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) GetPdnsId() *int64 {
	return s.PdnsId
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) GetSecretKey() *string {
	return s.SecretKey
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) GetServiceType() *string {
	return s.ServiceType
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) GetState() *string {
	return s.State
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) GetStatisticSwitchStatus() *string {
	return s.StatisticSwitchStatus
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) GetStoppedService() *string {
	return s.StoppedService
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) SetAvailableAccessSecurityType(v string) *DescribePdnsUserInfoResponseBodyUserInfo {
	s.AvailableAccessSecurityType = &v
	return s
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) SetAvailableService(v string) *DescribePdnsUserInfoResponseBodyUserInfo {
	s.AvailableService = &v
	return s
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) SetPdnsId(v int64) *DescribePdnsUserInfoResponseBodyUserInfo {
	s.PdnsId = &v
	return s
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) SetSecretKey(v string) *DescribePdnsUserInfoResponseBodyUserInfo {
	s.SecretKey = &v
	return s
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) SetServiceType(v string) *DescribePdnsUserInfoResponseBodyUserInfo {
	s.ServiceType = &v
	return s
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) SetState(v string) *DescribePdnsUserInfoResponseBodyUserInfo {
	s.State = &v
	return s
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) SetStatisticSwitchStatus(v string) *DescribePdnsUserInfoResponseBodyUserInfo {
	s.StatisticSwitchStatus = &v
	return s
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) SetStoppedService(v string) *DescribePdnsUserInfoResponseBodyUserInfo {
	s.StoppedService = &v
	return s
}

func (s *DescribePdnsUserInfoResponseBodyUserInfo) Validate() error {
	return dara.Validate(s)
}
