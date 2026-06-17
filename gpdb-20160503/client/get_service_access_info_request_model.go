// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceAccessInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetServiceAccessInfoRequest
	GetRegionId() *string
	SetServiceId(v string) *GetServiceAccessInfoRequest
	GetServiceId() *string
}

type GetServiceAccessInfoRequest struct {
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agdb-xxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s GetServiceAccessInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceAccessInfoRequest) GoString() string {
	return s.String()
}

func (s *GetServiceAccessInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceAccessInfoRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceAccessInfoRequest) SetRegionId(v string) *GetServiceAccessInfoRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceAccessInfoRequest) SetServiceId(v string) *GetServiceAccessInfoRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceAccessInfoRequest) Validate() error {
	return dara.Validate(s)
}
