// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveStreamCountResponseBody
	GetRequestId() *string
	SetStreamCountInfos(v *DescribeLiveStreamCountResponseBodyStreamCountInfos) *DescribeLiveStreamCountResponseBody
	GetStreamCountInfos() *DescribeLiveStreamCountResponseBodyStreamCountInfos
}

type DescribeLiveStreamCountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FCFFE4A4-F34F-4EEF-B401-36A01689AFBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics of the live streams.
	StreamCountInfos *DescribeLiveStreamCountResponseBodyStreamCountInfos `json:"StreamCountInfos,omitempty" xml:"StreamCountInfos,omitempty" type:"Struct"`
}

func (s DescribeLiveStreamCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamCountResponseBody) GetStreamCountInfos() *DescribeLiveStreamCountResponseBodyStreamCountInfos {
	return s.StreamCountInfos
}

func (s *DescribeLiveStreamCountResponseBody) SetRequestId(v string) *DescribeLiveStreamCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamCountResponseBody) SetStreamCountInfos(v *DescribeLiveStreamCountResponseBodyStreamCountInfos) *DescribeLiveStreamCountResponseBody {
	s.StreamCountInfos = v
	return s
}

func (s *DescribeLiveStreamCountResponseBody) Validate() error {
	if s.StreamCountInfos != nil {
		if err := s.StreamCountInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveStreamCountResponseBodyStreamCountInfos struct {
	StreamCountInfo []*DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo `json:"StreamCountInfo,omitempty" xml:"StreamCountInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamCountResponseBodyStreamCountInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamCountResponseBodyStreamCountInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfos) GetStreamCountInfo() []*DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo {
	return s.StreamCountInfo
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfos) SetStreamCountInfo(v []*DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo) *DescribeLiveStreamCountResponseBodyStreamCountInfos {
	s.StreamCountInfo = v
	return s
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfos) Validate() error {
	if s.StreamCountInfo != nil {
		for _, item := range s.StreamCountInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo struct {
	// The number of online streams.
	//
	// example:
	//
	// 3
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The maximum allowed number of concurrently ingested streams. This parameter is available only to users in the whitelist.
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The information about the live streams.
	StreamCountDetails *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetails `json:"StreamCountDetails,omitempty" xml:"StreamCountDetails,omitempty" type:"Struct"`
	// The type of the live stream. Valid values:
	//
	// 	- **raw**: source streams
	//
	// 	- **trans**: transcoded streams
	//
	// example:
	//
	// raw
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo) GetCount() *int64 {
	return s.Count
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo) GetStreamCountDetails() *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetails {
	return s.StreamCountDetails
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo) GetType() *string {
	return s.Type
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo) SetCount(v int64) *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo {
	s.Count = &v
	return s
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo) SetLimit(v int64) *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo {
	s.Limit = &v
	return s
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo) SetStreamCountDetails(v *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetails) *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo {
	s.StreamCountDetails = v
	return s
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo) SetType(v string) *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo {
	s.Type = &v
	return s
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfo) Validate() error {
	if s.StreamCountDetails != nil {
		if err := s.StreamCountDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetails struct {
	StreamCountDetail []*DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail `json:"StreamCountDetail,omitempty" xml:"StreamCountDetail,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetails) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetails) GetStreamCountDetail() []*DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail {
	return s.StreamCountDetail
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetails) SetStreamCountDetail(v []*DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail) *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetails {
	s.StreamCountDetail = v
	return s
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetails) Validate() error {
	if s.StreamCountDetail != nil {
		for _, item := range s.StreamCountDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail struct {
	// The number of online streams.
	//
	// example:
	//
	// 2
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The video codec. Valid values:
	//
	// 	- **h264**
	//
	// 	- **h265**
	//
	// example:
	//
	// h264
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The video bitrate. This parameter is returned only for transcoded streams.
	//
	// example:
	//
	// 390
	VideoDataRate *int64 `json:"VideoDataRate,omitempty" xml:"VideoDataRate,omitempty"`
}

func (s DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail) GetCount() *int64 {
	return s.Count
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail) GetFormat() *string {
	return s.Format
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail) GetVideoDataRate() *int64 {
	return s.VideoDataRate
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail) SetCount(v int64) *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail {
	s.Count = &v
	return s
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail) SetFormat(v string) *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail {
	s.Format = &v
	return s
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail) SetVideoDataRate(v int64) *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail {
	s.VideoDataRate = &v
	return s
}

func (s *DescribeLiveStreamCountResponseBodyStreamCountInfosStreamCountInfoStreamCountDetailsStreamCountDetail) Validate() error {
	return dara.Validate(s)
}
