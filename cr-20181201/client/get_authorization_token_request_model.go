// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorizationTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpiresInHours(v int32) *GetAuthorizationTokenRequest
	GetExpiresInHours() *int32
	SetInstanceId(v string) *GetAuthorizationTokenRequest
	GetInstanceId() *string
}

type GetAuthorizationTokenRequest struct {
	// The validity period of the temporary credential, in hours. Valid values: 1 to 24.
	//
	// example:
	//
	// 1
	ExpiresInHours *int32 `json:"ExpiresInHours,omitempty" xml:"ExpiresInHours,omitempty"`
	// The repository instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcvaduwb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAuthorizationTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationTokenRequest) GoString() string {
	return s.String()
}

func (s *GetAuthorizationTokenRequest) GetExpiresInHours() *int32 {
	return s.ExpiresInHours
}

func (s *GetAuthorizationTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAuthorizationTokenRequest) SetExpiresInHours(v int32) *GetAuthorizationTokenRequest {
	s.ExpiresInHours = &v
	return s
}

func (s *GetAuthorizationTokenRequest) SetInstanceId(v string) *GetAuthorizationTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAuthorizationTokenRequest) Validate() error {
	return dara.Validate(s)
}
