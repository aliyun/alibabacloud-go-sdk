// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeFileModerationResultRequest struct {
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
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
	Code    *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *DescribeFileModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	DataId     *string                                                   `json:"DataId,omitempty" xml:"DataId,omitempty"`
	DocType    *string                                                   `json:"DocType,omitempty" xml:"DocType,omitempty"`
	PageResult []*DescribeFileModerationResultResponseBodyDataPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Repeated"`
	Url        *string                                                   `json:"Url,omitempty" xml:"Url,omitempty"`
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

func (s *DescribeFileModerationResultResponseBodyData) SetUrl(v string) *DescribeFileModerationResultResponseBodyData {
	s.Url = &v
	return s
}

type DescribeFileModerationResultResponseBodyDataPageResult struct {
	ImageResult []*DescribeFileModerationResultResponseBodyDataPageResultImageResult `json:"ImageResult,omitempty" xml:"ImageResult,omitempty" type:"Repeated"`
	ImageUrl    *string                                                              `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	PageNum     *int32                                                               `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	TextResult  []*DescribeFileModerationResultResponseBodyDataPageResultTextResult  `json:"TextResult,omitempty" xml:"TextResult,omitempty" type:"Repeated"`
	TextUrl     *string                                                              `json:"TextUrl,omitempty" xml:"TextUrl,omitempty"`
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
	Description *string                                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	LabelResult []*DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult `json:"LabelResult,omitempty" xml:"LabelResult,omitempty" type:"Repeated"`
	Location    *DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation      `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	Service     *string                                                                         `json:"Service,omitempty" xml:"Service,omitempty"`
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

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResult) SetService(v string) *DescribeFileModerationResultResponseBodyDataPageResultImageResult {
	s.Service = &v
	return s
}

type DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult struct {
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Label      *string  `json:"Label,omitempty" xml:"Label,omitempty"`
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

func (s *DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult) SetLabel(v string) *DescribeFileModerationResultResponseBodyDataPageResultImageResultLabelResult {
	s.Label = &v
	return s
}

type DescribeFileModerationResultResponseBodyDataPageResultImageResultLocation struct {
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Labels      *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	RiskTips    *string `json:"RiskTips,omitempty" xml:"RiskTips,omitempty"`
	RiskWords   *string `json:"RiskWords,omitempty" xml:"RiskWords,omitempty"`
	Service     *string `json:"Service,omitempty" xml:"Service,omitempty"`
	Text        *string `json:"Text,omitempty" xml:"Text,omitempty"`
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

func (s *DescribeFileModerationResultResponseBodyDataPageResultTextResult) SetLabels(v string) *DescribeFileModerationResultResponseBodyDataPageResultTextResult {
	s.Labels = &v
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
	Code      *int32                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeImageModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                                        `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	DataId   *string                                                `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Frame    *string                                                `json:"Frame,omitempty" xml:"Frame,omitempty"`
	FrameNum *int32                                                 `json:"FrameNum,omitempty" xml:"FrameNum,omitempty"`
	ReqId    *string                                                `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	Result   []*DescribeImageModerationResultResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
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

func (s *DescribeImageModerationResultResponseBodyData) SetReqId(v string) *DescribeImageModerationResultResponseBodyData {
	s.ReqId = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyData) SetResult(v []*DescribeImageModerationResultResponseBodyDataResult) *DescribeImageModerationResultResponseBodyData {
	s.Result = v
	return s
}

type DescribeImageModerationResultResponseBodyDataResult struct {
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Label      *string  `json:"Label,omitempty" xml:"Label,omitempty"`
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

func (s *DescribeImageModerationResultResponseBodyDataResult) SetLabel(v string) *DescribeImageModerationResultResponseBodyDataResult {
	s.Label = &v
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
	InfoType *string `json:"InfoType,omitempty" xml:"InfoType,omitempty"`
	ReqId    *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
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
	Code      *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeImageResultExtResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                                 `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CustomImage  []*DescribeImageResultExtResponseBodyDataCustomImage  `json:"CustomImage,omitempty" xml:"CustomImage,omitempty" type:"Repeated"`
	PublicFigure []*DescribeImageResultExtResponseBodyDataPublicFigure `json:"PublicFigure,omitempty" xml:"PublicFigure,omitempty" type:"Repeated"`
	TextInImage  *DescribeImageResultExtResponseBodyDataTextInImage    `json:"TextInImage,omitempty" xml:"TextInImage,omitempty" type:"Struct"`
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
	// 图片ID。
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 图库ID。
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// 图库名。
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
	// 人物ID。
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
	OcrDatas  []*string `json:"OcrDatas,omitempty" xml:"OcrDatas,omitempty" type:"Repeated"`
	RiskWords []*string `json:"RiskWords,omitempty" xml:"RiskWords,omitempty" type:"Repeated"`
}

func (s DescribeImageResultExtResponseBodyDataTextInImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResultExtResponseBodyDataTextInImage) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtResponseBodyDataTextInImage) SetOcrDatas(v []*string) *DescribeImageResultExtResponseBodyDataTextInImage {
	s.OcrDatas = v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataTextInImage) SetRiskWords(v []*string) *DescribeImageResultExtResponseBodyDataTextInImage {
	s.RiskWords = v
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
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeUploadTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                              `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccessKeyId         *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret     *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	BucketName          *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Expiration          *int32  `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	FileNamePrefix      *string `json:"FileNamePrefix,omitempty" xml:"FileNamePrefix,omitempty"`
	OssInternalEndPoint *string `json:"OssInternalEndPoint,omitempty" xml:"OssInternalEndPoint,omitempty"`
	OssInternetEndPoint *string `json:"OssInternetEndPoint,omitempty" xml:"OssInternetEndPoint,omitempty"`
	SecurityToken       *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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
	Code      *int32                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeUrlModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                                      `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	DataId    *string                                               `json:"DataId,omitempty" xml:"DataId,omitempty"`
	ExtraInfo *DescribeUrlModerationResultResponseBodyDataExtraInfo `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty" type:"Struct"`
	ReqId     *string                                               `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	Result    []*DescribeUrlModerationResultResponseBodyDataResult  `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
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
	IcpNo   *string `json:"IcpNo,omitempty" xml:"IcpNo,omitempty"`
	IcpType *string `json:"IcpType,omitempty" xml:"IcpType,omitempty"`
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

type DescribeUrlModerationResultResponseBodyDataResult struct {
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Label      *string  `json:"Label,omitempty" xml:"Label,omitempty"`
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
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
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
	Code    *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *FileModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
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
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ImageAsyncModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                               `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	ReqId  *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
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

type ImageModerationRequest struct {
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
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
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ImageModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                          `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	DataId *string                                  `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Result []*ImageModerationResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
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

func (s *ImageModerationResponseBodyData) SetResult(v []*ImageModerationResponseBodyDataResult) *ImageModerationResponseBodyData {
	s.Result = v
	return s
}

type ImageModerationResponseBodyDataResult struct {
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Label      *string  `json:"Label,omitempty" xml:"Label,omitempty"`
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

func (s *ImageModerationResponseBodyDataResult) SetLabel(v string) *ImageModerationResponseBodyDataResult {
	s.Label = &v
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

type TextModerationRequest struct {
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
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
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TextModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	DeviceId  *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	Labels    *string `json:"labels,omitempty" xml:"labels,omitempty"`
	Reason    *string `json:"reason,omitempty" xml:"reason,omitempty"`
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

func (s *TextModerationResponseBodyData) SetDeviceId(v string) *TextModerationResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *TextModerationResponseBodyData) SetLabels(v string) *TextModerationResponseBodyData {
	s.Labels = &v
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
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
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
	Code    *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *TextModerationPlusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Advice []*TextModerationPlusResponseBodyDataAdvice `json:"Advice,omitempty" xml:"Advice,omitempty" type:"Repeated"`
	Result []*TextModerationPlusResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Score  *float32                                    `json:"Score,omitempty" xml:"Score,omitempty"`
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

func (s *TextModerationPlusResponseBodyData) SetResult(v []*TextModerationPlusResponseBodyDataResult) *TextModerationPlusResponseBodyData {
	s.Result = v
	return s
}

func (s *TextModerationPlusResponseBodyData) SetScore(v float32) *TextModerationPlusResponseBodyData {
	s.Score = &v
	return s
}

type TextModerationPlusResponseBodyDataAdvice struct {
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
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

type TextModerationPlusResponseBodyDataResult struct {
	Confidence    *float32                                                 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	CustomizedHit []*TextModerationPlusResponseBodyDataResultCustomizedHit `json:"CustomizedHit,omitempty" xml:"CustomizedHit,omitempty" type:"Repeated"`
	Label         *string                                                  `json:"Label,omitempty" xml:"Label,omitempty"`
	RiskWords     *string                                                  `json:"RiskWords,omitempty" xml:"RiskWords,omitempty"`
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

func (s *TextModerationPlusResponseBodyDataResult) SetLabel(v string) *TextModerationPlusResponseBodyDataResult {
	s.Label = &v
	return s
}

func (s *TextModerationPlusResponseBodyDataResult) SetRiskWords(v string) *TextModerationPlusResponseBodyDataResult {
	s.RiskWords = &v
	return s
}

type TextModerationPlusResponseBodyDataResultCustomizedHit struct {
	KeyWords *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	LibName  *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
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
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
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
	Code *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UrlAsyncModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg  *string                             `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Id of the request
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
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	ReqId  *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
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
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
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
	Code    *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *VideoModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s VideoModerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResponseBodyData) GoString() string {
	return s.String()
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
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
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
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
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
	Code    *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *VideoModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	AudioResult *VideoModerationResultResponseBodyDataAudioResult `json:"AudioResult,omitempty" xml:"AudioResult,omitempty" type:"Struct"`
	DataId      *string                                           `json:"DataId,omitempty" xml:"DataId,omitempty"`
	FrameResult *VideoModerationResultResponseBodyDataFrameResult `json:"FrameResult,omitempty" xml:"FrameResult,omitempty" type:"Struct"`
	LiveId      *string                                           `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
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

type VideoModerationResultResponseBodyDataAudioResult struct {
	AudioSummarys []*VideoModerationResultResponseBodyDataAudioResultAudioSummarys `json:"AudioSummarys,omitempty" xml:"AudioSummarys,omitempty" type:"Repeated"`
	SliceDetails  []*VideoModerationResultResponseBodyDataAudioResultSliceDetails  `json:"SliceDetails,omitempty" xml:"SliceDetails,omitempty" type:"Repeated"`
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

func (s *VideoModerationResultResponseBodyDataAudioResult) SetSliceDetails(v []*VideoModerationResultResponseBodyDataAudioResultSliceDetails) *VideoModerationResultResponseBodyDataAudioResult {
	s.SliceDetails = v
	return s
}

type VideoModerationResultResponseBodyDataAudioResultAudioSummarys struct {
	Label    *string `json:"Label,omitempty" xml:"Label,omitempty"`
	LabelSum *int32  `json:"LabelSum,omitempty" xml:"LabelSum,omitempty"`
}

func (s VideoModerationResultResponseBodyDataAudioResultAudioSummarys) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataAudioResultAudioSummarys) GoString() string {
	return s.String()
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
	EndTime        *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimestamp   *int64   `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	Extend         *string  `json:"Extend,omitempty" xml:"Extend,omitempty"`
	Labels         *string  `json:"Labels,omitempty" xml:"Labels,omitempty"`
	RiskTips       *string  `json:"RiskTips,omitempty" xml:"RiskTips,omitempty"`
	RiskWords      *string  `json:"RiskWords,omitempty" xml:"RiskWords,omitempty"`
	Score          *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	StartTime      *int64   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimestamp *int64   `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	Text           *string  `json:"Text,omitempty" xml:"Text,omitempty"`
	Url            *string  `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s VideoModerationResultResponseBodyDataAudioResultSliceDetails) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataAudioResultSliceDetails) GoString() string {
	return s.String()
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
	FrameNum      *int32                                                           `json:"FrameNum,omitempty" xml:"FrameNum,omitempty"`
	FrameSummarys []*VideoModerationResultResponseBodyDataFrameResultFrameSummarys `json:"FrameSummarys,omitempty" xml:"FrameSummarys,omitempty" type:"Repeated"`
	Frames        []*VideoModerationResultResponseBodyDataFrameResultFrames        `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
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

type VideoModerationResultResponseBodyDataFrameResultFrameSummarys struct {
	Label    *string `json:"Label,omitempty" xml:"Label,omitempty"`
	LabelSum *int32  `json:"LabelSum,omitempty" xml:"LabelSum,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFrameSummarys) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFrameSummarys) GoString() string {
	return s.String()
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
	Offset    *float32                                                         `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Results   []*VideoModerationResultResponseBodyDataFrameResultFramesResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	TempUrl   *string                                                          `json:"TempUrl,omitempty" xml:"TempUrl,omitempty"`
	Timestamp *int64                                                           `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
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

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) SetTempUrl(v string) *VideoModerationResultResponseBodyDataFrameResultFrames {
	s.TempUrl = &v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFrames) SetTimestamp(v int64) *VideoModerationResultResponseBodyDataFrameResultFrames {
	s.Timestamp = &v
	return s
}

type VideoModerationResultResponseBodyDataFrameResultFramesResults struct {
	Result  []*VideoModerationResultResponseBodyDataFrameResultFramesResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Service *string                                                                `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResults) String() string {
	return tea.Prettify(s)
}

func (s VideoModerationResultResponseBodyDataFrameResultFramesResults) GoString() string {
	return s.String()
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) SetResult(v []*VideoModerationResultResponseBodyDataFrameResultFramesResultsResult) *VideoModerationResultResponseBodyDataFrameResultFramesResults {
	s.Result = v
	return s
}

func (s *VideoModerationResultResponseBodyDataFrameResultFramesResults) SetService(v string) *VideoModerationResultResponseBodyDataFrameResultFramesResults {
	s.Service = &v
	return s
}

type VideoModerationResultResponseBodyDataFrameResultFramesResultsResult struct {
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Label      *string  `json:"Label,omitempty" xml:"Label,omitempty"`
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
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
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
	Code    *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *VoiceModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s VoiceModerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResponseBodyData) GoString() string {
	return s.String()
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
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
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
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
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
	Code    *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *VoiceModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	LiveId       *string                                              `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	SliceDetails []*VoiceModerationResultResponseBodyDataSliceDetails `json:"SliceDetails,omitempty" xml:"SliceDetails,omitempty" type:"Repeated"`
	TaskId       *string                                              `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Url          *string                                              `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s VoiceModerationResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *VoiceModerationResultResponseBodyData) SetLiveId(v string) *VoiceModerationResultResponseBodyData {
	s.LiveId = &v
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
	EndTime          *int64                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimestamp     *int64                 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	Extend           *string                `json:"Extend,omitempty" xml:"Extend,omitempty"`
	Labels           *string                `json:"Labels,omitempty" xml:"Labels,omitempty"`
	OriginAlgoResult map[string]interface{} `json:"OriginAlgoResult,omitempty" xml:"OriginAlgoResult,omitempty"`
	RiskTips         *string                `json:"RiskTips,omitempty" xml:"RiskTips,omitempty"`
	RiskWords        *string                `json:"RiskWords,omitempty" xml:"RiskWords,omitempty"`
	Score            *float32               `json:"Score,omitempty" xml:"Score,omitempty"`
	StartTime        *int64                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimestamp   *int64                 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	Text             *string                `json:"Text,omitempty" xml:"Text,omitempty"`
	Url              *string                `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s VoiceModerationResultResponseBodyDataSliceDetails) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResultResponseBodyDataSliceDetails) GoString() string {
	return s.String()
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
