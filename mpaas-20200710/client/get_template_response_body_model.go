// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTemplateResponseBody
	GetCode() *string
	SetData(v *GetTemplateResponseBodyData) *GetTemplateResponseBody
	GetData() *GetTemplateResponseBodyData
	SetMsg(v string) *GetTemplateResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTemplateResponseBody
	GetSuccess() *bool
}

type GetTemplateResponseBody struct {
	// example:
	//
	// 100
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
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
}

func (s GetTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTemplateResponseBody) GetData() *GetTemplateResponseBodyData {
	return s.Data
}

func (s *GetTemplateResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTemplateResponseBody) SetCode(v string) *GetTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetTemplateResponseBody) SetData(v *GetTemplateResponseBodyData) *GetTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetTemplateResponseBody) SetMsg(v string) *GetTemplateResponseBody {
	s.Msg = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetSuccess(v bool) *GetTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTemplateResponseBodyData struct {
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
	// 1745337419999
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// /
	IconUrls *string `json:"IconUrls,omitempty" xml:"IconUrls,omitempty"`
	// example:
	//
	// 123
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

func (s GetTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyData) GetAction() *string {
	return s.Action
}

func (s *GetTemplateResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetTemplateResponseBodyData) GetDescInfo() *string {
	return s.DescInfo
}

func (s *GetTemplateResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetTemplateResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetTemplateResponseBodyData) GetIconUrls() *string {
	return s.IconUrls
}

func (s *GetTemplateResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetTemplateResponseBodyData) GetImageUrls() *string {
	return s.ImageUrls
}

func (s *GetTemplateResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetTemplateResponseBodyData) GetPushStyle() *string {
	return s.PushStyle
}

func (s *GetTemplateResponseBodyData) GetShowStyle() *string {
	return s.ShowStyle
}

func (s *GetTemplateResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetTemplateResponseBodyData) GetUri() *string {
	return s.Uri
}

func (s *GetTemplateResponseBodyData) GetVariables() *string {
	return s.Variables
}

func (s *GetTemplateResponseBodyData) SetAction(v string) *GetTemplateResponseBodyData {
	s.Action = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetContent(v string) *GetTemplateResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetDescInfo(v string) *GetTemplateResponseBodyData {
	s.DescInfo = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetGmtCreate(v string) *GetTemplateResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetGmtModified(v string) *GetTemplateResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetIconUrls(v string) *GetTemplateResponseBodyData {
	s.IconUrls = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetId(v string) *GetTemplateResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetImageUrls(v string) *GetTemplateResponseBodyData {
	s.ImageUrls = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetName(v string) *GetTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetPushStyle(v string) *GetTemplateResponseBodyData {
	s.PushStyle = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetShowStyle(v string) *GetTemplateResponseBodyData {
	s.ShowStyle = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetTitle(v string) *GetTemplateResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetUri(v string) *GetTemplateResponseBodyData {
	s.Uri = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetVariables(v string) *GetTemplateResponseBodyData {
	s.Variables = &v
	return s
}

func (s *GetTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
