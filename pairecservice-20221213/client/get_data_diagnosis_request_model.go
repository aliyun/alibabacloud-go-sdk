// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetDataDiagnosisRequest
	GetInstanceId() *string
}

type GetDataDiagnosisRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// learn-pairec-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDataDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *GetDataDiagnosisRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDataDiagnosisRequest) SetInstanceId(v string) *GetDataDiagnosisRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDataDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}
