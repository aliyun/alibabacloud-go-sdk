// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageVulWhitelistTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpdateImageVulWhitelistTargetRequest
	GetId() *int64
	SetLang(v string) *UpdateImageVulWhitelistTargetRequest
	GetLang() *string
	SetReason(v string) *UpdateImageVulWhitelistTargetRequest
	GetReason() *string
	SetSource(v string) *UpdateImageVulWhitelistTargetRequest
	GetSource() *string
	SetTarget(v string) *UpdateImageVulWhitelistTargetRequest
	GetTarget() *string
}

type UpdateImageVulWhitelistTargetRequest struct {
	// The whitelist ID.
	//
	// example:
	//
	// 2000083
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// The reason why you add the vulnerability to the whitelist.
	//
	// example:
	//
	// ignore
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The source of the whitelist. Valid values:
	//
	// 	- **image**
	//
	// 	- **agentless**
	//
	// example:
	//
	// image
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The vulnerability that you want to add to the whitelist. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **type**: The type of the vulnerability. The value is fixed to repo.
	//
	// 	- **target**: The content of the vulnerability. The value is in the format of Namespace/Image repository.
	//
	// example:
	//
	// {\\"type\\":\\"repo\\",\\"target\\":[\\"sas_test/script_0209\\",\\"sas_test/script\\"]}
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s UpdateImageVulWhitelistTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageVulWhitelistTargetRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageVulWhitelistTargetRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateImageVulWhitelistTargetRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateImageVulWhitelistTargetRequest) GetReason() *string {
	return s.Reason
}

func (s *UpdateImageVulWhitelistTargetRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateImageVulWhitelistTargetRequest) GetTarget() *string {
	return s.Target
}

func (s *UpdateImageVulWhitelistTargetRequest) SetId(v int64) *UpdateImageVulWhitelistTargetRequest {
	s.Id = &v
	return s
}

func (s *UpdateImageVulWhitelistTargetRequest) SetLang(v string) *UpdateImageVulWhitelistTargetRequest {
	s.Lang = &v
	return s
}

func (s *UpdateImageVulWhitelistTargetRequest) SetReason(v string) *UpdateImageVulWhitelistTargetRequest {
	s.Reason = &v
	return s
}

func (s *UpdateImageVulWhitelistTargetRequest) SetSource(v string) *UpdateImageVulWhitelistTargetRequest {
	s.Source = &v
	return s
}

func (s *UpdateImageVulWhitelistTargetRequest) SetTarget(v string) *UpdateImageVulWhitelistTargetRequest {
	s.Target = &v
	return s
}

func (s *UpdateImageVulWhitelistTargetRequest) Validate() error {
	return dara.Validate(s)
}
