// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainApplicationTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ObtainApplicationTokenRequest
	GetApplicationId() *string
	SetApplicationTokenId(v string) *ObtainApplicationTokenRequest
	GetApplicationTokenId() *string
	SetInstanceId(v string) *ObtainApplicationTokenRequest
	GetInstanceId() *string
}

type ObtainApplicationTokenRequest struct {
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

func (s ObtainApplicationTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s ObtainApplicationTokenRequest) GoString() string {
	return s.String()
}

func (s *ObtainApplicationTokenRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ObtainApplicationTokenRequest) GetApplicationTokenId() *string {
	return s.ApplicationTokenId
}

func (s *ObtainApplicationTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ObtainApplicationTokenRequest) SetApplicationId(v string) *ObtainApplicationTokenRequest {
	s.ApplicationId = &v
	return s
}

func (s *ObtainApplicationTokenRequest) SetApplicationTokenId(v string) *ObtainApplicationTokenRequest {
	s.ApplicationTokenId = &v
	return s
}

func (s *ObtainApplicationTokenRequest) SetInstanceId(v string) *ObtainApplicationTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *ObtainApplicationTokenRequest) Validate() error {
	return dara.Validate(s)
}
