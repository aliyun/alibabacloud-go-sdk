// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DeleteApplicationTokenRequest
	GetApplicationId() *string
	SetApplicationTokenId(v string) *DeleteApplicationTokenRequest
	GetApplicationTokenId() *string
	SetInstanceId(v string) *DeleteApplicationTokenRequest
	GetInstanceId() *string
}

type DeleteApplicationTokenRequest struct {
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

func (s DeleteApplicationTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationTokenRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationTokenRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeleteApplicationTokenRequest) GetApplicationTokenId() *string {
	return s.ApplicationTokenId
}

func (s *DeleteApplicationTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteApplicationTokenRequest) SetApplicationId(v string) *DeleteApplicationTokenRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeleteApplicationTokenRequest) SetApplicationTokenId(v string) *DeleteApplicationTokenRequest {
	s.ApplicationTokenId = &v
	return s
}

func (s *DeleteApplicationTokenRequest) SetInstanceId(v string) *DeleteApplicationTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteApplicationTokenRequest) Validate() error {
	return dara.Validate(s)
}
