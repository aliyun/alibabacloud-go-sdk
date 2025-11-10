// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTracesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTracesResponseBody
	GetCode() *string
	SetData(v *ListTracesResponseBodyData) *ListTracesResponseBody
	GetData() *ListTracesResponseBodyData
	SetDynamicCode(v string) *ListTracesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListTracesResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListTracesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListTracesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTracesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTracesResponseBody
	GetSuccess() *bool
}

type ListTracesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	Data *ListTracesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// InstanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// EDFF77E1-1ED1-5389-B6A8-651D9433BBE5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTracesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTracesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTracesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTracesResponseBody) GetData() *ListTracesResponseBodyData {
	return s.Data
}

func (s *ListTracesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListTracesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListTracesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTracesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTracesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTracesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTracesResponseBody) SetCode(v string) *ListTracesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTracesResponseBody) SetData(v *ListTracesResponseBodyData) *ListTracesResponseBody {
	s.Data = v
	return s
}

func (s *ListTracesResponseBody) SetDynamicCode(v string) *ListTracesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListTracesResponseBody) SetDynamicMessage(v string) *ListTracesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListTracesResponseBody) SetHttpStatusCode(v int32) *ListTracesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTracesResponseBody) SetMessage(v string) *ListTracesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTracesResponseBody) SetRequestId(v string) *ListTracesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTracesResponseBody) SetSuccess(v bool) *ListTracesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTracesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTracesResponseBodyData struct {
	// Trace list.
	List []*ListTracesResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListTracesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTracesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTracesResponseBodyData) GetList() []*ListTracesResponseBodyDataList {
	return s.List
}

func (s *ListTracesResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListTracesResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTracesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTracesResponseBodyData) SetList(v []*ListTracesResponseBodyDataList) *ListTracesResponseBodyData {
	s.List = v
	return s
}

func (s *ListTracesResponseBodyData) SetPageNumber(v int64) *ListTracesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListTracesResponseBodyData) SetPageSize(v int64) *ListTracesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTracesResponseBodyData) SetTotalCount(v int64) *ListTracesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListTracesResponseBodyData) Validate() error {
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

type ListTracesResponseBodyDataList struct {
	// Message born time.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	BornTime *string `json:"bornTime,omitempty" xml:"bornTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Message id.
	//
	// example:
	//
	// 7F00000100207A4F0F294A938F7807AE
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// Message keys.
	MessageKeys []*string `json:"messageKeys,omitempty" xml:"messageKeys,omitempty" type:"Repeated"`
	// Message tag.
	//
	// example:
	//
	// xx
	MessageTag *string `json:"messageTag,omitempty" xml:"messageTag,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s ListTracesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListTracesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListTracesResponseBodyDataList) GetBornTime() *string {
	return s.BornTime
}

func (s *ListTracesResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTracesResponseBodyDataList) GetMessageId() *string {
	return s.MessageId
}

func (s *ListTracesResponseBodyDataList) GetMessageKeys() []*string {
	return s.MessageKeys
}

func (s *ListTracesResponseBodyDataList) GetMessageTag() *string {
	return s.MessageTag
}

func (s *ListTracesResponseBodyDataList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTracesResponseBodyDataList) GetTopicName() *string {
	return s.TopicName
}

func (s *ListTracesResponseBodyDataList) SetBornTime(v string) *ListTracesResponseBodyDataList {
	s.BornTime = &v
	return s
}

func (s *ListTracesResponseBodyDataList) SetInstanceId(v string) *ListTracesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListTracesResponseBodyDataList) SetMessageId(v string) *ListTracesResponseBodyDataList {
	s.MessageId = &v
	return s
}

func (s *ListTracesResponseBodyDataList) SetMessageKeys(v []*string) *ListTracesResponseBodyDataList {
	s.MessageKeys = v
	return s
}

func (s *ListTracesResponseBodyDataList) SetMessageTag(v string) *ListTracesResponseBodyDataList {
	s.MessageTag = &v
	return s
}

func (s *ListTracesResponseBodyDataList) SetRegionId(v string) *ListTracesResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListTracesResponseBodyDataList) SetTopicName(v string) *ListTracesResponseBodyDataList {
	s.TopicName = &v
	return s
}

func (s *ListTracesResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
