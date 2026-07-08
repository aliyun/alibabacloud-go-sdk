// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfig(v *DescribeVsDomainDetailResponseBodyDomainConfig) *DescribeVsDomainDetailResponseBody
	GetDomainConfig() *DescribeVsDomainDetailResponseBodyDomainConfig
	SetRequestId(v string) *DescribeVsDomainDetailResponseBody
	GetRequestId() *string
}

type DescribeVsDomainDetailResponseBody struct {
	// Domain configuration details.
	DomainConfig *DescribeVsDomainDetailResponseBodyDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9BEC5E85-C76B-56EF-A922-860EFDB8B64B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsDomainDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainDetailResponseBody) GetDomainConfig() *DescribeVsDomainDetailResponseBodyDomainConfig {
	return s.DomainConfig
}

func (s *DescribeVsDomainDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsDomainDetailResponseBody) SetDomainConfig(v *DescribeVsDomainDetailResponseBodyDomainConfig) *DescribeVsDomainDetailResponseBody {
	s.DomainConfig = v
	return s
}

func (s *DescribeVsDomainDetailResponseBody) SetRequestId(v string) *DescribeVsDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBody) Validate() error {
	if s.DomainConfig != nil {
		if err := s.DomainConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVsDomainDetailResponseBodyDomainConfig struct {
	// The CNAME assigned to the Visual Edge Computing Service domain. You must configure your DNS provider to point your domain to this CNAME.
	//
	// example:
	//
	// example.aliyundoc.com.*****.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The domain description.
	//
	// example:
	//
	// 测试使用
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Visual Edge Computing Service domain name.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The status of the Visual Edge Computing Service accelerated domain. Valid values:
	//
	// - **online**: Enabled.
	//
	// - **offline**: Disabled.
	//
	// - **configuring**: Being configured.
	//
	// example:
	//
	// online
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The domain type.
	//
	// > Static value: vs
	//
	// example:
	//
	// vs
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-07-19T10:27:23Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The most recent modification time.
	//
	// example:
	//
	// 2021-07-19T10:27:23Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The region where the domain is located.
	//
	// example:
	//
	// cn-qingdao
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Indicates whether HTTPS is enabled. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off*	- (default): Disabled.
	//
	// example:
	//
	// off
	SSLProtocol *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	// The acceleration region. Valid values:
	//
	// - **domestic**
	//
	// - **overseas**
	//
	// - **global**
	//
	// example:
	//
	// domestic
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s DescribeVsDomainDetailResponseBodyDomainConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainDetailResponseBodyDomainConfig) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) GetCname() *string {
	return s.Cname
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) GetDescription() *string {
	return s.Description
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) GetDomainType() *string {
	return s.DomainType
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) GetRegion() *string {
	return s.Region
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) GetScope() *string {
	return s.Scope
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetCname(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.Cname = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetDescription(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.Description = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetDomainName(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetDomainStatus(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.DomainStatus = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetDomainType(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.DomainType = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetGmtCreated(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.GmtCreated = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetGmtModified(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.GmtModified = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetRegion(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.Region = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetSSLProtocol(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) SetScope(v string) *DescribeVsDomainDetailResponseBodyDomainConfig {
	s.Scope = &v
	return s
}

func (s *DescribeVsDomainDetailResponseBodyDomainConfig) Validate() error {
	return dara.Validate(s)
}
