// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeAnomalyDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuid(v string) *InvokeAnomalyDiagnosisRequest
	GetUuid() *string
}

type InvokeAnomalyDiagnosisRequest struct {
	// example:
	//
	// 8047d763-5465-4a8c-b1cd-23f5a8ba2594
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s InvokeAnomalyDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeAnomalyDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *InvokeAnomalyDiagnosisRequest) GetUuid() *string {
	return s.Uuid
}

func (s *InvokeAnomalyDiagnosisRequest) SetUuid(v string) *InvokeAnomalyDiagnosisRequest {
	s.Uuid = &v
	return s
}

func (s *InvokeAnomalyDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}
