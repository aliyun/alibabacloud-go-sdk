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

type DataProductListLogMapValue struct {
	LogCode         *string                                      `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	LogName         *string                                      `json:"LogName,omitempty" xml:"LogName,omitempty"`
	LogNameEn       *string                                      `json:"LogNameEn,omitempty" xml:"LogNameEn,omitempty"`
	LogNameKey      *string                                      `json:"LogNameKey,omitempty" xml:"LogNameKey,omitempty"`
	Status          *bool                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	CanOperateOrNot *bool                                        `json:"CanOperateOrNot,omitempty" xml:"CanOperateOrNot,omitempty"`
	Topic           *string                                      `json:"Topic,omitempty" xml:"Topic,omitempty"`
	ExtraParameters []*DataProductListLogMapValueExtraParameters `json:"ExtraParameters,omitempty" xml:"ExtraParameters,omitempty" type:"Repeated"`
}

func (s DataProductListLogMapValue) String() string {
	return tea.Prettify(s)
}

func (s DataProductListLogMapValue) GoString() string {
	return s.String()
}

func (s *DataProductListLogMapValue) SetLogCode(v string) *DataProductListLogMapValue {
	s.LogCode = &v
	return s
}

func (s *DataProductListLogMapValue) SetLogName(v string) *DataProductListLogMapValue {
	s.LogName = &v
	return s
}

func (s *DataProductListLogMapValue) SetLogNameEn(v string) *DataProductListLogMapValue {
	s.LogNameEn = &v
	return s
}

func (s *DataProductListLogMapValue) SetLogNameKey(v string) *DataProductListLogMapValue {
	s.LogNameKey = &v
	return s
}

func (s *DataProductListLogMapValue) SetStatus(v bool) *DataProductListLogMapValue {
	s.Status = &v
	return s
}

func (s *DataProductListLogMapValue) SetCanOperateOrNot(v bool) *DataProductListLogMapValue {
	s.CanOperateOrNot = &v
	return s
}

func (s *DataProductListLogMapValue) SetTopic(v string) *DataProductListLogMapValue {
	s.Topic = &v
	return s
}

func (s *DataProductListLogMapValue) SetExtraParameters(v []*DataProductListLogMapValueExtraParameters) *DataProductListLogMapValue {
	s.ExtraParameters = v
	return s
}

type DataProductListLogMapValueExtraParameters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DataProductListLogMapValueExtraParameters) String() string {
	return tea.Prettify(s)
}

func (s DataProductListLogMapValueExtraParameters) GoString() string {
	return s.String()
}

func (s *DataProductListLogMapValueExtraParameters) SetKey(v string) *DataProductListLogMapValueExtraParameters {
	s.Key = &v
	return s
}

func (s *DataProductListLogMapValueExtraParameters) SetValue(v string) *DataProductListLogMapValueExtraParameters {
	s.Value = &v
	return s
}

type BatchJobCheckRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubmitId *string `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
}

func (s BatchJobCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchJobCheckRequest) GoString() string {
	return s.String()
}

func (s *BatchJobCheckRequest) SetRegionId(v string) *BatchJobCheckRequest {
	s.RegionId = &v
	return s
}

func (s *BatchJobCheckRequest) SetSubmitId(v string) *BatchJobCheckRequest {
	s.SubmitId = &v
	return s
}

type BatchJobCheckResponseBody struct {
	Code      *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *BatchJobCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode   *string                        `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchJobCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchJobCheckResponseBody) GoString() string {
	return s.String()
}

func (s *BatchJobCheckResponseBody) SetCode(v int32) *BatchJobCheckResponseBody {
	s.Code = &v
	return s
}

func (s *BatchJobCheckResponseBody) SetData(v *BatchJobCheckResponseBodyData) *BatchJobCheckResponseBody {
	s.Data = v
	return s
}

func (s *BatchJobCheckResponseBody) SetErrCode(v string) *BatchJobCheckResponseBody {
	s.ErrCode = &v
	return s
}

func (s *BatchJobCheckResponseBody) SetMessage(v string) *BatchJobCheckResponseBody {
	s.Message = &v
	return s
}

func (s *BatchJobCheckResponseBody) SetRequestId(v string) *BatchJobCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchJobCheckResponseBody) SetSuccess(v bool) *BatchJobCheckResponseBody {
	s.Success = &v
	return s
}

type BatchJobCheckResponseBodyData struct {
	ConfigId    *string                                     `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ErrTaskList []*BatchJobCheckResponseBodyDataErrTaskList `json:"ErrTaskList,omitempty" xml:"ErrTaskList,omitempty" type:"Repeated"`
	FailedCount *int32                                      `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	FinishCount *int32                                      `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	FolderId    *string                                     `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	TaskCount   *int32                                      `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
	TaskStatus  *string                                     `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s BatchJobCheckResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchJobCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchJobCheckResponseBodyData) SetConfigId(v string) *BatchJobCheckResponseBodyData {
	s.ConfigId = &v
	return s
}

func (s *BatchJobCheckResponseBodyData) SetErrTaskList(v []*BatchJobCheckResponseBodyDataErrTaskList) *BatchJobCheckResponseBodyData {
	s.ErrTaskList = v
	return s
}

func (s *BatchJobCheckResponseBodyData) SetFailedCount(v int32) *BatchJobCheckResponseBodyData {
	s.FailedCount = &v
	return s
}

func (s *BatchJobCheckResponseBodyData) SetFinishCount(v int32) *BatchJobCheckResponseBodyData {
	s.FinishCount = &v
	return s
}

func (s *BatchJobCheckResponseBodyData) SetFolderId(v string) *BatchJobCheckResponseBodyData {
	s.FolderId = &v
	return s
}

func (s *BatchJobCheckResponseBodyData) SetTaskCount(v int32) *BatchJobCheckResponseBodyData {
	s.TaskCount = &v
	return s
}

func (s *BatchJobCheckResponseBodyData) SetTaskStatus(v string) *BatchJobCheckResponseBodyData {
	s.TaskStatus = &v
	return s
}

type BatchJobCheckResponseBodyDataErrTaskList struct {
	ProductList []*BatchJobCheckResponseBodyDataErrTaskListProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Repeated"`
	UserId      *int64                                                 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BatchJobCheckResponseBodyDataErrTaskList) String() string {
	return tea.Prettify(s)
}

func (s BatchJobCheckResponseBodyDataErrTaskList) GoString() string {
	return s.String()
}

func (s *BatchJobCheckResponseBodyDataErrTaskList) SetProductList(v []*BatchJobCheckResponseBodyDataErrTaskListProductList) *BatchJobCheckResponseBodyDataErrTaskList {
	s.ProductList = v
	return s
}

func (s *BatchJobCheckResponseBodyDataErrTaskList) SetUserId(v int64) *BatchJobCheckResponseBodyDataErrTaskList {
	s.UserId = &v
	return s
}

type BatchJobCheckResponseBodyDataErrTaskListProductList struct {
	LogList     []*BatchJobCheckResponseBodyDataErrTaskListProductListLogList `json:"LogList,omitempty" xml:"LogList,omitempty" type:"Repeated"`
	ProductCode *string                                                       `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s BatchJobCheckResponseBodyDataErrTaskListProductList) String() string {
	return tea.Prettify(s)
}

func (s BatchJobCheckResponseBodyDataErrTaskListProductList) GoString() string {
	return s.String()
}

func (s *BatchJobCheckResponseBodyDataErrTaskListProductList) SetLogList(v []*BatchJobCheckResponseBodyDataErrTaskListProductListLogList) *BatchJobCheckResponseBodyDataErrTaskListProductList {
	s.LogList = v
	return s
}

func (s *BatchJobCheckResponseBodyDataErrTaskListProductList) SetProductCode(v string) *BatchJobCheckResponseBodyDataErrTaskListProductList {
	s.ProductCode = &v
	return s
}

type BatchJobCheckResponseBodyDataErrTaskListProductListLogList struct {
	ErrorCode           *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	LogCode             *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	LogStoreNamePattern *string `json:"LogStoreNamePattern,omitempty" xml:"LogStoreNamePattern,omitempty"`
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProjectNamePattern  *string `json:"ProjectNamePattern,omitempty" xml:"ProjectNamePattern,omitempty"`
	RegionCode          *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s BatchJobCheckResponseBodyDataErrTaskListProductListLogList) String() string {
	return tea.Prettify(s)
}

func (s BatchJobCheckResponseBodyDataErrTaskListProductListLogList) GoString() string {
	return s.String()
}

func (s *BatchJobCheckResponseBodyDataErrTaskListProductListLogList) SetErrorCode(v string) *BatchJobCheckResponseBodyDataErrTaskListProductListLogList {
	s.ErrorCode = &v
	return s
}

func (s *BatchJobCheckResponseBodyDataErrTaskListProductListLogList) SetLogCode(v string) *BatchJobCheckResponseBodyDataErrTaskListProductListLogList {
	s.LogCode = &v
	return s
}

func (s *BatchJobCheckResponseBodyDataErrTaskListProductListLogList) SetLogStoreNamePattern(v string) *BatchJobCheckResponseBodyDataErrTaskListProductListLogList {
	s.LogStoreNamePattern = &v
	return s
}

func (s *BatchJobCheckResponseBodyDataErrTaskListProductListLogList) SetProductCode(v string) *BatchJobCheckResponseBodyDataErrTaskListProductListLogList {
	s.ProductCode = &v
	return s
}

func (s *BatchJobCheckResponseBodyDataErrTaskListProductListLogList) SetProjectNamePattern(v string) *BatchJobCheckResponseBodyDataErrTaskListProductListLogList {
	s.ProjectNamePattern = &v
	return s
}

func (s *BatchJobCheckResponseBodyDataErrTaskListProductListLogList) SetRegionCode(v string) *BatchJobCheckResponseBodyDataErrTaskListProductListLogList {
	s.RegionCode = &v
	return s
}

type BatchJobCheckResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchJobCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchJobCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchJobCheckResponse) GoString() string {
	return s.String()
}

func (s *BatchJobCheckResponse) SetHeaders(v map[string]*string) *BatchJobCheckResponse {
	s.Headers = v
	return s
}

func (s *BatchJobCheckResponse) SetStatusCode(v int32) *BatchJobCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchJobCheckResponse) SetBody(v *BatchJobCheckResponseBody) *BatchJobCheckResponse {
	s.Body = v
	return s
}

type BatchJobSubmitRequest struct {
	JsonConfig *string `json:"JsonConfig,omitempty" xml:"JsonConfig,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BatchJobSubmitRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchJobSubmitRequest) GoString() string {
	return s.String()
}

func (s *BatchJobSubmitRequest) SetJsonConfig(v string) *BatchJobSubmitRequest {
	s.JsonConfig = &v
	return s
}

func (s *BatchJobSubmitRequest) SetRegionId(v string) *BatchJobSubmitRequest {
	s.RegionId = &v
	return s
}

type BatchJobSubmitResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *BatchJobSubmitResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode   *string                         `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchJobSubmitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchJobSubmitResponseBody) GoString() string {
	return s.String()
}

func (s *BatchJobSubmitResponseBody) SetCode(v int32) *BatchJobSubmitResponseBody {
	s.Code = &v
	return s
}

func (s *BatchJobSubmitResponseBody) SetData(v *BatchJobSubmitResponseBodyData) *BatchJobSubmitResponseBody {
	s.Data = v
	return s
}

func (s *BatchJobSubmitResponseBody) SetErrCode(v string) *BatchJobSubmitResponseBody {
	s.ErrCode = &v
	return s
}

func (s *BatchJobSubmitResponseBody) SetMessage(v string) *BatchJobSubmitResponseBody {
	s.Message = &v
	return s
}

func (s *BatchJobSubmitResponseBody) SetRequestId(v string) *BatchJobSubmitResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchJobSubmitResponseBody) SetSuccess(v bool) *BatchJobSubmitResponseBody {
	s.Success = &v
	return s
}

type BatchJobSubmitResponseBodyData struct {
	ConfigId   *string                                     `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigList []*BatchJobSubmitResponseBodyDataConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	SubmitId   *string                                     `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
	TaskCount  *int32                                      `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
}

func (s BatchJobSubmitResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchJobSubmitResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchJobSubmitResponseBodyData) SetConfigId(v string) *BatchJobSubmitResponseBodyData {
	s.ConfigId = &v
	return s
}

func (s *BatchJobSubmitResponseBodyData) SetConfigList(v []*BatchJobSubmitResponseBodyDataConfigList) *BatchJobSubmitResponseBodyData {
	s.ConfigList = v
	return s
}

func (s *BatchJobSubmitResponseBodyData) SetSubmitId(v string) *BatchJobSubmitResponseBodyData {
	s.SubmitId = &v
	return s
}

func (s *BatchJobSubmitResponseBodyData) SetTaskCount(v int32) *BatchJobSubmitResponseBodyData {
	s.TaskCount = &v
	return s
}

type BatchJobSubmitResponseBodyDataConfigList struct {
	ProductList []*BatchJobSubmitResponseBodyDataConfigListProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Repeated"`
	UserId      *int64                                                 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BatchJobSubmitResponseBodyDataConfigList) String() string {
	return tea.Prettify(s)
}

func (s BatchJobSubmitResponseBodyDataConfigList) GoString() string {
	return s.String()
}

func (s *BatchJobSubmitResponseBodyDataConfigList) SetProductList(v []*BatchJobSubmitResponseBodyDataConfigListProductList) *BatchJobSubmitResponseBodyDataConfigList {
	s.ProductList = v
	return s
}

func (s *BatchJobSubmitResponseBodyDataConfigList) SetUserId(v int64) *BatchJobSubmitResponseBodyDataConfigList {
	s.UserId = &v
	return s
}

type BatchJobSubmitResponseBodyDataConfigListProductList struct {
	LogList     []*BatchJobSubmitResponseBodyDataConfigListProductListLogList `json:"LogList,omitempty" xml:"LogList,omitempty" type:"Repeated"`
	ProductCode *string                                                       `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s BatchJobSubmitResponseBodyDataConfigListProductList) String() string {
	return tea.Prettify(s)
}

func (s BatchJobSubmitResponseBodyDataConfigListProductList) GoString() string {
	return s.String()
}

func (s *BatchJobSubmitResponseBodyDataConfigListProductList) SetLogList(v []*BatchJobSubmitResponseBodyDataConfigListProductListLogList) *BatchJobSubmitResponseBodyDataConfigListProductList {
	s.LogList = v
	return s
}

func (s *BatchJobSubmitResponseBodyDataConfigListProductList) SetProductCode(v string) *BatchJobSubmitResponseBodyDataConfigListProductList {
	s.ProductCode = &v
	return s
}

type BatchJobSubmitResponseBodyDataConfigListProductListLogList struct {
	ErrorCode           *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	LogCode             *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	LogStoreNamePattern *string `json:"LogStoreNamePattern,omitempty" xml:"LogStoreNamePattern,omitempty"`
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProjectNamePattern  *string `json:"ProjectNamePattern,omitempty" xml:"ProjectNamePattern,omitempty"`
	RegionCode          *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s BatchJobSubmitResponseBodyDataConfigListProductListLogList) String() string {
	return tea.Prettify(s)
}

func (s BatchJobSubmitResponseBodyDataConfigListProductListLogList) GoString() string {
	return s.String()
}

func (s *BatchJobSubmitResponseBodyDataConfigListProductListLogList) SetErrorCode(v string) *BatchJobSubmitResponseBodyDataConfigListProductListLogList {
	s.ErrorCode = &v
	return s
}

func (s *BatchJobSubmitResponseBodyDataConfigListProductListLogList) SetLogCode(v string) *BatchJobSubmitResponseBodyDataConfigListProductListLogList {
	s.LogCode = &v
	return s
}

func (s *BatchJobSubmitResponseBodyDataConfigListProductListLogList) SetLogStoreNamePattern(v string) *BatchJobSubmitResponseBodyDataConfigListProductListLogList {
	s.LogStoreNamePattern = &v
	return s
}

func (s *BatchJobSubmitResponseBodyDataConfigListProductListLogList) SetProductCode(v string) *BatchJobSubmitResponseBodyDataConfigListProductListLogList {
	s.ProductCode = &v
	return s
}

func (s *BatchJobSubmitResponseBodyDataConfigListProductListLogList) SetProjectNamePattern(v string) *BatchJobSubmitResponseBodyDataConfigListProductListLogList {
	s.ProjectNamePattern = &v
	return s
}

func (s *BatchJobSubmitResponseBodyDataConfigListProductListLogList) SetRegionCode(v string) *BatchJobSubmitResponseBodyDataConfigListProductListLogList {
	s.RegionCode = &v
	return s
}

type BatchJobSubmitResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchJobSubmitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchJobSubmitResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchJobSubmitResponse) GoString() string {
	return s.String()
}

func (s *BatchJobSubmitResponse) SetHeaders(v map[string]*string) *BatchJobSubmitResponse {
	s.Headers = v
	return s
}

func (s *BatchJobSubmitResponse) SetStatusCode(v int32) *BatchJobSubmitResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchJobSubmitResponse) SetBody(v *BatchJobSubmitResponseBody) *BatchJobSubmitResponse {
	s.Body = v
	return s
}

type CloseDeliveryRequest struct {
	LogCode     *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CloseDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseDeliveryRequest) GoString() string {
	return s.String()
}

func (s *CloseDeliveryRequest) SetLogCode(v string) *CloseDeliveryRequest {
	s.LogCode = &v
	return s
}

func (s *CloseDeliveryRequest) SetProductCode(v string) *CloseDeliveryRequest {
	s.ProductCode = &v
	return s
}

func (s *CloseDeliveryRequest) SetRegionId(v string) *CloseDeliveryRequest {
	s.RegionId = &v
	return s
}

type CloseDeliveryResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	DyCode    *string `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *CloseDeliveryResponseBody) SetCode(v int32) *CloseDeliveryResponseBody {
	s.Code = &v
	return s
}

func (s *CloseDeliveryResponseBody) SetData(v bool) *CloseDeliveryResponseBody {
	s.Data = &v
	return s
}

func (s *CloseDeliveryResponseBody) SetDyCode(v string) *CloseDeliveryResponseBody {
	s.DyCode = &v
	return s
}

func (s *CloseDeliveryResponseBody) SetDyMessage(v string) *CloseDeliveryResponseBody {
	s.DyMessage = &v
	return s
}

func (s *CloseDeliveryResponseBody) SetErrCode(v string) *CloseDeliveryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CloseDeliveryResponseBody) SetMessage(v string) *CloseDeliveryResponseBody {
	s.Message = &v
	return s
}

func (s *CloseDeliveryResponseBody) SetRequestId(v string) *CloseDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseDeliveryResponseBody) SetSuccess(v bool) *CloseDeliveryResponseBody {
	s.Success = &v
	return s
}

type CloseDeliveryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CloseDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseDeliveryResponse) GoString() string {
	return s.String()
}

func (s *CloseDeliveryResponse) SetHeaders(v map[string]*string) *CloseDeliveryResponse {
	s.Headers = v
	return s
}

func (s *CloseDeliveryResponse) SetStatusCode(v int32) *CloseDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseDeliveryResponse) SetBody(v *CloseDeliveryResponseBody) *CloseDeliveryResponse {
	s.Body = v
	return s
}

type DeleteAutomateResponseConfigRequest struct {
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAutomateResponseConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutomateResponseConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutomateResponseConfigRequest) SetId(v int64) *DeleteAutomateResponseConfigRequest {
	s.Id = &v
	return s
}

func (s *DeleteAutomateResponseConfigRequest) SetRegionId(v string) *DeleteAutomateResponseConfigRequest {
	s.RegionId = &v
	return s
}

type DeleteAutomateResponseConfigResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAutomateResponseConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutomateResponseConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutomateResponseConfigResponseBody) SetCode(v int32) *DeleteAutomateResponseConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAutomateResponseConfigResponseBody) SetData(v string) *DeleteAutomateResponseConfigResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAutomateResponseConfigResponseBody) SetMessage(v string) *DeleteAutomateResponseConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAutomateResponseConfigResponseBody) SetRequestId(v string) *DeleteAutomateResponseConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAutomateResponseConfigResponseBody) SetSuccess(v bool) *DeleteAutomateResponseConfigResponseBody {
	s.Success = &v
	return s
}

type DeleteAutomateResponseConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAutomateResponseConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAutomateResponseConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutomateResponseConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutomateResponseConfigResponse) SetHeaders(v map[string]*string) *DeleteAutomateResponseConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutomateResponseConfigResponse) SetStatusCode(v int32) *DeleteAutomateResponseConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutomateResponseConfigResponse) SetBody(v *DeleteAutomateResponseConfigResponseBody) *DeleteAutomateResponseConfigResponse {
	s.Body = v
	return s
}

type DeleteCustomizeRuleRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleId   *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteCustomizeRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizeRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomizeRuleRequest) SetRegionId(v string) *DeleteCustomizeRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomizeRuleRequest) SetRuleId(v int64) *DeleteCustomizeRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteCustomizeRuleResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int32  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomizeRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomizeRuleResponseBody) SetCode(v int32) *DeleteCustomizeRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomizeRuleResponseBody) SetData(v int32) *DeleteCustomizeRuleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCustomizeRuleResponseBody) SetMessage(v string) *DeleteCustomizeRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomizeRuleResponseBody) SetRequestId(v string) *DeleteCustomizeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomizeRuleResponseBody) SetSuccess(v bool) *DeleteCustomizeRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteCustomizeRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCustomizeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCustomizeRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomizeRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomizeRuleResponse) SetHeaders(v map[string]*string) *DeleteCustomizeRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomizeRuleResponse) SetStatusCode(v int32) *DeleteCustomizeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomizeRuleResponse) SetBody(v *DeleteCustomizeRuleResponseBody) *DeleteCustomizeRuleResponse {
	s.Body = v
	return s
}

type DeleteQuickQueryRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s DeleteQuickQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQuickQueryRequest) GoString() string {
	return s.String()
}

func (s *DeleteQuickQueryRequest) SetRegionId(v string) *DeleteQuickQueryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteQuickQueryRequest) SetSearchName(v string) *DeleteQuickQueryRequest {
	s.SearchName = &v
	return s
}

type DeleteQuickQueryResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	DyCode    *string `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQuickQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQuickQueryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQuickQueryResponseBody) SetCode(v int32) *DeleteQuickQueryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQuickQueryResponseBody) SetData(v bool) *DeleteQuickQueryResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQuickQueryResponseBody) SetDyCode(v string) *DeleteQuickQueryResponseBody {
	s.DyCode = &v
	return s
}

func (s *DeleteQuickQueryResponseBody) SetDyMessage(v string) *DeleteQuickQueryResponseBody {
	s.DyMessage = &v
	return s
}

func (s *DeleteQuickQueryResponseBody) SetErrCode(v string) *DeleteQuickQueryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteQuickQueryResponseBody) SetMessage(v string) *DeleteQuickQueryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteQuickQueryResponseBody) SetRequestId(v string) *DeleteQuickQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQuickQueryResponseBody) SetSuccess(v bool) *DeleteQuickQueryResponseBody {
	s.Success = &v
	return s
}

type DeleteQuickQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteQuickQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteQuickQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQuickQueryResponse) GoString() string {
	return s.String()
}

func (s *DeleteQuickQueryResponse) SetHeaders(v map[string]*string) *DeleteQuickQueryResponse {
	s.Headers = v
	return s
}

func (s *DeleteQuickQueryResponse) SetStatusCode(v int32) *DeleteQuickQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQuickQueryResponse) SetBody(v *DeleteQuickQueryResponseBody) *DeleteQuickQueryResponse {
	s.Body = v
	return s
}

type DeleteWhiteRuleListRequest struct {
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteWhiteRuleListRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWhiteRuleListRequest) GoString() string {
	return s.String()
}

func (s *DeleteWhiteRuleListRequest) SetId(v int64) *DeleteWhiteRuleListRequest {
	s.Id = &v
	return s
}

func (s *DeleteWhiteRuleListRequest) SetRegionId(v string) *DeleteWhiteRuleListRequest {
	s.RegionId = &v
	return s
}

type DeleteWhiteRuleListResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWhiteRuleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWhiteRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWhiteRuleListResponseBody) SetCode(v int32) *DeleteWhiteRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWhiteRuleListResponseBody) SetData(v interface{}) *DeleteWhiteRuleListResponseBody {
	s.Data = v
	return s
}

func (s *DeleteWhiteRuleListResponseBody) SetMessage(v string) *DeleteWhiteRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWhiteRuleListResponseBody) SetRequestId(v string) *DeleteWhiteRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWhiteRuleListResponseBody) SetSuccess(v bool) *DeleteWhiteRuleListResponseBody {
	s.Success = &v
	return s
}

type DeleteWhiteRuleListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteWhiteRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWhiteRuleListResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWhiteRuleListResponse) GoString() string {
	return s.String()
}

func (s *DeleteWhiteRuleListResponse) SetHeaders(v map[string]*string) *DeleteWhiteRuleListResponse {
	s.Headers = v
	return s
}

func (s *DeleteWhiteRuleListResponse) SetStatusCode(v int32) *DeleteWhiteRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWhiteRuleListResponse) SetBody(v *DeleteWhiteRuleListResponseBody) *DeleteWhiteRuleListResponse {
	s.Body = v
	return s
}

type DescribeAggregateFunctionRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAggregateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAggregateFunctionRequest) GoString() string {
	return s.String()
}

func (s *DescribeAggregateFunctionRequest) SetRegionId(v string) *DescribeAggregateFunctionRequest {
	s.RegionId = &v
	return s
}

type DescribeAggregateFunctionResponseBody struct {
	Code      *int32                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeAggregateFunctionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAggregateFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAggregateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAggregateFunctionResponseBody) SetCode(v int32) *DescribeAggregateFunctionResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAggregateFunctionResponseBody) SetData(v []*DescribeAggregateFunctionResponseBodyData) *DescribeAggregateFunctionResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAggregateFunctionResponseBody) SetMessage(v string) *DescribeAggregateFunctionResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAggregateFunctionResponseBody) SetRequestId(v string) *DescribeAggregateFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAggregateFunctionResponseBody) SetSuccess(v bool) *DescribeAggregateFunctionResponseBody {
	s.Success = &v
	return s
}

type DescribeAggregateFunctionResponseBodyData struct {
	Function     *string `json:"Function,omitempty" xml:"Function,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s DescribeAggregateFunctionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAggregateFunctionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAggregateFunctionResponseBodyData) SetFunction(v string) *DescribeAggregateFunctionResponseBodyData {
	s.Function = &v
	return s
}

func (s *DescribeAggregateFunctionResponseBodyData) SetFunctionName(v string) *DescribeAggregateFunctionResponseBodyData {
	s.FunctionName = &v
	return s
}

type DescribeAggregateFunctionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAggregateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAggregateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAggregateFunctionResponse) GoString() string {
	return s.String()
}

func (s *DescribeAggregateFunctionResponse) SetHeaders(v map[string]*string) *DescribeAggregateFunctionResponse {
	s.Headers = v
	return s
}

func (s *DescribeAggregateFunctionResponse) SetStatusCode(v int32) *DescribeAggregateFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAggregateFunctionResponse) SetBody(v *DescribeAggregateFunctionResponseBody) *DescribeAggregateFunctionResponse {
	s.Body = v
	return s
}

type DescribeAlertSceneByEventRequest struct {
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAlertSceneByEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSceneByEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneByEventRequest) SetIncidentUuid(v string) *DescribeAlertSceneByEventRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAlertSceneByEventRequest) SetRegionId(v string) *DescribeAlertSceneByEventRequest {
	s.RegionId = &v
	return s
}

type DescribeAlertSceneByEventResponseBody struct {
	Code      *int32                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeAlertSceneByEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertSceneByEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSceneByEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneByEventResponseBody) SetCode(v int32) *DescribeAlertSceneByEventResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBody) SetData(v []*DescribeAlertSceneByEventResponseBodyData) *DescribeAlertSceneByEventResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertSceneByEventResponseBody) SetMessage(v string) *DescribeAlertSceneByEventResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBody) SetRequestId(v string) *DescribeAlertSceneByEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBody) SetSuccess(v bool) *DescribeAlertSceneByEventResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertSceneByEventResponseBodyData struct {
	AlertName   *string                                             `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	AlertNameId *string                                             `json:"AlertNameId,omitempty" xml:"AlertNameId,omitempty"`
	AlertTile   *string                                             `json:"AlertTile,omitempty" xml:"AlertTile,omitempty"`
	AlertTileId *string                                             `json:"AlertTileId,omitempty" xml:"AlertTileId,omitempty"`
	AlertType   *string                                             `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AlertTypeId *string                                             `json:"AlertTypeId,omitempty" xml:"AlertTypeId,omitempty"`
	Targets     []*DescribeAlertSceneByEventResponseBodyDataTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s DescribeAlertSceneByEventResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSceneByEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertName(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertNameId(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertNameId = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertTile(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertTile = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertTileId(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertTileId = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertType(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetAlertTypeId(v string) *DescribeAlertSceneByEventResponseBodyData {
	s.AlertTypeId = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyData) SetTargets(v []*DescribeAlertSceneByEventResponseBodyDataTargets) *DescribeAlertSceneByEventResponseBodyData {
	s.Targets = v
	return s
}

type DescribeAlertSceneByEventResponseBodyDataTargets struct {
	Name   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Type   *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Value  *string   `json:"Value,omitempty" xml:"Value,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeAlertSceneByEventResponseBodyDataTargets) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSceneByEventResponseBodyDataTargets) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) SetName(v string) *DescribeAlertSceneByEventResponseBodyDataTargets {
	s.Name = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) SetType(v string) *DescribeAlertSceneByEventResponseBodyDataTargets {
	s.Type = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) SetValue(v string) *DescribeAlertSceneByEventResponseBodyDataTargets {
	s.Value = &v
	return s
}

func (s *DescribeAlertSceneByEventResponseBodyDataTargets) SetValues(v []*string) *DescribeAlertSceneByEventResponseBodyDataTargets {
	s.Values = v
	return s
}

type DescribeAlertSceneByEventResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAlertSceneByEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAlertSceneByEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSceneByEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneByEventResponse) SetHeaders(v map[string]*string) *DescribeAlertSceneByEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertSceneByEventResponse) SetStatusCode(v int32) *DescribeAlertSceneByEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertSceneByEventResponse) SetBody(v *DescribeAlertSceneByEventResponseBody) *DescribeAlertSceneByEventResponse {
	s.Body = v
	return s
}

type DescribeAlertSourceRequest struct {
	EndTime   *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Level     []*string `json:"Level,omitempty" xml:"Level,omitempty" type:"Repeated"`
	RegionId  *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAlertSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceRequest) SetEndTime(v int64) *DescribeAlertSourceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertSourceRequest) SetLevel(v []*string) *DescribeAlertSourceRequest {
	s.Level = v
	return s
}

func (s *DescribeAlertSourceRequest) SetRegionId(v string) *DescribeAlertSourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertSourceRequest) SetStartTime(v int64) *DescribeAlertSourceRequest {
	s.StartTime = &v
	return s
}

type DescribeAlertSourceResponseBody struct {
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeAlertSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceResponseBody) SetCode(v int32) *DescribeAlertSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertSourceResponseBody) SetData(v []*DescribeAlertSourceResponseBodyData) *DescribeAlertSourceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertSourceResponseBody) SetMessage(v string) *DescribeAlertSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertSourceResponseBody) SetRequestId(v string) *DescribeAlertSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertSourceResponseBody) SetSuccess(v bool) *DescribeAlertSourceResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertSourceResponseBodyData struct {
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
}

func (s DescribeAlertSourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceResponseBodyData) SetSource(v string) *DescribeAlertSourceResponseBodyData {
	s.Source = &v
	return s
}

func (s *DescribeAlertSourceResponseBodyData) SetSourceName(v string) *DescribeAlertSourceResponseBodyData {
	s.SourceName = &v
	return s
}

type DescribeAlertSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAlertSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAlertSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceResponse) SetHeaders(v map[string]*string) *DescribeAlertSourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertSourceResponse) SetStatusCode(v int32) *DescribeAlertSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertSourceResponse) SetBody(v *DescribeAlertSourceResponseBody) *DescribeAlertSourceResponse {
	s.Body = v
	return s
}

type DescribeAlertSourceWithEventRequest struct {
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAlertSourceWithEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSourceWithEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceWithEventRequest) SetIncidentUuid(v string) *DescribeAlertSourceWithEventRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAlertSourceWithEventRequest) SetRegionId(v string) *DescribeAlertSourceWithEventRequest {
	s.RegionId = &v
	return s
}

type DescribeAlertSourceWithEventResponseBody struct {
	Code      *int32                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeAlertSourceWithEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertSourceWithEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSourceWithEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceWithEventResponseBody) SetCode(v int32) *DescribeAlertSourceWithEventResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBody) SetData(v []*DescribeAlertSourceWithEventResponseBodyData) *DescribeAlertSourceWithEventResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBody) SetMessage(v string) *DescribeAlertSourceWithEventResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBody) SetRequestId(v string) *DescribeAlertSourceWithEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBody) SetSuccess(v bool) *DescribeAlertSourceWithEventResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertSourceWithEventResponseBodyData struct {
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
}

func (s DescribeAlertSourceWithEventResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSourceWithEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceWithEventResponseBodyData) SetSource(v string) *DescribeAlertSourceWithEventResponseBodyData {
	s.Source = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponseBodyData) SetSourceName(v string) *DescribeAlertSourceWithEventResponseBodyData {
	s.SourceName = &v
	return s
}

type DescribeAlertSourceWithEventResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAlertSourceWithEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAlertSourceWithEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertSourceWithEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceWithEventResponse) SetHeaders(v map[string]*string) *DescribeAlertSourceWithEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertSourceWithEventResponse) SetStatusCode(v int32) *DescribeAlertSourceWithEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertSourceWithEventResponse) SetBody(v *DescribeAlertSourceWithEventResponseBody) *DescribeAlertSourceWithEventResponse {
	s.Body = v
	return s
}

type DescribeAlertTypeRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAlertTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertTypeRequest) SetRegionId(v string) *DescribeAlertTypeRequest {
	s.RegionId = &v
	return s
}

type DescribeAlertTypeResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeAlertTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertTypeResponseBody) SetCode(v int32) *DescribeAlertTypeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertTypeResponseBody) SetData(v []*DescribeAlertTypeResponseBodyData) *DescribeAlertTypeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertTypeResponseBody) SetMessage(v string) *DescribeAlertTypeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertTypeResponseBody) SetRequestId(v string) *DescribeAlertTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertTypeResponseBody) SetSuccess(v bool) *DescribeAlertTypeResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertTypeResponseBodyData struct {
	AlertType    *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AlertTypeMds *string `json:"AlertTypeMds,omitempty" xml:"AlertTypeMds,omitempty"`
}

func (s DescribeAlertTypeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertTypeResponseBodyData) SetAlertType(v string) *DescribeAlertTypeResponseBodyData {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertTypeResponseBodyData) SetAlertTypeMds(v string) *DescribeAlertTypeResponseBodyData {
	s.AlertTypeMds = &v
	return s
}

type DescribeAlertTypeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAlertTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAlertTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertTypeResponse) SetHeaders(v map[string]*string) *DescribeAlertTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertTypeResponse) SetStatusCode(v int32) *DescribeAlertTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertTypeResponse) SetBody(v *DescribeAlertTypeResponseBody) *DescribeAlertTypeResponse {
	s.Body = v
	return s
}

type DescribeAlertsCountRequest struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAlertsCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertsCountRequest) SetEndTime(v int64) *DescribeAlertsCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertsCountRequest) SetRegionId(v string) *DescribeAlertsCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertsCountRequest) SetStartTime(v int64) *DescribeAlertsCountRequest {
	s.StartTime = &v
	return s
}

type DescribeAlertsCountResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeAlertsCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertsCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertsCountResponseBody) SetCode(v int32) *DescribeAlertsCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertsCountResponseBody) SetData(v *DescribeAlertsCountResponseBodyData) *DescribeAlertsCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertsCountResponseBody) SetMessage(v string) *DescribeAlertsCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertsCountResponseBody) SetRequestId(v string) *DescribeAlertsCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertsCountResponseBody) SetSuccess(v bool) *DescribeAlertsCountResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertsCountResponseBodyData struct {
	All        *int64 `json:"All,omitempty" xml:"All,omitempty"`
	High       *int64 `json:"High,omitempty" xml:"High,omitempty"`
	Low        *int64 `json:"Low,omitempty" xml:"Low,omitempty"`
	Medium     *int64 `json:"Medium,omitempty" xml:"Medium,omitempty"`
	ProductNum *int32 `json:"ProductNum,omitempty" xml:"ProductNum,omitempty"`
}

func (s DescribeAlertsCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertsCountResponseBodyData) SetAll(v int64) *DescribeAlertsCountResponseBodyData {
	s.All = &v
	return s
}

func (s *DescribeAlertsCountResponseBodyData) SetHigh(v int64) *DescribeAlertsCountResponseBodyData {
	s.High = &v
	return s
}

func (s *DescribeAlertsCountResponseBodyData) SetLow(v int64) *DescribeAlertsCountResponseBodyData {
	s.Low = &v
	return s
}

func (s *DescribeAlertsCountResponseBodyData) SetMedium(v int64) *DescribeAlertsCountResponseBodyData {
	s.Medium = &v
	return s
}

func (s *DescribeAlertsCountResponseBodyData) SetProductNum(v int32) *DescribeAlertsCountResponseBodyData {
	s.ProductNum = &v
	return s
}

type DescribeAlertsCountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAlertsCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAlertsCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertsCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertsCountResponse) SetHeaders(v map[string]*string) *DescribeAlertsCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertsCountResponse) SetStatusCode(v int32) *DescribeAlertsCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertsCountResponse) SetBody(v *DescribeAlertsCountResponseBody) *DescribeAlertsCountResponse {
	s.Body = v
	return s
}

type DescribeAttackTimeLineRequest struct {
	AssetName    *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAttackTimeLineRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAttackTimeLineRequest) GoString() string {
	return s.String()
}

func (s *DescribeAttackTimeLineRequest) SetAssetName(v string) *DescribeAttackTimeLineRequest {
	s.AssetName = &v
	return s
}

func (s *DescribeAttackTimeLineRequest) SetEndTime(v int64) *DescribeAttackTimeLineRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAttackTimeLineRequest) SetIncidentUuid(v string) *DescribeAttackTimeLineRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAttackTimeLineRequest) SetRegionId(v string) *DescribeAttackTimeLineRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAttackTimeLineRequest) SetStartTime(v int64) *DescribeAttackTimeLineRequest {
	s.StartTime = &v
	return s
}

type DescribeAttackTimeLineResponseBody struct {
	Code      *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeAttackTimeLineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAttackTimeLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAttackTimeLineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAttackTimeLineResponseBody) SetCode(v int32) *DescribeAttackTimeLineResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBody) SetData(v []*DescribeAttackTimeLineResponseBodyData) *DescribeAttackTimeLineResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAttackTimeLineResponseBody) SetMessage(v string) *DescribeAttackTimeLineResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBody) SetRequestId(v string) *DescribeAttackTimeLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBody) SetSuccess(v bool) *DescribeAttackTimeLineResponseBody {
	s.Success = &v
	return s
}

type DescribeAttackTimeLineResponseBodyData struct {
	AlertLevel         *string `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	AlertName          *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	AlertNameCode      *string `json:"AlertNameCode,omitempty" xml:"AlertNameCode,omitempty"`
	AlertNameEn        *string `json:"AlertNameEn,omitempty" xml:"AlertNameEn,omitempty"`
	AlertSrcProd       *string `json:"AlertSrcProd,omitempty" xml:"AlertSrcProd,omitempty"`
	AlertSrcProdModule *string `json:"AlertSrcProdModule,omitempty" xml:"AlertSrcProdModule,omitempty"`
	AlertTime          *int64  `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	AlertTitle         *string `json:"AlertTitle,omitempty" xml:"AlertTitle,omitempty"`
	AlertTitleEn       *string `json:"AlertTitleEn,omitempty" xml:"AlertTitleEn,omitempty"`
	AlertType          *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AlertTypeCode      *string `json:"AlertTypeCode,omitempty" xml:"AlertTypeCode,omitempty"`
	AlertTypeEn        *string `json:"AlertTypeEn,omitempty" xml:"AlertTypeEn,omitempty"`
	AlertUuid          *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	AssetId            *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	AssetList          *string `json:"AssetList,omitempty" xml:"AssetList,omitempty"`
	AssetName          *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	AttCk              *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	CloudCode          *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	IncidentUuid       *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	LogTime            *string `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
}

func (s DescribeAttackTimeLineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAttackTimeLineResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAlertLevel(v string) *DescribeAttackTimeLineResponseBodyData {
	s.AlertLevel = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAlertName(v string) *DescribeAttackTimeLineResponseBodyData {
	s.AlertName = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAlertNameCode(v string) *DescribeAttackTimeLineResponseBodyData {
	s.AlertNameCode = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAlertNameEn(v string) *DescribeAttackTimeLineResponseBodyData {
	s.AlertNameEn = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAlertSrcProd(v string) *DescribeAttackTimeLineResponseBodyData {
	s.AlertSrcProd = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAlertSrcProdModule(v string) *DescribeAttackTimeLineResponseBodyData {
	s.AlertSrcProdModule = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAlertTime(v int64) *DescribeAttackTimeLineResponseBodyData {
	s.AlertTime = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAlertTitle(v string) *DescribeAttackTimeLineResponseBodyData {
	s.AlertTitle = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAlertTitleEn(v string) *DescribeAttackTimeLineResponseBodyData {
	s.AlertTitleEn = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAlertType(v string) *DescribeAttackTimeLineResponseBodyData {
	s.AlertType = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAlertTypeCode(v string) *DescribeAttackTimeLineResponseBodyData {
	s.AlertTypeCode = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAlertTypeEn(v string) *DescribeAttackTimeLineResponseBodyData {
	s.AlertTypeEn = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAlertUuid(v string) *DescribeAttackTimeLineResponseBodyData {
	s.AlertUuid = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAssetId(v string) *DescribeAttackTimeLineResponseBodyData {
	s.AssetId = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAssetList(v string) *DescribeAttackTimeLineResponseBodyData {
	s.AssetList = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAssetName(v string) *DescribeAttackTimeLineResponseBodyData {
	s.AssetName = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetAttCk(v string) *DescribeAttackTimeLineResponseBodyData {
	s.AttCk = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetCloudCode(v string) *DescribeAttackTimeLineResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetIncidentUuid(v string) *DescribeAttackTimeLineResponseBodyData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAttackTimeLineResponseBodyData) SetLogTime(v string) *DescribeAttackTimeLineResponseBodyData {
	s.LogTime = &v
	return s
}

type DescribeAttackTimeLineResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAttackTimeLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAttackTimeLineResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAttackTimeLineResponse) GoString() string {
	return s.String()
}

func (s *DescribeAttackTimeLineResponse) SetHeaders(v map[string]*string) *DescribeAttackTimeLineResponse {
	s.Headers = v
	return s
}

func (s *DescribeAttackTimeLineResponse) SetStatusCode(v int32) *DescribeAttackTimeLineResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAttackTimeLineResponse) SetBody(v *DescribeAttackTimeLineResponseBody) *DescribeAttackTimeLineResponse {
	s.Body = v
	return s
}

type DescribeAutomateResponseConfigCounterRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAutomateResponseConfigCounterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigCounterRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigCounterRequest) SetRegionId(v string) *DescribeAutomateResponseConfigCounterRequest {
	s.RegionId = &v
	return s
}

type DescribeAutomateResponseConfigCounterResponseBody struct {
	Code      *int32                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeAutomateResponseConfigCounterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAutomateResponseConfigCounterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigCounterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) SetCode(v int32) *DescribeAutomateResponseConfigCounterResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) SetData(v *DescribeAutomateResponseConfigCounterResponseBodyData) *DescribeAutomateResponseConfigCounterResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) SetMessage(v string) *DescribeAutomateResponseConfigCounterResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) SetRequestId(v string) *DescribeAutomateResponseConfigCounterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBody) SetSuccess(v bool) *DescribeAutomateResponseConfigCounterResponseBody {
	s.Success = &v
	return s
}

type DescribeAutomateResponseConfigCounterResponseBodyData struct {
	All    *int64 `json:"All,omitempty" xml:"All,omitempty"`
	Online *int64 `json:"Online,omitempty" xml:"Online,omitempty"`
}

func (s DescribeAutomateResponseConfigCounterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigCounterResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigCounterResponseBodyData) SetAll(v int64) *DescribeAutomateResponseConfigCounterResponseBodyData {
	s.All = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponseBodyData) SetOnline(v int64) *DescribeAutomateResponseConfigCounterResponseBodyData {
	s.Online = &v
	return s
}

type DescribeAutomateResponseConfigCounterResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAutomateResponseConfigCounterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutomateResponseConfigCounterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigCounterResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigCounterResponse) SetHeaders(v map[string]*string) *DescribeAutomateResponseConfigCounterResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponse) SetStatusCode(v int32) *DescribeAutomateResponseConfigCounterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutomateResponseConfigCounterResponse) SetBody(v *DescribeAutomateResponseConfigCounterResponseBody) *DescribeAutomateResponseConfigCounterResponse {
	s.Body = v
	return s
}

type DescribeAutomateResponseConfigFeatureRequest struct {
	AutoResponseType *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAutomateResponseConfigFeatureRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureRequest) SetAutoResponseType(v string) *DescribeAutomateResponseConfigFeatureRequest {
	s.AutoResponseType = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureRequest) SetRegionId(v string) *DescribeAutomateResponseConfigFeatureRequest {
	s.RegionId = &v
	return s
}

type DescribeAutomateResponseConfigFeatureResponseBody struct {
	Code      *int32                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeAutomateResponseConfigFeatureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAutomateResponseConfigFeatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) SetCode(v int32) *DescribeAutomateResponseConfigFeatureResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) SetData(v []*DescribeAutomateResponseConfigFeatureResponseBodyData) *DescribeAutomateResponseConfigFeatureResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) SetMessage(v string) *DescribeAutomateResponseConfigFeatureResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) SetRequestId(v string) *DescribeAutomateResponseConfigFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) SetSuccess(v bool) *DescribeAutomateResponseConfigFeatureResponseBody {
	s.Success = &v
	return s
}

type DescribeAutomateResponseConfigFeatureResponseBodyData struct {
	DataType         *string                                                                  `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Feature          *string                                                                  `json:"Feature,omitempty" xml:"Feature,omitempty"`
	RightValueEnums  []*DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums  `json:"RightValueEnums,omitempty" xml:"RightValueEnums,omitempty" type:"Repeated"`
	SupportOperators []*DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators `json:"SupportOperators,omitempty" xml:"SupportOperators,omitempty" type:"Repeated"`
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) SetDataType(v string) *DescribeAutomateResponseConfigFeatureResponseBodyData {
	s.DataType = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) SetFeature(v string) *DescribeAutomateResponseConfigFeatureResponseBodyData {
	s.Feature = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) SetRightValueEnums(v []*DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) *DescribeAutomateResponseConfigFeatureResponseBodyData {
	s.RightValueEnums = v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) SetSupportOperators(v []*DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) *DescribeAutomateResponseConfigFeatureResponseBodyData {
	s.SupportOperators = v
	return s
}

type DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums struct {
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueMds *string `json:"ValueMds,omitempty" xml:"ValueMds,omitempty"`
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) SetValue(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums {
	s.Value = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) SetValueMds(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums {
	s.ValueMds = &v
	return s
}

type DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators struct {
	HasRightValue   *bool     `json:"HasRightValue,omitempty" xml:"HasRightValue,omitempty"`
	Index           *int32    `json:"Index,omitempty" xml:"Index,omitempty"`
	Operator        *string   `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorDescCn  *string   `json:"OperatorDescCn,omitempty" xml:"OperatorDescCn,omitempty"`
	OperatorDescEn  *string   `json:"OperatorDescEn,omitempty" xml:"OperatorDescEn,omitempty"`
	OperatorName    *string   `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	SupportDataType *string   `json:"SupportDataType,omitempty" xml:"SupportDataType,omitempty"`
	SupportTag      []*string `json:"SupportTag,omitempty" xml:"SupportTag,omitempty" type:"Repeated"`
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetHasRightValue(v bool) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.HasRightValue = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetIndex(v int32) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.Index = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetOperator(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.Operator = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetOperatorDescCn(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.OperatorDescCn = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetOperatorDescEn(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.OperatorDescEn = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetOperatorName(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.OperatorName = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetSupportDataType(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.SupportDataType = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetSupportTag(v []*string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.SupportTag = v
	return s
}

type DescribeAutomateResponseConfigFeatureResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAutomateResponseConfigFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutomateResponseConfigFeatureResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureResponse) SetHeaders(v map[string]*string) *DescribeAutomateResponseConfigFeatureResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponse) SetStatusCode(v int32) *DescribeAutomateResponseConfigFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponse) SetBody(v *DescribeAutomateResponseConfigFeatureResponseBody) *DescribeAutomateResponseConfigFeatureResponse {
	s.Body = v
	return s
}

type DescribeAutomateResponseConfigPlayBooksRequest struct {
	AutoResponseType *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
	EntityType       *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAutomateResponseConfigPlayBooksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigPlayBooksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) SetAutoResponseType(v string) *DescribeAutomateResponseConfigPlayBooksRequest {
	s.AutoResponseType = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) SetEntityType(v string) *DescribeAutomateResponseConfigPlayBooksRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksRequest) SetRegionId(v string) *DescribeAutomateResponseConfigPlayBooksRequest {
	s.RegionId = &v
	return s
}

type DescribeAutomateResponseConfigPlayBooksResponseBody struct {
	Code      *int32                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeAutomateResponseConfigPlayBooksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAutomateResponseConfigPlayBooksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigPlayBooksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) SetCode(v int32) *DescribeAutomateResponseConfigPlayBooksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) SetData(v []*DescribeAutomateResponseConfigPlayBooksResponseBodyData) *DescribeAutomateResponseConfigPlayBooksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) SetMessage(v string) *DescribeAutomateResponseConfigPlayBooksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) SetRequestId(v string) *DescribeAutomateResponseConfigPlayBooksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) SetSuccess(v bool) *DescribeAutomateResponseConfigPlayBooksResponseBody {
	s.Success = &v
	return s
}

type DescribeAutomateResponseConfigPlayBooksResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParamType   *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	Uuid        *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeAutomateResponseConfigPlayBooksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigPlayBooksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) SetDescription(v string) *DescribeAutomateResponseConfigPlayBooksResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) SetDisplayName(v string) *DescribeAutomateResponseConfigPlayBooksResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) SetName(v string) *DescribeAutomateResponseConfigPlayBooksResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) SetParamType(v string) *DescribeAutomateResponseConfigPlayBooksResponseBodyData {
	s.ParamType = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) SetUuid(v string) *DescribeAutomateResponseConfigPlayBooksResponseBodyData {
	s.Uuid = &v
	return s
}

type DescribeAutomateResponseConfigPlayBooksResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAutomateResponseConfigPlayBooksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutomateResponseConfigPlayBooksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutomateResponseConfigPlayBooksResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigPlayBooksResponse) SetHeaders(v map[string]*string) *DescribeAutomateResponseConfigPlayBooksResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponse) SetStatusCode(v int32) *DescribeAutomateResponseConfigPlayBooksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponse) SetBody(v *DescribeAutomateResponseConfigPlayBooksResponseBody) *DescribeAutomateResponseConfigPlayBooksResponse {
	s.Body = v
	return s
}

type DescribeCloudSiemAssetsRequest struct {
	AssetType    *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudSiemAssetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsRequest) SetAssetType(v string) *DescribeCloudSiemAssetsRequest {
	s.AssetType = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetCurrentPage(v int32) *DescribeCloudSiemAssetsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetIncidentUuid(v string) *DescribeCloudSiemAssetsRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetPageSize(v int32) *DescribeCloudSiemAssetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudSiemAssetsRequest) SetRegionId(v string) *DescribeCloudSiemAssetsRequest {
	s.RegionId = &v
	return s
}

type DescribeCloudSiemAssetsResponseBody struct {
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeCloudSiemAssetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudSiemAssetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponseBody) SetCode(v int32) *DescribeCloudSiemAssetsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBody) SetData(v *DescribeCloudSiemAssetsResponseBodyData) *DescribeCloudSiemAssetsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBody) SetMessage(v string) *DescribeCloudSiemAssetsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBody) SetRequestId(v string) *DescribeCloudSiemAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBody) SetSuccess(v bool) *DescribeCloudSiemAssetsResponseBody {
	s.Success = &v
	return s
}

type DescribeCloudSiemAssetsResponseBodyData struct {
	PageInfo     *DescribeCloudSiemAssetsResponseBodyDataPageInfo       `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	ResponseData []*DescribeCloudSiemAssetsResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeCloudSiemAssetsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponseBodyData) SetPageInfo(v *DescribeCloudSiemAssetsResponseBodyDataPageInfo) *DescribeCloudSiemAssetsResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyData) SetResponseData(v []*DescribeCloudSiemAssetsResponseBodyDataResponseData) *DescribeCloudSiemAssetsResponseBodyData {
	s.ResponseData = v
	return s
}

type DescribeCloudSiemAssetsResponseBodyDataPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudSiemAssetsResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeCloudSiemAssetsResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeCloudSiemAssetsResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeCloudSiemAssetsResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeCloudSiemAssetsResponseBodyDataResponseData struct {
	AlertUuid    *string                                                         `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	Aliuid       *int64                                                          `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	AssetId      *string                                                         `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	AssetInfo    []*DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo `json:"AssetInfo,omitempty" xml:"AssetInfo,omitempty" type:"Repeated"`
	AssetName    *string                                                         `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	AssetType    *string                                                         `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	CloudCode    *string                                                         `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	GmtCreate    *string                                                         `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *string                                                         `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id           *int64                                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	IncidentUuid *string                                                         `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	SubUserId    *int64                                                          `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DescribeCloudSiemAssetsResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAlertUuid(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.AlertUuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAliuid(v int64) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAssetId(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.AssetId = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAssetInfo(v []*DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.AssetInfo = v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAssetName(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.AssetName = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAssetType(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.AssetType = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetCloudCode(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.CloudCode = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetGmtCreate(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetGmtModified(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetId(v int64) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetIncidentUuid(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetSubUserId(v int64) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

type DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo struct {
	Key     *string `json:"Key,omitempty" xml:"Key,omitempty"`
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	Values  *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) SetKey(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo {
	s.Key = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) SetKeyName(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo {
	s.KeyName = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) SetValues(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo {
	s.Values = &v
	return s
}

type DescribeCloudSiemAssetsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCloudSiemAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCloudSiemAssetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponse) SetHeaders(v map[string]*string) *DescribeCloudSiemAssetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudSiemAssetsResponse) SetStatusCode(v int32) *DescribeCloudSiemAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponse) SetBody(v *DescribeCloudSiemAssetsResponseBody) *DescribeCloudSiemAssetsResponse {
	s.Body = v
	return s
}

type DescribeCloudSiemAssetsCounterRequest struct {
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudSiemAssetsCounterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsCounterRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsCounterRequest) SetIncidentUuid(v string) *DescribeCloudSiemAssetsCounterRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterRequest) SetRegionId(v string) *DescribeCloudSiemAssetsCounterRequest {
	s.RegionId = &v
	return s
}

type DescribeCloudSiemAssetsCounterResponseBody struct {
	Code      *int32                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeCloudSiemAssetsCounterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudSiemAssetsCounterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsCounterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) SetCode(v int32) *DescribeCloudSiemAssetsCounterResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) SetData(v []*DescribeCloudSiemAssetsCounterResponseBodyData) *DescribeCloudSiemAssetsCounterResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) SetMessage(v string) *DescribeCloudSiemAssetsCounterResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) SetRequestId(v string) *DescribeCloudSiemAssetsCounterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) SetSuccess(v bool) *DescribeCloudSiemAssetsCounterResponseBody {
	s.Success = &v
	return s
}

type DescribeCloudSiemAssetsCounterResponseBodyData struct {
	AssetNum  *int32  `json:"AssetNum,omitempty" xml:"AssetNum,omitempty"`
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
}

func (s DescribeCloudSiemAssetsCounterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsCounterResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsCounterResponseBodyData) SetAssetNum(v int32) *DescribeCloudSiemAssetsCounterResponseBodyData {
	s.AssetNum = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBodyData) SetAssetType(v string) *DescribeCloudSiemAssetsCounterResponseBodyData {
	s.AssetType = &v
	return s
}

type DescribeCloudSiemAssetsCounterResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCloudSiemAssetsCounterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCloudSiemAssetsCounterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemAssetsCounterResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsCounterResponse) SetHeaders(v map[string]*string) *DescribeCloudSiemAssetsCounterResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponse) SetStatusCode(v int32) *DescribeCloudSiemAssetsCounterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponse) SetBody(v *DescribeCloudSiemAssetsCounterResponseBody) *DescribeCloudSiemAssetsCounterResponse {
	s.Body = v
	return s
}

type DescribeCloudSiemEventDetailRequest struct {
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudSiemEventDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventDetailRequest) SetIncidentUuid(v string) *DescribeCloudSiemEventDetailRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemEventDetailRequest) SetRegionId(v string) *DescribeCloudSiemEventDetailRequest {
	s.RegionId = &v
	return s
}

type DescribeCloudSiemEventDetailResponseBody struct {
	Code      *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeCloudSiemEventDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudSiemEventDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventDetailResponseBody) SetCode(v int32) *DescribeCloudSiemEventDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBody) SetData(v *DescribeCloudSiemEventDetailResponseBodyData) *DescribeCloudSiemEventDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBody) SetMessage(v string) *DescribeCloudSiemEventDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBody) SetRequestId(v string) *DescribeCloudSiemEventDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBody) SetSuccess(v bool) *DescribeCloudSiemEventDetailResponseBody {
	s.Success = &v
	return s
}

type DescribeCloudSiemEventDetailResponseBodyData struct {
	AlertNum       *int32    `json:"AlertNum,omitempty" xml:"AlertNum,omitempty"`
	Aliuid         *int64    `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	AssetNum       *int32    `json:"AssetNum,omitempty" xml:"AssetNum,omitempty"`
	AttCkLabels    []*string `json:"AttCkLabels,omitempty" xml:"AttCkLabels,omitempty" type:"Repeated"`
	DataSources    []*string `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	Description    *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DescriptionEn  *string   `json:"DescriptionEn,omitempty" xml:"DescriptionEn,omitempty"`
	ExtContent     *string   `json:"ExtContent,omitempty" xml:"ExtContent,omitempty"`
	GmtCreate      *string   `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified    *string   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	IncidentName   *string   `json:"IncidentName,omitempty" xml:"IncidentName,omitempty"`
	IncidentNameEn *string   `json:"IncidentNameEn,omitempty" xml:"IncidentNameEn,omitempty"`
	IncidentUuid   *string   `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	Remark         *string   `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status         *int32    `json:"Status,omitempty" xml:"Status,omitempty"`
	ThreatLevel    *string   `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	ThreatScore    *float32  `json:"ThreatScore,omitempty" xml:"ThreatScore,omitempty"`
}

func (s DescribeCloudSiemEventDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetAlertNum(v int32) *DescribeCloudSiemEventDetailResponseBodyData {
	s.AlertNum = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetAliuid(v int64) *DescribeCloudSiemEventDetailResponseBodyData {
	s.Aliuid = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetAssetNum(v int32) *DescribeCloudSiemEventDetailResponseBodyData {
	s.AssetNum = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetAttCkLabels(v []*string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.AttCkLabels = v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetDataSources(v []*string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.DataSources = v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetDescription(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetDescriptionEn(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.DescriptionEn = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetExtContent(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.ExtContent = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetGmtCreate(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetGmtModified(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetIncidentName(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.IncidentName = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetIncidentNameEn(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.IncidentNameEn = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetIncidentUuid(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetRemark(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.Remark = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetStatus(v int32) *DescribeCloudSiemEventDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetThreatLevel(v string) *DescribeCloudSiemEventDetailResponseBodyData {
	s.ThreatLevel = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponseBodyData) SetThreatScore(v float32) *DescribeCloudSiemEventDetailResponseBodyData {
	s.ThreatScore = &v
	return s
}

type DescribeCloudSiemEventDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCloudSiemEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCloudSiemEventDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventDetailResponse) SetHeaders(v map[string]*string) *DescribeCloudSiemEventDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudSiemEventDetailResponse) SetStatusCode(v int32) *DescribeCloudSiemEventDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponse) SetBody(v *DescribeCloudSiemEventDetailResponseBody) *DescribeCloudSiemEventDetailResponse {
	s.Body = v
	return s
}

type DescribeCloudSiemEventsRequest struct {
	AssetId      *string   `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	CurrentPage  *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime      *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventName    *string   `json:"EventName,omitempty" xml:"EventName,omitempty"`
	IncidentUuid *string   `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	Order        *string   `json:"Order,omitempty" xml:"Order,omitempty"`
	OrderField   *string   `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime    *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *int32    `json:"Status,omitempty" xml:"Status,omitempty"`
	ThreadLevel  []*string `json:"ThreadLevel,omitempty" xml:"ThreadLevel,omitempty" type:"Repeated"`
}

func (s DescribeCloudSiemEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsRequest) SetAssetId(v string) *DescribeCloudSiemEventsRequest {
	s.AssetId = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetCurrentPage(v int32) *DescribeCloudSiemEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetEndTime(v int64) *DescribeCloudSiemEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetEventName(v string) *DescribeCloudSiemEventsRequest {
	s.EventName = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetIncidentUuid(v string) *DescribeCloudSiemEventsRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetOrder(v string) *DescribeCloudSiemEventsRequest {
	s.Order = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetOrderField(v string) *DescribeCloudSiemEventsRequest {
	s.OrderField = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetPageSize(v int32) *DescribeCloudSiemEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetRegionId(v string) *DescribeCloudSiemEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetStartTime(v int64) *DescribeCloudSiemEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetStatus(v int32) *DescribeCloudSiemEventsRequest {
	s.Status = &v
	return s
}

func (s *DescribeCloudSiemEventsRequest) SetThreadLevel(v []*string) *DescribeCloudSiemEventsRequest {
	s.ThreadLevel = v
	return s
}

type DescribeCloudSiemEventsResponseBody struct {
	Code      *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeCloudSiemEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudSiemEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponseBody) SetCode(v int32) *DescribeCloudSiemEventsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBody) SetData(v *DescribeCloudSiemEventsResponseBodyData) *DescribeCloudSiemEventsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudSiemEventsResponseBody) SetMessage(v string) *DescribeCloudSiemEventsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBody) SetRequestId(v string) *DescribeCloudSiemEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBody) SetSuccess(v bool) *DescribeCloudSiemEventsResponseBody {
	s.Success = &v
	return s
}

type DescribeCloudSiemEventsResponseBodyData struct {
	PageInfo     *DescribeCloudSiemEventsResponseBodyDataPageInfo       `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	ResponseData []*DescribeCloudSiemEventsResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeCloudSiemEventsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponseBodyData) SetPageInfo(v *DescribeCloudSiemEventsResponseBodyDataPageInfo) *DescribeCloudSiemEventsResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyData) SetResponseData(v []*DescribeCloudSiemEventsResponseBodyDataResponseData) *DescribeCloudSiemEventsResponseBodyData {
	s.ResponseData = v
	return s
}

type DescribeCloudSiemEventsResponseBodyDataPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudSiemEventsResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventsResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeCloudSiemEventsResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeCloudSiemEventsResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeCloudSiemEventsResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeCloudSiemEventsResponseBodyDataResponseData struct {
	AlertNum       *int32    `json:"AlertNum,omitempty" xml:"AlertNum,omitempty"`
	Aliuid         *int64    `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	AssetNum       *int32    `json:"AssetNum,omitempty" xml:"AssetNum,omitempty"`
	AttCkLabels    []*string `json:"AttCkLabels,omitempty" xml:"AttCkLabels,omitempty" type:"Repeated"`
	DataSources    []*string `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	Description    *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DescriptionEn  *string   `json:"DescriptionEn,omitempty" xml:"DescriptionEn,omitempty"`
	ExtContent     *string   `json:"ExtContent,omitempty" xml:"ExtContent,omitempty"`
	GmtCreate      *string   `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified    *string   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	IncidentName   *string   `json:"IncidentName,omitempty" xml:"IncidentName,omitempty"`
	IncidentNameEn *string   `json:"IncidentNameEn,omitempty" xml:"IncidentNameEn,omitempty"`
	IncidentUuid   *string   `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	Remark         *string   `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status         *int32    `json:"Status,omitempty" xml:"Status,omitempty"`
	ThreatLevel    *string   `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	ThreatScore    *float32  `json:"ThreatScore,omitempty" xml:"ThreatScore,omitempty"`
}

func (s DescribeCloudSiemEventsResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventsResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetAlertNum(v int32) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.AlertNum = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetAliuid(v int64) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetAssetNum(v int32) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.AssetNum = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetAttCkLabels(v []*string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.AttCkLabels = v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetDataSources(v []*string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.DataSources = v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetDescription(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.Description = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetDescriptionEn(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.DescriptionEn = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetExtContent(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.ExtContent = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetGmtCreate(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetGmtModified(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetIncidentName(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.IncidentName = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetIncidentNameEn(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.IncidentNameEn = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetIncidentUuid(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetRemark(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.Remark = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetStatus(v int32) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetThreatLevel(v string) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.ThreatLevel = &v
	return s
}

func (s *DescribeCloudSiemEventsResponseBodyDataResponseData) SetThreatScore(v float32) *DescribeCloudSiemEventsResponseBodyDataResponseData {
	s.ThreatScore = &v
	return s
}

type DescribeCloudSiemEventsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCloudSiemEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCloudSiemEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudSiemEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponse) SetHeaders(v map[string]*string) *DescribeCloudSiemEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudSiemEventsResponse) SetStatusCode(v int32) *DescribeCloudSiemEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudSiemEventsResponse) SetBody(v *DescribeCloudSiemEventsResponseBody) *DescribeCloudSiemEventsResponse {
	s.Body = v
	return s
}

type DescribeCustomizeRuleRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleId   *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeCustomizeRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleRequest) SetRegionId(v string) *DescribeCustomizeRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomizeRuleRequest) SetRuleId(v int64) *DescribeCustomizeRuleRequest {
	s.RuleId = &v
	return s
}

type DescribeCustomizeRuleResponseBody struct {
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeCustomizeRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomizeRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleResponseBody) SetCode(v int32) *DescribeCustomizeRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBody) SetData(v *DescribeCustomizeRuleResponseBodyData) *DescribeCustomizeRuleResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCustomizeRuleResponseBody) SetMessage(v string) *DescribeCustomizeRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBody) SetRequestId(v string) *DescribeCustomizeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBody) SetSuccess(v bool) *DescribeCustomizeRuleResponseBody {
	s.Success = &v
	return s
}

type DescribeCustomizeRuleResponseBodyData struct {
	AlertType           *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AlertTypeMds        *string `json:"AlertTypeMds,omitempty" xml:"AlertTypeMds,omitempty"`
	Aliuid              *int64  `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	EventTransferExt    *string `json:"EventTransferExt,omitempty" xml:"EventTransferExt,omitempty"`
	EventTransferSwitch *int32  `json:"EventTransferSwitch,omitempty" xml:"EventTransferSwitch,omitempty"`
	EventTransferType   *string `json:"EventTransferType,omitempty" xml:"EventTransferType,omitempty"`
	GmtCreate           *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified         *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                  *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LogSource           *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	LogSourceMds        *string `json:"LogSourceMds,omitempty" xml:"LogSourceMds,omitempty"`
	LogType             *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	LogTypeMds          *string `json:"LogTypeMds,omitempty" xml:"LogTypeMds,omitempty"`
	QueryCycle          *string `json:"QueryCycle,omitempty" xml:"QueryCycle,omitempty"`
	RuleCondition       *string `json:"RuleCondition,omitempty" xml:"RuleCondition,omitempty"`
	RuleDesc            *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	RuleGroup           *string `json:"RuleGroup,omitempty" xml:"RuleGroup,omitempty"`
	RuleName            *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleThreshold       *string `json:"RuleThreshold,omitempty" xml:"RuleThreshold,omitempty"`
	RuleType            *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Status              *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ThreatLevel         *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
}

func (s DescribeCustomizeRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleResponseBodyData) SetAlertType(v string) *DescribeCustomizeRuleResponseBodyData {
	s.AlertType = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetAlertTypeMds(v string) *DescribeCustomizeRuleResponseBodyData {
	s.AlertTypeMds = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetAliuid(v int64) *DescribeCustomizeRuleResponseBodyData {
	s.Aliuid = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetEventTransferExt(v string) *DescribeCustomizeRuleResponseBodyData {
	s.EventTransferExt = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetEventTransferSwitch(v int32) *DescribeCustomizeRuleResponseBodyData {
	s.EventTransferSwitch = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetEventTransferType(v string) *DescribeCustomizeRuleResponseBodyData {
	s.EventTransferType = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetGmtCreate(v string) *DescribeCustomizeRuleResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetGmtModified(v string) *DescribeCustomizeRuleResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetId(v int64) *DescribeCustomizeRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetLogSource(v string) *DescribeCustomizeRuleResponseBodyData {
	s.LogSource = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetLogSourceMds(v string) *DescribeCustomizeRuleResponseBodyData {
	s.LogSourceMds = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetLogType(v string) *DescribeCustomizeRuleResponseBodyData {
	s.LogType = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetLogTypeMds(v string) *DescribeCustomizeRuleResponseBodyData {
	s.LogTypeMds = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetQueryCycle(v string) *DescribeCustomizeRuleResponseBodyData {
	s.QueryCycle = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetRuleCondition(v string) *DescribeCustomizeRuleResponseBodyData {
	s.RuleCondition = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetRuleDesc(v string) *DescribeCustomizeRuleResponseBodyData {
	s.RuleDesc = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetRuleGroup(v string) *DescribeCustomizeRuleResponseBodyData {
	s.RuleGroup = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetRuleName(v string) *DescribeCustomizeRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetRuleThreshold(v string) *DescribeCustomizeRuleResponseBodyData {
	s.RuleThreshold = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetRuleType(v string) *DescribeCustomizeRuleResponseBodyData {
	s.RuleType = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetStatus(v int32) *DescribeCustomizeRuleResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeCustomizeRuleResponseBodyData) SetThreatLevel(v string) *DescribeCustomizeRuleResponseBodyData {
	s.ThreatLevel = &v
	return s
}

type DescribeCustomizeRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCustomizeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCustomizeRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleResponse) SetHeaders(v map[string]*string) *DescribeCustomizeRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizeRuleResponse) SetStatusCode(v int32) *DescribeCustomizeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizeRuleResponse) SetBody(v *DescribeCustomizeRuleResponseBody) *DescribeCustomizeRuleResponse {
	s.Body = v
	return s
}

type DescribeCustomizeRuleCountRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCustomizeRuleCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleCountRequest) SetRegionId(v string) *DescribeCustomizeRuleCountRequest {
	s.RegionId = &v
	return s
}

type DescribeCustomizeRuleCountResponseBody struct {
	Code      *int32                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeCustomizeRuleCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomizeRuleCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleCountResponseBody) SetCode(v int32) *DescribeCustomizeRuleCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBody) SetData(v *DescribeCustomizeRuleCountResponseBodyData) *DescribeCustomizeRuleCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBody) SetMessage(v string) *DescribeCustomizeRuleCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBody) SetRequestId(v string) *DescribeCustomizeRuleCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBody) SetSuccess(v bool) *DescribeCustomizeRuleCountResponseBody {
	s.Success = &v
	return s
}

type DescribeCustomizeRuleCountResponseBodyData struct {
	HighRuleNum   *int32 `json:"HighRuleNum,omitempty" xml:"HighRuleNum,omitempty"`
	InUseRuleNum  *int32 `json:"InUseRuleNum,omitempty" xml:"InUseRuleNum,omitempty"`
	LowRuleNum    *int32 `json:"LowRuleNum,omitempty" xml:"LowRuleNum,omitempty"`
	MediumRuleNum *int32 `json:"MediumRuleNum,omitempty" xml:"MediumRuleNum,omitempty"`
}

func (s DescribeCustomizeRuleCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetHighRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.HighRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetInUseRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.InUseRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetLowRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.LowRuleNum = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponseBodyData) SetMediumRuleNum(v int32) *DescribeCustomizeRuleCountResponseBodyData {
	s.MediumRuleNum = &v
	return s
}

type DescribeCustomizeRuleCountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCustomizeRuleCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCustomizeRuleCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleCountResponse) SetHeaders(v map[string]*string) *DescribeCustomizeRuleCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizeRuleCountResponse) SetStatusCode(v int32) *DescribeCustomizeRuleCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponse) SetBody(v *DescribeCustomizeRuleCountResponseBody) *DescribeCustomizeRuleCountResponse {
	s.Body = v
	return s
}

type DescribeCustomizeRuleTestRequest struct {
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCustomizeRuleTestRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleTestRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestRequest) SetId(v int64) *DescribeCustomizeRuleTestRequest {
	s.Id = &v
	return s
}

func (s *DescribeCustomizeRuleTestRequest) SetRegionId(v string) *DescribeCustomizeRuleTestRequest {
	s.RegionId = &v
	return s
}

type DescribeCustomizeRuleTestResponseBody struct {
	Code      *int32                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeCustomizeRuleTestResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomizeRuleTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleTestResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestResponseBody) SetCode(v int32) *DescribeCustomizeRuleTestResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBody) SetData(v *DescribeCustomizeRuleTestResponseBodyData) *DescribeCustomizeRuleTestResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBody) SetMessage(v string) *DescribeCustomizeRuleTestResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBody) SetRequestId(v string) *DescribeCustomizeRuleTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBody) SetSuccess(v bool) *DescribeCustomizeRuleTestResponseBody {
	s.Success = &v
	return s
}

type DescribeCustomizeRuleTestResponseBodyData struct {
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	SimulateData *string `json:"SimulateData,omitempty" xml:"SimulateData,omitempty"`
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCustomizeRuleTestResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleTestResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestResponseBodyData) SetId(v int64) *DescribeCustomizeRuleTestResponseBodyData {
	s.Id = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBodyData) SetSimulateData(v string) *DescribeCustomizeRuleTestResponseBodyData {
	s.SimulateData = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponseBodyData) SetStatus(v int32) *DescribeCustomizeRuleTestResponseBodyData {
	s.Status = &v
	return s
}

type DescribeCustomizeRuleTestResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCustomizeRuleTestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCustomizeRuleTestResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleTestResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestResponse) SetHeaders(v map[string]*string) *DescribeCustomizeRuleTestResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizeRuleTestResponse) SetStatusCode(v int32) *DescribeCustomizeRuleTestResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizeRuleTestResponse) SetBody(v *DescribeCustomizeRuleTestResponseBody) *DescribeCustomizeRuleTestResponse {
	s.Body = v
	return s
}

type DescribeCustomizeRuleTestHistogramRequest struct {
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCustomizeRuleTestHistogramRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleTestHistogramRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestHistogramRequest) SetId(v int64) *DescribeCustomizeRuleTestHistogramRequest {
	s.Id = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramRequest) SetRegionId(v string) *DescribeCustomizeRuleTestHistogramRequest {
	s.RegionId = &v
	return s
}

type DescribeCustomizeRuleTestHistogramResponseBody struct {
	Code      *int32                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeCustomizeRuleTestHistogramResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomizeRuleTestHistogramResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleTestHistogramResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) SetCode(v int32) *DescribeCustomizeRuleTestHistogramResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) SetData(v []*DescribeCustomizeRuleTestHistogramResponseBodyData) *DescribeCustomizeRuleTestHistogramResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) SetMessage(v string) *DescribeCustomizeRuleTestHistogramResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) SetRequestId(v string) *DescribeCustomizeRuleTestHistogramResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBody) SetSuccess(v bool) *DescribeCustomizeRuleTestHistogramResponseBody {
	s.Success = &v
	return s
}

type DescribeCustomizeRuleTestHistogramResponseBodyData struct {
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	From  *int64 `json:"From,omitempty" xml:"From,omitempty"`
	To    *int64 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s DescribeCustomizeRuleTestHistogramResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleTestHistogramResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestHistogramResponseBodyData) SetCount(v int64) *DescribeCustomizeRuleTestHistogramResponseBodyData {
	s.Count = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBodyData) SetFrom(v int64) *DescribeCustomizeRuleTestHistogramResponseBodyData {
	s.From = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponseBodyData) SetTo(v int64) *DescribeCustomizeRuleTestHistogramResponseBodyData {
	s.To = &v
	return s
}

type DescribeCustomizeRuleTestHistogramResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCustomizeRuleTestHistogramResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCustomizeRuleTestHistogramResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomizeRuleTestHistogramResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestHistogramResponse) SetHeaders(v map[string]*string) *DescribeCustomizeRuleTestHistogramResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponse) SetStatusCode(v int32) *DescribeCustomizeRuleTestHistogramResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponse) SetBody(v *DescribeCustomizeRuleTestHistogramResponseBody) *DescribeCustomizeRuleTestHistogramResponse {
	s.Body = v
	return s
}

type DescribeDisposeAndPlaybookRequest struct {
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EntityType   *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDisposeAndPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeAndPlaybookRequest) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookRequest) SetCurrentPage(v int32) *DescribeDisposeAndPlaybookRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetEntityType(v string) *DescribeDisposeAndPlaybookRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetIncidentUuid(v string) *DescribeDisposeAndPlaybookRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetPageSize(v int32) *DescribeDisposeAndPlaybookRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDisposeAndPlaybookRequest) SetRegionId(v string) *DescribeDisposeAndPlaybookRequest {
	s.RegionId = &v
	return s
}

type DescribeDisposeAndPlaybookResponseBody struct {
	Code      *int32                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeDisposeAndPlaybookResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDisposeAndPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponseBody) SetCode(v int32) *DescribeDisposeAndPlaybookResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBody) SetData(v *DescribeDisposeAndPlaybookResponseBodyData) *DescribeDisposeAndPlaybookResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBody) SetMessage(v string) *DescribeDisposeAndPlaybookResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBody) SetRequestId(v string) *DescribeDisposeAndPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBody) SetSuccess(v bool) *DescribeDisposeAndPlaybookResponseBody {
	s.Success = &v
	return s
}

type DescribeDisposeAndPlaybookResponseBodyData struct {
	PageInfo     *DescribeDisposeAndPlaybookResponseBodyDataPageInfo       `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	ResponseData []*DescribeDisposeAndPlaybookResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeDisposeAndPlaybookResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponseBodyData) SetPageInfo(v *DescribeDisposeAndPlaybookResponseBodyDataPageInfo) *DescribeDisposeAndPlaybookResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyData) SetResponseData(v []*DescribeDisposeAndPlaybookResponseBodyDataResponseData) *DescribeDisposeAndPlaybookResponseBodyData {
	s.ResponseData = v
	return s
}

type DescribeDisposeAndPlaybookResponseBodyDataPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDisposeAndPlaybookResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeDisposeAndPlaybookResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeDisposeAndPlaybookResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeDisposeAndPlaybookResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeDisposeAndPlaybookResponseBodyDataResponseData struct {
	AlertNum     *int32                                                                `json:"AlertNum,omitempty" xml:"AlertNum,omitempty"`
	Dispose      *string                                                               `json:"Dispose,omitempty" xml:"Dispose,omitempty"`
	EntityId     *int64                                                                `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityInfo   map[string]interface{}                                                `json:"EntityInfo,omitempty" xml:"EntityInfo,omitempty"`
	OpcodeMap    map[string]*string                                                    `json:"OpcodeMap,omitempty" xml:"OpcodeMap,omitempty"`
	OpcodeSet    []*string                                                             `json:"OpcodeSet,omitempty" xml:"OpcodeSet,omitempty" type:"Repeated"`
	PlaybookList []*DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList `json:"PlaybookList,omitempty" xml:"PlaybookList,omitempty" type:"Repeated"`
	Scope        []interface{}                                                         `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
}

func (s DescribeDisposeAndPlaybookResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetAlertNum(v int32) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.AlertNum = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetDispose(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.Dispose = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetEntityId(v int64) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.EntityId = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetEntityInfo(v map[string]interface{}) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.EntityInfo = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetOpcodeMap(v map[string]*string) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.OpcodeMap = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetOpcodeSet(v []*string) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.OpcodeSet = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetPlaybookList(v []*DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.PlaybookList = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseData) SetScope(v []interface{}) *DescribeDisposeAndPlaybookResponseBodyDataResponseData {
	s.Scope = v
	return s
}

type DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OpCode      *string `json:"OpCode,omitempty" xml:"OpCode,omitempty"`
	OpLevel     *string `json:"OpLevel,omitempty" xml:"OpLevel,omitempty"`
	TaskConfig  *string `json:"TaskConfig,omitempty" xml:"TaskConfig,omitempty"`
	WafPlaybook *bool   `json:"WafPlaybook,omitempty" xml:"WafPlaybook,omitempty"`
}

func (s DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetDescription(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.Description = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetDisplayName(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.DisplayName = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetName(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.Name = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetOpCode(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.OpCode = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetOpLevel(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.OpLevel = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetTaskConfig(v string) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.TaskConfig = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList) SetWafPlaybook(v bool) *DescribeDisposeAndPlaybookResponseBodyDataResponseDataPlaybookList {
	s.WafPlaybook = &v
	return s
}

type DescribeDisposeAndPlaybookResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDisposeAndPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDisposeAndPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeAndPlaybookResponse) GoString() string {
	return s.String()
}

func (s *DescribeDisposeAndPlaybookResponse) SetHeaders(v map[string]*string) *DescribeDisposeAndPlaybookResponse {
	s.Headers = v
	return s
}

func (s *DescribeDisposeAndPlaybookResponse) SetStatusCode(v int32) *DescribeDisposeAndPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDisposeAndPlaybookResponse) SetBody(v *DescribeDisposeAndPlaybookResponseBody) *DescribeDisposeAndPlaybookResponse {
	s.Body = v
	return s
}

type DescribeDisposeStrategyPlaybookRequest struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDisposeStrategyPlaybookRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeStrategyPlaybookRequest) GoString() string {
	return s.String()
}

func (s *DescribeDisposeStrategyPlaybookRequest) SetEndTime(v int64) *DescribeDisposeStrategyPlaybookRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookRequest) SetRegionId(v string) *DescribeDisposeStrategyPlaybookRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookRequest) SetStartTime(v int64) *DescribeDisposeStrategyPlaybookRequest {
	s.StartTime = &v
	return s
}

type DescribeDisposeStrategyPlaybookResponseBody struct {
	Code      *int32                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeDisposeStrategyPlaybookResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDisposeStrategyPlaybookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeStrategyPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) SetCode(v int32) *DescribeDisposeStrategyPlaybookResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) SetData(v []*DescribeDisposeStrategyPlaybookResponseBodyData) *DescribeDisposeStrategyPlaybookResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) SetMessage(v string) *DescribeDisposeStrategyPlaybookResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) SetRequestId(v string) *DescribeDisposeStrategyPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) SetSuccess(v bool) *DescribeDisposeStrategyPlaybookResponseBody {
	s.Success = &v
	return s
}

type DescribeDisposeStrategyPlaybookResponseBodyData struct {
	PlaybookName *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeDisposeStrategyPlaybookResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeStrategyPlaybookResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDisposeStrategyPlaybookResponseBodyData) SetPlaybookName(v string) *DescribeDisposeStrategyPlaybookResponseBodyData {
	s.PlaybookName = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBodyData) SetPlaybookUuid(v string) *DescribeDisposeStrategyPlaybookResponseBodyData {
	s.PlaybookUuid = &v
	return s
}

type DescribeDisposeStrategyPlaybookResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDisposeStrategyPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDisposeStrategyPlaybookResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisposeStrategyPlaybookResponse) GoString() string {
	return s.String()
}

func (s *DescribeDisposeStrategyPlaybookResponse) SetHeaders(v map[string]*string) *DescribeDisposeStrategyPlaybookResponse {
	s.Headers = v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponse) SetStatusCode(v int32) *DescribeDisposeStrategyPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponse) SetBody(v *DescribeDisposeStrategyPlaybookResponseBody) *DescribeDisposeStrategyPlaybookResponse {
	s.Body = v
	return s
}

type DescribeEntityInfoRequest struct {
	EntityId       *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityIdentity *string `json:"EntityIdentity,omitempty" xml:"EntityIdentity,omitempty"`
	IncidentUuid   *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SophonTaskId   *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
}

func (s DescribeEntityInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEntityInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeEntityInfoRequest) SetEntityId(v int64) *DescribeEntityInfoRequest {
	s.EntityId = &v
	return s
}

func (s *DescribeEntityInfoRequest) SetEntityIdentity(v string) *DescribeEntityInfoRequest {
	s.EntityIdentity = &v
	return s
}

func (s *DescribeEntityInfoRequest) SetIncidentUuid(v string) *DescribeEntityInfoRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeEntityInfoRequest) SetRegionId(v string) *DescribeEntityInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEntityInfoRequest) SetSophonTaskId(v string) *DescribeEntityInfoRequest {
	s.SophonTaskId = &v
	return s
}

type DescribeEntityInfoResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeEntityInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEntityInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEntityInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEntityInfoResponseBody) SetCode(v int32) *DescribeEntityInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEntityInfoResponseBody) SetData(v *DescribeEntityInfoResponseBodyData) *DescribeEntityInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEntityInfoResponseBody) SetMessage(v string) *DescribeEntityInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEntityInfoResponseBody) SetRequestId(v string) *DescribeEntityInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEntityInfoResponseBody) SetSuccess(v bool) *DescribeEntityInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeEntityInfoResponseBodyData struct {
	EntityId   *int64                 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityInfo map[string]interface{} `json:"EntityInfo,omitempty" xml:"EntityInfo,omitempty"`
	EntityType *string                `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	TipInfo    map[string]interface{} `json:"TipInfo,omitempty" xml:"TipInfo,omitempty"`
}

func (s DescribeEntityInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEntityInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEntityInfoResponseBodyData) SetEntityId(v int64) *DescribeEntityInfoResponseBodyData {
	s.EntityId = &v
	return s
}

func (s *DescribeEntityInfoResponseBodyData) SetEntityInfo(v map[string]interface{}) *DescribeEntityInfoResponseBodyData {
	s.EntityInfo = v
	return s
}

func (s *DescribeEntityInfoResponseBodyData) SetEntityType(v string) *DescribeEntityInfoResponseBodyData {
	s.EntityType = &v
	return s
}

func (s *DescribeEntityInfoResponseBodyData) SetTipInfo(v map[string]interface{}) *DescribeEntityInfoResponseBodyData {
	s.TipInfo = v
	return s
}

type DescribeEntityInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEntityInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEntityInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEntityInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeEntityInfoResponse) SetHeaders(v map[string]*string) *DescribeEntityInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeEntityInfoResponse) SetStatusCode(v int32) *DescribeEntityInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEntityInfoResponse) SetBody(v *DescribeEntityInfoResponseBody) *DescribeEntityInfoResponse {
	s.Body = v
	return s
}

type DescribeEventCountByThreatLevelRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEventCountByThreatLevelRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventCountByThreatLevelRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventCountByThreatLevelRequest) SetRegionId(v string) *DescribeEventCountByThreatLevelRequest {
	s.RegionId = &v
	return s
}

type DescribeEventCountByThreatLevelResponseBody struct {
	Code      *int32                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeEventCountByThreatLevelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEventCountByThreatLevelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventCountByThreatLevelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventCountByThreatLevelResponseBody) SetCode(v int32) *DescribeEventCountByThreatLevelResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBody) SetData(v *DescribeEventCountByThreatLevelResponseBodyData) *DescribeEventCountByThreatLevelResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBody) SetMessage(v string) *DescribeEventCountByThreatLevelResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBody) SetRequestId(v string) *DescribeEventCountByThreatLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBody) SetSuccess(v bool) *DescribeEventCountByThreatLevelResponseBody {
	s.Success = &v
	return s
}

type DescribeEventCountByThreatLevelResponseBodyData struct {
	EventNum            *int64 `json:"EventNum,omitempty" xml:"EventNum,omitempty"`
	HighLevelEventNum   *int64 `json:"HighLevelEventNum,omitempty" xml:"HighLevelEventNum,omitempty"`
	LowLevelEventNum    *int64 `json:"LowLevelEventNum,omitempty" xml:"LowLevelEventNum,omitempty"`
	MediumLevelEventNum *int64 `json:"MediumLevelEventNum,omitempty" xml:"MediumLevelEventNum,omitempty"`
	UndealEventNum      *int64 `json:"UndealEventNum,omitempty" xml:"UndealEventNum,omitempty"`
}

func (s DescribeEventCountByThreatLevelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventCountByThreatLevelResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.EventNum = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetHighLevelEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.HighLevelEventNum = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetLowLevelEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.LowLevelEventNum = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetMediumLevelEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.MediumLevelEventNum = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetUndealEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.UndealEventNum = &v
	return s
}

type DescribeEventCountByThreatLevelResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEventCountByThreatLevelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEventCountByThreatLevelResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventCountByThreatLevelResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventCountByThreatLevelResponse) SetHeaders(v map[string]*string) *DescribeEventCountByThreatLevelResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventCountByThreatLevelResponse) SetStatusCode(v int32) *DescribeEventCountByThreatLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponse) SetBody(v *DescribeEventCountByThreatLevelResponseBody) *DescribeEventCountByThreatLevelResponse {
	s.Body = v
	return s
}

type DescribeEventDisposeRequest struct {
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEventDisposeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDisposeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventDisposeRequest) SetCurrentPage(v int32) *DescribeEventDisposeRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventDisposeRequest) SetIncidentUuid(v string) *DescribeEventDisposeRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeEventDisposeRequest) SetPageSize(v int32) *DescribeEventDisposeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventDisposeRequest) SetRegionId(v string) *DescribeEventDisposeRequest {
	s.RegionId = &v
	return s
}

type DescribeEventDisposeResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeEventDisposeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEventDisposeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDisposeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventDisposeResponseBody) SetCode(v int32) *DescribeEventDisposeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventDisposeResponseBody) SetData(v *DescribeEventDisposeResponseBodyData) *DescribeEventDisposeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEventDisposeResponseBody) SetMessage(v string) *DescribeEventDisposeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventDisposeResponseBody) SetRequestId(v string) *DescribeEventDisposeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventDisposeResponseBody) SetSuccess(v bool) *DescribeEventDisposeResponseBody {
	s.Success = &v
	return s
}

type DescribeEventDisposeResponseBodyData struct {
	EventDispose []interface{}                                     `json:"EventDispose,omitempty" xml:"EventDispose,omitempty" type:"Repeated"`
	ReceiverInfo *DescribeEventDisposeResponseBodyDataReceiverInfo `json:"ReceiverInfo,omitempty" xml:"ReceiverInfo,omitempty" type:"Struct"`
	Remark       *string                                           `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status       *int32                                            `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventDisposeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDisposeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEventDisposeResponseBodyData) SetEventDispose(v []interface{}) *DescribeEventDisposeResponseBodyData {
	s.EventDispose = v
	return s
}

func (s *DescribeEventDisposeResponseBodyData) SetReceiverInfo(v *DescribeEventDisposeResponseBodyDataReceiverInfo) *DescribeEventDisposeResponseBodyData {
	s.ReceiverInfo = v
	return s
}

func (s *DescribeEventDisposeResponseBodyData) SetRemark(v string) *DescribeEventDisposeResponseBodyData {
	s.Remark = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyData) SetStatus(v int32) *DescribeEventDisposeResponseBodyData {
	s.Status = &v
	return s
}

type DescribeEventDisposeResponseBodyDataReceiverInfo struct {
	Channel      *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	GmtCreate    *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	MessageTitle *string `json:"MessageTitle,omitempty" xml:"MessageTitle,omitempty"`
	Receiver     *string `json:"Receiver,omitempty" xml:"Receiver,omitempty"`
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEventDisposeResponseBodyDataReceiverInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDisposeResponseBodyDataReceiverInfo) GoString() string {
	return s.String()
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetChannel(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.Channel = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetGmtCreate(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.GmtCreate = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetGmtModified(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.GmtModified = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetId(v int64) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.Id = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetIncidentUuid(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetMessageTitle(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.MessageTitle = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetReceiver(v string) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.Receiver = &v
	return s
}

func (s *DescribeEventDisposeResponseBodyDataReceiverInfo) SetStatus(v int32) *DescribeEventDisposeResponseBodyDataReceiverInfo {
	s.Status = &v
	return s
}

type DescribeEventDisposeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEventDisposeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEventDisposeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventDisposeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventDisposeResponse) SetHeaders(v map[string]*string) *DescribeEventDisposeResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventDisposeResponse) SetStatusCode(v int32) *DescribeEventDisposeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventDisposeResponse) SetBody(v *DescribeEventDisposeResponseBody) *DescribeEventDisposeResponse {
	s.Body = v
	return s
}

type DescribeJobStatusRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubmitId *string `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
}

func (s DescribeJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobStatusRequest) SetRegionId(v string) *DescribeJobStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeJobStatusRequest) SetSubmitId(v string) *DescribeJobStatusRequest {
	s.SubmitId = &v
	return s
}

type DescribeJobStatusResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeJobStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode   *string                            `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobStatusResponseBody) SetCode(v int32) *DescribeJobStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeJobStatusResponseBody) SetData(v *DescribeJobStatusResponseBodyData) *DescribeJobStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeJobStatusResponseBody) SetErrCode(v string) *DescribeJobStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeJobStatusResponseBody) SetMessage(v string) *DescribeJobStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeJobStatusResponseBody) SetRequestId(v string) *DescribeJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobStatusResponseBody) SetSuccess(v bool) *DescribeJobStatusResponseBody {
	s.Success = &v
	return s
}

type DescribeJobStatusResponseBodyData struct {
	ConfigId    *string                                         `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ErrTaskList []*DescribeJobStatusResponseBodyDataErrTaskList `json:"ErrTaskList,omitempty" xml:"ErrTaskList,omitempty" type:"Repeated"`
	FailedCount *int32                                          `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	FinishCount *int32                                          `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	FolderId    *string                                         `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	TaskCount   *int32                                          `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
	TaskStatus  *string                                         `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s DescribeJobStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeJobStatusResponseBodyData) SetConfigId(v string) *DescribeJobStatusResponseBodyData {
	s.ConfigId = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetErrTaskList(v []*DescribeJobStatusResponseBodyDataErrTaskList) *DescribeJobStatusResponseBodyData {
	s.ErrTaskList = v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetFailedCount(v int32) *DescribeJobStatusResponseBodyData {
	s.FailedCount = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetFinishCount(v int32) *DescribeJobStatusResponseBodyData {
	s.FinishCount = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetFolderId(v string) *DescribeJobStatusResponseBodyData {
	s.FolderId = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetTaskCount(v int32) *DescribeJobStatusResponseBodyData {
	s.TaskCount = &v
	return s
}

func (s *DescribeJobStatusResponseBodyData) SetTaskStatus(v string) *DescribeJobStatusResponseBodyData {
	s.TaskStatus = &v
	return s
}

type DescribeJobStatusResponseBodyDataErrTaskList struct {
	ProductList []*DescribeJobStatusResponseBodyDataErrTaskListProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Repeated"`
	UserId      *int64                                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeJobStatusResponseBodyDataErrTaskList) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobStatusResponseBodyDataErrTaskList) GoString() string {
	return s.String()
}

func (s *DescribeJobStatusResponseBodyDataErrTaskList) SetProductList(v []*DescribeJobStatusResponseBodyDataErrTaskListProductList) *DescribeJobStatusResponseBodyDataErrTaskList {
	s.ProductList = v
	return s
}

func (s *DescribeJobStatusResponseBodyDataErrTaskList) SetUserId(v int64) *DescribeJobStatusResponseBodyDataErrTaskList {
	s.UserId = &v
	return s
}

type DescribeJobStatusResponseBodyDataErrTaskListProductList struct {
	LogList     []*DescribeJobStatusResponseBodyDataErrTaskListProductListLogList `json:"LogList,omitempty" xml:"LogList,omitempty" type:"Repeated"`
	ProductCode *string                                                           `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s DescribeJobStatusResponseBodyDataErrTaskListProductList) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobStatusResponseBodyDataErrTaskListProductList) GoString() string {
	return s.String()
}

func (s *DescribeJobStatusResponseBodyDataErrTaskListProductList) SetLogList(v []*DescribeJobStatusResponseBodyDataErrTaskListProductListLogList) *DescribeJobStatusResponseBodyDataErrTaskListProductList {
	s.LogList = v
	return s
}

func (s *DescribeJobStatusResponseBodyDataErrTaskListProductList) SetProductCode(v string) *DescribeJobStatusResponseBodyDataErrTaskListProductList {
	s.ProductCode = &v
	return s
}

type DescribeJobStatusResponseBodyDataErrTaskListProductListLogList struct {
	ErrorCode           *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	LogCode             *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	LogStoreNamePattern *string `json:"LogStoreNamePattern,omitempty" xml:"LogStoreNamePattern,omitempty"`
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProjectNamePattern  *string `json:"ProjectNamePattern,omitempty" xml:"ProjectNamePattern,omitempty"`
	RegionCode          *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s DescribeJobStatusResponseBodyDataErrTaskListProductListLogList) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobStatusResponseBodyDataErrTaskListProductListLogList) GoString() string {
	return s.String()
}

func (s *DescribeJobStatusResponseBodyDataErrTaskListProductListLogList) SetErrorCode(v string) *DescribeJobStatusResponseBodyDataErrTaskListProductListLogList {
	s.ErrorCode = &v
	return s
}

func (s *DescribeJobStatusResponseBodyDataErrTaskListProductListLogList) SetLogCode(v string) *DescribeJobStatusResponseBodyDataErrTaskListProductListLogList {
	s.LogCode = &v
	return s
}

func (s *DescribeJobStatusResponseBodyDataErrTaskListProductListLogList) SetLogStoreNamePattern(v string) *DescribeJobStatusResponseBodyDataErrTaskListProductListLogList {
	s.LogStoreNamePattern = &v
	return s
}

func (s *DescribeJobStatusResponseBodyDataErrTaskListProductListLogList) SetProductCode(v string) *DescribeJobStatusResponseBodyDataErrTaskListProductListLogList {
	s.ProductCode = &v
	return s
}

func (s *DescribeJobStatusResponseBodyDataErrTaskListProductListLogList) SetProjectNamePattern(v string) *DescribeJobStatusResponseBodyDataErrTaskListProductListLogList {
	s.ProjectNamePattern = &v
	return s
}

func (s *DescribeJobStatusResponseBodyDataErrTaskListProductListLogList) SetRegionCode(v string) *DescribeJobStatusResponseBodyDataErrTaskListProductListLogList {
	s.RegionCode = &v
	return s
}

type DescribeJobStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeJobStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobStatusResponse) SetHeaders(v map[string]*string) *DescribeJobStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobStatusResponse) SetStatusCode(v int32) *DescribeJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobStatusResponse) SetBody(v *DescribeJobStatusResponseBody) *DescribeJobStatusResponse {
	s.Body = v
	return s
}

type DescribeLogFieldsRequest struct {
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	LogType   *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLogFieldsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogFieldsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogFieldsRequest) SetLogSource(v string) *DescribeLogFieldsRequest {
	s.LogSource = &v
	return s
}

func (s *DescribeLogFieldsRequest) SetLogType(v string) *DescribeLogFieldsRequest {
	s.LogType = &v
	return s
}

func (s *DescribeLogFieldsRequest) SetRegionId(v string) *DescribeLogFieldsRequest {
	s.RegionId = &v
	return s
}

type DescribeLogFieldsResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeLogFieldsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeLogFieldsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogFieldsResponseBody) SetCode(v int32) *DescribeLogFieldsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogFieldsResponseBody) SetData(v []*DescribeLogFieldsResponseBodyData) *DescribeLogFieldsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLogFieldsResponseBody) SetMessage(v string) *DescribeLogFieldsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLogFieldsResponseBody) SetRequestId(v string) *DescribeLogFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogFieldsResponseBody) SetSuccess(v bool) *DescribeLogFieldsResponseBody {
	s.Success = &v
	return s
}

type DescribeLogFieldsResponseBodyData struct {
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	FieldDesc    *string `json:"FieldDesc,omitempty" xml:"FieldDesc,omitempty"`
	FieldName    *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	FieldType    *string `json:"FieldType,omitempty" xml:"FieldType,omitempty"`
	LogCode      *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
}

func (s DescribeLogFieldsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogFieldsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLogFieldsResponseBodyData) SetActivityName(v string) *DescribeLogFieldsResponseBodyData {
	s.ActivityName = &v
	return s
}

func (s *DescribeLogFieldsResponseBodyData) SetFieldDesc(v string) *DescribeLogFieldsResponseBodyData {
	s.FieldDesc = &v
	return s
}

func (s *DescribeLogFieldsResponseBodyData) SetFieldName(v string) *DescribeLogFieldsResponseBodyData {
	s.FieldName = &v
	return s
}

func (s *DescribeLogFieldsResponseBodyData) SetFieldType(v string) *DescribeLogFieldsResponseBodyData {
	s.FieldType = &v
	return s
}

func (s *DescribeLogFieldsResponseBodyData) SetLogCode(v string) *DescribeLogFieldsResponseBodyData {
	s.LogCode = &v
	return s
}

type DescribeLogFieldsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLogFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogFieldsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogFieldsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogFieldsResponse) SetHeaders(v map[string]*string) *DescribeLogFieldsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogFieldsResponse) SetStatusCode(v int32) *DescribeLogFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogFieldsResponse) SetBody(v *DescribeLogFieldsResponseBody) *DescribeLogFieldsResponse {
	s.Body = v
	return s
}

type DescribeLogSourceRequest struct {
	LogType  *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLogSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogSourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogSourceRequest) SetLogType(v string) *DescribeLogSourceRequest {
	s.LogType = &v
	return s
}

func (s *DescribeLogSourceRequest) SetRegionId(v string) *DescribeLogSourceRequest {
	s.RegionId = &v
	return s
}

type DescribeLogSourceResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeLogSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeLogSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogSourceResponseBody) SetCode(v int32) *DescribeLogSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogSourceResponseBody) SetData(v []*DescribeLogSourceResponseBodyData) *DescribeLogSourceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLogSourceResponseBody) SetMessage(v string) *DescribeLogSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLogSourceResponseBody) SetRequestId(v string) *DescribeLogSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogSourceResponseBody) SetSuccess(v bool) *DescribeLogSourceResponseBody {
	s.Success = &v
	return s
}

type DescribeLogSourceResponseBodyData struct {
	LogSource     *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	LogSourceName *string `json:"LogSourceName,omitempty" xml:"LogSourceName,omitempty"`
}

func (s DescribeLogSourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLogSourceResponseBodyData) SetLogSource(v string) *DescribeLogSourceResponseBodyData {
	s.LogSource = &v
	return s
}

func (s *DescribeLogSourceResponseBodyData) SetLogSourceName(v string) *DescribeLogSourceResponseBodyData {
	s.LogSourceName = &v
	return s
}

type DescribeLogSourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLogSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogSourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogSourceResponse) SetHeaders(v map[string]*string) *DescribeLogSourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogSourceResponse) SetStatusCode(v int32) *DescribeLogSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogSourceResponse) SetBody(v *DescribeLogSourceResponseBody) *DescribeLogSourceResponse {
	s.Body = v
	return s
}

type DescribeLogStoreRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLogStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreRequest) SetRegionId(v string) *DescribeLogStoreRequest {
	s.RegionId = &v
	return s
}

type DescribeLogStoreResponseBody struct {
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeLogStoreResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DyCode    *string                           `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string                           `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string                           `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeLogStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreResponseBody) SetCode(v int32) *DescribeLogStoreResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogStoreResponseBody) SetData(v *DescribeLogStoreResponseBodyData) *DescribeLogStoreResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLogStoreResponseBody) SetDyCode(v string) *DescribeLogStoreResponseBody {
	s.DyCode = &v
	return s
}

func (s *DescribeLogStoreResponseBody) SetDyMessage(v string) *DescribeLogStoreResponseBody {
	s.DyMessage = &v
	return s
}

func (s *DescribeLogStoreResponseBody) SetErrCode(v string) *DescribeLogStoreResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeLogStoreResponseBody) SetMessage(v string) *DescribeLogStoreResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLogStoreResponseBody) SetRequestId(v string) *DescribeLogStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogStoreResponseBody) SetSuccess(v bool) *DescribeLogStoreResponseBody {
	s.Success = &v
	return s
}

type DescribeLogStoreResponseBodyData struct {
	AppendMeta     *bool   `json:"AppendMeta,omitempty" xml:"AppendMeta,omitempty"`
	AutoSplit      *bool   `json:"AutoSplit,omitempty" xml:"AutoSplit,omitempty"`
	EnableTracking *bool   `json:"EnableTracking,omitempty" xml:"EnableTracking,omitempty"`
	LogStoreName   *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	MaxSplitShard  *int32  `json:"MaxSplitShard,omitempty" xml:"MaxSplitShard,omitempty"`
	ShardCount     *int32  `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	Ttl            *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s DescribeLogStoreResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreResponseBodyData) SetAppendMeta(v bool) *DescribeLogStoreResponseBodyData {
	s.AppendMeta = &v
	return s
}

func (s *DescribeLogStoreResponseBodyData) SetAutoSplit(v bool) *DescribeLogStoreResponseBodyData {
	s.AutoSplit = &v
	return s
}

func (s *DescribeLogStoreResponseBodyData) SetEnableTracking(v bool) *DescribeLogStoreResponseBodyData {
	s.EnableTracking = &v
	return s
}

func (s *DescribeLogStoreResponseBodyData) SetLogStoreName(v string) *DescribeLogStoreResponseBodyData {
	s.LogStoreName = &v
	return s
}

func (s *DescribeLogStoreResponseBodyData) SetMaxSplitShard(v int32) *DescribeLogStoreResponseBodyData {
	s.MaxSplitShard = &v
	return s
}

func (s *DescribeLogStoreResponseBodyData) SetShardCount(v int32) *DescribeLogStoreResponseBodyData {
	s.ShardCount = &v
	return s
}

func (s *DescribeLogStoreResponseBodyData) SetTtl(v int32) *DescribeLogStoreResponseBodyData {
	s.Ttl = &v
	return s
}

type DescribeLogStoreResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLogStoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreResponse) SetHeaders(v map[string]*string) *DescribeLogStoreResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogStoreResponse) SetStatusCode(v int32) *DescribeLogStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogStoreResponse) SetBody(v *DescribeLogStoreResponseBody) *DescribeLogStoreResponse {
	s.Body = v
	return s
}

type DescribeLogTypeRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLogTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogTypeRequest) SetRegionId(v string) *DescribeLogTypeRequest {
	s.RegionId = &v
	return s
}

type DescribeLogTypeResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeLogTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeLogTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogTypeResponseBody) SetCode(v int32) *DescribeLogTypeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogTypeResponseBody) SetData(v []*DescribeLogTypeResponseBodyData) *DescribeLogTypeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLogTypeResponseBody) SetMessage(v string) *DescribeLogTypeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLogTypeResponseBody) SetRequestId(v string) *DescribeLogTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogTypeResponseBody) SetSuccess(v bool) *DescribeLogTypeResponseBody {
	s.Success = &v
	return s
}

type DescribeLogTypeResponseBodyData struct {
	LogType     *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	LogTypeName *string `json:"LogTypeName,omitempty" xml:"LogTypeName,omitempty"`
}

func (s DescribeLogTypeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLogTypeResponseBodyData) SetLogType(v string) *DescribeLogTypeResponseBodyData {
	s.LogType = &v
	return s
}

func (s *DescribeLogTypeResponseBodyData) SetLogTypeName(v string) *DescribeLogTypeResponseBodyData {
	s.LogTypeName = &v
	return s
}

type DescribeLogTypeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLogTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogTypeResponse) SetHeaders(v map[string]*string) *DescribeLogTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogTypeResponse) SetStatusCode(v int32) *DescribeLogTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogTypeResponse) SetBody(v *DescribeLogTypeResponseBody) *DescribeLogTypeResponse {
	s.Body = v
	return s
}

type DescribeOperatorsRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
}

func (s DescribeOperatorsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOperatorsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOperatorsRequest) SetRegionId(v string) *DescribeOperatorsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOperatorsRequest) SetSceneType(v string) *DescribeOperatorsRequest {
	s.SceneType = &v
	return s
}

type DescribeOperatorsResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeOperatorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeOperatorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOperatorsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOperatorsResponseBody) SetCode(v int32) *DescribeOperatorsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeOperatorsResponseBody) SetData(v []*DescribeOperatorsResponseBodyData) *DescribeOperatorsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOperatorsResponseBody) SetMessage(v string) *DescribeOperatorsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeOperatorsResponseBody) SetRequestId(v string) *DescribeOperatorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOperatorsResponseBody) SetSuccess(v bool) *DescribeOperatorsResponseBody {
	s.Success = &v
	return s
}

type DescribeOperatorsResponseBodyData struct {
	Index           *int32    `json:"Index,omitempty" xml:"Index,omitempty"`
	Operator        *string   `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorDescCn  *string   `json:"OperatorDescCn,omitempty" xml:"OperatorDescCn,omitempty"`
	OperatorDescEn  *string   `json:"OperatorDescEn,omitempty" xml:"OperatorDescEn,omitempty"`
	OperatorName    *string   `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	SupportDataType *string   `json:"SupportDataType,omitempty" xml:"SupportDataType,omitempty"`
	SupportTag      []*string `json:"SupportTag,omitempty" xml:"SupportTag,omitempty" type:"Repeated"`
}

func (s DescribeOperatorsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeOperatorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOperatorsResponseBodyData) SetIndex(v int32) *DescribeOperatorsResponseBodyData {
	s.Index = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetOperator(v string) *DescribeOperatorsResponseBodyData {
	s.Operator = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetOperatorDescCn(v string) *DescribeOperatorsResponseBodyData {
	s.OperatorDescCn = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetOperatorDescEn(v string) *DescribeOperatorsResponseBodyData {
	s.OperatorDescEn = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetOperatorName(v string) *DescribeOperatorsResponseBodyData {
	s.OperatorName = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetSupportDataType(v string) *DescribeOperatorsResponseBodyData {
	s.SupportDataType = &v
	return s
}

func (s *DescribeOperatorsResponseBodyData) SetSupportTag(v []*string) *DescribeOperatorsResponseBodyData {
	s.SupportTag = v
	return s
}

type DescribeOperatorsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOperatorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOperatorsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOperatorsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOperatorsResponse) SetHeaders(v map[string]*string) *DescribeOperatorsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOperatorsResponse) SetStatusCode(v int32) *DescribeOperatorsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOperatorsResponse) SetBody(v *DescribeOperatorsResponseBody) *DescribeOperatorsResponse {
	s.Body = v
	return s
}

type DescribeScopeUsersRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeScopeUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScopeUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeScopeUsersRequest) SetRegionId(v string) *DescribeScopeUsersRequest {
	s.RegionId = &v
	return s
}

type DescribeScopeUsersResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeScopeUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeScopeUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScopeUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScopeUsersResponseBody) SetCode(v int32) *DescribeScopeUsersResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeScopeUsersResponseBody) SetData(v []*DescribeScopeUsersResponseBodyData) *DescribeScopeUsersResponseBody {
	s.Data = v
	return s
}

func (s *DescribeScopeUsersResponseBody) SetMessage(v string) *DescribeScopeUsersResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeScopeUsersResponseBody) SetRequestId(v string) *DescribeScopeUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScopeUsersResponseBody) SetSuccess(v bool) *DescribeScopeUsersResponseBody {
	s.Success = &v
	return s
}

type DescribeScopeUsersResponseBodyData struct {
	AliUid     *int64    `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Domains    []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserName   *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeScopeUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScopeUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeScopeUsersResponseBodyData) SetAliUid(v int64) *DescribeScopeUsersResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *DescribeScopeUsersResponseBodyData) SetDomains(v []*string) *DescribeScopeUsersResponseBodyData {
	s.Domains = v
	return s
}

func (s *DescribeScopeUsersResponseBodyData) SetInstanceId(v string) *DescribeScopeUsersResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeScopeUsersResponseBodyData) SetUserName(v string) *DescribeScopeUsersResponseBodyData {
	s.UserName = &v
	return s
}

type DescribeScopeUsersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScopeUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScopeUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScopeUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeScopeUsersResponse) SetHeaders(v map[string]*string) *DescribeScopeUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeScopeUsersResponse) SetStatusCode(v int32) *DescribeScopeUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScopeUsersResponse) SetBody(v *DescribeScopeUsersResponseBody) *DescribeScopeUsersResponse {
	s.Body = v
	return s
}

type DescribeStorageRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageRequest) GoString() string {
	return s.String()
}

func (s *DescribeStorageRequest) SetRegionId(v string) *DescribeStorageRequest {
	s.RegionId = &v
	return s
}

type DescribeStorageResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	DyCode    *string `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStorageResponseBody) SetCode(v int32) *DescribeStorageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeStorageResponseBody) SetData(v bool) *DescribeStorageResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeStorageResponseBody) SetDyCode(v string) *DescribeStorageResponseBody {
	s.DyCode = &v
	return s
}

func (s *DescribeStorageResponseBody) SetDyMessage(v string) *DescribeStorageResponseBody {
	s.DyMessage = &v
	return s
}

func (s *DescribeStorageResponseBody) SetErrCode(v string) *DescribeStorageResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeStorageResponseBody) SetMessage(v string) *DescribeStorageResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeStorageResponseBody) SetRequestId(v string) *DescribeStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStorageResponseBody) SetSuccess(v bool) *DescribeStorageResponseBody {
	s.Success = &v
	return s
}

type DescribeStorageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStorageResponse) GoString() string {
	return s.String()
}

func (s *DescribeStorageResponse) SetHeaders(v map[string]*string) *DescribeStorageResponse {
	s.Headers = v
	return s
}

func (s *DescribeStorageResponse) SetStatusCode(v int32) *DescribeStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStorageResponse) SetBody(v *DescribeStorageResponseBody) *DescribeStorageResponse {
	s.Body = v
	return s
}

type DescribeWafScopeRequest struct {
	EntityId *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeWafScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafScopeRequest) GoString() string {
	return s.String()
}

func (s *DescribeWafScopeRequest) SetEntityId(v int64) *DescribeWafScopeRequest {
	s.EntityId = &v
	return s
}

func (s *DescribeWafScopeRequest) SetRegionId(v string) *DescribeWafScopeRequest {
	s.RegionId = &v
	return s
}

type DescribeWafScopeResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*DescribeWafScopeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeWafScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafScopeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWafScopeResponseBody) SetCode(v int32) *DescribeWafScopeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeWafScopeResponseBody) SetData(v []*DescribeWafScopeResponseBodyData) *DescribeWafScopeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeWafScopeResponseBody) SetMessage(v string) *DescribeWafScopeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeWafScopeResponseBody) SetRequestId(v string) *DescribeWafScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWafScopeResponseBody) SetSuccess(v bool) *DescribeWafScopeResponseBody {
	s.Success = &v
	return s
}

type DescribeWafScopeResponseBodyData struct {
	Aliuid     *int64    `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	Domains    []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeWafScopeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafScopeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeWafScopeResponseBodyData) SetAliuid(v int64) *DescribeWafScopeResponseBodyData {
	s.Aliuid = &v
	return s
}

func (s *DescribeWafScopeResponseBodyData) SetDomains(v []*string) *DescribeWafScopeResponseBodyData {
	s.Domains = v
	return s
}

func (s *DescribeWafScopeResponseBodyData) SetInstanceId(v string) *DescribeWafScopeResponseBodyData {
	s.InstanceId = &v
	return s
}

type DescribeWafScopeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeWafScopeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWafScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafScopeResponse) GoString() string {
	return s.String()
}

func (s *DescribeWafScopeResponse) SetHeaders(v map[string]*string) *DescribeWafScopeResponse {
	s.Headers = v
	return s
}

func (s *DescribeWafScopeResponse) SetStatusCode(v int32) *DescribeWafScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWafScopeResponse) SetBody(v *DescribeWafScopeResponseBody) *DescribeWafScopeResponse {
	s.Body = v
	return s
}

type DoQuickFieldRequest struct {
	From     *int32  `json:"From,omitempty" xml:"From,omitempty"`
	Index    *string `json:"Index,omitempty" xml:"Index,omitempty"`
	Page     *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Reverse  *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	Size     *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	To       *int32  `json:"To,omitempty" xml:"To,omitempty"`
}

func (s DoQuickFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s DoQuickFieldRequest) GoString() string {
	return s.String()
}

func (s *DoQuickFieldRequest) SetFrom(v int32) *DoQuickFieldRequest {
	s.From = &v
	return s
}

func (s *DoQuickFieldRequest) SetIndex(v string) *DoQuickFieldRequest {
	s.Index = &v
	return s
}

func (s *DoQuickFieldRequest) SetPage(v int32) *DoQuickFieldRequest {
	s.Page = &v
	return s
}

func (s *DoQuickFieldRequest) SetRegionId(v string) *DoQuickFieldRequest {
	s.RegionId = &v
	return s
}

func (s *DoQuickFieldRequest) SetReverse(v bool) *DoQuickFieldRequest {
	s.Reverse = &v
	return s
}

func (s *DoQuickFieldRequest) SetSize(v int32) *DoQuickFieldRequest {
	s.Size = &v
	return s
}

func (s *DoQuickFieldRequest) SetTo(v int32) *DoQuickFieldRequest {
	s.To = &v
	return s
}

type DoQuickFieldResponseBody struct {
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DoQuickFieldResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DyCode    *string                       `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string                       `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string                       `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DoQuickFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DoQuickFieldResponseBody) GoString() string {
	return s.String()
}

func (s *DoQuickFieldResponseBody) SetCode(v int32) *DoQuickFieldResponseBody {
	s.Code = &v
	return s
}

func (s *DoQuickFieldResponseBody) SetData(v *DoQuickFieldResponseBodyData) *DoQuickFieldResponseBody {
	s.Data = v
	return s
}

func (s *DoQuickFieldResponseBody) SetDyCode(v string) *DoQuickFieldResponseBody {
	s.DyCode = &v
	return s
}

func (s *DoQuickFieldResponseBody) SetDyMessage(v string) *DoQuickFieldResponseBody {
	s.DyMessage = &v
	return s
}

func (s *DoQuickFieldResponseBody) SetErrCode(v string) *DoQuickFieldResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DoQuickFieldResponseBody) SetMessage(v string) *DoQuickFieldResponseBody {
	s.Message = &v
	return s
}

func (s *DoQuickFieldResponseBody) SetRequestId(v string) *DoQuickFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *DoQuickFieldResponseBody) SetSuccess(v bool) *DoQuickFieldResponseBody {
	s.Success = &v
	return s
}

type DoQuickFieldResponseBodyData struct {
	AggQueryd     *string       `json:"AggQueryd,omitempty" xml:"AggQueryd,omitempty"`
	CompleteOrNot *bool         `json:"CompleteOrNot,omitempty" xml:"CompleteOrNot,omitempty"`
	Count         *int32        `json:"Count,omitempty" xml:"Count,omitempty"`
	HasSQL        *bool         `json:"HasSQL,omitempty" xml:"HasSQL,omitempty"`
	Keys          []*string     `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	Limited       *int64        `json:"Limited,omitempty" xml:"Limited,omitempty"`
	Logs          []interface{} `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	PQuery        *string       `json:"PQuery,omitempty" xml:"PQuery,omitempty"`
	ProcessedRows *int64        `json:"ProcessedRows,omitempty" xml:"ProcessedRows,omitempty"`
	QueryMode     *int32        `json:"QueryMode,omitempty" xml:"QueryMode,omitempty"`
	WhereQuery    *string       `json:"WhereQuery,omitempty" xml:"WhereQuery,omitempty"`
}

func (s DoQuickFieldResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DoQuickFieldResponseBodyData) GoString() string {
	return s.String()
}

func (s *DoQuickFieldResponseBodyData) SetAggQueryd(v string) *DoQuickFieldResponseBodyData {
	s.AggQueryd = &v
	return s
}

func (s *DoQuickFieldResponseBodyData) SetCompleteOrNot(v bool) *DoQuickFieldResponseBodyData {
	s.CompleteOrNot = &v
	return s
}

func (s *DoQuickFieldResponseBodyData) SetCount(v int32) *DoQuickFieldResponseBodyData {
	s.Count = &v
	return s
}

func (s *DoQuickFieldResponseBodyData) SetHasSQL(v bool) *DoQuickFieldResponseBodyData {
	s.HasSQL = &v
	return s
}

func (s *DoQuickFieldResponseBodyData) SetKeys(v []*string) *DoQuickFieldResponseBodyData {
	s.Keys = v
	return s
}

func (s *DoQuickFieldResponseBodyData) SetLimited(v int64) *DoQuickFieldResponseBodyData {
	s.Limited = &v
	return s
}

func (s *DoQuickFieldResponseBodyData) SetLogs(v []interface{}) *DoQuickFieldResponseBodyData {
	s.Logs = v
	return s
}

func (s *DoQuickFieldResponseBodyData) SetPQuery(v string) *DoQuickFieldResponseBodyData {
	s.PQuery = &v
	return s
}

func (s *DoQuickFieldResponseBodyData) SetProcessedRows(v int64) *DoQuickFieldResponseBodyData {
	s.ProcessedRows = &v
	return s
}

func (s *DoQuickFieldResponseBodyData) SetQueryMode(v int32) *DoQuickFieldResponseBodyData {
	s.QueryMode = &v
	return s
}

func (s *DoQuickFieldResponseBodyData) SetWhereQuery(v string) *DoQuickFieldResponseBodyData {
	s.WhereQuery = &v
	return s
}

type DoQuickFieldResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DoQuickFieldResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DoQuickFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s DoQuickFieldResponse) GoString() string {
	return s.String()
}

func (s *DoQuickFieldResponse) SetHeaders(v map[string]*string) *DoQuickFieldResponse {
	s.Headers = v
	return s
}

func (s *DoQuickFieldResponse) SetStatusCode(v int32) *DoQuickFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *DoQuickFieldResponse) SetBody(v *DoQuickFieldResponseBody) *DoQuickFieldResponse {
	s.Body = v
	return s
}

type DoSelfDelegateRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	DelegateOrNot *int32  `json:"DelegateOrNot,omitempty" xml:"DelegateOrNot,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DoSelfDelegateRequest) String() string {
	return tea.Prettify(s)
}

func (s DoSelfDelegateRequest) GoString() string {
	return s.String()
}

func (s *DoSelfDelegateRequest) SetAliUid(v int64) *DoSelfDelegateRequest {
	s.AliUid = &v
	return s
}

func (s *DoSelfDelegateRequest) SetDelegateOrNot(v int32) *DoSelfDelegateRequest {
	s.DelegateOrNot = &v
	return s
}

func (s *DoSelfDelegateRequest) SetRegionId(v string) *DoSelfDelegateRequest {
	s.RegionId = &v
	return s
}

type DoSelfDelegateResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	DyCode    *string `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DoSelfDelegateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DoSelfDelegateResponseBody) GoString() string {
	return s.String()
}

func (s *DoSelfDelegateResponseBody) SetCode(v int32) *DoSelfDelegateResponseBody {
	s.Code = &v
	return s
}

func (s *DoSelfDelegateResponseBody) SetData(v bool) *DoSelfDelegateResponseBody {
	s.Data = &v
	return s
}

func (s *DoSelfDelegateResponseBody) SetDyCode(v string) *DoSelfDelegateResponseBody {
	s.DyCode = &v
	return s
}

func (s *DoSelfDelegateResponseBody) SetDyMessage(v string) *DoSelfDelegateResponseBody {
	s.DyMessage = &v
	return s
}

func (s *DoSelfDelegateResponseBody) SetErrCode(v string) *DoSelfDelegateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DoSelfDelegateResponseBody) SetMessage(v string) *DoSelfDelegateResponseBody {
	s.Message = &v
	return s
}

func (s *DoSelfDelegateResponseBody) SetRequestId(v string) *DoSelfDelegateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DoSelfDelegateResponseBody) SetSuccess(v bool) *DoSelfDelegateResponseBody {
	s.Success = &v
	return s
}

type DoSelfDelegateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DoSelfDelegateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DoSelfDelegateResponse) String() string {
	return tea.Prettify(s)
}

func (s DoSelfDelegateResponse) GoString() string {
	return s.String()
}

func (s *DoSelfDelegateResponse) SetHeaders(v map[string]*string) *DoSelfDelegateResponse {
	s.Headers = v
	return s
}

func (s *DoSelfDelegateResponse) SetStatusCode(v int32) *DoSelfDelegateResponse {
	s.StatusCode = &v
	return s
}

func (s *DoSelfDelegateResponse) SetBody(v *DoSelfDelegateResponseBody) *DoSelfDelegateResponse {
	s.Body = v
	return s
}

type GetCapacityRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetCapacityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCapacityRequest) GoString() string {
	return s.String()
}

func (s *GetCapacityRequest) SetRegionId(v string) *GetCapacityRequest {
	s.RegionId = &v
	return s
}

type GetCapacityResponseBody struct {
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetCapacityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DyCode    *string                      `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string                      `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string                      `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCapacityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *GetCapacityResponseBody) SetCode(v int32) *GetCapacityResponseBody {
	s.Code = &v
	return s
}

func (s *GetCapacityResponseBody) SetData(v *GetCapacityResponseBodyData) *GetCapacityResponseBody {
	s.Data = v
	return s
}

func (s *GetCapacityResponseBody) SetDyCode(v string) *GetCapacityResponseBody {
	s.DyCode = &v
	return s
}

func (s *GetCapacityResponseBody) SetDyMessage(v string) *GetCapacityResponseBody {
	s.DyMessage = &v
	return s
}

func (s *GetCapacityResponseBody) SetErrCode(v string) *GetCapacityResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetCapacityResponseBody) SetMessage(v string) *GetCapacityResponseBody {
	s.Message = &v
	return s
}

func (s *GetCapacityResponseBody) SetRequestId(v string) *GetCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCapacityResponseBody) SetSuccess(v bool) *GetCapacityResponseBody {
	s.Success = &v
	return s
}

type GetCapacityResponseBodyData struct {
	ExistLogStore     *bool    `json:"ExistLogStore,omitempty" xml:"ExistLogStore,omitempty"`
	PreservedCapacity *int64   `json:"PreservedCapacity,omitempty" xml:"PreservedCapacity,omitempty"`
	UsedCapacity      *float64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s GetCapacityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCapacityResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCapacityResponseBodyData) SetExistLogStore(v bool) *GetCapacityResponseBodyData {
	s.ExistLogStore = &v
	return s
}

func (s *GetCapacityResponseBodyData) SetPreservedCapacity(v int64) *GetCapacityResponseBodyData {
	s.PreservedCapacity = &v
	return s
}

func (s *GetCapacityResponseBodyData) SetUsedCapacity(v float64) *GetCapacityResponseBodyData {
	s.UsedCapacity = &v
	return s
}

type GetCapacityResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCapacityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCapacityResponse) GoString() string {
	return s.String()
}

func (s *GetCapacityResponse) SetHeaders(v map[string]*string) *GetCapacityResponse {
	s.Headers = v
	return s
}

func (s *GetCapacityResponse) SetStatusCode(v int32) *GetCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCapacityResponse) SetBody(v *GetCapacityResponseBody) *GetCapacityResponse {
	s.Body = v
	return s
}

type GetHistogramsRequest struct {
	From     *int32  `json:"From,omitempty" xml:"From,omitempty"`
	Query    *string `json:"Query,omitempty" xml:"Query,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	To       *int32  `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetHistogramsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHistogramsRequest) GoString() string {
	return s.String()
}

func (s *GetHistogramsRequest) SetFrom(v int32) *GetHistogramsRequest {
	s.From = &v
	return s
}

func (s *GetHistogramsRequest) SetQuery(v string) *GetHistogramsRequest {
	s.Query = &v
	return s
}

func (s *GetHistogramsRequest) SetRegionId(v string) *GetHistogramsRequest {
	s.RegionId = &v
	return s
}

func (s *GetHistogramsRequest) SetTo(v int32) *GetHistogramsRequest {
	s.To = &v
	return s
}

type GetHistogramsResponseBody struct {
	Code      *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetHistogramsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DyCode    *string                        `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string                        `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string                        `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHistogramsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHistogramsResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistogramsResponseBody) SetCode(v int32) *GetHistogramsResponseBody {
	s.Code = &v
	return s
}

func (s *GetHistogramsResponseBody) SetData(v *GetHistogramsResponseBodyData) *GetHistogramsResponseBody {
	s.Data = v
	return s
}

func (s *GetHistogramsResponseBody) SetDyCode(v string) *GetHistogramsResponseBody {
	s.DyCode = &v
	return s
}

func (s *GetHistogramsResponseBody) SetDyMessage(v string) *GetHistogramsResponseBody {
	s.DyMessage = &v
	return s
}

func (s *GetHistogramsResponseBody) SetErrCode(v string) *GetHistogramsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetHistogramsResponseBody) SetMessage(v string) *GetHistogramsResponseBody {
	s.Message = &v
	return s
}

func (s *GetHistogramsResponseBody) SetRequestId(v string) *GetHistogramsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHistogramsResponseBody) SetSuccess(v bool) *GetHistogramsResponseBody {
	s.Success = &v
	return s
}

type GetHistogramsResponseBodyData struct {
	Histograms []*GetHistogramsResponseBodyDataHistograms `json:"Histograms,omitempty" xml:"Histograms,omitempty" type:"Repeated"`
	Server     *string                                    `json:"Server,omitempty" xml:"Server,omitempty"`
	TotalCount *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetHistogramsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHistogramsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHistogramsResponseBodyData) SetHistograms(v []*GetHistogramsResponseBodyDataHistograms) *GetHistogramsResponseBodyData {
	s.Histograms = v
	return s
}

func (s *GetHistogramsResponseBodyData) SetServer(v string) *GetHistogramsResponseBodyData {
	s.Server = &v
	return s
}

func (s *GetHistogramsResponseBodyData) SetTotalCount(v int64) *GetHistogramsResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetHistogramsResponseBodyDataHistograms struct {
	CompletedOrNot *bool  `json:"CompletedOrNot,omitempty" xml:"CompletedOrNot,omitempty"`
	Count          *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	From           *int32 `json:"From,omitempty" xml:"From,omitempty"`
	To             *int32 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s GetHistogramsResponseBodyDataHistograms) String() string {
	return tea.Prettify(s)
}

func (s GetHistogramsResponseBodyDataHistograms) GoString() string {
	return s.String()
}

func (s *GetHistogramsResponseBodyDataHistograms) SetCompletedOrNot(v bool) *GetHistogramsResponseBodyDataHistograms {
	s.CompletedOrNot = &v
	return s
}

func (s *GetHistogramsResponseBodyDataHistograms) SetCount(v int64) *GetHistogramsResponseBodyDataHistograms {
	s.Count = &v
	return s
}

func (s *GetHistogramsResponseBodyDataHistograms) SetFrom(v int32) *GetHistogramsResponseBodyDataHistograms {
	s.From = &v
	return s
}

func (s *GetHistogramsResponseBodyDataHistograms) SetTo(v int32) *GetHistogramsResponseBodyDataHistograms {
	s.To = &v
	return s
}

type GetHistogramsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHistogramsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHistogramsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHistogramsResponse) GoString() string {
	return s.String()
}

func (s *GetHistogramsResponse) SetHeaders(v map[string]*string) *GetHistogramsResponse {
	s.Headers = v
	return s
}

func (s *GetHistogramsResponse) SetStatusCode(v int32) *GetHistogramsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistogramsResponse) SetBody(v *GetHistogramsResponseBody) *GetHistogramsResponse {
	s.Body = v
	return s
}

type GetLogsRequest struct {
	From         *int32  `json:"From,omitempty" xml:"From,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query        *string `json:"Query,omitempty" xml:"Query,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReverseOrNot *bool   `json:"ReverseOrNot,omitempty" xml:"ReverseOrNot,omitempty"`
	To           *int32  `json:"To,omitempty" xml:"To,omitempty"`
	Total        *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogsRequest) GoString() string {
	return s.String()
}

func (s *GetLogsRequest) SetFrom(v int32) *GetLogsRequest {
	s.From = &v
	return s
}

func (s *GetLogsRequest) SetPageIndex(v int32) *GetLogsRequest {
	s.PageIndex = &v
	return s
}

func (s *GetLogsRequest) SetPageSize(v int32) *GetLogsRequest {
	s.PageSize = &v
	return s
}

func (s *GetLogsRequest) SetQuery(v string) *GetLogsRequest {
	s.Query = &v
	return s
}

func (s *GetLogsRequest) SetRegionId(v string) *GetLogsRequest {
	s.RegionId = &v
	return s
}

func (s *GetLogsRequest) SetReverseOrNot(v bool) *GetLogsRequest {
	s.ReverseOrNot = &v
	return s
}

func (s *GetLogsRequest) SetTo(v int32) *GetLogsRequest {
	s.To = &v
	return s
}

func (s *GetLogsRequest) SetTotal(v int64) *GetLogsRequest {
	s.Total = &v
	return s
}

type GetLogsResponseBody struct {
	Code      *int32                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogsResponseBody) SetCode(v int32) *GetLogsResponseBody {
	s.Code = &v
	return s
}

func (s *GetLogsResponseBody) SetData(v *GetLogsResponseBodyData) *GetLogsResponseBody {
	s.Data = v
	return s
}

func (s *GetLogsResponseBody) SetMessage(v string) *GetLogsResponseBody {
	s.Message = &v
	return s
}

func (s *GetLogsResponseBody) SetRequestId(v string) *GetLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogsResponseBody) SetSuccess(v bool) *GetLogsResponseBody {
	s.Success = &v
	return s
}

type GetLogsResponseBodyData struct {
	PageInfo     *GetLogsResponseBodyDataPageInfo     `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	ResponseData *GetLogsResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Struct"`
}

func (s GetLogsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLogsResponseBodyData) SetPageInfo(v *GetLogsResponseBodyDataPageInfo) *GetLogsResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *GetLogsResponseBodyData) SetResponseData(v *GetLogsResponseBodyDataResponseData) *GetLogsResponseBodyData {
	s.ResponseData = v
	return s
}

type GetLogsResponseBodyDataPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetLogsResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetLogsResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *GetLogsResponseBodyDataPageInfo) SetCurrentPage(v int32) *GetLogsResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *GetLogsResponseBodyDataPageInfo) SetPageSize(v int32) *GetLogsResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetLogsResponseBodyDataPageInfo) SetTotalCount(v int64) *GetLogsResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type GetLogsResponseBodyDataResponseData struct {
	CompleteOrNot *bool         `json:"CompleteOrNot,omitempty" xml:"CompleteOrNot,omitempty"`
	Cost          *int64        `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Count         *int32        `json:"Count,omitempty" xml:"Count,omitempty"`
	HasSql        *bool         `json:"HasSql,omitempty" xml:"HasSql,omitempty"`
	Keys          []*string     `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	Lines         []interface{} `json:"Lines,omitempty" xml:"Lines,omitempty" type:"Repeated"`
}

func (s GetLogsResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetLogsResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *GetLogsResponseBodyDataResponseData) SetCompleteOrNot(v bool) *GetLogsResponseBodyDataResponseData {
	s.CompleteOrNot = &v
	return s
}

func (s *GetLogsResponseBodyDataResponseData) SetCost(v int64) *GetLogsResponseBodyDataResponseData {
	s.Cost = &v
	return s
}

func (s *GetLogsResponseBodyDataResponseData) SetCount(v int32) *GetLogsResponseBodyDataResponseData {
	s.Count = &v
	return s
}

func (s *GetLogsResponseBodyDataResponseData) SetHasSql(v bool) *GetLogsResponseBodyDataResponseData {
	s.HasSql = &v
	return s
}

func (s *GetLogsResponseBodyDataResponseData) SetKeys(v []*string) *GetLogsResponseBodyDataResponseData {
	s.Keys = v
	return s
}

func (s *GetLogsResponseBodyDataResponseData) SetLines(v []interface{}) *GetLogsResponseBodyDataResponseData {
	s.Lines = v
	return s
}

type GetLogsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogsResponse) GoString() string {
	return s.String()
}

func (s *GetLogsResponse) SetHeaders(v map[string]*string) *GetLogsResponse {
	s.Headers = v
	return s
}

func (s *GetLogsResponse) SetStatusCode(v int32) *GetLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogsResponse) SetBody(v *GetLogsResponseBody) *GetLogsResponse {
	s.Body = v
	return s
}

type GetQuickQueryRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s GetQuickQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuickQueryRequest) GoString() string {
	return s.String()
}

func (s *GetQuickQueryRequest) SetRegionId(v string) *GetQuickQueryRequest {
	s.RegionId = &v
	return s
}

func (s *GetQuickQueryRequest) SetSearchName(v string) *GetQuickQueryRequest {
	s.SearchName = &v
	return s
}

type GetQuickQueryResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DyCode    *string `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQuickQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuickQueryResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuickQueryResponseBody) SetCode(v int32) *GetQuickQueryResponseBody {
	s.Code = &v
	return s
}

func (s *GetQuickQueryResponseBody) SetData(v string) *GetQuickQueryResponseBody {
	s.Data = &v
	return s
}

func (s *GetQuickQueryResponseBody) SetDyCode(v string) *GetQuickQueryResponseBody {
	s.DyCode = &v
	return s
}

func (s *GetQuickQueryResponseBody) SetDyMessage(v string) *GetQuickQueryResponseBody {
	s.DyMessage = &v
	return s
}

func (s *GetQuickQueryResponseBody) SetErrCode(v string) *GetQuickQueryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetQuickQueryResponseBody) SetMessage(v string) *GetQuickQueryResponseBody {
	s.Message = &v
	return s
}

func (s *GetQuickQueryResponseBody) SetRequestId(v string) *GetQuickQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuickQueryResponseBody) SetSuccess(v bool) *GetQuickQueryResponseBody {
	s.Success = &v
	return s
}

type GetQuickQueryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQuickQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQuickQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuickQueryResponse) GoString() string {
	return s.String()
}

func (s *GetQuickQueryResponse) SetHeaders(v map[string]*string) *GetQuickQueryResponse {
	s.Headers = v
	return s
}

func (s *GetQuickQueryResponse) SetStatusCode(v int32) *GetQuickQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuickQueryResponse) SetBody(v *GetQuickQueryResponseBody) *GetQuickQueryResponse {
	s.Body = v
	return s
}

type GetStorageRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStorageRequest) GoString() string {
	return s.String()
}

func (s *GetStorageRequest) SetRegionId(v string) *GetStorageRequest {
	s.RegionId = &v
	return s
}

type GetStorageResponseBody struct {
	Code      *int32                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetStorageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DyCode    *string                     `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string                     `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string                     `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStorageResponseBody) GoString() string {
	return s.String()
}

func (s *GetStorageResponseBody) SetCode(v int32) *GetStorageResponseBody {
	s.Code = &v
	return s
}

func (s *GetStorageResponseBody) SetData(v *GetStorageResponseBodyData) *GetStorageResponseBody {
	s.Data = v
	return s
}

func (s *GetStorageResponseBody) SetDyCode(v string) *GetStorageResponseBody {
	s.DyCode = &v
	return s
}

func (s *GetStorageResponseBody) SetDyMessage(v string) *GetStorageResponseBody {
	s.DyMessage = &v
	return s
}

func (s *GetStorageResponseBody) SetErrCode(v string) *GetStorageResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetStorageResponseBody) SetMessage(v string) *GetStorageResponseBody {
	s.Message = &v
	return s
}

func (s *GetStorageResponseBody) SetRequestId(v string) *GetStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStorageResponseBody) SetSuccess(v bool) *GetStorageResponseBody {
	s.Success = &v
	return s
}

type GetStorageResponseBodyData struct {
	CanOperate    *bool   `json:"CanOperate,omitempty" xml:"CanOperate,omitempty"`
	DisplayRegion *bool   `json:"DisplayRegion,omitempty" xml:"DisplayRegion,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Ttl           *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s GetStorageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetStorageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStorageResponseBodyData) SetCanOperate(v bool) *GetStorageResponseBodyData {
	s.CanOperate = &v
	return s
}

func (s *GetStorageResponseBodyData) SetDisplayRegion(v bool) *GetStorageResponseBodyData {
	s.DisplayRegion = &v
	return s
}

func (s *GetStorageResponseBodyData) SetRegion(v string) *GetStorageResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetStorageResponseBodyData) SetTtl(v int32) *GetStorageResponseBodyData {
	s.Ttl = &v
	return s
}

type GetStorageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStorageResponse) GoString() string {
	return s.String()
}

func (s *GetStorageResponse) SetHeaders(v map[string]*string) *GetStorageResponse {
	s.Headers = v
	return s
}

func (s *GetStorageResponse) SetStatusCode(v int32) *GetStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStorageResponse) SetBody(v *GetStorageResponseBody) *GetStorageResponse {
	s.Body = v
	return s
}

type ListAutomateResponseConfigsRequest struct {
	ActionType       *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	AutoResponseType *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
	CurrentPage      *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlaybookUuid     *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleName         *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SubUserId        *int64  `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListAutomateResponseConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAutomateResponseConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsRequest) SetActionType(v string) *ListAutomateResponseConfigsRequest {
	s.ActionType = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetAutoResponseType(v string) *ListAutomateResponseConfigsRequest {
	s.AutoResponseType = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetCurrentPage(v int32) *ListAutomateResponseConfigsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetId(v int64) *ListAutomateResponseConfigsRequest {
	s.Id = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetPageSize(v int32) *ListAutomateResponseConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetPlaybookUuid(v string) *ListAutomateResponseConfigsRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetRegionId(v string) *ListAutomateResponseConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetRuleName(v string) *ListAutomateResponseConfigsRequest {
	s.RuleName = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetStatus(v int32) *ListAutomateResponseConfigsRequest {
	s.Status = &v
	return s
}

func (s *ListAutomateResponseConfigsRequest) SetSubUserId(v int64) *ListAutomateResponseConfigsRequest {
	s.SubUserId = &v
	return s
}

type ListAutomateResponseConfigsResponseBody struct {
	Code      *int32                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListAutomateResponseConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAutomateResponseConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAutomateResponseConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsResponseBody) SetCode(v int32) *ListAutomateResponseConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBody) SetData(v *ListAutomateResponseConfigsResponseBodyData) *ListAutomateResponseConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListAutomateResponseConfigsResponseBody) SetMessage(v string) *ListAutomateResponseConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBody) SetRequestId(v string) *ListAutomateResponseConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBody) SetSuccess(v bool) *ListAutomateResponseConfigsResponseBody {
	s.Success = &v
	return s
}

type ListAutomateResponseConfigsResponseBodyData struct {
	PageInfo     *ListAutomateResponseConfigsResponseBodyDataPageInfo       `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	ResponseData []*ListAutomateResponseConfigsResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListAutomateResponseConfigsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAutomateResponseConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsResponseBodyData) SetPageInfo(v *ListAutomateResponseConfigsResponseBodyDataPageInfo) *ListAutomateResponseConfigsResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyData) SetResponseData(v []*ListAutomateResponseConfigsResponseBodyDataResponseData) *ListAutomateResponseConfigsResponseBodyData {
	s.ResponseData = v
	return s
}

type ListAutomateResponseConfigsResponseBodyDataPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAutomateResponseConfigsResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAutomateResponseConfigsResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListAutomateResponseConfigsResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataPageInfo) SetPageSize(v int32) *ListAutomateResponseConfigsResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataPageInfo) SetTotalCount(v int64) *ListAutomateResponseConfigsResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type ListAutomateResponseConfigsResponseBodyDataResponseData struct {
	ActionConfig       *string `json:"ActionConfig,omitempty" xml:"ActionConfig,omitempty"`
	ActionType         *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	Aliuid             *int64  `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	AutoResponseType   *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
	ExecutionCondition *string `json:"ExecutionCondition,omitempty" xml:"ExecutionCondition,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RuleName           *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Status             *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SubUserId          *int64  `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListAutomateResponseConfigsResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListAutomateResponseConfigsResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetActionConfig(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.ActionConfig = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetActionType(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.ActionType = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetAliuid(v int64) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetAutoResponseType(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.AutoResponseType = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetExecutionCondition(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.ExecutionCondition = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetGmtCreate(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetGmtModified(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetId(v int64) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetRuleName(v string) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.RuleName = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetStatus(v int32) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *ListAutomateResponseConfigsResponseBodyDataResponseData) SetSubUserId(v int64) *ListAutomateResponseConfigsResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

type ListAutomateResponseConfigsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAutomateResponseConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAutomateResponseConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAutomateResponseConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListAutomateResponseConfigsResponse) SetHeaders(v map[string]*string) *ListAutomateResponseConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListAutomateResponseConfigsResponse) SetStatusCode(v int32) *ListAutomateResponseConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutomateResponseConfigsResponse) SetBody(v *ListAutomateResponseConfigsResponseBody) *ListAutomateResponseConfigsResponse {
	s.Body = v
	return s
}

type ListCloudSiemCustomizeRulesRequest struct {
	AlertType   *string   `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	CurrentPage *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime     *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id          *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	PageSize    *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleName    *string   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType    *string   `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	StartTime   *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status      *int32    `json:"Status,omitempty" xml:"Status,omitempty"`
	ThreatLevel []*string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty" type:"Repeated"`
}

func (s ListCloudSiemCustomizeRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesRequest) SetAlertType(v string) *ListCloudSiemCustomizeRulesRequest {
	s.AlertType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetCurrentPage(v int32) *ListCloudSiemCustomizeRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetEndTime(v int64) *ListCloudSiemCustomizeRulesRequest {
	s.EndTime = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetId(v string) *ListCloudSiemCustomizeRulesRequest {
	s.Id = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetPageSize(v int32) *ListCloudSiemCustomizeRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetRegionId(v string) *ListCloudSiemCustomizeRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetRuleName(v string) *ListCloudSiemCustomizeRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetRuleType(v string) *ListCloudSiemCustomizeRulesRequest {
	s.RuleType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetStartTime(v int64) *ListCloudSiemCustomizeRulesRequest {
	s.StartTime = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetStatus(v int32) *ListCloudSiemCustomizeRulesRequest {
	s.Status = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesRequest) SetThreatLevel(v []*string) *ListCloudSiemCustomizeRulesRequest {
	s.ThreatLevel = v
	return s
}

type ListCloudSiemCustomizeRulesResponseBody struct {
	Code      *int32                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListCloudSiemCustomizeRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCloudSiemCustomizeRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesResponseBody) SetCode(v int32) *ListCloudSiemCustomizeRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBody) SetData(v *ListCloudSiemCustomizeRulesResponseBodyData) *ListCloudSiemCustomizeRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBody) SetMessage(v string) *ListCloudSiemCustomizeRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBody) SetRequestId(v string) *ListCloudSiemCustomizeRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBody) SetSuccess(v bool) *ListCloudSiemCustomizeRulesResponseBody {
	s.Success = &v
	return s
}

type ListCloudSiemCustomizeRulesResponseBodyData struct {
	PageInfo     *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo       `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	ResponseData []*ListCloudSiemCustomizeRulesResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListCloudSiemCustomizeRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesResponseBodyData) SetPageInfo(v *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) *ListCloudSiemCustomizeRulesResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyData) SetResponseData(v []*ListCloudSiemCustomizeRulesResponseBodyDataResponseData) *ListCloudSiemCustomizeRulesResponseBodyData {
	s.ResponseData = v
	return s
}

type ListCloudSiemCustomizeRulesResponseBodyDataPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) SetPageSize(v int32) *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) SetTotalCount(v int64) *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type ListCloudSiemCustomizeRulesResponseBodyDataResponseData struct {
	AlertType           *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AlertTypeMds        *string `json:"AlertTypeMds,omitempty" xml:"AlertTypeMds,omitempty"`
	Aliuid              *int64  `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	EventTransferExt    *string `json:"EventTransferExt,omitempty" xml:"EventTransferExt,omitempty"`
	EventTransferSwitch *int32  `json:"EventTransferSwitch,omitempty" xml:"EventTransferSwitch,omitempty"`
	EventTransferType   *string `json:"EventTransferType,omitempty" xml:"EventTransferType,omitempty"`
	GmtCreate           *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified         *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                  *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LogSource           *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	LogSourceMds        *string `json:"LogSourceMds,omitempty" xml:"LogSourceMds,omitempty"`
	LogType             *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	LogTypeMds          *string `json:"LogTypeMds,omitempty" xml:"LogTypeMds,omitempty"`
	QueryCycle          *string `json:"QueryCycle,omitempty" xml:"QueryCycle,omitempty"`
	RuleCondition       *string `json:"RuleCondition,omitempty" xml:"RuleCondition,omitempty"`
	RuleDesc            *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	RuleGroup           *string `json:"RuleGroup,omitempty" xml:"RuleGroup,omitempty"`
	RuleName            *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleThreshold       *string `json:"RuleThreshold,omitempty" xml:"RuleThreshold,omitempty"`
	RuleType            *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Status              *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ThreatLevel         *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
}

func (s ListCloudSiemCustomizeRulesResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetAlertType(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.AlertType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetAlertTypeMds(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.AlertTypeMds = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetAliuid(v int64) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetEventTransferExt(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.EventTransferExt = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetEventTransferSwitch(v int32) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.EventTransferSwitch = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetEventTransferType(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.EventTransferType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetGmtCreate(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetGmtModified(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetId(v int64) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetLogSource(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.LogSource = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetLogSourceMds(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.LogSourceMds = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetLogType(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.LogType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetLogTypeMds(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.LogTypeMds = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetQueryCycle(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.QueryCycle = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleCondition(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleCondition = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleDesc(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleDesc = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleGroup(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleGroup = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleName(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleName = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleThreshold(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleThreshold = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleType(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetStatus(v int32) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetThreatLevel(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.ThreatLevel = &v
	return s
}

type ListCloudSiemCustomizeRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCloudSiemCustomizeRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCloudSiemCustomizeRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesResponse) SetHeaders(v map[string]*string) *ListCloudSiemCustomizeRulesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponse) SetStatusCode(v int32) *ListCloudSiemCustomizeRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponse) SetBody(v *ListCloudSiemCustomizeRulesResponseBody) *ListCloudSiemCustomizeRulesResponse {
	s.Body = v
	return s
}

type ListCloudSiemPredefinedRulesRequest struct {
	AlertType   *string   `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	CurrentPage *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime     *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id          *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	PageSize    *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleName    *string   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType    *string   `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	StartTime   *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status      *int32    `json:"Status,omitempty" xml:"Status,omitempty"`
	ThreatLevel []*string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty" type:"Repeated"`
}

func (s ListCloudSiemPredefinedRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesRequest) SetAlertType(v string) *ListCloudSiemPredefinedRulesRequest {
	s.AlertType = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetCurrentPage(v int32) *ListCloudSiemPredefinedRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetEndTime(v int64) *ListCloudSiemPredefinedRulesRequest {
	s.EndTime = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetId(v string) *ListCloudSiemPredefinedRulesRequest {
	s.Id = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetPageSize(v int32) *ListCloudSiemPredefinedRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetRegionId(v string) *ListCloudSiemPredefinedRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetRuleName(v string) *ListCloudSiemPredefinedRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetRuleType(v string) *ListCloudSiemPredefinedRulesRequest {
	s.RuleType = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetStartTime(v int64) *ListCloudSiemPredefinedRulesRequest {
	s.StartTime = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetStatus(v int32) *ListCloudSiemPredefinedRulesRequest {
	s.Status = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesRequest) SetThreatLevel(v []*string) *ListCloudSiemPredefinedRulesRequest {
	s.ThreatLevel = v
	return s
}

type ListCloudSiemPredefinedRulesResponseBody struct {
	Code      *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListCloudSiemPredefinedRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCloudSiemPredefinedRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesResponseBody) SetCode(v int32) *ListCloudSiemPredefinedRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBody) SetData(v *ListCloudSiemPredefinedRulesResponseBodyData) *ListCloudSiemPredefinedRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBody) SetMessage(v string) *ListCloudSiemPredefinedRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBody) SetRequestId(v string) *ListCloudSiemPredefinedRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBody) SetSuccess(v bool) *ListCloudSiemPredefinedRulesResponseBody {
	s.Success = &v
	return s
}

type ListCloudSiemPredefinedRulesResponseBodyData struct {
	PageInfo     *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo       `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	ResponseData []*ListCloudSiemPredefinedRulesResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListCloudSiemPredefinedRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesResponseBodyData) SetPageInfo(v *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) *ListCloudSiemPredefinedRulesResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyData) SetResponseData(v []*ListCloudSiemPredefinedRulesResponseBodyDataResponseData) *ListCloudSiemPredefinedRulesResponseBodyData {
	s.ResponseData = v
	return s
}

type ListCloudSiemPredefinedRulesResponseBodyDataPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) SetPageSize(v int32) *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo) SetTotalCount(v int64) *ListCloudSiemPredefinedRulesResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type ListCloudSiemPredefinedRulesResponseBodyDataResponseData struct {
	AlertType   *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RuleDescMds *string `json:"RuleDescMds,omitempty" xml:"RuleDescMds,omitempty"`
	RuleName    *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleNameMds *string `json:"RuleNameMds,omitempty" xml:"RuleNameMds,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
}

func (s ListCloudSiemPredefinedRulesResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetAlertType(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.AlertType = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetGmtCreate(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetGmtModified(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetId(v int64) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetRuleDescMds(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.RuleDescMds = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetRuleName(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.RuleName = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetRuleNameMds(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.RuleNameMds = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetSource(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.Source = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetStatus(v int32) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponseBodyDataResponseData) SetThreatLevel(v string) *ListCloudSiemPredefinedRulesResponseBodyDataResponseData {
	s.ThreatLevel = &v
	return s
}

type ListCloudSiemPredefinedRulesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCloudSiemPredefinedRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCloudSiemPredefinedRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesResponse) SetHeaders(v map[string]*string) *ListCloudSiemPredefinedRulesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponse) SetStatusCode(v int32) *ListCloudSiemPredefinedRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponse) SetBody(v *ListCloudSiemPredefinedRulesResponseBody) *ListCloudSiemPredefinedRulesResponse {
	s.Body = v
	return s
}

type ListCustomizeRuleTestResultRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCustomizeRuleTestResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomizeRuleTestResultRequest) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultRequest) SetCurrentPage(v int32) *ListCustomizeRuleTestResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetId(v int64) *ListCustomizeRuleTestResultRequest {
	s.Id = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetPageSize(v int32) *ListCustomizeRuleTestResultRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomizeRuleTestResultRequest) SetRegionId(v string) *ListCustomizeRuleTestResultRequest {
	s.RegionId = &v
	return s
}

type ListCustomizeRuleTestResultResponseBody struct {
	Code      *int32                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListCustomizeRuleTestResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCustomizeRuleTestResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomizeRuleTestResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultResponseBody) SetCode(v int32) *ListCustomizeRuleTestResultResponseBody {
	s.Code = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBody) SetData(v *ListCustomizeRuleTestResultResponseBodyData) *ListCustomizeRuleTestResultResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBody) SetMessage(v string) *ListCustomizeRuleTestResultResponseBody {
	s.Message = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBody) SetRequestId(v string) *ListCustomizeRuleTestResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBody) SetSuccess(v bool) *ListCustomizeRuleTestResultResponseBody {
	s.Success = &v
	return s
}

type ListCustomizeRuleTestResultResponseBodyData struct {
	PageInfo     *ListCustomizeRuleTestResultResponseBodyDataPageInfo       `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	ResponseData []*ListCustomizeRuleTestResultResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListCustomizeRuleTestResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCustomizeRuleTestResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultResponseBodyData) SetPageInfo(v *ListCustomizeRuleTestResultResponseBodyDataPageInfo) *ListCustomizeRuleTestResultResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyData) SetResponseData(v []*ListCustomizeRuleTestResultResponseBodyDataResponseData) *ListCustomizeRuleTestResultResponseBodyData {
	s.ResponseData = v
	return s
}

type ListCustomizeRuleTestResultResponseBodyDataPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomizeRuleTestResultResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListCustomizeRuleTestResultResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListCustomizeRuleTestResultResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataPageInfo) SetPageSize(v int32) *ListCustomizeRuleTestResultResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataPageInfo) SetTotalCount(v int64) *ListCustomizeRuleTestResultResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type ListCustomizeRuleTestResultResponseBodyDataResponseData struct {
	AlertDesc          *string `json:"AlertDesc,omitempty" xml:"AlertDesc,omitempty"`
	AlertDetail        *string `json:"AlertDetail,omitempty" xml:"AlertDetail,omitempty"`
	AlertSrcProd       *string `json:"AlertSrcProd,omitempty" xml:"AlertSrcProd,omitempty"`
	AlertSrcProdModule *string `json:"AlertSrcProdModule,omitempty" xml:"AlertSrcProdModule,omitempty"`
	AttCk              *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	EventName          *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	EventType          *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Level              *string `json:"Level,omitempty" xml:"Level,omitempty"`
	LogSource          *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	LogTime            *string `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	LogType            *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	MainUserId         *string `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	OnlineStatus       *string `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	SubUserId          *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	Uuid               *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListCustomizeRuleTestResultResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListCustomizeRuleTestResultResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetAlertDesc(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.AlertDesc = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetAlertDetail(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.AlertDetail = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetAlertSrcProd(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.AlertSrcProd = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetAlertSrcProdModule(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.AlertSrcProdModule = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetAttCk(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.AttCk = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetEventName(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.EventName = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetEventType(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.EventType = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetLevel(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.Level = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetLogSource(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.LogSource = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetLogTime(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.LogTime = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetLogType(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.LogType = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetMainUserId(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.MainUserId = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetOnlineStatus(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.OnlineStatus = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetSubUserId(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetUuid(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.Uuid = &v
	return s
}

type ListCustomizeRuleTestResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCustomizeRuleTestResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCustomizeRuleTestResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomizeRuleTestResultResponse) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultResponse) SetHeaders(v map[string]*string) *ListCustomizeRuleTestResultResponse {
	s.Headers = v
	return s
}

func (s *ListCustomizeRuleTestResultResponse) SetStatusCode(v int32) *ListCustomizeRuleTestResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponse) SetBody(v *ListCustomizeRuleTestResultResponseBody) *ListCustomizeRuleTestResultResponse {
	s.Body = v
	return s
}

type ListDeliveryRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryRequest) GoString() string {
	return s.String()
}

func (s *ListDeliveryRequest) SetRegionId(v string) *ListDeliveryRequest {
	s.RegionId = &v
	return s
}

type ListDeliveryResponseBody struct {
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListDeliveryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DyCode    *string                       `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string                       `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string                       `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponseBody) SetCode(v int32) *ListDeliveryResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeliveryResponseBody) SetData(v *ListDeliveryResponseBodyData) *ListDeliveryResponseBody {
	s.Data = v
	return s
}

func (s *ListDeliveryResponseBody) SetDyCode(v string) *ListDeliveryResponseBody {
	s.DyCode = &v
	return s
}

func (s *ListDeliveryResponseBody) SetDyMessage(v string) *ListDeliveryResponseBody {
	s.DyMessage = &v
	return s
}

func (s *ListDeliveryResponseBody) SetErrCode(v string) *ListDeliveryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListDeliveryResponseBody) SetMessage(v string) *ListDeliveryResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeliveryResponseBody) SetRequestId(v string) *ListDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeliveryResponseBody) SetSuccess(v bool) *ListDeliveryResponseBody {
	s.Success = &v
	return s
}

type ListDeliveryResponseBodyData struct {
	DashboardUrl       *string                                    `json:"DashboardUrl,omitempty" xml:"DashboardUrl,omitempty"`
	DisplaySwitchOrNot *bool                                      `json:"DisplaySwitchOrNot,omitempty" xml:"DisplaySwitchOrNot,omitempty"`
	LogStoreName       *string                                    `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	ProductList        []*ListDeliveryResponseBodyDataProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Repeated"`
	ProjectName        *string                                    `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SearchUrl          *string                                    `json:"SearchUrl,omitempty" xml:"SearchUrl,omitempty"`
}

func (s ListDeliveryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponseBodyData) SetDashboardUrl(v string) *ListDeliveryResponseBodyData {
	s.DashboardUrl = &v
	return s
}

func (s *ListDeliveryResponseBodyData) SetDisplaySwitchOrNot(v bool) *ListDeliveryResponseBodyData {
	s.DisplaySwitchOrNot = &v
	return s
}

func (s *ListDeliveryResponseBodyData) SetLogStoreName(v string) *ListDeliveryResponseBodyData {
	s.LogStoreName = &v
	return s
}

func (s *ListDeliveryResponseBodyData) SetProductList(v []*ListDeliveryResponseBodyDataProductList) *ListDeliveryResponseBodyData {
	s.ProductList = v
	return s
}

func (s *ListDeliveryResponseBodyData) SetProjectName(v string) *ListDeliveryResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *ListDeliveryResponseBodyData) SetSearchUrl(v string) *ListDeliveryResponseBodyData {
	s.SearchUrl = &v
	return s
}

type ListDeliveryResponseBodyDataProductList struct {
	LogList     []*ListDeliveryResponseBodyDataProductListLogList `json:"LogList,omitempty" xml:"LogList,omitempty" type:"Repeated"`
	LogMap      map[string][]*DataProductListLogMapValue          `json:"LogMap,omitempty" xml:"LogMap,omitempty"`
	ProductCode *string                                           `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName *string                                           `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s ListDeliveryResponseBodyDataProductList) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryResponseBodyDataProductList) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponseBodyDataProductList) SetLogList(v []*ListDeliveryResponseBodyDataProductListLogList) *ListDeliveryResponseBodyDataProductList {
	s.LogList = v
	return s
}

func (s *ListDeliveryResponseBodyDataProductList) SetLogMap(v map[string][]*DataProductListLogMapValue) *ListDeliveryResponseBodyDataProductList {
	s.LogMap = v
	return s
}

func (s *ListDeliveryResponseBodyDataProductList) SetProductCode(v string) *ListDeliveryResponseBodyDataProductList {
	s.ProductCode = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductList) SetProductName(v string) *ListDeliveryResponseBodyDataProductList {
	s.ProductName = &v
	return s
}

type ListDeliveryResponseBodyDataProductListLogList struct {
	CanOperateOrNot *bool                                                            `json:"CanOperateOrNot,omitempty" xml:"CanOperateOrNot,omitempty"`
	ExtraParameters []*ListDeliveryResponseBodyDataProductListLogListExtraParameters `json:"ExtraParameters,omitempty" xml:"ExtraParameters,omitempty" type:"Repeated"`
	LogCode         *string                                                          `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	LogName         *string                                                          `json:"LogName,omitempty" xml:"LogName,omitempty"`
	LogNameEn       *string                                                          `json:"LogNameEn,omitempty" xml:"LogNameEn,omitempty"`
	LogNameKey      *string                                                          `json:"LogNameKey,omitempty" xml:"LogNameKey,omitempty"`
	Status          *bool                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	Topic           *string                                                          `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ListDeliveryResponseBodyDataProductListLogList) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryResponseBodyDataProductListLogList) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponseBodyDataProductListLogList) SetCanOperateOrNot(v bool) *ListDeliveryResponseBodyDataProductListLogList {
	s.CanOperateOrNot = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogList) SetExtraParameters(v []*ListDeliveryResponseBodyDataProductListLogListExtraParameters) *ListDeliveryResponseBodyDataProductListLogList {
	s.ExtraParameters = v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogList) SetLogCode(v string) *ListDeliveryResponseBodyDataProductListLogList {
	s.LogCode = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogList) SetLogName(v string) *ListDeliveryResponseBodyDataProductListLogList {
	s.LogName = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogList) SetLogNameEn(v string) *ListDeliveryResponseBodyDataProductListLogList {
	s.LogNameEn = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogList) SetLogNameKey(v string) *ListDeliveryResponseBodyDataProductListLogList {
	s.LogNameKey = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogList) SetStatus(v bool) *ListDeliveryResponseBodyDataProductListLogList {
	s.Status = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogList) SetTopic(v string) *ListDeliveryResponseBodyDataProductListLogList {
	s.Topic = &v
	return s
}

type ListDeliveryResponseBodyDataProductListLogListExtraParameters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDeliveryResponseBodyDataProductListLogListExtraParameters) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryResponseBodyDataProductListLogListExtraParameters) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponseBodyDataProductListLogListExtraParameters) SetKey(v string) *ListDeliveryResponseBodyDataProductListLogListExtraParameters {
	s.Key = &v
	return s
}

func (s *ListDeliveryResponseBodyDataProductListLogListExtraParameters) SetValue(v string) *ListDeliveryResponseBodyDataProductListLogListExtraParameters {
	s.Value = &v
	return s
}

type ListDeliveryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryResponse) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponse) SetHeaders(v map[string]*string) *ListDeliveryResponse {
	s.Headers = v
	return s
}

func (s *ListDeliveryResponse) SetStatusCode(v int32) *ListDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeliveryResponse) SetBody(v *ListDeliveryResponseBody) *ListDeliveryResponse {
	s.Body = v
	return s
}

type ListDisposeStrategyRequest struct {
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EffectiveStatus *int32  `json:"EffectiveStatus,omitempty" xml:"EffectiveStatus,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EntityIdentity  *string `json:"EntityIdentity,omitempty" xml:"EntityIdentity,omitempty"`
	EntityType      *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Order           *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OrderField      *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlaybookName    *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	PlaybookTypes   *string `json:"PlaybookTypes,omitempty" xml:"PlaybookTypes,omitempty"`
	PlaybookUuid    *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SophonTaskId    *string `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDisposeStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDisposeStrategyRequest) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyRequest) SetCurrentPage(v int32) *ListDisposeStrategyRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetEffectiveStatus(v int32) *ListDisposeStrategyRequest {
	s.EffectiveStatus = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetEndTime(v int64) *ListDisposeStrategyRequest {
	s.EndTime = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetEntityIdentity(v string) *ListDisposeStrategyRequest {
	s.EntityIdentity = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetEntityType(v string) *ListDisposeStrategyRequest {
	s.EntityType = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetOrder(v string) *ListDisposeStrategyRequest {
	s.Order = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetOrderField(v string) *ListDisposeStrategyRequest {
	s.OrderField = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetPageSize(v int32) *ListDisposeStrategyRequest {
	s.PageSize = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetPlaybookName(v string) *ListDisposeStrategyRequest {
	s.PlaybookName = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetPlaybookTypes(v string) *ListDisposeStrategyRequest {
	s.PlaybookTypes = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetPlaybookUuid(v string) *ListDisposeStrategyRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetRegionId(v string) *ListDisposeStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetSophonTaskId(v string) *ListDisposeStrategyRequest {
	s.SophonTaskId = &v
	return s
}

func (s *ListDisposeStrategyRequest) SetStartTime(v int64) *ListDisposeStrategyRequest {
	s.StartTime = &v
	return s
}

type ListDisposeStrategyResponseBody struct {
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListDisposeStrategyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDisposeStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDisposeStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponseBody) SetCode(v int32) *ListDisposeStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *ListDisposeStrategyResponseBody) SetData(v *ListDisposeStrategyResponseBodyData) *ListDisposeStrategyResponseBody {
	s.Data = v
	return s
}

func (s *ListDisposeStrategyResponseBody) SetMessage(v string) *ListDisposeStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *ListDisposeStrategyResponseBody) SetRequestId(v string) *ListDisposeStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDisposeStrategyResponseBody) SetSuccess(v bool) *ListDisposeStrategyResponseBody {
	s.Success = &v
	return s
}

type ListDisposeStrategyResponseBodyData struct {
	PageInfo     *ListDisposeStrategyResponseBodyDataPageInfo       `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	ResponseData []*ListDisposeStrategyResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListDisposeStrategyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDisposeStrategyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponseBodyData) SetPageInfo(v *ListDisposeStrategyResponseBodyDataPageInfo) *ListDisposeStrategyResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListDisposeStrategyResponseBodyData) SetResponseData(v []*ListDisposeStrategyResponseBodyDataResponseData) *ListDisposeStrategyResponseBodyData {
	s.ResponseData = v
	return s
}

type ListDisposeStrategyResponseBodyDataPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDisposeStrategyResponseBodyDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDisposeStrategyResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListDisposeStrategyResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataPageInfo) SetPageSize(v int32) *ListDisposeStrategyResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataPageInfo) SetTotalCount(v int64) *ListDisposeStrategyResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

type ListDisposeStrategyResponseBodyDataResponseData struct {
	AlertUuid       *string       `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	Aliuid          *int64        `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	EffectiveStatus *int32        `json:"EffectiveStatus,omitempty" xml:"EffectiveStatus,omitempty"`
	Entity          []interface{} `json:"Entity,omitempty" xml:"Entity,omitempty" type:"Repeated"`
	EntityId        *int64        `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType      *string       `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	ErrorMessage    *string       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FinishTime      *string       `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	GmtCreate       *string       `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *string       `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id              *int64        `json:"Id,omitempty" xml:"Id,omitempty"`
	IncidentName    *string       `json:"IncidentName,omitempty" xml:"IncidentName,omitempty"`
	IncidentUuid    *string       `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	PlaybookName    *string       `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	PlaybookType    *string       `json:"PlaybookType,omitempty" xml:"PlaybookType,omitempty"`
	PlaybookUuid    *string       `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	Scope           []interface{} `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
	SophonTaskId    *string       `json:"SophonTaskId,omitempty" xml:"SophonTaskId,omitempty"`
	Status          *int32        `json:"Status,omitempty" xml:"Status,omitempty"`
	SubAliuid       *int64        `json:"SubAliuid,omitempty" xml:"SubAliuid,omitempty"`
	TaskParam       *string       `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
}

func (s ListDisposeStrategyResponseBodyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListDisposeStrategyResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetAlertUuid(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.AlertUuid = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetAliuid(v int64) *ListDisposeStrategyResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetEffectiveStatus(v int32) *ListDisposeStrategyResponseBodyDataResponseData {
	s.EffectiveStatus = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetEntity(v []interface{}) *ListDisposeStrategyResponseBodyDataResponseData {
	s.Entity = v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetEntityId(v int64) *ListDisposeStrategyResponseBodyDataResponseData {
	s.EntityId = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetEntityType(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.EntityType = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetErrorMessage(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.ErrorMessage = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetFinishTime(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.FinishTime = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetGmtCreate(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetGmtModified(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetId(v int64) *ListDisposeStrategyResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetIncidentName(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.IncidentName = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetIncidentUuid(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetPlaybookName(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.PlaybookName = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetPlaybookType(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.PlaybookType = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetPlaybookUuid(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.PlaybookUuid = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetScope(v []interface{}) *ListDisposeStrategyResponseBodyDataResponseData {
	s.Scope = v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetSophonTaskId(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.SophonTaskId = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetStatus(v int32) *ListDisposeStrategyResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetSubAliuid(v int64) *ListDisposeStrategyResponseBodyDataResponseData {
	s.SubAliuid = &v
	return s
}

func (s *ListDisposeStrategyResponseBodyDataResponseData) SetTaskParam(v string) *ListDisposeStrategyResponseBodyDataResponseData {
	s.TaskParam = &v
	return s
}

type ListDisposeStrategyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDisposeStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDisposeStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDisposeStrategyResponse) GoString() string {
	return s.String()
}

func (s *ListDisposeStrategyResponse) SetHeaders(v map[string]*string) *ListDisposeStrategyResponse {
	s.Headers = v
	return s
}

func (s *ListDisposeStrategyResponse) SetStatusCode(v int32) *ListDisposeStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDisposeStrategyResponse) SetBody(v *ListDisposeStrategyResponseBody) *ListDisposeStrategyResponse {
	s.Body = v
	return s
}

type ListOperationRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOperationRequest) GoString() string {
	return s.String()
}

func (s *ListOperationRequest) SetRegionId(v string) *ListOperationRequest {
	s.RegionId = &v
	return s
}

type ListOperationResponseBody struct {
	Code      *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListOperationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DyCode    *string                        `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string                        `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string                        `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOperationResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationResponseBody) SetCode(v int32) *ListOperationResponseBody {
	s.Code = &v
	return s
}

func (s *ListOperationResponseBody) SetData(v *ListOperationResponseBodyData) *ListOperationResponseBody {
	s.Data = v
	return s
}

func (s *ListOperationResponseBody) SetDyCode(v string) *ListOperationResponseBody {
	s.DyCode = &v
	return s
}

func (s *ListOperationResponseBody) SetDyMessage(v string) *ListOperationResponseBody {
	s.DyMessage = &v
	return s
}

func (s *ListOperationResponseBody) SetErrCode(v string) *ListOperationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListOperationResponseBody) SetMessage(v string) *ListOperationResponseBody {
	s.Message = &v
	return s
}

func (s *ListOperationResponseBody) SetRequestId(v string) *ListOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationResponseBody) SetSuccess(v bool) *ListOperationResponseBody {
	s.Success = &v
	return s
}

type ListOperationResponseBodyData struct {
	AdminOrNot    *bool     `json:"AdminOrNot,omitempty" xml:"AdminOrNot,omitempty"`
	OperationList []*string `json:"OperationList,omitempty" xml:"OperationList,omitempty" type:"Repeated"`
}

func (s ListOperationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListOperationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOperationResponseBodyData) SetAdminOrNot(v bool) *ListOperationResponseBodyData {
	s.AdminOrNot = &v
	return s
}

func (s *ListOperationResponseBodyData) SetOperationList(v []*string) *ListOperationResponseBodyData {
	s.OperationList = v
	return s
}

type ListOperationResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOperationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOperationResponse) GoString() string {
	return s.String()
}

func (s *ListOperationResponse) SetHeaders(v map[string]*string) *ListOperationResponse {
	s.Headers = v
	return s
}

func (s *ListOperationResponse) SetStatusCode(v int32) *ListOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationResponse) SetBody(v *ListOperationResponseBody) *ListOperationResponse {
	s.Body = v
	return s
}

type ListQuickQueryRequest struct {
	Offset   *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListQuickQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuickQueryRequest) GoString() string {
	return s.String()
}

func (s *ListQuickQueryRequest) SetOffset(v int32) *ListQuickQueryRequest {
	s.Offset = &v
	return s
}

func (s *ListQuickQueryRequest) SetPageSize(v int32) *ListQuickQueryRequest {
	s.PageSize = &v
	return s
}

func (s *ListQuickQueryRequest) SetRegionId(v string) *ListQuickQueryRequest {
	s.RegionId = &v
	return s
}

type ListQuickQueryResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListQuickQueryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DyCode    *string                         `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string                         `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string                         `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListQuickQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuickQueryResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuickQueryResponseBody) SetCode(v int32) *ListQuickQueryResponseBody {
	s.Code = &v
	return s
}

func (s *ListQuickQueryResponseBody) SetData(v *ListQuickQueryResponseBodyData) *ListQuickQueryResponseBody {
	s.Data = v
	return s
}

func (s *ListQuickQueryResponseBody) SetDyCode(v string) *ListQuickQueryResponseBody {
	s.DyCode = &v
	return s
}

func (s *ListQuickQueryResponseBody) SetDyMessage(v string) *ListQuickQueryResponseBody {
	s.DyMessage = &v
	return s
}

func (s *ListQuickQueryResponseBody) SetErrCode(v string) *ListQuickQueryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListQuickQueryResponseBody) SetMessage(v string) *ListQuickQueryResponseBody {
	s.Message = &v
	return s
}

func (s *ListQuickQueryResponseBody) SetRequestId(v string) *ListQuickQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuickQueryResponseBody) SetSuccess(v bool) *ListQuickQueryResponseBody {
	s.Success = &v
	return s
}

type ListQuickQueryResponseBodyData struct {
	Count          *int32                                          `json:"Count,omitempty" xml:"Count,omitempty"`
	QuickQueryList []*ListQuickQueryResponseBodyDataQuickQueryList `json:"QuickQueryList,omitempty" xml:"QuickQueryList,omitempty" type:"Repeated"`
	Total          *int32                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListQuickQueryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQuickQueryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQuickQueryResponseBodyData) SetCount(v int32) *ListQuickQueryResponseBodyData {
	s.Count = &v
	return s
}

func (s *ListQuickQueryResponseBodyData) SetQuickQueryList(v []*ListQuickQueryResponseBodyDataQuickQueryList) *ListQuickQueryResponseBodyData {
	s.QuickQueryList = v
	return s
}

func (s *ListQuickQueryResponseBodyData) SetTotal(v int32) *ListQuickQueryResponseBodyData {
	s.Total = &v
	return s
}

type ListQuickQueryResponseBodyDataQuickQueryList struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Query       *string `json:"Query,omitempty" xml:"Query,omitempty"`
	SearchName  *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s ListQuickQueryResponseBodyDataQuickQueryList) String() string {
	return tea.Prettify(s)
}

func (s ListQuickQueryResponseBodyDataQuickQueryList) GoString() string {
	return s.String()
}

func (s *ListQuickQueryResponseBodyDataQuickQueryList) SetDisplayName(v string) *ListQuickQueryResponseBodyDataQuickQueryList {
	s.DisplayName = &v
	return s
}

func (s *ListQuickQueryResponseBodyDataQuickQueryList) SetQuery(v string) *ListQuickQueryResponseBodyDataQuickQueryList {
	s.Query = &v
	return s
}

func (s *ListQuickQueryResponseBodyDataQuickQueryList) SetSearchName(v string) *ListQuickQueryResponseBodyDataQuickQueryList {
	s.SearchName = &v
	return s
}

type ListQuickQueryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQuickQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQuickQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQuickQueryResponse) GoString() string {
	return s.String()
}

func (s *ListQuickQueryResponse) SetHeaders(v map[string]*string) *ListQuickQueryResponse {
	s.Headers = v
	return s
}

func (s *ListQuickQueryResponse) SetStatusCode(v int32) *ListQuickQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuickQueryResponse) SetBody(v *ListQuickQueryResponseBody) *ListQuickQueryResponse {
	s.Body = v
	return s
}

type OpenDeliveryRequest struct {
	LogCode     *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpenDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenDeliveryRequest) GoString() string {
	return s.String()
}

func (s *OpenDeliveryRequest) SetLogCode(v string) *OpenDeliveryRequest {
	s.LogCode = &v
	return s
}

func (s *OpenDeliveryRequest) SetProductCode(v string) *OpenDeliveryRequest {
	s.ProductCode = &v
	return s
}

func (s *OpenDeliveryRequest) SetRegionId(v string) *OpenDeliveryRequest {
	s.RegionId = &v
	return s
}

type OpenDeliveryResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	DyCode    *string `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OpenDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *OpenDeliveryResponseBody) SetCode(v int32) *OpenDeliveryResponseBody {
	s.Code = &v
	return s
}

func (s *OpenDeliveryResponseBody) SetData(v bool) *OpenDeliveryResponseBody {
	s.Data = &v
	return s
}

func (s *OpenDeliveryResponseBody) SetDyCode(v string) *OpenDeliveryResponseBody {
	s.DyCode = &v
	return s
}

func (s *OpenDeliveryResponseBody) SetDyMessage(v string) *OpenDeliveryResponseBody {
	s.DyMessage = &v
	return s
}

func (s *OpenDeliveryResponseBody) SetErrCode(v string) *OpenDeliveryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *OpenDeliveryResponseBody) SetMessage(v string) *OpenDeliveryResponseBody {
	s.Message = &v
	return s
}

func (s *OpenDeliveryResponseBody) SetRequestId(v string) *OpenDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenDeliveryResponseBody) SetSuccess(v bool) *OpenDeliveryResponseBody {
	s.Success = &v
	return s
}

type OpenDeliveryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenDeliveryResponse) GoString() string {
	return s.String()
}

func (s *OpenDeliveryResponse) SetHeaders(v map[string]*string) *OpenDeliveryResponse {
	s.Headers = v
	return s
}

func (s *OpenDeliveryResponse) SetStatusCode(v int32) *OpenDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenDeliveryResponse) SetBody(v *OpenDeliveryResponseBody) *OpenDeliveryResponse {
	s.Body = v
	return s
}

type PostAutomateResponseConfigRequest struct {
	ActionConfig       *string `json:"ActionConfig,omitempty" xml:"ActionConfig,omitempty"`
	ActionType         *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	AutoResponseType   *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
	ExecutionCondition *string `json:"ExecutionCondition,omitempty" xml:"ExecutionCondition,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleName           *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	SubUserId          *int64  `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s PostAutomateResponseConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s PostAutomateResponseConfigRequest) GoString() string {
	return s.String()
}

func (s *PostAutomateResponseConfigRequest) SetActionConfig(v string) *PostAutomateResponseConfigRequest {
	s.ActionConfig = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetActionType(v string) *PostAutomateResponseConfigRequest {
	s.ActionType = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetAutoResponseType(v string) *PostAutomateResponseConfigRequest {
	s.AutoResponseType = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetExecutionCondition(v string) *PostAutomateResponseConfigRequest {
	s.ExecutionCondition = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetId(v int64) *PostAutomateResponseConfigRequest {
	s.Id = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetRegionId(v string) *PostAutomateResponseConfigRequest {
	s.RegionId = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetRuleName(v string) *PostAutomateResponseConfigRequest {
	s.RuleName = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetSubUserId(v int64) *PostAutomateResponseConfigRequest {
	s.SubUserId = &v
	return s
}

type PostAutomateResponseConfigResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostAutomateResponseConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostAutomateResponseConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PostAutomateResponseConfigResponseBody) SetCode(v int32) *PostAutomateResponseConfigResponseBody {
	s.Code = &v
	return s
}

func (s *PostAutomateResponseConfigResponseBody) SetData(v string) *PostAutomateResponseConfigResponseBody {
	s.Data = &v
	return s
}

func (s *PostAutomateResponseConfigResponseBody) SetMessage(v string) *PostAutomateResponseConfigResponseBody {
	s.Message = &v
	return s
}

func (s *PostAutomateResponseConfigResponseBody) SetRequestId(v string) *PostAutomateResponseConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostAutomateResponseConfigResponseBody) SetSuccess(v bool) *PostAutomateResponseConfigResponseBody {
	s.Success = &v
	return s
}

type PostAutomateResponseConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PostAutomateResponseConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PostAutomateResponseConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s PostAutomateResponseConfigResponse) GoString() string {
	return s.String()
}

func (s *PostAutomateResponseConfigResponse) SetHeaders(v map[string]*string) *PostAutomateResponseConfigResponse {
	s.Headers = v
	return s
}

func (s *PostAutomateResponseConfigResponse) SetStatusCode(v int32) *PostAutomateResponseConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PostAutomateResponseConfigResponse) SetBody(v *PostAutomateResponseConfigResponseBody) *PostAutomateResponseConfigResponse {
	s.Body = v
	return s
}

type PostCustomizeRuleRequest struct {
	AlertType           *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AlertTypeMds        *string `json:"AlertTypeMds,omitempty" xml:"AlertTypeMds,omitempty"`
	EventTransferExt    *string `json:"EventTransferExt,omitempty" xml:"EventTransferExt,omitempty"`
	EventTransferSwitch *int32  `json:"EventTransferSwitch,omitempty" xml:"EventTransferSwitch,omitempty"`
	EventTransferType   *string `json:"EventTransferType,omitempty" xml:"EventTransferType,omitempty"`
	Id                  *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LogSource           *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	LogSourceMds        *string `json:"LogSourceMds,omitempty" xml:"LogSourceMds,omitempty"`
	LogType             *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	LogTypeMds          *string `json:"LogTypeMds,omitempty" xml:"LogTypeMds,omitempty"`
	QueryCycle          *string `json:"QueryCycle,omitempty" xml:"QueryCycle,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleCondition       *string `json:"RuleCondition,omitempty" xml:"RuleCondition,omitempty"`
	RuleDesc            *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	RuleGroup           *string `json:"RuleGroup,omitempty" xml:"RuleGroup,omitempty"`
	RuleName            *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleThreshold       *string `json:"RuleThreshold,omitempty" xml:"RuleThreshold,omitempty"`
	ThreatLevel         *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
}

func (s PostCustomizeRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s PostCustomizeRuleRequest) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleRequest) SetAlertType(v string) *PostCustomizeRuleRequest {
	s.AlertType = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetAlertTypeMds(v string) *PostCustomizeRuleRequest {
	s.AlertTypeMds = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetEventTransferExt(v string) *PostCustomizeRuleRequest {
	s.EventTransferExt = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetEventTransferSwitch(v int32) *PostCustomizeRuleRequest {
	s.EventTransferSwitch = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetEventTransferType(v string) *PostCustomizeRuleRequest {
	s.EventTransferType = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetId(v int64) *PostCustomizeRuleRequest {
	s.Id = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetLogSource(v string) *PostCustomizeRuleRequest {
	s.LogSource = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetLogSourceMds(v string) *PostCustomizeRuleRequest {
	s.LogSourceMds = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetLogType(v string) *PostCustomizeRuleRequest {
	s.LogType = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetLogTypeMds(v string) *PostCustomizeRuleRequest {
	s.LogTypeMds = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetQueryCycle(v string) *PostCustomizeRuleRequest {
	s.QueryCycle = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRegionId(v string) *PostCustomizeRuleRequest {
	s.RegionId = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRuleCondition(v string) *PostCustomizeRuleRequest {
	s.RuleCondition = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRuleDesc(v string) *PostCustomizeRuleRequest {
	s.RuleDesc = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRuleGroup(v string) *PostCustomizeRuleRequest {
	s.RuleGroup = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRuleName(v string) *PostCustomizeRuleRequest {
	s.RuleName = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRuleThreshold(v string) *PostCustomizeRuleRequest {
	s.RuleThreshold = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetThreatLevel(v string) *PostCustomizeRuleRequest {
	s.ThreatLevel = &v
	return s
}

type PostCustomizeRuleResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PostCustomizeRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostCustomizeRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostCustomizeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleResponseBody) SetCode(v int32) *PostCustomizeRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PostCustomizeRuleResponseBody) SetData(v *PostCustomizeRuleResponseBodyData) *PostCustomizeRuleResponseBody {
	s.Data = v
	return s
}

func (s *PostCustomizeRuleResponseBody) SetMessage(v string) *PostCustomizeRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PostCustomizeRuleResponseBody) SetRequestId(v string) *PostCustomizeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostCustomizeRuleResponseBody) SetSuccess(v bool) *PostCustomizeRuleResponseBody {
	s.Success = &v
	return s
}

type PostCustomizeRuleResponseBodyData struct {
	AlertType           *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AlertTypeMds        *string `json:"AlertTypeMds,omitempty" xml:"AlertTypeMds,omitempty"`
	Aliuid              *int64  `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	EventTransferExt    *string `json:"EventTransferExt,omitempty" xml:"EventTransferExt,omitempty"`
	EventTransferSwitch *int32  `json:"EventTransferSwitch,omitempty" xml:"EventTransferSwitch,omitempty"`
	EventTransferType   *string `json:"EventTransferType,omitempty" xml:"EventTransferType,omitempty"`
	GmtCreate           *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified         *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                  *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LogSource           *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	LogSourceMds        *string `json:"LogSourceMds,omitempty" xml:"LogSourceMds,omitempty"`
	LogType             *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	LogTypeMds          *string `json:"LogTypeMds,omitempty" xml:"LogTypeMds,omitempty"`
	QueryCycle          *string `json:"QueryCycle,omitempty" xml:"QueryCycle,omitempty"`
	RuleCondition       *string `json:"RuleCondition,omitempty" xml:"RuleCondition,omitempty"`
	RuleDesc            *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	RuleGroup           *string `json:"RuleGroup,omitempty" xml:"RuleGroup,omitempty"`
	RuleName            *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleThreshold       *string `json:"RuleThreshold,omitempty" xml:"RuleThreshold,omitempty"`
	RuleType            *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Status              *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ThreatLevel         *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
}

func (s PostCustomizeRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PostCustomizeRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleResponseBodyData) SetAlertType(v string) *PostCustomizeRuleResponseBodyData {
	s.AlertType = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetAlertTypeMds(v string) *PostCustomizeRuleResponseBodyData {
	s.AlertTypeMds = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetAliuid(v int64) *PostCustomizeRuleResponseBodyData {
	s.Aliuid = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetEventTransferExt(v string) *PostCustomizeRuleResponseBodyData {
	s.EventTransferExt = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetEventTransferSwitch(v int32) *PostCustomizeRuleResponseBodyData {
	s.EventTransferSwitch = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetEventTransferType(v string) *PostCustomizeRuleResponseBodyData {
	s.EventTransferType = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetGmtCreate(v string) *PostCustomizeRuleResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetGmtModified(v string) *PostCustomizeRuleResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetId(v int64) *PostCustomizeRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetLogSource(v string) *PostCustomizeRuleResponseBodyData {
	s.LogSource = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetLogSourceMds(v string) *PostCustomizeRuleResponseBodyData {
	s.LogSourceMds = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetLogType(v string) *PostCustomizeRuleResponseBodyData {
	s.LogType = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetLogTypeMds(v string) *PostCustomizeRuleResponseBodyData {
	s.LogTypeMds = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetQueryCycle(v string) *PostCustomizeRuleResponseBodyData {
	s.QueryCycle = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleCondition(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleCondition = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleDesc(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleDesc = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleGroup(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleGroup = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleName(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleThreshold(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleThreshold = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleType(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleType = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetStatus(v int32) *PostCustomizeRuleResponseBodyData {
	s.Status = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetThreatLevel(v string) *PostCustomizeRuleResponseBodyData {
	s.ThreatLevel = &v
	return s
}

type PostCustomizeRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PostCustomizeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PostCustomizeRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PostCustomizeRuleResponse) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleResponse) SetHeaders(v map[string]*string) *PostCustomizeRuleResponse {
	s.Headers = v
	return s
}

func (s *PostCustomizeRuleResponse) SetStatusCode(v int32) *PostCustomizeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PostCustomizeRuleResponse) SetBody(v *PostCustomizeRuleResponseBody) *PostCustomizeRuleResponse {
	s.Body = v
	return s
}

type PostCustomizeRuleTestRequest struct {
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SimulatedData *string `json:"SimulatedData,omitempty" xml:"SimulatedData,omitempty"`
	TestType      *string `json:"TestType,omitempty" xml:"TestType,omitempty"`
}

func (s PostCustomizeRuleTestRequest) String() string {
	return tea.Prettify(s)
}

func (s PostCustomizeRuleTestRequest) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleTestRequest) SetId(v int64) *PostCustomizeRuleTestRequest {
	s.Id = &v
	return s
}

func (s *PostCustomizeRuleTestRequest) SetRegionId(v string) *PostCustomizeRuleTestRequest {
	s.RegionId = &v
	return s
}

func (s *PostCustomizeRuleTestRequest) SetSimulatedData(v string) *PostCustomizeRuleTestRequest {
	s.SimulatedData = &v
	return s
}

func (s *PostCustomizeRuleTestRequest) SetTestType(v string) *PostCustomizeRuleTestRequest {
	s.TestType = &v
	return s
}

type PostCustomizeRuleTestResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostCustomizeRuleTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostCustomizeRuleTestResponseBody) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleTestResponseBody) SetCode(v int32) *PostCustomizeRuleTestResponseBody {
	s.Code = &v
	return s
}

func (s *PostCustomizeRuleTestResponseBody) SetData(v interface{}) *PostCustomizeRuleTestResponseBody {
	s.Data = v
	return s
}

func (s *PostCustomizeRuleTestResponseBody) SetMessage(v string) *PostCustomizeRuleTestResponseBody {
	s.Message = &v
	return s
}

func (s *PostCustomizeRuleTestResponseBody) SetRequestId(v string) *PostCustomizeRuleTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostCustomizeRuleTestResponseBody) SetSuccess(v bool) *PostCustomizeRuleTestResponseBody {
	s.Success = &v
	return s
}

type PostCustomizeRuleTestResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PostCustomizeRuleTestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PostCustomizeRuleTestResponse) String() string {
	return tea.Prettify(s)
}

func (s PostCustomizeRuleTestResponse) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleTestResponse) SetHeaders(v map[string]*string) *PostCustomizeRuleTestResponse {
	s.Headers = v
	return s
}

func (s *PostCustomizeRuleTestResponse) SetStatusCode(v int32) *PostCustomizeRuleTestResponse {
	s.StatusCode = &v
	return s
}

func (s *PostCustomizeRuleTestResponse) SetBody(v *PostCustomizeRuleTestResponseBody) *PostCustomizeRuleTestResponse {
	s.Body = v
	return s
}

type PostEventDisposeAndWhiteruleListRequest struct {
	EventDispose *string `json:"EventDispose,omitempty" xml:"EventDispose,omitempty"`
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	ReceiverInfo *string `json:"ReceiverInfo,omitempty" xml:"ReceiverInfo,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PostEventDisposeAndWhiteruleListRequest) String() string {
	return tea.Prettify(s)
}

func (s PostEventDisposeAndWhiteruleListRequest) GoString() string {
	return s.String()
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetEventDispose(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.EventDispose = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetIncidentUuid(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.IncidentUuid = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetReceiverInfo(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.ReceiverInfo = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetRegionId(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.RegionId = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetRemark(v string) *PostEventDisposeAndWhiteruleListRequest {
	s.Remark = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListRequest) SetStatus(v int32) *PostEventDisposeAndWhiteruleListRequest {
	s.Status = &v
	return s
}

type PostEventDisposeAndWhiteruleListResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostEventDisposeAndWhiteruleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostEventDisposeAndWhiteruleListResponseBody) GoString() string {
	return s.String()
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) SetCode(v int32) *PostEventDisposeAndWhiteruleListResponseBody {
	s.Code = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) SetData(v string) *PostEventDisposeAndWhiteruleListResponseBody {
	s.Data = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) SetMessage(v string) *PostEventDisposeAndWhiteruleListResponseBody {
	s.Message = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) SetRequestId(v string) *PostEventDisposeAndWhiteruleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponseBody) SetSuccess(v bool) *PostEventDisposeAndWhiteruleListResponseBody {
	s.Success = &v
	return s
}

type PostEventDisposeAndWhiteruleListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PostEventDisposeAndWhiteruleListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PostEventDisposeAndWhiteruleListResponse) String() string {
	return tea.Prettify(s)
}

func (s PostEventDisposeAndWhiteruleListResponse) GoString() string {
	return s.String()
}

func (s *PostEventDisposeAndWhiteruleListResponse) SetHeaders(v map[string]*string) *PostEventDisposeAndWhiteruleListResponse {
	s.Headers = v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponse) SetStatusCode(v int32) *PostEventDisposeAndWhiteruleListResponse {
	s.StatusCode = &v
	return s
}

func (s *PostEventDisposeAndWhiteruleListResponse) SetBody(v *PostEventDisposeAndWhiteruleListResponseBody) *PostEventDisposeAndWhiteruleListResponse {
	s.Body = v
	return s
}

type PostEventWhiteruleListRequest struct {
	IncidentUuid  *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WhiteruleList *string `json:"WhiteruleList,omitempty" xml:"WhiteruleList,omitempty"`
}

func (s PostEventWhiteruleListRequest) String() string {
	return tea.Prettify(s)
}

func (s PostEventWhiteruleListRequest) GoString() string {
	return s.String()
}

func (s *PostEventWhiteruleListRequest) SetIncidentUuid(v string) *PostEventWhiteruleListRequest {
	s.IncidentUuid = &v
	return s
}

func (s *PostEventWhiteruleListRequest) SetRegionId(v string) *PostEventWhiteruleListRequest {
	s.RegionId = &v
	return s
}

func (s *PostEventWhiteruleListRequest) SetWhiteruleList(v string) *PostEventWhiteruleListRequest {
	s.WhiteruleList = &v
	return s
}

type PostEventWhiteruleListResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostEventWhiteruleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostEventWhiteruleListResponseBody) GoString() string {
	return s.String()
}

func (s *PostEventWhiteruleListResponseBody) SetCode(v int32) *PostEventWhiteruleListResponseBody {
	s.Code = &v
	return s
}

func (s *PostEventWhiteruleListResponseBody) SetData(v string) *PostEventWhiteruleListResponseBody {
	s.Data = &v
	return s
}

func (s *PostEventWhiteruleListResponseBody) SetMessage(v string) *PostEventWhiteruleListResponseBody {
	s.Message = &v
	return s
}

func (s *PostEventWhiteruleListResponseBody) SetRequestId(v string) *PostEventWhiteruleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostEventWhiteruleListResponseBody) SetSuccess(v bool) *PostEventWhiteruleListResponseBody {
	s.Success = &v
	return s
}

type PostEventWhiteruleListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PostEventWhiteruleListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PostEventWhiteruleListResponse) String() string {
	return tea.Prettify(s)
}

func (s PostEventWhiteruleListResponse) GoString() string {
	return s.String()
}

func (s *PostEventWhiteruleListResponse) SetHeaders(v map[string]*string) *PostEventWhiteruleListResponse {
	s.Headers = v
	return s
}

func (s *PostEventWhiteruleListResponse) SetStatusCode(v int32) *PostEventWhiteruleListResponse {
	s.StatusCode = &v
	return s
}

func (s *PostEventWhiteruleListResponse) SetBody(v *PostEventWhiteruleListResponseBody) *PostEventWhiteruleListResponse {
	s.Body = v
	return s
}

type PostFinishCustomizeRuleTestRequest struct {
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s PostFinishCustomizeRuleTestRequest) String() string {
	return tea.Prettify(s)
}

func (s PostFinishCustomizeRuleTestRequest) GoString() string {
	return s.String()
}

func (s *PostFinishCustomizeRuleTestRequest) SetId(v int64) *PostFinishCustomizeRuleTestRequest {
	s.Id = &v
	return s
}

func (s *PostFinishCustomizeRuleTestRequest) SetRegionId(v string) *PostFinishCustomizeRuleTestRequest {
	s.RegionId = &v
	return s
}

type PostFinishCustomizeRuleTestResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostFinishCustomizeRuleTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostFinishCustomizeRuleTestResponseBody) GoString() string {
	return s.String()
}

func (s *PostFinishCustomizeRuleTestResponseBody) SetCode(v int32) *PostFinishCustomizeRuleTestResponseBody {
	s.Code = &v
	return s
}

func (s *PostFinishCustomizeRuleTestResponseBody) SetData(v interface{}) *PostFinishCustomizeRuleTestResponseBody {
	s.Data = v
	return s
}

func (s *PostFinishCustomizeRuleTestResponseBody) SetMessage(v string) *PostFinishCustomizeRuleTestResponseBody {
	s.Message = &v
	return s
}

func (s *PostFinishCustomizeRuleTestResponseBody) SetRequestId(v string) *PostFinishCustomizeRuleTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostFinishCustomizeRuleTestResponseBody) SetSuccess(v bool) *PostFinishCustomizeRuleTestResponseBody {
	s.Success = &v
	return s
}

type PostFinishCustomizeRuleTestResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PostFinishCustomizeRuleTestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PostFinishCustomizeRuleTestResponse) String() string {
	return tea.Prettify(s)
}

func (s PostFinishCustomizeRuleTestResponse) GoString() string {
	return s.String()
}

func (s *PostFinishCustomizeRuleTestResponse) SetHeaders(v map[string]*string) *PostFinishCustomizeRuleTestResponse {
	s.Headers = v
	return s
}

func (s *PostFinishCustomizeRuleTestResponse) SetStatusCode(v int32) *PostFinishCustomizeRuleTestResponse {
	s.StatusCode = &v
	return s
}

func (s *PostFinishCustomizeRuleTestResponse) SetBody(v *PostFinishCustomizeRuleTestResponseBody) *PostFinishCustomizeRuleTestResponse {
	s.Body = v
	return s
}

type PostRuleStatusChangeRequest struct {
	Ids      *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	InUse    *bool   `json:"InUse,omitempty" xml:"InUse,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s PostRuleStatusChangeRequest) String() string {
	return tea.Prettify(s)
}

func (s PostRuleStatusChangeRequest) GoString() string {
	return s.String()
}

func (s *PostRuleStatusChangeRequest) SetIds(v string) *PostRuleStatusChangeRequest {
	s.Ids = &v
	return s
}

func (s *PostRuleStatusChangeRequest) SetInUse(v bool) *PostRuleStatusChangeRequest {
	s.InUse = &v
	return s
}

func (s *PostRuleStatusChangeRequest) SetRegionId(v string) *PostRuleStatusChangeRequest {
	s.RegionId = &v
	return s
}

func (s *PostRuleStatusChangeRequest) SetRuleType(v string) *PostRuleStatusChangeRequest {
	s.RuleType = &v
	return s
}

type PostRuleStatusChangeResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PostRuleStatusChangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PostRuleStatusChangeResponseBody) GoString() string {
	return s.String()
}

func (s *PostRuleStatusChangeResponseBody) SetCode(v int32) *PostRuleStatusChangeResponseBody {
	s.Code = &v
	return s
}

func (s *PostRuleStatusChangeResponseBody) SetData(v interface{}) *PostRuleStatusChangeResponseBody {
	s.Data = v
	return s
}

func (s *PostRuleStatusChangeResponseBody) SetMessage(v string) *PostRuleStatusChangeResponseBody {
	s.Message = &v
	return s
}

func (s *PostRuleStatusChangeResponseBody) SetRequestId(v string) *PostRuleStatusChangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostRuleStatusChangeResponseBody) SetSuccess(v bool) *PostRuleStatusChangeResponseBody {
	s.Success = &v
	return s
}

type PostRuleStatusChangeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PostRuleStatusChangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PostRuleStatusChangeResponse) String() string {
	return tea.Prettify(s)
}

func (s PostRuleStatusChangeResponse) GoString() string {
	return s.String()
}

func (s *PostRuleStatusChangeResponse) SetHeaders(v map[string]*string) *PostRuleStatusChangeResponse {
	s.Headers = v
	return s
}

func (s *PostRuleStatusChangeResponse) SetStatusCode(v int32) *PostRuleStatusChangeResponse {
	s.StatusCode = &v
	return s
}

func (s *PostRuleStatusChangeResponse) SetBody(v *PostRuleStatusChangeResponseBody) *PostRuleStatusChangeResponse {
	s.Body = v
	return s
}

type RestoreCapacityRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RestoreCapacityRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreCapacityRequest) GoString() string {
	return s.String()
}

func (s *RestoreCapacityRequest) SetRegionId(v string) *RestoreCapacityRequest {
	s.RegionId = &v
	return s
}

type RestoreCapacityResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	DyCode    *string `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestoreCapacityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreCapacityResponseBody) SetCode(v int32) *RestoreCapacityResponseBody {
	s.Code = &v
	return s
}

func (s *RestoreCapacityResponseBody) SetData(v bool) *RestoreCapacityResponseBody {
	s.Data = &v
	return s
}

func (s *RestoreCapacityResponseBody) SetDyCode(v string) *RestoreCapacityResponseBody {
	s.DyCode = &v
	return s
}

func (s *RestoreCapacityResponseBody) SetDyMessage(v string) *RestoreCapacityResponseBody {
	s.DyMessage = &v
	return s
}

func (s *RestoreCapacityResponseBody) SetErrCode(v string) *RestoreCapacityResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RestoreCapacityResponseBody) SetMessage(v string) *RestoreCapacityResponseBody {
	s.Message = &v
	return s
}

func (s *RestoreCapacityResponseBody) SetRequestId(v string) *RestoreCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreCapacityResponseBody) SetSuccess(v bool) *RestoreCapacityResponseBody {
	s.Success = &v
	return s
}

type RestoreCapacityResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestoreCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestoreCapacityResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreCapacityResponse) GoString() string {
	return s.String()
}

func (s *RestoreCapacityResponse) SetHeaders(v map[string]*string) *RestoreCapacityResponse {
	s.Headers = v
	return s
}

func (s *RestoreCapacityResponse) SetStatusCode(v int32) *RestoreCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreCapacityResponse) SetBody(v *RestoreCapacityResponseBody) *RestoreCapacityResponse {
	s.Body = v
	return s
}

type SaveQuickQueryRequest struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Query       *string `json:"Query,omitempty" xml:"Query,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SaveQuickQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveQuickQueryRequest) GoString() string {
	return s.String()
}

func (s *SaveQuickQueryRequest) SetDisplayName(v string) *SaveQuickQueryRequest {
	s.DisplayName = &v
	return s
}

func (s *SaveQuickQueryRequest) SetQuery(v string) *SaveQuickQueryRequest {
	s.Query = &v
	return s
}

func (s *SaveQuickQueryRequest) SetRegionId(v string) *SaveQuickQueryRequest {
	s.RegionId = &v
	return s
}

type SaveQuickQueryResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	DyCode    *string `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveQuickQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveQuickQueryResponseBody) GoString() string {
	return s.String()
}

func (s *SaveQuickQueryResponseBody) SetCode(v int32) *SaveQuickQueryResponseBody {
	s.Code = &v
	return s
}

func (s *SaveQuickQueryResponseBody) SetData(v bool) *SaveQuickQueryResponseBody {
	s.Data = &v
	return s
}

func (s *SaveQuickQueryResponseBody) SetDyCode(v string) *SaveQuickQueryResponseBody {
	s.DyCode = &v
	return s
}

func (s *SaveQuickQueryResponseBody) SetDyMessage(v string) *SaveQuickQueryResponseBody {
	s.DyMessage = &v
	return s
}

func (s *SaveQuickQueryResponseBody) SetErrCode(v string) *SaveQuickQueryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SaveQuickQueryResponseBody) SetMessage(v string) *SaveQuickQueryResponseBody {
	s.Message = &v
	return s
}

func (s *SaveQuickQueryResponseBody) SetRequestId(v string) *SaveQuickQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveQuickQueryResponseBody) SetSuccess(v bool) *SaveQuickQueryResponseBody {
	s.Success = &v
	return s
}

type SaveQuickQueryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveQuickQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveQuickQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveQuickQueryResponse) GoString() string {
	return s.String()
}

func (s *SaveQuickQueryResponse) SetHeaders(v map[string]*string) *SaveQuickQueryResponse {
	s.Headers = v
	return s
}

func (s *SaveQuickQueryResponse) SetStatusCode(v int32) *SaveQuickQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveQuickQueryResponse) SetBody(v *SaveQuickQueryResponseBody) *SaveQuickQueryResponse {
	s.Body = v
	return s
}

type SetStorageRequest struct {
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Ttl      *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s SetStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s SetStorageRequest) GoString() string {
	return s.String()
}

func (s *SetStorageRequest) SetRegion(v string) *SetStorageRequest {
	s.Region = &v
	return s
}

func (s *SetStorageRequest) SetRegionId(v string) *SetStorageRequest {
	s.RegionId = &v
	return s
}

func (s *SetStorageRequest) SetTtl(v int32) *SetStorageRequest {
	s.Ttl = &v
	return s
}

type SetStorageResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	DyCode    *string `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetStorageResponseBody) GoString() string {
	return s.String()
}

func (s *SetStorageResponseBody) SetCode(v int32) *SetStorageResponseBody {
	s.Code = &v
	return s
}

func (s *SetStorageResponseBody) SetData(v bool) *SetStorageResponseBody {
	s.Data = &v
	return s
}

func (s *SetStorageResponseBody) SetDyCode(v string) *SetStorageResponseBody {
	s.DyCode = &v
	return s
}

func (s *SetStorageResponseBody) SetDyMessage(v string) *SetStorageResponseBody {
	s.DyMessage = &v
	return s
}

func (s *SetStorageResponseBody) SetErrCode(v string) *SetStorageResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SetStorageResponseBody) SetMessage(v string) *SetStorageResponseBody {
	s.Message = &v
	return s
}

func (s *SetStorageResponseBody) SetRequestId(v string) *SetStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetStorageResponseBody) SetSuccess(v bool) *SetStorageResponseBody {
	s.Success = &v
	return s
}

type SetStorageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s SetStorageResponse) GoString() string {
	return s.String()
}

func (s *SetStorageResponse) SetHeaders(v map[string]*string) *SetStorageResponse {
	s.Headers = v
	return s
}

func (s *SetStorageResponse) SetStatusCode(v int32) *SetStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *SetStorageResponse) SetBody(v *SetStorageResponseBody) *SetStorageResponse {
	s.Body = v
	return s
}

type ShowQuickAnalysisRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ShowQuickAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s ShowQuickAnalysisRequest) GoString() string {
	return s.String()
}

func (s *ShowQuickAnalysisRequest) SetRegionId(v string) *ShowQuickAnalysisRequest {
	s.RegionId = &v
	return s
}

type ShowQuickAnalysisResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ShowQuickAnalysisResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DyCode    *string                            `json:"DyCode,omitempty" xml:"DyCode,omitempty"`
	DyMessage *string                            `json:"DyMessage,omitempty" xml:"DyMessage,omitempty"`
	ErrCode   *string                            `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ShowQuickAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ShowQuickAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *ShowQuickAnalysisResponseBody) SetCode(v int32) *ShowQuickAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *ShowQuickAnalysisResponseBody) SetData(v *ShowQuickAnalysisResponseBodyData) *ShowQuickAnalysisResponseBody {
	s.Data = v
	return s
}

func (s *ShowQuickAnalysisResponseBody) SetDyCode(v string) *ShowQuickAnalysisResponseBody {
	s.DyCode = &v
	return s
}

func (s *ShowQuickAnalysisResponseBody) SetDyMessage(v string) *ShowQuickAnalysisResponseBody {
	s.DyMessage = &v
	return s
}

func (s *ShowQuickAnalysisResponseBody) SetErrCode(v string) *ShowQuickAnalysisResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ShowQuickAnalysisResponseBody) SetMessage(v string) *ShowQuickAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *ShowQuickAnalysisResponseBody) SetRequestId(v string) *ShowQuickAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *ShowQuickAnalysisResponseBody) SetSuccess(v bool) *ShowQuickAnalysisResponseBody {
	s.Success = &v
	return s
}

type ShowQuickAnalysisResponseBodyData struct {
	IndexList []*string `json:"IndexList,omitempty" xml:"IndexList,omitempty" type:"Repeated"`
}

func (s ShowQuickAnalysisResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ShowQuickAnalysisResponseBodyData) GoString() string {
	return s.String()
}

func (s *ShowQuickAnalysisResponseBodyData) SetIndexList(v []*string) *ShowQuickAnalysisResponseBodyData {
	s.IndexList = v
	return s
}

type ShowQuickAnalysisResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ShowQuickAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ShowQuickAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s ShowQuickAnalysisResponse) GoString() string {
	return s.String()
}

func (s *ShowQuickAnalysisResponse) SetHeaders(v map[string]*string) *ShowQuickAnalysisResponse {
	s.Headers = v
	return s
}

func (s *ShowQuickAnalysisResponse) SetStatusCode(v int32) *ShowQuickAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *ShowQuickAnalysisResponse) SetBody(v *ShowQuickAnalysisResponseBody) *ShowQuickAnalysisResponse {
	s.Body = v
	return s
}

type UpdateAutomateResponseConfigStatusRequest struct {
	Ids      *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	InUse    *bool   `json:"InUse,omitempty" xml:"InUse,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateAutomateResponseConfigStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutomateResponseConfigStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutomateResponseConfigStatusRequest) SetIds(v string) *UpdateAutomateResponseConfigStatusRequest {
	s.Ids = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusRequest) SetInUse(v bool) *UpdateAutomateResponseConfigStatusRequest {
	s.InUse = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusRequest) SetRegionId(v string) *UpdateAutomateResponseConfigStatusRequest {
	s.RegionId = &v
	return s
}

type UpdateAutomateResponseConfigStatusResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAutomateResponseConfigStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutomateResponseConfigStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) SetCode(v int32) *UpdateAutomateResponseConfigStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) SetData(v string) *UpdateAutomateResponseConfigStatusResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) SetMessage(v string) *UpdateAutomateResponseConfigStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) SetRequestId(v string) *UpdateAutomateResponseConfigStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) SetSuccess(v bool) *UpdateAutomateResponseConfigStatusResponseBody {
	s.Success = &v
	return s
}

type UpdateAutomateResponseConfigStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAutomateResponseConfigStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAutomateResponseConfigStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutomateResponseConfigStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutomateResponseConfigStatusResponse) SetHeaders(v map[string]*string) *UpdateAutomateResponseConfigStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponse) SetStatusCode(v int32) *UpdateAutomateResponseConfigStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponse) SetBody(v *UpdateAutomateResponseConfigStatusResponseBody) *UpdateAutomateResponseConfigStatusResponse {
	s.Body = v
	return s
}

type UpdateWhiteRuleListRequest struct {
	Expression   *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WhiteRuleId  *int64  `json:"WhiteRuleId,omitempty" xml:"WhiteRuleId,omitempty"`
}

func (s UpdateWhiteRuleListRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWhiteRuleListRequest) GoString() string {
	return s.String()
}

func (s *UpdateWhiteRuleListRequest) SetExpression(v string) *UpdateWhiteRuleListRequest {
	s.Expression = &v
	return s
}

func (s *UpdateWhiteRuleListRequest) SetIncidentUuid(v string) *UpdateWhiteRuleListRequest {
	s.IncidentUuid = &v
	return s
}

func (s *UpdateWhiteRuleListRequest) SetRegionId(v string) *UpdateWhiteRuleListRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateWhiteRuleListRequest) SetWhiteRuleId(v int64) *UpdateWhiteRuleListRequest {
	s.WhiteRuleId = &v
	return s
}

type UpdateWhiteRuleListResponseBody struct {
	Code      *int32      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWhiteRuleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWhiteRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWhiteRuleListResponseBody) SetCode(v int32) *UpdateWhiteRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWhiteRuleListResponseBody) SetData(v interface{}) *UpdateWhiteRuleListResponseBody {
	s.Data = v
	return s
}

func (s *UpdateWhiteRuleListResponseBody) SetMessage(v string) *UpdateWhiteRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWhiteRuleListResponseBody) SetRequestId(v string) *UpdateWhiteRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWhiteRuleListResponseBody) SetSuccess(v bool) *UpdateWhiteRuleListResponseBody {
	s.Success = &v
	return s
}

type UpdateWhiteRuleListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateWhiteRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWhiteRuleListResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWhiteRuleListResponse) GoString() string {
	return s.String()
}

func (s *UpdateWhiteRuleListResponse) SetHeaders(v map[string]*string) *UpdateWhiteRuleListResponse {
	s.Headers = v
	return s
}

func (s *UpdateWhiteRuleListResponse) SetStatusCode(v int32) *UpdateWhiteRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWhiteRuleListResponse) SetBody(v *UpdateWhiteRuleListResponseBody) *UpdateWhiteRuleListResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloud-siem"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BatchJobCheckWithOptions(request *BatchJobCheckRequest, runtime *util.RuntimeOptions) (_result *BatchJobCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SubmitId)) {
		body["SubmitId"] = request.SubmitId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchJobCheck"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchJobCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchJobCheck(request *BatchJobCheckRequest) (_result *BatchJobCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchJobCheckResponse{}
	_body, _err := client.BatchJobCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchJobSubmitWithOptions(request *BatchJobSubmitRequest, runtime *util.RuntimeOptions) (_result *BatchJobSubmitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JsonConfig)) {
		body["JsonConfig"] = request.JsonConfig
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchJobSubmit"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchJobSubmitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchJobSubmit(request *BatchJobSubmitRequest) (_result *BatchJobSubmitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchJobSubmitResponse{}
	_body, _err := client.BatchJobSubmitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseDeliveryWithOptions(request *CloseDeliveryRequest, runtime *util.RuntimeOptions) (_result *CloseDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogCode)) {
		body["LogCode"] = request.LogCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseDelivery"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseDelivery(request *CloseDeliveryRequest) (_result *CloseDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseDeliveryResponse{}
	_body, _err := client.CloseDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAutomateResponseConfigWithOptions(request *DeleteAutomateResponseConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteAutomateResponseConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAutomateResponseConfig"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAutomateResponseConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAutomateResponseConfig(request *DeleteAutomateResponseConfigRequest) (_result *DeleteAutomateResponseConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAutomateResponseConfigResponse{}
	_body, _err := client.DeleteAutomateResponseConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCustomizeRuleWithOptions(request *DeleteCustomizeRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomizeRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomizeRule"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomizeRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCustomizeRule(request *DeleteCustomizeRuleRequest) (_result *DeleteCustomizeRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomizeRuleResponse{}
	_body, _err := client.DeleteCustomizeRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteQuickQueryWithOptions(request *DeleteQuickQueryRequest, runtime *util.RuntimeOptions) (_result *DeleteQuickQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchName)) {
		body["SearchName"] = request.SearchName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteQuickQuery"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteQuickQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteQuickQuery(request *DeleteQuickQueryRequest) (_result *DeleteQuickQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteQuickQueryResponse{}
	_body, _err := client.DeleteQuickQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWhiteRuleListWithOptions(request *DeleteWhiteRuleListRequest, runtime *util.RuntimeOptions) (_result *DeleteWhiteRuleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWhiteRuleList"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWhiteRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWhiteRuleList(request *DeleteWhiteRuleListRequest) (_result *DeleteWhiteRuleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWhiteRuleListResponse{}
	_body, _err := client.DeleteWhiteRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAggregateFunctionWithOptions(request *DescribeAggregateFunctionRequest, runtime *util.RuntimeOptions) (_result *DescribeAggregateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAggregateFunction"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAggregateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAggregateFunction(request *DescribeAggregateFunctionRequest) (_result *DescribeAggregateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAggregateFunctionResponse{}
	_body, _err := client.DescribeAggregateFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAlertSceneByEventWithOptions(request *DescribeAlertSceneByEventRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertSceneByEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlertSceneByEvent"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlertSceneByEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAlertSceneByEvent(request *DescribeAlertSceneByEventRequest) (_result *DescribeAlertSceneByEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertSceneByEventResponse{}
	_body, _err := client.DescribeAlertSceneByEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAlertSourceWithOptions(request *DescribeAlertSourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlertSource"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlertSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAlertSource(request *DescribeAlertSourceRequest) (_result *DescribeAlertSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertSourceResponse{}
	_body, _err := client.DescribeAlertSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAlertSourceWithEventWithOptions(request *DescribeAlertSourceWithEventRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertSourceWithEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlertSourceWithEvent"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlertSourceWithEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAlertSourceWithEvent(request *DescribeAlertSourceWithEventRequest) (_result *DescribeAlertSourceWithEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertSourceWithEventResponse{}
	_body, _err := client.DescribeAlertSourceWithEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAlertTypeWithOptions(request *DescribeAlertTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlertType"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlertTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAlertType(request *DescribeAlertTypeRequest) (_result *DescribeAlertTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertTypeResponse{}
	_body, _err := client.DescribeAlertTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAlertsCountWithOptions(request *DescribeAlertsCountRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertsCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlertsCount"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlertsCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAlertsCount(request *DescribeAlertsCountRequest) (_result *DescribeAlertsCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertsCountResponse{}
	_body, _err := client.DescribeAlertsCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAttackTimeLineWithOptions(request *DescribeAttackTimeLineRequest, runtime *util.RuntimeOptions) (_result *DescribeAttackTimeLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetName)) {
		body["AssetName"] = request.AssetName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAttackTimeLine"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAttackTimeLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAttackTimeLine(request *DescribeAttackTimeLineRequest) (_result *DescribeAttackTimeLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAttackTimeLineResponse{}
	_body, _err := client.DescribeAttackTimeLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutomateResponseConfigCounterWithOptions(request *DescribeAutomateResponseConfigCounterRequest, runtime *util.RuntimeOptions) (_result *DescribeAutomateResponseConfigCounterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAutomateResponseConfigCounter"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAutomateResponseConfigCounterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutomateResponseConfigCounter(request *DescribeAutomateResponseConfigCounterRequest) (_result *DescribeAutomateResponseConfigCounterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutomateResponseConfigCounterResponse{}
	_body, _err := client.DescribeAutomateResponseConfigCounterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutomateResponseConfigFeatureWithOptions(request *DescribeAutomateResponseConfigFeatureRequest, runtime *util.RuntimeOptions) (_result *DescribeAutomateResponseConfigFeatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoResponseType)) {
		body["AutoResponseType"] = request.AutoResponseType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAutomateResponseConfigFeature"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAutomateResponseConfigFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutomateResponseConfigFeature(request *DescribeAutomateResponseConfigFeatureRequest) (_result *DescribeAutomateResponseConfigFeatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutomateResponseConfigFeatureResponse{}
	_body, _err := client.DescribeAutomateResponseConfigFeatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutomateResponseConfigPlayBooksWithOptions(request *DescribeAutomateResponseConfigPlayBooksRequest, runtime *util.RuntimeOptions) (_result *DescribeAutomateResponseConfigPlayBooksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoResponseType)) {
		body["AutoResponseType"] = request.AutoResponseType
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		body["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAutomateResponseConfigPlayBooks"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAutomateResponseConfigPlayBooksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutomateResponseConfigPlayBooks(request *DescribeAutomateResponseConfigPlayBooksRequest) (_result *DescribeAutomateResponseConfigPlayBooksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutomateResponseConfigPlayBooksResponse{}
	_body, _err := client.DescribeAutomateResponseConfigPlayBooksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCloudSiemAssetsWithOptions(request *DescribeCloudSiemAssetsRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudSiemAssetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetType)) {
		body["AssetType"] = request.AssetType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudSiemAssets"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudSiemAssetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCloudSiemAssets(request *DescribeCloudSiemAssetsRequest) (_result *DescribeCloudSiemAssetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudSiemAssetsResponse{}
	_body, _err := client.DescribeCloudSiemAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCloudSiemAssetsCounterWithOptions(request *DescribeCloudSiemAssetsCounterRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudSiemAssetsCounterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudSiemAssetsCounter"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudSiemAssetsCounterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCloudSiemAssetsCounter(request *DescribeCloudSiemAssetsCounterRequest) (_result *DescribeCloudSiemAssetsCounterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudSiemAssetsCounterResponse{}
	_body, _err := client.DescribeCloudSiemAssetsCounterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCloudSiemEventDetailWithOptions(request *DescribeCloudSiemEventDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudSiemEventDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudSiemEventDetail"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudSiemEventDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCloudSiemEventDetail(request *DescribeCloudSiemEventDetailRequest) (_result *DescribeCloudSiemEventDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudSiemEventDetailResponse{}
	_body, _err := client.DescribeCloudSiemEventDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCloudSiemEventsWithOptions(request *DescribeCloudSiemEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudSiemEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetId)) {
		body["AssetId"] = request.AssetId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EventName)) {
		body["EventName"] = request.EventName
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderField)) {
		body["OrderField"] = request.OrderField
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThreadLevel)) {
		body["ThreadLevel"] = request.ThreadLevel
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudSiemEvents"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudSiemEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCloudSiemEvents(request *DescribeCloudSiemEventsRequest) (_result *DescribeCloudSiemEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudSiemEventsResponse{}
	_body, _err := client.DescribeCloudSiemEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomizeRuleWithOptions(request *DescribeCustomizeRuleRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomizeRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCustomizeRule"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCustomizeRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomizeRule(request *DescribeCustomizeRuleRequest) (_result *DescribeCustomizeRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomizeRuleResponse{}
	_body, _err := client.DescribeCustomizeRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomizeRuleCountWithOptions(request *DescribeCustomizeRuleCountRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomizeRuleCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCustomizeRuleCount"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCustomizeRuleCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomizeRuleCount(request *DescribeCustomizeRuleCountRequest) (_result *DescribeCustomizeRuleCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomizeRuleCountResponse{}
	_body, _err := client.DescribeCustomizeRuleCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomizeRuleTestWithOptions(request *DescribeCustomizeRuleTestRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomizeRuleTestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCustomizeRuleTest"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCustomizeRuleTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomizeRuleTest(request *DescribeCustomizeRuleTestRequest) (_result *DescribeCustomizeRuleTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomizeRuleTestResponse{}
	_body, _err := client.DescribeCustomizeRuleTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomizeRuleTestHistogramWithOptions(request *DescribeCustomizeRuleTestHistogramRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomizeRuleTestHistogramResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCustomizeRuleTestHistogram"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCustomizeRuleTestHistogramResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomizeRuleTestHistogram(request *DescribeCustomizeRuleTestHistogramRequest) (_result *DescribeCustomizeRuleTestHistogramResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomizeRuleTestHistogramResponse{}
	_body, _err := client.DescribeCustomizeRuleTestHistogramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDisposeAndPlaybookWithOptions(request *DescribeDisposeAndPlaybookRequest, runtime *util.RuntimeOptions) (_result *DescribeDisposeAndPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		body["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDisposeAndPlaybook"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDisposeAndPlaybookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDisposeAndPlaybook(request *DescribeDisposeAndPlaybookRequest) (_result *DescribeDisposeAndPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDisposeAndPlaybookResponse{}
	_body, _err := client.DescribeDisposeAndPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDisposeStrategyPlaybookWithOptions(request *DescribeDisposeStrategyPlaybookRequest, runtime *util.RuntimeOptions) (_result *DescribeDisposeStrategyPlaybookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDisposeStrategyPlaybook"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDisposeStrategyPlaybookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDisposeStrategyPlaybook(request *DescribeDisposeStrategyPlaybookRequest) (_result *DescribeDisposeStrategyPlaybookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDisposeStrategyPlaybookResponse{}
	_body, _err := client.DescribeDisposeStrategyPlaybookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEntityInfoWithOptions(request *DescribeEntityInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeEntityInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityIdentity)) {
		body["EntityIdentity"] = request.EntityIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SophonTaskId)) {
		body["SophonTaskId"] = request.SophonTaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEntityInfo"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEntityInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEntityInfo(request *DescribeEntityInfoRequest) (_result *DescribeEntityInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEntityInfoResponse{}
	_body, _err := client.DescribeEntityInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEventCountByThreatLevelWithOptions(request *DescribeEventCountByThreatLevelRequest, runtime *util.RuntimeOptions) (_result *DescribeEventCountByThreatLevelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEventCountByThreatLevel"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEventCountByThreatLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEventCountByThreatLevel(request *DescribeEventCountByThreatLevelRequest) (_result *DescribeEventCountByThreatLevelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventCountByThreatLevelResponse{}
	_body, _err := client.DescribeEventCountByThreatLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEventDisposeWithOptions(request *DescribeEventDisposeRequest, runtime *util.RuntimeOptions) (_result *DescribeEventDisposeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEventDispose"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEventDisposeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEventDispose(request *DescribeEventDisposeRequest) (_result *DescribeEventDisposeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventDisposeResponse{}
	_body, _err := client.DescribeEventDisposeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeJobStatusWithOptions(request *DescribeJobStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SubmitId)) {
		body["SubmitId"] = request.SubmitId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeJobStatus"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeJobStatus(request *DescribeJobStatusRequest) (_result *DescribeJobStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeJobStatusResponse{}
	_body, _err := client.DescribeJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogFieldsWithOptions(request *DescribeLogFieldsRequest, runtime *util.RuntimeOptions) (_result *DescribeLogFieldsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogSource)) {
		body["LogSource"] = request.LogSource
	}

	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		body["LogType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogFields"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogFields(request *DescribeLogFieldsRequest) (_result *DescribeLogFieldsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogFieldsResponse{}
	_body, _err := client.DescribeLogFieldsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogSourceWithOptions(request *DescribeLogSourceRequest, runtime *util.RuntimeOptions) (_result *DescribeLogSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		body["LogType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogSource"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogSource(request *DescribeLogSourceRequest) (_result *DescribeLogSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogSourceResponse{}
	_body, _err := client.DescribeLogSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogStoreWithOptions(request *DescribeLogStoreRequest, runtime *util.RuntimeOptions) (_result *DescribeLogStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogStore"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogStore(request *DescribeLogStoreRequest) (_result *DescribeLogStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogStoreResponse{}
	_body, _err := client.DescribeLogStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogTypeWithOptions(request *DescribeLogTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeLogTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogType"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogType(request *DescribeLogTypeRequest) (_result *DescribeLogTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogTypeResponse{}
	_body, _err := client.DescribeLogTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOperatorsWithOptions(request *DescribeOperatorsRequest, runtime *util.RuntimeOptions) (_result *DescribeOperatorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneType)) {
		body["SceneType"] = request.SceneType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOperators"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOperatorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOperators(request *DescribeOperatorsRequest) (_result *DescribeOperatorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOperatorsResponse{}
	_body, _err := client.DescribeOperatorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScopeUsersWithOptions(request *DescribeScopeUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeScopeUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScopeUsers"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScopeUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScopeUsers(request *DescribeScopeUsersRequest) (_result *DescribeScopeUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScopeUsersResponse{}
	_body, _err := client.DescribeScopeUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStorageWithOptions(request *DescribeStorageRequest, runtime *util.RuntimeOptions) (_result *DescribeStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStorage"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStorage(request *DescribeStorageRequest) (_result *DescribeStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStorageResponse{}
	_body, _err := client.DescribeStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWafScopeWithOptions(request *DescribeWafScopeRequest, runtime *util.RuntimeOptions) (_result *DescribeWafScopeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWafScope"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWafScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWafScope(request *DescribeWafScopeRequest) (_result *DescribeWafScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWafScopeResponse{}
	_body, _err := client.DescribeWafScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DoQuickFieldWithOptions(request *DoQuickFieldRequest, runtime *util.RuntimeOptions) (_result *DoQuickFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		body["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.Index)) {
		body["Index"] = request.Index
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		body["Reverse"] = request.Reverse
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		body["To"] = request.To
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DoQuickField"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DoQuickFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoQuickField(request *DoQuickFieldRequest) (_result *DoQuickFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoQuickFieldResponse{}
	_body, _err := client.DoQuickFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DoSelfDelegateWithOptions(request *DoSelfDelegateRequest, runtime *util.RuntimeOptions) (_result *DoSelfDelegateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		body["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.DelegateOrNot)) {
		body["DelegateOrNot"] = request.DelegateOrNot
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DoSelfDelegate"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DoSelfDelegateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoSelfDelegate(request *DoSelfDelegateRequest) (_result *DoSelfDelegateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoSelfDelegateResponse{}
	_body, _err := client.DoSelfDelegateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCapacityWithOptions(request *GetCapacityRequest, runtime *util.RuntimeOptions) (_result *GetCapacityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCapacity"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCapacity(request *GetCapacityRequest) (_result *GetCapacityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCapacityResponse{}
	_body, _err := client.GetCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHistogramsWithOptions(request *GetHistogramsRequest, runtime *util.RuntimeOptions) (_result *GetHistogramsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		body["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		body["To"] = request.To
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHistograms"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHistogramsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHistograms(request *GetHistogramsRequest) (_result *GetHistogramsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHistogramsResponse{}
	_body, _err := client.GetHistogramsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLogsWithOptions(request *GetLogsRequest, runtime *util.RuntimeOptions) (_result *GetLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		body["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReverseOrNot)) {
		body["ReverseOrNot"] = request.ReverseOrNot
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		body["To"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.Total)) {
		body["Total"] = request.Total
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLogs"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogs(request *GetLogsRequest) (_result *GetLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLogsResponse{}
	_body, _err := client.GetLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQuickQueryWithOptions(request *GetQuickQueryRequest, runtime *util.RuntimeOptions) (_result *GetQuickQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchName)) {
		body["SearchName"] = request.SearchName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQuickQuery"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQuickQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQuickQuery(request *GetQuickQueryRequest) (_result *GetQuickQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQuickQueryResponse{}
	_body, _err := client.GetQuickQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStorageWithOptions(request *GetStorageRequest, runtime *util.RuntimeOptions) (_result *GetStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStorage"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStorage(request *GetStorageRequest) (_result *GetStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStorageResponse{}
	_body, _err := client.GetStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAutomateResponseConfigsWithOptions(request *ListAutomateResponseConfigsRequest, runtime *util.RuntimeOptions) (_result *ListAutomateResponseConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		body["ActionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.AutoResponseType)) {
		body["AutoResponseType"] = request.AutoResponseType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		body["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		body["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAutomateResponseConfigs"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAutomateResponseConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAutomateResponseConfigs(request *ListAutomateResponseConfigsRequest) (_result *ListAutomateResponseConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAutomateResponseConfigsResponse{}
	_body, _err := client.ListAutomateResponseConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCloudSiemCustomizeRulesWithOptions(request *ListCloudSiemCustomizeRulesRequest, runtime *util.RuntimeOptions) (_result *ListCloudSiemCustomizeRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertType)) {
		body["AlertType"] = request.AlertType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		body["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		body["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThreatLevel)) {
		body["ThreatLevel"] = request.ThreatLevel
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCloudSiemCustomizeRules"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCloudSiemCustomizeRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCloudSiemCustomizeRules(request *ListCloudSiemCustomizeRulesRequest) (_result *ListCloudSiemCustomizeRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCloudSiemCustomizeRulesResponse{}
	_body, _err := client.ListCloudSiemCustomizeRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCloudSiemPredefinedRulesWithOptions(request *ListCloudSiemPredefinedRulesRequest, runtime *util.RuntimeOptions) (_result *ListCloudSiemPredefinedRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertType)) {
		body["AlertType"] = request.AlertType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		body["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		body["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThreatLevel)) {
		body["ThreatLevel"] = request.ThreatLevel
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCloudSiemPredefinedRules"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCloudSiemPredefinedRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCloudSiemPredefinedRules(request *ListCloudSiemPredefinedRulesRequest) (_result *ListCloudSiemPredefinedRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCloudSiemPredefinedRulesResponse{}
	_body, _err := client.ListCloudSiemPredefinedRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCustomizeRuleTestResultWithOptions(request *ListCustomizeRuleTestResultRequest, runtime *util.RuntimeOptions) (_result *ListCustomizeRuleTestResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomizeRuleTestResult"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomizeRuleTestResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCustomizeRuleTestResult(request *ListCustomizeRuleTestResultRequest) (_result *ListCustomizeRuleTestResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCustomizeRuleTestResultResponse{}
	_body, _err := client.ListCustomizeRuleTestResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeliveryWithOptions(request *ListDeliveryRequest, runtime *util.RuntimeOptions) (_result *ListDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDelivery"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDelivery(request *ListDeliveryRequest) (_result *ListDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeliveryResponse{}
	_body, _err := client.ListDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDisposeStrategyWithOptions(request *ListDisposeStrategyRequest, runtime *util.RuntimeOptions) (_result *ListDisposeStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EffectiveStatus)) {
		body["EffectiveStatus"] = request.EffectiveStatus
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EntityIdentity)) {
		body["EntityIdentity"] = request.EntityIdentity
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		body["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderField)) {
		body["OrderField"] = request.OrderField
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookName)) {
		body["PlaybookName"] = request.PlaybookName
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookTypes)) {
		body["PlaybookTypes"] = request.PlaybookTypes
	}

	if !tea.BoolValue(util.IsUnset(request.PlaybookUuid)) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SophonTaskId)) {
		body["SophonTaskId"] = request.SophonTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDisposeStrategy"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDisposeStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDisposeStrategy(request *ListDisposeStrategyRequest) (_result *ListDisposeStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDisposeStrategyResponse{}
	_body, _err := client.ListDisposeStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOperationWithOptions(request *ListOperationRequest, runtime *util.RuntimeOptions) (_result *ListOperationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOperation"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOperation(request *ListOperationRequest) (_result *ListOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOperationResponse{}
	_body, _err := client.ListOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQuickQueryWithOptions(request *ListQuickQueryRequest, runtime *util.RuntimeOptions) (_result *ListQuickQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		body["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQuickQuery"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQuickQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQuickQuery(request *ListQuickQueryRequest) (_result *ListQuickQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQuickQueryResponse{}
	_body, _err := client.ListQuickQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenDeliveryWithOptions(request *OpenDeliveryRequest, runtime *util.RuntimeOptions) (_result *OpenDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogCode)) {
		body["LogCode"] = request.LogCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenDelivery"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenDelivery(request *OpenDeliveryRequest) (_result *OpenDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenDeliveryResponse{}
	_body, _err := client.OpenDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostAutomateResponseConfigWithOptions(request *PostAutomateResponseConfigRequest, runtime *util.RuntimeOptions) (_result *PostAutomateResponseConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionConfig)) {
		body["ActionConfig"] = request.ActionConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		body["ActionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.AutoResponseType)) {
		body["AutoResponseType"] = request.AutoResponseType
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionCondition)) {
		body["ExecutionCondition"] = request.ExecutionCondition
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		body["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		body["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostAutomateResponseConfig"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostAutomateResponseConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostAutomateResponseConfig(request *PostAutomateResponseConfigRequest) (_result *PostAutomateResponseConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostAutomateResponseConfigResponse{}
	_body, _err := client.PostAutomateResponseConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostCustomizeRuleWithOptions(request *PostCustomizeRuleRequest, runtime *util.RuntimeOptions) (_result *PostCustomizeRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertType)) {
		body["AlertType"] = request.AlertType
	}

	if !tea.BoolValue(util.IsUnset(request.AlertTypeMds)) {
		body["AlertTypeMds"] = request.AlertTypeMds
	}

	if !tea.BoolValue(util.IsUnset(request.EventTransferExt)) {
		body["EventTransferExt"] = request.EventTransferExt
	}

	if !tea.BoolValue(util.IsUnset(request.EventTransferSwitch)) {
		body["EventTransferSwitch"] = request.EventTransferSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.EventTransferType)) {
		body["EventTransferType"] = request.EventTransferType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.LogSource)) {
		body["LogSource"] = request.LogSource
	}

	if !tea.BoolValue(util.IsUnset(request.LogSourceMds)) {
		body["LogSourceMds"] = request.LogSourceMds
	}

	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		body["LogType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.LogTypeMds)) {
		body["LogTypeMds"] = request.LogTypeMds
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCycle)) {
		body["QueryCycle"] = request.QueryCycle
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleCondition)) {
		body["RuleCondition"] = request.RuleCondition
	}

	if !tea.BoolValue(util.IsUnset(request.RuleDesc)) {
		body["RuleDesc"] = request.RuleDesc
	}

	if !tea.BoolValue(util.IsUnset(request.RuleGroup)) {
		body["RuleGroup"] = request.RuleGroup
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		body["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleThreshold)) {
		body["RuleThreshold"] = request.RuleThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.ThreatLevel)) {
		body["ThreatLevel"] = request.ThreatLevel
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostCustomizeRule"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostCustomizeRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostCustomizeRule(request *PostCustomizeRuleRequest) (_result *PostCustomizeRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostCustomizeRuleResponse{}
	_body, _err := client.PostCustomizeRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostCustomizeRuleTestWithOptions(request *PostCustomizeRuleTestRequest, runtime *util.RuntimeOptions) (_result *PostCustomizeRuleTestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SimulatedData)) {
		body["SimulatedData"] = request.SimulatedData
	}

	if !tea.BoolValue(util.IsUnset(request.TestType)) {
		body["TestType"] = request.TestType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostCustomizeRuleTest"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostCustomizeRuleTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostCustomizeRuleTest(request *PostCustomizeRuleTestRequest) (_result *PostCustomizeRuleTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostCustomizeRuleTestResponse{}
	_body, _err := client.PostCustomizeRuleTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostEventDisposeAndWhiteruleListWithOptions(request *PostEventDisposeAndWhiteruleListRequest, runtime *util.RuntimeOptions) (_result *PostEventDisposeAndWhiteruleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventDispose)) {
		body["EventDispose"] = request.EventDispose
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverInfo)) {
		body["ReceiverInfo"] = request.ReceiverInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostEventDisposeAndWhiteruleList"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostEventDisposeAndWhiteruleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostEventDisposeAndWhiteruleList(request *PostEventDisposeAndWhiteruleListRequest) (_result *PostEventDisposeAndWhiteruleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostEventDisposeAndWhiteruleListResponse{}
	_body, _err := client.PostEventDisposeAndWhiteruleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostEventWhiteruleListWithOptions(request *PostEventWhiteruleListRequest, runtime *util.RuntimeOptions) (_result *PostEventWhiteruleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteruleList)) {
		body["WhiteruleList"] = request.WhiteruleList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostEventWhiteruleList"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostEventWhiteruleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostEventWhiteruleList(request *PostEventWhiteruleListRequest) (_result *PostEventWhiteruleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostEventWhiteruleListResponse{}
	_body, _err := client.PostEventWhiteruleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostFinishCustomizeRuleTestWithOptions(request *PostFinishCustomizeRuleTestRequest, runtime *util.RuntimeOptions) (_result *PostFinishCustomizeRuleTestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostFinishCustomizeRuleTest"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostFinishCustomizeRuleTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostFinishCustomizeRuleTest(request *PostFinishCustomizeRuleTestRequest) (_result *PostFinishCustomizeRuleTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostFinishCustomizeRuleTestResponse{}
	_body, _err := client.PostFinishCustomizeRuleTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostRuleStatusChangeWithOptions(request *PostRuleStatusChangeRequest, runtime *util.RuntimeOptions) (_result *PostRuleStatusChangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		body["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.InUse)) {
		body["InUse"] = request.InUse
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		body["RuleType"] = request.RuleType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PostRuleStatusChange"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PostRuleStatusChangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostRuleStatusChange(request *PostRuleStatusChangeRequest) (_result *PostRuleStatusChangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PostRuleStatusChangeResponse{}
	_body, _err := client.PostRuleStatusChangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestoreCapacityWithOptions(request *RestoreCapacityRequest, runtime *util.RuntimeOptions) (_result *RestoreCapacityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RestoreCapacity"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestoreCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestoreCapacity(request *RestoreCapacityRequest) (_result *RestoreCapacityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestoreCapacityResponse{}
	_body, _err := client.RestoreCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveQuickQueryWithOptions(request *SaveQuickQueryRequest, runtime *util.RuntimeOptions) (_result *SaveQuickQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveQuickQuery"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveQuickQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveQuickQuery(request *SaveQuickQueryRequest) (_result *SaveQuickQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveQuickQueryResponse{}
	_body, _err := client.SaveQuickQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetStorageWithOptions(request *SetStorageRequest, runtime *util.RuntimeOptions) (_result *SetStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		body["Ttl"] = request.Ttl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetStorage"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetStorage(request *SetStorageRequest) (_result *SetStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetStorageResponse{}
	_body, _err := client.SetStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ShowQuickAnalysisWithOptions(request *ShowQuickAnalysisRequest, runtime *util.RuntimeOptions) (_result *ShowQuickAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ShowQuickAnalysis"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ShowQuickAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ShowQuickAnalysis(request *ShowQuickAnalysisRequest) (_result *ShowQuickAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ShowQuickAnalysisResponse{}
	_body, _err := client.ShowQuickAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAutomateResponseConfigStatusWithOptions(request *UpdateAutomateResponseConfigStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateAutomateResponseConfigStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		body["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.InUse)) {
		body["InUse"] = request.InUse
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAutomateResponseConfigStatus"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAutomateResponseConfigStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAutomateResponseConfigStatus(request *UpdateAutomateResponseConfigStatusRequest) (_result *UpdateAutomateResponseConfigStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAutomateResponseConfigStatusResponse{}
	_body, _err := client.UpdateAutomateResponseConfigStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWhiteRuleListWithOptions(request *UpdateWhiteRuleListRequest, runtime *util.RuntimeOptions) (_result *UpdateWhiteRuleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		body["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.IncidentUuid)) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteRuleId)) {
		body["WhiteRuleId"] = request.WhiteRuleId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWhiteRuleList"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWhiteRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWhiteRuleList(request *UpdateWhiteRuleListRequest) (_result *UpdateWhiteRuleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWhiteRuleListResponse{}
	_body, _err := client.UpdateWhiteRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
