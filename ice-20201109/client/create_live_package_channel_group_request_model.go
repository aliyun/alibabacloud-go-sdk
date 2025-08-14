// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePackageChannelGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateLivePackageChannelGroupRequest
	GetClientToken() *string
	SetDescription(v string) *CreateLivePackageChannelGroupRequest
	GetDescription() *string
	SetGroupName(v string) *CreateLivePackageChannelGroupRequest
	GetGroupName() *string
}

type CreateLivePackageChannelGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ****12e8864746a0a398****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The channel group description. It can be up to 1,000 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The channel group name. It can contain letters, digits, hyphens (-), and underscores (_). The name must be 1 to 200 characters in length. Format: [A-Za-z0-9_-]+
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-group-01
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s CreateLivePackageChannelGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePackageChannelGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateLivePackageChannelGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateLivePackageChannelGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLivePackageChannelGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateLivePackageChannelGroupRequest) SetClientToken(v string) *CreateLivePackageChannelGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLivePackageChannelGroupRequest) SetDescription(v string) *CreateLivePackageChannelGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateLivePackageChannelGroupRequest) SetGroupName(v string) *CreateLivePackageChannelGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateLivePackageChannelGroupRequest) Validate() error {
	return dara.Validate(s)
}
