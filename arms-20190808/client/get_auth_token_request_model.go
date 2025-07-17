// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetAuthTokenRequest
	GetClusterId() *string
	SetRegionId(v string) *GetAuthTokenRequest
	GetRegionId() *string
}

type GetAuthTokenRequest struct {
	// The ID of the ACK cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAuthTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *GetAuthTokenRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetAuthTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAuthTokenRequest) SetClusterId(v string) *GetAuthTokenRequest {
	s.ClusterId = &v
	return s
}

func (s *GetAuthTokenRequest) SetRegionId(v string) *GetAuthTokenRequest {
	s.RegionId = &v
	return s
}

func (s *GetAuthTokenRequest) Validate() error {
	return dara.Validate(s)
}
