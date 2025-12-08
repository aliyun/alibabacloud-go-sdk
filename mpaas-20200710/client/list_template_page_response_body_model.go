// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatePageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTemplatePageResponseBody
	GetCode() *string
	SetCurrentPage(v int32) *ListTemplatePageResponseBody
	GetCurrentPage() *int32
	SetData(v []*ListTemplatePageResponseBodyData) *ListTemplatePageResponseBody
	GetData() []*ListTemplatePageResponseBodyData
	SetMsg(v string) *ListTemplatePageResponseBody
	GetMsg() *string
	SetPageSize(v int32) *ListTemplatePageResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTemplatePageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTemplatePageResponseBody
	GetSuccess() *bool
	SetTotalSize(v int32) *ListTemplatePageResponseBody
	GetTotalSize() *int32
}

type ListTemplatePageResponseBody struct {
	// example:
	//
	// 100
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*ListTemplatePageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// SUCCESS
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 20
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListTemplatePageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatePageResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatePageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTemplatePageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTemplatePageResponseBody) GetData() []*ListTemplatePageResponseBodyData {
	return s.Data
}

func (s *ListTemplatePageResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ListTemplatePageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTemplatePageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplatePageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTemplatePageResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListTemplatePageResponseBody) SetCode(v string) *ListTemplatePageResponseBody {
	s.Code = &v
	return s
}

func (s *ListTemplatePageResponseBody) SetCurrentPage(v int32) *ListTemplatePageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListTemplatePageResponseBody) SetData(v []*ListTemplatePageResponseBodyData) *ListTemplatePageResponseBody {
	s.Data = v
	return s
}

func (s *ListTemplatePageResponseBody) SetMsg(v string) *ListTemplatePageResponseBody {
	s.Msg = &v
	return s
}

func (s *ListTemplatePageResponseBody) SetPageSize(v int32) *ListTemplatePageResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTemplatePageResponseBody) SetRequestId(v string) *ListTemplatePageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatePageResponseBody) SetSuccess(v bool) *ListTemplatePageResponseBody {
	s.Success = &v
	return s
}

func (s *ListTemplatePageResponseBody) SetTotalSize(v int32) *ListTemplatePageResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListTemplatePageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTemplatePageResponseBodyData struct {
	// example:
	//
	// 1
	Action   *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DescInfo *string `json:"DescInfo,omitempty" xml:"DescInfo,omitempty"`
	// example:
	//
	// 1740479834
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1722564835000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// /
	IconUrls *string `json:"IconUrls,omitempty" xml:"IconUrls,omitempty"`
	// example:
	//
	// 10029984
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// /
	ImageUrls *string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0
	PushStyle *string `json:"PushStyle,omitempty" xml:"PushStyle,omitempty"`
	// example:
	//
	// 0
	ShowStyle *string `json:"ShowStyle,omitempty" xml:"ShowStyle,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// /
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// example:
	//
	// title,content
	Variables *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s ListTemplatePageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatePageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTemplatePageResponseBodyData) GetAction() *string {
	return s.Action
}

func (s *ListTemplatePageResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *ListTemplatePageResponseBodyData) GetDescInfo() *string {
	return s.DescInfo
}

func (s *ListTemplatePageResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListTemplatePageResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListTemplatePageResponseBodyData) GetIconUrls() *string {
	return s.IconUrls
}

func (s *ListTemplatePageResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListTemplatePageResponseBodyData) GetImageUrls() *string {
	return s.ImageUrls
}

func (s *ListTemplatePageResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListTemplatePageResponseBodyData) GetPushStyle() *string {
	return s.PushStyle
}

func (s *ListTemplatePageResponseBodyData) GetShowStyle() *string {
	return s.ShowStyle
}

func (s *ListTemplatePageResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListTemplatePageResponseBodyData) GetUri() *string {
	return s.Uri
}

func (s *ListTemplatePageResponseBodyData) GetVariables() *string {
	return s.Variables
}

func (s *ListTemplatePageResponseBodyData) SetAction(v string) *ListTemplatePageResponseBodyData {
	s.Action = &v
	return s
}

func (s *ListTemplatePageResponseBodyData) SetContent(v string) *ListTemplatePageResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListTemplatePageResponseBodyData) SetDescInfo(v string) *ListTemplatePageResponseBodyData {
	s.DescInfo = &v
	return s
}

func (s *ListTemplatePageResponseBodyData) SetGmtCreate(v string) *ListTemplatePageResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListTemplatePageResponseBodyData) SetGmtModified(v string) *ListTemplatePageResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListTemplatePageResponseBodyData) SetIconUrls(v string) *ListTemplatePageResponseBodyData {
	s.IconUrls = &v
	return s
}

func (s *ListTemplatePageResponseBodyData) SetId(v string) *ListTemplatePageResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListTemplatePageResponseBodyData) SetImageUrls(v string) *ListTemplatePageResponseBodyData {
	s.ImageUrls = &v
	return s
}

func (s *ListTemplatePageResponseBodyData) SetName(v string) *ListTemplatePageResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListTemplatePageResponseBodyData) SetPushStyle(v string) *ListTemplatePageResponseBodyData {
	s.PushStyle = &v
	return s
}

func (s *ListTemplatePageResponseBodyData) SetShowStyle(v string) *ListTemplatePageResponseBodyData {
	s.ShowStyle = &v
	return s
}

func (s *ListTemplatePageResponseBodyData) SetTitle(v string) *ListTemplatePageResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListTemplatePageResponseBodyData) SetUri(v string) *ListTemplatePageResponseBodyData {
	s.Uri = &v
	return s
}

func (s *ListTemplatePageResponseBodyData) SetVariables(v string) *ListTemplatePageResponseBodyData {
	s.Variables = &v
	return s
}

func (s *ListTemplatePageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
