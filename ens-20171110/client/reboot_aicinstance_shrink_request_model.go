// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootAICInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RebootAICInstanceShrinkRequest
	GetInstanceId() *string
	SetInstanceIdsShrink(v string) *RebootAICInstanceShrinkRequest
	GetInstanceIdsShrink() *string
	SetServerId(v string) *RebootAICInstanceShrinkRequest
	GetServerId() *string
}

type RebootAICInstanceShrinkRequest struct {
	// The ID of the AIC instance.
	//
	// example:
	//
	// aic-instance****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the AIC instance groups.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The ID of the server.
	//
	// example:
	//
	// cas-instance****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
}

func (s RebootAICInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootAICInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *RebootAICInstanceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RebootAICInstanceShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *RebootAICInstanceShrinkRequest) GetServerId() *string {
	return s.ServerId
}

func (s *RebootAICInstanceShrinkRequest) SetInstanceId(v string) *RebootAICInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *RebootAICInstanceShrinkRequest) SetInstanceIdsShrink(v string) *RebootAICInstanceShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *RebootAICInstanceShrinkRequest) SetServerId(v string) *RebootAICInstanceShrinkRequest {
	s.ServerId = &v
	return s
}

func (s *RebootAICInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
