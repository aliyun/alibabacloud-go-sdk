// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEmgVulSubmitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyEmgVulSubmitRequest
	GetLang() *string
	SetName(v string) *ModifyEmgVulSubmitRequest
	GetName() *string
	SetResourceDirectoryAccountId(v int64) *ModifyEmgVulSubmitRequest
	GetResourceDirectoryAccountId() *int64
	SetUserAgreement(v string) *ModifyEmgVulSubmitRequest
	GetUserAgreement() *string
}

type ModifyEmgVulSubmitRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the vulnerability.
	//
	// This parameter is required.
	//
	// example:
	//
	// scan:ASCV-2019-032401
	Name                       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceDirectoryAccountId *int64  `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// Specifies whether to scan for urgent vulnerabilities. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no**
	//
	// This parameter is required.
	//
	// example:
	//
	// yes
	UserAgreement *string `json:"UserAgreement,omitempty" xml:"UserAgreement,omitempty"`
}

func (s ModifyEmgVulSubmitRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmgVulSubmitRequest) GoString() string {
	return s.String()
}

func (s *ModifyEmgVulSubmitRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyEmgVulSubmitRequest) GetName() *string {
	return s.Name
}

func (s *ModifyEmgVulSubmitRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *ModifyEmgVulSubmitRequest) GetUserAgreement() *string {
	return s.UserAgreement
}

func (s *ModifyEmgVulSubmitRequest) SetLang(v string) *ModifyEmgVulSubmitRequest {
	s.Lang = &v
	return s
}

func (s *ModifyEmgVulSubmitRequest) SetName(v string) *ModifyEmgVulSubmitRequest {
	s.Name = &v
	return s
}

func (s *ModifyEmgVulSubmitRequest) SetResourceDirectoryAccountId(v int64) *ModifyEmgVulSubmitRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *ModifyEmgVulSubmitRequest) SetUserAgreement(v string) *ModifyEmgVulSubmitRequest {
	s.UserAgreement = &v
	return s
}

func (s *ModifyEmgVulSubmitRequest) Validate() error {
	return dara.Validate(s)
}
