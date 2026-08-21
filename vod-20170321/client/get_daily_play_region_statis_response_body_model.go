// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDailyPlayRegionStatisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDailyPlayRegionStatisList(v []*GetDailyPlayRegionStatisResponseBodyDailyPlayRegionStatisList) *GetDailyPlayRegionStatisResponseBody
	GetDailyPlayRegionStatisList() []*GetDailyPlayRegionStatisResponseBodyDailyPlayRegionStatisList
	SetEmptyDates(v []*string) *GetDailyPlayRegionStatisResponseBody
	GetEmptyDates() []*string
	SetFailDates(v []*string) *GetDailyPlayRegionStatisResponseBody
	GetFailDates() []*string
	SetRequestId(v string) *GetDailyPlayRegionStatisResponseBody
	GetRequestId() *string
}

type GetDailyPlayRegionStatisResponseBody struct {
	// The CSV file information returned.
	DailyPlayRegionStatisList []*GetDailyPlayRegionStatisResponseBodyDailyPlayRegionStatisList `json:"DailyPlayRegionStatisList,omitempty" xml:"DailyPlayRegionStatisList,omitempty" type:"Repeated"`
	// If no playback data exists on a specific day, the date is displayed in this field.
	EmptyDates []*string `json:"EmptyDates,omitempty" xml:"EmptyDates,omitempty" type:"Repeated"`
	// If the playback data export fails on a specific day, the date is displayed in this field.
	FailDates []*string `json:"FailDates,omitempty" xml:"FailDates,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 67720502-CDDB-3F1C-8A91-12258*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDailyPlayRegionStatisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDailyPlayRegionStatisResponseBody) GoString() string {
	return s.String()
}

func (s *GetDailyPlayRegionStatisResponseBody) GetDailyPlayRegionStatisList() []*GetDailyPlayRegionStatisResponseBodyDailyPlayRegionStatisList {
	return s.DailyPlayRegionStatisList
}

func (s *GetDailyPlayRegionStatisResponseBody) GetEmptyDates() []*string {
	return s.EmptyDates
}

func (s *GetDailyPlayRegionStatisResponseBody) GetFailDates() []*string {
	return s.FailDates
}

func (s *GetDailyPlayRegionStatisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDailyPlayRegionStatisResponseBody) SetDailyPlayRegionStatisList(v []*GetDailyPlayRegionStatisResponseBodyDailyPlayRegionStatisList) *GetDailyPlayRegionStatisResponseBody {
	s.DailyPlayRegionStatisList = v
	return s
}

func (s *GetDailyPlayRegionStatisResponseBody) SetEmptyDates(v []*string) *GetDailyPlayRegionStatisResponseBody {
	s.EmptyDates = v
	return s
}

func (s *GetDailyPlayRegionStatisResponseBody) SetFailDates(v []*string) *GetDailyPlayRegionStatisResponseBody {
	s.FailDates = v
	return s
}

func (s *GetDailyPlayRegionStatisResponseBody) SetRequestId(v string) *GetDailyPlayRegionStatisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDailyPlayRegionStatisResponseBody) Validate() error {
	if s.DailyPlayRegionStatisList != nil {
		for _, item := range s.DailyPlayRegionStatisList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDailyPlayRegionStatisResponseBodyDailyPlayRegionStatisList struct {
	// The playback date, in days. Format: yyyy-MM-dd.
	//
	// example:
	//
	// 2025-03-20
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The download URL of the CSV file.
	//
	// example:
	//
	// https://outin-e70266d4ed*******0163e1403e7.oss-cn-shanghai.aliyuncs.com/dataexport/play/cn_hangzhou_20250320_video_traffic.csv?*******
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s GetDailyPlayRegionStatisResponseBodyDailyPlayRegionStatisList) String() string {
	return dara.Prettify(s)
}

func (s GetDailyPlayRegionStatisResponseBodyDailyPlayRegionStatisList) GoString() string {
	return s.String()
}

func (s *GetDailyPlayRegionStatisResponseBodyDailyPlayRegionStatisList) GetDate() *string {
	return s.Date
}

func (s *GetDailyPlayRegionStatisResponseBodyDailyPlayRegionStatisList) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetDailyPlayRegionStatisResponseBodyDailyPlayRegionStatisList) SetDate(v string) *GetDailyPlayRegionStatisResponseBodyDailyPlayRegionStatisList {
	s.Date = &v
	return s
}

func (s *GetDailyPlayRegionStatisResponseBodyDailyPlayRegionStatisList) SetFileUrl(v string) *GetDailyPlayRegionStatisResponseBodyDailyPlayRegionStatisList {
	s.FileUrl = &v
	return s
}

func (s *GetDailyPlayRegionStatisResponseBodyDailyPlayRegionStatisList) Validate() error {
	return dara.Validate(s)
}
