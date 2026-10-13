// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPortalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPortalResponseBody
	GetCode() *string
	SetData(v *GetPortalResponseBodyData) *GetPortalResponseBody
	GetData() *GetPortalResponseBodyData
	SetMessage(v string) *GetPortalResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPortalResponseBody
	GetRequestId() *string
}

type GetPortalResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetPortalResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPortalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPortalResponseBody) GoString() string {
	return s.String()
}

func (s *GetPortalResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPortalResponseBody) GetData() *GetPortalResponseBodyData {
	return s.Data
}

func (s *GetPortalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPortalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPortalResponseBody) SetCode(v string) *GetPortalResponseBody {
	s.Code = &v
	return s
}

func (s *GetPortalResponseBody) SetData(v *GetPortalResponseBodyData) *GetPortalResponseBody {
	s.Data = v
	return s
}

func (s *GetPortalResponseBody) SetMessage(v string) *GetPortalResponseBody {
	s.Message = &v
	return s
}

func (s *GetPortalResponseBody) SetRequestId(v string) *GetPortalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPortalResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPortalResponseBodyData struct {
	// The creation time in ISO 8601 (RFC 3339) format in the UTC time zone, such as 2026-05-15T10:00:00Z.
	//
	// example:
	//
	// 2026-05-15T10:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The portal description, which is the Description property of the portal resource.
	//
	// example:
	//
	// HiMarket portal for internal developers
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The portal name, which is the Name property of the portal resource.
	//
	// example:
	//
	// developer-portal
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of portal domain name configurations. The list element is named domainConfig.
	PortalDomainConfig []*HiMarketPortalDomainConfig `json:"portalDomainConfig,omitempty" xml:"portalDomainConfig,omitempty" type:"Repeated"`
	// The portal ID, which is the primary ID of the portal resource and points to the same portal as the portalId request parameter.
	//
	// example:
	//
	// portal-xxxxxxxx
	PortalId *string `json:"portalId,omitempty" xml:"portalId,omitempty"`
	// The portal settings, including whether to enable username and password authentication, whether to automatically approve developer registrations, and whether to automatically approve subscription requests.
	PortalSettingConfig *HiMarketPortalSettingConfig `json:"portalSettingConfig,omitempty" xml:"portalSettingConfig,omitempty"`
	// The region ID, which is the RegionId property of the portal resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetPortalResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPortalResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPortalResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetPortalResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetPortalResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetPortalResponseBodyData) GetPortalDomainConfig() []*HiMarketPortalDomainConfig {
	return s.PortalDomainConfig
}

func (s *GetPortalResponseBodyData) GetPortalId() *string {
	return s.PortalId
}

func (s *GetPortalResponseBodyData) GetPortalSettingConfig() *HiMarketPortalSettingConfig {
	return s.PortalSettingConfig
}

func (s *GetPortalResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPortalResponseBodyData) SetCreatedAt(v string) *GetPortalResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetPortalResponseBodyData) SetDescription(v string) *GetPortalResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetPortalResponseBodyData) SetName(v string) *GetPortalResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetPortalResponseBodyData) SetPortalDomainConfig(v []*HiMarketPortalDomainConfig) *GetPortalResponseBodyData {
	s.PortalDomainConfig = v
	return s
}

func (s *GetPortalResponseBodyData) SetPortalId(v string) *GetPortalResponseBodyData {
	s.PortalId = &v
	return s
}

func (s *GetPortalResponseBodyData) SetPortalSettingConfig(v *HiMarketPortalSettingConfig) *GetPortalResponseBodyData {
	s.PortalSettingConfig = v
	return s
}

func (s *GetPortalResponseBodyData) SetRegionId(v string) *GetPortalResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetPortalResponseBodyData) Validate() error {
	if s.PortalDomainConfig != nil {
		for _, item := range s.PortalDomainConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PortalSettingConfig != nil {
		if err := s.PortalSettingConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
