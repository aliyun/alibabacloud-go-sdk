// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAnycastEipAddressAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnycastId(v string) *ModifyAnycastEipAddressAttributeRequest
	GetAnycastId() *string
	SetDescription(v string) *ModifyAnycastEipAddressAttributeRequest
	GetDescription() *string
	SetName(v string) *ModifyAnycastEipAddressAttributeRequest
	GetName() *string
}

type ModifyAnycastEipAddressAttributeRequest struct {
	// The ID of the Anycast EIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// aeip-bp1ix34fralt4ykf3****
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The description of the Anycast EIP.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// docdesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the Anycast EIP.
	//
	// The name must be 0 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// docname
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyAnycastEipAddressAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAnycastEipAddressAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAnycastEipAddressAttributeRequest) GetAnycastId() *string {
	return s.AnycastId
}

func (s *ModifyAnycastEipAddressAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAnycastEipAddressAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyAnycastEipAddressAttributeRequest) SetAnycastId(v string) *ModifyAnycastEipAddressAttributeRequest {
	s.AnycastId = &v
	return s
}

func (s *ModifyAnycastEipAddressAttributeRequest) SetDescription(v string) *ModifyAnycastEipAddressAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyAnycastEipAddressAttributeRequest) SetName(v string) *ModifyAnycastEipAddressAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyAnycastEipAddressAttributeRequest) Validate() error {
	return dara.Validate(s)
}
