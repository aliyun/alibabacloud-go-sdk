// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListQueuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkListQueuesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkListQueuesResponseBody
	GetCode() *string
	SetData(v *ClinkListQueuesResponseBodyData) *ClinkListQueuesResponseBody
	GetData() *ClinkListQueuesResponseBodyData
	SetMessage(v string) *ClinkListQueuesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkListQueuesResponseBody
	GetRequestId() *string
}

type ClinkListQueuesResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkListQueuesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9FF70B74-1B3C-44C1-ACDF-8DF962988F0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkListQueuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkListQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkListQueuesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkListQueuesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkListQueuesResponseBody) GetData() *ClinkListQueuesResponseBodyData {
	return s.Data
}

func (s *ClinkListQueuesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkListQueuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkListQueuesResponseBody) SetAccessDeniedDetail(v string) *ClinkListQueuesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkListQueuesResponseBody) SetCode(v string) *ClinkListQueuesResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkListQueuesResponseBody) SetData(v *ClinkListQueuesResponseBodyData) *ClinkListQueuesResponseBody {
	s.Data = v
	return s
}

func (s *ClinkListQueuesResponseBody) SetMessage(v string) *ClinkListQueuesResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkListQueuesResponseBody) SetRequestId(v string) *ClinkListQueuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkListQueuesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkListQueuesResponseBodyData struct {
	// 本次请求id
	//
	// example:
	//
	// xxx
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// queues
	Queues []*ClinkListQueuesResponseBodyDataQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
	// 总数
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ClinkListQueuesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkListQueuesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkListQueuesResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkListQueuesResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ClinkListQueuesResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ClinkListQueuesResponseBodyData) GetQueues() []*ClinkListQueuesResponseBodyDataQueues {
	return s.Queues
}

func (s *ClinkListQueuesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ClinkListQueuesResponseBodyData) SetClinkRequestId(v string) *ClinkListQueuesResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkListQueuesResponseBodyData) SetPageNumber(v int64) *ClinkListQueuesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ClinkListQueuesResponseBodyData) SetPageSize(v int64) *ClinkListQueuesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ClinkListQueuesResponseBodyData) SetQueues(v []*ClinkListQueuesResponseBodyDataQueues) *ClinkListQueuesResponseBodyData {
	s.Queues = v
	return s
}

func (s *ClinkListQueuesResponseBodyData) SetTotalCount(v int64) *ClinkListQueuesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ClinkListQueuesResponseBodyData) Validate() error {
	if s.Queues != nil {
		for _, item := range s.Queues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkListQueuesResponseBodyDataQueues struct {
	// 呼叫中心座席数
	//
	// example:
	//
	// 51
	CallClientNum *int64 `json:"CallClientNum,omitempty" xml:"CallClientNum,omitempty"`
	// 在线客服数
	//
	// example:
	//
	// 72
	ChatClientNum *int64 `json:"ChatClientNum,omitempty" xml:"ChatClientNum,omitempty"`
	// 队列 Id
	//
	// example:
	//
	// 3
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 队列名
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 全渠道座席数
	//
	// example:
	//
	// 87
	OmniClientNum *int64 `json:"OmniClientNum,omitempty" xml:"OmniClientNum,omitempty"`
	// 队列号
	//
	// example:
	//
	// 1221
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
}

func (s ClinkListQueuesResponseBodyDataQueues) String() string {
	return dara.Prettify(s)
}

func (s ClinkListQueuesResponseBodyDataQueues) GoString() string {
	return s.String()
}

func (s *ClinkListQueuesResponseBodyDataQueues) GetCallClientNum() *int64 {
	return s.CallClientNum
}

func (s *ClinkListQueuesResponseBodyDataQueues) GetChatClientNum() *int64 {
	return s.ChatClientNum
}

func (s *ClinkListQueuesResponseBodyDataQueues) GetId() *int64 {
	return s.Id
}

func (s *ClinkListQueuesResponseBodyDataQueues) GetName() *string {
	return s.Name
}

func (s *ClinkListQueuesResponseBodyDataQueues) GetOmniClientNum() *int64 {
	return s.OmniClientNum
}

func (s *ClinkListQueuesResponseBodyDataQueues) GetQno() *string {
	return s.Qno
}

func (s *ClinkListQueuesResponseBodyDataQueues) SetCallClientNum(v int64) *ClinkListQueuesResponseBodyDataQueues {
	s.CallClientNum = &v
	return s
}

func (s *ClinkListQueuesResponseBodyDataQueues) SetChatClientNum(v int64) *ClinkListQueuesResponseBodyDataQueues {
	s.ChatClientNum = &v
	return s
}

func (s *ClinkListQueuesResponseBodyDataQueues) SetId(v int64) *ClinkListQueuesResponseBodyDataQueues {
	s.Id = &v
	return s
}

func (s *ClinkListQueuesResponseBodyDataQueues) SetName(v string) *ClinkListQueuesResponseBodyDataQueues {
	s.Name = &v
	return s
}

func (s *ClinkListQueuesResponseBodyDataQueues) SetOmniClientNum(v int64) *ClinkListQueuesResponseBodyDataQueues {
	s.OmniClientNum = &v
	return s
}

func (s *ClinkListQueuesResponseBodyDataQueues) SetQno(v string) *ClinkListQueuesResponseBodyDataQueues {
	s.Qno = &v
	return s
}

func (s *ClinkListQueuesResponseBodyDataQueues) Validate() error {
	return dara.Validate(s)
}
