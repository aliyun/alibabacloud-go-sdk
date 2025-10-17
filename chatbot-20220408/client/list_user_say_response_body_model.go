// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserSayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListUserSayResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserSayResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListUserSayResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListUserSayResponseBody
	GetTotalCount() *int32
	SetUserSays(v []*ListUserSayResponseBodyUserSays) *ListUserSayResponseBody
	GetUserSays() []*ListUserSayResponseBodyUserSays
}

type ListUserSayResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// fs1fg4512v43572v23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserSays   []*ListUserSayResponseBodyUserSays `json:"UserSays,omitempty" xml:"UserSays,omitempty" type:"Repeated"`
}

func (s ListUserSayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserSayResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserSayResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserSayResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserSayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserSayResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserSayResponseBody) GetUserSays() []*ListUserSayResponseBodyUserSays {
	return s.UserSays
}

func (s *ListUserSayResponseBody) SetPageNumber(v int32) *ListUserSayResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUserSayResponseBody) SetPageSize(v int32) *ListUserSayResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserSayResponseBody) SetRequestId(v string) *ListUserSayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserSayResponseBody) SetTotalCount(v int32) *ListUserSayResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserSayResponseBody) SetUserSays(v []*ListUserSayResponseBodyUserSays) *ListUserSayResponseBody {
	s.UserSays = v
	return s
}

func (s *ListUserSayResponseBody) Validate() error {
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

type ListUserSayResponseBodyUserSays struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 235564564
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	ModifyTime *string                                     `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	SlotInfos  []*ListUserSayResponseBodyUserSaysSlotInfos `json:"SlotInfos,omitempty" xml:"SlotInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 3453452138
	UserSayId *int64 `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s ListUserSayResponseBodyUserSays) String() string {
	return dara.Prettify(s)
}

func (s ListUserSayResponseBodyUserSays) GoString() string {
	return s.String()
}

func (s *ListUserSayResponseBodyUserSays) GetContent() *string {
	return s.Content
}

func (s *ListUserSayResponseBodyUserSays) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUserSayResponseBodyUserSays) GetIntentId() *int64 {
	return s.IntentId
}

func (s *ListUserSayResponseBodyUserSays) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListUserSayResponseBodyUserSays) GetSlotInfos() []*ListUserSayResponseBodyUserSaysSlotInfos {
	return s.SlotInfos
}

func (s *ListUserSayResponseBodyUserSays) GetUserSayId() *int64 {
	return s.UserSayId
}

func (s *ListUserSayResponseBodyUserSays) SetContent(v string) *ListUserSayResponseBodyUserSays {
	s.Content = &v
	return s
}

func (s *ListUserSayResponseBodyUserSays) SetCreateTime(v string) *ListUserSayResponseBodyUserSays {
	s.CreateTime = &v
	return s
}

func (s *ListUserSayResponseBodyUserSays) SetIntentId(v int64) *ListUserSayResponseBodyUserSays {
	s.IntentId = &v
	return s
}

func (s *ListUserSayResponseBodyUserSays) SetModifyTime(v string) *ListUserSayResponseBodyUserSays {
	s.ModifyTime = &v
	return s
}

func (s *ListUserSayResponseBodyUserSays) SetSlotInfos(v []*ListUserSayResponseBodyUserSaysSlotInfos) *ListUserSayResponseBodyUserSays {
	s.SlotInfos = v
	return s
}

func (s *ListUserSayResponseBodyUserSays) SetUserSayId(v int64) *ListUserSayResponseBodyUserSays {
	s.UserSayId = &v
	return s
}

func (s *ListUserSayResponseBodyUserSays) Validate() error {
	if s.SlotInfos != nil {
		for _, item := range s.SlotInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserSayResponseBodyUserSaysSlotInfos struct {
	// example:
	//
	// 4
	EndIndex *int32 `json:"EndIndex,omitempty" xml:"EndIndex,omitempty"`
	// example:
	//
	// 3456sdfg3tu
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	// example:
	//
	// 2
	StartIndex *int32 `json:"StartIndex,omitempty" xml:"StartIndex,omitempty"`
}

func (s ListUserSayResponseBodyUserSaysSlotInfos) String() string {
	return dara.Prettify(s)
}

func (s ListUserSayResponseBodyUserSaysSlotInfos) GoString() string {
	return s.String()
}

func (s *ListUserSayResponseBodyUserSaysSlotInfos) GetEndIndex() *int32 {
	return s.EndIndex
}

func (s *ListUserSayResponseBodyUserSaysSlotInfos) GetSlotId() *string {
	return s.SlotId
}

func (s *ListUserSayResponseBodyUserSaysSlotInfos) GetStartIndex() *int32 {
	return s.StartIndex
}

func (s *ListUserSayResponseBodyUserSaysSlotInfos) SetEndIndex(v int32) *ListUserSayResponseBodyUserSaysSlotInfos {
	s.EndIndex = &v
	return s
}

func (s *ListUserSayResponseBodyUserSaysSlotInfos) SetSlotId(v string) *ListUserSayResponseBodyUserSaysSlotInfos {
	s.SlotId = &v
	return s
}

func (s *ListUserSayResponseBodyUserSaysSlotInfos) SetStartIndex(v int32) *ListUserSayResponseBodyUserSaysSlotInfos {
	s.StartIndex = &v
	return s
}

func (s *ListUserSayResponseBodyUserSaysSlotInfos) Validate() error {
	return dara.Validate(s)
}
