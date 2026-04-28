// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppAccessStrategy interface {
	dara.Model
	String() string
	GoString() string
	SetEffect(v string) *AppAccessStrategy
	GetEffect() *string
	SetExceptAppIdList(v []*string) *AppAccessStrategy
	GetExceptAppIdList() []*string
}

type AppAccessStrategy struct {
	// The global access policy of the application. Valid values:
	//
	// 	- allow: The domain allows access from all applications.
	//
	// 	- deny: The domain denies access from all apps. This is the default value.
	//
	// Recommended settings:
	//
	// 1.  Set effect to deny.
	//
	// 2.  Specify except_app_id_list to allow specific applications to access the domain. Example: ["appid1", "appid2"].
	//
	// example:
	//
	// deny
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	// The IDs of applications excluded from the global access policy.
	//
	// 	- If you set effect to allow, which indicates that the domain allows access from all applications, the applications specified by this parameter value cannot access the domain.
	//
	// 	- If you set effect to deny, which indicates that the domain denies access from all applications, the applications specified by this parameter value can access the domain.
	ExceptAppIdList []*string `json:"except_app_id_list,omitempty" xml:"except_app_id_list,omitempty" type:"Repeated"`
}

func (s AppAccessStrategy) String() string {
	return dara.Prettify(s)
}

func (s AppAccessStrategy) GoString() string {
	return s.String()
}

func (s *AppAccessStrategy) GetEffect() *string {
	return s.Effect
}

func (s *AppAccessStrategy) GetExceptAppIdList() []*string {
	return s.ExceptAppIdList
}

func (s *AppAccessStrategy) SetEffect(v string) *AppAccessStrategy {
	s.Effect = &v
	return s
}

func (s *AppAccessStrategy) SetExceptAppIdList(v []*string) *AppAccessStrategy {
	s.ExceptAppIdList = v
	return s
}

func (s *AppAccessStrategy) Validate() error {
	return dara.Validate(s)
}
