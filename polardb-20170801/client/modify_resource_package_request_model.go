// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourcePackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoQuota(v bool) *ModifyResourcePackageRequest
	GetAutoQuota() *bool
	SetResourcePackageId(v string) *ModifyResourcePackageRequest
	GetResourcePackageId() *string
}

type ModifyResourcePackageRequest struct {
	// example:
	//
	// true
	AutoQuota *bool `json:"AutoQuota,omitempty" xml:"AutoQuota,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pm-bp11b0i9389******
	ResourcePackageId *string `json:"ResourcePackageId,omitempty" xml:"ResourcePackageId,omitempty"`
}

func (s ModifyResourcePackageRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourcePackageRequest) GoString() string {
	return s.String()
}

func (s *ModifyResourcePackageRequest) GetAutoQuota() *bool {
	return s.AutoQuota
}

func (s *ModifyResourcePackageRequest) GetResourcePackageId() *string {
	return s.ResourcePackageId
}

func (s *ModifyResourcePackageRequest) SetAutoQuota(v bool) *ModifyResourcePackageRequest {
	s.AutoQuota = &v
	return s
}

func (s *ModifyResourcePackageRequest) SetResourcePackageId(v string) *ModifyResourcePackageRequest {
	s.ResourcePackageId = &v
	return s
}

func (s *ModifyResourcePackageRequest) Validate() error {
	return dara.Validate(s)
}
