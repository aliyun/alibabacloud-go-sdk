// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataDiagnosisId(v string) *CreateDataDiagnosisResponseBody
	GetDataDiagnosisId() *string
	SetRequestId(v string) *CreateDataDiagnosisResponseBody
	GetRequestId() *string
}

type CreateDataDiagnosisResponseBody struct {
	// example:
	//
	// 1
	DataDiagnosisId *string `json:"DataDiagnosisId,omitempty" xml:"DataDiagnosisId,omitempty"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataDiagnosisResponseBody) GetDataDiagnosisId() *string {
	return s.DataDiagnosisId
}

func (s *CreateDataDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataDiagnosisResponseBody) SetDataDiagnosisId(v string) *CreateDataDiagnosisResponseBody {
	s.DataDiagnosisId = &v
	return s
}

func (s *CreateDataDiagnosisResponseBody) SetRequestId(v string) *CreateDataDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataDiagnosisResponseBody) Validate() error {
	return dara.Validate(s)
}
