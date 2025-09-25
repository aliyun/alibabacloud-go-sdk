// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenDocQaResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GenDocQaResultResponseBody
	GetCost() *int64
	SetData(v *GenDocQaResultResponseBodyData) *GenDocQaResultResponseBody
	GetData() *GenDocQaResultResponseBodyData
	SetDataType(v string) *GenDocQaResultResponseBody
	GetDataType() *string
	SetErrCode(v string) *GenDocQaResultResponseBody
	GetErrCode() *string
	SetMessage(v string) *GenDocQaResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenDocQaResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenDocQaResultResponseBody
	GetSuccess() *bool
	SetTime(v string) *GenDocQaResultResponseBody
	GetTime() *string
}

type GenDocQaResultResponseBody struct {
	// example:
	//
	// null
	Cost *int64                          `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GenDocQaResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 44BD277A-87F9-5310-8D63-3E6645F1DA85
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GenDocQaResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenDocQaResultResponseBody) GoString() string {
	return s.String()
}

func (s *GenDocQaResultResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GenDocQaResultResponseBody) GetData() *GenDocQaResultResponseBodyData {
	return s.Data
}

func (s *GenDocQaResultResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GenDocQaResultResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GenDocQaResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenDocQaResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenDocQaResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenDocQaResultResponseBody) GetTime() *string {
	return s.Time
}

func (s *GenDocQaResultResponseBody) SetCost(v int64) *GenDocQaResultResponseBody {
	s.Cost = &v
	return s
}

func (s *GenDocQaResultResponseBody) SetData(v *GenDocQaResultResponseBodyData) *GenDocQaResultResponseBody {
	s.Data = v
	return s
}

func (s *GenDocQaResultResponseBody) SetDataType(v string) *GenDocQaResultResponseBody {
	s.DataType = &v
	return s
}

func (s *GenDocQaResultResponseBody) SetErrCode(v string) *GenDocQaResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GenDocQaResultResponseBody) SetMessage(v string) *GenDocQaResultResponseBody {
	s.Message = &v
	return s
}

func (s *GenDocQaResultResponseBody) SetRequestId(v string) *GenDocQaResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenDocQaResultResponseBody) SetSuccess(v bool) *GenDocQaResultResponseBody {
	s.Success = &v
	return s
}

func (s *GenDocQaResultResponseBody) SetTime(v string) *GenDocQaResultResponseBody {
	s.Time = &v
	return s
}

func (s *GenDocQaResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GenDocQaResultResponseBodyData struct {
	// example:
	//
	// PROCESSING
	CurrentStatus *string `json:"currentStatus,omitempty" xml:"currentStatus,omitempty"`
	// example:
	//
	// 873648346573245
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// 7wxwrjpabj
	LibraryId      *string                                         `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	ParseQaResults []*GenDocQaResultResponseBodyDataParseQaResults `json:"parseQaResults,omitempty" xml:"parseQaResults,omitempty" type:"Repeated"`
}

func (s GenDocQaResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenDocQaResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenDocQaResultResponseBodyData) GetCurrentStatus() *string {
	return s.CurrentStatus
}

func (s *GenDocQaResultResponseBodyData) GetDocId() *string {
	return s.DocId
}

func (s *GenDocQaResultResponseBodyData) GetLibraryId() *string {
	return s.LibraryId
}

func (s *GenDocQaResultResponseBodyData) GetParseQaResults() []*GenDocQaResultResponseBodyDataParseQaResults {
	return s.ParseQaResults
}

func (s *GenDocQaResultResponseBodyData) SetCurrentStatus(v string) *GenDocQaResultResponseBodyData {
	s.CurrentStatus = &v
	return s
}

func (s *GenDocQaResultResponseBodyData) SetDocId(v string) *GenDocQaResultResponseBodyData {
	s.DocId = &v
	return s
}

func (s *GenDocQaResultResponseBodyData) SetLibraryId(v string) *GenDocQaResultResponseBodyData {
	s.LibraryId = &v
	return s
}

func (s *GenDocQaResultResponseBodyData) SetParseQaResults(v []*GenDocQaResultResponseBodyDataParseQaResults) *GenDocQaResultResponseBodyData {
	s.ParseQaResults = v
	return s
}

func (s *GenDocQaResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GenDocQaResultResponseBodyDataParseQaResults struct {
	Answer   *string `json:"answer,omitempty" xml:"answer,omitempty"`
	Question *string `json:"question,omitempty" xml:"question,omitempty"`
}

func (s GenDocQaResultResponseBodyDataParseQaResults) String() string {
	return dara.Prettify(s)
}

func (s GenDocQaResultResponseBodyDataParseQaResults) GoString() string {
	return s.String()
}

func (s *GenDocQaResultResponseBodyDataParseQaResults) GetAnswer() *string {
	return s.Answer
}

func (s *GenDocQaResultResponseBodyDataParseQaResults) GetQuestion() *string {
	return s.Question
}

func (s *GenDocQaResultResponseBodyDataParseQaResults) SetAnswer(v string) *GenDocQaResultResponseBodyDataParseQaResults {
	s.Answer = &v
	return s
}

func (s *GenDocQaResultResponseBodyDataParseQaResults) SetQuestion(v string) *GenDocQaResultResponseBodyDataParseQaResults {
	s.Question = &v
	return s
}

func (s *GenDocQaResultResponseBodyDataParseQaResults) Validate() error {
	return dara.Validate(s)
}
