// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDynamicDictRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArgKeywords(v string) *GenerateDynamicDictRequest
	GetArgKeywords() *string
	SetDomains(v string) *GenerateDynamicDictRequest
	GetDomains() *string
	SetNames(v string) *GenerateDynamicDictRequest
	GetNames() *string
	SetSourceIp(v string) *GenerateDynamicDictRequest
	GetSourceIp() *string
}

type GenerateDynamicDictRequest struct {
	// The keyword of the dictionary.
	//
	// example:
	//
	// keyword
	ArgKeywords *string `json:"ArgKeywords,omitempty" xml:"ArgKeywords,omitempty"`
	// The domain name for custom weak passwords.
	//
	// example:
	//
	// https://www.aliyun.com
	Domains *string `json:"Domains,omitempty" xml:"Domains,omitempty"`
	// The company name for custom weak passwords.
	//
	// example:
	//
	// Alibaba
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 58.248.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s GenerateDynamicDictRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateDynamicDictRequest) GoString() string {
	return s.String()
}

func (s *GenerateDynamicDictRequest) GetArgKeywords() *string {
	return s.ArgKeywords
}

func (s *GenerateDynamicDictRequest) GetDomains() *string {
	return s.Domains
}

func (s *GenerateDynamicDictRequest) GetNames() *string {
	return s.Names
}

func (s *GenerateDynamicDictRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *GenerateDynamicDictRequest) SetArgKeywords(v string) *GenerateDynamicDictRequest {
	s.ArgKeywords = &v
	return s
}

func (s *GenerateDynamicDictRequest) SetDomains(v string) *GenerateDynamicDictRequest {
	s.Domains = &v
	return s
}

func (s *GenerateDynamicDictRequest) SetNames(v string) *GenerateDynamicDictRequest {
	s.Names = &v
	return s
}

func (s *GenerateDynamicDictRequest) SetSourceIp(v string) *GenerateDynamicDictRequest {
	s.SourceIp = &v
	return s
}

func (s *GenerateDynamicDictRequest) Validate() error {
	return dara.Validate(s)
}
