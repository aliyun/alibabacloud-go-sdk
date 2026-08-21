// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteWorkspaceQueueRequest
	GetRegionId() *string
}

type DeleteWorkspaceQueueRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DeleteWorkspaceQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceQueueRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceQueueRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteWorkspaceQueueRequest) SetRegionId(v string) *DeleteWorkspaceQueueRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteWorkspaceQueueRequest) Validate() error {
	return dara.Validate(s)
}
