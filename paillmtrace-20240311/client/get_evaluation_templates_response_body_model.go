// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEvaluationTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEvaluationTemplatesResponseBody
	GetCode() *string
	SetEvaluationTemplates(v []interface{}) *GetEvaluationTemplatesResponseBody
	GetEvaluationTemplates() []interface{}
	SetMessage(v string) *GetEvaluationTemplatesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEvaluationTemplatesResponseBody
	GetRequestId() *string
}

type GetEvaluationTemplatesResponseBody struct {
	// Internal error code. Set only when the response has an error.
	//
	// example:
	//
	// ExecutionFailure
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A series of templates used internally by the evaluation system to construct LLM interaction information.
	EvaluationTemplates []interface{} `json:"EvaluationTemplates,omitempty" xml:"EvaluationTemplates,omitempty" type:"Repeated"`
	// Response error message. Set only when the response has an error.
	//
	// example:
	//
	// cannot get data back.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEvaluationTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluationTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *GetEvaluationTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEvaluationTemplatesResponseBody) GetEvaluationTemplates() []interface{} {
	return s.EvaluationTemplates
}

func (s *GetEvaluationTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEvaluationTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEvaluationTemplatesResponseBody) SetCode(v string) *GetEvaluationTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *GetEvaluationTemplatesResponseBody) SetEvaluationTemplates(v []interface{}) *GetEvaluationTemplatesResponseBody {
	s.EvaluationTemplates = v
	return s
}

func (s *GetEvaluationTemplatesResponseBody) SetMessage(v string) *GetEvaluationTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *GetEvaluationTemplatesResponseBody) SetRequestId(v string) *GetEvaluationTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEvaluationTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}
