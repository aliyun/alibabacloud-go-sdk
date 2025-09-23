// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainAccessModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainList(v []*string) *DescribeDomainAccessModeRequest
	GetDomainList() []*string
	SetSourceIp(v string) *DescribeDomainAccessModeRequest
	GetSourceIp() *string
}

type DescribeDomainAccessModeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	DomainList []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainAccessModeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAccessModeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainAccessModeRequest) GetDomainList() []*string {
	return s.DomainList
}

func (s *DescribeDomainAccessModeRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDomainAccessModeRequest) SetDomainList(v []*string) *DescribeDomainAccessModeRequest {
	s.DomainList = v
	return s
}

func (s *DescribeDomainAccessModeRequest) SetSourceIp(v string) *DescribeDomainAccessModeRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainAccessModeRequest) Validate() error {
	return dara.Validate(s)
}
