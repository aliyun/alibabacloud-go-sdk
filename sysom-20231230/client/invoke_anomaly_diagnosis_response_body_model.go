// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeAnomalyDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvokeAnomalyDiagnosisResponseBody
	GetCode() *string
	SetMessage(v string) *InvokeAnomalyDiagnosisResponseBody
	GetMessage() *string
	SetRequestId(v string) *InvokeAnomalyDiagnosisResponseBody
	GetRequestId() *string
}

type InvokeAnomalyDiagnosisResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s InvokeAnomalyDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokeAnomalyDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeAnomalyDiagnosisResponseBody) GetCode() *string {
	return s.Code
}

func (s *InvokeAnomalyDiagnosisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvokeAnomalyDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokeAnomalyDiagnosisResponseBody) SetCode(v string) *InvokeAnomalyDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *InvokeAnomalyDiagnosisResponseBody) SetMessage(v string) *InvokeAnomalyDiagnosisResponseBody {
	s.Message = &v
	return s
}

func (s *InvokeAnomalyDiagnosisResponseBody) SetRequestId(v string) *InvokeAnomalyDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeAnomalyDiagnosisResponseBody) Validate() error {
	return dara.Validate(s)
}
