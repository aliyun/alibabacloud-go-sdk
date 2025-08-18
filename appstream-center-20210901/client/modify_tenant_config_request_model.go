// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTenantConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupExpireRemind(v bool) *ModifyTenantConfigRequest
	GetAppInstanceGroupExpireRemind() *bool
}

type ModifyTenantConfigRequest struct {
	// Specifies whether to enable the resource expiration reminder feature.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AppInstanceGroupExpireRemind *bool `json:"AppInstanceGroupExpireRemind,omitempty" xml:"AppInstanceGroupExpireRemind,omitempty"`
}

func (s ModifyTenantConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTenantConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyTenantConfigRequest) GetAppInstanceGroupExpireRemind() *bool {
	return s.AppInstanceGroupExpireRemind
}

func (s *ModifyTenantConfigRequest) SetAppInstanceGroupExpireRemind(v bool) *ModifyTenantConfigRequest {
	s.AppInstanceGroupExpireRemind = &v
	return s
}

func (s *ModifyTenantConfigRequest) Validate() error {
	return dara.Validate(s)
}
