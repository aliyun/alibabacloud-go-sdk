// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteDataDiagnosisRequest
	GetInstanceId() *string
}

type DeleteDataDiagnosisRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// learn-pairec-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDataDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataDiagnosisRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDataDiagnosisRequest) SetInstanceId(v string) *DeleteDataDiagnosisRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDataDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}
