// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDatasetResponseBody
	GetCode() *string
	SetData(v *UpdateDatasetResponseBodyData) *UpdateDatasetResponseBody
	GetData() *UpdateDatasetResponseBodyData
	SetHttpStatusCode(v int32) *UpdateDatasetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateDatasetResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDatasetResponseBody
	GetSuccess() *bool
}

type UpdateDatasetResponseBody struct {
	// example:
	//
	// NoData
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateDatasetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
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

func (s UpdateDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatasetResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDatasetResponseBody) GetData() *UpdateDatasetResponseBodyData {
	return s.Data
}

func (s *UpdateDatasetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateDatasetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDatasetResponseBody) SetCode(v string) *UpdateDatasetResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDatasetResponseBody) SetData(v *UpdateDatasetResponseBodyData) *UpdateDatasetResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDatasetResponseBody) SetHttpStatusCode(v int32) *UpdateDatasetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDatasetResponseBody) SetMessage(v string) *UpdateDatasetResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDatasetResponseBody) SetRequestId(v string) *UpdateDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDatasetResponseBody) SetSuccess(v bool) *UpdateDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDatasetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetResponseBodyData struct {
	// example:
	//
	// 2024-11-12 21:46:24
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// xxx
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// xxx
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// xxx
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// CustomSemanticSearch
	DatasetType        *string                                            `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	NewsArticleResults []*UpdateDatasetResponseBodyDataNewsArticleResults `json:"NewsArticleResults,omitempty" xml:"NewsArticleResults,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	SearchDatasetEnable *int32 `json:"SearchDatasetEnable,omitempty" xml:"SearchDatasetEnable,omitempty"`
}

func (s UpdateDatasetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDatasetResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateDatasetResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *UpdateDatasetResponseBodyData) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *UpdateDatasetResponseBodyData) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *UpdateDatasetResponseBodyData) GetDatasetName() *string {
	return s.DatasetName
}

func (s *UpdateDatasetResponseBodyData) GetDatasetType() *string {
	return s.DatasetType
}

func (s *UpdateDatasetResponseBodyData) GetNewsArticleResults() []*UpdateDatasetResponseBodyDataNewsArticleResults {
	return s.NewsArticleResults
}

func (s *UpdateDatasetResponseBodyData) GetSearchDatasetEnable() *int32 {
	return s.SearchDatasetEnable
}

func (s *UpdateDatasetResponseBodyData) SetCreateTime(v string) *UpdateDatasetResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateDatasetResponseBodyData) SetCreateUser(v string) *UpdateDatasetResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *UpdateDatasetResponseBodyData) SetDatasetDescription(v string) *UpdateDatasetResponseBodyData {
	s.DatasetDescription = &v
	return s
}

func (s *UpdateDatasetResponseBodyData) SetDatasetId(v int64) *UpdateDatasetResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *UpdateDatasetResponseBodyData) SetDatasetName(v string) *UpdateDatasetResponseBodyData {
	s.DatasetName = &v
	return s
}

func (s *UpdateDatasetResponseBodyData) SetDatasetType(v string) *UpdateDatasetResponseBodyData {
	s.DatasetType = &v
	return s
}

func (s *UpdateDatasetResponseBodyData) SetNewsArticleResults(v []*UpdateDatasetResponseBodyDataNewsArticleResults) *UpdateDatasetResponseBodyData {
	s.NewsArticleResults = v
	return s
}

func (s *UpdateDatasetResponseBodyData) SetSearchDatasetEnable(v int32) *UpdateDatasetResponseBodyData {
	s.SearchDatasetEnable = &v
	return s
}

func (s *UpdateDatasetResponseBodyData) Validate() error {
	if s.NewsArticleResults != nil {
		for _, item := range s.NewsArticleResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDatasetResponseBodyDataNewsArticleResults struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Current *int32                                                 `json:"Current,omitempty" xml:"Current,omitempty"`
	Data    []*UpdateDatasetResponseBodyDataNewsArticleResultsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s UpdateDatasetResponseBodyDataNewsArticleResults) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetResponseBodyDataNewsArticleResults) GoString() string {
	return s.String()
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResults) GetCode() *string {
	return s.Code
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResults) GetCurrent() *int32 {
	return s.Current
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResults) GetData() []*UpdateDatasetResponseBodyDataNewsArticleResultsData {
	return s.Data
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResults) GetMessage() *string {
	return s.Message
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResults) GetSize() *int32 {
	return s.Size
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResults) GetTotal() *int32 {
	return s.Total
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResults) SetCode(v string) *UpdateDatasetResponseBodyDataNewsArticleResults {
	s.Code = &v
	return s
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResults) SetCurrent(v int32) *UpdateDatasetResponseBodyDataNewsArticleResults {
	s.Current = &v
	return s
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResults) SetData(v []*UpdateDatasetResponseBodyDataNewsArticleResultsData) *UpdateDatasetResponseBodyDataNewsArticleResults {
	s.Data = v
	return s
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResults) SetMessage(v string) *UpdateDatasetResponseBodyDataNewsArticleResults {
	s.Message = &v
	return s
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResults) SetSize(v int32) *UpdateDatasetResponseBodyDataNewsArticleResults {
	s.Size = &v
	return s
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResults) SetTotal(v int32) *UpdateDatasetResponseBodyDataNewsArticleResults {
	s.Total = &v
	return s
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResults) Validate() error {
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

type UpdateDatasetResponseBodyDataNewsArticleResultsData struct {
	// example:
	//
	// 文章内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2024-11-12 15:12:14
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 央视网
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.docx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateDatasetResponseBodyDataNewsArticleResultsData) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetResponseBodyDataNewsArticleResultsData) GoString() string {
	return s.String()
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResultsData) GetContent() *string {
	return s.Content
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResultsData) GetPubTime() *string {
	return s.PubTime
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResultsData) GetSource() *string {
	return s.Source
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResultsData) GetSummary() *string {
	return s.Summary
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResultsData) GetTitle() *string {
	return s.Title
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResultsData) GetUrl() *string {
	return s.Url
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResultsData) SetContent(v string) *UpdateDatasetResponseBodyDataNewsArticleResultsData {
	s.Content = &v
	return s
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResultsData) SetPubTime(v string) *UpdateDatasetResponseBodyDataNewsArticleResultsData {
	s.PubTime = &v
	return s
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResultsData) SetSource(v string) *UpdateDatasetResponseBodyDataNewsArticleResultsData {
	s.Source = &v
	return s
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResultsData) SetSummary(v string) *UpdateDatasetResponseBodyDataNewsArticleResultsData {
	s.Summary = &v
	return s
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResultsData) SetTitle(v string) *UpdateDatasetResponseBodyDataNewsArticleResultsData {
	s.Title = &v
	return s
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResultsData) SetUrl(v string) *UpdateDatasetResponseBodyDataNewsArticleResultsData {
	s.Url = &v
	return s
}

func (s *UpdateDatasetResponseBodyDataNewsArticleResultsData) Validate() error {
	return dara.Validate(s)
}
