// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPptArtifactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPptArtifactsResponseBody
	GetCode() *string
	SetCurrent(v int32) *ListPptArtifactsResponseBody
	GetCurrent() *int32
	SetData(v []*ListPptArtifactsResponseBodyData) *ListPptArtifactsResponseBody
	GetData() []*ListPptArtifactsResponseBodyData
	SetHttpStatusCode(v int32) *ListPptArtifactsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListPptArtifactsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListPptArtifactsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListPptArtifactsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPptArtifactsResponseBody
	GetRequestId() *string
	SetSize(v int32) *ListPptArtifactsResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ListPptArtifactsResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListPptArtifactsResponseBody
	GetTotal() *int32
}

type ListPptArtifactsResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Current *int32                              `json:"Current,omitempty" xml:"Current,omitempty"`
	Data    []*ListPptArtifactsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// cEoBWREAXdxaOyjq/cqAbg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListPptArtifactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPptArtifactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPptArtifactsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPptArtifactsResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *ListPptArtifactsResponseBody) GetData() []*ListPptArtifactsResponseBodyData {
	return s.Data
}

func (s *ListPptArtifactsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPptArtifactsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPptArtifactsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPptArtifactsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPptArtifactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPptArtifactsResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ListPptArtifactsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPptArtifactsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListPptArtifactsResponseBody) SetCode(v string) *ListPptArtifactsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPptArtifactsResponseBody) SetCurrent(v int32) *ListPptArtifactsResponseBody {
	s.Current = &v
	return s
}

func (s *ListPptArtifactsResponseBody) SetData(v []*ListPptArtifactsResponseBodyData) *ListPptArtifactsResponseBody {
	s.Data = v
	return s
}

func (s *ListPptArtifactsResponseBody) SetHttpStatusCode(v int32) *ListPptArtifactsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPptArtifactsResponseBody) SetMaxResults(v int32) *ListPptArtifactsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPptArtifactsResponseBody) SetMessage(v string) *ListPptArtifactsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPptArtifactsResponseBody) SetNextToken(v string) *ListPptArtifactsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPptArtifactsResponseBody) SetRequestId(v string) *ListPptArtifactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPptArtifactsResponseBody) SetSize(v int32) *ListPptArtifactsResponseBody {
	s.Size = &v
	return s
}

func (s *ListPptArtifactsResponseBody) SetSuccess(v bool) *ListPptArtifactsResponseBody {
	s.Success = &v
	return s
}

func (s *ListPptArtifactsResponseBody) SetTotal(v int32) *ListPptArtifactsResponseBody {
	s.Total = &v
	return s
}

func (s *ListPptArtifactsResponseBody) Validate() error {
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

type ListPptArtifactsResponseBodyData struct {
	// example:
	//
	// 2024-01-04 11:46:07
	CreateTime *string                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FileAttr   *ListPptArtifactsResponseBodyDataFileAttr `json:"FileAttr,omitempty" xml:"FileAttr,omitempty" type:"Struct"`
	// example:
	//
	// http://www.example.com/xxx.jpg
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// example:
	//
	// 10
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2025-04-14 19:59:53
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListPptArtifactsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPptArtifactsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPptArtifactsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPptArtifactsResponseBodyData) GetFileAttr() *ListPptArtifactsResponseBodyDataFileAttr {
	return s.FileAttr
}

func (s *ListPptArtifactsResponseBodyData) GetFileKey() *string {
	return s.FileKey
}

func (s *ListPptArtifactsResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListPptArtifactsResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListPptArtifactsResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListPptArtifactsResponseBodyData) SetCreateTime(v string) *ListPptArtifactsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListPptArtifactsResponseBodyData) SetFileAttr(v *ListPptArtifactsResponseBodyDataFileAttr) *ListPptArtifactsResponseBodyData {
	s.FileAttr = v
	return s
}

func (s *ListPptArtifactsResponseBodyData) SetFileKey(v string) *ListPptArtifactsResponseBodyData {
	s.FileKey = &v
	return s
}

func (s *ListPptArtifactsResponseBodyData) SetId(v int64) *ListPptArtifactsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListPptArtifactsResponseBodyData) SetTitle(v string) *ListPptArtifactsResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListPptArtifactsResponseBodyData) SetUpdateTime(v string) *ListPptArtifactsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListPptArtifactsResponseBodyData) Validate() error {
	if s.FileAttr != nil {
		if err := s.FileAttr.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPptArtifactsResponseBodyDataFileAttr struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 500
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// http://www.example.com/xxx.jpg
	TmpUrl *string `json:"TmpUrl,omitempty" xml:"TmpUrl,omitempty"`
	// example:
	//
	// 100
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListPptArtifactsResponseBodyDataFileAttr) String() string {
	return dara.Prettify(s)
}

func (s ListPptArtifactsResponseBodyDataFileAttr) GoString() string {
	return s.String()
}

func (s *ListPptArtifactsResponseBodyDataFileAttr) GetFileName() *string {
	return s.FileName
}

func (s *ListPptArtifactsResponseBodyDataFileAttr) GetHeight() *int32 {
	return s.Height
}

func (s *ListPptArtifactsResponseBodyDataFileAttr) GetTmpUrl() *string {
	return s.TmpUrl
}

func (s *ListPptArtifactsResponseBodyDataFileAttr) GetWidth() *int32 {
	return s.Width
}

func (s *ListPptArtifactsResponseBodyDataFileAttr) SetFileName(v string) *ListPptArtifactsResponseBodyDataFileAttr {
	s.FileName = &v
	return s
}

func (s *ListPptArtifactsResponseBodyDataFileAttr) SetHeight(v int32) *ListPptArtifactsResponseBodyDataFileAttr {
	s.Height = &v
	return s
}

func (s *ListPptArtifactsResponseBodyDataFileAttr) SetTmpUrl(v string) *ListPptArtifactsResponseBodyDataFileAttr {
	s.TmpUrl = &v
	return s
}

func (s *ListPptArtifactsResponseBodyDataFileAttr) SetWidth(v int32) *ListPptArtifactsResponseBodyDataFileAttr {
	s.Width = &v
	return s
}

func (s *ListPptArtifactsResponseBodyDataFileAttr) Validate() error {
	return dara.Validate(s)
}
