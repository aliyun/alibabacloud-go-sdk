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

type ScanImageRequest struct {
	Scene []*string               `json:"Scene,omitempty" xml:"Scene,omitempty" type:"Repeated"`
	Task  []*ScanImageRequestTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s ScanImageRequest) String() string {
	return tea.Prettify(s)
}

func (s ScanImageRequest) GoString() string {
	return s.String()
}

func (s *ScanImageRequest) SetScene(v []*string) *ScanImageRequest {
	s.Scene = v
	return s
}

func (s *ScanImageRequest) SetTask(v []*ScanImageRequestTask) *ScanImageRequest {
	s.Task = v
	return s
}

type ScanImageRequestTask struct {
	DataId               *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	ImageTimeMillisecond *int64  `json:"ImageTimeMillisecond,omitempty" xml:"ImageTimeMillisecond,omitempty"`
	ImageURL             *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Interval             *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	MaxFrames            *int32  `json:"MaxFrames,omitempty" xml:"MaxFrames,omitempty"`
}

func (s ScanImageRequestTask) String() string {
	return tea.Prettify(s)
}

func (s ScanImageRequestTask) GoString() string {
	return s.String()
}

func (s *ScanImageRequestTask) SetDataId(v string) *ScanImageRequestTask {
	s.DataId = &v
	return s
}

func (s *ScanImageRequestTask) SetImageTimeMillisecond(v int64) *ScanImageRequestTask {
	s.ImageTimeMillisecond = &v
	return s
}

func (s *ScanImageRequestTask) SetImageURL(v string) *ScanImageRequestTask {
	s.ImageURL = &v
	return s
}

func (s *ScanImageRequestTask) SetInterval(v int32) *ScanImageRequestTask {
	s.Interval = &v
	return s
}

func (s *ScanImageRequestTask) SetMaxFrames(v int32) *ScanImageRequestTask {
	s.MaxFrames = &v
	return s
}

type ScanImageAdvanceRequest struct {
	Scene []*string                      `json:"Scene,omitempty" xml:"Scene,omitempty" type:"Repeated"`
	Task  []*ScanImageAdvanceRequestTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s ScanImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ScanImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ScanImageAdvanceRequest) SetScene(v []*string) *ScanImageAdvanceRequest {
	s.Scene = v
	return s
}

func (s *ScanImageAdvanceRequest) SetTask(v []*ScanImageAdvanceRequestTask) *ScanImageAdvanceRequest {
	s.Task = v
	return s
}

type ScanImageAdvanceRequestTask struct {
	DataId               *string   `json:"DataId,omitempty" xml:"DataId,omitempty"`
	ImageTimeMillisecond *int64    `json:"ImageTimeMillisecond,omitempty" xml:"ImageTimeMillisecond,omitempty"`
	ImageURLObject       io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Interval             *int32    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	MaxFrames            *int32    `json:"MaxFrames,omitempty" xml:"MaxFrames,omitempty"`
}

func (s ScanImageAdvanceRequestTask) String() string {
	return tea.Prettify(s)
}

func (s ScanImageAdvanceRequestTask) GoString() string {
	return s.String()
}

func (s *ScanImageAdvanceRequestTask) SetDataId(v string) *ScanImageAdvanceRequestTask {
	s.DataId = &v
	return s
}

func (s *ScanImageAdvanceRequestTask) SetImageTimeMillisecond(v int64) *ScanImageAdvanceRequestTask {
	s.ImageTimeMillisecond = &v
	return s
}

func (s *ScanImageAdvanceRequestTask) SetImageURLObject(v io.Reader) *ScanImageAdvanceRequestTask {
	s.ImageURLObject = v
	return s
}

func (s *ScanImageAdvanceRequestTask) SetInterval(v int32) *ScanImageAdvanceRequestTask {
	s.Interval = &v
	return s
}

func (s *ScanImageAdvanceRequestTask) SetMaxFrames(v int32) *ScanImageAdvanceRequestTask {
	s.MaxFrames = &v
	return s
}

type ScanImageResponseBody struct {
	Data      *ScanImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScanImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBody) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBody) SetData(v *ScanImageResponseBodyData) *ScanImageResponseBody {
	s.Data = v
	return s
}

func (s *ScanImageResponseBody) SetRequestId(v string) *ScanImageResponseBody {
	s.RequestId = &v
	return s
}

type ScanImageResponseBodyData struct {
	Results []*ScanImageResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s ScanImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyData) SetResults(v []*ScanImageResponseBodyDataResults) *ScanImageResponseBodyData {
	s.Results = v
	return s
}

type ScanImageResponseBodyDataResults struct {
	DataId     *string                                       `json:"DataId,omitempty" xml:"DataId,omitempty"`
	ImageURL   *string                                       `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	SubResults []*ScanImageResponseBodyDataResultsSubResults `json:"SubResults,omitempty" xml:"SubResults,omitempty" type:"Repeated"`
	TaskId     *string                                       `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ScanImageResponseBodyDataResults) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResults) SetDataId(v string) *ScanImageResponseBodyDataResults {
	s.DataId = &v
	return s
}

func (s *ScanImageResponseBodyDataResults) SetImageURL(v string) *ScanImageResponseBodyDataResults {
	s.ImageURL = &v
	return s
}

func (s *ScanImageResponseBodyDataResults) SetSubResults(v []*ScanImageResponseBodyDataResultsSubResults) *ScanImageResponseBodyDataResults {
	s.SubResults = v
	return s
}

func (s *ScanImageResponseBodyDataResults) SetTaskId(v string) *ScanImageResponseBodyDataResults {
	s.TaskId = &v
	return s
}

type ScanImageResponseBodyDataResultsSubResults struct {
	Frames              []*ScanImageResponseBodyDataResultsSubResultsFrames              `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
	HintWordsInfoList   []*ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList   `json:"HintWordsInfoList,omitempty" xml:"HintWordsInfoList,omitempty" type:"Repeated"`
	Label               *string                                                          `json:"Label,omitempty" xml:"Label,omitempty"`
	LogoDataList        []*ScanImageResponseBodyDataResultsSubResultsLogoDataList        `json:"LogoDataList,omitempty" xml:"LogoDataList,omitempty" type:"Repeated"`
	OCRDataList         []*string                                                        `json:"OCRDataList,omitempty" xml:"OCRDataList,omitempty" type:"Repeated"`
	ProgramCodeDataList []*ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList `json:"ProgramCodeDataList,omitempty" xml:"ProgramCodeDataList,omitempty" type:"Repeated"`
	Rate                *float32                                                         `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Scene               *string                                                          `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SfaceDataList       []*ScanImageResponseBodyDataResultsSubResultsSfaceDataList       `json:"SfaceDataList,omitempty" xml:"SfaceDataList,omitempty" type:"Repeated"`
	Suggestion          *string                                                          `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResults) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResults) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetFrames(v []*ScanImageResponseBodyDataResultsSubResultsFrames) *ScanImageResponseBodyDataResultsSubResults {
	s.Frames = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetHintWordsInfoList(v []*ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList) *ScanImageResponseBodyDataResultsSubResults {
	s.HintWordsInfoList = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetLabel(v string) *ScanImageResponseBodyDataResultsSubResults {
	s.Label = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetLogoDataList(v []*ScanImageResponseBodyDataResultsSubResultsLogoDataList) *ScanImageResponseBodyDataResultsSubResults {
	s.LogoDataList = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetOCRDataList(v []*string) *ScanImageResponseBodyDataResultsSubResults {
	s.OCRDataList = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetProgramCodeDataList(v []*ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) *ScanImageResponseBodyDataResultsSubResults {
	s.ProgramCodeDataList = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetRate(v float32) *ScanImageResponseBodyDataResultsSubResults {
	s.Rate = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetScene(v string) *ScanImageResponseBodyDataResultsSubResults {
	s.Scene = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetSfaceDataList(v []*ScanImageResponseBodyDataResultsSubResultsSfaceDataList) *ScanImageResponseBodyDataResultsSubResults {
	s.SfaceDataList = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetSuggestion(v string) *ScanImageResponseBodyDataResultsSubResults {
	s.Suggestion = &v
	return s
}

type ScanImageResponseBodyDataResultsSubResultsFrames struct {
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	URL  *string  `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsFrames) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsFrames) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsFrames) SetRate(v float32) *ScanImageResponseBodyDataResultsSubResultsFrames {
	s.Rate = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsFrames) SetURL(v string) *ScanImageResponseBodyDataResultsSubResultsFrames {
	s.URL = &v
	return s
}

type ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList struct {
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList) SetContext(v string) *ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList {
	s.Context = &v
	return s
}

type ScanImageResponseBodyDataResultsSubResultsLogoDataList struct {
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Type   *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Width  *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y      *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsLogoDataList) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsLogoDataList) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetHeight(v float32) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.Height = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetName(v string) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.Name = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetType(v string) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.Type = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetWidth(v float32) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.Width = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetX(v float32) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.X = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetY(v float32) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.Y = &v
	return s
}

type ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList struct {
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y      *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) SetHeight(v float32) *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList {
	s.Height = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) SetWidth(v float32) *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList {
	s.Width = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) SetX(v float32) *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList {
	s.X = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) SetY(v float32) *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList {
	s.Y = &v
	return s
}

type ScanImageResponseBodyDataResultsSubResultsSfaceDataList struct {
	Faces  []*ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	Height *float32                                                        `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *float32                                                        `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *float32                                                        `json:"X,omitempty" xml:"X,omitempty"`
	Y      *float32                                                        `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsSfaceDataList) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsSfaceDataList) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) SetFaces(v []*ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) *ScanImageResponseBodyDataResultsSubResultsSfaceDataList {
	s.Faces = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) SetHeight(v float32) *ScanImageResponseBodyDataResultsSubResultsSfaceDataList {
	s.Height = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) SetWidth(v float32) *ScanImageResponseBodyDataResultsSubResultsSfaceDataList {
	s.Width = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) SetX(v float32) *ScanImageResponseBodyDataResultsSubResultsSfaceDataList {
	s.X = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) SetY(v float32) *ScanImageResponseBodyDataResultsSubResultsSfaceDataList {
	s.Y = &v
	return s
}

type ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces struct {
	Id   *string  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) SetId(v string) *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces {
	s.Id = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) SetName(v string) *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces {
	s.Name = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) SetRate(v float32) *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces {
	s.Rate = &v
	return s
}

type ScanImageResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ScanImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScanImageResponse) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponse) GoString() string {
	return s.String()
}

func (s *ScanImageResponse) SetHeaders(v map[string]*string) *ScanImageResponse {
	s.Headers = v
	return s
}

func (s *ScanImageResponse) SetStatusCode(v int32) *ScanImageResponse {
	s.StatusCode = &v
	return s
}

func (s *ScanImageResponse) SetBody(v *ScanImageResponseBody) *ScanImageResponse {
	s.Body = v
	return s
}

type ScanTextRequest struct {
	Labels []*ScanTextRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Tasks  []*ScanTextRequestTasks  `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ScanTextRequest) String() string {
	return tea.Prettify(s)
}

func (s ScanTextRequest) GoString() string {
	return s.String()
}

func (s *ScanTextRequest) SetLabels(v []*ScanTextRequestLabels) *ScanTextRequest {
	s.Labels = v
	return s
}

func (s *ScanTextRequest) SetTasks(v []*ScanTextRequestTasks) *ScanTextRequest {
	s.Tasks = v
	return s
}

type ScanTextRequestLabels struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ScanTextRequestLabels) String() string {
	return tea.Prettify(s)
}

func (s ScanTextRequestLabels) GoString() string {
	return s.String()
}

func (s *ScanTextRequestLabels) SetLabel(v string) *ScanTextRequestLabels {
	s.Label = &v
	return s
}

type ScanTextRequestTasks struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ScanTextRequestTasks) String() string {
	return tea.Prettify(s)
}

func (s ScanTextRequestTasks) GoString() string {
	return s.String()
}

func (s *ScanTextRequestTasks) SetContent(v string) *ScanTextRequestTasks {
	s.Content = &v
	return s
}

type ScanTextResponseBody struct {
	Data      *ScanTextResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScanTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponseBody) GoString() string {
	return s.String()
}

func (s *ScanTextResponseBody) SetData(v *ScanTextResponseBodyData) *ScanTextResponseBody {
	s.Data = v
	return s
}

func (s *ScanTextResponseBody) SetRequestId(v string) *ScanTextResponseBody {
	s.RequestId = &v
	return s
}

type ScanTextResponseBodyData struct {
	Elements []*ScanTextResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s ScanTextResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponseBodyData) GoString() string {
	return s.String()
}

func (s *ScanTextResponseBodyData) SetElements(v []*ScanTextResponseBodyDataElements) *ScanTextResponseBodyData {
	s.Elements = v
	return s
}

type ScanTextResponseBodyDataElements struct {
	Results []*ScanTextResponseBodyDataElementsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	TaskId  *string                                    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ScanTextResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *ScanTextResponseBodyDataElements) SetResults(v []*ScanTextResponseBodyDataElementsResults) *ScanTextResponseBodyDataElements {
	s.Results = v
	return s
}

func (s *ScanTextResponseBodyDataElements) SetTaskId(v string) *ScanTextResponseBodyDataElements {
	s.TaskId = &v
	return s
}

type ScanTextResponseBodyDataElementsResults struct {
	Details    []*ScanTextResponseBodyDataElementsResultsDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	Label      *string                                           `json:"Label,omitempty" xml:"Label,omitempty"`
	Rate       *float32                                          `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Suggestion *string                                           `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s ScanTextResponseBodyDataElementsResults) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponseBodyDataElementsResults) GoString() string {
	return s.String()
}

func (s *ScanTextResponseBodyDataElementsResults) SetDetails(v []*ScanTextResponseBodyDataElementsResultsDetails) *ScanTextResponseBodyDataElementsResults {
	s.Details = v
	return s
}

func (s *ScanTextResponseBodyDataElementsResults) SetLabel(v string) *ScanTextResponseBodyDataElementsResults {
	s.Label = &v
	return s
}

func (s *ScanTextResponseBodyDataElementsResults) SetRate(v float32) *ScanTextResponseBodyDataElementsResults {
	s.Rate = &v
	return s
}

func (s *ScanTextResponseBodyDataElementsResults) SetSuggestion(v string) *ScanTextResponseBodyDataElementsResults {
	s.Suggestion = &v
	return s
}

type ScanTextResponseBodyDataElementsResultsDetails struct {
	Contexts []*ScanTextResponseBodyDataElementsResultsDetailsContexts `json:"Contexts,omitempty" xml:"Contexts,omitempty" type:"Repeated"`
	Label    *string                                                   `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ScanTextResponseBodyDataElementsResultsDetails) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponseBodyDataElementsResultsDetails) GoString() string {
	return s.String()
}

func (s *ScanTextResponseBodyDataElementsResultsDetails) SetContexts(v []*ScanTextResponseBodyDataElementsResultsDetailsContexts) *ScanTextResponseBodyDataElementsResultsDetails {
	s.Contexts = v
	return s
}

func (s *ScanTextResponseBodyDataElementsResultsDetails) SetLabel(v string) *ScanTextResponseBodyDataElementsResultsDetails {
	s.Label = &v
	return s
}

type ScanTextResponseBodyDataElementsResultsDetailsContexts struct {
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
}

func (s ScanTextResponseBodyDataElementsResultsDetailsContexts) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponseBodyDataElementsResultsDetailsContexts) GoString() string {
	return s.String()
}

func (s *ScanTextResponseBodyDataElementsResultsDetailsContexts) SetContext(v string) *ScanTextResponseBodyDataElementsResultsDetailsContexts {
	s.Context = &v
	return s
}

type ScanTextResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ScanTextResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScanTextResponse) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponse) GoString() string {
	return s.String()
}

func (s *ScanTextResponse) SetHeaders(v map[string]*string) *ScanTextResponse {
	s.Headers = v
	return s
}

func (s *ScanTextResponse) SetStatusCode(v int32) *ScanTextResponse {
	s.StatusCode = &v
	return s
}

func (s *ScanTextResponse) SetBody(v *ScanTextResponseBody) *ScanTextResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("imageaudit"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ScanImageWithOptions(request *ScanImageRequest, runtime *util.RuntimeOptions) (_result *ScanImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.Task)) {
		body["Task"] = request.Task
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ScanImage"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ScanImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScanImage(request *ScanImageRequest) (_result *ScanImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScanImageResponse{}
	_body, _err := client.ScanImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScanImageAdvance(request *ScanImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *ScanImageResponse, _err error) {
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
		Product:  tea.String("imageaudit"),
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
	scanImageReq := &ScanImageRequest{}
	openapiutil.Convert(request, scanImageReq)
	if !tea.BoolValue(util.IsUnset(request.Task)) {
		i := tea.Int(0)
		for _, item0 := range request.Task {
			if !tea.BoolValue(util.IsUnset(item0.ImageURLObject)) {
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
					Content:     item0.ImageURLObject,
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
				tmp := scanImageReq.Task[tea.IntValue(i)]
				tmp.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
				i = number.Ltoi(number.Add(number.Itol(i), number.Itol(tea.Int(1))))
			}

		}
	}

	scanImageResp, _err := client.ScanImageWithOptions(scanImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = scanImageResp
	return _result, _err
}

func (client *Client) ScanTextWithOptions(request *ScanTextRequest, runtime *util.RuntimeOptions) (_result *ScanTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Tasks)) {
		body["Tasks"] = request.Tasks
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ScanText"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ScanTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScanText(request *ScanTextRequest) (_result *ScanTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScanTextResponse{}
	_body, _err := client.ScanTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
