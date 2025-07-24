// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterCloneMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetClusterCloneMetaRequest
	GetClusterId() *string
	SetRegionId(v string) *GetClusterCloneMetaRequest
	GetRegionId() *string
}

type GetClusterCloneMetaRequest struct {
	// The cluster ID.
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

func (s GetClusterCloneMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCloneMetaRequest) GoString() string {
	return s.String()
}

func (s *GetClusterCloneMetaRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterCloneMetaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetClusterCloneMetaRequest) SetClusterId(v string) *GetClusterCloneMetaRequest {
	s.ClusterId = &v
	return s
}

func (s *GetClusterCloneMetaRequest) SetRegionId(v string) *GetClusterCloneMetaRequest {
	s.RegionId = &v
	return s
}

func (s *GetClusterCloneMetaRequest) Validate() error {
	return dara.Validate(s)
}
