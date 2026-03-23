// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNotifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryNotifyResponseBodyData) *QueryNotifyResponseBody
	GetData() *QueryNotifyResponseBodyData
	SetRequestId(v string) *QueryNotifyResponseBody
	GetRequestId() *string
}

type QueryNotifyResponseBody struct {
	Data      *QueryNotifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryNotifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryNotifyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryNotifyResponseBody) GetData() *QueryNotifyResponseBodyData {
	return s.Data
}

func (s *QueryNotifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryNotifyResponseBody) SetData(v *QueryNotifyResponseBodyData) *QueryNotifyResponseBody {
	s.Data = v
	return s
}

func (s *QueryNotifyResponseBody) SetRequestId(v string) *QueryNotifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryNotifyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryNotifyResponseBodyData struct {
	NotifyItemList   []*QueryNotifyResponseBodyDataNotifyItemList `json:"NotifyItemList,omitempty" xml:"NotifyItemList,omitempty" type:"Repeated"`
	PageNumber       *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalRecordCount *int32                                       `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s QueryNotifyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryNotifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryNotifyResponseBodyData) GetNotifyItemList() []*QueryNotifyResponseBodyDataNotifyItemList {
	return s.NotifyItemList
}

func (s *QueryNotifyResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryNotifyResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryNotifyResponseBodyData) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *QueryNotifyResponseBodyData) SetNotifyItemList(v []*QueryNotifyResponseBodyDataNotifyItemList) *QueryNotifyResponseBodyData {
	s.NotifyItemList = v
	return s
}

func (s *QueryNotifyResponseBodyData) SetPageNumber(v int32) *QueryNotifyResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *QueryNotifyResponseBodyData) SetPageSize(v int32) *QueryNotifyResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryNotifyResponseBodyData) SetTotalRecordCount(v int32) *QueryNotifyResponseBodyData {
	s.TotalRecordCount = &v
	return s
}

func (s *QueryNotifyResponseBodyData) Validate() error {
	if s.NotifyItemList != nil {
		for _, item := range s.NotifyItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryNotifyResponseBodyDataNotifyItemList struct {
	AliUid          *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ConfirmFlag     *bool   `json:"ConfirmFlag,omitempty" xml:"ConfirmFlag,omitempty"`
	Confirmor       *int64  `json:"Confirmor,omitempty" xml:"Confirmor,omitempty"`
	GmtCreated      *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IdempotentCount *string `json:"IdempotentCount,omitempty" xml:"IdempotentCount,omitempty"`
	IdempotentId    *string `json:"IdempotentId,omitempty" xml:"IdempotentId,omitempty"`
	Level           *string `json:"Level,omitempty" xml:"Level,omitempty"`
	NotifyElement   *string `json:"NotifyElement,omitempty" xml:"NotifyElement,omitempty"`
	TemplateName    *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryNotifyResponseBodyDataNotifyItemList) String() string {
	return dara.Prettify(s)
}

func (s QueryNotifyResponseBodyDataNotifyItemList) GoString() string {
	return s.String()
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetConfirmFlag() *bool {
	return s.ConfirmFlag
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetConfirmor() *int64 {
	return s.Confirmor
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetId() *int64 {
	return s.Id
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetIdempotentCount() *string {
	return s.IdempotentCount
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetIdempotentId() *string {
	return s.IdempotentId
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetLevel() *string {
	return s.Level
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetNotifyElement() *string {
	return s.NotifyElement
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) GetType() *string {
	return s.Type
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetAliUid(v int64) *QueryNotifyResponseBodyDataNotifyItemList {
	s.AliUid = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetConfirmFlag(v bool) *QueryNotifyResponseBodyDataNotifyItemList {
	s.ConfirmFlag = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetConfirmor(v int64) *QueryNotifyResponseBodyDataNotifyItemList {
	s.Confirmor = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetGmtCreated(v string) *QueryNotifyResponseBodyDataNotifyItemList {
	s.GmtCreated = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetGmtModified(v string) *QueryNotifyResponseBodyDataNotifyItemList {
	s.GmtModified = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetId(v int64) *QueryNotifyResponseBodyDataNotifyItemList {
	s.Id = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetIdempotentCount(v string) *QueryNotifyResponseBodyDataNotifyItemList {
	s.IdempotentCount = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetIdempotentId(v string) *QueryNotifyResponseBodyDataNotifyItemList {
	s.IdempotentId = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetLevel(v string) *QueryNotifyResponseBodyDataNotifyItemList {
	s.Level = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetNotifyElement(v string) *QueryNotifyResponseBodyDataNotifyItemList {
	s.NotifyElement = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetTemplateName(v string) *QueryNotifyResponseBodyDataNotifyItemList {
	s.TemplateName = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) SetType(v string) *QueryNotifyResponseBodyDataNotifyItemList {
	s.Type = &v
	return s
}

func (s *QueryNotifyResponseBodyDataNotifyItemList) Validate() error {
	return dara.Validate(s)
}
