// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreateApplicationTokenRequest
	GetApplicationId() *string
	SetApplicationTokenType(v string) *CreateApplicationTokenRequest
	GetApplicationTokenType() *string
	SetExpirationTime(v int64) *CreateApplicationTokenRequest
	GetExpirationTime() *int64
	SetInstanceId(v string) *CreateApplicationTokenRequest
	GetInstanceId() *string
}

type CreateApplicationTokenRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application token type.
	//
	// This parameter is required.
	//
	// example:
	//
	// bearer_token
	ApplicationTokenType *string `json:"ApplicationTokenType,omitempty" xml:"ApplicationTokenType,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 1735530123762
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateApplicationTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationTokenRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateApplicationTokenRequest) GetApplicationTokenType() *string {
	return s.ApplicationTokenType
}

func (s *CreateApplicationTokenRequest) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *CreateApplicationTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateApplicationTokenRequest) SetApplicationId(v string) *CreateApplicationTokenRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationTokenRequest) SetApplicationTokenType(v string) *CreateApplicationTokenRequest {
	s.ApplicationTokenType = &v
	return s
}

func (s *CreateApplicationTokenRequest) SetExpirationTime(v int64) *CreateApplicationTokenRequest {
	s.ExpirationTime = &v
	return s
}

func (s *CreateApplicationTokenRequest) SetInstanceId(v string) *CreateApplicationTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateApplicationTokenRequest) Validate() error {
	return dara.Validate(s)
}
