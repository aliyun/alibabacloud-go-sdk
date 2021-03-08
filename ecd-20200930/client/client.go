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

type CreateADConnectorDirectoryRequest struct {
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DomainName          *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainUserName      *string   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	DomainPassword      *string   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	DirectoryName       *string   `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	EnableAdminAccess   *bool     `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	DesktopAccessType   *string   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	SubDomainName       *string   `json:"SubDomainName,omitempty" xml:"SubDomainName,omitempty"`
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
	DirectoryId   *string                                               `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *CreateADConnectorDirectoryResponseBody) SetDirectoryId(v string) *CreateADConnectorDirectoryResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateADConnectorDirectoryResponseBody) SetRequestId(v string) *CreateADConnectorDirectoryResponseBody {
	s.RequestId = &v
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

type CreateDesktopsRequest struct {
	RegionId       *string                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	GroupId        *string                     `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	BundleId       *string                     `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	SystemDiskSize *int32                      `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	DataDiskSize   *int32                      `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DesktopName    *string                     `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	UserName       *string                     `json:"UserName,omitempty" xml:"UserName,omitempty"`
	VpcId          *string                     `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Amount         *int32                      `json:"Amount,omitempty" xml:"Amount,omitempty"`
	DirectoryId    *string                     `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	PolicyGroupId  *string                     `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	ChargeType     *string                     `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Period         *int32                      `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit     *string                     `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoPay        *bool                       `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
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

func (s *CreateDesktopsRequest) SetSystemDiskSize(v int32) *CreateDesktopsRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateDesktopsRequest) SetDataDiskSize(v int32) *CreateDesktopsRequest {
	s.DataDiskSize = &v
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
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	OrderId   *string   `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDesktopsResponseBody) SetRequestId(v string) *CreateDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDesktopsResponseBody) SetDesktopId(v []*string) *CreateDesktopsResponseBody {
	s.DesktopId = v
	return s
}

func (s *CreateDesktopsResponseBody) SetOrderId(v string) *CreateDesktopsResponseBody {
	s.OrderId = &v
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

type CreateImageRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	ImageName   *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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

type CreateImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s CreateImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageResponseBody) SetRequestId(v string) *CreateImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageResponseBody) SetImageId(v string) *CreateImageResponseBody {
	s.ImageId = &v
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

type CreatePolicyGroupRequest struct {
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Clipboard             *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	LocalDrive            *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UsbRedirect           *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	Watermark             *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	WatermarkType         *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	WatermarkCustomText   *string `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	WatermarkTransparency *string `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
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

type CreatePolicyGroupResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
}

func (s CreatePolicyGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupResponseBody) SetRequestId(v string) *CreatePolicyGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyGroupResponseBody) SetPolicyGroupId(v string) *CreatePolicyGroupResponseBody {
	s.PolicyGroupId = &v
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

type DescribeBundlesRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	UserName   *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Category   *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	BundleType *string   `json:"BundleType,omitempty" xml:"BundleType,omitempty"`
	BundleId   []*string `json:"BundleId,omitempty" xml:"BundleId,omitempty" type:"Repeated"`
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

func (s *DescribeBundlesRequest) SetUserName(v string) *DescribeBundlesRequest {
	s.UserName = &v
	return s
}

func (s *DescribeBundlesRequest) SetCategory(v string) *DescribeBundlesRequest {
	s.Category = &v
	return s
}

func (s *DescribeBundlesRequest) SetBundleType(v string) *DescribeBundlesRequest {
	s.BundleType = &v
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
	Disks                []*DescribeBundlesResponseBodyBundlesDisks              `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	BundleType           *string                                                 `json:"BundleType,omitempty" xml:"BundleType,omitempty"`
	Description          *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopType          *string                                                 `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	DesktopTypeAttribute *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute `json:"DesktopTypeAttribute,omitempty" xml:"DesktopTypeAttribute,omitempty" type:"Struct"`
	BundleId             *string                                                 `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	ImageId              *string                                                 `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	BundleName           *string                                                 `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
}

func (s DescribeBundlesResponseBodyBundles) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponseBodyBundles) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBodyBundles) SetDisks(v []*DescribeBundlesResponseBodyBundlesDisks) *DescribeBundlesResponseBodyBundles {
	s.Disks = v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetBundleType(v string) *DescribeBundlesResponseBodyBundles {
	s.BundleType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDescription(v string) *DescribeBundlesResponseBodyBundles {
	s.Description = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDesktopType(v string) *DescribeBundlesResponseBodyBundles {
	s.DesktopType = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundles) SetDesktopTypeAttribute(v *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) *DescribeBundlesResponseBodyBundles {
	s.DesktopTypeAttribute = v
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

type DescribeBundlesResponseBodyBundlesDisks struct {
	DiskSize *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
}

func (s DescribeBundlesResponseBodyBundlesDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeBundlesResponseBodyBundlesDisks) GoString() string {
	return s.String()
}

func (s *DescribeBundlesResponseBodyBundlesDisks) SetDiskSize(v int32) *DescribeBundlesResponseBodyBundlesDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeBundlesResponseBodyBundlesDisks) SetDiskType(v string) *DescribeBundlesResponseBodyBundlesDisks {
	s.DiskType = &v
	return s
}

type DescribeBundlesResponseBodyBundlesDesktopTypeAttribute struct {
	CpuCount   *int32  `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	GpuCount   *int32  `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	GpuSpec    *string `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	MemorySize *int32  `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
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

func (s *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute) SetGpuCount(v int32) *DescribeBundlesResponseBodyBundlesDesktopTypeAttribute {
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

type DescribeClientEventsRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	EndUserId   *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	DesktopId   *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopIp   *string `json:"DesktopIp,omitempty" xml:"DesktopIp,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EventType   *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	BytesReceived *string `json:"BytesReceived,omitempty" xml:"BytesReceived,omitempty"`
	DesktopIp     *string `json:"DesktopIp,omitempty" xml:"DesktopIp,omitempty"`
	EventTime     *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	BytesSend     *string `json:"BytesSend,omitempty" xml:"BytesSend,omitempty"`
	AliUid        *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	DesktopId     *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	EventId       *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	DirectoryType *string `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	EventType     *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	EndUserId     *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ClientIp      *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientOS      *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	DirectoryId   *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
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
	PolicyGroupId *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
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

func (s *DescribeDesktopsRequest) SetPolicyGroupId(v string) *DescribeDesktopsRequest {
	s.PolicyGroupId = &v
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
	CreationTime       *string                                      `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ChargeType         *string                                      `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DesktopName        *string                                      `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	Tags               []*DescribeDesktopsResponseBodyDesktopsTags  `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Disks              []*DescribeDesktopsResponseBodyDesktopsDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	SystemDiskSize     *int32                                       `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	PolicyGroupId      *string                                      `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	DesktopStatus      *string                                      `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	DesktopType        *string                                      `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	GpuCount           *int32                                       `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	Memory             *int64                                       `json:"Memory,omitempty" xml:"Memory,omitempty"`
	GpuSpec            *string                                      `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	ImageId            *string                                      `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	DirectoryId        *string                                      `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DataDiskCategory   *string                                      `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	SystemDiskCategory *string                                      `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	NetworkInterfaceId *int64                                       `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	DataDiskSize       *string                                      `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DesktopId          *string                                      `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	EndUserIds         []*string                                    `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	StartTime          *string                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Cpu                *int32                                       `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	ExpiredTime        *string                                      `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ConnectionStatus   *string                                      `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
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

func (s *DescribeDesktopsResponseBodyDesktops) SetTags(v []*DescribeDesktopsResponseBodyDesktopsTags) *DescribeDesktopsResponseBodyDesktops {
	s.Tags = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDisks(v []*DescribeDesktopsResponseBodyDesktopsDisks) *DescribeDesktopsResponseBodyDesktops {
	s.Disks = v
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

func (s *DescribeDesktopsResponseBodyDesktops) SetGpuCount(v int32) *DescribeDesktopsResponseBodyDesktops {
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

func (s *DescribeDesktopsResponseBodyDesktops) SetImageId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDirectoryId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DirectoryId = &v
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

func (s *DescribeDesktopsResponseBodyDesktops) SetNetworkInterfaceId(v int64) *DescribeDesktopsResponseBodyDesktops {
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

func (s *DescribeDesktopsResponseBodyDesktops) SetEndUserIds(v []*string) *DescribeDesktopsResponseBodyDesktops {
	s.EndUserIds = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetStartTime(v string) *DescribeDesktopsResponseBodyDesktops {
	s.StartTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetCpu(v int32) *DescribeDesktopsResponseBodyDesktops {
	s.Cpu = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetExpiredTime(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetConnectionStatus(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ConnectionStatus = &v
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

type DescribeDesktopsResponseBodyDesktopsDisks struct {
	DiskSize *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	DiskId   *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktopsDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsDisks) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetDiskSize(v int32) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetDiskType(v string) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.DiskType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetDiskId(v string) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.DiskId = &v
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

type DescribeDesktopTypesRequest struct {
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopTypeId      *string `json:"DesktopTypeId,omitempty" xml:"DesktopTypeId,omitempty"`
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
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
	GpuCount           *int32                                                       `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
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

func (s *DescribeDesktopTypesResponseBodyDesktopTypes) SetGpuCount(v int32) *DescribeDesktopTypesResponseBodyDesktopTypes {
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
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	DataDiskSize   *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
}

func (s DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize) SetSystemDiskSize(v int32) *DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize) SetDataDiskSize(v int32) *DescribeDesktopTypesResponseBodyDesktopTypesAllowDiskSize {
	s.DataDiskSize = &v
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
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DirectoryType *string   `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	MaxResults    *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	DirectoryId   []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" type:"Repeated"`
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
	Directories []*DescribeDirectoriesResponseBodyDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	NextToken   *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBody) SetDirectories(v []*DescribeDirectoriesResponseBodyDirectories) *DescribeDirectoriesResponseBody {
	s.Directories = v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetNextToken(v string) *DescribeDirectoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetRequestId(v string) *DescribeDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDirectoriesResponseBodyDirectories struct {
	EnableInternetAccess  *bool                                                     `json:"EnableInternetAccess,omitempty" xml:"EnableInternetAccess,omitempty"`
	VpcId                 *string                                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	CreationTime          *string                                                   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Status                *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	DomainPassword        *string                                                   `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	CustomSecurityGroupId *string                                                   `json:"CustomSecurityGroupId,omitempty" xml:"CustomSecurityGroupId,omitempty"`
	DomainUserName        *string                                                   `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	DesktopVpcEndpoint    *string                                                   `json:"DesktopVpcEndpoint,omitempty" xml:"DesktopVpcEndpoint,omitempty"`
	DesktopAccessType     *string                                                   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	DomainName            *string                                                   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DirectoryType         *string                                                   `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	DnsUserName           *string                                                   `json:"DnsUserName,omitempty" xml:"DnsUserName,omitempty"`
	Logs                  []*DescribeDirectoriesResponseBodyDirectoriesLogs         `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	TrustPassword         *string                                                   `json:"TrustPassword,omitempty" xml:"TrustPassword,omitempty"`
	VSwitchIds            []*string                                                 `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	Name                  *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	ADConnectors          []*DescribeDirectoriesResponseBodyDirectoriesADConnectors `json:"ADConnectors,omitempty" xml:"ADConnectors,omitempty" type:"Repeated"`
	DnsAddress            []*string                                                 `json:"DnsAddress,omitempty" xml:"DnsAddress,omitempty" type:"Repeated"`
	DirectoryId           *string                                                   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
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

func (s *DescribeDirectoriesResponseBodyDirectories) SetCustomSecurityGroupId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.CustomSecurityGroupId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDomainUserName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DomainUserName = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDesktopVpcEndpoint(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DesktopVpcEndpoint = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDesktopAccessType(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDomainName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DomainName = &v
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

func (s *DescribeDirectoriesResponseBodyDirectories) SetLogs(v []*DescribeDirectoriesResponseBodyDirectoriesLogs) *DescribeDirectoriesResponseBodyDirectories {
	s.Logs = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetTrustPassword(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.TrustPassword = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetVSwitchIds(v []*string) *DescribeDirectoriesResponseBodyDirectories {
	s.VSwitchIds = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetName(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.Name = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetADConnectors(v []*DescribeDirectoriesResponseBodyDirectoriesADConnectors) *DescribeDirectoriesResponseBodyDirectories {
	s.ADConnectors = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDnsAddress(v []*string) *DescribeDirectoriesResponseBodyDirectories {
	s.DnsAddress = v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDirectoryId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DirectoryId = &v
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

type DescribeDirectoriesResponseBodyDirectoriesADConnectors struct {
	ConnectorStatus    *string `json:"ConnectorStatus,omitempty" xml:"ConnectorStatus,omitempty"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	ADConnectorAddress *string `json:"ADConnectorAddress,omitempty" xml:"ADConnectorAddress,omitempty"`
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

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetNetworkInterfaceId(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectoriesADConnectors) SetADConnectorAddress(v string) *DescribeDirectoriesResponseBodyDirectoriesADConnectors {
	s.ADConnectorAddress = &v
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

type DescribeImagesRequest struct {
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults  *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ImageType   *string   `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageStatus *string   `json:"ImageStatus,omitempty" xml:"ImageStatus,omitempty"`
	GpuCategory *bool     `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
	ImageId     []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
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
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ImageType    *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Size         *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	OsType       *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	DataDiskSize *int32  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	GpuCategory  *bool   `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s DescribeImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImages) SetStatus(v string) *DescribeImagesResponseBodyImages {
	s.Status = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetCreationTime(v string) *DescribeImagesResponseBodyImages {
	s.CreationTime = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetImageType(v string) *DescribeImagesResponseBodyImages {
	s.ImageType = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetProgress(v string) *DescribeImagesResponseBodyImages {
	s.Progress = &v
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

func (s *DescribeImagesResponseBodyImages) SetDataDiskSize(v int32) *DescribeImagesResponseBodyImages {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetName(v string) *DescribeImagesResponseBodyImages {
	s.Name = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetGpuCategory(v bool) *DescribeImagesResponseBodyImages {
	s.GpuCategory = &v
	return s
}

func (s *DescribeImagesResponseBodyImages) SetImageId(v string) *DescribeImagesResponseBodyImages {
	s.ImageId = &v
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
	InvokeDesktops   []*DescribeInvocationsResponseBodyInvocationsInvokeDesktops `json:"InvokeDesktops,omitempty" xml:"InvokeDesktops,omitempty" type:"Repeated"`
	CommandContent   *string                                                     `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	InvokeId         *string                                                     `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	CommandType      *string                                                     `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
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

func (s *DescribeInvocationsResponseBodyInvocations) SetInvokeDesktops(v []*DescribeInvocationsResponseBodyInvocationsInvokeDesktops) *DescribeInvocationsResponseBodyInvocations {
	s.InvokeDesktops = v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCommandContent(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CommandContent = &v
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

type DescribeInvocationsResponseBodyInvocationsInvokeDesktops struct {
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	UpdateTime       *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	FinishTime       *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	Repeats          *int32  `json:"Repeats,omitempty" xml:"Repeats,omitempty"`
	Output           *string `json:"Output,omitempty" xml:"Output,omitempty"`
	DesktopId        *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
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

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetUpdateTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.UpdateTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetFinishTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.FinishTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetRepeats(v int32) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.Repeats = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetOutput(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.Output = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetDesktopId(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.DesktopId = &v
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
	PolicyStatus          *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	WatermarkType         *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	PolicyGroupId         *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	WatermarkTransparency *string `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
	UsbRedirect           *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	PolicyGroupType       *string `json:"PolicyGroupType,omitempty" xml:"PolicyGroupType,omitempty"`
	WatermarkCustomText   *string `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	Watermark             *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	Clipboard             *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	LocalDrive            *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
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

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkType = &v
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

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetUsbRedirect(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.UsbRedirect = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetPolicyGroupType(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.PolicyGroupType = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermarkCustomText(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.WatermarkCustomText = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetWatermark(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Watermark = &v
	return s
}

func (s *DescribePolicyGroupsResponseBodyDescribePolicyGroups) SetClipboard(v string) *DescribePolicyGroupsResponseBodyDescribePolicyGroups {
	s.Clipboard = &v
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

type DescribeVirtualMFADevicesRequest struct {
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MaxResults  *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	DirectoryId *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId   []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
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

func (s *DescribeVirtualMFADevicesRequest) SetEndUserId(v []*string) *DescribeVirtualMFADevicesRequest {
	s.EndUserId = v
	return s
}

type DescribeVirtualMFADevicesResponseBody struct {
	VirtualMFADevices []*DescribeVirtualMFADevicesResponseBodyVirtualMFADevices `json:"VirtualMFADevices,omitempty" xml:"VirtualMFADevices,omitempty" type:"Repeated"`
	NextToken         *string                                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVirtualMFADevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualMFADevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesResponseBody) SetVirtualMFADevices(v []*DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) *DescribeVirtualMFADevicesResponseBody {
	s.VirtualMFADevices = v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBody) SetNextToken(v string) *DescribeVirtualMFADevicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBody) SetRequestId(v string) *DescribeVirtualMFADevicesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVirtualMFADevicesResponseBodyVirtualMFADevices struct {
	GmtUnlock        *string `json:"GmtUnlock,omitempty" xml:"GmtUnlock,omitempty"`
	SerialNumber     *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	EndUserId        *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ConsecutiveFails *int32  `json:"ConsecutiveFails,omitempty" xml:"ConsecutiveFails,omitempty"`
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

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetGmtUnlock(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.GmtUnlock = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetSerialNumber(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.SerialNumber = &v
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

type ListTagResourcesRequest struct {
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
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
	Message   *int32  `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *ModifyDesktopsPolicyGroupResponseBodyModifyResults) SetMessage(v int32) *ModifyDesktopsPolicyGroupResponseBodyModifyResults {
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

type ModifyPolicyGroupRequest struct {
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PolicyGroupId         *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Clipboard             *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	LocalDrive            *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UsbRedirect           *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	Watermark             *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	WatermarkType         *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
	WatermarkCustomText   *string `json:"WatermarkCustomText,omitempty" xml:"WatermarkCustomText,omitempty"`
	WatermarkTransparency *string `json:"WatermarkTransparency,omitempty" xml:"WatermarkTransparency,omitempty"`
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RenewDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *RenewDesktopsResponseBody) SetRequestId(v string) *RenewDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewDesktopsResponseBody) SetOrderId(v string) *RenewDesktopsResponseBody {
	s.OrderId = &v
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InvokeId  *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
}

func (s RunCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunCommandResponseBody) SetRequestId(v string) *RunCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCommandResponseBody) SetInvokeId(v string) *RunCommandResponseBody {
	s.InvokeId = &v
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

type StopDesktopsRequest struct {
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
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
