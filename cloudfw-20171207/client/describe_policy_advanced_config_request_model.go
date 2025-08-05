// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyAdvancedConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribePolicyAdvancedConfigRequest
	GetLang() *string
	SetSourceIp(v string) *DescribePolicyAdvancedConfigRequest
	GetSourceIp() *string
}

type DescribePolicyAdvancedConfigRequest struct {
	// The natural language of the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	//
	// The source IP address of the request.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribePolicyAdvancedConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyAdvancedConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyAdvancedConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePolicyAdvancedConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribePolicyAdvancedConfigRequest) SetLang(v string) *DescribePolicyAdvancedConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribePolicyAdvancedConfigRequest) SetSourceIp(v string) *DescribePolicyAdvancedConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePolicyAdvancedConfigRequest) Validate() error {
	return dara.Validate(s)
}
