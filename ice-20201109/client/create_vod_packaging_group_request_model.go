// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodPackagingGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateVodPackagingGroupRequest
	GetDescription() *string
	SetGroupName(v string) *CreateVodPackagingGroupRequest
	GetGroupName() *string
}

type CreateVodPackagingGroupRequest struct {
	// The packaging group description.
	//
	// example:
	//
	// vod hls packaging
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the packaging group. The name must be unique in an account and can be up to 128 characters in length. Letters, digits, underscores (_), and hyphens (-) are supported.
	//
	// example:
	//
	// vod_hls
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s CreateVodPackagingGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVodPackagingGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateVodPackagingGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVodPackagingGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateVodPackagingGroupRequest) SetDescription(v string) *CreateVodPackagingGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateVodPackagingGroupRequest) SetGroupName(v string) *CreateVodPackagingGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateVodPackagingGroupRequest) Validate() error {
	return dara.Validate(s)
}
