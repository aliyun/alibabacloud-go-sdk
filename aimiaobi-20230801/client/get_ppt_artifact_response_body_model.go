// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPptArtifactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPptArtifactResponseBody
	GetCode() *string
	SetData(v *GetPptArtifactResponseBodyData) *GetPptArtifactResponseBody
	GetData() *GetPptArtifactResponseBodyData
	SetHttpStatusCode(v int32) *GetPptArtifactResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPptArtifactResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPptArtifactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPptArtifactResponseBody
	GetSuccess() *bool
}

type GetPptArtifactResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetPptArtifactResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPptArtifactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPptArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *GetPptArtifactResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPptArtifactResponseBody) GetData() *GetPptArtifactResponseBodyData {
	return s.Data
}

func (s *GetPptArtifactResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPptArtifactResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPptArtifactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPptArtifactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPptArtifactResponseBody) SetCode(v string) *GetPptArtifactResponseBody {
	s.Code = &v
	return s
}

func (s *GetPptArtifactResponseBody) SetData(v *GetPptArtifactResponseBodyData) *GetPptArtifactResponseBody {
	s.Data = v
	return s
}

func (s *GetPptArtifactResponseBody) SetHttpStatusCode(v int32) *GetPptArtifactResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPptArtifactResponseBody) SetMessage(v string) *GetPptArtifactResponseBody {
	s.Message = &v
	return s
}

func (s *GetPptArtifactResponseBody) SetRequestId(v string) *GetPptArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPptArtifactResponseBody) SetSuccess(v bool) *GetPptArtifactResponseBody {
	s.Success = &v
	return s
}

func (s *GetPptArtifactResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPptArtifactResponseBodyData struct {
	// example:
	//
	// 2024-11-25 11:40:50
	CreateTime *string                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FileAttr   *GetPptArtifactResponseBodyDataFileAttr `json:"FileAttr,omitempty" xml:"FileAttr,omitempty" type:"Struct"`
	// example:
	//
	// oss://default/oss-bucket-name/aimiaobi/2021/07/01/1625126400000/1.docx
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// example:
	//
	// 10
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2024-11-25 11:40:50
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetPptArtifactResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPptArtifactResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPptArtifactResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPptArtifactResponseBodyData) GetFileAttr() *GetPptArtifactResponseBodyDataFileAttr {
	return s.FileAttr
}

func (s *GetPptArtifactResponseBodyData) GetFileKey() *string {
	return s.FileKey
}

func (s *GetPptArtifactResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetPptArtifactResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetPptArtifactResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetPptArtifactResponseBodyData) SetCreateTime(v string) *GetPptArtifactResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetPptArtifactResponseBodyData) SetFileAttr(v *GetPptArtifactResponseBodyDataFileAttr) *GetPptArtifactResponseBodyData {
	s.FileAttr = v
	return s
}

func (s *GetPptArtifactResponseBodyData) SetFileKey(v string) *GetPptArtifactResponseBodyData {
	s.FileKey = &v
	return s
}

func (s *GetPptArtifactResponseBodyData) SetId(v int64) *GetPptArtifactResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetPptArtifactResponseBodyData) SetTitle(v string) *GetPptArtifactResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetPptArtifactResponseBodyData) SetUpdateTime(v string) *GetPptArtifactResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetPptArtifactResponseBodyData) Validate() error {
	if s.FileAttr != nil {
		if err := s.FileAttr.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPptArtifactResponseBodyDataFileAttr struct {
	// example:
	//
	// xxx.pptx
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 600
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// http://www.example.com/xxx.pptx
	TmpUrl *string `json:"TmpUrl,omitempty" xml:"TmpUrl,omitempty"`
	// example:
	//
	// 800
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetPptArtifactResponseBodyDataFileAttr) String() string {
	return dara.Prettify(s)
}

func (s GetPptArtifactResponseBodyDataFileAttr) GoString() string {
	return s.String()
}

func (s *GetPptArtifactResponseBodyDataFileAttr) GetFileName() *string {
	return s.FileName
}

func (s *GetPptArtifactResponseBodyDataFileAttr) GetHeight() *int32 {
	return s.Height
}

func (s *GetPptArtifactResponseBodyDataFileAttr) GetTmpUrl() *string {
	return s.TmpUrl
}

func (s *GetPptArtifactResponseBodyDataFileAttr) GetWidth() *int32 {
	return s.Width
}

func (s *GetPptArtifactResponseBodyDataFileAttr) SetFileName(v string) *GetPptArtifactResponseBodyDataFileAttr {
	s.FileName = &v
	return s
}

func (s *GetPptArtifactResponseBodyDataFileAttr) SetHeight(v int32) *GetPptArtifactResponseBodyDataFileAttr {
	s.Height = &v
	return s
}

func (s *GetPptArtifactResponseBodyDataFileAttr) SetTmpUrl(v string) *GetPptArtifactResponseBodyDataFileAttr {
	s.TmpUrl = &v
	return s
}

func (s *GetPptArtifactResponseBodyDataFileAttr) SetWidth(v int32) *GetPptArtifactResponseBodyDataFileAttr {
	s.Width = &v
	return s
}

func (s *GetPptArtifactResponseBodyDataFileAttr) Validate() error {
	return dara.Validate(s)
}
