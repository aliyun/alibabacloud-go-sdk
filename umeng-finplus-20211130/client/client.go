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

type CancelComputeTaskByBcIdRequest struct {
	BcId *int64 `json:"bcId,omitempty" xml:"bcId,omitempty"`
}

func (s CancelComputeTaskByBcIdRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelComputeTaskByBcIdRequest) GoString() string {
	return s.String()
}

func (s *CancelComputeTaskByBcIdRequest) SetBcId(v int64) *CancelComputeTaskByBcIdRequest {
	s.BcId = &v
	return s
}

type CancelComputeTaskByBcIdResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelComputeTaskByBcIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelComputeTaskByBcIdResponseBody) GoString() string {
	return s.String()
}

func (s *CancelComputeTaskByBcIdResponseBody) SetCode(v string) *CancelComputeTaskByBcIdResponseBody {
	s.Code = &v
	return s
}

func (s *CancelComputeTaskByBcIdResponseBody) SetData(v bool) *CancelComputeTaskByBcIdResponseBody {
	s.Data = &v
	return s
}

func (s *CancelComputeTaskByBcIdResponseBody) SetMsg(v string) *CancelComputeTaskByBcIdResponseBody {
	s.Msg = &v
	return s
}

func (s *CancelComputeTaskByBcIdResponseBody) SetRequestId(v string) *CancelComputeTaskByBcIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelComputeTaskByBcIdResponseBody) SetSuccess(v bool) *CancelComputeTaskByBcIdResponseBody {
	s.Success = &v
	return s
}

type CancelComputeTaskByBcIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelComputeTaskByBcIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelComputeTaskByBcIdResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelComputeTaskByBcIdResponse) GoString() string {
	return s.String()
}

func (s *CancelComputeTaskByBcIdResponse) SetHeaders(v map[string]*string) *CancelComputeTaskByBcIdResponse {
	s.Headers = v
	return s
}

func (s *CancelComputeTaskByBcIdResponse) SetStatusCode(v int32) *CancelComputeTaskByBcIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelComputeTaskByBcIdResponse) SetBody(v *CancelComputeTaskByBcIdResponseBody) *CancelComputeTaskByBcIdResponse {
	s.Body = v
	return s
}

type CreateComputeTaskByDataSetIdRequest struct {
	BatchInfoForm []*CreateComputeTaskByDataSetIdRequestBatchInfoForm `json:"BatchInfoForm,omitempty" xml:"BatchInfoForm,omitempty" type:"Repeated"`
	DatasetId     *int64                                              `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	Name          *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Remarks       *string                                             `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
}

func (s CreateComputeTaskByDataSetIdRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeTaskByDataSetIdRequest) GoString() string {
	return s.String()
}

func (s *CreateComputeTaskByDataSetIdRequest) SetBatchInfoForm(v []*CreateComputeTaskByDataSetIdRequestBatchInfoForm) *CreateComputeTaskByDataSetIdRequest {
	s.BatchInfoForm = v
	return s
}

func (s *CreateComputeTaskByDataSetIdRequest) SetDatasetId(v int64) *CreateComputeTaskByDataSetIdRequest {
	s.DatasetId = &v
	return s
}

func (s *CreateComputeTaskByDataSetIdRequest) SetName(v string) *CreateComputeTaskByDataSetIdRequest {
	s.Name = &v
	return s
}

func (s *CreateComputeTaskByDataSetIdRequest) SetRemarks(v string) *CreateComputeTaskByDataSetIdRequest {
	s.Remarks = &v
	return s
}

type CreateComputeTaskByDataSetIdRequestBatchInfoForm struct {
	AppId      *int64                                                        `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CuVersions []*CreateComputeTaskByDataSetIdRequestBatchInfoFormCuVersions `json:"CuVersions,omitempty" xml:"CuVersions,omitempty" type:"Repeated"`
}

func (s CreateComputeTaskByDataSetIdRequestBatchInfoForm) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeTaskByDataSetIdRequestBatchInfoForm) GoString() string {
	return s.String()
}

func (s *CreateComputeTaskByDataSetIdRequestBatchInfoForm) SetAppId(v int64) *CreateComputeTaskByDataSetIdRequestBatchInfoForm {
	s.AppId = &v
	return s
}

func (s *CreateComputeTaskByDataSetIdRequestBatchInfoForm) SetCuVersions(v []*CreateComputeTaskByDataSetIdRequestBatchInfoFormCuVersions) *CreateComputeTaskByDataSetIdRequestBatchInfoForm {
	s.CuVersions = v
	return s
}

type CreateComputeTaskByDataSetIdRequestBatchInfoFormCuVersions struct {
	CuId      *string `json:"CuId,omitempty" xml:"CuId,omitempty"`
	CuVersion *string `json:"CuVersion,omitempty" xml:"CuVersion,omitempty"`
}

func (s CreateComputeTaskByDataSetIdRequestBatchInfoFormCuVersions) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeTaskByDataSetIdRequestBatchInfoFormCuVersions) GoString() string {
	return s.String()
}

func (s *CreateComputeTaskByDataSetIdRequestBatchInfoFormCuVersions) SetCuId(v string) *CreateComputeTaskByDataSetIdRequestBatchInfoFormCuVersions {
	s.CuId = &v
	return s
}

func (s *CreateComputeTaskByDataSetIdRequestBatchInfoFormCuVersions) SetCuVersion(v string) *CreateComputeTaskByDataSetIdRequestBatchInfoFormCuVersions {
	s.CuVersion = &v
	return s
}

type CreateComputeTaskByDataSetIdResponseBody struct {
	Code      *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Msg       *string  `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateComputeTaskByDataSetIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeTaskByDataSetIdResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComputeTaskByDataSetIdResponseBody) SetCode(v string) *CreateComputeTaskByDataSetIdResponseBody {
	s.Code = &v
	return s
}

func (s *CreateComputeTaskByDataSetIdResponseBody) SetData(v []*int64) *CreateComputeTaskByDataSetIdResponseBody {
	s.Data = v
	return s
}

func (s *CreateComputeTaskByDataSetIdResponseBody) SetMsg(v string) *CreateComputeTaskByDataSetIdResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateComputeTaskByDataSetIdResponseBody) SetRequestId(v string) *CreateComputeTaskByDataSetIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateComputeTaskByDataSetIdResponseBody) SetSuccess(v bool) *CreateComputeTaskByDataSetIdResponseBody {
	s.Success = &v
	return s
}

type CreateComputeTaskByDataSetIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComputeTaskByDataSetIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComputeTaskByDataSetIdResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeTaskByDataSetIdResponse) GoString() string {
	return s.String()
}

func (s *CreateComputeTaskByDataSetIdResponse) SetHeaders(v map[string]*string) *CreateComputeTaskByDataSetIdResponse {
	s.Headers = v
	return s
}

func (s *CreateComputeTaskByDataSetIdResponse) SetStatusCode(v int32) *CreateComputeTaskByDataSetIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComputeTaskByDataSetIdResponse) SetBody(v *CreateComputeTaskByDataSetIdResponseBody) *CreateComputeTaskByDataSetIdResponse {
	s.Body = v
	return s
}

type CreateComputeTaskByMultiDataSetIdRequest struct {
	AppId      *int64  `json:"appId,omitempty" xml:"appId,omitempty"`
	DataSetIds *string `json:"dataSetIds,omitempty" xml:"dataSetIds,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Remarks    *string `json:"remarks,omitempty" xml:"remarks,omitempty"`
}

func (s CreateComputeTaskByMultiDataSetIdRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeTaskByMultiDataSetIdRequest) GoString() string {
	return s.String()
}

func (s *CreateComputeTaskByMultiDataSetIdRequest) SetAppId(v int64) *CreateComputeTaskByMultiDataSetIdRequest {
	s.AppId = &v
	return s
}

func (s *CreateComputeTaskByMultiDataSetIdRequest) SetDataSetIds(v string) *CreateComputeTaskByMultiDataSetIdRequest {
	s.DataSetIds = &v
	return s
}

func (s *CreateComputeTaskByMultiDataSetIdRequest) SetName(v string) *CreateComputeTaskByMultiDataSetIdRequest {
	s.Name = &v
	return s
}

func (s *CreateComputeTaskByMultiDataSetIdRequest) SetRemarks(v string) *CreateComputeTaskByMultiDataSetIdRequest {
	s.Remarks = &v
	return s
}

type CreateComputeTaskByMultiDataSetIdResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateComputeTaskByMultiDataSetIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeTaskByMultiDataSetIdResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComputeTaskByMultiDataSetIdResponseBody) SetCode(v string) *CreateComputeTaskByMultiDataSetIdResponseBody {
	s.Code = &v
	return s
}

func (s *CreateComputeTaskByMultiDataSetIdResponseBody) SetData(v int64) *CreateComputeTaskByMultiDataSetIdResponseBody {
	s.Data = &v
	return s
}

func (s *CreateComputeTaskByMultiDataSetIdResponseBody) SetMsg(v string) *CreateComputeTaskByMultiDataSetIdResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateComputeTaskByMultiDataSetIdResponseBody) SetRequestId(v string) *CreateComputeTaskByMultiDataSetIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateComputeTaskByMultiDataSetIdResponseBody) SetSuccess(v bool) *CreateComputeTaskByMultiDataSetIdResponseBody {
	s.Success = &v
	return s
}

type CreateComputeTaskByMultiDataSetIdResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComputeTaskByMultiDataSetIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComputeTaskByMultiDataSetIdResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeTaskByMultiDataSetIdResponse) GoString() string {
	return s.String()
}

func (s *CreateComputeTaskByMultiDataSetIdResponse) SetHeaders(v map[string]*string) *CreateComputeTaskByMultiDataSetIdResponse {
	s.Headers = v
	return s
}

func (s *CreateComputeTaskByMultiDataSetIdResponse) SetStatusCode(v int32) *CreateComputeTaskByMultiDataSetIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComputeTaskByMultiDataSetIdResponse) SetBody(v *CreateComputeTaskByMultiDataSetIdResponseBody) *CreateComputeTaskByMultiDataSetIdResponse {
	s.Body = v
	return s
}

type GetAvailableDataSetListResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetAvailableDataSetListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Msg       *string                                    `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAvailableDataSetListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAvailableDataSetListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAvailableDataSetListResponseBody) SetCode(v string) *GetAvailableDataSetListResponseBody {
	s.Code = &v
	return s
}

func (s *GetAvailableDataSetListResponseBody) SetData(v []*GetAvailableDataSetListResponseBodyData) *GetAvailableDataSetListResponseBody {
	s.Data = v
	return s
}

func (s *GetAvailableDataSetListResponseBody) SetMsg(v string) *GetAvailableDataSetListResponseBody {
	s.Msg = &v
	return s
}

func (s *GetAvailableDataSetListResponseBody) SetRequestId(v string) *GetAvailableDataSetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAvailableDataSetListResponseBody) SetSuccess(v string) *GetAvailableDataSetListResponseBody {
	s.Success = &v
	return s
}

type GetAvailableDataSetListResponseBodyData struct {
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DataSetType *int32  `json:"dataSetType,omitempty" xml:"dataSetType,omitempty"`
	DatasetId   *int64  `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	IdTypeDesc  *string `json:"idTypeDesc,omitempty" xml:"idTypeDesc,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	StatusDesc  *string `json:"statusDesc,omitempty" xml:"statusDesc,omitempty"`
}

func (s GetAvailableDataSetListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAvailableDataSetListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAvailableDataSetListResponseBodyData) SetCreateTime(v string) *GetAvailableDataSetListResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetAvailableDataSetListResponseBodyData) SetDataSetType(v int32) *GetAvailableDataSetListResponseBodyData {
	s.DataSetType = &v
	return s
}

func (s *GetAvailableDataSetListResponseBodyData) SetDatasetId(v int64) *GetAvailableDataSetListResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *GetAvailableDataSetListResponseBodyData) SetIdTypeDesc(v string) *GetAvailableDataSetListResponseBodyData {
	s.IdTypeDesc = &v
	return s
}

func (s *GetAvailableDataSetListResponseBodyData) SetName(v string) *GetAvailableDataSetListResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAvailableDataSetListResponseBodyData) SetStatusDesc(v string) *GetAvailableDataSetListResponseBodyData {
	s.StatusDesc = &v
	return s
}

type GetAvailableDataSetListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAvailableDataSetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAvailableDataSetListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAvailableDataSetListResponse) GoString() string {
	return s.String()
}

func (s *GetAvailableDataSetListResponse) SetHeaders(v map[string]*string) *GetAvailableDataSetListResponse {
	s.Headers = v
	return s
}

func (s *GetAvailableDataSetListResponse) SetStatusCode(v int32) *GetAvailableDataSetListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAvailableDataSetListResponse) SetBody(v *GetAvailableDataSetListResponseBody) *GetAvailableDataSetListResponse {
	s.Body = v
	return s
}

type GetComputeResultRequest struct {
	BcId *int64 `json:"bcId,omitempty" xml:"bcId,omitempty"`
}

func (s GetComputeResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetComputeResultRequest) GoString() string {
	return s.String()
}

func (s *GetComputeResultRequest) SetBcId(v int64) *GetComputeResultRequest {
	s.BcId = &v
	return s
}

type GetComputeResultResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetComputeResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                           `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetComputeResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetComputeResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetComputeResultResponseBody) SetCode(v string) *GetComputeResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetComputeResultResponseBody) SetData(v *GetComputeResultResponseBodyData) *GetComputeResultResponseBody {
	s.Data = v
	return s
}

func (s *GetComputeResultResponseBody) SetMsg(v string) *GetComputeResultResponseBody {
	s.Msg = &v
	return s
}

func (s *GetComputeResultResponseBody) SetRequestId(v string) *GetComputeResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetComputeResultResponseBody) SetSuccess(v bool) *GetComputeResultResponseBody {
	s.Success = &v
	return s
}

type GetComputeResultResponseBodyData struct {
	BcId           *int64  `json:"bcId,omitempty" xml:"bcId,omitempty"`
	BilledNum      *int64  `json:"billedNum,omitempty" xml:"billedNum,omitempty"`
	Code10000Num   *int64  `json:"code10000Num,omitempty" xml:"code10000Num,omitempty"`
	Code108Num     *int64  `json:"code108Num,omitempty" xml:"code108Num,omitempty"`
	Code109Num     *int64  `json:"code109Num,omitempty" xml:"code109Num,omitempty"`
	ExportFileName *string `json:"exportFileName,omitempty" xml:"exportFileName,omitempty"`
	FileLineNumber *int64  `json:"fileLineNumber,omitempty" xml:"fileLineNumber,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Password       *string `json:"password,omitempty" xml:"password,omitempty"`
	State          *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s GetComputeResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetComputeResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetComputeResultResponseBodyData) SetBcId(v int64) *GetComputeResultResponseBodyData {
	s.BcId = &v
	return s
}

func (s *GetComputeResultResponseBodyData) SetBilledNum(v int64) *GetComputeResultResponseBodyData {
	s.BilledNum = &v
	return s
}

func (s *GetComputeResultResponseBodyData) SetCode10000Num(v int64) *GetComputeResultResponseBodyData {
	s.Code10000Num = &v
	return s
}

func (s *GetComputeResultResponseBodyData) SetCode108Num(v int64) *GetComputeResultResponseBodyData {
	s.Code108Num = &v
	return s
}

func (s *GetComputeResultResponseBodyData) SetCode109Num(v int64) *GetComputeResultResponseBodyData {
	s.Code109Num = &v
	return s
}

func (s *GetComputeResultResponseBodyData) SetExportFileName(v string) *GetComputeResultResponseBodyData {
	s.ExportFileName = &v
	return s
}

func (s *GetComputeResultResponseBodyData) SetFileLineNumber(v int64) *GetComputeResultResponseBodyData {
	s.FileLineNumber = &v
	return s
}

func (s *GetComputeResultResponseBodyData) SetName(v string) *GetComputeResultResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetComputeResultResponseBodyData) SetPassword(v string) *GetComputeResultResponseBodyData {
	s.Password = &v
	return s
}

func (s *GetComputeResultResponseBodyData) SetState(v string) *GetComputeResultResponseBodyData {
	s.State = &v
	return s
}

type GetComputeResultResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComputeResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComputeResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetComputeResultResponse) GoString() string {
	return s.String()
}

func (s *GetComputeResultResponse) SetHeaders(v map[string]*string) *GetComputeResultResponse {
	s.Headers = v
	return s
}

func (s *GetComputeResultResponse) SetStatusCode(v int32) *GetComputeResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComputeResultResponse) SetBody(v *GetComputeResultResponseBody) *GetComputeResultResponse {
	s.Body = v
	return s
}

type GetDataSetStatusRequest struct {
	DataSetId *int64 `json:"dataSetId,omitempty" xml:"dataSetId,omitempty"`
}

func (s GetDataSetStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataSetStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDataSetStatusRequest) SetDataSetId(v int64) *GetDataSetStatusRequest {
	s.DataSetId = &v
	return s
}

type GetDataSetStatusResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDataSetStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                           `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDataSetStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataSetStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSetStatusResponseBody) SetCode(v string) *GetDataSetStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataSetStatusResponseBody) SetData(v *GetDataSetStatusResponseBodyData) *GetDataSetStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetDataSetStatusResponseBody) SetMsg(v string) *GetDataSetStatusResponseBody {
	s.Msg = &v
	return s
}

func (s *GetDataSetStatusResponseBody) SetSuccess(v bool) *GetDataSetStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataSetStatusResponseBody) SetRequestId(v string) *GetDataSetStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetDataSetStatusResponseBodyData struct {
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DataSetType *int32  `json:"dataSetType,omitempty" xml:"dataSetType,omitempty"`
	DatasetId   *int64  `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	IdTypeDesc  *string `json:"idTypeDesc,omitempty" xml:"idTypeDesc,omitempty"`
	Message     *string `json:"message,omitempty" xml:"message,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	StatusDesc  *string `json:"statusDesc,omitempty" xml:"statusDesc,omitempty"`
}

func (s GetDataSetStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDataSetStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataSetStatusResponseBodyData) SetCreateTime(v string) *GetDataSetStatusResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetDataSetStatusResponseBodyData) SetDataSetType(v int32) *GetDataSetStatusResponseBodyData {
	s.DataSetType = &v
	return s
}

func (s *GetDataSetStatusResponseBodyData) SetDatasetId(v int64) *GetDataSetStatusResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *GetDataSetStatusResponseBodyData) SetIdTypeDesc(v string) *GetDataSetStatusResponseBodyData {
	s.IdTypeDesc = &v
	return s
}

func (s *GetDataSetStatusResponseBodyData) SetMessage(v string) *GetDataSetStatusResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetDataSetStatusResponseBodyData) SetName(v string) *GetDataSetStatusResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetDataSetStatusResponseBodyData) SetStatusDesc(v string) *GetDataSetStatusResponseBodyData {
	s.StatusDesc = &v
	return s
}

type GetDataSetStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataSetStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataSetStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataSetStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDataSetStatusResponse) SetHeaders(v map[string]*string) *GetDataSetStatusResponse {
	s.Headers = v
	return s
}

func (s *GetDataSetStatusResponse) SetStatusCode(v int32) *GetDataSetStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataSetStatusResponse) SetBody(v *GetDataSetStatusResponseBody) *GetDataSetStatusResponse {
	s.Body = v
	return s
}

type GetDataSetStsAKResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetDataSetStsAKResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                          `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataSetStsAKResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataSetStsAKResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSetStsAKResponseBody) SetCode(v string) *GetDataSetStsAKResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataSetStsAKResponseBody) SetData(v *GetDataSetStsAKResponseBodyData) *GetDataSetStsAKResponseBody {
	s.Data = v
	return s
}

func (s *GetDataSetStsAKResponseBody) SetMsg(v string) *GetDataSetStsAKResponseBody {
	s.Msg = &v
	return s
}

func (s *GetDataSetStsAKResponseBody) SetRequestId(v string) *GetDataSetStsAKResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataSetStsAKResponseBody) SetSuccess(v bool) *GetDataSetStsAKResponseBody {
	s.Success = &v
	return s
}

type GetDataSetStsAKResponseBodyData struct {
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Secret   *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	Token    *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetDataSetStsAKResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDataSetStsAKResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataSetStsAKResponseBodyData) SetBucket(v string) *GetDataSetStsAKResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *GetDataSetStsAKResponseBodyData) SetEndpoint(v string) *GetDataSetStsAKResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetDataSetStsAKResponseBodyData) SetId(v string) *GetDataSetStsAKResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetDataSetStsAKResponseBodyData) SetPath(v string) *GetDataSetStsAKResponseBodyData {
	s.Path = &v
	return s
}

func (s *GetDataSetStsAKResponseBodyData) SetSecret(v string) *GetDataSetStsAKResponseBodyData {
	s.Secret = &v
	return s
}

func (s *GetDataSetStsAKResponseBodyData) SetToken(v string) *GetDataSetStsAKResponseBodyData {
	s.Token = &v
	return s
}

type GetDataSetStsAKResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataSetStsAKResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataSetStsAKResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataSetStsAKResponse) GoString() string {
	return s.String()
}

func (s *GetDataSetStsAKResponse) SetHeaders(v map[string]*string) *GetDataSetStsAKResponse {
	s.Headers = v
	return s
}

func (s *GetDataSetStsAKResponse) SetStatusCode(v int32) *GetDataSetStsAKResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataSetStsAKResponse) SetBody(v *GetDataSetStsAKResponseBody) *GetDataSetStsAKResponse {
	s.Body = v
	return s
}

type SubmitDataSetTaskRequest struct {
	DataSetType *int32  `json:"dataSetType,omitempty" xml:"dataSetType,omitempty"`
	IdType      *int32  `json:"idType,omitempty" xml:"idType,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OssFileName *string `json:"ossFileName,omitempty" xml:"ossFileName,omitempty"`
}

func (s SubmitDataSetTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDataSetTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitDataSetTaskRequest) SetDataSetType(v int32) *SubmitDataSetTaskRequest {
	s.DataSetType = &v
	return s
}

func (s *SubmitDataSetTaskRequest) SetIdType(v int32) *SubmitDataSetTaskRequest {
	s.IdType = &v
	return s
}

func (s *SubmitDataSetTaskRequest) SetName(v string) *SubmitDataSetTaskRequest {
	s.Name = &v
	return s
}

func (s *SubmitDataSetTaskRequest) SetOssFileName(v string) *SubmitDataSetTaskRequest {
	s.OssFileName = &v
	return s
}

type SubmitDataSetTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SubmitDataSetTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitDataSetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDataSetTaskResponseBody) SetCode(v string) *SubmitDataSetTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitDataSetTaskResponseBody) SetData(v int64) *SubmitDataSetTaskResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitDataSetTaskResponseBody) SetMsg(v string) *SubmitDataSetTaskResponseBody {
	s.Msg = &v
	return s
}

func (s *SubmitDataSetTaskResponseBody) SetSuccess(v bool) *SubmitDataSetTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitDataSetTaskResponseBody) SetRequestId(v string) *SubmitDataSetTaskResponseBody {
	s.RequestId = &v
	return s
}

type SubmitDataSetTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDataSetTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDataSetTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitDataSetTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitDataSetTaskResponse) SetHeaders(v map[string]*string) *SubmitDataSetTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitDataSetTaskResponse) SetStatusCode(v int32) *SubmitDataSetTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDataSetTaskResponse) SetBody(v *SubmitDataSetTaskResponseBody) *SubmitDataSetTaskResponse {
	s.Body = v
	return s
}

type CancelYydTaskByBcIdRequest struct {
	BcId *int64 `json:"bcId,omitempty" xml:"bcId,omitempty"`
}

func (s CancelYydTaskByBcIdRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelYydTaskByBcIdRequest) GoString() string {
	return s.String()
}

func (s *CancelYydTaskByBcIdRequest) SetBcId(v int64) *CancelYydTaskByBcIdRequest {
	s.BcId = &v
	return s
}

type CancelYydTaskByBcIdResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelYydTaskByBcIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelYydTaskByBcIdResponseBody) GoString() string {
	return s.String()
}

func (s *CancelYydTaskByBcIdResponseBody) SetCode(v string) *CancelYydTaskByBcIdResponseBody {
	s.Code = &v
	return s
}

func (s *CancelYydTaskByBcIdResponseBody) SetData(v bool) *CancelYydTaskByBcIdResponseBody {
	s.Data = &v
	return s
}

func (s *CancelYydTaskByBcIdResponseBody) SetMsg(v string) *CancelYydTaskByBcIdResponseBody {
	s.Msg = &v
	return s
}

func (s *CancelYydTaskByBcIdResponseBody) SetRequestId(v string) *CancelYydTaskByBcIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelYydTaskByBcIdResponseBody) SetSuccess(v bool) *CancelYydTaskByBcIdResponseBody {
	s.Success = &v
	return s
}

type CancelYydTaskByBcIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelYydTaskByBcIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelYydTaskByBcIdResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelYydTaskByBcIdResponse) GoString() string {
	return s.String()
}

func (s *CancelYydTaskByBcIdResponse) SetHeaders(v map[string]*string) *CancelYydTaskByBcIdResponse {
	s.Headers = v
	return s
}

func (s *CancelYydTaskByBcIdResponse) SetStatusCode(v int32) *CancelYydTaskByBcIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelYydTaskByBcIdResponse) SetBody(v *CancelYydTaskByBcIdResponseBody) *CancelYydTaskByBcIdResponse {
	s.Body = v
	return s
}

type CreateYydComputeTaskRequest struct {
	DatasetId *int64  `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Remarks   *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	AppId     *int64  `json:"appId,omitempty" xml:"appId,omitempty"`
}

func (s CreateYydComputeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateYydComputeTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateYydComputeTaskRequest) SetDatasetId(v int64) *CreateYydComputeTaskRequest {
	s.DatasetId = &v
	return s
}

func (s *CreateYydComputeTaskRequest) SetName(v string) *CreateYydComputeTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateYydComputeTaskRequest) SetRemarks(v string) *CreateYydComputeTaskRequest {
	s.Remarks = &v
	return s
}

func (s *CreateYydComputeTaskRequest) SetAppId(v int64) *CreateYydComputeTaskRequest {
	s.AppId = &v
	return s
}

type CreateYydComputeTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateYydComputeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateYydComputeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateYydComputeTaskResponseBody) SetCode(v string) *CreateYydComputeTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateYydComputeTaskResponseBody) SetData(v int64) *CreateYydComputeTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateYydComputeTaskResponseBody) SetMsg(v string) *CreateYydComputeTaskResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateYydComputeTaskResponseBody) SetRequestId(v string) *CreateYydComputeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateYydComputeTaskResponseBody) SetSuccess(v bool) *CreateYydComputeTaskResponseBody {
	s.Success = &v
	return s
}

type CreateYydComputeTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateYydComputeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateYydComputeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateYydComputeTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateYydComputeTaskResponse) SetHeaders(v map[string]*string) *CreateYydComputeTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateYydComputeTaskResponse) SetStatusCode(v int32) *CreateYydComputeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateYydComputeTaskResponse) SetBody(v *CreateYydComputeTaskResponseBody) *CreateYydComputeTaskResponse {
	s.Body = v
	return s
}

type CreateYydDataSetRequest struct {
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OssFileName *string `json:"ossFileName,omitempty" xml:"ossFileName,omitempty"`
}

func (s CreateYydDataSetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateYydDataSetRequest) GoString() string {
	return s.String()
}

func (s *CreateYydDataSetRequest) SetName(v string) *CreateYydDataSetRequest {
	s.Name = &v
	return s
}

func (s *CreateYydDataSetRequest) SetOssFileName(v string) *CreateYydDataSetRequest {
	s.OssFileName = &v
	return s
}

type CreateYydDataSetResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateYydDataSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateYydDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateYydDataSetResponseBody) SetCode(v string) *CreateYydDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *CreateYydDataSetResponseBody) SetData(v int64) *CreateYydDataSetResponseBody {
	s.Data = &v
	return s
}

func (s *CreateYydDataSetResponseBody) SetMsg(v string) *CreateYydDataSetResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateYydDataSetResponseBody) SetSuccess(v bool) *CreateYydDataSetResponseBody {
	s.Success = &v
	return s
}

func (s *CreateYydDataSetResponseBody) SetRequestId(v string) *CreateYydDataSetResponseBody {
	s.RequestId = &v
	return s
}

type CreateYydDataSetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateYydDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateYydDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateYydDataSetResponse) GoString() string {
	return s.String()
}

func (s *CreateYydDataSetResponse) SetHeaders(v map[string]*string) *CreateYydDataSetResponse {
	s.Headers = v
	return s
}

func (s *CreateYydDataSetResponse) SetStatusCode(v int32) *CreateYydDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateYydDataSetResponse) SetBody(v *CreateYydDataSetResponseBody) *CreateYydDataSetResponse {
	s.Body = v
	return s
}

type ListYydComputeTaskListResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListYydComputeTaskListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Msg       *string                                   `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListYydComputeTaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListYydComputeTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *ListYydComputeTaskListResponseBody) SetCode(v string) *ListYydComputeTaskListResponseBody {
	s.Code = &v
	return s
}

func (s *ListYydComputeTaskListResponseBody) SetData(v []*ListYydComputeTaskListResponseBodyData) *ListYydComputeTaskListResponseBody {
	s.Data = v
	return s
}

func (s *ListYydComputeTaskListResponseBody) SetMsg(v string) *ListYydComputeTaskListResponseBody {
	s.Msg = &v
	return s
}

func (s *ListYydComputeTaskListResponseBody) SetRequestId(v string) *ListYydComputeTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListYydComputeTaskListResponseBody) SetSuccess(v bool) *ListYydComputeTaskListResponseBody {
	s.Success = &v
	return s
}

type ListYydComputeTaskListResponseBodyData struct {
	AppId       *int64  `json:"appId,omitempty" xml:"appId,omitempty"`
	BcId        *int64  `json:"bcId,omitempty" xml:"bcId,omitempty"`
	GmtCreate   *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Remarks     *string `json:"remarks,omitempty" xml:"remarks,omitempty"`
	State       *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListYydComputeTaskListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListYydComputeTaskListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListYydComputeTaskListResponseBodyData) SetAppId(v int64) *ListYydComputeTaskListResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListYydComputeTaskListResponseBodyData) SetBcId(v int64) *ListYydComputeTaskListResponseBodyData {
	s.BcId = &v
	return s
}

func (s *ListYydComputeTaskListResponseBodyData) SetGmtCreate(v string) *ListYydComputeTaskListResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListYydComputeTaskListResponseBodyData) SetGmtModified(v string) *ListYydComputeTaskListResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListYydComputeTaskListResponseBodyData) SetName(v string) *ListYydComputeTaskListResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListYydComputeTaskListResponseBodyData) SetRemarks(v string) *ListYydComputeTaskListResponseBodyData {
	s.Remarks = &v
	return s
}

func (s *ListYydComputeTaskListResponseBodyData) SetState(v string) *ListYydComputeTaskListResponseBodyData {
	s.State = &v
	return s
}

type ListYydComputeTaskListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListYydComputeTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListYydComputeTaskListResponse) String() string {
	return tea.Prettify(s)
}

func (s ListYydComputeTaskListResponse) GoString() string {
	return s.String()
}

func (s *ListYydComputeTaskListResponse) SetHeaders(v map[string]*string) *ListYydComputeTaskListResponse {
	s.Headers = v
	return s
}

func (s *ListYydComputeTaskListResponse) SetStatusCode(v int32) *ListYydComputeTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *ListYydComputeTaskListResponse) SetBody(v *ListYydComputeTaskListResponseBody) *ListYydComputeTaskListResponse {
	s.Body = v
	return s
}

type ListYydDataSetResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListYydDataSetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Msg       *string                           `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListYydDataSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListYydDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *ListYydDataSetResponseBody) SetCode(v string) *ListYydDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *ListYydDataSetResponseBody) SetData(v []*ListYydDataSetResponseBodyData) *ListYydDataSetResponseBody {
	s.Data = v
	return s
}

func (s *ListYydDataSetResponseBody) SetMsg(v string) *ListYydDataSetResponseBody {
	s.Msg = &v
	return s
}

func (s *ListYydDataSetResponseBody) SetSuccess(v bool) *ListYydDataSetResponseBody {
	s.Success = &v
	return s
}

func (s *ListYydDataSetResponseBody) SetRequestId(v string) *ListYydDataSetResponseBody {
	s.RequestId = &v
	return s
}

type ListYydDataSetResponseBodyData struct {
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DataSetType *int32  `json:"dataSetType,omitempty" xml:"dataSetType,omitempty"`
	DatasetId   *int64  `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	IdTypeDesc  *string `json:"idTypeDesc,omitempty" xml:"idTypeDesc,omitempty"`
	Message     *string `json:"message,omitempty" xml:"message,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	StatusDesc  *string `json:"statusDesc,omitempty" xml:"statusDesc,omitempty"`
}

func (s ListYydDataSetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListYydDataSetResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListYydDataSetResponseBodyData) SetCreateTime(v string) *ListYydDataSetResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListYydDataSetResponseBodyData) SetDataSetType(v int32) *ListYydDataSetResponseBodyData {
	s.DataSetType = &v
	return s
}

func (s *ListYydDataSetResponseBodyData) SetDatasetId(v int64) *ListYydDataSetResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *ListYydDataSetResponseBodyData) SetIdTypeDesc(v string) *ListYydDataSetResponseBodyData {
	s.IdTypeDesc = &v
	return s
}

func (s *ListYydDataSetResponseBodyData) SetMessage(v string) *ListYydDataSetResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListYydDataSetResponseBodyData) SetName(v string) *ListYydDataSetResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListYydDataSetResponseBodyData) SetStatusDesc(v string) *ListYydDataSetResponseBodyData {
	s.StatusDesc = &v
	return s
}

type ListYydDataSetResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListYydDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListYydDataSetResponse) String() string {
	return tea.Prettify(s)
}

func (s ListYydDataSetResponse) GoString() string {
	return s.String()
}

func (s *ListYydDataSetResponse) SetHeaders(v map[string]*string) *ListYydDataSetResponse {
	s.Headers = v
	return s
}

func (s *ListYydDataSetResponse) SetStatusCode(v int32) *ListYydDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListYydDataSetResponse) SetBody(v *ListYydDataSetResponseBody) *ListYydDataSetResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("umeng-finplus"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CancelComputeTaskByBcIdWithOptions(request *CancelComputeTaskByBcIdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelComputeTaskByBcIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BcId)) {
		query["bcId"] = request.BcId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelComputeTaskByBcId"),
		Version:     tea.String("2021-11-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/batch_compute/cancelComputeTaskByBcId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelComputeTaskByBcIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelComputeTaskByBcId(request *CancelComputeTaskByBcIdRequest) (_result *CancelComputeTaskByBcIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelComputeTaskByBcIdResponse{}
	_body, _err := client.CancelComputeTaskByBcIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateComputeTaskByDataSetIdWithOptions(request *CreateComputeTaskByDataSetIdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateComputeTaskByDataSetIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BatchInfoForm)) {
		body["BatchInfoForm"] = request.BatchInfoForm
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		body["DatasetId"] = request.DatasetId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Remarks)) {
		body["Remarks"] = request.Remarks
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateComputeTaskByDataSetId"),
		Version:     tea.String("2021-11-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/batch_compute/createComputeTaskByDataSetId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateComputeTaskByDataSetIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateComputeTaskByDataSetId(request *CreateComputeTaskByDataSetIdRequest) (_result *CreateComputeTaskByDataSetIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateComputeTaskByDataSetIdResponse{}
	_body, _err := client.CreateComputeTaskByDataSetIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateComputeTaskByMultiDataSetIdWithOptions(request *CreateComputeTaskByMultiDataSetIdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateComputeTaskByMultiDataSetIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.DataSetIds)) {
		body["dataSetIds"] = request.DataSetIds
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Remarks)) {
		body["remarks"] = request.Remarks
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateComputeTaskByMultiDataSetId"),
		Version:     tea.String("2021-11-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/batch_compute/createComputeTaskByMultiDataSetId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateComputeTaskByMultiDataSetIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateComputeTaskByMultiDataSetId(request *CreateComputeTaskByMultiDataSetIdRequest) (_result *CreateComputeTaskByMultiDataSetIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateComputeTaskByMultiDataSetIdResponse{}
	_body, _err := client.CreateComputeTaskByMultiDataSetIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAvailableDataSetListWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAvailableDataSetListResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAvailableDataSetList"),
		Version:     tea.String("2021-11-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/batch_compute/getAvailableDataSetList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAvailableDataSetListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAvailableDataSetList() (_result *GetAvailableDataSetListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAvailableDataSetListResponse{}
	_body, _err := client.GetAvailableDataSetListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetComputeResultWithOptions(request *GetComputeResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetComputeResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BcId)) {
		query["bcId"] = request.BcId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetComputeResult"),
		Version:     tea.String("2021-11-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/batch_compute/getComputeTaskResult"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetComputeResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetComputeResult(request *GetComputeResultRequest) (_result *GetComputeResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetComputeResultResponse{}
	_body, _err := client.GetComputeResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataSetStatusWithOptions(request *GetDataSetStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDataSetStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSetId)) {
		query["dataSetId"] = request.DataSetId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataSetStatus"),
		Version:     tea.String("2021-11-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/batch_compute/getDataSetStatus"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataSetStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataSetStatus(request *GetDataSetStatusRequest) (_result *GetDataSetStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDataSetStatusResponse{}
	_body, _err := client.GetDataSetStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataSetStsAKWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDataSetStsAKResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataSetStsAK"),
		Version:     tea.String("2021-11-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/batch_compute/getDataSetStsAk"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataSetStsAKResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataSetStsAK() (_result *GetDataSetStsAKResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDataSetStsAKResponse{}
	_body, _err := client.GetDataSetStsAKWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDataSetTaskWithOptions(request *SubmitDataSetTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitDataSetTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSetType)) {
		body["dataSetType"] = request.DataSetType
	}

	if !tea.BoolValue(util.IsUnset(request.IdType)) {
		body["idType"] = request.IdType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OssFileName)) {
		body["ossFileName"] = request.OssFileName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitDataSetTask"),
		Version:     tea.String("2021-11-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/batch_compute/submitDataSetTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitDataSetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitDataSetTask(request *SubmitDataSetTaskRequest) (_result *SubmitDataSetTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitDataSetTaskResponse{}
	_body, _err := client.SubmitDataSetTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelYydTaskByBcIdWithOptions(request *CancelYydTaskByBcIdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelYydTaskByBcIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BcId)) {
		query["bcId"] = request.BcId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("cancelYydTaskByBcId"),
		Version:     tea.String("2021-11-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/batch_compute/yyd/task/cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelYydTaskByBcIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelYydTaskByBcId(request *CancelYydTaskByBcIdRequest) (_result *CancelYydTaskByBcIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelYydTaskByBcIdResponse{}
	_body, _err := client.CancelYydTaskByBcIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateYydComputeTaskWithOptions(request *CreateYydComputeTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateYydComputeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		body["DatasetId"] = request.DatasetId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Remarks)) {
		body["Remarks"] = request.Remarks
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("createYydComputeTask"),
		Version:     tea.String("2021-11-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/batch_compute/yyd/task/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateYydComputeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateYydComputeTask(request *CreateYydComputeTaskRequest) (_result *CreateYydComputeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateYydComputeTaskResponse{}
	_body, _err := client.CreateYydComputeTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateYydDataSetWithOptions(request *CreateYydDataSetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateYydDataSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OssFileName)) {
		body["ossFileName"] = request.OssFileName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("createYydDataSet"),
		Version:     tea.String("2021-11-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/batch_compute/yyd/dataset/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateYydDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateYydDataSet(request *CreateYydDataSetRequest) (_result *CreateYydDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateYydDataSetResponse{}
	_body, _err := client.CreateYydDataSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListYydComputeTaskListWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListYydComputeTaskListResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("listYydComputeTaskList"),
		Version:     tea.String("2021-11-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/batch_compute/yyd/task/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListYydComputeTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListYydComputeTaskList() (_result *ListYydComputeTaskListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListYydComputeTaskListResponse{}
	_body, _err := client.ListYydComputeTaskListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListYydDataSetWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListYydDataSetResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("listYydDataSet"),
		Version:     tea.String("2021-11-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/batch_compute/yyd/dataset/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListYydDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListYydDataSet() (_result *ListYydDataSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListYydDataSetResponse{}
	_body, _err := client.ListYydDataSetWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
