// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOfficeSiteSsoStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOfficeSiteId(v string) *GetOfficeSiteSsoStatusRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *GetOfficeSiteSsoStatusRequest
	GetRegionId() *string
}

type GetOfficeSiteSsoStatusRequest struct {
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetOfficeSiteSsoStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOfficeSiteSsoStatusRequest) GoString() string {
	return s.String()
}

func (s *GetOfficeSiteSsoStatusRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *GetOfficeSiteSsoStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetOfficeSiteSsoStatusRequest) SetOfficeSiteId(v string) *GetOfficeSiteSsoStatusRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *GetOfficeSiteSsoStatusRequest) SetRegionId(v string) *GetOfficeSiteSsoStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetOfficeSiteSsoStatusRequest) Validate() error {
	return dara.Validate(s)
}
