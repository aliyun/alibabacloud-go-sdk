// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableResourceServerCustomSubjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DisableResourceServerCustomSubjectRequest
	GetApplicationId() *string
	SetInstanceId(v string) *DisableResourceServerCustomSubjectRequest
	GetInstanceId() *string
}

type DisableResourceServerCustomSubjectRequest struct {
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

func (s DisableResourceServerCustomSubjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableResourceServerCustomSubjectRequest) GoString() string {
	return s.String()
}

func (s *DisableResourceServerCustomSubjectRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DisableResourceServerCustomSubjectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableResourceServerCustomSubjectRequest) SetApplicationId(v string) *DisableResourceServerCustomSubjectRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisableResourceServerCustomSubjectRequest) SetInstanceId(v string) *DisableResourceServerCustomSubjectRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableResourceServerCustomSubjectRequest) Validate() error {
	return dara.Validate(s)
}
