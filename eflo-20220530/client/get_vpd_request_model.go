// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetVpdRequest
	GetRegionId() *string
	SetVpdId(v string) *GetVpdRequest
	GetVpdId() *string
}

type GetVpdRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the VPD instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-ze3na0wf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s GetVpdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVpdRequest) GoString() string {
	return s.String()
}

func (s *GetVpdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpdRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *GetVpdRequest) SetRegionId(v string) *GetVpdRequest {
	s.RegionId = &v
	return s
}

func (s *GetVpdRequest) SetVpdId(v string) *GetVpdRequest {
	s.VpdId = &v
	return s
}

func (s *GetVpdRequest) Validate() error {
	return dara.Validate(s)
}
