// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DisableApplicationTokenRequest
	GetApplicationId() *string
	SetApplicationTokenId(v string) *DisableApplicationTokenRequest
	GetApplicationTokenId() *string
	SetInstanceId(v string) *DisableApplicationTokenRequest
	GetInstanceId() *string
}

type DisableApplicationTokenRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// IDaaS的应用资源TokenID。
	//
	// This parameter is required.
	//
	// example:
	//
	// token_sfrwerxxxxxxxxxxxxxx
	ApplicationTokenId *string `json:"ApplicationTokenId,omitempty" xml:"ApplicationTokenId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableApplicationTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationTokenRequest) GoString() string {
	return s.String()
}

func (s *DisableApplicationTokenRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DisableApplicationTokenRequest) GetApplicationTokenId() *string {
	return s.ApplicationTokenId
}

func (s *DisableApplicationTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableApplicationTokenRequest) SetApplicationId(v string) *DisableApplicationTokenRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisableApplicationTokenRequest) SetApplicationTokenId(v string) *DisableApplicationTokenRequest {
	s.ApplicationTokenId = &v
	return s
}

func (s *DisableApplicationTokenRequest) SetInstanceId(v string) *DisableApplicationTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableApplicationTokenRequest) Validate() error {
	return dara.Validate(s)
}
