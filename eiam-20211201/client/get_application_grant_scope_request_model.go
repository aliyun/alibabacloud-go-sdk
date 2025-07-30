// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationGrantScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetApplicationGrantScopeRequest
	GetApplicationId() *string
	SetInstanceId(v string) *GetApplicationGrantScopeRequest
	GetInstanceId() *string
}

type GetApplicationGrantScopeRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetApplicationGrantScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationGrantScopeRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationGrantScopeRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationGrantScopeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApplicationGrantScopeRequest) SetApplicationId(v string) *GetApplicationGrantScopeRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationGrantScopeRequest) SetInstanceId(v string) *GetApplicationGrantScopeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationGrantScopeRequest) Validate() error {
	return dara.Validate(s)
}
