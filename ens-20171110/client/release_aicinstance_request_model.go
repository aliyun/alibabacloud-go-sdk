// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseAICInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServerId(v string) *ReleaseAICInstanceRequest
	GetServerId() *string
}

type ReleaseAICInstanceRequest struct {
	// The ID of the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cas-instance****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
}

func (s ReleaseAICInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseAICInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseAICInstanceRequest) GetServerId() *string {
	return s.ServerId
}

func (s *ReleaseAICInstanceRequest) SetServerId(v string) *ReleaseAICInstanceRequest {
	s.ServerId = &v
	return s
}

func (s *ReleaseAICInstanceRequest) Validate() error {
	return dara.Validate(s)
}
