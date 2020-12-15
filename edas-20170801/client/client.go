// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	roautil "github.com/alibabacloud-go/tea-roa-utils/service"
	roa "github.com/alibabacloud-go/tea-roa/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateModuleQuery struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty" require:"true"`
}

func (s CreateModuleQuery) String() string {
	return tea.Prettify(s)
}

func (s CreateModuleQuery) GoString() string {
	return s.String()
}

func (s *CreateModuleQuery) SetClusterId(v string) *CreateModuleQuery {
	s.ClusterId = &v
	return s
}

func (s *CreateModuleQuery) SetModuleName(v string) *CreateModuleQuery {
	s.ModuleName = &v
	return s
}

type CreateModuleBody struct {
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s CreateModuleBody) String() string {
	return tea.Prettify(s)
}

func (s CreateModuleBody) GoString() string {
	return s.String()
}

func (s *CreateModuleBody) SetValues(v string) *CreateModuleBody {
	s.Values = &v
	return s
}

type CreateModuleRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *CreateModuleQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
	Body    *CreateModuleBody  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateModuleRequest) GoString() string {
	return s.String()
}

func (s *CreateModuleRequest) SetHeaders(v map[string]*string) *CreateModuleRequest {
	s.Headers = v
	return s
}

func (s *CreateModuleRequest) SetQuery(v *CreateModuleQuery) *CreateModuleRequest {
	s.Query = v
	return s
}

func (s *CreateModuleRequest) SetBody(v *CreateModuleBody) *CreateModuleRequest {
	s.Body = v
	return s
}

type CreateModuleResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty" require:"true"`
}

func (s CreateModuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateModuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModuleResponseBody) SetCode(v int) *CreateModuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateModuleResponseBody) SetMessage(v string) *CreateModuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateModuleResponseBody) SetRequestId(v string) *CreateModuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModuleResponseBody) SetResult(v string) *CreateModuleResponseBody {
	s.Result = &v
	return s
}

type CreateModuleResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateModuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateModuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateModuleResponse) GoString() string {
	return s.String()
}

func (s *CreateModuleResponse) SetHeaders(v map[string]*string) *CreateModuleResponse {
	s.Headers = v
	return s
}

func (s *CreateModuleResponse) SetBody(v *CreateModuleResponseBody) *CreateModuleResponse {
	s.Body = v
	return s
}

type UpdateK8sApplicationBaseInfoQuery struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Owner       *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s UpdateK8sApplicationBaseInfoQuery) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sApplicationBaseInfoQuery) GoString() string {
	return s.String()
}

func (s *UpdateK8sApplicationBaseInfoQuery) SetAppId(v string) *UpdateK8sApplicationBaseInfoQuery {
	s.AppId = &v
	return s
}

func (s *UpdateK8sApplicationBaseInfoQuery) SetDescription(v string) *UpdateK8sApplicationBaseInfoQuery {
	s.Description = &v
	return s
}

func (s *UpdateK8sApplicationBaseInfoQuery) SetOwner(v string) *UpdateK8sApplicationBaseInfoQuery {
	s.Owner = &v
	return s
}

type UpdateK8sApplicationBaseInfoRequest struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *UpdateK8sApplicationBaseInfoQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s UpdateK8sApplicationBaseInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sApplicationBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateK8sApplicationBaseInfoRequest) SetHeaders(v map[string]*string) *UpdateK8sApplicationBaseInfoRequest {
	s.Headers = v
	return s
}

func (s *UpdateK8sApplicationBaseInfoRequest) SetQuery(v *UpdateK8sApplicationBaseInfoQuery) *UpdateK8sApplicationBaseInfoRequest {
	s.Query = v
	return s
}

type UpdateK8sApplicationBaseInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty" require:"true"`
}

func (s UpdateK8sApplicationBaseInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sApplicationBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateK8sApplicationBaseInfoResponseBody) SetRequestId(v string) *UpdateK8sApplicationBaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateK8sApplicationBaseInfoResponseBody) SetCode(v int) *UpdateK8sApplicationBaseInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateK8sApplicationBaseInfoResponseBody) SetMessage(v string) *UpdateK8sApplicationBaseInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateK8sApplicationBaseInfoResponseBody) SetResult(v string) *UpdateK8sApplicationBaseInfoResponseBody {
	s.Result = &v
	return s
}

type UpdateK8sApplicationBaseInfoResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateK8sApplicationBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateK8sApplicationBaseInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sApplicationBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateK8sApplicationBaseInfoResponse) SetHeaders(v map[string]*string) *UpdateK8sApplicationBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateK8sApplicationBaseInfoResponse) SetBody(v *UpdateK8sApplicationBaseInfoResponseBody) *UpdateK8sApplicationBaseInfoResponse {
	s.Body = v
	return s
}

type ScaleoutApplicationWithNewInstancesQuery struct {
	AppId                    *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	GroupId                  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	ScalingNum               *int    `json:"ScalingNum,omitempty" xml:"ScalingNum,omitempty" require:"true"`
	TemplateId               *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateVersion          *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	ScalingPolicy            *string `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty"`
	TemplateInstanceId       *string `json:"TemplateInstanceId,omitempty" xml:"TemplateInstanceId,omitempty"`
	ClusterId                *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	InstanceChargeType       *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceChargePeriodUnit *string `json:"InstanceChargePeriodUnit,omitempty" xml:"InstanceChargePeriodUnit,omitempty"`
	InstanceChargePeriod     *int    `json:"InstanceChargePeriod,omitempty" xml:"InstanceChargePeriod,omitempty"`
	AutoRenew                *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod          *int    `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
}

func (s ScaleoutApplicationWithNewInstancesQuery) String() string {
	return tea.Prettify(s)
}

func (s ScaleoutApplicationWithNewInstancesQuery) GoString() string {
	return s.String()
}

func (s *ScaleoutApplicationWithNewInstancesQuery) SetAppId(v string) *ScaleoutApplicationWithNewInstancesQuery {
	s.AppId = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesQuery) SetGroupId(v string) *ScaleoutApplicationWithNewInstancesQuery {
	s.GroupId = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesQuery) SetScalingNum(v int) *ScaleoutApplicationWithNewInstancesQuery {
	s.ScalingNum = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesQuery) SetTemplateId(v string) *ScaleoutApplicationWithNewInstancesQuery {
	s.TemplateId = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesQuery) SetTemplateVersion(v string) *ScaleoutApplicationWithNewInstancesQuery {
	s.TemplateVersion = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesQuery) SetScalingPolicy(v string) *ScaleoutApplicationWithNewInstancesQuery {
	s.ScalingPolicy = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesQuery) SetTemplateInstanceId(v string) *ScaleoutApplicationWithNewInstancesQuery {
	s.TemplateInstanceId = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesQuery) SetClusterId(v string) *ScaleoutApplicationWithNewInstancesQuery {
	s.ClusterId = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesQuery) SetInstanceChargeType(v string) *ScaleoutApplicationWithNewInstancesQuery {
	s.InstanceChargeType = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesQuery) SetInstanceChargePeriodUnit(v string) *ScaleoutApplicationWithNewInstancesQuery {
	s.InstanceChargePeriodUnit = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesQuery) SetInstanceChargePeriod(v int) *ScaleoutApplicationWithNewInstancesQuery {
	s.InstanceChargePeriod = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesQuery) SetAutoRenew(v bool) *ScaleoutApplicationWithNewInstancesQuery {
	s.AutoRenew = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesQuery) SetAutoRenewPeriod(v int) *ScaleoutApplicationWithNewInstancesQuery {
	s.AutoRenewPeriod = &v
	return s
}

type ScaleoutApplicationWithNewInstancesRequest struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ScaleoutApplicationWithNewInstancesQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ScaleoutApplicationWithNewInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleoutApplicationWithNewInstancesRequest) GoString() string {
	return s.String()
}

func (s *ScaleoutApplicationWithNewInstancesRequest) SetHeaders(v map[string]*string) *ScaleoutApplicationWithNewInstancesRequest {
	s.Headers = v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesRequest) SetQuery(v *ScaleoutApplicationWithNewInstancesQuery) *ScaleoutApplicationWithNewInstancesRequest {
	s.Query = v
	return s
}

type ScaleoutApplicationWithNewInstancesResponseBody struct {
	RequestId     *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code          *int      `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string   `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ChangeOrderId *string   `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	InstanceIds   []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" require:"true" type:"Repeated"`
}

func (s ScaleoutApplicationWithNewInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleoutApplicationWithNewInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleoutApplicationWithNewInstancesResponseBody) SetRequestId(v string) *ScaleoutApplicationWithNewInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesResponseBody) SetCode(v int) *ScaleoutApplicationWithNewInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesResponseBody) SetMessage(v string) *ScaleoutApplicationWithNewInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesResponseBody) SetChangeOrderId(v string) *ScaleoutApplicationWithNewInstancesResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesResponseBody) SetInstanceIds(v []*string) *ScaleoutApplicationWithNewInstancesResponseBody {
	s.InstanceIds = v
	return s
}

type ScaleoutApplicationWithNewInstancesResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ScaleoutApplicationWithNewInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScaleoutApplicationWithNewInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleoutApplicationWithNewInstancesResponse) GoString() string {
	return s.String()
}

func (s *ScaleoutApplicationWithNewInstancesResponse) SetHeaders(v map[string]*string) *ScaleoutApplicationWithNewInstancesResponse {
	s.Headers = v
	return s
}

func (s *ScaleoutApplicationWithNewInstancesResponse) SetBody(v *ScaleoutApplicationWithNewInstancesResponseBody) *ScaleoutApplicationWithNewInstancesResponse {
	s.Body = v
	return s
}

type GetK8sClusterQuery struct {
	RegionTag   *string `json:"RegionTag,omitempty" xml:"RegionTag,omitempty" require:"true"`
	CurrentPage *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ClusterType *int    `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
}

func (s GetK8sClusterQuery) String() string {
	return tea.Prettify(s)
}

func (s GetK8sClusterQuery) GoString() string {
	return s.String()
}

func (s *GetK8sClusterQuery) SetRegionTag(v string) *GetK8sClusterQuery {
	s.RegionTag = &v
	return s
}

func (s *GetK8sClusterQuery) SetCurrentPage(v int) *GetK8sClusterQuery {
	s.CurrentPage = &v
	return s
}

func (s *GetK8sClusterQuery) SetPageSize(v int) *GetK8sClusterQuery {
	s.PageSize = &v
	return s
}

func (s *GetK8sClusterQuery) SetClusterType(v int) *GetK8sClusterQuery {
	s.ClusterType = &v
	return s
}

type GetK8sClusterRequest struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *GetK8sClusterQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s GetK8sClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetK8sClusterRequest) GoString() string {
	return s.String()
}

func (s *GetK8sClusterRequest) SetHeaders(v map[string]*string) *GetK8sClusterRequest {
	s.Headers = v
	return s
}

func (s *GetK8sClusterRequest) SetQuery(v *GetK8sClusterQuery) *GetK8sClusterRequest {
	s.Query = v
	return s
}

type GetK8sClusterResponseBody struct {
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code        *int                                  `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message     *string                               `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ClusterPage *GetK8sClusterResponseBodyClusterPage `json:"ClusterPage,omitempty" xml:"ClusterPage,omitempty" require:"true" type:"Struct"`
}

func (s GetK8sClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetK8sClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetK8sClusterResponseBody) SetRequestId(v string) *GetK8sClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetK8sClusterResponseBody) SetCode(v int) *GetK8sClusterResponseBody {
	s.Code = &v
	return s
}

func (s *GetK8sClusterResponseBody) SetMessage(v string) *GetK8sClusterResponseBody {
	s.Message = &v
	return s
}

func (s *GetK8sClusterResponseBody) SetClusterPage(v *GetK8sClusterResponseBodyClusterPage) *GetK8sClusterResponseBody {
	s.ClusterPage = v
	return s
}

type GetK8sClusterResponseBodyClusterPage struct {
	CurrentPage *int                                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize    *int                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalSize   *int                                             `json:"TotalSize,omitempty" xml:"TotalSize,omitempty" require:"true"`
	ClusterList *GetK8sClusterResponseBodyClusterPageClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" require:"true" type:"Struct"`
}

func (s GetK8sClusterResponseBodyClusterPage) String() string {
	return tea.Prettify(s)
}

func (s GetK8sClusterResponseBodyClusterPage) GoString() string {
	return s.String()
}

func (s *GetK8sClusterResponseBodyClusterPage) SetCurrentPage(v int) *GetK8sClusterResponseBodyClusterPage {
	s.CurrentPage = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPage) SetPageSize(v int) *GetK8sClusterResponseBodyClusterPage {
	s.PageSize = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPage) SetTotalSize(v int) *GetK8sClusterResponseBodyClusterPage {
	s.TotalSize = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPage) SetClusterList(v *GetK8sClusterResponseBodyClusterPageClusterList) *GetK8sClusterResponseBodyClusterPage {
	s.ClusterList = v
	return s
}

type GetK8sClusterResponseBodyClusterPageClusterList struct {
	Cluster []*GetK8sClusterResponseBodyClusterPageClusterListCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" require:"true" type:"Repeated"`
}

func (s GetK8sClusterResponseBodyClusterPageClusterList) String() string {
	return tea.Prettify(s)
}

func (s GetK8sClusterResponseBodyClusterPageClusterList) GoString() string {
	return s.String()
}

func (s *GetK8sClusterResponseBodyClusterPageClusterList) SetCluster(v []*GetK8sClusterResponseBodyClusterPageClusterListCluster) *GetK8sClusterResponseBodyClusterPageClusterList {
	s.Cluster = v
	return s
}

type GetK8sClusterResponseBodyClusterPageClusterListCluster struct {
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	ClusterImportStatus *int    `json:"ClusterImportStatus,omitempty" xml:"ClusterImportStatus,omitempty" require:"true"`
	ClusterName         *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty" require:"true"`
	ClusterStatus       *int    `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty" require:"true"`
	ClusterType         *int    `json:"ClusterType,omitempty" xml:"ClusterType,omitempty" require:"true"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	VpcId               *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	VswitchId           *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty" require:"true"`
	SubNetCidr          *string `json:"SubNetCidr,omitempty" xml:"SubNetCidr,omitempty" require:"true"`
	CsClusterStatus     *string `json:"CsClusterStatus,omitempty" xml:"CsClusterStatus,omitempty" require:"true"`
	CsClusterId         *string `json:"CsClusterId,omitempty" xml:"CsClusterId,omitempty" require:"true"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	NodeNum             *int    `json:"NodeNum,omitempty" xml:"NodeNum,omitempty" require:"true"`
	Cpu                 *int    `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Mem                 *int    `json:"Mem,omitempty" xml:"Mem,omitempty" require:"true"`
	NetworkMode         *int    `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty" require:"true"`
}

func (s GetK8sClusterResponseBodyClusterPageClusterListCluster) String() string {
	return tea.Prettify(s)
}

func (s GetK8sClusterResponseBodyClusterPageClusterListCluster) GoString() string {
	return s.String()
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetClusterId(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.ClusterId = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetClusterImportStatus(v int) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.ClusterImportStatus = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetClusterName(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.ClusterName = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetClusterStatus(v int) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.ClusterStatus = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetClusterType(v int) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.ClusterType = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetRegionId(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.RegionId = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetVpcId(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.VpcId = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetVswitchId(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.VswitchId = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetSubNetCidr(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.SubNetCidr = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetCsClusterStatus(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.CsClusterStatus = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetCsClusterId(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.CsClusterId = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetDescription(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.Description = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetNodeNum(v int) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.NodeNum = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetCpu(v int) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.Cpu = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetMem(v int) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.Mem = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetNetworkMode(v int) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.NetworkMode = &v
	return s
}

type GetK8sClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetK8sClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetK8sClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetK8sClusterResponse) GoString() string {
	return s.String()
}

func (s *GetK8sClusterResponse) SetHeaders(v map[string]*string) *GetK8sClusterResponse {
	s.Headers = v
	return s
}

func (s *GetK8sClusterResponse) SetBody(v *GetK8sClusterResponseBody) *GetK8sClusterResponse {
	s.Body = v
	return s
}

type CreateIDCImportCommandBody struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
}

func (s CreateIDCImportCommandBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIDCImportCommandBody) GoString() string {
	return s.String()
}

func (s *CreateIDCImportCommandBody) SetClusterId(v string) *CreateIDCImportCommandBody {
	s.ClusterId = &v
	return s
}

type CreateIDCImportCommandRequest struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *CreateIDCImportCommandBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIDCImportCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIDCImportCommandRequest) GoString() string {
	return s.String()
}

func (s *CreateIDCImportCommandRequest) SetHeaders(v map[string]*string) *CreateIDCImportCommandRequest {
	s.Headers = v
	return s
}

func (s *CreateIDCImportCommandRequest) SetBody(v *CreateIDCImportCommandBody) *CreateIDCImportCommandRequest {
	s.Body = v
	return s
}

type CreateIDCImportCommandResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s CreateIDCImportCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIDCImportCommandResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIDCImportCommandResponseBody) SetRequestId(v string) *CreateIDCImportCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIDCImportCommandResponseBody) SetCode(v string) *CreateIDCImportCommandResponseBody {
	s.Code = &v
	return s
}

func (s *CreateIDCImportCommandResponseBody) SetMessage(v string) *CreateIDCImportCommandResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIDCImportCommandResponseBody) SetData(v string) *CreateIDCImportCommandResponseBody {
	s.Data = &v
	return s
}

type CreateIDCImportCommandResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateIDCImportCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIDCImportCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIDCImportCommandResponse) GoString() string {
	return s.String()
}

func (s *CreateIDCImportCommandResponse) SetHeaders(v map[string]*string) *CreateIDCImportCommandResponse {
	s.Headers = v
	return s
}

func (s *CreateIDCImportCommandResponse) SetBody(v *CreateIDCImportCommandResponseBody) *CreateIDCImportCommandResponse {
	s.Body = v
	return s
}

type ListOperationLogsQuery struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BeginTime   *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	CurrentPage *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOperationLogsQuery) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsQuery) GoString() string {
	return s.String()
}

func (s *ListOperationLogsQuery) SetAppId(v string) *ListOperationLogsQuery {
	s.AppId = &v
	return s
}

func (s *ListOperationLogsQuery) SetBeginTime(v int64) *ListOperationLogsQuery {
	s.BeginTime = &v
	return s
}

func (s *ListOperationLogsQuery) SetEndTime(v int64) *ListOperationLogsQuery {
	s.EndTime = &v
	return s
}

func (s *ListOperationLogsQuery) SetUserId(v string) *ListOperationLogsQuery {
	s.UserId = &v
	return s
}

func (s *ListOperationLogsQuery) SetCurrentPage(v int) *ListOperationLogsQuery {
	s.CurrentPage = &v
	return s
}

func (s *ListOperationLogsQuery) SetPageSize(v int) *ListOperationLogsQuery {
	s.PageSize = &v
	return s
}

type ListOperationLogsRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListOperationLogsQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s ListOperationLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsRequest) GoString() string {
	return s.String()
}

func (s *ListOperationLogsRequest) SetHeaders(v map[string]*string) *ListOperationLogsRequest {
	s.Headers = v
	return s
}

func (s *ListOperationLogsRequest) SetQuery(v *ListOperationLogsQuery) *ListOperationLogsRequest {
	s.Query = v
	return s
}

type ListOperationLogsResponseBody struct {
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code        *int                                    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message     *string                                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	TotalSize   *int                                    `json:"TotalSize,omitempty" xml:"TotalSize,omitempty" require:"true"`
	CurrentPage *int                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize    *int                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	LogList     []*ListOperationLogsResponseBodyLogList `json:"LogList,omitempty" xml:"LogList,omitempty" require:"true" type:"Repeated"`
}

func (s ListOperationLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationLogsResponseBody) SetRequestId(v string) *ListOperationLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationLogsResponseBody) SetCode(v int) *ListOperationLogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListOperationLogsResponseBody) SetMessage(v string) *ListOperationLogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListOperationLogsResponseBody) SetTotalSize(v int) *ListOperationLogsResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListOperationLogsResponseBody) SetCurrentPage(v int) *ListOperationLogsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListOperationLogsResponseBody) SetPageSize(v int) *ListOperationLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOperationLogsResponseBody) SetLogList(v []*ListOperationLogsResponseBodyLogList) *ListOperationLogsResponseBody {
	s.LogList = v
	return s
}

type ListOperationLogsResponseBodyLogList struct {
	ActionGroup     *string `json:"ActionGroup,omitempty" xml:"ActionGroup,omitempty" require:"true"`
	ActionName      *string `json:"ActionName,omitempty" xml:"ActionName,omitempty" require:"true"`
	BeginTime       *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty" require:"true"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	OperatorName    *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty" require:"true"`
	OperatorId      *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty" require:"true"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty" require:"true"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ExtraParameters *string `json:"ExtraParameters,omitempty" xml:"ExtraParameters,omitempty" require:"true"`
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s ListOperationLogsResponseBodyLogList) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsResponseBodyLogList) GoString() string {
	return s.String()
}

func (s *ListOperationLogsResponseBodyLogList) SetActionGroup(v string) *ListOperationLogsResponseBodyLogList {
	s.ActionGroup = &v
	return s
}

func (s *ListOperationLogsResponseBodyLogList) SetActionName(v string) *ListOperationLogsResponseBodyLogList {
	s.ActionName = &v
	return s
}

func (s *ListOperationLogsResponseBodyLogList) SetBeginTime(v int64) *ListOperationLogsResponseBodyLogList {
	s.BeginTime = &v
	return s
}

func (s *ListOperationLogsResponseBodyLogList) SetEndTime(v int64) *ListOperationLogsResponseBodyLogList {
	s.EndTime = &v
	return s
}

func (s *ListOperationLogsResponseBodyLogList) SetOperatorName(v string) *ListOperationLogsResponseBodyLogList {
	s.OperatorName = &v
	return s
}

func (s *ListOperationLogsResponseBodyLogList) SetOperatorId(v string) *ListOperationLogsResponseBodyLogList {
	s.OperatorId = &v
	return s
}

func (s *ListOperationLogsResponseBodyLogList) SetSource(v string) *ListOperationLogsResponseBodyLogList {
	s.Source = &v
	return s
}

func (s *ListOperationLogsResponseBodyLogList) SetStatus(v string) *ListOperationLogsResponseBodyLogList {
	s.Status = &v
	return s
}

func (s *ListOperationLogsResponseBodyLogList) SetMessage(v string) *ListOperationLogsResponseBodyLogList {
	s.Message = &v
	return s
}

func (s *ListOperationLogsResponseBodyLogList) SetExtraParameters(v string) *ListOperationLogsResponseBodyLogList {
	s.ExtraParameters = &v
	return s
}

func (s *ListOperationLogsResponseBodyLogList) SetAppId(v string) *ListOperationLogsResponseBodyLogList {
	s.AppId = &v
	return s
}

type ListOperationLogsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOperationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOperationLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsResponse) GoString() string {
	return s.String()
}

func (s *ListOperationLogsResponse) SetHeaders(v map[string]*string) *ListOperationLogsResponse {
	s.Headers = v
	return s
}

func (s *ListOperationLogsResponse) SetBody(v *ListOperationLogsResponseBody) *ListOperationLogsResponse {
	s.Body = v
	return s
}

type ListRootStacksQuery struct {
	CurrentPage *int `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRootStacksQuery) String() string {
	return tea.Prettify(s)
}

func (s ListRootStacksQuery) GoString() string {
	return s.String()
}

func (s *ListRootStacksQuery) SetCurrentPage(v int) *ListRootStacksQuery {
	s.CurrentPage = &v
	return s
}

func (s *ListRootStacksQuery) SetPageSize(v int) *ListRootStacksQuery {
	s.PageSize = &v
	return s
}

type ListRootStacksRequest struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListRootStacksQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s ListRootStacksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRootStacksRequest) GoString() string {
	return s.String()
}

func (s *ListRootStacksRequest) SetHeaders(v map[string]*string) *ListRootStacksRequest {
	s.Headers = v
	return s
}

func (s *ListRootStacksRequest) SetQuery(v *ListRootStacksQuery) *ListRootStacksRequest {
	s.Query = v
	return s
}

type ListRootStacksResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *int                            `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *ListRootStacksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListRootStacksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRootStacksResponseBody) GoString() string {
	return s.String()
}

func (s *ListRootStacksResponseBody) SetRequestId(v string) *ListRootStacksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRootStacksResponseBody) SetMessage(v string) *ListRootStacksResponseBody {
	s.Message = &v
	return s
}

func (s *ListRootStacksResponseBody) SetCode(v int) *ListRootStacksResponseBody {
	s.Code = &v
	return s
}

func (s *ListRootStacksResponseBody) SetData(v *ListRootStacksResponseBodyData) *ListRootStacksResponseBody {
	s.Data = v
	return s
}

type ListRootStacksResponseBodyData struct {
	CurrentPage *int                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize    *int                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalSize   *int                                    `json:"TotalSize,omitempty" xml:"TotalSize,omitempty" require:"true"`
	Result      []*ListRootStacksResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" require:"true" type:"Repeated"`
}

func (s ListRootStacksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRootStacksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRootStacksResponseBodyData) SetCurrentPage(v int) *ListRootStacksResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListRootStacksResponseBodyData) SetPageSize(v int) *ListRootStacksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListRootStacksResponseBodyData) SetTotalSize(v int) *ListRootStacksResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListRootStacksResponseBodyData) SetResult(v []*ListRootStacksResponseBodyDataResult) *ListRootStacksResponseBodyData {
	s.Result = v
	return s
}

type ListRootStacksResponseBodyDataResult struct {
	Children []*ListRootStacksResponseBodyDataResultChildren `json:"Children,omitempty" xml:"Children,omitempty" require:"true" type:"Repeated"`
	Root     *ListRootStacksResponseBodyDataResultRoot       `json:"Root,omitempty" xml:"Root,omitempty" require:"true" type:"Struct"`
}

func (s ListRootStacksResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ListRootStacksResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListRootStacksResponseBodyDataResult) SetChildren(v []*ListRootStacksResponseBodyDataResultChildren) *ListRootStacksResponseBodyDataResult {
	s.Children = v
	return s
}

func (s *ListRootStacksResponseBodyDataResult) SetRoot(v *ListRootStacksResponseBodyDataResultRoot) *ListRootStacksResponseBodyDataResult {
	s.Root = v
	return s
}

type ListRootStacksResponseBodyDataResultChildren struct {
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Icon    *string `json:"Icon,omitempty" xml:"Icon,omitempty" require:"true"`
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty" require:"true"`
}

func (s ListRootStacksResponseBodyDataResultChildren) String() string {
	return tea.Prettify(s)
}

func (s ListRootStacksResponseBodyDataResultChildren) GoString() string {
	return s.String()
}

func (s *ListRootStacksResponseBodyDataResultChildren) SetId(v int64) *ListRootStacksResponseBodyDataResultChildren {
	s.Id = &v
	return s
}

func (s *ListRootStacksResponseBodyDataResultChildren) SetName(v string) *ListRootStacksResponseBodyDataResultChildren {
	s.Name = &v
	return s
}

func (s *ListRootStacksResponseBodyDataResultChildren) SetIcon(v string) *ListRootStacksResponseBodyDataResultChildren {
	s.Icon = &v
	return s
}

func (s *ListRootStacksResponseBodyDataResultChildren) SetComment(v string) *ListRootStacksResponseBodyDataResultChildren {
	s.Comment = &v
	return s
}

type ListRootStacksResponseBodyDataResultRoot struct {
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s ListRootStacksResponseBodyDataResultRoot) String() string {
	return tea.Prettify(s)
}

func (s ListRootStacksResponseBodyDataResultRoot) GoString() string {
	return s.String()
}

func (s *ListRootStacksResponseBodyDataResultRoot) SetId(v int64) *ListRootStacksResponseBodyDataResultRoot {
	s.Id = &v
	return s
}

func (s *ListRootStacksResponseBodyDataResultRoot) SetName(v string) *ListRootStacksResponseBodyDataResultRoot {
	s.Name = &v
	return s
}

type ListRootStacksResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRootStacksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRootStacksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRootStacksResponse) GoString() string {
	return s.String()
}

func (s *ListRootStacksResponse) SetHeaders(v map[string]*string) *ListRootStacksResponse {
	s.Headers = v
	return s
}

func (s *ListRootStacksResponse) SetBody(v *ListRootStacksResponseBody) *ListRootStacksResponse {
	s.Body = v
	return s
}

type ListChildrenStacksQuery struct {
	StackId     *int64 `json:"StackId,omitempty" xml:"StackId,omitempty" require:"true"`
	CurrentPage *int   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListChildrenStacksQuery) String() string {
	return tea.Prettify(s)
}

func (s ListChildrenStacksQuery) GoString() string {
	return s.String()
}

func (s *ListChildrenStacksQuery) SetStackId(v int64) *ListChildrenStacksQuery {
	s.StackId = &v
	return s
}

func (s *ListChildrenStacksQuery) SetCurrentPage(v int) *ListChildrenStacksQuery {
	s.CurrentPage = &v
	return s
}

func (s *ListChildrenStacksQuery) SetPageSize(v int) *ListChildrenStacksQuery {
	s.PageSize = &v
	return s
}

type ListChildrenStacksRequest struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListChildrenStacksQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ListChildrenStacksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChildrenStacksRequest) GoString() string {
	return s.String()
}

func (s *ListChildrenStacksRequest) SetHeaders(v map[string]*string) *ListChildrenStacksRequest {
	s.Headers = v
	return s
}

func (s *ListChildrenStacksRequest) SetQuery(v *ListChildrenStacksQuery) *ListChildrenStacksRequest {
	s.Query = v
	return s
}

type ListChildrenStacksResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                                `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *ListChildrenStacksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ListChildrenStacksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChildrenStacksResponseBody) GoString() string {
	return s.String()
}

func (s *ListChildrenStacksResponseBody) SetRequestId(v string) *ListChildrenStacksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChildrenStacksResponseBody) SetCode(v int) *ListChildrenStacksResponseBody {
	s.Code = &v
	return s
}

func (s *ListChildrenStacksResponseBody) SetMessage(v string) *ListChildrenStacksResponseBody {
	s.Message = &v
	return s
}

func (s *ListChildrenStacksResponseBody) SetData(v *ListChildrenStacksResponseBodyData) *ListChildrenStacksResponseBody {
	s.Data = v
	return s
}

type ListChildrenStacksResponseBodyData struct {
	PageSize    *int                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalSize   *int                                        `json:"TotalSize,omitempty" xml:"TotalSize,omitempty" require:"true"`
	CurrentPage *int                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	Result      []*ListChildrenStacksResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" require:"true" type:"Repeated"`
}

func (s ListChildrenStacksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListChildrenStacksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListChildrenStacksResponseBodyData) SetPageSize(v int) *ListChildrenStacksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListChildrenStacksResponseBodyData) SetTotalSize(v int) *ListChildrenStacksResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListChildrenStacksResponseBodyData) SetCurrentPage(v int) *ListChildrenStacksResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListChildrenStacksResponseBodyData) SetResult(v []*ListChildrenStacksResponseBodyDataResult) *ListChildrenStacksResponseBodyData {
	s.Result = v
	return s
}

type ListChildrenStacksResponseBodyDataResult struct {
	Id         *int64    `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	Name       *string   `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Comment    *string   `json:"Comment,omitempty" xml:"Comment,omitempty" require:"true"`
	Preferred  *bool     `json:"Preferred,omitempty" xml:"Preferred,omitempty" require:"true"`
	Latest     *bool     `json:"Latest,omitempty" xml:"Latest,omitempty" require:"true"`
	BuildTypes []*string `json:"BuildTypes,omitempty" xml:"BuildTypes,omitempty" require:"true" type:"Repeated"`
}

func (s ListChildrenStacksResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ListChildrenStacksResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListChildrenStacksResponseBodyDataResult) SetId(v int64) *ListChildrenStacksResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *ListChildrenStacksResponseBodyDataResult) SetName(v string) *ListChildrenStacksResponseBodyDataResult {
	s.Name = &v
	return s
}

func (s *ListChildrenStacksResponseBodyDataResult) SetComment(v string) *ListChildrenStacksResponseBodyDataResult {
	s.Comment = &v
	return s
}

func (s *ListChildrenStacksResponseBodyDataResult) SetPreferred(v bool) *ListChildrenStacksResponseBodyDataResult {
	s.Preferred = &v
	return s
}

func (s *ListChildrenStacksResponseBodyDataResult) SetLatest(v bool) *ListChildrenStacksResponseBodyDataResult {
	s.Latest = &v
	return s
}

func (s *ListChildrenStacksResponseBodyDataResult) SetBuildTypes(v []*string) *ListChildrenStacksResponseBodyDataResult {
	s.BuildTypes = v
	return s
}

type ListChildrenStacksResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListChildrenStacksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListChildrenStacksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChildrenStacksResponse) GoString() string {
	return s.String()
}

func (s *ListChildrenStacksResponse) SetHeaders(v map[string]*string) *ListChildrenStacksResponse {
	s.Headers = v
	return s
}

func (s *ListChildrenStacksResponse) SetBody(v *ListChildrenStacksResponseBody) *ListChildrenStacksResponse {
	s.Body = v
	return s
}

type StartK8sApplicationQuery struct {
	Replicas *int    `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	Timeout  *int    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s StartK8sApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s StartK8sApplicationQuery) GoString() string {
	return s.String()
}

func (s *StartK8sApplicationQuery) SetReplicas(v int) *StartK8sApplicationQuery {
	s.Replicas = &v
	return s
}

func (s *StartK8sApplicationQuery) SetTimeout(v int) *StartK8sApplicationQuery {
	s.Timeout = &v
	return s
}

func (s *StartK8sApplicationQuery) SetAppId(v string) *StartK8sApplicationQuery {
	s.AppId = &v
	return s
}

type StartK8sApplicationRequest struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *StartK8sApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s StartK8sApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s StartK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *StartK8sApplicationRequest) SetHeaders(v map[string]*string) *StartK8sApplicationRequest {
	s.Headers = v
	return s
}

func (s *StartK8sApplicationRequest) SetQuery(v *StartK8sApplicationQuery) *StartK8sApplicationRequest {
	s.Query = v
	return s
}

type StartK8sApplicationResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s StartK8sApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartK8sApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *StartK8sApplicationResponseBody) SetRequestId(v string) *StartK8sApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartK8sApplicationResponseBody) SetCode(v int) *StartK8sApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *StartK8sApplicationResponseBody) SetMessage(v string) *StartK8sApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *StartK8sApplicationResponseBody) SetChangeOrderId(v string) *StartK8sApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

type StartK8sApplicationResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartK8sApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartK8sApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s StartK8sApplicationResponse) GoString() string {
	return s.String()
}

func (s *StartK8sApplicationResponse) SetHeaders(v map[string]*string) *StartK8sApplicationResponse {
	s.Headers = v
	return s
}

func (s *StartK8sApplicationResponse) SetBody(v *StartK8sApplicationResponseBody) *StartK8sApplicationResponse {
	s.Body = v
	return s
}

type RestartK8sApplicationQuery struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Timeout *int    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s RestartK8sApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s RestartK8sApplicationQuery) GoString() string {
	return s.String()
}

func (s *RestartK8sApplicationQuery) SetAppId(v string) *RestartK8sApplicationQuery {
	s.AppId = &v
	return s
}

func (s *RestartK8sApplicationQuery) SetTimeout(v int) *RestartK8sApplicationQuery {
	s.Timeout = &v
	return s
}

type RestartK8sApplicationRequest struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *RestartK8sApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s RestartK8sApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *RestartK8sApplicationRequest) SetHeaders(v map[string]*string) *RestartK8sApplicationRequest {
	s.Headers = v
	return s
}

func (s *RestartK8sApplicationRequest) SetQuery(v *RestartK8sApplicationQuery) *RestartK8sApplicationRequest {
	s.Query = v
	return s
}

type RestartK8sApplicationResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s RestartK8sApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartK8sApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RestartK8sApplicationResponseBody) SetRequestId(v string) *RestartK8sApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartK8sApplicationResponseBody) SetCode(v int) *RestartK8sApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *RestartK8sApplicationResponseBody) SetMessage(v string) *RestartK8sApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *RestartK8sApplicationResponseBody) SetChangeOrderId(v string) *RestartK8sApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

type RestartK8sApplicationResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartK8sApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartK8sApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartK8sApplicationResponse) GoString() string {
	return s.String()
}

func (s *RestartK8sApplicationResponse) SetHeaders(v map[string]*string) *RestartK8sApplicationResponse {
	s.Headers = v
	return s
}

func (s *RestartK8sApplicationResponse) SetBody(v *RestartK8sApplicationResponseBody) *RestartK8sApplicationResponse {
	s.Body = v
	return s
}

type BindEcsSlbQuery struct {
	AppId                  *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	SlbId                  *string `json:"SlbId,omitempty" xml:"SlbId,omitempty" require:"true"`
	ListenerPort           *int    `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty" require:"true"`
	VServerGroupId         *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	ListenerProtocol       *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty" require:"true"`
	DeployGroupId          *string `json:"DeployGroupId,omitempty" xml:"DeployGroupId,omitempty"`
	VServerGroupName       *string `json:"VServerGroupName,omitempty" xml:"VServerGroupName,omitempty"`
	ListenerHealthCheckUrl *string `json:"ListenerHealthCheckUrl,omitempty" xml:"ListenerHealthCheckUrl,omitempty"`
	VForwardingUrlRule     *string `json:"VForwardingUrlRule,omitempty" xml:"VForwardingUrlRule,omitempty"`
}

func (s BindEcsSlbQuery) String() string {
	return tea.Prettify(s)
}

func (s BindEcsSlbQuery) GoString() string {
	return s.String()
}

func (s *BindEcsSlbQuery) SetAppId(v string) *BindEcsSlbQuery {
	s.AppId = &v
	return s
}

func (s *BindEcsSlbQuery) SetSlbId(v string) *BindEcsSlbQuery {
	s.SlbId = &v
	return s
}

func (s *BindEcsSlbQuery) SetListenerPort(v int) *BindEcsSlbQuery {
	s.ListenerPort = &v
	return s
}

func (s *BindEcsSlbQuery) SetVServerGroupId(v string) *BindEcsSlbQuery {
	s.VServerGroupId = &v
	return s
}

func (s *BindEcsSlbQuery) SetListenerProtocol(v string) *BindEcsSlbQuery {
	s.ListenerProtocol = &v
	return s
}

func (s *BindEcsSlbQuery) SetDeployGroupId(v string) *BindEcsSlbQuery {
	s.DeployGroupId = &v
	return s
}

func (s *BindEcsSlbQuery) SetVServerGroupName(v string) *BindEcsSlbQuery {
	s.VServerGroupName = &v
	return s
}

func (s *BindEcsSlbQuery) SetListenerHealthCheckUrl(v string) *BindEcsSlbQuery {
	s.ListenerHealthCheckUrl = &v
	return s
}

func (s *BindEcsSlbQuery) SetVForwardingUrlRule(v string) *BindEcsSlbQuery {
	s.VForwardingUrlRule = &v
	return s
}

type BindEcsSlbRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *BindEcsSlbQuery   `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s BindEcsSlbRequest) String() string {
	return tea.Prettify(s)
}

func (s BindEcsSlbRequest) GoString() string {
	return s.String()
}

func (s *BindEcsSlbRequest) SetHeaders(v map[string]*string) *BindEcsSlbRequest {
	s.Headers = v
	return s
}

func (s *BindEcsSlbRequest) SetQuery(v *BindEcsSlbQuery) *BindEcsSlbRequest {
	s.Query = v
	return s
}

type BindEcsSlbResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s BindEcsSlbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindEcsSlbResponseBody) GoString() string {
	return s.String()
}

func (s *BindEcsSlbResponseBody) SetRequestId(v string) *BindEcsSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindEcsSlbResponseBody) SetMessage(v string) *BindEcsSlbResponseBody {
	s.Message = &v
	return s
}

func (s *BindEcsSlbResponseBody) SetCode(v int) *BindEcsSlbResponseBody {
	s.Code = &v
	return s
}

func (s *BindEcsSlbResponseBody) SetChangeOrderId(v string) *BindEcsSlbResponseBody {
	s.ChangeOrderId = &v
	return s
}

type BindEcsSlbResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindEcsSlbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindEcsSlbResponse) String() string {
	return tea.Prettify(s)
}

func (s BindEcsSlbResponse) GoString() string {
	return s.String()
}

func (s *BindEcsSlbResponse) SetHeaders(v map[string]*string) *BindEcsSlbResponse {
	s.Headers = v
	return s
}

func (s *BindEcsSlbResponse) SetBody(v *BindEcsSlbResponseBody) *BindEcsSlbResponse {
	s.Body = v
	return s
}

type StopK8sApplicationQuery struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Timeout *int    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s StopK8sApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s StopK8sApplicationQuery) GoString() string {
	return s.String()
}

func (s *StopK8sApplicationQuery) SetAppId(v string) *StopK8sApplicationQuery {
	s.AppId = &v
	return s
}

func (s *StopK8sApplicationQuery) SetTimeout(v int) *StopK8sApplicationQuery {
	s.Timeout = &v
	return s
}

type StopK8sApplicationRequest struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *StopK8sApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s StopK8sApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s StopK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *StopK8sApplicationRequest) SetHeaders(v map[string]*string) *StopK8sApplicationRequest {
	s.Headers = v
	return s
}

func (s *StopK8sApplicationRequest) SetQuery(v *StopK8sApplicationQuery) *StopK8sApplicationRequest {
	s.Query = v
	return s
}

type StopK8sApplicationResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s StopK8sApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopK8sApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *StopK8sApplicationResponseBody) SetRequestId(v string) *StopK8sApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopK8sApplicationResponseBody) SetCode(v int) *StopK8sApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *StopK8sApplicationResponseBody) SetMessage(v string) *StopK8sApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *StopK8sApplicationResponseBody) SetChangeOrderId(v string) *StopK8sApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

type StopK8sApplicationResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopK8sApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopK8sApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s StopK8sApplicationResponse) GoString() string {
	return s.String()
}

func (s *StopK8sApplicationResponse) SetHeaders(v map[string]*string) *StopK8sApplicationResponse {
	s.Headers = v
	return s
}

func (s *StopK8sApplicationResponse) SetBody(v *StopK8sApplicationResponseBody) *StopK8sApplicationResponse {
	s.Body = v
	return s
}

type ConvertK8sResourceQuery struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty" require:"true"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty" require:"true"`
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
}

func (s ConvertK8sResourceQuery) String() string {
	return tea.Prettify(s)
}

func (s ConvertK8sResourceQuery) GoString() string {
	return s.String()
}

func (s *ConvertK8sResourceQuery) SetResourceType(v string) *ConvertK8sResourceQuery {
	s.ResourceType = &v
	return s
}

func (s *ConvertK8sResourceQuery) SetResourceName(v string) *ConvertK8sResourceQuery {
	s.ResourceName = &v
	return s
}

func (s *ConvertK8sResourceQuery) SetNamespace(v string) *ConvertK8sResourceQuery {
	s.Namespace = &v
	return s
}

func (s *ConvertK8sResourceQuery) SetClusterId(v string) *ConvertK8sResourceQuery {
	s.ClusterId = &v
	return s
}

type ConvertK8sResourceRequest struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ConvertK8sResourceQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ConvertK8sResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertK8sResourceRequest) GoString() string {
	return s.String()
}

func (s *ConvertK8sResourceRequest) SetHeaders(v map[string]*string) *ConvertK8sResourceRequest {
	s.Headers = v
	return s
}

func (s *ConvertK8sResourceRequest) SetQuery(v *ConvertK8sResourceQuery) *ConvertK8sResourceRequest {
	s.Query = v
	return s
}

type ConvertK8sResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
}

func (s ConvertK8sResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConvertK8sResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertK8sResourceResponseBody) SetRequestId(v string) *ConvertK8sResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertK8sResourceResponseBody) SetMessage(v string) *ConvertK8sResourceResponseBody {
	s.Message = &v
	return s
}

func (s *ConvertK8sResourceResponseBody) SetCode(v int) *ConvertK8sResourceResponseBody {
	s.Code = &v
	return s
}

type ConvertK8sResourceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConvertK8sResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConvertK8sResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertK8sResourceResponse) GoString() string {
	return s.String()
}

func (s *ConvertK8sResourceResponse) SetHeaders(v map[string]*string) *ConvertK8sResourceResponse {
	s.Headers = v
	return s
}

func (s *ConvertK8sResourceResponse) SetBody(v *ConvertK8sResourceResponseBody) *ConvertK8sResourceResponse {
	s.Body = v
	return s
}

type UpdateSlsLogStoreBody struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Configs *string `json:"Configs,omitempty" xml:"Configs,omitempty" require:"true"`
}

func (s UpdateSlsLogStoreBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSlsLogStoreBody) GoString() string {
	return s.String()
}

func (s *UpdateSlsLogStoreBody) SetAppId(v string) *UpdateSlsLogStoreBody {
	s.AppId = &v
	return s
}

func (s *UpdateSlsLogStoreBody) SetConfigs(v string) *UpdateSlsLogStoreBody {
	s.Configs = &v
	return s
}

type UpdateSlsLogStoreRequest struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *UpdateSlsLogStoreBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSlsLogStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSlsLogStoreRequest) GoString() string {
	return s.String()
}

func (s *UpdateSlsLogStoreRequest) SetHeaders(v map[string]*string) *UpdateSlsLogStoreRequest {
	s.Headers = v
	return s
}

func (s *UpdateSlsLogStoreRequest) SetBody(v *UpdateSlsLogStoreBody) *UpdateSlsLogStoreRequest {
	s.Body = v
	return s
}

type UpdateSlsLogStoreResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
}

func (s UpdateSlsLogStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSlsLogStoreResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSlsLogStoreResponseBody) SetRequestId(v string) *UpdateSlsLogStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSlsLogStoreResponseBody) SetMessage(v string) *UpdateSlsLogStoreResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSlsLogStoreResponseBody) SetCode(v int) *UpdateSlsLogStoreResponseBody {
	s.Code = &v
	return s
}

type UpdateSlsLogStoreResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSlsLogStoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSlsLogStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSlsLogStoreResponse) GoString() string {
	return s.String()
}

func (s *UpdateSlsLogStoreResponse) SetHeaders(v map[string]*string) *UpdateSlsLogStoreResponse {
	s.Headers = v
	return s
}

func (s *UpdateSlsLogStoreResponse) SetBody(v *UpdateSlsLogStoreResponseBody) *UpdateSlsLogStoreResponse {
	s.Body = v
	return s
}

type QueryK8sClusterLogProjectInfoQuery struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
}

func (s QueryK8sClusterLogProjectInfoQuery) String() string {
	return tea.Prettify(s)
}

func (s QueryK8sClusterLogProjectInfoQuery) GoString() string {
	return s.String()
}

func (s *QueryK8sClusterLogProjectInfoQuery) SetClusterId(v string) *QueryK8sClusterLogProjectInfoQuery {
	s.ClusterId = &v
	return s
}

type QueryK8sClusterLogProjectInfoRequest struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *QueryK8sClusterLogProjectInfoQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s QueryK8sClusterLogProjectInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryK8sClusterLogProjectInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryK8sClusterLogProjectInfoRequest) SetHeaders(v map[string]*string) *QueryK8sClusterLogProjectInfoRequest {
	s.Headers = v
	return s
}

func (s *QueryK8sClusterLogProjectInfoRequest) SetQuery(v *QueryK8sClusterLogProjectInfoQuery) *QueryK8sClusterLogProjectInfoRequest {
	s.Query = v
	return s
}

type QueryK8sClusterLogProjectInfoResponseBody struct {
	RequestId   *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code        *int                                                  `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message     *string                                               `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ProjectInfo *QueryK8sClusterLogProjectInfoResponseBodyProjectInfo `json:"ProjectInfo,omitempty" xml:"ProjectInfo,omitempty" require:"true" type:"Struct"`
}

func (s QueryK8sClusterLogProjectInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryK8sClusterLogProjectInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryK8sClusterLogProjectInfoResponseBody) SetRequestId(v string) *QueryK8sClusterLogProjectInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryK8sClusterLogProjectInfoResponseBody) SetCode(v int) *QueryK8sClusterLogProjectInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryK8sClusterLogProjectInfoResponseBody) SetMessage(v string) *QueryK8sClusterLogProjectInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryK8sClusterLogProjectInfoResponseBody) SetProjectInfo(v *QueryK8sClusterLogProjectInfoResponseBodyProjectInfo) *QueryK8sClusterLogProjectInfoResponseBody {
	s.ProjectInfo = v
	return s
}

type QueryK8sClusterLogProjectInfoResponseBodyProjectInfo struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty" require:"true"`
}

func (s QueryK8sClusterLogProjectInfoResponseBodyProjectInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryK8sClusterLogProjectInfoResponseBodyProjectInfo) GoString() string {
	return s.String()
}

func (s *QueryK8sClusterLogProjectInfoResponseBodyProjectInfo) SetProjectName(v string) *QueryK8sClusterLogProjectInfoResponseBodyProjectInfo {
	s.ProjectName = &v
	return s
}

type QueryK8sClusterLogProjectInfoResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryK8sClusterLogProjectInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryK8sClusterLogProjectInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryK8sClusterLogProjectInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryK8sClusterLogProjectInfoResponse) SetHeaders(v map[string]*string) *QueryK8sClusterLogProjectInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryK8sClusterLogProjectInfoResponse) SetBody(v *QueryK8sClusterLogProjectInfoResponseBody) *QueryK8sClusterLogProjectInfoResponse {
	s.Body = v
	return s
}

type QuerySlsLogStoreListQuery struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	PageSize    *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QuerySlsLogStoreListQuery) String() string {
	return tea.Prettify(s)
}

func (s QuerySlsLogStoreListQuery) GoString() string {
	return s.String()
}

func (s *QuerySlsLogStoreListQuery) SetAppId(v string) *QuerySlsLogStoreListQuery {
	s.AppId = &v
	return s
}

func (s *QuerySlsLogStoreListQuery) SetType(v string) *QuerySlsLogStoreListQuery {
	s.Type = &v
	return s
}

func (s *QuerySlsLogStoreListQuery) SetPageSize(v int) *QuerySlsLogStoreListQuery {
	s.PageSize = &v
	return s
}

func (s *QuerySlsLogStoreListQuery) SetCurrentPage(v int) *QuerySlsLogStoreListQuery {
	s.CurrentPage = &v
	return s
}

type QuerySlsLogStoreListRequest struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *QuerySlsLogStoreListQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s QuerySlsLogStoreListRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySlsLogStoreListRequest) GoString() string {
	return s.String()
}

func (s *QuerySlsLogStoreListRequest) SetHeaders(v map[string]*string) *QuerySlsLogStoreListRequest {
	s.Headers = v
	return s
}

func (s *QuerySlsLogStoreListRequest) SetQuery(v *QuerySlsLogStoreListQuery) *QuerySlsLogStoreListRequest {
	s.Query = v
	return s
}

type QuerySlsLogStoreListResponseBody struct {
	TotalSize *int                                      `json:"TotalSize,omitempty" xml:"TotalSize,omitempty" require:"true"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                                      `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Result    []*QuerySlsLogStoreListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" require:"true" type:"Repeated"`
}

func (s QuerySlsLogStoreListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySlsLogStoreListResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySlsLogStoreListResponseBody) SetTotalSize(v int) *QuerySlsLogStoreListResponseBody {
	s.TotalSize = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBody) SetRequestId(v string) *QuerySlsLogStoreListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBody) SetCode(v int) *QuerySlsLogStoreListResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBody) SetMessage(v string) *QuerySlsLogStoreListResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBody) SetResult(v []*QuerySlsLogStoreListResponseBodyResult) *QuerySlsLogStoreListResponseBody {
	s.Result = v
	return s
}

type QuerySlsLogStoreListResponseBodyResult struct {
	Logstore     *string `json:"Logstore,omitempty" xml:"Logstore,omitempty" require:"true"`
	Project      *string `json:"Project,omitempty" xml:"Project,omitempty" require:"true"`
	Link         *string `json:"Link,omitempty" xml:"Link,omitempty" require:"true"`
	ConsumerSide *string `json:"ConsumerSide,omitempty" xml:"ConsumerSide,omitempty" require:"true"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	Source       *string `json:"Source,omitempty" xml:"Source,omitempty" require:"true"`
}

func (s QuerySlsLogStoreListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QuerySlsLogStoreListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QuerySlsLogStoreListResponseBodyResult) SetLogstore(v string) *QuerySlsLogStoreListResponseBodyResult {
	s.Logstore = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBodyResult) SetProject(v string) *QuerySlsLogStoreListResponseBodyResult {
	s.Project = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBodyResult) SetLink(v string) *QuerySlsLogStoreListResponseBodyResult {
	s.Link = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBodyResult) SetConsumerSide(v string) *QuerySlsLogStoreListResponseBodyResult {
	s.ConsumerSide = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBodyResult) SetCreateTime(v string) *QuerySlsLogStoreListResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *QuerySlsLogStoreListResponseBodyResult) SetSource(v string) *QuerySlsLogStoreListResponseBodyResult {
	s.Source = &v
	return s
}

type QuerySlsLogStoreListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySlsLogStoreListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySlsLogStoreListResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySlsLogStoreListResponse) GoString() string {
	return s.String()
}

func (s *QuerySlsLogStoreListResponse) SetHeaders(v map[string]*string) *QuerySlsLogStoreListResponse {
	s.Headers = v
	return s
}

func (s *QuerySlsLogStoreListResponse) SetBody(v *QuerySlsLogStoreListResponseBody) *QuerySlsLogStoreListResponse {
	s.Body = v
	return s
}

type GetSubAccountInfoQuery struct {
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty" require:"true"`
}

func (s GetSubAccountInfoQuery) String() string {
	return tea.Prettify(s)
}

func (s GetSubAccountInfoQuery) GoString() string {
	return s.String()
}

func (s *GetSubAccountInfoQuery) SetTargetUserId(v string) *GetSubAccountInfoQuery {
	s.TargetUserId = &v
	return s
}

type GetSubAccountInfoRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *GetSubAccountInfoQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s GetSubAccountInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSubAccountInfoRequest) SetHeaders(v map[string]*string) *GetSubAccountInfoRequest {
	s.Headers = v
	return s
}

func (s *GetSubAccountInfoRequest) SetQuery(v *GetSubAccountInfoQuery) *GetSubAccountInfoRequest {
	s.Query = v
	return s
}

type GetSubAccountInfoResponseBody struct {
	Code          *int                                        `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string                                     `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Authorization *GetSubAccountInfoResponseBodyAuthorization `json:"Authorization,omitempty" xml:"Authorization,omitempty" require:"true" type:"Struct"`
}

func (s GetSubAccountInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubAccountInfoResponseBody) SetCode(v int) *GetSubAccountInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubAccountInfoResponseBody) SetMessage(v string) *GetSubAccountInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubAccountInfoResponseBody) SetRequestId(v string) *GetSubAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubAccountInfoResponseBody) SetAuthorization(v *GetSubAccountInfoResponseBodyAuthorization) *GetSubAccountInfoResponseBody {
	s.Authorization = v
	return s
}

type GetSubAccountInfoResponseBodyAuthorization struct {
	AdminUserId    *string `json:"AdminUserId,omitempty" xml:"AdminUserId,omitempty" require:"true"`
	AdminEdasId    *string `json:"AdminEdasId,omitempty" xml:"AdminEdasId,omitempty" require:"true"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	EdasId         *string `json:"EdasId,omitempty" xml:"EdasId,omitempty" require:"true"`
	AppIdData      *string `json:"AppIdData,omitempty" xml:"AppIdData,omitempty" require:"true"`
	RoleIdData     *string `json:"RoleIdData,omitempty" xml:"RoleIdData,omitempty" require:"true"`
	CreateTime     *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime     *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	ResGroupId     *int64  `json:"ResGroupId,omitempty" xml:"ResGroupId,omitempty" require:"true"`
	ResGroupIdData *string `json:"ResGroupIdData,omitempty" xml:"ResGroupIdData,omitempty" require:"true"`
	IsRamSlave     *bool   `json:"IsRamSlave,omitempty" xml:"IsRamSlave,omitempty" require:"true"`
	IsRamDel       *bool   `json:"IsRamDel,omitempty" xml:"IsRamDel,omitempty" require:"true"`
	SubUserKp      *string `json:"SubUserKp,omitempty" xml:"SubUserKp,omitempty" require:"true"`
	RamOperation   *bool   `json:"RamOperation,omitempty" xml:"RamOperation,omitempty" require:"true"`
	DelegateAdmin  *bool   `json:"DelegateAdmin,omitempty" xml:"DelegateAdmin,omitempty" require:"true"`
	Sts            *bool   `json:"Sts,omitempty" xml:"Sts,omitempty" require:"true"`
}

func (s GetSubAccountInfoResponseBodyAuthorization) String() string {
	return tea.Prettify(s)
}

func (s GetSubAccountInfoResponseBodyAuthorization) GoString() string {
	return s.String()
}

func (s *GetSubAccountInfoResponseBodyAuthorization) SetAdminUserId(v string) *GetSubAccountInfoResponseBodyAuthorization {
	s.AdminUserId = &v
	return s
}

func (s *GetSubAccountInfoResponseBodyAuthorization) SetAdminEdasId(v string) *GetSubAccountInfoResponseBodyAuthorization {
	s.AdminEdasId = &v
	return s
}

func (s *GetSubAccountInfoResponseBodyAuthorization) SetUserId(v string) *GetSubAccountInfoResponseBodyAuthorization {
	s.UserId = &v
	return s
}

func (s *GetSubAccountInfoResponseBodyAuthorization) SetEdasId(v string) *GetSubAccountInfoResponseBodyAuthorization {
	s.EdasId = &v
	return s
}

func (s *GetSubAccountInfoResponseBodyAuthorization) SetAppIdData(v string) *GetSubAccountInfoResponseBodyAuthorization {
	s.AppIdData = &v
	return s
}

func (s *GetSubAccountInfoResponseBodyAuthorization) SetRoleIdData(v string) *GetSubAccountInfoResponseBodyAuthorization {
	s.RoleIdData = &v
	return s
}

func (s *GetSubAccountInfoResponseBodyAuthorization) SetCreateTime(v int64) *GetSubAccountInfoResponseBodyAuthorization {
	s.CreateTime = &v
	return s
}

func (s *GetSubAccountInfoResponseBodyAuthorization) SetUpdateTime(v int64) *GetSubAccountInfoResponseBodyAuthorization {
	s.UpdateTime = &v
	return s
}

func (s *GetSubAccountInfoResponseBodyAuthorization) SetResGroupId(v int64) *GetSubAccountInfoResponseBodyAuthorization {
	s.ResGroupId = &v
	return s
}

func (s *GetSubAccountInfoResponseBodyAuthorization) SetResGroupIdData(v string) *GetSubAccountInfoResponseBodyAuthorization {
	s.ResGroupIdData = &v
	return s
}

func (s *GetSubAccountInfoResponseBodyAuthorization) SetIsRamSlave(v bool) *GetSubAccountInfoResponseBodyAuthorization {
	s.IsRamSlave = &v
	return s
}

func (s *GetSubAccountInfoResponseBodyAuthorization) SetIsRamDel(v bool) *GetSubAccountInfoResponseBodyAuthorization {
	s.IsRamDel = &v
	return s
}

func (s *GetSubAccountInfoResponseBodyAuthorization) SetSubUserKp(v string) *GetSubAccountInfoResponseBodyAuthorization {
	s.SubUserKp = &v
	return s
}

func (s *GetSubAccountInfoResponseBodyAuthorization) SetRamOperation(v bool) *GetSubAccountInfoResponseBodyAuthorization {
	s.RamOperation = &v
	return s
}

func (s *GetSubAccountInfoResponseBodyAuthorization) SetDelegateAdmin(v bool) *GetSubAccountInfoResponseBodyAuthorization {
	s.DelegateAdmin = &v
	return s
}

func (s *GetSubAccountInfoResponseBodyAuthorization) SetSts(v bool) *GetSubAccountInfoResponseBodyAuthorization {
	s.Sts = &v
	return s
}

type GetSubAccountInfoResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSubAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSubAccountInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSubAccountInfoResponse) SetHeaders(v map[string]*string) *GetSubAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSubAccountInfoResponse) SetBody(v *GetSubAccountInfoResponseBody) *GetSubAccountInfoResponse {
	s.Body = v
	return s
}

type DelegateAdminRoleQuery struct {
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty" require:"true"`
}

func (s DelegateAdminRoleQuery) String() string {
	return tea.Prettify(s)
}

func (s DelegateAdminRoleQuery) GoString() string {
	return s.String()
}

func (s *DelegateAdminRoleQuery) SetTargetUserId(v string) *DelegateAdminRoleQuery {
	s.TargetUserId = &v
	return s
}

type DelegateAdminRoleRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DelegateAdminRoleQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DelegateAdminRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s DelegateAdminRoleRequest) GoString() string {
	return s.String()
}

func (s *DelegateAdminRoleRequest) SetHeaders(v map[string]*string) *DelegateAdminRoleRequest {
	s.Headers = v
	return s
}

func (s *DelegateAdminRoleRequest) SetQuery(v *DelegateAdminRoleQuery) *DelegateAdminRoleRequest {
	s.Query = v
	return s
}

type DelegateAdminRoleResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DelegateAdminRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DelegateAdminRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DelegateAdminRoleResponseBody) SetCode(v int) *DelegateAdminRoleResponseBody {
	s.Code = &v
	return s
}

func (s *DelegateAdminRoleResponseBody) SetMessage(v string) *DelegateAdminRoleResponseBody {
	s.Message = &v
	return s
}

func (s *DelegateAdminRoleResponseBody) SetRequestId(v string) *DelegateAdminRoleResponseBody {
	s.RequestId = &v
	return s
}

type DelegateAdminRoleResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DelegateAdminRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DelegateAdminRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s DelegateAdminRoleResponse) GoString() string {
	return s.String()
}

func (s *DelegateAdminRoleResponse) SetHeaders(v map[string]*string) *DelegateAdminRoleResponse {
	s.Headers = v
	return s
}

func (s *DelegateAdminRoleResponse) SetBody(v *DelegateAdminRoleResponseBody) *DelegateAdminRoleResponse {
	s.Body = v
	return s
}

type RestartApplicationQuery struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	EccInfo *string `json:"EccInfo,omitempty" xml:"EccInfo,omitempty"`
}

func (s RestartApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s RestartApplicationQuery) GoString() string {
	return s.String()
}

func (s *RestartApplicationQuery) SetAppId(v string) *RestartApplicationQuery {
	s.AppId = &v
	return s
}

func (s *RestartApplicationQuery) SetEccInfo(v string) *RestartApplicationQuery {
	s.EccInfo = &v
	return s
}

type RestartApplicationRequest struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *RestartApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s RestartApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartApplicationRequest) GoString() string {
	return s.String()
}

func (s *RestartApplicationRequest) SetHeaders(v map[string]*string) *RestartApplicationRequest {
	s.Headers = v
	return s
}

func (s *RestartApplicationRequest) SetQuery(v *RestartApplicationQuery) *RestartApplicationRequest {
	s.Query = v
	return s
}

type RestartApplicationResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s RestartApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RestartApplicationResponseBody) SetRequestId(v string) *RestartApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartApplicationResponseBody) SetCode(v int) *RestartApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *RestartApplicationResponseBody) SetMessage(v string) *RestartApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *RestartApplicationResponseBody) SetChangeOrderId(v string) *RestartApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

type RestartApplicationResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartApplicationResponse) GoString() string {
	return s.String()
}

func (s *RestartApplicationResponse) SetHeaders(v map[string]*string) *RestartApplicationResponse {
	s.Headers = v
	return s
}

func (s *RestartApplicationResponse) SetBody(v *RestartApplicationResponseBody) *RestartApplicationResponse {
	s.Body = v
	return s
}

type DeleteLogPathQuery struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Path  *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DeleteLogPathQuery) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogPathQuery) GoString() string {
	return s.String()
}

func (s *DeleteLogPathQuery) SetAppId(v string) *DeleteLogPathQuery {
	s.AppId = &v
	return s
}

func (s *DeleteLogPathQuery) SetPath(v string) *DeleteLogPathQuery {
	s.Path = &v
	return s
}

type DeleteLogPathRequest struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeleteLogPathQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DeleteLogPathRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogPathRequest) GoString() string {
	return s.String()
}

func (s *DeleteLogPathRequest) SetHeaders(v map[string]*string) *DeleteLogPathRequest {
	s.Headers = v
	return s
}

func (s *DeleteLogPathRequest) SetQuery(v *DeleteLogPathQuery) *DeleteLogPathRequest {
	s.Query = v
	return s
}

type DeleteLogPathResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s DeleteLogPathResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogPathResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLogPathResponseBody) SetRequestId(v string) *DeleteLogPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLogPathResponseBody) SetCode(v string) *DeleteLogPathResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLogPathResponseBody) SetMessage(v string) *DeleteLogPathResponseBody {
	s.Message = &v
	return s
}

type DeleteLogPathResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLogPathResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLogPathResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogPathResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogPathResponse) SetHeaders(v map[string]*string) *DeleteLogPathResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogPathResponse) SetBody(v *DeleteLogPathResponseBody) *DeleteLogPathResponse {
	s.Body = v
	return s
}

type AddLogPathBody struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Path  *string `json:"Path,omitempty" xml:"Path,omitempty" require:"true"`
}

func (s AddLogPathBody) String() string {
	return tea.Prettify(s)
}

func (s AddLogPathBody) GoString() string {
	return s.String()
}

func (s *AddLogPathBody) SetAppId(v string) *AddLogPathBody {
	s.AppId = &v
	return s
}

func (s *AddLogPathBody) SetPath(v string) *AddLogPathBody {
	s.Path = &v
	return s
}

type AddLogPathRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *AddLogPathBody    `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddLogPathRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLogPathRequest) GoString() string {
	return s.String()
}

func (s *AddLogPathRequest) SetHeaders(v map[string]*string) *AddLogPathRequest {
	s.Headers = v
	return s
}

func (s *AddLogPathRequest) SetBody(v *AddLogPathBody) *AddLogPathRequest {
	s.Body = v
	return s
}

type AddLogPathResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s AddLogPathResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddLogPathResponseBody) GoString() string {
	return s.String()
}

func (s *AddLogPathResponseBody) SetRequestId(v string) *AddLogPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLogPathResponseBody) SetCode(v int) *AddLogPathResponseBody {
	s.Code = &v
	return s
}

func (s *AddLogPathResponseBody) SetMessage(v string) *AddLogPathResponseBody {
	s.Message = &v
	return s
}

type AddLogPathResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddLogPathResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddLogPathResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLogPathResponse) GoString() string {
	return s.String()
}

func (s *AddLogPathResponse) SetHeaders(v map[string]*string) *AddLogPathResponse {
	s.Headers = v
	return s
}

func (s *AddLogPathResponse) SetBody(v *AddLogPathResponseBody) *AddLogPathResponse {
	s.Body = v
	return s
}

type UpdateK8sResourceBody struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ResourceContent *string `json:"ResourceContent,omitempty" xml:"ResourceContent,omitempty"`
}

func (s UpdateK8sResourceBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sResourceBody) GoString() string {
	return s.String()
}

func (s *UpdateK8sResourceBody) SetClusterId(v string) *UpdateK8sResourceBody {
	s.ClusterId = &v
	return s
}

func (s *UpdateK8sResourceBody) SetNamespace(v string) *UpdateK8sResourceBody {
	s.Namespace = &v
	return s
}

func (s *UpdateK8sResourceBody) SetResourceContent(v string) *UpdateK8sResourceBody {
	s.ResourceContent = &v
	return s
}

type UpdateK8sResourceRequest struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	Body    *UpdateK8sResourceBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateK8sResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateK8sResourceRequest) SetHeaders(v map[string]*string) *UpdateK8sResourceRequest {
	s.Headers = v
	return s
}

func (s *UpdateK8sResourceRequest) SetBody(v *UpdateK8sResourceBody) *UpdateK8sResourceRequest {
	s.Body = v
	return s
}

type UpdateK8sResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s UpdateK8sResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateK8sResourceResponseBody) SetRequestId(v string) *UpdateK8sResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateK8sResourceResponseBody) SetCode(v int) *UpdateK8sResourceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateK8sResourceResponseBody) SetMessage(v string) *UpdateK8sResourceResponseBody {
	s.Message = &v
	return s
}

type UpdateK8sResourceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateK8sResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateK8sResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateK8sResourceResponse) SetHeaders(v map[string]*string) *UpdateK8sResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateK8sResourceResponse) SetBody(v *UpdateK8sResourceResponseBody) *UpdateK8sResourceResponse {
	s.Body = v
	return s
}

type AbortChangeOrderQuery struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s AbortChangeOrderQuery) String() string {
	return tea.Prettify(s)
}

func (s AbortChangeOrderQuery) GoString() string {
	return s.String()
}

func (s *AbortChangeOrderQuery) SetChangeOrderId(v string) *AbortChangeOrderQuery {
	s.ChangeOrderId = &v
	return s
}

type AbortChangeOrderRequest struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *AbortChangeOrderQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s AbortChangeOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s AbortChangeOrderRequest) GoString() string {
	return s.String()
}

func (s *AbortChangeOrderRequest) SetHeaders(v map[string]*string) *AbortChangeOrderRequest {
	s.Headers = v
	return s
}

func (s *AbortChangeOrderRequest) SetQuery(v *AbortChangeOrderQuery) *AbortChangeOrderRequest {
	s.Query = v
	return s
}

type AbortChangeOrderResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                              `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ErrorCode *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty" require:"true"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	TraceId   *string                           `json:"TraceId,omitempty" xml:"TraceId,omitempty" require:"true"`
	Data      *AbortChangeOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s AbortChangeOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbortChangeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *AbortChangeOrderResponseBody) SetRequestId(v string) *AbortChangeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetCode(v int) *AbortChangeOrderResponseBody {
	s.Code = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetErrorCode(v string) *AbortChangeOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetMessage(v string) *AbortChangeOrderResponseBody {
	s.Message = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetTraceId(v string) *AbortChangeOrderResponseBody {
	s.TraceId = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetData(v *AbortChangeOrderResponseBodyData) *AbortChangeOrderResponseBody {
	s.Data = v
	return s
}

type AbortChangeOrderResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s AbortChangeOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AbortChangeOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *AbortChangeOrderResponseBodyData) SetChangeOrderId(v string) *AbortChangeOrderResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type AbortChangeOrderResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AbortChangeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AbortChangeOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s AbortChangeOrderResponse) GoString() string {
	return s.String()
}

func (s *AbortChangeOrderResponse) SetHeaders(v map[string]*string) *AbortChangeOrderResponse {
	s.Headers = v
	return s
}

func (s *AbortChangeOrderResponse) SetBody(v *AbortChangeOrderResponseBody) *AbortChangeOrderResponse {
	s.Body = v
	return s
}

type RollbackChangeOrderQuery struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s RollbackChangeOrderQuery) String() string {
	return tea.Prettify(s)
}

func (s RollbackChangeOrderQuery) GoString() string {
	return s.String()
}

func (s *RollbackChangeOrderQuery) SetChangeOrderId(v string) *RollbackChangeOrderQuery {
	s.ChangeOrderId = &v
	return s
}

type RollbackChangeOrderRequest struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *RollbackChangeOrderQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s RollbackChangeOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackChangeOrderRequest) GoString() string {
	return s.String()
}

func (s *RollbackChangeOrderRequest) SetHeaders(v map[string]*string) *RollbackChangeOrderRequest {
	s.Headers = v
	return s
}

func (s *RollbackChangeOrderRequest) SetQuery(v *RollbackChangeOrderQuery) *RollbackChangeOrderRequest {
	s.Query = v
	return s
}

type RollbackChangeOrderResponseBody struct {
	Code      *int                                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ErrorCode *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty" require:"true"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TraceId   *string                              `json:"TraceId,omitempty" xml:"TraceId,omitempty" require:"true"`
	Data      *RollbackChangeOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s RollbackChangeOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackChangeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackChangeOrderResponseBody) SetCode(v int) *RollbackChangeOrderResponseBody {
	s.Code = &v
	return s
}

func (s *RollbackChangeOrderResponseBody) SetErrorCode(v string) *RollbackChangeOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RollbackChangeOrderResponseBody) SetMessage(v string) *RollbackChangeOrderResponseBody {
	s.Message = &v
	return s
}

func (s *RollbackChangeOrderResponseBody) SetRequestId(v string) *RollbackChangeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackChangeOrderResponseBody) SetTraceId(v string) *RollbackChangeOrderResponseBody {
	s.TraceId = &v
	return s
}

func (s *RollbackChangeOrderResponseBody) SetData(v *RollbackChangeOrderResponseBodyData) *RollbackChangeOrderResponseBody {
	s.Data = v
	return s
}

type RollbackChangeOrderResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s RollbackChangeOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RollbackChangeOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *RollbackChangeOrderResponseBodyData) SetChangeOrderId(v string) *RollbackChangeOrderResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type RollbackChangeOrderResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RollbackChangeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackChangeOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackChangeOrderResponse) GoString() string {
	return s.String()
}

func (s *RollbackChangeOrderResponse) SetHeaders(v map[string]*string) *RollbackChangeOrderResponse {
	s.Headers = v
	return s
}

func (s *RollbackChangeOrderResponse) SetBody(v *RollbackChangeOrderResponseBody) *RollbackChangeOrderResponse {
	s.Body = v
	return s
}

type AbortAndRollbackChangeOrderQuery struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s AbortAndRollbackChangeOrderQuery) String() string {
	return tea.Prettify(s)
}

func (s AbortAndRollbackChangeOrderQuery) GoString() string {
	return s.String()
}

func (s *AbortAndRollbackChangeOrderQuery) SetChangeOrderId(v string) *AbortAndRollbackChangeOrderQuery {
	s.ChangeOrderId = &v
	return s
}

type AbortAndRollbackChangeOrderRequest struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *AbortAndRollbackChangeOrderQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s AbortAndRollbackChangeOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s AbortAndRollbackChangeOrderRequest) GoString() string {
	return s.String()
}

func (s *AbortAndRollbackChangeOrderRequest) SetHeaders(v map[string]*string) *AbortAndRollbackChangeOrderRequest {
	s.Headers = v
	return s
}

func (s *AbortAndRollbackChangeOrderRequest) SetQuery(v *AbortAndRollbackChangeOrderQuery) *AbortAndRollbackChangeOrderRequest {
	s.Query = v
	return s
}

type AbortAndRollbackChangeOrderResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                                         `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	ErrorCode *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty" require:"true"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	TraceId   *string                                      `json:"TraceId,omitempty" xml:"TraceId,omitempty" require:"true"`
	Data      *AbortAndRollbackChangeOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s AbortAndRollbackChangeOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbortAndRollbackChangeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetRequestId(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetCode(v int) *AbortAndRollbackChangeOrderResponseBody {
	s.Code = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetErrorCode(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetMessage(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.Message = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetTraceId(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.TraceId = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetData(v *AbortAndRollbackChangeOrderResponseBodyData) *AbortAndRollbackChangeOrderResponseBody {
	s.Data = v
	return s
}

type AbortAndRollbackChangeOrderResponseBodyData struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s AbortAndRollbackChangeOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AbortAndRollbackChangeOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *AbortAndRollbackChangeOrderResponseBodyData) SetChangeOrderId(v string) *AbortAndRollbackChangeOrderResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

type AbortAndRollbackChangeOrderResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AbortAndRollbackChangeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AbortAndRollbackChangeOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s AbortAndRollbackChangeOrderResponse) GoString() string {
	return s.String()
}

func (s *AbortAndRollbackChangeOrderResponse) SetHeaders(v map[string]*string) *AbortAndRollbackChangeOrderResponse {
	s.Headers = v
	return s
}

func (s *AbortAndRollbackChangeOrderResponse) SetBody(v *AbortAndRollbackChangeOrderResponseBody) *AbortAndRollbackChangeOrderResponse {
	s.Body = v
	return s
}

type GetScalingRulesQuery struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	Mode    *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s GetScalingRulesQuery) String() string {
	return tea.Prettify(s)
}

func (s GetScalingRulesQuery) GoString() string {
	return s.String()
}

func (s *GetScalingRulesQuery) SetAppId(v string) *GetScalingRulesQuery {
	s.AppId = &v
	return s
}

func (s *GetScalingRulesQuery) SetGroupId(v string) *GetScalingRulesQuery {
	s.GroupId = &v
	return s
}

func (s *GetScalingRulesQuery) SetMode(v string) *GetScalingRulesQuery {
	s.Mode = &v
	return s
}

type GetScalingRulesRequest struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *GetScalingRulesQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s GetScalingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScalingRulesRequest) GoString() string {
	return s.String()
}

func (s *GetScalingRulesRequest) SetHeaders(v map[string]*string) *GetScalingRulesRequest {
	s.Headers = v
	return s
}

func (s *GetScalingRulesRequest) SetQuery(v *GetScalingRulesQuery) *GetScalingRulesRequest {
	s.Query = v
	return s
}

type GetScalingRulesResponseBody struct {
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code       *int                             `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message    *string                          `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	UpdateTime *int64                           `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	Data       *GetScalingRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetScalingRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScalingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *GetScalingRulesResponseBody) SetRequestId(v string) *GetScalingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScalingRulesResponseBody) SetCode(v int) *GetScalingRulesResponseBody {
	s.Code = &v
	return s
}

func (s *GetScalingRulesResponseBody) SetMessage(v string) *GetScalingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *GetScalingRulesResponseBody) SetUpdateTime(v int64) *GetScalingRulesResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetScalingRulesResponseBody) SetData(v *GetScalingRulesResponseBodyData) *GetScalingRulesResponseBody {
	s.Data = v
	return s
}

type GetScalingRulesResponseBodyData struct {
	ClusterType    *int                                     `json:"ClusterType,omitempty" xml:"ClusterType,omitempty" require:"true"`
	OversoldFactor *int                                     `json:"OversoldFactor,omitempty" xml:"OversoldFactor,omitempty" require:"true"`
	VpcId          *string                                  `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	UpdateTime     *int64                                   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	RuleList       *GetScalingRulesResponseBodyDataRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" require:"true" type:"Struct"`
}

func (s GetScalingRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetScalingRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScalingRulesResponseBodyData) SetClusterType(v int) *GetScalingRulesResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *GetScalingRulesResponseBodyData) SetOversoldFactor(v int) *GetScalingRulesResponseBodyData {
	s.OversoldFactor = &v
	return s
}

func (s *GetScalingRulesResponseBodyData) SetVpcId(v string) *GetScalingRulesResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *GetScalingRulesResponseBodyData) SetUpdateTime(v int64) *GetScalingRulesResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetScalingRulesResponseBodyData) SetRuleList(v *GetScalingRulesResponseBodyDataRuleList) *GetScalingRulesResponseBodyData {
	s.RuleList = v
	return s
}

type GetScalingRulesResponseBodyDataRuleList struct {
	Rule []*GetScalingRulesResponseBodyDataRuleListRule `json:"Rule,omitempty" xml:"Rule,omitempty" require:"true" type:"Repeated"`
}

func (s GetScalingRulesResponseBodyDataRuleList) String() string {
	return tea.Prettify(s)
}

func (s GetScalingRulesResponseBodyDataRuleList) GoString() string {
	return s.String()
}

func (s *GetScalingRulesResponseBodyDataRuleList) SetRule(v []*GetScalingRulesResponseBodyDataRuleListRule) *GetScalingRulesResponseBodyDataRuleList {
	s.Rule = v
	return s
}

type GetScalingRulesResponseBodyDataRuleListRule struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	Enable          *bool   `json:"Enable,omitempty" xml:"Enable,omitempty" require:"true"`
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty" require:"true"`
	Cond            *string `json:"Cond,omitempty" xml:"Cond,omitempty" require:"true"`
	Cpu             *int    `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Rt              *int    `json:"Rt,omitempty" xml:"Rt,omitempty" require:"true"`
	ResourceFrom    *string `json:"ResourceFrom,omitempty" xml:"ResourceFrom,omitempty" require:"true"`
	Step            *int    `json:"Step,omitempty" xml:"Step,omitempty" require:"true"`
	InstNum         *int    `json:"InstNum,omitempty" xml:"InstNum,omitempty" require:"true"`
	LoadNum         *int    `json:"LoadNum,omitempty" xml:"LoadNum,omitempty" require:"true"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty" require:"true"`
	TemplateVersion *int    `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty" require:"true"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	MetricType      *string `json:"MetricType,omitempty" xml:"MetricType,omitempty" require:"true"`
	VSwitchIds      *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" require:"true"`
	MultiAzPolicy   *string `json:"MultiAzPolicy,omitempty" xml:"MultiAzPolicy,omitempty" require:"true"`
	SpecId          *string `json:"SpecId,omitempty" xml:"SpecId,omitempty" require:"true"`
	Duration        *int    `json:"Duration,omitempty" xml:"Duration,omitempty" require:"true"`
	CreateTime      *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime      *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
}

func (s GetScalingRulesResponseBodyDataRuleListRule) String() string {
	return tea.Prettify(s)
}

func (s GetScalingRulesResponseBodyDataRuleListRule) GoString() string {
	return s.String()
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetAppId(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.AppId = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetGroupId(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.GroupId = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetEnable(v bool) *GetScalingRulesResponseBodyDataRuleListRule {
	s.Enable = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetMode(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.Mode = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetCond(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.Cond = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetCpu(v int) *GetScalingRulesResponseBodyDataRuleListRule {
	s.Cpu = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetRt(v int) *GetScalingRulesResponseBodyDataRuleListRule {
	s.Rt = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetResourceFrom(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.ResourceFrom = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetStep(v int) *GetScalingRulesResponseBodyDataRuleListRule {
	s.Step = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetInstNum(v int) *GetScalingRulesResponseBodyDataRuleListRule {
	s.InstNum = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetLoadNum(v int) *GetScalingRulesResponseBodyDataRuleListRule {
	s.LoadNum = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetTemplateId(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.TemplateId = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetTemplateVersion(v int) *GetScalingRulesResponseBodyDataRuleListRule {
	s.TemplateVersion = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetVpcId(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.VpcId = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetMetricType(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.MetricType = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetVSwitchIds(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.VSwitchIds = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetMultiAzPolicy(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.MultiAzPolicy = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetSpecId(v string) *GetScalingRulesResponseBodyDataRuleListRule {
	s.SpecId = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetDuration(v int) *GetScalingRulesResponseBodyDataRuleListRule {
	s.Duration = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetCreateTime(v int64) *GetScalingRulesResponseBodyDataRuleListRule {
	s.CreateTime = &v
	return s
}

func (s *GetScalingRulesResponseBodyDataRuleListRule) SetUpdateTime(v int64) *GetScalingRulesResponseBodyDataRuleListRule {
	s.UpdateTime = &v
	return s
}

type GetScalingRulesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetScalingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetScalingRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScalingRulesResponse) GoString() string {
	return s.String()
}

func (s *GetScalingRulesResponse) SetHeaders(v map[string]*string) *GetScalingRulesResponse {
	s.Headers = v
	return s
}

func (s *GetScalingRulesResponse) SetBody(v *GetScalingRulesResponseBody) *GetScalingRulesResponse {
	s.Body = v
	return s
}

type ModifyScalingRuleQuery struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	InCondition          *string `json:"InCondition,omitempty" xml:"InCondition,omitempty"`
	InEnable             *bool   `json:"InEnable,omitempty" xml:"InEnable,omitempty"`
	InCpu                *int    `json:"InCpu,omitempty" xml:"InCpu,omitempty"`
	InDuration           *int    `json:"InDuration,omitempty" xml:"InDuration,omitempty"`
	InInstanceNum        *int    `json:"InInstanceNum,omitempty" xml:"InInstanceNum,omitempty"`
	InLoad               *int    `json:"InLoad,omitempty" xml:"InLoad,omitempty"`
	InRT                 *int    `json:"InRT,omitempty" xml:"InRT,omitempty"`
	InStep               *int    `json:"InStep,omitempty" xml:"InStep,omitempty"`
	OutCondition         *string `json:"OutCondition,omitempty" xml:"OutCondition,omitempty"`
	OutCPU               *int    `json:"OutCPU,omitempty" xml:"OutCPU,omitempty"`
	OutDuration          *int    `json:"OutDuration,omitempty" xml:"OutDuration,omitempty"`
	OutEnable            *bool   `json:"OutEnable,omitempty" xml:"OutEnable,omitempty"`
	OutInstanceNum       *int    `json:"OutInstanceNum,omitempty" xml:"OutInstanceNum,omitempty"`
	OutLoad              *int    `json:"OutLoad,omitempty" xml:"OutLoad,omitempty"`
	OutRT                *int    `json:"OutRT,omitempty" xml:"OutRT,omitempty"`
	OutStep              *int    `json:"OutStep,omitempty" xml:"OutStep,omitempty"`
	ResourceFrom         *string `json:"ResourceFrom,omitempty" xml:"ResourceFrom,omitempty"`
	MultiAzPolicy        *string `json:"MultiAzPolicy,omitempty" xml:"MultiAzPolicy,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchIds           *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	ScalingPolicy        *string `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty"`
	TemplateInstanceId   *string `json:"TemplateInstanceId,omitempty" xml:"TemplateInstanceId,omitempty"`
	TemplateInstanceName *string `json:"TemplateInstanceName,omitempty" xml:"TemplateInstanceName,omitempty"`
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	KeyPairName          *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	AcceptEULA           *bool   `json:"AcceptEULA,omitempty" xml:"AcceptEULA,omitempty"`
	TemplateId           *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateVersion      *int    `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s ModifyScalingRuleQuery) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingRuleQuery) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleQuery) SetAppId(v string) *ModifyScalingRuleQuery {
	s.AppId = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetGroupId(v string) *ModifyScalingRuleQuery {
	s.GroupId = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetInCondition(v string) *ModifyScalingRuleQuery {
	s.InCondition = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetInEnable(v bool) *ModifyScalingRuleQuery {
	s.InEnable = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetInCpu(v int) *ModifyScalingRuleQuery {
	s.InCpu = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetInDuration(v int) *ModifyScalingRuleQuery {
	s.InDuration = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetInInstanceNum(v int) *ModifyScalingRuleQuery {
	s.InInstanceNum = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetInLoad(v int) *ModifyScalingRuleQuery {
	s.InLoad = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetInRT(v int) *ModifyScalingRuleQuery {
	s.InRT = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetInStep(v int) *ModifyScalingRuleQuery {
	s.InStep = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetOutCondition(v string) *ModifyScalingRuleQuery {
	s.OutCondition = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetOutCPU(v int) *ModifyScalingRuleQuery {
	s.OutCPU = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetOutDuration(v int) *ModifyScalingRuleQuery {
	s.OutDuration = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetOutEnable(v bool) *ModifyScalingRuleQuery {
	s.OutEnable = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetOutInstanceNum(v int) *ModifyScalingRuleQuery {
	s.OutInstanceNum = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetOutLoad(v int) *ModifyScalingRuleQuery {
	s.OutLoad = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetOutRT(v int) *ModifyScalingRuleQuery {
	s.OutRT = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetOutStep(v int) *ModifyScalingRuleQuery {
	s.OutStep = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetResourceFrom(v string) *ModifyScalingRuleQuery {
	s.ResourceFrom = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetMultiAzPolicy(v string) *ModifyScalingRuleQuery {
	s.MultiAzPolicy = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetVpcId(v string) *ModifyScalingRuleQuery {
	s.VpcId = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetVSwitchIds(v string) *ModifyScalingRuleQuery {
	s.VSwitchIds = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetScalingPolicy(v string) *ModifyScalingRuleQuery {
	s.ScalingPolicy = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetTemplateInstanceId(v string) *ModifyScalingRuleQuery {
	s.TemplateInstanceId = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetTemplateInstanceName(v string) *ModifyScalingRuleQuery {
	s.TemplateInstanceName = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetPassword(v string) *ModifyScalingRuleQuery {
	s.Password = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetKeyPairName(v string) *ModifyScalingRuleQuery {
	s.KeyPairName = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetAcceptEULA(v bool) *ModifyScalingRuleQuery {
	s.AcceptEULA = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetTemplateId(v string) *ModifyScalingRuleQuery {
	s.TemplateId = &v
	return s
}

func (s *ModifyScalingRuleQuery) SetTemplateVersion(v int) *ModifyScalingRuleQuery {
	s.TemplateVersion = &v
	return s
}

type ModifyScalingRuleRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ModifyScalingRuleQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ModifyScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleRequest) SetHeaders(v map[string]*string) *ModifyScalingRuleRequest {
	s.Headers = v
	return s
}

func (s *ModifyScalingRuleRequest) SetQuery(v *ModifyScalingRuleQuery) *ModifyScalingRuleRequest {
	s.Query = v
	return s
}

type ModifyScalingRuleResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleResponseBody) SetCode(v int) *ModifyScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyScalingRuleResponseBody) SetMessage(v string) *ModifyScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyScalingRuleResponseBody) SetRequestId(v string) *ModifyScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyScalingRuleResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleResponse) SetHeaders(v map[string]*string) *ModifyScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyScalingRuleResponse) SetBody(v *ModifyScalingRuleResponseBody) *ModifyScalingRuleResponse {
	s.Body = v
	return s
}

type ListMethodsQuery struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" require:"true"`
}

func (s ListMethodsQuery) String() string {
	return tea.Prettify(s)
}

func (s ListMethodsQuery) GoString() string {
	return s.String()
}

func (s *ListMethodsQuery) SetAppId(v string) *ListMethodsQuery {
	s.AppId = &v
	return s
}

func (s *ListMethodsQuery) SetServiceName(v string) *ListMethodsQuery {
	s.ServiceName = &v
	return s
}

type ListMethodsRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListMethodsQuery  `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ListMethodsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMethodsRequest) GoString() string {
	return s.String()
}

func (s *ListMethodsRequest) SetHeaders(v map[string]*string) *ListMethodsRequest {
	s.Headers = v
	return s
}

func (s *ListMethodsRequest) SetQuery(v *ListMethodsQuery) *ListMethodsRequest {
	s.Query = v
	return s
}

type ListMethodsResponseBody struct {
	Code              *int                                      `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message           *string                                   `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId         *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ServiceMethodList *ListMethodsResponseBodyServiceMethodList `json:"ServiceMethodList,omitempty" xml:"ServiceMethodList,omitempty" require:"true" type:"Struct"`
}

func (s ListMethodsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMethodsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMethodsResponseBody) SetCode(v int) *ListMethodsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMethodsResponseBody) SetMessage(v string) *ListMethodsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMethodsResponseBody) SetRequestId(v string) *ListMethodsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMethodsResponseBody) SetServiceMethodList(v *ListMethodsResponseBodyServiceMethodList) *ListMethodsResponseBody {
	s.ServiceMethodList = v
	return s
}

type ListMethodsResponseBodyServiceMethodList struct {
	ServiceMethod []*ListMethodsResponseBodyServiceMethodListServiceMethod `json:"ServiceMethod,omitempty" xml:"ServiceMethod,omitempty" require:"true" type:"Repeated"`
}

func (s ListMethodsResponseBodyServiceMethodList) String() string {
	return tea.Prettify(s)
}

func (s ListMethodsResponseBodyServiceMethodList) GoString() string {
	return s.String()
}

func (s *ListMethodsResponseBodyServiceMethodList) SetServiceMethod(v []*ListMethodsResponseBodyServiceMethodListServiceMethod) *ListMethodsResponseBodyServiceMethodList {
	s.ServiceMethod = v
	return s
}

type ListMethodsResponseBodyServiceMethodListServiceMethod struct {
	AppName     *string                                                           `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	ServiceName *string                                                           `json:"ServiceName,omitempty" xml:"ServiceName,omitempty" require:"true"`
	MethodName  *string                                                           `json:"MethodName,omitempty" xml:"MethodName,omitempty" require:"true"`
	Output      *string                                                           `json:"Output,omitempty" xml:"Output,omitempty" require:"true"`
	InputParams *ListMethodsResponseBodyServiceMethodListServiceMethodInputParams `json:"InputParams,omitempty" xml:"InputParams,omitempty" require:"true" type:"Struct"`
	ParamTypes  *ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes  `json:"ParamTypes,omitempty" xml:"ParamTypes,omitempty" require:"true" type:"Struct"`
}

func (s ListMethodsResponseBodyServiceMethodListServiceMethod) String() string {
	return tea.Prettify(s)
}

func (s ListMethodsResponseBodyServiceMethodListServiceMethod) GoString() string {
	return s.String()
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) SetAppName(v string) *ListMethodsResponseBodyServiceMethodListServiceMethod {
	s.AppName = &v
	return s
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) SetServiceName(v string) *ListMethodsResponseBodyServiceMethodListServiceMethod {
	s.ServiceName = &v
	return s
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) SetMethodName(v string) *ListMethodsResponseBodyServiceMethodListServiceMethod {
	s.MethodName = &v
	return s
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) SetOutput(v string) *ListMethodsResponseBodyServiceMethodListServiceMethod {
	s.Output = &v
	return s
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) SetInputParams(v *ListMethodsResponseBodyServiceMethodListServiceMethodInputParams) *ListMethodsResponseBodyServiceMethodListServiceMethod {
	s.InputParams = v
	return s
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) SetParamTypes(v *ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes) *ListMethodsResponseBodyServiceMethodListServiceMethod {
	s.ParamTypes = v
	return s
}

type ListMethodsResponseBodyServiceMethodListServiceMethodInputParams struct {
	InputParam []*string `json:"InputParam,omitempty" xml:"InputParam,omitempty" require:"true" type:"Repeated"`
}

func (s ListMethodsResponseBodyServiceMethodListServiceMethodInputParams) String() string {
	return tea.Prettify(s)
}

func (s ListMethodsResponseBodyServiceMethodListServiceMethodInputParams) GoString() string {
	return s.String()
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethodInputParams) SetInputParam(v []*string) *ListMethodsResponseBodyServiceMethodListServiceMethodInputParams {
	s.InputParam = v
	return s
}

type ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes struct {
	ParamType []*string `json:"ParamType,omitempty" xml:"ParamType,omitempty" require:"true" type:"Repeated"`
}

func (s ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes) String() string {
	return tea.Prettify(s)
}

func (s ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes) GoString() string {
	return s.String()
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes) SetParamType(v []*string) *ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes {
	s.ParamType = v
	return s
}

type ListMethodsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMethodsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMethodsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMethodsResponse) GoString() string {
	return s.String()
}

func (s *ListMethodsResponse) SetHeaders(v map[string]*string) *ListMethodsResponse {
	s.Headers = v
	return s
}

func (s *ListMethodsResponse) SetBody(v *ListMethodsResponseBody) *ListMethodsResponse {
	s.Body = v
	return s
}

type GetK8sApplicationQuery struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	From  *string `json:"From,omitempty" xml:"From,omitempty"`
}

func (s GetK8sApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s GetK8sApplicationQuery) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationQuery) SetAppId(v string) *GetK8sApplicationQuery {
	s.AppId = &v
	return s
}

func (s *GetK8sApplicationQuery) SetFrom(v string) *GetK8sApplicationQuery {
	s.From = &v
	return s
}

type GetK8sApplicationRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *GetK8sApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s GetK8sApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationRequest) SetHeaders(v map[string]*string) *GetK8sApplicationRequest {
	s.Headers = v
	return s
}

func (s *GetK8sApplicationRequest) SetQuery(v *GetK8sApplicationQuery) *GetK8sApplicationRequest {
	s.Query = v
	return s
}

type GetK8sApplicationResponseBody struct {
	Code       *int                                     `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Applcation *GetK8sApplicationResponseBodyApplcation `json:"Applcation,omitempty" xml:"Applcation,omitempty" require:"true" type:"Struct"`
}

func (s GetK8sApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetK8sApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBody) SetCode(v int) *GetK8sApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *GetK8sApplicationResponseBody) SetMessage(v string) *GetK8sApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *GetK8sApplicationResponseBody) SetRequestId(v string) *GetK8sApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetK8sApplicationResponseBody) SetApplcation(v *GetK8sApplicationResponseBodyApplcation) *GetK8sApplicationResponseBody {
	s.Applcation = v
	return s
}

type GetK8sApplicationResponseBodyApplcation struct {
	AppId        *string                                              `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	DeployGroups *GetK8sApplicationResponseBodyApplcationDeployGroups `json:"DeployGroups,omitempty" xml:"DeployGroups,omitempty" require:"true" type:"Struct"`
	ImageInfo    *GetK8sApplicationResponseBodyApplcationImageInfo    `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" require:"true" type:"Struct"`
	App          *GetK8sApplicationResponseBodyApplcationApp          `json:"App,omitempty" xml:"App,omitempty" require:"true" type:"Struct"`
	Conf         *GetK8sApplicationResponseBodyApplcationConf         `json:"Conf,omitempty" xml:"Conf,omitempty" require:"true" type:"Struct"`
}

func (s GetK8sApplicationResponseBodyApplcation) String() string {
	return tea.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcation) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcation) SetAppId(v string) *GetK8sApplicationResponseBodyApplcation {
	s.AppId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcation) SetDeployGroups(v *GetK8sApplicationResponseBodyApplcationDeployGroups) *GetK8sApplicationResponseBodyApplcation {
	s.DeployGroups = v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcation) SetImageInfo(v *GetK8sApplicationResponseBodyApplcationImageInfo) *GetK8sApplicationResponseBodyApplcation {
	s.ImageInfo = v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcation) SetApp(v *GetK8sApplicationResponseBodyApplcationApp) *GetK8sApplicationResponseBodyApplcation {
	s.App = v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcation) SetConf(v *GetK8sApplicationResponseBodyApplcationConf) *GetK8sApplicationResponseBodyApplcation {
	s.Conf = v
	return s
}

type GetK8sApplicationResponseBodyApplcationDeployGroups struct {
	DeployGroup []*GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup `json:"DeployGroup,omitempty" xml:"DeployGroup,omitempty" require:"true" type:"Repeated"`
}

func (s GetK8sApplicationResponseBodyApplcationDeployGroups) String() string {
	return tea.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationDeployGroups) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroups) SetDeployGroup(v []*GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup) *GetK8sApplicationResponseBodyApplcationDeployGroups {
	s.DeployGroup = v
	return s
}

type GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup struct {
	Components *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents `json:"Components,omitempty" xml:"Components,omitempty" require:"true" type:"Struct"`
}

func (s GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup) String() string {
	return tea.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup) SetComponents(v *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents) *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup {
	s.Components = v
	return s
}

type GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents struct {
	Components []*GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents `json:"Components,omitempty" xml:"Components,omitempty" require:"true" type:"Repeated"`
}

func (s GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents) String() string {
	return tea.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents) SetComponents(v []*GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents) *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents {
	s.Components = v
	return s
}

type GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents struct {
	ComponentId  *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" require:"true"`
	ComponentKey *string `json:"ComponentKey,omitempty" xml:"ComponentKey,omitempty" require:"true"`
}

func (s GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents) String() string {
	return tea.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents) SetComponentId(v string) *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents {
	s.ComponentId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents) SetComponentKey(v string) *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents {
	s.ComponentKey = &v
	return s
}

type GetK8sApplicationResponseBodyApplcationImageInfo struct {
	ImageUrl       *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" require:"true"`
	RepoName       *string `json:"RepoName,omitempty" xml:"RepoName,omitempty" require:"true"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	RepoId         *string `json:"RepoId,omitempty" xml:"RepoId,omitempty" require:"true"`
	RepoNamespace  *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty" require:"true"`
	RepoOriginType *string `json:"RepoOriginType,omitempty" xml:"RepoOriginType,omitempty" require:"true"`
	Tag            *string `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true"`
}

func (s GetK8sApplicationResponseBodyApplcationImageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationImageInfo) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) SetImageUrl(v string) *GetK8sApplicationResponseBodyApplcationImageInfo {
	s.ImageUrl = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) SetRepoName(v string) *GetK8sApplicationResponseBodyApplcationImageInfo {
	s.RepoName = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) SetRegionId(v string) *GetK8sApplicationResponseBodyApplcationImageInfo {
	s.RegionId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) SetRepoId(v string) *GetK8sApplicationResponseBodyApplcationImageInfo {
	s.RepoId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) SetRepoNamespace(v string) *GetK8sApplicationResponseBodyApplcationImageInfo {
	s.RepoNamespace = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) SetRepoOriginType(v string) *GetK8sApplicationResponseBodyApplcationImageInfo {
	s.RepoOriginType = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) SetTag(v string) *GetK8sApplicationResponseBodyApplcationImageInfo {
	s.Tag = &v
	return s
}

type GetK8sApplicationResponseBodyApplcationApp struct {
	ApplicationType        *string                                            `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty" require:"true"`
	Cmd                    *string                                            `json:"Cmd,omitempty" xml:"Cmd,omitempty" require:"true"`
	DeployType             *string                                            `json:"DeployType,omitempty" xml:"DeployType,omitempty" require:"true"`
	EdasContainerVersion   *string                                            `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty" require:"true"`
	BuildpackId            *int                                               `json:"BuildpackId,omitempty" xml:"BuildpackId,omitempty" require:"true"`
	TomcatVersion          *string                                            `json:"TomcatVersion,omitempty" xml:"TomcatVersion,omitempty" require:"true"`
	InstancesBeforeScaling *int                                               `json:"InstancesBeforeScaling,omitempty" xml:"InstancesBeforeScaling,omitempty" require:"true"`
	AppId                  *string                                            `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	ApplicationName        *string                                            `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty" require:"true"`
	ClusterId              *string                                            `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	Instances              *int                                               `json:"Instances,omitempty" xml:"Instances,omitempty" require:"true"`
	CsClusterId            *string                                            `json:"CsClusterId,omitempty" xml:"CsClusterId,omitempty" require:"true"`
	EnvList                *GetK8sApplicationResponseBodyApplcationAppEnvList `json:"EnvList,omitempty" xml:"EnvList,omitempty" require:"true" type:"Struct"`
	CmdArgs                *GetK8sApplicationResponseBodyApplcationAppCmdArgs `json:"CmdArgs,omitempty" xml:"CmdArgs,omitempty" require:"true" type:"Struct"`
}

func (s GetK8sApplicationResponseBodyApplcationApp) String() string {
	return tea.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationApp) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetApplicationType(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.ApplicationType = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetCmd(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.Cmd = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetDeployType(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.DeployType = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetEdasContainerVersion(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.EdasContainerVersion = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetBuildpackId(v int) *GetK8sApplicationResponseBodyApplcationApp {
	s.BuildpackId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetTomcatVersion(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.TomcatVersion = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetInstancesBeforeScaling(v int) *GetK8sApplicationResponseBodyApplcationApp {
	s.InstancesBeforeScaling = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetAppId(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.AppId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetApplicationName(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.ApplicationName = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetClusterId(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.ClusterId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetInstances(v int) *GetK8sApplicationResponseBodyApplcationApp {
	s.Instances = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetCsClusterId(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.CsClusterId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetEnvList(v *GetK8sApplicationResponseBodyApplcationAppEnvList) *GetK8sApplicationResponseBodyApplcationApp {
	s.EnvList = v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetCmdArgs(v *GetK8sApplicationResponseBodyApplcationAppCmdArgs) *GetK8sApplicationResponseBodyApplcationApp {
	s.CmdArgs = v
	return s
}

type GetK8sApplicationResponseBodyApplcationAppEnvList struct {
	Env []*GetK8sApplicationResponseBodyApplcationAppEnvListEnv `json:"Env,omitempty" xml:"Env,omitempty" require:"true" type:"Repeated"`
}

func (s GetK8sApplicationResponseBodyApplcationAppEnvList) String() string {
	return tea.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationAppEnvList) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationAppEnvList) SetEnv(v []*GetK8sApplicationResponseBodyApplcationAppEnvListEnv) *GetK8sApplicationResponseBodyApplcationAppEnvList {
	s.Env = v
	return s
}

type GetK8sApplicationResponseBodyApplcationAppEnvListEnv struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s GetK8sApplicationResponseBodyApplcationAppEnvListEnv) String() string {
	return tea.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationAppEnvListEnv) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationAppEnvListEnv) SetName(v string) *GetK8sApplicationResponseBodyApplcationAppEnvListEnv {
	s.Name = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationAppEnvListEnv) SetValue(v string) *GetK8sApplicationResponseBodyApplcationAppEnvListEnv {
	s.Value = &v
	return s
}

type GetK8sApplicationResponseBodyApplcationAppCmdArgs struct {
	CmdArg []*string `json:"CmdArg,omitempty" xml:"CmdArg,omitempty" require:"true" type:"Repeated"`
}

func (s GetK8sApplicationResponseBodyApplcationAppCmdArgs) String() string {
	return tea.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationAppCmdArgs) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationAppCmdArgs) SetCmdArg(v []*string) *GetK8sApplicationResponseBodyApplcationAppCmdArgs {
	s.CmdArg = v
	return s
}

type GetK8sApplicationResponseBodyApplcationConf struct {
	JarStartArgs       *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty" require:"true"`
	JarStartOptions    *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty" require:"true"`
	K8sCmd             *string `json:"K8sCmd,omitempty" xml:"K8sCmd,omitempty" require:"true"`
	K8sCmdArgs         *string `json:"K8sCmdArgs,omitempty" xml:"K8sCmdArgs,omitempty" require:"true"`
	K8sLocalvolumeInfo *string `json:"K8sLocalvolumeInfo,omitempty" xml:"K8sLocalvolumeInfo,omitempty" require:"true"`
	K8sNasInfo         *string `json:"K8sNasInfo,omitempty" xml:"K8sNasInfo,omitempty" require:"true"`
	K8sVolumeInfo      *string `json:"K8sVolumeInfo,omitempty" xml:"K8sVolumeInfo,omitempty" require:"true"`
	Liveness           *string `json:"Liveness,omitempty" xml:"Liveness,omitempty" require:"true"`
	PostStart          *string `json:"PostStart,omitempty" xml:"PostStart,omitempty" require:"true"`
	PreStop            *string `json:"PreStop,omitempty" xml:"PreStop,omitempty" require:"true"`
	Readiness          *string `json:"Readiness,omitempty" xml:"Readiness,omitempty" require:"true"`
	RuntimeClassName   *string `json:"RuntimeClassName,omitempty" xml:"RuntimeClassName,omitempty" require:"true"`
	DeployAcrossZones  *string `json:"DeployAcrossZones,omitempty" xml:"DeployAcrossZones,omitempty" require:"true"`
	AhasEnabled        *bool   `json:"AhasEnabled,omitempty" xml:"AhasEnabled,omitempty" require:"true"`
	DeployAcrossNodes  *string `json:"DeployAcrossNodes,omitempty" xml:"DeployAcrossNodes,omitempty" require:"true"`
}

func (s GetK8sApplicationResponseBodyApplcationConf) String() string {
	return tea.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationConf) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetJarStartArgs(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.JarStartArgs = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetJarStartOptions(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.JarStartOptions = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetK8sCmd(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.K8sCmd = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetK8sCmdArgs(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.K8sCmdArgs = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetK8sLocalvolumeInfo(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.K8sLocalvolumeInfo = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetK8sNasInfo(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.K8sNasInfo = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetK8sVolumeInfo(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.K8sVolumeInfo = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetLiveness(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.Liveness = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetPostStart(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.PostStart = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetPreStop(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.PreStop = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetReadiness(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.Readiness = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetRuntimeClassName(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.RuntimeClassName = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetDeployAcrossZones(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.DeployAcrossZones = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetAhasEnabled(v bool) *GetK8sApplicationResponseBodyApplcationConf {
	s.AhasEnabled = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetDeployAcrossNodes(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.DeployAcrossNodes = &v
	return s
}

type GetK8sApplicationResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetK8sApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetK8sApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetK8sApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponse) SetHeaders(v map[string]*string) *GetK8sApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetK8sApplicationResponse) SetBody(v *GetK8sApplicationResponseBody) *GetK8sApplicationResponse {
	s.Body = v
	return s
}

type QueryEccInfoQuery struct {
	EccId *string `json:"EccId,omitempty" xml:"EccId,omitempty" require:"true"`
}

func (s QueryEccInfoQuery) String() string {
	return tea.Prettify(s)
}

func (s QueryEccInfoQuery) GoString() string {
	return s.String()
}

func (s *QueryEccInfoQuery) SetEccId(v string) *QueryEccInfoQuery {
	s.EccId = &v
	return s
}

type QueryEccInfoRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *QueryEccInfoQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s QueryEccInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEccInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryEccInfoRequest) SetHeaders(v map[string]*string) *QueryEccInfoRequest {
	s.Headers = v
	return s
}

func (s *QueryEccInfoRequest) SetQuery(v *QueryEccInfoQuery) *QueryEccInfoRequest {
	s.Query = v
	return s
}

type QueryEccInfoResponseBody struct {
	Code      *int                             `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	EccInfo   *QueryEccInfoResponseBodyEccInfo `json:"EccInfo,omitempty" xml:"EccInfo,omitempty" require:"true" type:"Struct"`
}

func (s QueryEccInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEccInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEccInfoResponseBody) SetCode(v int) *QueryEccInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEccInfoResponseBody) SetMessage(v string) *QueryEccInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryEccInfoResponseBody) SetRequestId(v string) *QueryEccInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEccInfoResponseBody) SetEccInfo(v *QueryEccInfoResponseBodyEccInfo) *QueryEccInfoResponseBody {
	s.EccInfo = v
	return s
}

type QueryEccInfoResponseBodyEccInfo struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	EccId          *string `json:"EccId,omitempty" xml:"EccId,omitempty" require:"true"`
	EcuId          *string `json:"EcuId,omitempty" xml:"EcuId,omitempty" require:"true"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty" require:"true"`
	PackageMd5     *string `json:"PackageMd5,omitempty" xml:"PackageMd5,omitempty" require:"true"`
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty" require:"true"`
	VpcId          *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
}

func (s QueryEccInfoResponseBodyEccInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryEccInfoResponseBodyEccInfo) GoString() string {
	return s.String()
}

func (s *QueryEccInfoResponseBodyEccInfo) SetAppId(v string) *QueryEccInfoResponseBodyEccInfo {
	s.AppId = &v
	return s
}

func (s *QueryEccInfoResponseBodyEccInfo) SetEccId(v string) *QueryEccInfoResponseBodyEccInfo {
	s.EccId = &v
	return s
}

func (s *QueryEccInfoResponseBodyEccInfo) SetEcuId(v string) *QueryEccInfoResponseBodyEccInfo {
	s.EcuId = &v
	return s
}

func (s *QueryEccInfoResponseBodyEccInfo) SetGroupId(v string) *QueryEccInfoResponseBodyEccInfo {
	s.GroupId = &v
	return s
}

func (s *QueryEccInfoResponseBodyEccInfo) SetGroupName(v string) *QueryEccInfoResponseBodyEccInfo {
	s.GroupName = &v
	return s
}

func (s *QueryEccInfoResponseBodyEccInfo) SetPackageMd5(v string) *QueryEccInfoResponseBodyEccInfo {
	s.PackageMd5 = &v
	return s
}

func (s *QueryEccInfoResponseBodyEccInfo) SetPackageVersion(v string) *QueryEccInfoResponseBodyEccInfo {
	s.PackageVersion = &v
	return s
}

func (s *QueryEccInfoResponseBodyEccInfo) SetVpcId(v string) *QueryEccInfoResponseBodyEccInfo {
	s.VpcId = &v
	return s
}

type QueryEccInfoResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryEccInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryEccInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEccInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryEccInfoResponse) SetHeaders(v map[string]*string) *QueryEccInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryEccInfoResponse) SetBody(v *QueryEccInfoResponseBody) *QueryEccInfoResponse {
	s.Body = v
	return s
}

type ContinuePipelineQuery struct {
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty" require:"true"`
	Confirm    *bool   `json:"Confirm,omitempty" xml:"Confirm,omitempty"`
}

func (s ContinuePipelineQuery) String() string {
	return tea.Prettify(s)
}

func (s ContinuePipelineQuery) GoString() string {
	return s.String()
}

func (s *ContinuePipelineQuery) SetPipelineId(v string) *ContinuePipelineQuery {
	s.PipelineId = &v
	return s
}

func (s *ContinuePipelineQuery) SetConfirm(v bool) *ContinuePipelineQuery {
	s.Confirm = &v
	return s
}

type ContinuePipelineRequest struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ContinuePipelineQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ContinuePipelineRequest) String() string {
	return tea.Prettify(s)
}

func (s ContinuePipelineRequest) GoString() string {
	return s.String()
}

func (s *ContinuePipelineRequest) SetHeaders(v map[string]*string) *ContinuePipelineRequest {
	s.Headers = v
	return s
}

func (s *ContinuePipelineRequest) SetQuery(v *ContinuePipelineQuery) *ContinuePipelineRequest {
	s.Query = v
	return s
}

type ContinuePipelineResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ContinuePipelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContinuePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *ContinuePipelineResponseBody) SetCode(v int) *ContinuePipelineResponseBody {
	s.Code = &v
	return s
}

func (s *ContinuePipelineResponseBody) SetMessage(v string) *ContinuePipelineResponseBody {
	s.Message = &v
	return s
}

func (s *ContinuePipelineResponseBody) SetRequestId(v string) *ContinuePipelineResponseBody {
	s.RequestId = &v
	return s
}

type ContinuePipelineResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ContinuePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ContinuePipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s ContinuePipelineResponse) GoString() string {
	return s.String()
}

func (s *ContinuePipelineResponse) SetHeaders(v map[string]*string) *ContinuePipelineResponse {
	s.Headers = v
	return s
}

func (s *ContinuePipelineResponse) SetBody(v *ContinuePipelineResponseBody) *ContinuePipelineResponse {
	s.Body = v
	return s
}

type ChangeDeployGroupQuery struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	EccInfo     *string `json:"EccInfo,omitempty" xml:"EccInfo,omitempty" require:"true"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty" require:"true"`
	ForceStatus *bool   `json:"ForceStatus,omitempty" xml:"ForceStatus,omitempty"`
}

func (s ChangeDeployGroupQuery) String() string {
	return tea.Prettify(s)
}

func (s ChangeDeployGroupQuery) GoString() string {
	return s.String()
}

func (s *ChangeDeployGroupQuery) SetAppId(v string) *ChangeDeployGroupQuery {
	s.AppId = &v
	return s
}

func (s *ChangeDeployGroupQuery) SetEccInfo(v string) *ChangeDeployGroupQuery {
	s.EccInfo = &v
	return s
}

func (s *ChangeDeployGroupQuery) SetGroupName(v string) *ChangeDeployGroupQuery {
	s.GroupName = &v
	return s
}

func (s *ChangeDeployGroupQuery) SetForceStatus(v bool) *ChangeDeployGroupQuery {
	s.ForceStatus = &v
	return s
}

type ChangeDeployGroupRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ChangeDeployGroupQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ChangeDeployGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeDeployGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeDeployGroupRequest) SetHeaders(v map[string]*string) *ChangeDeployGroupRequest {
	s.Headers = v
	return s
}

func (s *ChangeDeployGroupRequest) SetQuery(v *ChangeDeployGroupQuery) *ChangeDeployGroupRequest {
	s.Query = v
	return s
}

type ChangeDeployGroupResponseBody struct {
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ChangeDeployGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeDeployGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeDeployGroupResponseBody) SetCode(v int) *ChangeDeployGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeDeployGroupResponseBody) SetMessage(v string) *ChangeDeployGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeDeployGroupResponseBody) SetChangeOrderId(v string) *ChangeDeployGroupResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *ChangeDeployGroupResponseBody) SetRequestId(v string) *ChangeDeployGroupResponseBody {
	s.RequestId = &v
	return s
}

type ChangeDeployGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChangeDeployGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeDeployGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeDeployGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeDeployGroupResponse) SetHeaders(v map[string]*string) *ChangeDeployGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeDeployGroupResponse) SetBody(v *ChangeDeployGroupResponseBody) *ChangeDeployGroupResponse {
	s.Body = v
	return s
}

type GetClusterQuery struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
}

func (s GetClusterQuery) String() string {
	return tea.Prettify(s)
}

func (s GetClusterQuery) GoString() string {
	return s.String()
}

func (s *GetClusterQuery) SetClusterId(v string) *GetClusterQuery {
	s.ClusterId = &v
	return s
}

type GetClusterRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *GetClusterQuery   `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s GetClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRequest) GoString() string {
	return s.String()
}

func (s *GetClusterRequest) SetHeaders(v map[string]*string) *GetClusterRequest {
	s.Headers = v
	return s
}

func (s *GetClusterRequest) SetQuery(v *GetClusterQuery) *GetClusterRequest {
	s.Query = v
	return s
}

type GetClusterResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Cluster   *GetClusterResponseBodyCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" require:"true" type:"Struct"`
}

func (s GetClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBody) SetRequestId(v string) *GetClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterResponseBody) SetCode(v int) *GetClusterResponseBody {
	s.Code = &v
	return s
}

func (s *GetClusterResponseBody) SetMessage(v string) *GetClusterResponseBody {
	s.Message = &v
	return s
}

func (s *GetClusterResponseBody) SetCluster(v *GetClusterResponseBodyCluster) *GetClusterResponseBody {
	s.Cluster = v
	return s
}

type GetClusterResponseBodyCluster struct {
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	ClusterName         *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty" require:"true"`
	ClusterType         *int    `json:"ClusterType,omitempty" xml:"ClusterType,omitempty" require:"true"`
	OversoldFactor      *int    `json:"OversoldFactor,omitempty" xml:"OversoldFactor,omitempty" require:"true"`
	NetworkMode         *int    `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty" require:"true"`
	VpcId               *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	NodeNum             *int    `json:"NodeNum,omitempty" xml:"NodeNum,omitempty" require:"true"`
	Cpu                 *int    `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Mem                 *int    `json:"Mem,omitempty" xml:"Mem,omitempty" require:"true"`
	CpuUsed             *int    `json:"CpuUsed,omitempty" xml:"CpuUsed,omitempty" require:"true"`
	MemUsed             *int    `json:"MemUsed,omitempty" xml:"MemUsed,omitempty" require:"true"`
	CreateTime          *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime          *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	IaasProvider        *string `json:"IaasProvider,omitempty" xml:"IaasProvider,omitempty" require:"true"`
	CsClusterId         *string `json:"CsClusterId,omitempty" xml:"CsClusterId,omitempty" require:"true"`
	ClusterImportStatus *int    `json:"ClusterImportStatus,omitempty" xml:"ClusterImportStatus,omitempty" require:"true"`
}

func (s GetClusterResponseBodyCluster) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyCluster) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyCluster) SetClusterId(v string) *GetClusterResponseBodyCluster {
	s.ClusterId = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetRegionId(v string) *GetClusterResponseBodyCluster {
	s.RegionId = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetDescription(v string) *GetClusterResponseBodyCluster {
	s.Description = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetClusterName(v string) *GetClusterResponseBodyCluster {
	s.ClusterName = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetClusterType(v int) *GetClusterResponseBodyCluster {
	s.ClusterType = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetOversoldFactor(v int) *GetClusterResponseBodyCluster {
	s.OversoldFactor = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetNetworkMode(v int) *GetClusterResponseBodyCluster {
	s.NetworkMode = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetVpcId(v string) *GetClusterResponseBodyCluster {
	s.VpcId = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetNodeNum(v int) *GetClusterResponseBodyCluster {
	s.NodeNum = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetCpu(v int) *GetClusterResponseBodyCluster {
	s.Cpu = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetMem(v int) *GetClusterResponseBodyCluster {
	s.Mem = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetCpuUsed(v int) *GetClusterResponseBodyCluster {
	s.CpuUsed = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetMemUsed(v int) *GetClusterResponseBodyCluster {
	s.MemUsed = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetCreateTime(v int64) *GetClusterResponseBodyCluster {
	s.CreateTime = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetUpdateTime(v int64) *GetClusterResponseBodyCluster {
	s.UpdateTime = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetIaasProvider(v string) *GetClusterResponseBodyCluster {
	s.IaasProvider = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetCsClusterId(v string) *GetClusterResponseBodyCluster {
	s.CsClusterId = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetClusterImportStatus(v int) *GetClusterResponseBodyCluster {
	s.ClusterImportStatus = &v
	return s
}

type GetClusterResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponse) GoString() string {
	return s.String()
}

func (s *GetClusterResponse) SetHeaders(v map[string]*string) *GetClusterResponse {
	s.Headers = v
	return s
}

func (s *GetClusterResponse) SetBody(v *GetClusterResponseBody) *GetClusterResponse {
	s.Body = v
	return s
}

type QueryRegionConfigRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s QueryRegionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRegionConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryRegionConfigRequest) SetHeaders(v map[string]*string) *QueryRegionConfigRequest {
	s.Headers = v
	return s
}

type QueryRegionConfigResponseBody struct {
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code         *int                                       `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message      *string                                    `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RegionConfig *QueryRegionConfigResponseBodyRegionConfig `json:"RegionConfig,omitempty" xml:"RegionConfig,omitempty" require:"true" type:"Struct"`
}

func (s QueryRegionConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRegionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRegionConfigResponseBody) SetRequestId(v string) *QueryRegionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRegionConfigResponseBody) SetCode(v int) *QueryRegionConfigResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRegionConfigResponseBody) SetMessage(v string) *QueryRegionConfigResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRegionConfigResponseBody) SetRegionConfig(v *QueryRegionConfigResponseBodyRegionConfig) *QueryRegionConfigResponseBody {
	s.RegionConfig = v
	return s
}

type QueryRegionConfigResponseBodyRegionConfig struct {
	AddressServerHost  *string                                                    `json:"AddressServerHost,omitempty" xml:"AddressServerHost,omitempty" require:"true"`
	AgentInstallScript *string                                                    `json:"AgentInstallScript,omitempty" xml:"AgentInstallScript,omitempty" require:"true"`
	FileServerType     *string                                                    `json:"FileServerType,omitempty" xml:"FileServerType,omitempty" require:"true"`
	Id                 *string                                                    `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	ImageId            *string                                                    `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	Name               *string                                                    `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	No                 *int                                                       `json:"No,omitempty" xml:"No,omitempty" require:"true"`
	Tag                *string                                                    `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true"`
	FileServerConfig   *QueryRegionConfigResponseBodyRegionConfigFileServerConfig `json:"FileServerConfig,omitempty" xml:"FileServerConfig,omitempty" require:"true" type:"Struct"`
}

func (s QueryRegionConfigResponseBodyRegionConfig) String() string {
	return tea.Prettify(s)
}

func (s QueryRegionConfigResponseBodyRegionConfig) GoString() string {
	return s.String()
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetAddressServerHost(v string) *QueryRegionConfigResponseBodyRegionConfig {
	s.AddressServerHost = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetAgentInstallScript(v string) *QueryRegionConfigResponseBodyRegionConfig {
	s.AgentInstallScript = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetFileServerType(v string) *QueryRegionConfigResponseBodyRegionConfig {
	s.FileServerType = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetId(v string) *QueryRegionConfigResponseBodyRegionConfig {
	s.Id = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetImageId(v string) *QueryRegionConfigResponseBodyRegionConfig {
	s.ImageId = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetName(v string) *QueryRegionConfigResponseBodyRegionConfig {
	s.Name = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetNo(v int) *QueryRegionConfigResponseBodyRegionConfig {
	s.No = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetTag(v string) *QueryRegionConfigResponseBodyRegionConfig {
	s.Tag = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfig) SetFileServerConfig(v *QueryRegionConfigResponseBodyRegionConfigFileServerConfig) *QueryRegionConfigResponseBodyRegionConfig {
	s.FileServerConfig = v
	return s
}

type QueryRegionConfigResponseBodyRegionConfigFileServerConfig struct {
	Bucket      *string `json:"Bucket,omitempty" xml:"Bucket,omitempty" require:"true"`
	InternalUrl *string `json:"InternalUrl,omitempty" xml:"InternalUrl,omitempty" require:"true"`
	PublicUrl   *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty" require:"true"`
	VpcUrl      *string `json:"VpcUrl,omitempty" xml:"VpcUrl,omitempty" require:"true"`
}

func (s QueryRegionConfigResponseBodyRegionConfigFileServerConfig) String() string {
	return tea.Prettify(s)
}

func (s QueryRegionConfigResponseBodyRegionConfigFileServerConfig) GoString() string {
	return s.String()
}

func (s *QueryRegionConfigResponseBodyRegionConfigFileServerConfig) SetBucket(v string) *QueryRegionConfigResponseBodyRegionConfigFileServerConfig {
	s.Bucket = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfigFileServerConfig) SetInternalUrl(v string) *QueryRegionConfigResponseBodyRegionConfigFileServerConfig {
	s.InternalUrl = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfigFileServerConfig) SetPublicUrl(v string) *QueryRegionConfigResponseBodyRegionConfigFileServerConfig {
	s.PublicUrl = &v
	return s
}

func (s *QueryRegionConfigResponseBodyRegionConfigFileServerConfig) SetVpcUrl(v string) *QueryRegionConfigResponseBodyRegionConfigFileServerConfig {
	s.VpcUrl = &v
	return s
}

type QueryRegionConfigResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRegionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRegionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRegionConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryRegionConfigResponse) SetHeaders(v map[string]*string) *QueryRegionConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryRegionConfigResponse) SetBody(v *QueryRegionConfigResponseBody) *QueryRegionConfigResponse {
	s.Body = v
	return s
}

type SynchronizeResourceQuery struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
}

func (s SynchronizeResourceQuery) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeResourceQuery) GoString() string {
	return s.String()
}

func (s *SynchronizeResourceQuery) SetType(v string) *SynchronizeResourceQuery {
	s.Type = &v
	return s
}

type SynchronizeResourceRequest struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *SynchronizeResourceQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s SynchronizeResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeResourceRequest) GoString() string {
	return s.String()
}

func (s *SynchronizeResourceRequest) SetHeaders(v map[string]*string) *SynchronizeResourceRequest {
	s.Headers = v
	return s
}

func (s *SynchronizeResourceRequest) SetQuery(v *SynchronizeResourceQuery) *SynchronizeResourceRequest {
	s.Query = v
	return s
}

type SynchronizeResourceResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SynchronizeResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeResourceResponseBody) GoString() string {
	return s.String()
}

func (s *SynchronizeResourceResponseBody) SetCode(v int) *SynchronizeResourceResponseBody {
	s.Code = &v
	return s
}

func (s *SynchronizeResourceResponseBody) SetMessage(v string) *SynchronizeResourceResponseBody {
	s.Message = &v
	return s
}

func (s *SynchronizeResourceResponseBody) SetSuccess(v bool) *SynchronizeResourceResponseBody {
	s.Success = &v
	return s
}

func (s *SynchronizeResourceResponseBody) SetData(v string) *SynchronizeResourceResponseBody {
	s.Data = &v
	return s
}

func (s *SynchronizeResourceResponseBody) SetRequestId(v string) *SynchronizeResourceResponseBody {
	s.RequestId = &v
	return s
}

type SynchronizeResourceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SynchronizeResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SynchronizeResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeResourceResponse) GoString() string {
	return s.String()
}

func (s *SynchronizeResourceResponse) SetHeaders(v map[string]*string) *SynchronizeResourceResponse {
	s.Headers = v
	return s
}

func (s *SynchronizeResourceResponse) SetBody(v *SynchronizeResourceResponseBody) *SynchronizeResourceResponse {
	s.Body = v
	return s
}

type InstallAgentQuery struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" require:"true"`
	DoAsync     *bool   `json:"DoAsync,omitempty" xml:"DoAsync,omitempty"`
}

func (s InstallAgentQuery) String() string {
	return tea.Prettify(s)
}

func (s InstallAgentQuery) GoString() string {
	return s.String()
}

func (s *InstallAgentQuery) SetClusterId(v string) *InstallAgentQuery {
	s.ClusterId = &v
	return s
}

func (s *InstallAgentQuery) SetInstanceIds(v string) *InstallAgentQuery {
	s.InstanceIds = &v
	return s
}

func (s *InstallAgentQuery) SetDoAsync(v bool) *InstallAgentQuery {
	s.DoAsync = &v
	return s
}

type InstallAgentRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *InstallAgentQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s InstallAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallAgentRequest) GoString() string {
	return s.String()
}

func (s *InstallAgentRequest) SetHeaders(v map[string]*string) *InstallAgentRequest {
	s.Headers = v
	return s
}

func (s *InstallAgentRequest) SetQuery(v *InstallAgentQuery) *InstallAgentRequest {
	s.Query = v
	return s
}

type InstallAgentResponseBody struct {
	RequestId           *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code                *int                                         `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message             *string                                      `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ExecutionResultList *InstallAgentResponseBodyExecutionResultList `json:"ExecutionResultList,omitempty" xml:"ExecutionResultList,omitempty" require:"true" type:"Struct"`
}

func (s InstallAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallAgentResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAgentResponseBody) SetRequestId(v string) *InstallAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallAgentResponseBody) SetCode(v int) *InstallAgentResponseBody {
	s.Code = &v
	return s
}

func (s *InstallAgentResponseBody) SetMessage(v string) *InstallAgentResponseBody {
	s.Message = &v
	return s
}

func (s *InstallAgentResponseBody) SetExecutionResultList(v *InstallAgentResponseBodyExecutionResultList) *InstallAgentResponseBody {
	s.ExecutionResultList = v
	return s
}

type InstallAgentResponseBodyExecutionResultList struct {
	ExecutionResult []*InstallAgentResponseBodyExecutionResultListExecutionResult `json:"ExecutionResult,omitempty" xml:"ExecutionResult,omitempty" require:"true" type:"Repeated"`
}

func (s InstallAgentResponseBodyExecutionResultList) String() string {
	return tea.Prettify(s)
}

func (s InstallAgentResponseBodyExecutionResultList) GoString() string {
	return s.String()
}

func (s *InstallAgentResponseBodyExecutionResultList) SetExecutionResult(v []*InstallAgentResponseBodyExecutionResultListExecutionResult) *InstallAgentResponseBodyExecutionResultList {
	s.ExecutionResult = v
	return s
}

type InstallAgentResponseBodyExecutionResultListExecutionResult struct {
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	FinishedTime       *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty" require:"true"`
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty" require:"true"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s InstallAgentResponseBodyExecutionResultListExecutionResult) String() string {
	return tea.Prettify(s)
}

func (s InstallAgentResponseBodyExecutionResultListExecutionResult) GoString() string {
	return s.String()
}

func (s *InstallAgentResponseBodyExecutionResultListExecutionResult) SetInstanceId(v string) *InstallAgentResponseBodyExecutionResultListExecutionResult {
	s.InstanceId = &v
	return s
}

func (s *InstallAgentResponseBodyExecutionResultListExecutionResult) SetStatus(v string) *InstallAgentResponseBodyExecutionResultListExecutionResult {
	s.Status = &v
	return s
}

func (s *InstallAgentResponseBodyExecutionResultListExecutionResult) SetFinishedTime(v string) *InstallAgentResponseBodyExecutionResultListExecutionResult {
	s.FinishedTime = &v
	return s
}

func (s *InstallAgentResponseBodyExecutionResultListExecutionResult) SetInvokeRecordStatus(v string) *InstallAgentResponseBodyExecutionResultListExecutionResult {
	s.InvokeRecordStatus = &v
	return s
}

func (s *InstallAgentResponseBodyExecutionResultListExecutionResult) SetSuccess(v bool) *InstallAgentResponseBodyExecutionResultListExecutionResult {
	s.Success = &v
	return s
}

type InstallAgentResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InstallAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallAgentResponse) GoString() string {
	return s.String()
}

func (s *InstallAgentResponse) SetHeaders(v map[string]*string) *InstallAgentResponse {
	s.Headers = v
	return s
}

func (s *InstallAgentResponse) SetBody(v *InstallAgentResponseBody) *InstallAgentResponse {
	s.Body = v
	return s
}

type ListComponentsRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ListComponentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsRequest) GoString() string {
	return s.String()
}

func (s *ListComponentsRequest) SetHeaders(v map[string]*string) *ListComponentsRequest {
	s.Headers = v
	return s
}

type ListComponentsResponseBody struct {
	Code          *int                                     `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string                                  `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ComponentList *ListComponentsResponseBodyComponentList `json:"ComponentList,omitempty" xml:"ComponentList,omitempty" require:"true" type:"Struct"`
}

func (s ListComponentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBody) SetCode(v int) *ListComponentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListComponentsResponseBody) SetMessage(v string) *ListComponentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListComponentsResponseBody) SetComponentList(v *ListComponentsResponseBodyComponentList) *ListComponentsResponseBody {
	s.ComponentList = v
	return s
}

type ListComponentsResponseBodyComponentList struct {
	Component []*ListComponentsResponseBodyComponentListComponent `json:"Component,omitempty" xml:"Component,omitempty" require:"true" type:"Repeated"`
}

func (s ListComponentsResponseBodyComponentList) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyComponentList) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyComponentList) SetComponent(v []*ListComponentsResponseBodyComponentListComponent) *ListComponentsResponseBodyComponentList {
	s.Component = v
	return s
}

type ListComponentsResponseBodyComponentListComponent struct {
	ComponentId  *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" require:"true"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Expired      *bool   `json:"Expired,omitempty" xml:"Expired,omitempty" require:"true"`
	ComponentKey *string `json:"ComponentKey,omitempty" xml:"ComponentKey,omitempty" require:"true"`
	Desc         *string `json:"Desc,omitempty" xml:"Desc,omitempty" require:"true"`
}

func (s ListComponentsResponseBodyComponentListComponent) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyComponentListComponent) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyComponentListComponent) SetComponentId(v string) *ListComponentsResponseBodyComponentListComponent {
	s.ComponentId = &v
	return s
}

func (s *ListComponentsResponseBodyComponentListComponent) SetType(v string) *ListComponentsResponseBodyComponentListComponent {
	s.Type = &v
	return s
}

func (s *ListComponentsResponseBodyComponentListComponent) SetVersion(v string) *ListComponentsResponseBodyComponentListComponent {
	s.Version = &v
	return s
}

func (s *ListComponentsResponseBodyComponentListComponent) SetExpired(v bool) *ListComponentsResponseBodyComponentListComponent {
	s.Expired = &v
	return s
}

func (s *ListComponentsResponseBodyComponentListComponent) SetComponentKey(v string) *ListComponentsResponseBodyComponentListComponent {
	s.ComponentKey = &v
	return s
}

func (s *ListComponentsResponseBodyComponentListComponent) SetDesc(v string) *ListComponentsResponseBodyComponentListComponent {
	s.Desc = &v
	return s
}

type ListComponentsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListComponentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponse) GoString() string {
	return s.String()
}

func (s *ListComponentsResponse) SetHeaders(v map[string]*string) *ListComponentsResponse {
	s.Headers = v
	return s
}

func (s *ListComponentsResponse) SetBody(v *ListComponentsResponseBody) *ListComponentsResponse {
	s.Body = v
	return s
}

type GetPackageStorageCredentialRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s GetPackageStorageCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPackageStorageCredentialRequest) GoString() string {
	return s.String()
}

func (s *GetPackageStorageCredentialRequest) SetHeaders(v map[string]*string) *GetPackageStorageCredentialRequest {
	s.Headers = v
	return s
}

type GetPackageStorageCredentialResponseBody struct {
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code       *int                                               `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message    *string                                            `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Credential *GetPackageStorageCredentialResponseBodyCredential `json:"Credential,omitempty" xml:"Credential,omitempty" require:"true" type:"Struct"`
}

func (s GetPackageStorageCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPackageStorageCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetPackageStorageCredentialResponseBody) SetRequestId(v string) *GetPackageStorageCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBody) SetCode(v int) *GetPackageStorageCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBody) SetMessage(v string) *GetPackageStorageCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBody) SetCredential(v *GetPackageStorageCredentialResponseBodyCredential) *GetPackageStorageCredentialResponseBody {
	s.Credential = v
	return s
}

type GetPackageStorageCredentialResponseBodyCredential struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty" require:"true"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty" require:"true"`
	Bucket          *string `json:"Bucket,omitempty" xml:"Bucket,omitempty" require:"true"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty" require:"true"`
	KeyPrefix       *string `json:"KeyPrefix,omitempty" xml:"KeyPrefix,omitempty" require:"true"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty" require:"true"`
}

func (s GetPackageStorageCredentialResponseBodyCredential) String() string {
	return tea.Prettify(s)
}

func (s GetPackageStorageCredentialResponseBodyCredential) GoString() string {
	return s.String()
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetAccessKeyId(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.AccessKeyId = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetAccessKeySecret(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.AccessKeySecret = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetBucket(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.Bucket = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetExpiration(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.Expiration = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetKeyPrefix(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.KeyPrefix = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetRegionId(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.RegionId = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetSecurityToken(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.SecurityToken = &v
	return s
}

type GetPackageStorageCredentialResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPackageStorageCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPackageStorageCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPackageStorageCredentialResponse) GoString() string {
	return s.String()
}

func (s *GetPackageStorageCredentialResponse) SetHeaders(v map[string]*string) *GetPackageStorageCredentialResponse {
	s.Headers = v
	return s
}

func (s *GetPackageStorageCredentialResponse) SetBody(v *GetPackageStorageCredentialResponseBody) *GetPackageStorageCredentialResponse {
	s.Body = v
	return s
}

type ListEcsNotInClusterQuery struct {
	NetworkMode *int    `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty" require:"true"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListEcsNotInClusterQuery) String() string {
	return tea.Prettify(s)
}

func (s ListEcsNotInClusterQuery) GoString() string {
	return s.String()
}

func (s *ListEcsNotInClusterQuery) SetNetworkMode(v int) *ListEcsNotInClusterQuery {
	s.NetworkMode = &v
	return s
}

func (s *ListEcsNotInClusterQuery) SetVpcId(v string) *ListEcsNotInClusterQuery {
	s.VpcId = &v
	return s
}

type ListEcsNotInClusterRequest struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListEcsNotInClusterQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ListEcsNotInClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEcsNotInClusterRequest) GoString() string {
	return s.String()
}

func (s *ListEcsNotInClusterRequest) SetHeaders(v map[string]*string) *ListEcsNotInClusterRequest {
	s.Headers = v
	return s
}

func (s *ListEcsNotInClusterRequest) SetQuery(v *ListEcsNotInClusterQuery) *ListEcsNotInClusterRequest {
	s.Query = v
	return s
}

type ListEcsNotInClusterResponseBody struct {
	Code          *int                                          `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string                                       `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	EcsEntityList *ListEcsNotInClusterResponseBodyEcsEntityList `json:"EcsEntityList,omitempty" xml:"EcsEntityList,omitempty" require:"true" type:"Struct"`
}

func (s ListEcsNotInClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEcsNotInClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ListEcsNotInClusterResponseBody) SetCode(v int) *ListEcsNotInClusterResponseBody {
	s.Code = &v
	return s
}

func (s *ListEcsNotInClusterResponseBody) SetMessage(v string) *ListEcsNotInClusterResponseBody {
	s.Message = &v
	return s
}

func (s *ListEcsNotInClusterResponseBody) SetRequestId(v string) *ListEcsNotInClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEcsNotInClusterResponseBody) SetEcsEntityList(v *ListEcsNotInClusterResponseBodyEcsEntityList) *ListEcsNotInClusterResponseBody {
	s.EcsEntityList = v
	return s
}

type ListEcsNotInClusterResponseBodyEcsEntityList struct {
	EcsEntity []*ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity `json:"EcsEntity,omitempty" xml:"EcsEntity,omitempty" require:"true" type:"Repeated"`
}

func (s ListEcsNotInClusterResponseBodyEcsEntityList) String() string {
	return tea.Prettify(s)
}

func (s ListEcsNotInClusterResponseBodyEcsEntityList) GoString() string {
	return s.String()
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityList) SetEcsEntity(v []*ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) *ListEcsNotInClusterResponseBodyEcsEntityList {
	s.EcsEntity = v
	return s
}

type ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty" require:"true"`
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	VpcName      *string `json:"VpcName,omitempty" xml:"VpcName,omitempty" require:"true"`
	Expired      *bool   `json:"Expired,omitempty" xml:"Expired,omitempty" require:"true"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Cpu          *int    `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Mem          *int    `json:"Mem,omitempty" xml:"Mem,omitempty" require:"true"`
	PublicIp     *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty" require:"true"`
	InnerIp      *string `json:"InnerIp,omitempty" xml:"InnerIp,omitempty" require:"true"`
	PrivateIp    *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty" require:"true"`
	Eip          *string `json:"Eip,omitempty" xml:"Eip,omitempty" require:"true"`
}

func (s ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) String() string {
	return tea.Prettify(s)
}

func (s ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) GoString() string {
	return s.String()
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetInstanceId(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.InstanceId = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetInstanceName(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.InstanceName = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetVpcId(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.VpcId = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetVpcName(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.VpcName = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetExpired(v bool) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.Expired = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetStatus(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.Status = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetRegionId(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.RegionId = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetCpu(v int) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.Cpu = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetMem(v int) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.Mem = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetPublicIp(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.PublicIp = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetInnerIp(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.InnerIp = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetPrivateIp(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.PrivateIp = &v
	return s
}

func (s *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity) SetEip(v string) *ListEcsNotInClusterResponseBodyEcsEntityListEcsEntity {
	s.Eip = &v
	return s
}

type ListEcsNotInClusterResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEcsNotInClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEcsNotInClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEcsNotInClusterResponse) GoString() string {
	return s.String()
}

func (s *ListEcsNotInClusterResponse) SetHeaders(v map[string]*string) *ListEcsNotInClusterResponse {
	s.Headers = v
	return s
}

func (s *ListEcsNotInClusterResponse) SetBody(v *ListEcsNotInClusterResponseBody) *ListEcsNotInClusterResponse {
	s.Body = v
	return s
}

type InsertK8sApplicationQuery struct {
	IntranetTargetPort     *int    `json:"IntranetTargetPort,omitempty" xml:"IntranetTargetPort,omitempty"`
	IntranetSlbPort        *int    `json:"IntranetSlbPort,omitempty" xml:"IntranetSlbPort,omitempty"`
	ImageUrl               *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ApplicationDescription *string `json:"ApplicationDescription,omitempty" xml:"ApplicationDescription,omitempty"`
	RepoId                 *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	Replicas               *int    `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	LimitCpu               *int    `json:"LimitCpu,omitempty" xml:"LimitCpu,omitempty"`
	LimitMem               *int    `json:"LimitMem,omitempty" xml:"LimitMem,omitempty"`
	RequestsCpu            *int    `json:"RequestsCpu,omitempty" xml:"RequestsCpu,omitempty"`
	RequestsMem            *int    `json:"RequestsMem,omitempty" xml:"RequestsMem,omitempty"`
	Command                *string `json:"Command,omitempty" xml:"Command,omitempty"`
	CommandArgs            *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	AppName                *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	IntranetSlbProtocol    *string `json:"IntranetSlbProtocol,omitempty" xml:"IntranetSlbProtocol,omitempty"`
	IntranetSlbId          *string `json:"IntranetSlbId,omitempty" xml:"IntranetSlbId,omitempty"`
	ClusterId              *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	InternetSlbId          *string `json:"InternetSlbId,omitempty" xml:"InternetSlbId,omitempty"`
	InternetSlbProtocol    *string `json:"InternetSlbProtocol,omitempty" xml:"InternetSlbProtocol,omitempty"`
	InternetSlbPort        *int    `json:"InternetSlbPort,omitempty" xml:"InternetSlbPort,omitempty"`
	InternetTargetPort     *int    `json:"InternetTargetPort,omitempty" xml:"InternetTargetPort,omitempty"`
	Envs                   *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	PreStop                *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	PostStart              *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	Liveness               *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	Readiness              *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	NasId                  *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	MountDescs             *string `json:"MountDescs,omitempty" xml:"MountDescs,omitempty"`
	StorageType            *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	LocalVolume            *string `json:"LocalVolume,omitempty" xml:"LocalVolume,omitempty"`
	Namespace              *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	LogicalRegionId        *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	PackageType            *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	PackageUrl             *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	PackageVersion         *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	JDK                    *string `json:"JDK,omitempty" xml:"JDK,omitempty"`
	WebContainer           *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
	EdasContainerVersion   *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	UriEncoding            *string `json:"UriEncoding,omitempty" xml:"UriEncoding,omitempty"`
	UseBodyEncoding        *bool   `json:"UseBodyEncoding,omitempty" xml:"UseBodyEncoding,omitempty"`
	RequestsmCpu           *int    `json:"RequestsmCpu,omitempty" xml:"RequestsmCpu,omitempty"`
	LimitmCpu              *int    `json:"LimitmCpu,omitempty" xml:"LimitmCpu,omitempty"`
	RuntimeClassName       *string `json:"RuntimeClassName,omitempty" xml:"RuntimeClassName,omitempty"`
	DeployAcrossZones      *string `json:"DeployAcrossZones,omitempty" xml:"DeployAcrossZones,omitempty"`
	Timeout                *int    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	EnableAhas             *bool   `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
	WebContainerConfig     *string `json:"WebContainerConfig,omitempty" xml:"WebContainerConfig,omitempty"`
	JavaStartUpConfig      *string `json:"JavaStartUpConfig,omitempty" xml:"JavaStartUpConfig,omitempty"`
	SlsConfigs             *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	IsMultilingualApp      *bool   `json:"IsMultilingualApp,omitempty" xml:"IsMultilingualApp,omitempty"`
	DeployAcrossNodes      *string `json:"DeployAcrossNodes,omitempty" xml:"DeployAcrossNodes,omitempty"`
}

func (s InsertK8sApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s InsertK8sApplicationQuery) GoString() string {
	return s.String()
}

func (s *InsertK8sApplicationQuery) SetIntranetTargetPort(v int) *InsertK8sApplicationQuery {
	s.IntranetTargetPort = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetIntranetSlbPort(v int) *InsertK8sApplicationQuery {
	s.IntranetSlbPort = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetImageUrl(v string) *InsertK8sApplicationQuery {
	s.ImageUrl = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetApplicationDescription(v string) *InsertK8sApplicationQuery {
	s.ApplicationDescription = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetRepoId(v string) *InsertK8sApplicationQuery {
	s.RepoId = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetReplicas(v int) *InsertK8sApplicationQuery {
	s.Replicas = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetLimitCpu(v int) *InsertK8sApplicationQuery {
	s.LimitCpu = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetLimitMem(v int) *InsertK8sApplicationQuery {
	s.LimitMem = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetRequestsCpu(v int) *InsertK8sApplicationQuery {
	s.RequestsCpu = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetRequestsMem(v int) *InsertK8sApplicationQuery {
	s.RequestsMem = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetCommand(v string) *InsertK8sApplicationQuery {
	s.Command = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetCommandArgs(v string) *InsertK8sApplicationQuery {
	s.CommandArgs = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetAppName(v string) *InsertK8sApplicationQuery {
	s.AppName = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetIntranetSlbProtocol(v string) *InsertK8sApplicationQuery {
	s.IntranetSlbProtocol = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetIntranetSlbId(v string) *InsertK8sApplicationQuery {
	s.IntranetSlbId = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetClusterId(v string) *InsertK8sApplicationQuery {
	s.ClusterId = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetInternetSlbId(v string) *InsertK8sApplicationQuery {
	s.InternetSlbId = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetInternetSlbProtocol(v string) *InsertK8sApplicationQuery {
	s.InternetSlbProtocol = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetInternetSlbPort(v int) *InsertK8sApplicationQuery {
	s.InternetSlbPort = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetInternetTargetPort(v int) *InsertK8sApplicationQuery {
	s.InternetTargetPort = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetEnvs(v string) *InsertK8sApplicationQuery {
	s.Envs = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetPreStop(v string) *InsertK8sApplicationQuery {
	s.PreStop = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetPostStart(v string) *InsertK8sApplicationQuery {
	s.PostStart = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetLiveness(v string) *InsertK8sApplicationQuery {
	s.Liveness = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetReadiness(v string) *InsertK8sApplicationQuery {
	s.Readiness = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetNasId(v string) *InsertK8sApplicationQuery {
	s.NasId = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetMountDescs(v string) *InsertK8sApplicationQuery {
	s.MountDescs = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetStorageType(v string) *InsertK8sApplicationQuery {
	s.StorageType = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetLocalVolume(v string) *InsertK8sApplicationQuery {
	s.LocalVolume = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetNamespace(v string) *InsertK8sApplicationQuery {
	s.Namespace = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetLogicalRegionId(v string) *InsertK8sApplicationQuery {
	s.LogicalRegionId = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetPackageType(v string) *InsertK8sApplicationQuery {
	s.PackageType = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetPackageUrl(v string) *InsertK8sApplicationQuery {
	s.PackageUrl = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetPackageVersion(v string) *InsertK8sApplicationQuery {
	s.PackageVersion = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetJDK(v string) *InsertK8sApplicationQuery {
	s.JDK = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetWebContainer(v string) *InsertK8sApplicationQuery {
	s.WebContainer = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetEdasContainerVersion(v string) *InsertK8sApplicationQuery {
	s.EdasContainerVersion = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetUriEncoding(v string) *InsertK8sApplicationQuery {
	s.UriEncoding = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetUseBodyEncoding(v bool) *InsertK8sApplicationQuery {
	s.UseBodyEncoding = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetRequestsmCpu(v int) *InsertK8sApplicationQuery {
	s.RequestsmCpu = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetLimitmCpu(v int) *InsertK8sApplicationQuery {
	s.LimitmCpu = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetRuntimeClassName(v string) *InsertK8sApplicationQuery {
	s.RuntimeClassName = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetDeployAcrossZones(v string) *InsertK8sApplicationQuery {
	s.DeployAcrossZones = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetTimeout(v int) *InsertK8sApplicationQuery {
	s.Timeout = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetEnableAhas(v bool) *InsertK8sApplicationQuery {
	s.EnableAhas = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetWebContainerConfig(v string) *InsertK8sApplicationQuery {
	s.WebContainerConfig = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetJavaStartUpConfig(v string) *InsertK8sApplicationQuery {
	s.JavaStartUpConfig = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetSlsConfigs(v string) *InsertK8sApplicationQuery {
	s.SlsConfigs = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetIsMultilingualApp(v bool) *InsertK8sApplicationQuery {
	s.IsMultilingualApp = &v
	return s
}

func (s *InsertK8sApplicationQuery) SetDeployAcrossNodes(v string) *InsertK8sApplicationQuery {
	s.DeployAcrossNodes = &v
	return s
}

type InsertK8sApplicationRequest struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *InsertK8sApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s InsertK8sApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *InsertK8sApplicationRequest) SetHeaders(v map[string]*string) *InsertK8sApplicationRequest {
	s.Headers = v
	return s
}

func (s *InsertK8sApplicationRequest) SetQuery(v *InsertK8sApplicationQuery) *InsertK8sApplicationRequest {
	s.Query = v
	return s
}

type InsertK8sApplicationResponseBody struct {
	Code            *int                                             `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message         *string                                          `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId       *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ApplicationInfo *InsertK8sApplicationResponseBodyApplicationInfo `json:"ApplicationInfo,omitempty" xml:"ApplicationInfo,omitempty" require:"true" type:"Struct"`
}

func (s InsertK8sApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertK8sApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *InsertK8sApplicationResponseBody) SetCode(v int) *InsertK8sApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *InsertK8sApplicationResponseBody) SetMessage(v string) *InsertK8sApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *InsertK8sApplicationResponseBody) SetRequestId(v string) *InsertK8sApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertK8sApplicationResponseBody) SetApplicationInfo(v *InsertK8sApplicationResponseBodyApplicationInfo) *InsertK8sApplicationResponseBody {
	s.ApplicationInfo = v
	return s
}

type InsertK8sApplicationResponseBodyApplicationInfo struct {
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	EdasId        *string `json:"EdasId,omitempty" xml:"EdasId,omitempty" require:"true"`
	Owner         *string `json:"Owner,omitempty" xml:"Owner,omitempty" require:"true"`
	Dockerize     *bool   `json:"Dockerize,omitempty" xml:"Dockerize,omitempty" require:"true"`
	ClusterType   *int    `json:"ClusterType,omitempty" xml:"ClusterType,omitempty" require:"true"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s InsertK8sApplicationResponseBodyApplicationInfo) String() string {
	return tea.Prettify(s)
}

func (s InsertK8sApplicationResponseBodyApplicationInfo) GoString() string {
	return s.String()
}

func (s *InsertK8sApplicationResponseBodyApplicationInfo) SetAppName(v string) *InsertK8sApplicationResponseBodyApplicationInfo {
	s.AppName = &v
	return s
}

func (s *InsertK8sApplicationResponseBodyApplicationInfo) SetAppId(v string) *InsertK8sApplicationResponseBodyApplicationInfo {
	s.AppId = &v
	return s
}

func (s *InsertK8sApplicationResponseBodyApplicationInfo) SetUserId(v string) *InsertK8sApplicationResponseBodyApplicationInfo {
	s.UserId = &v
	return s
}

func (s *InsertK8sApplicationResponseBodyApplicationInfo) SetEdasId(v string) *InsertK8sApplicationResponseBodyApplicationInfo {
	s.EdasId = &v
	return s
}

func (s *InsertK8sApplicationResponseBodyApplicationInfo) SetOwner(v string) *InsertK8sApplicationResponseBodyApplicationInfo {
	s.Owner = &v
	return s
}

func (s *InsertK8sApplicationResponseBodyApplicationInfo) SetDockerize(v bool) *InsertK8sApplicationResponseBodyApplicationInfo {
	s.Dockerize = &v
	return s
}

func (s *InsertK8sApplicationResponseBodyApplicationInfo) SetClusterType(v int) *InsertK8sApplicationResponseBodyApplicationInfo {
	s.ClusterType = &v
	return s
}

func (s *InsertK8sApplicationResponseBodyApplicationInfo) SetRegionId(v string) *InsertK8sApplicationResponseBodyApplicationInfo {
	s.RegionId = &v
	return s
}

func (s *InsertK8sApplicationResponseBodyApplicationInfo) SetChangeOrderId(v string) *InsertK8sApplicationResponseBodyApplicationInfo {
	s.ChangeOrderId = &v
	return s
}

type InsertK8sApplicationResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertK8sApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertK8sApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertK8sApplicationResponse) GoString() string {
	return s.String()
}

func (s *InsertK8sApplicationResponse) SetHeaders(v map[string]*string) *InsertK8sApplicationResponse {
	s.Headers = v
	return s
}

func (s *InsertK8sApplicationResponse) SetBody(v *InsertK8sApplicationResponseBody) *InsertK8sApplicationResponse {
	s.Body = v
	return s
}

type ImportK8sClusterQuery struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	Mode        *int    `json:"Mode,omitempty" xml:"Mode,omitempty"`
	EnableAsm   *bool   `json:"EnableAsm,omitempty" xml:"EnableAsm,omitempty"`
}

func (s ImportK8sClusterQuery) String() string {
	return tea.Prettify(s)
}

func (s ImportK8sClusterQuery) GoString() string {
	return s.String()
}

func (s *ImportK8sClusterQuery) SetClusterId(v string) *ImportK8sClusterQuery {
	s.ClusterId = &v
	return s
}

func (s *ImportK8sClusterQuery) SetNamespaceId(v string) *ImportK8sClusterQuery {
	s.NamespaceId = &v
	return s
}

func (s *ImportK8sClusterQuery) SetMode(v int) *ImportK8sClusterQuery {
	s.Mode = &v
	return s
}

func (s *ImportK8sClusterQuery) SetEnableAsm(v bool) *ImportK8sClusterQuery {
	s.EnableAsm = &v
	return s
}

type ImportK8sClusterRequest struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ImportK8sClusterQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ImportK8sClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportK8sClusterRequest) GoString() string {
	return s.String()
}

func (s *ImportK8sClusterRequest) SetHeaders(v map[string]*string) *ImportK8sClusterRequest {
	s.Headers = v
	return s
}

func (s *ImportK8sClusterRequest) SetQuery(v *ImportK8sClusterQuery) *ImportK8sClusterRequest {
	s.Query = v
	return s
}

type ImportK8sClusterResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ImportK8sClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportK8sClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ImportK8sClusterResponseBody) SetCode(v int) *ImportK8sClusterResponseBody {
	s.Code = &v
	return s
}

func (s *ImportK8sClusterResponseBody) SetMessage(v string) *ImportK8sClusterResponseBody {
	s.Message = &v
	return s
}

func (s *ImportK8sClusterResponseBody) SetData(v string) *ImportK8sClusterResponseBody {
	s.Data = &v
	return s
}

func (s *ImportK8sClusterResponseBody) SetRequestId(v string) *ImportK8sClusterResponseBody {
	s.RequestId = &v
	return s
}

type ImportK8sClusterResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportK8sClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportK8sClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportK8sClusterResponse) GoString() string {
	return s.String()
}

func (s *ImportK8sClusterResponse) SetHeaders(v map[string]*string) *ImportK8sClusterResponse {
	s.Headers = v
	return s
}

func (s *ImportK8sClusterResponse) SetBody(v *ImportK8sClusterResponseBody) *ImportK8sClusterResponse {
	s.Body = v
	return s
}

type UpdateK8sApplicationConfigQuery struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	CpuLimit      *string `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty" require:"true"`
	MemoryLimit   *string `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty" require:"true"`
	McpuLimit     *string `json:"McpuLimit,omitempty" xml:"McpuLimit,omitempty"`
	CpuRequest    *string `json:"CpuRequest,omitempty" xml:"CpuRequest,omitempty"`
	McpuRequest   *string `json:"McpuRequest,omitempty" xml:"McpuRequest,omitempty"`
	MemoryRequest *string `json:"MemoryRequest,omitempty" xml:"MemoryRequest,omitempty"`
	Timeout       *int    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateK8sApplicationConfigQuery) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sApplicationConfigQuery) GoString() string {
	return s.String()
}

func (s *UpdateK8sApplicationConfigQuery) SetAppId(v string) *UpdateK8sApplicationConfigQuery {
	s.AppId = &v
	return s
}

func (s *UpdateK8sApplicationConfigQuery) SetClusterId(v string) *UpdateK8sApplicationConfigQuery {
	s.ClusterId = &v
	return s
}

func (s *UpdateK8sApplicationConfigQuery) SetCpuLimit(v string) *UpdateK8sApplicationConfigQuery {
	s.CpuLimit = &v
	return s
}

func (s *UpdateK8sApplicationConfigQuery) SetMemoryLimit(v string) *UpdateK8sApplicationConfigQuery {
	s.MemoryLimit = &v
	return s
}

func (s *UpdateK8sApplicationConfigQuery) SetMcpuLimit(v string) *UpdateK8sApplicationConfigQuery {
	s.McpuLimit = &v
	return s
}

func (s *UpdateK8sApplicationConfigQuery) SetCpuRequest(v string) *UpdateK8sApplicationConfigQuery {
	s.CpuRequest = &v
	return s
}

func (s *UpdateK8sApplicationConfigQuery) SetMcpuRequest(v string) *UpdateK8sApplicationConfigQuery {
	s.McpuRequest = &v
	return s
}

func (s *UpdateK8sApplicationConfigQuery) SetMemoryRequest(v string) *UpdateK8sApplicationConfigQuery {
	s.MemoryRequest = &v
	return s
}

func (s *UpdateK8sApplicationConfigQuery) SetTimeout(v int) *UpdateK8sApplicationConfigQuery {
	s.Timeout = &v
	return s
}

type UpdateK8sApplicationConfigRequest struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *UpdateK8sApplicationConfigQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s UpdateK8sApplicationConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sApplicationConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateK8sApplicationConfigRequest) SetHeaders(v map[string]*string) *UpdateK8sApplicationConfigRequest {
	s.Headers = v
	return s
}

func (s *UpdateK8sApplicationConfigRequest) SetQuery(v *UpdateK8sApplicationConfigQuery) *UpdateK8sApplicationConfigRequest {
	s.Query = v
	return s
}

type UpdateK8sApplicationConfigResponseBody struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateK8sApplicationConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sApplicationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateK8sApplicationConfigResponseBody) SetChangeOrderId(v string) *UpdateK8sApplicationConfigResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *UpdateK8sApplicationConfigResponseBody) SetCode(v int) *UpdateK8sApplicationConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateK8sApplicationConfigResponseBody) SetMessage(v string) *UpdateK8sApplicationConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateK8sApplicationConfigResponseBody) SetRequestId(v string) *UpdateK8sApplicationConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateK8sApplicationConfigResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateK8sApplicationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateK8sApplicationConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sApplicationConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateK8sApplicationConfigResponse) SetHeaders(v map[string]*string) *UpdateK8sApplicationConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateK8sApplicationConfigResponse) SetBody(v *UpdateK8sApplicationConfigResponseBody) *UpdateK8sApplicationConfigResponse {
	s.Body = v
	return s
}

type DeployK8sApplicationQuery struct {
	PreStop                *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	Envs                   *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	ImageTag               *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	BatchWaitTime          *int    `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	Command                *string `json:"Command,omitempty" xml:"Command,omitempty"`
	AppId                  *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	PostStart              *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	Readiness              *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	Liveness               *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	Args                   *string `json:"Args,omitempty" xml:"Args,omitempty"`
	Replicas               *int    `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	Image                  *string `json:"Image,omitempty" xml:"Image,omitempty"`
	CpuLimit               *int    `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	MemoryLimit            *int    `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	CpuRequest             *int    `json:"CpuRequest,omitempty" xml:"CpuRequest,omitempty"`
	MemoryRequest          *int    `json:"MemoryRequest,omitempty" xml:"MemoryRequest,omitempty"`
	NasId                  *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	MountDescs             *string `json:"MountDescs,omitempty" xml:"MountDescs,omitempty"`
	StorageType            *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	LocalVolume            *string `json:"LocalVolume,omitempty" xml:"LocalVolume,omitempty"`
	PackageUrl             *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	PackageVersion         *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	JDK                    *string `json:"JDK,omitempty" xml:"JDK,omitempty"`
	WebContainer           *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
	EdasContainerVersion   *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	UriEncoding            *string `json:"UriEncoding,omitempty" xml:"UriEncoding,omitempty"`
	UseBodyEncoding        *bool   `json:"UseBodyEncoding,omitempty" xml:"UseBodyEncoding,omitempty"`
	UpdateStrategy         *string `json:"UpdateStrategy,omitempty" xml:"UpdateStrategy,omitempty"`
	McpuRequest            *int    `json:"McpuRequest,omitempty" xml:"McpuRequest,omitempty"`
	McpuLimit              *int    `json:"McpuLimit,omitempty" xml:"McpuLimit,omitempty"`
	VolumesStr             *string `json:"VolumesStr,omitempty" xml:"VolumesStr,omitempty"`
	PackageVersionId       *string `json:"PackageVersionId,omitempty" xml:"PackageVersionId,omitempty"`
	ChangeOrderDesc        *string `json:"ChangeOrderDesc,omitempty" xml:"ChangeOrderDesc,omitempty"`
	RuntimeClassName       *string `json:"RuntimeClassName,omitempty" xml:"RuntimeClassName,omitempty"`
	DeployAcrossZones      *string `json:"DeployAcrossZones,omitempty" xml:"DeployAcrossZones,omitempty"`
	BatchTimeout           *int    `json:"BatchTimeout,omitempty" xml:"BatchTimeout,omitempty"`
	EnableAhas             *bool   `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
	WebContainerConfig     *string `json:"WebContainerConfig,omitempty" xml:"WebContainerConfig,omitempty"`
	JavaStartUpConfig      *string `json:"JavaStartUpConfig,omitempty" xml:"JavaStartUpConfig,omitempty"`
	SlsConfigs             *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	DeployAcrossNodes      *string `json:"DeployAcrossNodes,omitempty" xml:"DeployAcrossNodes,omitempty"`
	TrafficControlStrategy *string `json:"TrafficControlStrategy,omitempty" xml:"TrafficControlStrategy,omitempty"`
}

func (s DeployK8sApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s DeployK8sApplicationQuery) GoString() string {
	return s.String()
}

func (s *DeployK8sApplicationQuery) SetPreStop(v string) *DeployK8sApplicationQuery {
	s.PreStop = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetEnvs(v string) *DeployK8sApplicationQuery {
	s.Envs = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetImageTag(v string) *DeployK8sApplicationQuery {
	s.ImageTag = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetBatchWaitTime(v int) *DeployK8sApplicationQuery {
	s.BatchWaitTime = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetCommand(v string) *DeployK8sApplicationQuery {
	s.Command = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetAppId(v string) *DeployK8sApplicationQuery {
	s.AppId = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetPostStart(v string) *DeployK8sApplicationQuery {
	s.PostStart = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetReadiness(v string) *DeployK8sApplicationQuery {
	s.Readiness = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetLiveness(v string) *DeployK8sApplicationQuery {
	s.Liveness = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetArgs(v string) *DeployK8sApplicationQuery {
	s.Args = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetReplicas(v int) *DeployK8sApplicationQuery {
	s.Replicas = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetImage(v string) *DeployK8sApplicationQuery {
	s.Image = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetCpuLimit(v int) *DeployK8sApplicationQuery {
	s.CpuLimit = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetMemoryLimit(v int) *DeployK8sApplicationQuery {
	s.MemoryLimit = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetCpuRequest(v int) *DeployK8sApplicationQuery {
	s.CpuRequest = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetMemoryRequest(v int) *DeployK8sApplicationQuery {
	s.MemoryRequest = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetNasId(v string) *DeployK8sApplicationQuery {
	s.NasId = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetMountDescs(v string) *DeployK8sApplicationQuery {
	s.MountDescs = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetStorageType(v string) *DeployK8sApplicationQuery {
	s.StorageType = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetLocalVolume(v string) *DeployK8sApplicationQuery {
	s.LocalVolume = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetPackageUrl(v string) *DeployK8sApplicationQuery {
	s.PackageUrl = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetPackageVersion(v string) *DeployK8sApplicationQuery {
	s.PackageVersion = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetJDK(v string) *DeployK8sApplicationQuery {
	s.JDK = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetWebContainer(v string) *DeployK8sApplicationQuery {
	s.WebContainer = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetEdasContainerVersion(v string) *DeployK8sApplicationQuery {
	s.EdasContainerVersion = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetUriEncoding(v string) *DeployK8sApplicationQuery {
	s.UriEncoding = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetUseBodyEncoding(v bool) *DeployK8sApplicationQuery {
	s.UseBodyEncoding = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetUpdateStrategy(v string) *DeployK8sApplicationQuery {
	s.UpdateStrategy = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetMcpuRequest(v int) *DeployK8sApplicationQuery {
	s.McpuRequest = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetMcpuLimit(v int) *DeployK8sApplicationQuery {
	s.McpuLimit = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetVolumesStr(v string) *DeployK8sApplicationQuery {
	s.VolumesStr = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetPackageVersionId(v string) *DeployK8sApplicationQuery {
	s.PackageVersionId = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetChangeOrderDesc(v string) *DeployK8sApplicationQuery {
	s.ChangeOrderDesc = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetRuntimeClassName(v string) *DeployK8sApplicationQuery {
	s.RuntimeClassName = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetDeployAcrossZones(v string) *DeployK8sApplicationQuery {
	s.DeployAcrossZones = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetBatchTimeout(v int) *DeployK8sApplicationQuery {
	s.BatchTimeout = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetEnableAhas(v bool) *DeployK8sApplicationQuery {
	s.EnableAhas = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetWebContainerConfig(v string) *DeployK8sApplicationQuery {
	s.WebContainerConfig = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetJavaStartUpConfig(v string) *DeployK8sApplicationQuery {
	s.JavaStartUpConfig = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetSlsConfigs(v string) *DeployK8sApplicationQuery {
	s.SlsConfigs = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetDeployAcrossNodes(v string) *DeployK8sApplicationQuery {
	s.DeployAcrossNodes = &v
	return s
}

func (s *DeployK8sApplicationQuery) SetTrafficControlStrategy(v string) *DeployK8sApplicationQuery {
	s.TrafficControlStrategy = &v
	return s
}

type DeployK8sApplicationRequest struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeployK8sApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DeployK8sApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeployK8sApplicationRequest) SetHeaders(v map[string]*string) *DeployK8sApplicationRequest {
	s.Headers = v
	return s
}

func (s *DeployK8sApplicationRequest) SetQuery(v *DeployK8sApplicationQuery) *DeployK8sApplicationRequest {
	s.Query = v
	return s
}

type DeployK8sApplicationResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s DeployK8sApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployK8sApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeployK8sApplicationResponseBody) SetRequestId(v string) *DeployK8sApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployK8sApplicationResponseBody) SetCode(v int) *DeployK8sApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeployK8sApplicationResponseBody) SetMessage(v string) *DeployK8sApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeployK8sApplicationResponseBody) SetChangeOrderId(v string) *DeployK8sApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

type DeployK8sApplicationResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeployK8sApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployK8sApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployK8sApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeployK8sApplicationResponse) SetHeaders(v map[string]*string) *DeployK8sApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeployK8sApplicationResponse) SetBody(v *DeployK8sApplicationResponseBody) *DeployK8sApplicationResponse {
	s.Body = v
	return s
}

type ScaleK8sApplicationQuery struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Replicas *int    `json:"Replicas,omitempty" xml:"Replicas,omitempty" require:"true"`
	Timeout  *int    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ScaleK8sApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s ScaleK8sApplicationQuery) GoString() string {
	return s.String()
}

func (s *ScaleK8sApplicationQuery) SetAppId(v string) *ScaleK8sApplicationQuery {
	s.AppId = &v
	return s
}

func (s *ScaleK8sApplicationQuery) SetReplicas(v int) *ScaleK8sApplicationQuery {
	s.Replicas = &v
	return s
}

func (s *ScaleK8sApplicationQuery) SetTimeout(v int) *ScaleK8sApplicationQuery {
	s.Timeout = &v
	return s
}

type ScaleK8sApplicationRequest struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ScaleK8sApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ScaleK8sApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *ScaleK8sApplicationRequest) SetHeaders(v map[string]*string) *ScaleK8sApplicationRequest {
	s.Headers = v
	return s
}

func (s *ScaleK8sApplicationRequest) SetQuery(v *ScaleK8sApplicationQuery) *ScaleK8sApplicationRequest {
	s.Query = v
	return s
}

type ScaleK8sApplicationResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s ScaleK8sApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleK8sApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleK8sApplicationResponseBody) SetRequestId(v string) *ScaleK8sApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleK8sApplicationResponseBody) SetCode(v int) *ScaleK8sApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ScaleK8sApplicationResponseBody) SetMessage(v string) *ScaleK8sApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ScaleK8sApplicationResponseBody) SetChangeOrderId(v string) *ScaleK8sApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

type ScaleK8sApplicationResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ScaleK8sApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScaleK8sApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleK8sApplicationResponse) GoString() string {
	return s.String()
}

func (s *ScaleK8sApplicationResponse) SetHeaders(v map[string]*string) *ScaleK8sApplicationResponse {
	s.Headers = v
	return s
}

func (s *ScaleK8sApplicationResponse) SetBody(v *ScaleK8sApplicationResponseBody) *ScaleK8sApplicationResponse {
	s.Body = v
	return s
}

type DeleteK8sApplicationQuery struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s DeleteK8sApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s DeleteK8sApplicationQuery) GoString() string {
	return s.String()
}

func (s *DeleteK8sApplicationQuery) SetAppId(v string) *DeleteK8sApplicationQuery {
	s.AppId = &v
	return s
}

type DeleteK8sApplicationRequest struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeleteK8sApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DeleteK8sApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteK8sApplicationRequest) SetHeaders(v map[string]*string) *DeleteK8sApplicationRequest {
	s.Headers = v
	return s
}

func (s *DeleteK8sApplicationRequest) SetQuery(v *DeleteK8sApplicationQuery) *DeleteK8sApplicationRequest {
	s.Query = v
	return s
}

type DeleteK8sApplicationResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s DeleteK8sApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteK8sApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteK8sApplicationResponseBody) SetRequestId(v string) *DeleteK8sApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteK8sApplicationResponseBody) SetCode(v int) *DeleteK8sApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteK8sApplicationResponseBody) SetMessage(v string) *DeleteK8sApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteK8sApplicationResponseBody) SetChangeOrderId(v string) *DeleteK8sApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

type DeleteK8sApplicationResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteK8sApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteK8sApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteK8sApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteK8sApplicationResponse) SetHeaders(v map[string]*string) *DeleteK8sApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteK8sApplicationResponse) SetBody(v *DeleteK8sApplicationResponseBody) *DeleteK8sApplicationResponse {
	s.Body = v
	return s
}

type UnbindK8sSlbQuery struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
}

func (s UnbindK8sSlbQuery) String() string {
	return tea.Prettify(s)
}

func (s UnbindK8sSlbQuery) GoString() string {
	return s.String()
}

func (s *UnbindK8sSlbQuery) SetAppId(v string) *UnbindK8sSlbQuery {
	s.AppId = &v
	return s
}

func (s *UnbindK8sSlbQuery) SetClusterId(v string) *UnbindK8sSlbQuery {
	s.ClusterId = &v
	return s
}

func (s *UnbindK8sSlbQuery) SetType(v string) *UnbindK8sSlbQuery {
	s.Type = &v
	return s
}

type UnbindK8sSlbRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *UnbindK8sSlbQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s UnbindK8sSlbRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindK8sSlbRequest) GoString() string {
	return s.String()
}

func (s *UnbindK8sSlbRequest) SetHeaders(v map[string]*string) *UnbindK8sSlbRequest {
	s.Headers = v
	return s
}

func (s *UnbindK8sSlbRequest) SetQuery(v *UnbindK8sSlbQuery) *UnbindK8sSlbRequest {
	s.Query = v
	return s
}

type UnbindK8sSlbResponseBody struct {
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UnbindK8sSlbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindK8sSlbResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindK8sSlbResponseBody) SetCode(v int) *UnbindK8sSlbResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindK8sSlbResponseBody) SetMessage(v string) *UnbindK8sSlbResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindK8sSlbResponseBody) SetChangeOrderId(v string) *UnbindK8sSlbResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *UnbindK8sSlbResponseBody) SetRequestId(v string) *UnbindK8sSlbResponseBody {
	s.RequestId = &v
	return s
}

type UnbindK8sSlbResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindK8sSlbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindK8sSlbResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindK8sSlbResponse) GoString() string {
	return s.String()
}

func (s *UnbindK8sSlbResponse) SetHeaders(v map[string]*string) *UnbindK8sSlbResponse {
	s.Headers = v
	return s
}

func (s *UnbindK8sSlbResponse) SetBody(v *UnbindK8sSlbResponseBody) *UnbindK8sSlbResponse {
	s.Body = v
	return s
}

type BindK8sSlbQuery struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	SlbId            *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	SlbProtocol      *string `json:"SlbProtocol,omitempty" xml:"SlbProtocol,omitempty"`
	TargetPort       *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ServicePortInfos *string `json:"ServicePortInfos,omitempty" xml:"ServicePortInfos,omitempty"`
	Specification    *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	Scheduler        *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
}

func (s BindK8sSlbQuery) String() string {
	return tea.Prettify(s)
}

func (s BindK8sSlbQuery) GoString() string {
	return s.String()
}

func (s *BindK8sSlbQuery) SetAppId(v string) *BindK8sSlbQuery {
	s.AppId = &v
	return s
}

func (s *BindK8sSlbQuery) SetClusterId(v string) *BindK8sSlbQuery {
	s.ClusterId = &v
	return s
}

func (s *BindK8sSlbQuery) SetType(v string) *BindK8sSlbQuery {
	s.Type = &v
	return s
}

func (s *BindK8sSlbQuery) SetSlbId(v string) *BindK8sSlbQuery {
	s.SlbId = &v
	return s
}

func (s *BindK8sSlbQuery) SetSlbProtocol(v string) *BindK8sSlbQuery {
	s.SlbProtocol = &v
	return s
}

func (s *BindK8sSlbQuery) SetTargetPort(v string) *BindK8sSlbQuery {
	s.TargetPort = &v
	return s
}

func (s *BindK8sSlbQuery) SetPort(v string) *BindK8sSlbQuery {
	s.Port = &v
	return s
}

func (s *BindK8sSlbQuery) SetServicePortInfos(v string) *BindK8sSlbQuery {
	s.ServicePortInfos = &v
	return s
}

func (s *BindK8sSlbQuery) SetSpecification(v string) *BindK8sSlbQuery {
	s.Specification = &v
	return s
}

func (s *BindK8sSlbQuery) SetScheduler(v string) *BindK8sSlbQuery {
	s.Scheduler = &v
	return s
}

type BindK8sSlbRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *BindK8sSlbQuery   `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s BindK8sSlbRequest) String() string {
	return tea.Prettify(s)
}

func (s BindK8sSlbRequest) GoString() string {
	return s.String()
}

func (s *BindK8sSlbRequest) SetHeaders(v map[string]*string) *BindK8sSlbRequest {
	s.Headers = v
	return s
}

func (s *BindK8sSlbRequest) SetQuery(v *BindK8sSlbQuery) *BindK8sSlbRequest {
	s.Query = v
	return s
}

type BindK8sSlbResponseBody struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s BindK8sSlbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindK8sSlbResponseBody) GoString() string {
	return s.String()
}

func (s *BindK8sSlbResponseBody) SetChangeOrderId(v string) *BindK8sSlbResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *BindK8sSlbResponseBody) SetCode(v int) *BindK8sSlbResponseBody {
	s.Code = &v
	return s
}

func (s *BindK8sSlbResponseBody) SetMessage(v string) *BindK8sSlbResponseBody {
	s.Message = &v
	return s
}

func (s *BindK8sSlbResponseBody) SetRequestId(v string) *BindK8sSlbResponseBody {
	s.RequestId = &v
	return s
}

type BindK8sSlbResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindK8sSlbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindK8sSlbResponse) String() string {
	return tea.Prettify(s)
}

func (s BindK8sSlbResponse) GoString() string {
	return s.String()
}

func (s *BindK8sSlbResponse) SetHeaders(v map[string]*string) *BindK8sSlbResponse {
	s.Headers = v
	return s
}

func (s *BindK8sSlbResponse) SetBody(v *BindK8sSlbResponseBody) *BindK8sSlbResponse {
	s.Body = v
	return s
}

type UpdateK8sSlbQuery struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	SlbProtocol      *string `json:"SlbProtocol,omitempty" xml:"SlbProtocol,omitempty"`
	TargetPort       *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ServicePortInfos *string `json:"ServicePortInfos,omitempty" xml:"ServicePortInfos,omitempty"`
	Specification    *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	Scheduler        *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
}

func (s UpdateK8sSlbQuery) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sSlbQuery) GoString() string {
	return s.String()
}

func (s *UpdateK8sSlbQuery) SetAppId(v string) *UpdateK8sSlbQuery {
	s.AppId = &v
	return s
}

func (s *UpdateK8sSlbQuery) SetClusterId(v string) *UpdateK8sSlbQuery {
	s.ClusterId = &v
	return s
}

func (s *UpdateK8sSlbQuery) SetType(v string) *UpdateK8sSlbQuery {
	s.Type = &v
	return s
}

func (s *UpdateK8sSlbQuery) SetSlbProtocol(v string) *UpdateK8sSlbQuery {
	s.SlbProtocol = &v
	return s
}

func (s *UpdateK8sSlbQuery) SetTargetPort(v string) *UpdateK8sSlbQuery {
	s.TargetPort = &v
	return s
}

func (s *UpdateK8sSlbQuery) SetPort(v string) *UpdateK8sSlbQuery {
	s.Port = &v
	return s
}

func (s *UpdateK8sSlbQuery) SetServicePortInfos(v string) *UpdateK8sSlbQuery {
	s.ServicePortInfos = &v
	return s
}

func (s *UpdateK8sSlbQuery) SetSpecification(v string) *UpdateK8sSlbQuery {
	s.Specification = &v
	return s
}

func (s *UpdateK8sSlbQuery) SetScheduler(v string) *UpdateK8sSlbQuery {
	s.Scheduler = &v
	return s
}

type UpdateK8sSlbRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *UpdateK8sSlbQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s UpdateK8sSlbRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sSlbRequest) GoString() string {
	return s.String()
}

func (s *UpdateK8sSlbRequest) SetHeaders(v map[string]*string) *UpdateK8sSlbRequest {
	s.Headers = v
	return s
}

func (s *UpdateK8sSlbRequest) SetQuery(v *UpdateK8sSlbQuery) *UpdateK8sSlbRequest {
	s.Query = v
	return s
}

type UpdateK8sSlbResponseBody struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateK8sSlbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sSlbResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateK8sSlbResponseBody) SetChangeOrderId(v string) *UpdateK8sSlbResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *UpdateK8sSlbResponseBody) SetCode(v int) *UpdateK8sSlbResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateK8sSlbResponseBody) SetMessage(v string) *UpdateK8sSlbResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateK8sSlbResponseBody) SetRequestId(v string) *UpdateK8sSlbResponseBody {
	s.RequestId = &v
	return s
}

type UpdateK8sSlbResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateK8sSlbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateK8sSlbResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateK8sSlbResponse) GoString() string {
	return s.String()
}

func (s *UpdateK8sSlbResponse) SetHeaders(v map[string]*string) *UpdateK8sSlbResponse {
	s.Headers = v
	return s
}

func (s *UpdateK8sSlbResponse) SetBody(v *UpdateK8sSlbResponseBody) *UpdateK8sSlbResponse {
	s.Body = v
	return s
}

type GetSecureTokenQuery struct {
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty" require:"true"`
}

func (s GetSecureTokenQuery) String() string {
	return tea.Prettify(s)
}

func (s GetSecureTokenQuery) GoString() string {
	return s.String()
}

func (s *GetSecureTokenQuery) SetNamespaceId(v string) *GetSecureTokenQuery {
	s.NamespaceId = &v
	return s
}

type GetSecureTokenRequest struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *GetSecureTokenQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s GetSecureTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecureTokenRequest) GoString() string {
	return s.String()
}

func (s *GetSecureTokenRequest) SetHeaders(v map[string]*string) *GetSecureTokenRequest {
	s.Headers = v
	return s
}

func (s *GetSecureTokenRequest) SetQuery(v *GetSecureTokenQuery) *GetSecureTokenRequest {
	s.Query = v
	return s
}

type GetSecureTokenResponseBody struct {
	Code        *int                                   `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message     *string                                `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SecureToken *GetSecureTokenResponseBodySecureToken `json:"SecureToken,omitempty" xml:"SecureToken,omitempty" require:"true" type:"Struct"`
}

func (s GetSecureTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecureTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecureTokenResponseBody) SetCode(v int) *GetSecureTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecureTokenResponseBody) SetMessage(v string) *GetSecureTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecureTokenResponseBody) SetRequestId(v string) *GetSecureTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecureTokenResponseBody) SetSecureToken(v *GetSecureTokenResponseBodySecureToken) *GetSecureTokenResponseBody {
	s.SecureToken = v
	return s
}

type GetSecureTokenResponseBodySecureToken struct {
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	EdasId            *string `json:"EdasId,omitempty" xml:"EdasId,omitempty" require:"true"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	RegionName        *string `json:"RegionName,omitempty" xml:"RegionName,omitempty" require:"true"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	BelongRegion      *string `json:"BelongRegion,omitempty" xml:"BelongRegion,omitempty" require:"true"`
	AccessKey         *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" require:"true"`
	SecretKey         *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty" require:"true"`
	TenantId          *string `json:"TenantId,omitempty" xml:"TenantId,omitempty" require:"true"`
	AddressServerHost *string `json:"AddressServerHost,omitempty" xml:"AddressServerHost,omitempty" require:"true"`
}

func (s GetSecureTokenResponseBodySecureToken) String() string {
	return tea.Prettify(s)
}

func (s GetSecureTokenResponseBodySecureToken) GoString() string {
	return s.String()
}

func (s *GetSecureTokenResponseBodySecureToken) SetId(v int64) *GetSecureTokenResponseBodySecureToken {
	s.Id = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetUserId(v string) *GetSecureTokenResponseBodySecureToken {
	s.UserId = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetEdasId(v string) *GetSecureTokenResponseBodySecureToken {
	s.EdasId = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetRegionId(v string) *GetSecureTokenResponseBodySecureToken {
	s.RegionId = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetRegionName(v string) *GetSecureTokenResponseBodySecureToken {
	s.RegionName = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetDescription(v string) *GetSecureTokenResponseBodySecureToken {
	s.Description = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetBelongRegion(v string) *GetSecureTokenResponseBodySecureToken {
	s.BelongRegion = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetAccessKey(v string) *GetSecureTokenResponseBodySecureToken {
	s.AccessKey = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetSecretKey(v string) *GetSecureTokenResponseBodySecureToken {
	s.SecretKey = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetTenantId(v string) *GetSecureTokenResponseBodySecureToken {
	s.TenantId = &v
	return s
}

func (s *GetSecureTokenResponseBodySecureToken) SetAddressServerHost(v string) *GetSecureTokenResponseBodySecureToken {
	s.AddressServerHost = &v
	return s
}

type GetSecureTokenResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSecureTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSecureTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecureTokenResponse) GoString() string {
	return s.String()
}

func (s *GetSecureTokenResponse) SetHeaders(v map[string]*string) *GetSecureTokenResponse {
	s.Headers = v
	return s
}

func (s *GetSecureTokenResponse) SetBody(v *GetSecureTokenResponseBody) *GetSecureTokenResponse {
	s.Body = v
	return s
}

type TransformClusterMemberQuery struct {
	InstanceIds     *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" require:"true"`
	TargetClusterId *string `json:"TargetClusterId,omitempty" xml:"TargetClusterId,omitempty" require:"true"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty" require:"true"`
}

func (s TransformClusterMemberQuery) String() string {
	return tea.Prettify(s)
}

func (s TransformClusterMemberQuery) GoString() string {
	return s.String()
}

func (s *TransformClusterMemberQuery) SetInstanceIds(v string) *TransformClusterMemberQuery {
	s.InstanceIds = &v
	return s
}

func (s *TransformClusterMemberQuery) SetTargetClusterId(v string) *TransformClusterMemberQuery {
	s.TargetClusterId = &v
	return s
}

func (s *TransformClusterMemberQuery) SetPassword(v string) *TransformClusterMemberQuery {
	s.Password = &v
	return s
}

type TransformClusterMemberRequest struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *TransformClusterMemberQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s TransformClusterMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s TransformClusterMemberRequest) GoString() string {
	return s.String()
}

func (s *TransformClusterMemberRequest) SetHeaders(v map[string]*string) *TransformClusterMemberRequest {
	s.Headers = v
	return s
}

func (s *TransformClusterMemberRequest) SetQuery(v *TransformClusterMemberQuery) *TransformClusterMemberRequest {
	s.Query = v
	return s
}

type TransformClusterMemberResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s TransformClusterMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransformClusterMemberResponseBody) GoString() string {
	return s.String()
}

func (s *TransformClusterMemberResponseBody) SetCode(v int) *TransformClusterMemberResponseBody {
	s.Code = &v
	return s
}

func (s *TransformClusterMemberResponseBody) SetMessage(v string) *TransformClusterMemberResponseBody {
	s.Message = &v
	return s
}

func (s *TransformClusterMemberResponseBody) SetData(v string) *TransformClusterMemberResponseBody {
	s.Data = &v
	return s
}

func (s *TransformClusterMemberResponseBody) SetRequestId(v string) *TransformClusterMemberResponseBody {
	s.RequestId = &v
	return s
}

type TransformClusterMemberResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TransformClusterMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransformClusterMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s TransformClusterMemberResponse) GoString() string {
	return s.String()
}

func (s *TransformClusterMemberResponse) SetHeaders(v map[string]*string) *TransformClusterMemberResponse {
	s.Headers = v
	return s
}

func (s *TransformClusterMemberResponse) SetBody(v *TransformClusterMemberResponseBody) *TransformClusterMemberResponse {
	s.Body = v
	return s
}

type ListConvertableEcuQuery struct {
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty" require:"true"`
}

func (s ListConvertableEcuQuery) String() string {
	return tea.Prettify(s)
}

func (s ListConvertableEcuQuery) GoString() string {
	return s.String()
}

func (s *ListConvertableEcuQuery) SetClusterId(v string) *ListConvertableEcuQuery {
	s.ClusterId = &v
	return s
}

type ListConvertableEcuRequest struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListConvertableEcuQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ListConvertableEcuRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConvertableEcuRequest) GoString() string {
	return s.String()
}

func (s *ListConvertableEcuRequest) SetHeaders(v map[string]*string) *ListConvertableEcuRequest {
	s.Headers = v
	return s
}

func (s *ListConvertableEcuRequest) SetQuery(v *ListConvertableEcuQuery) *ListConvertableEcuRequest {
	s.Query = v
	return s
}

type ListConvertableEcuResponseBody struct {
	Code         *int                                        `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message      *string                                     `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	InstanceList *ListConvertableEcuResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" require:"true" type:"Struct"`
}

func (s ListConvertableEcuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConvertableEcuResponseBody) GoString() string {
	return s.String()
}

func (s *ListConvertableEcuResponseBody) SetCode(v int) *ListConvertableEcuResponseBody {
	s.Code = &v
	return s
}

func (s *ListConvertableEcuResponseBody) SetMessage(v string) *ListConvertableEcuResponseBody {
	s.Message = &v
	return s
}

func (s *ListConvertableEcuResponseBody) SetRequestId(v string) *ListConvertableEcuResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConvertableEcuResponseBody) SetInstanceList(v *ListConvertableEcuResponseBodyInstanceList) *ListConvertableEcuResponseBody {
	s.InstanceList = v
	return s
}

type ListConvertableEcuResponseBodyInstanceList struct {
	Instance []*ListConvertableEcuResponseBodyInstanceListInstance `json:"Instance,omitempty" xml:"Instance,omitempty" require:"true" type:"Repeated"`
}

func (s ListConvertableEcuResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s ListConvertableEcuResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *ListConvertableEcuResponseBodyInstanceList) SetInstance(v []*ListConvertableEcuResponseBodyInstanceListInstance) *ListConvertableEcuResponseBodyInstanceList {
	s.Instance = v
	return s
}

type ListConvertableEcuResponseBodyInstanceListInstance struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty" require:"true"`
	EcuId        *string `json:"EcuId,omitempty" xml:"EcuId,omitempty" require:"true"`
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	VpcName      *string `json:"VpcName,omitempty" xml:"VpcName,omitempty" require:"true"`
	Expired      *bool   `json:"Expired,omitempty" xml:"Expired,omitempty" require:"true"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Cpu          *int    `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Mem          *int    `json:"Mem,omitempty" xml:"Mem,omitempty" require:"true"`
	PublicIp     *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty" require:"true"`
	InnerIp      *string `json:"InnerIp,omitempty" xml:"InnerIp,omitempty" require:"true"`
	PrivateIp    *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty" require:"true"`
	Eip          *string `json:"Eip,omitempty" xml:"Eip,omitempty" require:"true"`
}

func (s ListConvertableEcuResponseBodyInstanceListInstance) String() string {
	return tea.Prettify(s)
}

func (s ListConvertableEcuResponseBodyInstanceListInstance) GoString() string {
	return s.String()
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetInstanceId(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.InstanceId = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetInstanceName(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.InstanceName = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetEcuId(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.EcuId = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetVpcId(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.VpcId = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetVpcName(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.VpcName = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetExpired(v bool) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.Expired = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetStatus(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.Status = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetRegionId(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.RegionId = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetCpu(v int) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.Cpu = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetMem(v int) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.Mem = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetPublicIp(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.PublicIp = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetInnerIp(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.InnerIp = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetPrivateIp(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.PrivateIp = &v
	return s
}

func (s *ListConvertableEcuResponseBodyInstanceListInstance) SetEip(v string) *ListConvertableEcuResponseBodyInstanceListInstance {
	s.Eip = &v
	return s
}

type ListConvertableEcuResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConvertableEcuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConvertableEcuResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConvertableEcuResponse) GoString() string {
	return s.String()
}

func (s *ListConvertableEcuResponse) SetHeaders(v map[string]*string) *ListConvertableEcuResponse {
	s.Headers = v
	return s
}

func (s *ListConvertableEcuResponse) SetBody(v *ListConvertableEcuResponseBody) *ListConvertableEcuResponse {
	s.Body = v
	return s
}

type InsertClusterMemberQuery struct {
	ClusterId   *string `json:"clusterId,omitempty" xml:"clusterId,omitempty" require:"true"`
	InstanceIds *string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty" require:"true"`
	Password    *string `json:"password,omitempty" xml:"password,omitempty" require:"true"`
}

func (s InsertClusterMemberQuery) String() string {
	return tea.Prettify(s)
}

func (s InsertClusterMemberQuery) GoString() string {
	return s.String()
}

func (s *InsertClusterMemberQuery) SetClusterId(v string) *InsertClusterMemberQuery {
	s.ClusterId = &v
	return s
}

func (s *InsertClusterMemberQuery) SetInstanceIds(v string) *InsertClusterMemberQuery {
	s.InstanceIds = &v
	return s
}

func (s *InsertClusterMemberQuery) SetPassword(v string) *InsertClusterMemberQuery {
	s.Password = &v
	return s
}

type InsertClusterMemberRequest struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *InsertClusterMemberQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s InsertClusterMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertClusterMemberRequest) GoString() string {
	return s.String()
}

func (s *InsertClusterMemberRequest) SetHeaders(v map[string]*string) *InsertClusterMemberRequest {
	s.Headers = v
	return s
}

func (s *InsertClusterMemberRequest) SetQuery(v *InsertClusterMemberQuery) *InsertClusterMemberRequest {
	s.Query = v
	return s
}

type InsertClusterMemberResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s InsertClusterMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertClusterMemberResponseBody) GoString() string {
	return s.String()
}

func (s *InsertClusterMemberResponseBody) SetCode(v int) *InsertClusterMemberResponseBody {
	s.Code = &v
	return s
}

func (s *InsertClusterMemberResponseBody) SetMessage(v string) *InsertClusterMemberResponseBody {
	s.Message = &v
	return s
}

func (s *InsertClusterMemberResponseBody) SetData(v string) *InsertClusterMemberResponseBody {
	s.Data = &v
	return s
}

func (s *InsertClusterMemberResponseBody) SetRequestId(v string) *InsertClusterMemberResponseBody {
	s.RequestId = &v
	return s
}

type InsertClusterMemberResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertClusterMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertClusterMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertClusterMemberResponse) GoString() string {
	return s.String()
}

func (s *InsertClusterMemberResponse) SetHeaders(v map[string]*string) *InsertClusterMemberResponse {
	s.Headers = v
	return s
}

func (s *InsertClusterMemberResponse) SetBody(v *InsertClusterMemberResponseBody) *InsertClusterMemberResponse {
	s.Body = v
	return s
}

type UpdateRoleQuery struct {
	RoleId     *int    `json:"RoleId,omitempty" xml:"RoleId,omitempty" require:"true"`
	ActionData *string `json:"ActionData,omitempty" xml:"ActionData,omitempty" require:"true"`
}

func (s UpdateRoleQuery) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleQuery) GoString() string {
	return s.String()
}

func (s *UpdateRoleQuery) SetRoleId(v int) *UpdateRoleQuery {
	s.RoleId = &v
	return s
}

func (s *UpdateRoleQuery) SetActionData(v string) *UpdateRoleQuery {
	s.ActionData = &v
	return s
}

type UpdateRoleRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *UpdateRoleQuery   `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s UpdateRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoleRequest) SetHeaders(v map[string]*string) *UpdateRoleRequest {
	s.Headers = v
	return s
}

func (s *UpdateRoleRequest) SetQuery(v *UpdateRoleQuery) *UpdateRoleRequest {
	s.Query = v
	return s
}

type UpdateRoleResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponseBody) SetCode(v int) *UpdateRoleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRoleResponseBody) SetMessage(v string) *UpdateRoleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRoleResponseBody) SetRequestId(v string) *UpdateRoleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRoleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponse) SetHeaders(v map[string]*string) *UpdateRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoleResponse) SetBody(v *UpdateRoleResponseBody) *UpdateRoleResponse {
	s.Body = v
	return s
}

type UpdateJvmConfigurationQuery struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Options     *string `json:"Options,omitempty" xml:"Options,omitempty"`
	MinHeapSize *int    `json:"MinHeapSize,omitempty" xml:"MinHeapSize,omitempty"`
	MaxPermSize *int    `json:"MaxPermSize,omitempty" xml:"MaxPermSize,omitempty"`
	MaxHeapSize *int    `json:"MaxHeapSize,omitempty" xml:"MaxHeapSize,omitempty"`
}

func (s UpdateJvmConfigurationQuery) String() string {
	return tea.Prettify(s)
}

func (s UpdateJvmConfigurationQuery) GoString() string {
	return s.String()
}

func (s *UpdateJvmConfigurationQuery) SetAppId(v string) *UpdateJvmConfigurationQuery {
	s.AppId = &v
	return s
}

func (s *UpdateJvmConfigurationQuery) SetGroupId(v string) *UpdateJvmConfigurationQuery {
	s.GroupId = &v
	return s
}

func (s *UpdateJvmConfigurationQuery) SetOptions(v string) *UpdateJvmConfigurationQuery {
	s.Options = &v
	return s
}

func (s *UpdateJvmConfigurationQuery) SetMinHeapSize(v int) *UpdateJvmConfigurationQuery {
	s.MinHeapSize = &v
	return s
}

func (s *UpdateJvmConfigurationQuery) SetMaxPermSize(v int) *UpdateJvmConfigurationQuery {
	s.MaxPermSize = &v
	return s
}

func (s *UpdateJvmConfigurationQuery) SetMaxHeapSize(v int) *UpdateJvmConfigurationQuery {
	s.MaxHeapSize = &v
	return s
}

type UpdateJvmConfigurationRequest struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *UpdateJvmConfigurationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s UpdateJvmConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateJvmConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateJvmConfigurationRequest) SetHeaders(v map[string]*string) *UpdateJvmConfigurationRequest {
	s.Headers = v
	return s
}

func (s *UpdateJvmConfigurationRequest) SetQuery(v *UpdateJvmConfigurationQuery) *UpdateJvmConfigurationRequest {
	s.Query = v
	return s
}

type UpdateJvmConfigurationResponseBody struct {
	Code             *int                                                `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message          *string                                             `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	JvmConfiguration *UpdateJvmConfigurationResponseBodyJvmConfiguration `json:"JvmConfiguration,omitempty" xml:"JvmConfiguration,omitempty" require:"true" type:"Struct"`
}

func (s UpdateJvmConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateJvmConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJvmConfigurationResponseBody) SetCode(v int) *UpdateJvmConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateJvmConfigurationResponseBody) SetMessage(v string) *UpdateJvmConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateJvmConfigurationResponseBody) SetRequestId(v string) *UpdateJvmConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateJvmConfigurationResponseBody) SetJvmConfiguration(v *UpdateJvmConfigurationResponseBodyJvmConfiguration) *UpdateJvmConfigurationResponseBody {
	s.JvmConfiguration = v
	return s
}

type UpdateJvmConfigurationResponseBodyJvmConfiguration struct {
	MaxHeapSize *int    `json:"MaxHeapSize,omitempty" xml:"MaxHeapSize,omitempty" require:"true"`
	MaxPermSize *int    `json:"MaxPermSize,omitempty" xml:"MaxPermSize,omitempty" require:"true"`
	MinHeapSize *int    `json:"MinHeapSize,omitempty" xml:"MinHeapSize,omitempty" require:"true"`
	Options     *string `json:"Options,omitempty" xml:"Options,omitempty" require:"true"`
}

func (s UpdateJvmConfigurationResponseBodyJvmConfiguration) String() string {
	return tea.Prettify(s)
}

func (s UpdateJvmConfigurationResponseBodyJvmConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateJvmConfigurationResponseBodyJvmConfiguration) SetMaxHeapSize(v int) *UpdateJvmConfigurationResponseBodyJvmConfiguration {
	s.MaxHeapSize = &v
	return s
}

func (s *UpdateJvmConfigurationResponseBodyJvmConfiguration) SetMaxPermSize(v int) *UpdateJvmConfigurationResponseBodyJvmConfiguration {
	s.MaxPermSize = &v
	return s
}

func (s *UpdateJvmConfigurationResponseBodyJvmConfiguration) SetMinHeapSize(v int) *UpdateJvmConfigurationResponseBodyJvmConfiguration {
	s.MinHeapSize = &v
	return s
}

func (s *UpdateJvmConfigurationResponseBodyJvmConfiguration) SetOptions(v string) *UpdateJvmConfigurationResponseBodyJvmConfiguration {
	s.Options = &v
	return s
}

type UpdateJvmConfigurationResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateJvmConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateJvmConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateJvmConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateJvmConfigurationResponse) SetHeaders(v map[string]*string) *UpdateJvmConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateJvmConfigurationResponse) SetBody(v *UpdateJvmConfigurationResponseBody) *UpdateJvmConfigurationResponse {
	s.Body = v
	return s
}

type UpdateHealthCheckUrlQuery struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	HcURL *string `json:"hcURL,omitempty" xml:"hcURL,omitempty"`
}

func (s UpdateHealthCheckUrlQuery) String() string {
	return tea.Prettify(s)
}

func (s UpdateHealthCheckUrlQuery) GoString() string {
	return s.String()
}

func (s *UpdateHealthCheckUrlQuery) SetAppId(v string) *UpdateHealthCheckUrlQuery {
	s.AppId = &v
	return s
}

func (s *UpdateHealthCheckUrlQuery) SetHcURL(v string) *UpdateHealthCheckUrlQuery {
	s.HcURL = &v
	return s
}

type UpdateHealthCheckUrlRequest struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *UpdateHealthCheckUrlQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s UpdateHealthCheckUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHealthCheckUrlRequest) GoString() string {
	return s.String()
}

func (s *UpdateHealthCheckUrlRequest) SetHeaders(v map[string]*string) *UpdateHealthCheckUrlRequest {
	s.Headers = v
	return s
}

func (s *UpdateHealthCheckUrlRequest) SetQuery(v *UpdateHealthCheckUrlQuery) *UpdateHealthCheckUrlRequest {
	s.Query = v
	return s
}

type UpdateHealthCheckUrlResponseBody struct {
	Code           *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	HealthCheckURL *string `json:"HealthCheckURL,omitempty" xml:"HealthCheckURL,omitempty" require:"true"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateHealthCheckUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHealthCheckUrlResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHealthCheckUrlResponseBody) SetCode(v int) *UpdateHealthCheckUrlResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHealthCheckUrlResponseBody) SetHealthCheckURL(v string) *UpdateHealthCheckUrlResponseBody {
	s.HealthCheckURL = &v
	return s
}

func (s *UpdateHealthCheckUrlResponseBody) SetMessage(v string) *UpdateHealthCheckUrlResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHealthCheckUrlResponseBody) SetRequestId(v string) *UpdateHealthCheckUrlResponseBody {
	s.RequestId = &v
	return s
}

type UpdateHealthCheckUrlResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateHealthCheckUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateHealthCheckUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHealthCheckUrlResponse) GoString() string {
	return s.String()
}

func (s *UpdateHealthCheckUrlResponse) SetHeaders(v map[string]*string) *UpdateHealthCheckUrlResponse {
	s.Headers = v
	return s
}

func (s *UpdateHealthCheckUrlResponse) SetBody(v *UpdateHealthCheckUrlResponseBody) *UpdateHealthCheckUrlResponse {
	s.Body = v
	return s
}

type UpdateFlowControlQuery struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RuleId        *string `json:"RuleId,omitempty" xml:"RuleId,omitempty" require:"true"`
	ConsumerAppId *string `json:"ConsumerAppId,omitempty" xml:"ConsumerAppId,omitempty"`
	Granularity   *string `json:"Granularity,omitempty" xml:"Granularity,omitempty" require:"true"`
	MethodName    *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	RuleType      *string `json:"RuleType,omitempty" xml:"RuleType,omitempty" require:"true"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Strategy      *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	Threshold     *int    `json:"Threshold,omitempty" xml:"Threshold,omitempty" require:"true"`
	UrlVar        *string `json:"UrlVar,omitempty" xml:"UrlVar,omitempty"`
}

func (s UpdateFlowControlQuery) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowControlQuery) GoString() string {
	return s.String()
}

func (s *UpdateFlowControlQuery) SetAppId(v string) *UpdateFlowControlQuery {
	s.AppId = &v
	return s
}

func (s *UpdateFlowControlQuery) SetRuleId(v string) *UpdateFlowControlQuery {
	s.RuleId = &v
	return s
}

func (s *UpdateFlowControlQuery) SetConsumerAppId(v string) *UpdateFlowControlQuery {
	s.ConsumerAppId = &v
	return s
}

func (s *UpdateFlowControlQuery) SetGranularity(v string) *UpdateFlowControlQuery {
	s.Granularity = &v
	return s
}

func (s *UpdateFlowControlQuery) SetMethodName(v string) *UpdateFlowControlQuery {
	s.MethodName = &v
	return s
}

func (s *UpdateFlowControlQuery) SetRuleType(v string) *UpdateFlowControlQuery {
	s.RuleType = &v
	return s
}

func (s *UpdateFlowControlQuery) SetServiceName(v string) *UpdateFlowControlQuery {
	s.ServiceName = &v
	return s
}

func (s *UpdateFlowControlQuery) SetStrategy(v string) *UpdateFlowControlQuery {
	s.Strategy = &v
	return s
}

func (s *UpdateFlowControlQuery) SetThreshold(v int) *UpdateFlowControlQuery {
	s.Threshold = &v
	return s
}

func (s *UpdateFlowControlQuery) SetUrlVar(v string) *UpdateFlowControlQuery {
	s.UrlVar = &v
	return s
}

type UpdateFlowControlRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *UpdateFlowControlQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s UpdateFlowControlRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowControlRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowControlRequest) SetHeaders(v map[string]*string) *UpdateFlowControlRequest {
	s.Headers = v
	return s
}

func (s *UpdateFlowControlRequest) SetQuery(v *UpdateFlowControlQuery) *UpdateFlowControlRequest {
	s.Query = v
	return s
}

type UpdateFlowControlResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateFlowControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowControlResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlowControlResponseBody) SetCode(v int) *UpdateFlowControlResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFlowControlResponseBody) SetMessage(v string) *UpdateFlowControlResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFlowControlResponseBody) SetRequestId(v string) *UpdateFlowControlResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFlowControlResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFlowControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFlowControlResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowControlResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlowControlResponse) SetHeaders(v map[string]*string) *UpdateFlowControlResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlowControlResponse) SetBody(v *UpdateFlowControlResponseBody) *UpdateFlowControlResponse {
	s.Body = v
	return s
}

type UpdateDegradeControlQuery struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RuleId      *string `json:"RuleId,omitempty" xml:"RuleId,omitempty" require:"true"`
	Duration    *int    `json:"Duration,omitempty" xml:"Duration,omitempty" require:"true"`
	MethodName  *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	RtThreshold *int    `json:"RtThreshold,omitempty" xml:"RtThreshold,omitempty" require:"true"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	UrlVar      *string `json:"UrlVar,omitempty" xml:"UrlVar,omitempty"`
	RuleType    *string `json:"RuleType,omitempty" xml:"RuleType,omitempty" require:"true"`
}

func (s UpdateDegradeControlQuery) String() string {
	return tea.Prettify(s)
}

func (s UpdateDegradeControlQuery) GoString() string {
	return s.String()
}

func (s *UpdateDegradeControlQuery) SetAppId(v string) *UpdateDegradeControlQuery {
	s.AppId = &v
	return s
}

func (s *UpdateDegradeControlQuery) SetRuleId(v string) *UpdateDegradeControlQuery {
	s.RuleId = &v
	return s
}

func (s *UpdateDegradeControlQuery) SetDuration(v int) *UpdateDegradeControlQuery {
	s.Duration = &v
	return s
}

func (s *UpdateDegradeControlQuery) SetMethodName(v string) *UpdateDegradeControlQuery {
	s.MethodName = &v
	return s
}

func (s *UpdateDegradeControlQuery) SetRtThreshold(v int) *UpdateDegradeControlQuery {
	s.RtThreshold = &v
	return s
}

func (s *UpdateDegradeControlQuery) SetServiceName(v string) *UpdateDegradeControlQuery {
	s.ServiceName = &v
	return s
}

func (s *UpdateDegradeControlQuery) SetUrlVar(v string) *UpdateDegradeControlQuery {
	s.UrlVar = &v
	return s
}

func (s *UpdateDegradeControlQuery) SetRuleType(v string) *UpdateDegradeControlQuery {
	s.RuleType = &v
	return s
}

type UpdateDegradeControlRequest struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *UpdateDegradeControlQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s UpdateDegradeControlRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDegradeControlRequest) GoString() string {
	return s.String()
}

func (s *UpdateDegradeControlRequest) SetHeaders(v map[string]*string) *UpdateDegradeControlRequest {
	s.Headers = v
	return s
}

func (s *UpdateDegradeControlRequest) SetQuery(v *UpdateDegradeControlQuery) *UpdateDegradeControlRequest {
	s.Query = v
	return s
}

type UpdateDegradeControlResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateDegradeControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDegradeControlResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDegradeControlResponseBody) SetCode(v int) *UpdateDegradeControlResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDegradeControlResponseBody) SetMessage(v string) *UpdateDegradeControlResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDegradeControlResponseBody) SetRequestId(v string) *UpdateDegradeControlResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDegradeControlResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDegradeControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDegradeControlResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDegradeControlResponse) GoString() string {
	return s.String()
}

func (s *UpdateDegradeControlResponse) SetHeaders(v map[string]*string) *UpdateDegradeControlResponse {
	s.Headers = v
	return s
}

func (s *UpdateDegradeControlResponse) SetBody(v *UpdateDegradeControlResponseBody) *UpdateDegradeControlResponse {
	s.Body = v
	return s
}

type UpdateContainerConfigurationQuery struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ContextPath     *string `json:"ContextPath,omitempty" xml:"ContextPath,omitempty"`
	HttpPort        *int    `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	MaxThreads      *int    `json:"MaxThreads,omitempty" xml:"MaxThreads,omitempty"`
	URIEncoding     *string `json:"URIEncoding,omitempty" xml:"URIEncoding,omitempty"`
	UseBodyEncoding *bool   `json:"UseBodyEncoding,omitempty" xml:"UseBodyEncoding,omitempty"`
}

func (s UpdateContainerConfigurationQuery) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerConfigurationQuery) GoString() string {
	return s.String()
}

func (s *UpdateContainerConfigurationQuery) SetAppId(v string) *UpdateContainerConfigurationQuery {
	s.AppId = &v
	return s
}

func (s *UpdateContainerConfigurationQuery) SetGroupId(v string) *UpdateContainerConfigurationQuery {
	s.GroupId = &v
	return s
}

func (s *UpdateContainerConfigurationQuery) SetContextPath(v string) *UpdateContainerConfigurationQuery {
	s.ContextPath = &v
	return s
}

func (s *UpdateContainerConfigurationQuery) SetHttpPort(v int) *UpdateContainerConfigurationQuery {
	s.HttpPort = &v
	return s
}

func (s *UpdateContainerConfigurationQuery) SetMaxThreads(v int) *UpdateContainerConfigurationQuery {
	s.MaxThreads = &v
	return s
}

func (s *UpdateContainerConfigurationQuery) SetURIEncoding(v string) *UpdateContainerConfigurationQuery {
	s.URIEncoding = &v
	return s
}

func (s *UpdateContainerConfigurationQuery) SetUseBodyEncoding(v bool) *UpdateContainerConfigurationQuery {
	s.UseBodyEncoding = &v
	return s
}

type UpdateContainerConfigurationRequest struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *UpdateContainerConfigurationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s UpdateContainerConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateContainerConfigurationRequest) SetHeaders(v map[string]*string) *UpdateContainerConfigurationRequest {
	s.Headers = v
	return s
}

func (s *UpdateContainerConfigurationRequest) SetQuery(v *UpdateContainerConfigurationQuery) *UpdateContainerConfigurationRequest {
	s.Query = v
	return s
}

type UpdateContainerConfigurationResponseBody struct {
	Code                   *int                                                            `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message                *string                                                         `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId              *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ContainerConfiguration *UpdateContainerConfigurationResponseBodyContainerConfiguration `json:"ContainerConfiguration,omitempty" xml:"ContainerConfiguration,omitempty" require:"true" type:"Struct"`
}

func (s UpdateContainerConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContainerConfigurationResponseBody) SetCode(v int) *UpdateContainerConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateContainerConfigurationResponseBody) SetMessage(v string) *UpdateContainerConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateContainerConfigurationResponseBody) SetRequestId(v string) *UpdateContainerConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContainerConfigurationResponseBody) SetContainerConfiguration(v *UpdateContainerConfigurationResponseBodyContainerConfiguration) *UpdateContainerConfigurationResponseBody {
	s.ContainerConfiguration = v
	return s
}

type UpdateContainerConfigurationResponseBodyContainerConfiguration struct {
	ContextPath     *string `json:"ContextPath,omitempty" xml:"ContextPath,omitempty" require:"true"`
	HttpPort        *int    `json:"HttpPort,omitempty" xml:"HttpPort,omitempty" require:"true"`
	MaxThreads      *int    `json:"MaxThreads,omitempty" xml:"MaxThreads,omitempty" require:"true"`
	URIEncoding     *string `json:"URIEncoding,omitempty" xml:"URIEncoding,omitempty" require:"true"`
	UseBodyEncoding *bool   `json:"UseBodyEncoding,omitempty" xml:"UseBodyEncoding,omitempty" require:"true"`
}

func (s UpdateContainerConfigurationResponseBodyContainerConfiguration) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerConfigurationResponseBodyContainerConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateContainerConfigurationResponseBodyContainerConfiguration) SetContextPath(v string) *UpdateContainerConfigurationResponseBodyContainerConfiguration {
	s.ContextPath = &v
	return s
}

func (s *UpdateContainerConfigurationResponseBodyContainerConfiguration) SetHttpPort(v int) *UpdateContainerConfigurationResponseBodyContainerConfiguration {
	s.HttpPort = &v
	return s
}

func (s *UpdateContainerConfigurationResponseBodyContainerConfiguration) SetMaxThreads(v int) *UpdateContainerConfigurationResponseBodyContainerConfiguration {
	s.MaxThreads = &v
	return s
}

func (s *UpdateContainerConfigurationResponseBodyContainerConfiguration) SetURIEncoding(v string) *UpdateContainerConfigurationResponseBodyContainerConfiguration {
	s.URIEncoding = &v
	return s
}

func (s *UpdateContainerConfigurationResponseBodyContainerConfiguration) SetUseBodyEncoding(v bool) *UpdateContainerConfigurationResponseBodyContainerConfiguration {
	s.UseBodyEncoding = &v
	return s
}

type UpdateContainerConfigurationResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateContainerConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateContainerConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateContainerConfigurationResponse) SetHeaders(v map[string]*string) *UpdateContainerConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateContainerConfigurationResponse) SetBody(v *UpdateContainerConfigurationResponseBody) *UpdateContainerConfigurationResponse {
	s.Body = v
	return s
}

type UpdateApplicationBaseInfoQuery struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Desc    *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Owner   *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s UpdateApplicationBaseInfoQuery) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationBaseInfoQuery) GoString() string {
	return s.String()
}

func (s *UpdateApplicationBaseInfoQuery) SetAppId(v string) *UpdateApplicationBaseInfoQuery {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationBaseInfoQuery) SetAppName(v string) *UpdateApplicationBaseInfoQuery {
	s.AppName = &v
	return s
}

func (s *UpdateApplicationBaseInfoQuery) SetDesc(v string) *UpdateApplicationBaseInfoQuery {
	s.Desc = &v
	return s
}

func (s *UpdateApplicationBaseInfoQuery) SetOwner(v string) *UpdateApplicationBaseInfoQuery {
	s.Owner = &v
	return s
}

type UpdateApplicationBaseInfoRequest struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *UpdateApplicationBaseInfoQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s UpdateApplicationBaseInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationBaseInfoRequest) SetHeaders(v map[string]*string) *UpdateApplicationBaseInfoRequest {
	s.Headers = v
	return s
}

func (s *UpdateApplicationBaseInfoRequest) SetQuery(v *UpdateApplicationBaseInfoQuery) *UpdateApplicationBaseInfoRequest {
	s.Query = v
	return s
}

type UpdateApplicationBaseInfoResponseBody struct {
	Code       *int                                             `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message    *string                                          `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Applcation *UpdateApplicationBaseInfoResponseBodyApplcation `json:"Applcation,omitempty" xml:"Applcation,omitempty" require:"true" type:"Struct"`
}

func (s UpdateApplicationBaseInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationBaseInfoResponseBody) SetCode(v int) *UpdateApplicationBaseInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBody) SetMessage(v string) *UpdateApplicationBaseInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBody) SetRequestId(v string) *UpdateApplicationBaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBody) SetApplcation(v *UpdateApplicationBaseInfoResponseBodyApplcation) *UpdateApplicationBaseInfoResponseBody {
	s.Applcation = v
	return s
}

type UpdateApplicationBaseInfoResponseBodyApplcation struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Owner                *string `json:"Owner,omitempty" xml:"Owner,omitempty" require:"true"`
	InstanceCount        *int    `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty" require:"true"`
	RunningInstanceCount *int    `json:"RunningInstanceCount,omitempty" xml:"RunningInstanceCount,omitempty" require:"true"`
	Port                 *int    `json:"Port,omitempty" xml:"Port,omitempty" require:"true"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	SlbId                *string `json:"SlbId,omitempty" xml:"SlbId,omitempty" require:"true"`
	SlbIp                *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty" require:"true"`
	SlbPort              *int    `json:"SlbPort,omitempty" xml:"SlbPort,omitempty" require:"true"`
	ExtSlbId             *string `json:"ExtSlbId,omitempty" xml:"ExtSlbId,omitempty" require:"true"`
	ExtSlbIp             *string `json:"ExtSlbIp,omitempty" xml:"ExtSlbIp,omitempty" require:"true"`
	SlbName              *string `json:"SlbName,omitempty" xml:"SlbName,omitempty" require:"true"`
	ExtSlbName           *string `json:"ExtSlbName,omitempty" xml:"ExtSlbName,omitempty" require:"true"`
	ApplicationType      *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty" require:"true"`
	ClusterType          *int    `json:"ClusterType,omitempty" xml:"ClusterType,omitempty" require:"true"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	Dockerize            *bool   `json:"Dockerize,omitempty" xml:"Dockerize,omitempty" require:"true"`
	Cpu                  *int    `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Memory               *int    `json:"Memory,omitempty" xml:"Memory,omitempty" require:"true"`
	HealthCheckUrl       *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty" require:"true"`
	BuildPackageId       *int64  `json:"BuildPackageId,omitempty" xml:"BuildPackageId,omitempty" require:"true"`
	CreateTime           *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
}

func (s UpdateApplicationBaseInfoResponseBodyApplcation) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationBaseInfoResponseBodyApplcation) GoString() string {
	return s.String()
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetAppId(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetName(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.Name = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetRegionId(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.RegionId = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetDescription(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.Description = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetOwner(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.Owner = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetInstanceCount(v int) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.InstanceCount = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetRunningInstanceCount(v int) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.RunningInstanceCount = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetPort(v int) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.Port = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetUserId(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.UserId = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetSlbId(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.SlbId = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetSlbIp(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.SlbIp = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetSlbPort(v int) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.SlbPort = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetExtSlbId(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.ExtSlbId = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetExtSlbIp(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.ExtSlbIp = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetSlbName(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.SlbName = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetExtSlbName(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.ExtSlbName = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetApplicationType(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.ApplicationType = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetClusterType(v int) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.ClusterType = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetClusterId(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.ClusterId = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetDockerize(v bool) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.Dockerize = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetCpu(v int) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.Cpu = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetMemory(v int) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.Memory = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetHealthCheckUrl(v string) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.HealthCheckUrl = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetBuildPackageId(v int64) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.BuildPackageId = &v
	return s
}

func (s *UpdateApplicationBaseInfoResponseBodyApplcation) SetCreateTime(v int64) *UpdateApplicationBaseInfoResponseBodyApplcation {
	s.CreateTime = &v
	return s
}

type UpdateApplicationBaseInfoResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateApplicationBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateApplicationBaseInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationBaseInfoResponse) SetHeaders(v map[string]*string) *UpdateApplicationBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationBaseInfoResponse) SetBody(v *UpdateApplicationBaseInfoResponseBody) *UpdateApplicationBaseInfoResponse {
	s.Body = v
	return s
}

type UpdateAccountInfoQuery struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Telephone *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	Email     *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s UpdateAccountInfoQuery) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountInfoQuery) GoString() string {
	return s.String()
}

func (s *UpdateAccountInfoQuery) SetName(v string) *UpdateAccountInfoQuery {
	s.Name = &v
	return s
}

func (s *UpdateAccountInfoQuery) SetTelephone(v string) *UpdateAccountInfoQuery {
	s.Telephone = &v
	return s
}

func (s *UpdateAccountInfoQuery) SetEmail(v string) *UpdateAccountInfoQuery {
	s.Email = &v
	return s
}

type UpdateAccountInfoRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *UpdateAccountInfoQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s UpdateAccountInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccountInfoRequest) SetHeaders(v map[string]*string) *UpdateAccountInfoRequest {
	s.Headers = v
	return s
}

func (s *UpdateAccountInfoRequest) SetQuery(v *UpdateAccountInfoQuery) *UpdateAccountInfoRequest {
	s.Query = v
	return s
}

type UpdateAccountInfoResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateAccountInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccountInfoResponseBody) SetCode(v int) *UpdateAccountInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAccountInfoResponseBody) SetMessage(v string) *UpdateAccountInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAccountInfoResponseBody) SetRequestId(v string) *UpdateAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAccountInfoResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAccountInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccountInfoResponse) SetHeaders(v map[string]*string) *UpdateAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccountInfoResponse) SetBody(v *UpdateAccountInfoResponseBody) *UpdateAccountInfoResponse {
	s.Body = v
	return s
}

type UnbindSlbQuery struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty" require:"true"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
}

func (s UnbindSlbQuery) String() string {
	return tea.Prettify(s)
}

func (s UnbindSlbQuery) GoString() string {
	return s.String()
}

func (s *UnbindSlbQuery) SetAppId(v string) *UnbindSlbQuery {
	s.AppId = &v
	return s
}

func (s *UnbindSlbQuery) SetSlbId(v string) *UnbindSlbQuery {
	s.SlbId = &v
	return s
}

func (s *UnbindSlbQuery) SetType(v string) *UnbindSlbQuery {
	s.Type = &v
	return s
}

type UnbindSlbRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *UnbindSlbQuery    `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s UnbindSlbRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindSlbRequest) GoString() string {
	return s.String()
}

func (s *UnbindSlbRequest) SetHeaders(v map[string]*string) *UnbindSlbRequest {
	s.Headers = v
	return s
}

func (s *UnbindSlbRequest) SetQuery(v *UnbindSlbQuery) *UnbindSlbRequest {
	s.Query = v
	return s
}

type UnbindSlbResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UnbindSlbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindSlbResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindSlbResponseBody) SetCode(v int) *UnbindSlbResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindSlbResponseBody) SetData(v string) *UnbindSlbResponseBody {
	s.Data = &v
	return s
}

func (s *UnbindSlbResponseBody) SetMessage(v string) *UnbindSlbResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindSlbResponseBody) SetRequestId(v string) *UnbindSlbResponseBody {
	s.RequestId = &v
	return s
}

type UnbindSlbResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindSlbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindSlbResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindSlbResponse) GoString() string {
	return s.String()
}

func (s *UnbindSlbResponse) SetHeaders(v map[string]*string) *UnbindSlbResponse {
	s.Headers = v
	return s
}

func (s *UnbindSlbResponse) SetBody(v *UnbindSlbResponseBody) *UnbindSlbResponse {
	s.Body = v
	return s
}

type RollbackApplicationQuery struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	HistoryVersion *string `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty" require:"true"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	Batch          *int    `json:"Batch,omitempty" xml:"Batch,omitempty"`
	BatchWaitTime  *int    `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
}

func (s RollbackApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s RollbackApplicationQuery) GoString() string {
	return s.String()
}

func (s *RollbackApplicationQuery) SetAppId(v string) *RollbackApplicationQuery {
	s.AppId = &v
	return s
}

func (s *RollbackApplicationQuery) SetHistoryVersion(v string) *RollbackApplicationQuery {
	s.HistoryVersion = &v
	return s
}

func (s *RollbackApplicationQuery) SetGroupId(v string) *RollbackApplicationQuery {
	s.GroupId = &v
	return s
}

func (s *RollbackApplicationQuery) SetBatch(v int) *RollbackApplicationQuery {
	s.Batch = &v
	return s
}

func (s *RollbackApplicationQuery) SetBatchWaitTime(v int) *RollbackApplicationQuery {
	s.BatchWaitTime = &v
	return s
}

type RollbackApplicationRequest struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *RollbackApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s RollbackApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackApplicationRequest) GoString() string {
	return s.String()
}

func (s *RollbackApplicationRequest) SetHeaders(v map[string]*string) *RollbackApplicationRequest {
	s.Headers = v
	return s
}

func (s *RollbackApplicationRequest) SetQuery(v *RollbackApplicationQuery) *RollbackApplicationRequest {
	s.Query = v
	return s
}

type RollbackApplicationResponseBody struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RollbackApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackApplicationResponseBody) SetChangeOrderId(v string) *RollbackApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *RollbackApplicationResponseBody) SetCode(v int) *RollbackApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *RollbackApplicationResponseBody) SetMessage(v string) *RollbackApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *RollbackApplicationResponseBody) SetRequestId(v string) *RollbackApplicationResponseBody {
	s.RequestId = &v
	return s
}

type RollbackApplicationResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RollbackApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackApplicationResponse) GoString() string {
	return s.String()
}

func (s *RollbackApplicationResponse) SetHeaders(v map[string]*string) *RollbackApplicationResponse {
	s.Headers = v
	return s
}

func (s *RollbackApplicationResponse) SetBody(v *RollbackApplicationResponseBody) *RollbackApplicationResponse {
	s.Body = v
	return s
}

type QueryMonitorInfoQuery struct {
	Start      *int64                 `json:"Start,omitempty" xml:"Start,omitempty" require:"true"`
	End        *int64                 `json:"End,omitempty" xml:"End,omitempty" require:"true"`
	Metric     *string                `json:"Metric,omitempty" xml:"Metric,omitempty" require:"true"`
	Tags       map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true"`
	Aggregator *string                `json:"Aggregator,omitempty" xml:"Aggregator,omitempty" require:"true"`
	Interval   *string                `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s QueryMonitorInfoQuery) String() string {
	return tea.Prettify(s)
}

func (s QueryMonitorInfoQuery) GoString() string {
	return s.String()
}

func (s *QueryMonitorInfoQuery) SetStart(v int64) *QueryMonitorInfoQuery {
	s.Start = &v
	return s
}

func (s *QueryMonitorInfoQuery) SetEnd(v int64) *QueryMonitorInfoQuery {
	s.End = &v
	return s
}

func (s *QueryMonitorInfoQuery) SetMetric(v string) *QueryMonitorInfoQuery {
	s.Metric = &v
	return s
}

func (s *QueryMonitorInfoQuery) SetTags(v map[string]interface{}) *QueryMonitorInfoQuery {
	s.Tags = v
	return s
}

func (s *QueryMonitorInfoQuery) SetAggregator(v string) *QueryMonitorInfoQuery {
	s.Aggregator = &v
	return s
}

func (s *QueryMonitorInfoQuery) SetInterval(v string) *QueryMonitorInfoQuery {
	s.Interval = &v
	return s
}

type QueryMonitorInfoShrinkQuery struct {
	Start      *int64  `json:"Start,omitempty" xml:"Start,omitempty" require:"true"`
	End        *int64  `json:"End,omitempty" xml:"End,omitempty" require:"true"`
	Metric     *string `json:"Metric,omitempty" xml:"Metric,omitempty" require:"true"`
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true"`
	Aggregator *string `json:"Aggregator,omitempty" xml:"Aggregator,omitempty" require:"true"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s QueryMonitorInfoShrinkQuery) String() string {
	return tea.Prettify(s)
}

func (s QueryMonitorInfoShrinkQuery) GoString() string {
	return s.String()
}

func (s *QueryMonitorInfoShrinkQuery) SetStart(v int64) *QueryMonitorInfoShrinkQuery {
	s.Start = &v
	return s
}

func (s *QueryMonitorInfoShrinkQuery) SetEnd(v int64) *QueryMonitorInfoShrinkQuery {
	s.End = &v
	return s
}

func (s *QueryMonitorInfoShrinkQuery) SetMetric(v string) *QueryMonitorInfoShrinkQuery {
	s.Metric = &v
	return s
}

func (s *QueryMonitorInfoShrinkQuery) SetTagsShrink(v string) *QueryMonitorInfoShrinkQuery {
	s.TagsShrink = &v
	return s
}

func (s *QueryMonitorInfoShrinkQuery) SetAggregator(v string) *QueryMonitorInfoShrinkQuery {
	s.Aggregator = &v
	return s
}

func (s *QueryMonitorInfoShrinkQuery) SetInterval(v string) *QueryMonitorInfoShrinkQuery {
	s.Interval = &v
	return s
}

type QueryMonitorInfoRequest struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *QueryMonitorInfoQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s QueryMonitorInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMonitorInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryMonitorInfoRequest) SetHeaders(v map[string]*string) *QueryMonitorInfoRequest {
	s.Headers = v
	return s
}

func (s *QueryMonitorInfoRequest) SetQuery(v *QueryMonitorInfoQuery) *QueryMonitorInfoRequest {
	s.Query = v
	return s
}

type QueryMonitorInfoShrinkRequest struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *QueryMonitorInfoShrinkQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s QueryMonitorInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMonitorInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMonitorInfoShrinkRequest) SetHeaders(v map[string]*string) *QueryMonitorInfoShrinkRequest {
	s.Headers = v
	return s
}

func (s *QueryMonitorInfoShrinkRequest) SetQuery(v *QueryMonitorInfoShrinkQuery) *QueryMonitorInfoShrinkRequest {
	s.Query = v
	return s
}

type QueryMonitorInfoResponseBody struct {
	Code        *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	MonitorInfo *string `json:"MonitorInfo,omitempty" xml:"MonitorInfo,omitempty" require:"true"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s QueryMonitorInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMonitorInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMonitorInfoResponseBody) SetCode(v int) *QueryMonitorInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMonitorInfoResponseBody) SetMessage(v string) *QueryMonitorInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMonitorInfoResponseBody) SetMonitorInfo(v string) *QueryMonitorInfoResponseBody {
	s.MonitorInfo = &v
	return s
}

func (s *QueryMonitorInfoResponseBody) SetRequestId(v string) *QueryMonitorInfoResponseBody {
	s.RequestId = &v
	return s
}

type QueryMonitorInfoResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMonitorInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMonitorInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMonitorInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryMonitorInfoResponse) SetHeaders(v map[string]*string) *QueryMonitorInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryMonitorInfoResponse) SetBody(v *QueryMonitorInfoResponseBody) *QueryMonitorInfoResponse {
	s.Body = v
	return s
}

type QueryMigrateRegionListQuery struct {
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
}

func (s QueryMigrateRegionListQuery) String() string {
	return tea.Prettify(s)
}

func (s QueryMigrateRegionListQuery) GoString() string {
	return s.String()
}

func (s *QueryMigrateRegionListQuery) SetLogicalRegionId(v string) *QueryMigrateRegionListQuery {
	s.LogicalRegionId = &v
	return s
}

type QueryMigrateRegionListRequest struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *QueryMigrateRegionListQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s QueryMigrateRegionListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMigrateRegionListRequest) GoString() string {
	return s.String()
}

func (s *QueryMigrateRegionListRequest) SetHeaders(v map[string]*string) *QueryMigrateRegionListRequest {
	s.Headers = v
	return s
}

func (s *QueryMigrateRegionListRequest) SetQuery(v *QueryMigrateRegionListQuery) *QueryMigrateRegionListRequest {
	s.Query = v
	return s
}

type QueryMigrateRegionListResponseBody struct {
	Code             *int                                                `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message          *string                                             `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionEntityList *QueryMigrateRegionListResponseBodyRegionEntityList `json:"RegionEntityList,omitempty" xml:"RegionEntityList,omitempty" require:"true" type:"Struct"`
}

func (s QueryMigrateRegionListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMigrateRegionListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMigrateRegionListResponseBody) SetCode(v int) *QueryMigrateRegionListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMigrateRegionListResponseBody) SetMessage(v string) *QueryMigrateRegionListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMigrateRegionListResponseBody) SetRequestId(v string) *QueryMigrateRegionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMigrateRegionListResponseBody) SetRegionEntityList(v *QueryMigrateRegionListResponseBodyRegionEntityList) *QueryMigrateRegionListResponseBody {
	s.RegionEntityList = v
	return s
}

type QueryMigrateRegionListResponseBodyRegionEntityList struct {
	RegionEntity []*QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity `json:"RegionEntity,omitempty" xml:"RegionEntity,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMigrateRegionListResponseBodyRegionEntityList) String() string {
	return tea.Prettify(s)
}

func (s QueryMigrateRegionListResponseBodyRegionEntityList) GoString() string {
	return s.String()
}

func (s *QueryMigrateRegionListResponseBodyRegionEntityList) SetRegionEntity(v []*QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity) *QueryMigrateRegionListResponseBodyRegionEntityList {
	s.RegionEntity = v
	return s
}

type QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity struct {
	RegionNo   *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty" require:"true"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty" require:"true"`
}

func (s QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity) String() string {
	return tea.Prettify(s)
}

func (s QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity) GoString() string {
	return s.String()
}

func (s *QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity) SetRegionNo(v string) *QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity {
	s.RegionNo = &v
	return s
}

func (s *QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity) SetRegionName(v string) *QueryMigrateRegionListResponseBodyRegionEntityListRegionEntity {
	s.RegionName = &v
	return s
}

type QueryMigrateRegionListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMigrateRegionListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMigrateRegionListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMigrateRegionListResponse) GoString() string {
	return s.String()
}

func (s *QueryMigrateRegionListResponse) SetHeaders(v map[string]*string) *QueryMigrateRegionListResponse {
	s.Headers = v
	return s
}

func (s *QueryMigrateRegionListResponse) SetBody(v *QueryMigrateRegionListResponseBody) *QueryMigrateRegionListResponse {
	s.Body = v
	return s
}

type QueryMigrateEcuListQuery struct {
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
}

func (s QueryMigrateEcuListQuery) String() string {
	return tea.Prettify(s)
}

func (s QueryMigrateEcuListQuery) GoString() string {
	return s.String()
}

func (s *QueryMigrateEcuListQuery) SetLogicalRegionId(v string) *QueryMigrateEcuListQuery {
	s.LogicalRegionId = &v
	return s
}

type QueryMigrateEcuListRequest struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *QueryMigrateEcuListQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s QueryMigrateEcuListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMigrateEcuListRequest) GoString() string {
	return s.String()
}

func (s *QueryMigrateEcuListRequest) SetHeaders(v map[string]*string) *QueryMigrateEcuListRequest {
	s.Headers = v
	return s
}

func (s *QueryMigrateEcuListRequest) SetQuery(v *QueryMigrateEcuListQuery) *QueryMigrateEcuListRequest {
	s.Query = v
	return s
}

type QueryMigrateEcuListResponseBody struct {
	Code          *int                                          `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string                                       `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	EcuEntityList *QueryMigrateEcuListResponseBodyEcuEntityList `json:"EcuEntityList,omitempty" xml:"EcuEntityList,omitempty" require:"true" type:"Struct"`
}

func (s QueryMigrateEcuListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMigrateEcuListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMigrateEcuListResponseBody) SetCode(v int) *QueryMigrateEcuListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMigrateEcuListResponseBody) SetMessage(v string) *QueryMigrateEcuListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMigrateEcuListResponseBody) SetRequestId(v string) *QueryMigrateEcuListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMigrateEcuListResponseBody) SetEcuEntityList(v *QueryMigrateEcuListResponseBodyEcuEntityList) *QueryMigrateEcuListResponseBody {
	s.EcuEntityList = v
	return s
}

type QueryMigrateEcuListResponseBodyEcuEntityList struct {
	EcuEntity []*QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity `json:"EcuEntity,omitempty" xml:"EcuEntity,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMigrateEcuListResponseBodyEcuEntityList) String() string {
	return tea.Prettify(s)
}

func (s QueryMigrateEcuListResponseBodyEcuEntityList) GoString() string {
	return s.String()
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityList) SetEcuEntity(v []*QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) *QueryMigrateEcuListResponseBodyEcuEntityList {
	s.EcuEntity = v
	return s
}

type QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity struct {
	EcuId         *string `json:"EcuId,omitempty" xml:"EcuId,omitempty" require:"true"`
	Online        *bool   `json:"Online,omitempty" xml:"Online,omitempty" require:"true"`
	DockerEnv     *bool   `json:"DockerEnv,omitempty" xml:"DockerEnv,omitempty" require:"true"`
	CreateTime    *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime    *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	IpAddr        *string `json:"IpAddr,omitempty" xml:"IpAddr,omitempty" require:"true"`
	HeartbeatTime *int64  `json:"HeartbeatTime,omitempty" xml:"HeartbeatTime,omitempty" require:"true"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	ZoneId        *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	AvailableCpu  *int    `json:"AvailableCpu,omitempty" xml:"AvailableCpu,omitempty" require:"true"`
	AvailableMem  *int    `json:"AvailableMem,omitempty" xml:"AvailableMem,omitempty" require:"true"`
	Cpu           *int    `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Mem           *int    `json:"Mem,omitempty" xml:"Mem,omitempty" require:"true"`
}

func (s QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) String() string {
	return tea.Prettify(s)
}

func (s QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) GoString() string {
	return s.String()
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetEcuId(v string) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.EcuId = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetOnline(v bool) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.Online = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetDockerEnv(v bool) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.DockerEnv = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetCreateTime(v int64) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.CreateTime = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetUpdateTime(v int64) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.UpdateTime = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetIpAddr(v string) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.IpAddr = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetHeartbeatTime(v int64) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.HeartbeatTime = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetUserId(v string) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.UserId = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetName(v string) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.Name = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetZoneId(v string) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.ZoneId = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetRegionId(v string) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.RegionId = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetInstanceId(v string) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.InstanceId = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetVpcId(v string) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.VpcId = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetAvailableCpu(v int) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.AvailableCpu = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetAvailableMem(v int) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.AvailableMem = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetCpu(v int) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.Cpu = &v
	return s
}

func (s *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity) SetMem(v int) *QueryMigrateEcuListResponseBodyEcuEntityListEcuEntity {
	s.Mem = &v
	return s
}

type QueryMigrateEcuListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMigrateEcuListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMigrateEcuListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMigrateEcuListResponse) GoString() string {
	return s.String()
}

func (s *QueryMigrateEcuListResponse) SetHeaders(v map[string]*string) *QueryMigrateEcuListResponse {
	s.Headers = v
	return s
}

func (s *QueryMigrateEcuListResponse) SetBody(v *QueryMigrateEcuListResponseBody) *QueryMigrateEcuListResponse {
	s.Body = v
	return s
}

type QueryConfigCenterQuery struct {
	DataId          *string `json:"DataId,omitempty" xml:"DataId,omitempty" require:"true"`
	Group           *string `json:"Group,omitempty" xml:"Group,omitempty" require:"true"`
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty" require:"true"`
}

func (s QueryConfigCenterQuery) String() string {
	return tea.Prettify(s)
}

func (s QueryConfigCenterQuery) GoString() string {
	return s.String()
}

func (s *QueryConfigCenterQuery) SetDataId(v string) *QueryConfigCenterQuery {
	s.DataId = &v
	return s
}

func (s *QueryConfigCenterQuery) SetGroup(v string) *QueryConfigCenterQuery {
	s.Group = &v
	return s
}

func (s *QueryConfigCenterQuery) SetLogicalRegionId(v string) *QueryConfigCenterQuery {
	s.LogicalRegionId = &v
	return s
}

type QueryConfigCenterRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *QueryConfigCenterQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s QueryConfigCenterRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConfigCenterRequest) GoString() string {
	return s.String()
}

func (s *QueryConfigCenterRequest) SetHeaders(v map[string]*string) *QueryConfigCenterRequest {
	s.Headers = v
	return s
}

func (s *QueryConfigCenterRequest) SetQuery(v *QueryConfigCenterQuery) *QueryConfigCenterRequest {
	s.Query = v
	return s
}

type QueryConfigCenterResponseBody struct {
	Code             *int                                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message          *string                                        `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId        *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ConfigCenterInfo *QueryConfigCenterResponseBodyConfigCenterInfo `json:"configCenterInfo,omitempty" xml:"configCenterInfo,omitempty" require:"true" type:"Struct"`
}

func (s QueryConfigCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryConfigCenterResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConfigCenterResponseBody) SetCode(v int) *QueryConfigCenterResponseBody {
	s.Code = &v
	return s
}

func (s *QueryConfigCenterResponseBody) SetMessage(v string) *QueryConfigCenterResponseBody {
	s.Message = &v
	return s
}

func (s *QueryConfigCenterResponseBody) SetRequestId(v string) *QueryConfigCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryConfigCenterResponseBody) SetConfigCenterInfo(v *QueryConfigCenterResponseBodyConfigCenterInfo) *QueryConfigCenterResponseBody {
	s.ConfigCenterInfo = v
	return s
}

type QueryConfigCenterResponseBodyConfigCenterInfo struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty" require:"true"`
	DataId  *string `json:"DataId,omitempty" xml:"DataId,omitempty" require:"true"`
	Group   *string `json:"Group,omitempty" xml:"Group,omitempty" require:"true"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
}

func (s QueryConfigCenterResponseBodyConfigCenterInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryConfigCenterResponseBodyConfigCenterInfo) GoString() string {
	return s.String()
}

func (s *QueryConfigCenterResponseBodyConfigCenterInfo) SetAppName(v string) *QueryConfigCenterResponseBodyConfigCenterInfo {
	s.AppName = &v
	return s
}

func (s *QueryConfigCenterResponseBodyConfigCenterInfo) SetContent(v string) *QueryConfigCenterResponseBodyConfigCenterInfo {
	s.Content = &v
	return s
}

func (s *QueryConfigCenterResponseBodyConfigCenterInfo) SetDataId(v string) *QueryConfigCenterResponseBodyConfigCenterInfo {
	s.DataId = &v
	return s
}

func (s *QueryConfigCenterResponseBodyConfigCenterInfo) SetGroup(v string) *QueryConfigCenterResponseBodyConfigCenterInfo {
	s.Group = &v
	return s
}

func (s *QueryConfigCenterResponseBodyConfigCenterInfo) SetId(v string) *QueryConfigCenterResponseBodyConfigCenterInfo {
	s.Id = &v
	return s
}

type QueryConfigCenterResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryConfigCenterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryConfigCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryConfigCenterResponse) GoString() string {
	return s.String()
}

func (s *QueryConfigCenterResponse) SetHeaders(v map[string]*string) *QueryConfigCenterResponse {
	s.Headers = v
	return s
}

func (s *QueryConfigCenterResponse) SetBody(v *QueryConfigCenterResponseBody) *QueryConfigCenterResponse {
	s.Body = v
	return s
}

type QueryApplicationStatusQuery struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s QueryApplicationStatusQuery) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationStatusQuery) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusQuery) SetAppId(v string) *QueryApplicationStatusQuery {
	s.AppId = &v
	return s
}

type QueryApplicationStatusRequest struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *QueryApplicationStatusQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s QueryApplicationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusRequest) SetHeaders(v map[string]*string) *QueryApplicationStatusRequest {
	s.Headers = v
	return s
}

func (s *QueryApplicationStatusRequest) SetQuery(v *QueryApplicationStatusQuery) *QueryApplicationStatusRequest {
	s.Query = v
	return s
}

type QueryApplicationStatusResponseBody struct {
	Code      *int                                       `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AppInfo   *QueryApplicationStatusResponseBodyAppInfo `json:"AppInfo,omitempty" xml:"AppInfo,omitempty" require:"true" type:"Struct"`
}

func (s QueryApplicationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBody) SetCode(v int) *QueryApplicationStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryApplicationStatusResponseBody) SetMessage(v string) *QueryApplicationStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryApplicationStatusResponseBody) SetRequestId(v string) *QueryApplicationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryApplicationStatusResponseBody) SetAppInfo(v *QueryApplicationStatusResponseBodyAppInfo) *QueryApplicationStatusResponseBody {
	s.AppInfo = v
	return s
}

type QueryApplicationStatusResponseBodyAppInfo struct {
	EcuList          *QueryApplicationStatusResponseBodyAppInfoEcuList          `json:"EcuList,omitempty" xml:"EcuList,omitempty" require:"true" type:"Struct"`
	EccList          *QueryApplicationStatusResponseBodyAppInfoEccList          `json:"EccList,omitempty" xml:"EccList,omitempty" require:"true" type:"Struct"`
	GroupList        *QueryApplicationStatusResponseBodyAppInfoGroupList        `json:"GroupList,omitempty" xml:"GroupList,omitempty" require:"true" type:"Struct"`
	DeployRecordList *QueryApplicationStatusResponseBodyAppInfoDeployRecordList `json:"DeployRecordList,omitempty" xml:"DeployRecordList,omitempty" require:"true" type:"Struct"`
	Application      *QueryApplicationStatusResponseBodyAppInfoApplication      `json:"Application,omitempty" xml:"Application,omitempty" require:"true" type:"Struct"`
}

func (s QueryApplicationStatusResponseBodyAppInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfo) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfo) SetEcuList(v *QueryApplicationStatusResponseBodyAppInfoEcuList) *QueryApplicationStatusResponseBodyAppInfo {
	s.EcuList = v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfo) SetEccList(v *QueryApplicationStatusResponseBodyAppInfoEccList) *QueryApplicationStatusResponseBodyAppInfo {
	s.EccList = v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfo) SetGroupList(v *QueryApplicationStatusResponseBodyAppInfoGroupList) *QueryApplicationStatusResponseBodyAppInfo {
	s.GroupList = v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfo) SetDeployRecordList(v *QueryApplicationStatusResponseBodyAppInfoDeployRecordList) *QueryApplicationStatusResponseBodyAppInfo {
	s.DeployRecordList = v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfo) SetApplication(v *QueryApplicationStatusResponseBodyAppInfoApplication) *QueryApplicationStatusResponseBodyAppInfo {
	s.Application = v
	return s
}

type QueryApplicationStatusResponseBodyAppInfoEcuList struct {
	Ecu []*QueryApplicationStatusResponseBodyAppInfoEcuListEcu `json:"Ecu,omitempty" xml:"Ecu,omitempty" require:"true" type:"Repeated"`
}

func (s QueryApplicationStatusResponseBodyAppInfoEcuList) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoEcuList) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuList) SetEcu(v []*QueryApplicationStatusResponseBodyAppInfoEcuListEcu) *QueryApplicationStatusResponseBodyAppInfoEcuList {
	s.Ecu = v
	return s
}

type QueryApplicationStatusResponseBodyAppInfoEcuListEcu struct {
	EcuId         *string `json:"EcuId,omitempty" xml:"EcuId,omitempty" require:"true"`
	Online        *bool   `json:"Online,omitempty" xml:"Online,omitempty" require:"true"`
	DockerEnv     *bool   `json:"DockerEnv,omitempty" xml:"DockerEnv,omitempty" require:"true"`
	CreateTime    *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime    *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	IpAddr        *string `json:"IpAddr,omitempty" xml:"IpAddr,omitempty" require:"true"`
	HeartbeatTime *int64  `json:"HeartbeatTime,omitempty" xml:"HeartbeatTime,omitempty" require:"true"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	ZoneId        *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	AvailableCpu  *int    `json:"AvailableCpu,omitempty" xml:"AvailableCpu,omitempty" require:"true"`
	AvailableMem  *int    `json:"AvailableMem,omitempty" xml:"AvailableMem,omitempty" require:"true"`
}

func (s QueryApplicationStatusResponseBodyAppInfoEcuListEcu) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoEcuListEcu) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetEcuId(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.EcuId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetOnline(v bool) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.Online = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetDockerEnv(v bool) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.DockerEnv = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetCreateTime(v int64) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.CreateTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetUpdateTime(v int64) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.UpdateTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetIpAddr(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.IpAddr = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetHeartbeatTime(v int64) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.HeartbeatTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetUserId(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.UserId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetGroupId(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.GroupId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetName(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.Name = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetZoneId(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.ZoneId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetRegionId(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.RegionId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetInstanceId(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.InstanceId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetVpcId(v string) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.VpcId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetAvailableCpu(v int) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.AvailableCpu = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEcuListEcu) SetAvailableMem(v int) *QueryApplicationStatusResponseBodyAppInfoEcuListEcu {
	s.AvailableMem = &v
	return s
}

type QueryApplicationStatusResponseBodyAppInfoEccList struct {
	Ecc []*QueryApplicationStatusResponseBodyAppInfoEccListEcc `json:"Ecc,omitempty" xml:"Ecc,omitempty" require:"true" type:"Repeated"`
}

func (s QueryApplicationStatusResponseBodyAppInfoEccList) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoEccList) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccList) SetEcc(v []*QueryApplicationStatusResponseBodyAppInfoEccListEcc) *QueryApplicationStatusResponseBodyAppInfoEccList {
	s.Ecc = v
	return s
}

type QueryApplicationStatusResponseBodyAppInfoEccListEcc struct {
	EccId           *string `json:"EccId,omitempty" xml:"EccId,omitempty" require:"true"`
	EcuId           *string `json:"EcuId,omitempty" xml:"EcuId,omitempty" require:"true"`
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	AppState        *int    `json:"AppState,omitempty" xml:"AppState,omitempty" require:"true"`
	TaskState       *int    `json:"TaskState,omitempty" xml:"TaskState,omitempty" require:"true"`
	CreateTime      *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime      *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty" require:"true"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	ContainerStatus *string `json:"ContainerStatus,omitempty" xml:"ContainerStatus,omitempty" require:"true"`
}

func (s QueryApplicationStatusResponseBodyAppInfoEccListEcc) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoEccListEcc) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetEccId(v string) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.EccId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetEcuId(v string) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.EcuId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetAppId(v string) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.AppId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetAppState(v int) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.AppState = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetTaskState(v int) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.TaskState = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetCreateTime(v int64) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.CreateTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetUpdateTime(v int64) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.UpdateTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetIp(v string) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.Ip = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetVpcId(v string) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.VpcId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetGroupId(v string) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.GroupId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoEccListEcc) SetContainerStatus(v string) *QueryApplicationStatusResponseBodyAppInfoEccListEcc {
	s.ContainerStatus = &v
	return s
}

type QueryApplicationStatusResponseBodyAppInfoGroupList struct {
	Group []*QueryApplicationStatusResponseBodyAppInfoGroupListGroup `json:"Group,omitempty" xml:"Group,omitempty" require:"true" type:"Repeated"`
}

func (s QueryApplicationStatusResponseBodyAppInfoGroupList) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoGroupList) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupList) SetGroup(v []*QueryApplicationStatusResponseBodyAppInfoGroupListGroup) *QueryApplicationStatusResponseBodyAppInfoGroupList {
	s.Group = v
	return s
}

type QueryApplicationStatusResponseBodyAppInfoGroupListGroup struct {
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	GroupName        *string `json:"GroupName,omitempty" xml:"GroupName,omitempty" require:"true"`
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	PackageVersionId *string `json:"PackageVersionId,omitempty" xml:"PackageVersionId,omitempty" require:"true"`
	AppVersionId     *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty" require:"true"`
	GroupType        *int    `json:"GroupType,omitempty" xml:"GroupType,omitempty" require:"true"`
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	CreateTime       *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime       *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
}

func (s QueryApplicationStatusResponseBodyAppInfoGroupListGroup) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoGroupListGroup) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetGroupId(v string) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.GroupId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetGroupName(v string) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.GroupName = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetAppId(v string) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.AppId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetPackageVersionId(v string) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.PackageVersionId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetAppVersionId(v string) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.AppVersionId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetGroupType(v int) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.GroupType = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetClusterId(v string) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.ClusterId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetCreateTime(v int64) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.CreateTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoGroupListGroup) SetUpdateTime(v int64) *QueryApplicationStatusResponseBodyAppInfoGroupListGroup {
	s.UpdateTime = &v
	return s
}

type QueryApplicationStatusResponseBodyAppInfoDeployRecordList struct {
	DeployRecord []*QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord `json:"DeployRecord,omitempty" xml:"DeployRecord,omitempty" require:"true" type:"Repeated"`
}

func (s QueryApplicationStatusResponseBodyAppInfoDeployRecordList) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoDeployRecordList) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordList) SetDeployRecord(v []*QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) *QueryApplicationStatusResponseBodyAppInfoDeployRecordList {
	s.DeployRecord = v
	return s
}

type QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord struct {
	DeployRecordId   *string `json:"DeployRecordId,omitempty" xml:"DeployRecordId,omitempty" require:"true"`
	EccId            *string `json:"EccId,omitempty" xml:"EccId,omitempty" require:"true"`
	EcuId            *string `json:"EcuId,omitempty" xml:"EcuId,omitempty" require:"true"`
	PackageVersionId *string `json:"PackageVersionId,omitempty" xml:"PackageVersionId,omitempty" require:"true"`
	PackageMd5       *string `json:"PackageMd5,omitempty" xml:"PackageMd5,omitempty" require:"true"`
	CreateTime       *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
}

func (s QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) SetDeployRecordId(v string) *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord {
	s.DeployRecordId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) SetEccId(v string) *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord {
	s.EccId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) SetEcuId(v string) *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord {
	s.EcuId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) SetPackageVersionId(v string) *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord {
	s.PackageVersionId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) SetPackageMd5(v string) *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord {
	s.PackageMd5 = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord) SetCreateTime(v int64) *QueryApplicationStatusResponseBodyAppInfoDeployRecordListDeployRecord {
	s.CreateTime = &v
	return s
}

type QueryApplicationStatusResponseBodyAppInfoApplication struct {
	ApplicationId        *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty" require:"true"`
	BuildPackageId       *int    `json:"BuildPackageId,omitempty" xml:"BuildPackageId,omitempty" require:"true"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	Cpu                  *int    `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	CreateTime           *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	Dockerize            *bool   `json:"Dockerize,omitempty" xml:"Dockerize,omitempty" require:"true"`
	Email                *string `json:"Email,omitempty" xml:"Email,omitempty" require:"true"`
	HealthCheckUrl       *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty" require:"true"`
	InstanceCount        *int    `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty" require:"true"`
	LaunchTime           *int64  `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty" require:"true"`
	Memory               *int    `json:"Memory,omitempty" xml:"Memory,omitempty" require:"true"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Owner                *string `json:"Owner,omitempty" xml:"Owner,omitempty" require:"true"`
	Phone                *string `json:"Phone,omitempty" xml:"Phone,omitempty" require:"true"`
	Port                 *int    `json:"Port,omitempty" xml:"Port,omitempty" require:"true"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	RunningInstanceCount *int    `json:"RunningInstanceCount,omitempty" xml:"RunningInstanceCount,omitempty" require:"true"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
}

func (s QueryApplicationStatusResponseBodyAppInfoApplication) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationStatusResponseBodyAppInfoApplication) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetApplicationId(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.ApplicationId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetBuildPackageId(v int) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.BuildPackageId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetClusterId(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.ClusterId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetCpu(v int) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.Cpu = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetCreateTime(v int64) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.CreateTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetDockerize(v bool) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.Dockerize = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetEmail(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.Email = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetHealthCheckUrl(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.HealthCheckUrl = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetInstanceCount(v int) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.InstanceCount = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetLaunchTime(v int64) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.LaunchTime = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetMemory(v int) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.Memory = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetName(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.Name = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetOwner(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.Owner = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetPhone(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.Phone = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetPort(v int) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.Port = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetRegionId(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.RegionId = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetRunningInstanceCount(v int) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.RunningInstanceCount = &v
	return s
}

func (s *QueryApplicationStatusResponseBodyAppInfoApplication) SetUserId(v string) *QueryApplicationStatusResponseBodyAppInfoApplication {
	s.UserId = &v
	return s
}

type QueryApplicationStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryApplicationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryApplicationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryApplicationStatusResponse) SetHeaders(v map[string]*string) *QueryApplicationStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryApplicationStatusResponse) SetBody(v *QueryApplicationStatusResponseBody) *QueryApplicationStatusResponse {
	s.Body = v
	return s
}

type MigrateEcuQuery struct {
	InstanceIds     *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" require:"true"`
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
}

func (s MigrateEcuQuery) String() string {
	return tea.Prettify(s)
}

func (s MigrateEcuQuery) GoString() string {
	return s.String()
}

func (s *MigrateEcuQuery) SetInstanceIds(v string) *MigrateEcuQuery {
	s.InstanceIds = &v
	return s
}

func (s *MigrateEcuQuery) SetLogicalRegionId(v string) *MigrateEcuQuery {
	s.LogicalRegionId = &v
	return s
}

type MigrateEcuRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *MigrateEcuQuery   `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s MigrateEcuRequest) String() string {
	return tea.Prettify(s)
}

func (s MigrateEcuRequest) GoString() string {
	return s.String()
}

func (s *MigrateEcuRequest) SetHeaders(v map[string]*string) *MigrateEcuRequest {
	s.Headers = v
	return s
}

func (s *MigrateEcuRequest) SetQuery(v *MigrateEcuQuery) *MigrateEcuRequest {
	s.Query = v
	return s
}

type MigrateEcuResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s MigrateEcuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MigrateEcuResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateEcuResponseBody) SetCode(v int) *MigrateEcuResponseBody {
	s.Code = &v
	return s
}

func (s *MigrateEcuResponseBody) SetMessage(v string) *MigrateEcuResponseBody {
	s.Message = &v
	return s
}

func (s *MigrateEcuResponseBody) SetData(v string) *MigrateEcuResponseBody {
	s.Data = &v
	return s
}

func (s *MigrateEcuResponseBody) SetRequestId(v string) *MigrateEcuResponseBody {
	s.RequestId = &v
	return s
}

type MigrateEcuResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MigrateEcuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MigrateEcuResponse) String() string {
	return tea.Prettify(s)
}

func (s MigrateEcuResponse) GoString() string {
	return s.String()
}

func (s *MigrateEcuResponse) SetHeaders(v map[string]*string) *MigrateEcuResponse {
	s.Headers = v
	return s
}

func (s *MigrateEcuResponse) SetBody(v *MigrateEcuResponseBody) *MigrateEcuResponse {
	s.Body = v
	return s
}

type ListUserDefineRegionQuery struct {
	DebugEnable *bool `json:"DebugEnable,omitempty" xml:"DebugEnable,omitempty"`
}

func (s ListUserDefineRegionQuery) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefineRegionQuery) GoString() string {
	return s.String()
}

func (s *ListUserDefineRegionQuery) SetDebugEnable(v bool) *ListUserDefineRegionQuery {
	s.DebugEnable = &v
	return s
}

type ListUserDefineRegionRequest struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListUserDefineRegionQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s ListUserDefineRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefineRegionRequest) GoString() string {
	return s.String()
}

func (s *ListUserDefineRegionRequest) SetHeaders(v map[string]*string) *ListUserDefineRegionRequest {
	s.Headers = v
	return s
}

func (s *ListUserDefineRegionRequest) SetQuery(v *ListUserDefineRegionQuery) *ListUserDefineRegionRequest {
	s.Query = v
	return s
}

type ListUserDefineRegionResponseBody struct {
	Code                 *int                                                  `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message              *string                                               `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId            *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UserDefineRegionList *ListUserDefineRegionResponseBodyUserDefineRegionList `json:"UserDefineRegionList,omitempty" xml:"UserDefineRegionList,omitempty" require:"true" type:"Struct"`
}

func (s ListUserDefineRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefineRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserDefineRegionResponseBody) SetCode(v int) *ListUserDefineRegionResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserDefineRegionResponseBody) SetMessage(v string) *ListUserDefineRegionResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserDefineRegionResponseBody) SetRequestId(v string) *ListUserDefineRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserDefineRegionResponseBody) SetUserDefineRegionList(v *ListUserDefineRegionResponseBodyUserDefineRegionList) *ListUserDefineRegionResponseBody {
	s.UserDefineRegionList = v
	return s
}

type ListUserDefineRegionResponseBodyUserDefineRegionList struct {
	UserDefineRegionEntity []*ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity `json:"UserDefineRegionEntity,omitempty" xml:"UserDefineRegionEntity,omitempty" require:"true" type:"Repeated"`
}

func (s ListUserDefineRegionResponseBodyUserDefineRegionList) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefineRegionResponseBodyUserDefineRegionList) GoString() string {
	return s.String()
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionList) SetUserDefineRegionEntity(v []*ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) *ListUserDefineRegionResponseBodyUserDefineRegionList {
	s.UserDefineRegionEntity = v
	return s
}

type ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity struct {
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	RegionName   *string `json:"RegionName,omitempty" xml:"RegionName,omitempty" require:"true"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	BelongRegion *string `json:"BelongRegion,omitempty" xml:"BelongRegion,omitempty" require:"true"`
	DebugEnable  *bool   `json:"DebugEnable,omitempty" xml:"DebugEnable,omitempty" require:"true"`
}

func (s ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) GoString() string {
	return s.String()
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) SetId(v int64) *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	s.Id = &v
	return s
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) SetUserId(v string) *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	s.UserId = &v
	return s
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) SetRegionId(v string) *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	s.RegionId = &v
	return s
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) SetRegionName(v string) *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	s.RegionName = &v
	return s
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) SetDescription(v string) *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	s.Description = &v
	return s
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) SetBelongRegion(v string) *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	s.BelongRegion = &v
	return s
}

func (s *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity) SetDebugEnable(v bool) *ListUserDefineRegionResponseBodyUserDefineRegionListUserDefineRegionEntity {
	s.DebugEnable = &v
	return s
}

type ListUserDefineRegionResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserDefineRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserDefineRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserDefineRegionResponse) GoString() string {
	return s.String()
}

func (s *ListUserDefineRegionResponse) SetHeaders(v map[string]*string) *ListUserDefineRegionResponse {
	s.Headers = v
	return s
}

func (s *ListUserDefineRegionResponse) SetBody(v *ListUserDefineRegionResponseBody) *ListUserDefineRegionResponse {
	s.Body = v
	return s
}

type ListSubAccountRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ListSubAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubAccountRequest) GoString() string {
	return s.String()
}

func (s *ListSubAccountRequest) SetHeaders(v map[string]*string) *ListSubAccountRequest {
	s.Headers = v
	return s
}

type ListSubAccountResponseBody struct {
	Code           *int                                      `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SubAccountList *ListSubAccountResponseBodySubAccountList `json:"SubAccountList,omitempty" xml:"SubAccountList,omitempty" require:"true" type:"Struct"`
}

func (s ListSubAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubAccountResponseBody) SetCode(v int) *ListSubAccountResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubAccountResponseBody) SetMessage(v string) *ListSubAccountResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubAccountResponseBody) SetRequestId(v string) *ListSubAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubAccountResponseBody) SetSubAccountList(v *ListSubAccountResponseBodySubAccountList) *ListSubAccountResponseBody {
	s.SubAccountList = v
	return s
}

type ListSubAccountResponseBodySubAccountList struct {
	SubAccount []*ListSubAccountResponseBodySubAccountListSubAccount `json:"SubAccount,omitempty" xml:"SubAccount,omitempty" require:"true" type:"Repeated"`
}

func (s ListSubAccountResponseBodySubAccountList) String() string {
	return tea.Prettify(s)
}

func (s ListSubAccountResponseBodySubAccountList) GoString() string {
	return s.String()
}

func (s *ListSubAccountResponseBodySubAccountList) SetSubAccount(v []*ListSubAccountResponseBodySubAccountListSubAccount) *ListSubAccountResponseBodySubAccountList {
	s.SubAccount = v
	return s
}

type ListSubAccountResponseBodySubAccountListSubAccount struct {
	AdminUserId *string `json:"AdminUserId,omitempty" xml:"AdminUserId,omitempty" require:"true"`
	SubUserId   *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty" require:"true"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty" require:"true"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty" require:"true"`
	AdminUserKp *string `json:"AdminUserKp,omitempty" xml:"AdminUserKp,omitempty" require:"true"`
	SubUserKp   *string `json:"SubUserKp,omitempty" xml:"SubUserKp,omitempty" require:"true"`
	AdminEdasId *string `json:"AdminEdasId,omitempty" xml:"AdminEdasId,omitempty" require:"true"`
	SubEdasId   *string `json:"SubEdasId,omitempty" xml:"SubEdasId,omitempty" require:"true"`
}

func (s ListSubAccountResponseBodySubAccountListSubAccount) String() string {
	return tea.Prettify(s)
}

func (s ListSubAccountResponseBodySubAccountListSubAccount) GoString() string {
	return s.String()
}

func (s *ListSubAccountResponseBodySubAccountListSubAccount) SetAdminUserId(v string) *ListSubAccountResponseBodySubAccountListSubAccount {
	s.AdminUserId = &v
	return s
}

func (s *ListSubAccountResponseBodySubAccountListSubAccount) SetSubUserId(v string) *ListSubAccountResponseBodySubAccountListSubAccount {
	s.SubUserId = &v
	return s
}

func (s *ListSubAccountResponseBodySubAccountListSubAccount) SetEmail(v string) *ListSubAccountResponseBodySubAccountListSubAccount {
	s.Email = &v
	return s
}

func (s *ListSubAccountResponseBodySubAccountListSubAccount) SetPhone(v string) *ListSubAccountResponseBodySubAccountListSubAccount {
	s.Phone = &v
	return s
}

func (s *ListSubAccountResponseBodySubAccountListSubAccount) SetAdminUserKp(v string) *ListSubAccountResponseBodySubAccountListSubAccount {
	s.AdminUserKp = &v
	return s
}

func (s *ListSubAccountResponseBodySubAccountListSubAccount) SetSubUserKp(v string) *ListSubAccountResponseBodySubAccountListSubAccount {
	s.SubUserKp = &v
	return s
}

func (s *ListSubAccountResponseBodySubAccountListSubAccount) SetAdminEdasId(v string) *ListSubAccountResponseBodySubAccountListSubAccount {
	s.AdminEdasId = &v
	return s
}

func (s *ListSubAccountResponseBodySubAccountListSubAccount) SetSubEdasId(v string) *ListSubAccountResponseBodySubAccountListSubAccount {
	s.SubEdasId = &v
	return s
}

type ListSubAccountResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSubAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubAccountResponse) GoString() string {
	return s.String()
}

func (s *ListSubAccountResponse) SetHeaders(v map[string]*string) *ListSubAccountResponse {
	s.Headers = v
	return s
}

func (s *ListSubAccountResponse) SetBody(v *ListSubAccountResponseBody) *ListSubAccountResponse {
	s.Body = v
	return s
}

type ListSlbRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ListSlbRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSlbRequest) GoString() string {
	return s.String()
}

func (s *ListSlbRequest) SetHeaders(v map[string]*string) *ListSlbRequest {
	s.Headers = v
	return s
}

type ListSlbResponseBody struct {
	Code      *int                        `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SlbList   *ListSlbResponseBodySlbList `json:"SlbList,omitempty" xml:"SlbList,omitempty" require:"true" type:"Struct"`
}

func (s ListSlbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSlbResponseBody) GoString() string {
	return s.String()
}

func (s *ListSlbResponseBody) SetCode(v int) *ListSlbResponseBody {
	s.Code = &v
	return s
}

func (s *ListSlbResponseBody) SetMessage(v string) *ListSlbResponseBody {
	s.Message = &v
	return s
}

func (s *ListSlbResponseBody) SetRequestId(v string) *ListSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSlbResponseBody) SetSlbList(v *ListSlbResponseBodySlbList) *ListSlbResponseBody {
	s.SlbList = v
	return s
}

type ListSlbResponseBodySlbList struct {
	SlbEntity []*ListSlbResponseBodySlbListSlbEntity `json:"SlbEntity,omitempty" xml:"SlbEntity,omitempty" require:"true" type:"Repeated"`
}

func (s ListSlbResponseBodySlbList) String() string {
	return tea.Prettify(s)
}

func (s ListSlbResponseBodySlbList) GoString() string {
	return s.String()
}

func (s *ListSlbResponseBodySlbList) SetSlbEntity(v []*ListSlbResponseBodySlbListSlbEntity) *ListSlbResponseBodySlbList {
	s.SlbEntity = v
	return s
}

type ListSlbResponseBodySlbListSlbEntity struct {
	SlbId       *string `json:"SlbId,omitempty" xml:"SlbId,omitempty" require:"true"`
	SlbName     *string `json:"SlbName,omitempty" xml:"SlbName,omitempty" require:"true"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	Address     *string `json:"Address,omitempty" xml:"Address,omitempty" require:"true"`
	SlbStatus   *string `json:"SlbStatus,omitempty" xml:"SlbStatus,omitempty" require:"true"`
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty" require:"true"`
	VswitchId   *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty" require:"true"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" require:"true"`
	GroupId     *int    `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	Expired     *bool   `json:"Expired,omitempty" xml:"Expired,omitempty" require:"true"`
}

func (s ListSlbResponseBodySlbListSlbEntity) String() string {
	return tea.Prettify(s)
}

func (s ListSlbResponseBodySlbListSlbEntity) GoString() string {
	return s.String()
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetSlbId(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.SlbId = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetSlbName(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.SlbName = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetRegionId(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.RegionId = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetUserId(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.UserId = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetAddress(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.Address = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetSlbStatus(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.SlbStatus = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetAddressType(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.AddressType = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetVswitchId(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.VswitchId = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetVpcId(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.VpcId = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetNetworkType(v string) *ListSlbResponseBodySlbListSlbEntity {
	s.NetworkType = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetGroupId(v int) *ListSlbResponseBodySlbListSlbEntity {
	s.GroupId = &v
	return s
}

func (s *ListSlbResponseBodySlbListSlbEntity) SetExpired(v bool) *ListSlbResponseBodySlbListSlbEntity {
	s.Expired = &v
	return s
}

type ListSlbResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSlbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSlbResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSlbResponse) GoString() string {
	return s.String()
}

func (s *ListSlbResponse) SetHeaders(v map[string]*string) *ListSlbResponse {
	s.Headers = v
	return s
}

func (s *ListSlbResponse) SetBody(v *ListSlbResponseBody) *ListSlbResponse {
	s.Body = v
	return s
}

type ListServiceGroupsRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ListServiceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsRequest) SetHeaders(v map[string]*string) *ListServiceGroupsRequest {
	s.Headers = v
	return s
}

type ListServiceGroupsResponseBody struct {
	Code              *int                                            `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message           *string                                         `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId         *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ServiceGroupsList *ListServiceGroupsResponseBodyServiceGroupsList `json:"ServiceGroupsList,omitempty" xml:"ServiceGroupsList,omitempty" require:"true" type:"Struct"`
}

func (s ListServiceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsResponseBody) SetCode(v int) *ListServiceGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListServiceGroupsResponseBody) SetMessage(v string) *ListServiceGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListServiceGroupsResponseBody) SetRequestId(v string) *ListServiceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceGroupsResponseBody) SetServiceGroupsList(v *ListServiceGroupsResponseBodyServiceGroupsList) *ListServiceGroupsResponseBody {
	s.ServiceGroupsList = v
	return s
}

type ListServiceGroupsResponseBodyServiceGroupsList struct {
	ListServiceGroups []*ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups `json:"ListServiceGroups,omitempty" xml:"ListServiceGroups,omitempty" require:"true" type:"Repeated"`
}

func (s ListServiceGroupsResponseBodyServiceGroupsList) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupsResponseBodyServiceGroupsList) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsResponseBodyServiceGroupsList) SetListServiceGroups(v []*ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups) *ListServiceGroupsResponseBodyServiceGroupsList {
	s.ListServiceGroups = v
	return s
}

type ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty" require:"true"`
}

func (s ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups) SetCreateTime(v string) *ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups {
	s.CreateTime = &v
	return s
}

func (s *ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups) SetGroupId(v string) *ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups {
	s.GroupId = &v
	return s
}

func (s *ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups) SetGroupName(v string) *ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups {
	s.GroupName = &v
	return s
}

type ListServiceGroupsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListServiceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServiceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsResponse) SetHeaders(v map[string]*string) *ListServiceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceGroupsResponse) SetBody(v *ListServiceGroupsResponseBody) *ListServiceGroupsResponse {
	s.Body = v
	return s
}

type ListScaleOutEcuQuery struct {
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Cpu             *int    `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Mem             *int    `json:"Mem,omitempty" xml:"Mem,omitempty"`
	InstanceNum     *int    `json:"InstanceNum,omitempty" xml:"InstanceNum,omitempty"`
}

func (s ListScaleOutEcuQuery) String() string {
	return tea.Prettify(s)
}

func (s ListScaleOutEcuQuery) GoString() string {
	return s.String()
}

func (s *ListScaleOutEcuQuery) SetLogicalRegionId(v string) *ListScaleOutEcuQuery {
	s.LogicalRegionId = &v
	return s
}

func (s *ListScaleOutEcuQuery) SetClusterId(v string) *ListScaleOutEcuQuery {
	s.ClusterId = &v
	return s
}

func (s *ListScaleOutEcuQuery) SetAppId(v string) *ListScaleOutEcuQuery {
	s.AppId = &v
	return s
}

func (s *ListScaleOutEcuQuery) SetGroupId(v string) *ListScaleOutEcuQuery {
	s.GroupId = &v
	return s
}

func (s *ListScaleOutEcuQuery) SetCpu(v int) *ListScaleOutEcuQuery {
	s.Cpu = &v
	return s
}

func (s *ListScaleOutEcuQuery) SetMem(v int) *ListScaleOutEcuQuery {
	s.Mem = &v
	return s
}

func (s *ListScaleOutEcuQuery) SetInstanceNum(v int) *ListScaleOutEcuQuery {
	s.InstanceNum = &v
	return s
}

type ListScaleOutEcuRequest struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListScaleOutEcuQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s ListScaleOutEcuRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScaleOutEcuRequest) GoString() string {
	return s.String()
}

func (s *ListScaleOutEcuRequest) SetHeaders(v map[string]*string) *ListScaleOutEcuRequest {
	s.Headers = v
	return s
}

func (s *ListScaleOutEcuRequest) SetQuery(v *ListScaleOutEcuQuery) *ListScaleOutEcuRequest {
	s.Query = v
	return s
}

type ListScaleOutEcuResponseBody struct {
	Code        *int                                    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message     *string                                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	EcuInfoList *ListScaleOutEcuResponseBodyEcuInfoList `json:"EcuInfoList,omitempty" xml:"EcuInfoList,omitempty" require:"true" type:"Struct"`
}

func (s ListScaleOutEcuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScaleOutEcuResponseBody) GoString() string {
	return s.String()
}

func (s *ListScaleOutEcuResponseBody) SetCode(v int) *ListScaleOutEcuResponseBody {
	s.Code = &v
	return s
}

func (s *ListScaleOutEcuResponseBody) SetMessage(v string) *ListScaleOutEcuResponseBody {
	s.Message = &v
	return s
}

func (s *ListScaleOutEcuResponseBody) SetRequestId(v string) *ListScaleOutEcuResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScaleOutEcuResponseBody) SetEcuInfoList(v *ListScaleOutEcuResponseBodyEcuInfoList) *ListScaleOutEcuResponseBody {
	s.EcuInfoList = v
	return s
}

type ListScaleOutEcuResponseBodyEcuInfoList struct {
	EcuInfo []*ListScaleOutEcuResponseBodyEcuInfoListEcuInfo `json:"EcuInfo,omitempty" xml:"EcuInfo,omitempty" require:"true" type:"Repeated"`
}

func (s ListScaleOutEcuResponseBodyEcuInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListScaleOutEcuResponseBodyEcuInfoList) GoString() string {
	return s.String()
}

func (s *ListScaleOutEcuResponseBodyEcuInfoList) SetEcuInfo(v []*ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) *ListScaleOutEcuResponseBodyEcuInfoList {
	s.EcuInfo = v
	return s
}

type ListScaleOutEcuResponseBodyEcuInfoListEcuInfo struct {
	EcuId         *string `json:"EcuId,omitempty" xml:"EcuId,omitempty" require:"true"`
	Online        *bool   `json:"Online,omitempty" xml:"Online,omitempty" require:"true"`
	DockerEnv     *bool   `json:"DockerEnv,omitempty" xml:"DockerEnv,omitempty" require:"true"`
	CreateTime    *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime    *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	IpAddr        *string `json:"IpAddr,omitempty" xml:"IpAddr,omitempty" require:"true"`
	HeartbeatTime *int64  `json:"HeartbeatTime,omitempty" xml:"HeartbeatTime,omitempty" require:"true"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	ZoneId        *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	AvailableCpu  *int    `json:"AvailableCpu,omitempty" xml:"AvailableCpu,omitempty" require:"true"`
	AvailableMem  *int    `json:"AvailableMem,omitempty" xml:"AvailableMem,omitempty" require:"true"`
}

func (s ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) String() string {
	return tea.Prettify(s)
}

func (s ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) GoString() string {
	return s.String()
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetEcuId(v string) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.EcuId = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetOnline(v bool) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.Online = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetDockerEnv(v bool) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.DockerEnv = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetCreateTime(v int64) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.CreateTime = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetUpdateTime(v int64) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.UpdateTime = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetIpAddr(v string) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.IpAddr = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetHeartbeatTime(v int64) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.HeartbeatTime = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetUserId(v string) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.UserId = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetName(v string) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.Name = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetZoneId(v string) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.ZoneId = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetRegionId(v string) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.RegionId = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetInstanceId(v string) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.InstanceId = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetVpcId(v string) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.VpcId = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetAvailableCpu(v int) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.AvailableCpu = &v
	return s
}

func (s *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo) SetAvailableMem(v int) *ListScaleOutEcuResponseBodyEcuInfoListEcuInfo {
	s.AvailableMem = &v
	return s
}

type ListScaleOutEcuResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListScaleOutEcuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListScaleOutEcuResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScaleOutEcuResponse) GoString() string {
	return s.String()
}

func (s *ListScaleOutEcuResponse) SetHeaders(v map[string]*string) *ListScaleOutEcuResponse {
	s.Headers = v
	return s
}

func (s *ListScaleOutEcuResponse) SetBody(v *ListScaleOutEcuResponseBody) *ListScaleOutEcuResponse {
	s.Body = v
	return s
}

type ListRoleRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ListRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRoleRequest) GoString() string {
	return s.String()
}

func (s *ListRoleRequest) SetHeaders(v map[string]*string) *ListRoleRequest {
	s.Headers = v
	return s
}

type ListRoleResponseBody struct {
	Code      *int                          `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RoleList  *ListRoleResponseBodyRoleList `json:"RoleList,omitempty" xml:"RoleList,omitempty" require:"true" type:"Struct"`
}

func (s ListRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRoleResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoleResponseBody) SetCode(v int) *ListRoleResponseBody {
	s.Code = &v
	return s
}

func (s *ListRoleResponseBody) SetMessage(v string) *ListRoleResponseBody {
	s.Message = &v
	return s
}

func (s *ListRoleResponseBody) SetRequestId(v string) *ListRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoleResponseBody) SetRoleList(v *ListRoleResponseBodyRoleList) *ListRoleResponseBody {
	s.RoleList = v
	return s
}

type ListRoleResponseBodyRoleList struct {
	RoleItem []*ListRoleResponseBodyRoleListRoleItem `json:"RoleItem,omitempty" xml:"RoleItem,omitempty" require:"true" type:"Repeated"`
}

func (s ListRoleResponseBodyRoleList) String() string {
	return tea.Prettify(s)
}

func (s ListRoleResponseBodyRoleList) GoString() string {
	return s.String()
}

func (s *ListRoleResponseBodyRoleList) SetRoleItem(v []*ListRoleResponseBodyRoleListRoleItem) *ListRoleResponseBodyRoleList {
	s.RoleItem = v
	return s
}

type ListRoleResponseBodyRoleListRoleItem struct {
	ActionList *ListRoleResponseBodyRoleListRoleItemActionList `json:"ActionList,omitempty" xml:"ActionList,omitempty" require:"true" type:"Struct"`
	Role       *ListRoleResponseBodyRoleListRoleItemRole       `json:"Role,omitempty" xml:"Role,omitempty" require:"true" type:"Struct"`
}

func (s ListRoleResponseBodyRoleListRoleItem) String() string {
	return tea.Prettify(s)
}

func (s ListRoleResponseBodyRoleListRoleItem) GoString() string {
	return s.String()
}

func (s *ListRoleResponseBodyRoleListRoleItem) SetActionList(v *ListRoleResponseBodyRoleListRoleItemActionList) *ListRoleResponseBodyRoleListRoleItem {
	s.ActionList = v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItem) SetRole(v *ListRoleResponseBodyRoleListRoleItemRole) *ListRoleResponseBodyRoleListRoleItem {
	s.Role = v
	return s
}

type ListRoleResponseBodyRoleListRoleItemActionList struct {
	Action []*ListRoleResponseBodyRoleListRoleItemActionListAction `json:"Action,omitempty" xml:"Action,omitempty" require:"true" type:"Repeated"`
}

func (s ListRoleResponseBodyRoleListRoleItemActionList) String() string {
	return tea.Prettify(s)
}

func (s ListRoleResponseBodyRoleListRoleItemActionList) GoString() string {
	return s.String()
}

func (s *ListRoleResponseBodyRoleListRoleItemActionList) SetAction(v []*ListRoleResponseBodyRoleListRoleItemActionListAction) *ListRoleResponseBodyRoleListRoleItemActionList {
	s.Action = v
	return s
}

type ListRoleResponseBodyRoleListRoleItemActionListAction struct {
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
}

func (s ListRoleResponseBodyRoleListRoleItemActionListAction) String() string {
	return tea.Prettify(s)
}

func (s ListRoleResponseBodyRoleListRoleItemActionListAction) GoString() string {
	return s.String()
}

func (s *ListRoleResponseBodyRoleListRoleItemActionListAction) SetGroupId(v string) *ListRoleResponseBodyRoleListRoleItemActionListAction {
	s.GroupId = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemActionListAction) SetCode(v string) *ListRoleResponseBodyRoleListRoleItemActionListAction {
	s.Code = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemActionListAction) SetName(v string) *ListRoleResponseBodyRoleListRoleItemActionListAction {
	s.Name = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemActionListAction) SetDescription(v string) *ListRoleResponseBodyRoleListRoleItemActionListAction {
	s.Description = &v
	return s
}

type ListRoleResponseBodyRoleListRoleItemRole struct {
	Id          *int    `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	AdminUserId *string `json:"AdminUserId,omitempty" xml:"AdminUserId,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	IsDefault   *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty" require:"true"`
}

func (s ListRoleResponseBodyRoleListRoleItemRole) String() string {
	return tea.Prettify(s)
}

func (s ListRoleResponseBodyRoleListRoleItemRole) GoString() string {
	return s.String()
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) SetId(v int) *ListRoleResponseBodyRoleListRoleItemRole {
	s.Id = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) SetAdminUserId(v string) *ListRoleResponseBodyRoleListRoleItemRole {
	s.AdminUserId = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) SetName(v string) *ListRoleResponseBodyRoleListRoleItemRole {
	s.Name = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) SetCreateTime(v int64) *ListRoleResponseBodyRoleListRoleItemRole {
	s.CreateTime = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) SetUpdateTime(v int64) *ListRoleResponseBodyRoleListRoleItemRole {
	s.UpdateTime = &v
	return s
}

func (s *ListRoleResponseBodyRoleListRoleItemRole) SetIsDefault(v bool) *ListRoleResponseBodyRoleListRoleItemRole {
	s.IsDefault = &v
	return s
}

type ListRoleResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRoleResponse) GoString() string {
	return s.String()
}

func (s *ListRoleResponse) SetHeaders(v map[string]*string) *ListRoleResponse {
	s.Headers = v
	return s
}

func (s *ListRoleResponse) SetBody(v *ListRoleResponseBody) *ListRoleResponse {
	s.Body = v
	return s
}

type ListResourceGroupRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ListResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupRequest) SetHeaders(v map[string]*string) *ListResourceGroupRequest {
	s.Headers = v
	return s
}

type ListResourceGroupResponseBody struct {
	Code              *int                                            `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message           *string                                         `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId         *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResourceGroupList *ListResourceGroupResponseBodyResourceGroupList `json:"ResourceGroupList,omitempty" xml:"ResourceGroupList,omitempty" require:"true" type:"Struct"`
}

func (s ListResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBody) SetCode(v int) *ListResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListResourceGroupResponseBody) SetMessage(v string) *ListResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListResourceGroupResponseBody) SetRequestId(v string) *ListResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupResponseBody) SetResourceGroupList(v *ListResourceGroupResponseBodyResourceGroupList) *ListResourceGroupResponseBody {
	s.ResourceGroupList = v
	return s
}

type ListResourceGroupResponseBodyResourceGroupList struct {
	ResGroupEntity []*ListResourceGroupResponseBodyResourceGroupListResGroupEntity `json:"ResGroupEntity,omitempty" xml:"ResGroupEntity,omitempty" require:"true" type:"Repeated"`
}

func (s ListResourceGroupResponseBodyResourceGroupList) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupResponseBodyResourceGroupList) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBodyResourceGroupList) SetResGroupEntity(v []*ListResourceGroupResponseBodyResourceGroupListResGroupEntity) *ListResourceGroupResponseBodyResourceGroupList {
	s.ResGroupEntity = v
	return s
}

type ListResourceGroupResponseBodyResourceGroupListResGroupEntity struct {
	Id          *int64                                                               `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	Name        *string                                                              `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Description *string                                                              `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	AdminUserId *string                                                              `json:"AdminUserId,omitempty" xml:"AdminUserId,omitempty" require:"true"`
	CreateTime  *int64                                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime  *int64                                                               `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	RegionId    *string                                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	EcsList     *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList `json:"ecsList,omitempty" xml:"ecsList,omitempty" require:"true" type:"Struct"`
	SlbList     *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList `json:"SlbList,omitempty" xml:"SlbList,omitempty" require:"true" type:"Struct"`
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntity) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntity) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetId(v int64) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.Id = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetName(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.Name = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetDescription(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.Description = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetAdminUserId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.AdminUserId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetCreateTime(v int64) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.CreateTime = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetUpdateTime(v int64) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.UpdateTime = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetRegionId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.RegionId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetEcsList(v *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.EcsList = v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntity) SetSlbList(v *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList) *ListResourceGroupResponseBodyResourceGroupListResGroupEntity {
	s.SlbList = v
	return s
}

type ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList struct {
	EcsEntity []*ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity `json:"EcsEntity,omitempty" xml:"EcsEntity,omitempty" require:"true" type:"Repeated"`
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList) SetEcsEntity(v []*ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsList {
	s.EcsEntity = v
	return s
}

type ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity struct {
	InstanceId   *string                                                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	InstanceName *string                                                                                `json:"InstanceName,omitempty" xml:"InstanceName,omitempty" require:"true"`
	HostName     *string                                                                                `json:"HostName,omitempty" xml:"HostName,omitempty" require:"true"`
	Description  *string                                                                                `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Status       *string                                                                                `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	PublicIp     *string                                                                                `json:"PublicIp,omitempty" xml:"PublicIp,omitempty" require:"true"`
	InnerIp      *string                                                                                `json:"InnerIp,omitempty" xml:"InnerIp,omitempty" require:"true"`
	PrivateIp    *string                                                                                `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty" require:"true"`
	Eip          *string                                                                                `json:"Eip,omitempty" xml:"Eip,omitempty" require:"true"`
	SerialNum    *string                                                                                `json:"SerialNum,omitempty" xml:"SerialNum,omitempty" require:"true"`
	UserId       *string                                                                                `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	ZoneId       *string                                                                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
	RegionId     *string                                                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Expired      *bool                                                                                  `json:"Expired,omitempty" xml:"Expired,omitempty" require:"true"`
	SgId         *string                                                                                `json:"SgId,omitempty" xml:"SgId,omitempty" require:"true"`
	VpcId        *string                                                                                `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	GroupId      *string                                                                                `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	Cpu          *int                                                                                   `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Mem          *int                                                                                   `json:"Mem,omitempty" xml:"Mem,omitempty" require:"true"`
	EcuEntity    *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity `json:"EcuEntity,omitempty" xml:"EcuEntity,omitempty" require:"true" type:"Struct"`
	VpcEntity    *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity `json:"VpcEntity,omitempty" xml:"VpcEntity,omitempty" require:"true" type:"Struct"`
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetInstanceId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.InstanceId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetInstanceName(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.InstanceName = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetHostName(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.HostName = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetDescription(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.Description = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetStatus(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.Status = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetPublicIp(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.PublicIp = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetInnerIp(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.InnerIp = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetPrivateIp(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.PrivateIp = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetEip(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.Eip = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetSerialNum(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.SerialNum = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetUserId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.UserId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetZoneId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.ZoneId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetRegionId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.RegionId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetExpired(v bool) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.Expired = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetSgId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.SgId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetVpcId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.VpcId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetGroupId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.GroupId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetCpu(v int) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.Cpu = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetMem(v int) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.Mem = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetEcuEntity(v *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.EcuEntity = v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity) SetVpcEntity(v *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntity {
	s.VpcEntity = v
	return s
}

type ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity struct {
	EcuId         *string `json:"EcuId,omitempty" xml:"EcuId,omitempty" require:"true"`
	Online        *bool   `json:"Online,omitempty" xml:"Online,omitempty" require:"true"`
	DockerEnv     *bool   `json:"DockerEnv,omitempty" xml:"DockerEnv,omitempty" require:"true"`
	CreateTime    *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime    *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	IpAddr        *string `json:"IpAddr,omitempty" xml:"IpAddr,omitempty" require:"true"`
	HeartbeatTime *int64  `json:"HeartbeatTime,omitempty" xml:"HeartbeatTime,omitempty" require:"true"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	ZoneId        *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	AvailableCpu  *int    `json:"AvailableCpu,omitempty" xml:"AvailableCpu,omitempty" require:"true"`
	AvailableMem  *int    `json:"AvailableMem,omitempty" xml:"AvailableMem,omitempty" require:"true"`
	Cpu           *int    `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Mem           *int    `json:"Mem,omitempty" xml:"Mem,omitempty" require:"true"`
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetEcuId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.EcuId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetOnline(v bool) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.Online = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetDockerEnv(v bool) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.DockerEnv = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetCreateTime(v int64) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.CreateTime = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetUpdateTime(v int64) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.UpdateTime = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetIpAddr(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.IpAddr = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetHeartbeatTime(v int64) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.HeartbeatTime = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetUserId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.UserId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetName(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.Name = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetZoneId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.ZoneId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetRegionId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.RegionId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetInstanceId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.InstanceId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetVpcId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.VpcId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetAvailableCpu(v int) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.AvailableCpu = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetAvailableMem(v int) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.AvailableMem = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetCpu(v int) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.Cpu = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity) SetMem(v int) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityEcuEntity {
	s.Mem = &v
	return s
}

type ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity struct {
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	VpcName     *string `json:"VpcName,omitempty" xml:"VpcName,omitempty" require:"true"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	Cidrblock   *string `json:"Cidrblock,omitempty" xml:"Cidrblock,omitempty" require:"true"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Expired     *bool   `json:"Expired,omitempty" xml:"Expired,omitempty" require:"true"`
	EcsNum      *int    `json:"EcsNum,omitempty" xml:"EcsNum,omitempty" require:"true"`
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetVpcId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.VpcId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetVpcName(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.VpcName = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetRegionId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.RegionId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetUserId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.UserId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetCidrblock(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.Cidrblock = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetStatus(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.Status = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetDescription(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.Description = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetExpired(v bool) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.Expired = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity) SetEcsNum(v int) *ListResourceGroupResponseBodyResourceGroupListResGroupEntityEcsListEcsEntityVpcEntity {
	s.EcsNum = &v
	return s
}

type ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList struct {
	SlbEntity []*ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity `json:"SlbEntity,omitempty" xml:"SlbEntity,omitempty" require:"true" type:"Repeated"`
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList) SetSlbEntity(v []*ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbList {
	s.SlbEntity = v
	return s
}

type ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity struct {
	SlbId       *string `json:"SlbId,omitempty" xml:"SlbId,omitempty" require:"true"`
	SlbName     *string `json:"SlbName,omitempty" xml:"SlbName,omitempty" require:"true"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	Address     *string `json:"Address,omitempty" xml:"Address,omitempty" require:"true"`
	SlbStatus   *string `json:"SlbStatus,omitempty" xml:"SlbStatus,omitempty" require:"true"`
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty" require:"true"`
	VswitchId   *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty" require:"true"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty" require:"true"`
	GroupId     *int    `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	Expired     *bool   `json:"Expired,omitempty" xml:"Expired,omitempty" require:"true"`
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetSlbId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.SlbId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetSlbName(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.SlbName = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetRegionId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.RegionId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetUserId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.UserId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetAddress(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.Address = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetSlbStatus(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.SlbStatus = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetAddressType(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.AddressType = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetVswitchId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.VswitchId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetVpcId(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.VpcId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetNetworkType(v string) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.NetworkType = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetGroupId(v int) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.GroupId = &v
	return s
}

func (s *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity) SetExpired(v bool) *ListResourceGroupResponseBodyResourceGroupListResGroupEntitySlbListSlbEntity {
	s.Expired = &v
	return s
}

type ListResourceGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponse) SetHeaders(v map[string]*string) *ListResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ListResourceGroupResponse) SetBody(v *ListResourceGroupResponseBody) *ListResourceGroupResponse {
	s.Body = v
	return s
}

type ListRecentChangeOrderQuery struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s ListRecentChangeOrderQuery) String() string {
	return tea.Prettify(s)
}

func (s ListRecentChangeOrderQuery) GoString() string {
	return s.String()
}

func (s *ListRecentChangeOrderQuery) SetAppId(v string) *ListRecentChangeOrderQuery {
	s.AppId = &v
	return s
}

type ListRecentChangeOrderRequest struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListRecentChangeOrderQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ListRecentChangeOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecentChangeOrderRequest) GoString() string {
	return s.String()
}

func (s *ListRecentChangeOrderRequest) SetHeaders(v map[string]*string) *ListRecentChangeOrderRequest {
	s.Headers = v
	return s
}

func (s *ListRecentChangeOrderRequest) SetQuery(v *ListRecentChangeOrderQuery) *ListRecentChangeOrderRequest {
	s.Query = v
	return s
}

type ListRecentChangeOrderResponseBody struct {
	Code            *int                                              `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message         *string                                           `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ChangeOrderList *ListRecentChangeOrderResponseBodyChangeOrderList `json:"ChangeOrderList,omitempty" xml:"ChangeOrderList,omitempty" require:"true" type:"Struct"`
}

func (s ListRecentChangeOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecentChangeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecentChangeOrderResponseBody) SetCode(v int) *ListRecentChangeOrderResponseBody {
	s.Code = &v
	return s
}

func (s *ListRecentChangeOrderResponseBody) SetMessage(v string) *ListRecentChangeOrderResponseBody {
	s.Message = &v
	return s
}

func (s *ListRecentChangeOrderResponseBody) SetRequestId(v string) *ListRecentChangeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecentChangeOrderResponseBody) SetChangeOrderList(v *ListRecentChangeOrderResponseBodyChangeOrderList) *ListRecentChangeOrderResponseBody {
	s.ChangeOrderList = v
	return s
}

type ListRecentChangeOrderResponseBodyChangeOrderList struct {
	ChangeOrder []*ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder `json:"ChangeOrder,omitempty" xml:"ChangeOrder,omitempty" require:"true" type:"Repeated"`
}

func (s ListRecentChangeOrderResponseBodyChangeOrderList) String() string {
	return tea.Prettify(s)
}

func (s ListRecentChangeOrderResponseBodyChangeOrderList) GoString() string {
	return s.String()
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderList) SetChangeOrder(v []*ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) *ListRecentChangeOrderResponseBodyChangeOrderList {
	s.ChangeOrder = v
	return s
}

type ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder struct {
	ChangeOrderId          *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	BatchType              *string `json:"BatchType,omitempty" xml:"BatchType,omitempty" require:"true"`
	BatchCount             *int    `json:"BatchCount,omitempty" xml:"BatchCount,omitempty" require:"true"`
	AppId                  *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	GroupId                *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	Status                 *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	FinishTime             *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty" require:"true"`
	CoType                 *string `json:"CoType,omitempty" xml:"CoType,omitempty" require:"true"`
	CreateUserId           *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty" require:"true"`
	CoTypeCode             *string `json:"CoTypeCode,omitempty" xml:"CoTypeCode,omitempty" require:"true"`
	Source                 *string `json:"Source,omitempty" xml:"Source,omitempty" require:"true"`
	ChangeOrderDescription *string `json:"ChangeOrderDescription,omitempty" xml:"ChangeOrderDescription,omitempty" require:"true"`
}

func (s ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) String() string {
	return tea.Prettify(s)
}

func (s ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) GoString() string {
	return s.String()
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetChangeOrderId(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.ChangeOrderId = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetUserId(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.UserId = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetBatchType(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.BatchType = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetBatchCount(v int) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.BatchCount = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetAppId(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.AppId = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetGroupId(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.GroupId = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetStatus(v int) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.Status = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetCreateTime(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.CreateTime = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetFinishTime(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.FinishTime = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetCoType(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.CoType = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetCreateUserId(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.CreateUserId = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetCoTypeCode(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.CoTypeCode = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetSource(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.Source = &v
	return s
}

func (s *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder) SetChangeOrderDescription(v string) *ListRecentChangeOrderResponseBodyChangeOrderListChangeOrder {
	s.ChangeOrderDescription = &v
	return s
}

type ListRecentChangeOrderResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRecentChangeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRecentChangeOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecentChangeOrderResponse) GoString() string {
	return s.String()
}

func (s *ListRecentChangeOrderResponse) SetHeaders(v map[string]*string) *ListRecentChangeOrderResponse {
	s.Headers = v
	return s
}

func (s *ListRecentChangeOrderResponse) SetBody(v *ListRecentChangeOrderResponseBody) *ListRecentChangeOrderResponse {
	s.Body = v
	return s
}

type ListPublishedServicesQuery struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s ListPublishedServicesQuery) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedServicesQuery) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesQuery) SetAppId(v string) *ListPublishedServicesQuery {
	s.AppId = &v
	return s
}

type ListPublishedServicesRequest struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListPublishedServicesQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ListPublishedServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedServicesRequest) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesRequest) SetHeaders(v map[string]*string) *ListPublishedServicesRequest {
	s.Headers = v
	return s
}

func (s *ListPublishedServicesRequest) SetQuery(v *ListPublishedServicesQuery) *ListPublishedServicesRequest {
	s.Query = v
	return s
}

type ListPublishedServicesResponseBody struct {
	Code                  *int                                                    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message               *string                                                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId             *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PublishedServicesList *ListPublishedServicesResponseBodyPublishedServicesList `json:"PublishedServicesList,omitempty" xml:"PublishedServicesList,omitempty" require:"true" type:"Struct"`
}

func (s ListPublishedServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponseBody) SetCode(v int) *ListPublishedServicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPublishedServicesResponseBody) SetMessage(v string) *ListPublishedServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPublishedServicesResponseBody) SetRequestId(v string) *ListPublishedServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublishedServicesResponseBody) SetPublishedServicesList(v *ListPublishedServicesResponseBodyPublishedServicesList) *ListPublishedServicesResponseBody {
	s.PublishedServicesList = v
	return s
}

type ListPublishedServicesResponseBodyPublishedServicesList struct {
	ListPublishedServices []*ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices `json:"ListPublishedServices,omitempty" xml:"ListPublishedServices,omitempty" require:"true" type:"Repeated"`
}

func (s ListPublishedServicesResponseBodyPublishedServicesList) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedServicesResponseBodyPublishedServicesList) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponseBodyPublishedServicesList) SetListPublishedServices(v []*ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) *ListPublishedServicesResponseBodyPublishedServicesList {
	s.ListPublishedServices = v
	return s
}

type ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices struct {
	AppId             *string                                                                            `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	DockerApplication *bool                                                                              `json:"DockerApplication,omitempty" xml:"DockerApplication,omitempty" require:"true"`
	Group2Ip          *string                                                                            `json:"Group2Ip,omitempty" xml:"Group2Ip,omitempty" require:"true"`
	Name              *string                                                                            `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Type              *string                                                                            `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	Version           *string                                                                            `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Groups            *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups `json:"Groups,omitempty" xml:"Groups,omitempty" require:"true" type:"Struct"`
	Ips               *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps    `json:"Ips,omitempty" xml:"Ips,omitempty" require:"true" type:"Struct"`
}

func (s ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) SetAppId(v string) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	s.AppId = &v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) SetDockerApplication(v bool) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	s.DockerApplication = &v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) SetGroup2Ip(v string) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	s.Group2Ip = &v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) SetName(v string) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	s.Name = &v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) SetType(v string) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	s.Type = &v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) SetVersion(v string) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	s.Version = &v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) SetGroups(v *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	s.Groups = v
	return s
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices) SetIps(v *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServices {
	s.Ips = v
	return s
}

type ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups struct {
	Group []*string `json:"group,omitempty" xml:"group,omitempty" require:"true" type:"Repeated"`
}

func (s ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups) SetGroup(v []*string) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesGroups {
	s.Group = v
	return s
}

type ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps struct {
	Ip []*string `json:"ip,omitempty" xml:"ip,omitempty" require:"true" type:"Repeated"`
}

func (s ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps) SetIp(v []*string) *ListPublishedServicesResponseBodyPublishedServicesListListPublishedServicesIps {
	s.Ip = v
	return s
}

type ListPublishedServicesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPublishedServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPublishedServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedServicesResponse) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponse) SetHeaders(v map[string]*string) *ListPublishedServicesResponse {
	s.Headers = v
	return s
}

func (s *ListPublishedServicesResponse) SetBody(v *ListPublishedServicesResponseBody) *ListPublishedServicesResponse {
	s.Body = v
	return s
}

type ListHistoryDeployVersionQuery struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s ListHistoryDeployVersionQuery) String() string {
	return tea.Prettify(s)
}

func (s ListHistoryDeployVersionQuery) GoString() string {
	return s.String()
}

func (s *ListHistoryDeployVersionQuery) SetAppId(v string) *ListHistoryDeployVersionQuery {
	s.AppId = &v
	return s
}

type ListHistoryDeployVersionRequest struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListHistoryDeployVersionQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ListHistoryDeployVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHistoryDeployVersionRequest) GoString() string {
	return s.String()
}

func (s *ListHistoryDeployVersionRequest) SetHeaders(v map[string]*string) *ListHistoryDeployVersionRequest {
	s.Headers = v
	return s
}

func (s *ListHistoryDeployVersionRequest) SetQuery(v *ListHistoryDeployVersionQuery) *ListHistoryDeployVersionRequest {
	s.Query = v
	return s
}

type ListHistoryDeployVersionResponseBody struct {
	Code               *int                                                    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message            *string                                                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PackageVersionList *ListHistoryDeployVersionResponseBodyPackageVersionList `json:"PackageVersionList,omitempty" xml:"PackageVersionList,omitempty" require:"true" type:"Struct"`
}

func (s ListHistoryDeployVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHistoryDeployVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListHistoryDeployVersionResponseBody) SetCode(v int) *ListHistoryDeployVersionResponseBody {
	s.Code = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBody) SetMessage(v string) *ListHistoryDeployVersionResponseBody {
	s.Message = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBody) SetRequestId(v string) *ListHistoryDeployVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBody) SetPackageVersionList(v *ListHistoryDeployVersionResponseBodyPackageVersionList) *ListHistoryDeployVersionResponseBody {
	s.PackageVersionList = v
	return s
}

type ListHistoryDeployVersionResponseBodyPackageVersionList struct {
	PackageVersion []*ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty" require:"true" type:"Repeated"`
}

func (s ListHistoryDeployVersionResponseBodyPackageVersionList) String() string {
	return tea.Prettify(s)
}

func (s ListHistoryDeployVersionResponseBodyPackageVersionList) GoString() string {
	return s.String()
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionList) SetPackageVersion(v []*ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) *ListHistoryDeployVersionResponseBodyPackageVersionList {
	s.PackageVersion = v
	return s
}

type ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion struct {
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty" require:"true"`
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	WarUrl         *string `json:"WarUrl,omitempty" xml:"WarUrl,omitempty" require:"true"`
	CreateTime     *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime     *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	PublicUrl      *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty" require:"true"`
}

func (s ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) String() string {
	return tea.Prettify(s)
}

func (s ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) GoString() string {
	return s.String()
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetId(v string) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.Id = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetPackageVersion(v string) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.PackageVersion = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetAppId(v string) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.AppId = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetDescription(v string) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.Description = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetWarUrl(v string) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.WarUrl = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetCreateTime(v int64) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.CreateTime = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetUpdateTime(v int64) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.UpdateTime = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetType(v string) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.Type = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetPublicUrl(v string) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.PublicUrl = &v
	return s
}

type ListHistoryDeployVersionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHistoryDeployVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHistoryDeployVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHistoryDeployVersionResponse) GoString() string {
	return s.String()
}

func (s *ListHistoryDeployVersionResponse) SetHeaders(v map[string]*string) *ListHistoryDeployVersionResponse {
	s.Headers = v
	return s
}

func (s *ListHistoryDeployVersionResponse) SetBody(v *ListHistoryDeployVersionResponseBody) *ListHistoryDeployVersionResponse {
	s.Body = v
	return s
}

type ListFlowControlsQuery struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s ListFlowControlsQuery) String() string {
	return tea.Prettify(s)
}

func (s ListFlowControlsQuery) GoString() string {
	return s.String()
}

func (s *ListFlowControlsQuery) SetAppId(v string) *ListFlowControlsQuery {
	s.AppId = &v
	return s
}

type ListFlowControlsRequest struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListFlowControlsQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ListFlowControlsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowControlsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowControlsRequest) SetHeaders(v map[string]*string) *ListFlowControlsRequest {
	s.Headers = v
	return s
}

func (s *ListFlowControlsRequest) SetQuery(v *ListFlowControlsQuery) *ListFlowControlsRequest {
	s.Query = v
	return s
}

type ListFlowControlsResponseBody struct {
	Code            *int                                         `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message         *string                                      `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId       *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	FlowControlsMap *ListFlowControlsResponseBodyFlowControlsMap `json:"FlowControlsMap,omitempty" xml:"FlowControlsMap,omitempty" require:"true" type:"Struct"`
}

func (s ListFlowControlsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowControlsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowControlsResponseBody) SetCode(v int) *ListFlowControlsResponseBody {
	s.Code = &v
	return s
}

func (s *ListFlowControlsResponseBody) SetMessage(v string) *ListFlowControlsResponseBody {
	s.Message = &v
	return s
}

func (s *ListFlowControlsResponseBody) SetRequestId(v string) *ListFlowControlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowControlsResponseBody) SetFlowControlsMap(v *ListFlowControlsResponseBodyFlowControlsMap) *ListFlowControlsResponseBody {
	s.FlowControlsMap = v
	return s
}

type ListFlowControlsResponseBodyFlowControlsMap struct {
	AppId            *string                                                      `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	AppName          *string                                                      `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	InterfaceMethods *ListFlowControlsResponseBodyFlowControlsMapInterfaceMethods `json:"InterfaceMethods,omitempty" xml:"InterfaceMethods,omitempty" require:"true" type:"Struct"`
	AppList          *ListFlowControlsResponseBodyFlowControlsMapAppList          `json:"AppList,omitempty" xml:"AppList,omitempty" require:"true" type:"Struct"`
	RuleList         *ListFlowControlsResponseBodyFlowControlsMapRuleList         `json:"RuleList,omitempty" xml:"RuleList,omitempty" require:"true" type:"Struct"`
}

func (s ListFlowControlsResponseBodyFlowControlsMap) String() string {
	return tea.Prettify(s)
}

func (s ListFlowControlsResponseBodyFlowControlsMap) GoString() string {
	return s.String()
}

func (s *ListFlowControlsResponseBodyFlowControlsMap) SetAppId(v string) *ListFlowControlsResponseBodyFlowControlsMap {
	s.AppId = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMap) SetAppName(v string) *ListFlowControlsResponseBodyFlowControlsMap {
	s.AppName = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMap) SetInterfaceMethods(v *ListFlowControlsResponseBodyFlowControlsMapInterfaceMethods) *ListFlowControlsResponseBodyFlowControlsMap {
	s.InterfaceMethods = v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMap) SetAppList(v *ListFlowControlsResponseBodyFlowControlsMapAppList) *ListFlowControlsResponseBodyFlowControlsMap {
	s.AppList = v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMap) SetRuleList(v *ListFlowControlsResponseBodyFlowControlsMapRuleList) *ListFlowControlsResponseBodyFlowControlsMap {
	s.RuleList = v
	return s
}

type ListFlowControlsResponseBodyFlowControlsMapInterfaceMethods struct {
	InterfaceMethod []*ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod `json:"InterfaceMethod,omitempty" xml:"InterfaceMethod,omitempty" require:"true" type:"Repeated"`
}

func (s ListFlowControlsResponseBodyFlowControlsMapInterfaceMethods) String() string {
	return tea.Prettify(s)
}

func (s ListFlowControlsResponseBodyFlowControlsMapInterfaceMethods) GoString() string {
	return s.String()
}

func (s *ListFlowControlsResponseBodyFlowControlsMapInterfaceMethods) SetInterfaceMethod(v []*ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod) *ListFlowControlsResponseBodyFlowControlsMapInterfaceMethods {
	s.InterfaceMethod = v
	return s
}

type ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod struct {
	Name    *string                                                                            `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	Version *string                                                                            `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Methods *ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethodMethods `json:"Methods,omitempty" xml:"Methods,omitempty" require:"true" type:"Struct"`
}

func (s ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod) String() string {
	return tea.Prettify(s)
}

func (s ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod) GoString() string {
	return s.String()
}

func (s *ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod) SetName(v string) *ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod {
	s.Name = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod) SetVersion(v string) *ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod {
	s.Version = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod) SetMethods(v *ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethodMethods) *ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod {
	s.Methods = v
	return s
}

type ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethodMethods struct {
	Method []*string `json:"Method,omitempty" xml:"Method,omitempty" require:"true" type:"Repeated"`
}

func (s ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethodMethods) String() string {
	return tea.Prettify(s)
}

func (s ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethodMethods) GoString() string {
	return s.String()
}

func (s *ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethodMethods) SetMethod(v []*string) *ListFlowControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethodMethods {
	s.Method = v
	return s
}

type ListFlowControlsResponseBodyFlowControlsMapAppList struct {
	App []*ListFlowControlsResponseBodyFlowControlsMapAppListApp `json:"App,omitempty" xml:"App,omitempty" require:"true" type:"Repeated"`
}

func (s ListFlowControlsResponseBodyFlowControlsMapAppList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowControlsResponseBodyFlowControlsMapAppList) GoString() string {
	return s.String()
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppList) SetApp(v []*ListFlowControlsResponseBodyFlowControlsMapAppListApp) *ListFlowControlsResponseBodyFlowControlsMapAppList {
	s.App = v
	return s
}

type ListFlowControlsResponseBodyFlowControlsMapAppListApp struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Owner                *string `json:"Owner,omitempty" xml:"Owner,omitempty" require:"true"`
	InstanceCount        *int    `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty" require:"true"`
	RunningInstanceCount *int    `json:"RunningInstanceCount,omitempty" xml:"RunningInstanceCount,omitempty" require:"true"`
	Port                 *int    `json:"Port,omitempty" xml:"Port,omitempty" require:"true"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	SlbId                *string `json:"SlbId,omitempty" xml:"SlbId,omitempty" require:"true"`
	SlbIp                *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty" require:"true"`
	SlbPort              *int    `json:"SlbPort,omitempty" xml:"SlbPort,omitempty" require:"true"`
	ExtSlbId             *string `json:"ExtSlbId,omitempty" xml:"ExtSlbId,omitempty" require:"true"`
	ExtSlbIp             *string `json:"ExtSlbIp,omitempty" xml:"ExtSlbIp,omitempty" require:"true"`
	ApplicationType      *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty" require:"true"`
	ClusterType          *int    `json:"ClusterType,omitempty" xml:"ClusterType,omitempty" require:"true"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	Dockerize            *bool   `json:"Dockerize,omitempty" xml:"Dockerize,omitempty" require:"true"`
	Cpu                  *int    `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Memory               *int    `json:"Memory,omitempty" xml:"Memory,omitempty" require:"true"`
	HealthCheckUrl       *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty" require:"true"`
	BuildPackageId       *int64  `json:"BuildPackageId,omitempty" xml:"BuildPackageId,omitempty" require:"true"`
	CreateTime           *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
}

func (s ListFlowControlsResponseBodyFlowControlsMapAppListApp) String() string {
	return tea.Prettify(s)
}

func (s ListFlowControlsResponseBodyFlowControlsMapAppListApp) GoString() string {
	return s.String()
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetAppId(v string) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.AppId = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetName(v string) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.Name = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetRegionId(v string) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.RegionId = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetDescription(v string) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.Description = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetOwner(v string) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.Owner = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetInstanceCount(v int) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.InstanceCount = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetRunningInstanceCount(v int) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.RunningInstanceCount = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetPort(v int) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.Port = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetUserId(v string) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.UserId = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetSlbId(v string) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.SlbId = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetSlbIp(v string) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.SlbIp = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetSlbPort(v int) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.SlbPort = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetExtSlbId(v string) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.ExtSlbId = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetExtSlbIp(v string) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.ExtSlbIp = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetApplicationType(v string) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.ApplicationType = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetClusterType(v int) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.ClusterType = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetClusterId(v string) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.ClusterId = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetDockerize(v bool) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.Dockerize = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetCpu(v int) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.Cpu = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetMemory(v int) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.Memory = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetHealthCheckUrl(v string) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.HealthCheckUrl = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetBuildPackageId(v int64) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.BuildPackageId = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapAppListApp) SetCreateTime(v int64) *ListFlowControlsResponseBodyFlowControlsMapAppListApp {
	s.CreateTime = &v
	return s
}

type ListFlowControlsResponseBodyFlowControlsMapRuleList struct {
	CurrentPage    *int                                                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize       *int                                                               `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
	TotalSize      *int                                                               `json:"TotalSize,omitempty" xml:"TotalSize,omitempty" require:"true"`
	RuleResultList *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultList `json:"RuleResultList,omitempty" xml:"RuleResultList,omitempty" require:"true" type:"Struct"`
}

func (s ListFlowControlsResponseBodyFlowControlsMapRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowControlsResponseBodyFlowControlsMapRuleList) GoString() string {
	return s.String()
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleList) SetCurrentPage(v int) *ListFlowControlsResponseBodyFlowControlsMapRuleList {
	s.CurrentPage = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleList) SetPageSize(v int) *ListFlowControlsResponseBodyFlowControlsMapRuleList {
	s.PageSize = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleList) SetTotalSize(v int) *ListFlowControlsResponseBodyFlowControlsMapRuleList {
	s.TotalSize = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleList) SetRuleResultList(v *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultList) *ListFlowControlsResponseBodyFlowControlsMapRuleList {
	s.RuleResultList = v
	return s
}

type ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultList struct {
	Rule []*ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule `json:"Rule,omitempty" xml:"Rule,omitempty" require:"true" type:"Repeated"`
}

func (s ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultList) GoString() string {
	return s.String()
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultList) SetRule(v []*ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultList {
	s.Rule = v
	return s
}

type ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	ConsumerAppId *string `json:"ConsumerAppId,omitempty" xml:"ConsumerAppId,omitempty" require:"true"`
	CreateTime    *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	Granularity   *string `json:"Granularity,omitempty" xml:"Granularity,omitempty" require:"true"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	Resource      *string `json:"Resource,omitempty" xml:"Resource,omitempty" require:"true"`
	RuleId        *string `json:"RuleId,omitempty" xml:"RuleId,omitempty" require:"true"`
	RuleType      *string `json:"RuleType,omitempty" xml:"RuleType,omitempty" require:"true"`
	State         *int    `json:"State,omitempty" xml:"State,omitempty" require:"true"`
	Strategy      *string `json:"Strategy,omitempty" xml:"Strategy,omitempty" require:"true"`
	Threshold     *int    `json:"Threshold,omitempty" xml:"Threshold,omitempty" require:"true"`
	UpdateTime    *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
}

func (s ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) String() string {
	return tea.Prettify(s)
}

func (s ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) GoString() string {
	return s.String()
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetAppId(v string) *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.AppId = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetConsumerAppId(v string) *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.ConsumerAppId = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetCreateTime(v int64) *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.CreateTime = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetGranularity(v string) *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.Granularity = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetId(v string) *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.Id = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetResource(v string) *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.Resource = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetRuleId(v string) *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.RuleId = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetRuleType(v string) *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.RuleType = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetState(v int) *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.State = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetStrategy(v string) *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.Strategy = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetThreshold(v int) *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.Threshold = &v
	return s
}

func (s *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetUpdateTime(v int64) *ListFlowControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.UpdateTime = &v
	return s
}

type ListFlowControlsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowControlsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowControlsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowControlsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowControlsResponse) SetHeaders(v map[string]*string) *ListFlowControlsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowControlsResponse) SetBody(v *ListFlowControlsResponseBody) *ListFlowControlsResponse {
	s.Body = v
	return s
}

type ListEcuByRegionQuery struct {
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	Act             *string `json:"Act,omitempty" xml:"Act,omitempty" require:"true"`
}

func (s ListEcuByRegionQuery) String() string {
	return tea.Prettify(s)
}

func (s ListEcuByRegionQuery) GoString() string {
	return s.String()
}

func (s *ListEcuByRegionQuery) SetLogicalRegionId(v string) *ListEcuByRegionQuery {
	s.LogicalRegionId = &v
	return s
}

func (s *ListEcuByRegionQuery) SetAct(v string) *ListEcuByRegionQuery {
	s.Act = &v
	return s
}

type ListEcuByRegionRequest struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListEcuByRegionQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ListEcuByRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEcuByRegionRequest) GoString() string {
	return s.String()
}

func (s *ListEcuByRegionRequest) SetHeaders(v map[string]*string) *ListEcuByRegionRequest {
	s.Headers = v
	return s
}

func (s *ListEcuByRegionRequest) SetQuery(v *ListEcuByRegionQuery) *ListEcuByRegionRequest {
	s.Query = v
	return s
}

type ListEcuByRegionResponseBody struct {
	Code          *int                                      `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string                                   `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	EcuEntityList *ListEcuByRegionResponseBodyEcuEntityList `json:"EcuEntityList,omitempty" xml:"EcuEntityList,omitempty" require:"true" type:"Struct"`
}

func (s ListEcuByRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEcuByRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListEcuByRegionResponseBody) SetCode(v int) *ListEcuByRegionResponseBody {
	s.Code = &v
	return s
}

func (s *ListEcuByRegionResponseBody) SetMessage(v string) *ListEcuByRegionResponseBody {
	s.Message = &v
	return s
}

func (s *ListEcuByRegionResponseBody) SetRequestId(v string) *ListEcuByRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEcuByRegionResponseBody) SetEcuEntityList(v *ListEcuByRegionResponseBodyEcuEntityList) *ListEcuByRegionResponseBody {
	s.EcuEntityList = v
	return s
}

type ListEcuByRegionResponseBodyEcuEntityList struct {
	EcuEntity []*ListEcuByRegionResponseBodyEcuEntityListEcuEntity `json:"EcuEntity,omitempty" xml:"EcuEntity,omitempty" require:"true" type:"Repeated"`
}

func (s ListEcuByRegionResponseBodyEcuEntityList) String() string {
	return tea.Prettify(s)
}

func (s ListEcuByRegionResponseBodyEcuEntityList) GoString() string {
	return s.String()
}

func (s *ListEcuByRegionResponseBodyEcuEntityList) SetEcuEntity(v []*ListEcuByRegionResponseBodyEcuEntityListEcuEntity) *ListEcuByRegionResponseBodyEcuEntityList {
	s.EcuEntity = v
	return s
}

type ListEcuByRegionResponseBodyEcuEntityListEcuEntity struct {
	EcuId         *string `json:"EcuId,omitempty" xml:"EcuId,omitempty" require:"true"`
	Online        *bool   `json:"Online,omitempty" xml:"Online,omitempty" require:"true"`
	DockerEnv     *bool   `json:"DockerEnv,omitempty" xml:"DockerEnv,omitempty" require:"true"`
	CreateTime    *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime    *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	IpAddr        *string `json:"IpAddr,omitempty" xml:"IpAddr,omitempty" require:"true"`
	HeartbeatTime *int64  `json:"HeartbeatTime,omitempty" xml:"HeartbeatTime,omitempty" require:"true"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	ZoneId        *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	AvailableCpu  *int    `json:"AvailableCpu,omitempty" xml:"AvailableCpu,omitempty" require:"true"`
	AvailableMem  *int    `json:"AvailableMem,omitempty" xml:"AvailableMem,omitempty" require:"true"`
	Cpu           *int    `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Mem           *int    `json:"Mem,omitempty" xml:"Mem,omitempty" require:"true"`
}

func (s ListEcuByRegionResponseBodyEcuEntityListEcuEntity) String() string {
	return tea.Prettify(s)
}

func (s ListEcuByRegionResponseBodyEcuEntityListEcuEntity) GoString() string {
	return s.String()
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetEcuId(v string) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.EcuId = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetOnline(v bool) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.Online = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetDockerEnv(v bool) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.DockerEnv = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetCreateTime(v int64) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.CreateTime = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetUpdateTime(v int64) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.UpdateTime = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetIpAddr(v string) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.IpAddr = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetHeartbeatTime(v int64) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.HeartbeatTime = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetUserId(v string) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.UserId = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetName(v string) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.Name = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetZoneId(v string) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.ZoneId = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetRegionId(v string) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.RegionId = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetInstanceId(v string) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.InstanceId = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetVpcId(v string) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.VpcId = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetAvailableCpu(v int) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.AvailableCpu = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetAvailableMem(v int) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.AvailableMem = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetCpu(v int) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.Cpu = &v
	return s
}

func (s *ListEcuByRegionResponseBodyEcuEntityListEcuEntity) SetMem(v int) *ListEcuByRegionResponseBodyEcuEntityListEcuEntity {
	s.Mem = &v
	return s
}

type ListEcuByRegionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEcuByRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEcuByRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEcuByRegionResponse) GoString() string {
	return s.String()
}

func (s *ListEcuByRegionResponse) SetHeaders(v map[string]*string) *ListEcuByRegionResponse {
	s.Headers = v
	return s
}

func (s *ListEcuByRegionResponse) SetBody(v *ListEcuByRegionResponseBody) *ListEcuByRegionResponse {
	s.Body = v
	return s
}

type ListDegradeControlsQuery struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s ListDegradeControlsQuery) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeControlsQuery) GoString() string {
	return s.String()
}

func (s *ListDegradeControlsQuery) SetAppId(v string) *ListDegradeControlsQuery {
	s.AppId = &v
	return s
}

type ListDegradeControlsRequest struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListDegradeControlsQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ListDegradeControlsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeControlsRequest) GoString() string {
	return s.String()
}

func (s *ListDegradeControlsRequest) SetHeaders(v map[string]*string) *ListDegradeControlsRequest {
	s.Headers = v
	return s
}

func (s *ListDegradeControlsRequest) SetQuery(v *ListDegradeControlsQuery) *ListDegradeControlsRequest {
	s.Query = v
	return s
}

type ListDegradeControlsResponseBody struct {
	Code            *int                                            `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message         *string                                         `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId       *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	FlowControlsMap *ListDegradeControlsResponseBodyFlowControlsMap `json:"FlowControlsMap,omitempty" xml:"FlowControlsMap,omitempty" require:"true" type:"Struct"`
}

func (s ListDegradeControlsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeControlsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDegradeControlsResponseBody) SetCode(v int) *ListDegradeControlsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDegradeControlsResponseBody) SetMessage(v string) *ListDegradeControlsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDegradeControlsResponseBody) SetRequestId(v string) *ListDegradeControlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDegradeControlsResponseBody) SetFlowControlsMap(v *ListDegradeControlsResponseBodyFlowControlsMap) *ListDegradeControlsResponseBody {
	s.FlowControlsMap = v
	return s
}

type ListDegradeControlsResponseBodyFlowControlsMap struct {
	AppId            *string                                                         `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	AppName          *string                                                         `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	InterfaceMethods *ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethods `json:"InterfaceMethods,omitempty" xml:"InterfaceMethods,omitempty" require:"true" type:"Struct"`
	RuleList         *ListDegradeControlsResponseBodyFlowControlsMapRuleList         `json:"RuleList,omitempty" xml:"RuleList,omitempty" require:"true" type:"Struct"`
}

func (s ListDegradeControlsResponseBodyFlowControlsMap) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeControlsResponseBodyFlowControlsMap) GoString() string {
	return s.String()
}

func (s *ListDegradeControlsResponseBodyFlowControlsMap) SetAppId(v string) *ListDegradeControlsResponseBodyFlowControlsMap {
	s.AppId = &v
	return s
}

func (s *ListDegradeControlsResponseBodyFlowControlsMap) SetAppName(v string) *ListDegradeControlsResponseBodyFlowControlsMap {
	s.AppName = &v
	return s
}

func (s *ListDegradeControlsResponseBodyFlowControlsMap) SetInterfaceMethods(v *ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethods) *ListDegradeControlsResponseBodyFlowControlsMap {
	s.InterfaceMethods = v
	return s
}

func (s *ListDegradeControlsResponseBodyFlowControlsMap) SetRuleList(v *ListDegradeControlsResponseBodyFlowControlsMapRuleList) *ListDegradeControlsResponseBodyFlowControlsMap {
	s.RuleList = v
	return s
}

type ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethods struct {
	InterfaceMethod []*ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod `json:"InterfaceMethod,omitempty" xml:"InterfaceMethod,omitempty" require:"true" type:"Repeated"`
}

func (s ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethods) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethods) GoString() string {
	return s.String()
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethods) SetInterfaceMethod(v []*ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod) *ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethods {
	s.InterfaceMethod = v
	return s
}

type ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod struct {
	Name    *string                                                                               `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	Version *string                                                                               `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Methods *ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethodMethods `json:"Methods,omitempty" xml:"Methods,omitempty" require:"true" type:"Struct"`
}

func (s ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod) GoString() string {
	return s.String()
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod) SetName(v string) *ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod {
	s.Name = &v
	return s
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod) SetVersion(v string) *ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod {
	s.Version = &v
	return s
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod) SetMethods(v *ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethodMethods) *ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethod {
	s.Methods = v
	return s
}

type ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethodMethods struct {
	Method []*string `json:"Method,omitempty" xml:"Method,omitempty" require:"true" type:"Repeated"`
}

func (s ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethodMethods) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethodMethods) GoString() string {
	return s.String()
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethodMethods) SetMethod(v []*string) *ListDegradeControlsResponseBodyFlowControlsMapInterfaceMethodsInterfaceMethodMethods {
	s.Method = v
	return s
}

type ListDegradeControlsResponseBodyFlowControlsMapRuleList struct {
	CurrentPage    *int                                                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize       *int                                                                  `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
	TotalSize      *int                                                                  `json:"TotalSize,omitempty" xml:"TotalSize,omitempty" require:"true"`
	RuleResultList *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultList `json:"RuleResultList,omitempty" xml:"RuleResultList,omitempty" require:"true" type:"Struct"`
}

func (s ListDegradeControlsResponseBodyFlowControlsMapRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeControlsResponseBodyFlowControlsMapRuleList) GoString() string {
	return s.String()
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapRuleList) SetCurrentPage(v int) *ListDegradeControlsResponseBodyFlowControlsMapRuleList {
	s.CurrentPage = &v
	return s
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapRuleList) SetPageSize(v int) *ListDegradeControlsResponseBodyFlowControlsMapRuleList {
	s.PageSize = &v
	return s
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapRuleList) SetTotalSize(v int) *ListDegradeControlsResponseBodyFlowControlsMapRuleList {
	s.TotalSize = &v
	return s
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapRuleList) SetRuleResultList(v *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultList) *ListDegradeControlsResponseBodyFlowControlsMapRuleList {
	s.RuleResultList = v
	return s
}

type ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultList struct {
	Rule []*ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule `json:"Rule,omitempty" xml:"Rule,omitempty" require:"true" type:"Repeated"`
}

func (s ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultList) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultList) GoString() string {
	return s.String()
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultList) SetRule(v []*ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultList {
	s.Rule = v
	return s
}

type ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RuleId      *string `json:"RuleId,omitempty" xml:"RuleId,omitempty" require:"true"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	Resource    *string `json:"Resource,omitempty" xml:"Resource,omitempty" require:"true"`
	RtThreshold *int    `json:"RtThreshold,omitempty" xml:"RtThreshold,omitempty" require:"true"`
	Duration    *int    `json:"Duration,omitempty" xml:"Duration,omitempty" require:"true"`
	State       *int    `json:"State,omitempty" xml:"State,omitempty" require:"true"`
	RuleType    *string `json:"RuleType,omitempty" xml:"RuleType,omitempty" require:"true"`
}

func (s ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) GoString() string {
	return s.String()
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetAppId(v string) *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.AppId = &v
	return s
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetRuleId(v string) *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.RuleId = &v
	return s
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetCreateTime(v int64) *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.CreateTime = &v
	return s
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetUpdateTime(v int64) *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.UpdateTime = &v
	return s
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetResource(v string) *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.Resource = &v
	return s
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetRtThreshold(v int) *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.RtThreshold = &v
	return s
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetDuration(v int) *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.Duration = &v
	return s
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetState(v int) *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.State = &v
	return s
}

func (s *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule) SetRuleType(v string) *ListDegradeControlsResponseBodyFlowControlsMapRuleListRuleResultListRule {
	s.RuleType = &v
	return s
}

type ListDegradeControlsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDegradeControlsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDegradeControlsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDegradeControlsResponse) GoString() string {
	return s.String()
}

func (s *ListDegradeControlsResponse) SetHeaders(v map[string]*string) *ListDegradeControlsResponse {
	s.Headers = v
	return s
}

func (s *ListDegradeControlsResponse) SetBody(v *ListDegradeControlsResponseBody) *ListDegradeControlsResponse {
	s.Body = v
	return s
}

type ListConsumedServicesQuery struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s ListConsumedServicesQuery) String() string {
	return tea.Prettify(s)
}

func (s ListConsumedServicesQuery) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesQuery) SetAppId(v string) *ListConsumedServicesQuery {
	s.AppId = &v
	return s
}

type ListConsumedServicesRequest struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListConsumedServicesQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ListConsumedServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConsumedServicesRequest) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesRequest) SetHeaders(v map[string]*string) *ListConsumedServicesRequest {
	s.Headers = v
	return s
}

func (s *ListConsumedServicesRequest) SetQuery(v *ListConsumedServicesQuery) *ListConsumedServicesRequest {
	s.Query = v
	return s
}

type ListConsumedServicesResponseBody struct {
	Code                 *int                                                  `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message              *string                                               `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId            *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ConsumedServicesList *ListConsumedServicesResponseBodyConsumedServicesList `json:"ConsumedServicesList,omitempty" xml:"ConsumedServicesList,omitempty" require:"true" type:"Struct"`
}

func (s ListConsumedServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConsumedServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponseBody) SetCode(v int) *ListConsumedServicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumedServicesResponseBody) SetMessage(v string) *ListConsumedServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumedServicesResponseBody) SetRequestId(v string) *ListConsumedServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumedServicesResponseBody) SetConsumedServicesList(v *ListConsumedServicesResponseBodyConsumedServicesList) *ListConsumedServicesResponseBody {
	s.ConsumedServicesList = v
	return s
}

type ListConsumedServicesResponseBodyConsumedServicesList struct {
	ListConsumedServices []*ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices `json:"ListConsumedServices,omitempty" xml:"ListConsumedServices,omitempty" require:"true" type:"Repeated"`
}

func (s ListConsumedServicesResponseBodyConsumedServicesList) String() string {
	return tea.Prettify(s)
}

func (s ListConsumedServicesResponseBodyConsumedServicesList) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponseBodyConsumedServicesList) SetListConsumedServices(v []*ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) *ListConsumedServicesResponseBodyConsumedServicesList {
	s.ListConsumedServices = v
	return s
}

type ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices struct {
	AppId             *string                                                                         `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	DockerApplication *bool                                                                           `json:"DockerApplication,omitempty" xml:"DockerApplication,omitempty" require:"true"`
	Group2Ip          *string                                                                         `json:"Group2Ip,omitempty" xml:"Group2Ip,omitempty" require:"true"`
	Name              *string                                                                         `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Type              *string                                                                         `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	Version           *string                                                                         `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
	Groups            *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups `json:"Groups,omitempty" xml:"Groups,omitempty" require:"true" type:"Struct"`
	Ips               *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps    `json:"Ips,omitempty" xml:"Ips,omitempty" require:"true" type:"Struct"`
}

func (s ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) String() string {
	return tea.Prettify(s)
}

func (s ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) SetAppId(v string) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	s.AppId = &v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) SetDockerApplication(v bool) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	s.DockerApplication = &v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) SetGroup2Ip(v string) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	s.Group2Ip = &v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) SetName(v string) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	s.Name = &v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) SetType(v string) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	s.Type = &v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) SetVersion(v string) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	s.Version = &v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) SetGroups(v *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	s.Groups = v
	return s
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices) SetIps(v *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServices {
	s.Ips = v
	return s
}

type ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups struct {
	Group []*string `json:"group,omitempty" xml:"group,omitempty" require:"true" type:"Repeated"`
}

func (s ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups) String() string {
	return tea.Prettify(s)
}

func (s ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups) SetGroup(v []*string) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesGroups {
	s.Group = v
	return s
}

type ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps struct {
	Ip []*string `json:"ip,omitempty" xml:"ip,omitempty" require:"true" type:"Repeated"`
}

func (s ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps) String() string {
	return tea.Prettify(s)
}

func (s ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps) SetIp(v []*string) *ListConsumedServicesResponseBodyConsumedServicesListListConsumedServicesIps {
	s.Ip = v
	return s
}

type ListConsumedServicesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConsumedServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConsumedServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConsumedServicesResponse) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponse) SetHeaders(v map[string]*string) *ListConsumedServicesResponse {
	s.Headers = v
	return s
}

func (s *ListConsumedServicesResponse) SetBody(v *ListConsumedServicesResponseBody) *ListConsumedServicesResponse {
	s.Body = v
	return s
}

type ListConfigCentersQuery struct {
	DataIdPattern   *string `json:"DataIdPattern,omitempty" xml:"DataIdPattern,omitempty"`
	Group           *string `json:"Group,omitempty" xml:"Group,omitempty" require:"true"`
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty" require:"true"`
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s ListConfigCentersQuery) String() string {
	return tea.Prettify(s)
}

func (s ListConfigCentersQuery) GoString() string {
	return s.String()
}

func (s *ListConfigCentersQuery) SetDataIdPattern(v string) *ListConfigCentersQuery {
	s.DataIdPattern = &v
	return s
}

func (s *ListConfigCentersQuery) SetGroup(v string) *ListConfigCentersQuery {
	s.Group = &v
	return s
}

func (s *ListConfigCentersQuery) SetLogicalRegionId(v string) *ListConfigCentersQuery {
	s.LogicalRegionId = &v
	return s
}

func (s *ListConfigCentersQuery) SetAppName(v string) *ListConfigCentersQuery {
	s.AppName = &v
	return s
}

type ListConfigCentersRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListConfigCentersQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ListConfigCentersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConfigCentersRequest) GoString() string {
	return s.String()
}

func (s *ListConfigCentersRequest) SetHeaders(v map[string]*string) *ListConfigCentersRequest {
	s.Headers = v
	return s
}

func (s *ListConfigCentersRequest) SetQuery(v *ListConfigCentersQuery) *ListConfigCentersRequest {
	s.Query = v
	return s
}

type ListConfigCentersResponseBody struct {
	Code              *int                                            `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message           *string                                         `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId         *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ConfigCentersList *ListConfigCentersResponseBodyConfigCentersList `json:"ConfigCentersList,omitempty" xml:"ConfigCentersList,omitempty" require:"true" type:"Struct"`
}

func (s ListConfigCentersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConfigCentersResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigCentersResponseBody) SetCode(v int) *ListConfigCentersResponseBody {
	s.Code = &v
	return s
}

func (s *ListConfigCentersResponseBody) SetMessage(v string) *ListConfigCentersResponseBody {
	s.Message = &v
	return s
}

func (s *ListConfigCentersResponseBody) SetRequestId(v string) *ListConfigCentersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConfigCentersResponseBody) SetConfigCentersList(v *ListConfigCentersResponseBodyConfigCentersList) *ListConfigCentersResponseBody {
	s.ConfigCentersList = v
	return s
}

type ListConfigCentersResponseBodyConfigCentersList struct {
	ListConfigCenters []*ListConfigCentersResponseBodyConfigCentersListListConfigCenters `json:"ListConfigCenters,omitempty" xml:"ListConfigCenters,omitempty" require:"true" type:"Repeated"`
}

func (s ListConfigCentersResponseBodyConfigCentersList) String() string {
	return tea.Prettify(s)
}

func (s ListConfigCentersResponseBodyConfigCentersList) GoString() string {
	return s.String()
}

func (s *ListConfigCentersResponseBodyConfigCentersList) SetListConfigCenters(v []*ListConfigCentersResponseBodyConfigCentersListListConfigCenters) *ListConfigCentersResponseBodyConfigCentersList {
	s.ListConfigCenters = v
	return s
}

type ListConfigCentersResponseBodyConfigCentersListListConfigCenters struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	DataId  *string `json:"DataId,omitempty" xml:"DataId,omitempty" require:"true"`
	Group   *string `json:"Group,omitempty" xml:"Group,omitempty" require:"true"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
}

func (s ListConfigCentersResponseBodyConfigCentersListListConfigCenters) String() string {
	return tea.Prettify(s)
}

func (s ListConfigCentersResponseBodyConfigCentersListListConfigCenters) GoString() string {
	return s.String()
}

func (s *ListConfigCentersResponseBodyConfigCentersListListConfigCenters) SetAppName(v string) *ListConfigCentersResponseBodyConfigCentersListListConfigCenters {
	s.AppName = &v
	return s
}

func (s *ListConfigCentersResponseBodyConfigCentersListListConfigCenters) SetDataId(v string) *ListConfigCentersResponseBodyConfigCentersListListConfigCenters {
	s.DataId = &v
	return s
}

func (s *ListConfigCentersResponseBodyConfigCentersListListConfigCenters) SetGroup(v string) *ListConfigCentersResponseBodyConfigCentersListListConfigCenters {
	s.Group = &v
	return s
}

func (s *ListConfigCentersResponseBodyConfigCentersListListConfigCenters) SetId(v string) *ListConfigCentersResponseBodyConfigCentersListListConfigCenters {
	s.Id = &v
	return s
}

type ListConfigCentersResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConfigCentersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConfigCentersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConfigCentersResponse) GoString() string {
	return s.String()
}

func (s *ListConfigCentersResponse) SetHeaders(v map[string]*string) *ListConfigCentersResponse {
	s.Headers = v
	return s
}

func (s *ListConfigCentersResponse) SetBody(v *ListConfigCentersResponseBody) *ListConfigCentersResponse {
	s.Body = v
	return s
}

type ListClusterMembersQuery struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	CurrentPage *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	EcsList     *string `json:"EcsList,omitempty" xml:"EcsList,omitempty"`
}

func (s ListClusterMembersQuery) String() string {
	return tea.Prettify(s)
}

func (s ListClusterMembersQuery) GoString() string {
	return s.String()
}

func (s *ListClusterMembersQuery) SetClusterId(v string) *ListClusterMembersQuery {
	s.ClusterId = &v
	return s
}

func (s *ListClusterMembersQuery) SetCurrentPage(v int) *ListClusterMembersQuery {
	s.CurrentPage = &v
	return s
}

func (s *ListClusterMembersQuery) SetPageSize(v int) *ListClusterMembersQuery {
	s.PageSize = &v
	return s
}

func (s *ListClusterMembersQuery) SetEcsList(v string) *ListClusterMembersQuery {
	s.EcsList = &v
	return s
}

type ListClusterMembersRequest struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListClusterMembersQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ListClusterMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterMembersRequest) GoString() string {
	return s.String()
}

func (s *ListClusterMembersRequest) SetHeaders(v map[string]*string) *ListClusterMembersRequest {
	s.Headers = v
	return s
}

func (s *ListClusterMembersRequest) SetQuery(v *ListClusterMembersQuery) *ListClusterMembersRequest {
	s.Query = v
	return s
}

type ListClusterMembersResponseBody struct {
	Code              *int                                             `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message           *string                                          `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId         *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ClusterMemberPage *ListClusterMembersResponseBodyClusterMemberPage `json:"ClusterMemberPage,omitempty" xml:"ClusterMemberPage,omitempty" require:"true" type:"Struct"`
}

func (s ListClusterMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterMembersResponseBody) SetCode(v int) *ListClusterMembersResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterMembersResponseBody) SetMessage(v string) *ListClusterMembersResponseBody {
	s.Message = &v
	return s
}

func (s *ListClusterMembersResponseBody) SetRequestId(v string) *ListClusterMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterMembersResponseBody) SetClusterMemberPage(v *ListClusterMembersResponseBodyClusterMemberPage) *ListClusterMembersResponseBody {
	s.ClusterMemberPage = v
	return s
}

type ListClusterMembersResponseBodyClusterMemberPage struct {
	CurrentPage       *int                                                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize          *int                                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalSize         *int                                                              `json:"TotalSize,omitempty" xml:"TotalSize,omitempty" require:"true"`
	ClusterMemberList *ListClusterMembersResponseBodyClusterMemberPageClusterMemberList `json:"ClusterMemberList,omitempty" xml:"ClusterMemberList,omitempty" require:"true" type:"Struct"`
}

func (s ListClusterMembersResponseBodyClusterMemberPage) String() string {
	return tea.Prettify(s)
}

func (s ListClusterMembersResponseBodyClusterMemberPage) GoString() string {
	return s.String()
}

func (s *ListClusterMembersResponseBodyClusterMemberPage) SetCurrentPage(v int) *ListClusterMembersResponseBodyClusterMemberPage {
	s.CurrentPage = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPage) SetPageSize(v int) *ListClusterMembersResponseBodyClusterMemberPage {
	s.PageSize = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPage) SetTotalSize(v int) *ListClusterMembersResponseBodyClusterMemberPage {
	s.TotalSize = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPage) SetClusterMemberList(v *ListClusterMembersResponseBodyClusterMemberPageClusterMemberList) *ListClusterMembersResponseBodyClusterMemberPage {
	s.ClusterMemberList = v
	return s
}

type ListClusterMembersResponseBodyClusterMemberPageClusterMemberList struct {
	ClusterMember []*ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember `json:"ClusterMember,omitempty" xml:"ClusterMember,omitempty" require:"true" type:"Repeated"`
}

func (s ListClusterMembersResponseBodyClusterMemberPageClusterMemberList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterMembersResponseBodyClusterMemberPageClusterMemberList) GoString() string {
	return s.String()
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberList) SetClusterMember(v []*ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberList {
	s.ClusterMember = v
	return s
}

type ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember struct {
	ClusterMemberId *string `json:"ClusterMemberId,omitempty" xml:"ClusterMemberId,omitempty" require:"true"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	EcuId           *string `json:"EcuId,omitempty" xml:"EcuId,omitempty" require:"true"`
	EcsId           *string `json:"EcsId,omitempty" xml:"EcsId,omitempty" require:"true"`
	Status          *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	CreateTime      *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime      *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
}

func (s ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) String() string {
	return tea.Prettify(s)
}

func (s ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) GoString() string {
	return s.String()
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) SetClusterMemberId(v string) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember {
	s.ClusterMemberId = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) SetClusterId(v string) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember {
	s.ClusterId = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) SetEcuId(v string) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember {
	s.EcuId = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) SetEcsId(v string) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember {
	s.EcsId = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) SetStatus(v int) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember {
	s.Status = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) SetCreateTime(v int64) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember {
	s.CreateTime = &v
	return s
}

func (s *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember) SetUpdateTime(v int64) *ListClusterMembersResponseBodyClusterMemberPageClusterMemberListClusterMember {
	s.UpdateTime = &v
	return s
}

type ListClusterMembersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterMembersResponse) GoString() string {
	return s.String()
}

func (s *ListClusterMembersResponse) SetHeaders(v map[string]*string) *ListClusterMembersResponse {
	s.Headers = v
	return s
}

func (s *ListClusterMembersResponse) SetBody(v *ListClusterMembersResponseBody) *ListClusterMembersResponse {
	s.Body = v
	return s
}

type ListAuthorityRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ListAuthorityRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorityRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorityRequest) SetHeaders(v map[string]*string) *ListAuthorityRequest {
	s.Headers = v
	return s
}

type ListAuthorityResponseBody struct {
	Code          *int                                    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string                                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AuthorityList *ListAuthorityResponseBodyAuthorityList `json:"AuthorityList,omitempty" xml:"AuthorityList,omitempty" require:"true" type:"Struct"`
}

func (s ListAuthorityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorityResponseBody) SetCode(v int) *ListAuthorityResponseBody {
	s.Code = &v
	return s
}

func (s *ListAuthorityResponseBody) SetMessage(v string) *ListAuthorityResponseBody {
	s.Message = &v
	return s
}

func (s *ListAuthorityResponseBody) SetRequestId(v string) *ListAuthorityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorityResponseBody) SetAuthorityList(v *ListAuthorityResponseBodyAuthorityList) *ListAuthorityResponseBody {
	s.AuthorityList = v
	return s
}

type ListAuthorityResponseBodyAuthorityList struct {
	Authority []*ListAuthorityResponseBodyAuthorityListAuthority `json:"Authority,omitempty" xml:"Authority,omitempty" require:"true" type:"Repeated"`
}

func (s ListAuthorityResponseBodyAuthorityList) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorityResponseBodyAuthorityList) GoString() string {
	return s.String()
}

func (s *ListAuthorityResponseBodyAuthorityList) SetAuthority(v []*ListAuthorityResponseBodyAuthorityListAuthority) *ListAuthorityResponseBodyAuthorityList {
	s.Authority = v
	return s
}

type ListAuthorityResponseBodyAuthorityListAuthority struct {
	GroupId     *string                                                    `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	Name        *string                                                    `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Description *string                                                    `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	ActionList  *ListAuthorityResponseBodyAuthorityListAuthorityActionList `json:"ActionList,omitempty" xml:"ActionList,omitempty" require:"true" type:"Struct"`
}

func (s ListAuthorityResponseBodyAuthorityListAuthority) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorityResponseBodyAuthorityListAuthority) GoString() string {
	return s.String()
}

func (s *ListAuthorityResponseBodyAuthorityListAuthority) SetGroupId(v string) *ListAuthorityResponseBodyAuthorityListAuthority {
	s.GroupId = &v
	return s
}

func (s *ListAuthorityResponseBodyAuthorityListAuthority) SetName(v string) *ListAuthorityResponseBodyAuthorityListAuthority {
	s.Name = &v
	return s
}

func (s *ListAuthorityResponseBodyAuthorityListAuthority) SetDescription(v string) *ListAuthorityResponseBodyAuthorityListAuthority {
	s.Description = &v
	return s
}

func (s *ListAuthorityResponseBodyAuthorityListAuthority) SetActionList(v *ListAuthorityResponseBodyAuthorityListAuthorityActionList) *ListAuthorityResponseBodyAuthorityListAuthority {
	s.ActionList = v
	return s
}

type ListAuthorityResponseBodyAuthorityListAuthorityActionList struct {
	Action []*ListAuthorityResponseBodyAuthorityListAuthorityActionListAction `json:"Action,omitempty" xml:"Action,omitempty" require:"true" type:"Repeated"`
}

func (s ListAuthorityResponseBodyAuthorityListAuthorityActionList) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorityResponseBodyAuthorityListAuthorityActionList) GoString() string {
	return s.String()
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionList) SetAction(v []*ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) *ListAuthorityResponseBodyAuthorityListAuthorityActionList {
	s.Action = v
	return s
}

type ListAuthorityResponseBodyAuthorityListAuthorityActionListAction struct {
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
}

func (s ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) GoString() string {
	return s.String()
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) SetGroupId(v string) *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction {
	s.GroupId = &v
	return s
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) SetCode(v string) *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction {
	s.Code = &v
	return s
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) SetName(v string) *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction {
	s.Name = &v
	return s
}

func (s *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction) SetDescription(v string) *ListAuthorityResponseBodyAuthorityListAuthorityActionListAction {
	s.Description = &v
	return s
}

type ListAuthorityResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAuthorityResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorityResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorityResponse) SetHeaders(v map[string]*string) *ListAuthorityResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorityResponse) SetBody(v *ListAuthorityResponseBody) *ListAuthorityResponse {
	s.Body = v
	return s
}

type ListAliyunRegionRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ListAliyunRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAliyunRegionRequest) GoString() string {
	return s.String()
}

func (s *ListAliyunRegionRequest) SetHeaders(v map[string]*string) *ListAliyunRegionRequest {
	s.Headers = v
	return s
}

type ListAliyunRegionResponseBody struct {
	Code             *int                                          `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message          *string                                       `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId        *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionEntityList *ListAliyunRegionResponseBodyRegionEntityList `json:"RegionEntityList,omitempty" xml:"RegionEntityList,omitempty" require:"true" type:"Struct"`
}

func (s ListAliyunRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAliyunRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListAliyunRegionResponseBody) SetCode(v int) *ListAliyunRegionResponseBody {
	s.Code = &v
	return s
}

func (s *ListAliyunRegionResponseBody) SetMessage(v string) *ListAliyunRegionResponseBody {
	s.Message = &v
	return s
}

func (s *ListAliyunRegionResponseBody) SetRequestId(v string) *ListAliyunRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAliyunRegionResponseBody) SetRegionEntityList(v *ListAliyunRegionResponseBodyRegionEntityList) *ListAliyunRegionResponseBody {
	s.RegionEntityList = v
	return s
}

type ListAliyunRegionResponseBodyRegionEntityList struct {
	RegionEntity []*ListAliyunRegionResponseBodyRegionEntityListRegionEntity `json:"RegionEntity,omitempty" xml:"RegionEntity,omitempty" require:"true" type:"Repeated"`
}

func (s ListAliyunRegionResponseBodyRegionEntityList) String() string {
	return tea.Prettify(s)
}

func (s ListAliyunRegionResponseBodyRegionEntityList) GoString() string {
	return s.String()
}

func (s *ListAliyunRegionResponseBodyRegionEntityList) SetRegionEntity(v []*ListAliyunRegionResponseBodyRegionEntityListRegionEntity) *ListAliyunRegionResponseBodyRegionEntityList {
	s.RegionEntity = v
	return s
}

type ListAliyunRegionResponseBodyRegionEntityListRegionEntity struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s ListAliyunRegionResponseBodyRegionEntityListRegionEntity) String() string {
	return tea.Prettify(s)
}

func (s ListAliyunRegionResponseBodyRegionEntityListRegionEntity) GoString() string {
	return s.String()
}

func (s *ListAliyunRegionResponseBodyRegionEntityListRegionEntity) SetId(v string) *ListAliyunRegionResponseBodyRegionEntityListRegionEntity {
	s.Id = &v
	return s
}

func (s *ListAliyunRegionResponseBodyRegionEntityListRegionEntity) SetName(v string) *ListAliyunRegionResponseBodyRegionEntityListRegionEntity {
	s.Name = &v
	return s
}

type ListAliyunRegionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAliyunRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAliyunRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAliyunRegionResponse) GoString() string {
	return s.String()
}

func (s *ListAliyunRegionResponse) SetHeaders(v map[string]*string) *ListAliyunRegionResponse {
	s.Headers = v
	return s
}

func (s *ListAliyunRegionResponse) SetBody(v *ListAliyunRegionResponseBody) *ListAliyunRegionResponse {
	s.Body = v
	return s
}

type InsertServiceGroupQuery struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty" require:"true"`
}

func (s InsertServiceGroupQuery) String() string {
	return tea.Prettify(s)
}

func (s InsertServiceGroupQuery) GoString() string {
	return s.String()
}

func (s *InsertServiceGroupQuery) SetGroupName(v string) *InsertServiceGroupQuery {
	s.GroupName = &v
	return s
}

type InsertServiceGroupRequest struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *InsertServiceGroupQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s InsertServiceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *InsertServiceGroupRequest) SetHeaders(v map[string]*string) *InsertServiceGroupRequest {
	s.Headers = v
	return s
}

func (s *InsertServiceGroupRequest) SetQuery(v *InsertServiceGroupQuery) *InsertServiceGroupRequest {
	s.Query = v
	return s
}

type InsertServiceGroupResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s InsertServiceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertServiceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *InsertServiceGroupResponseBody) SetCode(v int) *InsertServiceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *InsertServiceGroupResponseBody) SetMessage(v string) *InsertServiceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *InsertServiceGroupResponseBody) SetRequestId(v string) *InsertServiceGroupResponseBody {
	s.RequestId = &v
	return s
}

type InsertServiceGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertServiceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertServiceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *InsertServiceGroupResponse) SetHeaders(v map[string]*string) *InsertServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *InsertServiceGroupResponse) SetBody(v *InsertServiceGroupResponseBody) *InsertServiceGroupResponse {
	s.Body = v
	return s
}

type InsertRoleQuery struct {
	RoleName   *string `json:"RoleName,omitempty" xml:"RoleName,omitempty" require:"true"`
	ActionData *string `json:"ActionData,omitempty" xml:"ActionData,omitempty" require:"true"`
}

func (s InsertRoleQuery) String() string {
	return tea.Prettify(s)
}

func (s InsertRoleQuery) GoString() string {
	return s.String()
}

func (s *InsertRoleQuery) SetRoleName(v string) *InsertRoleQuery {
	s.RoleName = &v
	return s
}

func (s *InsertRoleQuery) SetActionData(v string) *InsertRoleQuery {
	s.ActionData = &v
	return s
}

type InsertRoleRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *InsertRoleQuery   `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s InsertRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertRoleRequest) GoString() string {
	return s.String()
}

func (s *InsertRoleRequest) SetHeaders(v map[string]*string) *InsertRoleRequest {
	s.Headers = v
	return s
}

func (s *InsertRoleRequest) SetQuery(v *InsertRoleQuery) *InsertRoleRequest {
	s.Query = v
	return s
}

type InsertRoleResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RoleId    *int    `json:"RoleId,omitempty" xml:"RoleId,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s InsertRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertRoleResponseBody) GoString() string {
	return s.String()
}

func (s *InsertRoleResponseBody) SetCode(v int) *InsertRoleResponseBody {
	s.Code = &v
	return s
}

func (s *InsertRoleResponseBody) SetMessage(v string) *InsertRoleResponseBody {
	s.Message = &v
	return s
}

func (s *InsertRoleResponseBody) SetRoleId(v int) *InsertRoleResponseBody {
	s.RoleId = &v
	return s
}

func (s *InsertRoleResponseBody) SetRequestId(v string) *InsertRoleResponseBody {
	s.RequestId = &v
	return s
}

type InsertRoleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertRoleResponse) GoString() string {
	return s.String()
}

func (s *InsertRoleResponse) SetHeaders(v map[string]*string) *InsertRoleResponse {
	s.Headers = v
	return s
}

func (s *InsertRoleResponse) SetBody(v *InsertRoleResponseBody) *InsertRoleResponse {
	s.Body = v
	return s
}

type InsertOrUpdateRegionQuery struct {
	RegionTag   *string `json:"RegionTag,omitempty" xml:"RegionTag,omitempty" require:"true"`
	RegionName  *string `json:"RegionName,omitempty" xml:"RegionName,omitempty" require:"true"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	DebugEnable *bool   `json:"DebugEnable,omitempty" xml:"DebugEnable,omitempty"`
}

func (s InsertOrUpdateRegionQuery) String() string {
	return tea.Prettify(s)
}

func (s InsertOrUpdateRegionQuery) GoString() string {
	return s.String()
}

func (s *InsertOrUpdateRegionQuery) SetRegionTag(v string) *InsertOrUpdateRegionQuery {
	s.RegionTag = &v
	return s
}

func (s *InsertOrUpdateRegionQuery) SetRegionName(v string) *InsertOrUpdateRegionQuery {
	s.RegionName = &v
	return s
}

func (s *InsertOrUpdateRegionQuery) SetDescription(v string) *InsertOrUpdateRegionQuery {
	s.Description = &v
	return s
}

func (s *InsertOrUpdateRegionQuery) SetId(v int64) *InsertOrUpdateRegionQuery {
	s.Id = &v
	return s
}

func (s *InsertOrUpdateRegionQuery) SetDebugEnable(v bool) *InsertOrUpdateRegionQuery {
	s.DebugEnable = &v
	return s
}

type InsertOrUpdateRegionRequest struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *InsertOrUpdateRegionQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s InsertOrUpdateRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertOrUpdateRegionRequest) GoString() string {
	return s.String()
}

func (s *InsertOrUpdateRegionRequest) SetHeaders(v map[string]*string) *InsertOrUpdateRegionRequest {
	s.Headers = v
	return s
}

func (s *InsertOrUpdateRegionRequest) SetQuery(v *InsertOrUpdateRegionQuery) *InsertOrUpdateRegionRequest {
	s.Query = v
	return s
}

type InsertOrUpdateRegionResponseBody struct {
	Code                   *int                                                    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message                *string                                                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId              *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UserDefineRegionEntity *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity `json:"UserDefineRegionEntity,omitempty" xml:"UserDefineRegionEntity,omitempty" require:"true" type:"Struct"`
}

func (s InsertOrUpdateRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertOrUpdateRegionResponseBody) GoString() string {
	return s.String()
}

func (s *InsertOrUpdateRegionResponseBody) SetCode(v int) *InsertOrUpdateRegionResponseBody {
	s.Code = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBody) SetMessage(v string) *InsertOrUpdateRegionResponseBody {
	s.Message = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBody) SetRequestId(v string) *InsertOrUpdateRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBody) SetUserDefineRegionEntity(v *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) *InsertOrUpdateRegionResponseBody {
	s.UserDefineRegionEntity = v
	return s
}

type InsertOrUpdateRegionResponseBodyUserDefineRegionEntity struct {
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	RegionName   *string `json:"RegionName,omitempty" xml:"RegionName,omitempty" require:"true"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	BelongRegion *string `json:"BelongRegion,omitempty" xml:"BelongRegion,omitempty" require:"true"`
	DebugEnable  *bool   `json:"DebugEnable,omitempty" xml:"DebugEnable,omitempty" require:"true"`
}

func (s InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) String() string {
	return tea.Prettify(s)
}

func (s InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) GoString() string {
	return s.String()
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) SetId(v int64) *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity {
	s.Id = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) SetUserId(v string) *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity {
	s.UserId = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) SetRegionId(v string) *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity {
	s.RegionId = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) SetRegionName(v string) *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity {
	s.RegionName = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) SetDescription(v string) *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity {
	s.Description = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) SetBelongRegion(v string) *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity {
	s.BelongRegion = &v
	return s
}

func (s *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity) SetDebugEnable(v bool) *InsertOrUpdateRegionResponseBodyUserDefineRegionEntity {
	s.DebugEnable = &v
	return s
}

type InsertOrUpdateRegionResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertOrUpdateRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertOrUpdateRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertOrUpdateRegionResponse) GoString() string {
	return s.String()
}

func (s *InsertOrUpdateRegionResponse) SetHeaders(v map[string]*string) *InsertOrUpdateRegionResponse {
	s.Headers = v
	return s
}

func (s *InsertOrUpdateRegionResponse) SetBody(v *InsertOrUpdateRegionResponseBody) *InsertOrUpdateRegionResponse {
	s.Body = v
	return s
}

type InsertConfigCenterQuery struct {
	DataId          *string `json:"DataId,omitempty" xml:"DataId,omitempty" require:"true"`
	Group           *string `json:"Group,omitempty" xml:"Group,omitempty" require:"true"`
	Data            *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty" require:"true"`
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s InsertConfigCenterQuery) String() string {
	return tea.Prettify(s)
}

func (s InsertConfigCenterQuery) GoString() string {
	return s.String()
}

func (s *InsertConfigCenterQuery) SetDataId(v string) *InsertConfigCenterQuery {
	s.DataId = &v
	return s
}

func (s *InsertConfigCenterQuery) SetGroup(v string) *InsertConfigCenterQuery {
	s.Group = &v
	return s
}

func (s *InsertConfigCenterQuery) SetData(v string) *InsertConfigCenterQuery {
	s.Data = &v
	return s
}

func (s *InsertConfigCenterQuery) SetLogicalRegionId(v string) *InsertConfigCenterQuery {
	s.LogicalRegionId = &v
	return s
}

func (s *InsertConfigCenterQuery) SetAppName(v string) *InsertConfigCenterQuery {
	s.AppName = &v
	return s
}

type InsertConfigCenterRequest struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *InsertConfigCenterQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s InsertConfigCenterRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertConfigCenterRequest) GoString() string {
	return s.String()
}

func (s *InsertConfigCenterRequest) SetHeaders(v map[string]*string) *InsertConfigCenterRequest {
	s.Headers = v
	return s
}

func (s *InsertConfigCenterRequest) SetQuery(v *InsertConfigCenterQuery) *InsertConfigCenterRequest {
	s.Query = v
	return s
}

type InsertConfigCenterResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s InsertConfigCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertConfigCenterResponseBody) GoString() string {
	return s.String()
}

func (s *InsertConfigCenterResponseBody) SetCode(v int) *InsertConfigCenterResponseBody {
	s.Code = &v
	return s
}

func (s *InsertConfigCenterResponseBody) SetMessage(v string) *InsertConfigCenterResponseBody {
	s.Message = &v
	return s
}

func (s *InsertConfigCenterResponseBody) SetRequestId(v string) *InsertConfigCenterResponseBody {
	s.RequestId = &v
	return s
}

type InsertConfigCenterResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertConfigCenterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertConfigCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertConfigCenterResponse) GoString() string {
	return s.String()
}

func (s *InsertConfigCenterResponse) SetHeaders(v map[string]*string) *InsertConfigCenterResponse {
	s.Headers = v
	return s
}

func (s *InsertConfigCenterResponse) SetBody(v *InsertConfigCenterResponseBody) *InsertConfigCenterResponse {
	s.Body = v
	return s
}

type InsertFlowControlQuery struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	ConsumerAppId *string `json:"ConsumerAppId,omitempty" xml:"ConsumerAppId,omitempty"`
	Granularity   *string `json:"Granularity,omitempty" xml:"Granularity,omitempty" require:"true"`
	MethodName    *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	RuleType      *string `json:"RuleType,omitempty" xml:"RuleType,omitempty" require:"true"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Strategy      *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	Threshold     *int    `json:"Threshold,omitempty" xml:"Threshold,omitempty" require:"true"`
	UrlVar        *string `json:"UrlVar,omitempty" xml:"UrlVar,omitempty"`
}

func (s InsertFlowControlQuery) String() string {
	return tea.Prettify(s)
}

func (s InsertFlowControlQuery) GoString() string {
	return s.String()
}

func (s *InsertFlowControlQuery) SetAppId(v string) *InsertFlowControlQuery {
	s.AppId = &v
	return s
}

func (s *InsertFlowControlQuery) SetConsumerAppId(v string) *InsertFlowControlQuery {
	s.ConsumerAppId = &v
	return s
}

func (s *InsertFlowControlQuery) SetGranularity(v string) *InsertFlowControlQuery {
	s.Granularity = &v
	return s
}

func (s *InsertFlowControlQuery) SetMethodName(v string) *InsertFlowControlQuery {
	s.MethodName = &v
	return s
}

func (s *InsertFlowControlQuery) SetRuleType(v string) *InsertFlowControlQuery {
	s.RuleType = &v
	return s
}

func (s *InsertFlowControlQuery) SetServiceName(v string) *InsertFlowControlQuery {
	s.ServiceName = &v
	return s
}

func (s *InsertFlowControlQuery) SetStrategy(v string) *InsertFlowControlQuery {
	s.Strategy = &v
	return s
}

func (s *InsertFlowControlQuery) SetThreshold(v int) *InsertFlowControlQuery {
	s.Threshold = &v
	return s
}

func (s *InsertFlowControlQuery) SetUrlVar(v string) *InsertFlowControlQuery {
	s.UrlVar = &v
	return s
}

type InsertFlowControlRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *InsertFlowControlQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s InsertFlowControlRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertFlowControlRequest) GoString() string {
	return s.String()
}

func (s *InsertFlowControlRequest) SetHeaders(v map[string]*string) *InsertFlowControlRequest {
	s.Headers = v
	return s
}

func (s *InsertFlowControlRequest) SetQuery(v *InsertFlowControlQuery) *InsertFlowControlRequest {
	s.Query = v
	return s
}

type InsertFlowControlResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s InsertFlowControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertFlowControlResponseBody) GoString() string {
	return s.String()
}

func (s *InsertFlowControlResponseBody) SetCode(v int) *InsertFlowControlResponseBody {
	s.Code = &v
	return s
}

func (s *InsertFlowControlResponseBody) SetMessage(v string) *InsertFlowControlResponseBody {
	s.Message = &v
	return s
}

func (s *InsertFlowControlResponseBody) SetRequestId(v string) *InsertFlowControlResponseBody {
	s.RequestId = &v
	return s
}

type InsertFlowControlResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertFlowControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertFlowControlResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertFlowControlResponse) GoString() string {
	return s.String()
}

func (s *InsertFlowControlResponse) SetHeaders(v map[string]*string) *InsertFlowControlResponse {
	s.Headers = v
	return s
}

func (s *InsertFlowControlResponse) SetBody(v *InsertFlowControlResponseBody) *InsertFlowControlResponse {
	s.Body = v
	return s
}

type InsertDeployGroupQuery struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty" require:"true"`
	InitPackageVersionId *string `json:"InitPackageVersionId,omitempty" xml:"InitPackageVersionId,omitempty"`
}

func (s InsertDeployGroupQuery) String() string {
	return tea.Prettify(s)
}

func (s InsertDeployGroupQuery) GoString() string {
	return s.String()
}

func (s *InsertDeployGroupQuery) SetAppId(v string) *InsertDeployGroupQuery {
	s.AppId = &v
	return s
}

func (s *InsertDeployGroupQuery) SetGroupName(v string) *InsertDeployGroupQuery {
	s.GroupName = &v
	return s
}

func (s *InsertDeployGroupQuery) SetInitPackageVersionId(v string) *InsertDeployGroupQuery {
	s.InitPackageVersionId = &v
	return s
}

type InsertDeployGroupRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *InsertDeployGroupQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s InsertDeployGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertDeployGroupRequest) GoString() string {
	return s.String()
}

func (s *InsertDeployGroupRequest) SetHeaders(v map[string]*string) *InsertDeployGroupRequest {
	s.Headers = v
	return s
}

func (s *InsertDeployGroupRequest) SetQuery(v *InsertDeployGroupQuery) *InsertDeployGroupRequest {
	s.Query = v
	return s
}

type InsertDeployGroupResponseBody struct {
	Code              *int                                            `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message           *string                                         `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId         *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DeployGroupEntity *InsertDeployGroupResponseBodyDeployGroupEntity `json:"DeployGroupEntity,omitempty" xml:"DeployGroupEntity,omitempty" require:"true" type:"Struct"`
}

func (s InsertDeployGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertDeployGroupResponseBody) GoString() string {
	return s.String()
}

func (s *InsertDeployGroupResponseBody) SetCode(v int) *InsertDeployGroupResponseBody {
	s.Code = &v
	return s
}

func (s *InsertDeployGroupResponseBody) SetMessage(v string) *InsertDeployGroupResponseBody {
	s.Message = &v
	return s
}

func (s *InsertDeployGroupResponseBody) SetRequestId(v string) *InsertDeployGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertDeployGroupResponseBody) SetDeployGroupEntity(v *InsertDeployGroupResponseBodyDeployGroupEntity) *InsertDeployGroupResponseBody {
	s.DeployGroupEntity = v
	return s
}

type InsertDeployGroupResponseBodyDeployGroupEntity struct {
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	GroupName        *string `json:"GroupName,omitempty" xml:"GroupName,omitempty" require:"true"`
	GroupType        *int    `json:"GroupType,omitempty" xml:"GroupType,omitempty" require:"true"`
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	PackageVersionId *string `json:"PackageVersionId,omitempty" xml:"PackageVersionId,omitempty" require:"true"`
	AppVersionId     *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty" require:"true"`
	CreateTime       *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime       *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
}

func (s InsertDeployGroupResponseBodyDeployGroupEntity) String() string {
	return tea.Prettify(s)
}

func (s InsertDeployGroupResponseBodyDeployGroupEntity) GoString() string {
	return s.String()
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetId(v string) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.Id = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetGroupName(v string) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.GroupName = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetGroupType(v int) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.GroupType = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetAppId(v string) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.AppId = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetClusterId(v string) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.ClusterId = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetPackageVersionId(v string) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.PackageVersionId = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetAppVersionId(v string) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.AppVersionId = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetCreateTime(v int64) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.CreateTime = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetUpdateTime(v int64) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.UpdateTime = &v
	return s
}

type InsertDeployGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertDeployGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertDeployGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertDeployGroupResponse) GoString() string {
	return s.String()
}

func (s *InsertDeployGroupResponse) SetHeaders(v map[string]*string) *InsertDeployGroupResponse {
	s.Headers = v
	return s
}

func (s *InsertDeployGroupResponse) SetBody(v *InsertDeployGroupResponseBody) *InsertDeployGroupResponse {
	s.Body = v
	return s
}

type InsertDegradeControlQuery struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Duration    *int    `json:"Duration,omitempty" xml:"Duration,omitempty" require:"true"`
	MethodName  *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	RtThreshold *int    `json:"RtThreshold,omitempty" xml:"RtThreshold,omitempty" require:"true"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	UrlVar      *string `json:"UrlVar,omitempty" xml:"UrlVar,omitempty"`
	RuleType    *string `json:"RuleType,omitempty" xml:"RuleType,omitempty" require:"true"`
}

func (s InsertDegradeControlQuery) String() string {
	return tea.Prettify(s)
}

func (s InsertDegradeControlQuery) GoString() string {
	return s.String()
}

func (s *InsertDegradeControlQuery) SetAppId(v string) *InsertDegradeControlQuery {
	s.AppId = &v
	return s
}

func (s *InsertDegradeControlQuery) SetDuration(v int) *InsertDegradeControlQuery {
	s.Duration = &v
	return s
}

func (s *InsertDegradeControlQuery) SetMethodName(v string) *InsertDegradeControlQuery {
	s.MethodName = &v
	return s
}

func (s *InsertDegradeControlQuery) SetRtThreshold(v int) *InsertDegradeControlQuery {
	s.RtThreshold = &v
	return s
}

func (s *InsertDegradeControlQuery) SetServiceName(v string) *InsertDegradeControlQuery {
	s.ServiceName = &v
	return s
}

func (s *InsertDegradeControlQuery) SetUrlVar(v string) *InsertDegradeControlQuery {
	s.UrlVar = &v
	return s
}

func (s *InsertDegradeControlQuery) SetRuleType(v string) *InsertDegradeControlQuery {
	s.RuleType = &v
	return s
}

type InsertDegradeControlRequest struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *InsertDegradeControlQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s InsertDegradeControlRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertDegradeControlRequest) GoString() string {
	return s.String()
}

func (s *InsertDegradeControlRequest) SetHeaders(v map[string]*string) *InsertDegradeControlRequest {
	s.Headers = v
	return s
}

func (s *InsertDegradeControlRequest) SetQuery(v *InsertDegradeControlQuery) *InsertDegradeControlRequest {
	s.Query = v
	return s
}

type InsertDegradeControlResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s InsertDegradeControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertDegradeControlResponseBody) GoString() string {
	return s.String()
}

func (s *InsertDegradeControlResponseBody) SetCode(v int) *InsertDegradeControlResponseBody {
	s.Code = &v
	return s
}

func (s *InsertDegradeControlResponseBody) SetMessage(v string) *InsertDegradeControlResponseBody {
	s.Message = &v
	return s
}

func (s *InsertDegradeControlResponseBody) SetRequestId(v string) *InsertDegradeControlResponseBody {
	s.RequestId = &v
	return s
}

type InsertDegradeControlResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertDegradeControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertDegradeControlResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertDegradeControlResponse) GoString() string {
	return s.String()
}

func (s *InsertDegradeControlResponse) SetHeaders(v map[string]*string) *InsertDegradeControlResponse {
	s.Headers = v
	return s
}

func (s *InsertDegradeControlResponse) SetBody(v *InsertDegradeControlResponseBody) *InsertDegradeControlResponse {
	s.Body = v
	return s
}

type InsertClusterQuery struct {
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	ClusterName     *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty" require:"true"`
	ClusterType     *int    `json:"ClusterType,omitempty" xml:"ClusterType,omitempty" require:"true"`
	NetworkMode     *int    `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty" require:"true"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	OversoldFactor  *int    `json:"OversoldFactor,omitempty" xml:"OversoldFactor,omitempty"`
	IaasProvider    *string `json:"IaasProvider,omitempty" xml:"IaasProvider,omitempty"`
}

func (s InsertClusterQuery) String() string {
	return tea.Prettify(s)
}

func (s InsertClusterQuery) GoString() string {
	return s.String()
}

func (s *InsertClusterQuery) SetLogicalRegionId(v string) *InsertClusterQuery {
	s.LogicalRegionId = &v
	return s
}

func (s *InsertClusterQuery) SetClusterName(v string) *InsertClusterQuery {
	s.ClusterName = &v
	return s
}

func (s *InsertClusterQuery) SetClusterType(v int) *InsertClusterQuery {
	s.ClusterType = &v
	return s
}

func (s *InsertClusterQuery) SetNetworkMode(v int) *InsertClusterQuery {
	s.NetworkMode = &v
	return s
}

func (s *InsertClusterQuery) SetVpcId(v string) *InsertClusterQuery {
	s.VpcId = &v
	return s
}

func (s *InsertClusterQuery) SetOversoldFactor(v int) *InsertClusterQuery {
	s.OversoldFactor = &v
	return s
}

func (s *InsertClusterQuery) SetIaasProvider(v string) *InsertClusterQuery {
	s.IaasProvider = &v
	return s
}

type InsertClusterRequest struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *InsertClusterQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s InsertClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertClusterRequest) GoString() string {
	return s.String()
}

func (s *InsertClusterRequest) SetHeaders(v map[string]*string) *InsertClusterRequest {
	s.Headers = v
	return s
}

func (s *InsertClusterRequest) SetQuery(v *InsertClusterQuery) *InsertClusterRequest {
	s.Query = v
	return s
}

type InsertClusterResponseBody struct {
	Code      *int                              `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Cluster   *InsertClusterResponseBodyCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" require:"true" type:"Struct"`
}

func (s InsertClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertClusterResponseBody) GoString() string {
	return s.String()
}

func (s *InsertClusterResponseBody) SetCode(v int) *InsertClusterResponseBody {
	s.Code = &v
	return s
}

func (s *InsertClusterResponseBody) SetMessage(v string) *InsertClusterResponseBody {
	s.Message = &v
	return s
}

func (s *InsertClusterResponseBody) SetRequestId(v string) *InsertClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertClusterResponseBody) SetCluster(v *InsertClusterResponseBodyCluster) *InsertClusterResponseBody {
	s.Cluster = v
	return s
}

type InsertClusterResponseBodyCluster struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ClusterName    *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty" require:"true"`
	ClusterType    *int    `json:"ClusterType,omitempty" xml:"ClusterType,omitempty" require:"true"`
	OversoldFactor *int    `json:"OversoldFactor,omitempty" xml:"OversoldFactor,omitempty" require:"true"`
	NetworkMode    *int    `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty" require:"true"`
	VpcId          *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	IaasProvider   *string `json:"IaasProvider,omitempty" xml:"IaasProvider,omitempty" require:"true"`
}

func (s InsertClusterResponseBodyCluster) String() string {
	return tea.Prettify(s)
}

func (s InsertClusterResponseBodyCluster) GoString() string {
	return s.String()
}

func (s *InsertClusterResponseBodyCluster) SetClusterId(v string) *InsertClusterResponseBodyCluster {
	s.ClusterId = &v
	return s
}

func (s *InsertClusterResponseBodyCluster) SetRegionId(v string) *InsertClusterResponseBodyCluster {
	s.RegionId = &v
	return s
}

func (s *InsertClusterResponseBodyCluster) SetClusterName(v string) *InsertClusterResponseBodyCluster {
	s.ClusterName = &v
	return s
}

func (s *InsertClusterResponseBodyCluster) SetClusterType(v int) *InsertClusterResponseBodyCluster {
	s.ClusterType = &v
	return s
}

func (s *InsertClusterResponseBodyCluster) SetOversoldFactor(v int) *InsertClusterResponseBodyCluster {
	s.OversoldFactor = &v
	return s
}

func (s *InsertClusterResponseBodyCluster) SetNetworkMode(v int) *InsertClusterResponseBodyCluster {
	s.NetworkMode = &v
	return s
}

func (s *InsertClusterResponseBodyCluster) SetVpcId(v string) *InsertClusterResponseBodyCluster {
	s.VpcId = &v
	return s
}

func (s *InsertClusterResponseBodyCluster) SetIaasProvider(v string) *InsertClusterResponseBodyCluster {
	s.IaasProvider = &v
	return s
}

type InsertClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertClusterResponse) GoString() string {
	return s.String()
}

func (s *InsertClusterResponse) SetHeaders(v map[string]*string) *InsertClusterResponse {
	s.Headers = v
	return s
}

func (s *InsertClusterResponse) SetBody(v *InsertClusterResponseBody) *InsertClusterResponse {
	s.Body = v
	return s
}

type GetJvmConfigurationQuery struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s GetJvmConfigurationQuery) String() string {
	return tea.Prettify(s)
}

func (s GetJvmConfigurationQuery) GoString() string {
	return s.String()
}

func (s *GetJvmConfigurationQuery) SetAppId(v string) *GetJvmConfigurationQuery {
	s.AppId = &v
	return s
}

func (s *GetJvmConfigurationQuery) SetGroupId(v string) *GetJvmConfigurationQuery {
	s.GroupId = &v
	return s
}

type GetJvmConfigurationRequest struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *GetJvmConfigurationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s GetJvmConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJvmConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetJvmConfigurationRequest) SetHeaders(v map[string]*string) *GetJvmConfigurationRequest {
	s.Headers = v
	return s
}

func (s *GetJvmConfigurationRequest) SetQuery(v *GetJvmConfigurationQuery) *GetJvmConfigurationRequest {
	s.Query = v
	return s
}

type GetJvmConfigurationResponseBody struct {
	Code             *int                                             `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message          *string                                          `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	JvmConfiguration *GetJvmConfigurationResponseBodyJvmConfiguration `json:"JvmConfiguration,omitempty" xml:"JvmConfiguration,omitempty" require:"true" type:"Struct"`
}

func (s GetJvmConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJvmConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetJvmConfigurationResponseBody) SetCode(v int) *GetJvmConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *GetJvmConfigurationResponseBody) SetMessage(v string) *GetJvmConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *GetJvmConfigurationResponseBody) SetRequestId(v string) *GetJvmConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJvmConfigurationResponseBody) SetJvmConfiguration(v *GetJvmConfigurationResponseBodyJvmConfiguration) *GetJvmConfigurationResponseBody {
	s.JvmConfiguration = v
	return s
}

type GetJvmConfigurationResponseBodyJvmConfiguration struct {
	Options     *string `json:"Options,omitempty" xml:"Options,omitempty" require:"true"`
	MinHeapSize *int    `json:"MinHeapSize,omitempty" xml:"MinHeapSize,omitempty" require:"true"`
	MaxPermSize *int    `json:"MaxPermSize,omitempty" xml:"MaxPermSize,omitempty" require:"true"`
	MaxHeapSize *int    `json:"MaxHeapSize,omitempty" xml:"MaxHeapSize,omitempty" require:"true"`
}

func (s GetJvmConfigurationResponseBodyJvmConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetJvmConfigurationResponseBodyJvmConfiguration) GoString() string {
	return s.String()
}

func (s *GetJvmConfigurationResponseBodyJvmConfiguration) SetOptions(v string) *GetJvmConfigurationResponseBodyJvmConfiguration {
	s.Options = &v
	return s
}

func (s *GetJvmConfigurationResponseBodyJvmConfiguration) SetMinHeapSize(v int) *GetJvmConfigurationResponseBodyJvmConfiguration {
	s.MinHeapSize = &v
	return s
}

func (s *GetJvmConfigurationResponseBodyJvmConfiguration) SetMaxPermSize(v int) *GetJvmConfigurationResponseBodyJvmConfiguration {
	s.MaxPermSize = &v
	return s
}

func (s *GetJvmConfigurationResponseBodyJvmConfiguration) SetMaxHeapSize(v int) *GetJvmConfigurationResponseBodyJvmConfiguration {
	s.MaxHeapSize = &v
	return s
}

type GetJvmConfigurationResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJvmConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJvmConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJvmConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetJvmConfigurationResponse) SetHeaders(v map[string]*string) *GetJvmConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetJvmConfigurationResponse) SetBody(v *GetJvmConfigurationResponseBody) *GetJvmConfigurationResponse {
	s.Body = v
	return s
}

type GetContainerConfigurationQuery struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s GetContainerConfigurationQuery) String() string {
	return tea.Prettify(s)
}

func (s GetContainerConfigurationQuery) GoString() string {
	return s.String()
}

func (s *GetContainerConfigurationQuery) SetAppId(v string) *GetContainerConfigurationQuery {
	s.AppId = &v
	return s
}

func (s *GetContainerConfigurationQuery) SetGroupId(v string) *GetContainerConfigurationQuery {
	s.GroupId = &v
	return s
}

type GetContainerConfigurationRequest struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *GetContainerConfigurationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s GetContainerConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetContainerConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetContainerConfigurationRequest) SetHeaders(v map[string]*string) *GetContainerConfigurationRequest {
	s.Headers = v
	return s
}

func (s *GetContainerConfigurationRequest) SetQuery(v *GetContainerConfigurationQuery) *GetContainerConfigurationRequest {
	s.Query = v
	return s
}

type GetContainerConfigurationResponseBody struct {
	Code                   *int                                                         `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message                *string                                                      `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId              *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ContainerConfiguration *GetContainerConfigurationResponseBodyContainerConfiguration `json:"ContainerConfiguration,omitempty" xml:"ContainerConfiguration,omitempty" require:"true" type:"Struct"`
}

func (s GetContainerConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetContainerConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetContainerConfigurationResponseBody) SetCode(v int) *GetContainerConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *GetContainerConfigurationResponseBody) SetMessage(v string) *GetContainerConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *GetContainerConfigurationResponseBody) SetRequestId(v string) *GetContainerConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContainerConfigurationResponseBody) SetContainerConfiguration(v *GetContainerConfigurationResponseBodyContainerConfiguration) *GetContainerConfigurationResponseBody {
	s.ContainerConfiguration = v
	return s
}

type GetContainerConfigurationResponseBodyContainerConfiguration struct {
	ContextPath     *string `json:"ContextPath,omitempty" xml:"ContextPath,omitempty" require:"true"`
	HttpPort        *int    `json:"HttpPort,omitempty" xml:"HttpPort,omitempty" require:"true"`
	MaxThreads      *int    `json:"MaxThreads,omitempty" xml:"MaxThreads,omitempty" require:"true"`
	URIEncoding     *string `json:"URIEncoding,omitempty" xml:"URIEncoding,omitempty" require:"true"`
	UseBodyEncoding *bool   `json:"UseBodyEncoding,omitempty" xml:"UseBodyEncoding,omitempty" require:"true"`
}

func (s GetContainerConfigurationResponseBodyContainerConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetContainerConfigurationResponseBodyContainerConfiguration) GoString() string {
	return s.String()
}

func (s *GetContainerConfigurationResponseBodyContainerConfiguration) SetContextPath(v string) *GetContainerConfigurationResponseBodyContainerConfiguration {
	s.ContextPath = &v
	return s
}

func (s *GetContainerConfigurationResponseBodyContainerConfiguration) SetHttpPort(v int) *GetContainerConfigurationResponseBodyContainerConfiguration {
	s.HttpPort = &v
	return s
}

func (s *GetContainerConfigurationResponseBodyContainerConfiguration) SetMaxThreads(v int) *GetContainerConfigurationResponseBodyContainerConfiguration {
	s.MaxThreads = &v
	return s
}

func (s *GetContainerConfigurationResponseBodyContainerConfiguration) SetURIEncoding(v string) *GetContainerConfigurationResponseBodyContainerConfiguration {
	s.URIEncoding = &v
	return s
}

func (s *GetContainerConfigurationResponseBodyContainerConfiguration) SetUseBodyEncoding(v bool) *GetContainerConfigurationResponseBodyContainerConfiguration {
	s.UseBodyEncoding = &v
	return s
}

type GetContainerConfigurationResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetContainerConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetContainerConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetContainerConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetContainerConfigurationResponse) SetHeaders(v map[string]*string) *GetContainerConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetContainerConfigurationResponse) SetBody(v *GetContainerConfigurationResponseBody) *GetContainerConfigurationResponse {
	s.Body = v
	return s
}

type GetApplicationQuery struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s GetApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationQuery) GoString() string {
	return s.String()
}

func (s *GetApplicationQuery) SetAppId(v string) *GetApplicationQuery {
	s.AppId = &v
	return s
}

type GetApplicationRequest struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *GetApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s GetApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationRequest) SetHeaders(v map[string]*string) *GetApplicationRequest {
	s.Headers = v
	return s
}

func (s *GetApplicationRequest) SetQuery(v *GetApplicationQuery) *GetApplicationRequest {
	s.Query = v
	return s
}

type GetApplicationResponseBody struct {
	Code       *int                                  `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Applcation *GetApplicationResponseBodyApplcation `json:"Applcation,omitempty" xml:"Applcation,omitempty" require:"true" type:"Struct"`
}

func (s GetApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) SetCode(v int) *GetApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *GetApplicationResponseBody) SetMessage(v string) *GetApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *GetApplicationResponseBody) SetRequestId(v string) *GetApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationResponseBody) SetApplcation(v *GetApplicationResponseBodyApplcation) *GetApplicationResponseBody {
	s.Applcation = v
	return s
}

type GetApplicationResponseBodyApplcation struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Owner                *string `json:"Owner,omitempty" xml:"Owner,omitempty" require:"true"`
	InstanceCount        *int    `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty" require:"true"`
	RunningInstanceCount *int    `json:"RunningInstanceCount,omitempty" xml:"RunningInstanceCount,omitempty" require:"true"`
	Port                 *int    `json:"Port,omitempty" xml:"Port,omitempty" require:"true"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	SlbId                *string `json:"SlbId,omitempty" xml:"SlbId,omitempty" require:"true"`
	SlbIp                *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty" require:"true"`
	SlbPort              *int    `json:"SlbPort,omitempty" xml:"SlbPort,omitempty" require:"true"`
	ExtSlbId             *string `json:"ExtSlbId,omitempty" xml:"ExtSlbId,omitempty" require:"true"`
	ExtSlbIp             *string `json:"ExtSlbIp,omitempty" xml:"ExtSlbIp,omitempty" require:"true"`
	SlbName              *string `json:"SlbName,omitempty" xml:"SlbName,omitempty" require:"true"`
	ExtSlbName           *string `json:"ExtSlbName,omitempty" xml:"ExtSlbName,omitempty" require:"true"`
	ApplicationType      *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty" require:"true"`
	ClusterType          *int    `json:"ClusterType,omitempty" xml:"ClusterType,omitempty" require:"true"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	Dockerize            *bool   `json:"Dockerize,omitempty" xml:"Dockerize,omitempty" require:"true"`
	Cpu                  *int    `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Memory               *int    `json:"Memory,omitempty" xml:"Memory,omitempty" require:"true"`
	HealthCheckUrl       *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty" require:"true"`
	BuildPackageId       *int64  `json:"BuildPackageId,omitempty" xml:"BuildPackageId,omitempty" require:"true"`
	CreateTime           *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	NameSpace            *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty" require:"true"`
	SlbInfo              *string `json:"SlbInfo,omitempty" xml:"SlbInfo,omitempty" require:"true"`
	Email                *string `json:"Email,omitempty" xml:"Email,omitempty" require:"true"`
}

func (s GetApplicationResponseBodyApplcation) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyApplcation) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplcation) SetAppId(v string) *GetApplicationResponseBodyApplcation {
	s.AppId = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetName(v string) *GetApplicationResponseBodyApplcation {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetRegionId(v string) *GetApplicationResponseBodyApplcation {
	s.RegionId = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetDescription(v string) *GetApplicationResponseBodyApplcation {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetOwner(v string) *GetApplicationResponseBodyApplcation {
	s.Owner = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetInstanceCount(v int) *GetApplicationResponseBodyApplcation {
	s.InstanceCount = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetRunningInstanceCount(v int) *GetApplicationResponseBodyApplcation {
	s.RunningInstanceCount = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetPort(v int) *GetApplicationResponseBodyApplcation {
	s.Port = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetUserId(v string) *GetApplicationResponseBodyApplcation {
	s.UserId = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetSlbId(v string) *GetApplicationResponseBodyApplcation {
	s.SlbId = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetSlbIp(v string) *GetApplicationResponseBodyApplcation {
	s.SlbIp = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetSlbPort(v int) *GetApplicationResponseBodyApplcation {
	s.SlbPort = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetExtSlbId(v string) *GetApplicationResponseBodyApplcation {
	s.ExtSlbId = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetExtSlbIp(v string) *GetApplicationResponseBodyApplcation {
	s.ExtSlbIp = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetSlbName(v string) *GetApplicationResponseBodyApplcation {
	s.SlbName = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetExtSlbName(v string) *GetApplicationResponseBodyApplcation {
	s.ExtSlbName = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetApplicationType(v string) *GetApplicationResponseBodyApplcation {
	s.ApplicationType = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetClusterType(v int) *GetApplicationResponseBodyApplcation {
	s.ClusterType = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetClusterId(v string) *GetApplicationResponseBodyApplcation {
	s.ClusterId = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetDockerize(v bool) *GetApplicationResponseBodyApplcation {
	s.Dockerize = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetCpu(v int) *GetApplicationResponseBodyApplcation {
	s.Cpu = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetMemory(v int) *GetApplicationResponseBodyApplcation {
	s.Memory = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetHealthCheckUrl(v string) *GetApplicationResponseBodyApplcation {
	s.HealthCheckUrl = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetBuildPackageId(v int64) *GetApplicationResponseBodyApplcation {
	s.BuildPackageId = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetCreateTime(v int64) *GetApplicationResponseBodyApplcation {
	s.CreateTime = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetNameSpace(v string) *GetApplicationResponseBodyApplcation {
	s.NameSpace = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetSlbInfo(v string) *GetApplicationResponseBodyApplcation {
	s.SlbInfo = &v
	return s
}

func (s *GetApplicationResponseBodyApplcation) SetEmail(v string) *GetApplicationResponseBodyApplcation {
	s.Email = &v
	return s
}

type GetApplicationResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationResponse) SetHeaders(v map[string]*string) *GetApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationResponse) SetBody(v *GetApplicationResponseBody) *GetApplicationResponse {
	s.Body = v
	return s
}

type EnableFlowControlQuery struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty" require:"true"`
}

func (s EnableFlowControlQuery) String() string {
	return tea.Prettify(s)
}

func (s EnableFlowControlQuery) GoString() string {
	return s.String()
}

func (s *EnableFlowControlQuery) SetAppId(v string) *EnableFlowControlQuery {
	s.AppId = &v
	return s
}

func (s *EnableFlowControlQuery) SetRuleId(v string) *EnableFlowControlQuery {
	s.RuleId = &v
	return s
}

type EnableFlowControlRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *EnableFlowControlQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s EnableFlowControlRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableFlowControlRequest) GoString() string {
	return s.String()
}

func (s *EnableFlowControlRequest) SetHeaders(v map[string]*string) *EnableFlowControlRequest {
	s.Headers = v
	return s
}

func (s *EnableFlowControlRequest) SetQuery(v *EnableFlowControlQuery) *EnableFlowControlRequest {
	s.Query = v
	return s
}

type EnableFlowControlResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s EnableFlowControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableFlowControlResponseBody) GoString() string {
	return s.String()
}

func (s *EnableFlowControlResponseBody) SetCode(v int) *EnableFlowControlResponseBody {
	s.Code = &v
	return s
}

func (s *EnableFlowControlResponseBody) SetMessage(v string) *EnableFlowControlResponseBody {
	s.Message = &v
	return s
}

func (s *EnableFlowControlResponseBody) SetRequestId(v string) *EnableFlowControlResponseBody {
	s.RequestId = &v
	return s
}

type EnableFlowControlResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableFlowControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableFlowControlResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableFlowControlResponse) GoString() string {
	return s.String()
}

func (s *EnableFlowControlResponse) SetHeaders(v map[string]*string) *EnableFlowControlResponse {
	s.Headers = v
	return s
}

func (s *EnableFlowControlResponse) SetBody(v *EnableFlowControlResponseBody) *EnableFlowControlResponse {
	s.Body = v
	return s
}

type EnableDegradeControlQuery struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty" require:"true"`
}

func (s EnableDegradeControlQuery) String() string {
	return tea.Prettify(s)
}

func (s EnableDegradeControlQuery) GoString() string {
	return s.String()
}

func (s *EnableDegradeControlQuery) SetAppId(v string) *EnableDegradeControlQuery {
	s.AppId = &v
	return s
}

func (s *EnableDegradeControlQuery) SetRuleId(v string) *EnableDegradeControlQuery {
	s.RuleId = &v
	return s
}

type EnableDegradeControlRequest struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *EnableDegradeControlQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s EnableDegradeControlRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableDegradeControlRequest) GoString() string {
	return s.String()
}

func (s *EnableDegradeControlRequest) SetHeaders(v map[string]*string) *EnableDegradeControlRequest {
	s.Headers = v
	return s
}

func (s *EnableDegradeControlRequest) SetQuery(v *EnableDegradeControlQuery) *EnableDegradeControlRequest {
	s.Query = v
	return s
}

type EnableDegradeControlResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s EnableDegradeControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableDegradeControlResponseBody) GoString() string {
	return s.String()
}

func (s *EnableDegradeControlResponseBody) SetCode(v int) *EnableDegradeControlResponseBody {
	s.Code = &v
	return s
}

func (s *EnableDegradeControlResponseBody) SetMessage(v string) *EnableDegradeControlResponseBody {
	s.Message = &v
	return s
}

func (s *EnableDegradeControlResponseBody) SetRequestId(v string) *EnableDegradeControlResponseBody {
	s.RequestId = &v
	return s
}

type EnableDegradeControlResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableDegradeControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableDegradeControlResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableDegradeControlResponse) GoString() string {
	return s.String()
}

func (s *EnableDegradeControlResponse) SetHeaders(v map[string]*string) *EnableDegradeControlResponse {
	s.Headers = v
	return s
}

func (s *EnableDegradeControlResponse) SetBody(v *EnableDegradeControlResponseBody) *EnableDegradeControlResponse {
	s.Body = v
	return s
}

type DisableFlowControlQuery struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty" require:"true"`
}

func (s DisableFlowControlQuery) String() string {
	return tea.Prettify(s)
}

func (s DisableFlowControlQuery) GoString() string {
	return s.String()
}

func (s *DisableFlowControlQuery) SetAppId(v string) *DisableFlowControlQuery {
	s.AppId = &v
	return s
}

func (s *DisableFlowControlQuery) SetRuleId(v string) *DisableFlowControlQuery {
	s.RuleId = &v
	return s
}

type DisableFlowControlRequest struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DisableFlowControlQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DisableFlowControlRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableFlowControlRequest) GoString() string {
	return s.String()
}

func (s *DisableFlowControlRequest) SetHeaders(v map[string]*string) *DisableFlowControlRequest {
	s.Headers = v
	return s
}

func (s *DisableFlowControlRequest) SetQuery(v *DisableFlowControlQuery) *DisableFlowControlRequest {
	s.Query = v
	return s
}

type DisableFlowControlResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DisableFlowControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableFlowControlResponseBody) GoString() string {
	return s.String()
}

func (s *DisableFlowControlResponseBody) SetCode(v int) *DisableFlowControlResponseBody {
	s.Code = &v
	return s
}

func (s *DisableFlowControlResponseBody) SetMessage(v string) *DisableFlowControlResponseBody {
	s.Message = &v
	return s
}

func (s *DisableFlowControlResponseBody) SetRequestId(v string) *DisableFlowControlResponseBody {
	s.RequestId = &v
	return s
}

type DisableFlowControlResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableFlowControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableFlowControlResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableFlowControlResponse) GoString() string {
	return s.String()
}

func (s *DisableFlowControlResponse) SetHeaders(v map[string]*string) *DisableFlowControlResponse {
	s.Headers = v
	return s
}

func (s *DisableFlowControlResponse) SetBody(v *DisableFlowControlResponseBody) *DisableFlowControlResponse {
	s.Body = v
	return s
}

type DisableDegradeControlQuery struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty" require:"true"`
}

func (s DisableDegradeControlQuery) String() string {
	return tea.Prettify(s)
}

func (s DisableDegradeControlQuery) GoString() string {
	return s.String()
}

func (s *DisableDegradeControlQuery) SetAppId(v string) *DisableDegradeControlQuery {
	s.AppId = &v
	return s
}

func (s *DisableDegradeControlQuery) SetRuleId(v string) *DisableDegradeControlQuery {
	s.RuleId = &v
	return s
}

type DisableDegradeControlRequest struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DisableDegradeControlQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DisableDegradeControlRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableDegradeControlRequest) GoString() string {
	return s.String()
}

func (s *DisableDegradeControlRequest) SetHeaders(v map[string]*string) *DisableDegradeControlRequest {
	s.Headers = v
	return s
}

func (s *DisableDegradeControlRequest) SetQuery(v *DisableDegradeControlQuery) *DisableDegradeControlRequest {
	s.Query = v
	return s
}

type DisableDegradeControlResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DisableDegradeControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableDegradeControlResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDegradeControlResponseBody) SetCode(v int) *DisableDegradeControlResponseBody {
	s.Code = &v
	return s
}

func (s *DisableDegradeControlResponseBody) SetMessage(v string) *DisableDegradeControlResponseBody {
	s.Message = &v
	return s
}

func (s *DisableDegradeControlResponseBody) SetRequestId(v string) *DisableDegradeControlResponseBody {
	s.RequestId = &v
	return s
}

type DisableDegradeControlResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableDegradeControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableDegradeControlResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableDegradeControlResponse) GoString() string {
	return s.String()
}

func (s *DisableDegradeControlResponse) SetHeaders(v map[string]*string) *DisableDegradeControlResponse {
	s.Headers = v
	return s
}

func (s *DisableDegradeControlResponse) SetBody(v *DisableDegradeControlResponseBody) *DisableDegradeControlResponse {
	s.Body = v
	return s
}

type DeleteUserDefineRegionQuery struct {
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	RegionTag *string `json:"RegionTag,omitempty" xml:"RegionTag,omitempty"`
}

func (s DeleteUserDefineRegionQuery) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDefineRegionQuery) GoString() string {
	return s.String()
}

func (s *DeleteUserDefineRegionQuery) SetId(v int64) *DeleteUserDefineRegionQuery {
	s.Id = &v
	return s
}

func (s *DeleteUserDefineRegionQuery) SetRegionTag(v string) *DeleteUserDefineRegionQuery {
	s.RegionTag = &v
	return s
}

type DeleteUserDefineRegionRequest struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeleteUserDefineRegionQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DeleteUserDefineRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDefineRegionRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserDefineRegionRequest) SetHeaders(v map[string]*string) *DeleteUserDefineRegionRequest {
	s.Headers = v
	return s
}

func (s *DeleteUserDefineRegionRequest) SetQuery(v *DeleteUserDefineRegionQuery) *DeleteUserDefineRegionRequest {
	s.Query = v
	return s
}

type DeleteUserDefineRegionResponseBody struct {
	Code         *int                                            `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message      *string                                         `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionDefine *DeleteUserDefineRegionResponseBodyRegionDefine `json:"RegionDefine,omitempty" xml:"RegionDefine,omitempty" require:"true" type:"Struct"`
}

func (s DeleteUserDefineRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDefineRegionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserDefineRegionResponseBody) SetCode(v int) *DeleteUserDefineRegionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBody) SetMessage(v string) *DeleteUserDefineRegionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBody) SetRequestId(v string) *DeleteUserDefineRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBody) SetRegionDefine(v *DeleteUserDefineRegionResponseBodyRegionDefine) *DeleteUserDefineRegionResponseBody {
	s.RegionDefine = v
	return s
}

type DeleteUserDefineRegionResponseBodyRegionDefine struct {
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	RegionName   *string `json:"RegionName,omitempty" xml:"RegionName,omitempty" require:"true"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	BelongRegion *string `json:"BelongRegion,omitempty" xml:"BelongRegion,omitempty" require:"true"`
}

func (s DeleteUserDefineRegionResponseBodyRegionDefine) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDefineRegionResponseBodyRegionDefine) GoString() string {
	return s.String()
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) SetId(v int64) *DeleteUserDefineRegionResponseBodyRegionDefine {
	s.Id = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) SetUserId(v string) *DeleteUserDefineRegionResponseBodyRegionDefine {
	s.UserId = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) SetRegionId(v string) *DeleteUserDefineRegionResponseBodyRegionDefine {
	s.RegionId = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) SetRegionName(v string) *DeleteUserDefineRegionResponseBodyRegionDefine {
	s.RegionName = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) SetDescription(v string) *DeleteUserDefineRegionResponseBodyRegionDefine {
	s.Description = &v
	return s
}

func (s *DeleteUserDefineRegionResponseBodyRegionDefine) SetBelongRegion(v string) *DeleteUserDefineRegionResponseBodyRegionDefine {
	s.BelongRegion = &v
	return s
}

type DeleteUserDefineRegionResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserDefineRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserDefineRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDefineRegionResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserDefineRegionResponse) SetHeaders(v map[string]*string) *DeleteUserDefineRegionResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserDefineRegionResponse) SetBody(v *DeleteUserDefineRegionResponseBody) *DeleteUserDefineRegionResponse {
	s.Body = v
	return s
}

type DeleteServiceGroupQuery struct {
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
}

func (s DeleteServiceGroupQuery) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceGroupQuery) GoString() string {
	return s.String()
}

func (s *DeleteServiceGroupQuery) SetGroupId(v string) *DeleteServiceGroupQuery {
	s.GroupId = &v
	return s
}

type DeleteServiceGroupRequest struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeleteServiceGroupQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DeleteServiceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceGroupRequest) SetHeaders(v map[string]*string) *DeleteServiceGroupRequest {
	s.Headers = v
	return s
}

func (s *DeleteServiceGroupRequest) SetQuery(v *DeleteServiceGroupQuery) *DeleteServiceGroupRequest {
	s.Query = v
	return s
}

type DeleteServiceGroupResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteServiceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceGroupResponseBody) SetCode(v int) *DeleteServiceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteServiceGroupResponseBody) SetMessage(v string) *DeleteServiceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceGroupResponseBody) SetRequestId(v string) *DeleteServiceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteServiceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceGroupResponse) SetHeaders(v map[string]*string) *DeleteServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceGroupResponse) SetBody(v *DeleteServiceGroupResponseBody) *DeleteServiceGroupResponse {
	s.Body = v
	return s
}

type DeleteRoleQuery struct {
	RoleId *int `json:"RoleId,omitempty" xml:"RoleId,omitempty" require:"true"`
}

func (s DeleteRoleQuery) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoleQuery) GoString() string {
	return s.String()
}

func (s *DeleteRoleQuery) SetRoleId(v int) *DeleteRoleQuery {
	s.RoleId = &v
	return s
}

type DeleteRoleRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeleteRoleQuery   `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DeleteRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoleRequest) SetHeaders(v map[string]*string) *DeleteRoleRequest {
	s.Headers = v
	return s
}

func (s *DeleteRoleRequest) SetQuery(v *DeleteRoleQuery) *DeleteRoleRequest {
	s.Query = v
	return s
}

type DeleteRoleResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoleResponseBody) SetCode(v int) *DeleteRoleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRoleResponseBody) SetMessage(v string) *DeleteRoleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRoleResponseBody) SetRequestId(v string) *DeleteRoleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRoleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoleResponse) SetHeaders(v map[string]*string) *DeleteRoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoleResponse) SetBody(v *DeleteRoleResponseBody) *DeleteRoleResponse {
	s.Body = v
	return s
}

type DeleteFlowControlQuery struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty" require:"true"`
}

func (s DeleteFlowControlQuery) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowControlQuery) GoString() string {
	return s.String()
}

func (s *DeleteFlowControlQuery) SetAppId(v string) *DeleteFlowControlQuery {
	s.AppId = &v
	return s
}

func (s *DeleteFlowControlQuery) SetRuleId(v string) *DeleteFlowControlQuery {
	s.RuleId = &v
	return s
}

type DeleteFlowControlRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeleteFlowControlQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DeleteFlowControlRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowControlRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowControlRequest) SetHeaders(v map[string]*string) *DeleteFlowControlRequest {
	s.Headers = v
	return s
}

func (s *DeleteFlowControlRequest) SetQuery(v *DeleteFlowControlQuery) *DeleteFlowControlRequest {
	s.Query = v
	return s
}

type DeleteFlowControlResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteFlowControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowControlResponseBody) SetCode(v int) *DeleteFlowControlResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFlowControlResponseBody) SetMessage(v string) *DeleteFlowControlResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFlowControlResponseBody) SetRequestId(v string) *DeleteFlowControlResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFlowControlResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowControlResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowControlResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowControlResponse) SetHeaders(v map[string]*string) *DeleteFlowControlResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowControlResponse) SetBody(v *DeleteFlowControlResponseBody) *DeleteFlowControlResponse {
	s.Body = v
	return s
}

type DeleteEcuQuery struct {
	EcuId *string `json:"EcuId,omitempty" xml:"EcuId,omitempty" require:"true"`
}

func (s DeleteEcuQuery) String() string {
	return tea.Prettify(s)
}

func (s DeleteEcuQuery) GoString() string {
	return s.String()
}

func (s *DeleteEcuQuery) SetEcuId(v string) *DeleteEcuQuery {
	s.EcuId = &v
	return s
}

type DeleteEcuRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeleteEcuQuery    `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DeleteEcuRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEcuRequest) GoString() string {
	return s.String()
}

func (s *DeleteEcuRequest) SetHeaders(v map[string]*string) *DeleteEcuRequest {
	s.Headers = v
	return s
}

func (s *DeleteEcuRequest) SetQuery(v *DeleteEcuQuery) *DeleteEcuRequest {
	s.Query = v
	return s
}

type DeleteEcuResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteEcuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEcuResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEcuResponseBody) SetCode(v int) *DeleteEcuResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEcuResponseBody) SetMessage(v string) *DeleteEcuResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEcuResponseBody) SetData(v string) *DeleteEcuResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteEcuResponseBody) SetRequestId(v string) *DeleteEcuResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEcuResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEcuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEcuResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEcuResponse) GoString() string {
	return s.String()
}

func (s *DeleteEcuResponse) SetHeaders(v map[string]*string) *DeleteEcuResponse {
	s.Headers = v
	return s
}

func (s *DeleteEcuResponse) SetBody(v *DeleteEcuResponseBody) *DeleteEcuResponse {
	s.Body = v
	return s
}

type DeleteDeployGroupQuery struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty" require:"true"`
}

func (s DeleteDeployGroupQuery) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeployGroupQuery) GoString() string {
	return s.String()
}

func (s *DeleteDeployGroupQuery) SetAppId(v string) *DeleteDeployGroupQuery {
	s.AppId = &v
	return s
}

func (s *DeleteDeployGroupQuery) SetGroupName(v string) *DeleteDeployGroupQuery {
	s.GroupName = &v
	return s
}

type DeleteDeployGroupRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeleteDeployGroupQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DeleteDeployGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeployGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeployGroupRequest) SetHeaders(v map[string]*string) *DeleteDeployGroupRequest {
	s.Headers = v
	return s
}

func (s *DeleteDeployGroupRequest) SetQuery(v *DeleteDeployGroupQuery) *DeleteDeployGroupRequest {
	s.Query = v
	return s
}

type DeleteDeployGroupResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteDeployGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeployGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeployGroupResponseBody) SetCode(v int) *DeleteDeployGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDeployGroupResponseBody) SetMessage(v string) *DeleteDeployGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDeployGroupResponseBody) SetData(v string) *DeleteDeployGroupResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDeployGroupResponseBody) SetRequestId(v string) *DeleteDeployGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDeployGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDeployGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeployGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeployGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeployGroupResponse) SetHeaders(v map[string]*string) *DeleteDeployGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeployGroupResponse) SetBody(v *DeleteDeployGroupResponseBody) *DeleteDeployGroupResponse {
	s.Body = v
	return s
}

type DeleteDegradeControlQuery struct {
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty" require:"true"`
}

func (s DeleteDegradeControlQuery) String() string {
	return tea.Prettify(s)
}

func (s DeleteDegradeControlQuery) GoString() string {
	return s.String()
}

func (s *DeleteDegradeControlQuery) SetAppId(v string) *DeleteDegradeControlQuery {
	s.AppId = &v
	return s
}

func (s *DeleteDegradeControlQuery) SetRuleId(v string) *DeleteDegradeControlQuery {
	s.RuleId = &v
	return s
}

type DeleteDegradeControlRequest struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeleteDegradeControlQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DeleteDegradeControlRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDegradeControlRequest) GoString() string {
	return s.String()
}

func (s *DeleteDegradeControlRequest) SetHeaders(v map[string]*string) *DeleteDegradeControlRequest {
	s.Headers = v
	return s
}

func (s *DeleteDegradeControlRequest) SetQuery(v *DeleteDegradeControlQuery) *DeleteDegradeControlRequest {
	s.Query = v
	return s
}

type DeleteDegradeControlResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteDegradeControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDegradeControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDegradeControlResponseBody) SetCode(v int) *DeleteDegradeControlResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDegradeControlResponseBody) SetMessage(v string) *DeleteDegradeControlResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDegradeControlResponseBody) SetRequestId(v string) *DeleteDegradeControlResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDegradeControlResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDegradeControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDegradeControlResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDegradeControlResponse) GoString() string {
	return s.String()
}

func (s *DeleteDegradeControlResponse) SetHeaders(v map[string]*string) *DeleteDegradeControlResponse {
	s.Headers = v
	return s
}

func (s *DeleteDegradeControlResponse) SetBody(v *DeleteDegradeControlResponseBody) *DeleteDegradeControlResponse {
	s.Body = v
	return s
}

type DeleteConfigCenterQuery struct {
	DataId          *string `json:"DataId,omitempty" xml:"DataId,omitempty" require:"true"`
	Group           *string `json:"Group,omitempty" xml:"Group,omitempty" require:"true"`
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty" require:"true"`
}

func (s DeleteConfigCenterQuery) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigCenterQuery) GoString() string {
	return s.String()
}

func (s *DeleteConfigCenterQuery) SetDataId(v string) *DeleteConfigCenterQuery {
	s.DataId = &v
	return s
}

func (s *DeleteConfigCenterQuery) SetGroup(v string) *DeleteConfigCenterQuery {
	s.Group = &v
	return s
}

func (s *DeleteConfigCenterQuery) SetLogicalRegionId(v string) *DeleteConfigCenterQuery {
	s.LogicalRegionId = &v
	return s
}

type DeleteConfigCenterRequest struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeleteConfigCenterQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DeleteConfigCenterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigCenterRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigCenterRequest) SetHeaders(v map[string]*string) *DeleteConfigCenterRequest {
	s.Headers = v
	return s
}

func (s *DeleteConfigCenterRequest) SetQuery(v *DeleteConfigCenterQuery) *DeleteConfigCenterRequest {
	s.Query = v
	return s
}

type DeleteConfigCenterResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteConfigCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigCenterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigCenterResponseBody) SetCode(v int) *DeleteConfigCenterResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConfigCenterResponseBody) SetMessage(v string) *DeleteConfigCenterResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConfigCenterResponseBody) SetRequestId(v string) *DeleteConfigCenterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteConfigCenterResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteConfigCenterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConfigCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigCenterResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigCenterResponse) SetHeaders(v map[string]*string) *DeleteConfigCenterResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigCenterResponse) SetBody(v *DeleteConfigCenterResponseBody) *DeleteConfigCenterResponse {
	s.Body = v
	return s
}

type DeleteClusterMemberQuery struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	ClusterMemberId *string `json:"ClusterMemberId,omitempty" xml:"ClusterMemberId,omitempty" require:"true"`
}

func (s DeleteClusterMemberQuery) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterMemberQuery) GoString() string {
	return s.String()
}

func (s *DeleteClusterMemberQuery) SetClusterId(v string) *DeleteClusterMemberQuery {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterMemberQuery) SetClusterMemberId(v string) *DeleteClusterMemberQuery {
	s.ClusterMemberId = &v
	return s
}

type DeleteClusterMemberRequest struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeleteClusterMemberQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DeleteClusterMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterMemberRequest) SetHeaders(v map[string]*string) *DeleteClusterMemberRequest {
	s.Headers = v
	return s
}

func (s *DeleteClusterMemberRequest) SetQuery(v *DeleteClusterMemberQuery) *DeleteClusterMemberRequest {
	s.Query = v
	return s
}

type DeleteClusterMemberResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteClusterMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterMemberResponseBody) SetCode(v int) *DeleteClusterMemberResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteClusterMemberResponseBody) SetMessage(v string) *DeleteClusterMemberResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteClusterMemberResponseBody) SetData(v bool) *DeleteClusterMemberResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteClusterMemberResponseBody) SetRequestId(v string) *DeleteClusterMemberResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterMemberResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteClusterMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteClusterMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterMemberResponse) SetHeaders(v map[string]*string) *DeleteClusterMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterMemberResponse) SetBody(v *DeleteClusterMemberResponseBody) *DeleteClusterMemberResponse {
	s.Body = v
	return s
}

type DeleteClusterQuery struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	Mode      *int    `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s DeleteClusterQuery) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterQuery) GoString() string {
	return s.String()
}

func (s *DeleteClusterQuery) SetClusterId(v string) *DeleteClusterQuery {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterQuery) SetMode(v int) *DeleteClusterQuery {
	s.Mode = &v
	return s
}

type DeleteClusterRequest struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeleteClusterQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetHeaders(v map[string]*string) *DeleteClusterRequest {
	s.Headers = v
	return s
}

func (s *DeleteClusterRequest) SetQuery(v *DeleteClusterQuery) *DeleteClusterRequest {
	s.Query = v
	return s
}

type DeleteClusterResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetCode(v int) *DeleteClusterResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteClusterResponseBody) SetMessage(v string) *DeleteClusterResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteClusterResponseBody) SetData(v bool) *DeleteClusterResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type BindSlbQuery struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	SlbId          *string `json:"SlbId,omitempty" xml:"SlbId,omitempty" require:"true"`
	SlbIp          *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty" require:"true"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	ListenerPort   *int    `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s BindSlbQuery) String() string {
	return tea.Prettify(s)
}

func (s BindSlbQuery) GoString() string {
	return s.String()
}

func (s *BindSlbQuery) SetAppId(v string) *BindSlbQuery {
	s.AppId = &v
	return s
}

func (s *BindSlbQuery) SetSlbId(v string) *BindSlbQuery {
	s.SlbId = &v
	return s
}

func (s *BindSlbQuery) SetSlbIp(v string) *BindSlbQuery {
	s.SlbIp = &v
	return s
}

func (s *BindSlbQuery) SetType(v string) *BindSlbQuery {
	s.Type = &v
	return s
}

func (s *BindSlbQuery) SetListenerPort(v int) *BindSlbQuery {
	s.ListenerPort = &v
	return s
}

func (s *BindSlbQuery) SetVServerGroupId(v string) *BindSlbQuery {
	s.VServerGroupId = &v
	return s
}

type BindSlbRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *BindSlbQuery      `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s BindSlbRequest) String() string {
	return tea.Prettify(s)
}

func (s BindSlbRequest) GoString() string {
	return s.String()
}

func (s *BindSlbRequest) SetHeaders(v map[string]*string) *BindSlbRequest {
	s.Headers = v
	return s
}

func (s *BindSlbRequest) SetQuery(v *BindSlbQuery) *BindSlbRequest {
	s.Query = v
	return s
}

type BindSlbResponseBody struct {
	Code      *int                     `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *BindSlbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s BindSlbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindSlbResponseBody) GoString() string {
	return s.String()
}

func (s *BindSlbResponseBody) SetCode(v int) *BindSlbResponseBody {
	s.Code = &v
	return s
}

func (s *BindSlbResponseBody) SetMessage(v string) *BindSlbResponseBody {
	s.Message = &v
	return s
}

func (s *BindSlbResponseBody) SetRequestId(v string) *BindSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindSlbResponseBody) SetData(v *BindSlbResponseBodyData) *BindSlbResponseBody {
	s.Data = v
	return s
}

type BindSlbResponseBodyData struct {
	ExtSlbId          *string `json:"ExtSlbId,omitempty" xml:"ExtSlbId,omitempty" require:"true"`
	ExtSlbIp          *string `json:"ExtSlbIp,omitempty" xml:"ExtSlbIp,omitempty" require:"true"`
	ExtSlbName        *string `json:"ExtSlbName,omitempty" xml:"ExtSlbName,omitempty" require:"true"`
	ExtVServerGroupId *string `json:"ExtVServerGroupId,omitempty" xml:"ExtVServerGroupId,omitempty" require:"true"`
	SlbId             *string `json:"SlbId,omitempty" xml:"SlbId,omitempty" require:"true"`
	SlbIp             *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty" require:"true"`
	SlbName           *string `json:"SlbName,omitempty" xml:"SlbName,omitempty" require:"true"`
	SlbPort           *int    `json:"SlbPort,omitempty" xml:"SlbPort,omitempty" require:"true"`
	VServerGroupId    *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty" require:"true"`
}

func (s BindSlbResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BindSlbResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindSlbResponseBodyData) SetExtSlbId(v string) *BindSlbResponseBodyData {
	s.ExtSlbId = &v
	return s
}

func (s *BindSlbResponseBodyData) SetExtSlbIp(v string) *BindSlbResponseBodyData {
	s.ExtSlbIp = &v
	return s
}

func (s *BindSlbResponseBodyData) SetExtSlbName(v string) *BindSlbResponseBodyData {
	s.ExtSlbName = &v
	return s
}

func (s *BindSlbResponseBodyData) SetExtVServerGroupId(v string) *BindSlbResponseBodyData {
	s.ExtVServerGroupId = &v
	return s
}

func (s *BindSlbResponseBodyData) SetSlbId(v string) *BindSlbResponseBodyData {
	s.SlbId = &v
	return s
}

func (s *BindSlbResponseBodyData) SetSlbIp(v string) *BindSlbResponseBodyData {
	s.SlbIp = &v
	return s
}

func (s *BindSlbResponseBodyData) SetSlbName(v string) *BindSlbResponseBodyData {
	s.SlbName = &v
	return s
}

func (s *BindSlbResponseBodyData) SetSlbPort(v int) *BindSlbResponseBodyData {
	s.SlbPort = &v
	return s
}

func (s *BindSlbResponseBodyData) SetVServerGroupId(v string) *BindSlbResponseBodyData {
	s.VServerGroupId = &v
	return s
}

type BindSlbResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindSlbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindSlbResponse) String() string {
	return tea.Prettify(s)
}

func (s BindSlbResponse) GoString() string {
	return s.String()
}

func (s *BindSlbResponse) SetHeaders(v map[string]*string) *BindSlbResponse {
	s.Headers = v
	return s
}

func (s *BindSlbResponse) SetBody(v *BindSlbResponseBody) *BindSlbResponse {
	s.Body = v
	return s
}

type AuthorizeRoleQuery struct {
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty" require:"true"`
	RoleIds      *string `json:"RoleIds,omitempty" xml:"RoleIds,omitempty" require:"true"`
}

func (s AuthorizeRoleQuery) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeRoleQuery) GoString() string {
	return s.String()
}

func (s *AuthorizeRoleQuery) SetTargetUserId(v string) *AuthorizeRoleQuery {
	s.TargetUserId = &v
	return s
}

func (s *AuthorizeRoleQuery) SetRoleIds(v string) *AuthorizeRoleQuery {
	s.RoleIds = &v
	return s
}

type AuthorizeRoleRequest struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *AuthorizeRoleQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s AuthorizeRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeRoleRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeRoleRequest) SetHeaders(v map[string]*string) *AuthorizeRoleRequest {
	s.Headers = v
	return s
}

func (s *AuthorizeRoleRequest) SetQuery(v *AuthorizeRoleQuery) *AuthorizeRoleRequest {
	s.Query = v
	return s
}

type AuthorizeRoleResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AuthorizeRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeRoleResponseBody) SetCode(v int) *AuthorizeRoleResponseBody {
	s.Code = &v
	return s
}

func (s *AuthorizeRoleResponseBody) SetMessage(v string) *AuthorizeRoleResponseBody {
	s.Message = &v
	return s
}

func (s *AuthorizeRoleResponseBody) SetRequestId(v string) *AuthorizeRoleResponseBody {
	s.RequestId = &v
	return s
}

type AuthorizeRoleResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AuthorizeRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthorizeRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeRoleResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeRoleResponse) SetHeaders(v map[string]*string) *AuthorizeRoleResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeRoleResponse) SetBody(v *AuthorizeRoleResponseBody) *AuthorizeRoleResponse {
	s.Body = v
	return s
}

type AuthorizeResourceGroupQuery struct {
	TargetUserId     *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty" require:"true"`
	ResourceGroupIds *string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" require:"true"`
}

func (s AuthorizeResourceGroupQuery) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeResourceGroupQuery) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceGroupQuery) SetTargetUserId(v string) *AuthorizeResourceGroupQuery {
	s.TargetUserId = &v
	return s
}

func (s *AuthorizeResourceGroupQuery) SetResourceGroupIds(v string) *AuthorizeResourceGroupQuery {
	s.ResourceGroupIds = &v
	return s
}

type AuthorizeResourceGroupRequest struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *AuthorizeResourceGroupQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s AuthorizeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceGroupRequest) SetHeaders(v map[string]*string) *AuthorizeResourceGroupRequest {
	s.Headers = v
	return s
}

func (s *AuthorizeResourceGroupRequest) SetQuery(v *AuthorizeResourceGroupQuery) *AuthorizeResourceGroupRequest {
	s.Query = v
	return s
}

type AuthorizeResourceGroupResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AuthorizeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceGroupResponseBody) SetCode(v int) *AuthorizeResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AuthorizeResourceGroupResponseBody) SetMessage(v string) *AuthorizeResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AuthorizeResourceGroupResponseBody) SetRequestId(v string) *AuthorizeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type AuthorizeResourceGroupResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AuthorizeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthorizeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceGroupResponse) SetHeaders(v map[string]*string) *AuthorizeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeResourceGroupResponse) SetBody(v *AuthorizeResourceGroupResponseBody) *AuthorizeResourceGroupResponse {
	s.Body = v
	return s
}

type AuthorizeApplicationQuery struct {
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty" require:"true"`
	AppIds       *string `json:"AppIds,omitempty" xml:"AppIds,omitempty" require:"true"`
}

func (s AuthorizeApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeApplicationQuery) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationQuery) SetTargetUserId(v string) *AuthorizeApplicationQuery {
	s.TargetUserId = &v
	return s
}

func (s *AuthorizeApplicationQuery) SetAppIds(v string) *AuthorizeApplicationQuery {
	s.AppIds = &v
	return s
}

type AuthorizeApplicationRequest struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *AuthorizeApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s AuthorizeApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeApplicationRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationRequest) SetHeaders(v map[string]*string) *AuthorizeApplicationRequest {
	s.Headers = v
	return s
}

func (s *AuthorizeApplicationRequest) SetQuery(v *AuthorizeApplicationQuery) *AuthorizeApplicationRequest {
	s.Query = v
	return s
}

type AuthorizeApplicationResponseBody struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AuthorizeApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationResponseBody) SetCode(v int) *AuthorizeApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *AuthorizeApplicationResponseBody) SetMessage(v string) *AuthorizeApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *AuthorizeApplicationResponseBody) SetRequestId(v string) *AuthorizeApplicationResponseBody {
	s.RequestId = &v
	return s
}

type AuthorizeApplicationResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AuthorizeApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthorizeApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeApplicationResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationResponse) SetHeaders(v map[string]*string) *AuthorizeApplicationResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeApplicationResponse) SetBody(v *AuthorizeApplicationResponseBody) *AuthorizeApplicationResponse {
	s.Body = v
	return s
}

type ListVpcRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ListVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpcRequest) GoString() string {
	return s.String()
}

func (s *ListVpcRequest) SetHeaders(v map[string]*string) *ListVpcRequest {
	s.Headers = v
	return s
}

type ListVpcResponseBody struct {
	Code      *int                        `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	VpcList   *ListVpcResponseBodyVpcList `json:"VpcList,omitempty" xml:"VpcList,omitempty" require:"true" type:"Struct"`
}

func (s ListVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcResponseBody) SetCode(v int) *ListVpcResponseBody {
	s.Code = &v
	return s
}

func (s *ListVpcResponseBody) SetMessage(v string) *ListVpcResponseBody {
	s.Message = &v
	return s
}

func (s *ListVpcResponseBody) SetRequestId(v string) *ListVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcResponseBody) SetVpcList(v *ListVpcResponseBodyVpcList) *ListVpcResponseBody {
	s.VpcList = v
	return s
}

type ListVpcResponseBodyVpcList struct {
	VpcEntity []*ListVpcResponseBodyVpcListVpcEntity `json:"VpcEntity,omitempty" xml:"VpcEntity,omitempty" require:"true" type:"Repeated"`
}

func (s ListVpcResponseBodyVpcList) String() string {
	return tea.Prettify(s)
}

func (s ListVpcResponseBodyVpcList) GoString() string {
	return s.String()
}

func (s *ListVpcResponseBodyVpcList) SetVpcEntity(v []*ListVpcResponseBodyVpcListVpcEntity) *ListVpcResponseBodyVpcList {
	s.VpcEntity = v
	return s
}

type ListVpcResponseBodyVpcListVpcEntity struct {
	VpcId    *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	VpcName  *string `json:"VpcName,omitempty" xml:"VpcName,omitempty" require:"true"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	Expired  *bool   `json:"Expired,omitempty" xml:"Expired,omitempty" require:"true"`
	EcsNum   *int    `json:"EcsNum,omitempty" xml:"EcsNum,omitempty" require:"true"`
}

func (s ListVpcResponseBodyVpcListVpcEntity) String() string {
	return tea.Prettify(s)
}

func (s ListVpcResponseBodyVpcListVpcEntity) GoString() string {
	return s.String()
}

func (s *ListVpcResponseBodyVpcListVpcEntity) SetVpcId(v string) *ListVpcResponseBodyVpcListVpcEntity {
	s.VpcId = &v
	return s
}

func (s *ListVpcResponseBodyVpcListVpcEntity) SetVpcName(v string) *ListVpcResponseBodyVpcListVpcEntity {
	s.VpcName = &v
	return s
}

func (s *ListVpcResponseBodyVpcListVpcEntity) SetRegionId(v string) *ListVpcResponseBodyVpcListVpcEntity {
	s.RegionId = &v
	return s
}

func (s *ListVpcResponseBodyVpcListVpcEntity) SetUserId(v string) *ListVpcResponseBodyVpcListVpcEntity {
	s.UserId = &v
	return s
}

func (s *ListVpcResponseBodyVpcListVpcEntity) SetExpired(v bool) *ListVpcResponseBodyVpcListVpcEntity {
	s.Expired = &v
	return s
}

func (s *ListVpcResponseBodyVpcListVpcEntity) SetEcsNum(v int) *ListVpcResponseBodyVpcListVpcEntity {
	s.EcsNum = &v
	return s
}

type ListVpcResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVpcResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcResponse) GoString() string {
	return s.String()
}

func (s *ListVpcResponse) SetHeaders(v map[string]*string) *ListVpcResponse {
	s.Headers = v
	return s
}

func (s *ListVpcResponse) SetBody(v *ListVpcResponseBody) *ListVpcResponse {
	s.Body = v
	return s
}

type ScaleOutApplicationQuery struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	EcuInfo     *string `json:"EcuInfo,omitempty" xml:"EcuInfo,omitempty" require:"true"`
	DeployGroup *string `json:"DeployGroup,omitempty" xml:"DeployGroup,omitempty" require:"true"`
}

func (s ScaleOutApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutApplicationQuery) GoString() string {
	return s.String()
}

func (s *ScaleOutApplicationQuery) SetAppId(v string) *ScaleOutApplicationQuery {
	s.AppId = &v
	return s
}

func (s *ScaleOutApplicationQuery) SetEcuInfo(v string) *ScaleOutApplicationQuery {
	s.EcuInfo = &v
	return s
}

func (s *ScaleOutApplicationQuery) SetDeployGroup(v string) *ScaleOutApplicationQuery {
	s.DeployGroup = &v
	return s
}

type ScaleOutApplicationRequest struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ScaleOutApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ScaleOutApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutApplicationRequest) GoString() string {
	return s.String()
}

func (s *ScaleOutApplicationRequest) SetHeaders(v map[string]*string) *ScaleOutApplicationRequest {
	s.Headers = v
	return s
}

func (s *ScaleOutApplicationRequest) SetQuery(v *ScaleOutApplicationQuery) *ScaleOutApplicationRequest {
	s.Query = v
	return s
}

type ScaleOutApplicationResponseBody struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ScaleOutApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleOutApplicationResponseBody) SetChangeOrderId(v string) *ScaleOutApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *ScaleOutApplicationResponseBody) SetCode(v int) *ScaleOutApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ScaleOutApplicationResponseBody) SetMessage(v string) *ScaleOutApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ScaleOutApplicationResponseBody) SetRequestId(v string) *ScaleOutApplicationResponseBody {
	s.RequestId = &v
	return s
}

type ScaleOutApplicationResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ScaleOutApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScaleOutApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleOutApplicationResponse) GoString() string {
	return s.String()
}

func (s *ScaleOutApplicationResponse) SetHeaders(v map[string]*string) *ScaleOutApplicationResponse {
	s.Headers = v
	return s
}

func (s *ScaleOutApplicationResponse) SetBody(v *ScaleOutApplicationResponseBody) *ScaleOutApplicationResponse {
	s.Body = v
	return s
}

type ScaleInApplicationQuery struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	EccInfo     *string `json:"EccInfo,omitempty" xml:"EccInfo,omitempty" require:"true"`
	ForceStatus *bool   `json:"ForceStatus,omitempty" xml:"ForceStatus,omitempty"`
}

func (s ScaleInApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s ScaleInApplicationQuery) GoString() string {
	return s.String()
}

func (s *ScaleInApplicationQuery) SetAppId(v string) *ScaleInApplicationQuery {
	s.AppId = &v
	return s
}

func (s *ScaleInApplicationQuery) SetEccInfo(v string) *ScaleInApplicationQuery {
	s.EccInfo = &v
	return s
}

func (s *ScaleInApplicationQuery) SetForceStatus(v bool) *ScaleInApplicationQuery {
	s.ForceStatus = &v
	return s
}

type ScaleInApplicationRequest struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ScaleInApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ScaleInApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleInApplicationRequest) GoString() string {
	return s.String()
}

func (s *ScaleInApplicationRequest) SetHeaders(v map[string]*string) *ScaleInApplicationRequest {
	s.Headers = v
	return s
}

func (s *ScaleInApplicationRequest) SetQuery(v *ScaleInApplicationQuery) *ScaleInApplicationRequest {
	s.Query = v
	return s
}

type ScaleInApplicationResponseBody struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s ScaleInApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleInApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleInApplicationResponseBody) SetChangeOrderId(v string) *ScaleInApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *ScaleInApplicationResponseBody) SetCode(v int) *ScaleInApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ScaleInApplicationResponseBody) SetMessage(v string) *ScaleInApplicationResponseBody {
	s.Message = &v
	return s
}

type ScaleInApplicationResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ScaleInApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScaleInApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleInApplicationResponse) GoString() string {
	return s.String()
}

func (s *ScaleInApplicationResponse) SetHeaders(v map[string]*string) *ScaleInApplicationResponse {
	s.Headers = v
	return s
}

func (s *ScaleInApplicationResponse) SetBody(v *ScaleInApplicationResponseBody) *ScaleInApplicationResponse {
	s.Body = v
	return s
}

type DeployApplicationQuery struct {
	AppId                  *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	PackageVersion         *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty" require:"true"`
	Desc                   *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	DeployType             *string `json:"DeployType,omitempty" xml:"DeployType,omitempty" require:"true"`
	WarUrl                 *string `json:"WarUrl,omitempty" xml:"WarUrl,omitempty"`
	ImageUrl               *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	GroupId                *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	Batch                  *int    `json:"Batch,omitempty" xml:"Batch,omitempty"`
	BatchWaitTime          *int    `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	AppEnv                 *string `json:"AppEnv,omitempty" xml:"AppEnv,omitempty"`
	BuildPackId            *int64  `json:"BuildPackId,omitempty" xml:"BuildPackId,omitempty"`
	ComponentIds           *string `json:"ComponentIds,omitempty" xml:"ComponentIds,omitempty"`
	ReleaseType            *int64  `json:"ReleaseType,omitempty" xml:"ReleaseType,omitempty"`
	Gray                   *bool   `json:"Gray,omitempty" xml:"Gray,omitempty"`
	TrafficControlStrategy *string `json:"TrafficControlStrategy,omitempty" xml:"TrafficControlStrategy,omitempty"`
}

func (s DeployApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s DeployApplicationQuery) GoString() string {
	return s.String()
}

func (s *DeployApplicationQuery) SetAppId(v string) *DeployApplicationQuery {
	s.AppId = &v
	return s
}

func (s *DeployApplicationQuery) SetPackageVersion(v string) *DeployApplicationQuery {
	s.PackageVersion = &v
	return s
}

func (s *DeployApplicationQuery) SetDesc(v string) *DeployApplicationQuery {
	s.Desc = &v
	return s
}

func (s *DeployApplicationQuery) SetDeployType(v string) *DeployApplicationQuery {
	s.DeployType = &v
	return s
}

func (s *DeployApplicationQuery) SetWarUrl(v string) *DeployApplicationQuery {
	s.WarUrl = &v
	return s
}

func (s *DeployApplicationQuery) SetImageUrl(v string) *DeployApplicationQuery {
	s.ImageUrl = &v
	return s
}

func (s *DeployApplicationQuery) SetGroupId(v string) *DeployApplicationQuery {
	s.GroupId = &v
	return s
}

func (s *DeployApplicationQuery) SetBatch(v int) *DeployApplicationQuery {
	s.Batch = &v
	return s
}

func (s *DeployApplicationQuery) SetBatchWaitTime(v int) *DeployApplicationQuery {
	s.BatchWaitTime = &v
	return s
}

func (s *DeployApplicationQuery) SetAppEnv(v string) *DeployApplicationQuery {
	s.AppEnv = &v
	return s
}

func (s *DeployApplicationQuery) SetBuildPackId(v int64) *DeployApplicationQuery {
	s.BuildPackId = &v
	return s
}

func (s *DeployApplicationQuery) SetComponentIds(v string) *DeployApplicationQuery {
	s.ComponentIds = &v
	return s
}

func (s *DeployApplicationQuery) SetReleaseType(v int64) *DeployApplicationQuery {
	s.ReleaseType = &v
	return s
}

func (s *DeployApplicationQuery) SetGray(v bool) *DeployApplicationQuery {
	s.Gray = &v
	return s
}

func (s *DeployApplicationQuery) SetTrafficControlStrategy(v string) *DeployApplicationQuery {
	s.TrafficControlStrategy = &v
	return s
}

type DeployApplicationRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeployApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DeployApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeployApplicationRequest) SetHeaders(v map[string]*string) *DeployApplicationRequest {
	s.Headers = v
	return s
}

func (s *DeployApplicationRequest) SetQuery(v *DeployApplicationQuery) *DeployApplicationRequest {
	s.Query = v
	return s
}

type DeployApplicationResponseBody struct {
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeployApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeployApplicationResponseBody) SetCode(v int) *DeployApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeployApplicationResponseBody) SetMessage(v string) *DeployApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeployApplicationResponseBody) SetChangeOrderId(v string) *DeployApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *DeployApplicationResponseBody) SetRequestId(v string) *DeployApplicationResponseBody {
	s.RequestId = &v
	return s
}

type DeployApplicationResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeployApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeployApplicationResponse) SetHeaders(v map[string]*string) *DeployApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeployApplicationResponse) SetBody(v *DeployApplicationResponseBody) *DeployApplicationResponse {
	s.Body = v
	return s
}

type InsertApplicationQuery struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	BuildPackId     *int    `json:"BuildPackId,omitempty" xml:"BuildPackId,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty" require:"true"`
	ReservedPortStr *string `json:"ReservedPortStr,omitempty" xml:"ReservedPortStr,omitempty"`
	EcuInfo         *string `json:"EcuInfo,omitempty" xml:"EcuInfo,omitempty"`
	Cpu             *int    `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Mem             *int    `json:"Mem,omitempty" xml:"Mem,omitempty"`
	HealthCheckURL  *string `json:"HealthCheckURL,omitempty" xml:"HealthCheckURL,omitempty"`
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	Jdk             *string `json:"Jdk,omitempty" xml:"Jdk,omitempty"`
	WebContainer    *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
	PackageType     *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	ComponentIds    *string `json:"ComponentIds,omitempty" xml:"ComponentIds,omitempty"`
}

func (s InsertApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s InsertApplicationQuery) GoString() string {
	return s.String()
}

func (s *InsertApplicationQuery) SetClusterId(v string) *InsertApplicationQuery {
	s.ClusterId = &v
	return s
}

func (s *InsertApplicationQuery) SetBuildPackId(v int) *InsertApplicationQuery {
	s.BuildPackId = &v
	return s
}

func (s *InsertApplicationQuery) SetDescription(v string) *InsertApplicationQuery {
	s.Description = &v
	return s
}

func (s *InsertApplicationQuery) SetApplicationName(v string) *InsertApplicationQuery {
	s.ApplicationName = &v
	return s
}

func (s *InsertApplicationQuery) SetReservedPortStr(v string) *InsertApplicationQuery {
	s.ReservedPortStr = &v
	return s
}

func (s *InsertApplicationQuery) SetEcuInfo(v string) *InsertApplicationQuery {
	s.EcuInfo = &v
	return s
}

func (s *InsertApplicationQuery) SetCpu(v int) *InsertApplicationQuery {
	s.Cpu = &v
	return s
}

func (s *InsertApplicationQuery) SetMem(v int) *InsertApplicationQuery {
	s.Mem = &v
	return s
}

func (s *InsertApplicationQuery) SetHealthCheckURL(v string) *InsertApplicationQuery {
	s.HealthCheckURL = &v
	return s
}

func (s *InsertApplicationQuery) SetLogicalRegionId(v string) *InsertApplicationQuery {
	s.LogicalRegionId = &v
	return s
}

func (s *InsertApplicationQuery) SetJdk(v string) *InsertApplicationQuery {
	s.Jdk = &v
	return s
}

func (s *InsertApplicationQuery) SetWebContainer(v string) *InsertApplicationQuery {
	s.WebContainer = &v
	return s
}

func (s *InsertApplicationQuery) SetPackageType(v string) *InsertApplicationQuery {
	s.PackageType = &v
	return s
}

func (s *InsertApplicationQuery) SetComponentIds(v string) *InsertApplicationQuery {
	s.ComponentIds = &v
	return s
}

type InsertApplicationRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *InsertApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s InsertApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertApplicationRequest) GoString() string {
	return s.String()
}

func (s *InsertApplicationRequest) SetHeaders(v map[string]*string) *InsertApplicationRequest {
	s.Headers = v
	return s
}

func (s *InsertApplicationRequest) SetQuery(v *InsertApplicationQuery) *InsertApplicationRequest {
	s.Query = v
	return s
}

type InsertApplicationResponseBody struct {
	Code            *int                                          `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message         *string                                       `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId       *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ApplicationInfo *InsertApplicationResponseBodyApplicationInfo `json:"ApplicationInfo,omitempty" xml:"ApplicationInfo,omitempty" require:"true" type:"Struct"`
}

func (s InsertApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *InsertApplicationResponseBody) SetCode(v int) *InsertApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *InsertApplicationResponseBody) SetMessage(v string) *InsertApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *InsertApplicationResponseBody) SetRequestId(v string) *InsertApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertApplicationResponseBody) SetApplicationInfo(v *InsertApplicationResponseBodyApplicationInfo) *InsertApplicationResponseBody {
	s.ApplicationInfo = v
	return s
}

type InsertApplicationResponseBodyApplicationInfo struct {
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	Owner         *string `json:"Owner,omitempty" xml:"Owner,omitempty" require:"true"`
	Dockerize     *bool   `json:"Dockerize,omitempty" xml:"Dockerize,omitempty" require:"true"`
	Port          *int    `json:"Port,omitempty" xml:"Port,omitempty" require:"true"`
	RegionName    *string `json:"RegionName,omitempty" xml:"RegionName,omitempty" require:"true"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s InsertApplicationResponseBodyApplicationInfo) String() string {
	return tea.Prettify(s)
}

func (s InsertApplicationResponseBodyApplicationInfo) GoString() string {
	return s.String()
}

func (s *InsertApplicationResponseBodyApplicationInfo) SetAppName(v string) *InsertApplicationResponseBodyApplicationInfo {
	s.AppName = &v
	return s
}

func (s *InsertApplicationResponseBodyApplicationInfo) SetAppId(v string) *InsertApplicationResponseBodyApplicationInfo {
	s.AppId = &v
	return s
}

func (s *InsertApplicationResponseBodyApplicationInfo) SetUserId(v string) *InsertApplicationResponseBodyApplicationInfo {
	s.UserId = &v
	return s
}

func (s *InsertApplicationResponseBodyApplicationInfo) SetOwner(v string) *InsertApplicationResponseBodyApplicationInfo {
	s.Owner = &v
	return s
}

func (s *InsertApplicationResponseBodyApplicationInfo) SetDockerize(v bool) *InsertApplicationResponseBodyApplicationInfo {
	s.Dockerize = &v
	return s
}

func (s *InsertApplicationResponseBodyApplicationInfo) SetPort(v int) *InsertApplicationResponseBodyApplicationInfo {
	s.Port = &v
	return s
}

func (s *InsertApplicationResponseBodyApplicationInfo) SetRegionName(v string) *InsertApplicationResponseBodyApplicationInfo {
	s.RegionName = &v
	return s
}

func (s *InsertApplicationResponseBodyApplicationInfo) SetChangeOrderId(v string) *InsertApplicationResponseBodyApplicationInfo {
	s.ChangeOrderId = &v
	return s
}

type InsertApplicationResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertApplicationResponse) GoString() string {
	return s.String()
}

func (s *InsertApplicationResponse) SetHeaders(v map[string]*string) *InsertApplicationResponse {
	s.Headers = v
	return s
}

func (s *InsertApplicationResponse) SetBody(v *InsertApplicationResponseBody) *InsertApplicationResponse {
	s.Body = v
	return s
}

type DeleteApplicationQuery struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s DeleteApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationQuery) GoString() string {
	return s.String()
}

func (s *DeleteApplicationQuery) SetAppId(v string) *DeleteApplicationQuery {
	s.AppId = &v
	return s
}

type DeleteApplicationRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeleteApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s DeleteApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRequest) SetHeaders(v map[string]*string) *DeleteApplicationRequest {
	s.Headers = v
	return s
}

func (s *DeleteApplicationRequest) SetQuery(v *DeleteApplicationQuery) *DeleteApplicationRequest {
	s.Query = v
	return s
}

type DeleteApplicationResponseBody struct {
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponseBody) SetCode(v int) *DeleteApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetMessage(v string) *DeleteApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetChangeOrderId(v string) *DeleteApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetRequestId(v string) *DeleteApplicationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApplicationResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponse) SetHeaders(v map[string]*string) *DeleteApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationResponse) SetBody(v *DeleteApplicationResponseBody) *DeleteApplicationResponse {
	s.Body = v
	return s
}

type UpdateContainerQuery struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	BuildPackId *int    `json:"BuildPackId,omitempty" xml:"BuildPackId,omitempty" require:"true"`
}

func (s UpdateContainerQuery) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerQuery) GoString() string {
	return s.String()
}

func (s *UpdateContainerQuery) SetAppId(v string) *UpdateContainerQuery {
	s.AppId = &v
	return s
}

func (s *UpdateContainerQuery) SetBuildPackId(v int) *UpdateContainerQuery {
	s.BuildPackId = &v
	return s
}

type UpdateContainerRequest struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *UpdateContainerQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s UpdateContainerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerRequest) GoString() string {
	return s.String()
}

func (s *UpdateContainerRequest) SetHeaders(v map[string]*string) *UpdateContainerRequest {
	s.Headers = v
	return s
}

func (s *UpdateContainerRequest) SetQuery(v *UpdateContainerQuery) *UpdateContainerRequest {
	s.Query = v
	return s
}

type UpdateContainerResponseBody struct {
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateContainerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContainerResponseBody) SetCode(v int) *UpdateContainerResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateContainerResponseBody) SetMessage(v string) *UpdateContainerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateContainerResponseBody) SetChangeOrderId(v string) *UpdateContainerResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *UpdateContainerResponseBody) SetRequestId(v string) *UpdateContainerResponseBody {
	s.RequestId = &v
	return s
}

type UpdateContainerResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateContainerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateContainerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerResponse) GoString() string {
	return s.String()
}

func (s *UpdateContainerResponse) SetHeaders(v map[string]*string) *UpdateContainerResponse {
	s.Headers = v
	return s
}

func (s *UpdateContainerResponse) SetBody(v *UpdateContainerResponseBody) *UpdateContainerResponse {
	s.Body = v
	return s
}

type StopApplicationQuery struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	EccInfo *string `json:"EccInfo,omitempty" xml:"EccInfo,omitempty"`
}

func (s StopApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s StopApplicationQuery) GoString() string {
	return s.String()
}

func (s *StopApplicationQuery) SetAppId(v string) *StopApplicationQuery {
	s.AppId = &v
	return s
}

func (s *StopApplicationQuery) SetEccInfo(v string) *StopApplicationQuery {
	s.EccInfo = &v
	return s
}

type StopApplicationRequest struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *StopApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s StopApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s StopApplicationRequest) GoString() string {
	return s.String()
}

func (s *StopApplicationRequest) SetHeaders(v map[string]*string) *StopApplicationRequest {
	s.Headers = v
	return s
}

func (s *StopApplicationRequest) SetQuery(v *StopApplicationQuery) *StopApplicationRequest {
	s.Query = v
	return s
}

type StopApplicationResponseBody struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StopApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *StopApplicationResponseBody) SetChangeOrderId(v string) *StopApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *StopApplicationResponseBody) SetCode(v int) *StopApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *StopApplicationResponseBody) SetMessage(v string) *StopApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *StopApplicationResponseBody) SetRequestId(v string) *StopApplicationResponseBody {
	s.RequestId = &v
	return s
}

type StopApplicationResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s StopApplicationResponse) GoString() string {
	return s.String()
}

func (s *StopApplicationResponse) SetHeaders(v map[string]*string) *StopApplicationResponse {
	s.Headers = v
	return s
}

func (s *StopApplicationResponse) SetBody(v *StopApplicationResponseBody) *StopApplicationResponse {
	s.Body = v
	return s
}

type StartApplicationQuery struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	EccInfo *string `json:"EccInfo,omitempty" xml:"EccInfo,omitempty"`
}

func (s StartApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s StartApplicationQuery) GoString() string {
	return s.String()
}

func (s *StartApplicationQuery) SetAppId(v string) *StartApplicationQuery {
	s.AppId = &v
	return s
}

func (s *StartApplicationQuery) SetEccInfo(v string) *StartApplicationQuery {
	s.EccInfo = &v
	return s
}

type StartApplicationRequest struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *StartApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s StartApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s StartApplicationRequest) GoString() string {
	return s.String()
}

func (s *StartApplicationRequest) SetHeaders(v map[string]*string) *StartApplicationRequest {
	s.Headers = v
	return s
}

func (s *StartApplicationRequest) SetQuery(v *StartApplicationQuery) *StartApplicationRequest {
	s.Query = v
	return s
}

type StartApplicationResponseBody struct {
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s StartApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *StartApplicationResponseBody) SetCode(v int) *StartApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *StartApplicationResponseBody) SetMessage(v string) *StartApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *StartApplicationResponseBody) SetChangeOrderId(v string) *StartApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *StartApplicationResponseBody) SetRequestId(v string) *StartApplicationResponseBody {
	s.RequestId = &v
	return s
}

type StartApplicationResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s StartApplicationResponse) GoString() string {
	return s.String()
}

func (s *StartApplicationResponse) SetHeaders(v map[string]*string) *StartApplicationResponse {
	s.Headers = v
	return s
}

func (s *StartApplicationResponse) SetBody(v *StartApplicationResponseBody) *StartApplicationResponse {
	s.Body = v
	return s
}

type ResetApplicationQuery struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	EccInfo *string `json:"EccInfo,omitempty" xml:"EccInfo,omitempty" require:"true"`
}

func (s ResetApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s ResetApplicationQuery) GoString() string {
	return s.String()
}

func (s *ResetApplicationQuery) SetAppId(v string) *ResetApplicationQuery {
	s.AppId = &v
	return s
}

func (s *ResetApplicationQuery) SetEccInfo(v string) *ResetApplicationQuery {
	s.EccInfo = &v
	return s
}

type ResetApplicationRequest struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ResetApplicationQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ResetApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetApplicationRequest) GoString() string {
	return s.String()
}

func (s *ResetApplicationRequest) SetHeaders(v map[string]*string) *ResetApplicationRequest {
	s.Headers = v
	return s
}

func (s *ResetApplicationRequest) SetQuery(v *ResetApplicationQuery) *ResetApplicationRequest {
	s.Query = v
	return s
}

type ResetApplicationResponseBody struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	Code          *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ResetApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ResetApplicationResponseBody) SetChangeOrderId(v string) *ResetApplicationResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *ResetApplicationResponseBody) SetCode(v int) *ResetApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ResetApplicationResponseBody) SetMessage(v string) *ResetApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ResetApplicationResponseBody) SetRequestId(v string) *ResetApplicationResponseBody {
	s.RequestId = &v
	return s
}

type ResetApplicationResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetApplicationResponse) GoString() string {
	return s.String()
}

func (s *ResetApplicationResponse) SetHeaders(v map[string]*string) *ResetApplicationResponse {
	s.Headers = v
	return s
}

func (s *ResetApplicationResponse) SetBody(v *ResetApplicationResponseBody) *ResetApplicationResponse {
	s.Body = v
	return s
}

type ListDeployGroupQuery struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s ListDeployGroupQuery) String() string {
	return tea.Prettify(s)
}

func (s ListDeployGroupQuery) GoString() string {
	return s.String()
}

func (s *ListDeployGroupQuery) SetAppId(v string) *ListDeployGroupQuery {
	s.AppId = &v
	return s
}

type ListDeployGroupRequest struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListDeployGroupQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s ListDeployGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeployGroupRequest) GoString() string {
	return s.String()
}

func (s *ListDeployGroupRequest) SetHeaders(v map[string]*string) *ListDeployGroupRequest {
	s.Headers = v
	return s
}

func (s *ListDeployGroupRequest) SetQuery(v *ListDeployGroupQuery) *ListDeployGroupRequest {
	s.Query = v
	return s
}

type ListDeployGroupResponseBody struct {
	Code            *int                                        `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message         *string                                     `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId       *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DeployGroupList *ListDeployGroupResponseBodyDeployGroupList `json:"DeployGroupList,omitempty" xml:"DeployGroupList,omitempty" require:"true" type:"Struct"`
}

func (s ListDeployGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeployGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeployGroupResponseBody) SetCode(v int) *ListDeployGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeployGroupResponseBody) SetMessage(v string) *ListDeployGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeployGroupResponseBody) SetRequestId(v string) *ListDeployGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeployGroupResponseBody) SetDeployGroupList(v *ListDeployGroupResponseBodyDeployGroupList) *ListDeployGroupResponseBody {
	s.DeployGroupList = v
	return s
}

type ListDeployGroupResponseBodyDeployGroupList struct {
	DeployGroup []*ListDeployGroupResponseBodyDeployGroupListDeployGroup `json:"DeployGroup,omitempty" xml:"DeployGroup,omitempty" require:"true" type:"Repeated"`
}

func (s ListDeployGroupResponseBodyDeployGroupList) String() string {
	return tea.Prettify(s)
}

func (s ListDeployGroupResponseBodyDeployGroupList) GoString() string {
	return s.String()
}

func (s *ListDeployGroupResponseBodyDeployGroupList) SetDeployGroup(v []*ListDeployGroupResponseBodyDeployGroupListDeployGroup) *ListDeployGroupResponseBodyDeployGroupList {
	s.DeployGroup = v
	return s
}

type ListDeployGroupResponseBodyDeployGroupListDeployGroup struct {
	GroupId               *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
	GroupName             *string `json:"GroupName,omitempty" xml:"GroupName,omitempty" require:"true"`
	AppId                 *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	PackageVersionId      *string `json:"PackageVersionId,omitempty" xml:"PackageVersionId,omitempty" require:"true"`
	AppVersionId          *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty" require:"true"`
	GroupType             *int    `json:"GroupType,omitempty" xml:"GroupType,omitempty" require:"true"`
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime            *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	NameSpace             *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty" require:"true"`
	ClusterName           *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty" require:"true"`
	LastUpdateTime        *int64  `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty" require:"true"`
	PreStop               *string `json:"PreStop,omitempty" xml:"PreStop,omitempty" require:"true"`
	PostStart             *string `json:"PostStart,omitempty" xml:"PostStart,omitempty" require:"true"`
	PackageUrl            *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty" require:"true"`
	Env                   *string `json:"Env,omitempty" xml:"Env,omitempty" require:"true"`
	Labels                *string `json:"Labels,omitempty" xml:"Labels,omitempty" require:"true"`
	Selector              *string `json:"Selector,omitempty" xml:"Selector,omitempty" require:"true"`
	Strategy              *string `json:"Strategy,omitempty" xml:"Strategy,omitempty" require:"true"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	Reversion             *string `json:"Reversion,omitempty" xml:"Reversion,omitempty" require:"true"`
	CsClusterId           *string `json:"CsClusterId,omitempty" xml:"CsClusterId,omitempty" require:"true"`
	BaseComponentMetaName *string `json:"BaseComponentMetaName,omitempty" xml:"BaseComponentMetaName,omitempty" require:"true"`
	DeploymentName        *string `json:"DeploymentName,omitempty" xml:"DeploymentName,omitempty" require:"true"`
	CpuLimit              *string `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty" require:"true"`
	MemoryLimit           *string `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty" require:"true"`
	PackagePublicUrl      *string `json:"PackagePublicUrl,omitempty" xml:"PackagePublicUrl,omitempty" require:"true"`
	PackageVersion        *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty" require:"true"`
	CpuRequest            *string `json:"CpuRequest,omitempty" xml:"CpuRequest,omitempty" require:"true"`
	MemoryRequest         *string `json:"MemoryRequest,omitempty" xml:"MemoryRequest,omitempty" require:"true"`
	VServerGroupId        *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty" require:"true"`
	VExtServerGroupId     *string `json:"VExtServerGroupId,omitempty" xml:"VExtServerGroupId,omitempty" require:"true"`
}

func (s ListDeployGroupResponseBodyDeployGroupListDeployGroup) String() string {
	return tea.Prettify(s)
}

func (s ListDeployGroupResponseBodyDeployGroupListDeployGroup) GoString() string {
	return s.String()
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetGroupId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.GroupId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetGroupName(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.GroupName = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetAppId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.AppId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPackageVersionId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PackageVersionId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetAppVersionId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.AppVersionId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetGroupType(v int) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.GroupType = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetClusterId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.ClusterId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetCreateTime(v int64) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.CreateTime = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetUpdateTime(v int64) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.UpdateTime = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetNameSpace(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.NameSpace = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetClusterName(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.ClusterName = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetLastUpdateTime(v int64) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.LastUpdateTime = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPreStop(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PreStop = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPostStart(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PostStart = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPackageUrl(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PackageUrl = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetEnv(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Env = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetLabels(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Labels = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetSelector(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Selector = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetStrategy(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Strategy = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetStatus(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Status = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetReversion(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Reversion = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetCsClusterId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.CsClusterId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetBaseComponentMetaName(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.BaseComponentMetaName = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetDeploymentName(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.DeploymentName = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetCpuLimit(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.CpuLimit = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetMemoryLimit(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.MemoryLimit = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPackagePublicUrl(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PackagePublicUrl = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPackageVersion(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PackageVersion = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetCpuRequest(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.CpuRequest = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetMemoryRequest(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.MemoryRequest = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetVServerGroupId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.VServerGroupId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetVExtServerGroupId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.VExtServerGroupId = &v
	return s
}

type ListDeployGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeployGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeployGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeployGroupResponse) GoString() string {
	return s.String()
}

func (s *ListDeployGroupResponse) SetHeaders(v map[string]*string) *ListDeployGroupResponse {
	s.Headers = v
	return s
}

func (s *ListDeployGroupResponse) SetBody(v *ListDeployGroupResponseBody) *ListDeployGroupResponse {
	s.Body = v
	return s
}

type ListBuildPackRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ListBuildPackRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBuildPackRequest) GoString() string {
	return s.String()
}

func (s *ListBuildPackRequest) SetHeaders(v map[string]*string) *ListBuildPackRequest {
	s.Headers = v
	return s
}

type ListBuildPackResponseBody struct {
	Code          *int                                    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message       *string                                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BuildPackList *ListBuildPackResponseBodyBuildPackList `json:"BuildPackList,omitempty" xml:"BuildPackList,omitempty" require:"true" type:"Struct"`
}

func (s ListBuildPackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBuildPackResponseBody) GoString() string {
	return s.String()
}

func (s *ListBuildPackResponseBody) SetCode(v int) *ListBuildPackResponseBody {
	s.Code = &v
	return s
}

func (s *ListBuildPackResponseBody) SetMessage(v string) *ListBuildPackResponseBody {
	s.Message = &v
	return s
}

func (s *ListBuildPackResponseBody) SetRequestId(v string) *ListBuildPackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBuildPackResponseBody) SetBuildPackList(v *ListBuildPackResponseBodyBuildPackList) *ListBuildPackResponseBody {
	s.BuildPackList = v
	return s
}

type ListBuildPackResponseBodyBuildPackList struct {
	BuildPack []*ListBuildPackResponseBodyBuildPackListBuildPack `json:"BuildPack,omitempty" xml:"BuildPack,omitempty" require:"true" type:"Repeated"`
}

func (s ListBuildPackResponseBodyBuildPackList) String() string {
	return tea.Prettify(s)
}

func (s ListBuildPackResponseBodyBuildPackList) GoString() string {
	return s.String()
}

func (s *ListBuildPackResponseBodyBuildPackList) SetBuildPack(v []*ListBuildPackResponseBodyBuildPackListBuildPack) *ListBuildPackResponseBodyBuildPackList {
	s.BuildPack = v
	return s
}

type ListBuildPackResponseBodyBuildPackListBuildPack struct {
	ConfigId           *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty" require:"true"`
	PackVersion        *string `json:"PackVersion,omitempty" xml:"PackVersion,omitempty" require:"true"`
	TomcatDesc         *string `json:"TomcatDesc,omitempty" xml:"TomcatDesc,omitempty" require:"true"`
	TomcatVersion      *string `json:"TomcatVersion,omitempty" xml:"TomcatVersion,omitempty" require:"true"`
	TomcatDownloadUrl  *string `json:"TomcatDownloadUrl,omitempty" xml:"TomcatDownloadUrl,omitempty" require:"true"`
	PandoraVersion     *string `json:"PandoraVersion,omitempty" xml:"PandoraVersion,omitempty" require:"true"`
	PandoraDownloadUrl *string `json:"PandoraDownloadUrl,omitempty" xml:"PandoraDownloadUrl,omitempty" require:"true"`
	PandoraDesc        *string `json:"PandoraDesc,omitempty" xml:"PandoraDesc,omitempty" require:"true"`
	PluginInfo         *string `json:"PluginInfo,omitempty" xml:"PluginInfo,omitempty" require:"true"`
	TomcatPath         *string `json:"TomcatPath,omitempty" xml:"TomcatPath,omitempty" require:"true"`
	ImageId            *string `json:"ImageId,omitempty" xml:"ImageId,omitempty" require:"true"`
	TengineImageId     *string `json:"TengineImageId,omitempty" xml:"TengineImageId,omitempty" require:"true"`
	MultipleTenant     *bool   `json:"MultipleTenant,omitempty" xml:"MultipleTenant,omitempty" require:"true"`
	WithTengine        *bool   `json:"WithTengine,omitempty" xml:"WithTengine,omitempty" require:"true"`
	TengineDownloadUrl *string `json:"TengineDownloadUrl,omitempty" xml:"TengineDownloadUrl,omitempty" require:"true"`
	ScriptName         *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty" require:"true"`
	ScriptVersion      *string `json:"ScriptVersion,omitempty" xml:"ScriptVersion,omitempty" require:"true"`
	Feature            *string `json:"Feature,omitempty" xml:"Feature,omitempty" require:"true"`
	SupportFeatures    *string `json:"SupportFeatures,omitempty" xml:"SupportFeatures,omitempty" require:"true"`
	Disabled           *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty" require:"true"`
}

func (s ListBuildPackResponseBodyBuildPackListBuildPack) String() string {
	return tea.Prettify(s)
}

func (s ListBuildPackResponseBodyBuildPackListBuildPack) GoString() string {
	return s.String()
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetConfigId(v int64) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.ConfigId = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetPackVersion(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.PackVersion = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetTomcatDesc(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.TomcatDesc = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetTomcatVersion(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.TomcatVersion = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetTomcatDownloadUrl(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.TomcatDownloadUrl = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetPandoraVersion(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.PandoraVersion = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetPandoraDownloadUrl(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.PandoraDownloadUrl = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetPandoraDesc(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.PandoraDesc = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetPluginInfo(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.PluginInfo = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetTomcatPath(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.TomcatPath = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetImageId(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.ImageId = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetTengineImageId(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.TengineImageId = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetMultipleTenant(v bool) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.MultipleTenant = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetWithTengine(v bool) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.WithTengine = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetTengineDownloadUrl(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.TengineDownloadUrl = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetScriptName(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.ScriptName = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetScriptVersion(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.ScriptVersion = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetFeature(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.Feature = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetSupportFeatures(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.SupportFeatures = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetDisabled(v bool) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.Disabled = &v
	return s
}

type ListBuildPackResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBuildPackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBuildPackResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBuildPackResponse) GoString() string {
	return s.String()
}

func (s *ListBuildPackResponse) SetHeaders(v map[string]*string) *ListBuildPackResponse {
	s.Headers = v
	return s
}

func (s *ListBuildPackResponse) SetBody(v *ListBuildPackResponseBody) *ListBuildPackResponse {
	s.Body = v
	return s
}

type ListApplicationEcuRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s ListApplicationEcuRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationEcuRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationEcuRequest) SetHeaders(v map[string]*string) *ListApplicationEcuRequest {
	s.Headers = v
	return s
}

type ListApplicationEcuResponseBody struct {
	Code        *int                                       `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message     *string                                    `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId   *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	EcuInfoList *ListApplicationEcuResponseBodyEcuInfoList `json:"EcuInfoList,omitempty" xml:"EcuInfoList,omitempty" require:"true" type:"Struct"`
}

func (s ListApplicationEcuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationEcuResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationEcuResponseBody) SetCode(v int) *ListApplicationEcuResponseBody {
	s.Code = &v
	return s
}

func (s *ListApplicationEcuResponseBody) SetMessage(v string) *ListApplicationEcuResponseBody {
	s.Message = &v
	return s
}

func (s *ListApplicationEcuResponseBody) SetRequestId(v string) *ListApplicationEcuResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationEcuResponseBody) SetEcuInfoList(v *ListApplicationEcuResponseBodyEcuInfoList) *ListApplicationEcuResponseBody {
	s.EcuInfoList = v
	return s
}

type ListApplicationEcuResponseBodyEcuInfoList struct {
	EcuEntity []*ListApplicationEcuResponseBodyEcuInfoListEcuEntity `json:"EcuEntity,omitempty" xml:"EcuEntity,omitempty" require:"true" type:"Repeated"`
}

func (s ListApplicationEcuResponseBodyEcuInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationEcuResponseBodyEcuInfoList) GoString() string {
	return s.String()
}

func (s *ListApplicationEcuResponseBodyEcuInfoList) SetEcuEntity(v []*ListApplicationEcuResponseBodyEcuInfoListEcuEntity) *ListApplicationEcuResponseBodyEcuInfoList {
	s.EcuEntity = v
	return s
}

type ListApplicationEcuResponseBodyEcuInfoListEcuEntity struct {
	EcuId         *string `json:"EcuId,omitempty" xml:"EcuId,omitempty" require:"true"`
	Online        *bool   `json:"Online,omitempty" xml:"Online,omitempty" require:"true"`
	DockerEnv     *bool   `json:"DockerEnv,omitempty" xml:"DockerEnv,omitempty" require:"true"`
	CreateTime    *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime    *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	IpAddr        *string `json:"IpAddr,omitempty" xml:"IpAddr,omitempty" require:"true"`
	HeartbeatTime *int64  `json:"HeartbeatTime,omitempty" xml:"HeartbeatTime,omitempty" require:"true"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	ZoneId        *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	AvailableCpu  *int    `json:"AvailableCpu,omitempty" xml:"AvailableCpu,omitempty" require:"true"`
	AvailableMem  *int    `json:"AvailableMem,omitempty" xml:"AvailableMem,omitempty" require:"true"`
	Cpu           *int    `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Mem           *int    `json:"Mem,omitempty" xml:"Mem,omitempty" require:"true"`
}

func (s ListApplicationEcuResponseBodyEcuInfoListEcuEntity) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationEcuResponseBodyEcuInfoListEcuEntity) GoString() string {
	return s.String()
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetEcuId(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.EcuId = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetOnline(v bool) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.Online = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetDockerEnv(v bool) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.DockerEnv = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetCreateTime(v int64) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetUpdateTime(v int64) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.UpdateTime = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetIpAddr(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.IpAddr = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetHeartbeatTime(v int64) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.HeartbeatTime = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetUserId(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.UserId = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetName(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.Name = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetZoneId(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.ZoneId = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetRegionId(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.RegionId = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetInstanceId(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetVpcId(v string) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.VpcId = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetAvailableCpu(v int) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.AvailableCpu = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetAvailableMem(v int) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.AvailableMem = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetCpu(v int) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.Cpu = &v
	return s
}

func (s *ListApplicationEcuResponseBodyEcuInfoListEcuEntity) SetMem(v int) *ListApplicationEcuResponseBodyEcuInfoListEcuEntity {
	s.Mem = &v
	return s
}

type ListApplicationEcuResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListApplicationEcuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApplicationEcuResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationEcuResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationEcuResponse) SetHeaders(v map[string]*string) *ListApplicationEcuResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationEcuResponse) SetBody(v *ListApplicationEcuResponseBody) *ListApplicationEcuResponse {
	s.Body = v
	return s
}

type GetChangeOrderInfoQuery struct {
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
}

func (s GetChangeOrderInfoQuery) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoQuery) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoQuery) SetChangeOrderId(v string) *GetChangeOrderInfoQuery {
	s.ChangeOrderId = &v
	return s
}

type GetChangeOrderInfoRequest struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *GetChangeOrderInfoQuery `json:"query,omitempty" xml:"query,omitempty" require:"true"`
}

func (s GetChangeOrderInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoRequest) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoRequest) SetHeaders(v map[string]*string) *GetChangeOrderInfoRequest {
	s.Headers = v
	return s
}

func (s *GetChangeOrderInfoRequest) SetQuery(v *GetChangeOrderInfoQuery) *GetChangeOrderInfoRequest {
	s.Query = v
	return s
}

type GetChangeOrderInfoResponseBody struct {
	Code            *int                                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message         *string                                        `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId       *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ChangeOrderInfo *GetChangeOrderInfoResponseBodyChangeOrderInfo `json:"changeOrderInfo,omitempty" xml:"changeOrderInfo,omitempty" require:"true" type:"Struct"`
}

func (s GetChangeOrderInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBody) SetCode(v int) *GetChangeOrderInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetChangeOrderInfoResponseBody) SetMessage(v string) *GetChangeOrderInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetChangeOrderInfoResponseBody) SetRequestId(v string) *GetChangeOrderInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBody) SetChangeOrderInfo(v *GetChangeOrderInfoResponseBodyChangeOrderInfo) *GetChangeOrderInfoResponseBody {
	s.ChangeOrderInfo = v
	return s
}

type GetChangeOrderInfoResponseBodyChangeOrderInfo struct {
	ChangeOrderId          *string                                                        `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty" require:"true"`
	CreateUserId           *string                                                        `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty" require:"true"`
	Desc                   *string                                                        `json:"Desc,omitempty" xml:"Desc,omitempty" require:"true"`
	BatchCount             *int                                                           `json:"BatchCount,omitempty" xml:"BatchCount,omitempty" require:"true"`
	BatchType              *string                                                        `json:"BatchType,omitempty" xml:"BatchType,omitempty" require:"true"`
	Status                 *int                                                           `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	CoType                 *string                                                        `json:"CoType,omitempty" xml:"CoType,omitempty" require:"true"`
	CreateTime             *string                                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	SupportRollback        *bool                                                          `json:"SupportRollback,omitempty" xml:"SupportRollback,omitempty" require:"true"`
	ChangeOrderDescription *string                                                        `json:"ChangeOrderDescription,omitempty" xml:"ChangeOrderDescription,omitempty" require:"true"`
	PipelineInfoList       *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList `json:"PipelineInfoList,omitempty" xml:"PipelineInfoList,omitempty" require:"true" type:"Struct"`
	TrafficControl         *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl   `json:"TrafficControl,omitempty" xml:"TrafficControl,omitempty" require:"true" type:"Struct"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfo) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfo) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetChangeOrderId(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.ChangeOrderId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetCreateUserId(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.CreateUserId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetDesc(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.Desc = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetBatchCount(v int) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.BatchCount = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetBatchType(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.BatchType = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetStatus(v int) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.Status = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetCoType(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.CoType = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetCreateTime(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.CreateTime = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetSupportRollback(v bool) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.SupportRollback = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetChangeOrderDescription(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.ChangeOrderDescription = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetPipelineInfoList(v *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.PipelineInfoList = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfo) SetTrafficControl(v *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl) *GetChangeOrderInfoResponseBodyChangeOrderInfo {
	s.TrafficControl = v
	return s
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList struct {
	PipelineInfo []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo `json:"PipelineInfo,omitempty" xml:"PipelineInfo,omitempty" require:"true" type:"Repeated"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList) SetPipelineInfo(v []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoList {
	s.PipelineInfo = v
	return s
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo struct {
	PipelineId      *string                                                                                   `json:"PipelineId,omitempty" xml:"PipelineId,omitempty" require:"true"`
	PipelineName    *string                                                                                   `json:"PipelineName,omitempty" xml:"PipelineName,omitempty" require:"true"`
	PipelineStatus  *int                                                                                      `json:"PipelineStatus,omitempty" xml:"PipelineStatus,omitempty" require:"true"`
	StartTime       *string                                                                                   `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	UpdateTime      *string                                                                                   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	StageList       *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList       `json:"StageList,omitempty" xml:"StageList,omitempty" require:"true" type:"Struct"`
	StageDetailList *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList `json:"StageDetailList,omitempty" xml:"StageDetailList,omitempty" require:"true" type:"Struct"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) SetPipelineId(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo {
	s.PipelineId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) SetPipelineName(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo {
	s.PipelineName = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) SetPipelineStatus(v int) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo {
	s.PipelineStatus = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) SetStartTime(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo {
	s.StartTime = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) SetUpdateTime(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo {
	s.UpdateTime = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) SetStageList(v *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo {
	s.StageList = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo) SetStageDetailList(v *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfo {
	s.StageDetailList = v
	return s
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList struct {
	StageInfoDTO []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO `json:"StageInfoDTO,omitempty" xml:"StageInfoDTO,omitempty" require:"true" type:"Repeated"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList) SetStageInfoDTO(v []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageList {
	s.StageInfoDTO = v
	return s
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO struct {
	StageId        *string                                                                                                       `json:"StageId,omitempty" xml:"StageId,omitempty" require:"true"`
	StageName      *string                                                                                                       `json:"StageName,omitempty" xml:"StageName,omitempty" require:"true"`
	Status         *int                                                                                                          `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	StageResultDTO *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO `json:"StageResultDTO,omitempty" xml:"StageResultDTO,omitempty" require:"true" type:"Struct"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) SetStageId(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO {
	s.StageId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) SetStageName(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO {
	s.StageName = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) SetStatus(v int) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO {
	s.Status = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO) SetStageResultDTO(v *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTO {
	s.StageResultDTO = v
	return s
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO struct {
	InstanceDTOList *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList `json:"InstanceDTOList,omitempty" xml:"InstanceDTOList,omitempty" require:"true" type:"Struct"`
	ServiceStage    *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage    `json:"ServiceStage,omitempty" xml:"ServiceStage,omitempty" require:"true" type:"Struct"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO) SetInstanceDTOList(v *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO {
	s.InstanceDTOList = v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO) SetServiceStage(v *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTO {
	s.ServiceStage = v
	return s
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList struct {
	InstanceDTO []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO `json:"InstanceDTO,omitempty" xml:"InstanceDTO,omitempty" require:"true" type:"Repeated"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList) SetInstanceDTO(v []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOList {
	s.InstanceDTO = v
	return s
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO struct {
	InstanceName         *string                                                                                                                                                     `json:"InstanceName,omitempty" xml:"InstanceName,omitempty" require:"true"`
	InstanceIp           *string                                                                                                                                                     `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty" require:"true"`
	Status               *int                                                                                                                                                        `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	PodName              *string                                                                                                                                                     `json:"PodName,omitempty" xml:"PodName,omitempty" require:"true"`
	PodStatus            *string                                                                                                                                                     `json:"PodStatus,omitempty" xml:"PodStatus,omitempty" require:"true"`
	InstanceStageDTOList *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList `json:"InstanceStageDTOList,omitempty" xml:"InstanceStageDTOList,omitempty" require:"true" type:"Struct"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) SetInstanceName(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO {
	s.InstanceName = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) SetInstanceIp(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO {
	s.InstanceIp = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) SetStatus(v int) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO {
	s.Status = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) SetPodName(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO {
	s.PodName = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) SetPodStatus(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO {
	s.PodStatus = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO) SetInstanceStageDTOList(v *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTO {
	s.InstanceStageDTOList = v
	return s
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList struct {
	InstanceStageDTO []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO `json:"InstanceStageDTO,omitempty" xml:"InstanceStageDTO,omitempty" require:"true" type:"Repeated"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList) SetInstanceStageDTO(v []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOList {
	s.InstanceStageDTO = v
	return s
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO struct {
	StageId      *string `json:"StageId,omitempty" xml:"StageId,omitempty" require:"true"`
	StageName    *string `json:"StageName,omitempty" xml:"StageName,omitempty" require:"true"`
	Status       *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	StageMessage *string `json:"StageMessage,omitempty" xml:"StageMessage,omitempty" require:"true"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	FinishTime   *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty" require:"true"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) SetStageId(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO {
	s.StageId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) SetStageName(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO {
	s.StageName = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) SetStatus(v int) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO {
	s.Status = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) SetStageMessage(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO {
	s.StageMessage = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) SetStartTime(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO {
	s.StartTime = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO) SetFinishTime(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOInstanceDTOListInstanceDTOInstanceStageDTOListInstanceStageDTO {
	s.FinishTime = &v
	return s
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage struct {
	StageId   *string `json:"StageId,omitempty" xml:"StageId,omitempty" require:"true"`
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty" require:"true"`
	Status    *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) SetStageId(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage {
	s.StageId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) SetStageName(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage {
	s.StageName = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) SetStatus(v int) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage {
	s.Status = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage) SetMessage(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageListStageInfoDTOStageResultDTOServiceStage {
	s.Message = &v
	return s
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList struct {
	StageDetailDTO []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO `json:"StageDetailDTO,omitempty" xml:"StageDetailDTO,omitempty" require:"true" type:"Repeated"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList) SetStageDetailDTO(v []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailList {
	s.StageDetailDTO = v
	return s
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO struct {
	StageId     *string                                                                                                         `json:"StageId,omitempty" xml:"StageId,omitempty" require:"true"`
	StageName   *string                                                                                                         `json:"StageName,omitempty" xml:"StageName,omitempty" require:"true"`
	StageStatus *int                                                                                                            `json:"StageStatus,omitempty" xml:"StageStatus,omitempty" require:"true"`
	TaskList    *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" require:"true" type:"Struct"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) SetStageId(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO {
	s.StageId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) SetStageName(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO {
	s.StageName = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) SetStageStatus(v int) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO {
	s.StageStatus = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO) SetTaskList(v *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTO {
	s.TaskList = v
	return s
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList struct {
	TaskInfoDTO []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO `json:"TaskInfoDTO,omitempty" xml:"TaskInfoDTO,omitempty" require:"true" type:"Repeated"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList) SetTaskInfoDTO(v []*GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskList {
	s.TaskInfoDTO = v
	return s
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO struct {
	TaskName            *string `json:"TaskName,omitempty" xml:"TaskName,omitempty" require:"true"`
	TaskStatus          *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty" require:"true"`
	TaskMessage         *string `json:"TaskMessage,omitempty" xml:"TaskMessage,omitempty" require:"true"`
	TaskId              *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	TaskErrorCode       *string `json:"TaskErrorCode,omitempty" xml:"TaskErrorCode,omitempty" require:"true"`
	TaskErrorMessage    *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty" require:"true"`
	ShowManualIgnorance *bool   `json:"ShowManualIgnorance,omitempty" xml:"ShowManualIgnorance,omitempty" require:"true"`
	TaskErrorIgnorance  *int    `json:"TaskErrorIgnorance,omitempty" xml:"TaskErrorIgnorance,omitempty" require:"true"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetTaskName(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.TaskName = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetTaskStatus(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.TaskStatus = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetTaskMessage(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.TaskMessage = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetTaskId(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.TaskId = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetTaskErrorCode(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.TaskErrorCode = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetTaskErrorMessage(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.TaskErrorMessage = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetShowManualIgnorance(v bool) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.ShowManualIgnorance = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO) SetTaskErrorIgnorance(v int) *GetChangeOrderInfoResponseBodyChangeOrderInfoPipelineInfoListPipelineInfoStageDetailListStageDetailDTOTaskListTaskInfoDTO {
	s.TaskErrorIgnorance = &v
	return s
}

type GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl struct {
	Rules  *string `json:"Rules,omitempty" xml:"Rules,omitempty" require:"true"`
	Routes *string `json:"Routes,omitempty" xml:"Routes,omitempty" require:"true"`
	Tips   *string `json:"Tips,omitempty" xml:"Tips,omitempty" require:"true"`
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl) SetRules(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl {
	s.Rules = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl) SetRoutes(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl {
	s.Routes = &v
	return s
}

func (s *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl) SetTips(v string) *GetChangeOrderInfoResponseBodyChangeOrderInfoTrafficControl {
	s.Tips = &v
	return s
}

type GetChangeOrderInfoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetChangeOrderInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetChangeOrderInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChangeOrderInfoResponse) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponse) SetHeaders(v map[string]*string) *GetChangeOrderInfoResponse {
	s.Headers = v
	return s
}

func (s *GetChangeOrderInfoResponse) SetBody(v *GetChangeOrderInfoResponseBody) *GetChangeOrderInfoResponse {
	s.Body = v
	return s
}

type ListClusterQuery struct {
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListClusterQuery) String() string {
	return tea.Prettify(s)
}

func (s ListClusterQuery) GoString() string {
	return s.String()
}

func (s *ListClusterQuery) SetLogicalRegionId(v string) *ListClusterQuery {
	s.LogicalRegionId = &v
	return s
}

func (s *ListClusterQuery) SetResourceGroupId(v string) *ListClusterQuery {
	s.ResourceGroupId = &v
	return s
}

type ListClusterRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListClusterQuery  `json:"query,omitempty" xml:"query,omitempty"`
}

func (s ListClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterRequest) GoString() string {
	return s.String()
}

func (s *ListClusterRequest) SetHeaders(v map[string]*string) *ListClusterRequest {
	s.Headers = v
	return s
}

func (s *ListClusterRequest) SetQuery(v *ListClusterQuery) *ListClusterRequest {
	s.Query = v
	return s
}

type ListClusterResponseBody struct {
	Code        *int                                `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message     *string                             `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ClusterList *ListClusterResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" require:"true" type:"Struct"`
}

func (s ListClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterResponseBody) SetCode(v int) *ListClusterResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterResponseBody) SetMessage(v string) *ListClusterResponseBody {
	s.Message = &v
	return s
}

func (s *ListClusterResponseBody) SetRequestId(v string) *ListClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterResponseBody) SetClusterList(v *ListClusterResponseBodyClusterList) *ListClusterResponseBody {
	s.ClusterList = v
	return s
}

type ListClusterResponseBodyClusterList struct {
	Cluster []*ListClusterResponseBodyClusterListCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" require:"true" type:"Repeated"`
}

func (s ListClusterResponseBodyClusterList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *ListClusterResponseBodyClusterList) SetCluster(v []*ListClusterResponseBodyClusterListCluster) *ListClusterResponseBodyClusterList {
	s.Cluster = v
	return s
}

type ListClusterResponseBodyClusterListCluster struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	ClusterName     *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty" require:"true"`
	ClusterType     *int    `json:"ClusterType,omitempty" xml:"ClusterType,omitempty" require:"true"`
	OversoldFactor  *int    `json:"OversoldFactor,omitempty" xml:"OversoldFactor,omitempty" require:"true"`
	NetworkMode     *int    `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty" require:"true"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	NodeNum         *int    `json:"NodeNum,omitempty" xml:"NodeNum,omitempty" require:"true"`
	Cpu             *int    `json:"Cpu,omitempty" xml:"Cpu,omitempty" require:"true"`
	Mem             *int    `json:"Mem,omitempty" xml:"Mem,omitempty" require:"true"`
	CpuUsed         *int    `json:"CpuUsed,omitempty" xml:"CpuUsed,omitempty" require:"true"`
	MemUsed         *int    `json:"MemUsed,omitempty" xml:"MemUsed,omitempty" require:"true"`
	CreateTime      *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime      *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	IaasProvider    *string `json:"IaasProvider,omitempty" xml:"IaasProvider,omitempty" require:"true"`
	CsClusterId     *string `json:"CsClusterId,omitempty" xml:"CsClusterId,omitempty" require:"true"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty" require:"true"`
}

func (s ListClusterResponseBodyClusterListCluster) String() string {
	return tea.Prettify(s)
}

func (s ListClusterResponseBodyClusterListCluster) GoString() string {
	return s.String()
}

func (s *ListClusterResponseBodyClusterListCluster) SetClusterId(v string) *ListClusterResponseBodyClusterListCluster {
	s.ClusterId = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetRegionId(v string) *ListClusterResponseBodyClusterListCluster {
	s.RegionId = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetDescription(v string) *ListClusterResponseBodyClusterListCluster {
	s.Description = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetClusterName(v string) *ListClusterResponseBodyClusterListCluster {
	s.ClusterName = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetClusterType(v int) *ListClusterResponseBodyClusterListCluster {
	s.ClusterType = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetOversoldFactor(v int) *ListClusterResponseBodyClusterListCluster {
	s.OversoldFactor = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetNetworkMode(v int) *ListClusterResponseBodyClusterListCluster {
	s.NetworkMode = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetVpcId(v string) *ListClusterResponseBodyClusterListCluster {
	s.VpcId = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetNodeNum(v int) *ListClusterResponseBodyClusterListCluster {
	s.NodeNum = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetCpu(v int) *ListClusterResponseBodyClusterListCluster {
	s.Cpu = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetMem(v int) *ListClusterResponseBodyClusterListCluster {
	s.Mem = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetCpuUsed(v int) *ListClusterResponseBodyClusterListCluster {
	s.CpuUsed = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetMemUsed(v int) *ListClusterResponseBodyClusterListCluster {
	s.MemUsed = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetCreateTime(v int64) *ListClusterResponseBodyClusterListCluster {
	s.CreateTime = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetUpdateTime(v int64) *ListClusterResponseBodyClusterListCluster {
	s.UpdateTime = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetIaasProvider(v string) *ListClusterResponseBodyClusterListCluster {
	s.IaasProvider = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetCsClusterId(v string) *ListClusterResponseBodyClusterListCluster {
	s.CsClusterId = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetResourceGroupId(v string) *ListClusterResponseBodyClusterListCluster {
	s.ResourceGroupId = &v
	return s
}

type ListClusterResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterResponse) GoString() string {
	return s.String()
}

func (s *ListClusterResponse) SetHeaders(v map[string]*string) *ListClusterResponse {
	s.Headers = v
	return s
}

func (s *ListClusterResponse) SetBody(v *ListClusterResponseBody) *ListClusterResponse {
	s.Body = v
	return s
}

type ListApplicationQuery struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListApplicationQuery) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationQuery) GoString() string {
	return s.String()
}

func (s *ListApplicationQuery) SetClusterId(v string) *ListApplicationQuery {
	s.ClusterId = &v
	return s
}

func (s *ListApplicationQuery) SetLogicalRegionId(v string) *ListApplicationQuery {
	s.LogicalRegionId = &v
	return s
}

func (s *ListApplicationQuery) SetAppName(v string) *ListApplicationQuery {
	s.AppName = &v
	return s
}

func (s *ListApplicationQuery) SetResourceGroupId(v string) *ListApplicationQuery {
	s.ResourceGroupId = &v
	return s
}

type ListApplicationRequest struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *ListApplicationQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s ListApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationRequest) SetHeaders(v map[string]*string) *ListApplicationRequest {
	s.Headers = v
	return s
}

func (s *ListApplicationRequest) SetQuery(v *ListApplicationQuery) *ListApplicationRequest {
	s.Query = v
	return s
}

type ListApplicationResponseBody struct {
	Code            *int                                        `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message         *string                                     `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId       *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ApplicationList *ListApplicationResponseBodyApplicationList `json:"ApplicationList,omitempty" xml:"ApplicationList,omitempty" require:"true" type:"Struct"`
}

func (s ListApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBody) SetCode(v int) *ListApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ListApplicationResponseBody) SetMessage(v string) *ListApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ListApplicationResponseBody) SetRequestId(v string) *ListApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationResponseBody) SetApplicationList(v *ListApplicationResponseBodyApplicationList) *ListApplicationResponseBody {
	s.ApplicationList = v
	return s
}

type ListApplicationResponseBodyApplicationList struct {
	Application []*ListApplicationResponseBodyApplicationListApplication `json:"Application,omitempty" xml:"Application,omitempty" require:"true" type:"Repeated"`
}

func (s ListApplicationResponseBodyApplicationList) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationResponseBodyApplicationList) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBodyApplicationList) SetApplication(v []*ListApplicationResponseBodyApplicationListApplication) *ListApplicationResponseBodyApplicationList {
	s.Application = v
	return s
}

type ListApplicationResponseBodyApplicationListApplication struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ApplicationType      *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty" require:"true"`
	ClusterType          *int    `json:"ClusterType,omitempty" xml:"ClusterType,omitempty" require:"true"`
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty" require:"true"`
	BuildPackageId       *int64  `json:"BuildPackageId,omitempty" xml:"BuildPackageId,omitempty" require:"true"`
	RunningInstanceCount *int    `json:"RunningInstanceCount,omitempty" xml:"RunningInstanceCount,omitempty" require:"true"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty" require:"true"`
}

func (s ListApplicationResponseBodyApplicationListApplication) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationResponseBodyApplicationListApplication) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetAppId(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.AppId = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetName(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.Name = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetRegionId(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.RegionId = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetApplicationType(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.ApplicationType = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetClusterType(v int) *ListApplicationResponseBodyApplicationListApplication {
	s.ClusterType = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetClusterId(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.ClusterId = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetBuildPackageId(v int64) *ListApplicationResponseBodyApplicationListApplication {
	s.BuildPackageId = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetRunningInstanceCount(v int) *ListApplicationResponseBodyApplicationListApplication {
	s.RunningInstanceCount = &v
	return s
}

func (s *ListApplicationResponseBodyApplicationListApplication) SetResourceGroupId(v string) *ListApplicationResponseBodyApplicationListApplication {
	s.ResourceGroupId = &v
	return s
}

type ListApplicationResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationResponse) SetHeaders(v map[string]*string) *ListApplicationResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationResponse) SetBody(v *ListApplicationResponseBody) *ListApplicationResponse {
	s.Body = v
	return s
}

type Client struct {
	roa.Client
}

func NewClient(config *roa.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *roa.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("edas.ap-northeast-1.aliyuncs.com"),
		"ap-south-1":                  tea.String("edas.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-3":              tea.String("edas.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-5":              tea.String("edas.ap-northeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("edas.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("edas.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("edas.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("edas.aliyuncs.com"),
		"cn-chengdu":                  tea.String("edas.aliyuncs.com"),
		"cn-edge-1":                   tea.String("edas.aliyuncs.com"),
		"cn-fujian":                   tea.String("edas.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("edas.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("edas.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("edas.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("edas.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("edas.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("edas.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("edas.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("edas.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("edas.aliyuncs.com"),
		"cn-huhehaote":                tea.String("edas.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("edas.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("edas.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("edas.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("edas.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("edas.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("edas.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("edas.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("edas.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("edas.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("edas.aliyuncs.com"),
		"cn-wuhan":                    tea.String("edas.aliyuncs.com"),
		"cn-yushanfang":               tea.String("edas.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("edas.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("edas.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("edas.aliyuncs.com"),
		"eu-west-1":                   tea.String("edas.ap-northeast-1.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("edas.ap-northeast-1.aliyuncs.com"),
		"me-east-1":                   tea.String("edas.ap-northeast-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("edas.ap-northeast-1.aliyuncs.com"),
		"us-west-1":                   tea.String("edas.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.EndpointHost, _err = client.GetEndpoint(tea.String("edas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.EndpointHost)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) CreateModuleWithOptions(request *CreateModuleRequest, runtime *util.RuntimeOptions) (_result *CreateModuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateModuleResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("CreateModule"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/oam/feature/module_actions"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, tea.ToMap(request.Body), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateModule(request *CreateModuleRequest) (_result *CreateModuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateModuleResponse{}
	_body, _err := client.CreateModuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateK8sApplicationBaseInfoWithOptions(request *UpdateK8sApplicationBaseInfoRequest, runtime *util.RuntimeOptions) (_result *UpdateK8sApplicationBaseInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateK8sApplicationBaseInfoResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateK8sApplicationBaseInfo"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/oam/update_app_basic_info"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateK8sApplicationBaseInfo(request *UpdateK8sApplicationBaseInfoRequest) (_result *UpdateK8sApplicationBaseInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateK8sApplicationBaseInfoResponse{}
	_body, _err := client.UpdateK8sApplicationBaseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScaleoutApplicationWithNewInstancesWithOptions(request *ScaleoutApplicationWithNewInstancesRequest, runtime *util.RuntimeOptions) (_result *ScaleoutApplicationWithNewInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ScaleoutApplicationWithNewInstancesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ScaleoutApplicationWithNewInstances"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/scaling/scale_out"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScaleoutApplicationWithNewInstances(request *ScaleoutApplicationWithNewInstancesRequest) (_result *ScaleoutApplicationWithNewInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScaleoutApplicationWithNewInstancesResponse{}
	_body, _err := client.ScaleoutApplicationWithNewInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetK8sClusterWithOptions(request *GetK8sClusterRequest, runtime *util.RuntimeOptions) (_result *GetK8sClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetK8sClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GetK8sCluster"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/k8s_clusters"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetK8sCluster(request *GetK8sClusterRequest) (_result *GetK8sClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetK8sClusterResponse{}
	_body, _err := client.GetK8sClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIDCImportCommandWithOptions(request *CreateIDCImportCommandRequest, runtime *util.RuntimeOptions) (_result *CreateIDCImportCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateIDCImportCommandResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("CreateIDCImportCommand"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/create_idc_import_command"), nil, request.Headers, tea.ToMap(request.Body), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIDCImportCommand(request *CreateIDCImportCommandRequest) (_result *CreateIDCImportCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIDCImportCommandResponse{}
	_body, _err := client.CreateIDCImportCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOperationLogsWithOptions(request *ListOperationLogsRequest, runtime *util.RuntimeOptions) (_result *ListOperationLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListOperationLogsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListOperationLogs"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/operationlog/log_list"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOperationLogs(request *ListOperationLogsRequest) (_result *ListOperationLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOperationLogsResponse{}
	_body, _err := client.ListOperationLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRootStacksWithOptions(request *ListRootStacksRequest, runtime *util.RuntimeOptions) (_result *ListRootStacksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListRootStacksResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListRootStacks"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/s2i/list_root_stack"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRootStacks(request *ListRootStacksRequest) (_result *ListRootStacksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRootStacksResponse{}
	_body, _err := client.ListRootStacksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListChildrenStacksWithOptions(request *ListChildrenStacksRequest, runtime *util.RuntimeOptions) (_result *ListChildrenStacksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListChildrenStacksResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListChildrenStacks"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/s2i/list_children_stack"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListChildrenStacks(request *ListChildrenStacksRequest) (_result *ListChildrenStacksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListChildrenStacksResponse{}
	_body, _err := client.ListChildrenStacksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartK8sApplicationWithOptions(request *StartK8sApplicationRequest, runtime *util.RuntimeOptions) (_result *StartK8sApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartK8sApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("StartK8sApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/k8s/acs/start_k8s_app"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartK8sApplication(request *StartK8sApplicationRequest) (_result *StartK8sApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartK8sApplicationResponse{}
	_body, _err := client.StartK8sApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartK8sApplicationWithOptions(request *RestartK8sApplicationRequest, runtime *util.RuntimeOptions) (_result *RestartK8sApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RestartK8sApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("RestartK8sApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/k8s/acs/restart_k8s_app"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartK8sApplication(request *RestartK8sApplicationRequest) (_result *RestartK8sApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartK8sApplicationResponse{}
	_body, _err := client.RestartK8sApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindEcsSlbWithOptions(request *BindEcsSlbRequest, runtime *util.RuntimeOptions) (_result *BindEcsSlbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &BindEcsSlbResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("BindEcsSlb"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/app/slb/bind_slb"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindEcsSlb(request *BindEcsSlbRequest) (_result *BindEcsSlbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindEcsSlbResponse{}
	_body, _err := client.BindEcsSlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopK8sApplicationWithOptions(request *StopK8sApplicationRequest, runtime *util.RuntimeOptions) (_result *StopK8sApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopK8sApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("StopK8sApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/k8s/acs/stop_k8s_app"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopK8sApplication(request *StopK8sApplicationRequest) (_result *StopK8sApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopK8sApplicationResponse{}
	_body, _err := client.StopK8sApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConvertK8sResourceWithOptions(request *ConvertK8sResourceRequest, runtime *util.RuntimeOptions) (_result *ConvertK8sResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ConvertK8sResourceResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ConvertK8sResource"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/oam/k8s_resource_convert"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConvertK8sResource(request *ConvertK8sResourceRequest) (_result *ConvertK8sResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConvertK8sResourceResponse{}
	_body, _err := client.ConvertK8sResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSlsLogStoreWithOptions(request *UpdateSlsLogStoreRequest, runtime *util.RuntimeOptions) (_result *UpdateSlsLogStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateSlsLogStoreResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateSlsLogStore"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/k8s/sls/update_sls_log_store"), nil, request.Headers, tea.ToMap(request.Body), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSlsLogStore(request *UpdateSlsLogStoreRequest) (_result *UpdateSlsLogStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSlsLogStoreResponse{}
	_body, _err := client.UpdateSlsLogStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryK8sClusterLogProjectInfoWithOptions(request *QueryK8sClusterLogProjectInfoRequest, runtime *util.RuntimeOptions) (_result *QueryK8sClusterLogProjectInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryK8sClusterLogProjectInfoResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("QueryK8sClusterLogProjectInfo"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/k8s/sls/project"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryK8sClusterLogProjectInfo(request *QueryK8sClusterLogProjectInfoRequest) (_result *QueryK8sClusterLogProjectInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryK8sClusterLogProjectInfoResponse{}
	_body, _err := client.QueryK8sClusterLogProjectInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySlsLogStoreListWithOptions(request *QuerySlsLogStoreListRequest, runtime *util.RuntimeOptions) (_result *QuerySlsLogStoreListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QuerySlsLogStoreListResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("QuerySlsLogStoreList"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/k8s/sls/query_sls_log_store_list"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySlsLogStoreList(request *QuerySlsLogStoreListRequest) (_result *QuerySlsLogStoreListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySlsLogStoreListResponse{}
	_body, _err := client.QuerySlsLogStoreListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSubAccountInfoWithOptions(request *GetSubAccountInfoRequest, runtime *util.RuntimeOptions) (_result *GetSubAccountInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSubAccountInfoResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GetSubAccountInfo"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/account/sub_account_info"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSubAccountInfo(request *GetSubAccountInfoRequest) (_result *GetSubAccountInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSubAccountInfoResponse{}
	_body, _err := client.GetSubAccountInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DelegateAdminRoleWithOptions(request *DelegateAdminRoleRequest, runtime *util.RuntimeOptions) (_result *DelegateAdminRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DelegateAdminRoleResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DelegateAdminRole"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/pop/v5/account/delegate_admin_role"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DelegateAdminRole(request *DelegateAdminRoleRequest) (_result *DelegateAdminRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DelegateAdminRoleResponse{}
	_body, _err := client.DelegateAdminRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartApplicationWithOptions(request *RestartApplicationRequest, runtime *util.RuntimeOptions) (_result *RestartApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RestartApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("RestartApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/changeorder/co_restart"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartApplication(request *RestartApplicationRequest) (_result *RestartApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartApplicationResponse{}
	_body, _err := client.RestartApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLogPathWithOptions(request *DeleteLogPathRequest, runtime *util.RuntimeOptions) (_result *DeleteLogPathResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLogPathResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteLogPath"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v5/log/popListLogDirs"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLogPath(request *DeleteLogPathRequest) (_result *DeleteLogPathResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLogPathResponse{}
	_body, _err := client.DeleteLogPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLogPathWithOptions(request *AddLogPathRequest, runtime *util.RuntimeOptions) (_result *AddLogPathResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddLogPathResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("AddLogPath"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/log/popListLogDirs"), nil, request.Headers, tea.ToMap(request.Body), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLogPath(request *AddLogPathRequest) (_result *AddLogPathResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLogPathResponse{}
	_body, _err := client.AddLogPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateK8sResourceWithOptions(request *UpdateK8sResourceRequest, runtime *util.RuntimeOptions) (_result *UpdateK8sResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateK8sResourceResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateK8sResource"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/oam/update_k8s_resource_config"), nil, request.Headers, tea.ToMap(request.Body), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateK8sResource(request *UpdateK8sResourceRequest) (_result *UpdateK8sResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateK8sResourceResponse{}
	_body, _err := client.UpdateK8sResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AbortChangeOrderWithOptions(request *AbortChangeOrderRequest, runtime *util.RuntimeOptions) (_result *AbortChangeOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AbortChangeOrderResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("AbortChangeOrder"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/pop/v5/changeorder/change_order_abort"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AbortChangeOrder(request *AbortChangeOrderRequest) (_result *AbortChangeOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbortChangeOrderResponse{}
	_body, _err := client.AbortChangeOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackChangeOrderWithOptions(request *RollbackChangeOrderRequest, runtime *util.RuntimeOptions) (_result *RollbackChangeOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RollbackChangeOrderResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("RollbackChangeOrder"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/pop/v5/oam/changeorder/rollback"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackChangeOrder(request *RollbackChangeOrderRequest) (_result *RollbackChangeOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackChangeOrderResponse{}
	_body, _err := client.RollbackChangeOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AbortAndRollbackChangeOrderWithOptions(request *AbortAndRollbackChangeOrderRequest, runtime *util.RuntimeOptions) (_result *AbortAndRollbackChangeOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AbortAndRollbackChangeOrderResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("AbortAndRollbackChangeOrder"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/pop/v5/changeorder/change_order_abort_and_rollback"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AbortAndRollbackChangeOrder(request *AbortAndRollbackChangeOrderRequest) (_result *AbortAndRollbackChangeOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbortAndRollbackChangeOrderResponse{}
	_body, _err := client.AbortAndRollbackChangeOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetScalingRulesWithOptions(request *GetScalingRulesRequest, runtime *util.RuntimeOptions) (_result *GetScalingRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetScalingRulesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GetScalingRules"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/app/scalingRules"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetScalingRules(request *GetScalingRulesRequest) (_result *GetScalingRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetScalingRulesResponse{}
	_body, _err := client.GetScalingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyScalingRuleWithOptions(request *ModifyScalingRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyScalingRuleResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ModifyScalingRule"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/app/scaling_rules"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyScalingRule(request *ModifyScalingRuleRequest) (_result *ModifyScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyScalingRuleResponse{}
	_body, _err := client.ModifyScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMethodsWithOptions(request *ListMethodsRequest, runtime *util.RuntimeOptions) (_result *ListMethodsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListMethodsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListMethods"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/service/list_methods"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMethods(request *ListMethodsRequest) (_result *ListMethodsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMethodsResponse{}
	_body, _err := client.ListMethodsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetK8sApplicationWithOptions(request *GetK8sApplicationRequest, runtime *util.RuntimeOptions) (_result *GetK8sApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetK8sApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GetK8sApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/changeorder/co_application"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetK8sApplication(request *GetK8sApplicationRequest) (_result *GetK8sApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetK8sApplicationResponse{}
	_body, _err := client.GetK8sApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEccInfoWithOptions(request *QueryEccInfoRequest, runtime *util.RuntimeOptions) (_result *QueryEccInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryEccInfoResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("QueryEccInfo"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/ecc"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEccInfo(request *QueryEccInfoRequest) (_result *QueryEccInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEccInfoResponse{}
	_body, _err := client.QueryEccInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ContinuePipelineWithOptions(request *ContinuePipelineRequest, runtime *util.RuntimeOptions) (_result *ContinuePipelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ContinuePipelineResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ContinuePipeline"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/changeorder/pipeline_batch_confirm"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ContinuePipeline(request *ContinuePipelineRequest) (_result *ContinuePipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContinuePipelineResponse{}
	_body, _err := client.ContinuePipelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeDeployGroupWithOptions(request *ChangeDeployGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeDeployGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ChangeDeployGroupResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ChangeDeployGroup"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/changeorder/co_change_group"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeDeployGroup(request *ChangeDeployGroupRequest) (_result *ChangeDeployGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeDeployGroupResponse{}
	_body, _err := client.ChangeDeployGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClusterWithOptions(request *GetClusterRequest, runtime *util.RuntimeOptions) (_result *GetClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GetCluster"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/resource/cluster"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCluster(request *GetClusterRequest) (_result *GetClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClusterResponse{}
	_body, _err := client.GetClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRegionConfigWithOptions(request *QueryRegionConfigRequest, runtime *util.RuntimeOptions) (_result *QueryRegionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryRegionConfigResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("QueryRegionConfig"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/region_config"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRegionConfig(request *QueryRegionConfigRequest) (_result *QueryRegionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRegionConfigResponse{}
	_body, _err := client.QueryRegionConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SynchronizeResourceWithOptions(request *SynchronizeResourceRequest, runtime *util.RuntimeOptions) (_result *SynchronizeResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SynchronizeResourceResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("SynchronizeResource"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/resource/pop_sync_resource"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SynchronizeResource(request *SynchronizeResourceRequest) (_result *SynchronizeResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SynchronizeResourceResponse{}
	_body, _err := client.SynchronizeResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallAgentWithOptions(request *InstallAgentRequest, runtime *util.RuntimeOptions) (_result *InstallAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InstallAgentResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("InstallAgent"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/ecss/install_agent"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallAgent(request *InstallAgentRequest) (_result *InstallAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallAgentResponse{}
	_body, _err := client.InstallAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListComponentsWithOptions(request *ListComponentsRequest, runtime *util.RuntimeOptions) (_result *ListComponentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListComponentsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListComponents"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/resource/components"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListComponents(request *ListComponentsRequest) (_result *ListComponentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListComponentsResponse{}
	_body, _err := client.ListComponentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPackageStorageCredentialWithOptions(request *GetPackageStorageCredentialRequest, runtime *util.RuntimeOptions) (_result *GetPackageStorageCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPackageStorageCredentialResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GetPackageStorageCredential"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/package_storage_credential"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPackageStorageCredential(request *GetPackageStorageCredentialRequest) (_result *GetPackageStorageCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPackageStorageCredentialResponse{}
	_body, _err := client.GetPackageStorageCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEcsNotInClusterWithOptions(request *ListEcsNotInClusterRequest, runtime *util.RuntimeOptions) (_result *ListEcsNotInClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListEcsNotInClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListEcsNotInCluster"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/resource/ecs_not_in_cluster"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEcsNotInCluster(request *ListEcsNotInClusterRequest) (_result *ListEcsNotInClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEcsNotInClusterResponse{}
	_body, _err := client.ListEcsNotInClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertK8sApplicationWithOptions(request *InsertK8sApplicationRequest, runtime *util.RuntimeOptions) (_result *InsertK8sApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InsertK8sApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("InsertK8sApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/k8s/acs/create_k8s_app"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertK8sApplication(request *InsertK8sApplicationRequest) (_result *InsertK8sApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertK8sApplicationResponse{}
	_body, _err := client.InsertK8sApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportK8sClusterWithOptions(request *ImportK8sClusterRequest, runtime *util.RuntimeOptions) (_result *ImportK8sClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ImportK8sClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ImportK8sCluster"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/import_k8s_cluster"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportK8sCluster(request *ImportK8sClusterRequest) (_result *ImportK8sClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportK8sClusterResponse{}
	_body, _err := client.ImportK8sClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateK8sApplicationConfigWithOptions(request *UpdateK8sApplicationConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateK8sApplicationConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateK8sApplicationConfigResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateK8sApplicationConfig"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/pop/v5/k8s/acs/k8s_app_configuration"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateK8sApplicationConfig(request *UpdateK8sApplicationConfigRequest) (_result *UpdateK8sApplicationConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateK8sApplicationConfigResponse{}
	_body, _err := client.UpdateK8sApplicationConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployK8sApplicationWithOptions(request *DeployK8sApplicationRequest, runtime *util.RuntimeOptions) (_result *DeployK8sApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeployK8sApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeployK8sApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/k8s/acs/k8s_apps"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployK8sApplication(request *DeployK8sApplicationRequest) (_result *DeployK8sApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployK8sApplicationResponse{}
	_body, _err := client.DeployK8sApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScaleK8sApplicationWithOptions(request *ScaleK8sApplicationRequest, runtime *util.RuntimeOptions) (_result *ScaleK8sApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ScaleK8sApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ScaleK8sApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/pop/v5/k8s/acs/k8s_apps"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScaleK8sApplication(request *ScaleK8sApplicationRequest) (_result *ScaleK8sApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScaleK8sApplicationResponse{}
	_body, _err := client.ScaleK8sApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteK8sApplicationWithOptions(request *DeleteK8sApplicationRequest, runtime *util.RuntimeOptions) (_result *DeleteK8sApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteK8sApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteK8sApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v5/k8s/acs/k8s_apps"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteK8sApplication(request *DeleteK8sApplicationRequest) (_result *DeleteK8sApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteK8sApplicationResponse{}
	_body, _err := client.DeleteK8sApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindK8sSlbWithOptions(request *UnbindK8sSlbRequest, runtime *util.RuntimeOptions) (_result *UnbindK8sSlbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UnbindK8sSlbResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UnbindK8sSlb"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v5/k8s/acs/k8s_slb_binding"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindK8sSlb(request *UnbindK8sSlbRequest) (_result *UnbindK8sSlbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindK8sSlbResponse{}
	_body, _err := client.UnbindK8sSlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindK8sSlbWithOptions(request *BindK8sSlbRequest, runtime *util.RuntimeOptions) (_result *BindK8sSlbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &BindK8sSlbResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("BindK8sSlb"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/k8s/acs/k8s_slb_binding"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindK8sSlb(request *BindK8sSlbRequest) (_result *BindK8sSlbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindK8sSlbResponse{}
	_body, _err := client.BindK8sSlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateK8sSlbWithOptions(request *UpdateK8sSlbRequest, runtime *util.RuntimeOptions) (_result *UpdateK8sSlbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateK8sSlbResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateK8sSlb"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/pop/v5/k8s/acs/k8s_slb_binding"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateK8sSlb(request *UpdateK8sSlbRequest) (_result *UpdateK8sSlbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateK8sSlbResponse{}
	_body, _err := client.UpdateK8sSlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSecureTokenWithOptions(request *GetSecureTokenRequest, runtime *util.RuntimeOptions) (_result *GetSecureTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSecureTokenResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GetSecureToken"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/secure_token"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSecureToken(request *GetSecureTokenRequest) (_result *GetSecureTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSecureTokenResponse{}
	_body, _err := client.GetSecureTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransformClusterMemberWithOptions(request *TransformClusterMemberRequest, runtime *util.RuntimeOptions) (_result *TransformClusterMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &TransformClusterMemberResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("TransformClusterMember"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/resource/transform_cluster_member"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransformClusterMember(request *TransformClusterMemberRequest) (_result *TransformClusterMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransformClusterMemberResponse{}
	_body, _err := client.TransformClusterMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConvertableEcuWithOptions(request *ListConvertableEcuRequest, runtime *util.RuntimeOptions) (_result *ListConvertableEcuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListConvertableEcuResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListConvertableEcu"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/resource/convertable_ecu_list"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConvertableEcu(request *ListConvertableEcuRequest) (_result *ListConvertableEcuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConvertableEcuResponse{}
	_body, _err := client.ListConvertableEcuWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertClusterMemberWithOptions(request *InsertClusterMemberRequest, runtime *util.RuntimeOptions) (_result *InsertClusterMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InsertClusterMemberResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("InsertClusterMember"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/resource/cluster_member"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertClusterMember(request *InsertClusterMemberRequest) (_result *InsertClusterMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertClusterMemberResponse{}
	_body, _err := client.InsertClusterMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRoleWithOptions(request *UpdateRoleRequest, runtime *util.RuntimeOptions) (_result *UpdateRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateRoleResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateRole"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/account/edit_role"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRole(request *UpdateRoleRequest) (_result *UpdateRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRoleResponse{}
	_body, _err := client.UpdateRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateJvmConfigurationWithOptions(request *UpdateJvmConfigurationRequest, runtime *util.RuntimeOptions) (_result *UpdateJvmConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateJvmConfigurationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateJvmConfiguration"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/app/app_jvm_config"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateJvmConfiguration(request *UpdateJvmConfigurationRequest) (_result *UpdateJvmConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateJvmConfigurationResponse{}
	_body, _err := client.UpdateJvmConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateHealthCheckUrlWithOptions(request *UpdateHealthCheckUrlRequest, runtime *util.RuntimeOptions) (_result *UpdateHealthCheckUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateHealthCheckUrlResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateHealthCheckUrl"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/app/modify_hc_url"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateHealthCheckUrl(request *UpdateHealthCheckUrlRequest) (_result *UpdateHealthCheckUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateHealthCheckUrlResponse{}
	_body, _err := client.UpdateHealthCheckUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFlowControlWithOptions(request *UpdateFlowControlRequest, runtime *util.RuntimeOptions) (_result *UpdateFlowControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateFlowControlResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateFlowControl"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/pop/v5/flowControl"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFlowControl(request *UpdateFlowControlRequest) (_result *UpdateFlowControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFlowControlResponse{}
	_body, _err := client.UpdateFlowControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDegradeControlWithOptions(request *UpdateDegradeControlRequest, runtime *util.RuntimeOptions) (_result *UpdateDegradeControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateDegradeControlResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateDegradeControl"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/pop/v5/degradeControl"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDegradeControl(request *UpdateDegradeControlRequest) (_result *UpdateDegradeControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDegradeControlResponse{}
	_body, _err := client.UpdateDegradeControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateContainerConfigurationWithOptions(request *UpdateContainerConfigurationRequest, runtime *util.RuntimeOptions) (_result *UpdateContainerConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateContainerConfigurationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateContainerConfiguration"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/app/container_config"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateContainerConfiguration(request *UpdateContainerConfigurationRequest) (_result *UpdateContainerConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateContainerConfigurationResponse{}
	_body, _err := client.UpdateContainerConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateApplicationBaseInfoWithOptions(request *UpdateApplicationBaseInfoRequest, runtime *util.RuntimeOptions) (_result *UpdateApplicationBaseInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateApplicationBaseInfoResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateApplicationBaseInfo"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/app/update_app_info"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApplicationBaseInfo(request *UpdateApplicationBaseInfoRequest) (_result *UpdateApplicationBaseInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateApplicationBaseInfoResponse{}
	_body, _err := client.UpdateApplicationBaseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAccountInfoWithOptions(request *UpdateAccountInfoRequest, runtime *util.RuntimeOptions) (_result *UpdateAccountInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateAccountInfoResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateAccountInfo"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/account/edit_account_info"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAccountInfo(request *UpdateAccountInfoRequest) (_result *UpdateAccountInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAccountInfoResponse{}
	_body, _err := client.UpdateAccountInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindSlbWithOptions(request *UnbindSlbRequest, runtime *util.RuntimeOptions) (_result *UnbindSlbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UnbindSlbResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UnbindSlb"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/app/unbind_slb_json"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindSlb(request *UnbindSlbRequest) (_result *UnbindSlbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindSlbResponse{}
	_body, _err := client.UnbindSlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackApplicationWithOptions(request *RollbackApplicationRequest, runtime *util.RuntimeOptions) (_result *RollbackApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RollbackApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("RollbackApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/changeorder/co_rollback"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackApplication(request *RollbackApplicationRequest) (_result *RollbackApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackApplicationResponse{}
	_body, _err := client.RollbackApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMonitorInfoWithOptions(tmp *QueryMonitorInfoRequest, runtime *util.RuntimeOptions) (_result *QueryMonitorInfoResponse, _err error) {
	_err = util.ValidateModel(tmp)
	if _err != nil {
		return _result, _err
	}
	request := &QueryMonitorInfoShrinkRequest{}
	roautil.Convert(tmp, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmp.Query))) {
		query := &QueryMonitorInfoShrinkQuery{}
		roautil.Convert(tmp.Query, query)
		if !tea.BoolValue(util.IsUnset(tmp.Query.Tags)) {
			query.TagsShrink = util.ToJSONString(tmp.Query.Tags)
		}

		request.Query = query
	}

	_result = &QueryMonitorInfoResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("QueryMonitorInfo"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/monitor/queryMonitorInfo"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMonitorInfo(request *QueryMonitorInfoRequest) (_result *QueryMonitorInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMonitorInfoResponse{}
	_body, _err := client.QueryMonitorInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMigrateRegionListWithOptions(request *QueryMigrateRegionListRequest, runtime *util.RuntimeOptions) (_result *QueryMigrateRegionListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryMigrateRegionListResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("QueryMigrateRegionList"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/resource/migrate_region_select"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMigrateRegionList(request *QueryMigrateRegionListRequest) (_result *QueryMigrateRegionListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMigrateRegionListResponse{}
	_body, _err := client.QueryMigrateRegionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMigrateEcuListWithOptions(request *QueryMigrateEcuListRequest, runtime *util.RuntimeOptions) (_result *QueryMigrateEcuListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryMigrateEcuListResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("QueryMigrateEcuList"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/resource/migrate_ecu_list"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMigrateEcuList(request *QueryMigrateEcuListRequest) (_result *QueryMigrateEcuListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMigrateEcuListResponse{}
	_body, _err := client.QueryMigrateEcuListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryConfigCenterWithOptions(request *QueryConfigCenterRequest, runtime *util.RuntimeOptions) (_result *QueryConfigCenterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryConfigCenterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("QueryConfigCenter"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/configCenter"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryConfigCenter(request *QueryConfigCenterRequest) (_result *QueryConfigCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryConfigCenterResponse{}
	_body, _err := client.QueryConfigCenterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryApplicationStatusWithOptions(request *QueryApplicationStatusRequest, runtime *util.RuntimeOptions) (_result *QueryApplicationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryApplicationStatusResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("QueryApplicationStatus"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/app/app_status"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryApplicationStatus(request *QueryApplicationStatusRequest) (_result *QueryApplicationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryApplicationStatusResponse{}
	_body, _err := client.QueryApplicationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MigrateEcuWithOptions(request *MigrateEcuRequest, runtime *util.RuntimeOptions) (_result *MigrateEcuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &MigrateEcuResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("MigrateEcu"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/resource/migrate_ecu"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MigrateEcu(request *MigrateEcuRequest) (_result *MigrateEcuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MigrateEcuResponse{}
	_body, _err := client.MigrateEcuWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserDefineRegionWithOptions(request *ListUserDefineRegionRequest, runtime *util.RuntimeOptions) (_result *ListUserDefineRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListUserDefineRegionResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListUserDefineRegion"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/user_region_defs"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserDefineRegion(request *ListUserDefineRegionRequest) (_result *ListUserDefineRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserDefineRegionResponse{}
	_body, _err := client.ListUserDefineRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubAccountWithOptions(request *ListSubAccountRequest, runtime *util.RuntimeOptions) (_result *ListSubAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListSubAccountResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListSubAccount"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/account/sub_account_list"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSubAccount(request *ListSubAccountRequest) (_result *ListSubAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSubAccountResponse{}
	_body, _err := client.ListSubAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSlbWithOptions(request *ListSlbRequest, runtime *util.RuntimeOptions) (_result *ListSlbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListSlbResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListSlb"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/slb_list"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSlb(request *ListSlbRequest) (_result *ListSlbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSlbResponse{}
	_body, _err := client.ListSlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServiceGroupsWithOptions(request *ListServiceGroupsRequest, runtime *util.RuntimeOptions) (_result *ListServiceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListServiceGroupsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListServiceGroups"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/service/serviceGroups"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceGroups(request *ListServiceGroupsRequest) (_result *ListServiceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceGroupsResponse{}
	_body, _err := client.ListServiceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListScaleOutEcuWithOptions(request *ListScaleOutEcuRequest, runtime *util.RuntimeOptions) (_result *ListScaleOutEcuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListScaleOutEcuResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListScaleOutEcu"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/resource/scale_out_ecu_list"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListScaleOutEcu(request *ListScaleOutEcuRequest) (_result *ListScaleOutEcuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScaleOutEcuResponse{}
	_body, _err := client.ListScaleOutEcuWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRoleWithOptions(request *ListRoleRequest, runtime *util.RuntimeOptions) (_result *ListRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListRoleResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListRole"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/account/role_list"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRole(request *ListRoleRequest) (_result *ListRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRoleResponse{}
	_body, _err := client.ListRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceGroupWithOptions(request *ListResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ListResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListResourceGroupResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListResourceGroup"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/resource/reg_group_list"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceGroup(request *ListResourceGroupRequest) (_result *ListResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceGroupResponse{}
	_body, _err := client.ListResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRecentChangeOrderWithOptions(request *ListRecentChangeOrderRequest, runtime *util.RuntimeOptions) (_result *ListRecentChangeOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListRecentChangeOrderResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListRecentChangeOrder"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/changeorder/change_order_list"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRecentChangeOrder(request *ListRecentChangeOrderRequest) (_result *ListRecentChangeOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRecentChangeOrderResponse{}
	_body, _err := client.ListRecentChangeOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPublishedServicesWithOptions(request *ListPublishedServicesRequest, runtime *util.RuntimeOptions) (_result *ListPublishedServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListPublishedServicesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListPublishedServices"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/service/listPublishedServices"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPublishedServices(request *ListPublishedServicesRequest) (_result *ListPublishedServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublishedServicesResponse{}
	_body, _err := client.ListPublishedServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHistoryDeployVersionWithOptions(request *ListHistoryDeployVersionRequest, runtime *util.RuntimeOptions) (_result *ListHistoryDeployVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListHistoryDeployVersionResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListHistoryDeployVersion"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/app/deploy_history_version_list"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHistoryDeployVersion(request *ListHistoryDeployVersionRequest) (_result *ListHistoryDeployVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHistoryDeployVersionResponse{}
	_body, _err := client.ListHistoryDeployVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowControlsWithOptions(request *ListFlowControlsRequest, runtime *util.RuntimeOptions) (_result *ListFlowControlsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListFlowControlsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListFlowControls"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/app/flowControls"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowControls(request *ListFlowControlsRequest) (_result *ListFlowControlsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowControlsResponse{}
	_body, _err := client.ListFlowControlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEcuByRegionWithOptions(request *ListEcuByRegionRequest, runtime *util.RuntimeOptions) (_result *ListEcuByRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListEcuByRegionResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListEcuByRegion"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/resource/ecu_list"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEcuByRegion(request *ListEcuByRegionRequest) (_result *ListEcuByRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEcuByRegionResponse{}
	_body, _err := client.ListEcuByRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDegradeControlsWithOptions(request *ListDegradeControlsRequest, runtime *util.RuntimeOptions) (_result *ListDegradeControlsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListDegradeControlsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListDegradeControls"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/app/degradeControls"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDegradeControls(request *ListDegradeControlsRequest) (_result *ListDegradeControlsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDegradeControlsResponse{}
	_body, _err := client.ListDegradeControlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConsumedServicesWithOptions(request *ListConsumedServicesRequest, runtime *util.RuntimeOptions) (_result *ListConsumedServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListConsumedServicesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListConsumedServices"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/service/listConsumedServices"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConsumedServices(request *ListConsumedServicesRequest) (_result *ListConsumedServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConsumedServicesResponse{}
	_body, _err := client.ListConsumedServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConfigCentersWithOptions(request *ListConfigCentersRequest, runtime *util.RuntimeOptions) (_result *ListConfigCentersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListConfigCentersResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListConfigCenters"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/configCenters"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConfigCenters(request *ListConfigCentersRequest) (_result *ListConfigCentersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConfigCentersResponse{}
	_body, _err := client.ListConfigCentersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterMembersWithOptions(request *ListClusterMembersRequest, runtime *util.RuntimeOptions) (_result *ListClusterMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListClusterMembersResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListClusterMembers"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/resource/cluster_member_list"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterMembers(request *ListClusterMembersRequest) (_result *ListClusterMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterMembersResponse{}
	_body, _err := client.ListClusterMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAuthorityWithOptions(request *ListAuthorityRequest, runtime *util.RuntimeOptions) (_result *ListAuthorityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListAuthorityResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListAuthority"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/account/authority_list"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAuthority(request *ListAuthorityRequest) (_result *ListAuthorityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAuthorityResponse{}
	_body, _err := client.ListAuthorityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAliyunRegionWithOptions(request *ListAliyunRegionRequest, runtime *util.RuntimeOptions) (_result *ListAliyunRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListAliyunRegionResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListAliyunRegion"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/resource/region_list"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAliyunRegion(request *ListAliyunRegionRequest) (_result *ListAliyunRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAliyunRegionResponse{}
	_body, _err := client.ListAliyunRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertServiceGroupWithOptions(request *InsertServiceGroupRequest, runtime *util.RuntimeOptions) (_result *InsertServiceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InsertServiceGroupResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("InsertServiceGroup"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/service/serviceGroups"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertServiceGroup(request *InsertServiceGroupRequest) (_result *InsertServiceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertServiceGroupResponse{}
	_body, _err := client.InsertServiceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertRoleWithOptions(request *InsertRoleRequest, runtime *util.RuntimeOptions) (_result *InsertRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InsertRoleResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("InsertRole"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/account/create_role"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertRole(request *InsertRoleRequest) (_result *InsertRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertRoleResponse{}
	_body, _err := client.InsertRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertOrUpdateRegionWithOptions(request *InsertOrUpdateRegionRequest, runtime *util.RuntimeOptions) (_result *InsertOrUpdateRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InsertOrUpdateRegionResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("InsertOrUpdateRegion"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/user_region_def"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertOrUpdateRegion(request *InsertOrUpdateRegionRequest) (_result *InsertOrUpdateRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertOrUpdateRegionResponse{}
	_body, _err := client.InsertOrUpdateRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertConfigCenterWithOptions(request *InsertConfigCenterRequest, runtime *util.RuntimeOptions) (_result *InsertConfigCenterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InsertConfigCenterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("InsertConfigCenter"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/configCenter"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertConfigCenter(request *InsertConfigCenterRequest) (_result *InsertConfigCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertConfigCenterResponse{}
	_body, _err := client.InsertConfigCenterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertFlowControlWithOptions(request *InsertFlowControlRequest, runtime *util.RuntimeOptions) (_result *InsertFlowControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InsertFlowControlResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("InsertFlowControl"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/flowControl"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertFlowControl(request *InsertFlowControlRequest) (_result *InsertFlowControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertFlowControlResponse{}
	_body, _err := client.InsertFlowControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertDeployGroupWithOptions(request *InsertDeployGroupRequest, runtime *util.RuntimeOptions) (_result *InsertDeployGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InsertDeployGroupResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("InsertDeployGroup"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/deploy_group"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertDeployGroup(request *InsertDeployGroupRequest) (_result *InsertDeployGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertDeployGroupResponse{}
	_body, _err := client.InsertDeployGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertDegradeControlWithOptions(request *InsertDegradeControlRequest, runtime *util.RuntimeOptions) (_result *InsertDegradeControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InsertDegradeControlResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("InsertDegradeControl"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/degradeControl"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertDegradeControl(request *InsertDegradeControlRequest) (_result *InsertDegradeControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertDegradeControlResponse{}
	_body, _err := client.InsertDegradeControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertClusterWithOptions(request *InsertClusterRequest, runtime *util.RuntimeOptions) (_result *InsertClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InsertClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("InsertCluster"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/resource/cluster"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertCluster(request *InsertClusterRequest) (_result *InsertClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertClusterResponse{}
	_body, _err := client.InsertClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJvmConfigurationWithOptions(request *GetJvmConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetJvmConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetJvmConfigurationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GetJvmConfiguration"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/app/app_jvm_config"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJvmConfiguration(request *GetJvmConfigurationRequest) (_result *GetJvmConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJvmConfigurationResponse{}
	_body, _err := client.GetJvmConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetContainerConfigurationWithOptions(request *GetContainerConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetContainerConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetContainerConfigurationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GetContainerConfiguration"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/app/container_config"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetContainerConfiguration(request *GetContainerConfigurationRequest) (_result *GetContainerConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetContainerConfigurationResponse{}
	_body, _err := client.GetContainerConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetApplicationWithOptions(request *GetApplicationRequest, runtime *util.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GetApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/app/app_info"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApplication(request *GetApplicationRequest) (_result *GetApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApplicationResponse{}
	_body, _err := client.GetApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableFlowControlWithOptions(request *EnableFlowControlRequest, runtime *util.RuntimeOptions) (_result *EnableFlowControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EnableFlowControlResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("EnableFlowControl"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/pop/v5/flowcontrol/enable"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableFlowControl(request *EnableFlowControlRequest) (_result *EnableFlowControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableFlowControlResponse{}
	_body, _err := client.EnableFlowControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableDegradeControlWithOptions(request *EnableDegradeControlRequest, runtime *util.RuntimeOptions) (_result *EnableDegradeControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EnableDegradeControlResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("EnableDegradeControl"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/pop/v5/degradecontrol/enable"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableDegradeControl(request *EnableDegradeControlRequest) (_result *EnableDegradeControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableDegradeControlResponse{}
	_body, _err := client.EnableDegradeControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableFlowControlWithOptions(request *DisableFlowControlRequest, runtime *util.RuntimeOptions) (_result *DisableFlowControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DisableFlowControlResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DisableFlowControl"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/pop/v5/flowcontrol/disable"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableFlowControl(request *DisableFlowControlRequest) (_result *DisableFlowControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableFlowControlResponse{}
	_body, _err := client.DisableFlowControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableDegradeControlWithOptions(request *DisableDegradeControlRequest, runtime *util.RuntimeOptions) (_result *DisableDegradeControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DisableDegradeControlResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DisableDegradeControl"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/pop/v5/degradecontrol/disable"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableDegradeControl(request *DisableDegradeControlRequest) (_result *DisableDegradeControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableDegradeControlResponse{}
	_body, _err := client.DisableDegradeControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserDefineRegionWithOptions(request *DeleteUserDefineRegionRequest, runtime *util.RuntimeOptions) (_result *DeleteUserDefineRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteUserDefineRegionResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteUserDefineRegion"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v5/user_region_def"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserDefineRegion(request *DeleteUserDefineRegionRequest) (_result *DeleteUserDefineRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserDefineRegionResponse{}
	_body, _err := client.DeleteUserDefineRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceGroupWithOptions(request *DeleteServiceGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteServiceGroupResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteServiceGroup"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v5/service/serviceGroups"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServiceGroup(request *DeleteServiceGroupRequest) (_result *DeleteServiceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServiceGroupResponse{}
	_body, _err := client.DeleteServiceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRoleWithOptions(request *DeleteRoleRequest, runtime *util.RuntimeOptions) (_result *DeleteRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteRoleResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteRole"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/account/delete_role"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRole(request *DeleteRoleRequest) (_result *DeleteRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRoleResponse{}
	_body, _err := client.DeleteRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowControlWithOptions(request *DeleteFlowControlRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteFlowControlResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteFlowControl"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v5/flowControl"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlowControl(request *DeleteFlowControlRequest) (_result *DeleteFlowControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowControlResponse{}
	_body, _err := client.DeleteFlowControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEcuWithOptions(request *DeleteEcuRequest, runtime *util.RuntimeOptions) (_result *DeleteEcuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteEcuResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteEcu"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/resource/delete_ecu"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEcu(request *DeleteEcuRequest) (_result *DeleteEcuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEcuResponse{}
	_body, _err := client.DeleteEcuWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeployGroupWithOptions(request *DeleteDeployGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteDeployGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDeployGroupResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteDeployGroup"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v5/deploy_group"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDeployGroup(request *DeleteDeployGroupRequest) (_result *DeleteDeployGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDeployGroupResponse{}
	_body, _err := client.DeleteDeployGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDegradeControlWithOptions(request *DeleteDegradeControlRequest, runtime *util.RuntimeOptions) (_result *DeleteDegradeControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDegradeControlResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteDegradeControl"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v5/degradeControl"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDegradeControl(request *DeleteDegradeControlRequest) (_result *DeleteDegradeControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDegradeControlResponse{}
	_body, _err := client.DeleteDegradeControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConfigCenterWithOptions(request *DeleteConfigCenterRequest, runtime *util.RuntimeOptions) (_result *DeleteConfigCenterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteConfigCenterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteConfigCenter"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v5/configCenter"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConfigCenter(request *DeleteConfigCenterRequest) (_result *DeleteConfigCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConfigCenterResponse{}
	_body, _err := client.DeleteConfigCenterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterMemberWithOptions(request *DeleteClusterMemberRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteClusterMemberResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteClusterMember"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v5/resource/cluster_member"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteClusterMember(request *DeleteClusterMemberRequest) (_result *DeleteClusterMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterMemberResponse{}
	_body, _err := client.DeleteClusterMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteCluster"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v5/resource/cluster"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindSlbWithOptions(request *BindSlbRequest, runtime *util.RuntimeOptions) (_result *BindSlbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &BindSlbResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("BindSlb"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/app/bind_slb_json"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindSlb(request *BindSlbRequest) (_result *BindSlbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindSlbResponse{}
	_body, _err := client.BindSlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthorizeRoleWithOptions(request *AuthorizeRoleRequest, runtime *util.RuntimeOptions) (_result *AuthorizeRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AuthorizeRoleResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("AuthorizeRole"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/account/authorize_role"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthorizeRole(request *AuthorizeRoleRequest) (_result *AuthorizeRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeRoleResponse{}
	_body, _err := client.AuthorizeRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthorizeResourceGroupWithOptions(request *AuthorizeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *AuthorizeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AuthorizeResourceGroupResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("AuthorizeResourceGroup"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/account/authorize_res_group"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthorizeResourceGroup(request *AuthorizeResourceGroupRequest) (_result *AuthorizeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeResourceGroupResponse{}
	_body, _err := client.AuthorizeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthorizeApplicationWithOptions(request *AuthorizeApplicationRequest, runtime *util.RuntimeOptions) (_result *AuthorizeApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AuthorizeApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("AuthorizeApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/account/authorize_app"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthorizeApplication(request *AuthorizeApplicationRequest) (_result *AuthorizeApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeApplicationResponse{}
	_body, _err := client.AuthorizeApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpcWithOptions(request *ListVpcRequest, runtime *util.RuntimeOptions) (_result *ListVpcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListVpcResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListVpc"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/pop/v5/vpc_list"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpc(request *ListVpcRequest) (_result *ListVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVpcResponse{}
	_body, _err := client.ListVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScaleOutApplicationWithOptions(request *ScaleOutApplicationRequest, runtime *util.RuntimeOptions) (_result *ScaleOutApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ScaleOutApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ScaleOutApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/changeorder/co_scale_out"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScaleOutApplication(request *ScaleOutApplicationRequest) (_result *ScaleOutApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScaleOutApplicationResponse{}
	_body, _err := client.ScaleOutApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScaleInApplicationWithOptions(request *ScaleInApplicationRequest, runtime *util.RuntimeOptions) (_result *ScaleInApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ScaleInApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ScaleInApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/changeorder/co_scale_in"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScaleInApplication(request *ScaleInApplicationRequest) (_result *ScaleInApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScaleInApplicationResponse{}
	_body, _err := client.ScaleInApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployApplicationWithOptions(request *DeployApplicationRequest, runtime *util.RuntimeOptions) (_result *DeployApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeployApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeployApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/changeorder/co_deploy"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployApplication(request *DeployApplicationRequest) (_result *DeployApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployApplicationResponse{}
	_body, _err := client.DeployApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertApplicationWithOptions(request *InsertApplicationRequest, runtime *util.RuntimeOptions) (_result *InsertApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InsertApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("InsertApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/changeorder/co_create_app"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertApplication(request *InsertApplicationRequest) (_result *InsertApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertApplicationResponse{}
	_body, _err := client.InsertApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, runtime *util.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/pop/v5/changeorder/co_delete_app"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApplication(request *DeleteApplicationRequest) (_result *DeleteApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DeleteApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateContainerWithOptions(request *UpdateContainerRequest, runtime *util.RuntimeOptions) (_result *UpdateContainerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateContainerResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("UpdateContainer"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/changeorder/co_update_container"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateContainer(request *UpdateContainerRequest) (_result *UpdateContainerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateContainerResponse{}
	_body, _err := client.UpdateContainerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopApplicationWithOptions(request *StopApplicationRequest, runtime *util.RuntimeOptions) (_result *StopApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("StopApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/changeorder/co_stop"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopApplication(request *StopApplicationRequest) (_result *StopApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopApplicationResponse{}
	_body, _err := client.StopApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartApplicationWithOptions(request *StartApplicationRequest, runtime *util.RuntimeOptions) (_result *StartApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("StartApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/changeorder/co_start"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartApplication(request *StartApplicationRequest) (_result *StartApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartApplicationResponse{}
	_body, _err := client.StartApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetApplicationWithOptions(request *ResetApplicationRequest, runtime *util.RuntimeOptions) (_result *ResetApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ResetApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ResetApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/changeorder/co_reset"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetApplication(request *ResetApplicationRequest) (_result *ResetApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetApplicationResponse{}
	_body, _err := client.ResetApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeployGroupWithOptions(request *ListDeployGroupRequest, runtime *util.RuntimeOptions) (_result *ListDeployGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListDeployGroupResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListDeployGroup"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/app/deploy_group_list"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeployGroup(request *ListDeployGroupRequest) (_result *ListDeployGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeployGroupResponse{}
	_body, _err := client.ListDeployGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBuildPackWithOptions(request *ListBuildPackRequest, runtime *util.RuntimeOptions) (_result *ListBuildPackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListBuildPackResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListBuildPack"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/app/build_pack_list"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBuildPack(request *ListBuildPackRequest) (_result *ListBuildPackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBuildPackResponse{}
	_body, _err := client.ListBuildPackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplicationEcuWithOptions(request *ListApplicationEcuRequest, runtime *util.RuntimeOptions) (_result *ListApplicationEcuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListApplicationEcuResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListApplicationEcu"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/resource/ecu_list"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplicationEcu(request *ListApplicationEcuRequest) (_result *ListApplicationEcuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationEcuResponse{}
	_body, _err := client.ListApplicationEcuWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetChangeOrderInfoWithOptions(request *GetChangeOrderInfoRequest, runtime *util.RuntimeOptions) (_result *GetChangeOrderInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetChangeOrderInfoResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GetChangeOrderInfo"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/changeorder/change_order_info"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetChangeOrderInfo(request *GetChangeOrderInfoRequest) (_result *GetChangeOrderInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetChangeOrderInfoResponse{}
	_body, _err := client.GetChangeOrderInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterWithOptions(request *ListClusterRequest, runtime *util.RuntimeOptions) (_result *ListClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListCluster"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/resource/cluster_list"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCluster(request *ListClusterRequest) (_result *ListClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterResponse{}
	_body, _err := client.ListClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplicationWithOptions(request *ListApplicationRequest, runtime *util.RuntimeOptions) (_result *ListApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListApplicationResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("ListApplication"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/pop/v5/app/app_list"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplication(request *ListApplicationRequest) (_result *ListApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationResponse{}
	_body, _err := client.ListApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
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
