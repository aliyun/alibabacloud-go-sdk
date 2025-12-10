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
	// The language in which you want to return the descriptions of the access control policies. Valid values:
	//
	// 	- zh-CN (default value): Chinese
	//
	// 	- en: English
	//
	// 	- ja: Japanese
	//
	// >  This parameter is valid only for system access control policies.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The ID of the object whose access control policies you want to query. Access control policies can be attached to the following objects:
	//
	// 	- Root folder
	//
	// 	- Subfolders of the Root folder
	//
	// 	- Members
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
