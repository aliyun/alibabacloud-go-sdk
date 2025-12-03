// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterDeletionProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyClusterDeletionProtectionRequest
	GetClusterId() *string
	SetProtection(v bool) *ModifyClusterDeletionProtectionRequest
	GetProtection() *bool
}

type ModifyClusterDeletionProtectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-****************
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Protection *bool `json:"Protection,omitempty" xml:"Protection,omitempty"`
}

func (s ModifyClusterDeletionProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterDeletionProtectionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyClusterDeletionProtectionRequest) GetProtection() *bool {
	return s.Protection
}

func (s *ModifyClusterDeletionProtectionRequest) SetClusterId(v string) *ModifyClusterDeletionProtectionRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterDeletionProtectionRequest) SetProtection(v bool) *ModifyClusterDeletionProtectionRequest {
	s.Protection = &v
	return s
}

func (s *ModifyClusterDeletionProtectionRequest) Validate() error {
	return dara.Validate(s)
}
