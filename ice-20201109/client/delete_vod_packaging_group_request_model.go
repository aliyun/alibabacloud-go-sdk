// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodPackagingGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *DeleteVodPackagingGroupRequest
	GetGroupName() *string
}

type DeleteVodPackagingGroupRequest struct {
	// The name of the packaging group.
	//
	// example:
	//
	// vod_hls
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DeleteVodPackagingGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodPackagingGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteVodPackagingGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteVodPackagingGroupRequest) SetGroupName(v string) *DeleteVodPackagingGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteVodPackagingGroupRequest) Validate() error {
	return dara.Validate(s)
}
