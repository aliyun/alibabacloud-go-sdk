// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateAndAnalyzeNetworkPathRequest struct {
	// The protocol type. Valid values:
	//
	// 	- **tcp**: Transmission Control Protocol (TCP)
	//
	// 	- **udp**: User Datagram Protocol (UDP)
	//
	// 	- **icmp**: Internet Control Message Protocol (ICMP)
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the region for which you want to initiate a task for analyzing network reachability.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the source resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-uf62y8khhbkbdrp6****
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	// The source port.
	//
	// example:
	//
	// 0
	SourcePort *int32 `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The type of the source resource. Valid values:
	//
	// 	- **ecs**: the Elastic Compute Service (ECS) instance
	//
	// 	- **internetIp**: the public IP address
	//
	// 	- **vsw**: the vSwitch
	//
	// 	- **vpn**: the VPN gateway
	//
	// 	- **vbr**: the virtual border router (VBR)
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The ID of the destination resource.
	//
	// example:
	//
	// i-m5eactvw7wtpktv5****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The destination IP address.
	//
	// example:
	//
	// 172.50.XX.XX
	TargetIpAddress *string `json:"TargetIpAddress,omitempty" xml:"TargetIpAddress,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 80
	TargetPort *int32 `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	// The type of the destination resource. Valid values:
	//
	// 	- **ecs**: the ECS instance
	//
	// 	- **internetIp**: the public IP address
	//
	// 	- **vsw**: the vSwitch
	//
	// 	- **vpn**: the VPN gateway
	//
	// 	- **vbr**: the VBR
	//
	// 	- **clb**: the Classic Load Balancer (CLB) instance
	//
	// example:
	//
	// ecs
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateAndAnalyzeNetworkPathRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAndAnalyzeNetworkPathRequest) GoString() string {
	return s.String()
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetProtocol(v string) *CreateAndAnalyzeNetworkPathRequest {
	s.Protocol = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetRegionId(v string) *CreateAndAnalyzeNetworkPathRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetSourceId(v string) *CreateAndAnalyzeNetworkPathRequest {
	s.SourceId = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetSourceIpAddress(v string) *CreateAndAnalyzeNetworkPathRequest {
	s.SourceIpAddress = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetSourcePort(v int32) *CreateAndAnalyzeNetworkPathRequest {
	s.SourcePort = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetSourceType(v string) *CreateAndAnalyzeNetworkPathRequest {
	s.SourceType = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetTargetId(v string) *CreateAndAnalyzeNetworkPathRequest {
	s.TargetId = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetTargetIpAddress(v string) *CreateAndAnalyzeNetworkPathRequest {
	s.TargetIpAddress = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetTargetPort(v int32) *CreateAndAnalyzeNetworkPathRequest {
	s.TargetPort = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetTargetType(v string) *CreateAndAnalyzeNetworkPathRequest {
	s.TargetType = &v
	return s
}

type CreateAndAnalyzeNetworkPathResponseBody struct {
	// The ID of the task for analyzing network reachability that you initiated.
	//
	// example:
	//
	// nra-dfe9e53d2b524568****
	NetworkReachableAnalysisId *string `json:"NetworkReachableAnalysisId,omitempty" xml:"NetworkReachableAnalysisId,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D5E98683-355B-5867-8D3D-A24755F6895B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the source resource.
	//
	// example:
	//
	// i-uf62y8khhbkbdrp6****
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	// The source port.
	//
	// example:
	//
	// 0
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The type of the source resource.
	//
	// example:
	//
	// ecs
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The ID of the destination resource.
	//
	// example:
	//
	// i-m5eactvw7wtpktv5****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The destination IP address.
	//
	// example:
	//
	// 172.50.XX.XX
	TargetIpAddress *string `json:"TargetIpAddress,omitempty" xml:"TargetIpAddress,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 80
	TargetPort *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	// The type of the destination resource.
	//
	// example:
	//
	// ecs
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateAndAnalyzeNetworkPathResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAndAnalyzeNetworkPathResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetNetworkReachableAnalysisId(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.NetworkReachableAnalysisId = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetProtocol(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.Protocol = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetRequestId(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetSourceId(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetSourceIpAddress(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.SourceIpAddress = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetSourcePort(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.SourcePort = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetSourceType(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.SourceType = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetTargetId(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.TargetId = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetTargetIpAddress(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.TargetIpAddress = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetTargetPort(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.TargetPort = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponseBody) SetTargetType(v string) *CreateAndAnalyzeNetworkPathResponseBody {
	s.TargetType = &v
	return s
}

type CreateAndAnalyzeNetworkPathResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAndAnalyzeNetworkPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAndAnalyzeNetworkPathResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAndAnalyzeNetworkPathResponse) GoString() string {
	return s.String()
}

func (s *CreateAndAnalyzeNetworkPathResponse) SetHeaders(v map[string]*string) *CreateAndAnalyzeNetworkPathResponse {
	s.Headers = v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponse) SetStatusCode(v int32) *CreateAndAnalyzeNetworkPathResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponse) SetBody(v *CreateAndAnalyzeNetworkPathResponseBody) *CreateAndAnalyzeNetworkPathResponse {
	s.Body = v
	return s
}

type CreateNetworkPathRequest struct {
	// The description of the network path.
	//
	// example:
	//
	// Analyze the path from ECS to ECS
	NetworkPathDescription *string `json:"NetworkPathDescription,omitempty" xml:"NetworkPathDescription,omitempty"`
	// The name of the network path.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs2PublicIp
	NetworkPathName *string `json:"NetworkPathName,omitempty" xml:"NetworkPathName,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- **tcp**: Transmission Control Protocol (TCP)
	//
	// 	- **udp**: User Datagram Protocol (UDP)
	//
	// 	- **icmp**: Internet Control Message Protocol (ICMP)
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region ID of the network path that you want to create.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfm27qsxjj****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the source resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2zef4ngqfarepyun****
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 172.17.XX.XX
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	// The source port.
	//
	// example:
	//
	// 443
	SourcePort *int32 `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The type of the source resource. Valid values:
	//
	// 	- **ecs**: the Elastic Compute Service (ECS) instance
	//
	// 	- **internetIp**: the public IP address
	//
	// 	- **vsw**: the vSwitch
	//
	// 	- **vpn**: the VPN gateway
	//
	// 	- **vbr**: the virtual border router (VBR)
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The tags to add to the resource.
	Tag []*CreateNetworkPathRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the destination resource.
	//
	// example:
	//
	// i-bp13d0e064gubm****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The destination IP address.
	//
	// example:
	//
	// 192.168.0.210
	TargetIpAddress *string `json:"TargetIpAddress,omitempty" xml:"TargetIpAddress,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 80
	TargetPort *int32 `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	// The type of the destination resource. Valid values:
	//
	// 	- **ecs**: the ECS instance
	//
	// 	- **internetIp**: the public IP address
	//
	// 	- **vsw**: the vSwitch
	//
	// 	- **vpn**: the VPN gateway
	//
	// 	- **vbr**: the VBR
	//
	// 	- **clb**: the Classic Load Balancer (CLB) instance
	//
	// example:
	//
	// ecs
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateNetworkPathRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkPathRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkPathRequest) SetNetworkPathDescription(v string) *CreateNetworkPathRequest {
	s.NetworkPathDescription = &v
	return s
}

func (s *CreateNetworkPathRequest) SetNetworkPathName(v string) *CreateNetworkPathRequest {
	s.NetworkPathName = &v
	return s
}

func (s *CreateNetworkPathRequest) SetProtocol(v string) *CreateNetworkPathRequest {
	s.Protocol = &v
	return s
}

func (s *CreateNetworkPathRequest) SetRegionId(v string) *CreateNetworkPathRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNetworkPathRequest) SetResourceGroupId(v string) *CreateNetworkPathRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateNetworkPathRequest) SetSourceId(v string) *CreateNetworkPathRequest {
	s.SourceId = &v
	return s
}

func (s *CreateNetworkPathRequest) SetSourceIpAddress(v string) *CreateNetworkPathRequest {
	s.SourceIpAddress = &v
	return s
}

func (s *CreateNetworkPathRequest) SetSourcePort(v int32) *CreateNetworkPathRequest {
	s.SourcePort = &v
	return s
}

func (s *CreateNetworkPathRequest) SetSourceType(v string) *CreateNetworkPathRequest {
	s.SourceType = &v
	return s
}

func (s *CreateNetworkPathRequest) SetTag(v []*CreateNetworkPathRequestTag) *CreateNetworkPathRequest {
	s.Tag = v
	return s
}

func (s *CreateNetworkPathRequest) SetTargetId(v string) *CreateNetworkPathRequest {
	s.TargetId = &v
	return s
}

func (s *CreateNetworkPathRequest) SetTargetIpAddress(v string) *CreateNetworkPathRequest {
	s.TargetIpAddress = &v
	return s
}

func (s *CreateNetworkPathRequest) SetTargetPort(v int32) *CreateNetworkPathRequest {
	s.TargetPort = &v
	return s
}

func (s *CreateNetworkPathRequest) SetTargetType(v string) *CreateNetworkPathRequest {
	s.TargetType = &v
	return s
}

type CreateNetworkPathRequestTag struct {
	// The key of tag N to add to the resource. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// You can add up to 20 tags in each call.
	//
	// example:
	//
	// role
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// ops
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNetworkPathRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkPathRequestTag) GoString() string {
	return s.String()
}

func (s *CreateNetworkPathRequestTag) SetKey(v string) *CreateNetworkPathRequestTag {
	s.Key = &v
	return s
}

func (s *CreateNetworkPathRequestTag) SetValue(v string) *CreateNetworkPathRequestTag {
	s.Value = &v
	return s
}

type CreateNetworkPathResponseBody struct {
	// The ID of the network path.
	//
	// example:
	//
	// np-4cbf598673d14d27****
	NetworkPathId *string `json:"NetworkPathId,omitempty" xml:"NetworkPathId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 92DD9FFB-06FB-56F7-83EF-5CEF98F5562A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkPathResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkPathResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkPathResponseBody) SetNetworkPathId(v string) *CreateNetworkPathResponseBody {
	s.NetworkPathId = &v
	return s
}

func (s *CreateNetworkPathResponseBody) SetRequestId(v string) *CreateNetworkPathResponseBody {
	s.RequestId = &v
	return s
}

type CreateNetworkPathResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkPathResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkPathResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkPathResponse) SetHeaders(v map[string]*string) *CreateNetworkPathResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkPathResponse) SetStatusCode(v int32) *CreateNetworkPathResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkPathResponse) SetBody(v *CreateNetworkPathResponseBody) *CreateNetworkPathResponse {
	s.Body = v
	return s
}

type CreateNetworkReachableAnalysisRequest struct {
	// The ID of the network path. You can call the [CreateNetworkPath](https://help.aliyun.com/document_detail/2366522.html) operation to obtain the ID of the network path.
	//
	// This parameter is required.
	//
	// example:
	//
	// np-b2f618ceb2c84057****
	NetworkPathId *string `json:"NetworkPathId,omitempty" xml:"NetworkPathId,omitempty"`
	// The ID of the region for which you want to create a task for analyzing network reachability.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags to add to the resource.
	Tag []*CreateNetworkReachableAnalysisRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateNetworkReachableAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkReachableAnalysisRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkReachableAnalysisRequest) SetNetworkPathId(v string) *CreateNetworkReachableAnalysisRequest {
	s.NetworkPathId = &v
	return s
}

func (s *CreateNetworkReachableAnalysisRequest) SetRegionId(v string) *CreateNetworkReachableAnalysisRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNetworkReachableAnalysisRequest) SetTag(v []*CreateNetworkReachableAnalysisRequestTag) *CreateNetworkReachableAnalysisRequest {
	s.Tag = v
	return s
}

type CreateNetworkReachableAnalysisRequestTag struct {
	// The key of the tag to add to the resource. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// You can add up to 20 tags in each call.
	//
	// example:
	//
	// Team
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to add to the resource. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`. The tag value can be an empty string.
	//
	// You can add up to 20 tag values in each call.
	//
	// example:
	//
	// ops
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNetworkReachableAnalysisRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkReachableAnalysisRequestTag) GoString() string {
	return s.String()
}

func (s *CreateNetworkReachableAnalysisRequestTag) SetKey(v string) *CreateNetworkReachableAnalysisRequestTag {
	s.Key = &v
	return s
}

func (s *CreateNetworkReachableAnalysisRequestTag) SetValue(v string) *CreateNetworkReachableAnalysisRequestTag {
	s.Value = &v
	return s
}

type CreateNetworkReachableAnalysisResponseBody struct {
	// The ID of the task for analyzing network reachability.
	//
	// example:
	//
	// nra-2fede05617494417****
	NetworkReachableAnalysisId *string `json:"NetworkReachableAnalysisId,omitempty" xml:"NetworkReachableAnalysisId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A7F0D6EC-E19E-58AC-AC9F-08036763960F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkReachableAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkReachableAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkReachableAnalysisResponseBody) SetNetworkReachableAnalysisId(v string) *CreateNetworkReachableAnalysisResponseBody {
	s.NetworkReachableAnalysisId = &v
	return s
}

func (s *CreateNetworkReachableAnalysisResponseBody) SetRequestId(v string) *CreateNetworkReachableAnalysisResponseBody {
	s.RequestId = &v
	return s
}

type CreateNetworkReachableAnalysisResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkReachableAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkReachableAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkReachableAnalysisResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkReachableAnalysisResponse) SetHeaders(v map[string]*string) *CreateNetworkReachableAnalysisResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkReachableAnalysisResponse) SetStatusCode(v int32) *CreateNetworkReachableAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkReachableAnalysisResponse) SetBody(v *CreateNetworkReachableAnalysisResponseBody) *CreateNetworkReachableAnalysisResponse {
	s.Body = v
	return s
}

type DeleteNetworkPathRequest struct {
	// The IDs of network paths.
	//
	// This parameter is required.
	NetworkPathIds []*string `json:"NetworkPathIds,omitempty" xml:"NetworkPathIds,omitempty" type:"Repeated"`
	// The region ID of the network path that you want to delete.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteNetworkPathRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkPathRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPathRequest) SetNetworkPathIds(v []*string) *DeleteNetworkPathRequest {
	s.NetworkPathIds = v
	return s
}

func (s *DeleteNetworkPathRequest) SetRegionId(v string) *DeleteNetworkPathRequest {
	s.RegionId = &v
	return s
}

type DeleteNetworkPathShrinkRequest struct {
	// The IDs of network paths.
	//
	// This parameter is required.
	NetworkPathIdsShrink *string `json:"NetworkPathIds,omitempty" xml:"NetworkPathIds,omitempty"`
	// The region ID of the network path that you want to delete.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteNetworkPathShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkPathShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPathShrinkRequest) SetNetworkPathIdsShrink(v string) *DeleteNetworkPathShrinkRequest {
	s.NetworkPathIdsShrink = &v
	return s
}

func (s *DeleteNetworkPathShrinkRequest) SetRegionId(v string) *DeleteNetworkPathShrinkRequest {
	s.RegionId = &v
	return s
}

type DeleteNetworkPathResponseBody struct {
	// Result of operation.
	//
	// - **true**: Delete Success.
	//
	// - **false**: Delete Fail.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C4331873-C534-590F-A905-F66C53B88A47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkPathResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkPathResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPathResponseBody) SetData(v bool) *DeleteNetworkPathResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteNetworkPathResponseBody) SetRequestId(v string) *DeleteNetworkPathResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNetworkPathResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkPathResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkPathResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPathResponse) SetHeaders(v map[string]*string) *DeleteNetworkPathResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkPathResponse) SetStatusCode(v int32) *DeleteNetworkPathResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkPathResponse) SetBody(v *DeleteNetworkPathResponseBody) *DeleteNetworkPathResponse {
	s.Body = v
	return s
}

type DeleteNetworkReachableAnalysisRequest struct {
	// The IDs of the tasks for analyzing network reachability.
	//
	// This parameter is required.
	NetworkReachableAnalysisIds []*string `json:"NetworkReachableAnalysisIds,omitempty" xml:"NetworkReachableAnalysisIds,omitempty" type:"Repeated"`
	// The ID of the region for which you want to delete a task for analyzing network reachability.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteNetworkReachableAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkReachableAnalysisRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkReachableAnalysisRequest) SetNetworkReachableAnalysisIds(v []*string) *DeleteNetworkReachableAnalysisRequest {
	s.NetworkReachableAnalysisIds = v
	return s
}

func (s *DeleteNetworkReachableAnalysisRequest) SetRegionId(v string) *DeleteNetworkReachableAnalysisRequest {
	s.RegionId = &v
	return s
}

type DeleteNetworkReachableAnalysisShrinkRequest struct {
	// The IDs of the tasks for analyzing network reachability.
	//
	// This parameter is required.
	NetworkReachableAnalysisIdsShrink *string `json:"NetworkReachableAnalysisIds,omitempty" xml:"NetworkReachableAnalysisIds,omitempty"`
	// The ID of the region for which you want to delete a task for analyzing network reachability.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteNetworkReachableAnalysisShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkReachableAnalysisShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkReachableAnalysisShrinkRequest) SetNetworkReachableAnalysisIdsShrink(v string) *DeleteNetworkReachableAnalysisShrinkRequest {
	s.NetworkReachableAnalysisIdsShrink = &v
	return s
}

func (s *DeleteNetworkReachableAnalysisShrinkRequest) SetRegionId(v string) *DeleteNetworkReachableAnalysisShrinkRequest {
	s.RegionId = &v
	return s
}

type DeleteNetworkReachableAnalysisResponseBody struct {
	// Result of operation.
	//
	// - **true**: Delete Success.
	//
	// - **false**: Delete Fail.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4838F3F2-30E1-5D82-B25A-B9FE33BC3E25
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkReachableAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkReachableAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkReachableAnalysisResponseBody) SetData(v bool) *DeleteNetworkReachableAnalysisResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteNetworkReachableAnalysisResponseBody) SetRequestId(v string) *DeleteNetworkReachableAnalysisResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNetworkReachableAnalysisResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkReachableAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkReachableAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkReachableAnalysisResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkReachableAnalysisResponse) SetHeaders(v map[string]*string) *DeleteNetworkReachableAnalysisResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkReachableAnalysisResponse) SetStatusCode(v int32) *DeleteNetworkReachableAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkReachableAnalysisResponse) SetBody(v *DeleteNetworkReachableAnalysisResponseBody) *DeleteNetworkReachableAnalysisResponse {
	s.Body = v
	return s
}

type GetInternetTupleRequest struct {
	// The IDs of member accounts.
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684373600099
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The local IP address.
	//
	// example:
	//
	// 112.74.XX.XX
	CloudIp *string `json:"CloudIp,omitempty" xml:"CloudIp,omitempty"`
	// The local IP addresses for filtering.
	CloudIpList []*string `json:"CloudIpList,omitempty" xml:"CloudIpList,omitempty" type:"Repeated"`
	// The local Internet service provider (ISP).
	//
	// >  In most cases, the value is Alibaba or Alibaba Cloud.
	//
	// example:
	//
	// Alibaba
	CloudIsp *string `json:"CloudIsp,omitempty" xml:"CloudIsp,omitempty"`
	// The local port.
	//
	// >  This parameter is required only if you set GroupBy to CloudPort.
	//
	// example:
	//
	// 443
	CloudPort *string `json:"CloudPort,omitempty" xml:"CloudPort,omitempty"`
	// The direction of the Internet traffic that you want to query. Valid values:
	//
	// 	- **in**: inbound
	//
	// 	- **out**: outbound
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684373700099
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Alibaba Cloud instance.
	//
	// example:
	//
	// eip-sample*
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance IDs for filtering.
	InstanceList []*string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The metric for data ranking. Default value: **ByteCount**. This value indicates that Internet traffic data is ranked by traffic volume.
	//
	// Valid values:
	//
	// 	- Rtt
	//
	// 	- ByteCount
	//
	// 	- PacketCount
	//
	// 	- RetransmitRate
	//
	// example:
	//
	// ByteCount
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The remote city.
	//
	// >  This parameter is required only if you set **TupleType*	- to **2*	- or **5**.
	//
	// example:
	//
	// Hangzhou
	OtherCity *string `json:"OtherCity,omitempty" xml:"OtherCity,omitempty"`
	// The remote country.
	//
	// >  This parameter is required only if you set **TupleType*	- to **2*	- or **5**.
	//
	// example:
	//
	// China
	OtherCountry *string `json:"OtherCountry,omitempty" xml:"OtherCountry,omitempty"`
	// The remote IP address.
	//
	// > This parameter is required only when you set **TupleType** to **2** or **5**.
	//
	// example:
	//
	// 122.112.XX.XX
	OtherIp *string `json:"OtherIp,omitempty" xml:"OtherIp,omitempty"`
	// The remote ISP.
	//
	// > This parameter is required if you want to view the information about the remote ISP.
	//
	// example:
	//
	// China Mobile
	OtherIsp *string `json:"OtherIsp,omitempty" xml:"OtherIsp,omitempty"`
	// The remote port.
	//
	// > This parameter is required only when you set **TupleType*	- to **5**.
	//
	// example:
	//
	// 40002
	OtherPort *string `json:"OtherPort,omitempty" xml:"OtherPort,omitempty"`
	// The protocol number.
	//
	// > All protocols are supported. This parameter is required only when you set **TupleType** to **5**.
	//
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the region for which you want to query the Internet traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The order in which instances are ranked by Internet traffic. Valid values:
	//
	// 	- **desc**: the descending order
	//
	// 	- **asc**: the ascending order
	//
	// example:
	//
	// desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Specifies top-N traffic data to display. Default value: **10**. This value specifies to display top-10 traffic data by default. Max value: **100**.
	//
	// example:
	//
	// 10
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
	// The type of the tuple. Valid values:
	//
	// 	- **1**: 1-tuple
	//
	// 	- **2**: 2-tuple
	//
	// 	- **5**: 5-tuple
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TupleType *int32 `json:"TupleType,omitempty" xml:"TupleType,omitempty"`
	// Specifies whether to enable the multi-account management feature. Default value: **false**. This value specifies that the multi-account management feature is disabled.
	//
	// >  By default, the multi-account management feature is not available. If you want to use this feature, contact your account manager to apply for permissions.
	//
	// example:
	//
	// false
	UseMultiAccount *bool `json:"UseMultiAccount,omitempty" xml:"UseMultiAccount,omitempty"`
}

func (s GetInternetTupleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInternetTupleRequest) GoString() string {
	return s.String()
}

func (s *GetInternetTupleRequest) SetAccountIds(v []*int64) *GetInternetTupleRequest {
	s.AccountIds = v
	return s
}

func (s *GetInternetTupleRequest) SetBeginTime(v int64) *GetInternetTupleRequest {
	s.BeginTime = &v
	return s
}

func (s *GetInternetTupleRequest) SetCloudIp(v string) *GetInternetTupleRequest {
	s.CloudIp = &v
	return s
}

func (s *GetInternetTupleRequest) SetCloudIpList(v []*string) *GetInternetTupleRequest {
	s.CloudIpList = v
	return s
}

func (s *GetInternetTupleRequest) SetCloudIsp(v string) *GetInternetTupleRequest {
	s.CloudIsp = &v
	return s
}

func (s *GetInternetTupleRequest) SetCloudPort(v string) *GetInternetTupleRequest {
	s.CloudPort = &v
	return s
}

func (s *GetInternetTupleRequest) SetDirection(v string) *GetInternetTupleRequest {
	s.Direction = &v
	return s
}

func (s *GetInternetTupleRequest) SetEndTime(v int64) *GetInternetTupleRequest {
	s.EndTime = &v
	return s
}

func (s *GetInternetTupleRequest) SetInstanceId(v string) *GetInternetTupleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInternetTupleRequest) SetInstanceList(v []*string) *GetInternetTupleRequest {
	s.InstanceList = v
	return s
}

func (s *GetInternetTupleRequest) SetOrderBy(v string) *GetInternetTupleRequest {
	s.OrderBy = &v
	return s
}

func (s *GetInternetTupleRequest) SetOtherCity(v string) *GetInternetTupleRequest {
	s.OtherCity = &v
	return s
}

func (s *GetInternetTupleRequest) SetOtherCountry(v string) *GetInternetTupleRequest {
	s.OtherCountry = &v
	return s
}

func (s *GetInternetTupleRequest) SetOtherIp(v string) *GetInternetTupleRequest {
	s.OtherIp = &v
	return s
}

func (s *GetInternetTupleRequest) SetOtherIsp(v string) *GetInternetTupleRequest {
	s.OtherIsp = &v
	return s
}

func (s *GetInternetTupleRequest) SetOtherPort(v string) *GetInternetTupleRequest {
	s.OtherPort = &v
	return s
}

func (s *GetInternetTupleRequest) SetProtocol(v string) *GetInternetTupleRequest {
	s.Protocol = &v
	return s
}

func (s *GetInternetTupleRequest) SetRegionId(v string) *GetInternetTupleRequest {
	s.RegionId = &v
	return s
}

func (s *GetInternetTupleRequest) SetSort(v string) *GetInternetTupleRequest {
	s.Sort = &v
	return s
}

func (s *GetInternetTupleRequest) SetTopN(v int32) *GetInternetTupleRequest {
	s.TopN = &v
	return s
}

func (s *GetInternetTupleRequest) SetTupleType(v int32) *GetInternetTupleRequest {
	s.TupleType = &v
	return s
}

func (s *GetInternetTupleRequest) SetUseMultiAccount(v bool) *GetInternetTupleRequest {
	s.UseMultiAccount = &v
	return s
}

type GetInternetTupleShrinkRequest struct {
	// The IDs of member accounts.
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684373600099
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The local IP address.
	//
	// example:
	//
	// 112.74.XX.XX
	CloudIp *string `json:"CloudIp,omitempty" xml:"CloudIp,omitempty"`
	// The local IP addresses for filtering.
	CloudIpListShrink *string `json:"CloudIpList,omitempty" xml:"CloudIpList,omitempty"`
	// The local Internet service provider (ISP).
	//
	// >  In most cases, the value is Alibaba or Alibaba Cloud.
	//
	// example:
	//
	// Alibaba
	CloudIsp *string `json:"CloudIsp,omitempty" xml:"CloudIsp,omitempty"`
	// The local port.
	//
	// >  This parameter is required only if you set GroupBy to CloudPort.
	//
	// example:
	//
	// 443
	CloudPort *string `json:"CloudPort,omitempty" xml:"CloudPort,omitempty"`
	// The direction of the Internet traffic that you want to query. Valid values:
	//
	// 	- **in**: inbound
	//
	// 	- **out**: outbound
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684373700099
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Alibaba Cloud instance.
	//
	// example:
	//
	// eip-sample*
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance IDs for filtering.
	InstanceListShrink *string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty"`
	// The metric for data ranking. Default value: **ByteCount**. This value indicates that Internet traffic data is ranked by traffic volume.
	//
	// Valid values:
	//
	// 	- Rtt
	//
	// 	- ByteCount
	//
	// 	- PacketCount
	//
	// 	- RetransmitRate
	//
	// example:
	//
	// ByteCount
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The remote city.
	//
	// >  This parameter is required only if you set **TupleType*	- to **2*	- or **5**.
	//
	// example:
	//
	// Hangzhou
	OtherCity *string `json:"OtherCity,omitempty" xml:"OtherCity,omitempty"`
	// The remote country.
	//
	// >  This parameter is required only if you set **TupleType*	- to **2*	- or **5**.
	//
	// example:
	//
	// China
	OtherCountry *string `json:"OtherCountry,omitempty" xml:"OtherCountry,omitempty"`
	// The remote IP address.
	//
	// > This parameter is required only when you set **TupleType** to **2** or **5**.
	//
	// example:
	//
	// 122.112.XX.XX
	OtherIp *string `json:"OtherIp,omitempty" xml:"OtherIp,omitempty"`
	// The remote ISP.
	//
	// > This parameter is required if you want to view the information about the remote ISP.
	//
	// example:
	//
	// China Mobile
	OtherIsp *string `json:"OtherIsp,omitempty" xml:"OtherIsp,omitempty"`
	// The remote port.
	//
	// > This parameter is required only when you set **TupleType*	- to **5**.
	//
	// example:
	//
	// 40002
	OtherPort *string `json:"OtherPort,omitempty" xml:"OtherPort,omitempty"`
	// The protocol number.
	//
	// > All protocols are supported. This parameter is required only when you set **TupleType** to **5**.
	//
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the region for which you want to query the Internet traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The order in which instances are ranked by Internet traffic. Valid values:
	//
	// 	- **desc**: the descending order
	//
	// 	- **asc**: the ascending order
	//
	// example:
	//
	// desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Specifies top-N traffic data to display. Default value: **10**. This value specifies to display top-10 traffic data by default. Max value: **100**.
	//
	// example:
	//
	// 10
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
	// The type of the tuple. Valid values:
	//
	// 	- **1**: 1-tuple
	//
	// 	- **2**: 2-tuple
	//
	// 	- **5**: 5-tuple
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TupleType *int32 `json:"TupleType,omitempty" xml:"TupleType,omitempty"`
	// Specifies whether to enable the multi-account management feature. Default value: **false**. This value specifies that the multi-account management feature is disabled.
	//
	// >  By default, the multi-account management feature is not available. If you want to use this feature, contact your account manager to apply for permissions.
	//
	// example:
	//
	// false
	UseMultiAccount *bool `json:"UseMultiAccount,omitempty" xml:"UseMultiAccount,omitempty"`
}

func (s GetInternetTupleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInternetTupleShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetInternetTupleShrinkRequest) SetAccountIds(v []*int64) *GetInternetTupleShrinkRequest {
	s.AccountIds = v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetBeginTime(v int64) *GetInternetTupleShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetCloudIp(v string) *GetInternetTupleShrinkRequest {
	s.CloudIp = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetCloudIpListShrink(v string) *GetInternetTupleShrinkRequest {
	s.CloudIpListShrink = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetCloudIsp(v string) *GetInternetTupleShrinkRequest {
	s.CloudIsp = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetCloudPort(v string) *GetInternetTupleShrinkRequest {
	s.CloudPort = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetDirection(v string) *GetInternetTupleShrinkRequest {
	s.Direction = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetEndTime(v int64) *GetInternetTupleShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetInstanceId(v string) *GetInternetTupleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetInstanceListShrink(v string) *GetInternetTupleShrinkRequest {
	s.InstanceListShrink = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetOrderBy(v string) *GetInternetTupleShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetOtherCity(v string) *GetInternetTupleShrinkRequest {
	s.OtherCity = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetOtherCountry(v string) *GetInternetTupleShrinkRequest {
	s.OtherCountry = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetOtherIp(v string) *GetInternetTupleShrinkRequest {
	s.OtherIp = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetOtherIsp(v string) *GetInternetTupleShrinkRequest {
	s.OtherIsp = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetOtherPort(v string) *GetInternetTupleShrinkRequest {
	s.OtherPort = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetProtocol(v string) *GetInternetTupleShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetRegionId(v string) *GetInternetTupleShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetSort(v string) *GetInternetTupleShrinkRequest {
	s.Sort = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetTopN(v int32) *GetInternetTupleShrinkRequest {
	s.TopN = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetTupleType(v int32) *GetInternetTupleShrinkRequest {
	s.TupleType = &v
	return s
}

func (s *GetInternetTupleShrinkRequest) SetUseMultiAccount(v bool) *GetInternetTupleShrinkRequest {
	s.UseMultiAccount = &v
	return s
}

type GetInternetTupleResponseBody struct {
	// The ranking result of Internet traffic data.
	Data []*GetInternetTupleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D5E98683-355B-5867-8D3D-A24755F6895B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInternetTupleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInternetTupleResponseBody) GoString() string {
	return s.String()
}

func (s *GetInternetTupleResponseBody) SetData(v []*GetInternetTupleResponseBodyData) *GetInternetTupleResponseBody {
	s.Data = v
	return s
}

func (s *GetInternetTupleResponseBody) SetRequestId(v string) *GetInternetTupleResponseBody {
	s.RequestId = &v
	return s
}

type GetInternetTupleResponseBodyData struct {
	// The access point of Alibaba Cloud.
	//
	// >  This parameter is valid only if you set **InstanceId*	- to the instance ID of an Anycast elastic IP address (EIP).
	//
	// example:
	//
	// cn-hongkong-pop
	AccessRegion *string `json:"AccessRegion,omitempty" xml:"AccessRegion,omitempty"`
	// The beginning of the time range that you queried. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1684373600099
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The traffic volume. Unit: bytes.
	//
	// example:
	//
	// 88
	ByteCount *float64 `json:"ByteCount,omitempty" xml:"ByteCount,omitempty"`
	// The local city.
	//
	// example:
	//
	// Nanjing
	CloudCity *string `json:"CloudCity,omitempty" xml:"CloudCity,omitempty"`
	// The local country or region.
	//
	// example:
	//
	// China
	CloudCountry *string `json:"CloudCountry,omitempty" xml:"CloudCountry,omitempty"`
	// The local IP address.
	//
	// example:
	//
	// 112.74.XX.XX
	CloudIp *string `json:"CloudIp,omitempty" xml:"CloudIp,omitempty"`
	// The local ISP.
	//
	// example:
	//
	// China Mobile
	CloudIsp *string `json:"CloudIsp,omitempty" xml:"CloudIsp,omitempty"`
	// The local port.
	//
	// example:
	//
	// 443
	CloudPort *string `json:"CloudPort,omitempty" xml:"CloudPort,omitempty"`
	// The service code of the instance to which the local IP address belongs.
	//
	// example:
	//
	// EIP
	CloudProduct *string `json:"CloudProduct,omitempty" xml:"CloudProduct,omitempty"`
	// The local province.
	//
	// example:
	//
	// Jiangsu
	CloudProvince *string `json:"CloudProvince,omitempty" xml:"CloudProvince,omitempty"`
	// The direction of Internet traffic. Valid values:
	//
	// 	- **in**: inbound
	//
	// 	- **out**: outbound
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The inbound traffic volume. Unit: bytes.
	//
	// example:
	//
	// 88
	InByteCount *float64 `json:"InByteCount,omitempty" xml:"InByteCount,omitempty"`
	// The number of inbound disordered packets.
	//
	// example:
	//
	// 2
	InOutOrderCount *float64 `json:"InOutOrderCount,omitempty" xml:"InOutOrderCount,omitempty"`
	// The number of inbound packets.
	//
	// example:
	//
	// 33
	InPacketCount *float64 `json:"InPacketCount,omitempty" xml:"InPacketCount,omitempty"`
	// The number of inbound repeated packets.
	//
	// example:
	//
	// 0
	InRetranCount *float64 `json:"InRetranCount,omitempty" xml:"InRetranCount,omitempty"`
	// The ID of the instance to which the local IP address belongs.
	//
	// example:
	//
	// eip-sample*
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The remote city. In most cases, this parameter is empty if you set **OtherCountry*	- to a country except China.
	//
	// example:
	//
	// Austin
	OtherCity *string `json:"OtherCity,omitempty" xml:"OtherCity,omitempty"`
	// The remote country or region.
	//
	// example:
	//
	// United States
	OtherCountry *string `json:"OtherCountry,omitempty" xml:"OtherCountry,omitempty"`
	// The remote IP address.
	//
	// example:
	//
	// 122.112.XX.XX
	OtherIp *string `json:"OtherIp,omitempty" xml:"OtherIp,omitempty"`
	// The remote ISP.
	//
	// example:
	//
	// amazon.com
	OtherIsp *string `json:"OtherIsp,omitempty" xml:"OtherIsp,omitempty"`
	// The remote port.
	//
	// example:
	//
	// 40002
	OtherPort *string `json:"OtherPort,omitempty" xml:"OtherPort,omitempty"`
	// The service code of the instance to which the remote IP address belongs. If the IP address is not on the cloud, this parameter is empty.
	//
	// example:
	//
	// ECS
	OtherProduct *string `json:"OtherProduct,omitempty" xml:"OtherProduct,omitempty"`
	// The remote province. In most cases, this parameter is empty if you set **OtherCountry*	- to a country except China.
	//
	// example:
	//
	// Texas
	OtherProvince *string `json:"OtherProvince,omitempty" xml:"OtherProvince,omitempty"`
	// The outbound traffic volume. Unit: bytes.
	//
	// example:
	//
	// 66
	OutByteCount *float64 `json:"OutByteCount,omitempty" xml:"OutByteCount,omitempty"`
	// The number of disordered packets.
	//
	// example:
	//
	// 1
	OutOrderCount *float64 `json:"OutOrderCount,omitempty" xml:"OutOrderCount,omitempty"`
	// The number of outbound disordered packets.
	//
	// example:
	//
	// 1
	OutOutOrderCount *float64 `json:"OutOutOrderCount,omitempty" xml:"OutOutOrderCount,omitempty"`
	// The number of outbound packets.
	//
	// example:
	//
	// 22
	OutPacketCount *float64 `json:"OutPacketCount,omitempty" xml:"OutPacketCount,omitempty"`
	// The number of outbound repeated packets.
	//
	// example:
	//
	// 1
	OutRetranCount *float64 `json:"OutRetranCount,omitempty" xml:"OutRetranCount,omitempty"`
	// The number of packets.
	//
	// example:
	//
	// 66
	PacketCount *float64 `json:"PacketCount,omitempty" xml:"PacketCount,omitempty"`
	// The protocol number.
	//
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The retransmission rate of TCP packets.
	//
	// example:
	//
	// 0.1
	RetransmitRate *float64 `json:"RetransmitRate,omitempty" xml:"RetransmitRate,omitempty"`
	// The round-trip time (RTT). Unit: milliseconds.
	//
	// example:
	//
	// 10000
	Rtt *float64 `json:"Rtt,omitempty" xml:"Rtt,omitempty"`
}

func (s GetInternetTupleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInternetTupleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInternetTupleResponseBodyData) SetAccessRegion(v string) *GetInternetTupleResponseBodyData {
	s.AccessRegion = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetBeginTime(v string) *GetInternetTupleResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetByteCount(v float64) *GetInternetTupleResponseBodyData {
	s.ByteCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetCloudCity(v string) *GetInternetTupleResponseBodyData {
	s.CloudCity = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetCloudCountry(v string) *GetInternetTupleResponseBodyData {
	s.CloudCountry = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetCloudIp(v string) *GetInternetTupleResponseBodyData {
	s.CloudIp = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetCloudIsp(v string) *GetInternetTupleResponseBodyData {
	s.CloudIsp = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetCloudPort(v string) *GetInternetTupleResponseBodyData {
	s.CloudPort = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetCloudProduct(v string) *GetInternetTupleResponseBodyData {
	s.CloudProduct = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetCloudProvince(v string) *GetInternetTupleResponseBodyData {
	s.CloudProvince = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetDirection(v string) *GetInternetTupleResponseBodyData {
	s.Direction = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetInByteCount(v float64) *GetInternetTupleResponseBodyData {
	s.InByteCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetInOutOrderCount(v float64) *GetInternetTupleResponseBodyData {
	s.InOutOrderCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetInPacketCount(v float64) *GetInternetTupleResponseBodyData {
	s.InPacketCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetInRetranCount(v float64) *GetInternetTupleResponseBodyData {
	s.InRetranCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetInstanceId(v string) *GetInternetTupleResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOtherCity(v string) *GetInternetTupleResponseBodyData {
	s.OtherCity = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOtherCountry(v string) *GetInternetTupleResponseBodyData {
	s.OtherCountry = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOtherIp(v string) *GetInternetTupleResponseBodyData {
	s.OtherIp = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOtherIsp(v string) *GetInternetTupleResponseBodyData {
	s.OtherIsp = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOtherPort(v string) *GetInternetTupleResponseBodyData {
	s.OtherPort = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOtherProduct(v string) *GetInternetTupleResponseBodyData {
	s.OtherProduct = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOtherProvince(v string) *GetInternetTupleResponseBodyData {
	s.OtherProvince = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOutByteCount(v float64) *GetInternetTupleResponseBodyData {
	s.OutByteCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOutOrderCount(v float64) *GetInternetTupleResponseBodyData {
	s.OutOrderCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOutOutOrderCount(v float64) *GetInternetTupleResponseBodyData {
	s.OutOutOrderCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOutPacketCount(v float64) *GetInternetTupleResponseBodyData {
	s.OutPacketCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetOutRetranCount(v float64) *GetInternetTupleResponseBodyData {
	s.OutRetranCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetPacketCount(v float64) *GetInternetTupleResponseBodyData {
	s.PacketCount = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetProtocol(v string) *GetInternetTupleResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetRetransmitRate(v float64) *GetInternetTupleResponseBodyData {
	s.RetransmitRate = &v
	return s
}

func (s *GetInternetTupleResponseBodyData) SetRtt(v float64) *GetInternetTupleResponseBodyData {
	s.Rtt = &v
	return s
}

type GetInternetTupleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInternetTupleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInternetTupleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInternetTupleResponse) GoString() string {
	return s.String()
}

func (s *GetInternetTupleResponse) SetHeaders(v map[string]*string) *GetInternetTupleResponse {
	s.Headers = v
	return s
}

func (s *GetInternetTupleResponse) SetStatusCode(v int32) *GetInternetTupleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInternetTupleResponse) SetBody(v *GetInternetTupleResponseBody) *GetInternetTupleResponse {
	s.Body = v
	return s
}

type GetNatTopNRequest struct {
	// The beginning of the time range to query in milliseconds. If you do not specify **EndTime**, the point in time specified by **BeginTime*	- is queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638239092000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end of the time range to query in milliseconds. The time range specified by **BeginTime*	- and **EndTime*	- cannot exceed **86400000*	- milliseconds (24 hours).
	//
	// example:
	//
	// 1638239093000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Query ranking statistics for a specific IP address. If you specify this parameter, you do not need to specify **TopN*	- or **OrderBy**.
	//
	// example:
	//
	// 192.168.156.101
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-sample***
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The metric that is used for real-time SNAT performance ranking. Valid values:
	//
	// 	- **InBps**: inbound data transfer. Unit: bit/s.
	//
	// 	- **OutBps**: outbound data transfer. Unit: bit/s.
	//
	// 	- **InPps**: inbound packet forwarding rate. Unit: packets per second.
	//
	// 	- **OutPps**: outbound packet forwarding rate. Unit: packets per second.
	//
	// 	- **NewSessionPerSecond**: new connection creation rate. Unit: connections per second.
	//
	// 	- **ActiveSessionCount**: number of concurrent connections. Unit: connections.
	//
	// example:
	//
	// InBps
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The ID of the region in which the NAT gateway is deployed.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of entries to return for real-time SNAT performance ranking. Valid values: **1 to 100**. Default value: **10**.
	//
	// example:
	//
	// 10
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
}

func (s GetNatTopNRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNatTopNRequest) GoString() string {
	return s.String()
}

func (s *GetNatTopNRequest) SetBeginTime(v int64) *GetNatTopNRequest {
	s.BeginTime = &v
	return s
}

func (s *GetNatTopNRequest) SetEndTime(v int64) *GetNatTopNRequest {
	s.EndTime = &v
	return s
}

func (s *GetNatTopNRequest) SetIp(v string) *GetNatTopNRequest {
	s.Ip = &v
	return s
}

func (s *GetNatTopNRequest) SetNatGatewayId(v string) *GetNatTopNRequest {
	s.NatGatewayId = &v
	return s
}

func (s *GetNatTopNRequest) SetOrderBy(v string) *GetNatTopNRequest {
	s.OrderBy = &v
	return s
}

func (s *GetNatTopNRequest) SetRegionId(v string) *GetNatTopNRequest {
	s.RegionId = &v
	return s
}

func (s *GetNatTopNRequest) SetTopN(v int32) *GetNatTopNRequest {
	s.TopN = &v
	return s
}

type GetNatTopNResponseBody struct {
	// Indicates whether Network Intelligence Service (NIS) is activated. The NatGatewayTopN parameter returns an empty array when NIS is not activated.
	//
	// 	- **true**: activated
	//
	// 	- **false**: not activated
	//
	// example:
	//
	// true
	IsTopNOpen *bool `json:"IsTopNOpen,omitempty" xml:"IsTopNOpen,omitempty"`
	// An array of statistics about real-time SNAT performance ranking.
	NatGatewayTopN []*GetNatTopNResponseBodyNatGatewayTopN `json:"NatGatewayTopN,omitempty" xml:"NatGatewayTopN,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 77C512B5-12f3-f892-BD94-88A98271C1A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNatTopNResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNatTopNResponseBody) GoString() string {
	return s.String()
}

func (s *GetNatTopNResponseBody) SetIsTopNOpen(v bool) *GetNatTopNResponseBody {
	s.IsTopNOpen = &v
	return s
}

func (s *GetNatTopNResponseBody) SetNatGatewayTopN(v []*GetNatTopNResponseBodyNatGatewayTopN) *GetNatTopNResponseBody {
	s.NatGatewayTopN = v
	return s
}

func (s *GetNatTopNResponseBody) SetRequestId(v string) *GetNatTopNResponseBody {
	s.RequestId = &v
	return s
}

type GetNatTopNResponseBodyNatGatewayTopN struct {
	// The number of concurrent connections. Unit: connections.
	//
	// example:
	//
	// 8
	ActiveSessionCount *float32 `json:"ActiveSessionCount,omitempty" xml:"ActiveSessionCount,omitempty"`
	// The inbound data transfer. Unit: bit/s.
	//
	// example:
	//
	// 100
	InBps *float32 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	// This field is reserved and not in use.
	//
	// example:
	//
	// 10
	InFlowPerMinute *float32 `json:"InFlowPerMinute,omitempty" xml:"InFlowPerMinute,omitempty"`
	// The inbound packet forwarding rate. Unit: packets per second.
	//
	// example:
	//
	// 10
	InPps *float32 `json:"InPps,omitempty" xml:"InPps,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 192.168.156.101
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The new connection creation rate. Unit: connections per second.
	//
	// example:
	//
	// 2
	NewSessionPerSecond *float32 `json:"NewSessionPerSecond,omitempty" xml:"NewSessionPerSecond,omitempty"`
	// The outbound data transfer. Unit: bit/s.
	//
	// example:
	//
	// 200
	OutBps *float32 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
	// This field is reserved and not in use.
	//
	// example:
	//
	// 10
	OutFlowPerMinute *float32 `json:"OutFlowPerMinute,omitempty" xml:"OutFlowPerMinute,omitempty"`
	// The outbound packet forwarding rate. Unit: packets per second.
	//
	// example:
	//
	// 20
	OutPps *float32 `json:"OutPps,omitempty" xml:"OutPps,omitempty"`
}

func (s GetNatTopNResponseBodyNatGatewayTopN) String() string {
	return tea.Prettify(s)
}

func (s GetNatTopNResponseBodyNatGatewayTopN) GoString() string {
	return s.String()
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetActiveSessionCount(v float32) *GetNatTopNResponseBodyNatGatewayTopN {
	s.ActiveSessionCount = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetInBps(v float32) *GetNatTopNResponseBodyNatGatewayTopN {
	s.InBps = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetInFlowPerMinute(v float32) *GetNatTopNResponseBodyNatGatewayTopN {
	s.InFlowPerMinute = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetInPps(v float32) *GetNatTopNResponseBodyNatGatewayTopN {
	s.InPps = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetIp(v string) *GetNatTopNResponseBodyNatGatewayTopN {
	s.Ip = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetNewSessionPerSecond(v float32) *GetNatTopNResponseBodyNatGatewayTopN {
	s.NewSessionPerSecond = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetOutBps(v float32) *GetNatTopNResponseBodyNatGatewayTopN {
	s.OutBps = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetOutFlowPerMinute(v float32) *GetNatTopNResponseBodyNatGatewayTopN {
	s.OutFlowPerMinute = &v
	return s
}

func (s *GetNatTopNResponseBodyNatGatewayTopN) SetOutPps(v float32) *GetNatTopNResponseBodyNatGatewayTopN {
	s.OutPps = &v
	return s
}

type GetNatTopNResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNatTopNResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNatTopNResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNatTopNResponse) GoString() string {
	return s.String()
}

func (s *GetNatTopNResponse) SetHeaders(v map[string]*string) *GetNatTopNResponse {
	s.Headers = v
	return s
}

func (s *GetNatTopNResponse) SetStatusCode(v int32) *GetNatTopNResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNatTopNResponse) SetBody(v *GetNatTopNResponseBody) *GetNatTopNResponse {
	s.Body = v
	return s
}

type GetNetworkReachableAnalysisRequest struct {
	// The ID of the task for analyzing network reachability. You can call the **CreateNetworkRearchableAnalysis*	- operation to obtain the ID of the task for analyzing network reachability.
	//
	// This parameter is required.
	//
	// example:
	//
	// nra-90eef36a9e6e4662****
	NetworkReachableAnalysisId *string `json:"NetworkReachableAnalysisId,omitempty" xml:"NetworkReachableAnalysisId,omitempty"`
	// The ID of the region for which you want to obtain the result of network reachability analysis.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetNetworkReachableAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkReachableAnalysisRequest) GoString() string {
	return s.String()
}

func (s *GetNetworkReachableAnalysisRequest) SetNetworkReachableAnalysisId(v string) *GetNetworkReachableAnalysisRequest {
	s.NetworkReachableAnalysisId = &v
	return s
}

func (s *GetNetworkReachableAnalysisRequest) SetRegionId(v string) *GetNetworkReachableAnalysisRequest {
	s.RegionId = &v
	return s
}

type GetNetworkReachableAnalysisResponseBody struct {
	// The unique ID (UID) of the Alibaba Cloud account.
	//
	// example:
	//
	// 123147627844****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The time when the network path was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-03-16T07:11:27Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The network path ID.
	//
	// example:
	//
	// np-2a1332214fa346b6****
	NetworkPathId *string `json:"NetworkPathId,omitempty" xml:"NetworkPathId,omitempty"`
	// The parameters of the network path.
	//
	// example:
	//
	// {
	//
	//   "sourceId": "i-bp100g5pbp6kj4p9****",
	//
	//   "sourceType": "ecs",
	//
	//   "targetId": "i-t4n4ltwgbbomzb0g****",
	//
	//   "targetType": "ecs"
	//
	// }
	NetworkPathParameter *string `json:"NetworkPathParameter,omitempty" xml:"NetworkPathParameter,omitempty"`
	// The ID of the task for analyzing network reachability.
	//
	// example:
	//
	// nra-8607514e71c1484****
	NetworkReachableAnalysisId *string `json:"NetworkReachableAnalysisId,omitempty" xml:"NetworkReachableAnalysisId,omitempty"`
	// The result of network reachability analysis, which includes the network topology, error codes of network unreachability, and rules of network unreachability.
	//
	// example:
	//
	// {
	//
	//   "errorCode": "",
	//
	//   "networkAclData": {
	//
	//     "networkAclItems": [
	//
	//
	//
	//     ]
	//
	//   },
	//
	//   "nraId": "nra-f2c8701a36424094****",
	//
	//   "requestId": "B931F8A0-620E-5230-B77F-3BD7F612****",
	//
	//   "routeData": {
	//
	//     "routeItems": [
	//
	//
	//
	//     ]
	//
	//   },
	//
	//   "securityGroupData": {
	//
	//     "policy": "accept",
	//
	//     "securityGroupItems": [
	//
	//       {
	//
	//         "description": "default_sg_access_rule",
	//
	//         "matchedRule": {
	//
	//           "bizProtocol": "ALL",
	//
	//           "creatingTime": "2022-11-10T03:24:49Z",
	//
	//           "description": "",
	//
	//           "destinationCidr": "",
	//
	//           "destinationGroupId": "sg-wz980j96p8y99co5****",
	//
	//           "direction": "egress",
	//
	//           "policy": "Accept",
	//
	//           "portRange": "-1/-1",
	//
	//           "priority": "1",
	//
	//           "sourceCidr": "",
	//
	//           "sourceGroupId": ""
	//
	//         },
	//
	//         "policy": "accept",
	//
	//         "resourceId": "eni-wz92ce4saz1jzazg****",
	//
	//         "securityGroupId": "sg-wz980j96p8y99co5****"
	//
	//       },
	//
	//       {
	//
	//         "description": "user_acl_drop_rule",
	//
	//         "matchedRule": {
	//
	//           "bizProtocol": "",
	//
	//           "creatingTime": "",
	//
	//           "description": "",
	//
	//           "destinationCidr": "",
	//
	//           "destinationGroupId": "",
	//
	//           "direction": "",
	//
	//           "policy": "",
	//
	//           "portRange": "",
	//
	//           "priority": "",
	//
	//           "sourceCidr": "",
	//
	//           "sourceGroupId": ""
	//
	//         },
	//
	//         "policy": "",
	//
	//         "resourceId": "eni-wz97vry93t6z4lbd****",
	//
	//         "securityGroupId": "sg-wz980j96p8y99co****"
	//
	//       }
	//
	//     ],
	//
	//     "securityGroupReportId": "sgr-4479d23bb37241aab****"
	//
	//   },
	//
	//   "status": "security_group_checking_target",
	//
	//   "topologyData": {
	//
	//     "positive": {
	//
	//       "linkList": [
	//
	//         {
	//
	//           "id": "i-wz91dk7bor557hp93zyv-->eni-wz92ce4saz1jzazg****",
	//
	//           "source": "i-wz91dk7bor557hp9****",
	//
	//           "target": "eni-wz92ce4saz1jzazg****"
	//
	//         },
	//
	//         {
	//
	//           "id": "eni-wz92ce4saz1jzazgi13d-->vsw-wz9slpwdcppwfrnee****",
	//
	//           "source": "eni-wz92ce4saz1jzazg****",
	//
	//           "target": "vsw-wz9slpwdcppwfrnee****"
	//
	//         },
	//
	//         {
	//
	//           "id": "vsw-wz9slpwdcppwfrneebcrp-->eni-wz97vry93t6z4lbd****",
	//
	//           "source": "vsw-wz9slpwdcppwfrnee****",
	//
	//           "target": "eni-wz97vry93t6z4lbd****"
	//
	//         },
	//
	//         {
	//
	//           "id": "eni-wz97vry93t6z4lbdgmfi-->i-wz91dk7bor557hp9****",
	//
	//           "source": "eni-wz97vry93t6z4lbd****",
	//
	//           "target": "i-wz91dk7bor557hp9****"
	//
	//         }
	//
	//       ],
	//
	//       "nodeList": [
	//
	//         {
	//
	//           "aZone": "cn-shenzhen-d",
	//
	//           "bizInsId": "i-wz91dk7bor557hp9****",
	//
	//           "id": "i-wz91dk7bor557hp9****",
	//
	//           "level": 1,
	//
	//           "matchedRoute": {
	//
	//             "nextHopSet": [
	//
	//
	//
	//             ]
	//
	//           },
	//
	//           "nodeType": "VM",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "trafficLogs": [
	//
	//
	//
	//           ]
	//
	//         },
	//
	//         {
	//
	//           "aZone": "cn-shenzhen-d",
	//
	//           "bizInsId": "i-wz91dk7bor557hp9****",
	//
	//           "id": "i-wz91dk7bor557hp9****",
	//
	//           "level": 3,
	//
	//           "matchedRoute": {
	//
	//             "nextHopSet": [
	//
	//
	//
	//             ]
	//
	//           },
	//
	//           "nodeType": "VM",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "trafficLogs": [
	//
	//
	//
	//           ]
	//
	//         },
	//
	//         {
	//
	//           "aZone": "cn-shenzhen-d",
	//
	//           "bizInsId": "vsw-wz9slpwdcppwfrnee****",
	//
	//           "cidr": "192.168.0.0/24",
	//
	//           "id": "vsw-wz9slpwdcppwfrnee****",
	//
	//           "level": 2,
	//
	//           "matchedRoute": {
	//
	//             "nextHopSet": [
	//
	//
	//
	//             ]
	//
	//           },
	//
	//           "nodeType": "VSW",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "trafficLogs": [
	//
	//
	//
	//           ]
	//
	//         },
	//
	//         {
	//
	//           "bizInsId": "eni-wz92ce4saz1jzazg****",
	//
	//           "id": "eni-wz92ce4saz1jzazg****",
	//
	//           "ip": "192.168.0.33",
	//
	//           "mac": "00:XXXX:3e:16:7c:50",
	//
	//           "matchedRoute": {
	//
	//             "nextHopSet": [
	//
	//
	//
	//             ]
	//
	//           },
	//
	//           "nodeType": "ENI",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "status": "InUse",
	//
	//           "trafficLogs": [
	//
	//
	//
	//           ]
	//
	//         },
	//
	//         {
	//
	//           "bizInsId": "eni-wz97vry93t6z4lbd****",
	//
	//           "id": "eni-wz97vry93t6z4lbd****",
	//
	//           "ip": "192.168.0.34",
	//
	//           "mac": "00:XXXX:3e:14:70:c2",
	//
	//           "matchedRoute": {
	//
	//             "nextHopSet": [
	//
	//
	//
	//             ]
	//
	//           },
	//
	//           "nodeType": "ENI",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "status": "InUse",
	//
	//           "trafficLogs": [
	//
	//
	//
	//           ]
	//
	//         }
	//
	//       ]
	//
	//     },
	//
	//     "reverse": {
	//
	//       "revLinkList": [
	//
	//         {
	//
	//           "id": "i-wz91dk7bor557hp93zys-->eni-wz97vry93t6z4lbd****",
	//
	//           "source": "i-wz91dk7bor557hp9****",
	//
	//           "target": "eni-wz97vry93t6z4lbd****"
	//
	//         },
	//
	//         {
	//
	//           "id": "eni-wz97vry93t6z4lbdgmfi-->vsw-wz9slpwdcppwfrnee****",
	//
	//           "source": "eni-wz97vry93t6z4lbd****",
	//
	//           "target": "vsw-wz9slpwdcppwfrnee****"
	//
	//         },
	//
	//         {
	//
	//           "id": "vsw-wz9slpwdcppwfrneebcrp-->eni-wz92ce4saz1jzazg****",
	//
	//           "source": "vsw-wz9slpwdcppwfrnee****",
	//
	//           "target": "eni-wz92ce4saz1jzazg****"
	//
	//         },
	//
	//         {
	//
	//           "id": "eni-wz92ce4saz1jzazgi13d-->i-wz91dk7bor557hp9****",
	//
	//           "source": "eni-wz92ce4saz1jzazg****",
	//
	//           "target": "i-wz91dk7bor557hp9****"
	//
	//         }
	//
	//       ],
	//
	//       "revNodeList": [
	//
	//         {
	//
	//           "aZone": "cn-shenzhen-d",
	//
	//           "bizInsId": "i-wz91dk7bor557hp9****",
	//
	//           "id": "i-wz91dk7bor557hp9****",
	//
	//           "level": 1,
	//
	//           "nodeType": "VM",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "revMatchedRoute": {
	//
	//             "revNextHopSet": [
	//
	//
	//
	//             ]
	//
	//           }
	//
	//         },
	//
	//         {
	//
	//           "aZone": "cn-shenzhen-d",
	//
	//           "bizInsId": "i-wz91dk7bor557hp9****",
	//
	//           "id": "i-wz91dk7bor557hp9****",
	//
	//           "level": 3,
	//
	//           "nodeType": "VM",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "revMatchedRoute": {
	//
	//             "revNextHopSet": [
	//
	//
	//
	//             ]
	//
	//           }
	//
	//         },
	//
	//         {
	//
	//           "aZone": "cn-shenzhen-d",
	//
	//           "bizInsId": "vsw-wz9slpwdcppwfrnee****",
	//
	//           "cidr": "192.168.0.0/24",
	//
	//           "id": "vsw-wz9slpwdcppwfrnee****",
	//
	//           "level": 2,
	//
	//           "nodeType": "VSW",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "revMatchedRoute": {
	//
	//             "revNextHopSet": [
	//
	//
	//
	//             ]
	//
	//           }
	//
	//         },
	//
	//         {
	//
	//           "bizInsId": "eni-wz97vry93t6z4lbd****",
	//
	//           "id": "eni-wz97vry93t6z4lbd****",
	//
	//           "ip": "192.168.0.34",
	//
	//           "mac": "00:XXXX:3e:14:70:c2",
	//
	//           "nodeType": "ENI",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "revMatchedRoute": {
	//
	//             "revNextHopSet": [
	//
	//
	//
	//             ]
	//
	//           },
	//
	//           "status": "InUse"
	//
	//         },
	//
	//         {
	//
	//           "bizInsId": "eni-wz92ce4saz1jzazg****",
	//
	//           "id": "eni-wz92ce4saz1jzazg****",
	//
	//           "ip": "192.168.0.33",
	//
	//           "mac": "00:XXXX:3e:16:7c:50",
	//
	//           "nodeType": "ENI",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "revMatchedRoute": {
	//
	//             "revNextHopSet": [
	//
	//
	//
	//             ]
	//
	//           },
	//
	//           "status": "InUse"
	//
	//         }
	//
	//       ]
	//
	//     },
	//
	//     "topologyReportId": "tpr-21cf60002715491b8****"
	//
	//   }
	//
	// }
	NetworkReachableAnalysisResult *string `json:"NetworkReachableAnalysisResult,omitempty" xml:"NetworkReachableAnalysisResult,omitempty"`
	// The state of the task for analyzing network reachability. Valid values:
	//
	// 	- **init**: The task is in progress.
	//
	// 	- **finish**: The task is complete.
	//
	// 	- **error**: An analysis error occurred.
	//
	// 	- **timeout**: The task timed out.
	//
	// example:
	//
	// finish
	NetworkReachableAnalysisStatus *string `json:"NetworkReachableAnalysisStatus,omitempty" xml:"NetworkReachableAnalysisStatus,omitempty"`
	// Indicates whether the network path is reachable. Valid values:
	//
	// 	- **true**: The network path is reachable.
	//
	// 	- **false**: The network path is unreachable.
	//
	// example:
	//
	// true
	Reachable *bool `json:"Reachable,omitempty" xml:"Reachable,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DEE0FEAF-59AE-5CDD-AA07-626BC365D571
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNetworkReachableAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkReachableAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetworkReachableAnalysisResponseBody) SetAliUid(v int64) *GetNetworkReachableAnalysisResponseBody {
	s.AliUid = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) SetCreateTime(v string) *GetNetworkReachableAnalysisResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) SetNetworkPathId(v string) *GetNetworkReachableAnalysisResponseBody {
	s.NetworkPathId = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) SetNetworkPathParameter(v string) *GetNetworkReachableAnalysisResponseBody {
	s.NetworkPathParameter = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) SetNetworkReachableAnalysisId(v string) *GetNetworkReachableAnalysisResponseBody {
	s.NetworkReachableAnalysisId = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) SetNetworkReachableAnalysisResult(v string) *GetNetworkReachableAnalysisResponseBody {
	s.NetworkReachableAnalysisResult = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) SetNetworkReachableAnalysisStatus(v string) *GetNetworkReachableAnalysisResponseBody {
	s.NetworkReachableAnalysisStatus = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) SetReachable(v bool) *GetNetworkReachableAnalysisResponseBody {
	s.Reachable = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) SetRequestId(v string) *GetNetworkReachableAnalysisResponseBody {
	s.RequestId = &v
	return s
}

type GetNetworkReachableAnalysisResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNetworkReachableAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNetworkReachableAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkReachableAnalysisResponse) GoString() string {
	return s.String()
}

func (s *GetNetworkReachableAnalysisResponse) SetHeaders(v map[string]*string) *GetNetworkReachableAnalysisResponse {
	s.Headers = v
	return s
}

func (s *GetNetworkReachableAnalysisResponse) SetStatusCode(v int32) *GetNetworkReachableAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponse) SetBody(v *GetNetworkReachableAnalysisResponseBody) *GetNetworkReachableAnalysisResponse {
	s.Body = v
	return s
}

type GetNisNetworkMetricsRequest struct {
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1638239092000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// This parameter is required.
	Dimensions []*GetNisNetworkMetricsRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// example:
	//
	// 1684373700099
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bps
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AccessInternetIPV4
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// TimestampAscending
	ScanBy *string `json:"ScanBy,omitempty" xml:"ScanBy,omitempty"`
	// example:
	//
	// false
	UseCrossAccount *bool `json:"UseCrossAccount,omitempty" xml:"UseCrossAccount,omitempty"`
}

func (s GetNisNetworkMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNisNetworkMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsRequest) SetAccountIds(v []*string) *GetNisNetworkMetricsRequest {
	s.AccountIds = v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetBeginTime(v int64) *GetNisNetworkMetricsRequest {
	s.BeginTime = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetDimensions(v []*GetNisNetworkMetricsRequestDimensions) *GetNisNetworkMetricsRequest {
	s.Dimensions = v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetEndTime(v int64) *GetNisNetworkMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetMetricName(v string) *GetNisNetworkMetricsRequest {
	s.MetricName = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetRegionNo(v string) *GetNisNetworkMetricsRequest {
	s.RegionNo = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetResourceType(v string) *GetNisNetworkMetricsRequest {
	s.ResourceType = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetScanBy(v string) *GetNisNetworkMetricsRequest {
	s.ScanBy = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetUseCrossAccount(v bool) *GetNisNetworkMetricsRequest {
	s.UseCrossAccount = &v
	return s
}

type GetNisNetworkMetricsRequestDimensions struct {
	// example:
	//
	// instanceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// eip-sample*
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetNisNetworkMetricsRequestDimensions) String() string {
	return tea.Prettify(s)
}

func (s GetNisNetworkMetricsRequestDimensions) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsRequestDimensions) SetName(v string) *GetNisNetworkMetricsRequestDimensions {
	s.Name = &v
	return s
}

func (s *GetNisNetworkMetricsRequestDimensions) SetValue(v string) *GetNisNetworkMetricsRequestDimensions {
	s.Value = &v
	return s
}

type GetNisNetworkMetricsShrinkRequest struct {
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1638239092000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// This parameter is required.
	DimensionsShrink *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// example:
	//
	// 1684373700099
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bps
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AccessInternetIPV4
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// TimestampAscending
	ScanBy *string `json:"ScanBy,omitempty" xml:"ScanBy,omitempty"`
	// example:
	//
	// false
	UseCrossAccount *bool `json:"UseCrossAccount,omitempty" xml:"UseCrossAccount,omitempty"`
}

func (s GetNisNetworkMetricsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNisNetworkMetricsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsShrinkRequest) SetAccountIds(v []*string) *GetNisNetworkMetricsShrinkRequest {
	s.AccountIds = v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetBeginTime(v int64) *GetNisNetworkMetricsShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetDimensionsShrink(v string) *GetNisNetworkMetricsShrinkRequest {
	s.DimensionsShrink = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetEndTime(v int64) *GetNisNetworkMetricsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetMetricName(v string) *GetNisNetworkMetricsShrinkRequest {
	s.MetricName = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetRegionNo(v string) *GetNisNetworkMetricsShrinkRequest {
	s.RegionNo = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetResourceType(v string) *GetNisNetworkMetricsShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetScanBy(v string) *GetNisNetworkMetricsShrinkRequest {
	s.ScanBy = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetUseCrossAccount(v bool) *GetNisNetworkMetricsShrinkRequest {
	s.UseCrossAccount = &v
	return s
}

type GetNisNetworkMetricsResponseBody struct {
	Data *GetNisNetworkMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D5E98683-355B-5867-8D3D-A24755F6895B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNisNetworkMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNisNetworkMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsResponseBody) SetData(v *GetNisNetworkMetricsResponseBodyData) *GetNisNetworkMetricsResponseBody {
	s.Data = v
	return s
}

func (s *GetNisNetworkMetricsResponseBody) SetRequestId(v string) *GetNisNetworkMetricsResponseBody {
	s.RequestId = &v
	return s
}

type GetNisNetworkMetricsResponseBodyData struct {
	Metrics []*GetNisNetworkMetricsResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// Bits/Second
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s GetNisNetworkMetricsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetNisNetworkMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsResponseBodyData) SetMetrics(v []*GetNisNetworkMetricsResponseBodyDataMetrics) *GetNisNetworkMetricsResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetNisNetworkMetricsResponseBodyData) SetUnit(v string) *GetNisNetworkMetricsResponseBodyData {
	s.Unit = &v
	return s
}

type GetNisNetworkMetricsResponseBodyDataMetrics struct {
	// example:
	//
	// 1690684091100
	TimeStamp *int64 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// example:
	//
	// 88
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetNisNetworkMetricsResponseBodyDataMetrics) String() string {
	return tea.Prettify(s)
}

func (s GetNisNetworkMetricsResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsResponseBodyDataMetrics) SetTimeStamp(v int64) *GetNisNetworkMetricsResponseBodyDataMetrics {
	s.TimeStamp = &v
	return s
}

func (s *GetNisNetworkMetricsResponseBodyDataMetrics) SetValue(v float64) *GetNisNetworkMetricsResponseBodyDataMetrics {
	s.Value = &v
	return s
}

type GetNisNetworkMetricsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNisNetworkMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNisNetworkMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNisNetworkMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsResponse) SetHeaders(v map[string]*string) *GetNisNetworkMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetNisNetworkMetricsResponse) SetStatusCode(v int32) *GetNisNetworkMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNisNetworkMetricsResponse) SetBody(v *GetNisNetworkMetricsResponseBody) *GetNisNetworkMetricsResponse {
	s.Body = v
	return s
}

type GetNisNetworkRankingRequest struct {
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1684373600099
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// 1684379093000
	EndTime *int64                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Filter  []*GetNisNetworkRankingRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Protocol
	GroupBy *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bps
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AccessInternetIpV4
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 10
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
	// example:
	//
	// false
	UseCrossAccount *bool `json:"UseCrossAccount,omitempty" xml:"UseCrossAccount,omitempty"`
}

func (s GetNisNetworkRankingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNisNetworkRankingRequest) GoString() string {
	return s.String()
}

func (s *GetNisNetworkRankingRequest) SetAccountIds(v []*string) *GetNisNetworkRankingRequest {
	s.AccountIds = v
	return s
}

func (s *GetNisNetworkRankingRequest) SetBeginTime(v int64) *GetNisNetworkRankingRequest {
	s.BeginTime = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetDirection(v string) *GetNisNetworkRankingRequest {
	s.Direction = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetEndTime(v int64) *GetNisNetworkRankingRequest {
	s.EndTime = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetFilter(v []*GetNisNetworkRankingRequestFilter) *GetNisNetworkRankingRequest {
	s.Filter = v
	return s
}

func (s *GetNisNetworkRankingRequest) SetGroupBy(v string) *GetNisNetworkRankingRequest {
	s.GroupBy = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetOrderBy(v string) *GetNisNetworkRankingRequest {
	s.OrderBy = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetRegionNo(v string) *GetNisNetworkRankingRequest {
	s.RegionNo = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetResourceType(v string) *GetNisNetworkRankingRequest {
	s.ResourceType = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetSort(v string) *GetNisNetworkRankingRequest {
	s.Sort = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetTopN(v int32) *GetNisNetworkRankingRequest {
	s.TopN = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetUseCrossAccount(v bool) *GetNisNetworkRankingRequest {
	s.UseCrossAccount = &v
	return s
}

type GetNisNetworkRankingRequestFilter struct {
	// example:
	//
	// instanceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// lb-2zxxxxz1d
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetNisNetworkRankingRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s GetNisNetworkRankingRequestFilter) GoString() string {
	return s.String()
}

func (s *GetNisNetworkRankingRequestFilter) SetName(v string) *GetNisNetworkRankingRequestFilter {
	s.Name = &v
	return s
}

func (s *GetNisNetworkRankingRequestFilter) SetValue(v string) *GetNisNetworkRankingRequestFilter {
	s.Value = &v
	return s
}

type GetNisNetworkRankingShrinkRequest struct {
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1684373600099
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// 1684379093000
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Protocol
	GroupBy *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bps
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AccessInternetIpV4
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 10
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
	// example:
	//
	// false
	UseCrossAccount *bool `json:"UseCrossAccount,omitempty" xml:"UseCrossAccount,omitempty"`
}

func (s GetNisNetworkRankingShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNisNetworkRankingShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetNisNetworkRankingShrinkRequest) SetAccountIds(v []*string) *GetNisNetworkRankingShrinkRequest {
	s.AccountIds = v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetBeginTime(v int64) *GetNisNetworkRankingShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetDirection(v string) *GetNisNetworkRankingShrinkRequest {
	s.Direction = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetEndTime(v int64) *GetNisNetworkRankingShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetFilterShrink(v string) *GetNisNetworkRankingShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetGroupBy(v string) *GetNisNetworkRankingShrinkRequest {
	s.GroupBy = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetOrderBy(v string) *GetNisNetworkRankingShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetRegionNo(v string) *GetNisNetworkRankingShrinkRequest {
	s.RegionNo = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetResourceType(v string) *GetNisNetworkRankingShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetSort(v string) *GetNisNetworkRankingShrinkRequest {
	s.Sort = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetTopN(v int32) *GetNisNetworkRankingShrinkRequest {
	s.TopN = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetUseCrossAccount(v bool) *GetNisNetworkRankingShrinkRequest {
	s.UseCrossAccount = &v
	return s
}

type GetNisNetworkRankingResponseBody struct {
	Data []*GetNisNetworkRankingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// A7F0D6EC-E19E-58AC-AC9F-08036763960F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNisNetworkRankingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNisNetworkRankingResponseBody) GoString() string {
	return s.String()
}

func (s *GetNisNetworkRankingResponseBody) SetData(v []*GetNisNetworkRankingResponseBodyData) *GetNisNetworkRankingResponseBody {
	s.Data = v
	return s
}

func (s *GetNisNetworkRankingResponseBody) SetRequestId(v string) *GetNisNetworkRankingResponseBody {
	s.RequestId = &v
	return s
}

type GetNisNetworkRankingResponseBodyData struct {
	// example:
	//
	// 66
	ActiveSessionCount *float64 `json:"ActiveSessionCount,omitempty" xml:"ActiveSessionCount,omitempty"`
	// example:
	//
	// 129103
	Asn *string `json:"Asn,omitempty" xml:"Asn,omitempty"`
	// example:
	//
	// tr-sample*
	AttachmentId *string `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	// example:
	//
	// cbwp-sample*
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// example:
	//
	// 1024
	ByteCount *float64 `json:"ByteCount,omitempty" xml:"ByteCount,omitempty"`
	City      *string  `json:"City,omitempty" xml:"City,omitempty"`
	Country   *string  `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// 2.2.XX.XX
	DestinationIp  *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	DestinationIsp *string `json:"DestinationIsp,omitempty" xml:"DestinationIsp,omitempty"`
	// example:
	//
	// 80
	DestinationPort *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	// example:
	//
	// cn-hangzhou
	DestinationRegionNo *string `json:"DestinationRegionNo,omitempty" xml:"DestinationRegionNo,omitempty"`
	// example:
	//
	// cn-hangzhou-b
	DestinationZone *string `json:"DestinationZone,omitempty" xml:"DestinationZone,omitempty"`
	// example:
	//
	// 120.238.XX.XX
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// example:
	//
	// 10
	InBps *float64 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	// example:
	//
	// 3
	InPps *float64 `json:"InPps,omitempty" xml:"InPps,omitempty"`
	// example:
	//
	// eip-sample*
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Isp        *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// example:
	//
	// 18
	NewSessionPerSecond *float64 `json:"NewSessionPerSecond,omitempty" xml:"NewSessionPerSecond,omitempty"`
	// example:
	//
	// 88
	OutBps *float64 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
	// example:
	//
	// 8
	OutPps *float64 `json:"OutPps,omitempty" xml:"OutPps,omitempty"`
	// example:
	//
	// 66
	PacketCount *float64 `json:"PacketCount,omitempty" xml:"PacketCount,omitempty"`
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// example:
	//
	// 23
	RTT *float64 `json:"RTT,omitempty" xml:"RTT,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// 0.1
	RetransmitRate *float64 `json:"RetransmitRate,omitempty" xml:"RetransmitRate,omitempty"`
	// example:
	//
	// 42.120.XX.XX
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	SourceIsp *string `json:"SourceIsp,omitempty" xml:"SourceIsp,omitempty"`
	// example:
	//
	// 443
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	SourceZone *string `json:"SourceZone,omitempty" xml:"SourceZone,omitempty"`
	// example:
	//
	// vbr-sample*
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
}

func (s GetNisNetworkRankingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetNisNetworkRankingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNisNetworkRankingResponseBodyData) SetActiveSessionCount(v float64) *GetNisNetworkRankingResponseBodyData {
	s.ActiveSessionCount = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetAsn(v string) *GetNisNetworkRankingResponseBodyData {
	s.Asn = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetAttachmentId(v string) *GetNisNetworkRankingResponseBodyData {
	s.AttachmentId = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetBandwidthPackageId(v string) *GetNisNetworkRankingResponseBodyData {
	s.BandwidthPackageId = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetByteCount(v float64) *GetNisNetworkRankingResponseBodyData {
	s.ByteCount = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetCity(v string) *GetNisNetworkRankingResponseBodyData {
	s.City = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetCountry(v string) *GetNisNetworkRankingResponseBodyData {
	s.Country = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetDestinationIp(v string) *GetNisNetworkRankingResponseBodyData {
	s.DestinationIp = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetDestinationIsp(v string) *GetNisNetworkRankingResponseBodyData {
	s.DestinationIsp = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetDestinationPort(v string) *GetNisNetworkRankingResponseBodyData {
	s.DestinationPort = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetDestinationRegionNo(v string) *GetNisNetworkRankingResponseBodyData {
	s.DestinationRegionNo = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetDestinationZone(v string) *GetNisNetworkRankingResponseBodyData {
	s.DestinationZone = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetIP(v string) *GetNisNetworkRankingResponseBodyData {
	s.IP = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetInBps(v float64) *GetNisNetworkRankingResponseBodyData {
	s.InBps = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetInPps(v float64) *GetNisNetworkRankingResponseBodyData {
	s.InPps = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetInstanceId(v string) *GetNisNetworkRankingResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetIsp(v string) *GetNisNetworkRankingResponseBodyData {
	s.Isp = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetNewSessionPerSecond(v float64) *GetNisNetworkRankingResponseBodyData {
	s.NewSessionPerSecond = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetOutBps(v float64) *GetNisNetworkRankingResponseBodyData {
	s.OutBps = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetOutPps(v float64) *GetNisNetworkRankingResponseBodyData {
	s.OutPps = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetPacketCount(v float64) *GetNisNetworkRankingResponseBodyData {
	s.PacketCount = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetProtocol(v string) *GetNisNetworkRankingResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetProvince(v string) *GetNisNetworkRankingResponseBodyData {
	s.Province = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetRTT(v float64) *GetNisNetworkRankingResponseBodyData {
	s.RTT = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetRegionNo(v string) *GetNisNetworkRankingResponseBodyData {
	s.RegionNo = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetRetransmitRate(v float64) *GetNisNetworkRankingResponseBodyData {
	s.RetransmitRate = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetSourceIp(v string) *GetNisNetworkRankingResponseBodyData {
	s.SourceIp = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetSourceIsp(v string) *GetNisNetworkRankingResponseBodyData {
	s.SourceIsp = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetSourcePort(v string) *GetNisNetworkRankingResponseBodyData {
	s.SourcePort = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetSourceZone(v string) *GetNisNetworkRankingResponseBodyData {
	s.SourceZone = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetVbrId(v string) *GetNisNetworkRankingResponseBodyData {
	s.VbrId = &v
	return s
}

type GetNisNetworkRankingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNisNetworkRankingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNisNetworkRankingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNisNetworkRankingResponse) GoString() string {
	return s.String()
}

func (s *GetNisNetworkRankingResponse) SetHeaders(v map[string]*string) *GetNisNetworkRankingResponse {
	s.Headers = v
	return s
}

func (s *GetNisNetworkRankingResponse) SetStatusCode(v int32) *GetNisNetworkRankingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNisNetworkRankingResponse) SetBody(v *GetNisNetworkRankingResponseBody) *GetNisNetworkRankingResponse {
	s.Body = v
	return s
}

type GetTransitRouterFlowTopNRequest struct {
	// The IDs of the member accounts.
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The ID of the CEN bandwidth plan.
	//
	// example:
	//
	// cenbwp-ia8kw1zjv4hyal*****
	BandwithPackageId *string `json:"BandwithPackageId,omitempty" xml:"BandwithPackageId,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. The maximum time range that you can query is 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684373600099
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The CEN instance ID.
	//
	// example:
	//
	// cen-ia8kw1zjv4hyal****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The direction of the inter-region traffic in the local regions or for the local IP addresses. Valid values:
	//
	// 	- **in**: inbound traffic
	//
	// 	- **out**: outbound traffic
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. The maximum time range that you can query is 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638239093000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The dimension for ranking inter-region traffic data. The value of this parameter is case-sensitive. Valid values:
	//
	// 	- **1Tuple**: queries the rankings of inter-region traffic data for the local regions, Cloud Enterprise Network (CEN) instances, and IP addresses.
	//
	// 	- **2Tuple**: queries the rankings of inter-region traffic data for the local and remote regions, and the local and remote IP addresses.
	//
	// 	- **5Tuple**: queries the rankings of inter-region traffic data for the local and remote IP addresses, local and remote ports, and protocols.
	//
	// 	- **Cen**: queries the rankings of inter-region traffic data for CEN instances.
	//
	// 	- **RegionPair**: queries the rankings of inter-region traffic data for the local and remote regions.
	//
	// 	- **Port**: queries the rankings of inter-region traffic data for the local and remote ports.
	//
	// 	- **Protocol**: queries the rankings of inter-region traffic data for the protocols.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1Tuple
	GroupBy *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// The metric for ranking inter-region traffic data. Default value: Bytes. This value specifies that inter-region traffic data is ranked by traffic volume.
	//
	// example:
	//
	// Bytes
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The remote IP address.
	//
	// example:
	//
	// 122.112.XX.XX
	OtherIp *string `json:"OtherIp,omitempty" xml:"OtherIp,omitempty"`
	// The remote port.
	//
	// example:
	//
	// 10869
	OtherPort *string `json:"OtherPort,omitempty" xml:"OtherPort,omitempty"`
	// The remote region.
	//
	// example:
	//
	// ap-southeast-1
	OtherRegion *string `json:"OtherRegion,omitempty" xml:"OtherRegion,omitempty"`
	// The protocol number.
	//
	// >  All protocols are supported. This parameter is required only if you set **GroupBy*	- to **5Tuple*	- or **Protocol**.
	//
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The order for ranking inter-region traffic data. Valid values:
	//
	// 	- **desc**: descending order
	//
	// 	- **asc**: ascending order
	//
	// example:
	//
	// desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The local IP address.
	//
	// example:
	//
	// 1.8.XX.XX
	ThisIp *string `json:"ThisIp,omitempty" xml:"ThisIp,omitempty"`
	// The local port.
	//
	// example:
	//
	// 80
	ThisPort *string `json:"ThisPort,omitempty" xml:"ThisPort,omitempty"`
	// The local region where the **local IP address*	- resides.
	//
	// example:
	//
	// cn-shanghai
	ThisRegion *string `json:"ThisRegion,omitempty" xml:"ThisRegion,omitempty"`
	// Specifies the maximum number of data entries to display. Default value: **10**. Maximum value: 100.
	//
	// example:
	//
	// 20
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
	// Specifies whether to enable the multi-account management feature. Default value: **false**. This value specifies that the multi-account management feature is disabled.
	//
	// >  By default, the multi-account management feature is not available. If you want to use this feature, contact your account manager to apply for permissions.
	//
	// example:
	//
	// false
	UseMultiAccount *bool `json:"UseMultiAccount,omitempty" xml:"UseMultiAccount,omitempty"`
}

func (s GetTransitRouterFlowTopNRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTransitRouterFlowTopNRequest) GoString() string {
	return s.String()
}

func (s *GetTransitRouterFlowTopNRequest) SetAccountIds(v []*int64) *GetTransitRouterFlowTopNRequest {
	s.AccountIds = v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetBandwithPackageId(v string) *GetTransitRouterFlowTopNRequest {
	s.BandwithPackageId = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetBeginTime(v int64) *GetTransitRouterFlowTopNRequest {
	s.BeginTime = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetCenId(v string) *GetTransitRouterFlowTopNRequest {
	s.CenId = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetDirection(v string) *GetTransitRouterFlowTopNRequest {
	s.Direction = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetEndTime(v int64) *GetTransitRouterFlowTopNRequest {
	s.EndTime = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetGroupBy(v string) *GetTransitRouterFlowTopNRequest {
	s.GroupBy = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetOrderBy(v string) *GetTransitRouterFlowTopNRequest {
	s.OrderBy = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetOtherIp(v string) *GetTransitRouterFlowTopNRequest {
	s.OtherIp = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetOtherPort(v string) *GetTransitRouterFlowTopNRequest {
	s.OtherPort = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetOtherRegion(v string) *GetTransitRouterFlowTopNRequest {
	s.OtherRegion = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetProtocol(v string) *GetTransitRouterFlowTopNRequest {
	s.Protocol = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetSort(v string) *GetTransitRouterFlowTopNRequest {
	s.Sort = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetThisIp(v string) *GetTransitRouterFlowTopNRequest {
	s.ThisIp = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetThisPort(v string) *GetTransitRouterFlowTopNRequest {
	s.ThisPort = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetThisRegion(v string) *GetTransitRouterFlowTopNRequest {
	s.ThisRegion = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetTopN(v int32) *GetTransitRouterFlowTopNRequest {
	s.TopN = &v
	return s
}

func (s *GetTransitRouterFlowTopNRequest) SetUseMultiAccount(v bool) *GetTransitRouterFlowTopNRequest {
	s.UseMultiAccount = &v
	return s
}

type GetTransitRouterFlowTopNShrinkRequest struct {
	// The IDs of the member accounts.
	AccountIdsShrink *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	// The ID of the CEN bandwidth plan.
	//
	// example:
	//
	// cenbwp-ia8kw1zjv4hyal*****
	BandwithPackageId *string `json:"BandwithPackageId,omitempty" xml:"BandwithPackageId,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. The maximum time range that you can query is 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684373600099
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The CEN instance ID.
	//
	// example:
	//
	// cen-ia8kw1zjv4hyal****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The direction of the inter-region traffic in the local regions or for the local IP addresses. Valid values:
	//
	// 	- **in**: inbound traffic
	//
	// 	- **out**: outbound traffic
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. The maximum time range that you can query is 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638239093000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The dimension for ranking inter-region traffic data. The value of this parameter is case-sensitive. Valid values:
	//
	// 	- **1Tuple**: queries the rankings of inter-region traffic data for the local regions, Cloud Enterprise Network (CEN) instances, and IP addresses.
	//
	// 	- **2Tuple**: queries the rankings of inter-region traffic data for the local and remote regions, and the local and remote IP addresses.
	//
	// 	- **5Tuple**: queries the rankings of inter-region traffic data for the local and remote IP addresses, local and remote ports, and protocols.
	//
	// 	- **Cen**: queries the rankings of inter-region traffic data for CEN instances.
	//
	// 	- **RegionPair**: queries the rankings of inter-region traffic data for the local and remote regions.
	//
	// 	- **Port**: queries the rankings of inter-region traffic data for the local and remote ports.
	//
	// 	- **Protocol**: queries the rankings of inter-region traffic data for the protocols.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1Tuple
	GroupBy *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// The metric for ranking inter-region traffic data. Default value: Bytes. This value specifies that inter-region traffic data is ranked by traffic volume.
	//
	// example:
	//
	// Bytes
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The remote IP address.
	//
	// example:
	//
	// 122.112.XX.XX
	OtherIp *string `json:"OtherIp,omitempty" xml:"OtherIp,omitempty"`
	// The remote port.
	//
	// example:
	//
	// 10869
	OtherPort *string `json:"OtherPort,omitempty" xml:"OtherPort,omitempty"`
	// The remote region.
	//
	// example:
	//
	// ap-southeast-1
	OtherRegion *string `json:"OtherRegion,omitempty" xml:"OtherRegion,omitempty"`
	// The protocol number.
	//
	// >  All protocols are supported. This parameter is required only if you set **GroupBy*	- to **5Tuple*	- or **Protocol**.
	//
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The order for ranking inter-region traffic data. Valid values:
	//
	// 	- **desc**: descending order
	//
	// 	- **asc**: ascending order
	//
	// example:
	//
	// desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The local IP address.
	//
	// example:
	//
	// 1.8.XX.XX
	ThisIp *string `json:"ThisIp,omitempty" xml:"ThisIp,omitempty"`
	// The local port.
	//
	// example:
	//
	// 80
	ThisPort *string `json:"ThisPort,omitempty" xml:"ThisPort,omitempty"`
	// The local region where the **local IP address*	- resides.
	//
	// example:
	//
	// cn-shanghai
	ThisRegion *string `json:"ThisRegion,omitempty" xml:"ThisRegion,omitempty"`
	// Specifies the maximum number of data entries to display. Default value: **10**. Maximum value: 100.
	//
	// example:
	//
	// 20
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
	// Specifies whether to enable the multi-account management feature. Default value: **false**. This value specifies that the multi-account management feature is disabled.
	//
	// >  By default, the multi-account management feature is not available. If you want to use this feature, contact your account manager to apply for permissions.
	//
	// example:
	//
	// false
	UseMultiAccount *bool `json:"UseMultiAccount,omitempty" xml:"UseMultiAccount,omitempty"`
}

func (s GetTransitRouterFlowTopNShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTransitRouterFlowTopNShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetAccountIdsShrink(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetBandwithPackageId(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.BandwithPackageId = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetBeginTime(v int64) *GetTransitRouterFlowTopNShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetCenId(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.CenId = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetDirection(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.Direction = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetEndTime(v int64) *GetTransitRouterFlowTopNShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetGroupBy(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.GroupBy = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetOrderBy(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetOtherIp(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.OtherIp = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetOtherPort(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.OtherPort = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetOtherRegion(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.OtherRegion = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetProtocol(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetSort(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.Sort = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetThisIp(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.ThisIp = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetThisPort(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.ThisPort = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetThisRegion(v string) *GetTransitRouterFlowTopNShrinkRequest {
	s.ThisRegion = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetTopN(v int32) *GetTransitRouterFlowTopNShrinkRequest {
	s.TopN = &v
	return s
}

func (s *GetTransitRouterFlowTopNShrinkRequest) SetUseMultiAccount(v bool) *GetTransitRouterFlowTopNShrinkRequest {
	s.UseMultiAccount = &v
	return s
}

type GetTransitRouterFlowTopNResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D5E98683-355B-5867-8D3D-A24755F6895B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ranking result of inter-region traffic data.
	TransitRouterFlowTopN []*GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN `json:"TransitRouterFlowTopN,omitempty" xml:"TransitRouterFlowTopN,omitempty" type:"Repeated"`
}

func (s GetTransitRouterFlowTopNResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTransitRouterFlowTopNResponseBody) GoString() string {
	return s.String()
}

func (s *GetTransitRouterFlowTopNResponseBody) SetRequestId(v string) *GetTransitRouterFlowTopNResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBody) SetTransitRouterFlowTopN(v []*GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) *GetTransitRouterFlowTopNResponseBody {
	s.TransitRouterFlowTopN = v
	return s
}

type GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN struct {
	// The account ID.
	//
	// example:
	//
	// 118639953821xxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the CEN bandwidth plan.
	//
	// example:
	//
	// cenbwp-ia8kw1zjv4hyal****
	BandwithPackageId *string `json:"BandwithPackageId,omitempty" xml:"BandwithPackageId,omitempty"`
	// The total volume of traffic in the specified time range.
	//
	// example:
	//
	// 188
	Bytes *float64 `json:"Bytes,omitempty" xml:"Bytes,omitempty"`
	// The CEN instance ID.
	//
	// example:
	//
	// cen-ia8kw1zjv4hyal****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The end of the time range that you queried. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-01-31T06:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The remote IP address.
	//
	// example:
	//
	// 47.216.XX.XX
	OtherIp *string `json:"OtherIp,omitempty" xml:"OtherIp,omitempty"`
	// The remote port.
	//
	// example:
	//
	// 53470
	OtherPort *string `json:"OtherPort,omitempty" xml:"OtherPort,omitempty"`
	// The remote region where the **remote IP address*	- resides.
	//
	// example:
	//
	// ap-southeast-1
	OtherRegion *string `json:"OtherRegion,omitempty" xml:"OtherRegion,omitempty"`
	// The total number of packets in the specified time range.
	//
	// example:
	//
	// 88
	Packets *float64 `json:"Packets,omitempty" xml:"Packets,omitempty"`
	// The protocol number.
	//
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The beginning of the time range that you queried. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-01-31T05:40:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The local IP address.
	//
	// example:
	//
	// 1.8.XX.XX
	ThisIp *string `json:"ThisIp,omitempty" xml:"ThisIp,omitempty"`
	// The local port.
	//
	// example:
	//
	// 80
	ThisPort *string `json:"ThisPort,omitempty" xml:"ThisPort,omitempty"`
	// The local region where the **local IP address*	- resides.
	//
	// example:
	//
	// cn-shanghai
	ThisRegion *string `json:"ThisRegion,omitempty" xml:"ThisRegion,omitempty"`
}

func (s GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) String() string {
	return tea.Prettify(s)
}

func (s GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) GoString() string {
	return s.String()
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetAccountId(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.AccountId = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetBandwithPackageId(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.BandwithPackageId = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetBytes(v float64) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.Bytes = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetCenId(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.CenId = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetEndTime(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.EndTime = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetOtherIp(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.OtherIp = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetOtherPort(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.OtherPort = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetOtherRegion(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.OtherRegion = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetPackets(v float64) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.Packets = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetProtocol(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.Protocol = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetStartTime(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.StartTime = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetThisIp(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.ThisIp = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetThisPort(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.ThisPort = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN) SetThisRegion(v string) *GetTransitRouterFlowTopNResponseBodyTransitRouterFlowTopN {
	s.ThisRegion = &v
	return s
}

type GetTransitRouterFlowTopNResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTransitRouterFlowTopNResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTransitRouterFlowTopNResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTransitRouterFlowTopNResponse) GoString() string {
	return s.String()
}

func (s *GetTransitRouterFlowTopNResponse) SetHeaders(v map[string]*string) *GetTransitRouterFlowTopNResponse {
	s.Headers = v
	return s
}

func (s *GetTransitRouterFlowTopNResponse) SetStatusCode(v int32) *GetTransitRouterFlowTopNResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTransitRouterFlowTopNResponse) SetBody(v *GetTransitRouterFlowTopNResponseBody) *GetTransitRouterFlowTopNResponse {
	s.Body = v
	return s
}

type GetVbrFlowTopNRequest struct {
	// The IDs of member accounts.
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The CEN connection ID.
	//
	// example:
	//
	// tr-attach-dnv870gmqzmb5u****
	AttachmentId *string `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. The maximum time range that you can query is 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638239092000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The CEN instance ID.
	//
	// example:
	//
	// cen-ia8kw1zjv4hyal****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The local IP address.
	//
	// example:
	//
	// 112.74.XX.XX
	CloudIp *string `json:"CloudIp,omitempty" xml:"CloudIp,omitempty"`
	// The local port.
	//
	// >  This parameter is required only if you set **GroupBy*	- to **CloudPort**.
	//
	// example:
	//
	// 443
	CloudPort *string `json:"CloudPort,omitempty" xml:"CloudPort,omitempty"`
	// The direction of the hybrid cloud traffic in the local regions or for the local IP addresses. Valid values:
	//
	// 	- **in**: traffic from a data center to Alibaba Cloud
	//
	// 	- **out**: traffic from Alibaba Cloud to a data center
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. The maximum time range that you can query is 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638239093000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The dimension for ranking hybrid cloud traffic data. The value of this parameter is case-sensitive. Valid values:
	//
	// 	- **1Tuple**: queries the rankings of hybrid cloud traffic data for the Cloud Enterprise Network (CEN) instances, CEN connections, virtual border routers (VBRs), and IP addresses.
	//
	// 	- **2Tuple**: queries the rankings of hybrid cloud traffic data for the local and remote IP addresses.
	//
	// 	- **5Tuple**: queries the rankings of hybrid cloud traffic data for the local and remote IP addresses, local and remote ports, and protocols.
	//
	// 	- **CloudPort**: queries the rankings of hybrid cloud traffic data for the local ports.
	//
	// 	- **OtherPort**: queries the rankings of hybrid cloud traffic data for the remote ports.
	//
	// 	- **Protocol**: queries the rankings of hybrid cloud traffic data for the protocols.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1Tuple
	GroupBy *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// The metric for ranking hybrid cloud traffic data. Default value: Bytes. This value specifies that hybrid cloud traffic data is ranked by traffic volumes.
	//
	// example:
	//
	// Bytes
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The remote IP address.
	//
	// example:
	//
	// 122.112.XX.XX
	OtherIp *string `json:"OtherIp,omitempty" xml:"OtherIp,omitempty"`
	// The remote port.
	//
	// >  This parameter is required only if you set **GroupBy*	- to **OtherPort**.
	//
	// example:
	//
	// 40002
	OtherPort *string `json:"OtherPort,omitempty" xml:"OtherPort,omitempty"`
	// The protocol number.
	//
	// >  All protocols are supported. This parameter is required only if you set **GroupBy*	- to **5Tuple*	- or **Protocol**.
	//
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The local region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The order for ranking hybrid cloud traffic data. Valid values:
	//
	// 	- **desc**: descending order
	//
	// 	- **asc**: ascending order
	//
	// example:
	//
	// desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Specifies top-N traffic data to display. Default value: **10**. This value specifies that top-10 traffic data is displayed by default. Maximum value: **100**.
	//
	// example:
	//
	// 10
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
	// Specifies whether to enable the multi-account management feature. Default value: **false**. This value specifies that the multi-account management feature is disabled.
	//
	// >  By default, the multi-account management feature is not available. If you want to use this feature, contact your account manager to apply for permissions.
	//
	// example:
	//
	// false
	UseMultiAccount *bool `json:"UseMultiAccount,omitempty" xml:"UseMultiAccount,omitempty"`
	// The ID of the VBR that is associated with the Express Connect circuit.
	//
	// example:
	//
	// vbr-k1atj46citwuek42j****
	VirtualBorderRouterId *string `json:"VirtualBorderRouterId,omitempty" xml:"VirtualBorderRouterId,omitempty"`
}

func (s GetVbrFlowTopNRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVbrFlowTopNRequest) GoString() string {
	return s.String()
}

func (s *GetVbrFlowTopNRequest) SetAccountIds(v []*int64) *GetVbrFlowTopNRequest {
	s.AccountIds = v
	return s
}

func (s *GetVbrFlowTopNRequest) SetAttachmentId(v string) *GetVbrFlowTopNRequest {
	s.AttachmentId = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetBeginTime(v int64) *GetVbrFlowTopNRequest {
	s.BeginTime = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetCenId(v string) *GetVbrFlowTopNRequest {
	s.CenId = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetCloudIp(v string) *GetVbrFlowTopNRequest {
	s.CloudIp = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetCloudPort(v string) *GetVbrFlowTopNRequest {
	s.CloudPort = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetDirection(v string) *GetVbrFlowTopNRequest {
	s.Direction = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetEndTime(v int64) *GetVbrFlowTopNRequest {
	s.EndTime = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetGroupBy(v string) *GetVbrFlowTopNRequest {
	s.GroupBy = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetOrderBy(v string) *GetVbrFlowTopNRequest {
	s.OrderBy = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetOtherIp(v string) *GetVbrFlowTopNRequest {
	s.OtherIp = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetOtherPort(v string) *GetVbrFlowTopNRequest {
	s.OtherPort = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetProtocol(v string) *GetVbrFlowTopNRequest {
	s.Protocol = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetRegionId(v string) *GetVbrFlowTopNRequest {
	s.RegionId = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetSort(v string) *GetVbrFlowTopNRequest {
	s.Sort = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetTopN(v int32) *GetVbrFlowTopNRequest {
	s.TopN = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetUseMultiAccount(v bool) *GetVbrFlowTopNRequest {
	s.UseMultiAccount = &v
	return s
}

func (s *GetVbrFlowTopNRequest) SetVirtualBorderRouterId(v string) *GetVbrFlowTopNRequest {
	s.VirtualBorderRouterId = &v
	return s
}

type GetVbrFlowTopNShrinkRequest struct {
	// The IDs of member accounts.
	AccountIdsShrink *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	// The CEN connection ID.
	//
	// example:
	//
	// tr-attach-dnv870gmqzmb5u****
	AttachmentId *string `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. The maximum time range that you can query is 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638239092000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The CEN instance ID.
	//
	// example:
	//
	// cen-ia8kw1zjv4hyal****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The local IP address.
	//
	// example:
	//
	// 112.74.XX.XX
	CloudIp *string `json:"CloudIp,omitempty" xml:"CloudIp,omitempty"`
	// The local port.
	//
	// >  This parameter is required only if you set **GroupBy*	- to **CloudPort**.
	//
	// example:
	//
	// 443
	CloudPort *string `json:"CloudPort,omitempty" xml:"CloudPort,omitempty"`
	// The direction of the hybrid cloud traffic in the local regions or for the local IP addresses. Valid values:
	//
	// 	- **in**: traffic from a data center to Alibaba Cloud
	//
	// 	- **out**: traffic from Alibaba Cloud to a data center
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. The maximum time range that you can query is 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638239093000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The dimension for ranking hybrid cloud traffic data. The value of this parameter is case-sensitive. Valid values:
	//
	// 	- **1Tuple**: queries the rankings of hybrid cloud traffic data for the Cloud Enterprise Network (CEN) instances, CEN connections, virtual border routers (VBRs), and IP addresses.
	//
	// 	- **2Tuple**: queries the rankings of hybrid cloud traffic data for the local and remote IP addresses.
	//
	// 	- **5Tuple**: queries the rankings of hybrid cloud traffic data for the local and remote IP addresses, local and remote ports, and protocols.
	//
	// 	- **CloudPort**: queries the rankings of hybrid cloud traffic data for the local ports.
	//
	// 	- **OtherPort**: queries the rankings of hybrid cloud traffic data for the remote ports.
	//
	// 	- **Protocol**: queries the rankings of hybrid cloud traffic data for the protocols.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1Tuple
	GroupBy *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// The metric for ranking hybrid cloud traffic data. Default value: Bytes. This value specifies that hybrid cloud traffic data is ranked by traffic volumes.
	//
	// example:
	//
	// Bytes
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The remote IP address.
	//
	// example:
	//
	// 122.112.XX.XX
	OtherIp *string `json:"OtherIp,omitempty" xml:"OtherIp,omitempty"`
	// The remote port.
	//
	// >  This parameter is required only if you set **GroupBy*	- to **OtherPort**.
	//
	// example:
	//
	// 40002
	OtherPort *string `json:"OtherPort,omitempty" xml:"OtherPort,omitempty"`
	// The protocol number.
	//
	// >  All protocols are supported. This parameter is required only if you set **GroupBy*	- to **5Tuple*	- or **Protocol**.
	//
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The local region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The order for ranking hybrid cloud traffic data. Valid values:
	//
	// 	- **desc**: descending order
	//
	// 	- **asc**: ascending order
	//
	// example:
	//
	// desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Specifies top-N traffic data to display. Default value: **10**. This value specifies that top-10 traffic data is displayed by default. Maximum value: **100**.
	//
	// example:
	//
	// 10
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
	// Specifies whether to enable the multi-account management feature. Default value: **false**. This value specifies that the multi-account management feature is disabled.
	//
	// >  By default, the multi-account management feature is not available. If you want to use this feature, contact your account manager to apply for permissions.
	//
	// example:
	//
	// false
	UseMultiAccount *bool `json:"UseMultiAccount,omitempty" xml:"UseMultiAccount,omitempty"`
	// The ID of the VBR that is associated with the Express Connect circuit.
	//
	// example:
	//
	// vbr-k1atj46citwuek42j****
	VirtualBorderRouterId *string `json:"VirtualBorderRouterId,omitempty" xml:"VirtualBorderRouterId,omitempty"`
}

func (s GetVbrFlowTopNShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVbrFlowTopNShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetVbrFlowTopNShrinkRequest) SetAccountIdsShrink(v string) *GetVbrFlowTopNShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetAttachmentId(v string) *GetVbrFlowTopNShrinkRequest {
	s.AttachmentId = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetBeginTime(v int64) *GetVbrFlowTopNShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetCenId(v string) *GetVbrFlowTopNShrinkRequest {
	s.CenId = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetCloudIp(v string) *GetVbrFlowTopNShrinkRequest {
	s.CloudIp = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetCloudPort(v string) *GetVbrFlowTopNShrinkRequest {
	s.CloudPort = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetDirection(v string) *GetVbrFlowTopNShrinkRequest {
	s.Direction = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetEndTime(v int64) *GetVbrFlowTopNShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetGroupBy(v string) *GetVbrFlowTopNShrinkRequest {
	s.GroupBy = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetOrderBy(v string) *GetVbrFlowTopNShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetOtherIp(v string) *GetVbrFlowTopNShrinkRequest {
	s.OtherIp = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetOtherPort(v string) *GetVbrFlowTopNShrinkRequest {
	s.OtherPort = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetProtocol(v string) *GetVbrFlowTopNShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetRegionId(v string) *GetVbrFlowTopNShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetSort(v string) *GetVbrFlowTopNShrinkRequest {
	s.Sort = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetTopN(v int32) *GetVbrFlowTopNShrinkRequest {
	s.TopN = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetUseMultiAccount(v bool) *GetVbrFlowTopNShrinkRequest {
	s.UseMultiAccount = &v
	return s
}

func (s *GetVbrFlowTopNShrinkRequest) SetVirtualBorderRouterId(v string) *GetVbrFlowTopNShrinkRequest {
	s.VirtualBorderRouterId = &v
	return s
}

type GetVbrFlowTopNResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A7F0D6EC-E19E-58AC-AC9F-08036763960F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ranking result of hybrid cloud traffic data.
	VirtualBorderRouterFlowlogTopN []*GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN `json:"VirtualBorderRouterFlowlogTopN,omitempty" xml:"VirtualBorderRouterFlowlogTopN,omitempty" type:"Repeated"`
}

func (s GetVbrFlowTopNResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVbrFlowTopNResponseBody) GoString() string {
	return s.String()
}

func (s *GetVbrFlowTopNResponseBody) SetRequestId(v string) *GetVbrFlowTopNResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVbrFlowTopNResponseBody) SetVirtualBorderRouterFlowlogTopN(v []*GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) *GetVbrFlowTopNResponseBody {
	s.VirtualBorderRouterFlowlogTopN = v
	return s
}

type GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN struct {
	// The account ID.
	//
	// example:
	//
	// 156237031628****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The CEN connection ID.
	//
	// example:
	//
	// tr-attach-u6v1j3jre0fe9h****
	AttachmentId *string `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	// The total volume of traffic in the specified time range.
	//
	// example:
	//
	// 108
	Bytes *float64 `json:"Bytes,omitempty" xml:"Bytes,omitempty"`
	// The local IP address.
	//
	// example:
	//
	// 120.24.X.X
	CloudIp *string `json:"CloudIp,omitempty" xml:"CloudIp,omitempty"`
	// The local port.
	//
	// example:
	//
	// 80
	CloudPort *string `json:"CloudPort,omitempty" xml:"CloudPort,omitempty"`
	// The local region where the local IP address resides.
	//
	// example:
	//
	// cn-shanghai
	CloudRegion *string `json:"CloudRegion,omitempty" xml:"CloudRegion,omitempty"`
	// The remote IP address.
	//
	// example:
	//
	// 222.85.X.X
	OtherIp *string `json:"OtherIp,omitempty" xml:"OtherIp,omitempty"`
	// The remote port.
	//
	// example:
	//
	// 10965
	OtherPort *string `json:"OtherPort,omitempty" xml:"OtherPort,omitempty"`
	// The total number of packets in the specified time range.
	//
	// example:
	//
	// 66
	Packets *float64 `json:"Packets,omitempty" xml:"Packets,omitempty"`
	// The protocol number.
	//
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the VBR that is associated with the Express Connect circuit.
	//
	// example:
	//
	// vbr-k1atj46citwuek42j****
	VirtualBorderRouterId *string `json:"VirtualBorderRouterId,omitempty" xml:"VirtualBorderRouterId,omitempty"`
}

func (s GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) String() string {
	return tea.Prettify(s)
}

func (s GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) GoString() string {
	return s.String()
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetAccountId(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.AccountId = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetAttachmentId(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.AttachmentId = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetBytes(v float64) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.Bytes = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetCloudIp(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.CloudIp = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetCloudPort(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.CloudPort = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetCloudRegion(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.CloudRegion = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetOtherIp(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.OtherIp = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetOtherPort(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.OtherPort = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetPackets(v float64) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.Packets = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetProtocol(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.Protocol = &v
	return s
}

func (s *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN) SetVirtualBorderRouterId(v string) *GetVbrFlowTopNResponseBodyVirtualBorderRouterFlowlogTopN {
	s.VirtualBorderRouterId = &v
	return s
}

type GetVbrFlowTopNResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVbrFlowTopNResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVbrFlowTopNResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVbrFlowTopNResponse) GoString() string {
	return s.String()
}

func (s *GetVbrFlowTopNResponse) SetHeaders(v map[string]*string) *GetVbrFlowTopNResponse {
	s.Headers = v
	return s
}

func (s *GetVbrFlowTopNResponse) SetStatusCode(v int32) *GetVbrFlowTopNResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVbrFlowTopNResponse) SetBody(v *GetVbrFlowTopNResponseBody) *GetVbrFlowTopNResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("nis"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Initiates a task for analyzing network reachability.
//
// Description:
//
// You can call this operation to initiate a task for analyzing network reachability by specifying only the information about the source and destination. You do not need to create a network path for reachability analysis. The analysis result is not recorded in the system. If you want to record the path parameters and analysis result in the Network Intelligence Service (NIS) console, we recommend that you call the **createNetworkReachableAnalysis*	- operation.
//
// @param request - CreateAndAnalyzeNetworkPathRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAndAnalyzeNetworkPathResponse
func (client *Client) CreateAndAnalyzeNetworkPathWithOptions(request *CreateAndAnalyzeNetworkPathRequest, runtime *util.RuntimeOptions) (_result *CreateAndAnalyzeNetworkPathResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAndAnalyzeNetworkPath"),
		Version:     tea.String("2021-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAndAnalyzeNetworkPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a task for analyzing network reachability.
//
// Description:
//
// You can call this operation to initiate a task for analyzing network reachability by specifying only the information about the source and destination. You do not need to create a network path for reachability analysis. The analysis result is not recorded in the system. If you want to record the path parameters and analysis result in the Network Intelligence Service (NIS) console, we recommend that you call the **createNetworkReachableAnalysis*	- operation.
//
// @param request - CreateAndAnalyzeNetworkPathRequest
//
// @return CreateAndAnalyzeNetworkPathResponse
func (client *Client) CreateAndAnalyzeNetworkPath(request *CreateAndAnalyzeNetworkPathRequest) (_result *CreateAndAnalyzeNetworkPathResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAndAnalyzeNetworkPathResponse{}
	_body, _err := client.CreateAndAnalyzeNetworkPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a network path in the cloud for reachability analysis.
//
// Description:
//
//   You can call the **CreateNetworkPath*	- operation to create network paths in multiple networking scenarios and between multiple resources. After a path is created, the path parameters are saved for repeated analysis.
//
// 	- You can create up to 100 network paths within one Alibaba Cloud account.
//
// @param request - CreateNetworkPathRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkPathResponse
func (client *Client) CreateNetworkPathWithOptions(request *CreateNetworkPathRequest, runtime *util.RuntimeOptions) (_result *CreateNetworkPathResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NetworkPathDescription)) {
		query["NetworkPathDescription"] = request.NetworkPathDescription
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkPathName)) {
		query["NetworkPathName"] = request.NetworkPathName
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["SourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIpAddress)) {
		query["SourceIpAddress"] = request.SourceIpAddress
	}

	if !tea.BoolValue(util.IsUnset(request.SourcePort)) {
		query["SourcePort"] = request.SourcePort
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetIpAddress)) {
		query["TargetIpAddress"] = request.TargetIpAddress
	}

	if !tea.BoolValue(util.IsUnset(request.TargetPort)) {
		query["TargetPort"] = request.TargetPort
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNetworkPath"),
		Version:     tea.String("2021-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNetworkPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a network path in the cloud for reachability analysis.
//
// Description:
//
//   You can call the **CreateNetworkPath*	- operation to create network paths in multiple networking scenarios and between multiple resources. After a path is created, the path parameters are saved for repeated analysis.
//
// 	- You can create up to 100 network paths within one Alibaba Cloud account.
//
// @param request - CreateNetworkPathRequest
//
// @return CreateNetworkPathResponse
func (client *Client) CreateNetworkPath(request *CreateNetworkPathRequest) (_result *CreateNetworkPathResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNetworkPathResponse{}
	_body, _err := client.CreateNetworkPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a task for analyzing network reachability.
//
// Description:
//
//   The **CreateNetworkReachableAnalysis*	- operation is used to create a task for analyzing the reachability of the network path that is created by calling the **CreateNetworkPath*	- operation and record the analysis results.
//
// 	- The **CreateNetworkReachableAnalysis*	- operation can be called to repeatedly analyze the reachability of a network path.
//
// 	- You can create up to 1,000 reachability analysis records within one Alibaba Cloud account.
//
// @param request - CreateNetworkReachableAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkReachableAnalysisResponse
func (client *Client) CreateNetworkReachableAnalysisWithOptions(request *CreateNetworkReachableAnalysisRequest, runtime *util.RuntimeOptions) (_result *CreateNetworkReachableAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NetworkPathId)) {
		query["NetworkPathId"] = request.NetworkPathId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNetworkReachableAnalysis"),
		Version:     tea.String("2021-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNetworkReachableAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task for analyzing network reachability.
//
// Description:
//
//   The **CreateNetworkReachableAnalysis*	- operation is used to create a task for analyzing the reachability of the network path that is created by calling the **CreateNetworkPath*	- operation and record the analysis results.
//
// 	- The **CreateNetworkReachableAnalysis*	- operation can be called to repeatedly analyze the reachability of a network path.
//
// 	- You can create up to 1,000 reachability analysis records within one Alibaba Cloud account.
//
// @param request - CreateNetworkReachableAnalysisRequest
//
// @return CreateNetworkReachableAnalysisResponse
func (client *Client) CreateNetworkReachableAnalysis(request *CreateNetworkReachableAnalysisRequest) (_result *CreateNetworkReachableAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNetworkReachableAnalysisResponse{}
	_body, _err := client.CreateNetworkReachableAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a network path.
//
// @param tmpReq - DeleteNetworkPathRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkPathResponse
func (client *Client) DeleteNetworkPathWithOptions(tmpReq *DeleteNetworkPathRequest, runtime *util.RuntimeOptions) (_result *DeleteNetworkPathResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteNetworkPathShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NetworkPathIds)) {
		request.NetworkPathIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetworkPathIds, tea.String("NetworkPathIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NetworkPathIdsShrink)) {
		query["NetworkPathIds"] = request.NetworkPathIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNetworkPath"),
		Version:     tea.String("2021-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNetworkPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a network path.
//
// @param request - DeleteNetworkPathRequest
//
// @return DeleteNetworkPathResponse
func (client *Client) DeleteNetworkPath(request *DeleteNetworkPathRequest) (_result *DeleteNetworkPathResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNetworkPathResponse{}
	_body, _err := client.DeleteNetworkPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a task for analyzing network reachability.
//
// @param tmpReq - DeleteNetworkReachableAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkReachableAnalysisResponse
func (client *Client) DeleteNetworkReachableAnalysisWithOptions(tmpReq *DeleteNetworkReachableAnalysisRequest, runtime *util.RuntimeOptions) (_result *DeleteNetworkReachableAnalysisResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteNetworkReachableAnalysisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NetworkReachableAnalysisIds)) {
		request.NetworkReachableAnalysisIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetworkReachableAnalysisIds, tea.String("NetworkReachableAnalysisIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NetworkReachableAnalysisIdsShrink)) {
		query["NetworkReachableAnalysisIds"] = request.NetworkReachableAnalysisIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNetworkReachableAnalysis"),
		Version:     tea.String("2021-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNetworkReachableAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a task for analyzing network reachability.
//
// @param request - DeleteNetworkReachableAnalysisRequest
//
// @return DeleteNetworkReachableAnalysisResponse
func (client *Client) DeleteNetworkReachableAnalysis(request *DeleteNetworkReachableAnalysisRequest) (_result *DeleteNetworkReachableAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNetworkReachableAnalysisResponse{}
	_body, _err := client.DeleteNetworkReachableAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the rankings of Internet traffic data in the form of 1-tuple, 2-tuple, or 5-tuple. Internet traffic data can be ranked by metrics such as traffic volumes and the number of packets.
//
// @param tmpReq - GetInternetTupleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInternetTupleResponse
func (client *Client) GetInternetTupleWithOptions(tmpReq *GetInternetTupleRequest, runtime *util.RuntimeOptions) (_result *GetInternetTupleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetInternetTupleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CloudIpList)) {
		request.CloudIpListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CloudIpList, tea.String("CloudIpList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceList)) {
		request.InstanceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceList, tea.String("InstanceList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountIds)) {
		query["AccountIds"] = request.AccountIds
	}

	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.CloudIp)) {
		query["CloudIp"] = request.CloudIp
	}

	if !tea.BoolValue(util.IsUnset(request.CloudIpListShrink)) {
		query["CloudIpList"] = request.CloudIpListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CloudIsp)) {
		query["CloudIsp"] = request.CloudIsp
	}

	if !tea.BoolValue(util.IsUnset(request.CloudPort)) {
		query["CloudPort"] = request.CloudPort
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceListShrink)) {
		query["InstanceList"] = request.InstanceListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OtherCity)) {
		query["OtherCity"] = request.OtherCity
	}

	if !tea.BoolValue(util.IsUnset(request.OtherCountry)) {
		query["OtherCountry"] = request.OtherCountry
	}

	if !tea.BoolValue(util.IsUnset(request.OtherIp)) {
		query["OtherIp"] = request.OtherIp
	}

	if !tea.BoolValue(util.IsUnset(request.OtherIsp)) {
		query["OtherIsp"] = request.OtherIsp
	}

	if !tea.BoolValue(util.IsUnset(request.OtherPort)) {
		query["OtherPort"] = request.OtherPort
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.TopN)) {
		query["TopN"] = request.TopN
	}

	if !tea.BoolValue(util.IsUnset(request.TupleType)) {
		query["TupleType"] = request.TupleType
	}

	if !tea.BoolValue(util.IsUnset(request.UseMultiAccount)) {
		query["UseMultiAccount"] = request.UseMultiAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInternetTuple"),
		Version:     tea.String("2021-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInternetTupleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the rankings of Internet traffic data in the form of 1-tuple, 2-tuple, or 5-tuple. Internet traffic data can be ranked by metrics such as traffic volumes and the number of packets.
//
// @param request - GetInternetTupleRequest
//
// @return GetInternetTupleResponse
func (client *Client) GetInternetTuple(request *GetInternetTupleRequest) (_result *GetInternetTupleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInternetTupleResponse{}
	_body, _err := client.GetInternetTupleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the real-time SNAT performance ranking of a NAT gateway.
//
// @param request - GetNatTopNRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNatTopNResponse
func (client *Client) GetNatTopNWithOptions(request *GetNatTopNRequest, runtime *util.RuntimeOptions) (_result *GetNatTopNResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.NatGatewayId)) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TopN)) {
		query["TopN"] = request.TopN
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNatTopN"),
		Version:     tea.String("2021-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNatTopNResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time SNAT performance ranking of a NAT gateway.
//
// @param request - GetNatTopNRequest
//
// @return GetNatTopNResponse
func (client *Client) GetNatTopN(request *GetNatTopNRequest) (_result *GetNatTopNResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNatTopNResponse{}
	_body, _err := client.GetNatTopNWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the results of network reachability analysis.
//
// Description:
//
// *GetNetworkReachableAnalysis*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can query the state of the task for analyzing network reachability.
//
// 	- The **init*	- state indicates that the task is in progress.
//
// 	- The **finish*	- state indicates that the task is complete. In this state, you can obtain the analysis result.
//
// @param request - GetNetworkReachableAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkReachableAnalysisResponse
func (client *Client) GetNetworkReachableAnalysisWithOptions(request *GetNetworkReachableAnalysisRequest, runtime *util.RuntimeOptions) (_result *GetNetworkReachableAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NetworkReachableAnalysisId)) {
		query["NetworkReachableAnalysisId"] = request.NetworkReachableAnalysisId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNetworkReachableAnalysis"),
		Version:     tea.String("2021-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNetworkReachableAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the results of network reachability analysis.
//
// Description:
//
// *GetNetworkReachableAnalysis*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can query the state of the task for analyzing network reachability.
//
// 	- The **init*	- state indicates that the task is in progress.
//
// 	- The **finish*	- state indicates that the task is complete. In this state, you can obtain the analysis result.
//
// @param request - GetNetworkReachableAnalysisRequest
//
// @return GetNetworkReachableAnalysisResponse
func (client *Client) GetNetworkReachableAnalysis(request *GetNetworkReachableAnalysisRequest) (_result *GetNetworkReachableAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNetworkReachableAnalysisResponse{}
	_body, _err := client.GetNetworkReachableAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取云网络指标趋势
//
// @param tmpReq - GetNisNetworkMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNisNetworkMetricsResponse
func (client *Client) GetNisNetworkMetricsWithOptions(tmpReq *GetNisNetworkMetricsRequest, runtime *util.RuntimeOptions) (_result *GetNisNetworkMetricsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetNisNetworkMetricsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Dimensions)) {
		request.DimensionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Dimensions, tea.String("Dimensions"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountIds)) {
		query["AccountIds"] = request.AccountIds
	}

	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.DimensionsShrink)) {
		query["Dimensions"] = request.DimensionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ScanBy)) {
		query["ScanBy"] = request.ScanBy
	}

	if !tea.BoolValue(util.IsUnset(request.UseCrossAccount)) {
		query["UseCrossAccount"] = request.UseCrossAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNisNetworkMetrics"),
		Version:     tea.String("2021-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNisNetworkMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取云网络指标趋势
//
// @param request - GetNisNetworkMetricsRequest
//
// @return GetNisNetworkMetricsResponse
func (client *Client) GetNisNetworkMetrics(request *GetNisNetworkMetricsRequest) (_result *GetNisNetworkMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNisNetworkMetricsResponse{}
	_body, _err := client.GetNisNetworkMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取云网络指标排名
//
// @param tmpReq - GetNisNetworkRankingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNisNetworkRankingResponse
func (client *Client) GetNisNetworkRankingWithOptions(tmpReq *GetNisNetworkRankingRequest, runtime *util.RuntimeOptions) (_result *GetNisNetworkRankingResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetNisNetworkRankingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Filter)) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, tea.String("Filter"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountIds)) {
		query["AccountIds"] = request.AccountIds
	}

	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FilterShrink)) {
		query["Filter"] = request.FilterShrink
	}

	if !tea.BoolValue(util.IsUnset(request.GroupBy)) {
		query["GroupBy"] = request.GroupBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.TopN)) {
		query["TopN"] = request.TopN
	}

	if !tea.BoolValue(util.IsUnset(request.UseCrossAccount)) {
		query["UseCrossAccount"] = request.UseCrossAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNisNetworkRanking"),
		Version:     tea.String("2021-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNisNetworkRankingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取云网络指标排名
//
// @param request - GetNisNetworkRankingRequest
//
// @return GetNisNetworkRankingResponse
func (client *Client) GetNisNetworkRanking(request *GetNisNetworkRankingRequest) (_result *GetNisNetworkRankingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNisNetworkRankingResponse{}
	_body, _err := client.GetNisNetworkRankingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the rankings of inter-region traffic data in the form of 1-tuple, 2-tuple, or 5-tuple. Inter-region traffic data can be ranked by metrics such as traffic volumes and the number of packets.
//
// @param tmpReq - GetTransitRouterFlowTopNRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTransitRouterFlowTopNResponse
func (client *Client) GetTransitRouterFlowTopNWithOptions(tmpReq *GetTransitRouterFlowTopNRequest, runtime *util.RuntimeOptions) (_result *GetTransitRouterFlowTopNResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetTransitRouterFlowTopNShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AccountIds)) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, tea.String("AccountIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountIdsShrink)) {
		query["AccountIds"] = request.AccountIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.BandwithPackageId)) {
		query["BandwithPackageId"] = request.BandwithPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		query["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.GroupBy)) {
		query["GroupBy"] = request.GroupBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OtherIp)) {
		query["OtherIp"] = request.OtherIp
	}

	if !tea.BoolValue(util.IsUnset(request.OtherPort)) {
		query["OtherPort"] = request.OtherPort
	}

	if !tea.BoolValue(util.IsUnset(request.OtherRegion)) {
		query["OtherRegion"] = request.OtherRegion
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.ThisIp)) {
		query["ThisIp"] = request.ThisIp
	}

	if !tea.BoolValue(util.IsUnset(request.ThisPort)) {
		query["ThisPort"] = request.ThisPort
	}

	if !tea.BoolValue(util.IsUnset(request.ThisRegion)) {
		query["ThisRegion"] = request.ThisRegion
	}

	if !tea.BoolValue(util.IsUnset(request.TopN)) {
		query["TopN"] = request.TopN
	}

	if !tea.BoolValue(util.IsUnset(request.UseMultiAccount)) {
		query["UseMultiAccount"] = request.UseMultiAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTransitRouterFlowTopN"),
		Version:     tea.String("2021-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTransitRouterFlowTopNResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the rankings of inter-region traffic data in the form of 1-tuple, 2-tuple, or 5-tuple. Inter-region traffic data can be ranked by metrics such as traffic volumes and the number of packets.
//
// @param request - GetTransitRouterFlowTopNRequest
//
// @return GetTransitRouterFlowTopNResponse
func (client *Client) GetTransitRouterFlowTopN(request *GetTransitRouterFlowTopNRequest) (_result *GetTransitRouterFlowTopNResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTransitRouterFlowTopNResponse{}
	_body, _err := client.GetTransitRouterFlowTopNWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the rankings of hybrid cloud traffic data in the form of 1-tuple, 2-tuple, or 5-tuple. Hybrid cloud traffic data can be ranked by metrics such as traffic volumes and the number of packets.
//
// @param tmpReq - GetVbrFlowTopNRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVbrFlowTopNResponse
func (client *Client) GetVbrFlowTopNWithOptions(tmpReq *GetVbrFlowTopNRequest, runtime *util.RuntimeOptions) (_result *GetVbrFlowTopNResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetVbrFlowTopNShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AccountIds)) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, tea.String("AccountIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountIdsShrink)) {
		query["AccountIds"] = request.AccountIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AttachmentId)) {
		query["AttachmentId"] = request.AttachmentId
	}

	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		query["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.CloudIp)) {
		query["CloudIp"] = request.CloudIp
	}

	if !tea.BoolValue(util.IsUnset(request.CloudPort)) {
		query["CloudPort"] = request.CloudPort
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.GroupBy)) {
		query["GroupBy"] = request.GroupBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OtherIp)) {
		query["OtherIp"] = request.OtherIp
	}

	if !tea.BoolValue(util.IsUnset(request.OtherPort)) {
		query["OtherPort"] = request.OtherPort
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.TopN)) {
		query["TopN"] = request.TopN
	}

	if !tea.BoolValue(util.IsUnset(request.UseMultiAccount)) {
		query["UseMultiAccount"] = request.UseMultiAccount
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualBorderRouterId)) {
		query["VirtualBorderRouterId"] = request.VirtualBorderRouterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVbrFlowTopN"),
		Version:     tea.String("2021-12-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVbrFlowTopNResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the rankings of hybrid cloud traffic data in the form of 1-tuple, 2-tuple, or 5-tuple. Hybrid cloud traffic data can be ranked by metrics such as traffic volumes and the number of packets.
//
// @param request - GetVbrFlowTopNRequest
//
// @return GetVbrFlowTopNResponse
func (client *Client) GetVbrFlowTopN(request *GetVbrFlowTopNRequest) (_result *GetVbrFlowTopNResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVbrFlowTopNResponse{}
	_body, _err := client.GetVbrFlowTopNWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
