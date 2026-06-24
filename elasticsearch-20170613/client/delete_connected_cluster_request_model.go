// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectedClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteConnectedClusterRequest
	GetClientToken() *string
	SetConnectedInstanceId(v string) *DeleteConnectedClusterRequest
	GetConnectedInstanceId() *string
}

type DeleteConnectedClusterRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The instance ID of the remote instance that is connected over the network.
	//
	// This parameter is required.
	//
	// example:
	//
	// es-cn-09k1rgid9000g****
	ConnectedInstanceId *string `json:"connectedInstanceId,omitempty" xml:"connectedInstanceId,omitempty"`
}

func (s DeleteConnectedClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectedClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteConnectedClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteConnectedClusterRequest) GetConnectedInstanceId() *string {
	return s.ConnectedInstanceId
}

func (s *DeleteConnectedClusterRequest) SetClientToken(v string) *DeleteConnectedClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteConnectedClusterRequest) SetConnectedInstanceId(v string) *DeleteConnectedClusterRequest {
	s.ConnectedInstanceId = &v
	return s
}

func (s *DeleteConnectedClusterRequest) Validate() error {
	return dara.Validate(s)
}
