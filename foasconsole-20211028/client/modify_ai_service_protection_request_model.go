// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAiServiceProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeletionProtection(v bool) *ModifyAiServiceProtectionRequest
	GetDeletionProtection() *bool
	SetRegion(v string) *ModifyAiServiceProtectionRequest
	GetRegion() *string
}

type ModifyAiServiceProtectionRequest struct {
	// Specifies whether to enable manual shutdown protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ModifyAiServiceProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAiServiceProtectionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAiServiceProtectionRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ModifyAiServiceProtectionRequest) GetRegion() *string {
	return s.Region
}

func (s *ModifyAiServiceProtectionRequest) SetDeletionProtection(v bool) *ModifyAiServiceProtectionRequest {
	s.DeletionProtection = &v
	return s
}

func (s *ModifyAiServiceProtectionRequest) SetRegion(v string) *ModifyAiServiceProtectionRequest {
	s.Region = &v
	return s
}

func (s *ModifyAiServiceProtectionRequest) Validate() error {
	return dara.Validate(s)
}
