// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetDcdnWafDomainConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIpTag(v string) *BatchSetDcdnWafDomainConfigsRequest
	GetClientIpTag() *string
	SetDefenseStatus(v string) *BatchSetDcdnWafDomainConfigsRequest
	GetDefenseStatus() *string
	SetDomainNames(v string) *BatchSetDcdnWafDomainConfigsRequest
	GetDomainNames() *string
}

type BatchSetDcdnWafDomainConfigsRequest struct {
	// Specifies the header that records the IP address to be obtained. If the default header is selected, the value of this parameter is empty. If a custom header is selected, the value of this parameter is the value specified by the user. Separate multiple values with commas (,). You can specify a maximum of five values.
	//
	// example:
	//
	// X-Forwarded-For
	ClientIpTag *string `json:"ClientIpTag,omitempty" xml:"ClientIpTag,omitempty"`
	// The protection status of the domain name. Valid values: on, off, and empty string.
	//
	// 	- When you add a domain name, the value of this parameter is **on**, and the value of ClientIpTag takes effect, which is empty if the default header is selected and is the value specified by the user if a custom header is selected.
	//
	// 	- When you delete a domain name, the value of this parameter is **off**, and the value of ClientIpTag does not take effect.
	//
	// 	- When you only modify the value of ClientIpTag, the value of DefenseStatus is an empty string.
	//
	// example:
	//
	// on
	DefenseStatus *string `json:"DefenseStatus,omitempty" xml:"DefenseStatus,omitempty"`
	// The protected domain names for which you want to change the protection status. You can specify up to 50 domain names. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com,example2.com
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
}

func (s BatchSetDcdnWafDomainConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSetDcdnWafDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnWafDomainConfigsRequest) GetClientIpTag() *string {
	return s.ClientIpTag
}

func (s *BatchSetDcdnWafDomainConfigsRequest) GetDefenseStatus() *string {
	return s.DefenseStatus
}

func (s *BatchSetDcdnWafDomainConfigsRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchSetDcdnWafDomainConfigsRequest) SetClientIpTag(v string) *BatchSetDcdnWafDomainConfigsRequest {
	s.ClientIpTag = &v
	return s
}

func (s *BatchSetDcdnWafDomainConfigsRequest) SetDefenseStatus(v string) *BatchSetDcdnWafDomainConfigsRequest {
	s.DefenseStatus = &v
	return s
}

func (s *BatchSetDcdnWafDomainConfigsRequest) SetDomainNames(v string) *BatchSetDcdnWafDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchSetDcdnWafDomainConfigsRequest) Validate() error {
	return dara.Validate(s)
}
