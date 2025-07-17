// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterAllUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetClusterAllUrlRequest
	GetClusterId() *string
	SetRegionId(v string) *GetClusterAllUrlRequest
	GetRegionId() *string
}

type GetClusterAllUrlRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID. Default value: cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetClusterAllUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClusterAllUrlRequest) GoString() string {
	return s.String()
}

func (s *GetClusterAllUrlRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterAllUrlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetClusterAllUrlRequest) SetClusterId(v string) *GetClusterAllUrlRequest {
	s.ClusterId = &v
	return s
}

func (s *GetClusterAllUrlRequest) SetRegionId(v string) *GetClusterAllUrlRequest {
	s.RegionId = &v
	return s
}

func (s *GetClusterAllUrlRequest) Validate() error {
	return dara.Validate(s)
}
