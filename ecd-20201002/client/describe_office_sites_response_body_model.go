// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOfficeSitesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOfficeSites(v []*DescribeOfficeSitesResponseBodyOfficeSites) *DescribeOfficeSitesResponseBody
	GetOfficeSites() []*DescribeOfficeSitesResponseBodyOfficeSites
	SetRequestId(v string) *DescribeOfficeSitesResponseBody
	GetRequestId() *string
}

type DescribeOfficeSitesResponseBody struct {
	OfficeSites []*DescribeOfficeSitesResponseBodyOfficeSites `json:"OfficeSites,omitempty" xml:"OfficeSites,omitempty" type:"Repeated"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOfficeSitesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfficeSitesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBody) GetOfficeSites() []*DescribeOfficeSitesResponseBodyOfficeSites {
	return s.OfficeSites
}

func (s *DescribeOfficeSitesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOfficeSitesResponseBody) SetOfficeSites(v []*DescribeOfficeSitesResponseBodyOfficeSites) *DescribeOfficeSitesResponseBody {
	s.OfficeSites = v
	return s
}

func (s *DescribeOfficeSitesResponseBody) SetRequestId(v string) *DescribeOfficeSitesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBody) Validate() error {
	if s.OfficeSites != nil {
		for _, item := range s.OfficeSites {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOfficeSitesResponseBodyOfficeSites struct {
	// example:
	//
	// VPC
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	// example:
	//
	// http://ep-bp1s2vmbj55r5rzc****.epsrv-bp1pcfhpwvlpny01****.cn-hangzhou.privatelink.aliyuncs.com
	DesktopVpcEndpoint *string `json:"DesktopVpcEndpoint,omitempty" xml:"DesktopVpcEndpoint,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// SIMPLE
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	// example:
	//
	// 268****
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// example:
	//
	// https://eds-cn-shanghai-67****
	SsoServiceUrl *string `json:"SsoServiceUrl,omitempty" xml:"SsoServiceUrl,omitempty"`
}

func (s DescribeOfficeSitesResponseBodyOfficeSites) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfficeSitesResponseBodyOfficeSites) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetDesktopAccessType() *string {
	return s.DesktopAccessType
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetDesktopVpcEndpoint() *string {
	return s.DesktopVpcEndpoint
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetOfficeSiteType() *string {
	return s.OfficeSiteType
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetProviderId() *string {
	return s.ProviderId
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) GetSsoServiceUrl() *string {
	return s.SsoServiceUrl
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDesktopAccessType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDesktopVpcEndpoint(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DesktopVpcEndpoint = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetOfficeSiteId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetOfficeSiteType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetProviderId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.ProviderId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetSsoServiceUrl(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.SsoServiceUrl = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) Validate() error {
	return dara.Validate(s)
}
