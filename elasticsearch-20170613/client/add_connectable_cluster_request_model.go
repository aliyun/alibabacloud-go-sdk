// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddConnectableClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *AddConnectableClusterRequest
	GetBody() *string
	SetClientToken(v string) *AddConnectableClusterRequest
	GetClientToken() *string
}

type AddConnectableClusterRequest struct {
	// The remote instance ID. Specifies the remote instance ID to establish network connectivity with. The remote instance must be in the same VPC as the current instance.
	//
	// example:
	//
	// {     "instanceId":"es-cn-09k1rgid9000g****" }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// A client-generated token used to ensure the idempotence of the request. The value must be unique across different requests and cannot exceed 64 ASCII characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s AddConnectableClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s AddConnectableClusterRequest) GoString() string {
	return s.String()
}

func (s *AddConnectableClusterRequest) GetBody() *string {
	return s.Body
}

func (s *AddConnectableClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddConnectableClusterRequest) SetBody(v string) *AddConnectableClusterRequest {
	s.Body = &v
	return s
}

func (s *AddConnectableClusterRequest) SetClientToken(v string) *AddConnectableClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *AddConnectableClusterRequest) Validate() error {
	return dara.Validate(s)
}
