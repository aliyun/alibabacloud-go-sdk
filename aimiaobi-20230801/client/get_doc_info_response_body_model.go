// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDocInfoResponseBody
	GetCode() *string
	SetData(v *GetDocInfoResponseBodyData) *GetDocInfoResponseBody
	GetData() *GetDocInfoResponseBodyData
	SetHttpStatusCode(v int32) *GetDocInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDocInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDocInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDocInfoResponseBody
	GetSuccess() *bool
}

type GetDocInfoResponseBody struct {
	// example:
	//
	// successful
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDocInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDocInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDocInfoResponseBody) GetData() *GetDocInfoResponseBodyData {
	return s.Data
}

func (s *GetDocInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDocInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDocInfoResponseBody) SetCode(v string) *GetDocInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocInfoResponseBody) SetData(v *GetDocInfoResponseBodyData) *GetDocInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetDocInfoResponseBody) SetHttpStatusCode(v int32) *GetDocInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDocInfoResponseBody) SetMessage(v string) *GetDocInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocInfoResponseBody) SetRequestId(v string) *GetDocInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocInfoResponseBody) SetSuccess(v bool) *GetDocInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDocInfoResponseBodyData struct {
	// example:
	//
	// default
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 電視廣播2020年報
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// example:
	//
	// pdf
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// http://xxx/xxx.pdf
	FileUrl  *string                             `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	PageInfo *GetDocInfoResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 导入成功
	StatusMessage *string   `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	VideoContents []*string `json:"VideoContents,omitempty" xml:"VideoContents,omitempty" type:"Repeated"`
}

func (s GetDocInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDocInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocInfoResponseBodyData) GetCategoryId() *string {
	return s.CategoryId
}

func (s *GetDocInfoResponseBodyData) GetDocName() *string {
	return s.DocName
}

func (s *GetDocInfoResponseBodyData) GetDocType() *string {
	return s.DocType
}

func (s *GetDocInfoResponseBodyData) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetDocInfoResponseBodyData) GetPageInfo() *GetDocInfoResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *GetDocInfoResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetDocInfoResponseBodyData) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *GetDocInfoResponseBodyData) GetVideoContents() []*string {
	return s.VideoContents
}

func (s *GetDocInfoResponseBodyData) SetCategoryId(v string) *GetDocInfoResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *GetDocInfoResponseBodyData) SetDocName(v string) *GetDocInfoResponseBodyData {
	s.DocName = &v
	return s
}

func (s *GetDocInfoResponseBodyData) SetDocType(v string) *GetDocInfoResponseBodyData {
	s.DocType = &v
	return s
}

func (s *GetDocInfoResponseBodyData) SetFileUrl(v string) *GetDocInfoResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *GetDocInfoResponseBodyData) SetPageInfo(v *GetDocInfoResponseBodyDataPageInfo) *GetDocInfoResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *GetDocInfoResponseBodyData) SetStatus(v int32) *GetDocInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDocInfoResponseBodyData) SetStatusMessage(v string) *GetDocInfoResponseBodyData {
	s.StatusMessage = &v
	return s
}

func (s *GetDocInfoResponseBodyData) SetVideoContents(v []*string) *GetDocInfoResponseBodyData {
	s.VideoContents = v
	return s
}

func (s *GetDocInfoResponseBodyData) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDocInfoResponseBodyDataPageInfo struct {
	// example:
	//
	// 200
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 100
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetDocInfoResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDocInfoResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *GetDocInfoResponseBodyDataPageInfo) GetHeight() *int32 {
	return s.Height
}

func (s *GetDocInfoResponseBodyDataPageInfo) GetWidth() *int32 {
	return s.Width
}

func (s *GetDocInfoResponseBodyDataPageInfo) SetHeight(v int32) *GetDocInfoResponseBodyDataPageInfo {
	s.Height = &v
	return s
}

func (s *GetDocInfoResponseBodyDataPageInfo) SetWidth(v int32) *GetDocInfoResponseBodyDataPageInfo {
	s.Width = &v
	return s
}

func (s *GetDocInfoResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}
