// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddAnswerSampleRequest struct {
	// example:
	//
	// alxxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// data/xxx.xlsx
	SampleObject *string `json:"SampleObject,omitempty" xml:"SampleObject,omitempty"`
	Samples      *string `json:"Samples,omitempty" xml:"Samples,omitempty"`
}

func (s AddAnswerSampleRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAnswerSampleRequest) GoString() string {
	return s.String()
}

func (s *AddAnswerSampleRequest) SetLibId(v string) *AddAnswerSampleRequest {
	s.LibId = &v
	return s
}

func (s *AddAnswerSampleRequest) SetRegionId(v string) *AddAnswerSampleRequest {
	s.RegionId = &v
	return s
}

func (s *AddAnswerSampleRequest) SetSampleObject(v string) *AddAnswerSampleRequest {
	s.SampleObject = &v
	return s
}

func (s *AddAnswerSampleRequest) SetSamples(v string) *AddAnswerSampleRequest {
	s.Samples = &v
	return s
}

type AddAnswerSampleResponseBody struct {
	// example:
	//
	// alxxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AddAnswerSampleResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// alAxbbxxxx-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddAnswerSampleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAnswerSampleResponseBody) GoString() string {
	return s.String()
}

func (s *AddAnswerSampleResponseBody) SetLibId(v string) *AddAnswerSampleResponseBody {
	s.LibId = &v
	return s
}

func (s *AddAnswerSampleResponseBody) SetRequestId(v string) *AddAnswerSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAnswerSampleResponseBody) SetResult(v *AddAnswerSampleResponseBodyResult) *AddAnswerSampleResponseBody {
	s.Result = v
	return s
}

func (s *AddAnswerSampleResponseBody) SetTaskId(v string) *AddAnswerSampleResponseBody {
	s.TaskId = &v
	return s
}

type AddAnswerSampleResponseBodyResult struct {
	// example:
	//
	// xxx
	I18nKey              *string   `json:"I18nKey,omitempty" xml:"I18nKey,omitempty"`
	IllegalLengthSamples []*string `json:"IllegalLengthSamples,omitempty" xml:"IllegalLengthSamples,omitempty" type:"Repeated"`
	// example:
	//
	// 118
	InvalidCount *int32 `json:"InvalidCount,omitempty" xml:"InvalidCount,omitempty"`
	// example:
	//
	// alxxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 98
	RepeatCount   *int32    `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	RepeatSamples []*string `json:"RepeatSamples,omitempty" xml:"RepeatSamples,omitempty" type:"Repeated"`
	// example:
	//
	// 318
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// alAxbbxxxx-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 534
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s AddAnswerSampleResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddAnswerSampleResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddAnswerSampleResponseBodyResult) SetI18nKey(v string) *AddAnswerSampleResponseBodyResult {
	s.I18nKey = &v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetIllegalLengthSamples(v []*string) *AddAnswerSampleResponseBodyResult {
	s.IllegalLengthSamples = v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetInvalidCount(v int32) *AddAnswerSampleResponseBodyResult {
	s.InvalidCount = &v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetLibId(v string) *AddAnswerSampleResponseBodyResult {
	s.LibId = &v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetProgress(v int32) *AddAnswerSampleResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetRepeatCount(v int32) *AddAnswerSampleResponseBodyResult {
	s.RepeatCount = &v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetRepeatSamples(v []*string) *AddAnswerSampleResponseBodyResult {
	s.RepeatSamples = v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetSuccessCount(v int32) *AddAnswerSampleResponseBodyResult {
	s.SuccessCount = &v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetTaskId(v string) *AddAnswerSampleResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *AddAnswerSampleResponseBodyResult) SetTotalCount(v int32) *AddAnswerSampleResponseBodyResult {
	s.TotalCount = &v
	return s
}

type AddAnswerSampleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAnswerSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAnswerSampleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAnswerSampleResponse) GoString() string {
	return s.String()
}

func (s *AddAnswerSampleResponse) SetHeaders(v map[string]*string) *AddAnswerSampleResponse {
	s.Headers = v
	return s
}

func (s *AddAnswerSampleResponse) SetStatusCode(v int32) *AddAnswerSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAnswerSampleResponse) SetBody(v *AddAnswerSampleResponseBody) *AddAnswerSampleResponse {
	s.Body = v
	return s
}

type AddImageLibRequest struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddImageLibRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageLibRequest) GoString() string {
	return s.String()
}

func (s *AddImageLibRequest) SetComment(v string) *AddImageLibRequest {
	s.Comment = &v
	return s
}

func (s *AddImageLibRequest) SetLibName(v string) *AddImageLibRequest {
	s.LibName = &v
	return s
}

func (s *AddImageLibRequest) SetRegionId(v string) *AddImageLibRequest {
	s.RegionId = &v
	return s
}

type AddImageLibResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddImageLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddImageLibResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageLibResponseBody) SetCode(v int32) *AddImageLibResponseBody {
	s.Code = &v
	return s
}

func (s *AddImageLibResponseBody) SetData(v bool) *AddImageLibResponseBody {
	s.Data = &v
	return s
}

func (s *AddImageLibResponseBody) SetHttpStatusCode(v int32) *AddImageLibResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddImageLibResponseBody) SetMsg(v string) *AddImageLibResponseBody {
	s.Msg = &v
	return s
}

func (s *AddImageLibResponseBody) SetRequestId(v string) *AddImageLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImageLibResponseBody) SetSuccess(v bool) *AddImageLibResponseBody {
	s.Success = &v
	return s
}

type AddImageLibResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddImageLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddImageLibResponse) String() string {
	return tea.Prettify(s)
}

func (s AddImageLibResponse) GoString() string {
	return s.String()
}

func (s *AddImageLibResponse) SetHeaders(v map[string]*string) *AddImageLibResponse {
	s.Headers = v
	return s
}

func (s *AddImageLibResponse) SetStatusCode(v int32) *AddImageLibResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImageLibResponse) SetBody(v *AddImageLibResponseBody) *AddImageLibResponse {
	s.Body = v
	return s
}

type AddImages2LibRequest struct {
	// example:
	//
	// upload/ea7a98f9-f8bd-4905-a79b-963c9da419c5.jpg
	ImgUrl *string `json:"ImgUrl,omitempty" xml:"ImgUrl,omitempty"`
	// example:
	//
	// xxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddImages2LibRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImages2LibRequest) GoString() string {
	return s.String()
}

func (s *AddImages2LibRequest) SetImgUrl(v string) *AddImages2LibRequest {
	s.ImgUrl = &v
	return s
}

func (s *AddImages2LibRequest) SetLibId(v string) *AddImages2LibRequest {
	s.LibId = &v
	return s
}

func (s *AddImages2LibRequest) SetRegionId(v string) *AddImages2LibRequest {
	s.RegionId = &v
	return s
}

type AddImages2LibResponseBody struct {
	// example:
	//
	// 200
	Code *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddImages2LibResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddImages2LibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddImages2LibResponseBody) GoString() string {
	return s.String()
}

func (s *AddImages2LibResponseBody) SetCode(v int32) *AddImages2LibResponseBody {
	s.Code = &v
	return s
}

func (s *AddImages2LibResponseBody) SetData(v *AddImages2LibResponseBodyData) *AddImages2LibResponseBody {
	s.Data = v
	return s
}

func (s *AddImages2LibResponseBody) SetHttpStatusCode(v int32) *AddImages2LibResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddImages2LibResponseBody) SetMsg(v string) *AddImages2LibResponseBody {
	s.Msg = &v
	return s
}

func (s *AddImages2LibResponseBody) SetRequestId(v string) *AddImages2LibResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImages2LibResponseBody) SetSuccess(v bool) *AddImages2LibResponseBody {
	s.Success = &v
	return s
}

type AddImages2LibResponseBodyData struct {
	// example:
	//
	// 100001
	ImgId *string `json:"ImgId,omitempty" xml:"ImgId,omitempty"`
}

func (s AddImages2LibResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddImages2LibResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddImages2LibResponseBodyData) SetImgId(v string) *AddImages2LibResponseBodyData {
	s.ImgId = &v
	return s
}

type AddImages2LibResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddImages2LibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddImages2LibResponse) String() string {
	return tea.Prettify(s)
}

func (s AddImages2LibResponse) GoString() string {
	return s.String()
}

func (s *AddImages2LibResponse) SetHeaders(v map[string]*string) *AddImages2LibResponse {
	s.Headers = v
	return s
}

func (s *AddImages2LibResponse) SetStatusCode(v int32) *AddImages2LibResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImages2LibResponse) SetBody(v *AddImages2LibResponseBody) *AddImages2LibResponse {
	s.Body = v
	return s
}

type AddKeywordLibRequest struct {
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// example:
	//
	// upload/1e5353c0-0d91-40ba-9d41-ae7abd3fe561.txt
	KeywordsObject *string `json:"KeywordsObject,omitempty" xml:"KeywordsObject,omitempty"`
	LibName        *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddKeywordLibRequest) String() string {
	return tea.Prettify(s)
}

func (s AddKeywordLibRequest) GoString() string {
	return s.String()
}

func (s *AddKeywordLibRequest) SetKeywords(v string) *AddKeywordLibRequest {
	s.Keywords = &v
	return s
}

func (s *AddKeywordLibRequest) SetKeywordsObject(v string) *AddKeywordLibRequest {
	s.KeywordsObject = &v
	return s
}

func (s *AddKeywordLibRequest) SetLibName(v string) *AddKeywordLibRequest {
	s.LibName = &v
	return s
}

func (s *AddKeywordLibRequest) SetRegionId(v string) *AddKeywordLibRequest {
	s.RegionId = &v
	return s
}

type AddKeywordLibResponseBody struct {
	// example:
	//
	// 200
	Code *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddKeywordLibResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddKeywordLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddKeywordLibResponseBody) GoString() string {
	return s.String()
}

func (s *AddKeywordLibResponseBody) SetCode(v int32) *AddKeywordLibResponseBody {
	s.Code = &v
	return s
}

func (s *AddKeywordLibResponseBody) SetData(v *AddKeywordLibResponseBodyData) *AddKeywordLibResponseBody {
	s.Data = v
	return s
}

func (s *AddKeywordLibResponseBody) SetMsg(v string) *AddKeywordLibResponseBody {
	s.Msg = &v
	return s
}

func (s *AddKeywordLibResponseBody) SetRequestId(v string) *AddKeywordLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddKeywordLibResponseBody) SetSuccess(v bool) *AddKeywordLibResponseBody {
	s.Success = &v
	return s
}

type AddKeywordLibResponseBodyData struct {
	KeywordsResult *AddKeywordLibResponseBodyDataKeywordsResult `json:"KeywordsResult,omitempty" xml:"KeywordsResult,omitempty" type:"Struct"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddKeywordLibResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddKeywordLibResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddKeywordLibResponseBodyData) SetKeywordsResult(v *AddKeywordLibResponseBodyDataKeywordsResult) *AddKeywordLibResponseBodyData {
	s.KeywordsResult = v
	return s
}

func (s *AddKeywordLibResponseBodyData) SetLibId(v string) *AddKeywordLibResponseBodyData {
	s.LibId = &v
	return s
}

func (s *AddKeywordLibResponseBodyData) SetTaskId(v string) *AddKeywordLibResponseBodyData {
	s.TaskId = &v
	return s
}

type AddKeywordLibResponseBodyDataKeywordsResult struct {
	// example:
	//
	// xxx
	I18nKey               *string   `json:"I18nKey,omitempty" xml:"I18nKey,omitempty"`
	IllegalLengthKeywords []*string `json:"IllegalLengthKeywords,omitempty" xml:"IllegalLengthKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// 133
	InvalidCount    *int32    `json:"InvalidCount,omitempty" xml:"InvalidCount,omitempty"`
	InvalidKeywords []*string `json:"InvalidKeywords,omitempty" xml:"InvalidKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 118
	RepeatCount    *int32    `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	RepeatKeywords []*string `json:"RepeatKeywords,omitempty" xml:"RepeatKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// 278
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// xxx
	Tips *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
	// example:
	//
	// 529
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s AddKeywordLibResponseBodyDataKeywordsResult) String() string {
	return tea.Prettify(s)
}

func (s AddKeywordLibResponseBodyDataKeywordsResult) GoString() string {
	return s.String()
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetI18nKey(v string) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.I18nKey = &v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetIllegalLengthKeywords(v []*string) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.IllegalLengthKeywords = v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetInvalidCount(v int32) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.InvalidCount = &v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetInvalidKeywords(v []*string) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.InvalidKeywords = v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetLibId(v string) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.LibId = &v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetRepeatCount(v int32) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.RepeatCount = &v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetRepeatKeywords(v []*string) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.RepeatKeywords = v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetSuccessCount(v int32) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.SuccessCount = &v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetTips(v string) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.Tips = &v
	return s
}

func (s *AddKeywordLibResponseBodyDataKeywordsResult) SetTotalCount(v int32) *AddKeywordLibResponseBodyDataKeywordsResult {
	s.TotalCount = &v
	return s
}

type AddKeywordLibResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddKeywordLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddKeywordLibResponse) String() string {
	return tea.Prettify(s)
}

func (s AddKeywordLibResponse) GoString() string {
	return s.String()
}

func (s *AddKeywordLibResponse) SetHeaders(v map[string]*string) *AddKeywordLibResponse {
	s.Headers = v
	return s
}

func (s *AddKeywordLibResponse) SetStatusCode(v int32) *AddKeywordLibResponse {
	s.StatusCode = &v
	return s
}

func (s *AddKeywordLibResponse) SetBody(v *AddKeywordLibResponseBody) *AddKeywordLibResponse {
	s.Body = v
	return s
}

type AddKeywordsRequest struct {
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// example:
	//
	// upload/1e5353c0-0d91-40ba-9d41-ae7abd3fe561.txt
	KeywordsObject *string `json:"KeywordsObject,omitempty" xml:"KeywordsObject,omitempty"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddKeywordsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddKeywordsRequest) GoString() string {
	return s.String()
}

func (s *AddKeywordsRequest) SetKeywords(v string) *AddKeywordsRequest {
	s.Keywords = &v
	return s
}

func (s *AddKeywordsRequest) SetKeywordsObject(v string) *AddKeywordsRequest {
	s.KeywordsObject = &v
	return s
}

func (s *AddKeywordsRequest) SetLibId(v string) *AddKeywordsRequest {
	s.LibId = &v
	return s
}

func (s *AddKeywordsRequest) SetRegionId(v string) *AddKeywordsRequest {
	s.RegionId = &v
	return s
}

type AddKeywordsResponseBody struct {
	// example:
	//
	// 200
	Code *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddKeywordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddKeywordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddKeywordsResponseBody) GoString() string {
	return s.String()
}

func (s *AddKeywordsResponseBody) SetCode(v int32) *AddKeywordsResponseBody {
	s.Code = &v
	return s
}

func (s *AddKeywordsResponseBody) SetData(v *AddKeywordsResponseBodyData) *AddKeywordsResponseBody {
	s.Data = v
	return s
}

func (s *AddKeywordsResponseBody) SetMsg(v string) *AddKeywordsResponseBody {
	s.Msg = &v
	return s
}

func (s *AddKeywordsResponseBody) SetRequestId(v string) *AddKeywordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddKeywordsResponseBody) SetSuccess(v bool) *AddKeywordsResponseBody {
	s.Success = &v
	return s
}

type AddKeywordsResponseBodyData struct {
	KeywordsResult *AddKeywordsResponseBodyDataKeywordsResult `json:"KeywordsResult,omitempty" xml:"KeywordsResult,omitempty" type:"Struct"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddKeywordsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddKeywordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddKeywordsResponseBodyData) SetKeywordsResult(v *AddKeywordsResponseBodyDataKeywordsResult) *AddKeywordsResponseBodyData {
	s.KeywordsResult = v
	return s
}

func (s *AddKeywordsResponseBodyData) SetLibId(v string) *AddKeywordsResponseBodyData {
	s.LibId = &v
	return s
}

func (s *AddKeywordsResponseBodyData) SetTaskId(v string) *AddKeywordsResponseBodyData {
	s.TaskId = &v
	return s
}

type AddKeywordsResponseBodyDataKeywordsResult struct {
	// example:
	//
	// xxx
	I18nKey               *string   `json:"I18nKey,omitempty" xml:"I18nKey,omitempty"`
	IllegalLengthKeywords []*string `json:"IllegalLengthKeywords,omitempty" xml:"IllegalLengthKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	InvalidCount    *int32    `json:"InvalidCount,omitempty" xml:"InvalidCount,omitempty"`
	InvalidKeywords []*string `json:"InvalidKeywords,omitempty" xml:"InvalidKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 1
	RepeatCount    *int32    `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	RepeatKeywords []*string `json:"RepeatKeywords,omitempty" xml:"RepeatKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// 6
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// xxxxx
	Tips *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
	// example:
	//
	// 8
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s AddKeywordsResponseBodyDataKeywordsResult) String() string {
	return tea.Prettify(s)
}

func (s AddKeywordsResponseBodyDataKeywordsResult) GoString() string {
	return s.String()
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetI18nKey(v string) *AddKeywordsResponseBodyDataKeywordsResult {
	s.I18nKey = &v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetIllegalLengthKeywords(v []*string) *AddKeywordsResponseBodyDataKeywordsResult {
	s.IllegalLengthKeywords = v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetInvalidCount(v int32) *AddKeywordsResponseBodyDataKeywordsResult {
	s.InvalidCount = &v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetInvalidKeywords(v []*string) *AddKeywordsResponseBodyDataKeywordsResult {
	s.InvalidKeywords = v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetLibId(v string) *AddKeywordsResponseBodyDataKeywordsResult {
	s.LibId = &v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetProgress(v int32) *AddKeywordsResponseBodyDataKeywordsResult {
	s.Progress = &v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetRepeatCount(v int32) *AddKeywordsResponseBodyDataKeywordsResult {
	s.RepeatCount = &v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetRepeatKeywords(v []*string) *AddKeywordsResponseBodyDataKeywordsResult {
	s.RepeatKeywords = v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetSuccessCount(v int32) *AddKeywordsResponseBodyDataKeywordsResult {
	s.SuccessCount = &v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetTips(v string) *AddKeywordsResponseBodyDataKeywordsResult {
	s.Tips = &v
	return s
}

func (s *AddKeywordsResponseBodyDataKeywordsResult) SetTotalCount(v int32) *AddKeywordsResponseBodyDataKeywordsResult {
	s.TotalCount = &v
	return s
}

type AddKeywordsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddKeywordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddKeywordsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddKeywordsResponse) GoString() string {
	return s.String()
}

func (s *AddKeywordsResponse) SetHeaders(v map[string]*string) *AddKeywordsResponse {
	s.Headers = v
	return s
}

func (s *AddKeywordsResponse) SetStatusCode(v int32) *AddKeywordsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddKeywordsResponse) SetBody(v *AddKeywordsResponseBody) *AddKeywordsResponse {
	s.Body = v
	return s
}

type AddKeywordsToLibRequest struct {
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// example:
	//
	// upload/1e5353c0-0d91-40ba-9d41-ae7abd3fe561.txt
	KeywordsObject *string `json:"KeywordsObject,omitempty" xml:"KeywordsObject,omitempty"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddKeywordsToLibRequest) String() string {
	return tea.Prettify(s)
}

func (s AddKeywordsToLibRequest) GoString() string {
	return s.String()
}

func (s *AddKeywordsToLibRequest) SetKeywords(v string) *AddKeywordsToLibRequest {
	s.Keywords = &v
	return s
}

func (s *AddKeywordsToLibRequest) SetKeywordsObject(v string) *AddKeywordsToLibRequest {
	s.KeywordsObject = &v
	return s
}

func (s *AddKeywordsToLibRequest) SetLibId(v string) *AddKeywordsToLibRequest {
	s.LibId = &v
	return s
}

func (s *AddKeywordsToLibRequest) SetRegionId(v string) *AddKeywordsToLibRequest {
	s.RegionId = &v
	return s
}

type AddKeywordsToLibResponseBody struct {
	Data *AddKeywordsToLibResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddKeywordsToLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddKeywordsToLibResponseBody) GoString() string {
	return s.String()
}

func (s *AddKeywordsToLibResponseBody) SetData(v *AddKeywordsToLibResponseBodyData) *AddKeywordsToLibResponseBody {
	s.Data = v
	return s
}

func (s *AddKeywordsToLibResponseBody) SetRequestId(v string) *AddKeywordsToLibResponseBody {
	s.RequestId = &v
	return s
}

type AddKeywordsToLibResponseBodyData struct {
	KeywordsResult *AddKeywordsToLibResponseBodyDataKeywordsResult `json:"KeywordsResult,omitempty" xml:"KeywordsResult,omitempty" type:"Struct"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddKeywordsToLibResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddKeywordsToLibResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddKeywordsToLibResponseBodyData) SetKeywordsResult(v *AddKeywordsToLibResponseBodyDataKeywordsResult) *AddKeywordsToLibResponseBodyData {
	s.KeywordsResult = v
	return s
}

func (s *AddKeywordsToLibResponseBodyData) SetLibId(v string) *AddKeywordsToLibResponseBodyData {
	s.LibId = &v
	return s
}

func (s *AddKeywordsToLibResponseBodyData) SetTaskId(v string) *AddKeywordsToLibResponseBodyData {
	s.TaskId = &v
	return s
}

type AddKeywordsToLibResponseBodyDataKeywordsResult struct {
	// example:
	//
	// xxx
	I18nKey               *string   `json:"I18nKey,omitempty" xml:"I18nKey,omitempty"`
	IllegalLengthKeywords []*string `json:"IllegalLengthKeywords,omitempty" xml:"IllegalLengthKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	InvalidCount    *int32    `json:"InvalidCount,omitempty" xml:"InvalidCount,omitempty"`
	InvalidKeywords []*string `json:"InvalidKeywords,omitempty" xml:"InvalidKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 1
	RepeatCount    *int32    `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	RepeatKeywords []*string `json:"RepeatKeywords,omitempty" xml:"RepeatKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// 8
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s AddKeywordsToLibResponseBodyDataKeywordsResult) String() string {
	return tea.Prettify(s)
}

func (s AddKeywordsToLibResponseBodyDataKeywordsResult) GoString() string {
	return s.String()
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetI18nKey(v string) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.I18nKey = &v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetIllegalLengthKeywords(v []*string) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.IllegalLengthKeywords = v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetInvalidCount(v int32) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.InvalidCount = &v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetInvalidKeywords(v []*string) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.InvalidKeywords = v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetLibId(v string) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.LibId = &v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetProgress(v int32) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.Progress = &v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetRepeatCount(v int32) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.RepeatCount = &v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetRepeatKeywords(v []*string) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.RepeatKeywords = v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetSuccessCount(v int32) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.SuccessCount = &v
	return s
}

func (s *AddKeywordsToLibResponseBodyDataKeywordsResult) SetTotalCount(v int32) *AddKeywordsToLibResponseBodyDataKeywordsResult {
	s.TotalCount = &v
	return s
}

type AddKeywordsToLibResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddKeywordsToLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddKeywordsToLibResponse) String() string {
	return tea.Prettify(s)
}

func (s AddKeywordsToLibResponse) GoString() string {
	return s.String()
}

func (s *AddKeywordsToLibResponse) SetHeaders(v map[string]*string) *AddKeywordsToLibResponse {
	s.Headers = v
	return s
}

func (s *AddKeywordsToLibResponse) SetStatusCode(v int32) *AddKeywordsToLibResponse {
	s.StatusCode = &v
	return s
}

func (s *AddKeywordsToLibResponse) SetBody(v *AddKeywordsToLibResponseBody) *AddKeywordsToLibResponse {
	s.Body = v
	return s
}

type CancelStockOssCheckTaskRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// P_UNYVB
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelStockOssCheckTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelStockOssCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelStockOssCheckTaskRequest) SetRegionId(v string) *CancelStockOssCheckTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CancelStockOssCheckTaskRequest) SetTaskId(v string) *CancelStockOssCheckTaskRequest {
	s.TaskId = &v
	return s
}

type CancelStockOssCheckTaskResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelStockOssCheckTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelStockOssCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelStockOssCheckTaskResponseBody) SetData(v bool) *CancelStockOssCheckTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CancelStockOssCheckTaskResponseBody) SetRequestId(v string) *CancelStockOssCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

type CancelStockOssCheckTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelStockOssCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelStockOssCheckTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelStockOssCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelStockOssCheckTaskResponse) SetHeaders(v map[string]*string) *CancelStockOssCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelStockOssCheckTaskResponse) SetStatusCode(v int32) *CancelStockOssCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelStockOssCheckTaskResponse) SetBody(v *CancelStockOssCheckTaskResponseBody) *CancelStockOssCheckTaskResponse {
	s.Body = v
	return s
}

type CopyServiceConfigRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// nickname_detection
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ServiceDesc *string `json:"ServiceDesc,omitempty" xml:"ServiceDesc,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CopyServiceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *CopyServiceConfigRequest) SetRegionId(v string) *CopyServiceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *CopyServiceConfigRequest) SetResourceType(v string) *CopyServiceConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *CopyServiceConfigRequest) SetServiceCode(v string) *CopyServiceConfigRequest {
	s.ServiceCode = &v
	return s
}

func (s *CopyServiceConfigRequest) SetServiceDesc(v string) *CopyServiceConfigRequest {
	s.ServiceDesc = &v
	return s
}

func (s *CopyServiceConfigRequest) SetServiceName(v string) *CopyServiceConfigRequest {
	s.ServiceName = &v
	return s
}

type CopyServiceConfigResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyServiceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CopyServiceConfigResponseBody) SetData(v bool) *CopyServiceConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CopyServiceConfigResponseBody) SetRequestId(v string) *CopyServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

type CopyServiceConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyServiceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *CopyServiceConfigResponse) SetHeaders(v map[string]*string) *CopyServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *CopyServiceConfigResponse) SetStatusCode(v int32) *CopyServiceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyServiceConfigResponse) SetBody(v *CopyServiceConfigResponseBody) *CopyServiceConfigResponse {
	s.Body = v
	return s
}

type CreatStockOssCheckTaskRequest struct {
	// example:
	//
	// [{\\"Bucket\\":\\"bucket01-test\\",\\"Region\\":\\"cn-beijing\\"}]
	Buckets *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	// example:
	//
	// 1751
	CallbackId *string `json:"CallbackId,omitempty" xml:"CallbackId,omitempty"`
	// example:
	//
	// true
	DistinctHistoryTasks *bool `json:"DistinctHistoryTasks,omitempty" xml:"DistinctHistoryTasks,omitempty"`
	// example:
	//
	// 2023-12-18 10:08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	ExecuteDate *int32 `json:"ExecuteDate,omitempty" xml:"ExecuteDate,omitempty"`
	// example:
	//
	// 01:09:30-01:19:30
	ExecuteTime       *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	Freeze            *bool   `json:"Freeze,omitempty" xml:"Freeze,omitempty"`
	FreezeHighRisk1   *bool   `json:"FreezeHighRisk1,omitempty" xml:"FreezeHighRisk1,omitempty"`
	FreezeHighRisk2   *bool   `json:"FreezeHighRisk2,omitempty" xml:"FreezeHighRisk2,omitempty"`
	FreezeMediumRisk1 *bool   `json:"FreezeMediumRisk1,omitempty" xml:"FreezeMediumRisk1,omitempty"`
	FreezeMediumRisk2 *bool   `json:"FreezeMediumRisk2,omitempty" xml:"FreezeMediumRisk2,omitempty"`
	FreezeType        *string `json:"FreezeType,omitempty" xml:"FreezeType,omitempty"`
	// example:
	//
	// false
	IsInc *bool `json:"IsInc,omitempty" xml:"IsInc,omitempty"`
	// example:
	//
	// 1
	MediaType *int32 `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// all
	PrefixFilterType *string `json:"PrefixFilterType,omitempty" xml:"PrefixFilterType,omitempty"`
	// example:
	//
	// dir1,dir2
	PrefixFilters *string `json:"PrefixFilters,omitempty" xml:"PrefixFilters,omitempty"`
	// example:
	//
	// 0
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 10
	ScanLimit *int64 `json:"ScanLimit,omitempty" xml:"ScanLimit,omitempty"`
	// example:
	//
	// true
	ScanNoFileType *bool `json:"ScanNoFileType,omitempty" xml:"ScanNoFileType,omitempty"`
	// example:
	//
	// 0
	ScanResourceType *string `json:"ScanResourceType,omitempty" xml:"ScanResourceType,omitempty"`
	// example:
	//
	// baselineCheck
	ScanService *string `json:"ScanService,omitempty" xml:"ScanService,omitempty"`
	// example:
	//
	// 2023-12-17 10:08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 0
	TaskCycle *int32  `json:"TaskCycle,omitempty" xml:"TaskCycle,omitempty"`
	TaskName  *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// batch
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreatStockOssCheckTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatStockOssCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *CreatStockOssCheckTaskRequest) SetBuckets(v string) *CreatStockOssCheckTaskRequest {
	s.Buckets = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetCallbackId(v string) *CreatStockOssCheckTaskRequest {
	s.CallbackId = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetDistinctHistoryTasks(v bool) *CreatStockOssCheckTaskRequest {
	s.DistinctHistoryTasks = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetEndTime(v string) *CreatStockOssCheckTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetExecuteDate(v int32) *CreatStockOssCheckTaskRequest {
	s.ExecuteDate = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetExecuteTime(v string) *CreatStockOssCheckTaskRequest {
	s.ExecuteTime = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetFreeze(v bool) *CreatStockOssCheckTaskRequest {
	s.Freeze = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetFreezeHighRisk1(v bool) *CreatStockOssCheckTaskRequest {
	s.FreezeHighRisk1 = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetFreezeHighRisk2(v bool) *CreatStockOssCheckTaskRequest {
	s.FreezeHighRisk2 = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetFreezeMediumRisk1(v bool) *CreatStockOssCheckTaskRequest {
	s.FreezeMediumRisk1 = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetFreezeMediumRisk2(v bool) *CreatStockOssCheckTaskRequest {
	s.FreezeMediumRisk2 = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetFreezeType(v string) *CreatStockOssCheckTaskRequest {
	s.FreezeType = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetIsInc(v bool) *CreatStockOssCheckTaskRequest {
	s.IsInc = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetMediaType(v int32) *CreatStockOssCheckTaskRequest {
	s.MediaType = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetPrefixFilterType(v string) *CreatStockOssCheckTaskRequest {
	s.PrefixFilterType = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetPrefixFilters(v string) *CreatStockOssCheckTaskRequest {
	s.PrefixFilters = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetPriority(v int32) *CreatStockOssCheckTaskRequest {
	s.Priority = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetRegionId(v string) *CreatStockOssCheckTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetScanLimit(v int64) *CreatStockOssCheckTaskRequest {
	s.ScanLimit = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetScanNoFileType(v bool) *CreatStockOssCheckTaskRequest {
	s.ScanNoFileType = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetScanResourceType(v string) *CreatStockOssCheckTaskRequest {
	s.ScanResourceType = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetScanService(v string) *CreatStockOssCheckTaskRequest {
	s.ScanService = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetStartTime(v string) *CreatStockOssCheckTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetTaskCycle(v int32) *CreatStockOssCheckTaskRequest {
	s.TaskCycle = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetTaskName(v string) *CreatStockOssCheckTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreatStockOssCheckTaskRequest) SetTaskType(v string) *CreatStockOssCheckTaskRequest {
	s.TaskType = &v
	return s
}

type CreatStockOssCheckTaskResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatStockOssCheckTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatStockOssCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreatStockOssCheckTaskResponseBody) SetData(v bool) *CreatStockOssCheckTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreatStockOssCheckTaskResponseBody) SetRequestId(v string) *CreatStockOssCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreatStockOssCheckTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatStockOssCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatStockOssCheckTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatStockOssCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *CreatStockOssCheckTaskResponse) SetHeaders(v map[string]*string) *CreatStockOssCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *CreatStockOssCheckTaskResponse) SetStatusCode(v int32) *CreatStockOssCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatStockOssCheckTaskResponse) SetBody(v *CreatStockOssCheckTaskResponseBody) *CreatStockOssCheckTaskResponse {
	s.Body = v
	return s
}

type CreateAnswerLibRequest struct {
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// oss-cip-shanghai
	SampleBucket *string `json:"SampleBucket,omitempty" xml:"SampleBucket,omitempty"`
	// example:
	//
	// data/xxx.xlsx
	SampleObject *string `json:"SampleObject,omitempty" xml:"SampleObject,omitempty"`
	Samples      *string `json:"Samples,omitempty" xml:"Samples,omitempty"`
}

func (s CreateAnswerLibRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAnswerLibRequest) GoString() string {
	return s.String()
}

func (s *CreateAnswerLibRequest) SetLibName(v string) *CreateAnswerLibRequest {
	s.LibName = &v
	return s
}

func (s *CreateAnswerLibRequest) SetRegionId(v string) *CreateAnswerLibRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAnswerLibRequest) SetSampleBucket(v string) *CreateAnswerLibRequest {
	s.SampleBucket = &v
	return s
}

func (s *CreateAnswerLibRequest) SetSampleObject(v string) *CreateAnswerLibRequest {
	s.SampleObject = &v
	return s
}

func (s *CreateAnswerLibRequest) SetSamples(v string) *CreateAnswerLibRequest {
	s.Samples = &v
	return s
}

type CreateAnswerLibResponseBody struct {
	// example:
	//
	// alxxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CreateAnswerLibResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// alAxbbxxxx-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateAnswerLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAnswerLibResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAnswerLibResponseBody) SetLibId(v string) *CreateAnswerLibResponseBody {
	s.LibId = &v
	return s
}

func (s *CreateAnswerLibResponseBody) SetRequestId(v string) *CreateAnswerLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAnswerLibResponseBody) SetResult(v *CreateAnswerLibResponseBodyResult) *CreateAnswerLibResponseBody {
	s.Result = v
	return s
}

func (s *CreateAnswerLibResponseBody) SetTaskId(v string) *CreateAnswerLibResponseBody {
	s.TaskId = &v
	return s
}

type CreateAnswerLibResponseBodyResult struct {
	// example:
	//
	// xxx
	I18nKey              *string   `json:"I18nKey,omitempty" xml:"I18nKey,omitempty"`
	IllegalLengthSamples []*string `json:"IllegalLengthSamples,omitempty" xml:"IllegalLengthSamples,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	InvalidCount *int32 `json:"InvalidCount,omitempty" xml:"InvalidCount,omitempty"`
	// example:
	//
	// alxxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 1
	RepeatCount   *int32    `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	RepeatSamples []*string `json:"RepeatSamples,omitempty" xml:"RepeatSamples,omitempty" type:"Repeated"`
	// example:
	//
	// 8
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// alAaaaxxx-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CreateAnswerLibResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAnswerLibResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAnswerLibResponseBodyResult) SetI18nKey(v string) *CreateAnswerLibResponseBodyResult {
	s.I18nKey = &v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetIllegalLengthSamples(v []*string) *CreateAnswerLibResponseBodyResult {
	s.IllegalLengthSamples = v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetInvalidCount(v int32) *CreateAnswerLibResponseBodyResult {
	s.InvalidCount = &v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetLibId(v string) *CreateAnswerLibResponseBodyResult {
	s.LibId = &v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetProgress(v int32) *CreateAnswerLibResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetRepeatCount(v int32) *CreateAnswerLibResponseBodyResult {
	s.RepeatCount = &v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetRepeatSamples(v []*string) *CreateAnswerLibResponseBodyResult {
	s.RepeatSamples = v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetSuccessCount(v int32) *CreateAnswerLibResponseBodyResult {
	s.SuccessCount = &v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetTaskId(v string) *CreateAnswerLibResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *CreateAnswerLibResponseBodyResult) SetTotalCount(v int32) *CreateAnswerLibResponseBodyResult {
	s.TotalCount = &v
	return s
}

type CreateAnswerLibResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAnswerLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAnswerLibResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAnswerLibResponse) GoString() string {
	return s.String()
}

func (s *CreateAnswerLibResponse) SetHeaders(v map[string]*string) *CreateAnswerLibResponse {
	s.Headers = v
	return s
}

func (s *CreateAnswerLibResponse) SetStatusCode(v int32) *CreateAnswerLibResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAnswerLibResponse) SetBody(v *CreateAnswerLibResponseBody) *CreateAnswerLibResponse {
	s.Body = v
	return s
}

type CreateCallbackRequest struct {
	// example:
	//
	// SHA256
	CryptType *string `json:"CryptType,omitempty" xml:"CryptType,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// all
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// https://console.aliyun.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackRequest) GoString() string {
	return s.String()
}

func (s *CreateCallbackRequest) SetCryptType(v string) *CreateCallbackRequest {
	s.CryptType = &v
	return s
}

func (s *CreateCallbackRequest) SetName(v string) *CreateCallbackRequest {
	s.Name = &v
	return s
}

func (s *CreateCallbackRequest) SetRegionId(v string) *CreateCallbackRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCallbackRequest) SetScope(v string) *CreateCallbackRequest {
	s.Scope = &v
	return s
}

func (s *CreateCallbackRequest) SetUrl(v string) *CreateCallbackRequest {
	s.Url = &v
	return s
}

type CreateCallbackResponseBody struct {
	// example:
	//
	// True
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCallbackResponseBody) SetData(v int64) *CreateCallbackResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCallbackResponseBody) SetRequestId(v string) *CreateCallbackResponseBody {
	s.RequestId = &v
	return s
}

type CreateCallbackResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackResponse) GoString() string {
	return s.String()
}

func (s *CreateCallbackResponse) SetHeaders(v map[string]*string) *CreateCallbackResponse {
	s.Headers = v
	return s
}

func (s *CreateCallbackResponse) SetStatusCode(v int32) *CreateCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCallbackResponse) SetBody(v *CreateCallbackResponseBody) *CreateCallbackResponse {
	s.Body = v
	return s
}

type CreatePreCheckRequest struct {
	// example:
	//
	// [{\\"Bucket\\":\\"bucket01-test\\",\\"Region\\":\\"cn-beijing\\"}]
	Buckets *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	// example:
	//
	// true
	DistinctHistoryTasks *bool `json:"DistinctHistoryTasks,omitempty" xml:"DistinctHistoryTasks,omitempty"`
	// example:
	//
	// 2023-12-18 10:08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// false
	IsInc *bool `json:"IsInc,omitempty" xml:"IsInc,omitempty"`
	// example:
	//
	// 1
	MediaType *int32 `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// all
	PrefixFilterType *string `json:"PrefixFilterType,omitempty" xml:"PrefixFilterType,omitempty"`
	// example:
	//
	// dir1,dir2
	PrefixFilters *string `json:"PrefixFilters,omitempty" xml:"PrefixFilters,omitempty"`
	// example:
	//
	// 0
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 10
	ScanLimit *int64 `json:"ScanLimit,omitempty" xml:"ScanLimit,omitempty"`
	// example:
	//
	// true
	ScanNoFileType *bool `json:"ScanNoFileType,omitempty" xml:"ScanNoFileType,omitempty"`
	// example:
	//
	// baselineCheck
	ScanService *string `json:"ScanService,omitempty" xml:"ScanService,omitempty"`
	// example:
	//
	// 2023-12-17 10:08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskName  *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreatePreCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePreCheckRequest) GoString() string {
	return s.String()
}

func (s *CreatePreCheckRequest) SetBuckets(v string) *CreatePreCheckRequest {
	s.Buckets = &v
	return s
}

func (s *CreatePreCheckRequest) SetDistinctHistoryTasks(v bool) *CreatePreCheckRequest {
	s.DistinctHistoryTasks = &v
	return s
}

func (s *CreatePreCheckRequest) SetEndTime(v string) *CreatePreCheckRequest {
	s.EndTime = &v
	return s
}

func (s *CreatePreCheckRequest) SetIsInc(v bool) *CreatePreCheckRequest {
	s.IsInc = &v
	return s
}

func (s *CreatePreCheckRequest) SetMediaType(v int32) *CreatePreCheckRequest {
	s.MediaType = &v
	return s
}

func (s *CreatePreCheckRequest) SetPrefixFilterType(v string) *CreatePreCheckRequest {
	s.PrefixFilterType = &v
	return s
}

func (s *CreatePreCheckRequest) SetPrefixFilters(v string) *CreatePreCheckRequest {
	s.PrefixFilters = &v
	return s
}

func (s *CreatePreCheckRequest) SetPriority(v int32) *CreatePreCheckRequest {
	s.Priority = &v
	return s
}

func (s *CreatePreCheckRequest) SetRegionId(v string) *CreatePreCheckRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePreCheckRequest) SetScanLimit(v int64) *CreatePreCheckRequest {
	s.ScanLimit = &v
	return s
}

func (s *CreatePreCheckRequest) SetScanNoFileType(v bool) *CreatePreCheckRequest {
	s.ScanNoFileType = &v
	return s
}

func (s *CreatePreCheckRequest) SetScanService(v string) *CreatePreCheckRequest {
	s.ScanService = &v
	return s
}

func (s *CreatePreCheckRequest) SetStartTime(v string) *CreatePreCheckRequest {
	s.StartTime = &v
	return s
}

func (s *CreatePreCheckRequest) SetTaskName(v string) *CreatePreCheckRequest {
	s.TaskName = &v
	return s
}

type CreatePreCheckResponseBody struct {
	Data map[string]*bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePreCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePreCheckResponseBody) SetData(v map[string]*bool) *CreatePreCheckResponseBody {
	s.Data = v
	return s
}

func (s *CreatePreCheckResponseBody) SetRequestId(v string) *CreatePreCheckResponseBody {
	s.RequestId = &v
	return s
}

type CreatePreCheckResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePreCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePreCheckResponse) GoString() string {
	return s.String()
}

func (s *CreatePreCheckResponse) SetHeaders(v map[string]*string) *CreatePreCheckResponse {
	s.Headers = v
	return s
}

func (s *CreatePreCheckResponse) SetStatusCode(v int32) *CreatePreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePreCheckResponse) SetBody(v *CreatePreCheckResponseBody) *CreatePreCheckResponse {
	s.Body = v
	return s
}

type DeleteAnswerLibRequest struct {
	// example:
	//
	// alxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAnswerLibRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAnswerLibRequest) GoString() string {
	return s.String()
}

func (s *DeleteAnswerLibRequest) SetLibId(v string) *DeleteAnswerLibRequest {
	s.LibId = &v
	return s
}

func (s *DeleteAnswerLibRequest) SetRegionId(v string) *DeleteAnswerLibRequest {
	s.RegionId = &v
	return s
}

type DeleteAnswerLibResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAnswerLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAnswerLibResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAnswerLibResponseBody) SetData(v bool) *DeleteAnswerLibResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAnswerLibResponseBody) SetRequestId(v string) *DeleteAnswerLibResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAnswerLibResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAnswerLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAnswerLibResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAnswerLibResponse) GoString() string {
	return s.String()
}

func (s *DeleteAnswerLibResponse) SetHeaders(v map[string]*string) *DeleteAnswerLibResponse {
	s.Headers = v
	return s
}

func (s *DeleteAnswerLibResponse) SetStatusCode(v int32) *DeleteAnswerLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAnswerLibResponse) SetBody(v *DeleteAnswerLibResponseBody) *DeleteAnswerLibResponse {
	s.Body = v
	return s
}

type DeleteAnswerSampleRequest struct {
	// example:
	//
	// [15463605]
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// example:
	//
	// alxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAnswerSampleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAnswerSampleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAnswerSampleRequest) SetIds(v string) *DeleteAnswerSampleRequest {
	s.Ids = &v
	return s
}

func (s *DeleteAnswerSampleRequest) SetLibId(v string) *DeleteAnswerSampleRequest {
	s.LibId = &v
	return s
}

func (s *DeleteAnswerSampleRequest) SetRegionId(v string) *DeleteAnswerSampleRequest {
	s.RegionId = &v
	return s
}

type DeleteAnswerSampleResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAnswerSampleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAnswerSampleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAnswerSampleResponseBody) SetData(v bool) *DeleteAnswerSampleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAnswerSampleResponseBody) SetRequestId(v string) *DeleteAnswerSampleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAnswerSampleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAnswerSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAnswerSampleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAnswerSampleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAnswerSampleResponse) SetHeaders(v map[string]*string) *DeleteAnswerSampleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAnswerSampleResponse) SetStatusCode(v int32) *DeleteAnswerSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAnswerSampleResponse) SetBody(v *DeleteAnswerSampleResponseBody) *DeleteAnswerSampleResponse {
	s.Body = v
	return s
}

type DeleteCallbackRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1480
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCallbackRequest) GoString() string {
	return s.String()
}

func (s *DeleteCallbackRequest) SetId(v int64) *DeleteCallbackRequest {
	s.Id = &v
	return s
}

func (s *DeleteCallbackRequest) SetRegionId(v string) *DeleteCallbackRequest {
	s.RegionId = &v
	return s
}

type DeleteCallbackResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCallbackResponseBody) SetData(v bool) *DeleteCallbackResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCallbackResponseBody) SetRequestId(v string) *DeleteCallbackResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCallbackResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCallbackResponse) GoString() string {
	return s.String()
}

func (s *DeleteCallbackResponse) SetHeaders(v map[string]*string) *DeleteCallbackResponse {
	s.Headers = v
	return s
}

func (s *DeleteCallbackResponse) SetStatusCode(v int32) *DeleteCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCallbackResponse) SetBody(v *DeleteCallbackResponseBody) *DeleteCallbackResponse {
	s.Body = v
	return s
}

type DeleteFeatureConfigRequest struct {
	// Label value, customer-defined
	//
	// example:
	//
	// __config__
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// Region
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource type.
	//
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Service code.
	//
	// example:
	//
	// llm_query_moderation
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Type
	//
	// example:
	//
	// custom_llm_template
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteFeatureConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFeatureConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteFeatureConfigRequest) SetField(v string) *DeleteFeatureConfigRequest {
	s.Field = &v
	return s
}

func (s *DeleteFeatureConfigRequest) SetRegionId(v string) *DeleteFeatureConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFeatureConfigRequest) SetResourceType(v string) *DeleteFeatureConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *DeleteFeatureConfigRequest) SetServiceCode(v string) *DeleteFeatureConfigRequest {
	s.ServiceCode = &v
	return s
}

func (s *DeleteFeatureConfigRequest) SetType(v string) *DeleteFeatureConfigRequest {
	s.Type = &v
	return s
}

type DeleteFeatureConfigResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return result.
	//
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message of this request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// ID assigned by the backend, used to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success indicator.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFeatureConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFeatureConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFeatureConfigResponseBody) SetCode(v int32) *DeleteFeatureConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFeatureConfigResponseBody) SetData(v bool) *DeleteFeatureConfigResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFeatureConfigResponseBody) SetHttpStatusCode(v int32) *DeleteFeatureConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteFeatureConfigResponseBody) SetMsg(v string) *DeleteFeatureConfigResponseBody {
	s.Msg = &v
	return s
}

func (s *DeleteFeatureConfigResponseBody) SetRequestId(v string) *DeleteFeatureConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFeatureConfigResponseBody) SetSuccess(v bool) *DeleteFeatureConfigResponseBody {
	s.Success = &v
	return s
}

type DeleteFeatureConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFeatureConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFeatureConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFeatureConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteFeatureConfigResponse) SetHeaders(v map[string]*string) *DeleteFeatureConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteFeatureConfigResponse) SetStatusCode(v int32) *DeleteFeatureConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFeatureConfigResponse) SetBody(v *DeleteFeatureConfigResponseBody) *DeleteFeatureConfigResponse {
	s.Body = v
	return s
}

type DeleteImagesFromLibRequest struct {
	// example:
	//
	// [158794]
	ImageIds *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteImagesFromLibRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesFromLibRequest) GoString() string {
	return s.String()
}

func (s *DeleteImagesFromLibRequest) SetImageIds(v string) *DeleteImagesFromLibRequest {
	s.ImageIds = &v
	return s
}

func (s *DeleteImagesFromLibRequest) SetLibId(v string) *DeleteImagesFromLibRequest {
	s.LibId = &v
	return s
}

func (s *DeleteImagesFromLibRequest) SetRegionId(v string) *DeleteImagesFromLibRequest {
	s.RegionId = &v
	return s
}

type DeleteImagesFromLibResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteImagesFromLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesFromLibResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImagesFromLibResponseBody) SetCode(v int32) *DeleteImagesFromLibResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteImagesFromLibResponseBody) SetData(v bool) *DeleteImagesFromLibResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteImagesFromLibResponseBody) SetHttpStatusCode(v int32) *DeleteImagesFromLibResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteImagesFromLibResponseBody) SetMsg(v string) *DeleteImagesFromLibResponseBody {
	s.Msg = &v
	return s
}

func (s *DeleteImagesFromLibResponseBody) SetRequestId(v string) *DeleteImagesFromLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImagesFromLibResponseBody) SetSuccess(v bool) *DeleteImagesFromLibResponseBody {
	s.Success = &v
	return s
}

type DeleteImagesFromLibResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteImagesFromLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteImagesFromLibResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesFromLibResponse) GoString() string {
	return s.String()
}

func (s *DeleteImagesFromLibResponse) SetHeaders(v map[string]*string) *DeleteImagesFromLibResponse {
	s.Headers = v
	return s
}

func (s *DeleteImagesFromLibResponse) SetStatusCode(v int32) *DeleteImagesFromLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImagesFromLibResponse) SetBody(v *DeleteImagesFromLibResponseBody) *DeleteImagesFromLibResponse {
	s.Body = v
	return s
}

type DeleteKeywordRequest struct {
	KeywordIdList *string `json:"KeywordIdList,omitempty" xml:"KeywordIdList,omitempty"`
	// example:
	//
	// [16754493]
	KeywordIds *string `json:"KeywordIds,omitempty" xml:"KeywordIds,omitempty"`
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteKeywordRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeywordRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeywordRequest) SetKeywordIdList(v string) *DeleteKeywordRequest {
	s.KeywordIdList = &v
	return s
}

func (s *DeleteKeywordRequest) SetKeywordIds(v string) *DeleteKeywordRequest {
	s.KeywordIds = &v
	return s
}

func (s *DeleteKeywordRequest) SetLibId(v string) *DeleteKeywordRequest {
	s.LibId = &v
	return s
}

func (s *DeleteKeywordRequest) SetRegionId(v string) *DeleteKeywordRequest {
	s.RegionId = &v
	return s
}

type DeleteKeywordResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteKeywordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeywordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKeywordResponseBody) SetCode(v int32) *DeleteKeywordResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteKeywordResponseBody) SetData(v bool) *DeleteKeywordResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteKeywordResponseBody) SetMsg(v string) *DeleteKeywordResponseBody {
	s.Msg = &v
	return s
}

func (s *DeleteKeywordResponseBody) SetRequestId(v string) *DeleteKeywordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKeywordResponseBody) SetSuccess(v bool) *DeleteKeywordResponseBody {
	s.Success = &v
	return s
}

type DeleteKeywordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKeywordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKeywordResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeywordResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeywordResponse) SetHeaders(v map[string]*string) *DeleteKeywordResponse {
	s.Headers = v
	return s
}

func (s *DeleteKeywordResponse) SetStatusCode(v int32) *DeleteKeywordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKeywordResponse) SetBody(v *DeleteKeywordResponseBody) *DeleteKeywordResponse {
	s.Body = v
	return s
}

type DeleteKeywordLibRequest struct {
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteKeywordLibRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeywordLibRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeywordLibRequest) SetLibId(v string) *DeleteKeywordLibRequest {
	s.LibId = &v
	return s
}

func (s *DeleteKeywordLibRequest) SetRegionId(v string) *DeleteKeywordLibRequest {
	s.RegionId = &v
	return s
}

type DeleteKeywordLibResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteKeywordLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeywordLibResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKeywordLibResponseBody) SetCode(v int32) *DeleteKeywordLibResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteKeywordLibResponseBody) SetData(v bool) *DeleteKeywordLibResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteKeywordLibResponseBody) SetMsg(v string) *DeleteKeywordLibResponseBody {
	s.Msg = &v
	return s
}

func (s *DeleteKeywordLibResponseBody) SetRequestId(v string) *DeleteKeywordLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKeywordLibResponseBody) SetSuccess(v bool) *DeleteKeywordLibResponseBody {
	s.Success = &v
	return s
}

type DeleteKeywordLibResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKeywordLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKeywordLibResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeywordLibResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeywordLibResponse) SetHeaders(v map[string]*string) *DeleteKeywordLibResponse {
	s.Headers = v
	return s
}

func (s *DeleteKeywordLibResponse) SetStatusCode(v int32) *DeleteKeywordLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKeywordLibResponse) SetBody(v *DeleteKeywordLibResponseBody) *DeleteKeywordLibResponse {
	s.Body = v
	return s
}

type ExportAnswerSampleRequest struct {
	// example:
	//
	// alxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ExportAnswerSampleRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportAnswerSampleRequest) GoString() string {
	return s.String()
}

func (s *ExportAnswerSampleRequest) SetLibId(v string) *ExportAnswerSampleRequest {
	s.LibId = &v
	return s
}

func (s *ExportAnswerSampleRequest) SetRegionId(v string) *ExportAnswerSampleRequest {
	s.RegionId = &v
	return s
}

type ExportAnswerSampleResponseBody struct {
	// example:
	//
	// True
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportAnswerSampleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportAnswerSampleResponseBody) GoString() string {
	return s.String()
}

func (s *ExportAnswerSampleResponseBody) SetData(v string) *ExportAnswerSampleResponseBody {
	s.Data = &v
	return s
}

func (s *ExportAnswerSampleResponseBody) SetRequestId(v string) *ExportAnswerSampleResponseBody {
	s.RequestId = &v
	return s
}

type ExportAnswerSampleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportAnswerSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportAnswerSampleResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportAnswerSampleResponse) GoString() string {
	return s.String()
}

func (s *ExportAnswerSampleResponse) SetHeaders(v map[string]*string) *ExportAnswerSampleResponse {
	s.Headers = v
	return s
}

func (s *ExportAnswerSampleResponse) SetStatusCode(v int32) *ExportAnswerSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportAnswerSampleResponse) SetBody(v *ExportAnswerSampleResponseBody) *ExportAnswerSampleResponse {
	s.Body = v
	return s
}

type ExportCipStatsRequest struct {
	// example:
	//
	// true
	ByMonth *bool `json:"ByMonth,omitempty" xml:"ByMonth,omitempty"`
	// example:
	//
	// 2024-04-16 09:00:00
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// example:
	//
	// xx
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ServiceCode  *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// 2024-04-15 09:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 268220485413130979
	SubUid *string `json:"SubUid,omitempty" xml:"SubUid,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ExportCipStatsRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportCipStatsRequest) GoString() string {
	return s.String()
}

func (s *ExportCipStatsRequest) SetByMonth(v bool) *ExportCipStatsRequest {
	s.ByMonth = &v
	return s
}

func (s *ExportCipStatsRequest) SetEndDate(v string) *ExportCipStatsRequest {
	s.EndDate = &v
	return s
}

func (s *ExportCipStatsRequest) SetExportType(v string) *ExportCipStatsRequest {
	s.ExportType = &v
	return s
}

func (s *ExportCipStatsRequest) SetLabel(v string) *ExportCipStatsRequest {
	s.Label = &v
	return s
}

func (s *ExportCipStatsRequest) SetRegionId(v string) *ExportCipStatsRequest {
	s.RegionId = &v
	return s
}

func (s *ExportCipStatsRequest) SetResourceType(v string) *ExportCipStatsRequest {
	s.ResourceType = &v
	return s
}

func (s *ExportCipStatsRequest) SetServiceCode(v string) *ExportCipStatsRequest {
	s.ServiceCode = &v
	return s
}

func (s *ExportCipStatsRequest) SetStartDate(v string) *ExportCipStatsRequest {
	s.StartDate = &v
	return s
}

func (s *ExportCipStatsRequest) SetSubUid(v string) *ExportCipStatsRequest {
	s.SubUid = &v
	return s
}

func (s *ExportCipStatsRequest) SetType(v string) *ExportCipStatsRequest {
	s.Type = &v
	return s
}

type ExportCipStatsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/portal_data/production/cipStat/text/statistics1720597246783.xlsx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportCipStatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportCipStatsResponseBody) GoString() string {
	return s.String()
}

func (s *ExportCipStatsResponseBody) SetCode(v int32) *ExportCipStatsResponseBody {
	s.Code = &v
	return s
}

func (s *ExportCipStatsResponseBody) SetData(v string) *ExportCipStatsResponseBody {
	s.Data = &v
	return s
}

func (s *ExportCipStatsResponseBody) SetHttpStatusCode(v int32) *ExportCipStatsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExportCipStatsResponseBody) SetMsg(v string) *ExportCipStatsResponseBody {
	s.Msg = &v
	return s
}

func (s *ExportCipStatsResponseBody) SetRequestId(v string) *ExportCipStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportCipStatsResponseBody) SetSuccess(v bool) *ExportCipStatsResponseBody {
	s.Success = &v
	return s
}

type ExportCipStatsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportCipStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportCipStatsResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportCipStatsResponse) GoString() string {
	return s.String()
}

func (s *ExportCipStatsResponse) SetHeaders(v map[string]*string) *ExportCipStatsResponse {
	s.Headers = v
	return s
}

func (s *ExportCipStatsResponse) SetStatusCode(v int32) *ExportCipStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportCipStatsResponse) SetBody(v *ExportCipStatsResponseBody) *ExportCipStatsResponse {
	s.Body = v
	return s
}

type ExportKeywordRequest struct {
	// example:
	//
	// customxx_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ExportKeywordRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportKeywordRequest) GoString() string {
	return s.String()
}

func (s *ExportKeywordRequest) SetLibId(v string) *ExportKeywordRequest {
	s.LibId = &v
	return s
}

func (s *ExportKeywordRequest) SetRegionId(v string) *ExportKeywordRequest {
	s.RegionId = &v
	return s
}

type ExportKeywordResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/console_data/export/production/keyword/export_keywordO4ee1Bok1R8IIDVpcT9viU-1xxWr
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportKeywordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportKeywordResponseBody) GoString() string {
	return s.String()
}

func (s *ExportKeywordResponseBody) SetCode(v int32) *ExportKeywordResponseBody {
	s.Code = &v
	return s
}

func (s *ExportKeywordResponseBody) SetData(v string) *ExportKeywordResponseBody {
	s.Data = &v
	return s
}

func (s *ExportKeywordResponseBody) SetMsg(v string) *ExportKeywordResponseBody {
	s.Msg = &v
	return s
}

func (s *ExportKeywordResponseBody) SetRequestId(v string) *ExportKeywordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportKeywordResponseBody) SetSuccess(v bool) *ExportKeywordResponseBody {
	s.Success = &v
	return s
}

type ExportKeywordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportKeywordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportKeywordResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportKeywordResponse) GoString() string {
	return s.String()
}

func (s *ExportKeywordResponse) SetHeaders(v map[string]*string) *ExportKeywordResponse {
	s.Headers = v
	return s
}

func (s *ExportKeywordResponse) SetStatusCode(v int32) *ExportKeywordResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportKeywordResponse) SetBody(v *ExportKeywordResponseBody) *ExportKeywordResponse {
	s.Body = v
	return s
}

type ExportOssCheckStatRequest struct {
	// example:
	//
	// true
	ByMonth *bool `json:"ByMonth,omitempty" xml:"ByMonth,omitempty"`
	// example:
	//
	// 2024-03-11 10:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// P_UX0K5X
	ParentTaskId *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2024-03-10 10:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportOssCheckStatRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportOssCheckStatRequest) GoString() string {
	return s.String()
}

func (s *ExportOssCheckStatRequest) SetByMonth(v bool) *ExportOssCheckStatRequest {
	s.ByMonth = &v
	return s
}

func (s *ExportOssCheckStatRequest) SetEndDate(v string) *ExportOssCheckStatRequest {
	s.EndDate = &v
	return s
}

func (s *ExportOssCheckStatRequest) SetParentTaskId(v string) *ExportOssCheckStatRequest {
	s.ParentTaskId = &v
	return s
}

func (s *ExportOssCheckStatRequest) SetRegionId(v string) *ExportOssCheckStatRequest {
	s.RegionId = &v
	return s
}

func (s *ExportOssCheckStatRequest) SetStartDate(v string) *ExportOssCheckStatRequest {
	s.StartDate = &v
	return s
}

type ExportOssCheckStatResponseBody struct {
	// example:
	//
	// https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/console_data/production/scanResult/osscheck/ossCheckStat_aliUf5B3lJfOkLpqozLIn94Uy-1XxKyX.xlsx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportOssCheckStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportOssCheckStatResponseBody) GoString() string {
	return s.String()
}

func (s *ExportOssCheckStatResponseBody) SetData(v string) *ExportOssCheckStatResponseBody {
	s.Data = &v
	return s
}

func (s *ExportOssCheckStatResponseBody) SetRequestId(v string) *ExportOssCheckStatResponseBody {
	s.RequestId = &v
	return s
}

type ExportOssCheckStatResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportOssCheckStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportOssCheckStatResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportOssCheckStatResponse) GoString() string {
	return s.String()
}

func (s *ExportOssCheckStatResponse) SetHeaders(v map[string]*string) *ExportOssCheckStatResponse {
	s.Headers = v
	return s
}

func (s *ExportOssCheckStatResponse) SetStatusCode(v int32) *ExportOssCheckStatResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportOssCheckStatResponse) SetBody(v *ExportOssCheckStatResponseBody) *ExportOssCheckStatResponse {
	s.Body = v
	return s
}

type ExportResultRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// {"TaskId":"P_11TL5T"}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sort     map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportResultRequest) GoString() string {
	return s.String()
}

func (s *ExportResultRequest) SetCurrentPage(v int32) *ExportResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *ExportResultRequest) SetEndDate(v string) *ExportResultRequest {
	s.EndDate = &v
	return s
}

func (s *ExportResultRequest) SetPageSize(v int32) *ExportResultRequest {
	s.PageSize = &v
	return s
}

func (s *ExportResultRequest) SetQuery(v string) *ExportResultRequest {
	s.Query = &v
	return s
}

func (s *ExportResultRequest) SetRegionId(v string) *ExportResultRequest {
	s.RegionId = &v
	return s
}

func (s *ExportResultRequest) SetSort(v map[string]*string) *ExportResultRequest {
	s.Sort = v
	return s
}

func (s *ExportResultRequest) SetStartDate(v string) *ExportResultRequest {
	s.StartDate = &v
	return s
}

type ExportResultShrinkRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// {"TaskId":"P_11TL5T"}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportResultShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExportResultShrinkRequest) SetCurrentPage(v int32) *ExportResultShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ExportResultShrinkRequest) SetEndDate(v string) *ExportResultShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *ExportResultShrinkRequest) SetPageSize(v int32) *ExportResultShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ExportResultShrinkRequest) SetQuery(v string) *ExportResultShrinkRequest {
	s.Query = &v
	return s
}

func (s *ExportResultShrinkRequest) SetRegionId(v string) *ExportResultShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ExportResultShrinkRequest) SetSortShrink(v string) *ExportResultShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *ExportResultShrinkRequest) SetStartDate(v string) *ExportResultShrinkRequest {
	s.StartDate = &v
	return s
}

type ExportResultResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/console_data/production/scanResult/osscheck/ossCheckResult_aliiGGXhSMvmIvsS7Lrl3LaUZ-1A9%24oZ.xlsx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportResultResponseBody) GoString() string {
	return s.String()
}

func (s *ExportResultResponseBody) SetCode(v int32) *ExportResultResponseBody {
	s.Code = &v
	return s
}

func (s *ExportResultResponseBody) SetData(v string) *ExportResultResponseBody {
	s.Data = &v
	return s
}

func (s *ExportResultResponseBody) SetMsg(v string) *ExportResultResponseBody {
	s.Msg = &v
	return s
}

func (s *ExportResultResponseBody) SetRequestId(v string) *ExportResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportResultResponseBody) SetSuccess(v bool) *ExportResultResponseBody {
	s.Success = &v
	return s
}

type ExportResultResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportResultResponse) GoString() string {
	return s.String()
}

func (s *ExportResultResponse) SetHeaders(v map[string]*string) *ExportResultResponse {
	s.Headers = v
	return s
}

func (s *ExportResultResponse) SetStatusCode(v int32) *ExportResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportResultResponse) SetBody(v *ExportResultResponseBody) *ExportResultResponse {
	s.Body = v
	return s
}

type ExportScanResultRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2024-03-11 10:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 20
	PageSize *int32             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query    map[string]*string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// text
	ResourceType *string            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Sort         map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2024-03-10 10:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportScanResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportScanResultRequest) GoString() string {
	return s.String()
}

func (s *ExportScanResultRequest) SetCurrentPage(v int32) *ExportScanResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *ExportScanResultRequest) SetEndDate(v string) *ExportScanResultRequest {
	s.EndDate = &v
	return s
}

func (s *ExportScanResultRequest) SetPageSize(v int32) *ExportScanResultRequest {
	s.PageSize = &v
	return s
}

func (s *ExportScanResultRequest) SetQuery(v map[string]*string) *ExportScanResultRequest {
	s.Query = v
	return s
}

func (s *ExportScanResultRequest) SetRegionId(v string) *ExportScanResultRequest {
	s.RegionId = &v
	return s
}

func (s *ExportScanResultRequest) SetResourceType(v string) *ExportScanResultRequest {
	s.ResourceType = &v
	return s
}

func (s *ExportScanResultRequest) SetSort(v map[string]*string) *ExportScanResultRequest {
	s.Sort = v
	return s
}

func (s *ExportScanResultRequest) SetStartDate(v string) *ExportScanResultRequest {
	s.StartDate = &v
	return s
}

type ExportScanResultShrinkRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2024-03-11 10:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 20
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryShrink *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SortShrink   *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2024-03-10 10:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportScanResultShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportScanResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExportScanResultShrinkRequest) SetCurrentPage(v int32) *ExportScanResultShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ExportScanResultShrinkRequest) SetEndDate(v string) *ExportScanResultShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *ExportScanResultShrinkRequest) SetPageSize(v int32) *ExportScanResultShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ExportScanResultShrinkRequest) SetQueryShrink(v string) *ExportScanResultShrinkRequest {
	s.QueryShrink = &v
	return s
}

func (s *ExportScanResultShrinkRequest) SetRegionId(v string) *ExportScanResultShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ExportScanResultShrinkRequest) SetResourceType(v string) *ExportScanResultShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ExportScanResultShrinkRequest) SetSortShrink(v string) *ExportScanResultShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *ExportScanResultShrinkRequest) SetStartDate(v string) *ExportScanResultShrinkRequest {
	s.StartDate = &v
	return s
}

type ExportScanResultResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/portal_data/production/scanResult/text/textScanResult_aliow2MAdWXCakCxlitVY8Lnn-1A9KEw.xlsx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportScanResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportScanResultResponseBody) GoString() string {
	return s.String()
}

func (s *ExportScanResultResponseBody) SetCode(v int32) *ExportScanResultResponseBody {
	s.Code = &v
	return s
}

func (s *ExportScanResultResponseBody) SetData(v string) *ExportScanResultResponseBody {
	s.Data = &v
	return s
}

func (s *ExportScanResultResponseBody) SetHttpStatusCode(v int32) *ExportScanResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExportScanResultResponseBody) SetMsg(v string) *ExportScanResultResponseBody {
	s.Msg = &v
	return s
}

func (s *ExportScanResultResponseBody) SetRequestId(v string) *ExportScanResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportScanResultResponseBody) SetSuccess(v bool) *ExportScanResultResponseBody {
	s.Success = &v
	return s
}

type ExportScanResultResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportScanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportScanResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportScanResultResponse) GoString() string {
	return s.String()
}

func (s *ExportScanResultResponse) SetHeaders(v map[string]*string) *ExportScanResultResponse {
	s.Headers = v
	return s
}

func (s *ExportScanResultResponse) SetStatusCode(v int32) *ExportScanResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportScanResultResponse) SetBody(v *ExportScanResultResponseBody) *ExportScanResultResponse {
	s.Body = v
	return s
}

type ExportTextScanResultRequest struct {
	// example:
	//
	// 2024-03-11 10:00:00
	EndDate *string            `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Query   map[string]*string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2024-03-10 10:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportTextScanResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportTextScanResultRequest) GoString() string {
	return s.String()
}

func (s *ExportTextScanResultRequest) SetEndDate(v string) *ExportTextScanResultRequest {
	s.EndDate = &v
	return s
}

func (s *ExportTextScanResultRequest) SetQuery(v map[string]*string) *ExportTextScanResultRequest {
	s.Query = v
	return s
}

func (s *ExportTextScanResultRequest) SetRegionId(v string) *ExportTextScanResultRequest {
	s.RegionId = &v
	return s
}

func (s *ExportTextScanResultRequest) SetStartDate(v string) *ExportTextScanResultRequest {
	s.StartDate = &v
	return s
}

type ExportTextScanResultShrinkRequest struct {
	// example:
	//
	// 2024-03-11 10:00:00
	EndDate     *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	QueryShrink *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2024-03-10 10:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportTextScanResultShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportTextScanResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExportTextScanResultShrinkRequest) SetEndDate(v string) *ExportTextScanResultShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *ExportTextScanResultShrinkRequest) SetQueryShrink(v string) *ExportTextScanResultShrinkRequest {
	s.QueryShrink = &v
	return s
}

func (s *ExportTextScanResultShrinkRequest) SetRegionId(v string) *ExportTextScanResultShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ExportTextScanResultShrinkRequest) SetStartDate(v string) *ExportTextScanResultShrinkRequest {
	s.StartDate = &v
	return s
}

type ExportTextScanResultResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/portal_data/production/scanResult/text/textScanResult_aliow2MAdWXCakCxlitVY8Lnn-1A9KEw.xlsx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportTextScanResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportTextScanResultResponseBody) GoString() string {
	return s.String()
}

func (s *ExportTextScanResultResponseBody) SetCode(v int32) *ExportTextScanResultResponseBody {
	s.Code = &v
	return s
}

func (s *ExportTextScanResultResponseBody) SetData(v string) *ExportTextScanResultResponseBody {
	s.Data = &v
	return s
}

func (s *ExportTextScanResultResponseBody) SetMsg(v string) *ExportTextScanResultResponseBody {
	s.Msg = &v
	return s
}

func (s *ExportTextScanResultResponseBody) SetRequestId(v string) *ExportTextScanResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportTextScanResultResponseBody) SetSuccess(v bool) *ExportTextScanResultResponseBody {
	s.Success = &v
	return s
}

type ExportTextScanResultResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportTextScanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportTextScanResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportTextScanResultResponse) GoString() string {
	return s.String()
}

func (s *ExportTextScanResultResponse) SetHeaders(v map[string]*string) *ExportTextScanResultResponse {
	s.Headers = v
	return s
}

func (s *ExportTextScanResultResponse) SetStatusCode(v int32) *ExportTextScanResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportTextScanResultResponse) SetBody(v *ExportTextScanResultResponseBody) *ExportTextScanResultResponse {
	s.Body = v
	return s
}

type GetAnswerImportProgressRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// alAxbbxxxx-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAnswerImportProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAnswerImportProgressRequest) GoString() string {
	return s.String()
}

func (s *GetAnswerImportProgressRequest) SetRegionId(v string) *GetAnswerImportProgressRequest {
	s.RegionId = &v
	return s
}

func (s *GetAnswerImportProgressRequest) SetTaskId(v string) *GetAnswerImportProgressRequest {
	s.TaskId = &v
	return s
}

type GetAnswerImportProgressResponseBody struct {
	// example:
	//
	// xxx
	I18nKey              *string   `json:"I18nKey,omitempty" xml:"I18nKey,omitempty"`
	IllegalLengthSamples []*string `json:"IllegalLengthSamples,omitempty" xml:"IllegalLengthSamples,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	InvalidCount *int32 `json:"InvalidCount,omitempty" xml:"InvalidCount,omitempty"`
	// example:
	//
	// alxxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 1
	RepeatCount   *int32    `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	RepeatSamples []*string `json:"RepeatSamples,omitempty" xml:"RepeatSamples,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// alAxbbxxxx-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// xxxxx
	Tips *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAnswerImportProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAnswerImportProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetAnswerImportProgressResponseBody) SetI18nKey(v string) *GetAnswerImportProgressResponseBody {
	s.I18nKey = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetIllegalLengthSamples(v []*string) *GetAnswerImportProgressResponseBody {
	s.IllegalLengthSamples = v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetInvalidCount(v int32) *GetAnswerImportProgressResponseBody {
	s.InvalidCount = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetLibId(v string) *GetAnswerImportProgressResponseBody {
	s.LibId = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetProgress(v int32) *GetAnswerImportProgressResponseBody {
	s.Progress = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetRepeatCount(v int32) *GetAnswerImportProgressResponseBody {
	s.RepeatCount = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetRepeatSamples(v []*string) *GetAnswerImportProgressResponseBody {
	s.RepeatSamples = v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetRequestId(v string) *GetAnswerImportProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetSuccessCount(v int32) *GetAnswerImportProgressResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetTaskId(v string) *GetAnswerImportProgressResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetTips(v string) *GetAnswerImportProgressResponseBody {
	s.Tips = &v
	return s
}

func (s *GetAnswerImportProgressResponseBody) SetTotalCount(v int32) *GetAnswerImportProgressResponseBody {
	s.TotalCount = &v
	return s
}

type GetAnswerImportProgressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAnswerImportProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAnswerImportProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAnswerImportProgressResponse) GoString() string {
	return s.String()
}

func (s *GetAnswerImportProgressResponse) SetHeaders(v map[string]*string) *GetAnswerImportProgressResponse {
	s.Headers = v
	return s
}

func (s *GetAnswerImportProgressResponse) SetStatusCode(v int32) *GetAnswerImportProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAnswerImportProgressResponse) SetBody(v *GetAnswerImportProgressResponseBody) *GetAnswerImportProgressResponse {
	s.Body = v
	return s
}

type GetBackupBucketsListRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBackupBucketsListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBackupBucketsListRequest) GoString() string {
	return s.String()
}

func (s *GetBackupBucketsListRequest) SetRegionId(v string) *GetBackupBucketsListRequest {
	s.RegionId = &v
	return s
}

type GetBackupBucketsListResponseBody struct {
	Data []*GetBackupBucketsListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBackupBucketsListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBackupBucketsListResponseBody) GoString() string {
	return s.String()
}

func (s *GetBackupBucketsListResponseBody) SetData(v []*GetBackupBucketsListResponseBodyData) *GetBackupBucketsListResponseBody {
	s.Data = v
	return s
}

func (s *GetBackupBucketsListResponseBody) SetRequestId(v string) *GetBackupBucketsListResponseBody {
	s.RequestId = &v
	return s
}

type GetBackupBucketsListResponseBodyData struct {
	// example:
	//
	// gj-bucket1
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetBackupBucketsListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetBackupBucketsListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBackupBucketsListResponseBodyData) SetBucket(v string) *GetBackupBucketsListResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *GetBackupBucketsListResponseBodyData) SetRegion(v string) *GetBackupBucketsListResponseBodyData {
	s.Region = &v
	return s
}

type GetBackupBucketsListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBackupBucketsListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBackupBucketsListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBackupBucketsListResponse) GoString() string {
	return s.String()
}

func (s *GetBackupBucketsListResponse) SetHeaders(v map[string]*string) *GetBackupBucketsListResponse {
	s.Headers = v
	return s
}

func (s *GetBackupBucketsListResponse) SetStatusCode(v int32) *GetBackupBucketsListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBackupBucketsListResponse) SetBody(v *GetBackupBucketsListResponseBody) *GetBackupBucketsListResponse {
	s.Body = v
	return s
}

type GetBackupConfigRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s GetBackupConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBackupConfigRequest) GoString() string {
	return s.String()
}

func (s *GetBackupConfigRequest) SetRegionId(v string) *GetBackupConfigRequest {
	s.RegionId = &v
	return s
}

func (s *GetBackupConfigRequest) SetResourceType(v string) *GetBackupConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *GetBackupConfigRequest) SetServiceCode(v string) *GetBackupConfigRequest {
	s.ServiceCode = &v
	return s
}

type GetBackupConfigResponseBody struct {
	// example:
	//
	// 0
	BackupMode *int32 `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// example:
	//
	// buckect_test
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// true
	Enable            *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	EnableBackup      *bool `json:"EnableBackup,omitempty" xml:"EnableBackup,omitempty"`
	EnableBackupVoice *bool `json:"EnableBackupVoice,omitempty" xml:"EnableBackupVoice,omitempty"`
	// example:
	//
	// 300
	ExpireSeconds *int32 `json:"ExpireSeconds,omitempty" xml:"ExpireSeconds,omitempty"`
	// example:
	//
	// 2023-01-17 12:29:56
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// aliyun/template/
	Path      *string `json:"Path,omitempty" xml:"Path,omitempty"`
	PathVoice *string `json:"PathVoice,omitempty" xml:"PathVoice,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// UID。
	//
	// example:
	//
	// 1772612608370735
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetBackupConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBackupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetBackupConfigResponseBody) SetBackupMode(v int32) *GetBackupConfigResponseBody {
	s.BackupMode = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetBucket(v string) *GetBackupConfigResponseBody {
	s.Bucket = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetEnable(v bool) *GetBackupConfigResponseBody {
	s.Enable = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetEnableBackup(v bool) *GetBackupConfigResponseBody {
	s.EnableBackup = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetEnableBackupVoice(v bool) *GetBackupConfigResponseBody {
	s.EnableBackupVoice = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetExpireSeconds(v int32) *GetBackupConfigResponseBody {
	s.ExpireSeconds = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetGmtModified(v string) *GetBackupConfigResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetPath(v string) *GetBackupConfigResponseBody {
	s.Path = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetPathVoice(v string) *GetBackupConfigResponseBody {
	s.PathVoice = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetRegion(v string) *GetBackupConfigResponseBody {
	s.Region = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetRequestId(v string) *GetBackupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetResourceType(v string) *GetBackupConfigResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetServiceCode(v string) *GetBackupConfigResponseBody {
	s.ServiceCode = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetUid(v string) *GetBackupConfigResponseBody {
	s.Uid = &v
	return s
}

type GetBackupConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBackupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBackupConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBackupConfigResponse) GoString() string {
	return s.String()
}

func (s *GetBackupConfigResponse) SetHeaders(v map[string]*string) *GetBackupConfigResponse {
	s.Headers = v
	return s
}

func (s *GetBackupConfigResponse) SetStatusCode(v int32) *GetBackupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBackupConfigResponse) SetBody(v *GetBackupConfigResponseBody) *GetBackupConfigResponse {
	s.Body = v
	return s
}

type GetBackupStatusRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBackupStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBackupStatusRequest) GoString() string {
	return s.String()
}

func (s *GetBackupStatusRequest) SetRegionId(v string) *GetBackupStatusRequest {
	s.RegionId = &v
	return s
}

type GetBackupStatusResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBackupStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBackupStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetBackupStatusResponseBody) SetData(v bool) *GetBackupStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetBackupStatusResponseBody) SetRequestId(v string) *GetBackupStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetBackupStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBackupStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBackupStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBackupStatusResponse) GoString() string {
	return s.String()
}

func (s *GetBackupStatusResponse) SetHeaders(v map[string]*string) *GetBackupStatusResponse {
	s.Headers = v
	return s
}

func (s *GetBackupStatusResponse) SetStatusCode(v int32) *GetBackupStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBackupStatusResponse) SetBody(v *GetBackupStatusResponseBody) *GetBackupStatusResponse {
	s.Body = v
	return s
}

type GetBucketsListRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBucketsListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBucketsListRequest) GoString() string {
	return s.String()
}

func (s *GetBucketsListRequest) SetRegionId(v string) *GetBucketsListRequest {
	s.RegionId = &v
	return s
}

type GetBucketsListResponseBody struct {
	Data []*GetBucketsListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBucketsListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketsListResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketsListResponseBody) SetData(v []*GetBucketsListResponseBodyData) *GetBucketsListResponseBody {
	s.Data = v
	return s
}

func (s *GetBucketsListResponseBody) SetRequestId(v string) *GetBucketsListResponseBody {
	s.RequestId = &v
	return s
}

type GetBucketsListResponseBodyData struct {
	// example:
	//
	// bucket_test
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetBucketsListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetBucketsListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBucketsListResponseBodyData) SetBucket(v string) *GetBucketsListResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *GetBucketsListResponseBodyData) SetRegion(v string) *GetBucketsListResponseBodyData {
	s.Region = &v
	return s
}

type GetBucketsListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketsListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketsListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketsListResponse) GoString() string {
	return s.String()
}

func (s *GetBucketsListResponse) SetHeaders(v map[string]*string) *GetBucketsListResponse {
	s.Headers = v
	return s
}

func (s *GetBucketsListResponse) SetStatusCode(v int32) *GetBucketsListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketsListResponse) SetBody(v *GetBucketsListResponseBody) *GetBucketsListResponse {
	s.Body = v
	return s
}

type GetCipStatsRequest struct {
	// example:
	//
	// true
	ByMonth *bool `json:"ByMonth,omitempty" xml:"ByMonth,omitempty"`
	// example:
	//
	// 2024-03-11 10:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// xx
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ServiceCode  *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// 2024-03-10 10:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 253552244990701265
	SubUid *string `json:"SubUid,omitempty" xml:"SubUid,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCipStatsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCipStatsRequest) GoString() string {
	return s.String()
}

func (s *GetCipStatsRequest) SetByMonth(v bool) *GetCipStatsRequest {
	s.ByMonth = &v
	return s
}

func (s *GetCipStatsRequest) SetEndDate(v string) *GetCipStatsRequest {
	s.EndDate = &v
	return s
}

func (s *GetCipStatsRequest) SetLabel(v string) *GetCipStatsRequest {
	s.Label = &v
	return s
}

func (s *GetCipStatsRequest) SetRegionId(v string) *GetCipStatsRequest {
	s.RegionId = &v
	return s
}

func (s *GetCipStatsRequest) SetResourceType(v string) *GetCipStatsRequest {
	s.ResourceType = &v
	return s
}

func (s *GetCipStatsRequest) SetServiceCode(v string) *GetCipStatsRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetCipStatsRequest) SetStartDate(v string) *GetCipStatsRequest {
	s.StartDate = &v
	return s
}

func (s *GetCipStatsRequest) SetSubUid(v string) *GetCipStatsRequest {
	s.SubUid = &v
	return s
}

func (s *GetCipStatsRequest) SetType(v string) *GetCipStatsRequest {
	s.Type = &v
	return s
}

type GetCipStatsResponseBody struct {
	// example:
	//
	// 200
	Code *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetCipStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCipStatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCipStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBody) SetCode(v int32) *GetCipStatsResponseBody {
	s.Code = &v
	return s
}

func (s *GetCipStatsResponseBody) SetData(v *GetCipStatsResponseBodyData) *GetCipStatsResponseBody {
	s.Data = v
	return s
}

func (s *GetCipStatsResponseBody) SetHttpStatusCode(v int32) *GetCipStatsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCipStatsResponseBody) SetMsg(v string) *GetCipStatsResponseBody {
	s.Msg = &v
	return s
}

func (s *GetCipStatsResponseBody) SetRequestId(v string) *GetCipStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCipStatsResponseBody) SetSuccess(v bool) *GetCipStatsResponseBody {
	s.Success = &v
	return s
}

type GetCipStatsResponseBodyData struct {
	LabelStatChart []*GetCipStatsResponseBodyDataLabelStatChart `json:"LabelStatChart,omitempty" xml:"LabelStatChart,omitempty" type:"Repeated"`
	TotalStat      map[string]interface{}                       `json:"TotalStat,omitempty" xml:"TotalStat,omitempty"`
	Uids           []*string                                    `json:"Uids,omitempty" xml:"Uids,omitempty" type:"Repeated"`
	X              []*string                                    `json:"X,omitempty" xml:"X,omitempty" type:"Repeated"`
	Y              []*GetCipStatsResponseBodyDataY              `json:"Y,omitempty" xml:"Y,omitempty" type:"Repeated"`
	Z              []*GetCipStatsResponseBodyDataZ              `json:"Z,omitempty" xml:"Z,omitempty" type:"Repeated"`
}

func (s GetCipStatsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCipStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyData) SetLabelStatChart(v []*GetCipStatsResponseBodyDataLabelStatChart) *GetCipStatsResponseBodyData {
	s.LabelStatChart = v
	return s
}

func (s *GetCipStatsResponseBodyData) SetTotalStat(v map[string]interface{}) *GetCipStatsResponseBodyData {
	s.TotalStat = v
	return s
}

func (s *GetCipStatsResponseBodyData) SetUids(v []*string) *GetCipStatsResponseBodyData {
	s.Uids = v
	return s
}

func (s *GetCipStatsResponseBodyData) SetX(v []*string) *GetCipStatsResponseBodyData {
	s.X = v
	return s
}

func (s *GetCipStatsResponseBodyData) SetY(v []*GetCipStatsResponseBodyDataY) *GetCipStatsResponseBodyData {
	s.Y = v
	return s
}

func (s *GetCipStatsResponseBodyData) SetZ(v []*GetCipStatsResponseBodyDataZ) *GetCipStatsResponseBodyData {
	s.Z = v
	return s
}

type GetCipStatsResponseBodyDataLabelStatChart struct {
	ImageTreeChar []*GetCipStatsResponseBodyDataLabelStatChartImageTreeChar `json:"ImageTreeChar,omitempty" xml:"ImageTreeChar,omitempty" type:"Repeated"`
	// example:
	//
	// nickNameDetection
	ServiceCode   *string                                                   `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	TextTreeChart []*GetCipStatsResponseBodyDataLabelStatChartTextTreeChart `json:"TextTreeChart,omitempty" xml:"TextTreeChart,omitempty" type:"Repeated"`
	// example:
	//
	// 117
	TotalCount     *int64                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TreeChart      []*GetCipStatsResponseBodyDataLabelStatChartTreeChart      `json:"TreeChart,omitempty" xml:"TreeChart,omitempty" type:"Repeated"`
	VoiceTreeChart []*GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart `json:"VoiceTreeChart,omitempty" xml:"VoiceTreeChart,omitempty" type:"Repeated"`
	X              []*string                                                  `json:"X,omitempty" xml:"X,omitempty" type:"Repeated"`
	Y              []*GetCipStatsResponseBodyDataLabelStatChartY              `json:"Y,omitempty" xml:"Y,omitempty" type:"Repeated"`
}

func (s GetCipStatsResponseBodyDataLabelStatChart) String() string {
	return tea.Prettify(s)
}

func (s GetCipStatsResponseBodyDataLabelStatChart) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) SetImageTreeChar(v []*GetCipStatsResponseBodyDataLabelStatChartImageTreeChar) *GetCipStatsResponseBodyDataLabelStatChart {
	s.ImageTreeChar = v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) SetServiceCode(v string) *GetCipStatsResponseBodyDataLabelStatChart {
	s.ServiceCode = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) SetTextTreeChart(v []*GetCipStatsResponseBodyDataLabelStatChartTextTreeChart) *GetCipStatsResponseBodyDataLabelStatChart {
	s.TextTreeChart = v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) SetTotalCount(v int64) *GetCipStatsResponseBodyDataLabelStatChart {
	s.TotalCount = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) SetTreeChart(v []*GetCipStatsResponseBodyDataLabelStatChartTreeChart) *GetCipStatsResponseBodyDataLabelStatChart {
	s.TreeChart = v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) SetVoiceTreeChart(v []*GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart) *GetCipStatsResponseBodyDataLabelStatChart {
	s.VoiceTreeChart = v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) SetX(v []*string) *GetCipStatsResponseBodyDataLabelStatChart {
	s.X = v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChart) SetY(v []*GetCipStatsResponseBodyDataLabelStatChartY) *GetCipStatsResponseBodyDataLabelStatChart {
	s.Y = v
	return s
}

type GetCipStatsResponseBodyDataLabelStatChartImageTreeChar struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCipStatsResponseBodyDataLabelStatChartImageTreeChar) String() string {
	return tea.Prettify(s)
}

func (s GetCipStatsResponseBodyDataLabelStatChartImageTreeChar) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyDataLabelStatChartImageTreeChar) SetDescription(v string) *GetCipStatsResponseBodyDataLabelStatChartImageTreeChar {
	s.Description = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartImageTreeChar) SetName(v string) *GetCipStatsResponseBodyDataLabelStatChartImageTreeChar {
	s.Name = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartImageTreeChar) SetValue(v string) *GetCipStatsResponseBodyDataLabelStatChartImageTreeChar {
	s.Value = &v
	return s
}

type GetCipStatsResponseBodyDataLabelStatChartTextTreeChart struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCipStatsResponseBodyDataLabelStatChartTextTreeChart) String() string {
	return tea.Prettify(s)
}

func (s GetCipStatsResponseBodyDataLabelStatChartTextTreeChart) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTextTreeChart) SetDescription(v string) *GetCipStatsResponseBodyDataLabelStatChartTextTreeChart {
	s.Description = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTextTreeChart) SetName(v string) *GetCipStatsResponseBodyDataLabelStatChartTextTreeChart {
	s.Name = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTextTreeChart) SetValue(v string) *GetCipStatsResponseBodyDataLabelStatChartTextTreeChart {
	s.Value = &v
	return s
}

type GetCipStatsResponseBodyDataLabelStatChartTreeChart struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// nickNameDetection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 99.91
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCipStatsResponseBodyDataLabelStatChartTreeChart) String() string {
	return tea.Prettify(s)
}

func (s GetCipStatsResponseBodyDataLabelStatChartTreeChart) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTreeChart) SetDescription(v string) *GetCipStatsResponseBodyDataLabelStatChartTreeChart {
	s.Description = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTreeChart) SetName(v string) *GetCipStatsResponseBodyDataLabelStatChartTreeChart {
	s.Name = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartTreeChart) SetValue(v string) *GetCipStatsResponseBodyDataLabelStatChartTreeChart {
	s.Value = &v
	return s
}

type GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// nickNameDetection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 99.91
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart) String() string {
	return tea.Prettify(s)
}

func (s GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart) SetDescription(v string) *GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart {
	s.Description = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart) SetName(v string) *GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart {
	s.Name = &v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart) SetValue(v string) *GetCipStatsResponseBodyDataLabelStatChartVoiceTreeChart {
	s.Value = &v
	return s
}

type GetCipStatsResponseBodyDataLabelStatChartY struct {
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// nickNameDetection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetCipStatsResponseBodyDataLabelStatChartY) String() string {
	return tea.Prettify(s)
}

func (s GetCipStatsResponseBodyDataLabelStatChartY) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyDataLabelStatChartY) SetData(v []*int64) *GetCipStatsResponseBodyDataLabelStatChartY {
	s.Data = v
	return s
}

func (s *GetCipStatsResponseBodyDataLabelStatChartY) SetName(v string) *GetCipStatsResponseBodyDataLabelStatChartY {
	s.Name = &v
	return s
}

type GetCipStatsResponseBodyDataY struct {
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// nickNameDetection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetCipStatsResponseBodyDataY) String() string {
	return tea.Prettify(s)
}

func (s GetCipStatsResponseBodyDataY) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyDataY) SetData(v []*int64) *GetCipStatsResponseBodyDataY {
	s.Data = v
	return s
}

func (s *GetCipStatsResponseBodyDataY) SetName(v string) *GetCipStatsResponseBodyDataY {
	s.Name = &v
	return s
}

type GetCipStatsResponseBodyDataZ struct {
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// nickNameDetection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetCipStatsResponseBodyDataZ) String() string {
	return tea.Prettify(s)
}

func (s GetCipStatsResponseBodyDataZ) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponseBodyDataZ) SetData(v []*int64) *GetCipStatsResponseBodyDataZ {
	s.Data = v
	return s
}

func (s *GetCipStatsResponseBodyDataZ) SetName(v string) *GetCipStatsResponseBodyDataZ {
	s.Name = &v
	return s
}

type GetCipStatsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCipStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCipStatsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCipStatsResponse) GoString() string {
	return s.String()
}

func (s *GetCipStatsResponse) SetHeaders(v map[string]*string) *GetCipStatsResponse {
	s.Headers = v
	return s
}

func (s *GetCipStatsResponse) SetStatusCode(v int32) *GetCipStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCipStatsResponse) SetBody(v *GetCipStatsResponseBody) *GetCipStatsResponse {
	s.Body = v
	return s
}

type GetExecuteTimeRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetExecuteTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExecuteTimeRequest) GoString() string {
	return s.String()
}

func (s *GetExecuteTimeRequest) SetRegionId(v string) *GetExecuteTimeRequest {
	s.RegionId = &v
	return s
}

type GetExecuteTimeResponseBody struct {
	// example:
	//
	// 02:24:30
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExecuteTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExecuteTimeResponseBody) GoString() string {
	return s.String()
}

func (s *GetExecuteTimeResponseBody) SetData(v string) *GetExecuteTimeResponseBody {
	s.Data = &v
	return s
}

func (s *GetExecuteTimeResponseBody) SetRequestId(v string) *GetExecuteTimeResponseBody {
	s.RequestId = &v
	return s
}

type GetExecuteTimeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExecuteTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExecuteTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExecuteTimeResponse) GoString() string {
	return s.String()
}

func (s *GetExecuteTimeResponse) SetHeaders(v map[string]*string) *GetExecuteTimeResponse {
	s.Headers = v
	return s
}

func (s *GetExecuteTimeResponse) SetStatusCode(v int32) *GetExecuteTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExecuteTimeResponse) SetBody(v *GetExecuteTimeResponseBody) *GetExecuteTimeResponse {
	s.Body = v
	return s
}

type GetFeatureConfigRequest struct {
	// Region ID
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource type.
	//
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Service code.
	//
	// example:
	//
	// llm_query_moderation
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Type
	//
	// example:
	//
	// custom_llm_template
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetFeatureConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureConfigRequest) GoString() string {
	return s.String()
}

func (s *GetFeatureConfigRequest) SetRegionId(v string) *GetFeatureConfigRequest {
	s.RegionId = &v
	return s
}

func (s *GetFeatureConfigRequest) SetResourceType(v string) *GetFeatureConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *GetFeatureConfigRequest) SetServiceCode(v string) *GetFeatureConfigRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetFeatureConfigRequest) SetType(v string) *GetFeatureConfigRequest {
	s.Type = &v
	return s
}

type GetFeatureConfigResponseBody struct {
	// Status code
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	Data *GetFeatureConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message for this request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// ID assigned by the backend, used to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success indicator
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFeatureConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetFeatureConfigResponseBody) SetCode(v int32) *GetFeatureConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetFeatureConfigResponseBody) SetData(v *GetFeatureConfigResponseBodyData) *GetFeatureConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetFeatureConfigResponseBody) SetHttpStatusCode(v int32) *GetFeatureConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetFeatureConfigResponseBody) SetMsg(v string) *GetFeatureConfigResponseBody {
	s.Msg = &v
	return s
}

func (s *GetFeatureConfigResponseBody) SetRequestId(v string) *GetFeatureConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFeatureConfigResponseBody) SetSuccess(v bool) *GetFeatureConfigResponseBody {
	s.Success = &v
	return s
}

type GetFeatureConfigResponseBodyData struct {
	// List of feature configurations
	FeatureConf []map[string]interface{} `json:"FeatureConf,omitempty" xml:"FeatureConf,omitempty" type:"Repeated"`
	// Resource type.
	//
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Service code.
	//
	// example:
	//
	// llm_query_moderation
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Type
	//
	// example:
	//
	// custom_llm_template
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// UID.
	//
	// example:
	//
	// 1643953****74290
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetFeatureConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFeatureConfigResponseBodyData) SetFeatureConf(v []map[string]interface{}) *GetFeatureConfigResponseBodyData {
	s.FeatureConf = v
	return s
}

func (s *GetFeatureConfigResponseBodyData) SetResourceType(v string) *GetFeatureConfigResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *GetFeatureConfigResponseBodyData) SetServiceCode(v string) *GetFeatureConfigResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *GetFeatureConfigResponseBodyData) SetType(v string) *GetFeatureConfigResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetFeatureConfigResponseBodyData) SetUid(v string) *GetFeatureConfigResponseBodyData {
	s.Uid = &v
	return s
}

type GetFeatureConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFeatureConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFeatureConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureConfigResponse) GoString() string {
	return s.String()
}

func (s *GetFeatureConfigResponse) SetHeaders(v map[string]*string) *GetFeatureConfigResponse {
	s.Headers = v
	return s
}

func (s *GetFeatureConfigResponse) SetStatusCode(v int32) *GetFeatureConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFeatureConfigResponse) SetBody(v *GetFeatureConfigResponseBody) *GetFeatureConfigResponse {
	s.Body = v
	return s
}

type GetImageSceneLabelConfRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetImageSceneLabelConfRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageSceneLabelConfRequest) GoString() string {
	return s.String()
}

func (s *GetImageSceneLabelConfRequest) SetRegionId(v string) *GetImageSceneLabelConfRequest {
	s.RegionId = &v
	return s
}

type GetImageSceneLabelConfResponseBody struct {
	// example:
	//
	// 200
	Code *int32                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetImageSceneLabelConfResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageSceneLabelConfResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageSceneLabelConfResponseBody) SetCode(v int32) *GetImageSceneLabelConfResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageSceneLabelConfResponseBody) SetData(v []map[string]interface{}) *GetImageSceneLabelConfResponseBody {
	s.Data = v
	return s
}

func (s *GetImageSceneLabelConfResponseBody) SetHttpStatusCode(v int32) *GetImageSceneLabelConfResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetImageSceneLabelConfResponseBody) SetMsg(v string) *GetImageSceneLabelConfResponseBody {
	s.Msg = &v
	return s
}

func (s *GetImageSceneLabelConfResponseBody) SetRequestId(v string) *GetImageSceneLabelConfResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageSceneLabelConfResponseBody) SetSuccess(v bool) *GetImageSceneLabelConfResponseBody {
	s.Success = &v
	return s
}

type GetImageSceneLabelConfResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageSceneLabelConfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageSceneLabelConfResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageSceneLabelConfResponse) GoString() string {
	return s.String()
}

func (s *GetImageSceneLabelConfResponse) SetHeaders(v map[string]*string) *GetImageSceneLabelConfResponse {
	s.Headers = v
	return s
}

func (s *GetImageSceneLabelConfResponse) SetStatusCode(v int32) *GetImageSceneLabelConfResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageSceneLabelConfResponse) SetBody(v *GetImageSceneLabelConfResponseBody) *GetImageSceneLabelConfResponse {
	s.Body = v
	return s
}

type GetImageSceneLabelListConfRequest struct {
	// example:
	//
	// baselineCheck
	ImageServiceCode *string `json:"ImageServiceCode,omitempty" xml:"ImageServiceCode,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetImageSceneLabelListConfRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageSceneLabelListConfRequest) GoString() string {
	return s.String()
}

func (s *GetImageSceneLabelListConfRequest) SetImageServiceCode(v string) *GetImageSceneLabelListConfRequest {
	s.ImageServiceCode = &v
	return s
}

func (s *GetImageSceneLabelListConfRequest) SetRegionId(v string) *GetImageSceneLabelListConfRequest {
	s.RegionId = &v
	return s
}

type GetImageSceneLabelListConfResponseBody struct {
	Data []interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageSceneLabelListConfResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageSceneLabelListConfResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageSceneLabelListConfResponseBody) SetData(v []interface{}) *GetImageSceneLabelListConfResponseBody {
	s.Data = v
	return s
}

func (s *GetImageSceneLabelListConfResponseBody) SetRequestId(v string) *GetImageSceneLabelListConfResponseBody {
	s.RequestId = &v
	return s
}

type GetImageSceneLabelListConfResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageSceneLabelListConfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageSceneLabelListConfResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageSceneLabelListConfResponse) GoString() string {
	return s.String()
}

func (s *GetImageSceneLabelListConfResponse) SetHeaders(v map[string]*string) *GetImageSceneLabelListConfResponse {
	s.Headers = v
	return s
}

func (s *GetImageSceneLabelListConfResponse) SetStatusCode(v int32) *GetImageSceneLabelListConfResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageSceneLabelListConfResponse) SetBody(v *GetImageSceneLabelListConfResponseBody) *GetImageSceneLabelListConfResponse {
	s.Body = v
	return s
}

type GetJobNameListRequest struct {
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// {"TaskId":"P_11TL5T"}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sort     map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetJobNameListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobNameListRequest) GoString() string {
	return s.String()
}

func (s *GetJobNameListRequest) SetEndDate(v string) *GetJobNameListRequest {
	s.EndDate = &v
	return s
}

func (s *GetJobNameListRequest) SetQuery(v string) *GetJobNameListRequest {
	s.Query = &v
	return s
}

func (s *GetJobNameListRequest) SetRegionId(v string) *GetJobNameListRequest {
	s.RegionId = &v
	return s
}

func (s *GetJobNameListRequest) SetSort(v map[string]*string) *GetJobNameListRequest {
	s.Sort = v
	return s
}

func (s *GetJobNameListRequest) SetStartDate(v string) *GetJobNameListRequest {
	s.StartDate = &v
	return s
}

type GetJobNameListShrinkRequest struct {
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// {"TaskId":"P_11TL5T"}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetJobNameListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobNameListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetJobNameListShrinkRequest) SetEndDate(v string) *GetJobNameListShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetJobNameListShrinkRequest) SetQuery(v string) *GetJobNameListShrinkRequest {
	s.Query = &v
	return s
}

func (s *GetJobNameListShrinkRequest) SetRegionId(v string) *GetJobNameListShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetJobNameListShrinkRequest) SetSortShrink(v string) *GetJobNameListShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *GetJobNameListShrinkRequest) SetStartDate(v string) *GetJobNameListShrinkRequest {
	s.StartDate = &v
	return s
}

type GetJobNameListResponseBody struct {
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobNameListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobNameListResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobNameListResponseBody) SetData(v []*string) *GetJobNameListResponseBody {
	s.Data = v
	return s
}

func (s *GetJobNameListResponseBody) SetRequestId(v string) *GetJobNameListResponseBody {
	s.RequestId = &v
	return s
}

type GetJobNameListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobNameListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobNameListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobNameListResponse) GoString() string {
	return s.String()
}

func (s *GetJobNameListResponse) SetHeaders(v map[string]*string) *GetJobNameListResponse {
	s.Headers = v
	return s
}

func (s *GetJobNameListResponse) SetStatusCode(v int32) *GetJobNameListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobNameListResponse) SetBody(v *GetJobNameListResponseBody) *GetJobNameListResponse {
	s.Body = v
	return s
}

type GetKeywordImportResultRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// xxx-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetKeywordImportResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetKeywordImportResultRequest) GoString() string {
	return s.String()
}

func (s *GetKeywordImportResultRequest) SetRegionId(v string) *GetKeywordImportResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetKeywordImportResultRequest) SetTaskId(v string) *GetKeywordImportResultRequest {
	s.TaskId = &v
	return s
}

type GetKeywordImportResultResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetKeywordImportResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetKeywordImportResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetKeywordImportResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetKeywordImportResultResponseBody) SetCode(v int32) *GetKeywordImportResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetKeywordImportResultResponseBody) SetData(v *GetKeywordImportResultResponseBodyData) *GetKeywordImportResultResponseBody {
	s.Data = v
	return s
}

func (s *GetKeywordImportResultResponseBody) SetMsg(v string) *GetKeywordImportResultResponseBody {
	s.Msg = &v
	return s
}

func (s *GetKeywordImportResultResponseBody) SetRequestId(v string) *GetKeywordImportResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKeywordImportResultResponseBody) SetSuccess(v bool) *GetKeywordImportResultResponseBody {
	s.Success = &v
	return s
}

type GetKeywordImportResultResponseBodyData struct {
	// example:
	//
	// xxx
	I18nKey               *string   `json:"I18nKey,omitempty" xml:"I18nKey,omitempty"`
	IllegalLengthKeywords []*string `json:"IllegalLengthKeywords,omitempty" xml:"IllegalLengthKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	InvalidCount    *int32    `json:"InvalidCount,omitempty" xml:"InvalidCount,omitempty"`
	InvalidKeywords []*string `json:"InvalidKeywords,omitempty" xml:"InvalidKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// customxx_xxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 1
	RepeatCount    *int32    `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	RepeatKeywords []*string `json:"RepeatKeywords,omitempty" xml:"RepeatKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// 8
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// xxx
	Tips *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetKeywordImportResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetKeywordImportResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKeywordImportResultResponseBodyData) SetI18nKey(v string) *GetKeywordImportResultResponseBodyData {
	s.I18nKey = &v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetIllegalLengthKeywords(v []*string) *GetKeywordImportResultResponseBodyData {
	s.IllegalLengthKeywords = v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetInvalidCount(v int32) *GetKeywordImportResultResponseBodyData {
	s.InvalidCount = &v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetInvalidKeywords(v []*string) *GetKeywordImportResultResponseBodyData {
	s.InvalidKeywords = v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetLibId(v string) *GetKeywordImportResultResponseBodyData {
	s.LibId = &v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetProgress(v int32) *GetKeywordImportResultResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetRepeatCount(v int32) *GetKeywordImportResultResponseBodyData {
	s.RepeatCount = &v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetRepeatKeywords(v []*string) *GetKeywordImportResultResponseBodyData {
	s.RepeatKeywords = v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetSuccessCount(v int32) *GetKeywordImportResultResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetTips(v string) *GetKeywordImportResultResponseBodyData {
	s.Tips = &v
	return s
}

func (s *GetKeywordImportResultResponseBodyData) SetTotalCount(v int32) *GetKeywordImportResultResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetKeywordImportResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKeywordImportResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKeywordImportResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetKeywordImportResultResponse) GoString() string {
	return s.String()
}

func (s *GetKeywordImportResultResponse) SetHeaders(v map[string]*string) *GetKeywordImportResultResponse {
	s.Headers = v
	return s
}

func (s *GetKeywordImportResultResponse) SetStatusCode(v int32) *GetKeywordImportResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKeywordImportResultResponse) SetBody(v *GetKeywordImportResultResponseBody) *GetKeywordImportResultResponse {
	s.Body = v
	return s
}

type GetKeywordLibRequest struct {
	// example:
	//
	// customxx_xxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetKeywordLibRequest) String() string {
	return tea.Prettify(s)
}

func (s GetKeywordLibRequest) GoString() string {
	return s.String()
}

func (s *GetKeywordLibRequest) SetLibId(v string) *GetKeywordLibRequest {
	s.LibId = &v
	return s
}

func (s *GetKeywordLibRequest) SetRegionId(v string) *GetKeywordLibRequest {
	s.RegionId = &v
	return s
}

type GetKeywordLibResponseBody struct {
	// example:
	//
	// 200
	Code *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetKeywordLibResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetKeywordLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetKeywordLibResponseBody) GoString() string {
	return s.String()
}

func (s *GetKeywordLibResponseBody) SetCode(v int32) *GetKeywordLibResponseBody {
	s.Code = &v
	return s
}

func (s *GetKeywordLibResponseBody) SetData(v *GetKeywordLibResponseBodyData) *GetKeywordLibResponseBody {
	s.Data = v
	return s
}

func (s *GetKeywordLibResponseBody) SetMsg(v string) *GetKeywordLibResponseBody {
	s.Msg = &v
	return s
}

func (s *GetKeywordLibResponseBody) SetRequestId(v string) *GetKeywordLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKeywordLibResponseBody) SetSuccess(v bool) *GetKeywordLibResponseBody {
	s.Success = &v
	return s
}

type GetKeywordLibResponseBodyData struct {
	// example:
	//
	// 2024-01-29 10:26:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 100
	KeywordCount *string `json:"KeywordCount,omitempty" xml:"KeywordCount,omitempty"`
	// example:
	//
	// customxx_xxx
	LibId   *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// example:
	//
	// 1825457112123838
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetKeywordLibResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetKeywordLibResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKeywordLibResponseBodyData) SetGmtModified(v string) *GetKeywordLibResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetKeywordLibResponseBodyData) SetKeywordCount(v string) *GetKeywordLibResponseBodyData {
	s.KeywordCount = &v
	return s
}

func (s *GetKeywordLibResponseBodyData) SetLibId(v string) *GetKeywordLibResponseBodyData {
	s.LibId = &v
	return s
}

func (s *GetKeywordLibResponseBodyData) SetLibName(v string) *GetKeywordLibResponseBodyData {
	s.LibName = &v
	return s
}

func (s *GetKeywordLibResponseBodyData) SetUid(v string) *GetKeywordLibResponseBodyData {
	s.Uid = &v
	return s
}

type GetKeywordLibResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKeywordLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKeywordLibResponse) String() string {
	return tea.Prettify(s)
}

func (s GetKeywordLibResponse) GoString() string {
	return s.String()
}

func (s *GetKeywordLibResponse) SetHeaders(v map[string]*string) *GetKeywordLibResponse {
	s.Headers = v
	return s
}

func (s *GetKeywordLibResponse) SetStatusCode(v int32) *GetKeywordLibResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKeywordLibResponse) SetBody(v *GetKeywordLibResponseBody) *GetKeywordLibResponse {
	s.Body = v
	return s
}

type GetOssCheckStatRequest struct {
	// example:
	//
	// true
	ByMonth *bool `json:"ByMonth,omitempty" xml:"ByMonth,omitempty"`
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// P_UNHBH
	ParentTaskId *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetOssCheckStatRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOssCheckStatRequest) GoString() string {
	return s.String()
}

func (s *GetOssCheckStatRequest) SetByMonth(v bool) *GetOssCheckStatRequest {
	s.ByMonth = &v
	return s
}

func (s *GetOssCheckStatRequest) SetEndDate(v string) *GetOssCheckStatRequest {
	s.EndDate = &v
	return s
}

func (s *GetOssCheckStatRequest) SetParentTaskId(v string) *GetOssCheckStatRequest {
	s.ParentTaskId = &v
	return s
}

func (s *GetOssCheckStatRequest) SetRegionId(v string) *GetOssCheckStatRequest {
	s.RegionId = &v
	return s
}

func (s *GetOssCheckStatRequest) SetStartDate(v string) *GetOssCheckStatRequest {
	s.StartDate = &v
	return s
}

type GetOssCheckStatResponseBody struct {
	BarChart *GetOssCheckStatResponseBodyBarChart `json:"BarChart,omitempty" xml:"BarChart,omitempty" type:"Struct"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOssCheckStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOssCheckStatResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssCheckStatResponseBody) SetBarChart(v *GetOssCheckStatResponseBodyBarChart) *GetOssCheckStatResponseBody {
	s.BarChart = v
	return s
}

func (s *GetOssCheckStatResponseBody) SetRequestId(v string) *GetOssCheckStatResponseBody {
	s.RequestId = &v
	return s
}

type GetOssCheckStatResponseBodyBarChart struct {
	X []*string                               `json:"X,omitempty" xml:"X,omitempty" type:"Repeated"`
	Y []*GetOssCheckStatResponseBodyBarChartY `json:"Y,omitempty" xml:"Y,omitempty" type:"Repeated"`
}

func (s GetOssCheckStatResponseBodyBarChart) String() string {
	return tea.Prettify(s)
}

func (s GetOssCheckStatResponseBodyBarChart) GoString() string {
	return s.String()
}

func (s *GetOssCheckStatResponseBodyBarChart) SetX(v []*string) *GetOssCheckStatResponseBodyBarChart {
	s.X = v
	return s
}

func (s *GetOssCheckStatResponseBodyBarChart) SetY(v []*GetOssCheckStatResponseBodyBarChartY) *GetOssCheckStatResponseBodyBarChart {
	s.Y = v
	return s
}

type GetOssCheckStatResponseBodyBarChartY struct {
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// document_detection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOssCheckStatResponseBodyBarChartY) String() string {
	return tea.Prettify(s)
}

func (s GetOssCheckStatResponseBodyBarChartY) GoString() string {
	return s.String()
}

func (s *GetOssCheckStatResponseBodyBarChartY) SetData(v []*int64) *GetOssCheckStatResponseBodyBarChartY {
	s.Data = v
	return s
}

func (s *GetOssCheckStatResponseBodyBarChartY) SetName(v string) *GetOssCheckStatResponseBodyBarChartY {
	s.Name = &v
	return s
}

type GetOssCheckStatResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssCheckStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssCheckStatResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOssCheckStatResponse) GoString() string {
	return s.String()
}

func (s *GetOssCheckStatResponse) SetHeaders(v map[string]*string) *GetOssCheckStatResponse {
	s.Headers = v
	return s
}

func (s *GetOssCheckStatResponse) SetStatusCode(v int32) *GetOssCheckStatResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssCheckStatResponse) SetBody(v *GetOssCheckStatResponseBody) *GetOssCheckStatResponse {
	s.Body = v
	return s
}

type GetOssCheckStatusRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetOssCheckStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOssCheckStatusRequest) GoString() string {
	return s.String()
}

func (s *GetOssCheckStatusRequest) SetRegionId(v string) *GetOssCheckStatusRequest {
	s.RegionId = &v
	return s
}

type GetOssCheckStatusResponseBody struct {
	// Bid。
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// example:
	//
	// True
	Buy *bool `json:"Buy,omitempty" xml:"Buy,omitempty"`
	// example:
	//
	// xxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// False
	Indebt *bool `json:"Indebt,omitempty" xml:"Indebt,omitempty"`
	// example:
	//
	// True
	RamStatus *string `json:"RamStatus,omitempty" xml:"RamStatus,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	SlsStatus *string `json:"SlsStatus,omitempty" xml:"SlsStatus,omitempty"`
}

func (s GetOssCheckStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOssCheckStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssCheckStatusResponseBody) SetBid(v string) *GetOssCheckStatusResponseBody {
	s.Bid = &v
	return s
}

func (s *GetOssCheckStatusResponseBody) SetBuy(v bool) *GetOssCheckStatusResponseBody {
	s.Buy = &v
	return s
}

func (s *GetOssCheckStatusResponseBody) SetCommodityCode(v string) *GetOssCheckStatusResponseBody {
	s.CommodityCode = &v
	return s
}

func (s *GetOssCheckStatusResponseBody) SetIndebt(v bool) *GetOssCheckStatusResponseBody {
	s.Indebt = &v
	return s
}

func (s *GetOssCheckStatusResponseBody) SetRamStatus(v string) *GetOssCheckStatusResponseBody {
	s.RamStatus = &v
	return s
}

func (s *GetOssCheckStatusResponseBody) SetRequestId(v string) *GetOssCheckStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssCheckStatusResponseBody) SetSlsStatus(v string) *GetOssCheckStatusResponseBody {
	s.SlsStatus = &v
	return s
}

type GetOssCheckStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOssCheckStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOssCheckStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOssCheckStatusResponse) GoString() string {
	return s.String()
}

func (s *GetOssCheckStatusResponse) SetHeaders(v map[string]*string) *GetOssCheckStatusResponse {
	s.Headers = v
	return s
}

func (s *GetOssCheckStatusResponse) SetStatusCode(v int32) *GetOssCheckStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOssCheckStatusResponse) SetBody(v *GetOssCheckStatusResponseBody) *GetOssCheckStatusResponse {
	s.Body = v
	return s
}

type GetScanNumRequest struct {
	// example:
	//
	// tmpsample
	Buckets *string `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	// example:
	//
	// image
	MediaType *int32 `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetScanNumRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScanNumRequest) GoString() string {
	return s.String()
}

func (s *GetScanNumRequest) SetBuckets(v string) *GetScanNumRequest {
	s.Buckets = &v
	return s
}

func (s *GetScanNumRequest) SetMediaType(v int32) *GetScanNumRequest {
	s.MediaType = &v
	return s
}

func (s *GetScanNumRequest) SetRegionId(v string) *GetScanNumRequest {
	s.RegionId = &v
	return s
}

type GetScanNumResponseBody struct {
	// example:
	//
	// 10
	LimitNumber *int64 `json:"LimitNumber,omitempty" xml:"LimitNumber,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	ScanNumber *int64 `json:"ScanNumber,omitempty" xml:"ScanNumber,omitempty"`
	// example:
	//
	// 10
	SumNumber *int64 `json:"SumNumber,omitempty" xml:"SumNumber,omitempty"`
	// example:
	//
	// True
	Tag *bool `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetScanNumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScanNumResponseBody) GoString() string {
	return s.String()
}

func (s *GetScanNumResponseBody) SetLimitNumber(v int64) *GetScanNumResponseBody {
	s.LimitNumber = &v
	return s
}

func (s *GetScanNumResponseBody) SetRequestId(v string) *GetScanNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScanNumResponseBody) SetScanNumber(v int64) *GetScanNumResponseBody {
	s.ScanNumber = &v
	return s
}

func (s *GetScanNumResponseBody) SetSumNumber(v int64) *GetScanNumResponseBody {
	s.SumNumber = &v
	return s
}

func (s *GetScanNumResponseBody) SetTag(v bool) *GetScanNumResponseBody {
	s.Tag = &v
	return s
}

type GetScanNumResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScanNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScanNumResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScanNumResponse) GoString() string {
	return s.String()
}

func (s *GetScanNumResponse) SetHeaders(v map[string]*string) *GetScanNumResponse {
	s.Headers = v
	return s
}

func (s *GetScanNumResponse) SetStatusCode(v int32) *GetScanNumResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScanNumResponse) SetBody(v *GetScanNumResponseBody) *GetScanNumResponse {
	s.Body = v
	return s
}

type GetScanResultRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 10
	PageSize *int32             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query    map[string]*string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// image
	ResourceType *string            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Sort         map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetScanResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScanResultRequest) GoString() string {
	return s.String()
}

func (s *GetScanResultRequest) SetCurrentPage(v int32) *GetScanResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetScanResultRequest) SetEndDate(v string) *GetScanResultRequest {
	s.EndDate = &v
	return s
}

func (s *GetScanResultRequest) SetPageSize(v int32) *GetScanResultRequest {
	s.PageSize = &v
	return s
}

func (s *GetScanResultRequest) SetQuery(v map[string]*string) *GetScanResultRequest {
	s.Query = v
	return s
}

func (s *GetScanResultRequest) SetRegionId(v string) *GetScanResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetScanResultRequest) SetResourceType(v string) *GetScanResultRequest {
	s.ResourceType = &v
	return s
}

func (s *GetScanResultRequest) SetSort(v map[string]*string) *GetScanResultRequest {
	s.Sort = v
	return s
}

func (s *GetScanResultRequest) SetStartDate(v string) *GetScanResultRequest {
	s.StartDate = &v
	return s
}

type GetScanResultShrinkRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 10
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryShrink *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SortShrink   *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetScanResultShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScanResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetScanResultShrinkRequest) SetCurrentPage(v int32) *GetScanResultShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetScanResultShrinkRequest) SetEndDate(v string) *GetScanResultShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetScanResultShrinkRequest) SetPageSize(v int32) *GetScanResultShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetScanResultShrinkRequest) SetQueryShrink(v string) *GetScanResultShrinkRequest {
	s.QueryShrink = &v
	return s
}

func (s *GetScanResultShrinkRequest) SetRegionId(v string) *GetScanResultShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetScanResultShrinkRequest) SetResourceType(v string) *GetScanResultShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *GetScanResultShrinkRequest) SetSortShrink(v string) *GetScanResultShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *GetScanResultShrinkRequest) SetStartDate(v string) *GetScanResultShrinkRequest {
	s.StartDate = &v
	return s
}

type GetScanResultResponseBody struct {
	// example:
	//
	// 200
	Code *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetScanResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetScanResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScanResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetScanResultResponseBody) SetCode(v int32) *GetScanResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetScanResultResponseBody) SetData(v *GetScanResultResponseBodyData) *GetScanResultResponseBody {
	s.Data = v
	return s
}

func (s *GetScanResultResponseBody) SetHttpStatusCode(v int32) *GetScanResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetScanResultResponseBody) SetMsg(v string) *GetScanResultResponseBody {
	s.Msg = &v
	return s
}

func (s *GetScanResultResponseBody) SetRequestId(v string) *GetScanResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScanResultResponseBody) SetSuccess(v bool) *GetScanResultResponseBody {
	s.Success = &v
	return s
}

type GetScanResultResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int32                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*GetScanResultResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetScanResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetScanResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScanResultResponseBodyData) SetCurrentPage(v int32) *GetScanResultResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetScanResultResponseBodyData) SetItems(v []*GetScanResultResponseBodyDataItems) *GetScanResultResponseBodyData {
	s.Items = v
	return s
}

func (s *GetScanResultResponseBodyData) SetPageSize(v int32) *GetScanResultResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetScanResultResponseBodyData) SetTotalCount(v int64) *GetScanResultResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetScanResultResponseBodyDataItems struct {
	AttackLevel *string `json:"AttackLevel,omitempty" xml:"AttackLevel,omitempty"`
	// example:
	//
	// xxx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// dataId
	//
	// example:
	//
	// 4f27b8cc7c4544cb90b41882a5b36326
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// example:
	//
	// 22
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// xxx
	ExtFeedback *string `json:"ExtFeedback,omitempty" xml:"ExtFeedback,omitempty"`
	// example:
	//
	// {}
	Extra map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// 20
	FrameCount *int64 `json:"FrameCount,omitempty" xml:"FrameCount,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	GmtCreate   *string                  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ImageLabels []map[string]interface{} `json:"ImageLabels,omitempty" xml:"ImageLabels,omitempty" type:"Repeated"`
	// example:
	//
	// baselineCheck
	ImageService *string `json:"ImageService,omitempty" xml:"ImageService,omitempty"`
	// url
	//
	// example:
	//
	// https://www.aliyuncs.com/xxx.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// nonLabel
	Labels   *string   `json:"Labels,omitempty" xml:"Labels,omitempty"`
	NoLabels []*string `json:"NoLabels,omitempty" xml:"NoLabels,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// example:
	//
	// 1
	PageNum     *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	RequestFrom *string `json:"RequestFrom,omitempty" xml:"RequestFrom,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	RequestTime *string                                     `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	Result      []*GetScanResultResponseBodyDataItemsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	RiskLevel   *string                                     `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	RiskTips    *string                                     `json:"RiskTips,omitempty" xml:"RiskTips,omitempty"`
	RiskWords   *string                                     `json:"RiskWords,omitempty" xml:"RiskWords,omitempty"`
	// example:
	//
	// {}
	ScanResult *string `json:"ScanResult,omitempty" xml:"ScanResult,omitempty"`
	// example:
	//
	// 25
	Score          *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	SensitiveLevel *string  `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// 11
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// review
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// example:
	//
	// vi_s_EbrXb716LyBpkfwxyX5xyh-1A6RY9
	TaskId     *string                  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TextLabels []map[string]interface{} `json:"TextLabels,omitempty" xml:"TextLabels,omitempty" type:"Repeated"`
	// example:
	//
	// https://www.aliyuncs.com/xxx.png
	Thumbnail *string `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	// example:
	//
	// 00:00:40-00:00:42
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// example:
	//
	// https://www.aliyuncs.com/xxx.png
	Url         *string                  `json:"Url,omitempty" xml:"Url,omitempty"`
	VoiceLabels []map[string]interface{} `json:"VoiceLabels,omitempty" xml:"VoiceLabels,omitempty" type:"Repeated"`
	// example:
	//
	// True
	VoiceScanOpened *bool `json:"VoiceScanOpened,omitempty" xml:"VoiceScanOpened,omitempty"`
	// example:
	//
	// live_stream_detection
	VoiceService *string `json:"VoiceService,omitempty" xml:"VoiceService,omitempty"`
}

func (s GetScanResultResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s GetScanResultResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetScanResultResponseBodyDataItems) SetAttackLevel(v string) *GetScanResultResponseBodyDataItems {
	s.AttackLevel = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetContent(v string) *GetScanResultResponseBodyDataItems {
	s.Content = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetDataId(v string) *GetScanResultResponseBodyDataItems {
	s.DataId = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetEndTime(v string) *GetScanResultResponseBodyDataItems {
	s.EndTime = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetExtFeedback(v string) *GetScanResultResponseBodyDataItems {
	s.ExtFeedback = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetExtra(v map[string]interface{}) *GetScanResultResponseBodyDataItems {
	s.Extra = v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetFrameCount(v int64) *GetScanResultResponseBodyDataItems {
	s.FrameCount = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetGmtCreate(v string) *GetScanResultResponseBodyDataItems {
	s.GmtCreate = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetImageLabels(v []map[string]interface{}) *GetScanResultResponseBodyDataItems {
	s.ImageLabels = v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetImageService(v string) *GetScanResultResponseBodyDataItems {
	s.ImageService = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetImageUrl(v string) *GetScanResultResponseBodyDataItems {
	s.ImageUrl = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetLabels(v string) *GetScanResultResponseBodyDataItems {
	s.Labels = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetNoLabels(v []*string) *GetScanResultResponseBodyDataItems {
	s.NoLabels = v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetOffset(v int64) *GetScanResultResponseBodyDataItems {
	s.Offset = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetPageNum(v int64) *GetScanResultResponseBodyDataItems {
	s.PageNum = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetRequestFrom(v string) *GetScanResultResponseBodyDataItems {
	s.RequestFrom = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetRequestId(v string) *GetScanResultResponseBodyDataItems {
	s.RequestId = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetRequestTime(v string) *GetScanResultResponseBodyDataItems {
	s.RequestTime = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetResult(v []*GetScanResultResponseBodyDataItemsResult) *GetScanResultResponseBodyDataItems {
	s.Result = v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetRiskLevel(v string) *GetScanResultResponseBodyDataItems {
	s.RiskLevel = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetRiskTips(v string) *GetScanResultResponseBodyDataItems {
	s.RiskTips = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetRiskWords(v string) *GetScanResultResponseBodyDataItems {
	s.RiskWords = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetScanResult(v string) *GetScanResultResponseBodyDataItems {
	s.ScanResult = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetScore(v float32) *GetScanResultResponseBodyDataItems {
	s.Score = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetSensitiveLevel(v string) *GetScanResultResponseBodyDataItems {
	s.SensitiveLevel = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetServiceCode(v string) *GetScanResultResponseBodyDataItems {
	s.ServiceCode = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetStartTime(v string) *GetScanResultResponseBodyDataItems {
	s.StartTime = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetSuggestion(v string) *GetScanResultResponseBodyDataItems {
	s.Suggestion = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetTaskId(v string) *GetScanResultResponseBodyDataItems {
	s.TaskId = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetTextLabels(v []map[string]interface{}) *GetScanResultResponseBodyDataItems {
	s.TextLabels = v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetThumbnail(v string) *GetScanResultResponseBodyDataItems {
	s.Thumbnail = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetTimeStamp(v string) *GetScanResultResponseBodyDataItems {
	s.TimeStamp = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetUrl(v string) *GetScanResultResponseBodyDataItems {
	s.Url = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetVoiceLabels(v []map[string]interface{}) *GetScanResultResponseBodyDataItems {
	s.VoiceLabels = v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetVoiceScanOpened(v bool) *GetScanResultResponseBodyDataItems {
	s.VoiceScanOpened = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetVoiceService(v string) *GetScanResultResponseBodyDataItems {
	s.VoiceService = &v
	return s
}

type GetScanResultResponseBodyDataItemsResult struct {
	// example:
	//
	// 50.0
	Confidence  *string `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// politics
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetScanResultResponseBodyDataItemsResult) String() string {
	return tea.Prettify(s)
}

func (s GetScanResultResponseBodyDataItemsResult) GoString() string {
	return s.String()
}

func (s *GetScanResultResponseBodyDataItemsResult) SetConfidence(v string) *GetScanResultResponseBodyDataItemsResult {
	s.Confidence = &v
	return s
}

func (s *GetScanResultResponseBodyDataItemsResult) SetDescription(v string) *GetScanResultResponseBodyDataItemsResult {
	s.Description = &v
	return s
}

func (s *GetScanResultResponseBodyDataItemsResult) SetLabel(v string) *GetScanResultResponseBodyDataItemsResult {
	s.Label = &v
	return s
}

type GetScanResultResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScanResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScanResultResponse) GoString() string {
	return s.String()
}

func (s *GetScanResultResponse) SetHeaders(v map[string]*string) *GetScanResultResponse {
	s.Headers = v
	return s
}

func (s *GetScanResultResponse) SetStatusCode(v int32) *GetScanResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScanResultResponse) SetBody(v *GetScanResultResponseBody) *GetScanResultResponse {
	s.Body = v
	return s
}

type GetServiceConfRequest struct {
	// example:
	//
	// False
	ByDefault *bool `json:"ByDefault,omitempty" xml:"ByDefault,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// pornographic
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s GetServiceConfRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceConfRequest) GoString() string {
	return s.String()
}

func (s *GetServiceConfRequest) SetByDefault(v bool) *GetServiceConfRequest {
	s.ByDefault = &v
	return s
}

func (s *GetServiceConfRequest) SetRegionId(v string) *GetServiceConfRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceConfRequest) SetResourceType(v string) *GetServiceConfRequest {
	s.ResourceType = &v
	return s
}

func (s *GetServiceConfRequest) SetScene(v string) *GetServiceConfRequest {
	s.Scene = &v
	return s
}

func (s *GetServiceConfRequest) SetServiceCode(v string) *GetServiceConfRequest {
	s.ServiceCode = &v
	return s
}

type GetServiceConfResponseBody struct {
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {}
	CustomServiceConf map[string]interface{} `json:"CustomServiceConf,omitempty" xml:"CustomServiceConf,omitempty"`
	// example:
	//
	// 2023-01-17 12:29:56
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// {}
	Option map[string]interface{} `json:"Option,omitempty" xml:"Option,omitempty"`
	// example:
	//
	// 6CF2815C-****-****-B52E-FF6E2****492
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// UID。
	//
	// example:
	//
	// 17726*****370735
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetServiceConfResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceConfResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceConfResponseBody) SetClassify(v string) *GetServiceConfResponseBody {
	s.Classify = &v
	return s
}

func (s *GetServiceConfResponseBody) SetCode(v int32) *GetServiceConfResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceConfResponseBody) SetCustomServiceConf(v map[string]interface{}) *GetServiceConfResponseBody {
	s.CustomServiceConf = v
	return s
}

func (s *GetServiceConfResponseBody) SetGmtModified(v string) *GetServiceConfResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetServiceConfResponseBody) SetMsg(v string) *GetServiceConfResponseBody {
	s.Msg = &v
	return s
}

func (s *GetServiceConfResponseBody) SetOption(v map[string]interface{}) *GetServiceConfResponseBody {
	s.Option = v
	return s
}

func (s *GetServiceConfResponseBody) SetRequestId(v string) *GetServiceConfResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceConfResponseBody) SetResourceType(v string) *GetServiceConfResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetServiceConfResponseBody) SetServiceCode(v string) *GetServiceConfResponseBody {
	s.ServiceCode = &v
	return s
}

func (s *GetServiceConfResponseBody) SetSuccess(v bool) *GetServiceConfResponseBody {
	s.Success = &v
	return s
}

func (s *GetServiceConfResponseBody) SetUid(v string) *GetServiceConfResponseBody {
	s.Uid = &v
	return s
}

type GetServiceConfResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceConfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceConfResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceConfResponse) GoString() string {
	return s.String()
}

func (s *GetServiceConfResponse) SetHeaders(v map[string]*string) *GetServiceConfResponse {
	s.Headers = v
	return s
}

func (s *GetServiceConfResponse) SetStatusCode(v int32) *GetServiceConfResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceConfResponse) SetBody(v *GetServiceConfResponseBody) *GetServiceConfResponse {
	s.Body = v
	return s
}

type GetServiceConfigRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// nickname_detection
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s GetServiceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *GetServiceConfigRequest) SetRegionId(v string) *GetServiceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceConfigRequest) SetResourceType(v string) *GetServiceConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *GetServiceConfigRequest) SetServiceCode(v string) *GetServiceConfigRequest {
	s.ServiceCode = &v
	return s
}

type GetServiceConfigResponseBody struct {
	// example:
	//
	// 200
	Code *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetServiceConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetServiceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceConfigResponseBody) SetCode(v int32) *GetServiceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceConfigResponseBody) SetData(v *GetServiceConfigResponseBodyData) *GetServiceConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceConfigResponseBody) SetMsg(v string) *GetServiceConfigResponseBody {
	s.Msg = &v
	return s
}

func (s *GetServiceConfigResponseBody) SetRequestId(v string) *GetServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceConfigResponseBody) SetSuccess(v bool) *GetServiceConfigResponseBody {
	s.Success = &v
	return s
}

type GetServiceConfigResponseBodyData struct {
	CustomServiceConf *GetServiceConfigResponseBodyDataCustomServiceConf `json:"CustomServiceConf,omitempty" xml:"CustomServiceConf,omitempty" type:"Struct"`
	// example:
	//
	// 2024-05-06 03:07:44
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// nickname_detection
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// UID。
	//
	// example:
	//
	// 165379****31937
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetServiceConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetServiceConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceConfigResponseBodyData) SetCustomServiceConf(v *GetServiceConfigResponseBodyDataCustomServiceConf) *GetServiceConfigResponseBodyData {
	s.CustomServiceConf = v
	return s
}

func (s *GetServiceConfigResponseBodyData) SetGmtModified(v string) *GetServiceConfigResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetServiceConfigResponseBodyData) SetResourceType(v string) *GetServiceConfigResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *GetServiceConfigResponseBodyData) SetServiceCode(v string) *GetServiceConfigResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *GetServiceConfigResponseBodyData) SetUid(v string) *GetServiceConfigResponseBodyData {
	s.Uid = &v
	return s
}

type GetServiceConfigResponseBodyDataCustomServiceConf struct {
	KeywordFilterLibs  []*string `json:"KeywordFilterLibs,omitempty" xml:"KeywordFilterLibs,omitempty" type:"Repeated"`
	KeywordHitLibs     []*string `json:"KeywordHitLibs,omitempty" xml:"KeywordHitLibs,omitempty" type:"Repeated"`
	SimilarTextHitLibs []*string `json:"SimilarTextHitLibs,omitempty" xml:"SimilarTextHitLibs,omitempty" type:"Repeated"`
}

func (s GetServiceConfigResponseBodyDataCustomServiceConf) String() string {
	return tea.Prettify(s)
}

func (s GetServiceConfigResponseBodyDataCustomServiceConf) GoString() string {
	return s.String()
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConf) SetKeywordFilterLibs(v []*string) *GetServiceConfigResponseBodyDataCustomServiceConf {
	s.KeywordFilterLibs = v
	return s
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConf) SetKeywordHitLibs(v []*string) *GetServiceConfigResponseBodyDataCustomServiceConf {
	s.KeywordHitLibs = v
	return s
}

func (s *GetServiceConfigResponseBodyDataCustomServiceConf) SetSimilarTextHitLibs(v []*string) *GetServiceConfigResponseBodyDataCustomServiceConf {
	s.SimilarTextHitLibs = v
	return s
}

type GetServiceConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *GetServiceConfigResponse) SetHeaders(v map[string]*string) *GetServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *GetServiceConfigResponse) SetStatusCode(v int32) *GetServiceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceConfigResponse) SetBody(v *GetServiceConfigResponseBody) *GetServiceConfigResponse {
	s.Body = v
	return s
}

type GetServiceLabelConfigRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// nickname_detection
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s GetServiceLabelConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLabelConfigRequest) GoString() string {
	return s.String()
}

func (s *GetServiceLabelConfigRequest) SetRegionId(v string) *GetServiceLabelConfigRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceLabelConfigRequest) SetResourceType(v string) *GetServiceLabelConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *GetServiceLabelConfigRequest) SetServiceCode(v string) *GetServiceLabelConfigRequest {
	s.ServiceCode = &v
	return s
}

type GetServiceLabelConfigResponseBody struct {
	Data []interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetServiceLabelConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLabelConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceLabelConfigResponseBody) SetData(v []interface{}) *GetServiceLabelConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceLabelConfigResponseBody) SetRequestId(v string) *GetServiceLabelConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetServiceLabelConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceLabelConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceLabelConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceLabelConfigResponse) GoString() string {
	return s.String()
}

func (s *GetServiceLabelConfigResponse) SetHeaders(v map[string]*string) *GetServiceLabelConfigResponse {
	s.Headers = v
	return s
}

func (s *GetServiceLabelConfigResponse) SetStatusCode(v int32) *GetServiceLabelConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceLabelConfigResponse) SetBody(v *GetServiceLabelConfigResponseBody) *GetServiceLabelConfigResponse {
	s.Body = v
	return s
}

type GetStockOssCheckTasksListRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2023-06-18 02:08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// false
	IsInc *bool `json:"IsInc,omitempty" xml:"IsInc,omitempty"`
	// example:
	//
	// image
	MediaType *int32 `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sort     map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-06-17 02:08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// batch
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetStockOssCheckTasksListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStockOssCheckTasksListRequest) GoString() string {
	return s.String()
}

func (s *GetStockOssCheckTasksListRequest) SetCurrentPage(v int32) *GetStockOssCheckTasksListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetStockOssCheckTasksListRequest) SetEndTime(v string) *GetStockOssCheckTasksListRequest {
	s.EndTime = &v
	return s
}

func (s *GetStockOssCheckTasksListRequest) SetIsInc(v bool) *GetStockOssCheckTasksListRequest {
	s.IsInc = &v
	return s
}

func (s *GetStockOssCheckTasksListRequest) SetMediaType(v int32) *GetStockOssCheckTasksListRequest {
	s.MediaType = &v
	return s
}

func (s *GetStockOssCheckTasksListRequest) SetPageSize(v int32) *GetStockOssCheckTasksListRequest {
	s.PageSize = &v
	return s
}

func (s *GetStockOssCheckTasksListRequest) SetRegionId(v string) *GetStockOssCheckTasksListRequest {
	s.RegionId = &v
	return s
}

func (s *GetStockOssCheckTasksListRequest) SetSort(v map[string]*string) *GetStockOssCheckTasksListRequest {
	s.Sort = v
	return s
}

func (s *GetStockOssCheckTasksListRequest) SetStartTime(v string) *GetStockOssCheckTasksListRequest {
	s.StartTime = &v
	return s
}

func (s *GetStockOssCheckTasksListRequest) SetTaskType(v string) *GetStockOssCheckTasksListRequest {
	s.TaskType = &v
	return s
}

type GetStockOssCheckTasksListShrinkRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2023-06-18 02:08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// false
	IsInc *bool `json:"IsInc,omitempty" xml:"IsInc,omitempty"`
	// example:
	//
	// image
	MediaType *int32 `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-06-17 02:08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// batch
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetStockOssCheckTasksListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStockOssCheckTasksListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetCurrentPage(v int32) *GetStockOssCheckTasksListShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetEndTime(v string) *GetStockOssCheckTasksListShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetIsInc(v bool) *GetStockOssCheckTasksListShrinkRequest {
	s.IsInc = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetMediaType(v int32) *GetStockOssCheckTasksListShrinkRequest {
	s.MediaType = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetPageSize(v int32) *GetStockOssCheckTasksListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetRegionId(v string) *GetStockOssCheckTasksListShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetSortShrink(v string) *GetStockOssCheckTasksListShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetStartTime(v string) *GetStockOssCheckTasksListShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *GetStockOssCheckTasksListShrinkRequest) SetTaskType(v string) *GetStockOssCheckTasksListShrinkRequest {
	s.TaskType = &v
	return s
}

type GetStockOssCheckTasksListResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*GetStockOssCheckTasksListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetStockOssCheckTasksListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStockOssCheckTasksListResponseBody) GoString() string {
	return s.String()
}

func (s *GetStockOssCheckTasksListResponseBody) SetCurrentPage(v int32) *GetStockOssCheckTasksListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBody) SetItems(v []*GetStockOssCheckTasksListResponseBodyItems) *GetStockOssCheckTasksListResponseBody {
	s.Items = v
	return s
}

func (s *GetStockOssCheckTasksListResponseBody) SetPageSize(v int32) *GetStockOssCheckTasksListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBody) SetRequestId(v string) *GetStockOssCheckTasksListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBody) SetTotalCount(v int64) *GetStockOssCheckTasksListResponseBody {
	s.TotalCount = &v
	return s
}

type GetStockOssCheckTasksListResponseBodyItems struct {
	// example:
	//
	// tmp
	Buckets *string                                           `json:"Buckets,omitempty" xml:"Buckets,omitempty"`
	Config  *GetStockOssCheckTasksListResponseBodyItemsConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// example:
	//
	// 2024-01-10 11:42:31
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2
	FinishNum *int64 `json:"FinishNum,omitempty" xml:"FinishNum,omitempty"`
	// example:
	//
	// false
	IsInc *bool `json:"IsInc,omitempty" xml:"IsInc,omitempty"`
	// example:
	//
	// 02:00:00
	LastExecuteDate *string `json:"LastExecuteDate,omitempty" xml:"LastExecuteDate,omitempty"`
	// example:
	//
	// video
	MediaType *int32 `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 02:00:00
	NextExecuteDate *string `json:"NextExecuteDate,omitempty" xml:"NextExecuteDate,omitempty"`
	// example:
	//
	// 10
	ObjectNum *int64 `json:"ObjectNum,omitempty" xml:"ObjectNum,omitempty"`
	// example:
	//
	// 10
	SearchNum *int64 `json:"SearchNum,omitempty" xml:"SearchNum,omitempty"`
	// example:
	//
	// 2023-12-21 15:30:19
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// P_XHDUS
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// batch
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetStockOssCheckTasksListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s GetStockOssCheckTasksListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetBuckets(v string) *GetStockOssCheckTasksListResponseBodyItems {
	s.Buckets = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetConfig(v *GetStockOssCheckTasksListResponseBodyItemsConfig) *GetStockOssCheckTasksListResponseBodyItems {
	s.Config = v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetEndTime(v string) *GetStockOssCheckTasksListResponseBodyItems {
	s.EndTime = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetFinishNum(v int64) *GetStockOssCheckTasksListResponseBodyItems {
	s.FinishNum = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetIsInc(v bool) *GetStockOssCheckTasksListResponseBodyItems {
	s.IsInc = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetLastExecuteDate(v string) *GetStockOssCheckTasksListResponseBodyItems {
	s.LastExecuteDate = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetMediaType(v int32) *GetStockOssCheckTasksListResponseBodyItems {
	s.MediaType = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetNextExecuteDate(v string) *GetStockOssCheckTasksListResponseBodyItems {
	s.NextExecuteDate = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetObjectNum(v int64) *GetStockOssCheckTasksListResponseBodyItems {
	s.ObjectNum = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetSearchNum(v int64) *GetStockOssCheckTasksListResponseBodyItems {
	s.SearchNum = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetStartTime(v string) *GetStockOssCheckTasksListResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetStatus(v int32) *GetStockOssCheckTasksListResponseBodyItems {
	s.Status = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetTaskId(v string) *GetStockOssCheckTasksListResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetTaskName(v string) *GetStockOssCheckTasksListResponseBodyItems {
	s.TaskName = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItems) SetTaskType(v string) *GetStockOssCheckTasksListResponseBodyItems {
	s.TaskType = &v
	return s
}

type GetStockOssCheckTasksListResponseBodyItemsConfig struct {
	CallbackId *int64 `json:"CallbackId,omitempty" xml:"CallbackId,omitempty"`
	// example:
	//
	// false
	DistinctHistoryTasks *bool `json:"DistinctHistoryTasks,omitempty" xml:"DistinctHistoryTasks,omitempty"`
	// example:
	//
	// 2024-01-10 11:42:31
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	ExecuteDate *int32 `json:"ExecuteDate,omitempty" xml:"ExecuteDate,omitempty"`
	// example:
	//
	// 02:00:00
	ExecuteTime       *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	Freeze            *bool   `json:"Freeze,omitempty" xml:"Freeze,omitempty"`
	FreezeHighRisk1   *bool   `json:"FreezeHighRisk1,omitempty" xml:"FreezeHighRisk1,omitempty"`
	FreezeHighRisk2   *bool   `json:"FreezeHighRisk2,omitempty" xml:"FreezeHighRisk2,omitempty"`
	FreezeMediumRisk1 *bool   `json:"FreezeMediumRisk1,omitempty" xml:"FreezeMediumRisk1,omitempty"`
	FreezeMediumRisk2 *bool   `json:"FreezeMediumRisk2,omitempty" xml:"FreezeMediumRisk2,omitempty"`
	FreezeType        *string `json:"FreezeType,omitempty" xml:"FreezeType,omitempty"`
	// example:
	//
	// all
	PrefixFilterType *string   `json:"PrefixFilterType,omitempty" xml:"PrefixFilterType,omitempty"`
	PrefixFilters    []*string `json:"PrefixFilters,omitempty" xml:"PrefixFilters,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 10
	ScanLimit *int64 `json:"ScanLimit,omitempty" xml:"ScanLimit,omitempty"`
	// example:
	//
	// true
	ScanNoFileType *bool `json:"ScanNoFileType,omitempty" xml:"ScanNoFileType,omitempty"`
	// example:
	//
	// 0
	ScanResourceType *int32                                                              `json:"ScanResourceType,omitempty" xml:"ScanResourceType,omitempty"`
	ScanService      []*string                                                           `json:"ScanService,omitempty" xml:"ScanService,omitempty" type:"Repeated"`
	ScanServiceInfos []*GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos `json:"ScanServiceInfos,omitempty" xml:"ScanServiceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-12-21 15:30:19
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 0
	TaskCycle *int32 `json:"TaskCycle,omitempty" xml:"TaskCycle,omitempty"`
}

func (s GetStockOssCheckTasksListResponseBodyItemsConfig) String() string {
	return tea.Prettify(s)
}

func (s GetStockOssCheckTasksListResponseBodyItemsConfig) GoString() string {
	return s.String()
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetCallbackId(v int64) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.CallbackId = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetDistinctHistoryTasks(v bool) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.DistinctHistoryTasks = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetEndTime(v string) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.EndTime = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetExecuteDate(v int32) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.ExecuteDate = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetExecuteTime(v string) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.ExecuteTime = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetFreeze(v bool) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.Freeze = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetFreezeHighRisk1(v bool) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.FreezeHighRisk1 = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetFreezeHighRisk2(v bool) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.FreezeHighRisk2 = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetFreezeMediumRisk1(v bool) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.FreezeMediumRisk1 = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetFreezeMediumRisk2(v bool) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.FreezeMediumRisk2 = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetFreezeType(v string) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.FreezeType = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetPrefixFilterType(v string) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.PrefixFilterType = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetPrefixFilters(v []*string) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.PrefixFilters = v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetPriority(v int32) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.Priority = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetScanLimit(v int64) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.ScanLimit = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetScanNoFileType(v bool) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.ScanNoFileType = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetScanResourceType(v int32) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.ScanResourceType = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetScanService(v []*string) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.ScanService = v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetScanServiceInfos(v []*GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.ScanServiceInfos = v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetStartTime(v string) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.StartTime = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfig) SetTaskCycle(v int32) *GetStockOssCheckTasksListResponseBodyItemsConfig {
	s.TaskCycle = &v
	return s
}

type GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos struct {
	// example:
	//
	// baselineCheck
	CopyFrom *string `json:"CopyFrom,omitempty" xml:"CopyFrom,omitempty"`
	// example:
	//
	// false
	IsCopy *bool `json:"IsCopy,omitempty" xml:"IsCopy,omitempty"`
	// example:
	//
	// baselineCheck_01
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) GoString() string {
	return s.String()
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) SetCopyFrom(v string) *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos {
	s.CopyFrom = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) SetIsCopy(v bool) *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos {
	s.IsCopy = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) SetServiceCode(v string) *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos {
	s.ServiceCode = &v
	return s
}

func (s *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos) SetServiceName(v string) *GetStockOssCheckTasksListResponseBodyItemsConfigScanServiceInfos {
	s.ServiceName = &v
	return s
}

type GetStockOssCheckTasksListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStockOssCheckTasksListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStockOssCheckTasksListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStockOssCheckTasksListResponse) GoString() string {
	return s.String()
}

func (s *GetStockOssCheckTasksListResponse) SetHeaders(v map[string]*string) *GetStockOssCheckTasksListResponse {
	s.Headers = v
	return s
}

func (s *GetStockOssCheckTasksListResponse) SetStatusCode(v int32) *GetStockOssCheckTasksListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStockOssCheckTasksListResponse) SetBody(v *GetStockOssCheckTasksListResponseBody) *GetStockOssCheckTasksListResponse {
	s.Body = v
	return s
}

type GetTextScanResultRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 10
	PageSize *int32             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query    map[string]*string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sort     map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetTextScanResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTextScanResultRequest) GoString() string {
	return s.String()
}

func (s *GetTextScanResultRequest) SetCurrentPage(v int32) *GetTextScanResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetTextScanResultRequest) SetEndDate(v string) *GetTextScanResultRequest {
	s.EndDate = &v
	return s
}

func (s *GetTextScanResultRequest) SetPageSize(v int32) *GetTextScanResultRequest {
	s.PageSize = &v
	return s
}

func (s *GetTextScanResultRequest) SetQuery(v map[string]*string) *GetTextScanResultRequest {
	s.Query = v
	return s
}

func (s *GetTextScanResultRequest) SetRegionId(v string) *GetTextScanResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetTextScanResultRequest) SetSort(v map[string]*string) *GetTextScanResultRequest {
	s.Sort = v
	return s
}

func (s *GetTextScanResultRequest) SetStartDate(v string) *GetTextScanResultRequest {
	s.StartDate = &v
	return s
}

type GetTextScanResultShrinkRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 10
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryShrink *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetTextScanResultShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTextScanResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetTextScanResultShrinkRequest) SetCurrentPage(v int32) *GetTextScanResultShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetTextScanResultShrinkRequest) SetEndDate(v string) *GetTextScanResultShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetTextScanResultShrinkRequest) SetPageSize(v int32) *GetTextScanResultShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetTextScanResultShrinkRequest) SetQueryShrink(v string) *GetTextScanResultShrinkRequest {
	s.QueryShrink = &v
	return s
}

func (s *GetTextScanResultShrinkRequest) SetRegionId(v string) *GetTextScanResultShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetTextScanResultShrinkRequest) SetSortShrink(v string) *GetTextScanResultShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *GetTextScanResultShrinkRequest) SetStartDate(v string) *GetTextScanResultShrinkRequest {
	s.StartDate = &v
	return s
}

type GetTextScanResultResponseBody struct {
	// example:
	//
	// 200
	Code *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTextScanResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTextScanResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTextScanResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTextScanResultResponseBody) SetCode(v int32) *GetTextScanResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetTextScanResultResponseBody) SetData(v *GetTextScanResultResponseBodyData) *GetTextScanResultResponseBody {
	s.Data = v
	return s
}

func (s *GetTextScanResultResponseBody) SetMsg(v string) *GetTextScanResultResponseBody {
	s.Msg = &v
	return s
}

func (s *GetTextScanResultResponseBody) SetRequestId(v string) *GetTextScanResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTextScanResultResponseBody) SetSuccess(v bool) *GetTextScanResultResponseBody {
	s.Success = &v
	return s
}

type GetTextScanResultResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int32                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*GetTextScanResultResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetTextScanResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTextScanResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTextScanResultResponseBodyData) SetCurrentPage(v int32) *GetTextScanResultResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetTextScanResultResponseBodyData) SetItems(v []*GetTextScanResultResponseBodyDataItems) *GetTextScanResultResponseBodyData {
	s.Items = v
	return s
}

func (s *GetTextScanResultResponseBodyData) SetPageSize(v int32) *GetTextScanResultResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetTextScanResultResponseBodyData) SetTotalCount(v int64) *GetTextScanResultResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetTextScanResultResponseBodyDataItems struct {
	BailianRequestId *string `json:"BailianRequestId,omitempty" xml:"BailianRequestId,omitempty"`
	Content          *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// miss
	ExtFeedback *string `json:"ExtFeedback,omitempty" xml:"ExtFeedback,omitempty"`
	// example:
	//
	// {}
	Extra map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// 2023-07-11 14:21:36
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// nonLabel
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2023-07-11 14:21:36
	RequestTime *string                                         `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	Result      []*GetTextScanResultResponseBodyDataItemsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	RiskLevel   *string                                         `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// {}
	ScanResult *string `json:"ScanResult,omitempty" xml:"ScanResult,omitempty"`
	// example:
	//
	// 20
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// nickname_detection
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// review
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// example:
	//
	// txtwkgb******AsYNXoJswy-1Aa1Qk
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTextScanResultResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s GetTextScanResultResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetTextScanResultResponseBodyDataItems) SetBailianRequestId(v string) *GetTextScanResultResponseBodyDataItems {
	s.BailianRequestId = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetContent(v string) *GetTextScanResultResponseBodyDataItems {
	s.Content = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetExtFeedback(v string) *GetTextScanResultResponseBodyDataItems {
	s.ExtFeedback = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetExtra(v map[string]interface{}) *GetTextScanResultResponseBodyDataItems {
	s.Extra = v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetGmtCreate(v string) *GetTextScanResultResponseBodyDataItems {
	s.GmtCreate = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetLabels(v string) *GetTextScanResultResponseBodyDataItems {
	s.Labels = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetRequestId(v string) *GetTextScanResultResponseBodyDataItems {
	s.RequestId = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetRequestTime(v string) *GetTextScanResultResponseBodyDataItems {
	s.RequestTime = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetResult(v []*GetTextScanResultResponseBodyDataItemsResult) *GetTextScanResultResponseBodyDataItems {
	s.Result = v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetRiskLevel(v string) *GetTextScanResultResponseBodyDataItems {
	s.RiskLevel = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetScanResult(v string) *GetTextScanResultResponseBodyDataItems {
	s.ScanResult = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetScore(v float32) *GetTextScanResultResponseBodyDataItems {
	s.Score = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetServiceCode(v string) *GetTextScanResultResponseBodyDataItems {
	s.ServiceCode = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetSuggestion(v string) *GetTextScanResultResponseBodyDataItems {
	s.Suggestion = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetTaskId(v string) *GetTextScanResultResponseBodyDataItems {
	s.TaskId = &v
	return s
}

type GetTextScanResultResponseBodyDataItemsResult struct {
	// example:
	//
	// 25.0
	Confidence  *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// political_n
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetTextScanResultResponseBodyDataItemsResult) String() string {
	return tea.Prettify(s)
}

func (s GetTextScanResultResponseBodyDataItemsResult) GoString() string {
	return s.String()
}

func (s *GetTextScanResultResponseBodyDataItemsResult) SetConfidence(v float32) *GetTextScanResultResponseBodyDataItemsResult {
	s.Confidence = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItemsResult) SetDescription(v string) *GetTextScanResultResponseBodyDataItemsResult {
	s.Description = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItemsResult) SetLabel(v string) *GetTextScanResultResponseBodyDataItemsResult {
	s.Label = &v
	return s
}

type GetTextScanResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTextScanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTextScanResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTextScanResultResponse) GoString() string {
	return s.String()
}

func (s *GetTextScanResultResponse) SetHeaders(v map[string]*string) *GetTextScanResultResponse {
	s.Headers = v
	return s
}

func (s *GetTextScanResultResponse) SetStatusCode(v int32) *GetTextScanResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTextScanResultResponse) SetBody(v *GetTextScanResultResponseBody) *GetTextScanResultResponse {
	s.Body = v
	return s
}

type GetUploadInfoRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetUploadInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUploadInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUploadInfoRequest) SetName(v string) *GetUploadInfoRequest {
	s.Name = &v
	return s
}

func (s *GetUploadInfoRequest) SetRegionId(v string) *GetUploadInfoRequest {
	s.RegionId = &v
	return s
}

func (s *GetUploadInfoRequest) SetResourceType(v string) *GetUploadInfoRequest {
	s.ResourceType = &v
	return s
}

type GetUploadInfoResponseBody struct {
	// example:
	//
	// LTAI5t9HM*****EXQmw3DVH
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 900
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// example:
	//
	// image/upload/xxx
	Folder *string `json:"Folder,omitempty" xml:"Folder,omitempty"`
	// example:
	//
	// https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// image/upload/xxx
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// OK
	Msg  *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// xxxx
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// iyu7VHblYj+mEF9p46cdGOlNPAw=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUploadInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUploadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadInfoResponseBody) SetAccessId(v string) *GetUploadInfoResponseBody {
	s.AccessId = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetCode(v int32) *GetUploadInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetExpire(v int64) *GetUploadInfoResponseBody {
	s.Expire = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetFolder(v string) *GetUploadInfoResponseBody {
	s.Folder = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetHost(v string) *GetUploadInfoResponseBody {
	s.Host = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetHttpStatusCode(v int32) *GetUploadInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetKey(v string) *GetUploadInfoResponseBody {
	s.Key = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetMsg(v string) *GetUploadInfoResponseBody {
	s.Msg = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetName(v string) *GetUploadInfoResponseBody {
	s.Name = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetPolicy(v string) *GetUploadInfoResponseBody {
	s.Policy = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetRequestId(v string) *GetUploadInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetSignature(v string) *GetUploadInfoResponseBody {
	s.Signature = &v
	return s
}

func (s *GetUploadInfoResponseBody) SetSuccess(v bool) *GetUploadInfoResponseBody {
	s.Success = &v
	return s
}

type GetUploadInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUploadInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUploadInfoResponse) SetHeaders(v map[string]*string) *GetUploadInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUploadInfoResponse) SetStatusCode(v int32) *GetUploadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadInfoResponse) SetBody(v *GetUploadInfoResponseBody) *GetUploadInfoResponse {
	s.Body = v
	return s
}

type GetUserBuyStatusRequest struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetUserBuyStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserBuyStatusRequest) GoString() string {
	return s.String()
}

func (s *GetUserBuyStatusRequest) SetCommodityCode(v string) *GetUserBuyStatusRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetUserBuyStatusRequest) SetRegionId(v string) *GetUserBuyStatusRequest {
	s.RegionId = &v
	return s
}

type GetUserBuyStatusResponseBody struct {
	// example:
	//
	// 200
	Code *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetUserBuyStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserBuyStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserBuyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserBuyStatusResponseBody) SetCode(v int32) *GetUserBuyStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserBuyStatusResponseBody) SetData(v *GetUserBuyStatusResponseBodyData) *GetUserBuyStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetUserBuyStatusResponseBody) SetMsg(v string) *GetUserBuyStatusResponseBody {
	s.Msg = &v
	return s
}

func (s *GetUserBuyStatusResponseBody) SetRequestId(v string) *GetUserBuyStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserBuyStatusResponseBody) SetSuccess(v bool) *GetUserBuyStatusResponseBody {
	s.Success = &v
	return s
}

type GetUserBuyStatusResponseBodyData struct {
	// Bid。
	//
	// example:
	//
	// 26842
	Bid *int64 `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// example:
	//
	// True
	Buy *bool `json:"Buy,omitempty" xml:"Buy,omitempty"`
	// example:
	//
	// False
	Indebt *bool   `json:"Indebt,omitempty" xml:"Indebt,omitempty"`
	Tag    *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetUserBuyStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserBuyStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserBuyStatusResponseBodyData) SetBid(v int64) *GetUserBuyStatusResponseBodyData {
	s.Bid = &v
	return s
}

func (s *GetUserBuyStatusResponseBodyData) SetBuy(v bool) *GetUserBuyStatusResponseBodyData {
	s.Buy = &v
	return s
}

func (s *GetUserBuyStatusResponseBodyData) SetIndebt(v bool) *GetUserBuyStatusResponseBodyData {
	s.Indebt = &v
	return s
}

func (s *GetUserBuyStatusResponseBodyData) SetTag(v string) *GetUserBuyStatusResponseBodyData {
	s.Tag = &v
	return s
}

type GetUserBuyStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserBuyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserBuyStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserBuyStatusResponse) GoString() string {
	return s.String()
}

func (s *GetUserBuyStatusResponse) SetHeaders(v map[string]*string) *GetUserBuyStatusResponse {
	s.Headers = v
	return s
}

func (s *GetUserBuyStatusResponse) SetStatusCode(v int32) *GetUserBuyStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserBuyStatusResponse) SetBody(v *GetUserBuyStatusResponseBody) *GetUserBuyStatusResponse {
	s.Body = v
	return s
}

type ListAnswerLibRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAnswerLibRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAnswerLibRequest) GoString() string {
	return s.String()
}

func (s *ListAnswerLibRequest) SetRegionId(v string) *ListAnswerLibRequest {
	s.RegionId = &v
	return s
}

type ListAnswerLibResponseBody struct {
	Data []*ListAnswerLibResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAnswerLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAnswerLibResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnswerLibResponseBody) SetData(v []*ListAnswerLibResponseBodyData) *ListAnswerLibResponseBody {
	s.Data = v
	return s
}

func (s *ListAnswerLibResponseBody) SetRequestId(v string) *ListAnswerLibResponseBody {
	s.RequestId = &v
	return s
}

type ListAnswerLibResponseBodyData struct {
	// example:
	//
	// 100
	AnswerCount *int32 `json:"AnswerCount,omitempty" xml:"AnswerCount,omitempty"`
	// example:
	//
	// 2024-06-03 18:15:01
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// alxxx
	LibId   *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// UID。
	//
	// example:
	//
	// 1643953****74290
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListAnswerLibResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAnswerLibResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAnswerLibResponseBodyData) SetAnswerCount(v int32) *ListAnswerLibResponseBodyData {
	s.AnswerCount = &v
	return s
}

func (s *ListAnswerLibResponseBodyData) SetGmtModified(v string) *ListAnswerLibResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListAnswerLibResponseBodyData) SetLibId(v string) *ListAnswerLibResponseBodyData {
	s.LibId = &v
	return s
}

func (s *ListAnswerLibResponseBodyData) SetLibName(v string) *ListAnswerLibResponseBodyData {
	s.LibName = &v
	return s
}

func (s *ListAnswerLibResponseBodyData) SetUid(v string) *ListAnswerLibResponseBodyData {
	s.Uid = &v
	return s
}

type ListAnswerLibResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAnswerLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAnswerLibResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAnswerLibResponse) GoString() string {
	return s.String()
}

func (s *ListAnswerLibResponse) SetHeaders(v map[string]*string) *ListAnswerLibResponse {
	s.Headers = v
	return s
}

func (s *ListAnswerLibResponse) SetStatusCode(v int32) *ListAnswerLibResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAnswerLibResponse) SetBody(v *ListAnswerLibResponseBody) *ListAnswerLibResponse {
	s.Body = v
	return s
}

type ListCallbackRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCallbackRequest) GoString() string {
	return s.String()
}

func (s *ListCallbackRequest) SetRegionId(v string) *ListCallbackRequest {
	s.RegionId = &v
	return s
}

type ListCallbackResponseBody struct {
	Data []*ListCallbackResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *ListCallbackResponseBody) SetData(v []*ListCallbackResponseBodyData) *ListCallbackResponseBody {
	s.Data = v
	return s
}

func (s *ListCallbackResponseBody) SetRequestId(v string) *ListCallbackResponseBody {
	s.RequestId = &v
	return s
}

type ListCallbackResponseBodyData struct {
	// example:
	//
	// SHA256
	CryptType *string `json:"CryptType,omitempty" xml:"CryptType,omitempty"`
	// example:
	//
	// 2024-06-03 15:20:14
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-06-03 15:20:14
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 11234
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// all
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// Seed。
	//
	// example:
	//
	// cbupVnpBjkgjFxfINMHKkrHS-1zZPUm
	Seed *string `json:"Seed,omitempty" xml:"Seed,omitempty"`
	// UID。
	//
	// example:
	//
	// 16537*****831937
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// example:
	//
	// https://console.aliyun.com/
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListCallbackResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCallbackResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCallbackResponseBodyData) SetCryptType(v string) *ListCallbackResponseBodyData {
	s.CryptType = &v
	return s
}

func (s *ListCallbackResponseBodyData) SetGmtCreate(v string) *ListCallbackResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListCallbackResponseBodyData) SetGmtModified(v string) *ListCallbackResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListCallbackResponseBodyData) SetId(v int64) *ListCallbackResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListCallbackResponseBodyData) SetName(v string) *ListCallbackResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListCallbackResponseBodyData) SetScope(v string) *ListCallbackResponseBodyData {
	s.Scope = &v
	return s
}

func (s *ListCallbackResponseBodyData) SetSeed(v string) *ListCallbackResponseBodyData {
	s.Seed = &v
	return s
}

func (s *ListCallbackResponseBodyData) SetUid(v string) *ListCallbackResponseBodyData {
	s.Uid = &v
	return s
}

func (s *ListCallbackResponseBodyData) SetUrl(v string) *ListCallbackResponseBodyData {
	s.Url = &v
	return s
}

type ListCallbackResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCallbackResponse) GoString() string {
	return s.String()
}

func (s *ListCallbackResponse) SetHeaders(v map[string]*string) *ListCallbackResponse {
	s.Headers = v
	return s
}

func (s *ListCallbackResponse) SetStatusCode(v int32) *ListCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCallbackResponse) SetBody(v *ListCallbackResponseBody) *ListCallbackResponse {
	s.Body = v
	return s
}

type ListImageLibRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListImageLibRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImageLibRequest) GoString() string {
	return s.String()
}

func (s *ListImageLibRequest) SetRegionId(v string) *ListImageLibRequest {
	s.RegionId = &v
	return s
}

type ListImageLibResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	LibList        []*ListImageLibResponseBodyLibList `json:"LibList,omitempty" xml:"LibList,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListImageLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImageLibResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageLibResponseBody) SetCode(v int32) *ListImageLibResponseBody {
	s.Code = &v
	return s
}

func (s *ListImageLibResponseBody) SetHttpStatusCode(v int32) *ListImageLibResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListImageLibResponseBody) SetLibList(v []*ListImageLibResponseBodyLibList) *ListImageLibResponseBody {
	s.LibList = v
	return s
}

func (s *ListImageLibResponseBody) SetMsg(v string) *ListImageLibResponseBody {
	s.Msg = &v
	return s
}

func (s *ListImageLibResponseBody) SetRequestId(v string) *ListImageLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageLibResponseBody) SetSuccess(v bool) *ListImageLibResponseBody {
	s.Success = &v
	return s
}

type ListImageLibResponseBodyLibList struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 1
	FreeInspection *int32 `json:"FreeInspection,omitempty" xml:"FreeInspection,omitempty"`
	// example:
	//
	// 2024-06-03 15:20:14
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-06-03 15:20:14
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1
	ImageNum *int64 `json:"ImageNum,omitempty" xml:"ImageNum,omitempty"`
	// example:
	//
	// custom_xxxx
	LibId   *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s ListImageLibResponseBodyLibList) String() string {
	return tea.Prettify(s)
}

func (s ListImageLibResponseBodyLibList) GoString() string {
	return s.String()
}

func (s *ListImageLibResponseBodyLibList) SetComment(v string) *ListImageLibResponseBodyLibList {
	s.Comment = &v
	return s
}

func (s *ListImageLibResponseBodyLibList) SetFreeInspection(v int32) *ListImageLibResponseBodyLibList {
	s.FreeInspection = &v
	return s
}

func (s *ListImageLibResponseBodyLibList) SetGmtCreate(v string) *ListImageLibResponseBodyLibList {
	s.GmtCreate = &v
	return s
}

func (s *ListImageLibResponseBodyLibList) SetGmtModified(v string) *ListImageLibResponseBodyLibList {
	s.GmtModified = &v
	return s
}

func (s *ListImageLibResponseBodyLibList) SetImageNum(v int64) *ListImageLibResponseBodyLibList {
	s.ImageNum = &v
	return s
}

func (s *ListImageLibResponseBodyLibList) SetLibId(v string) *ListImageLibResponseBodyLibList {
	s.LibId = &v
	return s
}

func (s *ListImageLibResponseBodyLibList) SetLibName(v string) *ListImageLibResponseBodyLibList {
	s.LibName = &v
	return s
}

type ListImageLibResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImageLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImageLibResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImageLibResponse) GoString() string {
	return s.String()
}

func (s *ListImageLibResponse) SetHeaders(v map[string]*string) *ListImageLibResponse {
	s.Headers = v
	return s
}

func (s *ListImageLibResponse) SetStatusCode(v int32) *ListImageLibResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageLibResponse) SetBody(v *ListImageLibResponseBody) *ListImageLibResponse {
	s.Body = v
	return s
}

type ListImagesFromLibRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 112
	ImgId *string `json:"ImgId,omitempty" xml:"ImgId,omitempty"`
	// example:
	//
	// custom_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sort     map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ListImagesFromLibRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesFromLibRequest) GoString() string {
	return s.String()
}

func (s *ListImagesFromLibRequest) SetCurrentPage(v int32) *ListImagesFromLibRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListImagesFromLibRequest) SetEndDate(v string) *ListImagesFromLibRequest {
	s.EndDate = &v
	return s
}

func (s *ListImagesFromLibRequest) SetImgId(v string) *ListImagesFromLibRequest {
	s.ImgId = &v
	return s
}

func (s *ListImagesFromLibRequest) SetLibId(v string) *ListImagesFromLibRequest {
	s.LibId = &v
	return s
}

func (s *ListImagesFromLibRequest) SetPageSize(v int32) *ListImagesFromLibRequest {
	s.PageSize = &v
	return s
}

func (s *ListImagesFromLibRequest) SetRegionId(v string) *ListImagesFromLibRequest {
	s.RegionId = &v
	return s
}

func (s *ListImagesFromLibRequest) SetSort(v map[string]*string) *ListImagesFromLibRequest {
	s.Sort = v
	return s
}

func (s *ListImagesFromLibRequest) SetStartDate(v string) *ListImagesFromLibRequest {
	s.StartDate = &v
	return s
}

type ListImagesFromLibShrinkRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 112
	ImgId *string `json:"ImgId,omitempty" xml:"ImgId,omitempty"`
	// example:
	//
	// custom_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ListImagesFromLibShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesFromLibShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListImagesFromLibShrinkRequest) SetCurrentPage(v int32) *ListImagesFromLibShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListImagesFromLibShrinkRequest) SetEndDate(v string) *ListImagesFromLibShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *ListImagesFromLibShrinkRequest) SetImgId(v string) *ListImagesFromLibShrinkRequest {
	s.ImgId = &v
	return s
}

func (s *ListImagesFromLibShrinkRequest) SetLibId(v string) *ListImagesFromLibShrinkRequest {
	s.LibId = &v
	return s
}

func (s *ListImagesFromLibShrinkRequest) SetPageSize(v int32) *ListImagesFromLibShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListImagesFromLibShrinkRequest) SetRegionId(v string) *ListImagesFromLibShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListImagesFromLibShrinkRequest) SetSortShrink(v string) *ListImagesFromLibShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *ListImagesFromLibShrinkRequest) SetStartDate(v string) *ListImagesFromLibShrinkRequest {
	s.StartDate = &v
	return s
}

type ListImagesFromLibResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          []*ListImagesFromLibResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListImagesFromLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesFromLibResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesFromLibResponseBody) SetCode(v int32) *ListImagesFromLibResponseBody {
	s.Code = &v
	return s
}

func (s *ListImagesFromLibResponseBody) SetCurrentPage(v int32) *ListImagesFromLibResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListImagesFromLibResponseBody) SetHttpStatusCode(v int32) *ListImagesFromLibResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListImagesFromLibResponseBody) SetItems(v []*ListImagesFromLibResponseBodyItems) *ListImagesFromLibResponseBody {
	s.Items = v
	return s
}

func (s *ListImagesFromLibResponseBody) SetMsg(v string) *ListImagesFromLibResponseBody {
	s.Msg = &v
	return s
}

func (s *ListImagesFromLibResponseBody) SetPageSize(v int32) *ListImagesFromLibResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListImagesFromLibResponseBody) SetRequestId(v string) *ListImagesFromLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesFromLibResponseBody) SetSuccess(v bool) *ListImagesFromLibResponseBody {
	s.Success = &v
	return s
}

func (s *ListImagesFromLibResponseBody) SetTotalCount(v int32) *ListImagesFromLibResponseBody {
	s.TotalCount = &v
	return s
}

type ListImagesFromLibResponseBodyItems struct {
	// example:
	//
	// 2022-11-30 16:30:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 112
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/image/upload/IMG_2123.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// https://oss-cip-shanghai.oss-cn-shanghai.aliyuncs.com/image/upload/IMG_2123.jpg
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
}

func (s ListImagesFromLibResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListImagesFromLibResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListImagesFromLibResponseBodyItems) SetGmtCreate(v string) *ListImagesFromLibResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *ListImagesFromLibResponseBodyItems) SetImageId(v string) *ListImagesFromLibResponseBodyItems {
	s.ImageId = &v
	return s
}

func (s *ListImagesFromLibResponseBodyItems) SetImageUrl(v string) *ListImagesFromLibResponseBodyItems {
	s.ImageUrl = &v
	return s
}

func (s *ListImagesFromLibResponseBodyItems) SetThumbnailUrl(v string) *ListImagesFromLibResponseBodyItems {
	s.ThumbnailUrl = &v
	return s
}

type ListImagesFromLibResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImagesFromLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImagesFromLibResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesFromLibResponse) GoString() string {
	return s.String()
}

func (s *ListImagesFromLibResponse) SetHeaders(v map[string]*string) *ListImagesFromLibResponse {
	s.Headers = v
	return s
}

func (s *ListImagesFromLibResponse) SetStatusCode(v int32) *ListImagesFromLibResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesFromLibResponse) SetBody(v *ListImagesFromLibResponseBody) *ListImagesFromLibResponse {
	s.Body = v
	return s
}

type ListKeywordLibsRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListKeywordLibsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKeywordLibsRequest) GoString() string {
	return s.String()
}

func (s *ListKeywordLibsRequest) SetRegionId(v string) *ListKeywordLibsRequest {
	s.RegionId = &v
	return s
}

type ListKeywordLibsResponseBody struct {
	// example:
	//
	// 200
	Code *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListKeywordLibsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListKeywordLibsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListKeywordLibsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKeywordLibsResponseBody) SetCode(v int32) *ListKeywordLibsResponseBody {
	s.Code = &v
	return s
}

func (s *ListKeywordLibsResponseBody) SetData(v []*ListKeywordLibsResponseBodyData) *ListKeywordLibsResponseBody {
	s.Data = v
	return s
}

func (s *ListKeywordLibsResponseBody) SetMsg(v string) *ListKeywordLibsResponseBody {
	s.Msg = &v
	return s
}

func (s *ListKeywordLibsResponseBody) SetRequestId(v string) *ListKeywordLibsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKeywordLibsResponseBody) SetSuccess(v bool) *ListKeywordLibsResponseBody {
	s.Success = &v
	return s
}

type ListKeywordLibsResponseBodyData struct {
	// example:
	//
	// 2022-11-30 16:30:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 10
	KeywordCount *string `json:"KeywordCount,omitempty" xml:"KeywordCount,omitempty"`
	// example:
	//
	// custom_xxxxx
	LibId        *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	LibName      *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	ServiceCodes *string `json:"ServiceCodes,omitempty" xml:"ServiceCodes,omitempty"`
	// UID。
	//
	// example:
	//
	// 19964*****086772
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListKeywordLibsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListKeywordLibsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKeywordLibsResponseBodyData) SetGmtModified(v string) *ListKeywordLibsResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListKeywordLibsResponseBodyData) SetKeywordCount(v string) *ListKeywordLibsResponseBodyData {
	s.KeywordCount = &v
	return s
}

func (s *ListKeywordLibsResponseBodyData) SetLibId(v string) *ListKeywordLibsResponseBodyData {
	s.LibId = &v
	return s
}

func (s *ListKeywordLibsResponseBodyData) SetLibName(v string) *ListKeywordLibsResponseBodyData {
	s.LibName = &v
	return s
}

func (s *ListKeywordLibsResponseBodyData) SetServiceCodes(v string) *ListKeywordLibsResponseBodyData {
	s.ServiceCodes = &v
	return s
}

func (s *ListKeywordLibsResponseBodyData) SetUid(v string) *ListKeywordLibsResponseBodyData {
	s.Uid = &v
	return s
}

type ListKeywordLibsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKeywordLibsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKeywordLibsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListKeywordLibsResponse) GoString() string {
	return s.String()
}

func (s *ListKeywordLibsResponse) SetHeaders(v map[string]*string) *ListKeywordLibsResponse {
	s.Headers = v
	return s
}

func (s *ListKeywordLibsResponse) SetStatusCode(v int32) *ListKeywordLibsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKeywordLibsResponse) SetBody(v *ListKeywordLibsResponseBody) *ListKeywordLibsResponse {
	s.Body = v
	return s
}

type ListKeywordsRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// custom_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sort     map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	Word     *string            `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s ListKeywordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKeywordsRequest) GoString() string {
	return s.String()
}

func (s *ListKeywordsRequest) SetCurrentPage(v int32) *ListKeywordsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListKeywordsRequest) SetLibId(v string) *ListKeywordsRequest {
	s.LibId = &v
	return s
}

func (s *ListKeywordsRequest) SetPageSize(v int32) *ListKeywordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListKeywordsRequest) SetRegionId(v string) *ListKeywordsRequest {
	s.RegionId = &v
	return s
}

func (s *ListKeywordsRequest) SetSort(v map[string]*string) *ListKeywordsRequest {
	s.Sort = v
	return s
}

func (s *ListKeywordsRequest) SetWord(v string) *ListKeywordsRequest {
	s.Word = &v
	return s
}

type ListKeywordsShrinkRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// custom_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	Word       *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s ListKeywordsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKeywordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListKeywordsShrinkRequest) SetCurrentPage(v int32) *ListKeywordsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListKeywordsShrinkRequest) SetLibId(v string) *ListKeywordsShrinkRequest {
	s.LibId = &v
	return s
}

func (s *ListKeywordsShrinkRequest) SetPageSize(v int32) *ListKeywordsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListKeywordsShrinkRequest) SetRegionId(v string) *ListKeywordsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListKeywordsShrinkRequest) SetSortShrink(v string) *ListKeywordsShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *ListKeywordsShrinkRequest) SetWord(v string) *ListKeywordsShrinkRequest {
	s.Word = &v
	return s
}

type ListKeywordsResponseBody struct {
	// example:
	//
	// 200
	Code *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListKeywordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListKeywordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListKeywordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKeywordsResponseBody) SetCode(v int32) *ListKeywordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListKeywordsResponseBody) SetData(v *ListKeywordsResponseBodyData) *ListKeywordsResponseBody {
	s.Data = v
	return s
}

func (s *ListKeywordsResponseBody) SetMsg(v string) *ListKeywordsResponseBody {
	s.Msg = &v
	return s
}

func (s *ListKeywordsResponseBody) SetRequestId(v string) *ListKeywordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKeywordsResponseBody) SetSuccess(v bool) *ListKeywordsResponseBody {
	s.Success = &v
	return s
}

type ListKeywordsResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int32                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*ListKeywordsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKeywordsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListKeywordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKeywordsResponseBodyData) SetCurrentPage(v int32) *ListKeywordsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListKeywordsResponseBodyData) SetItems(v []*ListKeywordsResponseBodyDataItems) *ListKeywordsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListKeywordsResponseBodyData) SetPageSize(v int32) *ListKeywordsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListKeywordsResponseBodyData) SetTotalCount(v int32) *ListKeywordsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListKeywordsResponseBodyDataItems struct {
	// example:
	//
	// 2023-06-03 14:43:03
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2023-06-03 14:43:03
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 112
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// custom_xxxx
	KeywordLibId *string `json:"KeywordLibId,omitempty" xml:"KeywordLibId,omitempty"`
	// example:
	//
	// 4205334
	KeywordMd5Id *int64  `json:"KeywordMd5Id,omitempty" xml:"KeywordMd5Id,omitempty"`
	Word         *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s ListKeywordsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListKeywordsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListKeywordsResponseBodyDataItems) SetGmtCreate(v string) *ListKeywordsResponseBodyDataItems {
	s.GmtCreate = &v
	return s
}

func (s *ListKeywordsResponseBodyDataItems) SetGmtModified(v string) *ListKeywordsResponseBodyDataItems {
	s.GmtModified = &v
	return s
}

func (s *ListKeywordsResponseBodyDataItems) SetId(v int64) *ListKeywordsResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListKeywordsResponseBodyDataItems) SetKeywordLibId(v string) *ListKeywordsResponseBodyDataItems {
	s.KeywordLibId = &v
	return s
}

func (s *ListKeywordsResponseBodyDataItems) SetKeywordMd5Id(v int64) *ListKeywordsResponseBodyDataItems {
	s.KeywordMd5Id = &v
	return s
}

func (s *ListKeywordsResponseBodyDataItems) SetWord(v string) *ListKeywordsResponseBodyDataItems {
	s.Word = &v
	return s
}

type ListKeywordsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKeywordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKeywordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListKeywordsResponse) GoString() string {
	return s.String()
}

func (s *ListKeywordsResponse) SetHeaders(v map[string]*string) *ListKeywordsResponse {
	s.Headers = v
	return s
}

func (s *ListKeywordsResponse) SetStatusCode(v int32) *ListKeywordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKeywordsResponse) SetBody(v *ListKeywordsResponseBody) *ListKeywordsResponse {
	s.Body = v
	return s
}

type ListOssCheckResultRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 2
	FinishNum *int64 `json:"FinishNum,omitempty" xml:"FinishNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// {"TaskId":"P_11TL5T"}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sort     map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListOssCheckResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOssCheckResultRequest) GoString() string {
	return s.String()
}

func (s *ListOssCheckResultRequest) SetCurrentPage(v int32) *ListOssCheckResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOssCheckResultRequest) SetEndDate(v string) *ListOssCheckResultRequest {
	s.EndDate = &v
	return s
}

func (s *ListOssCheckResultRequest) SetFinishNum(v int64) *ListOssCheckResultRequest {
	s.FinishNum = &v
	return s
}

func (s *ListOssCheckResultRequest) SetPageSize(v int32) *ListOssCheckResultRequest {
	s.PageSize = &v
	return s
}

func (s *ListOssCheckResultRequest) SetQuery(v string) *ListOssCheckResultRequest {
	s.Query = &v
	return s
}

func (s *ListOssCheckResultRequest) SetRegionId(v string) *ListOssCheckResultRequest {
	s.RegionId = &v
	return s
}

func (s *ListOssCheckResultRequest) SetSort(v map[string]*string) *ListOssCheckResultRequest {
	s.Sort = v
	return s
}

func (s *ListOssCheckResultRequest) SetStartDate(v string) *ListOssCheckResultRequest {
	s.StartDate = &v
	return s
}

func (s *ListOssCheckResultRequest) SetStatus(v int32) *ListOssCheckResultRequest {
	s.Status = &v
	return s
}

type ListOssCheckResultShrinkRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 2
	FinishNum *int64 `json:"FinishNum,omitempty" xml:"FinishNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// {"TaskId":"P_11TL5T"}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListOssCheckResultShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOssCheckResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListOssCheckResultShrinkRequest) SetCurrentPage(v int32) *ListOssCheckResultShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) SetEndDate(v string) *ListOssCheckResultShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) SetFinishNum(v int64) *ListOssCheckResultShrinkRequest {
	s.FinishNum = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) SetPageSize(v int32) *ListOssCheckResultShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) SetQuery(v string) *ListOssCheckResultShrinkRequest {
	s.Query = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) SetRegionId(v string) *ListOssCheckResultShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) SetSortShrink(v string) *ListOssCheckResultShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) SetStartDate(v string) *ListOssCheckResultShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) SetStatus(v int32) *ListOssCheckResultShrinkRequest {
	s.Status = &v
	return s
}

type ListOssCheckResultResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*ListOssCheckResultResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 13
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOssCheckResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOssCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListOssCheckResultResponseBody) SetCurrentPage(v int32) *ListOssCheckResultResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListOssCheckResultResponseBody) SetItems(v []*ListOssCheckResultResponseBodyItems) *ListOssCheckResultResponseBody {
	s.Items = v
	return s
}

func (s *ListOssCheckResultResponseBody) SetPageSize(v int32) *ListOssCheckResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOssCheckResultResponseBody) SetRequestId(v string) *ListOssCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOssCheckResultResponseBody) SetTotalCount(v int64) *ListOssCheckResultResponseBody {
	s.TotalCount = &v
	return s
}

type ListOssCheckResultResponseBodyItems struct {
	// example:
	//
	// tmp
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// audio
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// audio_media_detection
	CopyFrom     *string `json:"CopyFrom,omitempty" xml:"CopyFrom,omitempty"`
	FreezeStatus *string `json:"FreezeStatus,omitempty" xml:"FreezeStatus,omitempty"`
	FreezeType   *string `json:"FreezeType,omitempty" xml:"FreezeType,omitempty"`
	// example:
	//
	// http://www.aliyuncs.com/test.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// true
	IsCopy *bool `json:"IsCopy,omitempty" xml:"IsCopy,omitempty"`
	// example:
	//
	// dhT20X2310
	JobName *string   `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Labels  []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Labels2 []*string `json:"Labels2,omitempty" xml:"Labels2,omitempty" type:"Repeated"`
	// example:
	//
	// 54416c9b159df4a60ae03c04ccb94cb5
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 1713014531569_958.png.jpeg
	Object     *string `json:"Object,omitempty" xml:"Object,omitempty"`
	RiskLevel  *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	RiskLevel0 *string `json:"RiskLevel0,omitempty" xml:"RiskLevel0,omitempty"`
	RiskLevel2 *string `json:"RiskLevel2,omitempty" xml:"RiskLevel2,omitempty"`
	// example:
	//
	// {}
	ScanResult *string `json:"ScanResult,omitempty" xml:"ScanResult,omitempty"`
	// example:
	//
	// audio_media_detection_01
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// example:
	//
	// EP6TI7_au_Zo25ITvCbkocNuF801QOQX
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// http://www.aliyuncs.com/test.mp3
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListOssCheckResultResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListOssCheckResultResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListOssCheckResultResponseBodyItems) SetBucket(v string) *ListOssCheckResultResponseBodyItems {
	s.Bucket = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetCode(v string) *ListOssCheckResultResponseBodyItems {
	s.Code = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetContentType(v string) *ListOssCheckResultResponseBodyItems {
	s.ContentType = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetCopyFrom(v string) *ListOssCheckResultResponseBodyItems {
	s.CopyFrom = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetFreezeStatus(v string) *ListOssCheckResultResponseBodyItems {
	s.FreezeStatus = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetFreezeType(v string) *ListOssCheckResultResponseBodyItems {
	s.FreezeType = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetImageUrl(v string) *ListOssCheckResultResponseBodyItems {
	s.ImageUrl = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetIsCopy(v bool) *ListOssCheckResultResponseBodyItems {
	s.IsCopy = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetJobName(v string) *ListOssCheckResultResponseBodyItems {
	s.JobName = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetLabels(v []*string) *ListOssCheckResultResponseBodyItems {
	s.Labels = v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetLabels2(v []*string) *ListOssCheckResultResponseBodyItems {
	s.Labels2 = v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetMd5(v string) *ListOssCheckResultResponseBodyItems {
	s.Md5 = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetMsg(v string) *ListOssCheckResultResponseBodyItems {
	s.Msg = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetObject(v string) *ListOssCheckResultResponseBodyItems {
	s.Object = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetRiskLevel(v string) *ListOssCheckResultResponseBodyItems {
	s.RiskLevel = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetRiskLevel0(v string) *ListOssCheckResultResponseBodyItems {
	s.RiskLevel0 = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetRiskLevel2(v string) *ListOssCheckResultResponseBodyItems {
	s.RiskLevel2 = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetScanResult(v string) *ListOssCheckResultResponseBodyItems {
	s.ScanResult = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetServiceCode(v string) *ListOssCheckResultResponseBodyItems {
	s.ServiceCode = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetServiceName(v string) *ListOssCheckResultResponseBodyItems {
	s.ServiceName = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetTaskId(v string) *ListOssCheckResultResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetUrl(v string) *ListOssCheckResultResponseBodyItems {
	s.Url = &v
	return s
}

type ListOssCheckResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOssCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOssCheckResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOssCheckResultResponse) GoString() string {
	return s.String()
}

func (s *ListOssCheckResultResponse) SetHeaders(v map[string]*string) *ListOssCheckResultResponse {
	s.Headers = v
	return s
}

func (s *ListOssCheckResultResponse) SetStatusCode(v int32) *ListOssCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOssCheckResultResponse) SetBody(v *ListOssCheckResultResponseBody) *ListOssCheckResultResponse {
	s.Body = v
	return s
}

type ListServiceConfigsRequest struct {
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	UseStatus    *string `json:"UseStatus,omitempty" xml:"UseStatus,omitempty"`
}

func (s ListServiceConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceConfigsRequest) SetClassify(v string) *ListServiceConfigsRequest {
	s.Classify = &v
	return s
}

func (s *ListServiceConfigsRequest) SetRegionId(v string) *ListServiceConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceConfigsRequest) SetResourceType(v string) *ListServiceConfigsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListServiceConfigsRequest) SetUseStatus(v string) *ListServiceConfigsRequest {
	s.UseStatus = &v
	return s
}

type ListServiceConfigsResponseBody struct {
	// example:
	//
	// 400
	Code *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListServiceConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListServiceConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceConfigsResponseBody) SetCode(v int32) *ListServiceConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListServiceConfigsResponseBody) SetData(v []*ListServiceConfigsResponseBodyData) *ListServiceConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListServiceConfigsResponseBody) SetMsg(v string) *ListServiceConfigsResponseBody {
	s.Msg = &v
	return s
}

func (s *ListServiceConfigsResponseBody) SetRequestId(v string) *ListServiceConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceConfigsResponseBody) SetSuccess(v bool) *ListServiceConfigsResponseBody {
	s.Success = &v
	return s
}

type ListServiceConfigsResponseBodyData struct {
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// example:
	//
	// nickname_detection
	CopyFrom          *string                                              `json:"CopyFrom,omitempty" xml:"CopyFrom,omitempty"`
	CustomServiceConf *ListServiceConfigsResponseBodyDataCustomServiceConf `json:"CustomServiceConf,omitempty" xml:"CustomServiceConf,omitempty" type:"Struct"`
	// example:
	//
	// 2023-07-11 15:40:04
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// {}
	Option map[string]interface{} `json:"Option,omitempty" xml:"Option,omitempty"`
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// nickname_detection
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ServiceDesc *string `json:"ServiceDesc,omitempty" xml:"ServiceDesc,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// example:
	//
	// plus
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// UID。
	//
	// example:
	//
	// 1674*****0071291
	Uid       *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UseStatus *string `json:"UseStatus,omitempty" xml:"UseStatus,omitempty"`
}

func (s ListServiceConfigsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListServiceConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServiceConfigsResponseBodyData) SetClassify(v string) *ListServiceConfigsResponseBodyData {
	s.Classify = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetCopyFrom(v string) *ListServiceConfigsResponseBodyData {
	s.CopyFrom = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetCustomServiceConf(v *ListServiceConfigsResponseBodyDataCustomServiceConf) *ListServiceConfigsResponseBodyData {
	s.CustomServiceConf = v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetGmtModified(v string) *ListServiceConfigsResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetOption(v map[string]interface{}) *ListServiceConfigsResponseBodyData {
	s.Option = v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetResourceType(v string) *ListServiceConfigsResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetServiceCode(v string) *ListServiceConfigsResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetServiceDesc(v string) *ListServiceConfigsResponseBodyData {
	s.ServiceDesc = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetServiceName(v string) *ListServiceConfigsResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetServiceType(v string) *ListServiceConfigsResponseBodyData {
	s.ServiceType = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetUid(v string) *ListServiceConfigsResponseBodyData {
	s.Uid = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetUseStatus(v string) *ListServiceConfigsResponseBodyData {
	s.UseStatus = &v
	return s
}

type ListServiceConfigsResponseBodyDataCustomServiceConf struct {
	// example:
	//
	// audio_media_detection
	AudioService       *string                                                     `json:"AudioService,omitempty" xml:"AudioService,omitempty"`
	ImageService       []*string                                                   `json:"ImageService,omitempty" xml:"ImageService,omitempty" type:"Repeated"`
	KeywordFilterLibs  []*string                                                   `json:"KeywordFilterLibs,omitempty" xml:"KeywordFilterLibs,omitempty" type:"Repeated"`
	KeywordHitLibs     []*string                                                   `json:"KeywordHitLibs,omitempty" xml:"KeywordHitLibs,omitempty" type:"Repeated"`
	Rules              []*ListServiceConfigsResponseBodyDataCustomServiceConfRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	SimilarTextHitLibs []*string                                                   `json:"SimilarTextHitLibs,omitempty" xml:"SimilarTextHitLibs,omitempty" type:"Repeated"`
}

func (s ListServiceConfigsResponseBodyDataCustomServiceConf) String() string {
	return tea.Prettify(s)
}

func (s ListServiceConfigsResponseBodyDataCustomServiceConf) GoString() string {
	return s.String()
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) SetAudioService(v string) *ListServiceConfigsResponseBodyDataCustomServiceConf {
	s.AudioService = &v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) SetImageService(v []*string) *ListServiceConfigsResponseBodyDataCustomServiceConf {
	s.ImageService = v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) SetKeywordFilterLibs(v []*string) *ListServiceConfigsResponseBodyDataCustomServiceConf {
	s.KeywordFilterLibs = v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) SetKeywordHitLibs(v []*string) *ListServiceConfigsResponseBodyDataCustomServiceConf {
	s.KeywordHitLibs = v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) SetRules(v []*ListServiceConfigsResponseBodyDataCustomServiceConfRules) *ListServiceConfigsResponseBodyDataCustomServiceConf {
	s.Rules = v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) SetSimilarTextHitLibs(v []*string) *ListServiceConfigsResponseBodyDataCustomServiceConf {
	s.SimilarTextHitLibs = v
	return s
}

type ListServiceConfigsResponseBodyDataCustomServiceConfRules struct {
	ImageScanRule *ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule `json:"ImageScanRule,omitempty" xml:"ImageScanRule,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Index        *int32                                                                `json:"Index,omitempty" xml:"Index,omitempty"`
	TextScanRule *ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule `json:"TextScanRule,omitempty" xml:"TextScanRule,omitempty" type:"Struct"`
}

func (s ListServiceConfigsResponseBodyDataCustomServiceConfRules) String() string {
	return tea.Prettify(s)
}

func (s ListServiceConfigsResponseBodyDataCustomServiceConfRules) GoString() string {
	return s.String()
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRules) SetImageScanRule(v *ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule) *ListServiceConfigsResponseBodyDataCustomServiceConfRules {
	s.ImageScanRule = v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRules) SetIndex(v int32) *ListServiceConfigsResponseBodyDataCustomServiceConfRules {
	s.Index = &v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRules) SetTextScanRule(v *ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule) *ListServiceConfigsResponseBodyDataCustomServiceConfRules {
	s.TextScanRule = v
	return s
}

type ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule struct {
	Services []*string `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule) String() string {
	return tea.Prettify(s)
}

func (s ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule) GoString() string {
	return s.String()
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule) SetServices(v []*string) *ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule {
	s.Services = v
	return s
}

type ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule struct {
	Services []*string `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule) String() string {
	return tea.Prettify(s)
}

func (s ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule) GoString() string {
	return s.String()
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule) SetServices(v []*string) *ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule {
	s.Services = v
	return s
}

type ListServiceConfigsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceConfigsResponse) SetHeaders(v map[string]*string) *ListServiceConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceConfigsResponse) SetStatusCode(v int32) *ListServiceConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceConfigsResponse) SetBody(v *ListServiceConfigsResponseBody) *ListServiceConfigsResponse {
	s.Body = v
	return s
}

type ModifyAnswerLibRequest struct {
	// example:
	//
	// custom_xxxx
	LibId   *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAnswerLibRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAnswerLibRequest) GoString() string {
	return s.String()
}

func (s *ModifyAnswerLibRequest) SetLibId(v string) *ModifyAnswerLibRequest {
	s.LibId = &v
	return s
}

func (s *ModifyAnswerLibRequest) SetLibName(v string) *ModifyAnswerLibRequest {
	s.LibName = &v
	return s
}

func (s *ModifyAnswerLibRequest) SetRegionId(v string) *ModifyAnswerLibRequest {
	s.RegionId = &v
	return s
}

type ModifyAnswerLibResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAnswerLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAnswerLibResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAnswerLibResponseBody) SetData(v bool) *ModifyAnswerLibResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyAnswerLibResponseBody) SetRequestId(v string) *ModifyAnswerLibResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAnswerLibResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAnswerLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAnswerLibResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAnswerLibResponse) GoString() string {
	return s.String()
}

func (s *ModifyAnswerLibResponse) SetHeaders(v map[string]*string) *ModifyAnswerLibResponse {
	s.Headers = v
	return s
}

func (s *ModifyAnswerLibResponse) SetStatusCode(v int32) *ModifyAnswerLibResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAnswerLibResponse) SetBody(v *ModifyAnswerLibResponseBody) *ModifyAnswerLibResponse {
	s.Body = v
	return s
}

type ModifyCallbackRequest struct {
	// example:
	//
	// SHA256
	CryptType *string `json:"CryptType,omitempty" xml:"CryptType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 112
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// all
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// https://www.aliyuncs.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ModifyCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCallbackRequest) GoString() string {
	return s.String()
}

func (s *ModifyCallbackRequest) SetCryptType(v string) *ModifyCallbackRequest {
	s.CryptType = &v
	return s
}

func (s *ModifyCallbackRequest) SetId(v int64) *ModifyCallbackRequest {
	s.Id = &v
	return s
}

func (s *ModifyCallbackRequest) SetName(v string) *ModifyCallbackRequest {
	s.Name = &v
	return s
}

func (s *ModifyCallbackRequest) SetRegionId(v string) *ModifyCallbackRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCallbackRequest) SetScope(v string) *ModifyCallbackRequest {
	s.Scope = &v
	return s
}

func (s *ModifyCallbackRequest) SetUrl(v string) *ModifyCallbackRequest {
	s.Url = &v
	return s
}

type ModifyCallbackResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCallbackResponseBody) SetData(v bool) *ModifyCallbackResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyCallbackResponseBody) SetRequestId(v string) *ModifyCallbackResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCallbackResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCallbackResponse) GoString() string {
	return s.String()
}

func (s *ModifyCallbackResponse) SetHeaders(v map[string]*string) *ModifyCallbackResponse {
	s.Headers = v
	return s
}

func (s *ModifyCallbackResponse) SetStatusCode(v int32) *ModifyCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCallbackResponse) SetBody(v *ModifyCallbackResponseBody) *ModifyCallbackResponse {
	s.Body = v
	return s
}

type ModifyFeatureConfigRequest struct {
	// Configuration, in JSON format
	//
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Label meaning
	//
	// example:
	//
	// 标签2
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Label value, customer-defined
	//
	// example:
	//
	// __config__
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// Region
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource type.
	//
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Service code.
	//
	// example:
	//
	// llm_query_moderation
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Type
	//
	// example:
	//
	// custom_llm_template
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyFeatureConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFeatureConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyFeatureConfigRequest) SetConfig(v string) *ModifyFeatureConfigRequest {
	s.Config = &v
	return s
}

func (s *ModifyFeatureConfigRequest) SetDescription(v string) *ModifyFeatureConfigRequest {
	s.Description = &v
	return s
}

func (s *ModifyFeatureConfigRequest) SetField(v string) *ModifyFeatureConfigRequest {
	s.Field = &v
	return s
}

func (s *ModifyFeatureConfigRequest) SetRegionId(v string) *ModifyFeatureConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFeatureConfigRequest) SetResourceType(v string) *ModifyFeatureConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyFeatureConfigRequest) SetServiceCode(v string) *ModifyFeatureConfigRequest {
	s.ServiceCode = &v
	return s
}

func (s *ModifyFeatureConfigRequest) SetType(v string) *ModifyFeatureConfigRequest {
	s.Type = &v
	return s
}

type ModifyFeatureConfigResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data
	//
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message of this request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// ID assigned by the backend to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success indicator.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyFeatureConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFeatureConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFeatureConfigResponseBody) SetCode(v int32) *ModifyFeatureConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyFeatureConfigResponseBody) SetData(v bool) *ModifyFeatureConfigResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyFeatureConfigResponseBody) SetHttpStatusCode(v int32) *ModifyFeatureConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyFeatureConfigResponseBody) SetMsg(v string) *ModifyFeatureConfigResponseBody {
	s.Msg = &v
	return s
}

func (s *ModifyFeatureConfigResponseBody) SetRequestId(v string) *ModifyFeatureConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFeatureConfigResponseBody) SetSuccess(v bool) *ModifyFeatureConfigResponseBody {
	s.Success = &v
	return s
}

type ModifyFeatureConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFeatureConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFeatureConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFeatureConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyFeatureConfigResponse) SetHeaders(v map[string]*string) *ModifyFeatureConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyFeatureConfigResponse) SetStatusCode(v int32) *ModifyFeatureConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFeatureConfigResponse) SetBody(v *ModifyFeatureConfigResponseBody) *ModifyFeatureConfigResponse {
	s.Body = v
	return s
}

type ModifyServiceInfoRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ServiceDesc *string `json:"ServiceDesc,omitempty" xml:"ServiceDesc,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ModifyServiceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyServiceInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyServiceInfoRequest) SetRegionId(v string) *ModifyServiceInfoRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyServiceInfoRequest) SetResourceType(v string) *ModifyServiceInfoRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyServiceInfoRequest) SetServiceCode(v string) *ModifyServiceInfoRequest {
	s.ServiceCode = &v
	return s
}

func (s *ModifyServiceInfoRequest) SetServiceDesc(v string) *ModifyServiceInfoRequest {
	s.ServiceDesc = &v
	return s
}

func (s *ModifyServiceInfoRequest) SetServiceName(v string) *ModifyServiceInfoRequest {
	s.ServiceName = &v
	return s
}

type ModifyServiceInfoResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyServiceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyServiceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyServiceInfoResponseBody) SetData(v bool) *ModifyServiceInfoResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyServiceInfoResponseBody) SetRequestId(v string) *ModifyServiceInfoResponseBody {
	s.RequestId = &v
	return s
}

type ModifyServiceInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyServiceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyServiceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyServiceInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyServiceInfoResponse) SetHeaders(v map[string]*string) *ModifyServiceInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyServiceInfoResponse) SetStatusCode(v int32) *ModifyServiceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyServiceInfoResponse) SetBody(v *ModifyServiceInfoResponseBody) *ModifyServiceInfoResponse {
	s.Body = v
	return s
}

type QueryAnswerSampleByPageRequest struct {
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// custom_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sort     map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s QueryAnswerSampleByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAnswerSampleByPageRequest) GoString() string {
	return s.String()
}

func (s *QueryAnswerSampleByPageRequest) SetAnswer(v string) *QueryAnswerSampleByPageRequest {
	s.Answer = &v
	return s
}

func (s *QueryAnswerSampleByPageRequest) SetCurrentPage(v int32) *QueryAnswerSampleByPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryAnswerSampleByPageRequest) SetLibId(v string) *QueryAnswerSampleByPageRequest {
	s.LibId = &v
	return s
}

func (s *QueryAnswerSampleByPageRequest) SetPageSize(v int32) *QueryAnswerSampleByPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAnswerSampleByPageRequest) SetRegionId(v string) *QueryAnswerSampleByPageRequest {
	s.RegionId = &v
	return s
}

func (s *QueryAnswerSampleByPageRequest) SetSort(v map[string]*string) *QueryAnswerSampleByPageRequest {
	s.Sort = v
	return s
}

type QueryAnswerSampleByPageShrinkRequest struct {
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// custom_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s QueryAnswerSampleByPageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAnswerSampleByPageShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryAnswerSampleByPageShrinkRequest) SetAnswer(v string) *QueryAnswerSampleByPageShrinkRequest {
	s.Answer = &v
	return s
}

func (s *QueryAnswerSampleByPageShrinkRequest) SetCurrentPage(v int32) *QueryAnswerSampleByPageShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryAnswerSampleByPageShrinkRequest) SetLibId(v string) *QueryAnswerSampleByPageShrinkRequest {
	s.LibId = &v
	return s
}

func (s *QueryAnswerSampleByPageShrinkRequest) SetPageSize(v int32) *QueryAnswerSampleByPageShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAnswerSampleByPageShrinkRequest) SetRegionId(v string) *QueryAnswerSampleByPageShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *QueryAnswerSampleByPageShrinkRequest) SetSortShrink(v string) *QueryAnswerSampleByPageShrinkRequest {
	s.SortShrink = &v
	return s
}

type QueryAnswerSampleByPageResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*QueryAnswerSampleByPageResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 58
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryAnswerSampleByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAnswerSampleByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAnswerSampleByPageResponseBody) SetCurrentPage(v int32) *QueryAnswerSampleByPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryAnswerSampleByPageResponseBody) SetItems(v []*QueryAnswerSampleByPageResponseBodyItems) *QueryAnswerSampleByPageResponseBody {
	s.Items = v
	return s
}

func (s *QueryAnswerSampleByPageResponseBody) SetPageSize(v int32) *QueryAnswerSampleByPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryAnswerSampleByPageResponseBody) SetRequestId(v string) *QueryAnswerSampleByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAnswerSampleByPageResponseBody) SetTotalCount(v int64) *QueryAnswerSampleByPageResponseBody {
	s.TotalCount = &v
	return s
}

type QueryAnswerSampleByPageResponseBodyItems struct {
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// example:
	//
	// 2023-07-31 06:16:06
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1666
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// custom_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// UID。
	//
	// example:
	//
	// 104813*****2399
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s QueryAnswerSampleByPageResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s QueryAnswerSampleByPageResponseBodyItems) GoString() string {
	return s.String()
}

func (s *QueryAnswerSampleByPageResponseBodyItems) SetAnswer(v string) *QueryAnswerSampleByPageResponseBodyItems {
	s.Answer = &v
	return s
}

func (s *QueryAnswerSampleByPageResponseBodyItems) SetGmtCreate(v string) *QueryAnswerSampleByPageResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *QueryAnswerSampleByPageResponseBodyItems) SetId(v int64) *QueryAnswerSampleByPageResponseBodyItems {
	s.Id = &v
	return s
}

func (s *QueryAnswerSampleByPageResponseBodyItems) SetLibId(v string) *QueryAnswerSampleByPageResponseBodyItems {
	s.LibId = &v
	return s
}

func (s *QueryAnswerSampleByPageResponseBodyItems) SetUid(v string) *QueryAnswerSampleByPageResponseBodyItems {
	s.Uid = &v
	return s
}

type QueryAnswerSampleByPageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAnswerSampleByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAnswerSampleByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAnswerSampleByPageResponse) GoString() string {
	return s.String()
}

func (s *QueryAnswerSampleByPageResponse) SetHeaders(v map[string]*string) *QueryAnswerSampleByPageResponse {
	s.Headers = v
	return s
}

func (s *QueryAnswerSampleByPageResponse) SetStatusCode(v int32) *QueryAnswerSampleByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAnswerSampleByPageResponse) SetBody(v *QueryAnswerSampleByPageResponseBody) *QueryAnswerSampleByPageResponse {
	s.Body = v
	return s
}

type QueryCallbackRequest struct {
	// example:
	//
	// true
	CheckForOss *bool `json:"CheckForOss,omitempty" xml:"CheckForOss,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCallbackRequest) GoString() string {
	return s.String()
}

func (s *QueryCallbackRequest) SetCheckForOss(v bool) *QueryCallbackRequest {
	s.CheckForOss = &v
	return s
}

func (s *QueryCallbackRequest) SetId(v int64) *QueryCallbackRequest {
	s.Id = &v
	return s
}

func (s *QueryCallbackRequest) SetRegionId(v string) *QueryCallbackRequest {
	s.RegionId = &v
	return s
}

type QueryCallbackResponseBody struct {
	// example:
	//
	// SHA256
	CryptType *string `json:"CryptType,omitempty" xml:"CryptType,omitempty"`
	// example:
	//
	// false
	ExistsOssCheckTask *bool `json:"ExistsOssCheckTask,omitempty" xml:"ExistsOssCheckTask,omitempty"`
	// example:
	//
	// 2022-11-30 16:30:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-06-03 15:20:14
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 11234
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// all
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// Seed。
	//
	// example:
	//
	// cb2MysbJTAAIf6gB3u4vpIEU-1ySnnf
	Seed *string `json:"Seed,omitempty" xml:"Seed,omitempty"`
	// UID。
	//
	// example:
	//
	// 19964*****086772
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// example:
	//
	// https://www.aliyuncs.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallbackResponseBody) SetCryptType(v string) *QueryCallbackResponseBody {
	s.CryptType = &v
	return s
}

func (s *QueryCallbackResponseBody) SetExistsOssCheckTask(v bool) *QueryCallbackResponseBody {
	s.ExistsOssCheckTask = &v
	return s
}

func (s *QueryCallbackResponseBody) SetGmtCreate(v string) *QueryCallbackResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *QueryCallbackResponseBody) SetGmtModified(v string) *QueryCallbackResponseBody {
	s.GmtModified = &v
	return s
}

func (s *QueryCallbackResponseBody) SetId(v int64) *QueryCallbackResponseBody {
	s.Id = &v
	return s
}

func (s *QueryCallbackResponseBody) SetName(v string) *QueryCallbackResponseBody {
	s.Name = &v
	return s
}

func (s *QueryCallbackResponseBody) SetRequestId(v string) *QueryCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCallbackResponseBody) SetScope(v string) *QueryCallbackResponseBody {
	s.Scope = &v
	return s
}

func (s *QueryCallbackResponseBody) SetSeed(v string) *QueryCallbackResponseBody {
	s.Seed = &v
	return s
}

func (s *QueryCallbackResponseBody) SetUid(v string) *QueryCallbackResponseBody {
	s.Uid = &v
	return s
}

func (s *QueryCallbackResponseBody) SetUrl(v string) *QueryCallbackResponseBody {
	s.Url = &v
	return s
}

type QueryCallbackResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCallbackResponse) GoString() string {
	return s.String()
}

func (s *QueryCallbackResponse) SetHeaders(v map[string]*string) *QueryCallbackResponse {
	s.Headers = v
	return s
}

func (s *QueryCallbackResponse) SetStatusCode(v int32) *QueryCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCallbackResponse) SetBody(v *QueryCallbackResponseBody) *QueryCallbackResponse {
	s.Body = v
	return s
}

type QueryCallbackByPageRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryCallbackByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCallbackByPageRequest) GoString() string {
	return s.String()
}

func (s *QueryCallbackByPageRequest) SetCurrentPage(v int32) *QueryCallbackByPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryCallbackByPageRequest) SetPageSize(v int32) *QueryCallbackByPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCallbackByPageRequest) SetRegionId(v string) *QueryCallbackByPageRequest {
	s.RegionId = &v
	return s
}

type QueryCallbackByPageResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*QueryCallbackByPageResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryCallbackByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCallbackByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallbackByPageResponseBody) SetCurrentPage(v int32) *QueryCallbackByPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryCallbackByPageResponseBody) SetItems(v []*QueryCallbackByPageResponseBodyItems) *QueryCallbackByPageResponseBody {
	s.Items = v
	return s
}

func (s *QueryCallbackByPageResponseBody) SetPageSize(v int32) *QueryCallbackByPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryCallbackByPageResponseBody) SetRequestId(v string) *QueryCallbackByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCallbackByPageResponseBody) SetTotalCount(v int64) *QueryCallbackByPageResponseBody {
	s.TotalCount = &v
	return s
}

type QueryCallbackByPageResponseBodyItems struct {
	// example:
	//
	// SHA256
	CryptType *string `json:"CryptType,omitempty" xml:"CryptType,omitempty"`
	// example:
	//
	// 2024-06-03 15:20:14
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-06-03 15:20:14
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1697
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// all
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// Seed。
	//
	// example:
	//
	// cb6gYS8GXj4Vn4Y4FN0Y8R5M-1x46Mq
	Seed *string `json:"Seed,omitempty" xml:"Seed,omitempty"`
	// UID。
	//
	// example:
	//
	// 12161*****398900
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// example:
	//
	// https://console.aliyun.com/
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryCallbackByPageResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s QueryCallbackByPageResponseBodyItems) GoString() string {
	return s.String()
}

func (s *QueryCallbackByPageResponseBodyItems) SetCryptType(v string) *QueryCallbackByPageResponseBodyItems {
	s.CryptType = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) SetGmtCreate(v string) *QueryCallbackByPageResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) SetGmtModified(v string) *QueryCallbackByPageResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) SetId(v int64) *QueryCallbackByPageResponseBodyItems {
	s.Id = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) SetName(v string) *QueryCallbackByPageResponseBodyItems {
	s.Name = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) SetScope(v string) *QueryCallbackByPageResponseBodyItems {
	s.Scope = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) SetSeed(v string) *QueryCallbackByPageResponseBodyItems {
	s.Seed = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) SetUid(v string) *QueryCallbackByPageResponseBodyItems {
	s.Uid = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) SetUrl(v string) *QueryCallbackByPageResponseBodyItems {
	s.Url = &v
	return s
}

type QueryCallbackByPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCallbackByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCallbackByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCallbackByPageResponse) GoString() string {
	return s.String()
}

func (s *QueryCallbackByPageResponse) SetHeaders(v map[string]*string) *QueryCallbackByPageResponse {
	s.Headers = v
	return s
}

func (s *QueryCallbackByPageResponse) SetStatusCode(v int32) *QueryCallbackByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCallbackByPageResponse) SetBody(v *QueryCallbackByPageResponseBody) *QueryCallbackByPageResponse {
	s.Body = v
	return s
}

type UpdateBackupConfigRequest struct {
	// example:
	//
	// {}
	BackupConfig *string `json:"BackupConfig,omitempty" xml:"BackupConfig,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// video
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// videoDetection
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s UpdateBackupConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateBackupConfigRequest) SetBackupConfig(v string) *UpdateBackupConfigRequest {
	s.BackupConfig = &v
	return s
}

func (s *UpdateBackupConfigRequest) SetRegionId(v string) *UpdateBackupConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBackupConfigRequest) SetResourceType(v string) *UpdateBackupConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateBackupConfigRequest) SetServiceCode(v string) *UpdateBackupConfigRequest {
	s.ServiceCode = &v
	return s
}

type UpdateBackupConfigResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBackupConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBackupConfigResponseBody) SetData(v bool) *UpdateBackupConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateBackupConfigResponseBody) SetRequestId(v string) *UpdateBackupConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateBackupConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBackupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBackupConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBackupConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateBackupConfigResponse) SetHeaders(v map[string]*string) *UpdateBackupConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateBackupConfigResponse) SetStatusCode(v int32) *UpdateBackupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBackupConfigResponse) SetBody(v *UpdateBackupConfigResponseBody) *UpdateBackupConfigResponse {
	s.Body = v
	return s
}

type UpdateImageLibRequest struct {
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 0
	FreeInspection *int32 `json:"FreeInspection,omitempty" xml:"FreeInspection,omitempty"`
	// example:
	//
	// custom_xxxx
	LibId   *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateImageLibRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageLibRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageLibRequest) SetComment(v string) *UpdateImageLibRequest {
	s.Comment = &v
	return s
}

func (s *UpdateImageLibRequest) SetFreeInspection(v int32) *UpdateImageLibRequest {
	s.FreeInspection = &v
	return s
}

func (s *UpdateImageLibRequest) SetLibId(v string) *UpdateImageLibRequest {
	s.LibId = &v
	return s
}

func (s *UpdateImageLibRequest) SetLibName(v string) *UpdateImageLibRequest {
	s.LibName = &v
	return s
}

func (s *UpdateImageLibRequest) SetRegionId(v string) *UpdateImageLibRequest {
	s.RegionId = &v
	return s
}

type UpdateImageLibResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateImageLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageLibResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageLibResponseBody) SetCode(v int32) *UpdateImageLibResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateImageLibResponseBody) SetData(v bool) *UpdateImageLibResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateImageLibResponseBody) SetHttpStatusCode(v int32) *UpdateImageLibResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateImageLibResponseBody) SetMsg(v string) *UpdateImageLibResponseBody {
	s.Msg = &v
	return s
}

func (s *UpdateImageLibResponseBody) SetRequestId(v string) *UpdateImageLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageLibResponseBody) SetSuccess(v bool) *UpdateImageLibResponseBody {
	s.Success = &v
	return s
}

type UpdateImageLibResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateImageLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateImageLibResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageLibResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageLibResponse) SetHeaders(v map[string]*string) *UpdateImageLibResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageLibResponse) SetStatusCode(v int32) *UpdateImageLibResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageLibResponse) SetBody(v *UpdateImageLibResponseBody) *UpdateImageLibResponse {
	s.Body = v
	return s
}

type UpdateImageLibFreeInspectionRequest struct {
	Config map[string]*int32 `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateImageLibFreeInspectionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageLibFreeInspectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageLibFreeInspectionRequest) SetConfig(v map[string]*int32) *UpdateImageLibFreeInspectionRequest {
	s.Config = v
	return s
}

func (s *UpdateImageLibFreeInspectionRequest) SetRegionId(v string) *UpdateImageLibFreeInspectionRequest {
	s.RegionId = &v
	return s
}

type UpdateImageLibFreeInspectionShrinkRequest struct {
	ConfigShrink *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateImageLibFreeInspectionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageLibFreeInspectionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageLibFreeInspectionShrinkRequest) SetConfigShrink(v string) *UpdateImageLibFreeInspectionShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *UpdateImageLibFreeInspectionShrinkRequest) SetRegionId(v string) *UpdateImageLibFreeInspectionShrinkRequest {
	s.RegionId = &v
	return s
}

type UpdateImageLibFreeInspectionResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateImageLibFreeInspectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageLibFreeInspectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageLibFreeInspectionResponseBody) SetCode(v int32) *UpdateImageLibFreeInspectionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateImageLibFreeInspectionResponseBody) SetData(v bool) *UpdateImageLibFreeInspectionResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateImageLibFreeInspectionResponseBody) SetHttpStatusCode(v int32) *UpdateImageLibFreeInspectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateImageLibFreeInspectionResponseBody) SetMsg(v string) *UpdateImageLibFreeInspectionResponseBody {
	s.Msg = &v
	return s
}

func (s *UpdateImageLibFreeInspectionResponseBody) SetRequestId(v string) *UpdateImageLibFreeInspectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageLibFreeInspectionResponseBody) SetSuccess(v bool) *UpdateImageLibFreeInspectionResponseBody {
	s.Success = &v
	return s
}

type UpdateImageLibFreeInspectionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateImageLibFreeInspectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateImageLibFreeInspectionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageLibFreeInspectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageLibFreeInspectionResponse) SetHeaders(v map[string]*string) *UpdateImageLibFreeInspectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageLibFreeInspectionResponse) SetStatusCode(v int32) *UpdateImageLibFreeInspectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageLibFreeInspectionResponse) SetBody(v *UpdateImageLibFreeInspectionResponseBody) *UpdateImageLibFreeInspectionResponse {
	s.Body = v
	return s
}

type UpdateKeywordLibRequest struct {
	// example:
	//
	// custom_xxxx
	LibId   *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateKeywordLibRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKeywordLibRequest) GoString() string {
	return s.String()
}

func (s *UpdateKeywordLibRequest) SetLibId(v string) *UpdateKeywordLibRequest {
	s.LibId = &v
	return s
}

func (s *UpdateKeywordLibRequest) SetLibName(v string) *UpdateKeywordLibRequest {
	s.LibName = &v
	return s
}

func (s *UpdateKeywordLibRequest) SetRegionId(v string) *UpdateKeywordLibRequest {
	s.RegionId = &v
	return s
}

type UpdateKeywordLibResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateKeywordLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateKeywordLibResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKeywordLibResponseBody) SetCode(v int32) *UpdateKeywordLibResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateKeywordLibResponseBody) SetData(v bool) *UpdateKeywordLibResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateKeywordLibResponseBody) SetMsg(v string) *UpdateKeywordLibResponseBody {
	s.Msg = &v
	return s
}

func (s *UpdateKeywordLibResponseBody) SetRequestId(v string) *UpdateKeywordLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKeywordLibResponseBody) SetSuccess(v bool) *UpdateKeywordLibResponseBody {
	s.Success = &v
	return s
}

type UpdateKeywordLibResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKeywordLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKeywordLibResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateKeywordLibResponse) GoString() string {
	return s.String()
}

func (s *UpdateKeywordLibResponse) SetHeaders(v map[string]*string) *UpdateKeywordLibResponse {
	s.Headers = v
	return s
}

func (s *UpdateKeywordLibResponse) SetStatusCode(v int32) *UpdateKeywordLibResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKeywordLibResponse) SetBody(v *UpdateKeywordLibResponseBody) *UpdateKeywordLibResponse {
	s.Body = v
	return s
}

type UpdateScanResultFeedbackRequest struct {
	// example:
	//
	// missOut
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// example:
	//
	// 46232656-984E-****-A648-B1D0667B6C3E
	QueryRequestId *string `json:"QueryRequestId,omitempty" xml:"QueryRequestId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s UpdateScanResultFeedbackRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateScanResultFeedbackRequest) GoString() string {
	return s.String()
}

func (s *UpdateScanResultFeedbackRequest) SetFeedback(v string) *UpdateScanResultFeedbackRequest {
	s.Feedback = &v
	return s
}

func (s *UpdateScanResultFeedbackRequest) SetQueryRequestId(v string) *UpdateScanResultFeedbackRequest {
	s.QueryRequestId = &v
	return s
}

func (s *UpdateScanResultFeedbackRequest) SetRegionId(v string) *UpdateScanResultFeedbackRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateScanResultFeedbackRequest) SetResourceType(v string) *UpdateScanResultFeedbackRequest {
	s.ResourceType = &v
	return s
}

type UpdateScanResultFeedbackResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateScanResultFeedbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateScanResultFeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScanResultFeedbackResponseBody) SetData(v bool) *UpdateScanResultFeedbackResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateScanResultFeedbackResponseBody) SetRequestId(v string) *UpdateScanResultFeedbackResponseBody {
	s.RequestId = &v
	return s
}

type UpdateScanResultFeedbackResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScanResultFeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScanResultFeedbackResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateScanResultFeedbackResponse) GoString() string {
	return s.String()
}

func (s *UpdateScanResultFeedbackResponse) SetHeaders(v map[string]*string) *UpdateScanResultFeedbackResponse {
	s.Headers = v
	return s
}

func (s *UpdateScanResultFeedbackResponse) SetStatusCode(v int32) *UpdateScanResultFeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScanResultFeedbackResponse) SetBody(v *UpdateScanResultFeedbackResponseBody) *UpdateScanResultFeedbackResponse {
	s.Body = v
	return s
}

type UpdateServiceConfigRequest struct {
	// example:
	//
	// {}
	FileConfig *string `json:"FileConfig,omitempty" xml:"FileConfig,omitempty"`
	// example:
	//
	// []
	KeywordFilterLibs *string `json:"KeywordFilterLibs,omitempty" xml:"KeywordFilterLibs,omitempty"`
	// example:
	//
	// []
	KeywordHitLibs *string `json:"KeywordHitLibs,omitempty" xml:"KeywordHitLibs,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// pornographic
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// example:
	//
	// {}
	SceneConfig *string `json:"SceneConfig,omitempty" xml:"SceneConfig,omitempty"`
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// {}
	VideoConfig *string `json:"VideoConfig,omitempty" xml:"VideoConfig,omitempty"`
}

func (s UpdateServiceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceConfigRequest) SetFileConfig(v string) *UpdateServiceConfigRequest {
	s.FileConfig = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetKeywordFilterLibs(v string) *UpdateServiceConfigRequest {
	s.KeywordFilterLibs = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetKeywordHitLibs(v string) *UpdateServiceConfigRequest {
	s.KeywordHitLibs = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetRegionId(v string) *UpdateServiceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetResourceType(v string) *UpdateServiceConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetScene(v string) *UpdateServiceConfigRequest {
	s.Scene = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetSceneConfig(v string) *UpdateServiceConfigRequest {
	s.SceneConfig = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetServiceCode(v string) *UpdateServiceConfigRequest {
	s.ServiceCode = &v
	return s
}

func (s *UpdateServiceConfigRequest) SetVideoConfig(v string) *UpdateServiceConfigRequest {
	s.VideoConfig = &v
	return s
}

type UpdateServiceConfigResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateServiceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceConfigResponseBody) SetCode(v int32) *UpdateServiceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateServiceConfigResponseBody) SetData(v bool) *UpdateServiceConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateServiceConfigResponseBody) SetHttpStatusCode(v int32) *UpdateServiceConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateServiceConfigResponseBody) SetMsg(v string) *UpdateServiceConfigResponseBody {
	s.Msg = &v
	return s
}

func (s *UpdateServiceConfigResponseBody) SetRequestId(v string) *UpdateServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceConfigResponseBody) SetSuccess(v bool) *UpdateServiceConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateServiceConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceConfigResponse) SetHeaders(v map[string]*string) *UpdateServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceConfigResponse) SetStatusCode(v int32) *UpdateServiceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceConfigResponse) SetBody(v *UpdateServiceConfigResponseBody) *UpdateServiceConfigResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"cn-chengdu":            tea.String("green.aliyuncs.com"),
		"cn-hongkong":           tea.String("green.aliyuncs.com"),
		"cn-huhehaote":          tea.String("green.aliyuncs.com"),
		"cn-qingdao":            tea.String("green.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("green.aliyuncs.com"),
		"eu-central-1":          tea.String("green.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("green.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             tea.String("green.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             tea.String("green.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("green.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("green.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("green.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("green.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("green"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加代答样本
//
// @param request - AddAnswerSampleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAnswerSampleResponse
func (client *Client) AddAnswerSampleWithOptions(request *AddAnswerSampleRequest, runtime *util.RuntimeOptions) (_result *AddAnswerSampleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		query["LibId"] = request.LibId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SampleObject)) {
		query["SampleObject"] = request.SampleObject
	}

	if !tea.BoolValue(util.IsUnset(request.Samples)) {
		query["Samples"] = request.Samples
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddAnswerSample"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddAnswerSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加代答样本
//
// @param request - AddAnswerSampleRequest
//
// @return AddAnswerSampleResponse
func (client *Client) AddAnswerSample(request *AddAnswerSampleRequest) (_result *AddAnswerSampleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddAnswerSampleResponse{}
	_body, _err := client.AddAnswerSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建图库
//
// @param request - AddImageLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddImageLibResponse
func (client *Client) AddImageLibWithOptions(request *AddImageLibRequest, runtime *util.RuntimeOptions) (_result *AddImageLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		body["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.LibName)) {
		body["LibName"] = request.LibName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddImageLib"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddImageLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建图库
//
// @param request - AddImageLibRequest
//
// @return AddImageLibResponse
func (client *Client) AddImageLib(request *AddImageLibRequest) (_result *AddImageLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddImageLibResponse{}
	_body, _err := client.AddImageLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量添加图片
//
// @param request - AddImages2LibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddImages2LibResponse
func (client *Client) AddImages2LibWithOptions(request *AddImages2LibRequest, runtime *util.RuntimeOptions) (_result *AddImages2LibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImgUrl)) {
		body["ImgUrl"] = request.ImgUrl
	}

	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		body["LibId"] = request.LibId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddImages2Lib"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddImages2LibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量添加图片
//
// @param request - AddImages2LibRequest
//
// @return AddImages2LibResponse
func (client *Client) AddImages2Lib(request *AddImages2LibRequest) (_result *AddImages2LibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddImages2LibResponse{}
	_body, _err := client.AddImages2LibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建关键词库
//
// @param request - AddKeywordLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddKeywordLibResponse
func (client *Client) AddKeywordLibWithOptions(request *AddKeywordLibRequest, runtime *util.RuntimeOptions) (_result *AddKeywordLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		body["Keywords"] = request.Keywords
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordsObject)) {
		body["KeywordsObject"] = request.KeywordsObject
	}

	if !tea.BoolValue(util.IsUnset(request.LibName)) {
		body["LibName"] = request.LibName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddKeywordLib"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddKeywordLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建关键词库
//
// @param request - AddKeywordLibRequest
//
// @return AddKeywordLibResponse
func (client *Client) AddKeywordLib(request *AddKeywordLibRequest) (_result *AddKeywordLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddKeywordLibResponse{}
	_body, _err := client.AddKeywordLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加关键词
//
// @param request - AddKeywordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddKeywordsResponse
func (client *Client) AddKeywordsWithOptions(request *AddKeywordsRequest, runtime *util.RuntimeOptions) (_result *AddKeywordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		body["Keywords"] = request.Keywords
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordsObject)) {
		body["KeywordsObject"] = request.KeywordsObject
	}

	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		body["LibId"] = request.LibId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddKeywords"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddKeywordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加关键词
//
// @param request - AddKeywordsRequest
//
// @return AddKeywordsResponse
func (client *Client) AddKeywords(request *AddKeywordsRequest) (_result *AddKeywordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddKeywordsResponse{}
	_body, _err := client.AddKeywordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加关键词
//
// @param request - AddKeywordsToLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddKeywordsToLibResponse
func (client *Client) AddKeywordsToLibWithOptions(request *AddKeywordsToLibRequest, runtime *util.RuntimeOptions) (_result *AddKeywordsToLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		body["Keywords"] = request.Keywords
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordsObject)) {
		body["KeywordsObject"] = request.KeywordsObject
	}

	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		body["LibId"] = request.LibId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddKeywordsToLib"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddKeywordsToLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加关键词
//
// @param request - AddKeywordsToLibRequest
//
// @return AddKeywordsToLibResponse
func (client *Client) AddKeywordsToLib(request *AddKeywordsToLibRequest) (_result *AddKeywordsToLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddKeywordsToLibResponse{}
	_body, _err := client.AddKeywordsToLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消oss扫描任务
//
// @param request - CancelStockOssCheckTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelStockOssCheckTaskResponse
func (client *Client) CancelStockOssCheckTaskWithOptions(request *CancelStockOssCheckTaskRequest, runtime *util.RuntimeOptions) (_result *CancelStockOssCheckTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelStockOssCheckTask"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelStockOssCheckTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消oss扫描任务
//
// @param request - CancelStockOssCheckTaskRequest
//
// @return CancelStockOssCheckTaskResponse
func (client *Client) CancelStockOssCheckTask(request *CancelStockOssCheckTaskRequest) (_result *CancelStockOssCheckTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelStockOssCheckTaskResponse{}
	_body, _err := client.CancelStockOssCheckTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 复制服务
//
// @param request - CopyServiceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyServiceConfigResponse
func (client *Client) CopyServiceConfigWithOptions(request *CopyServiceConfigRequest, runtime *util.RuntimeOptions) (_result *CopyServiceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceDesc)) {
		body["ServiceDesc"] = request.ServiceDesc
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["ServiceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CopyServiceConfig"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyServiceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 复制服务
//
// @param request - CopyServiceConfigRequest
//
// @return CopyServiceConfigResponse
func (client *Client) CopyServiceConfig(request *CopyServiceConfigRequest) (_result *CopyServiceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CopyServiceConfigResponse{}
	_body, _err := client.CopyServiceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建oss扫描任务
//
// @param request - CreatStockOssCheckTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatStockOssCheckTaskResponse
func (client *Client) CreatStockOssCheckTaskWithOptions(request *CreatStockOssCheckTaskRequest, runtime *util.RuntimeOptions) (_result *CreatStockOssCheckTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Buckets)) {
		query["Buckets"] = request.Buckets
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackId)) {
		query["CallbackId"] = request.CallbackId
	}

	if !tea.BoolValue(util.IsUnset(request.DistinctHistoryTasks)) {
		query["DistinctHistoryTasks"] = request.DistinctHistoryTasks
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteDate)) {
		query["ExecuteDate"] = request.ExecuteDate
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteTime)) {
		query["ExecuteTime"] = request.ExecuteTime
	}

	if !tea.BoolValue(util.IsUnset(request.Freeze)) {
		query["Freeze"] = request.Freeze
	}

	if !tea.BoolValue(util.IsUnset(request.FreezeHighRisk1)) {
		query["FreezeHighRisk1"] = request.FreezeHighRisk1
	}

	if !tea.BoolValue(util.IsUnset(request.FreezeHighRisk2)) {
		query["FreezeHighRisk2"] = request.FreezeHighRisk2
	}

	if !tea.BoolValue(util.IsUnset(request.FreezeMediumRisk1)) {
		query["FreezeMediumRisk1"] = request.FreezeMediumRisk1
	}

	if !tea.BoolValue(util.IsUnset(request.FreezeMediumRisk2)) {
		query["FreezeMediumRisk2"] = request.FreezeMediumRisk2
	}

	if !tea.BoolValue(util.IsUnset(request.FreezeType)) {
		query["FreezeType"] = request.FreezeType
	}

	if !tea.BoolValue(util.IsUnset(request.IsInc)) {
		query["IsInc"] = request.IsInc
	}

	if !tea.BoolValue(util.IsUnset(request.MediaType)) {
		query["MediaType"] = request.MediaType
	}

	if !tea.BoolValue(util.IsUnset(request.PrefixFilterType)) {
		query["PrefixFilterType"] = request.PrefixFilterType
	}

	if !tea.BoolValue(util.IsUnset(request.PrefixFilters)) {
		query["PrefixFilters"] = request.PrefixFilters
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ScanLimit)) {
		query["ScanLimit"] = request.ScanLimit
	}

	if !tea.BoolValue(util.IsUnset(request.ScanNoFileType)) {
		query["ScanNoFileType"] = request.ScanNoFileType
	}

	if !tea.BoolValue(util.IsUnset(request.ScanResourceType)) {
		query["ScanResourceType"] = request.ScanResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ScanService)) {
		query["ScanService"] = request.ScanService
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskCycle)) {
		query["TaskCycle"] = request.TaskCycle
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatStockOssCheckTask"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatStockOssCheckTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建oss扫描任务
//
// @param request - CreatStockOssCheckTaskRequest
//
// @return CreatStockOssCheckTaskResponse
func (client *Client) CreatStockOssCheckTask(request *CreatStockOssCheckTaskRequest) (_result *CreatStockOssCheckTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatStockOssCheckTaskResponse{}
	_body, _err := client.CreatStockOssCheckTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建代答库
//
// @param request - CreateAnswerLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAnswerLibResponse
func (client *Client) CreateAnswerLibWithOptions(request *CreateAnswerLibRequest, runtime *util.RuntimeOptions) (_result *CreateAnswerLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibName)) {
		body["LibName"] = request.LibName
	}

	if !tea.BoolValue(util.IsUnset(request.SampleBucket)) {
		body["SampleBucket"] = request.SampleBucket
	}

	if !tea.BoolValue(util.IsUnset(request.SampleObject)) {
		body["SampleObject"] = request.SampleObject
	}

	if !tea.BoolValue(util.IsUnset(request.Samples)) {
		body["Samples"] = request.Samples
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAnswerLib"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAnswerLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建代答库
//
// @param request - CreateAnswerLibRequest
//
// @return CreateAnswerLibResponse
func (client *Client) CreateAnswerLib(request *CreateAnswerLibRequest) (_result *CreateAnswerLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAnswerLibResponse{}
	_body, _err := client.CreateAnswerLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建消息通知
//
// @param request - CreateCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCallbackResponse
func (client *Client) CreateCallbackWithOptions(request *CreateCallbackRequest, runtime *util.RuntimeOptions) (_result *CreateCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CryptType)) {
		body["CryptType"] = request.CryptType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCallback"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建消息通知
//
// @param request - CreateCallbackRequest
//
// @return CreateCallbackResponse
func (client *Client) CreateCallback(request *CreateCallbackRequest) (_result *CreateCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCallbackResponse{}
	_body, _err := client.CreateCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建oss扫描任务前检查
//
// @param request - CreatePreCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePreCheckResponse
func (client *Client) CreatePreCheckWithOptions(request *CreatePreCheckRequest, runtime *util.RuntimeOptions) (_result *CreatePreCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Buckets)) {
		body["Buckets"] = request.Buckets
	}

	if !tea.BoolValue(util.IsUnset(request.DistinctHistoryTasks)) {
		body["DistinctHistoryTasks"] = request.DistinctHistoryTasks
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsInc)) {
		body["IsInc"] = request.IsInc
	}

	if !tea.BoolValue(util.IsUnset(request.MediaType)) {
		body["MediaType"] = request.MediaType
	}

	if !tea.BoolValue(util.IsUnset(request.PrefixFilterType)) {
		body["PrefixFilterType"] = request.PrefixFilterType
	}

	if !tea.BoolValue(util.IsUnset(request.PrefixFilters)) {
		body["PrefixFilters"] = request.PrefixFilters
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.ScanLimit)) {
		body["ScanLimit"] = request.ScanLimit
	}

	if !tea.BoolValue(util.IsUnset(request.ScanNoFileType)) {
		body["ScanNoFileType"] = request.ScanNoFileType
	}

	if !tea.BoolValue(util.IsUnset(request.ScanService)) {
		body["ScanService"] = request.ScanService
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePreCheck"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePreCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建oss扫描任务前检查
//
// @param request - CreatePreCheckRequest
//
// @return CreatePreCheckResponse
func (client *Client) CreatePreCheck(request *CreatePreCheckRequest) (_result *CreatePreCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePreCheckResponse{}
	_body, _err := client.CreatePreCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除代答库
//
// @param request - DeleteAnswerLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAnswerLibResponse
func (client *Client) DeleteAnswerLibWithOptions(request *DeleteAnswerLibRequest, runtime *util.RuntimeOptions) (_result *DeleteAnswerLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		query["LibId"] = request.LibId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAnswerLib"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAnswerLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除代答库
//
// @param request - DeleteAnswerLibRequest
//
// @return DeleteAnswerLibResponse
func (client *Client) DeleteAnswerLib(request *DeleteAnswerLibRequest) (_result *DeleteAnswerLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAnswerLibResponse{}
	_body, _err := client.DeleteAnswerLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除代答答案
//
// @param request - DeleteAnswerSampleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAnswerSampleResponse
func (client *Client) DeleteAnswerSampleWithOptions(request *DeleteAnswerSampleRequest, runtime *util.RuntimeOptions) (_result *DeleteAnswerSampleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		body["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		body["LibId"] = request.LibId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAnswerSample"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAnswerSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除代答答案
//
// @param request - DeleteAnswerSampleRequest
//
// @return DeleteAnswerSampleResponse
func (client *Client) DeleteAnswerSample(request *DeleteAnswerSampleRequest) (_result *DeleteAnswerSampleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAnswerSampleResponse{}
	_body, _err := client.DeleteAnswerSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除消息通知
//
// @param request - DeleteCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCallbackResponse
func (client *Client) DeleteCallbackWithOptions(request *DeleteCallbackRequest, runtime *util.RuntimeOptions) (_result *DeleteCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCallback"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除消息通知
//
// @param request - DeleteCallbackRequest
//
// @return DeleteCallbackResponse
func (client *Client) DeleteCallback(request *DeleteCallbackRequest) (_result *DeleteCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCallbackResponse{}
	_body, _err := client.DeleteCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete feature configuration
//
// @param request - DeleteFeatureConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFeatureConfigResponse
func (client *Client) DeleteFeatureConfigWithOptions(request *DeleteFeatureConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteFeatureConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Field)) {
		body["Field"] = request.Field
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFeatureConfig"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFeatureConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete feature configuration
//
// @param request - DeleteFeatureConfigRequest
//
// @return DeleteFeatureConfigResponse
func (client *Client) DeleteFeatureConfig(request *DeleteFeatureConfigRequest) (_result *DeleteFeatureConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFeatureConfigResponse{}
	_body, _err := client.DeleteFeatureConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除
//
// @param request - DeleteImagesFromLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteImagesFromLibResponse
func (client *Client) DeleteImagesFromLibWithOptions(request *DeleteImagesFromLibRequest, runtime *util.RuntimeOptions) (_result *DeleteImagesFromLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageIds)) {
		body["ImageIds"] = request.ImageIds
	}

	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		body["LibId"] = request.LibId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteImagesFromLib"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteImagesFromLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除
//
// @param request - DeleteImagesFromLibRequest
//
// @return DeleteImagesFromLibResponse
func (client *Client) DeleteImagesFromLib(request *DeleteImagesFromLibRequest) (_result *DeleteImagesFromLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImagesFromLibResponse{}
	_body, _err := client.DeleteImagesFromLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除关键词
//
// @param request - DeleteKeywordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKeywordResponse
func (client *Client) DeleteKeywordWithOptions(request *DeleteKeywordRequest, runtime *util.RuntimeOptions) (_result *DeleteKeywordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeywordIdList)) {
		body["KeywordIdList"] = request.KeywordIdList
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordIds)) {
		body["KeywordIds"] = request.KeywordIds
	}

	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		body["LibId"] = request.LibId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteKeyword"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteKeywordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除关键词
//
// @param request - DeleteKeywordRequest
//
// @return DeleteKeywordResponse
func (client *Client) DeleteKeyword(request *DeleteKeywordRequest) (_result *DeleteKeywordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteKeywordResponse{}
	_body, _err := client.DeleteKeywordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除关键词库
//
// @param request - DeleteKeywordLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKeywordLibResponse
func (client *Client) DeleteKeywordLibWithOptions(request *DeleteKeywordLibRequest, runtime *util.RuntimeOptions) (_result *DeleteKeywordLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		body["LibId"] = request.LibId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteKeywordLib"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteKeywordLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除关键词库
//
// @param request - DeleteKeywordLibRequest
//
// @return DeleteKeywordLibResponse
func (client *Client) DeleteKeywordLib(request *DeleteKeywordLibRequest) (_result *DeleteKeywordLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteKeywordLibResponse{}
	_body, _err := client.DeleteKeywordLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出代答答案
//
// @param request - ExportAnswerSampleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportAnswerSampleResponse
func (client *Client) ExportAnswerSampleWithOptions(request *ExportAnswerSampleRequest, runtime *util.RuntimeOptions) (_result *ExportAnswerSampleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		body["LibId"] = request.LibId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportAnswerSample"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportAnswerSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出代答答案
//
// @param request - ExportAnswerSampleRequest
//
// @return ExportAnswerSampleResponse
func (client *Client) ExportAnswerSample(request *ExportAnswerSampleRequest) (_result *ExportAnswerSampleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportAnswerSampleResponse{}
	_body, _err := client.ExportAnswerSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出调用量
//
// @param request - ExportCipStatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportCipStatsResponse
func (client *Client) ExportCipStatsWithOptions(request *ExportCipStatsRequest, runtime *util.RuntimeOptions) (_result *ExportCipStatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ByMonth)) {
		body["ByMonth"] = request.ByMonth
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ExportType)) {
		body["ExportType"] = request.ExportType
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		body["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SubUid)) {
		body["SubUid"] = request.SubUid
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportCipStats"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportCipStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出调用量
//
// @param request - ExportCipStatsRequest
//
// @return ExportCipStatsResponse
func (client *Client) ExportCipStats(request *ExportCipStatsRequest) (_result *ExportCipStatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportCipStatsResponse{}
	_body, _err := client.ExportCipStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出关键词
//
// @param request - ExportKeywordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportKeywordResponse
func (client *Client) ExportKeywordWithOptions(request *ExportKeywordRequest, runtime *util.RuntimeOptions) (_result *ExportKeywordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		body["LibId"] = request.LibId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportKeyword"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportKeywordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出关键词
//
// @param request - ExportKeywordRequest
//
// @return ExportKeywordResponse
func (client *Client) ExportKeyword(request *ExportKeywordRequest) (_result *ExportKeywordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportKeywordResponse{}
	_body, _err := client.ExportKeywordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// oss用量统计导出
//
// @param request - ExportOssCheckStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportOssCheckStatResponse
func (client *Client) ExportOssCheckStatWithOptions(request *ExportOssCheckStatRequest, runtime *util.RuntimeOptions) (_result *ExportOssCheckStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ByMonth)) {
		body["ByMonth"] = request.ByMonth
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ParentTaskId)) {
		body["ParentTaskId"] = request.ParentTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportOssCheckStat"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportOssCheckStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// oss用量统计导出
//
// @param request - ExportOssCheckStatRequest
//
// @return ExportOssCheckStatResponse
func (client *Client) ExportOssCheckStat(request *ExportOssCheckStatRequest) (_result *ExportOssCheckStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportOssCheckStatResponse{}
	_body, _err := client.ExportOssCheckStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出oss扫描结果
//
// @param tmpReq - ExportResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportResultResponse
func (client *Client) ExportResultWithOptions(tmpReq *ExportResultRequest, runtime *util.RuntimeOptions) (_result *ExportResultResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExportResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Sort)) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, tea.String("Sort"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.SortShrink)) {
		body["Sort"] = request.SortShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportResult"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出oss扫描结果
//
// @param request - ExportResultRequest
//
// @return ExportResultResponse
func (client *Client) ExportResult(request *ExportResultRequest) (_result *ExportResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportResultResponse{}
	_body, _err := client.ExportResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出调用结果，excel文件
//
// @param tmpReq - ExportScanResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportScanResultResponse
func (client *Client) ExportScanResultWithOptions(tmpReq *ExportScanResultRequest, runtime *util.RuntimeOptions) (_result *ExportScanResultResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExportScanResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Query)) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, tea.String("Query"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Sort)) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, tea.String("Sort"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryShrink)) {
		body["Query"] = request.QueryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SortShrink)) {
		body["Sort"] = request.SortShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportScanResult"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportScanResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出调用结果，excel文件
//
// @param request - ExportScanResultRequest
//
// @return ExportScanResultResponse
func (client *Client) ExportScanResult(request *ExportScanResultRequest) (_result *ExportScanResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportScanResultResponse{}
	_body, _err := client.ExportScanResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出调用结果，excel文件
//
// @param tmpReq - ExportTextScanResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportTextScanResultResponse
func (client *Client) ExportTextScanResultWithOptions(tmpReq *ExportTextScanResultRequest, runtime *util.RuntimeOptions) (_result *ExportTextScanResultResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExportTextScanResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Query)) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, tea.String("Query"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.QueryShrink)) {
		body["Query"] = request.QueryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportTextScanResult"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportTextScanResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出调用结果，excel文件
//
// @param request - ExportTextScanResultRequest
//
// @return ExportTextScanResultResponse
func (client *Client) ExportTextScanResult(request *ExportTextScanResultRequest) (_result *ExportTextScanResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportTextScanResultResponse{}
	_body, _err := client.ExportTextScanResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取代答样本导入进度
//
// @param request - GetAnswerImportProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAnswerImportProgressResponse
func (client *Client) GetAnswerImportProgressWithOptions(request *GetAnswerImportProgressRequest, runtime *util.RuntimeOptions) (_result *GetAnswerImportProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAnswerImportProgress"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAnswerImportProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取代答样本导入进度
//
// @param request - GetAnswerImportProgressRequest
//
// @return GetAnswerImportProgressResponse
func (client *Client) GetAnswerImportProgress(request *GetAnswerImportProgressRequest) (_result *GetAnswerImportProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAnswerImportProgressResponse{}
	_body, _err := client.GetAnswerImportProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 证据转存获取用户bucket列表
//
// @param request - GetBackupBucketsListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBackupBucketsListResponse
func (client *Client) GetBackupBucketsListWithOptions(request *GetBackupBucketsListRequest, runtime *util.RuntimeOptions) (_result *GetBackupBucketsListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBackupBucketsList"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBackupBucketsListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 证据转存获取用户bucket列表
//
// @param request - GetBackupBucketsListRequest
//
// @return GetBackupBucketsListResponse
func (client *Client) GetBackupBucketsList(request *GetBackupBucketsListRequest) (_result *GetBackupBucketsListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBackupBucketsListResponse{}
	_body, _err := client.GetBackupBucketsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取证据转存配置
//
// @param request - GetBackupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBackupConfigResponse
func (client *Client) GetBackupConfigWithOptions(request *GetBackupConfigRequest, runtime *util.RuntimeOptions) (_result *GetBackupConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		query["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBackupConfig"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBackupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取证据转存配置
//
// @param request - GetBackupConfigRequest
//
// @return GetBackupConfigResponse
func (client *Client) GetBackupConfig(request *GetBackupConfigRequest) (_result *GetBackupConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBackupConfigResponse{}
	_body, _err := client.GetBackupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户授权校验
//
// @param request - GetBackupStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBackupStatusResponse
func (client *Client) GetBackupStatusWithOptions(request *GetBackupStatusRequest, runtime *util.RuntimeOptions) (_result *GetBackupStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBackupStatus"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBackupStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户授权校验
//
// @param request - GetBackupStatusRequest
//
// @return GetBackupStatusResponse
func (client *Client) GetBackupStatus(request *GetBackupStatusRequest) (_result *GetBackupStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBackupStatusResponse{}
	_body, _err := client.GetBackupStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// bucket列表
//
// @param request - GetBucketsListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBucketsListResponse
func (client *Client) GetBucketsListWithOptions(request *GetBucketsListRequest, runtime *util.RuntimeOptions) (_result *GetBucketsListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketsList"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBucketsListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// bucket列表
//
// @param request - GetBucketsListRequest
//
// @return GetBucketsListResponse
func (client *Client) GetBucketsList(request *GetBucketsListRequest) (_result *GetBucketsListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBucketsListResponse{}
	_body, _err := client.GetBucketsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询调用量
//
// @param request - GetCipStatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCipStatsResponse
func (client *Client) GetCipStatsWithOptions(request *GetCipStatsRequest, runtime *util.RuntimeOptions) (_result *GetCipStatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ByMonth)) {
		body["ByMonth"] = request.ByMonth
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		body["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SubUid)) {
		body["SubUid"] = request.SubUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCipStats"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCipStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询调用量
//
// @param request - GetCipStatsRequest
//
// @return GetCipStatsResponse
func (client *Client) GetCipStats(request *GetCipStatsRequest) (_result *GetCipStatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCipStatsResponse{}
	_body, _err := client.GetCipStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取定时任务预计执行时间
//
// @param request - GetExecuteTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExecuteTimeResponse
func (client *Client) GetExecuteTimeWithOptions(request *GetExecuteTimeRequest, runtime *util.RuntimeOptions) (_result *GetExecuteTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExecuteTime"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExecuteTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取定时任务预计执行时间
//
// @param request - GetExecuteTimeRequest
//
// @return GetExecuteTimeResponse
func (client *Client) GetExecuteTime(request *GetExecuteTimeRequest) (_result *GetExecuteTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExecuteTimeResponse{}
	_body, _err := client.GetExecuteTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Feature Configuration
//
// @param request - GetFeatureConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFeatureConfigResponse
func (client *Client) GetFeatureConfigWithOptions(request *GetFeatureConfigRequest, runtime *util.RuntimeOptions) (_result *GetFeatureConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFeatureConfig"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFeatureConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Feature Configuration
//
// @param request - GetFeatureConfigRequest
//
// @return GetFeatureConfigResponse
func (client *Client) GetFeatureConfig(request *GetFeatureConfigRequest) (_result *GetFeatureConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFeatureConfigResponse{}
	_body, _err := client.GetFeatureConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取图片规则标签信息
//
// @param request - GetImageSceneLabelConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageSceneLabelConfResponse
func (client *Client) GetImageSceneLabelConfWithOptions(request *GetImageSceneLabelConfRequest, runtime *util.RuntimeOptions) (_result *GetImageSceneLabelConfResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImageSceneLabelConf"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImageSceneLabelConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取图片规则标签信息
//
// @param request - GetImageSceneLabelConfRequest
//
// @return GetImageSceneLabelConfResponse
func (client *Client) GetImageSceneLabelConf(request *GetImageSceneLabelConfRequest) (_result *GetImageSceneLabelConfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImageSceneLabelConfResponse{}
	_body, _err := client.GetImageSceneLabelConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取图片规则标签信息
//
// @param request - GetImageSceneLabelListConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageSceneLabelListConfResponse
func (client *Client) GetImageSceneLabelListConfWithOptions(request *GetImageSceneLabelListConfRequest, runtime *util.RuntimeOptions) (_result *GetImageSceneLabelListConfResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageServiceCode)) {
		query["ImageServiceCode"] = request.ImageServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImageSceneLabelListConf"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImageSceneLabelListConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取图片规则标签信息
//
// @param request - GetImageSceneLabelListConfRequest
//
// @return GetImageSceneLabelListConfResponse
func (client *Client) GetImageSceneLabelListConf(request *GetImageSceneLabelListConfRequest) (_result *GetImageSceneLabelListConfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImageSceneLabelListConfResponse{}
	_body, _err := client.GetImageSceneLabelListConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// oss定时扫描检测周期查询
//
// @param tmpReq - GetJobNameListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobNameListResponse
func (client *Client) GetJobNameListWithOptions(tmpReq *GetJobNameListRequest, runtime *util.RuntimeOptions) (_result *GetJobNameListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetJobNameListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Sort)) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, tea.String("Sort"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SortShrink)) {
		query["Sort"] = request.SortShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobNameList"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobNameListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// oss定时扫描检测周期查询
//
// @param request - GetJobNameListRequest
//
// @return GetJobNameListResponse
func (client *Client) GetJobNameList(request *GetJobNameListRequest) (_result *GetJobNameListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobNameListResponse{}
	_body, _err := client.GetJobNameListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询导入关键词结果
//
// @param request - GetKeywordImportResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKeywordImportResultResponse
func (client *Client) GetKeywordImportResultWithOptions(request *GetKeywordImportResultRequest, runtime *util.RuntimeOptions) (_result *GetKeywordImportResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetKeywordImportResult"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetKeywordImportResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询导入关键词结果
//
// @param request - GetKeywordImportResultRequest
//
// @return GetKeywordImportResultResponse
func (client *Client) GetKeywordImportResult(request *GetKeywordImportResultRequest) (_result *GetKeywordImportResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetKeywordImportResultResponse{}
	_body, _err := client.GetKeywordImportResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关键词库信息
//
// @param request - GetKeywordLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKeywordLibResponse
func (client *Client) GetKeywordLibWithOptions(request *GetKeywordLibRequest, runtime *util.RuntimeOptions) (_result *GetKeywordLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		body["LibId"] = request.LibId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetKeywordLib"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetKeywordLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关键词库信息
//
// @param request - GetKeywordLibRequest
//
// @return GetKeywordLibResponse
func (client *Client) GetKeywordLib(request *GetKeywordLibRequest) (_result *GetKeywordLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetKeywordLibResponse{}
	_body, _err := client.GetKeywordLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// oss用量统计
//
// @param request - GetOssCheckStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssCheckStatResponse
func (client *Client) GetOssCheckStatWithOptions(request *GetOssCheckStatRequest, runtime *util.RuntimeOptions) (_result *GetOssCheckStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ByMonth)) {
		body["ByMonth"] = request.ByMonth
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ParentTaskId)) {
		body["ParentTaskId"] = request.ParentTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOssCheckStat"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOssCheckStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// oss用量统计
//
// @param request - GetOssCheckStatRequest
//
// @return GetOssCheckStatResponse
func (client *Client) GetOssCheckStat(request *GetOssCheckStatRequest) (_result *GetOssCheckStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOssCheckStatResponse{}
	_body, _err := client.GetOssCheckStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取OSS检测用户状态
//
// @param request - GetOssCheckStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssCheckStatusResponse
func (client *Client) GetOssCheckStatusWithOptions(request *GetOssCheckStatusRequest, runtime *util.RuntimeOptions) (_result *GetOssCheckStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOssCheckStatus"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOssCheckStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取OSS检测用户状态
//
// @param request - GetOssCheckStatusRequest
//
// @return GetOssCheckStatusResponse
func (client *Client) GetOssCheckStatus(request *GetOssCheckStatusRequest) (_result *GetOssCheckStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOssCheckStatusResponse{}
	_body, _err := client.GetOssCheckStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户待检测信息
//
// @param request - GetScanNumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScanNumResponse
func (client *Client) GetScanNumWithOptions(request *GetScanNumRequest, runtime *util.RuntimeOptions) (_result *GetScanNumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Buckets)) {
		query["Buckets"] = request.Buckets
	}

	if !tea.BoolValue(util.IsUnset(request.MediaType)) {
		query["MediaType"] = request.MediaType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetScanNum"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetScanNumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户待检测信息
//
// @param request - GetScanNumRequest
//
// @return GetScanNumResponse
func (client *Client) GetScanNum(request *GetScanNumRequest) (_result *GetScanNumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetScanNumResponse{}
	_body, _err := client.GetScanNumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询调用结果
//
// @param tmpReq - GetScanResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScanResultResponse
func (client *Client) GetScanResultWithOptions(tmpReq *GetScanResultRequest, runtime *util.RuntimeOptions) (_result *GetScanResultResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetScanResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Query)) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, tea.String("Query"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Sort)) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, tea.String("Sort"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryShrink)) {
		body["Query"] = request.QueryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SortShrink)) {
		body["Sort"] = request.SortShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetScanResult"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetScanResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询调用结果
//
// @param request - GetScanResultRequest
//
// @return GetScanResultResponse
func (client *Client) GetScanResult(request *GetScanResultRequest) (_result *GetScanResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetScanResultResponse{}
	_body, _err := client.GetScanResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取单个服务
//
// @param request - GetServiceConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceConfResponse
func (client *Client) GetServiceConfWithOptions(request *GetServiceConfRequest, runtime *util.RuntimeOptions) (_result *GetServiceConfResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ByDefault)) {
		body["ByDefault"] = request.ByDefault
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceConf"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单个服务
//
// @param request - GetServiceConfRequest
//
// @return GetServiceConfResponse
func (client *Client) GetServiceConf(request *GetServiceConfRequest) (_result *GetServiceConfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceConfResponse{}
	_body, _err := client.GetServiceConfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取单个服务
//
// @param request - GetServiceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceConfigResponse
func (client *Client) GetServiceConfigWithOptions(request *GetServiceConfigRequest, runtime *util.RuntimeOptions) (_result *GetServiceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceConfig"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单个服务
//
// @param request - GetServiceConfigRequest
//
// @return GetServiceConfigResponse
func (client *Client) GetServiceConfig(request *GetServiceConfigRequest) (_result *GetServiceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceConfigResponse{}
	_body, _err := client.GetServiceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取单个服务的标签配置
//
// @param request - GetServiceLabelConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceLabelConfigResponse
func (client *Client) GetServiceLabelConfigWithOptions(request *GetServiceLabelConfigRequest, runtime *util.RuntimeOptions) (_result *GetServiceLabelConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceLabelConfig"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceLabelConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单个服务的标签配置
//
// @param request - GetServiceLabelConfigRequest
//
// @return GetServiceLabelConfigResponse
func (client *Client) GetServiceLabelConfig(request *GetServiceLabelConfigRequest) (_result *GetServiceLabelConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceLabelConfigResponse{}
	_body, _err := client.GetServiceLabelConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询oss扫描任务列表
//
// @param tmpReq - GetStockOssCheckTasksListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStockOssCheckTasksListResponse
func (client *Client) GetStockOssCheckTasksListWithOptions(tmpReq *GetStockOssCheckTasksListRequest, runtime *util.RuntimeOptions) (_result *GetStockOssCheckTasksListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetStockOssCheckTasksListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Sort)) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, tea.String("Sort"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsInc)) {
		query["IsInc"] = request.IsInc
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MediaType)) {
		body["MediaType"] = request.MediaType
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortShrink)) {
		body["Sort"] = request.SortShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStockOssCheckTasksList"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStockOssCheckTasksListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询oss扫描任务列表
//
// @param request - GetStockOssCheckTasksListRequest
//
// @return GetStockOssCheckTasksListResponse
func (client *Client) GetStockOssCheckTasksList(request *GetStockOssCheckTasksListRequest) (_result *GetStockOssCheckTasksListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStockOssCheckTasksListResponse{}
	_body, _err := client.GetStockOssCheckTasksListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询调用结果
//
// @param tmpReq - GetTextScanResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextScanResultResponse
func (client *Client) GetTextScanResultWithOptions(tmpReq *GetTextScanResultRequest, runtime *util.RuntimeOptions) (_result *GetTextScanResultResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetTextScanResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Query)) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, tea.String("Query"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Sort)) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, tea.String("Sort"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryShrink)) {
		body["Query"] = request.QueryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SortShrink)) {
		body["Sort"] = request.SortShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTextScanResult"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTextScanResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询调用结果
//
// @param request - GetTextScanResultRequest
//
// @return GetTextScanResultResponse
func (client *Client) GetTextScanResult(request *GetTextScanResultRequest) (_result *GetTextScanResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTextScanResultResponse{}
	_body, _err := client.GetTextScanResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文件上传获取相应信息
//
// @param request - GetUploadInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadInfoResponse
func (client *Client) GetUploadInfoWithOptions(request *GetUploadInfoRequest, runtime *util.RuntimeOptions) (_result *GetUploadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUploadInfo"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUploadInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文件上传获取相应信息
//
// @param request - GetUploadInfoRequest
//
// @return GetUploadInfoResponse
func (client *Client) GetUploadInfo(request *GetUploadInfoRequest) (_result *GetUploadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUploadInfoResponse{}
	_body, _err := client.GetUploadInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户购买状态
//
// @param request - GetUserBuyStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserBuyStatusResponse
func (client *Client) GetUserBuyStatusWithOptions(request *GetUserBuyStatusRequest, runtime *util.RuntimeOptions) (_result *GetUserBuyStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		body["CommodityCode"] = request.CommodityCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserBuyStatus"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserBuyStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户购买状态
//
// @param request - GetUserBuyStatusRequest
//
// @return GetUserBuyStatusResponse
func (client *Client) GetUserBuyStatus(request *GetUserBuyStatusRequest) (_result *GetUserBuyStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserBuyStatusResponse{}
	_body, _err := client.GetUserBuyStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 代答库列表
//
// @param request - ListAnswerLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAnswerLibResponse
func (client *Client) ListAnswerLibWithOptions(request *ListAnswerLibRequest, runtime *util.RuntimeOptions) (_result *ListAnswerLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAnswerLib"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAnswerLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 代答库列表
//
// @param request - ListAnswerLibRequest
//
// @return ListAnswerLibResponse
func (client *Client) ListAnswerLib(request *ListAnswerLibRequest) (_result *ListAnswerLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAnswerLibResponse{}
	_body, _err := client.ListAnswerLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 消息通知列表
//
// @param request - ListCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCallbackResponse
func (client *Client) ListCallbackWithOptions(request *ListCallbackRequest, runtime *util.RuntimeOptions) (_result *ListCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCallback"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 消息通知列表
//
// @param request - ListCallbackRequest
//
// @return ListCallbackResponse
func (client *Client) ListCallback(request *ListCallbackRequest) (_result *ListCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCallbackResponse{}
	_body, _err := client.ListCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 图库列表
//
// @param request - ListImageLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImageLibResponse
func (client *Client) ListImageLibWithOptions(request *ListImageLibRequest, runtime *util.RuntimeOptions) (_result *ListImageLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImageLib"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImageLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图库列表
//
// @param request - ListImageLibRequest
//
// @return ListImageLibResponse
func (client *Client) ListImageLib(request *ListImageLibRequest) (_result *ListImageLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImageLibResponse{}
	_body, _err := client.ListImageLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 图片列表 分页
//
// @param tmpReq - ListImagesFromLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImagesFromLibResponse
func (client *Client) ListImagesFromLibWithOptions(tmpReq *ListImagesFromLibRequest, runtime *util.RuntimeOptions) (_result *ListImagesFromLibResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListImagesFromLibShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Sort)) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, tea.String("Sort"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ImgId)) {
		body["ImgId"] = request.ImgId
	}

	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		body["LibId"] = request.LibId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortShrink)) {
		body["Sort"] = request.SortShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImagesFromLib"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImagesFromLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片列表 分页
//
// @param request - ListImagesFromLibRequest
//
// @return ListImagesFromLibResponse
func (client *Client) ListImagesFromLib(request *ListImagesFromLibRequest) (_result *ListImagesFromLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImagesFromLibResponse{}
	_body, _err := client.ListImagesFromLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关键词库列表
//
// @param request - ListKeywordLibsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKeywordLibsResponse
func (client *Client) ListKeywordLibsWithOptions(request *ListKeywordLibsRequest, runtime *util.RuntimeOptions) (_result *ListKeywordLibsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListKeywordLibs"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListKeywordLibsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关键词库列表
//
// @param request - ListKeywordLibsRequest
//
// @return ListKeywordLibsResponse
func (client *Client) ListKeywordLibs(request *ListKeywordLibsRequest) (_result *ListKeywordLibsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListKeywordLibsResponse{}
	_body, _err := client.ListKeywordLibsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询关键词列表
//
// @param tmpReq - ListKeywordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKeywordsResponse
func (client *Client) ListKeywordsWithOptions(tmpReq *ListKeywordsRequest, runtime *util.RuntimeOptions) (_result *ListKeywordsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListKeywordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Sort)) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, tea.String("Sort"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		body["LibId"] = request.LibId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortShrink)) {
		body["Sort"] = request.SortShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Word)) {
		body["Word"] = request.Word
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListKeywords"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListKeywordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询关键词列表
//
// @param request - ListKeywordsRequest
//
// @return ListKeywordsResponse
func (client *Client) ListKeywords(request *ListKeywordsRequest) (_result *ListKeywordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListKeywordsResponse{}
	_body, _err := client.ListKeywordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// oss扫描结果查询
//
// @param tmpReq - ListOssCheckResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOssCheckResultResponse
func (client *Client) ListOssCheckResultWithOptions(tmpReq *ListOssCheckResultRequest, runtime *util.RuntimeOptions) (_result *ListOssCheckResultResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListOssCheckResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Sort)) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, tea.String("Sort"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.FinishNum)) {
		query["FinishNum"] = request.FinishNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SortShrink)) {
		query["Sort"] = request.SortShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOssCheckResult"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOssCheckResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// oss扫描结果查询
//
// @param request - ListOssCheckResultRequest
//
// @return ListOssCheckResultResponse
func (client *Client) ListOssCheckResult(request *ListOssCheckResultRequest) (_result *ListOssCheckResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOssCheckResultResponse{}
	_body, _err := client.ListOssCheckResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取服务列表
//
// @param request - ListServiceConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceConfigsResponse
func (client *Client) ListServiceConfigsWithOptions(request *ListServiceConfigsRequest, runtime *util.RuntimeOptions) (_result *ListServiceConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Classify)) {
		query["Classify"] = request.Classify
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UseStatus)) {
		query["UseStatus"] = request.UseStatus
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceConfigs"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取服务列表
//
// @param request - ListServiceConfigsRequest
//
// @return ListServiceConfigsResponse
func (client *Client) ListServiceConfigs(request *ListServiceConfigsRequest) (_result *ListServiceConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceConfigsResponse{}
	_body, _err := client.ListServiceConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新代答库
//
// @param request - ModifyAnswerLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAnswerLibResponse
func (client *Client) ModifyAnswerLibWithOptions(request *ModifyAnswerLibRequest, runtime *util.RuntimeOptions) (_result *ModifyAnswerLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		query["LibId"] = request.LibId
	}

	if !tea.BoolValue(util.IsUnset(request.LibName)) {
		query["LibName"] = request.LibName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAnswerLib"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAnswerLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新代答库
//
// @param request - ModifyAnswerLibRequest
//
// @return ModifyAnswerLibResponse
func (client *Client) ModifyAnswerLib(request *ModifyAnswerLibRequest) (_result *ModifyAnswerLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAnswerLibResponse{}
	_body, _err := client.ModifyAnswerLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改消息通知
//
// @param request - ModifyCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCallbackResponse
func (client *Client) ModifyCallbackWithOptions(request *ModifyCallbackRequest, runtime *util.RuntimeOptions) (_result *ModifyCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CryptType)) {
		body["CryptType"] = request.CryptType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyCallback"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改消息通知
//
// @param request - ModifyCallbackRequest
//
// @return ModifyCallbackResponse
func (client *Client) ModifyCallback(request *ModifyCallbackRequest) (_result *ModifyCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCallbackResponse{}
	_body, _err := client.ModifyCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Save Feature Configuration
//
// @param request - ModifyFeatureConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFeatureConfigResponse
func (client *Client) ModifyFeatureConfigWithOptions(request *ModifyFeatureConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyFeatureConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Field)) {
		body["Field"] = request.Field
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFeatureConfig"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFeatureConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Save Feature Configuration
//
// @param request - ModifyFeatureConfigRequest
//
// @return ModifyFeatureConfigResponse
func (client *Client) ModifyFeatureConfig(request *ModifyFeatureConfigRequest) (_result *ModifyFeatureConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFeatureConfigResponse{}
	_body, _err := client.ModifyFeatureConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑服务
//
// @param request - ModifyServiceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyServiceInfoResponse
func (client *Client) ModifyServiceInfoWithOptions(request *ModifyServiceInfoRequest, runtime *util.RuntimeOptions) (_result *ModifyServiceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceDesc)) {
		body["ServiceDesc"] = request.ServiceDesc
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["ServiceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyServiceInfo"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyServiceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑服务
//
// @param request - ModifyServiceInfoRequest
//
// @return ModifyServiceInfoResponse
func (client *Client) ModifyServiceInfo(request *ModifyServiceInfoRequest) (_result *ModifyServiceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyServiceInfoResponse{}
	_body, _err := client.ModifyServiceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询代答样本
//
// @param tmpReq - QueryAnswerSampleByPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAnswerSampleByPageResponse
func (client *Client) QueryAnswerSampleByPageWithOptions(tmpReq *QueryAnswerSampleByPageRequest, runtime *util.RuntimeOptions) (_result *QueryAnswerSampleByPageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryAnswerSampleByPageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Sort)) {
		request.SortShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sort, tea.String("Sort"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Answer)) {
		query["Answer"] = request.Answer
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		query["LibId"] = request.LibId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SortShrink)) {
		query["Sort"] = request.SortShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAnswerSampleByPage"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAnswerSampleByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询代答样本
//
// @param request - QueryAnswerSampleByPageRequest
//
// @return QueryAnswerSampleByPageResponse
func (client *Client) QueryAnswerSampleByPage(request *QueryAnswerSampleByPageRequest) (_result *QueryAnswerSampleByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAnswerSampleByPageResponse{}
	_body, _err := client.QueryAnswerSampleByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个回调配置
//
// @param request - QueryCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCallbackResponse
func (client *Client) QueryCallbackWithOptions(request *QueryCallbackRequest, runtime *util.RuntimeOptions) (_result *QueryCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckForOss)) {
		body["CheckForOss"] = request.CheckForOss
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCallback"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个回调配置
//
// @param request - QueryCallbackRequest
//
// @return QueryCallbackResponse
func (client *Client) QueryCallback(request *QueryCallbackRequest) (_result *QueryCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCallbackResponse{}
	_body, _err := client.QueryCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 消息通知
//
// @param request - QueryCallbackByPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCallbackByPageResponse
func (client *Client) QueryCallbackByPageWithOptions(request *QueryCallbackByPageRequest, runtime *util.RuntimeOptions) (_result *QueryCallbackByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCallbackByPage"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCallbackByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 消息通知
//
// @param request - QueryCallbackByPageRequest
//
// @return QueryCallbackByPageResponse
func (client *Client) QueryCallbackByPage(request *QueryCallbackByPageRequest) (_result *QueryCallbackByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCallbackByPageResponse{}
	_body, _err := client.QueryCallbackByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新证据转存配置
//
// @param request - UpdateBackupConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBackupConfigResponse
func (client *Client) UpdateBackupConfigWithOptions(request *UpdateBackupConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateBackupConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupConfig)) {
		query["BackupConfig"] = request.BackupConfig
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		query["ServiceCode"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBackupConfig"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBackupConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新证据转存配置
//
// @param request - UpdateBackupConfigRequest
//
// @return UpdateBackupConfigResponse
func (client *Client) UpdateBackupConfig(request *UpdateBackupConfigRequest) (_result *UpdateBackupConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBackupConfigResponse{}
	_body, _err := client.UpdateBackupConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑图库
//
// @param request - UpdateImageLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateImageLibResponse
func (client *Client) UpdateImageLibWithOptions(request *UpdateImageLibRequest, runtime *util.RuntimeOptions) (_result *UpdateImageLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		body["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.FreeInspection)) {
		body["FreeInspection"] = request.FreeInspection
	}

	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		body["LibId"] = request.LibId
	}

	if !tea.BoolValue(util.IsUnset(request.LibName)) {
		body["LibName"] = request.LibName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateImageLib"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateImageLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑图库
//
// @param request - UpdateImageLibRequest
//
// @return UpdateImageLibResponse
func (client *Client) UpdateImageLib(request *UpdateImageLibRequest) (_result *UpdateImageLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateImageLibResponse{}
	_body, _err := client.UpdateImageLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑图库免检配置
//
// @param tmpReq - UpdateImageLibFreeInspectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateImageLibFreeInspectionResponse
func (client *Client) UpdateImageLibFreeInspectionWithOptions(tmpReq *UpdateImageLibFreeInspectionRequest, runtime *util.RuntimeOptions) (_result *UpdateImageLibFreeInspectionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateImageLibFreeInspectionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Config)) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, tea.String("Config"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigShrink)) {
		body["Config"] = request.ConfigShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateImageLibFreeInspection"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateImageLibFreeInspectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑图库免检配置
//
// @param request - UpdateImageLibFreeInspectionRequest
//
// @return UpdateImageLibFreeInspectionResponse
func (client *Client) UpdateImageLibFreeInspection(request *UpdateImageLibFreeInspectionRequest) (_result *UpdateImageLibFreeInspectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateImageLibFreeInspectionResponse{}
	_body, _err := client.UpdateImageLibFreeInspectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑关键词库
//
// @param request - UpdateKeywordLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKeywordLibResponse
func (client *Client) UpdateKeywordLibWithOptions(request *UpdateKeywordLibRequest, runtime *util.RuntimeOptions) (_result *UpdateKeywordLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibId)) {
		body["LibId"] = request.LibId
	}

	if !tea.BoolValue(util.IsUnset(request.LibName)) {
		body["LibName"] = request.LibName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateKeywordLib"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateKeywordLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑关键词库
//
// @param request - UpdateKeywordLibRequest
//
// @return UpdateKeywordLibResponse
func (client *Client) UpdateKeywordLib(request *UpdateKeywordLibRequest) (_result *UpdateKeywordLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateKeywordLibResponse{}
	_body, _err := client.UpdateKeywordLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检测结果反馈
//
// @param request - UpdateScanResultFeedbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScanResultFeedbackResponse
func (client *Client) UpdateScanResultFeedbackWithOptions(request *UpdateScanResultFeedbackRequest, runtime *util.RuntimeOptions) (_result *UpdateScanResultFeedbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Feedback)) {
		body["Feedback"] = request.Feedback
	}

	if !tea.BoolValue(util.IsUnset(request.QueryRequestId)) {
		body["QueryRequestId"] = request.QueryRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateScanResultFeedback"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateScanResultFeedbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检测结果反馈
//
// @param request - UpdateScanResultFeedbackRequest
//
// @return UpdateScanResultFeedbackResponse
func (client *Client) UpdateScanResultFeedback(request *UpdateScanResultFeedbackRequest) (_result *UpdateScanResultFeedbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateScanResultFeedbackResponse{}
	_body, _err := client.UpdateScanResultFeedbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新服务
//
// @param request - UpdateServiceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceConfigResponse
func (client *Client) UpdateServiceConfigWithOptions(request *UpdateServiceConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileConfig)) {
		body["FileConfig"] = request.FileConfig
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordFilterLibs)) {
		body["KeywordFilterLibs"] = request.KeywordFilterLibs
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordHitLibs)) {
		body["KeywordHitLibs"] = request.KeywordHitLibs
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.SceneConfig)) {
		body["SceneConfig"] = request.SceneConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.VideoConfig)) {
		body["VideoConfig"] = request.VideoConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceConfig"),
		Version:     tea.String("2022-09-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新服务
//
// @param request - UpdateServiceConfigRequest
//
// @return UpdateServiceConfigResponse
func (client *Client) UpdateServiceConfig(request *UpdateServiceConfigRequest) (_result *UpdateServiceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServiceConfigResponse{}
	_body, _err := client.UpdateServiceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
