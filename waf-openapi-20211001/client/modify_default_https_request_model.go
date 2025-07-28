// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefaultHttpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *ModifyDefaultHttpsRequest
	GetCertId() *string
	SetCipherSuite(v int32) *ModifyDefaultHttpsRequest
	GetCipherSuite() *int32
	SetCustomCiphers(v []*string) *ModifyDefaultHttpsRequest
	GetCustomCiphers() []*string
	SetEnableTLSv3(v bool) *ModifyDefaultHttpsRequest
	GetEnableTLSv3() *bool
	SetInstanceId(v string) *ModifyDefaultHttpsRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyDefaultHttpsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyDefaultHttpsRequest
	GetResourceManagerResourceGroupId() *string
	SetTLSVersion(v string) *ModifyDefaultHttpsRequest
	GetTLSVersion() *string
}

type ModifyDefaultHttpsRequest struct {
	// The ID of the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123-cn-hangzhou
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The type of the cipher suites. Valid values:
	//
	// 	- **1**: all cipher suites.
	//
	// 	- **2**: strong cipher suites.
	//
	// 	- **99**: custom cipher suites.
	//
	// example:
	//
	// 0
	CipherSuite *int32 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suites that you want to add. This parameter is available only if you set **CipherSuite*	- to **99**.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// Specifies whether to support TLS 1.3. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 阿里云资源组ID。
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The version of the TLS protocol. Valid values:
	//
	// 	- **tlsv1**
	//
	// 	- **tlsv1.1**
	//
	// 	- **tlsv1.2**
	//
	// This parameter is required.
	//
	// example:
	//
	// tlsv1
	TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
}

func (s ModifyDefaultHttpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefaultHttpsRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefaultHttpsRequest) GetCertId() *string {
	return s.CertId
}

func (s *ModifyDefaultHttpsRequest) GetCipherSuite() *int32 {
	return s.CipherSuite
}

func (s *ModifyDefaultHttpsRequest) GetCustomCiphers() []*string {
	return s.CustomCiphers
}

func (s *ModifyDefaultHttpsRequest) GetEnableTLSv3() *bool {
	return s.EnableTLSv3
}

func (s *ModifyDefaultHttpsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDefaultHttpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDefaultHttpsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyDefaultHttpsRequest) GetTLSVersion() *string {
	return s.TLSVersion
}

func (s *ModifyDefaultHttpsRequest) SetCertId(v string) *ModifyDefaultHttpsRequest {
	s.CertId = &v
	return s
}

func (s *ModifyDefaultHttpsRequest) SetCipherSuite(v int32) *ModifyDefaultHttpsRequest {
	s.CipherSuite = &v
	return s
}

func (s *ModifyDefaultHttpsRequest) SetCustomCiphers(v []*string) *ModifyDefaultHttpsRequest {
	s.CustomCiphers = v
	return s
}

func (s *ModifyDefaultHttpsRequest) SetEnableTLSv3(v bool) *ModifyDefaultHttpsRequest {
	s.EnableTLSv3 = &v
	return s
}

func (s *ModifyDefaultHttpsRequest) SetInstanceId(v string) *ModifyDefaultHttpsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefaultHttpsRequest) SetRegionId(v string) *ModifyDefaultHttpsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefaultHttpsRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefaultHttpsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefaultHttpsRequest) SetTLSVersion(v string) *ModifyDefaultHttpsRequest {
	s.TLSVersion = &v
	return s
}

func (s *ModifyDefaultHttpsRequest) Validate() error {
	return dara.Validate(s)
}
