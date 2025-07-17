// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSyntheticTaskMonitorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetSyntheticTaskMonitorsRequest
	GetRegionId() *string
}

type GetSyntheticTaskMonitorsRequest struct {
	// The ID of the region in which the application is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetSyntheticTaskMonitorsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskMonitorsRequest) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskMonitorsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSyntheticTaskMonitorsRequest) SetRegionId(v string) *GetSyntheticTaskMonitorsRequest {
	s.RegionId = &v
	return s
}

func (s *GetSyntheticTaskMonitorsRequest) Validate() error {
	return dara.Validate(s)
}
