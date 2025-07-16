// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnosisList(v []*DescribeServiceDiagnosisResponseBodyDiagnosisList) *DescribeServiceDiagnosisResponseBody
	GetDiagnosisList() []*DescribeServiceDiagnosisResponseBodyDiagnosisList
	SetRequestId(v string) *DescribeServiceDiagnosisResponseBody
	GetRequestId() *string
}

type DescribeServiceDiagnosisResponseBody struct {
	// The diagnostics list.
	DiagnosisList []*DescribeServiceDiagnosisResponseBodyDiagnosisList `json:"DiagnosisList,omitempty" xml:"DiagnosisList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceDiagnosisResponseBody) GetDiagnosisList() []*DescribeServiceDiagnosisResponseBodyDiagnosisList {
	return s.DiagnosisList
}

func (s *DescribeServiceDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceDiagnosisResponseBody) SetDiagnosisList(v []*DescribeServiceDiagnosisResponseBodyDiagnosisList) *DescribeServiceDiagnosisResponseBody {
	s.DiagnosisList = v
	return s
}

func (s *DescribeServiceDiagnosisResponseBody) SetRequestId(v string) *DescribeServiceDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceDiagnosisResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceDiagnosisResponseBodyDiagnosisList struct {
	// The suggestions about how to handle the errors.
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

func (s DescribeServiceDiagnosisResponseBodyDiagnosisList) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceDiagnosisResponseBodyDiagnosisList) GoString() string {
	return s.String()
}

func (s *DescribeServiceDiagnosisResponseBodyDiagnosisList) GetAdvices() []*string {
	return s.Advices
}

func (s *DescribeServiceDiagnosisResponseBodyDiagnosisList) GetCauses() []*string {
	return s.Causes
}

func (s *DescribeServiceDiagnosisResponseBodyDiagnosisList) GetError() *string {
	return s.Error
}

func (s *DescribeServiceDiagnosisResponseBodyDiagnosisList) SetAdvices(v []*string) *DescribeServiceDiagnosisResponseBodyDiagnosisList {
	s.Advices = v
	return s
}

func (s *DescribeServiceDiagnosisResponseBodyDiagnosisList) SetCauses(v []*string) *DescribeServiceDiagnosisResponseBodyDiagnosisList {
	s.Causes = v
	return s
}

func (s *DescribeServiceDiagnosisResponseBodyDiagnosisList) SetError(v string) *DescribeServiceDiagnosisResponseBodyDiagnosisList {
	s.Error = &v
	return s
}

func (s *DescribeServiceDiagnosisResponseBodyDiagnosisList) Validate() error {
	return dara.Validate(s)
}
