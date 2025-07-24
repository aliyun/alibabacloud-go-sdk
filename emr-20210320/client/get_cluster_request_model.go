// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetClusterRequest
	GetClusterId() *string
	SetRegionId(v string) *GetClusterRequest
	GetRegionId() *string
}

type GetClusterRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRequest) GoString() string {
	return s.String()
}

func (s *GetClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetClusterRequest) SetClusterId(v string) *GetClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *GetClusterRequest) SetRegionId(v string) *GetClusterRequest {
	s.RegionId = &v
	return s
}

func (s *GetClusterRequest) Validate() error {
	return dara.Validate(s)
}
