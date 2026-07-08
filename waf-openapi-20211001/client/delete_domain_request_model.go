// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *DeleteDomainRequest
	GetAccessType() *string
	SetDomain(v string) *DeleteDomainRequest
	GetDomain() *string
	SetDomainId(v string) *DeleteDomainRequest
	GetDomainId() *string
	SetInstanceId(v string) *DeleteDomainRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteDomainRequest
	GetRegionId() *string
}

type DeleteDomainRequest struct {
	// The access type of the WAF instance. Valid values:
	//
	// - **share*	- (default): CNAME access.
	//
	// - **hybrid_cloud_cname**: hybrid cloud reverse proxy access.
	//
	// example:
	//
	// share
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The domain name that has been connected to WAF.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The domain name ID.
	//
	// example:
	//
	// www.aliyundoc.com-waf
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) GetAccessType() *string {
	return s.AccessType
}

func (s *DeleteDomainRequest) GetDomain() *string {
	return s.Domain
}

func (s *DeleteDomainRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *DeleteDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDomainRequest) SetAccessType(v string) *DeleteDomainRequest {
	s.AccessType = &v
	return s
}

func (s *DeleteDomainRequest) SetDomain(v string) *DeleteDomainRequest {
	s.Domain = &v
	return s
}

func (s *DeleteDomainRequest) SetDomainId(v string) *DeleteDomainRequest {
	s.DomainId = &v
	return s
}

func (s *DeleteDomainRequest) SetInstanceId(v string) *DeleteDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDomainRequest) SetRegionId(v string) *DeleteDomainRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDomainRequest) Validate() error {
	return dara.Validate(s)
}
