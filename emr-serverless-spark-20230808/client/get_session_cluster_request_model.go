// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetSessionClusterRequest
	GetRegionId() *string
}

type GetSessionClusterRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetSessionClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSessionClusterRequest) GoString() string {
	return s.String()
}

func (s *GetSessionClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSessionClusterRequest) SetRegionId(v string) *GetSessionClusterRequest {
	s.RegionId = &v
	return s
}

func (s *GetSessionClusterRequest) Validate() error {
	return dara.Validate(s)
}
