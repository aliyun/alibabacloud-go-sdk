// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ScanTextRequest struct {
	Tasks  []*ScanTextRequestTasks  `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	Labels []*ScanTextRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s ScanTextRequest) String() string {
	return tea.Prettify(s)
}

func (s ScanTextRequest) GoString() string {
	return s.String()
}

func (s *ScanTextRequest) SetTasks(v []*ScanTextRequestTasks) *ScanTextRequest {
	s.Tasks = v
	return s
}

func (s *ScanTextRequest) SetLabels(v []*ScanTextRequestLabels) *ScanTextRequest {
	s.Labels = v
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

type ScanTextResponseBody struct {
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ScanTextResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ScanTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponseBody) GoString() string {
	return s.String()
}

func (s *ScanTextResponseBody) SetRequestId(v string) *ScanTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScanTextResponseBody) SetData(v *ScanTextResponseBodyData) *ScanTextResponseBody {
	s.Data = v
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
	TaskId  *string                                    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Results []*ScanTextResponseBodyDataElementsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s ScanTextResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *ScanTextResponseBodyDataElements) SetTaskId(v string) *ScanTextResponseBodyDataElements {
	s.TaskId = &v
	return s
}

func (s *ScanTextResponseBodyDataElements) SetResults(v []*ScanTextResponseBodyDataElementsResults) *ScanTextResponseBodyDataElements {
	s.Results = v
	return s
}

type ScanTextResponseBodyDataElementsResults struct {
	Suggestion *string                                           `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	Label      *string                                           `json:"Label,omitempty" xml:"Label,omitempty"`
	Rate       *float32                                          `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Details    []*ScanTextResponseBodyDataElementsResultsDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
}

func (s ScanTextResponseBodyDataElementsResults) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponseBodyDataElementsResults) GoString() string {
	return s.String()
}

func (s *ScanTextResponseBodyDataElementsResults) SetSuggestion(v string) *ScanTextResponseBodyDataElementsResults {
	s.Suggestion = &v
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

func (s *ScanTextResponseBodyDataElementsResults) SetDetails(v []*ScanTextResponseBodyDataElementsResultsDetails) *ScanTextResponseBodyDataElementsResults {
	s.Details = v
	return s
}

type ScanTextResponseBodyDataElementsResultsDetails struct {
	Label    *string                                                   `json:"Label,omitempty" xml:"Label,omitempty"`
	Contexts []*ScanTextResponseBodyDataElementsResultsDetailsContexts `json:"Contexts,omitempty" xml:"Contexts,omitempty" type:"Repeated"`
}

func (s ScanTextResponseBodyDataElementsResultsDetails) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponseBodyDataElementsResultsDetails) GoString() string {
	return s.String()
}

func (s *ScanTextResponseBodyDataElementsResultsDetails) SetLabel(v string) *ScanTextResponseBodyDataElementsResultsDetails {
	s.Label = &v
	return s
}

func (s *ScanTextResponseBodyDataElementsResultsDetails) SetContexts(v []*ScanTextResponseBodyDataElementsResultsDetailsContexts) *ScanTextResponseBodyDataElementsResultsDetails {
	s.Contexts = v
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
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ScanTextResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ScanTextResponse) SetBody(v *ScanTextResponseBody) *ScanTextResponse {
	s.Body = v
	return s
}

type ScanImageRequest struct {
	Task  []*ScanImageRequestTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
	Scene []*string               `json:"Scene,omitempty" xml:"Scene,omitempty" type:"Repeated"`
}

func (s ScanImageRequest) String() string {
	return tea.Prettify(s)
}

func (s ScanImageRequest) GoString() string {
	return s.String()
}

func (s *ScanImageRequest) SetTask(v []*ScanImageRequestTask) *ScanImageRequest {
	s.Task = v
	return s
}

func (s *ScanImageRequest) SetScene(v []*string) *ScanImageRequest {
	s.Scene = v
	return s
}

type ScanImageRequestTask struct {
	ImageTimeMillisecond *int64  `json:"ImageTimeMillisecond,omitempty" xml:"ImageTimeMillisecond,omitempty"`
	Interval             *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ImageURL             *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	MaxFrames            *int32  `json:"MaxFrames,omitempty" xml:"MaxFrames,omitempty"`
	DataId               *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
}

func (s ScanImageRequestTask) String() string {
	return tea.Prettify(s)
}

func (s ScanImageRequestTask) GoString() string {
	return s.String()
}

func (s *ScanImageRequestTask) SetImageTimeMillisecond(v int64) *ScanImageRequestTask {
	s.ImageTimeMillisecond = &v
	return s
}

func (s *ScanImageRequestTask) SetInterval(v int32) *ScanImageRequestTask {
	s.Interval = &v
	return s
}

func (s *ScanImageRequestTask) SetImageURL(v string) *ScanImageRequestTask {
	s.ImageURL = &v
	return s
}

func (s *ScanImageRequestTask) SetMaxFrames(v int32) *ScanImageRequestTask {
	s.MaxFrames = &v
	return s
}

func (s *ScanImageRequestTask) SetDataId(v string) *ScanImageRequestTask {
	s.DataId = &v
	return s
}

type ScanImageResponseBody struct {
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ScanImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ScanImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBody) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBody) SetRequestId(v string) *ScanImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScanImageResponseBody) SetData(v *ScanImageResponseBodyData) *ScanImageResponseBody {
	s.Data = v
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
	ImageURL   *string                                       `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	DataId     *string                                       `json:"DataId,omitempty" xml:"DataId,omitempty"`
	SubResults []*ScanImageResponseBodyDataResultsSubResults `json:"SubResults,omitempty" xml:"SubResults,omitempty" type:"Repeated"`
	TaskId     *string                                       `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ScanImageResponseBodyDataResults) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResults) SetImageURL(v string) *ScanImageResponseBodyDataResults {
	s.ImageURL = &v
	return s
}

func (s *ScanImageResponseBodyDataResults) SetDataId(v string) *ScanImageResponseBodyDataResults {
	s.DataId = &v
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
	SfaceDataList       []*ScanImageResponseBodyDataResultsSubResultsSfaceDataList       `json:"SfaceDataList,omitempty" xml:"SfaceDataList,omitempty" type:"Repeated"`
	HintWordsInfoList   []*ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList   `json:"HintWordsInfoList,omitempty" xml:"HintWordsInfoList,omitempty" type:"Repeated"`
	Suggestion          *string                                                          `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	ProgramCodeDataList []*ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList `json:"ProgramCodeDataList,omitempty" xml:"ProgramCodeDataList,omitempty" type:"Repeated"`
	OCRDataList         []*string                                                        `json:"OCRDataList,omitempty" xml:"OCRDataList,omitempty" type:"Repeated"`
	Frames              []*ScanImageResponseBodyDataResultsSubResultsFrames              `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
	LogoDataList        []*ScanImageResponseBodyDataResultsSubResultsLogoDataList        `json:"LogoDataList,omitempty" xml:"LogoDataList,omitempty" type:"Repeated"`
	Label               *string                                                          `json:"Label,omitempty" xml:"Label,omitempty"`
	Scene               *string                                                          `json:"Scene,omitempty" xml:"Scene,omitempty"`
	Rate                *float32                                                         `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResults) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResults) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetSfaceDataList(v []*ScanImageResponseBodyDataResultsSubResultsSfaceDataList) *ScanImageResponseBodyDataResultsSubResults {
	s.SfaceDataList = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetHintWordsInfoList(v []*ScanImageResponseBodyDataResultsSubResultsHintWordsInfoList) *ScanImageResponseBodyDataResultsSubResults {
	s.HintWordsInfoList = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetSuggestion(v string) *ScanImageResponseBodyDataResultsSubResults {
	s.Suggestion = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetProgramCodeDataList(v []*ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) *ScanImageResponseBodyDataResultsSubResults {
	s.ProgramCodeDataList = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetOCRDataList(v []*string) *ScanImageResponseBodyDataResultsSubResults {
	s.OCRDataList = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetFrames(v []*ScanImageResponseBodyDataResultsSubResultsFrames) *ScanImageResponseBodyDataResultsSubResults {
	s.Frames = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetLogoDataList(v []*ScanImageResponseBodyDataResultsSubResultsLogoDataList) *ScanImageResponseBodyDataResultsSubResults {
	s.LogoDataList = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetLabel(v string) *ScanImageResponseBodyDataResultsSubResults {
	s.Label = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetScene(v string) *ScanImageResponseBodyDataResultsSubResults {
	s.Scene = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResults) SetRate(v float32) *ScanImageResponseBodyDataResultsSubResults {
	s.Rate = &v
	return s
}

type ScanImageResponseBodyDataResultsSubResultsSfaceDataList struct {
	Width  *float32                                                        `json:"Width,omitempty" xml:"Width,omitempty"`
	Faces  []*ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	Height *float32                                                        `json:"Height,omitempty" xml:"Height,omitempty"`
	Y      *float32                                                        `json:"Y,omitempty" xml:"Y,omitempty"`
	X      *float32                                                        `json:"X,omitempty" xml:"X,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsSfaceDataList) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsSfaceDataList) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) SetWidth(v float32) *ScanImageResponseBodyDataResultsSubResultsSfaceDataList {
	s.Width = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) SetFaces(v []*ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) *ScanImageResponseBodyDataResultsSubResultsSfaceDataList {
	s.Faces = v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) SetHeight(v float32) *ScanImageResponseBodyDataResultsSubResultsSfaceDataList {
	s.Height = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) SetY(v float32) *ScanImageResponseBodyDataResultsSubResultsSfaceDataList {
	s.Y = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataList) SetX(v float32) *ScanImageResponseBodyDataResultsSubResultsSfaceDataList {
	s.X = &v
	return s
}

type ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces struct {
	Name *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *string  `json:"Id,omitempty" xml:"Id,omitempty"`
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) SetName(v string) *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces {
	s.Name = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) SetId(v string) *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces {
	s.Id = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces) SetRate(v float32) *ScanImageResponseBodyDataResultsSubResultsSfaceDataListFaces {
	s.Rate = &v
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

type ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList struct {
	Width  *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y      *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	X      *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) SetWidth(v float32) *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList {
	s.Width = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) SetHeight(v float32) *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList {
	s.Height = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) SetY(v float32) *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList {
	s.Y = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList) SetX(v float32) *ScanImageResponseBodyDataResultsSubResultsProgramCodeDataList {
	s.X = &v
	return s
}

type ScanImageResponseBodyDataResultsSubResultsFrames struct {
	URL  *string  `json:"URL,omitempty" xml:"URL,omitempty"`
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsFrames) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsFrames) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsFrames) SetURL(v string) *ScanImageResponseBodyDataResultsSubResultsFrames {
	s.URL = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsFrames) SetRate(v float32) *ScanImageResponseBodyDataResultsSubResultsFrames {
	s.Rate = &v
	return s
}

type ScanImageResponseBodyDataResultsSubResultsLogoDataList struct {
	Type   *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Width  *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y      *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	X      *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s ScanImageResponseBodyDataResultsSubResultsLogoDataList) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseBodyDataResultsSubResultsLogoDataList) GoString() string {
	return s.String()
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetType(v string) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.Type = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetWidth(v float32) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.Width = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetHeight(v float32) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.Height = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetY(v float32) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.Y = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetName(v string) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.Name = &v
	return s
}

func (s *ScanImageResponseBodyDataResultsSubResultsLogoDataList) SetX(v float32) *ScanImageResponseBodyDataResultsSubResultsLogoDataList {
	s.X = &v
	return s
}

type ScanImageResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ScanImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ScanImageResponse) SetBody(v *ScanImageResponseBody) *ScanImageResponse {
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

func (client *Client) ScanTextWithOptions(request *ScanTextRequest, runtime *util.RuntimeOptions) (_result *ScanTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ScanTextResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ScanText"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ScanImageWithOptions(request *ScanImageRequest, runtime *util.RuntimeOptions) (_result *ScanImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ScanImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ScanImage"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
