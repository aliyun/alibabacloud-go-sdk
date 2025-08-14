// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVodPackagingGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *GetVodPackagingGroupRequest
	GetGroupName() *string
}

type GetVodPackagingGroupRequest struct {
	// The name of the packaging group. The name must be unique and can be up to 128 characters in length. Letters, digits, underscores (_), and hyphens (-) are supported.
	//
	// example:
	//
	// vod_hls
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetVodPackagingGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVodPackagingGroupRequest) GoString() string {
	return s.String()
}

func (s *GetVodPackagingGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *GetVodPackagingGroupRequest) SetGroupName(v string) *GetVodPackagingGroupRequest {
	s.GroupName = &v
	return s
}

func (s *GetVodPackagingGroupRequest) Validate() error {
	return dara.Validate(s)
}
