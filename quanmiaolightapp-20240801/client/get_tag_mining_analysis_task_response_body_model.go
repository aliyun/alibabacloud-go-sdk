// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTagMiningAnalysisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTagMiningAnalysisTaskResponseBody
	GetCode() *string
	SetData(v *GetTagMiningAnalysisTaskResponseBodyData) *GetTagMiningAnalysisTaskResponseBody
	GetData() *GetTagMiningAnalysisTaskResponseBodyData
	SetHttpStatusCode(v string) *GetTagMiningAnalysisTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *GetTagMiningAnalysisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTagMiningAnalysisTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTagMiningAnalysisTaskResponseBody
	GetSuccess() *bool
}

type GetTagMiningAnalysisTaskResponseBody struct {
	// example:
	//
	// successful
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetTagMiningAnalysisTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// DataNotExists
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// requestId
	//
	// example:
	//
	// 085BE2D2-BB7E-59A6-B688-F2CB32124E7F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTagMiningAnalysisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTagMiningAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTagMiningAnalysisTaskResponseBody) GetData() *GetTagMiningAnalysisTaskResponseBodyData {
	return s.Data
}

func (s *GetTagMiningAnalysisTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetTagMiningAnalysisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTagMiningAnalysisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTagMiningAnalysisTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTagMiningAnalysisTaskResponseBody) SetCode(v string) *GetTagMiningAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBody) SetData(v *GetTagMiningAnalysisTaskResponseBodyData) *GetTagMiningAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBody) SetHttpStatusCode(v string) *GetTagMiningAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBody) SetMessage(v string) *GetTagMiningAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBody) SetRequestId(v string) *GetTagMiningAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBody) SetSuccess(v bool) *GetTagMiningAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTagMiningAnalysisTaskResponseBodyData struct {
	ErrorCode    *string                                            `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                                            `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Results      []*GetTagMiningAnalysisTaskResponseBodyDataResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	// example:
	//
	// RUNNIN
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetTagMiningAnalysisTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTagMiningAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTagMiningAnalysisTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTagMiningAnalysisTaskResponseBodyData) GetResults() []*GetTagMiningAnalysisTaskResponseBodyDataResults {
	return s.Results
}

func (s *GetTagMiningAnalysisTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetTagMiningAnalysisTaskResponseBodyData) SetErrorCode(v string) *GetTagMiningAnalysisTaskResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyData) SetErrorMessage(v string) *GetTagMiningAnalysisTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyData) SetResults(v []*GetTagMiningAnalysisTaskResponseBodyDataResults) *GetTagMiningAnalysisTaskResponseBodyData {
	s.Results = v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyData) SetStatus(v string) *GetTagMiningAnalysisTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyData) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTagMiningAnalysisTaskResponseBodyDataResults struct {
	// example:
	//
	// 1
	CustomId *string                                                 `json:"customId,omitempty" xml:"customId,omitempty"`
	Header   *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload  *GetTagMiningAnalysisTaskResponseBodyDataResultsPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResults) GetCustomId() *string {
	return s.CustomId
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResults) GetHeader() *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader {
	return s.Header
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResults) GetPayload() *GetTagMiningAnalysisTaskResponseBodyDataResultsPayload {
	return s.Payload
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResults) SetCustomId(v string) *GetTagMiningAnalysisTaskResponseBodyDataResults {
	s.CustomId = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResults) SetHeader(v *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) *GetTagMiningAnalysisTaskResponseBodyDataResults {
	s.Header = v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResults) SetPayload(v *GetTagMiningAnalysisTaskResponseBodyDataResultsPayload) *GetTagMiningAnalysisTaskResponseBodyDataResults {
	s.Payload = v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResults) Validate() error {
	if s.Header != nil {
		if err := s.Header.Validate(); err != nil {
			return err
		}
	}
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTagMiningAnalysisTaskResponseBodyDataResultsHeader struct {
	// example:
	//
	// DataNotExists
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// task-finished
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// 085BE2D2-BB7E-59A6-B688-F2CB32124E7F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) String() string {
	return dara.Prettify(s)
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) GetEvent() *string {
	return s.Event
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) SetErrorCode(v string) *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader {
	s.ErrorCode = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) SetErrorMessage(v string) *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader {
	s.ErrorMessage = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) SetEvent(v string) *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader {
	s.Event = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) SetRequestId(v string) *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader {
	s.RequestId = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsHeader) Validate() error {
	return dara.Validate(s)
}

type GetTagMiningAnalysisTaskResponseBodyDataResultsPayload struct {
	Output *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResultsPayload) String() string {
	return dara.Prettify(s)
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResultsPayload) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayload) GetOutput() *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput {
	return s.Output
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayload) GetUsage() *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage {
	return s.Usage
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayload) SetOutput(v *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput) *GetTagMiningAnalysisTaskResponseBodyDataResultsPayload {
	s.Output = v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayload) SetUsage(v *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage) *GetTagMiningAnalysisTaskResponseBodyDataResultsPayload {
	s.Usage = v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayload) Validate() error {
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput struct {
	// example:
	//
	// xxxx
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput) GetText() *string {
	return s.Text
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput) SetText(v string) *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput {
	s.Text = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage struct {
	// example:
	//
	// 100
	InputToken *int64 `json:"inputToken,omitempty" xml:"inputToken,omitempty"`
	// example:
	//
	// 200
	OutputToken *int64 `json:"outputToken,omitempty" xml:"outputToken,omitempty"`
	// example:
	//
	// 300
	TotalToken *int64 `json:"totalToken,omitempty" xml:"totalToken,omitempty"`
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage) GetInputToken() *int64 {
	return s.InputToken
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage) GetOutputToken() *int64 {
	return s.OutputToken
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage) GetTotalToken() *int64 {
	return s.TotalToken
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage) SetInputToken(v int64) *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage {
	s.InputToken = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage) SetOutputToken(v int64) *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage {
	s.OutputToken = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage) SetTotalToken(v int64) *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage {
	s.TotalToken = &v
	return s
}

func (s *GetTagMiningAnalysisTaskResponseBodyDataResultsPayloadUsage) Validate() error {
	return dara.Validate(s)
}
