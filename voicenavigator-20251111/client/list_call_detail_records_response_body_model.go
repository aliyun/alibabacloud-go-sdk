// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallDetailRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCallDetailRecordsResponseBody
	GetCode() *string
	SetData(v *ListCallDetailRecordsResponseBodyData) *ListCallDetailRecordsResponseBody
	GetData() *ListCallDetailRecordsResponseBodyData
	SetHttpStatusCode(v int32) *ListCallDetailRecordsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCallDetailRecordsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListCallDetailRecordsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListCallDetailRecordsResponseBody
	GetRequestId() *string
}

type ListCallDetailRecordsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListCallDetailRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// A8AED3C8-F57B-5D71-9A34-4A170287533F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCallDetailRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCallDetailRecordsResponseBody) GetData() *ListCallDetailRecordsResponseBodyData {
	return s.Data
}

func (s *ListCallDetailRecordsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCallDetailRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCallDetailRecordsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListCallDetailRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCallDetailRecordsResponseBody) SetCode(v string) *ListCallDetailRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCallDetailRecordsResponseBody) SetData(v *ListCallDetailRecordsResponseBodyData) *ListCallDetailRecordsResponseBody {
	s.Data = v
	return s
}

func (s *ListCallDetailRecordsResponseBody) SetHttpStatusCode(v int32) *ListCallDetailRecordsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCallDetailRecordsResponseBody) SetMessage(v string) *ListCallDetailRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCallDetailRecordsResponseBody) SetParams(v []*string) *ListCallDetailRecordsResponseBody {
	s.Params = v
	return s
}

func (s *ListCallDetailRecordsResponseBody) SetRequestId(v string) *ListCallDetailRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCallDetailRecordsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCallDetailRecordsResponseBodyData struct {
	CallDetailRecords []*ListCallDetailRecordsResponseBodyDataCallDetailRecords `json:"CallDetailRecords,omitempty" xml:"CallDetailRecords,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 362
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCallDetailRecordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsResponseBodyData) GetCallDetailRecords() []*ListCallDetailRecordsResponseBodyDataCallDetailRecords {
	return s.CallDetailRecords
}

func (s *ListCallDetailRecordsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCallDetailRecordsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCallDetailRecordsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCallDetailRecordsResponseBodyData) SetCallDetailRecords(v []*ListCallDetailRecordsResponseBodyDataCallDetailRecords) *ListCallDetailRecordsResponseBodyData {
	s.CallDetailRecords = v
	return s
}

func (s *ListCallDetailRecordsResponseBodyData) SetPageNumber(v int32) *ListCallDetailRecordsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyData) SetPageSize(v int32) *ListCallDetailRecordsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyData) SetTotalCount(v int32) *ListCallDetailRecordsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyData) Validate() error {
	if s.CallDetailRecords != nil {
		for _, item := range s.CallDetailRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCallDetailRecordsResponseBodyDataCallDetailRecords struct {
	// example:
	//
	// 15100000001
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// example:
	//
	// 4001027277
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// example:
	//
	// Answered
	DispositionCode *string `json:"DispositionCode,omitempty" xml:"DispositionCode,omitempty"`
	// example:
	//
	// ConcurrentLimitExceeded
	DispositionReason *string `json:"DispositionReason,omitempty" xml:"DispositionReason,omitempty"`
	// example:
	//
	// 265
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 1767315903531
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// false
	IssueResolved *bool `json:"IssueResolved,omitempty" xml:"IssueResolved,omitempty"`
	// example:
	//
	// Customer
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// example:
	//
	// unknown
	ResolutionEvidence *string `json:"ResolutionEvidence,omitempty" xml:"ResolutionEvidence,omitempty"`
	// example:
	//
	// 99eebcba-7f8a-4a54-b0da-603ada79d8e2
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 1767315903531
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 245
	TalkTime *int32 `json:"TalkTime,omitempty" xml:"TalkTime,omitempty"`
	// example:
	//
	// 1
	TalkTurns *int32 `json:"TalkTurns,omitempty" xml:"TalkTurns,omitempty"`
	// example:
	//
	// default@testinstance1
	TransferTarget *string `json:"TransferTarget,omitempty" xml:"TransferTarget,omitempty"`
	// example:
	//
	// Agent
	TransferType *string `json:"TransferType,omitempty" xml:"TransferType,omitempty"`
}

func (s ListCallDetailRecordsResponseBodyDataCallDetailRecords) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsResponseBodyDataCallDetailRecords) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) GetCallee() *string {
	return s.Callee
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) GetCaller() *string {
	return s.Caller
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) GetDispositionCode() *string {
	return s.DispositionCode
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) GetDispositionReason() *string {
	return s.DispositionReason
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) GetDuration() *int32 {
	return s.Duration
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) GetIssueResolved() *bool {
	return s.IssueResolved
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) GetResolutionEvidence() *string {
	return s.ResolutionEvidence
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) GetSessionId() *string {
	return s.SessionId
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) GetTalkTime() *int32 {
	return s.TalkTime
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) GetTalkTurns() *int32 {
	return s.TalkTurns
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) GetTransferTarget() *string {
	return s.TransferTarget
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) GetTransferType() *string {
	return s.TransferType
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) SetCallee(v string) *ListCallDetailRecordsResponseBodyDataCallDetailRecords {
	s.Callee = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) SetCaller(v string) *ListCallDetailRecordsResponseBodyDataCallDetailRecords {
	s.Caller = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) SetDispositionCode(v string) *ListCallDetailRecordsResponseBodyDataCallDetailRecords {
	s.DispositionCode = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) SetDispositionReason(v string) *ListCallDetailRecordsResponseBodyDataCallDetailRecords {
	s.DispositionReason = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) SetDuration(v int32) *ListCallDetailRecordsResponseBodyDataCallDetailRecords {
	s.Duration = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) SetEndTime(v int64) *ListCallDetailRecordsResponseBodyDataCallDetailRecords {
	s.EndTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) SetIssueResolved(v bool) *ListCallDetailRecordsResponseBodyDataCallDetailRecords {
	s.IssueResolved = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) SetReleaseInitiator(v string) *ListCallDetailRecordsResponseBodyDataCallDetailRecords {
	s.ReleaseInitiator = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) SetResolutionEvidence(v string) *ListCallDetailRecordsResponseBodyDataCallDetailRecords {
	s.ResolutionEvidence = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) SetSessionId(v string) *ListCallDetailRecordsResponseBodyDataCallDetailRecords {
	s.SessionId = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) SetStartTime(v int64) *ListCallDetailRecordsResponseBodyDataCallDetailRecords {
	s.StartTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) SetTalkTime(v int32) *ListCallDetailRecordsResponseBodyDataCallDetailRecords {
	s.TalkTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) SetTalkTurns(v int32) *ListCallDetailRecordsResponseBodyDataCallDetailRecords {
	s.TalkTurns = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) SetTransferTarget(v string) *ListCallDetailRecordsResponseBodyDataCallDetailRecords {
	s.TransferTarget = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) SetTransferType(v string) *ListCallDetailRecordsResponseBodyDataCallDetailRecords {
	s.TransferType = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataCallDetailRecords) Validate() error {
	return dara.Validate(s)
}
