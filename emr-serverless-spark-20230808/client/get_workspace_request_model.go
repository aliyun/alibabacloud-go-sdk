// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetWorkspaceRequest
	GetRegionId() *string
}

type GetWorkspaceRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetWorkspaceRequest) SetRegionId(v string) *GetWorkspaceRequest {
	s.RegionId = &v
	return s
}

func (s *GetWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
