// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceInstanceDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnosis(v *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) *DescribeServiceInstanceDiagnosisResponseBody
	GetDiagnosis() *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis
	SetRequestId(v string) *DescribeServiceInstanceDiagnosisResponseBody
	GetRequestId() *string
}

type DescribeServiceInstanceDiagnosisResponseBody struct {
	// The diagnostics information.
	Diagnosis *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis `json:"Diagnosis,omitempty" xml:"Diagnosis,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceInstanceDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceInstanceDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceInstanceDiagnosisResponseBody) GetDiagnosis() *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis {
	return s.Diagnosis
}

func (s *DescribeServiceInstanceDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceInstanceDiagnosisResponseBody) SetDiagnosis(v *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) *DescribeServiceInstanceDiagnosisResponseBody {
	s.Diagnosis = v
	return s
}

func (s *DescribeServiceInstanceDiagnosisResponseBody) SetRequestId(v string) *DescribeServiceInstanceDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceInstanceDiagnosisResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceInstanceDiagnosisResponseBodyDiagnosis struct {
	// The solutions to the errors.
	Advices []*string `json:"Advices,omitempty" xml:"Advices,omitempty" type:"Repeated"`
	// The causes of the errors.
	Causes []*string `json:"Causes,omitempty" xml:"Causes,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// Container worker0 failed to pull image.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
}

func (s DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) GoString() string {
	return s.String()
}

func (s *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) GetAdvices() []*string {
	return s.Advices
}

func (s *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) GetCauses() []*string {
	return s.Causes
}

func (s *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) GetError() *string {
	return s.Error
}

func (s *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) SetAdvices(v []*string) *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis {
	s.Advices = v
	return s
}

func (s *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) SetCauses(v []*string) *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis {
	s.Causes = v
	return s
}

func (s *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) SetError(v string) *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis {
	s.Error = &v
	return s
}

func (s *DescribeServiceInstanceDiagnosisResponseBodyDiagnosis) Validate() error {
	return dara.Validate(s)
}
