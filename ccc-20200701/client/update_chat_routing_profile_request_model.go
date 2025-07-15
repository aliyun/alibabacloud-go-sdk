// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatRoutingProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateChatRoutingProfileRequest
	GetInstanceId() *string
	SetRoutingProfiles(v string) *UpdateChatRoutingProfileRequest
	GetRoutingProfiles() *string
}

type UpdateChatRoutingProfileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 9cfad875-6260-4a53-ab6e-b13e3fb31f7d
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "RoutingType": "Automatic",
	//
	//     "AgentConcurrencySettings": {
	//
	//         "AllowExceedingLimitWhenTransferring": false,
	//
	//         "ConcurrencyLimit": 4,
	//
	//         "AllowExceedingLimitWhenClaiming": true,
	//
	//         "Enabled": true
	//
	//     },
	//
	//     "ChatSettings": {
	//
	//         "IdleChatTimeoutSeconds": 300
	//
	//     },
	//
	//     "DistributionSettings": {
	//
	//         "Enabled": true,
	//
	//         "AgentRingTimeoutSeconds": 119,
	//
	//         "MaxNumberOfConversationsAgentCanMiss": 5,
	//
	//         "PostAgentMissingConversionsAction": "Nothing"
	//
	//     }
	//
	// }
	RoutingProfiles *string `json:"RoutingProfiles,omitempty" xml:"RoutingProfiles,omitempty"`
}

func (s UpdateChatRoutingProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatRoutingProfileRequest) GoString() string {
	return s.String()
}

func (s *UpdateChatRoutingProfileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateChatRoutingProfileRequest) GetRoutingProfiles() *string {
	return s.RoutingProfiles
}

func (s *UpdateChatRoutingProfileRequest) SetInstanceId(v string) *UpdateChatRoutingProfileRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateChatRoutingProfileRequest) SetRoutingProfiles(v string) *UpdateChatRoutingProfileRequest {
	s.RoutingProfiles = &v
	return s
}

func (s *UpdateChatRoutingProfileRequest) Validate() error {
	return dara.Validate(s)
}
