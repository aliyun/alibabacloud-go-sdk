// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallChatListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CallChatListResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int64) *CallChatListResponseBody
	GetCode() *int64
	SetMessage(v string) *CallChatListResponseBody
	GetMessage() *string
	SetModel(v []*CallChatListResponseBodyModel) *CallChatListResponseBody
	GetModel() []*CallChatListResponseBodyModel
	SetRequestId(v string) *CallChatListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CallChatListResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *CallChatListResponseBody
	GetTimestamp() *int64
}

type CallChatListResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   []*CallChatListResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 81
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s CallChatListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CallChatListResponseBody) GoString() string {
	return s.String()
}

func (s *CallChatListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CallChatListResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CallChatListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CallChatListResponseBody) GetModel() []*CallChatListResponseBodyModel {
	return s.Model
}

func (s *CallChatListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CallChatListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CallChatListResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *CallChatListResponseBody) SetAccessDeniedDetail(v string) *CallChatListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CallChatListResponseBody) SetCode(v int64) *CallChatListResponseBody {
	s.Code = &v
	return s
}

func (s *CallChatListResponseBody) SetMessage(v string) *CallChatListResponseBody {
	s.Message = &v
	return s
}

func (s *CallChatListResponseBody) SetModel(v []*CallChatListResponseBodyModel) *CallChatListResponseBody {
	s.Model = v
	return s
}

func (s *CallChatListResponseBody) SetRequestId(v string) *CallChatListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CallChatListResponseBody) SetSuccess(v bool) *CallChatListResponseBody {
	s.Success = &v
	return s
}

func (s *CallChatListResponseBody) SetTimestamp(v int64) *CallChatListResponseBody {
	s.Timestamp = &v
	return s
}

func (s *CallChatListResponseBody) Validate() error {
	if s.Model != nil {
		for _, item := range s.Model {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CallChatListResponseBodyModel struct {
	// 聊天内容id
	//
	// example:
	//
	// 79
	ChatId *int64 `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// 说话内容
	//
	// example:
	//
	// 1
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 说话时间
	//
	// example:
	//
	// 2019-01-09 14:14:19
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 说话号码，其中0为AI,1-用户,2-坐席
	//
	// example:
	//
	// 4
	FromType *int64 `json:"FromType,omitempty" xml:"FromType,omitempty"`
}

func (s CallChatListResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s CallChatListResponseBodyModel) GoString() string {
	return s.String()
}

func (s *CallChatListResponseBodyModel) GetChatId() *int64 {
	return s.ChatId
}

func (s *CallChatListResponseBodyModel) GetContent() *string {
	return s.Content
}

func (s *CallChatListResponseBodyModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CallChatListResponseBodyModel) GetFromType() *int64 {
	return s.FromType
}

func (s *CallChatListResponseBodyModel) SetChatId(v int64) *CallChatListResponseBodyModel {
	s.ChatId = &v
	return s
}

func (s *CallChatListResponseBodyModel) SetContent(v string) *CallChatListResponseBodyModel {
	s.Content = &v
	return s
}

func (s *CallChatListResponseBodyModel) SetCreateTime(v string) *CallChatListResponseBodyModel {
	s.CreateTime = &v
	return s
}

func (s *CallChatListResponseBodyModel) SetFromType(v int64) *CallChatListResponseBodyModel {
	s.FromType = &v
	return s
}

func (s *CallChatListResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
