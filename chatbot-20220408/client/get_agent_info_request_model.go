// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetAgentInfoRequest
	GetInstanceId() *string
}

type GetAgentInfoRequest struct {
	// example:
	//
	// beebot_bot_public_cn-ca36x8v3n1x
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAgentInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAgentInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAgentInfoRequest) SetInstanceId(v string) *GetAgentInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentInfoRequest) Validate() error {
	return dara.Validate(s)
}
