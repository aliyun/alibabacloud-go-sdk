// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpgradeVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUpgradeVersionResponseBody
	GetRequestId() *string
	SetVersion(v *DescribeUpgradeVersionResponseBodyVersion) *DescribeUpgradeVersionResponseBody
	GetVersion() *DescribeUpgradeVersionResponseBodyVersion
}

type DescribeUpgradeVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 11fd0027-c27e-41bb-a565-75583054****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The version information.
	Version *DescribeUpgradeVersionResponseBodyVersion `json:"Version,omitempty" xml:"Version,omitempty" type:"Struct"`
}

func (s DescribeUpgradeVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUpgradeVersionResponseBody) GetVersion() *DescribeUpgradeVersionResponseBodyVersion {
	return s.Version
}

func (s *DescribeUpgradeVersionResponseBody) SetRequestId(v string) *DescribeUpgradeVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUpgradeVersionResponseBody) SetVersion(v *DescribeUpgradeVersionResponseBodyVersion) *DescribeUpgradeVersionResponseBody {
	s.Version = v
	return s
}

func (s *DescribeUpgradeVersionResponseBody) Validate() error {
	if s.Version != nil {
		if err := s.Version.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUpgradeVersionResponseBodyVersion struct {
	// The version of the ASM instance.
	//
	// example:
	//
	// v1.17.2.42-gf7619883-aliyun
	IstioOperatorVersion *string `json:"IstioOperatorVersion,omitempty" xml:"IstioOperatorVersion,omitempty"`
	// The Istio version.
	//
	// example:
	//
	// 1.17.2
	IstioVersion *string `json:"IstioVersion,omitempty" xml:"IstioVersion,omitempty"`
	// The Kubernetes version.
	//
	// example:
	//
	// v1.24.6-aliyun.1
	KubernetesVersion *string `json:"KubernetesVersion,omitempty" xml:"KubernetesVersion,omitempty"`
}

func (s DescribeUpgradeVersionResponseBodyVersion) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeVersionResponseBodyVersion) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeVersionResponseBodyVersion) GetIstioOperatorVersion() *string {
	return s.IstioOperatorVersion
}

func (s *DescribeUpgradeVersionResponseBodyVersion) GetIstioVersion() *string {
	return s.IstioVersion
}

func (s *DescribeUpgradeVersionResponseBodyVersion) GetKubernetesVersion() *string {
	return s.KubernetesVersion
}

func (s *DescribeUpgradeVersionResponseBodyVersion) SetIstioOperatorVersion(v string) *DescribeUpgradeVersionResponseBodyVersion {
	s.IstioOperatorVersion = &v
	return s
}

func (s *DescribeUpgradeVersionResponseBodyVersion) SetIstioVersion(v string) *DescribeUpgradeVersionResponseBodyVersion {
	s.IstioVersion = &v
	return s
}

func (s *DescribeUpgradeVersionResponseBodyVersion) SetKubernetesVersion(v string) *DescribeUpgradeVersionResponseBodyVersion {
	s.KubernetesVersion = &v
	return s
}

func (s *DescribeUpgradeVersionResponseBodyVersion) Validate() error {
	return dara.Validate(s)
}
