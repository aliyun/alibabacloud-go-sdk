// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListConsumerGroupsResponseBody
	GetCode() *string
	SetData(v *ListConsumerGroupsResponseBodyData) *ListConsumerGroupsResponseBody
	GetData() *ListConsumerGroupsResponseBodyData
	SetDynamicCode(v string) *ListConsumerGroupsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListConsumerGroupsResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListConsumerGroupsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListConsumerGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListConsumerGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListConsumerGroupsResponseBody
	GetSuccess() *bool
}

type ListConsumerGroupsResponseBody struct {
	// Error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListConsumerGroupsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Dynamic error code.
	//
	// example:
	//
	// xxx
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// xxx
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// Parameter InstanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, each request has a unique ID that can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// 5503A460-98ED-5543-92CF-4853DE28****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the execution was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListConsumerGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListConsumerGroupsResponseBody) GetData() *ListConsumerGroupsResponseBodyData {
	return s.Data
}

func (s *ListConsumerGroupsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListConsumerGroupsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListConsumerGroupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListConsumerGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConsumerGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConsumerGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListConsumerGroupsResponseBody) SetCode(v string) *ListConsumerGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetData(v *ListConsumerGroupsResponseBodyData) *ListConsumerGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetDynamicCode(v string) *ListConsumerGroupsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetDynamicMessage(v string) *ListConsumerGroupsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetHttpStatusCode(v int32) *ListConsumerGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetMessage(v string) *ListConsumerGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetRequestId(v string) *ListConsumerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetSuccess(v bool) *ListConsumerGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConsumerGroupsResponseBodyData struct {
	// The consumer groups.
	List []*ListConsumerGroupsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// Current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total number of returned results.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListConsumerGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsResponseBodyData) GetList() []*ListConsumerGroupsResponseBodyDataList {
	return s.List
}

func (s *ListConsumerGroupsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListConsumerGroupsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListConsumerGroupsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListConsumerGroupsResponseBodyData) SetList(v []*ListConsumerGroupsResponseBodyDataList) *ListConsumerGroupsResponseBodyData {
	s.List = v
	return s
}

func (s *ListConsumerGroupsResponseBodyData) SetPageNumber(v int64) *ListConsumerGroupsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyData) SetPageSize(v int64) *ListConsumerGroupsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyData) SetTotalCount(v int64) *ListConsumerGroupsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyData) Validate() error {
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

type ListConsumerGroupsResponseBodyDataList struct {
	// Consumer group ID.
	//
	// example:
	//
	// GID-TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// Creation time of the consumer group.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The maximum TPS for message sending.
	//
	// example:
	//
	// 1000
	MaxReceiveTps *int64 `json:"maxReceiveTps,omitempty" xml:"maxReceiveTps,omitempty"`
	// example:
	//
	// LITE_SELECTIVE
	MessageModel *string `json:"messageModel,omitempty" xml:"messageModel,omitempty"`
	// The region ID to which the instance belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Remark information of the consumer group.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// Status of the consumer group.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// test1
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
	// Last update time of the consumer group.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListConsumerGroupsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsResponseBodyDataList) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *ListConsumerGroupsResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListConsumerGroupsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConsumerGroupsResponseBodyDataList) GetMaxReceiveTps() *int64 {
	return s.MaxReceiveTps
}

func (s *ListConsumerGroupsResponseBodyDataList) GetMessageModel() *string {
	return s.MessageModel
}

func (s *ListConsumerGroupsResponseBodyDataList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListConsumerGroupsResponseBodyDataList) GetRemark() *string {
	return s.Remark
}

func (s *ListConsumerGroupsResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *ListConsumerGroupsResponseBodyDataList) GetTopicName() *string {
	return s.TopicName
}

func (s *ListConsumerGroupsResponseBodyDataList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListConsumerGroupsResponseBodyDataList) SetConsumerGroupId(v string) *ListConsumerGroupsResponseBodyDataList {
	s.ConsumerGroupId = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetCreateTime(v string) *ListConsumerGroupsResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetInstanceId(v string) *ListConsumerGroupsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetMaxReceiveTps(v int64) *ListConsumerGroupsResponseBodyDataList {
	s.MaxReceiveTps = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetMessageModel(v string) *ListConsumerGroupsResponseBodyDataList {
	s.MessageModel = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetRegionId(v string) *ListConsumerGroupsResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetRemark(v string) *ListConsumerGroupsResponseBodyDataList {
	s.Remark = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetStatus(v string) *ListConsumerGroupsResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetTopicName(v string) *ListConsumerGroupsResponseBodyDataList {
	s.TopicName = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetUpdateTime(v string) *ListConsumerGroupsResponseBodyDataList {
	s.UpdateTime = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
