// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorizationServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationServerId(v string) *GetAuthorizationServerRequest
	GetAuthorizationServerId() *string
	SetInstanceId(v string) *GetAuthorizationServerRequest
	GetInstanceId() *string
}

type GetAuthorizationServerRequest struct {
	// The authorization server ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// iauths_mkv7rgt4d7i4u7zqtzev2mxxxx
	AuthorizationServerId *string `json:"AuthorizationServerId,omitempty" xml:"AuthorizationServerId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAuthorizationServerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationServerRequest) GoString() string {
	return s.String()
}

func (s *GetAuthorizationServerRequest) GetAuthorizationServerId() *string {
	return s.AuthorizationServerId
}

func (s *GetAuthorizationServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAuthorizationServerRequest) SetAuthorizationServerId(v string) *GetAuthorizationServerRequest {
	s.AuthorizationServerId = &v
	return s
}

func (s *GetAuthorizationServerRequest) SetInstanceId(v string) *GetAuthorizationServerRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAuthorizationServerRequest) Validate() error {
	return dara.Validate(s)
}
