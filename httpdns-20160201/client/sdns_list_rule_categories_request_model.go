// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdnsListRuleCategoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SdnsListRuleCategoriesRequest
	GetDomainName() *string
}

type SdnsListRuleCategoriesRequest struct {
	// The domain name that is connected to HTTPDNS under the current account. This parameter is required.
	//
	// example:
	//
	// api.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s SdnsListRuleCategoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s SdnsListRuleCategoriesRequest) GoString() string {
	return s.String()
}

func (s *SdnsListRuleCategoriesRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SdnsListRuleCategoriesRequest) SetDomainName(v string) *SdnsListRuleCategoriesRequest {
	s.DomainName = &v
	return s
}

func (s *SdnsListRuleCategoriesRequest) Validate() error {
	return dara.Validate(s)
}
