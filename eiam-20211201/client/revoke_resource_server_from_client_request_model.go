// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourceServerFromClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientApplicationId(v string) *RevokeResourceServerFromClientRequest
	GetClientApplicationId() *string
	SetInstanceId(v string) *RevokeResourceServerFromClientRequest
	GetInstanceId() *string
	SetResourceServerApplicationId(v string) *RevokeResourceServerFromClientRequest
	GetResourceServerApplicationId() *string
}

type RevokeResourceServerFromClientRequest struct {
	// The ID of the Client application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientApplicationId *string `json:"ClientApplicationId,omitempty" xml:"ClientApplicationId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the ResourceServer application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ResourceServerApplicationId *string `json:"ResourceServerApplicationId,omitempty" xml:"ResourceServerApplicationId,omitempty"`
}

func (s RevokeResourceServerFromClientRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourceServerFromClientRequest) GoString() string {
	return s.String()
}

func (s *RevokeResourceServerFromClientRequest) GetClientApplicationId() *string {
	return s.ClientApplicationId
}

func (s *RevokeResourceServerFromClientRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RevokeResourceServerFromClientRequest) GetResourceServerApplicationId() *string {
	return s.ResourceServerApplicationId
}

func (s *RevokeResourceServerFromClientRequest) SetClientApplicationId(v string) *RevokeResourceServerFromClientRequest {
	s.ClientApplicationId = &v
	return s
}

func (s *RevokeResourceServerFromClientRequest) SetInstanceId(v string) *RevokeResourceServerFromClientRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeResourceServerFromClientRequest) SetResourceServerApplicationId(v string) *RevokeResourceServerFromClientRequest {
	s.ResourceServerApplicationId = &v
	return s
}

func (s *RevokeResourceServerFromClientRequest) Validate() error {
	return dara.Validate(s)
}
