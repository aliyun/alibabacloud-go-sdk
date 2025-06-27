// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeFileModerationResultRequest struct {
	// The type of the moderation service.
	//
	// example:
	//
	// document_detection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// 	- taskId: required. The URL of the object that you want to moderate. Make sure that the URL can be accessed over the Internet.
	//
	// example:
	//
	// {\\"taskId\\":\\"vi_f_hPgx9PFIQISdlfA888hOFG-1yJq8v\\"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s DescribeFileModerationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileModerationResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultRequest) SetService(v string) *DescribeFileModerationResultRequest {
	s.Service = &v
	return s
}

func (s *DescribeFileModerationResultRequest) SetServiceParameters(v string) *DescribeFileModerationResultRequest {
	s.ServiceParameters = &v
	return s
}

type DescribeFileModerationResultResponseBody struct {
	// The returned HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribeFileModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6CF2815C-C8C7-4A01-B52E-FF6E24F53492
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFileModerationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBody) SetCode(v int32) *DescribeFileModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFileModerationResultResponseBody) SetData(v *DescribeFileModerationResultResponseBodyData) *DescribeFileModerationResultResponseBody {
	s.Data = v
	return s
}

func (s *DescribeFileModerationResultResponseBody) SetMessage(v string) *DescribeFileModerationResultResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFileModerationResultResponseBody) SetRequestId(v string) *DescribeFileModerationResultResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFileModerationResultResponseBodyData struct {
	// The ID of the moderated object.
	//
	// example:
	//
	// 26769ada6e264e7ba9aa048241e12be9
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// Optional. The document type.
	//
	// example:
	//
	// doc
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// An array that consists of the moderation results.
	PageResult []*DescribeFileModerationResultResponseBodyDataPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Repeated"`
	// Summary of results
	PageSummary *DescribeFileModerationResultResponseBodyDataPageSummary `json:"PageSummary,omitempty" xml:"PageSummary,omitempty" type:"Struct"`
	// Risk Level
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The URL of the moderated object.
	//
	// example:
	//
	// https://detect-obj.oss-cn-hangzhou.aliyuncs.com/sample/xxxx.pdf
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeFileModerationResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyData) SetDataId(v string) *DescribeFileModerationResultResponseBodyData {
	s.DataId = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyData) SetDocType(v string) *DescribeFileModerationResultResponseBodyData {
	s.DocType = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyData) SetPageResult(v []*DescribeFileModerationResultResponseBodyDataPageResult) *DescribeFileModerationResultResponseBodyData {
	s.PageResult = v
	return s
}

func (s *DescribeFileModerationResultResponseBodyData) SetPageSummary(v *DescribeFileModerationResultResponseBodyDataPageSummary) *DescribeFileModerationResultResponseBodyData {
	s.PageSummary = v
	return s
}

func (s *DescribeFileModerationResultResponseBodyData) SetRiskLevel(v string) *DescribeFileModerationResultResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyData) SetUrl(v string) *DescribeFileModerationResultResponseBodyData {
	s.Url = &v
	return s
}

type DescribeFileModerationResultResponseBodyDataPageResult struct {
	// The image moderation results.
	ImageResult []*DescribeFileModerationResultResponseBodyDataPageResultImageResult `json:"ImageResult,omitempty" xml:"ImageResult,omitempty" type:"Repeated"`
	// The image URL.
	//
	// example:
	//
	// https://detect-obj.oss-cn-hangzhou.aliyuncs.com/sample/xxxx.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The text moderation results.
	TextResult []*DescribeFileModerationResultResponseBodyDataPageResultTextResult `json:"TextResult,omitempty" xml:"TextResult,omitempty" type:"Repeated"`
	// The text URL.
	//
	// example:
	//
	// https://detect-obj.oss-cn-hangzhou.aliyuncs.com/sample/xxxx.txt
	TextUrl *string `json:"TextUrl,omitempty" xml:"TextUrl,omitempty"`
}

func (s DescribeFileModerationResultResponseBodyDataPageResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageResult) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageResult) SetImageResult(v []*DescribeFileModerationResultResponseBodyDataPageResultImageResult) *DescribeFileModerationResultResponseBodyDataPageResult {
	s.ImageResult = v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResult) SetImageUrl(v string) *DescribeFileModerationResultResponseBodyDataPageResult {
	s.ImageUrl = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResult) SetPageNum(v int32) *DescribeFileModerationResultResponseBodyDataPageResult {
	s.PageNum = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResult) SetTextResult(v []*DescribeFileModerationResultResponseBodyDataPageResultTextResult) *DescribeFileModerationResultResponseBodyDataPageResult {
	s.TextResult = v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResult) SetTextUrl(v string) *DescribeFileModerationResultResponseBodyDataPageResult {
	s.TextUrl = &v
	return s
}

type DescribeFileModerationResultResponseBodyDataPageResultImageResult struct {
	// The description.
	//
	// example:
	//
	// This is a title.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label information.
	LabelResult []*DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult `json:"LabelResult,omitempty" xml:"LabelResult,omitempty" type:"Repeated"`
	// The location information
	Location *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// Risk Level
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The moderation service.
	//
	// example:
	//
	// baselineCheck
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s DescribeFileModerationResultResponseBodyDataPageResultImageResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageResultImageResult) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResult) SetDescription(v string) *DescribeFileModerationResultResponseBodyDataPageResultImageResult {
	s.Description = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResult) SetLabelResult(v []*DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult) *DescribeFileModerationResultResponseBodyDataPageResultImageResult {
	s.LabelResult = v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResult) SetLocation(v *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation) *DescribeFileModerationResultResponseBodyDataPageResultImageResult {
	s.Location = v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResult) SetRiskLevel(v string) *DescribeFileModerationResultResponseBodyDataPageResultImageResult {
	s.RiskLevel = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResult) SetService(v string) *DescribeFileModerationResultResponseBodyDataPageResultImageResult {
	s.Service = &v
	return s
}

type DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult struct {
	// The score of the confidence level. Valid values: 0 to 100. The value is accurate to two decimal places.
	//
	// example:
	//
	// 25.0
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The description.
	//
	// example:
	//
	// This is a title.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The details of the labels.
	//
	// example:
	//
	// nonlabel
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult) SetConfidence(v float32) *DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult {
	s.Confidence = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult) SetDescription(v string) *DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult {
	s.Description = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult) SetLabel(v string) *DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult {
	s.Label = &v
	return s
}

type DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation struct {
	// The H value of the coordinate point.
	//
	// example:
	//
	// 44
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The W value of the coordinate point.
	//
	// example:
	//
	// 33
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The X value of the coordinate point.
	//
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The Y value of the coordinate point.
	//
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation) SetH(v int32) *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation {
	s.H = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation) SetW(v int32) *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation {
	s.W = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation) SetX(v int32) *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation {
	s.X = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation) SetY(v int32) *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation {
	s.Y = &v
	return s
}

type DescribeFileModerationResultResponseBodyDataPageResultTextResult struct {
	// The description.
	//
	// example:
	//
	// This is a title.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The description of the labels.
	//
	// example:
	//
	// no risk
	Descriptions *string `json:"Descriptions,omitempty" xml:"Descriptions,omitempty"`
	// The details of the labels.
	//
	// example:
	//
	// porn
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// Risk Level
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The risk details that are hit.
	//
	// example:
	//
	// xxx
	RiskTips *string `json:"RiskTips,omitempty" xml:"RiskTips,omitempty"`
	// The risk words that are hit.
	//
	// example:
	//
	// xxx
	RiskWords *string `json:"RiskWords,omitempty" xml:"RiskWords,omitempty"`
	// The moderation service.
	//
	// example:
	//
	// chat_detection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The text content.
	//
	// example:
	//
	// This is a text.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The text segmentation information.
	//
	// example:
	//
	// [0,999]
	TextSegment *string `json:"TextSegment,omitempty" xml:"TextSegment,omitempty"`
}

func (s DescribeFileModerationResultResponseBodyDataPageResultTextResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageResultTextResult) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) SetDescription(v string) *DescribeFileModerationResultResponseBodyDataPageResultTextResult {
	s.Description = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) SetDescriptions(v string) *DescribeFileModerationResultResponseBodyDataPageResultTextResult {
	s.Descriptions = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) SetLabels(v string) *DescribeFileModerationResultResponseBodyDataPageResultTextResult {
	s.Labels = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) SetRiskLevel(v string) *DescribeFileModerationResultResponseBodyDataPageResultTextResult {
	s.RiskLevel = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) SetRiskTips(v string) *DescribeFileModerationResultResponseBodyDataPageResultTextResult {
	s.RiskTips = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) SetRiskWords(v string) *DescribeFileModerationResultResponseBodyDataPageResultTextResult {
	s.RiskWords = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) SetService(v string) *DescribeFileModerationResultResponseBodyDataPageResultTextResult {
	s.Service = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) SetText(v string) *DescribeFileModerationResultResponseBodyDataPageResultTextResult {
	s.Text = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) SetTextSegment(v string) *DescribeFileModerationResultResponseBodyDataPageResultTextResult {
	s.TextSegment = &v
	return s
}

type DescribeFileModerationResultResponseBodyDataPageSummary struct {
	// Image Results Summary
	ImageSummary *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary `json:"ImageSummary,omitempty" xml:"ImageSummary,omitempty" type:"Struct"`
	// Number of pages
	//
	// example:
	//
	// 1
	PageSum *int32 `json:"PageSum,omitempty" xml:"PageSum,omitempty"`
	// Text Results Summary
	TextSummary *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary `json:"TextSummary,omitempty" xml:"TextSummary,omitempty" type:"Struct"`
}

func (s DescribeFileModerationResultResponseBodyDataPageSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageSummary) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummary) SetImageSummary(v *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary) *DescribeFileModerationResultResponseBodyDataPageSummary {
	s.ImageSummary = v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummary) SetPageSum(v int32) *DescribeFileModerationResultResponseBodyDataPageSummary {
	s.PageSum = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummary) SetTextSummary(v *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary) *DescribeFileModerationResultResponseBodyDataPageSummary {
	s.TextSummary = v
	return s
}

type DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary struct {
	// Image Label
	ImageLabels []*DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels `json:"ImageLabels,omitempty" xml:"ImageLabels,omitempty" type:"Repeated"`
	// Risk Level
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary) SetImageLabels(v []*DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels) *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary {
	s.ImageLabels = v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary) SetRiskLevel(v string) *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary {
	s.RiskLevel = &v
	return s
}

type DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels struct {
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The details of the labels.
	//
	// example:
	//
	// contraband
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The number of times that the label is matched.
	//
	// example:
	//
	// 1
	LabelSum *int32 `json:"LabelSum,omitempty" xml:"LabelSum,omitempty"`
}

func (s DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels) SetDescription(v string) *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels {
	s.Description = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels) SetLabel(v string) *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels {
	s.Label = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels) SetLabelSum(v int32) *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels {
	s.LabelSum = &v
	return s
}

type DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary struct {
	// Risk Level
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// Text Label
	TextLabels []*DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels `json:"TextLabels,omitempty" xml:"TextLabels,omitempty" type:"Repeated"`
}

func (s DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary) SetRiskLevel(v string) *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary {
	s.RiskLevel = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary) SetTextLabels(v []*DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels) *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary {
	s.TextLabels = v
	return s
}

type DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels struct {
	// The description of the labels.
	//
	// example:
	//
	// no risk
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The details of the labels.
	//
	// example:
	//
	// contraband
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The number of times that the label is matched.
	//
	// example:
	//
	// 1
	LabelSum *int32 `json:"LabelSum,omitempty" xml:"LabelSum,omitempty"`
}

func (s DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels) SetDescription(v string) *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels {
	s.Description = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels) SetLabel(v string) *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels {
	s.Label = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels) SetLabelSum(v int32) *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels {
	s.LabelSum = &v
	return s
}

type DescribeFileModerationResultResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFileModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFileModerationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileModerationResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponse) SetHeaders(v map[string]*string) *DescribeFileModerationResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileModerationResultResponse) SetStatusCode(v int32) *DescribeFileModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileModerationResultResponse) SetBody(v *DescribeFileModerationResultResponseBody) *DescribeFileModerationResultResponse {
	s.Body = v
	return s
}

type DescribeImageModerationResultRequest struct {
	// The ReqId field returned by the asynchronous Image Moderation 2.0 API.
	//
	// example:
	//
	// B0963D30-BAB4-562F-9ED0-7A23AEC51C7C
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
}

func (s DescribeImageModerationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageModerationResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageModerationResultRequest) SetReqId(v string) *DescribeImageModerationResultRequest {
	s.ReqId = &v
	return s
}

type DescribeImageModerationResultResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The image moderation results.
	Data *DescribeImageModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2881AD4F-638B-52A3-BA20-F74C5B1CEAE3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageModerationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageModerationResultResponseBody) SetCode(v int32) *DescribeImageModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImageModerationResultResponseBody) SetData(v *DescribeImageModerationResultResponseBodyData) *DescribeImageModerationResultResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImageModerationResultResponseBody) SetMsg(v string) *DescribeImageModerationResultResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeImageModerationResultResponseBody) SetRequestId(v string) *DescribeImageModerationResultResponseBody {
	s.RequestId = &v
	return s
}

type DescribeImageModerationResultResponseBodyData struct {
	// The value of dataId that is specified in the API request. If this parameter is not specified in the API request, this field is not available in the response.
	//
	// example:
	//
	// 2a5389eb-4ff8-4584-ac99-644e2a539aa1
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The information about the captured frames.
	//
	// example:
	//
	// [{"result":[{"confidence":81.22,"label":"violent_explosion"}]}]
	Frame *string `json:"Frame,omitempty" xml:"Frame,omitempty"`
	// The number of frames.
	//
	// example:
	//
	// 1
	FrameNum     *int32  `json:"FrameNum,omitempty" xml:"FrameNum,omitempty"`
	ManualTaskId *string `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// The reqId field returned by the Image Async Moderation API.
	//
	// example:
	//
	// B0963D30-BAB4-562F-9ED0-7A23AEC51C7C
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	// The results of image moderation parameters such as the label parameter and the confidence parameter.
	Result []*DescribeImageModerationResultResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Risk Level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeImageModerationResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageModerationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImageModerationResultResponseBodyData) SetDataId(v string) *DescribeImageModerationResultResponseBodyData {
	s.DataId = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyData) SetFrame(v string) *DescribeImageModerationResultResponseBodyData {
	s.Frame = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyData) SetFrameNum(v int32) *DescribeImageModerationResultResponseBodyData {
	s.FrameNum = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyData) SetManualTaskId(v string) *DescribeImageModerationResultResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyData) SetReqId(v string) *DescribeImageModerationResultResponseBodyData {
	s.ReqId = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyData) SetResult(v []*DescribeImageModerationResultResponseBodyDataResult) *DescribeImageModerationResultResponseBodyData {
	s.Result = v
	return s
}

func (s *DescribeImageModerationResultResponseBodyData) SetRiskLevel(v string) *DescribeImageModerationResultResponseBodyData {
	s.RiskLevel = &v
	return s
}

type DescribeImageModerationResultResponseBodyDataResult struct {
	// The score of the confidence level. Valid values: 0 to 100. The value is accurate to two decimal places.
	//
	// example:
	//
	// 81.22
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The description of the result.
	//
	// example:
	//
	// no risk
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The labels returned after the image moderation.
	//
	// example:
	//
	// violent_explosion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Risk Level
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeImageModerationResultResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageModerationResultResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DescribeImageModerationResultResponseBodyDataResult) SetConfidence(v float32) *DescribeImageModerationResultResponseBodyDataResult {
	s.Confidence = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyDataResult) SetDescription(v string) *DescribeImageModerationResultResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyDataResult) SetLabel(v string) *DescribeImageModerationResultResponseBodyDataResult {
	s.Label = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyDataResult) SetRiskLevel(v string) *DescribeImageModerationResultResponseBodyDataResult {
	s.RiskLevel = &v
	return s
}

type DescribeImageModerationResultResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageModerationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageModerationResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageModerationResultResponse) SetHeaders(v map[string]*string) *DescribeImageModerationResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageModerationResultResponse) SetStatusCode(v int32) *DescribeImageModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageModerationResultResponse) SetBody(v *DescribeImageModerationResultResponseBody) *DescribeImageModerationResultResponse {
	s.Body = v
	return s
}

type DescribeImageResultExtRequest struct {
	// The content of the information to be obtained. Multiple values are separated by commas.
	//
	// example:
	//
	// customImage,textInImage
	InfoType *string `json:"InfoType,omitempty" xml:"InfoType,omitempty"`
	// The reqId field returned by the Url Async Moderation API.
	//
	// example:
	//
	// 638EDDC65C82AB39319A9F60
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
}

func (s DescribeImageResultExtRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResultExtRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtRequest) SetInfoType(v string) *DescribeImageResultExtRequest {
	s.InfoType = &v
	return s
}

func (s *DescribeImageResultExtRequest) SetReqId(v string) *DescribeImageResultExtRequest {
	s.ReqId = &v
	return s
}

type DescribeImageResultExtResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribeImageResultExtResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6CF2815C-C8C7-4A01-B52E-FF6E24F53492
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageResultExtResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResultExtResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtResponseBody) SetCode(v int32) *DescribeImageResultExtResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImageResultExtResponseBody) SetData(v *DescribeImageResultExtResponseBodyData) *DescribeImageResultExtResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImageResultExtResponseBody) SetMsg(v string) *DescribeImageResultExtResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeImageResultExtResponseBody) SetRequestId(v string) *DescribeImageResultExtResponseBody {
	s.RequestId = &v
	return s
}

type DescribeImageResultExtResponseBodyData struct {
	// If a custom image library is hit, information about the hit custom image library is returned.
	CustomImage []*DescribeImageResultExtResponseBodyDataCustomImage `json:"CustomImage,omitempty" xml:"CustomImage,omitempty" type:"Repeated"`
	// Person information list.
	PublicFigure []*DescribeImageResultExtResponseBodyDataPublicFigure `json:"PublicFigure,omitempty" xml:"PublicFigure,omitempty" type:"Repeated"`
	// Returns the text information in the hit image.
	TextInImage *DescribeImageResultExtResponseBodyDataTextInImage `json:"TextInImage,omitempty" xml:"TextInImage,omitempty" type:"Struct"`
}

func (s DescribeImageResultExtResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResultExtResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtResponseBodyData) SetCustomImage(v []*DescribeImageResultExtResponseBodyDataCustomImage) *DescribeImageResultExtResponseBodyData {
	s.CustomImage = v
	return s
}

func (s *DescribeImageResultExtResponseBodyData) SetPublicFigure(v []*DescribeImageResultExtResponseBodyDataPublicFigure) *DescribeImageResultExtResponseBodyData {
	s.PublicFigure = v
	return s
}

func (s *DescribeImageResultExtResponseBodyData) SetTextInImage(v *DescribeImageResultExtResponseBodyDataTextInImage) *DescribeImageResultExtResponseBodyData {
	s.TextInImage = v
	return s
}

type DescribeImageResultExtResponseBodyDataCustomImage struct {
	// The image ID.
	//
	// example:
	//
	// 123456
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image library ID.
	//
	// example:
	//
	// 123456
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// The image library name.
	//
	// example:
	//
	// 图库123
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s DescribeImageResultExtResponseBodyDataCustomImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResultExtResponseBodyDataCustomImage) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtResponseBodyDataCustomImage) SetImageId(v string) *DescribeImageResultExtResponseBodyDataCustomImage {
	s.ImageId = &v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataCustomImage) SetLibId(v string) *DescribeImageResultExtResponseBodyDataCustomImage {
	s.LibId = &v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataCustomImage) SetLibName(v string) *DescribeImageResultExtResponseBodyDataCustomImage {
	s.LibName = &v
	return s
}

type DescribeImageResultExtResponseBodyDataPublicFigure struct {
	// Identified person coding information.
	//
	// example:
	//
	// yzazhzou
	FigureId *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
}

func (s DescribeImageResultExtResponseBodyDataPublicFigure) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResultExtResponseBodyDataPublicFigure) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtResponseBodyDataPublicFigure) SetFigureId(v string) *DescribeImageResultExtResponseBodyDataPublicFigure {
	s.FigureId = &v
	return s
}

type DescribeImageResultExtResponseBodyDataTextInImage struct {
	// When a custom text library is hit, the custom library ID, custom library name, and custom word are returned.
	CustomTexts []*DescribeImageResultExtResponseBodyDataTextInImageCustomTexts `json:"CustomTexts,omitempty" xml:"CustomTexts,omitempty" type:"Repeated"`
	// Returns the text information in the recognized image.
	OcrDatas []*string `json:"OcrDatas,omitempty" xml:"OcrDatas,omitempty" type:"Repeated"`
	// The risk words that are hit. Multiple words are separated by commas (,).
	RiskWords []*string `json:"RiskWords,omitempty" xml:"RiskWords,omitempty" type:"Repeated"`
}

func (s DescribeImageResultExtResponseBodyDataTextInImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResultExtResponseBodyDataTextInImage) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtResponseBodyDataTextInImage) SetCustomTexts(v []*DescribeImageResultExtResponseBodyDataTextInImageCustomTexts) *DescribeImageResultExtResponseBodyDataTextInImage {
	s.CustomTexts = v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataTextInImage) SetOcrDatas(v []*string) *DescribeImageResultExtResponseBodyDataTextInImage {
	s.OcrDatas = v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataTextInImage) SetRiskWords(v []*string) *DescribeImageResultExtResponseBodyDataTextInImage {
	s.RiskWords = v
	return s
}

type DescribeImageResultExtResponseBodyDataTextInImageCustomTexts struct {
	// Custom words, multiple words separated by commas.
	//
	// example:
	//
	// aaa,bbb
	KeyWords *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// Custom library ID.
	//
	// example:
	//
	// 123456
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// Custom library name.
	//
	// example:
	//
	// test
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s DescribeImageResultExtResponseBodyDataTextInImageCustomTexts) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResultExtResponseBodyDataTextInImageCustomTexts) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtResponseBodyDataTextInImageCustomTexts) SetKeyWords(v string) *DescribeImageResultExtResponseBodyDataTextInImageCustomTexts {
	s.KeyWords = &v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataTextInImageCustomTexts) SetLibId(v string) *DescribeImageResultExtResponseBodyDataTextInImageCustomTexts {
	s.LibId = &v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataTextInImageCustomTexts) SetLibName(v string) *DescribeImageResultExtResponseBodyDataTextInImageCustomTexts {
	s.LibName = &v
	return s
}

type DescribeImageResultExtResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageResultExtResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageResultExtResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResultExtResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtResponse) SetHeaders(v map[string]*string) *DescribeImageResultExtResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageResultExtResponse) SetStatusCode(v int32) *DescribeImageResultExtResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageResultExtResponse) SetBody(v *DescribeImageResultExtResponseBody) *DescribeImageResultExtResponse {
	s.Body = v
	return s
}

type DescribeUploadTokenResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribeUploadTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUploadTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUploadTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUploadTokenResponseBody) SetCode(v int32) *DescribeUploadTokenResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUploadTokenResponseBody) SetData(v *DescribeUploadTokenResponseBodyData) *DescribeUploadTokenResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUploadTokenResponseBody) SetMsg(v string) *DescribeUploadTokenResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeUploadTokenResponseBody) SetRequestId(v string) *DescribeUploadTokenResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUploadTokenResponseBodyData struct {
	// The AccessKey ID.
	//
	// example:
	//
	// STS.NUEUjvDqMuvH6oQA1TXxxH4wVR
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret.
	//
	// example:
	//
	// xxxx
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The bucket name.
	//
	// example:
	//
	// oss-cip-shanghai
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The time when the file sharing link expires.
	//
	// example:
	//
	// 1720577200
	Expiration *int32 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The file prefix.
	//
	// example:
	//
	// upload/1xxb89/
	FileNamePrefix *string `json:"FileNamePrefix,omitempty" xml:"FileNamePrefix,omitempty"`
	// the oss intranet point.
	//
	// example:
	//
	// https://oss-cn-shanghai-internal.aliyuncs.com
	OssInternalEndPoint *string `json:"OssInternalEndPoint,omitempty" xml:"OssInternalEndPoint,omitempty"`
	// the oss internet point.
	//
	// example:
	//
	// https://oss-cn-shanghai.aliyuncs.com
	OssInternetEndPoint *string `json:"OssInternetEndPoint,omitempty" xml:"OssInternetEndPoint,omitempty"`
	// The security token.
	//
	// example:
	//
	// xxxx
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeUploadTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeUploadTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeUploadTokenResponseBodyData) SetAccessKeyId(v string) *DescribeUploadTokenResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeUploadTokenResponseBodyData) SetAccessKeySecret(v string) *DescribeUploadTokenResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *DescribeUploadTokenResponseBodyData) SetBucketName(v string) *DescribeUploadTokenResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *DescribeUploadTokenResponseBodyData) SetExpiration(v int32) *DescribeUploadTokenResponseBodyData {
	s.Expiration = &v
	return s
}

func (s *DescribeUploadTokenResponseBodyData) SetFileNamePrefix(v string) *DescribeUploadTokenResponseBodyData {
	s.FileNamePrefix = &v
	return s
}

func (s *DescribeUploadTokenResponseBodyData) SetOssInternalEndPoint(v string) *DescribeUploadTokenResponseBodyData {
	s.OssInternalEndPoint = &v
	return s
}

func (s *DescribeUploadTokenResponseBodyData) SetOssInternetEndPoint(v string) *DescribeUploadTokenResponseBodyData {
	s.OssInternetEndPoint = &v
	return s
}

func (s *DescribeUploadTokenResponseBodyData) SetSecurityToken(v string) *DescribeUploadTokenResponseBodyData {
	s.SecurityToken = &v
	return s
}

type DescribeUploadTokenResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUploadTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUploadTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUploadTokenResponse) GoString() string {
	return s.String()
}

func (s *DescribeUploadTokenResponse) SetHeaders(v map[string]*string) *DescribeUploadTokenResponse {
	s.Headers = v
	return s
}

func (s *DescribeUploadTokenResponse) SetStatusCode(v int32) *DescribeUploadTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUploadTokenResponse) SetBody(v *DescribeUploadTokenResponseBody) *DescribeUploadTokenResponse {
	s.Body = v
	return s
}

type DescribeUrlModerationResultRequest struct {
	// The ReqId field returned by an asynchronous URL moderation operation.
	//
	// example:
	//
	// B0963D30-BAB4-562F-9ED0-7A23AEC51C7C
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
}

func (s DescribeUrlModerationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUrlModerationResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeUrlModerationResultRequest) SetReqId(v string) *DescribeUrlModerationResultRequest {
	s.ReqId = &v
	return s
}

type DescribeUrlModerationResultResponseBody struct {
	// The returned HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribeUrlModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01F9144A-2088-5D87-935B-2DB865284B1A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUrlModerationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUrlModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUrlModerationResultResponseBody) SetCode(v int32) *DescribeUrlModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBody) SetData(v *DescribeUrlModerationResultResponseBodyData) *DescribeUrlModerationResultResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUrlModerationResultResponseBody) SetMsg(v string) *DescribeUrlModerationResultResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBody) SetRequestId(v string) *DescribeUrlModerationResultResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUrlModerationResultResponseBodyData struct {
	// The value of dataId that is specified in the API request. If this parameter is not specified in the API request, this field is not available in the response.
	//
	// example:
	//
	// 26769ada6e264e7ba9aa048241e12be9
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The supplementary information.
	ExtraInfo *DescribeUrlModerationResultResponseBodyDataExtraInfo `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty" type:"Struct"`
	// The ReqId field returned by an asynchronous URL moderation operation.
	//
	// example:
	//
	// B0963D30-BAB4-562F-9ED0-7A23AEC51C7C
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	// The returned results.
	Result []*DescribeUrlModerationResultResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeUrlModerationResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeUrlModerationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeUrlModerationResultResponseBodyData) SetDataId(v string) *DescribeUrlModerationResultResponseBodyData {
	s.DataId = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBodyData) SetExtraInfo(v *DescribeUrlModerationResultResponseBodyDataExtraInfo) *DescribeUrlModerationResultResponseBodyData {
	s.ExtraInfo = v
	return s
}

func (s *DescribeUrlModerationResultResponseBodyData) SetReqId(v string) *DescribeUrlModerationResultResponseBodyData {
	s.ReqId = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBodyData) SetResult(v []*DescribeUrlModerationResultResponseBodyDataResult) *DescribeUrlModerationResultResponseBodyData {
	s.Result = v
	return s
}

type DescribeUrlModerationResultResponseBodyDataExtraInfo struct {
	// The ICP number.
	//
	// example:
	//
	// xx
	IcpNo *string `json:"IcpNo,omitempty" xml:"IcpNo,omitempty"`
	// The type of the ICP filing.
	//
	// example:
	//
	// xx
	IcpType *string `json:"IcpType,omitempty" xml:"IcpType,omitempty"`
	// The type of site
	//
	// example:
	//
	// game
	SiteType *string `json:"SiteType,omitempty" xml:"SiteType,omitempty"`
}

func (s DescribeUrlModerationResultResponseBodyDataExtraInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUrlModerationResultResponseBodyDataExtraInfo) GoString() string {
	return s.String()
}

func (s *DescribeUrlModerationResultResponseBodyDataExtraInfo) SetIcpNo(v string) *DescribeUrlModerationResultResponseBodyDataExtraInfo {
	s.IcpNo = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBodyDataExtraInfo) SetIcpType(v string) *DescribeUrlModerationResultResponseBodyDataExtraInfo {
	s.IcpType = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBodyDataExtraInfo) SetSiteType(v string) *DescribeUrlModerationResultResponseBodyDataExtraInfo {
	s.SiteType = &v
	return s
}

type DescribeUrlModerationResultResponseBodyDataResult struct {
	// The score of the confidence level. Valid values: 0 to 100. The value is accurate to two decimal places.
	//
	// example:
	//
	// 81.22
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The labels returned after the asynchronous URL moderation.
	//
	// example:
	//
	// sexual_url
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s DescribeUrlModerationResultResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeUrlModerationResultResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DescribeUrlModerationResultResponseBodyDataResult) SetConfidence(v float32) *DescribeUrlModerationResultResponseBodyDataResult {
	s.Confidence = &v
	return s
}

func (s *DescribeUrlModerationResultResponseBodyDataResult) SetLabel(v string) *DescribeUrlModerationResultResponseBodyDataResult {
	s.Label = &v
	return s
}

type DescribeUrlModerationResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUrlModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUrlModerationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUrlModerationResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeUrlModerationResultResponse) SetHeaders(v map[string]*string) *DescribeUrlModerationResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeUrlModerationResultResponse) SetStatusCode(v int32) *DescribeUrlModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUrlModerationResultResponse) SetBody(v *DescribeUrlModerationResultResponseBody) *DescribeUrlModerationResultResponse {
	s.Body = v
	return s
}

type FileModerationRequest struct {
	// The type of the moderation service.
	//
	// example:
	//
	// document_detection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// example:
	//
	// {"url":"https://detect-obj.oss-cn-hangzhou.aliyuncs.com/sample/xxxx.pdf"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s FileModerationRequest) String() string {
	return tea.Prettify(s)
}

func (s FileModerationRequest) GoString() string {
	return s.String()
}

func (s *FileModerationRequest) SetService(v string) *FileModerationRequest {
	s.Service = &v
	return s
}

func (s *FileModerationRequest) SetServiceParameters(v string) *FileModerationRequest {
	s.ServiceParameters = &v
	return s
}

type FileModerationResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *FileModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FileModerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FileModerationResponseBody) GoString() string {
	return s.String()
}

func (s *FileModerationResponseBody) SetCode(v int32) *FileModerationResponseBody {
	s.Code = &v
	return s
}

func (s *FileModerationResponseBody) SetData(v *FileModerationResponseBodyData) *FileModerationResponseBody {
	s.Data = v
	return s
}

func (s *FileModerationResponseBody) SetMessage(v string) *FileModerationResponseBody {
	s.Message = &v
	return s
}

func (s *FileModerationResponseBody) SetRequestId(v string) *FileModerationResponseBody {
	s.RequestId = &v
	return s
}

type FileModerationResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FileModerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FileModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *FileModerationResponseBodyData) SetTaskId(v string) *FileModerationResponseBodyData {
	s.TaskId = &v
	return s
}

type FileModerationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FileModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FileModerationResponse) String() string {
	return tea.Prettify(s)
}

func (s FileModerationResponse) GoString() string {
	return s.String()
}

func (s *FileModerationResponse) SetHeaders(v map[string]*string) *FileModerationResponse {
	s.Headers = v
	return s
}

func (s *FileModerationResponse) SetStatusCode(v int32) *FileModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *FileModerationResponse) SetBody(v *FileModerationResponseBody) *FileModerationResponse {
	s.Body = v
	return s
}

type ImageAsyncModerationRequest struct {
	// The type of the moderation service.
	//
	// example:
	//
	// baselineCheck
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// example:
	//
	// {"imageUrl":"https://img.alicdn.com/tfs/TB1U4r9AeH2gK0jSZJnXXaT1FXa-2880-480.png","dataId":"img123****"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ImageAsyncModerationRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageAsyncModerationRequest) GoString() string {
	return s.String()
}

func (s *ImageAsyncModerationRequest) SetService(v string) *ImageAsyncModerationRequest {
	s.Service = &v
	return s
}

func (s *ImageAsyncModerationRequest) SetServiceParameters(v string) *ImageAsyncModerationRequest {
	s.ServiceParameters = &v
	return s
}

type ImageAsyncModerationResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ImageAsyncModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4A926AE2-4C96-573F-824F-0532960799F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageAsyncModerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImageAsyncModerationResponseBody) GoString() string {
	return s.String()
}

func (s *ImageAsyncModerationResponseBody) SetCode(v int32) *ImageAsyncModerationResponseBody {
	s.Code = &v
	return s
}

func (s *ImageAsyncModerationResponseBody) SetData(v *ImageAsyncModerationResponseBodyData) *ImageAsyncModerationResponseBody {
	s.Data = v
	return s
}

func (s *ImageAsyncModerationResponseBody) SetMsg(v string) *ImageAsyncModerationResponseBody {
	s.Msg = &v
	return s
}

func (s *ImageAsyncModerationResponseBody) SetRequestId(v string) *ImageAsyncModerationResponseBody {
	s.RequestId = &v
	return s
}

type ImageAsyncModerationResponseBodyData struct {
	// The ID of the moderated object.
	//
	// example:
	//
	// fb5ffab1-993b-449f-b8d6-b97d5e3331f2
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The reqId field returned by the Image Async Moderation API. You can use this field to query the detection results.
	//
	// example:
	//
	// A07B3DB9-D762-5C56-95B1-8EC55CF176D2
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
}

func (s ImageAsyncModerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImageAsyncModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageAsyncModerationResponseBodyData) SetDataId(v string) *ImageAsyncModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *ImageAsyncModerationResponseBodyData) SetReqId(v string) *ImageAsyncModerationResponseBodyData {
	s.ReqId = &v
	return s
}

type ImageAsyncModerationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageAsyncModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageAsyncModerationResponse) String() string {
	return tea.Prettify(s)
}

func (s ImageAsyncModerationResponse) GoString() string {
	return s.String()
}

func (s *ImageAsyncModerationResponse) SetHeaders(v map[string]*string) *ImageAsyncModerationResponse {
	s.Headers = v
	return s
}

func (s *ImageAsyncModerationResponse) SetStatusCode(v int32) *ImageAsyncModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageAsyncModerationResponse) SetBody(v *ImageAsyncModerationResponseBody) *ImageAsyncModerationResponse {
	s.Body = v
	return s
}

type ImageBatchModerationRequest struct {
	// The types of detection supported by the enhanced image review, separated by English commas. Values:
	//
	// - baselineCheck：General Baseline Detection
	//
	// - baselineCheck_pro：General Baseline Detection_Pro Edition
	//
	// - tonalityImprove：Content governance monitoring
	//
	// - aigcCheck：AIGC image detection
	//
	// example:
	//
	// baselineCheck,tonalityImprove
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The set of relevant parameters for content detection objects.
	//
	// example:
	//
	// {
	//
	//         "imageUrl": "https://img.alicdn.com/tfs/TB1U4r9AeH2gK0jSZJnXXaT1FXa-2880-480.png",
	//
	//         "dataId": "img123****"
	//
	//     }
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ImageBatchModerationRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationRequest) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationRequest) SetService(v string) *ImageBatchModerationRequest {
	s.Service = &v
	return s
}

func (s *ImageBatchModerationRequest) SetServiceParameters(v string) *ImageBatchModerationRequest {
	s.ServiceParameters = &v
	return s
}

type ImageBatchModerationResponseBody struct {
	// Return code. A return of 200 represents success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of image content detection.
	Data *ImageBatchModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message for this request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The ID of this invocation request, generated by Alibaba Cloud as a unique identifier for the request, can be used for troubleshooting and pinpointing issues.
	//
	// example:
	//
	// 6CF2815C-C8C7-4A01-B52E-FF6E24F53492
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageBatchModerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponseBody) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBody) SetCode(v int32) *ImageBatchModerationResponseBody {
	s.Code = &v
	return s
}

func (s *ImageBatchModerationResponseBody) SetData(v *ImageBatchModerationResponseBodyData) *ImageBatchModerationResponseBody {
	s.Data = v
	return s
}

func (s *ImageBatchModerationResponseBody) SetMsg(v string) *ImageBatchModerationResponseBody {
	s.Msg = &v
	return s
}

func (s *ImageBatchModerationResponseBody) SetRequestId(v string) *ImageBatchModerationResponseBody {
	s.RequestId = &v
	return s
}

type ImageBatchModerationResponseBodyData struct {
	// To detect the data ID corresponding to the object.
	//
	// example:
	//
	// 26769ada6e264e7ba9aa048241e12be9
	DataId       *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	ManualTaskId *string `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// The risk labels, confidence scores, and other parameters of image detection results, in an array structure.
	Result []*ImageBatchModerationResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The risk labels, confidence scores, and other parameters for each service\\"s image detection, in an array structure.
	Results []*ImageBatchModerationResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// Risk level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ImageBatchModerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyData) SetDataId(v string) *ImageBatchModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *ImageBatchModerationResponseBodyData) SetManualTaskId(v string) *ImageBatchModerationResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *ImageBatchModerationResponseBodyData) SetResult(v []*ImageBatchModerationResponseBodyDataResult) *ImageBatchModerationResponseBodyData {
	s.Result = v
	return s
}

func (s *ImageBatchModerationResponseBodyData) SetResults(v []*ImageBatchModerationResponseBodyDataResults) *ImageBatchModerationResponseBodyData {
	s.Results = v
	return s
}

func (s *ImageBatchModerationResponseBodyData) SetRiskLevel(v string) *ImageBatchModerationResponseBodyData {
	s.RiskLevel = &v
	return s
}

type ImageBatchModerationResponseBodyDataResult struct {
	// Confidence score, ranging from 0 to 100, retained to two decimal places. Some labels do not have a confidence score.
	//
	// example:
	//
	// 81.22
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// Description.
	//
	// example:
	//
	// No risk detected
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The labels returned after image content detection. A single image may be associated with multiple labels and corresponding scores.
	//
	// example:
	//
	// violent_explosion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResult) SetConfidence(v float32) *ImageBatchModerationResponseBodyDataResult {
	s.Confidence = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResult) SetDescription(v string) *ImageBatchModerationResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResult) SetLabel(v string) *ImageBatchModerationResponseBodyDataResult {
	s.Label = &v
	return s
}

type ImageBatchModerationResponseBodyDataResults struct {
	// Image supplementary reference information.
	Ext *ImageBatchModerationResponseBodyDataResultsExt `json:"Ext,omitempty" xml:"Ext,omitempty" type:"Struct"`
	// The risk labels, confidence scores, and other parameters of image detection results, in an array structure.
	Result []*ImageBatchModerationResponseBodyDataResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Risk level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The enhanced image detection service supports various detection services.
	//
	// example:
	//
	// baselineCheck
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResults) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResults) SetExt(v *ImageBatchModerationResponseBodyDataResultsExt) *ImageBatchModerationResponseBodyDataResults {
	s.Ext = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResults) SetResult(v []*ImageBatchModerationResponseBodyDataResultsResult) *ImageBatchModerationResponseBodyDataResults {
	s.Result = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResults) SetRiskLevel(v string) *ImageBatchModerationResponseBodyDataResults {
	s.RiskLevel = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResults) SetService(v string) *ImageBatchModerationResponseBodyDataResults {
	s.Service = &v
	return s
}

type ImageBatchModerationResponseBodyDataResultsExt struct {
	// Custom image library hit information list.
	CustomImage []*ImageBatchModerationResponseBodyDataResultsExtCustomImage `json:"CustomImage,omitempty" xml:"CustomImage,omitempty" type:"Repeated"`
	// Logo identification information.
	LogoData *ImageBatchModerationResponseBodyDataResultsExtLogoData `json:"LogoData,omitempty" xml:"LogoData,omitempty" type:"Struct"`
	// List of character information.
	PublicFigure []*ImageBatchModerationResponseBodyDataResultsExtPublicFigure `json:"PublicFigure,omitempty" xml:"PublicFigure,omitempty" type:"Repeated"`
	// Return the text information from the recognized images.
	TextInImage *ImageBatchModerationResponseBodyDataResultsExtTextInImage `json:"TextInImage,omitempty" xml:"TextInImage,omitempty" type:"Struct"`
}

func (s ImageBatchModerationResponseBodyDataResultsExt) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExt) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExt) SetCustomImage(v []*ImageBatchModerationResponseBodyDataResultsExtCustomImage) *ImageBatchModerationResponseBodyDataResultsExt {
	s.CustomImage = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExt) SetLogoData(v *ImageBatchModerationResponseBodyDataResultsExtLogoData) *ImageBatchModerationResponseBodyDataResultsExt {
	s.LogoData = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExt) SetPublicFigure(v []*ImageBatchModerationResponseBodyDataResultsExtPublicFigure) *ImageBatchModerationResponseBodyDataResultsExt {
	s.PublicFigure = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExt) SetTextInImage(v *ImageBatchModerationResponseBodyDataResultsExtTextInImage) *ImageBatchModerationResponseBodyDataResultsExt {
	s.TextInImage = v
	return s
}

type ImageBatchModerationResponseBodyDataResultsExtCustomImage struct {
	// The ID of the hit custom image.
	//
	// example:
	//
	// 1965304870002
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Custom library ID
	//
	// example:
	//
	// 1965304870002
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// The name of the hit custom gallery.
	//
	// example:
	//
	// Whitelist
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtCustomImage) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtCustomImage) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtCustomImage) SetImageId(v string) *ImageBatchModerationResponseBodyDataResultsExtCustomImage {
	s.ImageId = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtCustomImage) SetLibId(v string) *ImageBatchModerationResponseBodyDataResultsExtCustomImage {
	s.LibId = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtCustomImage) SetLibName(v string) *ImageBatchModerationResponseBodyDataResultsExtCustomImage {
	s.LibName = &v
	return s
}

type ImageBatchModerationResponseBodyDataResultsExtLogoData struct {
	// The location information of the identifier.
	Location *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// identification information
	Logo []*ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo `json:"Logo,omitempty" xml:"Logo,omitempty" type:"Repeated"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtLogoData) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtLogoData) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoData) SetLocation(v *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) *ImageBatchModerationResponseBodyDataResultsExtLogoData {
	s.Location = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoData) SetLogo(v []*ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo) *ImageBatchModerationResponseBodyDataResultsExtLogoData {
	s.Logo = v
	return s
}

type ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation struct {
	// The width of the text area, in pixels.
	//
	// example:
	//
	// 440
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The height of the text area, in pixels.
	//
	// example:
	//
	// 330
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The distance from the top-left corner of the text area to the y-axis, with the top-left corner of the image as the origin, in pixels.
	//
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The distance from the top-left corner of the text area to the x-axis, with the top-left corner of the image as the origin, in pixels.
	//
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) SetH(v int32) *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation {
	s.H = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) SetW(v int32) *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation {
	s.W = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) SetX(v int32) *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation {
	s.X = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation) SetY(v int32) *ImageBatchModerationResponseBodyDataResultsExtLogoDataLocation {
	s.Y = &v
	return s
}

type ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo struct {
	// Confidence score, from 0 to 100, rounded to two decimal places.
	//
	// example:
	//
	// 99.1
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// Identify the category.
	//
	// example:
	//
	// logo_sns
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// identifier  name.
	//
	// example:
	//
	// Alibaba Cloud
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo) SetConfidence(v float32) *ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo {
	s.Confidence = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo) SetLabel(v string) *ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo {
	s.Label = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo) SetName(v string) *ImageBatchModerationResponseBodyDataResultsExtLogoDataLogo {
	s.Name = &v
	return s
}

type ImageBatchModerationResponseBodyDataResultsExtPublicFigure struct {
	// Identify the encoded information of the person.
	//
	// example:
	//
	// 12324222
	FigureId *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
	// The identified person\\"s name information.
	//
	// example:
	//
	// xxxxx
	FigureName *string `json:"FigureName,omitempty" xml:"FigureName,omitempty"`
	// The location information of the identifier.
	Location []*ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Repeated"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtPublicFigure) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtPublicFigure) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigure) SetFigureId(v string) *ImageBatchModerationResponseBodyDataResultsExtPublicFigure {
	s.FigureId = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigure) SetFigureName(v string) *ImageBatchModerationResponseBodyDataResultsExtPublicFigure {
	s.FigureName = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigure) SetLocation(v []*ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) *ImageBatchModerationResponseBodyDataResultsExtPublicFigure {
	s.Location = v
	return s
}

type ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation struct {
	// The width of the text area, in pixels.
	//
	// example:
	//
	// 440
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The height of the text area, in pixels.
	//
	// example:
	//
	// 330
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The distance from the top-left corner of the text area to the y-axis, with the top-left corner of the image as the origin, in pixels.
	//
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The distance from the top-left corner of the text area to the x-axis, with the top-left corner of the image as the origin, in pixels.
	//
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) SetH(v int32) *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation {
	s.H = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) SetW(v int32) *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation {
	s.W = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) SetX(v int32) *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation {
	s.X = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation) SetY(v int32) *ImageBatchModerationResponseBodyDataResultsExtPublicFigureLocation {
	s.Y = &v
	return s
}

type ImageBatchModerationResponseBodyDataResultsExtTextInImage struct {
	// When a custom text library is matched, return the custom library ID, custom library name, and custom words.
	CustomText []*ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText `json:"CustomText,omitempty" xml:"CustomText,omitempty" type:"Repeated"`
	// Return the text information of each line recognized in the image.
	OcrResult []*ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult `json:"OcrResult,omitempty" xml:"OcrResult,omitempty" type:"Repeated"`
	// hit risk keywords
	RiskWord []*string `json:"RiskWord,omitempty" xml:"RiskWord,omitempty" type:"Repeated"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtTextInImage) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtTextInImage) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImage) SetCustomText(v []*ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText) *ImageBatchModerationResponseBodyDataResultsExtTextInImage {
	s.CustomText = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImage) SetOcrResult(v []*ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult) *ImageBatchModerationResponseBodyDataResultsExtTextInImage {
	s.OcrResult = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImage) SetRiskWord(v []*string) *ImageBatchModerationResponseBodyDataResultsExtTextInImage {
	s.RiskWord = v
	return s
}

type ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText struct {
	// Custom words, separate multiple words with commas.
	//
	// example:
	//
	// Custom Word 1  and  Custom Word 2
	KeyWords *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// Custom library ID.
	//
	// example:
	//
	// 123456
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// Custom library name.
	//
	// example:
	//
	// Custom Library 1
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText) SetKeyWords(v string) *ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText {
	s.KeyWords = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText) SetLibId(v string) *ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText {
	s.LibId = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText) SetLibName(v string) *ImageBatchModerationResponseBodyDataResultsExtTextInImageCustomText {
	s.LibName = &v
	return s
}

type ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult struct {
	// Text line coordinate information.
	Location *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// Text information.
	//
	// example:
	//
	// abcd
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult) SetLocation(v *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult {
	s.Location = v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult) SetText(v string) *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResult {
	s.Text = &v
	return s
}

type ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation struct {
	// The height of the text area, in pixels.
	//
	// example:
	//
	// 33
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width of the text area, in pixels.
	//
	// example:
	//
	// 44
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The distance from the top-left corner of the text area to the y-axis, with the top-left corner of the image as the origin, in pixels.
	//
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The distance from the top-left corner of the text area to the x-axis, with the top-left corner of the image as the origin, in pixels.
	//
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) SetH(v int32) *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation {
	s.H = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) SetW(v int32) *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation {
	s.W = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) SetX(v int32) *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation {
	s.X = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation) SetY(v int32) *ImageBatchModerationResponseBodyDataResultsExtTextInImageOcrResultLocation {
	s.Y = &v
	return s
}

type ImageBatchModerationResponseBodyDataResultsResult struct {
	// Confidence score, ranging from 0 to 100, rounded to two decimal places. Some labels do not have a confidence score.
	//
	// example:
	//
	// 81.22
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// Description.
	//
	// example:
	//
	// No risk detected
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The labels returned after image content detection. A single image may have multiple labels and corresponding scores detected.
	//
	// example:
	//
	// violent_explosion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ImageBatchModerationResponseBodyDataResultsResult) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponseBodyDataResultsResult) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponseBodyDataResultsResult) SetConfidence(v float32) *ImageBatchModerationResponseBodyDataResultsResult {
	s.Confidence = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsResult) SetDescription(v string) *ImageBatchModerationResponseBodyDataResultsResult {
	s.Description = &v
	return s
}

func (s *ImageBatchModerationResponseBodyDataResultsResult) SetLabel(v string) *ImageBatchModerationResponseBodyDataResultsResult {
	s.Label = &v
	return s
}

type ImageBatchModerationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageBatchModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageBatchModerationResponse) String() string {
	return tea.Prettify(s)
}

func (s ImageBatchModerationResponse) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationResponse) SetHeaders(v map[string]*string) *ImageBatchModerationResponse {
	s.Headers = v
	return s
}

func (s *ImageBatchModerationResponse) SetStatusCode(v int32) *ImageBatchModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageBatchModerationResponse) SetBody(v *ImageBatchModerationResponseBody) *ImageBatchModerationResponse {
	s.Body = v
	return s
}

type ImageModerationRequest struct {
	// The moderation services supported by Image Moderation 2.0. Valid values:
	//
	// 	- baselineCheck: common baseline moderation
	//
	// 	- baselineCheck_pro: common baseline moderation_Professional
	//
	// 	- baselineCheck_cb: common baseline moderation_For regions outside the Chinese mainland
	//
	// 	- tonalityImprove: content governance moderation
	//
	// 	- aigcCheck: AI-generated image identification
	//
	// 	- profilePhotoCheck: avatar image moderation
	//
	// 	- advertisingCheck: marketing material identification
	//
	// 	- liveStreamCheck: moderation of screenshots of videos and live streams
	//
	// Valid values:
	//
	// 	- liveStreamCheck: moderation of screenshots of videos and live streams
	//
	// 	- baselineCheck: common baseline moderation
	//
	// 	- aigcCheck: AI-generated image identification
	//
	// 	- baselineCheck_pro: common baseline moderation_Professional
	//
	// 	- advertisingCheck: marketing material identification
	//
	// 	- baselineCheck_cb: common baseline moderation_For regions outside the Chinese mainland
	//
	// 	- tonalityImprove: content governance moderation
	//
	// 	- profilePhotoCheck: avatar image moderation
	//
	// example:
	//
	// baselineCheck
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// 	- imageUrl: the URL of the object that you want to moderate. This parameter is required.
	//
	// 	- dataId: the ID of the object that you want to moderate. This parameter is optional.
	//
	// example:
	//
	// {"imageUrl":"https://www.aliyun.com/test.jpg","dataId":"img1234567"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ImageModerationRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationRequest) GoString() string {
	return s.String()
}

func (s *ImageModerationRequest) SetService(v string) *ImageModerationRequest {
	s.Service = &v
	return s
}

func (s *ImageModerationRequest) SetServiceParameters(v string) *ImageModerationRequest {
	s.ServiceParameters = &v
	return s
}

type ImageModerationResponseBody struct {
	// The returned HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The moderation results.
	Data *ImageModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6CF2815C-C8C7-4A01-B52E-FF6E24F53492
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageModerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBody) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBody) SetCode(v int32) *ImageModerationResponseBody {
	s.Code = &v
	return s
}

func (s *ImageModerationResponseBody) SetData(v *ImageModerationResponseBodyData) *ImageModerationResponseBody {
	s.Data = v
	return s
}

func (s *ImageModerationResponseBody) SetMsg(v string) *ImageModerationResponseBody {
	s.Msg = &v
	return s
}

func (s *ImageModerationResponseBody) SetRequestId(v string) *ImageModerationResponseBody {
	s.RequestId = &v
	return s
}

type ImageModerationResponseBodyData struct {
	// The ID of the moderated object.
	//
	// >  If you specify the dataId parameter in the request, the value of the dataId parameter is returned in the response.
	//
	// example:
	//
	// fb5ffab1-993b-449f-b8d6-b97d5e3331f2
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// Auxiliary reference information.
	Ext          *ImageModerationResponseBodyDataExt `json:"Ext,omitempty" xml:"Ext,omitempty" type:"Struct"`
	ManualTaskId *string                             `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// The results of image moderation parameters such as the label parameter and the confidence parameter, which are an array structure.
	Result []*ImageModerationResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Risk Level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ImageModerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyData) SetDataId(v string) *ImageModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *ImageModerationResponseBodyData) SetExt(v *ImageModerationResponseBodyDataExt) *ImageModerationResponseBodyData {
	s.Ext = v
	return s
}

func (s *ImageModerationResponseBodyData) SetManualTaskId(v string) *ImageModerationResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *ImageModerationResponseBodyData) SetResult(v []*ImageModerationResponseBodyDataResult) *ImageModerationResponseBodyData {
	s.Result = v
	return s
}

func (s *ImageModerationResponseBodyData) SetRiskLevel(v string) *ImageModerationResponseBodyData {
	s.RiskLevel = &v
	return s
}

type ImageModerationResponseBodyDataExt struct {
	// If a custom image library is hit, information about the hit custom image library is returned.
	CustomImage []*ImageModerationResponseBodyDataExtCustomImage `json:"CustomImage,omitempty" xml:"CustomImage,omitempty" type:"Repeated"`
	// The returned face attribute information
	FaceData []*ImageModerationResponseBodyDataExtFaceData `json:"FaceData,omitempty" xml:"FaceData,omitempty" type:"Repeated"`
	// Logo information.
	LogoData []*ImageModerationResponseBodyDataExtLogoData `json:"LogoData,omitempty" xml:"LogoData,omitempty" type:"Repeated"`
	// Returns the text information in the recognized image.
	OcrResult []*ImageModerationResponseBodyDataExtOcrResult `json:"OcrResult,omitempty" xml:"OcrResult,omitempty" type:"Repeated"`
	// Person information list.
	PublicFigure []*ImageModerationResponseBodyDataExtPublicFigure `json:"PublicFigure,omitempty" xml:"PublicFigure,omitempty" type:"Repeated"`
	// The result of image recognition.
	Recognition []*ImageModerationResponseBodyDataExtRecognition `json:"Recognition,omitempty" xml:"Recognition,omitempty" type:"Repeated"`
	// Returns the text information in the hit image.
	TextInImage *ImageModerationResponseBodyDataExtTextInImage `json:"TextInImage,omitempty" xml:"TextInImage,omitempty" type:"Struct"`
	// the vl output content
	VlContent *ImageModerationResponseBodyDataExtVlContent `json:"VlContent,omitempty" xml:"VlContent,omitempty" type:"Struct"`
}

func (s ImageModerationResponseBodyDataExt) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExt) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExt) SetCustomImage(v []*ImageModerationResponseBodyDataExtCustomImage) *ImageModerationResponseBodyDataExt {
	s.CustomImage = v
	return s
}

func (s *ImageModerationResponseBodyDataExt) SetFaceData(v []*ImageModerationResponseBodyDataExtFaceData) *ImageModerationResponseBodyDataExt {
	s.FaceData = v
	return s
}

func (s *ImageModerationResponseBodyDataExt) SetLogoData(v []*ImageModerationResponseBodyDataExtLogoData) *ImageModerationResponseBodyDataExt {
	s.LogoData = v
	return s
}

func (s *ImageModerationResponseBodyDataExt) SetOcrResult(v []*ImageModerationResponseBodyDataExtOcrResult) *ImageModerationResponseBodyDataExt {
	s.OcrResult = v
	return s
}

func (s *ImageModerationResponseBodyDataExt) SetPublicFigure(v []*ImageModerationResponseBodyDataExtPublicFigure) *ImageModerationResponseBodyDataExt {
	s.PublicFigure = v
	return s
}

func (s *ImageModerationResponseBodyDataExt) SetRecognition(v []*ImageModerationResponseBodyDataExtRecognition) *ImageModerationResponseBodyDataExt {
	s.Recognition = v
	return s
}

func (s *ImageModerationResponseBodyDataExt) SetTextInImage(v *ImageModerationResponseBodyDataExtTextInImage) *ImageModerationResponseBodyDataExt {
	s.TextInImage = v
	return s
}

func (s *ImageModerationResponseBodyDataExt) SetVlContent(v *ImageModerationResponseBodyDataExtVlContent) *ImageModerationResponseBodyDataExt {
	s.VlContent = v
	return s
}

type ImageModerationResponseBodyDataExtCustomImage struct {
	// The image ID.
	//
	// example:
	//
	// 123456
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image library ID.
	//
	// example:
	//
	// lib_123456
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// The image library name.
	//
	// example:
	//
	// xx
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s ImageModerationResponseBodyDataExtCustomImage) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtCustomImage) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtCustomImage) SetImageId(v string) *ImageModerationResponseBodyDataExtCustomImage {
	s.ImageId = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtCustomImage) SetLibId(v string) *ImageModerationResponseBodyDataExtCustomImage {
	s.LibId = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtCustomImage) SetLibName(v string) *ImageModerationResponseBodyDataExtCustomImage {
	s.LibName = &v
	return s
}

type ImageModerationResponseBodyDataExtFaceData struct {
	// The age recognition result.
	//
	// example:
	//
	// 18
	Age *int32 `json:"Age,omitempty" xml:"Age,omitempty"`
	// Indicates whether the recognition result of bangs is available.
	Bang *ImageModerationResponseBodyDataExtFaceDataBang `json:"Bang,omitempty" xml:"Bang,omitempty" type:"Struct"`
	// The gender recognition result.
	Gender *ImageModerationResponseBodyDataExtFaceDataGender `json:"Gender,omitempty" xml:"Gender,omitempty" type:"Struct"`
	// The recognition result of whether to wear glasses.
	//
	// - None: No glasses.
	//
	// - Wear: Wear glasses.
	//
	// - Sunglass: Wear sunglasses.
	//
	// example:
	//
	// Common
	Glasses *string `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	// The hairstyle recognition result.
	Hairstyle *ImageModerationResponseBodyDataExtFaceDataHairstyle `json:"Hairstyle,omitempty" xml:"Hairstyle,omitempty" type:"Struct"`
	// The recognition result of whether to wear a hat.
	Hat *ImageModerationResponseBodyDataExtFaceDataHat `json:"Hat,omitempty" xml:"Hat,omitempty" type:"Struct"`
	// The location of the face.
	Location *ImageModerationResponseBodyDataExtFaceDataLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// The recognition result of whether to wear a mask.
	Mask *ImageModerationResponseBodyDataExtFaceDataMask `json:"Mask,omitempty" xml:"Mask,omitempty" type:"Struct"`
	// The identification result of whether there is a beard.
	Mustache *ImageModerationResponseBodyDataExtFaceDataMustache `json:"Mustache,omitempty" xml:"Mustache,omitempty" type:"Struct"`
	// The quality information of the face image.
	Quality *ImageModerationResponseBodyDataExtFaceDataQuality `json:"Quality,omitempty" xml:"Quality,omitempty" type:"Struct"`
	// The smiling degree of the face.
	//
	// example:
	//
	// 85.88
	Smile *float32 `json:"Smile,omitempty" xml:"Smile,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceData) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceData) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetAge(v int32) *ImageModerationResponseBodyDataExtFaceData {
	s.Age = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetBang(v *ImageModerationResponseBodyDataExtFaceDataBang) *ImageModerationResponseBodyDataExtFaceData {
	s.Bang = v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetGender(v *ImageModerationResponseBodyDataExtFaceDataGender) *ImageModerationResponseBodyDataExtFaceData {
	s.Gender = v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetGlasses(v string) *ImageModerationResponseBodyDataExtFaceData {
	s.Glasses = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetHairstyle(v *ImageModerationResponseBodyDataExtFaceDataHairstyle) *ImageModerationResponseBodyDataExtFaceData {
	s.Hairstyle = v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetHat(v *ImageModerationResponseBodyDataExtFaceDataHat) *ImageModerationResponseBodyDataExtFaceData {
	s.Hat = v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetLocation(v *ImageModerationResponseBodyDataExtFaceDataLocation) *ImageModerationResponseBodyDataExtFaceData {
	s.Location = v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetMask(v *ImageModerationResponseBodyDataExtFaceDataMask) *ImageModerationResponseBodyDataExtFaceData {
	s.Mask = v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetMustache(v *ImageModerationResponseBodyDataExtFaceDataMustache) *ImageModerationResponseBodyDataExtFaceData {
	s.Mustache = v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetQuality(v *ImageModerationResponseBodyDataExtFaceDataQuality) *ImageModerationResponseBodyDataExtFaceData {
	s.Quality = v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceData) SetSmile(v float32) *ImageModerationResponseBodyDataExtFaceData {
	s.Smile = &v
	return s
}

type ImageModerationResponseBodyDataExtFaceDataBang struct {
	// The confidence level of the bang recognition result. Valid values: 0 to 100. A higher value indicates a more credible result.
	//
	// example:
	//
	// 81.88
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// Indicates whether the recognition result of bangs is available.
	//
	// example:
	//
	// Has
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceDataBang) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceDataBang) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceDataBang) SetConfidence(v float32) *ImageModerationResponseBodyDataExtFaceDataBang {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataBang) SetValue(v string) *ImageModerationResponseBodyDataExtFaceDataBang {
	s.Value = &v
	return s
}

type ImageModerationResponseBodyDataExtFaceDataGender struct {
	// The confidence level of the gender recognition result. Valid values: 0 to 100. A higher value indicates a more credible result.
	//
	// example:
	//
	// 81.88
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The gender recognition result. Valid values:
	//
	// - Male
	//
	// - FeMale
	//
	// example:
	//
	// FeMale
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceDataGender) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceDataGender) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceDataGender) SetConfidence(v float32) *ImageModerationResponseBodyDataExtFaceDataGender {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataGender) SetValue(v string) *ImageModerationResponseBodyDataExtFaceDataGender {
	s.Value = &v
	return s
}

type ImageModerationResponseBodyDataExtFaceDataHairstyle struct {
	// The confidence level of the hairstyle recognition result. Valid values: 0 to 100. A higher value indicates a more credible result.
	//
	// example:
	//
	// 81.88
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The hairstyle recognition result. Valid values:
	//
	// - Bald: bald head.
	//
	// - Long: Long hair.
	//
	// - Short: Short hair.
	//
	// example:
	//
	// Short
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceDataHairstyle) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceDataHairstyle) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceDataHairstyle) SetConfidence(v float32) *ImageModerationResponseBodyDataExtFaceDataHairstyle {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataHairstyle) SetValue(v string) *ImageModerationResponseBodyDataExtFaceDataHairstyle {
	s.Value = &v
	return s
}

type ImageModerationResponseBodyDataExtFaceDataHat struct {
	// The confidence level of the result of wearing the hat. Valid values: 0 to 100. A higher value indicates a more credible result.
	//
	// example:
	//
	// 88.88
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The recognition result of whether to wear the hat. Valid values:
	//
	// - Wear: Wear a hat.
	//
	// - None: No hat.
	//
	// example:
	//
	// Wear
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceDataHat) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceDataHat) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceDataHat) SetConfidence(v float32) *ImageModerationResponseBodyDataExtFaceDataHat {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataHat) SetValue(v string) *ImageModerationResponseBodyDataExtFaceDataHat {
	s.Value = &v
	return s
}

type ImageModerationResponseBodyDataExtFaceDataLocation struct {
	// The height of the face area. Unit: pixels.
	//
	// example:
	//
	// 26
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width of the face area. Unit: pixels.
	//
	// example:
	//
	// 83
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The distance from the upper-left corner of the face area to the y-axis with the upper-left corner of the image as the coordinate origin. Unit: pixels.
	//
	// example:
	//
	// 41
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The distance from the upper-left corner of the face area to the x-axis with the upper-left corner of the image as the coordinate origin. Unit: pixels.
	//
	// example:
	//
	// 84
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceDataLocation) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceDataLocation) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceDataLocation) SetH(v int32) *ImageModerationResponseBodyDataExtFaceDataLocation {
	s.H = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataLocation) SetW(v int32) *ImageModerationResponseBodyDataExtFaceDataLocation {
	s.W = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataLocation) SetX(v int32) *ImageModerationResponseBodyDataExtFaceDataLocation {
	s.X = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataLocation) SetY(v int32) *ImageModerationResponseBodyDataExtFaceDataLocation {
	s.Y = &v
	return s
}

type ImageModerationResponseBodyDataExtFaceDataMask struct {
	// The confidence level of the result of wearing the mask. Valid values: 0 to 100. A higher value indicates a more credible result.
	//
	// example:
	//
	// 99.99
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The recognition result of whether to wear a mask. Valid values:
	//
	// - Wear a mask.
	//
	//  - None: No mask.
	//
	// example:
	//
	// Wear
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceDataMask) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceDataMask) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceDataMask) SetConfidence(v float32) *ImageModerationResponseBodyDataExtFaceDataMask {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataMask) SetValue(v string) *ImageModerationResponseBodyDataExtFaceDataMask {
	s.Value = &v
	return s
}

type ImageModerationResponseBodyDataExtFaceDataMustache struct {
	// The confidence level of the result of the beard. Valid values: 0 to 100. A higher value indicates a more credible result.
	//
	// example:
	//
	// 99.99
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The identification result of whether there is a beard.Valid values:
	//
	// - Has:have a beard.
	//
	// - None:No beard.
	//
	// example:
	//
	// Has
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceDataMustache) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceDataMustache) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceDataMustache) SetConfidence(v float32) *ImageModerationResponseBodyDataExtFaceDataMustache {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataMustache) SetValue(v string) *ImageModerationResponseBodyDataExtFaceDataMustache {
	s.Value = &v
	return s
}

type ImageModerationResponseBodyDataExtFaceDataQuality struct {
	// The blur of the face image. Valid values: 0 to 100. The higher the score, the more fuzzy it is.
	//
	// Recommended values: 0 to 25.
	//
	// example:
	//
	// 5.88
	Blur *float32 `json:"Blur,omitempty" xml:"Blur,omitempty"`
	// The integrity of the human face. Recommended values:80 to 100.
	//
	// example:
	//
	// 100.0
	Integrity *float32 `json:"Integrity,omitempty" xml:"Integrity,omitempty"`
	// The head-up or head-down angle of the face.
	//
	// Recommended values:-30 to 30.
	//
	// example:
	//
	// 5.88
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	// The plane rotation angle of the face.
	//
	// Recommended values:-30 to 30.
	//
	// example:
	//
	// 5.18
	Roll *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	// The left and right shaking angle of the human face.
	//
	// Recommended values:-30 to 30.
	//
	// example:
	//
	// 5.18
	Yaw *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s ImageModerationResponseBodyDataExtFaceDataQuality) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtFaceDataQuality) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtFaceDataQuality) SetBlur(v float32) *ImageModerationResponseBodyDataExtFaceDataQuality {
	s.Blur = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataQuality) SetIntegrity(v float32) *ImageModerationResponseBodyDataExtFaceDataQuality {
	s.Integrity = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataQuality) SetPitch(v float32) *ImageModerationResponseBodyDataExtFaceDataQuality {
	s.Pitch = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataQuality) SetRoll(v float32) *ImageModerationResponseBodyDataExtFaceDataQuality {
	s.Roll = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtFaceDataQuality) SetYaw(v float32) *ImageModerationResponseBodyDataExtFaceDataQuality {
	s.Yaw = &v
	return s
}

type ImageModerationResponseBodyDataExtLogoData struct {
	// Location information.
	Location *ImageModerationResponseBodyDataExtLogoDataLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// Logo information.
	Logo []*ImageModerationResponseBodyDataExtLogoDataLogo `json:"Logo,omitempty" xml:"Logo,omitempty" type:"Repeated"`
}

func (s ImageModerationResponseBodyDataExtLogoData) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtLogoData) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtLogoData) SetLocation(v *ImageModerationResponseBodyDataExtLogoDataLocation) *ImageModerationResponseBodyDataExtLogoData {
	s.Location = v
	return s
}

func (s *ImageModerationResponseBodyDataExtLogoData) SetLogo(v []*ImageModerationResponseBodyDataExtLogoDataLogo) *ImageModerationResponseBodyDataExtLogoData {
	s.Logo = v
	return s
}

type ImageModerationResponseBodyDataExtLogoDataLocation struct {
	// The height of the text area, in pixels.
	//
	// example:
	//
	// 44
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width of the text area, in pixels.
	//
	// example:
	//
	// 33
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The distance between the upper-left corner of the text area and the y-axis, using the upper-left corner of the image as the coordinate origin, in pixels.
	//
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The distance between the upper left corner of the text area and the x-axis, with the upper left corner of the image as the coordinate origin, in pixels.
	//
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageModerationResponseBodyDataExtLogoDataLocation) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtLogoDataLocation) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtLogoDataLocation) SetH(v int32) *ImageModerationResponseBodyDataExtLogoDataLocation {
	s.H = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtLogoDataLocation) SetW(v int32) *ImageModerationResponseBodyDataExtLogoDataLocation {
	s.W = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtLogoDataLocation) SetX(v int32) *ImageModerationResponseBodyDataExtLogoDataLocation {
	s.X = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtLogoDataLocation) SetY(v int32) *ImageModerationResponseBodyDataExtLogoDataLocation {
	s.Y = &v
	return s
}

type ImageModerationResponseBodyDataExtLogoDataLogo struct {
	// The score of the confidence level. Valid values: 0 to 100. The value is accurate to two decimal places. Some labels do not have scores of confidence levels.
	//
	// example:
	//
	// 99.01
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// Logo category.
	//
	// example:
	//
	// xx
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Logo name.
	//
	// example:
	//
	// xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ImageModerationResponseBodyDataExtLogoDataLogo) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtLogoDataLogo) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtLogoDataLogo) SetConfidence(v float32) *ImageModerationResponseBodyDataExtLogoDataLogo {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtLogoDataLogo) SetLabel(v string) *ImageModerationResponseBodyDataExtLogoDataLogo {
	s.Label = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtLogoDataLogo) SetName(v string) *ImageModerationResponseBodyDataExtLogoDataLogo {
	s.Name = &v
	return s
}

type ImageModerationResponseBodyDataExtOcrResult struct {
	// Location information.
	Location *ImageModerationResponseBodyDataExtOcrResultLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// The text information in the recognized image.
	//
	// example:
	//
	// xx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ImageModerationResponseBodyDataExtOcrResult) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtOcrResult) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtOcrResult) SetLocation(v *ImageModerationResponseBodyDataExtOcrResultLocation) *ImageModerationResponseBodyDataExtOcrResult {
	s.Location = v
	return s
}

func (s *ImageModerationResponseBodyDataExtOcrResult) SetText(v string) *ImageModerationResponseBodyDataExtOcrResult {
	s.Text = &v
	return s
}

type ImageModerationResponseBodyDataExtOcrResultLocation struct {
	// The height of the text area, in pixels.
	//
	// example:
	//
	// 44
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width of the text area, in pixels.
	//
	// example:
	//
	// 33
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The distance between the upper-left corner of the text area and the y-axis, using the upper-left corner of the image as the coordinate origin, in pixels.
	//
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The distance between the upper left corner of the text area and the x-axis, with the upper left corner of the image as the coordinate origin, in pixels.
	//
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageModerationResponseBodyDataExtOcrResultLocation) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtOcrResultLocation) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtOcrResultLocation) SetH(v int32) *ImageModerationResponseBodyDataExtOcrResultLocation {
	s.H = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtOcrResultLocation) SetW(v int32) *ImageModerationResponseBodyDataExtOcrResultLocation {
	s.W = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtOcrResultLocation) SetX(v int32) *ImageModerationResponseBodyDataExtOcrResultLocation {
	s.X = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtOcrResultLocation) SetY(v int32) *ImageModerationResponseBodyDataExtOcrResultLocation {
	s.Y = &v
	return s
}

type ImageModerationResponseBodyDataExtPublicFigure struct {
	// Identified person coding information.
	//
	// example:
	//
	// xx
	FigureId *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
	// Identified person name information.
	//
	// example:
	//
	// xx
	FigureName *string `json:"FigureName,omitempty" xml:"FigureName,omitempty"`
	// the data array of location info
	Location []*ImageModerationResponseBodyDataExtPublicFigureLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Repeated"`
}

func (s ImageModerationResponseBodyDataExtPublicFigure) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtPublicFigure) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtPublicFigure) SetFigureId(v string) *ImageModerationResponseBodyDataExtPublicFigure {
	s.FigureId = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtPublicFigure) SetFigureName(v string) *ImageModerationResponseBodyDataExtPublicFigure {
	s.FigureName = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtPublicFigure) SetLocation(v []*ImageModerationResponseBodyDataExtPublicFigureLocation) *ImageModerationResponseBodyDataExtPublicFigure {
	s.Location = v
	return s
}

type ImageModerationResponseBodyDataExtPublicFigureLocation struct {
	// The height
	//
	// example:
	//
	// 44
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The weight
	//
	// example:
	//
	// 33
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// X coordinate
	//
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// Y coordinate
	//
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageModerationResponseBodyDataExtPublicFigureLocation) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtPublicFigureLocation) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtPublicFigureLocation) SetH(v int32) *ImageModerationResponseBodyDataExtPublicFigureLocation {
	s.H = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtPublicFigureLocation) SetW(v int32) *ImageModerationResponseBodyDataExtPublicFigureLocation {
	s.W = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtPublicFigureLocation) SetX(v int32) *ImageModerationResponseBodyDataExtPublicFigureLocation {
	s.X = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtPublicFigureLocation) SetY(v int32) *ImageModerationResponseBodyDataExtPublicFigureLocation {
	s.Y = &v
	return s
}

type ImageModerationResponseBodyDataExtRecognition struct {
	// The category of image recognition.
	//
	// example:
	//
	// xx
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	// The score of the confidence level. Valid values: 0 to 100. The value is accurate to two decimal places. Some labels do not have scores of confidence levels.
	//
	// example:
	//
	// 99.01
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
}

func (s ImageModerationResponseBodyDataExtRecognition) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtRecognition) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtRecognition) SetClassification(v string) *ImageModerationResponseBodyDataExtRecognition {
	s.Classification = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtRecognition) SetConfidence(v float32) *ImageModerationResponseBodyDataExtRecognition {
	s.Confidence = &v
	return s
}

type ImageModerationResponseBodyDataExtTextInImage struct {
	// When a custom text library is hit, the custom library ID, custom library name, and custom word are returned.
	CustomText []*ImageModerationResponseBodyDataExtTextInImageCustomText `json:"CustomText,omitempty" xml:"CustomText,omitempty" type:"Repeated"`
	// Returns the text information in the recognized image.
	OcrResult []*ImageModerationResponseBodyDataExtTextInImageOcrResult `json:"OcrResult,omitempty" xml:"OcrResult,omitempty" type:"Repeated"`
	// The risk words that are hit. Multiple words are separated by commas (,).
	RiskWord []*string `json:"RiskWord,omitempty" xml:"RiskWord,omitempty" type:"Repeated"`
}

func (s ImageModerationResponseBodyDataExtTextInImage) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtTextInImage) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtTextInImage) SetCustomText(v []*ImageModerationResponseBodyDataExtTextInImageCustomText) *ImageModerationResponseBodyDataExtTextInImage {
	s.CustomText = v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImage) SetOcrResult(v []*ImageModerationResponseBodyDataExtTextInImageOcrResult) *ImageModerationResponseBodyDataExtTextInImage {
	s.OcrResult = v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImage) SetRiskWord(v []*string) *ImageModerationResponseBodyDataExtTextInImage {
	s.RiskWord = v
	return s
}

type ImageModerationResponseBodyDataExtTextInImageCustomText struct {
	// Custom words, multiple words separated by commas.
	//
	// example:
	//
	// aaa,bbb
	KeyWords *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// Custom library ID.
	//
	// example:
	//
	// 123456
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// Custom library name.
	//
	// example:
	//
	// xxxx
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s ImageModerationResponseBodyDataExtTextInImageCustomText) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtTextInImageCustomText) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtTextInImageCustomText) SetKeyWords(v string) *ImageModerationResponseBodyDataExtTextInImageCustomText {
	s.KeyWords = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImageCustomText) SetLibId(v string) *ImageModerationResponseBodyDataExtTextInImageCustomText {
	s.LibId = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImageCustomText) SetLibName(v string) *ImageModerationResponseBodyDataExtTextInImageCustomText {
	s.LibName = &v
	return s
}

type ImageModerationResponseBodyDataExtTextInImageOcrResult struct {
	// Location information.
	Location *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// The text information in the recognized image.
	//
	// example:
	//
	// xx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ImageModerationResponseBodyDataExtTextInImageOcrResult) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtTextInImageOcrResult) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResult) SetLocation(v *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) *ImageModerationResponseBodyDataExtTextInImageOcrResult {
	s.Location = v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResult) SetText(v string) *ImageModerationResponseBodyDataExtTextInImageOcrResult {
	s.Text = &v
	return s
}

type ImageModerationResponseBodyDataExtTextInImageOcrResultLocation struct {
	// The height of the text area, in pixels.
	//
	// example:
	//
	// 33
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width of the text area, in pixels.
	//
	// example:
	//
	// 44
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The distance between the upper-left corner of the text area and the y-axis, using the upper-left corner of the image as the coordinate origin, in pixels.
	//
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The distance between the upper left corner of the text area and the x-axis, with the upper left corner of the image as the coordinate origin, in pixels.
	//
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) SetH(v int32) *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation {
	s.H = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) SetW(v int32) *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation {
	s.W = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) SetX(v int32) *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation {
	s.X = &v
	return s
}

func (s *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation) SetY(v int32) *ImageModerationResponseBodyDataExtTextInImageOcrResultLocation {
	s.Y = &v
	return s
}

type ImageModerationResponseBodyDataExtVlContent struct {
	// the vl output content
	//
	// example:
	//
	// this is a cat
	OutputText *string `json:"OutputText,omitempty" xml:"OutputText,omitempty"`
}

func (s ImageModerationResponseBodyDataExtVlContent) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataExtVlContent) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataExtVlContent) SetOutputText(v string) *ImageModerationResponseBodyDataExtVlContent {
	s.OutputText = &v
	return s
}

type ImageModerationResponseBodyDataResult struct {
	// The score of the confidence level. Valid values: 0 to 100. The value is accurate to two decimal places. Some labels do not have scores of confidence levels.
	//
	// example:
	//
	// 81.22
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The description of the result.
	//
	// example:
	//
	// no risk
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The labels returned after the image moderation. Multiple risk labels and the corresponding scores of confidence levels may be returned for an image.
	//
	// example:
	//
	// violent_explosion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Risk Level
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ImageModerationResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataResult) SetConfidence(v float32) *ImageModerationResponseBodyDataResult {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataResult) SetDescription(v string) *ImageModerationResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *ImageModerationResponseBodyDataResult) SetLabel(v string) *ImageModerationResponseBodyDataResult {
	s.Label = &v
	return s
}

func (s *ImageModerationResponseBodyDataResult) SetRiskLevel(v string) *ImageModerationResponseBodyDataResult {
	s.RiskLevel = &v
	return s
}

type ImageModerationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageModerationResponse) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponse) GoString() string {
	return s.String()
}

func (s *ImageModerationResponse) SetHeaders(v map[string]*string) *ImageModerationResponse {
	s.Headers = v
	return s
}

func (s *ImageModerationResponse) SetStatusCode(v int32) *ImageModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageModerationResponse) SetBody(v *ImageModerationResponseBody) *ImageModerationResponse {
	s.Body = v
	return s
}

type ManualCallbackRequest struct {
	// Channel field
	//
	// example:
	//
	// ant
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// Checksum.
	//
	// example:
	//
	// abc
	Checksum *string `json:"Checksum,omitempty" xml:"Checksum,omitempty"`
	// Code value
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	//
	// example:
	//
	// {\\"Result\\": [{\\"Confidence\\": 100.0, \\"CustomizedHit\\": [{\\"KeyWords\\": u\\"\\u4fdd\\u969c,\\u6700\\u5927,\\u9ad8\\u7ea7\\", \\"LibName\\": u\\"\\u4f18\\u8def\\u654f\\u611f\\u8bcd\\"}], \\"Label\\": \\"customized\\"}]}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Message information
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Platform request ID, used for troubleshooting assistance
	//
	// example:
	//
	// B0963D30-BAB4-562F-9ED0-7A23AEC51C7C
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
}

func (s ManualCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s ManualCallbackRequest) GoString() string {
	return s.String()
}

func (s *ManualCallbackRequest) SetChannel(v string) *ManualCallbackRequest {
	s.Channel = &v
	return s
}

func (s *ManualCallbackRequest) SetChecksum(v string) *ManualCallbackRequest {
	s.Checksum = &v
	return s
}

func (s *ManualCallbackRequest) SetCode(v string) *ManualCallbackRequest {
	s.Code = &v
	return s
}

func (s *ManualCallbackRequest) SetData(v string) *ManualCallbackRequest {
	s.Data = &v
	return s
}

func (s *ManualCallbackRequest) SetMsg(v string) *ManualCallbackRequest {
	s.Msg = &v
	return s
}

func (s *ManualCallbackRequest) SetReqId(v string) *ManualCallbackRequest {
	s.ReqId = &v
	return s
}

type ManualCallbackResponseBody struct {
	// Error code
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message information
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ManualCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ManualCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *ManualCallbackResponseBody) SetCode(v int32) *ManualCallbackResponseBody {
	s.Code = &v
	return s
}

func (s *ManualCallbackResponseBody) SetMessage(v string) *ManualCallbackResponseBody {
	s.Message = &v
	return s
}

func (s *ManualCallbackResponseBody) SetRequestId(v string) *ManualCallbackResponseBody {
	s.RequestId = &v
	return s
}

type ManualCallbackResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManualCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManualCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s ManualCallbackResponse) GoString() string {
	return s.String()
}

func (s *ManualCallbackResponse) SetHeaders(v map[string]*string) *ManualCallbackResponse {
	s.Headers = v
	return s
}

func (s *ManualCallbackResponse) SetStatusCode(v int32) *ManualCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *ManualCallbackResponse) SetBody(v *ManualCallbackResponseBody) *ManualCallbackResponse {
	s.Body = v
	return s
}

type ManualModerationRequest struct {
	// Service.
	//
	// example:
	//
	// imageManualCheck
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// Parameter set required for the review service, in JSON string format.
	//
	// - url: The URL of the object to be checked. Please ensure that this URL is publicly accessible.
	//
	// - dataId: Optional, the data ID corresponding to the object being checked.
	//
	// example:
	//
	// {"url": "https://talesofai.oss-cn-shanghai.aliyuncs.com/xxx.mp4", "dataId": "data1234"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ManualModerationRequest) String() string {
	return tea.Prettify(s)
}

func (s ManualModerationRequest) GoString() string {
	return s.String()
}

func (s *ManualModerationRequest) SetService(v string) *ManualModerationRequest {
	s.Service = &v
	return s
}

func (s *ManualModerationRequest) SetServiceParameters(v string) *ManualModerationRequest {
	s.ServiceParameters = &v
	return s
}

type ManualModerationResponseBody struct {
	// Status code
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	Data *ManualModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error message
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ManualModerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ManualModerationResponseBody) GoString() string {
	return s.String()
}

func (s *ManualModerationResponseBody) SetCode(v int32) *ManualModerationResponseBody {
	s.Code = &v
	return s
}

func (s *ManualModerationResponseBody) SetData(v *ManualModerationResponseBodyData) *ManualModerationResponseBody {
	s.Data = v
	return s
}

func (s *ManualModerationResponseBody) SetMessage(v string) *ManualModerationResponseBody {
	s.Message = &v
	return s
}

func (s *ManualModerationResponseBody) SetRequestId(v string) *ManualModerationResponseBody {
	s.RequestId = &v
	return s
}

type ManualModerationResponseBodyData struct {
	// The value of dataId passed during the API request. This field will not be present if it was not provided during the request.
	//
	// example:
	//
	// 2a5389eb-4ff8-4584-ac99-644e2a539aa1
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// Task ID
	//
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ManualModerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ManualModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ManualModerationResponseBodyData) SetDataId(v string) *ManualModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *ManualModerationResponseBodyData) SetTaskId(v string) *ManualModerationResponseBodyData {
	s.TaskId = &v
	return s
}

type ManualModerationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManualModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManualModerationResponse) String() string {
	return tea.Prettify(s)
}

func (s ManualModerationResponse) GoString() string {
	return s.String()
}

func (s *ManualModerationResponse) SetHeaders(v map[string]*string) *ManualModerationResponse {
	s.Headers = v
	return s
}

func (s *ManualModerationResponse) SetStatusCode(v int32) *ManualModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *ManualModerationResponse) SetBody(v *ManualModerationResponseBody) *ManualModerationResponse {
	s.Body = v
	return s
}

type ManualModerationResultRequest struct {
	// Set of parameters required by the service, in JSON string format.
	//
	// - TaskId: The task ID returned when the task was submitted.
	//
	// example:
	//
	// {\\"TaskId\\":\\"e5f2d886-4c23-440d-999c-bd98acde11b6\\"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ManualModerationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ManualModerationResultRequest) GoString() string {
	return s.String()
}

func (s *ManualModerationResultRequest) SetServiceParameters(v string) *ManualModerationResultRequest {
	s.ServiceParameters = &v
	return s
}

type ManualModerationResultResponseBody struct {
	// Error code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	Data *ManualModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error message
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ManualModerationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ManualModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *ManualModerationResultResponseBody) SetCode(v int32) *ManualModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *ManualModerationResultResponseBody) SetData(v *ManualModerationResultResponseBodyData) *ManualModerationResultResponseBody {
	s.Data = v
	return s
}

func (s *ManualModerationResultResponseBody) SetMessage(v string) *ManualModerationResultResponseBody {
	s.Message = &v
	return s
}

func (s *ManualModerationResultResponseBody) SetRequestId(v string) *ManualModerationResultResponseBody {
	s.RequestId = &v
	return s
}

type ManualModerationResultResponseBodyData struct {
	// The value of dataId passed during the API request. This field will not be present if it was not provided during the request.
	//
	// example:
	//
	// data1234
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// Detailed label results.
	Result []*ManualModerationResultResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Risk level, returned based on the set high and low risk scores. Possible values include:
	//
	// - high: High risk
	//
	//
	//
	// - low: Low risk
	//
	//  - none: No risk detected
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// Task ID
	//
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ManualModerationResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ManualModerationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *ManualModerationResultResponseBodyData) SetDataId(v string) *ManualModerationResultResponseBodyData {
	s.DataId = &v
	return s
}

func (s *ManualModerationResultResponseBodyData) SetResult(v []*ManualModerationResultResponseBodyDataResult) *ManualModerationResultResponseBodyData {
	s.Result = v
	return s
}

func (s *ManualModerationResultResponseBodyData) SetRiskLevel(v string) *ManualModerationResultResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *ManualModerationResultResponseBodyData) SetTaskId(v string) *ManualModerationResultResponseBodyData {
	s.TaskId = &v
	return s
}

type ManualModerationResultResponseBodyDataResult struct {
	// Label description
	//
	// example:
	//
	// no risk
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Risk label
	//
	// example:
	//
	// violent_explosion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ManualModerationResultResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ManualModerationResultResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ManualModerationResultResponseBodyDataResult) SetDescription(v string) *ManualModerationResultResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *ManualModerationResultResponseBodyDataResult) SetLabel(v string) *ManualModerationResultResponseBodyDataResult {
	s.Label = &v
	return s
}

type ManualModerationResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManualModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManualModerationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ManualModerationResultResponse) GoString() string {
	return s.String()
}

func (s *ManualModerationResultResponse) SetHeaders(v map[string]*string) *ManualModerationResultResponse {
	s.Headers = v
	return s
}

func (s *ManualModerationResultResponse) SetStatusCode(v int32) *ManualModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ManualModerationResultResponse) SetBody(v *ManualModerationResultResponseBody) *ManualModerationResultResponse {
	s.Body = v
	return s
}

type MultiModalGuardRequest struct {
	// example:
	//
	// query_security_check
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s MultiModalGuardRequest) String() string {
	return tea.Prettify(s)
}

func (s MultiModalGuardRequest) GoString() string {
	return s.String()
}

func (s *MultiModalGuardRequest) SetService(v string) *MultiModalGuardRequest {
	s.Service = &v
	return s
}

func (s *MultiModalGuardRequest) SetServiceParameters(v string) *MultiModalGuardRequest {
	s.ServiceParameters = &v
	return s
}

type MultiModalGuardResponseBody struct {
	// example:
	//
	// 200
	Code *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *MultiModalGuardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MultiModalGuardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MultiModalGuardResponseBody) GoString() string {
	return s.String()
}

func (s *MultiModalGuardResponseBody) SetCode(v int32) *MultiModalGuardResponseBody {
	s.Code = &v
	return s
}

func (s *MultiModalGuardResponseBody) SetData(v *MultiModalGuardResponseBodyData) *MultiModalGuardResponseBody {
	s.Data = v
	return s
}

func (s *MultiModalGuardResponseBody) SetMessage(v string) *MultiModalGuardResponseBody {
	s.Message = &v
	return s
}

func (s *MultiModalGuardResponseBody) SetRequestId(v string) *MultiModalGuardResponseBody {
	s.RequestId = &v
	return s
}

type MultiModalGuardResponseBodyData struct {
	// example:
	//
	// data1234
	DataId *string                                  `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Detail []*MultiModalGuardResponseBodyDataDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s MultiModalGuardResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MultiModalGuardResponseBodyData) GoString() string {
	return s.String()
}

func (s *MultiModalGuardResponseBodyData) SetDataId(v string) *MultiModalGuardResponseBodyData {
	s.DataId = &v
	return s
}

func (s *MultiModalGuardResponseBodyData) SetDetail(v []*MultiModalGuardResponseBodyDataDetail) *MultiModalGuardResponseBodyData {
	s.Detail = v
	return s
}

func (s *MultiModalGuardResponseBodyData) SetSuggestion(v string) *MultiModalGuardResponseBodyData {
	s.Suggestion = &v
	return s
}

type MultiModalGuardResponseBodyDataDetail struct {
	// example:
	//
	// none
	Level  *string                                        `json:"Level,omitempty" xml:"Level,omitempty"`
	Result []*MultiModalGuardResponseBodyDataDetailResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// example:
	//
	// contentModeration
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MultiModalGuardResponseBodyDataDetail) String() string {
	return tea.Prettify(s)
}

func (s MultiModalGuardResponseBodyDataDetail) GoString() string {
	return s.String()
}

func (s *MultiModalGuardResponseBodyDataDetail) SetLevel(v string) *MultiModalGuardResponseBodyDataDetail {
	s.Level = &v
	return s
}

func (s *MultiModalGuardResponseBodyDataDetail) SetResult(v []*MultiModalGuardResponseBodyDataDetailResult) *MultiModalGuardResponseBodyDataDetail {
	s.Result = v
	return s
}

func (s *MultiModalGuardResponseBodyDataDetail) SetSuggestion(v string) *MultiModalGuardResponseBodyDataDetail {
	s.Suggestion = &v
	return s
}

func (s *MultiModalGuardResponseBodyDataDetail) SetType(v string) *MultiModalGuardResponseBodyDataDetail {
	s.Type = &v
	return s
}

type MultiModalGuardResponseBodyDataDetailResult struct {
	// example:
	//
	// 0
	Confidence  *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {}
	Ext interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// contraband_act
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// none
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s MultiModalGuardResponseBodyDataDetailResult) String() string {
	return tea.Prettify(s)
}

func (s MultiModalGuardResponseBodyDataDetailResult) GoString() string {
	return s.String()
}

func (s *MultiModalGuardResponseBodyDataDetailResult) SetConfidence(v float32) *MultiModalGuardResponseBodyDataDetailResult {
	s.Confidence = &v
	return s
}

func (s *MultiModalGuardResponseBodyDataDetailResult) SetDescription(v string) *MultiModalGuardResponseBodyDataDetailResult {
	s.Description = &v
	return s
}

func (s *MultiModalGuardResponseBodyDataDetailResult) SetExt(v interface{}) *MultiModalGuardResponseBodyDataDetailResult {
	s.Ext = v
	return s
}

func (s *MultiModalGuardResponseBodyDataDetailResult) SetLabel(v string) *MultiModalGuardResponseBodyDataDetailResult {
	s.Label = &v
	return s
}

func (s *MultiModalGuardResponseBodyDataDetailResult) SetLevel(v string) *MultiModalGuardResponseBodyDataDetailResult {
	s.Level = &v
	return s
}

type MultiModalGuardResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MultiModalGuardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MultiModalGuardResponse) String() string {
	return tea.Prettify(s)
}

func (s MultiModalGuardResponse) GoString() string {
	return s.String()
}

func (s *MultiModalGuardResponse) SetHeaders(v map[string]*string) *MultiModalGuardResponse {
	s.Headers = v
	return s
}

func (s *MultiModalGuardResponse) SetStatusCode(v int32) *MultiModalGuardResponse {
	s.StatusCode = &v
	return s
}

func (s *MultiModalGuardResponse) SetBody(v *MultiModalGuardResponseBody) *MultiModalGuardResponse {
	s.Body = v
	return s
}

type TextModerationRequest struct {
	// The type of the moderation service. Valid values: nickname_detection: user nickname chat_detection: chat interactions comment_detection: dynamic comments pgc_detection: professionally-generated content (PGC) teaching materials
	//
	// Valid values:
	//
	// 	- pgc_detection: moderation of PGC teaching materials
	//
	// 	- nickname_detection: user nickname moderation
	//
	// 	- comment_multilingual_pro: multi-language moderation in international business scenarios
	//
	// 	- chat_detection: moderation of interactive content of private chats
	//
	// 	- ad_compliance_detection: advertising law compliance identification
	//
	// 	- comment_detection: moderation of comment content of public chats
	//
	// 	- ai_art_detection: AI-generated text identfication
	//
	// example:
	//
	// nickname_detection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// example:
	//
	// {"content":"Content to be moderated"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s TextModerationRequest) String() string {
	return tea.Prettify(s)
}

func (s TextModerationRequest) GoString() string {
	return s.String()
}

func (s *TextModerationRequest) SetService(v string) *TextModerationRequest {
	s.Service = &v
	return s
}

func (s *TextModerationRequest) SetServiceParameters(v string) *TextModerationRequest {
	s.ServiceParameters = &v
	return s
}

type TextModerationResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The moderation results.
	Data *TextModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TextModerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TextModerationResponseBody) GoString() string {
	return s.String()
}

func (s *TextModerationResponseBody) SetCode(v int32) *TextModerationResponseBody {
	s.Code = &v
	return s
}

func (s *TextModerationResponseBody) SetData(v *TextModerationResponseBodyData) *TextModerationResponseBody {
	s.Data = v
	return s
}

func (s *TextModerationResponseBody) SetMessage(v string) *TextModerationResponseBody {
	s.Message = &v
	return s
}

func (s *TextModerationResponseBody) SetRequestId(v string) *TextModerationResponseBody {
	s.RequestId = &v
	return s
}

type TextModerationResponseBodyData struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 123456
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// The ID of the moderated object.
	//
	// example:
	//
	// text1234
	DataId *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// The description of the labels.
	//
	// example:
	//
	// no risk
	Descriptions *string `json:"descriptions,omitempty" xml:"descriptions,omitempty"`
	// The device ID.
	//
	// example:
	//
	// xxxxxx
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// The labels. Multiple labels are separated by commas (,). Valid values: ad: ad violation profanity: abuse contraband: contraband sexual_content: pornography violence: violence nonsense: irrigation spam: spam negative_content: undesirable content cyberbullying: cyberbullying C_customized: custom library that is hit
	//
	// example:
	//
	// porn
	Labels       *string `json:"labels,omitempty" xml:"labels,omitempty"`
	ManualTaskId *string `json:"manualTaskId,omitempty" xml:"manualTaskId,omitempty"`
	// The JSON string used to locate the cause. Valid values: riskTips: subcategory label riskWords: risk words adNums: hit advertising number customizedWords: customized words customizedLibs: customized libraries
	//
	// example:
	//
	// {\\"detectedLanguage\\":\\"ar\\",\\"riskTips\\":\\"sexuality_Suggestive\\",\\"riskWords\\":\\"pxxxxy\\",\\"translatedContent\\":\\"pxxxxy sxxxx\\"}
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s TextModerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TextModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *TextModerationResponseBodyData) SetAccountId(v string) *TextModerationResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *TextModerationResponseBodyData) SetDataId(v string) *TextModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *TextModerationResponseBodyData) SetDescriptions(v string) *TextModerationResponseBodyData {
	s.Descriptions = &v
	return s
}

func (s *TextModerationResponseBodyData) SetDeviceId(v string) *TextModerationResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *TextModerationResponseBodyData) SetLabels(v string) *TextModerationResponseBodyData {
	s.Labels = &v
	return s
}

func (s *TextModerationResponseBodyData) SetManualTaskId(v string) *TextModerationResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *TextModerationResponseBodyData) SetReason(v string) *TextModerationResponseBodyData {
	s.Reason = &v
	return s
}

type TextModerationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TextModerationResponse) String() string {
	return tea.Prettify(s)
}

func (s TextModerationResponse) GoString() string {
	return s.String()
}

func (s *TextModerationResponse) SetHeaders(v map[string]*string) *TextModerationResponse {
	s.Headers = v
	return s
}

func (s *TextModerationResponse) SetStatusCode(v int32) *TextModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *TextModerationResponse) SetBody(v *TextModerationResponseBody) *TextModerationResponse {
	s.Body = v
	return s
}

type TextModerationPlusRequest struct {
	// The moderation service.
	//
	// Valid values:
	//
	// 	- chat_detection_pro: moderation of interactive content of private chats_Professional
	//
	// 	- llm_response_moderation: moderation of text generated by LLMs
	//
	// 	- llm_query_moderation: moderation of input text of LLMs
	//
	// 	- nickname_detection_pro: moderation of user nicknames_Professional
	//
	// 	- comment_detection_pro: moderation of comment content of public chats_Professional
	//
	// example:
	//
	// llm_query_moderation
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// example:
	//
	// {"content":"Content to be moderated"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s TextModerationPlusRequest) String() string {
	return tea.Prettify(s)
}

func (s TextModerationPlusRequest) GoString() string {
	return s.String()
}

func (s *TextModerationPlusRequest) SetService(v string) *TextModerationPlusRequest {
	s.Service = &v
	return s
}

func (s *TextModerationPlusRequest) SetServiceParameters(v string) *TextModerationPlusRequest {
	s.ServiceParameters = &v
	return s
}

type TextModerationPlusResponseBody struct {
	// The returned HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The moderation results.
	Data *TextModerationPlusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TextModerationPlusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TextModerationPlusResponseBody) GoString() string {
	return s.String()
}

func (s *TextModerationPlusResponseBody) SetCode(v int32) *TextModerationPlusResponseBody {
	s.Code = &v
	return s
}

func (s *TextModerationPlusResponseBody) SetData(v *TextModerationPlusResponseBodyData) *TextModerationPlusResponseBody {
	s.Data = v
	return s
}

func (s *TextModerationPlusResponseBody) SetMessage(v string) *TextModerationPlusResponseBody {
	s.Message = &v
	return s
}

func (s *TextModerationPlusResponseBody) SetRequestId(v string) *TextModerationPlusResponseBody {
	s.RequestId = &v
	return s
}

type TextModerationPlusResponseBodyData struct {
	// The suggestion.
	Advice []*TextModerationPlusResponseBodyDataAdvice `json:"Advice,omitempty" xml:"Advice,omitempty" type:"Repeated"`
	// The level of prompt attack
	//
	// example:
	//
	// none
	AttackLevel *string `json:"AttackLevel,omitempty" xml:"AttackLevel,omitempty"`
	// The result of prompt attack detect
	AttackResult []*TextModerationPlusResponseBodyDataAttackResult `json:"AttackResult,omitempty" xml:"AttackResult,omitempty" type:"Repeated"`
	// The id of data
	//
	// example:
	//
	// text1234
	DataId       *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	ManualTaskId *string `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// The results.
	Result []*TextModerationPlusResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Risk Level
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The score.
	//
	// example:
	//
	// 1
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The level of sensitivity data
	//
	// example:
	//
	// S0
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	// The result of sensitivity data detect
	SensitiveResult []*TextModerationPlusResponseBodyDataSensitiveResult `json:"SensitiveResult,omitempty" xml:"SensitiveResult,omitempty" type:"Repeated"`
}

func (s TextModerationPlusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TextModerationPlusResponseBodyData) GoString() string {
	return s.String()
}

func (s *TextModerationPlusResponseBodyData) SetAdvice(v []*TextModerationPlusResponseBodyDataAdvice) *TextModerationPlusResponseBodyData {
	s.Advice = v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetAttackLevel(v string) *TextModerationPlusResponseBodyData {
	s.AttackLevel = &v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetAttackResult(v []*TextModerationPlusResponseBodyDataAttackResult) *TextModerationPlusResponseBodyData {
	s.AttackResult = v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetDataId(v string) *TextModerationPlusResponseBodyData {
	s.DataId = &v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetManualTaskId(v string) *TextModerationPlusResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetResult(v []*TextModerationPlusResponseBodyDataResult) *TextModerationPlusResponseBodyData {
	s.Result = v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetRiskLevel(v string) *TextModerationPlusResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetScore(v float32) *TextModerationPlusResponseBodyData {
	s.Score = &v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetSensitiveLevel(v string) *TextModerationPlusResponseBodyData {
	s.SensitiveLevel = &v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetSensitiveResult(v []*TextModerationPlusResponseBodyDataSensitiveResult) *TextModerationPlusResponseBodyData {
	s.SensitiveResult = v
	return s
}

type TextModerationPlusResponseBodyDataAdvice struct {
	// The answer.
	//
	// example:
	//
	// XXX
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// Hit Label
	//
	// example:
	//
	// xxx
	HitLabel *string `json:"HitLabel,omitempty" xml:"HitLabel,omitempty"`
	// Hit Library Name
	//
	// example:
	//
	// xxx
	HitLibName *string `json:"HitLibName,omitempty" xml:"HitLibName,omitempty"`
}

func (s TextModerationPlusResponseBodyDataAdvice) String() string {
	return tea.Prettify(s)
}

func (s TextModerationPlusResponseBodyDataAdvice) GoString() string {
	return s.String()
}

func (s *TextModerationPlusResponseBodyDataAdvice) SetAnswer(v string) *TextModerationPlusResponseBodyDataAdvice {
	s.Answer = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataAdvice) SetHitLabel(v string) *TextModerationPlusResponseBodyDataAdvice {
	s.HitLabel = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataAdvice) SetHitLibName(v string) *TextModerationPlusResponseBodyDataAdvice {
	s.HitLibName = &v
	return s
}

type TextModerationPlusResponseBodyDataAttackResult struct {
	// The level of prompt attack
	//
	// example:
	//
	// none
	AttackLevel *string `json:"AttackLevel,omitempty" xml:"AttackLevel,omitempty"`
	// The confidence
	//
	// example:
	//
	// 0
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// Description
	//
	// example:
	//
	// safe
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label
	//
	// example:
	//
	// safe
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s TextModerationPlusResponseBodyDataAttackResult) String() string {
	return tea.Prettify(s)
}

func (s TextModerationPlusResponseBodyDataAttackResult) GoString() string {
	return s.String()
}

func (s *TextModerationPlusResponseBodyDataAttackResult) SetAttackLevel(v string) *TextModerationPlusResponseBodyDataAttackResult {
	s.AttackLevel = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataAttackResult) SetConfidence(v float32) *TextModerationPlusResponseBodyDataAttackResult {
	s.Confidence = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataAttackResult) SetDescription(v string) *TextModerationPlusResponseBodyDataAttackResult {
	s.Description = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataAttackResult) SetLabel(v string) *TextModerationPlusResponseBodyDataAttackResult {
	s.Label = &v
	return s
}

type TextModerationPlusResponseBodyDataResult struct {
	// The score of the confidence level. Valid values: 0 to 100. The value is accurate to two decimal places.
	//
	// example:
	//
	// 81.22
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The custom term hit by the moderated content.
	CustomizedHit []*TextModerationPlusResponseBodyDataResultCustomizedHit `json:"CustomizedHit,omitempty" xml:"CustomizedHit,omitempty" type:"Repeated"`
	// The description of the label.
	//
	// example:
	//
	// none
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The term hit by the moderated content.
	//
	// example:
	//
	// XXX
	RiskWords *string `json:"RiskWords,omitempty" xml:"RiskWords,omitempty"`
}

func (s TextModerationPlusResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s TextModerationPlusResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *TextModerationPlusResponseBodyDataResult) SetConfidence(v float32) *TextModerationPlusResponseBodyDataResult {
	s.Confidence = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataResult) SetCustomizedHit(v []*TextModerationPlusResponseBodyDataResultCustomizedHit) *TextModerationPlusResponseBodyDataResult {
	s.CustomizedHit = v
	return s
}

func (s *TextModerationPlusResponseBodyDataResult) SetDescription(v string) *TextModerationPlusResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataResult) SetLabel(v string) *TextModerationPlusResponseBodyDataResult {
	s.Label = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataResult) SetRiskWords(v string) *TextModerationPlusResponseBodyDataResult {
	s.RiskWords = &v
	return s
}

type TextModerationPlusResponseBodyDataResultCustomizedHit struct {
	// The terms that are hit. Multiple terms are separated by commas (,).
	//
	// example:
	//
	// xxx
	KeyWords *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// The library name.
	//
	// example:
	//
	// test
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s TextModerationPlusResponseBodyDataResultCustomizedHit) String() string {
	return tea.Prettify(s)
}

func (s TextModerationPlusResponseBodyDataResultCustomizedHit) GoString() string {
	return s.String()
}

func (s *TextModerationPlusResponseBodyDataResultCustomizedHit) SetKeyWords(v string) *TextModerationPlusResponseBodyDataResultCustomizedHit {
	s.KeyWords = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataResultCustomizedHit) SetLibName(v string) *TextModerationPlusResponseBodyDataResultCustomizedHit {
	s.LibName = &v
	return s
}

type TextModerationPlusResponseBodyDataSensitiveResult struct {
	// Description
	//
	// example:
	//
	// xxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label
	//
	// example:
	//
	// 1234
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The sensitive data.
	SensitiveData []*string `json:"SensitiveData,omitempty" xml:"SensitiveData,omitempty" type:"Repeated"`
	// The level of sensitivity data
	//
	// example:
	//
	// S1
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
}

func (s TextModerationPlusResponseBodyDataSensitiveResult) String() string {
	return tea.Prettify(s)
}

func (s TextModerationPlusResponseBodyDataSensitiveResult) GoString() string {
	return s.String()
}

func (s *TextModerationPlusResponseBodyDataSensitiveResult) SetDescription(v string) *TextModerationPlusResponseBodyDataSensitiveResult {
	s.Description = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataSensitiveResult) SetLabel(v string) *TextModerationPlusResponseBodyDataSensitiveResult {
	s.Label = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataSensitiveResult) SetSensitiveData(v []*string) *TextModerationPlusResponseBodyDataSensitiveResult {
	s.SensitiveData = v
	return s
}

func (s *TextModerationPlusResponseBodyDataSensitiveResult) SetSensitiveLevel(v string) *TextModerationPlusResponseBodyDataSensitiveResult {
	s.SensitiveLevel = &v
	return s
}

type TextModerationPlusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextModerationPlusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TextModerationPlusResponse) String() string {
	return tea.Prettify(s)
}

func (s TextModerationPlusResponse) GoString() string {
	return s.String()
}

func (s *TextModerationPlusResponse) SetHeaders(v map[string]*string) *TextModerationPlusResponse {
	s.Headers = v
	return s
}

func (s *TextModerationPlusResponse) SetStatusCode(v int32) *TextModerationPlusResponse {
	s.StatusCode = &v
	return s
}

func (s *TextModerationPlusResponse) SetBody(v *TextModerationPlusResponseBody) *TextModerationPlusResponse {
	s.Body = v
	return s
}

type UrlAsyncModerationRequest struct {
	// The type of the moderation service.
	//
	// example:
	//
	// url_detection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// example:
	//
	// {
	//
	//         "url": "https://help.aliyun.com/",
	//
	//         "dataId": "url123******"
	//
	// }
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s UrlAsyncModerationRequest) String() string {
	return tea.Prettify(s)
}

func (s UrlAsyncModerationRequest) GoString() string {
	return s.String()
}

func (s *UrlAsyncModerationRequest) SetService(v string) *UrlAsyncModerationRequest {
	s.Service = &v
	return s
}

func (s *UrlAsyncModerationRequest) SetServiceParameters(v string) *UrlAsyncModerationRequest {
	s.ServiceParameters = &v
	return s
}

type UrlAsyncModerationResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *UrlAsyncModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6CF2815C-****-****-B52E-FF6E2****492
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UrlAsyncModerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UrlAsyncModerationResponseBody) GoString() string {
	return s.String()
}

func (s *UrlAsyncModerationResponseBody) SetCode(v int32) *UrlAsyncModerationResponseBody {
	s.Code = &v
	return s
}

func (s *UrlAsyncModerationResponseBody) SetData(v *UrlAsyncModerationResponseBodyData) *UrlAsyncModerationResponseBody {
	s.Data = v
	return s
}

func (s *UrlAsyncModerationResponseBody) SetMsg(v string) *UrlAsyncModerationResponseBody {
	s.Msg = &v
	return s
}

func (s *UrlAsyncModerationResponseBody) SetRequestId(v string) *UrlAsyncModerationResponseBody {
	s.RequestId = &v
	return s
}

type UrlAsyncModerationResponseBodyData struct {
	// The ID of the moderated object.
	//
	// example:
	//
	// 26769ada6e264e7ba9aa048241e12be9
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The reqId field returned by the Url Async Moderation API.
	//
	// example:
	//
	// A07B3DB9-D762-5C56-95B1-8EC55CF176D2
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
}

func (s UrlAsyncModerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UrlAsyncModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *UrlAsyncModerationResponseBodyData) SetDataId(v string) *UrlAsyncModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *UrlAsyncModerationResponseBodyData) SetReqId(v string) *UrlAsyncModerationResponseBodyData {
	s.ReqId = &v
	return s
}

type UrlAsyncModerationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UrlAsyncModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UrlAsyncModerationResponse) String() string {
	return tea.Prettify(s)
}

func (s UrlAsyncModerationResponse) GoString() string {
	return s.String()
}

func (s *UrlAsyncModerationResponse) SetHeaders(v map[string]*string) *UrlAsyncModerationResponse {
	s.Headers = v
	return s
}

func (s *UrlAsyncModerationResponse) SetStatusCode(v int32) *UrlAsyncModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *UrlAsyncModerationResponse) SetBody(v *UrlAsyncModerationResponseBody) *UrlAsyncModerationResponse {
	s.Body = v
	return s
}

type VideoModerationRequest struct {
	// The type of the moderation service.
	//
	// example:
	//
	// videoDetection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// example:
	//
	// {\\"url\\": \\"https://talesofai.oss-cn-shanghai.aliyuncs.com/xxx.mp4\\", \\"dataId\\": \\"94db0b88-f521-11ed-806e-fae21c1f239c\\"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s VideoModerationRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationRequest) GoString() string {
	return s.String()
}

func (s *VideoModerationRequest) SetService(v string) *VideoModerationRequest {
	s.Service = &v
	return s
}

func (s *VideoModerationRequest) SetServiceParameters(v string) *VideoModerationRequest {
	s.ServiceParameters = &v
	return s
}

type VideoModerationResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *VideoModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VideoModerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResponseBody) GoString() string {
	return s.String()
}

func (s *VideoModerationResponseBody) SetCode(v int32) *VideoModerationResponseBody {
	s.Code = &v
	return s
}

func (s *VideoModerationResponseBody) SetData(v *VideoModerationResponseBodyData) *VideoModerationResponseBody {
	s.Data = v
	return s
}

func (s *VideoModerationResponseBody) SetMessage(v string) *VideoModerationResponseBody {
	s.Message = &v
	return s
}

func (s *VideoModerationResponseBody) SetRequestId(v string) *VideoModerationResponseBody {
	s.RequestId = &v
	return s
}

type VideoModerationResponseBodyData struct {
	// The ID of the moderated object.
	//
	// example:
	//
	// data1234
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s VideoModerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *VideoModerationResponseBodyData) SetDataId(v string) *VideoModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *VideoModerationResponseBodyData) SetTaskId(v string) *VideoModerationResponseBodyData {
	s.TaskId = &v
	return s
}

type VideoModerationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VideoModerationResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResponse) GoString() string {
	return s.String()
}

func (s *VideoModerationResponse) SetHeaders(v map[string]*string) *VideoModerationResponse {
	s.Headers = v
	return s
}

func (s *VideoModerationResponse) SetStatusCode(v int32) *VideoModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoModerationResponse) SetBody(v *VideoModerationResponseBody) *VideoModerationResponse {
	s.Body = v
	return s
}

type VideoModerationCancelRequest struct {
	// The type of the moderation service.
	//
	// example:
	//
	// videoDetection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// example:
	//
	// {\\"taskId\\":\\"vi_s_4O9gp7GfNQdx9GOqdekFmk-1z2RJT\\"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s VideoModerationCancelRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationCancelRequest) GoString() string {
	return s.String()
}

func (s *VideoModerationCancelRequest) SetService(v string) *VideoModerationCancelRequest {
	s.Service = &v
	return s
}

func (s *VideoModerationCancelRequest) SetServiceParameters(v string) *VideoModerationCancelRequest {
	s.ServiceParameters = &v
	return s
}

type VideoModerationCancelResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6CF2815C-****-****-B52E-FF6E2****492
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VideoModerationCancelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationCancelResponseBody) GoString() string {
	return s.String()
}

func (s *VideoModerationCancelResponseBody) SetCode(v int32) *VideoModerationCancelResponseBody {
	s.Code = &v
	return s
}

func (s *VideoModerationCancelResponseBody) SetMessage(v string) *VideoModerationCancelResponseBody {
	s.Message = &v
	return s
}

func (s *VideoModerationCancelResponseBody) SetRequestId(v string) *VideoModerationCancelResponseBody {
	s.RequestId = &v
	return s
}

type VideoModerationCancelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoModerationCancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VideoModerationCancelResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationCancelResponse) GoString() string {
	return s.String()
}

func (s *VideoModerationCancelResponse) SetHeaders(v map[string]*string) *VideoModerationCancelResponse {
	s.Headers = v
	return s
}

func (s *VideoModerationCancelResponse) SetStatusCode(v int32) *VideoModerationCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoModerationCancelResponse) SetBody(v *VideoModerationCancelResponseBody) *VideoModerationCancelResponse {
	s.Body = v
	return s
}

type VideoModerationResultRequest struct {
	// The type of the moderation service.
	//
	// Valid values:
	//
	// 	- liveStreamDetection: live stream moderation
	//
	// 	- videoDetection: video file moderation
	//
	// 	- liveStreamDetection_cb: live stream moderation_For regions outside the Chinese mainland
	//
	// 	- videoDetection_cb: video file moderation_For regions outside the Chinese mainland.
	//
	// example:
	//
	// videoDetection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The ID of the task that you want to query. You can specify one task ID at a time.
	//
	// example:
	//
	// {\\"taskId\\":\\"au_f_8PoWiZKoLbczp5HRn69VdT-1y8@U5\\"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s VideoModerationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultRequest) GoString() string {
	return s.String()
}

func (s *VideoModerationResultRequest) SetService(v string) *VideoModerationResultRequest {
	s.Service = &v
	return s
}

func (s *VideoModerationResultRequest) SetServiceParameters(v string) *VideoModerationResultRequest {
	s.ServiceParameters = &v
	return s
}

type VideoModerationResultResponseBody struct {
	// The returned HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The moderation results.
	Data *VideoModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// success finished
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6CF2815C-C8C7-4A01-B52E-FF6E24F53492
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VideoModerationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBody) SetCode(v int32) *VideoModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *VideoModerationResultResponseBody) SetData(v *VideoModerationResultResponseBodyData) *VideoModerationResultResponseBody {
	s.Data = v
	return s
}

func (s *VideoModerationResultResponseBody) SetMessage(v string) *VideoModerationResultResponseBody {
	s.Message = &v
	return s
}

func (s *VideoModerationResultResponseBody) SetRequestId(v string) *VideoModerationResultResponseBody {
	s.RequestId = &v
	return s
}

type VideoModerationResultResponseBodyData struct {
	// The voice moderation results. The moderation results contain a structure.
	AudioResult *VideoModerationResultResponseBodyDataAudioResult `json:"AudioResult,omitempty" xml:"AudioResult,omitempty" type:"Struct"`
	// The value of dataId that is specified in the API request. If this parameter is not specified in the API request, the dataId field is not available in the response.
	//
	// example:
	//
	// product_content-2055763
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The image moderation results. If the call is successful, the HTTP status code 200 and moderation results are returned. The moderation results contain a structure.
	FrameResult *VideoModerationResultResponseBodyDataFrameResult `json:"FrameResult,omitempty" xml:"FrameResult,omitempty" type:"Struct"`
	// The unique ID of the live stream.
	//
	// example:
	//
	// liveId
	LiveId       *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	ManualTaskId *string `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// Risk Level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The task ID.
	//
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s VideoModerationResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyData) SetAudioResult(v *VideoModerationResultResponseBodyDataAudioResult) *VideoModerationResultResponseBodyData {
	s.AudioResult = v
	return s
}

func (s *VideoModerationResultResponseBodyData) SetDataId(v string) *VideoModerationResultResponseBodyData {
	s.DataId = &v
	return s
}

func (s *VideoModerationResultResponseBodyData) SetFrameResult(v *VideoModerationResultResponseBodyDataFrameResult) *VideoModerationResultResponseBodyData {
	s.FrameResult = v
	return s
}

func (s *VideoModerationResultResponseBodyData) SetLiveId(v string) *VideoModerationResultResponseBodyData {
	s.LiveId = &v
	return s
}

func (s *VideoModerationResultResponseBodyData) SetManualTaskId(v string) *VideoModerationResultResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *VideoModerationResultResponseBodyData) SetRiskLevel(v string) *VideoModerationResultResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *VideoModerationResultResponseBodyData) SetTaskId(v string) *VideoModerationResultResponseBodyData {
	s.TaskId = &v
	return s
}

type VideoModerationResultResponseBodyDataAudioResult struct {
	// Summary of voice labels.
	AudioSummarys []*VideoModerationResultResponseBodyDataAudioResultAudioSummarys `json:"AudioSummarys,omitempty" xml:"AudioSummarys,omitempty" type:"Repeated"`
	// Risk Level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The details about the text in the moderated voice. The value is a JSON array that contains one or more elements. Each element corresponds to a text entry.
	SliceDetails []*VideoModerationResultResponseBodyDataAudioResultSliceDetails `json:"SliceDetails,omitempty" xml:"SliceDetails,omitempty" type:"Repeated"`
}

func (s VideoModerationResultResponseBodyDataAudioResult) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataAudioResult) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataAudioResult) SetAudioSummarys(v []*VideoModerationResultResponseBodyDataAudioResultAudioSummarys) *VideoModerationResultResponseBodyDataAudioResult {
	s.AudioSummarys = v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResult) SetRiskLevel(v string) *VideoModerationResultResponseBodyDataAudioResult {
	s.RiskLevel = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResult) SetSliceDetails(v []*VideoModerationResultResponseBodyDataAudioResultSliceDetails) *VideoModerationResultResponseBodyDataAudioResult {
	s.SliceDetails = v
	return s
}

type VideoModerationResultResponseBodyDataAudioResultAudioSummarys struct {
	// The description of the labels.
	//
	// example:
	//
	// no risk
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The voice label.
	//
	// example:
	//
	// profanity
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The number of times that the label is matched.
	//
	// example:
	//
	// 8
	LabelSum *int32 `json:"LabelSum,omitempty" xml:"LabelSum,omitempty"`
}

func (s VideoModerationResultResponseBodyDataAudioResultAudioSummarys) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataAudioResultAudioSummarys) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataAudioResultAudioSummarys) SetDescription(v string) *VideoModerationResultResponseBodyDataAudioResultAudioSummarys {
	s.Description = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultAudioSummarys) SetLabel(v string) *VideoModerationResultResponseBodyDataAudioResultAudioSummarys {
	s.Label = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultAudioSummarys) SetLabelSum(v int32) *VideoModerationResultResponseBodyDataAudioResultAudioSummarys {
	s.LabelSum = &v
	return s
}

type VideoModerationResultResponseBodyDataAudioResultSliceDetails struct {
	// The description of the labels.
	//
	// example:
	//
	// no risk
	Descriptions *string `json:"Descriptions,omitempty" xml:"Descriptions,omitempty"`
	// The end time of the text after voice-to-text conversion. Unit: seconds.
	//
	// example:
	//
	// 30
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The end timestamp of the segment. Unit: milliseconds.
	//
	// example:
	//
	// 1685245261939
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// {\\"consoleProduct\\":\\"slbnext\\"}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The details of the labels.
	//
	// example:
	//
	// porn
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// Risk Level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// Subcategory labels. Multiple labels are separated by commas (,).
	//
	// example:
	//
	// ""
	RiskTips *string `json:"RiskTips,omitempty" xml:"RiskTips,omitempty"`
	// The risk words that are hit. Multiple words are separated by commas (,).
	//
	// example:
	//
	// ""
	RiskWords *string `json:"RiskWords,omitempty" xml:"RiskWords,omitempty"`
	// The risk score. Default range: 0 to 99.
	//
	// example:
	//
	// 5
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The start time of the text after voice-to-text conversion. Unit: seconds.
	//
	// example:
	//
	// 0
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The start timestamp of the segment. Unit: milliseconds.
	//
	// example:
	//
	// 1659935002123
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The text converted from voice.
	//
	// example:
	//
	// Disgusting
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// If the moderation object is a voice stream, this parameter indicates the temporary access URL of the voice stream to which the text entry corresponds. The validity period of the URL is 30 minutes. You must prepare another URL to store the voice stream at the earliest opportunity.
	//
	// example:
	//
	// http://xxxx.abc.img
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s VideoModerationResultResponseBodyDataAudioResultSliceDetails) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataAudioResultSliceDetails) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetDescriptions(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.Descriptions = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetEndTime(v int64) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.EndTime = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetEndTimestamp(v int64) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.EndTimestamp = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetExtend(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.Extend = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetLabels(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.Labels = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetRiskLevel(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.RiskLevel = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetRiskTips(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.RiskTips = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetRiskWords(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.RiskWords = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetScore(v float32) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.Score = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetStartTime(v int64) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.StartTime = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetStartTimestamp(v int64) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.StartTimestamp = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetText(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.Text = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataAudioResultSliceDetails) SetUrl(v string) *VideoModerationResultResponseBodyDataAudioResultSliceDetails {
	s.Url = &v
	return s
}

type VideoModerationResultResponseBodyDataFrameResult struct {
	// The number of captured frames that are returned for the video file.
	//
	// example:
	//
	// 10
	FrameNum *int32 `json:"FrameNum,omitempty" xml:"FrameNum,omitempty"`
	// The summary of the labels against which captured frames are matched.
	FrameSummarys []*VideoModerationResultResponseBodyDataFrameResultFrameSummarys `json:"FrameSummarys,omitempty" xml:"FrameSummarys,omitempty" type:"Repeated"`
	// The information about the frames that match the labels.
	Frames []*VideoModerationResultResponseBodyDataFrameResultFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
	// Risk Level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResult) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResult) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResult) SetFrameNum(v int32) *VideoModerationResultResponseBodyDataFrameResult {
	s.FrameNum = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResult) SetFrameSummarys(v []*VideoModerationResultResponseBodyDataFrameResultFrameSummarys) *VideoModerationResultResponseBodyDataFrameResult {
	s.FrameSummarys = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResult) SetFrames(v []*VideoModerationResultResponseBodyDataFrameResultFrames) *VideoModerationResultResponseBodyDataFrameResult {
	s.Frames = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResult) SetRiskLevel(v string) *VideoModerationResultResponseBodyDataFrameResult {
	s.RiskLevel = &v
	return s
}

type VideoModerationResultResponseBodyDataFrameResultFrameSummarys struct {
	// The description of the result.
	//
	// example:
	//
	// no risk
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label against which a captured frame is matched.
	//
	// example:
	//
	// violent_armedForces
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The number of times that the label is matched.
	//
	// example:
	//
	// 8
	LabelSum *int32 `json:"LabelSum,omitempty" xml:"LabelSum,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFrameSummarys) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFrameSummarys) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrameSummarys) SetDescription(v string) *VideoModerationResultResponseBodyDataFrameResultFrameSummarys {
	s.Description = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrameSummarys) SetLabel(v string) *VideoModerationResultResponseBodyDataFrameResultFrameSummarys {
	s.Label = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrameSummarys) SetLabelSum(v int32) *VideoModerationResultResponseBodyDataFrameResultFrameSummarys {
	s.LabelSum = &v
	return s
}

type VideoModerationResultResponseBodyDataFrameResultFrames struct {
	// The interval between the start of the video file and the captured frame. Unit: seconds.
	//
	// example:
	//
	// 338
	Offset *float32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The results of frame moderation parameters such as the label parameter and the confidence parameter.
	Results []*VideoModerationResultResponseBodyDataFrameResultFramesResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// Risk Level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The temporary URL of a captured frame.
	//
	// example:
	//
	// http://xxxx.abc.jpg
	TempUrl *string `json:"TempUrl,omitempty" xml:"TempUrl,omitempty"`
	// The absolute timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1684559739000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFrames) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFrames) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) SetOffset(v float32) *VideoModerationResultResponseBodyDataFrameResultFrames {
	s.Offset = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) SetResults(v []*VideoModerationResultResponseBodyDataFrameResultFramesResults) *VideoModerationResultResponseBodyDataFrameResultFrames {
	s.Results = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) SetRiskLevel(v string) *VideoModerationResultResponseBodyDataFrameResultFrames {
	s.RiskLevel = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) SetTempUrl(v string) *VideoModerationResultResponseBodyDataFrameResultFrames {
	s.TempUrl = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) SetTimestamp(v int64) *VideoModerationResultResponseBodyDataFrameResultFrames {
	s.Timestamp = &v
	return s
}

type VideoModerationResultResponseBodyDataFrameResultFramesResults struct {
	// If a custom image library is hit, information about the custom image library is returned.
	CustomImage []*VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage `json:"CustomImage,omitempty" xml:"CustomImage,omitempty" type:"Repeated"`
	// Returns logo information when the video contains a logo.
	LogoData []*VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData `json:"LogoData,omitempty" xml:"LogoData,omitempty" type:"Repeated"`
	// If the video contains a specific figure, the code of the identified figure is returned.
	PublicFigure []*VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure `json:"PublicFigure,omitempty" xml:"PublicFigure,omitempty" type:"Repeated"`
	// The results of frame moderation parameters such as the label parameter and the confidence parameter.
	Result []*VideoModerationResultResponseBodyDataFrameResultFramesResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The moderation service that is called.
	//
	// example:
	//
	// tonalityImprove
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The information about the text hit in the image is returned.
	TextInImage map[string]interface{} `json:"TextInImage,omitempty" xml:"TextInImage,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResults) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResults) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) SetCustomImage(v []*VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage) *VideoModerationResultResponseBodyDataFrameResultFramesResults {
	s.CustomImage = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) SetLogoData(v []*VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData) *VideoModerationResultResponseBodyDataFrameResultFramesResults {
	s.LogoData = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) SetPublicFigure(v []*VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure) *VideoModerationResultResponseBodyDataFrameResultFramesResults {
	s.PublicFigure = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) SetResult(v []*VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) *VideoModerationResultResponseBodyDataFrameResultFramesResults {
	s.Result = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) SetService(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResults {
	s.Service = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) SetTextInImage(v map[string]interface{}) *VideoModerationResultResponseBodyDataFrameResultFramesResults {
	s.TextInImage = v
	return s
}

type VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage struct {
	// The ID of the custom image that is hit.
	//
	// example:
	//
	// 1234
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the custom image library that is hit.
	//
	// example:
	//
	// 12345678
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage) SetImageId(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage {
	s.ImageId = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage) SetLibId(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsCustomImage {
	s.LibId = &v
	return s
}

type VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData struct {
	// The location of the logo.
	Location *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// Logo information.
	Logo []*VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo `json:"Logo,omitempty" xml:"Logo,omitempty" type:"Repeated"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData) SetLocation(v *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData {
	s.Location = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData) SetLogo(v []*VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoData {
	s.Logo = v
	return s
}

type VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation struct {
	// The height of the text area. Unit: pixels.
	//
	// example:
	//
	// 111
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// The width of the text area. Unit: pixels.
	//
	// example:
	//
	// 111
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// The distance from the top-left corner of the text area to the y-axis, with the top-left corner of the image as the origin. Unit: pixels.
	//
	// example:
	//
	// 111
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// The distance from the top-left corner of the text area to the x-axis, with the top-left corner of the image as the origin. Unit: pixels.
	//
	// example:
	//
	// 222
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) SetH(v int32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation {
	s.H = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) SetW(v int32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation {
	s.W = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) SetX(v int32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation {
	s.X = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation) SetY(v int32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLocation {
	s.Y = &v
	return s
}

type VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo struct {
	// Confidence score, ranging from 0 to 100, with two decimal places.
	//
	// example:
	//
	// 99.10
	Confidence *int64 `json:"confidence,omitempty" xml:"confidence,omitempty"`
	// label
	//
	// example:
	//
	// pt_logotoSocialNetwork
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// Logo name.
	//
	// example:
	//
	// **logo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo) SetConfidence(v int64) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo {
	s.Confidence = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo) SetLabel(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo {
	s.Label = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo) SetName(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsLogoDataLogo {
	s.Name = &v
	return s
}

type VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure struct {
	// The information about the code of the identified figure.
	//
	// example:
	//
	// xxx001
	FigureId   *string                                                                              `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
	FigureName *string                                                                              `json:"FigureName,omitempty" xml:"FigureName,omitempty"`
	Location   []*VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Repeated"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure) SetFigureId(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure {
	s.FigureId = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure) SetFigureName(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure {
	s.FigureName = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure) SetLocation(v []*VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigure {
	s.Location = v
	return s
}

type VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation struct {
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) SetH(v int32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation {
	s.H = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) SetW(v int32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation {
	s.W = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) SetX(v int32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation {
	s.X = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation) SetY(v int32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsPublicFigureLocation {
	s.Y = &v
	return s
}

type VideoModerationResultResponseBodyDataFrameResultFramesResultsResult struct {
	// The score of the confidence level. Valid values: 0 to 100. The value is accurate to two decimal places.
	//
	// example:
	//
	// 50
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The description of the result.
	//
	// example:
	//
	// no risk
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label returned after a frame is moderated. Multiple risk labels and the corresponding scores of confidence levels may be returned for a frame.
	//
	// example:
	//
	// bloody
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) SetConfidence(v float32) *VideoModerationResultResponseBodyDataFrameResultFramesResultsResult {
	s.Confidence = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) SetDescription(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsResult {
	s.Description = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) SetLabel(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResultsResult {
	s.Label = &v
	return s
}

type VideoModerationResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VideoModerationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponse) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponse) SetHeaders(v map[string]*string) *VideoModerationResultResponse {
	s.Headers = v
	return s
}

func (s *VideoModerationResultResponse) SetStatusCode(v int32) *VideoModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoModerationResultResponse) SetBody(v *VideoModerationResultResponseBody) *VideoModerationResultResponse {
	s.Body = v
	return s
}

type VoiceModerationRequest struct {
	// The type of the moderation service.
	//
	// example:
	//
	// nickname_detection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// example:
	//
	// {"url": "http://aliyundoc.com/test.flv", "dataId": "data1234"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s VoiceModerationRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationRequest) GoString() string {
	return s.String()
}

func (s *VoiceModerationRequest) SetService(v string) *VoiceModerationRequest {
	s.Service = &v
	return s
}

func (s *VoiceModerationRequest) SetServiceParameters(v string) *VoiceModerationRequest {
	s.ServiceParameters = &v
	return s
}

type VoiceModerationResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *VoiceModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceModerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceModerationResponseBody) SetCode(v int32) *VoiceModerationResponseBody {
	s.Code = &v
	return s
}

func (s *VoiceModerationResponseBody) SetData(v *VoiceModerationResponseBodyData) *VoiceModerationResponseBody {
	s.Data = v
	return s
}

func (s *VoiceModerationResponseBody) SetMessage(v string) *VoiceModerationResponseBody {
	s.Message = &v
	return s
}

func (s *VoiceModerationResponseBody) SetRequestId(v string) *VoiceModerationResponseBody {
	s.RequestId = &v
	return s
}

type VoiceModerationResponseBodyData struct {
	// The ID of the moderated object.
	//
	// example:
	//
	// data1234
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s VoiceModerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *VoiceModerationResponseBodyData) SetDataId(v string) *VoiceModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *VoiceModerationResponseBodyData) SetTaskId(v string) *VoiceModerationResponseBodyData {
	s.TaskId = &v
	return s
}

type VoiceModerationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VoiceModerationResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResponse) GoString() string {
	return s.String()
}

func (s *VoiceModerationResponse) SetHeaders(v map[string]*string) *VoiceModerationResponse {
	s.Headers = v
	return s
}

func (s *VoiceModerationResponse) SetStatusCode(v int32) *VoiceModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceModerationResponse) SetBody(v *VoiceModerationResponseBody) *VoiceModerationResponse {
	s.Body = v
	return s
}

type VoiceModerationCancelRequest struct {
	// The type of the moderation service.
	//
	// example:
	//
	// nickname_detection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// example:
	//
	// {"taskId":"xxxxx-xxxx"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s VoiceModerationCancelRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationCancelRequest) GoString() string {
	return s.String()
}

func (s *VoiceModerationCancelRequest) SetService(v string) *VoiceModerationCancelRequest {
	s.Service = &v
	return s
}

func (s *VoiceModerationCancelRequest) SetServiceParameters(v string) *VoiceModerationCancelRequest {
	s.ServiceParameters = &v
	return s
}

type VoiceModerationCancelResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4A926AE2-4C96-573F-824F-0532960799F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceModerationCancelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationCancelResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceModerationCancelResponseBody) SetCode(v int32) *VoiceModerationCancelResponseBody {
	s.Code = &v
	return s
}

func (s *VoiceModerationCancelResponseBody) SetMessage(v string) *VoiceModerationCancelResponseBody {
	s.Message = &v
	return s
}

func (s *VoiceModerationCancelResponseBody) SetRequestId(v string) *VoiceModerationCancelResponseBody {
	s.RequestId = &v
	return s
}

type VoiceModerationCancelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceModerationCancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VoiceModerationCancelResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationCancelResponse) GoString() string {
	return s.String()
}

func (s *VoiceModerationCancelResponse) SetHeaders(v map[string]*string) *VoiceModerationCancelResponse {
	s.Headers = v
	return s
}

func (s *VoiceModerationCancelResponse) SetStatusCode(v int32) *VoiceModerationCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceModerationCancelResponse) SetBody(v *VoiceModerationCancelResponseBody) *VoiceModerationCancelResponse {
	s.Body = v
	return s
}

type VoiceModerationResultRequest struct {
	// The type of the moderation service. Valid values: nickname_detection: user nickname
	//
	// example:
	//
	// nickname_detection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters of API requests that are sent from API Gateway to the backend service.
	//
	// For more information, see [ServiceParameter](https://help.aliyun.com/document_detail/43988.html).
	//
	// example:
	//
	// {"taskId":"xxxxx-xxxx"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s VoiceModerationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResultRequest) GoString() string {
	return s.String()
}

func (s *VoiceModerationResultRequest) SetService(v string) *VoiceModerationResultRequest {
	s.Service = &v
	return s
}

func (s *VoiceModerationResultRequest) SetServiceParameters(v string) *VoiceModerationResultRequest {
	s.ServiceParameters = &v
	return s
}

type VoiceModerationResultResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *VoiceModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2881AD4F-638B-52A3-BA20-F74C5B1CEAE3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceModerationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceModerationResultResponseBody) SetCode(v int32) *VoiceModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *VoiceModerationResultResponseBody) SetData(v *VoiceModerationResultResponseBodyData) *VoiceModerationResultResponseBody {
	s.Data = v
	return s
}

func (s *VoiceModerationResultResponseBody) SetMessage(v string) *VoiceModerationResultResponseBody {
	s.Message = &v
	return s
}

func (s *VoiceModerationResultResponseBody) SetRequestId(v string) *VoiceModerationResultResponseBody {
	s.RequestId = &v
	return s
}

type VoiceModerationResultResponseBodyData struct {
	// The ID of the moderated object.
	//
	// example:
	//
	// 26769ada6e264e7ba9aa048241e12be9
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The unique ID of the live stream.
	//
	// example:
	//
	// liveId
	LiveId       *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	ManualTaskId *string `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// Risk Level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The moderation results of audio segments.
	SliceDetails []*VoiceModerationResultResponseBodyDataSliceDetails `json:"SliceDetails,omitempty" xml:"SliceDetails,omitempty" type:"Repeated"`
	// The task ID.
	//
	// example:
	//
	// kw24ihd0WGkdi5nniVZM@qOj-1x5Ibb
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The URL of the moderated content.
	//
	// example:
	//
	// https://aliyundoc.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s VoiceModerationResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *VoiceModerationResultResponseBodyData) SetDataId(v string) *VoiceModerationResultResponseBodyData {
	s.DataId = &v
	return s
}

func (s *VoiceModerationResultResponseBodyData) SetLiveId(v string) *VoiceModerationResultResponseBodyData {
	s.LiveId = &v
	return s
}

func (s *VoiceModerationResultResponseBodyData) SetManualTaskId(v string) *VoiceModerationResultResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *VoiceModerationResultResponseBodyData) SetRiskLevel(v string) *VoiceModerationResultResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *VoiceModerationResultResponseBodyData) SetSliceDetails(v []*VoiceModerationResultResponseBodyDataSliceDetails) *VoiceModerationResultResponseBodyData {
	s.SliceDetails = v
	return s
}

func (s *VoiceModerationResultResponseBodyData) SetTaskId(v string) *VoiceModerationResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *VoiceModerationResultResponseBodyData) SetUrl(v string) *VoiceModerationResultResponseBodyData {
	s.Url = &v
	return s
}

type VoiceModerationResultResponseBodyDataSliceDetails struct {
	// The description of the labels.
	//
	// example:
	//
	// no risk
	Descriptions *string `json:"Descriptions,omitempty" xml:"Descriptions,omitempty"`
	// The end time of the audio segment in seconds.
	//
	// example:
	//
	// 10
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The end timestamp of the segment. Unit: milliseconds.
	//
	// example:
	//
	// 1678854649720
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// Extended fields.
	//
	// example:
	//
	// {\\"riskTips\\":\\"sexuality_Suggestive\\",\\"riskWords\\":\\"pxxxxy\\"}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The details of the labels.
	//
	// example:
	//
	// sexual_sounds
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// {}
	OriginAlgoResult map[string]interface{} `json:"OriginAlgoResult,omitempty" xml:"OriginAlgoResult,omitempty"`
	// Risk Level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The details of the risky content.
	//
	// example:
	//
	// sexuality_Suggestive
	RiskTips *string `json:"RiskTips,omitempty" xml:"RiskTips,omitempty"`
	// The term hit by the risky content.
	//
	// example:
	//
	// AAA,BBB,CCC
	RiskWords *string `json:"RiskWords,omitempty" xml:"RiskWords,omitempty"`
	// The risk score. Default range: 0 to 99.
	//
	// example:
	//
	// 87.01
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The start time of the audio segment in seconds.
	//
	// example:
	//
	// 0
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The start timestamp of the segment. Unit: milliseconds.
	//
	// example:
	//
	// 1678854649720
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The text converted from the audio segment.
	//
	// example:
	//
	// Disgusting
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The temporary URL of the audio segment.
	//
	// example:
	//
	// https://aliyundoc.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s VoiceModerationResultResponseBodyDataSliceDetails) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResultResponseBodyDataSliceDetails) GoString() string {
	return s.String()
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetDescriptions(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Descriptions = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetEndTime(v int64) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.EndTime = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetEndTimestamp(v int64) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.EndTimestamp = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetExtend(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Extend = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetLabels(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Labels = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetOriginAlgoResult(v map[string]interface{}) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.OriginAlgoResult = v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetRiskLevel(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.RiskLevel = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetRiskTips(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.RiskTips = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetRiskWords(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.RiskWords = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetScore(v float32) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Score = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetStartTime(v int64) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.StartTime = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetStartTimestamp(v int64) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.StartTimestamp = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetText(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Text = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetUrl(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Url = &v
	return s
}

type VoiceModerationResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VoiceModerationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResultResponse) GoString() string {
	return s.String()
}

func (s *VoiceModerationResultResponse) SetHeaders(v map[string]*string) *VoiceModerationResultResponse {
	s.Headers = v
	return s
}

func (s *VoiceModerationResultResponse) SetStatusCode(v int32) *VoiceModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceModerationResultResponse) SetBody(v *VoiceModerationResultResponseBody) *VoiceModerationResultResponse {
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
// 文档审核结果
//
// @param request - DescribeFileModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileModerationResultResponse
func (client *Client) DescribeFileModerationResultWithOptions(request *DescribeFileModerationResultRequest, runtime *util.RuntimeOptions) (_result *DescribeFileModerationResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFileModerationResult"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFileModerationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档审核结果
//
// @param request - DescribeFileModerationResultRequest
//
// @return DescribeFileModerationResultResponse
func (client *Client) DescribeFileModerationResult(request *DescribeFileModerationResultRequest) (_result *DescribeFileModerationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFileModerationResultResponse{}
	_body, _err := client.DescribeFileModerationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the moderation results of an Image Moderation 2.0 task.
//
// Description:
//
//	  Billing: This operation is free of charge.
//
//		- QPS limit: You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeImageModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImageModerationResultResponse
func (client *Client) DescribeImageModerationResultWithOptions(request *DescribeImageModerationResultRequest, runtime *util.RuntimeOptions) (_result *DescribeImageModerationResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReqId)) {
		query["ReqId"] = request.ReqId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImageModerationResult"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImageModerationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the moderation results of an Image Moderation 2.0 task.
//
// Description:
//
//	  Billing: This operation is free of charge.
//
//		- QPS limit: You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeImageModerationResultRequest
//
// @return DescribeImageModerationResultResponse
func (client *Client) DescribeImageModerationResult(request *DescribeImageModerationResultRequest) (_result *DescribeImageModerationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageModerationResultResponse{}
	_body, _err := client.DescribeImageModerationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询检测结果辅助信息
//
// @param request - DescribeImageResultExtRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImageResultExtResponse
func (client *Client) DescribeImageResultExtWithOptions(request *DescribeImageResultExtRequest, runtime *util.RuntimeOptions) (_result *DescribeImageResultExtResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InfoType)) {
		body["InfoType"] = request.InfoType
	}

	if !tea.BoolValue(util.IsUnset(request.ReqId)) {
		body["ReqId"] = request.ReqId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImageResultExt"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImageResultExtResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询检测结果辅助信息
//
// @param request - DescribeImageResultExtRequest
//
// @return DescribeImageResultExtResponse
func (client *Client) DescribeImageResultExt(request *DescribeImageResultExtRequest) (_result *DescribeImageResultExtResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageResultExtResponse{}
	_body, _err := client.DescribeImageResultExtWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询上传token
//
// @param request - DescribeUploadTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUploadTokenResponse
func (client *Client) DescribeUploadTokenWithOptions(runtime *util.RuntimeOptions) (_result *DescribeUploadTokenResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeUploadToken"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUploadTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询上传token
//
// @return DescribeUploadTokenResponse
func (client *Client) DescribeUploadToken() (_result *DescribeUploadTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUploadTokenResponse{}
	_body, _err := client.DescribeUploadTokenWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the moderation results based on the ReqId returned by asynchronous URL moderation.
//
// Description:
//
//	  Billing: This operation is free of charge.
//
//		- Query timeout: We recommend that you query moderation results at least 480 seconds after you send an asynchronous moderation request. Content Moderation retains moderation results for up to 3 days. After 3 days, the results are deleted.
//
//		- You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeUrlModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUrlModerationResultResponse
func (client *Client) DescribeUrlModerationResultWithOptions(request *DescribeUrlModerationResultRequest, runtime *util.RuntimeOptions) (_result *DescribeUrlModerationResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReqId)) {
		body["ReqId"] = request.ReqId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUrlModerationResult"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUrlModerationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the moderation results based on the ReqId returned by asynchronous URL moderation.
//
// Description:
//
//	  Billing: This operation is free of charge.
//
//		- Query timeout: We recommend that you query moderation results at least 480 seconds after you send an asynchronous moderation request. Content Moderation retains moderation results for up to 3 days. After 3 days, the results are deleted.
//
//		- You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeUrlModerationResultRequest
//
// @return DescribeUrlModerationResultResponse
func (client *Client) DescribeUrlModerationResult(request *DescribeUrlModerationResultRequest) (_result *DescribeUrlModerationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUrlModerationResultResponse{}
	_body, _err := client.DescribeUrlModerationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档审核
//
// @param request - FileModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileModerationResponse
func (client *Client) FileModerationWithOptions(request *FileModerationRequest, runtime *util.RuntimeOptions) (_result *FileModerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FileModeration"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FileModerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档审核
//
// @param request - FileModerationRequest
//
// @return FileModerationResponse
func (client *Client) FileModeration(request *FileModerationRequest) (_result *FileModerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FileModerationResponse{}
	_body, _err := client.FileModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 图片异步检测
//
// @param request - ImageAsyncModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImageAsyncModerationResponse
func (client *Client) ImageAsyncModerationWithOptions(request *ImageAsyncModerationRequest, runtime *util.RuntimeOptions) (_result *ImageAsyncModerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		query["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		query["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImageAsyncModeration"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImageAsyncModerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片异步检测
//
// @param request - ImageAsyncModerationRequest
//
// @return ImageAsyncModerationResponse
func (client *Client) ImageAsyncModeration(request *ImageAsyncModerationRequest) (_result *ImageAsyncModerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImageAsyncModerationResponse{}
	_body, _err := client.ImageAsyncModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 图片批量调用
//
// @param request - ImageBatchModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImageBatchModerationResponse
func (client *Client) ImageBatchModerationWithOptions(request *ImageBatchModerationRequest, runtime *util.RuntimeOptions) (_result *ImageBatchModerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		query["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		query["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImageBatchModeration"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImageBatchModerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片批量调用
//
// @param request - ImageBatchModerationRequest
//
// @return ImageBatchModerationResponse
func (client *Client) ImageBatchModeration(request *ImageBatchModerationRequest) (_result *ImageBatchModerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImageBatchModerationResponse{}
	_body, _err := client.ImageBatchModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Identifies whether an image contains content or elements that violate relevant regulations on network content dissemination, affect the content order of a specific platform, or affect user experience. Image Moderation 2.0 supports over 90 content risk labels and over 100 risk control items. Image Moderation 2.0 of Content Moderation allows you to develop further moderation or governance measures for specific image content based on business scenarios, platform-specific content governance rules, or rich risk labels and scores of confidence levels returned by API calls.
//
// Description:
//
// *Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/467826.html)[](https://www.aliyun.com/price/product?#/lvwang/detail/cdibag) of Image Moderation 2.0.
//
// @param request - ImageModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImageModerationResponse
func (client *Client) ImageModerationWithOptions(request *ImageModerationRequest, runtime *util.RuntimeOptions) (_result *ImageModerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImageModeration"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImageModerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Identifies whether an image contains content or elements that violate relevant regulations on network content dissemination, affect the content order of a specific platform, or affect user experience. Image Moderation 2.0 supports over 90 content risk labels and over 100 risk control items. Image Moderation 2.0 of Content Moderation allows you to develop further moderation or governance measures for specific image content based on business scenarios, platform-specific content governance rules, or rich risk labels and scores of confidence levels returned by API calls.
//
// Description:
//
// *Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/467826.html)[](https://www.aliyun.com/price/product?#/lvwang/detail/cdibag) of Image Moderation 2.0.
//
// @param request - ImageModerationRequest
//
// @return ImageModerationResponse
func (client *Client) ImageModeration(request *ImageModerationRequest) (_result *ImageModerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImageModerationResponse{}
	_body, _err := client.ImageModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Content Security Manual Review Result Callback Interface
//
// @param request - ManualCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManualCallbackResponse
func (client *Client) ManualCallbackWithOptions(request *ManualCallbackRequest, runtime *util.RuntimeOptions) (_result *ManualCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["Channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.Checksum)) {
		body["Checksum"] = request.Checksum
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Msg)) {
		body["Msg"] = request.Msg
	}

	if !tea.BoolValue(util.IsUnset(request.ReqId)) {
		body["ReqId"] = request.ReqId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ManualCallback"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ManualCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Content Security Manual Review Result Callback Interface
//
// @param request - ManualCallbackRequest
//
// @return ManualCallbackResponse
func (client *Client) ManualCallback(request *ManualCallbackRequest) (_result *ManualCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ManualCallbackResponse{}
	_body, _err := client.ManualCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Content Security Manual Review Request Interface
//
// @param request - ManualModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManualModerationResponse
func (client *Client) ManualModerationWithOptions(request *ManualModerationRequest, runtime *util.RuntimeOptions) (_result *ManualModerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ManualModeration"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ManualModerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Content Security Manual Review Request Interface
//
// @param request - ManualModerationRequest
//
// @return ManualModerationResponse
func (client *Client) ManualModeration(request *ManualModerationRequest) (_result *ManualModerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ManualModerationResponse{}
	_body, _err := client.ManualModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Retrieve manual review results
//
// @param request - ManualModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManualModerationResultResponse
func (client *Client) ManualModerationResultWithOptions(request *ManualModerationResultRequest, runtime *util.RuntimeOptions) (_result *ManualModerationResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ManualModerationResult"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ManualModerationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve manual review results
//
// @param request - ManualModerationResultRequest
//
// @return ManualModerationResultResponse
func (client *Client) ManualModerationResult(request *ManualModerationResultRequest) (_result *ManualModerationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ManualModerationResultResponse{}
	_body, _err := client.ManualModerationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步检测接口
//
// @param request - MultiModalGuardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MultiModalGuardResponse
func (client *Client) MultiModalGuardWithOptions(request *MultiModalGuardRequest, runtime *util.RuntimeOptions) (_result *MultiModalGuardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MultiModalGuard"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MultiModalGuardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步检测接口
//
// @param request - MultiModalGuardRequest
//
// @return MultiModalGuardResponse
func (client *Client) MultiModalGuard(request *MultiModalGuardRequest) (_result *MultiModalGuardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MultiModalGuardResponse{}
	_body, _err := client.MultiModalGuardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Provides moderation services for multiple business scenarios and identifies various violation risks.
//
// Description:
//
// Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/464388.html?#section-itm-m2s-ugq) of Text Moderation 2.0.
//
// @param request - TextModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TextModerationResponse
func (client *Client) TextModerationWithOptions(request *TextModerationRequest, runtime *util.RuntimeOptions) (_result *TextModerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TextModeration"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TextModerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides moderation services for multiple business scenarios and identifies various violation risks.
//
// Description:
//
// Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/464388.html?#section-itm-m2s-ugq) of Text Moderation 2.0.
//
// @param request - TextModerationRequest
//
// @return TextModerationResponse
func (client *Client) TextModeration(request *TextModerationRequest) (_result *TextModerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TextModerationResponse{}
	_body, _err := client.TextModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moderates the input command and generated text of large language models (LLMs). Specific model input commands can be used to retrieve standard answers. The feature of enabling and disabling the moderation labels is also available.
//
// Description:
//
// Before you call this operation, make sure that you have [activated the Content Moderation 2.0 service](https://common-buy.aliyun.com/?commodityCode=lvwang_cip_public_cn) and are familiar with the [billing](https://help.aliyun.com/document_detail/2671445.html?#section-6od-32j-99n) of the Text Moderation 2.0 Plus service.
//
// @param request - TextModerationPlusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TextModerationPlusResponse
func (client *Client) TextModerationPlusWithOptions(request *TextModerationPlusRequest, runtime *util.RuntimeOptions) (_result *TextModerationPlusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TextModerationPlus"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TextModerationPlusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moderates the input command and generated text of large language models (LLMs). Specific model input commands can be used to retrieve standard answers. The feature of enabling and disabling the moderation labels is also available.
//
// Description:
//
// Before you call this operation, make sure that you have [activated the Content Moderation 2.0 service](https://common-buy.aliyun.com/?commodityCode=lvwang_cip_public_cn) and are familiar with the [billing](https://help.aliyun.com/document_detail/2671445.html?#section-6od-32j-99n) of the Text Moderation 2.0 Plus service.
//
// @param request - TextModerationPlusRequest
//
// @return TextModerationPlusResponse
func (client *Client) TextModerationPlus(request *TextModerationPlusRequest) (_result *TextModerationPlusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TextModerationPlusResponse{}
	_body, _err := client.TextModerationPlusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// url异步检测
//
// @param request - UrlAsyncModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UrlAsyncModerationResponse
func (client *Client) UrlAsyncModerationWithOptions(request *UrlAsyncModerationRequest, runtime *util.RuntimeOptions) (_result *UrlAsyncModerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		query["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		query["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UrlAsyncModeration"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UrlAsyncModerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// url异步检测
//
// @param request - UrlAsyncModerationRequest
//
// @return UrlAsyncModerationResponse
func (client *Client) UrlAsyncModeration(request *UrlAsyncModerationRequest) (_result *UrlAsyncModerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UrlAsyncModerationResponse{}
	_body, _err := client.UrlAsyncModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 视频检测任务提交
//
// @param request - VideoModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VideoModerationResponse
func (client *Client) VideoModerationWithOptions(request *VideoModerationRequest, runtime *util.RuntimeOptions) (_result *VideoModerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VideoModeration"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VideoModerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频检测任务提交
//
// @param request - VideoModerationRequest
//
// @return VideoModerationResponse
func (client *Client) VideoModeration(request *VideoModerationRequest) (_result *VideoModerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VideoModerationResponse{}
	_body, _err := client.VideoModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消视频直播流检测
//
// @param request - VideoModerationCancelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VideoModerationCancelResponse
func (client *Client) VideoModerationCancelWithOptions(request *VideoModerationCancelRequest, runtime *util.RuntimeOptions) (_result *VideoModerationCancelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VideoModerationCancel"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VideoModerationCancelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消视频直播流检测
//
// @param request - VideoModerationCancelRequest
//
// @return VideoModerationCancelResponse
func (client *Client) VideoModerationCancel(request *VideoModerationCancelRequest) (_result *VideoModerationCancelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VideoModerationCancelResponse{}
	_body, _err := client.VideoModerationCancelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Video Detection Results
//
// Description:
//
// This operation is free of charge. We recommend that you query moderation results at least 30 seconds after you send an asynchronous moderation request. Content Moderation retains moderation results for at most 24 hours. After 24 hours, the results are deleted.
//
// @param request - VideoModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VideoModerationResultResponse
func (client *Client) VideoModerationResultWithOptions(request *VideoModerationResultRequest, runtime *util.RuntimeOptions) (_result *VideoModerationResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VideoModerationResult"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VideoModerationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Video Detection Results
//
// Description:
//
// This operation is free of charge. We recommend that you query moderation results at least 30 seconds after you send an asynchronous moderation request. Content Moderation retains moderation results for at most 24 hours. After 24 hours, the results are deleted.
//
// @param request - VideoModerationResultRequest
//
// @return VideoModerationResultResponse
func (client *Client) VideoModerationResult(request *VideoModerationResultRequest) (_result *VideoModerationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VideoModerationResultResponse{}
	_body, _err := client.VideoModerationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 语音审核
//
// @param request - VoiceModerationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceModerationResponse
func (client *Client) VoiceModerationWithOptions(request *VoiceModerationRequest, runtime *util.RuntimeOptions) (_result *VoiceModerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VoiceModeration"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceModerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语音审核
//
// @param request - VoiceModerationRequest
//
// @return VoiceModerationResponse
func (client *Client) VoiceModeration(request *VoiceModerationRequest) (_result *VoiceModerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VoiceModerationResponse{}
	_body, _err := client.VoiceModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消检测
//
// @param request - VoiceModerationCancelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceModerationCancelResponse
func (client *Client) VoiceModerationCancelWithOptions(request *VoiceModerationCancelRequest, runtime *util.RuntimeOptions) (_result *VoiceModerationCancelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VoiceModerationCancel"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceModerationCancelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消检测
//
// @param request - VoiceModerationCancelRequest
//
// @return VoiceModerationCancelResponse
func (client *Client) VoiceModerationCancel(request *VoiceModerationCancelRequest) (_result *VoiceModerationCancelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VoiceModerationCancelResponse{}
	_body, _err := client.VoiceModerationCancelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the moderation results of a Voice Moderation 2.0 task.
//
// @param request - VoiceModerationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceModerationResultResponse
func (client *Client) VoiceModerationResultWithOptions(request *VoiceModerationResultRequest, runtime *util.RuntimeOptions) (_result *VoiceModerationResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VoiceModerationResult"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceModerationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the moderation results of a Voice Moderation 2.0 task.
//
// @param request - VoiceModerationResultRequest
//
// @return VoiceModerationResultResponse
func (client *Client) VoiceModerationResult(request *VoiceModerationResultRequest) (_result *VoiceModerationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VoiceModerationResultResponse{}
	_body, _err := client.VoiceModerationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
