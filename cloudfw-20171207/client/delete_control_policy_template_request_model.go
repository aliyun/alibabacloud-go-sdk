// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteControlPolicyTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteControlPolicyTemplateRequest
	GetLang() *string
	SetSourceIp(v string) *DeleteControlPolicyTemplateRequest
	GetSourceIp() *string
	SetTemplateId(v string) *DeleteControlPolicyTemplateRequest
	GetTemplateId() *string
}

type DeleteControlPolicyTemplateRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 61.178.12.52
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the access control policy template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 719ce620-ae23-4e42-9f93-9191b4400b55
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteControlPolicyTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteControlPolicyTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyTemplateRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteControlPolicyTemplateRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DeleteControlPolicyTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteControlPolicyTemplateRequest) SetLang(v string) *DeleteControlPolicyTemplateRequest {
	s.Lang = &v
	return s
}

func (s *DeleteControlPolicyTemplateRequest) SetSourceIp(v string) *DeleteControlPolicyTemplateRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteControlPolicyTemplateRequest) SetTemplateId(v string) *DeleteControlPolicyTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteControlPolicyTemplateRequest) Validate() error {
	return dara.Validate(s)
}
