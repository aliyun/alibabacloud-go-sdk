// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAICInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ResetAICInstanceRequest
	GetInstanceId() *string
	SetServerId(v string) *ResetAICInstanceRequest
	GetServerId() *string
}

type ResetAICInstanceRequest struct {
	// The ID of the AIC instance.
	//
	// example:
	//
	// aic-instance****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the server.
	//
	// example:
	//
	// cas-instance****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
}

func (s ResetAICInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAICInstanceRequest) GoString() string {
	return s.String()
}

func (s *ResetAICInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetAICInstanceRequest) GetServerId() *string {
	return s.ServerId
}

func (s *ResetAICInstanceRequest) SetInstanceId(v string) *ResetAICInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetAICInstanceRequest) SetServerId(v string) *ResetAICInstanceRequest {
	s.ServerId = &v
	return s
}

func (s *ResetAICInstanceRequest) Validate() error {
	return dara.Validate(s)
}
