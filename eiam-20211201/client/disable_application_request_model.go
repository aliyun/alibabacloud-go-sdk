// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DisableApplicationRequest
	GetApplicationId() *string
	SetInstanceId(v string) *DisableApplicationRequest
	GetInstanceId() *string
}

type DisableApplicationRequest struct {
	// The ID of the application that you want to disable.
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

func (s DisableApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationRequest) GoString() string {
	return s.String()
}

func (s *DisableApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DisableApplicationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableApplicationRequest) SetApplicationId(v string) *DisableApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisableApplicationRequest) SetInstanceId(v string) *DisableApplicationRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableApplicationRequest) Validate() error {
	return dara.Validate(s)
}
