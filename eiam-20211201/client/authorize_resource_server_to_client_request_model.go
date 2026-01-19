// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceServerToClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientApplicationId(v string) *AuthorizeResourceServerToClientRequest
	GetClientApplicationId() *string
	SetInstanceId(v string) *AuthorizeResourceServerToClientRequest
	GetInstanceId() *string
	SetResourceServerApplicationId(v string) *AuthorizeResourceServerToClientRequest
	GetResourceServerApplicationId() *string
}

type AuthorizeResourceServerToClientRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientApplicationId *string `json:"ClientApplicationId,omitempty" xml:"ClientApplicationId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ResourceServerApplicationId *string `json:"ResourceServerApplicationId,omitempty" xml:"ResourceServerApplicationId,omitempty"`
}

func (s AuthorizeResourceServerToClientRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceServerToClientRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceServerToClientRequest) GetClientApplicationId() *string {
	return s.ClientApplicationId
}

func (s *AuthorizeResourceServerToClientRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AuthorizeResourceServerToClientRequest) GetResourceServerApplicationId() *string {
	return s.ResourceServerApplicationId
}

func (s *AuthorizeResourceServerToClientRequest) SetClientApplicationId(v string) *AuthorizeResourceServerToClientRequest {
	s.ClientApplicationId = &v
	return s
}

func (s *AuthorizeResourceServerToClientRequest) SetInstanceId(v string) *AuthorizeResourceServerToClientRequest {
	s.InstanceId = &v
	return s
}

func (s *AuthorizeResourceServerToClientRequest) SetResourceServerApplicationId(v string) *AuthorizeResourceServerToClientRequest {
	s.ResourceServerApplicationId = &v
	return s
}

func (s *AuthorizeResourceServerToClientRequest) Validate() error {
	return dara.Validate(s)
}
