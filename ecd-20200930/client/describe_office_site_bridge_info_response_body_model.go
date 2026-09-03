// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOfficeSiteBridgeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBridge(v *DescribeOfficeSiteBridgeInfoResponseBodyBridge) *DescribeOfficeSiteBridgeInfoResponseBody
	GetBridge() *DescribeOfficeSiteBridgeInfoResponseBodyBridge
	SetRequestId(v string) *DescribeOfficeSiteBridgeInfoResponseBody
	GetRequestId() *string
}

type DescribeOfficeSiteBridgeInfoResponseBody struct {
	// The virtual bridge information.
	Bridge *DescribeOfficeSiteBridgeInfoResponseBodyBridge `json:"Bridge,omitempty" xml:"Bridge,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F7E4322D-D679-5ACB-A909-490D2F0E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOfficeSiteBridgeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfficeSiteBridgeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSiteBridgeInfoResponseBody) GetBridge() *DescribeOfficeSiteBridgeInfoResponseBodyBridge {
	return s.Bridge
}

func (s *DescribeOfficeSiteBridgeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOfficeSiteBridgeInfoResponseBody) SetBridge(v *DescribeOfficeSiteBridgeInfoResponseBodyBridge) *DescribeOfficeSiteBridgeInfoResponseBody {
	s.Bridge = v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponseBody) SetRequestId(v string) *DescribeOfficeSiteBridgeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponseBody) Validate() error {
	if s.Bridge != nil {
		if err := s.Bridge.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOfficeSiteBridgeInfoResponseBodyBridge struct {
	// The access type of the management page.
	//
	// example:
	//
	// intranet
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The virtual bridge ID.
	//
	// example:
	//
	// vb-sofiahfish***
	BridgeId *string `json:"BridgeId,omitempty" xml:"BridgeId,omitempty"`
	// The virtual bridge specifications.
	//
	// example:
	//
	// vb.pro
	BridgeLevel *string `json:"BridgeLevel,omitempty" xml:"BridgeLevel,omitempty"`
	// The virtual bridge status.
	//
	// example:
	//
	// inuse
	BridgeStatus *string `json:"BridgeStatus,omitempty" xml:"BridgeStatus,omitempty"`
	// The third-party plugin type of the virtual bridge.
	//
	// example:
	//
	// unsr
	BridgeType *string `json:"BridgeType,omitempty" xml:"BridgeType,omitempty"`
	// The default password for the administrator page.
	//
	// example:
	//
	// password
	DefaultPassword *string `json:"DefaultPassword,omitempty" xml:"DefaultPassword,omitempty"`
	// The default account for the administrator page.
	//
	// example:
	//
	// user
	DefaultUser *string `json:"DefaultUser,omitempty" xml:"DefaultUser,omitempty"`
	// The deployment time of the virtual bridge. The time is in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-11-07T02:02:00Z
	DeployTime *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	// The expiration time of the virtual bridge. The time is in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2026-03-18T00:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The public network address.
	//
	// example:
	//
	// http://8.*.*.*:8080
	InternetUrl *string `json:"InternetUrl,omitempty" xml:"InternetUrl,omitempty"`
	// The internal network address.
	//
	// example:
	//
	// http://10.0.0.0:8080
	IntranetUrl *string `json:"IntranetUrl,omitempty" xml:"IntranetUrl,omitempty"`
	// The ID of the locked convenience office network.
	//
	// example:
	//
	// cn-beijing+dir-0211574032
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The office network name. The name must be 2 to 255 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter or Chinese character and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// P0801-1
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The start time of the virtual bridge. The time is in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-11-07T02:02:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOfficeSiteBridgeInfoResponseBodyBridge) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfficeSiteBridgeInfoResponseBodyBridge) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) GetAccessType() *string {
	return s.AccessType
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) GetBridgeId() *string {
	return s.BridgeId
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) GetBridgeLevel() *string {
	return s.BridgeLevel
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) GetBridgeStatus() *string {
	return s.BridgeStatus
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) GetBridgeType() *string {
	return s.BridgeType
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) GetDefaultPassword() *string {
	return s.DefaultPassword
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) GetDefaultUser() *string {
	return s.DefaultUser
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) GetDeployTime() *string {
	return s.DeployTime
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) GetInternetUrl() *string {
	return s.InternetUrl
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) GetIntranetUrl() *string {
	return s.IntranetUrl
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) SetAccessType(v string) *DescribeOfficeSiteBridgeInfoResponseBodyBridge {
	s.AccessType = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) SetBridgeId(v string) *DescribeOfficeSiteBridgeInfoResponseBodyBridge {
	s.BridgeId = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) SetBridgeLevel(v string) *DescribeOfficeSiteBridgeInfoResponseBodyBridge {
	s.BridgeLevel = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) SetBridgeStatus(v string) *DescribeOfficeSiteBridgeInfoResponseBodyBridge {
	s.BridgeStatus = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) SetBridgeType(v string) *DescribeOfficeSiteBridgeInfoResponseBodyBridge {
	s.BridgeType = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) SetDefaultPassword(v string) *DescribeOfficeSiteBridgeInfoResponseBodyBridge {
	s.DefaultPassword = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) SetDefaultUser(v string) *DescribeOfficeSiteBridgeInfoResponseBodyBridge {
	s.DefaultUser = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) SetDeployTime(v string) *DescribeOfficeSiteBridgeInfoResponseBodyBridge {
	s.DeployTime = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) SetExpireTime(v string) *DescribeOfficeSiteBridgeInfoResponseBodyBridge {
	s.ExpireTime = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) SetInternetUrl(v string) *DescribeOfficeSiteBridgeInfoResponseBodyBridge {
	s.InternetUrl = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) SetIntranetUrl(v string) *DescribeOfficeSiteBridgeInfoResponseBodyBridge {
	s.IntranetUrl = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) SetOfficeSiteId(v string) *DescribeOfficeSiteBridgeInfoResponseBodyBridge {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) SetOfficeSiteName(v string) *DescribeOfficeSiteBridgeInfoResponseBodyBridge {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) SetStartTime(v string) *DescribeOfficeSiteBridgeInfoResponseBodyBridge {
	s.StartTime = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoResponseBodyBridge) Validate() error {
	return dara.Validate(s)
}
