// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetRuleHitCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclUuid(v string) *ResetRuleHitCountRequest
	GetAclUuid() *string
	SetLang(v string) *ResetRuleHitCountRequest
	GetLang() *string
	SetSourceIp(v string) *ResetRuleHitCountRequest
	GetSourceIp() *string
}

type ResetRuleHitCountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 211fd804-30f5-470f-ab26-c465a4****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 112.64.126.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ResetRuleHitCountRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetRuleHitCountRequest) GoString() string {
	return s.String()
}

func (s *ResetRuleHitCountRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *ResetRuleHitCountRequest) GetLang() *string {
	return s.Lang
}

func (s *ResetRuleHitCountRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ResetRuleHitCountRequest) SetAclUuid(v string) *ResetRuleHitCountRequest {
	s.AclUuid = &v
	return s
}

func (s *ResetRuleHitCountRequest) SetLang(v string) *ResetRuleHitCountRequest {
	s.Lang = &v
	return s
}

func (s *ResetRuleHitCountRequest) SetSourceIp(v string) *ResetRuleHitCountRequest {
	s.SourceIp = &v
	return s
}

func (s *ResetRuleHitCountRequest) Validate() error {
	return dara.Validate(s)
}
