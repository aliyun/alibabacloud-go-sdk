// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualBridgesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBridges(v []*ListVirtualBridgesResponseBodyBridges) *ListVirtualBridgesResponseBody
	GetBridges() []*ListVirtualBridgesResponseBodyBridges
	SetMaxResults(v int32) *ListVirtualBridgesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListVirtualBridgesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVirtualBridgesResponseBody
	GetRequestId() *string
}

type ListVirtualBridgesResponseBody struct {
	// The virtual bridge information.
	Bridges []*ListVirtualBridgesResponseBodyBridges `json:"Bridges,omitempty" xml:"Bridges,omitempty" type:"Repeated"`
	// The maximum number of entries returned. Valid values: 1 to 100. If this parameter is not specified, the default value 100 is used. The number of returned entries can be less than but cannot be greater than the specified number.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVirtualBridgesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualBridgesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirtualBridgesResponseBody) GetBridges() []*ListVirtualBridgesResponseBodyBridges {
	return s.Bridges
}

func (s *ListVirtualBridgesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVirtualBridgesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVirtualBridgesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVirtualBridgesResponseBody) SetBridges(v []*ListVirtualBridgesResponseBodyBridges) *ListVirtualBridgesResponseBody {
	s.Bridges = v
	return s
}

func (s *ListVirtualBridgesResponseBody) SetMaxResults(v int32) *ListVirtualBridgesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVirtualBridgesResponseBody) SetNextToken(v string) *ListVirtualBridgesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVirtualBridgesResponseBody) SetRequestId(v string) *ListVirtualBridgesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirtualBridgesResponseBody) Validate() error {
	if s.Bridges != nil {
		for _, item := range s.Bridges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVirtualBridgesResponseBodyBridges struct {
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
	// vb-sfijaosifiosjfoij***
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
	// panbit
	BridgeType *string `json:"BridgeType,omitempty" xml:"BridgeType,omitempty"`
	// The expiration time. The time is in the ISO 8601 standard in the UTC format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2025-11-07T02:02:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The internal network address.
	//
	// example:
	//
	// http://10.0.0.0:8080
	IntranetUrl *string `json:"IntranetUrl,omitempty" xml:"IntranetUrl,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-shenzhen+dir-8485473914
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The office network name.
	//
	// example:
	//
	// hk11
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
}

func (s ListVirtualBridgesResponseBodyBridges) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualBridgesResponseBodyBridges) GoString() string {
	return s.String()
}

func (s *ListVirtualBridgesResponseBodyBridges) GetAccessType() *string {
	return s.AccessType
}

func (s *ListVirtualBridgesResponseBodyBridges) GetBridgeId() *string {
	return s.BridgeId
}

func (s *ListVirtualBridgesResponseBodyBridges) GetBridgeLevel() *string {
	return s.BridgeLevel
}

func (s *ListVirtualBridgesResponseBodyBridges) GetBridgeStatus() *string {
	return s.BridgeStatus
}

func (s *ListVirtualBridgesResponseBodyBridges) GetBridgeType() *string {
	return s.BridgeType
}

func (s *ListVirtualBridgesResponseBodyBridges) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListVirtualBridgesResponseBodyBridges) GetIntranetUrl() *string {
	return s.IntranetUrl
}

func (s *ListVirtualBridgesResponseBodyBridges) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ListVirtualBridgesResponseBodyBridges) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *ListVirtualBridgesResponseBodyBridges) SetAccessType(v string) *ListVirtualBridgesResponseBodyBridges {
	s.AccessType = &v
	return s
}

func (s *ListVirtualBridgesResponseBodyBridges) SetBridgeId(v string) *ListVirtualBridgesResponseBodyBridges {
	s.BridgeId = &v
	return s
}

func (s *ListVirtualBridgesResponseBodyBridges) SetBridgeLevel(v string) *ListVirtualBridgesResponseBodyBridges {
	s.BridgeLevel = &v
	return s
}

func (s *ListVirtualBridgesResponseBodyBridges) SetBridgeStatus(v string) *ListVirtualBridgesResponseBodyBridges {
	s.BridgeStatus = &v
	return s
}

func (s *ListVirtualBridgesResponseBodyBridges) SetBridgeType(v string) *ListVirtualBridgesResponseBodyBridges {
	s.BridgeType = &v
	return s
}

func (s *ListVirtualBridgesResponseBodyBridges) SetExpireTime(v string) *ListVirtualBridgesResponseBodyBridges {
	s.ExpireTime = &v
	return s
}

func (s *ListVirtualBridgesResponseBodyBridges) SetIntranetUrl(v string) *ListVirtualBridgesResponseBodyBridges {
	s.IntranetUrl = &v
	return s
}

func (s *ListVirtualBridgesResponseBodyBridges) SetOfficeSiteId(v string) *ListVirtualBridgesResponseBodyBridges {
	s.OfficeSiteId = &v
	return s
}

func (s *ListVirtualBridgesResponseBodyBridges) SetOfficeSiteName(v string) *ListVirtualBridgesResponseBodyBridges {
	s.OfficeSiteName = &v
	return s
}

func (s *ListVirtualBridgesResponseBodyBridges) Validate() error {
	return dara.Validate(s)
}
