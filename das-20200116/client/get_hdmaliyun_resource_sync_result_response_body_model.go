// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHDMAliyunResourceSyncResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHDMAliyunResourceSyncResultResponseBody
	GetCode() *string
	SetData(v *GetHDMAliyunResourceSyncResultResponseBodyData) *GetHDMAliyunResourceSyncResultResponseBody
	GetData() *GetHDMAliyunResourceSyncResultResponseBodyData
	SetMessage(v string) *GetHDMAliyunResourceSyncResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHDMAliyunResourceSyncResultResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetHDMAliyunResourceSyncResultResponseBody
	GetSuccess() *string
	SetSynchro(v string) *GetHDMAliyunResourceSyncResultResponseBody
	GetSynchro() *string
}

type GetHDMAliyunResourceSyncResultResponseBody struct {
	Code      *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetHDMAliyunResourceSyncResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string                                         `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetHDMAliyunResourceSyncResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHDMAliyunResourceSyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) GetData() *GetHDMAliyunResourceSyncResultResponseBodyData {
	return s.Data
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) GetSynchro() *string {
	return s.Synchro
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) SetCode(v string) *GetHDMAliyunResourceSyncResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) SetData(v *GetHDMAliyunResourceSyncResultResponseBodyData) *GetHDMAliyunResourceSyncResultResponseBody {
	s.Data = v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) SetMessage(v string) *GetHDMAliyunResourceSyncResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) SetRequestId(v string) *GetHDMAliyunResourceSyncResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) SetSuccess(v string) *GetHDMAliyunResourceSyncResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) SetSynchro(v string) *GetHDMAliyunResourceSyncResultResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHDMAliyunResourceSyncResultResponseBodyData struct {
	ErrorMsg   *string                                                   `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Results    *string                                                   `json:"Results,omitempty" xml:"Results,omitempty"`
	SubResults *GetHDMAliyunResourceSyncResultResponseBodyDataSubResults `json:"SubResults,omitempty" xml:"SubResults,omitempty" type:"Struct"`
	SyncStatus *string                                                   `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
}

func (s GetHDMAliyunResourceSyncResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHDMAliyunResourceSyncResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyData) GetResults() *string {
	return s.Results
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyData) GetSubResults() *GetHDMAliyunResourceSyncResultResponseBodyDataSubResults {
	return s.SubResults
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyData) GetSyncStatus() *string {
	return s.SyncStatus
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyData) SetErrorMsg(v string) *GetHDMAliyunResourceSyncResultResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyData) SetResults(v string) *GetHDMAliyunResourceSyncResultResponseBodyData {
	s.Results = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyData) SetSubResults(v *GetHDMAliyunResourceSyncResultResponseBodyDataSubResults) *GetHDMAliyunResourceSyncResultResponseBodyData {
	s.SubResults = v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyData) SetSyncStatus(v string) *GetHDMAliyunResourceSyncResultResponseBodyData {
	s.SyncStatus = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetHDMAliyunResourceSyncResultResponseBodyDataSubResults struct {
	ResourceSyncSubResult []*GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult `json:"ResourceSyncSubResult,omitempty" xml:"ResourceSyncSubResult,omitempty" type:"Repeated"`
}

func (s GetHDMAliyunResourceSyncResultResponseBodyDataSubResults) String() string {
	return dara.Prettify(s)
}

func (s GetHDMAliyunResourceSyncResultResponseBodyDataSubResults) GoString() string {
	return s.String()
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResults) GetResourceSyncSubResult() []*GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	return s.ResourceSyncSubResult
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResults) SetResourceSyncSubResult(v []*GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) *GetHDMAliyunResourceSyncResultResponseBodyDataSubResults {
	s.ResourceSyncSubResult = v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResults) Validate() error {
	return dara.Validate(s)
}

type GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult struct {
	ErrMsg       *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	SyncCount    *int32  `json:"SyncCount,omitempty" xml:"SyncCount,omitempty"`
}

func (s GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) String() string {
	return dara.Prettify(s)
}

func (s GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) GoString() string {
	return s.String()
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) GetSuccess() *bool {
	return s.Success
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) GetSyncCount() *int32 {
	return s.SyncCount
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) SetErrMsg(v string) *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	s.ErrMsg = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) SetResourceType(v string) *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	s.ResourceType = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) SetSuccess(v bool) *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	s.Success = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) SetSyncCount(v int32) *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	s.SyncCount = &v
	return s
}

func (s *GetHDMAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) Validate() error {
	return dara.Validate(s)
}
