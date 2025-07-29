// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEssayCorrectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEssayCorrectionTaskResponseBody
	GetCode() *string
	SetData(v *GetEssayCorrectionTaskResponseBodyData) *GetEssayCorrectionTaskResponseBody
	GetData() *GetEssayCorrectionTaskResponseBodyData
	SetHttpStatusCode(v int32) *GetEssayCorrectionTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetEssayCorrectionTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEssayCorrectionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetEssayCorrectionTaskResponseBody
	GetSuccess() *bool
}

type GetEssayCorrectionTaskResponseBody struct {
	// example:
	//
	// successful
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetEssayCorrectionTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetEssayCorrectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEssayCorrectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetEssayCorrectionTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEssayCorrectionTaskResponseBody) GetData() *GetEssayCorrectionTaskResponseBodyData {
	return s.Data
}

func (s *GetEssayCorrectionTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetEssayCorrectionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEssayCorrectionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEssayCorrectionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEssayCorrectionTaskResponseBody) SetCode(v string) *GetEssayCorrectionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetEssayCorrectionTaskResponseBody) SetData(v *GetEssayCorrectionTaskResponseBodyData) *GetEssayCorrectionTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetEssayCorrectionTaskResponseBody) SetHttpStatusCode(v int32) *GetEssayCorrectionTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetEssayCorrectionTaskResponseBody) SetMessage(v string) *GetEssayCorrectionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetEssayCorrectionTaskResponseBody) SetRequestId(v string) *GetEssayCorrectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEssayCorrectionTaskResponseBody) SetSuccess(v bool) *GetEssayCorrectionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetEssayCorrectionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetEssayCorrectionTaskResponseBodyData struct {
	ErrorMessage *string                                          `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Results      []*GetEssayCorrectionTaskResponseBodyDataResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetEssayCorrectionTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEssayCorrectionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEssayCorrectionTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetEssayCorrectionTaskResponseBodyData) GetResults() []*GetEssayCorrectionTaskResponseBodyDataResults {
	return s.Results
}

func (s *GetEssayCorrectionTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetEssayCorrectionTaskResponseBodyData) SetErrorMessage(v string) *GetEssayCorrectionTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetEssayCorrectionTaskResponseBodyData) SetResults(v []*GetEssayCorrectionTaskResponseBodyDataResults) *GetEssayCorrectionTaskResponseBodyData {
	s.Results = v
	return s
}

func (s *GetEssayCorrectionTaskResponseBodyData) SetStatus(v string) *GetEssayCorrectionTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetEssayCorrectionTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetEssayCorrectionTaskResponseBodyDataResults struct {
	// xxx
	//
	// example:
	//
	// 1
	CustomId *string `json:"customId,omitempty" xml:"customId,omitempty"`
	Result   *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 58
	Score *int32 `json:"score,omitempty" xml:"score,omitempty"`
}

func (s GetEssayCorrectionTaskResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s GetEssayCorrectionTaskResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *GetEssayCorrectionTaskResponseBodyDataResults) GetCustomId() *string {
	return s.CustomId
}

func (s *GetEssayCorrectionTaskResponseBodyDataResults) GetResult() *string {
	return s.Result
}

func (s *GetEssayCorrectionTaskResponseBodyDataResults) GetScore() *int32 {
	return s.Score
}

func (s *GetEssayCorrectionTaskResponseBodyDataResults) SetCustomId(v string) *GetEssayCorrectionTaskResponseBodyDataResults {
	s.CustomId = &v
	return s
}

func (s *GetEssayCorrectionTaskResponseBodyDataResults) SetResult(v string) *GetEssayCorrectionTaskResponseBodyDataResults {
	s.Result = &v
	return s
}

func (s *GetEssayCorrectionTaskResponseBodyDataResults) SetScore(v int32) *GetEssayCorrectionTaskResponseBodyDataResults {
	s.Score = &v
	return s
}

func (s *GetEssayCorrectionTaskResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}
