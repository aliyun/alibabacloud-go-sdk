// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIstioVersion(v string) *DescribeCrTemplatesRequest
	GetIstioVersion() *string
	SetKind(v string) *DescribeCrTemplatesRequest
	GetKind() *string
}

type DescribeCrTemplatesRequest struct {
	// The version of Istio used by the ASM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1.9.7.31-g24cdcb43-aliyun
	IstioVersion *string `json:"IstioVersion,omitempty" xml:"IstioVersion,omitempty"`
	// The type of Istio resource whose common YAML templates you want to query. Valid values:
	//
	// 	- AuthorizationPolicy
	//
	// 	- RequestAuthentication
	//
	// 	- PeerAuthentication
	//
	// 	- WorkloadGroup
	//
	// 	- WorkloadEntry
	//
	// 	- Sidecar
	//
	// 	- EnvoyFilter
	//
	// 	- ServiceEntry
	//
	// 	- Gateway
	//
	// 	- DestinationRule
	//
	// 	- VirtualService
	//
	// This parameter is required.
	//
	// example:
	//
	// VirtualService
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
}

func (s DescribeCrTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCrTemplatesRequest) GetIstioVersion() *string {
	return s.IstioVersion
}

func (s *DescribeCrTemplatesRequest) GetKind() *string {
	return s.Kind
}

func (s *DescribeCrTemplatesRequest) SetIstioVersion(v string) *DescribeCrTemplatesRequest {
	s.IstioVersion = &v
	return s
}

func (s *DescribeCrTemplatesRequest) SetKind(v string) *DescribeCrTemplatesRequest {
	s.Kind = &v
	return s
}

func (s *DescribeCrTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
