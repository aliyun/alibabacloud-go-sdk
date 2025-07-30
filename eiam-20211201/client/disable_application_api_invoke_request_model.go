// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationApiInvokeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DisableApplicationApiInvokeRequest
	GetApplicationId() *string
	SetInstanceId(v string) *DisableApplicationApiInvokeRequest
	GetInstanceId() *string
}

type DisableApplicationApiInvokeRequest struct {
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

func (s DisableApplicationApiInvokeRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationApiInvokeRequest) GoString() string {
	return s.String()
}

func (s *DisableApplicationApiInvokeRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DisableApplicationApiInvokeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableApplicationApiInvokeRequest) SetApplicationId(v string) *DisableApplicationApiInvokeRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisableApplicationApiInvokeRequest) SetInstanceId(v string) *DisableApplicationApiInvokeRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableApplicationApiInvokeRequest) Validate() error {
	return dara.Validate(s)
}
