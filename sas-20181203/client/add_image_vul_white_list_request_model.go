// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageVulWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *AddImageVulWhiteListRequest
	GetLang() *string
	SetReason(v string) *AddImageVulWhiteListRequest
	GetReason() *string
	SetSource(v string) *AddImageVulWhiteListRequest
	GetSource() *string
	SetTarget(v string) *AddImageVulWhiteListRequest
	GetTarget() *string
	SetType(v string) *AddImageVulWhiteListRequest
	GetType() *string
	SetWhitelist(v string) *AddImageVulWhiteListRequest
	GetWhitelist() *string
}

type AddImageVulWhiteListRequest struct {
	// The language of the content within the request and response. Default value: zh. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The reason why you add the vulnerability to the whitelist.
	//
	// example:
	//
	// already config in another way
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The source of the whitelist. Valid values:
	//
	// - **image**
	//
	// - **agentless**
	//
	// example:
	//
	// image
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The object on which you want to perform the operation. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **type**: the object type. The value is fixed to repo.
	//
	// 	- **target**: the object content. The value is in the Namespace/Image repository format.
	//
	// example:
	//
	// {\\"type\\":\\"repo\\",\\"target\\":[\\"sas_test/script_0209\\",\\"sas_test/script\\"]}
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: system vulnerability
	//
	// 	- **sca**: application vulnerability
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The whitelist. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **Type**: the vulnerability type. Valid values: cve and sca.
	//
	// 	- **Name**: the name of the vulnerability that is specified in Common Vulnerabilities and Exposures (CVE).
	//
	// 	- **AliasName**: the alias of the vulnerability that is specified in CVE.
	//
	// example:
	//
	// [{\\"Type\\":\\"sca\\",\\"Name\\":\\"imgsca:java:spring-core:AVD-2022-1124599\\",\\"AliasName\\":\\"Spring Framework JDK >= 9 (CVE-2022-22965)\\"}]
	Whitelist *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
}

func (s AddImageVulWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s AddImageVulWhiteListRequest) GoString() string {
	return s.String()
}

func (s *AddImageVulWhiteListRequest) GetLang() *string {
	return s.Lang
}

func (s *AddImageVulWhiteListRequest) GetReason() *string {
	return s.Reason
}

func (s *AddImageVulWhiteListRequest) GetSource() *string {
	return s.Source
}

func (s *AddImageVulWhiteListRequest) GetTarget() *string {
	return s.Target
}

func (s *AddImageVulWhiteListRequest) GetType() *string {
	return s.Type
}

func (s *AddImageVulWhiteListRequest) GetWhitelist() *string {
	return s.Whitelist
}

func (s *AddImageVulWhiteListRequest) SetLang(v string) *AddImageVulWhiteListRequest {
	s.Lang = &v
	return s
}

func (s *AddImageVulWhiteListRequest) SetReason(v string) *AddImageVulWhiteListRequest {
	s.Reason = &v
	return s
}

func (s *AddImageVulWhiteListRequest) SetSource(v string) *AddImageVulWhiteListRequest {
	s.Source = &v
	return s
}

func (s *AddImageVulWhiteListRequest) SetTarget(v string) *AddImageVulWhiteListRequest {
	s.Target = &v
	return s
}

func (s *AddImageVulWhiteListRequest) SetType(v string) *AddImageVulWhiteListRequest {
	s.Type = &v
	return s
}

func (s *AddImageVulWhiteListRequest) SetWhitelist(v string) *AddImageVulWhiteListRequest {
	s.Whitelist = &v
	return s
}

func (s *AddImageVulWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
