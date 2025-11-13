// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunLegalAdviceConsultationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunLegalAdviceConsultationResponseBody
	GetCode() *string
	SetMessage(v string) *RunLegalAdviceConsultationResponseBody
	GetMessage() *string
	SetRequestId(v string) *RunLegalAdviceConsultationResponseBody
	GetRequestId() *string
	SetResponseMarkdown(v string) *RunLegalAdviceConsultationResponseBody
	GetResponseMarkdown() *string
	SetRound(v int32) *RunLegalAdviceConsultationResponseBody
	GetRound() *int32
	SetStatus(v string) *RunLegalAdviceConsultationResponseBody
	GetStatus() *string
	SetSuccess(v bool) *RunLegalAdviceConsultationResponseBody
	GetSuccess() *bool
	SetUsage(v *RunLegalAdviceConsultationResponseBodyUsage) *RunLegalAdviceConsultationResponseBody
	GetUsage() *RunLegalAdviceConsultationResponseBodyUsage
	SetContents(v string) *RunLegalAdviceConsultationResponseBody
	GetContents() *string
	SetExtra(v string) *RunLegalAdviceConsultationResponseBody
	GetExtra() *string
	SetHttpStatusCode(v string) *RunLegalAdviceConsultationResponseBody
	GetHttpStatusCode() *string
}

type RunLegalAdviceConsultationResponseBody struct {
	// example:
	//
	// Request.Signature.Error
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 744419D0-671A-5997-9840-E8AE48356194
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseMarkdown *string `json:"ResponseMarkdown,omitempty" xml:"ResponseMarkdown,omitempty"`
	// example:
	//
	// 1
	Round  *int32  `json:"Round,omitempty" xml:"Round,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// True
	Success  *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	Usage    *RunLegalAdviceConsultationResponseBodyUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
	Contents *string                                      `json:"contents,omitempty" xml:"contents,omitempty"`
	Extra    *string                                      `json:"extra,omitempty" xml:"extra,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
}

func (s RunLegalAdviceConsultationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunLegalAdviceConsultationResponseBody) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunLegalAdviceConsultationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunLegalAdviceConsultationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunLegalAdviceConsultationResponseBody) GetResponseMarkdown() *string {
	return s.ResponseMarkdown
}

func (s *RunLegalAdviceConsultationResponseBody) GetRound() *int32 {
	return s.Round
}

func (s *RunLegalAdviceConsultationResponseBody) GetStatus() *string {
	return s.Status
}

func (s *RunLegalAdviceConsultationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunLegalAdviceConsultationResponseBody) GetUsage() *RunLegalAdviceConsultationResponseBodyUsage {
	return s.Usage
}

func (s *RunLegalAdviceConsultationResponseBody) GetContents() *string {
	return s.Contents
}

func (s *RunLegalAdviceConsultationResponseBody) GetExtra() *string {
	return s.Extra
}

func (s *RunLegalAdviceConsultationResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *RunLegalAdviceConsultationResponseBody) SetCode(v string) *RunLegalAdviceConsultationResponseBody {
	s.Code = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetMessage(v string) *RunLegalAdviceConsultationResponseBody {
	s.Message = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetRequestId(v string) *RunLegalAdviceConsultationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetResponseMarkdown(v string) *RunLegalAdviceConsultationResponseBody {
	s.ResponseMarkdown = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetRound(v int32) *RunLegalAdviceConsultationResponseBody {
	s.Round = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetStatus(v string) *RunLegalAdviceConsultationResponseBody {
	s.Status = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetSuccess(v bool) *RunLegalAdviceConsultationResponseBody {
	s.Success = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetUsage(v *RunLegalAdviceConsultationResponseBodyUsage) *RunLegalAdviceConsultationResponseBody {
	s.Usage = v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetContents(v string) *RunLegalAdviceConsultationResponseBody {
	s.Contents = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetExtra(v string) *RunLegalAdviceConsultationResponseBody {
	s.Extra = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetHttpStatusCode(v string) *RunLegalAdviceConsultationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) Validate() error {
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunLegalAdviceConsultationResponseBodyUsage struct {
	// example:
	//
	// 500
	InputTokens *int32 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 700
	OutputTokens *int32 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 1200
	TotalTokens *int32 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunLegalAdviceConsultationResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s RunLegalAdviceConsultationResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationResponseBodyUsage) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *RunLegalAdviceConsultationResponseBodyUsage) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *RunLegalAdviceConsultationResponseBodyUsage) GetTotalTokens() *int32 {
	return s.TotalTokens
}

func (s *RunLegalAdviceConsultationResponseBodyUsage) SetInputTokens(v int32) *RunLegalAdviceConsultationResponseBodyUsage {
	s.InputTokens = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBodyUsage) SetOutputTokens(v int32) *RunLegalAdviceConsultationResponseBodyUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBodyUsage) SetTotalTokens(v int32) *RunLegalAdviceConsultationResponseBodyUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
