// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentProfilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentProfileIds(v []*string) *DeleteAgentProfilesRequest
	GetAgentProfileIds() []*string
	SetAppIp(v string) *DeleteAgentProfilesRequest
	GetAppIp() *string
}

type DeleteAgentProfilesRequest struct {
	// Collection of agent configuration IDs
	AgentProfileIds []*string `json:"AgentProfileIds,omitempty" xml:"AgentProfileIds,omitempty" type:"Repeated"`
	// app_ip (system field, optional)
	//
	// example:
	//
	// 127.0.0.1
	AppIp *string `json:"AppIp,omitempty" xml:"AppIp,omitempty"`
}

func (s DeleteAgentProfilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentProfilesRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentProfilesRequest) GetAgentProfileIds() []*string {
	return s.AgentProfileIds
}

func (s *DeleteAgentProfilesRequest) GetAppIp() *string {
	return s.AppIp
}

func (s *DeleteAgentProfilesRequest) SetAgentProfileIds(v []*string) *DeleteAgentProfilesRequest {
	s.AgentProfileIds = v
	return s
}

func (s *DeleteAgentProfilesRequest) SetAppIp(v string) *DeleteAgentProfilesRequest {
	s.AppIp = &v
	return s
}

func (s *DeleteAgentProfilesRequest) Validate() error {
	return dara.Validate(s)
}
