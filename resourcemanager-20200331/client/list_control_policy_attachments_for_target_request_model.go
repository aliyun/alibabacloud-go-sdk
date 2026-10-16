// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListControlPolicyAttachmentsForTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *ListControlPolicyAttachmentsForTargetRequest
	GetLanguage() *string
	SetTargetId(v string) *ListControlPolicyAttachmentsForTargetRequest
	GetTargetId() *string
}

type ListControlPolicyAttachmentsForTargetRequest struct {
	// The language in which the control policy descriptions are returned. Valid values:
	//
	// - zh-CN (default): Chinese.
	//
	// - en: English.
	//
	// - ja: Japanese.
	//
	// > This parameter is valid only for system control policies.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The target node. Valid values:
	//
	// - Root ID: queries the control policies attached to the Root node.
	//
	// - Resource folder ID: queries the control policies attached to a resource folder.
	//
	// - Member ID: queries the control policies attached to a member.
	//
	// This parameter is required.
	//
	// example:
	//
	// fd-ZDNPiT****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s ListControlPolicyAttachmentsForTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetRequest) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListControlPolicyAttachmentsForTargetRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *ListControlPolicyAttachmentsForTargetRequest) SetLanguage(v string) *ListControlPolicyAttachmentsForTargetRequest {
	s.Language = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetRequest) SetTargetId(v string) *ListControlPolicyAttachmentsForTargetRequest {
	s.TargetId = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetRequest) Validate() error {
	return dara.Validate(s)
}
