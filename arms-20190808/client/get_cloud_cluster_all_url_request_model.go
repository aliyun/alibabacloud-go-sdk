// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudClusterAllUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetCloudClusterAllUrlRequest
	GetClusterId() *string
	SetRegionId(v string) *GetCloudClusterAllUrlRequest
	GetRegionId() *string
}

type GetCloudClusterAllUrlRequest struct {
	// The ID of the CloudMonitor instance.
	//
	// example:
	//
	// ca9676014babd4
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetCloudClusterAllUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCloudClusterAllUrlRequest) GoString() string {
	return s.String()
}

func (s *GetCloudClusterAllUrlRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetCloudClusterAllUrlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCloudClusterAllUrlRequest) SetClusterId(v string) *GetCloudClusterAllUrlRequest {
	s.ClusterId = &v
	return s
}

func (s *GetCloudClusterAllUrlRequest) SetRegionId(v string) *GetCloudClusterAllUrlRequest {
	s.RegionId = &v
	return s
}

func (s *GetCloudClusterAllUrlRequest) Validate() error {
	return dara.Validate(s)
}
