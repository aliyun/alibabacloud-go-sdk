// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type GetOpPvCustomValuesRequest struct {
	PidLoopParameterId *string `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
	PidLoopId          *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
}

func (s GetOpPvCustomValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOpPvCustomValuesRequest) GoString() string {
	return s.String()
}

func (s *GetOpPvCustomValuesRequest) SetPidLoopParameterId(v string) *GetOpPvCustomValuesRequest {
	s.PidLoopParameterId = &v
	return s
}

func (s *GetOpPvCustomValuesRequest) SetPidLoopId(v string) *GetOpPvCustomValuesRequest {
	s.PidLoopId = &v
	return s
}

type GetOpPvCustomValuesResponse struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Data      *GetOpPvCustomValuesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetOpPvCustomValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOpPvCustomValuesResponse) GoString() string {
	return s.String()
}

func (s *GetOpPvCustomValuesResponse) SetRequestId(v string) *GetOpPvCustomValuesResponse {
	s.RequestId = &v
	return s
}

func (s *GetOpPvCustomValuesResponse) SetMessage(v string) *GetOpPvCustomValuesResponse {
	s.Message = &v
	return s
}

func (s *GetOpPvCustomValuesResponse) SetCode(v string) *GetOpPvCustomValuesResponse {
	s.Code = &v
	return s
}

func (s *GetOpPvCustomValuesResponse) SetSuccess(v bool) *GetOpPvCustomValuesResponse {
	s.Success = &v
	return s
}

func (s *GetOpPvCustomValuesResponse) SetData(v *GetOpPvCustomValuesResponseData) *GetOpPvCustomValuesResponse {
	s.Data = v
	return s
}

type GetOpPvCustomValuesResponseData struct {
	Status             *bool                                              `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	OppvCustomDataInfo *GetOpPvCustomValuesResponseDataOppvCustomDataInfo `json:"OppvCustomDataInfo,omitempty" xml:"OppvCustomDataInfo,omitempty" require:"true" type:"Struct"`
}

func (s GetOpPvCustomValuesResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetOpPvCustomValuesResponseData) GoString() string {
	return s.String()
}

func (s *GetOpPvCustomValuesResponseData) SetStatus(v bool) *GetOpPvCustomValuesResponseData {
	s.Status = &v
	return s
}

func (s *GetOpPvCustomValuesResponseData) SetOppvCustomDataInfo(v *GetOpPvCustomValuesResponseDataOppvCustomDataInfo) *GetOpPvCustomValuesResponseData {
	s.OppvCustomDataInfo = v
	return s
}

type GetOpPvCustomValuesResponseDataOppvCustomDataInfo struct {
	Opzdy []*GetOpPvCustomValuesResponseDataOppvCustomDataInfoOpzdy `json:"Opzdy,omitempty" xml:"Opzdy,omitempty" require:"true" type:"Repeated"`
	Pvzdy []*GetOpPvCustomValuesResponseDataOppvCustomDataInfoPvzdy `json:"Pvzdy,omitempty" xml:"Pvzdy,omitempty" require:"true" type:"Repeated"`
}

func (s GetOpPvCustomValuesResponseDataOppvCustomDataInfo) String() string {
	return tea.Prettify(s)
}

func (s GetOpPvCustomValuesResponseDataOppvCustomDataInfo) GoString() string {
	return s.String()
}

func (s *GetOpPvCustomValuesResponseDataOppvCustomDataInfo) SetOpzdy(v []*GetOpPvCustomValuesResponseDataOppvCustomDataInfoOpzdy) *GetOpPvCustomValuesResponseDataOppvCustomDataInfo {
	s.Opzdy = v
	return s
}

func (s *GetOpPvCustomValuesResponseDataOppvCustomDataInfo) SetPvzdy(v []*GetOpPvCustomValuesResponseDataOppvCustomDataInfoPvzdy) *GetOpPvCustomValuesResponseDataOppvCustomDataInfo {
	s.Pvzdy = v
	return s
}

type GetOpPvCustomValuesResponseDataOppvCustomDataInfoOpzdy struct {
	Xlabel  *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel  *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name    *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Quality *int     `json:"Quality,omitempty" xml:"Quality,omitempty" require:"true"`
}

func (s GetOpPvCustomValuesResponseDataOppvCustomDataInfoOpzdy) String() string {
	return tea.Prettify(s)
}

func (s GetOpPvCustomValuesResponseDataOppvCustomDataInfoOpzdy) GoString() string {
	return s.String()
}

func (s *GetOpPvCustomValuesResponseDataOppvCustomDataInfoOpzdy) SetXlabel(v float32) *GetOpPvCustomValuesResponseDataOppvCustomDataInfoOpzdy {
	s.Xlabel = &v
	return s
}

func (s *GetOpPvCustomValuesResponseDataOppvCustomDataInfoOpzdy) SetYlabel(v float32) *GetOpPvCustomValuesResponseDataOppvCustomDataInfoOpzdy {
	s.Ylabel = &v
	return s
}

func (s *GetOpPvCustomValuesResponseDataOppvCustomDataInfoOpzdy) SetName(v string) *GetOpPvCustomValuesResponseDataOppvCustomDataInfoOpzdy {
	s.Name = &v
	return s
}

func (s *GetOpPvCustomValuesResponseDataOppvCustomDataInfoOpzdy) SetQuality(v int) *GetOpPvCustomValuesResponseDataOppvCustomDataInfoOpzdy {
	s.Quality = &v
	return s
}

type GetOpPvCustomValuesResponseDataOppvCustomDataInfoPvzdy struct {
	Xlabel  *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel  *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name    *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Quality *int     `json:"Quality,omitempty" xml:"Quality,omitempty" require:"true"`
}

func (s GetOpPvCustomValuesResponseDataOppvCustomDataInfoPvzdy) String() string {
	return tea.Prettify(s)
}

func (s GetOpPvCustomValuesResponseDataOppvCustomDataInfoPvzdy) GoString() string {
	return s.String()
}

func (s *GetOpPvCustomValuesResponseDataOppvCustomDataInfoPvzdy) SetXlabel(v float32) *GetOpPvCustomValuesResponseDataOppvCustomDataInfoPvzdy {
	s.Xlabel = &v
	return s
}

func (s *GetOpPvCustomValuesResponseDataOppvCustomDataInfoPvzdy) SetYlabel(v float32) *GetOpPvCustomValuesResponseDataOppvCustomDataInfoPvzdy {
	s.Ylabel = &v
	return s
}

func (s *GetOpPvCustomValuesResponseDataOppvCustomDataInfoPvzdy) SetName(v string) *GetOpPvCustomValuesResponseDataOppvCustomDataInfoPvzdy {
	s.Name = &v
	return s
}

func (s *GetOpPvCustomValuesResponseDataOppvCustomDataInfoPvzdy) SetQuality(v int) *GetOpPvCustomValuesResponseDataOppvCustomDataInfoPvzdy {
	s.Quality = &v
	return s
}

type SubmitPidLoopEvaluationRequest struct {
	PidLoopIdList map[string]interface{} `json:"PidLoopIdList,omitempty" xml:"PidLoopIdList,omitempty" require:"true"`
	PidProjectId  *string                `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	StartTime     *string                `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime       *string                `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s SubmitPidLoopEvaluationRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitPidLoopEvaluationRequest) GoString() string {
	return s.String()
}

func (s *SubmitPidLoopEvaluationRequest) SetPidLoopIdList(v map[string]interface{}) *SubmitPidLoopEvaluationRequest {
	s.PidLoopIdList = v
	return s
}

func (s *SubmitPidLoopEvaluationRequest) SetPidProjectId(v string) *SubmitPidLoopEvaluationRequest {
	s.PidProjectId = &v
	return s
}

func (s *SubmitPidLoopEvaluationRequest) SetStartTime(v string) *SubmitPidLoopEvaluationRequest {
	s.StartTime = &v
	return s
}

func (s *SubmitPidLoopEvaluationRequest) SetEndTime(v string) *SubmitPidLoopEvaluationRequest {
	s.EndTime = &v
	return s
}

type SubmitPidLoopEvaluationResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s SubmitPidLoopEvaluationResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitPidLoopEvaluationResponse) GoString() string {
	return s.String()
}

func (s *SubmitPidLoopEvaluationResponse) SetRequestId(v string) *SubmitPidLoopEvaluationResponse {
	s.RequestId = &v
	return s
}

func (s *SubmitPidLoopEvaluationResponse) SetCode(v string) *SubmitPidLoopEvaluationResponse {
	s.Code = &v
	return s
}

func (s *SubmitPidLoopEvaluationResponse) SetMessage(v string) *SubmitPidLoopEvaluationResponse {
	s.Message = &v
	return s
}

func (s *SubmitPidLoopEvaluationResponse) SetSuccess(v bool) *SubmitPidLoopEvaluationResponse {
	s.Success = &v
	return s
}

type GetDefaultAdjustValuesRequest struct {
	PidLoopParameterId *string `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty"`
	PidLoopId          *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty"`
}

func (s GetDefaultAdjustValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultAdjustValuesRequest) GoString() string {
	return s.String()
}

func (s *GetDefaultAdjustValuesRequest) SetPidLoopParameterId(v string) *GetDefaultAdjustValuesRequest {
	s.PidLoopParameterId = &v
	return s
}

func (s *GetDefaultAdjustValuesRequest) SetPidLoopId(v string) *GetDefaultAdjustValuesRequest {
	s.PidLoopId = &v
	return s
}

type GetDefaultAdjustValuesResponse struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Data      *GetDefaultAdjustValuesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetDefaultAdjustValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultAdjustValuesResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultAdjustValuesResponse) SetRequestId(v string) *GetDefaultAdjustValuesResponse {
	s.RequestId = &v
	return s
}

func (s *GetDefaultAdjustValuesResponse) SetMessage(v string) *GetDefaultAdjustValuesResponse {
	s.Message = &v
	return s
}

func (s *GetDefaultAdjustValuesResponse) SetCode(v string) *GetDefaultAdjustValuesResponse {
	s.Code = &v
	return s
}

func (s *GetDefaultAdjustValuesResponse) SetSuccess(v bool) *GetDefaultAdjustValuesResponse {
	s.Success = &v
	return s
}

func (s *GetDefaultAdjustValuesResponse) SetData(v *GetDefaultAdjustValuesResponseData) *GetDefaultAdjustValuesResponse {
	s.Data = v
	return s
}

type GetDefaultAdjustValuesResponseData struct {
	Status                *bool                                                    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	DefaultAdjustDataInfo *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfo `json:"DefaultAdjustDataInfo,omitempty" xml:"DefaultAdjustDataInfo,omitempty" require:"true" type:"Struct"`
}

func (s GetDefaultAdjustValuesResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultAdjustValuesResponseData) GoString() string {
	return s.String()
}

func (s *GetDefaultAdjustValuesResponseData) SetStatus(v bool) *GetDefaultAdjustValuesResponseData {
	s.Status = &v
	return s
}

func (s *GetDefaultAdjustValuesResponseData) SetDefaultAdjustDataInfo(v *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfo) *GetDefaultAdjustValuesResponseData {
	s.DefaultAdjustDataInfo = v
	return s
}

type GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfo struct {
	ManualCtrl    *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualCtrl    `json:"ManualCtrl,omitempty" xml:"ManualCtrl,omitempty" require:"true" type:"Struct"`
	Perform       *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform       `json:"Perform,omitempty" xml:"Perform,omitempty" require:"true" type:"Struct"`
	ManualPerform *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform `json:"ManualPerform,omitempty" xml:"ManualPerform,omitempty" require:"true" type:"Struct"`
}

func (s GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfo) GoString() string {
	return s.String()
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfo) SetManualCtrl(v *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualCtrl) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfo {
	s.ManualCtrl = v
	return s
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfo) SetPerform(v *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfo {
	s.Perform = v
	return s
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfo) SetManualPerform(v *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfo {
	s.ManualPerform = v
	return s
}

type GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualCtrl struct {
	Kp *float32 `json:"Kp,omitempty" xml:"Kp,omitempty" require:"true"`
	Ti *float32 `json:"Ti,omitempty" xml:"Ti,omitempty" require:"true"`
	Td *float32 `json:"Td,omitempty" xml:"Td,omitempty" require:"true"`
}

func (s GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualCtrl) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualCtrl) GoString() string {
	return s.String()
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualCtrl) SetKp(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualCtrl {
	s.Kp = &v
	return s
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualCtrl) SetTi(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualCtrl {
	s.Ti = &v
	return s
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualCtrl) SetTd(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualCtrl {
	s.Td = &v
	return s
}

type GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform struct {
	RaiseTime   *float32 `json:"RaiseTime,omitempty" xml:"RaiseTime,omitempty" require:"true"`
	FinalValue  *float32 `json:"FinalValue,omitempty" xml:"FinalValue,omitempty" require:"true"`
	OverValue   *float32 `json:"OverValue,omitempty" xml:"OverValue,omitempty" require:"true"`
	StableError *float32 `json:"StableError,omitempty" xml:"StableError,omitempty" require:"true"`
	AdjustTime  *float32 `json:"AdjustTime,omitempty" xml:"AdjustTime,omitempty" require:"true"`
	Dynamic     *float32 `json:"Dynamic,omitempty" xml:"Dynamic,omitempty" require:"true"`
	Robust      *float32 `json:"Robust,omitempty" xml:"Robust,omitempty" require:"true"`
}

func (s GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform) GoString() string {
	return s.String()
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform) SetRaiseTime(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform {
	s.RaiseTime = &v
	return s
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform) SetFinalValue(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform {
	s.FinalValue = &v
	return s
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform) SetOverValue(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform {
	s.OverValue = &v
	return s
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform) SetStableError(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform {
	s.StableError = &v
	return s
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform) SetAdjustTime(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform {
	s.AdjustTime = &v
	return s
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform) SetDynamic(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform {
	s.Dynamic = &v
	return s
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform) SetRobust(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoPerform {
	s.Robust = &v
	return s
}

type GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform struct {
	RaiseTime   *float32 `json:"RaiseTime,omitempty" xml:"RaiseTime,omitempty" require:"true"`
	FinalValue  *float32 `json:"FinalValue,omitempty" xml:"FinalValue,omitempty" require:"true"`
	OverValue   *float32 `json:"OverValue,omitempty" xml:"OverValue,omitempty" require:"true"`
	StableError *float32 `json:"StableError,omitempty" xml:"StableError,omitempty" require:"true"`
	AdjustTime  *float32 `json:"AdjustTime,omitempty" xml:"AdjustTime,omitempty" require:"true"`
	Dynamic     *float32 `json:"Dynamic,omitempty" xml:"Dynamic,omitempty" require:"true"`
	Robust      *float32 `json:"Robust,omitempty" xml:"Robust,omitempty" require:"true"`
}

func (s GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform) GoString() string {
	return s.String()
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform) SetRaiseTime(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform {
	s.RaiseTime = &v
	return s
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform) SetFinalValue(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform {
	s.FinalValue = &v
	return s
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform) SetOverValue(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform {
	s.OverValue = &v
	return s
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform) SetStableError(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform {
	s.StableError = &v
	return s
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform) SetAdjustTime(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform {
	s.AdjustTime = &v
	return s
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform) SetDynamic(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform {
	s.Dynamic = &v
	return s
}

func (s *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform) SetRobust(v float32) *GetDefaultAdjustValuesResponseDataDefaultAdjustDataInfoManualPerform {
	s.Robust = &v
	return s
}

type GetPvOpAdjustValuesRequest struct {
	PidLoopParameterId *string `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
	PidLoopId          *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
}

func (s GetPvOpAdjustValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPvOpAdjustValuesRequest) GoString() string {
	return s.String()
}

func (s *GetPvOpAdjustValuesRequest) SetPidLoopParameterId(v string) *GetPvOpAdjustValuesRequest {
	s.PidLoopParameterId = &v
	return s
}

func (s *GetPvOpAdjustValuesRequest) SetPidLoopId(v string) *GetPvOpAdjustValuesRequest {
	s.PidLoopId = &v
	return s
}

type GetPvOpAdjustValuesResponse struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Data      *GetPvOpAdjustValuesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetPvOpAdjustValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPvOpAdjustValuesResponse) GoString() string {
	return s.String()
}

func (s *GetPvOpAdjustValuesResponse) SetRequestId(v string) *GetPvOpAdjustValuesResponse {
	s.RequestId = &v
	return s
}

func (s *GetPvOpAdjustValuesResponse) SetMessage(v string) *GetPvOpAdjustValuesResponse {
	s.Message = &v
	return s
}

func (s *GetPvOpAdjustValuesResponse) SetCode(v string) *GetPvOpAdjustValuesResponse {
	s.Code = &v
	return s
}

func (s *GetPvOpAdjustValuesResponse) SetSuccess(v bool) *GetPvOpAdjustValuesResponse {
	s.Success = &v
	return s
}

func (s *GetPvOpAdjustValuesResponse) SetData(v *GetPvOpAdjustValuesResponseData) *GetPvOpAdjustValuesResponse {
	s.Data = v
	return s
}

type GetPvOpAdjustValuesResponseData struct {
	Status             *bool                                              `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	PvopAdjustDataInfo *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfo `json:"PvopAdjustDataInfo,omitempty" xml:"PvopAdjustDataInfo,omitempty" require:"true" type:"Struct"`
}

func (s GetPvOpAdjustValuesResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetPvOpAdjustValuesResponseData) GoString() string {
	return s.String()
}

func (s *GetPvOpAdjustValuesResponseData) SetStatus(v bool) *GetPvOpAdjustValuesResponseData {
	s.Status = &v
	return s
}

func (s *GetPvOpAdjustValuesResponseData) SetPvopAdjustDataInfo(v *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfo) *GetPvOpAdjustValuesResponseData {
	s.PvopAdjustDataInfo = v
	return s
}

type GetPvOpAdjustValuesResponseDataPvopAdjustDataInfo struct {
	Op []*GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoOp `json:"Op,omitempty" xml:"Op,omitempty" require:"true" type:"Repeated"`
	Pv []*GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoPv `json:"Pv,omitempty" xml:"Pv,omitempty" require:"true" type:"Repeated"`
	Sp []*GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoSp `json:"Sp,omitempty" xml:"Sp,omitempty" require:"true" type:"Repeated"`
}

func (s GetPvOpAdjustValuesResponseDataPvopAdjustDataInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPvOpAdjustValuesResponseDataPvopAdjustDataInfo) GoString() string {
	return s.String()
}

func (s *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfo) SetOp(v []*GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoOp) *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfo {
	s.Op = v
	return s
}

func (s *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfo) SetPv(v []*GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoPv) *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfo {
	s.Pv = v
	return s
}

func (s *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfo) SetSp(v []*GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoSp) *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfo {
	s.Sp = v
	return s
}

type GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoOp struct {
	Xlabel  *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel  *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name    *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Quality *int     `json:"Quality,omitempty" xml:"Quality,omitempty" require:"true"`
}

func (s GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoOp) String() string {
	return tea.Prettify(s)
}

func (s GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoOp) GoString() string {
	return s.String()
}

func (s *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoOp) SetXlabel(v float32) *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoOp {
	s.Xlabel = &v
	return s
}

func (s *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoOp) SetYlabel(v float32) *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoOp {
	s.Ylabel = &v
	return s
}

func (s *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoOp) SetName(v string) *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoOp {
	s.Name = &v
	return s
}

func (s *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoOp) SetQuality(v int) *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoOp {
	s.Quality = &v
	return s
}

type GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoPv struct {
	Xlabel  *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel  *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name    *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Quality *int     `json:"Quality,omitempty" xml:"Quality,omitempty" require:"true"`
}

func (s GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoPv) String() string {
	return tea.Prettify(s)
}

func (s GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoPv) GoString() string {
	return s.String()
}

func (s *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoPv) SetXlabel(v float32) *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoPv {
	s.Xlabel = &v
	return s
}

func (s *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoPv) SetYlabel(v float32) *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoPv {
	s.Ylabel = &v
	return s
}

func (s *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoPv) SetName(v string) *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoPv {
	s.Name = &v
	return s
}

func (s *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoPv) SetQuality(v int) *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoPv {
	s.Quality = &v
	return s
}

type GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoSp struct {
	Xlabel  *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel  *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name    *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Quality *int     `json:"Quality,omitempty" xml:"Quality,omitempty" require:"true"`
}

func (s GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoSp) String() string {
	return tea.Prettify(s)
}

func (s GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoSp) GoString() string {
	return s.String()
}

func (s *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoSp) SetXlabel(v float32) *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoSp {
	s.Xlabel = &v
	return s
}

func (s *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoSp) SetYlabel(v float32) *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoSp {
	s.Ylabel = &v
	return s
}

func (s *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoSp) SetName(v string) *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoSp {
	s.Name = &v
	return s
}

func (s *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoSp) SetQuality(v int) *GetPvOpAdjustValuesResponseDataPvopAdjustDataInfoSp {
	s.Quality = &v
	return s
}

type GetPvCustomSimulationValuesRequest struct {
	PidLoopParameterId *string `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty"`
	PidLoopId          *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty"`
}

func (s GetPvCustomSimulationValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPvCustomSimulationValuesRequest) GoString() string {
	return s.String()
}

func (s *GetPvCustomSimulationValuesRequest) SetPidLoopParameterId(v string) *GetPvCustomSimulationValuesRequest {
	s.PidLoopParameterId = &v
	return s
}

func (s *GetPvCustomSimulationValuesRequest) SetPidLoopId(v string) *GetPvCustomSimulationValuesRequest {
	s.PidLoopId = &v
	return s
}

type GetPvCustomSimulationValuesResponse struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Data      *GetPvCustomSimulationValuesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetPvCustomSimulationValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPvCustomSimulationValuesResponse) GoString() string {
	return s.String()
}

func (s *GetPvCustomSimulationValuesResponse) SetRequestId(v string) *GetPvCustomSimulationValuesResponse {
	s.RequestId = &v
	return s
}

func (s *GetPvCustomSimulationValuesResponse) SetMessage(v string) *GetPvCustomSimulationValuesResponse {
	s.Message = &v
	return s
}

func (s *GetPvCustomSimulationValuesResponse) SetCode(v string) *GetPvCustomSimulationValuesResponse {
	s.Code = &v
	return s
}

func (s *GetPvCustomSimulationValuesResponse) SetSuccess(v bool) *GetPvCustomSimulationValuesResponse {
	s.Success = &v
	return s
}

func (s *GetPvCustomSimulationValuesResponse) SetData(v *GetPvCustomSimulationValuesResponseData) *GetPvCustomSimulationValuesResponse {
	s.Data = v
	return s
}

type GetPvCustomSimulationValuesResponseData struct {
	Status                   *bool                                                            `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	PvCustomSimulateDataInfo *GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfo `json:"PvCustomSimulateDataInfo,omitempty" xml:"PvCustomSimulateDataInfo,omitempty" require:"true" type:"Struct"`
}

func (s GetPvCustomSimulationValuesResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetPvCustomSimulationValuesResponseData) GoString() string {
	return s.String()
}

func (s *GetPvCustomSimulationValuesResponseData) SetStatus(v bool) *GetPvCustomSimulationValuesResponseData {
	s.Status = &v
	return s
}

func (s *GetPvCustomSimulationValuesResponseData) SetPvCustomSimulateDataInfo(v *GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfo) *GetPvCustomSimulationValuesResponseData {
	s.PvCustomSimulateDataInfo = v
	return s
}

type GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfo struct {
	Pv []*GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfoPv `json:"Pv,omitempty" xml:"Pv,omitempty" require:"true" type:"Repeated"`
}

func (s GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfo) GoString() string {
	return s.String()
}

func (s *GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfo) SetPv(v []*GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfoPv) *GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfo {
	s.Pv = v
	return s
}

type GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfoPv struct {
	Xlabel  *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel  *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name    *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Quality *int     `json:"Quality,omitempty" xml:"Quality,omitempty" require:"true"`
}

func (s GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfoPv) String() string {
	return tea.Prettify(s)
}

func (s GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfoPv) GoString() string {
	return s.String()
}

func (s *GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfoPv) SetXlabel(v float32) *GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfoPv {
	s.Xlabel = &v
	return s
}

func (s *GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfoPv) SetYlabel(v float32) *GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfoPv {
	s.Ylabel = &v
	return s
}

func (s *GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfoPv) SetName(v string) *GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfoPv {
	s.Name = &v
	return s
}

func (s *GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfoPv) SetQuality(v int) *GetPvCustomSimulationValuesResponseDataPvCustomSimulateDataInfoPv {
	s.Quality = &v
	return s
}

type GetIdentificateValuesRequest struct {
	PidLoopParameterId *string `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
	PidLoopId          *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
}

func (s GetIdentificateValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIdentificateValuesRequest) GoString() string {
	return s.String()
}

func (s *GetIdentificateValuesRequest) SetPidLoopParameterId(v string) *GetIdentificateValuesRequest {
	s.PidLoopParameterId = &v
	return s
}

func (s *GetIdentificateValuesRequest) SetPidLoopId(v string) *GetIdentificateValuesRequest {
	s.PidLoopId = &v
	return s
}

type GetIdentificateValuesResponse struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Data      *GetIdentificateValuesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetIdentificateValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIdentificateValuesResponse) GoString() string {
	return s.String()
}

func (s *GetIdentificateValuesResponse) SetRequestId(v string) *GetIdentificateValuesResponse {
	s.RequestId = &v
	return s
}

func (s *GetIdentificateValuesResponse) SetMessage(v string) *GetIdentificateValuesResponse {
	s.Message = &v
	return s
}

func (s *GetIdentificateValuesResponse) SetCode(v string) *GetIdentificateValuesResponse {
	s.Code = &v
	return s
}

func (s *GetIdentificateValuesResponse) SetSuccess(v bool) *GetIdentificateValuesResponse {
	s.Success = &v
	return s
}

func (s *GetIdentificateValuesResponse) SetData(v *GetIdentificateValuesResponseData) *GetIdentificateValuesResponse {
	s.Data = v
	return s
}

type GetIdentificateValuesResponseData struct {
	Status               *bool                                                  `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	IdentificateDataInfo *GetIdentificateValuesResponseDataIdentificateDataInfo `json:"IdentificateDataInfo,omitempty" xml:"IdentificateDataInfo,omitempty" require:"true" type:"Struct"`
}

func (s GetIdentificateValuesResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetIdentificateValuesResponseData) GoString() string {
	return s.String()
}

func (s *GetIdentificateValuesResponseData) SetStatus(v bool) *GetIdentificateValuesResponseData {
	s.Status = &v
	return s
}

func (s *GetIdentificateValuesResponseData) SetIdentificateDataInfo(v *GetIdentificateValuesResponseDataIdentificateDataInfo) *GetIdentificateValuesResponseData {
	s.IdentificateDataInfo = v
	return s
}

type GetIdentificateValuesResponseDataIdentificateDataInfo struct {
	Dj []*GetIdentificateValuesResponseDataIdentificateDataInfoDj `json:"Dj,omitempty" xml:"Dj,omitempty" require:"true" type:"Repeated"`
	Gj []*GetIdentificateValuesResponseDataIdentificateDataInfoGj `json:"Gj,omitempty" xml:"Gj,omitempty" require:"true" type:"Repeated"`
	Bs []*GetIdentificateValuesResponseDataIdentificateDataInfoBs `json:"Bs,omitempty" xml:"Bs,omitempty" require:"true" type:"Repeated"`
}

func (s GetIdentificateValuesResponseDataIdentificateDataInfo) String() string {
	return tea.Prettify(s)
}

func (s GetIdentificateValuesResponseDataIdentificateDataInfo) GoString() string {
	return s.String()
}

func (s *GetIdentificateValuesResponseDataIdentificateDataInfo) SetDj(v []*GetIdentificateValuesResponseDataIdentificateDataInfoDj) *GetIdentificateValuesResponseDataIdentificateDataInfo {
	s.Dj = v
	return s
}

func (s *GetIdentificateValuesResponseDataIdentificateDataInfo) SetGj(v []*GetIdentificateValuesResponseDataIdentificateDataInfoGj) *GetIdentificateValuesResponseDataIdentificateDataInfo {
	s.Gj = v
	return s
}

func (s *GetIdentificateValuesResponseDataIdentificateDataInfo) SetBs(v []*GetIdentificateValuesResponseDataIdentificateDataInfoBs) *GetIdentificateValuesResponseDataIdentificateDataInfo {
	s.Bs = v
	return s
}

type GetIdentificateValuesResponseDataIdentificateDataInfoDj struct {
	Xlabel  *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel  *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name    *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Quality *int     `json:"Quality,omitempty" xml:"Quality,omitempty" require:"true"`
}

func (s GetIdentificateValuesResponseDataIdentificateDataInfoDj) String() string {
	return tea.Prettify(s)
}

func (s GetIdentificateValuesResponseDataIdentificateDataInfoDj) GoString() string {
	return s.String()
}

func (s *GetIdentificateValuesResponseDataIdentificateDataInfoDj) SetXlabel(v float32) *GetIdentificateValuesResponseDataIdentificateDataInfoDj {
	s.Xlabel = &v
	return s
}

func (s *GetIdentificateValuesResponseDataIdentificateDataInfoDj) SetYlabel(v float32) *GetIdentificateValuesResponseDataIdentificateDataInfoDj {
	s.Ylabel = &v
	return s
}

func (s *GetIdentificateValuesResponseDataIdentificateDataInfoDj) SetName(v string) *GetIdentificateValuesResponseDataIdentificateDataInfoDj {
	s.Name = &v
	return s
}

func (s *GetIdentificateValuesResponseDataIdentificateDataInfoDj) SetQuality(v int) *GetIdentificateValuesResponseDataIdentificateDataInfoDj {
	s.Quality = &v
	return s
}

type GetIdentificateValuesResponseDataIdentificateDataInfoGj struct {
	Xlabel  *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel  *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name    *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Quality *int     `json:"Quality,omitempty" xml:"Quality,omitempty" require:"true"`
}

func (s GetIdentificateValuesResponseDataIdentificateDataInfoGj) String() string {
	return tea.Prettify(s)
}

func (s GetIdentificateValuesResponseDataIdentificateDataInfoGj) GoString() string {
	return s.String()
}

func (s *GetIdentificateValuesResponseDataIdentificateDataInfoGj) SetXlabel(v float32) *GetIdentificateValuesResponseDataIdentificateDataInfoGj {
	s.Xlabel = &v
	return s
}

func (s *GetIdentificateValuesResponseDataIdentificateDataInfoGj) SetYlabel(v float32) *GetIdentificateValuesResponseDataIdentificateDataInfoGj {
	s.Ylabel = &v
	return s
}

func (s *GetIdentificateValuesResponseDataIdentificateDataInfoGj) SetName(v string) *GetIdentificateValuesResponseDataIdentificateDataInfoGj {
	s.Name = &v
	return s
}

func (s *GetIdentificateValuesResponseDataIdentificateDataInfoGj) SetQuality(v int) *GetIdentificateValuesResponseDataIdentificateDataInfoGj {
	s.Quality = &v
	return s
}

type GetIdentificateValuesResponseDataIdentificateDataInfoBs struct {
	Xlabel  *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel  *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name    *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Quality *int     `json:"Quality,omitempty" xml:"Quality,omitempty" require:"true"`
}

func (s GetIdentificateValuesResponseDataIdentificateDataInfoBs) String() string {
	return tea.Prettify(s)
}

func (s GetIdentificateValuesResponseDataIdentificateDataInfoBs) GoString() string {
	return s.String()
}

func (s *GetIdentificateValuesResponseDataIdentificateDataInfoBs) SetXlabel(v float32) *GetIdentificateValuesResponseDataIdentificateDataInfoBs {
	s.Xlabel = &v
	return s
}

func (s *GetIdentificateValuesResponseDataIdentificateDataInfoBs) SetYlabel(v float32) *GetIdentificateValuesResponseDataIdentificateDataInfoBs {
	s.Ylabel = &v
	return s
}

func (s *GetIdentificateValuesResponseDataIdentificateDataInfoBs) SetName(v string) *GetIdentificateValuesResponseDataIdentificateDataInfoBs {
	s.Name = &v
	return s
}

func (s *GetIdentificateValuesResponseDataIdentificateDataInfoBs) SetQuality(v int) *GetIdentificateValuesResponseDataIdentificateDataInfoBs {
	s.Quality = &v
	return s
}

type GetPvIdentSimulationValuesRequest struct {
	PidLoopId          *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
	PidLoopParameterId *string `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
}

func (s GetPvIdentSimulationValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPvIdentSimulationValuesRequest) GoString() string {
	return s.String()
}

func (s *GetPvIdentSimulationValuesRequest) SetPidLoopId(v string) *GetPvIdentSimulationValuesRequest {
	s.PidLoopId = &v
	return s
}

func (s *GetPvIdentSimulationValuesRequest) SetPidLoopParameterId(v string) *GetPvIdentSimulationValuesRequest {
	s.PidLoopParameterId = &v
	return s
}

type GetPvIdentSimulationValuesResponse struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Data      *GetPvIdentSimulationValuesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetPvIdentSimulationValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPvIdentSimulationValuesResponse) GoString() string {
	return s.String()
}

func (s *GetPvIdentSimulationValuesResponse) SetRequestId(v string) *GetPvIdentSimulationValuesResponse {
	s.RequestId = &v
	return s
}

func (s *GetPvIdentSimulationValuesResponse) SetCode(v string) *GetPvIdentSimulationValuesResponse {
	s.Code = &v
	return s
}

func (s *GetPvIdentSimulationValuesResponse) SetMessage(v string) *GetPvIdentSimulationValuesResponse {
	s.Message = &v
	return s
}

func (s *GetPvIdentSimulationValuesResponse) SetSuccess(v bool) *GetPvIdentSimulationValuesResponse {
	s.Success = &v
	return s
}

func (s *GetPvIdentSimulationValuesResponse) SetData(v *GetPvIdentSimulationValuesResponseData) *GetPvIdentSimulationValuesResponse {
	s.Data = v
	return s
}

type GetPvIdentSimulationValuesResponseData struct {
	Status                  *bool                                                          `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	PvIdentSimulateDataInfo *GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfo `json:"PvIdentSimulateDataInfo,omitempty" xml:"PvIdentSimulateDataInfo,omitempty" require:"true" type:"Struct"`
}

func (s GetPvIdentSimulationValuesResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetPvIdentSimulationValuesResponseData) GoString() string {
	return s.String()
}

func (s *GetPvIdentSimulationValuesResponseData) SetStatus(v bool) *GetPvIdentSimulationValuesResponseData {
	s.Status = &v
	return s
}

func (s *GetPvIdentSimulationValuesResponseData) SetPvIdentSimulateDataInfo(v *GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfo) *GetPvIdentSimulationValuesResponseData {
	s.PvIdentSimulateDataInfo = v
	return s
}

type GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfo struct {
	Pv []*GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfoPv `json:"Pv,omitempty" xml:"Pv,omitempty" require:"true" type:"Repeated"`
}

func (s GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfo) GoString() string {
	return s.String()
}

func (s *GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfo) SetPv(v []*GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfoPv) *GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfo {
	s.Pv = v
	return s
}

type GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfoPv struct {
	Xlabel  *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel  *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name    *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Quality *int     `json:"Quality,omitempty" xml:"Quality,omitempty" require:"true"`
}

func (s GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfoPv) String() string {
	return tea.Prettify(s)
}

func (s GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfoPv) GoString() string {
	return s.String()
}

func (s *GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfoPv) SetXlabel(v float32) *GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfoPv {
	s.Xlabel = &v
	return s
}

func (s *GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfoPv) SetYlabel(v float32) *GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfoPv {
	s.Ylabel = &v
	return s
}

func (s *GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfoPv) SetName(v string) *GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfoPv {
	s.Name = &v
	return s
}

func (s *GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfoPv) SetQuality(v int) *GetPvIdentSimulationValuesResponseDataPvIdentSimulateDataInfoPv {
	s.Quality = &v
	return s
}

type GetCustomIdentValuesRequest struct {
	PidLoopParameterId *string `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
	PidLoopId          *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
}

func (s GetCustomIdentValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomIdentValuesRequest) GoString() string {
	return s.String()
}

func (s *GetCustomIdentValuesRequest) SetPidLoopParameterId(v string) *GetCustomIdentValuesRequest {
	s.PidLoopParameterId = &v
	return s
}

func (s *GetCustomIdentValuesRequest) SetPidLoopId(v string) *GetCustomIdentValuesRequest {
	s.PidLoopId = &v
	return s
}

type GetCustomIdentValuesResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Data      *GetCustomIdentValuesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetCustomIdentValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomIdentValuesResponse) GoString() string {
	return s.String()
}

func (s *GetCustomIdentValuesResponse) SetRequestId(v string) *GetCustomIdentValuesResponse {
	s.RequestId = &v
	return s
}

func (s *GetCustomIdentValuesResponse) SetCode(v string) *GetCustomIdentValuesResponse {
	s.Code = &v
	return s
}

func (s *GetCustomIdentValuesResponse) SetMessage(v string) *GetCustomIdentValuesResponse {
	s.Message = &v
	return s
}

func (s *GetCustomIdentValuesResponse) SetSuccess(v bool) *GetCustomIdentValuesResponse {
	s.Success = &v
	return s
}

func (s *GetCustomIdentValuesResponse) SetData(v *GetCustomIdentValuesResponseData) *GetCustomIdentValuesResponse {
	s.Data = v
	return s
}

type GetCustomIdentValuesResponseData struct {
	Status              *bool                                                `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	CustomIdentDataInfo *GetCustomIdentValuesResponseDataCustomIdentDataInfo `json:"CustomIdentDataInfo,omitempty" xml:"CustomIdentDataInfo,omitempty" require:"true" type:"Struct"`
}

func (s GetCustomIdentValuesResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetCustomIdentValuesResponseData) GoString() string {
	return s.String()
}

func (s *GetCustomIdentValuesResponseData) SetStatus(v bool) *GetCustomIdentValuesResponseData {
	s.Status = &v
	return s
}

func (s *GetCustomIdentValuesResponseData) SetCustomIdentDataInfo(v *GetCustomIdentValuesResponseDataCustomIdentDataInfo) *GetCustomIdentValuesResponseData {
	s.CustomIdentDataInfo = v
	return s
}

type GetCustomIdentValuesResponseDataCustomIdentDataInfo struct {
	CustomModelTrend []*GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomModelTrend `json:"CustomModelTrend,omitempty" xml:"CustomModelTrend,omitempty" require:"true" type:"Repeated"`
	CustomResidual   []*GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomResidual   `json:"CustomResidual,omitempty" xml:"CustomResidual,omitempty" require:"true" type:"Repeated"`
}

func (s GetCustomIdentValuesResponseDataCustomIdentDataInfo) String() string {
	return tea.Prettify(s)
}

func (s GetCustomIdentValuesResponseDataCustomIdentDataInfo) GoString() string {
	return s.String()
}

func (s *GetCustomIdentValuesResponseDataCustomIdentDataInfo) SetCustomModelTrend(v []*GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomModelTrend) *GetCustomIdentValuesResponseDataCustomIdentDataInfo {
	s.CustomModelTrend = v
	return s
}

func (s *GetCustomIdentValuesResponseDataCustomIdentDataInfo) SetCustomResidual(v []*GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomResidual) *GetCustomIdentValuesResponseDataCustomIdentDataInfo {
	s.CustomResidual = v
	return s
}

type GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomModelTrend struct {
	Xlabel  *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel  *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name    *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Quality *int     `json:"Quality,omitempty" xml:"Quality,omitempty" require:"true"`
}

func (s GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomModelTrend) String() string {
	return tea.Prettify(s)
}

func (s GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomModelTrend) GoString() string {
	return s.String()
}

func (s *GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomModelTrend) SetXlabel(v float32) *GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomModelTrend {
	s.Xlabel = &v
	return s
}

func (s *GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomModelTrend) SetYlabel(v float32) *GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomModelTrend {
	s.Ylabel = &v
	return s
}

func (s *GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomModelTrend) SetName(v string) *GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomModelTrend {
	s.Name = &v
	return s
}

func (s *GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomModelTrend) SetQuality(v int) *GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomModelTrend {
	s.Quality = &v
	return s
}

type GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomResidual struct {
	Xlabel  *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel  *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name    *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Quality *int     `json:"Quality,omitempty" xml:"Quality,omitempty" require:"true"`
}

func (s GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomResidual) String() string {
	return tea.Prettify(s)
}

func (s GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomResidual) GoString() string {
	return s.String()
}

func (s *GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomResidual) SetXlabel(v float32) *GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomResidual {
	s.Xlabel = &v
	return s
}

func (s *GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomResidual) SetYlabel(v float32) *GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomResidual {
	s.Ylabel = &v
	return s
}

func (s *GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomResidual) SetName(v string) *GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomResidual {
	s.Name = &v
	return s
}

func (s *GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomResidual) SetQuality(v int) *GetCustomIdentValuesResponseDataCustomIdentDataInfoCustomResidual {
	s.Quality = &v
	return s
}

type GetPidProjectReportOverviewRequest struct {
	PidProjectId *string `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s GetPidProjectReportOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPidProjectReportOverviewRequest) GoString() string {
	return s.String()
}

func (s *GetPidProjectReportOverviewRequest) SetPidProjectId(v string) *GetPidProjectReportOverviewRequest {
	s.PidProjectId = &v
	return s
}

func (s *GetPidProjectReportOverviewRequest) SetStartTime(v string) *GetPidProjectReportOverviewRequest {
	s.StartTime = &v
	return s
}

func (s *GetPidProjectReportOverviewRequest) SetEndTime(v string) *GetPidProjectReportOverviewRequest {
	s.EndTime = &v
	return s
}

type GetPidProjectReportOverviewResponse struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *GetPidProjectReportOverviewResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetPidProjectReportOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPidProjectReportOverviewResponse) GoString() string {
	return s.String()
}

func (s *GetPidProjectReportOverviewResponse) SetRequestId(v string) *GetPidProjectReportOverviewResponse {
	s.RequestId = &v
	return s
}

func (s *GetPidProjectReportOverviewResponse) SetCode(v string) *GetPidProjectReportOverviewResponse {
	s.Code = &v
	return s
}

func (s *GetPidProjectReportOverviewResponse) SetSuccess(v bool) *GetPidProjectReportOverviewResponse {
	s.Success = &v
	return s
}

func (s *GetPidProjectReportOverviewResponse) SetMessage(v string) *GetPidProjectReportOverviewResponse {
	s.Message = &v
	return s
}

func (s *GetPidProjectReportOverviewResponse) SetData(v *GetPidProjectReportOverviewResponseData) *GetPidProjectReportOverviewResponse {
	s.Data = v
	return s
}

type GetPidProjectReportOverviewResponseData struct {
	PerformMetrics        *float32                                                        `json:"PerformMetrics,omitempty" xml:"PerformMetrics,omitempty" require:"true"`
	OperationRate         *int64                                                          `json:"OperationRate,omitempty" xml:"OperationRate,omitempty" require:"true"`
	AllPerformMetricsList []*GetPidProjectReportOverviewResponseDataAllPerformMetricsList `json:"AllPerformMetricsList,omitempty" xml:"AllPerformMetricsList,omitempty" require:"true" type:"Repeated"`
	KeyPerformMetricsList []*GetPidProjectReportOverviewResponseDataKeyPerformMetricsList `json:"KeyPerformMetricsList,omitempty" xml:"KeyPerformMetricsList,omitempty" require:"true" type:"Repeated"`
	AllOperationRateList  []*GetPidProjectReportOverviewResponseDataAllOperationRateList  `json:"AllOperationRateList,omitempty" xml:"AllOperationRateList,omitempty" require:"true" type:"Repeated"`
	KeyOperationRateList  []*GetPidProjectReportOverviewResponseDataKeyOperationRateList  `json:"KeyOperationRateList,omitempty" xml:"KeyOperationRateList,omitempty" require:"true" type:"Repeated"`
	LoopScoreList         []*GetPidProjectReportOverviewResponseDataLoopScoreList         `json:"LoopScoreList,omitempty" xml:"LoopScoreList,omitempty" require:"true" type:"Repeated"`
	LoopOperationRateList []*GetPidProjectReportOverviewResponseDataLoopOperationRateList `json:"LoopOperationRateList,omitempty" xml:"LoopOperationRateList,omitempty" require:"true" type:"Repeated"`
	Status                *GetPidProjectReportOverviewResponseDataStatus                  `json:"Status,omitempty" xml:"Status,omitempty" require:"true" type:"Struct"`
}

func (s GetPidProjectReportOverviewResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetPidProjectReportOverviewResponseData) GoString() string {
	return s.String()
}

func (s *GetPidProjectReportOverviewResponseData) SetPerformMetrics(v float32) *GetPidProjectReportOverviewResponseData {
	s.PerformMetrics = &v
	return s
}

func (s *GetPidProjectReportOverviewResponseData) SetOperationRate(v int64) *GetPidProjectReportOverviewResponseData {
	s.OperationRate = &v
	return s
}

func (s *GetPidProjectReportOverviewResponseData) SetAllPerformMetricsList(v []*GetPidProjectReportOverviewResponseDataAllPerformMetricsList) *GetPidProjectReportOverviewResponseData {
	s.AllPerformMetricsList = v
	return s
}

func (s *GetPidProjectReportOverviewResponseData) SetKeyPerformMetricsList(v []*GetPidProjectReportOverviewResponseDataKeyPerformMetricsList) *GetPidProjectReportOverviewResponseData {
	s.KeyPerformMetricsList = v
	return s
}

func (s *GetPidProjectReportOverviewResponseData) SetAllOperationRateList(v []*GetPidProjectReportOverviewResponseDataAllOperationRateList) *GetPidProjectReportOverviewResponseData {
	s.AllOperationRateList = v
	return s
}

func (s *GetPidProjectReportOverviewResponseData) SetKeyOperationRateList(v []*GetPidProjectReportOverviewResponseDataKeyOperationRateList) *GetPidProjectReportOverviewResponseData {
	s.KeyOperationRateList = v
	return s
}

func (s *GetPidProjectReportOverviewResponseData) SetLoopScoreList(v []*GetPidProjectReportOverviewResponseDataLoopScoreList) *GetPidProjectReportOverviewResponseData {
	s.LoopScoreList = v
	return s
}

func (s *GetPidProjectReportOverviewResponseData) SetLoopOperationRateList(v []*GetPidProjectReportOverviewResponseDataLoopOperationRateList) *GetPidProjectReportOverviewResponseData {
	s.LoopOperationRateList = v
	return s
}

func (s *GetPidProjectReportOverviewResponseData) SetStatus(v *GetPidProjectReportOverviewResponseDataStatus) *GetPidProjectReportOverviewResponseData {
	s.Status = v
	return s
}

type GetPidProjectReportOverviewResponseDataAllPerformMetricsList struct {
	Xlabel *string `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *string `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
}

func (s GetPidProjectReportOverviewResponseDataAllPerformMetricsList) String() string {
	return tea.Prettify(s)
}

func (s GetPidProjectReportOverviewResponseDataAllPerformMetricsList) GoString() string {
	return s.String()
}

func (s *GetPidProjectReportOverviewResponseDataAllPerformMetricsList) SetXlabel(v string) *GetPidProjectReportOverviewResponseDataAllPerformMetricsList {
	s.Xlabel = &v
	return s
}

func (s *GetPidProjectReportOverviewResponseDataAllPerformMetricsList) SetYlabel(v string) *GetPidProjectReportOverviewResponseDataAllPerformMetricsList {
	s.Ylabel = &v
	return s
}

type GetPidProjectReportOverviewResponseDataKeyPerformMetricsList struct {
	Xlabel *string `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *string `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
}

func (s GetPidProjectReportOverviewResponseDataKeyPerformMetricsList) String() string {
	return tea.Prettify(s)
}

func (s GetPidProjectReportOverviewResponseDataKeyPerformMetricsList) GoString() string {
	return s.String()
}

func (s *GetPidProjectReportOverviewResponseDataKeyPerformMetricsList) SetXlabel(v string) *GetPidProjectReportOverviewResponseDataKeyPerformMetricsList {
	s.Xlabel = &v
	return s
}

func (s *GetPidProjectReportOverviewResponseDataKeyPerformMetricsList) SetYlabel(v string) *GetPidProjectReportOverviewResponseDataKeyPerformMetricsList {
	s.Ylabel = &v
	return s
}

type GetPidProjectReportOverviewResponseDataAllOperationRateList struct {
	Xlabel *string `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *string `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
}

func (s GetPidProjectReportOverviewResponseDataAllOperationRateList) String() string {
	return tea.Prettify(s)
}

func (s GetPidProjectReportOverviewResponseDataAllOperationRateList) GoString() string {
	return s.String()
}

func (s *GetPidProjectReportOverviewResponseDataAllOperationRateList) SetXlabel(v string) *GetPidProjectReportOverviewResponseDataAllOperationRateList {
	s.Xlabel = &v
	return s
}

func (s *GetPidProjectReportOverviewResponseDataAllOperationRateList) SetYlabel(v string) *GetPidProjectReportOverviewResponseDataAllOperationRateList {
	s.Ylabel = &v
	return s
}

type GetPidProjectReportOverviewResponseDataKeyOperationRateList struct {
	Xlabel *string `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *string `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
}

func (s GetPidProjectReportOverviewResponseDataKeyOperationRateList) String() string {
	return tea.Prettify(s)
}

func (s GetPidProjectReportOverviewResponseDataKeyOperationRateList) GoString() string {
	return s.String()
}

func (s *GetPidProjectReportOverviewResponseDataKeyOperationRateList) SetXlabel(v string) *GetPidProjectReportOverviewResponseDataKeyOperationRateList {
	s.Xlabel = &v
	return s
}

func (s *GetPidProjectReportOverviewResponseDataKeyOperationRateList) SetYlabel(v string) *GetPidProjectReportOverviewResponseDataKeyOperationRateList {
	s.Ylabel = &v
	return s
}

type GetPidProjectReportOverviewResponseDataLoopScoreList struct {
	Xlabel *string `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *string `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
}

func (s GetPidProjectReportOverviewResponseDataLoopScoreList) String() string {
	return tea.Prettify(s)
}

func (s GetPidProjectReportOverviewResponseDataLoopScoreList) GoString() string {
	return s.String()
}

func (s *GetPidProjectReportOverviewResponseDataLoopScoreList) SetXlabel(v string) *GetPidProjectReportOverviewResponseDataLoopScoreList {
	s.Xlabel = &v
	return s
}

func (s *GetPidProjectReportOverviewResponseDataLoopScoreList) SetYlabel(v string) *GetPidProjectReportOverviewResponseDataLoopScoreList {
	s.Ylabel = &v
	return s
}

type GetPidProjectReportOverviewResponseDataLoopOperationRateList struct {
	Xlabel *string `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *string `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
}

func (s GetPidProjectReportOverviewResponseDataLoopOperationRateList) String() string {
	return tea.Prettify(s)
}

func (s GetPidProjectReportOverviewResponseDataLoopOperationRateList) GoString() string {
	return s.String()
}

func (s *GetPidProjectReportOverviewResponseDataLoopOperationRateList) SetXlabel(v string) *GetPidProjectReportOverviewResponseDataLoopOperationRateList {
	s.Xlabel = &v
	return s
}

func (s *GetPidProjectReportOverviewResponseDataLoopOperationRateList) SetYlabel(v string) *GetPidProjectReportOverviewResponseDataLoopOperationRateList {
	s.Ylabel = &v
	return s
}

type GetPidProjectReportOverviewResponseDataStatus struct {
	All *GetPidProjectReportOverviewResponseDataStatusAll `json:"All,omitempty" xml:"All,omitempty" require:"true" type:"Struct"`
	Key *GetPidProjectReportOverviewResponseDataStatusKey `json:"Key,omitempty" xml:"Key,omitempty" require:"true" type:"Struct"`
}

func (s GetPidProjectReportOverviewResponseDataStatus) String() string {
	return tea.Prettify(s)
}

func (s GetPidProjectReportOverviewResponseDataStatus) GoString() string {
	return s.String()
}

func (s *GetPidProjectReportOverviewResponseDataStatus) SetAll(v *GetPidProjectReportOverviewResponseDataStatusAll) *GetPidProjectReportOverviewResponseDataStatus {
	s.All = v
	return s
}

func (s *GetPidProjectReportOverviewResponseDataStatus) SetKey(v *GetPidProjectReportOverviewResponseDataStatusKey) *GetPidProjectReportOverviewResponseDataStatus {
	s.Key = v
	return s
}

type GetPidProjectReportOverviewResponseDataStatusAll struct {
	Best     *int64 `json:"Best,omitempty" xml:"Best,omitempty" require:"true"`
	Good     *int64 `json:"Good,omitempty" xml:"Good,omitempty" require:"true"`
	Middle   *int64 `json:"Middle,omitempty" xml:"Middle,omitempty" require:"true"`
	Bad      *int64 `json:"Bad,omitempty" xml:"Bad,omitempty" require:"true"`
	OpenLoop *int64 `json:"OpenLoop,omitempty" xml:"OpenLoop,omitempty" require:"true"`
}

func (s GetPidProjectReportOverviewResponseDataStatusAll) String() string {
	return tea.Prettify(s)
}

func (s GetPidProjectReportOverviewResponseDataStatusAll) GoString() string {
	return s.String()
}

func (s *GetPidProjectReportOverviewResponseDataStatusAll) SetBest(v int64) *GetPidProjectReportOverviewResponseDataStatusAll {
	s.Best = &v
	return s
}

func (s *GetPidProjectReportOverviewResponseDataStatusAll) SetGood(v int64) *GetPidProjectReportOverviewResponseDataStatusAll {
	s.Good = &v
	return s
}

func (s *GetPidProjectReportOverviewResponseDataStatusAll) SetMiddle(v int64) *GetPidProjectReportOverviewResponseDataStatusAll {
	s.Middle = &v
	return s
}

func (s *GetPidProjectReportOverviewResponseDataStatusAll) SetBad(v int64) *GetPidProjectReportOverviewResponseDataStatusAll {
	s.Bad = &v
	return s
}

func (s *GetPidProjectReportOverviewResponseDataStatusAll) SetOpenLoop(v int64) *GetPidProjectReportOverviewResponseDataStatusAll {
	s.OpenLoop = &v
	return s
}

type GetPidProjectReportOverviewResponseDataStatusKey struct {
	Best     *int64 `json:"Best,omitempty" xml:"Best,omitempty" require:"true"`
	Good     *int64 `json:"Good,omitempty" xml:"Good,omitempty" require:"true"`
	Middle   *int64 `json:"Middle,omitempty" xml:"Middle,omitempty" require:"true"`
	Bad      *int64 `json:"Bad,omitempty" xml:"Bad,omitempty" require:"true"`
	OpenLoop *int64 `json:"OpenLoop,omitempty" xml:"OpenLoop,omitempty" require:"true"`
}

func (s GetPidProjectReportOverviewResponseDataStatusKey) String() string {
	return tea.Prettify(s)
}

func (s GetPidProjectReportOverviewResponseDataStatusKey) GoString() string {
	return s.String()
}

func (s *GetPidProjectReportOverviewResponseDataStatusKey) SetBest(v int64) *GetPidProjectReportOverviewResponseDataStatusKey {
	s.Best = &v
	return s
}

func (s *GetPidProjectReportOverviewResponseDataStatusKey) SetGood(v int64) *GetPidProjectReportOverviewResponseDataStatusKey {
	s.Good = &v
	return s
}

func (s *GetPidProjectReportOverviewResponseDataStatusKey) SetMiddle(v int64) *GetPidProjectReportOverviewResponseDataStatusKey {
	s.Middle = &v
	return s
}

func (s *GetPidProjectReportOverviewResponseDataStatusKey) SetBad(v int64) *GetPidProjectReportOverviewResponseDataStatusKey {
	s.Bad = &v
	return s
}

func (s *GetPidProjectReportOverviewResponseDataStatusKey) SetOpenLoop(v int64) *GetPidProjectReportOverviewResponseDataStatusKey {
	s.OpenLoop = &v
	return s
}

type GetLowModelPerformValuesRequest struct {
	PidLoopParameterId *string `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
	PidLoopId          *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
}

func (s GetLowModelPerformValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLowModelPerformValuesRequest) GoString() string {
	return s.String()
}

func (s *GetLowModelPerformValuesRequest) SetPidLoopParameterId(v string) *GetLowModelPerformValuesRequest {
	s.PidLoopParameterId = &v
	return s
}

func (s *GetLowModelPerformValuesRequest) SetPidLoopId(v string) *GetLowModelPerformValuesRequest {
	s.PidLoopId = &v
	return s
}

type GetLowModelPerformValuesResponse struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   map[string]interface{}                `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Data      *GetLowModelPerformValuesResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetLowModelPerformValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLowModelPerformValuesResponse) GoString() string {
	return s.String()
}

func (s *GetLowModelPerformValuesResponse) SetRequestId(v string) *GetLowModelPerformValuesResponse {
	s.RequestId = &v
	return s
}

func (s *GetLowModelPerformValuesResponse) SetMessage(v string) *GetLowModelPerformValuesResponse {
	s.Message = &v
	return s
}

func (s *GetLowModelPerformValuesResponse) SetCode(v string) *GetLowModelPerformValuesResponse {
	s.Code = &v
	return s
}

func (s *GetLowModelPerformValuesResponse) SetSuccess(v map[string]interface{}) *GetLowModelPerformValuesResponse {
	s.Success = v
	return s
}

func (s *GetLowModelPerformValuesResponse) SetData(v *GetLowModelPerformValuesResponseData) *GetLowModelPerformValuesResponse {
	s.Data = v
	return s
}

type GetLowModelPerformValuesResponseData struct {
	Status                  *bool                                                        `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	LowModelPerformDataInfo *GetLowModelPerformValuesResponseDataLowModelPerformDataInfo `json:"LowModelPerformDataInfo,omitempty" xml:"LowModelPerformDataInfo,omitempty" require:"true" type:"Struct"`
}

func (s GetLowModelPerformValuesResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetLowModelPerformValuesResponseData) GoString() string {
	return s.String()
}

func (s *GetLowModelPerformValuesResponseData) SetStatus(v bool) *GetLowModelPerformValuesResponseData {
	s.Status = &v
	return s
}

func (s *GetLowModelPerformValuesResponseData) SetLowModelPerformDataInfo(v *GetLowModelPerformValuesResponseDataLowModelPerformDataInfo) *GetLowModelPerformValuesResponseData {
	s.LowModelPerformDataInfo = v
	return s
}

type GetLowModelPerformValuesResponseDataLowModelPerformDataInfo struct {
	Confidence  *float32                                                                `json:"Confidence,omitempty" xml:"Confidence,omitempty" require:"true"`
	FitDegree   *float32                                                                `json:"FitDegree,omitempty" xml:"FitDegree,omitempty" require:"true"`
	ManualModel *GetLowModelPerformValuesResponseDataLowModelPerformDataInfoManualModel `json:"ManualModel,omitempty" xml:"ManualModel,omitempty" require:"true" type:"Struct"`
}

func (s GetLowModelPerformValuesResponseDataLowModelPerformDataInfo) String() string {
	return tea.Prettify(s)
}

func (s GetLowModelPerformValuesResponseDataLowModelPerformDataInfo) GoString() string {
	return s.String()
}

func (s *GetLowModelPerformValuesResponseDataLowModelPerformDataInfo) SetConfidence(v float32) *GetLowModelPerformValuesResponseDataLowModelPerformDataInfo {
	s.Confidence = &v
	return s
}

func (s *GetLowModelPerformValuesResponseDataLowModelPerformDataInfo) SetFitDegree(v float32) *GetLowModelPerformValuesResponseDataLowModelPerformDataInfo {
	s.FitDegree = &v
	return s
}

func (s *GetLowModelPerformValuesResponseDataLowModelPerformDataInfo) SetManualModel(v *GetLowModelPerformValuesResponseDataLowModelPerformDataInfoManualModel) *GetLowModelPerformValuesResponseDataLowModelPerformDataInfo {
	s.ManualModel = v
	return s
}

type GetLowModelPerformValuesResponseDataLowModelPerformDataInfoManualModel struct {
	K   *float32 `json:"K,omitempty" xml:"K,omitempty" require:"true"`
	Tau *float32 `json:"Tau,omitempty" xml:"Tau,omitempty" require:"true"`
	T1  *float32 `json:"T1,omitempty" xml:"T1,omitempty" require:"true"`
	T2  *float32 `json:"T2,omitempty" xml:"T2,omitempty" require:"true"`
}

func (s GetLowModelPerformValuesResponseDataLowModelPerformDataInfoManualModel) String() string {
	return tea.Prettify(s)
}

func (s GetLowModelPerformValuesResponseDataLowModelPerformDataInfoManualModel) GoString() string {
	return s.String()
}

func (s *GetLowModelPerformValuesResponseDataLowModelPerformDataInfoManualModel) SetK(v float32) *GetLowModelPerformValuesResponseDataLowModelPerformDataInfoManualModel {
	s.K = &v
	return s
}

func (s *GetLowModelPerformValuesResponseDataLowModelPerformDataInfoManualModel) SetTau(v float32) *GetLowModelPerformValuesResponseDataLowModelPerformDataInfoManualModel {
	s.Tau = &v
	return s
}

func (s *GetLowModelPerformValuesResponseDataLowModelPerformDataInfoManualModel) SetT1(v float32) *GetLowModelPerformValuesResponseDataLowModelPerformDataInfoManualModel {
	s.T1 = &v
	return s
}

func (s *GetLowModelPerformValuesResponseDataLowModelPerformDataInfoManualModel) SetT2(v float32) *GetLowModelPerformValuesResponseDataLowModelPerformDataInfoManualModel {
	s.T2 = &v
	return s
}

type ListPidLoopEvaluationsRequest struct {
	CurrentPage     *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PidProjectId    *string `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	Order           *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OrderByProperty *string `json:"OrderByProperty,omitempty" xml:"OrderByProperty,omitempty"`
	PidLoopName     *string `json:"PidLoopName,omitempty" xml:"PidLoopName,omitempty"`
	Grade           *string `json:"Grade,omitempty" xml:"Grade,omitempty"`
	Date            *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s ListPidLoopEvaluationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPidLoopEvaluationsRequest) GoString() string {
	return s.String()
}

func (s *ListPidLoopEvaluationsRequest) SetCurrentPage(v int) *ListPidLoopEvaluationsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListPidLoopEvaluationsRequest) SetPageSize(v int) *ListPidLoopEvaluationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPidLoopEvaluationsRequest) SetPidProjectId(v string) *ListPidLoopEvaluationsRequest {
	s.PidProjectId = &v
	return s
}

func (s *ListPidLoopEvaluationsRequest) SetOrder(v string) *ListPidLoopEvaluationsRequest {
	s.Order = &v
	return s
}

func (s *ListPidLoopEvaluationsRequest) SetOrderByProperty(v string) *ListPidLoopEvaluationsRequest {
	s.OrderByProperty = &v
	return s
}

func (s *ListPidLoopEvaluationsRequest) SetPidLoopName(v string) *ListPidLoopEvaluationsRequest {
	s.PidLoopName = &v
	return s
}

func (s *ListPidLoopEvaluationsRequest) SetGrade(v string) *ListPidLoopEvaluationsRequest {
	s.Grade = &v
	return s
}

func (s *ListPidLoopEvaluationsRequest) SetDate(v string) *ListPidLoopEvaluationsRequest {
	s.Date = &v
	return s
}

type ListPidLoopEvaluationsResponse struct {
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code        *string                               `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message     *string                               `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Success     *bool                                 `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	CurrentPage *int                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize    *int                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount  *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Data        []*ListPidLoopEvaluationsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s ListPidLoopEvaluationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPidLoopEvaluationsResponse) GoString() string {
	return s.String()
}

func (s *ListPidLoopEvaluationsResponse) SetRequestId(v string) *ListPidLoopEvaluationsResponse {
	s.RequestId = &v
	return s
}

func (s *ListPidLoopEvaluationsResponse) SetCode(v string) *ListPidLoopEvaluationsResponse {
	s.Code = &v
	return s
}

func (s *ListPidLoopEvaluationsResponse) SetMessage(v string) *ListPidLoopEvaluationsResponse {
	s.Message = &v
	return s
}

func (s *ListPidLoopEvaluationsResponse) SetSuccess(v bool) *ListPidLoopEvaluationsResponse {
	s.Success = &v
	return s
}

func (s *ListPidLoopEvaluationsResponse) SetCurrentPage(v int) *ListPidLoopEvaluationsResponse {
	s.CurrentPage = &v
	return s
}

func (s *ListPidLoopEvaluationsResponse) SetPageSize(v int) *ListPidLoopEvaluationsResponse {
	s.PageSize = &v
	return s
}

func (s *ListPidLoopEvaluationsResponse) SetTotalCount(v int64) *ListPidLoopEvaluationsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListPidLoopEvaluationsResponse) SetData(v []*ListPidLoopEvaluationsResponseData) *ListPidLoopEvaluationsResponse {
	s.Data = v
	return s
}

type ListPidLoopEvaluationsResponseData struct {
	AfterPidParameters  map[string]interface{} `json:"AfterPidParameters,omitempty" xml:"AfterPidParameters,omitempty" require:"true"`
	BeforePidParameters map[string]interface{} `json:"BeforePidParameters,omitempty" xml:"BeforePidParameters,omitempty" require:"true"`
	PidLoopRemark       *string                `json:"PidLoopRemark,omitempty" xml:"PidLoopRemark,omitempty" require:"true"`
	AssessTime          *int64                 `json:"AssessTime,omitempty" xml:"AssessTime,omitempty" require:"true"`
	Robust              *float32               `json:"Robust,omitempty" xml:"Robust,omitempty" require:"true"`
	ValidOperationRate  *float32               `json:"ValidOperationRate,omitempty" xml:"ValidOperationRate,omitempty" require:"true"`
	OperationRate       *float32               `json:"OperationRate,omitempty" xml:"OperationRate,omitempty" require:"true"`
	PerformMetrics      *float32               `json:"PerformMetrics,omitempty" xml:"PerformMetrics,omitempty" require:"true"`
	MultipleScore       *float32               `json:"MultipleScore,omitempty" xml:"MultipleScore,omitempty" require:"true"`
	Grade               *string                `json:"Grade,omitempty" xml:"Grade,omitempty" require:"true"`
	PidLoopName         *string                `json:"PidLoopName,omitempty" xml:"PidLoopName,omitempty" require:"true"`
	PidProjectId        *string                `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	PidLoopId           *string                `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
	PidLoopParameterId  *string                `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
}

func (s ListPidLoopEvaluationsResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListPidLoopEvaluationsResponseData) GoString() string {
	return s.String()
}

func (s *ListPidLoopEvaluationsResponseData) SetAfterPidParameters(v map[string]interface{}) *ListPidLoopEvaluationsResponseData {
	s.AfterPidParameters = v
	return s
}

func (s *ListPidLoopEvaluationsResponseData) SetBeforePidParameters(v map[string]interface{}) *ListPidLoopEvaluationsResponseData {
	s.BeforePidParameters = v
	return s
}

func (s *ListPidLoopEvaluationsResponseData) SetPidLoopRemark(v string) *ListPidLoopEvaluationsResponseData {
	s.PidLoopRemark = &v
	return s
}

func (s *ListPidLoopEvaluationsResponseData) SetAssessTime(v int64) *ListPidLoopEvaluationsResponseData {
	s.AssessTime = &v
	return s
}

func (s *ListPidLoopEvaluationsResponseData) SetRobust(v float32) *ListPidLoopEvaluationsResponseData {
	s.Robust = &v
	return s
}

func (s *ListPidLoopEvaluationsResponseData) SetValidOperationRate(v float32) *ListPidLoopEvaluationsResponseData {
	s.ValidOperationRate = &v
	return s
}

func (s *ListPidLoopEvaluationsResponseData) SetOperationRate(v float32) *ListPidLoopEvaluationsResponseData {
	s.OperationRate = &v
	return s
}

func (s *ListPidLoopEvaluationsResponseData) SetPerformMetrics(v float32) *ListPidLoopEvaluationsResponseData {
	s.PerformMetrics = &v
	return s
}

func (s *ListPidLoopEvaluationsResponseData) SetMultipleScore(v float32) *ListPidLoopEvaluationsResponseData {
	s.MultipleScore = &v
	return s
}

func (s *ListPidLoopEvaluationsResponseData) SetGrade(v string) *ListPidLoopEvaluationsResponseData {
	s.Grade = &v
	return s
}

func (s *ListPidLoopEvaluationsResponseData) SetPidLoopName(v string) *ListPidLoopEvaluationsResponseData {
	s.PidLoopName = &v
	return s
}

func (s *ListPidLoopEvaluationsResponseData) SetPidProjectId(v string) *ListPidLoopEvaluationsResponseData {
	s.PidProjectId = &v
	return s
}

func (s *ListPidLoopEvaluationsResponseData) SetPidLoopId(v string) *ListPidLoopEvaluationsResponseData {
	s.PidLoopId = &v
	return s
}

func (s *ListPidLoopEvaluationsResponseData) SetPidLoopParameterId(v string) *ListPidLoopEvaluationsResponseData {
	s.PidLoopParameterId = &v
	return s
}

type ListLoopParameterTagValuesRequest struct {
	PidLoopParameterId *string `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
	DataStartTime      *string `json:"DataStartTime,omitempty" xml:"DataStartTime,omitempty"`
	DataEndTime        *string `json:"DataEndTime,omitempty" xml:"DataEndTime,omitempty"`
	PidLoopId          *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
}

func (s ListLoopParameterTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLoopParameterTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListLoopParameterTagValuesRequest) SetPidLoopParameterId(v string) *ListLoopParameterTagValuesRequest {
	s.PidLoopParameterId = &v
	return s
}

func (s *ListLoopParameterTagValuesRequest) SetDataStartTime(v string) *ListLoopParameterTagValuesRequest {
	s.DataStartTime = &v
	return s
}

func (s *ListLoopParameterTagValuesRequest) SetDataEndTime(v string) *ListLoopParameterTagValuesRequest {
	s.DataEndTime = &v
	return s
}

func (s *ListLoopParameterTagValuesRequest) SetPidLoopId(v string) *ListLoopParameterTagValuesRequest {
	s.PidLoopId = &v
	return s
}

type ListLoopParameterTagValuesResponse struct {
	RequestId     *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message       *string                                          `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code          *string                                          `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success       *bool                                            `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	TagsValuesRsp *ListLoopParameterTagValuesResponseTagsValuesRsp `json:"TagsValuesRsp,omitempty" xml:"TagsValuesRsp,omitempty" require:"true" type:"Struct"`
}

func (s ListLoopParameterTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLoopParameterTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListLoopParameterTagValuesResponse) SetRequestId(v string) *ListLoopParameterTagValuesResponse {
	s.RequestId = &v
	return s
}

func (s *ListLoopParameterTagValuesResponse) SetMessage(v string) *ListLoopParameterTagValuesResponse {
	s.Message = &v
	return s
}

func (s *ListLoopParameterTagValuesResponse) SetCode(v string) *ListLoopParameterTagValuesResponse {
	s.Code = &v
	return s
}

func (s *ListLoopParameterTagValuesResponse) SetSuccess(v bool) *ListLoopParameterTagValuesResponse {
	s.Success = &v
	return s
}

func (s *ListLoopParameterTagValuesResponse) SetTagsValuesRsp(v *ListLoopParameterTagValuesResponseTagsValuesRsp) *ListLoopParameterTagValuesResponse {
	s.TagsValuesRsp = v
	return s
}

type ListLoopParameterTagValuesResponseTagsValuesRsp struct {
	Status   *bool                                                    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	DataList *ListLoopParameterTagValuesResponseTagsValuesRspDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Struct"`
}

func (s ListLoopParameterTagValuesResponseTagsValuesRsp) String() string {
	return tea.Prettify(s)
}

func (s ListLoopParameterTagValuesResponseTagsValuesRsp) GoString() string {
	return s.String()
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRsp) SetStatus(v bool) *ListLoopParameterTagValuesResponseTagsValuesRsp {
	s.Status = &v
	return s
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRsp) SetDataList(v *ListLoopParameterTagValuesResponseTagsValuesRspDataList) *ListLoopParameterTagValuesResponseTagsValuesRsp {
	s.DataList = v
	return s
}

type ListLoopParameterTagValuesResponseTagsValuesRspDataList struct {
	Op  []*ListLoopParameterTagValuesResponseTagsValuesRspDataListOp  `json:"Op,omitempty" xml:"Op,omitempty" require:"true" type:"Repeated"`
	Op1 []*ListLoopParameterTagValuesResponseTagsValuesRspDataListOp1 `json:"Op1,omitempty" xml:"Op1,omitempty" require:"true" type:"Repeated"`
	Op2 []*ListLoopParameterTagValuesResponseTagsValuesRspDataListOp2 `json:"Op2,omitempty" xml:"Op2,omitempty" require:"true" type:"Repeated"`
	Sp  []*ListLoopParameterTagValuesResponseTagsValuesRspDataListSp  `json:"Sp,omitempty" xml:"Sp,omitempty" require:"true" type:"Repeated"`
	Pv  []*ListLoopParameterTagValuesResponseTagsValuesRspDataListPv  `json:"Pv,omitempty" xml:"Pv,omitempty" require:"true" type:"Repeated"`
}

func (s ListLoopParameterTagValuesResponseTagsValuesRspDataList) String() string {
	return tea.Prettify(s)
}

func (s ListLoopParameterTagValuesResponseTagsValuesRspDataList) GoString() string {
	return s.String()
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataList) SetOp(v []*ListLoopParameterTagValuesResponseTagsValuesRspDataListOp) *ListLoopParameterTagValuesResponseTagsValuesRspDataList {
	s.Op = v
	return s
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataList) SetOp1(v []*ListLoopParameterTagValuesResponseTagsValuesRspDataListOp1) *ListLoopParameterTagValuesResponseTagsValuesRspDataList {
	s.Op1 = v
	return s
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataList) SetOp2(v []*ListLoopParameterTagValuesResponseTagsValuesRspDataListOp2) *ListLoopParameterTagValuesResponseTagsValuesRspDataList {
	s.Op2 = v
	return s
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataList) SetSp(v []*ListLoopParameterTagValuesResponseTagsValuesRspDataListSp) *ListLoopParameterTagValuesResponseTagsValuesRspDataList {
	s.Sp = v
	return s
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataList) SetPv(v []*ListLoopParameterTagValuesResponseTagsValuesRspDataListPv) *ListLoopParameterTagValuesResponseTagsValuesRspDataList {
	s.Pv = v
	return s
}

type ListLoopParameterTagValuesResponseTagsValuesRspDataListOp struct {
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s ListLoopParameterTagValuesResponseTagsValuesRspDataListOp) String() string {
	return tea.Prettify(s)
}

func (s ListLoopParameterTagValuesResponseTagsValuesRspDataListOp) GoString() string {
	return s.String()
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp) SetXlabel(v float32) *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp {
	s.Xlabel = &v
	return s
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp) SetYlabel(v float32) *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp {
	s.Ylabel = &v
	return s
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp) SetName(v string) *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp {
	s.Name = &v
	return s
}

type ListLoopParameterTagValuesResponseTagsValuesRspDataListOp1 struct {
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s ListLoopParameterTagValuesResponseTagsValuesRspDataListOp1) String() string {
	return tea.Prettify(s)
}

func (s ListLoopParameterTagValuesResponseTagsValuesRspDataListOp1) GoString() string {
	return s.String()
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp1) SetXlabel(v float32) *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp1 {
	s.Xlabel = &v
	return s
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp1) SetYlabel(v float32) *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp1 {
	s.Ylabel = &v
	return s
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp1) SetName(v string) *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp1 {
	s.Name = &v
	return s
}

type ListLoopParameterTagValuesResponseTagsValuesRspDataListOp2 struct {
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s ListLoopParameterTagValuesResponseTagsValuesRspDataListOp2) String() string {
	return tea.Prettify(s)
}

func (s ListLoopParameterTagValuesResponseTagsValuesRspDataListOp2) GoString() string {
	return s.String()
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp2) SetXlabel(v float32) *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp2 {
	s.Xlabel = &v
	return s
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp2) SetYlabel(v float32) *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp2 {
	s.Ylabel = &v
	return s
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp2) SetName(v string) *ListLoopParameterTagValuesResponseTagsValuesRspDataListOp2 {
	s.Name = &v
	return s
}

type ListLoopParameterTagValuesResponseTagsValuesRspDataListSp struct {
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s ListLoopParameterTagValuesResponseTagsValuesRspDataListSp) String() string {
	return tea.Prettify(s)
}

func (s ListLoopParameterTagValuesResponseTagsValuesRspDataListSp) GoString() string {
	return s.String()
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataListSp) SetXlabel(v float32) *ListLoopParameterTagValuesResponseTagsValuesRspDataListSp {
	s.Xlabel = &v
	return s
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataListSp) SetYlabel(v float32) *ListLoopParameterTagValuesResponseTagsValuesRspDataListSp {
	s.Ylabel = &v
	return s
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataListSp) SetName(v string) *ListLoopParameterTagValuesResponseTagsValuesRspDataListSp {
	s.Name = &v
	return s
}

type ListLoopParameterTagValuesResponseTagsValuesRspDataListPv struct {
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s ListLoopParameterTagValuesResponseTagsValuesRspDataListPv) String() string {
	return tea.Prettify(s)
}

func (s ListLoopParameterTagValuesResponseTagsValuesRspDataListPv) GoString() string {
	return s.String()
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataListPv) SetXlabel(v float32) *ListLoopParameterTagValuesResponseTagsValuesRspDataListPv {
	s.Xlabel = &v
	return s
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataListPv) SetYlabel(v float32) *ListLoopParameterTagValuesResponseTagsValuesRspDataListPv {
	s.Ylabel = &v
	return s
}

func (s *ListLoopParameterTagValuesResponseTagsValuesRspDataListPv) SetName(v string) *ListLoopParameterTagValuesResponseTagsValuesRspDataListPv {
	s.Name = &v
	return s
}

type GetPidLoopParameterRequest struct {
	PidLoopParameterId *string `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
	PidLoopId          *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
}

func (s GetPidLoopParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopParameterRequest) GoString() string {
	return s.String()
}

func (s *GetPidLoopParameterRequest) SetPidLoopParameterId(v string) *GetPidLoopParameterRequest {
	s.PidLoopParameterId = &v
	return s
}

func (s *GetPidLoopParameterRequest) SetPidLoopId(v string) *GetPidLoopParameterRequest {
	s.PidLoopId = &v
	return s
}

type GetPidLoopParameterResponse struct {
	RequestId        *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message          *string                                      `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code             *string                                      `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success          *bool                                        `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	PidLoopParameter *GetPidLoopParameterResponsePidLoopParameter `json:"PidLoopParameter,omitempty" xml:"PidLoopParameter,omitempty" require:"true" type:"Struct"`
}

func (s GetPidLoopParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopParameterResponse) GoString() string {
	return s.String()
}

func (s *GetPidLoopParameterResponse) SetRequestId(v string) *GetPidLoopParameterResponse {
	s.RequestId = &v
	return s
}

func (s *GetPidLoopParameterResponse) SetMessage(v string) *GetPidLoopParameterResponse {
	s.Message = &v
	return s
}

func (s *GetPidLoopParameterResponse) SetCode(v string) *GetPidLoopParameterResponse {
	s.Code = &v
	return s
}

func (s *GetPidLoopParameterResponse) SetSuccess(v bool) *GetPidLoopParameterResponse {
	s.Success = &v
	return s
}

func (s *GetPidLoopParameterResponse) SetPidLoopParameter(v *GetPidLoopParameterResponsePidLoopParameter) *GetPidLoopParameterResponse {
	s.PidLoopParameter = v
	return s
}

type GetPidLoopParameterResponsePidLoopParameter struct {
	PidLoopParameterId *string                                                   `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
	PidProjectId       *string                                                   `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	PidLoopId          *string                                                   `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
	PidIdentTaskDate   *string                                                   `json:"PidIdentTaskDate,omitempty" xml:"PidIdentTaskDate,omitempty" require:"true"`
	PidResetTaskDate   *string                                                   `json:"PidResetTaskDate,omitempty" xml:"PidResetTaskDate,omitempty" require:"true"`
	PidIdentParam      *GetPidLoopParameterResponsePidLoopParameterPidIdentParam `json:"PidIdentParam,omitempty" xml:"PidIdentParam,omitempty" require:"true" type:"Struct"`
	PIdResetParam      *GetPidLoopParameterResponsePidLoopParameterPIdResetParam `json:"PIdResetParam,omitempty" xml:"PIdResetParam,omitempty" require:"true" type:"Struct"`
}

func (s GetPidLoopParameterResponsePidLoopParameter) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopParameterResponsePidLoopParameter) GoString() string {
	return s.String()
}

func (s *GetPidLoopParameterResponsePidLoopParameter) SetPidLoopParameterId(v string) *GetPidLoopParameterResponsePidLoopParameter {
	s.PidLoopParameterId = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameter) SetPidProjectId(v string) *GetPidLoopParameterResponsePidLoopParameter {
	s.PidProjectId = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameter) SetPidLoopId(v string) *GetPidLoopParameterResponsePidLoopParameter {
	s.PidLoopId = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameter) SetPidIdentTaskDate(v string) *GetPidLoopParameterResponsePidLoopParameter {
	s.PidIdentTaskDate = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameter) SetPidResetTaskDate(v string) *GetPidLoopParameterResponsePidLoopParameter {
	s.PidResetTaskDate = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameter) SetPidIdentParam(v *GetPidLoopParameterResponsePidLoopParameterPidIdentParam) *GetPidLoopParameterResponsePidLoopParameter {
	s.PidIdentParam = v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameter) SetPIdResetParam(v *GetPidLoopParameterResponsePidLoopParameterPIdResetParam) *GetPidLoopParameterResponsePidLoopParameter {
	s.PIdResetParam = v
	return s
}

type GetPidLoopParameterResponsePidLoopParameterPidIdentParam struct {
	ModelType        *int                                                                 `json:"ModelType,omitempty" xml:"ModelType,omitempty" require:"true"`
	Integral         *bool                                                                `json:"Integral,omitempty" xml:"Integral,omitempty" require:"true"`
	TrendControl     *bool                                                                `json:"TrendControl,omitempty" xml:"TrendControl,omitempty" require:"true"`
	Delay            *int                                                                 `json:"Delay,omitempty" xml:"Delay,omitempty" require:"true"`
	Smooth           *int                                                                 `json:"Smooth,omitempty" xml:"Smooth,omitempty" require:"true"`
	Robust           *int                                                                 `json:"Robust,omitempty" xml:"Robust,omitempty" require:"true"`
	LimitSw          *bool                                                                `json:"LimitSw,omitempty" xml:"LimitSw,omitempty" require:"true"`
	ManualTrend      *bool                                                                `json:"ManualTrend,omitempty" xml:"ManualTrend,omitempty" require:"true"`
	StartTime        *string                                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime          *string                                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PidDataProcessId *int64                                                               `json:"PidDataProcessId,omitempty" xml:"PidDataProcessId,omitempty" require:"true"`
	LimitMax         *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMax    `json:"LimitMax,omitempty" xml:"LimitMax,omitempty" require:"true" type:"Struct"`
	LimitMin         *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMin    `json:"LimitMin,omitempty" xml:"LimitMin,omitempty" require:"true" type:"Struct"`
	ManualModel      *GetPidLoopParameterResponsePidLoopParameterPidIdentParamManualModel `json:"ManualModel,omitempty" xml:"ManualModel,omitempty" require:"true" type:"Struct"`
}

func (s GetPidLoopParameterResponsePidLoopParameterPidIdentParam) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopParameterResponsePidLoopParameterPidIdentParam) GoString() string {
	return s.String()
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParam) SetModelType(v int) *GetPidLoopParameterResponsePidLoopParameterPidIdentParam {
	s.ModelType = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParam) SetIntegral(v bool) *GetPidLoopParameterResponsePidLoopParameterPidIdentParam {
	s.Integral = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParam) SetTrendControl(v bool) *GetPidLoopParameterResponsePidLoopParameterPidIdentParam {
	s.TrendControl = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParam) SetDelay(v int) *GetPidLoopParameterResponsePidLoopParameterPidIdentParam {
	s.Delay = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParam) SetSmooth(v int) *GetPidLoopParameterResponsePidLoopParameterPidIdentParam {
	s.Smooth = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParam) SetRobust(v int) *GetPidLoopParameterResponsePidLoopParameterPidIdentParam {
	s.Robust = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParam) SetLimitSw(v bool) *GetPidLoopParameterResponsePidLoopParameterPidIdentParam {
	s.LimitSw = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParam) SetManualTrend(v bool) *GetPidLoopParameterResponsePidLoopParameterPidIdentParam {
	s.ManualTrend = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParam) SetStartTime(v string) *GetPidLoopParameterResponsePidLoopParameterPidIdentParam {
	s.StartTime = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParam) SetEndTime(v string) *GetPidLoopParameterResponsePidLoopParameterPidIdentParam {
	s.EndTime = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParam) SetPidDataProcessId(v int64) *GetPidLoopParameterResponsePidLoopParameterPidIdentParam {
	s.PidDataProcessId = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParam) SetLimitMax(v *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMax) *GetPidLoopParameterResponsePidLoopParameterPidIdentParam {
	s.LimitMax = v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParam) SetLimitMin(v *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMin) *GetPidLoopParameterResponsePidLoopParameterPidIdentParam {
	s.LimitMin = v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParam) SetManualModel(v *GetPidLoopParameterResponsePidLoopParameterPidIdentParamManualModel) *GetPidLoopParameterResponsePidLoopParameterPidIdentParam {
	s.ManualModel = v
	return s
}

type GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMax struct {
	K   *float32 `json:"K,omitempty" xml:"K,omitempty" require:"true"`
	Tau *float32 `json:"Tau,omitempty" xml:"Tau,omitempty" require:"true"`
	T1  *float32 `json:"T1,omitempty" xml:"T1,omitempty" require:"true"`
	T2  *float32 `json:"T2,omitempty" xml:"T2,omitempty" require:"true"`
}

func (s GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMax) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMax) GoString() string {
	return s.String()
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMax) SetK(v float32) *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMax {
	s.K = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMax) SetTau(v float32) *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMax {
	s.Tau = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMax) SetT1(v float32) *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMax {
	s.T1 = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMax) SetT2(v float32) *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMax {
	s.T2 = &v
	return s
}

type GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMin struct {
	K   *float32 `json:"K,omitempty" xml:"K,omitempty" require:"true"`
	Tau *float32 `json:"Tau,omitempty" xml:"Tau,omitempty" require:"true"`
	T1  *float32 `json:"T1,omitempty" xml:"T1,omitempty" require:"true"`
	T2  *float32 `json:"T2,omitempty" xml:"T2,omitempty" require:"true"`
}

func (s GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMin) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMin) GoString() string {
	return s.String()
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMin) SetK(v float32) *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMin {
	s.K = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMin) SetTau(v float32) *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMin {
	s.Tau = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMin) SetT1(v float32) *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMin {
	s.T1 = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMin) SetT2(v float32) *GetPidLoopParameterResponsePidLoopParameterPidIdentParamLimitMin {
	s.T2 = &v
	return s
}

type GetPidLoopParameterResponsePidLoopParameterPidIdentParamManualModel struct {
	K   *float32 `json:"K,omitempty" xml:"K,omitempty" require:"true"`
	Tau *float32 `json:"Tau,omitempty" xml:"Tau,omitempty" require:"true"`
	T1  *float32 `json:"T1,omitempty" xml:"T1,omitempty" require:"true"`
	T2  *float32 `json:"T2,omitempty" xml:"T2,omitempty" require:"true"`
}

func (s GetPidLoopParameterResponsePidLoopParameterPidIdentParamManualModel) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopParameterResponsePidLoopParameterPidIdentParamManualModel) GoString() string {
	return s.String()
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParamManualModel) SetK(v float32) *GetPidLoopParameterResponsePidLoopParameterPidIdentParamManualModel {
	s.K = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParamManualModel) SetTau(v float32) *GetPidLoopParameterResponsePidLoopParameterPidIdentParamManualModel {
	s.Tau = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParamManualModel) SetT1(v float32) *GetPidLoopParameterResponsePidLoopParameterPidIdentParamManualModel {
	s.T1 = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPidIdentParamManualModel) SetT2(v float32) *GetPidLoopParameterResponsePidLoopParameterPidIdentParamManualModel {
	s.T2 = &v
	return s
}

type GetPidLoopParameterResponsePidLoopParameterPIdResetParam struct {
	DcsType            *string                                                                     `json:"DcsType,omitempty" xml:"DcsType,omitempty" require:"true"`
	Model              *string                                                                     `json:"Model,omitempty" xml:"Model,omitempty" require:"true"`
	ModelType          *int                                                                        `json:"ModelType,omitempty" xml:"ModelType,omitempty" require:"true"`
	Integral           *bool                                                                       `json:"Integral,omitempty" xml:"Integral,omitempty" require:"true"`
	CtrlMode           *int                                                                        `json:"CtrlMode,omitempty" xml:"CtrlMode,omitempty" require:"true"`
	CtrlStuc           *int                                                                        `json:"CtrlStuc,omitempty" xml:"CtrlStuc,omitempty" require:"true"`
	Robust             *int                                                                        `json:"Robust,omitempty" xml:"Robust,omitempty" require:"true"`
	Dynamic            *float32                                                                    `json:"Dynamic,omitempty" xml:"Dynamic,omitempty" require:"true"`
	ManualTrend        *bool                                                                       `json:"ManualTrend,omitempty" xml:"ManualTrend,omitempty" require:"true"`
	MeasuredValueRange *GetPidLoopParameterResponsePidLoopParameterPIdResetParamMeasuredValueRange `json:"MeasuredValueRange,omitempty" xml:"MeasuredValueRange,omitempty" require:"true" type:"Struct"`
	ValvePositionRange *GetPidLoopParameterResponsePidLoopParameterPIdResetParamValvePositionRange `json:"ValvePositionRange,omitempty" xml:"ValvePositionRange,omitempty" require:"true" type:"Struct"`
	ManualCtrl         *GetPidLoopParameterResponsePidLoopParameterPIdResetParamManualCtrl         `json:"ManualCtrl,omitempty" xml:"ManualCtrl,omitempty" require:"true" type:"Struct"`
}

func (s GetPidLoopParameterResponsePidLoopParameterPIdResetParam) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopParameterResponsePidLoopParameterPIdResetParam) GoString() string {
	return s.String()
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParam) SetDcsType(v string) *GetPidLoopParameterResponsePidLoopParameterPIdResetParam {
	s.DcsType = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParam) SetModel(v string) *GetPidLoopParameterResponsePidLoopParameterPIdResetParam {
	s.Model = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParam) SetModelType(v int) *GetPidLoopParameterResponsePidLoopParameterPIdResetParam {
	s.ModelType = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParam) SetIntegral(v bool) *GetPidLoopParameterResponsePidLoopParameterPIdResetParam {
	s.Integral = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParam) SetCtrlMode(v int) *GetPidLoopParameterResponsePidLoopParameterPIdResetParam {
	s.CtrlMode = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParam) SetCtrlStuc(v int) *GetPidLoopParameterResponsePidLoopParameterPIdResetParam {
	s.CtrlStuc = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParam) SetRobust(v int) *GetPidLoopParameterResponsePidLoopParameterPIdResetParam {
	s.Robust = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParam) SetDynamic(v float32) *GetPidLoopParameterResponsePidLoopParameterPIdResetParam {
	s.Dynamic = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParam) SetManualTrend(v bool) *GetPidLoopParameterResponsePidLoopParameterPIdResetParam {
	s.ManualTrend = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParam) SetMeasuredValueRange(v *GetPidLoopParameterResponsePidLoopParameterPIdResetParamMeasuredValueRange) *GetPidLoopParameterResponsePidLoopParameterPIdResetParam {
	s.MeasuredValueRange = v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParam) SetValvePositionRange(v *GetPidLoopParameterResponsePidLoopParameterPIdResetParamValvePositionRange) *GetPidLoopParameterResponsePidLoopParameterPIdResetParam {
	s.ValvePositionRange = v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParam) SetManualCtrl(v *GetPidLoopParameterResponsePidLoopParameterPIdResetParamManualCtrl) *GetPidLoopParameterResponsePidLoopParameterPIdResetParam {
	s.ManualCtrl = v
	return s
}

type GetPidLoopParameterResponsePidLoopParameterPIdResetParamMeasuredValueRange struct {
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty" require:"true"`
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s GetPidLoopParameterResponsePidLoopParameterPIdResetParamMeasuredValueRange) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopParameterResponsePidLoopParameterPIdResetParamMeasuredValueRange) GoString() string {
	return s.String()
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParamMeasuredValueRange) SetMin(v float32) *GetPidLoopParameterResponsePidLoopParameterPIdResetParamMeasuredValueRange {
	s.Min = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParamMeasuredValueRange) SetMax(v float32) *GetPidLoopParameterResponsePidLoopParameterPIdResetParamMeasuredValueRange {
	s.Max = &v
	return s
}

type GetPidLoopParameterResponsePidLoopParameterPIdResetParamValvePositionRange struct {
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty" require:"true"`
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s GetPidLoopParameterResponsePidLoopParameterPIdResetParamValvePositionRange) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopParameterResponsePidLoopParameterPIdResetParamValvePositionRange) GoString() string {
	return s.String()
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParamValvePositionRange) SetMin(v float32) *GetPidLoopParameterResponsePidLoopParameterPIdResetParamValvePositionRange {
	s.Min = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParamValvePositionRange) SetMax(v float32) *GetPidLoopParameterResponsePidLoopParameterPIdResetParamValvePositionRange {
	s.Max = &v
	return s
}

type GetPidLoopParameterResponsePidLoopParameterPIdResetParamManualCtrl struct {
	Kp *float32 `json:"Kp,omitempty" xml:"Kp,omitempty" require:"true"`
	Ti *float32 `json:"Ti,omitempty" xml:"Ti,omitempty" require:"true"`
	Td *float32 `json:"Td,omitempty" xml:"Td,omitempty" require:"true"`
}

func (s GetPidLoopParameterResponsePidLoopParameterPIdResetParamManualCtrl) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopParameterResponsePidLoopParameterPIdResetParamManualCtrl) GoString() string {
	return s.String()
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParamManualCtrl) SetKp(v float32) *GetPidLoopParameterResponsePidLoopParameterPIdResetParamManualCtrl {
	s.Kp = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParamManualCtrl) SetTi(v float32) *GetPidLoopParameterResponsePidLoopParameterPIdResetParamManualCtrl {
	s.Ti = &v
	return s
}

func (s *GetPidLoopParameterResponsePidLoopParameterPIdResetParamManualCtrl) SetTd(v float32) *GetPidLoopParameterResponsePidLoopParameterPIdResetParamManualCtrl {
	s.Td = &v
	return s
}

type GetDailyReportInfoRequest struct {
	PidLoopId  *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
	ReportDate *string `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
}

func (s GetDailyReportInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDailyReportInfoRequest) SetPidLoopId(v string) *GetDailyReportInfoRequest {
	s.PidLoopId = &v
	return s
}

func (s *GetDailyReportInfoRequest) SetReportDate(v string) *GetDailyReportInfoRequest {
	s.ReportDate = &v
	return s
}

type GetDailyReportInfoResponse struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	DayResultRsp *GetDailyReportInfoResponseDayResultRsp `json:"DayResultRsp,omitempty" xml:"DayResultRsp,omitempty" require:"true" type:"Struct"`
}

func (s GetDailyReportInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDailyReportInfoResponse) SetRequestId(v string) *GetDailyReportInfoResponse {
	s.RequestId = &v
	return s
}

func (s *GetDailyReportInfoResponse) SetMessage(v string) *GetDailyReportInfoResponse {
	s.Message = &v
	return s
}

func (s *GetDailyReportInfoResponse) SetCode(v string) *GetDailyReportInfoResponse {
	s.Code = &v
	return s
}

func (s *GetDailyReportInfoResponse) SetSuccess(v bool) *GetDailyReportInfoResponse {
	s.Success = &v
	return s
}

func (s *GetDailyReportInfoResponse) SetDayResultRsp(v *GetDailyReportInfoResponseDayResultRsp) *GetDailyReportInfoResponse {
	s.DayResultRsp = v
	return s
}

type GetDailyReportInfoResponseDayResultRsp struct {
	Status        *bool                                                `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	DayResultData *GetDailyReportInfoResponseDayResultRspDayResultData `json:"DayResultData,omitempty" xml:"DayResultData,omitempty" require:"true" type:"Struct"`
}

func (s GetDailyReportInfoResponseDayResultRsp) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportInfoResponseDayResultRsp) GoString() string {
	return s.String()
}

func (s *GetDailyReportInfoResponseDayResultRsp) SetStatus(v bool) *GetDailyReportInfoResponseDayResultRsp {
	s.Status = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRsp) SetDayResultData(v *GetDailyReportInfoResponseDayResultRspDayResultData) *GetDailyReportInfoResponseDayResultRsp {
	s.DayResultData = v
	return s
}

type GetDailyReportInfoResponseDayResultRspDayResultData struct {
	LoopName           *string `json:"LoopName,omitempty" xml:"LoopName,omitempty" require:"true"`
	ReportDate         *string `json:"ReportDate,omitempty" xml:"ReportDate,omitempty" require:"true"`
	MultipleScore      *string `json:"MultipleScore,omitempty" xml:"MultipleScore,omitempty" require:"true"`
	PerformMetrics     *string `json:"PerformMetrics,omitempty" xml:"PerformMetrics,omitempty" require:"true"`
	MultipleScoreIndex *string `json:"MultipleScoreIndex,omitempty" xml:"MultipleScoreIndex,omitempty" require:"true"`
	OperationRate      *string `json:"OperationRate,omitempty" xml:"OperationRate,omitempty" require:"true"`
	ValidOperationRate *string `json:"ValidOperationRate,omitempty" xml:"ValidOperationRate,omitempty" require:"true"`
	OscillationIndex   *string `json:"OscillationIndex,omitempty" xml:"OscillationIndex,omitempty" require:"true"`
	ValveStickIndex    *string `json:"ValveStickIndex,omitempty" xml:"ValveStickIndex,omitempty" require:"true"`
	PvMean             *string `json:"PvMean,omitempty" xml:"PvMean,omitempty" require:"true"`
	PvStd              *string `json:"PvStd,omitempty" xml:"PvStd,omitempty" require:"true"`
	PvMax              *string `json:"PvMax,omitempty" xml:"PvMax,omitempty" require:"true"`
	PvMin              *string `json:"PvMin,omitempty" xml:"PvMin,omitempty" require:"true"`
	ErrMean            *string `json:"ErrMean,omitempty" xml:"ErrMean,omitempty" require:"true"`
	ErrStd             *string `json:"ErrStd,omitempty" xml:"ErrStd,omitempty" require:"true"`
	ErrMax             *string `json:"ErrMax,omitempty" xml:"ErrMax,omitempty" require:"true"`
	ErrMin             *string `json:"ErrMin,omitempty" xml:"ErrMin,omitempty" require:"true"`
	OpMean             *string `json:"OpMean,omitempty" xml:"OpMean,omitempty" require:"true"`
	OpStd              *string `json:"OpStd,omitempty" xml:"OpStd,omitempty" require:"true"`
	OpMax              *string `json:"OpMax,omitempty" xml:"OpMax,omitempty" require:"true"`
	OpMin              *string `json:"OpMin,omitempty" xml:"OpMin,omitempty" require:"true"`
	LoopTrans          *string `json:"LoopTrans,omitempty" xml:"LoopTrans,omitempty" require:"true"`
	SpCross            *string `json:"SpCross,omitempty" xml:"SpCross,omitempty" require:"true"`
	OpTurn             *string `json:"OpTurn,omitempty" xml:"OpTurn,omitempty" require:"true"`
	OpSum              *string `json:"OpSum,omitempty" xml:"OpSum,omitempty" require:"true"`
	GoodRate           *string `json:"GoodRate,omitempty" xml:"GoodRate,omitempty" require:"true"`
	SatuRate           *string `json:"SatuRate,omitempty" xml:"SatuRate,omitempty" require:"true"`
}

func (s GetDailyReportInfoResponseDayResultRspDayResultData) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportInfoResponseDayResultRspDayResultData) GoString() string {
	return s.String()
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetLoopName(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.LoopName = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetReportDate(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.ReportDate = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetMultipleScore(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.MultipleScore = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetPerformMetrics(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.PerformMetrics = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetMultipleScoreIndex(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.MultipleScoreIndex = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetOperationRate(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.OperationRate = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetValidOperationRate(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.ValidOperationRate = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetOscillationIndex(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.OscillationIndex = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetValveStickIndex(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.ValveStickIndex = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetPvMean(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.PvMean = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetPvStd(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.PvStd = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetPvMax(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.PvMax = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetPvMin(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.PvMin = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetErrMean(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.ErrMean = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetErrStd(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.ErrStd = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetErrMax(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.ErrMax = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetErrMin(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.ErrMin = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetOpMean(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.OpMean = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetOpStd(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.OpStd = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetOpMax(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.OpMax = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetOpMin(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.OpMin = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetLoopTrans(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.LoopTrans = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetSpCross(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.SpCross = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetOpTurn(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.OpTurn = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetOpSum(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.OpSum = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetGoodRate(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.GoodRate = &v
	return s
}

func (s *GetDailyReportInfoResponseDayResultRspDayResultData) SetSatuRate(v string) *GetDailyReportInfoResponseDayResultRspDayResultData {
	s.SatuRate = &v
	return s
}

type GetPidLoopRequest struct {
	PidLoopId *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
}

func (s GetPidLoopRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopRequest) GoString() string {
	return s.String()
}

func (s *GetPidLoopRequest) SetPidLoopId(v string) *GetPidLoopRequest {
	s.PidLoopId = &v
	return s
}

type GetPidLoopResponse struct {
	RequestId         *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message           *string                              `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code              *string                              `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success           *bool                                `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	LoopConfiguration *GetPidLoopResponseLoopConfiguration `json:"LoopConfiguration,omitempty" xml:"LoopConfiguration,omitempty" require:"true" type:"Struct"`
}

func (s GetPidLoopResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponse) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponse) SetRequestId(v string) *GetPidLoopResponse {
	s.RequestId = &v
	return s
}

func (s *GetPidLoopResponse) SetMessage(v string) *GetPidLoopResponse {
	s.Message = &v
	return s
}

func (s *GetPidLoopResponse) SetCode(v string) *GetPidLoopResponse {
	s.Code = &v
	return s
}

func (s *GetPidLoopResponse) SetSuccess(v bool) *GetPidLoopResponse {
	s.Success = &v
	return s
}

func (s *GetPidLoopResponse) SetLoopConfiguration(v *GetPidLoopResponseLoopConfiguration) *GetPidLoopResponse {
	s.LoopConfiguration = v
	return s
}

type GetPidLoopResponseLoopConfiguration struct {
	BaseParam  *GetPidLoopResponseLoopConfigurationBaseParam  `json:"BaseParam,omitempty" xml:"BaseParam,omitempty" require:"true" type:"Struct"`
	IdentParam *GetPidLoopResponseLoopConfigurationIdentParam `json:"IdentParam,omitempty" xml:"IdentParam,omitempty" require:"true" type:"Struct"`
	ResetParam *GetPidLoopResponseLoopConfigurationResetParam `json:"ResetParam,omitempty" xml:"ResetParam,omitempty" require:"true" type:"Struct"`
}

func (s GetPidLoopResponseLoopConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfiguration) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfiguration) SetBaseParam(v *GetPidLoopResponseLoopConfigurationBaseParam) *GetPidLoopResponseLoopConfiguration {
	s.BaseParam = v
	return s
}

func (s *GetPidLoopResponseLoopConfiguration) SetIdentParam(v *GetPidLoopResponseLoopConfigurationIdentParam) *GetPidLoopResponseLoopConfiguration {
	s.IdentParam = v
	return s
}

func (s *GetPidLoopResponseLoopConfiguration) SetResetParam(v *GetPidLoopResponseLoopConfigurationResetParam) *GetPidLoopResponseLoopConfiguration {
	s.ResetParam = v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParam struct {
	OpenLoopTime      *int                                                   `json:"OpenLoopTime,omitempty" xml:"OpenLoopTime,omitempty" require:"true"`
	SampleTime        *int                                                   `json:"SampleTime,omitempty" xml:"SampleTime,omitempty" require:"true"`
	SuitCtrlTime      *int                                                   `json:"SuitCtrlTime,omitempty" xml:"SuitCtrlTime,omitempty" require:"true"`
	Pv                *string                                                `json:"Pv,omitempty" xml:"Pv,omitempty" require:"true"`
	Sp                *string                                                `json:"Sp,omitempty" xml:"Sp,omitempty" require:"true"`
	SplitRangeControl *bool                                                  `json:"SplitRangeControl,omitempty" xml:"SplitRangeControl,omitempty" require:"true"`
	Op                *string                                                `json:"Op,omitempty" xml:"Op,omitempty" require:"true"`
	Op1               *string                                                `json:"Op1,omitempty" xml:"Op1,omitempty" require:"true"`
	Op2               *string                                                `json:"Op2,omitempty" xml:"Op2,omitempty" require:"true"`
	ControlSwitch     *string                                                `json:"ControlSwitch,omitempty" xml:"ControlSwitch,omitempty" require:"true"`
	Mv                *string                                                `json:"Mv,omitempty" xml:"Mv,omitempty" require:"true"`
	ForwardController *bool                                                  `json:"ForwardController,omitempty" xml:"ForwardController,omitempty" require:"true"`
	ForwardVariable   *string                                                `json:"ForwardVariable,omitempty" xml:"ForwardVariable,omitempty" require:"true"`
	Integral          *bool                                                  `json:"Integral,omitempty" xml:"Integral,omitempty" require:"true"`
	PvRange           *GetPidLoopResponseLoopConfigurationBaseParamPvRange   `json:"PvRange,omitempty" xml:"PvRange,omitempty" require:"true" type:"Struct"`
	SpOperate         *GetPidLoopResponseLoopConfigurationBaseParamSpOperate `json:"SpOperate,omitempty" xml:"SpOperate,omitempty" require:"true" type:"Struct"`
	OpParam           *GetPidLoopResponseLoopConfigurationBaseParamOpParam   `json:"OpParam,omitempty" xml:"OpParam,omitempty" require:"true" type:"Struct"`
	Op1Param          *GetPidLoopResponseLoopConfigurationBaseParamOp1Param  `json:"Op1Param,omitempty" xml:"Op1Param,omitempty" require:"true" type:"Struct"`
	Op2Param          *GetPidLoopResponseLoopConfigurationBaseParamOp2Param  `json:"Op2Param,omitempty" xml:"Op2Param,omitempty" require:"true" type:"Struct"`
	Kp                *GetPidLoopResponseLoopConfigurationBaseParamKp        `json:"Kp,omitempty" xml:"Kp,omitempty" require:"true" type:"Struct"`
	Td                *GetPidLoopResponseLoopConfigurationBaseParamTd        `json:"Td,omitempty" xml:"Td,omitempty" require:"true" type:"Struct"`
	Ti                *GetPidLoopResponseLoopConfigurationBaseParamTi        `json:"Ti,omitempty" xml:"Ti,omitempty" require:"true" type:"Struct"`
	Kd                *GetPidLoopResponseLoopConfigurationBaseParamKd        `json:"Kd,omitempty" xml:"Kd,omitempty" require:"true" type:"Struct"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParam) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParam) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetOpenLoopTime(v int) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.OpenLoopTime = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetSampleTime(v int) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.SampleTime = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetSuitCtrlTime(v int) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.SuitCtrlTime = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetPv(v string) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.Pv = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetSp(v string) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.Sp = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetSplitRangeControl(v bool) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.SplitRangeControl = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetOp(v string) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.Op = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetOp1(v string) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.Op1 = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetOp2(v string) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.Op2 = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetControlSwitch(v string) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.ControlSwitch = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetMv(v string) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.Mv = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetForwardController(v bool) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.ForwardController = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetForwardVariable(v string) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.ForwardVariable = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetIntegral(v bool) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.Integral = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetPvRange(v *GetPidLoopResponseLoopConfigurationBaseParamPvRange) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.PvRange = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetSpOperate(v *GetPidLoopResponseLoopConfigurationBaseParamSpOperate) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.SpOperate = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetOpParam(v *GetPidLoopResponseLoopConfigurationBaseParamOpParam) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.OpParam = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetOp1Param(v *GetPidLoopResponseLoopConfigurationBaseParamOp1Param) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.Op1Param = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetOp2Param(v *GetPidLoopResponseLoopConfigurationBaseParamOp2Param) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.Op2Param = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetKp(v *GetPidLoopResponseLoopConfigurationBaseParamKp) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.Kp = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetTd(v *GetPidLoopResponseLoopConfigurationBaseParamTd) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.Td = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetTi(v *GetPidLoopResponseLoopConfigurationBaseParamTi) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.Ti = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParam) SetKd(v *GetPidLoopResponseLoopConfigurationBaseParamKd) *GetPidLoopResponseLoopConfigurationBaseParam {
	s.Kd = v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamPvRange struct {
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty" require:"true"`
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamPvRange) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamPvRange) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamPvRange) SetMin(v float32) *GetPidLoopResponseLoopConfigurationBaseParamPvRange {
	s.Min = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamPvRange) SetMax(v float32) *GetPidLoopResponseLoopConfigurationBaseParamPvRange {
	s.Max = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamSpOperate struct {
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty" require:"true"`
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamSpOperate) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamSpOperate) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamSpOperate) SetMin(v float32) *GetPidLoopResponseLoopConfigurationBaseParamSpOperate {
	s.Min = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamSpOperate) SetMax(v float32) *GetPidLoopResponseLoopConfigurationBaseParamSpOperate {
	s.Max = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamOpParam struct {
	Trend     *int                                                          `json:"Trend,omitempty" xml:"Trend,omitempty" require:"true"`
	OpScope   *GetPidLoopResponseLoopConfigurationBaseParamOpParamOpScope   `json:"OpScope,omitempty" xml:"OpScope,omitempty" require:"true" type:"Struct"`
	Range     *GetPidLoopResponseLoopConfigurationBaseParamOpParamRange     `json:"Range,omitempty" xml:"Range,omitempty" require:"true" type:"Struct"`
	Operate   *GetPidLoopResponseLoopConfigurationBaseParamOpParamOperate   `json:"Operate,omitempty" xml:"Operate,omitempty" require:"true" type:"Struct"`
	Increment *GetPidLoopResponseLoopConfigurationBaseParamOpParamIncrement `json:"Increment,omitempty" xml:"Increment,omitempty" require:"true" type:"Struct"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOpParam) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOpParam) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOpParam) SetTrend(v int) *GetPidLoopResponseLoopConfigurationBaseParamOpParam {
	s.Trend = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOpParam) SetOpScope(v *GetPidLoopResponseLoopConfigurationBaseParamOpParamOpScope) *GetPidLoopResponseLoopConfigurationBaseParamOpParam {
	s.OpScope = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOpParam) SetRange(v *GetPidLoopResponseLoopConfigurationBaseParamOpParamRange) *GetPidLoopResponseLoopConfigurationBaseParamOpParam {
	s.Range = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOpParam) SetOperate(v *GetPidLoopResponseLoopConfigurationBaseParamOpParamOperate) *GetPidLoopResponseLoopConfigurationBaseParamOpParam {
	s.Operate = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOpParam) SetIncrement(v *GetPidLoopResponseLoopConfigurationBaseParamOpParamIncrement) *GetPidLoopResponseLoopConfigurationBaseParamOpParam {
	s.Increment = v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamOpParamOpScope struct {
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty" require:"true"`
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOpParamOpScope) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOpParamOpScope) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOpParamOpScope) SetMin(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOpParamOpScope {
	s.Min = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOpParamOpScope) SetMax(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOpParamOpScope {
	s.Max = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamOpParamRange struct {
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty" require:"true"`
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOpParamRange) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOpParamRange) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOpParamRange) SetMin(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOpParamRange {
	s.Min = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOpParamRange) SetMax(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOpParamRange {
	s.Max = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamOpParamOperate struct {
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty" require:"true"`
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOpParamOperate) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOpParamOperate) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOpParamOperate) SetMin(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOpParamOperate {
	s.Min = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOpParamOperate) SetMax(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOpParamOperate {
	s.Max = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamOpParamIncrement struct {
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOpParamIncrement) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOpParamIncrement) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOpParamIncrement) SetMax(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOpParamIncrement {
	s.Max = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamOp1Param struct {
	Trend     *int                                                           `json:"Trend,omitempty" xml:"Trend,omitempty" require:"true"`
	OpScope   *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOpScope   `json:"OpScope,omitempty" xml:"OpScope,omitempty" require:"true" type:"Struct"`
	Range     *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamRange     `json:"Range,omitempty" xml:"Range,omitempty" require:"true" type:"Struct"`
	Operate   *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOperate   `json:"Operate,omitempty" xml:"Operate,omitempty" require:"true" type:"Struct"`
	Increment *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamIncrement `json:"Increment,omitempty" xml:"Increment,omitempty" require:"true" type:"Struct"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp1Param) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp1Param) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp1Param) SetTrend(v int) *GetPidLoopResponseLoopConfigurationBaseParamOp1Param {
	s.Trend = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp1Param) SetOpScope(v *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOpScope) *GetPidLoopResponseLoopConfigurationBaseParamOp1Param {
	s.OpScope = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp1Param) SetRange(v *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamRange) *GetPidLoopResponseLoopConfigurationBaseParamOp1Param {
	s.Range = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp1Param) SetOperate(v *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOperate) *GetPidLoopResponseLoopConfigurationBaseParamOp1Param {
	s.Operate = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp1Param) SetIncrement(v *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamIncrement) *GetPidLoopResponseLoopConfigurationBaseParamOp1Param {
	s.Increment = v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOpScope struct {
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty" require:"true"`
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOpScope) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOpScope) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOpScope) SetMin(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOpScope {
	s.Min = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOpScope) SetMax(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOpScope {
	s.Max = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamOp1ParamRange struct {
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty" require:"true"`
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp1ParamRange) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp1ParamRange) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamRange) SetMin(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamRange {
	s.Min = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamRange) SetMax(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamRange {
	s.Max = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOperate struct {
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty" require:"true"`
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOperate) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOperate) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOperate) SetMin(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOperate {
	s.Min = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOperate) SetMax(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamOperate {
	s.Max = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamOp1ParamIncrement struct {
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp1ParamIncrement) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp1ParamIncrement) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamIncrement) SetMax(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOp1ParamIncrement {
	s.Max = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamOp2Param struct {
	Trend     *int                                                           `json:"Trend,omitempty" xml:"Trend,omitempty" require:"true"`
	OpScope   *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOpScope   `json:"OpScope,omitempty" xml:"OpScope,omitempty" require:"true" type:"Struct"`
	Range     *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamRange     `json:"Range,omitempty" xml:"Range,omitempty" require:"true" type:"Struct"`
	Operate   *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOperate   `json:"Operate,omitempty" xml:"Operate,omitempty" require:"true" type:"Struct"`
	Increment *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamIncrement `json:"Increment,omitempty" xml:"Increment,omitempty" require:"true" type:"Struct"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp2Param) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp2Param) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp2Param) SetTrend(v int) *GetPidLoopResponseLoopConfigurationBaseParamOp2Param {
	s.Trend = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp2Param) SetOpScope(v *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOpScope) *GetPidLoopResponseLoopConfigurationBaseParamOp2Param {
	s.OpScope = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp2Param) SetRange(v *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamRange) *GetPidLoopResponseLoopConfigurationBaseParamOp2Param {
	s.Range = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp2Param) SetOperate(v *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOperate) *GetPidLoopResponseLoopConfigurationBaseParamOp2Param {
	s.Operate = v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp2Param) SetIncrement(v *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamIncrement) *GetPidLoopResponseLoopConfigurationBaseParamOp2Param {
	s.Increment = v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOpScope struct {
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty" require:"true"`
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOpScope) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOpScope) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOpScope) SetMin(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOpScope {
	s.Min = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOpScope) SetMax(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOpScope {
	s.Max = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamOp2ParamRange struct {
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty" require:"true"`
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp2ParamRange) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp2ParamRange) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamRange) SetMin(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamRange {
	s.Min = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamRange) SetMax(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamRange {
	s.Max = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOperate struct {
	Min *float32 `json:"Min,omitempty" xml:"Min,omitempty" require:"true"`
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOperate) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOperate) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOperate) SetMin(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOperate {
	s.Min = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOperate) SetMax(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamOperate {
	s.Max = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamOp2ParamIncrement struct {
	Max *float32 `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp2ParamIncrement) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamOp2ParamIncrement) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamIncrement) SetMax(v float32) *GetPidLoopResponseLoopConfigurationBaseParamOp2ParamIncrement {
	s.Max = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamKp struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamKp) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamKp) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamKp) SetTagKey(v string) *GetPidLoopResponseLoopConfigurationBaseParamKp {
	s.TagKey = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamKp) SetTagValue(v string) *GetPidLoopResponseLoopConfigurationBaseParamKp {
	s.TagValue = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamTd struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamTd) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamTd) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamTd) SetTagKey(v string) *GetPidLoopResponseLoopConfigurationBaseParamTd {
	s.TagKey = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamTd) SetTagValue(v string) *GetPidLoopResponseLoopConfigurationBaseParamTd {
	s.TagValue = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamTi struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamTi) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamTi) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamTi) SetTagKey(v string) *GetPidLoopResponseLoopConfigurationBaseParamTi {
	s.TagKey = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamTi) SetTagValue(v string) *GetPidLoopResponseLoopConfigurationBaseParamTi {
	s.TagValue = &v
	return s
}

type GetPidLoopResponseLoopConfigurationBaseParamKd struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationBaseParamKd) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationBaseParamKd) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamKd) SetTagKey(v string) *GetPidLoopResponseLoopConfigurationBaseParamKd {
	s.TagKey = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationBaseParamKd) SetTagValue(v string) *GetPidLoopResponseLoopConfigurationBaseParamKd {
	s.TagValue = &v
	return s
}

type GetPidLoopResponseLoopConfigurationIdentParam struct {
	ModelType *int `json:"ModelType,omitempty" xml:"ModelType,omitempty" require:"true"`
	Delay     *int `json:"Delay,omitempty" xml:"Delay,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationIdentParam) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationIdentParam) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationIdentParam) SetModelType(v int) *GetPidLoopResponseLoopConfigurationIdentParam {
	s.ModelType = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationIdentParam) SetDelay(v int) *GetPidLoopResponseLoopConfigurationIdentParam {
	s.Delay = &v
	return s
}

type GetPidLoopResponseLoopConfigurationResetParam struct {
	CtrlMode *int `json:"CtrlMode,omitempty" xml:"CtrlMode,omitempty" require:"true"`
	CtrlStuc *int `json:"CtrlStuc,omitempty" xml:"CtrlStuc,omitempty" require:"true"`
}

func (s GetPidLoopResponseLoopConfigurationResetParam) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopResponseLoopConfigurationResetParam) GoString() string {
	return s.String()
}

func (s *GetPidLoopResponseLoopConfigurationResetParam) SetCtrlMode(v int) *GetPidLoopResponseLoopConfigurationResetParam {
	s.CtrlMode = &v
	return s
}

func (s *GetPidLoopResponseLoopConfigurationResetParam) SetCtrlStuc(v int) *GetPidLoopResponseLoopConfigurationResetParam {
	s.CtrlStuc = &v
	return s
}

type CreatePidProjectRequest struct {
	PidProjectName    *string `json:"PidProjectName,omitempty" xml:"PidProjectName,omitempty" require:"true"`
	PidOrganisationId *string `json:"PidOrganisationId,omitempty" xml:"PidOrganisationId,omitempty" require:"true"`
	PidProjectDesc    *string `json:"PidProjectDesc,omitempty" xml:"PidProjectDesc,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty" require:"true"`
}

func (s CreatePidProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePidProjectRequest) GoString() string {
	return s.String()
}

func (s *CreatePidProjectRequest) SetPidProjectName(v string) *CreatePidProjectRequest {
	s.PidProjectName = &v
	return s
}

func (s *CreatePidProjectRequest) SetPidOrganisationId(v string) *CreatePidProjectRequest {
	s.PidOrganisationId = &v
	return s
}

func (s *CreatePidProjectRequest) SetPidProjectDesc(v string) *CreatePidProjectRequest {
	s.PidProjectDesc = &v
	return s
}

func (s *CreatePidProjectRequest) SetClientToken(v string) *CreatePidProjectRequest {
	s.ClientToken = &v
	return s
}

type CreatePidProjectResponse struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	PidProjectId *string `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
}

func (s CreatePidProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePidProjectResponse) GoString() string {
	return s.String()
}

func (s *CreatePidProjectResponse) SetRequestId(v string) *CreatePidProjectResponse {
	s.RequestId = &v
	return s
}

func (s *CreatePidProjectResponse) SetMessage(v string) *CreatePidProjectResponse {
	s.Message = &v
	return s
}

func (s *CreatePidProjectResponse) SetCode(v string) *CreatePidProjectResponse {
	s.Code = &v
	return s
}

func (s *CreatePidProjectResponse) SetSuccess(v bool) *CreatePidProjectResponse {
	s.Success = &v
	return s
}

func (s *CreatePidProjectResponse) SetPidProjectId(v string) *CreatePidProjectResponse {
	s.PidProjectId = &v
	return s
}

type ListIdentModelsRequest struct {
	PidLoopParameterId *string `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
	PidLoopId          *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
}

func (s ListIdentModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIdentModelsRequest) GoString() string {
	return s.String()
}

func (s *ListIdentModelsRequest) SetPidLoopParameterId(v string) *ListIdentModelsRequest {
	s.PidLoopParameterId = &v
	return s
}

func (s *ListIdentModelsRequest) SetPidLoopId(v string) *ListIdentModelsRequest {
	s.PidLoopId = &v
	return s
}

type ListIdentModelsResponse struct {
	RequestId         *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message           *string                                     `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code              *string                                     `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success           *bool                                       `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	PidIdentModelList []*ListIdentModelsResponsePidIdentModelList `json:"PidIdentModelList,omitempty" xml:"PidIdentModelList,omitempty" require:"true" type:"Repeated"`
}

func (s ListIdentModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIdentModelsResponse) GoString() string {
	return s.String()
}

func (s *ListIdentModelsResponse) SetRequestId(v string) *ListIdentModelsResponse {
	s.RequestId = &v
	return s
}

func (s *ListIdentModelsResponse) SetMessage(v string) *ListIdentModelsResponse {
	s.Message = &v
	return s
}

func (s *ListIdentModelsResponse) SetCode(v string) *ListIdentModelsResponse {
	s.Code = &v
	return s
}

func (s *ListIdentModelsResponse) SetSuccess(v bool) *ListIdentModelsResponse {
	s.Success = &v
	return s
}

func (s *ListIdentModelsResponse) SetPidIdentModelList(v []*ListIdentModelsResponsePidIdentModelList) *ListIdentModelsResponse {
	s.PidIdentModelList = v
	return s
}

type ListIdentModelsResponsePidIdentModelList struct {
	ModelId *int                                           `json:"ModelId,omitempty" xml:"ModelId,omitempty" require:"true"`
	Desc    *string                                        `json:"Desc,omitempty" xml:"Desc,omitempty" require:"true"`
	Model   *ListIdentModelsResponsePidIdentModelListModel `json:"Model,omitempty" xml:"Model,omitempty" require:"true" type:"Struct"`
}

func (s ListIdentModelsResponsePidIdentModelList) String() string {
	return tea.Prettify(s)
}

func (s ListIdentModelsResponsePidIdentModelList) GoString() string {
	return s.String()
}

func (s *ListIdentModelsResponsePidIdentModelList) SetModelId(v int) *ListIdentModelsResponsePidIdentModelList {
	s.ModelId = &v
	return s
}

func (s *ListIdentModelsResponsePidIdentModelList) SetDesc(v string) *ListIdentModelsResponsePidIdentModelList {
	s.Desc = &v
	return s
}

func (s *ListIdentModelsResponsePidIdentModelList) SetModel(v *ListIdentModelsResponsePidIdentModelListModel) *ListIdentModelsResponsePidIdentModelList {
	s.Model = v
	return s
}

type ListIdentModelsResponsePidIdentModelListModel struct {
	K   *float32 `json:"K,omitempty" xml:"K,omitempty" require:"true"`
	Tau *float32 `json:"Tau,omitempty" xml:"Tau,omitempty" require:"true"`
	T1  *float32 `json:"T1,omitempty" xml:"T1,omitempty" require:"true"`
	T2  *float32 `json:"T2,omitempty" xml:"T2,omitempty" require:"true"`
}

func (s ListIdentModelsResponsePidIdentModelListModel) String() string {
	return tea.Prettify(s)
}

func (s ListIdentModelsResponsePidIdentModelListModel) GoString() string {
	return s.String()
}

func (s *ListIdentModelsResponsePidIdentModelListModel) SetK(v float32) *ListIdentModelsResponsePidIdentModelListModel {
	s.K = &v
	return s
}

func (s *ListIdentModelsResponsePidIdentModelListModel) SetTau(v float32) *ListIdentModelsResponsePidIdentModelListModel {
	s.Tau = &v
	return s
}

func (s *ListIdentModelsResponsePidIdentModelListModel) SetT1(v float32) *ListIdentModelsResponsePidIdentModelListModel {
	s.T1 = &v
	return s
}

func (s *ListIdentModelsResponsePidIdentModelListModel) SetT2(v float32) *ListIdentModelsResponsePidIdentModelListModel {
	s.T2 = &v
	return s
}

type ListPidDataProcessRequest struct {
	PidLoopId          *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
	OnlyCompleteStatus *bool   `json:"OnlyCompleteStatus,omitempty" xml:"OnlyCompleteStatus,omitempty" require:"true"`
	PidDataProcessType *string `json:"PidDataProcessType,omitempty" xml:"PidDataProcessType,omitempty" require:"true"`
}

func (s ListPidDataProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPidDataProcessRequest) GoString() string {
	return s.String()
}

func (s *ListPidDataProcessRequest) SetPidLoopId(v string) *ListPidDataProcessRequest {
	s.PidLoopId = &v
	return s
}

func (s *ListPidDataProcessRequest) SetOnlyCompleteStatus(v bool) *ListPidDataProcessRequest {
	s.OnlyCompleteStatus = &v
	return s
}

func (s *ListPidDataProcessRequest) SetPidDataProcessType(v string) *ListPidDataProcessRequest {
	s.PidDataProcessType = &v
	return s
}

type ListPidDataProcessResponse struct {
	RequestId          *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message            *string                                         `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code               *string                                         `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success            *bool                                           `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	PidDataProcessList []*ListPidDataProcessResponsePidDataProcessList `json:"PidDataProcessList,omitempty" xml:"PidDataProcessList,omitempty" require:"true" type:"Repeated"`
}

func (s ListPidDataProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPidDataProcessResponse) GoString() string {
	return s.String()
}

func (s *ListPidDataProcessResponse) SetRequestId(v string) *ListPidDataProcessResponse {
	s.RequestId = &v
	return s
}

func (s *ListPidDataProcessResponse) SetMessage(v string) *ListPidDataProcessResponse {
	s.Message = &v
	return s
}

func (s *ListPidDataProcessResponse) SetCode(v string) *ListPidDataProcessResponse {
	s.Code = &v
	return s
}

func (s *ListPidDataProcessResponse) SetSuccess(v bool) *ListPidDataProcessResponse {
	s.Success = &v
	return s
}

func (s *ListPidDataProcessResponse) SetPidDataProcessList(v []*ListPidDataProcessResponsePidDataProcessList) *ListPidDataProcessResponse {
	s.PidDataProcessList = v
	return s
}

type ListPidDataProcessResponsePidDataProcessList struct {
	PidDataProcessId       *string `json:"PidDataProcessId,omitempty" xml:"PidDataProcessId,omitempty" require:"true"`
	StartTime              *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime                *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PidDataProcessStatus   *string `json:"PidDataProcessStatus,omitempty" xml:"PidDataProcessStatus,omitempty" require:"true"`
	PidDataProcessProgress *string `json:"PidDataProcessProgress,omitempty" xml:"PidDataProcessProgress,omitempty" require:"true"`
}

func (s ListPidDataProcessResponsePidDataProcessList) String() string {
	return tea.Prettify(s)
}

func (s ListPidDataProcessResponsePidDataProcessList) GoString() string {
	return s.String()
}

func (s *ListPidDataProcessResponsePidDataProcessList) SetPidDataProcessId(v string) *ListPidDataProcessResponsePidDataProcessList {
	s.PidDataProcessId = &v
	return s
}

func (s *ListPidDataProcessResponsePidDataProcessList) SetStartTime(v string) *ListPidDataProcessResponsePidDataProcessList {
	s.StartTime = &v
	return s
}

func (s *ListPidDataProcessResponsePidDataProcessList) SetEndTime(v string) *ListPidDataProcessResponsePidDataProcessList {
	s.EndTime = &v
	return s
}

func (s *ListPidDataProcessResponsePidDataProcessList) SetPidDataProcessStatus(v string) *ListPidDataProcessResponsePidDataProcessList {
	s.PidDataProcessStatus = &v
	return s
}

func (s *ListPidDataProcessResponsePidDataProcessList) SetPidDataProcessProgress(v string) *ListPidDataProcessResponsePidDataProcessList {
	s.PidDataProcessProgress = &v
	return s
}

type AddCustomIdentModelRequest struct {
	PidLoopParameterId *string  `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
	Key                *string  `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Value              *float32 `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
	PidLoopId          *string  `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
}

func (s AddCustomIdentModelRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCustomIdentModelRequest) GoString() string {
	return s.String()
}

func (s *AddCustomIdentModelRequest) SetPidLoopParameterId(v string) *AddCustomIdentModelRequest {
	s.PidLoopParameterId = &v
	return s
}

func (s *AddCustomIdentModelRequest) SetKey(v string) *AddCustomIdentModelRequest {
	s.Key = &v
	return s
}

func (s *AddCustomIdentModelRequest) SetValue(v float32) *AddCustomIdentModelRequest {
	s.Value = &v
	return s
}

func (s *AddCustomIdentModelRequest) SetPidLoopId(v string) *AddCustomIdentModelRequest {
	s.PidLoopId = &v
	return s
}

type AddCustomIdentModelResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s AddCustomIdentModelResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCustomIdentModelResponse) GoString() string {
	return s.String()
}

func (s *AddCustomIdentModelResponse) SetRequestId(v string) *AddCustomIdentModelResponse {
	s.RequestId = &v
	return s
}

func (s *AddCustomIdentModelResponse) SetMessage(v string) *AddCustomIdentModelResponse {
	s.Message = &v
	return s
}

func (s *AddCustomIdentModelResponse) SetCode(v string) *AddCustomIdentModelResponse {
	s.Code = &v
	return s
}

func (s *AddCustomIdentModelResponse) SetSuccess(v bool) *AddCustomIdentModelResponse {
	s.Success = &v
	return s
}

type GetLoopParameterStepRequest struct {
	PidLoopParameterId *string `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
	PidLoopId          *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
}

func (s GetLoopParameterStepRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoopParameterStepRequest) GoString() string {
	return s.String()
}

func (s *GetLoopParameterStepRequest) SetPidLoopParameterId(v string) *GetLoopParameterStepRequest {
	s.PidLoopParameterId = &v
	return s
}

func (s *GetLoopParameterStepRequest) SetPidLoopId(v string) *GetLoopParameterStepRequest {
	s.PidLoopId = &v
	return s
}

type GetLoopParameterStepResponse struct {
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message                 *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code                    *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success                 *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	PidLoopParamTurningStep *int    `json:"PidLoopParamTurningStep,omitempty" xml:"PidLoopParamTurningStep,omitempty" require:"true"`
}

func (s GetLoopParameterStepResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoopParameterStepResponse) GoString() string {
	return s.String()
}

func (s *GetLoopParameterStepResponse) SetRequestId(v string) *GetLoopParameterStepResponse {
	s.RequestId = &v
	return s
}

func (s *GetLoopParameterStepResponse) SetMessage(v string) *GetLoopParameterStepResponse {
	s.Message = &v
	return s
}

func (s *GetLoopParameterStepResponse) SetCode(v string) *GetLoopParameterStepResponse {
	s.Code = &v
	return s
}

func (s *GetLoopParameterStepResponse) SetSuccess(v bool) *GetLoopParameterStepResponse {
	s.Success = &v
	return s
}

func (s *GetLoopParameterStepResponse) SetPidLoopParamTurningStep(v int) *GetLoopParameterStepResponse {
	s.PidLoopParamTurningStep = &v
	return s
}

type ListPidLoopsRequest struct {
	PidProjectId *string `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	PidLoopName  *string `json:"PidLoopName,omitempty" xml:"PidLoopName,omitempty"`
	CurrentPage  *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize     *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPidLoopsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPidLoopsRequest) GoString() string {
	return s.String()
}

func (s *ListPidLoopsRequest) SetPidProjectId(v string) *ListPidLoopsRequest {
	s.PidProjectId = &v
	return s
}

func (s *ListPidLoopsRequest) SetPidLoopName(v string) *ListPidLoopsRequest {
	s.PidLoopName = &v
	return s
}

func (s *ListPidLoopsRequest) SetCurrentPage(v int) *ListPidLoopsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListPidLoopsRequest) SetPageSize(v int) *ListPidLoopsRequest {
	s.PageSize = &v
	return s
}

type ListPidLoopsResponse struct {
	RequestId   *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message     *string                            `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code        *string                            `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success     *bool                              `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	CurrentPage *int                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize    *int                               `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount  *int64                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PidLoopList []*ListPidLoopsResponsePidLoopList `json:"PidLoopList,omitempty" xml:"PidLoopList,omitempty" require:"true" type:"Repeated"`
}

func (s ListPidLoopsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPidLoopsResponse) GoString() string {
	return s.String()
}

func (s *ListPidLoopsResponse) SetRequestId(v string) *ListPidLoopsResponse {
	s.RequestId = &v
	return s
}

func (s *ListPidLoopsResponse) SetMessage(v string) *ListPidLoopsResponse {
	s.Message = &v
	return s
}

func (s *ListPidLoopsResponse) SetCode(v string) *ListPidLoopsResponse {
	s.Code = &v
	return s
}

func (s *ListPidLoopsResponse) SetSuccess(v bool) *ListPidLoopsResponse {
	s.Success = &v
	return s
}

func (s *ListPidLoopsResponse) SetCurrentPage(v int) *ListPidLoopsResponse {
	s.CurrentPage = &v
	return s
}

func (s *ListPidLoopsResponse) SetPageSize(v int) *ListPidLoopsResponse {
	s.PageSize = &v
	return s
}

func (s *ListPidLoopsResponse) SetTotalCount(v int64) *ListPidLoopsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListPidLoopsResponse) SetPidLoopList(v []*ListPidLoopsResponsePidLoopList) *ListPidLoopsResponse {
	s.PidLoopList = v
	return s
}

type ListPidLoopsResponsePidLoopList struct {
	PidLoopId              *string   `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
	PidLoopName            *string   `json:"PidLoopName,omitempty" xml:"PidLoopName,omitempty" require:"true"`
	Crucial                *bool     `json:"Crucial,omitempty" xml:"Crucial,omitempty" require:"true"`
	PidPublishStatus       *string   `json:"PidPublishStatus,omitempty" xml:"PidPublishStatus,omitempty" require:"true"`
	ParameterTunningId     *string   `json:"ParameterTunningId,omitempty" xml:"ParameterTunningId,omitempty" require:"true"`
	PidLoopDcsType         *string   `json:"PidLoopDcsType,omitempty" xml:"PidLoopDcsType,omitempty" require:"true"`
	PidLoopType            *string   `json:"PidLoopType,omitempty" xml:"PidLoopType,omitempty" require:"true"`
	PidLoopProgress        *string   `json:"PidLoopProgress,omitempty" xml:"PidLoopProgress,omitempty" require:"true"`
	PidLoopReportProgress  *string   `json:"PidLoopReportProgress,omitempty" xml:"PidLoopReportProgress,omitempty" require:"true"`
	PidLoopDataProgress    *string   `json:"PidLoopDataProgress,omitempty" xml:"PidLoopDataProgress,omitempty" require:"true"`
	ParameterTunningIdList []*string `json:"ParameterTunningIdList,omitempty" xml:"ParameterTunningIdList,omitempty" require:"true" type:"Repeated"`
}

func (s ListPidLoopsResponsePidLoopList) String() string {
	return tea.Prettify(s)
}

func (s ListPidLoopsResponsePidLoopList) GoString() string {
	return s.String()
}

func (s *ListPidLoopsResponsePidLoopList) SetPidLoopId(v string) *ListPidLoopsResponsePidLoopList {
	s.PidLoopId = &v
	return s
}

func (s *ListPidLoopsResponsePidLoopList) SetPidLoopName(v string) *ListPidLoopsResponsePidLoopList {
	s.PidLoopName = &v
	return s
}

func (s *ListPidLoopsResponsePidLoopList) SetCrucial(v bool) *ListPidLoopsResponsePidLoopList {
	s.Crucial = &v
	return s
}

func (s *ListPidLoopsResponsePidLoopList) SetPidPublishStatus(v string) *ListPidLoopsResponsePidLoopList {
	s.PidPublishStatus = &v
	return s
}

func (s *ListPidLoopsResponsePidLoopList) SetParameterTunningId(v string) *ListPidLoopsResponsePidLoopList {
	s.ParameterTunningId = &v
	return s
}

func (s *ListPidLoopsResponsePidLoopList) SetPidLoopDcsType(v string) *ListPidLoopsResponsePidLoopList {
	s.PidLoopDcsType = &v
	return s
}

func (s *ListPidLoopsResponsePidLoopList) SetPidLoopType(v string) *ListPidLoopsResponsePidLoopList {
	s.PidLoopType = &v
	return s
}

func (s *ListPidLoopsResponsePidLoopList) SetPidLoopProgress(v string) *ListPidLoopsResponsePidLoopList {
	s.PidLoopProgress = &v
	return s
}

func (s *ListPidLoopsResponsePidLoopList) SetPidLoopReportProgress(v string) *ListPidLoopsResponsePidLoopList {
	s.PidLoopReportProgress = &v
	return s
}

func (s *ListPidLoopsResponsePidLoopList) SetPidLoopDataProgress(v string) *ListPidLoopsResponsePidLoopList {
	s.PidLoopDataProgress = &v
	return s
}

func (s *ListPidLoopsResponsePidLoopList) SetParameterTunningIdList(v []*string) *ListPidLoopsResponsePidLoopList {
	s.ParameterTunningIdList = v
	return s
}

type MovePidOrganizationRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty" require:"true"`
	MoveType       *string `json:"MoveType,omitempty" xml:"MoveType,omitempty" require:"true"`
}

func (s MovePidOrganizationRequest) String() string {
	return tea.Prettify(s)
}

func (s MovePidOrganizationRequest) GoString() string {
	return s.String()
}

func (s *MovePidOrganizationRequest) SetOrganizationId(v string) *MovePidOrganizationRequest {
	s.OrganizationId = &v
	return s
}

func (s *MovePidOrganizationRequest) SetMoveType(v string) *MovePidOrganizationRequest {
	s.MoveType = &v
	return s
}

type MovePidOrganizationResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s MovePidOrganizationResponse) String() string {
	return tea.Prettify(s)
}

func (s MovePidOrganizationResponse) GoString() string {
	return s.String()
}

func (s *MovePidOrganizationResponse) SetRequestId(v string) *MovePidOrganizationResponse {
	s.RequestId = &v
	return s
}

func (s *MovePidOrganizationResponse) SetMessage(v string) *MovePidOrganizationResponse {
	s.Message = &v
	return s
}

func (s *MovePidOrganizationResponse) SetCode(v string) *MovePidOrganizationResponse {
	s.Code = &v
	return s
}

func (s *MovePidOrganizationResponse) SetSuccess(v bool) *MovePidOrganizationResponse {
	s.Success = &v
	return s
}

type UpdatePidLoopRequest struct {
	PidLoopId            *string                `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
	PidLoopName          *string                `json:"PidLoopName,omitempty" xml:"PidLoopName,omitempty" require:"true"`
	Crucial              *bool                  `json:"Crucial,omitempty" xml:"Crucial,omitempty" require:"true"`
	PidLoopDesc          *string                `json:"PidLoopDesc,omitempty" xml:"PidLoopDesc,omitempty"`
	PidProjectId         *string                `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	PidLoopDcsType       *string                `json:"PidLoopDcsType,omitempty" xml:"PidLoopDcsType,omitempty" require:"true"`
	PidLoopType          *string                `json:"PidLoopType,omitempty" xml:"PidLoopType,omitempty" require:"true"`
	PidLoopConfiguration map[string]interface{} `json:"PidLoopConfiguration,omitempty" xml:"PidLoopConfiguration,omitempty" require:"true"`
}

func (s UpdatePidLoopRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePidLoopRequest) GoString() string {
	return s.String()
}

func (s *UpdatePidLoopRequest) SetPidLoopId(v string) *UpdatePidLoopRequest {
	s.PidLoopId = &v
	return s
}

func (s *UpdatePidLoopRequest) SetPidLoopName(v string) *UpdatePidLoopRequest {
	s.PidLoopName = &v
	return s
}

func (s *UpdatePidLoopRequest) SetCrucial(v bool) *UpdatePidLoopRequest {
	s.Crucial = &v
	return s
}

func (s *UpdatePidLoopRequest) SetPidLoopDesc(v string) *UpdatePidLoopRequest {
	s.PidLoopDesc = &v
	return s
}

func (s *UpdatePidLoopRequest) SetPidProjectId(v string) *UpdatePidLoopRequest {
	s.PidProjectId = &v
	return s
}

func (s *UpdatePidLoopRequest) SetPidLoopDcsType(v string) *UpdatePidLoopRequest {
	s.PidLoopDcsType = &v
	return s
}

func (s *UpdatePidLoopRequest) SetPidLoopType(v string) *UpdatePidLoopRequest {
	s.PidLoopType = &v
	return s
}

func (s *UpdatePidLoopRequest) SetPidLoopConfiguration(v map[string]interface{}) *UpdatePidLoopRequest {
	s.PidLoopConfiguration = v
	return s
}

type UpdatePidLoopResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s UpdatePidLoopResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePidLoopResponse) GoString() string {
	return s.String()
}

func (s *UpdatePidLoopResponse) SetRequestId(v string) *UpdatePidLoopResponse {
	s.RequestId = &v
	return s
}

func (s *UpdatePidLoopResponse) SetMessage(v string) *UpdatePidLoopResponse {
	s.Message = &v
	return s
}

func (s *UpdatePidLoopResponse) SetCode(v string) *UpdatePidLoopResponse {
	s.Code = &v
	return s
}

func (s *UpdatePidLoopResponse) SetSuccess(v string) *UpdatePidLoopResponse {
	s.Success = &v
	return s
}

type AddPidLoopToScheduleRequest struct {
	PidLoopIdList map[string]interface{} `json:"PidLoopIdList,omitempty" xml:"PidLoopIdList,omitempty" require:"true"`
}

func (s AddPidLoopToScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s AddPidLoopToScheduleRequest) GoString() string {
	return s.String()
}

func (s *AddPidLoopToScheduleRequest) SetPidLoopIdList(v map[string]interface{}) *AddPidLoopToScheduleRequest {
	s.PidLoopIdList = v
	return s
}

type AddPidLoopToScheduleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s AddPidLoopToScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddPidLoopToScheduleResponse) GoString() string {
	return s.String()
}

func (s *AddPidLoopToScheduleResponse) SetRequestId(v string) *AddPidLoopToScheduleResponse {
	s.RequestId = &v
	return s
}

func (s *AddPidLoopToScheduleResponse) SetMessage(v string) *AddPidLoopToScheduleResponse {
	s.Message = &v
	return s
}

func (s *AddPidLoopToScheduleResponse) SetCode(v string) *AddPidLoopToScheduleResponse {
	s.Code = &v
	return s
}

func (s *AddPidLoopToScheduleResponse) SetSuccess(v string) *AddPidLoopToScheduleResponse {
	s.Success = &v
	return s
}

type GetParsingTagTaskRequest struct {
	PidProjectId *string `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
}

func (s GetParsingTagTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetParsingTagTaskRequest) GoString() string {
	return s.String()
}

func (s *GetParsingTagTaskRequest) SetPidProjectId(v string) *GetParsingTagTaskRequest {
	s.PidProjectId = &v
	return s
}

type GetParsingTagTaskResponse struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	PidTagTaskStatus *bool   `json:"PidTagTaskStatus,omitempty" xml:"PidTagTaskStatus,omitempty" require:"true"`
}

func (s GetParsingTagTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetParsingTagTaskResponse) GoString() string {
	return s.String()
}

func (s *GetParsingTagTaskResponse) SetRequestId(v string) *GetParsingTagTaskResponse {
	s.RequestId = &v
	return s
}

func (s *GetParsingTagTaskResponse) SetMessage(v string) *GetParsingTagTaskResponse {
	s.Message = &v
	return s
}

func (s *GetParsingTagTaskResponse) SetCode(v string) *GetParsingTagTaskResponse {
	s.Code = &v
	return s
}

func (s *GetParsingTagTaskResponse) SetSuccess(v bool) *GetParsingTagTaskResponse {
	s.Success = &v
	return s
}

func (s *GetParsingTagTaskResponse) SetPidTagTaskStatus(v bool) *GetParsingTagTaskResponse {
	s.PidTagTaskStatus = &v
	return s
}

type GetPidLoopLatestTaskStatusRequest struct {
	PidLoopParameterId *string `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
	Type               *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	PidLoopId          *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
}

func (s GetPidLoopLatestTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopLatestTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPidLoopLatestTaskStatusRequest) SetPidLoopParameterId(v string) *GetPidLoopLatestTaskStatusRequest {
	s.PidLoopParameterId = &v
	return s
}

func (s *GetPidLoopLatestTaskStatusRequest) SetType(v string) *GetPidLoopLatestTaskStatusRequest {
	s.Type = &v
	return s
}

func (s *GetPidLoopLatestTaskStatusRequest) SetPidLoopId(v string) *GetPidLoopLatestTaskStatusRequest {
	s.PidLoopId = &v
	return s
}

type GetPidLoopLatestTaskStatusResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Status    *bool   `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s GetPidLoopLatestTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPidLoopLatestTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPidLoopLatestTaskStatusResponse) SetRequestId(v string) *GetPidLoopLatestTaskStatusResponse {
	s.RequestId = &v
	return s
}

func (s *GetPidLoopLatestTaskStatusResponse) SetCode(v string) *GetPidLoopLatestTaskStatusResponse {
	s.Code = &v
	return s
}

func (s *GetPidLoopLatestTaskStatusResponse) SetSuccess(v bool) *GetPidLoopLatestTaskStatusResponse {
	s.Success = &v
	return s
}

func (s *GetPidLoopLatestTaskStatusResponse) SetMessage(v string) *GetPidLoopLatestTaskStatusResponse {
	s.Message = &v
	return s
}

func (s *GetPidLoopLatestTaskStatusResponse) SetStatus(v bool) *GetPidLoopLatestTaskStatusResponse {
	s.Status = &v
	return s
}

type DeletePidLoopRequest struct {
	PidLoopId *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty"`
}

func (s DeletePidLoopRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePidLoopRequest) GoString() string {
	return s.String()
}

func (s *DeletePidLoopRequest) SetPidLoopId(v string) *DeletePidLoopRequest {
	s.PidLoopId = &v
	return s
}

type DeletePidLoopResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s DeletePidLoopResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePidLoopResponse) GoString() string {
	return s.String()
}

func (s *DeletePidLoopResponse) SetRequestId(v string) *DeletePidLoopResponse {
	s.RequestId = &v
	return s
}

func (s *DeletePidLoopResponse) SetMessage(v string) *DeletePidLoopResponse {
	s.Message = &v
	return s
}

func (s *DeletePidLoopResponse) SetCode(v string) *DeletePidLoopResponse {
	s.Code = &v
	return s
}

func (s *DeletePidLoopResponse) SetSuccess(v bool) *DeletePidLoopResponse {
	s.Success = &v
	return s
}

type GetSummaryReportInfoRequest struct {
	PidLoopId *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s GetSummaryReportInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSummaryReportInfoRequest) SetPidLoopId(v string) *GetSummaryReportInfoRequest {
	s.PidLoopId = &v
	return s
}

func (s *GetSummaryReportInfoRequest) SetStartTime(v string) *GetSummaryReportInfoRequest {
	s.StartTime = &v
	return s
}

func (s *GetSummaryReportInfoRequest) SetEndTime(v string) *GetSummaryReportInfoRequest {
	s.EndTime = &v
	return s
}

type GetSummaryReportInfoResponse struct {
	Code             *string                                       `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message          *string                                       `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId        *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success          *bool                                         `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	SummaryResultRsp *GetSummaryReportInfoResponseSummaryResultRsp `json:"SummaryResultRsp,omitempty" xml:"SummaryResultRsp,omitempty" require:"true" type:"Struct"`
}

func (s GetSummaryReportInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSummaryReportInfoResponse) SetCode(v string) *GetSummaryReportInfoResponse {
	s.Code = &v
	return s
}

func (s *GetSummaryReportInfoResponse) SetMessage(v string) *GetSummaryReportInfoResponse {
	s.Message = &v
	return s
}

func (s *GetSummaryReportInfoResponse) SetRequestId(v string) *GetSummaryReportInfoResponse {
	s.RequestId = &v
	return s
}

func (s *GetSummaryReportInfoResponse) SetSuccess(v bool) *GetSummaryReportInfoResponse {
	s.Success = &v
	return s
}

func (s *GetSummaryReportInfoResponse) SetSummaryResultRsp(v *GetSummaryReportInfoResponseSummaryResultRsp) *GetSummaryReportInfoResponse {
	s.SummaryResultRsp = v
	return s
}

type GetSummaryReportInfoResponseSummaryResultRsp struct {
	Status            *bool                                                          `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	SummaryResultData *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData `json:"SummaryResultData,omitempty" xml:"SummaryResultData,omitempty" require:"true" type:"Struct"`
}

func (s GetSummaryReportInfoResponseSummaryResultRsp) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportInfoResponseSummaryResultRsp) GoString() string {
	return s.String()
}

func (s *GetSummaryReportInfoResponseSummaryResultRsp) SetStatus(v bool) *GetSummaryReportInfoResponseSummaryResultRsp {
	s.Status = &v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRsp) SetSummaryResultData(v *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData) *GetSummaryReportInfoResponseSummaryResultRsp {
	s.SummaryResultData = v
	return s
}

type GetSummaryReportInfoResponseSummaryResultRspSummaryResultData struct {
	LoopName              *string                                                                             `json:"LoopName,omitempty" xml:"LoopName,omitempty" require:"true"`
	MultipleScoreIndexRsp *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreIndexRsp `json:"MultipleScoreIndexRsp,omitempty" xml:"MultipleScoreIndexRsp,omitempty" require:"true" type:"Struct"`
	MultipleScoreRsp      *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreRsp      `json:"MultipleScoreRsp,omitempty" xml:"MultipleScoreRsp,omitempty" require:"true" type:"Struct"`
	OperationRate         *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOperationRate         `json:"OperationRate,omitempty" xml:"OperationRate,omitempty" require:"true" type:"Struct"`
	OscillationIndex      *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOscillationIndex      `json:"OscillationIndex,omitempty" xml:"OscillationIndex,omitempty" require:"true" type:"Struct"`
	PerformMetrics        *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataPerformMetrics        `json:"PerformMetrics,omitempty" xml:"PerformMetrics,omitempty" require:"true" type:"Struct"`
	ValidOperationRate    *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValidOperationRate    `json:"ValidOperationRate,omitempty" xml:"ValidOperationRate,omitempty" require:"true" type:"Struct"`
	ValveStickIndex       *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValveStickIndex       `json:"ValveStickIndex,omitempty" xml:"ValveStickIndex,omitempty" require:"true" type:"Struct"`
}

func (s GetSummaryReportInfoResponseSummaryResultRspSummaryResultData) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportInfoResponseSummaryResultRspSummaryResultData) GoString() string {
	return s.String()
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData) SetLoopName(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData {
	s.LoopName = &v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData) SetMultipleScoreIndexRsp(v *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreIndexRsp) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData {
	s.MultipleScoreIndexRsp = v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData) SetMultipleScoreRsp(v *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreRsp) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData {
	s.MultipleScoreRsp = v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData) SetOperationRate(v *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOperationRate) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData {
	s.OperationRate = v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData) SetOscillationIndex(v *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOscillationIndex) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData {
	s.OscillationIndex = v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData) SetPerformMetrics(v *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataPerformMetrics) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData {
	s.PerformMetrics = v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData) SetValidOperationRate(v *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValidOperationRate) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData {
	s.ValidOperationRate = v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData) SetValveStickIndex(v *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValveStickIndex) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultData {
	s.ValveStickIndex = v
	return s
}

type GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreIndexRsp struct {
	Average *string `json:"Average,omitempty" xml:"Average,omitempty" require:"true"`
	Best    *string `json:"Best,omitempty" xml:"Best,omitempty" require:"true"`
	Worst   *string `json:"Worst,omitempty" xml:"Worst,omitempty" require:"true"`
}

func (s GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreIndexRsp) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreIndexRsp) GoString() string {
	return s.String()
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreIndexRsp) SetAverage(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreIndexRsp {
	s.Average = &v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreIndexRsp) SetBest(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreIndexRsp {
	s.Best = &v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreIndexRsp) SetWorst(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreIndexRsp {
	s.Worst = &v
	return s
}

type GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreRsp struct {
	Average *string `json:"Average,omitempty" xml:"Average,omitempty" require:"true"`
	Best    *string `json:"Best,omitempty" xml:"Best,omitempty" require:"true"`
	Worst   *string `json:"Worst,omitempty" xml:"Worst,omitempty" require:"true"`
}

func (s GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreRsp) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreRsp) GoString() string {
	return s.String()
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreRsp) SetAverage(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreRsp {
	s.Average = &v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreRsp) SetBest(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreRsp {
	s.Best = &v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreRsp) SetWorst(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataMultipleScoreRsp {
	s.Worst = &v
	return s
}

type GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOperationRate struct {
	Average *string `json:"Average,omitempty" xml:"Average,omitempty" require:"true"`
	Best    *string `json:"Best,omitempty" xml:"Best,omitempty" require:"true"`
	Worst   *string `json:"Worst,omitempty" xml:"Worst,omitempty" require:"true"`
}

func (s GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOperationRate) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOperationRate) GoString() string {
	return s.String()
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOperationRate) SetAverage(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOperationRate {
	s.Average = &v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOperationRate) SetBest(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOperationRate {
	s.Best = &v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOperationRate) SetWorst(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOperationRate {
	s.Worst = &v
	return s
}

type GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOscillationIndex struct {
	Average *string `json:"Average,omitempty" xml:"Average,omitempty" require:"true"`
	Best    *string `json:"Best,omitempty" xml:"Best,omitempty" require:"true"`
	Worst   *string `json:"Worst,omitempty" xml:"Worst,omitempty" require:"true"`
}

func (s GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOscillationIndex) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOscillationIndex) GoString() string {
	return s.String()
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOscillationIndex) SetAverage(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOscillationIndex {
	s.Average = &v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOscillationIndex) SetBest(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOscillationIndex {
	s.Best = &v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOscillationIndex) SetWorst(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataOscillationIndex {
	s.Worst = &v
	return s
}

type GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataPerformMetrics struct {
	Average *string `json:"Average,omitempty" xml:"Average,omitempty" require:"true"`
	Best    *string `json:"Best,omitempty" xml:"Best,omitempty" require:"true"`
	Worst   *string `json:"Worst,omitempty" xml:"Worst,omitempty" require:"true"`
}

func (s GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataPerformMetrics) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataPerformMetrics) GoString() string {
	return s.String()
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataPerformMetrics) SetAverage(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataPerformMetrics {
	s.Average = &v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataPerformMetrics) SetBest(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataPerformMetrics {
	s.Best = &v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataPerformMetrics) SetWorst(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataPerformMetrics {
	s.Worst = &v
	return s
}

type GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValidOperationRate struct {
	Average *string `json:"Average,omitempty" xml:"Average,omitempty" require:"true"`
	Best    *string `json:"Best,omitempty" xml:"Best,omitempty" require:"true"`
	Worst   *string `json:"Worst,omitempty" xml:"Worst,omitempty" require:"true"`
}

func (s GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValidOperationRate) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValidOperationRate) GoString() string {
	return s.String()
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValidOperationRate) SetAverage(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValidOperationRate {
	s.Average = &v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValidOperationRate) SetBest(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValidOperationRate {
	s.Best = &v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValidOperationRate) SetWorst(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValidOperationRate {
	s.Worst = &v
	return s
}

type GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValveStickIndex struct {
	Average *string `json:"Average,omitempty" xml:"Average,omitempty" require:"true"`
	Best    *string `json:"Best,omitempty" xml:"Best,omitempty" require:"true"`
	Worst   *string `json:"Worst,omitempty" xml:"Worst,omitempty" require:"true"`
}

func (s GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValveStickIndex) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValveStickIndex) GoString() string {
	return s.String()
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValveStickIndex) SetAverage(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValveStickIndex {
	s.Average = &v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValveStickIndex) SetBest(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValveStickIndex {
	s.Best = &v
	return s
}

func (s *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValveStickIndex) SetWorst(v string) *GetSummaryReportInfoResponseSummaryResultRspSummaryResultDataValveStickIndex {
	s.Worst = &v
	return s
}

type GetSummaryReportDataTrendRequest struct {
	PidLoopId *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	TrendType *string `json:"TrendType,omitempty" xml:"TrendType,omitempty" require:"true"`
}

func (s GetSummaryReportDataTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportDataTrendRequest) GoString() string {
	return s.String()
}

func (s *GetSummaryReportDataTrendRequest) SetPidLoopId(v string) *GetSummaryReportDataTrendRequest {
	s.PidLoopId = &v
	return s
}

func (s *GetSummaryReportDataTrendRequest) SetStartTime(v string) *GetSummaryReportDataTrendRequest {
	s.StartTime = &v
	return s
}

func (s *GetSummaryReportDataTrendRequest) SetEndTime(v string) *GetSummaryReportDataTrendRequest {
	s.EndTime = &v
	return s
}

func (s *GetSummaryReportDataTrendRequest) SetTrendType(v string) *GetSummaryReportDataTrendRequest {
	s.TrendType = &v
	return s
}

type GetSummaryReportDataTrendResponse struct {
	Code             *string                                            `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message          *string                                            `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success          *bool                                              `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	SummaryTrendData *GetSummaryReportDataTrendResponseSummaryTrendData `json:"SummaryTrendData,omitempty" xml:"SummaryTrendData,omitempty" require:"true" type:"Struct"`
}

func (s GetSummaryReportDataTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportDataTrendResponse) GoString() string {
	return s.String()
}

func (s *GetSummaryReportDataTrendResponse) SetCode(v string) *GetSummaryReportDataTrendResponse {
	s.Code = &v
	return s
}

func (s *GetSummaryReportDataTrendResponse) SetMessage(v string) *GetSummaryReportDataTrendResponse {
	s.Message = &v
	return s
}

func (s *GetSummaryReportDataTrendResponse) SetRequestId(v string) *GetSummaryReportDataTrendResponse {
	s.RequestId = &v
	return s
}

func (s *GetSummaryReportDataTrendResponse) SetSuccess(v bool) *GetSummaryReportDataTrendResponse {
	s.Success = &v
	return s
}

func (s *GetSummaryReportDataTrendResponse) SetSummaryTrendData(v *GetSummaryReportDataTrendResponseSummaryTrendData) *GetSummaryReportDataTrendResponse {
	s.SummaryTrendData = v
	return s
}

type GetSummaryReportDataTrendResponseSummaryTrendData struct {
	Status *bool                                                  `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	Data   *GetSummaryReportDataTrendResponseSummaryTrendDataData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetSummaryReportDataTrendResponseSummaryTrendData) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportDataTrendResponseSummaryTrendData) GoString() string {
	return s.String()
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendData) SetStatus(v bool) *GetSummaryReportDataTrendResponseSummaryTrendData {
	s.Status = &v
	return s
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendData) SetData(v *GetSummaryReportDataTrendResponseSummaryTrendDataData) *GetSummaryReportDataTrendResponseSummaryTrendData {
	s.Data = v
	return s
}

type GetSummaryReportDataTrendResponseSummaryTrendDataData struct {
	Op2 []*GetSummaryReportDataTrendResponseSummaryTrendDataDataOp2 `json:"Op2,omitempty" xml:"Op2,omitempty" require:"true" type:"Repeated"`
	Op  []*GetSummaryReportDataTrendResponseSummaryTrendDataDataOp  `json:"Op,omitempty" xml:"Op,omitempty" require:"true" type:"Repeated"`
	Op1 []*GetSummaryReportDataTrendResponseSummaryTrendDataDataOp1 `json:"Op1,omitempty" xml:"Op1,omitempty" require:"true" type:"Repeated"`
	Pv  []*GetSummaryReportDataTrendResponseSummaryTrendDataDataPv  `json:"Pv,omitempty" xml:"Pv,omitempty" require:"true" type:"Repeated"`
	Sp  []*GetSummaryReportDataTrendResponseSummaryTrendDataDataSp  `json:"Sp,omitempty" xml:"Sp,omitempty" require:"true" type:"Repeated"`
}

func (s GetSummaryReportDataTrendResponseSummaryTrendDataData) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportDataTrendResponseSummaryTrendDataData) GoString() string {
	return s.String()
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataData) SetOp2(v []*GetSummaryReportDataTrendResponseSummaryTrendDataDataOp2) *GetSummaryReportDataTrendResponseSummaryTrendDataData {
	s.Op2 = v
	return s
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataData) SetOp(v []*GetSummaryReportDataTrendResponseSummaryTrendDataDataOp) *GetSummaryReportDataTrendResponseSummaryTrendDataData {
	s.Op = v
	return s
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataData) SetOp1(v []*GetSummaryReportDataTrendResponseSummaryTrendDataDataOp1) *GetSummaryReportDataTrendResponseSummaryTrendDataData {
	s.Op1 = v
	return s
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataData) SetPv(v []*GetSummaryReportDataTrendResponseSummaryTrendDataDataPv) *GetSummaryReportDataTrendResponseSummaryTrendDataData {
	s.Pv = v
	return s
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataData) SetSp(v []*GetSummaryReportDataTrendResponseSummaryTrendDataDataSp) *GetSummaryReportDataTrendResponseSummaryTrendDataData {
	s.Sp = v
	return s
}

type GetSummaryReportDataTrendResponseSummaryTrendDataDataOp2 struct {
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
}

func (s GetSummaryReportDataTrendResponseSummaryTrendDataDataOp2) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportDataTrendResponseSummaryTrendDataDataOp2) GoString() string {
	return s.String()
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp2) SetName(v string) *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp2 {
	s.Name = &v
	return s
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp2) SetXlabel(v float32) *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp2 {
	s.Xlabel = &v
	return s
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp2) SetYlabel(v float32) *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp2 {
	s.Ylabel = &v
	return s
}

type GetSummaryReportDataTrendResponseSummaryTrendDataDataOp struct {
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
}

func (s GetSummaryReportDataTrendResponseSummaryTrendDataDataOp) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportDataTrendResponseSummaryTrendDataDataOp) GoString() string {
	return s.String()
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp) SetName(v string) *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp {
	s.Name = &v
	return s
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp) SetXlabel(v float32) *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp {
	s.Xlabel = &v
	return s
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp) SetYlabel(v float32) *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp {
	s.Ylabel = &v
	return s
}

type GetSummaryReportDataTrendResponseSummaryTrendDataDataOp1 struct {
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
}

func (s GetSummaryReportDataTrendResponseSummaryTrendDataDataOp1) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportDataTrendResponseSummaryTrendDataDataOp1) GoString() string {
	return s.String()
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp1) SetName(v string) *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp1 {
	s.Name = &v
	return s
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp1) SetXlabel(v float32) *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp1 {
	s.Xlabel = &v
	return s
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp1) SetYlabel(v float32) *GetSummaryReportDataTrendResponseSummaryTrendDataDataOp1 {
	s.Ylabel = &v
	return s
}

type GetSummaryReportDataTrendResponseSummaryTrendDataDataPv struct {
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
}

func (s GetSummaryReportDataTrendResponseSummaryTrendDataDataPv) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportDataTrendResponseSummaryTrendDataDataPv) GoString() string {
	return s.String()
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataDataPv) SetName(v string) *GetSummaryReportDataTrendResponseSummaryTrendDataDataPv {
	s.Name = &v
	return s
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataDataPv) SetXlabel(v float32) *GetSummaryReportDataTrendResponseSummaryTrendDataDataPv {
	s.Xlabel = &v
	return s
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataDataPv) SetYlabel(v float32) *GetSummaryReportDataTrendResponseSummaryTrendDataDataPv {
	s.Ylabel = &v
	return s
}

type GetSummaryReportDataTrendResponseSummaryTrendDataDataSp struct {
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
}

func (s GetSummaryReportDataTrendResponseSummaryTrendDataDataSp) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportDataTrendResponseSummaryTrendDataDataSp) GoString() string {
	return s.String()
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataDataSp) SetName(v string) *GetSummaryReportDataTrendResponseSummaryTrendDataDataSp {
	s.Name = &v
	return s
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataDataSp) SetXlabel(v float32) *GetSummaryReportDataTrendResponseSummaryTrendDataDataSp {
	s.Xlabel = &v
	return s
}

func (s *GetSummaryReportDataTrendResponseSummaryTrendDataDataSp) SetYlabel(v float32) *GetSummaryReportDataTrendResponseSummaryTrendDataDataSp {
	s.Ylabel = &v
	return s
}

type GetSummaryReportChartRequest struct {
	PidLoopId *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s GetSummaryReportChartRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportChartRequest) GoString() string {
	return s.String()
}

func (s *GetSummaryReportChartRequest) SetPidLoopId(v string) *GetSummaryReportChartRequest {
	s.PidLoopId = &v
	return s
}

func (s *GetSummaryReportChartRequest) SetStartTime(v string) *GetSummaryReportChartRequest {
	s.StartTime = &v
	return s
}

func (s *GetSummaryReportChartRequest) SetEndTime(v string) *GetSummaryReportChartRequest {
	s.EndTime = &v
	return s
}

type GetSummaryReportChartResponse struct {
	Code                    *string                                               `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message                 *string                                               `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RequestId               *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success                 *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	SummaryLineChartDataRsp *GetSummaryReportChartResponseSummaryLineChartDataRsp `json:"SummaryLineChartDataRsp,omitempty" xml:"SummaryLineChartDataRsp,omitempty" require:"true" type:"Struct"`
}

func (s GetSummaryReportChartResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportChartResponse) GoString() string {
	return s.String()
}

func (s *GetSummaryReportChartResponse) SetCode(v string) *GetSummaryReportChartResponse {
	s.Code = &v
	return s
}

func (s *GetSummaryReportChartResponse) SetMessage(v string) *GetSummaryReportChartResponse {
	s.Message = &v
	return s
}

func (s *GetSummaryReportChartResponse) SetRequestId(v string) *GetSummaryReportChartResponse {
	s.RequestId = &v
	return s
}

func (s *GetSummaryReportChartResponse) SetSuccess(v bool) *GetSummaryReportChartResponse {
	s.Success = &v
	return s
}

func (s *GetSummaryReportChartResponse) SetSummaryLineChartDataRsp(v *GetSummaryReportChartResponseSummaryLineChartDataRsp) *GetSummaryReportChartResponse {
	s.SummaryLineChartDataRsp = v
	return s
}

type GetSummaryReportChartResponseSummaryLineChartDataRsp struct {
	Status               *bool                                                                     `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	SummaryLineChartData *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartData `json:"SummaryLineChartData,omitempty" xml:"SummaryLineChartData,omitempty" require:"true" type:"Struct"`
}

func (s GetSummaryReportChartResponseSummaryLineChartDataRsp) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportChartResponseSummaryLineChartDataRsp) GoString() string {
	return s.String()
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRsp) SetStatus(v bool) *GetSummaryReportChartResponseSummaryLineChartDataRsp {
	s.Status = &v
	return s
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRsp) SetSummaryLineChartData(v *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartData) *GetSummaryReportChartResponseSummaryLineChartDataRsp {
	s.SummaryLineChartData = v
	return s
}

type GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartData struct {
	MultipleScoreList      []*GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataMultipleScoreList      `json:"MultipleScoreList,omitempty" xml:"MultipleScoreList,omitempty" require:"true" type:"Repeated"`
	OperationRateList      []*GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataOperationRateList      `json:"OperationRateList,omitempty" xml:"OperationRateList,omitempty" require:"true" type:"Repeated"`
	PerformMetricsList     []*GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataPerformMetricsList     `json:"PerformMetricsList,omitempty" xml:"PerformMetricsList,omitempty" require:"true" type:"Repeated"`
	ValidOperationRateList []*GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataValidOperationRateList `json:"ValidOperationRateList,omitempty" xml:"ValidOperationRateList,omitempty" require:"true" type:"Repeated"`
}

func (s GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartData) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartData) GoString() string {
	return s.String()
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartData) SetMultipleScoreList(v []*GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataMultipleScoreList) *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartData {
	s.MultipleScoreList = v
	return s
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartData) SetOperationRateList(v []*GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataOperationRateList) *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartData {
	s.OperationRateList = v
	return s
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartData) SetPerformMetricsList(v []*GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataPerformMetricsList) *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartData {
	s.PerformMetricsList = v
	return s
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartData) SetValidOperationRateList(v []*GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataValidOperationRateList) *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartData {
	s.ValidOperationRateList = v
	return s
}

type GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataMultipleScoreList struct {
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
}

func (s GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataMultipleScoreList) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataMultipleScoreList) GoString() string {
	return s.String()
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataMultipleScoreList) SetName(v string) *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataMultipleScoreList {
	s.Name = &v
	return s
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataMultipleScoreList) SetXlabel(v float32) *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataMultipleScoreList {
	s.Xlabel = &v
	return s
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataMultipleScoreList) SetYlabel(v float32) *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataMultipleScoreList {
	s.Ylabel = &v
	return s
}

type GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataOperationRateList struct {
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
}

func (s GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataOperationRateList) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataOperationRateList) GoString() string {
	return s.String()
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataOperationRateList) SetName(v string) *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataOperationRateList {
	s.Name = &v
	return s
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataOperationRateList) SetXlabel(v float32) *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataOperationRateList {
	s.Xlabel = &v
	return s
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataOperationRateList) SetYlabel(v float32) *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataOperationRateList {
	s.Ylabel = &v
	return s
}

type GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataPerformMetricsList struct {
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
}

func (s GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataPerformMetricsList) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataPerformMetricsList) GoString() string {
	return s.String()
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataPerformMetricsList) SetName(v string) *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataPerformMetricsList {
	s.Name = &v
	return s
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataPerformMetricsList) SetXlabel(v float32) *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataPerformMetricsList {
	s.Xlabel = &v
	return s
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataPerformMetricsList) SetYlabel(v float32) *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataPerformMetricsList {
	s.Ylabel = &v
	return s
}

type GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataValidOperationRateList struct {
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
}

func (s GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataValidOperationRateList) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataValidOperationRateList) GoString() string {
	return s.String()
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataValidOperationRateList) SetName(v string) *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataValidOperationRateList {
	s.Name = &v
	return s
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataValidOperationRateList) SetXlabel(v float32) *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataValidOperationRateList {
	s.Xlabel = &v
	return s
}

func (s *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataValidOperationRateList) SetYlabel(v float32) *GetSummaryReportChartResponseSummaryLineChartDataRspSummaryLineChartDataValidOperationRateList {
	s.Ylabel = &v
	return s
}

type GetDailyReportDataTrendRequest struct {
	PidLoopId  *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
	ReportDate *string `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
}

func (s GetDailyReportDataTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportDataTrendRequest) GoString() string {
	return s.String()
}

func (s *GetDailyReportDataTrendRequest) SetPidLoopId(v string) *GetDailyReportDataTrendRequest {
	s.PidLoopId = &v
	return s
}

func (s *GetDailyReportDataTrendRequest) SetReportDate(v string) *GetDailyReportDataTrendRequest {
	s.ReportDate = &v
	return s
}

type GetDailyReportDataTrendResponse struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message       *string                                       `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code          *string                                       `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success       *bool                                         `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	TagsValuesRsp *GetDailyReportDataTrendResponseTagsValuesRsp `json:"TagsValuesRsp,omitempty" xml:"TagsValuesRsp,omitempty" require:"true" type:"Struct"`
}

func (s GetDailyReportDataTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportDataTrendResponse) GoString() string {
	return s.String()
}

func (s *GetDailyReportDataTrendResponse) SetRequestId(v string) *GetDailyReportDataTrendResponse {
	s.RequestId = &v
	return s
}

func (s *GetDailyReportDataTrendResponse) SetMessage(v string) *GetDailyReportDataTrendResponse {
	s.Message = &v
	return s
}

func (s *GetDailyReportDataTrendResponse) SetCode(v string) *GetDailyReportDataTrendResponse {
	s.Code = &v
	return s
}

func (s *GetDailyReportDataTrendResponse) SetSuccess(v bool) *GetDailyReportDataTrendResponse {
	s.Success = &v
	return s
}

func (s *GetDailyReportDataTrendResponse) SetTagsValuesRsp(v *GetDailyReportDataTrendResponseTagsValuesRsp) *GetDailyReportDataTrendResponse {
	s.TagsValuesRsp = v
	return s
}

type GetDailyReportDataTrendResponseTagsValuesRsp struct {
	Status *bool                                             `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	Data   *GetDailyReportDataTrendResponseTagsValuesRspData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetDailyReportDataTrendResponseTagsValuesRsp) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportDataTrendResponseTagsValuesRsp) GoString() string {
	return s.String()
}

func (s *GetDailyReportDataTrendResponseTagsValuesRsp) SetStatus(v bool) *GetDailyReportDataTrendResponseTagsValuesRsp {
	s.Status = &v
	return s
}

func (s *GetDailyReportDataTrendResponseTagsValuesRsp) SetData(v *GetDailyReportDataTrendResponseTagsValuesRspData) *GetDailyReportDataTrendResponseTagsValuesRsp {
	s.Data = v
	return s
}

type GetDailyReportDataTrendResponseTagsValuesRspData struct {
	Op  []*GetDailyReportDataTrendResponseTagsValuesRspDataOp  `json:"Op,omitempty" xml:"Op,omitempty" require:"true" type:"Repeated"`
	Op1 []*GetDailyReportDataTrendResponseTagsValuesRspDataOp1 `json:"Op1,omitempty" xml:"Op1,omitempty" require:"true" type:"Repeated"`
	Op2 []*GetDailyReportDataTrendResponseTagsValuesRspDataOp2 `json:"Op2,omitempty" xml:"Op2,omitempty" require:"true" type:"Repeated"`
	Sp  []*GetDailyReportDataTrendResponseTagsValuesRspDataSp  `json:"Sp,omitempty" xml:"Sp,omitempty" require:"true" type:"Repeated"`
	Pv  []*GetDailyReportDataTrendResponseTagsValuesRspDataPv  `json:"Pv,omitempty" xml:"Pv,omitempty" require:"true" type:"Repeated"`
}

func (s GetDailyReportDataTrendResponseTagsValuesRspData) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportDataTrendResponseTagsValuesRspData) GoString() string {
	return s.String()
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspData) SetOp(v []*GetDailyReportDataTrendResponseTagsValuesRspDataOp) *GetDailyReportDataTrendResponseTagsValuesRspData {
	s.Op = v
	return s
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspData) SetOp1(v []*GetDailyReportDataTrendResponseTagsValuesRspDataOp1) *GetDailyReportDataTrendResponseTagsValuesRspData {
	s.Op1 = v
	return s
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspData) SetOp2(v []*GetDailyReportDataTrendResponseTagsValuesRspDataOp2) *GetDailyReportDataTrendResponseTagsValuesRspData {
	s.Op2 = v
	return s
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspData) SetSp(v []*GetDailyReportDataTrendResponseTagsValuesRspDataSp) *GetDailyReportDataTrendResponseTagsValuesRspData {
	s.Sp = v
	return s
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspData) SetPv(v []*GetDailyReportDataTrendResponseTagsValuesRspDataPv) *GetDailyReportDataTrendResponseTagsValuesRspData {
	s.Pv = v
	return s
}

type GetDailyReportDataTrendResponseTagsValuesRspDataOp struct {
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s GetDailyReportDataTrendResponseTagsValuesRspDataOp) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportDataTrendResponseTagsValuesRspDataOp) GoString() string {
	return s.String()
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspDataOp) SetXlabel(v float32) *GetDailyReportDataTrendResponseTagsValuesRspDataOp {
	s.Xlabel = &v
	return s
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspDataOp) SetYlabel(v float32) *GetDailyReportDataTrendResponseTagsValuesRspDataOp {
	s.Ylabel = &v
	return s
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspDataOp) SetName(v string) *GetDailyReportDataTrendResponseTagsValuesRspDataOp {
	s.Name = &v
	return s
}

type GetDailyReportDataTrendResponseTagsValuesRspDataOp1 struct {
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s GetDailyReportDataTrendResponseTagsValuesRspDataOp1) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportDataTrendResponseTagsValuesRspDataOp1) GoString() string {
	return s.String()
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspDataOp1) SetXlabel(v float32) *GetDailyReportDataTrendResponseTagsValuesRspDataOp1 {
	s.Xlabel = &v
	return s
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspDataOp1) SetYlabel(v float32) *GetDailyReportDataTrendResponseTagsValuesRspDataOp1 {
	s.Ylabel = &v
	return s
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspDataOp1) SetName(v string) *GetDailyReportDataTrendResponseTagsValuesRspDataOp1 {
	s.Name = &v
	return s
}

type GetDailyReportDataTrendResponseTagsValuesRspDataOp2 struct {
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s GetDailyReportDataTrendResponseTagsValuesRspDataOp2) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportDataTrendResponseTagsValuesRspDataOp2) GoString() string {
	return s.String()
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspDataOp2) SetXlabel(v float32) *GetDailyReportDataTrendResponseTagsValuesRspDataOp2 {
	s.Xlabel = &v
	return s
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspDataOp2) SetYlabel(v float32) *GetDailyReportDataTrendResponseTagsValuesRspDataOp2 {
	s.Ylabel = &v
	return s
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspDataOp2) SetName(v string) *GetDailyReportDataTrendResponseTagsValuesRspDataOp2 {
	s.Name = &v
	return s
}

type GetDailyReportDataTrendResponseTagsValuesRspDataSp struct {
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s GetDailyReportDataTrendResponseTagsValuesRspDataSp) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportDataTrendResponseTagsValuesRspDataSp) GoString() string {
	return s.String()
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspDataSp) SetXlabel(v float32) *GetDailyReportDataTrendResponseTagsValuesRspDataSp {
	s.Xlabel = &v
	return s
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspDataSp) SetYlabel(v float32) *GetDailyReportDataTrendResponseTagsValuesRspDataSp {
	s.Ylabel = &v
	return s
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspDataSp) SetName(v string) *GetDailyReportDataTrendResponseTagsValuesRspDataSp {
	s.Name = &v
	return s
}

type GetDailyReportDataTrendResponseTagsValuesRspDataPv struct {
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s GetDailyReportDataTrendResponseTagsValuesRspDataPv) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportDataTrendResponseTagsValuesRspDataPv) GoString() string {
	return s.String()
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspDataPv) SetXlabel(v float32) *GetDailyReportDataTrendResponseTagsValuesRspDataPv {
	s.Xlabel = &v
	return s
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspDataPv) SetYlabel(v float32) *GetDailyReportDataTrendResponseTagsValuesRspDataPv {
	s.Ylabel = &v
	return s
}

func (s *GetDailyReportDataTrendResponseTagsValuesRspDataPv) SetName(v string) *GetDailyReportDataTrendResponseTagsValuesRspDataPv {
	s.Name = &v
	return s
}

type ListTagValueTrendRequest struct {
	PidProjectId *string                           `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	PidTag       []*ListTagValueTrendRequestPidTag `json:"PidTag,omitempty" xml:"PidTag,omitempty" require:"true" type:"Repeated"`
}

func (s ListTagValueTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValueTrendRequest) GoString() string {
	return s.String()
}

func (s *ListTagValueTrendRequest) SetPidProjectId(v string) *ListTagValueTrendRequest {
	s.PidProjectId = &v
	return s
}

func (s *ListTagValueTrendRequest) SetPidTag(v []*ListTagValueTrendRequestPidTag) *ListTagValueTrendRequest {
	s.PidTag = v
	return s
}

type ListTagValueTrendRequestPidTag struct {
	PidTagId *string `json:"PidTagId,omitempty" xml:"PidTagId,omitempty" require:"true"`
}

func (s ListTagValueTrendRequestPidTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagValueTrendRequestPidTag) GoString() string {
	return s.String()
}

func (s *ListTagValueTrendRequestPidTag) SetPidTagId(v string) *ListTagValueTrendRequestPidTag {
	s.PidTagId = &v
	return s
}

type ListTagValueTrendResponse struct {
	RequestId       *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message         *string                                     `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code            *string                                     `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success         *bool                                       `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	PidTagTrendList []*ListTagValueTrendResponsePidTagTrendList `json:"PidTagTrendList,omitempty" xml:"PidTagTrendList,omitempty" require:"true" type:"Repeated"`
}

func (s ListTagValueTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValueTrendResponse) GoString() string {
	return s.String()
}

func (s *ListTagValueTrendResponse) SetRequestId(v string) *ListTagValueTrendResponse {
	s.RequestId = &v
	return s
}

func (s *ListTagValueTrendResponse) SetMessage(v string) *ListTagValueTrendResponse {
	s.Message = &v
	return s
}

func (s *ListTagValueTrendResponse) SetCode(v string) *ListTagValueTrendResponse {
	s.Code = &v
	return s
}

func (s *ListTagValueTrendResponse) SetSuccess(v bool) *ListTagValueTrendResponse {
	s.Success = &v
	return s
}

func (s *ListTagValueTrendResponse) SetPidTagTrendList(v []*ListTagValueTrendResponsePidTagTrendList) *ListTagValueTrendResponse {
	s.PidTagTrendList = v
	return s
}

type ListTagValueTrendResponsePidTagTrendList struct {
	PidTagName      *string                                                    `json:"PidTagName,omitempty" xml:"PidTagName,omitempty" require:"true"`
	PidTagValueList []*ListTagValueTrendResponsePidTagTrendListPidTagValueList `json:"PidTagValueList,omitempty" xml:"PidTagValueList,omitempty" require:"true" type:"Repeated"`
}

func (s ListTagValueTrendResponsePidTagTrendList) String() string {
	return tea.Prettify(s)
}

func (s ListTagValueTrendResponsePidTagTrendList) GoString() string {
	return s.String()
}

func (s *ListTagValueTrendResponsePidTagTrendList) SetPidTagName(v string) *ListTagValueTrendResponsePidTagTrendList {
	s.PidTagName = &v
	return s
}

func (s *ListTagValueTrendResponsePidTagTrendList) SetPidTagValueList(v []*ListTagValueTrendResponsePidTagTrendListPidTagValueList) *ListTagValueTrendResponsePidTagTrendList {
	s.PidTagValueList = v
	return s
}

type ListTagValueTrendResponsePidTagTrendListPidTagValueList struct {
	Time  *int64  `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s ListTagValueTrendResponsePidTagTrendListPidTagValueList) String() string {
	return tea.Prettify(s)
}

func (s ListTagValueTrendResponsePidTagTrendListPidTagValueList) GoString() string {
	return s.String()
}

func (s *ListTagValueTrendResponsePidTagTrendListPidTagValueList) SetTime(v int64) *ListTagValueTrendResponsePidTagTrendListPidTagValueList {
	s.Time = &v
	return s
}

func (s *ListTagValueTrendResponsePidTagTrendListPidTagValueList) SetValue(v string) *ListTagValueTrendResponsePidTagTrendListPidTagValueList {
	s.Value = &v
	return s
}

type GetDailyReportChartRequest struct {
	PidLoopId  *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
	ReportDate *string `json:"ReportDate,omitempty" xml:"ReportDate,omitempty"`
}

func (s GetDailyReportChartRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportChartRequest) GoString() string {
	return s.String()
}

func (s *GetDailyReportChartRequest) SetPidLoopId(v string) *GetDailyReportChartRequest {
	s.PidLoopId = &v
	return s
}

func (s *GetDailyReportChartRequest) SetReportDate(v string) *GetDailyReportChartRequest {
	s.ReportDate = &v
	return s
}

type GetDailyReportChartResponse struct {
	RequestId           *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message             *string                                         `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code                *string                                         `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success             *bool                                           `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	DayLineChartDataRsp *GetDailyReportChartResponseDayLineChartDataRsp `json:"DayLineChartDataRsp,omitempty" xml:"DayLineChartDataRsp,omitempty" require:"true" type:"Struct"`
}

func (s GetDailyReportChartResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportChartResponse) GoString() string {
	return s.String()
}

func (s *GetDailyReportChartResponse) SetRequestId(v string) *GetDailyReportChartResponse {
	s.RequestId = &v
	return s
}

func (s *GetDailyReportChartResponse) SetMessage(v string) *GetDailyReportChartResponse {
	s.Message = &v
	return s
}

func (s *GetDailyReportChartResponse) SetCode(v string) *GetDailyReportChartResponse {
	s.Code = &v
	return s
}

func (s *GetDailyReportChartResponse) SetSuccess(v bool) *GetDailyReportChartResponse {
	s.Success = &v
	return s
}

func (s *GetDailyReportChartResponse) SetDayLineChartDataRsp(v *GetDailyReportChartResponseDayLineChartDataRsp) *GetDailyReportChartResponse {
	s.DayLineChartDataRsp = v
	return s
}

type GetDailyReportChartResponseDayLineChartDataRsp struct {
	Status           *bool                                                           `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	DayLineChartData *GetDailyReportChartResponseDayLineChartDataRspDayLineChartData `json:"DayLineChartData,omitempty" xml:"DayLineChartData,omitempty" require:"true" type:"Struct"`
}

func (s GetDailyReportChartResponseDayLineChartDataRsp) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportChartResponseDayLineChartDataRsp) GoString() string {
	return s.String()
}

func (s *GetDailyReportChartResponseDayLineChartDataRsp) SetStatus(v bool) *GetDailyReportChartResponseDayLineChartDataRsp {
	s.Status = &v
	return s
}

func (s *GetDailyReportChartResponseDayLineChartDataRsp) SetDayLineChartData(v *GetDailyReportChartResponseDayLineChartDataRspDayLineChartData) *GetDailyReportChartResponseDayLineChartDataRsp {
	s.DayLineChartData = v
	return s
}

type GetDailyReportChartResponseDayLineChartDataRspDayLineChartData struct {
	Harris        []*GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataHarris        `json:"Harris,omitempty" xml:"Harris,omitempty" require:"true" type:"Repeated"`
	CloseLoop     []*GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataCloseLoop     `json:"CloseLoop,omitempty" xml:"CloseLoop,omitempty" require:"true" type:"Repeated"`
	Bode          []*GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataBode          `json:"Bode,omitempty" xml:"Bode,omitempty" require:"true" type:"Repeated"`
	ResidualStage []*GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataResidualStage `json:"ResidualStage,omitempty" xml:"ResidualStage,omitempty" require:"true" type:"Repeated"`
}

func (s GetDailyReportChartResponseDayLineChartDataRspDayLineChartData) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportChartResponseDayLineChartDataRspDayLineChartData) GoString() string {
	return s.String()
}

func (s *GetDailyReportChartResponseDayLineChartDataRspDayLineChartData) SetHarris(v []*GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataHarris) *GetDailyReportChartResponseDayLineChartDataRspDayLineChartData {
	s.Harris = v
	return s
}

func (s *GetDailyReportChartResponseDayLineChartDataRspDayLineChartData) SetCloseLoop(v []*GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataCloseLoop) *GetDailyReportChartResponseDayLineChartDataRspDayLineChartData {
	s.CloseLoop = v
	return s
}

func (s *GetDailyReportChartResponseDayLineChartDataRspDayLineChartData) SetBode(v []*GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataBode) *GetDailyReportChartResponseDayLineChartDataRspDayLineChartData {
	s.Bode = v
	return s
}

func (s *GetDailyReportChartResponseDayLineChartDataRspDayLineChartData) SetResidualStage(v []*GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataResidualStage) *GetDailyReportChartResponseDayLineChartDataRspDayLineChartData {
	s.ResidualStage = v
	return s
}

type GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataHarris struct {
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataHarris) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataHarris) GoString() string {
	return s.String()
}

func (s *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataHarris) SetXlabel(v float32) *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataHarris {
	s.Xlabel = &v
	return s
}

func (s *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataHarris) SetYlabel(v float32) *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataHarris {
	s.Ylabel = &v
	return s
}

func (s *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataHarris) SetName(v string) *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataHarris {
	s.Name = &v
	return s
}

type GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataCloseLoop struct {
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataCloseLoop) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataCloseLoop) GoString() string {
	return s.String()
}

func (s *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataCloseLoop) SetXlabel(v float32) *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataCloseLoop {
	s.Xlabel = &v
	return s
}

func (s *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataCloseLoop) SetYlabel(v float32) *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataCloseLoop {
	s.Ylabel = &v
	return s
}

func (s *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataCloseLoop) SetName(v string) *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataCloseLoop {
	s.Name = &v
	return s
}

type GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataBode struct {
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataBode) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataBode) GoString() string {
	return s.String()
}

func (s *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataBode) SetXlabel(v float32) *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataBode {
	s.Xlabel = &v
	return s
}

func (s *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataBode) SetYlabel(v float32) *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataBode {
	s.Ylabel = &v
	return s
}

func (s *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataBode) SetName(v string) *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataBode {
	s.Name = &v
	return s
}

type GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataResidualStage struct {
	Xlabel *float32 `json:"Xlabel,omitempty" xml:"Xlabel,omitempty" require:"true"`
	Ylabel *float32 `json:"Ylabel,omitempty" xml:"Ylabel,omitempty" require:"true"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataResidualStage) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataResidualStage) GoString() string {
	return s.String()
}

func (s *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataResidualStage) SetXlabel(v float32) *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataResidualStage {
	s.Xlabel = &v
	return s
}

func (s *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataResidualStage) SetYlabel(v float32) *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataResidualStage {
	s.Ylabel = &v
	return s
}

func (s *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataResidualStage) SetName(v string) *GetDailyReportChartResponseDayLineChartDataRspDayLineChartDataResidualStage {
	s.Name = &v
	return s
}

type ListPidOrganizationsRequest struct {
	ParentOrganizationId *string `json:"ParentOrganizationId,omitempty" xml:"ParentOrganizationId,omitempty"`
}

func (s ListPidOrganizationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPidOrganizationsRequest) GoString() string {
	return s.String()
}

func (s *ListPidOrganizationsRequest) SetParentOrganizationId(v string) *ListPidOrganizationsRequest {
	s.ParentOrganizationId = &v
	return s
}

type ListPidOrganizationsResponse struct {
	RequestId        *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success          *bool                                           `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Code             *string                                         `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message          *string                                         `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	OrganizationList []*ListPidOrganizationsResponseOrganizationList `json:"OrganizationList,omitempty" xml:"OrganizationList,omitempty" require:"true" type:"Repeated"`
}

func (s ListPidOrganizationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPidOrganizationsResponse) GoString() string {
	return s.String()
}

func (s *ListPidOrganizationsResponse) SetRequestId(v string) *ListPidOrganizationsResponse {
	s.RequestId = &v
	return s
}

func (s *ListPidOrganizationsResponse) SetSuccess(v bool) *ListPidOrganizationsResponse {
	s.Success = &v
	return s
}

func (s *ListPidOrganizationsResponse) SetCode(v string) *ListPidOrganizationsResponse {
	s.Code = &v
	return s
}

func (s *ListPidOrganizationsResponse) SetMessage(v string) *ListPidOrganizationsResponse {
	s.Message = &v
	return s
}

func (s *ListPidOrganizationsResponse) SetOrganizationList(v []*ListPidOrganizationsResponseOrganizationList) *ListPidOrganizationsResponse {
	s.OrganizationList = v
	return s
}

type ListPidOrganizationsResponseOrganizationList struct {
	OrganizationId       *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty" require:"true"`
	OrganizationName     *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty" require:"true"`
	OrganizationLevel    *int    `json:"OrganizationLevel,omitempty" xml:"OrganizationLevel,omitempty" require:"true"`
	ParentOrganizationId *string `json:"ParentOrganizationId,omitempty" xml:"ParentOrganizationId,omitempty" require:"true"`
	LevelName            *string `json:"LevelName,omitempty" xml:"LevelName,omitempty" require:"true"`
}

func (s ListPidOrganizationsResponseOrganizationList) String() string {
	return tea.Prettify(s)
}

func (s ListPidOrganizationsResponseOrganizationList) GoString() string {
	return s.String()
}

func (s *ListPidOrganizationsResponseOrganizationList) SetOrganizationId(v string) *ListPidOrganizationsResponseOrganizationList {
	s.OrganizationId = &v
	return s
}

func (s *ListPidOrganizationsResponseOrganizationList) SetOrganizationName(v string) *ListPidOrganizationsResponseOrganizationList {
	s.OrganizationName = &v
	return s
}

func (s *ListPidOrganizationsResponseOrganizationList) SetOrganizationLevel(v int) *ListPidOrganizationsResponseOrganizationList {
	s.OrganizationLevel = &v
	return s
}

func (s *ListPidOrganizationsResponseOrganizationList) SetParentOrganizationId(v string) *ListPidOrganizationsResponseOrganizationList {
	s.ParentOrganizationId = &v
	return s
}

func (s *ListPidOrganizationsResponseOrganizationList) SetLevelName(v string) *ListPidOrganizationsResponseOrganizationList {
	s.LevelName = &v
	return s
}

type CreatePidDataProcessesRequest struct {
	PidProjectId    *string                                         `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	PidLoopId       *string                                         `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
	DataProcessTime []*CreatePidDataProcessesRequestDataProcessTime `json:"DataProcessTime,omitempty" xml:"DataProcessTime,omitempty" require:"true" type:"Repeated"`
	ClientToken     *string                                         `json:"ClientToken,omitempty" xml:"ClientToken,omitempty" require:"true"`
}

func (s CreatePidDataProcessesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePidDataProcessesRequest) GoString() string {
	return s.String()
}

func (s *CreatePidDataProcessesRequest) SetPidProjectId(v string) *CreatePidDataProcessesRequest {
	s.PidProjectId = &v
	return s
}

func (s *CreatePidDataProcessesRequest) SetPidLoopId(v string) *CreatePidDataProcessesRequest {
	s.PidLoopId = &v
	return s
}

func (s *CreatePidDataProcessesRequest) SetDataProcessTime(v []*CreatePidDataProcessesRequestDataProcessTime) *CreatePidDataProcessesRequest {
	s.DataProcessTime = v
	return s
}

func (s *CreatePidDataProcessesRequest) SetClientToken(v string) *CreatePidDataProcessesRequest {
	s.ClientToken = &v
	return s
}

type CreatePidDataProcessesRequestDataProcessTime struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s CreatePidDataProcessesRequestDataProcessTime) String() string {
	return tea.Prettify(s)
}

func (s CreatePidDataProcessesRequestDataProcessTime) GoString() string {
	return s.String()
}

func (s *CreatePidDataProcessesRequestDataProcessTime) SetStartTime(v string) *CreatePidDataProcessesRequestDataProcessTime {
	s.StartTime = &v
	return s
}

func (s *CreatePidDataProcessesRequestDataProcessTime) SetEndTime(v string) *CreatePidDataProcessesRequestDataProcessTime {
	s.EndTime = &v
	return s
}

type CreatePidDataProcessesResponse struct {
	RequestId            *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message              *string   `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code                 *string   `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success              *bool     `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	PidDataProcessIdList []*string `json:"PidDataProcessIdList,omitempty" xml:"PidDataProcessIdList,omitempty" require:"true" type:"Repeated"`
}

func (s CreatePidDataProcessesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePidDataProcessesResponse) GoString() string {
	return s.String()
}

func (s *CreatePidDataProcessesResponse) SetRequestId(v string) *CreatePidDataProcessesResponse {
	s.RequestId = &v
	return s
}

func (s *CreatePidDataProcessesResponse) SetMessage(v string) *CreatePidDataProcessesResponse {
	s.Message = &v
	return s
}

func (s *CreatePidDataProcessesResponse) SetCode(v string) *CreatePidDataProcessesResponse {
	s.Code = &v
	return s
}

func (s *CreatePidDataProcessesResponse) SetSuccess(v bool) *CreatePidDataProcessesResponse {
	s.Success = &v
	return s
}

func (s *CreatePidDataProcessesResponse) SetPidDataProcessIdList(v []*string) *CreatePidDataProcessesResponse {
	s.PidDataProcessIdList = v
	return s
}

type GetDailyReportPidParamRequest struct {
	PidLoopId     *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
	PidReportDate *string `json:"PidReportDate,omitempty" xml:"PidReportDate,omitempty"`
}

func (s GetDailyReportPidParamRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportPidParamRequest) GoString() string {
	return s.String()
}

func (s *GetDailyReportPidParamRequest) SetPidLoopId(v string) *GetDailyReportPidParamRequest {
	s.PidLoopId = &v
	return s
}

func (s *GetDailyReportPidParamRequest) SetPidReportDate(v string) *GetDailyReportPidParamRequest {
	s.PidReportDate = &v
	return s
}

type GetDailyReportPidParamResponse struct {
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code           *string                                       `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success        *bool                                         `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	DayPidParamRsp *GetDailyReportPidParamResponseDayPidParamRsp `json:"DayPidParamRsp,omitempty" xml:"DayPidParamRsp,omitempty" require:"true" type:"Struct"`
}

func (s GetDailyReportPidParamResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportPidParamResponse) GoString() string {
	return s.String()
}

func (s *GetDailyReportPidParamResponse) SetRequestId(v string) *GetDailyReportPidParamResponse {
	s.RequestId = &v
	return s
}

func (s *GetDailyReportPidParamResponse) SetMessage(v string) *GetDailyReportPidParamResponse {
	s.Message = &v
	return s
}

func (s *GetDailyReportPidParamResponse) SetCode(v string) *GetDailyReportPidParamResponse {
	s.Code = &v
	return s
}

func (s *GetDailyReportPidParamResponse) SetSuccess(v bool) *GetDailyReportPidParamResponse {
	s.Success = &v
	return s
}

func (s *GetDailyReportPidParamResponse) SetDayPidParamRsp(v *GetDailyReportPidParamResponseDayPidParamRsp) *GetDailyReportPidParamResponse {
	s.DayPidParamRsp = v
	return s
}

type GetDailyReportPidParamResponseDayPidParamRsp struct {
	Status      *bool                                                    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	DayPidParam *GetDailyReportPidParamResponseDayPidParamRspDayPidParam `json:"DayPidParam,omitempty" xml:"DayPidParam,omitempty" require:"true" type:"Struct"`
}

func (s GetDailyReportPidParamResponseDayPidParamRsp) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportPidParamResponseDayPidParamRsp) GoString() string {
	return s.String()
}

func (s *GetDailyReportPidParamResponseDayPidParamRsp) SetStatus(v bool) *GetDailyReportPidParamResponseDayPidParamRsp {
	s.Status = &v
	return s
}

func (s *GetDailyReportPidParamResponseDayPidParamRsp) SetDayPidParam(v *GetDailyReportPidParamResponseDayPidParamRspDayPidParam) *GetDailyReportPidParamResponseDayPidParamRsp {
	s.DayPidParam = v
	return s
}

type GetDailyReportPidParamResponseDayPidParamRspDayPidParam struct {
	Kp             *float32 `json:"Kp,omitempty" xml:"Kp,omitempty" require:"true"`
	Ti             *float32 `json:"Ti,omitempty" xml:"Ti,omitempty" require:"true"`
	Td             *float32 `json:"Td,omitempty" xml:"Td,omitempty" require:"true"`
	RecommendIndex *float32 `json:"RecommendIndex,omitempty" xml:"RecommendIndex,omitempty" require:"true"`
}

func (s GetDailyReportPidParamResponseDayPidParamRspDayPidParam) String() string {
	return tea.Prettify(s)
}

func (s GetDailyReportPidParamResponseDayPidParamRspDayPidParam) GoString() string {
	return s.String()
}

func (s *GetDailyReportPidParamResponseDayPidParamRspDayPidParam) SetKp(v float32) *GetDailyReportPidParamResponseDayPidParamRspDayPidParam {
	s.Kp = &v
	return s
}

func (s *GetDailyReportPidParamResponseDayPidParamRspDayPidParam) SetTi(v float32) *GetDailyReportPidParamResponseDayPidParamRspDayPidParam {
	s.Ti = &v
	return s
}

func (s *GetDailyReportPidParamResponseDayPidParamRspDayPidParam) SetTd(v float32) *GetDailyReportPidParamResponseDayPidParamRspDayPidParam {
	s.Td = &v
	return s
}

func (s *GetDailyReportPidParamResponseDayPidParamRspDayPidParam) SetRecommendIndex(v float32) *GetDailyReportPidParamResponseDayPidParamRspDayPidParam {
	s.RecommendIndex = &v
	return s
}

type AddParameterToPidLoopRequest struct {
	PidProjectId       *string                `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	PidLoopId          *string                `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
	PidLoopParameterId *string                `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
	AdjustType         *int                   `json:"AdjustType,omitempty" xml:"AdjustType,omitempty" require:"true"`
	ModelDistinguish   map[string]interface{} `json:"ModelDistinguish,omitempty" xml:"ModelDistinguish,omitempty"`
	ParameterTuning    map[string]interface{} `json:"ParameterTuning,omitempty" xml:"ParameterTuning,omitempty"`
}

func (s AddParameterToPidLoopRequest) String() string {
	return tea.Prettify(s)
}

func (s AddParameterToPidLoopRequest) GoString() string {
	return s.String()
}

func (s *AddParameterToPidLoopRequest) SetPidProjectId(v string) *AddParameterToPidLoopRequest {
	s.PidProjectId = &v
	return s
}

func (s *AddParameterToPidLoopRequest) SetPidLoopId(v string) *AddParameterToPidLoopRequest {
	s.PidLoopId = &v
	return s
}

func (s *AddParameterToPidLoopRequest) SetPidLoopParameterId(v string) *AddParameterToPidLoopRequest {
	s.PidLoopParameterId = &v
	return s
}

func (s *AddParameterToPidLoopRequest) SetAdjustType(v int) *AddParameterToPidLoopRequest {
	s.AdjustType = &v
	return s
}

func (s *AddParameterToPidLoopRequest) SetModelDistinguish(v map[string]interface{}) *AddParameterToPidLoopRequest {
	s.ModelDistinguish = v
	return s
}

func (s *AddParameterToPidLoopRequest) SetParameterTuning(v map[string]interface{}) *AddParameterToPidLoopRequest {
	s.ParameterTuning = v
	return s
}

type AddParameterToPidLoopResponse struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	PidLoopParameterId *string `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
}

func (s AddParameterToPidLoopResponse) String() string {
	return tea.Prettify(s)
}

func (s AddParameterToPidLoopResponse) GoString() string {
	return s.String()
}

func (s *AddParameterToPidLoopResponse) SetRequestId(v string) *AddParameterToPidLoopResponse {
	s.RequestId = &v
	return s
}

func (s *AddParameterToPidLoopResponse) SetMessage(v string) *AddParameterToPidLoopResponse {
	s.Message = &v
	return s
}

func (s *AddParameterToPidLoopResponse) SetCode(v string) *AddParameterToPidLoopResponse {
	s.Code = &v
	return s
}

func (s *AddParameterToPidLoopResponse) SetSuccess(v bool) *AddParameterToPidLoopResponse {
	s.Success = &v
	return s
}

func (s *AddParameterToPidLoopResponse) SetPidLoopParameterId(v string) *AddParameterToPidLoopResponse {
	s.PidLoopParameterId = &v
	return s
}

type ListPidProjectsRequest struct {
	PidProjectName    *string `json:"PidProjectName,omitempty" xml:"PidProjectName,omitempty"`
	PidOrganisationId *string `json:"PidOrganisationId,omitempty" xml:"PidOrganisationId,omitempty"`
	PageSize          *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage       *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s ListPidProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPidProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListPidProjectsRequest) SetPidProjectName(v string) *ListPidProjectsRequest {
	s.PidProjectName = &v
	return s
}

func (s *ListPidProjectsRequest) SetPidOrganisationId(v string) *ListPidProjectsRequest {
	s.PidOrganisationId = &v
	return s
}

func (s *ListPidProjectsRequest) SetPageSize(v int) *ListPidProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPidProjectsRequest) SetCurrentPage(v int) *ListPidProjectsRequest {
	s.CurrentPage = &v
	return s
}

type ListPidProjectsResponse struct {
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	PageSize       *int                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	CurrentPage    *int                                     `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	TotalCount     *int64                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PidProjectList []*ListPidProjectsResponsePidProjectList `json:"PidProjectList,omitempty" xml:"PidProjectList,omitempty" require:"true" type:"Repeated"`
}

func (s ListPidProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPidProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListPidProjectsResponse) SetRequestId(v string) *ListPidProjectsResponse {
	s.RequestId = &v
	return s
}

func (s *ListPidProjectsResponse) SetMessage(v string) *ListPidProjectsResponse {
	s.Message = &v
	return s
}

func (s *ListPidProjectsResponse) SetCode(v string) *ListPidProjectsResponse {
	s.Code = &v
	return s
}

func (s *ListPidProjectsResponse) SetSuccess(v bool) *ListPidProjectsResponse {
	s.Success = &v
	return s
}

func (s *ListPidProjectsResponse) SetPageSize(v int) *ListPidProjectsResponse {
	s.PageSize = &v
	return s
}

func (s *ListPidProjectsResponse) SetCurrentPage(v int) *ListPidProjectsResponse {
	s.CurrentPage = &v
	return s
}

func (s *ListPidProjectsResponse) SetTotalCount(v int64) *ListPidProjectsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListPidProjectsResponse) SetPidProjectList(v []*ListPidProjectsResponsePidProjectList) *ListPidProjectsResponse {
	s.PidProjectList = v
	return s
}

type ListPidProjectsResponsePidProjectList struct {
	PidProjectId      *string `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	PidProjectName    *string `json:"PidProjectName,omitempty" xml:"PidProjectName,omitempty" require:"true"`
	PidProjectDesc    *string `json:"PidProjectDesc,omitempty" xml:"PidProjectDesc,omitempty" require:"true"`
	PidOrganisationId *string `json:"PidOrganisationId,omitempty" xml:"PidOrganisationId,omitempty" require:"true"`
}

func (s ListPidProjectsResponsePidProjectList) String() string {
	return tea.Prettify(s)
}

func (s ListPidProjectsResponsePidProjectList) GoString() string {
	return s.String()
}

func (s *ListPidProjectsResponsePidProjectList) SetPidProjectId(v string) *ListPidProjectsResponsePidProjectList {
	s.PidProjectId = &v
	return s
}

func (s *ListPidProjectsResponsePidProjectList) SetPidProjectName(v string) *ListPidProjectsResponsePidProjectList {
	s.PidProjectName = &v
	return s
}

func (s *ListPidProjectsResponsePidProjectList) SetPidProjectDesc(v string) *ListPidProjectsResponsePidProjectList {
	s.PidProjectDesc = &v
	return s
}

func (s *ListPidProjectsResponsePidProjectList) SetPidOrganisationId(v string) *ListPidProjectsResponsePidProjectList {
	s.PidOrganisationId = &v
	return s
}

type DeletePidOrganizationRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty" require:"true"`
}

func (s DeletePidOrganizationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePidOrganizationRequest) GoString() string {
	return s.String()
}

func (s *DeletePidOrganizationRequest) SetOrganizationId(v string) *DeletePidOrganizationRequest {
	s.OrganizationId = &v
	return s
}

type DeletePidOrganizationResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s DeletePidOrganizationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePidOrganizationResponse) GoString() string {
	return s.String()
}

func (s *DeletePidOrganizationResponse) SetRequestId(v string) *DeletePidOrganizationResponse {
	s.RequestId = &v
	return s
}

func (s *DeletePidOrganizationResponse) SetMessage(v string) *DeletePidOrganizationResponse {
	s.Message = &v
	return s
}

func (s *DeletePidOrganizationResponse) SetCode(v string) *DeletePidOrganizationResponse {
	s.Code = &v
	return s
}

func (s *DeletePidOrganizationResponse) SetSuccess(v bool) *DeletePidOrganizationResponse {
	s.Success = &v
	return s
}

type SetPidLoopDefaultParameterRequest struct {
	PidLoopId          *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
	PidLoopParameterId *string `json:"PidLoopParameterId,omitempty" xml:"PidLoopParameterId,omitempty" require:"true"`
}

func (s SetPidLoopDefaultParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPidLoopDefaultParameterRequest) GoString() string {
	return s.String()
}

func (s *SetPidLoopDefaultParameterRequest) SetPidLoopId(v string) *SetPidLoopDefaultParameterRequest {
	s.PidLoopId = &v
	return s
}

func (s *SetPidLoopDefaultParameterRequest) SetPidLoopParameterId(v string) *SetPidLoopDefaultParameterRequest {
	s.PidLoopParameterId = &v
	return s
}

type SetPidLoopDefaultParameterResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s SetPidLoopDefaultParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s SetPidLoopDefaultParameterResponse) GoString() string {
	return s.String()
}

func (s *SetPidLoopDefaultParameterResponse) SetRequestId(v string) *SetPidLoopDefaultParameterResponse {
	s.RequestId = &v
	return s
}

func (s *SetPidLoopDefaultParameterResponse) SetMessage(v string) *SetPidLoopDefaultParameterResponse {
	s.Message = &v
	return s
}

func (s *SetPidLoopDefaultParameterResponse) SetCode(v string) *SetPidLoopDefaultParameterResponse {
	s.Code = &v
	return s
}

func (s *SetPidLoopDefaultParameterResponse) SetSuccess(v string) *SetPidLoopDefaultParameterResponse {
	s.Success = &v
	return s
}

type GetPidOrganizationParentIdsRequest struct {
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty" require:"true"`
}

func (s GetPidOrganizationParentIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPidOrganizationParentIdsRequest) GoString() string {
	return s.String()
}

func (s *GetPidOrganizationParentIdsRequest) SetOrganizationId(v string) *GetPidOrganizationParentIdsRequest {
	s.OrganizationId = &v
	return s
}

type GetPidOrganizationParentIdsResponse struct {
	RequestId          *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message            *string   `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code               *string   `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success            *bool     `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	OrganizationIdList []*string `json:"OrganizationIdList,omitempty" xml:"OrganizationIdList,omitempty" require:"true" type:"Repeated"`
}

func (s GetPidOrganizationParentIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPidOrganizationParentIdsResponse) GoString() string {
	return s.String()
}

func (s *GetPidOrganizationParentIdsResponse) SetRequestId(v string) *GetPidOrganizationParentIdsResponse {
	s.RequestId = &v
	return s
}

func (s *GetPidOrganizationParentIdsResponse) SetMessage(v string) *GetPidOrganizationParentIdsResponse {
	s.Message = &v
	return s
}

func (s *GetPidOrganizationParentIdsResponse) SetCode(v string) *GetPidOrganizationParentIdsResponse {
	s.Code = &v
	return s
}

func (s *GetPidOrganizationParentIdsResponse) SetSuccess(v bool) *GetPidOrganizationParentIdsResponse {
	s.Success = &v
	return s
}

func (s *GetPidOrganizationParentIdsResponse) SetOrganizationIdList(v []*string) *GetPidOrganizationParentIdsResponse {
	s.OrganizationIdList = v
	return s
}

type CreatePidOrganizationRequest struct {
	OrganizationName     *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty" require:"true"`
	ParentOrganizationId *string `json:"ParentOrganizationId,omitempty" xml:"ParentOrganizationId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty" require:"true"`
}

func (s CreatePidOrganizationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePidOrganizationRequest) GoString() string {
	return s.String()
}

func (s *CreatePidOrganizationRequest) SetOrganizationName(v string) *CreatePidOrganizationRequest {
	s.OrganizationName = &v
	return s
}

func (s *CreatePidOrganizationRequest) SetParentOrganizationId(v string) *CreatePidOrganizationRequest {
	s.ParentOrganizationId = &v
	return s
}

func (s *CreatePidOrganizationRequest) SetClientToken(v string) *CreatePidOrganizationRequest {
	s.ClientToken = &v
	return s
}

type CreatePidOrganizationResponse struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty" require:"true"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s CreatePidOrganizationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePidOrganizationResponse) GoString() string {
	return s.String()
}

func (s *CreatePidOrganizationResponse) SetRequestId(v string) *CreatePidOrganizationResponse {
	s.RequestId = &v
	return s
}

func (s *CreatePidOrganizationResponse) SetOrganizationId(v string) *CreatePidOrganizationResponse {
	s.OrganizationId = &v
	return s
}

func (s *CreatePidOrganizationResponse) SetCode(v string) *CreatePidOrganizationResponse {
	s.Code = &v
	return s
}

func (s *CreatePidOrganizationResponse) SetSuccess(v bool) *CreatePidOrganizationResponse {
	s.Success = &v
	return s
}

func (s *CreatePidOrganizationResponse) SetMessage(v string) *CreatePidOrganizationResponse {
	s.Message = &v
	return s
}

type UpdatePidOrganizationRequest struct {
	OrganizationId   *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty" require:"true"`
	OrganizationName *string `json:"OrganizationName,omitempty" xml:"OrganizationName,omitempty" require:"true"`
}

func (s UpdatePidOrganizationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePidOrganizationRequest) GoString() string {
	return s.String()
}

func (s *UpdatePidOrganizationRequest) SetOrganizationId(v string) *UpdatePidOrganizationRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdatePidOrganizationRequest) SetOrganizationName(v string) *UpdatePidOrganizationRequest {
	s.OrganizationName = &v
	return s
}

type UpdatePidOrganizationResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s UpdatePidOrganizationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePidOrganizationResponse) GoString() string {
	return s.String()
}

func (s *UpdatePidOrganizationResponse) SetRequestId(v string) *UpdatePidOrganizationResponse {
	s.RequestId = &v
	return s
}

func (s *UpdatePidOrganizationResponse) SetMessage(v string) *UpdatePidOrganizationResponse {
	s.Message = &v
	return s
}

func (s *UpdatePidOrganizationResponse) SetCode(v string) *UpdatePidOrganizationResponse {
	s.Code = &v
	return s
}

func (s *UpdatePidOrganizationResponse) SetSuccess(v bool) *UpdatePidOrganizationResponse {
	s.Success = &v
	return s
}

type CreatePidLoopRequest struct {
	PidLoopName          *string                `json:"PidLoopName,omitempty" xml:"PidLoopName,omitempty" require:"true"`
	IsCrucialPidLoop     *bool                  `json:"IsCrucialPidLoop,omitempty" xml:"IsCrucialPidLoop,omitempty" require:"true"`
	PidLoopDesc          *string                `json:"PidLoopDesc,omitempty" xml:"PidLoopDesc,omitempty"`
	PidProjectId         *string                `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	PidLoopDcsType       *string                `json:"PidLoopDcsType,omitempty" xml:"PidLoopDcsType,omitempty" require:"true"`
	PidLoopType          *string                `json:"PidLoopType,omitempty" xml:"PidLoopType,omitempty" require:"true"`
	PidLoopConfiguration map[string]interface{} `json:"PidLoopConfiguration,omitempty" xml:"PidLoopConfiguration,omitempty" require:"true"`
	ClientToken          *string                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty" require:"true"`
}

func (s CreatePidLoopRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePidLoopRequest) GoString() string {
	return s.String()
}

func (s *CreatePidLoopRequest) SetPidLoopName(v string) *CreatePidLoopRequest {
	s.PidLoopName = &v
	return s
}

func (s *CreatePidLoopRequest) SetIsCrucialPidLoop(v bool) *CreatePidLoopRequest {
	s.IsCrucialPidLoop = &v
	return s
}

func (s *CreatePidLoopRequest) SetPidLoopDesc(v string) *CreatePidLoopRequest {
	s.PidLoopDesc = &v
	return s
}

func (s *CreatePidLoopRequest) SetPidProjectId(v string) *CreatePidLoopRequest {
	s.PidProjectId = &v
	return s
}

func (s *CreatePidLoopRequest) SetPidLoopDcsType(v string) *CreatePidLoopRequest {
	s.PidLoopDcsType = &v
	return s
}

func (s *CreatePidLoopRequest) SetPidLoopType(v string) *CreatePidLoopRequest {
	s.PidLoopType = &v
	return s
}

func (s *CreatePidLoopRequest) SetPidLoopConfiguration(v map[string]interface{}) *CreatePidLoopRequest {
	s.PidLoopConfiguration = v
	return s
}

func (s *CreatePidLoopRequest) SetClientToken(v string) *CreatePidLoopRequest {
	s.ClientToken = &v
	return s
}

type CreatePidLoopResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	PidLoopId *string `json:"PidLoopId,omitempty" xml:"PidLoopId,omitempty" require:"true"`
}

func (s CreatePidLoopResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePidLoopResponse) GoString() string {
	return s.String()
}

func (s *CreatePidLoopResponse) SetRequestId(v string) *CreatePidLoopResponse {
	s.RequestId = &v
	return s
}

func (s *CreatePidLoopResponse) SetMessage(v string) *CreatePidLoopResponse {
	s.Message = &v
	return s
}

func (s *CreatePidLoopResponse) SetCode(v string) *CreatePidLoopResponse {
	s.Code = &v
	return s
}

func (s *CreatePidLoopResponse) SetSuccess(v bool) *CreatePidLoopResponse {
	s.Success = &v
	return s
}

func (s *CreatePidLoopResponse) SetPidLoopId(v string) *CreatePidLoopResponse {
	s.PidLoopId = &v
	return s
}

type UpdatePidProjectRequest struct {
	PidProjectId      *string `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	PidProjectName    *string `json:"PidProjectName,omitempty" xml:"PidProjectName,omitempty"`
	PidProjectDesc    *string `json:"PidProjectDesc,omitempty" xml:"PidProjectDesc,omitempty"`
	PidOrganisationId *string `json:"PidOrganisationId,omitempty" xml:"PidOrganisationId,omitempty" require:"true"`
}

func (s UpdatePidProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePidProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdatePidProjectRequest) SetPidProjectId(v string) *UpdatePidProjectRequest {
	s.PidProjectId = &v
	return s
}

func (s *UpdatePidProjectRequest) SetPidProjectName(v string) *UpdatePidProjectRequest {
	s.PidProjectName = &v
	return s
}

func (s *UpdatePidProjectRequest) SetPidProjectDesc(v string) *UpdatePidProjectRequest {
	s.PidProjectDesc = &v
	return s
}

func (s *UpdatePidProjectRequest) SetPidOrganisationId(v string) *UpdatePidProjectRequest {
	s.PidOrganisationId = &v
	return s
}

type UpdatePidProjectResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s UpdatePidProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePidProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdatePidProjectResponse) SetRequestId(v string) *UpdatePidProjectResponse {
	s.RequestId = &v
	return s
}

func (s *UpdatePidProjectResponse) SetMessage(v string) *UpdatePidProjectResponse {
	s.Message = &v
	return s
}

func (s *UpdatePidProjectResponse) SetCode(v string) *UpdatePidProjectResponse {
	s.Code = &v
	return s
}

func (s *UpdatePidProjectResponse) SetSuccess(v bool) *UpdatePidProjectResponse {
	s.Success = &v
	return s
}

type ListPidTagsRequest struct {
	PidProjectId *string `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	PidTagName   *string `json:"PidTagName,omitempty" xml:"PidTagName,omitempty"`
	CurrentPage  *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize     *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPidTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPidTagsRequest) GoString() string {
	return s.String()
}

func (s *ListPidTagsRequest) SetPidProjectId(v string) *ListPidTagsRequest {
	s.PidProjectId = &v
	return s
}

func (s *ListPidTagsRequest) SetPidTagName(v string) *ListPidTagsRequest {
	s.PidTagName = &v
	return s
}

func (s *ListPidTagsRequest) SetCurrentPage(v int) *ListPidTagsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListPidTagsRequest) SetPageSize(v int) *ListPidTagsRequest {
	s.PageSize = &v
	return s
}

type ListPidTagsResponse struct {
	RequestId   *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message     *string                          `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code        *string                          `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success     *bool                            `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	CurrentPage *int                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize    *int                             `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount  *int64                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PidTagList  []*ListPidTagsResponsePidTagList `json:"PidTagList,omitempty" xml:"PidTagList,omitempty" require:"true" type:"Repeated"`
}

func (s ListPidTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPidTagsResponse) GoString() string {
	return s.String()
}

func (s *ListPidTagsResponse) SetRequestId(v string) *ListPidTagsResponse {
	s.RequestId = &v
	return s
}

func (s *ListPidTagsResponse) SetMessage(v string) *ListPidTagsResponse {
	s.Message = &v
	return s
}

func (s *ListPidTagsResponse) SetCode(v string) *ListPidTagsResponse {
	s.Code = &v
	return s
}

func (s *ListPidTagsResponse) SetSuccess(v bool) *ListPidTagsResponse {
	s.Success = &v
	return s
}

func (s *ListPidTagsResponse) SetCurrentPage(v int) *ListPidTagsResponse {
	s.CurrentPage = &v
	return s
}

func (s *ListPidTagsResponse) SetPageSize(v int) *ListPidTagsResponse {
	s.PageSize = &v
	return s
}

func (s *ListPidTagsResponse) SetTotalCount(v int64) *ListPidTagsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListPidTagsResponse) SetPidTagList(v []*ListPidTagsResponsePidTagList) *ListPidTagsResponse {
	s.PidTagList = v
	return s
}

type ListPidTagsResponsePidTagList struct {
	PidTagId        *string `json:"PidTagId,omitempty" xml:"PidTagId,omitempty" require:"true"`
	PidProjectId    *string `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	PidTagName      *string `json:"PidTagName,omitempty" xml:"PidTagName,omitempty" require:"true"`
	PidTagType      *string `json:"PidTagType,omitempty" xml:"PidTagType,omitempty" require:"true"`
	UploadProjectId *int64  `json:"UploadProjectId,omitempty" xml:"UploadProjectId,omitempty" require:"true"`
}

func (s ListPidTagsResponsePidTagList) String() string {
	return tea.Prettify(s)
}

func (s ListPidTagsResponsePidTagList) GoString() string {
	return s.String()
}

func (s *ListPidTagsResponsePidTagList) SetPidTagId(v string) *ListPidTagsResponsePidTagList {
	s.PidTagId = &v
	return s
}

func (s *ListPidTagsResponsePidTagList) SetPidProjectId(v string) *ListPidTagsResponsePidTagList {
	s.PidProjectId = &v
	return s
}

func (s *ListPidTagsResponsePidTagList) SetPidTagName(v string) *ListPidTagsResponsePidTagList {
	s.PidTagName = &v
	return s
}

func (s *ListPidTagsResponsePidTagList) SetPidTagType(v string) *ListPidTagsResponsePidTagList {
	s.PidTagType = &v
	return s
}

func (s *ListPidTagsResponsePidTagList) SetUploadProjectId(v int64) *ListPidTagsResponsePidTagList {
	s.UploadProjectId = &v
	return s
}

type DeletePidDataProcessRequest struct {
	PidDataProcessId *string `json:"PidDataProcessId,omitempty" xml:"PidDataProcessId,omitempty" require:"true"`
}

func (s DeletePidDataProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePidDataProcessRequest) GoString() string {
	return s.String()
}

func (s *DeletePidDataProcessRequest) SetPidDataProcessId(v string) *DeletePidDataProcessRequest {
	s.PidDataProcessId = &v
	return s
}

type DeletePidDataProcessResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	PidData   *string `json:"PidData,omitempty" xml:"PidData,omitempty" require:"true"`
}

func (s DeletePidDataProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePidDataProcessResponse) GoString() string {
	return s.String()
}

func (s *DeletePidDataProcessResponse) SetRequestId(v string) *DeletePidDataProcessResponse {
	s.RequestId = &v
	return s
}

func (s *DeletePidDataProcessResponse) SetMessage(v string) *DeletePidDataProcessResponse {
	s.Message = &v
	return s
}

func (s *DeletePidDataProcessResponse) SetCode(v string) *DeletePidDataProcessResponse {
	s.Code = &v
	return s
}

func (s *DeletePidDataProcessResponse) SetSuccess(v bool) *DeletePidDataProcessResponse {
	s.Success = &v
	return s
}

func (s *DeletePidDataProcessResponse) SetPidData(v string) *DeletePidDataProcessResponse {
	s.PidData = &v
	return s
}

type CreatePidDataSourceRequest struct {
	PidProjectId   *string `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	OssPath        *string `json:"OssPath,omitempty" xml:"OssPath,omitempty" require:"true"`
	FileName       *string `json:"FileName,omitempty" xml:"FileName,omitempty" require:"true"`
	OpenUpload     *int    `json:"OpenUpload,omitempty" xml:"OpenUpload,omitempty"`
	NeedCreateTime *int    `json:"NeedCreateTime,omitempty" xml:"NeedCreateTime,omitempty" require:"true"`
	StartDate      *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	SampleTime     *int    `json:"SampleTime,omitempty" xml:"SampleTime,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty" require:"true"`
}

func (s CreatePidDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePidDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreatePidDataSourceRequest) SetPidProjectId(v string) *CreatePidDataSourceRequest {
	s.PidProjectId = &v
	return s
}

func (s *CreatePidDataSourceRequest) SetType(v string) *CreatePidDataSourceRequest {
	s.Type = &v
	return s
}

func (s *CreatePidDataSourceRequest) SetOssPath(v string) *CreatePidDataSourceRequest {
	s.OssPath = &v
	return s
}

func (s *CreatePidDataSourceRequest) SetFileName(v string) *CreatePidDataSourceRequest {
	s.FileName = &v
	return s
}

func (s *CreatePidDataSourceRequest) SetOpenUpload(v int) *CreatePidDataSourceRequest {
	s.OpenUpload = &v
	return s
}

func (s *CreatePidDataSourceRequest) SetNeedCreateTime(v int) *CreatePidDataSourceRequest {
	s.NeedCreateTime = &v
	return s
}

func (s *CreatePidDataSourceRequest) SetStartDate(v string) *CreatePidDataSourceRequest {
	s.StartDate = &v
	return s
}

func (s *CreatePidDataSourceRequest) SetSampleTime(v int) *CreatePidDataSourceRequest {
	s.SampleTime = &v
	return s
}

func (s *CreatePidDataSourceRequest) SetClientToken(v string) *CreatePidDataSourceRequest {
	s.ClientToken = &v
	return s
}

type CreatePidDataSourceResponse struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success         *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	PidDataSourceId *string `json:"PidDataSourceId,omitempty" xml:"PidDataSourceId,omitempty" require:"true"`
}

func (s CreatePidDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePidDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CreatePidDataSourceResponse) SetRequestId(v string) *CreatePidDataSourceResponse {
	s.RequestId = &v
	return s
}

func (s *CreatePidDataSourceResponse) SetSuccess(v bool) *CreatePidDataSourceResponse {
	s.Success = &v
	return s
}

func (s *CreatePidDataSourceResponse) SetCode(v string) *CreatePidDataSourceResponse {
	s.Code = &v
	return s
}

func (s *CreatePidDataSourceResponse) SetMessage(v string) *CreatePidDataSourceResponse {
	s.Message = &v
	return s
}

func (s *CreatePidDataSourceResponse) SetPidDataSourceId(v string) *CreatePidDataSourceResponse {
	s.PidDataSourceId = &v
	return s
}

type CreatePidDataSourceAdvanceRequest struct {
	OssPathObject  io.Reader `json:"OssPathObject,omitempty" xml:"OssPathObject,omitempty" require:"true"`
	PidProjectId   *string   `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	Type           *string   `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	FileName       *string   `json:"FileName,omitempty" xml:"FileName,omitempty" require:"true"`
	OpenUpload     *int      `json:"OpenUpload,omitempty" xml:"OpenUpload,omitempty"`
	NeedCreateTime *int      `json:"NeedCreateTime,omitempty" xml:"NeedCreateTime,omitempty" require:"true"`
	StartDate      *string   `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	SampleTime     *int      `json:"SampleTime,omitempty" xml:"SampleTime,omitempty"`
	ClientToken    *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty" require:"true"`
}

func (s CreatePidDataSourceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePidDataSourceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreatePidDataSourceAdvanceRequest) SetOssPathObject(v io.Reader) *CreatePidDataSourceAdvanceRequest {
	s.OssPathObject = v
	return s
}

func (s *CreatePidDataSourceAdvanceRequest) SetPidProjectId(v string) *CreatePidDataSourceAdvanceRequest {
	s.PidProjectId = &v
	return s
}

func (s *CreatePidDataSourceAdvanceRequest) SetType(v string) *CreatePidDataSourceAdvanceRequest {
	s.Type = &v
	return s
}

func (s *CreatePidDataSourceAdvanceRequest) SetFileName(v string) *CreatePidDataSourceAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *CreatePidDataSourceAdvanceRequest) SetOpenUpload(v int) *CreatePidDataSourceAdvanceRequest {
	s.OpenUpload = &v
	return s
}

func (s *CreatePidDataSourceAdvanceRequest) SetNeedCreateTime(v int) *CreatePidDataSourceAdvanceRequest {
	s.NeedCreateTime = &v
	return s
}

func (s *CreatePidDataSourceAdvanceRequest) SetStartDate(v string) *CreatePidDataSourceAdvanceRequest {
	s.StartDate = &v
	return s
}

func (s *CreatePidDataSourceAdvanceRequest) SetSampleTime(v int) *CreatePidDataSourceAdvanceRequest {
	s.SampleTime = &v
	return s
}

func (s *CreatePidDataSourceAdvanceRequest) SetClientToken(v string) *CreatePidDataSourceAdvanceRequest {
	s.ClientToken = &v
	return s
}

type DeletePidProjectRequest struct {
	PidProjectId *string `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
}

func (s DeletePidProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePidProjectRequest) GoString() string {
	return s.String()
}

func (s *DeletePidProjectRequest) SetPidProjectId(v string) *DeletePidProjectRequest {
	s.PidProjectId = &v
	return s
}

type DeletePidProjectResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s DeletePidProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePidProjectResponse) GoString() string {
	return s.String()
}

func (s *DeletePidProjectResponse) SetRequestId(v string) *DeletePidProjectResponse {
	s.RequestId = &v
	return s
}

func (s *DeletePidProjectResponse) SetMessage(v string) *DeletePidProjectResponse {
	s.Message = &v
	return s
}

func (s *DeletePidProjectResponse) SetCode(v string) *DeletePidProjectResponse {
	s.Code = &v
	return s
}

func (s *DeletePidProjectResponse) SetSuccess(v bool) *DeletePidProjectResponse {
	s.Success = &v
	return s
}

type DeletePidTagRequest struct {
	PidProjectId *string `json:"PidProjectId,omitempty" xml:"PidProjectId,omitempty" require:"true"`
	PidTagId     *string `json:"PidTagId,omitempty" xml:"PidTagId,omitempty" require:"true"`
}

func (s DeletePidTagRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePidTagRequest) GoString() string {
	return s.String()
}

func (s *DeletePidTagRequest) SetPidProjectId(v string) *DeletePidTagRequest {
	s.PidProjectId = &v
	return s
}

func (s *DeletePidTagRequest) SetPidTagId(v string) *DeletePidTagRequest {
	s.PidTagId = &v
	return s
}

type DeletePidTagResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s DeletePidTagResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePidTagResponse) GoString() string {
	return s.String()
}

func (s *DeletePidTagResponse) SetRequestId(v string) *DeletePidTagResponse {
	s.RequestId = &v
	return s
}

func (s *DeletePidTagResponse) SetMessage(v string) *DeletePidTagResponse {
	s.Message = &v
	return s
}

func (s *DeletePidTagResponse) SetCode(v string) *DeletePidTagResponse {
	s.Code = &v
	return s
}

func (s *DeletePidTagResponse) SetSuccess(v bool) *DeletePidTagResponse {
	s.Success = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("brain-industrial"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetOpPvCustomValues(request *GetOpPvCustomValuesRequest, runtime *util.RuntimeOptions) (_result *GetOpPvCustomValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetOpPvCustomValuesResponse{}
	_body, _err := client.DoRequest(tea.String("GetOpPvCustomValues"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitPidLoopEvaluation(request *SubmitPidLoopEvaluationRequest, runtime *util.RuntimeOptions) (_result *SubmitPidLoopEvaluationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SubmitPidLoopEvaluationResponse{}
	_body, _err := client.DoRequest(tea.String("SubmitPidLoopEvaluation"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDefaultAdjustValues(request *GetDefaultAdjustValuesRequest, runtime *util.RuntimeOptions) (_result *GetDefaultAdjustValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDefaultAdjustValuesResponse{}
	_body, _err := client.DoRequest(tea.String("GetDefaultAdjustValues"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPvOpAdjustValues(request *GetPvOpAdjustValuesRequest, runtime *util.RuntimeOptions) (_result *GetPvOpAdjustValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPvOpAdjustValuesResponse{}
	_body, _err := client.DoRequest(tea.String("GetPvOpAdjustValues"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPvCustomSimulationValues(request *GetPvCustomSimulationValuesRequest, runtime *util.RuntimeOptions) (_result *GetPvCustomSimulationValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPvCustomSimulationValuesResponse{}
	_body, _err := client.DoRequest(tea.String("GetPvCustomSimulationValues"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIdentificateValues(request *GetIdentificateValuesRequest, runtime *util.RuntimeOptions) (_result *GetIdentificateValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetIdentificateValuesResponse{}
	_body, _err := client.DoRequest(tea.String("GetIdentificateValues"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPvIdentSimulationValues(request *GetPvIdentSimulationValuesRequest, runtime *util.RuntimeOptions) (_result *GetPvIdentSimulationValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPvIdentSimulationValuesResponse{}
	_body, _err := client.DoRequest(tea.String("GetPvIdentSimulationValues"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCustomIdentValues(request *GetCustomIdentValuesRequest, runtime *util.RuntimeOptions) (_result *GetCustomIdentValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetCustomIdentValuesResponse{}
	_body, _err := client.DoRequest(tea.String("GetCustomIdentValues"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPidProjectReportOverview(request *GetPidProjectReportOverviewRequest, runtime *util.RuntimeOptions) (_result *GetPidProjectReportOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPidProjectReportOverviewResponse{}
	_body, _err := client.DoRequest(tea.String("GetPidProjectReportOverview"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLowModelPerformValues(request *GetLowModelPerformValuesRequest, runtime *util.RuntimeOptions) (_result *GetLowModelPerformValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetLowModelPerformValuesResponse{}
	_body, _err := client.DoRequest(tea.String("GetLowModelPerformValues"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPidLoopEvaluations(request *ListPidLoopEvaluationsRequest, runtime *util.RuntimeOptions) (_result *ListPidLoopEvaluationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListPidLoopEvaluationsResponse{}
	_body, _err := client.DoRequest(tea.String("ListPidLoopEvaluations"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLoopParameterTagValues(request *ListLoopParameterTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListLoopParameterTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListLoopParameterTagValuesResponse{}
	_body, _err := client.DoRequest(tea.String("ListLoopParameterTagValues"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPidLoopParameter(request *GetPidLoopParameterRequest, runtime *util.RuntimeOptions) (_result *GetPidLoopParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPidLoopParameterResponse{}
	_body, _err := client.DoRequest(tea.String("GetPidLoopParameter"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDailyReportInfo(request *GetDailyReportInfoRequest, runtime *util.RuntimeOptions) (_result *GetDailyReportInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDailyReportInfoResponse{}
	_body, _err := client.DoRequest(tea.String("GetDailyReportInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPidLoop(request *GetPidLoopRequest, runtime *util.RuntimeOptions) (_result *GetPidLoopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPidLoopResponse{}
	_body, _err := client.DoRequest(tea.String("GetPidLoop"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePidProject(request *CreatePidProjectRequest, runtime *util.RuntimeOptions) (_result *CreatePidProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreatePidProjectResponse{}
	_body, _err := client.DoRequest(tea.String("CreatePidProject"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIdentModels(request *ListIdentModelsRequest, runtime *util.RuntimeOptions) (_result *ListIdentModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListIdentModelsResponse{}
	_body, _err := client.DoRequest(tea.String("ListIdentModels"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPidDataProcess(request *ListPidDataProcessRequest, runtime *util.RuntimeOptions) (_result *ListPidDataProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListPidDataProcessResponse{}
	_body, _err := client.DoRequest(tea.String("ListPidDataProcess"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCustomIdentModel(request *AddCustomIdentModelRequest, runtime *util.RuntimeOptions) (_result *AddCustomIdentModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddCustomIdentModelResponse{}
	_body, _err := client.DoRequest(tea.String("AddCustomIdentModel"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLoopParameterStep(request *GetLoopParameterStepRequest, runtime *util.RuntimeOptions) (_result *GetLoopParameterStepResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetLoopParameterStepResponse{}
	_body, _err := client.DoRequest(tea.String("GetLoopParameterStep"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPidLoops(request *ListPidLoopsRequest, runtime *util.RuntimeOptions) (_result *ListPidLoopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListPidLoopsResponse{}
	_body, _err := client.DoRequest(tea.String("ListPidLoops"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MovePidOrganization(request *MovePidOrganizationRequest, runtime *util.RuntimeOptions) (_result *MovePidOrganizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &MovePidOrganizationResponse{}
	_body, _err := client.DoRequest(tea.String("MovePidOrganization"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePidLoop(request *UpdatePidLoopRequest, runtime *util.RuntimeOptions) (_result *UpdatePidLoopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdatePidLoopResponse{}
	_body, _err := client.DoRequest(tea.String("UpdatePidLoop"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddPidLoopToSchedule(request *AddPidLoopToScheduleRequest, runtime *util.RuntimeOptions) (_result *AddPidLoopToScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddPidLoopToScheduleResponse{}
	_body, _err := client.DoRequest(tea.String("AddPidLoopToSchedule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetParsingTagTask(request *GetParsingTagTaskRequest, runtime *util.RuntimeOptions) (_result *GetParsingTagTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetParsingTagTaskResponse{}
	_body, _err := client.DoRequest(tea.String("GetParsingTagTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPidLoopLatestTaskStatus(request *GetPidLoopLatestTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetPidLoopLatestTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPidLoopLatestTaskStatusResponse{}
	_body, _err := client.DoRequest(tea.String("GetPidLoopLatestTaskStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePidLoop(request *DeletePidLoopRequest, runtime *util.RuntimeOptions) (_result *DeletePidLoopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeletePidLoopResponse{}
	_body, _err := client.DoRequest(tea.String("DeletePidLoop"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSummaryReportInfo(request *GetSummaryReportInfoRequest, runtime *util.RuntimeOptions) (_result *GetSummaryReportInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSummaryReportInfoResponse{}
	_body, _err := client.DoRequest(tea.String("GetSummaryReportInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSummaryReportDataTrend(request *GetSummaryReportDataTrendRequest, runtime *util.RuntimeOptions) (_result *GetSummaryReportDataTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSummaryReportDataTrendResponse{}
	_body, _err := client.DoRequest(tea.String("GetSummaryReportDataTrend"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSummaryReportChart(request *GetSummaryReportChartRequest, runtime *util.RuntimeOptions) (_result *GetSummaryReportChartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSummaryReportChartResponse{}
	_body, _err := client.DoRequest(tea.String("GetSummaryReportChart"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDailyReportDataTrend(request *GetDailyReportDataTrendRequest, runtime *util.RuntimeOptions) (_result *GetDailyReportDataTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDailyReportDataTrendResponse{}
	_body, _err := client.DoRequest(tea.String("GetDailyReportDataTrend"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagValueTrend(request *ListTagValueTrendRequest, runtime *util.RuntimeOptions) (_result *ListTagValueTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListTagValueTrendResponse{}
	_body, _err := client.DoRequest(tea.String("ListTagValueTrend"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDailyReportChart(request *GetDailyReportChartRequest, runtime *util.RuntimeOptions) (_result *GetDailyReportChartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDailyReportChartResponse{}
	_body, _err := client.DoRequest(tea.String("GetDailyReportChart"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPidOrganizations(request *ListPidOrganizationsRequest, runtime *util.RuntimeOptions) (_result *ListPidOrganizationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListPidOrganizationsResponse{}
	_body, _err := client.DoRequest(tea.String("ListPidOrganizations"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePidDataProcesses(request *CreatePidDataProcessesRequest, runtime *util.RuntimeOptions) (_result *CreatePidDataProcessesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreatePidDataProcessesResponse{}
	_body, _err := client.DoRequest(tea.String("CreatePidDataProcesses"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDailyReportPidParam(request *GetDailyReportPidParamRequest, runtime *util.RuntimeOptions) (_result *GetDailyReportPidParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDailyReportPidParamResponse{}
	_body, _err := client.DoRequest(tea.String("GetDailyReportPidParam"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddParameterToPidLoop(request *AddParameterToPidLoopRequest, runtime *util.RuntimeOptions) (_result *AddParameterToPidLoopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddParameterToPidLoopResponse{}
	_body, _err := client.DoRequest(tea.String("AddParameterToPidLoop"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPidProjects(request *ListPidProjectsRequest, runtime *util.RuntimeOptions) (_result *ListPidProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListPidProjectsResponse{}
	_body, _err := client.DoRequest(tea.String("ListPidProjects"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePidOrganization(request *DeletePidOrganizationRequest, runtime *util.RuntimeOptions) (_result *DeletePidOrganizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeletePidOrganizationResponse{}
	_body, _err := client.DoRequest(tea.String("DeletePidOrganization"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetPidLoopDefaultParameter(request *SetPidLoopDefaultParameterRequest, runtime *util.RuntimeOptions) (_result *SetPidLoopDefaultParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetPidLoopDefaultParameterResponse{}
	_body, _err := client.DoRequest(tea.String("SetPidLoopDefaultParameter"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPidOrganizationParentIds(request *GetPidOrganizationParentIdsRequest, runtime *util.RuntimeOptions) (_result *GetPidOrganizationParentIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPidOrganizationParentIdsResponse{}
	_body, _err := client.DoRequest(tea.String("GetPidOrganizationParentIds"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePidOrganization(request *CreatePidOrganizationRequest, runtime *util.RuntimeOptions) (_result *CreatePidOrganizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreatePidOrganizationResponse{}
	_body, _err := client.DoRequest(tea.String("CreatePidOrganization"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePidOrganization(request *UpdatePidOrganizationRequest, runtime *util.RuntimeOptions) (_result *UpdatePidOrganizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdatePidOrganizationResponse{}
	_body, _err := client.DoRequest(tea.String("UpdatePidOrganization"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePidLoop(request *CreatePidLoopRequest, runtime *util.RuntimeOptions) (_result *CreatePidLoopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreatePidLoopResponse{}
	_body, _err := client.DoRequest(tea.String("CreatePidLoop"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePidProject(request *UpdatePidProjectRequest, runtime *util.RuntimeOptions) (_result *UpdatePidProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdatePidProjectResponse{}
	_body, _err := client.DoRequest(tea.String("UpdatePidProject"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPidTags(request *ListPidTagsRequest, runtime *util.RuntimeOptions) (_result *ListPidTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListPidTagsResponse{}
	_body, _err := client.DoRequest(tea.String("ListPidTags"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePidDataProcess(request *DeletePidDataProcessRequest, runtime *util.RuntimeOptions) (_result *DeletePidDataProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeletePidDataProcessResponse{}
	_body, _err := client.DoRequest(tea.String("DeletePidDataProcess"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePidDataSource(request *CreatePidDataSourceRequest, runtime *util.RuntimeOptions) (_result *CreatePidDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreatePidDataSourceResponse{}
	_body, _err := client.DoRequest(tea.String("CreatePidDataSource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePidDataSourceAdvance(request *CreatePidDataSourceAdvanceRequest, runtime *util.RuntimeOptions) (_result *CreatePidDataSourceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("brain-industrial"),
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
	rpcutil.Convert(runtime, ossRuntime)
	createPidDataSourcereq := &CreatePidDataSourceRequest{}
	rpcutil.Convert(request, createPidDataSourcereq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.OssPathObject,
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
	createPidDataSourcereq.OssPath = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	createPidDataSourceResp, _err := client.CreatePidDataSource(createPidDataSourcereq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = createPidDataSourceResp
	return _result, _err
}

func (client *Client) DeletePidProject(request *DeletePidProjectRequest, runtime *util.RuntimeOptions) (_result *DeletePidProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeletePidProjectResponse{}
	_body, _err := client.DoRequest(tea.String("DeletePidProject"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePidTag(request *DeletePidTagRequest, runtime *util.RuntimeOptions) (_result *DeletePidTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeletePidTagResponse{}
	_body, _err := client.DoRequest(tea.String("DeletePidTag"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-09-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
