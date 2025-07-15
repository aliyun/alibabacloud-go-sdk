// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMaterialDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListMaterialDocumentsRequest
	GetAgentKey() *string
	SetContent(v string) *ListMaterialDocumentsRequest
	GetContent() *string
	SetCreateTimeEnd(v string) *ListMaterialDocumentsRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *ListMaterialDocumentsRequest
	GetCreateTimeStart() *string
	SetCurrent(v int32) *ListMaterialDocumentsRequest
	GetCurrent() *int32
	SetDocType(v string) *ListMaterialDocumentsRequest
	GetDocType() *string
	SetDocTypeList(v []*string) *ListMaterialDocumentsRequest
	GetDocTypeList() []*string
	SetGeneratePublicUrl(v bool) *ListMaterialDocumentsRequest
	GetGeneratePublicUrl() *bool
	SetId(v int64) *ListMaterialDocumentsRequest
	GetId() *int64
	SetKeywords(v []*string) *ListMaterialDocumentsRequest
	GetKeywords() []*string
	SetQuery(v string) *ListMaterialDocumentsRequest
	GetQuery() *string
	SetShareAttr(v int32) *ListMaterialDocumentsRequest
	GetShareAttr() *int32
	SetSize(v int32) *ListMaterialDocumentsRequest
	GetSize() *int32
	SetTitle(v string) *ListMaterialDocumentsRequest
	GetTitle() *string
	SetUpdateTimeEnd(v string) *ListMaterialDocumentsRequest
	GetUpdateTimeEnd() *string
	SetUpdateTimeStart(v string) *ListMaterialDocumentsRequest
	GetUpdateTimeStart() *string
}

type ListMaterialDocumentsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 33a2658aaabf4c24b45d50e575125311_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2023-03-18 02:00:00
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// example:
	//
	// 2023-02-19 07:28:11
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// jsonLine
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// excel
	DocTypeList []*string `json:"DocTypeList,omitempty" xml:"DocTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	GeneratePublicUrl *bool `json:"GeneratePublicUrl,omitempty" xml:"GeneratePublicUrl,omitempty"`
	// example:
	//
	// 69
	Id       *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Keywords []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	Query    *string   `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 1
	ShareAttr *int32 `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	// example:
	//
	// 10
	Size  *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2023-03-18 03:00:00
	UpdateTimeEnd *string `json:"UpdateTimeEnd,omitempty" xml:"UpdateTimeEnd,omitempty"`
	// example:
	//
	// 2023-03-18 02:00:00
	UpdateTimeStart *string `json:"UpdateTimeStart,omitempty" xml:"UpdateTimeStart,omitempty"`
}

func (s ListMaterialDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMaterialDocumentsRequest) GoString() string {
	return s.String()
}

func (s *ListMaterialDocumentsRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListMaterialDocumentsRequest) GetContent() *string {
	return s.Content
}

func (s *ListMaterialDocumentsRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *ListMaterialDocumentsRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *ListMaterialDocumentsRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListMaterialDocumentsRequest) GetDocType() *string {
	return s.DocType
}

func (s *ListMaterialDocumentsRequest) GetDocTypeList() []*string {
	return s.DocTypeList
}

func (s *ListMaterialDocumentsRequest) GetGeneratePublicUrl() *bool {
	return s.GeneratePublicUrl
}

func (s *ListMaterialDocumentsRequest) GetId() *int64 {
	return s.Id
}

func (s *ListMaterialDocumentsRequest) GetKeywords() []*string {
	return s.Keywords
}

func (s *ListMaterialDocumentsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListMaterialDocumentsRequest) GetShareAttr() *int32 {
	return s.ShareAttr
}

func (s *ListMaterialDocumentsRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListMaterialDocumentsRequest) GetTitle() *string {
	return s.Title
}

func (s *ListMaterialDocumentsRequest) GetUpdateTimeEnd() *string {
	return s.UpdateTimeEnd
}

func (s *ListMaterialDocumentsRequest) GetUpdateTimeStart() *string {
	return s.UpdateTimeStart
}

func (s *ListMaterialDocumentsRequest) SetAgentKey(v string) *ListMaterialDocumentsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetContent(v string) *ListMaterialDocumentsRequest {
	s.Content = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetCreateTimeEnd(v string) *ListMaterialDocumentsRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetCreateTimeStart(v string) *ListMaterialDocumentsRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetCurrent(v int32) *ListMaterialDocumentsRequest {
	s.Current = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetDocType(v string) *ListMaterialDocumentsRequest {
	s.DocType = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetDocTypeList(v []*string) *ListMaterialDocumentsRequest {
	s.DocTypeList = v
	return s
}

func (s *ListMaterialDocumentsRequest) SetGeneratePublicUrl(v bool) *ListMaterialDocumentsRequest {
	s.GeneratePublicUrl = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetId(v int64) *ListMaterialDocumentsRequest {
	s.Id = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetKeywords(v []*string) *ListMaterialDocumentsRequest {
	s.Keywords = v
	return s
}

func (s *ListMaterialDocumentsRequest) SetQuery(v string) *ListMaterialDocumentsRequest {
	s.Query = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetShareAttr(v int32) *ListMaterialDocumentsRequest {
	s.ShareAttr = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetSize(v int32) *ListMaterialDocumentsRequest {
	s.Size = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetTitle(v string) *ListMaterialDocumentsRequest {
	s.Title = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetUpdateTimeEnd(v string) *ListMaterialDocumentsRequest {
	s.UpdateTimeEnd = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetUpdateTimeStart(v string) *ListMaterialDocumentsRequest {
	s.UpdateTimeStart = &v
	return s
}

func (s *ListMaterialDocumentsRequest) Validate() error {
	return dara.Validate(s)
}
