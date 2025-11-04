// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterImsEditUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeMeterImsEditUsageResponseBodyData) *DescribeMeterImsEditUsageResponseBody
	GetData() []*DescribeMeterImsEditUsageResponseBodyData
	SetRequestId(v string) *DescribeMeterImsEditUsageResponseBody
	GetRequestId() *string
}

type DescribeMeterImsEditUsageResponseBody struct {
	// The usage statistics of IMS on VOD editing.
	Data []*DescribeMeterImsEditUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7F3AE2C6-5CC6-5712-BAC5-5A735A157687
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeterImsEditUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsEditUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsEditUsageResponseBody) GetData() []*DescribeMeterImsEditUsageResponseBodyData {
	return s.Data
}

func (s *DescribeMeterImsEditUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMeterImsEditUsageResponseBody) SetData(v []*DescribeMeterImsEditUsageResponseBodyData) *DescribeMeterImsEditUsageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMeterImsEditUsageResponseBody) SetRequestId(v string) *DescribeMeterImsEditUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMeterImsEditUsageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMeterImsEditUsageResponseBodyData struct {
	// The usage duration, in minutes.
	//
	// example:
	//
	// 1.23
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The video profile.
	//
	// example:
	//
	// 1080P
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The beginning time of usage. The value is a 10-digit timestamp.
	//
	// example:
	//
	// 1656950400
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeMeterImsEditUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterImsEditUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMeterImsEditUsageResponseBodyData) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeMeterImsEditUsageResponseBodyData) GetProfile() *string {
	return s.Profile
}

func (s *DescribeMeterImsEditUsageResponseBodyData) GetTime() *int64 {
	return s.Time
}

func (s *DescribeMeterImsEditUsageResponseBodyData) SetDuration(v int64) *DescribeMeterImsEditUsageResponseBodyData {
	s.Duration = &v
	return s
}

func (s *DescribeMeterImsEditUsageResponseBodyData) SetProfile(v string) *DescribeMeterImsEditUsageResponseBodyData {
	s.Profile = &v
	return s
}

func (s *DescribeMeterImsEditUsageResponseBodyData) SetTime(v int64) *DescribeMeterImsEditUsageResponseBodyData {
	s.Time = &v
	return s
}

func (s *DescribeMeterImsEditUsageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
