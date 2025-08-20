// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorizationTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetAuthorizationTokenRequest
	GetInstanceId() *string
}

type GetAuthorizationTokenRequest struct {
	// The ID of the request.
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

func (s *GetAuthorizationTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAuthorizationTokenRequest) SetInstanceId(v string) *GetAuthorizationTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAuthorizationTokenRequest) Validate() error {
	return dara.Validate(s)
}
