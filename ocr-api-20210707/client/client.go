// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type DataSubImagesFigureInfoValue struct {
	// example:
	//
	// 3
	FigureCount   *int32                                       `json:"FigureCount,omitempty" xml:"FigureCount,omitempty"`
	FigureDetails []*DataSubImagesFigureInfoValueFigureDetails `json:"FigureDetails,omitempty" xml:"FigureDetails,omitempty" type:"Repeated"`
}

func (s DataSubImagesFigureInfoValue) String() string {
	return tea.Prettify(s)
}

func (s DataSubImagesFigureInfoValue) GoString() string {
	return s.String()
}

func (s *DataSubImagesFigureInfoValue) SetFigureCount(v int32) *DataSubImagesFigureInfoValue {
	s.FigureCount = &v
	return s
}

func (s *DataSubImagesFigureInfoValue) SetFigureDetails(v []*DataSubImagesFigureInfoValueFigureDetails) *DataSubImagesFigureInfoValue {
	s.FigureDetails = v
	return s
}

type DataSubImagesFigureInfoValueFigureDetails struct {
	// example:
	//
	// face
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// “”
	Data         interface{}                                              `json:"Data,omitempty" xml:"Data,omitempty"`
	FigurePoints []*DataSubImagesFigureInfoValueFigureDetailsFigurePoints `json:"FigurePoints,omitempty" xml:"FigurePoints,omitempty" type:"Repeated"`
	FigureRect   *DataSubImagesFigureInfoValueFigureDetailsFigureRect     `json:"FigureRect,omitempty" xml:"FigureRect,omitempty" type:"Struct"`
	// example:
	//
	// 0
	FigureAngle *int32 `json:"FigureAngle,omitempty" xml:"FigureAngle,omitempty"`
}

func (s DataSubImagesFigureInfoValueFigureDetails) String() string {
	return tea.Prettify(s)
}

func (s DataSubImagesFigureInfoValueFigureDetails) GoString() string {
	return s.String()
}

func (s *DataSubImagesFigureInfoValueFigureDetails) SetType(v string) *DataSubImagesFigureInfoValueFigureDetails {
	s.Type = &v
	return s
}

func (s *DataSubImagesFigureInfoValueFigureDetails) SetData(v interface{}) *DataSubImagesFigureInfoValueFigureDetails {
	s.Data = v
	return s
}

func (s *DataSubImagesFigureInfoValueFigureDetails) SetFigurePoints(v []*DataSubImagesFigureInfoValueFigureDetailsFigurePoints) *DataSubImagesFigureInfoValueFigureDetails {
	s.FigurePoints = v
	return s
}

func (s *DataSubImagesFigureInfoValueFigureDetails) SetFigureRect(v *DataSubImagesFigureInfoValueFigureDetailsFigureRect) *DataSubImagesFigureInfoValueFigureDetails {
	s.FigureRect = v
	return s
}

func (s *DataSubImagesFigureInfoValueFigureDetails) SetFigureAngle(v int32) *DataSubImagesFigureInfoValueFigureDetails {
	s.FigureAngle = &v
	return s
}

type DataSubImagesFigureInfoValueFigureDetailsFigurePoints struct {
	// example:
	//
	// 100
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 200
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DataSubImagesFigureInfoValueFigureDetailsFigurePoints) String() string {
	return tea.Prettify(s)
}

func (s DataSubImagesFigureInfoValueFigureDetailsFigurePoints) GoString() string {
	return s.String()
}

func (s *DataSubImagesFigureInfoValueFigureDetailsFigurePoints) SetX(v int32) *DataSubImagesFigureInfoValueFigureDetailsFigurePoints {
	s.X = &v
	return s
}

func (s *DataSubImagesFigureInfoValueFigureDetailsFigurePoints) SetY(v int32) *DataSubImagesFigureInfoValueFigureDetailsFigurePoints {
	s.Y = &v
	return s
}

type DataSubImagesFigureInfoValueFigureDetailsFigureRect struct {
	// example:
	//
	// 100
	CenterX *int32 `json:"CenterX,omitempty" xml:"CenterX,omitempty"`
	// example:
	//
	// 200
	CenterY *int32 `json:"CenterY,omitempty" xml:"CenterY,omitempty"`
	// example:
	//
	// 50
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 50
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s DataSubImagesFigureInfoValueFigureDetailsFigureRect) String() string {
	return tea.Prettify(s)
}

func (s DataSubImagesFigureInfoValueFigureDetailsFigureRect) GoString() string {
	return s.String()
}

func (s *DataSubImagesFigureInfoValueFigureDetailsFigureRect) SetCenterX(v int32) *DataSubImagesFigureInfoValueFigureDetailsFigureRect {
	s.CenterX = &v
	return s
}

func (s *DataSubImagesFigureInfoValueFigureDetailsFigureRect) SetCenterY(v int32) *DataSubImagesFigureInfoValueFigureDetailsFigureRect {
	s.CenterY = &v
	return s
}

func (s *DataSubImagesFigureInfoValueFigureDetailsFigureRect) SetWidth(v int32) *DataSubImagesFigureInfoValueFigureDetailsFigureRect {
	s.Width = &v
	return s
}

func (s *DataSubImagesFigureInfoValueFigureDetailsFigureRect) SetHeight(v int32) *DataSubImagesFigureInfoValueFigureDetailsFigureRect {
	s.Height = &v
	return s
}

type DataSubImagesKvInfoKvDetailsValue struct {
	// example:
	//
	// "address"
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// example:
	//
	// 100
	KeyConfidence *int32  `json:"KeyConfidence,omitempty" xml:"KeyConfidence,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// 98
	ValueConfidence *int32                                          `json:"ValueConfidence,omitempty" xml:"ValueConfidence,omitempty"`
	ValuePoints     []*DataSubImagesKvInfoKvDetailsValueValuePoints `json:"ValuePoints,omitempty" xml:"ValuePoints,omitempty" type:"Repeated"`
	ValueRect       *DataSubImagesKvInfoKvDetailsValueValueRect     `json:"ValueRect,omitempty" xml:"ValueRect,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ValueAngle *int32 `json:"ValueAngle,omitempty" xml:"ValueAngle,omitempty"`
}

func (s DataSubImagesKvInfoKvDetailsValue) String() string {
	return tea.Prettify(s)
}

func (s DataSubImagesKvInfoKvDetailsValue) GoString() string {
	return s.String()
}

func (s *DataSubImagesKvInfoKvDetailsValue) SetKeyName(v string) *DataSubImagesKvInfoKvDetailsValue {
	s.KeyName = &v
	return s
}

func (s *DataSubImagesKvInfoKvDetailsValue) SetKeyConfidence(v int32) *DataSubImagesKvInfoKvDetailsValue {
	s.KeyConfidence = &v
	return s
}

func (s *DataSubImagesKvInfoKvDetailsValue) SetValue(v string) *DataSubImagesKvInfoKvDetailsValue {
	s.Value = &v
	return s
}

func (s *DataSubImagesKvInfoKvDetailsValue) SetValueConfidence(v int32) *DataSubImagesKvInfoKvDetailsValue {
	s.ValueConfidence = &v
	return s
}

func (s *DataSubImagesKvInfoKvDetailsValue) SetValuePoints(v []*DataSubImagesKvInfoKvDetailsValueValuePoints) *DataSubImagesKvInfoKvDetailsValue {
	s.ValuePoints = v
	return s
}

func (s *DataSubImagesKvInfoKvDetailsValue) SetValueRect(v *DataSubImagesKvInfoKvDetailsValueValueRect) *DataSubImagesKvInfoKvDetailsValue {
	s.ValueRect = v
	return s
}

func (s *DataSubImagesKvInfoKvDetailsValue) SetValueAngle(v int32) *DataSubImagesKvInfoKvDetailsValue {
	s.ValueAngle = &v
	return s
}

type DataSubImagesKvInfoKvDetailsValueValuePoints struct {
	// example:
	//
	// 100
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 200
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DataSubImagesKvInfoKvDetailsValueValuePoints) String() string {
	return tea.Prettify(s)
}

func (s DataSubImagesKvInfoKvDetailsValueValuePoints) GoString() string {
	return s.String()
}

func (s *DataSubImagesKvInfoKvDetailsValueValuePoints) SetX(v int32) *DataSubImagesKvInfoKvDetailsValueValuePoints {
	s.X = &v
	return s
}

func (s *DataSubImagesKvInfoKvDetailsValueValuePoints) SetY(v int32) *DataSubImagesKvInfoKvDetailsValueValuePoints {
	s.Y = &v
	return s
}

type DataSubImagesKvInfoKvDetailsValueValueRect struct {
	// example:
	//
	// 100
	CenterX *int32 `json:"CenterX,omitempty" xml:"CenterX,omitempty"`
	// example:
	//
	// 200
	CenterY *int32 `json:"CenterY,omitempty" xml:"CenterY,omitempty"`
	// example:
	//
	// 50
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 50
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s DataSubImagesKvInfoKvDetailsValueValueRect) String() string {
	return tea.Prettify(s)
}

func (s DataSubImagesKvInfoKvDetailsValueValueRect) GoString() string {
	return s.String()
}

func (s *DataSubImagesKvInfoKvDetailsValueValueRect) SetCenterX(v int32) *DataSubImagesKvInfoKvDetailsValueValueRect {
	s.CenterX = &v
	return s
}

func (s *DataSubImagesKvInfoKvDetailsValueValueRect) SetCenterY(v int32) *DataSubImagesKvInfoKvDetailsValueValueRect {
	s.CenterY = &v
	return s
}

func (s *DataSubImagesKvInfoKvDetailsValueValueRect) SetWidth(v int32) *DataSubImagesKvInfoKvDetailsValueValueRect {
	s.Width = &v
	return s
}

func (s *DataSubImagesKvInfoKvDetailsValueValueRect) SetHeight(v int32) *DataSubImagesKvInfoKvDetailsValueValueRect {
	s.Height = &v
	return s
}

type RecognizeAdvancedRequest struct {
	// example:
	//
	// false
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// example:
	//
	// false
	NeedSortPage *bool `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
	// example:
	//
	// false
	NoStamp *bool `json:"NoStamp,omitempty" xml:"NoStamp,omitempty"`
	// example:
	//
	// false
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// example:
	//
	// false
	OutputFigure *bool `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	// example:
	//
	// false
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// example:
	//
	// false
	Paragraph *bool `json:"Paragraph,omitempty" xml:"Paragraph,omitempty"`
	// example:
	//
	// false
	Row *bool `json:"Row,omitempty" xml:"Row,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeAdvancedRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAdvancedRequest) GoString() string {
	return s.String()
}

func (s *RecognizeAdvancedRequest) SetNeedRotate(v bool) *RecognizeAdvancedRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetNeedSortPage(v bool) *RecognizeAdvancedRequest {
	s.NeedSortPage = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetNoStamp(v bool) *RecognizeAdvancedRequest {
	s.NoStamp = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetOutputCharInfo(v bool) *RecognizeAdvancedRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetOutputFigure(v bool) *RecognizeAdvancedRequest {
	s.OutputFigure = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetOutputTable(v bool) *RecognizeAdvancedRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetParagraph(v bool) *RecognizeAdvancedRequest {
	s.Paragraph = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetRow(v bool) *RecognizeAdvancedRequest {
	s.Row = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetUrl(v string) *RecognizeAdvancedRequest {
	s.Url = &v
	return s
}

func (s *RecognizeAdvancedRequest) SetBody(v io.Reader) *RecognizeAdvancedRequest {
	s.Body = v
	return s
}

type RecognizeAdvancedResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// { 	"content": "2017年河北区实验小学", 	"height": 3509, 	"orgHeight": 3509, 	"orgWidth": 2512, 	"prism_version": "1.0.9", 	"prism_wnum": 126, 	"prism_wordsInfo": [{ 		"angle": -89, 		"direction": 0, 		"height": 541, 		"pos": [{ 			"x": 982, 			"y": 223 		}, { 			"x": 1522, 			"y": 223 		}, { 			"x": 1522, 			"y": 266 		}, { 			"x": 982, 			"y": 266 		}], 		"prob": 99, 		"width": 43, 		"word": "2017年河北区实验小学", 		"x": 1230, 		"y": -26 	}], 	"width": 2512 }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeAdvancedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAdvancedResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeAdvancedResponseBody) SetCode(v string) *RecognizeAdvancedResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeAdvancedResponseBody) SetData(v string) *RecognizeAdvancedResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeAdvancedResponseBody) SetMessage(v string) *RecognizeAdvancedResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeAdvancedResponseBody) SetRequestId(v string) *RecognizeAdvancedResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeAdvancedResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeAdvancedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeAdvancedResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAdvancedResponse) GoString() string {
	return s.String()
}

func (s *RecognizeAdvancedResponse) SetHeaders(v map[string]*string) *RecognizeAdvancedResponse {
	s.Headers = v
	return s
}

func (s *RecognizeAdvancedResponse) SetStatusCode(v int32) *RecognizeAdvancedResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeAdvancedResponse) SetBody(v *RecognizeAdvancedResponseBody) *RecognizeAdvancedResponse {
	s.Body = v
	return s
}

type RecognizeAirItineraryRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1hBCIcBr0gK0jSZFnXXbRRXXa-1833-785.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeAirItineraryRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAirItineraryRequest) GoString() string {
	return s.String()
}

func (s *RecognizeAirItineraryRequest) SetUrl(v string) *RecognizeAirItineraryRequest {
	s.Url = &v
	return s
}

func (s *RecognizeAirItineraryRequest) SetBody(v io.Reader) *RecognizeAirItineraryRequest {
	s.Body = v
	return s
}

type RecognizeAirItineraryResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeAirItineraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAirItineraryResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeAirItineraryResponseBody) SetCode(v string) *RecognizeAirItineraryResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeAirItineraryResponseBody) SetData(v string) *RecognizeAirItineraryResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeAirItineraryResponseBody) SetMessage(v string) *RecognizeAirItineraryResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeAirItineraryResponseBody) SetRequestId(v string) *RecognizeAirItineraryResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeAirItineraryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeAirItineraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeAirItineraryResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAirItineraryResponse) GoString() string {
	return s.String()
}

func (s *RecognizeAirItineraryResponse) SetHeaders(v map[string]*string) *RecognizeAirItineraryResponse {
	s.Headers = v
	return s
}

func (s *RecognizeAirItineraryResponse) SetStatusCode(v int32) *RecognizeAirItineraryResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeAirItineraryResponse) SetBody(v *RecognizeAirItineraryResponseBody) *RecognizeAirItineraryResponse {
	s.Body = v
	return s
}

type RecognizeAllTextRequest struct {
	AdvancedConfig                     *RecognizeAllTextRequestAdvancedConfig                     `json:"AdvancedConfig,omitempty" xml:"AdvancedConfig,omitempty" type:"Struct"`
	IdCardConfig                       *RecognizeAllTextRequestIdCardConfig                       `json:"IdCardConfig,omitempty" xml:"IdCardConfig,omitempty" type:"Struct"`
	InternationalBusinessLicenseConfig *RecognizeAllTextRequestInternationalBusinessLicenseConfig `json:"InternationalBusinessLicenseConfig,omitempty" xml:"InternationalBusinessLicenseConfig,omitempty" type:"Struct"`
	InternationalIdCardConfig          *RecognizeAllTextRequestInternationalIdCardConfig          `json:"InternationalIdCardConfig,omitempty" xml:"InternationalIdCardConfig,omitempty" type:"Struct"`
	MultiLanConfig                     *RecognizeAllTextRequestMultiLanConfig                     `json:"MultiLanConfig,omitempty" xml:"MultiLanConfig,omitempty" type:"Struct"`
	// example:
	//
	// false
	OutputBarCode *bool `json:"OutputBarCode,omitempty" xml:"OutputBarCode,omitempty"`
	// example:
	//
	// false
	OutputCoordinate *string `json:"OutputCoordinate,omitempty" xml:"OutputCoordinate,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// false
	OutputFigure *bool `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	// example:
	//
	// false
	OutputKVExcel *bool `json:"OutputKVExcel,omitempty" xml:"OutputKVExcel,omitempty"`
	// example:
	//
	// false
	OutputOricoord *bool `json:"OutputOricoord,omitempty" xml:"OutputOricoord,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// false
	OutputQrcode *bool `json:"OutputQrcode,omitempty" xml:"OutputQrcode,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// false
	OutputStamp *bool `json:"OutputStamp,omitempty" xml:"OutputStamp,omitempty"`
	// example:
	//
	// 1
	PageNo      *int32                              `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	TableConfig *RecognizeAllTextRequestTableConfig `json:"TableConfig,omitempty" xml:"TableConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// Advanced
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://example.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeAllTextRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextRequest) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextRequest) SetAdvancedConfig(v *RecognizeAllTextRequestAdvancedConfig) *RecognizeAllTextRequest {
	s.AdvancedConfig = v
	return s
}

func (s *RecognizeAllTextRequest) SetIdCardConfig(v *RecognizeAllTextRequestIdCardConfig) *RecognizeAllTextRequest {
	s.IdCardConfig = v
	return s
}

func (s *RecognizeAllTextRequest) SetInternationalBusinessLicenseConfig(v *RecognizeAllTextRequestInternationalBusinessLicenseConfig) *RecognizeAllTextRequest {
	s.InternationalBusinessLicenseConfig = v
	return s
}

func (s *RecognizeAllTextRequest) SetInternationalIdCardConfig(v *RecognizeAllTextRequestInternationalIdCardConfig) *RecognizeAllTextRequest {
	s.InternationalIdCardConfig = v
	return s
}

func (s *RecognizeAllTextRequest) SetMultiLanConfig(v *RecognizeAllTextRequestMultiLanConfig) *RecognizeAllTextRequest {
	s.MultiLanConfig = v
	return s
}

func (s *RecognizeAllTextRequest) SetOutputBarCode(v bool) *RecognizeAllTextRequest {
	s.OutputBarCode = &v
	return s
}

func (s *RecognizeAllTextRequest) SetOutputCoordinate(v string) *RecognizeAllTextRequest {
	s.OutputCoordinate = &v
	return s
}

func (s *RecognizeAllTextRequest) SetOutputFigure(v bool) *RecognizeAllTextRequest {
	s.OutputFigure = &v
	return s
}

func (s *RecognizeAllTextRequest) SetOutputKVExcel(v bool) *RecognizeAllTextRequest {
	s.OutputKVExcel = &v
	return s
}

func (s *RecognizeAllTextRequest) SetOutputOricoord(v bool) *RecognizeAllTextRequest {
	s.OutputOricoord = &v
	return s
}

func (s *RecognizeAllTextRequest) SetOutputQrcode(v bool) *RecognizeAllTextRequest {
	s.OutputQrcode = &v
	return s
}

func (s *RecognizeAllTextRequest) SetOutputStamp(v bool) *RecognizeAllTextRequest {
	s.OutputStamp = &v
	return s
}

func (s *RecognizeAllTextRequest) SetPageNo(v int32) *RecognizeAllTextRequest {
	s.PageNo = &v
	return s
}

func (s *RecognizeAllTextRequest) SetTableConfig(v *RecognizeAllTextRequestTableConfig) *RecognizeAllTextRequest {
	s.TableConfig = v
	return s
}

func (s *RecognizeAllTextRequest) SetType(v string) *RecognizeAllTextRequest {
	s.Type = &v
	return s
}

func (s *RecognizeAllTextRequest) SetUrl(v string) *RecognizeAllTextRequest {
	s.Url = &v
	return s
}

func (s *RecognizeAllTextRequest) SetBody(v io.Reader) *RecognizeAllTextRequest {
	s.Body = v
	return s
}

type RecognizeAllTextRequestAdvancedConfig struct {
	// example:
	//
	// false
	IsHandWritingTable *bool `json:"IsHandWritingTable,omitempty" xml:"IsHandWritingTable,omitempty"`
	// example:
	//
	// false
	IsLineLessTable *bool `json:"IsLineLessTable,omitempty" xml:"IsLineLessTable,omitempty"`
	// example:
	//
	// false
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// example:
	//
	// false
	OutputParagraph *bool `json:"OutputParagraph,omitempty" xml:"OutputParagraph,omitempty"`
	// example:
	//
	// false
	OutputRow *bool `json:"OutputRow,omitempty" xml:"OutputRow,omitempty"`
	// example:
	//
	// false
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// example:
	//
	// false
	OutputTableExcel *bool `json:"OutputTableExcel,omitempty" xml:"OutputTableExcel,omitempty"`
	// example:
	//
	// false
	OutputTableHtml *bool `json:"OutputTableHtml,omitempty" xml:"OutputTableHtml,omitempty"`
}

func (s RecognizeAllTextRequestAdvancedConfig) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextRequestAdvancedConfig) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextRequestAdvancedConfig) SetIsHandWritingTable(v bool) *RecognizeAllTextRequestAdvancedConfig {
	s.IsHandWritingTable = &v
	return s
}

func (s *RecognizeAllTextRequestAdvancedConfig) SetIsLineLessTable(v bool) *RecognizeAllTextRequestAdvancedConfig {
	s.IsLineLessTable = &v
	return s
}

func (s *RecognizeAllTextRequestAdvancedConfig) SetOutputCharInfo(v bool) *RecognizeAllTextRequestAdvancedConfig {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeAllTextRequestAdvancedConfig) SetOutputParagraph(v bool) *RecognizeAllTextRequestAdvancedConfig {
	s.OutputParagraph = &v
	return s
}

func (s *RecognizeAllTextRequestAdvancedConfig) SetOutputRow(v bool) *RecognizeAllTextRequestAdvancedConfig {
	s.OutputRow = &v
	return s
}

func (s *RecognizeAllTextRequestAdvancedConfig) SetOutputTable(v bool) *RecognizeAllTextRequestAdvancedConfig {
	s.OutputTable = &v
	return s
}

func (s *RecognizeAllTextRequestAdvancedConfig) SetOutputTableExcel(v bool) *RecognizeAllTextRequestAdvancedConfig {
	s.OutputTableExcel = &v
	return s
}

func (s *RecognizeAllTextRequestAdvancedConfig) SetOutputTableHtml(v bool) *RecognizeAllTextRequestAdvancedConfig {
	s.OutputTableHtml = &v
	return s
}

type RecognizeAllTextRequestIdCardConfig struct {
	// example:
	//
	// false
	OutputIdCardQuality *bool `json:"OutputIdCardQuality,omitempty" xml:"OutputIdCardQuality,omitempty"`
}

func (s RecognizeAllTextRequestIdCardConfig) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextRequestIdCardConfig) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextRequestIdCardConfig) SetOutputIdCardQuality(v bool) *RecognizeAllTextRequestIdCardConfig {
	s.OutputIdCardQuality = &v
	return s
}

type RecognizeAllTextRequestInternationalBusinessLicenseConfig struct {
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
}

func (s RecognizeAllTextRequestInternationalBusinessLicenseConfig) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextRequestInternationalBusinessLicenseConfig) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextRequestInternationalBusinessLicenseConfig) SetCountry(v string) *RecognizeAllTextRequestInternationalBusinessLicenseConfig {
	s.Country = &v
	return s
}

type RecognizeAllTextRequestInternationalIdCardConfig struct {
	// example:
	//
	// India
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
}

func (s RecognizeAllTextRequestInternationalIdCardConfig) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextRequestInternationalIdCardConfig) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextRequestInternationalIdCardConfig) SetCountry(v string) *RecognizeAllTextRequestInternationalIdCardConfig {
	s.Country = &v
	return s
}

type RecognizeAllTextRequestMultiLanConfig struct {
	// example:
	//
	// eng,chn
	Languages *string `json:"Languages,omitempty" xml:"Languages,omitempty"`
}

func (s RecognizeAllTextRequestMultiLanConfig) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextRequestMultiLanConfig) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextRequestMultiLanConfig) SetLanguages(v string) *RecognizeAllTextRequestMultiLanConfig {
	s.Languages = &v
	return s
}

type RecognizeAllTextRequestTableConfig struct {
	IsHandWritingTable *bool `json:"IsHandWritingTable,omitempty" xml:"IsHandWritingTable,omitempty"`
	IsLineLessTable    *bool `json:"IsLineLessTable,omitempty" xml:"IsLineLessTable,omitempty"`
	OutputTableExcel   *bool `json:"OutputTableExcel,omitempty" xml:"OutputTableExcel,omitempty"`
	OutputTableHtml    *bool `json:"OutputTableHtml,omitempty" xml:"OutputTableHtml,omitempty"`
}

func (s RecognizeAllTextRequestTableConfig) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextRequestTableConfig) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextRequestTableConfig) SetIsHandWritingTable(v bool) *RecognizeAllTextRequestTableConfig {
	s.IsHandWritingTable = &v
	return s
}

func (s *RecognizeAllTextRequestTableConfig) SetIsLineLessTable(v bool) *RecognizeAllTextRequestTableConfig {
	s.IsLineLessTable = &v
	return s
}

func (s *RecognizeAllTextRequestTableConfig) SetOutputTableExcel(v bool) *RecognizeAllTextRequestTableConfig {
	s.OutputTableExcel = &v
	return s
}

func (s *RecognizeAllTextRequestTableConfig) SetOutputTableHtml(v bool) *RecognizeAllTextRequestTableConfig {
	s.OutputTableHtml = &v
	return s
}

type RecognizeAllTextShrinkRequest struct {
	AdvancedConfigShrink                     *string `json:"AdvancedConfig,omitempty" xml:"AdvancedConfig,omitempty"`
	IdCardConfigShrink                       *string `json:"IdCardConfig,omitempty" xml:"IdCardConfig,omitempty"`
	InternationalBusinessLicenseConfigShrink *string `json:"InternationalBusinessLicenseConfig,omitempty" xml:"InternationalBusinessLicenseConfig,omitempty"`
	InternationalIdCardConfigShrink          *string `json:"InternationalIdCardConfig,omitempty" xml:"InternationalIdCardConfig,omitempty"`
	MultiLanConfigShrink                     *string `json:"MultiLanConfig,omitempty" xml:"MultiLanConfig,omitempty"`
	// example:
	//
	// false
	OutputBarCode *bool `json:"OutputBarCode,omitempty" xml:"OutputBarCode,omitempty"`
	// example:
	//
	// false
	OutputCoordinate *string `json:"OutputCoordinate,omitempty" xml:"OutputCoordinate,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// false
	OutputFigure *bool `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	// example:
	//
	// false
	OutputKVExcel *bool `json:"OutputKVExcel,omitempty" xml:"OutputKVExcel,omitempty"`
	// example:
	//
	// false
	OutputOricoord *bool `json:"OutputOricoord,omitempty" xml:"OutputOricoord,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// false
	OutputQrcode *bool `json:"OutputQrcode,omitempty" xml:"OutputQrcode,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// false
	OutputStamp *bool `json:"OutputStamp,omitempty" xml:"OutputStamp,omitempty"`
	// example:
	//
	// 1
	PageNo            *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	TableConfigShrink *string `json:"TableConfig,omitempty" xml:"TableConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Advanced
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://example.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeAllTextShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextShrinkRequest) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextShrinkRequest) SetAdvancedConfigShrink(v string) *RecognizeAllTextShrinkRequest {
	s.AdvancedConfigShrink = &v
	return s
}

func (s *RecognizeAllTextShrinkRequest) SetIdCardConfigShrink(v string) *RecognizeAllTextShrinkRequest {
	s.IdCardConfigShrink = &v
	return s
}

func (s *RecognizeAllTextShrinkRequest) SetInternationalBusinessLicenseConfigShrink(v string) *RecognizeAllTextShrinkRequest {
	s.InternationalBusinessLicenseConfigShrink = &v
	return s
}

func (s *RecognizeAllTextShrinkRequest) SetInternationalIdCardConfigShrink(v string) *RecognizeAllTextShrinkRequest {
	s.InternationalIdCardConfigShrink = &v
	return s
}

func (s *RecognizeAllTextShrinkRequest) SetMultiLanConfigShrink(v string) *RecognizeAllTextShrinkRequest {
	s.MultiLanConfigShrink = &v
	return s
}

func (s *RecognizeAllTextShrinkRequest) SetOutputBarCode(v bool) *RecognizeAllTextShrinkRequest {
	s.OutputBarCode = &v
	return s
}

func (s *RecognizeAllTextShrinkRequest) SetOutputCoordinate(v string) *RecognizeAllTextShrinkRequest {
	s.OutputCoordinate = &v
	return s
}

func (s *RecognizeAllTextShrinkRequest) SetOutputFigure(v bool) *RecognizeAllTextShrinkRequest {
	s.OutputFigure = &v
	return s
}

func (s *RecognizeAllTextShrinkRequest) SetOutputKVExcel(v bool) *RecognizeAllTextShrinkRequest {
	s.OutputKVExcel = &v
	return s
}

func (s *RecognizeAllTextShrinkRequest) SetOutputOricoord(v bool) *RecognizeAllTextShrinkRequest {
	s.OutputOricoord = &v
	return s
}

func (s *RecognizeAllTextShrinkRequest) SetOutputQrcode(v bool) *RecognizeAllTextShrinkRequest {
	s.OutputQrcode = &v
	return s
}

func (s *RecognizeAllTextShrinkRequest) SetOutputStamp(v bool) *RecognizeAllTextShrinkRequest {
	s.OutputStamp = &v
	return s
}

func (s *RecognizeAllTextShrinkRequest) SetPageNo(v int32) *RecognizeAllTextShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *RecognizeAllTextShrinkRequest) SetTableConfigShrink(v string) *RecognizeAllTextShrinkRequest {
	s.TableConfigShrink = &v
	return s
}

func (s *RecognizeAllTextShrinkRequest) SetType(v string) *RecognizeAllTextShrinkRequest {
	s.Type = &v
	return s
}

func (s *RecognizeAllTextShrinkRequest) SetUrl(v string) *RecognizeAllTextShrinkRequest {
	s.Url = &v
	return s
}

func (s *RecognizeAllTextShrinkRequest) SetBody(v io.Reader) *RecognizeAllTextShrinkRequest {
	s.Body = v
	return s
}

type RecognizeAllTextResponseBody struct {
	// example:
	//
	// 400
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *RecognizeAllTextResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// illegalImageUrl
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E2A98925-DC2C-18FB-995F-BAF507XXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeAllTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBody) SetCode(v string) *RecognizeAllTextResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeAllTextResponseBody) SetData(v *RecognizeAllTextResponseBodyData) *RecognizeAllTextResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeAllTextResponseBody) SetMessage(v string) *RecognizeAllTextResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeAllTextResponseBody) SetRequestId(v string) *RecognizeAllTextResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeAllTextResponseBodyData struct {
	AlgoServer []*string `json:"AlgoServer,omitempty" xml:"AlgoServer,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	AlgoVersion *string `json:"AlgoVersion,omitempty" xml:"AlgoVersion,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// ""
	DebugInfo interface{} `json:"DebugInfo,omitempty" xml:"DebugInfo,omitempty"`
	// example:
	//
	// 2000
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// false
	IsMixedMode *bool `json:"IsMixedMode,omitempty" xml:"IsMixedMode,omitempty"`
	// example:
	//
	// https://example.xlsx
	KvExcelUrl *string `json:"KvExcelUrl,omitempty" xml:"KvExcelUrl,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 2
	SubImageCount *int32                                       `json:"SubImageCount,omitempty" xml:"SubImageCount,omitempty"`
	SubImages     []*RecognizeAllTextResponseBodyDataSubImages `json:"SubImages,omitempty" xml:"SubImages,omitempty" type:"Repeated"`
	// example:
	//
	// 1000
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// ""
	XmlResult *string `json:"XmlResult,omitempty" xml:"XmlResult,omitempty"`
}

func (s RecognizeAllTextResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyData) SetAlgoServer(v []*string) *RecognizeAllTextResponseBodyData {
	s.AlgoServer = v
	return s
}

func (s *RecognizeAllTextResponseBodyData) SetAlgoVersion(v string) *RecognizeAllTextResponseBodyData {
	s.AlgoVersion = &v
	return s
}

func (s *RecognizeAllTextResponseBodyData) SetContent(v string) *RecognizeAllTextResponseBodyData {
	s.Content = &v
	return s
}

func (s *RecognizeAllTextResponseBodyData) SetDebugInfo(v interface{}) *RecognizeAllTextResponseBodyData {
	s.DebugInfo = v
	return s
}

func (s *RecognizeAllTextResponseBodyData) SetHeight(v int32) *RecognizeAllTextResponseBodyData {
	s.Height = &v
	return s
}

func (s *RecognizeAllTextResponseBodyData) SetIsMixedMode(v bool) *RecognizeAllTextResponseBodyData {
	s.IsMixedMode = &v
	return s
}

func (s *RecognizeAllTextResponseBodyData) SetKvExcelUrl(v string) *RecognizeAllTextResponseBodyData {
	s.KvExcelUrl = &v
	return s
}

func (s *RecognizeAllTextResponseBodyData) SetPageNo(v int32) *RecognizeAllTextResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *RecognizeAllTextResponseBodyData) SetSubImageCount(v int32) *RecognizeAllTextResponseBodyData {
	s.SubImageCount = &v
	return s
}

func (s *RecognizeAllTextResponseBodyData) SetSubImages(v []*RecognizeAllTextResponseBodyDataSubImages) *RecognizeAllTextResponseBodyData {
	s.SubImages = v
	return s
}

func (s *RecognizeAllTextResponseBodyData) SetWidth(v int32) *RecognizeAllTextResponseBodyData {
	s.Width = &v
	return s
}

func (s *RecognizeAllTextResponseBodyData) SetXmlResult(v string) *RecognizeAllTextResponseBodyData {
	s.XmlResult = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImages struct {
	// example:
	//
	// 0
	Angle         *int32                                                  `json:"Angle,omitempty" xml:"Angle,omitempty"`
	BarCodeInfo   *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfo   `json:"BarCodeInfo,omitempty" xml:"BarCodeInfo,omitempty" type:"Struct"`
	BlockInfo     *RecognizeAllTextResponseBodyDataSubImagesBlockInfo     `json:"BlockInfo,omitempty" xml:"BlockInfo,omitempty" type:"Struct"`
	FigureInfo    map[string]*DataSubImagesFigureInfoValue                `json:"FigureInfo,omitempty" xml:"FigureInfo,omitempty"`
	KvInfo        *RecognizeAllTextResponseBodyDataSubImagesKvInfo        `json:"KvInfo,omitempty" xml:"KvInfo,omitempty" type:"Struct"`
	ParagraphInfo *RecognizeAllTextResponseBodyDataSubImagesParagraphInfo `json:"ParagraphInfo,omitempty" xml:"ParagraphInfo,omitempty" type:"Struct"`
	QrCodeInfo    *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfo    `json:"QrCodeInfo,omitempty" xml:"QrCodeInfo,omitempty" type:"Struct"`
	QualityInfo   *RecognizeAllTextResponseBodyDataSubImagesQualityInfo   `json:"QualityInfo,omitempty" xml:"QualityInfo,omitempty" type:"Struct"`
	RowInfo       *RecognizeAllTextResponseBodyDataSubImagesRowInfo       `json:"RowInfo,omitempty" xml:"RowInfo,omitempty" type:"Struct"`
	StampInfo     *RecognizeAllTextResponseBodyDataSubImagesStampInfo     `json:"StampInfo,omitempty" xml:"StampInfo,omitempty" type:"Struct"`
	// example:
	//
	// 0
	SubImageId     *int32                                                     `json:"SubImageId,omitempty" xml:"SubImageId,omitempty"`
	SubImagePoints []*RecognizeAllTextResponseBodyDataSubImagesSubImagePoints `json:"SubImagePoints,omitempty" xml:"SubImagePoints,omitempty" type:"Repeated"`
	SubImageRect   *RecognizeAllTextResponseBodyDataSubImagesSubImageRect     `json:"SubImageRect,omitempty" xml:"SubImageRect,omitempty" type:"Struct"`
	TableInfo      *RecognizeAllTextResponseBodyDataSubImagesTableInfo        `json:"TableInfo,omitempty" xml:"TableInfo,omitempty" type:"Struct"`
	Type           *string                                                    `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImages) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImages) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImages) SetAngle(v int32) *RecognizeAllTextResponseBodyDataSubImages {
	s.Angle = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImages) SetBarCodeInfo(v *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfo) *RecognizeAllTextResponseBodyDataSubImages {
	s.BarCodeInfo = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImages) SetBlockInfo(v *RecognizeAllTextResponseBodyDataSubImagesBlockInfo) *RecognizeAllTextResponseBodyDataSubImages {
	s.BlockInfo = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImages) SetFigureInfo(v map[string]*DataSubImagesFigureInfoValue) *RecognizeAllTextResponseBodyDataSubImages {
	s.FigureInfo = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImages) SetKvInfo(v *RecognizeAllTextResponseBodyDataSubImagesKvInfo) *RecognizeAllTextResponseBodyDataSubImages {
	s.KvInfo = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImages) SetParagraphInfo(v *RecognizeAllTextResponseBodyDataSubImagesParagraphInfo) *RecognizeAllTextResponseBodyDataSubImages {
	s.ParagraphInfo = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImages) SetQrCodeInfo(v *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfo) *RecognizeAllTextResponseBodyDataSubImages {
	s.QrCodeInfo = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImages) SetQualityInfo(v *RecognizeAllTextResponseBodyDataSubImagesQualityInfo) *RecognizeAllTextResponseBodyDataSubImages {
	s.QualityInfo = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImages) SetRowInfo(v *RecognizeAllTextResponseBodyDataSubImagesRowInfo) *RecognizeAllTextResponseBodyDataSubImages {
	s.RowInfo = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImages) SetStampInfo(v *RecognizeAllTextResponseBodyDataSubImagesStampInfo) *RecognizeAllTextResponseBodyDataSubImages {
	s.StampInfo = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImages) SetSubImageId(v int32) *RecognizeAllTextResponseBodyDataSubImages {
	s.SubImageId = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImages) SetSubImagePoints(v []*RecognizeAllTextResponseBodyDataSubImagesSubImagePoints) *RecognizeAllTextResponseBodyDataSubImages {
	s.SubImagePoints = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImages) SetSubImageRect(v *RecognizeAllTextResponseBodyDataSubImagesSubImageRect) *RecognizeAllTextResponseBodyDataSubImages {
	s.SubImageRect = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImages) SetTableInfo(v *RecognizeAllTextResponseBodyDataSubImagesTableInfo) *RecognizeAllTextResponseBodyDataSubImages {
	s.TableInfo = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImages) SetType(v string) *RecognizeAllTextResponseBodyDataSubImages {
	s.Type = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesBarCodeInfo struct {
	// example:
	//
	// 2
	BarCodeCount   *int32                                                                `json:"BarCodeCount,omitempty" xml:"BarCodeCount,omitempty"`
	BarCodeDetails []*RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetails `json:"BarCodeDetails,omitempty" xml:"BarCodeDetails,omitempty" type:"Repeated"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesBarCodeInfo) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesBarCodeInfo) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfo) SetBarCodeCount(v int32) *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfo {
	s.BarCodeCount = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfo) SetBarCodeDetails(v []*RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetails) *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfo {
	s.BarCodeDetails = v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetails struct {
	// example:
	//
	// 0
	BarCodeAngle  *int32                                                                             `json:"BarCodeAngle,omitempty" xml:"BarCodeAngle,omitempty"`
	BarCodePoints []*RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodePoints `json:"BarCodePoints,omitempty" xml:"BarCodePoints,omitempty" type:"Repeated"`
	BarCodeRect   *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodeRect     `json:"BarCodeRect,omitempty" xml:"BarCodeRect,omitempty" type:"Struct"`
	// example:
	//
	// "1100011XXXXXX"
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Code128
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetails) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetails) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetails) SetBarCodeAngle(v int32) *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetails {
	s.BarCodeAngle = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetails) SetBarCodePoints(v []*RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodePoints) *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetails {
	s.BarCodePoints = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetails) SetBarCodeRect(v *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodeRect) *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetails {
	s.BarCodeRect = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetails) SetData(v interface{}) *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetails {
	s.Data = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetails) SetType(v string) *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetails {
	s.Type = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodePoints struct {
	// example:
	//
	// 100
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 200
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodePoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodePoints) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodePoints) SetX(v int32) *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodePoints {
	s.X = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodePoints) SetY(v int32) *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodePoints {
	s.Y = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodeRect struct {
	// example:
	//
	// 100
	CenterX *int32 `json:"CenterX,omitempty" xml:"CenterX,omitempty"`
	// example:
	//
	// 200
	CenterY *int32 `json:"CenterY,omitempty" xml:"CenterY,omitempty"`
	// example:
	//
	// 10
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 100
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodeRect) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodeRect) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodeRect) SetCenterX(v int32) *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodeRect {
	s.CenterX = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodeRect) SetCenterY(v int32) *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodeRect {
	s.CenterY = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodeRect) SetHeight(v int32) *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodeRect {
	s.Height = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodeRect) SetWidth(v int32) *RecognizeAllTextResponseBodyDataSubImagesBarCodeInfoBarCodeDetailsBarCodeRect {
	s.Width = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesBlockInfo struct {
	// example:
	//
	// 12
	BlockCount   *int32                                                            `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
	BlockDetails []*RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails `json:"BlockDetails,omitempty" xml:"BlockDetails,omitempty" type:"Repeated"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesBlockInfo) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesBlockInfo) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfo) SetBlockCount(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfo {
	s.BlockCount = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfo) SetBlockDetails(v []*RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails) *RecognizeAllTextResponseBodyDataSubImagesBlockInfo {
	s.BlockDetails = v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails struct {
	// example:
	//
	// 0
	BlockAngle *int32 `json:"BlockAngle,omitempty" xml:"BlockAngle,omitempty"`
	// example:
	//
	// 98
	BlockConfidence *int32  `json:"BlockConfidence,omitempty" xml:"BlockConfidence,omitempty"`
	BlockContent    *string `json:"BlockContent,omitempty" xml:"BlockContent,omitempty"`
	// example:
	//
	// 0
	BlockId     *int32                                                                       `json:"BlockId,omitempty" xml:"BlockId,omitempty"`
	BlockPoints []*RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockPoints `json:"BlockPoints,omitempty" xml:"BlockPoints,omitempty" type:"Repeated"`
	BlockRect   *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockRect     `json:"BlockRect,omitempty" xml:"BlockRect,omitempty" type:"Struct"`
	CharInfos   []*RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfos   `json:"CharInfos,omitempty" xml:"CharInfos,omitempty" type:"Repeated"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails) SetBlockAngle(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails {
	s.BlockAngle = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails) SetBlockConfidence(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails {
	s.BlockConfidence = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails) SetBlockContent(v string) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails {
	s.BlockContent = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails) SetBlockId(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails {
	s.BlockId = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails) SetBlockPoints(v []*RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockPoints) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails {
	s.BlockPoints = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails) SetBlockRect(v *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockRect) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails {
	s.BlockRect = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails) SetCharInfos(v []*RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfos) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetails {
	s.CharInfos = v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockPoints struct {
	// example:
	//
	// 100
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 200
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockPoints) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockPoints) SetX(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockPoints {
	s.X = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockPoints) SetY(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockPoints {
	s.Y = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockRect struct {
	// example:
	//
	// 100
	CenterX *int32 `json:"CenterX,omitempty" xml:"CenterX,omitempty"`
	// example:
	//
	// 200
	CenterY *int32 `json:"CenterY,omitempty" xml:"CenterY,omitempty"`
	// example:
	//
	// 10
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 50
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockRect) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockRect) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockRect) SetCenterX(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockRect {
	s.CenterX = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockRect) SetCenterY(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockRect {
	s.CenterY = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockRect) SetHeight(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockRect {
	s.Height = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockRect) SetWidth(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsBlockRect {
	s.Width = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfos struct {
	// example:
	//
	// 95
	CharConfidence *int32  `json:"CharConfidence,omitempty" xml:"CharConfidence,omitempty"`
	CharContent    *string `json:"CharContent,omitempty" xml:"CharContent,omitempty"`
	// example:
	//
	// 0
	CharId     *int32                                                                               `json:"CharId,omitempty" xml:"CharId,omitempty"`
	CharPoints []*RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharPoints `json:"CharPoints,omitempty" xml:"CharPoints,omitempty" type:"Repeated"`
	CharRect   *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharRect     `json:"CharRect,omitempty" xml:"CharRect,omitempty" type:"Struct"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfos) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfos) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfos) SetCharConfidence(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfos {
	s.CharConfidence = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfos) SetCharContent(v string) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfos {
	s.CharContent = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfos) SetCharId(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfos {
	s.CharId = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfos) SetCharPoints(v []*RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharPoints) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfos {
	s.CharPoints = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfos) SetCharRect(v *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharRect) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfos {
	s.CharRect = v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharPoints struct {
	// example:
	//
	// 100
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 200
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharPoints) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharPoints) SetX(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharPoints {
	s.X = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharPoints) SetY(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharPoints {
	s.Y = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharRect struct {
	// example:
	//
	// 100
	CenterX *int32 `json:"CenterX,omitempty" xml:"CenterX,omitempty"`
	// example:
	//
	// 200
	CenterY *int32 `json:"CenterY,omitempty" xml:"CenterY,omitempty"`
	// example:
	//
	// 10
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 10
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharRect) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharRect) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharRect) SetCenterX(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharRect {
	s.CenterX = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharRect) SetCenterY(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharRect {
	s.CenterY = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharRect) SetHeight(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharRect {
	s.Height = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharRect) SetWidth(v int32) *RecognizeAllTextResponseBodyDataSubImagesBlockInfoBlockDetailsCharInfosCharRect {
	s.Width = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesKvInfo struct {
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 6
	KvCount   *int32                                        `json:"KvCount,omitempty" xml:"KvCount,omitempty"`
	KvDetails map[string]*DataSubImagesKvInfoKvDetailsValue `json:"KvDetails,omitempty" xml:"KvDetails,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesKvInfo) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesKvInfo) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesKvInfo) SetData(v interface{}) *RecognizeAllTextResponseBodyDataSubImagesKvInfo {
	s.Data = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesKvInfo) SetKvCount(v int32) *RecognizeAllTextResponseBodyDataSubImagesKvInfo {
	s.KvCount = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesKvInfo) SetKvDetails(v map[string]*DataSubImagesKvInfoKvDetailsValue) *RecognizeAllTextResponseBodyDataSubImagesKvInfo {
	s.KvDetails = v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesParagraphInfo struct {
	// example:
	//
	// 11
	ParagraphCount   *int32                                                                    `json:"ParagraphCount,omitempty" xml:"ParagraphCount,omitempty"`
	ParagraphDetails []*RecognizeAllTextResponseBodyDataSubImagesParagraphInfoParagraphDetails `json:"ParagraphDetails,omitempty" xml:"ParagraphDetails,omitempty" type:"Repeated"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesParagraphInfo) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesParagraphInfo) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesParagraphInfo) SetParagraphCount(v int32) *RecognizeAllTextResponseBodyDataSubImagesParagraphInfo {
	s.ParagraphCount = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesParagraphInfo) SetParagraphDetails(v []*RecognizeAllTextResponseBodyDataSubImagesParagraphInfoParagraphDetails) *RecognizeAllTextResponseBodyDataSubImagesParagraphInfo {
	s.ParagraphDetails = v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesParagraphInfoParagraphDetails struct {
	BlockList        []*int32 `json:"BlockList,omitempty" xml:"BlockList,omitempty" type:"Repeated"`
	ParagraphContent *string  `json:"ParagraphContent,omitempty" xml:"ParagraphContent,omitempty"`
	// example:
	//
	// 0
	ParagraphId *int32 `json:"ParagraphId,omitempty" xml:"ParagraphId,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesParagraphInfoParagraphDetails) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesParagraphInfoParagraphDetails) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesParagraphInfoParagraphDetails) SetBlockList(v []*int32) *RecognizeAllTextResponseBodyDataSubImagesParagraphInfoParagraphDetails {
	s.BlockList = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesParagraphInfoParagraphDetails) SetParagraphContent(v string) *RecognizeAllTextResponseBodyDataSubImagesParagraphInfoParagraphDetails {
	s.ParagraphContent = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesParagraphInfoParagraphDetails) SetParagraphId(v int32) *RecognizeAllTextResponseBodyDataSubImagesParagraphInfoParagraphDetails {
	s.ParagraphId = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesQrCodeInfo struct {
	// example:
	//
	// 1
	QrCodeCount   *int32                                                              `json:"QrCodeCount,omitempty" xml:"QrCodeCount,omitempty"`
	QrCodeDetails []*RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetails `json:"QrCodeDetails,omitempty" xml:"QrCodeDetails,omitempty" type:"Repeated"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesQrCodeInfo) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesQrCodeInfo) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfo) SetQrCodeCount(v int32) *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfo {
	s.QrCodeCount = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfo) SetQrCodeDetails(v []*RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetails) *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfo {
	s.QrCodeDetails = v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetails struct {
	// example:
	//
	// “http://www.gsxt.gov.cn/indeXXX”
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 0
	QrCodeAngle  *int32                                                                          `json:"QrCodeAngle,omitempty" xml:"QrCodeAngle,omitempty"`
	QrCodePoints []*RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodePoints `json:"QrCodePoints,omitempty" xml:"QrCodePoints,omitempty" type:"Repeated"`
	QrCodeRect   *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodeRect     `json:"QrCodeRect,omitempty" xml:"QrCodeRect,omitempty" type:"Struct"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetails) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetails) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetails) SetData(v interface{}) *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetails {
	s.Data = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetails) SetQrCodeAngle(v int32) *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetails {
	s.QrCodeAngle = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetails) SetQrCodePoints(v []*RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodePoints) *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetails {
	s.QrCodePoints = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetails) SetQrCodeRect(v *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodeRect) *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetails {
	s.QrCodeRect = v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodePoints struct {
	// example:
	//
	// 100
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 200
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodePoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodePoints) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodePoints) SetX(v int32) *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodePoints {
	s.X = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodePoints) SetY(v int32) *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodePoints {
	s.Y = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodeRect struct {
	// example:
	//
	// 100
	CenterX *int32 `json:"CenterX,omitempty" xml:"CenterX,omitempty"`
	// example:
	//
	// 200
	CenterY *int32 `json:"CenterY,omitempty" xml:"CenterY,omitempty"`
	// example:
	//
	// 100
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 100
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodeRect) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodeRect) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodeRect) SetCenterX(v int32) *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodeRect {
	s.CenterX = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodeRect) SetCenterY(v int32) *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodeRect {
	s.CenterY = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodeRect) SetHeight(v int32) *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodeRect {
	s.Height = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodeRect) SetWidth(v int32) *RecognizeAllTextResponseBodyDataSubImagesQrCodeInfoQrCodeDetailsQrCodeRect {
	s.Width = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesQualityInfo struct {
	// example:
	//
	// 90.5
	CompletenessScore *float32 `json:"CompletenessScore,omitempty" xml:"CompletenessScore,omitempty"`
	// example:
	//
	// false
	IsCopy *bool `json:"IsCopy,omitempty" xml:"IsCopy,omitempty"`
	// example:
	//
	// false
	IsReshoot *bool `json:"IsReshoot,omitempty" xml:"IsReshoot,omitempty"`
	// example:
	//
	// 80.5
	QualityScore *float32 `json:"QualityScore,omitempty" xml:"QualityScore,omitempty"`
	// example:
	//
	// 10.5
	TamperScore *float32 `json:"TamperScore,omitempty" xml:"TamperScore,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesQualityInfo) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesQualityInfo) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQualityInfo) SetCompletenessScore(v float32) *RecognizeAllTextResponseBodyDataSubImagesQualityInfo {
	s.CompletenessScore = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQualityInfo) SetIsCopy(v bool) *RecognizeAllTextResponseBodyDataSubImagesQualityInfo {
	s.IsCopy = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQualityInfo) SetIsReshoot(v bool) *RecognizeAllTextResponseBodyDataSubImagesQualityInfo {
	s.IsReshoot = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQualityInfo) SetQualityScore(v float32) *RecognizeAllTextResponseBodyDataSubImagesQualityInfo {
	s.QualityScore = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesQualityInfo) SetTamperScore(v float32) *RecognizeAllTextResponseBodyDataSubImagesQualityInfo {
	s.TamperScore = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesRowInfo struct {
	// example:
	//
	// 9
	RowCount   *int32                                                        `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	RowDetails []*RecognizeAllTextResponseBodyDataSubImagesRowInfoRowDetails `json:"RowDetails,omitempty" xml:"RowDetails,omitempty" type:"Repeated"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesRowInfo) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesRowInfo) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesRowInfo) SetRowCount(v int32) *RecognizeAllTextResponseBodyDataSubImagesRowInfo {
	s.RowCount = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesRowInfo) SetRowDetails(v []*RecognizeAllTextResponseBodyDataSubImagesRowInfoRowDetails) *RecognizeAllTextResponseBodyDataSubImagesRowInfo {
	s.RowDetails = v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesRowInfoRowDetails struct {
	BlockList  []*int32 `json:"BlockList,omitempty" xml:"BlockList,omitempty" type:"Repeated"`
	RowContent *string  `json:"RowContent,omitempty" xml:"RowContent,omitempty"`
	// example:
	//
	// 0
	RowId *int32 `json:"RowId,omitempty" xml:"RowId,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesRowInfoRowDetails) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesRowInfoRowDetails) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesRowInfoRowDetails) SetBlockList(v []*int32) *RecognizeAllTextResponseBodyDataSubImagesRowInfoRowDetails {
	s.BlockList = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesRowInfoRowDetails) SetRowContent(v string) *RecognizeAllTextResponseBodyDataSubImagesRowInfoRowDetails {
	s.RowContent = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesRowInfoRowDetails) SetRowId(v int32) *RecognizeAllTextResponseBodyDataSubImagesRowInfoRowDetails {
	s.RowId = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesStampInfo struct {
	// example:
	//
	// 2
	StampCount   *int32                                                            `json:"StampCount,omitempty" xml:"StampCount,omitempty"`
	StampDetails []*RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetails `json:"StampDetails,omitempty" xml:"StampDetails,omitempty" type:"Repeated"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesStampInfo) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesStampInfo) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfo) SetStampCount(v int32) *RecognizeAllTextResponseBodyDataSubImagesStampInfo {
	s.StampCount = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfo) SetStampDetails(v []*RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetails) *RecognizeAllTextResponseBodyDataSubImagesStampInfo {
	s.StampDetails = v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetails struct {
	Data *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	StampAngle  *int32                                                                       `json:"StampAngle,omitempty" xml:"StampAngle,omitempty"`
	StampPoints []*RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampPoints `json:"StampPoints,omitempty" xml:"StampPoints,omitempty" type:"Repeated"`
	StampRect   *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampRect     `json:"StampRect,omitempty" xml:"StampRect,omitempty" type:"Struct"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetails) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetails) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetails) SetData(v *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetails {
	s.Data = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetails) SetStampAngle(v int32) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetails {
	s.StampAngle = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetails) SetStampPoints(v []*RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampPoints) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetails {
	s.StampPoints = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetails) SetStampRect(v *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampRect) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetails {
	s.StampRect = v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData struct {
	// example:
	//
	// "3205823XXXXXX"
	AntiFakeCode *string `json:"AntiFakeCode,omitempty" xml:"AntiFakeCode,omitempty"`
	// example:
	//
	// "XXX"
	CompanyId        *string `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty"`
	// example:
	//
	// ""
	OrganizationNameEng *string `json:"OrganizationNameEng,omitempty" xml:"OrganizationNameEng,omitempty"`
	// example:
	//
	// "3205823XXXXXX"
	OtherText *string `json:"OtherText,omitempty" xml:"OtherText,omitempty"`
	// example:
	//
	// ""
	TaxpayerId *string `json:"TaxpayerId,omitempty" xml:"TaxpayerId,omitempty"`
	TopText    *string `json:"TopText,omitempty" xml:"TopText,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData) SetAntiFakeCode(v string) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData {
	s.AntiFakeCode = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData) SetCompanyId(v string) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData {
	s.CompanyId = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData) SetOrganizationName(v string) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData {
	s.OrganizationName = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData) SetOrganizationNameEng(v string) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData {
	s.OrganizationNameEng = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData) SetOtherText(v string) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData {
	s.OtherText = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData) SetTaxpayerId(v string) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData {
	s.TaxpayerId = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData) SetTopText(v string) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsData {
	s.TopText = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampPoints struct {
	// example:
	//
	// 100
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 200
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampPoints) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampPoints) SetX(v int32) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampPoints {
	s.X = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampPoints) SetY(v int32) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampPoints {
	s.Y = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampRect struct {
	// example:
	//
	// 100
	CenterX *int32 `json:"CenterX,omitempty" xml:"CenterX,omitempty"`
	// example:
	//
	// 200
	CenterY *int32 `json:"CenterY,omitempty" xml:"CenterY,omitempty"`
	// example:
	//
	// 50
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 50
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampRect) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampRect) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampRect) SetCenterX(v int32) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampRect {
	s.CenterX = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampRect) SetCenterY(v int32) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampRect {
	s.CenterY = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampRect) SetHeight(v int32) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampRect {
	s.Height = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampRect) SetWidth(v int32) *RecognizeAllTextResponseBodyDataSubImagesStampInfoStampDetailsStampRect {
	s.Width = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesSubImagePoints struct {
	// example:
	//
	// 100
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 200
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesSubImagePoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesSubImagePoints) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesSubImagePoints) SetX(v int32) *RecognizeAllTextResponseBodyDataSubImagesSubImagePoints {
	s.X = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesSubImagePoints) SetY(v int32) *RecognizeAllTextResponseBodyDataSubImagesSubImagePoints {
	s.Y = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesSubImageRect struct {
	// example:
	//
	// 100
	CenterX *int32 `json:"CenterX,omitempty" xml:"CenterX,omitempty"`
	// example:
	//
	// 200
	CenterY *int32 `json:"CenterY,omitempty" xml:"CenterY,omitempty"`
	// example:
	//
	// 2000
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 1000
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesSubImageRect) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesSubImageRect) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesSubImageRect) SetCenterX(v int32) *RecognizeAllTextResponseBodyDataSubImagesSubImageRect {
	s.CenterX = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesSubImageRect) SetCenterY(v int32) *RecognizeAllTextResponseBodyDataSubImagesSubImageRect {
	s.CenterY = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesSubImageRect) SetHeight(v int32) *RecognizeAllTextResponseBodyDataSubImagesSubImageRect {
	s.Height = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesSubImageRect) SetWidth(v int32) *RecognizeAllTextResponseBodyDataSubImagesSubImageRect {
	s.Width = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesTableInfo struct {
	// example:
	//
	// 2
	TableCount   *int32                                                            `json:"TableCount,omitempty" xml:"TableCount,omitempty"`
	TableDetails []*RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails `json:"TableDetails,omitempty" xml:"TableDetails,omitempty" type:"Repeated"`
	// example:
	//
	// https://example.xlsx
	TableExcel *string `json:"TableExcel,omitempty" xml:"TableExcel,omitempty"`
	// example:
	//
	// https://example.html
	TableHtml *string `json:"TableHtml,omitempty" xml:"TableHtml,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfo) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfo) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfo) SetTableCount(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfo {
	s.TableCount = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfo) SetTableDetails(v []*RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails) *RecognizeAllTextResponseBodyDataSubImagesTableInfo {
	s.TableDetails = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfo) SetTableExcel(v string) *RecognizeAllTextResponseBodyDataSubImagesTableInfo {
	s.TableExcel = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfo) SetTableHtml(v string) *RecognizeAllTextResponseBodyDataSubImagesTableInfo {
	s.TableHtml = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails struct {
	// example:
	//
	// 29
	CellCount   *int32                                                                       `json:"CellCount,omitempty" xml:"CellCount,omitempty"`
	CellDetails []*RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails `json:"CellDetails,omitempty" xml:"CellDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	ColumnCount *int32                                                                `json:"ColumnCount,omitempty" xml:"ColumnCount,omitempty"`
	Footer      *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsFooter `json:"Footer,omitempty" xml:"Footer,omitempty" type:"Struct"`
	Header      *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsHeader `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	// example:
	//
	// 10
	RowCount *int32 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// example:
	//
	// 0
	TableId     *int32                                                                       `json:"TableId,omitempty" xml:"TableId,omitempty"`
	TablePoints []*RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTablePoints `json:"TablePoints,omitempty" xml:"TablePoints,omitempty" type:"Repeated"`
	TableRect   *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTableRect     `json:"TableRect,omitempty" xml:"TableRect,omitempty" type:"Struct"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails) SetCellCount(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails {
	s.CellCount = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails) SetCellDetails(v []*RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails {
	s.CellDetails = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails) SetColumnCount(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails {
	s.ColumnCount = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails) SetFooter(v *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsFooter) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails {
	s.Footer = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails) SetHeader(v *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsHeader) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails {
	s.Header = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails) SetRowCount(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails {
	s.RowCount = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails) SetTableId(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails {
	s.TableId = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails) SetTablePoints(v []*RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTablePoints) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails {
	s.TablePoints = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails) SetTableRect(v *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTableRect) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetails {
	s.TableRect = v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails struct {
	BlockList []*int32 `json:"BlockList,omitempty" xml:"BlockList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	CellAngle   *int32  `json:"CellAngle,omitempty" xml:"CellAngle,omitempty"`
	CellContent *string `json:"CellContent,omitempty" xml:"CellContent,omitempty"`
	// example:
	//
	// 0
	CellId     *int32                                                                                 `json:"CellId,omitempty" xml:"CellId,omitempty"`
	CellPoints []*RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellPoints `json:"CellPoints,omitempty" xml:"CellPoints,omitempty" type:"Repeated"`
	CellRect   *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellRect     `json:"CellRect,omitempty" xml:"CellRect,omitempty" type:"Struct"`
	// example:
	//
	// 5
	ColumnEnd *int32 `json:"ColumnEnd,omitempty" xml:"ColumnEnd,omitempty"`
	// example:
	//
	// 2
	ColumnStart *int32 `json:"ColumnStart,omitempty" xml:"ColumnStart,omitempty"`
	// example:
	//
	// 0
	RowEnd *int32 `json:"RowEnd,omitempty" xml:"RowEnd,omitempty"`
	// example:
	//
	// 0
	RowStart *int32 `json:"RowStart,omitempty" xml:"RowStart,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails) SetBlockList(v []*int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails {
	s.BlockList = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails) SetCellAngle(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails {
	s.CellAngle = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails) SetCellContent(v string) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails {
	s.CellContent = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails) SetCellId(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails {
	s.CellId = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails) SetCellPoints(v []*RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellPoints) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails {
	s.CellPoints = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails) SetCellRect(v *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellRect) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails {
	s.CellRect = v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails) SetColumnEnd(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails {
	s.ColumnEnd = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails) SetColumnStart(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails {
	s.ColumnStart = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails) SetRowEnd(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails {
	s.RowEnd = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails) SetRowStart(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetails {
	s.RowStart = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellPoints struct {
	// example:
	//
	// 100
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 200
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellPoints) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellPoints) SetX(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellPoints {
	s.X = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellPoints) SetY(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellPoints {
	s.Y = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellRect struct {
	// example:
	//
	// 100
	CenterX *int32 `json:"CenterX,omitempty" xml:"CenterX,omitempty"`
	// example:
	//
	// 200
	CenterY *int32 `json:"CenterY,omitempty" xml:"CenterY,omitempty"`
	// example:
	//
	// 20
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 20
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellRect) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellRect) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellRect) SetCenterX(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellRect {
	s.CenterX = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellRect) SetCenterY(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellRect {
	s.CenterY = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellRect) SetHeight(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellRect {
	s.Height = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellRect) SetWidth(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsCellDetailsCellRect {
	s.Width = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsFooter struct {
	// example:
	//
	// 0
	BlockId  *int32    `json:"BlockId,omitempty" xml:"BlockId,omitempty"`
	Contents []*string `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsFooter) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsFooter) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsFooter) SetBlockId(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsFooter {
	s.BlockId = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsFooter) SetContents(v []*string) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsFooter {
	s.Contents = v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsHeader struct {
	// example:
	//
	// 0
	BlockId  *int32    `json:"BlockId,omitempty" xml:"BlockId,omitempty"`
	Contents []*string `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsHeader) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsHeader) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsHeader) SetBlockId(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsHeader {
	s.BlockId = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsHeader) SetContents(v []*string) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsHeader {
	s.Contents = v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTablePoints struct {
	// example:
	//
	// 100
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 200
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTablePoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTablePoints) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTablePoints) SetX(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTablePoints {
	s.X = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTablePoints) SetY(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTablePoints {
	s.Y = &v
	return s
}

type RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTableRect struct {
	// example:
	//
	// 100
	CenterX *int32 `json:"CenterX,omitempty" xml:"CenterX,omitempty"`
	// example:
	//
	// 200
	CenterY *int32 `json:"CenterY,omitempty" xml:"CenterY,omitempty"`
	// example:
	//
	// 100
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 100
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTableRect) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTableRect) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTableRect) SetCenterX(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTableRect {
	s.CenterX = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTableRect) SetCenterY(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTableRect {
	s.CenterY = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTableRect) SetHeight(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTableRect {
	s.Height = &v
	return s
}

func (s *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTableRect) SetWidth(v int32) *RecognizeAllTextResponseBodyDataSubImagesTableInfoTableDetailsTableRect {
	s.Width = &v
	return s
}

type RecognizeAllTextResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeAllTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeAllTextResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAllTextResponse) GoString() string {
	return s.String()
}

func (s *RecognizeAllTextResponse) SetHeaders(v map[string]*string) *RecognizeAllTextResponse {
	s.Headers = v
	return s
}

func (s *RecognizeAllTextResponse) SetStatusCode(v int32) *RecognizeAllTextResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeAllTextResponse) SetBody(v *RecognizeAllTextResponseBody) *RecognizeAllTextResponse {
	s.Body = v
	return s
}

type RecognizeBankAcceptanceRequest struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i1/O1CN016eNk0d1ubhKP4y6gK_!!6000000006056-2-tps-631-570.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBankAcceptanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankAcceptanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBankAcceptanceRequest) SetUrl(v string) *RecognizeBankAcceptanceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeBankAcceptanceRequest) SetBody(v io.Reader) *RecognizeBankAcceptanceRequest {
	s.Body = v
	return s
}

type RecognizeBankAcceptanceResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"data": {"出票日期": "2021-03-31", "到期日期": "2021-06-30", "票据状态": "提示收票已签收", "票据号码": "1306595000098202103", "出票人全称": "广东格林精密部件股份有限公司", "出票人账号": "9550880016631800646", "出票人开户银行": "广发银行股份有限公司惠州江北支行", "收票人全称": "限公司", "收票人账号": "2008022009200267322", "收票人开户银行": "中国工商银行惠州仲恺高新区支行", "票据金额大写": "贰拾万捌仟捌佰肆拾伍元整", "票据金额小写": "208845.00", "承兑人全称": "广发银行股份有限公司惠州江北支行", "承兑人账号": "", "承兑人开户行行号": "306595000098", "承兑人开户行名称": "广发银行股份有限公司惠州江北支行", "交易合同号": "", "能否转让": "可转让", "承兑日期": "2021-03-31"}, "ftype": 0, "height": 570, "orgHeight": 570, "orgWidth": 631, "prism_keyValueInfo": [{"key": "出票日期", "keyProb": 100, "value": "2021-03-31", "valuePos": [{"x": 148, "y": 37}, {"x": 148, "y": 48}, {"x": 86, "y": 48}, {"x": 86, "y": 37}], "valueProb": 100}, {"key": "到期日期", "keyProb": 100, "value": "2021-06-30", "valuePos": [{"x": 150, "y": 54}, {"x": 150, "y": 66}, {"x": 86, "y": 66}, {"x": 86, "y": 54}], "valueProb": 100}, {"key": "票据状态", "keyProb": 100, "value": "提示收票已签收", "valuePos": [{"x": 466, "y": 35}, {"x": 466, "y": 50}, {"x": 379, "y": 50}, {"x": 379, "y": 35}], "valueProb": 100}, {"key": "票据号码", "keyProb": 96, "value": "1306595000098202103", "valuePos": [{"x": 509, "y": 55}, {"x": 509, "y": 66}, {"x": 379, "y": 66}, {"x": 379, "y": 54}], "valueProb": 96}, {"key": "出票人全称", "keyProb": 100, "value": "广东格林精密部件股份有限公司", "valuePos": [{"x": 274, "y": 73}, {"x": 274, "y": 88}, {"x": 102, "y": 88}, {"x": 102, "y": 73}], "valueProb": 100}, {"key": "出票人账号", "keyProb": 97, "value": "9550880016631800646", "valuePos": [{"x": 220, "y": 94}, {"x": 220, "y": 106}, {"x": 104, "y": 106}, {"x": 104, "y": 94}], "valueProb": 97}, {"key": "出票人开户银行", "keyProb": 100, "value": "广发银行股份有限公司惠州江北支行", "valuePos": [{"x": 297, "y": 119}, {"x": 297, "y": 134}, {"x": 105, "y": 134}, {"x": 105, "y": 118}], "valueProb": 100}, {"key": "收票人全称", "keyProb": 100, "value": "限公司", "valuePos": [{"x": 548, "y": 75}, {"x": 588, "y": 74}, {"x": 589, "y": 86}, {"x": 548, "y": 88}], "valueProb": 100}, {"key": "收票人账号", "keyProb": 99, "value": "2008022009200267322", "valuePos": [{"x": 536, "y": 96}, {"x": 536, "y": 106}, {"x": 418, "y": 106}, {"x": 418, "y": 96}], "valueProb": 99}, {"key": "收票人开户银行", "keyProb": 100, "value": "中国工商银行惠州仲恺高新区支行", "valuePos": [{"x": 585, "y": 111}, {"x": 586, "y": 136}, {"x": 420, "y": 137}, {"x": 419, "y": 113}], "valueProb": 100}, {"key": "票据金额大写", "keyProb": 100, "value": "贰拾万捌仟捌佰肆拾伍元整", "valuePos": [{"x": 299, "y": 162}, {"x": 299, "y": 178}, {"x": 152, "y": 178}, {"x": 152, "y": 162}], "valueProb": 100}, {"key": "票据金额小写", "keyProb": 100, "value": "208845.00", "valuePos": [{"x": 299, "y": 162}, {"x": 299, "y": 178}, {"x": 152, "y": 178}, {"x": 152, "y": 162}], "valueProb": 100}, {"key": "承兑人全称", "keyProb": 100, "value": "广发银行股份有限公司惠州江北支行", "valuePos": [{"x": 309, "y": 208}, {"x": 309, "y": 234}, {"x": 178, "y": 234}, {"x": 178, "y": 208}], "valueProb": 100}, {"key": "承兑人账号", "keyProb": 98, "value": "", "valuePos": [{"x": 187, "y": 247}, {"x": 187, "y": 258}, {"x": 180, "y": 258}, {"x": 180, "y": 247}], "valueProb": 98}, {"key": "承兑人开户行行号", "keyProb": 100, "value": "306595000098", "valuePos": [{"x": 493, "y": 216}, {"x": 493, "y": 227}, {"x": 420, "y": 227}, {"x": 420, "y": 216}], "valueProb": 100}, {"key": "承兑人开户行名称", "keyProb": 100, "value": "广发银行股份有限公司惠州江北支行", "valuePos": [{"x": 419, "y": 239}, {"x": 586, "y": 239}, {"x": 586, "y": 264}, {"x": 419, "y": 264}], "valueProb": 100}, {"key": "交易合同号", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "能否转让", "keyProb": 100, "value": "可转让", "valuePos": [{"x": 143, "y": 307}, {"x": 143, "y": 322}, {"x": 105, "y": 322}, {"x": 105, "y": 307}], "valueProb": 100}, {"key": "承兑日期", "keyProb": 100, "value": "2021-03-31", "valuePos": [{"x": 404, "y": 314}, {"x": 465, "y": 314}, {"x": 465, "y": 326}, {"x": 404, "y": 326}], "valueProb": 100}], "sliceRect": {"x0": 11, "y0": 90, "x1": 614, "y1": 93, "x2": 614, "y2": 490, "x3": 10, "y3": 489}, "width": 631}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBankAcceptanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankAcceptanceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBankAcceptanceResponseBody) SetCode(v string) *RecognizeBankAcceptanceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBankAcceptanceResponseBody) SetData(v string) *RecognizeBankAcceptanceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBankAcceptanceResponseBody) SetMessage(v string) *RecognizeBankAcceptanceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeBankAcceptanceResponseBody) SetRequestId(v string) *RecognizeBankAcceptanceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBankAcceptanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeBankAcceptanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBankAcceptanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankAcceptanceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBankAcceptanceResponse) SetHeaders(v map[string]*string) *RecognizeBankAcceptanceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBankAcceptanceResponse) SetStatusCode(v int32) *RecognizeBankAcceptanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBankAcceptanceResponse) SetBody(v *RecognizeBankAcceptanceResponseBody) *RecognizeBankAcceptanceResponse {
	s.Body = v
	return s
}

type RecognizeBankAccountLicenseRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB17liGda67gK0jSZFHXXa9jVXa-1375-1000.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBankAccountLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankAccountLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBankAccountLicenseRequest) SetUrl(v string) *RecognizeBankAccountLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeBankAccountLicenseRequest) SetBody(v io.Reader) *RecognizeBankAccountLicenseRequest {
	s.Body = v
	return s
}

type RecognizeBankAccountLicenseResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBankAccountLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankAccountLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBankAccountLicenseResponseBody) SetCode(v string) *RecognizeBankAccountLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBankAccountLicenseResponseBody) SetData(v string) *RecognizeBankAccountLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBankAccountLicenseResponseBody) SetMessage(v string) *RecognizeBankAccountLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeBankAccountLicenseResponseBody) SetRequestId(v string) *RecognizeBankAccountLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBankAccountLicenseResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeBankAccountLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBankAccountLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankAccountLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBankAccountLicenseResponse) SetHeaders(v map[string]*string) *RecognizeBankAccountLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBankAccountLicenseResponse) SetStatusCode(v int32) *RecognizeBankAccountLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBankAccountLicenseResponse) SetBody(v *RecognizeBankAccountLicenseResponseBody) *RecognizeBankAccountLicenseResponse {
	s.Body = v
	return s
}

type RecognizeBankCardRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1fL.fiCzqK1RjSZPcXXbTepXa-3116-2139.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBankCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardRequest) SetUrl(v string) *RecognizeBankCardRequest {
	s.Url = &v
	return s
}

func (s *RecognizeBankCardRequest) SetBody(v io.Reader) *RecognizeBankCardRequest {
	s.Body = v
	return s
}

type RecognizeBankCardResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBankCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardResponseBody) SetCode(v string) *RecognizeBankCardResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBankCardResponseBody) SetData(v string) *RecognizeBankCardResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBankCardResponseBody) SetMessage(v string) *RecognizeBankCardResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeBankCardResponseBody) SetRequestId(v string) *RecognizeBankCardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBankCardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeBankCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBankCardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankCardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardResponse) SetHeaders(v map[string]*string) *RecognizeBankCardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBankCardResponse) SetStatusCode(v int32) *RecognizeBankCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBankCardResponse) SetBody(v *RecognizeBankCardResponseBody) *RecognizeBankCardResponse {
	s.Body = v
	return s
}

type RecognizeBasicRequest struct {
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBasicRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBasicRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBasicRequest) SetNeedRotate(v bool) *RecognizeBasicRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeBasicRequest) SetUrl(v string) *RecognizeBasicRequest {
	s.Url = &v
	return s
}

func (s *RecognizeBasicRequest) SetBody(v io.Reader) *RecognizeBasicRequest {
	s.Body = v
	return s
}

type RecognizeBasicResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBasicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBasicResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBasicResponseBody) SetCode(v string) *RecognizeBasicResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBasicResponseBody) SetData(v string) *RecognizeBasicResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBasicResponseBody) SetMessage(v string) *RecognizeBasicResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeBasicResponseBody) SetRequestId(v string) *RecognizeBasicResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBasicResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeBasicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBasicResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBasicResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBasicResponse) SetHeaders(v map[string]*string) *RecognizeBasicResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBasicResponse) SetStatusCode(v int32) *RecognizeBasicResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBasicResponse) SetBody(v *RecognizeBasicResponseBody) *RecognizeBasicResponse {
	s.Body = v
	return s
}

type RecognizeBirthCertificationRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1P6Yll8Bh1e4jSZFhXXcC9VXa-1381-962.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBirthCertificationRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBirthCertificationRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBirthCertificationRequest) SetUrl(v string) *RecognizeBirthCertificationRequest {
	s.Url = &v
	return s
}

func (s *RecognizeBirthCertificationRequest) SetBody(v io.Reader) *RecognizeBirthCertificationRequest {
	s.Body = v
	return s
}

type RecognizeBirthCertificationResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBirthCertificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBirthCertificationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBirthCertificationResponseBody) SetCode(v string) *RecognizeBirthCertificationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBirthCertificationResponseBody) SetData(v string) *RecognizeBirthCertificationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBirthCertificationResponseBody) SetMessage(v string) *RecognizeBirthCertificationResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeBirthCertificationResponseBody) SetRequestId(v string) *RecognizeBirthCertificationResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBirthCertificationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeBirthCertificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBirthCertificationResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBirthCertificationResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBirthCertificationResponse) SetHeaders(v map[string]*string) *RecognizeBirthCertificationResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBirthCertificationResponse) SetStatusCode(v int32) *RecognizeBirthCertificationResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBirthCertificationResponse) SetBody(v *RecognizeBirthCertificationResponseBody) *RecognizeBirthCertificationResponse {
	s.Body = v
	return s
}

type RecognizeBusShipTicketRequest struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i2/O1CN010iDcM7218ZQJtJyGX_!!6000000006940-0-tps-936-541.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBusShipTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusShipTicketRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBusShipTicketRequest) SetUrl(v string) *RecognizeBusShipTicketRequest {
	s.Url = &v
	return s
}

func (s *RecognizeBusShipTicketRequest) SetBody(v io.Reader) *RecognizeBusShipTicketRequest {
	s.Body = v
	return s
}

type RecognizeBusShipTicketResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"angle":0,"data":{"title":"南通汽运实业集团有限公司旅客运输专用发票","formType":"发票联","invoiceCode":"132061981313","invoiceNumber":"05591493","date":"2020-01-20","time":"12:30","departureStation":"南通东站","arrivalStation":"上海总站","totalAmount":"56.00","passengerName":"颜跃第","idcardNo":"3210****2218"},"ftype":0,"height":541,"orgHeight":541,"orgWidth":936,"prism_keyValueInfo":[{"key":"title","keyProb":97,"value":"南通汽运实业集团有限公司旅客运输专用发票","valuePos":[{"x":508,"y":16},{"x":509,"y":94},{"x":91,"y":95},{"x":90,"y":18}],"valueProb":98},{"key":"formType","keyProb":100,"value":"发票联","valuePos":[{"x":388,"y":119},{"x":388,"y":157},{"x":209,"y":157},{"x":209,"y":118}],"valueProb":100},{"key":"invoiceCode","keyProb":100,"value":"132061981313","valuePos":[{"x":929,"y":127},{"x":929,"y":161},{"x":699,"y":162},{"x":698,"y":128}],"valueProb":100},{"key":"invoiceNumber","keyProb":100,"value":"05591493","valuePos":[{"x":851,"y":167},{"x":851,"y":199},{"x":696,"y":201},{"x":695,"y":168}],"valueProb":100},{"key":"date","keyProb":100,"value":"2020-01-20","valuePos":[{"x":185,"y":356},{"x":186,"y":384},{"x":62,"y":385},{"x":62,"y":358}],"valueProb":100},{"key":"time","keyProb":100,"value":"12:30","valuePos":[{"x":186,"y":385},{"x":186,"y":358},{"x":264,"y":359},{"x":264,"y":386}],"valueProb":100},{"key":"departureStation","keyProb":100,"value":"南通东站","valuePos":[{"x":66,"y":304},{"x":66,"y":271},{"x":187,"y":274},{"x":186,"y":308}],"valueProb":100},{"key":"arrivalStation","keyProb":100,"value":"上海总站","valuePos":[{"x":205,"y":306},{"x":205,"y":273},{"x":326,"y":276},{"x":325,"y":308}],"valueProb":100},{"key":"totalAmount","keyProb":100,"value":"56.00","valuePos":[{"x":402,"y":278},{"x":402,"y":306},{"x":366,"y":306},{"x":366,"y":278}],"valueProb":100},{"key":"passengerName","keyProb":97,"value":"颜跃第","valuePos":[{"x":426,"y":466},{"x":427,"y":434},{"x":516,"y":435},{"x":516,"y":468}],"valueProb":97},{"key":"idcardNo","keyProb":100,"value":"3210****2218","valuePos":[{"x":729,"y":441},{"x":729,"y":468},{"x":548,"y":468},{"x":548,"y":441}],"valueProb":100}],"sliceRect":{"x0":0,"y0":14,"x1":934,"y1":18,"x2":936,"y2":541,"x3":0,"y3":541},"width":936}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBusShipTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusShipTicketResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBusShipTicketResponseBody) SetCode(v string) *RecognizeBusShipTicketResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBusShipTicketResponseBody) SetData(v string) *RecognizeBusShipTicketResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBusShipTicketResponseBody) SetMessage(v string) *RecognizeBusShipTicketResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeBusShipTicketResponseBody) SetRequestId(v string) *RecognizeBusShipTicketResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBusShipTicketResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeBusShipTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBusShipTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusShipTicketResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBusShipTicketResponse) SetHeaders(v map[string]*string) *RecognizeBusShipTicketResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBusShipTicketResponse) SetStatusCode(v int32) *RecognizeBusShipTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBusShipTicketResponse) SetBody(v *RecognizeBusShipTicketResponseBody) *RecognizeBusShipTicketResponse {
	s.Body = v
	return s
}

type RecognizeBusinessLicenseRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1nnHJNSrqK1RjSZK9XXXyypXa-564-829.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBusinessLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseRequest) SetUrl(v string) *RecognizeBusinessLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeBusinessLicenseRequest) SetBody(v io.Reader) *RecognizeBusinessLicenseRequest {
	s.Body = v
	return s
}

type RecognizeBusinessLicenseResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponseBody) SetCode(v string) *RecognizeBusinessLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBody) SetData(v string) *RecognizeBusinessLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBody) SetMessage(v string) *RecognizeBusinessLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBody) SetRequestId(v string) *RecognizeBusinessLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBusinessLicenseResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeBusinessLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBusinessLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponse) SetHeaders(v map[string]*string) *RecognizeBusinessLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBusinessLicenseResponse) SetStatusCode(v int32) *RecognizeBusinessLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBusinessLicenseResponse) SetBody(v *RecognizeBusinessLicenseResponseBody) *RecognizeBusinessLicenseResponse {
	s.Body = v
	return s
}

type RecognizeCarInvoiceRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1hC7bXCzqK1RjSZPcXXbTepXa-832-616.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCarInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCarInvoiceRequest) SetUrl(v string) *RecognizeCarInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeCarInvoiceRequest) SetBody(v io.Reader) *RecognizeCarInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeCarInvoiceResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeCarInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCarInvoiceResponseBody) SetCode(v string) *RecognizeCarInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCarInvoiceResponseBody) SetData(v string) *RecognizeCarInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCarInvoiceResponseBody) SetMessage(v string) *RecognizeCarInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeCarInvoiceResponseBody) SetRequestId(v string) *RecognizeCarInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeCarInvoiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeCarInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCarInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeCarInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeCarInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeCarInvoiceResponse) SetStatusCode(v int32) *RecognizeCarInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeCarInvoiceResponse) SetBody(v *RecognizeCarInvoiceResponseBody) *RecognizeCarInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeCarNumberRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCarNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarNumberRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCarNumberRequest) SetUrl(v string) *RecognizeCarNumberRequest {
	s.Url = &v
	return s
}

func (s *RecognizeCarNumberRequest) SetBody(v io.Reader) *RecognizeCarNumberRequest {
	s.Body = v
	return s
}

type RecognizeCarNumberResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeCarNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarNumberResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCarNumberResponseBody) SetCode(v string) *RecognizeCarNumberResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCarNumberResponseBody) SetData(v string) *RecognizeCarNumberResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCarNumberResponseBody) SetMessage(v string) *RecognizeCarNumberResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeCarNumberResponseBody) SetRequestId(v string) *RecognizeCarNumberResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeCarNumberResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeCarNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCarNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarNumberResponse) GoString() string {
	return s.String()
}

func (s *RecognizeCarNumberResponse) SetHeaders(v map[string]*string) *RecognizeCarNumberResponse {
	s.Headers = v
	return s
}

func (s *RecognizeCarNumberResponse) SetStatusCode(v int32) *RecognizeCarNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeCarNumberResponse) SetBody(v *RecognizeCarNumberResponseBody) *RecognizeCarNumberResponse {
	s.Body = v
	return s
}

type RecognizeCarVinCodeRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCarVinCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarVinCodeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCarVinCodeRequest) SetUrl(v string) *RecognizeCarVinCodeRequest {
	s.Url = &v
	return s
}

func (s *RecognizeCarVinCodeRequest) SetBody(v io.Reader) *RecognizeCarVinCodeRequest {
	s.Body = v
	return s
}

type RecognizeCarVinCodeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeCarVinCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarVinCodeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCarVinCodeResponseBody) SetCode(v string) *RecognizeCarVinCodeResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCarVinCodeResponseBody) SetData(v string) *RecognizeCarVinCodeResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCarVinCodeResponseBody) SetMessage(v string) *RecognizeCarVinCodeResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeCarVinCodeResponseBody) SetRequestId(v string) *RecognizeCarVinCodeResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeCarVinCodeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeCarVinCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCarVinCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCarVinCodeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeCarVinCodeResponse) SetHeaders(v map[string]*string) *RecognizeCarVinCodeResponse {
	s.Headers = v
	return s
}

func (s *RecognizeCarVinCodeResponse) SetStatusCode(v int32) *RecognizeCarVinCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeCarVinCodeResponse) SetBody(v *RecognizeCarVinCodeResponseBody) *RecognizeCarVinCodeResponse {
	s.Body = v
	return s
}

type RecognizeChinesePassportRequest struct {
	// example:
	//
	// true
	OutputFigure *bool `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i2/O1CN01yaQKCT1PrUsTWqgSK_!!6000000001894-0-tps-271-186.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeChinesePassportRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeChinesePassportRequest) GoString() string {
	return s.String()
}

func (s *RecognizeChinesePassportRequest) SetOutputFigure(v bool) *RecognizeChinesePassportRequest {
	s.OutputFigure = &v
	return s
}

func (s *RecognizeChinesePassportRequest) SetUrl(v string) *RecognizeChinesePassportRequest {
	s.Url = &v
	return s
}

func (s *RecognizeChinesePassportRequest) SetBody(v io.Reader) *RecognizeChinesePassportRequest {
	s.Body = v
	return s
}

type RecognizeChinesePassportResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"data":{"passportType":"P","countryCode":"CHN","passportNumber":"E90000082","nameEn":",ZHENGJIANGANGUX","name":"","sex":"3.7F","birthPlace":"北京/BEIJIHG","nationality":"","issuePlace":"山东/SHANDON","issueAuthority":"公安部出入境管理局","mrzLine1":"POCHNZHENGJIAN<<YANGBEN<<<<<<<<<<<<<<<<<<<<<","mrzLine2":"E900000821CHN8108038F2110189NGKELMPONBPJB978","validToDate":"2921.DCF.3B","birthDate":"08.1981","issueDate":"91.1010.19"},"ftype":0,"height":186,"orgHeight":186,"orgWidth":271,"prism_keyValueInfo":[{"key":"passportType","keyProb":100,"value":"P","valuePos":[{"x":93,"y":26},{"x":93,"y":33},{"x":87,"y":33},{"x":87,"y":26}],"valueProb":100},{"key":"countryCode","keyProb":92,"value":"CHN","valuePos":[{"x":143,"y":26},{"x":143,"y":33},{"x":126,"y":33},{"x":126,"y":26}],"valueProb":92},{"key":"passportNumber","keyProb":100,"value":"E90000082","valuePos":[{"x":173,"y":29},{"x":230,"y":28},{"x":230,"y":35},{"x":174,"y":37}],"valueProb":100},{"key":"nameEn","keyProb":87,"value":",ZHENGJIANGANGUX","valuePos":[{"x":88,"y":55},{"x":89,"y":48},{"x":166,"y":49},{"x":166,"y":57}],"valueProb":87},{"key":"name","keyProb":100,"value":"","valueProb":100},{"key":"sex","keyProb":99,"value":"3.7F","valuePos":[{"x":103,"y":67},{"x":103,"y":74},{"x":87,"y":74},{"x":87,"y":67}],"valueProb":99},{"key":"birthPlace","keyProb":98,"value":"北京/BEIJIHG","valuePos":[{"x":133,"y":83},{"x":133,"y":91},{"x":87,"y":91},{"x":87,"y":83}],"valueProb":98},{"key":"nationality","keyProb":100,"value":"","valueProb":100},{"key":"issuePlace","keyProb":99,"value":"山东/SHANDON","valuePos":[{"x":136,"y":100},{"x":136,"y":108},{"x":88,"y":108},{"x":88,"y":100}],"valueProb":99},{"key":"issueAuthority","keyProb":79,"value":"公安部出入境管理局","valuePos":[{"x":87,"y":118},{"x":142,"y":118},{"x":142,"y":125},{"x":87,"y":125}],"valueProb":79},{"key":"mrzLine1","keyProb":100,"value":"POCHNZHENGJIAN<<YANGBEN<<<<<<<<<<<<<<<<<<<<<","valuePos":[{"x":12,"y":153},{"x":252,"y":152},{"x":252,"y":159},{"x":12,"y":161}],"valueProb":100},{"key":"mrzLine2","keyProb":99,"value":"E900000821CHN8108038F2110189NGKELMPONBPJB978","valuePos":[{"x":11,"y":166},{"x":253,"y":165},{"x":253,"y":173},{"x":12,"y":175}],"valueProb":99},{"key":"validToDate","keyProb":60,"value":"2921.DCF.3B","valuePos":[{"x":170,"y":107},{"x":171,"y":99},{"x":226,"y":101},{"x":225,"y":108}],"valueProb":86},{"key":"birthDate","keyProb":100,"value":"08.1981","valuePos":[{"x":209,"y":67},{"x":209,"y":74},{"x":181,"y":74},{"x":181,"y":67}],"valueProb":99},{"key":"issueDate","keyProb":82,"value":"91.1010.19","valuePos":[{"x":226,"y":83},{"x":226,"y":90},{"x":170,"y":90},{"x":170,"y":83}],"valueProb":84}],"sliceRect":{"x0":1,"y0":1,"x1":269,"y1":1,"x2":269,"y2":184,"x3":1,"y3":183},"width":271}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeChinesePassportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeChinesePassportResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeChinesePassportResponseBody) SetCode(v string) *RecognizeChinesePassportResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeChinesePassportResponseBody) SetData(v string) *RecognizeChinesePassportResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeChinesePassportResponseBody) SetMessage(v string) *RecognizeChinesePassportResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeChinesePassportResponseBody) SetRequestId(v string) *RecognizeChinesePassportResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeChinesePassportResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeChinesePassportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeChinesePassportResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeChinesePassportResponse) GoString() string {
	return s.String()
}

func (s *RecognizeChinesePassportResponse) SetHeaders(v map[string]*string) *RecognizeChinesePassportResponse {
	s.Headers = v
	return s
}

func (s *RecognizeChinesePassportResponse) SetStatusCode(v int32) *RecognizeChinesePassportResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeChinesePassportResponse) SetBody(v *RecognizeChinesePassportResponseBody) *RecognizeChinesePassportResponse {
	s.Body = v
	return s
}

type RecognizeCommonPrintedInvoiceRequest struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i2/O1CN01XU9dTh1O4CdHxXhMw_!!6000000001651-0-tps-1437-909.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCommonPrintedInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCommonPrintedInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCommonPrintedInvoiceRequest) SetUrl(v string) *RecognizeCommonPrintedInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeCommonPrintedInvoiceRequest) SetBody(v io.Reader) *RecognizeCommonPrintedInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeCommonPrintedInvoiceResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"angle":0,"data":{"title":"浙江通用机打发票","formType":"发票联","invoiceCode":"133041930432","invoiceNumber":"01488558","printedInvoiceCode":"","printedInvoiceNumber":"","invoiceDate":"2019-11-19","totalAmount":"170.00","sellerName":"嘉兴市南湖区余新镇瘦汁味餐饮店","sellerTaxNumber":"92330402MA28B4LL4B","purchaserName":"阿里巴巴俪人购(上海)电子商务有限公司","purchaserTaxNumber":"91310114312356647G","drawer":"高伟","recipient":"","remarks":"","invoiceDetails":[{"itemName":"餐饮费","unit":"","quantity":"1","unitPrice":"170.00","amount":"170.00"}]},"ftype":0,"height":909,"orgHeight":909,"orgWidth":1437,"prism_keyValueInfo":[{"key":"title","keyProb":100,"value":"浙江通用机打发票","valuePos":[{"x":431,"y":68},{"x":843,"y":62},{"x":843,"y":125},{"x":431,"y":130}],"valueProb":100},{"key":"formType","keyProb":100,"value":"发票联","valuePos":[{"x":507,"y":154},{"x":767,"y":152},{"x":768,"y":214},{"x":508,"y":215}],"valueProb":100},{"key":"invoiceCode","keyProb":100,"value":"133041930432","valuePos":[{"x":990,"y":134},{"x":1283,"y":131},{"x":1283,"y":167},{"x":991,"y":171}],"valueProb":100},{"key":"invoiceNumber","keyProb":100,"value":"01488558","valuePos":[{"x":999,"y":195},{"x":1197,"y":193},{"x":1198,"y":234},{"x":999,"y":235}],"valueProb":100},{"key":"printedInvoiceCode","keyProb":100,"value":"","valueProb":100},{"key":"printedInvoiceNumber","keyProb":100,"value":"","valueProb":100},{"key":"invoiceDate","keyProb":100,"value":"2019-11-19","valuePos":[{"x":153,"y":280},{"x":351,"y":278},{"x":351,"y":309},{"x":154,"y":312}],"valueProb":100},{"key":"totalAmount","keyProb":100,"value":"170.00","valuePos":[{"x":300,"y":752},{"x":461,"y":749},{"x":462,"y":786},{"x":300,"y":788}],"valueProb":100},{"key":"sellerName","keyProb":100,"value":"嘉兴市南湖区余新镇瘦汁味餐饮店","valuePos":[{"x":220,"y":455},{"x":612,"y":450},{"x":612,"y":482},{"x":221,"y":488}],"valueProb":100},{"key":"sellerTaxNumber","keyProb":97,"value":"92330402MA28B4LL4B","valuePos":[{"x":224,"y":511},{"x":476,"y":509},{"x":477,"y":537},{"x":225,"y":539}],"valueProb":97},{"key":"purchaserName","keyProb":98,"value":"阿里巴巴俪人购(上海)电子商务有限公司","valuePos":[{"x":213,"y":327},{"x":714,"y":324},{"x":715,"y":359},{"x":214,"y":363}],"valueProb":98},{"key":"purchaserTaxNumber","keyProb":100,"value":"91310114312356647G","valuePos":[{"x":221,"y":406},{"x":480,"y":402},{"x":481,"y":432},{"x":221,"y":435}],"valueProb":100},{"key":"drawer","keyProb":100,"value":"高伟","valuePos":[{"x":680,"y":819},{"x":680,"y":850},{"x":627,"y":850},{"x":627,"y":819}],"valueProb":100},{"key":"recipient","keyProb":100,"value":"","valueProb":100},{"key":"remarks","keyProb":100,"value":"","valueProb":100},{"key":"invoiceDetails","keyProb":100,"value":"[{\"itemName\":\"餐饮费\",\"unit\":\"\",\"quantity\":\"1\",\"unitPrice\":\"170.00\",\"amount\":\"170.00\"}]","valueProb":100}],"sliceRect":{"x0":0,"y0":7,"x1":1416,"y1":0,"x2":1421,"y2":907,"x3":0,"y3":904},"width":1437}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeCommonPrintedInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCommonPrintedInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCommonPrintedInvoiceResponseBody) SetCode(v string) *RecognizeCommonPrintedInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCommonPrintedInvoiceResponseBody) SetData(v string) *RecognizeCommonPrintedInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCommonPrintedInvoiceResponseBody) SetMessage(v string) *RecognizeCommonPrintedInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeCommonPrintedInvoiceResponseBody) SetRequestId(v string) *RecognizeCommonPrintedInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeCommonPrintedInvoiceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeCommonPrintedInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCommonPrintedInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCommonPrintedInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeCommonPrintedInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeCommonPrintedInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeCommonPrintedInvoiceResponse) SetStatusCode(v int32) *RecognizeCommonPrintedInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeCommonPrintedInvoiceResponse) SetBody(v *RecognizeCommonPrintedInvoiceResponseBody) *RecognizeCommonPrintedInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeCosmeticProduceLicenseRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCosmeticProduceLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCosmeticProduceLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCosmeticProduceLicenseRequest) SetUrl(v string) *RecognizeCosmeticProduceLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeCosmeticProduceLicenseRequest) SetBody(v io.Reader) *RecognizeCosmeticProduceLicenseRequest {
	s.Body = v
	return s
}

type RecognizeCosmeticProduceLicenseResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AA91C84E-7DB9-1951-B8FE-D830076A0473
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeCosmeticProduceLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCosmeticProduceLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCosmeticProduceLicenseResponseBody) SetCode(v string) *RecognizeCosmeticProduceLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCosmeticProduceLicenseResponseBody) SetData(v string) *RecognizeCosmeticProduceLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCosmeticProduceLicenseResponseBody) SetMessage(v string) *RecognizeCosmeticProduceLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeCosmeticProduceLicenseResponseBody) SetRequestId(v string) *RecognizeCosmeticProduceLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeCosmeticProduceLicenseResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeCosmeticProduceLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCosmeticProduceLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCosmeticProduceLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeCosmeticProduceLicenseResponse) SetHeaders(v map[string]*string) *RecognizeCosmeticProduceLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeCosmeticProduceLicenseResponse) SetStatusCode(v int32) *RecognizeCosmeticProduceLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeCosmeticProduceLicenseResponse) SetBody(v *RecognizeCosmeticProduceLicenseResponseBody) *RecognizeCosmeticProduceLicenseResponse {
	s.Body = v
	return s
}

type RecognizeCovidTestReportRequest struct {
	// example:
	//
	// false
	MultipleResult *bool `json:"MultipleResult,omitempty" xml:"MultipleResult,omitempty"`
	// example:
	//
	// http://duguang-database-public.oss-cn-hangzhou.aliyuncs.com/covid_init_covid_test_report/test_report__data_pool_15a4f85478cb1bd69a5d631b182aba69.jpg_item_0_cls_covid_test_report.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCovidTestReportRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCovidTestReportRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCovidTestReportRequest) SetMultipleResult(v bool) *RecognizeCovidTestReportRequest {
	s.MultipleResult = &v
	return s
}

func (s *RecognizeCovidTestReportRequest) SetUrl(v string) *RecognizeCovidTestReportRequest {
	s.Url = &v
	return s
}

func (s *RecognizeCovidTestReportRequest) SetBody(v io.Reader) *RecognizeCovidTestReportRequest {
	s.Body = v
	return s
}

type RecognizeCovidTestReportResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"data": {"name": "张德周", "idNumber": "612401********22010", "samplingDate": "2022-03-30", "samplingTime": "330", "testOrganization": "", "testItem": "", "testResult": ""}, "ftype": 0, "height": 991, "orgHeight": 998, "orgWidth": 1076, "prism_keyValueInfo": [{"key": "name", "keyProb": 100, "value": "张德周", "valuePos": [{"x": 291, "y": 465}, {"x": 473, "y": 463}, {"x": 474, "y": 526}, {"x": 291, "y": 527}], "valueProb": 100}, {"key": "idNumber", "keyProb": 91, "value": "612401********22010", "valuePos": [{"x": 791, "y": 180}, {"x": 791, "y": 227}, {"x": 300, "y": 226}, {"x": 300, "y": 179}], "valueProb": 91}, {"key": "samplingDate", "keyProb": 100, "value": "2022-03-30", "valuePos": [{"x": 597, "y": 775}, {"x": 597, "y": 826}, {"x": 296, "y": 826}, {"x": 296, "y": 775}], "valueProb": 100}, {"key": "samplingTime", "keyProb": 100, "value": "330", "valuePos": [{"x": 412, "y": 684}, {"x": 413, "y": 741}, {"x": 268, "y": 742}, {"x": 268, "y": 686}], "valueProb": 100}, {"key": "testOrganization", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "testItem", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "testResult", "keyProb": 28, "value": "", "valuePos": [{"x": 417, "y": 873}, {"x": 417, "y": 941}, {"x": 298, "y": 941}, {"x": 298, "y": 873}], "valueProb": 28}], "sliceRect": {"x0": 0, "y0": 10, "x1": 1076, "y1": 6, "x2": 1076, "y2": 995, "x3": 0, "y3": 996}, "width": 1076}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeCovidTestReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCovidTestReportResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCovidTestReportResponseBody) SetCode(v string) *RecognizeCovidTestReportResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCovidTestReportResponseBody) SetData(v string) *RecognizeCovidTestReportResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCovidTestReportResponseBody) SetMessage(v string) *RecognizeCovidTestReportResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeCovidTestReportResponseBody) SetRequestId(v string) *RecognizeCovidTestReportResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeCovidTestReportResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeCovidTestReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCovidTestReportResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCovidTestReportResponse) GoString() string {
	return s.String()
}

func (s *RecognizeCovidTestReportResponse) SetHeaders(v map[string]*string) *RecognizeCovidTestReportResponse {
	s.Headers = v
	return s
}

func (s *RecognizeCovidTestReportResponse) SetStatusCode(v int32) *RecognizeCovidTestReportResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeCovidTestReportResponse) SetBody(v *RecognizeCovidTestReportResponseBody) *RecognizeCovidTestReportResponse {
	s.Body = v
	return s
}

type RecognizeCtwoMedicalDeviceManageLicenseRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Hyx0MEH1gK0jSZSyXXXtlpXa-750-1000.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCtwoMedicalDeviceManageLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCtwoMedicalDeviceManageLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseRequest) SetUrl(v string) *RecognizeCtwoMedicalDeviceManageLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseRequest) SetBody(v io.Reader) *RecognizeCtwoMedicalDeviceManageLicenseRequest {
	s.Body = v
	return s
}

type RecognizeCtwoMedicalDeviceManageLicenseResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeCtwoMedicalDeviceManageLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCtwoMedicalDeviceManageLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponseBody) SetCode(v string) *RecognizeCtwoMedicalDeviceManageLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponseBody) SetData(v string) *RecognizeCtwoMedicalDeviceManageLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponseBody) SetMessage(v string) *RecognizeCtwoMedicalDeviceManageLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponseBody) SetRequestId(v string) *RecognizeCtwoMedicalDeviceManageLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeCtwoMedicalDeviceManageLicenseResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeCtwoMedicalDeviceManageLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCtwoMedicalDeviceManageLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCtwoMedicalDeviceManageLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponse) SetHeaders(v map[string]*string) *RecognizeCtwoMedicalDeviceManageLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponse) SetStatusCode(v int32) *RecognizeCtwoMedicalDeviceManageLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeCtwoMedicalDeviceManageLicenseResponse) SetBody(v *RecognizeCtwoMedicalDeviceManageLicenseResponseBody) *RecognizeCtwoMedicalDeviceManageLicenseResponse {
	s.Body = v
	return s
}

type RecognizeDocumentStructureRequest struct {
	// example:
	//
	// false
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// example:
	//
	// false
	NeedSortPage *bool `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
	// example:
	//
	// false
	NoStamp *bool `json:"NoStamp,omitempty" xml:"NoStamp,omitempty"`
	// example:
	//
	// false
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// example:
	//
	// false
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// example:
	//
	// false
	Page *bool `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// false
	Paragraph *bool `json:"Paragraph,omitempty" xml:"Paragraph,omitempty"`
	// example:
	//
	// false
	Row *bool `json:"Row,omitempty" xml:"Row,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i4/O1CN01amMFBF1GUki3NHNzI_!!6000000000626-2-tps-978-1346.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// false
	UseNewStyleOutput *bool     `json:"UseNewStyleOutput,omitempty" xml:"UseNewStyleOutput,omitempty"`
	Body              io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeDocumentStructureRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDocumentStructureRequest) GoString() string {
	return s.String()
}

func (s *RecognizeDocumentStructureRequest) SetNeedRotate(v bool) *RecognizeDocumentStructureRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetNeedSortPage(v bool) *RecognizeDocumentStructureRequest {
	s.NeedSortPage = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetNoStamp(v bool) *RecognizeDocumentStructureRequest {
	s.NoStamp = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetOutputCharInfo(v bool) *RecognizeDocumentStructureRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetOutputTable(v bool) *RecognizeDocumentStructureRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetPage(v bool) *RecognizeDocumentStructureRequest {
	s.Page = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetParagraph(v bool) *RecognizeDocumentStructureRequest {
	s.Paragraph = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetRow(v bool) *RecognizeDocumentStructureRequest {
	s.Row = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetUrl(v string) *RecognizeDocumentStructureRequest {
	s.Url = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetUseNewStyleOutput(v bool) *RecognizeDocumentStructureRequest {
	s.UseNewStyleOutput = &v
	return s
}

func (s *RecognizeDocumentStructureRequest) SetBody(v io.Reader) *RecognizeDocumentStructureRequest {
	s.Body = v
	return s
}

type RecognizeDocumentStructureResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// { 	"content": "2017年河北区实验小学", 	"height": 3509, 	"orgHeight": 3509, 	"orgWidth": 2512, 	"prism_version": "1.0.9", 	"prism_wnum": 126, 	"prism_wordsInfo": [{ 		"angle": -89, 		"direction": 0, 		"height": 541, 		"pos": [{ 			"x": 982, 			"y": 223 		}, { 			"x": 1522, 			"y": 223 		}, { 			"x": 1522, 			"y": 266 		}, { 			"x": 982, 			"y": 266 		}], 		"prob": 99, 		"width": 43, 		"word": "2017年河北区实验小学", 		"x": 1230, 		"y": -26 	}], 	"width": 2512 }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeDocumentStructureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDocumentStructureResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeDocumentStructureResponseBody) SetCode(v string) *RecognizeDocumentStructureResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeDocumentStructureResponseBody) SetData(v string) *RecognizeDocumentStructureResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeDocumentStructureResponseBody) SetMessage(v string) *RecognizeDocumentStructureResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeDocumentStructureResponseBody) SetRequestId(v string) *RecognizeDocumentStructureResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeDocumentStructureResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeDocumentStructureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeDocumentStructureResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDocumentStructureResponse) GoString() string {
	return s.String()
}

func (s *RecognizeDocumentStructureResponse) SetHeaders(v map[string]*string) *RecognizeDocumentStructureResponse {
	s.Headers = v
	return s
}

func (s *RecognizeDocumentStructureResponse) SetStatusCode(v int32) *RecognizeDocumentStructureResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeDocumentStructureResponse) SetBody(v *RecognizeDocumentStructureResponseBody) *RecognizeDocumentStructureResponse {
	s.Body = v
	return s
}

type RecognizeDrivingLicenseRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB18sTuNSzqK1RjSZPxXXc4tVXa-629-416.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeDrivingLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseRequest) SetUrl(v string) *RecognizeDrivingLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeDrivingLicenseRequest) SetBody(v io.Reader) *RecognizeDrivingLicenseRequest {
	s.Body = v
	return s
}

type RecognizeDrivingLicenseResponseBody struct {
	// example:
	//
	// unmatchedImageType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// The type of image didn\\"t match the api.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeDrivingLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseResponseBody) SetCode(v string) *RecognizeDrivingLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBody) SetData(v string) *RecognizeDrivingLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBody) SetMessage(v string) *RecognizeDrivingLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBody) SetRequestId(v string) *RecognizeDrivingLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeDrivingLicenseResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeDrivingLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeDrivingLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseResponse) SetHeaders(v map[string]*string) *RecognizeDrivingLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeDrivingLicenseResponse) SetStatusCode(v int32) *RecognizeDrivingLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeDrivingLicenseResponse) SetBody(v *RecognizeDrivingLicenseResponseBody) *RecognizeDrivingLicenseResponse {
	s.Body = v
	return s
}

type RecognizeEduFormulaRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduFormulaRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduFormulaRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEduFormulaRequest) SetUrl(v string) *RecognizeEduFormulaRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEduFormulaRequest) SetBody(v io.Reader) *RecognizeEduFormulaRequest {
	s.Body = v
	return s
}

type RecognizeEduFormulaResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeEduFormulaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduFormulaResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduFormulaResponseBody) SetCode(v string) *RecognizeEduFormulaResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduFormulaResponseBody) SetData(v string) *RecognizeEduFormulaResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduFormulaResponseBody) SetMessage(v string) *RecognizeEduFormulaResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeEduFormulaResponseBody) SetRequestId(v string) *RecognizeEduFormulaResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeEduFormulaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeEduFormulaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduFormulaResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduFormulaResponse) GoString() string {
	return s.String()
}

func (s *RecognizeEduFormulaResponse) SetHeaders(v map[string]*string) *RecognizeEduFormulaResponse {
	s.Headers = v
	return s
}

func (s *RecognizeEduFormulaResponse) SetStatusCode(v int32) *RecognizeEduFormulaResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeEduFormulaResponse) SetBody(v *RecognizeEduFormulaResponseBody) *RecognizeEduFormulaResponse {
	s.Body = v
	return s
}

type RecognizeEduOralCalculationRequest struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i4/O1CN01diDxZe21hNSkCBf5n_!!6000000007016-0-tps-2268-3024.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduOralCalculationRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduOralCalculationRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEduOralCalculationRequest) SetUrl(v string) *RecognizeEduOralCalculationRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEduOralCalculationRequest) SetBody(v io.Reader) *RecognizeEduOralCalculationRequest {
	s.Body = v
	return s
}

type RecognizeEduOralCalculationResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"height":3024,"mathsInfo":[{"pos":[{"x":128,"y":456},{"x":481,"y":425},{"x":479,"y":526},{"x":127,"y":523}],"result":"right","title":"5 9 - 2 5 = 3 4"}],"orgHeight":3024,"orgWidth":2268,"prism_version":"1.0.9","prism_wnum":0,"prism_wordsInfo":[],"width":2268}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeEduOralCalculationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduOralCalculationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduOralCalculationResponseBody) SetCode(v string) *RecognizeEduOralCalculationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduOralCalculationResponseBody) SetData(v string) *RecognizeEduOralCalculationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduOralCalculationResponseBody) SetMessage(v string) *RecognizeEduOralCalculationResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeEduOralCalculationResponseBody) SetRequestId(v string) *RecognizeEduOralCalculationResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeEduOralCalculationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeEduOralCalculationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduOralCalculationResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduOralCalculationResponse) GoString() string {
	return s.String()
}

func (s *RecognizeEduOralCalculationResponse) SetHeaders(v map[string]*string) *RecognizeEduOralCalculationResponse {
	s.Headers = v
	return s
}

func (s *RecognizeEduOralCalculationResponse) SetStatusCode(v int32) *RecognizeEduOralCalculationResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeEduOralCalculationResponse) SetBody(v *RecognizeEduOralCalculationResponseBody) *RecognizeEduOralCalculationResponse {
	s.Body = v
	return s
}

type RecognizeEduPaperCutRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// question：题目， answer：答案
	CutType *string `json:"CutType,omitempty" xml:"CutType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// scan：扫描图， photo：实拍图
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// example:
	//
	// default:默认, Math:数学, PrimarySchool_Math:小学数学, JHighSchool_Math: 初中数学, Chinese:语文, PrimarySchool_Chinese:小学语文, JHighSchool_Chinese:初中语文, English:英语, PrimarySchool_English:小学英语, JHighSchool_English:初中英语, Physics:物理, JHighSchool_Physics:初中物理, Chemistry: 化学, JHighSchool_Chemistry:初中化学, Biology:生物, JHighSchool_Biology:初中生物, History:历史, JHighSchool_History:初中历史, Geography:地理, JHighSchool_Geography:初中地理, Politics:政治, JHighSchool_Politics:初中政治
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduPaperCutRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperCutRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperCutRequest) SetCutType(v string) *RecognizeEduPaperCutRequest {
	s.CutType = &v
	return s
}

func (s *RecognizeEduPaperCutRequest) SetImageType(v string) *RecognizeEduPaperCutRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeEduPaperCutRequest) SetSubject(v string) *RecognizeEduPaperCutRequest {
	s.Subject = &v
	return s
}

func (s *RecognizeEduPaperCutRequest) SetUrl(v string) *RecognizeEduPaperCutRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEduPaperCutRequest) SetBody(v io.Reader) *RecognizeEduPaperCutRequest {
	s.Body = v
	return s
}

type RecognizeEduPaperCutResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeEduPaperCutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperCutResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperCutResponseBody) SetCode(v string) *RecognizeEduPaperCutResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduPaperCutResponseBody) SetData(v string) *RecognizeEduPaperCutResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduPaperCutResponseBody) SetMessage(v string) *RecognizeEduPaperCutResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeEduPaperCutResponseBody) SetRequestId(v string) *RecognizeEduPaperCutResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeEduPaperCutResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeEduPaperCutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduPaperCutResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperCutResponse) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperCutResponse) SetHeaders(v map[string]*string) *RecognizeEduPaperCutResponse {
	s.Headers = v
	return s
}

func (s *RecognizeEduPaperCutResponse) SetStatusCode(v int32) *RecognizeEduPaperCutResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeEduPaperCutResponse) SetBody(v *RecognizeEduPaperCutResponseBody) *RecognizeEduPaperCutResponse {
	s.Body = v
	return s
}

type RecognizeEduPaperOcrRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// scan：扫描图， photo：实拍图
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// example:
	//
	// false
	OutputOricoord *bool `json:"OutputOricoord,omitempty" xml:"OutputOricoord,omitempty"`
	// example:
	//
	// default:默认, Math:数学, PrimarySchool_Math:小学数学, JHighSchool_Math: 初中数学, Chinese:语文, PrimarySchool_Chinese:小学语文, JHighSchool_Chinese:初中语文, English:英语, PrimarySchool_English:小学英语, JHighSchool_English:初中英语, Physics:物理, JHighSchool_Physics:初中物理, Chemistry: 化学, JHighSchool_Chemistry:初中化学, Biology:生物, JHighSchool_Biology:初中生物, History:历史, JHighSchool_History:初中历史, Geography:地理, JHighSchool_Geography:初中地理, Politics:政治, JHighSchool_Politics:初中政治
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduPaperOcrRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperOcrRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperOcrRequest) SetImageType(v string) *RecognizeEduPaperOcrRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeEduPaperOcrRequest) SetOutputOricoord(v bool) *RecognizeEduPaperOcrRequest {
	s.OutputOricoord = &v
	return s
}

func (s *RecognizeEduPaperOcrRequest) SetSubject(v string) *RecognizeEduPaperOcrRequest {
	s.Subject = &v
	return s
}

func (s *RecognizeEduPaperOcrRequest) SetUrl(v string) *RecognizeEduPaperOcrRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEduPaperOcrRequest) SetBody(v io.Reader) *RecognizeEduPaperOcrRequest {
	s.Body = v
	return s
}

type RecognizeEduPaperOcrResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeEduPaperOcrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperOcrResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperOcrResponseBody) SetCode(v string) *RecognizeEduPaperOcrResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduPaperOcrResponseBody) SetData(v string) *RecognizeEduPaperOcrResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduPaperOcrResponseBody) SetMessage(v string) *RecognizeEduPaperOcrResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeEduPaperOcrResponseBody) SetRequestId(v string) *RecognizeEduPaperOcrResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeEduPaperOcrResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeEduPaperOcrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduPaperOcrResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperOcrResponse) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperOcrResponse) SetHeaders(v map[string]*string) *RecognizeEduPaperOcrResponse {
	s.Headers = v
	return s
}

func (s *RecognizeEduPaperOcrResponse) SetStatusCode(v int32) *RecognizeEduPaperOcrResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeEduPaperOcrResponse) SetBody(v *RecognizeEduPaperOcrResponseBody) *RecognizeEduPaperOcrResponse {
	s.Body = v
	return s
}

type RecognizeEduPaperStructedRequest struct {
	// example:
	//
	// false
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// example:
	//
	// default:默认, Math:数学, PrimarySchool_Math:小学数学, JHighSchool_Math: 初中数学, Chinese:语文, PrimarySchool_Chinese:小学语文, //JHighSchool_Chinese:初中语文, English:英语, PrimarySchool_English:小学英语, JHighSchool_English:初中英语, Physics:物理, JHighSchool_Physics:初中物理   //Chemistry: 化学, JHighSchool_Chemistry:初中化学, Biology:生物, JHighSchool_Biology:初中生物, History:历史, JHighSchool_History:初中历史, Geography:地理,   //JHighSchool_Geography:初中地理, Politics:政治, JHighSchool_Politics:初中政治   "templateType": "Math"
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduPaperStructedRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperStructedRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperStructedRequest) SetNeedRotate(v bool) *RecognizeEduPaperStructedRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeEduPaperStructedRequest) SetSubject(v string) *RecognizeEduPaperStructedRequest {
	s.Subject = &v
	return s
}

func (s *RecognizeEduPaperStructedRequest) SetUrl(v string) *RecognizeEduPaperStructedRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEduPaperStructedRequest) SetBody(v io.Reader) *RecognizeEduPaperStructedRequest {
	s.Body = v
	return s
}

type RecognizeEduPaperStructedResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeEduPaperStructedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperStructedResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperStructedResponseBody) SetCode(v string) *RecognizeEduPaperStructedResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduPaperStructedResponseBody) SetData(v string) *RecognizeEduPaperStructedResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduPaperStructedResponseBody) SetMessage(v string) *RecognizeEduPaperStructedResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeEduPaperStructedResponseBody) SetRequestId(v string) *RecognizeEduPaperStructedResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeEduPaperStructedResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeEduPaperStructedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduPaperStructedResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduPaperStructedResponse) GoString() string {
	return s.String()
}

func (s *RecognizeEduPaperStructedResponse) SetHeaders(v map[string]*string) *RecognizeEduPaperStructedResponse {
	s.Headers = v
	return s
}

func (s *RecognizeEduPaperStructedResponse) SetStatusCode(v int32) *RecognizeEduPaperStructedResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeEduPaperStructedResponse) SetBody(v *RecognizeEduPaperStructedResponseBody) *RecognizeEduPaperStructedResponse {
	s.Body = v
	return s
}

type RecognizeEduQuestionOcrRequest struct {
	// example:
	//
	// false
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduQuestionOcrRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduQuestionOcrRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEduQuestionOcrRequest) SetNeedRotate(v bool) *RecognizeEduQuestionOcrRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeEduQuestionOcrRequest) SetUrl(v string) *RecognizeEduQuestionOcrRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEduQuestionOcrRequest) SetBody(v io.Reader) *RecognizeEduQuestionOcrRequest {
	s.Body = v
	return s
}

type RecognizeEduQuestionOcrResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeEduQuestionOcrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduQuestionOcrResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEduQuestionOcrResponseBody) SetCode(v string) *RecognizeEduQuestionOcrResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEduQuestionOcrResponseBody) SetData(v string) *RecognizeEduQuestionOcrResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEduQuestionOcrResponseBody) SetMessage(v string) *RecognizeEduQuestionOcrResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeEduQuestionOcrResponseBody) SetRequestId(v string) *RecognizeEduQuestionOcrResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeEduQuestionOcrResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeEduQuestionOcrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEduQuestionOcrResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEduQuestionOcrResponse) GoString() string {
	return s.String()
}

func (s *RecognizeEduQuestionOcrResponse) SetHeaders(v map[string]*string) *RecognizeEduQuestionOcrResponse {
	s.Headers = v
	return s
}

func (s *RecognizeEduQuestionOcrResponse) SetStatusCode(v int32) *RecognizeEduQuestionOcrResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeEduQuestionOcrResponse) SetBody(v *RecognizeEduQuestionOcrResponseBody) *RecognizeEduQuestionOcrResponse {
	s.Body = v
	return s
}

type RecognizeEnglishRequest struct {
	// example:
	//
	// false
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// example:
	//
	// false
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEnglishRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEnglishRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEnglishRequest) SetNeedRotate(v bool) *RecognizeEnglishRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeEnglishRequest) SetOutputTable(v bool) *RecognizeEnglishRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeEnglishRequest) SetUrl(v string) *RecognizeEnglishRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEnglishRequest) SetBody(v io.Reader) *RecognizeEnglishRequest {
	s.Body = v
	return s
}

type RecognizeEnglishResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeEnglishResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEnglishResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEnglishResponseBody) SetCode(v string) *RecognizeEnglishResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEnglishResponseBody) SetData(v string) *RecognizeEnglishResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEnglishResponseBody) SetMessage(v string) *RecognizeEnglishResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeEnglishResponseBody) SetRequestId(v string) *RecognizeEnglishResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeEnglishResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeEnglishResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEnglishResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEnglishResponse) GoString() string {
	return s.String()
}

func (s *RecognizeEnglishResponse) SetHeaders(v map[string]*string) *RecognizeEnglishResponse {
	s.Headers = v
	return s
}

func (s *RecognizeEnglishResponse) SetStatusCode(v int32) *RecognizeEnglishResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeEnglishResponse) SetBody(v *RecognizeEnglishResponseBody) *RecognizeEnglishResponse {
	s.Body = v
	return s
}

type RecognizeEstateCertificationRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1idy2XDZmx1VjSZFGXXax2XXa-713-1133.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEstateCertificationRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEstateCertificationRequest) GoString() string {
	return s.String()
}

func (s *RecognizeEstateCertificationRequest) SetUrl(v string) *RecognizeEstateCertificationRequest {
	s.Url = &v
	return s
}

func (s *RecognizeEstateCertificationRequest) SetBody(v io.Reader) *RecognizeEstateCertificationRequest {
	s.Body = v
	return s
}

type RecognizeEstateCertificationResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeEstateCertificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEstateCertificationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeEstateCertificationResponseBody) SetCode(v string) *RecognizeEstateCertificationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeEstateCertificationResponseBody) SetData(v string) *RecognizeEstateCertificationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeEstateCertificationResponseBody) SetMessage(v string) *RecognizeEstateCertificationResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeEstateCertificationResponseBody) SetRequestId(v string) *RecognizeEstateCertificationResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeEstateCertificationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeEstateCertificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeEstateCertificationResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeEstateCertificationResponse) GoString() string {
	return s.String()
}

func (s *RecognizeEstateCertificationResponse) SetHeaders(v map[string]*string) *RecognizeEstateCertificationResponse {
	s.Headers = v
	return s
}

func (s *RecognizeEstateCertificationResponse) SetStatusCode(v int32) *RecognizeEstateCertificationResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeEstateCertificationResponse) SetBody(v *RecognizeEstateCertificationResponseBody) *RecognizeEstateCertificationResponse {
	s.Body = v
	return s
}

type RecognizeExitEntryPermitToHKRequest struct {
	// example:
	//
	// true/false
	OutputFigure *bool `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i2/O1CN01Rs4C321G2oTD7Dg1U_!!6000000000565-0-tps-1024-692.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeExitEntryPermitToHKRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExitEntryPermitToHKRequest) GoString() string {
	return s.String()
}

func (s *RecognizeExitEntryPermitToHKRequest) SetOutputFigure(v bool) *RecognizeExitEntryPermitToHKRequest {
	s.OutputFigure = &v
	return s
}

func (s *RecognizeExitEntryPermitToHKRequest) SetUrl(v string) *RecognizeExitEntryPermitToHKRequest {
	s.Url = &v
	return s
}

func (s *RecognizeExitEntryPermitToHKRequest) SetBody(v io.Reader) *RecognizeExitEntryPermitToHKRequest {
	s.Body = v
	return s
}

type RecognizeExitEntryPermitToHKResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"data":{"permitType":"往来港澳通行证","nameCn":"朱伟","nameEn":"ZHU,WEI","birthDate":"2021.01.01","sex":"男","validPeriod":"2018.06.11-2028.06.10","issueAuthority":"公安部出入境管理局","issuePlace":"江苏","permitNumber":"C88600000","mrzCode":"CSC886084772<2800800<8200000<6"},"figure":[{"type":"face","x":160,"y":271,"w":190,"h":248,"box":{"x":254,"y":394,"w":186,"h":244,"angle":0},"points":[{"x":160,"y":272},{"x":347,"y":271},{"x":348,"y":516},{"x":161,"y":517}]},{"type":"face","x":711,"y":355,"w":80,"h":103,"box":{"x":750,"y":405,"w":75,"h":99,"angle":-1},"points":[{"x":711,"y":357},{"x":787,"y":355},{"x":789,"y":454},{"x":713,"y":456}]}],"ftype":0,"height":692,"orgHeight":692,"orgWidth":1024,"prism_keyValueInfo":[{"key":"permitType","keyProb":100,"value":"往来港澳通行证","valuePos":[{"x":142,"y":39},{"x":476,"y":35},{"x":477,"y":75},{"x":142,"y":79}],"valueProb":100},{"key":"nameCn","keyProb":100,"value":"朱伟","valuePos":[{"x":272,"y":126},{"x":346,"y":124},{"x":347,"y":160},{"x":272,"y":161}],"valueProb":100},{"key":"nameEn","keyProb":100,"value":"ZHU,WEI","valuePos":[{"x":273,"y":168},{"x":403,"y":167},{"x":403,"y":194},{"x":274,"y":196}],"valueProb":100},{"key":"birthDate","keyProb":100,"value":"2021.01.01","valuePos":[{"x":421,"y":240},{"x":421,"y":269},{"x":281,"y":269},{"x":281,"y":240}],"valueProb":100},{"key":"sex","keyProb":100,"value":"男","valuePos":[{"x":502,"y":240},{"x":502,"y":270},{"x":474,"y":270},{"x":474,"y":240}],"valueProb":100},{"key":"validPeriod","keyProb":100,"value":"2018.06.11-2028.06.10","valuePos":[{"x":579,"y":301},{"x":579,"y":328},{"x":275,"y":328},{"x":275,"y":301}],"valueProb":100},{"key":"issueAuthority","keyProb":100,"value":"公安部出入境管理局","valuePos":[{"x":278,"y":361},{"x":524,"y":361},{"x":524,"y":391},{"x":278,"y":391}],"valueProb":100},{"key":"issuePlace","keyProb":100,"value":"江苏","valuePos":[{"x":619,"y":361},{"x":619,"y":391},{"x":561,"y":391},{"x":561,"y":361}],"valueProb":100},{"key":"permitNumber","keyProb":100,"value":"C88600000","valuePos":[{"x":524,"y":61},{"x":727,"y":60},{"x":728,"y":92},{"x":524,"y":94}],"valueProb":100},{"key":"mrzCode","keyProb":98,"value":"CSC886084772<2800800<8200000<6","valuePos":[{"x":714,"y":421},{"x":714,"y":449},{"x":65,"y":449},{"x":65,"y":421}],"valueProb":98}],"sliceRect":{"x0":107,"y0":135,"x1":880,"y1":134,"x2":874,"y2":616,"x3":117,"y3":624},"width":1024}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C99EABB8-9FCB-5E5E-B4D9-AFCFA6C8B3FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeExitEntryPermitToHKResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExitEntryPermitToHKResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeExitEntryPermitToHKResponseBody) SetCode(v string) *RecognizeExitEntryPermitToHKResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeExitEntryPermitToHKResponseBody) SetData(v string) *RecognizeExitEntryPermitToHKResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeExitEntryPermitToHKResponseBody) SetMessage(v string) *RecognizeExitEntryPermitToHKResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeExitEntryPermitToHKResponseBody) SetRequestId(v string) *RecognizeExitEntryPermitToHKResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeExitEntryPermitToHKResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeExitEntryPermitToHKResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeExitEntryPermitToHKResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExitEntryPermitToHKResponse) GoString() string {
	return s.String()
}

func (s *RecognizeExitEntryPermitToHKResponse) SetHeaders(v map[string]*string) *RecognizeExitEntryPermitToHKResponse {
	s.Headers = v
	return s
}

func (s *RecognizeExitEntryPermitToHKResponse) SetStatusCode(v int32) *RecognizeExitEntryPermitToHKResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeExitEntryPermitToHKResponse) SetBody(v *RecognizeExitEntryPermitToHKResponseBody) *RecognizeExitEntryPermitToHKResponse {
	s.Body = v
	return s
}

type RecognizeExitEntryPermitToMainlandRequest struct {
	// example:
	//
	// true/false
	OutputFigure *bool `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i2/O1CN01VpucoK1PtmovU859J_!!6000000001899-0-tps-928-626.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeExitEntryPermitToMainlandRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExitEntryPermitToMainlandRequest) GoString() string {
	return s.String()
}

func (s *RecognizeExitEntryPermitToMainlandRequest) SetOutputFigure(v bool) *RecognizeExitEntryPermitToMainlandRequest {
	s.OutputFigure = &v
	return s
}

func (s *RecognizeExitEntryPermitToMainlandRequest) SetUrl(v string) *RecognizeExitEntryPermitToMainlandRequest {
	s.Url = &v
	return s
}

func (s *RecognizeExitEntryPermitToMainlandRequest) SetBody(v io.Reader) *RecognizeExitEntryPermitToMainlandRequest {
	s.Body = v
	return s
}

type RecognizeExitEntryPermitToMainlandResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"data":{"permitType":"港澳居民来往内地通行证","nameCn":"何郑","nameEn":"HE,CHENG","birthDate":"2000.01.01","sex":"男","validPeriod":"2014.04.10-2019.04.09","issueAuthority":"公安部出入境管理局","issuePlace":"","permitNumber":"H10387877","issueCount":"01"},"figure":[{"type":"face","x":80,"y":164,"w":192,"h":273,"box":{"x":175,"y":300,"w":187,"h":269,"angle":0},"points":[{"x":80,"y":166},{"x":268,"y":164},{"x":270,"y":433},{"x":82,"y":435}]}],"ftype":0,"height":626,"orgHeight":626,"orgWidth":928,"prism_keyValueInfo":[{"key":"permitType","keyProb":100,"value":"港澳居民来往内地通行证","valuePos":[{"x":680,"y":41},{"x":681,"y":83},{"x":177,"y":86},{"x":176,"y":44}],"valueProb":100},{"key":"nameCn","keyProb":100,"value":"何郑","valuePos":[{"x":346,"y":119},{"x":346,"y":153},{"x":269,"y":153},{"x":269,"y":119}],"valueProb":100},{"key":"nameEn","keyProb":100,"value":"HE,CHENG","valuePos":[{"x":452,"y":166},{"x":452,"y":195},{"x":270,"y":195},{"x":270,"y":166}],"valueProb":100},{"key":"birthDate","keyProb":100,"value":"2000.01.01","valuePos":[{"x":273,"y":226},{"x":414,"y":226},{"x":414,"y":254},{"x":273,"y":254}],"valueProb":100},{"key":"sex","keyProb":100,"value":"男","valuePos":[{"x":594,"y":234},{"x":594,"y":268},{"x":562,"y":268},{"x":562,"y":234}],"valueProb":100},{"key":"validPeriod","keyProb":100,"value":"2014.04.10-2019.04.09","valuePos":[{"x":700,"y":295},{"x":700,"y":323},{"x":267,"y":324},{"x":267,"y":296}],"valueProb":100},{"key":"issueAuthority","keyProb":100,"value":"公安部出入境管理局","valuePos":[{"x":264,"y":386},{"x":265,"y":353},{"x":536,"y":357},{"x":536,"y":390}],"valueProb":100},{"key":"issuePlace","keyProb":100,"value":"","valueProb":100},{"key":"permitNumber","keyProb":100,"value":"H10387877","valuePos":[{"x":489,"y":424},{"x":489,"y":457},{"x":268,"y":457},{"x":268,"y":424}],"valueProb":100},{"key":"issueCount","keyProb":100,"value":"01","valuePos":[{"x":601,"y":425},{"x":601,"y":456},{"x":555,"y":456},{"x":555,"y":425}],"valueProb":100}],"sliceRect":{"x0":46,"y0":30,"x1":887,"y1":38,"x2":892,"y2":564,"x3":39,"y3":567},"width":928}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeExitEntryPermitToMainlandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExitEntryPermitToMainlandResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeExitEntryPermitToMainlandResponseBody) SetCode(v string) *RecognizeExitEntryPermitToMainlandResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeExitEntryPermitToMainlandResponseBody) SetData(v string) *RecognizeExitEntryPermitToMainlandResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeExitEntryPermitToMainlandResponseBody) SetMessage(v string) *RecognizeExitEntryPermitToMainlandResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeExitEntryPermitToMainlandResponseBody) SetRequestId(v string) *RecognizeExitEntryPermitToMainlandResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeExitEntryPermitToMainlandResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeExitEntryPermitToMainlandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeExitEntryPermitToMainlandResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeExitEntryPermitToMainlandResponse) GoString() string {
	return s.String()
}

func (s *RecognizeExitEntryPermitToMainlandResponse) SetHeaders(v map[string]*string) *RecognizeExitEntryPermitToMainlandResponse {
	s.Headers = v
	return s
}

func (s *RecognizeExitEntryPermitToMainlandResponse) SetStatusCode(v int32) *RecognizeExitEntryPermitToMainlandResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeExitEntryPermitToMainlandResponse) SetBody(v *RecognizeExitEntryPermitToMainlandResponseBody) *RecognizeExitEntryPermitToMainlandResponse {
	s.Body = v
	return s
}

type RecognizeFoodManageLicenseRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeFoodManageLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodManageLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFoodManageLicenseRequest) SetUrl(v string) *RecognizeFoodManageLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeFoodManageLicenseRequest) SetBody(v io.Reader) *RecognizeFoodManageLicenseRequest {
	s.Body = v
	return s
}

type RecognizeFoodManageLicenseResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeFoodManageLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodManageLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeFoodManageLicenseResponseBody) SetCode(v string) *RecognizeFoodManageLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeFoodManageLicenseResponseBody) SetData(v string) *RecognizeFoodManageLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeFoodManageLicenseResponseBody) SetMessage(v string) *RecognizeFoodManageLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeFoodManageLicenseResponseBody) SetRequestId(v string) *RecognizeFoodManageLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeFoodManageLicenseResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeFoodManageLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeFoodManageLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodManageLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFoodManageLicenseResponse) SetHeaders(v map[string]*string) *RecognizeFoodManageLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeFoodManageLicenseResponse) SetStatusCode(v int32) *RecognizeFoodManageLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeFoodManageLicenseResponse) SetBody(v *RecognizeFoodManageLicenseResponseBody) *RecognizeFoodManageLicenseResponse {
	s.Body = v
	return s
}

type RecognizeFoodProduceLicenseRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1YaMhXKT2gK0jSZFvXXXnFXXa-1414-1000.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeFoodProduceLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodProduceLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFoodProduceLicenseRequest) SetUrl(v string) *RecognizeFoodProduceLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeFoodProduceLicenseRequest) SetBody(v io.Reader) *RecognizeFoodProduceLicenseRequest {
	s.Body = v
	return s
}

type RecognizeFoodProduceLicenseResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeFoodProduceLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodProduceLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeFoodProduceLicenseResponseBody) SetCode(v string) *RecognizeFoodProduceLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeFoodProduceLicenseResponseBody) SetData(v string) *RecognizeFoodProduceLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeFoodProduceLicenseResponseBody) SetMessage(v string) *RecognizeFoodProduceLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeFoodProduceLicenseResponseBody) SetRequestId(v string) *RecognizeFoodProduceLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeFoodProduceLicenseResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeFoodProduceLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeFoodProduceLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodProduceLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFoodProduceLicenseResponse) SetHeaders(v map[string]*string) *RecognizeFoodProduceLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeFoodProduceLicenseResponse) SetStatusCode(v int32) *RecognizeFoodProduceLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeFoodProduceLicenseResponse) SetBody(v *RecognizeFoodProduceLicenseResponseBody) *RecognizeFoodProduceLicenseResponse {
	s.Body = v
	return s
}

type RecognizeGeneralRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeGeneralRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeGeneralRequest) GoString() string {
	return s.String()
}

func (s *RecognizeGeneralRequest) SetUrl(v string) *RecognizeGeneralRequest {
	s.Url = &v
	return s
}

func (s *RecognizeGeneralRequest) SetBody(v io.Reader) *RecognizeGeneralRequest {
	s.Body = v
	return s
}

type RecognizeGeneralResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeGeneralResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeGeneralResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeGeneralResponseBody) SetCode(v string) *RecognizeGeneralResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeGeneralResponseBody) SetData(v string) *RecognizeGeneralResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeGeneralResponseBody) SetMessage(v string) *RecognizeGeneralResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeGeneralResponseBody) SetRequestId(v string) *RecognizeGeneralResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeGeneralResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeGeneralResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeGeneralResponse) GoString() string {
	return s.String()
}

func (s *RecognizeGeneralResponse) SetHeaders(v map[string]*string) *RecognizeGeneralResponse {
	s.Headers = v
	return s
}

func (s *RecognizeGeneralResponse) SetStatusCode(v int32) *RecognizeGeneralResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeGeneralResponse) SetBody(v *RecognizeGeneralResponseBody) *RecognizeGeneralResponse {
	s.Body = v
	return s
}

type RecognizeHKIdcardRequest struct {
	// example:
	//
	// https://example.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeHKIdcardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHKIdcardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeHKIdcardRequest) SetUrl(v string) *RecognizeHKIdcardRequest {
	s.Url = &v
	return s
}

func (s *RecognizeHKIdcardRequest) SetBody(v io.Reader) *RecognizeHKIdcardRequest {
	s.Body = v
	return s
}

type RecognizeHKIdcardResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeHKIdcardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHKIdcardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeHKIdcardResponseBody) SetCode(v string) *RecognizeHKIdcardResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeHKIdcardResponseBody) SetData(v string) *RecognizeHKIdcardResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeHKIdcardResponseBody) SetMessage(v string) *RecognizeHKIdcardResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeHKIdcardResponseBody) SetRequestId(v string) *RecognizeHKIdcardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeHKIdcardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeHKIdcardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeHKIdcardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHKIdcardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeHKIdcardResponse) SetHeaders(v map[string]*string) *RecognizeHKIdcardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeHKIdcardResponse) SetStatusCode(v int32) *RecognizeHKIdcardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeHKIdcardResponse) SetBody(v *RecognizeHKIdcardResponseBody) *RecognizeHKIdcardResponse {
	s.Body = v
	return s
}

type RecognizeHandwritingRequest struct {
	// example:
	//
	// false
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// example:
	//
	// false
	NeedSortPage *bool `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
	// example:
	//
	// false
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// example:
	//
	// false
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	Paragraph   *bool `json:"Paragraph,omitempty" xml:"Paragraph,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeHandwritingRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHandwritingRequest) GoString() string {
	return s.String()
}

func (s *RecognizeHandwritingRequest) SetNeedRotate(v bool) *RecognizeHandwritingRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeHandwritingRequest) SetNeedSortPage(v bool) *RecognizeHandwritingRequest {
	s.NeedSortPage = &v
	return s
}

func (s *RecognizeHandwritingRequest) SetOutputCharInfo(v bool) *RecognizeHandwritingRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeHandwritingRequest) SetOutputTable(v bool) *RecognizeHandwritingRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeHandwritingRequest) SetParagraph(v bool) *RecognizeHandwritingRequest {
	s.Paragraph = &v
	return s
}

func (s *RecognizeHandwritingRequest) SetUrl(v string) *RecognizeHandwritingRequest {
	s.Url = &v
	return s
}

func (s *RecognizeHandwritingRequest) SetBody(v io.Reader) *RecognizeHandwritingRequest {
	s.Body = v
	return s
}

type RecognizeHandwritingResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeHandwritingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHandwritingResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeHandwritingResponseBody) SetCode(v string) *RecognizeHandwritingResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeHandwritingResponseBody) SetData(v string) *RecognizeHandwritingResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeHandwritingResponseBody) SetMessage(v string) *RecognizeHandwritingResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeHandwritingResponseBody) SetRequestId(v string) *RecognizeHandwritingResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeHandwritingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeHandwritingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeHandwritingResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHandwritingResponse) GoString() string {
	return s.String()
}

func (s *RecognizeHandwritingResponse) SetHeaders(v map[string]*string) *RecognizeHandwritingResponse {
	s.Headers = v
	return s
}

func (s *RecognizeHandwritingResponse) SetStatusCode(v int32) *RecognizeHandwritingResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeHandwritingResponse) SetBody(v *RecognizeHandwritingResponseBody) *RecognizeHandwritingResponse {
	s.Body = v
	return s
}

type RecognizeHealthCodeRequest struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i3/O1CN01ME0L7j29f6VRZKo5e_!!6000000008094-0-tps-1237-1981.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeHealthCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHealthCodeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeHealthCodeRequest) SetUrl(v string) *RecognizeHealthCodeRequest {
	s.Url = &v
	return s
}

func (s *RecognizeHealthCodeRequest) SetBody(v io.Reader) *RecognizeHealthCodeRequest {
	s.Body = v
	return s
}

type RecognizeHealthCodeResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"data":{"permitType":"往来港澳通行证","nameCn":"朱伟","nameEn":"ZHU,WEI","birthDate":"2021.01.01","sex":"男","validPeriod":"2018.06.11-2028.06.10","issueAuthority":"公安部出入境管理局","issuePlace":"江苏","permitNumber":"C88600000","mrzCode":"CSC886084772<2800800<8200000<6"},"figure":[{"type":"face","x":160,"y":271,"w":190,"h":248,"box":{"x":254,"y":394,"w":186,"h":244,"angle":0},"points":[{"x":160,"y":272},{"x":347,"y":271},{"x":348,"y":516},{"x":161,"y":517}]},{"type":"face","x":711,"y":355,"w":80,"h":103,"box":{"x":750,"y":405,"w":75,"h":99,"angle":-1},"points":[{"x":711,"y":357},{"x":787,"y":355},{"x":789,"y":454},{"x":713,"y":456}]}],"ftype":0,"height":692,"orgHeight":692,"orgWidth":1024,"prism_keyValueInfo":[{"key":"permitType","keyProb":100,"value":"往来港澳通行证","valuePos":[{"x":142,"y":39},{"x":476,"y":35},{"x":477,"y":75},{"x":142,"y":79}],"valueProb":100},{"key":"nameCn","keyProb":100,"value":"朱伟","valuePos":[{"x":272,"y":126},{"x":346,"y":124},{"x":347,"y":160},{"x":272,"y":161}],"valueProb":100},{"key":"nameEn","keyProb":100,"value":"ZHU,WEI","valuePos":[{"x":273,"y":168},{"x":403,"y":167},{"x":403,"y":194},{"x":274,"y":196}],"valueProb":100},{"key":"birthDate","keyProb":100,"value":"2021.01.01","valuePos":[{"x":421,"y":240},{"x":421,"y":269},{"x":281,"y":269},{"x":281,"y":240}],"valueProb":100},{"key":"sex","keyProb":100,"value":"男","valuePos":[{"x":502,"y":240},{"x":502,"y":270},{"x":474,"y":270},{"x":474,"y":240}],"valueProb":100},{"key":"validPeriod","keyProb":100,"value":"2018.06.11-2028.06.10","valuePos":[{"x":579,"y":301},{"x":579,"y":328},{"x":275,"y":328},{"x":275,"y":301}],"valueProb":100},{"key":"issueAuthority","keyProb":100,"value":"公安部出入境管理局","valuePos":[{"x":278,"y":361},{"x":524,"y":361},{"x":524,"y":391},{"x":278,"y":391}],"valueProb":100},{"key":"issuePlace","keyProb":100,"value":"江苏","valuePos":[{"x":619,"y":361},{"x":619,"y":391},{"x":561,"y":391},{"x":561,"y":361}],"valueProb":100},{"key":"permitNumber","keyProb":100,"value":"C88600000","valuePos":[{"x":524,"y":61},{"x":727,"y":60},{"x":728,"y":92},{"x":524,"y":94}],"valueProb":100},{"key":"mrzCode","keyProb":98,"value":"CSC886084772<2800800<8200000<6","valuePos":[{"x":714,"y":421},{"x":714,"y":449},{"x":65,"y":449},{"x":65,"y":421}],"valueProb":98}],"sliceRect":{"x0":107,"y0":135,"x1":880,"y1":134,"x2":874,"y2":616,"x3":117,"y3":624},"width":1024}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeHealthCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHealthCodeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeHealthCodeResponseBody) SetCode(v string) *RecognizeHealthCodeResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeHealthCodeResponseBody) SetData(v string) *RecognizeHealthCodeResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeHealthCodeResponseBody) SetMessage(v string) *RecognizeHealthCodeResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeHealthCodeResponseBody) SetRequestId(v string) *RecognizeHealthCodeResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeHealthCodeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeHealthCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeHealthCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHealthCodeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeHealthCodeResponse) SetHeaders(v map[string]*string) *RecognizeHealthCodeResponse {
	s.Headers = v
	return s
}

func (s *RecognizeHealthCodeResponse) SetStatusCode(v int32) *RecognizeHealthCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeHealthCodeResponse) SetBody(v *RecognizeHealthCodeResponseBody) *RecognizeHealthCodeResponse {
	s.Body = v
	return s
}

type RecognizeHotelConsumeRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeHotelConsumeRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHotelConsumeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeHotelConsumeRequest) SetUrl(v string) *RecognizeHotelConsumeRequest {
	s.Url = &v
	return s
}

func (s *RecognizeHotelConsumeRequest) SetBody(v io.Reader) *RecognizeHotelConsumeRequest {
	s.Body = v
	return s
}

type RecognizeHotelConsumeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeHotelConsumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHotelConsumeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeHotelConsumeResponseBody) SetCode(v string) *RecognizeHotelConsumeResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeHotelConsumeResponseBody) SetData(v string) *RecognizeHotelConsumeResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeHotelConsumeResponseBody) SetMessage(v string) *RecognizeHotelConsumeResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeHotelConsumeResponseBody) SetRequestId(v string) *RecognizeHotelConsumeResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeHotelConsumeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeHotelConsumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeHotelConsumeResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHotelConsumeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeHotelConsumeResponse) SetHeaders(v map[string]*string) *RecognizeHotelConsumeResponse {
	s.Headers = v
	return s
}

func (s *RecognizeHotelConsumeResponse) SetStatusCode(v int32) *RecognizeHotelConsumeResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeHotelConsumeResponse) SetBody(v *RecognizeHotelConsumeResponseBody) *RecognizeHotelConsumeResponse {
	s.Body = v
	return s
}

type RecognizeHouseholdRequest struct {
	// example:
	//
	// false
	IsResidentPage *bool `json:"IsResidentPage,omitempty" xml:"IsResidentPage,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB11ZxTMxD1gK0jSZFsXXbldVXa-920-606.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeHouseholdRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHouseholdRequest) GoString() string {
	return s.String()
}

func (s *RecognizeHouseholdRequest) SetIsResidentPage(v bool) *RecognizeHouseholdRequest {
	s.IsResidentPage = &v
	return s
}

func (s *RecognizeHouseholdRequest) SetUrl(v string) *RecognizeHouseholdRequest {
	s.Url = &v
	return s
}

func (s *RecognizeHouseholdRequest) SetBody(v io.Reader) *RecognizeHouseholdRequest {
	s.Body = v
	return s
}

type RecognizeHouseholdResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeHouseholdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHouseholdResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeHouseholdResponseBody) SetCode(v string) *RecognizeHouseholdResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeHouseholdResponseBody) SetData(v string) *RecognizeHouseholdResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeHouseholdResponseBody) SetMessage(v string) *RecognizeHouseholdResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeHouseholdResponseBody) SetRequestId(v string) *RecognizeHouseholdResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeHouseholdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeHouseholdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeHouseholdResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeHouseholdResponse) GoString() string {
	return s.String()
}

func (s *RecognizeHouseholdResponse) SetHeaders(v map[string]*string) *RecognizeHouseholdResponse {
	s.Headers = v
	return s
}

func (s *RecognizeHouseholdResponse) SetStatusCode(v int32) *RecognizeHouseholdResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeHouseholdResponse) SetBody(v *RecognizeHouseholdResponseBody) *RecognizeHouseholdResponse {
	s.Body = v
	return s
}

type RecognizeIdcardRequest struct {
	// example:
	//
	// false
	OutputFigure *bool `json:"OutputFigure,omitempty" xml:"OutputFigure,omitempty"`
	// example:
	//
	// false
	OutputQualityInfo *bool `json:"OutputQualityInfo,omitempty" xml:"OutputQualityInfo,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1q5IeXAvoK1RjSZFNXXcxMVXa-483-307.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeIdcardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdcardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeIdcardRequest) SetOutputFigure(v bool) *RecognizeIdcardRequest {
	s.OutputFigure = &v
	return s
}

func (s *RecognizeIdcardRequest) SetOutputQualityInfo(v bool) *RecognizeIdcardRequest {
	s.OutputQualityInfo = &v
	return s
}

func (s *RecognizeIdcardRequest) SetUrl(v string) *RecognizeIdcardRequest {
	s.Url = &v
	return s
}

func (s *RecognizeIdcardRequest) SetBody(v io.Reader) *RecognizeIdcardRequest {
	s.Body = v
	return s
}

type RecognizeIdcardResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86B83935-DD36-195B-B6E4-D07BE370C8B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeIdcardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdcardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeIdcardResponseBody) SetCode(v string) *RecognizeIdcardResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeIdcardResponseBody) SetData(v string) *RecognizeIdcardResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeIdcardResponseBody) SetMessage(v string) *RecognizeIdcardResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeIdcardResponseBody) SetRequestId(v string) *RecognizeIdcardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeIdcardResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeIdcardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeIdcardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdcardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeIdcardResponse) SetHeaders(v map[string]*string) *RecognizeIdcardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeIdcardResponse) SetStatusCode(v int32) *RecognizeIdcardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeIdcardResponse) SetBody(v *RecognizeIdcardResponseBody) *RecognizeIdcardResponse {
	s.Body = v
	return s
}

type RecognizeInternationalBusinessLicenseRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// India
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// https://www.example.com
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeInternationalBusinessLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInternationalBusinessLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeInternationalBusinessLicenseRequest) SetCountry(v string) *RecognizeInternationalBusinessLicenseRequest {
	s.Country = &v
	return s
}

func (s *RecognizeInternationalBusinessLicenseRequest) SetUrl(v string) *RecognizeInternationalBusinessLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeInternationalBusinessLicenseRequest) SetBody(v io.Reader) *RecognizeInternationalBusinessLicenseRequest {
	s.Body = v
	return s
}

type RecognizeInternationalBusinessLicenseResponseBody struct {
	// example:
	//
	// unmatchedImageType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"algo_version": "b16f86189b72c2d726865272c98e8e58156a41c7;b16f86189b72c2d726865272c98e8e58156a41c7", "data": {"certificateType": "간이과세자", "issuanceNo": "", "processingTime": "", "companyNameEn": "", "companyName": "", "registrationNo": "135-31-78773", "nameOfRepresentativeEn": "", "nameOfRepresentative": "테라", "residentRegistrationNo": "", "businessAddressEn": "", "businessAddress": "경기도 수원시 영통구 영통로 498, 143동 1806흐(영통동, 황골마을 주공아파트)", "businessCommencementDate": "1972-01-10", "businessRegistrationDate": "", "businessTypeEn": "", "businessType": "", "businessItemEn": "", "businessItem": "스매업 전자상거래업(의류)", "jointCompanyName": "", "jointCompanyRegistrationNo": "", "issueDate": "2015-10-28", "issuer": "동수원세무서장"}, "ftype": 0, "height": 2988, "orgHeight": 2988, "orgWidth": 5312, "prism_keyValueInfo": [{"key": "certificateType", "keyProb": 100, "value": "간이과세자", "valuePos": [{"x": 621, "y": 1768}, {"x": 615, "y": 1221}, {"x": 720, "y": 1220}, {"x": 726, "y": 1767}], "valueProb": 100}, {"key": "issuanceNo", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "processingTime", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "companyNameEn", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "companyName", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "registrationNo", "keyProb": 100, "value": "135-31-78773", "valuePos": [{"x": 773, "y": 1517}, {"x": 763, "y": 881}, {"x": 861, "y": 880}, {"x": 870, "y": 1515}], "valueProb": 100}, {"key": "nameOfRepresentativeEn", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "nameOfRepresentative", "keyProb": 90, "value": "테라", "valuePos": [{"x": 946, "y": 2201}, {"x": 946, "y": 2047}, {"x": 1022, "y": 2047}, {"x": 1022, "y": 2201}], "valueProb": 90}, {"key": "residentRegistrationNo", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "businessAddressEn", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "businessAddress", "keyProb": 96, "value": "경기도 수원시 영통구 영통로 498, 143동 1806흐(영통동, 황골마을 주공아파트)", "valuePos": [{"x": 1346, "y": 2200}, {"x": 1321, "y": 736}, {"x": 1499, "y": 733}, {"x": 1523, "y": 2197}], "valueProb": 96}, {"key": "businessCommencementDate", "keyProb": 100, "value": "1972-01-10", "valuePos": [{"x": 1055, "y": 788}, {"x": 1046, "y": 62}, {"x": 1127, "y": 62}, {"x": 1135, "y": 787}], "valueProb": 100}, {"key": "businessRegistrationDate", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "businessTypeEn", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "businessType", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "businessItemEn", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "businessItem", "keyProb": 100, "value": "스매업 전자상거래업(의류)", "valuePos": [{"x": 1590, "y": 1982}, {"x": 1561, "y": 293}, {"x": 1659, "y": 291}, {"x": 1688, "y": 1980}], "valueProb": 100}, {"key": "jointCompanyName", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "jointCompanyRegistrationNo", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "issueDate", "keyProb": 100, "value": "2015-10-28", "valuePos": [{"x": 3755, "y": 1938}, {"x": 3749, "y": 1057}, {"x": 3842, "y": 1056}, {"x": 3848, "y": 1937}], "valueProb": 100}, {"key": "issuer", "keyProb": 100, "value": "동수원세무서장", "valuePos": [{"x": 3978, "y": 1997}, {"x": 3970, "y": 982}, {"x": 4099, "y": 980}, {"x": 4107, "y": 1996}], "valueProb": 100}], "sliceRect": {"x0": 8, "y0": 0, "x1": 4695, "y1": 0, "x2": 4737, "y2": 2976, "x3": 12, "y3": 2988}, "width": 5312}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// The type of image didn\\"t match the api.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeInternationalBusinessLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInternationalBusinessLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeInternationalBusinessLicenseResponseBody) SetCode(v string) *RecognizeInternationalBusinessLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeInternationalBusinessLicenseResponseBody) SetData(v string) *RecognizeInternationalBusinessLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeInternationalBusinessLicenseResponseBody) SetMessage(v string) *RecognizeInternationalBusinessLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeInternationalBusinessLicenseResponseBody) SetRequestId(v string) *RecognizeInternationalBusinessLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeInternationalBusinessLicenseResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeInternationalBusinessLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeInternationalBusinessLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInternationalBusinessLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeInternationalBusinessLicenseResponse) SetHeaders(v map[string]*string) *RecognizeInternationalBusinessLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeInternationalBusinessLicenseResponse) SetStatusCode(v int32) *RecognizeInternationalBusinessLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeInternationalBusinessLicenseResponse) SetBody(v *RecognizeInternationalBusinessLicenseResponseBody) *RecognizeInternationalBusinessLicenseResponse {
	s.Body = v
	return s
}

type RecognizeInternationalIdcardRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Vietnam
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// http://example.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeInternationalIdcardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInternationalIdcardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeInternationalIdcardRequest) SetCountry(v string) *RecognizeInternationalIdcardRequest {
	s.Country = &v
	return s
}

func (s *RecognizeInternationalIdcardRequest) SetUrl(v string) *RecognizeInternationalIdcardRequest {
	s.Url = &v
	return s
}

func (s *RecognizeInternationalIdcardRequest) SetBody(v io.Reader) *RecognizeInternationalIdcardRequest {
	s.Body = v
	return s
}

type RecognizeInternationalIdcardResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeInternationalIdcardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInternationalIdcardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeInternationalIdcardResponseBody) SetCode(v string) *RecognizeInternationalIdcardResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeInternationalIdcardResponseBody) SetData(v string) *RecognizeInternationalIdcardResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeInternationalIdcardResponseBody) SetMessage(v string) *RecognizeInternationalIdcardResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeInternationalIdcardResponseBody) SetRequestId(v string) *RecognizeInternationalIdcardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeInternationalIdcardResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeInternationalIdcardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeInternationalIdcardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInternationalIdcardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeInternationalIdcardResponse) SetHeaders(v map[string]*string) *RecognizeInternationalIdcardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeInternationalIdcardResponse) SetStatusCode(v int32) *RecognizeInternationalIdcardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeInternationalIdcardResponse) SetBody(v *RecognizeInternationalIdcardResponseBody) *RecognizeInternationalIdcardResponse {
	s.Body = v
	return s
}

type RecognizeInvoiceRequest struct {
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1qIIfXAPoK1RjSZKbXXX1IXXa-808-523.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeInvoiceRequest) SetPageNo(v int32) *RecognizeInvoiceRequest {
	s.PageNo = &v
	return s
}

func (s *RecognizeInvoiceRequest) SetUrl(v string) *RecognizeInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeInvoiceRequest) SetBody(v io.Reader) *RecognizeInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeInvoiceResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeInvoiceResponseBody) SetCode(v string) *RecognizeInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeInvoiceResponseBody) SetData(v string) *RecognizeInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeInvoiceResponseBody) SetMessage(v string) *RecognizeInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeInvoiceResponseBody) SetRequestId(v string) *RecognizeInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeInvoiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeInvoiceResponse) SetStatusCode(v int32) *RecognizeInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeInvoiceResponse) SetBody(v *RecognizeInvoiceResponseBody) *RecognizeInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeJanpaneseRequest struct {
	// example:
	//
	// false
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// example:
	//
	// false
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// example:
	//
	// false
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeJanpaneseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeJanpaneseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeJanpaneseRequest) SetNeedRotate(v bool) *RecognizeJanpaneseRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeJanpaneseRequest) SetOutputCharInfo(v bool) *RecognizeJanpaneseRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeJanpaneseRequest) SetOutputTable(v bool) *RecognizeJanpaneseRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeJanpaneseRequest) SetUrl(v string) *RecognizeJanpaneseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeJanpaneseRequest) SetBody(v io.Reader) *RecognizeJanpaneseRequest {
	s.Body = v
	return s
}

type RecognizeJanpaneseResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeJanpaneseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeJanpaneseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeJanpaneseResponseBody) SetCode(v string) *RecognizeJanpaneseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeJanpaneseResponseBody) SetData(v string) *RecognizeJanpaneseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeJanpaneseResponseBody) SetMessage(v string) *RecognizeJanpaneseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeJanpaneseResponseBody) SetRequestId(v string) *RecognizeJanpaneseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeJanpaneseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeJanpaneseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeJanpaneseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeJanpaneseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeJanpaneseResponse) SetHeaders(v map[string]*string) *RecognizeJanpaneseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeJanpaneseResponse) SetStatusCode(v int32) *RecognizeJanpaneseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeJanpaneseResponse) SetBody(v *RecognizeJanpaneseResponseBody) *RecognizeJanpaneseResponse {
	s.Body = v
	return s
}

type RecognizeKoreanRequest struct {
	// example:
	//
	// false
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// example:
	//
	// false
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// example:
	//
	// false
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeKoreanRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeKoreanRequest) GoString() string {
	return s.String()
}

func (s *RecognizeKoreanRequest) SetNeedRotate(v bool) *RecognizeKoreanRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeKoreanRequest) SetOutputCharInfo(v bool) *RecognizeKoreanRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeKoreanRequest) SetOutputTable(v bool) *RecognizeKoreanRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeKoreanRequest) SetUrl(v string) *RecognizeKoreanRequest {
	s.Url = &v
	return s
}

func (s *RecognizeKoreanRequest) SetBody(v io.Reader) *RecognizeKoreanRequest {
	s.Body = v
	return s
}

type RecognizeKoreanResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"content":"위 기자재는 [전파법] 제58조의2 제3항에 따라 등록되었음을 증명합니다.  Itis verified thatforegoing equipment has bee en registered underthe Clause 3, Article 58-2 of Radio Waves Act.  y0 13년(Year)_08월(Month) 16일(Date) 국립전 파연구 국립전파연7 구원장 인 Dlrector General ofNatlonal Radio Research Agency    적합등록 방송통신기자재는 반드시\\"적합성평가표: .시\\"를 부착하여 유통하여야 합니다.  위반시 과태료 처분 및등록이 취소될 수 있습니다.  ","height":499,"orgHeight":499,"orgWidth":1153,"prism_version":"1.0.9","prism_wnum":19,"prism_wordsInfo":[{"angle":-90,"direction":0,"height":587,"pos":[{"x":61,"y":18},{"x":647,"y":16},{"x":647,"y":43},{"x":61,"y":45}],"prob":98,"width":27,"word":"위 기자재는 [전파법] 제58조의2 제3항에 따라","x":341,"y":-263}],"width":1153}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeKoreanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeKoreanResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeKoreanResponseBody) SetCode(v string) *RecognizeKoreanResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeKoreanResponseBody) SetData(v string) *RecognizeKoreanResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeKoreanResponseBody) SetMessage(v string) *RecognizeKoreanResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeKoreanResponseBody) SetRequestId(v string) *RecognizeKoreanResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeKoreanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeKoreanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeKoreanResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeKoreanResponse) GoString() string {
	return s.String()
}

func (s *RecognizeKoreanResponse) SetHeaders(v map[string]*string) *RecognizeKoreanResponse {
	s.Headers = v
	return s
}

func (s *RecognizeKoreanResponse) SetStatusCode(v int32) *RecognizeKoreanResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeKoreanResponse) SetBody(v *RecognizeKoreanResponseBody) *RecognizeKoreanResponse {
	s.Body = v
	return s
}

type RecognizeLatinRequest struct {
	// example:
	//
	// false
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// example:
	//
	// false
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// example:
	//
	// false
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeLatinRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLatinRequest) GoString() string {
	return s.String()
}

func (s *RecognizeLatinRequest) SetNeedRotate(v bool) *RecognizeLatinRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeLatinRequest) SetOutputCharInfo(v bool) *RecognizeLatinRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeLatinRequest) SetOutputTable(v bool) *RecognizeLatinRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeLatinRequest) SetUrl(v string) *RecognizeLatinRequest {
	s.Url = &v
	return s
}

func (s *RecognizeLatinRequest) SetBody(v io.Reader) *RecognizeLatinRequest {
	s.Body = v
	return s
}

type RecognizeLatinResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"angle":1,"content":"Đậm Phong Cách Khác Biêt  Trên tay chế tác nguyên khối dẫn đầu xu hướng với thiết kế thần máy liền mạch, độ mông ấn tượng 8.5mm cùng   kiểu dáng mặt kinh bóng mượt, sang trọng từ Galaxy M30. Vừa vặn hoền hẩo trong lông bần tay, tho thích thể hiện   phong cách thời thượng với hai phiên bản màu Đen hoặc Xanh cắ tính.  xanh Ngân Hà   Đen Ngả Khói  OC S   ","height":821,"orgHeight":803,"orgWidth":1075,"prism_version":"1.0.9","prism_wnum":9,"prism_wordsInfo":[{"angle":0,"direction":0,"height":37,"pos":[{"x":293,"y":37},{"x":776,"y":29},{"x":777,"y":66},{"x":294,"y":74}],"prob":99,"width":484,"word":"Đậm Phong Cách","x":292,"y":24}],"width":1088}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeLatinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLatinResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeLatinResponseBody) SetCode(v string) *RecognizeLatinResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeLatinResponseBody) SetData(v string) *RecognizeLatinResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeLatinResponseBody) SetMessage(v string) *RecognizeLatinResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeLatinResponseBody) SetRequestId(v string) *RecognizeLatinResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeLatinResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeLatinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeLatinResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLatinResponse) GoString() string {
	return s.String()
}

func (s *RecognizeLatinResponse) SetHeaders(v map[string]*string) *RecognizeLatinResponse {
	s.Headers = v
	return s
}

func (s *RecognizeLatinResponse) SetStatusCode(v int32) *RecognizeLatinResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeLatinResponse) SetBody(v *RecognizeLatinResponseBody) *RecognizeLatinResponse {
	s.Body = v
	return s
}

type RecognizeMedicalDeviceManageLicenseRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1ZrF.MuL2gK0jSZFmXXc7iXXa-1417-995.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeMedicalDeviceManageLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMedicalDeviceManageLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeMedicalDeviceManageLicenseRequest) SetUrl(v string) *RecognizeMedicalDeviceManageLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeMedicalDeviceManageLicenseRequest) SetBody(v io.Reader) *RecognizeMedicalDeviceManageLicenseRequest {
	s.Body = v
	return s
}

type RecognizeMedicalDeviceManageLicenseResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeMedicalDeviceManageLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMedicalDeviceManageLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeMedicalDeviceManageLicenseResponseBody) SetCode(v string) *RecognizeMedicalDeviceManageLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeMedicalDeviceManageLicenseResponseBody) SetData(v string) *RecognizeMedicalDeviceManageLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeMedicalDeviceManageLicenseResponseBody) SetMessage(v string) *RecognizeMedicalDeviceManageLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeMedicalDeviceManageLicenseResponseBody) SetRequestId(v string) *RecognizeMedicalDeviceManageLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeMedicalDeviceManageLicenseResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeMedicalDeviceManageLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeMedicalDeviceManageLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMedicalDeviceManageLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeMedicalDeviceManageLicenseResponse) SetHeaders(v map[string]*string) *RecognizeMedicalDeviceManageLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeMedicalDeviceManageLicenseResponse) SetStatusCode(v int32) *RecognizeMedicalDeviceManageLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeMedicalDeviceManageLicenseResponse) SetBody(v *RecognizeMedicalDeviceManageLicenseResponseBody) *RecognizeMedicalDeviceManageLicenseResponse {
	s.Body = v
	return s
}

type RecognizeMedicalDeviceProduceLicenseRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB13MJ.MuT2gK0jSZFvXXXnFXXa-1417-994.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeMedicalDeviceProduceLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMedicalDeviceProduceLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeMedicalDeviceProduceLicenseRequest) SetUrl(v string) *RecognizeMedicalDeviceProduceLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeMedicalDeviceProduceLicenseRequest) SetBody(v io.Reader) *RecognizeMedicalDeviceProduceLicenseRequest {
	s.Body = v
	return s
}

type RecognizeMedicalDeviceProduceLicenseResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeMedicalDeviceProduceLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMedicalDeviceProduceLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeMedicalDeviceProduceLicenseResponseBody) SetCode(v string) *RecognizeMedicalDeviceProduceLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeMedicalDeviceProduceLicenseResponseBody) SetData(v string) *RecognizeMedicalDeviceProduceLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeMedicalDeviceProduceLicenseResponseBody) SetMessage(v string) *RecognizeMedicalDeviceProduceLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeMedicalDeviceProduceLicenseResponseBody) SetRequestId(v string) *RecognizeMedicalDeviceProduceLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeMedicalDeviceProduceLicenseResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeMedicalDeviceProduceLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeMedicalDeviceProduceLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMedicalDeviceProduceLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeMedicalDeviceProduceLicenseResponse) SetHeaders(v map[string]*string) *RecognizeMedicalDeviceProduceLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeMedicalDeviceProduceLicenseResponse) SetStatusCode(v int32) *RecognizeMedicalDeviceProduceLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeMedicalDeviceProduceLicenseResponse) SetBody(v *RecognizeMedicalDeviceProduceLicenseResponseBody) *RecognizeMedicalDeviceProduceLicenseResponse {
	s.Body = v
	return s
}

type RecognizeMixedInvoicesRequest struct {
	MergePdfPages *bool  `json:"MergePdfPages,omitempty" xml:"MergePdfPages,omitempty"`
	PageNo        *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1.bnGbRWD3KVjSZFsXXcqkpXa-1654-2341.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeMixedInvoicesRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMixedInvoicesRequest) GoString() string {
	return s.String()
}

func (s *RecognizeMixedInvoicesRequest) SetMergePdfPages(v bool) *RecognizeMixedInvoicesRequest {
	s.MergePdfPages = &v
	return s
}

func (s *RecognizeMixedInvoicesRequest) SetPageNo(v int32) *RecognizeMixedInvoicesRequest {
	s.PageNo = &v
	return s
}

func (s *RecognizeMixedInvoicesRequest) SetUrl(v string) *RecognizeMixedInvoicesRequest {
	s.Url = &v
	return s
}

func (s *RecognizeMixedInvoicesRequest) SetBody(v io.Reader) *RecognizeMixedInvoicesRequest {
	s.Body = v
	return s
}

type RecognizeMixedInvoicesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeMixedInvoicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMixedInvoicesResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeMixedInvoicesResponseBody) SetCode(v string) *RecognizeMixedInvoicesResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeMixedInvoicesResponseBody) SetData(v string) *RecognizeMixedInvoicesResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeMixedInvoicesResponseBody) SetMessage(v string) *RecognizeMixedInvoicesResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeMixedInvoicesResponseBody) SetRequestId(v string) *RecognizeMixedInvoicesResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeMixedInvoicesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeMixedInvoicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeMixedInvoicesResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMixedInvoicesResponse) GoString() string {
	return s.String()
}

func (s *RecognizeMixedInvoicesResponse) SetHeaders(v map[string]*string) *RecognizeMixedInvoicesResponse {
	s.Headers = v
	return s
}

func (s *RecognizeMixedInvoicesResponse) SetStatusCode(v int32) *RecognizeMixedInvoicesResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeMixedInvoicesResponse) SetBody(v *RecognizeMixedInvoicesResponseBody) *RecognizeMixedInvoicesResponse {
	s.Body = v
	return s
}

type RecognizeMultiLanguageRequest struct {
	// This parameter is required.
	Languages []*string `json:"Languages,omitempty" xml:"Languages,omitempty" type:"Repeated"`
	// example:
	//
	// false
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// example:
	//
	// false
	NeedSortPage *bool `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
	// example:
	//
	// false
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// example:
	//
	// false
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeMultiLanguageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMultiLanguageRequest) GoString() string {
	return s.String()
}

func (s *RecognizeMultiLanguageRequest) SetLanguages(v []*string) *RecognizeMultiLanguageRequest {
	s.Languages = v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetNeedRotate(v bool) *RecognizeMultiLanguageRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetNeedSortPage(v bool) *RecognizeMultiLanguageRequest {
	s.NeedSortPage = &v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetOutputCharInfo(v bool) *RecognizeMultiLanguageRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetOutputTable(v bool) *RecognizeMultiLanguageRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetUrl(v string) *RecognizeMultiLanguageRequest {
	s.Url = &v
	return s
}

func (s *RecognizeMultiLanguageRequest) SetBody(v io.Reader) *RecognizeMultiLanguageRequest {
	s.Body = v
	return s
}

type RecognizeMultiLanguageShrinkRequest struct {
	// This parameter is required.
	LanguagesShrink *string `json:"Languages,omitempty" xml:"Languages,omitempty"`
	// example:
	//
	// false
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// example:
	//
	// false
	NeedSortPage *bool `json:"NeedSortPage,omitempty" xml:"NeedSortPage,omitempty"`
	// example:
	//
	// false
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// example:
	//
	// false
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeMultiLanguageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMultiLanguageShrinkRequest) GoString() string {
	return s.String()
}

func (s *RecognizeMultiLanguageShrinkRequest) SetLanguagesShrink(v string) *RecognizeMultiLanguageShrinkRequest {
	s.LanguagesShrink = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetNeedRotate(v bool) *RecognizeMultiLanguageShrinkRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetNeedSortPage(v bool) *RecognizeMultiLanguageShrinkRequest {
	s.NeedSortPage = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetOutputCharInfo(v bool) *RecognizeMultiLanguageShrinkRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetOutputTable(v bool) *RecognizeMultiLanguageShrinkRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetUrl(v string) *RecognizeMultiLanguageShrinkRequest {
	s.Url = &v
	return s
}

func (s *RecognizeMultiLanguageShrinkRequest) SetBody(v io.Reader) *RecognizeMultiLanguageShrinkRequest {
	s.Body = v
	return s
}

type RecognizeMultiLanguageResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"content":"PACKING  Shipper/Export   Invoice No & Date  STM TECHNOLOGY INC. 20140730- ST44 ","height":1753,"orgHeight":1753,"orgWidth":1240,"prism_version":"1.0.9","prism_wnum":71,"prism_wordsInfo":[{"angle":0,"direction":0,"height":33,"pos":[{"x":348,"y":137},{"x":531,"y":135},{"x":532,"y":168},{"x":348,"y":170}],"prob":99,"recClassify":1,"width":184,"word":"PACKING","x":348,"y":135}],"width":1240}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeMultiLanguageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMultiLanguageResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeMultiLanguageResponseBody) SetCode(v string) *RecognizeMultiLanguageResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeMultiLanguageResponseBody) SetData(v string) *RecognizeMultiLanguageResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeMultiLanguageResponseBody) SetMessage(v string) *RecognizeMultiLanguageResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeMultiLanguageResponseBody) SetRequestId(v string) *RecognizeMultiLanguageResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeMultiLanguageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeMultiLanguageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeMultiLanguageResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMultiLanguageResponse) GoString() string {
	return s.String()
}

func (s *RecognizeMultiLanguageResponse) SetHeaders(v map[string]*string) *RecognizeMultiLanguageResponse {
	s.Headers = v
	return s
}

func (s *RecognizeMultiLanguageResponse) SetStatusCode(v int32) *RecognizeMultiLanguageResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeMultiLanguageResponse) SetBody(v *RecognizeMultiLanguageResponseBody) *RecognizeMultiLanguageResponse {
	s.Body = v
	return s
}

type RecognizeNonTaxInvoiceRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeNonTaxInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeNonTaxInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeNonTaxInvoiceRequest) SetUrl(v string) *RecognizeNonTaxInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeNonTaxInvoiceRequest) SetBody(v io.Reader) *RecognizeNonTaxInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeNonTaxInvoiceResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeNonTaxInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeNonTaxInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeNonTaxInvoiceResponseBody) SetCode(v string) *RecognizeNonTaxInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeNonTaxInvoiceResponseBody) SetData(v string) *RecognizeNonTaxInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeNonTaxInvoiceResponseBody) SetMessage(v string) *RecognizeNonTaxInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeNonTaxInvoiceResponseBody) SetRequestId(v string) *RecognizeNonTaxInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeNonTaxInvoiceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeNonTaxInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeNonTaxInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeNonTaxInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeNonTaxInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeNonTaxInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeNonTaxInvoiceResponse) SetStatusCode(v int32) *RecognizeNonTaxInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeNonTaxInvoiceResponse) SetBody(v *RecognizeNonTaxInvoiceResponseBody) *RecognizeNonTaxInvoiceResponse {
	s.Body = v
	return s
}

type RecognizePassportRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1uHglUgHqK1RjSZFEXXcGMXXa-800-502.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizePassportRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizePassportRequest) GoString() string {
	return s.String()
}

func (s *RecognizePassportRequest) SetUrl(v string) *RecognizePassportRequest {
	s.Url = &v
	return s
}

func (s *RecognizePassportRequest) SetBody(v io.Reader) *RecognizePassportRequest {
	s.Body = v
	return s
}

type RecognizePassportResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizePassportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizePassportResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizePassportResponseBody) SetCode(v string) *RecognizePassportResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizePassportResponseBody) SetData(v string) *RecognizePassportResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizePassportResponseBody) SetMessage(v string) *RecognizePassportResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizePassportResponseBody) SetRequestId(v string) *RecognizePassportResponseBody {
	s.RequestId = &v
	return s
}

type RecognizePassportResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizePassportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizePassportResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizePassportResponse) GoString() string {
	return s.String()
}

func (s *RecognizePassportResponse) SetHeaders(v map[string]*string) *RecognizePassportResponse {
	s.Headers = v
	return s
}

func (s *RecognizePassportResponse) SetStatusCode(v int32) *RecognizePassportResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizePassportResponse) SetBody(v *RecognizePassportResponseBody) *RecognizePassportResponse {
	s.Body = v
	return s
}

type RecognizePaymentRecordRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizePaymentRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizePaymentRecordRequest) GoString() string {
	return s.String()
}

func (s *RecognizePaymentRecordRequest) SetUrl(v string) *RecognizePaymentRecordRequest {
	s.Url = &v
	return s
}

func (s *RecognizePaymentRecordRequest) SetBody(v io.Reader) *RecognizePaymentRecordRequest {
	s.Body = v
	return s
}

type RecognizePaymentRecordResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizePaymentRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizePaymentRecordResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizePaymentRecordResponseBody) SetCode(v string) *RecognizePaymentRecordResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizePaymentRecordResponseBody) SetData(v string) *RecognizePaymentRecordResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizePaymentRecordResponseBody) SetMessage(v string) *RecognizePaymentRecordResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizePaymentRecordResponseBody) SetRequestId(v string) *RecognizePaymentRecordResponseBody {
	s.RequestId = &v
	return s
}

type RecognizePaymentRecordResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizePaymentRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizePaymentRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizePaymentRecordResponse) GoString() string {
	return s.String()
}

func (s *RecognizePaymentRecordResponse) SetHeaders(v map[string]*string) *RecognizePaymentRecordResponse {
	s.Headers = v
	return s
}

func (s *RecognizePaymentRecordResponse) SetStatusCode(v int32) *RecognizePaymentRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizePaymentRecordResponse) SetBody(v *RecognizePaymentRecordResponseBody) *RecognizePaymentRecordResponse {
	s.Body = v
	return s
}

type RecognizePurchaseRecordRequest struct {
	// example:
	//
	// false
	OutputMultiOrders *bool `json:"OutputMultiOrders,omitempty" xml:"OutputMultiOrders,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizePurchaseRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizePurchaseRecordRequest) GoString() string {
	return s.String()
}

func (s *RecognizePurchaseRecordRequest) SetOutputMultiOrders(v bool) *RecognizePurchaseRecordRequest {
	s.OutputMultiOrders = &v
	return s
}

func (s *RecognizePurchaseRecordRequest) SetUrl(v string) *RecognizePurchaseRecordRequest {
	s.Url = &v
	return s
}

func (s *RecognizePurchaseRecordRequest) SetBody(v io.Reader) *RecognizePurchaseRecordRequest {
	s.Body = v
	return s
}

type RecognizePurchaseRecordResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizePurchaseRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizePurchaseRecordResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizePurchaseRecordResponseBody) SetCode(v string) *RecognizePurchaseRecordResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizePurchaseRecordResponseBody) SetData(v string) *RecognizePurchaseRecordResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizePurchaseRecordResponseBody) SetMessage(v string) *RecognizePurchaseRecordResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizePurchaseRecordResponseBody) SetRequestId(v string) *RecognizePurchaseRecordResponseBody {
	s.RequestId = &v
	return s
}

type RecognizePurchaseRecordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizePurchaseRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizePurchaseRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizePurchaseRecordResponse) GoString() string {
	return s.String()
}

func (s *RecognizePurchaseRecordResponse) SetHeaders(v map[string]*string) *RecognizePurchaseRecordResponse {
	s.Headers = v
	return s
}

func (s *RecognizePurchaseRecordResponse) SetStatusCode(v int32) *RecognizePurchaseRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizePurchaseRecordResponse) SetBody(v *RecognizePurchaseRecordResponseBody) *RecognizePurchaseRecordResponse {
	s.Body = v
	return s
}

type RecognizeQuotaInvoiceRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1SwAeXHr1gK0jSZR0XXbP8XXa-870-604.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeQuotaInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQuotaInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceRequest) SetUrl(v string) *RecognizeQuotaInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeQuotaInvoiceRequest) SetBody(v io.Reader) *RecognizeQuotaInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeQuotaInvoiceResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeQuotaInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponseBody) SetCode(v string) *RecognizeQuotaInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBody) SetData(v string) *RecognizeQuotaInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBody) SetMessage(v string) *RecognizeQuotaInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBody) SetRequestId(v string) *RecognizeQuotaInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeQuotaInvoiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeQuotaInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeQuotaInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeQuotaInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeQuotaInvoiceResponse) SetStatusCode(v int32) *RecognizeQuotaInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponse) SetBody(v *RecognizeQuotaInvoiceResponseBody) *RecognizeQuotaInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeRideHailingItineraryRequest struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i1/O1CN01ePLJiZ1n8CTylKsn3_!!6000000005044-2-tps-194-260.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeRideHailingItineraryRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRideHailingItineraryRequest) GoString() string {
	return s.String()
}

func (s *RecognizeRideHailingItineraryRequest) SetUrl(v string) *RecognizeRideHailingItineraryRequest {
	s.Url = &v
	return s
}

func (s *RecognizeRideHailingItineraryRequest) SetBody(v io.Reader) *RecognizeRideHailingItineraryRequest {
	s.Body = v
	return s
}

type RecognizeRideHailingItineraryResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"data":{"serviceProvider":"滴滴出行","applicationDate":"","startTime":"","endTime":"","phoneNumber":"","totalAmount":"","rideDetails":[{"Number":"","carType":"","pickUpTime":"","city":"","startPlace":"","endPlace":"","mileage":"","amount":"","remarks":""}]},"ftype":0,"height":260,"orgHeight":260,"orgWidth":194,"prism_keyValueInfo":[{"key":"serviceProvider","keyProb":99,"value":"滴滴出行","valuePos":[{"x":120,"y":11},{"x":120,"y":21},{"x":57,"y":20},{"x":57,"y":10}],"valueProb":99},{"key":"applicationDate","keyProb":100,"value":"","valueProb":100},{"key":"startTime","keyProb":91,"value":"","valuePos":[{"x":94,"y":46},{"x":94,"y":50},{"x":75,"y":50},{"x":75,"y":46}],"valueProb":91},{"key":"endTime","keyProb":65,"value":"","valuePos":[{"x":112,"y":46},{"x":112,"y":50},{"x":95,"y":50},{"x":95,"y":46}],"valueProb":65},{"key":"phoneNumber","keyProb":100,"value":"","valueProb":100},{"key":"totalAmount","keyProb":100,"value":"","valueProb":100},{"key":"rideDetails","keyProb":100,"value":"[{\"Number\":\"\",\"carType\":\"\",\"pickUpTime\":\"\",\"city\":\"\",\"startPlace\":\"\",\"endPlace\":\"\",\"mileage\":\"\",\"amount\":\"\",\"remarks\":\"\"}]","valueProb":100}],"sliceRect":{"x0":6,"y0":72,"x1":186,"y1":72,"x2":186,"y2":156,"x3":6,"y3":156},"width":194}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeRideHailingItineraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRideHailingItineraryResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeRideHailingItineraryResponseBody) SetCode(v string) *RecognizeRideHailingItineraryResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeRideHailingItineraryResponseBody) SetData(v string) *RecognizeRideHailingItineraryResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeRideHailingItineraryResponseBody) SetMessage(v string) *RecognizeRideHailingItineraryResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeRideHailingItineraryResponseBody) SetRequestId(v string) *RecognizeRideHailingItineraryResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeRideHailingItineraryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeRideHailingItineraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeRideHailingItineraryResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRideHailingItineraryResponse) GoString() string {
	return s.String()
}

func (s *RecognizeRideHailingItineraryResponse) SetHeaders(v map[string]*string) *RecognizeRideHailingItineraryResponse {
	s.Headers = v
	return s
}

func (s *RecognizeRideHailingItineraryResponse) SetStatusCode(v int32) *RecognizeRideHailingItineraryResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeRideHailingItineraryResponse) SetBody(v *RecognizeRideHailingItineraryResponseBody) *RecognizeRideHailingItineraryResponse {
	s.Body = v
	return s
}

type RecognizeRollTicketRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Y2ryJKT2gK0jSZFvXXXnFXXa-438-934.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeRollTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRollTicketRequest) GoString() string {
	return s.String()
}

func (s *RecognizeRollTicketRequest) SetUrl(v string) *RecognizeRollTicketRequest {
	s.Url = &v
	return s
}

func (s *RecognizeRollTicketRequest) SetBody(v io.Reader) *RecognizeRollTicketRequest {
	s.Body = v
	return s
}

type RecognizeRollTicketResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeRollTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRollTicketResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeRollTicketResponseBody) SetCode(v string) *RecognizeRollTicketResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeRollTicketResponseBody) SetData(v string) *RecognizeRollTicketResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeRollTicketResponseBody) SetMessage(v string) *RecognizeRollTicketResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeRollTicketResponseBody) SetRequestId(v string) *RecognizeRollTicketResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeRollTicketResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeRollTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeRollTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRollTicketResponse) GoString() string {
	return s.String()
}

func (s *RecognizeRollTicketResponse) SetHeaders(v map[string]*string) *RecognizeRollTicketResponse {
	s.Headers = v
	return s
}

func (s *RecognizeRollTicketResponse) SetStatusCode(v int32) *RecognizeRollTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeRollTicketResponse) SetBody(v *RecognizeRollTicketResponseBody) *RecognizeRollTicketResponse {
	s.Body = v
	return s
}

type RecognizeRussianRequest struct {
	// example:
	//
	// false
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// example:
	//
	// false
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// example:
	//
	// false
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeRussianRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussianRequest) GoString() string {
	return s.String()
}

func (s *RecognizeRussianRequest) SetNeedRotate(v bool) *RecognizeRussianRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeRussianRequest) SetOutputCharInfo(v bool) *RecognizeRussianRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeRussianRequest) SetOutputTable(v bool) *RecognizeRussianRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeRussianRequest) SetUrl(v string) *RecognizeRussianRequest {
	s.Url = &v
	return s
}

func (s *RecognizeRussianRequest) SetBody(v io.Reader) *RecognizeRussianRequest {
	s.Body = v
	return s
}

type RecognizeRussianResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"content":"Тэбако (коробочка для косметики) с рисунком в виде колес повозки","height":199,"orgHeight":199,"orgWidth":766,"prism_version":"1.0.9","prism_wnum":6,"prism_wordsInfo":[{"angle":-89,"direction":0,"height":722,"pos":[{"x":6,"y":23},{"x":728,"y":26},{"x":727,"y":43},{"x":5,"y":41}],"prob":99,"width":17,"word":"Тэбако (коробочка для косметики) с рисунком в виде колес повозки， покрытая","x":358,"y":-327}],"width":766}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeRussianResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussianResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeRussianResponseBody) SetCode(v string) *RecognizeRussianResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeRussianResponseBody) SetData(v string) *RecognizeRussianResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeRussianResponseBody) SetMessage(v string) *RecognizeRussianResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeRussianResponseBody) SetRequestId(v string) *RecognizeRussianResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeRussianResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeRussianResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeRussianResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussianResponse) GoString() string {
	return s.String()
}

func (s *RecognizeRussianResponse) SetHeaders(v map[string]*string) *RecognizeRussianResponse {
	s.Headers = v
	return s
}

func (s *RecognizeRussianResponse) SetStatusCode(v int32) *RecognizeRussianResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeRussianResponse) SetBody(v *RecognizeRussianResponseBody) *RecognizeRussianResponse {
	s.Body = v
	return s
}

type RecognizeShoppingReceiptRequest struct {
	// example:
	//
	// http://duguang-database-public.oss-cn-hangzhou.aliyuncs.com/multi_receipt_shopping_receipt/shop_receipt__ticket_2020-05-14-11-59-30.540668_01_List.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeShoppingReceiptRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeShoppingReceiptRequest) GoString() string {
	return s.String()
}

func (s *RecognizeShoppingReceiptRequest) SetUrl(v string) *RecognizeShoppingReceiptRequest {
	s.Url = &v
	return s
}

func (s *RecognizeShoppingReceiptRequest) SetBody(v io.Reader) *RecognizeShoppingReceiptRequest {
	s.Body = v
	return s
}

type RecognizeShoppingReceiptResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"data": {"shopName": "世纪联华椒江市府大道店", "receiptDate": "2020-04-23", "receiptTime": "20:26:00", "contactNumber": "88068111", "shopAddress": "", "totalAmount": "566.67"}, "ftype": 0, "height": 1047, "orgHeight": 1055, "orgWidth": 690, "prism_keyValueInfo": [{"key": "shopName", "keyProb": 98, "value": "世纪联华椒江市府大道店", "valuePos": [{"x": 51, "y": 239}, {"x": 53, "y": 208}, {"x": 438, "y": 231}, {"x": 436, "y": 262}], "valueProb": 98}, {"key": "receiptDate", "keyProb": 100, "value": "2020-04-23", "valuePos": [{"x": 292, "y": 677}, {"x": 293, "y": 649}, {"x": 428, "y": 651}, {"x": 428, "y": 680}], "valueProb": 100}, {"key": "receiptTime", "keyProb": 100, "value": "20:26:00", "valuePos": [{"x": 435, "y": 681}, {"x": 435, "y": 652}, {"x": 548, "y": 656}, {"x": 547, "y": 684}], "valueProb": 100}, {"key": "contactNumber", "keyProb": 100, "value": "88068111", "valuePos": [{"x": 52, "y": 271}, {"x": 52, "y": 242}, {"x": 160, "y": 246}, {"x": 159, "y": 274}], "valueProb": 100}, {"key": "shopAddress", "keyProb": 100, "value": "", "valueProb": 100}, {"key": "totalAmount", "keyProb": 100, "value": "566.67", "valuePos": [{"x": 206, "y": 522}, {"x": 206, "y": 493}, {"x": 313, "y": 495}, {"x": 313, "y": 524}], "valueProb": 100}], "sliceRect": {"x0": 17, "y0": 8, "x1": 690, "y1": 42, "x2": 690, "y2": 1054, "x3": 6, "y3": 1053}, "width": 684}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeShoppingReceiptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeShoppingReceiptResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeShoppingReceiptResponseBody) SetCode(v string) *RecognizeShoppingReceiptResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeShoppingReceiptResponseBody) SetData(v string) *RecognizeShoppingReceiptResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeShoppingReceiptResponseBody) SetMessage(v string) *RecognizeShoppingReceiptResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeShoppingReceiptResponseBody) SetRequestId(v string) *RecognizeShoppingReceiptResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeShoppingReceiptResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeShoppingReceiptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeShoppingReceiptResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeShoppingReceiptResponse) GoString() string {
	return s.String()
}

func (s *RecognizeShoppingReceiptResponse) SetHeaders(v map[string]*string) *RecognizeShoppingReceiptResponse {
	s.Headers = v
	return s
}

func (s *RecognizeShoppingReceiptResponse) SetStatusCode(v int32) *RecognizeShoppingReceiptResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeShoppingReceiptResponse) SetBody(v *RecognizeShoppingReceiptResponseBody) *RecognizeShoppingReceiptResponse {
	s.Body = v
	return s
}

type RecognizeSocialSecurityCardRequest struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i4/O1CN01zpM9bJ1Pa5pCwJat7_!!6000000001856-0-tps-282-179.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeSocialSecurityCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSocialSecurityCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeSocialSecurityCardRequest) SetUrl(v string) *RecognizeSocialSecurityCardRequest {
	s.Url = &v
	return s
}

func (s *RecognizeSocialSecurityCardRequest) SetBody(v io.Reader) *RecognizeSocialSecurityCardRequest {
	s.Body = v
	return s
}

type RecognizeSocialSecurityCardResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"angle":0,"data":{"issueDate":"20168月4日","certificateNumber":"2014100285","taxAuthorityName":"格","formType":"第一联","taxNumbe":"","name":"","totalAmountInWords":"肆佰陆拾陆元叁角玖分","totalAmount":"466.39","drawer":"","remarks":"(20141)鄂国证00285001正常申报一般申报滞纳金自行申报松滋市街河市镇现:主管税务所(科、分局):松滋市国家税务局办税服票价格:4615.38、车辆厂牌:铃木牌/SUZUKIHJ125K-车辆型号:铃木牌/SUZUKIHJ125K-2A、车辆识别代号:LC6PCJ2Y5F1014537","taxClearanceDetails":[{"voucherNumber":"320160804000005082","taxType":"车辆购置税","itemName":"车辆购置税","taxPeriod":"2016-08-04至2016-08-04","date":"2016-08-04461.54","amount":""},{"voucherNumber":"320160804000005082","taxType":"车辆购置税","itemName":"滞纳金","taxPeriod":"2016-08-04至2016-08-04","date":"2016-08-044.85","amount":""}]},"ftype":0,"height":712,"orgHeight":712,"orgWidth":1080,"prism_keyValueInfo":[{"key":"issueDate","keyProb":100,"value":"20168月4日","valuePos":[{"x":458,"y":129},{"x":458,"y":110},{"x":639,"y":113},{"x":638,"y":131}],"valueProb":100},{"key":"certificateNumber","keyProb":99,"value":"2014100285","valuePos":[{"x":810,"y":87},{"x":997,"y":83},{"x":997,"y":103},{"x":810,"y":106}],"valueProb":99},{"key":"taxAuthorityName","keyProb":87,"value":"格","valuePos":[{"x":840,"y":103},{"x":840,"y":128},{"x":825,"y":128},{"x":825,"y":103}],"valueProb":87},{"key":"formType","keyProb":100,"value":"第一联","valuePos":[{"x":1036,"y":247},{"x":1051,"y":247},{"x":1051,"y":289},{"x":1036,"y":289}],"valueProb":100},{"key":"taxNumbe","keyProb":100,"value":"","valueProb":100},{"key":"name","keyProb":100,"value":"","valueProb":100},{"key":"totalAmountInWords","keyProb":100,"value":"肆佰陆拾陆元叁角玖分","valuePos":[{"x":239,"y":498},{"x":395,"y":496},{"x":395,"y":514},{"x":239,"y":515}],"valueProb":100},{"key":"totalAmount","keyProb":100,"value":"466.39","valuePos":[{"x":892,"y":494},{"x":957,"y":493},{"x":957,"y":508},{"x":893,"y":510}],"valueProb":100},{"key":"drawer","keyProb":100,"value":"","valueProb":100},{"key":"remarks","keyProb":100,"value":"(20141)鄂国证00285001正常申报一般申报滞纳金自行申报松滋市街河市镇现:主管税务所(科、分局):松滋市国家税务局办税服票价格:4615.38、车辆厂牌:铃木牌/SUZUKIHJ125K-车辆型号:铃木牌/SUZUKIHJ125K-2A、车辆识别代号:LC6PCJ2Y5F1014537","valuePos":[{"x":966,"y":538},{"x":966,"y":663},{"x":610,"y":663},{"x":610,"y":538}],"valueProb":100},{"key":"taxClearanceDetails","keyProb":100,"value":"[{\"voucherNumber\":\"320160804000005082\",\"taxType\":\"车辆购置税\",\"itemName\":\"车辆购置税\",\"taxPeriod\":\"2016-08-04至2016-08-04\",\"date\":\"2016-08-04461.54\",\"amount\":\"\"},{\"voucherNumber\":\"320160804000005082\",\"taxType\":\"车辆购置税\",\"itemName\":\"滞纳金\",\"taxPeriod\":\"2016-08-04至2016-08-04\",\"date\":\"2016-08-044.85\",\"amount\":\"\"}]","valueProb":100}],"sliceRect":{"x0":0,"y0":0,"x1":1077,"y1":0,"x2":1078,"y2":709,"x3":0,"y3":704},"width":1080}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeSocialSecurityCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSocialSecurityCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeSocialSecurityCardResponseBody) SetCode(v string) *RecognizeSocialSecurityCardResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeSocialSecurityCardResponseBody) SetData(v string) *RecognizeSocialSecurityCardResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeSocialSecurityCardResponseBody) SetMessage(v string) *RecognizeSocialSecurityCardResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeSocialSecurityCardResponseBody) SetRequestId(v string) *RecognizeSocialSecurityCardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeSocialSecurityCardResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeSocialSecurityCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeSocialSecurityCardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSocialSecurityCardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeSocialSecurityCardResponse) SetHeaders(v map[string]*string) *RecognizeSocialSecurityCardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeSocialSecurityCardResponse) SetStatusCode(v int32) *RecognizeSocialSecurityCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeSocialSecurityCardResponse) SetBody(v *RecognizeSocialSecurityCardResponseBody) *RecognizeSocialSecurityCardResponse {
	s.Body = v
	return s
}

type RecognizeSocialSecurityCardVersionIIRequest struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i4/O1CN01zpM9bJ1Pa5pCwJat7_!!6000000001856-0-tps-282-179.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeSocialSecurityCardVersionIIRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSocialSecurityCardVersionIIRequest) GoString() string {
	return s.String()
}

func (s *RecognizeSocialSecurityCardVersionIIRequest) SetUrl(v string) *RecognizeSocialSecurityCardVersionIIRequest {
	s.Url = &v
	return s
}

func (s *RecognizeSocialSecurityCardVersionIIRequest) SetBody(v io.Reader) *RecognizeSocialSecurityCardVersionIIRequest {
	s.Body = v
	return s
}

type RecognizeSocialSecurityCardVersionIIResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"angle":0,"data":{"issueDate":"20168月4日","certificateNumber":"2014100285","taxAuthorityName":"格","formType":"第一联","taxNumbe":"","name":"","totalAmountInWords":"肆佰陆拾陆元叁角玖分","totalAmount":"466.39","drawer":"","remarks":"(20141)鄂国证00285001正常申报一般申报滞纳金自行申报松滋市街河市镇现:主管税务所(科、分局):松滋市国家税务局办税服票价格:4615.38、车辆厂牌:铃木牌/SUZUKIHJ125K-车辆型号:铃木牌/SUZUKIHJ125K-2A、车辆识别代号:LC6PCJ2Y5F1014537","taxClearanceDetails":[{"voucherNumber":"320160804000005082","taxType":"车辆购置税","itemName":"车辆购置税","taxPeriod":"2016-08-04至2016-08-04","date":"2016-08-04461.54","amount":""},{"voucherNumber":"320160804000005082","taxType":"车辆购置税","itemName":"滞纳金","taxPeriod":"2016-08-04至2016-08-04","date":"2016-08-044.85","amount":""}]},"ftype":0,"height":712,"orgHeight":712,"orgWidth":1080,"prism_keyValueInfo":[{"key":"issueDate","keyProb":100,"value":"20168月4日","valuePos":[{"x":458,"y":129},{"x":458,"y":110},{"x":639,"y":113},{"x":638,"y":131}],"valueProb":100},{"key":"certificateNumber","keyProb":99,"value":"2014100285","valuePos":[{"x":810,"y":87},{"x":997,"y":83},{"x":997,"y":103},{"x":810,"y":106}],"valueProb":99},{"key":"taxAuthorityName","keyProb":87,"value":"格","valuePos":[{"x":840,"y":103},{"x":840,"y":128},{"x":825,"y":128},{"x":825,"y":103}],"valueProb":87},{"key":"formType","keyProb":100,"value":"第一联","valuePos":[{"x":1036,"y":247},{"x":1051,"y":247},{"x":1051,"y":289},{"x":1036,"y":289}],"valueProb":100},{"key":"taxNumbe","keyProb":100,"value":"","valueProb":100},{"key":"name","keyProb":100,"value":"","valueProb":100},{"key":"totalAmountInWords","keyProb":100,"value":"肆佰陆拾陆元叁角玖分","valuePos":[{"x":239,"y":498},{"x":395,"y":496},{"x":395,"y":514},{"x":239,"y":515}],"valueProb":100},{"key":"totalAmount","keyProb":100,"value":"466.39","valuePos":[{"x":892,"y":494},{"x":957,"y":493},{"x":957,"y":508},{"x":893,"y":510}],"valueProb":100},{"key":"drawer","keyProb":100,"value":"","valueProb":100},{"key":"remarks","keyProb":100,"value":"(20141)鄂国证00285001正常申报一般申报滞纳金自行申报松滋市街河市镇现:主管税务所(科、分局):松滋市国家税务局办税服票价格:4615.38、车辆厂牌:铃木牌/SUZUKIHJ125K-车辆型号:铃木牌/SUZUKIHJ125K-2A、车辆识别代号:LC6PCJ2Y5F1014537","valuePos":[{"x":966,"y":538},{"x":966,"y":663},{"x":610,"y":663},{"x":610,"y":538}],"valueProb":100},{"key":"taxClearanceDetails","keyProb":100,"value":"[{\"voucherNumber\":\"320160804000005082\",\"taxType\":\"车辆购置税\",\"itemName\":\"车辆购置税\",\"taxPeriod\":\"2016-08-04至2016-08-04\",\"date\":\"2016-08-04461.54\",\"amount\":\"\"},{\"voucherNumber\":\"320160804000005082\",\"taxType\":\"车辆购置税\",\"itemName\":\"滞纳金\",\"taxPeriod\":\"2016-08-04至2016-08-04\",\"date\":\"2016-08-044.85\",\"amount\":\"\"}]","valueProb":100}],"sliceRect":{"x0":0,"y0":0,"x1":1077,"y1":0,"x2":1078,"y2":709,"x3":0,"y3":704},"width":1080}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeSocialSecurityCardVersionIIResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSocialSecurityCardVersionIIResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeSocialSecurityCardVersionIIResponseBody) SetCode(v string) *RecognizeSocialSecurityCardVersionIIResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeSocialSecurityCardVersionIIResponseBody) SetData(v string) *RecognizeSocialSecurityCardVersionIIResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeSocialSecurityCardVersionIIResponseBody) SetMessage(v string) *RecognizeSocialSecurityCardVersionIIResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeSocialSecurityCardVersionIIResponseBody) SetRequestId(v string) *RecognizeSocialSecurityCardVersionIIResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeSocialSecurityCardVersionIIResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeSocialSecurityCardVersionIIResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeSocialSecurityCardVersionIIResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSocialSecurityCardVersionIIResponse) GoString() string {
	return s.String()
}

func (s *RecognizeSocialSecurityCardVersionIIResponse) SetHeaders(v map[string]*string) *RecognizeSocialSecurityCardVersionIIResponse {
	s.Headers = v
	return s
}

func (s *RecognizeSocialSecurityCardVersionIIResponse) SetStatusCode(v int32) *RecognizeSocialSecurityCardVersionIIResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeSocialSecurityCardVersionIIResponse) SetBody(v *RecognizeSocialSecurityCardVersionIIResponseBody) *RecognizeSocialSecurityCardVersionIIResponse {
	s.Body = v
	return s
}

type RecognizeTableOcrRequest struct {
	// example:
	//
	// "false"
	IsHandWriting *string `json:"IsHandWriting,omitempty" xml:"IsHandWriting,omitempty"`
	// example:
	//
	// false
	LineLess *bool `json:"LineLess,omitempty" xml:"LineLess,omitempty"`
	// example:
	//
	// true
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// example:
	//
	// false
	SkipDetection *bool `json:"SkipDetection,omitempty" xml:"SkipDetection,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTableOcrRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableOcrRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTableOcrRequest) SetIsHandWriting(v string) *RecognizeTableOcrRequest {
	s.IsHandWriting = &v
	return s
}

func (s *RecognizeTableOcrRequest) SetLineLess(v bool) *RecognizeTableOcrRequest {
	s.LineLess = &v
	return s
}

func (s *RecognizeTableOcrRequest) SetNeedRotate(v bool) *RecognizeTableOcrRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeTableOcrRequest) SetSkipDetection(v bool) *RecognizeTableOcrRequest {
	s.SkipDetection = &v
	return s
}

func (s *RecognizeTableOcrRequest) SetUrl(v string) *RecognizeTableOcrRequest {
	s.Url = &v
	return s
}

func (s *RecognizeTableOcrRequest) SetBody(v io.Reader) *RecognizeTableOcrRequest {
	s.Body = v
	return s
}

type RecognizeTableOcrResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTableOcrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableOcrResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTableOcrResponseBody) SetCode(v string) *RecognizeTableOcrResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTableOcrResponseBody) SetData(v string) *RecognizeTableOcrResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTableOcrResponseBody) SetMessage(v string) *RecognizeTableOcrResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeTableOcrResponseBody) SetRequestId(v string) *RecognizeTableOcrResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeTableOcrResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeTableOcrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTableOcrResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableOcrResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTableOcrResponse) SetHeaders(v map[string]*string) *RecognizeTableOcrResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTableOcrResponse) SetStatusCode(v int32) *RecognizeTableOcrResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTableOcrResponse) SetBody(v *RecognizeTableOcrResponseBody) *RecognizeTableOcrResponse {
	s.Body = v
	return s
}

type RecognizeTaxClearanceCertificateRequest struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i1/O1CN0131X3Xs1d1CHG8oypS_!!6000000003675-0-tps-1080-712.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTaxClearanceCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxClearanceCertificateRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTaxClearanceCertificateRequest) SetUrl(v string) *RecognizeTaxClearanceCertificateRequest {
	s.Url = &v
	return s
}

func (s *RecognizeTaxClearanceCertificateRequest) SetBody(v io.Reader) *RecognizeTaxClearanceCertificateRequest {
	s.Body = v
	return s
}

type RecognizeTaxClearanceCertificateResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"angle":0,"data":{"issueDate":"20168月4日","certificateNumber":"2014100285","taxAuthorityName":"格","formType":"第一联","taxNumbe":"","name":"","totalAmountInWords":"肆佰陆拾陆元叁角玖分","totalAmount":"466.39","drawer":"","remarks":"(20141)鄂国证00285001正常申报一般申报滞纳金自行申报松滋市街河市镇现:主管税务所(科、分局):松滋市国家税务局办税服票价格:4615.38、车辆厂牌:铃木牌/SUZUKIHJ125K-车辆型号:铃木牌/SUZUKIHJ125K-2A、车辆识别代号:LC6PCJ2Y5F1014537","taxClearanceDetails":[{"voucherNumber":"320160804000005082","taxType":"车辆购置税","itemName":"车辆购置税","taxPeriod":"2016-08-04至2016-08-04","date":"2016-08-04461.54","amount":""},{"voucherNumber":"320160804000005082","taxType":"车辆购置税","itemName":"滞纳金","taxPeriod":"2016-08-04至2016-08-04","date":"2016-08-044.85","amount":""}]},"ftype":0,"height":712,"orgHeight":712,"orgWidth":1080,"prism_keyValueInfo":[{"key":"issueDate","keyProb":100,"value":"20168月4日","valuePos":[{"x":458,"y":129},{"x":458,"y":110},{"x":639,"y":113},{"x":638,"y":131}],"valueProb":100},{"key":"certificateNumber","keyProb":99,"value":"2014100285","valuePos":[{"x":810,"y":87},{"x":997,"y":83},{"x":997,"y":103},{"x":810,"y":106}],"valueProb":99},{"key":"taxAuthorityName","keyProb":87,"value":"格","valuePos":[{"x":840,"y":103},{"x":840,"y":128},{"x":825,"y":128},{"x":825,"y":103}],"valueProb":87},{"key":"formType","keyProb":100,"value":"第一联","valuePos":[{"x":1036,"y":247},{"x":1051,"y":247},{"x":1051,"y":289},{"x":1036,"y":289}],"valueProb":100},{"key":"taxNumbe","keyProb":100,"value":"","valueProb":100},{"key":"name","keyProb":100,"value":"","valueProb":100},{"key":"totalAmountInWords","keyProb":100,"value":"肆佰陆拾陆元叁角玖分","valuePos":[{"x":239,"y":498},{"x":395,"y":496},{"x":395,"y":514},{"x":239,"y":515}],"valueProb":100},{"key":"totalAmount","keyProb":100,"value":"466.39","valuePos":[{"x":892,"y":494},{"x":957,"y":493},{"x":957,"y":508},{"x":893,"y":510}],"valueProb":100},{"key":"drawer","keyProb":100,"value":"","valueProb":100},{"key":"remarks","keyProb":100,"value":"(20141)鄂国证00285001正常申报一般申报滞纳金自行申报松滋市街河市镇现:主管税务所(科、分局):松滋市国家税务局办税服票价格:4615.38、车辆厂牌:铃木牌/SUZUKIHJ125K-车辆型号:铃木牌/SUZUKIHJ125K-2A、车辆识别代号:LC6PCJ2Y5F1014537","valuePos":[{"x":966,"y":538},{"x":966,"y":663},{"x":610,"y":663},{"x":610,"y":538}],"valueProb":100},{"key":"taxClearanceDetails","keyProb":100,"value":"[{\"voucherNumber\":\"320160804000005082\",\"taxType\":\"车辆购置税\",\"itemName\":\"车辆购置税\",\"taxPeriod\":\"2016-08-04至2016-08-04\",\"date\":\"2016-08-04461.54\",\"amount\":\"\"},{\"voucherNumber\":\"320160804000005082\",\"taxType\":\"车辆购置税\",\"itemName\":\"滞纳金\",\"taxPeriod\":\"2016-08-04至2016-08-04\",\"date\":\"2016-08-044.85\",\"amount\":\"\"}]","valueProb":100}],"sliceRect":{"x0":0,"y0":0,"x1":1077,"y1":0,"x2":1078,"y2":709,"x3":0,"y3":704},"width":1080}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTaxClearanceCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxClearanceCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTaxClearanceCertificateResponseBody) SetCode(v string) *RecognizeTaxClearanceCertificateResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTaxClearanceCertificateResponseBody) SetData(v string) *RecognizeTaxClearanceCertificateResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTaxClearanceCertificateResponseBody) SetMessage(v string) *RecognizeTaxClearanceCertificateResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeTaxClearanceCertificateResponseBody) SetRequestId(v string) *RecognizeTaxClearanceCertificateResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeTaxClearanceCertificateResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeTaxClearanceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTaxClearanceCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxClearanceCertificateResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTaxClearanceCertificateResponse) SetHeaders(v map[string]*string) *RecognizeTaxClearanceCertificateResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTaxClearanceCertificateResponse) SetStatusCode(v int32) *RecognizeTaxClearanceCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTaxClearanceCertificateResponse) SetBody(v *RecognizeTaxClearanceCertificateResponseBody) *RecognizeTaxClearanceCertificateResponse {
	s.Body = v
	return s
}

type RecognizeTaxiInvoiceRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1.OicXebviK0jSZFNXXaApXXa-364-982.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTaxiInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceRequest) SetUrl(v string) *RecognizeTaxiInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeTaxiInvoiceRequest) SetBody(v io.Reader) *RecognizeTaxiInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeTaxiInvoiceResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// { 	"data": { 		"licensePlateNumber": "B-30T76", 		"date": "2018-09-28", 		"invoiceCode": "150001583910", 		"invoiceNumber": "22566685", 		"mileage": "22.8", 		"fare": "¥57.00", 		"dropOffTime": "01：40", 		"pickUpTime": "01：19" 	}, 	"ftype": 0, 	"height": 982, 	"orgHeight": 982, 	"orgWidth": 364,  	"width": 364 }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBody) SetCode(v string) *RecognizeTaxiInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBody) SetData(v string) *RecognizeTaxiInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBody) SetMessage(v string) *RecognizeTaxiInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBody) SetRequestId(v string) *RecognizeTaxiInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeTaxiInvoiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeTaxiInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTaxiInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeTaxiInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTaxiInvoiceResponse) SetStatusCode(v int32) *RecognizeTaxiInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponse) SetBody(v *RecognizeTaxiInvoiceResponseBody) *RecognizeTaxiInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeThaiRequest struct {
	// example:
	//
	// false
	NeedRotate *bool `json:"NeedRotate,omitempty" xml:"NeedRotate,omitempty"`
	// example:
	//
	// false
	OutputCharInfo *bool `json:"OutputCharInfo,omitempty" xml:"OutputCharInfo,omitempty"`
	// example:
	//
	// false
	OutputTable *bool `json:"OutputTable,omitempty" xml:"OutputTable,omitempty"`
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeThaiRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeThaiRequest) GoString() string {
	return s.String()
}

func (s *RecognizeThaiRequest) SetNeedRotate(v bool) *RecognizeThaiRequest {
	s.NeedRotate = &v
	return s
}

func (s *RecognizeThaiRequest) SetOutputCharInfo(v bool) *RecognizeThaiRequest {
	s.OutputCharInfo = &v
	return s
}

func (s *RecognizeThaiRequest) SetOutputTable(v bool) *RecognizeThaiRequest {
	s.OutputTable = &v
	return s
}

func (s *RecognizeThaiRequest) SetUrl(v string) *RecognizeThaiRequest {
	s.Url = &v
	return s
}

func (s *RecognizeThaiRequest) SetBody(v io.Reader) *RecognizeThaiRequest {
	s.Body = v
	return s
}

type RecognizeThaiResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"angle":0,"content":"4สัป าR  ๗  เริมเห็นผิวที่เรียบเบียน  วิรีการใช้ LEshop uA","height":887,"orgHeight":887,"orgWidth":790,"prism_version":"1.0.9","prism_wnum":26,"prism_wordsInfo":[{"angle":-89,"direction":0,"height":210,"pos":[{"x":285,"y":14},{"x":495,"y":14},{"x":495,"y":63},{"x":285,"y":63}],"prob":85,"width":48,"word":"4สัป าR ","x":365,"y":-66}],"width":790}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeThaiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeThaiResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeThaiResponseBody) SetCode(v string) *RecognizeThaiResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeThaiResponseBody) SetData(v string) *RecognizeThaiResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeThaiResponseBody) SetMessage(v string) *RecognizeThaiResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeThaiResponseBody) SetRequestId(v string) *RecognizeThaiResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeThaiResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeThaiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeThaiResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeThaiResponse) GoString() string {
	return s.String()
}

func (s *RecognizeThaiResponse) SetHeaders(v map[string]*string) *RecognizeThaiResponse {
	s.Headers = v
	return s
}

func (s *RecognizeThaiResponse) SetStatusCode(v int32) *RecognizeThaiResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeThaiResponse) SetBody(v *RecognizeThaiResponseBody) *RecognizeThaiResponse {
	s.Body = v
	return s
}

type RecognizeTollInvoiceRequest struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i3/O1CN01uUHo411DCwPsBWDMJ_!!6000000000181-0-tps-199-254.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTollInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTollInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTollInvoiceRequest) SetUrl(v string) *RecognizeTollInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeTollInvoiceRequest) SetBody(v io.Reader) *RecognizeTollInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeTollInvoiceResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"angle":0,"data":{"title":"苏宁用打发","formType":"发票联","invoiceCode":"132001681414","invoiceNumber":"53184969","date":"","time":"","vehicleType":"客1","entranceName":"江","exitName":"","totalAmount":"0.00"},"ftype":0,"height":254,"orgHeight":254,"orgWidth":199,"prism_keyValueInfo":[{"key":"title","keyProb":98,"value":"苏宁用打发","valuePos":[{"x":174,"y":20},{"x":174,"y":35},{"x":24,"y":34},{"x":24,"y":19}],"valueProb":98},{"key":"formType","keyProb":89,"value":"发票联","valuePos":[{"x":50,"y":41},{"x":131,"y":37},{"x":131,"y":52},{"x":50,"y":56}],"valueProb":89},{"key":"invoiceCode","keyProb":100,"value":"132001681414","valuePos":[{"x":150,"y":94},{"x":150,"y":105},{"x":63,"y":105},{"x":63,"y":94}],"valueProb":100},{"key":"invoiceNumber","keyProb":100,"value":"53184969","valuePos":[{"x":119,"y":109},{"x":119,"y":120},{"x":63,"y":120},{"x":63,"y":109}],"valueProb":100},{"key":"date","keyProb":100,"value":"","valueProb":100},{"key":"time","keyProb":100,"value":"","valueProb":100},{"key":"vehicleType","keyProb":95,"value":"客1","valuePos":[{"x":40,"y":180},{"x":40,"y":192},{"x":28,"y":192},{"x":28,"y":180}],"valueProb":95},{"key":"entranceName","keyProb":98,"value":"江","valuePos":[{"x":96,"y":128},{"x":96,"y":140},{"x":39,"y":140},{"x":39,"y":128}],"valueProb":98},{"key":"exitName","keyProb":100,"value":"","valueProb":100},{"key":"totalAmount","keyProb":85,"value":"0.00","valuePos":[{"x":70,"y":181},{"x":70,"y":190},{"x":55,"y":190},{"x":55,"y":181}],"valueProb":85}],"sliceRect":{"x0":0,"y0":2,"x1":196,"y1":1,"x2":198,"y2":251,"x3":0,"y3":252},"width":199}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTollInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTollInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTollInvoiceResponseBody) SetCode(v string) *RecognizeTollInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTollInvoiceResponseBody) SetData(v string) *RecognizeTollInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTollInvoiceResponseBody) SetMessage(v string) *RecognizeTollInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeTollInvoiceResponseBody) SetRequestId(v string) *RecognizeTollInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeTollInvoiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeTollInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTollInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTollInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTollInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeTollInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTollInvoiceResponse) SetStatusCode(v int32) *RecognizeTollInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTollInvoiceResponse) SetBody(v *RecognizeTollInvoiceResponseBody) *RecognizeTollInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeTradeMarkCertificationRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1SZiGdfb2gK0jSZK9XXaEgFXa-1654-2340.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTradeMarkCertificationRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTradeMarkCertificationRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTradeMarkCertificationRequest) SetUrl(v string) *RecognizeTradeMarkCertificationRequest {
	s.Url = &v
	return s
}

func (s *RecognizeTradeMarkCertificationRequest) SetBody(v io.Reader) *RecognizeTradeMarkCertificationRequest {
	s.Body = v
	return s
}

type RecognizeTradeMarkCertificationResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTradeMarkCertificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTradeMarkCertificationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTradeMarkCertificationResponseBody) SetCode(v string) *RecognizeTradeMarkCertificationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTradeMarkCertificationResponseBody) SetData(v string) *RecognizeTradeMarkCertificationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTradeMarkCertificationResponseBody) SetMessage(v string) *RecognizeTradeMarkCertificationResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeTradeMarkCertificationResponseBody) SetRequestId(v string) *RecognizeTradeMarkCertificationResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeTradeMarkCertificationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeTradeMarkCertificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTradeMarkCertificationResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTradeMarkCertificationResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTradeMarkCertificationResponse) SetHeaders(v map[string]*string) *RecognizeTradeMarkCertificationResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTradeMarkCertificationResponse) SetStatusCode(v int32) *RecognizeTradeMarkCertificationResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTradeMarkCertificationResponse) SetBody(v *RecognizeTradeMarkCertificationResponseBody) *RecognizeTradeMarkCertificationResponse {
	s.Body = v
	return s
}

type RecognizeTrainInvoiceRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1u1HrUmzqK1RjSZFpXXakSXXa-1200-900.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTrainInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTrainInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTrainInvoiceRequest) SetUrl(v string) *RecognizeTrainInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeTrainInvoiceRequest) SetBody(v io.Reader) *RecognizeTrainInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeTrainInvoiceResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTrainInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTrainInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTrainInvoiceResponseBody) SetCode(v string) *RecognizeTrainInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeTrainInvoiceResponseBody) SetData(v string) *RecognizeTrainInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeTrainInvoiceResponseBody) SetMessage(v string) *RecognizeTrainInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeTrainInvoiceResponseBody) SetRequestId(v string) *RecognizeTrainInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeTrainInvoiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeTrainInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTrainInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTrainInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTrainInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeTrainInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTrainInvoiceResponse) SetStatusCode(v int32) *RecognizeTrainInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTrainInvoiceResponse) SetBody(v *RecognizeTrainInvoiceResponseBody) *RecognizeTrainInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeUsedCarInvoiceRequest struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i4/O1CN01NiY6e220zrtvT6dFJ_!!6000000006921-0-tps-3468-4624.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeUsedCarInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUsedCarInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeUsedCarInvoiceRequest) SetUrl(v string) *RecognizeUsedCarInvoiceRequest {
	s.Url = &v
	return s
}

func (s *RecognizeUsedCarInvoiceRequest) SetBody(v io.Reader) *RecognizeUsedCarInvoiceRequest {
	s.Body = v
	return s
}

type RecognizeUsedCarInvoiceResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"angle":270,"data":{"title":"二手车销售统一发票","formType":"发票联","invoiceDate":"2021-03-19","invoiceCode":"021002000117","invoiceNumber":"00713899","printedInvoiceCode":"021002000117","printedInvoiceNumber":"00713899","taxCode":"03300173880207579449","purchaserName":"李壮","purchaserCode":"210105198712104354","purchaserAddress":"沈阳市皇姑区泰山路69-50号3-1-3","purchaserPhoneNumber":"18947857689","sellerName":"张鹏","sellerCode":"210105197807263716","sellerAddress":"沈阳市皇姑区宁山东路甲2号3-5-2","sellerPhoneNumber":"17641587456","licensePlateNumber":"辽A6L136","certificateNumber":"210008639051","vehicleType":"小型轿车","vinCode":"WAUYGB4H3FN031973","brandMode":"奥迪WAUYGB4H","vehicleAdministrationName":"沈阳市车管所","totalAmountInWords":"壹万圆整","totalAmount":"10000.00","marketName":"沈阳车顶尚二手车交易市场有限公司","marketTaxNumber":"91210106MA0TURHE35","marketAddress":"沈阳市铁西区北二西路29甲4号(9门)","marketBankAccountInfo":"葫芦岛银行股份有限公司沈阳分行20008411159000000025","marketPhoneNumber":"15940287043","remarks":"","drawer":"张丽"},"ftype":0,"height":4624,"orgHeight":4624,"orgWidth":3468,"prism_keyValueInfo":[{"key":"title","keyProb":100,"value":"二手车销售统一发票","valuePos":[{"x":2715,"y":228},{"x":2715,"y":347},{"x":1283,"y":352},{"x":1282,"y":233}],"valueProb":100},{"key":"formType","keyProb":85,"value":"发票联","valuePos":[{"x":2289,"y":401},{"x":2290,"y":510},{"x":1701,"y":512},{"x":1701,"y":403}],"valueProb":85},{"key":"invoiceDate","keyProb":100,"value":"2021-03-19","valuePos":[{"x":728,"y":568},{"x":729,"y":504},{"x":1142,"y":509},{"x":1141,"y":573}],"valueProb":100},{"key":"invoiceCode","keyProb":100,"value":"021002000117","valuePos":[{"x":3090,"y":376},{"x":3676,"y":359},{"x":3678,"y":432},{"x":3093,"y":450}],"valueProb":100},{"key":"invoiceNumber","keyProb":100,"value":"00713899","valuePos":[{"x":3099,"y":457},{"x":3470,"y":449},{"x":3472,"y":523},{"x":3100,"y":530}],"valueProb":100},{"key":"printedInvoiceCode","keyProb":100,"value":"021002000117","valuePos":[{"x":1307,"y":621},{"x":1308,"y":683},{"x":812,"y":688},{"x":812,"y":626}],"valueProb":100},{"key":"printedInvoiceNumber","keyProb":100,"value":"00713899","valuePos":[{"x":811,"y":797},{"x":812,"y":731},{"x":1155,"y":738},{"x":1153,"y":803}],"valueProb":100},{"key":"taxCode","keyProb":100,"value":"03300173880207579449","valuePos":[{"x":3005,"y":755},{"x":3005,"y":818},{"x":2184,"y":825},{"x":2183,"y":761}],"valueProb":100},{"key":"purchaserName","keyProb":100,"value":"李壮","valuePos":[{"x":1139,"y":977},{"x":1260,"y":977},{"x":1260,"y":1044},{"x":1139,"y":1044}],"valueProb":100},{"key":"purchaserCode","keyProb":100,"value":"210105198712104354","valuePos":[{"x":3502,"y":992},{"x":3502,"y":1054},{"x":2802,"y":1054},{"x":2802,"y":992}],"valueProb":100},{"key":"purchaserAddress","keyProb":100,"value":"沈阳市皇姑区泰山路69-50号3-1-3","valuePos":[{"x":1138,"y":1105},{"x":1988,"y":1105},{"x":1988,"y":1176},{"x":1138,"y":1176}],"valueProb":100},{"key":"purchaserPhoneNumber","keyProb":100,"value":"18947857689","valuePos":[{"x":2996,"y":1115},{"x":3466,"y":1115},{"x":3466,"y":1181},{"x":2996,"y":1181}],"valueProb":100},{"key":"sellerName","keyProb":100,"value":"张鹏","valuePos":[{"x":1137,"y":1227},{"x":1259,"y":1227},{"x":1259,"y":1296},{"x":1137,"y":1296}],"valueProb":100},{"key":"sellerCode","keyProb":100,"value":"210105197807263716","valuePos":[{"x":3501,"y":1245},{"x":3501,"y":1305},{"x":2807,"y":1307},{"x":2806,"y":1247}],"valueProb":100},{"key":"sellerAddress","keyProb":100,"value":"沈阳市皇姑区宁山东路甲2号3-5-2","valuePos":[{"x":1991,"y":1353},{"x":1991,"y":1422},{"x":1137,"y":1426},{"x":1136,"y":1356}],"valueProb":100},{"key":"sellerPhoneNumber","keyProb":100,"value":"17641587456","valuePos":[{"x":3460,"y":1372},{"x":3461,"y":1433},{"x":2996,"y":1435},{"x":2996,"y":1373}],"valueProb":100},{"key":"licensePlateNumber","keyProb":100,"value":"辽A6L136","valuePos":[{"x":1470,"y":1471},{"x":1471,"y":1541},{"x":1140,"y":1544},{"x":1139,"y":1474}],"valueProb":100},{"key":"certificateNumber","keyProb":100,"value":"210008639051","valuePos":[{"x":2433,"y":1489},{"x":2433,"y":1549},{"x":1981,"y":1553},{"x":1981,"y":1493}],"valueProb":100},{"key":"vehicleType","keyProb":100,"value":"小型轿车","valuePos":[{"x":2994,"y":1498},{"x":3229,"y":1498},{"x":3229,"y":1562},{"x":2994,"y":1562}],"valueProb":100},{"key":"vinCode","keyProb":100,"value":"WAUYGB4H3FN031973","valuePos":[{"x":1601,"y":1587},{"x":1601,"y":1633},{"x":1138,"y":1638},{"x":1137,"y":1591}],"valueProb":100},{"key":"brandMode","keyProb":100,"value":"奥迪WAUYGB4H","valuePos":[{"x":2330,"y":1616},{"x":2330,"y":1677},{"x":1986,"y":1677},{"x":1986,"y":1616}],"valueProb":100},{"key":"vehicleAdministrationName","keyProb":100,"value":"沈阳市车管所","valuePos":[{"x":3347,"y":1621},{"x":3347,"y":1690},{"x":2989,"y":1693},{"x":2989,"y":1624}],"valueProb":100},{"key":"totalAmountInWords","keyProb":100,"value":"壹万圆整","valuePos":[{"x":1528,"y":1730},{"x":1529,"y":1799},{"x":1292,"y":1801},{"x":1291,"y":1732}],"valueProb":100},{"key":"totalAmount","keyProb":100,"value":"10000.00","valuePos":[{"x":3479,"y":1746},{"x":3479,"y":1816},{"x":3048,"y":1820},{"x":3047,"y":1749}],"valueProb":100},{"key":"marketName","keyProb":100,"value":"沈阳车顶尚二手车交易市场有限公司","valuePos":[{"x":2037,"y":2282},{"x":2037,"y":2354},{"x":1124,"y":2362},{"x":1124,"y":2290}],"valueProb":100},{"key":"marketTaxNumber","keyProb":96,"value":"91210106MA0TURHE35","valuePos":[{"x":3079,"y":2255},{"x":3079,"y":2314},{"x":2397,"y":2321},{"x":2396,"y":2261}],"valueProb":96},{"key":"marketAddress","keyProb":100,"value":"沈阳市铁西区北二西路29甲4号(9门)","valuePos":[{"x":3306,"y":2378},{"x":3307,"y":2445},{"x":2399,"y":2453},{"x":2399,"y":2387}],"valueProb":100},{"key":"marketBankAccountInfo","keyProb":100,"value":"葫芦岛银行股份有限公司沈阳分行20008411159000000025","valuePos":[{"x":2522,"y":2480},{"x":2523,"y":2554},{"x":1109,"y":2567},{"x":1109,"y":2494}],"valueProb":100},{"key":"marketPhoneNumber","keyProb":100,"value":"15940287043","valuePos":[{"x":3172,"y":2579},{"x":3173,"y":2518},{"x":3603,"y":2530},{"x":3601,"y":2590}],"valueProb":100},{"key":"remarks","keyProb":100,"value":"","valueProb":100},{"key":"drawer","keyProb":100,"value":"张丽","valuePos":[{"x":2787,"y":2819},{"x":2789,"y":2756},{"x":2914,"y":2761},{"x":2911,"y":2823}],"valueProb":100}],"sliceRect":{"x0":103,"y0":372,"x1":3174,"y1":428,"x2":3041,"y2":4364,"x3":161,"y3":4360},"width":3468}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeUsedCarInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUsedCarInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeUsedCarInvoiceResponseBody) SetCode(v string) *RecognizeUsedCarInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeUsedCarInvoiceResponseBody) SetData(v string) *RecognizeUsedCarInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeUsedCarInvoiceResponseBody) SetMessage(v string) *RecognizeUsedCarInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeUsedCarInvoiceResponseBody) SetRequestId(v string) *RecognizeUsedCarInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeUsedCarInvoiceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeUsedCarInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeUsedCarInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUsedCarInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeUsedCarInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeUsedCarInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeUsedCarInvoiceResponse) SetStatusCode(v int32) *RecognizeUsedCarInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeUsedCarInvoiceResponse) SetBody(v *RecognizeUsedCarInvoiceResponseBody) *RecognizeUsedCarInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeVehicleCertificationRequest struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i1/O1CN0196uE7i1FXD9TpYqLy_!!6000000000496-0-tps-3024-4032.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeVehicleCertificationRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleCertificationRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleCertificationRequest) SetUrl(v string) *RecognizeVehicleCertificationRequest {
	s.Url = &v
	return s
}

func (s *RecognizeVehicleCertificationRequest) SetBody(v io.Reader) *RecognizeVehicleCertificationRequest {
	s.Body = v
	return s
}

type RecognizeVehicleCertificationResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"data":{"certificateNumber":"YG170ZLM1234567","issueDate":"2021年01月01日","manufactureName":"中国重汽集团济南卡车股份有限公司","vehicleBrand":"豪沃牌","vehicleName":"自卸汽车","vehicleModel":"ZZ3257N414GE1","vinCode":"LZZ1ELSEXLW644557","vehicleColor":"水晶红","chassisModel":"ZZ3257N384GE1","chassisId":"2578516","chassisCertificateNumber":"","engineModel":"WP10H400E50","engineNumber":"7520K064819","fuelType":"柴油","displacement":"9500","power":"294","emissionStandard":"GB17691-2005国V","fuelConsumption":"","overallDimension":"8920×2550×3450","containerDimension":"6000×2350×1500","springNumber":"11/12","tireNumber":"10","tireSize":"12.00R2016PR","frontWheelTrack":"2022","rearWheelTrack":"1850/1850","wheelbase":"4125+1350","axleLoad":"7000/18000(二轴组)","axleNumber":"3","steeringForm":"方向盘","totalWeight":"25000","equipmentWeight":"12500","maximumLadenMass":"12370","massUtilizationCoefficient":"1.00","tractionWeight":"","MaximumLoadMass":"","cabPassengerCapacity":"2","passengerCapacity":"","maxDesignSpeed":"80","manufactureDate":"2020年12月03日","remarks":"备注:货厢自卸方式为后卸。"},"ftype":1,"height":4032,"orgHeight":4032,"orgWidth":3024,"prism_keyValueInfo":[{"key":"certificateNumber","keyProb":100,"value":"YG170ZLM1234567","valuePos":[{"x":554,"y":85},{"x":932,"y":84},{"x":932,"y":133},{"x":554,"y":135}],"valueProb":100},{"key":"issueDate","keyProb":100,"value":"2021年01月01日","valuePos":[{"x":1637,"y":132},{"x":1639,"y":82},{"x":2002,"y":91},{"x":2001,"y":142}],"valueProb":100},{"key":"manufactureName","keyProb":100,"value":"中国重汽集团济南卡车股份有限公司","valuePos":[{"x":552,"y":212},{"x":554,"y":164},{"x":1265,"y":180},{"x":1264,"y":229}],"valueProb":100},{"key":"vehicleBrand","keyProb":100,"value":"豪沃牌","valuePos":[{"x":554,"y":292},{"x":556,"y":240},{"x":693,"y":243},{"x":692,"y":296}],"valueProb":100},{"key":"vehicleName","keyProb":100,"value":"自卸汽车","valuePos":[{"x":1338,"y":257},{"x":1338,"y":307},{"x":1161,"y":307},{"x":1161,"y":257}],"valueProb":100},{"key":"vehicleModel","keyProb":100,"value":"ZZ3257N414GE1","valuePos":[{"x":550,"y":366},{"x":551,"y":319},{"x":846,"y":325},{"x":845,"y":372}],"valueProb":100},{"key":"vinCode","keyProb":100,"value":"LZZ1ELSEXLW644557","valuePos":[{"x":1636,"y":373},{"x":1638,"y":328},{"x":2016,"y":352},{"x":2013,"y":397}],"valueProb":100},{"key":"vehicleColor","keyProb":100,"value":"水晶红","valuePos":[{"x":554,"y":447},{"x":554,"y":395},{"x":690,"y":398},{"x":689,"y":449}],"valueProb":100},{"key":"chassisModel","keyProb":100,"value":"ZZ3257N384GE1","valuePos":[{"x":550,"y":521},{"x":550,"y":474},{"x":848,"y":480},{"x":847,"y":526}],"valueProb":100},{"key":"chassisId","keyProb":100,"value":"2578516","valuePos":[{"x":1635,"y":529},{"x":1637,"y":485},{"x":1801,"y":489},{"x":1800,"y":534}],"valueProb":100},{"key":"chassisCertificateNumber","keyProb":100,"value":"","valueProb":100},{"key":"engineModel","keyProb":100,"value":"WP10H400E50","valuePos":[{"x":1634,"y":607},{"x":1635,"y":562},{"x":1886,"y":570},{"x":1884,"y":614}],"valueProb":100},{"key":"engineNumber","keyProb":100,"value":"7520K064819","valuePos":[{"x":548,"y":672},{"x":549,"y":631},{"x":804,"y":635},{"x":804,"y":676}],"valueProb":100},{"key":"fuelType","keyProb":100,"value":"柴油","valuePos":[{"x":641,"y":705},{"x":641,"y":755},{"x":550,"y":755},{"x":550,"y":705}],"valueProb":100},{"key":"displacement","keyProb":100,"value":"9500","valuePos":[{"x":1631,"y":760},{"x":1631,"y":719},{"x":1728,"y":722},{"x":1727,"y":762}],"valueProb":100},{"key":"power","keyProb":100,"value":"294","valuePos":[{"x":2002,"y":729},{"x":2002,"y":769},{"x":1930,"y":769},{"x":1930,"y":729}],"valueProb":100},{"key":"emissionStandard","keyProb":100,"value":"GB17691-2005国V","valuePos":[{"x":545,"y":828},{"x":545,"y":782},{"x":904,"y":789},{"x":903,"y":835}],"valueProb":100},{"key":"fuelConsumption","keyProb":100,"value":"","valueProb":100},{"key":"overallDimension","keyProb":100,"value":"8920×2550×3450","valuePos":[{"x":547,"y":979},{"x":548,"y":939},{"x":1042,"y":950},{"x":1041,"y":989}],"valueProb":100},{"key":"containerDimension","keyProb":100,"value":"6000×2350×1500","valuePos":[{"x":1628,"y":992},{"x":1629,"y":949},{"x":2119,"y":962},{"x":2117,"y":1005}],"valueProb":100},{"key":"springNumber","keyProb":100,"value":"11/12","valuePos":[{"x":662,"y":1017},{"x":663,"y":1059},{"x":549,"y":1060},{"x":548,"y":1018}],"valueProb":100},{"key":"tireNumber","keyProb":100,"value":"10","valuePos":[{"x":1676,"y":1032},{"x":1676,"y":1073},{"x":1628,"y":1073},{"x":1628,"y":1032}],"valueProb":100},{"key":"tireSize","keyProb":100,"value":"12.00R2016PR","valuePos":[{"x":545,"y":1133},{"x":546,"y":1094},{"x":839,"y":1099},{"x":839,"y":1139}],"valueProb":100},{"key":"frontWheelTrack","keyProb":100,"value":"2022","valuePos":[{"x":640,"y":1169},{"x":640,"y":1208},{"x":545,"y":1210},{"x":544,"y":1170}],"valueProb":100},{"key":"rearWheelTrack","keyProb":100,"value":"1850/1850","valuePos":[{"x":1148,"y":1223},{"x":1149,"y":1183},{"x":1349,"y":1186},{"x":1349,"y":1227}],"valueProb":100},{"key":"wheelbase","keyProb":100,"value":"4125+1350","valuePos":[{"x":546,"y":1286},{"x":547,"y":1244},{"x":752,"y":1248},{"x":751,"y":1290}],"valueProb":100},{"key":"axleLoad","keyProb":100,"value":"7000/18000(二轴组)","valuePos":[{"x":539,"y":1364},{"x":539,"y":1316},{"x":946,"y":1325},{"x":945,"y":1372}],"valueProb":100},{"key":"axleNumber","keyProb":100,"value":"3","valuePos":[{"x":567,"y":1398},{"x":567,"y":1438},{"x":541,"y":1438},{"x":541,"y":1398}],"valueProb":100},{"key":"steeringForm","keyProb":100,"value":"方向盘","valuePos":[{"x":1757,"y":1412},{"x":1757,"y":1463},{"x":1622,"y":1464},{"x":1622,"y":1413}],"valueProb":100},{"key":"totalWeight","keyProb":100,"value":"25000","valuePos":[{"x":536,"y":1512},{"x":538,"y":1471},{"x":658,"y":1475},{"x":657,"y":1515}],"valueProb":100},{"key":"equipmentWeight","keyProb":100,"value":"12500","valuePos":[{"x":1735,"y":1491},{"x":1736,"y":1532},{"x":1620,"y":1534},{"x":1620,"y":1492}],"valueProb":100},{"key":"maximumLadenMass","keyProb":100,"value":"12370","valuePos":[{"x":539,"y":1590},{"x":539,"y":1547},{"x":656,"y":1549},{"x":656,"y":1592}],"valueProb":100},{"key":"massUtilizationCoefficient","keyProb":100,"value":"1.00","valuePos":[{"x":1712,"y":1568},{"x":1712,"y":1608},{"x":1617,"y":1610},{"x":1616,"y":1569}],"valueProb":100},{"key":"tractionWeight","keyProb":100,"value":"","valueProb":100},{"key":"MaximumLoadMass","keyProb":100,"value":"","valueProb":100},{"key":"cabPassengerCapacity","keyProb":100,"value":"2","valuePos":[{"x":560,"y":1777},{"x":560,"y":1817},{"x":532,"y":1817},{"x":532,"y":1777}],"valueProb":100},{"key":"passengerCapacity","keyProb":100,"value":"","valueProb":100},{"key":"maxDesignSpeed","keyProb":100,"value":"80","valuePos":[{"x":581,"y":1931},{"x":581,"y":1971},{"x":530,"y":1971},{"x":530,"y":1931}],"valueProb":100},{"key":"manufactureDate","keyProb":100,"value":"2020年12月03日","valuePos":[{"x":840,"y":2003},{"x":841,"y":2048},{"x":523,"y":2052},{"x":522,"y":2006}],"valueProb":100},{"key":"remarks","keyProb":100,"value":"备注:货厢自卸方式为后卸。","valuePos":[{"x":620,"y":2080},{"x":620,"y":2130},{"x":54,"y":2134},{"x":53,"y":2083}],"valueProb":100}],"sliceRect":{"x0":330,"y0":466,"x1":2530,"y1":420,"x2":2544,"y2":3811,"x3":229,"y3":3746},"width":3024}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVehicleCertificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleCertificationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleCertificationResponseBody) SetCode(v string) *RecognizeVehicleCertificationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeVehicleCertificationResponseBody) SetData(v string) *RecognizeVehicleCertificationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeVehicleCertificationResponseBody) SetMessage(v string) *RecognizeVehicleCertificationResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeVehicleCertificationResponseBody) SetRequestId(v string) *RecognizeVehicleCertificationResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeVehicleCertificationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeVehicleCertificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeVehicleCertificationResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleCertificationResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleCertificationResponse) SetHeaders(v map[string]*string) *RecognizeVehicleCertificationResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVehicleCertificationResponse) SetStatusCode(v int32) *RecognizeVehicleCertificationResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVehicleCertificationResponse) SetBody(v *RecognizeVehicleCertificationResponseBody) *RecognizeVehicleCertificationResponse {
	s.Body = v
	return s
}

type RecognizeVehicleLicenseRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1Wo7eXAvoK1RjSZFDXXXY3pXa-2512-3509.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeVehicleLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleLicenseRequest) SetUrl(v string) *RecognizeVehicleLicenseRequest {
	s.Url = &v
	return s
}

func (s *RecognizeVehicleLicenseRequest) SetBody(v io.Reader) *RecognizeVehicleLicenseRequest {
	s.Body = v
	return s
}

type RecognizeVehicleLicenseResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVehicleLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleLicenseResponseBody) SetCode(v string) *RecognizeVehicleLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeVehicleLicenseResponseBody) SetData(v string) *RecognizeVehicleLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeVehicleLicenseResponseBody) SetMessage(v string) *RecognizeVehicleLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeVehicleLicenseResponseBody) SetRequestId(v string) *RecognizeVehicleLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeVehicleLicenseResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeVehicleLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeVehicleLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleLicenseResponse) SetHeaders(v map[string]*string) *RecognizeVehicleLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVehicleLicenseResponse) SetStatusCode(v int32) *RecognizeVehicleLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVehicleLicenseResponse) SetBody(v *RecognizeVehicleLicenseResponseBody) *RecognizeVehicleLicenseResponse {
	s.Body = v
	return s
}

type RecognizeVehicleRegistrationRequest struct {
	// example:
	//
	// https://img.alicdn.com/imgextra/i1/O1CN01NA1F7A1cSO8cnFQ7m_!!6000000003599-0-tps-844-1125.jpg
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeVehicleRegistrationRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleRegistrationRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleRegistrationRequest) SetUrl(v string) *RecognizeVehicleRegistrationRequest {
	s.Url = &v
	return s
}

func (s *RecognizeVehicleRegistrationRequest) SetBody(v io.Reader) *RecognizeVehicleRegistrationRequest {
	s.Body = v
	return s
}

type RecognizeVehicleRegistrationResponseBody struct {
	// example:
	//
	// noPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"codes":[{"data":"310007798232","points":[{"x":502,"y":6},{"x":768,"y":6},{"x":768,"y":52},{"x":502,"y":52}],"type":"Codabar"}],"data":{"barCode":"310007798232","vehicleOwnerInfo":"某某某限有限公司/统一社会信用代码/12345678682296194","registrationAuthority":"上海市公安局交通警察总队车辆管理所","registrationDate":"2021-04-28","registrationNumber":"沪AG12345","vehicleType":"小型轿车","vehicleBrand":"大众汽车牌","vehicleModel":"SVW7142BPV","vehicleColor":"","vinCode":"LSVCY6C49MN027789","isDomestic":"国产","engineNumber":"035154","engineType":"DUK","fuelType":"混合动力","displacement":"1395","power":"110","manufactureName":"上汽大众汽车有限公司","steeringForm":"方向盘","frontWheelTrack":"1584","rearWheelTrack":"1570","tireNumber":"4","tireSize":"215/60R1695V","springNumber":"","wheelbase":"2871","axleNumber":"2","overallDimension":"4948×1836×1469","containerDimension":"","totalWeight":"2190","permittedWeight":"","passengerCapacity":"","tractionWeight":"","cabPassengerCapacity":"","useNature":"租赁","acquisitionMethod":"购买","manufactureDate":"2021-03-16","issueAuthority":"上海市公安局交通警察总队","issueDate":"2021-04-28"},"ftype":0,"height":1125,"orgHeight":1125,"orgWidth":844,"prism_keyValueInfo":[{"key":"barCode","keyProb":96,"value":"310007798232","valuePos":[{"x":545,"y":45},{"x":735,"y":42},{"x":736,"y":53},{"x":545,"y":55}],"valueProb":96},{"key":"vehicleOwnerInfo","keyProb":100,"value":"某某某限有限公司/统一社会信用代码/12345678682296194","valuePos":[{"x":317,"y":70},{"x":723,"y":68},{"x":723,"y":84},{"x":318,"y":87}],"valueProb":100},{"key":"registrationAuthority","keyProb":100,"value":"上海市公安局交通警察总队车辆管理所","valuePos":[{"x":165,"y":89},{"x":369,"y":87},{"x":370,"y":112},{"x":166,"y":114}],"valueProb":100},{"key":"registrationDate","keyProb":100,"value":"2021-04-28","valuePos":[{"x":463,"y":93},{"x":538,"y":92},{"x":538,"y":104},{"x":464,"y":106}],"valueProb":100},{"key":"registrationNumber","keyProb":100,"value":"沪AG12345","valuePos":[{"x":733,"y":93},{"x":733,"y":107},{"x":669,"y":107},{"x":669,"y":93}],"valueProb":100},{"key":"vehicleType","keyProb":84,"value":"小型轿车","valuePos":[{"x":262,"y":588},{"x":262,"y":603},{"x":205,"y":603},{"x":205,"y":588}],"valueProb":84},{"key":"vehicleBrand","keyProb":100,"value":"大众汽车牌","valuePos":[{"x":569,"y":606},{"x":570,"y":592},{"x":643,"y":594},{"x":642,"y":608}],"valueProb":100},{"key":"vehicleModel","keyProb":99,"value":"SVW7142BPV","valuePos":[{"x":277,"y":616},{"x":277,"y":630},{"x":206,"y":630},{"x":206,"y":616}],"valueProb":99},{"key":"vehicleColor","keyProb":77,"value":"","valuePos":[{"x":585,"y":620},{"x":585,"y":635},{"x":569,"y":635},{"x":569,"y":620}],"valueProb":77},{"key":"vinCode","keyProb":100,"value":"LSVCY6C49MN027789","valuePos":[{"x":324,"y":645},{"x":324,"y":659},{"x":204,"y":659},{"x":204,"y":645}],"valueProb":100},{"key":"isDomestic","keyProb":96,"value":"国产","valuePos":[{"x":568,"y":662},{"x":569,"y":649},{"x":599,"y":650},{"x":599,"y":664}],"valueProb":96},{"key":"engineNumber","keyProb":100,"value":"035154","valuePos":[{"x":203,"y":686},{"x":204,"y":671},{"x":250,"y":672},{"x":250,"y":688}],"valueProb":100},{"key":"engineType","keyProb":100,"value":"DUK","valuePos":[{"x":594,"y":678},{"x":594,"y":692},{"x":568,"y":692},{"x":568,"y":678}],"valueProb":100},{"key":"fuelType","keyProb":100,"value":"混合动力","valuePos":[{"x":260,"y":702},{"x":260,"y":717},{"x":204,"y":717},{"x":204,"y":702}],"valueProb":100},{"key":"displacement","keyProb":100,"value":"1395","valuePos":[{"x":600,"y":707},{"x":600,"y":722},{"x":569,"y":722},{"x":569,"y":707}],"valueProb":100},{"key":"power","keyProb":100,"value":"110","valuePos":[{"x":687,"y":708},{"x":687,"y":723},{"x":663,"y":723},{"x":663,"y":708}],"valueProb":100},{"key":"manufactureName","keyProb":100,"value":"上汽大众汽车有限公司","valuePos":[{"x":342,"y":731},{"x":342,"y":746},{"x":205,"y":746},{"x":205,"y":731}],"valueProb":100},{"key":"steeringForm","keyProb":100,"value":"方向盘","valueProb":100},{"key":"frontWheelTrack","keyProb":100,"value":"1584","valuePos":[{"x":252,"y":760},{"x":252,"y":774},{"x":222,"y":774},{"x":222,"y":760}],"valueProb":100},{"key":"rearWheelTrack","keyProb":100,"value":"1570","valuePos":[{"x":370,"y":761},{"x":370,"y":775},{"x":340,"y":775},{"x":340,"y":761}],"valueProb":100},{"key":"tireNumber","keyProb":100,"value":"4","valuePos":[{"x":580,"y":766},{"x":580,"y":781},{"x":568,"y":781},{"x":568,"y":766}],"valueProb":100},{"key":"tireSize","keyProb":100,"value":"215/60R1695V","valuePos":[{"x":302,"y":788},{"x":302,"y":803},{"x":203,"y":803},{"x":203,"y":788}],"valueProb":100},{"key":"springNumber","keyProb":100,"value":"","valueProb":100},{"key":"wheelbase","keyProb":100,"value":"2871","valuePos":[{"x":232,"y":817},{"x":232,"y":831},{"x":202,"y":831},{"x":202,"y":817}],"valueProb":100},{"key":"axleNumber","keyProb":92,"value":"2","valuePos":[{"x":578,"y":825},{"x":578,"y":839},{"x":569,"y":839},{"x":569,"y":825}],"valueProb":92},{"key":"overallDimension","keyProb":100,"value":"4948×1836×1469","valuePos":[{"x":221,"y":857},{"x":222,"y":845},{"x":475,"y":850},{"x":474,"y":862}],"valueProb":100},{"key":"containerDimension","keyProb":100,"value":"","valueProb":100},{"key":"totalWeight","keyProb":100,"value":"2190","valuePos":[{"x":232,"y":904},{"x":232,"y":918},{"x":203,"y":918},{"x":203,"y":904}],"valueProb":100},{"key":"permittedWeight","keyProb":100,"value":"","valueProb":100},{"key":"passengerCapacity","keyProb":100,"value":"","valueProb":100},{"key":"tractionWeight","keyProb":100,"value":"","valueProb":100},{"key":"cabPassengerCapacity","keyProb":100,"value":"","valueProb":100},{"key":"useNature","keyProb":97,"value":"租赁","valuePos":[{"x":487,"y":968},{"x":487,"y":984},{"x":457,"y":984},{"x":457,"y":968}],"valueProb":97},{"key":"acquisitionMethod","keyProb":100,"value":"购买","valuePos":[{"x":230,"y":992},{"x":230,"y":1008},{"x":200,"y":1008},{"x":200,"y":992}],"valueProb":100},{"key":"manufactureDate","keyProb":100,"value":"2021-03-16","valuePos":[{"x":455,"y":1012},{"x":456,"y":999},{"x":529,"y":1000},{"x":529,"y":1013}],"valueProb":100},{"key":"issueAuthority","keyProb":100,"value":"上海市公安局交通警察总队","valuePos":[{"x":684,"y":895},{"x":684,"y":980},{"x":599,"y":980},{"x":599,"y":895}],"valueProb":100},{"key":"issueDate","keyProb":100,"value":"2021-04-28","valuePos":[{"x":642,"y":1018},{"x":642,"y":1002},{"x":719,"y":1007},{"x":718,"y":1022}],"valueProb":100}],"sliceRect":{"x0":23,"y0":44,"x1":795,"y1":38,"x2":793,"y2":1124,"x3":12,"y3":1106},"width":844}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVehicleRegistrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleRegistrationResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleRegistrationResponseBody) SetCode(v string) *RecognizeVehicleRegistrationResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeVehicleRegistrationResponseBody) SetData(v string) *RecognizeVehicleRegistrationResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeVehicleRegistrationResponseBody) SetMessage(v string) *RecognizeVehicleRegistrationResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeVehicleRegistrationResponseBody) SetRequestId(v string) *RecognizeVehicleRegistrationResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeVehicleRegistrationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeVehicleRegistrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeVehicleRegistrationResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleRegistrationResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleRegistrationResponse) SetHeaders(v map[string]*string) *RecognizeVehicleRegistrationResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVehicleRegistrationResponse) SetStatusCode(v int32) *RecognizeVehicleRegistrationResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVehicleRegistrationResponse) SetBody(v *RecognizeVehicleRegistrationResponseBody) *RecognizeVehicleRegistrationResponse {
	s.Body = v
	return s
}

type RecognizeWaybillRequest struct {
	// example:
	//
	// https://img.alicdn.com/tfs/TB1lOe6VqL7gK0jSZFBXXXZZpXa-480-640.png
	Url  *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeWaybillRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeWaybillRequest) GoString() string {
	return s.String()
}

func (s *RecognizeWaybillRequest) SetUrl(v string) *RecognizeWaybillRequest {
	s.Url = &v
	return s
}

func (s *RecognizeWaybillRequest) SetBody(v io.Reader) *RecognizeWaybillRequest {
	s.Body = v
	return s
}

type RecognizeWaybillResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeWaybillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeWaybillResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeWaybillResponseBody) SetCode(v string) *RecognizeWaybillResponseBody {
	s.Code = &v
	return s
}

func (s *RecognizeWaybillResponseBody) SetData(v string) *RecognizeWaybillResponseBody {
	s.Data = &v
	return s
}

func (s *RecognizeWaybillResponseBody) SetMessage(v string) *RecognizeWaybillResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeWaybillResponseBody) SetRequestId(v string) *RecognizeWaybillResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeWaybillResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeWaybillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeWaybillResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeWaybillResponse) GoString() string {
	return s.String()
}

func (s *RecognizeWaybillResponse) SetHeaders(v map[string]*string) *RecognizeWaybillResponse {
	s.Headers = v
	return s
}

func (s *RecognizeWaybillResponse) SetStatusCode(v int32) *RecognizeWaybillResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeWaybillResponse) SetBody(v *RecognizeWaybillResponseBody) *RecognizeWaybillResponse {
	s.Body = v
	return s
}

type VerifyBusinessLicenseRequest struct {
	// This parameter is required.
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// This parameter is required.
	CreditCode *string `json:"CreditCode,omitempty" xml:"CreditCode,omitempty"`
	// This parameter is required.
	LegalPerson *string `json:"LegalPerson,omitempty" xml:"LegalPerson,omitempty"`
}

func (s VerifyBusinessLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyBusinessLicenseRequest) GoString() string {
	return s.String()
}

func (s *VerifyBusinessLicenseRequest) SetCompanyName(v string) *VerifyBusinessLicenseRequest {
	s.CompanyName = &v
	return s
}

func (s *VerifyBusinessLicenseRequest) SetCreditCode(v string) *VerifyBusinessLicenseRequest {
	s.CreditCode = &v
	return s
}

func (s *VerifyBusinessLicenseRequest) SetLegalPerson(v string) *VerifyBusinessLicenseRequest {
	s.LegalPerson = &v
	return s
}

type VerifyBusinessLicenseResponseBody struct {
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyBusinessLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyBusinessLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyBusinessLicenseResponseBody) SetData(v string) *VerifyBusinessLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *VerifyBusinessLicenseResponseBody) SetRequestId(v string) *VerifyBusinessLicenseResponseBody {
	s.RequestId = &v
	return s
}

type VerifyBusinessLicenseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyBusinessLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyBusinessLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyBusinessLicenseResponse) GoString() string {
	return s.String()
}

func (s *VerifyBusinessLicenseResponse) SetHeaders(v map[string]*string) *VerifyBusinessLicenseResponse {
	s.Headers = v
	return s
}

func (s *VerifyBusinessLicenseResponse) SetStatusCode(v int32) *VerifyBusinessLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyBusinessLicenseResponse) SetBody(v *VerifyBusinessLicenseResponseBody) *VerifyBusinessLicenseResponse {
	s.Body = v
	return s
}

type VerifyVATInvoiceRequest struct {
	InvoiceCode *string `json:"InvoiceCode,omitempty" xml:"InvoiceCode,omitempty"`
	// This parameter is required.
	InvoiceDate *string `json:"InvoiceDate,omitempty" xml:"InvoiceDate,omitempty"`
	InvoiceKind *int32  `json:"InvoiceKind,omitempty" xml:"InvoiceKind,omitempty"`
	// This parameter is required.
	InvoiceNo  *string `json:"InvoiceNo,omitempty" xml:"InvoiceNo,omitempty"`
	InvoiceSum *string `json:"InvoiceSum,omitempty" xml:"InvoiceSum,omitempty"`
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s VerifyVATInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyVATInvoiceRequest) GoString() string {
	return s.String()
}

func (s *VerifyVATInvoiceRequest) SetInvoiceCode(v string) *VerifyVATInvoiceRequest {
	s.InvoiceCode = &v
	return s
}

func (s *VerifyVATInvoiceRequest) SetInvoiceDate(v string) *VerifyVATInvoiceRequest {
	s.InvoiceDate = &v
	return s
}

func (s *VerifyVATInvoiceRequest) SetInvoiceKind(v int32) *VerifyVATInvoiceRequest {
	s.InvoiceKind = &v
	return s
}

func (s *VerifyVATInvoiceRequest) SetInvoiceNo(v string) *VerifyVATInvoiceRequest {
	s.InvoiceNo = &v
	return s
}

func (s *VerifyVATInvoiceRequest) SetInvoiceSum(v string) *VerifyVATInvoiceRequest {
	s.InvoiceSum = &v
	return s
}

func (s *VerifyVATInvoiceRequest) SetVerifyCode(v string) *VerifyVATInvoiceRequest {
	s.VerifyCode = &v
	return s
}

type VerifyVATInvoiceResponseBody struct {
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyVATInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyVATInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyVATInvoiceResponseBody) SetData(v string) *VerifyVATInvoiceResponseBody {
	s.Data = &v
	return s
}

func (s *VerifyVATInvoiceResponseBody) SetRequestId(v string) *VerifyVATInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type VerifyVATInvoiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyVATInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyVATInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyVATInvoiceResponse) GoString() string {
	return s.String()
}

func (s *VerifyVATInvoiceResponse) SetHeaders(v map[string]*string) *VerifyVATInvoiceResponse {
	s.Headers = v
	return s
}

func (s *VerifyVATInvoiceResponse) SetStatusCode(v int32) *VerifyVATInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyVATInvoiceResponse) SetBody(v *VerifyVATInvoiceResponseBody) *VerifyVATInvoiceResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ocr-api"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 全文识别高精版
//
// @param request - RecognizeAdvancedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeAdvancedResponse
func (client *Client) RecognizeAdvancedWithOptions(request *RecognizeAdvancedRequest, runtime *util.RuntimeOptions) (_result *RecognizeAdvancedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.NeedSortPage)) {
		query["NeedSortPage"] = request.NeedSortPage
	}

	if !tea.BoolValue(util.IsUnset(request.NoStamp)) {
		query["NoStamp"] = request.NoStamp
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputFigure)) {
		query["OutputFigure"] = request.OutputFigure
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Paragraph)) {
		query["Paragraph"] = request.Paragraph
	}

	if !tea.BoolValue(util.IsUnset(request.Row)) {
		query["Row"] = request.Row
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeAdvanced"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeAdvancedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 全文识别高精版
//
// @param request - RecognizeAdvancedRequest
//
// @return RecognizeAdvancedResponse
func (client *Client) RecognizeAdvanced(request *RecognizeAdvancedRequest) (_result *RecognizeAdvancedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeAdvancedResponse{}
	_body, _err := client.RecognizeAdvancedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 航空行程单
//
// @param request - RecognizeAirItineraryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeAirItineraryResponse
func (client *Client) RecognizeAirItineraryWithOptions(request *RecognizeAirItineraryRequest, runtime *util.RuntimeOptions) (_result *RecognizeAirItineraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeAirItinerary"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeAirItineraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航空行程单
//
// @param request - RecognizeAirItineraryRequest
//
// @return RecognizeAirItineraryResponse
func (client *Client) RecognizeAirItinerary(request *RecognizeAirItineraryRequest) (_result *RecognizeAirItineraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeAirItineraryResponse{}
	_body, _err := client.RecognizeAirItineraryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 统一Api
//
// @param tmpReq - RecognizeAllTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeAllTextResponse
func (client *Client) RecognizeAllTextWithOptions(tmpReq *RecognizeAllTextRequest, runtime *util.RuntimeOptions) (_result *RecognizeAllTextResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RecognizeAllTextShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AdvancedConfig)) {
		request.AdvancedConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdvancedConfig, tea.String("AdvancedConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.IdCardConfig)) {
		request.IdCardConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IdCardConfig, tea.String("IdCardConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.InternationalBusinessLicenseConfig)) {
		request.InternationalBusinessLicenseConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InternationalBusinessLicenseConfig, tea.String("InternationalBusinessLicenseConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.InternationalIdCardConfig)) {
		request.InternationalIdCardConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InternationalIdCardConfig, tea.String("InternationalIdCardConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.MultiLanConfig)) {
		request.MultiLanConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MultiLanConfig, tea.String("MultiLanConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TableConfig)) {
		request.TableConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableConfig, tea.String("TableConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdvancedConfigShrink)) {
		query["AdvancedConfig"] = request.AdvancedConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardConfigShrink)) {
		query["IdCardConfig"] = request.IdCardConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InternationalBusinessLicenseConfigShrink)) {
		query["InternationalBusinessLicenseConfig"] = request.InternationalBusinessLicenseConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InternationalIdCardConfigShrink)) {
		query["InternationalIdCardConfig"] = request.InternationalIdCardConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MultiLanConfigShrink)) {
		query["MultiLanConfig"] = request.MultiLanConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OutputBarCode)) {
		query["OutputBarCode"] = request.OutputBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCoordinate)) {
		query["OutputCoordinate"] = request.OutputCoordinate
	}

	if !tea.BoolValue(util.IsUnset(request.OutputFigure)) {
		query["OutputFigure"] = request.OutputFigure
	}

	if !tea.BoolValue(util.IsUnset(request.OutputKVExcel)) {
		query["OutputKVExcel"] = request.OutputKVExcel
	}

	if !tea.BoolValue(util.IsUnset(request.OutputOricoord)) {
		query["OutputOricoord"] = request.OutputOricoord
	}

	if !tea.BoolValue(util.IsUnset(request.OutputQrcode)) {
		query["OutputQrcode"] = request.OutputQrcode
	}

	if !tea.BoolValue(util.IsUnset(request.OutputStamp)) {
		query["OutputStamp"] = request.OutputStamp
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.TableConfigShrink)) {
		query["TableConfig"] = request.TableConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: tmpReq.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeAllText"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeAllTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 统一Api
//
// @param request - RecognizeAllTextRequest
//
// @return RecognizeAllTextResponse
func (client *Client) RecognizeAllText(request *RecognizeAllTextRequest) (_result *RecognizeAllTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeAllTextResponse{}
	_body, _err := client.RecognizeAllTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 银承汇票识别
//
// @param request - RecognizeBankAcceptanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeBankAcceptanceResponse
func (client *Client) RecognizeBankAcceptanceWithOptions(request *RecognizeBankAcceptanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeBankAcceptanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBankAcceptance"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBankAcceptanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 银承汇票识别
//
// @param request - RecognizeBankAcceptanceRequest
//
// @return RecognizeBankAcceptanceResponse
func (client *Client) RecognizeBankAcceptance(request *RecognizeBankAcceptanceRequest) (_result *RecognizeBankAcceptanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBankAcceptanceResponse{}
	_body, _err := client.RecognizeBankAcceptanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 银行开户许可证识别
//
// @param request - RecognizeBankAccountLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeBankAccountLicenseResponse
func (client *Client) RecognizeBankAccountLicenseWithOptions(request *RecognizeBankAccountLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeBankAccountLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBankAccountLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBankAccountLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 银行开户许可证识别
//
// @param request - RecognizeBankAccountLicenseRequest
//
// @return RecognizeBankAccountLicenseResponse
func (client *Client) RecognizeBankAccountLicense(request *RecognizeBankAccountLicenseRequest) (_result *RecognizeBankAccountLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBankAccountLicenseResponse{}
	_body, _err := client.RecognizeBankAccountLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 银行卡识别
//
// @param request - RecognizeBankCardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeBankCardResponse
func (client *Client) RecognizeBankCardWithOptions(request *RecognizeBankCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeBankCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBankCard"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBankCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 银行卡识别
//
// @param request - RecognizeBankCardRequest
//
// @return RecognizeBankCardResponse
func (client *Client) RecognizeBankCard(request *RecognizeBankCardRequest) (_result *RecognizeBankCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBankCardResponse{}
	_body, _err := client.RecognizeBankCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 电商图片文字识别
//
// @param request - RecognizeBasicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeBasicResponse
func (client *Client) RecognizeBasicWithOptions(request *RecognizeBasicRequest, runtime *util.RuntimeOptions) (_result *RecognizeBasicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBasic"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBasicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 电商图片文字识别
//
// @param request - RecognizeBasicRequest
//
// @return RecognizeBasicResponse
func (client *Client) RecognizeBasic(request *RecognizeBasicRequest) (_result *RecognizeBasicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBasicResponse{}
	_body, _err := client.RecognizeBasicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 出生证明
//
// @param request - RecognizeBirthCertificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeBirthCertificationResponse
func (client *Client) RecognizeBirthCertificationWithOptions(request *RecognizeBirthCertificationRequest, runtime *util.RuntimeOptions) (_result *RecognizeBirthCertificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBirthCertification"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBirthCertificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 出生证明
//
// @param request - RecognizeBirthCertificationRequest
//
// @return RecognizeBirthCertificationResponse
func (client *Client) RecognizeBirthCertification(request *RecognizeBirthCertificationRequest) (_result *RecognizeBirthCertificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBirthCertificationResponse{}
	_body, _err := client.RecognizeBirthCertificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 客运车船票识别
//
// @param request - RecognizeBusShipTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeBusShipTicketResponse
func (client *Client) RecognizeBusShipTicketWithOptions(request *RecognizeBusShipTicketRequest, runtime *util.RuntimeOptions) (_result *RecognizeBusShipTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBusShipTicket"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBusShipTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 客运车船票识别
//
// @param request - RecognizeBusShipTicketRequest
//
// @return RecognizeBusShipTicketResponse
func (client *Client) RecognizeBusShipTicket(request *RecognizeBusShipTicketRequest) (_result *RecognizeBusShipTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBusShipTicketResponse{}
	_body, _err := client.RecognizeBusShipTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 营业执照识别
//
// @param request - RecognizeBusinessLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeBusinessLicenseResponse
func (client *Client) RecognizeBusinessLicenseWithOptions(request *RecognizeBusinessLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeBusinessLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBusinessLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBusinessLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 营业执照识别
//
// @param request - RecognizeBusinessLicenseRequest
//
// @return RecognizeBusinessLicenseResponse
func (client *Client) RecognizeBusinessLicense(request *RecognizeBusinessLicenseRequest) (_result *RecognizeBusinessLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBusinessLicenseResponse{}
	_body, _err := client.RecognizeBusinessLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机动车销售发票
//
// @param request - RecognizeCarInvoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeCarInvoiceResponse
func (client *Client) RecognizeCarInvoiceWithOptions(request *RecognizeCarInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeCarInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeCarInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeCarInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机动车销售发票
//
// @param request - RecognizeCarInvoiceRequest
//
// @return RecognizeCarInvoiceResponse
func (client *Client) RecognizeCarInvoice(request *RecognizeCarInvoiceRequest) (_result *RecognizeCarInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeCarInvoiceResponse{}
	_body, _err := client.RecognizeCarInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 车牌识别
//
// @param request - RecognizeCarNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeCarNumberResponse
func (client *Client) RecognizeCarNumberWithOptions(request *RecognizeCarNumberRequest, runtime *util.RuntimeOptions) (_result *RecognizeCarNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeCarNumber"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeCarNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 车牌识别
//
// @param request - RecognizeCarNumberRequest
//
// @return RecognizeCarNumberResponse
func (client *Client) RecognizeCarNumber(request *RecognizeCarNumberRequest) (_result *RecognizeCarNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeCarNumberResponse{}
	_body, _err := client.RecognizeCarNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 车辆vin码识别
//
// @param request - RecognizeCarVinCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeCarVinCodeResponse
func (client *Client) RecognizeCarVinCodeWithOptions(request *RecognizeCarVinCodeRequest, runtime *util.RuntimeOptions) (_result *RecognizeCarVinCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeCarVinCode"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeCarVinCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 车辆vin码识别
//
// @param request - RecognizeCarVinCodeRequest
//
// @return RecognizeCarVinCodeResponse
func (client *Client) RecognizeCarVinCode(request *RecognizeCarVinCodeRequest) (_result *RecognizeCarVinCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeCarVinCodeResponse{}
	_body, _err := client.RecognizeCarVinCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 中国护照识别
//
// @param request - RecognizeChinesePassportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeChinesePassportResponse
func (client *Client) RecognizeChinesePassportWithOptions(request *RecognizeChinesePassportRequest, runtime *util.RuntimeOptions) (_result *RecognizeChinesePassportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutputFigure)) {
		query["OutputFigure"] = request.OutputFigure
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeChinesePassport"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeChinesePassportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 中国护照识别
//
// @param request - RecognizeChinesePassportRequest
//
// @return RecognizeChinesePassportResponse
func (client *Client) RecognizeChinesePassport(request *RecognizeChinesePassportRequest) (_result *RecognizeChinesePassportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeChinesePassportResponse{}
	_body, _err := client.RecognizeChinesePassportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通用机打发票识别
//
// @param request - RecognizeCommonPrintedInvoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeCommonPrintedInvoiceResponse
func (client *Client) RecognizeCommonPrintedInvoiceWithOptions(request *RecognizeCommonPrintedInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeCommonPrintedInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeCommonPrintedInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeCommonPrintedInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用机打发票识别
//
// @param request - RecognizeCommonPrintedInvoiceRequest
//
// @return RecognizeCommonPrintedInvoiceResponse
func (client *Client) RecognizeCommonPrintedInvoice(request *RecognizeCommonPrintedInvoiceRequest) (_result *RecognizeCommonPrintedInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeCommonPrintedInvoiceResponse{}
	_body, _err := client.RecognizeCommonPrintedInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 化妆品生产许可证识别
//
// @param request - RecognizeCosmeticProduceLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeCosmeticProduceLicenseResponse
func (client *Client) RecognizeCosmeticProduceLicenseWithOptions(request *RecognizeCosmeticProduceLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeCosmeticProduceLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeCosmeticProduceLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeCosmeticProduceLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 化妆品生产许可证识别
//
// @param request - RecognizeCosmeticProduceLicenseRequest
//
// @return RecognizeCosmeticProduceLicenseResponse
func (client *Client) RecognizeCosmeticProduceLicense(request *RecognizeCosmeticProduceLicenseRequest) (_result *RecognizeCosmeticProduceLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeCosmeticProduceLicenseResponse{}
	_body, _err := client.RecognizeCosmeticProduceLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 核算检测报告识别
//
// @param request - RecognizeCovidTestReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeCovidTestReportResponse
func (client *Client) RecognizeCovidTestReportWithOptions(request *RecognizeCovidTestReportRequest, runtime *util.RuntimeOptions) (_result *RecognizeCovidTestReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MultipleResult)) {
		query["MultipleResult"] = request.MultipleResult
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeCovidTestReport"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeCovidTestReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 核算检测报告识别
//
// @param request - RecognizeCovidTestReportRequest
//
// @return RecognizeCovidTestReportResponse
func (client *Client) RecognizeCovidTestReport(request *RecognizeCovidTestReportRequest) (_result *RecognizeCovidTestReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeCovidTestReportResponse{}
	_body, _err := client.RecognizeCovidTestReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 第二类医疗器械经营备案凭证
//
// @param request - RecognizeCtwoMedicalDeviceManageLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeCtwoMedicalDeviceManageLicenseResponse
func (client *Client) RecognizeCtwoMedicalDeviceManageLicenseWithOptions(request *RecognizeCtwoMedicalDeviceManageLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeCtwoMedicalDeviceManageLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeCtwoMedicalDeviceManageLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeCtwoMedicalDeviceManageLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 第二类医疗器械经营备案凭证
//
// @param request - RecognizeCtwoMedicalDeviceManageLicenseRequest
//
// @return RecognizeCtwoMedicalDeviceManageLicenseResponse
func (client *Client) RecognizeCtwoMedicalDeviceManageLicense(request *RecognizeCtwoMedicalDeviceManageLicenseRequest) (_result *RecognizeCtwoMedicalDeviceManageLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeCtwoMedicalDeviceManageLicenseResponse{}
	_body, _err := client.RecognizeCtwoMedicalDeviceManageLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档结构化识别
//
// @param request - RecognizeDocumentStructureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeDocumentStructureResponse
func (client *Client) RecognizeDocumentStructureWithOptions(request *RecognizeDocumentStructureRequest, runtime *util.RuntimeOptions) (_result *RecognizeDocumentStructureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.NeedSortPage)) {
		query["NeedSortPage"] = request.NeedSortPage
	}

	if !tea.BoolValue(util.IsUnset(request.NoStamp)) {
		query["NoStamp"] = request.NoStamp
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Paragraph)) {
		query["Paragraph"] = request.Paragraph
	}

	if !tea.BoolValue(util.IsUnset(request.Row)) {
		query["Row"] = request.Row
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.UseNewStyleOutput)) {
		query["UseNewStyleOutput"] = request.UseNewStyleOutput
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeDocumentStructure"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeDocumentStructureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档结构化识别
//
// @param request - RecognizeDocumentStructureRequest
//
// @return RecognizeDocumentStructureResponse
func (client *Client) RecognizeDocumentStructure(request *RecognizeDocumentStructureRequest) (_result *RecognizeDocumentStructureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeDocumentStructureResponse{}
	_body, _err := client.RecognizeDocumentStructureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 驾驶证识别
//
// @param request - RecognizeDrivingLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeDrivingLicenseResponse
func (client *Client) RecognizeDrivingLicenseWithOptions(request *RecognizeDrivingLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeDrivingLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeDrivingLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeDrivingLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 驾驶证识别
//
// @param request - RecognizeDrivingLicenseRequest
//
// @return RecognizeDrivingLicenseResponse
func (client *Client) RecognizeDrivingLicense(request *RecognizeDrivingLicenseRequest) (_result *RecognizeDrivingLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeDrivingLicenseResponse{}
	_body, _err := client.RecognizeDrivingLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 印刷体数学公式识别
//
// @param request - RecognizeEduFormulaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeEduFormulaResponse
func (client *Client) RecognizeEduFormulaWithOptions(request *RecognizeEduFormulaRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduFormulaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeEduFormula"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeEduFormulaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 印刷体数学公式识别
//
// @param request - RecognizeEduFormulaRequest
//
// @return RecognizeEduFormulaResponse
func (client *Client) RecognizeEduFormula(request *RecognizeEduFormulaRequest) (_result *RecognizeEduFormulaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeEduFormulaResponse{}
	_body, _err := client.RecognizeEduFormulaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 口算判题
//
// @param request - RecognizeEduOralCalculationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeEduOralCalculationResponse
func (client *Client) RecognizeEduOralCalculationWithOptions(request *RecognizeEduOralCalculationRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduOralCalculationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeEduOralCalculation"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeEduOralCalculationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 口算判题
//
// @param request - RecognizeEduOralCalculationRequest
//
// @return RecognizeEduOralCalculationResponse
func (client *Client) RecognizeEduOralCalculation(request *RecognizeEduOralCalculationRequest) (_result *RecognizeEduOralCalculationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeEduOralCalculationResponse{}
	_body, _err := client.RecognizeEduOralCalculationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 试卷切题识别
//
// @param request - RecognizeEduPaperCutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeEduPaperCutResponse
func (client *Client) RecognizeEduPaperCutWithOptions(request *RecognizeEduPaperCutRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduPaperCutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CutType)) {
		query["CutType"] = request.CutType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageType)) {
		query["ImageType"] = request.ImageType
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		query["Subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeEduPaperCut"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeEduPaperCutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 试卷切题识别
//
// @param request - RecognizeEduPaperCutRequest
//
// @return RecognizeEduPaperCutResponse
func (client *Client) RecognizeEduPaperCut(request *RecognizeEduPaperCutRequest) (_result *RecognizeEduPaperCutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeEduPaperCutResponse{}
	_body, _err := client.RecognizeEduPaperCutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 整页试卷识别
//
// @param request - RecognizeEduPaperOcrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeEduPaperOcrResponse
func (client *Client) RecognizeEduPaperOcrWithOptions(request *RecognizeEduPaperOcrRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduPaperOcrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageType)) {
		query["ImageType"] = request.ImageType
	}

	if !tea.BoolValue(util.IsUnset(request.OutputOricoord)) {
		query["OutputOricoord"] = request.OutputOricoord
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		query["Subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeEduPaperOcr"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeEduPaperOcrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 整页试卷识别
//
// @param request - RecognizeEduPaperOcrRequest
//
// @return RecognizeEduPaperOcrResponse
func (client *Client) RecognizeEduPaperOcr(request *RecognizeEduPaperOcrRequest) (_result *RecognizeEduPaperOcrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeEduPaperOcrResponse{}
	_body, _err := client.RecognizeEduPaperOcrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 精细版结构化切题
//
// @param request - RecognizeEduPaperStructedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeEduPaperStructedResponse
func (client *Client) RecognizeEduPaperStructedWithOptions(request *RecognizeEduPaperStructedRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduPaperStructedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.Subject)) {
		query["Subject"] = request.Subject
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeEduPaperStructed"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeEduPaperStructedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 精细版结构化切题
//
// @param request - RecognizeEduPaperStructedRequest
//
// @return RecognizeEduPaperStructedResponse
func (client *Client) RecognizeEduPaperStructed(request *RecognizeEduPaperStructedRequest) (_result *RecognizeEduPaperStructedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeEduPaperStructedResponse{}
	_body, _err := client.RecognizeEduPaperStructedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 题目识别
//
// @param request - RecognizeEduQuestionOcrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeEduQuestionOcrResponse
func (client *Client) RecognizeEduQuestionOcrWithOptions(request *RecognizeEduQuestionOcrRequest, runtime *util.RuntimeOptions) (_result *RecognizeEduQuestionOcrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeEduQuestionOcr"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeEduQuestionOcrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 题目识别
//
// @param request - RecognizeEduQuestionOcrRequest
//
// @return RecognizeEduQuestionOcrResponse
func (client *Client) RecognizeEduQuestionOcr(request *RecognizeEduQuestionOcrRequest) (_result *RecognizeEduQuestionOcrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeEduQuestionOcrResponse{}
	_body, _err := client.RecognizeEduQuestionOcrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 英语专项识别
//
// @param request - RecognizeEnglishRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeEnglishResponse
func (client *Client) RecognizeEnglishWithOptions(request *RecognizeEnglishRequest, runtime *util.RuntimeOptions) (_result *RecognizeEnglishResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeEnglish"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeEnglishResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 英语专项识别
//
// @param request - RecognizeEnglishRequest
//
// @return RecognizeEnglishResponse
func (client *Client) RecognizeEnglish(request *RecognizeEnglishRequest) (_result *RecognizeEnglishResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeEnglishResponse{}
	_body, _err := client.RecognizeEnglishWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 不动产权证
//
// @param request - RecognizeEstateCertificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeEstateCertificationResponse
func (client *Client) RecognizeEstateCertificationWithOptions(request *RecognizeEstateCertificationRequest, runtime *util.RuntimeOptions) (_result *RecognizeEstateCertificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeEstateCertification"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeEstateCertificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 不动产权证
//
// @param request - RecognizeEstateCertificationRequest
//
// @return RecognizeEstateCertificationResponse
func (client *Client) RecognizeEstateCertification(request *RecognizeEstateCertificationRequest) (_result *RecognizeEstateCertificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeEstateCertificationResponse{}
	_body, _err := client.RecognizeEstateCertificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 来往港澳台通行证识别
//
// @param request - RecognizeExitEntryPermitToHKRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeExitEntryPermitToHKResponse
func (client *Client) RecognizeExitEntryPermitToHKWithOptions(request *RecognizeExitEntryPermitToHKRequest, runtime *util.RuntimeOptions) (_result *RecognizeExitEntryPermitToHKResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutputFigure)) {
		query["OutputFigure"] = request.OutputFigure
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeExitEntryPermitToHK"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeExitEntryPermitToHKResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 来往港澳台通行证识别
//
// @param request - RecognizeExitEntryPermitToHKRequest
//
// @return RecognizeExitEntryPermitToHKResponse
func (client *Client) RecognizeExitEntryPermitToHK(request *RecognizeExitEntryPermitToHKRequest) (_result *RecognizeExitEntryPermitToHKResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeExitEntryPermitToHKResponse{}
	_body, _err := client.RecognizeExitEntryPermitToHKWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 来往大陆(内地)通行证识别
//
// @param request - RecognizeExitEntryPermitToMainlandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeExitEntryPermitToMainlandResponse
func (client *Client) RecognizeExitEntryPermitToMainlandWithOptions(request *RecognizeExitEntryPermitToMainlandRequest, runtime *util.RuntimeOptions) (_result *RecognizeExitEntryPermitToMainlandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutputFigure)) {
		query["OutputFigure"] = request.OutputFigure
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeExitEntryPermitToMainland"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeExitEntryPermitToMainlandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 来往大陆(内地)通行证识别
//
// @param request - RecognizeExitEntryPermitToMainlandRequest
//
// @return RecognizeExitEntryPermitToMainlandResponse
func (client *Client) RecognizeExitEntryPermitToMainland(request *RecognizeExitEntryPermitToMainlandRequest) (_result *RecognizeExitEntryPermitToMainlandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeExitEntryPermitToMainlandResponse{}
	_body, _err := client.RecognizeExitEntryPermitToMainlandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 食品经营许可证
//
// @param request - RecognizeFoodManageLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeFoodManageLicenseResponse
func (client *Client) RecognizeFoodManageLicenseWithOptions(request *RecognizeFoodManageLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeFoodManageLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeFoodManageLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeFoodManageLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 食品经营许可证
//
// @param request - RecognizeFoodManageLicenseRequest
//
// @return RecognizeFoodManageLicenseResponse
func (client *Client) RecognizeFoodManageLicense(request *RecognizeFoodManageLicenseRequest) (_result *RecognizeFoodManageLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeFoodManageLicenseResponse{}
	_body, _err := client.RecognizeFoodManageLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 食品生产许可证
//
// @param request - RecognizeFoodProduceLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeFoodProduceLicenseResponse
func (client *Client) RecognizeFoodProduceLicenseWithOptions(request *RecognizeFoodProduceLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeFoodProduceLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeFoodProduceLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeFoodProduceLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 食品生产许可证
//
// @param request - RecognizeFoodProduceLicenseRequest
//
// @return RecognizeFoodProduceLicenseResponse
func (client *Client) RecognizeFoodProduceLicense(request *RecognizeFoodProduceLicenseRequest) (_result *RecognizeFoodProduceLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeFoodProduceLicenseResponse{}
	_body, _err := client.RecognizeFoodProduceLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通用文字识别
//
// @param request - RecognizeGeneralRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeGeneralResponse
func (client *Client) RecognizeGeneralWithOptions(request *RecognizeGeneralRequest, runtime *util.RuntimeOptions) (_result *RecognizeGeneralResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeGeneral"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeGeneralResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用文字识别
//
// @param request - RecognizeGeneralRequest
//
// @return RecognizeGeneralResponse
func (client *Client) RecognizeGeneral(request *RecognizeGeneralRequest) (_result *RecognizeGeneralResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeGeneralResponse{}
	_body, _err := client.RecognizeGeneralWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 香港身份证识别
//
// @param request - RecognizeHKIdcardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeHKIdcardResponse
func (client *Client) RecognizeHKIdcardWithOptions(request *RecognizeHKIdcardRequest, runtime *util.RuntimeOptions) (_result *RecognizeHKIdcardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeHKIdcard"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeHKIdcardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 香港身份证识别
//
// @param request - RecognizeHKIdcardRequest
//
// @return RecognizeHKIdcardResponse
func (client *Client) RecognizeHKIdcard(request *RecognizeHKIdcardRequest) (_result *RecognizeHKIdcardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeHKIdcardResponse{}
	_body, _err := client.RecognizeHKIdcardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通用手写体识别
//
// @param request - RecognizeHandwritingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeHandwritingResponse
func (client *Client) RecognizeHandwritingWithOptions(request *RecognizeHandwritingRequest, runtime *util.RuntimeOptions) (_result *RecognizeHandwritingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.NeedSortPage)) {
		query["NeedSortPage"] = request.NeedSortPage
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Paragraph)) {
		query["Paragraph"] = request.Paragraph
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeHandwriting"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeHandwritingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用手写体识别
//
// @param request - RecognizeHandwritingRequest
//
// @return RecognizeHandwritingResponse
func (client *Client) RecognizeHandwriting(request *RecognizeHandwritingRequest) (_result *RecognizeHandwritingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeHandwritingResponse{}
	_body, _err := client.RecognizeHandwritingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 防疫健康码识别
//
// @param request - RecognizeHealthCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeHealthCodeResponse
func (client *Client) RecognizeHealthCodeWithOptions(request *RecognizeHealthCodeRequest, runtime *util.RuntimeOptions) (_result *RecognizeHealthCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeHealthCode"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeHealthCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 防疫健康码识别
//
// @param request - RecognizeHealthCodeRequest
//
// @return RecognizeHealthCodeResponse
func (client *Client) RecognizeHealthCode(request *RecognizeHealthCodeRequest) (_result *RecognizeHealthCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeHealthCodeResponse{}
	_body, _err := client.RecognizeHealthCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 酒店流水识别
//
// @param request - RecognizeHotelConsumeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeHotelConsumeResponse
func (client *Client) RecognizeHotelConsumeWithOptions(request *RecognizeHotelConsumeRequest, runtime *util.RuntimeOptions) (_result *RecognizeHotelConsumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeHotelConsume"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeHotelConsumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 酒店流水识别
//
// @param request - RecognizeHotelConsumeRequest
//
// @return RecognizeHotelConsumeResponse
func (client *Client) RecognizeHotelConsume(request *RecognizeHotelConsumeRequest) (_result *RecognizeHotelConsumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeHotelConsumeResponse{}
	_body, _err := client.RecognizeHotelConsumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 户口本识别
//
// @param request - RecognizeHouseholdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeHouseholdResponse
func (client *Client) RecognizeHouseholdWithOptions(request *RecognizeHouseholdRequest, runtime *util.RuntimeOptions) (_result *RecognizeHouseholdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsResidentPage)) {
		query["IsResidentPage"] = request.IsResidentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeHousehold"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeHouseholdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 户口本识别
//
// @param request - RecognizeHouseholdRequest
//
// @return RecognizeHouseholdResponse
func (client *Client) RecognizeHousehold(request *RecognizeHouseholdRequest) (_result *RecognizeHouseholdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeHouseholdResponse{}
	_body, _err := client.RecognizeHouseholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 身份证识别
//
// @param request - RecognizeIdcardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeIdcardResponse
func (client *Client) RecognizeIdcardWithOptions(request *RecognizeIdcardRequest, runtime *util.RuntimeOptions) (_result *RecognizeIdcardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutputFigure)) {
		query["OutputFigure"] = request.OutputFigure
	}

	if !tea.BoolValue(util.IsUnset(request.OutputQualityInfo)) {
		query["OutputQualityInfo"] = request.OutputQualityInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeIdcard"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeIdcardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 身份证识别
//
// @param request - RecognizeIdcardRequest
//
// @return RecognizeIdcardResponse
func (client *Client) RecognizeIdcard(request *RecognizeIdcardRequest) (_result *RecognizeIdcardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeIdcardResponse{}
	_body, _err := client.RecognizeIdcardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 国际营业执照识别
//
// @param request - RecognizeInternationalBusinessLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeInternationalBusinessLicenseResponse
func (client *Client) RecognizeInternationalBusinessLicenseWithOptions(request *RecognizeInternationalBusinessLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeInternationalBusinessLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeInternationalBusinessLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeInternationalBusinessLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际营业执照识别
//
// @param request - RecognizeInternationalBusinessLicenseRequest
//
// @return RecognizeInternationalBusinessLicenseResponse
func (client *Client) RecognizeInternationalBusinessLicense(request *RecognizeInternationalBusinessLicenseRequest) (_result *RecognizeInternationalBusinessLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeInternationalBusinessLicenseResponse{}
	_body, _err := client.RecognizeInternationalBusinessLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 国际身份证识别
//
// @param request - RecognizeInternationalIdcardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeInternationalIdcardResponse
func (client *Client) RecognizeInternationalIdcardWithOptions(request *RecognizeInternationalIdcardRequest, runtime *util.RuntimeOptions) (_result *RecognizeInternationalIdcardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeInternationalIdcard"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeInternationalIdcardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际身份证识别
//
// @param request - RecognizeInternationalIdcardRequest
//
// @return RecognizeInternationalIdcardResponse
func (client *Client) RecognizeInternationalIdcard(request *RecognizeInternationalIdcardRequest) (_result *RecognizeInternationalIdcardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeInternationalIdcardResponse{}
	_body, _err := client.RecognizeInternationalIdcardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 增值税发票识别
//
// @param request - RecognizeInvoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeInvoiceResponse
func (client *Client) RecognizeInvoiceWithOptions(request *RecognizeInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增值税发票识别
//
// @param request - RecognizeInvoiceRequest
//
// @return RecognizeInvoiceResponse
func (client *Client) RecognizeInvoice(request *RecognizeInvoiceRequest) (_result *RecognizeInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeInvoiceResponse{}
	_body, _err := client.RecognizeInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 日语识别
//
// @param request - RecognizeJanpaneseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeJanpaneseResponse
func (client *Client) RecognizeJanpaneseWithOptions(request *RecognizeJanpaneseRequest, runtime *util.RuntimeOptions) (_result *RecognizeJanpaneseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeJanpanese"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeJanpaneseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 日语识别
//
// @param request - RecognizeJanpaneseRequest
//
// @return RecognizeJanpaneseResponse
func (client *Client) RecognizeJanpanese(request *RecognizeJanpaneseRequest) (_result *RecognizeJanpaneseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeJanpaneseResponse{}
	_body, _err := client.RecognizeJanpaneseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 韩语识别
//
// @param request - RecognizeKoreanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeKoreanResponse
func (client *Client) RecognizeKoreanWithOptions(request *RecognizeKoreanRequest, runtime *util.RuntimeOptions) (_result *RecognizeKoreanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeKorean"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeKoreanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 韩语识别
//
// @param request - RecognizeKoreanRequest
//
// @return RecognizeKoreanResponse
func (client *Client) RecognizeKorean(request *RecognizeKoreanRequest) (_result *RecognizeKoreanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeKoreanResponse{}
	_body, _err := client.RecognizeKoreanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拉丁语识别
//
// @param request - RecognizeLatinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeLatinResponse
func (client *Client) RecognizeLatinWithOptions(request *RecognizeLatinRequest, runtime *util.RuntimeOptions) (_result *RecognizeLatinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeLatin"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeLatinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拉丁语识别
//
// @param request - RecognizeLatinRequest
//
// @return RecognizeLatinResponse
func (client *Client) RecognizeLatin(request *RecognizeLatinRequest) (_result *RecognizeLatinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeLatinResponse{}
	_body, _err := client.RecognizeLatinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 医疗器械经营许可证
//
// @param request - RecognizeMedicalDeviceManageLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeMedicalDeviceManageLicenseResponse
func (client *Client) RecognizeMedicalDeviceManageLicenseWithOptions(request *RecognizeMedicalDeviceManageLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeMedicalDeviceManageLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeMedicalDeviceManageLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeMedicalDeviceManageLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 医疗器械经营许可证
//
// @param request - RecognizeMedicalDeviceManageLicenseRequest
//
// @return RecognizeMedicalDeviceManageLicenseResponse
func (client *Client) RecognizeMedicalDeviceManageLicense(request *RecognizeMedicalDeviceManageLicenseRequest) (_result *RecognizeMedicalDeviceManageLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeMedicalDeviceManageLicenseResponse{}
	_body, _err := client.RecognizeMedicalDeviceManageLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 医疗器械生产许可证
//
// @param request - RecognizeMedicalDeviceProduceLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeMedicalDeviceProduceLicenseResponse
func (client *Client) RecognizeMedicalDeviceProduceLicenseWithOptions(request *RecognizeMedicalDeviceProduceLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeMedicalDeviceProduceLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeMedicalDeviceProduceLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeMedicalDeviceProduceLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 医疗器械生产许可证
//
// @param request - RecognizeMedicalDeviceProduceLicenseRequest
//
// @return RecognizeMedicalDeviceProduceLicenseResponse
func (client *Client) RecognizeMedicalDeviceProduceLicense(request *RecognizeMedicalDeviceProduceLicenseRequest) (_result *RecognizeMedicalDeviceProduceLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeMedicalDeviceProduceLicenseResponse{}
	_body, _err := client.RecognizeMedicalDeviceProduceLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 混贴发票识别
//
// @param request - RecognizeMixedInvoicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeMixedInvoicesResponse
func (client *Client) RecognizeMixedInvoicesWithOptions(request *RecognizeMixedInvoicesRequest, runtime *util.RuntimeOptions) (_result *RecognizeMixedInvoicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MergePdfPages)) {
		query["MergePdfPages"] = request.MergePdfPages
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeMixedInvoices"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeMixedInvoicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 混贴发票识别
//
// @param request - RecognizeMixedInvoicesRequest
//
// @return RecognizeMixedInvoicesResponse
func (client *Client) RecognizeMixedInvoices(request *RecognizeMixedInvoicesRequest) (_result *RecognizeMixedInvoicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeMixedInvoicesResponse{}
	_body, _err := client.RecognizeMixedInvoicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通用多语言识别
//
// @param tmpReq - RecognizeMultiLanguageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeMultiLanguageResponse
func (client *Client) RecognizeMultiLanguageWithOptions(tmpReq *RecognizeMultiLanguageRequest, runtime *util.RuntimeOptions) (_result *RecognizeMultiLanguageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RecognizeMultiLanguageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Languages)) {
		request.LanguagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Languages, tea.String("Languages"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LanguagesShrink)) {
		query["Languages"] = request.LanguagesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.NeedSortPage)) {
		query["NeedSortPage"] = request.NeedSortPage
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: tmpReq.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeMultiLanguage"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeMultiLanguageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用多语言识别
//
// @param request - RecognizeMultiLanguageRequest
//
// @return RecognizeMultiLanguageResponse
func (client *Client) RecognizeMultiLanguage(request *RecognizeMultiLanguageRequest) (_result *RecognizeMultiLanguageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeMultiLanguageResponse{}
	_body, _err := client.RecognizeMultiLanguageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 非税收入票据识别
//
// @param request - RecognizeNonTaxInvoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeNonTaxInvoiceResponse
func (client *Client) RecognizeNonTaxInvoiceWithOptions(request *RecognizeNonTaxInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeNonTaxInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeNonTaxInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeNonTaxInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 非税收入票据识别
//
// @param request - RecognizeNonTaxInvoiceRequest
//
// @return RecognizeNonTaxInvoiceResponse
func (client *Client) RecognizeNonTaxInvoice(request *RecognizeNonTaxInvoiceRequest) (_result *RecognizeNonTaxInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeNonTaxInvoiceResponse{}
	_body, _err := client.RecognizeNonTaxInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 护照识别
//
// @param request - RecognizePassportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizePassportResponse
func (client *Client) RecognizePassportWithOptions(request *RecognizePassportRequest, runtime *util.RuntimeOptions) (_result *RecognizePassportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizePassport"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizePassportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 护照识别
//
// @param request - RecognizePassportRequest
//
// @return RecognizePassportResponse
func (client *Client) RecognizePassport(request *RecognizePassportRequest) (_result *RecognizePassportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizePassportResponse{}
	_body, _err := client.RecognizePassportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 支付详情页识别
//
// @param request - RecognizePaymentRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizePaymentRecordResponse
func (client *Client) RecognizePaymentRecordWithOptions(request *RecognizePaymentRecordRequest, runtime *util.RuntimeOptions) (_result *RecognizePaymentRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizePaymentRecord"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizePaymentRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 支付详情页识别
//
// @param request - RecognizePaymentRecordRequest
//
// @return RecognizePaymentRecordResponse
func (client *Client) RecognizePaymentRecord(request *RecognizePaymentRecordRequest) (_result *RecognizePaymentRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizePaymentRecordResponse{}
	_body, _err := client.RecognizePaymentRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 电商订单页识别
//
// @param request - RecognizePurchaseRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizePurchaseRecordResponse
func (client *Client) RecognizePurchaseRecordWithOptions(request *RecognizePurchaseRecordRequest, runtime *util.RuntimeOptions) (_result *RecognizePurchaseRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutputMultiOrders)) {
		query["OutputMultiOrders"] = request.OutputMultiOrders
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizePurchaseRecord"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizePurchaseRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 电商订单页识别
//
// @param request - RecognizePurchaseRecordRequest
//
// @return RecognizePurchaseRecordResponse
func (client *Client) RecognizePurchaseRecord(request *RecognizePurchaseRecordRequest) (_result *RecognizePurchaseRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizePurchaseRecordResponse{}
	_body, _err := client.RecognizePurchaseRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 定额发票
//
// @param request - RecognizeQuotaInvoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeQuotaInvoiceResponse
func (client *Client) RecognizeQuotaInvoiceWithOptions(request *RecognizeQuotaInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeQuotaInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeQuotaInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeQuotaInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 定额发票
//
// @param request - RecognizeQuotaInvoiceRequest
//
// @return RecognizeQuotaInvoiceResponse
func (client *Client) RecognizeQuotaInvoice(request *RecognizeQuotaInvoiceRequest) (_result *RecognizeQuotaInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeQuotaInvoiceResponse{}
	_body, _err := client.RecognizeQuotaInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 网约车行程单识别
//
// @param request - RecognizeRideHailingItineraryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeRideHailingItineraryResponse
func (client *Client) RecognizeRideHailingItineraryWithOptions(request *RecognizeRideHailingItineraryRequest, runtime *util.RuntimeOptions) (_result *RecognizeRideHailingItineraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeRideHailingItinerary"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeRideHailingItineraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 网约车行程单识别
//
// @param request - RecognizeRideHailingItineraryRequest
//
// @return RecognizeRideHailingItineraryResponse
func (client *Client) RecognizeRideHailingItinerary(request *RecognizeRideHailingItineraryRequest) (_result *RecognizeRideHailingItineraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeRideHailingItineraryResponse{}
	_body, _err := client.RecognizeRideHailingItineraryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 增值税发票卷票
//
// @param request - RecognizeRollTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeRollTicketResponse
func (client *Client) RecognizeRollTicketWithOptions(request *RecognizeRollTicketRequest, runtime *util.RuntimeOptions) (_result *RecognizeRollTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeRollTicket"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeRollTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增值税发票卷票
//
// @param request - RecognizeRollTicketRequest
//
// @return RecognizeRollTicketResponse
func (client *Client) RecognizeRollTicket(request *RecognizeRollTicketRequest) (_result *RecognizeRollTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeRollTicketResponse{}
	_body, _err := client.RecognizeRollTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 俄语识别
//
// @param request - RecognizeRussianRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeRussianResponse
func (client *Client) RecognizeRussianWithOptions(request *RecognizeRussianRequest, runtime *util.RuntimeOptions) (_result *RecognizeRussianResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeRussian"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeRussianResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 俄语识别
//
// @param request - RecognizeRussianRequest
//
// @return RecognizeRussianResponse
func (client *Client) RecognizeRussian(request *RecognizeRussianRequest) (_result *RecognizeRussianResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeRussianResponse{}
	_body, _err := client.RecognizeRussianWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 购物小票识别
//
// @param request - RecognizeShoppingReceiptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeShoppingReceiptResponse
func (client *Client) RecognizeShoppingReceiptWithOptions(request *RecognizeShoppingReceiptRequest, runtime *util.RuntimeOptions) (_result *RecognizeShoppingReceiptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeShoppingReceipt"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeShoppingReceiptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 购物小票识别
//
// @param request - RecognizeShoppingReceiptRequest
//
// @return RecognizeShoppingReceiptResponse
func (client *Client) RecognizeShoppingReceipt(request *RecognizeShoppingReceiptRequest) (_result *RecognizeShoppingReceiptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeShoppingReceiptResponse{}
	_body, _err := client.RecognizeShoppingReceiptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 社会保障卡识别
//
// @param request - RecognizeSocialSecurityCardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeSocialSecurityCardResponse
func (client *Client) RecognizeSocialSecurityCardWithOptions(request *RecognizeSocialSecurityCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeSocialSecurityCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeSocialSecurityCard"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeSocialSecurityCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 社会保障卡识别
//
// @param request - RecognizeSocialSecurityCardRequest
//
// @return RecognizeSocialSecurityCardResponse
func (client *Client) RecognizeSocialSecurityCard(request *RecognizeSocialSecurityCardRequest) (_result *RecognizeSocialSecurityCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeSocialSecurityCardResponse{}
	_body, _err := client.RecognizeSocialSecurityCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 社保卡识别
//
// @param request - RecognizeSocialSecurityCardVersionIIRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeSocialSecurityCardVersionIIResponse
func (client *Client) RecognizeSocialSecurityCardVersionIIWithOptions(request *RecognizeSocialSecurityCardVersionIIRequest, runtime *util.RuntimeOptions) (_result *RecognizeSocialSecurityCardVersionIIResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeSocialSecurityCardVersionII"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeSocialSecurityCardVersionIIResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 社保卡识别
//
// @param request - RecognizeSocialSecurityCardVersionIIRequest
//
// @return RecognizeSocialSecurityCardVersionIIResponse
func (client *Client) RecognizeSocialSecurityCardVersionII(request *RecognizeSocialSecurityCardVersionIIRequest) (_result *RecognizeSocialSecurityCardVersionIIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeSocialSecurityCardVersionIIResponse{}
	_body, _err := client.RecognizeSocialSecurityCardVersionIIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 表格识别
//
// @param request - RecognizeTableOcrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeTableOcrResponse
func (client *Client) RecognizeTableOcrWithOptions(request *RecognizeTableOcrRequest, runtime *util.RuntimeOptions) (_result *RecognizeTableOcrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsHandWriting)) {
		query["IsHandWriting"] = request.IsHandWriting
	}

	if !tea.BoolValue(util.IsUnset(request.LineLess)) {
		query["LineLess"] = request.LineLess
	}

	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.SkipDetection)) {
		query["SkipDetection"] = request.SkipDetection
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTableOcr"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTableOcrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 表格识别
//
// @param request - RecognizeTableOcrRequest
//
// @return RecognizeTableOcrResponse
func (client *Client) RecognizeTableOcr(request *RecognizeTableOcrRequest) (_result *RecognizeTableOcrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTableOcrResponse{}
	_body, _err := client.RecognizeTableOcrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 税收完税证明识别
//
// @param request - RecognizeTaxClearanceCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeTaxClearanceCertificateResponse
func (client *Client) RecognizeTaxClearanceCertificateWithOptions(request *RecognizeTaxClearanceCertificateRequest, runtime *util.RuntimeOptions) (_result *RecognizeTaxClearanceCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTaxClearanceCertificate"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTaxClearanceCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 税收完税证明识别
//
// @param request - RecognizeTaxClearanceCertificateRequest
//
// @return RecognizeTaxClearanceCertificateResponse
func (client *Client) RecognizeTaxClearanceCertificate(request *RecognizeTaxClearanceCertificateRequest) (_result *RecognizeTaxClearanceCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTaxClearanceCertificateResponse{}
	_body, _err := client.RecognizeTaxClearanceCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 出租车发票
//
// @param request - RecognizeTaxiInvoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeTaxiInvoiceResponse
func (client *Client) RecognizeTaxiInvoiceWithOptions(request *RecognizeTaxiInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTaxiInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTaxiInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTaxiInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 出租车发票
//
// @param request - RecognizeTaxiInvoiceRequest
//
// @return RecognizeTaxiInvoiceResponse
func (client *Client) RecognizeTaxiInvoice(request *RecognizeTaxiInvoiceRequest) (_result *RecognizeTaxiInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTaxiInvoiceResponse{}
	_body, _err := client.RecognizeTaxiInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 泰语识别
//
// @param request - RecognizeThaiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeThaiResponse
func (client *Client) RecognizeThaiWithOptions(request *RecognizeThaiRequest, runtime *util.RuntimeOptions) (_result *RecognizeThaiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedRotate)) {
		query["NeedRotate"] = request.NeedRotate
	}

	if !tea.BoolValue(util.IsUnset(request.OutputCharInfo)) {
		query["OutputCharInfo"] = request.OutputCharInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OutputTable)) {
		query["OutputTable"] = request.OutputTable
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeThai"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeThaiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 泰语识别
//
// @param request - RecognizeThaiRequest
//
// @return RecognizeThaiResponse
func (client *Client) RecognizeThai(request *RecognizeThaiRequest) (_result *RecognizeThaiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeThaiResponse{}
	_body, _err := client.RecognizeThaiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 过路过桥费发票识别
//
// @param request - RecognizeTollInvoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeTollInvoiceResponse
func (client *Client) RecognizeTollInvoiceWithOptions(request *RecognizeTollInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTollInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTollInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTollInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 过路过桥费发票识别
//
// @param request - RecognizeTollInvoiceRequest
//
// @return RecognizeTollInvoiceResponse
func (client *Client) RecognizeTollInvoice(request *RecognizeTollInvoiceRequest) (_result *RecognizeTollInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTollInvoiceResponse{}
	_body, _err := client.RecognizeTollInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 商标注册证
//
// @param request - RecognizeTradeMarkCertificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeTradeMarkCertificationResponse
func (client *Client) RecognizeTradeMarkCertificationWithOptions(request *RecognizeTradeMarkCertificationRequest, runtime *util.RuntimeOptions) (_result *RecognizeTradeMarkCertificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTradeMarkCertification"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTradeMarkCertificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 商标注册证
//
// @param request - RecognizeTradeMarkCertificationRequest
//
// @return RecognizeTradeMarkCertificationResponse
func (client *Client) RecognizeTradeMarkCertification(request *RecognizeTradeMarkCertificationRequest) (_result *RecognizeTradeMarkCertificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTradeMarkCertificationResponse{}
	_body, _err := client.RecognizeTradeMarkCertificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 火车票
//
// @param request - RecognizeTrainInvoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeTrainInvoiceResponse
func (client *Client) RecognizeTrainInvoiceWithOptions(request *RecognizeTrainInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTrainInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTrainInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTrainInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 火车票
//
// @param request - RecognizeTrainInvoiceRequest
//
// @return RecognizeTrainInvoiceResponse
func (client *Client) RecognizeTrainInvoice(request *RecognizeTrainInvoiceRequest) (_result *RecognizeTrainInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTrainInvoiceResponse{}
	_body, _err := client.RecognizeTrainInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 二手车统一销售发票识别
//
// @param request - RecognizeUsedCarInvoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeUsedCarInvoiceResponse
func (client *Client) RecognizeUsedCarInvoiceWithOptions(request *RecognizeUsedCarInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeUsedCarInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeUsedCarInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeUsedCarInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 二手车统一销售发票识别
//
// @param request - RecognizeUsedCarInvoiceRequest
//
// @return RecognizeUsedCarInvoiceResponse
func (client *Client) RecognizeUsedCarInvoice(request *RecognizeUsedCarInvoiceRequest) (_result *RecognizeUsedCarInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeUsedCarInvoiceResponse{}
	_body, _err := client.RecognizeUsedCarInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 车辆合格证识别
//
// @param request - RecognizeVehicleCertificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeVehicleCertificationResponse
func (client *Client) RecognizeVehicleCertificationWithOptions(request *RecognizeVehicleCertificationRequest, runtime *util.RuntimeOptions) (_result *RecognizeVehicleCertificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVehicleCertification"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVehicleCertificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 车辆合格证识别
//
// @param request - RecognizeVehicleCertificationRequest
//
// @return RecognizeVehicleCertificationResponse
func (client *Client) RecognizeVehicleCertification(request *RecognizeVehicleCertificationRequest) (_result *RecognizeVehicleCertificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVehicleCertificationResponse{}
	_body, _err := client.RecognizeVehicleCertificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 行驶证识别
//
// @param request - RecognizeVehicleLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeVehicleLicenseResponse
func (client *Client) RecognizeVehicleLicenseWithOptions(request *RecognizeVehicleLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeVehicleLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVehicleLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVehicleLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 行驶证识别
//
// @param request - RecognizeVehicleLicenseRequest
//
// @return RecognizeVehicleLicenseResponse
func (client *Client) RecognizeVehicleLicense(request *RecognizeVehicleLicenseRequest) (_result *RecognizeVehicleLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVehicleLicenseResponse{}
	_body, _err := client.RecognizeVehicleLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机动车注册登记证识别
//
// @param request - RecognizeVehicleRegistrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeVehicleRegistrationResponse
func (client *Client) RecognizeVehicleRegistrationWithOptions(request *RecognizeVehicleRegistrationRequest, runtime *util.RuntimeOptions) (_result *RecognizeVehicleRegistrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVehicleRegistration"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVehicleRegistrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机动车注册登记证识别
//
// @param request - RecognizeVehicleRegistrationRequest
//
// @return RecognizeVehicleRegistrationResponse
func (client *Client) RecognizeVehicleRegistration(request *RecognizeVehicleRegistrationRequest) (_result *RecognizeVehicleRegistrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVehicleRegistrationResponse{}
	_body, _err := client.RecognizeVehicleRegistrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 电子面单识别
//
// @param request - RecognizeWaybillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeWaybillResponse
func (client *Client) RecognizeWaybillWithOptions(request *RecognizeWaybillRequest, runtime *util.RuntimeOptions) (_result *RecognizeWaybillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(query),
		Body:   request.Body,
		Stream: request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeWaybill"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeWaybillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 电子面单识别
//
// @param request - RecognizeWaybillRequest
//
// @return RecognizeWaybillResponse
func (client *Client) RecognizeWaybill(request *RecognizeWaybillRequest) (_result *RecognizeWaybillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeWaybillResponse{}
	_body, _err := client.RecognizeWaybillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 营业执照核验
//
// @param request - VerifyBusinessLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyBusinessLicenseResponse
func (client *Client) VerifyBusinessLicenseWithOptions(request *VerifyBusinessLicenseRequest, runtime *util.RuntimeOptions) (_result *VerifyBusinessLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompanyName)) {
		query["CompanyName"] = request.CompanyName
	}

	if !tea.BoolValue(util.IsUnset(request.CreditCode)) {
		query["CreditCode"] = request.CreditCode
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPerson)) {
		query["LegalPerson"] = request.LegalPerson
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyBusinessLicense"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyBusinessLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 营业执照核验
//
// @param request - VerifyBusinessLicenseRequest
//
// @return VerifyBusinessLicenseResponse
func (client *Client) VerifyBusinessLicense(request *VerifyBusinessLicenseRequest) (_result *VerifyBusinessLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyBusinessLicenseResponse{}
	_body, _err := client.VerifyBusinessLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 增值税发票核验
//
// @param request - VerifyVATInvoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyVATInvoiceResponse
func (client *Client) VerifyVATInvoiceWithOptions(request *VerifyVATInvoiceRequest, runtime *util.RuntimeOptions) (_result *VerifyVATInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InvoiceCode)) {
		query["InvoiceCode"] = request.InvoiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceDate)) {
		query["InvoiceDate"] = request.InvoiceDate
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceKind)) {
		query["InvoiceKind"] = request.InvoiceKind
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceNo)) {
		query["InvoiceNo"] = request.InvoiceNo
	}

	if !tea.BoolValue(util.IsUnset(request.InvoiceSum)) {
		query["InvoiceSum"] = request.InvoiceSum
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyCode)) {
		query["VerifyCode"] = request.VerifyCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyVATInvoice"),
		Version:     tea.String("2021-07-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyVATInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增值税发票核验
//
// @param request - VerifyVATInvoiceRequest
//
// @return VerifyVATInvoiceResponse
func (client *Client) VerifyVATInvoice(request *VerifyVATInvoiceRequest) (_result *VerifyVATInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyVATInvoiceResponse{}
	_body, _err := client.VerifyVATInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
