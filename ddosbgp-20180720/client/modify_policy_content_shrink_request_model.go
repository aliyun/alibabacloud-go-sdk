// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyContentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentShrink(v string) *ModifyPolicyContentShrinkRequest
	GetContentShrink() *string
	SetId(v string) *ModifyPolicyContentShrinkRequest
	GetId() *string
	SetName(v string) *ModifyPolicyContentShrinkRequest
	GetName() *string
	SetPortVersion(v string) *ModifyPolicyContentShrinkRequest
	GetPortVersion() *string
}

type ModifyPolicyContentShrinkRequest struct {
	// The policy content.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 83967609-7ea5-4f6d-a6ea-380b09e****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// demo**
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PortVersion *string `json:"PortVersion,omitempty" xml:"PortVersion,omitempty"`
}

func (s ModifyPolicyContentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *ModifyPolicyContentShrinkRequest) GetId() *string {
	return s.Id
}

func (s *ModifyPolicyContentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ModifyPolicyContentShrinkRequest) GetPortVersion() *string {
	return s.PortVersion
}

func (s *ModifyPolicyContentShrinkRequest) SetContentShrink(v string) *ModifyPolicyContentShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *ModifyPolicyContentShrinkRequest) SetId(v string) *ModifyPolicyContentShrinkRequest {
	s.Id = &v
	return s
}

func (s *ModifyPolicyContentShrinkRequest) SetName(v string) *ModifyPolicyContentShrinkRequest {
	s.Name = &v
	return s
}

func (s *ModifyPolicyContentShrinkRequest) SetPortVersion(v string) *ModifyPolicyContentShrinkRequest {
	s.PortVersion = &v
	return s
}

func (s *ModifyPolicyContentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
