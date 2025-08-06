// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDailyPlayStatisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDailyPlayStatisList(v []*GetDailyPlayStatisResponseBodyDailyPlayStatisList) *GetDailyPlayStatisResponseBody
	GetDailyPlayStatisList() []*GetDailyPlayStatisResponseBodyDailyPlayStatisList
	SetMediaId(v string) *GetDailyPlayStatisResponseBody
	GetMediaId() *string
	SetRequestId(v string) *GetDailyPlayStatisResponseBody
	GetRequestId() *string
}

type GetDailyPlayStatisResponseBody struct {
	DailyPlayStatisList []*GetDailyPlayStatisResponseBodyDailyPlayStatisList `json:"DailyPlayStatisList,omitempty" xml:"DailyPlayStatisList,omitempty" type:"Repeated"`
	MediaId             *string                                              `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	RequestId           *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDailyPlayStatisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDailyPlayStatisResponseBody) GoString() string {
	return s.String()
}

func (s *GetDailyPlayStatisResponseBody) GetDailyPlayStatisList() []*GetDailyPlayStatisResponseBodyDailyPlayStatisList {
	return s.DailyPlayStatisList
}

func (s *GetDailyPlayStatisResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *GetDailyPlayStatisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDailyPlayStatisResponseBody) SetDailyPlayStatisList(v []*GetDailyPlayStatisResponseBodyDailyPlayStatisList) *GetDailyPlayStatisResponseBody {
	s.DailyPlayStatisList = v
	return s
}

func (s *GetDailyPlayStatisResponseBody) SetMediaId(v string) *GetDailyPlayStatisResponseBody {
	s.MediaId = &v
	return s
}

func (s *GetDailyPlayStatisResponseBody) SetRequestId(v string) *GetDailyPlayStatisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDailyPlayStatisResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDailyPlayStatisResponseBodyDailyPlayStatisList struct {
	AppId     *int64   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Date      *string  `json:"Date,omitempty" xml:"Date,omitempty"`
	Flow      *float64 `json:"Flow,omitempty" xml:"Flow,omitempty"`
	MediaId   *string  `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	PlayCount *int64   `json:"PlayCount,omitempty" xml:"PlayCount,omitempty"`
	UserId    *int64   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetDailyPlayStatisResponseBodyDailyPlayStatisList) String() string {
	return dara.Prettify(s)
}

func (s GetDailyPlayStatisResponseBodyDailyPlayStatisList) GoString() string {
	return s.String()
}

func (s *GetDailyPlayStatisResponseBodyDailyPlayStatisList) GetAppId() *int64 {
	return s.AppId
}

func (s *GetDailyPlayStatisResponseBodyDailyPlayStatisList) GetDate() *string {
	return s.Date
}

func (s *GetDailyPlayStatisResponseBodyDailyPlayStatisList) GetFlow() *float64 {
	return s.Flow
}

func (s *GetDailyPlayStatisResponseBodyDailyPlayStatisList) GetMediaId() *string {
	return s.MediaId
}

func (s *GetDailyPlayStatisResponseBodyDailyPlayStatisList) GetPlayCount() *int64 {
	return s.PlayCount
}

func (s *GetDailyPlayStatisResponseBodyDailyPlayStatisList) GetUserId() *int64 {
	return s.UserId
}

func (s *GetDailyPlayStatisResponseBodyDailyPlayStatisList) SetAppId(v int64) *GetDailyPlayStatisResponseBodyDailyPlayStatisList {
	s.AppId = &v
	return s
}

func (s *GetDailyPlayStatisResponseBodyDailyPlayStatisList) SetDate(v string) *GetDailyPlayStatisResponseBodyDailyPlayStatisList {
	s.Date = &v
	return s
}

func (s *GetDailyPlayStatisResponseBodyDailyPlayStatisList) SetFlow(v float64) *GetDailyPlayStatisResponseBodyDailyPlayStatisList {
	s.Flow = &v
	return s
}

func (s *GetDailyPlayStatisResponseBodyDailyPlayStatisList) SetMediaId(v string) *GetDailyPlayStatisResponseBodyDailyPlayStatisList {
	s.MediaId = &v
	return s
}

func (s *GetDailyPlayStatisResponseBodyDailyPlayStatisList) SetPlayCount(v int64) *GetDailyPlayStatisResponseBodyDailyPlayStatisList {
	s.PlayCount = &v
	return s
}

func (s *GetDailyPlayStatisResponseBodyDailyPlayStatisList) SetUserId(v int64) *GetDailyPlayStatisResponseBodyDailyPlayStatisList {
	s.UserId = &v
	return s
}

func (s *GetDailyPlayStatisResponseBodyDailyPlayStatisList) Validate() error {
	return dara.Validate(s)
}
