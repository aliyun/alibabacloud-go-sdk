// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateApplicationDescriptionRequest
	GetApplicationId() *string
	SetDescription(v string) *UpdateApplicationDescriptionRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateApplicationDescriptionRequest
	GetInstanceId() *string
}

type UpdateApplicationDescriptionRequest struct {
	// The ID of the application that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// A demo application that is used for test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk2676xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateApplicationDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationDescriptionRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateApplicationDescriptionRequest) SetApplicationId(v string) *UpdateApplicationDescriptionRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationDescriptionRequest) SetDescription(v string) *UpdateApplicationDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateApplicationDescriptionRequest) SetInstanceId(v string) *UpdateApplicationDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateApplicationDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
