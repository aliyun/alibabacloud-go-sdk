// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTopicsResponseBody
	GetCode() *string
	SetData(v *ListTopicsResponseBodyData) *ListTopicsResponseBody
	GetData() *ListTopicsResponseBodyData
	SetDynamicCode(v string) *ListTopicsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListTopicsResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListTopicsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListTopicsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTopicsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTopicsResponseBody
	GetSuccess() *bool
}

type ListTopicsResponseBody struct {
	// Error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListTopicsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Dynamic error code.
	//
	// example:
	//
	// TopicName
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// Dynamic error message.
	//
	// example:
	//
	// topicName
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
	// The topic cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, each request has a unique ID that can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the execution was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTopicsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTopicsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTopicsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTopicsResponseBody) GetData() *ListTopicsResponseBodyData {
	return s.Data
}

func (s *ListTopicsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListTopicsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListTopicsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTopicsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTopicsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTopicsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTopicsResponseBody) SetCode(v string) *ListTopicsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTopicsResponseBody) SetData(v *ListTopicsResponseBodyData) *ListTopicsResponseBody {
	s.Data = v
	return s
}

func (s *ListTopicsResponseBody) SetDynamicCode(v string) *ListTopicsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListTopicsResponseBody) SetDynamicMessage(v string) *ListTopicsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListTopicsResponseBody) SetHttpStatusCode(v int32) *ListTopicsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTopicsResponseBody) SetMessage(v string) *ListTopicsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTopicsResponseBody) SetRequestId(v string) *ListTopicsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTopicsResponseBody) SetSuccess(v bool) *ListTopicsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTopicsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTopicsResponseBodyData struct {
	// The topics.
	List []*ListTopicsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
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
	// 3
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total number of results returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListTopicsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTopicsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTopicsResponseBodyData) GetList() []*ListTopicsResponseBodyDataList {
	return s.List
}

func (s *ListTopicsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListTopicsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTopicsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTopicsResponseBodyData) SetList(v []*ListTopicsResponseBodyDataList) *ListTopicsResponseBodyData {
	s.List = v
	return s
}

func (s *ListTopicsResponseBodyData) SetPageNumber(v int64) *ListTopicsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListTopicsResponseBodyData) SetPageSize(v int64) *ListTopicsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTopicsResponseBodyData) SetTotalCount(v int64) *ListTopicsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListTopicsResponseBodyData) Validate() error {
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

type ListTopicsResponseBodyDataList struct {
	// Creation time.
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
	// example:
	//
	// 20
	LiteTopicExpiration *int64 `json:"liteTopicExpiration,omitempty" xml:"liteTopicExpiration,omitempty"`
	// The maximum TPS for message sending.
	//
	// example:
	//
	// 1000
	MaxSendTps *int64 `json:"maxSendTps,omitempty" xml:"maxSendTps,omitempty"`
	// The type of messages in the topic.
	//
	// Valid values:
	//
	// 	- TRANSACTION: transactional messages
	//
	// 	- FIFO: ordered messages
	//
	// 	- DELAY: scheduled or delayed messages
	//
	// 	- NORMAL: normal messages
	//
	// example:
	//
	// NORMAL
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// The region ID to which the instance belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Remark information of the topic.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The topic status.
	//
	// Valid values:
	//
	// 	- RUNNING
	//
	// 	- CREATING
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Topic name.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
	// Last update time of the topic.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListTopicsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListTopicsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListTopicsResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTopicsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTopicsResponseBodyDataList) GetLiteTopicExpiration() *int64 {
	return s.LiteTopicExpiration
}

func (s *ListTopicsResponseBodyDataList) GetMaxSendTps() *int64 {
	return s.MaxSendTps
}

func (s *ListTopicsResponseBodyDataList) GetMessageType() *string {
	return s.MessageType
}

func (s *ListTopicsResponseBodyDataList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTopicsResponseBodyDataList) GetRemark() *string {
	return s.Remark
}

func (s *ListTopicsResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *ListTopicsResponseBodyDataList) GetTopicName() *string {
	return s.TopicName
}

func (s *ListTopicsResponseBodyDataList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListTopicsResponseBodyDataList) SetCreateTime(v string) *ListTopicsResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetInstanceId(v string) *ListTopicsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetLiteTopicExpiration(v int64) *ListTopicsResponseBodyDataList {
	s.LiteTopicExpiration = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetMaxSendTps(v int64) *ListTopicsResponseBodyDataList {
	s.MaxSendTps = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetMessageType(v string) *ListTopicsResponseBodyDataList {
	s.MessageType = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetRegionId(v string) *ListTopicsResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetRemark(v string) *ListTopicsResponseBodyDataList {
	s.Remark = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetStatus(v string) *ListTopicsResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetTopicName(v string) *ListTopicsResponseBodyDataList {
	s.TopicName = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetUpdateTime(v string) *ListTopicsResponseBodyDataList {
	s.UpdateTime = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
