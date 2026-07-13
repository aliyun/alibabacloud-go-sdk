// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestModelProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TestModelProviderResponseBody
	GetCode() *string
	SetData(v *TestModelProviderResponseBodyData) *TestModelProviderResponseBody
	GetData() *TestModelProviderResponseBodyData
	SetMessage(v string) *TestModelProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *TestModelProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TestModelProviderResponseBody
	GetSuccess() *bool
}

type TestModelProviderResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TestModelProviderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TestModelProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestModelProviderResponseBody) GoString() string {
	return s.String()
}

func (s *TestModelProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *TestModelProviderResponseBody) GetData() *TestModelProviderResponseBodyData {
	return s.Data
}

func (s *TestModelProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TestModelProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestModelProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TestModelProviderResponseBody) SetCode(v string) *TestModelProviderResponseBody {
	s.Code = &v
	return s
}

func (s *TestModelProviderResponseBody) SetData(v *TestModelProviderResponseBodyData) *TestModelProviderResponseBody {
	s.Data = v
	return s
}

func (s *TestModelProviderResponseBody) SetMessage(v string) *TestModelProviderResponseBody {
	s.Message = &v
	return s
}

func (s *TestModelProviderResponseBody) SetRequestId(v string) *TestModelProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestModelProviderResponseBody) SetSuccess(v bool) *TestModelProviderResponseBody {
	s.Success = &v
	return s
}

func (s *TestModelProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TestModelProviderResponseBodyData struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	InputTokens  *int64  `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	LatencyMs    *int64  `json:"LatencyMs,omitempty" xml:"LatencyMs,omitempty"`
	OutputTokens *int64  `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	Response     *string `json:"Response,omitempty" xml:"Response,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TestModelProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TestModelProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *TestModelProviderResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *TestModelProviderResponseBodyData) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *TestModelProviderResponseBodyData) GetLatencyMs() *int64 {
	return s.LatencyMs
}

func (s *TestModelProviderResponseBodyData) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *TestModelProviderResponseBodyData) GetResponse() *string {
	return s.Response
}

func (s *TestModelProviderResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *TestModelProviderResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *TestModelProviderResponseBodyData) SetErrorMessage(v string) *TestModelProviderResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *TestModelProviderResponseBodyData) SetInputTokens(v int64) *TestModelProviderResponseBodyData {
	s.InputTokens = &v
	return s
}

func (s *TestModelProviderResponseBodyData) SetLatencyMs(v int64) *TestModelProviderResponseBodyData {
	s.LatencyMs = &v
	return s
}

func (s *TestModelProviderResponseBodyData) SetOutputTokens(v int64) *TestModelProviderResponseBodyData {
	s.OutputTokens = &v
	return s
}

func (s *TestModelProviderResponseBodyData) SetResponse(v string) *TestModelProviderResponseBodyData {
	s.Response = &v
	return s
}

func (s *TestModelProviderResponseBodyData) SetStatus(v string) *TestModelProviderResponseBodyData {
	s.Status = &v
	return s
}

func (s *TestModelProviderResponseBodyData) SetSuccess(v bool) *TestModelProviderResponseBodyData {
	s.Success = &v
	return s
}

func (s *TestModelProviderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
