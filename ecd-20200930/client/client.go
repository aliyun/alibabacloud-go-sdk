// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddUserToDesktopGroupRequest struct {
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopGroupId *string   `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	ClientToken    *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndUserIds     []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
}

func (s AddUserToDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserToDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUserToDesktopGroupRequest) SetRegionId(v string) *AddUserToDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetDesktopGroupId(v string) *AddUserToDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetClientToken(v string) *AddUserToDesktopGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AddUserToDesktopGroupRequest) SetEndUserIds(v []*string) *AddUserToDesktopGroupRequest {
	s.EndUserIds = v
	return s
}

type AddUserToDesktopGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserToDesktopGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserToDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserToDesktopGroupResponseBody) SetRequestId(v string) *AddUserToDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddUserToDesktopGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddUserToDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUserToDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserToDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *AddUserToDesktopGroupResponse) SetHeaders(v map[string]*string) *AddUserToDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *AddUserToDesktopGroupResponse) SetBody(v *AddUserToDesktopGroupResponseBody) *AddUserToDesktopGroupResponse {
	s.Body = v
	return s
}

type AddUserToSecurityCenterWhiteListRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddUserToSecurityCenterWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserToSecurityCenterWhiteListRequest) GoString() string {
	return s.String()
}

func (s *AddUserToSecurityCenterWhiteListRequest) SetRegionId(v string) *AddUserToSecurityCenterWhiteListRequest {
	s.RegionId = &v
	return s
}

type AddUserToSecurityCenterWhiteListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserToSecurityCenterWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserToSecurityCenterWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserToSecurityCenterWhiteListResponseBody) SetRequestId(v string) *AddUserToSecurityCenterWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type AddUserToSecurityCenterWhiteListResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddUserToSecurityCenterWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUserToSecurityCenterWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserToSecurityCenterWhiteListResponse) GoString() string {
	return s.String()
}

func (s *AddUserToSecurityCenterWhiteListResponse) SetHeaders(v map[string]*string) *AddUserToSecurityCenterWhiteListResponse {
	s.Headers = v
	return s
}

func (s *AddUserToSecurityCenterWhiteListResponse) SetBody(v *AddUserToSecurityCenterWhiteListResponseBody) *AddUserToSecurityCenterWhiteListResponse {
	s.Body = v
	return s
}

type AttachCenRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	CenId        *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
}

func (s AttachCenRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachCenRequest) GoString() string {
	return s.String()
}

func (s *AttachCenRequest) SetRegionId(v string) *AttachCenRequest {
	s.RegionId = &v
	return s
}

func (s *AttachCenRequest) SetCenId(v string) *AttachCenRequest {
	s.CenId = &v
	return s
}

func (s *AttachCenRequest) SetOfficeSiteId(v string) *AttachCenRequest {
	s.OfficeSiteId = &v
	return s
}

type AttachCenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachCenResponseBody) GoString() string {
	return s.String()
}

func (s *AttachCenResponseBody) SetRequestId(v string) *AttachCenResponseBody {
	s.RequestId = &v
	return s
}

type AttachCenResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachCenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachCenResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachCenResponse) GoString() string {
	return s.String()
}

func (s *AttachCenResponse) SetHeaders(v map[string]*string) *AttachCenResponse {
	s.Headers = v
	return s
}

func (s *AttachCenResponse) SetBody(v *AttachCenResponseBody) *AttachCenResponse {
	s.Body = v
	return s
}

type CheckUserInSecurityCenterWhiteListRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckUserInSecurityCenterWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckUserInSecurityCenterWhiteListRequest) GoString() string {
	return s.String()
}

func (s *CheckUserInSecurityCenterWhiteListRequest) SetRegionId(v string) *CheckUserInSecurityCenterWhiteListRequest {
	s.RegionId = &v
	return s
}

type CheckUserInSecurityCenterWhiteListResponseBody struct {
	InWhiteList *bool   `json:"InWhiteList,omitempty" xml:"InWhiteList,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckUserInSecurityCenterWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckUserInSecurityCenterWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserInSecurityCenterWhiteListResponseBody) SetInWhiteList(v bool) *CheckUserInSecurityCenterWhiteListResponseBody {
	s.InWhiteList = &v
	return s
}

func (s *CheckUserInSecurityCenterWhiteListResponseBody) SetRequestId(v string) *CheckUserInSecurityCenterWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type CheckUserInSecurityCenterWhiteListResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckUserInSecurityCenterWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckUserInSecurityCenterWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUserInSecurityCenterWhiteListResponse) GoString() string {
	return s.String()
}

func (s *CheckUserInSecurityCenterWhiteListResponse) SetHeaders(v map[string]*string) *CheckUserInSecurityCenterWhiteListResponse {
	s.Headers = v
	return s
}

func (s *CheckUserInSecurityCenterWhiteListResponse) SetBody(v *CheckUserInSecurityCenterWhiteListResponseBody) *CheckUserInSecurityCenterWhiteListResponse {
	s.Body = v
	return s
}

type ClonePolicyGroupRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ClonePolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ClonePolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *ClonePolicyGroupRequest) SetRegionId(v string) *ClonePolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ClonePolicyGroupRequest) SetPolicyGroupId(v string) *ClonePolicyGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ClonePolicyGroupRequest) SetName(v string) *ClonePolicyGroupRequest {
	s.Name = &v
	return s
}

type ClonePolicyGroupResponseBody struct {
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClonePolicyGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClonePolicyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ClonePolicyGroupResponseBody) SetPolicyGroupId(v string) *ClonePolicyGroupResponseBody {
	s.PolicyGroupId = &v
	return s
}

func (s *ClonePolicyGroupResponseBody) SetRequestId(v string) *ClonePolicyGroupResponseBody {
	s.RequestId = &v
	return s
}

type ClonePolicyGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClonePolicyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClonePolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ClonePolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *ClonePolicyGroupResponse) SetHeaders(v map[string]*string) *ClonePolicyGroupResponse {
	s.Headers = v
	return s
}

func (s *ClonePolicyGroupResponse) SetBody(v *ClonePolicyGroupResponseBody) *ClonePolicyGroupResponse {
	s.Body = v
	return s
}

type CreateADConnectorDirectoryRequest struct {
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DomainName          *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainUserName      *string   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	DomainPassword      *string   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	DirectoryName       *string   `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	EnableAdminAccess   *bool     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	DesktopAccessType   *string   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	SubDomainName       *string   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	MfaEnabled          *bool     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	DnsAddress          []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	VSwitchId           []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
	SubDomainDnsAddress []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
}

func (s CreateADConnectorDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryRequest) SetRegionId(v string) *CreateADConnectorDirectoryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDomainName(v string) *CreateADConnectorDirectoryRequest {
	s.DomainName = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDomainUserName(v string) *CreateADConnectorDirectoryRequest {
	s.DomainUserName = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDomainPassword(v string) *CreateADConnectorDirectoryRequest {
	s.DomainPassword = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDirectoryName(v string) *CreateADConnectorDirectoryRequest {
	s.DirectoryName = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetEnableAdminAccess(v bool) *CreateADConnectorDirectoryRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDesktopAccessType(v string) *CreateADConnectorDirectoryRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetSubDomainName(v string) *CreateADConnectorDirectoryRequest {
	s.SubDomainName = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetMfaEnabled(v bool) *CreateADConnectorDirectoryRequest {
	s.MfaEnabled = &v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetDnsAddress(v []*string) *CreateADConnectorDirectoryRequest {
	s.DnsAddress = v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetVSwitchId(v []*string) *CreateADConnectorDirectoryRequest {
	s.VSwitchId = v
	return s
}

func (s *CreateADConnectorDirectoryRequest) SetSubDomainDnsAddress(v []*string) *CreateADConnectorDirectoryRequest {
	s.SubDomainDnsAddress = v
	return s
}

type CreateADConnectorDirectoryResponseBody struct {
	TrustPassword *string                                               `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty"`
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DirectoryId   *string                                               `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	AdConnectors  []*CreateADConnectorDirectoryResponseBodyAdConnectors `json:"AdConnectors,omitempty" xml:"AdConnectors,omitempty" type:"Repeated"`
}

func (s CreateADConnectorDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryResponseBody) SetTrustPassword(v string) *CreateADConnectorDirectoryResponseBody {
	s.TrustPassword = &v
	return s
}

func (s *CreateADConnectorDirectoryResponseBody) SetRequestId(v string) *CreateADConnectorDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateADConnectorDirectoryResponseBody) SetDirectoryId(v string) *CreateADConnectorDirectoryResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateADConnectorDirectoryResponseBody) SetAdConnectors(v []*CreateADConnectorDirectoryResponseBodyAdConnectors) *CreateADConnectorDirectoryResponseBody {
	s.AdConnectors = v
	return s
}

type CreateADConnectorDirectoryResponseBodyAdConnectors struct {
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
}

func (s CreateADConnectorDirectoryResponseBodyAdConnectors) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorDirectoryResponseBodyAdConnectors) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryResponseBodyAdConnectors) SetAddress(v string) *CreateADConnectorDirectoryResponseBodyAdConnectors {
	s.Address = &v
	return s
}

type CreateADConnectorDirectoryResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateADConnectorDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateADConnectorDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateADConnectorDirectoryResponse) SetHeaders(v map[string]*string) *CreateADConnectorDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreateADConnectorDirectoryResponse) SetBody(v *CreateADConnectorDirectoryResponseBody) *CreateADConnectorDirectoryResponse {
	s.Body = v
	return s
}

type CreateADConnectorOfficeSiteRequest struct {
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	CidrBlock            *string   `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	CenId                *string   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	Bandwidth            *int32    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	DomainName           *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainUserName       *string   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	DomainPassword       *string   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	OfficeSiteName       *string   `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	EnableAdminAccess    *bool     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	DesktopAccessType    *string   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	EnableInternetAccess *bool     `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	SubDomainName        *string   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	MfaEnabled           *bool     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	DnsAddress           []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	SubDomainDnsAddress  []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
}

func (s CreateADConnectorOfficeSiteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorOfficeSiteRequest) GoString() string {
	return s.String()
}

func (s *CreateADConnectorOfficeSiteRequest) SetRegionId(v string) *CreateADConnectorOfficeSiteRequest {
	s.RegionId = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetCidrBlock(v string) *CreateADConnectorOfficeSiteRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetCenId(v string) *CreateADConnectorOfficeSiteRequest {
	s.CenId = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetBandwidth(v int32) *CreateADConnectorOfficeSiteRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDomainName(v string) *CreateADConnectorOfficeSiteRequest {
	s.DomainName = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDomainUserName(v string) *CreateADConnectorOfficeSiteRequest {
	s.DomainUserName = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDomainPassword(v string) *CreateADConnectorOfficeSiteRequest {
	s.DomainPassword = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetOfficeSiteName(v string) *CreateADConnectorOfficeSiteRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetEnableAdminAccess(v bool) *CreateADConnectorOfficeSiteRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDesktopAccessType(v string) *CreateADConnectorOfficeSiteRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetEnableInternetAccess(v bool) *CreateADConnectorOfficeSiteRequest {
	s.EnableInternetAccess = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetSubDomainName(v string) *CreateADConnectorOfficeSiteRequest {
	s.SubDomainName = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetMfaEnabled(v bool) *CreateADConnectorOfficeSiteRequest {
	s.MfaEnabled = &v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetDnsAddress(v []*string) *CreateADConnectorOfficeSiteRequest {
	s.DnsAddress = v
	return s
}

func (s *CreateADConnectorOfficeSiteRequest) SetSubDomainDnsAddress(v []*string) *CreateADConnectorOfficeSiteRequest {
	s.SubDomainDnsAddress = v
	return s
}

type CreateADConnectorOfficeSiteResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
}

func (s CreateADConnectorOfficeSiteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorOfficeSiteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateADConnectorOfficeSiteResponseBody) SetRequestId(v string) *CreateADConnectorOfficeSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateADConnectorOfficeSiteResponseBody) SetOfficeSiteId(v string) *CreateADConnectorOfficeSiteResponseBody {
	s.OfficeSiteId = &v
	return s
}

type CreateADConnectorOfficeSiteResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateADConnectorOfficeSiteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateADConnectorOfficeSiteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateADConnectorOfficeSiteResponse) GoString() string {
	return s.String()
}

func (s *CreateADConnectorOfficeSiteResponse) SetHeaders(v map[string]*string) *CreateADConnectorOfficeSiteResponse {
	s.Headers = v
	return s
}

func (s *CreateADConnectorOfficeSiteResponse) SetBody(v *CreateADConnectorOfficeSiteResponseBody) *CreateADConnectorOfficeSiteResponse {
	s.Body = v
	return s
}

type CreateBundleRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	DesktopType     *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	RootDiskSizeGib *int32  `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	BundleName      *string `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UserDiskSizeGib []*int  `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty" type:"Repeated"`
}

func (s CreateBundleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBundleRequest) GoString() string {
	return s.String()
}

func (s *CreateBundleRequest) SetRegionId(v string) *CreateBundleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBundleRequest) SetImageId(v string) *CreateBundleRequest {
	s.ImageId = &v
	return s
}

func (s *CreateBundleRequest) SetDesktopType(v string) *CreateBundleRequest {
	s.DesktopType = &v
	return s
}

func (s *CreateBundleRequest) SetRootDiskSizeGib(v int32) *CreateBundleRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *CreateBundleRequest) SetBundleName(v string) *CreateBundleRequest {
	s.BundleName = &v
	return s
}

func (s *CreateBundleRequest) SetDescription(v string) *CreateBundleRequest {
	s.Description = &v
	return s
}

func (s *CreateBundleRequest) SetUserDiskSizeGib(v []*int) *CreateBundleRequest {
	s.UserDiskSizeGib = v
	return s
}

type CreateBundleResponseBody struct {
	BundleId  *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBundleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBundleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBundleResponseBody) SetBundleId(v string) *CreateBundleResponseBody {
	s.BundleId = &v
	return s
}

func (s *CreateBundleResponseBody) SetRequestId(v string) *CreateBundleResponseBody {
	s.RequestId = &v
	return s
}

type CreateBundleResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBundleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBundleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBundleResponse) GoString() string {
	return s.String()
}

func (s *CreateBundleResponse) SetHeaders(v map[string]*string) *CreateBundleResponse {
	s.Headers = v
	return s
}

func (s *CreateBundleResponse) SetBody(v *CreateBundleResponseBody) *CreateBundleResponse {
	s.Body = v
	return s
}

type CreateDesktopGroupRequest struct {
	RegionId                *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BundleId                *string   `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	OfficeSiteId            *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	PolicyGroupId           *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	DesktopGroupName        *string   `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	DirectoryId             *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	ScaleStrategyId         *string   `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty"`
	VpcId                   *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	DefaultInitDesktopCount *int32    `json:"DefaultInitDesktopCount,omitempty" xml:"DefaultInitDesktopCount,omitempty"`
	KeepDuration            *int64    `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	ChargeType              *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Period                  *int32    `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit              *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	OwnType                 *int32    `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	AutoPay                 *bool     `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	Comments                *string   `json:"Comments,omitempty" xml:"Comments,omitempty"`
	MinDesktopsCount        *int32    `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	MaxDesktopsCount        *int32    `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	AllowAutoSetup          *int32    `json:"AllowAutoSetup,omitempty" xml:"AllowAutoSetup,omitempty"`
	AllowBufferCount        *int32    `json:"AllowBufferCount,omitempty" xml:"AllowBufferCount,omitempty"`
	ClientToken             *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndUserIds              []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
}

func (s CreateDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDesktopGroupRequest) SetRegionId(v string) *CreateDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetBundleId(v string) *CreateDesktopGroupRequest {
	s.BundleId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetOfficeSiteId(v string) *CreateDesktopGroupRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetPolicyGroupId(v string) *CreateDesktopGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDesktopGroupName(v string) *CreateDesktopGroupRequest {
	s.DesktopGroupName = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDirectoryId(v string) *CreateDesktopGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetScaleStrategyId(v string) *CreateDesktopGroupRequest {
	s.ScaleStrategyId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetVpcId(v string) *CreateDesktopGroupRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetDefaultInitDesktopCount(v int32) *CreateDesktopGroupRequest {
	s.DefaultInitDesktopCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetKeepDuration(v int64) *CreateDesktopGroupRequest {
	s.KeepDuration = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetChargeType(v string) *CreateDesktopGroupRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetPeriod(v int32) *CreateDesktopGroupRequest {
	s.Period = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetPeriodUnit(v string) *CreateDesktopGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetOwnType(v int32) *CreateDesktopGroupRequest {
	s.OwnType = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetAutoPay(v bool) *CreateDesktopGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetComments(v string) *CreateDesktopGroupRequest {
	s.Comments = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetMinDesktopsCount(v int32) *CreateDesktopGroupRequest {
	s.MinDesktopsCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetMaxDesktopsCount(v int32) *CreateDesktopGroupRequest {
	s.MaxDesktopsCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetAllowAutoSetup(v int32) *CreateDesktopGroupRequest {
	s.AllowAutoSetup = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetAllowBufferCount(v int32) *CreateDesktopGroupRequest {
	s.AllowBufferCount = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetClientToken(v string) *CreateDesktopGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDesktopGroupRequest) SetEndUserIds(v []*string) *CreateDesktopGroupRequest {
	s.EndUserIds = v
	return s
}

type CreateDesktopGroupResponseBody struct {
	DesktopGroupId *string   `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderIds       []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
}

func (s CreateDesktopGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDesktopGroupResponseBody) SetDesktopGroupId(v string) *CreateDesktopGroupResponseBody {
	s.DesktopGroupId = &v
	return s
}

func (s *CreateDesktopGroupResponseBody) SetRequestId(v string) *CreateDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDesktopGroupResponseBody) SetOrderIds(v []*string) *CreateDesktopGroupResponseBody {
	s.OrderIds = v
	return s
}

type CreateDesktopGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDesktopGroupResponse) SetHeaders(v map[string]*string) *CreateDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDesktopGroupResponse) SetBody(v *CreateDesktopGroupResponseBody) *CreateDesktopGroupResponse {
	s.Body = v
	return s
}

type CreateDesktopsRequest struct {
	RegionId       *string                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	GroupId        *string                     `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	BundleId       *string                     `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	DesktopName    *string                     `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	UserName       *string                     `json:"UserName,omitempty" xml:"UserName,omitempty"`
	VpcId          *string                     `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Amount         *int32                      `json:"Amount,omitempty" xml:"Amount,omitempty"`
	DirectoryId    *string                     `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OfficeSiteId   *string                     `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	PolicyGroupId  *string                     `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	ChargeType     *string                     `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Period         *int32                      `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit     *string                     `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoPay        *bool                       `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew      *bool                       `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	PromotionId    *string                     `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	UserAssignMode *string                     `json:"UserAssignMode,omitempty" xml:"UserAssignMode,omitempty"`
	EndUserId      []*string                   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	Tag            []*CreateDesktopsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsRequest) GoString() string {
	return s.String()
}

func (s *CreateDesktopsRequest) SetRegionId(v string) *CreateDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDesktopsRequest) SetGroupId(v string) *CreateDesktopsRequest {
	s.GroupId = &v
	return s
}

func (s *CreateDesktopsRequest) SetBundleId(v string) *CreateDesktopsRequest {
	s.BundleId = &v
	return s
}

func (s *CreateDesktopsRequest) SetDesktopName(v string) *CreateDesktopsRequest {
	s.DesktopName = &v
	return s
}

func (s *CreateDesktopsRequest) SetUserName(v string) *CreateDesktopsRequest {
	s.UserName = &v
	return s
}

func (s *CreateDesktopsRequest) SetVpcId(v string) *CreateDesktopsRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDesktopsRequest) SetAmount(v int32) *CreateDesktopsRequest {
	s.Amount = &v
	return s
}

func (s *CreateDesktopsRequest) SetDirectoryId(v string) *CreateDesktopsRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateDesktopsRequest) SetOfficeSiteId(v string) *CreateDesktopsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateDesktopsRequest) SetPolicyGroupId(v string) *CreateDesktopsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateDesktopsRequest) SetChargeType(v string) *CreateDesktopsRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDesktopsRequest) SetPeriod(v int32) *CreateDesktopsRequest {
	s.Period = &v
	return s
}

func (s *CreateDesktopsRequest) SetPeriodUnit(v string) *CreateDesktopsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDesktopsRequest) SetAutoPay(v bool) *CreateDesktopsRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateDesktopsRequest) SetAutoRenew(v bool) *CreateDesktopsRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDesktopsRequest) SetPromotionId(v string) *CreateDesktopsRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateDesktopsRequest) SetUserAssignMode(v string) *CreateDesktopsRequest {
	s.UserAssignMode = &v
	return s
}

func (s *CreateDesktopsRequest) SetEndUserId(v []*string) *CreateDesktopsRequest {
	s.EndUserId = v
	return s
}

func (s *CreateDesktopsRequest) SetTag(v []*CreateDesktopsRequestTag) *CreateDesktopsRequest {
	s.Tag = v
	return s
}

type CreateDesktopsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDesktopsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDesktopsRequestTag) SetKey(v string) *CreateDesktopsRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDesktopsRequestTag) SetValue(v string) *CreateDesktopsRequestTag {
	s.Value = &v
	return s
}

type CreateDesktopsResponseBody struct {
	OrderId   *string   `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s CreateDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDesktopsResponseBody) SetOrderId(v string) *CreateDesktopsResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDesktopsResponseBody) SetRequestId(v string) *CreateDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDesktopsResponseBody) SetDesktopId(v []*string) *CreateDesktopsResponseBody {
	s.DesktopId = v
	return s
}

type CreateDesktopsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsResponse) GoString() string {
	return s.String()
}

func (s *CreateDesktopsResponse) SetHeaders(v map[string]*string) *CreateDesktopsResponse {
	s.Headers = v
	return s
}

func (s *CreateDesktopsResponse) SetBody(v *CreateDesktopsResponseBody) *CreateDesktopsResponse {
	s.Body = v
	return s
}

type CreateDesktopsLiteRequest struct {
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BundleId             *string   `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	UserAssignMode       *string   `json:"UserAssignMode,omitempty" xml:"UserAssignMode,omitempty"`
	Amount               *int32    `json:"Amount,omitempty" xml:"Amount,omitempty"`
	EnableInternetAccess *bool     `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	Bandwidth            *int32    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	PeriodUnit           *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Period               *int32    `json:"Period,omitempty" xml:"Period,omitempty"`
	EndUserId            []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
}

func (s CreateDesktopsLiteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsLiteRequest) GoString() string {
	return s.String()
}

func (s *CreateDesktopsLiteRequest) SetRegionId(v string) *CreateDesktopsLiteRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDesktopsLiteRequest) SetBundleId(v string) *CreateDesktopsLiteRequest {
	s.BundleId = &v
	return s
}

func (s *CreateDesktopsLiteRequest) SetUserAssignMode(v string) *CreateDesktopsLiteRequest {
	s.UserAssignMode = &v
	return s
}

func (s *CreateDesktopsLiteRequest) SetAmount(v int32) *CreateDesktopsLiteRequest {
	s.Amount = &v
	return s
}

func (s *CreateDesktopsLiteRequest) SetEnableInternetAccess(v bool) *CreateDesktopsLiteRequest {
	s.EnableInternetAccess = &v
	return s
}

func (s *CreateDesktopsLiteRequest) SetBandwidth(v int32) *CreateDesktopsLiteRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateDesktopsLiteRequest) SetPeriodUnit(v string) *CreateDesktopsLiteRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDesktopsLiteRequest) SetPeriod(v int32) *CreateDesktopsLiteRequest {
	s.Period = &v
	return s
}

func (s *CreateDesktopsLiteRequest) SetEndUserId(v []*string) *CreateDesktopsLiteRequest {
	s.EndUserId = v
	return s
}

type CreateDesktopsLiteResponseBody struct {
	OrderId   *string   `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s CreateDesktopsLiteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsLiteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDesktopsLiteResponseBody) SetOrderId(v string) *CreateDesktopsLiteResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDesktopsLiteResponseBody) SetRequestId(v string) *CreateDesktopsLiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDesktopsLiteResponseBody) SetDesktopId(v []*string) *CreateDesktopsLiteResponseBody {
	s.DesktopId = v
	return s
}

type CreateDesktopsLiteResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDesktopsLiteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDesktopsLiteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsLiteResponse) GoString() string {
	return s.String()
}

func (s *CreateDesktopsLiteResponse) SetHeaders(v map[string]*string) *CreateDesktopsLiteResponse {
	s.Headers = v
	return s
}

func (s *CreateDesktopsLiteResponse) SetBody(v *CreateDesktopsLiteResponseBody) *CreateDesktopsLiteResponse {
	s.Body = v
	return s
}

type CreateImageRequest struct {
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopId         *string   `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	ImageName         *string   `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	Description       *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	SnapshotId        *string   `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	ImageResourceType *string   `json:"ImageResourceType,omitempty" xml:"ImageResourceType,omitempty"`
	SnapshotIds       []*string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty" type:"Repeated"`
}

func (s CreateImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageRequest) GoString() string {
	return s.String()
}

func (s *CreateImageRequest) SetRegionId(v string) *CreateImageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageRequest) SetDesktopId(v string) *CreateImageRequest {
	s.DesktopId = &v
	return s
}

func (s *CreateImageRequest) SetImageName(v string) *CreateImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateImageRequest) SetDescription(v string) *CreateImageRequest {
	s.Description = &v
	return s
}

func (s *CreateImageRequest) SetSnapshotId(v string) *CreateImageRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateImageRequest) SetImageResourceType(v string) *CreateImageRequest {
	s.ImageResourceType = &v
	return s
}

func (s *CreateImageRequest) SetSnapshotIds(v []*string) *CreateImageRequest {
	s.SnapshotIds = v
	return s
}

type CreateImageResponseBody struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageResponseBody) SetImageId(v string) *CreateImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateImageResponseBody) SetRequestId(v string) *CreateImageResponseBody {
	s.RequestId = &v
	return s
}

type CreateImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponse) GoString() string {
	return s.String()
}

func (s *CreateImageResponse) SetHeaders(v map[string]*string) *CreateImageResponse {
	s.Headers = v
	return s
}

func (s *CreateImageResponse) SetBody(v *CreateImageResponseBody) *CreateImageResponse {
	s.Body = v
	return s
}

type CreateNASFileSystemRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateNASFileSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNASFileSystemRequest) GoString() string {
	return s.String()
}

func (s *CreateNASFileSystemRequest) SetRegionId(v string) *CreateNASFileSystemRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNASFileSystemRequest) SetOfficeSiteId(v string) *CreateNASFileSystemRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateNASFileSystemRequest) SetName(v string) *CreateNASFileSystemRequest {
	s.Name = &v
	return s
}

func (s *CreateNASFileSystemRequest) SetDescription(v string) *CreateNASFileSystemRequest {
	s.Description = &v
	return s
}

type CreateNASFileSystemResponseBody struct {
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemName    *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OfficeSiteId      *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
}

func (s CreateNASFileSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNASFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNASFileSystemResponseBody) SetFileSystemId(v string) *CreateNASFileSystemResponseBody {
	s.FileSystemId = &v
	return s
}

func (s *CreateNASFileSystemResponseBody) SetFileSystemName(v string) *CreateNASFileSystemResponseBody {
	s.FileSystemName = &v
	return s
}

func (s *CreateNASFileSystemResponseBody) SetMountTargetDomain(v string) *CreateNASFileSystemResponseBody {
	s.MountTargetDomain = &v
	return s
}

func (s *CreateNASFileSystemResponseBody) SetRequestId(v string) *CreateNASFileSystemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNASFileSystemResponseBody) SetOfficeSiteId(v string) *CreateNASFileSystemResponseBody {
	s.OfficeSiteId = &v
	return s
}

type CreateNASFileSystemResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateNASFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateNASFileSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNASFileSystemResponse) GoString() string {
	return s.String()
}

func (s *CreateNASFileSystemResponse) SetHeaders(v map[string]*string) *CreateNASFileSystemResponse {
	s.Headers = v
	return s
}

func (s *CreateNASFileSystemResponse) SetBody(v *CreateNASFileSystemResponseBody) *CreateNASFileSystemResponse {
	s.Body = v
	return s
}

type CreateNetworkPackageRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Bandwidth          *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	OfficeSiteId       *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	Period             *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit         *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoPay            *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew          *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
}

func (s CreateNetworkPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkPackageRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkPackageRequest) SetRegionId(v string) *CreateNetworkPackageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetBandwidth(v int32) *CreateNetworkPackageRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetOfficeSiteId(v string) *CreateNetworkPackageRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetInternetChargeType(v string) *CreateNetworkPackageRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetPeriod(v int32) *CreateNetworkPackageRequest {
	s.Period = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetPeriodUnit(v string) *CreateNetworkPackageRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetAutoPay(v bool) *CreateNetworkPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetAutoRenew(v bool) *CreateNetworkPackageRequest {
	s.AutoRenew = &v
	return s
}

type CreateNetworkPackageResponseBody struct {
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId          *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateNetworkPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkPackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkPackageResponseBody) SetNetworkPackageId(v string) *CreateNetworkPackageResponseBody {
	s.NetworkPackageId = &v
	return s
}

func (s *CreateNetworkPackageResponseBody) SetRequestId(v string) *CreateNetworkPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkPackageResponseBody) SetOrderId(v string) *CreateNetworkPackageResponseBody {
	s.OrderId = &v
	return s
}

type CreateNetworkPackageResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateNetworkPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateNetworkPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkPackageResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkPackageResponse) SetHeaders(v map[string]*string) *CreateNetworkPackageResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkPackageResponse) SetBody(v *CreateNetworkPackageResponseBody) *CreateNetworkPackageResponse {
	s.Body = v
	return s
}

type CreatePolicyGroupRequest struct {
	RegionId                    *string                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Clipboard                   *string                                                `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	LocalDrive                  *string                                                `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UsbRedirect                 *string                                                `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	VisualQuality               *string                                                `json:"VisualQuality,omitempty" xml:"VisualQuality,omitempty"`
	Html5Access                 *string                                                `json:"Html5Access,omitempty" xml:"Html5Access,omitempty"`
	Html5FileTransfer           *string                                                `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	Watermark                   *string                                                `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	Name                        *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	WatermarkType               *string                                                `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	WatermarkCustomText         *string                                                `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	WatermarkTransparency       *string                                                `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
	PreemptLogin                *string                                                `json:"PreemptLogin,omitempty" xml:"PreemptLogin,omitempty"`
	DomainList                  *string                                                `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	PreemptLoginUser            []*string                                              `json:"PreemptLoginUser,omitempty" xml:"PreemptLoginUser,omitempty" type:"Repeated"`
	AuthorizeSecurityPolicyRule []*CreatePolicyGroupRequestAuthorizeSecurityPolicyRule `json:"AuthorizeSecurityPolicyRule,omitempty" xml:"AuthorizeSecurityPolicyRule,omitempty" type:"Repeated"`
	AuthorizeAccessPolicyRule   []*CreatePolicyGroupRequestAuthorizeAccessPolicyRule   `json:"AuthorizeAccessPolicyRule,omitempty" xml:"AuthorizeAccessPolicyRule,omitempty" type:"Repeated"`
}

func (s CreatePolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequest) SetRegionId(v string) *CreatePolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetClipboard(v string) *CreatePolicyGroupRequest {
	s.Clipboard = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetLocalDrive(v string) *CreatePolicyGroupRequest {
	s.LocalDrive = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetUsbRedirect(v string) *CreatePolicyGroupRequest {
	s.UsbRedirect = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetVisualQuality(v string) *CreatePolicyGroupRequest {
	s.VisualQuality = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetHtml5Access(v string) *CreatePolicyGroupRequest {
	s.Html5Access = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetHtml5FileTransfer(v string) *CreatePolicyGroupRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermark(v string) *CreatePolicyGroupRequest {
	s.Watermark = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetName(v string) *CreatePolicyGroupRequest {
	s.Name = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkType(v string) *CreatePolicyGroupRequest {
	s.WatermarkType = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkCustomText(v string) *CreatePolicyGroupRequest {
	s.WatermarkCustomText = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetWatermarkTransparency(v string) *CreatePolicyGroupRequest {
	s.WatermarkTransparency = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetPreemptLogin(v string) *CreatePolicyGroupRequest {
	s.PreemptLogin = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetDomainList(v string) *CreatePolicyGroupRequest {
	s.DomainList = &v
	return s
}

func (s *CreatePolicyGroupRequest) SetPreemptLoginUser(v []*string) *CreatePolicyGroupRequest {
	s.PreemptLoginUser = v
	return s
}

func (s *CreatePolicyGroupRequest) SetAuthorizeSecurityPolicyRule(v []*CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) *CreatePolicyGroupRequest {
	s.AuthorizeSecurityPolicyRule = v
	return s
}

func (s *CreatePolicyGroupRequest) SetAuthorizeAccessPolicyRule(v []*CreatePolicyGroupRequestAuthorizeAccessPolicyRule) *CreatePolicyGroupRequest {
	s.AuthorizeAccessPolicyRule = v
	return s
}

type CreatePolicyGroupRequestAuthorizeSecurityPolicyRule struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PortRange   *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IpProtocol  *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Priority    *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
}

func (s CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetType(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Type = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetPolicy(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Policy = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetPortRange(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.PortRange = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetDescription(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Description = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetIpProtocol(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.IpProtocol = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetPriority(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Priority = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule) SetCidrIp(v string) *CreatePolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.CidrIp = &v
	return s
}

type CreatePolicyGroupRequestAuthorizeAccessPolicyRule struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
}

func (s CreatePolicyGroupRequestAuthorizeAccessPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupRequestAuthorizeAccessPolicyRule) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupRequestAuthorizeAccessPolicyRule) SetDescription(v string) *CreatePolicyGroupRequestAuthorizeAccessPolicyRule {
	s.Description = &v
	return s
}

func (s *CreatePolicyGroupRequestAuthorizeAccessPolicyRule) SetCidrIp(v string) *CreatePolicyGroupRequestAuthorizeAccessPolicyRule {
	s.CidrIp = &v
	return s
}

type CreatePolicyGroupResponseBody struct {
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolicyGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupResponseBody) SetPolicyGroupId(v string) *CreatePolicyGroupResponseBody {
	s.PolicyGroupId = &v
	return s
}

func (s *CreatePolicyGroupResponseBody) SetRequestId(v string) *CreatePolicyGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreatePolicyGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePolicyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupResponse) SetHeaders(v map[string]*string) *CreatePolicyGroupResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicyGroupResponse) SetBody(v *CreatePolicyGroupResponseBody) *CreatePolicyGroupResponse {
	s.Body = v
	return s
}

type CreateRAMDirectoryRequest struct {
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DirectoryName        *string   `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	EnableInternetAccess *bool     `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	EnableAdminAccess    *bool     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	DesktopAccessType    *string   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	VSwitchId            []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
}

func (s CreateRAMDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRAMDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateRAMDirectoryRequest) SetRegionId(v string) *CreateRAMDirectoryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetDirectoryName(v string) *CreateRAMDirectoryRequest {
	s.DirectoryName = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetEnableInternetAccess(v bool) *CreateRAMDirectoryRequest {
	s.EnableInternetAccess = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetEnableAdminAccess(v bool) *CreateRAMDirectoryRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetDesktopAccessType(v string) *CreateRAMDirectoryRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *CreateRAMDirectoryRequest) SetVSwitchId(v []*string) *CreateRAMDirectoryRequest {
	s.VSwitchId = v
	return s
}

type CreateRAMDirectoryResponseBody struct {
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRAMDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRAMDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRAMDirectoryResponseBody) SetDirectoryId(v string) *CreateRAMDirectoryResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateRAMDirectoryResponseBody) SetRequestId(v string) *CreateRAMDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type CreateRAMDirectoryResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRAMDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRAMDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRAMDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateRAMDirectoryResponse) SetHeaders(v map[string]*string) *CreateRAMDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreateRAMDirectoryResponse) SetBody(v *CreateRAMDirectoryResponseBody) *CreateRAMDirectoryResponse {
	s.Body = v
	return s
}

type CreateScaleStrategyRequest struct {
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScaleStrategyName         *string `json:"ScaleStrategyName,omitempty" xml:"ScaleStrategyName,omitempty"`
	ScaleStrategyType         *string `json:"ScaleStrategyType,omitempty" xml:"ScaleStrategyType,omitempty"`
	PayType                   *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	MinDesktopsCount          *int32  `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	MaxDesktopsCount          *int32  `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	MinAvailableDesktopsCount *int32  `json:"MinAvailableDesktopsCount,omitempty" xml:"MinAvailableDesktopsCount,omitempty"`
	MaxAvailableDesktopsCount *int32  `json:"MaxAvailableDesktopsCount,omitempty" xml:"MaxAvailableDesktopsCount,omitempty"`
	ScaleStep                 *int32  `json:"ScaleStep,omitempty" xml:"ScaleStep,omitempty"`
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateScaleStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScaleStrategyRequest) GoString() string {
	return s.String()
}

func (s *CreateScaleStrategyRequest) SetRegionId(v string) *CreateScaleStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetScaleStrategyName(v string) *CreateScaleStrategyRequest {
	s.ScaleStrategyName = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetScaleStrategyType(v string) *CreateScaleStrategyRequest {
	s.ScaleStrategyType = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetPayType(v string) *CreateScaleStrategyRequest {
	s.PayType = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetMinDesktopsCount(v int32) *CreateScaleStrategyRequest {
	s.MinDesktopsCount = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetMaxDesktopsCount(v int32) *CreateScaleStrategyRequest {
	s.MaxDesktopsCount = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetMinAvailableDesktopsCount(v int32) *CreateScaleStrategyRequest {
	s.MinAvailableDesktopsCount = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetMaxAvailableDesktopsCount(v int32) *CreateScaleStrategyRequest {
	s.MaxAvailableDesktopsCount = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetScaleStep(v int32) *CreateScaleStrategyRequest {
	s.ScaleStep = &v
	return s
}

func (s *CreateScaleStrategyRequest) SetClientToken(v string) *CreateScaleStrategyRequest {
	s.ClientToken = &v
	return s
}

type CreateScaleStrategyResponseBody struct {
	ScaleStrategyId *string `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateScaleStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScaleStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScaleStrategyResponseBody) SetScaleStrategyId(v string) *CreateScaleStrategyResponseBody {
	s.ScaleStrategyId = &v
	return s
}

func (s *CreateScaleStrategyResponseBody) SetRequestId(v string) *CreateScaleStrategyResponseBody {
	s.RequestId = &v
	return s
}

type CreateScaleStrategyResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateScaleStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateScaleStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScaleStrategyResponse) GoString() string {
	return s.String()
}

func (s *CreateScaleStrategyResponse) SetHeaders(v map[string]*string) *CreateScaleStrategyResponse {
	s.Headers = v
	return s
}

func (s *CreateScaleStrategyResponse) SetBody(v *CreateScaleStrategyResponseBody) *CreateScaleStrategyResponse {
	s.Body = v
	return s
}

type CreateServiceLinkedRoleRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) SetRegionId(v string) *CreateServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

type CreateServiceLinkedRoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceLinkedRoleResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *CreateServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceLinkedRoleResponse) SetBody(v *CreateServiceLinkedRoleResponseBody) *CreateServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type CreateSimpleOfficeSiteRequest struct {
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	CidrBlock            *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	Bandwidth            *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	OfficeSiteName       *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	EnableInternetAccess *bool   `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	EnableAdminAccess    *bool   `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	DesktopAccessType    *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
}

func (s CreateSimpleOfficeSiteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSimpleOfficeSiteRequest) GoString() string {
	return s.String()
}

func (s *CreateSimpleOfficeSiteRequest) SetRegionId(v string) *CreateSimpleOfficeSiteRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetCidrBlock(v string) *CreateSimpleOfficeSiteRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetCenId(v string) *CreateSimpleOfficeSiteRequest {
	s.CenId = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetBandwidth(v int32) *CreateSimpleOfficeSiteRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetOfficeSiteName(v string) *CreateSimpleOfficeSiteRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetEnableInternetAccess(v bool) *CreateSimpleOfficeSiteRequest {
	s.EnableInternetAccess = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetEnableAdminAccess(v bool) *CreateSimpleOfficeSiteRequest {
	s.EnableAdminAccess = &v
	return s
}

func (s *CreateSimpleOfficeSiteRequest) SetDesktopAccessType(v string) *CreateSimpleOfficeSiteRequest {
	s.DesktopAccessType = &v
	return s
}

type CreateSimpleOfficeSiteResponseBody struct {
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSimpleOfficeSiteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSimpleOfficeSiteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSimpleOfficeSiteResponseBody) SetOfficeSiteId(v string) *CreateSimpleOfficeSiteResponseBody {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateSimpleOfficeSiteResponseBody) SetRequestId(v string) *CreateSimpleOfficeSiteResponseBody {
	s.RequestId = &v
	return s
}

type CreateSimpleOfficeSiteResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSimpleOfficeSiteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSimpleOfficeSiteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSimpleOfficeSiteResponse) GoString() string {
	return s.String()
}

func (s *CreateSimpleOfficeSiteResponse) SetHeaders(v map[string]*string) *CreateSimpleOfficeSiteResponse {
	s.Headers = v
	return s
}

func (s *CreateSimpleOfficeSiteResponse) SetBody(v *CreateSimpleOfficeSiteResponseBody) *CreateSimpleOfficeSiteResponse {
	s.Body = v
	return s
}

type CreateSnapshotRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	SnapshotName   *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) SetRegionId(v string) *CreateSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSnapshotRequest) SetDesktopId(v string) *CreateSnapshotRequest {
	s.DesktopId = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

func (s *CreateSnapshotRequest) SetDescription(v string) *CreateSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateSnapshotRequest) SetSourceDiskType(v string) *CreateSnapshotRequest {
	s.SourceDiskType = &v
	return s
}

type CreateSnapshotResponseBody struct {
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBody) SetSnapshotId(v string) *CreateSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetRequestId(v string) *CreateSnapshotResponseBody {
	s.RequestId = &v
	return s
}

type CreateSnapshotResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponse) SetHeaders(v map[string]*string) *CreateSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateSnapshotResponse) SetBody(v *CreateSnapshotResponseBody) *CreateSnapshotResponse {
	s.Body = v
	return s
}

type DeleteBundlesRequest struct {
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BundleId []*string `json:"BundleId,omitempty" xml:"BundleId,omitempty" type:"Repeated"`
}

func (s DeleteBundlesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBundlesRequest) GoString() string {
	return s.String()
}

func (s *DeleteBundlesRequest) SetRegionId(v string) *DeleteBundlesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBundlesRequest) SetBundleId(v []*string) *DeleteBundlesRequest {
	s.BundleId = v
	return s
}

type DeleteBundlesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBundlesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBundlesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBundlesResponseBody) SetRequestId(v string) *DeleteBundlesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBundlesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteBundlesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBundlesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBundlesResponse) GoString() string {
	return s.String()
}

func (s *DeleteBundlesResponse) SetHeaders(v map[string]*string) *DeleteBundlesResponse {
	s.Headers = v
	return s
}

func (s *DeleteBundlesResponse) SetBody(v *DeleteBundlesResponseBody) *DeleteBundlesResponse {
	s.Body = v
	return s
}

type DeleteDesktopGroupRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
}

func (s DeleteDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDesktopGroupRequest) SetRegionId(v string) *DeleteDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDesktopGroupRequest) SetDesktopGroupId(v string) *DeleteDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

type DeleteDesktopGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDesktopGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDesktopGroupResponseBody) SetRequestId(v string) *DeleteDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDesktopGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDesktopGroupResponse) SetHeaders(v map[string]*string) *DeleteDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDesktopGroupResponse) SetBody(v *DeleteDesktopGroupResponseBody) *DeleteDesktopGroupResponse {
	s.Body = v
	return s
}

type DeleteDesktopsRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s DeleteDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDesktopsRequest) SetRegionId(v string) *DeleteDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDesktopsRequest) SetDesktopId(v []*string) *DeleteDesktopsRequest {
	s.DesktopId = v
	return s
}

type DeleteDesktopsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDesktopsResponseBody) SetRequestId(v string) *DeleteDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDesktopsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDesktopsResponse) SetHeaders(v map[string]*string) *DeleteDesktopsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDesktopsResponse) SetBody(v *DeleteDesktopsResponseBody) *DeleteDesktopsResponse {
	s.Body = v
	return s
}

type DeleteDirectoriesRequest struct {
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DirectoryId []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" type:"Repeated"`
}

func (s DeleteDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DeleteDirectoriesRequest) SetRegionId(v string) *DeleteDirectoriesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDirectoriesRequest) SetDirectoryId(v []*string) *DeleteDirectoriesRequest {
	s.DirectoryId = v
	return s
}

type DeleteDirectoriesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDirectoriesResponseBody) SetRequestId(v string) *DeleteDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDirectoriesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *DeleteDirectoriesResponse) SetHeaders(v map[string]*string) *DeleteDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *DeleteDirectoriesResponse) SetBody(v *DeleteDirectoriesResponseBody) *DeleteDirectoriesResponse {
	s.Body = v
	return s
}

type DeleteImagesRequest struct {
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ImageId  []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
}

func (s DeleteImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesRequest) GoString() string {
	return s.String()
}

func (s *DeleteImagesRequest) SetRegionId(v string) *DeleteImagesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteImagesRequest) SetImageId(v []*string) *DeleteImagesRequest {
	s.ImageId = v
	return s
}

type DeleteImagesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImagesResponseBody) SetRequestId(v string) *DeleteImagesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImagesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesResponse) GoString() string {
	return s.String()
}

func (s *DeleteImagesResponse) SetHeaders(v map[string]*string) *DeleteImagesResponse {
	s.Headers = v
	return s
}

func (s *DeleteImagesResponse) SetBody(v *DeleteImagesResponseBody) *DeleteImagesResponse {
	s.Body = v
	return s
}

type DeleteNASFileSystemsRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FileSystemId []*string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty" type:"Repeated"`
}

func (s DeleteNASFileSystemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNASFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *DeleteNASFileSystemsRequest) SetRegionId(v string) *DeleteNASFileSystemsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNASFileSystemsRequest) SetFileSystemId(v []*string) *DeleteNASFileSystemsRequest {
	s.FileSystemId = v
	return s
}

type DeleteNASFileSystemsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNASFileSystemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNASFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNASFileSystemsResponseBody) SetRequestId(v string) *DeleteNASFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNASFileSystemsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNASFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNASFileSystemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNASFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *DeleteNASFileSystemsResponse) SetHeaders(v map[string]*string) *DeleteNASFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *DeleteNASFileSystemsResponse) SetBody(v *DeleteNASFileSystemsResponseBody) *DeleteNASFileSystemsResponse {
	s.Body = v
	return s
}

type DeleteNetworkPackagesRequest struct {
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	NetworkPackageId []*string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" type:"Repeated"`
}

func (s DeleteNetworkPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkPackagesRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPackagesRequest) SetRegionId(v string) *DeleteNetworkPackagesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNetworkPackagesRequest) SetNetworkPackageId(v []*string) *DeleteNetworkPackagesRequest {
	s.NetworkPackageId = v
	return s
}

type DeleteNetworkPackagesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPackagesResponseBody) SetRequestId(v string) *DeleteNetworkPackagesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNetworkPackagesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNetworkPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNetworkPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkPackagesResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPackagesResponse) SetHeaders(v map[string]*string) *DeleteNetworkPackagesResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkPackagesResponse) SetBody(v *DeleteNetworkPackagesResponseBody) *DeleteNetworkPackagesResponse {
	s.Body = v
	return s
}

type DeleteOfficeSitesRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OfficeSiteId []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
}

func (s DeleteOfficeSitesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOfficeSitesRequest) GoString() string {
	return s.String()
}

func (s *DeleteOfficeSitesRequest) SetRegionId(v string) *DeleteOfficeSitesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteOfficeSitesRequest) SetOfficeSiteId(v []*string) *DeleteOfficeSitesRequest {
	s.OfficeSiteId = v
	return s
}

type DeleteOfficeSitesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOfficeSitesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOfficeSitesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOfficeSitesResponseBody) SetRequestId(v string) *DeleteOfficeSitesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteOfficeSitesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteOfficeSitesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteOfficeSitesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOfficeSitesResponse) GoString() string {
	return s.String()
}

func (s *DeleteOfficeSitesResponse) SetHeaders(v map[string]*string) *DeleteOfficeSitesResponse {
	s.Headers = v
	return s
}

func (s *DeleteOfficeSitesResponse) SetBody(v *DeleteOfficeSitesResponseBody) *DeleteOfficeSitesResponse {
	s.Body = v
	return s
}

type DeletePolicyGroupsRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PolicyGroupId []*string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" type:"Repeated"`
}

func (s DeletePolicyGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyGroupsRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupsRequest) SetRegionId(v string) *DeletePolicyGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePolicyGroupsRequest) SetPolicyGroupId(v []*string) *DeletePolicyGroupsRequest {
	s.PolicyGroupId = v
	return s
}

type DeletePolicyGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolicyGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupsResponseBody) SetRequestId(v string) *DeletePolicyGroupsResponseBody {
	s.RequestId = &v
	return s
}

type DeletePolicyGroupsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePolicyGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePolicyGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyGroupsResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyGroupsResponse) SetHeaders(v map[string]*string) *DeletePolicyGroupsResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyGroupsResponse) SetBody(v *DeletePolicyGroupsResponseBody) *DeletePolicyGroupsResponse {
	s.Body = v
	return s
}

type DeleteScaleStrategyRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScaleStrategyId *string `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty"`
}

func (s DeleteScaleStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScaleStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteScaleStrategyRequest) SetRegionId(v string) *DeleteScaleStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteScaleStrategyRequest) SetScaleStrategyId(v string) *DeleteScaleStrategyRequest {
	s.ScaleStrategyId = &v
	return s
}

type DeleteScaleStrategyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScaleStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScaleStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScaleStrategyResponseBody) SetRequestId(v string) *DeleteScaleStrategyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScaleStrategyResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteScaleStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScaleStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScaleStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteScaleStrategyResponse) SetHeaders(v map[string]*string) *DeleteScaleStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteScaleStrategyResponse) SetBody(v *DeleteScaleStrategyResponseBody) *DeleteScaleStrategyResponse {
	s.Body = v
	return s
}

type DeleteSnapshotRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotId []*string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty" type:"Repeated"`
}

func (s DeleteSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) SetRegionId(v string) *DeleteSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v []*string) *DeleteSnapshotRequest {
	s.SnapshotId = v
	return s
}

type DeleteSnapshotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponseBody) SetRequestId(v string) *DeleteSnapshotResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSnapshotResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponse) SetHeaders(v map[string]*string) *DeleteSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotResponse) SetBody(v *DeleteSnapshotResponseBody) *DeleteSnapshotResponse {
	s.Body = v
	return s
}

type DeleteVirtualMFADeviceRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s DeleteVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceRequest) SetRegionId(v string) *DeleteVirtualMFADeviceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVirtualMFADeviceRequest) SetSerialNumber(v string) *DeleteVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

type DeleteVirtualMFADeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVirtualMFADeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceResponseBody) SetRequestId(v string) *DeleteVirtualMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVirtualMFADeviceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVirtualMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceResponse) SetHeaders(v map[string]*string) *DeleteVirtualMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirtualMFADeviceResponse) SetBody(v *DeleteVirtualMFADeviceResponseBody) *DeleteVirtualMFADeviceResponse {
	s.Body = v
	return s
}

type DescribeAlarmEventStackInfoRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopId  *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	EventName  *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	UniqueInfo *string `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeAlarmEventStackInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventStackInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventStackInfoRequest) SetRegionId(v string) *DescribeAlarmEventStackInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetDesktopId(v string) *DescribeAlarmEventStackInfoRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetEventName(v string) *DescribeAlarmEventStackInfoRequest {
	s.EventName = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetUniqueInfo(v string) *DescribeAlarmEventStackInfoRequest {
	s.UniqueInfo = &v
	return s
}

func (s *DescribeAlarmEventStackInfoRequest) SetLang(v string) *DescribeAlarmEventStackInfoRequest {
	s.Lang = &v
	return s
}

type DescribeAlarmEventStackInfoResponseBody struct {
	StackInfo *string `json:"StackInfo,omitempty" xml:"StackInfo,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAlarmEventStackInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventStackInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventStackInfoResponseBody) SetStackInfo(v string) *DescribeAlarmEventStackInfoResponseBody {
	s.StackInfo = &v
	return s
}

func (s *DescribeAlarmEventStackInfoResponseBody) SetRequestId(v string) *DescribeAlarmEventStackInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAlarmEventStackInfoResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAlarmEventStackInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAlarmEventStackInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventStackInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventStackInfoResponse) SetHeaders(v map[string]*string) *DescribeAlarmEventStackInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlarmEventStackInfoResponse) SetBody(v *DescribeAlarmEventStackInfoResponseBody) *DescribeAlarmEventStackInfoResponse {
	s.Body = v
	return s
}

type DescribeBundlesRequest struct {
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults        *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	BundleType        *string   `json:"BundleType,omitempty" xml:"BundleType,omitempty"`
	DesktopTypeFamily *string   `json:"DesktopTypeFamily,omitempty" xml:"DesktopTypeFamily,omitempty"`
	CpuCount          *int32    `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	MemorySize        *int32    `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	GpuCount          *float32  `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	CheckStock        *bool     `json:"CheckStock,omitempty" xml:"CheckStock,omitempty"`
	FromDesktopGroup  *bool     `json:"FromDesktopGroup,omitempty" xml:"FromDesktopGroup,omitempty"`
	ProtocolType      *string   `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	BundleId          []*string `json:"BundleId,omitempty" xml:"BundleId,omitempty" type:"Repeated"`
}

func (s DescribeBundlesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBundlesRequest) SetRegionId(v string) *DescribeBundlesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBundlesRequest) SetMaxResults(v int32) *DescribeBundlesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeBundlesRequest) SetNextToken(v string) *DescribeBundlesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeBundlesRequest) SetBundleType(v string) *DescribeBundlesRequest {
	s.BundleType = &v
	return s
}

func (s *DescribeBundlesRequest) SetDesktopTypeFamily(v string) *DescribeBundlesRequest {
	s.DesktopTypeFamily = &v
	return s
}

func (s *DescribeBundlesRequest) SetCpuCount(v int32) *DescribeBundlesRequest {
	s.CpuCount = &v
	return s
}

func (s *DescribeBundlesRequest) SetMemorySize(v int32) *DescribeBundlesRequest {
	s.MemorySize = &v
	return s
}

func (s *DescribeBundlesRequest) SetGpuCount(v float32) *DescribeBundlesRequest {
	s.GpuCount = &v
	return s
}

func (s *DescribeBundlesRequest) SetCheckStock(v bool) *DescribeBundlesRequest {
	s.CheckStock = &v
	return s
}

func (s *DescribeBundlesRequest) SetFromDesktopGroup(v bool) *DescribeBundlesRequest {
	s.FromDesktopGroup = &v
	return s
}

func (s *DescribeBundlesRequest) SetProtocolType(v string) *DescribeBundlesRequest {
	s.ProtocolType = &v
	return s
}

func (s *DescribeBundlesRequest) SetBundleId(v []*string) *DescribeBundlesRequest {
	s.BundleId = v
	return s
}

type DescribeBundlesResponseBody struct {
	NextToken *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Bundles   []*DescribeBundlesResponseBodyBundles `json:"Bundles,omitempty" xml:"Bundles,omitempty" type:"Repeated"`
}

func (s DescribeBundlesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBody) SetNextToken(v string) *DescribeBundlesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeBundlesResponseBody) SetRequestId(v string) *DescribeBundlesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBundlesResponseBody) SetBundles(v []*DescribeBundlesResponseBodyBundles) *DescribeBundlesResponseBody {
	s.Bundles = v
	return s
}

type DescribeBundlesResponseBodyBundles struct {
	Description          *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	BundleType           *string                                                 `json:"BundleType,omitempty" xml:"BundleType,omitempty"`
	OsType               *string                                                 `json:"OsType,omitempty" xml:"OsType,omitempty"`
	StockState           *string                                                 `json:"StockState,omitempty" xml:"StockState,omitempty"`
	DesktopType          *string                                                 `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	ProtocolType         *string                                                 `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	BundleId             *string                                                 `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	ImageId              *string                                                 `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	BundleName           *string                                                 `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
	Disks                []*DescribeBundlesResponseBodyBundlesDisks              `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	DesktopTypeAttribute *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute `json:"DesktopTypeAttribute,omitempty" xml:"DesktopTypeAttribute,omitempty" type:"Struct"`
}

func (s DescribeBundlesResponseBodyBundles) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponseBodyBundles) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBodyBundles) SetDescription(v string) *DescribeBundlesResponseBodyBundles {
	s.Description = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetBundleType(v string) *DescribeBundlesResponseBodyBundles {
	s.BundleType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetOsType(v string) *DescribeBundlesResponseBodyBundles {
	s.OsType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetStockState(v string) *DescribeBundlesResponseBodyBundles {
	s.StockState = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDesktopType(v string) *DescribeBundlesResponseBodyBundles {
	s.DesktopType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetProtocolType(v string) *DescribeBundlesResponseBodyBundles {
	s.ProtocolType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetBundleId(v string) *DescribeBundlesResponseBodyBundles {
	s.BundleId = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetImageId(v string) *DescribeBundlesResponseBodyBundles {
	s.ImageId = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetBundleName(v string) *DescribeBundlesResponseBodyBundles {
	s.BundleName = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDisks(v []*DescribeBundlesResponseBodyBundlesDisks) *DescribeBundlesResponseBodyBundles {
	s.Disks = v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDesktopTypeAttribute(v *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) *DescribeBundlesResponseBodyBundles {
	s.DesktopTypeAttribute = v
	return s
}

type DescribeBundlesResponseBodyBundlesDisks struct {
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	DiskSize *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
}

func (s DescribeBundlesResponseBodyBundlesDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponseBodyBundlesDisks) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBodyBundlesDisks) SetDiskType(v string) *DescribeBundlesResponseBodyBundlesDisks {
	s.DiskType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDisks) SetDiskSize(v int32) *DescribeBundlesResponseBodyBundlesDisks {
	s.DiskSize = &v
	return s
}

type DescribeBundlesResponseBodyBundlesDesktopTypeAttribute struct {
	CpuCount   *int32   `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	GpuCount   *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	GpuSpec    *string  `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	MemorySize *int32   `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
}

func (s DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) SetCpuCount(v int32) *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute {
	s.CpuCount = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) SetGpuCount(v float32) *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute {
	s.GpuCount = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) SetGpuSpec(v string) *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute {
	s.GpuSpec = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) SetMemorySize(v int32) *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute {
	s.MemorySize = &v
	return s
}

type DescribeBundlesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBundlesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBundlesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponse) SetHeaders(v map[string]*string) *DescribeBundlesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBundlesResponse) SetBody(v *DescribeBundlesResponseBody) *DescribeBundlesResponse {
	s.Body = v
	return s
}

type DescribeCensRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeCensRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensRequest) GoString() string {
	return s.String()
}

func (s *DescribeCensRequest) SetRegionId(v string) *DescribeCensRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCensRequest) SetPageSize(v int32) *DescribeCensRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCensRequest) SetPageNumber(v int32) *DescribeCensRequest {
	s.PageNumber = &v
	return s
}

type DescribeCensResponseBody struct {
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Cens       []*DescribeCensResponseBodyCens `json:"Cens,omitempty" xml:"Cens,omitempty" type:"Repeated"`
}

func (s DescribeCensResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBody) SetPageSize(v int32) *DescribeCensResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCensResponseBody) SetRequestId(v string) *DescribeCensResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCensResponseBody) SetPageNumber(v int32) *DescribeCensResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCensResponseBody) SetTotalCount(v int32) *DescribeCensResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCensResponseBody) SetCens(v []*DescribeCensResponseBodyCens) *DescribeCensResponseBody {
	s.Cens = v
	return s
}

type DescribeCensResponseBodyCens struct {
	Status          *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	CreationTime    *string                                   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Ipv6Level       *string                                   `json:"Ipv6Level,omitempty" xml:"Ipv6Level,omitempty"`
	Description     *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	CenId           *string                                   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ProtectionLevel *string                                   `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	Name            *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Tags            []*DescribeCensResponseBodyCensTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	PackageIds      []*DescribeCensResponseBodyCensPackageIds `json:"PackageIds,omitempty" xml:"PackageIds,omitempty" type:"Repeated"`
}

func (s DescribeCensResponseBodyCens) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBodyCens) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCens) SetStatus(v string) *DescribeCensResponseBodyCens {
	s.Status = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetCreationTime(v string) *DescribeCensResponseBodyCens {
	s.CreationTime = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetIpv6Level(v string) *DescribeCensResponseBodyCens {
	s.Ipv6Level = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetDescription(v string) *DescribeCensResponseBodyCens {
	s.Description = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetCenId(v string) *DescribeCensResponseBodyCens {
	s.CenId = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetProtectionLevel(v string) *DescribeCensResponseBodyCens {
	s.ProtectionLevel = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetName(v string) *DescribeCensResponseBodyCens {
	s.Name = &v
	return s
}

func (s *DescribeCensResponseBodyCens) SetTags(v []*DescribeCensResponseBodyCensTags) *DescribeCensResponseBodyCens {
	s.Tags = v
	return s
}

func (s *DescribeCensResponseBodyCens) SetPackageIds(v []*DescribeCensResponseBodyCensPackageIds) *DescribeCensResponseBodyCens {
	s.PackageIds = v
	return s
}

type DescribeCensResponseBodyCensTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCensResponseBodyCensTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBodyCensTags) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCensTags) SetKey(v string) *DescribeCensResponseBodyCensTags {
	s.Key = &v
	return s
}

func (s *DescribeCensResponseBodyCensTags) SetValue(v string) *DescribeCensResponseBodyCensTags {
	s.Value = &v
	return s
}

type DescribeCensResponseBodyCensPackageIds struct {
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
}

func (s DescribeCensResponseBodyCensPackageIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponseBodyCensPackageIds) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBodyCensPackageIds) SetPackageId(v string) *DescribeCensResponseBodyCensPackageIds {
	s.PackageId = &v
	return s
}

type DescribeCensResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCensResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCensResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCensResponse) GoString() string {
	return s.String()
}

func (s *DescribeCensResponse) SetHeaders(v map[string]*string) *DescribeCensResponse {
	s.Headers = v
	return s
}

func (s *DescribeCensResponse) SetBody(v *DescribeCensResponseBody) *DescribeCensResponse {
	s.Body = v
	return s
}

type DescribeClientEventsRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	EndUserId    *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	DesktopId    *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopIp    *string `json:"DesktopIp,omitempty" xml:"DesktopIp,omitempty"`
	DirectoryId  *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	EventType    *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeClientEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsRequest) SetRegionId(v string) *DescribeClientEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetEndUserId(v string) *DescribeClientEventsRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetDesktopId(v string) *DescribeClientEventsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetDesktopIp(v string) *DescribeClientEventsRequest {
	s.DesktopIp = &v
	return s
}

func (s *DescribeClientEventsRequest) SetDirectoryId(v string) *DescribeClientEventsRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetOfficeSiteId(v string) *DescribeClientEventsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeClientEventsRequest) SetEventType(v string) *DescribeClientEventsRequest {
	s.EventType = &v
	return s
}

func (s *DescribeClientEventsRequest) SetStartTime(v string) *DescribeClientEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeClientEventsRequest) SetEndTime(v string) *DescribeClientEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeClientEventsRequest) SetMaxResults(v int32) *DescribeClientEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeClientEventsRequest) SetNextToken(v string) *DescribeClientEventsRequest {
	s.NextToken = &v
	return s
}

type DescribeClientEventsResponseBody struct {
	NextToken *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Events    []*DescribeClientEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
}

func (s DescribeClientEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsResponseBody) SetNextToken(v string) *DescribeClientEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeClientEventsResponseBody) SetRequestId(v string) *DescribeClientEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClientEventsResponseBody) SetEvents(v []*DescribeClientEventsResponseBodyEvents) *DescribeClientEventsResponseBody {
	s.Events = v
	return s
}

type DescribeClientEventsResponseBodyEvents struct {
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	BytesReceived  *string `json:"BytesReceived,omitempty" xml:"BytesReceived,omitempty"`
	DesktopIp      *string `json:"DesktopIp,omitempty" xml:"DesktopIp,omitempty"`
	EventTime      *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	BytesSend      *string `json:"BytesSend,omitempty" xml:"BytesSend,omitempty"`
	OfficeSiteId   *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	AliUid         *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	EventId        *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	DirectoryType  *string `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	EventType      *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	EndUserId      *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ClientIp       *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientOS       *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	DirectoryId    *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	ClientVersion  *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
}

func (s DescribeClientEventsResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsResponseBodyEvents) SetStatus(v string) *DescribeClientEventsResponseBodyEvents {
	s.Status = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetBytesReceived(v string) *DescribeClientEventsResponseBodyEvents {
	s.BytesReceived = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDesktopIp(v string) *DescribeClientEventsResponseBodyEvents {
	s.DesktopIp = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetEventTime(v string) *DescribeClientEventsResponseBodyEvents {
	s.EventTime = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetBytesSend(v string) *DescribeClientEventsResponseBodyEvents {
	s.BytesSend = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetOfficeSiteId(v string) *DescribeClientEventsResponseBodyEvents {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetAliUid(v string) *DescribeClientEventsResponseBodyEvents {
	s.AliUid = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDesktopId(v string) *DescribeClientEventsResponseBodyEvents {
	s.DesktopId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetRegionId(v string) *DescribeClientEventsResponseBodyEvents {
	s.RegionId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetEventId(v string) *DescribeClientEventsResponseBodyEvents {
	s.EventId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDirectoryType(v string) *DescribeClientEventsResponseBodyEvents {
	s.DirectoryType = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetEventType(v string) *DescribeClientEventsResponseBodyEvents {
	s.EventType = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetEndUserId(v string) *DescribeClientEventsResponseBodyEvents {
	s.EndUserId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetClientIp(v string) *DescribeClientEventsResponseBodyEvents {
	s.ClientIp = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetClientOS(v string) *DescribeClientEventsResponseBodyEvents {
	s.ClientOS = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetOfficeSiteType(v string) *DescribeClientEventsResponseBodyEvents {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetDirectoryId(v string) *DescribeClientEventsResponseBodyEvents {
	s.DirectoryId = &v
	return s
}

func (s *DescribeClientEventsResponseBodyEvents) SetClientVersion(v string) *DescribeClientEventsResponseBodyEvents {
	s.ClientVersion = &v
	return s
}

type DescribeClientEventsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClientEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClientEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClientEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientEventsResponse) SetHeaders(v map[string]*string) *DescribeClientEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientEventsResponse) SetBody(v *DescribeClientEventsResponseBody) *DescribeClientEventsResponse {
	s.Body = v
	return s
}

type DescribeDesktopIdsByVulNamesRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type         *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	OfficeSiteId *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	Necessity    *string   `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	VulName      []*string `json:"VulName,omitempty" xml:"VulName,omitempty" type:"Repeated"`
}

func (s DescribeDesktopIdsByVulNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopIdsByVulNamesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetRegionId(v string) *DescribeDesktopIdsByVulNamesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetType(v string) *DescribeDesktopIdsByVulNamesRequest {
	s.Type = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetOfficeSiteId(v string) *DescribeDesktopIdsByVulNamesRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetNecessity(v string) *DescribeDesktopIdsByVulNamesRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetVulName(v []*string) *DescribeDesktopIdsByVulNamesRequest {
	s.VulName = v
	return s
}

type DescribeDesktopIdsByVulNamesResponseBody struct {
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DesktopItems []*DescribeDesktopIdsByVulNamesResponseBodyDesktopItems `json:"DesktopItems,omitempty" xml:"DesktopItems,omitempty" type:"Repeated"`
}

func (s DescribeDesktopIdsByVulNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopIdsByVulNamesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopIdsByVulNamesResponseBody) SetRequestId(v string) *DescribeDesktopIdsByVulNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesResponseBody) SetDesktopItems(v []*DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) *DescribeDesktopIdsByVulNamesResponseBody {
	s.DesktopItems = v
	return s
}

type DescribeDesktopIdsByVulNamesResponseBodyDesktopItems struct {
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
}

func (s DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) GoString() string {
	return s.String()
}

func (s *DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) SetDesktopName(v string) *DescribeDesktopIdsByVulNamesResponseBodyDesktopItems {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) SetDesktopId(v string) *DescribeDesktopIdsByVulNamesResponseBodyDesktopItems {
	s.DesktopId = &v
	return s
}

type DescribeDesktopIdsByVulNamesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDesktopIdsByVulNamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDesktopIdsByVulNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopIdsByVulNamesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopIdsByVulNamesResponse) SetHeaders(v map[string]*string) *DescribeDesktopIdsByVulNamesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopIdsByVulNamesResponse) SetBody(v *DescribeDesktopIdsByVulNamesResponseBody) *DescribeDesktopIdsByVulNamesResponse {
	s.Body = v
	return s
}

type DescribeDesktopPolicysRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	DesktopId  []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s DescribeDesktopPolicysRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopPolicysRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopPolicysRequest) SetRegionId(v string) *DescribeDesktopPolicysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopPolicysRequest) SetNextToken(v string) *DescribeDesktopPolicysRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopPolicysRequest) SetMaxResults(v int32) *DescribeDesktopPolicysRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopPolicysRequest) SetDesktopId(v []*string) *DescribeDesktopPolicysRequest {
	s.DesktopId = v
	return s
}

type DescribeDesktopPolicysResponseBody struct {
	NextToken              *string                                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId              *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DescribeDesktopPolicys []*DescribeDesktopPolicysResponseBodyDescribeDesktopPolicys `json:"DescribeDesktopPolicys,omitempty" xml:"DescribeDesktopPolicys,omitempty" type:"Repeated"`
}

func (s DescribeDesktopPolicysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopPolicysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopPolicysResponseBody) SetNextToken(v string) *DescribeDesktopPolicysResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopPolicysResponseBody) SetRequestId(v string) *DescribeDesktopPolicysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopPolicysResponseBody) SetDescribeDesktopPolicys(v []*DescribeDesktopPolicysResponseBodyDescribeDesktopPolicys) *DescribeDesktopPolicysResponseBody {
	s.DescribeDesktopPolicys = v
	return s
}

type DescribeDesktopPolicysResponseBodyDescribeDesktopPolicys struct {
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	Watermark   *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	Clipboard   *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	LocalDrive  *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
}

func (s DescribeDesktopPolicysResponseBodyDescribeDesktopPolicys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopPolicysResponseBodyDescribeDesktopPolicys) GoString() string {
	return s.String()
}

func (s *DescribeDesktopPolicysResponseBodyDescribeDesktopPolicys) SetUsbRedirect(v string) *DescribeDesktopPolicysResponseBodyDescribeDesktopPolicys {
	s.UsbRedirect = &v
	return s
}

func (s *DescribeDesktopPolicysResponseBodyDescribeDesktopPolicys) SetDesktopId(v string) *DescribeDesktopPolicysResponseBodyDescribeDesktopPolicys {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopPolicysResponseBodyDescribeDesktopPolicys) SetWatermark(v string) *DescribeDesktopPolicysResponseBodyDescribeDesktopPolicys {
	s.Watermark = &v
	return s
}

func (s *DescribeDesktopPolicysResponseBodyDescribeDesktopPolicys) SetClipboard(v string) *DescribeDesktopPolicysResponseBodyDescribeDesktopPolicys {
	s.Clipboard = &v
	return s
}

func (s *DescribeDesktopPolicysResponseBodyDescribeDesktopPolicys) SetLocalDrive(v string) *DescribeDesktopPolicysResponseBodyDescribeDesktopPolicys {
	s.LocalDrive = &v
	return s
}

type DescribeDesktopPolicysResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDesktopPolicysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDesktopPolicysResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopPolicysResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopPolicysResponse) SetHeaders(v map[string]*string) *DescribeDesktopPolicysResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopPolicysResponse) SetBody(v *DescribeDesktopPolicysResponseBody) *DescribeDesktopPolicysResponse {
	s.Body = v
	return s
}

type DescribeDesktopsRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	GroupId       *string   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	DesktopStatus *string   `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	MaxResults    *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	UserName      *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
	DesktopName   *string   `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DirectoryId   *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OfficeSiteId  *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	PolicyGroupId *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	ChargeType    *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ExpiredTime   *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	DesktopId     []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	EndUserId     []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
}

func (s DescribeDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsRequest) SetRegionId(v string) *DescribeDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetGroupId(v string) *DescribeDesktopsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopStatus(v string) *DescribeDesktopsRequest {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsRequest) SetMaxResults(v int32) *DescribeDesktopsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopsRequest) SetNextToken(v string) *DescribeDesktopsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsRequest) SetUserName(v string) *DescribeDesktopsRequest {
	s.UserName = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopName(v string) *DescribeDesktopsRequest {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDirectoryId(v string) *DescribeDesktopsRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetOfficeSiteId(v string) *DescribeDesktopsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetPolicyGroupId(v string) *DescribeDesktopsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopsRequest) SetChargeType(v string) *DescribeDesktopsRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeDesktopsRequest) SetExpiredTime(v string) *DescribeDesktopsRequest {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDesktopId(v []*string) *DescribeDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *DescribeDesktopsRequest) SetEndUserId(v []*string) *DescribeDesktopsRequest {
	s.EndUserId = v
	return s
}

type DescribeDesktopsResponseBody struct {
	NextToken *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Desktops  []*DescribeDesktopsResponseBodyDesktops `json:"Desktops,omitempty" xml:"Desktops,omitempty" type:"Repeated"`
}

func (s DescribeDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBody) SetNextToken(v string) *DescribeDesktopsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsResponseBody) SetRequestId(v string) *DescribeDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopsResponseBody) SetDesktops(v []*DescribeDesktopsResponseBodyDesktops) *DescribeDesktopsResponseBody {
	s.Desktops = v
	return s
}

type DescribeDesktopsResponseBodyDesktops struct {
	CreationTime       *string                                         `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ChargeType         *string                                         `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DesktopName        *string                                         `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	PolicyGroupName    *string                                         `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	SystemDiskSize     *int32                                          `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	PolicyGroupId      *string                                         `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	DesktopStatus      *string                                         `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	DesktopType        *string                                         `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	GpuCount           *float32                                        `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	Memory             *int64                                          `json:"Memory,omitempty" xml:"Memory,omitempty"`
	GpuSpec            *string                                         `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	DirectoryId        *string                                         `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	ImageId            *string                                         `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ManagementFlag     *string                                         `json:"ManagementFlag,omitempty" xml:"ManagementFlag,omitempty"`
	DataDiskCategory   *string                                         `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	SystemDiskCategory *string                                         `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	OfficeSiteId       *string                                         `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	NetworkInterfaceId *string                                         `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	DataDiskSize       *string                                         `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DesktopId          *string                                         `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	OfficeSiteName     *string                                         `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	StartTime          *string                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DirectoryType      *string                                         `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	Cpu                *int32                                          `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	NetworkInterfaceIp *string                                         `json:"NetworkInterfaceIp,omitempty" xml:"NetworkInterfaceIp,omitempty"`
	ExpiredTime        *string                                         `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	OsType             *string                                         `json:"OsType,omitempty" xml:"OsType,omitempty"`
	ConnectionStatus   *string                                         `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	BundleId           *string                                         `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	OfficeSiteType     *string                                         `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	Disks              []*DescribeDesktopsResponseBodyDesktopsDisks    `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	Tags               []*DescribeDesktopsResponseBodyDesktopsTags     `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Sessions           []*DescribeDesktopsResponseBodyDesktopsSessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	EndUserIds         []*string                                       `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
}

func (s DescribeDesktopsResponseBodyDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktops) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktops) SetCreationTime(v string) *DescribeDesktopsResponseBodyDesktops {
	s.CreationTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetChargeType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ChargeType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopName(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetPolicyGroupName(v string) *DescribeDesktopsResponseBodyDesktops {
	s.PolicyGroupName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetSystemDiskSize(v int32) *DescribeDesktopsResponseBodyDesktops {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetPolicyGroupId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopStatus(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetGpuCount(v float32) *DescribeDesktopsResponseBodyDesktops {
	s.GpuCount = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetMemory(v int64) *DescribeDesktopsResponseBodyDesktops {
	s.Memory = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetGpuSpec(v string) *DescribeDesktopsResponseBodyDesktops {
	s.GpuSpec = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDirectoryId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetImageId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetManagementFlag(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ManagementFlag = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDataDiskCategory(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetSystemDiskCategory(v string) *DescribeDesktopsResponseBodyDesktops {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetOfficeSiteId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetNetworkInterfaceId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDataDiskSize(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetOfficeSiteName(v string) *DescribeDesktopsResponseBodyDesktops {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetStartTime(v string) *DescribeDesktopsResponseBodyDesktops {
	s.StartTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDirectoryType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetCpu(v int32) *DescribeDesktopsResponseBodyDesktops {
	s.Cpu = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetNetworkInterfaceIp(v string) *DescribeDesktopsResponseBodyDesktops {
	s.NetworkInterfaceIp = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetExpiredTime(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetOsType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.OsType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetConnectionStatus(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetBundleId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.BundleId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetOfficeSiteType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDisks(v []*DescribeDesktopsResponseBodyDesktopsDisks) *DescribeDesktopsResponseBodyDesktops {
	s.Disks = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetTags(v []*DescribeDesktopsResponseBodyDesktopsTags) *DescribeDesktopsResponseBodyDesktops {
	s.Tags = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetSessions(v []*DescribeDesktopsResponseBodyDesktopsSessions) *DescribeDesktopsResponseBodyDesktops {
	s.Sessions = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetEndUserIds(v []*string) *DescribeDesktopsResponseBodyDesktops {
	s.EndUserIds = v
	return s
}

type DescribeDesktopsResponseBodyDesktopsDisks struct {
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	DiskId   *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskSize *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktopsDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsDisks) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetDiskType(v string) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.DiskType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetDiskId(v string) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetDiskSize(v int32) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.DiskSize = &v
	return s
}

type DescribeDesktopsResponseBodyDesktopsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktopsTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsTags) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsTags) SetKey(v string) *DescribeDesktopsResponseBodyDesktopsTags {
	s.Key = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsTags) SetValue(v string) *DescribeDesktopsResponseBodyDesktopsTags {
	s.Value = &v
	return s
}

type DescribeDesktopsResponseBodyDesktopsSessions struct {
	EndUserId         *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EstablishmentTime *string `json:"EstablishmentTime,omitempty" xml:"EstablishmentTime,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktopsSessions) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsSessions) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsSessions) SetEndUserId(v string) *DescribeDesktopsResponseBodyDesktopsSessions {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsSessions) SetEstablishmentTime(v string) *DescribeDesktopsResponseBodyDesktopsSessions {
	s.EstablishmentTime = &v
	return s
}

type DescribeDesktopsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponse) SetHeaders(v map[string]*string) *DescribeDesktopsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopsResponse) SetBody(v *DescribeDesktopsResponseBody) *DescribeDesktopsResponse {
	s.Body = v
	return s
}

type DescribeDesktopsInGroupRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	PayType        *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeDesktopsInGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsInGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupRequest) SetRegionId(v string) *DescribeDesktopsInGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetDesktopGroupId(v string) *DescribeDesktopsInGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetPayType(v string) *DescribeDesktopsInGroupRequest {
	s.PayType = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetMaxResults(v int32) *DescribeDesktopsInGroupRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopsInGroupRequest) SetNextToken(v string) *DescribeDesktopsInGroupRequest {
	s.NextToken = &v
	return s
}

type DescribeDesktopsInGroupResponseBody struct {
	PostPaidDesktopsCount       *int32                                                 `json:"PostPaidDesktopsCount,omitempty" xml:"PostPaidDesktopsCount,omitempty"`
	NextToken                   *string                                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PaidDesktopsCount           *int32                                                 `json:"PaidDesktopsCount,omitempty" xml:"PaidDesktopsCount,omitempty"`
	RequestId                   *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PostPaidDesktopsTotalAmount *int32                                                 `json:"PostPaidDesktopsTotalAmount,omitempty" xml:"PostPaidDesktopsTotalAmount,omitempty"`
	PaidDesktops                []*DescribeDesktopsInGroupResponseBodyPaidDesktops     `json:"PaidDesktops,omitempty" xml:"PaidDesktops,omitempty" type:"Repeated"`
	PostPaidDesktops            []*DescribeDesktopsInGroupResponseBodyPostPaidDesktops `json:"PostPaidDesktops,omitempty" xml:"PostPaidDesktops,omitempty" type:"Repeated"`
}

func (s DescribeDesktopsInGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupResponseBody) SetPostPaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody {
	s.PostPaidDesktopsCount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetNextToken(v string) *DescribeDesktopsInGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetPaidDesktopsCount(v int32) *DescribeDesktopsInGroupResponseBody {
	s.PaidDesktopsCount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetRequestId(v string) *DescribeDesktopsInGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetPostPaidDesktopsTotalAmount(v int32) *DescribeDesktopsInGroupResponseBody {
	s.PostPaidDesktopsTotalAmount = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetPaidDesktops(v []*DescribeDesktopsInGroupResponseBodyPaidDesktops) *DescribeDesktopsInGroupResponseBody {
	s.PaidDesktops = v
	return s
}

func (s *DescribeDesktopsInGroupResponseBody) SetPostPaidDesktops(v []*DescribeDesktopsInGroupResponseBodyPostPaidDesktops) *DescribeDesktopsInGroupResponseBody {
	s.PostPaidDesktops = v
	return s
}

type DescribeDesktopsInGroupResponseBodyPaidDesktops struct {
	EndUserId        *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	DesktopStatus    *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	DesktopName      *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	DesktopId        *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	EndUserName      *string `json:"EndUserName,omitempty" xml:"EndUserName,omitempty"`
}

func (s DescribeDesktopsInGroupResponseBodyPaidDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsInGroupResponseBodyPaidDesktops) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetEndUserId(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetDesktopStatus(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetDesktopName(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetConnectionStatus(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetDesktopId(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPaidDesktops) SetEndUserName(v string) *DescribeDesktopsInGroupResponseBodyPaidDesktops {
	s.EndUserName = &v
	return s
}

type DescribeDesktopsInGroupResponseBodyPostPaidDesktops struct {
	CreateDuration   *string `json:"CreateDuration,omitempty" xml:"CreateDuration,omitempty"`
	EndUserId        *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	DesktopStatus    *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ReleaseTime      *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	DesktopName      *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	DesktopId        *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	EndUserName      *string `json:"EndUserName,omitempty" xml:"EndUserName,omitempty"`
}

func (s DescribeDesktopsInGroupResponseBodyPostPaidDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsInGroupResponseBodyPostPaidDesktops) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetCreateDuration(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.CreateDuration = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetEndUserId(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetDesktopStatus(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetCreateTime(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.CreateTime = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetReleaseTime(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetDesktopName(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetConnectionStatus(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetDesktopId(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopsInGroupResponseBodyPostPaidDesktops) SetEndUserName(v string) *DescribeDesktopsInGroupResponseBodyPostPaidDesktops {
	s.EndUserName = &v
	return s
}

type DescribeDesktopsInGroupResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDesktopsInGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDesktopsInGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsInGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupResponse) SetHeaders(v map[string]*string) *DescribeDesktopsInGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopsInGroupResponse) SetBody(v *DescribeDesktopsInGroupResponseBody) *DescribeDesktopsInGroupResponse {
	s.Body = v
	return s
}

type DescribeDesktopTypesRequest struct {
	RegionId           *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopTypeId      *string  `json:"DesktopTypeId,omitempty" xml:"DesktopTypeId,omitempty"`
	InstanceTypeFamily *string  `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	CpuCount           *int32   `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	MemorySize         *int32   `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	GpuCount           *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
}

func (s DescribeDesktopTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesRequest) SetRegionId(v string) *DescribeDesktopTypesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetDesktopTypeId(v string) *DescribeDesktopTypesRequest {
	s.DesktopTypeId = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetInstanceTypeFamily(v string) *DescribeDesktopTypesRequest {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetCpuCount(v int32) *DescribeDesktopTypesRequest {
	s.CpuCount = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetMemorySize(v int32) *DescribeDesktopTypesRequest {
	s.MemorySize = &v
	return s
}

func (s *DescribeDesktopTypesRequest) SetGpuCount(v float32) *DescribeDesktopTypesRequest {
	s.GpuCount = &v
	return s
}

type DescribeDesktopTypesResponseBody struct {
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DesktopTypes []*DescribeDesktopTypesResponseBodyDesktopTypes `json:"DesktopTypes,omitempty" xml:"DesktopTypes,omitempty" type:"Repeated"`
}

func (s DescribeDesktopTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponseBody) SetRequestId(v string) *DescribeDesktopTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopTypesResponseBody) SetDesktopTypes(v []*DescribeDesktopTypesResponseBodyDesktopTypes) *DescribeDesktopTypesResponseBody {
	s.DesktopTypes = v
	return s
}

type DescribeDesktopTypesResponseBodyDesktopTypes struct {
	SystemDiskSize     *string                                                      `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	DesktopTypeId      *string                                                      `json:"DesktopTypeId,omitempty" xml:"DesktopTypeId,omitempty"`
	DataDiskSize       *string                                                      `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	CpuCount           *string                                                      `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	GpuCount           *float32                                                     `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	GpuSpec            *string                                                      `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	InstanceTypeFamily *string                                                      `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	MemorySize         *string                                                      `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	AllowDiskSize      []*DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize `json:"AllowDiskSize,omitempty" xml:"AllowDiskSize,omitempty" type:"Repeated"`
}

func (s DescribeDesktopTypesResponseBodyDesktopTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesResponseBodyDesktopTypes) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetSystemDiskSize(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetDesktopTypeId(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.DesktopTypeId = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetDataDiskSize(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetCpuCount(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.CpuCount = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetGpuCount(v float32) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.GpuCount = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetGpuSpec(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.GpuSpec = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetInstanceTypeFamily(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetMemorySize(v string) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.MemorySize = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetAllowDiskSize(v []*DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize) *DescribeDesktopTypesResponseBodyDesktopTypes {
	s.AllowDiskSize = v
	return s
}

type DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize struct {
	DataDiskSize   *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize) SetDataDiskSize(v int32) *DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize) SetSystemDiskSize(v int32) *DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize {
	s.SystemDiskSize = &v
	return s
}

type DescribeDesktopTypesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDesktopTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDesktopTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponse) SetHeaders(v map[string]*string) *DescribeDesktopTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopTypesResponse) SetBody(v *DescribeDesktopTypesResponseBody) *DescribeDesktopTypesResponse {
	s.Body = v
	return s
}

type DescribeDirectoriesRequest struct {
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DirectoryType   *string   `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	DirectoryStatus *string   `json:"DirectoryStatus,omitempty" xml:"DirectoryStatus,omitempty"`
	MaxResults      *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	DirectoryId     []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" type:"Repeated"`
}

func (s DescribeDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesRequest) SetRegionId(v string) *DescribeDirectoriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryType(v string) *DescribeDirectoriesRequest {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryStatus(v string) *DescribeDirectoriesRequest {
	s.DirectoryStatus = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetMaxResults(v int32) *DescribeDirectoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetNextToken(v string) *DescribeDirectoriesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryId(v []*string) *DescribeDirectoriesRequest {
	s.DirectoryId = v
	return s
}

type DescribeDirectoriesResponseBody struct {
	NextToken   *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Directories []*DescribeDirectoriesResponseBodyDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
}

func (s DescribeDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBody) SetNextToken(v string) *DescribeDirectoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetRequestId(v string) *DescribeDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetDirectories(v []*DescribeDirectoriesResponseBodyDirectories) *DescribeDirectoriesResponseBody {
	s.Directories = v
	return s
}

type DescribeDirectoriesResponseBodyDirectories struct {
	EnableInternetAccess     *bool                                                     `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	VpcId                    *string                                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	CreationTime             *string                                                   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Status                   *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	DomainPassword           *string                                                   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	EnableAdminAccess        *bool                                                     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	SubDomainName            *string                                                   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	DomainUserName           *string                                                   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	EnableCrossDesktopAccess *bool                                                     `json:"EnableCrossDesktopAccess,omitempty" xml:"EnableCrossDesktopAccess,omitempty"`
	CustomSecurityGroupId    *string                                                   `json:"CustomSecurityGroupId,omitempty" xml:"CustomSecurityGroupId,omitempty"`
	DesktopVpcEndpoint       *string                                                   `json:"DesktopVpcEndpoint,omitempty" xml:"DesktopVpcEndpoint,omitempty"`
	SsoEnabled               *bool                                                     `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty"`
	DomainName               *string                                                   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DesktopAccessType        *string                                                   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	MfaEnabled               *bool                                                     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	DirectoryType            *string                                                   `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	DnsUserName              *string                                                   `json:"DnsUserName,omitempty" xml:"DnsUserName,omitempty"`
	TrustPassword            *string                                                   `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty"`
	Name                     *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	DirectoryId              *string                                                   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	ADConnectors             []*DescribeDirectoriesResponseBodyDirectoriesADConnectors `json:"ADConnectors,omitempty" xml:"ADConnectors,omitempty" type:"Repeated"`
	Logs                     []*DescribeDirectoriesResponseBodyDirectoriesLogs         `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	VSwitchIds               []*string                                                 `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	FileSystemIds            []*string                                                 `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty" type:"Repeated"`
	SubDnsAddress            []*string                                                 `json:"SubDnsAddress,omitempty" xml:"SubDnsAddress,omitempty" type:"Repeated"`
	DnsAddress               []*string                                                 `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
}

func (s DescribeDirectoriesResponseBodyDirectories) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBodyDirectories) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetEnableInternetAccess(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.EnableInternetAccess = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetVpcId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.VpcId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetCreationTime(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.CreationTime = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetStatus(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.Status = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDomainPassword(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DomainPassword = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetEnableAdminAccess(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.EnableAdminAccess = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetSubDomainName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.SubDomainName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDomainUserName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DomainUserName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetEnableCrossDesktopAccess(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.EnableCrossDesktopAccess = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetCustomSecurityGroupId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.CustomSecurityGroupId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDesktopVpcEndpoint(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DesktopVpcEndpoint = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetSsoEnabled(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.SsoEnabled = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDomainName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DomainName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDesktopAccessType(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetMfaEnabled(v bool) *DescribeDirectoriesResponseBodyDirectories {
	s.MfaEnabled = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDirectoryType(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDnsUserName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DnsUserName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetTrustPassword(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.TrustPassword = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.Name = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDirectoryId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetADConnectors(v []*DescribeDirectoriesResponseBodyDirectoriesADConnectors) *DescribeDirectoriesResponseBodyDirectories {
	s.ADConnectors = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetLogs(v []*DescribeDirectoriesResponseBodyDirectoriesLogs) *DescribeDirectoriesResponseBodyDirectories {
	s.Logs = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetVSwitchIds(v []*string) *DescribeDirectoriesResponseBodyDirectories {
	s.VSwitchIds = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetFileSystemIds(v []*string) *DescribeDirectoriesResponseBodyDirectories {
	s.FileSystemIds = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetSubDnsAddress(v []*string) *DescribeDirectoriesResponseBodyDirectories {
	s.SubDnsAddress = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDnsAddress(v []*string) *DescribeDirectoriesResponseBodyDirectories {
	s.DnsAddress = v
	return s
}

type DescribeDirectoriesResponseBodyDirectoriesADConnectors struct {
	ConnectorStatus    *string `json:"ConnectorStatus,omitempty" xml:"ConnectorStatus,omitempty"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ADConnectorAddress *string `json:"ADConnectorAddress,omitempty" xml:"ADConnectorAddress,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
}

func (s DescribeDirectoriesResponseBodyDirectoriesADConnectors) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBodyDirectoriesADConnectors) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetConnectorStatus(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.ConnectorStatus = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetVSwitchId(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetADConnectorAddress(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.ADConnectorAddress = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetNetworkInterfaceId(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.NetworkInterfaceId = &v
	return s
}

type DescribeDirectoriesResponseBodyDirectoriesLogs struct {
	Step      *string `json:"Step,omitempty" xml:"Step,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Level     *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeDirectoriesResponseBodyDirectoriesLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBodyDirectoriesLogs) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) SetStep(v string) *DescribeDirectoriesResponseBodyDirectoriesLogs {
	s.Step = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) SetMessage(v string) *DescribeDirectoriesResponseBodyDirectoriesLogs {
	s.Message = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) SetTimeStamp(v string) *DescribeDirectoriesResponseBodyDirectoriesLogs {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesLogs) SetLevel(v string) *DescribeDirectoriesResponseBodyDirectoriesLogs {
	s.Level = &v
	return s
}

type DescribeDirectoriesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponse) SetHeaders(v map[string]*string) *DescribeDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDirectoriesResponse) SetBody(v *DescribeDirectoriesResponseBody) *DescribeDirectoriesResponse {
	s.Body = v
	return s
}

type DescribeFrontVulPatchListRequest struct {
	RegionId    *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OperateType *string                                    `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	Type        *string                                    `json:"Type,omitempty" xml:"Type,omitempty"`
	VulInfo     []*DescribeFrontVulPatchListRequestVulInfo `json:"VulInfo,omitempty" xml:"VulInfo,omitempty" type:"Repeated"`
}

func (s DescribeFrontVulPatchListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListRequest) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListRequest) SetRegionId(v string) *DescribeFrontVulPatchListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFrontVulPatchListRequest) SetOperateType(v string) *DescribeFrontVulPatchListRequest {
	s.OperateType = &v
	return s
}

func (s *DescribeFrontVulPatchListRequest) SetType(v string) *DescribeFrontVulPatchListRequest {
	s.Type = &v
	return s
}

func (s *DescribeFrontVulPatchListRequest) SetVulInfo(v []*DescribeFrontVulPatchListRequestVulInfo) *DescribeFrontVulPatchListRequest {
	s.VulInfo = v
	return s
}

type DescribeFrontVulPatchListRequestVulInfo struct {
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	Tag       *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFrontVulPatchListRequestVulInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListRequestVulInfo) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListRequestVulInfo) SetDesktopId(v string) *DescribeFrontVulPatchListRequestVulInfo {
	s.DesktopId = &v
	return s
}

func (s *DescribeFrontVulPatchListRequestVulInfo) SetTag(v string) *DescribeFrontVulPatchListRequestVulInfo {
	s.Tag = &v
	return s
}

func (s *DescribeFrontVulPatchListRequestVulInfo) SetName(v string) *DescribeFrontVulPatchListRequestVulInfo {
	s.Name = &v
	return s
}

type DescribeFrontVulPatchListResponseBody struct {
	RequestId      *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FrontPatchList []*DescribeFrontVulPatchListResponseBodyFrontPatchList `json:"FrontPatchList,omitempty" xml:"FrontPatchList,omitempty" type:"Repeated"`
}

func (s DescribeFrontVulPatchListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListResponseBody) SetRequestId(v string) *DescribeFrontVulPatchListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFrontVulPatchListResponseBody) SetFrontPatchList(v []*DescribeFrontVulPatchListResponseBodyFrontPatchList) *DescribeFrontVulPatchListResponseBody {
	s.FrontPatchList = v
	return s
}

type DescribeFrontVulPatchListResponseBodyFrontPatchList struct {
	DesktopId *string                                                         `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	PatchList []*DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList `json:"PatchList,omitempty" xml:"PatchList,omitempty" type:"Repeated"`
}

func (s DescribeFrontVulPatchListResponseBodyFrontPatchList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListResponseBodyFrontPatchList) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchList) SetDesktopId(v string) *DescribeFrontVulPatchListResponseBodyFrontPatchList {
	s.DesktopId = &v
	return s
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchList) SetPatchList(v []*DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) *DescribeFrontVulPatchListResponseBodyFrontPatchList {
	s.PatchList = v
	return s
}

type DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
}

func (s DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) SetName(v string) *DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList {
	s.Name = &v
	return s
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) SetAliasName(v string) *DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList {
	s.AliasName = &v
	return s
}

type DescribeFrontVulPatchListResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFrontVulPatchListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFrontVulPatchListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFrontVulPatchListResponse) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListResponse) SetHeaders(v map[string]*string) *DescribeFrontVulPatchListResponse {
	s.Headers = v
	return s
}

func (s *DescribeFrontVulPatchListResponse) SetBody(v *DescribeFrontVulPatchListResponseBody) *DescribeFrontVulPatchListResponse {
	s.Body = v
	return s
}

type DescribeGroupedVulRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	Necessity    *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	Dealed       *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGroupedVulRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedVulRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulRequest) SetRegionId(v string) *DescribeGroupedVulRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetLang(v string) *DescribeGroupedVulRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetType(v string) *DescribeGroupedVulRequest {
	s.Type = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetOfficeSiteId(v string) *DescribeGroupedVulRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetNecessity(v string) *DescribeGroupedVulRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetDealed(v string) *DescribeGroupedVulRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetCurrentPage(v int32) *DescribeGroupedVulRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetPageSize(v int32) *DescribeGroupedVulRequest {
	s.PageSize = &v
	return s
}

type DescribeGroupedVulResponseBody struct {
	CurrentPage     *int32                                           `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId       *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize        *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount      *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	GroupedVulItems []*DescribeGroupedVulResponseBodyGroupedVulItems `json:"GroupedVulItems,omitempty" xml:"GroupedVulItems,omitempty" type:"Repeated"`
}

func (s DescribeGroupedVulResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedVulResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulResponseBody) SetCurrentPage(v int32) *DescribeGroupedVulResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedVulResponseBody) SetRequestId(v string) *DescribeGroupedVulResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupedVulResponseBody) SetPageSize(v int32) *DescribeGroupedVulResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedVulResponseBody) SetTotalCount(v int32) *DescribeGroupedVulResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBody) SetGroupedVulItems(v []*DescribeGroupedVulResponseBodyGroupedVulItems) *DescribeGroupedVulResponseBody {
	s.GroupedVulItems = v
	return s
}

type DescribeGroupedVulResponseBodyGroupedVulItems struct {
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	NntfCount    *int32  `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
	HandledCount *int32  `json:"HandledCount,omitempty" xml:"HandledCount,omitempty"`
	GmtLast      *string `json:"GmtLast,omitempty" xml:"GmtLast,omitempty"`
	Tags         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	LaterCount   *int32  `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	AliasName    *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AsapCount    *int32  `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
}

func (s DescribeGroupedVulResponseBodyGroupedVulItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedVulResponseBodyGroupedVulItems) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetType(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.Type = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetNntfCount(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.NntfCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetHandledCount(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.HandledCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetGmtLast(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.GmtLast = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetTags(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.Tags = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetLaterCount(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.LaterCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetAliasName(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.AliasName = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetName(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.Name = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetAsapCount(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.AsapCount = &v
	return s
}

type DescribeGroupedVulResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGroupedVulResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGroupedVulResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedVulResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulResponse) SetHeaders(v map[string]*string) *DescribeGroupedVulResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupedVulResponse) SetBody(v *DescribeGroupedVulResponseBody) *DescribeGroupedVulResponse {
	s.Body = v
	return s
}

type DescribeImagesRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults   *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ImageType    *string   `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageStatus  *string   `json:"ImageStatus,omitempty" xml:"ImageStatus,omitempty"`
	GpuCategory  *bool     `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
	ProtocolType *string   `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	ImageId      []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
}

func (s DescribeImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagesRequest) SetRegionId(v string) *DescribeImagesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImagesRequest) SetMaxResults(v int32) *DescribeImagesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeImagesRequest) SetNextToken(v string) *DescribeImagesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeImagesRequest) SetImageType(v string) *DescribeImagesRequest {
	s.ImageType = &v
	return s
}

func (s *DescribeImagesRequest) SetImageStatus(v string) *DescribeImagesRequest {
	s.ImageStatus = &v
	return s
}

func (s *DescribeImagesRequest) SetGpuCategory(v bool) *DescribeImagesRequest {
	s.GpuCategory = &v
	return s
}

func (s *DescribeImagesRequest) SetProtocolType(v string) *DescribeImagesRequest {
	s.ProtocolType = &v
	return s
}

func (s *DescribeImagesRequest) SetImageId(v []*string) *DescribeImagesRequest {
	s.ImageId = v
	return s
}

type DescribeImagesResponseBody struct {
	NextToken *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Images    []*DescribeImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
}

func (s DescribeImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBody) SetNextToken(v string) *DescribeImagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeImagesResponseBody) SetRequestId(v string) *DescribeImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImagesResponseBody) SetImages(v []*DescribeImagesResponseBodyImages) *DescribeImagesResponseBody {
	s.Images = v
	return s
}

type DescribeImagesResponseBodyImages struct {
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	DataDiskSize *int32  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	ImageType    *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Size         *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	OsType       *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	GpuCategory  *bool   `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
}

func (s DescribeImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImages) SetCreationTime(v string) *DescribeImagesResponseBodyImages {
	s.CreationTime = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetStatus(v string) *DescribeImagesResponseBodyImages {
	s.Status = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetProgress(v string) *DescribeImagesResponseBodyImages {
	s.Progress = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetDataDiskSize(v int32) *DescribeImagesResponseBodyImages {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetImageType(v string) *DescribeImagesResponseBodyImages {
	s.ImageType = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetDescription(v string) *DescribeImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetSize(v int32) *DescribeImagesResponseBodyImages {
	s.Size = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetOsType(v string) *DescribeImagesResponseBodyImages {
	s.OsType = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetProtocolType(v string) *DescribeImagesResponseBodyImages {
	s.ProtocolType = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetName(v string) *DescribeImagesResponseBodyImages {
	s.Name = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetImageId(v string) *DescribeImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetGpuCategory(v bool) *DescribeImagesResponseBodyImages {
	s.GpuCategory = &v
	return s
}

type DescribeImagesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponse) SetHeaders(v map[string]*string) *DescribeImagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeImagesResponse) SetBody(v *DescribeImagesResponseBody) *DescribeImagesResponse {
	s.Body = v
	return s
}

type DescribeInvocationsRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InvokeId        *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	CommandType     *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	InvokeStatus    *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	DesktopId       *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	IncludeOutput   *bool   `json:"IncludeOutput,omitempty" xml:"IncludeOutput,omitempty"`
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeInvocationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequest) SetRegionId(v string) *DescribeInvocationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeId(v string) *DescribeInvocationsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetCommandType(v string) *DescribeInvocationsRequest {
	s.CommandType = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeStatus(v string) *DescribeInvocationsRequest {
	s.InvokeStatus = &v
	return s
}

func (s *DescribeInvocationsRequest) SetDesktopId(v string) *DescribeInvocationsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetIncludeOutput(v bool) *DescribeInvocationsRequest {
	s.IncludeOutput = &v
	return s
}

func (s *DescribeInvocationsRequest) SetContentEncoding(v string) *DescribeInvocationsRequest {
	s.ContentEncoding = &v
	return s
}

func (s *DescribeInvocationsRequest) SetMaxResults(v int32) *DescribeInvocationsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInvocationsRequest) SetNextToken(v string) *DescribeInvocationsRequest {
	s.NextToken = &v
	return s
}

type DescribeInvocationsResponseBody struct {
	NextToken   *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Invocations []*DescribeInvocationsResponseBodyInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" type:"Repeated"`
}

func (s DescribeInvocationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBody) SetNextToken(v string) *DescribeInvocationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetRequestId(v string) *DescribeInvocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetInvocations(v []*DescribeInvocationsResponseBodyInvocations) *DescribeInvocationsResponseBody {
	s.Invocations = v
	return s
}

type DescribeInvocationsResponseBodyInvocations struct {
	CreationTime     *string                                                     `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	InvocationStatus *string                                                     `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	InvokeId         *string                                                     `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	CommandType      *string                                                     `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	CommandContent   *string                                                     `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	InvokeDesktops   []*DescribeInvocationsResponseBodyInvocationsInvokeDesktops `json:"InvokeDesktops,omitempty" xml:"InvokeDesktops,omitempty" type:"Repeated"`
}

func (s DescribeInvocationsResponseBodyInvocations) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocations) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCreationTime(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocations {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvokeId(v string) *DescribeInvocationsResponseBodyInvocations {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCommandType(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CommandType = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCommandContent(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CommandContent = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvokeDesktops(v []*DescribeInvocationsResponseBodyInvocationsInvokeDesktops) *DescribeInvocationsResponseBodyInvocations {
	s.InvokeDesktops = v
	return s
}

type DescribeInvocationsResponseBodyInvocationsInvokeDesktops struct {
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	FinishTime       *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	UpdateTime       *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Repeats          *int32  `json:"Repeats,omitempty" xml:"Repeats,omitempty"`
	DesktopId        *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	Output           *string `json:"Output,omitempty" xml:"Output,omitempty"`
	Dropped          *int32  `json:"Dropped,omitempty" xml:"Dropped,omitempty"`
	StopTime         *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	ExitCode         *int64  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ErrorInfo        *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	ErrorCode        *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvokeDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetCreationTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetFinishTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.FinishTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetUpdateTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.UpdateTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetRepeats(v int32) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.Repeats = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetDesktopId(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetOutput(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.Output = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetDropped(v int32) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.Dropped = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetStopTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.StopTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetExitCode(v int64) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.ExitCode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetStartTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.StartTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetErrorInfo(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetErrorCode(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.ErrorCode = &v
	return s
}

type DescribeInvocationsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInvocationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInvocationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponse) SetHeaders(v map[string]*string) *DescribeInvocationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvocationsResponse) SetBody(v *DescribeInvocationsResponseBody) *DescribeInvocationsResponse {
	s.Body = v
	return s
}

type DescribeModificationPriceRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RootDiskSizeGib *int32  `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	UserDiskSizeGib *int32  `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty"`
}

func (s DescribeModificationPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceRequest) SetRegionId(v string) *DescribeModificationPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetInstanceId(v string) *DescribeModificationPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetInstanceType(v string) *DescribeModificationPriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetRootDiskSizeGib(v int32) *DescribeModificationPriceRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *DescribeModificationPriceRequest) SetUserDiskSizeGib(v int32) *DescribeModificationPriceRequest {
	s.UserDiskSizeGib = &v
	return s
}

type DescribeModificationPriceResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PriceInfo *DescribeModificationPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
}

func (s DescribeModificationPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponseBody) SetRequestId(v string) *DescribeModificationPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModificationPriceResponseBody) SetPriceInfo(v *DescribeModificationPriceResponseBodyPriceInfo) *DescribeModificationPriceResponseBody {
	s.PriceInfo = v
	return s
}

type DescribeModificationPriceResponseBodyPriceInfo struct {
	Rules []*DescribeModificationPriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	Price *DescribeModificationPriceResponseBodyPriceInfoPrice   `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
}

func (s DescribeModificationPriceResponseBodyPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponseBodyPriceInfo) SetRules(v []*DescribeModificationPriceResponseBodyPriceInfoRules) *DescribeModificationPriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfo) SetPrice(v *DescribeModificationPriceResponseBodyPriceInfoPrice) *DescribeModificationPriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

type DescribeModificationPriceResponseBodyPriceInfoRules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeModificationPriceResponseBodyPriceInfoRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponseBodyPriceInfoRules) SetDescription(v string) *DescribeModificationPriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoRules) SetRuleId(v int64) *DescribeModificationPriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

type DescribeModificationPriceResponseBodyPriceInfoPrice struct {
	OriginalPrice *float32                                                         `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	DiscountPrice *float32                                                         `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	Currency      *string                                                          `json:"Currency,omitempty" xml:"Currency,omitempty"`
	TradePrice    *float32                                                         `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
	Promotions    []*DescribeModificationPriceResponseBodyPriceInfoPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
}

func (s DescribeModificationPriceResponseBodyPriceInfoPrice) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribeModificationPriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribeModificationPriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribeModificationPriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribeModificationPriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) SetPromotions(v []*DescribeModificationPriceResponseBodyPriceInfoPricePromotions) *DescribeModificationPriceResponseBodyPriceInfoPrice {
	s.Promotions = v
	return s
}

type DescribeModificationPriceResponseBodyPriceInfoPricePromotions struct {
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	OptionCode    *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	Selected      *bool   `json:"Selected,omitempty" xml:"Selected,omitempty"`
	PromotionId   *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
}

func (s DescribeModificationPriceResponseBodyPriceInfoPricePromotions) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponseBodyPriceInfoPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) SetPromotionDesc(v string) *DescribeModificationPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) SetOptionCode(v string) *DescribeModificationPriceResponseBodyPriceInfoPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) SetSelected(v bool) *DescribeModificationPriceResponseBodyPriceInfoPricePromotions {
	s.Selected = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) SetPromotionId(v string) *DescribeModificationPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) SetPromotionName(v string) *DescribeModificationPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionName = &v
	return s
}

type DescribeModificationPriceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeModificationPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeModificationPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeModificationPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponse) SetHeaders(v map[string]*string) *DescribeModificationPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeModificationPriceResponse) SetBody(v *DescribeModificationPriceResponseBody) *DescribeModificationPriceResponse {
	s.Body = v
	return s
}

type DescribeNASFileSystemsRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OfficeSiteId *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	MaxResults   *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	FileSystemId []*string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty" type:"Repeated"`
}

func (s DescribeNASFileSystemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNASFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsRequest) SetRegionId(v string) *DescribeNASFileSystemsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetOfficeSiteId(v string) *DescribeNASFileSystemsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetMaxResults(v int32) *DescribeNASFileSystemsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetNextToken(v string) *DescribeNASFileSystemsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeNASFileSystemsRequest) SetFileSystemId(v []*string) *DescribeNASFileSystemsRequest {
	s.FileSystemId = v
	return s
}

type DescribeNASFileSystemsResponseBody struct {
	NextToken   *string                                          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FileSystems []*DescribeNASFileSystemsResponseBodyFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Repeated"`
}

func (s DescribeNASFileSystemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNASFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsResponseBody) SetNextToken(v string) *DescribeNASFileSystemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBody) SetRequestId(v string) *DescribeNASFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBody) SetFileSystems(v []*DescribeNASFileSystemsResponseBodyFileSystems) *DescribeNASFileSystemsResponseBody {
	s.FileSystems = v
	return s
}

type DescribeNASFileSystemsResponseBodyFileSystems struct {
	Capacity          *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	MountTargetStatus *string `json:"MountTargetStatus,omitempty" xml:"MountTargetStatus,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	OfficeSiteId      *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	SupportAcl        *bool   `json:"SupportAcl,omitempty" xml:"SupportAcl,omitempty"`
	StorageType       *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	OfficeSiteName    *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemType    *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	FileSystemName    *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	MeteredSize       *int64  `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ZoneId            *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	FileSystemStatus  *string `json:"FileSystemStatus,omitempty" xml:"FileSystemStatus,omitempty"`
}

func (s DescribeNASFileSystemsResponseBodyFileSystems) String() string {
	return tea.Prettify(s)
}

func (s DescribeNASFileSystemsResponseBodyFileSystems) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetCapacity(v int64) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.Capacity = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetMountTargetStatus(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.MountTargetStatus = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetCreateTime(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.CreateTime = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetOfficeSiteId(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetSupportAcl(v bool) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.SupportAcl = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetStorageType(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.StorageType = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetOfficeSiteName(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetRegionId(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.RegionId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetFileSystemId(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.FileSystemId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetFileSystemType(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.FileSystemType = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetFileSystemName(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.FileSystemName = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetMeteredSize(v int64) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.MeteredSize = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetMountTargetDomain(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetDescription(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.Description = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetZoneId(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.ZoneId = &v
	return s
}

func (s *DescribeNASFileSystemsResponseBodyFileSystems) SetFileSystemStatus(v string) *DescribeNASFileSystemsResponseBodyFileSystems {
	s.FileSystemStatus = &v
	return s
}

type DescribeNASFileSystemsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNASFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNASFileSystemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNASFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNASFileSystemsResponse) SetHeaders(v map[string]*string) *DescribeNASFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNASFileSystemsResponse) SetBody(v *DescribeNASFileSystemsResponseBody) *DescribeNASFileSystemsResponse {
	s.Body = v
	return s
}

type DescribeNetworkPackagesRequest struct {
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults       *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	NetworkPackageId []*string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" type:"Repeated"`
}

func (s DescribeNetworkPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkPackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesRequest) SetRegionId(v string) *DescribeNetworkPackagesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetMaxResults(v int32) *DescribeNetworkPackagesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetNextToken(v string) *DescribeNetworkPackagesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeNetworkPackagesRequest) SetNetworkPackageId(v []*string) *DescribeNetworkPackagesRequest {
	s.NetworkPackageId = v
	return s
}

type DescribeNetworkPackagesResponseBody struct {
	NextToken       *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NetworkPackages []*DescribeNetworkPackagesResponseBodyNetworkPackages `json:"NetworkPackages,omitempty" xml:"NetworkPackages,omitempty" type:"Repeated"`
}

func (s DescribeNetworkPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesResponseBody) SetNextToken(v string) *DescribeNetworkPackagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBody) SetRequestId(v string) *DescribeNetworkPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBody) SetNetworkPackages(v []*DescribeNetworkPackagesResponseBodyNetworkPackages) *DescribeNetworkPackagesResponseBody {
	s.NetworkPackages = v
	return s
}

type DescribeNetworkPackagesResponseBodyNetworkPackages struct {
	NetworkPackageId     *string   `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	Bandwidth            *int32    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ExpiredTime          *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	CreateTime           *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	OfficeSiteId         *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	InternetChargeType   *string   `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	NetworkPackageStatus *string   `json:"NetworkPackageStatus,omitempty" xml:"NetworkPackageStatus,omitempty"`
	OfficeSiteName       *string   `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	EipAddresses         []*string `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty" type:"Repeated"`
}

func (s DescribeNetworkPackagesResponseBodyNetworkPackages) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkPackagesResponseBodyNetworkPackages) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetNetworkPackageId(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.NetworkPackageId = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetBandwidth(v int32) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.Bandwidth = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetExpiredTime(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetCreateTime(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.CreateTime = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetOfficeSiteId(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetInternetChargeType(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetNetworkPackageStatus(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.NetworkPackageStatus = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetOfficeSiteName(v string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeNetworkPackagesResponseBodyNetworkPackages) SetEipAddresses(v []*string) *DescribeNetworkPackagesResponseBodyNetworkPackages {
	s.EipAddresses = v
	return s
}

type DescribeNetworkPackagesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNetworkPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNetworkPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetworkPackagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesResponse) SetHeaders(v map[string]*string) *DescribeNetworkPackagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkPackagesResponse) SetBody(v *DescribeNetworkPackagesResponseBody) *DescribeNetworkPackagesResponse {
	s.Body = v
	return s
}

type DescribeOfficeSitesRequest struct {
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OfficeSiteType *string   `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	MaxResults     *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSiteId   []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
}

func (s DescribeOfficeSitesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesRequest) SetRegionId(v string) *DescribeOfficeSitesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetOfficeSiteType(v string) *DescribeOfficeSitesRequest {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetMaxResults(v int32) *DescribeOfficeSitesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetNextToken(v string) *DescribeOfficeSitesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetOfficeSiteId(v []*string) *DescribeOfficeSitesRequest {
	s.OfficeSiteId = v
	return s
}

type DescribeOfficeSitesResponseBody struct {
	NextToken   *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OfficeSites []*DescribeOfficeSitesResponseBodyOfficeSites `json:"OfficeSites,omitempty" xml:"OfficeSites,omitempty" type:"Repeated"`
}

func (s DescribeOfficeSitesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBody) SetNextToken(v string) *DescribeOfficeSitesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeOfficeSitesResponseBody) SetRequestId(v string) *DescribeOfficeSitesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBody) SetOfficeSites(v []*DescribeOfficeSitesResponseBodyOfficeSites) *DescribeOfficeSitesResponseBody {
	s.OfficeSites = v
	return s
}

type DescribeOfficeSitesResponseBodyOfficeSites struct {
	Status                   *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	CreationTime             *string                                                   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	VpcId                    *string                                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	EnableAdminAccess        *bool                                                     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	EnableCrossDesktopAccess *bool                                                     `json:"EnableCrossDesktopAccess,omitempty" xml:"EnableCrossDesktopAccess,omitempty"`
	DesktopVpcEndpoint       *string                                                   `json:"DesktopVpcEndpoint,omitempty" xml:"DesktopVpcEndpoint,omitempty"`
	DesktopAccessType        *string                                                   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	DomainName               *string                                                   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	SsoEnabled               *bool                                                     `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty"`
	CidrBlock                *string                                                   `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	Bandwidth                *int32                                                    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	TrustPassword            *string                                                   `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty"`
	Name                     *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	EnableInternetAccess     *bool                                                     `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	DomainPassword           *string                                                   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	CustomSecurityGroupId    *string                                                   `json:"CustomSecurityGroupId,omitempty" xml:"CustomSecurityGroupId,omitempty"`
	DomainUserName           *string                                                   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	SubDomainName            *string                                                   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	OfficeSiteId             *string                                                   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	CenId                    *string                                                   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	MfaEnabled               *bool                                                     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	NetworkPackageId         *string                                                   `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	DnsUserName              *string                                                   `json:"DnsUserName,omitempty" xml:"DnsUserName,omitempty"`
	OfficeSiteType           *string                                                   `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	ADConnectors             []*DescribeOfficeSitesResponseBodyOfficeSitesADConnectors `json:"ADConnectors,omitempty" xml:"ADConnectors,omitempty" type:"Repeated"`
	Logs                     []*DescribeOfficeSitesResponseBodyOfficeSitesLogs         `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	VSwitchIds               []*string                                                 `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	FileSystemIds            []*string                                                 `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty" type:"Repeated"`
	SubDnsAddress            []*string                                                 `json:"SubDnsAddress,omitempty" xml:"SubDnsAddress,omitempty" type:"Repeated"`
	DnsAddress               []*string                                                 `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
}

func (s DescribeOfficeSitesResponseBodyOfficeSites) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponseBodyOfficeSites) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetStatus(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.Status = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCreationTime(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CreationTime = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetVpcId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.VpcId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetEnableAdminAccess(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.EnableAdminAccess = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetEnableCrossDesktopAccess(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.EnableCrossDesktopAccess = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDesktopVpcEndpoint(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DesktopVpcEndpoint = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDesktopAccessType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDomainName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DomainName = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetSsoEnabled(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.SsoEnabled = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCidrBlock(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CidrBlock = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetBandwidth(v int32) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.Bandwidth = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetTrustPassword(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.TrustPassword = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.Name = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetEnableInternetAccess(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.EnableInternetAccess = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDomainPassword(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DomainPassword = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCustomSecurityGroupId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CustomSecurityGroupId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDomainUserName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DomainUserName = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetSubDomainName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.SubDomainName = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetOfficeSiteId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetCenId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.CenId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetMfaEnabled(v bool) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.MfaEnabled = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetNetworkPackageId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.NetworkPackageId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDnsUserName(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DnsUserName = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetOfficeSiteType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetADConnectors(v []*DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.ADConnectors = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetLogs(v []*DescribeOfficeSitesResponseBodyOfficeSitesLogs) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.Logs = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetVSwitchIds(v []*string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.VSwitchIds = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetFileSystemIds(v []*string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.FileSystemIds = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetSubDnsAddress(v []*string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.SubDnsAddress = v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDnsAddress(v []*string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DnsAddress = v
	return s
}

type DescribeOfficeSitesResponseBodyOfficeSitesADConnectors struct {
	ConnectorStatus    *string `json:"ConnectorStatus,omitempty" xml:"ConnectorStatus,omitempty"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ADConnectorAddress *string `json:"ADConnectorAddress,omitempty" xml:"ADConnectorAddress,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
}

func (s DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) SetConnectorStatus(v string) *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	s.ConnectorStatus = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) SetVSwitchId(v string) *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	s.VSwitchId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) SetADConnectorAddress(v string) *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	s.ADConnectorAddress = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors) SetNetworkInterfaceId(v string) *DescribeOfficeSitesResponseBodyOfficeSitesADConnectors {
	s.NetworkInterfaceId = &v
	return s
}

type DescribeOfficeSitesResponseBodyOfficeSitesLogs struct {
	Step      *string `json:"Step,omitempty" xml:"Step,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Level     *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeOfficeSitesResponseBodyOfficeSitesLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponseBodyOfficeSitesLogs) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) SetStep(v string) *DescribeOfficeSitesResponseBodyOfficeSitesLogs {
	s.Step = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) SetMessage(v string) *DescribeOfficeSitesResponseBodyOfficeSitesLogs {
	s.Message = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) SetTimeStamp(v string) *DescribeOfficeSitesResponseBodyOfficeSitesLogs {
	s.TimeStamp = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSitesLogs) SetLevel(v string) *DescribeOfficeSitesResponseBodyOfficeSitesLogs {
	s.Level = &v
	return s
}

type DescribeOfficeSitesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOfficeSitesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOfficeSitesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponse) SetHeaders(v map[string]*string) *DescribeOfficeSitesResponse {
	s.Headers = v
	return s
}

func (s *DescribeOfficeSitesResponse) SetBody(v *DescribeOfficeSitesResponseBody) *DescribeOfficeSitesResponse {
	s.Body = v
	return s
}

type DescribePolicyGroupsRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults    *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PolicyGroupId []*string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" type:"Repeated"`
}

func (s DescribePolicyGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsRequest) SetRegionId(v string) *DescribePolicyGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePolicyGroupsRequest) SetMaxResults(v int32) *DescribePolicyGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePolicyGroupsRequest) SetNextToken(v string) *DescribePolicyGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePolicyGroupsRequest) SetPolicyGroupId(v []*string) *DescribePolicyGroupsRequest {
	s.PolicyGroupId = v
	return s
}

type DescribePolicyGroupsResponseBody struct {
	NextToken            *string                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId            *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DescribePolicyGroups []*DescribePolicyGroupsResponseBodyDescribePolicyGroups `json:"DescribePolicyGroups,omitempty" xml:"DescribePolicyGroups,omitempty" type:"Repeated"`
}

func (s DescribePolicyGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBody) SetNextToken(v string) *DescribePolicyGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribePolicyGroupsResponseBody) SetRequestId(v string) *DescribePolicyGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolicyGroupsResponseBody) SetDescribePolicyGroups(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroups) *DescribePolicyGroupsResponseBody {
	s.DescribePolicyGroups = v
	return s
}

type DescribePolicyGroupsResponseBodyDescribePolicyGroups struct {
	PolicyStatus                 *string                                                                             `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	Html5Access                  *string                                                                             `json:"Html5Access,omitempty" xml:"Html5Access,omitempty"`
	WatermarkType                *string                                                                             `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	PreemptLogin                 *string                                                                             `json:"PreemptLogin,omitempty" xml:"PreemptLogin,omitempty"`
	WatermarkCustomText          *string                                                                             `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	Clipboard                    *string                                                                             `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	DomainList                   *string                                                                             `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	PolicyGroupId                *string                                                                             `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	WatermarkTransparency        *string                                                                             `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
	Html5FileTransfer            *string                                                                             `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	UsbRedirect                  *string                                                                             `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	PolicyGroupType              *string                                                                             `json:"PolicyGroupType,omitempty" xml:"PolicyGroupType,omitempty"`
	Watermark                    *string                                                                             `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	VisualQuality                *string                                                                             `json:"VisualQuality,omitempty" xml:"VisualQuality,omitempty"`
	EdsCount                     *int32                                                                              `json:"EdsCount,omitempty" xml:"EdsCount,omitempty"`
	Name                         *string                                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	LocalDrive                   *string                                                                             `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	AuthorizeSecurityPolicyRules []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules `json:"AuthorizeSecurityPolicyRules,omitempty" xml:"AuthorizeSecurityPolicyRules,omitempty" type:"Repeated"`
	AuthorizeAccessPolicyRules   []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules   `json:"AuthorizeAccessPolicyRules,omitempty" xml:"AuthorizeAccessPolicyRules,omitempty" type:"Repeated"`
	PreemptLoginUsers            []*string                                                                           `json:"PreemptLoginUsers,omitempty" xml:"PreemptLoginUsers,omitempty" type:"Repeated"`
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroups) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPolicyStatus(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PolicyStatus = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetHtml5Access(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Html5Access = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPreemptLogin(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PreemptLogin = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkCustomText(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkCustomText = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetClipboard(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Clipboard = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetDomainList(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.DomainList = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPolicyGroupId(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkTransparency(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkTransparency = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetHtml5FileTransfer(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Html5FileTransfer = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetUsbRedirect(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.UsbRedirect = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPolicyGroupType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PolicyGroupType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermark(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Watermark = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetVisualQuality(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.VisualQuality = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetEdsCount(v int32) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.EdsCount = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetName(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Name = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetLocalDrive(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.LocalDrive = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetAuthorizeSecurityPolicyRules(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.AuthorizeSecurityPolicyRules = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetAuthorizeAccessPolicyRules(v []*DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.AuthorizeAccessPolicyRules = v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPreemptLoginUsers(v []*string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PreemptLoginUsers = v
	return s
}

type DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PortRange   *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	IpProtocol  *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Priority    *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Type = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPolicy(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Policy = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetDescription(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Description = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPortRange(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.PortRange = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetIpProtocol(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.IpProtocol = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetPriority(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.Priority = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules) SetCidrIp(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeSecurityPolicyRules {
	s.CidrIp = &v
	return s
}

type DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) SetDescription(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules {
	s.Description = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules) SetCidrIp(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroupsAuthorizeAccessPolicyRules {
	s.CidrIp = &v
	return s
}

type DescribePolicyGroupsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePolicyGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePolicyGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponse) SetHeaders(v map[string]*string) *DescribePolicyGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyGroupsResponse) SetBody(v *DescribePolicyGroupsResponseBody) *DescribePolicyGroupsResponse {
	s.Body = v
	return s
}

type DescribePostPaidDesktopBillsRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	BillStartTime  *string `json:"BillStartTime,omitempty" xml:"BillStartTime,omitempty"`
	BillEndTime    *string `json:"BillEndTime,omitempty" xml:"BillEndTime,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribePostPaidDesktopBillsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePostPaidDesktopBillsRequest) GoString() string {
	return s.String()
}

func (s *DescribePostPaidDesktopBillsRequest) SetRegionId(v string) *DescribePostPaidDesktopBillsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePostPaidDesktopBillsRequest) SetDesktopGroupId(v string) *DescribePostPaidDesktopBillsRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribePostPaidDesktopBillsRequest) SetBillStartTime(v string) *DescribePostPaidDesktopBillsRequest {
	s.BillStartTime = &v
	return s
}

func (s *DescribePostPaidDesktopBillsRequest) SetBillEndTime(v string) *DescribePostPaidDesktopBillsRequest {
	s.BillEndTime = &v
	return s
}

func (s *DescribePostPaidDesktopBillsRequest) SetMaxResults(v int32) *DescribePostPaidDesktopBillsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePostPaidDesktopBillsRequest) SetNextToken(v string) *DescribePostPaidDesktopBillsRequest {
	s.NextToken = &v
	return s
}

type DescribePostPaidDesktopBillsResponseBody struct {
	PostPaidDesktopsBillsUrl    *string                                          `json:"PostPaidDesktopsBillsUrl,omitempty" xml:"PostPaidDesktopsBillsUrl,omitempty"`
	PostPaidDesktopsCount       *int32                                           `json:"PostPaidDesktopsCount,omitempty" xml:"PostPaidDesktopsCount,omitempty"`
	NextToken                   *string                                          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PostPaidDesktopsTotalAmount *float32                                         `json:"PostPaidDesktopsTotalAmount,omitempty" xml:"PostPaidDesktopsTotalAmount,omitempty"`
	Bills                       []*DescribePostPaidDesktopBillsResponseBodyBills `json:"Bills,omitempty" xml:"Bills,omitempty" type:"Repeated"`
}

func (s DescribePostPaidDesktopBillsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePostPaidDesktopBillsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePostPaidDesktopBillsResponseBody) SetPostPaidDesktopsBillsUrl(v string) *DescribePostPaidDesktopBillsResponseBody {
	s.PostPaidDesktopsBillsUrl = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBody) SetPostPaidDesktopsCount(v int32) *DescribePostPaidDesktopBillsResponseBody {
	s.PostPaidDesktopsCount = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBody) SetNextToken(v string) *DescribePostPaidDesktopBillsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBody) SetRequestId(v string) *DescribePostPaidDesktopBillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBody) SetPostPaidDesktopsTotalAmount(v float32) *DescribePostPaidDesktopBillsResponseBody {
	s.PostPaidDesktopsTotalAmount = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBody) SetBills(v []*DescribePostPaidDesktopBillsResponseBodyBills) *DescribePostPaidDesktopBillsResponseBody {
	s.Bills = v
	return s
}

type DescribePostPaidDesktopBillsResponseBodyBills struct {
	BillId            *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	DiscountPrice     *string `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	Product           *string `json:"Product,omitempty" xml:"Product,omitempty"`
	PriceUnit         *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	CashPayment       *string `json:"CashPayment,omitempty" xml:"CashPayment,omitempty"`
	Payment           *string `json:"Payment,omitempty" xml:"Payment,omitempty"`
	OriginalPrice     *string `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProductDetail     *string `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	Usage             *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
	GoldNote          *string `json:"GoldNote,omitempty" xml:"GoldNote,omitempty"`
	UsageUnit         *string `json:"UsageUnit,omitempty" xml:"UsageUnit,omitempty"`
	Price             *string `json:"Price,omitempty" xml:"Price,omitempty"`
	BillStartTime     *string `json:"BillStartTime,omitempty" xml:"BillStartTime,omitempty"`
	BillType          *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ChargeItem        *string `json:"ChargeItem,omitempty" xml:"ChargeItem,omitempty"`
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	ConsumeTime       *string `json:"ConsumeTime,omitempty" xml:"ConsumeTime,omitempty"`
	ConsumeType       *string `json:"ConsumeType,omitempty" xml:"ConsumeType,omitempty"`
	BillEndTime       *string `json:"BillEndTime,omitempty" xml:"BillEndTime,omitempty"`
}

func (s DescribePostPaidDesktopBillsResponseBodyBills) String() string {
	return tea.Prettify(s)
}

func (s DescribePostPaidDesktopBillsResponseBodyBills) GoString() string {
	return s.String()
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetBillId(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.BillId = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetDiscountPrice(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetProduct(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.Product = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetPriceUnit(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.PriceUnit = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetCashPayment(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.CashPayment = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetPayment(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.Payment = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetOriginalPrice(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetInstanceId(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.InstanceId = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetProductDetail(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.ProductDetail = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetUsage(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.Usage = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetGoldNote(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.GoldNote = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetUsageUnit(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.UsageUnit = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetPrice(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.Price = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetBillStartTime(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.BillStartTime = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetBillType(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.BillType = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetResourceGroupId(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetChargeItem(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.ChargeItem = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetResourceGroupName(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetConsumeTime(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.ConsumeTime = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetConsumeType(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.ConsumeType = &v
	return s
}

func (s *DescribePostPaidDesktopBillsResponseBodyBills) SetBillEndTime(v string) *DescribePostPaidDesktopBillsResponseBodyBills {
	s.BillEndTime = &v
	return s
}

type DescribePostPaidDesktopBillsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePostPaidDesktopBillsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePostPaidDesktopBillsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePostPaidDesktopBillsResponse) GoString() string {
	return s.String()
}

func (s *DescribePostPaidDesktopBillsResponse) SetHeaders(v map[string]*string) *DescribePostPaidDesktopBillsResponse {
	s.Headers = v
	return s
}

func (s *DescribePostPaidDesktopBillsResponse) SetBody(v *DescribePostPaidDesktopBillsResponseBody) *DescribePostPaidDesktopBillsResponse {
	s.Body = v
	return s
}

type DescribePriceRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	InstanceType       *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RootDiskSizeGib    *int32  `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	UserDiskSizeGib    *int32  `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty"`
	OsType             *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PeriodUnit         *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Period             *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	Amount             *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	PromotionId        *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	Bandwidth          *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
}

func (s DescribePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) SetRegionId(v string) *DescribePriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePriceRequest) SetResourceType(v string) *DescribePriceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceType(v string) *DescribePriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribePriceRequest) SetRootDiskSizeGib(v int32) *DescribePriceRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *DescribePriceRequest) SetUserDiskSizeGib(v int32) *DescribePriceRequest {
	s.UserDiskSizeGib = &v
	return s
}

func (s *DescribePriceRequest) SetOsType(v string) *DescribePriceRequest {
	s.OsType = &v
	return s
}

func (s *DescribePriceRequest) SetPeriodUnit(v string) *DescribePriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribePriceRequest) SetPeriod(v int32) *DescribePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribePriceRequest) SetAmount(v int32) *DescribePriceRequest {
	s.Amount = &v
	return s
}

func (s *DescribePriceRequest) SetPromotionId(v string) *DescribePriceRequest {
	s.PromotionId = &v
	return s
}

func (s *DescribePriceRequest) SetInternetChargeType(v string) *DescribePriceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribePriceRequest) SetBandwidth(v int32) *DescribePriceRequest {
	s.Bandwidth = &v
	return s
}

type DescribePriceResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PriceInfo *DescribePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
}

func (s DescribePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBody) SetRequestId(v string) *DescribePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePriceResponseBody) SetPriceInfo(v *DescribePriceResponseBodyPriceInfo) *DescribePriceResponseBody {
	s.PriceInfo = v
	return s
}

type DescribePriceResponseBodyPriceInfo struct {
	Rules []*DescribePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	Price *DescribePriceResponseBodyPriceInfoPrice   `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
}

func (s DescribePriceResponseBodyPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfo) SetRules(v []*DescribePriceResponseBodyPriceInfoRules) *DescribePriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetPrice(v *DescribePriceResponseBodyPriceInfoPrice) *DescribePriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

type DescribePriceResponseBodyPriceInfoRules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoRules) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoRules) SetDescription(v string) *DescribePriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoRules) SetRuleId(v int64) *DescribePriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

type DescribePriceResponseBodyPriceInfoPrice struct {
	OriginalPrice *float32                                             `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	DiscountPrice *float32                                             `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	Currency      *string                                              `json:"Currency,omitempty" xml:"Currency,omitempty"`
	TradePrice    *float32                                             `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
	Promotions    []*DescribePriceResponseBodyPriceInfoPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyPriceInfoPrice) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribePriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetPromotions(v []*DescribePriceResponseBodyPriceInfoPricePromotions) *DescribePriceResponseBodyPriceInfoPrice {
	s.Promotions = v
	return s
}

type DescribePriceResponseBodyPriceInfoPricePromotions struct {
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	OptionCode    *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	Selected      *bool   `json:"Selected,omitempty" xml:"Selected,omitempty"`
	PromotionId   *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoPricePromotions) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetPromotionDesc(v string) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetOptionCode(v string) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetSelected(v bool) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.Selected = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetPromotionId(v string) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetPromotionName(v string) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.PromotionName = &v
	return s
}

type DescribePriceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribePriceResponse) SetHeaders(v map[string]*string) *DescribePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribePriceResponse) SetBody(v *DescribePriceResponseBody) *DescribePriceResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeScaleStrategysRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScaleStrategyName *string `json:"ScaleStrategyName,omitempty" xml:"ScaleStrategyName,omitempty"`
	MaxResults        *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeScaleStrategysRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScaleStrategysRequest) GoString() string {
	return s.String()
}

func (s *DescribeScaleStrategysRequest) SetRegionId(v string) *DescribeScaleStrategysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScaleStrategysRequest) SetScaleStrategyName(v string) *DescribeScaleStrategysRequest {
	s.ScaleStrategyName = &v
	return s
}

func (s *DescribeScaleStrategysRequest) SetMaxResults(v int32) *DescribeScaleStrategysRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeScaleStrategysRequest) SetNextToken(v string) *DescribeScaleStrategysRequest {
	s.NextToken = &v
	return s
}

type DescribeScaleStrategysResponseBody struct {
	NextToken      *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScaleStrategys []*DescribeScaleStrategysResponseBodyScaleStrategys `json:"ScaleStrategys,omitempty" xml:"ScaleStrategys,omitempty" type:"Repeated"`
}

func (s DescribeScaleStrategysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScaleStrategysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScaleStrategysResponseBody) SetNextToken(v string) *DescribeScaleStrategysResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeScaleStrategysResponseBody) SetRequestId(v string) *DescribeScaleStrategysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScaleStrategysResponseBody) SetScaleStrategys(v []*DescribeScaleStrategysResponseBodyScaleStrategys) *DescribeScaleStrategysResponseBody {
	s.ScaleStrategys = v
	return s
}

type DescribeScaleStrategysResponseBodyScaleStrategys struct {
	ScaleStrategyId           *string `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty"`
	MaxDesktopsCount          *int32  `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	MaxAvailableDesktopsCount *int32  `json:"MaxAvailableDesktopsCount,omitempty" xml:"MaxAvailableDesktopsCount,omitempty"`
	ScaleStrategyName         *string `json:"ScaleStrategyName,omitempty" xml:"ScaleStrategyName,omitempty"`
	ScaleStrategyType         *string `json:"ScaleStrategyType,omitempty" xml:"ScaleStrategyType,omitempty"`
	MinDesktopsCount          *int32  `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	MinAvailableDesktopsCount *int32  `json:"MinAvailableDesktopsCount,omitempty" xml:"MinAvailableDesktopsCount,omitempty"`
	ScaleStep                 *int32  `json:"ScaleStep,omitempty" xml:"ScaleStep,omitempty"`
}

func (s DescribeScaleStrategysResponseBodyScaleStrategys) String() string {
	return tea.Prettify(s)
}

func (s DescribeScaleStrategysResponseBodyScaleStrategys) GoString() string {
	return s.String()
}

func (s *DescribeScaleStrategysResponseBodyScaleStrategys) SetScaleStrategyId(v string) *DescribeScaleStrategysResponseBodyScaleStrategys {
	s.ScaleStrategyId = &v
	return s
}

func (s *DescribeScaleStrategysResponseBodyScaleStrategys) SetMaxDesktopsCount(v int32) *DescribeScaleStrategysResponseBodyScaleStrategys {
	s.MaxDesktopsCount = &v
	return s
}

func (s *DescribeScaleStrategysResponseBodyScaleStrategys) SetMaxAvailableDesktopsCount(v int32) *DescribeScaleStrategysResponseBodyScaleStrategys {
	s.MaxAvailableDesktopsCount = &v
	return s
}

func (s *DescribeScaleStrategysResponseBodyScaleStrategys) SetScaleStrategyName(v string) *DescribeScaleStrategysResponseBodyScaleStrategys {
	s.ScaleStrategyName = &v
	return s
}

func (s *DescribeScaleStrategysResponseBodyScaleStrategys) SetScaleStrategyType(v string) *DescribeScaleStrategysResponseBodyScaleStrategys {
	s.ScaleStrategyType = &v
	return s
}

func (s *DescribeScaleStrategysResponseBodyScaleStrategys) SetMinDesktopsCount(v int32) *DescribeScaleStrategysResponseBodyScaleStrategys {
	s.MinDesktopsCount = &v
	return s
}

func (s *DescribeScaleStrategysResponseBodyScaleStrategys) SetMinAvailableDesktopsCount(v int32) *DescribeScaleStrategysResponseBodyScaleStrategys {
	s.MinAvailableDesktopsCount = &v
	return s
}

func (s *DescribeScaleStrategysResponseBodyScaleStrategys) SetScaleStep(v int32) *DescribeScaleStrategysResponseBodyScaleStrategys {
	s.ScaleStep = &v
	return s
}

type DescribeScaleStrategysResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScaleStrategysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScaleStrategysResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScaleStrategysResponse) GoString() string {
	return s.String()
}

func (s *DescribeScaleStrategysResponse) SetHeaders(v map[string]*string) *DescribeScaleStrategysResponse {
	s.Headers = v
	return s
}

func (s *DescribeScaleStrategysResponse) SetBody(v *DescribeScaleStrategysResponseBody) *DescribeScaleStrategysResponse {
	s.Body = v
	return s
}

type DescribeScanTaskProgressRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskId   *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeScanTaskProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScanTaskProgressRequest) GoString() string {
	return s.String()
}

func (s *DescribeScanTaskProgressRequest) SetRegionId(v string) *DescribeScanTaskProgressRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScanTaskProgressRequest) SetTaskId(v int64) *DescribeScanTaskProgressRequest {
	s.TaskId = &v
	return s
}

type DescribeScanTaskProgressResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s DescribeScanTaskProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScanTaskProgressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScanTaskProgressResponseBody) SetRequestId(v string) *DescribeScanTaskProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScanTaskProgressResponseBody) SetTaskStatus(v string) *DescribeScanTaskProgressResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *DescribeScanTaskProgressResponseBody) SetCreateTime(v string) *DescribeScanTaskProgressResponseBody {
	s.CreateTime = &v
	return s
}

type DescribeScanTaskProgressResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScanTaskProgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScanTaskProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScanTaskProgressResponse) GoString() string {
	return s.String()
}

func (s *DescribeScanTaskProgressResponse) SetHeaders(v map[string]*string) *DescribeScanTaskProgressResponse {
	s.Headers = v
	return s
}

func (s *DescribeScanTaskProgressResponse) SetBody(v *DescribeScanTaskProgressResponseBody) *DescribeScanTaskProgressResponse {
	s.Body = v
	return s
}

type DescribeSecurityEventOperationsRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityEventId *int64  `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
}

func (s DescribeSecurityEventOperationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsRequest) SetRegionId(v string) *DescribeSecurityEventOperationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityEventOperationsRequest) SetSecurityEventId(v int64) *DescribeSecurityEventOperationsRequest {
	s.SecurityEventId = &v
	return s
}

type DescribeSecurityEventOperationsResponseBody struct {
	RequestId               *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityEventOperations []*DescribeSecurityEventOperationsResponseBodySecurityEventOperations `json:"SecurityEventOperations,omitempty" xml:"SecurityEventOperations,omitempty" type:"Repeated"`
}

func (s DescribeSecurityEventOperationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponseBody) SetRequestId(v string) *DescribeSecurityEventOperationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBody) SetSecurityEventOperations(v []*DescribeSecurityEventOperationsResponseBodySecurityEventOperations) *DescribeSecurityEventOperationsResponseBody {
	s.SecurityEventOperations = v
	return s
}

type DescribeSecurityEventOperationsResponseBodySecurityEventOperations struct {
	OperationParams *string `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	OperationCode   *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	UserCanOperate  *bool   `json:"UserCanOperate,omitempty" xml:"UserCanOperate,omitempty"`
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperations) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperations) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperations) SetOperationParams(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperations {
	s.OperationParams = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperations) SetOperationCode(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperations {
	s.OperationCode = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperations) SetUserCanOperate(v bool) *DescribeSecurityEventOperationsResponseBodySecurityEventOperations {
	s.UserCanOperate = &v
	return s
}

type DescribeSecurityEventOperationsResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecurityEventOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityEventOperationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponse) SetHeaders(v map[string]*string) *DescribeSecurityEventOperationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityEventOperationsResponse) SetBody(v *DescribeSecurityEventOperationsResponseBody) *DescribeSecurityEventOperationsResponse {
	s.Body = v
	return s
}

type DescribeSecurityEventOperationStatusRequest struct {
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskId          *int64    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	SecurityEventId []*string `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty" type:"Repeated"`
}

func (s DescribeSecurityEventOperationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusRequest) SetRegionId(v string) *DescribeSecurityEventOperationStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusRequest) SetTaskId(v int64) *DescribeSecurityEventOperationStatusRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusRequest) SetSecurityEventId(v []*string) *DescribeSecurityEventOperationStatusRequest {
	s.SecurityEventId = v
	return s
}

type DescribeSecurityEventOperationStatusResponseBody struct {
	TaskStatus                     *string                                                                           `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	RequestId                      *string                                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityEventOperationStatuses []*DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses `json:"SecurityEventOperationStatuses,omitempty" xml:"SecurityEventOperationStatuses,omitempty" type:"Repeated"`
}

func (s DescribeSecurityEventOperationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusResponseBody) SetTaskStatus(v string) *DescribeSecurityEventOperationStatusResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBody) SetRequestId(v string) *DescribeSecurityEventOperationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBody) SetSecurityEventOperationStatuses(v []*DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses) *DescribeSecurityEventOperationStatusResponseBody {
	s.SecurityEventOperationStatuses = v
	return s
}

type DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses struct {
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SecurityEventId *int64  `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	ErrorCode       *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses) SetStatus(v string) *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses {
	s.Status = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses) SetSecurityEventId(v int64) *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses {
	s.SecurityEventId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses) SetErrorCode(v string) *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatuses {
	s.ErrorCode = &v
	return s
}

type DescribeSecurityEventOperationStatusResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecurityEventOperationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityEventOperationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusResponse) SetHeaders(v map[string]*string) *DescribeSecurityEventOperationStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponse) SetBody(v *DescribeSecurityEventOperationStatusResponseBody) *DescribeSecurityEventOperationStatusResponse {
	s.Body = v
	return s
}

type DescribeSnapshotsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopId  *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequest) SetRegionId(v string) *DescribeSnapshotsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetDesktopId(v string) *DescribeSnapshotsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotId(v string) *DescribeSnapshotsRequest {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetMaxResults(v int32) *DescribeSnapshotsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetNextToken(v string) *DescribeSnapshotsRequest {
	s.NextToken = &v
	return s
}

type DescribeSnapshotsResponseBody struct {
	NextToken *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshots []*DescribeSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBody) SetNextToken(v string) *DescribeSnapshotsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetRequestId(v string) *DescribeSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetSnapshots(v []*DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

type DescribeSnapshotsResponseBodySnapshots struct {
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SnapshotType   *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	SnapshotName   *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	Progress       *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SnapshotId     *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	RemainTime     *int32  `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	SourceDiskSize *string `json:"SourceDiskSize,omitempty" xml:"SourceDiskSize,omitempty"`
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
}

func (s DescribeSnapshotsResponseBodySnapshots) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetCreationTime(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.CreationTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetStatus(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Status = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotType(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotName(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetProgress(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Progress = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetDescription(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Description = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetRemainTime(v int32) *DescribeSnapshotsResponseBodySnapshots {
	s.RemainTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSourceDiskSize(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SourceDiskSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSourceDiskType(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SourceDiskType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetDesktopId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.DesktopId = &v
	return s
}

type DescribeSnapshotsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponse) SetHeaders(v map[string]*string) *DescribeSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnapshotsResponse) SetBody(v *DescribeSnapshotsResponseBody) *DescribeSnapshotsResponse {
	s.Body = v
	return s
}

type DescribeSuspEventOverviewRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSuspEventOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventOverviewRequest) SetRegionId(v string) *DescribeSuspEventOverviewRequest {
	s.RegionId = &v
	return s
}

type DescribeSuspEventOverviewResponseBody struct {
	SuspiciousCount *int32  `json:"SuspiciousCount,omitempty" xml:"SuspiciousCount,omitempty"`
	RemindCount     *int32  `json:"RemindCount,omitempty" xml:"RemindCount,omitempty"`
	SeriousCount    *int32  `json:"SeriousCount,omitempty" xml:"SeriousCount,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSuspEventOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventOverviewResponseBody) SetSuspiciousCount(v int32) *DescribeSuspEventOverviewResponseBody {
	s.SuspiciousCount = &v
	return s
}

func (s *DescribeSuspEventOverviewResponseBody) SetRemindCount(v int32) *DescribeSuspEventOverviewResponseBody {
	s.RemindCount = &v
	return s
}

func (s *DescribeSuspEventOverviewResponseBody) SetSeriousCount(v int32) *DescribeSuspEventOverviewResponseBody {
	s.SeriousCount = &v
	return s
}

func (s *DescribeSuspEventOverviewResponseBody) SetRequestId(v string) *DescribeSuspEventOverviewResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSuspEventOverviewResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSuspEventOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSuspEventOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventOverviewResponse) SetHeaders(v map[string]*string) *DescribeSuspEventOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspEventOverviewResponse) SetBody(v *DescribeSuspEventOverviewResponseBody) *DescribeSuspEventOverviewResponse {
	s.Body = v
	return s
}

type DescribeSuspEventQuaraFilesRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSuspEventQuaraFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesRequest) SetRegionId(v string) *DescribeSuspEventQuaraFilesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetStatus(v string) *DescribeSuspEventQuaraFilesRequest {
	s.Status = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetOfficeSiteId(v string) *DescribeSuspEventQuaraFilesRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetCurrentPage(v int32) *DescribeSuspEventQuaraFilesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetPageSize(v int32) *DescribeSuspEventQuaraFilesRequest {
	s.PageSize = &v
	return s
}

type DescribeSuspEventQuaraFilesResponseBody struct {
	CurrentPage *int32                                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	QuaraFiles  []*DescribeSuspEventQuaraFilesResponseBodyQuaraFiles `json:"QuaraFiles,omitempty" xml:"QuaraFiles,omitempty" type:"Repeated"`
}

func (s DescribeSuspEventQuaraFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetCurrentPage(v int32) *DescribeSuspEventQuaraFilesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetRequestId(v string) *DescribeSuspEventQuaraFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetPageSize(v int32) *DescribeSuspEventQuaraFilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetTotalCount(v int32) *DescribeSuspEventQuaraFilesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetQuaraFiles(v []*DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) *DescribeSuspEventQuaraFilesResponseBody {
	s.QuaraFiles = v
	return s
}

type DescribeSuspEventQuaraFilesResponseBodyQuaraFiles struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	EventName   *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	EventType   *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty"`
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	Md5         *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Tag         *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	Id          *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetStatus(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Status = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetEventName(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.EventName = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetEventType(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.EventType = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetPath(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Path = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetDesktopName(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.DesktopName = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetMd5(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Md5 = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetTag(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Tag = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetDesktopId(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.DesktopId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetId(v int32) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetModifyTime(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.ModifyTime = &v
	return s
}

type DescribeSuspEventQuaraFilesResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSuspEventQuaraFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSuspEventQuaraFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesResponse) SetHeaders(v map[string]*string) *DescribeSuspEventQuaraFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponse) SetBody(v *DescribeSuspEventQuaraFilesResponseBody) *DescribeSuspEventQuaraFilesResponse {
	s.Body = v
	return s
}

type DescribeSuspEventsRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OfficeSiteId    *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	Dealed          *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	Levels          *string `json:"Levels,omitempty" xml:"Levels,omitempty"`
	ParentEventType *string `json:"ParentEventType,omitempty" xml:"ParentEventType,omitempty"`
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSuspEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsRequest) SetRegionId(v string) *DescribeSuspEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetLang(v string) *DescribeSuspEventsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetOfficeSiteId(v string) *DescribeSuspEventsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetDealed(v string) *DescribeSuspEventsRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetLevels(v string) *DescribeSuspEventsRequest {
	s.Levels = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetParentEventType(v string) *DescribeSuspEventsRequest {
	s.ParentEventType = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetAlarmUniqueInfo(v string) *DescribeSuspEventsRequest {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetCurrentPage(v int32) *DescribeSuspEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetPageSize(v int32) *DescribeSuspEventsRequest {
	s.PageSize = &v
	return s
}

type DescribeSuspEventsResponseBody struct {
	CurrentPage *int32                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *string                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	SuspEvents  []*DescribeSuspEventsResponseBodySuspEvents `json:"SuspEvents,omitempty" xml:"SuspEvents,omitempty" type:"Repeated"`
}

func (s DescribeSuspEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBody) SetCurrentPage(v int32) *DescribeSuspEventsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetRequestId(v string) *DescribeSuspEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetPageSize(v string) *DescribeSuspEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetTotalCount(v int32) *DescribeSuspEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetSuspEvents(v []*DescribeSuspEventsResponseBodySuspEvents) *DescribeSuspEventsResponseBody {
	s.SuspEvents = v
	return s
}

type DescribeSuspEventsResponseBodySuspEvents struct {
	DataSource            *string                                            `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	EventSubType          *string                                            `json:"EventSubType,omitempty" xml:"EventSubType,omitempty"`
	DesktopName           *string                                            `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DesktopId             *string                                            `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	OccurrenceTime        *string                                            `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	AlarmEventTypeDisplay *string                                            `json:"AlarmEventTypeDisplay,omitempty" xml:"AlarmEventTypeDisplay,omitempty"`
	AlarmUniqueInfo       *string                                            `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	Desc                  *string                                            `json:"Desc,omitempty" xml:"Desc,omitempty"`
	AlarmEventNameDisplay *string                                            `json:"AlarmEventNameDisplay,omitempty" xml:"AlarmEventNameDisplay,omitempty"`
	CanCancelFault        *bool                                              `json:"CanCancelFault,omitempty" xml:"CanCancelFault,omitempty"`
	LastTime              *string                                            `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	OperateMsg            *string                                            `json:"OperateMsg,omitempty" xml:"OperateMsg,omitempty"`
	CanBeDealOnLine       *string                                            `json:"CanBeDealOnLine,omitempty" xml:"CanBeDealOnLine,omitempty"`
	AlarmEventType        *string                                            `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty"`
	EventStatus           *int32                                             `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	OperateErrorCode      *string                                            `json:"OperateErrorCode,omitempty" xml:"OperateErrorCode,omitempty"`
	AlarmEventName        *string                                            `json:"AlarmEventName,omitempty" xml:"AlarmEventName,omitempty"`
	Name                  *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	UniqueInfo            *string                                            `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
	Level                 *string                                            `json:"Level,omitempty" xml:"Level,omitempty"`
	Id                    *int64                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	Details               []*DescribeSuspEventsResponseBodySuspEventsDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
}

func (s DescribeSuspEventsResponseBodySuspEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsResponseBodySuspEvents) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDataSource(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.DataSource = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetEventSubType(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.EventSubType = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDesktopName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.DesktopName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDesktopId(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.DesktopId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetOccurrenceTime(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.OccurrenceTime = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmEventTypeDisplay(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmEventTypeDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmUniqueInfo(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDesc(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.Desc = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmEventNameDisplay(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmEventNameDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetCanCancelFault(v bool) *DescribeSuspEventsResponseBodySuspEvents {
	s.CanCancelFault = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetLastTime(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.LastTime = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetOperateMsg(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.OperateMsg = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetCanBeDealOnLine(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.CanBeDealOnLine = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmEventType(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmEventType = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetEventStatus(v int32) *DescribeSuspEventsResponseBodySuspEvents {
	s.EventStatus = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetOperateErrorCode(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.OperateErrorCode = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmEventName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmEventName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.Name = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetUniqueInfo(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.UniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetLevel(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.Level = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetId(v int64) *DescribeSuspEventsResponseBodySuspEvents {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDetails(v []*DescribeSuspEventsResponseBodySuspEventsDetails) *DescribeSuspEventsResponseBodySuspEvents {
	s.Details = v
	return s
}

type DescribeSuspEventsResponseBodySuspEventsDetails struct {
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	NameDisplay  *string `json:"NameDisplay,omitempty" xml:"NameDisplay,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ValueDisplay *string `json:"ValueDisplay,omitempty" xml:"ValueDisplay,omitempty"`
}

func (s DescribeSuspEventsResponseBodySuspEventsDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsResponseBodySuspEventsDetails) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetType(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.Type = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetValue(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.Value = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetNameDisplay(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.NameDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetName(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.Name = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetValueDisplay(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.ValueDisplay = &v
	return s
}

type DescribeSuspEventsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSuspEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSuspEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponse) SetHeaders(v map[string]*string) *DescribeSuspEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspEventsResponse) SetBody(v *DescribeSuspEventsResponseBody) *DescribeSuspEventsResponse {
	s.Body = v
	return s
}

type DescribeUserConnectionRecordsRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	EndUserId      *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EndUserType    *string `json:"EndUserType,omitempty" xml:"EndUserType,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeUserConnectionRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConnectionRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectionRecordsRequest) SetRegionId(v string) *DescribeUserConnectionRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetDesktopGroupId(v string) *DescribeUserConnectionRecordsRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetEndUserId(v string) *DescribeUserConnectionRecordsRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetEndUserType(v string) *DescribeUserConnectionRecordsRequest {
	s.EndUserType = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetMaxResults(v int32) *DescribeUserConnectionRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUserConnectionRecordsRequest) SetNextToken(v string) *DescribeUserConnectionRecordsRequest {
	s.NextToken = &v
	return s
}

type DescribeUserConnectionRecordsResponseBody struct {
	NextToken         *string                                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId         *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConnectionRecords []*DescribeUserConnectionRecordsResponseBodyConnectionRecords `json:"ConnectionRecords,omitempty" xml:"ConnectionRecords,omitempty" type:"Repeated"`
}

func (s DescribeUserConnectionRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConnectionRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectionRecordsResponseBody) SetNextToken(v string) *DescribeUserConnectionRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBody) SetRequestId(v string) *DescribeUserConnectionRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBody) SetConnectionRecords(v []*DescribeUserConnectionRecordsResponseBodyConnectionRecords) *DescribeUserConnectionRecordsResponseBody {
	s.ConnectionRecords = v
	return s
}

type DescribeUserConnectionRecordsResponseBodyConnectionRecords struct {
	ConnectionRecordId *string `json:"ConnectionRecordId,omitempty" xml:"ConnectionRecordId,omitempty"`
	ConnectStartTime   *string `json:"ConnectStartTime,omitempty" xml:"ConnectStartTime,omitempty"`
	DesktopName        *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	ConnectDuration    *string `json:"ConnectDuration,omitempty" xml:"ConnectDuration,omitempty"`
	ConnectEndTime     *string `json:"ConnectEndTime,omitempty" xml:"ConnectEndTime,omitempty"`
	DesktopId          *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
}

func (s DescribeUserConnectionRecordsResponseBodyConnectionRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConnectionRecordsResponseBodyConnectionRecords) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetConnectionRecordId(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.ConnectionRecordId = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetConnectStartTime(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.ConnectStartTime = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetDesktopName(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.DesktopName = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetConnectDuration(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.ConnectDuration = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetConnectEndTime(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.ConnectEndTime = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetDesktopId(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.DesktopId = &v
	return s
}

type DescribeUserConnectionRecordsResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserConnectionRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserConnectionRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserConnectionRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectionRecordsResponse) SetHeaders(v map[string]*string) *DescribeUserConnectionRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserConnectionRecordsResponse) SetBody(v *DescribeUserConnectionRecordsResponseBody) *DescribeUserConnectionRecordsResponse {
	s.Body = v
	return s
}

type DescribeUsersInGroupRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeUsersInGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersInGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupRequest) SetRegionId(v string) *DescribeUsersInGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetDesktopGroupId(v string) *DescribeUsersInGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetMaxResults(v int32) *DescribeUsersInGroupRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetNextToken(v string) *DescribeUsersInGroupRequest {
	s.NextToken = &v
	return s
}

type DescribeUsersInGroupResponseBody struct {
	NextToken        *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId        *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsersCount       *int32                                      `json:"UsersCount,omitempty" xml:"UsersCount,omitempty"`
	OnlineUsersCount *int32                                      `json:"OnlineUsersCount,omitempty" xml:"OnlineUsersCount,omitempty"`
	EndUsers         []*DescribeUsersInGroupResponseBodyEndUsers `json:"EndUsers,omitempty" xml:"EndUsers,omitempty" type:"Repeated"`
}

func (s DescribeUsersInGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupResponseBody) SetNextToken(v string) *DescribeUsersInGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeUsersInGroupResponseBody) SetRequestId(v string) *DescribeUsersInGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsersInGroupResponseBody) SetUsersCount(v int32) *DescribeUsersInGroupResponseBody {
	s.UsersCount = &v
	return s
}

func (s *DescribeUsersInGroupResponseBody) SetOnlineUsersCount(v int32) *DescribeUsersInGroupResponseBody {
	s.OnlineUsersCount = &v
	return s
}

func (s *DescribeUsersInGroupResponseBody) SetEndUsers(v []*DescribeUsersInGroupResponseBodyEndUsers) *DescribeUsersInGroupResponseBody {
	s.EndUsers = v
	return s
}

type DescribeUsersInGroupResponseBodyEndUsers struct {
	EndUserId        *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EndUserEmail     *string `json:"EndUserEmail,omitempty" xml:"EndUserEmail,omitempty"`
	DesktopName      *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	DesktopId        *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	EndUserType      *string `json:"EndUserType,omitempty" xml:"EndUserType,omitempty"`
	EndUserPhone     *string `json:"EndUserPhone,omitempty" xml:"EndUserPhone,omitempty"`
	EndUserName      *string `json:"EndUserName,omitempty" xml:"EndUserName,omitempty"`
}

func (s DescribeUsersInGroupResponseBodyEndUsers) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersInGroupResponseBodyEndUsers) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetEndUserId(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.EndUserId = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetEndUserEmail(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.EndUserEmail = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetDesktopName(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.DesktopName = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetConnectionStatus(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetDesktopId(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.DesktopId = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetEndUserType(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.EndUserType = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetEndUserPhone(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.EndUserPhone = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetEndUserName(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.EndUserName = &v
	return s
}

type DescribeUsersInGroupResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUsersInGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUsersInGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersInGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupResponse) SetHeaders(v map[string]*string) *DescribeUsersInGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsersInGroupResponse) SetBody(v *DescribeUsersInGroupResponseBody) *DescribeUsersInGroupResponse {
	s.Body = v
	return s
}

type DescribeVirtualMFADevicesRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults   *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	DirectoryId  *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OfficeSiteId *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	EndUserId    []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
}

func (s DescribeVirtualMFADevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualMFADevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesRequest) SetRegionId(v string) *DescribeVirtualMFADevicesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetMaxResults(v int32) *DescribeVirtualMFADevicesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetNextToken(v string) *DescribeVirtualMFADevicesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetDirectoryId(v string) *DescribeVirtualMFADevicesRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetOfficeSiteId(v string) *DescribeVirtualMFADevicesRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetEndUserId(v []*string) *DescribeVirtualMFADevicesRequest {
	s.EndUserId = v
	return s
}

type DescribeVirtualMFADevicesResponseBody struct {
	NextToken         *string                                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VirtualMFADevices []*DescribeVirtualMFADevicesResponseBodyVirtualMFADevices `json:"VirtualMFADevices,omitempty" xml:"VirtualMFADevices,omitempty" type:"Repeated"`
}

func (s DescribeVirtualMFADevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualMFADevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesResponseBody) SetNextToken(v string) *DescribeVirtualMFADevicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBody) SetRequestId(v string) *DescribeVirtualMFADevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBody) SetVirtualMFADevices(v []*DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) *DescribeVirtualMFADevicesResponseBody {
	s.VirtualMFADevices = v
	return s
}

type DescribeVirtualMFADevicesResponseBodyVirtualMFADevices struct {
	SerialNumber     *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	GmtUnlock        *string `json:"GmtUnlock,omitempty" xml:"GmtUnlock,omitempty"`
	EndUserId        *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ConsecutiveFails *int32  `json:"ConsecutiveFails,omitempty" xml:"ConsecutiveFails,omitempty"`
	OfficeSiteId     *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	Status           *string `json:"status,omitempty" xml:"status,omitempty"`
	DirectoryId      *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	GmtEnabled       *string `json:"GmtEnabled,omitempty" xml:"GmtEnabled,omitempty"`
}

func (s DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetSerialNumber(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.SerialNumber = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetGmtUnlock(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.GmtUnlock = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetEndUserId(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.EndUserId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetConsecutiveFails(v int32) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.ConsecutiveFails = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetOfficeSiteId(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetStatus(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.Status = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetDirectoryId(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.DirectoryId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetGmtEnabled(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.GmtEnabled = &v
	return s
}

type DescribeVirtualMFADevicesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVirtualMFADevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVirtualMFADevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualMFADevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesResponse) SetHeaders(v map[string]*string) *DescribeVirtualMFADevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVirtualMFADevicesResponse) SetBody(v *DescribeVirtualMFADevicesResponseBody) *DescribeVirtualMFADevicesResponse {
	s.Body = v
	return s
}

type DescribeVulDetailsRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
}

func (s DescribeVulDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsRequest) SetRegionId(v string) *DescribeVulDetailsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetLang(v string) *DescribeVulDetailsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetType(v string) *DescribeVulDetailsRequest {
	s.Type = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetName(v string) *DescribeVulDetailsRequest {
	s.Name = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetAliasName(v string) *DescribeVulDetailsRequest {
	s.AliasName = &v
	return s
}

type DescribeVulDetailsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Cves      []*DescribeVulDetailsResponseBodyCves `json:"Cves,omitempty" xml:"Cves,omitempty" type:"Repeated"`
}

func (s DescribeVulDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsResponseBody) SetRequestId(v string) *DescribeVulDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulDetailsResponseBody) SetCves(v []*DescribeVulDetailsResponseBodyCves) *DescribeVulDetailsResponseBody {
	s.Cves = v
	return s
}

type DescribeVulDetailsResponseBodyCves struct {
	CveId     *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	Summary   *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
	CvssScore *string `json:"CvssScore,omitempty" xml:"CvssScore,omitempty"`
}

func (s DescribeVulDetailsResponseBodyCves) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulDetailsResponseBodyCves) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsResponseBodyCves) SetCveId(v string) *DescribeVulDetailsResponseBodyCves {
	s.CveId = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetSummary(v string) *DescribeVulDetailsResponseBodyCves {
	s.Summary = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetTitle(v string) *DescribeVulDetailsResponseBodyCves {
	s.Title = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetCvssScore(v string) *DescribeVulDetailsResponseBodyCves {
	s.CvssScore = &v
	return s
}

type DescribeVulDetailsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVulDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVulDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsResponse) SetHeaders(v map[string]*string) *DescribeVulDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulDetailsResponse) SetBody(v *DescribeVulDetailsResponseBody) *DescribeVulDetailsResponse {
	s.Body = v
	return s
}

type DescribeVulListRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	Necessity    *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	AliasName    *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Dealed       *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeVulListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulListRequest) SetRegionId(v string) *DescribeVulListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVulListRequest) SetLang(v string) *DescribeVulListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVulListRequest) SetType(v string) *DescribeVulListRequest {
	s.Type = &v
	return s
}

func (s *DescribeVulListRequest) SetOfficeSiteId(v string) *DescribeVulListRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeVulListRequest) SetNecessity(v string) *DescribeVulListRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeVulListRequest) SetAliasName(v string) *DescribeVulListRequest {
	s.AliasName = &v
	return s
}

func (s *DescribeVulListRequest) SetDealed(v string) *DescribeVulListRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeVulListRequest) SetCurrentPage(v int32) *DescribeVulListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulListRequest) SetPageSize(v int32) *DescribeVulListRequest {
	s.PageSize = &v
	return s
}

type DescribeVulListResponseBody struct {
	CurrentPage *int32                                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VulRecords  []*DescribeVulListResponseBodyVulRecords `json:"VulRecords,omitempty" xml:"VulRecords,omitempty" type:"Repeated"`
}

func (s DescribeVulListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBody) SetCurrentPage(v int32) *DescribeVulListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulListResponseBody) SetRequestId(v string) *DescribeVulListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulListResponseBody) SetPageSize(v int32) *DescribeVulListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVulListResponseBody) SetTotalCount(v int32) *DescribeVulListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulListResponseBody) SetVulRecords(v []*DescribeVulListResponseBodyVulRecords) *DescribeVulListResponseBody {
	s.VulRecords = v
	return s
}

type DescribeVulListResponseBodyVulRecords struct {
	Status            *int32                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type              *string                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
	ModifyTs          *int64                                                  `json:"ModifyTs,omitempty" xml:"ModifyTs,omitempty"`
	DesktopName       *string                                                 `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	ResultCode        *string                                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	Tag               *string                                                 `json:"Tag,omitempty" xml:"Tag,omitempty"`
	DesktopId         *string                                                 `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	Related           *string                                                 `json:"Related,omitempty" xml:"Related,omitempty"`
	LastTs            *int64                                                  `json:"LastTs,omitempty" xml:"LastTs,omitempty"`
	FirstTs           *int64                                                  `json:"FirstTs,omitempty" xml:"FirstTs,omitempty"`
	Necessity         *string                                                 `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	RepairTs          *int64                                                  `json:"RepairTs,omitempty" xml:"RepairTs,omitempty"`
	Online            *bool                                                   `json:"Online,omitempty" xml:"Online,omitempty"`
	ResultMessage     *string                                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	OsVersion         *string                                                 `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	AliasName         *string                                                 `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Name              *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	ExtendContentJson *DescribeVulListResponseBodyVulRecordsExtendContentJson `json:"ExtendContentJson,omitempty" xml:"ExtendContentJson,omitempty" type:"Struct"`
}

func (s DescribeVulListResponseBodyVulRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponseBodyVulRecords) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBodyVulRecords) SetStatus(v int32) *DescribeVulListResponseBodyVulRecords {
	s.Status = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetType(v string) *DescribeVulListResponseBodyVulRecords {
	s.Type = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetModifyTs(v int64) *DescribeVulListResponseBodyVulRecords {
	s.ModifyTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetDesktopName(v string) *DescribeVulListResponseBodyVulRecords {
	s.DesktopName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetResultCode(v string) *DescribeVulListResponseBodyVulRecords {
	s.ResultCode = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetTag(v string) *DescribeVulListResponseBodyVulRecords {
	s.Tag = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetDesktopId(v string) *DescribeVulListResponseBodyVulRecords {
	s.DesktopId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetRelated(v string) *DescribeVulListResponseBodyVulRecords {
	s.Related = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetLastTs(v int64) *DescribeVulListResponseBodyVulRecords {
	s.LastTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetFirstTs(v int64) *DescribeVulListResponseBodyVulRecords {
	s.FirstTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetNecessity(v string) *DescribeVulListResponseBodyVulRecords {
	s.Necessity = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetRepairTs(v int64) *DescribeVulListResponseBodyVulRecords {
	s.RepairTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetOnline(v bool) *DescribeVulListResponseBodyVulRecords {
	s.Online = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetResultMessage(v string) *DescribeVulListResponseBodyVulRecords {
	s.ResultMessage = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetOsVersion(v string) *DescribeVulListResponseBodyVulRecords {
	s.OsVersion = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetAliasName(v string) *DescribeVulListResponseBodyVulRecords {
	s.AliasName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetName(v string) *DescribeVulListResponseBodyVulRecords {
	s.Name = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetExtendContentJson(v *DescribeVulListResponseBodyVulRecordsExtendContentJson) *DescribeVulListResponseBodyVulRecords {
	s.ExtendContentJson = v
	return s
}

type DescribeVulListResponseBodyVulRecordsExtendContentJson struct {
	RpmEntityList []*DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList `json:"RpmEntityList,omitempty" xml:"RpmEntityList,omitempty" type:"Repeated"`
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJson) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJson) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetRpmEntityList(v []*DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.RpmEntityList = v
	return s
}

type DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList struct {
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty"`
	UpdateCmd   *string `json:"UpdateCmd,omitempty" xml:"UpdateCmd,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	FullVersion *string `json:"FullVersion,omitempty" xml:"FullVersion,omitempty"`
	MatchDetail *string `json:"MatchDetail,omitempty" xml:"MatchDetail,omitempty"`
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetPath(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Path = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetUpdateCmd(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.UpdateCmd = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetName(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Name = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetFullVersion(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.FullVersion = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetMatchDetail(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.MatchDetail = &v
	return s
}

type DescribeVulListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVulListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVulListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponse) SetHeaders(v map[string]*string) *DescribeVulListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulListResponse) SetBody(v *DescribeVulListResponseBody) *DescribeVulListResponse {
	s.Body = v
	return s
}

type DescribeVulOverviewRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeVulOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulOverviewRequest) SetRegionId(v string) *DescribeVulOverviewRequest {
	s.RegionId = &v
	return s
}

type DescribeVulOverviewResponseBody struct {
	NntfCount  *int32  `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
	LaterCount *int32  `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AsapCount  *int32  `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
}

func (s DescribeVulOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulOverviewResponseBody) SetNntfCount(v int32) *DescribeVulOverviewResponseBody {
	s.NntfCount = &v
	return s
}

func (s *DescribeVulOverviewResponseBody) SetLaterCount(v int32) *DescribeVulOverviewResponseBody {
	s.LaterCount = &v
	return s
}

func (s *DescribeVulOverviewResponseBody) SetRequestId(v string) *DescribeVulOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulOverviewResponseBody) SetAsapCount(v int32) *DescribeVulOverviewResponseBody {
	s.AsapCount = &v
	return s
}

type DescribeVulOverviewResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVulOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVulOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulOverviewResponse) SetHeaders(v map[string]*string) *DescribeVulOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulOverviewResponse) SetBody(v *DescribeVulOverviewResponseBody) *DescribeVulOverviewResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

type DescribeZonesResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     []*DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v []*DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetZoneId(v string) *DescribeZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

type DescribeZonesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type DetachCenRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
}

func (s DetachCenRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachCenRequest) GoString() string {
	return s.String()
}

func (s *DetachCenRequest) SetRegionId(v string) *DetachCenRequest {
	s.RegionId = &v
	return s
}

func (s *DetachCenRequest) SetOfficeSiteId(v string) *DetachCenRequest {
	s.OfficeSiteId = &v
	return s
}

type DetachCenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachCenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachCenResponseBody) GoString() string {
	return s.String()
}

func (s *DetachCenResponseBody) SetRequestId(v string) *DetachCenResponseBody {
	s.RequestId = &v
	return s
}

type DetachCenResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachCenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachCenResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachCenResponse) GoString() string {
	return s.String()
}

func (s *DetachCenResponse) SetHeaders(v map[string]*string) *DetachCenResponse {
	s.Headers = v
	return s
}

func (s *DetachCenResponse) SetBody(v *DetachCenResponseBody) *DetachCenResponse {
	s.Body = v
	return s
}

type DoCheckResourceRequest struct {
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DoCheckResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DoCheckResourceRequest) GoString() string {
	return s.String()
}

func (s *DoCheckResourceRequest) SetInvoker(v string) *DoCheckResourceRequest {
	s.Invoker = &v
	return s
}

func (s *DoCheckResourceRequest) SetPk(v string) *DoCheckResourceRequest {
	s.Pk = &v
	return s
}

func (s *DoCheckResourceRequest) SetBid(v string) *DoCheckResourceRequest {
	s.Bid = &v
	return s
}

func (s *DoCheckResourceRequest) SetHid(v int64) *DoCheckResourceRequest {
	s.Hid = &v
	return s
}

func (s *DoCheckResourceRequest) SetCountry(v string) *DoCheckResourceRequest {
	s.Country = &v
	return s
}

func (s *DoCheckResourceRequest) SetTaskIdentifier(v string) *DoCheckResourceRequest {
	s.TaskIdentifier = &v
	return s
}

func (s *DoCheckResourceRequest) SetTaskExtraData(v string) *DoCheckResourceRequest {
	s.TaskExtraData = &v
	return s
}

func (s *DoCheckResourceRequest) SetGmtWakeup(v string) *DoCheckResourceRequest {
	s.GmtWakeup = &v
	return s
}

func (s *DoCheckResourceRequest) SetRegionId(v string) *DoCheckResourceRequest {
	s.RegionId = &v
	return s
}

type DoCheckResourceResponseBody struct {
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Prompt         *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	Level          *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
}

func (s DoCheckResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DoCheckResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DoCheckResourceResponseBody) SetGmtWakeup(v string) *DoCheckResourceResponseBody {
	s.GmtWakeup = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetHid(v int64) *DoCheckResourceResponseBody {
	s.Hid = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetMessage(v string) *DoCheckResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetTaskIdentifier(v string) *DoCheckResourceResponseBody {
	s.TaskIdentifier = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetRequestId(v string) *DoCheckResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetSuccess(v bool) *DoCheckResourceResponseBody {
	s.Success = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetUrl(v string) *DoCheckResourceResponseBody {
	s.Url = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetInterrupt(v bool) *DoCheckResourceResponseBody {
	s.Interrupt = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetInvoker(v string) *DoCheckResourceResponseBody {
	s.Invoker = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetTaskExtraData(v string) *DoCheckResourceResponseBody {
	s.TaskExtraData = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetCountry(v string) *DoCheckResourceResponseBody {
	s.Country = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetPrompt(v string) *DoCheckResourceResponseBody {
	s.Prompt = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetLevel(v int64) *DoCheckResourceResponseBody {
	s.Level = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetPk(v string) *DoCheckResourceResponseBody {
	s.Pk = &v
	return s
}

func (s *DoCheckResourceResponseBody) SetBid(v string) *DoCheckResourceResponseBody {
	s.Bid = &v
	return s
}

type DoCheckResourceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DoCheckResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DoCheckResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DoCheckResourceResponse) GoString() string {
	return s.String()
}

func (s *DoCheckResourceResponse) SetHeaders(v map[string]*string) *DoCheckResourceResponse {
	s.Headers = v
	return s
}

func (s *DoCheckResourceResponse) SetBody(v *DoCheckResourceResponseBody) *DoCheckResourceResponse {
	s.Body = v
	return s
}

type DoLogicalDeleteResourceRequest struct {
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DoLogicalDeleteResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DoLogicalDeleteResourceRequest) GoString() string {
	return s.String()
}

func (s *DoLogicalDeleteResourceRequest) SetInvoker(v string) *DoLogicalDeleteResourceRequest {
	s.Invoker = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetPk(v string) *DoLogicalDeleteResourceRequest {
	s.Pk = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetBid(v string) *DoLogicalDeleteResourceRequest {
	s.Bid = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetHid(v int64) *DoLogicalDeleteResourceRequest {
	s.Hid = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetCountry(v string) *DoLogicalDeleteResourceRequest {
	s.Country = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetTaskIdentifier(v string) *DoLogicalDeleteResourceRequest {
	s.TaskIdentifier = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetTaskExtraData(v string) *DoLogicalDeleteResourceRequest {
	s.TaskExtraData = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetGmtWakeup(v string) *DoLogicalDeleteResourceRequest {
	s.GmtWakeup = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetRegionId(v string) *DoLogicalDeleteResourceRequest {
	s.RegionId = &v
	return s
}

type DoLogicalDeleteResourceResponseBody struct {
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
}

func (s DoLogicalDeleteResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DoLogicalDeleteResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DoLogicalDeleteResourceResponseBody) SetGmtWakeup(v string) *DoLogicalDeleteResourceResponseBody {
	s.GmtWakeup = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetHid(v int64) *DoLogicalDeleteResourceResponseBody {
	s.Hid = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetMessage(v string) *DoLogicalDeleteResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetTaskIdentifier(v string) *DoLogicalDeleteResourceResponseBody {
	s.TaskIdentifier = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetRequestId(v string) *DoLogicalDeleteResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetInvoker(v string) *DoLogicalDeleteResourceResponseBody {
	s.Invoker = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetTaskExtraData(v string) *DoLogicalDeleteResourceResponseBody {
	s.TaskExtraData = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetCountry(v string) *DoLogicalDeleteResourceResponseBody {
	s.Country = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetPk(v string) *DoLogicalDeleteResourceResponseBody {
	s.Pk = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetBid(v string) *DoLogicalDeleteResourceResponseBody {
	s.Bid = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetSuccess(v bool) *DoLogicalDeleteResourceResponseBody {
	s.Success = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetInterrupt(v bool) *DoLogicalDeleteResourceResponseBody {
	s.Interrupt = &v
	return s
}

type DoLogicalDeleteResourceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DoLogicalDeleteResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DoLogicalDeleteResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DoLogicalDeleteResourceResponse) GoString() string {
	return s.String()
}

func (s *DoLogicalDeleteResourceResponse) SetHeaders(v map[string]*string) *DoLogicalDeleteResourceResponse {
	s.Headers = v
	return s
}

func (s *DoLogicalDeleteResourceResponse) SetBody(v *DoLogicalDeleteResourceResponseBody) *DoLogicalDeleteResourceResponse {
	s.Body = v
	return s
}

type DoPhysicalDeleteResourceRequest struct {
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DoPhysicalDeleteResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DoPhysicalDeleteResourceRequest) GoString() string {
	return s.String()
}

func (s *DoPhysicalDeleteResourceRequest) SetInvoker(v string) *DoPhysicalDeleteResourceRequest {
	s.Invoker = &v
	return s
}

func (s *DoPhysicalDeleteResourceRequest) SetPk(v string) *DoPhysicalDeleteResourceRequest {
	s.Pk = &v
	return s
}

func (s *DoPhysicalDeleteResourceRequest) SetBid(v string) *DoPhysicalDeleteResourceRequest {
	s.Bid = &v
	return s
}

func (s *DoPhysicalDeleteResourceRequest) SetHid(v int64) *DoPhysicalDeleteResourceRequest {
	s.Hid = &v
	return s
}

func (s *DoPhysicalDeleteResourceRequest) SetCountry(v string) *DoPhysicalDeleteResourceRequest {
	s.Country = &v
	return s
}

func (s *DoPhysicalDeleteResourceRequest) SetTaskIdentifier(v string) *DoPhysicalDeleteResourceRequest {
	s.TaskIdentifier = &v
	return s
}

func (s *DoPhysicalDeleteResourceRequest) SetTaskExtraData(v string) *DoPhysicalDeleteResourceRequest {
	s.TaskExtraData = &v
	return s
}

func (s *DoPhysicalDeleteResourceRequest) SetGmtWakeup(v string) *DoPhysicalDeleteResourceRequest {
	s.GmtWakeup = &v
	return s
}

func (s *DoPhysicalDeleteResourceRequest) SetRegionId(v string) *DoPhysicalDeleteResourceRequest {
	s.RegionId = &v
	return s
}

type DoPhysicalDeleteResourceResponseBody struct {
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
}

func (s DoPhysicalDeleteResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DoPhysicalDeleteResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DoPhysicalDeleteResourceResponseBody) SetGmtWakeup(v string) *DoPhysicalDeleteResourceResponseBody {
	s.GmtWakeup = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponseBody) SetHid(v int64) *DoPhysicalDeleteResourceResponseBody {
	s.Hid = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponseBody) SetMessage(v string) *DoPhysicalDeleteResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponseBody) SetTaskIdentifier(v string) *DoPhysicalDeleteResourceResponseBody {
	s.TaskIdentifier = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponseBody) SetRequestId(v string) *DoPhysicalDeleteResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponseBody) SetInvoker(v string) *DoPhysicalDeleteResourceResponseBody {
	s.Invoker = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponseBody) SetTaskExtraData(v string) *DoPhysicalDeleteResourceResponseBody {
	s.TaskExtraData = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponseBody) SetCountry(v string) *DoPhysicalDeleteResourceResponseBody {
	s.Country = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponseBody) SetPk(v string) *DoPhysicalDeleteResourceResponseBody {
	s.Pk = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponseBody) SetBid(v string) *DoPhysicalDeleteResourceResponseBody {
	s.Bid = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponseBody) SetSuccess(v bool) *DoPhysicalDeleteResourceResponseBody {
	s.Success = &v
	return s
}

func (s *DoPhysicalDeleteResourceResponseBody) SetInterrupt(v bool) *DoPhysicalDeleteResourceResponseBody {
	s.Interrupt = &v
	return s
}

type DoPhysicalDeleteResourceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DoPhysicalDeleteResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DoPhysicalDeleteResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DoPhysicalDeleteResourceResponse) GoString() string {
	return s.String()
}

func (s *DoPhysicalDeleteResourceResponse) SetHeaders(v map[string]*string) *DoPhysicalDeleteResourceResponse {
	s.Headers = v
	return s
}

func (s *DoPhysicalDeleteResourceResponse) SetBody(v *DoPhysicalDeleteResourceResponseBody) *DoPhysicalDeleteResourceResponse {
	s.Body = v
	return s
}

type GetConnectionTicketRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserName             *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	DesktopId            *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
}

func (s GetConnectionTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketRequest) SetOwnerId(v int64) *GetConnectionTicketRequest {
	s.OwnerId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetResourceOwnerAccount(v string) *GetConnectionTicketRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetConnectionTicketRequest) SetResourceOwnerId(v int64) *GetConnectionTicketRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetRegionId(v string) *GetConnectionTicketRequest {
	s.RegionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetInstanceId(v string) *GetConnectionTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetUserName(v string) *GetConnectionTicketRequest {
	s.UserName = &v
	return s
}

func (s *GetConnectionTicketRequest) SetPassword(v string) *GetConnectionTicketRequest {
	s.Password = &v
	return s
}

func (s *GetConnectionTicketRequest) SetTaskId(v string) *GetConnectionTicketRequest {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetDesktopId(v string) *GetConnectionTicketRequest {
	s.DesktopId = &v
	return s
}

type GetConnectionTicketResponseBody struct {
	Ticket     *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetConnectionTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBody) SetTicket(v string) *GetConnectionTicketResponseBody {
	s.Ticket = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskId(v string) *GetConnectionTicketResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetRequestId(v string) *GetConnectionTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskStatus(v string) *GetConnectionTicketResponseBody {
	s.TaskStatus = &v
	return s
}

type GetConnectionTicketResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetConnectionTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConnectionTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponse) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponse) SetHeaders(v map[string]*string) *GetConnectionTicketResponse {
	s.Headers = v
	return s
}

func (s *GetConnectionTicketResponse) SetBody(v *GetConnectionTicketResponseBody) *GetConnectionTicketResponse {
	s.Body = v
	return s
}

type GetDesktopGroupDetailRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
}

func (s GetDesktopGroupDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopGroupDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailRequest) SetRegionId(v string) *GetDesktopGroupDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetDesktopGroupDetailRequest) SetDesktopGroupId(v string) *GetDesktopGroupDetailRequest {
	s.DesktopGroupId = &v
	return s
}

type GetDesktopGroupDetailResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Desktops  []*GetDesktopGroupDetailResponseBodyDesktops `json:"Desktops,omitempty" xml:"Desktops,omitempty" type:"Repeated"`
}

func (s GetDesktopGroupDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopGroupDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailResponseBody) SetRequestId(v string) *GetDesktopGroupDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBody) SetDesktops(v []*GetDesktopGroupDetailResponseBodyDesktops) *GetDesktopGroupDetailResponseBody {
	s.Desktops = v
	return s
}

type GetDesktopGroupDetailResponseBodyDesktops struct {
	CreationTime       *string  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	PayType            *string  `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PolicyGroupName    *string  `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	Creator            *string  `json:"Creator,omitempty" xml:"Creator,omitempty"`
	MaxDesktopsCount   *int32   `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	AllowAutoSetup     *int32   `json:"AllowAutoSetup,omitempty" xml:"AllowAutoSetup,omitempty"`
	ResType            *int32   `json:"ResType,omitempty" xml:"ResType,omitempty"`
	SystemDiskSize     *int32   `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	PolicyGroupId      *string  `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	OwnBundleId        *string  `json:"OwnBundleId,omitempty" xml:"OwnBundleId,omitempty"`
	GpuCount           *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	AllowBufferCount   *int32   `json:"AllowBufferCount,omitempty" xml:"AllowBufferCount,omitempty"`
	Memory             *int64   `json:"Memory,omitempty" xml:"Memory,omitempty"`
	GpuSpec            *string  `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	DirectoryId        *string  `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OwnBundleName      *string  `json:"OwnBundleName,omitempty" xml:"OwnBundleName,omitempty"`
	DataDiskCategory   *string  `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	DesktopGroupName   *string  `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	SystemDiskCategory *string  `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	OfficeSiteId       *string  `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	KeepDuration       *int64   `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	MinDesktopsCount   *int32   `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	DataDiskSize       *string  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DesktopGroupId     *string  `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	OfficeSiteName     *string  `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	DirectoryType      *string  `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	Cpu                *int32   `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	ExpiredTime        *string  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Comments           *string  `json:"Comments,omitempty" xml:"Comments,omitempty"`
	OfficeSiteType     *string  `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
}

func (s GetDesktopGroupDetailResponseBodyDesktops) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopGroupDetailResponseBodyDesktops) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetCreationTime(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.CreationTime = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetPayType(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.PayType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetPolicyGroupName(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.PolicyGroupName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetCreator(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.Creator = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetMaxDesktopsCount(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.MaxDesktopsCount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetAllowAutoSetup(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.AllowAutoSetup = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetResType(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.ResType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetSystemDiskSize(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.SystemDiskSize = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetPolicyGroupId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.PolicyGroupId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOwnBundleId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OwnBundleId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetGpuCount(v float32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.GpuCount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetAllowBufferCount(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.AllowBufferCount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetMemory(v int64) *GetDesktopGroupDetailResponseBodyDesktops {
	s.Memory = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetGpuSpec(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.GpuSpec = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDirectoryId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DirectoryId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOwnBundleName(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OwnBundleName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDataDiskCategory(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DataDiskCategory = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDesktopGroupName(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DesktopGroupName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetSystemDiskCategory(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.SystemDiskCategory = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOfficeSiteId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OfficeSiteId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetKeepDuration(v int64) *GetDesktopGroupDetailResponseBodyDesktops {
	s.KeepDuration = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetMinDesktopsCount(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.MinDesktopsCount = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDataDiskSize(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DataDiskSize = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDesktopGroupId(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DesktopGroupId = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOfficeSiteName(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OfficeSiteName = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetDirectoryType(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.DirectoryType = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetCpu(v int32) *GetDesktopGroupDetailResponseBodyDesktops {
	s.Cpu = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetExpiredTime(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.ExpiredTime = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetComments(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.Comments = &v
	return s
}

func (s *GetDesktopGroupDetailResponseBodyDesktops) SetOfficeSiteType(v string) *GetDesktopGroupDetailResponseBodyDesktops {
	s.OfficeSiteType = &v
	return s
}

type GetDesktopGroupDetailResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDesktopGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDesktopGroupDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopGroupDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailResponse) SetHeaders(v map[string]*string) *GetDesktopGroupDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDesktopGroupDetailResponse) SetBody(v *GetDesktopGroupDetailResponseBody) *GetDesktopGroupDetailResponse {
	s.Body = v
	return s
}

type GetDesktopUsersRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
}

func (s GetDesktopUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopUsersRequest) GoString() string {
	return s.String()
}

func (s *GetDesktopUsersRequest) SetRegionId(v string) *GetDesktopUsersRequest {
	s.RegionId = &v
	return s
}

func (s *GetDesktopUsersRequest) SetDesktopId(v string) *GetDesktopUsersRequest {
	s.DesktopId = &v
	return s
}

type GetDesktopUsersResponseBody struct {
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
}

func (s GetDesktopUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopUsersResponseBody) GoString() string {
	return s.String()
}

func (s *GetDesktopUsersResponseBody) SetRequestId(v string) *GetDesktopUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDesktopUsersResponseBody) SetEndUserIds(v []*string) *GetDesktopUsersResponseBody {
	s.EndUserIds = v
	return s
}

type GetDesktopUsersResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDesktopUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDesktopUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDesktopUsersResponse) GoString() string {
	return s.String()
}

func (s *GetDesktopUsersResponse) SetHeaders(v map[string]*string) *GetDesktopUsersResponse {
	s.Headers = v
	return s
}

func (s *GetDesktopUsersResponse) SetBody(v *GetDesktopUsersResponseBody) *GetDesktopUsersResponse {
	s.Body = v
	return s
}

type GetDirectorySsoStatusRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetDirectorySsoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDirectorySsoStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDirectorySsoStatusRequest) SetRegionId(v string) *GetDirectorySsoStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetDirectorySsoStatusRequest) SetDirectoryId(v string) *GetDirectorySsoStatusRequest {
	s.DirectoryId = &v
	return s
}

type GetDirectorySsoStatusResponseBody struct {
	SsoStatus *bool   `json:"SsoStatus,omitempty" xml:"SsoStatus,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDirectorySsoStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDirectorySsoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetDirectorySsoStatusResponseBody) SetSsoStatus(v bool) *GetDirectorySsoStatusResponseBody {
	s.SsoStatus = &v
	return s
}

func (s *GetDirectorySsoStatusResponseBody) SetRequestId(v string) *GetDirectorySsoStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetDirectorySsoStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDirectorySsoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDirectorySsoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDirectorySsoStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDirectorySsoStatusResponse) SetHeaders(v map[string]*string) *GetDirectorySsoStatusResponse {
	s.Headers = v
	return s
}

func (s *GetDirectorySsoStatusResponse) SetBody(v *GetDirectorySsoStatusResponseBody) *GetDirectorySsoStatusResponse {
	s.Body = v
	return s
}

type GetOfficeSiteSsoStatusRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
}

func (s GetOfficeSiteSsoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeSiteSsoStatusRequest) GoString() string {
	return s.String()
}

func (s *GetOfficeSiteSsoStatusRequest) SetRegionId(v string) *GetOfficeSiteSsoStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetOfficeSiteSsoStatusRequest) SetOfficeSiteId(v string) *GetOfficeSiteSsoStatusRequest {
	s.OfficeSiteId = &v
	return s
}

type GetOfficeSiteSsoStatusResponseBody struct {
	SsoStatus *bool   `json:"SsoStatus,omitempty" xml:"SsoStatus,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOfficeSiteSsoStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeSiteSsoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfficeSiteSsoStatusResponseBody) SetSsoStatus(v bool) *GetOfficeSiteSsoStatusResponseBody {
	s.SsoStatus = &v
	return s
}

func (s *GetOfficeSiteSsoStatusResponseBody) SetRequestId(v string) *GetOfficeSiteSsoStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetOfficeSiteSsoStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOfficeSiteSsoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOfficeSiteSsoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeSiteSsoStatusResponse) GoString() string {
	return s.String()
}

func (s *GetOfficeSiteSsoStatusResponse) SetHeaders(v map[string]*string) *GetOfficeSiteSsoStatusResponse {
	s.Headers = v
	return s
}

func (s *GetOfficeSiteSsoStatusResponse) SetBody(v *GetOfficeSiteSsoStatusResponseBody) *GetOfficeSiteSsoStatusResponse {
	s.Body = v
	return s
}

type GetSpMetadataRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DirectoryId  *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
}

func (s GetSpMetadataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpMetadataRequest) GoString() string {
	return s.String()
}

func (s *GetSpMetadataRequest) SetRegionId(v string) *GetSpMetadataRequest {
	s.RegionId = &v
	return s
}

func (s *GetSpMetadataRequest) SetDirectoryId(v string) *GetSpMetadataRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetSpMetadataRequest) SetOfficeSiteId(v string) *GetSpMetadataRequest {
	s.OfficeSiteId = &v
	return s
}

type GetSpMetadataResponseBody struct {
	SpMetadata *string `json:"SpMetadata,omitempty" xml:"SpMetadata,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSpMetadataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSpMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpMetadataResponseBody) SetSpMetadata(v string) *GetSpMetadataResponseBody {
	s.SpMetadata = &v
	return s
}

func (s *GetSpMetadataResponseBody) SetRequestId(v string) *GetSpMetadataResponseBody {
	s.RequestId = &v
	return s
}

type GetSpMetadataResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSpMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSpMetadataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpMetadataResponse) GoString() string {
	return s.String()
}

func (s *GetSpMetadataResponse) SetHeaders(v map[string]*string) *GetSpMetadataResponse {
	s.Headers = v
	return s
}

func (s *GetSpMetadataResponse) SetBody(v *GetSpMetadataResponseBody) *GetSpMetadataResponse {
	s.Body = v
	return s
}

type HandleSecurityEventsRequest struct {
	RegionId        *string                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OperationCode   *string                                     `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	OperationParams *string                                     `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	SecurityEvent   []*HandleSecurityEventsRequestSecurityEvent `json:"SecurityEvent,omitempty" xml:"SecurityEvent,omitempty" type:"Repeated"`
}

func (s HandleSecurityEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s HandleSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsRequest) SetRegionId(v string) *HandleSecurityEventsRequest {
	s.RegionId = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetOperationCode(v string) *HandleSecurityEventsRequest {
	s.OperationCode = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetOperationParams(v string) *HandleSecurityEventsRequest {
	s.OperationParams = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetSecurityEvent(v []*HandleSecurityEventsRequestSecurityEvent) *HandleSecurityEventsRequest {
	s.SecurityEvent = v
	return s
}

type HandleSecurityEventsRequestSecurityEvent struct {
	SecurityEventId *string `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	DesktopId       *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
}

func (s HandleSecurityEventsRequestSecurityEvent) String() string {
	return tea.Prettify(s)
}

func (s HandleSecurityEventsRequestSecurityEvent) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsRequestSecurityEvent) SetSecurityEventId(v string) *HandleSecurityEventsRequestSecurityEvent {
	s.SecurityEventId = &v
	return s
}

func (s *HandleSecurityEventsRequestSecurityEvent) SetDesktopId(v string) *HandleSecurityEventsRequestSecurityEvent {
	s.DesktopId = &v
	return s
}

type HandleSecurityEventsResponseBody struct {
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s HandleSecurityEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HandleSecurityEventsResponseBody) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsResponseBody) SetTaskId(v int64) *HandleSecurityEventsResponseBody {
	s.TaskId = &v
	return s
}

func (s *HandleSecurityEventsResponseBody) SetRequestId(v string) *HandleSecurityEventsResponseBody {
	s.RequestId = &v
	return s
}

type HandleSecurityEventsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HandleSecurityEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HandleSecurityEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s HandleSecurityEventsResponse) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsResponse) SetHeaders(v map[string]*string) *HandleSecurityEventsResponse {
	s.Headers = v
	return s
}

func (s *HandleSecurityEventsResponse) SetBody(v *HandleSecurityEventsResponseBody) *HandleSecurityEventsResponse {
	s.Body = v
	return s
}

type ListDirectoryUsersRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Filter      *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListDirectoryUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoryUsersRequest) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersRequest) SetRegionId(v string) *ListDirectoryUsersRequest {
	s.RegionId = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetFilter(v string) *ListDirectoryUsersRequest {
	s.Filter = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetDirectoryId(v string) *ListDirectoryUsersRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetNextToken(v string) *ListDirectoryUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListDirectoryUsersRequest) SetMaxResults(v int32) *ListDirectoryUsersRequest {
	s.MaxResults = &v
	return s
}

type ListDirectoryUsersResponseBody struct {
	NextToken *string                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     []*ListDirectoryUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListDirectoryUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoryUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersResponseBody) SetNextToken(v string) *ListDirectoryUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDirectoryUsersResponseBody) SetRequestId(v string) *ListDirectoryUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDirectoryUsersResponseBody) SetUsers(v []*ListDirectoryUsersResponseBodyUsers) *ListDirectoryUsersResponseBody {
	s.Users = v
	return s
}

type ListDirectoryUsersResponseBodyUsers struct {
	EndUser *string `json:"EndUser,omitempty" xml:"EndUser,omitempty"`
}

func (s ListDirectoryUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoryUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersResponseBodyUsers) SetEndUser(v string) *ListDirectoryUsersResponseBodyUsers {
	s.EndUser = &v
	return s
}

type ListDirectoryUsersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDirectoryUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDirectoryUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDirectoryUsersResponse) GoString() string {
	return s.String()
}

func (s *ListDirectoryUsersResponse) SetHeaders(v map[string]*string) *ListDirectoryUsersResponse {
	s.Headers = v
	return s
}

func (s *ListDirectoryUsersResponse) SetBody(v *ListDirectoryUsersResponseBody) *ListDirectoryUsersResponse {
	s.Body = v
	return s
}

type ListOfficeSiteOverviewRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ForceRefresh *bool     `json:"ForceRefresh,omitempty" xml:"ForceRefresh,omitempty"`
	MaxResults   *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSiteId []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
}

func (s ListOfficeSiteOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteOverviewRequest) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteOverviewRequest) SetRegionId(v string) *ListOfficeSiteOverviewRequest {
	s.RegionId = &v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetForceRefresh(v bool) *ListOfficeSiteOverviewRequest {
	s.ForceRefresh = &v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetMaxResults(v int32) *ListOfficeSiteOverviewRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetNextToken(v string) *ListOfficeSiteOverviewRequest {
	s.NextToken = &v
	return s
}

func (s *ListOfficeSiteOverviewRequest) SetOfficeSiteId(v []*string) *ListOfficeSiteOverviewRequest {
	s.OfficeSiteId = v
	return s
}

type ListOfficeSiteOverviewResponseBody struct {
	NextToken                 *string                                                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                 *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OfficeSiteOverviewResults []*ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults `json:"OfficeSiteOverviewResults,omitempty" xml:"OfficeSiteOverviewResults,omitempty" type:"Repeated"`
}

func (s ListOfficeSiteOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteOverviewResponseBody) SetNextToken(v string) *ListOfficeSiteOverviewResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBody) SetRequestId(v string) *ListOfficeSiteOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBody) SetOfficeSiteOverviewResults(v []*ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) *ListOfficeSiteOverviewResponseBody {
	s.OfficeSiteOverviewResults = v
	return s
}

type ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults struct {
	OfficeSiteStatus    *string `json:"OfficeSiteStatus,omitempty" xml:"OfficeSiteStatus,omitempty"`
	TotalEdsCount       *int32  `json:"TotalEdsCount,omitempty" xml:"TotalEdsCount,omitempty"`
	WillExpiredEdsCount *int32  `json:"WillExpiredEdsCount,omitempty" xml:"WillExpiredEdsCount,omitempty"`
	OfficeSiteId        *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RunningEdsCount     *int32  `json:"RunningEdsCount,omitempty" xml:"RunningEdsCount,omitempty"`
	OfficeSiteName      *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	HasExpiredEdsCount  *int32  `json:"HasExpiredEdsCount,omitempty" xml:"HasExpiredEdsCount,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetOfficeSiteStatus(v string) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.OfficeSiteStatus = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetTotalEdsCount(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.TotalEdsCount = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetWillExpiredEdsCount(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.WillExpiredEdsCount = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetOfficeSiteId(v string) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.OfficeSiteId = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetRunningEdsCount(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.RunningEdsCount = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetOfficeSiteName(v string) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.OfficeSiteName = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetHasExpiredEdsCount(v int32) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.HasExpiredEdsCount = &v
	return s
}

func (s *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults) SetRegionId(v string) *ListOfficeSiteOverviewResponseBodyOfficeSiteOverviewResults {
	s.RegionId = &v
	return s
}

type ListOfficeSiteOverviewResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOfficeSiteOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOfficeSiteOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteOverviewResponse) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteOverviewResponse) SetHeaders(v map[string]*string) *ListOfficeSiteOverviewResponse {
	s.Headers = v
	return s
}

func (s *ListOfficeSiteOverviewResponse) SetBody(v *ListOfficeSiteOverviewResponseBody) *ListOfficeSiteOverviewResponse {
	s.Body = v
	return s
}

type ListOfficeSiteUsersRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Filter       *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListOfficeSiteUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteUsersRequest) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteUsersRequest) SetRegionId(v string) *ListOfficeSiteUsersRequest {
	s.RegionId = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetFilter(v string) *ListOfficeSiteUsersRequest {
	s.Filter = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetOfficeSiteId(v string) *ListOfficeSiteUsersRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetNextToken(v string) *ListOfficeSiteUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListOfficeSiteUsersRequest) SetMaxResults(v int32) *ListOfficeSiteUsersRequest {
	s.MaxResults = &v
	return s
}

type ListOfficeSiteUsersResponseBody struct {
	NextToken *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     []*ListOfficeSiteUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListOfficeSiteUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteUsersResponseBody) SetNextToken(v string) *ListOfficeSiteUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOfficeSiteUsersResponseBody) SetRequestId(v string) *ListOfficeSiteUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOfficeSiteUsersResponseBody) SetUsers(v []*ListOfficeSiteUsersResponseBodyUsers) *ListOfficeSiteUsersResponseBody {
	s.Users = v
	return s
}

type ListOfficeSiteUsersResponseBodyUsers struct {
	EndUser *string `json:"EndUser,omitempty" xml:"EndUser,omitempty"`
}

func (s ListOfficeSiteUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteUsersResponseBodyUsers) SetEndUser(v string) *ListOfficeSiteUsersResponseBodyUsers {
	s.EndUser = &v
	return s
}

type ListOfficeSiteUsersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOfficeSiteUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOfficeSiteUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOfficeSiteUsersResponse) GoString() string {
	return s.String()
}

func (s *ListOfficeSiteUsersResponse) SetHeaders(v map[string]*string) *ListOfficeSiteUsersResponse {
	s.Headers = v
	return s
}

func (s *ListOfficeSiteUsersResponse) SetBody(v *ListOfficeSiteUsersResponseBody) *ListOfficeSiteUsersResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	MaxResults   *int32                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetMaxResults(v int32) *ListTagResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type LockVirtualMFADeviceRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s LockVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s LockVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *LockVirtualMFADeviceRequest) SetRegionId(v string) *LockVirtualMFADeviceRequest {
	s.RegionId = &v
	return s
}

func (s *LockVirtualMFADeviceRequest) SetSerialNumber(v string) *LockVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

type LockVirtualMFADeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LockVirtualMFADeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LockVirtualMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *LockVirtualMFADeviceResponseBody) SetRequestId(v string) *LockVirtualMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

type LockVirtualMFADeviceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LockVirtualMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LockVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s LockVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *LockVirtualMFADeviceResponse) SetHeaders(v map[string]*string) *LockVirtualMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *LockVirtualMFADeviceResponse) SetBody(v *LockVirtualMFADeviceResponseBody) *LockVirtualMFADeviceResponse {
	s.Body = v
	return s
}

type ModifyADConnectorDirectoryRequest struct {
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DirectoryId         *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DomainName          *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainUserName      *string   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	DomainPassword      *string   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	DirectoryName       *string   `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	SubDomainName       *string   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	MfaEnabled          *bool     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	DnsAddress          []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	SubDomainDnsAddress []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
}

func (s ModifyADConnectorDirectoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorDirectoryRequest) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorDirectoryRequest) SetRegionId(v string) *ModifyADConnectorDirectoryRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDirectoryId(v string) *ModifyADConnectorDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDomainName(v string) *ModifyADConnectorDirectoryRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDomainUserName(v string) *ModifyADConnectorDirectoryRequest {
	s.DomainUserName = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDomainPassword(v string) *ModifyADConnectorDirectoryRequest {
	s.DomainPassword = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDirectoryName(v string) *ModifyADConnectorDirectoryRequest {
	s.DirectoryName = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetSubDomainName(v string) *ModifyADConnectorDirectoryRequest {
	s.SubDomainName = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetMfaEnabled(v bool) *ModifyADConnectorDirectoryRequest {
	s.MfaEnabled = &v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetDnsAddress(v []*string) *ModifyADConnectorDirectoryRequest {
	s.DnsAddress = v
	return s
}

func (s *ModifyADConnectorDirectoryRequest) SetSubDomainDnsAddress(v []*string) *ModifyADConnectorDirectoryRequest {
	s.SubDomainDnsAddress = v
	return s
}

type ModifyADConnectorDirectoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyADConnectorDirectoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorDirectoryResponseBody) SetRequestId(v string) *ModifyADConnectorDirectoryResponseBody {
	s.RequestId = &v
	return s
}

type ModifyADConnectorDirectoryResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyADConnectorDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyADConnectorDirectoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorDirectoryResponse) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorDirectoryResponse) SetHeaders(v map[string]*string) *ModifyADConnectorDirectoryResponse {
	s.Headers = v
	return s
}

func (s *ModifyADConnectorDirectoryResponse) SetBody(v *ModifyADConnectorDirectoryResponseBody) *ModifyADConnectorDirectoryResponse {
	s.Body = v
	return s
}

type ModifyADConnectorOfficeSiteRequest struct {
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OfficeSiteId        *string   `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	DomainName          *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainUserName      *string   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	DomainPassword      *string   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	OfficeSiteName      *string   `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	SubDomainName       *string   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
	MfaEnabled          *bool     `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	DnsAddress          []*string `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	SubDomainDnsAddress []*string `json:"SubDomainDnsAddress,omitempty" xml:"SubDomainDnsAddress,omitempty" type:"Repeated"`
}

func (s ModifyADConnectorOfficeSiteRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorOfficeSiteRequest) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorOfficeSiteRequest) SetRegionId(v string) *ModifyADConnectorOfficeSiteRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetOfficeSiteId(v string) *ModifyADConnectorOfficeSiteRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetDomainName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetDomainUserName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.DomainUserName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetDomainPassword(v string) *ModifyADConnectorOfficeSiteRequest {
	s.DomainPassword = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetOfficeSiteName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.OfficeSiteName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetSubDomainName(v string) *ModifyADConnectorOfficeSiteRequest {
	s.SubDomainName = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetMfaEnabled(v bool) *ModifyADConnectorOfficeSiteRequest {
	s.MfaEnabled = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetDnsAddress(v []*string) *ModifyADConnectorOfficeSiteRequest {
	s.DnsAddress = v
	return s
}

func (s *ModifyADConnectorOfficeSiteRequest) SetSubDomainDnsAddress(v []*string) *ModifyADConnectorOfficeSiteRequest {
	s.SubDomainDnsAddress = v
	return s
}

type ModifyADConnectorOfficeSiteResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyADConnectorOfficeSiteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorOfficeSiteResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorOfficeSiteResponseBody) SetRequestId(v string) *ModifyADConnectorOfficeSiteResponseBody {
	s.RequestId = &v
	return s
}

type ModifyADConnectorOfficeSiteResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyADConnectorOfficeSiteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyADConnectorOfficeSiteResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyADConnectorOfficeSiteResponse) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorOfficeSiteResponse) SetHeaders(v map[string]*string) *ModifyADConnectorOfficeSiteResponse {
	s.Headers = v
	return s
}

func (s *ModifyADConnectorOfficeSiteResponse) SetBody(v *ModifyADConnectorOfficeSiteResponseBody) *ModifyADConnectorOfficeSiteResponse {
	s.Body = v
	return s
}

type ModifyBundleRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BundleId    *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	BundleName  *string `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyBundleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBundleRequest) GoString() string {
	return s.String()
}

func (s *ModifyBundleRequest) SetRegionId(v string) *ModifyBundleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBundleRequest) SetBundleId(v string) *ModifyBundleRequest {
	s.BundleId = &v
	return s
}

func (s *ModifyBundleRequest) SetImageId(v string) *ModifyBundleRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyBundleRequest) SetBundleName(v string) *ModifyBundleRequest {
	s.BundleName = &v
	return s
}

func (s *ModifyBundleRequest) SetDescription(v string) *ModifyBundleRequest {
	s.Description = &v
	return s
}

type ModifyBundleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBundleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBundleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBundleResponseBody) SetRequestId(v string) *ModifyBundleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBundleResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBundleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBundleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBundleResponse) GoString() string {
	return s.String()
}

func (s *ModifyBundleResponse) SetHeaders(v map[string]*string) *ModifyBundleResponse {
	s.Headers = v
	return s
}

func (s *ModifyBundleResponse) SetBody(v *ModifyBundleResponseBody) *ModifyBundleResponse {
	s.Body = v
	return s
}

type ModifyDesktopGroupRequest struct {
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopGroupId   *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	OwnBundleId      *string `json:"OwnBundleId,omitempty" xml:"OwnBundleId,omitempty"`
	PolicyGroupId    *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	ScaleStrategyId  *string `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty"`
	KeepDuration     *int64  `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	Comments         *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	MinDesktopsCount *int32  `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	MaxDesktopsCount *int32  `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	AllowAutoSetup   *int32  `json:"AllowAutoSetup,omitempty" xml:"AllowAutoSetup,omitempty"`
	AllowBufferCount *int32  `json:"AllowBufferCount,omitempty" xml:"AllowBufferCount,omitempty"`
}

func (s ModifyDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopGroupRequest) SetRegionId(v string) *ModifyDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetDesktopGroupId(v string) *ModifyDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetOwnBundleId(v string) *ModifyDesktopGroupRequest {
	s.OwnBundleId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetPolicyGroupId(v string) *ModifyDesktopGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetDesktopGroupName(v string) *ModifyDesktopGroupRequest {
	s.DesktopGroupName = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetScaleStrategyId(v string) *ModifyDesktopGroupRequest {
	s.ScaleStrategyId = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetKeepDuration(v int64) *ModifyDesktopGroupRequest {
	s.KeepDuration = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetComments(v string) *ModifyDesktopGroupRequest {
	s.Comments = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetMinDesktopsCount(v int32) *ModifyDesktopGroupRequest {
	s.MinDesktopsCount = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetMaxDesktopsCount(v int32) *ModifyDesktopGroupRequest {
	s.MaxDesktopsCount = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetAllowAutoSetup(v int32) *ModifyDesktopGroupRequest {
	s.AllowAutoSetup = &v
	return s
}

func (s *ModifyDesktopGroupRequest) SetAllowBufferCount(v int32) *ModifyDesktopGroupRequest {
	s.AllowBufferCount = &v
	return s
}

type ModifyDesktopGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopGroupResponseBody) SetRequestId(v string) *ModifyDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDesktopGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopGroupResponse) SetHeaders(v map[string]*string) *ModifyDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopGroupResponse) SetBody(v *ModifyDesktopGroupResponseBody) *ModifyDesktopGroupResponse {
	s.Body = v
	return s
}

type ModifyDesktopNameRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	NewDesktopName *string `json:"NewDesktopName,omitempty" xml:"NewDesktopName,omitempty"`
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
}

func (s ModifyDesktopNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopNameRequest) SetRegionId(v string) *ModifyDesktopNameRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopNameRequest) SetNewDesktopName(v string) *ModifyDesktopNameRequest {
	s.NewDesktopName = &v
	return s
}

func (s *ModifyDesktopNameRequest) SetDesktopId(v string) *ModifyDesktopNameRequest {
	s.DesktopId = &v
	return s
}

type ModifyDesktopNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopNameResponseBody) SetRequestId(v string) *ModifyDesktopNameResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDesktopNameResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDesktopNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDesktopNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopNameResponse) SetHeaders(v map[string]*string) *ModifyDesktopNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopNameResponse) SetBody(v *ModifyDesktopNameResponseBody) *ModifyDesktopNameResponse {
	s.Body = v
	return s
}

type ModifyDesktopPolicysRequest struct {
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Clipboard   *string   `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	LocalDrive  *string   `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UsbRedirect *string   `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	Watermark   *string   `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	DesktopId   []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s ModifyDesktopPolicysRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopPolicysRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopPolicysRequest) SetRegionId(v string) *ModifyDesktopPolicysRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopPolicysRequest) SetClipboard(v string) *ModifyDesktopPolicysRequest {
	s.Clipboard = &v
	return s
}

func (s *ModifyDesktopPolicysRequest) SetLocalDrive(v string) *ModifyDesktopPolicysRequest {
	s.LocalDrive = &v
	return s
}

func (s *ModifyDesktopPolicysRequest) SetUsbRedirect(v string) *ModifyDesktopPolicysRequest {
	s.UsbRedirect = &v
	return s
}

func (s *ModifyDesktopPolicysRequest) SetWatermark(v string) *ModifyDesktopPolicysRequest {
	s.Watermark = &v
	return s
}

func (s *ModifyDesktopPolicysRequest) SetDesktopId(v []*string) *ModifyDesktopPolicysRequest {
	s.DesktopId = v
	return s
}

type ModifyDesktopPolicysResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*ModifyDesktopPolicysResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s ModifyDesktopPolicysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopPolicysResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopPolicysResponseBody) SetRequestId(v string) *ModifyDesktopPolicysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDesktopPolicysResponseBody) SetResults(v []*ModifyDesktopPolicysResponseBodyResults) *ModifyDesktopPolicysResponseBody {
	s.Results = v
	return s
}

type ModifyDesktopPolicysResponseBodyResults struct {
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
}

func (s ModifyDesktopPolicysResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopPolicysResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ModifyDesktopPolicysResponseBodyResults) SetSuccess(v string) *ModifyDesktopPolicysResponseBodyResults {
	s.Success = &v
	return s
}

func (s *ModifyDesktopPolicysResponseBodyResults) SetCode(v string) *ModifyDesktopPolicysResponseBodyResults {
	s.Code = &v
	return s
}

func (s *ModifyDesktopPolicysResponseBodyResults) SetMessage(v string) *ModifyDesktopPolicysResponseBodyResults {
	s.Message = &v
	return s
}

func (s *ModifyDesktopPolicysResponseBodyResults) SetDesktopId(v string) *ModifyDesktopPolicysResponseBodyResults {
	s.DesktopId = &v
	return s
}

type ModifyDesktopPolicysResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDesktopPolicysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDesktopPolicysResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopPolicysResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopPolicysResponse) SetHeaders(v map[string]*string) *ModifyDesktopPolicysResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopPolicysResponse) SetBody(v *ModifyDesktopPolicysResponseBody) *ModifyDesktopPolicysResponse {
	s.Body = v
	return s
}

type ModifyDesktopSpecRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopId       *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopType     *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	RootDiskSizeGib *int32  `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	UserDiskSizeGib *int32  `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty"`
	AutoPay         *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
}

func (s ModifyDesktopSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopSpecRequest) SetRegionId(v string) *ModifyDesktopSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetDesktopId(v string) *ModifyDesktopSpecRequest {
	s.DesktopId = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetDesktopType(v string) *ModifyDesktopSpecRequest {
	s.DesktopType = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetRootDiskSizeGib(v int32) *ModifyDesktopSpecRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetUserDiskSizeGib(v int32) *ModifyDesktopSpecRequest {
	s.UserDiskSizeGib = &v
	return s
}

func (s *ModifyDesktopSpecRequest) SetAutoPay(v bool) *ModifyDesktopSpecRequest {
	s.AutoPay = &v
	return s
}

type ModifyDesktopSpecResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopSpecResponseBody) SetOrderId(v string) *ModifyDesktopSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDesktopSpecResponseBody) SetRequestId(v string) *ModifyDesktopSpecResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDesktopSpecResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDesktopSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDesktopSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopSpecResponse) SetHeaders(v map[string]*string) *ModifyDesktopSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopSpecResponse) SetBody(v *ModifyDesktopSpecResponseBody) *ModifyDesktopSpecResponse {
	s.Body = v
	return s
}

type ModifyDesktopsPolicyGroupRequest struct {
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PolicyGroupId *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	DesktopId     []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s ModifyDesktopsPolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupRequest) SetRegionId(v string) *ModifyDesktopsPolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupRequest) SetPolicyGroupId(v string) *ModifyDesktopsPolicyGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupRequest) SetDesktopId(v []*string) *ModifyDesktopsPolicyGroupRequest {
	s.DesktopId = v
	return s
}

type ModifyDesktopsPolicyGroupResponseBody struct {
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ModifyResults []*ModifyDesktopsPolicyGroupResponseBodyModifyResults `json:"ModifyResults,omitempty" xml:"ModifyResults,omitempty" type:"Repeated"`
}

func (s ModifyDesktopsPolicyGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupResponseBody) SetRequestId(v string) *ModifyDesktopsPolicyGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponseBody) SetModifyResults(v []*ModifyDesktopsPolicyGroupResponseBodyModifyResults) *ModifyDesktopsPolicyGroupResponseBody {
	s.ModifyResults = v
	return s
}

type ModifyDesktopsPolicyGroupResponseBodyModifyResults struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
}

func (s ModifyDesktopsPolicyGroupResponseBodyModifyResults) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupResponseBodyModifyResults) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupResponseBodyModifyResults) SetCode(v string) *ModifyDesktopsPolicyGroupResponseBodyModifyResults {
	s.Code = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponseBodyModifyResults) SetMessage(v string) *ModifyDesktopsPolicyGroupResponseBodyModifyResults {
	s.Message = &v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponseBodyModifyResults) SetDesktopId(v string) *ModifyDesktopsPolicyGroupResponseBodyModifyResults {
	s.DesktopId = &v
	return s
}

type ModifyDesktopsPolicyGroupResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDesktopsPolicyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDesktopsPolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDesktopsPolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDesktopsPolicyGroupResponse) SetHeaders(v map[string]*string) *ModifyDesktopsPolicyGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDesktopsPolicyGroupResponse) SetBody(v *ModifyDesktopsPolicyGroupResponseBody) *ModifyDesktopsPolicyGroupResponse {
	s.Body = v
	return s
}

type ModifyEntitlementRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopId *string   `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
}

func (s ModifyEntitlementRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyEntitlementRequest) GoString() string {
	return s.String()
}

func (s *ModifyEntitlementRequest) SetRegionId(v string) *ModifyEntitlementRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyEntitlementRequest) SetDesktopId(v string) *ModifyEntitlementRequest {
	s.DesktopId = &v
	return s
}

func (s *ModifyEntitlementRequest) SetEndUserId(v []*string) *ModifyEntitlementRequest {
	s.EndUserId = v
	return s
}

type ModifyEntitlementResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEntitlementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyEntitlementResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEntitlementResponseBody) SetRequestId(v string) *ModifyEntitlementResponseBody {
	s.RequestId = &v
	return s
}

type ModifyEntitlementResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyEntitlementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyEntitlementResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyEntitlementResponse) GoString() string {
	return s.String()
}

func (s *ModifyEntitlementResponse) SetHeaders(v map[string]*string) *ModifyEntitlementResponse {
	s.Headers = v
	return s
}

func (s *ModifyEntitlementResponse) SetBody(v *ModifyEntitlementResponseBody) *ModifyEntitlementResponse {
	s.Body = v
	return s
}

type ModifyImageAttributeRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyImageAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeRequest) SetRegionId(v string) *ModifyImageAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetImageId(v string) *ModifyImageAttributeRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetName(v string) *ModifyImageAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetDescription(v string) *ModifyImageAttributeRequest {
	s.Description = &v
	return s
}

type ModifyImageAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyImageAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeResponseBody) SetRequestId(v string) *ModifyImageAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyImageAttributeResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyImageAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyImageAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeResponse) SetHeaders(v map[string]*string) *ModifyImageAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyImageAttributeResponse) SetBody(v *ModifyImageAttributeResponseBody) *ModifyImageAttributeResponse {
	s.Body = v
	return s
}

type ModifyNASDefaultMountTargetRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FileSystemId      *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
}

func (s ModifyNASDefaultMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNASDefaultMountTargetRequest) GoString() string {
	return s.String()
}

func (s *ModifyNASDefaultMountTargetRequest) SetRegionId(v string) *ModifyNASDefaultMountTargetRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNASDefaultMountTargetRequest) SetFileSystemId(v string) *ModifyNASDefaultMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyNASDefaultMountTargetRequest) SetMountTargetDomain(v string) *ModifyNASDefaultMountTargetRequest {
	s.MountTargetDomain = &v
	return s
}

type ModifyNASDefaultMountTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNASDefaultMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNASDefaultMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNASDefaultMountTargetResponseBody) SetRequestId(v string) *ModifyNASDefaultMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNASDefaultMountTargetResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyNASDefaultMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyNASDefaultMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNASDefaultMountTargetResponse) GoString() string {
	return s.String()
}

func (s *ModifyNASDefaultMountTargetResponse) SetHeaders(v map[string]*string) *ModifyNASDefaultMountTargetResponse {
	s.Headers = v
	return s
}

func (s *ModifyNASDefaultMountTargetResponse) SetBody(v *ModifyNASDefaultMountTargetResponseBody) *ModifyNASDefaultMountTargetResponse {
	s.Body = v
	return s
}

type ModifyNetworkPackageRequest struct {
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	Bandwidth        *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
}

func (s ModifyNetworkPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageRequest) SetRegionId(v string) *ModifyNetworkPackageRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNetworkPackageRequest) SetNetworkPackageId(v string) *ModifyNetworkPackageRequest {
	s.NetworkPackageId = &v
	return s
}

func (s *ModifyNetworkPackageRequest) SetBandwidth(v int32) *ModifyNetworkPackageRequest {
	s.Bandwidth = &v
	return s
}

type ModifyNetworkPackageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNetworkPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageResponseBody) SetRequestId(v string) *ModifyNetworkPackageResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNetworkPackageResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyNetworkPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyNetworkPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageResponse) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageResponse) SetHeaders(v map[string]*string) *ModifyNetworkPackageResponse {
	s.Headers = v
	return s
}

func (s *ModifyNetworkPackageResponse) SetBody(v *ModifyNetworkPackageResponseBody) *ModifyNetworkPackageResponse {
	s.Body = v
	return s
}

type ModifyNetworkPackageEnabledRequest struct {
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	Enabled          *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s ModifyNetworkPackageEnabledRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageEnabledRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageEnabledRequest) SetRegionId(v string) *ModifyNetworkPackageEnabledRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNetworkPackageEnabledRequest) SetNetworkPackageId(v string) *ModifyNetworkPackageEnabledRequest {
	s.NetworkPackageId = &v
	return s
}

func (s *ModifyNetworkPackageEnabledRequest) SetEnabled(v bool) *ModifyNetworkPackageEnabledRequest {
	s.Enabled = &v
	return s
}

type ModifyNetworkPackageEnabledResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNetworkPackageEnabledResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageEnabledResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageEnabledResponseBody) SetRequestId(v string) *ModifyNetworkPackageEnabledResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNetworkPackageEnabledResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyNetworkPackageEnabledResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyNetworkPackageEnabledResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNetworkPackageEnabledResponse) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageEnabledResponse) SetHeaders(v map[string]*string) *ModifyNetworkPackageEnabledResponse {
	s.Headers = v
	return s
}

func (s *ModifyNetworkPackageEnabledResponse) SetBody(v *ModifyNetworkPackageEnabledResponseBody) *ModifyNetworkPackageEnabledResponse {
	s.Body = v
	return s
}

type ModifyOfficeSiteAttributeRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OfficeSiteId      *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	OfficeSiteName    *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
}

func (s ModifyOfficeSiteAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteAttributeRequest) SetRegionId(v string) *ModifyOfficeSiteAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetOfficeSiteId(v string) *ModifyOfficeSiteAttributeRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetDesktopAccessType(v string) *ModifyOfficeSiteAttributeRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *ModifyOfficeSiteAttributeRequest) SetOfficeSiteName(v string) *ModifyOfficeSiteAttributeRequest {
	s.OfficeSiteName = &v
	return s
}

type ModifyOfficeSiteAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOfficeSiteAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteAttributeResponseBody) SetRequestId(v string) *ModifyOfficeSiteAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyOfficeSiteAttributeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyOfficeSiteAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyOfficeSiteAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteAttributeResponse) SetHeaders(v map[string]*string) *ModifyOfficeSiteAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyOfficeSiteAttributeResponse) SetBody(v *ModifyOfficeSiteAttributeResponseBody) *ModifyOfficeSiteAttributeResponse {
	s.Body = v
	return s
}

type ModifyOfficeSiteCrossDesktopAccessRequest struct {
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OfficeSiteId             *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	EnableCrossDesktopAccess *bool   `json:"EnableCrossDesktopAccess,omitempty" xml:"EnableCrossDesktopAccess,omitempty"`
}

func (s ModifyOfficeSiteCrossDesktopAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteCrossDesktopAccessRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteCrossDesktopAccessRequest) SetRegionId(v string) *ModifyOfficeSiteCrossDesktopAccessRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOfficeSiteCrossDesktopAccessRequest) SetOfficeSiteId(v string) *ModifyOfficeSiteCrossDesktopAccessRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyOfficeSiteCrossDesktopAccessRequest) SetEnableCrossDesktopAccess(v bool) *ModifyOfficeSiteCrossDesktopAccessRequest {
	s.EnableCrossDesktopAccess = &v
	return s
}

type ModifyOfficeSiteCrossDesktopAccessResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOfficeSiteCrossDesktopAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteCrossDesktopAccessResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponseBody) SetRequestId(v string) *ModifyOfficeSiteCrossDesktopAccessResponseBody {
	s.RequestId = &v
	return s
}

type ModifyOfficeSiteCrossDesktopAccessResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyOfficeSiteCrossDesktopAccessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyOfficeSiteCrossDesktopAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteCrossDesktopAccessResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponse) SetHeaders(v map[string]*string) *ModifyOfficeSiteCrossDesktopAccessResponse {
	s.Headers = v
	return s
}

func (s *ModifyOfficeSiteCrossDesktopAccessResponse) SetBody(v *ModifyOfficeSiteCrossDesktopAccessResponseBody) *ModifyOfficeSiteCrossDesktopAccessResponse {
	s.Body = v
	return s
}

type ModifyOfficeSiteMfaEnabledRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	MfaEnabled   *bool   `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
}

func (s ModifyOfficeSiteMfaEnabledRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteMfaEnabledRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteMfaEnabledRequest) SetRegionId(v string) *ModifyOfficeSiteMfaEnabledRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOfficeSiteMfaEnabledRequest) SetOfficeSiteId(v string) *ModifyOfficeSiteMfaEnabledRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyOfficeSiteMfaEnabledRequest) SetMfaEnabled(v bool) *ModifyOfficeSiteMfaEnabledRequest {
	s.MfaEnabled = &v
	return s
}

type ModifyOfficeSiteMfaEnabledResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOfficeSiteMfaEnabledResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteMfaEnabledResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteMfaEnabledResponseBody) SetRequestId(v string) *ModifyOfficeSiteMfaEnabledResponseBody {
	s.RequestId = &v
	return s
}

type ModifyOfficeSiteMfaEnabledResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyOfficeSiteMfaEnabledResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyOfficeSiteMfaEnabledResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOfficeSiteMfaEnabledResponse) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteMfaEnabledResponse) SetHeaders(v map[string]*string) *ModifyOfficeSiteMfaEnabledResponse {
	s.Headers = v
	return s
}

func (s *ModifyOfficeSiteMfaEnabledResponse) SetBody(v *ModifyOfficeSiteMfaEnabledResponseBody) *ModifyOfficeSiteMfaEnabledResponse {
	s.Body = v
	return s
}

type ModifyOperateVulRequest struct {
	RegionId    *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type        *string                           `json:"Type,omitempty" xml:"Type,omitempty"`
	OperateType *string                           `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	Reason      *string                           `json:"Reason,omitempty" xml:"Reason,omitempty"`
	VulInfo     []*ModifyOperateVulRequestVulInfo `json:"VulInfo,omitempty" xml:"VulInfo,omitempty" type:"Repeated"`
}

func (s ModifyOperateVulRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOperateVulRequest) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulRequest) SetRegionId(v string) *ModifyOperateVulRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOperateVulRequest) SetType(v string) *ModifyOperateVulRequest {
	s.Type = &v
	return s
}

func (s *ModifyOperateVulRequest) SetOperateType(v string) *ModifyOperateVulRequest {
	s.OperateType = &v
	return s
}

func (s *ModifyOperateVulRequest) SetReason(v string) *ModifyOperateVulRequest {
	s.Reason = &v
	return s
}

func (s *ModifyOperateVulRequest) SetVulInfo(v []*ModifyOperateVulRequestVulInfo) *ModifyOperateVulRequest {
	s.VulInfo = v
	return s
}

type ModifyOperateVulRequestVulInfo struct {
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	Tag       *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyOperateVulRequestVulInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifyOperateVulRequestVulInfo) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulRequestVulInfo) SetDesktopId(v string) *ModifyOperateVulRequestVulInfo {
	s.DesktopId = &v
	return s
}

func (s *ModifyOperateVulRequestVulInfo) SetTag(v string) *ModifyOperateVulRequestVulInfo {
	s.Tag = &v
	return s
}

func (s *ModifyOperateVulRequestVulInfo) SetName(v string) *ModifyOperateVulRequestVulInfo {
	s.Name = &v
	return s
}

type ModifyOperateVulResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOperateVulResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyOperateVulResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulResponseBody) SetRequestId(v string) *ModifyOperateVulResponseBody {
	s.RequestId = &v
	return s
}

type ModifyOperateVulResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyOperateVulResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyOperateVulResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOperateVulResponse) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulResponse) SetHeaders(v map[string]*string) *ModifyOperateVulResponse {
	s.Headers = v
	return s
}

func (s *ModifyOperateVulResponse) SetBody(v *ModifyOperateVulResponseBody) *ModifyOperateVulResponse {
	s.Body = v
	return s
}

type ModifyPolicyGroupRequest struct {
	RegionId                    *string                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PolicyGroupId               *string                                                `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	Name                        *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Clipboard                   *string                                                `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	LocalDrive                  *string                                                `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UsbRedirect                 *string                                                `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	VisualQuality               *string                                                `json:"VisualQuality,omitempty" xml:"VisualQuality,omitempty"`
	Html5Access                 *string                                                `json:"Html5Access,omitempty" xml:"Html5Access,omitempty"`
	Html5FileTransfer           *string                                                `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	Watermark                   *string                                                `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	WatermarkType               *string                                                `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	WatermarkCustomText         *string                                                `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	WatermarkTransparency       *string                                                `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
	PreemptLogin                *string                                                `json:"PreemptLogin,omitempty" xml:"PreemptLogin,omitempty"`
	DomainList                  *string                                                `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	PreemptLoginUser            []*string                                              `json:"PreemptLoginUser,omitempty" xml:"PreemptLoginUser,omitempty" type:"Repeated"`
	AuthorizeSecurityPolicyRule []*ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule `json:"AuthorizeSecurityPolicyRule,omitempty" xml:"AuthorizeSecurityPolicyRule,omitempty" type:"Repeated"`
	RevokeSecurityPolicyRule    []*ModifyPolicyGroupRequestRevokeSecurityPolicyRule    `json:"RevokeSecurityPolicyRule,omitempty" xml:"RevokeSecurityPolicyRule,omitempty" type:"Repeated"`
	AuthorizeAccessPolicyRule   []*ModifyPolicyGroupRequestAuthorizeAccessPolicyRule   `json:"AuthorizeAccessPolicyRule,omitempty" xml:"AuthorizeAccessPolicyRule,omitempty" type:"Repeated"`
	RevokeAccessPolicyRule      []*ModifyPolicyGroupRequestRevokeAccessPolicyRule      `json:"RevokeAccessPolicyRule,omitempty" xml:"RevokeAccessPolicyRule,omitempty" type:"Repeated"`
}

func (s ModifyPolicyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequest) SetRegionId(v string) *ModifyPolicyGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPolicyGroupId(v string) *ModifyPolicyGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetName(v string) *ModifyPolicyGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetClipboard(v string) *ModifyPolicyGroupRequest {
	s.Clipboard = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetLocalDrive(v string) *ModifyPolicyGroupRequest {
	s.LocalDrive = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetUsbRedirect(v string) *ModifyPolicyGroupRequest {
	s.UsbRedirect = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetVisualQuality(v string) *ModifyPolicyGroupRequest {
	s.VisualQuality = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetHtml5Access(v string) *ModifyPolicyGroupRequest {
	s.Html5Access = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetHtml5FileTransfer(v string) *ModifyPolicyGroupRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermark(v string) *ModifyPolicyGroupRequest {
	s.Watermark = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkType(v string) *ModifyPolicyGroupRequest {
	s.WatermarkType = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkCustomText(v string) *ModifyPolicyGroupRequest {
	s.WatermarkCustomText = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetWatermarkTransparency(v string) *ModifyPolicyGroupRequest {
	s.WatermarkTransparency = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPreemptLogin(v string) *ModifyPolicyGroupRequest {
	s.PreemptLogin = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetDomainList(v string) *ModifyPolicyGroupRequest {
	s.DomainList = &v
	return s
}

func (s *ModifyPolicyGroupRequest) SetPreemptLoginUser(v []*string) *ModifyPolicyGroupRequest {
	s.PreemptLoginUser = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetAuthorizeSecurityPolicyRule(v []*ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) *ModifyPolicyGroupRequest {
	s.AuthorizeSecurityPolicyRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRevokeSecurityPolicyRule(v []*ModifyPolicyGroupRequestRevokeSecurityPolicyRule) *ModifyPolicyGroupRequest {
	s.RevokeSecurityPolicyRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetAuthorizeAccessPolicyRule(v []*ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) *ModifyPolicyGroupRequest {
	s.AuthorizeAccessPolicyRule = v
	return s
}

func (s *ModifyPolicyGroupRequest) SetRevokeAccessPolicyRule(v []*ModifyPolicyGroupRequestRevokeAccessPolicyRule) *ModifyPolicyGroupRequest {
	s.RevokeAccessPolicyRule = v
	return s
}

type ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PortRange   *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IpProtocol  *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Priority    *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
}

func (s ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetType(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Type = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetPolicy(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Policy = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetPortRange(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.PortRange = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetDescription(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Description = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetIpProtocol(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.IpProtocol = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetPriority(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.Priority = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule) SetCidrIp(v string) *ModifyPolicyGroupRequestAuthorizeSecurityPolicyRule {
	s.CidrIp = &v
	return s
}

type ModifyPolicyGroupRequestRevokeSecurityPolicyRule struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	PortRange   *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IpProtocol  *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Priority    *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
}

func (s ModifyPolicyGroupRequestRevokeSecurityPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestRevokeSecurityPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetType(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.Type = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetPolicy(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.Policy = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetPortRange(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.PortRange = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetDescription(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.Description = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetIpProtocol(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.IpProtocol = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetPriority(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.Priority = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeSecurityPolicyRule) SetCidrIp(v string) *ModifyPolicyGroupRequestRevokeSecurityPolicyRule {
	s.CidrIp = &v
	return s
}

type ModifyPolicyGroupRequestAuthorizeAccessPolicyRule struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
}

func (s ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) SetDescription(v string) *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule {
	s.Description = &v
	return s
}

func (s *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule) SetCidrIp(v string) *ModifyPolicyGroupRequestAuthorizeAccessPolicyRule {
	s.CidrIp = &v
	return s
}

type ModifyPolicyGroupRequestRevokeAccessPolicyRule struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CidrIp      *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
}

func (s ModifyPolicyGroupRequestRevokeAccessPolicyRule) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupRequestRevokeAccessPolicyRule) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupRequestRevokeAccessPolicyRule) SetDescription(v string) *ModifyPolicyGroupRequestRevokeAccessPolicyRule {
	s.Description = &v
	return s
}

func (s *ModifyPolicyGroupRequestRevokeAccessPolicyRule) SetCidrIp(v string) *ModifyPolicyGroupRequestRevokeAccessPolicyRule {
	s.CidrIp = &v
	return s
}

type ModifyPolicyGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPolicyGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupResponseBody) SetRequestId(v string) *ModifyPolicyGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPolicyGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyPolicyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPolicyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupResponse) SetHeaders(v map[string]*string) *ModifyPolicyGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolicyGroupResponse) SetBody(v *ModifyPolicyGroupResponseBody) *ModifyPolicyGroupResponse {
	s.Body = v
	return s
}

type ModifyScaleStrategyRequest struct {
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScaleStrategyId           *string `json:"ScaleStrategyId,omitempty" xml:"ScaleStrategyId,omitempty"`
	ScaleStrategyName         *string `json:"ScaleStrategyName,omitempty" xml:"ScaleStrategyName,omitempty"`
	ScaleStrategyType         *string `json:"ScaleStrategyType,omitempty" xml:"ScaleStrategyType,omitempty"`
	MinDesktopsCount          *int32  `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	MaxDesktopsCount          *int32  `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	MinAvailableDesktopsCount *int32  `json:"MinAvailableDesktopsCount,omitempty" xml:"MinAvailableDesktopsCount,omitempty"`
	MaxAvailableDesktopsCount *int32  `json:"MaxAvailableDesktopsCount,omitempty" xml:"MaxAvailableDesktopsCount,omitempty"`
	ScaleStep                 *int32  `json:"ScaleStep,omitempty" xml:"ScaleStep,omitempty"`
}

func (s ModifyScaleStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScaleStrategyRequest) GoString() string {
	return s.String()
}

func (s *ModifyScaleStrategyRequest) SetRegionId(v string) *ModifyScaleStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyScaleStrategyRequest) SetScaleStrategyId(v string) *ModifyScaleStrategyRequest {
	s.ScaleStrategyId = &v
	return s
}

func (s *ModifyScaleStrategyRequest) SetScaleStrategyName(v string) *ModifyScaleStrategyRequest {
	s.ScaleStrategyName = &v
	return s
}

func (s *ModifyScaleStrategyRequest) SetScaleStrategyType(v string) *ModifyScaleStrategyRequest {
	s.ScaleStrategyType = &v
	return s
}

func (s *ModifyScaleStrategyRequest) SetMinDesktopsCount(v int32) *ModifyScaleStrategyRequest {
	s.MinDesktopsCount = &v
	return s
}

func (s *ModifyScaleStrategyRequest) SetMaxDesktopsCount(v int32) *ModifyScaleStrategyRequest {
	s.MaxDesktopsCount = &v
	return s
}

func (s *ModifyScaleStrategyRequest) SetMinAvailableDesktopsCount(v int32) *ModifyScaleStrategyRequest {
	s.MinAvailableDesktopsCount = &v
	return s
}

func (s *ModifyScaleStrategyRequest) SetMaxAvailableDesktopsCount(v int32) *ModifyScaleStrategyRequest {
	s.MaxAvailableDesktopsCount = &v
	return s
}

func (s *ModifyScaleStrategyRequest) SetScaleStep(v int32) *ModifyScaleStrategyRequest {
	s.ScaleStep = &v
	return s
}

type ModifyScaleStrategyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyScaleStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyScaleStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScaleStrategyResponseBody) SetRequestId(v string) *ModifyScaleStrategyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyScaleStrategyResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyScaleStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyScaleStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyScaleStrategyResponse) GoString() string {
	return s.String()
}

func (s *ModifyScaleStrategyResponse) SetHeaders(v map[string]*string) *ModifyScaleStrategyResponse {
	s.Headers = v
	return s
}

func (s *ModifyScaleStrategyResponse) SetBody(v *ModifyScaleStrategyResponseBody) *ModifyScaleStrategyResponse {
	s.Body = v
	return s
}

type ModifyUserToDesktopGroupRequest struct {
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopGroupId *string   `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	OldEndUserIds  []*string `json:"OldEndUserIds,omitempty" xml:"OldEndUserIds,omitempty" type:"Repeated"`
	NewEndUserIds  []*string `json:"NewEndUserIds,omitempty" xml:"NewEndUserIds,omitempty" type:"Repeated"`
}

func (s ModifyUserToDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserToDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserToDesktopGroupRequest) SetRegionId(v string) *ModifyUserToDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserToDesktopGroupRequest) SetDesktopGroupId(v string) *ModifyUserToDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *ModifyUserToDesktopGroupRequest) SetOldEndUserIds(v []*string) *ModifyUserToDesktopGroupRequest {
	s.OldEndUserIds = v
	return s
}

func (s *ModifyUserToDesktopGroupRequest) SetNewEndUserIds(v []*string) *ModifyUserToDesktopGroupRequest {
	s.NewEndUserIds = v
	return s
}

type ModifyUserToDesktopGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserToDesktopGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserToDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserToDesktopGroupResponseBody) SetRequestId(v string) *ModifyUserToDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserToDesktopGroupResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUserToDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserToDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserToDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserToDesktopGroupResponse) SetHeaders(v map[string]*string) *ModifyUserToDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserToDesktopGroupResponse) SetBody(v *ModifyUserToDesktopGroupResponseBody) *ModifyUserToDesktopGroupResponse {
	s.Body = v
	return s
}

type OperateVulsRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type         *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	OperateType  *string   `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	Reason       *string   `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Precondition *int32    `json:"Precondition,omitempty" xml:"Precondition,omitempty"`
	VulName      []*string `json:"VulName,omitempty" xml:"VulName,omitempty" type:"Repeated"`
	DesktopId    []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s OperateVulsRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateVulsRequest) GoString() string {
	return s.String()
}

func (s *OperateVulsRequest) SetRegionId(v string) *OperateVulsRequest {
	s.RegionId = &v
	return s
}

func (s *OperateVulsRequest) SetType(v string) *OperateVulsRequest {
	s.Type = &v
	return s
}

func (s *OperateVulsRequest) SetOperateType(v string) *OperateVulsRequest {
	s.OperateType = &v
	return s
}

func (s *OperateVulsRequest) SetReason(v string) *OperateVulsRequest {
	s.Reason = &v
	return s
}

func (s *OperateVulsRequest) SetPrecondition(v int32) *OperateVulsRequest {
	s.Precondition = &v
	return s
}

func (s *OperateVulsRequest) SetVulName(v []*string) *OperateVulsRequest {
	s.VulName = v
	return s
}

func (s *OperateVulsRequest) SetDesktopId(v []*string) *OperateVulsRequest {
	s.DesktopId = v
	return s
}

type OperateVulsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateVulsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateVulsResponseBody) GoString() string {
	return s.String()
}

func (s *OperateVulsResponseBody) SetRequestId(v string) *OperateVulsResponseBody {
	s.RequestId = &v
	return s
}

type OperateVulsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OperateVulsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OperateVulsResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateVulsResponse) GoString() string {
	return s.String()
}

func (s *OperateVulsResponse) SetHeaders(v map[string]*string) *OperateVulsResponse {
	s.Headers = v
	return s
}

func (s *OperateVulsResponse) SetBody(v *OperateVulsResponseBody) *OperateVulsResponse {
	s.Body = v
	return s
}

type PayOrderCallbackRequest struct {
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PayOrderCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s PayOrderCallbackRequest) GoString() string {
	return s.String()
}

func (s *PayOrderCallbackRequest) SetData(v string) *PayOrderCallbackRequest {
	s.Data = &v
	return s
}

type PayOrderCallbackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PayOrderCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PayOrderCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *PayOrderCallbackResponseBody) SetRequestId(v string) *PayOrderCallbackResponseBody {
	s.RequestId = &v
	return s
}

type PayOrderCallbackResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PayOrderCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PayOrderCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s PayOrderCallbackResponse) GoString() string {
	return s.String()
}

func (s *PayOrderCallbackResponse) SetHeaders(v map[string]*string) *PayOrderCallbackResponse {
	s.Headers = v
	return s
}

func (s *PayOrderCallbackResponse) SetBody(v *PayOrderCallbackResponseBody) *PayOrderCallbackResponse {
	s.Body = v
	return s
}

type RebootDesktopsRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s RebootDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RebootDesktopsRequest) SetRegionId(v string) *RebootDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *RebootDesktopsRequest) SetDesktopId(v []*string) *RebootDesktopsRequest {
	s.DesktopId = v
	return s
}

type RebootDesktopsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *RebootDesktopsResponseBody) SetRequestId(v string) *RebootDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type RebootDesktopsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RebootDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebootDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RebootDesktopsResponse) SetHeaders(v map[string]*string) *RebootDesktopsResponse {
	s.Headers = v
	return s
}

func (s *RebootDesktopsResponse) SetBody(v *RebootDesktopsResponseBody) *RebootDesktopsResponse {
	s.Body = v
	return s
}

type RebuildDesktopsRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s RebuildDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s RebuildDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsRequest) SetRegionId(v string) *RebuildDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *RebuildDesktopsRequest) SetDesktopId(v []*string) *RebuildDesktopsRequest {
	s.DesktopId = v
	return s
}

type RebuildDesktopsResponseBody struct {
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RebuildResults []*RebuildDesktopsResponseBodyRebuildResults `json:"RebuildResults,omitempty" xml:"RebuildResults,omitempty" type:"Repeated"`
}

func (s RebuildDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebuildDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsResponseBody) SetRequestId(v string) *RebuildDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebuildDesktopsResponseBody) SetRebuildResults(v []*RebuildDesktopsResponseBodyRebuildResults) *RebuildDesktopsResponseBody {
	s.RebuildResults = v
	return s
}

type RebuildDesktopsResponseBodyRebuildResults struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
}

func (s RebuildDesktopsResponseBodyRebuildResults) String() string {
	return tea.Prettify(s)
}

func (s RebuildDesktopsResponseBodyRebuildResults) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsResponseBodyRebuildResults) SetCode(v string) *RebuildDesktopsResponseBodyRebuildResults {
	s.Code = &v
	return s
}

func (s *RebuildDesktopsResponseBodyRebuildResults) SetMessage(v string) *RebuildDesktopsResponseBodyRebuildResults {
	s.Message = &v
	return s
}

func (s *RebuildDesktopsResponseBodyRebuildResults) SetDesktopId(v string) *RebuildDesktopsResponseBodyRebuildResults {
	s.DesktopId = &v
	return s
}

type RebuildDesktopsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RebuildDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebuildDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s RebuildDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsResponse) SetHeaders(v map[string]*string) *RebuildDesktopsResponse {
	s.Headers = v
	return s
}

func (s *RebuildDesktopsResponse) SetBody(v *RebuildDesktopsResponseBody) *RebuildDesktopsResponse {
	s.Body = v
	return s
}

type RecreateDesktopGroupRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	OwnBundleId    *string `json:"OwnBundleId,omitempty" xml:"OwnBundleId,omitempty"`
}

func (s RecreateDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RecreateDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *RecreateDesktopGroupRequest) SetRegionId(v string) *RecreateDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RecreateDesktopGroupRequest) SetDesktopGroupId(v string) *RecreateDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *RecreateDesktopGroupRequest) SetOwnBundleId(v string) *RecreateDesktopGroupRequest {
	s.OwnBundleId = &v
	return s
}

type RecreateDesktopGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecreateDesktopGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecreateDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RecreateDesktopGroupResponseBody) SetRequestId(v string) *RecreateDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

type RecreateDesktopGroupResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecreateDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecreateDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RecreateDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *RecreateDesktopGroupResponse) SetHeaders(v map[string]*string) *RecreateDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *RecreateDesktopGroupResponse) SetBody(v *RecreateDesktopGroupResponseBody) *RecreateDesktopGroupResponse {
	s.Body = v
	return s
}

type RemoveUserFromDesktopGroupRequest struct {
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopGroupId *string   `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	EndUserIds     []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
}

func (s RemoveUserFromDesktopGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromDesktopGroupRequest) SetRegionId(v string) *RemoveUserFromDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveUserFromDesktopGroupRequest) SetDesktopGroupId(v string) *RemoveUserFromDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *RemoveUserFromDesktopGroupRequest) SetEndUserIds(v []*string) *RemoveUserFromDesktopGroupRequest {
	s.EndUserIds = v
	return s
}

type RemoveUserFromDesktopGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUserFromDesktopGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserFromDesktopGroupResponseBody) SetRequestId(v string) *RemoveUserFromDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

type RemoveUserFromDesktopGroupResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveUserFromDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveUserFromDesktopGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserFromDesktopGroupResponse) SetHeaders(v map[string]*string) *RemoveUserFromDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserFromDesktopGroupResponse) SetBody(v *RemoveUserFromDesktopGroupResponseBody) *RemoveUserFromDesktopGroupResponse {
	s.Body = v
	return s
}

type RenewDesktopsRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Period     *int32    `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoPay    *bool     `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	DesktopId  []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s RenewDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RenewDesktopsRequest) SetRegionId(v string) *RenewDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *RenewDesktopsRequest) SetPeriod(v int32) *RenewDesktopsRequest {
	s.Period = &v
	return s
}

func (s *RenewDesktopsRequest) SetPeriodUnit(v string) *RenewDesktopsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewDesktopsRequest) SetAutoPay(v bool) *RenewDesktopsRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewDesktopsRequest) SetDesktopId(v []*string) *RenewDesktopsRequest {
	s.DesktopId = v
	return s
}

type RenewDesktopsResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *RenewDesktopsResponseBody) SetOrderId(v string) *RenewDesktopsResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewDesktopsResponseBody) SetRequestId(v string) *RenewDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type RenewDesktopsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RenewDesktopsResponse) SetHeaders(v map[string]*string) *RenewDesktopsResponse {
	s.Headers = v
	return s
}

func (s *RenewDesktopsResponse) SetBody(v *RenewDesktopsResponseBody) *RenewDesktopsResponse {
	s.Body = v
	return s
}

type ResetNASDefaultMountTargetRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s ResetNASDefaultMountTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetNASDefaultMountTargetRequest) GoString() string {
	return s.String()
}

func (s *ResetNASDefaultMountTargetRequest) SetRegionId(v string) *ResetNASDefaultMountTargetRequest {
	s.RegionId = &v
	return s
}

func (s *ResetNASDefaultMountTargetRequest) SetFileSystemId(v string) *ResetNASDefaultMountTargetRequest {
	s.FileSystemId = &v
	return s
}

type ResetNASDefaultMountTargetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetNASDefaultMountTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetNASDefaultMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ResetNASDefaultMountTargetResponseBody) SetRequestId(v string) *ResetNASDefaultMountTargetResponseBody {
	s.RequestId = &v
	return s
}

type ResetNASDefaultMountTargetResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetNASDefaultMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetNASDefaultMountTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetNASDefaultMountTargetResponse) GoString() string {
	return s.String()
}

func (s *ResetNASDefaultMountTargetResponse) SetHeaders(v map[string]*string) *ResetNASDefaultMountTargetResponse {
	s.Headers = v
	return s
}

func (s *ResetNASDefaultMountTargetResponse) SetBody(v *ResetNASDefaultMountTargetResponseBody) *ResetNASDefaultMountTargetResponse {
	s.Body = v
	return s
}

type ResetSnapshotRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ResetSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetSnapshotRequest) GoString() string {
	return s.String()
}

func (s *ResetSnapshotRequest) SetRegionId(v string) *ResetSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *ResetSnapshotRequest) SetSnapshotId(v string) *ResetSnapshotRequest {
	s.SnapshotId = &v
	return s
}

type ResetSnapshotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *ResetSnapshotResponseBody) SetRequestId(v string) *ResetSnapshotResponseBody {
	s.RequestId = &v
	return s
}

type ResetSnapshotResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetSnapshotResponse) GoString() string {
	return s.String()
}

func (s *ResetSnapshotResponse) SetHeaders(v map[string]*string) *ResetSnapshotResponse {
	s.Headers = v
	return s
}

func (s *ResetSnapshotResponse) SetBody(v *ResetSnapshotResponseBody) *ResetSnapshotResponse {
	s.Body = v
	return s
}

type RollbackSuspEventQuaraFileRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	QuaraFieldId *int32  `json:"QuaraFieldId,omitempty" xml:"QuaraFieldId,omitempty"`
	DesktopId    *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
}

func (s RollbackSuspEventQuaraFileRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackSuspEventQuaraFileRequest) GoString() string {
	return s.String()
}

func (s *RollbackSuspEventQuaraFileRequest) SetRegionId(v string) *RollbackSuspEventQuaraFileRequest {
	s.RegionId = &v
	return s
}

func (s *RollbackSuspEventQuaraFileRequest) SetQuaraFieldId(v int32) *RollbackSuspEventQuaraFileRequest {
	s.QuaraFieldId = &v
	return s
}

func (s *RollbackSuspEventQuaraFileRequest) SetDesktopId(v string) *RollbackSuspEventQuaraFileRequest {
	s.DesktopId = &v
	return s
}

type RollbackSuspEventQuaraFileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackSuspEventQuaraFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackSuspEventQuaraFileResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackSuspEventQuaraFileResponseBody) SetRequestId(v string) *RollbackSuspEventQuaraFileResponseBody {
	s.RequestId = &v
	return s
}

type RollbackSuspEventQuaraFileResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RollbackSuspEventQuaraFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackSuspEventQuaraFileResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackSuspEventQuaraFileResponse) GoString() string {
	return s.String()
}

func (s *RollbackSuspEventQuaraFileResponse) SetHeaders(v map[string]*string) *RollbackSuspEventQuaraFileResponse {
	s.Headers = v
	return s
}

func (s *RollbackSuspEventQuaraFileResponse) SetBody(v *RollbackSuspEventQuaraFileResponseBody) *RollbackSuspEventQuaraFileResponse {
	s.Body = v
	return s
}

type RunCommandRequest struct {
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type            *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	CommandContent  *string   `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	Timeout         *int64    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	ContentEncoding *string   `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	DesktopId       []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s RunCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCommandRequest) GoString() string {
	return s.String()
}

func (s *RunCommandRequest) SetRegionId(v string) *RunCommandRequest {
	s.RegionId = &v
	return s
}

func (s *RunCommandRequest) SetType(v string) *RunCommandRequest {
	s.Type = &v
	return s
}

func (s *RunCommandRequest) SetCommandContent(v string) *RunCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandRequest) SetTimeout(v int64) *RunCommandRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandRequest) SetContentEncoding(v string) *RunCommandRequest {
	s.ContentEncoding = &v
	return s
}

func (s *RunCommandRequest) SetDesktopId(v []*string) *RunCommandRequest {
	s.DesktopId = v
	return s
}

type RunCommandResponseBody struct {
	InvokeId  *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunCommandResponseBody) SetInvokeId(v string) *RunCommandResponseBody {
	s.InvokeId = &v
	return s
}

func (s *RunCommandResponseBody) SetRequestId(v string) *RunCommandResponseBody {
	s.RequestId = &v
	return s
}

type RunCommandResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponse) GoString() string {
	return s.String()
}

func (s *RunCommandResponse) SetHeaders(v map[string]*string) *RunCommandResponse {
	s.Headers = v
	return s
}

func (s *RunCommandResponse) SetBody(v *RunCommandResponseBody) *RunCommandResponse {
	s.Body = v
	return s
}

type SetDirectorySsoStatusRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EnableSso   *bool   `json:"EnableSso,omitempty" xml:"EnableSso,omitempty"`
}

func (s SetDirectorySsoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDirectorySsoStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDirectorySsoStatusRequest) SetRegionId(v string) *SetDirectorySsoStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetDirectorySsoStatusRequest) SetDirectoryId(v string) *SetDirectorySsoStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetDirectorySsoStatusRequest) SetEnableSso(v bool) *SetDirectorySsoStatusRequest {
	s.EnableSso = &v
	return s
}

type SetDirectorySsoStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDirectorySsoStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDirectorySsoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDirectorySsoStatusResponseBody) SetRequestId(v string) *SetDirectorySsoStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetDirectorySsoStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDirectorySsoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDirectorySsoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDirectorySsoStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDirectorySsoStatusResponse) SetHeaders(v map[string]*string) *SetDirectorySsoStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDirectorySsoStatusResponse) SetBody(v *SetDirectorySsoStatusResponseBody) *SetDirectorySsoStatusResponse {
	s.Body = v
	return s
}

type SetIdpMetadataRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DirectoryId  *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	IdpMetadata  *string `json:"IdpMetadata,omitempty" xml:"IdpMetadata,omitempty"`
}

func (s SetIdpMetadataRequest) String() string {
	return tea.Prettify(s)
}

func (s SetIdpMetadataRequest) GoString() string {
	return s.String()
}

func (s *SetIdpMetadataRequest) SetRegionId(v string) *SetIdpMetadataRequest {
	s.RegionId = &v
	return s
}

func (s *SetIdpMetadataRequest) SetDirectoryId(v string) *SetIdpMetadataRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetIdpMetadataRequest) SetOfficeSiteId(v string) *SetIdpMetadataRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *SetIdpMetadataRequest) SetIdpMetadata(v string) *SetIdpMetadataRequest {
	s.IdpMetadata = &v
	return s
}

type SetIdpMetadataResponseBody struct {
	IdpEntityId *string `json:"IdpEntityId,omitempty" xml:"IdpEntityId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetIdpMetadataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetIdpMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *SetIdpMetadataResponseBody) SetIdpEntityId(v string) *SetIdpMetadataResponseBody {
	s.IdpEntityId = &v
	return s
}

func (s *SetIdpMetadataResponseBody) SetRequestId(v string) *SetIdpMetadataResponseBody {
	s.RequestId = &v
	return s
}

type SetIdpMetadataResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetIdpMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetIdpMetadataResponse) String() string {
	return tea.Prettify(s)
}

func (s SetIdpMetadataResponse) GoString() string {
	return s.String()
}

func (s *SetIdpMetadataResponse) SetHeaders(v map[string]*string) *SetIdpMetadataResponse {
	s.Headers = v
	return s
}

func (s *SetIdpMetadataResponse) SetBody(v *SetIdpMetadataResponseBody) *SetIdpMetadataResponse {
	s.Body = v
	return s
}

type SetOfficeSiteSsoStatusRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	EnableSso    *bool   `json:"EnableSso,omitempty" xml:"EnableSso,omitempty"`
}

func (s SetOfficeSiteSsoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetOfficeSiteSsoStatusRequest) GoString() string {
	return s.String()
}

func (s *SetOfficeSiteSsoStatusRequest) SetRegionId(v string) *SetOfficeSiteSsoStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetOfficeSiteSsoStatusRequest) SetOfficeSiteId(v string) *SetOfficeSiteSsoStatusRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *SetOfficeSiteSsoStatusRequest) SetEnableSso(v bool) *SetOfficeSiteSsoStatusRequest {
	s.EnableSso = &v
	return s
}

type SetOfficeSiteSsoStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetOfficeSiteSsoStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetOfficeSiteSsoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetOfficeSiteSsoStatusResponseBody) SetRequestId(v string) *SetOfficeSiteSsoStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetOfficeSiteSsoStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetOfficeSiteSsoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetOfficeSiteSsoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetOfficeSiteSsoStatusResponse) GoString() string {
	return s.String()
}

func (s *SetOfficeSiteSsoStatusResponse) SetHeaders(v map[string]*string) *SetOfficeSiteSsoStatusResponse {
	s.Headers = v
	return s
}

func (s *SetOfficeSiteSsoStatusResponse) SetBody(v *SetOfficeSiteSsoStatusResponseBody) *SetOfficeSiteSsoStatusResponse {
	s.Body = v
	return s
}

type StartDesktopsRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s StartDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StartDesktopsRequest) SetRegionId(v string) *StartDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *StartDesktopsRequest) SetDesktopId(v []*string) *StartDesktopsRequest {
	s.DesktopId = v
	return s
}

type StartDesktopsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *StartDesktopsResponseBody) SetRequestId(v string) *StartDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type StartDesktopsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsResponse) GoString() string {
	return s.String()
}

func (s *StartDesktopsResponse) SetHeaders(v map[string]*string) *StartDesktopsResponse {
	s.Headers = v
	return s
}

func (s *StartDesktopsResponse) SetBody(v *StartDesktopsResponseBody) *StartDesktopsResponse {
	s.Body = v
	return s
}

type StartVirusScanTaskRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OfficeSiteId []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
	DesktopId    []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s StartVirusScanTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StartVirusScanTaskRequest) GoString() string {
	return s.String()
}

func (s *StartVirusScanTaskRequest) SetRegionId(v string) *StartVirusScanTaskRequest {
	s.RegionId = &v
	return s
}

func (s *StartVirusScanTaskRequest) SetOfficeSiteId(v []*string) *StartVirusScanTaskRequest {
	s.OfficeSiteId = v
	return s
}

func (s *StartVirusScanTaskRequest) SetDesktopId(v []*string) *StartVirusScanTaskRequest {
	s.DesktopId = v
	return s
}

type StartVirusScanTaskResponseBody struct {
	ScanTaskId *int64  `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartVirusScanTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartVirusScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartVirusScanTaskResponseBody) SetScanTaskId(v int64) *StartVirusScanTaskResponseBody {
	s.ScanTaskId = &v
	return s
}

func (s *StartVirusScanTaskResponseBody) SetRequestId(v string) *StartVirusScanTaskResponseBody {
	s.RequestId = &v
	return s
}

type StartVirusScanTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartVirusScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartVirusScanTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StartVirusScanTaskResponse) GoString() string {
	return s.String()
}

func (s *StartVirusScanTaskResponse) SetHeaders(v map[string]*string) *StartVirusScanTaskResponse {
	s.Headers = v
	return s
}

func (s *StartVirusScanTaskResponse) SetBody(v *StartVirusScanTaskResponseBody) *StartVirusScanTaskResponse {
	s.Body = v
	return s
}

type StopDesktopsRequest struct {
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StoppedMode *string   `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty"`
	DesktopId   []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s StopDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StopDesktopsRequest) SetRegionId(v string) *StopDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *StopDesktopsRequest) SetStoppedMode(v string) *StopDesktopsRequest {
	s.StoppedMode = &v
	return s
}

func (s *StopDesktopsRequest) SetDesktopId(v []*string) *StopDesktopsRequest {
	s.DesktopId = v
	return s
}

type StopDesktopsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *StopDesktopsResponseBody) SetRequestId(v string) *StopDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type StopDesktopsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsResponse) GoString() string {
	return s.String()
}

func (s *StopDesktopsResponse) SetHeaders(v map[string]*string) *StopDesktopsResponse {
	s.Headers = v
	return s
}

func (s *StopDesktopsResponse) SetBody(v *StopDesktopsResponseBody) *StopDesktopsResponse {
	s.Body = v
	return s
}

type StopInvocationRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InvokeId  *string   `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
}

func (s StopInvocationRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationRequest) GoString() string {
	return s.String()
}

func (s *StopInvocationRequest) SetRegionId(v string) *StopInvocationRequest {
	s.RegionId = &v
	return s
}

func (s *StopInvocationRequest) SetInvokeId(v string) *StopInvocationRequest {
	s.InvokeId = &v
	return s
}

func (s *StopInvocationRequest) SetDesktopId(v []*string) *StopInvocationRequest {
	s.DesktopId = v
	return s
}

type StopInvocationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInvocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationResponseBody) GoString() string {
	return s.String()
}

func (s *StopInvocationResponseBody) SetRequestId(v string) *StopInvocationResponseBody {
	s.RequestId = &v
	return s
}

type StopInvocationResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopInvocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopInvocationResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationResponse) GoString() string {
	return s.String()
}

func (s *StopInvocationResponse) SetHeaders(v map[string]*string) *StopInvocationResponse {
	s.Headers = v
	return s
}

func (s *StopInvocationResponse) SetBody(v *StopInvocationResponseBody) *StopInvocationResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnlockVirtualMFADeviceRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s UnlockVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnlockVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *UnlockVirtualMFADeviceRequest) SetRegionId(v string) *UnlockVirtualMFADeviceRequest {
	s.RegionId = &v
	return s
}

func (s *UnlockVirtualMFADeviceRequest) SetSerialNumber(v string) *UnlockVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

type UnlockVirtualMFADeviceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockVirtualMFADeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnlockVirtualMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockVirtualMFADeviceResponseBody) SetRequestId(v string) *UnlockVirtualMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

type UnlockVirtualMFADeviceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnlockVirtualMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnlockVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnlockVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *UnlockVirtualMFADeviceResponse) SetHeaders(v map[string]*string) *UnlockVirtualMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *UnlockVirtualMFADeviceResponse) SetBody(v *UnlockVirtualMFADeviceResponseBody) *UnlockVirtualMFADeviceResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ecd"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUserToDesktopGroupWithOptions(request *AddUserToDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *AddUserToDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddUserToDesktopGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddUserToDesktopGroup"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUserToDesktopGroup(request *AddUserToDesktopGroupRequest) (_result *AddUserToDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserToDesktopGroupResponse{}
	_body, _err := client.AddUserToDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUserToSecurityCenterWhiteListWithOptions(request *AddUserToSecurityCenterWhiteListRequest, runtime *util.RuntimeOptions) (_result *AddUserToSecurityCenterWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddUserToSecurityCenterWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddUserToSecurityCenterWhiteList"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUserToSecurityCenterWhiteList(request *AddUserToSecurityCenterWhiteListRequest) (_result *AddUserToSecurityCenterWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserToSecurityCenterWhiteListResponse{}
	_body, _err := client.AddUserToSecurityCenterWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachCenWithOptions(request *AttachCenRequest, runtime *util.RuntimeOptions) (_result *AttachCenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachCenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachCen"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachCen(request *AttachCenRequest) (_result *AttachCenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachCenResponse{}
	_body, _err := client.AttachCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckUserInSecurityCenterWhiteListWithOptions(request *CheckUserInSecurityCenterWhiteListRequest, runtime *util.RuntimeOptions) (_result *CheckUserInSecurityCenterWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckUserInSecurityCenterWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckUserInSecurityCenterWhiteList"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckUserInSecurityCenterWhiteList(request *CheckUserInSecurityCenterWhiteListRequest) (_result *CheckUserInSecurityCenterWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckUserInSecurityCenterWhiteListResponse{}
	_body, _err := client.CheckUserInSecurityCenterWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClonePolicyGroupWithOptions(request *ClonePolicyGroupRequest, runtime *util.RuntimeOptions) (_result *ClonePolicyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ClonePolicyGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ClonePolicyGroup"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClonePolicyGroup(request *ClonePolicyGroupRequest) (_result *ClonePolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClonePolicyGroupResponse{}
	_body, _err := client.ClonePolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateADConnectorDirectoryWithOptions(request *CreateADConnectorDirectoryRequest, runtime *util.RuntimeOptions) (_result *CreateADConnectorDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateADConnectorDirectoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateADConnectorDirectory"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateADConnectorDirectory(request *CreateADConnectorDirectoryRequest) (_result *CreateADConnectorDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateADConnectorDirectoryResponse{}
	_body, _err := client.CreateADConnectorDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateADConnectorOfficeSiteWithOptions(request *CreateADConnectorOfficeSiteRequest, runtime *util.RuntimeOptions) (_result *CreateADConnectorOfficeSiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateADConnectorOfficeSiteResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateADConnectorOfficeSite"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateADConnectorOfficeSite(request *CreateADConnectorOfficeSiteRequest) (_result *CreateADConnectorOfficeSiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateADConnectorOfficeSiteResponse{}
	_body, _err := client.CreateADConnectorOfficeSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBundleWithOptions(request *CreateBundleRequest, runtime *util.RuntimeOptions) (_result *CreateBundleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateBundleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateBundle"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBundle(request *CreateBundleRequest) (_result *CreateBundleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBundleResponse{}
	_body, _err := client.CreateBundleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDesktopGroupWithOptions(request *CreateDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *CreateDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDesktopGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDesktopGroup"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDesktopGroup(request *CreateDesktopGroupRequest) (_result *CreateDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDesktopGroupResponse{}
	_body, _err := client.CreateDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDesktopsWithOptions(request *CreateDesktopsRequest, runtime *util.RuntimeOptions) (_result *CreateDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDesktopsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDesktops"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDesktops(request *CreateDesktopsRequest) (_result *CreateDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDesktopsResponse{}
	_body, _err := client.CreateDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDesktopsLiteWithOptions(request *CreateDesktopsLiteRequest, runtime *util.RuntimeOptions) (_result *CreateDesktopsLiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDesktopsLiteResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDesktopsLite"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDesktopsLite(request *CreateDesktopsLiteRequest) (_result *CreateDesktopsLiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDesktopsLiteResponse{}
	_body, _err := client.CreateDesktopsLiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageWithOptions(request *CreateImageRequest, runtime *util.RuntimeOptions) (_result *CreateImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateImage"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImage(request *CreateImageRequest) (_result *CreateImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageResponse{}
	_body, _err := client.CreateImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNASFileSystemWithOptions(request *CreateNASFileSystemRequest, runtime *util.RuntimeOptions) (_result *CreateNASFileSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateNASFileSystemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateNASFileSystem"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNASFileSystem(request *CreateNASFileSystemRequest) (_result *CreateNASFileSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNASFileSystemResponse{}
	_body, _err := client.CreateNASFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNetworkPackageWithOptions(request *CreateNetworkPackageRequest, runtime *util.RuntimeOptions) (_result *CreateNetworkPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateNetworkPackageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateNetworkPackage"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNetworkPackage(request *CreateNetworkPackageRequest) (_result *CreateNetworkPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNetworkPackageResponse{}
	_body, _err := client.CreateNetworkPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePolicyGroupWithOptions(request *CreatePolicyGroupRequest, runtime *util.RuntimeOptions) (_result *CreatePolicyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreatePolicyGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreatePolicyGroup"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePolicyGroup(request *CreatePolicyGroupRequest) (_result *CreatePolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePolicyGroupResponse{}
	_body, _err := client.CreatePolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRAMDirectoryWithOptions(request *CreateRAMDirectoryRequest, runtime *util.RuntimeOptions) (_result *CreateRAMDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRAMDirectoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRAMDirectory"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRAMDirectory(request *CreateRAMDirectoryRequest) (_result *CreateRAMDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRAMDirectoryResponse{}
	_body, _err := client.CreateRAMDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateScaleStrategyWithOptions(request *CreateScaleStrategyRequest, runtime *util.RuntimeOptions) (_result *CreateScaleStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateScaleStrategyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateScaleStrategy"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateScaleStrategy(request *CreateScaleStrategyRequest) (_result *CreateScaleStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScaleStrategyResponse{}
	_body, _err := client.CreateScaleStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceLinkedRoleWithOptions(request *CreateServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateServiceLinkedRole"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (_result *CreateServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CreateServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSimpleOfficeSiteWithOptions(request *CreateSimpleOfficeSiteRequest, runtime *util.RuntimeOptions) (_result *CreateSimpleOfficeSiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSimpleOfficeSiteResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSimpleOfficeSite"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSimpleOfficeSite(request *CreateSimpleOfficeSiteRequest) (_result *CreateSimpleOfficeSiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSimpleOfficeSiteResponse{}
	_body, _err := client.CreateSimpleOfficeSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSnapshotWithOptions(request *CreateSnapshotRequest, runtime *util.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSnapshot"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSnapshot(request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBundlesWithOptions(request *DeleteBundlesRequest, runtime *util.RuntimeOptions) (_result *DeleteBundlesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteBundlesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteBundles"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBundles(request *DeleteBundlesRequest) (_result *DeleteBundlesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBundlesResponse{}
	_body, _err := client.DeleteBundlesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDesktopGroupWithOptions(request *DeleteDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDesktopGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDesktopGroup"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDesktopGroup(request *DeleteDesktopGroupRequest) (_result *DeleteDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDesktopGroupResponse{}
	_body, _err := client.DeleteDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDesktopsWithOptions(request *DeleteDesktopsRequest, runtime *util.RuntimeOptions) (_result *DeleteDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDesktopsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDesktops"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDesktops(request *DeleteDesktopsRequest) (_result *DeleteDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDesktopsResponse{}
	_body, _err := client.DeleteDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDirectoriesWithOptions(request *DeleteDirectoriesRequest, runtime *util.RuntimeOptions) (_result *DeleteDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDirectoriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDirectories"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDirectories(request *DeleteDirectoriesRequest) (_result *DeleteDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDirectoriesResponse{}
	_body, _err := client.DeleteDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImagesWithOptions(request *DeleteImagesRequest, runtime *util.RuntimeOptions) (_result *DeleteImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteImagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteImages"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImages(request *DeleteImagesRequest) (_result *DeleteImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImagesResponse{}
	_body, _err := client.DeleteImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNASFileSystemsWithOptions(request *DeleteNASFileSystemsRequest, runtime *util.RuntimeOptions) (_result *DeleteNASFileSystemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteNASFileSystemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteNASFileSystems"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNASFileSystems(request *DeleteNASFileSystemsRequest) (_result *DeleteNASFileSystemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNASFileSystemsResponse{}
	_body, _err := client.DeleteNASFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNetworkPackagesWithOptions(request *DeleteNetworkPackagesRequest, runtime *util.RuntimeOptions) (_result *DeleteNetworkPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteNetworkPackagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteNetworkPackages"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNetworkPackages(request *DeleteNetworkPackagesRequest) (_result *DeleteNetworkPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNetworkPackagesResponse{}
	_body, _err := client.DeleteNetworkPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteOfficeSitesWithOptions(request *DeleteOfficeSitesRequest, runtime *util.RuntimeOptions) (_result *DeleteOfficeSitesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteOfficeSitesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteOfficeSites"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteOfficeSites(request *DeleteOfficeSitesRequest) (_result *DeleteOfficeSitesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteOfficeSitesResponse{}
	_body, _err := client.DeleteOfficeSitesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePolicyGroupsWithOptions(request *DeletePolicyGroupsRequest, runtime *util.RuntimeOptions) (_result *DeletePolicyGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeletePolicyGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeletePolicyGroups"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePolicyGroups(request *DeletePolicyGroupsRequest) (_result *DeletePolicyGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePolicyGroupsResponse{}
	_body, _err := client.DeletePolicyGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteScaleStrategyWithOptions(request *DeleteScaleStrategyRequest, runtime *util.RuntimeOptions) (_result *DeleteScaleStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteScaleStrategyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteScaleStrategy"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteScaleStrategy(request *DeleteScaleStrategyRequest) (_result *DeleteScaleStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScaleStrategyResponse{}
	_body, _err := client.DeleteScaleStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSnapshotWithOptions(request *DeleteSnapshotRequest, runtime *util.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSnapshot"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (_result *DeleteSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DeleteSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVirtualMFADeviceWithOptions(request *DeleteVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVirtualMFADevice"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVirtualMFADevice(request *DeleteVirtualMFADeviceRequest) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.DeleteVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAlarmEventStackInfoWithOptions(request *DescribeAlarmEventStackInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeAlarmEventStackInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAlarmEventStackInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAlarmEventStackInfo"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAlarmEventStackInfo(request *DescribeAlarmEventStackInfoRequest) (_result *DescribeAlarmEventStackInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlarmEventStackInfoResponse{}
	_body, _err := client.DescribeAlarmEventStackInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBundlesWithOptions(request *DescribeBundlesRequest, runtime *util.RuntimeOptions) (_result *DescribeBundlesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBundlesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBundles"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBundles(request *DescribeBundlesRequest) (_result *DescribeBundlesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBundlesResponse{}
	_body, _err := client.DescribeBundlesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCensWithOptions(request *DescribeCensRequest, runtime *util.RuntimeOptions) (_result *DescribeCensResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCensResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCens"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCens(request *DescribeCensRequest) (_result *DescribeCensResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCensResponse{}
	_body, _err := client.DescribeCensWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClientEventsWithOptions(request *DescribeClientEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeClientEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeClientEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeClientEvents"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClientEvents(request *DescribeClientEventsRequest) (_result *DescribeClientEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClientEventsResponse{}
	_body, _err := client.DescribeClientEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopIdsByVulNamesWithOptions(request *DescribeDesktopIdsByVulNamesRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopIdsByVulNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDesktopIdsByVulNamesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDesktopIdsByVulNames"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktopIdsByVulNames(request *DescribeDesktopIdsByVulNamesRequest) (_result *DescribeDesktopIdsByVulNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopIdsByVulNamesResponse{}
	_body, _err := client.DescribeDesktopIdsByVulNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopPolicysWithOptions(request *DescribeDesktopPolicysRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopPolicysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDesktopPolicysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDesktopPolicys"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktopPolicys(request *DescribeDesktopPolicysRequest) (_result *DescribeDesktopPolicysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopPolicysResponse{}
	_body, _err := client.DescribeDesktopPolicysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopsWithOptions(request *DescribeDesktopsRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDesktopsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDesktops"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktops(request *DescribeDesktopsRequest) (_result *DescribeDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopsResponse{}
	_body, _err := client.DescribeDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopsInGroupWithOptions(request *DescribeDesktopsInGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopsInGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDesktopsInGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDesktopsInGroup"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktopsInGroup(request *DescribeDesktopsInGroupRequest) (_result *DescribeDesktopsInGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopsInGroupResponse{}
	_body, _err := client.DescribeDesktopsInGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDesktopTypesWithOptions(request *DescribeDesktopTypesRequest, runtime *util.RuntimeOptions) (_result *DescribeDesktopTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDesktopTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDesktopTypes"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDesktopTypes(request *DescribeDesktopTypesRequest) (_result *DescribeDesktopTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDesktopTypesResponse{}
	_body, _err := client.DescribeDesktopTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDirectoriesWithOptions(request *DescribeDirectoriesRequest, runtime *util.RuntimeOptions) (_result *DescribeDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDirectories"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDirectories(request *DescribeDirectoriesRequest) (_result *DescribeDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.DescribeDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFrontVulPatchListWithOptions(request *DescribeFrontVulPatchListRequest, runtime *util.RuntimeOptions) (_result *DescribeFrontVulPatchListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFrontVulPatchListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFrontVulPatchList"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFrontVulPatchList(request *DescribeFrontVulPatchListRequest) (_result *DescribeFrontVulPatchListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFrontVulPatchListResponse{}
	_body, _err := client.DescribeFrontVulPatchListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGroupedVulWithOptions(request *DescribeGroupedVulRequest, runtime *util.RuntimeOptions) (_result *DescribeGroupedVulResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGroupedVulResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGroupedVul"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGroupedVul(request *DescribeGroupedVulRequest) (_result *DescribeGroupedVulResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGroupedVulResponse{}
	_body, _err := client.DescribeGroupedVulWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImagesWithOptions(request *DescribeImagesRequest, runtime *util.RuntimeOptions) (_result *DescribeImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeImages"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImages(request *DescribeImagesRequest) (_result *DescribeImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DescribeImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInvocationsWithOptions(request *DescribeInvocationsRequest, runtime *util.RuntimeOptions) (_result *DescribeInvocationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInvocations"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInvocations(request *DescribeInvocationsRequest) (_result *DescribeInvocationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.DescribeInvocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeModificationPriceWithOptions(request *DescribeModificationPriceRequest, runtime *util.RuntimeOptions) (_result *DescribeModificationPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeModificationPriceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeModificationPrice"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeModificationPrice(request *DescribeModificationPriceRequest) (_result *DescribeModificationPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeModificationPriceResponse{}
	_body, _err := client.DescribeModificationPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNASFileSystemsWithOptions(request *DescribeNASFileSystemsRequest, runtime *util.RuntimeOptions) (_result *DescribeNASFileSystemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeNASFileSystemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNASFileSystems"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNASFileSystems(request *DescribeNASFileSystemsRequest) (_result *DescribeNASFileSystemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNASFileSystemsResponse{}
	_body, _err := client.DescribeNASFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNetworkPackagesWithOptions(request *DescribeNetworkPackagesRequest, runtime *util.RuntimeOptions) (_result *DescribeNetworkPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeNetworkPackagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNetworkPackages"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNetworkPackages(request *DescribeNetworkPackagesRequest) (_result *DescribeNetworkPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNetworkPackagesResponse{}
	_body, _err := client.DescribeNetworkPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOfficeSitesWithOptions(request *DescribeOfficeSitesRequest, runtime *util.RuntimeOptions) (_result *DescribeOfficeSitesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeOfficeSitesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeOfficeSites"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOfficeSites(request *DescribeOfficeSitesRequest) (_result *DescribeOfficeSitesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOfficeSitesResponse{}
	_body, _err := client.DescribeOfficeSitesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePolicyGroupsWithOptions(request *DescribePolicyGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribePolicyGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePolicyGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePolicyGroups"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePolicyGroups(request *DescribePolicyGroupsRequest) (_result *DescribePolicyGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePolicyGroupsResponse{}
	_body, _err := client.DescribePolicyGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePostPaidDesktopBillsWithOptions(request *DescribePostPaidDesktopBillsRequest, runtime *util.RuntimeOptions) (_result *DescribePostPaidDesktopBillsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePostPaidDesktopBillsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePostPaidDesktopBills"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePostPaidDesktopBills(request *DescribePostPaidDesktopBillsRequest) (_result *DescribePostPaidDesktopBillsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePostPaidDesktopBillsResponse{}
	_body, _err := client.DescribePostPaidDesktopBillsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePriceWithOptions(request *DescribePriceRequest, runtime *util.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePriceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePrice"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePrice(request *DescribePriceRequest) (_result *DescribePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePriceResponse{}
	_body, _err := client.DescribePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScaleStrategysWithOptions(request *DescribeScaleStrategysRequest, runtime *util.RuntimeOptions) (_result *DescribeScaleStrategysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScaleStrategysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScaleStrategys"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScaleStrategys(request *DescribeScaleStrategysRequest) (_result *DescribeScaleStrategysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScaleStrategysResponse{}
	_body, _err := client.DescribeScaleStrategysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScanTaskProgressWithOptions(request *DescribeScanTaskProgressRequest, runtime *util.RuntimeOptions) (_result *DescribeScanTaskProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScanTaskProgressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScanTaskProgress"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScanTaskProgress(request *DescribeScanTaskProgressRequest) (_result *DescribeScanTaskProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScanTaskProgressResponse{}
	_body, _err := client.DescribeScanTaskProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityEventOperationsWithOptions(request *DescribeSecurityEventOperationsRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityEventOperationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSecurityEventOperationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecurityEventOperations"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityEventOperations(request *DescribeSecurityEventOperationsRequest) (_result *DescribeSecurityEventOperationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityEventOperationsResponse{}
	_body, _err := client.DescribeSecurityEventOperationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityEventOperationStatusWithOptions(request *DescribeSecurityEventOperationStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityEventOperationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSecurityEventOperationStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecurityEventOperationStatus"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityEventOperationStatus(request *DescribeSecurityEventOperationStatusRequest) (_result *DescribeSecurityEventOperationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityEventOperationStatusResponse{}
	_body, _err := client.DescribeSecurityEventOperationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSnapshotsWithOptions(request *DescribeSnapshotsRequest, runtime *util.RuntimeOptions) (_result *DescribeSnapshotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSnapshots"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (_result *DescribeSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.DescribeSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSuspEventOverviewWithOptions(request *DescribeSuspEventOverviewRequest, runtime *util.RuntimeOptions) (_result *DescribeSuspEventOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSuspEventOverviewResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSuspEventOverview"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSuspEventOverview(request *DescribeSuspEventOverviewRequest) (_result *DescribeSuspEventOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSuspEventOverviewResponse{}
	_body, _err := client.DescribeSuspEventOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSuspEventQuaraFilesWithOptions(request *DescribeSuspEventQuaraFilesRequest, runtime *util.RuntimeOptions) (_result *DescribeSuspEventQuaraFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSuspEventQuaraFilesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSuspEventQuaraFiles"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSuspEventQuaraFiles(request *DescribeSuspEventQuaraFilesRequest) (_result *DescribeSuspEventQuaraFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSuspEventQuaraFilesResponse{}
	_body, _err := client.DescribeSuspEventQuaraFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSuspEventsWithOptions(request *DescribeSuspEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeSuspEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSuspEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSuspEvents"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSuspEvents(request *DescribeSuspEventsRequest) (_result *DescribeSuspEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSuspEventsResponse{}
	_body, _err := client.DescribeSuspEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserConnectionRecordsWithOptions(request *DescribeUserConnectionRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeUserConnectionRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserConnectionRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserConnectionRecords"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserConnectionRecords(request *DescribeUserConnectionRecordsRequest) (_result *DescribeUserConnectionRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserConnectionRecordsResponse{}
	_body, _err := client.DescribeUserConnectionRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUsersInGroupWithOptions(request *DescribeUsersInGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeUsersInGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUsersInGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUsersInGroup"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUsersInGroup(request *DescribeUsersInGroupRequest) (_result *DescribeUsersInGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUsersInGroupResponse{}
	_body, _err := client.DescribeUsersInGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVirtualMFADevicesWithOptions(request *DescribeVirtualMFADevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeVirtualMFADevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVirtualMFADevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVirtualMFADevices"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVirtualMFADevices(request *DescribeVirtualMFADevicesRequest) (_result *DescribeVirtualMFADevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVirtualMFADevicesResponse{}
	_body, _err := client.DescribeVirtualMFADevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVulDetailsWithOptions(request *DescribeVulDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeVulDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVulDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVulDetails"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVulDetails(request *DescribeVulDetailsRequest) (_result *DescribeVulDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVulDetailsResponse{}
	_body, _err := client.DescribeVulDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVulListWithOptions(request *DescribeVulListRequest, runtime *util.RuntimeOptions) (_result *DescribeVulListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVulListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVulList"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVulList(request *DescribeVulListRequest) (_result *DescribeVulListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVulListResponse{}
	_body, _err := client.DescribeVulListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVulOverviewWithOptions(request *DescribeVulOverviewRequest, runtime *util.RuntimeOptions) (_result *DescribeVulOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVulOverviewResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVulOverview"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVulOverview(request *DescribeVulOverviewRequest) (_result *DescribeVulOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVulOverviewResponse{}
	_body, _err := client.DescribeVulOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeZones"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachCenWithOptions(request *DetachCenRequest, runtime *util.RuntimeOptions) (_result *DetachCenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachCenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachCen"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachCen(request *DetachCenRequest) (_result *DetachCenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachCenResponse{}
	_body, _err := client.DetachCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DoCheckResourceWithOptions(request *DoCheckResourceRequest, runtime *util.RuntimeOptions) (_result *DoCheckResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DoCheckResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DoCheckResource"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoCheckResource(request *DoCheckResourceRequest) (_result *DoCheckResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoCheckResourceResponse{}
	_body, _err := client.DoCheckResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DoLogicalDeleteResourceWithOptions(request *DoLogicalDeleteResourceRequest, runtime *util.RuntimeOptions) (_result *DoLogicalDeleteResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DoLogicalDeleteResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DoLogicalDeleteResource"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoLogicalDeleteResource(request *DoLogicalDeleteResourceRequest) (_result *DoLogicalDeleteResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoLogicalDeleteResourceResponse{}
	_body, _err := client.DoLogicalDeleteResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DoPhysicalDeleteResourceWithOptions(request *DoPhysicalDeleteResourceRequest, runtime *util.RuntimeOptions) (_result *DoPhysicalDeleteResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DoPhysicalDeleteResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DoPhysicalDeleteResource"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoPhysicalDeleteResource(request *DoPhysicalDeleteResourceRequest) (_result *DoPhysicalDeleteResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoPhysicalDeleteResourceResponse{}
	_body, _err := client.DoPhysicalDeleteResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConnectionTicketWithOptions(request *GetConnectionTicketRequest, runtime *util.RuntimeOptions) (_result *GetConnectionTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetConnectionTicket"), tea.String("2020-09-30"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConnectionTicket(request *GetConnectionTicketRequest) (_result *GetConnectionTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.GetConnectionTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDesktopGroupDetailWithOptions(request *GetDesktopGroupDetailRequest, runtime *util.RuntimeOptions) (_result *GetDesktopGroupDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDesktopGroupDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDesktopGroupDetail"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDesktopGroupDetail(request *GetDesktopGroupDetailRequest) (_result *GetDesktopGroupDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDesktopGroupDetailResponse{}
	_body, _err := client.GetDesktopGroupDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDesktopUsersWithOptions(request *GetDesktopUsersRequest, runtime *util.RuntimeOptions) (_result *GetDesktopUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDesktopUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDesktopUsers"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDesktopUsers(request *GetDesktopUsersRequest) (_result *GetDesktopUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDesktopUsersResponse{}
	_body, _err := client.GetDesktopUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDirectorySsoStatusWithOptions(request *GetDirectorySsoStatusRequest, runtime *util.RuntimeOptions) (_result *GetDirectorySsoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDirectorySsoStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDirectorySsoStatus"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDirectorySsoStatus(request *GetDirectorySsoStatusRequest) (_result *GetDirectorySsoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDirectorySsoStatusResponse{}
	_body, _err := client.GetDirectorySsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOfficeSiteSsoStatusWithOptions(request *GetOfficeSiteSsoStatusRequest, runtime *util.RuntimeOptions) (_result *GetOfficeSiteSsoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetOfficeSiteSsoStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOfficeSiteSsoStatus"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOfficeSiteSsoStatus(request *GetOfficeSiteSsoStatusRequest) (_result *GetOfficeSiteSsoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOfficeSiteSsoStatusResponse{}
	_body, _err := client.GetOfficeSiteSsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSpMetadataWithOptions(request *GetSpMetadataRequest, runtime *util.RuntimeOptions) (_result *GetSpMetadataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSpMetadataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSpMetadata"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSpMetadata(request *GetSpMetadataRequest) (_result *GetSpMetadataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSpMetadataResponse{}
	_body, _err := client.GetSpMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HandleSecurityEventsWithOptions(request *HandleSecurityEventsRequest, runtime *util.RuntimeOptions) (_result *HandleSecurityEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &HandleSecurityEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("HandleSecurityEvents"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HandleSecurityEvents(request *HandleSecurityEventsRequest) (_result *HandleSecurityEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HandleSecurityEventsResponse{}
	_body, _err := client.HandleSecurityEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDirectoryUsersWithOptions(request *ListDirectoryUsersRequest, runtime *util.RuntimeOptions) (_result *ListDirectoryUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDirectoryUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDirectoryUsers"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDirectoryUsers(request *ListDirectoryUsersRequest) (_result *ListDirectoryUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDirectoryUsersResponse{}
	_body, _err := client.ListDirectoryUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOfficeSiteOverviewWithOptions(request *ListOfficeSiteOverviewRequest, runtime *util.RuntimeOptions) (_result *ListOfficeSiteOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOfficeSiteOverviewResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOfficeSiteOverview"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOfficeSiteOverview(request *ListOfficeSiteOverviewRequest) (_result *ListOfficeSiteOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOfficeSiteOverviewResponse{}
	_body, _err := client.ListOfficeSiteOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOfficeSiteUsersWithOptions(request *ListOfficeSiteUsersRequest, runtime *util.RuntimeOptions) (_result *ListOfficeSiteUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOfficeSiteUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOfficeSiteUsers"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOfficeSiteUsers(request *ListOfficeSiteUsersRequest) (_result *ListOfficeSiteUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOfficeSiteUsersResponse{}
	_body, _err := client.ListOfficeSiteUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LockVirtualMFADeviceWithOptions(request *LockVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *LockVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &LockVirtualMFADeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("LockVirtualMFADevice"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LockVirtualMFADevice(request *LockVirtualMFADeviceRequest) (_result *LockVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LockVirtualMFADeviceResponse{}
	_body, _err := client.LockVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyADConnectorDirectoryWithOptions(request *ModifyADConnectorDirectoryRequest, runtime *util.RuntimeOptions) (_result *ModifyADConnectorDirectoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyADConnectorDirectoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyADConnectorDirectory"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyADConnectorDirectory(request *ModifyADConnectorDirectoryRequest) (_result *ModifyADConnectorDirectoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyADConnectorDirectoryResponse{}
	_body, _err := client.ModifyADConnectorDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyADConnectorOfficeSiteWithOptions(request *ModifyADConnectorOfficeSiteRequest, runtime *util.RuntimeOptions) (_result *ModifyADConnectorOfficeSiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyADConnectorOfficeSiteResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyADConnectorOfficeSite"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyADConnectorOfficeSite(request *ModifyADConnectorOfficeSiteRequest) (_result *ModifyADConnectorOfficeSiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyADConnectorOfficeSiteResponse{}
	_body, _err := client.ModifyADConnectorOfficeSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBundleWithOptions(request *ModifyBundleRequest, runtime *util.RuntimeOptions) (_result *ModifyBundleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBundleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBundle"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBundle(request *ModifyBundleRequest) (_result *ModifyBundleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBundleResponse{}
	_body, _err := client.ModifyBundleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopGroupWithOptions(request *ModifyDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDesktopGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDesktopGroup"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopGroup(request *ModifyDesktopGroupRequest) (_result *ModifyDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopGroupResponse{}
	_body, _err := client.ModifyDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopNameWithOptions(request *ModifyDesktopNameRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDesktopNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDesktopName"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopName(request *ModifyDesktopNameRequest) (_result *ModifyDesktopNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopNameResponse{}
	_body, _err := client.ModifyDesktopNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopPolicysWithOptions(request *ModifyDesktopPolicysRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopPolicysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDesktopPolicysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDesktopPolicys"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopPolicys(request *ModifyDesktopPolicysRequest) (_result *ModifyDesktopPolicysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopPolicysResponse{}
	_body, _err := client.ModifyDesktopPolicysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopSpecWithOptions(request *ModifyDesktopSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDesktopSpecResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDesktopSpec"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopSpec(request *ModifyDesktopSpecRequest) (_result *ModifyDesktopSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopSpecResponse{}
	_body, _err := client.ModifyDesktopSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDesktopsPolicyGroupWithOptions(request *ModifyDesktopsPolicyGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDesktopsPolicyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDesktopsPolicyGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDesktopsPolicyGroup"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDesktopsPolicyGroup(request *ModifyDesktopsPolicyGroupRequest) (_result *ModifyDesktopsPolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDesktopsPolicyGroupResponse{}
	_body, _err := client.ModifyDesktopsPolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyEntitlementWithOptions(request *ModifyEntitlementRequest, runtime *util.RuntimeOptions) (_result *ModifyEntitlementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyEntitlementResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyEntitlement"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyEntitlement(request *ModifyEntitlementRequest) (_result *ModifyEntitlementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyEntitlementResponse{}
	_body, _err := client.ModifyEntitlementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyImageAttributeWithOptions(request *ModifyImageAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyImageAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyImageAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyImageAttribute"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyImageAttribute(request *ModifyImageAttributeRequest) (_result *ModifyImageAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyImageAttributeResponse{}
	_body, _err := client.ModifyImageAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNASDefaultMountTargetWithOptions(request *ModifyNASDefaultMountTargetRequest, runtime *util.RuntimeOptions) (_result *ModifyNASDefaultMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyNASDefaultMountTargetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyNASDefaultMountTarget"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNASDefaultMountTarget(request *ModifyNASDefaultMountTargetRequest) (_result *ModifyNASDefaultMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNASDefaultMountTargetResponse{}
	_body, _err := client.ModifyNASDefaultMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNetworkPackageWithOptions(request *ModifyNetworkPackageRequest, runtime *util.RuntimeOptions) (_result *ModifyNetworkPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyNetworkPackageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyNetworkPackage"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNetworkPackage(request *ModifyNetworkPackageRequest) (_result *ModifyNetworkPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNetworkPackageResponse{}
	_body, _err := client.ModifyNetworkPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNetworkPackageEnabledWithOptions(request *ModifyNetworkPackageEnabledRequest, runtime *util.RuntimeOptions) (_result *ModifyNetworkPackageEnabledResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyNetworkPackageEnabledResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyNetworkPackageEnabled"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNetworkPackageEnabled(request *ModifyNetworkPackageEnabledRequest) (_result *ModifyNetworkPackageEnabledResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNetworkPackageEnabledResponse{}
	_body, _err := client.ModifyNetworkPackageEnabledWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyOfficeSiteAttributeWithOptions(request *ModifyOfficeSiteAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyOfficeSiteAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyOfficeSiteAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyOfficeSiteAttribute"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyOfficeSiteAttribute(request *ModifyOfficeSiteAttributeRequest) (_result *ModifyOfficeSiteAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOfficeSiteAttributeResponse{}
	_body, _err := client.ModifyOfficeSiteAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyOfficeSiteCrossDesktopAccessWithOptions(request *ModifyOfficeSiteCrossDesktopAccessRequest, runtime *util.RuntimeOptions) (_result *ModifyOfficeSiteCrossDesktopAccessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyOfficeSiteCrossDesktopAccessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyOfficeSiteCrossDesktopAccess"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyOfficeSiteCrossDesktopAccess(request *ModifyOfficeSiteCrossDesktopAccessRequest) (_result *ModifyOfficeSiteCrossDesktopAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOfficeSiteCrossDesktopAccessResponse{}
	_body, _err := client.ModifyOfficeSiteCrossDesktopAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyOfficeSiteMfaEnabledWithOptions(request *ModifyOfficeSiteMfaEnabledRequest, runtime *util.RuntimeOptions) (_result *ModifyOfficeSiteMfaEnabledResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyOfficeSiteMfaEnabledResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyOfficeSiteMfaEnabled"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyOfficeSiteMfaEnabled(request *ModifyOfficeSiteMfaEnabledRequest) (_result *ModifyOfficeSiteMfaEnabledResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOfficeSiteMfaEnabledResponse{}
	_body, _err := client.ModifyOfficeSiteMfaEnabledWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyOperateVulWithOptions(request *ModifyOperateVulRequest, runtime *util.RuntimeOptions) (_result *ModifyOperateVulResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyOperateVulResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyOperateVul"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyOperateVul(request *ModifyOperateVulRequest) (_result *ModifyOperateVulResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOperateVulResponse{}
	_body, _err := client.ModifyOperateVulWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPolicyGroupWithOptions(request *ModifyPolicyGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyPolicyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyPolicyGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyPolicyGroup"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPolicyGroup(request *ModifyPolicyGroupRequest) (_result *ModifyPolicyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPolicyGroupResponse{}
	_body, _err := client.ModifyPolicyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyScaleStrategyWithOptions(request *ModifyScaleStrategyRequest, runtime *util.RuntimeOptions) (_result *ModifyScaleStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyScaleStrategyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyScaleStrategy"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyScaleStrategy(request *ModifyScaleStrategyRequest) (_result *ModifyScaleStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyScaleStrategyResponse{}
	_body, _err := client.ModifyScaleStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserToDesktopGroupWithOptions(request *ModifyUserToDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyUserToDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyUserToDesktopGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyUserToDesktopGroup"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserToDesktopGroup(request *ModifyUserToDesktopGroupRequest) (_result *ModifyUserToDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserToDesktopGroupResponse{}
	_body, _err := client.ModifyUserToDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OperateVulsWithOptions(request *OperateVulsRequest, runtime *util.RuntimeOptions) (_result *OperateVulsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OperateVulsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OperateVuls"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OperateVuls(request *OperateVulsRequest) (_result *OperateVulsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateVulsResponse{}
	_body, _err := client.OperateVulsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PayOrderCallbackWithOptions(request *PayOrderCallbackRequest, runtime *util.RuntimeOptions) (_result *PayOrderCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PayOrderCallbackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PayOrderCallback"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PayOrderCallback(request *PayOrderCallbackRequest) (_result *PayOrderCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PayOrderCallbackResponse{}
	_body, _err := client.PayOrderCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebootDesktopsWithOptions(request *RebootDesktopsRequest, runtime *util.RuntimeOptions) (_result *RebootDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RebootDesktops"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebootDesktops(request *RebootDesktopsRequest) (_result *RebootDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.RebootDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebuildDesktopsWithOptions(request *RebuildDesktopsRequest, runtime *util.RuntimeOptions) (_result *RebuildDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RebuildDesktopsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RebuildDesktops"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebuildDesktops(request *RebuildDesktopsRequest) (_result *RebuildDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebuildDesktopsResponse{}
	_body, _err := client.RebuildDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecreateDesktopGroupWithOptions(request *RecreateDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *RecreateDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecreateDesktopGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecreateDesktopGroup"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecreateDesktopGroup(request *RecreateDesktopGroupRequest) (_result *RecreateDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecreateDesktopGroupResponse{}
	_body, _err := client.RecreateDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveUserFromDesktopGroupWithOptions(request *RemoveUserFromDesktopGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveUserFromDesktopGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveUserFromDesktopGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveUserFromDesktopGroup"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveUserFromDesktopGroup(request *RemoveUserFromDesktopGroupRequest) (_result *RemoveUserFromDesktopGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUserFromDesktopGroupResponse{}
	_body, _err := client.RemoveUserFromDesktopGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewDesktopsWithOptions(request *RenewDesktopsRequest, runtime *util.RuntimeOptions) (_result *RenewDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenewDesktopsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenewDesktops"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewDesktops(request *RenewDesktopsRequest) (_result *RenewDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewDesktopsResponse{}
	_body, _err := client.RenewDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetNASDefaultMountTargetWithOptions(request *ResetNASDefaultMountTargetRequest, runtime *util.RuntimeOptions) (_result *ResetNASDefaultMountTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetNASDefaultMountTargetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetNASDefaultMountTarget"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetNASDefaultMountTarget(request *ResetNASDefaultMountTargetRequest) (_result *ResetNASDefaultMountTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetNASDefaultMountTargetResponse{}
	_body, _err := client.ResetNASDefaultMountTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetSnapshotWithOptions(request *ResetSnapshotRequest, runtime *util.RuntimeOptions) (_result *ResetSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetSnapshotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetSnapshot"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetSnapshot(request *ResetSnapshotRequest) (_result *ResetSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetSnapshotResponse{}
	_body, _err := client.ResetSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackSuspEventQuaraFileWithOptions(request *RollbackSuspEventQuaraFileRequest, runtime *util.RuntimeOptions) (_result *RollbackSuspEventQuaraFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RollbackSuspEventQuaraFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RollbackSuspEventQuaraFile"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackSuspEventQuaraFile(request *RollbackSuspEventQuaraFileRequest) (_result *RollbackSuspEventQuaraFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackSuspEventQuaraFileResponse{}
	_body, _err := client.RollbackSuspEventQuaraFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunCommandWithOptions(request *RunCommandRequest, runtime *util.RuntimeOptions) (_result *RunCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RunCommandResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RunCommand"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunCommand(request *RunCommandRequest) (_result *RunCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCommandResponse{}
	_body, _err := client.RunCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDirectorySsoStatusWithOptions(request *SetDirectorySsoStatusRequest, runtime *util.RuntimeOptions) (_result *SetDirectorySsoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetDirectorySsoStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetDirectorySsoStatus"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDirectorySsoStatus(request *SetDirectorySsoStatusRequest) (_result *SetDirectorySsoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDirectorySsoStatusResponse{}
	_body, _err := client.SetDirectorySsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetIdpMetadataWithOptions(request *SetIdpMetadataRequest, runtime *util.RuntimeOptions) (_result *SetIdpMetadataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetIdpMetadataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetIdpMetadata"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetIdpMetadata(request *SetIdpMetadataRequest) (_result *SetIdpMetadataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetIdpMetadataResponse{}
	_body, _err := client.SetIdpMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetOfficeSiteSsoStatusWithOptions(request *SetOfficeSiteSsoStatusRequest, runtime *util.RuntimeOptions) (_result *SetOfficeSiteSsoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetOfficeSiteSsoStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetOfficeSiteSsoStatus"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetOfficeSiteSsoStatus(request *SetOfficeSiteSsoStatusRequest) (_result *SetOfficeSiteSsoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetOfficeSiteSsoStatusResponse{}
	_body, _err := client.SetOfficeSiteSsoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartDesktopsWithOptions(request *StartDesktopsRequest, runtime *util.RuntimeOptions) (_result *StartDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartDesktopsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartDesktops"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartDesktops(request *StartDesktopsRequest) (_result *StartDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDesktopsResponse{}
	_body, _err := client.StartDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartVirusScanTaskWithOptions(request *StartVirusScanTaskRequest, runtime *util.RuntimeOptions) (_result *StartVirusScanTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartVirusScanTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartVirusScanTask"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartVirusScanTask(request *StartVirusScanTaskRequest) (_result *StartVirusScanTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartVirusScanTaskResponse{}
	_body, _err := client.StartVirusScanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopDesktopsWithOptions(request *StopDesktopsRequest, runtime *util.RuntimeOptions) (_result *StopDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopDesktopsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopDesktops"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDesktops(request *StopDesktopsRequest) (_result *StopDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDesktopsResponse{}
	_body, _err := client.StopDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopInvocationWithOptions(request *StopInvocationRequest, runtime *util.RuntimeOptions) (_result *StopInvocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopInvocationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopInvocation"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopInvocation(request *StopInvocationRequest) (_result *StopInvocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopInvocationResponse{}
	_body, _err := client.StopInvocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnlockVirtualMFADeviceWithOptions(request *UnlockVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *UnlockVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnlockVirtualMFADeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnlockVirtualMFADevice"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnlockVirtualMFADevice(request *UnlockVirtualMFADeviceRequest) (_result *UnlockVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnlockVirtualMFADeviceResponse{}
	_body, _err := client.UnlockVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2020-09-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
