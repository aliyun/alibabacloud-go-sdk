// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootAICInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RebootAICInstanceRequest
	GetInstanceId() *string
	SetInstanceIds(v []*string) *RebootAICInstanceRequest
	GetInstanceIds() []*string
	SetServerId(v string) *RebootAICInstanceRequest
	GetServerId() *string
}

type RebootAICInstanceRequest struct {
	// The ID of the AIC instance.
	//
	// example:
	//
	// aic-instance****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the AIC instance groups.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The ID of the server.
	//
	// example:
	//
	// cas-instance****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
}

func (s RebootAICInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootAICInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebootAICInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RebootAICInstanceRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RebootAICInstanceRequest) GetServerId() *string {
	return s.ServerId
}

func (s *RebootAICInstanceRequest) SetInstanceId(v string) *RebootAICInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RebootAICInstanceRequest) SetInstanceIds(v []*string) *RebootAICInstanceRequest {
	s.InstanceIds = v
	return s
}

func (s *RebootAICInstanceRequest) SetServerId(v string) *RebootAICInstanceRequest {
	s.ServerId = &v
	return s
}

func (s *RebootAICInstanceRequest) Validate() error {
	return dara.Validate(s)
}
