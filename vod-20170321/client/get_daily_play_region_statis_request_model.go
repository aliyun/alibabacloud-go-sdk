// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDailyPlayRegionStatisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *GetDailyPlayRegionStatisRequest
	GetDate() *string
	SetMediaRegion(v string) *GetDailyPlayRegionStatisRequest
	GetMediaRegion() *string
}

type GetDailyPlayRegionStatisRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2025-03-20
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	MediaRegion *string `json:"MediaRegion,omitempty" xml:"MediaRegion,omitempty"`
}

func (s GetDailyPlayRegionStatisRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDailyPlayRegionStatisRequest) GoString() string {
	return s.String()
}

func (s *GetDailyPlayRegionStatisRequest) GetDate() *string {
	return s.Date
}

func (s *GetDailyPlayRegionStatisRequest) GetMediaRegion() *string {
	return s.MediaRegion
}

func (s *GetDailyPlayRegionStatisRequest) SetDate(v string) *GetDailyPlayRegionStatisRequest {
	s.Date = &v
	return s
}

func (s *GetDailyPlayRegionStatisRequest) SetMediaRegion(v string) *GetDailyPlayRegionStatisRequest {
	s.MediaRegion = &v
	return s
}

func (s *GetDailyPlayRegionStatisRequest) Validate() error {
	return dara.Validate(s)
}
