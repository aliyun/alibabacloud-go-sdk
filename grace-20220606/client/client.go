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

type AnalysisConfig struct {
	BadThroughputThreshold           *float64                 `json:"badThroughputThreshold,omitempty" xml:"badThroughputThreshold,omitempty"`
	FullGCFrequentIntervalThreshold  *float64                 `json:"fullGCFrequentIntervalThreshold,omitempty" xml:"fullGCFrequentIntervalThreshold,omitempty"`
	HighHeapUsageThreshold           *float64                 `json:"highHeapUsageThreshold,omitempty" xml:"highHeapUsageThreshold,omitempty"`
	HighHumongousUsageThreshold      *float64                 `json:"highHumongousUsageThreshold,omitempty" xml:"highHumongousUsageThreshold,omitempty"`
	HighMetaspaceUsageThreshold      *float64                 `json:"highMetaspaceUsageThreshold,omitempty" xml:"highMetaspaceUsageThreshold,omitempty"`
	HighOldUsageThreshold            *float64                 `json:"highOldUsageThreshold,omitempty" xml:"highOldUsageThreshold,omitempty"`
	HighPromotionThreshold           *float64                 `json:"highPromotionThreshold,omitempty" xml:"highPromotionThreshold,omitempty"`
	HighSysThreshold                 *float64                 `json:"highSysThreshold,omitempty" xml:"highSysThreshold,omitempty"`
	LongConcurrentThreshold          *float64                 `json:"longConcurrentThreshold,omitempty" xml:"longConcurrentThreshold,omitempty"`
	LongPauseThreshold               *float64                 `json:"longPauseThreshold,omitempty" xml:"longPauseThreshold,omitempty"`
	LowUsrThreshold                  *float64                 `json:"lowUsrThreshold,omitempty" xml:"lowUsrThreshold,omitempty"`
	OldGCFrequentIntervalThreshold   *float64                 `json:"oldGCFrequentIntervalThreshold,omitempty" xml:"oldGCFrequentIntervalThreshold,omitempty"`
	SmallGenerationThreshold         *float64                 `json:"smallGenerationThreshold,omitempty" xml:"smallGenerationThreshold,omitempty"`
	TimeRange                        *AnalysisConfigTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
	TooManyOldGCThreshold            *float64                 `json:"tooManyOldGCThreshold,omitempty" xml:"tooManyOldGCThreshold,omitempty"`
	YoungGCFrequentIntervalThreshold *float64                 `json:"youngGCFrequentIntervalThreshold,omitempty" xml:"youngGCFrequentIntervalThreshold,omitempty"`
}

func (s AnalysisConfig) String() string {
	return tea.Prettify(s)
}

func (s AnalysisConfig) GoString() string {
	return s.String()
}

func (s *AnalysisConfig) SetBadThroughputThreshold(v float64) *AnalysisConfig {
	s.BadThroughputThreshold = &v
	return s
}

func (s *AnalysisConfig) SetFullGCFrequentIntervalThreshold(v float64) *AnalysisConfig {
	s.FullGCFrequentIntervalThreshold = &v
	return s
}

func (s *AnalysisConfig) SetHighHeapUsageThreshold(v float64) *AnalysisConfig {
	s.HighHeapUsageThreshold = &v
	return s
}

func (s *AnalysisConfig) SetHighHumongousUsageThreshold(v float64) *AnalysisConfig {
	s.HighHumongousUsageThreshold = &v
	return s
}

func (s *AnalysisConfig) SetHighMetaspaceUsageThreshold(v float64) *AnalysisConfig {
	s.HighMetaspaceUsageThreshold = &v
	return s
}

func (s *AnalysisConfig) SetHighOldUsageThreshold(v float64) *AnalysisConfig {
	s.HighOldUsageThreshold = &v
	return s
}

func (s *AnalysisConfig) SetHighPromotionThreshold(v float64) *AnalysisConfig {
	s.HighPromotionThreshold = &v
	return s
}

func (s *AnalysisConfig) SetHighSysThreshold(v float64) *AnalysisConfig {
	s.HighSysThreshold = &v
	return s
}

func (s *AnalysisConfig) SetLongConcurrentThreshold(v float64) *AnalysisConfig {
	s.LongConcurrentThreshold = &v
	return s
}

func (s *AnalysisConfig) SetLongPauseThreshold(v float64) *AnalysisConfig {
	s.LongPauseThreshold = &v
	return s
}

func (s *AnalysisConfig) SetLowUsrThreshold(v float64) *AnalysisConfig {
	s.LowUsrThreshold = &v
	return s
}

func (s *AnalysisConfig) SetOldGCFrequentIntervalThreshold(v float64) *AnalysisConfig {
	s.OldGCFrequentIntervalThreshold = &v
	return s
}

func (s *AnalysisConfig) SetSmallGenerationThreshold(v float64) *AnalysisConfig {
	s.SmallGenerationThreshold = &v
	return s
}

func (s *AnalysisConfig) SetTimeRange(v *AnalysisConfigTimeRange) *AnalysisConfig {
	s.TimeRange = v
	return s
}

func (s *AnalysisConfig) SetTooManyOldGCThreshold(v float64) *AnalysisConfig {
	s.TooManyOldGCThreshold = &v
	return s
}

func (s *AnalysisConfig) SetYoungGCFrequentIntervalThreshold(v float64) *AnalysisConfig {
	s.YoungGCFrequentIntervalThreshold = &v
	return s
}

type AnalysisConfigTimeRange struct {
	End   *float64 `json:"end,omitempty" xml:"end,omitempty"`
	Start *float64 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s AnalysisConfigTimeRange) String() string {
	return tea.Prettify(s)
}

func (s AnalysisConfigTimeRange) GoString() string {
	return s.String()
}

func (s *AnalysisConfigTimeRange) SetEnd(v float64) *AnalysisConfigTimeRange {
	s.End = &v
	return s
}

func (s *AnalysisConfigTimeRange) SetStart(v float64) *AnalysisConfigTimeRange {
	s.Start = &v
	return s
}

type FileInfo struct {
	AnalyzeProgress  *FileInfoAnalyzeProgress  `json:"analyzeProgress,omitempty" xml:"analyzeProgress,omitempty" type:"Struct"`
	CreationTime     *int64                    `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
	DisplayName      *string                   `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Labels           *string                   `json:"labels,omitempty" xml:"labels,omitempty"`
	Name             *string                   `json:"name,omitempty" xml:"name,omitempty"`
	RequestId        *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Shared           *bool                     `json:"shared,omitempty" xml:"shared,omitempty"`
	Size             *int64                    `json:"size,omitempty" xml:"size,omitempty"`
	TransferProgress *FileInfoTransferProgress `json:"transferProgress,omitempty" xml:"transferProgress,omitempty" type:"Struct"`
	TransferState    *string                   `json:"transferState,omitempty" xml:"transferState,omitempty"`
	Type             *string                   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FileInfo) String() string {
	return tea.Prettify(s)
}

func (s FileInfo) GoString() string {
	return s.String()
}

func (s *FileInfo) SetAnalyzeProgress(v *FileInfoAnalyzeProgress) *FileInfo {
	s.AnalyzeProgress = v
	return s
}

func (s *FileInfo) SetCreationTime(v int64) *FileInfo {
	s.CreationTime = &v
	return s
}

func (s *FileInfo) SetDisplayName(v string) *FileInfo {
	s.DisplayName = &v
	return s
}

func (s *FileInfo) SetLabels(v string) *FileInfo {
	s.Labels = &v
	return s
}

func (s *FileInfo) SetName(v string) *FileInfo {
	s.Name = &v
	return s
}

func (s *FileInfo) SetRequestId(v string) *FileInfo {
	s.RequestId = &v
	return s
}

func (s *FileInfo) SetShared(v bool) *FileInfo {
	s.Shared = &v
	return s
}

func (s *FileInfo) SetSize(v int64) *FileInfo {
	s.Size = &v
	return s
}

func (s *FileInfo) SetTransferProgress(v *FileInfoTransferProgress) *FileInfo {
	s.TransferProgress = v
	return s
}

func (s *FileInfo) SetTransferState(v string) *FileInfo {
	s.TransferState = &v
	return s
}

func (s *FileInfo) SetType(v string) *FileInfo {
	s.Type = &v
	return s
}

type FileInfoAnalyzeProgress struct {
	Message *string  `json:"message,omitempty" xml:"message,omitempty"`
	Percent *float64 `json:"percent,omitempty" xml:"percent,omitempty"`
	State   *string  `json:"state,omitempty" xml:"state,omitempty"`
}

func (s FileInfoAnalyzeProgress) String() string {
	return tea.Prettify(s)
}

func (s FileInfoAnalyzeProgress) GoString() string {
	return s.String()
}

func (s *FileInfoAnalyzeProgress) SetMessage(v string) *FileInfoAnalyzeProgress {
	s.Message = &v
	return s
}

func (s *FileInfoAnalyzeProgress) SetPercent(v float64) *FileInfoAnalyzeProgress {
	s.Percent = &v
	return s
}

func (s *FileInfoAnalyzeProgress) SetState(v string) *FileInfoAnalyzeProgress {
	s.State = &v
	return s
}

type FileInfoTransferProgress struct {
	TotalSize       *int64 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
	TransferredSize *int64 `json:"transferredSize,omitempty" xml:"transferredSize,omitempty"`
}

func (s FileInfoTransferProgress) String() string {
	return tea.Prettify(s)
}

func (s FileInfoTransferProgress) GoString() string {
	return s.String()
}

func (s *FileInfoTransferProgress) SetTotalSize(v int64) *FileInfoTransferProgress {
	s.TotalSize = &v
	return s
}

func (s *FileInfoTransferProgress) SetTransferredSize(v int64) *FileInfoTransferProgress {
	s.TransferredSize = &v
	return s
}

type PhaseStatisticItem struct {
	Count         *int32   `json:"count,omitempty" xml:"count,omitempty"`
	DurationAvg   *float64 `json:"durationAvg,omitempty" xml:"durationAvg,omitempty"`
	DurationMax   *float64 `json:"durationMax,omitempty" xml:"durationMax,omitempty"`
	DurationTotal *float64 `json:"durationTotal,omitempty" xml:"durationTotal,omitempty"`
	IntervalAvg   *float64 `json:"intervalAvg,omitempty" xml:"intervalAvg,omitempty"`
	IntervalMin   *float64 `json:"intervalMin,omitempty" xml:"intervalMin,omitempty"`
	Name          *string  `json:"name,omitempty" xml:"name,omitempty"`
}

func (s PhaseStatisticItem) String() string {
	return tea.Prettify(s)
}

func (s PhaseStatisticItem) GoString() string {
	return s.String()
}

func (s *PhaseStatisticItem) SetCount(v int32) *PhaseStatisticItem {
	s.Count = &v
	return s
}

func (s *PhaseStatisticItem) SetDurationAvg(v float64) *PhaseStatisticItem {
	s.DurationAvg = &v
	return s
}

func (s *PhaseStatisticItem) SetDurationMax(v float64) *PhaseStatisticItem {
	s.DurationMax = &v
	return s
}

func (s *PhaseStatisticItem) SetDurationTotal(v float64) *PhaseStatisticItem {
	s.DurationTotal = &v
	return s
}

func (s *PhaseStatisticItem) SetIntervalAvg(v float64) *PhaseStatisticItem {
	s.IntervalAvg = &v
	return s
}

func (s *PhaseStatisticItem) SetIntervalMin(v float64) *PhaseStatisticItem {
	s.IntervalMin = &v
	return s
}

func (s *PhaseStatisticItem) SetName(v string) *PhaseStatisticItem {
	s.Name = &v
	return s
}

type AnalyzeFileRequest struct {
	KeepUnreachableObjects *bool   `json:"keepUnreachableObjects,omitempty" xml:"keepUnreachableObjects,omitempty"`
	Name                   *string `json:"name,omitempty" xml:"name,omitempty"`
	Token                  *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s AnalyzeFileRequest) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeFileRequest) GoString() string {
	return s.String()
}

func (s *AnalyzeFileRequest) SetKeepUnreachableObjects(v bool) *AnalyzeFileRequest {
	s.KeepUnreachableObjects = &v
	return s
}

func (s *AnalyzeFileRequest) SetName(v string) *AnalyzeFileRequest {
	s.Name = &v
	return s
}

func (s *AnalyzeFileRequest) SetToken(v string) *AnalyzeFileRequest {
	s.Token = &v
	return s
}

type AnalyzeFileResponseBody struct {
	FileName  *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AnalyzeFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeFileResponseBody) GoString() string {
	return s.String()
}

func (s *AnalyzeFileResponseBody) SetFileName(v string) *AnalyzeFileResponseBody {
	s.FileName = &v
	return s
}

func (s *AnalyzeFileResponseBody) SetRequestId(v string) *AnalyzeFileResponseBody {
	s.RequestId = &v
	return s
}

type AnalyzeFileResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AnalyzeFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AnalyzeFileResponse) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeFileResponse) GoString() string {
	return s.String()
}

func (s *AnalyzeFileResponse) SetHeaders(v map[string]*string) *AnalyzeFileResponse {
	s.Headers = v
	return s
}

func (s *AnalyzeFileResponse) SetStatusCode(v int32) *AnalyzeFileResponse {
	s.StatusCode = &v
	return s
}

func (s *AnalyzeFileResponse) SetBody(v *AnalyzeFileResponseBody) *AnalyzeFileResponse {
	s.Body = v
	return s
}

type GetFileRequest struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GetFileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileRequest) GoString() string {
	return s.String()
}

func (s *GetFileRequest) SetName(v string) *GetFileRequest {
	s.Name = &v
	return s
}

func (s *GetFileRequest) SetToken(v string) *GetFileRequest {
	s.Token = &v
	return s
}

type GetFileResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FileInfo          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileResponse) GoString() string {
	return s.String()
}

func (s *GetFileResponse) SetHeaders(v map[string]*string) *GetFileResponse {
	s.Headers = v
	return s
}

func (s *GetFileResponse) SetStatusCode(v int32) *GetFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileResponse) SetBody(v *FileInfo) *GetFileResponse {
	s.Body = v
	return s
}

type UploadFileByOSSRequest struct {
	BucketName  *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// oss endpoint
	Endpoint   *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	ObjectName *string `json:"objectName,omitempty" xml:"objectName,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UploadFileByOSSRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadFileByOSSRequest) GoString() string {
	return s.String()
}

func (s *UploadFileByOSSRequest) SetBucketName(v string) *UploadFileByOSSRequest {
	s.BucketName = &v
	return s
}

func (s *UploadFileByOSSRequest) SetDisplayName(v string) *UploadFileByOSSRequest {
	s.DisplayName = &v
	return s
}

func (s *UploadFileByOSSRequest) SetEndpoint(v string) *UploadFileByOSSRequest {
	s.Endpoint = &v
	return s
}

func (s *UploadFileByOSSRequest) SetObjectName(v string) *UploadFileByOSSRequest {
	s.ObjectName = &v
	return s
}

func (s *UploadFileByOSSRequest) SetType(v string) *UploadFileByOSSRequest {
	s.Type = &v
	return s
}

type UploadFileByOSSResponseBody struct {
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UploadFileByOSSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadFileByOSSResponseBody) GoString() string {
	return s.String()
}

func (s *UploadFileByOSSResponseBody) SetName(v string) *UploadFileByOSSResponseBody {
	s.Name = &v
	return s
}

func (s *UploadFileByOSSResponseBody) SetRequestId(v string) *UploadFileByOSSResponseBody {
	s.RequestId = &v
	return s
}

type UploadFileByOSSResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadFileByOSSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadFileByOSSResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadFileByOSSResponse) GoString() string {
	return s.String()
}

func (s *UploadFileByOSSResponse) SetHeaders(v map[string]*string) *UploadFileByOSSResponse {
	s.Headers = v
	return s
}

func (s *UploadFileByOSSResponse) SetStatusCode(v int32) *UploadFileByOSSResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadFileByOSSResponse) SetBody(v *UploadFileByOSSResponseBody) *UploadFileByOSSResponse {
	s.Body = v
	return s
}

type UploadFileByURLRequest struct {
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Type        *string `json:"type,omitempty" xml:"type,omitempty"`
	Url         *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s UploadFileByURLRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadFileByURLRequest) GoString() string {
	return s.String()
}

func (s *UploadFileByURLRequest) SetDisplayName(v string) *UploadFileByURLRequest {
	s.DisplayName = &v
	return s
}

func (s *UploadFileByURLRequest) SetType(v string) *UploadFileByURLRequest {
	s.Type = &v
	return s
}

func (s *UploadFileByURLRequest) SetUrl(v string) *UploadFileByURLRequest {
	s.Url = &v
	return s
}

type UploadFileByURLResponseBody struct {
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UploadFileByURLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadFileByURLResponseBody) GoString() string {
	return s.String()
}

func (s *UploadFileByURLResponseBody) SetName(v string) *UploadFileByURLResponseBody {
	s.Name = &v
	return s
}

func (s *UploadFileByURLResponseBody) SetRequestId(v string) *UploadFileByURLResponseBody {
	s.RequestId = &v
	return s
}

type UploadFileByURLResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadFileByURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadFileByURLResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadFileByURLResponse) GoString() string {
	return s.String()
}

func (s *UploadFileByURLResponse) SetHeaders(v map[string]*string) *UploadFileByURLResponse {
	s.Headers = v
	return s
}

func (s *UploadFileByURLResponse) SetStatusCode(v int32) *UploadFileByURLResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadFileByURLResponse) SetBody(v *UploadFileByURLResponseBody) *UploadFileByURLResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("grace"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AnalyzeFileWithOptions(request *AnalyzeFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AnalyzeFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeepUnreachableObjects)) {
		query["keepUnreachableObjects"] = request.KeepUnreachableObjects
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AnalyzeFile"),
		Version:     tea.String("2022-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/AnalyzeFile"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AnalyzeFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AnalyzeFile(request *AnalyzeFileRequest) (_result *AnalyzeFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AnalyzeFileResponse{}
	_body, _err := client.AnalyzeFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFileWithOptions(request *GetFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFile"),
		Version:     tea.String("2022-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/GetFile"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFile(request *GetFileRequest) (_result *GetFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFileResponse{}
	_body, _err := client.GetFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadFileByOSSWithOptions(request *UploadFileByOSSRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadFileByOSSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["bucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Endpoint)) {
		query["endpoint"] = request.Endpoint
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectName)) {
		query["objectName"] = request.ObjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadFileByOSS"),
		Version:     tea.String("2022-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/UploadFileByOSS"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadFileByOSSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadFileByOSS(request *UploadFileByOSSRequest) (_result *UploadFileByOSSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadFileByOSSResponse{}
	_body, _err := client.UploadFileByOSSWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadFileByURLWithOptions(request *UploadFileByURLRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadFileByURLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadFileByURL"),
		Version:     tea.String("2022-06-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/UploadFileByURL"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadFileByURLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadFileByURL(request *UploadFileByURLRequest) (_result *UploadFileByURLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadFileByURLResponse{}
	_body, _err := client.UploadFileByURLWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
