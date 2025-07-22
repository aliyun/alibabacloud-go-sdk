// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHDMLastAliyunResourceSyncResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHDMLastAliyunResourceSyncResultResponseBody
	GetCode() *string
	SetData(v *GetHDMLastAliyunResourceSyncResultResponseBodyData) *GetHDMLastAliyunResourceSyncResultResponseBody
	GetData() *GetHDMLastAliyunResourceSyncResultResponseBodyData
	SetMessage(v string) *GetHDMLastAliyunResourceSyncResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHDMLastAliyunResourceSyncResultResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetHDMLastAliyunResourceSyncResultResponseBody
	GetSuccess() *string
	SetSynchro(v string) *GetHDMLastAliyunResourceSyncResultResponseBody
	GetSynchro() *string
}

type GetHDMLastAliyunResourceSyncResultResponseBody struct {
	Code      *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetHDMLastAliyunResourceSyncResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string                                             `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetHDMLastAliyunResourceSyncResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHDMLastAliyunResourceSyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) GetData() *GetHDMLastAliyunResourceSyncResultResponseBodyData {
	return s.Data
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) GetSynchro() *string {
	return s.Synchro
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) SetCode(v string) *GetHDMLastAliyunResourceSyncResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) SetData(v *GetHDMLastAliyunResourceSyncResultResponseBodyData) *GetHDMLastAliyunResourceSyncResultResponseBody {
	s.Data = v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) SetMessage(v string) *GetHDMLastAliyunResourceSyncResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) SetRequestId(v string) *GetHDMLastAliyunResourceSyncResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) SetSuccess(v string) *GetHDMLastAliyunResourceSyncResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) SetSynchro(v string) *GetHDMLastAliyunResourceSyncResultResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHDMLastAliyunResourceSyncResultResponseBodyData struct {
	ErrorMsg   *string                                                       `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Results    *string                                                       `json:"Results,omitempty" xml:"Results,omitempty"`
	SubResults *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults `json:"SubResults,omitempty" xml:"SubResults,omitempty" type:"Struct"`
	SyncStatus *string                                                       `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
}

func (s GetHDMLastAliyunResourceSyncResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHDMLastAliyunResourceSyncResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyData) GetResults() *string {
	return s.Results
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyData) GetSubResults() *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults {
	return s.SubResults
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyData) GetSyncStatus() *string {
	return s.SyncStatus
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyData) SetErrorMsg(v string) *GetHDMLastAliyunResourceSyncResultResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyData) SetResults(v string) *GetHDMLastAliyunResourceSyncResultResponseBodyData {
	s.Results = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyData) SetSubResults(v *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults) *GetHDMLastAliyunResourceSyncResultResponseBodyData {
	s.SubResults = v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyData) SetSyncStatus(v string) *GetHDMLastAliyunResourceSyncResultResponseBodyData {
	s.SyncStatus = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults struct {
	ResourceSyncSubResult []*GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult `json:"ResourceSyncSubResult,omitempty" xml:"ResourceSyncSubResult,omitempty" type:"Repeated"`
}

func (s GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults) String() string {
	return dara.Prettify(s)
}

func (s GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults) GoString() string {
	return s.String()
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults) GetResourceSyncSubResult() []*GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	return s.ResourceSyncSubResult
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults) SetResourceSyncSubResult(v []*GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults {
	s.ResourceSyncSubResult = v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResults) Validate() error {
	return dara.Validate(s)
}

type GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult struct {
	ErrMsg       *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	SyncCount    *int32  `json:"SyncCount,omitempty" xml:"SyncCount,omitempty"`
}

func (s GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) String() string {
	return dara.Prettify(s)
}

func (s GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) GoString() string {
	return s.String()
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) GetSuccess() *bool {
	return s.Success
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) GetSyncCount() *int32 {
	return s.SyncCount
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) SetErrMsg(v string) *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	s.ErrMsg = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) SetResourceType(v string) *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	s.ResourceType = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) SetSuccess(v bool) *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	s.Success = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) SetSyncCount(v int32) *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult {
	s.SyncCount = &v
	return s
}

func (s *GetHDMLastAliyunResourceSyncResultResponseBodyDataSubResultsResourceSyncSubResult) Validate() error {
	return dara.Validate(s)
}
