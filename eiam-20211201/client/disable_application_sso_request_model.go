// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationSsoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DisableApplicationSsoRequest
	GetApplicationId() *string
	SetInstanceId(v string) *DisableApplicationSsoRequest
	GetInstanceId() *string
}

type DisableApplicationSsoRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableApplicationSsoRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationSsoRequest) GoString() string {
	return s.String()
}

func (s *DisableApplicationSsoRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DisableApplicationSsoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableApplicationSsoRequest) SetApplicationId(v string) *DisableApplicationSsoRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisableApplicationSsoRequest) SetInstanceId(v string) *DisableApplicationSsoRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableApplicationSsoRequest) Validate() error {
	return dara.Validate(s)
}
