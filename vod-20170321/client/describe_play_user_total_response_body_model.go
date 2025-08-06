// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayUserTotalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePlayUserTotalResponseBody
	GetRequestId() *string
	SetUserPlayStatisTotals(v *DescribePlayUserTotalResponseBodyUserPlayStatisTotals) *DescribePlayUserTotalResponseBody
	GetUserPlayStatisTotals() *DescribePlayUserTotalResponseBodyUserPlayStatisTotals
}

type DescribePlayUserTotalResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1FAFB884-D5A7-47D1-****-8928AA9C8720
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The daily playback statistics.
	UserPlayStatisTotals *DescribePlayUserTotalResponseBodyUserPlayStatisTotals `json:"UserPlayStatisTotals,omitempty" xml:"UserPlayStatisTotals,omitempty" type:"Struct"`
}

func (s DescribePlayUserTotalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayUserTotalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlayUserTotalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlayUserTotalResponseBody) GetUserPlayStatisTotals() *DescribePlayUserTotalResponseBodyUserPlayStatisTotals {
	return s.UserPlayStatisTotals
}

func (s *DescribePlayUserTotalResponseBody) SetRequestId(v string) *DescribePlayUserTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlayUserTotalResponseBody) SetUserPlayStatisTotals(v *DescribePlayUserTotalResponseBodyUserPlayStatisTotals) *DescribePlayUserTotalResponseBody {
	s.UserPlayStatisTotals = v
	return s
}

func (s *DescribePlayUserTotalResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePlayUserTotalResponseBodyUserPlayStatisTotals struct {
	UserPlayStatisTotal []*DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal `json:"UserPlayStatisTotal,omitempty" xml:"UserPlayStatisTotal,omitempty" type:"Repeated"`
}

func (s DescribePlayUserTotalResponseBodyUserPlayStatisTotals) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayUserTotalResponseBodyUserPlayStatisTotals) GoString() string {
	return s.String()
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotals) GetUserPlayStatisTotal() []*DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal {
	return s.UserPlayStatisTotal
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotals) SetUserPlayStatisTotal(v []*DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) *DescribePlayUserTotalResponseBodyUserPlayStatisTotals {
	s.UserPlayStatisTotal = v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotals) Validate() error {
	return dara.Validate(s)
}

type DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal struct {
	// The date. The date is displayed in the yyyy-MM-dd format.
	//
	// example:
	//
	// 20170120
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The total playback duration. Unit: milliseconds.
	//
	// example:
	//
	// 9340070
	PlayDuration *string `json:"PlayDuration,omitempty" xml:"PlayDuration,omitempty"`
	// The distribution of the playback duration.
	//
	// example:
	//
	// "<=1m:74.3%;>1<=5m:22.8%;>5<=10m:1.0%;>10<=15m:1.0%;>15<=30m:1.0%"
	PlayRange *string `json:"PlayRange,omitempty" xml:"PlayRange,omitempty"`
	// The total number of unique visitors.
	UV *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV `json:"UV,omitempty" xml:"UV,omitempty" type:"Struct"`
	// The total number of video views.
	VV *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV `json:"VV,omitempty" xml:"VV,omitempty" type:"Struct"`
}

func (s DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) GoString() string {
	return s.String()
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) GetDate() *string {
	return s.Date
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) GetPlayDuration() *string {
	return s.PlayDuration
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) GetPlayRange() *string {
	return s.PlayRange
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) GetUV() *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV {
	return s.UV
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) GetVV() *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV {
	return s.VV
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) SetDate(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal {
	s.Date = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) SetPlayDuration(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal {
	s.PlayDuration = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) SetPlayRange(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal {
	s.PlayRange = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) SetUV(v *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal {
	s.UV = v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) SetVV(v *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal {
	s.VV = v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) Validate() error {
	return dara.Validate(s)
}

type DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV struct {
	// The total number of unique visitors who use ApsaraVideo Player SDK for Android.
	//
	// example:
	//
	// 2
	Android *string `json:"Android,omitempty" xml:"Android,omitempty"`
	// The total number of unique visitors who use the Flash player.
	//
	// example:
	//
	// 1
	Flash *string `json:"Flash,omitempty" xml:"Flash,omitempty"`
	// The total number of unique visitors who use the HTML5 player.
	//
	// example:
	//
	// 1
	HTML5 *string `json:"HTML5,omitempty" xml:"HTML5,omitempty"`
	// The total number of unique visitors who use ApsaraVideo Player SDK for iOS.
	//
	// example:
	//
	// 0
	IOS *string `json:"iOS,omitempty" xml:"iOS,omitempty"`
}

func (s DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) GoString() string {
	return s.String()
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) GetAndroid() *string {
	return s.Android
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) GetFlash() *string {
	return s.Flash
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) GetHTML5() *string {
	return s.HTML5
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) GetIOS() *string {
	return s.IOS
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) SetAndroid(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV {
	s.Android = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) SetFlash(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV {
	s.Flash = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) SetHTML5(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV {
	s.HTML5 = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) SetIOS(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV {
	s.IOS = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) Validate() error {
	return dara.Validate(s)
}

type DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV struct {
	// The total number of video views played by using ApsaraVideo Player SDK for Android.
	//
	// example:
	//
	// 161
	Android *string `json:"Android,omitempty" xml:"Android,omitempty"`
	// The total number of video views played by using the Flash player.
	//
	// example:
	//
	// 2
	Flash *string `json:"Flash,omitempty" xml:"Flash,omitempty"`
	// The total number of video views played by using the HTML5 player.
	//
	// example:
	//
	// 2
	HTML5 *string `json:"HTML5,omitempty" xml:"HTML5,omitempty"`
	// The total number of video views played by using ApsaraVideo Player SDK for iOS.
	//
	// example:
	//
	// 0
	IOS *string `json:"iOS,omitempty" xml:"iOS,omitempty"`
}

func (s DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) GoString() string {
	return s.String()
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) GetAndroid() *string {
	return s.Android
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) GetFlash() *string {
	return s.Flash
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) GetHTML5() *string {
	return s.HTML5
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) GetIOS() *string {
	return s.IOS
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) SetAndroid(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV {
	s.Android = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) SetFlash(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV {
	s.Flash = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) SetHTML5(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV {
	s.HTML5 = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) SetIOS(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV {
	s.IOS = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) Validate() error {
	return dara.Validate(s)
}
