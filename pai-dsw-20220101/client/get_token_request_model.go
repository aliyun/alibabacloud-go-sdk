// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireTime(v int32) *GetTokenRequest
	GetExpireTime() *int32
	SetInstanceId(v string) *GetTokenRequest
	GetInstanceId() *string
}

type GetTokenRequest struct {
	// The validity period. Unit: seconds.
	//
	// example:
	//
	// 60
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *GetTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTokenRequest) SetExpireTime(v int32) *GetTokenRequest {
	s.ExpireTime = &v
	return s
}

func (s *GetTokenRequest) SetInstanceId(v string) *GetTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTokenRequest) Validate() error {
	return dara.Validate(s)
}
