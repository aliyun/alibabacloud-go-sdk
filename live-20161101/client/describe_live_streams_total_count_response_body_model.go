// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsTotalCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveStreamsTotalCountResponseBody
	GetRequestId() *string
	SetStreamCountList(v *DescribeLiveStreamsTotalCountResponseBodyStreamCountList) *DescribeLiveStreamsTotalCountResponseBody
	GetStreamCountList() *DescribeLiveStreamsTotalCountResponseBodyStreamCountList
}

type DescribeLiveStreamsTotalCountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// String	FCFFE4A4-F34F-4EEF-B401-36A01689AFBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics about the live streams.
	StreamCountList *DescribeLiveStreamsTotalCountResponseBodyStreamCountList `json:"StreamCountList,omitempty" xml:"StreamCountList,omitempty" type:"Struct"`
}

func (s DescribeLiveStreamsTotalCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsTotalCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsTotalCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamsTotalCountResponseBody) GetStreamCountList() *DescribeLiveStreamsTotalCountResponseBodyStreamCountList {
	return s.StreamCountList
}

func (s *DescribeLiveStreamsTotalCountResponseBody) SetRequestId(v string) *DescribeLiveStreamsTotalCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamsTotalCountResponseBody) SetStreamCountList(v *DescribeLiveStreamsTotalCountResponseBodyStreamCountList) *DescribeLiveStreamsTotalCountResponseBody {
	s.StreamCountList = v
	return s
}

func (s *DescribeLiveStreamsTotalCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamsTotalCountResponseBodyStreamCountList struct {
	StreamCountInfos []*DescribeLiveStreamsTotalCountResponseBodyStreamCountListStreamCountInfos `json:"StreamCountInfos,omitempty" xml:"StreamCountInfos,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamsTotalCountResponseBodyStreamCountList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsTotalCountResponseBodyStreamCountList) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsTotalCountResponseBodyStreamCountList) GetStreamCountInfos() []*DescribeLiveStreamsTotalCountResponseBodyStreamCountListStreamCountInfos {
	return s.StreamCountInfos
}

func (s *DescribeLiveStreamsTotalCountResponseBodyStreamCountList) SetStreamCountInfos(v []*DescribeLiveStreamsTotalCountResponseBodyStreamCountListStreamCountInfos) *DescribeLiveStreamsTotalCountResponseBodyStreamCountList {
	s.StreamCountInfos = v
	return s
}

func (s *DescribeLiveStreamsTotalCountResponseBodyStreamCountList) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamsTotalCountResponseBodyStreamCountListStreamCountInfos struct {
	// The total number of live streams.
	//
	// example:
	//
	// 934
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The timestamp.
	//
	// example:
	//
	// 2023-07-24T16:00:00Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeLiveStreamsTotalCountResponseBodyStreamCountListStreamCountInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsTotalCountResponseBodyStreamCountListStreamCountInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsTotalCountResponseBodyStreamCountListStreamCountInfos) GetCount() *int32 {
	return s.Count
}

func (s *DescribeLiveStreamsTotalCountResponseBodyStreamCountListStreamCountInfos) GetTimestamp() *string {
	return s.Timestamp
}

func (s *DescribeLiveStreamsTotalCountResponseBodyStreamCountListStreamCountInfos) SetCount(v int32) *DescribeLiveStreamsTotalCountResponseBodyStreamCountListStreamCountInfos {
	s.Count = &v
	return s
}

func (s *DescribeLiveStreamsTotalCountResponseBodyStreamCountListStreamCountInfos) SetTimestamp(v string) *DescribeLiveStreamsTotalCountResponseBodyStreamCountListStreamCountInfos {
	s.Timestamp = &v
	return s
}

func (s *DescribeLiveStreamsTotalCountResponseBodyStreamCountListStreamCountInfos) Validate() error {
	return dara.Validate(s)
}
