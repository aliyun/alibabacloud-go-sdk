// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridMonitorNamespaceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHybridMonitorNamespaceListResponseBody
	GetCode() *string
	SetDescribeHybridMonitorNamespace(v []*DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) *DescribeHybridMonitorNamespaceListResponseBody
	GetDescribeHybridMonitorNamespace() []*DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace
	SetMessage(v string) *DescribeHybridMonitorNamespaceListResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeHybridMonitorNamespaceListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHybridMonitorNamespaceListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHybridMonitorNamespaceListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeHybridMonitorNamespaceListResponseBody
	GetSuccess() *string
	SetTotal(v int32) *DescribeHybridMonitorNamespaceListResponseBody
	GetTotal() *int32
}

type DescribeHybridMonitorNamespaceListResponseBody struct {
	// The response code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the namespaces.
	DescribeHybridMonitorNamespace []*DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace `json:"DescribeHybridMonitorNamespace,omitempty" xml:"DescribeHybridMonitorNamespace,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// Specified parameter PageSize is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1EC450A4-3221-5148-B77E-2827576CFE48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeHybridMonitorNamespaceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorNamespaceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) GetDescribeHybridMonitorNamespace() []*DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace {
	return s.DescribeHybridMonitorNamespace
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) SetCode(v string) *DescribeHybridMonitorNamespaceListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) SetDescribeHybridMonitorNamespace(v []*DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) *DescribeHybridMonitorNamespaceListResponseBody {
	s.DescribeHybridMonitorNamespace = v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) SetMessage(v string) *DescribeHybridMonitorNamespaceListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) SetPageNumber(v int32) *DescribeHybridMonitorNamespaceListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) SetPageSize(v int32) *DescribeHybridMonitorNamespaceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) SetRequestId(v string) *DescribeHybridMonitorNamespaceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) SetSuccess(v string) *DescribeHybridMonitorNamespaceListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) SetTotal(v int32) *DescribeHybridMonitorNamespaceListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace struct {
	// The configuration details of metric import tasks for Alibaba Cloud services.
	AliyunProductMetricList []*DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricList `json:"AliyunProductMetricList,omitempty" xml:"AliyunProductMetricList,omitempty" type:"Repeated"`
	// The timestamp that was generated when the namespace was created.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1652682744000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the namespace.
	//
	// example:
	//
	// aliyun-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The details of the data retention period.
	Detail *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	// The ID of the namespace.
	//
	// example:
	//
	// 3****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the namespace is deleted. Valid values:
	//
	// 	- 0: The namespace is not deleted.
	//
	// 	- 1: The namespace is deleted.
	//
	// example:
	//
	// 0
	IsDelete *int32 `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	// The timestamp that was generated when the namespace was last modified.
	//
	// example:
	//
	// 1652682744000
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// aliyun-test
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The storage scheme of metric data. Valid values:
	//
	// 	- m_prom_user: The metric data is stored in Simple Log Service.
	//
	// 	- m_prom_pool: The metric data is stored in the storage space provided by CloudMonitor.
	//
	// example:
	//
	// m_prom_user
	NamespaceType *string `json:"NamespaceType,omitempty" xml:"NamespaceType,omitempty"`
	// The number of metric import tasks for third-party services.
	//
	// example:
	//
	// 0
	NotAliyunTaskNumber *int64 `json:"NotAliyunTaskNumber,omitempty" xml:"NotAliyunTaskNumber,omitempty"`
}

func (s DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) GetAliyunProductMetricList() []*DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricList {
	return s.AliyunProductMetricList
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) GetDescription() *string {
	return s.Description
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) GetDetail() *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail {
	return s.Detail
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) GetId() *string {
	return s.Id
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) GetIsDelete() *int32 {
	return s.IsDelete
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) GetNamespaceType() *string {
	return s.NamespaceType
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) GetNotAliyunTaskNumber() *int64 {
	return s.NotAliyunTaskNumber
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) SetAliyunProductMetricList(v []*DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricList) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace {
	s.AliyunProductMetricList = v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) SetCreateTime(v string) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace {
	s.CreateTime = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) SetDescription(v string) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace {
	s.Description = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) SetDetail(v *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace {
	s.Detail = v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) SetId(v string) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace {
	s.Id = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) SetIsDelete(v int32) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace {
	s.IsDelete = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) SetModifyTime(v string) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace {
	s.ModifyTime = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) SetNamespace(v string) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace {
	s.Namespace = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) SetNamespaceType(v string) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace {
	s.NamespaceType = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) SetNotAliyunTaskNumber(v int64) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace {
	s.NotAliyunTaskNumber = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespace) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricList struct {
	// The namespaces.
	NamespaceList []*DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceList `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
	// The account that is used to create the namespace.
	//
	// example:
	//
	// 120886317861****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The configuration file of the Alibaba Cloud service that you want to monitor by using Hybrid Cloud Monitoring.
	//
	// 	- namespace: the namespace of the Alibaba Cloud service.
	//
	// 	- metric_list: the metrics of the Alibaba Cloud service.
	//
	// 	- dimension: the resources of the Alibaba Cloud service that you want to monitor by using Hybrid Cloud Monitoring. If you do not specify a dimension, all resources of the Alibaba Cloud service are monitored.
	//
	// example:
	//
	// products:- namespace: acs_ecs_dashboard metric_info: - metric_list: - cpu_total dimension: \\"\\"
	YAMLConfig *string `json:"YAMLConfig,omitempty" xml:"YAMLConfig,omitempty"`
}

func (s DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricList) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricList) GetNamespaceList() []*DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceList {
	return s.NamespaceList
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricList) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricList) GetYAMLConfig() *string {
	return s.YAMLConfig
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricList) SetNamespaceList(v []*DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceList) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricList {
	s.NamespaceList = v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricList) SetUserId(v int64) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricList {
	s.UserId = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricList) SetYAMLConfig(v string) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricList {
	s.YAMLConfig = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricList) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceList struct {
	// The metrics for the Alibaba Cloud service.
	MetricList []*DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceListMetricList `json:"MetricList,omitempty" xml:"MetricList,omitempty" type:"Repeated"`
	// The namespace for the Alibaba Cloud service.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceList) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceList) GetMetricList() []*DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceListMetricList {
	return s.MetricList
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceList) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceList) SetMetricList(v []*DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceListMetricList) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceList {
	s.MetricList = v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceList) SetNamespace(v string) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceList {
	s.Namespace = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceList) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceListMetricList struct {
	// The metrics.
	List []*string `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The collection period of the metric.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 60
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceListMetricList) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceListMetricList) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceListMetricList) GetList() []*string {
	return s.List
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceListMetricList) GetPeriod() *int64 {
	return s.Period
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceListMetricList) SetList(v []*string) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceListMetricList {
	s.List = v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceListMetricList) SetPeriod(v int64) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceListMetricList {
	s.Period = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceAliyunProductMetricListNamespaceListMetricList) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail struct {
	// The region where the metric data is stored.
	//
	// >  This parameter is returned if you select `m_prom_user` for `NamespaceType` when you create a namespace.
	//
	// example:
	//
	// cn-hangzhou
	NamespaceRegion      *string `json:"NamespaceRegion,omitempty" xml:"NamespaceRegion,omitempty"`
	PrometheusInstanceId *string `json:"PrometheusInstanceId,omitempty" xml:"PrometheusInstanceId,omitempty"`
	// The project where the metric data is located.
	//
	// >  This parameter is returned if you select `m_prom_user` for `NamespaceType` when you create a namespace.
	//
	// example:
	//
	// cms-hybrid-120886317861****-cn-hangzhou-a83d
	SLSProject *string `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
	// The data retention period. Valid values:
	//
	// 	- cms.s1.large (Retention Period 15 Days)
	//
	// 	- cms.s1.xlarge (Retention Period 32 Days)
	//
	// 	- cms.s1.2xlarge (Retention Period 63 Days)
	//
	// 	- cms.s1.3xlarge (Retention Period 93 Days)
	//
	// 	- cms.s1.6xlarge (Retention Period 185 Days)
	//
	// 	- cms.s1.12xlarge (Retention Period 367 Days)
	//
	// example:
	//
	// cms.s1.3xlarge
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail) GetNamespaceRegion() *string {
	return s.NamespaceRegion
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail) GetPrometheusInstanceId() *string {
	return s.PrometheusInstanceId
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail) GetSLSProject() *string {
	return s.SLSProject
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail) GetSpec() *string {
	return s.Spec
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail) SetNamespaceRegion(v string) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail {
	s.NamespaceRegion = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail) SetPrometheusInstanceId(v string) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail {
	s.PrometheusInstanceId = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail) SetSLSProject(v string) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail {
	s.SLSProject = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail) SetSpec(v string) *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail {
	s.Spec = &v
	return s
}

func (s *DescribeHybridMonitorNamespaceListResponseBodyDescribeHybridMonitorNamespaceDetail) Validate() error {
	return dara.Validate(s)
}
