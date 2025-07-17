// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelAuthTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DelAuthTokenRequest
	GetClusterId() *string
	SetRegionId(v string) *DelAuthTokenRequest
	GetRegionId() *string
}

type DelAuthTokenRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
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

func (s DelAuthTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s DelAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *DelAuthTokenRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DelAuthTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DelAuthTokenRequest) SetClusterId(v string) *DelAuthTokenRequest {
	s.ClusterId = &v
	return s
}

func (s *DelAuthTokenRequest) SetRegionId(v string) *DelAuthTokenRequest {
	s.RegionId = &v
	return s
}

func (s *DelAuthTokenRequest) Validate() error {
	return dara.Validate(s)
}
