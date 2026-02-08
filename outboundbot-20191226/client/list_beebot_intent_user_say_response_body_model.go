// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBeebotIntentUserSayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeebotRequestId(v string) *ListBeebotIntentUserSayResponseBody
	GetBeebotRequestId() *string
	SetCode(v string) *ListBeebotIntentUserSayResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListBeebotIntentUserSayResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBeebotIntentUserSayResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListBeebotIntentUserSayResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBeebotIntentUserSayResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListBeebotIntentUserSayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBeebotIntentUserSayResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListBeebotIntentUserSayResponseBody
	GetTotalCount() *int32
	SetUserSays(v []*ListBeebotIntentUserSayResponseBodyUserSays) *ListBeebotIntentUserSayResponseBody
	GetUserSays() []*ListBeebotIntentUserSayResponseBodyUserSays
}

type ListBeebotIntentUserSayResponseBody struct {
	// Internal request ID
	//
	// example:
	//
	// D7BBFCDF-59B0-1ADA-BCA3-4B77F642DDFB
	BeebotRequestId *string `json:"BeebotRequestId,omitempty" xml:"BeebotRequestId,omitempty"`
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of items per page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 0956D5DA-0978-5DC9-94B0-C68527DA7475
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total count
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// List of user utterances
	UserSays []*ListBeebotIntentUserSayResponseBodyUserSays `json:"UserSays,omitempty" xml:"UserSays,omitempty" type:"Repeated"`
}

func (s ListBeebotIntentUserSayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBeebotIntentUserSayResponseBody) GoString() string {
	return s.String()
}

func (s *ListBeebotIntentUserSayResponseBody) GetBeebotRequestId() *string {
	return s.BeebotRequestId
}

func (s *ListBeebotIntentUserSayResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBeebotIntentUserSayResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBeebotIntentUserSayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBeebotIntentUserSayResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBeebotIntentUserSayResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBeebotIntentUserSayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBeebotIntentUserSayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBeebotIntentUserSayResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBeebotIntentUserSayResponseBody) GetUserSays() []*ListBeebotIntentUserSayResponseBodyUserSays {
	return s.UserSays
}

func (s *ListBeebotIntentUserSayResponseBody) SetBeebotRequestId(v string) *ListBeebotIntentUserSayResponseBody {
	s.BeebotRequestId = &v
	return s
}

func (s *ListBeebotIntentUserSayResponseBody) SetCode(v string) *ListBeebotIntentUserSayResponseBody {
	s.Code = &v
	return s
}

func (s *ListBeebotIntentUserSayResponseBody) SetHttpStatusCode(v int32) *ListBeebotIntentUserSayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBeebotIntentUserSayResponseBody) SetMessage(v string) *ListBeebotIntentUserSayResponseBody {
	s.Message = &v
	return s
}

func (s *ListBeebotIntentUserSayResponseBody) SetPageNumber(v int32) *ListBeebotIntentUserSayResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBeebotIntentUserSayResponseBody) SetPageSize(v int32) *ListBeebotIntentUserSayResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBeebotIntentUserSayResponseBody) SetRequestId(v string) *ListBeebotIntentUserSayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBeebotIntentUserSayResponseBody) SetSuccess(v bool) *ListBeebotIntentUserSayResponseBody {
	s.Success = &v
	return s
}

func (s *ListBeebotIntentUserSayResponseBody) SetTotalCount(v int32) *ListBeebotIntentUserSayResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBeebotIntentUserSayResponseBody) SetUserSays(v []*ListBeebotIntentUserSayResponseBodyUserSays) *ListBeebotIntentUserSayResponseBody {
	s.UserSays = v
	return s
}

func (s *ListBeebotIntentUserSayResponseBody) Validate() error {
	if s.UserSays != nil {
		for _, item := range s.UserSays {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBeebotIntentUserSayResponseBodyUserSays struct {
	// Content
	//
	// example:
	//
	// 你知道xxxx么？
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 2025-04-21 14:16:05.+0800
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Intent ID
	//
	// example:
	//
	// 10717802
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// Updated At
	//
	// example:
	//
	// 2025-04-21 14:16:05.+0800
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// Utterance ID
	//
	// example:
	//
	// 17448458
	UserSayId *string `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s ListBeebotIntentUserSayResponseBodyUserSays) String() string {
	return dara.Prettify(s)
}

func (s ListBeebotIntentUserSayResponseBodyUserSays) GoString() string {
	return s.String()
}

func (s *ListBeebotIntentUserSayResponseBodyUserSays) GetContent() *string {
	return s.Content
}

func (s *ListBeebotIntentUserSayResponseBodyUserSays) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListBeebotIntentUserSayResponseBodyUserSays) GetIntentId() *int64 {
	return s.IntentId
}

func (s *ListBeebotIntentUserSayResponseBodyUserSays) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListBeebotIntentUserSayResponseBodyUserSays) GetUserSayId() *string {
	return s.UserSayId
}

func (s *ListBeebotIntentUserSayResponseBodyUserSays) SetContent(v string) *ListBeebotIntentUserSayResponseBodyUserSays {
	s.Content = &v
	return s
}

func (s *ListBeebotIntentUserSayResponseBodyUserSays) SetCreateTime(v string) *ListBeebotIntentUserSayResponseBodyUserSays {
	s.CreateTime = &v
	return s
}

func (s *ListBeebotIntentUserSayResponseBodyUserSays) SetIntentId(v int64) *ListBeebotIntentUserSayResponseBodyUserSays {
	s.IntentId = &v
	return s
}

func (s *ListBeebotIntentUserSayResponseBodyUserSays) SetModifyTime(v string) *ListBeebotIntentUserSayResponseBodyUserSays {
	s.ModifyTime = &v
	return s
}

func (s *ListBeebotIntentUserSayResponseBodyUserSays) SetUserSayId(v string) *ListBeebotIntentUserSayResponseBodyUserSays {
	s.UserSayId = &v
	return s
}

func (s *ListBeebotIntentUserSayResponseBodyUserSays) Validate() error {
	return dara.Validate(s)
}
