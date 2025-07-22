// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcUserCntDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRtcUserCntDataResponseBody
	GetRequestId() *string
	SetUserCntDataPerInterval(v *DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval) *DescribeRtcUserCntDataResponseBody
	GetUserCntDataPerInterval() *DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval
}

type DescribeRtcUserCntDataResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId              *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserCntDataPerInterval *DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval `json:"UserCntDataPerInterval,omitempty" xml:"UserCntDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeRtcUserCntDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcUserCntDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserCntDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRtcUserCntDataResponseBody) GetUserCntDataPerInterval() *DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval {
	return s.UserCntDataPerInterval
}

func (s *DescribeRtcUserCntDataResponseBody) SetRequestId(v string) *DescribeRtcUserCntDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcUserCntDataResponseBody) SetUserCntDataPerInterval(v *DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval) *DescribeRtcUserCntDataResponseBody {
	s.UserCntDataPerInterval = v
	return s
}

func (s *DescribeRtcUserCntDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval struct {
	UserCntModule []*DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule `json:"UserCntModule,omitempty" xml:"UserCntModule,omitempty" type:"Repeated"`
}

func (s DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval) GetUserCntModule() []*DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule {
	return s.UserCntModule
}

func (s *DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval) SetUserCntModule(v []*DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule) *DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval {
	s.UserCntModule = v
	return s
}

func (s *DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule struct {
	// example:
	//
	// 10
	ActiveUserCnt *int64 `json:"ActiveUserCnt,omitempty" xml:"ActiveUserCnt,omitempty"`
	// example:
	//
	// 2018-01-29T00:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule) GetActiveUserCnt() *int64 {
	return s.ActiveUserCnt
}

func (s *DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule) SetActiveUserCnt(v int64) *DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule {
	s.ActiveUserCnt = &v
	return s
}

func (s *DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule) SetTimeStamp(v string) *DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule) Validate() error {
	return dara.Validate(s)
}
