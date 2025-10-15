// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationTokenExpirationTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateApplicationTokenExpirationTimeRequest
	GetApplicationId() *string
	SetApplicationTokenId(v string) *UpdateApplicationTokenExpirationTimeRequest
	GetApplicationTokenId() *string
	SetExpirationTime(v int64) *UpdateApplicationTokenExpirationTimeRequest
	GetExpirationTime() *int64
	SetInstanceId(v string) *UpdateApplicationTokenExpirationTimeRequest
	GetInstanceId() *string
}

type UpdateApplicationTokenExpirationTimeRequest struct {
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
	// 不填，默认1年后到期
	//
	// This parameter is required.
	//
	// example:
	//
	// 1735530123762
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateApplicationTokenExpirationTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationTokenExpirationTimeRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationTokenExpirationTimeRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationTokenExpirationTimeRequest) GetApplicationTokenId() *string {
	return s.ApplicationTokenId
}

func (s *UpdateApplicationTokenExpirationTimeRequest) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *UpdateApplicationTokenExpirationTimeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateApplicationTokenExpirationTimeRequest) SetApplicationId(v string) *UpdateApplicationTokenExpirationTimeRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationTokenExpirationTimeRequest) SetApplicationTokenId(v string) *UpdateApplicationTokenExpirationTimeRequest {
	s.ApplicationTokenId = &v
	return s
}

func (s *UpdateApplicationTokenExpirationTimeRequest) SetExpirationTime(v int64) *UpdateApplicationTokenExpirationTimeRequest {
	s.ExpirationTime = &v
	return s
}

func (s *UpdateApplicationTokenExpirationTimeRequest) SetInstanceId(v string) *UpdateApplicationTokenExpirationTimeRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateApplicationTokenExpirationTimeRequest) Validate() error {
	return dara.Validate(s)
}
