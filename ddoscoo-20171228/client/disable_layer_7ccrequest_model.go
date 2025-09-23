// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableLayer7CCRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DisableLayer7CCRequest
	GetDomain() *string
	SetResourceGroupId(v string) *DisableLayer7CCRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DisableLayer7CCRequest
	GetSourceIp() *string
}

type DisableLayer7CCRequest struct {
	// This parameter is required.
	//
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

func (s DisableLayer7CCRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableLayer7CCRequest) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCRequest) GetDomain() *string {
	return s.Domain
}

func (s *DisableLayer7CCRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DisableLayer7CCRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DisableLayer7CCRequest) SetDomain(v string) *DisableLayer7CCRequest {
	s.Domain = &v
	return s
}

func (s *DisableLayer7CCRequest) SetResourceGroupId(v string) *DisableLayer7CCRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DisableLayer7CCRequest) SetSourceIp(v string) *DisableLayer7CCRequest {
	s.SourceIp = &v
	return s
}

func (s *DisableLayer7CCRequest) Validate() error {
	return dara.Validate(s)
}
