// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPromInstallStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *QueryPromInstallStatusRequest
	GetClusterId() *string
	SetRegionId(v string) *QueryPromInstallStatusRequest
	GetRegionId() *string
}

type QueryPromInstallStatusRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryPromInstallStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPromInstallStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryPromInstallStatusRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *QueryPromInstallStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryPromInstallStatusRequest) SetClusterId(v string) *QueryPromInstallStatusRequest {
	s.ClusterId = &v
	return s
}

func (s *QueryPromInstallStatusRequest) SetRegionId(v string) *QueryPromInstallStatusRequest {
	s.RegionId = &v
	return s
}

func (s *QueryPromInstallStatusRequest) Validate() error {
	return dara.Validate(s)
}
