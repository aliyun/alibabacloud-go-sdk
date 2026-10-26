// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstancesResponseBody
	GetCode() *string
	SetData(v *ListInstancesResponseBodyData) *ListInstancesResponseBody
	GetData() *ListInstancesResponseBodyData
	SetDynamicCode(v string) *ListInstancesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListInstancesResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListInstancesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstancesResponseBody
	GetSuccess() *bool
}

type ListInstancesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// MissingPageNumber
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response result.
	Data *ListInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// PageNumber
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// pageNumber
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter pageNumber is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID. Each request has a unique ID that you can use to troubleshoot and locate issues.
	//
	// example:
	//
	// 84445A20-2B50-5306-A3C0-AF99FC1833C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstancesResponseBody) GetData() *ListInstancesResponseBodyData {
	return s.Data
}

func (s *ListInstancesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListInstancesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListInstancesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstancesResponseBody) SetCode(v string) *ListInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesResponseBody) SetData(v *ListInstancesResponseBodyData) *ListInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListInstancesResponseBody) SetDynamicCode(v string) *ListInstancesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetDynamicMessage(v string) *ListInstancesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListInstancesResponseBody) SetHttpStatusCode(v int32) *ListInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetMessage(v string) *ListInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v bool) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstancesResponseBodyData struct {
	// The paged data.
	List []*ListInstancesResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned results.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyData) GetList() []*ListInstancesResponseBodyDataList {
	return s.List
}

func (s *ListInstancesResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListInstancesResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListInstancesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstancesResponseBodyData) SetList(v []*ListInstancesResponseBodyDataList) *ListInstancesResponseBodyData {
	s.List = v
	return s
}

func (s *ListInstancesResponseBodyData) SetPageNumber(v int64) *ListInstancesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetPageSize(v int64) *ListInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetTotalCount(v int64) *ListInstancesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyDataList struct {
	// The commodity code of the ApsaraMQ for RocketMQ 5.x series instance, such as ons_rmqsub_public_cn.
	//
	// example:
	//
	// ons_rmqsub_public_cn
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The time when the instance version was last updated, in YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2022-08-01 00:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The time when the instance expires, in YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2022-09-01 00:00:00
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The number of consumer groups created in the instance.
	//
	// example:
	//
	// 10
	GroupCount *int64 `json:"groupCount,omitempty" xml:"groupCount,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// test instance
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - PayAsYouGo: pay-as-you-go (postpaid).
	//
	// - Subscription: subscription (prepaid).
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The product information.
	ProductInfo *ListInstancesResponseBodyDataListProductInfo `json:"productInfo,omitempty" xml:"productInfo,omitempty" type:"Struct"`
	// The ID of the region to which the instance belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The time when the instance was released, in YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2022-09-07 00:00:00
	ReleaseTime *string `json:"releaseTime,omitempty" xml:"releaseTime,omitempty"`
	// The remarks for the instance.
	//
	// example:
	//
	// This is remark for instance.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmx7caj******
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The primary series code of the instance. Valid values:
	//
	// - standard: Standard Edition.
	//
	// - ultimate: Platinum Edition.
	//
	// - professional: Professional Edition.
	//
	// - shared: Shared Edition.
	//
	// - dedicated: Dedicated Edition.
	//
	// example:
	//
	// standard
	SeriesCode *string `json:"seriesCode,omitempty" xml:"seriesCode,omitempty"`
	// The service code of the instance. The service code for ApsaraMQ for RocketMQ is rmq.
	//
	// example:
	//
	// rmq
	ServiceCode *string `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	// The time when the instance was created, in YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2022-08-01 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The instance status. Valid values:
	//
	// - RELEASED: Released.
	//
	// - RUNNING: Running.
	//
	// - STOPPED: Stopped.
	//
	// - CHANGING: Changing.
	//
	// - CREATING: Creating.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The sub-series code of the instance. Valid values:
	//
	// - cluster_ha: cluster high-availability edition.
	//
	// - single_node: single-node edition.
	//
	// - serverless: serverless edition.
	//
	// example:
	//
	// cluster_ha
	SubSeriesCode *string `json:"subSeriesCode,omitempty" xml:"subSeriesCode,omitempty"`
	// The resource tag list.
	Tags []*ListInstancesResponseBodyDataListTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The number of topics created in the instance.
	//
	// example:
	//
	// 20
	TopicCount *int64 `json:"topicCount,omitempty" xml:"topicCount,omitempty"`
	// The time when the instance was last modified, in YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2022-08-02 00:00:00
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The ID of the user to whom the instance belongs.
	//
	// example:
	//
	// 6W0xz2uPfiwp****
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListInstancesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataList) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListInstancesResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListInstancesResponseBodyDataList) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListInstancesResponseBodyDataList) GetGroupCount() *int64 {
	return s.GroupCount
}

func (s *ListInstancesResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyDataList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesResponseBodyDataList) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListInstancesResponseBodyDataList) GetProductInfo() *ListInstancesResponseBodyDataListProductInfo {
	return s.ProductInfo
}

func (s *ListInstancesResponseBodyDataList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesResponseBodyDataList) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *ListInstancesResponseBodyDataList) GetRemark() *string {
	return s.Remark
}

func (s *ListInstancesResponseBodyDataList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesResponseBodyDataList) GetSeriesCode() *string {
	return s.SeriesCode
}

func (s *ListInstancesResponseBodyDataList) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *ListInstancesResponseBodyDataList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListInstancesResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyDataList) GetSubSeriesCode() *string {
	return s.SubSeriesCode
}

func (s *ListInstancesResponseBodyDataList) GetTags() []*ListInstancesResponseBodyDataListTags {
	return s.Tags
}

func (s *ListInstancesResponseBodyDataList) GetTopicCount() *int64 {
	return s.TopicCount
}

func (s *ListInstancesResponseBodyDataList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListInstancesResponseBodyDataList) GetUserId() *string {
	return s.UserId
}

func (s *ListInstancesResponseBodyDataList) SetCommodityCode(v string) *ListInstancesResponseBodyDataList {
	s.CommodityCode = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetCreateTime(v string) *ListInstancesResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetExpireTime(v string) *ListInstancesResponseBodyDataList {
	s.ExpireTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetGroupCount(v int64) *ListInstancesResponseBodyDataList {
	s.GroupCount = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetInstanceId(v string) *ListInstancesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetInstanceName(v string) *ListInstancesResponseBodyDataList {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetPaymentType(v string) *ListInstancesResponseBodyDataList {
	s.PaymentType = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetProductInfo(v *ListInstancesResponseBodyDataListProductInfo) *ListInstancesResponseBodyDataList {
	s.ProductInfo = v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetRegionId(v string) *ListInstancesResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetReleaseTime(v string) *ListInstancesResponseBodyDataList {
	s.ReleaseTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetRemark(v string) *ListInstancesResponseBodyDataList {
	s.Remark = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetResourceGroupId(v string) *ListInstancesResponseBodyDataList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetSeriesCode(v string) *ListInstancesResponseBodyDataList {
	s.SeriesCode = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetServiceCode(v string) *ListInstancesResponseBodyDataList {
	s.ServiceCode = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetStartTime(v string) *ListInstancesResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetStatus(v string) *ListInstancesResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetSubSeriesCode(v string) *ListInstancesResponseBodyDataList {
	s.SubSeriesCode = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetTags(v []*ListInstancesResponseBodyDataListTags) *ListInstancesResponseBodyDataList {
	s.Tags = v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetTopicCount(v int64) *ListInstancesResponseBodyDataList {
	s.TopicCount = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetUpdateTime(v string) *ListInstancesResponseBodyDataList {
	s.UpdateTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetUserId(v string) *ListInstancesResponseBodyDataList {
	s.UserId = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) Validate() error {
	if s.ProductInfo != nil {
		if err := s.ProductInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyDataListProductInfo struct {
	// The capacity mode. Valid values:
	//
	// - provisioned: reserved capacity with elastic scaling.
	//
	// - ondemand: pay-as-you-go.
	//
	// example:
	//
	// provisioned
	CapacityType *string `json:"capacityType,omitempty" xml:"capacityType,omitempty"`
	// Specifies whether the message trace feature is enabled.
	//
	// - true: The message trace feature is enabled.
	//
	// - false: The message trace feature is disabled.
	//
	// <props="intl">This parameter does not take effect. Regardless of whether this parameter is enabled, the instance supports the message trace feature by default.
	//
	// <props="china">This parameter takes effect only for Serverless instances. If the message trace feature is enabled, message trace fees are incurred. For more information about billing, see [Serverless billing description](https://help.aliyun.com/document_detail/2586280.html).
	//
	// <props="china">For pay-as-you-go and subscription instances, this parameter does not take effect. Regardless of whether this parameter is enabled, the message trace feature is supported by default.
	//
	// example:
	//
	// true
	TraceOn *bool `json:"traceOn,omitempty" xml:"traceOn,omitempty"`
}

func (s ListInstancesResponseBodyDataListProductInfo) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyDataListProductInfo) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataListProductInfo) GetCapacityType() *string {
	return s.CapacityType
}

func (s *ListInstancesResponseBodyDataListProductInfo) GetTraceOn() *bool {
	return s.TraceOn
}

func (s *ListInstancesResponseBodyDataListProductInfo) SetCapacityType(v string) *ListInstancesResponseBodyDataListProductInfo {
	s.CapacityType = &v
	return s
}

func (s *ListInstancesResponseBodyDataListProductInfo) SetTraceOn(v bool) *ListInstancesResponseBodyDataListProductInfo {
	s.TraceOn = &v
	return s
}

func (s *ListInstancesResponseBodyDataListProductInfo) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyDataListTags struct {
	// The key of the resource tag.
	//
	// example:
	//
	// key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the resource tag.
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListInstancesResponseBodyDataListTags) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyDataListTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataListTags) GetKey() *string {
	return s.Key
}

func (s *ListInstancesResponseBodyDataListTags) GetValue() *string {
	return s.Value
}

func (s *ListInstancesResponseBodyDataListTags) SetKey(v string) *ListInstancesResponseBodyDataListTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyDataListTags) SetValue(v string) *ListInstancesResponseBodyDataListTags {
	s.Value = &v
	return s
}

func (s *ListInstancesResponseBodyDataListTags) Validate() error {
	return dara.Validate(s)
}
