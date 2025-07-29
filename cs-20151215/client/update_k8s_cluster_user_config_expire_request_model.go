// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sClusterUserConfigExpireRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireHour(v int64) *UpdateK8sClusterUserConfigExpireRequest
	GetExpireHour() *int64
	SetUser(v string) *UpdateK8sClusterUserConfigExpireRequest
	GetUser() *string
}

type UpdateK8sClusterUserConfigExpireRequest struct {
	// Specifies the expiration time of the kubeconfig file. Unit: hours.
	//
	// Valid values: [1, 1876000]. The maximum value is 100 years.
	//
	// This parameter is required.
	//
	// example:
	//
	// 720
	ExpireHour *int64 `json:"expire_hour,omitempty" xml:"expire_hour,omitempty"`
	// The RAM user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// The ID of the Resource Access Management (RAM) user that you use.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s UpdateK8sClusterUserConfigExpireRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sClusterUserConfigExpireRequest) GoString() string {
	return s.String()
}

func (s *UpdateK8sClusterUserConfigExpireRequest) GetExpireHour() *int64 {
	return s.ExpireHour
}

func (s *UpdateK8sClusterUserConfigExpireRequest) GetUser() *string {
	return s.User
}

func (s *UpdateK8sClusterUserConfigExpireRequest) SetExpireHour(v int64) *UpdateK8sClusterUserConfigExpireRequest {
	s.ExpireHour = &v
	return s
}

func (s *UpdateK8sClusterUserConfigExpireRequest) SetUser(v string) *UpdateK8sClusterUserConfigExpireRequest {
	s.User = &v
	return s
}

func (s *UpdateK8sClusterUserConfigExpireRequest) Validate() error {
	return dara.Validate(s)
}
