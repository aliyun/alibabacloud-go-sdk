// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationResourceServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DisableApplicationResourceServerRequest
	GetApplicationId() *string
	SetInstanceId(v string) *DisableApplicationResourceServerRequest
	GetInstanceId() *string
}

type DisableApplicationResourceServerRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableApplicationResourceServerRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationResourceServerRequest) GoString() string {
	return s.String()
}

func (s *DisableApplicationResourceServerRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DisableApplicationResourceServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableApplicationResourceServerRequest) SetApplicationId(v string) *DisableApplicationResourceServerRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisableApplicationResourceServerRequest) SetInstanceId(v string) *DisableApplicationResourceServerRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableApplicationResourceServerRequest) Validate() error {
	return dara.Validate(s)
}
