// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDailyPlayStatisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetDailyPlayStatisRequest
	GetAppId() *string
	SetEndDate(v string) *GetDailyPlayStatisRequest
	GetEndDate() *string
	SetMediaId(v string) *GetDailyPlayStatisRequest
	GetMediaId() *string
	SetMediaRegion(v string) *GetDailyPlayStatisRequest
	GetMediaRegion() *string
	SetStartDate(v string) *GetDailyPlayStatisRequest
	GetStartDate() *string
}

type GetDailyPlayStatisRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	MediaId     *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	MediaRegion *string `json:"MediaRegion,omitempty" xml:"MediaRegion,omitempty"`
	// This parameter is required.
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetDailyPlayStatisRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDailyPlayStatisRequest) GoString() string {
	return s.String()
}

func (s *GetDailyPlayStatisRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetDailyPlayStatisRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetDailyPlayStatisRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetDailyPlayStatisRequest) GetMediaRegion() *string {
	return s.MediaRegion
}

func (s *GetDailyPlayStatisRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetDailyPlayStatisRequest) SetAppId(v string) *GetDailyPlayStatisRequest {
	s.AppId = &v
	return s
}

func (s *GetDailyPlayStatisRequest) SetEndDate(v string) *GetDailyPlayStatisRequest {
	s.EndDate = &v
	return s
}

func (s *GetDailyPlayStatisRequest) SetMediaId(v string) *GetDailyPlayStatisRequest {
	s.MediaId = &v
	return s
}

func (s *GetDailyPlayStatisRequest) SetMediaRegion(v string) *GetDailyPlayStatisRequest {
	s.MediaRegion = &v
	return s
}

func (s *GetDailyPlayStatisRequest) SetStartDate(v string) *GetDailyPlayStatisRequest {
	s.StartDate = &v
	return s
}

func (s *GetDailyPlayStatisRequest) Validate() error {
	return dara.Validate(s)
}
