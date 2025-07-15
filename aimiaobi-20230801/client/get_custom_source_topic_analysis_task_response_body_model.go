// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomSourceTopicAnalysisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCustomSourceTopicAnalysisTaskResponseBody
	GetCode() *string
	SetData(v *GetCustomSourceTopicAnalysisTaskResponseBodyData) *GetCustomSourceTopicAnalysisTaskResponseBody
	GetData() *GetCustomSourceTopicAnalysisTaskResponseBodyData
	SetHttpStatusCode(v int32) *GetCustomSourceTopicAnalysisTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCustomSourceTopicAnalysisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCustomSourceTopicAnalysisTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomSourceTopicAnalysisTaskResponseBody
	GetSuccess() *bool
}

type GetCustomSourceTopicAnalysisTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetCustomSourceTopicAnalysisTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetCustomSourceTopicAnalysisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomSourceTopicAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBody) GetData() *GetCustomSourceTopicAnalysisTaskResponseBodyData {
	return s.Data
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBody) SetCode(v string) *GetCustomSourceTopicAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBody) SetData(v *GetCustomSourceTopicAnalysisTaskResponseBodyData) *GetCustomSourceTopicAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *GetCustomSourceTopicAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBody) SetMessage(v string) *GetCustomSourceTopicAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBody) SetRequestId(v string) *GetCustomSourceTopicAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBody) SetSuccess(v bool) *GetCustomSourceTopicAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCustomSourceTopicAnalysisTaskResponseBodyData struct {
	// example:
	//
	// 5
	ClusterCount   *int32                                                            `json:"ClusterCount,omitempty" xml:"ClusterCount,omitempty"`
	ClusterResults []*GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResults `json:"ClusterResults,omitempty" xml:"ClusterResults,omitempty" type:"Repeated"`
	// example:
	//
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 8
	MaxClusteredTopicNewsSize *int32 `json:"MaxClusteredTopicNewsSize,omitempty" xml:"MaxClusteredTopicNewsSize,omitempty"`
	// example:
	//
	// 10
	ParsedNewsSize *int32 `json:"ParsedNewsSize,omitempty" xml:"ParsedNewsSize,omitempty"`
	// example:
	//
	// SUCCESSED
	Status *string           `json:"Status,omitempty" xml:"Status,omitempty"`
	Rt     *int64            `json:"rt,omitempty" xml:"rt,omitempty"`
	Usages map[string]*int64 `json:"usages,omitempty" xml:"usages,omitempty"`
}

func (s GetCustomSourceTopicAnalysisTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomSourceTopicAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) GetClusterCount() *int32 {
	return s.ClusterCount
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) GetClusterResults() []*GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResults {
	return s.ClusterResults
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) GetMaxClusteredTopicNewsSize() *int32 {
	return s.MaxClusteredTopicNewsSize
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) GetParsedNewsSize() *int32 {
	return s.ParsedNewsSize
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) GetRt() *int64 {
	return s.Rt
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) GetUsages() map[string]*int64 {
	return s.Usages
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) SetClusterCount(v int32) *GetCustomSourceTopicAnalysisTaskResponseBodyData {
	s.ClusterCount = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) SetClusterResults(v []*GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResults) *GetCustomSourceTopicAnalysisTaskResponseBodyData {
	s.ClusterResults = v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) SetErrorMessage(v string) *GetCustomSourceTopicAnalysisTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) SetMaxClusteredTopicNewsSize(v int32) *GetCustomSourceTopicAnalysisTaskResponseBodyData {
	s.MaxClusteredTopicNewsSize = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) SetParsedNewsSize(v int32) *GetCustomSourceTopicAnalysisTaskResponseBodyData {
	s.ParsedNewsSize = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) SetStatus(v string) *GetCustomSourceTopicAnalysisTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) SetRt(v int64) *GetCustomSourceTopicAnalysisTaskResponseBodyData {
	s.Rt = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) SetUsages(v map[string]*int64) *GetCustomSourceTopicAnalysisTaskResponseBodyData {
	s.Usages = v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResults struct {
	ClusterNews []*GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResultsClusterNews `json:"ClusterNews,omitempty" xml:"ClusterNews,omitempty" type:"Repeated"`
	Topic       *string                                                                      `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResults) String() string {
	return dara.Prettify(s)
}

func (s GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResults) GoString() string {
	return s.String()
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResults) GetClusterNews() []*GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResultsClusterNews {
	return s.ClusterNews
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResults) GetTopic() *string {
	return s.Topic
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResults) SetClusterNews(v []*GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResultsClusterNews) *GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResults {
	s.ClusterNews = v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResults) SetTopic(v string) *GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResults {
	s.Topic = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResults) Validate() error {
	return dara.Validate(s)
}

type GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResultsClusterNews struct {
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://www.example.com/xxx.html
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResultsClusterNews) String() string {
	return dara.Prettify(s)
}

func (s GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResultsClusterNews) GoString() string {
	return s.String()
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResultsClusterNews) GetTitle() *string {
	return s.Title
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResultsClusterNews) GetUrl() *string {
	return s.Url
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResultsClusterNews) SetTitle(v string) *GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResultsClusterNews {
	s.Title = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResultsClusterNews) SetUrl(v string) *GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResultsClusterNews {
	s.Url = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskResponseBodyDataClusterResultsClusterNews) Validate() error {
	return dara.Validate(s)
}
