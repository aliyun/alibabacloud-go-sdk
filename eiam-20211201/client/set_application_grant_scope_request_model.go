// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationGrantScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *SetApplicationGrantScopeRequest
	GetApplicationId() *string
	SetGrantScopes(v []*string) *SetApplicationGrantScopeRequest
	GetGrantScopes() []*string
	SetInstanceId(v string) *SetApplicationGrantScopeRequest
	GetInstanceId() *string
}

type SetApplicationGrantScopeRequest struct {
	// The ID of the application that you want to configure.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The permissions of the Developer API feature.
	GrantScopes []*string `json:"GrantScopes,omitempty" xml:"GrantScopes,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SetApplicationGrantScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationGrantScopeRequest) GoString() string {
	return s.String()
}

func (s *SetApplicationGrantScopeRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *SetApplicationGrantScopeRequest) GetGrantScopes() []*string {
	return s.GrantScopes
}

func (s *SetApplicationGrantScopeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetApplicationGrantScopeRequest) SetApplicationId(v string) *SetApplicationGrantScopeRequest {
	s.ApplicationId = &v
	return s
}

func (s *SetApplicationGrantScopeRequest) SetGrantScopes(v []*string) *SetApplicationGrantScopeRequest {
	s.GrantScopes = v
	return s
}

func (s *SetApplicationGrantScopeRequest) SetInstanceId(v string) *SetApplicationGrantScopeRequest {
	s.InstanceId = &v
	return s
}

func (s *SetApplicationGrantScopeRequest) Validate() error {
	return dara.Validate(s)
}
