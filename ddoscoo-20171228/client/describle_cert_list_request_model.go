// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribleCertListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribleCertListRequest
	GetDomain() *string
	SetResourceGroupId(v string) *DescribleCertListRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribleCertListRequest
	GetSourceIp() *string
}

type DescribleCertListRequest struct {
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribleCertListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribleCertListRequest) GoString() string {
	return s.String()
}

func (s *DescribleCertListRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribleCertListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribleCertListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribleCertListRequest) SetDomain(v string) *DescribleCertListRequest {
	s.Domain = &v
	return s
}

func (s *DescribleCertListRequest) SetResourceGroupId(v string) *DescribleCertListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribleCertListRequest) SetSourceIp(v string) *DescribleCertListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribleCertListRequest) Validate() error {
	return dara.Validate(s)
}
