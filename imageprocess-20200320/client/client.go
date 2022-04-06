// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type AnalyzeChestVesselRequest struct {
	DataFormat     *string                             `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	DataSourceType *string                             `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	OrgId          *string                             `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName        *string                             `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	URLList        []*AnalyzeChestVesselRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s AnalyzeChestVesselRequest) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeChestVesselRequest) GoString() string {
	return s.String()
}

func (s *AnalyzeChestVesselRequest) SetDataFormat(v string) *AnalyzeChestVesselRequest {
	s.DataFormat = &v
	return s
}

func (s *AnalyzeChestVesselRequest) SetDataSourceType(v string) *AnalyzeChestVesselRequest {
	s.DataSourceType = &v
	return s
}

func (s *AnalyzeChestVesselRequest) SetOrgId(v string) *AnalyzeChestVesselRequest {
	s.OrgId = &v
	return s
}

func (s *AnalyzeChestVesselRequest) SetOrgName(v string) *AnalyzeChestVesselRequest {
	s.OrgName = &v
	return s
}

func (s *AnalyzeChestVesselRequest) SetURLList(v []*AnalyzeChestVesselRequestURLList) *AnalyzeChestVesselRequest {
	s.URLList = v
	return s
}

type AnalyzeChestVesselRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s AnalyzeChestVesselRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeChestVesselRequestURLList) GoString() string {
	return s.String()
}

func (s *AnalyzeChestVesselRequestURLList) SetURL(v string) *AnalyzeChestVesselRequestURLList {
	s.URL = &v
	return s
}

type AnalyzeChestVesselResponseBody struct {
	Data *AnalyzeChestVesselResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AnalyzeChestVesselResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeChestVesselResponseBody) GoString() string {
	return s.String()
}

func (s *AnalyzeChestVesselResponseBody) SetData(v *AnalyzeChestVesselResponseBodyData) *AnalyzeChestVesselResponseBody {
	s.Data = v
	return s
}

func (s *AnalyzeChestVesselResponseBody) SetRequestId(v string) *AnalyzeChestVesselResponseBody {
	s.RequestId = &v
	return s
}

type AnalyzeChestVesselResponseBodyData struct {
	AortaInfo     *AnalyzeChestVesselResponseBodyDataAortaInfo     `json:"AortaInfo,omitempty" xml:"AortaInfo,omitempty" type:"Struct"`
	PulmonaryInfo *AnalyzeChestVesselResponseBodyDataPulmonaryInfo `json:"PulmonaryInfo,omitempty" xml:"PulmonaryInfo,omitempty" type:"Struct"`
	ResultURL     *string                                          `json:"ResultURL,omitempty" xml:"ResultURL,omitempty"`
}

func (s AnalyzeChestVesselResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeChestVesselResponseBodyData) GoString() string {
	return s.String()
}

func (s *AnalyzeChestVesselResponseBodyData) SetAortaInfo(v *AnalyzeChestVesselResponseBodyDataAortaInfo) *AnalyzeChestVesselResponseBodyData {
	s.AortaInfo = v
	return s
}

func (s *AnalyzeChestVesselResponseBodyData) SetPulmonaryInfo(v *AnalyzeChestVesselResponseBodyDataPulmonaryInfo) *AnalyzeChestVesselResponseBodyData {
	s.PulmonaryInfo = v
	return s
}

func (s *AnalyzeChestVesselResponseBodyData) SetResultURL(v string) *AnalyzeChestVesselResponseBodyData {
	s.ResultURL = &v
	return s
}

type AnalyzeChestVesselResponseBodyDataAortaInfo struct {
	Area         []*float32   `json:"Area,omitempty" xml:"Area,omitempty" type:"Repeated"`
	Coordinates  [][]*float32 `json:"Coordinates,omitempty" xml:"Coordinates,omitempty" type:"Repeated"`
	LabelValue   *int64       `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
	MaxArea      *float32     `json:"MaxArea,omitempty" xml:"MaxArea,omitempty"`
	MaxAreaIndex *int64       `json:"MaxAreaIndex,omitempty" xml:"MaxAreaIndex,omitempty"`
	MaxDiameter  *float32     `json:"MaxDiameter,omitempty" xml:"MaxDiameter,omitempty"`
}

func (s AnalyzeChestVesselResponseBodyDataAortaInfo) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeChestVesselResponseBodyDataAortaInfo) GoString() string {
	return s.String()
}

func (s *AnalyzeChestVesselResponseBodyDataAortaInfo) SetArea(v []*float32) *AnalyzeChestVesselResponseBodyDataAortaInfo {
	s.Area = v
	return s
}

func (s *AnalyzeChestVesselResponseBodyDataAortaInfo) SetCoordinates(v [][]*float32) *AnalyzeChestVesselResponseBodyDataAortaInfo {
	s.Coordinates = v
	return s
}

func (s *AnalyzeChestVesselResponseBodyDataAortaInfo) SetLabelValue(v int64) *AnalyzeChestVesselResponseBodyDataAortaInfo {
	s.LabelValue = &v
	return s
}

func (s *AnalyzeChestVesselResponseBodyDataAortaInfo) SetMaxArea(v float32) *AnalyzeChestVesselResponseBodyDataAortaInfo {
	s.MaxArea = &v
	return s
}

func (s *AnalyzeChestVesselResponseBodyDataAortaInfo) SetMaxAreaIndex(v int64) *AnalyzeChestVesselResponseBodyDataAortaInfo {
	s.MaxAreaIndex = &v
	return s
}

func (s *AnalyzeChestVesselResponseBodyDataAortaInfo) SetMaxDiameter(v float32) *AnalyzeChestVesselResponseBodyDataAortaInfo {
	s.MaxDiameter = &v
	return s
}

type AnalyzeChestVesselResponseBodyDataPulmonaryInfo struct {
	Area             []*float32   `json:"Area,omitempty" xml:"Area,omitempty" type:"Repeated"`
	Coordinates      [][]*float32 `json:"Coordinates,omitempty" xml:"Coordinates,omitempty" type:"Repeated"`
	LabelValue       *int64       `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
	MaxArea          *float32     `json:"MaxArea,omitempty" xml:"MaxArea,omitempty"`
	MaxAreaIndex     *int64       `json:"MaxAreaIndex,omitempty" xml:"MaxAreaIndex,omitempty"`
	MaxDiameter      *float32     `json:"MaxDiameter,omitempty" xml:"MaxDiameter,omitempty"`
	NearestAortaArea *float32     `json:"NearestAortaArea,omitempty" xml:"NearestAortaArea,omitempty"`
}

func (s AnalyzeChestVesselResponseBodyDataPulmonaryInfo) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeChestVesselResponseBodyDataPulmonaryInfo) GoString() string {
	return s.String()
}

func (s *AnalyzeChestVesselResponseBodyDataPulmonaryInfo) SetArea(v []*float32) *AnalyzeChestVesselResponseBodyDataPulmonaryInfo {
	s.Area = v
	return s
}

func (s *AnalyzeChestVesselResponseBodyDataPulmonaryInfo) SetCoordinates(v [][]*float32) *AnalyzeChestVesselResponseBodyDataPulmonaryInfo {
	s.Coordinates = v
	return s
}

func (s *AnalyzeChestVesselResponseBodyDataPulmonaryInfo) SetLabelValue(v int64) *AnalyzeChestVesselResponseBodyDataPulmonaryInfo {
	s.LabelValue = &v
	return s
}

func (s *AnalyzeChestVesselResponseBodyDataPulmonaryInfo) SetMaxArea(v float32) *AnalyzeChestVesselResponseBodyDataPulmonaryInfo {
	s.MaxArea = &v
	return s
}

func (s *AnalyzeChestVesselResponseBodyDataPulmonaryInfo) SetMaxAreaIndex(v int64) *AnalyzeChestVesselResponseBodyDataPulmonaryInfo {
	s.MaxAreaIndex = &v
	return s
}

func (s *AnalyzeChestVesselResponseBodyDataPulmonaryInfo) SetMaxDiameter(v float32) *AnalyzeChestVesselResponseBodyDataPulmonaryInfo {
	s.MaxDiameter = &v
	return s
}

func (s *AnalyzeChestVesselResponseBodyDataPulmonaryInfo) SetNearestAortaArea(v float32) *AnalyzeChestVesselResponseBodyDataPulmonaryInfo {
	s.NearestAortaArea = &v
	return s
}

type AnalyzeChestVesselResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AnalyzeChestVesselResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AnalyzeChestVesselResponse) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeChestVesselResponse) GoString() string {
	return s.String()
}

func (s *AnalyzeChestVesselResponse) SetHeaders(v map[string]*string) *AnalyzeChestVesselResponse {
	s.Headers = v
	return s
}

func (s *AnalyzeChestVesselResponse) SetBody(v *AnalyzeChestVesselResponseBody) *AnalyzeChestVesselResponse {
	s.Body = v
	return s
}

type CalcCACSRequest struct {
	DataFormat     *string                   `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	DataSourceType *string                   `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	OrgId          *string                   `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName        *string                   `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	URLList        []*CalcCACSRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s CalcCACSRequest) String() string {
	return tea.Prettify(s)
}

func (s CalcCACSRequest) GoString() string {
	return s.String()
}

func (s *CalcCACSRequest) SetDataFormat(v string) *CalcCACSRequest {
	s.DataFormat = &v
	return s
}

func (s *CalcCACSRequest) SetDataSourceType(v string) *CalcCACSRequest {
	s.DataSourceType = &v
	return s
}

func (s *CalcCACSRequest) SetOrgId(v string) *CalcCACSRequest {
	s.OrgId = &v
	return s
}

func (s *CalcCACSRequest) SetOrgName(v string) *CalcCACSRequest {
	s.OrgName = &v
	return s
}

func (s *CalcCACSRequest) SetURLList(v []*CalcCACSRequestURLList) *CalcCACSRequest {
	s.URLList = v
	return s
}

type CalcCACSRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s CalcCACSRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s CalcCACSRequestURLList) GoString() string {
	return s.String()
}

func (s *CalcCACSRequestURLList) SetURL(v string) *CalcCACSRequestURLList {
	s.URL = &v
	return s
}

type CalcCACSResponseBody struct {
	Data      *CalcCACSResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CalcCACSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CalcCACSResponseBody) GoString() string {
	return s.String()
}

func (s *CalcCACSResponseBody) SetData(v *CalcCACSResponseBodyData) *CalcCACSResponseBody {
	s.Data = v
	return s
}

func (s *CalcCACSResponseBody) SetRequestId(v string) *CalcCACSResponseBody {
	s.RequestId = &v
	return s
}

type CalcCACSResponseBodyData struct {
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s CalcCACSResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CalcCACSResponseBodyData) GoString() string {
	return s.String()
}

func (s *CalcCACSResponseBodyData) SetResultUrl(v string) *CalcCACSResponseBodyData {
	s.ResultUrl = &v
	return s
}

func (s *CalcCACSResponseBodyData) SetScore(v string) *CalcCACSResponseBodyData {
	s.Score = &v
	return s
}

type CalcCACSResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CalcCACSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CalcCACSResponse) String() string {
	return tea.Prettify(s)
}

func (s CalcCACSResponse) GoString() string {
	return s.String()
}

func (s *CalcCACSResponse) SetHeaders(v map[string]*string) *CalcCACSResponse {
	s.Headers = v
	return s
}

func (s *CalcCACSResponse) SetBody(v *CalcCACSResponseBody) *CalcCACSResponse {
	s.Body = v
	return s
}

type ClassifyFNFRequest struct {
	DataFormat *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	ImageUrl   *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	TracerId   *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s ClassifyFNFRequest) String() string {
	return tea.Prettify(s)
}

func (s ClassifyFNFRequest) GoString() string {
	return s.String()
}

func (s *ClassifyFNFRequest) SetDataFormat(v string) *ClassifyFNFRequest {
	s.DataFormat = &v
	return s
}

func (s *ClassifyFNFRequest) SetImageUrl(v string) *ClassifyFNFRequest {
	s.ImageUrl = &v
	return s
}

func (s *ClassifyFNFRequest) SetOrgId(v string) *ClassifyFNFRequest {
	s.OrgId = &v
	return s
}

func (s *ClassifyFNFRequest) SetOrgName(v string) *ClassifyFNFRequest {
	s.OrgName = &v
	return s
}

func (s *ClassifyFNFRequest) SetTracerId(v string) *ClassifyFNFRequest {
	s.TracerId = &v
	return s
}

type ClassifyFNFAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrlObject,omitempty" xml:"ImageUrlObject,omitempty" require:"true"`
	DataFormat     *string   `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId          *string   `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName        *string   `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	TracerId       *string   `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s ClassifyFNFAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ClassifyFNFAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ClassifyFNFAdvanceRequest) SetImageUrlObject(v io.Reader) *ClassifyFNFAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *ClassifyFNFAdvanceRequest) SetDataFormat(v string) *ClassifyFNFAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *ClassifyFNFAdvanceRequest) SetOrgId(v string) *ClassifyFNFAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *ClassifyFNFAdvanceRequest) SetOrgName(v string) *ClassifyFNFAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *ClassifyFNFAdvanceRequest) SetTracerId(v string) *ClassifyFNFAdvanceRequest {
	s.TracerId = &v
	return s
}

type ClassifyFNFResponseBody struct {
	Data      *ClassifyFNFResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClassifyFNFResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClassifyFNFResponseBody) GoString() string {
	return s.String()
}

func (s *ClassifyFNFResponseBody) SetData(v *ClassifyFNFResponseBodyData) *ClassifyFNFResponseBody {
	s.Data = v
	return s
}

func (s *ClassifyFNFResponseBody) SetRequestId(v string) *ClassifyFNFResponseBody {
	s.RequestId = &v
	return s
}

type ClassifyFNFResponseBodyData struct {
	Fractures []*ClassifyFNFResponseBodyDataFractures `json:"Fractures,omitempty" xml:"Fractures,omitempty" type:"Repeated"`
	ImageUrl  *string                                 `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	OrgId     *string                                 `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName   *string                                 `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
}

func (s ClassifyFNFResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ClassifyFNFResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClassifyFNFResponseBodyData) SetFractures(v []*ClassifyFNFResponseBodyDataFractures) *ClassifyFNFResponseBodyData {
	s.Fractures = v
	return s
}

func (s *ClassifyFNFResponseBodyData) SetImageUrl(v string) *ClassifyFNFResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *ClassifyFNFResponseBodyData) SetOrgId(v string) *ClassifyFNFResponseBodyData {
	s.OrgId = &v
	return s
}

func (s *ClassifyFNFResponseBodyData) SetOrgName(v string) *ClassifyFNFResponseBodyData {
	s.OrgName = &v
	return s
}

type ClassifyFNFResponseBodyDataFractures struct {
	Boxes []*int32                                 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Tag   *ClassifyFNFResponseBodyDataFracturesTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Struct"`
	Value *float32                                 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ClassifyFNFResponseBodyDataFractures) String() string {
	return tea.Prettify(s)
}

func (s ClassifyFNFResponseBodyDataFractures) GoString() string {
	return s.String()
}

func (s *ClassifyFNFResponseBodyDataFractures) SetBoxes(v []*int32) *ClassifyFNFResponseBodyDataFractures {
	s.Boxes = v
	return s
}

func (s *ClassifyFNFResponseBodyDataFractures) SetTag(v *ClassifyFNFResponseBodyDataFracturesTag) *ClassifyFNFResponseBodyDataFractures {
	s.Tag = v
	return s
}

func (s *ClassifyFNFResponseBodyDataFractures) SetValue(v float32) *ClassifyFNFResponseBodyDataFractures {
	s.Value = &v
	return s
}

type ClassifyFNFResponseBodyDataFracturesTag struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ClassifyFNFResponseBodyDataFracturesTag) String() string {
	return tea.Prettify(s)
}

func (s ClassifyFNFResponseBodyDataFracturesTag) GoString() string {
	return s.String()
}

func (s *ClassifyFNFResponseBodyDataFracturesTag) SetLabel(v string) *ClassifyFNFResponseBodyDataFracturesTag {
	s.Label = &v
	return s
}

type ClassifyFNFResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClassifyFNFResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClassifyFNFResponse) String() string {
	return tea.Prettify(s)
}

func (s ClassifyFNFResponse) GoString() string {
	return s.String()
}

func (s *ClassifyFNFResponse) SetHeaders(v map[string]*string) *ClassifyFNFResponse {
	s.Headers = v
	return s
}

func (s *ClassifyFNFResponse) SetBody(v *ClassifyFNFResponseBody) *ClassifyFNFResponse {
	s.Body = v
	return s
}

type DetectCovid19CadRequest struct {
	DataFormat *string                           `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId      *string                           `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string                           `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	URLList    []*DetectCovid19CadRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s DetectCovid19CadRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectCovid19CadRequest) GoString() string {
	return s.String()
}

func (s *DetectCovid19CadRequest) SetDataFormat(v string) *DetectCovid19CadRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectCovid19CadRequest) SetOrgId(v string) *DetectCovid19CadRequest {
	s.OrgId = &v
	return s
}

func (s *DetectCovid19CadRequest) SetOrgName(v string) *DetectCovid19CadRequest {
	s.OrgName = &v
	return s
}

func (s *DetectCovid19CadRequest) SetURLList(v []*DetectCovid19CadRequestURLList) *DetectCovid19CadRequest {
	s.URLList = v
	return s
}

type DetectCovid19CadRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s DetectCovid19CadRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectCovid19CadRequestURLList) GoString() string {
	return s.String()
}

func (s *DetectCovid19CadRequestURLList) SetURL(v string) *DetectCovid19CadRequestURLList {
	s.URL = &v
	return s
}

type DetectCovid19CadResponseBody struct {
	Data      *DetectCovid19CadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectCovid19CadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectCovid19CadResponseBody) GoString() string {
	return s.String()
}

func (s *DetectCovid19CadResponseBody) SetData(v *DetectCovid19CadResponseBodyData) *DetectCovid19CadResponseBody {
	s.Data = v
	return s
}

func (s *DetectCovid19CadResponseBody) SetRequestId(v string) *DetectCovid19CadResponseBody {
	s.RequestId = &v
	return s
}

type DetectCovid19CadResponseBodyData struct {
	LesionRatio       *string `json:"LesionRatio,omitempty" xml:"LesionRatio,omitempty"`
	Mask              *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	NewProbability    *string `json:"NewProbability,omitempty" xml:"NewProbability,omitempty"`
	NormalProbability *string `json:"NormalProbability,omitempty" xml:"NormalProbability,omitempty"`
	OtherProbability  *string `json:"OtherProbability,omitempty" xml:"OtherProbability,omitempty"`
}

func (s DetectCovid19CadResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectCovid19CadResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectCovid19CadResponseBodyData) SetLesionRatio(v string) *DetectCovid19CadResponseBodyData {
	s.LesionRatio = &v
	return s
}

func (s *DetectCovid19CadResponseBodyData) SetMask(v string) *DetectCovid19CadResponseBodyData {
	s.Mask = &v
	return s
}

func (s *DetectCovid19CadResponseBodyData) SetNewProbability(v string) *DetectCovid19CadResponseBodyData {
	s.NewProbability = &v
	return s
}

func (s *DetectCovid19CadResponseBodyData) SetNormalProbability(v string) *DetectCovid19CadResponseBodyData {
	s.NormalProbability = &v
	return s
}

func (s *DetectCovid19CadResponseBodyData) SetOtherProbability(v string) *DetectCovid19CadResponseBodyData {
	s.OtherProbability = &v
	return s
}

type DetectCovid19CadResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectCovid19CadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectCovid19CadResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectCovid19CadResponse) GoString() string {
	return s.String()
}

func (s *DetectCovid19CadResponse) SetHeaders(v map[string]*string) *DetectCovid19CadResponse {
	s.Headers = v
	return s
}

func (s *DetectCovid19CadResponse) SetBody(v *DetectCovid19CadResponseBody) *DetectCovid19CadResponse {
	s.Body = v
	return s
}

type DetectHipKeypointXRayRequest struct {
	DataFormat *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	ImageUrl   *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	TracerId   *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DetectHipKeypointXRayRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectHipKeypointXRayRequest) GoString() string {
	return s.String()
}

func (s *DetectHipKeypointXRayRequest) SetDataFormat(v string) *DetectHipKeypointXRayRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectHipKeypointXRayRequest) SetImageUrl(v string) *DetectHipKeypointXRayRequest {
	s.ImageUrl = &v
	return s
}

func (s *DetectHipKeypointXRayRequest) SetOrgId(v string) *DetectHipKeypointXRayRequest {
	s.OrgId = &v
	return s
}

func (s *DetectHipKeypointXRayRequest) SetOrgName(v string) *DetectHipKeypointXRayRequest {
	s.OrgName = &v
	return s
}

func (s *DetectHipKeypointXRayRequest) SetTracerId(v string) *DetectHipKeypointXRayRequest {
	s.TracerId = &v
	return s
}

type DetectHipKeypointXRayAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrlObject,omitempty" xml:"ImageUrlObject,omitempty" require:"true"`
	DataFormat     *string   `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId          *string   `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName        *string   `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	TracerId       *string   `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DetectHipKeypointXRayAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectHipKeypointXRayAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectHipKeypointXRayAdvanceRequest) SetImageUrlObject(v io.Reader) *DetectHipKeypointXRayAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *DetectHipKeypointXRayAdvanceRequest) SetDataFormat(v string) *DetectHipKeypointXRayAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectHipKeypointXRayAdvanceRequest) SetOrgId(v string) *DetectHipKeypointXRayAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *DetectHipKeypointXRayAdvanceRequest) SetOrgName(v string) *DetectHipKeypointXRayAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *DetectHipKeypointXRayAdvanceRequest) SetTracerId(v string) *DetectHipKeypointXRayAdvanceRequest {
	s.TracerId = &v
	return s
}

type DetectHipKeypointXRayResponseBody struct {
	Data      *DetectHipKeypointXRayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectHipKeypointXRayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectHipKeypointXRayResponseBody) GoString() string {
	return s.String()
}

func (s *DetectHipKeypointXRayResponseBody) SetData(v *DetectHipKeypointXRayResponseBodyData) *DetectHipKeypointXRayResponseBody {
	s.Data = v
	return s
}

func (s *DetectHipKeypointXRayResponseBody) SetRequestId(v string) *DetectHipKeypointXRayResponseBody {
	s.RequestId = &v
	return s
}

type DetectHipKeypointXRayResponseBodyData struct {
	ImageUrl  *string                                           `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	KeyPoints []*DetectHipKeypointXRayResponseBodyDataKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	OrgId     *string                                           `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName   *string                                           `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
}

func (s DetectHipKeypointXRayResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectHipKeypointXRayResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectHipKeypointXRayResponseBodyData) SetImageUrl(v string) *DetectHipKeypointXRayResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *DetectHipKeypointXRayResponseBodyData) SetKeyPoints(v []*DetectHipKeypointXRayResponseBodyDataKeyPoints) *DetectHipKeypointXRayResponseBodyData {
	s.KeyPoints = v
	return s
}

func (s *DetectHipKeypointXRayResponseBodyData) SetOrgId(v string) *DetectHipKeypointXRayResponseBodyData {
	s.OrgId = &v
	return s
}

func (s *DetectHipKeypointXRayResponseBodyData) SetOrgName(v string) *DetectHipKeypointXRayResponseBodyData {
	s.OrgName = &v
	return s
}

type DetectHipKeypointXRayResponseBodyDataKeyPoints struct {
	Coordinates []*int32                                           `json:"Coordinates,omitempty" xml:"Coordinates,omitempty" type:"Repeated"`
	Tag         *DetectHipKeypointXRayResponseBodyDataKeyPointsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Struct"`
	Value       *float32                                           `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DetectHipKeypointXRayResponseBodyDataKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s DetectHipKeypointXRayResponseBodyDataKeyPoints) GoString() string {
	return s.String()
}

func (s *DetectHipKeypointXRayResponseBodyDataKeyPoints) SetCoordinates(v []*int32) *DetectHipKeypointXRayResponseBodyDataKeyPoints {
	s.Coordinates = v
	return s
}

func (s *DetectHipKeypointXRayResponseBodyDataKeyPoints) SetTag(v *DetectHipKeypointXRayResponseBodyDataKeyPointsTag) *DetectHipKeypointXRayResponseBodyDataKeyPoints {
	s.Tag = v
	return s
}

func (s *DetectHipKeypointXRayResponseBodyDataKeyPoints) SetValue(v float32) *DetectHipKeypointXRayResponseBodyDataKeyPoints {
	s.Value = &v
	return s
}

type DetectHipKeypointXRayResponseBodyDataKeyPointsTag struct {
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s DetectHipKeypointXRayResponseBodyDataKeyPointsTag) String() string {
	return tea.Prettify(s)
}

func (s DetectHipKeypointXRayResponseBodyDataKeyPointsTag) GoString() string {
	return s.String()
}

func (s *DetectHipKeypointXRayResponseBodyDataKeyPointsTag) SetDirection(v string) *DetectHipKeypointXRayResponseBodyDataKeyPointsTag {
	s.Direction = &v
	return s
}

func (s *DetectHipKeypointXRayResponseBodyDataKeyPointsTag) SetLabel(v string) *DetectHipKeypointXRayResponseBodyDataKeyPointsTag {
	s.Label = &v
	return s
}

type DetectHipKeypointXRayResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectHipKeypointXRayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectHipKeypointXRayResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectHipKeypointXRayResponse) GoString() string {
	return s.String()
}

func (s *DetectHipKeypointXRayResponse) SetHeaders(v map[string]*string) *DetectHipKeypointXRayResponse {
	s.Headers = v
	return s
}

func (s *DetectHipKeypointXRayResponse) SetBody(v *DetectHipKeypointXRayResponseBody) *DetectHipKeypointXRayResponse {
	s.Body = v
	return s
}

type DetectKneeKeypointXRayRequest struct {
	DataFormat *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	ImageUrl   *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	TracerId   *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DetectKneeKeypointXRayRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeKeypointXRayRequest) GoString() string {
	return s.String()
}

func (s *DetectKneeKeypointXRayRequest) SetDataFormat(v string) *DetectKneeKeypointXRayRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectKneeKeypointXRayRequest) SetImageUrl(v string) *DetectKneeKeypointXRayRequest {
	s.ImageUrl = &v
	return s
}

func (s *DetectKneeKeypointXRayRequest) SetOrgId(v string) *DetectKneeKeypointXRayRequest {
	s.OrgId = &v
	return s
}

func (s *DetectKneeKeypointXRayRequest) SetOrgName(v string) *DetectKneeKeypointXRayRequest {
	s.OrgName = &v
	return s
}

func (s *DetectKneeKeypointXRayRequest) SetTracerId(v string) *DetectKneeKeypointXRayRequest {
	s.TracerId = &v
	return s
}

type DetectKneeKeypointXRayAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrlObject,omitempty" xml:"ImageUrlObject,omitempty" require:"true"`
	DataFormat     *string   `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId          *string   `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName        *string   `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	TracerId       *string   `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DetectKneeKeypointXRayAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeKeypointXRayAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectKneeKeypointXRayAdvanceRequest) SetImageUrlObject(v io.Reader) *DetectKneeKeypointXRayAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *DetectKneeKeypointXRayAdvanceRequest) SetDataFormat(v string) *DetectKneeKeypointXRayAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectKneeKeypointXRayAdvanceRequest) SetOrgId(v string) *DetectKneeKeypointXRayAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *DetectKneeKeypointXRayAdvanceRequest) SetOrgName(v string) *DetectKneeKeypointXRayAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *DetectKneeKeypointXRayAdvanceRequest) SetTracerId(v string) *DetectKneeKeypointXRayAdvanceRequest {
	s.TracerId = &v
	return s
}

type DetectKneeKeypointXRayResponseBody struct {
	Data      *DetectKneeKeypointXRayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectKneeKeypointXRayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeKeypointXRayResponseBody) GoString() string {
	return s.String()
}

func (s *DetectKneeKeypointXRayResponseBody) SetData(v *DetectKneeKeypointXRayResponseBodyData) *DetectKneeKeypointXRayResponseBody {
	s.Data = v
	return s
}

func (s *DetectKneeKeypointXRayResponseBody) SetRequestId(v string) *DetectKneeKeypointXRayResponseBody {
	s.RequestId = &v
	return s
}

type DetectKneeKeypointXRayResponseBodyData struct {
	ImageUrl  *string                                            `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	KeyPoints []*DetectKneeKeypointXRayResponseBodyDataKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	OrgId     *string                                            `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName   *string                                            `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
}

func (s DetectKneeKeypointXRayResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeKeypointXRayResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectKneeKeypointXRayResponseBodyData) SetImageUrl(v string) *DetectKneeKeypointXRayResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *DetectKneeKeypointXRayResponseBodyData) SetKeyPoints(v []*DetectKneeKeypointXRayResponseBodyDataKeyPoints) *DetectKneeKeypointXRayResponseBodyData {
	s.KeyPoints = v
	return s
}

func (s *DetectKneeKeypointXRayResponseBodyData) SetOrgId(v string) *DetectKneeKeypointXRayResponseBodyData {
	s.OrgId = &v
	return s
}

func (s *DetectKneeKeypointXRayResponseBodyData) SetOrgName(v string) *DetectKneeKeypointXRayResponseBodyData {
	s.OrgName = &v
	return s
}

type DetectKneeKeypointXRayResponseBodyDataKeyPoints struct {
	Coordinates []*int32                                            `json:"Coordinates,omitempty" xml:"Coordinates,omitempty" type:"Repeated"`
	Tag         *DetectKneeKeypointXRayResponseBodyDataKeyPointsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Struct"`
	Value       *float32                                            `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DetectKneeKeypointXRayResponseBodyDataKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeKeypointXRayResponseBodyDataKeyPoints) GoString() string {
	return s.String()
}

func (s *DetectKneeKeypointXRayResponseBodyDataKeyPoints) SetCoordinates(v []*int32) *DetectKneeKeypointXRayResponseBodyDataKeyPoints {
	s.Coordinates = v
	return s
}

func (s *DetectKneeKeypointXRayResponseBodyDataKeyPoints) SetTag(v *DetectKneeKeypointXRayResponseBodyDataKeyPointsTag) *DetectKneeKeypointXRayResponseBodyDataKeyPoints {
	s.Tag = v
	return s
}

func (s *DetectKneeKeypointXRayResponseBodyDataKeyPoints) SetValue(v float32) *DetectKneeKeypointXRayResponseBodyDataKeyPoints {
	s.Value = &v
	return s
}

type DetectKneeKeypointXRayResponseBodyDataKeyPointsTag struct {
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s DetectKneeKeypointXRayResponseBodyDataKeyPointsTag) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeKeypointXRayResponseBodyDataKeyPointsTag) GoString() string {
	return s.String()
}

func (s *DetectKneeKeypointXRayResponseBodyDataKeyPointsTag) SetDirection(v string) *DetectKneeKeypointXRayResponseBodyDataKeyPointsTag {
	s.Direction = &v
	return s
}

func (s *DetectKneeKeypointXRayResponseBodyDataKeyPointsTag) SetLabel(v string) *DetectKneeKeypointXRayResponseBodyDataKeyPointsTag {
	s.Label = &v
	return s
}

type DetectKneeKeypointXRayResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectKneeKeypointXRayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectKneeKeypointXRayResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeKeypointXRayResponse) GoString() string {
	return s.String()
}

func (s *DetectKneeKeypointXRayResponse) SetHeaders(v map[string]*string) *DetectKneeKeypointXRayResponse {
	s.Headers = v
	return s
}

func (s *DetectKneeKeypointXRayResponse) SetBody(v *DetectKneeKeypointXRayResponseBody) *DetectKneeKeypointXRayResponse {
	s.Body = v
	return s
}

type DetectKneeXRayRequest struct {
	DataFormat *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	Url        *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DetectKneeXRayRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeXRayRequest) GoString() string {
	return s.String()
}

func (s *DetectKneeXRayRequest) SetDataFormat(v string) *DetectKneeXRayRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectKneeXRayRequest) SetOrgId(v string) *DetectKneeXRayRequest {
	s.OrgId = &v
	return s
}

func (s *DetectKneeXRayRequest) SetOrgName(v string) *DetectKneeXRayRequest {
	s.OrgName = &v
	return s
}

func (s *DetectKneeXRayRequest) SetUrl(v string) *DetectKneeXRayRequest {
	s.Url = &v
	return s
}

type DetectKneeXRayAdvanceRequest struct {
	UrlObject  io.Reader `json:"UrlObject,omitempty" xml:"UrlObject,omitempty" require:"true"`
	DataFormat *string   `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId      *string   `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string   `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
}

func (s DetectKneeXRayAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeXRayAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectKneeXRayAdvanceRequest) SetUrlObject(v io.Reader) *DetectKneeXRayAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *DetectKneeXRayAdvanceRequest) SetDataFormat(v string) *DetectKneeXRayAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectKneeXRayAdvanceRequest) SetOrgId(v string) *DetectKneeXRayAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *DetectKneeXRayAdvanceRequest) SetOrgName(v string) *DetectKneeXRayAdvanceRequest {
	s.OrgName = &v
	return s
}

type DetectKneeXRayResponseBody struct {
	Data      *DetectKneeXRayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectKneeXRayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeXRayResponseBody) GoString() string {
	return s.String()
}

func (s *DetectKneeXRayResponseBody) SetData(v *DetectKneeXRayResponseBodyData) *DetectKneeXRayResponseBody {
	s.Data = v
	return s
}

func (s *DetectKneeXRayResponseBody) SetRequestId(v string) *DetectKneeXRayResponseBody {
	s.RequestId = &v
	return s
}

type DetectKneeXRayResponseBodyData struct {
	KLDetections []*DetectKneeXRayResponseBodyDataKLDetections `json:"KLDetections,omitempty" xml:"KLDetections,omitempty" type:"Repeated"`
}

func (s DetectKneeXRayResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeXRayResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectKneeXRayResponseBodyData) SetKLDetections(v []*DetectKneeXRayResponseBodyDataKLDetections) *DetectKneeXRayResponseBodyData {
	s.KLDetections = v
	return s
}

type DetectKneeXRayResponseBodyDataKLDetections struct {
	Detections []*float32 `json:"Detections,omitempty" xml:"Detections,omitempty" type:"Repeated"`
}

func (s DetectKneeXRayResponseBodyDataKLDetections) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeXRayResponseBodyDataKLDetections) GoString() string {
	return s.String()
}

func (s *DetectKneeXRayResponseBodyDataKLDetections) SetDetections(v []*float32) *DetectKneeXRayResponseBodyDataKLDetections {
	s.Detections = v
	return s
}

type DetectKneeXRayResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectKneeXRayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectKneeXRayResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeXRayResponse) GoString() string {
	return s.String()
}

func (s *DetectKneeXRayResponse) SetHeaders(v map[string]*string) *DetectKneeXRayResponse {
	s.Headers = v
	return s
}

func (s *DetectKneeXRayResponse) SetBody(v *DetectKneeXRayResponseBody) *DetectKneeXRayResponse {
	s.Body = v
	return s
}

type DetectLungNoduleRequest struct {
	DataFormat *string                           `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId      *string                           `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string                           `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	Threshold  *float32                          `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	URLList    []*DetectLungNoduleRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s DetectLungNoduleRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectLungNoduleRequest) GoString() string {
	return s.String()
}

func (s *DetectLungNoduleRequest) SetDataFormat(v string) *DetectLungNoduleRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectLungNoduleRequest) SetOrgId(v string) *DetectLungNoduleRequest {
	s.OrgId = &v
	return s
}

func (s *DetectLungNoduleRequest) SetOrgName(v string) *DetectLungNoduleRequest {
	s.OrgName = &v
	return s
}

func (s *DetectLungNoduleRequest) SetThreshold(v float32) *DetectLungNoduleRequest {
	s.Threshold = &v
	return s
}

func (s *DetectLungNoduleRequest) SetURLList(v []*DetectLungNoduleRequestURLList) *DetectLungNoduleRequest {
	s.URLList = v
	return s
}

type DetectLungNoduleRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s DetectLungNoduleRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectLungNoduleRequestURLList) GoString() string {
	return s.String()
}

func (s *DetectLungNoduleRequestURLList) SetURL(v string) *DetectLungNoduleRequestURLList {
	s.URL = &v
	return s
}

type DetectLungNoduleResponseBody struct {
	Data      *DetectLungNoduleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectLungNoduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectLungNoduleResponseBody) GoString() string {
	return s.String()
}

func (s *DetectLungNoduleResponseBody) SetData(v *DetectLungNoduleResponseBodyData) *DetectLungNoduleResponseBody {
	s.Data = v
	return s
}

func (s *DetectLungNoduleResponseBody) SetRequestId(v string) *DetectLungNoduleResponseBody {
	s.RequestId = &v
	return s
}

type DetectLungNoduleResponseBodyData struct {
	Series []*DetectLungNoduleResponseBodyDataSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
}

func (s DetectLungNoduleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectLungNoduleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectLungNoduleResponseBodyData) SetSeries(v []*DetectLungNoduleResponseBodyDataSeries) *DetectLungNoduleResponseBodyData {
	s.Series = v
	return s
}

type DetectLungNoduleResponseBodyDataSeries struct {
	Elements          []*DetectLungNoduleResponseBodyDataSeriesElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Origin            []*float32                                        `json:"Origin,omitempty" xml:"Origin,omitempty" type:"Repeated"`
	Report            *string                                           `json:"Report,omitempty" xml:"Report,omitempty"`
	SeriesInstanceUid *string                                           `json:"SeriesInstanceUid,omitempty" xml:"SeriesInstanceUid,omitempty"`
	Spacing           []*float32                                        `json:"Spacing,omitempty" xml:"Spacing,omitempty" type:"Repeated"`
}

func (s DetectLungNoduleResponseBodyDataSeries) String() string {
	return tea.Prettify(s)
}

func (s DetectLungNoduleResponseBodyDataSeries) GoString() string {
	return s.String()
}

func (s *DetectLungNoduleResponseBodyDataSeries) SetElements(v []*DetectLungNoduleResponseBodyDataSeriesElements) *DetectLungNoduleResponseBodyDataSeries {
	s.Elements = v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeries) SetOrigin(v []*float32) *DetectLungNoduleResponseBodyDataSeries {
	s.Origin = v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeries) SetReport(v string) *DetectLungNoduleResponseBodyDataSeries {
	s.Report = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeries) SetSeriesInstanceUid(v string) *DetectLungNoduleResponseBodyDataSeries {
	s.SeriesInstanceUid = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeries) SetSpacing(v []*float32) *DetectLungNoduleResponseBodyDataSeries {
	s.Spacing = v
	return s
}

type DetectLungNoduleResponseBodyDataSeriesElements struct {
	Category       *string  `json:"Category,omitempty" xml:"Category,omitempty"`
	Confidence     *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Diameter       *float32 `json:"Diameter,omitempty" xml:"Diameter,omitempty"`
	ImageX         *float32 `json:"ImageX,omitempty" xml:"ImageX,omitempty"`
	ImageY         *float32 `json:"ImageY,omitempty" xml:"ImageY,omitempty"`
	ImageZ         *float32 `json:"ImageZ,omitempty" xml:"ImageZ,omitempty"`
	Lobe           *string  `json:"Lobe,omitempty" xml:"Lobe,omitempty"`
	Lung           *string  `json:"Lung,omitempty" xml:"Lung,omitempty"`
	MeanValue      *float32 `json:"MeanValue,omitempty" xml:"MeanValue,omitempty"`
	SOPInstanceUID *string  `json:"SOPInstanceUID,omitempty" xml:"SOPInstanceUID,omitempty"`
	Volume         *float32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
	X              *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y              *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Z              *float32 `json:"Z,omitempty" xml:"Z,omitempty"`
}

func (s DetectLungNoduleResponseBodyDataSeriesElements) String() string {
	return tea.Prettify(s)
}

func (s DetectLungNoduleResponseBodyDataSeriesElements) GoString() string {
	return s.String()
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetCategory(v string) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.Category = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetConfidence(v float32) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.Confidence = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetDiameter(v float32) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.Diameter = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetImageX(v float32) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.ImageX = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetImageY(v float32) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.ImageY = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetImageZ(v float32) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.ImageZ = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetLobe(v string) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.Lobe = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetLung(v string) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.Lung = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetMeanValue(v float32) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.MeanValue = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetSOPInstanceUID(v string) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.SOPInstanceUID = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetVolume(v float32) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.Volume = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetX(v float32) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.X = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetY(v float32) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.Y = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetZ(v float32) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.Z = &v
	return s
}

type DetectLungNoduleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectLungNoduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectLungNoduleResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectLungNoduleResponse) GoString() string {
	return s.String()
}

func (s *DetectLungNoduleResponse) SetHeaders(v map[string]*string) *DetectLungNoduleResponse {
	s.Headers = v
	return s
}

func (s *DetectLungNoduleResponse) SetBody(v *DetectLungNoduleResponseBody) *DetectLungNoduleResponse {
	s.Body = v
	return s
}

type DetectRibFractureRequest struct {
	DataFormat *string                            `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId      *string                            `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string                            `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	SourceType *string                            `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	URLList    []*DetectRibFractureRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s DetectRibFractureRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectRibFractureRequest) GoString() string {
	return s.String()
}

func (s *DetectRibFractureRequest) SetDataFormat(v string) *DetectRibFractureRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectRibFractureRequest) SetOrgId(v string) *DetectRibFractureRequest {
	s.OrgId = &v
	return s
}

func (s *DetectRibFractureRequest) SetOrgName(v string) *DetectRibFractureRequest {
	s.OrgName = &v
	return s
}

func (s *DetectRibFractureRequest) SetSourceType(v string) *DetectRibFractureRequest {
	s.SourceType = &v
	return s
}

func (s *DetectRibFractureRequest) SetURLList(v []*DetectRibFractureRequestURLList) *DetectRibFractureRequest {
	s.URLList = v
	return s
}

type DetectRibFractureRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s DetectRibFractureRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectRibFractureRequestURLList) GoString() string {
	return s.String()
}

func (s *DetectRibFractureRequestURLList) SetURL(v string) *DetectRibFractureRequestURLList {
	s.URL = &v
	return s
}

type DetectRibFractureResponseBody struct {
	Data      *DetectRibFractureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectRibFractureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectRibFractureResponseBody) GoString() string {
	return s.String()
}

func (s *DetectRibFractureResponseBody) SetData(v *DetectRibFractureResponseBodyData) *DetectRibFractureResponseBody {
	s.Data = v
	return s
}

func (s *DetectRibFractureResponseBody) SetRequestId(v string) *DetectRibFractureResponseBody {
	s.RequestId = &v
	return s
}

type DetectRibFractureResponseBodyData struct {
	Detections []*DetectRibFractureResponseBodyDataDetections `json:"Detections,omitempty" xml:"Detections,omitempty" type:"Repeated"`
	Origin     []*float32                                     `json:"Origin,omitempty" xml:"Origin,omitempty" type:"Repeated"`
	ResultURL  *string                                        `json:"ResultURL,omitempty" xml:"ResultURL,omitempty"`
	Spacing    []*float32                                     `json:"Spacing,omitempty" xml:"Spacing,omitempty" type:"Repeated"`
}

func (s DetectRibFractureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectRibFractureResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectRibFractureResponseBodyData) SetDetections(v []*DetectRibFractureResponseBodyDataDetections) *DetectRibFractureResponseBodyData {
	s.Detections = v
	return s
}

func (s *DetectRibFractureResponseBodyData) SetOrigin(v []*float32) *DetectRibFractureResponseBodyData {
	s.Origin = v
	return s
}

func (s *DetectRibFractureResponseBodyData) SetResultURL(v string) *DetectRibFractureResponseBodyData {
	s.ResultURL = &v
	return s
}

func (s *DetectRibFractureResponseBodyData) SetSpacing(v []*float32) *DetectRibFractureResponseBodyData {
	s.Spacing = v
	return s
}

type DetectRibFractureResponseBodyDataDetections struct {
	CoordinateImage    []*int32 `json:"CoordinateImage,omitempty" xml:"CoordinateImage,omitempty" type:"Repeated"`
	Coordinates        []*int32 `json:"Coordinates,omitempty" xml:"Coordinates,omitempty" type:"Repeated"`
	FractureCategory   *string  `json:"FractureCategory,omitempty" xml:"FractureCategory,omitempty"`
	FractureConfidence *float32 `json:"FractureConfidence,omitempty" xml:"FractureConfidence,omitempty"`
	FractureId         *int32   `json:"FractureId,omitempty" xml:"FractureId,omitempty"`
	FractureLocation   *string  `json:"FractureLocation,omitempty" xml:"FractureLocation,omitempty"`
	FractureSegment    *int64   `json:"FractureSegment,omitempty" xml:"FractureSegment,omitempty"`
}

func (s DetectRibFractureResponseBodyDataDetections) String() string {
	return tea.Prettify(s)
}

func (s DetectRibFractureResponseBodyDataDetections) GoString() string {
	return s.String()
}

func (s *DetectRibFractureResponseBodyDataDetections) SetCoordinateImage(v []*int32) *DetectRibFractureResponseBodyDataDetections {
	s.CoordinateImage = v
	return s
}

func (s *DetectRibFractureResponseBodyDataDetections) SetCoordinates(v []*int32) *DetectRibFractureResponseBodyDataDetections {
	s.Coordinates = v
	return s
}

func (s *DetectRibFractureResponseBodyDataDetections) SetFractureCategory(v string) *DetectRibFractureResponseBodyDataDetections {
	s.FractureCategory = &v
	return s
}

func (s *DetectRibFractureResponseBodyDataDetections) SetFractureConfidence(v float32) *DetectRibFractureResponseBodyDataDetections {
	s.FractureConfidence = &v
	return s
}

func (s *DetectRibFractureResponseBodyDataDetections) SetFractureId(v int32) *DetectRibFractureResponseBodyDataDetections {
	s.FractureId = &v
	return s
}

func (s *DetectRibFractureResponseBodyDataDetections) SetFractureLocation(v string) *DetectRibFractureResponseBodyDataDetections {
	s.FractureLocation = &v
	return s
}

func (s *DetectRibFractureResponseBodyDataDetections) SetFractureSegment(v int64) *DetectRibFractureResponseBodyDataDetections {
	s.FractureSegment = &v
	return s
}

type DetectRibFractureResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectRibFractureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectRibFractureResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectRibFractureResponse) GoString() string {
	return s.String()
}

func (s *DetectRibFractureResponse) SetHeaders(v map[string]*string) *DetectRibFractureResponse {
	s.Headers = v
	return s
}

func (s *DetectRibFractureResponse) SetBody(v *DetectRibFractureResponseBody) *DetectRibFractureResponse {
	s.Body = v
	return s
}

type DetectSkinDiseaseRequest struct {
	OrgId   *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	Url     *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DetectSkinDiseaseRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectSkinDiseaseRequest) GoString() string {
	return s.String()
}

func (s *DetectSkinDiseaseRequest) SetOrgId(v string) *DetectSkinDiseaseRequest {
	s.OrgId = &v
	return s
}

func (s *DetectSkinDiseaseRequest) SetOrgName(v string) *DetectSkinDiseaseRequest {
	s.OrgName = &v
	return s
}

func (s *DetectSkinDiseaseRequest) SetUrl(v string) *DetectSkinDiseaseRequest {
	s.Url = &v
	return s
}

type DetectSkinDiseaseAdvanceRequest struct {
	UrlObject io.Reader `json:"UrlObject,omitempty" xml:"UrlObject,omitempty" require:"true"`
	OrgId     *string   `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName   *string   `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
}

func (s DetectSkinDiseaseAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectSkinDiseaseAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectSkinDiseaseAdvanceRequest) SetUrlObject(v io.Reader) *DetectSkinDiseaseAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *DetectSkinDiseaseAdvanceRequest) SetOrgId(v string) *DetectSkinDiseaseAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *DetectSkinDiseaseAdvanceRequest) SetOrgName(v string) *DetectSkinDiseaseAdvanceRequest {
	s.OrgName = &v
	return s
}

type DetectSkinDiseaseResponseBody struct {
	Data      *DetectSkinDiseaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectSkinDiseaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectSkinDiseaseResponseBody) GoString() string {
	return s.String()
}

func (s *DetectSkinDiseaseResponseBody) SetData(v *DetectSkinDiseaseResponseBodyData) *DetectSkinDiseaseResponseBody {
	s.Data = v
	return s
}

func (s *DetectSkinDiseaseResponseBody) SetRequestId(v string) *DetectSkinDiseaseResponseBody {
	s.RequestId = &v
	return s
}

type DetectSkinDiseaseResponseBodyData struct {
	Results map[string]interface{} `json:"Results,omitempty" xml:"Results,omitempty"`
}

func (s DetectSkinDiseaseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectSkinDiseaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectSkinDiseaseResponseBodyData) SetResults(v map[string]interface{}) *DetectSkinDiseaseResponseBodyData {
	s.Results = v
	return s
}

type DetectSkinDiseaseResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectSkinDiseaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectSkinDiseaseResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectSkinDiseaseResponse) GoString() string {
	return s.String()
}

func (s *DetectSkinDiseaseResponse) SetHeaders(v map[string]*string) *DetectSkinDiseaseResponse {
	s.Headers = v
	return s
}

func (s *DetectSkinDiseaseResponse) SetBody(v *DetectSkinDiseaseResponseBody) *DetectSkinDiseaseResponse {
	s.Body = v
	return s
}

type DetectSpineMRIRequest struct {
	DataFormat *string                         `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId      *string                         `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string                         `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	URLList    []*DetectSpineMRIRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s DetectSpineMRIRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectSpineMRIRequest) GoString() string {
	return s.String()
}

func (s *DetectSpineMRIRequest) SetDataFormat(v string) *DetectSpineMRIRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectSpineMRIRequest) SetOrgId(v string) *DetectSpineMRIRequest {
	s.OrgId = &v
	return s
}

func (s *DetectSpineMRIRequest) SetOrgName(v string) *DetectSpineMRIRequest {
	s.OrgName = &v
	return s
}

func (s *DetectSpineMRIRequest) SetURLList(v []*DetectSpineMRIRequestURLList) *DetectSpineMRIRequest {
	s.URLList = v
	return s
}

type DetectSpineMRIRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s DetectSpineMRIRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectSpineMRIRequestURLList) GoString() string {
	return s.String()
}

func (s *DetectSpineMRIRequestURLList) SetURL(v string) *DetectSpineMRIRequestURLList {
	s.URL = &v
	return s
}

type DetectSpineMRIResponseBody struct {
	Data      *DetectSpineMRIResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectSpineMRIResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectSpineMRIResponseBody) GoString() string {
	return s.String()
}

func (s *DetectSpineMRIResponseBody) SetData(v *DetectSpineMRIResponseBodyData) *DetectSpineMRIResponseBody {
	s.Data = v
	return s
}

func (s *DetectSpineMRIResponseBody) SetRequestId(v string) *DetectSpineMRIResponseBody {
	s.RequestId = &v
	return s
}

type DetectSpineMRIResponseBodyData struct {
	Discs     []*DetectSpineMRIResponseBodyDataDiscs     `json:"Discs,omitempty" xml:"Discs,omitempty" type:"Repeated"`
	Vertebras []*DetectSpineMRIResponseBodyDataVertebras `json:"Vertebras,omitempty" xml:"Vertebras,omitempty" type:"Repeated"`
}

func (s DetectSpineMRIResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectSpineMRIResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectSpineMRIResponseBodyData) SetDiscs(v []*DetectSpineMRIResponseBodyDataDiscs) *DetectSpineMRIResponseBodyData {
	s.Discs = v
	return s
}

func (s *DetectSpineMRIResponseBodyData) SetVertebras(v []*DetectSpineMRIResponseBodyDataVertebras) *DetectSpineMRIResponseBodyData {
	s.Vertebras = v
	return s
}

type DetectSpineMRIResponseBodyDataDiscs struct {
	Disease        *string    `json:"Disease,omitempty" xml:"Disease,omitempty"`
	Identification *string    `json:"Identification,omitempty" xml:"Identification,omitempty"`
	Location       []*float32 `json:"Location,omitempty" xml:"Location,omitempty" type:"Repeated"`
}

func (s DetectSpineMRIResponseBodyDataDiscs) String() string {
	return tea.Prettify(s)
}

func (s DetectSpineMRIResponseBodyDataDiscs) GoString() string {
	return s.String()
}

func (s *DetectSpineMRIResponseBodyDataDiscs) SetDisease(v string) *DetectSpineMRIResponseBodyDataDiscs {
	s.Disease = &v
	return s
}

func (s *DetectSpineMRIResponseBodyDataDiscs) SetIdentification(v string) *DetectSpineMRIResponseBodyDataDiscs {
	s.Identification = &v
	return s
}

func (s *DetectSpineMRIResponseBodyDataDiscs) SetLocation(v []*float32) *DetectSpineMRIResponseBodyDataDiscs {
	s.Location = v
	return s
}

type DetectSpineMRIResponseBodyDataVertebras struct {
	Disease        *string    `json:"Disease,omitempty" xml:"Disease,omitempty"`
	Identification *string    `json:"Identification,omitempty" xml:"Identification,omitempty"`
	Location       []*float32 `json:"Location,omitempty" xml:"Location,omitempty" type:"Repeated"`
}

func (s DetectSpineMRIResponseBodyDataVertebras) String() string {
	return tea.Prettify(s)
}

func (s DetectSpineMRIResponseBodyDataVertebras) GoString() string {
	return s.String()
}

func (s *DetectSpineMRIResponseBodyDataVertebras) SetDisease(v string) *DetectSpineMRIResponseBodyDataVertebras {
	s.Disease = &v
	return s
}

func (s *DetectSpineMRIResponseBodyDataVertebras) SetIdentification(v string) *DetectSpineMRIResponseBodyDataVertebras {
	s.Identification = &v
	return s
}

func (s *DetectSpineMRIResponseBodyDataVertebras) SetLocation(v []*float32) *DetectSpineMRIResponseBodyDataVertebras {
	s.Location = v
	return s
}

type DetectSpineMRIResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectSpineMRIResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectSpineMRIResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectSpineMRIResponse) GoString() string {
	return s.String()
}

func (s *DetectSpineMRIResponse) SetHeaders(v map[string]*string) *DetectSpineMRIResponse {
	s.Headers = v
	return s
}

func (s *DetectSpineMRIResponse) SetBody(v *DetectSpineMRIResponseBody) *DetectSpineMRIResponse {
	s.Body = v
	return s
}

type GetAsyncJobResultRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultRequest) SetJobId(v string) *GetAsyncJobResultRequest {
	s.JobId = &v
	return s
}

type GetAsyncJobResultResponseBody struct {
	Data      *GetAsyncJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAsyncJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBody) SetData(v *GetAsyncJobResultResponseBodyData) *GetAsyncJobResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAsyncJobResultResponseBody) SetRequestId(v string) *GetAsyncJobResultResponseBody {
	s.RequestId = &v
	return s
}

type GetAsyncJobResultResponseBodyData struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAsyncJobResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorCode(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorMessage(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetJobId(v string) *GetAsyncJobResultResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetResult(v string) *GetAsyncJobResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetStatus(v string) *GetAsyncJobResultResponseBodyData {
	s.Status = &v
	return s
}

type GetAsyncJobResultResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAsyncJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponse) SetHeaders(v map[string]*string) *GetAsyncJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncJobResultResponse) SetBody(v *GetAsyncJobResultResponseBody) *GetAsyncJobResultResponse {
	s.Body = v
	return s
}

type RunCTRegistrationRequest struct {
	DataFormat     *string                                  `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	DataSourceType *string                                  `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	FloatingList   []*RunCTRegistrationRequestFloatingList  `json:"FloatingList,omitempty" xml:"FloatingList,omitempty" type:"Repeated"`
	OrgId          *string                                  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName        *string                                  `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	ReferenceList  []*RunCTRegistrationRequestReferenceList `json:"ReferenceList,omitempty" xml:"ReferenceList,omitempty" type:"Repeated"`
}

func (s RunCTRegistrationRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCTRegistrationRequest) GoString() string {
	return s.String()
}

func (s *RunCTRegistrationRequest) SetDataFormat(v string) *RunCTRegistrationRequest {
	s.DataFormat = &v
	return s
}

func (s *RunCTRegistrationRequest) SetDataSourceType(v string) *RunCTRegistrationRequest {
	s.DataSourceType = &v
	return s
}

func (s *RunCTRegistrationRequest) SetFloatingList(v []*RunCTRegistrationRequestFloatingList) *RunCTRegistrationRequest {
	s.FloatingList = v
	return s
}

func (s *RunCTRegistrationRequest) SetOrgId(v string) *RunCTRegistrationRequest {
	s.OrgId = &v
	return s
}

func (s *RunCTRegistrationRequest) SetOrgName(v string) *RunCTRegistrationRequest {
	s.OrgName = &v
	return s
}

func (s *RunCTRegistrationRequest) SetReferenceList(v []*RunCTRegistrationRequestReferenceList) *RunCTRegistrationRequest {
	s.ReferenceList = v
	return s
}

type RunCTRegistrationRequestFloatingList struct {
	FloatingURL *string `json:"FloatingURL,omitempty" xml:"FloatingURL,omitempty"`
}

func (s RunCTRegistrationRequestFloatingList) String() string {
	return tea.Prettify(s)
}

func (s RunCTRegistrationRequestFloatingList) GoString() string {
	return s.String()
}

func (s *RunCTRegistrationRequestFloatingList) SetFloatingURL(v string) *RunCTRegistrationRequestFloatingList {
	s.FloatingURL = &v
	return s
}

type RunCTRegistrationRequestReferenceList struct {
	ReferenceURL *string `json:"ReferenceURL,omitempty" xml:"ReferenceURL,omitempty"`
}

func (s RunCTRegistrationRequestReferenceList) String() string {
	return tea.Prettify(s)
}

func (s RunCTRegistrationRequestReferenceList) GoString() string {
	return s.String()
}

func (s *RunCTRegistrationRequestReferenceList) SetReferenceURL(v string) *RunCTRegistrationRequestReferenceList {
	s.ReferenceURL = &v
	return s
}

type RunCTRegistrationResponseBody struct {
	Data      *RunCTRegistrationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCTRegistrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCTRegistrationResponseBody) GoString() string {
	return s.String()
}

func (s *RunCTRegistrationResponseBody) SetData(v *RunCTRegistrationResponseBodyData) *RunCTRegistrationResponseBody {
	s.Data = v
	return s
}

func (s *RunCTRegistrationResponseBody) SetRequestId(v string) *RunCTRegistrationResponseBody {
	s.RequestId = &v
	return s
}

type RunCTRegistrationResponseBodyData struct {
	DUrl *string `json:"DUrl,omitempty" xml:"DUrl,omitempty"`
	NUrl *string `json:"NUrl,omitempty" xml:"NUrl,omitempty"`
}

func (s RunCTRegistrationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RunCTRegistrationResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunCTRegistrationResponseBodyData) SetDUrl(v string) *RunCTRegistrationResponseBodyData {
	s.DUrl = &v
	return s
}

func (s *RunCTRegistrationResponseBodyData) SetNUrl(v string) *RunCTRegistrationResponseBodyData {
	s.NUrl = &v
	return s
}

type RunCTRegistrationResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunCTRegistrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunCTRegistrationResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCTRegistrationResponse) GoString() string {
	return s.String()
}

func (s *RunCTRegistrationResponse) SetHeaders(v map[string]*string) *RunCTRegistrationResponse {
	s.Headers = v
	return s
}

func (s *RunCTRegistrationResponse) SetBody(v *RunCTRegistrationResponseBody) *RunCTRegistrationResponse {
	s.Body = v
	return s
}

type RunMedQARequest struct {
	AnswerImageDataList []*RunMedQARequestAnswerImageDataList `json:"AnswerImageDataList,omitempty" xml:"AnswerImageDataList,omitempty" type:"Repeated"`
	AnswerImageURLList  []*RunMedQARequestAnswerImageURLList  `json:"AnswerImageURLList,omitempty" xml:"AnswerImageURLList,omitempty" type:"Repeated"`
	AnswerTextList      []*RunMedQARequestAnswerTextList      `json:"AnswerTextList,omitempty" xml:"AnswerTextList,omitempty" type:"Repeated"`
	Department          *string                               `json:"Department,omitempty" xml:"Department,omitempty"`
	OrgId               *string                               `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName             *string                               `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	QuestionType        *string                               `json:"QuestionType,omitempty" xml:"QuestionType,omitempty"`
	SessionId           *string                               `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s RunMedQARequest) String() string {
	return tea.Prettify(s)
}

func (s RunMedQARequest) GoString() string {
	return s.String()
}

func (s *RunMedQARequest) SetAnswerImageDataList(v []*RunMedQARequestAnswerImageDataList) *RunMedQARequest {
	s.AnswerImageDataList = v
	return s
}

func (s *RunMedQARequest) SetAnswerImageURLList(v []*RunMedQARequestAnswerImageURLList) *RunMedQARequest {
	s.AnswerImageURLList = v
	return s
}

func (s *RunMedQARequest) SetAnswerTextList(v []*RunMedQARequestAnswerTextList) *RunMedQARequest {
	s.AnswerTextList = v
	return s
}

func (s *RunMedQARequest) SetDepartment(v string) *RunMedQARequest {
	s.Department = &v
	return s
}

func (s *RunMedQARequest) SetOrgId(v string) *RunMedQARequest {
	s.OrgId = &v
	return s
}

func (s *RunMedQARequest) SetOrgName(v string) *RunMedQARequest {
	s.OrgName = &v
	return s
}

func (s *RunMedQARequest) SetQuestionType(v string) *RunMedQARequest {
	s.QuestionType = &v
	return s
}

func (s *RunMedQARequest) SetSessionId(v string) *RunMedQARequest {
	s.SessionId = &v
	return s
}

type RunMedQARequestAnswerImageDataList struct {
	AnswerImageData []byte `json:"AnswerImageData,omitempty" xml:"AnswerImageData,omitempty"`
}

func (s RunMedQARequestAnswerImageDataList) String() string {
	return tea.Prettify(s)
}

func (s RunMedQARequestAnswerImageDataList) GoString() string {
	return s.String()
}

func (s *RunMedQARequestAnswerImageDataList) SetAnswerImageData(v []byte) *RunMedQARequestAnswerImageDataList {
	s.AnswerImageData = v
	return s
}

type RunMedQARequestAnswerImageURLList struct {
	AnswerImageURL *string `json:"AnswerImageURL,omitempty" xml:"AnswerImageURL,omitempty"`
}

func (s RunMedQARequestAnswerImageURLList) String() string {
	return tea.Prettify(s)
}

func (s RunMedQARequestAnswerImageURLList) GoString() string {
	return s.String()
}

func (s *RunMedQARequestAnswerImageURLList) SetAnswerImageURL(v string) *RunMedQARequestAnswerImageURLList {
	s.AnswerImageURL = &v
	return s
}

type RunMedQARequestAnswerTextList struct {
	AnswerText *string `json:"AnswerText,omitempty" xml:"AnswerText,omitempty"`
}

func (s RunMedQARequestAnswerTextList) String() string {
	return tea.Prettify(s)
}

func (s RunMedQARequestAnswerTextList) GoString() string {
	return s.String()
}

func (s *RunMedQARequestAnswerTextList) SetAnswerText(v string) *RunMedQARequestAnswerTextList {
	s.AnswerText = &v
	return s
}

type RunMedQAResponseBody struct {
	Data      *RunMedQAResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunMedQAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunMedQAResponseBody) GoString() string {
	return s.String()
}

func (s *RunMedQAResponseBody) SetData(v *RunMedQAResponseBodyData) *RunMedQAResponseBody {
	s.Data = v
	return s
}

func (s *RunMedQAResponseBody) SetRequestId(v string) *RunMedQAResponseBody {
	s.RequestId = &v
	return s
}

type RunMedQAResponseBodyData struct {
	AnswerType   *string            `json:"AnswerType,omitempty" xml:"AnswerType,omitempty"`
	Options      []*string          `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
	Question     *string            `json:"Question,omitempty" xml:"Question,omitempty"`
	QuestionType *string            `json:"QuestionType,omitempty" xml:"QuestionType,omitempty"`
	Reports      map[string]*string `json:"Reports,omitempty" xml:"Reports,omitempty"`
	SessionId    *string            `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s RunMedQAResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RunMedQAResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunMedQAResponseBodyData) SetAnswerType(v string) *RunMedQAResponseBodyData {
	s.AnswerType = &v
	return s
}

func (s *RunMedQAResponseBodyData) SetOptions(v []*string) *RunMedQAResponseBodyData {
	s.Options = v
	return s
}

func (s *RunMedQAResponseBodyData) SetQuestion(v string) *RunMedQAResponseBodyData {
	s.Question = &v
	return s
}

func (s *RunMedQAResponseBodyData) SetQuestionType(v string) *RunMedQAResponseBodyData {
	s.QuestionType = &v
	return s
}

func (s *RunMedQAResponseBodyData) SetReports(v map[string]*string) *RunMedQAResponseBodyData {
	s.Reports = v
	return s
}

func (s *RunMedQAResponseBodyData) SetSessionId(v string) *RunMedQAResponseBodyData {
	s.SessionId = &v
	return s
}

type RunMedQAResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunMedQAResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunMedQAResponse) String() string {
	return tea.Prettify(s)
}

func (s RunMedQAResponse) GoString() string {
	return s.String()
}

func (s *RunMedQAResponse) SetHeaders(v map[string]*string) *RunMedQAResponse {
	s.Headers = v
	return s
}

func (s *RunMedQAResponse) SetBody(v *RunMedQAResponseBody) *RunMedQAResponse {
	s.Body = v
	return s
}

type ScreenChestCTRequest struct {
	DataFormat *string                        `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	Mask       *int64                         `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OrgId      *string                        `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string                        `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	URLList    []*ScreenChestCTRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s ScreenChestCTRequest) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTRequest) GoString() string {
	return s.String()
}

func (s *ScreenChestCTRequest) SetDataFormat(v string) *ScreenChestCTRequest {
	s.DataFormat = &v
	return s
}

func (s *ScreenChestCTRequest) SetMask(v int64) *ScreenChestCTRequest {
	s.Mask = &v
	return s
}

func (s *ScreenChestCTRequest) SetOrgId(v string) *ScreenChestCTRequest {
	s.OrgId = &v
	return s
}

func (s *ScreenChestCTRequest) SetOrgName(v string) *ScreenChestCTRequest {
	s.OrgName = &v
	return s
}

func (s *ScreenChestCTRequest) SetURLList(v []*ScreenChestCTRequestURLList) *ScreenChestCTRequest {
	s.URLList = v
	return s
}

type ScreenChestCTRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s ScreenChestCTRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTRequestURLList) GoString() string {
	return s.String()
}

func (s *ScreenChestCTRequestURLList) SetURL(v string) *ScreenChestCTRequestURLList {
	s.URL = &v
	return s
}

type ScreenChestCTResponseBody struct {
	Data      *ScreenChestCTResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScreenChestCTResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBody) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBody) SetData(v *ScreenChestCTResponseBodyData) *ScreenChestCTResponseBody {
	s.Data = v
	return s
}

func (s *ScreenChestCTResponseBody) SetRequestId(v string) *ScreenChestCTResponseBody {
	s.RequestId = &v
	return s
}

type ScreenChestCTResponseBodyData struct {
	AnalyzeChestVessel *ScreenChestCTResponseBodyDataAnalyzeChestVessel `json:"AnalyzeChestVessel,omitempty" xml:"AnalyzeChestVessel,omitempty" type:"Struct"`
	CACS               *ScreenChestCTResponseBodyDataCACS               `json:"CACS,omitempty" xml:"CACS,omitempty" type:"Struct"`
	Covid              *ScreenChestCTResponseBodyDataCovid              `json:"Covid,omitempty" xml:"Covid,omitempty" type:"Struct"`
	DetectRibFracture  *ScreenChestCTResponseBodyDataDetectRibFracture  `json:"DetectRibFracture,omitempty" xml:"DetectRibFracture,omitempty" type:"Struct"`
	ErrorMessage       *string                                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LungNodule         *ScreenChestCTResponseBodyDataLungNodule         `json:"LungNodule,omitempty" xml:"LungNodule,omitempty" type:"Struct"`
}

func (s ScreenChestCTResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyData) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyData) SetAnalyzeChestVessel(v *ScreenChestCTResponseBodyDataAnalyzeChestVessel) *ScreenChestCTResponseBodyData {
	s.AnalyzeChestVessel = v
	return s
}

func (s *ScreenChestCTResponseBodyData) SetCACS(v *ScreenChestCTResponseBodyDataCACS) *ScreenChestCTResponseBodyData {
	s.CACS = v
	return s
}

func (s *ScreenChestCTResponseBodyData) SetCovid(v *ScreenChestCTResponseBodyDataCovid) *ScreenChestCTResponseBodyData {
	s.Covid = v
	return s
}

func (s *ScreenChestCTResponseBodyData) SetDetectRibFracture(v *ScreenChestCTResponseBodyDataDetectRibFracture) *ScreenChestCTResponseBodyData {
	s.DetectRibFracture = v
	return s
}

func (s *ScreenChestCTResponseBodyData) SetErrorMessage(v string) *ScreenChestCTResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *ScreenChestCTResponseBodyData) SetLungNodule(v *ScreenChestCTResponseBodyDataLungNodule) *ScreenChestCTResponseBodyData {
	s.LungNodule = v
	return s
}

type ScreenChestCTResponseBodyDataAnalyzeChestVessel struct {
	AortaInfo     *ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo     `json:"AortaInfo,omitempty" xml:"AortaInfo,omitempty" type:"Struct"`
	PulmonaryInfo *ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo `json:"PulmonaryInfo,omitempty" xml:"PulmonaryInfo,omitempty" type:"Struct"`
	ResultURL     *string                                                       `json:"ResultURL,omitempty" xml:"ResultURL,omitempty"`
}

func (s ScreenChestCTResponseBodyDataAnalyzeChestVessel) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataAnalyzeChestVessel) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataAnalyzeChestVessel) SetAortaInfo(v *ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo) *ScreenChestCTResponseBodyDataAnalyzeChestVessel {
	s.AortaInfo = v
	return s
}

func (s *ScreenChestCTResponseBodyDataAnalyzeChestVessel) SetPulmonaryInfo(v *ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo) *ScreenChestCTResponseBodyDataAnalyzeChestVessel {
	s.PulmonaryInfo = v
	return s
}

func (s *ScreenChestCTResponseBodyDataAnalyzeChestVessel) SetResultURL(v string) *ScreenChestCTResponseBodyDataAnalyzeChestVessel {
	s.ResultURL = &v
	return s
}

type ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo struct {
	Area         []*float32   `json:"Area,omitempty" xml:"Area,omitempty" type:"Repeated"`
	Coordinates  [][]*float32 `json:"Coordinates,omitempty" xml:"Coordinates,omitempty" type:"Repeated"`
	LabelValue   *int64       `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
	MaxArea      *float32     `json:"MaxArea,omitempty" xml:"MaxArea,omitempty"`
	MaxAreaIndex *int64       `json:"MaxAreaIndex,omitempty" xml:"MaxAreaIndex,omitempty"`
	MaxDiameter  *float32     `json:"MaxDiameter,omitempty" xml:"MaxDiameter,omitempty"`
}

func (s ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo) SetArea(v []*float32) *ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo {
	s.Area = v
	return s
}

func (s *ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo) SetCoordinates(v [][]*float32) *ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo {
	s.Coordinates = v
	return s
}

func (s *ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo) SetLabelValue(v int64) *ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo {
	s.LabelValue = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo) SetMaxArea(v float32) *ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo {
	s.MaxArea = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo) SetMaxAreaIndex(v int64) *ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo {
	s.MaxAreaIndex = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo) SetMaxDiameter(v float32) *ScreenChestCTResponseBodyDataAnalyzeChestVesselAortaInfo {
	s.MaxDiameter = &v
	return s
}

type ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo struct {
	Area             []*float32   `json:"Area,omitempty" xml:"Area,omitempty" type:"Repeated"`
	Coordinates      [][]*float32 `json:"Coordinates,omitempty" xml:"Coordinates,omitempty" type:"Repeated"`
	LabelValue       *int64       `json:"LabelValue,omitempty" xml:"LabelValue,omitempty"`
	MaxArea          *float32     `json:"MaxArea,omitempty" xml:"MaxArea,omitempty"`
	MaxAreaIndex     *int64       `json:"MaxAreaIndex,omitempty" xml:"MaxAreaIndex,omitempty"`
	MaxDiameter      *float32     `json:"MaxDiameter,omitempty" xml:"MaxDiameter,omitempty"`
	NearestAortaArea *float32     `json:"NearestAortaArea,omitempty" xml:"NearestAortaArea,omitempty"`
}

func (s ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo) SetArea(v []*float32) *ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo {
	s.Area = v
	return s
}

func (s *ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo) SetCoordinates(v [][]*float32) *ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo {
	s.Coordinates = v
	return s
}

func (s *ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo) SetLabelValue(v int64) *ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo {
	s.LabelValue = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo) SetMaxArea(v float32) *ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo {
	s.MaxArea = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo) SetMaxAreaIndex(v int64) *ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo {
	s.MaxAreaIndex = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo) SetMaxDiameter(v float32) *ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo {
	s.MaxDiameter = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo) SetNearestAortaArea(v float32) *ScreenChestCTResponseBodyDataAnalyzeChestVesselPulmonaryInfo {
	s.NearestAortaArea = &v
	return s
}

type ScreenChestCTResponseBodyDataCACS struct {
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s ScreenChestCTResponseBodyDataCACS) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataCACS) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataCACS) SetResultUrl(v string) *ScreenChestCTResponseBodyDataCACS {
	s.ResultUrl = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataCACS) SetScore(v string) *ScreenChestCTResponseBodyDataCACS {
	s.Score = &v
	return s
}

type ScreenChestCTResponseBodyDataCovid struct {
	LesionRatio       *string `json:"LesionRatio,omitempty" xml:"LesionRatio,omitempty"`
	Mask              *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	NewProbability    *string `json:"NewProbability,omitempty" xml:"NewProbability,omitempty"`
	NormalProbability *string `json:"NormalProbability,omitempty" xml:"NormalProbability,omitempty"`
	OtherProbability  *string `json:"OtherProbability,omitempty" xml:"OtherProbability,omitempty"`
}

func (s ScreenChestCTResponseBodyDataCovid) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataCovid) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataCovid) SetLesionRatio(v string) *ScreenChestCTResponseBodyDataCovid {
	s.LesionRatio = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataCovid) SetMask(v string) *ScreenChestCTResponseBodyDataCovid {
	s.Mask = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataCovid) SetNewProbability(v string) *ScreenChestCTResponseBodyDataCovid {
	s.NewProbability = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataCovid) SetNormalProbability(v string) *ScreenChestCTResponseBodyDataCovid {
	s.NormalProbability = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataCovid) SetOtherProbability(v string) *ScreenChestCTResponseBodyDataCovid {
	s.OtherProbability = &v
	return s
}

type ScreenChestCTResponseBodyDataDetectRibFracture struct {
	Detections []*ScreenChestCTResponseBodyDataDetectRibFractureDetections `json:"Detections,omitempty" xml:"Detections,omitempty" type:"Repeated"`
	Origin     []*float32                                                  `json:"Origin,omitempty" xml:"Origin,omitempty" type:"Repeated"`
	ResultURL  *string                                                     `json:"ResultURL,omitempty" xml:"ResultURL,omitempty"`
	Spacing    []*float32                                                  `json:"Spacing,omitempty" xml:"Spacing,omitempty" type:"Repeated"`
}

func (s ScreenChestCTResponseBodyDataDetectRibFracture) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataDetectRibFracture) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataDetectRibFracture) SetDetections(v []*ScreenChestCTResponseBodyDataDetectRibFractureDetections) *ScreenChestCTResponseBodyDataDetectRibFracture {
	s.Detections = v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectRibFracture) SetOrigin(v []*float32) *ScreenChestCTResponseBodyDataDetectRibFracture {
	s.Origin = v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectRibFracture) SetResultURL(v string) *ScreenChestCTResponseBodyDataDetectRibFracture {
	s.ResultURL = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectRibFracture) SetSpacing(v []*float32) *ScreenChestCTResponseBodyDataDetectRibFracture {
	s.Spacing = v
	return s
}

type ScreenChestCTResponseBodyDataDetectRibFractureDetections struct {
	CoordinateImage    []*int64 `json:"CoordinateImage,omitempty" xml:"CoordinateImage,omitempty" type:"Repeated"`
	Coordinates        []*int64 `json:"Coordinates,omitempty" xml:"Coordinates,omitempty" type:"Repeated"`
	FractureCategory   *int64   `json:"FractureCategory,omitempty" xml:"FractureCategory,omitempty"`
	FractureConfidence *float32 `json:"FractureConfidence,omitempty" xml:"FractureConfidence,omitempty"`
	FractureId         *int64   `json:"FractureId,omitempty" xml:"FractureId,omitempty"`
	FractureLocation   *string  `json:"FractureLocation,omitempty" xml:"FractureLocation,omitempty"`
	FractureSegment    *int64   `json:"FractureSegment,omitempty" xml:"FractureSegment,omitempty"`
}

func (s ScreenChestCTResponseBodyDataDetectRibFractureDetections) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataDetectRibFractureDetections) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataDetectRibFractureDetections) SetCoordinateImage(v []*int64) *ScreenChestCTResponseBodyDataDetectRibFractureDetections {
	s.CoordinateImage = v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectRibFractureDetections) SetCoordinates(v []*int64) *ScreenChestCTResponseBodyDataDetectRibFractureDetections {
	s.Coordinates = v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectRibFractureDetections) SetFractureCategory(v int64) *ScreenChestCTResponseBodyDataDetectRibFractureDetections {
	s.FractureCategory = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectRibFractureDetections) SetFractureConfidence(v float32) *ScreenChestCTResponseBodyDataDetectRibFractureDetections {
	s.FractureConfidence = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectRibFractureDetections) SetFractureId(v int64) *ScreenChestCTResponseBodyDataDetectRibFractureDetections {
	s.FractureId = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectRibFractureDetections) SetFractureLocation(v string) *ScreenChestCTResponseBodyDataDetectRibFractureDetections {
	s.FractureLocation = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectRibFractureDetections) SetFractureSegment(v int64) *ScreenChestCTResponseBodyDataDetectRibFractureDetections {
	s.FractureSegment = &v
	return s
}

type ScreenChestCTResponseBodyDataLungNodule struct {
	Series []*ScreenChestCTResponseBodyDataLungNoduleSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
}

func (s ScreenChestCTResponseBodyDataLungNodule) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataLungNodule) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataLungNodule) SetSeries(v []*ScreenChestCTResponseBodyDataLungNoduleSeries) *ScreenChestCTResponseBodyDataLungNodule {
	s.Series = v
	return s
}

type ScreenChestCTResponseBodyDataLungNoduleSeries struct {
	Elements          []*ScreenChestCTResponseBodyDataLungNoduleSeriesElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Origin            []*float32                                               `json:"Origin,omitempty" xml:"Origin,omitempty" type:"Repeated"`
	Report            *string                                                  `json:"Report,omitempty" xml:"Report,omitempty"`
	SeriesInstanceUid *string                                                  `json:"SeriesInstanceUid,omitempty" xml:"SeriesInstanceUid,omitempty"`
	Spacing           []*float32                                               `json:"Spacing,omitempty" xml:"Spacing,omitempty" type:"Repeated"`
}

func (s ScreenChestCTResponseBodyDataLungNoduleSeries) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataLungNoduleSeries) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeries) SetElements(v []*ScreenChestCTResponseBodyDataLungNoduleSeriesElements) *ScreenChestCTResponseBodyDataLungNoduleSeries {
	s.Elements = v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeries) SetOrigin(v []*float32) *ScreenChestCTResponseBodyDataLungNoduleSeries {
	s.Origin = v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeries) SetReport(v string) *ScreenChestCTResponseBodyDataLungNoduleSeries {
	s.Report = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeries) SetSeriesInstanceUid(v string) *ScreenChestCTResponseBodyDataLungNoduleSeries {
	s.SeriesInstanceUid = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeries) SetSpacing(v []*float32) *ScreenChestCTResponseBodyDataLungNoduleSeries {
	s.Spacing = v
	return s
}

type ScreenChestCTResponseBodyDataLungNoduleSeriesElements struct {
	Category       *string  `json:"Category,omitempty" xml:"Category,omitempty"`
	Confidence     *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Diameter       *float32 `json:"Diameter,omitempty" xml:"Diameter,omitempty"`
	ImageX         *float32 `json:"ImageX,omitempty" xml:"ImageX,omitempty"`
	ImageY         *float32 `json:"ImageY,omitempty" xml:"ImageY,omitempty"`
	ImageZ         *float32 `json:"ImageZ,omitempty" xml:"ImageZ,omitempty"`
	Lobe           *string  `json:"Lobe,omitempty" xml:"Lobe,omitempty"`
	Lung           *string  `json:"Lung,omitempty" xml:"Lung,omitempty"`
	MeanValue      *float32 `json:"MeanValue,omitempty" xml:"MeanValue,omitempty"`
	SOPInstanceUID *string  `json:"SOPInstanceUID,omitempty" xml:"SOPInstanceUID,omitempty"`
	Volume         *float32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
	X              *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y              *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Z              *float32 `json:"Z,omitempty" xml:"Z,omitempty"`
}

func (s ScreenChestCTResponseBodyDataLungNoduleSeriesElements) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataLungNoduleSeriesElements) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetCategory(v string) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.Category = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetConfidence(v float32) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.Confidence = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetDiameter(v float32) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.Diameter = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetImageX(v float32) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.ImageX = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetImageY(v float32) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.ImageY = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetImageZ(v float32) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.ImageZ = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetLobe(v string) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.Lobe = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetLung(v string) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.Lung = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetMeanValue(v float32) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.MeanValue = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetSOPInstanceUID(v string) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.SOPInstanceUID = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetVolume(v float32) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.Volume = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetX(v float32) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.X = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetY(v float32) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.Y = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetZ(v float32) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.Z = &v
	return s
}

type ScreenChestCTResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ScreenChestCTResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScreenChestCTResponse) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponse) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponse) SetHeaders(v map[string]*string) *ScreenChestCTResponse {
	s.Headers = v
	return s
}

func (s *ScreenChestCTResponse) SetBody(v *ScreenChestCTResponseBody) *ScreenChestCTResponse {
	s.Body = v
	return s
}

type TranslateMedRequest struct {
	FromLanguage *string `json:"FromLanguage,omitempty" xml:"FromLanguage,omitempty"`
	Text         *string `json:"Text,omitempty" xml:"Text,omitempty"`
	ToLanguage   *string `json:"ToLanguage,omitempty" xml:"ToLanguage,omitempty"`
}

func (s TranslateMedRequest) String() string {
	return tea.Prettify(s)
}

func (s TranslateMedRequest) GoString() string {
	return s.String()
}

func (s *TranslateMedRequest) SetFromLanguage(v string) *TranslateMedRequest {
	s.FromLanguage = &v
	return s
}

func (s *TranslateMedRequest) SetText(v string) *TranslateMedRequest {
	s.Text = &v
	return s
}

func (s *TranslateMedRequest) SetToLanguage(v string) *TranslateMedRequest {
	s.ToLanguage = &v
	return s
}

type TranslateMedResponseBody struct {
	Data      *TranslateMedResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateMedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TranslateMedResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateMedResponseBody) SetData(v *TranslateMedResponseBodyData) *TranslateMedResponseBody {
	s.Data = v
	return s
}

func (s *TranslateMedResponseBody) SetRequestId(v string) *TranslateMedResponseBody {
	s.RequestId = &v
	return s
}

type TranslateMedResponseBodyData struct {
	Text  *string `json:"Text,omitempty" xml:"Text,omitempty"`
	Words *int64  `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s TranslateMedResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TranslateMedResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateMedResponseBodyData) SetText(v string) *TranslateMedResponseBodyData {
	s.Text = &v
	return s
}

func (s *TranslateMedResponseBodyData) SetWords(v int64) *TranslateMedResponseBodyData {
	s.Words = &v
	return s
}

type TranslateMedResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TranslateMedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TranslateMedResponse) String() string {
	return tea.Prettify(s)
}

func (s TranslateMedResponse) GoString() string {
	return s.String()
}

func (s *TranslateMedResponse) SetHeaders(v map[string]*string) *TranslateMedResponse {
	s.Headers = v
	return s
}

func (s *TranslateMedResponse) SetBody(v *TranslateMedResponseBody) *TranslateMedResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("imageprocess"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AnalyzeChestVesselWithOptions(request *AnalyzeChestVesselRequest, runtime *util.RuntimeOptions) (_result *AnalyzeChestVesselResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataFormat)) {
		body["DataFormat"] = request.DataFormat
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		body["URLList"] = request.URLList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AnalyzeChestVessel"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AnalyzeChestVesselResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AnalyzeChestVessel(request *AnalyzeChestVesselRequest) (_result *AnalyzeChestVesselResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AnalyzeChestVesselResponse{}
	_body, _err := client.AnalyzeChestVesselWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CalcCACSWithOptions(request *CalcCACSRequest, runtime *util.RuntimeOptions) (_result *CalcCACSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataFormat)) {
		body["DataFormat"] = request.DataFormat
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		body["URLList"] = request.URLList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CalcCACS"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CalcCACSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CalcCACS(request *CalcCACSRequest) (_result *CalcCACSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CalcCACSResponse{}
	_body, _err := client.CalcCACSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClassifyFNFWithOptions(request *ClassifyFNFRequest, runtime *util.RuntimeOptions) (_result *ClassifyFNFResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataFormat)) {
		body["DataFormat"] = request.DataFormat
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.TracerId)) {
		body["TracerId"] = request.TracerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ClassifyFNF"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClassifyFNFResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClassifyFNF(request *ClassifyFNFRequest) (_result *ClassifyFNFResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClassifyFNFResponse{}
	_body, _err := client.ClassifyFNFWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClassifyFNFAdvance(request *ClassifyFNFAdvanceRequest, runtime *util.RuntimeOptions) (_result *ClassifyFNFResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageprocess"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	classifyFNFReq := &ClassifyFNFRequest{}
	openapiutil.Convert(request, classifyFNFReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.ObjectKey,
			Content:     request.ImageUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.AccessKeyId,
			Policy:              authResponse.EncodedPolicy,
			Signature:           authResponse.Signature,
			Key:                 authResponse.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		classifyFNFReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	classifyFNFResp, _err := client.ClassifyFNFWithOptions(classifyFNFReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = classifyFNFResp
	return _result, _err
}

func (client *Client) DetectCovid19CadWithOptions(request *DetectCovid19CadRequest, runtime *util.RuntimeOptions) (_result *DetectCovid19CadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataFormat)) {
		body["DataFormat"] = request.DataFormat
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		body["URLList"] = request.URLList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectCovid19Cad"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectCovid19CadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectCovid19Cad(request *DetectCovid19CadRequest) (_result *DetectCovid19CadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectCovid19CadResponse{}
	_body, _err := client.DetectCovid19CadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectHipKeypointXRayWithOptions(request *DetectHipKeypointXRayRequest, runtime *util.RuntimeOptions) (_result *DetectHipKeypointXRayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataFormat)) {
		body["DataFormat"] = request.DataFormat
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.TracerId)) {
		body["TracerId"] = request.TracerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectHipKeypointXRay"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectHipKeypointXRayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectHipKeypointXRay(request *DetectHipKeypointXRayRequest) (_result *DetectHipKeypointXRayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectHipKeypointXRayResponse{}
	_body, _err := client.DetectHipKeypointXRayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectHipKeypointXRayAdvance(request *DetectHipKeypointXRayAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectHipKeypointXRayResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageprocess"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectHipKeypointXRayReq := &DetectHipKeypointXRayRequest{}
	openapiutil.Convert(request, detectHipKeypointXRayReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.ObjectKey,
			Content:     request.ImageUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.AccessKeyId,
			Policy:              authResponse.EncodedPolicy,
			Signature:           authResponse.Signature,
			Key:                 authResponse.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectHipKeypointXRayReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	detectHipKeypointXRayResp, _err := client.DetectHipKeypointXRayWithOptions(detectHipKeypointXRayReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectHipKeypointXRayResp
	return _result, _err
}

func (client *Client) DetectKneeKeypointXRayWithOptions(request *DetectKneeKeypointXRayRequest, runtime *util.RuntimeOptions) (_result *DetectKneeKeypointXRayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataFormat)) {
		body["DataFormat"] = request.DataFormat
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.TracerId)) {
		body["TracerId"] = request.TracerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectKneeKeypointXRay"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectKneeKeypointXRayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectKneeKeypointXRay(request *DetectKneeKeypointXRayRequest) (_result *DetectKneeKeypointXRayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectKneeKeypointXRayResponse{}
	_body, _err := client.DetectKneeKeypointXRayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectKneeKeypointXRayAdvance(request *DetectKneeKeypointXRayAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectKneeKeypointXRayResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageprocess"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectKneeKeypointXRayReq := &DetectKneeKeypointXRayRequest{}
	openapiutil.Convert(request, detectKneeKeypointXRayReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.ObjectKey,
			Content:     request.ImageUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.AccessKeyId,
			Policy:              authResponse.EncodedPolicy,
			Signature:           authResponse.Signature,
			Key:                 authResponse.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectKneeKeypointXRayReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	detectKneeKeypointXRayResp, _err := client.DetectKneeKeypointXRayWithOptions(detectKneeKeypointXRayReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectKneeKeypointXRayResp
	return _result, _err
}

func (client *Client) DetectKneeXRayWithOptions(request *DetectKneeXRayRequest, runtime *util.RuntimeOptions) (_result *DetectKneeXRayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataFormat)) {
		body["DataFormat"] = request.DataFormat
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectKneeXRay"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectKneeXRayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectKneeXRay(request *DetectKneeXRayRequest) (_result *DetectKneeXRayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectKneeXRayResponse{}
	_body, _err := client.DetectKneeXRayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectKneeXRayAdvance(request *DetectKneeXRayAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectKneeXRayResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageprocess"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectKneeXRayReq := &DetectKneeXRayRequest{}
	openapiutil.Convert(request, detectKneeXRayReq)
	if !tea.BoolValue(util.IsUnset(request.UrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.ObjectKey,
			Content:     request.UrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.AccessKeyId,
			Policy:              authResponse.EncodedPolicy,
			Signature:           authResponse.Signature,
			Key:                 authResponse.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectKneeXRayReq.Url = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	detectKneeXRayResp, _err := client.DetectKneeXRayWithOptions(detectKneeXRayReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectKneeXRayResp
	return _result, _err
}

func (client *Client) DetectLungNoduleWithOptions(request *DetectLungNoduleRequest, runtime *util.RuntimeOptions) (_result *DetectLungNoduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataFormat)) {
		body["DataFormat"] = request.DataFormat
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.Threshold)) {
		body["Threshold"] = request.Threshold
	}

	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		body["URLList"] = request.URLList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectLungNodule"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectLungNoduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectLungNodule(request *DetectLungNoduleRequest) (_result *DetectLungNoduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectLungNoduleResponse{}
	_body, _err := client.DetectLungNoduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectRibFractureWithOptions(request *DetectRibFractureRequest, runtime *util.RuntimeOptions) (_result *DetectRibFractureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataFormat)) {
		body["DataFormat"] = request.DataFormat
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		body["URLList"] = request.URLList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectRibFracture"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectRibFractureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectRibFracture(request *DetectRibFractureRequest) (_result *DetectRibFractureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectRibFractureResponse{}
	_body, _err := client.DetectRibFractureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectSkinDiseaseWithOptions(request *DetectSkinDiseaseRequest, runtime *util.RuntimeOptions) (_result *DetectSkinDiseaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectSkinDisease"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectSkinDiseaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectSkinDisease(request *DetectSkinDiseaseRequest) (_result *DetectSkinDiseaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectSkinDiseaseResponse{}
	_body, _err := client.DetectSkinDiseaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectSkinDiseaseAdvance(request *DetectSkinDiseaseAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectSkinDiseaseResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imageprocess"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectSkinDiseaseReq := &DetectSkinDiseaseRequest{}
	openapiutil.Convert(request, detectSkinDiseaseReq)
	if !tea.BoolValue(util.IsUnset(request.UrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.ObjectKey,
			Content:     request.UrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.AccessKeyId,
			Policy:              authResponse.EncodedPolicy,
			Signature:           authResponse.Signature,
			Key:                 authResponse.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectSkinDiseaseReq.Url = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	detectSkinDiseaseResp, _err := client.DetectSkinDiseaseWithOptions(detectSkinDiseaseReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectSkinDiseaseResp
	return _result, _err
}

func (client *Client) DetectSpineMRIWithOptions(request *DetectSpineMRIRequest, runtime *util.RuntimeOptions) (_result *DetectSpineMRIResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataFormat)) {
		body["DataFormat"] = request.DataFormat
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		body["URLList"] = request.URLList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectSpineMRI"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectSpineMRIResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectSpineMRI(request *DetectSpineMRIRequest) (_result *DetectSpineMRIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectSpineMRIResponse{}
	_body, _err := client.DetectSpineMRIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsyncJobResultWithOptions(request *GetAsyncJobResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncJobResult"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncJobResult(request *GetAsyncJobResultRequest) (_result *GetAsyncJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.GetAsyncJobResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunCTRegistrationWithOptions(request *RunCTRegistrationRequest, runtime *util.RuntimeOptions) (_result *RunCTRegistrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataFormat)) {
		body["DataFormat"] = request.DataFormat
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.FloatingList)) {
		body["FloatingList"] = request.FloatingList
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.ReferenceList)) {
		body["ReferenceList"] = request.ReferenceList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunCTRegistration"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunCTRegistrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunCTRegistration(request *RunCTRegistrationRequest) (_result *RunCTRegistrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCTRegistrationResponse{}
	_body, _err := client.RunCTRegistrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunMedQAWithOptions(request *RunMedQARequest, runtime *util.RuntimeOptions) (_result *RunMedQAResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnswerImageDataList)) {
		body["AnswerImageDataList"] = request.AnswerImageDataList
	}

	if !tea.BoolValue(util.IsUnset(request.AnswerImageURLList)) {
		body["AnswerImageURLList"] = request.AnswerImageURLList
	}

	if !tea.BoolValue(util.IsUnset(request.AnswerTextList)) {
		body["AnswerTextList"] = request.AnswerTextList
	}

	if !tea.BoolValue(util.IsUnset(request.Department)) {
		body["Department"] = request.Department
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.QuestionType)) {
		body["QuestionType"] = request.QuestionType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunMedQA"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunMedQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunMedQA(request *RunMedQARequest) (_result *RunMedQAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunMedQAResponse{}
	_body, _err := client.RunMedQAWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScreenChestCTWithOptions(request *ScreenChestCTRequest, runtime *util.RuntimeOptions) (_result *ScreenChestCTResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataFormat)) {
		body["DataFormat"] = request.DataFormat
	}

	if !tea.BoolValue(util.IsUnset(request.Mask)) {
		body["Mask"] = request.Mask
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		body["URLList"] = request.URLList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ScreenChestCT"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ScreenChestCTResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScreenChestCT(request *ScreenChestCTRequest) (_result *ScreenChestCTResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScreenChestCTResponse{}
	_body, _err := client.ScreenChestCTWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TranslateMedWithOptions(request *TranslateMedRequest, runtime *util.RuntimeOptions) (_result *TranslateMedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FromLanguage)) {
		body["FromLanguage"] = request.FromLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.ToLanguage)) {
		body["ToLanguage"] = request.ToLanguage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TranslateMed"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TranslateMedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TranslateMed(request *TranslateMedRequest) (_result *TranslateMedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TranslateMedResponse{}
	_body, _err := client.TranslateMedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
