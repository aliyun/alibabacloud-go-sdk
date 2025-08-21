// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverAICInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServerId(v string) *RecoverAICInstanceRequest
	GetServerId() *string
}

type RecoverAICInstanceRequest struct {
	// The ID of the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cas-instance****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
}

func (s RecoverAICInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecoverAICInstanceRequest) GoString() string {
	return s.String()
}

func (s *RecoverAICInstanceRequest) GetServerId() *string {
	return s.ServerId
}

func (s *RecoverAICInstanceRequest) SetServerId(v string) *RecoverAICInstanceRequest {
	s.ServerId = &v
	return s
}

func (s *RecoverAICInstanceRequest) Validate() error {
	return dara.Validate(s)
}
