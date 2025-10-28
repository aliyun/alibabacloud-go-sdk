// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileModerationResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeFileModerationResultResponseBody
	GetCode() *int32
	SetData(v *DescribeFileModerationResultResponseBodyData) *DescribeFileModerationResultResponseBody
	GetData() *DescribeFileModerationResultResponseBodyData
	SetMessage(v string) *DescribeFileModerationResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeFileModerationResultResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s DescribeFileModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeFileModerationResultResponseBody) GetData() *DescribeFileModerationResultResponseBodyData {
	return s.Data
}

func (s *DescribeFileModerationResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeFileModerationResultResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *DescribeFileModerationResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *DescribeFileModerationResultResponseBodyData) GetDocType() *string {
	return s.DocType
}

func (s *DescribeFileModerationResultResponseBodyData) GetPageResult() []*DescribeFileModerationResultResponseBodyDataPageResult {
	return s.PageResult
}

func (s *DescribeFileModerationResultResponseBodyData) GetPageSummary() *DescribeFileModerationResultResponseBodyDataPageSummary {
	return s.PageSummary
}

func (s *DescribeFileModerationResultResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeFileModerationResultResponseBodyData) GetUrl() *string {
	return s.Url
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

func (s *DescribeFileModerationResultResponseBodyData) Validate() error {
	if s.PageResult != nil {
		for _, item := range s.PageResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageSummary != nil {
		if err := s.PageSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageResult) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageResult) GetImageResult() []*DescribeFileModerationResultResponseBodyDataPageResultImageResult {
	return s.ImageResult
}

func (s *DescribeFileModerationResultResponseBodyDataPageResult) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeFileModerationResultResponseBodyDataPageResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeFileModerationResultResponseBodyDataPageResult) GetTextResult() []*DescribeFileModerationResultResponseBodyDataPageResultTextResult {
	return s.TextResult
}

func (s *DescribeFileModerationResultResponseBodyDataPageResult) GetTextUrl() *string {
	return s.TextUrl
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

func (s *DescribeFileModerationResultResponseBodyDataPageResult) Validate() error {
	if s.ImageResult != nil {
		for _, item := range s.ImageResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TextResult != nil {
		for _, item := range s.TextResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageResultImageResult) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResult) GetLabelResult() []*DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult {
	return s.LabelResult
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResult) GetLocation() *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation {
	return s.Location
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResult) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResult) GetService() *string {
	return s.Service
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

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResult) Validate() error {
	if s.LabelResult != nil {
		for _, item := range s.LabelResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult) GetLabel() *string {
	return s.Label
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

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation) GetH() *int32 {
	return s.H
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation) GetW() *int32 {
	return s.W
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation) GetX() *int32 {
	return s.X
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation) GetY() *int32 {
	return s.Y
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

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageResultTextResult) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) GetDescriptions() *string {
	return s.Descriptions
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) GetLabels() *string {
	return s.Labels
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) GetRiskTips() *string {
	return s.RiskTips
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) GetRiskWords() *string {
	return s.RiskWords
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) GetService() *string {
	return s.Service
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) GetText() *string {
	return s.Text
}

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) GetTextSegment() *string {
	return s.TextSegment
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

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageSummary) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummary) GetImageSummary() *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary {
	return s.ImageSummary
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummary) GetPageSum() *int32 {
	return s.PageSum
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummary) GetTextSummary() *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary {
	return s.TextSummary
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

func (s *DescribeFileModerationResultResponseBodyDataPageSummary) Validate() error {
	if s.ImageSummary != nil {
		if err := s.ImageSummary.Validate(); err != nil {
			return err
		}
	}
	if s.TextSummary != nil {
		if err := s.TextSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary) GetImageLabels() []*DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels {
	return s.ImageLabels
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary) SetImageLabels(v []*DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels) *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary {
	s.ImageLabels = v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary) SetRiskLevel(v string) *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary {
	s.RiskLevel = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummary) Validate() error {
	if s.ImageLabels != nil {
		for _, item := range s.ImageLabels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels) GetDescription() *string {
	return s.Description
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels) GetLabel() *string {
	return s.Label
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels) GetLabelSum() *int32 {
	return s.LabelSum
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

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryImageSummaryImageLabels) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary) GetTextLabels() []*DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels {
	return s.TextLabels
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary) SetRiskLevel(v string) *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary {
	s.RiskLevel = &v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary) SetTextLabels(v []*DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels) *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary {
	s.TextLabels = v
	return s
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummary) Validate() error {
	if s.TextLabels != nil {
		for _, item := range s.TextLabels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	return dara.Prettify(s)
}

func (s DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels) GoString() string {
	return s.String()
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels) GetDescription() *string {
	return s.Description
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels) GetLabel() *string {
	return s.Label
}

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels) GetLabelSum() *int32 {
	return s.LabelSum
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

func (s *DescribeFileModerationResultResponseBodyDataPageSummaryTextSummaryTextLabels) Validate() error {
	return dara.Validate(s)
}
