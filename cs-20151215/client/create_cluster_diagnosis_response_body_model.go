// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateClusterDiagnosisResponseBody
	GetClusterId() *string
	SetDiagnosisId(v string) *CreateClusterDiagnosisResponseBody
	GetDiagnosisId() *string
	SetRequestId(v string) *CreateClusterDiagnosisResponseBody
	GetRequestId() *string
}

type CreateClusterDiagnosisResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// c5cdf7e3938bc4f8eb0e44b21a80f****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The diagnostic ID.
	//
	// example:
	//
	// 6f719f23098240818eb26fe3a37d****
	DiagnosisId *string `json:"diagnosis_id,omitempty" xml:"diagnosis_id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 687C5BAA-D103-4993-884B-C35E4314****
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s CreateClusterDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterDiagnosisResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateClusterDiagnosisResponseBody) GetDiagnosisId() *string {
	return s.DiagnosisId
}

func (s *CreateClusterDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClusterDiagnosisResponseBody) SetClusterId(v string) *CreateClusterDiagnosisResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterDiagnosisResponseBody) SetDiagnosisId(v string) *CreateClusterDiagnosisResponseBody {
	s.DiagnosisId = &v
	return s
}

func (s *CreateClusterDiagnosisResponseBody) SetRequestId(v string) *CreateClusterDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterDiagnosisResponseBody) Validate() error {
	return dara.Validate(s)
}
