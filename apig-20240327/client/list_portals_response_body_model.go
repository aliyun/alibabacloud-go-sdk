// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPortalsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPortalsResponseBody
	GetCode() *string
	SetData(v *ListPortalsResponseBodyData) *ListPortalsResponseBody
	GetData() *ListPortalsResponseBodyData
	SetMaxResults(v int32) *ListPortalsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListPortalsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListPortalsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPortalsResponseBody
	GetRequestId() *string
}

type ListPortalsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListPortalsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The maximum number of entries returned in this request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// next-token-xxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPortalsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPortalsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPortalsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPortalsResponseBody) GetData() *ListPortalsResponseBodyData {
	return s.Data
}

func (s *ListPortalsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPortalsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPortalsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPortalsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPortalsResponseBody) SetCode(v string) *ListPortalsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPortalsResponseBody) SetData(v *ListPortalsResponseBodyData) *ListPortalsResponseBody {
	s.Data = v
	return s
}

func (s *ListPortalsResponseBody) SetMaxResults(v int32) *ListPortalsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPortalsResponseBody) SetMessage(v string) *ListPortalsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPortalsResponseBody) SetNextToken(v string) *ListPortalsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPortalsResponseBody) SetRequestId(v string) *ListPortalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPortalsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPortalsResponseBodyData struct {
	// The list of portals. The list element is named portalConfig, and each element contains the information of a HiMarket developer portal.
	Content []*ListPortalsResponseBodyDataContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPortalsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPortalsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPortalsResponseBodyData) GetContent() []*ListPortalsResponseBodyDataContent {
	return s.Content
}

func (s *ListPortalsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPortalsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPortalsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPortalsResponseBodyData) SetContent(v []*ListPortalsResponseBodyDataContent) *ListPortalsResponseBodyData {
	s.Content = v
	return s
}

func (s *ListPortalsResponseBodyData) SetPageNumber(v int32) *ListPortalsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPortalsResponseBodyData) SetPageSize(v int32) *ListPortalsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPortalsResponseBodyData) SetTotalCount(v int32) *ListPortalsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListPortalsResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPortalsResponseBodyDataContent struct {
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
	// The list of portal domain configurations. The list element is named domainConfig.
	PortalDomainConfig []*HiMarketPortalDomainConfig `json:"portalDomainConfig,omitempty" xml:"portalDomainConfig,omitempty" type:"Repeated"`
	// The portal ID, which is the primary ID of the portal resource. It can be used in GetPortal, UpdatePortal, and DeletePortal to locate the portal at /himarket/v1/portals/{portalId}.
	//
	// example:
	//
	// portal-xxxxxxxx
	PortalId *string `json:"portalId,omitempty" xml:"portalId,omitempty"`
	// The portal configuration, including whether username and password authentication is enabled, whether developer registrations are automatically approved, and whether subscription requests are automatically approved.
	PortalSettingConfig *HiMarketPortalSettingConfig `json:"portalSettingConfig,omitempty" xml:"portalSettingConfig,omitempty"`
	// The region ID, which corresponds to the RegionId property of the portal resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListPortalsResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListPortalsResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListPortalsResponseBodyDataContent) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListPortalsResponseBodyDataContent) GetDescription() *string {
	return s.Description
}

func (s *ListPortalsResponseBodyDataContent) GetName() *string {
	return s.Name
}

func (s *ListPortalsResponseBodyDataContent) GetPortalDomainConfig() []*HiMarketPortalDomainConfig {
	return s.PortalDomainConfig
}

func (s *ListPortalsResponseBodyDataContent) GetPortalId() *string {
	return s.PortalId
}

func (s *ListPortalsResponseBodyDataContent) GetPortalSettingConfig() *HiMarketPortalSettingConfig {
	return s.PortalSettingConfig
}

func (s *ListPortalsResponseBodyDataContent) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPortalsResponseBodyDataContent) SetCreatedAt(v string) *ListPortalsResponseBodyDataContent {
	s.CreatedAt = &v
	return s
}

func (s *ListPortalsResponseBodyDataContent) SetDescription(v string) *ListPortalsResponseBodyDataContent {
	s.Description = &v
	return s
}

func (s *ListPortalsResponseBodyDataContent) SetName(v string) *ListPortalsResponseBodyDataContent {
	s.Name = &v
	return s
}

func (s *ListPortalsResponseBodyDataContent) SetPortalDomainConfig(v []*HiMarketPortalDomainConfig) *ListPortalsResponseBodyDataContent {
	s.PortalDomainConfig = v
	return s
}

func (s *ListPortalsResponseBodyDataContent) SetPortalId(v string) *ListPortalsResponseBodyDataContent {
	s.PortalId = &v
	return s
}

func (s *ListPortalsResponseBodyDataContent) SetPortalSettingConfig(v *HiMarketPortalSettingConfig) *ListPortalsResponseBodyDataContent {
	s.PortalSettingConfig = v
	return s
}

func (s *ListPortalsResponseBodyDataContent) SetRegionId(v string) *ListPortalsResponseBodyDataContent {
	s.RegionId = &v
	return s
}

func (s *ListPortalsResponseBodyDataContent) Validate() error {
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
