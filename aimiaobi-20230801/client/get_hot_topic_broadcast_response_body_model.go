// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotTopicBroadcastResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotTopicBroadcastResponseBody
	GetCode() *string
	SetData(v *GetHotTopicBroadcastResponseBodyData) *GetHotTopicBroadcastResponseBody
	GetData() *GetHotTopicBroadcastResponseBodyData
	SetHttpStatusCode(v int32) *GetHotTopicBroadcastResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetHotTopicBroadcastResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotTopicBroadcastResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHotTopicBroadcastResponseBody
	GetSuccess() *bool
}

type GetHotTopicBroadcastResponseBody struct {
	// status code
	//
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Business Data
	Data *GetHotTopicBroadcastResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Fault description
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request UUID
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// is successful: true indicates Succeeded, false indicates failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotTopicBroadcastResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotTopicBroadcastResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotTopicBroadcastResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotTopicBroadcastResponseBody) GetData() *GetHotTopicBroadcastResponseBodyData {
	return s.Data
}

func (s *GetHotTopicBroadcastResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetHotTopicBroadcastResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotTopicBroadcastResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotTopicBroadcastResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHotTopicBroadcastResponseBody) SetCode(v string) *GetHotTopicBroadcastResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBody) SetData(v *GetHotTopicBroadcastResponseBodyData) *GetHotTopicBroadcastResponseBody {
	s.Data = v
	return s
}

func (s *GetHotTopicBroadcastResponseBody) SetHttpStatusCode(v int32) *GetHotTopicBroadcastResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBody) SetMessage(v string) *GetHotTopicBroadcastResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBody) SetRequestId(v string) *GetHotTopicBroadcastResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBody) SetSuccess(v bool) *GetHotTopicBroadcastResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotTopicBroadcastResponseBodyData struct {
	// List of hot spot bulletins
	Data []*GetHotTopicBroadcastResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Total count
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Estimated total number of tokens required for generation
	TotalTokenInfo *GetHotTopicBroadcastResponseBodyDataTotalTokenInfo `json:"TotalTokenInfo,omitempty" xml:"TotalTokenInfo,omitempty" type:"Struct"`
}

func (s GetHotTopicBroadcastResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHotTopicBroadcastResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotTopicBroadcastResponseBodyData) GetData() []*GetHotTopicBroadcastResponseBodyDataData {
	return s.Data
}

func (s *GetHotTopicBroadcastResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetHotTopicBroadcastResponseBodyData) GetTotalTokenInfo() *GetHotTopicBroadcastResponseBodyDataTotalTokenInfo {
	return s.TotalTokenInfo
}

func (s *GetHotTopicBroadcastResponseBodyData) SetData(v []*GetHotTopicBroadcastResponseBodyDataData) *GetHotTopicBroadcastResponseBodyData {
	s.Data = v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyData) SetTotalCount(v int32) *GetHotTopicBroadcastResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyData) SetTotalTokenInfo(v *GetHotTopicBroadcastResponseBodyDataTotalTokenInfo) *GetHotTopicBroadcastResponseBodyData {
	s.TotalTokenInfo = v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyData) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TotalTokenInfo != nil {
		if err := s.TotalTokenInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotTopicBroadcastResponseBodyDataData struct {
	// Hot topic category
	//
	// example:
	//
	// 热点话题分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Custom hotness value
	//
	// example:
	//
	// 34.7905341705522
	CustomHotValue *float64 `json:"CustomHotValue,omitempty" xml:"CustomHotValue,omitempty"`
	// Custom text summarization of the hot spot topic
	//
	// example:
	//
	// 自定义热点话题文本摘要
	CustomTextSummary *string `json:"CustomTextSummary,omitempty" xml:"CustomTextSummary,omitempty"`
	// Name of the hot spot topic
	//
	// example:
	//
	// 热点话题名称
	HotTopic *string `json:"HotTopic,omitempty" xml:"HotTopic,omitempty"`
	// hot spot topic summary Version
	//
	// example:
	//
	// 热点话题摘要版本
	HotTopicVersion *string `json:"HotTopicVersion,omitempty" xml:"HotTopicVersion,omitempty"`
	// Hotness value
	//
	// example:
	//
	// 1.4120480606282884
	HotValue *float64 `json:"HotValue,omitempty" xml:"HotValue,omitempty"`
	// hot spot topic ID
	//
	// example:
	//
	// 热点话题ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// List of hot spot topic images
	Images []*GetHotTopicBroadcastResponseBodyDataDataImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// Input Token
	//
	// example:
	//
	// 29
	InputToken *int32 `json:"InputToken,omitempty" xml:"InputToken,omitempty"`
	// List of Regions associated with the hot spot
	Locations []*string `json:"Locations,omitempty" xml:"Locations,omitempty" type:"Repeated"`
	// Article List
	News []*GetHotTopicBroadcastResponseBodyDataDataNews `json:"News,omitempty" xml:"News,omitempty" type:"Repeated"`
	// Output Token
	//
	// example:
	//
	// 22
	OutputToken *int32 `json:"OutputToken,omitempty" xml:"OutputToken,omitempty"`
	// Published At
	//
	// example:
	//
	// 2025-08-01 12:00:00
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// Structured summary of hot spot topics
	Summary *GetHotTopicBroadcastResponseBodyDataDataSummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Struct"`
	// Text summary of the hot topic
	//
	// example:
	//
	// 热点话题文本摘要
	TextSummary *string `json:"TextSummary,omitempty" xml:"TextSummary,omitempty"`
	// Hot list URL
	//
	// example:
	//
	// http://www.example.com/a.html
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetHotTopicBroadcastResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s GetHotTopicBroadcastResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetCategory() *string {
	return s.Category
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetCustomHotValue() *float64 {
	return s.CustomHotValue
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetCustomTextSummary() *string {
	return s.CustomTextSummary
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetHotTopic() *string {
	return s.HotTopic
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetHotTopicVersion() *string {
	return s.HotTopicVersion
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetHotValue() *float64 {
	return s.HotValue
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetId() *string {
	return s.Id
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetImages() []*GetHotTopicBroadcastResponseBodyDataDataImages {
	return s.Images
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetInputToken() *int32 {
	return s.InputToken
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetLocations() []*string {
	return s.Locations
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetNews() []*GetHotTopicBroadcastResponseBodyDataDataNews {
	return s.News
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetOutputToken() *int32 {
	return s.OutputToken
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetPubTime() *string {
	return s.PubTime
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetSummary() *GetHotTopicBroadcastResponseBodyDataDataSummary {
	return s.Summary
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetTextSummary() *string {
	return s.TextSummary
}

func (s *GetHotTopicBroadcastResponseBodyDataData) GetUrl() *string {
	return s.Url
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetCategory(v string) *GetHotTopicBroadcastResponseBodyDataData {
	s.Category = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetCreateTime(v string) *GetHotTopicBroadcastResponseBodyDataData {
	s.CreateTime = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetCustomHotValue(v float64) *GetHotTopicBroadcastResponseBodyDataData {
	s.CustomHotValue = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetCustomTextSummary(v string) *GetHotTopicBroadcastResponseBodyDataData {
	s.CustomTextSummary = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetHotTopic(v string) *GetHotTopicBroadcastResponseBodyDataData {
	s.HotTopic = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetHotTopicVersion(v string) *GetHotTopicBroadcastResponseBodyDataData {
	s.HotTopicVersion = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetHotValue(v float64) *GetHotTopicBroadcastResponseBodyDataData {
	s.HotValue = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetId(v string) *GetHotTopicBroadcastResponseBodyDataData {
	s.Id = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetImages(v []*GetHotTopicBroadcastResponseBodyDataDataImages) *GetHotTopicBroadcastResponseBodyDataData {
	s.Images = v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetInputToken(v int32) *GetHotTopicBroadcastResponseBodyDataData {
	s.InputToken = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetLocations(v []*string) *GetHotTopicBroadcastResponseBodyDataData {
	s.Locations = v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetNews(v []*GetHotTopicBroadcastResponseBodyDataDataNews) *GetHotTopicBroadcastResponseBodyDataData {
	s.News = v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetOutputToken(v int32) *GetHotTopicBroadcastResponseBodyDataData {
	s.OutputToken = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetPubTime(v string) *GetHotTopicBroadcastResponseBodyDataData {
	s.PubTime = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetSummary(v *GetHotTopicBroadcastResponseBodyDataDataSummary) *GetHotTopicBroadcastResponseBodyDataData {
	s.Summary = v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetTextSummary(v string) *GetHotTopicBroadcastResponseBodyDataData {
	s.TextSummary = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) SetUrl(v string) *GetHotTopicBroadcastResponseBodyDataData {
	s.Url = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataData) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.News != nil {
		for _, item := range s.News {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotTopicBroadcastResponseBodyDataDataImages struct {
	// URL link
	//
	// example:
	//
	// http://www.example.com/a.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetHotTopicBroadcastResponseBodyDataDataImages) String() string {
	return dara.Prettify(s)
}

func (s GetHotTopicBroadcastResponseBodyDataDataImages) GoString() string {
	return s.String()
}

func (s *GetHotTopicBroadcastResponseBodyDataDataImages) GetUrl() *string {
	return s.Url
}

func (s *GetHotTopicBroadcastResponseBodyDataDataImages) SetUrl(v string) *GetHotTopicBroadcastResponseBodyDataDataImages {
	s.Url = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataImages) Validate() error {
	return dara.Validate(s)
}

type GetHotTopicBroadcastResponseBodyDataDataNews struct {
	// Model categorization result
	//
	// example:
	//
	// 科技
	AnalysisCategory *string `json:"AnalysisCategory,omitempty" xml:"AnalysisCategory,omitempty"`
	// aggregated hot spot name
	//
	// example:
	//
	// 聚合后热点名称
	AnalysisTopic *string `json:"AnalysisTopic,omitempty" xml:"AnalysisTopic,omitempty"`
	// Author
	//
	// example:
	//
	// 作者
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// Categorization
	Category []*string `json:"Category,omitempty" xml:"Category,omitempty" type:"Repeated"`
	// News content
	Comments []*GetHotTopicBroadcastResponseBodyDataDataNewsComments `json:"Comments,omitempty" xml:"Comments,omitempty" type:"Repeated"`
	// News content
	//
	// example:
	//
	// 新闻内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Ingestion time
	//
	// example:
	//
	// 2024-06-13 08:45:05
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Source
	//
	// example:
	//
	// 夸克
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Date of entry
	//
	// example:
	//
	// 2024111110
	Dt *string `json:"Dt,omitempty" xml:"Dt,omitempty"`
	// original hot spot name
	//
	// example:
	//
	// 原始热点名称
	HotTopic *string `json:"HotTopic,omitempty" xml:"HotTopic,omitempty"`
	// Image list
	ImgList []*string `json:"ImgList,omitempty" xml:"ImgList,omitempty" type:"Repeated"`
	// logo
	//
	// example:
	//
	// https://www.example.com/a.png
	Logo *string `json:"Logo,omitempty" xml:"Logo,omitempty"`
	// Published At
	//
	// example:
	//
	// 2024-10-10 12:12:00
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// Summary
	//
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// title
	//
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// news URL
	//
	// example:
	//
	// http://www.example.com/a.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// Primary key ID
	//
	// example:
	//
	// 主键ID
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// website
	//
	// example:
	//
	// 网站
	Website *string `json:"Website,omitempty" xml:"Website,omitempty"`
}

func (s GetHotTopicBroadcastResponseBodyDataDataNews) String() string {
	return dara.Prettify(s)
}

func (s GetHotTopicBroadcastResponseBodyDataDataNews) GoString() string {
	return s.String()
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetAnalysisCategory() *string {
	return s.AnalysisCategory
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetAnalysisTopic() *string {
	return s.AnalysisTopic
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetAuthor() *string {
	return s.Author
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetCategory() []*string {
	return s.Category
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetComments() []*GetHotTopicBroadcastResponseBodyDataDataNewsComments {
	return s.Comments
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetContent() *string {
	return s.Content
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetDomain() *string {
	return s.Domain
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetDt() *string {
	return s.Dt
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetHotTopic() *string {
	return s.HotTopic
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetImgList() []*string {
	return s.ImgList
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetLogo() *string {
	return s.Logo
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetPubTime() *string {
	return s.PubTime
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetSummary() *string {
	return s.Summary
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetTitle() *string {
	return s.Title
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetUrl() *string {
	return s.Url
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetUuid() *string {
	return s.Uuid
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) GetWebsite() *string {
	return s.Website
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetAnalysisCategory(v string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.AnalysisCategory = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetAnalysisTopic(v string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.AnalysisTopic = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetAuthor(v string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.Author = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetCategory(v []*string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.Category = v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetComments(v []*GetHotTopicBroadcastResponseBodyDataDataNewsComments) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.Comments = v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetContent(v string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.Content = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetCreateTime(v string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.CreateTime = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetDomain(v string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.Domain = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetDt(v string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.Dt = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetHotTopic(v string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.HotTopic = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetImgList(v []*string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.ImgList = v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetLogo(v string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.Logo = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetPubTime(v string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.PubTime = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetSummary(v string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.Summary = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetTitle(v string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.Title = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetUrl(v string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.Url = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetUuid(v string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.Uuid = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) SetWebsite(v string) *GetHotTopicBroadcastResponseBodyDataDataNews {
	s.Website = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNews) Validate() error {
	if s.Comments != nil {
		for _, item := range s.Comments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHotTopicBroadcastResponseBodyDataDataNewsComments struct {
	// Content
	//
	// example:
	//
	// 内容
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// Username
	//
	// example:
	//
	// 用户名
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetHotTopicBroadcastResponseBodyDataDataNewsComments) String() string {
	return dara.Prettify(s)
}

func (s GetHotTopicBroadcastResponseBodyDataDataNewsComments) GoString() string {
	return s.String()
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNewsComments) GetText() *string {
	return s.Text
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNewsComments) GetUsername() *string {
	return s.Username
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNewsComments) SetText(v string) *GetHotTopicBroadcastResponseBodyDataDataNewsComments {
	s.Text = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNewsComments) SetUsername(v string) *GetHotTopicBroadcastResponseBodyDataDataNewsComments {
	s.Username = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataNewsComments) Validate() error {
	return dara.Validate(s)
}

type GetHotTopicBroadcastResponseBodyDataDataSummary struct {
	// Number of input tokens used to generate this summary
	//
	// example:
	//
	// 17
	InputToken *int32 `json:"InputToken,omitempty" xml:"InputToken,omitempty"`
	// Number of output tokens used to generate this summary
	//
	// example:
	//
	// 41
	OutputToken *int32 `json:"OutputToken,omitempty" xml:"OutputToken,omitempty"`
	// List of structured summaries
	Summaries []*GetHotTopicBroadcastResponseBodyDataDataSummarySummaries `json:"Summaries,omitempty" xml:"Summaries,omitempty" type:"Repeated"`
}

func (s GetHotTopicBroadcastResponseBodyDataDataSummary) String() string {
	return dara.Prettify(s)
}

func (s GetHotTopicBroadcastResponseBodyDataDataSummary) GoString() string {
	return s.String()
}

func (s *GetHotTopicBroadcastResponseBodyDataDataSummary) GetInputToken() *int32 {
	return s.InputToken
}

func (s *GetHotTopicBroadcastResponseBodyDataDataSummary) GetOutputToken() *int32 {
	return s.OutputToken
}

func (s *GetHotTopicBroadcastResponseBodyDataDataSummary) GetSummaries() []*GetHotTopicBroadcastResponseBodyDataDataSummarySummaries {
	return s.Summaries
}

func (s *GetHotTopicBroadcastResponseBodyDataDataSummary) SetInputToken(v int32) *GetHotTopicBroadcastResponseBodyDataDataSummary {
	s.InputToken = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataSummary) SetOutputToken(v int32) *GetHotTopicBroadcastResponseBodyDataDataSummary {
	s.OutputToken = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataSummary) SetSummaries(v []*GetHotTopicBroadcastResponseBodyDataDataSummarySummaries) *GetHotTopicBroadcastResponseBodyDataDataSummary {
	s.Summaries = v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataSummary) Validate() error {
	if s.Summaries != nil {
		for _, item := range s.Summaries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHotTopicBroadcastResponseBodyDataDataSummarySummaries struct {
	// Summary
	//
	// example:
	//
	// 摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// title
	//
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetHotTopicBroadcastResponseBodyDataDataSummarySummaries) String() string {
	return dara.Prettify(s)
}

func (s GetHotTopicBroadcastResponseBodyDataDataSummarySummaries) GoString() string {
	return s.String()
}

func (s *GetHotTopicBroadcastResponseBodyDataDataSummarySummaries) GetSummary() *string {
	return s.Summary
}

func (s *GetHotTopicBroadcastResponseBodyDataDataSummarySummaries) GetTitle() *string {
	return s.Title
}

func (s *GetHotTopicBroadcastResponseBodyDataDataSummarySummaries) SetSummary(v string) *GetHotTopicBroadcastResponseBodyDataDataSummarySummaries {
	s.Summary = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataSummarySummaries) SetTitle(v string) *GetHotTopicBroadcastResponseBodyDataDataSummarySummaries {
	s.Title = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataDataSummarySummaries) Validate() error {
	return dara.Validate(s)
}

type GetHotTopicBroadcastResponseBodyDataTotalTokenInfo struct {
	// Total number of hot spots
	//
	// example:
	//
	// 100
	HotTopicCount *int32 `json:"HotTopicCount,omitempty" xml:"HotTopicCount,omitempty"`
	// Estimated number of input tokens
	//
	// example:
	//
	// 100
	InputTokens *int32 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// Estimated number of output tokens
	//
	// example:
	//
	// 100
	OutputTokens *int32 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// Estimated total word count
	//
	// example:
	//
	// 100
	WordCount *int32 `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s GetHotTopicBroadcastResponseBodyDataTotalTokenInfo) String() string {
	return dara.Prettify(s)
}

func (s GetHotTopicBroadcastResponseBodyDataTotalTokenInfo) GoString() string {
	return s.String()
}

func (s *GetHotTopicBroadcastResponseBodyDataTotalTokenInfo) GetHotTopicCount() *int32 {
	return s.HotTopicCount
}

func (s *GetHotTopicBroadcastResponseBodyDataTotalTokenInfo) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *GetHotTopicBroadcastResponseBodyDataTotalTokenInfo) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *GetHotTopicBroadcastResponseBodyDataTotalTokenInfo) GetWordCount() *int32 {
	return s.WordCount
}

func (s *GetHotTopicBroadcastResponseBodyDataTotalTokenInfo) SetHotTopicCount(v int32) *GetHotTopicBroadcastResponseBodyDataTotalTokenInfo {
	s.HotTopicCount = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataTotalTokenInfo) SetInputTokens(v int32) *GetHotTopicBroadcastResponseBodyDataTotalTokenInfo {
	s.InputTokens = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataTotalTokenInfo) SetOutputTokens(v int32) *GetHotTopicBroadcastResponseBodyDataTotalTokenInfo {
	s.OutputTokens = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataTotalTokenInfo) SetWordCount(v int32) *GetHotTopicBroadcastResponseBodyDataTotalTokenInfo {
	s.WordCount = &v
	return s
}

func (s *GetHotTopicBroadcastResponseBodyDataTotalTokenInfo) Validate() error {
	return dara.Validate(s)
}
