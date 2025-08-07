// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadMessageContentResponseBody
	GetCode() *string
	SetData(v *ReadMessageContentResponseBodyData) *ReadMessageContentResponseBody
	GetData() *ReadMessageContentResponseBodyData
	SetMessage(v string) *ReadMessageContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadMessageContentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadMessageContentResponseBody
	GetSuccess() *bool
}

type ReadMessageContentResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ReadMessageContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadMessageContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageContentResponseBody) GoString() string {
	return s.String()
}

func (s *ReadMessageContentResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadMessageContentResponseBody) GetData() *ReadMessageContentResponseBodyData {
	return s.Data
}

func (s *ReadMessageContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadMessageContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadMessageContentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadMessageContentResponseBody) SetCode(v string) *ReadMessageContentResponseBody {
	s.Code = &v
	return s
}

func (s *ReadMessageContentResponseBody) SetData(v *ReadMessageContentResponseBodyData) *ReadMessageContentResponseBody {
	s.Data = v
	return s
}

func (s *ReadMessageContentResponseBody) SetMessage(v string) *ReadMessageContentResponseBody {
	s.Message = &v
	return s
}

func (s *ReadMessageContentResponseBody) SetRequestId(v string) *ReadMessageContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadMessageContentResponseBody) SetSuccess(v bool) *ReadMessageContentResponseBody {
	s.Success = &v
	return s
}

func (s *ReadMessageContentResponseBody) Validate() error {
	return dara.Validate(s)
}

type ReadMessageContentResponseBodyData struct {
	Datas *ReadMessageContentResponseBodyDataDatas `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Struct"`
}

func (s ReadMessageContentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadMessageContentResponseBodyData) GetDatas() *ReadMessageContentResponseBodyDataDatas {
	return s.Datas
}

func (s *ReadMessageContentResponseBodyData) SetDatas(v *ReadMessageContentResponseBodyDataDatas) *ReadMessageContentResponseBodyData {
	s.Datas = v
	return s
}

func (s *ReadMessageContentResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ReadMessageContentResponseBodyDataDatas struct {
	Item     []*ReadMessageContentResponseBodyDataDatasItem     `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	LastItem []*ReadMessageContentResponseBodyDataDatasLastItem `json:"LastItem,omitempty" xml:"LastItem,omitempty" type:"Repeated"`
	NextItem []*ReadMessageContentResponseBodyDataDatasNextItem `json:"NextItem,omitempty" xml:"NextItem,omitempty" type:"Repeated"`
}

func (s ReadMessageContentResponseBodyDataDatas) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageContentResponseBodyDataDatas) GoString() string {
	return s.String()
}

func (s *ReadMessageContentResponseBodyDataDatas) GetItem() []*ReadMessageContentResponseBodyDataDatasItem {
	return s.Item
}

func (s *ReadMessageContentResponseBodyDataDatas) GetLastItem() []*ReadMessageContentResponseBodyDataDatasLastItem {
	return s.LastItem
}

func (s *ReadMessageContentResponseBodyDataDatas) GetNextItem() []*ReadMessageContentResponseBodyDataDatasNextItem {
	return s.NextItem
}

func (s *ReadMessageContentResponseBodyDataDatas) SetItem(v []*ReadMessageContentResponseBodyDataDatasItem) *ReadMessageContentResponseBodyDataDatas {
	s.Item = v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatas) SetLastItem(v []*ReadMessageContentResponseBodyDataDatasLastItem) *ReadMessageContentResponseBodyDataDatas {
	s.LastItem = v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatas) SetNextItem(v []*ReadMessageContentResponseBodyDataDatasNextItem) *ReadMessageContentResponseBodyDataDatas {
	s.NextItem = v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatas) Validate() error {
	return dara.Validate(s)
}

type ReadMessageContentResponseBodyDataDatasItem struct {
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	ClassId      *int64  `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Deleted      *int32  `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	GmtCreated   *int64  `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtUpdate    *int64  `json:"GmtUpdate,omitempty" xml:"GmtUpdate,omitempty"`
	MassId       *int64  `json:"MassId,omitempty" xml:"MassId,omitempty"`
	Memo         *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	MsgId        *int64  `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Status       *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ReadMessageContentResponseBodyDataDatasItem) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageContentResponseBodyDataDatasItem) GoString() string {
	return s.String()
}

func (s *ReadMessageContentResponseBodyDataDatasItem) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ReadMessageContentResponseBodyDataDatasItem) GetClassId() *int64 {
	return s.ClassId
}

func (s *ReadMessageContentResponseBodyDataDatasItem) GetContent() *string {
	return s.Content
}

func (s *ReadMessageContentResponseBodyDataDatasItem) GetDeleted() *int32 {
	return s.Deleted
}

func (s *ReadMessageContentResponseBodyDataDatasItem) GetGmtCreated() *int64 {
	return s.GmtCreated
}

func (s *ReadMessageContentResponseBodyDataDatasItem) GetGmtUpdate() *int64 {
	return s.GmtUpdate
}

func (s *ReadMessageContentResponseBodyDataDatasItem) GetMassId() *int64 {
	return s.MassId
}

func (s *ReadMessageContentResponseBodyDataDatasItem) GetMemo() *string {
	return s.Memo
}

func (s *ReadMessageContentResponseBodyDataDatasItem) GetMsgId() *int64 {
	return s.MsgId
}

func (s *ReadMessageContentResponseBodyDataDatasItem) GetStatus() *int64 {
	return s.Status
}

func (s *ReadMessageContentResponseBodyDataDatasItem) GetTitle() *string {
	return s.Title
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetCategoryName(v string) *ReadMessageContentResponseBodyDataDatasItem {
	s.CategoryName = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetClassId(v int64) *ReadMessageContentResponseBodyDataDatasItem {
	s.ClassId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetContent(v string) *ReadMessageContentResponseBodyDataDatasItem {
	s.Content = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetDeleted(v int32) *ReadMessageContentResponseBodyDataDatasItem {
	s.Deleted = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetGmtCreated(v int64) *ReadMessageContentResponseBodyDataDatasItem {
	s.GmtCreated = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetGmtUpdate(v int64) *ReadMessageContentResponseBodyDataDatasItem {
	s.GmtUpdate = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetMassId(v int64) *ReadMessageContentResponseBodyDataDatasItem {
	s.MassId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetMemo(v string) *ReadMessageContentResponseBodyDataDatasItem {
	s.Memo = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetMsgId(v int64) *ReadMessageContentResponseBodyDataDatasItem {
	s.MsgId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetStatus(v int64) *ReadMessageContentResponseBodyDataDatasItem {
	s.Status = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) SetTitle(v string) *ReadMessageContentResponseBodyDataDatasItem {
	s.Title = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasItem) Validate() error {
	return dara.Validate(s)
}

type ReadMessageContentResponseBodyDataDatasLastItem struct {
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	ClassId      *int64  `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Deleted      *int32  `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	GmtCreated   *int64  `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtUpdate    *int64  `json:"GmtUpdate,omitempty" xml:"GmtUpdate,omitempty"`
	MassId       *int64  `json:"MassId,omitempty" xml:"MassId,omitempty"`
	Memo         *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	MsgId        *int64  `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Status       *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ReadMessageContentResponseBodyDataDatasLastItem) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageContentResponseBodyDataDatasLastItem) GoString() string {
	return s.String()
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) GetClassId() *int64 {
	return s.ClassId
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) GetContent() *string {
	return s.Content
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) GetDeleted() *int32 {
	return s.Deleted
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) GetGmtCreated() *int64 {
	return s.GmtCreated
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) GetGmtUpdate() *int64 {
	return s.GmtUpdate
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) GetMassId() *int64 {
	return s.MassId
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) GetMemo() *string {
	return s.Memo
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) GetMsgId() *int64 {
	return s.MsgId
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) GetStatus() *int64 {
	return s.Status
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) GetTitle() *string {
	return s.Title
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetCategoryName(v string) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.CategoryName = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetClassId(v int64) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.ClassId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetContent(v string) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.Content = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetDeleted(v int32) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.Deleted = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetGmtCreated(v int64) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.GmtCreated = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetGmtUpdate(v int64) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.GmtUpdate = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetMassId(v int64) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.MassId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetMemo(v string) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.Memo = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetMsgId(v int64) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.MsgId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetStatus(v int64) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.Status = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) SetTitle(v string) *ReadMessageContentResponseBodyDataDatasLastItem {
	s.Title = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasLastItem) Validate() error {
	return dara.Validate(s)
}

type ReadMessageContentResponseBodyDataDatasNextItem struct {
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	ClassId      *int64  `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Deleted      *int32  `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	GmtCreated   *int64  `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtUpdate    *int64  `json:"GmtUpdate,omitempty" xml:"GmtUpdate,omitempty"`
	MassId       *int64  `json:"MassId,omitempty" xml:"MassId,omitempty"`
	Memo         *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	MsgId        *int64  `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Status       *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ReadMessageContentResponseBodyDataDatasNextItem) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageContentResponseBodyDataDatasNextItem) GoString() string {
	return s.String()
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) GetClassId() *int64 {
	return s.ClassId
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) GetContent() *string {
	return s.Content
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) GetDeleted() *int32 {
	return s.Deleted
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) GetGmtCreated() *int64 {
	return s.GmtCreated
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) GetGmtUpdate() *int64 {
	return s.GmtUpdate
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) GetMassId() *int64 {
	return s.MassId
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) GetMemo() *string {
	return s.Memo
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) GetMsgId() *int64 {
	return s.MsgId
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) GetStatus() *int64 {
	return s.Status
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) GetTitle() *string {
	return s.Title
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetCategoryName(v string) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.CategoryName = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetClassId(v int64) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.ClassId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetContent(v string) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.Content = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetDeleted(v int32) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.Deleted = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetGmtCreated(v int64) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.GmtCreated = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetGmtUpdate(v int64) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.GmtUpdate = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetMassId(v int64) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.MassId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetMemo(v string) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.Memo = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetMsgId(v int64) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.MsgId = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetStatus(v int64) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.Status = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) SetTitle(v string) *ReadMessageContentResponseBodyDataDatasNextItem {
	s.Title = &v
	return s
}

func (s *ReadMessageContentResponseBodyDataDatasNextItem) Validate() error {
	return dara.Validate(s)
}
