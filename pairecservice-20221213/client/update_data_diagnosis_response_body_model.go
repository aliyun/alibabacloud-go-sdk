// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataDiagnosisResponseBody
	GetRequestId() *string
}

type UpdateDataDiagnosisResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDataDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataDiagnosisResponseBody) SetRequestId(v string) *UpdateDataDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataDiagnosisResponseBody) Validate() error {
	return dara.Validate(s)
}
