// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	number "github.com/alibabacloud-go/darabonba-number/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type AnalyzeChestVesselRequest struct {
	DataFormat     *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	OrgId          *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName        *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// 1
	URLList []*AnalyzeChestVesselRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
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

type AnalyzeChestVesselAdvanceRequest struct {
	DataFormat     *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	OrgId          *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName        *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// 1
	URLList []*AnalyzeChestVesselAdvanceRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s AnalyzeChestVesselAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeChestVesselAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AnalyzeChestVesselAdvanceRequest) SetDataFormat(v string) *AnalyzeChestVesselAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *AnalyzeChestVesselAdvanceRequest) SetDataSourceType(v string) *AnalyzeChestVesselAdvanceRequest {
	s.DataSourceType = &v
	return s
}

func (s *AnalyzeChestVesselAdvanceRequest) SetOrgId(v string) *AnalyzeChestVesselAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *AnalyzeChestVesselAdvanceRequest) SetOrgName(v string) *AnalyzeChestVesselAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *AnalyzeChestVesselAdvanceRequest) SetURLList(v []*AnalyzeChestVesselAdvanceRequestURLList) *AnalyzeChestVesselAdvanceRequest {
	s.URLList = v
	return s
}

type AnalyzeChestVesselAdvanceRequestURLList struct {
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s AnalyzeChestVesselAdvanceRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeChestVesselAdvanceRequestURLList) GoString() string {
	return s.String()
}

func (s *AnalyzeChestVesselAdvanceRequestURLList) SetURLObject(v io.Reader) *AnalyzeChestVesselAdvanceRequestURLList {
	s.URLObject = v
	return s
}

type AnalyzeChestVesselResponseBody struct {
	Data      *AnalyzeChestVesselResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *AnalyzeChestVesselResponseBody) SetMessage(v string) *AnalyzeChestVesselResponseBody {
	s.Message = &v
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
	// 1
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
	// 1
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AnalyzeChestVesselResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AnalyzeChestVesselResponse) SetStatusCode(v int32) *AnalyzeChestVesselResponse {
	s.StatusCode = &v
	return s
}

func (s *AnalyzeChestVesselResponse) SetBody(v *AnalyzeChestVesselResponseBody) *AnalyzeChestVesselResponse {
	s.Body = v
	return s
}

type CalcBMDRequest struct {
	DataFormat *string                  `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId      *string                  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string                  `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	SourceType *string                  `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	URLList    []*CalcBMDRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s CalcBMDRequest) String() string {
	return tea.Prettify(s)
}

func (s CalcBMDRequest) GoString() string {
	return s.String()
}

func (s *CalcBMDRequest) SetDataFormat(v string) *CalcBMDRequest {
	s.DataFormat = &v
	return s
}

func (s *CalcBMDRequest) SetOrgId(v string) *CalcBMDRequest {
	s.OrgId = &v
	return s
}

func (s *CalcBMDRequest) SetOrgName(v string) *CalcBMDRequest {
	s.OrgName = &v
	return s
}

func (s *CalcBMDRequest) SetSourceType(v string) *CalcBMDRequest {
	s.SourceType = &v
	return s
}

func (s *CalcBMDRequest) SetURLList(v []*CalcBMDRequestURLList) *CalcBMDRequest {
	s.URLList = v
	return s
}

type CalcBMDRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s CalcBMDRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s CalcBMDRequestURLList) GoString() string {
	return s.String()
}

func (s *CalcBMDRequestURLList) SetURL(v string) *CalcBMDRequestURLList {
	s.URL = &v
	return s
}

type CalcBMDAdvanceRequest struct {
	DataFormat *string                         `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId      *string                         `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string                         `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	SourceType *string                         `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	URLList    []*CalcBMDAdvanceRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s CalcBMDAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CalcBMDAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CalcBMDAdvanceRequest) SetDataFormat(v string) *CalcBMDAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *CalcBMDAdvanceRequest) SetOrgId(v string) *CalcBMDAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *CalcBMDAdvanceRequest) SetOrgName(v string) *CalcBMDAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *CalcBMDAdvanceRequest) SetSourceType(v string) *CalcBMDAdvanceRequest {
	s.SourceType = &v
	return s
}

func (s *CalcBMDAdvanceRequest) SetURLList(v []*CalcBMDAdvanceRequestURLList) *CalcBMDAdvanceRequest {
	s.URLList = v
	return s
}

type CalcBMDAdvanceRequestURLList struct {
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s CalcBMDAdvanceRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s CalcBMDAdvanceRequestURLList) GoString() string {
	return s.String()
}

func (s *CalcBMDAdvanceRequestURLList) SetURLObject(v io.Reader) *CalcBMDAdvanceRequestURLList {
	s.URLObject = v
	return s
}

type CalcBMDResponseBody struct {
	Data      *CalcBMDResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CalcBMDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CalcBMDResponseBody) GoString() string {
	return s.String()
}

func (s *CalcBMDResponseBody) SetData(v *CalcBMDResponseBodyData) *CalcBMDResponseBody {
	s.Data = v
	return s
}

func (s *CalcBMDResponseBody) SetMessage(v string) *CalcBMDResponseBody {
	s.Message = &v
	return s
}

func (s *CalcBMDResponseBody) SetRequestId(v string) *CalcBMDResponseBody {
	s.RequestId = &v
	return s
}

type CalcBMDResponseBodyData struct {
	Detections []*CalcBMDResponseBodyDataDetections `json:"Detections,omitempty" xml:"Detections,omitempty" type:"Repeated"`
	Origin     []*float32                           `json:"Origin,omitempty" xml:"Origin,omitempty" type:"Repeated"`
	ResultURL  *string                              `json:"ResultURL,omitempty" xml:"ResultURL,omitempty"`
	Spacing    []*float32                           `json:"Spacing,omitempty" xml:"Spacing,omitempty" type:"Repeated"`
}

func (s CalcBMDResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CalcBMDResponseBodyData) GoString() string {
	return s.String()
}

func (s *CalcBMDResponseBodyData) SetDetections(v []*CalcBMDResponseBodyDataDetections) *CalcBMDResponseBodyData {
	s.Detections = v
	return s
}

func (s *CalcBMDResponseBodyData) SetOrigin(v []*float32) *CalcBMDResponseBodyData {
	s.Origin = v
	return s
}

func (s *CalcBMDResponseBodyData) SetResultURL(v string) *CalcBMDResponseBodyData {
	s.ResultURL = &v
	return s
}

func (s *CalcBMDResponseBodyData) SetSpacing(v []*float32) *CalcBMDResponseBodyData {
	s.Spacing = v
	return s
}

type CalcBMDResponseBodyDataDetections struct {
	VertBMD      *float32 `json:"VertBMD,omitempty" xml:"VertBMD,omitempty"`
	VertCategory *float32 `json:"VertCategory,omitempty" xml:"VertCategory,omitempty"`
	VertId       *string  `json:"VertId,omitempty" xml:"VertId,omitempty"`
	VertTScore   *float32 `json:"VertTScore,omitempty" xml:"VertTScore,omitempty"`
	VertZScore   *float32 `json:"VertZScore,omitempty" xml:"VertZScore,omitempty"`
}

func (s CalcBMDResponseBodyDataDetections) String() string {
	return tea.Prettify(s)
}

func (s CalcBMDResponseBodyDataDetections) GoString() string {
	return s.String()
}

func (s *CalcBMDResponseBodyDataDetections) SetVertBMD(v float32) *CalcBMDResponseBodyDataDetections {
	s.VertBMD = &v
	return s
}

func (s *CalcBMDResponseBodyDataDetections) SetVertCategory(v float32) *CalcBMDResponseBodyDataDetections {
	s.VertCategory = &v
	return s
}

func (s *CalcBMDResponseBodyDataDetections) SetVertId(v string) *CalcBMDResponseBodyDataDetections {
	s.VertId = &v
	return s
}

func (s *CalcBMDResponseBodyDataDetections) SetVertTScore(v float32) *CalcBMDResponseBodyDataDetections {
	s.VertTScore = &v
	return s
}

func (s *CalcBMDResponseBodyDataDetections) SetVertZScore(v float32) *CalcBMDResponseBodyDataDetections {
	s.VertZScore = &v
	return s
}

type CalcBMDResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CalcBMDResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CalcBMDResponse) String() string {
	return tea.Prettify(s)
}

func (s CalcBMDResponse) GoString() string {
	return s.String()
}

func (s *CalcBMDResponse) SetHeaders(v map[string]*string) *CalcBMDResponse {
	s.Headers = v
	return s
}

func (s *CalcBMDResponse) SetStatusCode(v int32) *CalcBMDResponse {
	s.StatusCode = &v
	return s
}

func (s *CalcBMDResponse) SetBody(v *CalcBMDResponseBody) *CalcBMDResponse {
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

type CalcCACSAdvanceRequest struct {
	DataFormat     *string                          `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	DataSourceType *string                          `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	OrgId          *string                          `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName        *string                          `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	URLList        []*CalcCACSAdvanceRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s CalcCACSAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CalcCACSAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CalcCACSAdvanceRequest) SetDataFormat(v string) *CalcCACSAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *CalcCACSAdvanceRequest) SetDataSourceType(v string) *CalcCACSAdvanceRequest {
	s.DataSourceType = &v
	return s
}

func (s *CalcCACSAdvanceRequest) SetOrgId(v string) *CalcCACSAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *CalcCACSAdvanceRequest) SetOrgName(v string) *CalcCACSAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *CalcCACSAdvanceRequest) SetURLList(v []*CalcCACSAdvanceRequestURLList) *CalcCACSAdvanceRequest {
	s.URLList = v
	return s
}

type CalcCACSAdvanceRequestURLList struct {
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s CalcCACSAdvanceRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s CalcCACSAdvanceRequestURLList) GoString() string {
	return s.String()
}

func (s *CalcCACSAdvanceRequestURLList) SetURLObject(v io.Reader) *CalcCACSAdvanceRequestURLList {
	s.URLObject = v
	return s
}

type CalcCACSResponseBody struct {
	Data      *CalcCACSResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *CalcCACSResponseBody) SetMessage(v string) *CalcCACSResponseBody {
	s.Message = &v
	return s
}

func (s *CalcCACSResponseBody) SetRequestId(v string) *CalcCACSResponseBody {
	s.RequestId = &v
	return s
}

type CalcCACSResponseBodyData struct {
	Detections  []*CalcCACSResponseBodyDataDetections `json:"Detections,omitempty" xml:"Detections,omitempty" type:"Repeated"`
	ResultUrl   *string                               `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
	Score       *string                               `json:"Score,omitempty" xml:"Score,omitempty"`
	VolumeScore *string                               `json:"VolumeScore,omitempty" xml:"VolumeScore,omitempty"`
}

func (s CalcCACSResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CalcCACSResponseBodyData) GoString() string {
	return s.String()
}

func (s *CalcCACSResponseBodyData) SetDetections(v []*CalcCACSResponseBodyDataDetections) *CalcCACSResponseBodyData {
	s.Detections = v
	return s
}

func (s *CalcCACSResponseBodyData) SetResultUrl(v string) *CalcCACSResponseBodyData {
	s.ResultUrl = &v
	return s
}

func (s *CalcCACSResponseBodyData) SetScore(v string) *CalcCACSResponseBodyData {
	s.Score = &v
	return s
}

func (s *CalcCACSResponseBodyData) SetVolumeScore(v string) *CalcCACSResponseBodyData {
	s.VolumeScore = &v
	return s
}

type CalcCACSResponseBodyDataDetections struct {
	CalciumCenter []*int64 `json:"CalciumCenter,omitempty" xml:"CalciumCenter,omitempty" type:"Repeated"`
	CalciumId     *int64   `json:"CalciumId,omitempty" xml:"CalciumId,omitempty"`
	CalciumScore  *float32 `json:"CalciumScore,omitempty" xml:"CalciumScore,omitempty"`
	CalciumVolume *float32 `json:"CalciumVolume,omitempty" xml:"CalciumVolume,omitempty"`
}

func (s CalcCACSResponseBodyDataDetections) String() string {
	return tea.Prettify(s)
}

func (s CalcCACSResponseBodyDataDetections) GoString() string {
	return s.String()
}

func (s *CalcCACSResponseBodyDataDetections) SetCalciumCenter(v []*int64) *CalcCACSResponseBodyDataDetections {
	s.CalciumCenter = v
	return s
}

func (s *CalcCACSResponseBodyDataDetections) SetCalciumId(v int64) *CalcCACSResponseBodyDataDetections {
	s.CalciumId = &v
	return s
}

func (s *CalcCACSResponseBodyDataDetections) SetCalciumScore(v float32) *CalcCACSResponseBodyDataDetections {
	s.CalciumScore = &v
	return s
}

func (s *CalcCACSResponseBodyDataDetections) SetCalciumVolume(v float32) *CalcCACSResponseBodyDataDetections {
	s.CalciumVolume = &v
	return s
}

type CalcCACSResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CalcCACSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CalcCACSResponse) SetStatusCode(v int32) *CalcCACSResponse {
	s.StatusCode = &v
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
	DataFormat     *string   `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
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

func (s *ClassifyFNFAdvanceRequest) SetDataFormat(v string) *ClassifyFNFAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *ClassifyFNFAdvanceRequest) SetImageUrlObject(v io.Reader) *ClassifyFNFAdvanceRequest {
	s.ImageUrlObject = v
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ClassifyFNFResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ClassifyFNFResponse) SetStatusCode(v int32) *ClassifyFNFResponse {
	s.StatusCode = &v
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

type DetectCovid19CadAdvanceRequest struct {
	DataFormat *string                                  `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId      *string                                  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string                                  `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	URLList    []*DetectCovid19CadAdvanceRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s DetectCovid19CadAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectCovid19CadAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectCovid19CadAdvanceRequest) SetDataFormat(v string) *DetectCovid19CadAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectCovid19CadAdvanceRequest) SetOrgId(v string) *DetectCovid19CadAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *DetectCovid19CadAdvanceRequest) SetOrgName(v string) *DetectCovid19CadAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *DetectCovid19CadAdvanceRequest) SetURLList(v []*DetectCovid19CadAdvanceRequestURLList) *DetectCovid19CadAdvanceRequest {
	s.URLList = v
	return s
}

type DetectCovid19CadAdvanceRequestURLList struct {
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s DetectCovid19CadAdvanceRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectCovid19CadAdvanceRequestURLList) GoString() string {
	return s.String()
}

func (s *DetectCovid19CadAdvanceRequestURLList) SetURLObject(v io.Reader) *DetectCovid19CadAdvanceRequestURLList {
	s.URLObject = v
	return s
}

type DetectCovid19CadResponseBody struct {
	Data      *DetectCovid19CadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *DetectCovid19CadResponseBody) SetMessage(v string) *DetectCovid19CadResponseBody {
	s.Message = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectCovid19CadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectCovid19CadResponse) SetStatusCode(v int32) *DetectCovid19CadResponse {
	s.StatusCode = &v
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
	DataFormat     *string   `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
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

func (s *DetectHipKeypointXRayAdvanceRequest) SetDataFormat(v string) *DetectHipKeypointXRayAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectHipKeypointXRayAdvanceRequest) SetImageUrlObject(v io.Reader) *DetectHipKeypointXRayAdvanceRequest {
	s.ImageUrlObject = v
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
	// 1
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectHipKeypointXRayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectHipKeypointXRayResponse) SetStatusCode(v int32) *DetectHipKeypointXRayResponse {
	s.StatusCode = &v
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
	DataFormat     *string   `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
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

func (s *DetectKneeKeypointXRayAdvanceRequest) SetDataFormat(v string) *DetectKneeKeypointXRayAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectKneeKeypointXRayAdvanceRequest) SetImageUrlObject(v io.Reader) *DetectKneeKeypointXRayAdvanceRequest {
	s.ImageUrlObject = v
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
	// 1
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectKneeKeypointXRayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectKneeKeypointXRayResponse) SetStatusCode(v int32) *DetectKneeKeypointXRayResponse {
	s.StatusCode = &v
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
	DataFormat *string   `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId      *string   `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string   `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	UrlObject  io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DetectKneeXRayAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectKneeXRayAdvanceRequest) GoString() string {
	return s.String()
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

func (s *DetectKneeXRayAdvanceRequest) SetUrlObject(v io.Reader) *DetectKneeXRayAdvanceRequest {
	s.UrlObject = v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectKneeXRayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectKneeXRayResponse) SetStatusCode(v int32) *DetectKneeXRayResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectKneeXRayResponse) SetBody(v *DetectKneeXRayResponseBody) *DetectKneeXRayResponse {
	s.Body = v
	return s
}

type DetectLungNoduleRequest struct {
	DataFormat *string  `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId      *string  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string  `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	Threshold  *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// 1
	URLList []*DetectLungNoduleRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
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

type DetectLungNoduleAdvanceRequest struct {
	DataFormat *string  `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId      *string  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string  `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	Threshold  *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// 1
	URLList []*DetectLungNoduleAdvanceRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s DetectLungNoduleAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectLungNoduleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectLungNoduleAdvanceRequest) SetDataFormat(v string) *DetectLungNoduleAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectLungNoduleAdvanceRequest) SetOrgId(v string) *DetectLungNoduleAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *DetectLungNoduleAdvanceRequest) SetOrgName(v string) *DetectLungNoduleAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *DetectLungNoduleAdvanceRequest) SetThreshold(v float32) *DetectLungNoduleAdvanceRequest {
	s.Threshold = &v
	return s
}

func (s *DetectLungNoduleAdvanceRequest) SetURLList(v []*DetectLungNoduleAdvanceRequestURLList) *DetectLungNoduleAdvanceRequest {
	s.URLList = v
	return s
}

type DetectLungNoduleAdvanceRequestURLList struct {
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s DetectLungNoduleAdvanceRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectLungNoduleAdvanceRequestURLList) GoString() string {
	return s.String()
}

func (s *DetectLungNoduleAdvanceRequestURLList) SetURLObject(v io.Reader) *DetectLungNoduleAdvanceRequestURLList {
	s.URLObject = v
	return s
}

type DetectLungNoduleResponseBody struct {
	Data      *DetectLungNoduleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *DetectLungNoduleResponseBody) SetMessage(v string) *DetectLungNoduleResponseBody {
	s.Message = &v
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
	Elements []*DetectLungNoduleResponseBodyDataSeriesElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	// 1
	Origin            []*float32 `json:"Origin,omitempty" xml:"Origin,omitempty" type:"Repeated"`
	Report            *string    `json:"Report,omitempty" xml:"Report,omitempty"`
	SeriesInstanceUid *string    `json:"SeriesInstanceUid,omitempty" xml:"SeriesInstanceUid,omitempty"`
	// 1
	Spacing []*float32 `json:"Spacing,omitempty" xml:"Spacing,omitempty" type:"Repeated"`
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
	Category             *string    `json:"Category,omitempty" xml:"Category,omitempty"`
	Confidence           *float32   `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Diameter             *float32   `json:"Diameter,omitempty" xml:"Diameter,omitempty"`
	ImageX               *float32   `json:"ImageX,omitempty" xml:"ImageX,omitempty"`
	ImageY               *float32   `json:"ImageY,omitempty" xml:"ImageY,omitempty"`
	ImageZ               *float32   `json:"ImageZ,omitempty" xml:"ImageZ,omitempty"`
	Lobe                 *string    `json:"Lobe,omitempty" xml:"Lobe,omitempty"`
	Lung                 *string    `json:"Lung,omitempty" xml:"Lung,omitempty"`
	MajorAxis            []*float32 `json:"MajorAxis,omitempty" xml:"MajorAxis,omitempty" type:"Repeated"`
	MeanValue            *float32   `json:"MeanValue,omitempty" xml:"MeanValue,omitempty"`
	MinorAxis            []*float32 `json:"MinorAxis,omitempty" xml:"MinorAxis,omitempty" type:"Repeated"`
	RecistSOPInstanceUID *string    `json:"RecistSOPInstanceUID,omitempty" xml:"RecistSOPInstanceUID,omitempty"`
	Risk                 *float32   `json:"Risk,omitempty" xml:"Risk,omitempty"`
	SOPInstanceUID       *string    `json:"SOPInstanceUID,omitempty" xml:"SOPInstanceUID,omitempty"`
	Volume               *float32   `json:"Volume,omitempty" xml:"Volume,omitempty"`
	X                    *float32   `json:"X,omitempty" xml:"X,omitempty"`
	Y                    *float32   `json:"Y,omitempty" xml:"Y,omitempty"`
	Z                    *float32   `json:"Z,omitempty" xml:"Z,omitempty"`
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

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetMajorAxis(v []*float32) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.MajorAxis = v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetMeanValue(v float32) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.MeanValue = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetMinorAxis(v []*float32) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.MinorAxis = v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetRecistSOPInstanceUID(v string) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.RecistSOPInstanceUID = &v
	return s
}

func (s *DetectLungNoduleResponseBodyDataSeriesElements) SetRisk(v float32) *DetectLungNoduleResponseBodyDataSeriesElements {
	s.Risk = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectLungNoduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectLungNoduleResponse) SetStatusCode(v int32) *DetectLungNoduleResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectLungNoduleResponse) SetBody(v *DetectLungNoduleResponseBody) *DetectLungNoduleResponse {
	s.Body = v
	return s
}

type DetectLymphRequest struct {
	DataSourceType *string                      `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	URLList        []*DetectLymphRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s DetectLymphRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectLymphRequest) GoString() string {
	return s.String()
}

func (s *DetectLymphRequest) SetDataSourceType(v string) *DetectLymphRequest {
	s.DataSourceType = &v
	return s
}

func (s *DetectLymphRequest) SetURLList(v []*DetectLymphRequestURLList) *DetectLymphRequest {
	s.URLList = v
	return s
}

type DetectLymphRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s DetectLymphRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectLymphRequestURLList) GoString() string {
	return s.String()
}

func (s *DetectLymphRequestURLList) SetURL(v string) *DetectLymphRequestURLList {
	s.URL = &v
	return s
}

type DetectLymphAdvanceRequest struct {
	DataSourceType *string                             `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	URLList        []*DetectLymphAdvanceRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s DetectLymphAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectLymphAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectLymphAdvanceRequest) SetDataSourceType(v string) *DetectLymphAdvanceRequest {
	s.DataSourceType = &v
	return s
}

func (s *DetectLymphAdvanceRequest) SetURLList(v []*DetectLymphAdvanceRequestURLList) *DetectLymphAdvanceRequest {
	s.URLList = v
	return s
}

type DetectLymphAdvanceRequestURLList struct {
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s DetectLymphAdvanceRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectLymphAdvanceRequestURLList) GoString() string {
	return s.String()
}

func (s *DetectLymphAdvanceRequestURLList) SetURLObject(v io.Reader) *DetectLymphAdvanceRequestURLList {
	s.URLObject = v
	return s
}

type DetectLymphResponseBody struct {
	Data      *DetectLymphResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectLymphResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectLymphResponseBody) GoString() string {
	return s.String()
}

func (s *DetectLymphResponseBody) SetData(v *DetectLymphResponseBodyData) *DetectLymphResponseBody {
	s.Data = v
	return s
}

func (s *DetectLymphResponseBody) SetMessage(v string) *DetectLymphResponseBody {
	s.Message = &v
	return s
}

func (s *DetectLymphResponseBody) SetRequestId(v string) *DetectLymphResponseBody {
	s.RequestId = &v
	return s
}

type DetectLymphResponseBodyData struct {
	Lesions []*DetectLymphResponseBodyDataLesions `json:"Lesions,omitempty" xml:"Lesions,omitempty" type:"Repeated"`
}

func (s DetectLymphResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectLymphResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectLymphResponseBodyData) SetLesions(v []*DetectLymphResponseBodyDataLesions) *DetectLymphResponseBodyData {
	s.Lesions = v
	return s
}

type DetectLymphResponseBodyDataLesions struct {
	Boxes      []*float32   `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Diametermm []*float32   `json:"Diametermm,omitempty" xml:"Diametermm,omitempty" type:"Repeated"`
	KeySlice   *int64       `json:"KeySlice,omitempty" xml:"KeySlice,omitempty"`
	Recist     [][]*float32 `json:"Recist,omitempty" xml:"Recist,omitempty" type:"Repeated"`
	Score      *float32     `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s DetectLymphResponseBodyDataLesions) String() string {
	return tea.Prettify(s)
}

func (s DetectLymphResponseBodyDataLesions) GoString() string {
	return s.String()
}

func (s *DetectLymphResponseBodyDataLesions) SetBoxes(v []*float32) *DetectLymphResponseBodyDataLesions {
	s.Boxes = v
	return s
}

func (s *DetectLymphResponseBodyDataLesions) SetDiametermm(v []*float32) *DetectLymphResponseBodyDataLesions {
	s.Diametermm = v
	return s
}

func (s *DetectLymphResponseBodyDataLesions) SetKeySlice(v int64) *DetectLymphResponseBodyDataLesions {
	s.KeySlice = &v
	return s
}

func (s *DetectLymphResponseBodyDataLesions) SetRecist(v [][]*float32) *DetectLymphResponseBodyDataLesions {
	s.Recist = v
	return s
}

func (s *DetectLymphResponseBodyDataLesions) SetScore(v float32) *DetectLymphResponseBodyDataLesions {
	s.Score = &v
	return s
}

type DetectLymphResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectLymphResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectLymphResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectLymphResponse) GoString() string {
	return s.String()
}

func (s *DetectLymphResponse) SetHeaders(v map[string]*string) *DetectLymphResponse {
	s.Headers = v
	return s
}

func (s *DetectLymphResponse) SetStatusCode(v int32) *DetectLymphResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectLymphResponse) SetBody(v *DetectLymphResponseBody) *DetectLymphResponse {
	s.Body = v
	return s
}

type DetectPancRequest struct {
	DataSourceType *string                     `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	URLList        []*DetectPancRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s DetectPancRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectPancRequest) GoString() string {
	return s.String()
}

func (s *DetectPancRequest) SetDataSourceType(v string) *DetectPancRequest {
	s.DataSourceType = &v
	return s
}

func (s *DetectPancRequest) SetURLList(v []*DetectPancRequestURLList) *DetectPancRequest {
	s.URLList = v
	return s
}

type DetectPancRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s DetectPancRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectPancRequestURLList) GoString() string {
	return s.String()
}

func (s *DetectPancRequestURLList) SetURL(v string) *DetectPancRequestURLList {
	s.URL = &v
	return s
}

type DetectPancAdvanceRequest struct {
	DataSourceType *string                            `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	URLList        []*DetectPancAdvanceRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s DetectPancAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectPancAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectPancAdvanceRequest) SetDataSourceType(v string) *DetectPancAdvanceRequest {
	s.DataSourceType = &v
	return s
}

func (s *DetectPancAdvanceRequest) SetURLList(v []*DetectPancAdvanceRequestURLList) *DetectPancAdvanceRequest {
	s.URLList = v
	return s
}

type DetectPancAdvanceRequestURLList struct {
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s DetectPancAdvanceRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectPancAdvanceRequestURLList) GoString() string {
	return s.String()
}

func (s *DetectPancAdvanceRequestURLList) SetURLObject(v io.Reader) *DetectPancAdvanceRequestURLList {
	s.URLObject = v
	return s
}

type DetectPancResponseBody struct {
	Data      *DetectPancResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectPancResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectPancResponseBody) GoString() string {
	return s.String()
}

func (s *DetectPancResponseBody) SetData(v *DetectPancResponseBodyData) *DetectPancResponseBody {
	s.Data = v
	return s
}

func (s *DetectPancResponseBody) SetMessage(v string) *DetectPancResponseBody {
	s.Message = &v
	return s
}

func (s *DetectPancResponseBody) SetRequestId(v string) *DetectPancResponseBody {
	s.RequestId = &v
	return s
}

type DetectPancResponseBodyData struct {
	Lesion *DetectPancResponseBodyDataLesion `json:"Lesion,omitempty" xml:"Lesion,omitempty" type:"Struct"`
}

func (s DetectPancResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectPancResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectPancResponseBodyData) SetLesion(v *DetectPancResponseBodyDataLesion) *DetectPancResponseBodyData {
	s.Lesion = v
	return s
}

type DetectPancResponseBodyDataLesion struct {
	Mask          *string   `json:"Mask,omitempty" xml:"Mask,omitempty"`
	NonPdacVol    *string   `json:"NonPdacVol,omitempty" xml:"NonPdacVol,omitempty"`
	PancVol       *string   `json:"PancVol,omitempty" xml:"PancVol,omitempty"`
	PdacVol       *string   `json:"PdacVol,omitempty" xml:"PdacVol,omitempty"`
	Possibilities []*string `json:"Possibilities,omitempty" xml:"Possibilities,omitempty" type:"Repeated"`
}

func (s DetectPancResponseBodyDataLesion) String() string {
	return tea.Prettify(s)
}

func (s DetectPancResponseBodyDataLesion) GoString() string {
	return s.String()
}

func (s *DetectPancResponseBodyDataLesion) SetMask(v string) *DetectPancResponseBodyDataLesion {
	s.Mask = &v
	return s
}

func (s *DetectPancResponseBodyDataLesion) SetNonPdacVol(v string) *DetectPancResponseBodyDataLesion {
	s.NonPdacVol = &v
	return s
}

func (s *DetectPancResponseBodyDataLesion) SetPancVol(v string) *DetectPancResponseBodyDataLesion {
	s.PancVol = &v
	return s
}

func (s *DetectPancResponseBodyDataLesion) SetPdacVol(v string) *DetectPancResponseBodyDataLesion {
	s.PdacVol = &v
	return s
}

func (s *DetectPancResponseBodyDataLesion) SetPossibilities(v []*string) *DetectPancResponseBodyDataLesion {
	s.Possibilities = v
	return s
}

type DetectPancResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectPancResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectPancResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectPancResponse) GoString() string {
	return s.String()
}

func (s *DetectPancResponse) SetHeaders(v map[string]*string) *DetectPancResponse {
	s.Headers = v
	return s
}

func (s *DetectPancResponse) SetStatusCode(v int32) *DetectPancResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectPancResponse) SetBody(v *DetectPancResponseBody) *DetectPancResponse {
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

type DetectRibFractureAdvanceRequest struct {
	DataFormat *string                                   `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId      *string                                   `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string                                   `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	SourceType *string                                   `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	URLList    []*DetectRibFractureAdvanceRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s DetectRibFractureAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectRibFractureAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectRibFractureAdvanceRequest) SetDataFormat(v string) *DetectRibFractureAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectRibFractureAdvanceRequest) SetOrgId(v string) *DetectRibFractureAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *DetectRibFractureAdvanceRequest) SetOrgName(v string) *DetectRibFractureAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *DetectRibFractureAdvanceRequest) SetSourceType(v string) *DetectRibFractureAdvanceRequest {
	s.SourceType = &v
	return s
}

func (s *DetectRibFractureAdvanceRequest) SetURLList(v []*DetectRibFractureAdvanceRequestURLList) *DetectRibFractureAdvanceRequest {
	s.URLList = v
	return s
}

type DetectRibFractureAdvanceRequestURLList struct {
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s DetectRibFractureAdvanceRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectRibFractureAdvanceRequestURLList) GoString() string {
	return s.String()
}

func (s *DetectRibFractureAdvanceRequestURLList) SetURLObject(v io.Reader) *DetectRibFractureAdvanceRequestURLList {
	s.URLObject = v
	return s
}

type DetectRibFractureResponseBody struct {
	Data      *DetectRibFractureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *DetectRibFractureResponseBody) SetMessage(v string) *DetectRibFractureResponseBody {
	s.Message = &v
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
	FracSOPInstanceUID *string  `json:"FracSOPInstanceUID,omitempty" xml:"FracSOPInstanceUID,omitempty"`
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

func (s *DetectRibFractureResponseBodyDataDetections) SetFracSOPInstanceUID(v string) *DetectRibFractureResponseBodyDataDetections {
	s.FracSOPInstanceUID = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectRibFractureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectRibFractureResponse) SetStatusCode(v int32) *DetectRibFractureResponse {
	s.StatusCode = &v
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
	OrgId     *string   `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName   *string   `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DetectSkinDiseaseAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectSkinDiseaseAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectSkinDiseaseAdvanceRequest) SetOrgId(v string) *DetectSkinDiseaseAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *DetectSkinDiseaseAdvanceRequest) SetOrgName(v string) *DetectSkinDiseaseAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *DetectSkinDiseaseAdvanceRequest) SetUrlObject(v io.Reader) *DetectSkinDiseaseAdvanceRequest {
	s.UrlObject = v
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
	BodyPart       *string                `json:"BodyPart,omitempty" xml:"BodyPart,omitempty"`
	ImageQuality   *float32               `json:"ImageQuality,omitempty" xml:"ImageQuality,omitempty"`
	ImageType      *string                `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Results        map[string]interface{} `json:"Results,omitempty" xml:"Results,omitempty"`
	ResultsEnglish map[string]interface{} `json:"ResultsEnglish,omitempty" xml:"ResultsEnglish,omitempty"`
}

func (s DetectSkinDiseaseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectSkinDiseaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectSkinDiseaseResponseBodyData) SetBodyPart(v string) *DetectSkinDiseaseResponseBodyData {
	s.BodyPart = &v
	return s
}

func (s *DetectSkinDiseaseResponseBodyData) SetImageQuality(v float32) *DetectSkinDiseaseResponseBodyData {
	s.ImageQuality = &v
	return s
}

func (s *DetectSkinDiseaseResponseBodyData) SetImageType(v string) *DetectSkinDiseaseResponseBodyData {
	s.ImageType = &v
	return s
}

func (s *DetectSkinDiseaseResponseBodyData) SetResults(v map[string]interface{}) *DetectSkinDiseaseResponseBodyData {
	s.Results = v
	return s
}

func (s *DetectSkinDiseaseResponseBodyData) SetResultsEnglish(v map[string]interface{}) *DetectSkinDiseaseResponseBodyData {
	s.ResultsEnglish = v
	return s
}

type DetectSkinDiseaseResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectSkinDiseaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectSkinDiseaseResponse) SetStatusCode(v int32) *DetectSkinDiseaseResponse {
	s.StatusCode = &v
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

type DetectSpineMRIAdvanceRequest struct {
	DataFormat *string                                `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId      *string                                `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string                                `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	URLList    []*DetectSpineMRIAdvanceRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s DetectSpineMRIAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectSpineMRIAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectSpineMRIAdvanceRequest) SetDataFormat(v string) *DetectSpineMRIAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *DetectSpineMRIAdvanceRequest) SetOrgId(v string) *DetectSpineMRIAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *DetectSpineMRIAdvanceRequest) SetOrgName(v string) *DetectSpineMRIAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *DetectSpineMRIAdvanceRequest) SetURLList(v []*DetectSpineMRIAdvanceRequestURLList) *DetectSpineMRIAdvanceRequest {
	s.URLList = v
	return s
}

type DetectSpineMRIAdvanceRequestURLList struct {
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s DetectSpineMRIAdvanceRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s DetectSpineMRIAdvanceRequestURLList) GoString() string {
	return s.String()
}

func (s *DetectSpineMRIAdvanceRequestURLList) SetURLObject(v io.Reader) *DetectSpineMRIAdvanceRequestURLList {
	s.URLObject = v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectSpineMRIResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectSpineMRIResponse) SetStatusCode(v int32) *DetectSpineMRIResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectSpineMRIResponse) SetBody(v *DetectSpineMRIResponseBody) *DetectSpineMRIResponse {
	s.Body = v
	return s
}

type FeedbackSessionRequest struct {
	Feedback  *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s FeedbackSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s FeedbackSessionRequest) GoString() string {
	return s.String()
}

func (s *FeedbackSessionRequest) SetFeedback(v string) *FeedbackSessionRequest {
	s.Feedback = &v
	return s
}

func (s *FeedbackSessionRequest) SetSessionId(v string) *FeedbackSessionRequest {
	s.SessionId = &v
	return s
}

type FeedbackSessionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FeedbackSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FeedbackSessionResponseBody) GoString() string {
	return s.String()
}

func (s *FeedbackSessionResponseBody) SetRequestId(v string) *FeedbackSessionResponseBody {
	s.RequestId = &v
	return s
}

type FeedbackSessionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FeedbackSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FeedbackSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s FeedbackSessionResponse) GoString() string {
	return s.String()
}

func (s *FeedbackSessionResponse) SetHeaders(v map[string]*string) *FeedbackSessionResponse {
	s.Headers = v
	return s
}

func (s *FeedbackSessionResponse) SetStatusCode(v int32) *FeedbackSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *FeedbackSessionResponse) SetBody(v *FeedbackSessionResponseBody) *FeedbackSessionResponse {
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAsyncJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAsyncJobResultResponse) SetStatusCode(v int32) *GetAsyncJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncJobResultResponse) SetBody(v *GetAsyncJobResultResponseBody) *GetAsyncJobResultResponse {
	s.Body = v
	return s
}

type RunCTRegistrationRequest struct {
	// DICOM。
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

type RunCTRegistrationAdvanceRequest struct {
	// DICOM。
	DataFormat     *string                                         `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	DataSourceType *string                                         `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	FloatingList   []*RunCTRegistrationAdvanceRequestFloatingList  `json:"FloatingList,omitempty" xml:"FloatingList,omitempty" type:"Repeated"`
	OrgId          *string                                         `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName        *string                                         `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	ReferenceList  []*RunCTRegistrationAdvanceRequestReferenceList `json:"ReferenceList,omitempty" xml:"ReferenceList,omitempty" type:"Repeated"`
}

func (s RunCTRegistrationAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCTRegistrationAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RunCTRegistrationAdvanceRequest) SetDataFormat(v string) *RunCTRegistrationAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *RunCTRegistrationAdvanceRequest) SetDataSourceType(v string) *RunCTRegistrationAdvanceRequest {
	s.DataSourceType = &v
	return s
}

func (s *RunCTRegistrationAdvanceRequest) SetFloatingList(v []*RunCTRegistrationAdvanceRequestFloatingList) *RunCTRegistrationAdvanceRequest {
	s.FloatingList = v
	return s
}

func (s *RunCTRegistrationAdvanceRequest) SetOrgId(v string) *RunCTRegistrationAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *RunCTRegistrationAdvanceRequest) SetOrgName(v string) *RunCTRegistrationAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *RunCTRegistrationAdvanceRequest) SetReferenceList(v []*RunCTRegistrationAdvanceRequestReferenceList) *RunCTRegistrationAdvanceRequest {
	s.ReferenceList = v
	return s
}

type RunCTRegistrationAdvanceRequestFloatingList struct {
	FloatingURLObject io.Reader `json:"FloatingURL,omitempty" xml:"FloatingURL,omitempty"`
}

func (s RunCTRegistrationAdvanceRequestFloatingList) String() string {
	return tea.Prettify(s)
}

func (s RunCTRegistrationAdvanceRequestFloatingList) GoString() string {
	return s.String()
}

func (s *RunCTRegistrationAdvanceRequestFloatingList) SetFloatingURLObject(v io.Reader) *RunCTRegistrationAdvanceRequestFloatingList {
	s.FloatingURLObject = v
	return s
}

type RunCTRegistrationAdvanceRequestReferenceList struct {
	ReferenceURLObject io.Reader `json:"ReferenceURL,omitempty" xml:"ReferenceURL,omitempty"`
}

func (s RunCTRegistrationAdvanceRequestReferenceList) String() string {
	return tea.Prettify(s)
}

func (s RunCTRegistrationAdvanceRequestReferenceList) GoString() string {
	return s.String()
}

func (s *RunCTRegistrationAdvanceRequestReferenceList) SetReferenceURLObject(v io.Reader) *RunCTRegistrationAdvanceRequestReferenceList {
	s.ReferenceURLObject = v
	return s
}

type RunCTRegistrationResponseBody struct {
	Data      *RunCTRegistrationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *RunCTRegistrationResponseBody) SetMessage(v string) *RunCTRegistrationResponseBody {
	s.Message = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RunCTRegistrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RunCTRegistrationResponse) SetStatusCode(v int32) *RunCTRegistrationResponse {
	s.StatusCode = &v
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
	AnswerImageData *string `json:"AnswerImageData,omitempty" xml:"AnswerImageData,omitempty"`
}

func (s RunMedQARequestAnswerImageDataList) String() string {
	return tea.Prettify(s)
}

func (s RunMedQARequestAnswerImageDataList) GoString() string {
	return s.String()
}

func (s *RunMedQARequestAnswerImageDataList) SetAnswerImageData(v string) *RunMedQARequestAnswerImageDataList {
	s.AnswerImageData = &v
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

type RunMedQAAdvanceRequest struct {
	AnswerImageDataList []*RunMedQAAdvanceRequestAnswerImageDataList `json:"AnswerImageDataList,omitempty" xml:"AnswerImageDataList,omitempty" type:"Repeated"`
	AnswerImageURLList  []*RunMedQAAdvanceRequestAnswerImageURLList  `json:"AnswerImageURLList,omitempty" xml:"AnswerImageURLList,omitempty" type:"Repeated"`
	AnswerTextList      []*RunMedQAAdvanceRequestAnswerTextList      `json:"AnswerTextList,omitempty" xml:"AnswerTextList,omitempty" type:"Repeated"`
	Department          *string                                      `json:"Department,omitempty" xml:"Department,omitempty"`
	OrgId               *string                                      `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName             *string                                      `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	QuestionType        *string                                      `json:"QuestionType,omitempty" xml:"QuestionType,omitempty"`
	SessionId           *string                                      `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s RunMedQAAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RunMedQAAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RunMedQAAdvanceRequest) SetAnswerImageDataList(v []*RunMedQAAdvanceRequestAnswerImageDataList) *RunMedQAAdvanceRequest {
	s.AnswerImageDataList = v
	return s
}

func (s *RunMedQAAdvanceRequest) SetAnswerImageURLList(v []*RunMedQAAdvanceRequestAnswerImageURLList) *RunMedQAAdvanceRequest {
	s.AnswerImageURLList = v
	return s
}

func (s *RunMedQAAdvanceRequest) SetAnswerTextList(v []*RunMedQAAdvanceRequestAnswerTextList) *RunMedQAAdvanceRequest {
	s.AnswerTextList = v
	return s
}

func (s *RunMedQAAdvanceRequest) SetDepartment(v string) *RunMedQAAdvanceRequest {
	s.Department = &v
	return s
}

func (s *RunMedQAAdvanceRequest) SetOrgId(v string) *RunMedQAAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *RunMedQAAdvanceRequest) SetOrgName(v string) *RunMedQAAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *RunMedQAAdvanceRequest) SetQuestionType(v string) *RunMedQAAdvanceRequest {
	s.QuestionType = &v
	return s
}

func (s *RunMedQAAdvanceRequest) SetSessionId(v string) *RunMedQAAdvanceRequest {
	s.SessionId = &v
	return s
}

type RunMedQAAdvanceRequestAnswerImageDataList struct {
	AnswerImageData *string `json:"AnswerImageData,omitempty" xml:"AnswerImageData,omitempty"`
}

func (s RunMedQAAdvanceRequestAnswerImageDataList) String() string {
	return tea.Prettify(s)
}

func (s RunMedQAAdvanceRequestAnswerImageDataList) GoString() string {
	return s.String()
}

func (s *RunMedQAAdvanceRequestAnswerImageDataList) SetAnswerImageData(v string) *RunMedQAAdvanceRequestAnswerImageDataList {
	s.AnswerImageData = &v
	return s
}

type RunMedQAAdvanceRequestAnswerImageURLList struct {
	AnswerImageURLObject io.Reader `json:"AnswerImageURL,omitempty" xml:"AnswerImageURL,omitempty"`
}

func (s RunMedQAAdvanceRequestAnswerImageURLList) String() string {
	return tea.Prettify(s)
}

func (s RunMedQAAdvanceRequestAnswerImageURLList) GoString() string {
	return s.String()
}

func (s *RunMedQAAdvanceRequestAnswerImageURLList) SetAnswerImageURLObject(v io.Reader) *RunMedQAAdvanceRequestAnswerImageURLList {
	s.AnswerImageURLObject = v
	return s
}

type RunMedQAAdvanceRequestAnswerTextList struct {
	AnswerText *string `json:"AnswerText,omitempty" xml:"AnswerText,omitempty"`
}

func (s RunMedQAAdvanceRequestAnswerTextList) String() string {
	return tea.Prettify(s)
}

func (s RunMedQAAdvanceRequestAnswerTextList) GoString() string {
	return s.String()
}

func (s *RunMedQAAdvanceRequestAnswerTextList) SetAnswerText(v string) *RunMedQAAdvanceRequestAnswerTextList {
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RunMedQAResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RunMedQAResponse) SetStatusCode(v int32) *RunMedQAResponse {
	s.StatusCode = &v
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
	Verbose    *int64                         `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
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

func (s *ScreenChestCTRequest) SetVerbose(v int64) *ScreenChestCTRequest {
	s.Verbose = &v
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

type ScreenChestCTAdvanceRequest struct {
	DataFormat *string                               `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	Mask       *int64                                `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OrgId      *string                               `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string                               `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	URLList    []*ScreenChestCTAdvanceRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
	Verbose    *int64                                `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s ScreenChestCTAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ScreenChestCTAdvanceRequest) SetDataFormat(v string) *ScreenChestCTAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *ScreenChestCTAdvanceRequest) SetMask(v int64) *ScreenChestCTAdvanceRequest {
	s.Mask = &v
	return s
}

func (s *ScreenChestCTAdvanceRequest) SetOrgId(v string) *ScreenChestCTAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *ScreenChestCTAdvanceRequest) SetOrgName(v string) *ScreenChestCTAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *ScreenChestCTAdvanceRequest) SetURLList(v []*ScreenChestCTAdvanceRequestURLList) *ScreenChestCTAdvanceRequest {
	s.URLList = v
	return s
}

func (s *ScreenChestCTAdvanceRequest) SetVerbose(v int64) *ScreenChestCTAdvanceRequest {
	s.Verbose = &v
	return s
}

type ScreenChestCTAdvanceRequestURLList struct {
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s ScreenChestCTAdvanceRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTAdvanceRequestURLList) GoString() string {
	return s.String()
}

func (s *ScreenChestCTAdvanceRequestURLList) SetURLObject(v io.Reader) *ScreenChestCTAdvanceRequestURLList {
	s.URLObject = v
	return s
}

type ScreenChestCTResponseBody struct {
	Data      *ScreenChestCTResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *ScreenChestCTResponseBody) SetMessage(v string) *ScreenChestCTResponseBody {
	s.Message = &v
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
	DetectLymph        *ScreenChestCTResponseBodyDataDetectLymph        `json:"DetectLymph,omitempty" xml:"DetectLymph,omitempty" type:"Struct"`
	DetectPdac         *ScreenChestCTResponseBodyDataDetectPdac         `json:"DetectPdac,omitempty" xml:"DetectPdac,omitempty" type:"Struct"`
	DetectRibFracture  *ScreenChestCTResponseBodyDataDetectRibFracture  `json:"DetectRibFracture,omitempty" xml:"DetectRibFracture,omitempty" type:"Struct"`
	ErrorMessage       *string                                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LungNodule         *ScreenChestCTResponseBodyDataLungNodule         `json:"LungNodule,omitempty" xml:"LungNodule,omitempty" type:"Struct"`
	NestedUrlList      map[string]interface{}                           `json:"NestedUrlList,omitempty" xml:"NestedUrlList,omitempty"`
	ScreenEc           *ScreenChestCTResponseBodyDataScreenEc           `json:"ScreenEc,omitempty" xml:"ScreenEc,omitempty" type:"Struct"`
	URLList            map[string]interface{}                           `json:"URLList,omitempty" xml:"URLList,omitempty"`
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

func (s *ScreenChestCTResponseBodyData) SetDetectLymph(v *ScreenChestCTResponseBodyDataDetectLymph) *ScreenChestCTResponseBodyData {
	s.DetectLymph = v
	return s
}

func (s *ScreenChestCTResponseBodyData) SetDetectPdac(v *ScreenChestCTResponseBodyDataDetectPdac) *ScreenChestCTResponseBodyData {
	s.DetectPdac = v
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

func (s *ScreenChestCTResponseBodyData) SetNestedUrlList(v map[string]interface{}) *ScreenChestCTResponseBodyData {
	s.NestedUrlList = v
	return s
}

func (s *ScreenChestCTResponseBodyData) SetScreenEc(v *ScreenChestCTResponseBodyDataScreenEc) *ScreenChestCTResponseBodyData {
	s.ScreenEc = v
	return s
}

func (s *ScreenChestCTResponseBodyData) SetURLList(v map[string]interface{}) *ScreenChestCTResponseBodyData {
	s.URLList = v
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
	// 1
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
	// 1
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
	Detections        []*ScreenChestCTResponseBodyDataCACSDetections `json:"Detections,omitempty" xml:"Detections,omitempty" type:"Repeated"`
	ResultUrl         *string                                        `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
	Score             *string                                        `json:"Score,omitempty" xml:"Score,omitempty"`
	SeriesInstanceUID *string                                        `json:"SeriesInstanceUID,omitempty" xml:"SeriesInstanceUID,omitempty"`
	VolumeScore       *string                                        `json:"VolumeScore,omitempty" xml:"VolumeScore,omitempty"`
}

func (s ScreenChestCTResponseBodyDataCACS) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataCACS) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataCACS) SetDetections(v []*ScreenChestCTResponseBodyDataCACSDetections) *ScreenChestCTResponseBodyDataCACS {
	s.Detections = v
	return s
}

func (s *ScreenChestCTResponseBodyDataCACS) SetResultUrl(v string) *ScreenChestCTResponseBodyDataCACS {
	s.ResultUrl = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataCACS) SetScore(v string) *ScreenChestCTResponseBodyDataCACS {
	s.Score = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataCACS) SetSeriesInstanceUID(v string) *ScreenChestCTResponseBodyDataCACS {
	s.SeriesInstanceUID = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataCACS) SetVolumeScore(v string) *ScreenChestCTResponseBodyDataCACS {
	s.VolumeScore = &v
	return s
}

type ScreenChestCTResponseBodyDataCACSDetections struct {
	CalciumCenter []*int64 `json:"CalciumCenter,omitempty" xml:"CalciumCenter,omitempty" type:"Repeated"`
	CalciumId     *int64   `json:"CalciumId,omitempty" xml:"CalciumId,omitempty"`
	CalciumScore  *float32 `json:"CalciumScore,omitempty" xml:"CalciumScore,omitempty"`
	CalciumVolume *float32 `json:"CalciumVolume,omitempty" xml:"CalciumVolume,omitempty"`
}

func (s ScreenChestCTResponseBodyDataCACSDetections) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataCACSDetections) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataCACSDetections) SetCalciumCenter(v []*int64) *ScreenChestCTResponseBodyDataCACSDetections {
	s.CalciumCenter = v
	return s
}

func (s *ScreenChestCTResponseBodyDataCACSDetections) SetCalciumId(v int64) *ScreenChestCTResponseBodyDataCACSDetections {
	s.CalciumId = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataCACSDetections) SetCalciumScore(v float32) *ScreenChestCTResponseBodyDataCACSDetections {
	s.CalciumScore = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataCACSDetections) SetCalciumVolume(v float32) *ScreenChestCTResponseBodyDataCACSDetections {
	s.CalciumVolume = &v
	return s
}

type ScreenChestCTResponseBodyDataCovid struct {
	LesionRatio       *string `json:"LesionRatio,omitempty" xml:"LesionRatio,omitempty"`
	Mask              *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	NewProbability    *string `json:"NewProbability,omitempty" xml:"NewProbability,omitempty"`
	NormalProbability *string `json:"NormalProbability,omitempty" xml:"NormalProbability,omitempty"`
	OtherProbability  *string `json:"OtherProbability,omitempty" xml:"OtherProbability,omitempty"`
	SeriesInstanceUID *string `json:"SeriesInstanceUID,omitempty" xml:"SeriesInstanceUID,omitempty"`
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

func (s *ScreenChestCTResponseBodyDataCovid) SetSeriesInstanceUID(v string) *ScreenChestCTResponseBodyDataCovid {
	s.SeriesInstanceUID = &v
	return s
}

type ScreenChestCTResponseBodyDataDetectLymph struct {
	Lesions           []*ScreenChestCTResponseBodyDataDetectLymphLesions `json:"Lesions,omitempty" xml:"Lesions,omitempty" type:"Repeated"`
	SeriesInstanceUID *string                                            `json:"SeriesInstanceUID,omitempty" xml:"SeriesInstanceUID,omitempty"`
}

func (s ScreenChestCTResponseBodyDataDetectLymph) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataDetectLymph) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataDetectLymph) SetLesions(v []*ScreenChestCTResponseBodyDataDetectLymphLesions) *ScreenChestCTResponseBodyDataDetectLymph {
	s.Lesions = v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectLymph) SetSeriesInstanceUID(v string) *ScreenChestCTResponseBodyDataDetectLymph {
	s.SeriesInstanceUID = &v
	return s
}

type ScreenChestCTResponseBodyDataDetectLymphLesions struct {
	Boxes      []*float32   `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Diametermm []*float32   `json:"Diametermm,omitempty" xml:"Diametermm,omitempty" type:"Repeated"`
	KeySlice   *int64       `json:"KeySlice,omitempty" xml:"KeySlice,omitempty"`
	Recist     [][]*float32 `json:"Recist,omitempty" xml:"Recist,omitempty" type:"Repeated"`
	Score      *float32     `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s ScreenChestCTResponseBodyDataDetectLymphLesions) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataDetectLymphLesions) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataDetectLymphLesions) SetBoxes(v []*float32) *ScreenChestCTResponseBodyDataDetectLymphLesions {
	s.Boxes = v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectLymphLesions) SetDiametermm(v []*float32) *ScreenChestCTResponseBodyDataDetectLymphLesions {
	s.Diametermm = v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectLymphLesions) SetKeySlice(v int64) *ScreenChestCTResponseBodyDataDetectLymphLesions {
	s.KeySlice = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectLymphLesions) SetRecist(v [][]*float32) *ScreenChestCTResponseBodyDataDetectLymphLesions {
	s.Recist = v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectLymphLesions) SetScore(v float32) *ScreenChestCTResponseBodyDataDetectLymphLesions {
	s.Score = &v
	return s
}

type ScreenChestCTResponseBodyDataDetectPdac struct {
	Lesion            *ScreenChestCTResponseBodyDataDetectPdacLesion `json:"Lesion,omitempty" xml:"Lesion,omitempty" type:"Struct"`
	SeriesInstanceUID *string                                        `json:"SeriesInstanceUID,omitempty" xml:"SeriesInstanceUID,omitempty"`
}

func (s ScreenChestCTResponseBodyDataDetectPdac) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataDetectPdac) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataDetectPdac) SetLesion(v *ScreenChestCTResponseBodyDataDetectPdacLesion) *ScreenChestCTResponseBodyDataDetectPdac {
	s.Lesion = v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectPdac) SetSeriesInstanceUID(v string) *ScreenChestCTResponseBodyDataDetectPdac {
	s.SeriesInstanceUID = &v
	return s
}

type ScreenChestCTResponseBodyDataDetectPdacLesion struct {
	Mask          *string   `json:"Mask,omitempty" xml:"Mask,omitempty"`
	NonPdacVol    *string   `json:"NonPdacVol,omitempty" xml:"NonPdacVol,omitempty"`
	PancVol       *string   `json:"PancVol,omitempty" xml:"PancVol,omitempty"`
	PdacVol       *string   `json:"PdacVol,omitempty" xml:"PdacVol,omitempty"`
	Possibilities []*string `json:"Possibilities,omitempty" xml:"Possibilities,omitempty" type:"Repeated"`
}

func (s ScreenChestCTResponseBodyDataDetectPdacLesion) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataDetectPdacLesion) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataDetectPdacLesion) SetMask(v string) *ScreenChestCTResponseBodyDataDetectPdacLesion {
	s.Mask = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectPdacLesion) SetNonPdacVol(v string) *ScreenChestCTResponseBodyDataDetectPdacLesion {
	s.NonPdacVol = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectPdacLesion) SetPancVol(v string) *ScreenChestCTResponseBodyDataDetectPdacLesion {
	s.PancVol = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectPdacLesion) SetPdacVol(v string) *ScreenChestCTResponseBodyDataDetectPdacLesion {
	s.PdacVol = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectPdacLesion) SetPossibilities(v []*string) *ScreenChestCTResponseBodyDataDetectPdacLesion {
	s.Possibilities = v
	return s
}

type ScreenChestCTResponseBodyDataDetectRibFracture struct {
	Detections        []*ScreenChestCTResponseBodyDataDetectRibFractureDetections `json:"Detections,omitempty" xml:"Detections,omitempty" type:"Repeated"`
	FractureMaskURL   *string                                                     `json:"FractureMaskURL,omitempty" xml:"FractureMaskURL,omitempty"`
	Origin            []*float32                                                  `json:"Origin,omitempty" xml:"Origin,omitempty" type:"Repeated"`
	ResultURL         *string                                                     `json:"ResultURL,omitempty" xml:"ResultURL,omitempty"`
	RibSegmentMaskURL *string                                                     `json:"RibSegmentMaskURL,omitempty" xml:"RibSegmentMaskURL,omitempty"`
	SeriesInstanceUID *string                                                     `json:"SeriesInstanceUID,omitempty" xml:"SeriesInstanceUID,omitempty"`
	Spacing           []*float32                                                  `json:"Spacing,omitempty" xml:"Spacing,omitempty" type:"Repeated"`
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

func (s *ScreenChestCTResponseBodyDataDetectRibFracture) SetFractureMaskURL(v string) *ScreenChestCTResponseBodyDataDetectRibFracture {
	s.FractureMaskURL = &v
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

func (s *ScreenChestCTResponseBodyDataDetectRibFracture) SetRibSegmentMaskURL(v string) *ScreenChestCTResponseBodyDataDetectRibFracture {
	s.RibSegmentMaskURL = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectRibFracture) SetSeriesInstanceUID(v string) *ScreenChestCTResponseBodyDataDetectRibFracture {
	s.SeriesInstanceUID = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataDetectRibFracture) SetSpacing(v []*float32) *ScreenChestCTResponseBodyDataDetectRibFracture {
	s.Spacing = v
	return s
}

type ScreenChestCTResponseBodyDataDetectRibFractureDetections struct {
	CoordinateImage    []*int64 `json:"CoordinateImage,omitempty" xml:"CoordinateImage,omitempty" type:"Repeated"`
	Coordinates        []*int64 `json:"Coordinates,omitempty" xml:"Coordinates,omitempty" type:"Repeated"`
	FracSOPInstanceUID *string  `json:"FracSOPInstanceUID,omitempty" xml:"FracSOPInstanceUID,omitempty"`
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

func (s *ScreenChestCTResponseBodyDataDetectRibFractureDetections) SetFracSOPInstanceUID(v string) *ScreenChestCTResponseBodyDataDetectRibFractureDetections {
	s.FracSOPInstanceUID = &v
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
	Category             *string    `json:"Category,omitempty" xml:"Category,omitempty"`
	Confidence           *float32   `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Diameter             *float32   `json:"Diameter,omitempty" xml:"Diameter,omitempty"`
	ImageX               *float32   `json:"ImageX,omitempty" xml:"ImageX,omitempty"`
	ImageY               *float32   `json:"ImageY,omitempty" xml:"ImageY,omitempty"`
	ImageZ               *float32   `json:"ImageZ,omitempty" xml:"ImageZ,omitempty"`
	Lobe                 *string    `json:"Lobe,omitempty" xml:"Lobe,omitempty"`
	Lung                 *string    `json:"Lung,omitempty" xml:"Lung,omitempty"`
	MajorAxis            []*float32 `json:"MajorAxis,omitempty" xml:"MajorAxis,omitempty" type:"Repeated"`
	MeanValue            *float32   `json:"MeanValue,omitempty" xml:"MeanValue,omitempty"`
	MinorAxis            []*float32 `json:"MinorAxis,omitempty" xml:"MinorAxis,omitempty" type:"Repeated"`
	RecistSOPInstanceUID *string    `json:"RecistSOPInstanceUID,omitempty" xml:"RecistSOPInstanceUID,omitempty"`
	Risk                 *float32   `json:"Risk,omitempty" xml:"Risk,omitempty"`
	SOPInstanceUID       *string    `json:"SOPInstanceUID,omitempty" xml:"SOPInstanceUID,omitempty"`
	Volume               *float32   `json:"Volume,omitempty" xml:"Volume,omitempty"`
	X                    *float32   `json:"X,omitempty" xml:"X,omitempty"`
	Y                    *float32   `json:"Y,omitempty" xml:"Y,omitempty"`
	Z                    *float32   `json:"Z,omitempty" xml:"Z,omitempty"`
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

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetMajorAxis(v []*float32) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.MajorAxis = v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetMeanValue(v float32) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.MeanValue = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetMinorAxis(v []*float32) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.MinorAxis = v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetRecistSOPInstanceUID(v string) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.RecistSOPInstanceUID = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataLungNoduleSeriesElements) SetRisk(v float32) *ScreenChestCTResponseBodyDataLungNoduleSeriesElements {
	s.Risk = &v
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

type ScreenChestCTResponseBodyDataScreenEc struct {
	Lesion            *ScreenChestCTResponseBodyDataScreenEcLesion `json:"Lesion,omitempty" xml:"Lesion,omitempty" type:"Struct"`
	SeriesInstanceUid *string                                      `json:"SeriesInstanceUid,omitempty" xml:"SeriesInstanceUid,omitempty"`
}

func (s ScreenChestCTResponseBodyDataScreenEc) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataScreenEc) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataScreenEc) SetLesion(v *ScreenChestCTResponseBodyDataScreenEcLesion) *ScreenChestCTResponseBodyDataScreenEc {
	s.Lesion = v
	return s
}

func (s *ScreenChestCTResponseBodyDataScreenEc) SetSeriesInstanceUid(v string) *ScreenChestCTResponseBodyDataScreenEc {
	s.SeriesInstanceUid = &v
	return s
}

type ScreenChestCTResponseBodyDataScreenEcLesion struct {
	BenignVolume  *string   `json:"BenignVolume,omitempty" xml:"BenignVolume,omitempty"`
	EcVolume      *string   `json:"EcVolume,omitempty" xml:"EcVolume,omitempty"`
	EsoVolume     *string   `json:"EsoVolume,omitempty" xml:"EsoVolume,omitempty"`
	Mask          *string   `json:"Mask,omitempty" xml:"Mask,omitempty"`
	Possibilities []*string `json:"Possibilities,omitempty" xml:"Possibilities,omitempty" type:"Repeated"`
}

func (s ScreenChestCTResponseBodyDataScreenEcLesion) String() string {
	return tea.Prettify(s)
}

func (s ScreenChestCTResponseBodyDataScreenEcLesion) GoString() string {
	return s.String()
}

func (s *ScreenChestCTResponseBodyDataScreenEcLesion) SetBenignVolume(v string) *ScreenChestCTResponseBodyDataScreenEcLesion {
	s.BenignVolume = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataScreenEcLesion) SetEcVolume(v string) *ScreenChestCTResponseBodyDataScreenEcLesion {
	s.EcVolume = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataScreenEcLesion) SetEsoVolume(v string) *ScreenChestCTResponseBodyDataScreenEcLesion {
	s.EsoVolume = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataScreenEcLesion) SetMask(v string) *ScreenChestCTResponseBodyDataScreenEcLesion {
	s.Mask = &v
	return s
}

func (s *ScreenChestCTResponseBodyDataScreenEcLesion) SetPossibilities(v []*string) *ScreenChestCTResponseBodyDataScreenEcLesion {
	s.Possibilities = v
	return s
}

type ScreenChestCTResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ScreenChestCTResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ScreenChestCTResponse) SetStatusCode(v int32) *ScreenChestCTResponse {
	s.StatusCode = &v
	return s
}

func (s *ScreenChestCTResponse) SetBody(v *ScreenChestCTResponseBody) *ScreenChestCTResponse {
	s.Body = v
	return s
}

type ScreenECRequest struct {
	DataSourceType *string                   `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	URLList        []*ScreenECRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s ScreenECRequest) String() string {
	return tea.Prettify(s)
}

func (s ScreenECRequest) GoString() string {
	return s.String()
}

func (s *ScreenECRequest) SetDataSourceType(v string) *ScreenECRequest {
	s.DataSourceType = &v
	return s
}

func (s *ScreenECRequest) SetURLList(v []*ScreenECRequestURLList) *ScreenECRequest {
	s.URLList = v
	return s
}

type ScreenECRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s ScreenECRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s ScreenECRequestURLList) GoString() string {
	return s.String()
}

func (s *ScreenECRequestURLList) SetURL(v string) *ScreenECRequestURLList {
	s.URL = &v
	return s
}

type ScreenECResponseBody struct {
	Data      *ScreenECResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScreenECResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScreenECResponseBody) GoString() string {
	return s.String()
}

func (s *ScreenECResponseBody) SetData(v *ScreenECResponseBodyData) *ScreenECResponseBody {
	s.Data = v
	return s
}

func (s *ScreenECResponseBody) SetMessage(v string) *ScreenECResponseBody {
	s.Message = &v
	return s
}

func (s *ScreenECResponseBody) SetRequestId(v string) *ScreenECResponseBody {
	s.RequestId = &v
	return s
}

type ScreenECResponseBodyData struct {
	Lesion *ScreenECResponseBodyDataLesion `json:"Lesion,omitempty" xml:"Lesion,omitempty" type:"Struct"`
}

func (s ScreenECResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ScreenECResponseBodyData) GoString() string {
	return s.String()
}

func (s *ScreenECResponseBodyData) SetLesion(v *ScreenECResponseBodyDataLesion) *ScreenECResponseBodyData {
	s.Lesion = v
	return s
}

type ScreenECResponseBodyDataLesion struct {
	BenignVolume  *string   `json:"BenignVolume,omitempty" xml:"BenignVolume,omitempty"`
	EcVolume      *string   `json:"EcVolume,omitempty" xml:"EcVolume,omitempty"`
	EsoVolume     *string   `json:"EsoVolume,omitempty" xml:"EsoVolume,omitempty"`
	Mask          *string   `json:"Mask,omitempty" xml:"Mask,omitempty"`
	Possibilities []*string `json:"Possibilities,omitempty" xml:"Possibilities,omitempty" type:"Repeated"`
}

func (s ScreenECResponseBodyDataLesion) String() string {
	return tea.Prettify(s)
}

func (s ScreenECResponseBodyDataLesion) GoString() string {
	return s.String()
}

func (s *ScreenECResponseBodyDataLesion) SetBenignVolume(v string) *ScreenECResponseBodyDataLesion {
	s.BenignVolume = &v
	return s
}

func (s *ScreenECResponseBodyDataLesion) SetEcVolume(v string) *ScreenECResponseBodyDataLesion {
	s.EcVolume = &v
	return s
}

func (s *ScreenECResponseBodyDataLesion) SetEsoVolume(v string) *ScreenECResponseBodyDataLesion {
	s.EsoVolume = &v
	return s
}

func (s *ScreenECResponseBodyDataLesion) SetMask(v string) *ScreenECResponseBodyDataLesion {
	s.Mask = &v
	return s
}

func (s *ScreenECResponseBodyDataLesion) SetPossibilities(v []*string) *ScreenECResponseBodyDataLesion {
	s.Possibilities = v
	return s
}

type ScreenECResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ScreenECResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScreenECResponse) String() string {
	return tea.Prettify(s)
}

func (s ScreenECResponse) GoString() string {
	return s.String()
}

func (s *ScreenECResponse) SetHeaders(v map[string]*string) *ScreenECResponse {
	s.Headers = v
	return s
}

func (s *ScreenECResponse) SetStatusCode(v int32) *ScreenECResponse {
	s.StatusCode = &v
	return s
}

func (s *ScreenECResponse) SetBody(v *ScreenECResponseBody) *ScreenECResponse {
	s.Body = v
	return s
}

type SegmentLymphNodeRequest struct {
	BodyPart   *string                           `json:"BodyPart,omitempty" xml:"BodyPart,omitempty"`
	DataFormat *string                           `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId      *string                           `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string                           `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	URLList    []*SegmentLymphNodeRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s SegmentLymphNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentLymphNodeRequest) GoString() string {
	return s.String()
}

func (s *SegmentLymphNodeRequest) SetBodyPart(v string) *SegmentLymphNodeRequest {
	s.BodyPart = &v
	return s
}

func (s *SegmentLymphNodeRequest) SetDataFormat(v string) *SegmentLymphNodeRequest {
	s.DataFormat = &v
	return s
}

func (s *SegmentLymphNodeRequest) SetOrgId(v string) *SegmentLymphNodeRequest {
	s.OrgId = &v
	return s
}

func (s *SegmentLymphNodeRequest) SetOrgName(v string) *SegmentLymphNodeRequest {
	s.OrgName = &v
	return s
}

func (s *SegmentLymphNodeRequest) SetURLList(v []*SegmentLymphNodeRequestURLList) *SegmentLymphNodeRequest {
	s.URLList = v
	return s
}

type SegmentLymphNodeRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s SegmentLymphNodeRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s SegmentLymphNodeRequestURLList) GoString() string {
	return s.String()
}

func (s *SegmentLymphNodeRequestURLList) SetURL(v string) *SegmentLymphNodeRequestURLList {
	s.URL = &v
	return s
}

type SegmentLymphNodeAdvanceRequest struct {
	BodyPart   *string                                  `json:"BodyPart,omitempty" xml:"BodyPart,omitempty"`
	DataFormat *string                                  `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId      *string                                  `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string                                  `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	URLList    []*SegmentLymphNodeAdvanceRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s SegmentLymphNodeAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentLymphNodeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentLymphNodeAdvanceRequest) SetBodyPart(v string) *SegmentLymphNodeAdvanceRequest {
	s.BodyPart = &v
	return s
}

func (s *SegmentLymphNodeAdvanceRequest) SetDataFormat(v string) *SegmentLymphNodeAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *SegmentLymphNodeAdvanceRequest) SetOrgId(v string) *SegmentLymphNodeAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *SegmentLymphNodeAdvanceRequest) SetOrgName(v string) *SegmentLymphNodeAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *SegmentLymphNodeAdvanceRequest) SetURLList(v []*SegmentLymphNodeAdvanceRequestURLList) *SegmentLymphNodeAdvanceRequest {
	s.URLList = v
	return s
}

type SegmentLymphNodeAdvanceRequestURLList struct {
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s SegmentLymphNodeAdvanceRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s SegmentLymphNodeAdvanceRequestURLList) GoString() string {
	return s.String()
}

func (s *SegmentLymphNodeAdvanceRequestURLList) SetURLObject(v io.Reader) *SegmentLymphNodeAdvanceRequestURLList {
	s.URLObject = v
	return s
}

type SegmentLymphNodeResponseBody struct {
	Data      *SegmentLymphNodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentLymphNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentLymphNodeResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentLymphNodeResponseBody) SetData(v *SegmentLymphNodeResponseBodyData) *SegmentLymphNodeResponseBody {
	s.Data = v
	return s
}

func (s *SegmentLymphNodeResponseBody) SetMessage(v string) *SegmentLymphNodeResponseBody {
	s.Message = &v
	return s
}

func (s *SegmentLymphNodeResponseBody) SetRequestId(v string) *SegmentLymphNodeResponseBody {
	s.RequestId = &v
	return s
}

type SegmentLymphNodeResponseBodyData struct {
	ResultURL *string `json:"ResultURL,omitempty" xml:"ResultURL,omitempty"`
}

func (s SegmentLymphNodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentLymphNodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentLymphNodeResponseBodyData) SetResultURL(v string) *SegmentLymphNodeResponseBodyData {
	s.ResultURL = &v
	return s
}

type SegmentLymphNodeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentLymphNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SegmentLymphNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentLymphNodeResponse) GoString() string {
	return s.String()
}

func (s *SegmentLymphNodeResponse) SetHeaders(v map[string]*string) *SegmentLymphNodeResponse {
	s.Headers = v
	return s
}

func (s *SegmentLymphNodeResponse) SetStatusCode(v int32) *SegmentLymphNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentLymphNodeResponse) SetBody(v *SegmentLymphNodeResponseBody) *SegmentLymphNodeResponse {
	s.Body = v
	return s
}

type SegmentOARRequest struct {
	BodyPart   *string                     `json:"BodyPart,omitempty" xml:"BodyPart,omitempty"`
	Contrast   *bool                       `json:"Contrast,omitempty" xml:"Contrast,omitempty"`
	DataFormat *string                     `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	MaskList   []*int64                    `json:"MaskList,omitempty" xml:"MaskList,omitempty" type:"Repeated"`
	OrgId      *string                     `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string                     `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	URLList    []*SegmentOARRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s SegmentOARRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentOARRequest) GoString() string {
	return s.String()
}

func (s *SegmentOARRequest) SetBodyPart(v string) *SegmentOARRequest {
	s.BodyPart = &v
	return s
}

func (s *SegmentOARRequest) SetContrast(v bool) *SegmentOARRequest {
	s.Contrast = &v
	return s
}

func (s *SegmentOARRequest) SetDataFormat(v string) *SegmentOARRequest {
	s.DataFormat = &v
	return s
}

func (s *SegmentOARRequest) SetMaskList(v []*int64) *SegmentOARRequest {
	s.MaskList = v
	return s
}

func (s *SegmentOARRequest) SetOrgId(v string) *SegmentOARRequest {
	s.OrgId = &v
	return s
}

func (s *SegmentOARRequest) SetOrgName(v string) *SegmentOARRequest {
	s.OrgName = &v
	return s
}

func (s *SegmentOARRequest) SetURLList(v []*SegmentOARRequestURLList) *SegmentOARRequest {
	s.URLList = v
	return s
}

type SegmentOARRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s SegmentOARRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s SegmentOARRequestURLList) GoString() string {
	return s.String()
}

func (s *SegmentOARRequestURLList) SetURL(v string) *SegmentOARRequestURLList {
	s.URL = &v
	return s
}

type SegmentOARAdvanceRequest struct {
	BodyPart   *string                            `json:"BodyPart,omitempty" xml:"BodyPart,omitempty"`
	Contrast   *bool                              `json:"Contrast,omitempty" xml:"Contrast,omitempty"`
	DataFormat *string                            `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	MaskList   []*int64                           `json:"MaskList,omitempty" xml:"MaskList,omitempty" type:"Repeated"`
	OrgId      *string                            `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName    *string                            `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	URLList    []*SegmentOARAdvanceRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s SegmentOARAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentOARAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentOARAdvanceRequest) SetBodyPart(v string) *SegmentOARAdvanceRequest {
	s.BodyPart = &v
	return s
}

func (s *SegmentOARAdvanceRequest) SetContrast(v bool) *SegmentOARAdvanceRequest {
	s.Contrast = &v
	return s
}

func (s *SegmentOARAdvanceRequest) SetDataFormat(v string) *SegmentOARAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *SegmentOARAdvanceRequest) SetMaskList(v []*int64) *SegmentOARAdvanceRequest {
	s.MaskList = v
	return s
}

func (s *SegmentOARAdvanceRequest) SetOrgId(v string) *SegmentOARAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *SegmentOARAdvanceRequest) SetOrgName(v string) *SegmentOARAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *SegmentOARAdvanceRequest) SetURLList(v []*SegmentOARAdvanceRequestURLList) *SegmentOARAdvanceRequest {
	s.URLList = v
	return s
}

type SegmentOARAdvanceRequestURLList struct {
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s SegmentOARAdvanceRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s SegmentOARAdvanceRequestURLList) GoString() string {
	return s.String()
}

func (s *SegmentOARAdvanceRequestURLList) SetURLObject(v io.Reader) *SegmentOARAdvanceRequestURLList {
	s.URLObject = v
	return s
}

type SegmentOARResponseBody struct {
	Data      *SegmentOARResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentOARResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentOARResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentOARResponseBody) SetData(v *SegmentOARResponseBodyData) *SegmentOARResponseBody {
	s.Data = v
	return s
}

func (s *SegmentOARResponseBody) SetMessage(v string) *SegmentOARResponseBody {
	s.Message = &v
	return s
}

func (s *SegmentOARResponseBody) SetRequestId(v string) *SegmentOARResponseBody {
	s.RequestId = &v
	return s
}

type SegmentOARResponseBodyData struct {
	ResultURL *string `json:"ResultURL,omitempty" xml:"ResultURL,omitempty"`
}

func (s SegmentOARResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentOARResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentOARResponseBodyData) SetResultURL(v string) *SegmentOARResponseBodyData {
	s.ResultURL = &v
	return s
}

type SegmentOARResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SegmentOARResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SegmentOARResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentOARResponse) GoString() string {
	return s.String()
}

func (s *SegmentOARResponse) SetHeaders(v map[string]*string) *SegmentOARResponse {
	s.Headers = v
	return s
}

func (s *SegmentOARResponse) SetStatusCode(v int32) *SegmentOARResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentOARResponse) SetBody(v *SegmentOARResponseBody) *SegmentOARResponse {
	s.Body = v
	return s
}

type TargetVolumeSegmentRequest struct {
	CancerType       *string                              `json:"CancerType,omitempty" xml:"CancerType,omitempty"`
	DataFormat       *string                              `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId            *string                              `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName          *string                              `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	TargetVolumeType *string                              `json:"TargetVolumeType,omitempty" xml:"TargetVolumeType,omitempty"`
	URLList          []*TargetVolumeSegmentRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s TargetVolumeSegmentRequest) String() string {
	return tea.Prettify(s)
}

func (s TargetVolumeSegmentRequest) GoString() string {
	return s.String()
}

func (s *TargetVolumeSegmentRequest) SetCancerType(v string) *TargetVolumeSegmentRequest {
	s.CancerType = &v
	return s
}

func (s *TargetVolumeSegmentRequest) SetDataFormat(v string) *TargetVolumeSegmentRequest {
	s.DataFormat = &v
	return s
}

func (s *TargetVolumeSegmentRequest) SetOrgId(v string) *TargetVolumeSegmentRequest {
	s.OrgId = &v
	return s
}

func (s *TargetVolumeSegmentRequest) SetOrgName(v string) *TargetVolumeSegmentRequest {
	s.OrgName = &v
	return s
}

func (s *TargetVolumeSegmentRequest) SetTargetVolumeType(v string) *TargetVolumeSegmentRequest {
	s.TargetVolumeType = &v
	return s
}

func (s *TargetVolumeSegmentRequest) SetURLList(v []*TargetVolumeSegmentRequestURLList) *TargetVolumeSegmentRequest {
	s.URLList = v
	return s
}

type TargetVolumeSegmentRequestURLList struct {
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s TargetVolumeSegmentRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s TargetVolumeSegmentRequestURLList) GoString() string {
	return s.String()
}

func (s *TargetVolumeSegmentRequestURLList) SetURL(v string) *TargetVolumeSegmentRequestURLList {
	s.URL = &v
	return s
}

type TargetVolumeSegmentAdvanceRequest struct {
	CancerType       *string                                     `json:"CancerType,omitempty" xml:"CancerType,omitempty"`
	DataFormat       *string                                     `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	OrgId            *string                                     `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName          *string                                     `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	TargetVolumeType *string                                     `json:"TargetVolumeType,omitempty" xml:"TargetVolumeType,omitempty"`
	URLList          []*TargetVolumeSegmentAdvanceRequestURLList `json:"URLList,omitempty" xml:"URLList,omitempty" type:"Repeated"`
}

func (s TargetVolumeSegmentAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s TargetVolumeSegmentAdvanceRequest) GoString() string {
	return s.String()
}

func (s *TargetVolumeSegmentAdvanceRequest) SetCancerType(v string) *TargetVolumeSegmentAdvanceRequest {
	s.CancerType = &v
	return s
}

func (s *TargetVolumeSegmentAdvanceRequest) SetDataFormat(v string) *TargetVolumeSegmentAdvanceRequest {
	s.DataFormat = &v
	return s
}

func (s *TargetVolumeSegmentAdvanceRequest) SetOrgId(v string) *TargetVolumeSegmentAdvanceRequest {
	s.OrgId = &v
	return s
}

func (s *TargetVolumeSegmentAdvanceRequest) SetOrgName(v string) *TargetVolumeSegmentAdvanceRequest {
	s.OrgName = &v
	return s
}

func (s *TargetVolumeSegmentAdvanceRequest) SetTargetVolumeType(v string) *TargetVolumeSegmentAdvanceRequest {
	s.TargetVolumeType = &v
	return s
}

func (s *TargetVolumeSegmentAdvanceRequest) SetURLList(v []*TargetVolumeSegmentAdvanceRequestURLList) *TargetVolumeSegmentAdvanceRequest {
	s.URLList = v
	return s
}

type TargetVolumeSegmentAdvanceRequestURLList struct {
	URLObject io.Reader `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s TargetVolumeSegmentAdvanceRequestURLList) String() string {
	return tea.Prettify(s)
}

func (s TargetVolumeSegmentAdvanceRequestURLList) GoString() string {
	return s.String()
}

func (s *TargetVolumeSegmentAdvanceRequestURLList) SetURLObject(v io.Reader) *TargetVolumeSegmentAdvanceRequestURLList {
	s.URLObject = v
	return s
}

type TargetVolumeSegmentResponseBody struct {
	Data      *TargetVolumeSegmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TargetVolumeSegmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TargetVolumeSegmentResponseBody) GoString() string {
	return s.String()
}

func (s *TargetVolumeSegmentResponseBody) SetData(v *TargetVolumeSegmentResponseBodyData) *TargetVolumeSegmentResponseBody {
	s.Data = v
	return s
}

func (s *TargetVolumeSegmentResponseBody) SetMessage(v string) *TargetVolumeSegmentResponseBody {
	s.Message = &v
	return s
}

func (s *TargetVolumeSegmentResponseBody) SetRequestId(v string) *TargetVolumeSegmentResponseBody {
	s.RequestId = &v
	return s
}

type TargetVolumeSegmentResponseBodyData struct {
	ResultURL *string `json:"ResultURL,omitempty" xml:"ResultURL,omitempty"`
}

func (s TargetVolumeSegmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TargetVolumeSegmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *TargetVolumeSegmentResponseBodyData) SetResultURL(v string) *TargetVolumeSegmentResponseBodyData {
	s.ResultURL = &v
	return s
}

type TargetVolumeSegmentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TargetVolumeSegmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TargetVolumeSegmentResponse) String() string {
	return tea.Prettify(s)
}

func (s TargetVolumeSegmentResponse) GoString() string {
	return s.String()
}

func (s *TargetVolumeSegmentResponse) SetHeaders(v map[string]*string) *TargetVolumeSegmentResponse {
	s.Headers = v
	return s
}

func (s *TargetVolumeSegmentResponse) SetStatusCode(v int32) *TargetVolumeSegmentResponse {
	s.StatusCode = &v
	return s
}

func (s *TargetVolumeSegmentResponse) SetBody(v *TargetVolumeSegmentResponseBody) *TargetVolumeSegmentResponse {
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TranslateMedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TranslateMedResponse) SetStatusCode(v int32) *TranslateMedResponse {
	s.StatusCode = &v
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

func (client *Client) AnalyzeChestVesselAdvance(request *AnalyzeChestVesselAdvanceRequest, runtime *util.RuntimeOptions) (_result *AnalyzeChestVesselResponse, _err error) {
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

	authConfig := &openapi.Config{
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
	analyzeChestVesselReq := &AnalyzeChestVesselRequest{}
	openapiutil.Convert(request, analyzeChestVesselReq)
	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.URLList {
			if !tea.BoolValue(util.IsUnset(item0.URLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.URLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := analyzeChestVesselReq.URLList[tea.IntValue(i0)]
				tmp.URL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	analyzeChestVesselResp, _err := client.AnalyzeChestVesselWithOptions(analyzeChestVesselReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = analyzeChestVesselResp
	return _result, _err
}

func (client *Client) CalcBMDWithOptions(request *CalcBMDRequest, runtime *util.RuntimeOptions) (_result *CalcBMDResponse, _err error) {
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
		Action:      tea.String("CalcBMD"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CalcBMDResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CalcBMD(request *CalcBMDRequest) (_result *CalcBMDResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CalcBMDResponse{}
	_body, _err := client.CalcBMDWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CalcBMDAdvance(request *CalcBMDAdvanceRequest, runtime *util.RuntimeOptions) (_result *CalcBMDResponse, _err error) {
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

	authConfig := &openapi.Config{
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
	calcBMDReq := &CalcBMDRequest{}
	openapiutil.Convert(request, calcBMDReq)
	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.URLList {
			if !tea.BoolValue(util.IsUnset(item0.URLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.URLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := calcBMDReq.URLList[tea.IntValue(i0)]
				tmp.URL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	calcBMDResp, _err := client.CalcBMDWithOptions(calcBMDReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = calcBMDResp
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

func (client *Client) CalcCACSAdvance(request *CalcCACSAdvanceRequest, runtime *util.RuntimeOptions) (_result *CalcCACSResponse, _err error) {
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

	authConfig := &openapi.Config{
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
	calcCACSReq := &CalcCACSRequest{}
	openapiutil.Convert(request, calcCACSReq)
	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.URLList {
			if !tea.BoolValue(util.IsUnset(item0.URLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.URLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := calcCACSReq.URLList[tea.IntValue(i0)]
				tmp.URL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	calcCACSResp, _err := client.CalcCACSWithOptions(calcCACSReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = calcCACSResp
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

	authConfig := &openapi.Config{
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

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		classifyFNFReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
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

func (client *Client) DetectCovid19CadAdvance(request *DetectCovid19CadAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectCovid19CadResponse, _err error) {
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

	authConfig := &openapi.Config{
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
	detectCovid19CadReq := &DetectCovid19CadRequest{}
	openapiutil.Convert(request, detectCovid19CadReq)
	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.URLList {
			if !tea.BoolValue(util.IsUnset(item0.URLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.URLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := detectCovid19CadReq.URLList[tea.IntValue(i0)]
				tmp.URL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	detectCovid19CadResp, _err := client.DetectCovid19CadWithOptions(detectCovid19CadReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectCovid19CadResp
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

	authConfig := &openapi.Config{
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

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectHipKeypointXRayReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
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

	authConfig := &openapi.Config{
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

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectKneeKeypointXRayReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
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

	authConfig := &openapi.Config{
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

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.UrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectKneeXRayReq.Url = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
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

func (client *Client) DetectLungNoduleAdvance(request *DetectLungNoduleAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectLungNoduleResponse, _err error) {
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

	authConfig := &openapi.Config{
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
	detectLungNoduleReq := &DetectLungNoduleRequest{}
	openapiutil.Convert(request, detectLungNoduleReq)
	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.URLList {
			if !tea.BoolValue(util.IsUnset(item0.URLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.URLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := detectLungNoduleReq.URLList[tea.IntValue(i0)]
				tmp.URL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	detectLungNoduleResp, _err := client.DetectLungNoduleWithOptions(detectLungNoduleReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectLungNoduleResp
	return _result, _err
}

func (client *Client) DetectLymphWithOptions(request *DetectLymphRequest, runtime *util.RuntimeOptions) (_result *DetectLymphResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		body["URLList"] = request.URLList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectLymph"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectLymphResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectLymph(request *DetectLymphRequest) (_result *DetectLymphResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectLymphResponse{}
	_body, _err := client.DetectLymphWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectLymphAdvance(request *DetectLymphAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectLymphResponse, _err error) {
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

	authConfig := &openapi.Config{
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
	detectLymphReq := &DetectLymphRequest{}
	openapiutil.Convert(request, detectLymphReq)
	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.URLList {
			if !tea.BoolValue(util.IsUnset(item0.URLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.URLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := detectLymphReq.URLList[tea.IntValue(i0)]
				tmp.URL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	detectLymphResp, _err := client.DetectLymphWithOptions(detectLymphReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectLymphResp
	return _result, _err
}

func (client *Client) DetectPancWithOptions(request *DetectPancRequest, runtime *util.RuntimeOptions) (_result *DetectPancResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		body["URLList"] = request.URLList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectPanc"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectPancResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectPanc(request *DetectPancRequest) (_result *DetectPancResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectPancResponse{}
	_body, _err := client.DetectPancWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectPancAdvance(request *DetectPancAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectPancResponse, _err error) {
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

	authConfig := &openapi.Config{
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
	detectPancReq := &DetectPancRequest{}
	openapiutil.Convert(request, detectPancReq)
	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.URLList {
			if !tea.BoolValue(util.IsUnset(item0.URLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.URLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := detectPancReq.URLList[tea.IntValue(i0)]
				tmp.URL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	detectPancResp, _err := client.DetectPancWithOptions(detectPancReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectPancResp
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

func (client *Client) DetectRibFractureAdvance(request *DetectRibFractureAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectRibFractureResponse, _err error) {
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

	authConfig := &openapi.Config{
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
	detectRibFractureReq := &DetectRibFractureRequest{}
	openapiutil.Convert(request, detectRibFractureReq)
	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.URLList {
			if !tea.BoolValue(util.IsUnset(item0.URLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.URLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := detectRibFractureReq.URLList[tea.IntValue(i0)]
				tmp.URL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	detectRibFractureResp, _err := client.DetectRibFractureWithOptions(detectRibFractureReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectRibFractureResp
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

	authConfig := &openapi.Config{
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

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.UrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectSkinDiseaseReq.Url = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
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

func (client *Client) DetectSpineMRIAdvance(request *DetectSpineMRIAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectSpineMRIResponse, _err error) {
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

	authConfig := &openapi.Config{
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
	detectSpineMRIReq := &DetectSpineMRIRequest{}
	openapiutil.Convert(request, detectSpineMRIReq)
	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.URLList {
			if !tea.BoolValue(util.IsUnset(item0.URLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.URLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := detectSpineMRIReq.URLList[tea.IntValue(i0)]
				tmp.URL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	detectSpineMRIResp, _err := client.DetectSpineMRIWithOptions(detectSpineMRIReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectSpineMRIResp
	return _result, _err
}

func (client *Client) FeedbackSessionWithOptions(request *FeedbackSessionRequest, runtime *util.RuntimeOptions) (_result *FeedbackSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Feedback)) {
		body["Feedback"] = request.Feedback
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FeedbackSession"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FeedbackSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FeedbackSession(request *FeedbackSessionRequest) (_result *FeedbackSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FeedbackSessionResponse{}
	_body, _err := client.FeedbackSessionWithOptions(request, runtime)
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

func (client *Client) RunCTRegistrationAdvance(request *RunCTRegistrationAdvanceRequest, runtime *util.RuntimeOptions) (_result *RunCTRegistrationResponse, _err error) {
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

	authConfig := &openapi.Config{
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
	runCTRegistrationReq := &RunCTRegistrationRequest{}
	openapiutil.Convert(request, runCTRegistrationReq)
	if !tea.BoolValue(util.IsUnset(request.FloatingList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.FloatingList {
			if !tea.BoolValue(util.IsUnset(item0.FloatingURLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.FloatingURLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := runCTRegistrationReq.FloatingList[tea.IntValue(i0)]
				tmp.FloatingURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	if !tea.BoolValue(util.IsUnset(request.ReferenceList)) {
		i1 := tea.Int(0)
		for _, item0 := range request.ReferenceList {
			if !tea.BoolValue(util.IsUnset(item0.ReferenceURLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.ReferenceURLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := runCTRegistrationReq.ReferenceList[tea.IntValue(i1)]
				tmp.ReferenceURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i1 = number.Ltoi(number.Add(number.Itol(i1), number.Itol(tea.Int(1))))
			}

		}
	}

	runCTRegistrationResp, _err := client.RunCTRegistrationWithOptions(runCTRegistrationReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = runCTRegistrationResp
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

func (client *Client) RunMedQAAdvance(request *RunMedQAAdvanceRequest, runtime *util.RuntimeOptions) (_result *RunMedQAResponse, _err error) {
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

	authConfig := &openapi.Config{
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
	runMedQAReq := &RunMedQARequest{}
	openapiutil.Convert(request, runMedQAReq)
	if !tea.BoolValue(util.IsUnset(request.AnswerImageURLList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.AnswerImageURLList {
			if !tea.BoolValue(util.IsUnset(item0.AnswerImageURLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.AnswerImageURLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := runMedQAReq.AnswerImageURLList[tea.IntValue(i0)]
				tmp.AnswerImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	runMedQAResp, _err := client.RunMedQAWithOptions(runMedQAReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = runMedQAResp
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

	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		body["Verbose"] = request.Verbose
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

func (client *Client) ScreenChestCTAdvance(request *ScreenChestCTAdvanceRequest, runtime *util.RuntimeOptions) (_result *ScreenChestCTResponse, _err error) {
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

	authConfig := &openapi.Config{
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
	screenChestCTReq := &ScreenChestCTRequest{}
	openapiutil.Convert(request, screenChestCTReq)
	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.URLList {
			if !tea.BoolValue(util.IsUnset(item0.URLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.URLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := screenChestCTReq.URLList[tea.IntValue(i0)]
				tmp.URL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	screenChestCTResp, _err := client.ScreenChestCTWithOptions(screenChestCTReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = screenChestCTResp
	return _result, _err
}

func (client *Client) ScreenECWithOptions(request *ScreenECRequest, runtime *util.RuntimeOptions) (_result *ScreenECResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		body["URLList"] = request.URLList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ScreenEC"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ScreenECResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScreenEC(request *ScreenECRequest) (_result *ScreenECResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScreenECResponse{}
	_body, _err := client.ScreenECWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentLymphNodeWithOptions(request *SegmentLymphNodeRequest, runtime *util.RuntimeOptions) (_result *SegmentLymphNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyPart)) {
		body["BodyPart"] = request.BodyPart
	}

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
		Action:      tea.String("SegmentLymphNode"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentLymphNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentLymphNode(request *SegmentLymphNodeRequest) (_result *SegmentLymphNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentLymphNodeResponse{}
	_body, _err := client.SegmentLymphNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentLymphNodeAdvance(request *SegmentLymphNodeAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentLymphNodeResponse, _err error) {
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

	authConfig := &openapi.Config{
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
	segmentLymphNodeReq := &SegmentLymphNodeRequest{}
	openapiutil.Convert(request, segmentLymphNodeReq)
	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.URLList {
			if !tea.BoolValue(util.IsUnset(item0.URLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.URLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := segmentLymphNodeReq.URLList[tea.IntValue(i0)]
				tmp.URL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	segmentLymphNodeResp, _err := client.SegmentLymphNodeWithOptions(segmentLymphNodeReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentLymphNodeResp
	return _result, _err
}

func (client *Client) SegmentOARWithOptions(request *SegmentOARRequest, runtime *util.RuntimeOptions) (_result *SegmentOARResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyPart)) {
		body["BodyPart"] = request.BodyPart
	}

	if !tea.BoolValue(util.IsUnset(request.Contrast)) {
		body["Contrast"] = request.Contrast
	}

	if !tea.BoolValue(util.IsUnset(request.DataFormat)) {
		body["DataFormat"] = request.DataFormat
	}

	if !tea.BoolValue(util.IsUnset(request.MaskList)) {
		body["MaskList"] = request.MaskList
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
		Action:      tea.String("SegmentOAR"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SegmentOARResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentOAR(request *SegmentOARRequest) (_result *SegmentOARResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentOARResponse{}
	_body, _err := client.SegmentOARWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentOARAdvance(request *SegmentOARAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentOARResponse, _err error) {
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

	authConfig := &openapi.Config{
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
	segmentOARReq := &SegmentOARRequest{}
	openapiutil.Convert(request, segmentOARReq)
	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.URLList {
			if !tea.BoolValue(util.IsUnset(item0.URLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.URLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := segmentOARReq.URLList[tea.IntValue(i0)]
				tmp.URL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	segmentOARResp, _err := client.SegmentOARWithOptions(segmentOARReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentOARResp
	return _result, _err
}

func (client *Client) TargetVolumeSegmentWithOptions(request *TargetVolumeSegmentRequest, runtime *util.RuntimeOptions) (_result *TargetVolumeSegmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CancerType)) {
		body["CancerType"] = request.CancerType
	}

	if !tea.BoolValue(util.IsUnset(request.DataFormat)) {
		body["DataFormat"] = request.DataFormat
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetVolumeType)) {
		body["TargetVolumeType"] = request.TargetVolumeType
	}

	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		body["URLList"] = request.URLList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TargetVolumeSegment"),
		Version:     tea.String("2020-03-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TargetVolumeSegmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TargetVolumeSegment(request *TargetVolumeSegmentRequest) (_result *TargetVolumeSegmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TargetVolumeSegmentResponse{}
	_body, _err := client.TargetVolumeSegmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TargetVolumeSegmentAdvance(request *TargetVolumeSegmentAdvanceRequest, runtime *util.RuntimeOptions) (_result *TargetVolumeSegmentResponse, _err error) {
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

	authConfig := &openapi.Config{
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
	targetVolumeSegmentReq := &TargetVolumeSegmentRequest{}
	openapiutil.Convert(request, targetVolumeSegmentReq)
	if !tea.BoolValue(util.IsUnset(request.URLList)) {
		i0 := tea.Int(0)
		for _, item0 := range request.URLList {
			if !tea.BoolValue(util.IsUnset(item0.URLObject)) {
				authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
				if _err != nil {
					return _result, _err
				}

				ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
				ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
				ossClient, _err = oss.NewClient(ossConfig)
				if _err != nil {
					return _result, _err
				}

				fileObj = &fileform.FileField{
					Filename:    authResponse.Body.ObjectKey,
					Content:     item0.URLObject,
					ContentType: tea.String(""),
				}
				ossHeader = &oss.PostObjectRequestHeader{
					AccessKeyId:         authResponse.Body.AccessKeyId,
					Policy:              authResponse.Body.EncodedPolicy,
					Signature:           authResponse.Body.Signature,
					Key:                 authResponse.Body.ObjectKey,
					File:                fileObj,
					SuccessActionStatus: tea.String("201"),
				}
				uploadRequest = &oss.PostObjectRequest{
					BucketName: authResponse.Body.Bucket,
					Header:     ossHeader,
				}
				_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
				if _err != nil {
					return _result, _err
				}
				tmp := targetVolumeSegmentReq.URLList[tea.IntValue(i0)]
				tmp.URL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i0 = number.Ltoi(number.Add(number.Itol(i0), number.Itol(tea.Int(1))))
			}

		}
	}

	targetVolumeSegmentResp, _err := client.TargetVolumeSegmentWithOptions(targetVolumeSegmentReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = targetVolumeSegmentResp
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
