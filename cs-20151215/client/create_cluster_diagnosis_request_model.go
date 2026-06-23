// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTarget(v map[string]interface{}) *CreateClusterDiagnosisRequest
	GetTarget() map[string]interface{}
	SetType(v string) *CreateClusterDiagnosisRequest
	GetType() *string
}

type CreateClusterDiagnosisRequest struct {
	// The parameters that specify the diagnosis target. The following examples show the parameters for different diagnosis types:
	//
	// node:
	//
	// ```
	//
	// {"name": "cn-shanghai.10.10.10.107"}
	//
	// ```
	//
	// pod
	//
	// ```
	//
	// {"namespace": "kube-system", "name": "csi-plugin-2cg9f"}
	//
	// ```
	//
	// network
	//
	// ```
	//
	// {"src": "10.10.10.108", "dst": "10.11.247.16", "dport": "80"}
	//
	// ```
	//
	// ingress
	//
	// ```
	//
	// {"url": "https://example.com"}
	//
	// ```
	//
	// memory
	//
	// ```
	//
	// {"node":"ap-southeast-1.172.16.9.240"}
	//
	// ```
	//
	// service
	//
	// ```
	//
	// {"namespace": "kube-system", "name": "nginx-ingress-lb"}
	//
	// ```.
	//
	// example:
	//
	// {"namespace": "kube-system", "name": "csi-plugin-2cg9f"}
	Target map[string]interface{} `json:"target,omitempty" xml:"target,omitempty"`
	// The diagnosis type.
	//
	// example:
	//
	// node
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateClusterDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterDiagnosisRequest) GetTarget() map[string]interface{} {
	return s.Target
}

func (s *CreateClusterDiagnosisRequest) GetType() *string {
	return s.Type
}

func (s *CreateClusterDiagnosisRequest) SetTarget(v map[string]interface{}) *CreateClusterDiagnosisRequest {
	s.Target = v
	return s
}

func (s *CreateClusterDiagnosisRequest) SetType(v string) *CreateClusterDiagnosisRequest {
	s.Type = &v
	return s
}

func (s *CreateClusterDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}
