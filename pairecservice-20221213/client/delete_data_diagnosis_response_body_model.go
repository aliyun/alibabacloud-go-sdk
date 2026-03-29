// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataDiagnosisResponseBody
	GetRequestId() *string
}

type DeleteDataDiagnosisResponseBody struct {
	// example:
	//
	// 59CE7EC6-F268-5D71-9215-32922CC50D72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataDiagnosisResponseBody) SetRequestId(v string) *DeleteDataDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataDiagnosisResponseBody) Validate() error {
	return dara.Validate(s)
}
