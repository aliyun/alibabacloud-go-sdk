// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddIpRequest struct {
	// The ID of the Anti-DDoS Origin instance.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-npk1z7t9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP addresses that you want to add to the Anti-DDoS Origin instance. This parameter is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that includes the following field:
	//
	// 	- **ip**: required. The IP address that you want to add. Data type: string.
	//
	//     **
	//
	//     **Note*	- The IP address must be the IP address of an asset that belongs to the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"ip":"1.XX.XX.1"},{"ip":"2.XX.XX.2"}]
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. This parameter is empty by default, which indicates that the Anti-DDoS Origin instance belongs to the default resource group.
	//
	// For information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s AddIpRequest) String() string {
	return tea.Prettify(s)
}

func (s AddIpRequest) GoString() string {
	return s.String()
}

func (s *AddIpRequest) SetInstanceId(v string) *AddIpRequest {
	s.InstanceId = &v
	return s
}

func (s *AddIpRequest) SetIpList(v string) *AddIpRequest {
	s.IpList = &v
	return s
}

func (s *AddIpRequest) SetRegionId(v string) *AddIpRequest {
	s.RegionId = &v
	return s
}

func (s *AddIpRequest) SetResourceGroupId(v string) *AddIpRequest {
	s.ResourceGroupId = &v
	return s
}

type AddIpResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddIpResponseBody) GoString() string {
	return s.String()
}

func (s *AddIpResponseBody) SetRequestId(v string) *AddIpResponseBody {
	s.RequestId = &v
	return s
}

type AddIpResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddIpResponse) String() string {
	return tea.Prettify(s)
}

func (s AddIpResponse) GoString() string {
	return s.String()
}

func (s *AddIpResponse) SetHeaders(v map[string]*string) *AddIpResponse {
	s.Headers = v
	return s
}

func (s *AddIpResponse) SetStatusCode(v int32) *AddIpResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIpResponse) SetBody(v *AddIpResponseBody) *AddIpResponse {
	s.Body = v
	return s
}

type AddRdMemberListRequest struct {
	// The list of the members.
	//
	// This parameter is required.
	MemberList []*AddRdMemberListRequestMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
}

func (s AddRdMemberListRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRdMemberListRequest) GoString() string {
	return s.String()
}

func (s *AddRdMemberListRequest) SetMemberList(v []*AddRdMemberListRequestMemberList) *AddRdMemberListRequest {
	s.MemberList = v
	return s
}

type AddRdMemberListRequestMemberList struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 19510843762****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s AddRdMemberListRequestMemberList) String() string {
	return tea.Prettify(s)
}

func (s AddRdMemberListRequestMemberList) GoString() string {
	return s.String()
}

func (s *AddRdMemberListRequestMemberList) SetUid(v string) *AddRdMemberListRequestMemberList {
	s.Uid = &v
	return s
}

type AddRdMemberListShrinkRequest struct {
	// The list of the members.
	//
	// This parameter is required.
	MemberListShrink *string `json:"MemberList,omitempty" xml:"MemberList,omitempty"`
}

func (s AddRdMemberListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRdMemberListShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddRdMemberListShrinkRequest) SetMemberListShrink(v string) *AddRdMemberListShrinkRequest {
	s.MemberListShrink = &v
	return s
}

type AddRdMemberListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddRdMemberListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRdMemberListResponseBody) GoString() string {
	return s.String()
}

func (s *AddRdMemberListResponseBody) SetRequestId(v string) *AddRdMemberListResponseBody {
	s.RequestId = &v
	return s
}

type AddRdMemberListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRdMemberListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRdMemberListResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRdMemberListResponse) GoString() string {
	return s.String()
}

func (s *AddRdMemberListResponse) SetHeaders(v map[string]*string) *AddRdMemberListResponse {
	s.Headers = v
	return s
}

func (s *AddRdMemberListResponse) SetStatusCode(v int32) *AddRdMemberListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRdMemberListResponse) SetBody(v *AddRdMemberListResponseBody) *AddRdMemberListResponse {
	s.Body = v
	return s
}

type AttachAssetGroupToInstanceRequest struct {
	// The information about the asset to be associated.
	//
	// This parameter is required.
	AssetGroupList []*AttachAssetGroupToInstanceRequestAssetGroupList `json:"AssetGroupList,omitempty" xml:"AssetGroupList,omitempty" type:"Repeated"`
	// The ID of the instance to query.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances of paid editions.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachAssetGroupToInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachAssetGroupToInstanceRequest) GoString() string {
	return s.String()
}

func (s *AttachAssetGroupToInstanceRequest) SetAssetGroupList(v []*AttachAssetGroupToInstanceRequestAssetGroupList) *AttachAssetGroupToInstanceRequest {
	s.AssetGroupList = v
	return s
}

func (s *AttachAssetGroupToInstanceRequest) SetInstanceId(v string) *AttachAssetGroupToInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachAssetGroupToInstanceRequest) SetRegionId(v string) *AttachAssetGroupToInstanceRequest {
	s.RegionId = &v
	return s
}

type AttachAssetGroupToInstanceRequestAssetGroupList struct {
	// The UID of the member to which the asset belongs.
	//
	// example:
	//
	// 1743970208320***
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the asset that you want to add. If the asset is a Web Application Firewall (WAF) instance, specify the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-test-001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The type of the asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AttachAssetGroupToInstanceRequestAssetGroupList) String() string {
	return tea.Prettify(s)
}

func (s AttachAssetGroupToInstanceRequestAssetGroupList) GoString() string {
	return s.String()
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) SetMemberUid(v string) *AttachAssetGroupToInstanceRequestAssetGroupList {
	s.MemberUid = &v
	return s
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) SetName(v string) *AttachAssetGroupToInstanceRequestAssetGroupList {
	s.Name = &v
	return s
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) SetRegion(v string) *AttachAssetGroupToInstanceRequestAssetGroupList {
	s.Region = &v
	return s
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) SetType(v string) *AttachAssetGroupToInstanceRequestAssetGroupList {
	s.Type = &v
	return s
}

type AttachAssetGroupToInstanceShrinkRequest struct {
	// The information about the asset to be associated.
	//
	// This parameter is required.
	AssetGroupListShrink *string `json:"AssetGroupList,omitempty" xml:"AssetGroupList,omitempty"`
	// The ID of the instance to query.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances of paid editions.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachAssetGroupToInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachAssetGroupToInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachAssetGroupToInstanceShrinkRequest) SetAssetGroupListShrink(v string) *AttachAssetGroupToInstanceShrinkRequest {
	s.AssetGroupListShrink = &v
	return s
}

func (s *AttachAssetGroupToInstanceShrinkRequest) SetInstanceId(v string) *AttachAssetGroupToInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachAssetGroupToInstanceShrinkRequest) SetRegionId(v string) *AttachAssetGroupToInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

type AttachAssetGroupToInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 52B49F64-5A36-5CE0-BD00-765792C26AA9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachAssetGroupToInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachAssetGroupToInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AttachAssetGroupToInstanceResponseBody) SetRequestId(v string) *AttachAssetGroupToInstanceResponseBody {
	s.RequestId = &v
	return s
}

type AttachAssetGroupToInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachAssetGroupToInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachAssetGroupToInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachAssetGroupToInstanceResponse) GoString() string {
	return s.String()
}

func (s *AttachAssetGroupToInstanceResponse) SetHeaders(v map[string]*string) *AttachAssetGroupToInstanceResponse {
	s.Headers = v
	return s
}

func (s *AttachAssetGroupToInstanceResponse) SetStatusCode(v int32) *AttachAssetGroupToInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachAssetGroupToInstanceResponse) SetBody(v *AttachAssetGroupToInstanceResponseBody) *AttachAssetGroupToInstanceResponse {
	s.Body = v
	return s
}

type AttachToPolicyRequest struct {
	// The protected objects.
	//
	// This parameter is required.
	IpPortProtocolList []*AttachToPolicyRequestIpPortProtocolList `json:"IpPortProtocolList,omitempty" xml:"IpPortProtocolList,omitempty" type:"Repeated"`
	// The policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cd8b4d70-e4e0-413a-b390-e71d********
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s AttachToPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachToPolicyRequest) GoString() string {
	return s.String()
}

func (s *AttachToPolicyRequest) SetIpPortProtocolList(v []*AttachToPolicyRequestIpPortProtocolList) *AttachToPolicyRequest {
	s.IpPortProtocolList = v
	return s
}

func (s *AttachToPolicyRequest) SetPolicyId(v string) *AttachToPolicyRequest {
	s.PolicyId = &v
	return s
}

type AttachToPolicyRequestIpPortProtocolList struct {
	// The IP address of the protected object.
	//
	// This parameter is required.
	//
	// example:
	//
	// 112.124.241.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port number of the protected object.
	//
	// >  This parameter is available for only port-specific mitigation policies.
	//
	// example:
	//
	// 8*
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type of the protected object. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// >  This parameter is available for only port-specific mitigation policies.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s AttachToPolicyRequestIpPortProtocolList) String() string {
	return tea.Prettify(s)
}

func (s AttachToPolicyRequestIpPortProtocolList) GoString() string {
	return s.String()
}

func (s *AttachToPolicyRequestIpPortProtocolList) SetIp(v string) *AttachToPolicyRequestIpPortProtocolList {
	s.Ip = &v
	return s
}

func (s *AttachToPolicyRequestIpPortProtocolList) SetPort(v int32) *AttachToPolicyRequestIpPortProtocolList {
	s.Port = &v
	return s
}

func (s *AttachToPolicyRequestIpPortProtocolList) SetProtocol(v string) *AttachToPolicyRequestIpPortProtocolList {
	s.Protocol = &v
	return s
}

type AttachToPolicyShrinkRequest struct {
	// The protected objects.
	//
	// This parameter is required.
	IpPortProtocolListShrink *string `json:"IpPortProtocolList,omitempty" xml:"IpPortProtocolList,omitempty"`
	// The policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cd8b4d70-e4e0-413a-b390-e71d********
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s AttachToPolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachToPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachToPolicyShrinkRequest) SetIpPortProtocolListShrink(v string) *AttachToPolicyShrinkRequest {
	s.IpPortProtocolListShrink = &v
	return s
}

func (s *AttachToPolicyShrinkRequest) SetPolicyId(v string) *AttachToPolicyShrinkRequest {
	s.PolicyId = &v
	return s
}

type AttachToPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DC245DEE-9800-5579-BF99-189D6A5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachToPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachToPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AttachToPolicyResponseBody) SetRequestId(v string) *AttachToPolicyResponseBody {
	s.RequestId = &v
	return s
}

type AttachToPolicyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachToPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachToPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachToPolicyResponse) GoString() string {
	return s.String()
}

func (s *AttachToPolicyResponse) SetHeaders(v map[string]*string) *AttachToPolicyResponse {
	s.Headers = v
	return s
}

func (s *AttachToPolicyResponse) SetStatusCode(v int32) *AttachToPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachToPolicyResponse) SetBody(v *AttachToPolicyResponseBody) *AttachToPolicyResponse {
	s.Body = v
	return s
}

type CheckAccessLogAuthRequest struct {
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// For more information about the valid values of this parameter, see [Regions and zones](https://help.aliyun.com/document_detail/188196.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. This parameter is empty by default, which indicates that the Anti-DDoS Origin instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CheckAccessLogAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAccessLogAuthRequest) GoString() string {
	return s.String()
}

func (s *CheckAccessLogAuthRequest) SetRegionId(v string) *CheckAccessLogAuthRequest {
	s.RegionId = &v
	return s
}

func (s *CheckAccessLogAuthRequest) SetResourceGroupId(v string) *CheckAccessLogAuthRequest {
	s.ResourceGroupId = &v
	return s
}

type CheckAccessLogAuthResponseBody struct {
	// Indicates whether Anti-DDoS Origin was authorized to access Simple Log Service. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AccessLogAuth *bool `json:"AccessLogAuth,omitempty" xml:"AccessLogAuth,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 864FE2F4-CB2E-4024-B9EF-D59FD08ABD41
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckAccessLogAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckAccessLogAuthResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAccessLogAuthResponseBody) SetAccessLogAuth(v bool) *CheckAccessLogAuthResponseBody {
	s.AccessLogAuth = &v
	return s
}

func (s *CheckAccessLogAuthResponseBody) SetRequestId(v string) *CheckAccessLogAuthResponseBody {
	s.RequestId = &v
	return s
}

type CheckAccessLogAuthResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAccessLogAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAccessLogAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAccessLogAuthResponse) GoString() string {
	return s.String()
}

func (s *CheckAccessLogAuthResponse) SetHeaders(v map[string]*string) *CheckAccessLogAuthResponse {
	s.Headers = v
	return s
}

func (s *CheckAccessLogAuthResponse) SetStatusCode(v int32) *CheckAccessLogAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAccessLogAuthResponse) SetBody(v *CheckAccessLogAuthResponseBody) *CheckAccessLogAuthResponse {
	s.Body = v
	return s
}

type CheckGrantRequest struct {
	// Specifies whether to allow Anti-DDoS Origin to check the service-linked role. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsSlr *bool `json:"IsSlr,omitempty" xml:"IsSlr,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CheckGrantRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantRequest) GoString() string {
	return s.String()
}

func (s *CheckGrantRequest) SetIsSlr(v bool) *CheckGrantRequest {
	s.IsSlr = &v
	return s
}

func (s *CheckGrantRequest) SetRegionId(v string) *CheckGrantRequest {
	s.RegionId = &v
	return s
}

func (s *CheckGrantRequest) SetResourceGroupId(v string) *CheckGrantRequest {
	s.ResourceGroupId = &v
	return s
}

type CheckGrantResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DB002CE5-5E6C-5F11-AE15-B525299A40F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether Anti-DDoS Origin is authorized to obtain information about the assets within the current Alibaba Cloud account. Valid values:
	//
	// 	- **1**: Anti-DDoS Origin is authorized to obtain information about the assets within the current Alibaba Cloud account.
	//
	// 	- **0**: Anti-DDoS Origin is not authorized to obtain information about the assets within the current Alibaba Cloud account.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckGrantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantResponseBody) GoString() string {
	return s.String()
}

func (s *CheckGrantResponseBody) SetRequestId(v string) *CheckGrantResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckGrantResponseBody) SetStatus(v int32) *CheckGrantResponseBody {
	s.Status = &v
	return s
}

type CheckGrantResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckGrantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckGrantResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantResponse) GoString() string {
	return s.String()
}

func (s *CheckGrantResponse) SetHeaders(v map[string]*string) *CheckGrantResponse {
	s.Headers = v
	return s
}

func (s *CheckGrantResponse) SetStatusCode(v int32) *CheckGrantResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckGrantResponse) SetBody(v *CheckGrantResponseBody) *CheckGrantResponse {
	s.Body = v
	return s
}

type ConfigSchedruleOnDemandRequest struct {
	// The ID of the on-demand instance.
	//
	// >  You can call the [DescribeOnDemandInstance](https://help.aliyun.com/document_detail/152120.html) operation to query the IDs of all on-demand instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-z2q1qzxb****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the on-demand instance.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scheduling action. Set the value to **declare**, which indicates that the route is advertised.
	//
	// This parameter is required.
	//
	// example:
	//
	// declare
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// If the inbound bandwidth or packets consecutively exceed the threshold for the specified number of times, the scheduling rule is triggered and traffic is rerouted to the on-demand instance. The specified number of times is the value of this parameter.
	//
	// >  The threshold of inbound bandwidth is the value of **RuleConditionMbps**. The threshold of inbound packets is the value of **RuleConditionKpps**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	RuleConditionCnt *string `json:"RuleConditionCnt,omitempty" xml:"RuleConditionCnt,omitempty"`
	// The threshold of inbound packets. Unit: kilo packets per second (Kpps). Minimum value: **10**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	RuleConditionKpps *string `json:"RuleConditionKpps,omitempty" xml:"RuleConditionKpps,omitempty"`
	// The threshold of inbound bandwidth. Unit: Mbit/s. Minimum value: **100**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	RuleConditionMbps *string `json:"RuleConditionMbps,omitempty" xml:"RuleConditionMbps,omitempty"`
	// The name of the scheduling rule.
	//
	// The name can contain lowercase letters, digits, hyphens (-), and underscores (_). The name can be up to 32 characters in length. We recommend that you use the ID of the on-demand instance as the name. Example: `ddosbgp-cn-z2q1qzxb****`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-z2q1qzxb****
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Specifies whether the scheduling rule is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	RuleSwitch *string `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	// The start time of the period during which the scheduling rule is automatically stopped. The time must be in the 24-hour clock and in the `hh:mm` format.
	//
	// If the system detects that DDoS attacks stop, the system no longer reroutes traffic to the on-demand instance from the time you specified. We recommend that you set this parameter to a value that is defined as off-peak hours.
	//
	// >  This parameter takes effect only when the value of **RuleUndoMode*	- is **auto**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 03:00
	RuleUndoBeginTime *string `json:"RuleUndoBeginTime,omitempty" xml:"RuleUndoBeginTime,omitempty"`
	// The end time of the period during which the scheduling rule is automatically stopped. The time must be in the 24-hour clock and in the `hh:mm` format.
	//
	// example:
	//
	// 03:05
	RuleUndoEndTime *string `json:"RuleUndoEndTime,omitempty" xml:"RuleUndoEndTime,omitempty"`
	// The stop method of the scheduling rule. Valid values:
	//
	// 	- **auto**
	//
	// 	- **manual**
	//
	// This parameter is required.
	//
	// example:
	//
	// manual
	RuleUndoMode *string `json:"RuleUndoMode,omitempty" xml:"RuleUndoMode,omitempty"`
	// The time zone of the time when the scheduling rule automatically stops. The time zone must be in the `GMT-hh:mm` format.
	//
	// For example, the value `GMT-08:00` indicates that the time zone is UTC+8.
	//
	// >  This parameter takes effect only when the value of **RuleUndoMode*	- is **auto**.
	//
	// This parameter is required.
	//
	// example:
	//
	// GMT-08:00
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s ConfigSchedruleOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigSchedruleOnDemandRequest) GoString() string {
	return s.String()
}

func (s *ConfigSchedruleOnDemandRequest) SetInstanceId(v string) *ConfigSchedruleOnDemandRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRegionId(v string) *ConfigSchedruleOnDemandRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleAction(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleAction = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleConditionCnt(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleConditionCnt = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleConditionKpps(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleConditionKpps = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleConditionMbps(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleConditionMbps = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleName(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleName = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleSwitch(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleSwitch = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleUndoBeginTime(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleUndoBeginTime = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleUndoEndTime(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleUndoEndTime = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleUndoMode(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleUndoMode = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetTimeZone(v string) *ConfigSchedruleOnDemandRequest {
	s.TimeZone = &v
	return s
}

type ConfigSchedruleOnDemandResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD06F539-2FBE-450D-9391-7EFF787128F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigSchedruleOnDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigSchedruleOnDemandResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigSchedruleOnDemandResponseBody) SetRequestId(v string) *ConfigSchedruleOnDemandResponseBody {
	s.RequestId = &v
	return s
}

type ConfigSchedruleOnDemandResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigSchedruleOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigSchedruleOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigSchedruleOnDemandResponse) GoString() string {
	return s.String()
}

func (s *ConfigSchedruleOnDemandResponse) SetHeaders(v map[string]*string) *ConfigSchedruleOnDemandResponse {
	s.Headers = v
	return s
}

func (s *ConfigSchedruleOnDemandResponse) SetStatusCode(v int32) *ConfigSchedruleOnDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigSchedruleOnDemandResponse) SetBody(v *ConfigSchedruleOnDemandResponseBody) *ConfigSchedruleOnDemandResponse {
	s.Body = v
	return s
}

type CreatePolicyRequest struct {
	// The name of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **l3**: IP-specific mitigation policies.
	//
	// 	- **l4**: port-specific mitigation policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// l3
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequest) SetName(v string) *CreatePolicyRequest {
	s.Name = &v
	return s
}

func (s *CreatePolicyRequest) SetType(v string) *CreatePolicyRequest {
	s.Type = &v
	return s
}

type CreatePolicyResponseBody struct {
	// The ID of the policy.
	//
	// example:
	//
	// 83967609-7ea5-4f6d-a6ea-380b09e****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 864FE2F4-CB2E-4024-B9EF-D59FD08A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBody) SetId(v string) *CreatePolicyResponseBody {
	s.Id = &v
	return s
}

func (s *CreatePolicyResponseBody) SetRequestId(v string) *CreatePolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreatePolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponse) SetHeaders(v map[string]*string) *CreatePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicyResponse) SetStatusCode(v int32) *CreatePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolicyResponse) SetBody(v *CreatePolicyResponseBody) *CreatePolicyResponse {
	s.Body = v
	return s
}

type CreateSchedruleOnDemandRequest struct {
	// The ID of the on-demand instance.
	//
	// >  You can call the [DescribeOnDemandInstance](https://help.aliyun.com/document_detail/152120.html) operation to query the IDs of all on-demand instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-z2q1qzxb****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the on-demand instance.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scheduling action. Set the value to **declare**, which indicates that the route is advertised.
	//
	// This parameter is required.
	//
	// example:
	//
	// declare
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// If the inbound bandwidth or packets consecutively exceed the threshold for the specified number of times, the scheduling rule is triggered and traffic is rerouted to the on-demand instance. The specified number of times is the value of this parameter.
	//
	// >  The threshold of inbound bandwidth is the value of **RuleConditionMbps**. The threshold of inbound packets is the value of **RuleConditionKpps**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	RuleConditionCnt *string `json:"RuleConditionCnt,omitempty" xml:"RuleConditionCnt,omitempty"`
	// The threshold of inbound packets. Unit: kilo packets per second (Kpps). Minimum value: **10**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	RuleConditionKpps *string `json:"RuleConditionKpps,omitempty" xml:"RuleConditionKpps,omitempty"`
	// The threshold of inbound bandwidth. Unit: Mbit/s. Minimum value: **100**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	RuleConditionMbps *string `json:"RuleConditionMbps,omitempty" xml:"RuleConditionMbps,omitempty"`
	// The name of the scheduling rule.
	//
	// The name can contain lowercase letters, digits, hyphens (-), and underscores (_). The name can be up to 32 characters in length. We recommend that you use the ID of the on-demand instance as the name. Example: `ddosbgp-cn-z2q1qzxb****`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-z2q1qzxb****
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Specifies whether the scheduling rule is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	RuleSwitch *string `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	// The start time of the period during which the scheduling rule is automatically stopped. The time must be in the 24-hour clock and in the `hh:mm` format.
	//
	// If the system detects that DDoS attacks stop, the system no longer reroutes traffic to the on-demand instance from the time you specified. We recommend that you set this parameter to a value that is defined as off-peak hours.
	//
	// >  This parameter takes effect only when the value of **RuleUndoMode*	- is **auto**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 03:00
	RuleUndoBeginTime *string `json:"RuleUndoBeginTime,omitempty" xml:"RuleUndoBeginTime,omitempty"`
	// The end time of the period during which the scheduling rule is automatically stopped. The time must be in the 24-hour clock and in the `hh:mm` format.
	//
	// example:
	//
	// 03:05
	RuleUndoEndTime *string `json:"RuleUndoEndTime,omitempty" xml:"RuleUndoEndTime,omitempty"`
	// The stop method of the scheduling rule. Valid values:
	//
	// 	- **auto**
	//
	// 	- **manual**
	//
	// This parameter is required.
	//
	// example:
	//
	// auto
	RuleUndoMode *string `json:"RuleUndoMode,omitempty" xml:"RuleUndoMode,omitempty"`
	// The time zone of the time when the scheduling rule automatically stops. The time zone must be in the `GMT-hh:mm` format.
	//
	// For example, the value `GMT-08:00` indicates that the time zone is UTC+8.
	//
	// >  This parameter takes effect only when the value of **RuleUndoMode*	- is **auto**.
	//
	// This parameter is required.
	//
	// example:
	//
	// GMT-08:00
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s CreateSchedruleOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSchedruleOnDemandRequest) GoString() string {
	return s.String()
}

func (s *CreateSchedruleOnDemandRequest) SetInstanceId(v string) *CreateSchedruleOnDemandRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRegionId(v string) *CreateSchedruleOnDemandRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleAction(v string) *CreateSchedruleOnDemandRequest {
	s.RuleAction = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleConditionCnt(v string) *CreateSchedruleOnDemandRequest {
	s.RuleConditionCnt = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleConditionKpps(v string) *CreateSchedruleOnDemandRequest {
	s.RuleConditionKpps = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleConditionMbps(v string) *CreateSchedruleOnDemandRequest {
	s.RuleConditionMbps = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleName(v string) *CreateSchedruleOnDemandRequest {
	s.RuleName = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleSwitch(v string) *CreateSchedruleOnDemandRequest {
	s.RuleSwitch = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleUndoBeginTime(v string) *CreateSchedruleOnDemandRequest {
	s.RuleUndoBeginTime = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleUndoEndTime(v string) *CreateSchedruleOnDemandRequest {
	s.RuleUndoEndTime = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleUndoMode(v string) *CreateSchedruleOnDemandRequest {
	s.RuleUndoMode = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetTimeZone(v string) *CreateSchedruleOnDemandRequest {
	s.TimeZone = &v
	return s
}

type CreateSchedruleOnDemandResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD06F539-2FBE-450D-9391-7EFF787128F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSchedruleOnDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSchedruleOnDemandResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSchedruleOnDemandResponseBody) SetRequestId(v string) *CreateSchedruleOnDemandResponseBody {
	s.RequestId = &v
	return s
}

type CreateSchedruleOnDemandResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSchedruleOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSchedruleOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSchedruleOnDemandResponse) GoString() string {
	return s.String()
}

func (s *CreateSchedruleOnDemandResponse) SetHeaders(v map[string]*string) *CreateSchedruleOnDemandResponse {
	s.Headers = v
	return s
}

func (s *CreateSchedruleOnDemandResponse) SetStatusCode(v int32) *CreateSchedruleOnDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSchedruleOnDemandResponse) SetBody(v *CreateSchedruleOnDemandResponseBody) *CreateSchedruleOnDemandResponse {
	s.Body = v
	return s
}

type DeleteBlackholeRequest struct {
	// The ID of the Anti-DDoS Origin instance.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address for which you want to deactivate blackhole filtering.
	//
	// >  You can call the [DescribePackIpList](https://help.aliyun.com/document_detail/118701.html) operation to query all the IP addresses that are protected by the Anti-DDoS Origin instance and the protection status of the IP addresses. For example, you can query whether blackhole filtering is triggered for an IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47.89.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteBlackholeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeRequest) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeRequest) SetInstanceId(v string) *DeleteBlackholeRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteBlackholeRequest) SetIp(v string) *DeleteBlackholeRequest {
	s.Ip = &v
	return s
}

func (s *DeleteBlackholeRequest) SetRegionId(v string) *DeleteBlackholeRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBlackholeRequest) SetResourceGroupId(v string) *DeleteBlackholeRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteBlackholeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBlackholeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeResponseBody) SetRequestId(v string) *DeleteBlackholeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBlackholeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBlackholeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBlackholeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeResponse) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeResponse) SetHeaders(v map[string]*string) *DeleteBlackholeResponse {
	s.Headers = v
	return s
}

func (s *DeleteBlackholeResponse) SetStatusCode(v int32) *DeleteBlackholeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBlackholeResponse) SetBody(v *DeleteBlackholeResponseBody) *DeleteBlackholeResponse {
	s.Body = v
	return s
}

type DeleteIpRequest struct {
	// The ID of the Anti-DDoS Origin instance.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-npk1z7t9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP addresses that you want to remove from the Anti-DDoS Origin instance. This parameter is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that contains the following fields:
	//
	// 	- **ip**: required. The IP address that you want to remove. Data type: string.
	//
	//     **
	//
	//     **Note*	- The IP addresses that you want to remove must be protected by the Anti-DDoS Origin instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"ip":"1.XX.XX.1"},{"ip":"2.XX.XX.2"}]
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. This parameter is empty by default, which indicates that the Anti-DDoS Origin instance belongs to the default resource group.
	//
	// For information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpRequest) SetInstanceId(v string) *DeleteIpRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteIpRequest) SetIpList(v string) *DeleteIpRequest {
	s.IpList = &v
	return s
}

func (s *DeleteIpRequest) SetRegionId(v string) *DeleteIpRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpRequest) SetResourceGroupId(v string) *DeleteIpRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteIpResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpResponseBody) SetRequestId(v string) *DeleteIpResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIpResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpResponse) SetHeaders(v map[string]*string) *DeleteIpResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpResponse) SetStatusCode(v int32) *DeleteIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpResponse) SetBody(v *DeleteIpResponseBody) *DeleteIpResponse {
	s.Body = v
	return s
}

type DeletePolicyRequest struct {
	// The ID of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 90300b1a-ced8-4437-b4bf-f9a5*******
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeletePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyRequest) SetId(v string) *DeletePolicyRequest {
	s.Id = &v
	return s
}

type DeletePolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD06F539-2FBE-450D-9391-7EFF7871****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyResponseBody) SetRequestId(v string) *DeletePolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeletePolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyResponse) SetHeaders(v map[string]*string) *DeletePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyResponse) SetStatusCode(v int32) *DeletePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyResponse) SetBody(v *DeletePolicyResponseBody) *DeletePolicyResponse {
	s.Body = v
	return s
}

type DeleteRdMemberListRequest struct {
	// The list of the members.
	//
	// This parameter is required.
	MemberList []*DeleteRdMemberListRequestMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
}

func (s DeleteRdMemberListRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRdMemberListRequest) GoString() string {
	return s.String()
}

func (s *DeleteRdMemberListRequest) SetMemberList(v []*DeleteRdMemberListRequestMemberList) *DeleteRdMemberListRequest {
	s.MemberList = v
	return s
}

type DeleteRdMemberListRequestMemberList struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 136548010379****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DeleteRdMemberListRequestMemberList) String() string {
	return tea.Prettify(s)
}

func (s DeleteRdMemberListRequestMemberList) GoString() string {
	return s.String()
}

func (s *DeleteRdMemberListRequestMemberList) SetUid(v string) *DeleteRdMemberListRequestMemberList {
	s.Uid = &v
	return s
}

type DeleteRdMemberListShrinkRequest struct {
	// The list of the members.
	//
	// This parameter is required.
	MemberListShrink *string `json:"MemberList,omitempty" xml:"MemberList,omitempty"`
}

func (s DeleteRdMemberListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRdMemberListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteRdMemberListShrinkRequest) SetMemberListShrink(v string) *DeleteRdMemberListShrinkRequest {
	s.MemberListShrink = &v
	return s
}

type DeleteRdMemberListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A2D6D5FB-FA07-41A8-B093-A2B7B26E72F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRdMemberListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRdMemberListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRdMemberListResponseBody) SetRequestId(v string) *DeleteRdMemberListResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRdMemberListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRdMemberListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRdMemberListResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRdMemberListResponse) GoString() string {
	return s.String()
}

func (s *DeleteRdMemberListResponse) SetHeaders(v map[string]*string) *DeleteRdMemberListResponse {
	s.Headers = v
	return s
}

func (s *DeleteRdMemberListResponse) SetStatusCode(v int32) *DeleteRdMemberListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRdMemberListResponse) SetBody(v *DeleteRdMemberListResponseBody) *DeleteRdMemberListResponse {
	s.Body = v
	return s
}

type DeleteSchedruleOnDemandRequest struct {
	// The ID of the anti-DDoS diversion instance.
	//
	// >  You can call the [DescribeOnDemandInstance](https://help.aliyun.com/document_detail/152120.html) operation to query the IDs of all anti-DDoS diversion instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-z2q1qzxb****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the anti-DDoS diversion instance.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the scheduling rule that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// testrule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DeleteSchedruleOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchedruleOnDemandRequest) GoString() string {
	return s.String()
}

func (s *DeleteSchedruleOnDemandRequest) SetInstanceId(v string) *DeleteSchedruleOnDemandRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSchedruleOnDemandRequest) SetRegionId(v string) *DeleteSchedruleOnDemandRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSchedruleOnDemandRequest) SetRuleName(v string) *DeleteSchedruleOnDemandRequest {
	s.RuleName = &v
	return s
}

type DeleteSchedruleOnDemandResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD06F539-2FBE-450D-9391-7EFF787128F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSchedruleOnDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchedruleOnDemandResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSchedruleOnDemandResponseBody) SetRequestId(v string) *DeleteSchedruleOnDemandResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSchedruleOnDemandResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSchedruleOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSchedruleOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchedruleOnDemandResponse) GoString() string {
	return s.String()
}

func (s *DeleteSchedruleOnDemandResponse) SetHeaders(v map[string]*string) *DeleteSchedruleOnDemandResponse {
	s.Headers = v
	return s
}

func (s *DeleteSchedruleOnDemandResponse) SetStatusCode(v int32) *DeleteSchedruleOnDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSchedruleOnDemandResponse) SetBody(v *DeleteSchedruleOnDemandResponseBody) *DeleteSchedruleOnDemandResponse {
	s.Body = v
	return s
}

type DescribeAssetGroupRequest struct {
	// The ID of the asset. If the asset is a Web Application Firewall (WAF) instance, specify the ID of the WAF instance.
	//
	// example:
	//
	// waf_v2_public_cn-lbj382l****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- **waf**: WAF instance
	//
	// 	- **ga**: Global Accelerator (GA) instance
	//
	// This parameter is required.
	//
	// example:
	//
	// waf
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAssetGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupRequest) SetName(v string) *DescribeAssetGroupRequest {
	s.Name = &v
	return s
}

func (s *DescribeAssetGroupRequest) SetRegion(v string) *DescribeAssetGroupRequest {
	s.Region = &v
	return s
}

func (s *DescribeAssetGroupRequest) SetRegionId(v string) *DescribeAssetGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAssetGroupRequest) SetType(v string) *DescribeAssetGroupRequest {
	s.Type = &v
	return s
}

type DescribeAssetGroupResponseBody struct {
	// The information about the asset.
	AssetGroupList []*DescribeAssetGroupResponseBodyAssetGroupList `json:"AssetGroupList,omitempty" xml:"AssetGroupList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 487EC0F7-8D14-504E-914E-3A1BC314B581
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAssetGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupResponseBody) SetAssetGroupList(v []*DescribeAssetGroupResponseBodyAssetGroupList) *DescribeAssetGroupResponseBody {
	s.AssetGroupList = v
	return s
}

func (s *DescribeAssetGroupResponseBody) SetRequestId(v string) *DescribeAssetGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetGroupResponseBody) SetTotal(v int64) *DescribeAssetGroupResponseBody {
	s.Total = &v
	return s
}

type DescribeAssetGroupResponseBodyAssetGroupList struct {
	// The ID of the asset.
	//
	// example:
	//
	// waf_v2_public_cn-lbj382l****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region to which the asset belongs.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The type of the asset.
	//
	// example:
	//
	// waf
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAssetGroupResponseBodyAssetGroupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetGroupResponseBodyAssetGroupList) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupResponseBodyAssetGroupList) SetName(v string) *DescribeAssetGroupResponseBodyAssetGroupList {
	s.Name = &v
	return s
}

func (s *DescribeAssetGroupResponseBodyAssetGroupList) SetRegion(v string) *DescribeAssetGroupResponseBodyAssetGroupList {
	s.Region = &v
	return s
}

func (s *DescribeAssetGroupResponseBodyAssetGroupList) SetType(v string) *DescribeAssetGroupResponseBodyAssetGroupList {
	s.Type = &v
	return s
}

type DescribeAssetGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAssetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAssetGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupResponse) SetHeaders(v map[string]*string) *DescribeAssetGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetGroupResponse) SetStatusCode(v int32) *DescribeAssetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetGroupResponse) SetBody(v *DescribeAssetGroupResponseBody) *DescribeAssetGroupResponse {
	s.Body = v
	return s
}

type DescribeAssetGroupToInstanceRequest struct {
	// The ID of the instance to query.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances of paid editions.
	//
	// example:
	//
	// ddosbgp-cn-7212zaa5v***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The UID of the member to which the asset belongs.
	//
	// example:
	//
	// 170858869679****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the asset. If the asset is a Web Application Firewall (WAF) instance, specify the ID of the WAF instance.
	//
	// example:
	//
	// waf_v2_public_cn-lbj382l****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the asset.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- **waf**: WAF instance
	//
	// 	- **ga**: Global Accelerator (GA) instance
	//
	// example:
	//
	// waf
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAssetGroupToInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetGroupToInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupToInstanceRequest) SetInstanceId(v string) *DescribeAssetGroupToInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) SetMemberUid(v string) *DescribeAssetGroupToInstanceRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) SetName(v string) *DescribeAssetGroupToInstanceRequest {
	s.Name = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) SetRegion(v string) *DescribeAssetGroupToInstanceRequest {
	s.Region = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) SetRegionId(v string) *DescribeAssetGroupToInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) SetType(v string) *DescribeAssetGroupToInstanceRequest {
	s.Type = &v
	return s
}

type DescribeAssetGroupToInstanceResponseBody struct {
	// The response parameters.
	DataList []*DescribeAssetGroupToInstanceResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C73C59B9-9F5C-57FF-A394-13EC8FC3B2FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAssetGroupToInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetGroupToInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupToInstanceResponseBody) SetDataList(v []*DescribeAssetGroupToInstanceResponseBodyDataList) *DescribeAssetGroupToInstanceResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBody) SetRequestId(v string) *DescribeAssetGroupToInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBody) SetTotal(v int64) *DescribeAssetGroupToInstanceResponseBody {
	s.Total = &v
	return s
}

type DescribeAssetGroupToInstanceResponseBodyDataList struct {
	// The ID of the Anti-DDoS Origin instance of a paid edition.
	//
	// example:
	//
	// ddosbgp-cn-7212zaa5v***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The UID of the member to which the asset belongs.
	//
	// example:
	//
	// 170858869679****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the asset.
	//
	// example:
	//
	// waf_v2_public_cn-lbj382l****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the asset.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The type of the asset.
	//
	// example:
	//
	// waf
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAssetGroupToInstanceResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetGroupToInstanceResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) SetInstanceId(v string) *DescribeAssetGroupToInstanceResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) SetMemberUid(v string) *DescribeAssetGroupToInstanceResponseBodyDataList {
	s.MemberUid = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) SetName(v string) *DescribeAssetGroupToInstanceResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) SetRegion(v string) *DescribeAssetGroupToInstanceResponseBodyDataList {
	s.Region = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) SetType(v string) *DescribeAssetGroupToInstanceResponseBodyDataList {
	s.Type = &v
	return s
}

type DescribeAssetGroupToInstanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAssetGroupToInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAssetGroupToInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetGroupToInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupToInstanceResponse) SetHeaders(v map[string]*string) *DescribeAssetGroupToInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetGroupToInstanceResponse) SetStatusCode(v int32) *DescribeAssetGroupToInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponse) SetBody(v *DescribeAssetGroupToInstanceResponseBody) *DescribeAssetGroupToInstanceResponse {
	s.Body = v
	return s
}

type DescribeDdosEventRequest struct {
	// The end time of the DDoS attack events to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638288000
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Anti-DDoS Origin instance to query.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The attacked IP address to query.
	//
	// example:
	//
	// 47.89.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The start time of the DDoS attack events to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1633017600
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDdosEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventRequest) SetEndTime(v int32) *DescribeDdosEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventRequest) SetInstanceId(v string) *DescribeDdosEventRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDdosEventRequest) SetIp(v string) *DescribeDdosEventRequest {
	s.Ip = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPageNo(v int32) *DescribeDdosEventRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPageSize(v int32) *DescribeDdosEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDdosEventRequest) SetRegionId(v string) *DescribeDdosEventRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDdosEventRequest) SetResourceGroupId(v string) *DescribeDdosEventRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDdosEventRequest) SetStartTime(v int32) *DescribeDdosEventRequest {
	s.StartTime = &v
	return s
}

type DescribeDdosEventResponseBody struct {
	// The details about the DDoS attack event.
	Events []*DescribeDdosEventResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F3B6C3F9-6B21-519D-B976-A1E14166F909
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of DDoS attack events that are returned.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDdosEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponseBody) SetEvents(v []*DescribeDdosEventResponseBodyEvents) *DescribeDdosEventResponseBody {
	s.Events = v
	return s
}

func (s *DescribeDdosEventResponseBody) SetRequestId(v string) *DescribeDdosEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosEventResponseBody) SetTotal(v int64) *DescribeDdosEventResponseBody {
	s.Total = &v
	return s
}

type DescribeDdosEventResponseBodyEvents struct {
	// The time when the DDoS attack stopped. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1637554335
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The attacked IP address.
	//
	// example:
	//
	// 47.89.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The volume of the request traffic at the start of the DDoS attack. Unit: Mbit/s.
	//
	// example:
	//
	// 0
	Mbps *int32 `json:"Mbps,omitempty" xml:"Mbps,omitempty"`
	// The number of packets at the start of the DDoS attack. Unit: packets per second (PPS).
	//
	// example:
	//
	// 456
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The time when the DDoS attack started. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1637554034
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the DDoS attack event. Valid values:
	//
	// 	- **hole_begin**: indicates that blackhole filtering is triggered for the attacked IP address.
	//
	// 	- **hole_end**: indicates that blackhole filtering is deactivated for the attacked IP address.
	//
	// 	- **defense_begin**: indicates that attack traffic is being scrubbed.
	//
	// 	- **defense_end**: indicates that attack traffic is scrubbed.
	//
	// example:
	//
	// defense_end
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDdosEventResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponseBodyEvents) SetEndTime(v int32) *DescribeDdosEventResponseBodyEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetIp(v string) *DescribeDdosEventResponseBodyEvents {
	s.Ip = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetMbps(v int32) *DescribeDdosEventResponseBodyEvents {
	s.Mbps = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetPps(v int32) *DescribeDdosEventResponseBodyEvents {
	s.Pps = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetStartTime(v int32) *DescribeDdosEventResponseBodyEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetStatus(v string) *DescribeDdosEventResponseBodyEvents {
	s.Status = &v
	return s
}

type DescribeDdosEventResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDdosEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDdosEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponse) SetHeaders(v map[string]*string) *DescribeDdosEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosEventResponse) SetStatusCode(v int32) *DescribeDdosEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosEventResponse) SetBody(v *DescribeDdosEventResponseBody) *DescribeDdosEventResponse {
	s.Body = v
	return s
}

type DescribeDdosOriginInstanceBillRequest struct {
	// The end of the time range to query. The value is a timestamp. Unit: milliseconds. The time span between StartTime and EndTime cannot exceed 30 days.
	//
	// example:
	//
	// 1711382399410
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to display the bill details. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsShowList *bool `json:"IsShowList,omitempty" xml:"IsShowList,omitempty"`
	// The beginning of the time range to query. The value is a timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1711209600410
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The bill type. Valid values:
	//
	// 	- **flow_cn**: the bill for the clean bandwidth of elastic IP addresses (EIPs) with Anti-DDoS (Enhanced) enabled in the Chinese mainland.
	//
	// 	- **flow_ov**: the bill for the clean bandwidth of EIPs with Anti-DDoS (Enhanced) enabled outside the Chinese mainland.
	//
	// 	- **standard_assets_flow_cn**: the bill for the clean bandwidth of regular Alibaba Cloud services in the Chinese mainland.
	//
	// 	- **standard_assets_flow_ov**: the bill for the clean bandwidth of regular Alibaba Cloud services outside the Chinese mainland.
	//
	// 	- **function**: the bill for the basic fee.
	//
	// 	- **ip_count**: the bill for protected IP addresses.
	//
	// 	- **monthly_summary**: the monthly summary bill.
	//
	// example:
	//
	// function
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDdosOriginInstanceBillRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosOriginInstanceBillRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosOriginInstanceBillRequest) SetEndTime(v int64) *DescribeDdosOriginInstanceBillRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillRequest) SetIsShowList(v bool) *DescribeDdosOriginInstanceBillRequest {
	s.IsShowList = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillRequest) SetStartTime(v int64) *DescribeDdosOriginInstanceBillRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillRequest) SetType(v string) *DescribeDdosOriginInstanceBillRequest {
	s.Type = &v
	return s
}

type DescribeDdosOriginInstanceBillResponseBody struct {
	// The asset status.
	//
	// 	- **0**: No asset is added to the instance for protection.
	//
	// 	- **1**: Assets are added to the instance for protection.
	//
	// example:
	//
	// 0
	AssetStatus *int32 `json:"AssetStatus,omitempty" xml:"AssetStatus,omitempty"`
	// The payment status. Valid values:
	//
	// 	- **0**: The payment is not overdue.
	//
	// 	- **1**: The payment is overdue.
	//
	// example:
	//
	// 0
	DebtStatus *int64 `json:"DebtStatus,omitempty" xml:"DebtStatus,omitempty"`
	// The details about the traffic of EIPs with Anti-DDoS (Enhanced) enabled.
	FlowList []*DescribeDdosOriginInstanceBillResponseBodyFlowList `json:"FlowList,omitempty" xml:"FlowList,omitempty" type:"Repeated"`
	// The traffic distribution of EIPs with Anti-DDoS (Enhanced) enabled by region.
	//
	// example:
	//
	// {\\"cn-hongkong\\": 166491566}
	FlowRegion map[string]interface{} `json:"FlowRegion,omitempty" xml:"FlowRegion,omitempty"`
	// The ID of the Anti-DDoS Origin (Pay-as-you-go) instance to query.
	//
	// example:
	//
	// ddosorigin_cn-u7c3lcr9r02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of protected IP addresses.
	//
	// example:
	//
	// 15
	IpCount *int64 `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	// The protected IP addresses and enabled features.
	IpCountOrFunctionList []*DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList `json:"IpCountOrFunctionList,omitempty" xml:"IpCountOrFunctionList,omitempty" type:"Repeated"`
	// The IP address distribution. The JSON struct contains the following fields:
	//
	// 	- **eipCnIpCount**: the number of EIPs with Anti-DDoS (Enhanced) enabled in the Chinese mainland.
	//
	// 	- **eipOvIpCount**: the number of EIPs with Anti-DDoS (Enhanced) enabled outside the Chinese mainland.
	//
	// 	- **standardAssetsCnIpCount**: the number of IP addresses of regular Alibaba Cloud services in the Chinese mainland.
	//
	// 	- **standardAssetsOvIpCount**: the number of IP addresses of regular Alibaba Cloud services outside the Chinese mainland.
	//
	// example:
	//
	// {\\"eipCnIpCount\\":6,\\"eipOvIpCount\\":17,\\"standardAssetsCnIpCount\\":2,\\"standardAssetsOvIpCount\\":0}
	IpInfo *string `json:"IpInfo,omitempty" xml:"IpInfo,omitempty"`
	// The information about the monthly summary bills.
	MonthlySummaryList []*DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList `json:"MonthlySummaryList,omitempty" xml:"MonthlySummaryList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 72155560-F343-55C8-82FE-ED4D7E4AA97E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the traffic of regular Alibaba Cloud services.
	StandardAssetsFlowList []*DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList `json:"StandardAssetsFlowList,omitempty" xml:"StandardAssetsFlowList,omitempty" type:"Repeated"`
	// The traffic distribution of regular Alibaba Cloud services by region.
	//
	// example:
	//
	// {\\"cn-hongkong\\": 166491566}
	StandardAssetsFlowRegion map[string]interface{} `json:"StandardAssetsFlowRegion,omitempty" xml:"StandardAssetsFlowRegion,omitempty"`
	// The total traffic of regular Alibaba Cloud services in the Chinese mainland in the current month.
	//
	// example:
	//
	// 0
	StandardAssetsTotalFlowCn *int64 `json:"StandardAssetsTotalFlowCn,omitempty" xml:"StandardAssetsTotalFlowCn,omitempty"`
	// The total traffic of regular Alibaba Cloud services outside the Chinese mainland in the current month.
	//
	// example:
	//
	// 0
	StandardAssetsTotalFlowOv *int64 `json:"StandardAssetsTotalFlowOv,omitempty" xml:"StandardAssetsTotalFlowOv,omitempty"`
	// The instance status. Valid values:
	//
	// 	- **1**: normal
	//
	// 	- **2**: expired
	//
	// 	- **3**: released
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total traffic of EIPs with Anti-DDoS (Enhanced) enabled in the Chinese mainland in the current month. Unit: bytes.
	//
	// example:
	//
	// 6302081067
	TotalFlowCn *int64 `json:"TotalFlowCn,omitempty" xml:"TotalFlowCn,omitempty"`
	// The total traffic of EIPs with Anti-DDoS (Enhanced) enabled outside the Chinese mainland in the current month. Unit: bytes.
	//
	// example:
	//
	// 6204918019
	TotalFlowOv *int64 `json:"TotalFlowOv,omitempty" xml:"TotalFlowOv,omitempty"`
}

func (s DescribeDdosOriginInstanceBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosOriginInstanceBillResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetAssetStatus(v int32) *DescribeDdosOriginInstanceBillResponseBody {
	s.AssetStatus = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetDebtStatus(v int64) *DescribeDdosOriginInstanceBillResponseBody {
	s.DebtStatus = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetFlowList(v []*DescribeDdosOriginInstanceBillResponseBodyFlowList) *DescribeDdosOriginInstanceBillResponseBody {
	s.FlowList = v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetFlowRegion(v map[string]interface{}) *DescribeDdosOriginInstanceBillResponseBody {
	s.FlowRegion = v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetInstanceId(v string) *DescribeDdosOriginInstanceBillResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetIpCount(v int64) *DescribeDdosOriginInstanceBillResponseBody {
	s.IpCount = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetIpCountOrFunctionList(v []*DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) *DescribeDdosOriginInstanceBillResponseBody {
	s.IpCountOrFunctionList = v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetIpInfo(v string) *DescribeDdosOriginInstanceBillResponseBody {
	s.IpInfo = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetMonthlySummaryList(v []*DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) *DescribeDdosOriginInstanceBillResponseBody {
	s.MonthlySummaryList = v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetRequestId(v string) *DescribeDdosOriginInstanceBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetStandardAssetsFlowList(v []*DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) *DescribeDdosOriginInstanceBillResponseBody {
	s.StandardAssetsFlowList = v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetStandardAssetsFlowRegion(v map[string]interface{}) *DescribeDdosOriginInstanceBillResponseBody {
	s.StandardAssetsFlowRegion = v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetStandardAssetsTotalFlowCn(v int64) *DescribeDdosOriginInstanceBillResponseBody {
	s.StandardAssetsTotalFlowCn = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetStandardAssetsTotalFlowOv(v int64) *DescribeDdosOriginInstanceBillResponseBody {
	s.StandardAssetsTotalFlowOv = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetStatus(v int64) *DescribeDdosOriginInstanceBillResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetTotalFlowCn(v int64) *DescribeDdosOriginInstanceBillResponseBody {
	s.TotalFlowCn = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBody) SetTotalFlowOv(v int64) *DescribeDdosOriginInstanceBillResponseBody {
	s.TotalFlowOv = &v
	return s
}

type DescribeDdosOriginInstanceBillResponseBodyFlowList struct {
	// The traffic distribution by region. The JSON struct contains the following fields:
	//
	// 	- **bytes**: the traffic volume of EIPs with Anti-DDoS (Enhanced) enabled in a region. Unit: bytes.
	//
	// 	- **memberUid**: the owner account.
	//
	// 	- **instanceId**: the ID of the pay-as-you-go instance that protects the EIPs with Anti-DDoS (Enhanced) enabled.
	//
	// 	- **ip**: the EIPs with Anti-DDoS (Enhanced) enabled.
	//
	// 	- **region**: the region.
	//
	// >  If the memberUid field in the JSON struct is empty, the information about the current account is returned. The value of the bytes parameter in the outermost level of the JSON struct indicates the total traffic, and the values of the bytes parameters in inner levels indicate the traffic of the account.
	//
	// example:
	//
	// [{\\"bytes\\":79282719,\\"memberUid\\":\\"\\",\\"regionFlows\\":{\\"cn-hangzhou\\":[{\\"bytes\\":79282719,\\"instanceId\\":\\"ddosorigin_cn-u7c3lcr9r02\\",\\"ip\\":\\"47.118.168.57\\",\\"region\\":\\"cn-hangzhou\\"}]}}]
	MemberFlow *string `json:"MemberFlow,omitempty" xml:"MemberFlow,omitempty"`
	// The traffic distribution by region. The JSON struct contains the following fields:
	//
	// 	- **bytes**: the traffic volume of EIPs with Anti-DDoS (Enhanced) enabled in a region. Unit: bytes.
	//
	// 	- **instanceId**: the ID of the pay-as-you-go instance that protects the EIPs with Anti-DDoS (Enhanced) enabled.
	//
	// 	- **ip**: the EIPs with Anti-DDoS (Enhanced) enabled.
	//
	// 	- **region**: the region.
	//
	// example:
	//
	// {\\"cn-hangzhou\\":[{\\"bytes\\":0,\\"instanceId\\":\\"ddosorigin_cn-u7c3lcr9r02\\",\\"ip\\":\\"47.118.168.124\\",\\"region\\":\\"cn-hangzhou\\"}]}
	RegionFlow *string `json:"RegionFlow,omitempty" xml:"RegionFlow,omitempty"`
	// The timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1620951900
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The traffic of EIPs with Anti-DDoS (Enhanced) enabled. Unit: bytes.
	//
	// example:
	//
	// 6302081067
	TotalFlow *int64 `json:"TotalFlow,omitempty" xml:"TotalFlow,omitempty"`
}

func (s DescribeDdosOriginInstanceBillResponseBodyFlowList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosOriginInstanceBillResponseBodyFlowList) GoString() string {
	return s.String()
}

func (s *DescribeDdosOriginInstanceBillResponseBodyFlowList) SetMemberFlow(v string) *DescribeDdosOriginInstanceBillResponseBodyFlowList {
	s.MemberFlow = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyFlowList) SetRegionFlow(v string) *DescribeDdosOriginInstanceBillResponseBodyFlowList {
	s.RegionFlow = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyFlowList) SetTime(v int64) *DescribeDdosOriginInstanceBillResponseBodyFlowList {
	s.Time = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyFlowList) SetTotalFlow(v int64) *DescribeDdosOriginInstanceBillResponseBodyFlowList {
	s.TotalFlow = &v
	return s
}

type DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList struct {
	// The application scope of the instance. Valid values:
	//
	// 	- **only_mainland_china**: regions in the Chinese mainland.
	//
	// 	- **global**: all regions.
	//
	// 	- **international_and_hmt**: regions outside the Chinese mainland.
	//
	// example:
	//
	// global
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The number of IP addresses protected by the pay-as-you-go instance in the Chinese mainland.
	//
	// example:
	//
	// 5
	IpCntCn *int64 `json:"IpCntCn,omitempty" xml:"IpCntCn,omitempty"`
	// The number of IP addresses protected by the pay-as-you-go instance outside the Chinese mainland.
	//
	// example:
	//
	// 5
	IpCntOv *int64 `json:"IpCntOv,omitempty" xml:"IpCntOv,omitempty"`
	// The bill distribution by account. The JSON struct contains the following fields:
	//
	// 	- **eipCnIpCount**: the number of EIPs with Anti-DDoS (Enhanced) enabled in the Chinese mainland.
	//
	// 	- **eipOvIpCount**: the number of EIPs with Anti-DDoS (Enhanced) enabled outside the Chinese mainland.
	//
	// 	- **memberUid**: the owner account.
	//
	// 	- **standardAssetsCnIpCount**: the number of IP addresses of regular Alibaba Cloud services in the Chinese mainland.
	//
	// 	- **standardAssetsOvIpCount**: the number of IP addresses of regular Alibaba Cloud services outside the Chinese mainland.
	//
	// >  If the memberUid field in the JSON struct is empty, the information about the current account is returned.
	//
	// example:
	//
	// [{\\"eipCnIpCount\\":3,\\"eipOvIpCount\\":18,\\"memberUid\\":\\"\\",\\"standardAssetsCnIpCount\\":2,\\"standardAssetsOvIpCount\\":0},{\\"eipCnIpCount\\":3,\\"eipOvIpCount\\":0,\\"memberUid\\":\\"1776997906319249\\",\\"standardAssetsCnIpCount\\":0,\\"standardAssetsOvIpCount\\":0}]
	MemberIpCnt *string `json:"MemberIpCnt,omitempty" xml:"MemberIpCnt,omitempty"`
	// The billing time. Unit: milliseconds.
	//
	// example:
	//
	// 1680278400000
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) GoString() string {
	return s.String()
}

func (s *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) SetCoverage(v string) *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList {
	s.Coverage = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) SetIpCntCn(v int64) *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList {
	s.IpCntCn = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) SetIpCntOv(v int64) *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList {
	s.IpCntOv = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) SetMemberIpCnt(v string) *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList {
	s.MemberIpCnt = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList) SetTime(v int64) *DescribeDdosOriginInstanceBillResponseBodyIpCountOrFunctionList {
	s.Time = &v
	return s
}

type DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList struct {
	// The number of days that the instance is activated.
	//
	// example:
	//
	// 30
	EnableDays *int32 `json:"EnableDays,omitempty" xml:"EnableDays,omitempty"`
	// The total traffic of EIPs with Anti-DDoS (Enhanced) enabled in the Chinese mainland. Unit: bytes.
	//
	// example:
	//
	// 123456
	FlowCn *int64 `json:"FlowCn,omitempty" xml:"FlowCn,omitempty"`
	// The total traffic of EIPs with Anti-DDoS (Enhanced) enabled outside the Chinese mainland. Unit: bytes.
	//
	// example:
	//
	// 123456
	FlowIntl *int64 `json:"FlowIntl,omitempty" xml:"FlowIntl,omitempty"`
	// The total number of protected IP addresses in the Chinese mainland.
	//
	// >  The total number of protected IP addresses is the sum of the daily numbers of protected IP addresses in a month.
	//
	// example:
	//
	// 28
	IpCountCn *int32 `json:"IpCountCn,omitempty" xml:"IpCountCn,omitempty"`
	// The total number of protected IP addresses outside the Chinese mainland.
	//
	// >  The total number of protected IP addresses is the sum of the daily numbers of protected IP addresses in a month.
	//
	// example:
	//
	// 30
	IpCountIntl *int32 `json:"IpCountIntl,omitempty" xml:"IpCountIntl,omitempty"`
	// The ID of the member.
	//
	// example:
	//
	// 112873971277****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The total traffic of regular Alibaba Cloud services in the Chinese mainland. Unit: bytes.
	//
	// example:
	//
	// 123456
	StandardAssetsFlowCn *int64 `json:"StandardAssetsFlowCn,omitempty" xml:"StandardAssetsFlowCn,omitempty"`
	// The total traffic of regular Alibaba Cloud services outside the Chinese mainland. Unit: bytes.
	//
	// example:
	//
	// 123456
	StandardAssetsFlowIntl *int64 `json:"StandardAssetsFlowIntl,omitempty" xml:"StandardAssetsFlowIntl,omitempty"`
	// The ID of the administrator account.
	//
	// example:
	//
	// 102518028277****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) GoString() string {
	return s.String()
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetEnableDays(v int32) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.EnableDays = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetFlowCn(v int64) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.FlowCn = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetFlowIntl(v int64) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.FlowIntl = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetIpCountCn(v int32) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.IpCountCn = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetIpCountIntl(v int32) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.IpCountIntl = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetMemberUid(v string) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.MemberUid = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetStandardAssetsFlowCn(v int64) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.StandardAssetsFlowCn = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetStandardAssetsFlowIntl(v int64) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.StandardAssetsFlowIntl = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList) SetUid(v string) *DescribeDdosOriginInstanceBillResponseBodyMonthlySummaryList {
	s.Uid = &v
	return s
}

type DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList struct {
	// The traffic distribution by region. The JSON struct contains the following fields:
	//
	// 	- **bytes**: the traffic volume of regular Alibaba Cloud services in a region. Unit: bytes.
	//
	// 	- **memberUid**: the owner account.
	//
	// 	- **instanceId**: the ID of the pay-as-you-go instance that protects the regular Alibaba Cloud services.
	//
	// 	- **ip**: the IP address of the regular Alibaba Cloud service protected by the Anti-DDoS Origin instance.
	//
	// 	- **region**: the region.
	//
	// >  If the memberUid field in the JSON struct is empty, the information about the current account is returned. The value of the bytes parameter in the outermost level of the JSON struct indicates the total traffic, and the values of the bytes parameters in inner levels indicate the traffic of the account.
	//
	// example:
	//
	// [{\\"bytes\\":79282719,\\"memberUid\\":\\"\\",\\"regionFlows\\":{\\"cn-hangzhou\\":[{\\"bytes\\":79282719,\\"instanceId\\":\\"ddosorigin_cn-u7c3lcr9r02\\",\\"ip\\":\\"47.118.168.57\\",\\"region\\":\\"cn-hangzhou\\"}]}}]
	MemberFlow *string `json:"MemberFlow,omitempty" xml:"MemberFlow,omitempty"`
	// The traffic distribution by region. The JSON struct contains the following fields:
	//
	// 	- **bytes**: the traffic volume of regular Alibaba Cloud services in a region. Unit: bytes.
	//
	// 	- **instanceId**: the ID of the pay-as-you-go instance that protects the regular Alibaba Cloud services.
	//
	// 	- **ip**: the IP address protected by the Anti-DDoS Origin instance.
	//
	// 	- **region**: the region.
	//
	// example:
	//
	// {\\"cn-hangzhou\\":[{\\"bytes\\":0,\\"instanceId\\":\\"ddosorigin_cn-u7c3lcr9r02\\",\\"ip\\":\\"47.118.168.124\\",\\"region\\":\\"cn-hangzhou\\"}]}
	RegionFlow *string `json:"RegionFlow,omitempty" xml:"RegionFlow,omitempty"`
	// The timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1679846400000
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The traffic of regular Alibaba Cloud services. Unit: bytes.
	//
	// example:
	//
	// 6302081067
	TotalFlow *int64 `json:"TotalFlow,omitempty" xml:"TotalFlow,omitempty"`
}

func (s DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) GoString() string {
	return s.String()
}

func (s *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) SetMemberFlow(v string) *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList {
	s.MemberFlow = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) SetRegionFlow(v string) *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList {
	s.RegionFlow = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) SetTime(v int64) *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList {
	s.Time = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList) SetTotalFlow(v int64) *DescribeDdosOriginInstanceBillResponseBodyStandardAssetsFlowList {
	s.TotalFlow = &v
	return s
}

type DescribeDdosOriginInstanceBillResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDdosOriginInstanceBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDdosOriginInstanceBillResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosOriginInstanceBillResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosOriginInstanceBillResponse) SetHeaders(v map[string]*string) *DescribeDdosOriginInstanceBillResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponse) SetStatusCode(v int32) *DescribeDdosOriginInstanceBillResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosOriginInstanceBillResponse) SetBody(v *DescribeDdosOriginInstanceBillResponseBody) *DescribeDdosOriginInstanceBillResponse {
	s.Body = v
	return s
}

type DescribeExcpetionCountRequest struct {
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeExcpetionCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcpetionCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeExcpetionCountRequest) SetRegionId(v string) *DescribeExcpetionCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeExcpetionCountRequest) SetResourceGroupId(v string) *DescribeExcpetionCountRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeExcpetionCountResponseBody struct {
	// The number of assets that are in an abnormal state.
	//
	// example:
	//
	// 0
	ExceptionIpCount *int32 `json:"ExceptionIpCount,omitempty" xml:"ExceptionIpCount,omitempty"`
	// The number of Anti-DDoS Origin instances that are about to expire.
	//
	// example:
	//
	// 1
	ExpireTimeCount *int32 `json:"ExpireTimeCount,omitempty" xml:"ExpireTimeCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4B45279A-B1BE-5EEE-87CA-58AF4183EA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExcpetionCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcpetionCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExcpetionCountResponseBody) SetExceptionIpCount(v int32) *DescribeExcpetionCountResponseBody {
	s.ExceptionIpCount = &v
	return s
}

func (s *DescribeExcpetionCountResponseBody) SetExpireTimeCount(v int32) *DescribeExcpetionCountResponseBody {
	s.ExpireTimeCount = &v
	return s
}

func (s *DescribeExcpetionCountResponseBody) SetRequestId(v string) *DescribeExcpetionCountResponseBody {
	s.RequestId = &v
	return s
}

type DescribeExcpetionCountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExcpetionCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExcpetionCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcpetionCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeExcpetionCountResponse) SetHeaders(v map[string]*string) *DescribeExcpetionCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeExcpetionCountResponse) SetStatusCode(v int32) *DescribeExcpetionCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExcpetionCountResponse) SetBody(v *DescribeExcpetionCountResponseBody) *DescribeExcpetionCountResponse {
	s.Body = v
	return s
}

type DescribeInstanceListRequest struct {
	// The IDs of the Anti-DDoS Origin instances to query. Specify the value is in the `["<Instance ID 1>","<Instance ID 2>",……]` format.
	//
	// example:
	//
	// ["ddosbgp-cn-oew1pjrk****"]
	InstanceIdList *string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	// The mitigation plan of the Anti-DDoS Origin instance to query. Valid values:
	//
	// 	- **0**: the Professional mitigation plan
	//
	// 	- **1**: the Enterprise mitigation plan
	//
	// example:
	//
	// 0
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The mitigation plan of the Anti-DDoS Origin instance.
	InstanceTypeList []*string `json:"InstanceTypeList,omitempty" xml:"InstanceTypeList,omitempty" type:"Repeated"`
	// The IP address of the object that is protected by the Anti-DDoS Origin instance to query.
	//
	// example:
	//
	// 47.89.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The protocol type of the IP address asset that is protected by the Anti-DDoS Origin instance to query. Valid values:
	//
	// 	- **Ipv4**: IPv4
	//
	// 	- **Ipv6**: IPv6
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The field that is used to sort the Anti-DDoS Origin instances. Set the value to **expireTime**, which indicates that the instances are sorted based on the expiration time.
	//
	// You can set the **Orderdire*	- parameter to specify the sorting method.
	//
	// example:
	//
	// expireTime
	Orderby *string `json:"Orderby,omitempty" xml:"Orderby,omitempty"`
	// The sorting method. Valid values:
	//
	// 	- **desc**: the descending order. This is the default value.
	//
	// 	- **asc**: the ascending order.
	//
	// example:
	//
	// desc
	Orderdire *string `json:"Orderdire,omitempty" xml:"Orderdire,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance to query resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the Anti-DDoS Origin instance to query. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags that are added to the Anti-DDoS Origin instance.
	Tag []*DescribeInstanceListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListRequest) SetInstanceIdList(v string) *DescribeInstanceListRequest {
	s.InstanceIdList = &v
	return s
}

func (s *DescribeInstanceListRequest) SetInstanceType(v string) *DescribeInstanceListRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceListRequest) SetInstanceTypeList(v []*string) *DescribeInstanceListRequest {
	s.InstanceTypeList = v
	return s
}

func (s *DescribeInstanceListRequest) SetIp(v string) *DescribeInstanceListRequest {
	s.Ip = &v
	return s
}

func (s *DescribeInstanceListRequest) SetIpVersion(v string) *DescribeInstanceListRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeInstanceListRequest) SetOrderby(v string) *DescribeInstanceListRequest {
	s.Orderby = &v
	return s
}

func (s *DescribeInstanceListRequest) SetOrderdire(v string) *DescribeInstanceListRequest {
	s.Orderdire = &v
	return s
}

func (s *DescribeInstanceListRequest) SetPageNo(v int32) *DescribeInstanceListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeInstanceListRequest) SetPageSize(v int32) *DescribeInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceListRequest) SetRegionId(v string) *DescribeInstanceListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceListRequest) SetRemark(v string) *DescribeInstanceListRequest {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceListRequest) SetResourceGroupId(v string) *DescribeInstanceListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceListRequest) SetTag(v []*DescribeInstanceListRequestTag) *DescribeInstanceListRequest {
	s.Tag = v
	return s
}

type DescribeInstanceListRequestTag struct {
	// The key of the tag that is added to the Anti-DDoS Origin instance.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that is added to the Anti-DDoS Origin instance.
	//
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceListRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListRequestTag) SetKey(v string) *DescribeInstanceListRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstanceListRequestTag) SetValue(v string) *DescribeInstanceListRequestTag {
	s.Value = &v
	return s
}

type DescribeInstanceListResponseBody struct {
	// The details about the Anti-DDoS Origin instances.
	InstanceList []*DescribeInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 381D5D33-BB8F-395F-8EE4-AE3BB4B523C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of Anti-DDoS Origin instances.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponseBody) SetInstanceList(v []*DescribeInstanceListResponseBodyInstanceList) *DescribeInstanceListResponseBody {
	s.InstanceList = v
	return s
}

func (s *DescribeInstanceListResponseBody) SetRequestId(v string) *DescribeInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceListResponseBody) SetTotal(v int64) *DescribeInstanceListResponseBody {
	s.Total = &v
	return s
}

type DescribeInstanceListResponseBodyInstanceList struct {
	// The condition that triggers automatic association of the instance with an object.
	AutoProtectCondition *DescribeInstanceListResponseBodyInstanceListAutoProtectCondition `json:"AutoProtectCondition,omitempty" xml:"AutoProtectCondition,omitempty" type:"Struct"`
	// Indicates whether auto-renewal is enabled for the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The number of protected public IP addresses for which blackhole filtering is triggered.
	//
	// >  You can call the [DeleteBlackhole](https://help.aliyun.com/document_detail/118692.html) operation to deactivate blackhole filtering for a protected IP address.
	//
	// example:
	//
	// 0
	BlackholdingCount *string `json:"BlackholdingCount,omitempty" xml:"BlackholdingCount,omitempty"`
	// The type of the instance.
	//
	// 	- **ddos_ddosorigin_public_cn**: Anti-DDoS Origin 2.0 (Pay-as-you-go) on the China site (aliyun.com).
	//
	// 	- **ddos_ddosorigin_public_intl**: Anti-DDoS Origin 2.0 (Pay-as-you-go) on the International site (alibabacloud.com).
	//
	// example:
	//
	// ddos_ddosorigin_public_cn
	CommodityType *string `json:"CommodityType,omitempty" xml:"CommodityType,omitempty"`
	// The application scope of the instance.
	//
	// 	- **1**: The instance supports public IP addresses in all regions.
	//
	// 	- **2**: The instance supports public IP addresses in regions in the Chinese mainland.
	//
	// 	- **3**: The instance supports public IP addresses in regions outside the Chinese mainland.
	//
	// 	- **4**: The instance supports public IP addresses in a region in or outside the Chinese mainland.
	//
	// example:
	//
	// 1
	CoverageType *int32 `json:"CoverageType,omitempty" xml:"CoverageType,omitempty"`
	DebtStatus   *int64 `json:"DebtStatus,omitempty" xml:"DebtStatus,omitempty"`
	// The time when the instance expires. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1640275200000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the instance was purchased. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1592886047000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ddosbgp-cn-oew1pjrk****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The mitigation plan of the instance. Valid values:
	//
	// 	- **0**: the Professional mitigation plan
	//
	// 	- **1**: the Enterprise mitigation plan
	//
	// example:
	//
	// 1
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The protocol type of the IP address asset that is protected by the instance. Valid values:
	//
	// 	- **Ipv4**
	//
	// 	- **Ipv6**
	//
	// example:
	//
	// IPv4
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The type of the cloud service that is associated with the Anti-DDoS Origin instance By default, this parameter is not returned. If the Anti-DDoS Origin instance is created by using a different cloud service, the code of the cloud service is returned.
	//
	// Valid values:
	//
	// 	- **gamebox**: The Anti-DDoS Origin instance is created by using Game Security Box.
	//
	// 	- **eip**: The Anti-DDoS Origin instance is created by using an elastic IP address (EIP) for which Anti-DDoS (Enhanced Edition) is enabled.
	//
	// example:
	//
	// gamebox
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- **1**: normal
	//
	// 	- **2**: expired
	//
	// 	- **3**: released
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceListResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetAutoProtectCondition(v *DescribeInstanceListResponseBodyInstanceListAutoProtectCondition) *DescribeInstanceListResponseBodyInstanceList {
	s.AutoProtectCondition = v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetAutoRenewal(v bool) *DescribeInstanceListResponseBodyInstanceList {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetBlackholdingCount(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.BlackholdingCount = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetCommodityType(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.CommodityType = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetCoverageType(v int32) *DescribeInstanceListResponseBodyInstanceList {
	s.CoverageType = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetDebtStatus(v int64) *DescribeInstanceListResponseBodyInstanceList {
	s.DebtStatus = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetExpireTime(v int64) *DescribeInstanceListResponseBodyInstanceList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetGmtCreate(v int64) *DescribeInstanceListResponseBodyInstanceList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetInstanceId(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetInstanceType(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetIpType(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.IpType = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetProduct(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Product = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetRemark(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetStatus(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Status = &v
	return s
}

type DescribeInstanceListResponseBodyInstanceListAutoProtectCondition struct {
	// The events that trigger automatic association.
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
}

func (s DescribeInstanceListResponseBodyInstanceListAutoProtectCondition) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponseBodyInstanceListAutoProtectCondition) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponseBodyInstanceListAutoProtectCondition) SetEvents(v []*string) *DescribeInstanceListResponseBodyInstanceListAutoProtectCondition {
	s.Events = v
	return s
}

type DescribeInstanceListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponse) SetHeaders(v map[string]*string) *DescribeInstanceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceListResponse) SetStatusCode(v int32) *DescribeInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceListResponse) SetBody(v *DescribeInstanceListResponseBody) *DescribeInstanceListResponse {
	s.Body = v
	return s
}

type DescribeInstanceSpecsRequest struct {
	// The ID of the Anti-DDoS Origin instance. This parameter is a string that consists of JSON arrays. Each element in a JSON array indicates an instance ID. If you want to query more than one instance, separate instance IDs with commas (,).
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["ddosbgp-cn-n6w1r7nz****"]
	InstanceIdList *string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	// The region ID of the Anti-DDoS Origin instance. Default value: **cn-hangzhou**, which indicates the China (Hangzhou) region.
	//
	// >  If your instance does not reside in the China (Hangzhou) region, you must set this parameter to the region ID of your instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the regions of assets that can be protected by Anti-DDoS Origin in a specific region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. This parameter is empty by default, which indicates that the Anti-DDoS Origin instance belongs to the default resource group.
	//
	// For information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeInstanceSpecsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsRequest) SetInstanceIdList(v string) *DescribeInstanceSpecsRequest {
	s.InstanceIdList = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) SetRegionId(v string) *DescribeInstanceSpecsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) SetResourceGroupId(v string) *DescribeInstanceSpecsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeInstanceSpecsResponseBody struct {
	// The specifications of the Anti-DDoS Origin instance, including whether best-effort protection is enabled, the number of available best-effort protection sessions, and the number of used best-effort protection sessions.
	InstanceSpecs []*DescribeInstanceSpecsResponseBodyInstanceSpecs `json:"InstanceSpecs,omitempty" xml:"InstanceSpecs,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 5840AB9F-1419-4620-807D-5EA476090247
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBody) SetInstanceSpecs(v []*DescribeInstanceSpecsResponseBodyInstanceSpecs) *DescribeInstanceSpecsResponseBody {
	s.InstanceSpecs = v
	return s
}

func (s *DescribeInstanceSpecsResponseBody) SetRequestId(v string) *DescribeInstanceSpecsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceSpecsResponseBodyInstanceSpecs struct {
	// The available best-effort protection sessions.
	//
	// example:
	//
	// 2
	AvailableDefenseTimes *int32 `json:"AvailableDefenseTimes,omitempty" xml:"AvailableDefenseTimes,omitempty"`
	// The number of times that blackhole filtering can be deactivated.
	//
	// example:
	//
	// 100
	AvailableDeleteBlackholeCount *string `json:"AvailableDeleteBlackholeCount,omitempty" xml:"AvailableDeleteBlackholeCount,omitempty"`
	// The percentage of the used best-effort protection sessions. Unit: %.
	//
	// example:
	//
	// 30
	DefenseTimesPercent *int32 `json:"DefenseTimesPercent,omitempty" xml:"DefenseTimesPercent,omitempty"`
	// Indicates whether the instance is downgraded. Valid value:
	//
	// 	- **8**: The instance is downgraded due to excessive bandwidth usage.
	//
	// example:
	//
	// 8
	DowngradeStatus *int32 `json:"DowngradeStatus,omitempty" xml:"DowngradeStatus,omitempty"`
	// The ID of the Anti-DDoS Origin instance.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether best-effort protection is enabled. Valid values:
	//
	// 	- **0**: Best-effort protection is disabled.
	//
	// 	- **1**: Best-effort protection is enabled.
	//
	// example:
	//
	// 1
	IsFullDefenseMode *int32 `json:"IsFullDefenseMode,omitempty" xml:"IsFullDefenseMode,omitempty"`
	// The configurations of the Anti-DDoS Origin instance, including the number of protected IP addresses and protection bandwidth.
	PackConfig *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig `json:"PackConfig,omitempty" xml:"PackConfig,omitempty" type:"Struct"`
	// The region ID of the Anti-DDoS Origin instance.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the name of the region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The total best-effort protection sessions.
	//
	// example:
	//
	// 2
	TotalDefenseTimes *int32 `json:"TotalDefenseTimes,omitempty" xml:"TotalDefenseTimes,omitempty"`
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetAvailableDefenseTimes(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.AvailableDefenseTimes = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetAvailableDeleteBlackholeCount(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.AvailableDeleteBlackholeCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetDefenseTimesPercent(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.DefenseTimesPercent = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetDowngradeStatus(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.DowngradeStatus = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetInstanceId(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetIsFullDefenseMode(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.IsFullDefenseMode = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetPackConfig(v *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.PackConfig = v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetRegion(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.Region = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetTotalDefenseTimes(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.TotalDefenseTimes = &v
	return s
}

type DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig struct {
	// The bandwidth. Unit: Gbit/s.
	//
	// example:
	//
	// 2
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The number of IP addresses that are protected by the Anti-DDoS Origin Enterprise instance.
	//
	// example:
	//
	// 0
	BindIpCount *int32 `json:"BindIpCount,omitempty" xml:"BindIpCount,omitempty"`
	// The burstable clean bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	ElasticBwMbps *int32 `json:"ElasticBwMbps,omitempty" xml:"ElasticBwMbps,omitempty"`
	// The metering method of burstable clean bandwidth. Valid values:
	//
	// 	- **month**: the monthly 95th percentile metering method.
	//
	// 	- **day**: the daily 95th percentile metering method.
	//
	// example:
	//
	// day
	ElasticBwMode *string `json:"ElasticBwMode,omitempty" xml:"ElasticBwMode,omitempty"`
	// The burstable protection bandwidth of each protected IP address. Unit: Gbit/s.
	//
	// example:
	//
	// 300
	IpAdvanceThre *int32 `json:"IpAdvanceThre,omitempty" xml:"IpAdvanceThre,omitempty"`
	// The basic protection bandwidth of each protected IP address. Unit: Gbit/s.
	//
	// example:
	//
	// 20
	IpBasicThre *int32 `json:"IpBasicThre,omitempty" xml:"IpBasicThre,omitempty"`
	// The number of IP addresses that can be protected by the Anti-DDoS Origin Enterprise instance.
	//
	// example:
	//
	// 100
	IpSpec *int32 `json:"IpSpec,omitempty" xml:"IpSpec,omitempty"`
	// The clean bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 200
	NormalBandwidth *int32 `json:"NormalBandwidth,omitempty" xml:"NormalBandwidth,omitempty"`
	// The burstable protection bandwidth of the Anti-DDoS Origin instance. Unit: Gbit/s.
	//
	// example:
	//
	// 300
	PackAdvThre *int32 `json:"PackAdvThre,omitempty" xml:"PackAdvThre,omitempty"`
	// The basic protection bandwidth of the Anti-DDoS Origin instance. Unit: Gbit/s.
	//
	// example:
	//
	// 20
	PackBasicThre *int32 `json:"PackBasicThre,omitempty" xml:"PackBasicThre,omitempty"`
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetBandwidth(v int64) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.Bandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetBindIpCount(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.BindIpCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetElasticBwMbps(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.ElasticBwMbps = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetElasticBwMode(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.ElasticBwMode = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetIpAdvanceThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.IpAdvanceThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetIpBasicThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.IpBasicThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetIpSpec(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.IpSpec = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetNormalBandwidth(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.NormalBandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetPackAdvThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.PackAdvThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetPackBasicThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.PackBasicThre = &v
	return s
}

type DescribeInstanceSpecsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceSpecsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponse) SetHeaders(v map[string]*string) *DescribeInstanceSpecsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSpecsResponse) SetStatusCode(v int32) *DescribeInstanceSpecsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSpecsResponse) SetBody(v *DescribeInstanceSpecsResponseBody) *DescribeInstanceSpecsResponse {
	s.Body = v
	return s
}

type DescribeOnDemandDdosEventRequest struct {
	// The end time of the DDoS attack events to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1557909844
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the anti-DDoS diversion instance to query.
	//
	// >  You can call the [DescribeOnDemandInstance](https://help.aliyun.com/document_detail/152120.html) operation to query the IDs of all anti-DDoS diversion instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the anti-DDoS diversion instance to query.
	//
	// example:
	//
	// 192.XX.XX.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: **50**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the anti-DDoS diversion instance to query.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The start time of the DDoS attack events to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1557305044
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOnDemandDdosEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandDdosEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandDdosEventRequest) SetEndTime(v int32) *DescribeOnDemandDdosEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetInstanceId(v string) *DescribeOnDemandDdosEventRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetIp(v string) *DescribeOnDemandDdosEventRequest {
	s.Ip = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetPageNo(v int32) *DescribeOnDemandDdosEventRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetPageSize(v int32) *DescribeOnDemandDdosEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetRegionId(v string) *DescribeOnDemandDdosEventRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetResourceGroupId(v string) *DescribeOnDemandDdosEventRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetStartTime(v int32) *DescribeOnDemandDdosEventRequest {
	s.StartTime = &v
	return s
}

type DescribeOnDemandDdosEventResponseBody struct {
	// The details about the DDoS attack event.
	Events []*DescribeOnDemandDdosEventResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6A507DC8-F657-4C13-84E2-D1D1B9400753
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of DDoS attack events that are returned.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeOnDemandDdosEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandDdosEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandDdosEventResponseBody) SetEvents(v []*DescribeOnDemandDdosEventResponseBodyEvents) *DescribeOnDemandDdosEventResponseBody {
	s.Events = v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBody) SetRequestId(v string) *DescribeOnDemandDdosEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBody) SetTotal(v int64) *DescribeOnDemandDdosEventResponseBody {
	s.Total = &v
	return s
}

type DescribeOnDemandDdosEventResponseBodyEvents struct {
	// The time when the DDoS attack stopped. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1557891306
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The attacked IP address.
	//
	// example:
	//
	// 192.XX.XX.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The attack traffic. Unit: Mbit/s.
	//
	// example:
	//
	// 110000
	Mbps *int32 `json:"Mbps,omitempty" xml:"Mbps,omitempty"`
	// The packet forwarding rate of the DDoS attack. Unit: packets per second (PPS).
	//
	// example:
	//
	// 0
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The time when the DDoS attack started. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1557889506
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the DDoS attack event. Valid values:
	//
	// 	- **hole_begin**: indicates that blackhole filtering is triggered.
	//
	// 	- **hole_end**: indicates that tblackhole filtering is deactivated.
	//
	// 	- **defense_begin**: indicates that traffic scrubbing is in progress.
	//
	// 	- **defense_end**: indicates that traffic scrubbing is complete.
	//
	// example:
	//
	// defense_end
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeOnDemandDdosEventResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandDdosEventResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetEndTime(v int32) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetIp(v string) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.Ip = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetMbps(v int32) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.Mbps = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetPps(v int32) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.Pps = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetStartTime(v int32) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetStatus(v string) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.Status = &v
	return s
}

type DescribeOnDemandDdosEventResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOnDemandDdosEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOnDemandDdosEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandDdosEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandDdosEventResponse) SetHeaders(v map[string]*string) *DescribeOnDemandDdosEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeOnDemandDdosEventResponse) SetStatusCode(v int32) *DescribeOnDemandDdosEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponse) SetBody(v *DescribeOnDemandDdosEventResponseBody) *DescribeOnDemandDdosEventResponse {
	s.Body = v
	return s
}

type DescribeOnDemandInstanceStatusRequest struct {
	// The IDs of the anti-DDoS diversion instances.
	//
	// >  You can call the [DescribeOnDemandInstance](https://help.aliyun.com/document_detail/152120.html) operation to query the IDs of all anti-DDoS diversion instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-z2q1qzxb****
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	// The region ID of the anti-DDoS diversion instance.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query all regions that are supported by Anti-DDoS Origin.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOnDemandInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceStatusRequest) SetInstanceIdList(v []*string) *DescribeOnDemandInstanceStatusRequest {
	s.InstanceIdList = v
	return s
}

func (s *DescribeOnDemandInstanceStatusRequest) SetRegionId(v string) *DescribeOnDemandInstanceStatusRequest {
	s.RegionId = &v
	return s
}

type DescribeOnDemandInstanceStatusResponseBody struct {
	// The details of the anti-DDoS diversion instance.
	Instances []*DescribeOnDemandInstanceStatusResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// CC49FF51-612F-429B-AB1E-374B3F115396
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOnDemandInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceStatusResponseBody) SetInstances(v []*DescribeOnDemandInstanceStatusResponseBodyInstances) *DescribeOnDemandInstanceStatusResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBody) SetRequestId(v string) *DescribeOnDemandInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOnDemandInstanceStatusResponseBodyInstances struct {
	// The details of route advertisement for data centers outside the Chinese mainland. This parameter is a JSON string. The following fields are included in the value:
	//
	// 	- **region**: The code of the data center outside the Chinese mainland. The value is of the string type. For more information, see **Codes of data centers outside the Chinese mainland**.
	//
	// 	- **declared**: indicates whether the data center advertised the route. The value is of the string type. Valid values: **0*	- and **1**. The value of 0 indicates that the data center did not advertise the route. The value of 1 indicates that the data center advertised the route.
	//
	// example:
	//
	// [{\\"region\\":\\"oe24\\",\\"declared\\":0},{\\"region\\":\\"oe26\\",\\"declared\\":0},{\\"region\\":\\"oe28\\",\\"declared\\":0},{\\"region\\":\\"oi39\\",\\"declared\\":0},{\\"region\\":\\"us50\\",\\"declared\\":0},{\\"region\\":\\"jp141\\",\\"declared\\":0}]
	Declared *string `json:"Declared,omitempty" xml:"Declared,omitempty"`
	// The description of the anti-DDoS diversion instance.
	//
	// > This parameter is returned only when the information about multiple anti-DDoS diversion instances are returned. The value of this parameter is not returned because the information about only one anti-DDoS diversion instance is returned.
	//
	// example:
	//
	// test
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The ID of the anti-DDoS diversion instance.
	//
	// > This parameter is returned only when the information about multiple anti-DDoS diversion instances are returned. The value of this parameter is not returned because the information about only one anti-DDoS diversion instance is returned.
	//
	// example:
	//
	// ddosbgp-cn-z2q1qzxb****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The mode that is used to enable traffic rerouting to the anti-DDoS diversion instance. Valid values:
	//
	// 	- **manual**: The instance is manually started.
	//
	// 	- **netflow-auto**: The instance is automatically started by using NetFlow that monitors network traffic.
	//
	// example:
	//
	// netflow-auto
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The CIDR block of the anti-DDoS diversion instance.
	//
	// example:
	//
	// 47.***.***.0/24
	Net *string `json:"Net,omitempty" xml:"Net,omitempty"`
	// The number of the autonomous system (AS). Set the value to **0**, which indicates that AS is disabled.
	//
	// example:
	//
	// 0
	RegistedAs *string `json:"RegistedAs,omitempty" xml:"RegistedAs,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 171986973287****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeOnDemandInstanceStatusResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceStatusResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetDeclared(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.Declared = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetDesc(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.Desc = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetInstanceId(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetMode(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.Mode = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetNet(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.Net = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetRegistedAs(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.RegistedAs = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetUserId(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.UserId = &v
	return s
}

type DescribeOnDemandInstanceStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOnDemandInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOnDemandInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceStatusResponse) SetHeaders(v map[string]*string) *DescribeOnDemandInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponse) SetStatusCode(v int32) *DescribeOnDemandInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponse) SetBody(v *DescribeOnDemandInstanceStatusResponseBody) *DescribeOnDemandInstanceStatusResponse {
	s.Body = v
	return s
}

type DescribeOpEntitiesRequest struct {
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time. Operation logs that were generated before this time are queried.***	- The value is a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1640880000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance to query.
	//
	// > You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all instances.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OpAction   *int32  `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	// The sorting method of operation logs. Set the value to **opdate**, which indicates sorting based on the operation time.
	//
	// example:
	//
	// opdate
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The sort order of operation logs. Valid values:
	//
	// 	- **ASC**: the ascending order.
	//
	// 	- **DESC**: the descending order.
	//
	// Default value: **DESC**.
	//
	// example:
	//
	// ASC
	OrderDir *string `json:"OrderDir,omitempty" xml:"OrderDir,omitempty"`
	// The number of entries per page. Maximum value: 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the instance resides.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The start time. Operation logs that were generated after this time are queried.***	- The value is a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1609430400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOpEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesRequest) SetCurrentPage(v int32) *DescribeOpEntitiesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEndTime(v int64) *DescribeOpEntitiesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetInstanceId(v string) *DescribeOpEntitiesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetOpAction(v int32) *DescribeOpEntitiesRequest {
	s.OpAction = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetOrderBy(v string) *DescribeOpEntitiesRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetOrderDir(v string) *DescribeOpEntitiesRequest {
	s.OrderDir = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageSize(v int32) *DescribeOpEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetRegionId(v string) *DescribeOpEntitiesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetResourceGroupId(v string) *DescribeOpEntitiesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetStartTime(v int64) *DescribeOpEntitiesRequest {
	s.StartTime = &v
	return s
}

type DescribeOpEntitiesResponseBody struct {
	// The details of the operation log.
	OpEntities []*DescribeOpEntitiesResponseBodyOpEntities `json:"OpEntities,omitempty" xml:"OpEntities,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 52C8ECB0-0B1A-4E66-A31C-B6A855120E82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of operation logs.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOpEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBody) SetOpEntities(v []*DescribeOpEntitiesResponseBodyOpEntities) *DescribeOpEntitiesResponseBody {
	s.OpEntities = v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetRequestId(v string) *DescribeOpEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetTotalCount(v int32) *DescribeOpEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeOpEntitiesResponseBodyOpEntities struct {
	// The operation object, which is the ID of the instance.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	// The type of the operation object. The value is fixed as **1**, which indicates Anti-DDoS Origin instances.
	//
	// example:
	//
	// 1
	EntityType *int32 `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The time when the log was generated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1635818114000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the Alibaba Cloud account that performs the operation.
	//
	// > If the value is **system**, the operation is performed by Anti-DDoS Origin.
	//
	// example:
	//
	// 171986973287****
	OpAccount *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty"`
	// The type of operation. Valid values:
	//
	// 	- **3**: indicates an operation to add an IP address to the Anti-DDoS Origin instance for protection.
	//
	// 	- **4**: indicates an operation to remove a protected IP address from the Anti-DDoS Origin instance.
	//
	// 	- **5**: indicates an operation to downgrade the Anti-DDoS Origin instance.
	//
	// 	- **6**: indicates an operation to deactivate blackhole filtering for an IP address.
	//
	// 	- **7**: indicates an operation to reset the number of times that you can deactivate blackhole filtering.
	//
	// 	- **8**: indicates an operation to enable burstable protection.
	//
	// example:
	//
	// 8
	OpAction *int32 `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	// The details of the operation. The value is a string that consists of a JSON struct. The JSON struct contains the following fields:
	//
	// 	- **entity**: the operation object. Data type: object. The fields that are included in the value of the **entity*	- parameter vary based on the value of the **OpAction*	- parameter. Valid values:
	//
	//     	- If the value of the **OpAction*	- parameter is **3**, the value of the **entity*	- parameter consists of the following field:
	//
	//         	- **ips**: the public IP addresses that are protected by the Anti-DDoS Origin instance. Data type: array
	//
	//     	- If the value of the **OpAction*	- parameter is **4**, the value of the **entity*	- parameter consists of the following field:
	//
	//         	- **ips**: the public IP addresses that are no longer protected by the Anti-DDoS Origin instance. Data type: array.
	//
	//     	- If the value of the **OpAction*	- parameter is **5**, the value of the **entity*	- parameter consists of the following fields:
	//
	//         	- **baseBandwidth**: the basic protection bandwidth. Unit: Gbit/s. Data type: integer.
	//
	//         	- **elasticBandwidth**: the burstable protection bandwidth. Unit: Gbit/s. Data type: integer.
	//
	//         	- **opSource**: the source of the operation. The value is fixed as **1**, indicating that the operation is performed by Anti-DDoS Origin. Data type: integer.
	//
	//     	- If the value of the **OpAction*	- parameter is **6**, the value of the **entity*	- parameter consists of the following field:
	//
	//         	- **ips**: the public IP addresses for which to deactivate blackhole filtering. Data type: array.
	//
	//     	- If the value of the **OpAction*	- parameter is **7**, the **entity*	- parameter is not returned.
	//
	//     	- If the value of the **OpAction*	- parameter is **8**, the value of the **entity*	- parameter consists of the following fields:
	//
	//         	- **baseBandwidth**: the basic protection bandwidth. Unit: Gbit/s. Data type: integer.
	//
	//         	- **elasticBandwidth**: the burstable protection bandwidth. Unit: Gbit/s. Data type: integer.
	//
	// example:
	//
	// {"entity":{"baseBandwidth":20,"elasticBandwidth":20}}
	OpDesc *string `json:"OpDesc,omitempty" xml:"OpDesc,omitempty"`
}

func (s DescribeOpEntitiesResponseBodyOpEntities) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseBodyOpEntities) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityObject(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityType(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetGmtCreate(v int64) *DescribeOpEntitiesResponseBodyOpEntities {
	s.GmtCreate = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAccount(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAccount = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAction(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAction = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpDesc(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpDesc = &v
	return s
}

type DescribeOpEntitiesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponse) SetHeaders(v map[string]*string) *DescribeOpEntitiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpEntitiesResponse) SetStatusCode(v int32) *DescribeOpEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpEntitiesResponse) SetBody(v *DescribeOpEntitiesResponseBody) *DescribeOpEntitiesResponse {
	s.Body = v
	return s
}

type DescribePackIpListRequest struct {
	// The ID of the Anti-DDoS Origin instance to query.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The protected IP address to query.
	//
	// example:
	//
	// 47.98.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the member.
	//
	// example:
	//
	// 170858869679****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the cloud asset to which the protected IP address to query belongs. Valid values:
	//
	// 	- **ECS**: an Elastic Compute Service (ECS) instance.
	//
	// 	- **SLB**: a Classic Load Balancer (CLB) instance, originally called a Server Load Balancer (SLB) instance.
	//
	// 	- **EIP**: an elastic IP address (EIP). An Internet-facing Application Load Balancer (ALB) instance uses an EIP. If the IP address belongs to the Internet-facing ALB instance, set this parameter to EIP.
	//
	// 	- **WAF**: a Web Application Firewall (WAF) instance.
	//
	// example:
	//
	// ECS
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribePackIpListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListRequest) GoString() string {
	return s.String()
}

func (s *DescribePackIpListRequest) SetInstanceId(v string) *DescribePackIpListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePackIpListRequest) SetIp(v string) *DescribePackIpListRequest {
	s.Ip = &v
	return s
}

func (s *DescribePackIpListRequest) SetMemberUid(v string) *DescribePackIpListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribePackIpListRequest) SetPageNo(v int32) *DescribePackIpListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePackIpListRequest) SetPageSize(v int32) *DescribePackIpListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePackIpListRequest) SetProductName(v string) *DescribePackIpListRequest {
	s.ProductName = &v
	return s
}

func (s *DescribePackIpListRequest) SetRegionId(v string) *DescribePackIpListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePackIpListRequest) SetResourceGroupId(v string) *DescribePackIpListRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribePackIpListResponseBody struct {
	// The HTTP status code of the request.
	//
	// For more information about status codes, see [Common parameters](https://help.aliyun.com/document_detail/118841.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The IP addresses that are protected by the instance.
	IpList []*DescribePackIpListResponseBodyIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 4FD1578A-BD77-50B7-A969-45A374A7ED22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The call is successful.
	//
	// 	- **false**: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The number of protected IP addresses.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribePackIpListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackIpListResponseBody) SetCode(v string) *DescribePackIpListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePackIpListResponseBody) SetIpList(v []*DescribePackIpListResponseBodyIpList) *DescribePackIpListResponseBody {
	s.IpList = v
	return s
}

func (s *DescribePackIpListResponseBody) SetRequestId(v string) *DescribePackIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackIpListResponseBody) SetSuccess(v bool) *DescribePackIpListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePackIpListResponseBody) SetTotal(v int32) *DescribePackIpListResponseBody {
	s.Total = &v
	return s
}

type DescribePackIpListResponseBodyIpList struct {
	// The IP address.
	//
	// example:
	//
	// 47.98.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the member.
	//
	// example:
	//
	// 170858869679****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The time when the near-origin traffic diversion feature was disabled.
	//
	// example:
	//
	// 1715658000
	NsmExpireAt *int64 `json:"NsmExpireAt,omitempty" xml:"NsmExpireAt,omitempty"`
	// The time when the near-origin traffic diversion feature was enabled.
	//
	// example:
	//
	// 1715655000
	NsmStartAt *int64 `json:"NsmStartAt,omitempty" xml:"NsmStartAt,omitempty"`
	// The status of the near-origin traffic diversion feature. Valid values:
	//
	// 	- **1**: The near-origin traffic diversion feature is enabled.
	//
	// 	- **0**: The near-origin traffic diversion feature is disabled.
	//
	// example:
	//
	// 0
	NsmStatus *int32 `json:"NsmStatus,omitempty" xml:"NsmStatus,omitempty"`
	// The type of the cloud asset to which the IP address belongs. Valid values:
	//
	// 	- **ECS**: an ECS instance.
	//
	// 	- **SLB**: a CLB (formerly SLB) instance.
	//
	// 	- **EIP**: an EIP. If the IP address belongs to an ALB instance, the value EIP is returned.
	//
	// 	- **WAF**: a WAF instance.
	//
	// example:
	//
	// ECS
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region to which the protected IP address belongs.
	//
	// >  If the protected IP address is in the same region as the instance, this parameter is not returned.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The description of the cloud asset to which the IP address belongs. The asset can be an ECS instance or an SLB instance.
	//
	// >  If no descriptions are provided for the asset, this parameter is not returned.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the IP address. Valid values:
	//
	// 	- **normal**: The IP address is not under attack.
	//
	// 	- **hole_begin**: Blackhole filtering is triggered for the IP address.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePackIpListResponseBodyIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListResponseBodyIpList) GoString() string {
	return s.String()
}

func (s *DescribePackIpListResponseBodyIpList) SetIp(v string) *DescribePackIpListResponseBodyIpList {
	s.Ip = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetMemberUid(v string) *DescribePackIpListResponseBodyIpList {
	s.MemberUid = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetNsmExpireAt(v int64) *DescribePackIpListResponseBodyIpList {
	s.NsmExpireAt = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetNsmStartAt(v int64) *DescribePackIpListResponseBodyIpList {
	s.NsmStartAt = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetNsmStatus(v int32) *DescribePackIpListResponseBodyIpList {
	s.NsmStatus = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetProduct(v string) *DescribePackIpListResponseBodyIpList {
	s.Product = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetRegion(v string) *DescribePackIpListResponseBodyIpList {
	s.Region = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetRemark(v string) *DescribePackIpListResponseBodyIpList {
	s.Remark = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetStatus(v string) *DescribePackIpListResponseBodyIpList {
	s.Status = &v
	return s
}

type DescribePackIpListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePackIpListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePackIpListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListResponse) GoString() string {
	return s.String()
}

func (s *DescribePackIpListResponse) SetHeaders(v map[string]*string) *DescribePackIpListResponse {
	s.Headers = v
	return s
}

func (s *DescribePackIpListResponse) SetStatusCode(v int32) *DescribePackIpListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackIpListResponse) SetBody(v *DescribePackIpListResponseBody) *DescribePackIpListResponse {
	s.Body = v
	return s
}

type DescribeRdMemberListRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-x9bLhd
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
}

func (s DescribeRdMemberListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdMemberListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdMemberListRequest) SetPageNo(v int32) *DescribeRdMemberListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeRdMemberListRequest) SetPageSize(v int32) *DescribeRdMemberListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRdMemberListRequest) SetResourceDirectoryId(v string) *DescribeRdMemberListRequest {
	s.ResourceDirectoryId = &v
	return s
}

type DescribeRdMemberListResponseBody struct {
	// The list of the members.
	MemberList []*DescribeRdMemberListResponseBodyMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DC245DEE-9800-5579-BF99-189D6A5BA9FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRdMemberListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdMemberListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdMemberListResponseBody) SetMemberList(v []*DescribeRdMemberListResponseBodyMemberList) *DescribeRdMemberListResponseBody {
	s.MemberList = v
	return s
}

func (s *DescribeRdMemberListResponseBody) SetRequestId(v string) *DescribeRdMemberListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdMemberListResponseBody) SetTotal(v int64) *DescribeRdMemberListResponseBody {
	s.Total = &v
	return s
}

type DescribeRdMemberListResponseBodyMemberList struct {
	// The creation time.
	//
	// example:
	//
	// 1624954942000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The name of the member.
	//
	// example:
	//
	// test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 1960279802016267
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeRdMemberListResponseBodyMemberList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdMemberListResponseBodyMemberList) GoString() string {
	return s.String()
}

func (s *DescribeRdMemberListResponseBodyMemberList) SetGmtCreate(v int64) *DescribeRdMemberListResponseBodyMemberList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdMemberListResponseBodyMemberList) SetName(v string) *DescribeRdMemberListResponseBodyMemberList {
	s.Name = &v
	return s
}

func (s *DescribeRdMemberListResponseBodyMemberList) SetUid(v string) *DescribeRdMemberListResponseBodyMemberList {
	s.Uid = &v
	return s
}

type DescribeRdMemberListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRdMemberListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRdMemberListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdMemberListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdMemberListResponse) SetHeaders(v map[string]*string) *DescribeRdMemberListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdMemberListResponse) SetStatusCode(v int32) *DescribeRdMemberListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdMemberListResponse) SetBody(v *DescribeRdMemberListResponseBody) *DescribeRdMemberListResponse {
	s.Body = v
	return s
}

type DescribeRdStatusResponseBody struct {
	// The Alibaba Cloud account ID of the current account.
	//
	// example:
	//
	// 125085778340****
	CurrentUid *string `json:"CurrentUid,omitempty" xml:"CurrentUid,omitempty"`
	// The type of the Alibaba Cloud account. Valid values:
	//
	// 	- **MasterAccount**: management account
	//
	// 	- **DelegatedAdminAccount**: delegated administrator account
	//
	// 	- **MasterAccount**: member
	//
	// example:
	//
	// MemberAccount
	CurrentUidType *string `json:"CurrentUidType,omitempty" xml:"CurrentUidType,omitempty"`
	// Indicates whether the multi-account management feature is enabled for Anti-DDoS Origin.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether the multi-account management feature is enabled for the current account in Anti-DDoS Origin.
	//
	// example:
	//
	// false
	LocalEnable *bool `json:"LocalEnable,omitempty" xml:"LocalEnable,omitempty"`
	// The Alibaba Cloud account ID of the management account in the resource directory.
	//
	// example:
	//
	// 125085778340****
	MasterUid *string `json:"MasterUid,omitempty" xml:"MasterUid,omitempty"`
	// Indicates whether Resource Directory is enabled in the [Resource Management console](https://resourcemanager.console.aliyun.com).
	//
	// example:
	//
	// false
	RemoteEnable *bool `json:"RemoteEnable,omitempty" xml:"RemoteEnable,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1B0F7EC6-51D7-4D70-B0EC-CD8A9E998D86
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud account ID of the management account for which the multi-account management feature is enabled in Anti-DDoS Origin.
	//
	// example:
	//
	// 125085778340****
	RootUid *string `json:"RootUid,omitempty" xml:"RootUid,omitempty"`
	// Indicates whether the trusted service is enabled.
	//
	// example:
	//
	// false
	ServicePrincipalEnabled *bool `json:"ServicePrincipalEnabled,omitempty" xml:"ServicePrincipalEnabled,omitempty"`
}

func (s DescribeRdStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdStatusResponseBody) SetCurrentUid(v string) *DescribeRdStatusResponseBody {
	s.CurrentUid = &v
	return s
}

func (s *DescribeRdStatusResponseBody) SetCurrentUidType(v string) *DescribeRdStatusResponseBody {
	s.CurrentUidType = &v
	return s
}

func (s *DescribeRdStatusResponseBody) SetEnabled(v bool) *DescribeRdStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeRdStatusResponseBody) SetLocalEnable(v bool) *DescribeRdStatusResponseBody {
	s.LocalEnable = &v
	return s
}

func (s *DescribeRdStatusResponseBody) SetMasterUid(v string) *DescribeRdStatusResponseBody {
	s.MasterUid = &v
	return s
}

func (s *DescribeRdStatusResponseBody) SetRemoteEnable(v bool) *DescribeRdStatusResponseBody {
	s.RemoteEnable = &v
	return s
}

func (s *DescribeRdStatusResponseBody) SetRequestId(v string) *DescribeRdStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdStatusResponseBody) SetRootUid(v string) *DescribeRdStatusResponseBody {
	s.RootUid = &v
	return s
}

func (s *DescribeRdStatusResponseBody) SetServicePrincipalEnabled(v bool) *DescribeRdStatusResponseBody {
	s.ServicePrincipalEnabled = &v
	return s
}

type DescribeRdStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRdStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRdStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdStatusResponse) SetHeaders(v map[string]*string) *DescribeRdStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdStatusResponse) SetStatusCode(v int32) *DescribeRdStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdStatusResponse) SetBody(v *DescribeRdStatusResponseBody) *DescribeRdStatusResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The ID of the region. The default value is **cn-hangzhou**. If the default value is used, the regions of cloud assets that can be protected by Anti-DDoS Origin in the China (Hangzhou) region are queried.
	//
	// If you want to specify another value for **RegionId**, see [Regions and Zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. This parameter is empty by default, which indicates that the Anti-DDoS Origin instance belongs to the default resource group.
	//
	// For information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *DescribeRegionsRequest) SetResourceGroupId(v string) *DescribeRegionsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the regions of cloud assets that can be protected by Anti-DDoS Origin. The information includes region IDs and names.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F7CA8B4E-FB15-4336-A351-8DC29D66EA82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetCode(v string) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	// The English name of the region.
	//
	// example:
	//
	// China (Hangzhou)
	RegionEnName *string `json:"RegionEnName,omitempty" xml:"RegionEnName,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Chinese name of the region.
	//
	// example:
	//
	// 华东1（杭州）
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEnName(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEnName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionName(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionName = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeTrafficRequest struct {
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// If you do not specify this parameter, the current system time is used as the end time.
	//
	// example:
	//
	// 1563445054
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the traffic statistics to query. Valid values:
	//
	// 	- **max**: the peak traffic within the specified interval
	//
	// 	- **avg**: the average traffic within the specified interval
	//
	// example:
	//
	// max
	FlowType *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	// The ID of the Anti-DDoS Origin instance to query.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// If you specify an on-demand instance, you must configure the **Interval*	- parameter.
	//
	// example:
	//
	// ddosbgp-cn-n6w203qg****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The interval at which the traffic statistics are calculated. Unit: seconds. Default value: **5**.
	//
	// example:
	//
	// 5
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The public IP address of the asset to query. If you do not specify this parameter, the traffic statistics of all public IP addresses that are protected by the Anti-DDoS Origin instance are queried.
	//
	// >  The public IP address must be a protected object of the Anti-DDoS Origin instance. You can call the [DescribePackIpList](https://help.aliyun.com/document_detail/118701.html) operation to query all protected objects of the Anti-DDoS Origin instance.
	//
	// example:
	//
	// 39.XX.XX.96
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The Classless Inter-Domain Routing (CIDR) block of the on-demand instance that you want to query.
	//
	// example:
	//
	// 111.XX.XX.0/24
	Ipnet *string `json:"Ipnet,omitempty" xml:"Ipnet,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1619798400
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrafficRequest) SetEndTime(v int32) *DescribeTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTrafficRequest) SetFlowType(v string) *DescribeTrafficRequest {
	s.FlowType = &v
	return s
}

func (s *DescribeTrafficRequest) SetInstanceId(v string) *DescribeTrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTrafficRequest) SetInterval(v int32) *DescribeTrafficRequest {
	s.Interval = &v
	return s
}

func (s *DescribeTrafficRequest) SetIp(v string) *DescribeTrafficRequest {
	s.Ip = &v
	return s
}

func (s *DescribeTrafficRequest) SetIpnet(v string) *DescribeTrafficRequest {
	s.Ipnet = &v
	return s
}

func (s *DescribeTrafficRequest) SetRegionId(v string) *DescribeTrafficRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTrafficRequest) SetResourceGroupId(v string) *DescribeTrafficRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTrafficRequest) SetStartTime(v int32) *DescribeTrafficRequest {
	s.StartTime = &v
	return s
}

type DescribeTrafficResponseBody struct {
	// The queried traffic statistics.
	FlowList []*DescribeTrafficResponseBodyFlowList `json:"FlowList,omitempty" xml:"FlowList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 6A507DC8-F657-4C13-84E2-D1D1B9400753
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTrafficResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponseBody) SetFlowList(v []*DescribeTrafficResponseBodyFlowList) *DescribeTrafficResponseBody {
	s.FlowList = v
	return s
}

func (s *DescribeTrafficResponseBody) SetRequestId(v string) *DescribeTrafficResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTrafficResponseBodyFlowList struct {
	// The bandwidth of attack traffic. Unit: bit/s.
	//
	// >  This parameter is returned only if attack traffic exists.
	//
	// example:
	//
	// 0
	AttackBps *int64 `json:"AttackBps,omitempty" xml:"AttackBps,omitempty"`
	// The packet forwarding rate of attack traffic. Unit: packets per second.
	//
	// >  This parameter is returned only if attack traffic exists.
	//
	// example:
	//
	// 0
	AttackPps *int64 `json:"AttackPps,omitempty" xml:"AttackPps,omitempty"`
	// The type of the traffic statistics. Valid values:
	//
	// 	- **max**: the peak traffic within the specified interval
	//
	// 	- **avg**: the average traffic within the specified interval
	//
	// example:
	//
	// max
	FlowType *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	// The bandwidth of the total traffic. Unit: Kbit/s.
	//
	// example:
	//
	// 417
	Kbps *int32 `json:"Kbps,omitempty" xml:"Kbps,omitempty"`
	// The ID of the traffic statistics.
	//
	// example:
	//
	// 8e33f19e-5644-11eb-b5c1-d89d67182200
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The packet forwarding rate of the total traffic. Unit: packets per second.
	//
	// example:
	//
	// 274
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The time when the traffic statistics are calculated. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1620951900
	Time *int32 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeTrafficResponseBodyFlowList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponseBodyFlowList) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponseBodyFlowList) SetAttackBps(v int64) *DescribeTrafficResponseBodyFlowList {
	s.AttackBps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetAttackPps(v int64) *DescribeTrafficResponseBodyFlowList {
	s.AttackPps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetFlowType(v string) *DescribeTrafficResponseBodyFlowList {
	s.FlowType = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetKbps(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Kbps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetName(v string) *DescribeTrafficResponseBodyFlowList {
	s.Name = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetPps(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Pps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetTime(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Time = &v
	return s
}

type DescribeTrafficResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponse) SetHeaders(v map[string]*string) *DescribeTrafficResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrafficResponse) SetStatusCode(v int32) *DescribeTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrafficResponse) SetBody(v *DescribeTrafficResponseBody) *DescribeTrafficResponse {
	s.Body = v
	return s
}

type DetachFromPolicyRequest struct {
	// The protected objects.
	//
	// This parameter is required.
	IpPortProtocolList []*DetachFromPolicyRequestIpPortProtocolList `json:"IpPortProtocolList,omitempty" xml:"IpPortProtocolList,omitempty" type:"Repeated"`
	// The type of the policy. Valid values:
	//
	// 	- **default**: the default mitigation policies.
	//
	// 	- **l3**: IP-specific mitigation policies.
	//
	// 	- **l4**: port-specific mitigation policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// l3
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DetachFromPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachFromPolicyRequest) GoString() string {
	return s.String()
}

func (s *DetachFromPolicyRequest) SetIpPortProtocolList(v []*DetachFromPolicyRequestIpPortProtocolList) *DetachFromPolicyRequest {
	s.IpPortProtocolList = v
	return s
}

func (s *DetachFromPolicyRequest) SetPolicyType(v string) *DetachFromPolicyRequest {
	s.PolicyType = &v
	return s
}

type DetachFromPolicyRequestIpPortProtocolList struct {
	// The IP address of the protected object.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47.118.172.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port of the protected object.
	//
	// example:
	//
	// 8*
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type of the protected object. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DetachFromPolicyRequestIpPortProtocolList) String() string {
	return tea.Prettify(s)
}

func (s DetachFromPolicyRequestIpPortProtocolList) GoString() string {
	return s.String()
}

func (s *DetachFromPolicyRequestIpPortProtocolList) SetIp(v string) *DetachFromPolicyRequestIpPortProtocolList {
	s.Ip = &v
	return s
}

func (s *DetachFromPolicyRequestIpPortProtocolList) SetPort(v int32) *DetachFromPolicyRequestIpPortProtocolList {
	s.Port = &v
	return s
}

func (s *DetachFromPolicyRequestIpPortProtocolList) SetProtocol(v string) *DetachFromPolicyRequestIpPortProtocolList {
	s.Protocol = &v
	return s
}

type DetachFromPolicyShrinkRequest struct {
	// The protected objects.
	//
	// This parameter is required.
	IpPortProtocolListShrink *string `json:"IpPortProtocolList,omitempty" xml:"IpPortProtocolList,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **default**: the default mitigation policies.
	//
	// 	- **l3**: IP-specific mitigation policies.
	//
	// 	- **l4**: port-specific mitigation policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// l3
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DetachFromPolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachFromPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetachFromPolicyShrinkRequest) SetIpPortProtocolListShrink(v string) *DetachFromPolicyShrinkRequest {
	s.IpPortProtocolListShrink = &v
	return s
}

func (s *DetachFromPolicyShrinkRequest) SetPolicyType(v string) *DetachFromPolicyShrinkRequest {
	s.PolicyType = &v
	return s
}

type DetachFromPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1B0F7EC6-51D7-4D70-B0EC-CD8A9E99****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachFromPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachFromPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DetachFromPolicyResponseBody) SetRequestId(v string) *DetachFromPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DetachFromPolicyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachFromPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachFromPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachFromPolicyResponse) GoString() string {
	return s.String()
}

func (s *DetachFromPolicyResponse) SetHeaders(v map[string]*string) *DetachFromPolicyResponse {
	s.Headers = v
	return s
}

func (s *DetachFromPolicyResponse) SetStatusCode(v int32) *DetachFromPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachFromPolicyResponse) SetBody(v *DetachFromPolicyResponseBody) *DetachFromPolicyResponse {
	s.Body = v
	return s
}

type DettachAssetGroupToInstanceRequest struct {
	// The information about the asset that you want to dissociate.
	//
	// This parameter is required.
	AssetGroupList []*DettachAssetGroupToInstanceRequestAssetGroupList `json:"AssetGroupList,omitempty" xml:"AssetGroupList,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances of paid editions.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DettachAssetGroupToInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DettachAssetGroupToInstanceRequest) GoString() string {
	return s.String()
}

func (s *DettachAssetGroupToInstanceRequest) SetAssetGroupList(v []*DettachAssetGroupToInstanceRequestAssetGroupList) *DettachAssetGroupToInstanceRequest {
	s.AssetGroupList = v
	return s
}

func (s *DettachAssetGroupToInstanceRequest) SetInstanceId(v string) *DettachAssetGroupToInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DettachAssetGroupToInstanceRequest) SetRegionId(v string) *DettachAssetGroupToInstanceRequest {
	s.RegionId = &v
	return s
}

type DettachAssetGroupToInstanceRequestAssetGroupList struct {
	// The ID of the asset. If the asset is a Web Application Firewall (WAF) instance, specify the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-lbj382l****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- **waf**: WAF instance
	//
	// 	- **ga**: Global Accelerator (GA) instance
	//
	// This parameter is required.
	//
	// example:
	//
	// waf
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DettachAssetGroupToInstanceRequestAssetGroupList) String() string {
	return tea.Prettify(s)
}

func (s DettachAssetGroupToInstanceRequestAssetGroupList) GoString() string {
	return s.String()
}

func (s *DettachAssetGroupToInstanceRequestAssetGroupList) SetName(v string) *DettachAssetGroupToInstanceRequestAssetGroupList {
	s.Name = &v
	return s
}

func (s *DettachAssetGroupToInstanceRequestAssetGroupList) SetRegion(v string) *DettachAssetGroupToInstanceRequestAssetGroupList {
	s.Region = &v
	return s
}

func (s *DettachAssetGroupToInstanceRequestAssetGroupList) SetType(v string) *DettachAssetGroupToInstanceRequestAssetGroupList {
	s.Type = &v
	return s
}

type DettachAssetGroupToInstanceShrinkRequest struct {
	// The information about the asset that you want to dissociate.
	//
	// This parameter is required.
	AssetGroupListShrink *string `json:"AssetGroupList,omitempty" xml:"AssetGroupList,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances of paid editions.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DettachAssetGroupToInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DettachAssetGroupToInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DettachAssetGroupToInstanceShrinkRequest) SetAssetGroupListShrink(v string) *DettachAssetGroupToInstanceShrinkRequest {
	s.AssetGroupListShrink = &v
	return s
}

func (s *DettachAssetGroupToInstanceShrinkRequest) SetInstanceId(v string) *DettachAssetGroupToInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DettachAssetGroupToInstanceShrinkRequest) SetRegionId(v string) *DettachAssetGroupToInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

type DettachAssetGroupToInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E54BA258-9DE8-59BE-B7A8-DAD28E6E8DAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DettachAssetGroupToInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DettachAssetGroupToInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DettachAssetGroupToInstanceResponseBody) SetRequestId(v string) *DettachAssetGroupToInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DettachAssetGroupToInstanceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DettachAssetGroupToInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DettachAssetGroupToInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DettachAssetGroupToInstanceResponse) GoString() string {
	return s.String()
}

func (s *DettachAssetGroupToInstanceResponse) SetHeaders(v map[string]*string) *DettachAssetGroupToInstanceResponse {
	s.Headers = v
	return s
}

func (s *DettachAssetGroupToInstanceResponse) SetStatusCode(v int32) *DettachAssetGroupToInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DettachAssetGroupToInstanceResponse) SetBody(v *DettachAssetGroupToInstanceResponseBody) *DettachAssetGroupToInstanceResponse {
	s.Body = v
	return s
}

type GetSlsOpenStatusRequest struct {
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// For more information about the valid values of this parameter, see [Regions and zones](https://help.aliyun.com/document_detail/188196.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetSlsOpenStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSlsOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *GetSlsOpenStatusRequest) SetRegionId(v string) *GetSlsOpenStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetSlsOpenStatusRequest) SetResourceGroupId(v string) *GetSlsOpenStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type GetSlsOpenStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D01666F5-541B-4C78-98A6-D29E02DAAC7C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether Simple Log Service was activated. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	SlsOpenStatus *bool `json:"SlsOpenStatus,omitempty" xml:"SlsOpenStatus,omitempty"`
}

func (s GetSlsOpenStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSlsOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSlsOpenStatusResponseBody) SetRequestId(v string) *GetSlsOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSlsOpenStatusResponseBody) SetSlsOpenStatus(v bool) *GetSlsOpenStatusResponseBody {
	s.SlsOpenStatus = &v
	return s
}

type GetSlsOpenStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSlsOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSlsOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSlsOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSlsOpenStatusResponse) SetHeaders(v map[string]*string) *GetSlsOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSlsOpenStatusResponse) SetStatusCode(v int32) *GetSlsOpenStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSlsOpenStatusResponse) SetBody(v *GetSlsOpenStatusResponseBody) *GetSlsOpenStatusResponse {
	s.Body = v
	return s
}

type ListOpenedAccessLogInstancesRequest struct {
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListOpenedAccessLogInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOpenedAccessLogInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesRequest) SetPageNumber(v int32) *ListOpenedAccessLogInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOpenedAccessLogInstancesRequest) SetPageSize(v int32) *ListOpenedAccessLogInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOpenedAccessLogInstancesRequest) SetResourceGroupId(v string) *ListOpenedAccessLogInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

type ListOpenedAccessLogInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB64811-70A1-41C9-A0CE-CD8B260ED551
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configuration of log analysis for the Anti-DDoS Origin instance.
	SlsConfigStatus []*ListOpenedAccessLogInstancesResponseBodySlsConfigStatus `json:"SlsConfigStatus,omitempty" xml:"SlsConfigStatus,omitempty" type:"Repeated"`
	// The number of the Anti-DDoS Origin instances for which log analysis was enabled.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOpenedAccessLogInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOpenedAccessLogInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesResponseBody) SetRequestId(v string) *ListOpenedAccessLogInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponseBody) SetSlsConfigStatus(v []*ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) *ListOpenedAccessLogInstancesResponseBody {
	s.SlsConfigStatus = v
	return s
}

func (s *ListOpenedAccessLogInstancesResponseBody) SetTotalCount(v int32) *ListOpenedAccessLogInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListOpenedAccessLogInstancesResponseBodySlsConfigStatus struct {
	// Indicates whether log analysis was enabled for the Anti-DDoS Origin instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The ID of the Anti-DDoS Origin instance.
	//
	// example:
	//
	// ddosbgp-cn-m7r1zce2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) String() string {
	return tea.Prettify(s)
}

func (s ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) SetEnable(v bool) *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus {
	s.Enable = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) SetInstanceId(v string) *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus {
	s.InstanceId = &v
	return s
}

type ListOpenedAccessLogInstancesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOpenedAccessLogInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOpenedAccessLogInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOpenedAccessLogInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesResponse) SetHeaders(v map[string]*string) *ListOpenedAccessLogInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListOpenedAccessLogInstancesResponse) SetStatusCode(v int32) *ListOpenedAccessLogInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponse) SetBody(v *ListOpenedAccessLogInstancesResponseBody) *ListOpenedAccessLogInstancesResponse {
	s.Body = v
	return s
}

type ListPolicyRequest struct {
	// The name of the policy.
	//
	// example:
	//
	// test**
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The service type. Valid values:
	//
	// 	- **ecs**: Elastic Compute Service (ECS).
	//
	// 	- **slb**: Server Load Balancer (SLB).
	//
	// 	- **eip**: Elastic IP Address (EIP).
	//
	// 	- **gf-eip**: EIP with Anti-DDoS (Enhanced) enabled.
	//
	// >  This parameter is available only if Type is set to `default`.
	//
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **default**: the default mitigation policy.
	//
	// 	- **l3**: IP-specific mitigation policies.
	//
	// 	- **l4**: port-specific mitigation policies.
	//
	// example:
	//
	// l3
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyRequest) SetName(v string) *ListPolicyRequest {
	s.Name = &v
	return s
}

func (s *ListPolicyRequest) SetPageNo(v int64) *ListPolicyRequest {
	s.PageNo = &v
	return s
}

func (s *ListPolicyRequest) SetPageSize(v int64) *ListPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *ListPolicyRequest) SetProductType(v string) *ListPolicyRequest {
	s.ProductType = &v
	return s
}

func (s *ListPolicyRequest) SetType(v string) *ListPolicyRequest {
	s.Type = &v
	return s
}

type ListPolicyResponseBody struct {
	// The policies.
	PolicyList []*ListPolicyResponseBodyPolicyList `json:"PolicyList,omitempty" xml:"PolicyList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B4B379C2-9319-4C6B-B579-FE36831B09F4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of policies.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBody) SetPolicyList(v []*ListPolicyResponseBodyPolicyList) *ListPolicyResponseBody {
	s.PolicyList = v
	return s
}

func (s *ListPolicyResponseBody) SetRequestId(v string) *ListPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicyResponseBody) SetTotal(v int64) *ListPolicyResponseBody {
	s.Total = &v
	return s
}

type ListPolicyResponseBodyPolicyList struct {
	// The number of protected objects that are added to the policy.
	//
	// example:
	//
	// 0
	AttachedCount *int32 `json:"AttachedCount,omitempty" xml:"AttachedCount,omitempty"`
	// The content of the policy.
	Content *ListPolicyResponseBodyPolicyListContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the policy.
	//
	// example:
	//
	// 877afbdf-3982-4d36-9886-f043********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// test**
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The remarks of the policy.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **default**: the default mitigation policy.
	//
	// 	- **l3**: IP-specific mitigation policies.
	//
	// 	- **l4**: port-specific mitigation policies.
	//
	// example:
	//
	// l3
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPolicyResponseBodyPolicyList) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyList) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyList) SetAttachedCount(v int32) *ListPolicyResponseBodyPolicyList {
	s.AttachedCount = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyList) SetContent(v *ListPolicyResponseBodyPolicyListContent) *ListPolicyResponseBodyPolicyList {
	s.Content = v
	return s
}

func (s *ListPolicyResponseBodyPolicyList) SetId(v string) *ListPolicyResponseBodyPolicyList {
	s.Id = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyList) SetName(v string) *ListPolicyResponseBodyPolicyList {
	s.Name = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyList) SetRemark(v string) *ListPolicyResponseBodyPolicyList {
	s.Remark = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyList) SetType(v string) *ListPolicyResponseBodyPolicyList {
	s.Type = &v
	return s
}

type ListPolicyResponseBodyPolicyListContent struct {
	// The validity period of the IP address blacklist. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1716878000
	BlackIpListExpireAt *int64 `json:"BlackIpListExpireAt,omitempty" xml:"BlackIpListExpireAt,omitempty"`
	// Indicates whether ICMP blocking is enabled.
	//
	// example:
	//
	// false
	EnableDropIcmp *bool `json:"EnableDropIcmp,omitempty" xml:"EnableDropIcmp,omitempty"`
	// Indicates whether intelligent protection is enabled.
	//
	// example:
	//
	// true
	EnableIntelligence *bool `json:"EnableIntelligence,omitempty" xml:"EnableIntelligence,omitempty"`
	// Indicates whether port-specific mitigation is enabled.
	//
	// example:
	//
	// true
	EnableL4Defense *bool `json:"EnableL4Defense,omitempty" xml:"EnableL4Defense,omitempty"`
	// The byte-match filter rules.
	FingerPrintRuleList []*ListPolicyResponseBodyPolicyListContentFingerPrintRuleList `json:"FingerPrintRuleList,omitempty" xml:"FingerPrintRuleList,omitempty" type:"Repeated"`
	// The level of intelligent protection. Valid values:
	//
	// 	- **default**: normal.
	//
	// 	- **hard**: strict.
	//
	// 	- **weak**: loose.
	//
	// example:
	//
	// default
	IntelligenceLevel *string `json:"IntelligenceLevel,omitempty" xml:"IntelligenceLevel,omitempty"`
	// The port-specific mitigation rules.
	L4RuleList []*ListPolicyResponseBodyPolicyListContentL4RuleList `json:"L4RuleList,omitempty" xml:"L4RuleList,omitempty" type:"Repeated"`
	// The port blocking rules.
	PortRuleList []*ListPolicyResponseBodyPolicyListContentPortRuleList `json:"PortRuleList,omitempty" xml:"PortRuleList,omitempty" type:"Repeated"`
	// The ports whose traffic is filtered out by the filtering policies for UDP reflection attacks.
	ReflectBlockUdpPortList []*int32 `json:"ReflectBlockUdpPortList,omitempty" xml:"ReflectBlockUdpPortList,omitempty" type:"Repeated"`
	// The countries in the location blacklist.
	RegionBlockCountryList []*int32 `json:"RegionBlockCountryList,omitempty" xml:"RegionBlockCountryList,omitempty" type:"Repeated"`
	// The provinces in the location blacklist.
	RegionBlockProvinceList []*int32 `json:"RegionBlockProvinceList,omitempty" xml:"RegionBlockProvinceList,omitempty" type:"Repeated"`
	// The source IP addresses that are added to the blacklist.
	SourceBlockList []*ListPolicyResponseBodyPolicyListContentSourceBlockList `json:"SourceBlockList,omitempty" xml:"SourceBlockList,omitempty" type:"Repeated"`
	// The settings for source rate limiting.
	SourceLimit *ListPolicyResponseBodyPolicyListContentSourceLimit `json:"SourceLimit,omitempty" xml:"SourceLimit,omitempty" type:"Struct"`
	// Indicates whether back-to-origin CIDR blocks of Anti-DDoS Proxy are added to the whitelist.
	//
	// example:
	//
	// false
	WhitenGfbrNets *bool `json:"WhitenGfbrNets,omitempty" xml:"WhitenGfbrNets,omitempty"`
}

func (s ListPolicyResponseBodyPolicyListContent) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyListContent) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyListContent) SetBlackIpListExpireAt(v int64) *ListPolicyResponseBodyPolicyListContent {
	s.BlackIpListExpireAt = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetEnableDropIcmp(v bool) *ListPolicyResponseBodyPolicyListContent {
	s.EnableDropIcmp = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetEnableIntelligence(v bool) *ListPolicyResponseBodyPolicyListContent {
	s.EnableIntelligence = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetEnableL4Defense(v bool) *ListPolicyResponseBodyPolicyListContent {
	s.EnableL4Defense = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetFingerPrintRuleList(v []*ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) *ListPolicyResponseBodyPolicyListContent {
	s.FingerPrintRuleList = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetIntelligenceLevel(v string) *ListPolicyResponseBodyPolicyListContent {
	s.IntelligenceLevel = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetL4RuleList(v []*ListPolicyResponseBodyPolicyListContentL4RuleList) *ListPolicyResponseBodyPolicyListContent {
	s.L4RuleList = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetPortRuleList(v []*ListPolicyResponseBodyPolicyListContentPortRuleList) *ListPolicyResponseBodyPolicyListContent {
	s.PortRuleList = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetReflectBlockUdpPortList(v []*int32) *ListPolicyResponseBodyPolicyListContent {
	s.ReflectBlockUdpPortList = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetRegionBlockCountryList(v []*int32) *ListPolicyResponseBodyPolicyListContent {
	s.RegionBlockCountryList = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetRegionBlockProvinceList(v []*int32) *ListPolicyResponseBodyPolicyListContent {
	s.RegionBlockProvinceList = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetSourceBlockList(v []*ListPolicyResponseBodyPolicyListContentSourceBlockList) *ListPolicyResponseBodyPolicyListContent {
	s.SourceBlockList = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetSourceLimit(v *ListPolicyResponseBodyPolicyListContentSourceLimit) *ListPolicyResponseBodyPolicyListContent {
	s.SourceLimit = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetWhitenGfbrNets(v bool) *ListPolicyResponseBodyPolicyListContent {
	s.WhitenGfbrNets = &v
	return s
}

type ListPolicyResponseBodyPolicyListContentFingerPrintRuleList struct {
	// The end of the destination port range. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 65535
	DstPortEnd *int32 `json:"DstPortEnd,omitempty" xml:"DstPortEnd,omitempty"`
	// The start of the destination port range. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 0
	DstPortStart *int32 `json:"DstPortStart,omitempty" xml:"DstPortStart,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 2c0b09cd-a565-4481-9acb-418b********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The action triggered if the rule is matched. Valid values:
	//
	// 	- **accept**: allows the traffic that matches the conditions in the byte-match filter rule.
	//
	// 	- **drop**: discards the traffic that matches the conditions in the byte-match filter rule.
	//
	// 	- **ip_rate**: limits rates on the source IP address whose traffic matches the conditions in the byte-match filter rule. The rate limit is specified by **RateValue**.
	//
	// 	- **session_rate**: limits the number of sessions from the source IP address whose traffic matches the conditions in the byte-match filter rule. The rate limit is specified by **RateValue**.
	//
	// example:
	//
	// drop
	MatchAction *string `json:"MatchAction,omitempty" xml:"MatchAction,omitempty"`
	// The maximum packet length. Valid values: **1*	- to **1500**.
	//
	// example:
	//
	// 1500
	MaxPktLen *int32 `json:"MaxPktLen,omitempty" xml:"MaxPktLen,omitempty"`
	// The minimum packet length. Valid values: **1*	- to **1500**.
	//
	// example:
	//
	// 1
	MinPktLen *int32 `json:"MinPktLen,omitempty" xml:"MinPktLen,omitempty"`
	// The offset. Valid values: **0*	- to **1500**.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The payload. The value is a hexadecimal string.
	//
	// example:
	//
	// abcd
	PayloadBytes *string `json:"PayloadBytes,omitempty" xml:"PayloadBytes,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// example:
	//
	// udp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The rate limit. Valid values: **1*	- to **100000**.
	//
	// >  This parameter is required when **MatchAction*	- is set to **ip_rate*	- or **session_rate**.
	//
	// example:
	//
	// 1000
	RateValue *int32 `json:"RateValue,omitempty" xml:"RateValue,omitempty"`
	// The sequence number that indicates the order for the rule to take effect. The value is an integer.
	//
	// example:
	//
	// 1
	SeqNo *int32 `json:"SeqNo,omitempty" xml:"SeqNo,omitempty"`
	// The end of the source port range. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 65535
	SrcPortEnd *int32 `json:"SrcPortEnd,omitempty" xml:"SrcPortEnd,omitempty"`
	// The start of the source port range. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 0
	SrcPortStart *int32 `json:"SrcPortStart,omitempty" xml:"SrcPortStart,omitempty"`
}

func (s ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetDstPortEnd(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.DstPortEnd = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetDstPortStart(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.DstPortStart = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetId(v string) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.Id = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetMatchAction(v string) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.MatchAction = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetMaxPktLen(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.MaxPktLen = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetMinPktLen(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.MinPktLen = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetOffset(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.Offset = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetPayloadBytes(v string) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.PayloadBytes = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetProtocol(v string) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.Protocol = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetRateValue(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.RateValue = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetSeqNo(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.SeqNo = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetSrcPortEnd(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.SrcPortEnd = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetSrcPortStart(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.SrcPortStart = &v
	return s
}

type ListPolicyResponseBodyPolicyListContentL4RuleList struct {
	// The action that is specified in the rule. Valid value:
	//
	// 	- **2**: The traffic is discarded.
	//
	// example:
	//
	// 2
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The match conditions.
	ConditionList []*ListPolicyResponseBodyPolicyListContentL4RuleListConditionList `json:"ConditionList,omitempty" xml:"ConditionList,omitempty" type:"Repeated"`
	// The minimum number of bytes in a session to trigger matching. Valid values: **0*	- to **2048**.
	//
	// example:
	//
	// 0
	Limited *int32 `json:"Limited,omitempty" xml:"Limited,omitempty"`
	// The condition based on which an action is performed. Valid values:
	//
	// 	- **0**: If the rule is matched, the action specified in the rule is performed.
	//
	// 	- **1**: If the rule is not matched, the action specified in the rule is performed.
	//
	// example:
	//
	// 1
	Match *string `json:"Match,omitempty" xml:"Match,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **char**: string match.
	//
	// 	- **hex**: hexadecimal string match.
	//
	// example:
	//
	// char
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test**
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority of the rule.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s ListPolicyResponseBodyPolicyListContentL4RuleList) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyListContentL4RuleList) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) SetAction(v string) *ListPolicyResponseBodyPolicyListContentL4RuleList {
	s.Action = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) SetConditionList(v []*ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) *ListPolicyResponseBodyPolicyListContentL4RuleList {
	s.ConditionList = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) SetLimited(v int32) *ListPolicyResponseBodyPolicyListContentL4RuleList {
	s.Limited = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) SetMatch(v string) *ListPolicyResponseBodyPolicyListContentL4RuleList {
	s.Match = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) SetMethod(v string) *ListPolicyResponseBodyPolicyListContentL4RuleList {
	s.Method = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) SetName(v string) *ListPolicyResponseBodyPolicyListContentL4RuleList {
	s.Name = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) SetPriority(v int32) *ListPolicyResponseBodyPolicyListContentL4RuleList {
	s.Priority = &v
	return s
}

type ListPolicyResponseBodyPolicyListContentL4RuleListConditionList struct {
	// The term that is used for matching.
	//
	// >  If Method is set to **char**, the value of this parameter must be ASCII strings. If Method is set to **hex**, the value of this parameter must be hexadecimal strings. Maximum length: 2,048.
	//
	// example:
	//
	// test
	Arg *string `json:"Arg,omitempty" xml:"Arg,omitempty"`
	// The number of bytes from the start position for matching. Valid values: **1*	- to **2048**.
	//
	// example:
	//
	// 32
	Depth *int32 `json:"Depth,omitempty" xml:"Depth,omitempty"`
	// The start position for matching. Valid values: **0*	- to **2047**.
	//
	// example:
	//
	// 0
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) SetArg(v string) *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList {
	s.Arg = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) SetDepth(v int32) *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList {
	s.Depth = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) SetPosition(v int32) *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList {
	s.Position = &v
	return s
}

type ListPolicyResponseBodyPolicyListContentPortRuleList struct {
	// The end of the destination port range. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 65535
	DstPortEnd *int32 `json:"DstPortEnd,omitempty" xml:"DstPortEnd,omitempty"`
	// The start of the destination port range. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 0
	DstPortStart *int32 `json:"DstPortStart,omitempty" xml:"DstPortStart,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 8f3c3062-6c20-425d-8405-2bd1********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The action triggered if the rule is matched. Valid value:
	//
	// 	- **drop**: The traffic is discarded.
	//
	// example:
	//
	// drop
	MatchAction *string `json:"MatchAction,omitempty" xml:"MatchAction,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// example:
	//
	// udp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The sequence number that indicates the order for the rule to take effect. The value is an integer.
	//
	// example:
	//
	// 1
	SeqNo *int32 `json:"SeqNo,omitempty" xml:"SeqNo,omitempty"`
	// The end of the source port range. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 65535
	SrcPortEnd *int32 `json:"SrcPortEnd,omitempty" xml:"SrcPortEnd,omitempty"`
	// The start of the source port range. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 0
	SrcPortStart *int32 `json:"SrcPortStart,omitempty" xml:"SrcPortStart,omitempty"`
}

func (s ListPolicyResponseBodyPolicyListContentPortRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyListContentPortRuleList) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) SetDstPortEnd(v int32) *ListPolicyResponseBodyPolicyListContentPortRuleList {
	s.DstPortEnd = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) SetDstPortStart(v int32) *ListPolicyResponseBodyPolicyListContentPortRuleList {
	s.DstPortStart = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) SetId(v string) *ListPolicyResponseBodyPolicyListContentPortRuleList {
	s.Id = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) SetMatchAction(v string) *ListPolicyResponseBodyPolicyListContentPortRuleList {
	s.MatchAction = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) SetProtocol(v string) *ListPolicyResponseBodyPolicyListContentPortRuleList {
	s.Protocol = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) SetSeqNo(v int32) *ListPolicyResponseBodyPolicyListContentPortRuleList {
	s.SeqNo = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) SetSrcPortEnd(v int32) *ListPolicyResponseBodyPolicyListContentPortRuleList {
	s.SrcPortEnd = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) SetSrcPortStart(v int32) *ListPolicyResponseBodyPolicyListContentPortRuleList {
	s.SrcPortStart = &v
	return s
}

type ListPolicyResponseBodyPolicyListContentSourceBlockList struct {
	// The validity period of the blacklist to which the source IP address is added. Unit: seconds.
	//
	// example:
	//
	// 120
	BlockExpireSeconds *int32 `json:"BlockExpireSeconds,omitempty" xml:"BlockExpireSeconds,omitempty"`
	// The statistical period during which the system collects data on source IP addresses to determine whether to add the source IP addresses to the blacklist. Unit: seconds.
	//
	// example:
	//
	// 60
	EverySeconds *int32 `json:"EverySeconds,omitempty" xml:"EverySeconds,omitempty"`
	// The number of times that the source IP address exceeds a limit in a statistical period.
	//
	// example:
	//
	// 5
	ExceedLimitTimes *int32 `json:"ExceedLimitTimes,omitempty" xml:"ExceedLimitTimes,omitempty"`
	// The type of the source rate limit. Valid values:
	//
	// 	- **3**: the PPS limit on source IP addresses.
	//
	// 	- **4**: the bandwidth limit on source IP addresses.
	//
	// 	- **5**: the PPS limit on source SYN packets.
	//
	// 	- **6**: the bandwidth limit on source SYN packets.
	//
	// example:
	//
	// 3
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPolicyResponseBodyPolicyListContentSourceBlockList) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyListContentSourceBlockList) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyListContentSourceBlockList) SetBlockExpireSeconds(v int32) *ListPolicyResponseBodyPolicyListContentSourceBlockList {
	s.BlockExpireSeconds = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentSourceBlockList) SetEverySeconds(v int32) *ListPolicyResponseBodyPolicyListContentSourceBlockList {
	s.EverySeconds = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentSourceBlockList) SetExceedLimitTimes(v int32) *ListPolicyResponseBodyPolicyListContentSourceBlockList {
	s.ExceedLimitTimes = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentSourceBlockList) SetType(v int32) *ListPolicyResponseBodyPolicyListContentSourceBlockList {
	s.Type = &v
	return s
}

type ListPolicyResponseBodyPolicyListContentSourceLimit struct {
	// The bandwidth limit on source IP addresses. Unit: bytes per second.
	//
	// example:
	//
	// 2048
	Bps *int32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The packets per second (PPS) limit on source IP addresses.
	//
	// example:
	//
	// 64
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The bandwidth limit on source SYN packets. Unit: bytes per second.
	//
	// example:
	//
	// 2048
	SynBps *int32 `json:"SynBps,omitempty" xml:"SynBps,omitempty"`
	// The PPS limit on source SYN packets.
	//
	// example:
	//
	// 64
	SynPps *int32 `json:"SynPps,omitempty" xml:"SynPps,omitempty"`
}

func (s ListPolicyResponseBodyPolicyListContentSourceLimit) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyListContentSourceLimit) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyListContentSourceLimit) SetBps(v int32) *ListPolicyResponseBodyPolicyListContentSourceLimit {
	s.Bps = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentSourceLimit) SetPps(v int32) *ListPolicyResponseBodyPolicyListContentSourceLimit {
	s.Pps = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentSourceLimit) SetSynBps(v int32) *ListPolicyResponseBodyPolicyListContentSourceLimit {
	s.SynBps = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentSourceLimit) SetSynPps(v int32) *ListPolicyResponseBodyPolicyListContentSourceLimit {
	s.SynPps = &v
	return s
}

type ListPolicyResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListPolicyResponse) SetHeaders(v map[string]*string) *ListPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListPolicyResponse) SetStatusCode(v int32) *ListPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicyResponse) SetBody(v *ListPolicyResponseBody) *ListPolicyResponse {
	s.Body = v
	return s
}

type ListPolicyAttachmentRequest struct {
	// The protected objects.
	IpPortProtocolList []*ListPolicyAttachmentRequestIpPortProtocolList `json:"IpPortProtocolList,omitempty" xml:"IpPortProtocolList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// f38f6520-92b7-451e-b520-9ab3********
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **default**: the default mitigation policies.
	//
	// 	- **l3**: IP-specific mitigation policies.
	//
	// 	- **l4**: port-specific mitigation policies.
	//
	// example:
	//
	// l3
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPolicyAttachmentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyAttachmentRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentRequest) SetIpPortProtocolList(v []*ListPolicyAttachmentRequestIpPortProtocolList) *ListPolicyAttachmentRequest {
	s.IpPortProtocolList = v
	return s
}

func (s *ListPolicyAttachmentRequest) SetPageNo(v int64) *ListPolicyAttachmentRequest {
	s.PageNo = &v
	return s
}

func (s *ListPolicyAttachmentRequest) SetPageSize(v int64) *ListPolicyAttachmentRequest {
	s.PageSize = &v
	return s
}

func (s *ListPolicyAttachmentRequest) SetPolicyId(v string) *ListPolicyAttachmentRequest {
	s.PolicyId = &v
	return s
}

func (s *ListPolicyAttachmentRequest) SetPolicyType(v string) *ListPolicyAttachmentRequest {
	s.PolicyType = &v
	return s
}

type ListPolicyAttachmentRequestIpPortProtocolList struct {
	// The IP address of the protected object.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47.118.172.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port number of the protected object.
	//
	// example:
	//
	// 8*
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type of the protected object. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s ListPolicyAttachmentRequestIpPortProtocolList) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyAttachmentRequestIpPortProtocolList) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentRequestIpPortProtocolList) SetIp(v string) *ListPolicyAttachmentRequestIpPortProtocolList {
	s.Ip = &v
	return s
}

func (s *ListPolicyAttachmentRequestIpPortProtocolList) SetPort(v int32) *ListPolicyAttachmentRequestIpPortProtocolList {
	s.Port = &v
	return s
}

func (s *ListPolicyAttachmentRequestIpPortProtocolList) SetProtocol(v string) *ListPolicyAttachmentRequestIpPortProtocolList {
	s.Protocol = &v
	return s
}

type ListPolicyAttachmentShrinkRequest struct {
	// The protected objects.
	IpPortProtocolListShrink *string `json:"IpPortProtocolList,omitempty" xml:"IpPortProtocolList,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// f38f6520-92b7-451e-b520-9ab3********
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **default**: the default mitigation policies.
	//
	// 	- **l3**: IP-specific mitigation policies.
	//
	// 	- **l4**: port-specific mitigation policies.
	//
	// example:
	//
	// l3
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPolicyAttachmentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyAttachmentShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentShrinkRequest) SetIpPortProtocolListShrink(v string) *ListPolicyAttachmentShrinkRequest {
	s.IpPortProtocolListShrink = &v
	return s
}

func (s *ListPolicyAttachmentShrinkRequest) SetPageNo(v int64) *ListPolicyAttachmentShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *ListPolicyAttachmentShrinkRequest) SetPageSize(v int64) *ListPolicyAttachmentShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListPolicyAttachmentShrinkRequest) SetPolicyId(v string) *ListPolicyAttachmentShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *ListPolicyAttachmentShrinkRequest) SetPolicyType(v string) *ListPolicyAttachmentShrinkRequest {
	s.PolicyType = &v
	return s
}

type ListPolicyAttachmentResponseBody struct {
	// The records of attachments to the mitigation policy.
	AttachmentList []*ListPolicyAttachmentResponseBodyAttachmentList `json:"AttachmentList,omitempty" xml:"AttachmentList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B4B379C2-9319-4C6B-B579-FE36831B09F4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of attachments to the mitigation policy.
	//
	// example:
	//
	// 28
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListPolicyAttachmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentResponseBody) SetAttachmentList(v []*ListPolicyAttachmentResponseBodyAttachmentList) *ListPolicyAttachmentResponseBody {
	s.AttachmentList = v
	return s
}

func (s *ListPolicyAttachmentResponseBody) SetRequestId(v string) *ListPolicyAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicyAttachmentResponseBody) SetTotal(v int64) *ListPolicyAttachmentResponseBody {
	s.Total = &v
	return s
}

type ListPolicyAttachmentResponseBodyAttachmentList struct {
	// The IP address of the protected object.
	//
	// example:
	//
	// 147.139.183.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The UID of the member to which the IP address of the protected object belongs.
	//
	// example:
	//
	// 177699790631****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// 1b43f44e-65e1-411a-b0c0-d6c1********
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test**
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// test
	PolicyRemark *string `json:"PolicyRemark,omitempty" xml:"PolicyRemark,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **l3**: IP-specific mitigation policies.
	//
	// 	- **l4**: port-specific mitigation policies.
	//
	// example:
	//
	// l3
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The port number of the protected object.
	//
	// example:
	//
	// 8*
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type of the protected object. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// example:
	//
	// udp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region to which the IP address of the protected object belongs.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ListPolicyAttachmentResponseBodyAttachmentList) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyAttachmentResponseBodyAttachmentList) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetIp(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.Ip = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetMemberUid(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.MemberUid = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetPolicyId(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.PolicyId = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetPolicyName(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.PolicyName = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetPolicyRemark(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.PolicyRemark = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetPolicyType(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.PolicyType = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetPort(v int32) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.Port = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetProtocol(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.Protocol = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetRegion(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.Region = &v
	return s
}

type ListPolicyAttachmentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicyAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicyAttachmentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPolicyAttachmentResponse) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentResponse) SetHeaders(v map[string]*string) *ListPolicyAttachmentResponse {
	s.Headers = v
	return s
}

func (s *ListPolicyAttachmentResponse) SetStatusCode(v int32) *ListPolicyAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicyAttachmentResponse) SetBody(v *ListPolicyAttachmentResponseBody) *ListPolicyAttachmentResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Valid values: 1 to **50**. Default value: **10**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource type. Set the value to **INSTANCE**.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetCurrentPage(v int32) *ListTagKeysRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceGroupId(v string) *ListTagKeysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

type ListTagKeysResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 97935DF1-0289-4AA2-9DD1-72377838B16B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the tags.
	TagKeys []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	// The total number of tags returned.
	//
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetCurrentPage(v int32) *ListTagKeysResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysResponseBody) SetPageSize(v int32) *ListTagKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*ListTagKeysResponseBodyTagKeys) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

func (s *ListTagKeysResponseBody) SetTotalCount(v int32) *ListTagKeysResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagKeysResponseBodyTagKeys struct {
	// The total number of tag values that correspond to each key.
	//
	// example:
	//
	// 1
	TagCount *int32 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// The tag key.
	//
	// example:
	//
	// a
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagKeysResponseBodyTagKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagCount(v int32) *ListTagKeysResponseBodyTagKeys {
	s.TagCount = &v
	return s
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagKey(v string) *ListTagKeysResponseBodyTagKeys {
	s.TagKey = &v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// RGuYpqDdKhzXb8C3.D1BwQgc1tMBsoxdGiEKHHUUCf****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IDs of the Anti-DDoS Origin instances to query.
	//
	// >  The **ResourceId*	- parameter and the **key-value pair for the Tag parameter*	- cannot be left empty at the same time.
	//
	// example:
	//
	// ddosbgp-cn-v0h1fmwbc024
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource to query. Set the value to **INSTANCE**.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key-value pair of the tag to query.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceGroupId(v string) *ListTagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The key of the tag to query.
	//
	// >  The **ResourceId*	- parameter and the **key-value pair for the Tag parameter*	- cannot be left empty at the same time.
	//
	// example:
	//
	// testKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to query.
	//
	// >  The **ResourceId*	- parameter and the **key-value pair for the Tag parameter*	- cannot be left empty at the same time.
	//
	// example:
	//
	// testValue1
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
	// A pagination token.
	//
	// example:
	//
	// RGuYpqDdKhzXb8C3.D1BwQgc1tMBsoxdGiEKHHUUCf****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C3F7E6AE-43B2-4730-B6A3-FD17552B8F65
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags that are added to the Anti-DDoS Origin instance.
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
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

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	// The ID of the Anti-DDoS Origin instance.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. The value is set to **INSTANCE**.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag that is added to the instance.
	//
	// example:
	//
	// testKey1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag that is added to the instance.
	//
	// example:
	//
	// testValue1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyPolicyRequest struct {
	// The type of the action. Valid values:
	//
	// 	- **10**: modifies the name. If you specify this value, `Name` is required.
	//
	// 	- **11**: modifies the blacklist validity period. If you specify this value, `BlackIpListExpireAt` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **12**: changes the status of the feature of adding back-to-origin CIDR blocks of Anti-DDoS Proxy to the whitelist. If you specify this value, `WhitenGfbrNets` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **13**: changes the status of the ICMP blocking feature. If you specify this value, `EnableDropIcmp` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **20**: adds IP addresses to the blacklist or the whitelist. If you specify this value, you must specify at least one of `WhiteIpList` and `BlackIpList`. Only IP-specific mitigation policies support this value.
	//
	// 	- **21**: removes IP addresses from the blacklist or the whitelist. If you specify this value, at least one of `WhiteIpList` and `BlackIpList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **22**: clears the whitelist. Only IP-specific mitigation policies support this value.
	//
	// 	- **23**: clears the blacklist. Only IP-specific mitigation policies support this value.
	//
	// 	- **30**: modifies the status and level of intelligent protection. If you specify this value, `EnableIntelligence` and `IntelligenceLevel` are required. Only IP-specific mitigation policies support this value.
	//
	// 	- **31**: modifies the location blacklist settings. If you specify this value, one of `RegionBlockCountryList` and `RegionBlockProvinceList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **32**: modifies the settings for source rate limiting. If you specify this value, `SourceLimit` and `SourceBlockList` are required. Only IP-specific mitigation policies support this value.
	//
	// 	- **33**: modifies the settings for reflection attack filtering. If you specify this value, `ReflectBlockUdpPortList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **40**: creates a port blocking rule. If you specify this value, `PortRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **41**: modifies the port blocking rule. If you specify this value, `PortRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **42**: deletes the port blocking rule. If you specify this value, `PortRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **50**: creates a byte-match filter rule. If you specify this value, `FingerPrintRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **51**: modifies the byte-match filter rule. If you specify this value, `FingerPrintRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **52**: deletes the byte-match filter rule. If you specify this value, `FingerPrintRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **60**: changes the status of the port-specific mitigation feature. If you specify this value, `EnableL4Defense` is required. Only port-specific mitigation policies support this value.
	//
	// 	- **61**: creates a port-specific mitigation rule. If you specify this value, `L4RuleList` is required. Only port-specific mitigation policies support this value.
	//
	// 	- **62**: modifies the port-specific mitigation rule. If you specify this value, `L4RuleList` is required. Only port-specific mitigation policies support this value.
	//
	// 	- **63**: deletes the port-specific mitigation rule. If you specify this value, `L4RuleList` is required. Only port-specific mitigation policies support this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	ActionType *int32 `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The policy content.
	Content *ModifyPolicyRequestContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// c52c2fa6-fdac-40c4-8753-be7c********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// demo**
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequest) SetActionType(v int32) *ModifyPolicyRequest {
	s.ActionType = &v
	return s
}

func (s *ModifyPolicyRequest) SetContent(v *ModifyPolicyRequestContent) *ModifyPolicyRequest {
	s.Content = v
	return s
}

func (s *ModifyPolicyRequest) SetId(v string) *ModifyPolicyRequest {
	s.Id = &v
	return s
}

func (s *ModifyPolicyRequest) SetName(v string) *ModifyPolicyRequest {
	s.Name = &v
	return s
}

type ModifyPolicyRequestContent struct {
	// The IP addresses in the blacklist.
	BlackIpList []*string `json:"BlackIpList,omitempty" xml:"BlackIpList,omitempty" type:"Repeated"`
	// The validity period of the IP address blacklist. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1716878000
	BlackIpListExpireAt *int64 `json:"BlackIpListExpireAt,omitempty" xml:"BlackIpListExpireAt,omitempty"`
	// Specifies whether to enable ICMP blocking.
	//
	// example:
	//
	// true
	EnableDropIcmp *bool `json:"EnableDropIcmp,omitempty" xml:"EnableDropIcmp,omitempty"`
	// Specifies whether to enable intelligent protection.
	//
	// example:
	//
	// true
	EnableIntelligence *bool `json:"EnableIntelligence,omitempty" xml:"EnableIntelligence,omitempty"`
	// Specifies whether to enable port-specific mitigation.
	EnableL4Defense *bool `json:"EnableL4Defense,omitempty" xml:"EnableL4Defense,omitempty"`
	// The byte-match filter rules.
	FingerPrintRuleList []*ModifyPolicyRequestContentFingerPrintRuleList `json:"FingerPrintRuleList,omitempty" xml:"FingerPrintRuleList,omitempty" type:"Repeated"`
	// The level of intelligent protection. Valid values:
	//
	// 	- **default**: normal.
	//
	// 	- **hard**: strict.
	//
	// 	- **weak**: loose.
	//
	// example:
	//
	// default
	IntelligenceLevel *string `json:"IntelligenceLevel,omitempty" xml:"IntelligenceLevel,omitempty"`
	// The port-specific mitigation rules.
	L4RuleList []*ModifyPolicyRequestContentL4RuleList `json:"L4RuleList,omitempty" xml:"L4RuleList,omitempty" type:"Repeated"`
	// The port blocking rules.
	PortRuleList []*ModifyPolicyRequestContentPortRuleList `json:"PortRuleList,omitempty" xml:"PortRuleList,omitempty" type:"Repeated"`
	// The ports whose traffic is filtered out by the filtering policies for UDP reflection attacks.
	ReflectBlockUdpPortList []*int32 `json:"ReflectBlockUdpPortList,omitempty" xml:"ReflectBlockUdpPortList,omitempty" type:"Repeated"`
	// The countries in the location blacklist.
	RegionBlockCountryList []*int32 `json:"RegionBlockCountryList,omitempty" xml:"RegionBlockCountryList,omitempty" type:"Repeated"`
	// The provinces in the location blacklist.
	RegionBlockProvinceList []*int32 `json:"RegionBlockProvinceList,omitempty" xml:"RegionBlockProvinceList,omitempty" type:"Repeated"`
	// The source IP addresses that are added to the blacklist.
	SourceBlockList []*ModifyPolicyRequestContentSourceBlockList `json:"SourceBlockList,omitempty" xml:"SourceBlockList,omitempty" type:"Repeated"`
	// The settings for source rate limiting.
	SourceLimit *ModifyPolicyRequestContentSourceLimit `json:"SourceLimit,omitempty" xml:"SourceLimit,omitempty" type:"Struct"`
	// The IP addresses in the whitelist.
	WhiteIpList []*string `json:"WhiteIpList,omitempty" xml:"WhiteIpList,omitempty" type:"Repeated"`
	// Specifies whether to add back-to-origin CIDR blocks of Anti-DDoS Proxy to the whitelist.
	//
	// example:
	//
	// false
	WhitenGfbrNets *bool `json:"WhitenGfbrNets,omitempty" xml:"WhitenGfbrNets,omitempty"`
}

func (s ModifyPolicyRequestContent) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyRequestContent) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequestContent) SetBlackIpList(v []*string) *ModifyPolicyRequestContent {
	s.BlackIpList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetBlackIpListExpireAt(v int64) *ModifyPolicyRequestContent {
	s.BlackIpListExpireAt = &v
	return s
}

func (s *ModifyPolicyRequestContent) SetEnableDropIcmp(v bool) *ModifyPolicyRequestContent {
	s.EnableDropIcmp = &v
	return s
}

func (s *ModifyPolicyRequestContent) SetEnableIntelligence(v bool) *ModifyPolicyRequestContent {
	s.EnableIntelligence = &v
	return s
}

func (s *ModifyPolicyRequestContent) SetEnableL4Defense(v bool) *ModifyPolicyRequestContent {
	s.EnableL4Defense = &v
	return s
}

func (s *ModifyPolicyRequestContent) SetFingerPrintRuleList(v []*ModifyPolicyRequestContentFingerPrintRuleList) *ModifyPolicyRequestContent {
	s.FingerPrintRuleList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetIntelligenceLevel(v string) *ModifyPolicyRequestContent {
	s.IntelligenceLevel = &v
	return s
}

func (s *ModifyPolicyRequestContent) SetL4RuleList(v []*ModifyPolicyRequestContentL4RuleList) *ModifyPolicyRequestContent {
	s.L4RuleList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetPortRuleList(v []*ModifyPolicyRequestContentPortRuleList) *ModifyPolicyRequestContent {
	s.PortRuleList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetReflectBlockUdpPortList(v []*int32) *ModifyPolicyRequestContent {
	s.ReflectBlockUdpPortList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetRegionBlockCountryList(v []*int32) *ModifyPolicyRequestContent {
	s.RegionBlockCountryList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetRegionBlockProvinceList(v []*int32) *ModifyPolicyRequestContent {
	s.RegionBlockProvinceList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetSourceBlockList(v []*ModifyPolicyRequestContentSourceBlockList) *ModifyPolicyRequestContent {
	s.SourceBlockList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetSourceLimit(v *ModifyPolicyRequestContentSourceLimit) *ModifyPolicyRequestContent {
	s.SourceLimit = v
	return s
}

func (s *ModifyPolicyRequestContent) SetWhiteIpList(v []*string) *ModifyPolicyRequestContent {
	s.WhiteIpList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetWhitenGfbrNets(v bool) *ModifyPolicyRequestContent {
	s.WhitenGfbrNets = &v
	return s
}

type ModifyPolicyRequestContentFingerPrintRuleList struct {
	// The end of the destination port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65535
	DstPortEnd *int32 `json:"DstPortEnd,omitempty" xml:"DstPortEnd,omitempty"`
	// The start of the destination port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	DstPortStart *int32 `json:"DstPortStart,omitempty" xml:"DstPortStart,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 5fbe941f-a0cf-4a49-9c7c-8fac********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The action triggered if the rule is matched. Valid values:
	//
	// 	- **accept**: allows the traffic that matches the conditions in the byte-match filter rule.
	//
	// 	- **drop**: discards the traffic that matches the conditions in the byte-match filter rule.
	//
	// 	- **ip_rate**: limits rates on the source IP address whose traffic matches the conditions in the byte-match filter rule. The rate limit is specified by **RateValue**.
	//
	// 	- **session_rate**: limits the number of sessions from the source IP address whose traffic matches the conditions in the byte-match filter rule. The rate limit is specified by **RateValue**.
	//
	// This parameter is required.
	//
	// example:
	//
	// drop
	MatchAction *string `json:"MatchAction,omitempty" xml:"MatchAction,omitempty"`
	// The maximum packet length. Valid values: **1*	- to **1500**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1500
	MaxPktLen *int32 `json:"MaxPktLen,omitempty" xml:"MaxPktLen,omitempty"`
	// The minimum packet length. Valid values: **1*	- to **1500**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MinPktLen *int32 `json:"MinPktLen,omitempty" xml:"MinPktLen,omitempty"`
	// The offset. Valid values: **0*	- to **1500**.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The payload. The value is a hexadecimal string.
	//
	// example:
	//
	// abcd
	PayloadBytes *string `json:"PayloadBytes,omitempty" xml:"PayloadBytes,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// This parameter is required.
	//
	// example:
	//
	// udp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The rate limit. Valid values: **1*	- to **100000**.
	//
	// >  This parameter is required when **MatchAction*	- is set to **ip_rate*	- or **session_rate**.
	//
	// example:
	//
	// 100
	RateValue *int32 `json:"RateValue,omitempty" xml:"RateValue,omitempty"`
	// The sequence number that indicates the order for the rule to take effect. The value is an integer.
	//
	// >  A smaller number indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SeqNo *int32 `json:"SeqNo,omitempty" xml:"SeqNo,omitempty"`
	// The end of the source port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65535
	SrcPortEnd *int32 `json:"SrcPortEnd,omitempty" xml:"SrcPortEnd,omitempty"`
	// The start of the source port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	SrcPortStart *int32 `json:"SrcPortStart,omitempty" xml:"SrcPortStart,omitempty"`
}

func (s ModifyPolicyRequestContentFingerPrintRuleList) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyRequestContentFingerPrintRuleList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetDstPortEnd(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.DstPortEnd = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetDstPortStart(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.DstPortStart = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetId(v string) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.Id = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetMatchAction(v string) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.MatchAction = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetMaxPktLen(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.MaxPktLen = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetMinPktLen(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.MinPktLen = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetOffset(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.Offset = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetPayloadBytes(v string) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.PayloadBytes = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetProtocol(v string) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.Protocol = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetRateValue(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.RateValue = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetSeqNo(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.SeqNo = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetSrcPortEnd(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.SrcPortEnd = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetSrcPortStart(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.SrcPortStart = &v
	return s
}

type ModifyPolicyRequestContentL4RuleList struct {
	// The action that is specified in the rule. Valid value:
	//
	// 	- **2**: The traffic is discarded.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The match conditions.
	//
	// This parameter is required.
	ConditionList []*ModifyPolicyRequestContentL4RuleListConditionList `json:"ConditionList,omitempty" xml:"ConditionList,omitempty" type:"Repeated"`
	// The minimum number of bytes in a session to trigger matching. Valid values: **0*	- to **2048**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Limited *int32 `json:"Limited,omitempty" xml:"Limited,omitempty"`
	// The condition based on which an action is performed. Valid values:
	//
	// 	- **0**: If the rule is matched, the action specified in the rule is performed.
	//
	// 	- **1**: If the rule is not matched, the action specified in the rule is performed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Match *string `json:"Match,omitempty" xml:"Match,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **char**: string match.
	//
	// 	- **hex**: hexadecimal string match.
	//
	// This parameter is required.
	//
	// example:
	//
	// char
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// test****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority of the rule. Valid values: **1*	- to **100**.
	//
	// >  A smaller value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s ModifyPolicyRequestContentL4RuleList) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyRequestContentL4RuleList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequestContentL4RuleList) SetAction(v string) *ModifyPolicyRequestContentL4RuleList {
	s.Action = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleList) SetConditionList(v []*ModifyPolicyRequestContentL4RuleListConditionList) *ModifyPolicyRequestContentL4RuleList {
	s.ConditionList = v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleList) SetLimited(v int32) *ModifyPolicyRequestContentL4RuleList {
	s.Limited = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleList) SetMatch(v string) *ModifyPolicyRequestContentL4RuleList {
	s.Match = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleList) SetMethod(v string) *ModifyPolicyRequestContentL4RuleList {
	s.Method = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleList) SetName(v string) *ModifyPolicyRequestContentL4RuleList {
	s.Name = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleList) SetPriority(v int32) *ModifyPolicyRequestContentL4RuleList {
	s.Priority = &v
	return s
}

type ModifyPolicyRequestContentL4RuleListConditionList struct {
	// The term that is used for matching.
	//
	// >  If Method is set to **char**, the value of this parameter must be ASCII strings. If Method is set to **hex**, the value of this parameter must be hexadecimal strings. Maximum length: 2,048.
	//
	// This parameter is required.
	//
	// example:
	//
	// abcd
	Arg *string `json:"Arg,omitempty" xml:"Arg,omitempty"`
	// The number of bytes from the start position for matching. Valid values: **1*	- to **2048**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1200
	Depth *int32 `json:"Depth,omitempty" xml:"Depth,omitempty"`
	// The start position for matching. Valid values: **0*	- to **2047**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s ModifyPolicyRequestContentL4RuleListConditionList) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyRequestContentL4RuleListConditionList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) SetArg(v string) *ModifyPolicyRequestContentL4RuleListConditionList {
	s.Arg = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) SetDepth(v int32) *ModifyPolicyRequestContentL4RuleListConditionList {
	s.Depth = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) SetPosition(v int32) *ModifyPolicyRequestContentL4RuleListConditionList {
	s.Position = &v
	return s
}

type ModifyPolicyRequestContentPortRuleList struct {
	// The end of the destination port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65535
	DstPortEnd *int32 `json:"DstPortEnd,omitempty" xml:"DstPortEnd,omitempty"`
	// The start of the destination port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	DstPortStart *int32 `json:"DstPortStart,omitempty" xml:"DstPortStart,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// c52c2fa6-fdac-40c4-8753-be7c*********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The action triggered if the rule is matched. Valid values:
	//
	// 	- **drop**: The traffic is discarded.
	//
	// This parameter is required.
	//
	// example:
	//
	// drop
	MatchAction *string `json:"MatchAction,omitempty" xml:"MatchAction,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// This parameter is required.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The sequence number that indicates the order for the rule to take effect. The value is an integer.
	//
	// >  A smaller number indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SeqNo *int32 `json:"SeqNo,omitempty" xml:"SeqNo,omitempty"`
	// The end of the source port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65535
	SrcPortEnd *int32 `json:"SrcPortEnd,omitempty" xml:"SrcPortEnd,omitempty"`
	// The start of the source port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	SrcPortStart *int32 `json:"SrcPortStart,omitempty" xml:"SrcPortStart,omitempty"`
}

func (s ModifyPolicyRequestContentPortRuleList) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyRequestContentPortRuleList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequestContentPortRuleList) SetDstPortEnd(v int32) *ModifyPolicyRequestContentPortRuleList {
	s.DstPortEnd = &v
	return s
}

func (s *ModifyPolicyRequestContentPortRuleList) SetDstPortStart(v int32) *ModifyPolicyRequestContentPortRuleList {
	s.DstPortStart = &v
	return s
}

func (s *ModifyPolicyRequestContentPortRuleList) SetId(v string) *ModifyPolicyRequestContentPortRuleList {
	s.Id = &v
	return s
}

func (s *ModifyPolicyRequestContentPortRuleList) SetMatchAction(v string) *ModifyPolicyRequestContentPortRuleList {
	s.MatchAction = &v
	return s
}

func (s *ModifyPolicyRequestContentPortRuleList) SetProtocol(v string) *ModifyPolicyRequestContentPortRuleList {
	s.Protocol = &v
	return s
}

func (s *ModifyPolicyRequestContentPortRuleList) SetSeqNo(v int32) *ModifyPolicyRequestContentPortRuleList {
	s.SeqNo = &v
	return s
}

func (s *ModifyPolicyRequestContentPortRuleList) SetSrcPortEnd(v int32) *ModifyPolicyRequestContentPortRuleList {
	s.SrcPortEnd = &v
	return s
}

func (s *ModifyPolicyRequestContentPortRuleList) SetSrcPortStart(v int32) *ModifyPolicyRequestContentPortRuleList {
	s.SrcPortStart = &v
	return s
}

type ModifyPolicyRequestContentSourceBlockList struct {
	// The validity period of the blacklist to which the source IP address is added. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120
	BlockExpireSeconds *int32 `json:"BlockExpireSeconds,omitempty" xml:"BlockExpireSeconds,omitempty"`
	// The statistical period during which the system collects data on source IP addresses to determine whether to add the source IP addresses to the blacklist. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	EverySeconds *int32 `json:"EverySeconds,omitempty" xml:"EverySeconds,omitempty"`
	// The number of times that the source IP address exceeds a limit in a statistical period.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	ExceedLimitTimes *int32 `json:"ExceedLimitTimes,omitempty" xml:"ExceedLimitTimes,omitempty"`
	// The type of the source rate limit. Valid values:
	//
	// 	- **3**: the pps limit on source IP addresses.
	//
	// 	- **4**: the bandwidth limit on source IP addresses.
	//
	// 	- **5**: the pps limit on source SYN packets.
	//
	// 	- **6**: the bandwidth limit on source SYN packets.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyPolicyRequestContentSourceBlockList) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyRequestContentSourceBlockList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequestContentSourceBlockList) SetBlockExpireSeconds(v int32) *ModifyPolicyRequestContentSourceBlockList {
	s.BlockExpireSeconds = &v
	return s
}

func (s *ModifyPolicyRequestContentSourceBlockList) SetEverySeconds(v int32) *ModifyPolicyRequestContentSourceBlockList {
	s.EverySeconds = &v
	return s
}

func (s *ModifyPolicyRequestContentSourceBlockList) SetExceedLimitTimes(v int32) *ModifyPolicyRequestContentSourceBlockList {
	s.ExceedLimitTimes = &v
	return s
}

func (s *ModifyPolicyRequestContentSourceBlockList) SetType(v int32) *ModifyPolicyRequestContentSourceBlockList {
	s.Type = &v
	return s
}

type ModifyPolicyRequestContentSourceLimit struct {
	// The bandwidth limit on source IP addresses. Unit: bytes per second.
	//
	// example:
	//
	// 2048
	Bps *int32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The packets per second (pps) limit on source IP addresses.
	//
	// example:
	//
	// 64
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The bandwidth limit on source SYN packets. Unit: bytes per second.
	//
	// example:
	//
	// 2048
	SynBps *int32 `json:"SynBps,omitempty" xml:"SynBps,omitempty"`
	// The pps limit on source SYN packets.
	//
	// example:
	//
	// 64
	SynPps *int32 `json:"SynPps,omitempty" xml:"SynPps,omitempty"`
}

func (s ModifyPolicyRequestContentSourceLimit) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyRequestContentSourceLimit) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequestContentSourceLimit) SetBps(v int32) *ModifyPolicyRequestContentSourceLimit {
	s.Bps = &v
	return s
}

func (s *ModifyPolicyRequestContentSourceLimit) SetPps(v int32) *ModifyPolicyRequestContentSourceLimit {
	s.Pps = &v
	return s
}

func (s *ModifyPolicyRequestContentSourceLimit) SetSynBps(v int32) *ModifyPolicyRequestContentSourceLimit {
	s.SynBps = &v
	return s
}

func (s *ModifyPolicyRequestContentSourceLimit) SetSynPps(v int32) *ModifyPolicyRequestContentSourceLimit {
	s.SynPps = &v
	return s
}

type ModifyPolicyShrinkRequest struct {
	// The type of the action. Valid values:
	//
	// 	- **10**: modifies the name. If you specify this value, `Name` is required.
	//
	// 	- **11**: modifies the blacklist validity period. If you specify this value, `BlackIpListExpireAt` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **12**: changes the status of the feature of adding back-to-origin CIDR blocks of Anti-DDoS Proxy to the whitelist. If you specify this value, `WhitenGfbrNets` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **13**: changes the status of the ICMP blocking feature. If you specify this value, `EnableDropIcmp` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **20**: adds IP addresses to the blacklist or the whitelist. If you specify this value, you must specify at least one of `WhiteIpList` and `BlackIpList`. Only IP-specific mitigation policies support this value.
	//
	// 	- **21**: removes IP addresses from the blacklist or the whitelist. If you specify this value, at least one of `WhiteIpList` and `BlackIpList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **22**: clears the whitelist. Only IP-specific mitigation policies support this value.
	//
	// 	- **23**: clears the blacklist. Only IP-specific mitigation policies support this value.
	//
	// 	- **30**: modifies the status and level of intelligent protection. If you specify this value, `EnableIntelligence` and `IntelligenceLevel` are required. Only IP-specific mitigation policies support this value.
	//
	// 	- **31**: modifies the location blacklist settings. If you specify this value, one of `RegionBlockCountryList` and `RegionBlockProvinceList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **32**: modifies the settings for source rate limiting. If you specify this value, `SourceLimit` and `SourceBlockList` are required. Only IP-specific mitigation policies support this value.
	//
	// 	- **33**: modifies the settings for reflection attack filtering. If you specify this value, `ReflectBlockUdpPortList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **40**: creates a port blocking rule. If you specify this value, `PortRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **41**: modifies the port blocking rule. If you specify this value, `PortRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **42**: deletes the port blocking rule. If you specify this value, `PortRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **50**: creates a byte-match filter rule. If you specify this value, `FingerPrintRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **51**: modifies the byte-match filter rule. If you specify this value, `FingerPrintRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **52**: deletes the byte-match filter rule. If you specify this value, `FingerPrintRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **60**: changes the status of the port-specific mitigation feature. If you specify this value, `EnableL4Defense` is required. Only port-specific mitigation policies support this value.
	//
	// 	- **61**: creates a port-specific mitigation rule. If you specify this value, `L4RuleList` is required. Only port-specific mitigation policies support this value.
	//
	// 	- **62**: modifies the port-specific mitigation rule. If you specify this value, `L4RuleList` is required. Only port-specific mitigation policies support this value.
	//
	// 	- **63**: deletes the port-specific mitigation rule. If you specify this value, `L4RuleList` is required. Only port-specific mitigation policies support this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	ActionType *int32 `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The policy content.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// c52c2fa6-fdac-40c4-8753-be7c********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// demo**
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyPolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyShrinkRequest) SetActionType(v int32) *ModifyPolicyShrinkRequest {
	s.ActionType = &v
	return s
}

func (s *ModifyPolicyShrinkRequest) SetContentShrink(v string) *ModifyPolicyShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *ModifyPolicyShrinkRequest) SetId(v string) *ModifyPolicyShrinkRequest {
	s.Id = &v
	return s
}

func (s *ModifyPolicyShrinkRequest) SetName(v string) *ModifyPolicyShrinkRequest {
	s.Name = &v
	return s
}

type ModifyPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B4B379C2-9319-4C6B-B579-FE36831****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolicyResponseBody) SetRequestId(v string) *ModifyPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolicyResponse) SetHeaders(v map[string]*string) *ModifyPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolicyResponse) SetStatusCode(v int32) *ModifyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolicyResponse) SetBody(v *ModifyPolicyResponseBody) *ModifyPolicyResponse {
	s.Body = v
	return s
}

type ModifyPolicyContentRequest struct {
	// The policy content.
	Content *ModifyPolicyContentRequestContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 83967609-7ea5-4f6d-a6ea-380b09e****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// demo**
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyPolicyContentRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyContentRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequest) SetContent(v *ModifyPolicyContentRequestContent) *ModifyPolicyContentRequest {
	s.Content = v
	return s
}

func (s *ModifyPolicyContentRequest) SetId(v string) *ModifyPolicyContentRequest {
	s.Id = &v
	return s
}

func (s *ModifyPolicyContentRequest) SetName(v string) *ModifyPolicyContentRequest {
	s.Name = &v
	return s
}

type ModifyPolicyContentRequestContent struct {
	// The validity period of the IP address blacklist. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1716878000
	BlackIpListExpireAt *int64 `json:"BlackIpListExpireAt,omitempty" xml:"BlackIpListExpireAt,omitempty"`
	// Specifies whether to enable ICMP blocking.
	//
	// example:
	//
	// true
	EnableDropIcmp *bool `json:"EnableDropIcmp,omitempty" xml:"EnableDropIcmp,omitempty"`
	// Specifies whether to enable intelligent protection.
	//
	// example:
	//
	// true
	EnableIntelligence *bool `json:"EnableIntelligence,omitempty" xml:"EnableIntelligence,omitempty"`
	// Specifies whether to enable port-specific mitigation.
	EnableL4Defense *bool `json:"EnableL4Defense,omitempty" xml:"EnableL4Defense,omitempty"`
	// The byte-match filter rules.
	FingerPrintRuleList []*ModifyPolicyContentRequestContentFingerPrintRuleList `json:"FingerPrintRuleList,omitempty" xml:"FingerPrintRuleList,omitempty" type:"Repeated"`
	// The level of intelligent protection. Valid values:
	//
	// 	- **default**: normal.
	//
	// 	- **hard**: strict.
	//
	// 	- **weak**: loose.
	//
	// example:
	//
	// default
	IntelligenceLevel *string `json:"IntelligenceLevel,omitempty" xml:"IntelligenceLevel,omitempty"`
	// The port-specific mitigation rules.
	L4RuleList []*ModifyPolicyContentRequestContentL4RuleList `json:"L4RuleList,omitempty" xml:"L4RuleList,omitempty" type:"Repeated"`
	// The port blocking rules.
	PortRuleList []*ModifyPolicyContentRequestContentPortRuleList `json:"PortRuleList,omitempty" xml:"PortRuleList,omitempty" type:"Repeated"`
	// The ports whose traffic is filtered out by the filtering policies for UDP reflection attacks.
	ReflectBlockUdpPortList []*int32 `json:"ReflectBlockUdpPortList,omitempty" xml:"ReflectBlockUdpPortList,omitempty" type:"Repeated"`
	// The countries in the location blacklist.
	RegionBlockCountryList []*int32 `json:"RegionBlockCountryList,omitempty" xml:"RegionBlockCountryList,omitempty" type:"Repeated"`
	// The provinces in the location blacklist.
	RegionBlockProvinceList []*int32 `json:"RegionBlockProvinceList,omitempty" xml:"RegionBlockProvinceList,omitempty" type:"Repeated"`
	// The source IP addresses that are added to the blacklist.
	SourceBlockList []*ModifyPolicyContentRequestContentSourceBlockList `json:"SourceBlockList,omitempty" xml:"SourceBlockList,omitempty" type:"Repeated"`
	// The settings for source rate limiting.
	SourceLimit *ModifyPolicyContentRequestContentSourceLimit `json:"SourceLimit,omitempty" xml:"SourceLimit,omitempty" type:"Struct"`
	// Specifies whether to add back-to-origin CIDR blocks of Anti-DDoS Proxy to the whitelist.
	//
	// example:
	//
	// false
	WhitenGfbrNets *bool `json:"WhitenGfbrNets,omitempty" xml:"WhitenGfbrNets,omitempty"`
}

func (s ModifyPolicyContentRequestContent) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyContentRequestContent) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequestContent) SetBlackIpListExpireAt(v int64) *ModifyPolicyContentRequestContent {
	s.BlackIpListExpireAt = &v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetEnableDropIcmp(v bool) *ModifyPolicyContentRequestContent {
	s.EnableDropIcmp = &v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetEnableIntelligence(v bool) *ModifyPolicyContentRequestContent {
	s.EnableIntelligence = &v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetEnableL4Defense(v bool) *ModifyPolicyContentRequestContent {
	s.EnableL4Defense = &v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetFingerPrintRuleList(v []*ModifyPolicyContentRequestContentFingerPrintRuleList) *ModifyPolicyContentRequestContent {
	s.FingerPrintRuleList = v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetIntelligenceLevel(v string) *ModifyPolicyContentRequestContent {
	s.IntelligenceLevel = &v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetL4RuleList(v []*ModifyPolicyContentRequestContentL4RuleList) *ModifyPolicyContentRequestContent {
	s.L4RuleList = v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetPortRuleList(v []*ModifyPolicyContentRequestContentPortRuleList) *ModifyPolicyContentRequestContent {
	s.PortRuleList = v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetReflectBlockUdpPortList(v []*int32) *ModifyPolicyContentRequestContent {
	s.ReflectBlockUdpPortList = v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetRegionBlockCountryList(v []*int32) *ModifyPolicyContentRequestContent {
	s.RegionBlockCountryList = v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetRegionBlockProvinceList(v []*int32) *ModifyPolicyContentRequestContent {
	s.RegionBlockProvinceList = v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetSourceBlockList(v []*ModifyPolicyContentRequestContentSourceBlockList) *ModifyPolicyContentRequestContent {
	s.SourceBlockList = v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetSourceLimit(v *ModifyPolicyContentRequestContentSourceLimit) *ModifyPolicyContentRequestContent {
	s.SourceLimit = v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetWhitenGfbrNets(v bool) *ModifyPolicyContentRequestContent {
	s.WhitenGfbrNets = &v
	return s
}

type ModifyPolicyContentRequestContentFingerPrintRuleList struct {
	// The end of the destination port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65535
	DstPortEnd *int32 `json:"DstPortEnd,omitempty" xml:"DstPortEnd,omitempty"`
	// The start of the destination port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	DstPortStart *int32 `json:"DstPortStart,omitempty" xml:"DstPortStart,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 83967609-7ea5-4f6d-a6ea-380b09e****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The action triggered if the rule is matched. Valid values:
	//
	// 	- **permit**: allows the traffic that matches the conditions in the byte-match filter rule.
	//
	// 	- **drop**: discards the traffic that matches the conditions in the byte-match filter rule.
	//
	// 	- **ip_rate**: limits rates on the source IP address whose traffic matches the conditions in the byte-match filter rule. The rate limit is specified by **RateValue**.
	//
	// 	- **session_rate**: limits the number of sessions from the source IP address whose traffic matches the conditions in the byte-match filter rule. The rate limit is specified by **RateValue**.
	//
	// This parameter is required.
	//
	// example:
	//
	// drop
	MatchAction *string `json:"MatchAction,omitempty" xml:"MatchAction,omitempty"`
	// The maximum packet length. Valid values: **1*	- to **1500**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1500
	MaxPktLen *int32 `json:"MaxPktLen,omitempty" xml:"MaxPktLen,omitempty"`
	// The minimum packet length. Valid values: **1*	- to **1500**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MinPktLen *int32 `json:"MinPktLen,omitempty" xml:"MinPktLen,omitempty"`
	// The offset. Valid values: **0*	- to **1500**.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The payload. The value is a hexadecimal string.
	//
	// example:
	//
	// abcd
	PayloadBytes *string `json:"PayloadBytes,omitempty" xml:"PayloadBytes,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// This parameter is required.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The rate limit. Valid values: **1*	- to **100000**.
	//
	// >  This parameter is required when **MatchAction*	- is set to **ip_rate*	- or **session_rate**.
	//
	// example:
	//
	// 100
	RateValue *int32 `json:"RateValue,omitempty" xml:"RateValue,omitempty"`
	// The sequence number that indicates the order for the rule to take effect. The value is an integer.
	//
	// >  A smaller number indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SeqNo *int32 `json:"SeqNo,omitempty" xml:"SeqNo,omitempty"`
	// The end of the source port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65535
	SrcPortEnd *int32 `json:"SrcPortEnd,omitempty" xml:"SrcPortEnd,omitempty"`
	// The start of the source port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	SrcPortStart *int32 `json:"SrcPortStart,omitempty" xml:"SrcPortStart,omitempty"`
}

func (s ModifyPolicyContentRequestContentFingerPrintRuleList) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyContentRequestContentFingerPrintRuleList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetDstPortEnd(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.DstPortEnd = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetDstPortStart(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.DstPortStart = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetId(v string) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.Id = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetMatchAction(v string) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.MatchAction = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetMaxPktLen(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.MaxPktLen = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetMinPktLen(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.MinPktLen = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetOffset(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.Offset = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetPayloadBytes(v string) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.PayloadBytes = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetProtocol(v string) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.Protocol = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetRateValue(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.RateValue = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetSeqNo(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.SeqNo = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetSrcPortEnd(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.SrcPortEnd = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetSrcPortStart(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.SrcPortStart = &v
	return s
}

type ModifyPolicyContentRequestContentL4RuleList struct {
	// The action that is specified in the rule. Valid value:
	//
	// 	- **2**: The traffic is discarded.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The match conditions.
	//
	// This parameter is required.
	ConditionList []*ModifyPolicyContentRequestContentL4RuleListConditionList `json:"ConditionList,omitempty" xml:"ConditionList,omitempty" type:"Repeated"`
	// The minimum number of bytes in a session to trigger matching. Valid values: **0*	- to **2048**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Limited *int32 `json:"Limited,omitempty" xml:"Limited,omitempty"`
	// The condition based on which an action is performed. Valid values:
	//
	// 	- **0**: If the rule is matched, the action specified in the rule is performed.
	//
	// 	- **1**: If the rule is not matched, the action specified in the rule is performed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Match *string `json:"Match,omitempty" xml:"Match,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **char**: string match.
	//
	// 	- **hex**: hexadecimal string match.
	//
	// This parameter is required.
	//
	// example:
	//
	// char
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// test**
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority of the rule. Valid values: 1 to 100.
	//
	// >  A smaller value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s ModifyPolicyContentRequestContentL4RuleList) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyContentRequestContentL4RuleList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequestContentL4RuleList) SetAction(v string) *ModifyPolicyContentRequestContentL4RuleList {
	s.Action = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleList) SetConditionList(v []*ModifyPolicyContentRequestContentL4RuleListConditionList) *ModifyPolicyContentRequestContentL4RuleList {
	s.ConditionList = v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleList) SetLimited(v int32) *ModifyPolicyContentRequestContentL4RuleList {
	s.Limited = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleList) SetMatch(v string) *ModifyPolicyContentRequestContentL4RuleList {
	s.Match = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleList) SetMethod(v string) *ModifyPolicyContentRequestContentL4RuleList {
	s.Method = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleList) SetName(v string) *ModifyPolicyContentRequestContentL4RuleList {
	s.Name = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleList) SetPriority(v int32) *ModifyPolicyContentRequestContentL4RuleList {
	s.Priority = &v
	return s
}

type ModifyPolicyContentRequestContentL4RuleListConditionList struct {
	// The term that is used for matching.
	//
	// >  If Method is set to **char**, the value of this parameter must be ASCII strings. If Method is set to **hex**, the value of this parameter must be hexadecimal strings. Maximum length: 2,048.
	//
	// This parameter is required.
	//
	// example:
	//
	// abcd
	Arg *string `json:"Arg,omitempty" xml:"Arg,omitempty"`
	// The number of bytes from the start position for matching. Valid values: **1*	- to **2048**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1200
	Depth *int32 `json:"Depth,omitempty" xml:"Depth,omitempty"`
	// The start position for matching. Valid values: **0*	- to **2047**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s ModifyPolicyContentRequestContentL4RuleListConditionList) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyContentRequestContentL4RuleListConditionList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) SetArg(v string) *ModifyPolicyContentRequestContentL4RuleListConditionList {
	s.Arg = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) SetDepth(v int32) *ModifyPolicyContentRequestContentL4RuleListConditionList {
	s.Depth = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) SetPosition(v int32) *ModifyPolicyContentRequestContentL4RuleListConditionList {
	s.Position = &v
	return s
}

type ModifyPolicyContentRequestContentPortRuleList struct {
	// The end of the destination port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65535
	DstPortEnd *int32 `json:"DstPortEnd,omitempty" xml:"DstPortEnd,omitempty"`
	// The start of the destination port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	DstPortStart *int32 `json:"DstPortStart,omitempty" xml:"DstPortStart,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 412a7312-58ff-4e32-a202-0ab0*******
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The action triggered if the rule is matched. Valid values:
	//
	// 	- **drop**: The traffic is discarded.
	//
	// This parameter is required.
	//
	// example:
	//
	// drop
	MatchAction *string `json:"MatchAction,omitempty" xml:"MatchAction,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// This parameter is required.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The sequence number that indicates the order for the rule to take effect. The value is an integer.
	//
	// >  A smaller number indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SeqNo *int32 `json:"SeqNo,omitempty" xml:"SeqNo,omitempty"`
	// The end of the source port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65535
	SrcPortEnd *int32 `json:"SrcPortEnd,omitempty" xml:"SrcPortEnd,omitempty"`
	// The start of the source port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	SrcPortStart *int32 `json:"SrcPortStart,omitempty" xml:"SrcPortStart,omitempty"`
}

func (s ModifyPolicyContentRequestContentPortRuleList) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyContentRequestContentPortRuleList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequestContentPortRuleList) SetDstPortEnd(v int32) *ModifyPolicyContentRequestContentPortRuleList {
	s.DstPortEnd = &v
	return s
}

func (s *ModifyPolicyContentRequestContentPortRuleList) SetDstPortStart(v int32) *ModifyPolicyContentRequestContentPortRuleList {
	s.DstPortStart = &v
	return s
}

func (s *ModifyPolicyContentRequestContentPortRuleList) SetId(v string) *ModifyPolicyContentRequestContentPortRuleList {
	s.Id = &v
	return s
}

func (s *ModifyPolicyContentRequestContentPortRuleList) SetMatchAction(v string) *ModifyPolicyContentRequestContentPortRuleList {
	s.MatchAction = &v
	return s
}

func (s *ModifyPolicyContentRequestContentPortRuleList) SetProtocol(v string) *ModifyPolicyContentRequestContentPortRuleList {
	s.Protocol = &v
	return s
}

func (s *ModifyPolicyContentRequestContentPortRuleList) SetSeqNo(v int32) *ModifyPolicyContentRequestContentPortRuleList {
	s.SeqNo = &v
	return s
}

func (s *ModifyPolicyContentRequestContentPortRuleList) SetSrcPortEnd(v int32) *ModifyPolicyContentRequestContentPortRuleList {
	s.SrcPortEnd = &v
	return s
}

func (s *ModifyPolicyContentRequestContentPortRuleList) SetSrcPortStart(v int32) *ModifyPolicyContentRequestContentPortRuleList {
	s.SrcPortStart = &v
	return s
}

type ModifyPolicyContentRequestContentSourceBlockList struct {
	// The validity period of the blacklist to which the source IP address is added. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120
	BlockExpireSeconds *int32 `json:"BlockExpireSeconds,omitempty" xml:"BlockExpireSeconds,omitempty"`
	// The statistical period during which the system collects data on source IP addresses to determine whether to add the source IP addresses to the blacklist. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	EverySeconds *int32 `json:"EverySeconds,omitempty" xml:"EverySeconds,omitempty"`
	// The number of times that the source IP address exceeds a limit in a statistical period.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	ExceedLimitTimes *int32 `json:"ExceedLimitTimes,omitempty" xml:"ExceedLimitTimes,omitempty"`
	// The type of the source rate limit. Valid values:
	//
	// 	- **3**: the pps limit on source IP addresses.
	//
	// 	- **4**: the bandwidth limit on source IP addresses.
	//
	// 	- **5**: the pps limit on source SYN packets.
	//
	// 	- **6**: the bandwidth limit on source SYN packets.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyPolicyContentRequestContentSourceBlockList) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyContentRequestContentSourceBlockList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequestContentSourceBlockList) SetBlockExpireSeconds(v int32) *ModifyPolicyContentRequestContentSourceBlockList {
	s.BlockExpireSeconds = &v
	return s
}

func (s *ModifyPolicyContentRequestContentSourceBlockList) SetEverySeconds(v int32) *ModifyPolicyContentRequestContentSourceBlockList {
	s.EverySeconds = &v
	return s
}

func (s *ModifyPolicyContentRequestContentSourceBlockList) SetExceedLimitTimes(v int32) *ModifyPolicyContentRequestContentSourceBlockList {
	s.ExceedLimitTimes = &v
	return s
}

func (s *ModifyPolicyContentRequestContentSourceBlockList) SetType(v int32) *ModifyPolicyContentRequestContentSourceBlockList {
	s.Type = &v
	return s
}

type ModifyPolicyContentRequestContentSourceLimit struct {
	// The bandwidth limit on source IP addresses. Unit: bytes per second.
	//
	// example:
	//
	// 2048
	Bps *int32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The packets per second (pps) limit on source IP addresses.
	//
	// example:
	//
	// 64
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The bandwidth limit on source SYN packets. Unit: bytes per second.
	//
	// example:
	//
	// 2048
	SynBps *int32 `json:"SynBps,omitempty" xml:"SynBps,omitempty"`
	// The pps limit on source SYN packets.
	//
	// example:
	//
	// 64
	SynPps *int32 `json:"SynPps,omitempty" xml:"SynPps,omitempty"`
}

func (s ModifyPolicyContentRequestContentSourceLimit) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyContentRequestContentSourceLimit) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequestContentSourceLimit) SetBps(v int32) *ModifyPolicyContentRequestContentSourceLimit {
	s.Bps = &v
	return s
}

func (s *ModifyPolicyContentRequestContentSourceLimit) SetPps(v int32) *ModifyPolicyContentRequestContentSourceLimit {
	s.Pps = &v
	return s
}

func (s *ModifyPolicyContentRequestContentSourceLimit) SetSynBps(v int32) *ModifyPolicyContentRequestContentSourceLimit {
	s.SynBps = &v
	return s
}

func (s *ModifyPolicyContentRequestContentSourceLimit) SetSynPps(v int32) *ModifyPolicyContentRequestContentSourceLimit {
	s.SynPps = &v
	return s
}

type ModifyPolicyContentShrinkRequest struct {
	// The policy content.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 83967609-7ea5-4f6d-a6ea-380b09e****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// demo**
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyPolicyContentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentShrinkRequest) SetContentShrink(v string) *ModifyPolicyContentShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *ModifyPolicyContentShrinkRequest) SetId(v string) *ModifyPolicyContentShrinkRequest {
	s.Id = &v
	return s
}

func (s *ModifyPolicyContentShrinkRequest) SetName(v string) *ModifyPolicyContentShrinkRequest {
	s.Name = &v
	return s
}

type ModifyPolicyContentResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3777EF25-940B-51F4-BB1D-99B5********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPolicyContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyContentResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentResponseBody) SetRequestId(v string) *ModifyPolicyContentResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPolicyContentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPolicyContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPolicyContentResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyContentResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentResponse) SetHeaders(v map[string]*string) *ModifyPolicyContentResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolicyContentResponse) SetStatusCode(v int32) *ModifyPolicyContentResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolicyContentResponse) SetBody(v *ModifyPolicyContentResponseBody) *ModifyPolicyContentResponse {
	s.Body = v
	return s
}

type ModifyRemarkRequest struct {
	// The ID of the Anti-DDoS Origin instance for which you want to add remarks.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks for the Anti-DDoS Origin instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRemarkRequest) GoString() string {
	return s.String()
}

func (s *ModifyRemarkRequest) SetInstanceId(v string) *ModifyRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyRemarkRequest) SetRegionId(v string) *ModifyRemarkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRemarkRequest) SetRemark(v string) *ModifyRemarkRequest {
	s.Remark = &v
	return s
}

func (s *ModifyRemarkRequest) SetResourceGroupId(v string) *ModifyRemarkRequest {
	s.ResourceGroupId = &v
	return s
}

type ModifyRemarkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6AC3597B-7FD5-5E68-97C3-E11F4D010732
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRemarkResponseBody) SetRequestId(v string) *ModifyRemarkResponseBody {
	s.RequestId = &v
	return s
}

type ModifyRemarkResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRemarkResponse) GoString() string {
	return s.String()
}

func (s *ModifyRemarkResponse) SetHeaders(v map[string]*string) *ModifyRemarkResponse {
	s.Headers = v
	return s
}

func (s *ModifyRemarkResponse) SetStatusCode(v int32) *ModifyRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRemarkResponse) SetBody(v *ModifyRemarkResponseBody) *ModifyRemarkResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rg-acfm3peow3k****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ddos_originpre_public_cn-7213kxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetResourceGroupId(v string) *MoveResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceRegionId(v string) *MoveResourceGroupRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	// example:
	//
	// 16A78396-936F-5481-91D7-591BF7981246
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponse) SetHeaders(v map[string]*string) *MoveResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceGroupResponse) SetStatusCode(v int32) *MoveResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type QuerySchedruleOnDemandRequest struct {
	// The ID of the on-demand instance.
	//
	// >  You can call the [DescribeOnDemandInstance](https://help.aliyun.com/document_detail/152120.html) operation to query the IDs of all on-demand instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-z2q1qzxb****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the on-demand instance.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QuerySchedruleOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandRequest) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandRequest) SetInstanceId(v string) *QuerySchedruleOnDemandRequest {
	s.InstanceId = &v
	return s
}

func (s *QuerySchedruleOnDemandRequest) SetRegionId(v string) *QuerySchedruleOnDemandRequest {
	s.RegionId = &v
	return s
}

type QuerySchedruleOnDemandResponseBody struct {
	// The ID of the on-demand instance.
	//
	// example:
	//
	// ddosbgp-cn-z2q1qzxb****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4A8F9980-5ACB-497F-9F15-48E9D6B29028
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of the scheduling rule.
	RuleConfig []*QuerySchedruleOnDemandResponseBodyRuleConfig `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty" type:"Repeated"`
	// The status of the scheduling rule.
	RuleStatus []*QuerySchedruleOnDemandResponseBodyRuleStatus `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 171986973287****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuerySchedruleOnDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponseBody) SetInstanceId(v string) *QuerySchedruleOnDemandResponseBody {
	s.InstanceId = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBody) SetRequestId(v string) *QuerySchedruleOnDemandResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBody) SetRuleConfig(v []*QuerySchedruleOnDemandResponseBodyRuleConfig) *QuerySchedruleOnDemandResponseBody {
	s.RuleConfig = v
	return s
}

func (s *QuerySchedruleOnDemandResponseBody) SetRuleStatus(v []*QuerySchedruleOnDemandResponseBodyRuleStatus) *QuerySchedruleOnDemandResponseBody {
	s.RuleStatus = v
	return s
}

func (s *QuerySchedruleOnDemandResponseBody) SetUserId(v string) *QuerySchedruleOnDemandResponseBody {
	s.UserId = &v
	return s
}

type QuerySchedruleOnDemandResponseBodyRuleConfig struct {
	// The scheduling action. The value is set to **declare**, which indicates that the route is advertised.
	//
	// example:
	//
	// declare
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// If the inbound bandwidth or packets consecutively exceed the threshold for the specified number of times, the scheduling rule is triggered and traffic is rerouted to the on-demand instance. The specified number of times is the value of this parameter.
	//
	// >  The threshold of inbound bandwidth is the value of **RuleConditionMbps**. The threshold of inbound packets is the value of **RuleConditionKpps**.
	//
	// example:
	//
	// 3
	RuleConditionCnt *string `json:"RuleConditionCnt,omitempty" xml:"RuleConditionCnt,omitempty"`
	// The threshold of inbound packets. Unit: kilo packets per second (Kpps). Minimum value: **10**.
	//
	// example:
	//
	// 10
	RuleConditionKpps *string `json:"RuleConditionKpps,omitempty" xml:"RuleConditionKpps,omitempty"`
	// The threshold of inbound bandwidth. Unit: Mbit/s. Minimum value: **100**.
	//
	// example:
	//
	// 100
	RuleConditionMbps *string `json:"RuleConditionMbps,omitempty" xml:"RuleConditionMbps,omitempty"`
	// The name of the scheduling rule.
	//
	// example:
	//
	// ddosbgp-cn-z2q1qzxb****
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Indicates whether the scheduling rule is enabled. Valid values:
	//
	// 	- **on**: enabled.
	//
	// 	- **off**: disabled.
	//
	// example:
	//
	// on
	RuleSwitch *string `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	// The start time of the period during which the scheduling rule is automatically stopped. The time must be in the 24-hour clock and in the `hh:mm` format.
	//
	// If the system detects that DDoS attacks stop, the system no longer reroutes traffic to the on-demand instance from the time you specified. We recommend that you set this parameter to a value that is defined as off-peak hours.
	//
	// >  This parameter takes effect only when the value of **RuleUndoMode*	- is **auto**.
	//
	// example:
	//
	// 03:00
	RuleUndoBeginTime *string `json:"RuleUndoBeginTime,omitempty" xml:"RuleUndoBeginTime,omitempty"`
	// The end time of the period during which the scheduling rule is automatically stopped. The time must be in the 24-hour clock and in the `hh:mm` format.
	//
	// example:
	//
	// 03:05
	RuleUndoEndTime *string `json:"RuleUndoEndTime,omitempty" xml:"RuleUndoEndTime,omitempty"`
	// The stop method of the scheduling rule. Valid values:
	//
	// 	- **auto**
	//
	// 	- **manual**
	//
	// example:
	//
	// auto
	RuleUndoMode *string `json:"RuleUndoMode,omitempty" xml:"RuleUndoMode,omitempty"`
	// The time zone of the time when the scheduling rule automatically stops. The time zone must be in the `GMT-hh:mm` format.
	//
	// For example, the value `GMT-08:00` indicates that the time zone is UTC+8.
	//
	// >  This parameter takes effect only when the value of **RuleUndoMode*	- is **auto**.
	//
	// example:
	//
	// GMT-08:00
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s QuerySchedruleOnDemandResponseBodyRuleConfig) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponseBodyRuleConfig) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleAction(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleAction = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleConditionCnt(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleConditionCnt = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleConditionKpps(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleConditionKpps = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleConditionMbps(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleConditionMbps = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleName(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleName = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleSwitch(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleSwitch = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleUndoBeginTime(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleUndoBeginTime = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleUndoEndTime(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleUndoEndTime = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleUndoMode(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleUndoMode = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetTimeZone(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.TimeZone = &v
	return s
}

type QuerySchedruleOnDemandResponseBodyRuleStatus struct {
	// The CIDR block of the on-demand instance.
	//
	// example:
	//
	// 47.***.***.0/24
	Net *string `json:"Net,omitempty" xml:"Net,omitempty"`
	// The scheduling status. Valid values:
	//
	// 	- **scheduled**
	//
	// 	- **unscheduled**
	//
	// example:
	//
	// unscheduled
	RuleSchedStatus *string `json:"RuleSchedStatus,omitempty" xml:"RuleSchedStatus,omitempty"`
}

func (s QuerySchedruleOnDemandResponseBodyRuleStatus) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponseBodyRuleStatus) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponseBodyRuleStatus) SetNet(v string) *QuerySchedruleOnDemandResponseBodyRuleStatus {
	s.Net = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleStatus) SetRuleSchedStatus(v string) *QuerySchedruleOnDemandResponseBodyRuleStatus {
	s.RuleSchedStatus = &v
	return s
}

type QuerySchedruleOnDemandResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySchedruleOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySchedruleOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponse) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponse) SetHeaders(v map[string]*string) *QuerySchedruleOnDemandResponse {
	s.Headers = v
	return s
}

func (s *QuerySchedruleOnDemandResponse) SetStatusCode(v int32) *QuerySchedruleOnDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySchedruleOnDemandResponse) SetBody(v *QuerySchedruleOnDemandResponseBody) *QuerySchedruleOnDemandResponse {
	s.Body = v
	return s
}

type ReleaseDdosOriginInstanceRequest struct {
	// The ID of the Anti-DDoS Origin instance that you want to release.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosorigin_cn-pe335v7gs01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleaseDdosOriginInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseDdosOriginInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseDdosOriginInstanceRequest) SetInstanceId(v string) *ReleaseDdosOriginInstanceRequest {
	s.InstanceId = &v
	return s
}

type ReleaseDdosOriginInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B4B379C2-9319-4C6B-B579-FE36831B09F4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseDdosOriginInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseDdosOriginInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseDdosOriginInstanceResponseBody) SetRequestId(v string) *ReleaseDdosOriginInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseDdosOriginInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseDdosOriginInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseDdosOriginInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseDdosOriginInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseDdosOriginInstanceResponse) SetHeaders(v map[string]*string) *ReleaseDdosOriginInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseDdosOriginInstanceResponse) SetStatusCode(v int32) *ReleaseDdosOriginInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseDdosOriginInstanceResponse) SetBody(v *ReleaseDdosOriginInstanceResponseBody) *ReleaseDdosOriginInstanceResponse {
	s.Body = v
	return s
}

type SetInstanceModeOnDemandRequest struct {
	// The IDs of on-demand instances.
	//
	// > You can call the [DescribeOnDemandInstance](https://help.aliyun.com/document_detail/152120.html) operation to query the IDs of all on-demand instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-z2q1qzxb****
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	// Specifies the scheduling mode for on-demand instances. Valid values:
	//
	// 	- **manual**: manual scheduling
	//
	// 	- **netflow-auto**: automatic scheduling
	//
	// This parameter is required.
	//
	// example:
	//
	// netflow-auto
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The region ID of the on-demand instance.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query all regions that are supported by Anti-DDoS Origin.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetInstanceModeOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceModeOnDemandRequest) GoString() string {
	return s.String()
}

func (s *SetInstanceModeOnDemandRequest) SetInstanceIdList(v []*string) *SetInstanceModeOnDemandRequest {
	s.InstanceIdList = v
	return s
}

func (s *SetInstanceModeOnDemandRequest) SetMode(v string) *SetInstanceModeOnDemandRequest {
	s.Mode = &v
	return s
}

func (s *SetInstanceModeOnDemandRequest) SetRegionId(v string) *SetInstanceModeOnDemandRequest {
	s.RegionId = &v
	return s
}

type SetInstanceModeOnDemandResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// BD06F539-2FBE-450D-9391-7EFF787128F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetInstanceModeOnDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceModeOnDemandResponseBody) GoString() string {
	return s.String()
}

func (s *SetInstanceModeOnDemandResponseBody) SetRequestId(v string) *SetInstanceModeOnDemandResponseBody {
	s.RequestId = &v
	return s
}

type SetInstanceModeOnDemandResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetInstanceModeOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetInstanceModeOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceModeOnDemandResponse) GoString() string {
	return s.String()
}

func (s *SetInstanceModeOnDemandResponse) SetHeaders(v map[string]*string) *SetInstanceModeOnDemandResponse {
	s.Headers = v
	return s
}

func (s *SetInstanceModeOnDemandResponse) SetStatusCode(v int32) *SetInstanceModeOnDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *SetInstanceModeOnDemandResponse) SetBody(v *SetInstanceModeOnDemandResponseBody) *SetInstanceModeOnDemandResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The ID of the region in which the instance resides.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IDs of the instances to which you want to add tags. You can specify up to 51 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-v0h1fmwb****
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource to which you want to add tags. Set the value to **INSTANCE**, which indicates instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags to add. You can specify up to 21 tags.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *TagResourcesRequest) SetResourceGroupId(v string) *TagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The key of the tag to add.
	//
	// > If the specified key does not exist, a key is created.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to add.
	//
	// > If the specified tag value does not exist, the tag value is created.
	//
	// example:
	//
	// test-value
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
	// The ID of the request.
	//
	// example:
	//
	// 7078CD1E-F609-47A4-9C39-B288CC27C686
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the instances. Default value: No.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The ID of the region in which the instances reside.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IDs of the instances. Valid values of N: 0 to 49. You can specify up to 50 instances at a time. Example: ResourceId.0, ResourceId.1, ... , ResourceId.49.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-v0h1fmwbc024
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Set the value to **INSTANCE**.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The keys of the tag that you want to remove. Valid values of N: 0 to 19. You can specify up to 20 tag keys at a time. Example: Tag.0.Key, Tag.1.Key, ... , Tag.19.Key.
	//
	// example:
	//
	// testKey1
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceGroupId(v string) *UntagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F2D86AED-BA27-4584-BADC-B43BDA7EEBCA
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
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
	client.EndpointMap = map[string]*string{
		"cn-qingdao":            tea.String("ddosbgp.aliyuncs.com"),
		"cn-beijing":            tea.String("ddosbgp.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("ddosbgp.aliyuncs.com"),
		"cn-huhehaote":          tea.String("ddosbgp.aliyuncs.com"),
		"cn-hangzhou":           tea.String("ddosbgp.aliyuncs.com"),
		"cn-shanghai":           tea.String("ddosbgp.aliyuncs.com"),
		"cn-shenzhen":           tea.String("ddosbgp.aliyuncs.com"),
		"ap-northeast-1":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"cn-chengdu":            tea.String("ddosbgp.aliyuncs.com"),
		"eu-central-1":          tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("ddosbgp.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("ddosbgp.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("ddosbgp.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("ddosbgp.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ddosbgp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// Adds IP addresses to an Anti-DDoS Origin instance.
//
// @param request - AddIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddIpResponse
func (client *Client) AddIpWithOptions(request *AddIpRequest, runtime *util.RuntimeOptions) (_result *AddIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpList)) {
		query["IpList"] = request.IpList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddIp"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds IP addresses to an Anti-DDoS Origin instance.
//
// @param request - AddIpRequest
//
// @return AddIpResponse
func (client *Client) AddIp(request *AddIpRequest) (_result *AddIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddIpResponse{}
	_body, _err := client.AddIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds members to a resource directory.
//
// Description:
//
// Only a delegated administrator account or the management account of a resource directory can be used to add members to the resource directory.
//
// @param tmpReq - AddRdMemberListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRdMemberListResponse
func (client *Client) AddRdMemberListWithOptions(tmpReq *AddRdMemberListRequest, runtime *util.RuntimeOptions) (_result *AddRdMemberListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddRdMemberListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MemberList)) {
		request.MemberListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MemberList, tea.String("MemberList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberListShrink)) {
		query["MemberList"] = request.MemberListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddRdMemberList"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRdMemberListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds members to a resource directory.
//
// Description:
//
// Only a delegated administrator account or the management account of a resource directory can be used to add members to the resource directory.
//
// @param request - AddRdMemberListRequest
//
// @return AddRdMemberListResponse
func (client *Client) AddRdMemberList(request *AddRdMemberListRequest) (_result *AddRdMemberListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddRdMemberListResponse{}
	_body, _err := client.AddRdMemberListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates an asset with an Anti-DDoS Origin instance of a paid edition.
//
// @param tmpReq - AttachAssetGroupToInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachAssetGroupToInstanceResponse
func (client *Client) AttachAssetGroupToInstanceWithOptions(tmpReq *AttachAssetGroupToInstanceRequest, runtime *util.RuntimeOptions) (_result *AttachAssetGroupToInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AttachAssetGroupToInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AssetGroupList)) {
		request.AssetGroupListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetGroupList, tea.String("AssetGroupList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetGroupListShrink)) {
		query["AssetGroupList"] = request.AssetGroupListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachAssetGroupToInstance"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachAssetGroupToInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates an asset with an Anti-DDoS Origin instance of a paid edition.
//
// @param request - AttachAssetGroupToInstanceRequest
//
// @return AttachAssetGroupToInstanceResponse
func (client *Client) AttachAssetGroupToInstance(request *AttachAssetGroupToInstanceRequest) (_result *AttachAssetGroupToInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachAssetGroupToInstanceResponse{}
	_body, _err := client.AttachAssetGroupToInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a mitigation policy to a protected object.
//
// @param tmpReq - AttachToPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachToPolicyResponse
func (client *Client) AttachToPolicyWithOptions(tmpReq *AttachToPolicyRequest, runtime *util.RuntimeOptions) (_result *AttachToPolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AttachToPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.IpPortProtocolList)) {
		request.IpPortProtocolListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IpPortProtocolList, tea.String("IpPortProtocolList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpPortProtocolListShrink)) {
		query["IpPortProtocolList"] = request.IpPortProtocolListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachToPolicy"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachToPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a mitigation policy to a protected object.
//
// @param request - AttachToPolicyRequest
//
// @return AttachToPolicyResponse
func (client *Client) AttachToPolicy(request *AttachToPolicyRequest) (_result *AttachToPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachToPolicyResponse{}
	_body, _err := client.AttachToPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether Anti-DDoS Origin is authorized to access Simple Log Service.
//
// @param request - CheckAccessLogAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAccessLogAuthResponse
func (client *Client) CheckAccessLogAuthWithOptions(request *CheckAccessLogAuthRequest, runtime *util.RuntimeOptions) (_result *CheckAccessLogAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckAccessLogAuth"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckAccessLogAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether Anti-DDoS Origin is authorized to access Simple Log Service.
//
// @param request - CheckAccessLogAuthRequest
//
// @return CheckAccessLogAuthResponse
func (client *Client) CheckAccessLogAuth(request *CheckAccessLogAuthRequest) (_result *CheckAccessLogAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckAccessLogAuthResponse{}
	_body, _err := client.CheckAccessLogAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether Anti-DDoS Origin is authorized to obtain information about the assets within the current Alibaba Cloud account.
//
// Description:
//
// You can call the CheckGrant operation to query whether Anti-DDoS Origin is authorized to obtain information about the assets within the current Alibaba Cloud account.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CheckGrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckGrantResponse
func (client *Client) CheckGrantWithOptions(request *CheckGrantRequest, runtime *util.RuntimeOptions) (_result *CheckGrantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckGrant"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckGrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether Anti-DDoS Origin is authorized to obtain information about the assets within the current Alibaba Cloud account.
//
// Description:
//
// You can call the CheckGrant operation to query whether Anti-DDoS Origin is authorized to obtain information about the assets within the current Alibaba Cloud account.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CheckGrantRequest
//
// @return CheckGrantResponse
func (client *Client) CheckGrant(request *CheckGrantRequest) (_result *CheckGrantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckGrantResponse{}
	_body, _err := client.CheckGrantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a scheduling rule of an on-demand instance.
//
// @param request - ConfigSchedruleOnDemandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigSchedruleOnDemandResponse
func (client *Client) ConfigSchedruleOnDemandWithOptions(request *ConfigSchedruleOnDemandRequest, runtime *util.RuntimeOptions) (_result *ConfigSchedruleOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleAction)) {
		query["RuleAction"] = request.RuleAction
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionCnt)) {
		query["RuleConditionCnt"] = request.RuleConditionCnt
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionKpps)) {
		query["RuleConditionKpps"] = request.RuleConditionKpps
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionMbps)) {
		query["RuleConditionMbps"] = request.RuleConditionMbps
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleSwitch)) {
		query["RuleSwitch"] = request.RuleSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoBeginTime)) {
		query["RuleUndoBeginTime"] = request.RuleUndoBeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoEndTime)) {
		query["RuleUndoEndTime"] = request.RuleUndoEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoMode)) {
		query["RuleUndoMode"] = request.RuleUndoMode
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		query["TimeZone"] = request.TimeZone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigSchedruleOnDemand"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigSchedruleOnDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a scheduling rule of an on-demand instance.
//
// @param request - ConfigSchedruleOnDemandRequest
//
// @return ConfigSchedruleOnDemandResponse
func (client *Client) ConfigSchedruleOnDemand(request *ConfigSchedruleOnDemandRequest) (_result *ConfigSchedruleOnDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigSchedruleOnDemandResponse{}
	_body, _err := client.ConfigSchedruleOnDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a mitigation policy.
//
// @param request - CreatePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicyWithOptions(request *CreatePolicyRequest, runtime *util.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePolicy"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a mitigation policy.
//
// @param request - CreatePolicyRequest
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicy(request *CreatePolicyRequest) (_result *CreatePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CreatePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a scheduling rule for an on-demand instance.
//
// @param request - CreateSchedruleOnDemandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSchedruleOnDemandResponse
func (client *Client) CreateSchedruleOnDemandWithOptions(request *CreateSchedruleOnDemandRequest, runtime *util.RuntimeOptions) (_result *CreateSchedruleOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleAction)) {
		query["RuleAction"] = request.RuleAction
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionCnt)) {
		query["RuleConditionCnt"] = request.RuleConditionCnt
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionKpps)) {
		query["RuleConditionKpps"] = request.RuleConditionKpps
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionMbps)) {
		query["RuleConditionMbps"] = request.RuleConditionMbps
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleSwitch)) {
		query["RuleSwitch"] = request.RuleSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoBeginTime)) {
		query["RuleUndoBeginTime"] = request.RuleUndoBeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoEndTime)) {
		query["RuleUndoEndTime"] = request.RuleUndoEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoMode)) {
		query["RuleUndoMode"] = request.RuleUndoMode
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		query["TimeZone"] = request.TimeZone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSchedruleOnDemand"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSchedruleOnDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scheduling rule for an on-demand instance.
//
// @param request - CreateSchedruleOnDemandRequest
//
// @return CreateSchedruleOnDemandResponse
func (client *Client) CreateSchedruleOnDemand(request *CreateSchedruleOnDemandRequest) (_result *CreateSchedruleOnDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSchedruleOnDemandResponse{}
	_body, _err := client.CreateSchedruleOnDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deactivates blackhole filtering for a protected IP address.
//
// Description:
//
// You can call the DeleteBlackhole operation to deactivate blackhole filtering for a protected IP address.
//
// Before you call this operation, you can call the [DescribePackIpList](https://help.aliyun.com/document_detail/118701.html) operation to query the protection status of the IP addresses that are protected by your Anti-DDoS Origin instance. For example, you can query whether blackhole filtering is triggered for an IP address.
//
// ### [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteBlackholeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBlackholeResponse
func (client *Client) DeleteBlackholeWithOptions(request *DeleteBlackholeRequest, runtime *util.RuntimeOptions) (_result *DeleteBlackholeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBlackhole"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBlackholeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deactivates blackhole filtering for a protected IP address.
//
// Description:
//
// You can call the DeleteBlackhole operation to deactivate blackhole filtering for a protected IP address.
//
// Before you call this operation, you can call the [DescribePackIpList](https://help.aliyun.com/document_detail/118701.html) operation to query the protection status of the IP addresses that are protected by your Anti-DDoS Origin instance. For example, you can query whether blackhole filtering is triggered for an IP address.
//
// ### [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteBlackholeRequest
//
// @return DeleteBlackholeResponse
func (client *Client) DeleteBlackhole(request *DeleteBlackholeRequest) (_result *DeleteBlackholeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBlackholeResponse{}
	_body, _err := client.DeleteBlackholeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes specific IP addresses from an Anti-DDoS Origin instance.
//
// Description:
//
// The Anti-DDoS Origin Enterprise instance no longer protects the IP addresses that are removed.
//
// @param request - DeleteIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpResponse
func (client *Client) DeleteIpWithOptions(request *DeleteIpRequest, runtime *util.RuntimeOptions) (_result *DeleteIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpList)) {
		query["IpList"] = request.IpList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIp"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes specific IP addresses from an Anti-DDoS Origin instance.
//
// Description:
//
// The Anti-DDoS Origin Enterprise instance no longer protects the IP addresses that are removed.
//
// @param request - DeleteIpRequest
//
// @return DeleteIpResponse
func (client *Client) DeleteIp(request *DeleteIpRequest) (_result *DeleteIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpResponse{}
	_body, _err := client.DeleteIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a mitigation policy.
//
// Description:
//
// You cannot delete a mitigation policy to which a protected object is added.
//
// @param request - DeletePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicyWithOptions(request *DeletePolicyRequest, runtime *util.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePolicy"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a mitigation policy.
//
// Description:
//
// You cannot delete a mitigation policy to which a protected object is added.
//
// @param request - DeletePolicyRequest
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicy(request *DeletePolicyRequest) (_result *DeletePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePolicyResponse{}
	_body, _err := client.DeletePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes members.
//
// Description:
//
// Only a delegated administrator account or the management account of a resource directory can be used to delete members.
//
// @param tmpReq - DeleteRdMemberListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRdMemberListResponse
func (client *Client) DeleteRdMemberListWithOptions(tmpReq *DeleteRdMemberListRequest, runtime *util.RuntimeOptions) (_result *DeleteRdMemberListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteRdMemberListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MemberList)) {
		request.MemberListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MemberList, tea.String("MemberList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberListShrink)) {
		query["MemberList"] = request.MemberListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRdMemberList"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRdMemberListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes members.
//
// Description:
//
// Only a delegated administrator account or the management account of a resource directory can be used to delete members.
//
// @param request - DeleteRdMemberListRequest
//
// @return DeleteRdMemberListResponse
func (client *Client) DeleteRdMemberList(request *DeleteRdMemberListRequest) (_result *DeleteRdMemberListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRdMemberListResponse{}
	_body, _err := client.DeleteRdMemberListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a scheduling rule of an anti-DDoS diversion instance.
//
// @param request - DeleteSchedruleOnDemandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSchedruleOnDemandResponse
func (client *Client) DeleteSchedruleOnDemandWithOptions(request *DeleteSchedruleOnDemandRequest, runtime *util.RuntimeOptions) (_result *DeleteSchedruleOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSchedruleOnDemand"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSchedruleOnDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scheduling rule of an anti-DDoS diversion instance.
//
// @param request - DeleteSchedruleOnDemandRequest
//
// @return DeleteSchedruleOnDemandResponse
func (client *Client) DeleteSchedruleOnDemand(request *DeleteSchedruleOnDemandRequest) (_result *DeleteSchedruleOnDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSchedruleOnDemandResponse{}
	_body, _err := client.DeleteSchedruleOnDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the association between an asset and an Anti-DDoS Origin instance of a paid edition.
//
// @param request - DescribeAssetGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAssetGroupResponse
func (client *Client) DescribeAssetGroupWithOptions(request *DescribeAssetGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeAssetGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAssetGroup"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAssetGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the association between an asset and an Anti-DDoS Origin instance of a paid edition.
//
// @param request - DescribeAssetGroupRequest
//
// @return DescribeAssetGroupResponse
func (client *Client) DescribeAssetGroup(request *DescribeAssetGroupRequest) (_result *DescribeAssetGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAssetGroupResponse{}
	_body, _err := client.DescribeAssetGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the association between an asset and an Anti-DDoS Origin instance of a paid edition.
//
// @param request - DescribeAssetGroupToInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAssetGroupToInstanceResponse
func (client *Client) DescribeAssetGroupToInstanceWithOptions(request *DescribeAssetGroupToInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeAssetGroupToInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAssetGroupToInstance"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAssetGroupToInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the association between an asset and an Anti-DDoS Origin instance of a paid edition.
//
// @param request - DescribeAssetGroupToInstanceRequest
//
// @return DescribeAssetGroupToInstanceResponse
func (client *Client) DescribeAssetGroupToInstance(request *DescribeAssetGroupToInstanceRequest) (_result *DescribeAssetGroupToInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAssetGroupToInstanceResponse{}
	_body, _err := client.DescribeAssetGroupToInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about the DDoS attack events that occurred on an Anti-DDoS Origin instance.
//
// Description:
//
// You can call the DescribeDdosEvent operation to query the details about the DDoS attack events that occurred on an Anti-DDoS Origin instance by page. The details include the start time, end time, attacked IP address, and status of each event.
//
// ### [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDdosEventResponse
func (client *Client) DescribeDdosEventWithOptions(request *DescribeDdosEventRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDdosEvent"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDdosEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about the DDoS attack events that occurred on an Anti-DDoS Origin instance.
//
// Description:
//
// You can call the DescribeDdosEvent operation to query the details about the DDoS attack events that occurred on an Anti-DDoS Origin instance by page. The details include the start time, end time, attacked IP address, and status of each event.
//
// ### [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDdosEventRequest
//
// @return DescribeDdosEventResponse
func (client *Client) DescribeDdosEvent(request *DescribeDdosEventRequest) (_result *DescribeDdosEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosEventResponse{}
	_body, _err := client.DescribeDdosEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the bill of an Anti-DDoS Origin (Pay-as-you-go) instance.
//
// @param request - DescribeDdosOriginInstanceBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDdosOriginInstanceBillResponse
func (client *Client) DescribeDdosOriginInstanceBillWithOptions(request *DescribeDdosOriginInstanceBillRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosOriginInstanceBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsShowList)) {
		query["IsShowList"] = request.IsShowList
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDdosOriginInstanceBill"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDdosOriginInstanceBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bill of an Anti-DDoS Origin (Pay-as-you-go) instance.
//
// @param request - DescribeDdosOriginInstanceBillRequest
//
// @return DescribeDdosOriginInstanceBillResponse
func (client *Client) DescribeDdosOriginInstanceBill(request *DescribeDdosOriginInstanceBillRequest) (_result *DescribeDdosOriginInstanceBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosOriginInstanceBillResponse{}
	_body, _err := client.DescribeDdosOriginInstanceBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of assets that are in an abnormal state and the number of Anti-DDoS Origin instances that are about to expire. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses.
//
// Description:
//
// ## Usage notes
//
// You can call the DescribeExcpetionCount operation to query the number of assets that are in an abnormal state and the number of Anti-DDoS Origin instances that are about to expire in a specific region. For example, if blackhole filtering is triggered for an IP address, the IP address is in an abnormal state. An instance whose remaining validity period is less than seven days is considered as an instance that is about to expire.
//
// @param request - DescribeExcpetionCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExcpetionCountResponse
func (client *Client) DescribeExcpetionCountWithOptions(request *DescribeExcpetionCountRequest, runtime *util.RuntimeOptions) (_result *DescribeExcpetionCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExcpetionCount"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExcpetionCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of assets that are in an abnormal state and the number of Anti-DDoS Origin instances that are about to expire. The assets can be elastic IP addresses (EIPs). The assets can also be Elastic Compute Service (ECS) instances or Server Load Balancer (SLB) instances that are assigned public IP addresses.
//
// Description:
//
// ## Usage notes
//
// You can call the DescribeExcpetionCount operation to query the number of assets that are in an abnormal state and the number of Anti-DDoS Origin instances that are about to expire in a specific region. For example, if blackhole filtering is triggered for an IP address, the IP address is in an abnormal state. An instance whose remaining validity period is less than seven days is considered as an instance that is about to expire.
//
// @param request - DescribeExcpetionCountRequest
//
// @return DescribeExcpetionCountResponse
func (client *Client) DescribeExcpetionCount(request *DescribeExcpetionCountRequest) (_result *DescribeExcpetionCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExcpetionCountResponse{}
	_body, _err := client.DescribeExcpetionCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of all Anti-DDoS Origin instances.
//
// Description:
//
// You can call the DescribeInstanceList operation to query the details of all Anti-DDoS Origin instances within your Alibaba Cloud account by page. The details include the ID, validity period, and status of each instance.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceListResponse
func (client *Client) DescribeInstanceListWithOptions(request *DescribeInstanceListRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceTypeList)) {
		query["InstanceTypeList"] = request.InstanceTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Orderby)) {
		query["Orderby"] = request.Orderby
	}

	if !tea.BoolValue(util.IsUnset(request.Orderdire)) {
		query["Orderdire"] = request.Orderdire
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceList"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of all Anti-DDoS Origin instances.
//
// Description:
//
// You can call the DescribeInstanceList operation to query the details of all Anti-DDoS Origin instances within your Alibaba Cloud account by page. The details include the ID, validity period, and status of each instance.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceListRequest
//
// @return DescribeInstanceListResponse
func (client *Client) DescribeInstanceList(request *DescribeInstanceListRequest) (_result *DescribeInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceListResponse{}
	_body, _err := client.DescribeInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the specifications of a specific Anti-DDoS Origin instance.
//
// @param request - DescribeInstanceSpecsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceSpecsResponse
func (client *Client) DescribeInstanceSpecsWithOptions(request *DescribeInstanceSpecsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSpecsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceSpecs"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the specifications of a specific Anti-DDoS Origin instance.
//
// @param request - DescribeInstanceSpecsRequest
//
// @return DescribeInstanceSpecsResponse
func (client *Client) DescribeInstanceSpecs(request *DescribeInstanceSpecsRequest) (_result *DescribeInstanceSpecsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.DescribeInstanceSpecsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the DDoS attack events recorded for the IP address of an anti-DDoS diversion instance of Anti-DDoS Origin.
//
// Description:
//
// You can use this operation to query the details about the DDoS attack events that occurred on the IP address of an anti-DDoS diversion instance of Anti-DDoS Origin by page. The details include the start time, end time, volume of attack traffic, and status of each event.
//
// ### [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeOnDemandDdosEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOnDemandDdosEventResponse
func (client *Client) DescribeOnDemandDdosEventWithOptions(request *DescribeOnDemandDdosEventRequest, runtime *util.RuntimeOptions) (_result *DescribeOnDemandDdosEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOnDemandDdosEvent"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOnDemandDdosEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the DDoS attack events recorded for the IP address of an anti-DDoS diversion instance of Anti-DDoS Origin.
//
// Description:
//
// You can use this operation to query the details about the DDoS attack events that occurred on the IP address of an anti-DDoS diversion instance of Anti-DDoS Origin by page. The details include the start time, end time, volume of attack traffic, and status of each event.
//
// ### [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeOnDemandDdosEventRequest
//
// @return DescribeOnDemandDdosEventResponse
func (client *Client) DescribeOnDemandDdosEvent(request *DescribeOnDemandDdosEventRequest) (_result *DescribeOnDemandDdosEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOnDemandDdosEventResponse{}
	_body, _err := client.DescribeOnDemandDdosEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of anti-DDoS diversion instances.
//
// @param request - DescribeOnDemandInstanceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOnDemandInstanceStatusResponse
func (client *Client) DescribeOnDemandInstanceStatusWithOptions(request *DescribeOnDemandInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeOnDemandInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOnDemandInstanceStatus"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOnDemandInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of anti-DDoS diversion instances.
//
// @param request - DescribeOnDemandInstanceStatusRequest
//
// @return DescribeOnDemandInstanceStatusResponse
func (client *Client) DescribeOnDemandInstanceStatus(request *DescribeOnDemandInstanceStatusRequest) (_result *DescribeOnDemandInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOnDemandInstanceStatusResponse{}
	_body, _err := client.DescribeOnDemandInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operation logs of an Anti-DDoS Origin instance.
//
// Description:
//
// You can call the DescribeOpEntities operation to query the operation logs of an instance by page.
//
// ### Limit
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeOpEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOpEntitiesResponse
func (client *Client) DescribeOpEntitiesWithOptions(request *DescribeOpEntitiesRequest, runtime *util.RuntimeOptions) (_result *DescribeOpEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OpAction)) {
		query["OpAction"] = request.OpAction
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderDir)) {
		query["OrderDir"] = request.OrderDir
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOpEntities"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operation logs of an Anti-DDoS Origin instance.
//
// Description:
//
// You can call the DescribeOpEntities operation to query the operation logs of an instance by page.
//
// ### Limit
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeOpEntitiesRequest
//
// @return DescribeOpEntitiesResponse
func (client *Client) DescribeOpEntities(request *DescribeOpEntitiesRequest) (_result *DescribeOpEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.DescribeOpEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IP addresses that are protected by a specific Anti-DDoS Origin instance.
//
// Description:
//
// You can call the DescribePackIpList operation to query the details about each IP address that is protected by a specific Anti-DDoS Origin instance by page. The details include the IP address and the type of the cloud asset to which the IP address belongs. The details also include the status of the IP address, such as whether blackhole filtering is triggered for the IP address.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePackIpListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePackIpListResponse
func (client *Client) DescribePackIpListWithOptions(request *DescribePackIpListRequest, runtime *util.RuntimeOptions) (_result *DescribePackIpListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		query["ProductName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePackIpList"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePackIpListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP addresses that are protected by a specific Anti-DDoS Origin instance.
//
// Description:
//
// You can call the DescribePackIpList operation to query the details about each IP address that is protected by a specific Anti-DDoS Origin instance by page. The details include the IP address and the type of the cloud asset to which the IP address belongs. The details also include the status of the IP address, such as whether blackhole filtering is triggered for the IP address.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePackIpListRequest
//
// @return DescribePackIpListResponse
func (client *Client) DescribePackIpList(request *DescribePackIpListRequest) (_result *DescribePackIpListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePackIpListResponse{}
	_body, _err := client.DescribePackIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries members that are managed by using the multi-account management feature.
//
// @param request - DescribeRdMemberListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRdMemberListResponse
func (client *Client) DescribeRdMemberListWithOptions(request *DescribeRdMemberListRequest, runtime *util.RuntimeOptions) (_result *DescribeRdMemberListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceDirectoryId)) {
		query["ResourceDirectoryId"] = request.ResourceDirectoryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRdMemberList"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRdMemberListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries members that are managed by using the multi-account management feature.
//
// @param request - DescribeRdMemberListRequest
//
// @return DescribeRdMemberListResponse
func (client *Client) DescribeRdMemberList(request *DescribeRdMemberListRequest) (_result *DescribeRdMemberListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdMemberListResponse{}
	_body, _err := client.DescribeRdMemberListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the multi-account management feature.
//
// @param request - DescribeRdStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRdStatusResponse
func (client *Client) DescribeRdStatusWithOptions(runtime *util.RuntimeOptions) (_result *DescribeRdStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeRdStatus"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRdStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the multi-account management feature.
//
// @return DescribeRdStatusResponse
func (client *Client) DescribeRdStatus() (_result *DescribeRdStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdStatusResponse{}
	_body, _err := client.DescribeRdStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions of assets that can be protected by Anti-DDoS Origin Enterprise in a specific region.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions of assets that can be protected by Anti-DDoS Origin Enterprise in a specific region.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
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

// Summary:
//
// Queries traffic statistics of an Anti-DDoS Origin instance within a specific time period.
//
// Description:
//
// You can call the DescribeTraffic operation to query traffic statistics of an Anti-DDoS Origin instance within a specific time period.
//
// >  When you call this operation, you must configure the **InstanceId*	- parameter to specify the Anti-DDoS Origin instance whose traffic statistics you want to query.
//
// ## Limits
//
// You can call this operation once per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeTrafficRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrafficResponse
func (client *Client) DescribeTrafficWithOptions(request *DescribeTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribeTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FlowType)) {
		query["FlowType"] = request.FlowType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Ipnet)) {
		query["Ipnet"] = request.Ipnet
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTraffic"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTrafficResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries traffic statistics of an Anti-DDoS Origin instance within a specific time period.
//
// Description:
//
// You can call the DescribeTraffic operation to query traffic statistics of an Anti-DDoS Origin instance within a specific time period.
//
// >  When you call this operation, you must configure the **InstanceId*	- parameter to specify the Anti-DDoS Origin instance whose traffic statistics you want to query.
//
// ## Limits
//
// You can call this operation once per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeTrafficRequest
//
// @return DescribeTrafficResponse
func (client *Client) DescribeTraffic(request *DescribeTrafficRequest) (_result *DescribeTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTrafficResponse{}
	_body, _err := client.DescribeTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes protected objects from a mitigation policy.
//
// @param tmpReq - DetachFromPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachFromPolicyResponse
func (client *Client) DetachFromPolicyWithOptions(tmpReq *DetachFromPolicyRequest, runtime *util.RuntimeOptions) (_result *DetachFromPolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetachFromPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.IpPortProtocolList)) {
		request.IpPortProtocolListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IpPortProtocolList, tea.String("IpPortProtocolList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpPortProtocolListShrink)) {
		query["IpPortProtocolList"] = request.IpPortProtocolListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachFromPolicy"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachFromPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes protected objects from a mitigation policy.
//
// @param request - DetachFromPolicyRequest
//
// @return DetachFromPolicyResponse
func (client *Client) DetachFromPolicy(request *DetachFromPolicyRequest) (_result *DetachFromPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachFromPolicyResponse{}
	_body, _err := client.DetachFromPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Dissociates an asset from an Anti-DDoS Origin instance of a paid edition.
//
// @param tmpReq - DettachAssetGroupToInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DettachAssetGroupToInstanceResponse
func (client *Client) DettachAssetGroupToInstanceWithOptions(tmpReq *DettachAssetGroupToInstanceRequest, runtime *util.RuntimeOptions) (_result *DettachAssetGroupToInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DettachAssetGroupToInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AssetGroupList)) {
		request.AssetGroupListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetGroupList, tea.String("AssetGroupList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetGroupListShrink)) {
		query["AssetGroupList"] = request.AssetGroupListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DettachAssetGroupToInstance"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DettachAssetGroupToInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dissociates an asset from an Anti-DDoS Origin instance of a paid edition.
//
// @param request - DettachAssetGroupToInstanceRequest
//
// @return DettachAssetGroupToInstanceResponse
func (client *Client) DettachAssetGroupToInstance(request *DettachAssetGroupToInstanceRequest) (_result *DettachAssetGroupToInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DettachAssetGroupToInstanceResponse{}
	_body, _err := client.DettachAssetGroupToInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether Simple Log Service is activated.
//
// @param request - GetSlsOpenStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSlsOpenStatusResponse
func (client *Client) GetSlsOpenStatusWithOptions(request *GetSlsOpenStatusRequest, runtime *util.RuntimeOptions) (_result *GetSlsOpenStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSlsOpenStatus"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSlsOpenStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether Simple Log Service is activated.
//
// @param request - GetSlsOpenStatusRequest
//
// @return GetSlsOpenStatusResponse
func (client *Client) GetSlsOpenStatus(request *GetSlsOpenStatusRequest) (_result *GetSlsOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSlsOpenStatusResponse{}
	_body, _err := client.GetSlsOpenStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Anti-DDoS Origin instances for which log analysis is enabled.
//
// @param request - ListOpenedAccessLogInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOpenedAccessLogInstancesResponse
func (client *Client) ListOpenedAccessLogInstancesWithOptions(request *ListOpenedAccessLogInstancesRequest, runtime *util.RuntimeOptions) (_result *ListOpenedAccessLogInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOpenedAccessLogInstances"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOpenedAccessLogInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Anti-DDoS Origin instances for which log analysis is enabled.
//
// @param request - ListOpenedAccessLogInstancesRequest
//
// @return ListOpenedAccessLogInstancesResponse
func (client *Client) ListOpenedAccessLogInstances(request *ListOpenedAccessLogInstancesRequest) (_result *ListOpenedAccessLogInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOpenedAccessLogInstancesResponse{}
	_body, _err := client.ListOpenedAccessLogInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries mitigation policies.
//
// @param request - ListPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicyResponse
func (client *Client) ListPolicyWithOptions(request *ListPolicyRequest, runtime *util.RuntimeOptions) (_result *ListPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPolicy"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries mitigation policies.
//
// @param request - ListPolicyRequest
//
// @return ListPolicyResponse
func (client *Client) ListPolicy(request *ListPolicyRequest) (_result *ListPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPolicyResponse{}
	_body, _err := client.ListPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries attachments to mitigation policies.
//
// @param tmpReq - ListPolicyAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicyAttachmentResponse
func (client *Client) ListPolicyAttachmentWithOptions(tmpReq *ListPolicyAttachmentRequest, runtime *util.RuntimeOptions) (_result *ListPolicyAttachmentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListPolicyAttachmentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.IpPortProtocolList)) {
		request.IpPortProtocolListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IpPortProtocolList, tea.String("IpPortProtocolList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpPortProtocolListShrink)) {
		query["IpPortProtocolList"] = request.IpPortProtocolListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPolicyAttachment"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPolicyAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries attachments to mitigation policies.
//
// @param request - ListPolicyAttachmentRequest
//
// @return ListPolicyAttachmentResponse
func (client *Client) ListPolicyAttachment(request *ListPolicyAttachmentRequest) (_result *ListPolicyAttachmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPolicyAttachmentResponse{}
	_body, _err := client.ListPolicyAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all tags.
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all tags.
//
// @param request - ListTagKeysRequest
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the relationship between Anti-DDoS Origin instances and tags.
//
// Description:
//
// You can call the ListTagResources operation to query the tags that are added to Anti-DDoS Origin instances at a time.
//
// ### [](#qps-)Limits
//
// You can call this API operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the relationship between Anti-DDoS Origin instances and tags.
//
// Description:
//
// You can call the ListTagResources operation to query the tags that are added to Anti-DDoS Origin instances at a time.
//
// ### [](#qps-)Limits
//
// You can call this API operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
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

// Summary:
//
// Modifies a mitigation policy.
//
// @param tmpReq - ModifyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPolicyResponse
func (client *Client) ModifyPolicyWithOptions(tmpReq *ModifyPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyPolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Content)) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, tea.String("Content"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		query["ActionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.ContentShrink)) {
		query["Content"] = request.ContentShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPolicy"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a mitigation policy.
//
// @param request - ModifyPolicyRequest
//
// @return ModifyPolicyResponse
func (client *Client) ModifyPolicy(request *ModifyPolicyRequest) (_result *ModifyPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPolicyResponse{}
	_body, _err := client.ModifyPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the content of the mitigation policy.
//
// Description:
//
// Make sure that all request parameters are configured when you call this operation. If any parameter is left empty, the configuration is deleted.
//
// @param tmpReq - ModifyPolicyContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPolicyContentResponse
func (client *Client) ModifyPolicyContentWithOptions(tmpReq *ModifyPolicyContentRequest, runtime *util.RuntimeOptions) (_result *ModifyPolicyContentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyPolicyContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Content)) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, tea.String("Content"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentShrink)) {
		query["Content"] = request.ContentShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPolicyContent"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPolicyContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the content of the mitigation policy.
//
// Description:
//
// Make sure that all request parameters are configured when you call this operation. If any parameter is left empty, the configuration is deleted.
//
// @param request - ModifyPolicyContentRequest
//
// @return ModifyPolicyContentResponse
func (client *Client) ModifyPolicyContent(request *ModifyPolicyContentRequest) (_result *ModifyPolicyContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPolicyContentResponse{}
	_body, _err := client.ModifyPolicyContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds remarks for a single Anti-DDoS Origin instance.
//
// Description:
//
// You can call the ModifyRemark operation to add remarks for a single Anti-DDoS Origin instance.
//
// ### [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRemarkResponse
func (client *Client) ModifyRemarkWithOptions(request *ModifyRemarkRequest, runtime *util.RuntimeOptions) (_result *ModifyRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyRemark"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds remarks for a single Anti-DDoS Origin instance.
//
// Description:
//
// You can call the ModifyRemark operation to add remarks for a single Anti-DDoS Origin instance.
//
// ### [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyRemarkRequest
//
// @return ModifyRemarkResponse
func (client *Client) ModifyRemark(request *ModifyRemarkRequest) (_result *ModifyRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyRemarkResponse{}
	_body, _err := client.ModifyRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移动资源组
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResourceGroup"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移动资源组
//
// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the scheduling rule of an on-demand instance.
//
// @param request - QuerySchedruleOnDemandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySchedruleOnDemandResponse
func (client *Client) QuerySchedruleOnDemandWithOptions(request *QuerySchedruleOnDemandRequest, runtime *util.RuntimeOptions) (_result *QuerySchedruleOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySchedruleOnDemand"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySchedruleOnDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the scheduling rule of an on-demand instance.
//
// @param request - QuerySchedruleOnDemandRequest
//
// @return QuerySchedruleOnDemandResponse
func (client *Client) QuerySchedruleOnDemand(request *QuerySchedruleOnDemandRequest) (_result *QuerySchedruleOnDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySchedruleOnDemandResponse{}
	_body, _err := client.QuerySchedruleOnDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go Anti-DDoS Origin instance.
//
// @param request - ReleaseDdosOriginInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseDdosOriginInstanceResponse
func (client *Client) ReleaseDdosOriginInstanceWithOptions(request *ReleaseDdosOriginInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleaseDdosOriginInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseDdosOriginInstance"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseDdosOriginInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go Anti-DDoS Origin instance.
//
// @param request - ReleaseDdosOriginInstanceRequest
//
// @return ReleaseDdosOriginInstanceResponse
func (client *Client) ReleaseDdosOriginInstance(request *ReleaseDdosOriginInstanceRequest) (_result *ReleaseDdosOriginInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseDdosOriginInstanceResponse{}
	_body, _err := client.ReleaseDdosOriginInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies the scheduling mode for on-demand instances.
//
// @param request - SetInstanceModeOnDemandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetInstanceModeOnDemandResponse
func (client *Client) SetInstanceModeOnDemandWithOptions(request *SetInstanceModeOnDemandRequest, runtime *util.RuntimeOptions) (_result *SetInstanceModeOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetInstanceModeOnDemand"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetInstanceModeOnDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies the scheduling mode for on-demand instances.
//
// @param request - SetInstanceModeOnDemandRequest
//
// @return SetInstanceModeOnDemandResponse
func (client *Client) SetInstanceModeOnDemand(request *SetInstanceModeOnDemandRequest) (_result *SetInstanceModeOnDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetInstanceModeOnDemandResponse{}
	_body, _err := client.SetInstanceModeOnDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add tags to Anti-DDoS Origin instances.
//
// Description:
//
// You can call the TagResources operation to add tags to instances.
//
// ### Limit
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add tags to Anti-DDoS Origin instances.
//
// Description:
//
// You can call the TagResources operation to add tags to instances.
//
// ### Limit
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
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

// Summary:
//
// Removes tags from Anti-DDoS Origin instances.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from Anti-DDoS Origin instances.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
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
