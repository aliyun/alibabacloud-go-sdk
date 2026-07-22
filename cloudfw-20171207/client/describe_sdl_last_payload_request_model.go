// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlLastPayloadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDstIp(v string) *DescribeSdlLastPayloadRequest
	GetDstIp() *string
	SetEndTime(v int64) *DescribeSdlLastPayloadRequest
	GetEndTime() *int64
	SetLang(v string) *DescribeSdlLastPayloadRequest
	GetLang() *string
	SetSensitiveCategory(v string) *DescribeSdlLastPayloadRequest
	GetSensitiveCategory() *string
	SetSrcIp(v string) *DescribeSdlLastPayloadRequest
	GetSrcIp() *string
	SetStartTime(v int64) *DescribeSdlLastPayloadRequest
	GetStartTime() *int64
}

type DescribeSdlLastPayloadRequest struct {
	// The destination IP address. This is an optional parameter used to filter by destination IP address.
	//
	// example:
	//
	// 47.100.102.XXX
	DstIp *string `json:"DstIp,omitempty" xml:"DstIp,omitempty"`
	// The end time of the query (UNIX timestamp in seconds). This parameter is required. If this parameter is not specified, the API returns an error.
	//
	// example:
	//
	// 1534408267
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of sensitive credential. This parameter is required. If this parameter is not specified, the API returns an error.
	//
	// example:
	//
	// id_card
	SensitiveCategory *string `json:"SensitiveCategory,omitempty" xml:"SensitiveCategory,omitempty"`
	// The source IP address. This is an optional parameter used to filter by source IP address.
	//
	// example:
	//
	// 121.40.84.XXX
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// The start time of the query (UNIX timestamp in seconds). This parameter is required. If this parameter is not specified, the API returns an error.
	//
	// example:
	//
	// 1656837360
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSdlLastPayloadRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlLastPayloadRequest) GoString() string {
	return s.String()
}

func (s *DescribeSdlLastPayloadRequest) GetDstIp() *string {
	return s.DstIp
}

func (s *DescribeSdlLastPayloadRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSdlLastPayloadRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSdlLastPayloadRequest) GetSensitiveCategory() *string {
	return s.SensitiveCategory
}

func (s *DescribeSdlLastPayloadRequest) GetSrcIp() *string {
	return s.SrcIp
}

func (s *DescribeSdlLastPayloadRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSdlLastPayloadRequest) SetDstIp(v string) *DescribeSdlLastPayloadRequest {
	s.DstIp = &v
	return s
}

func (s *DescribeSdlLastPayloadRequest) SetEndTime(v int64) *DescribeSdlLastPayloadRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSdlLastPayloadRequest) SetLang(v string) *DescribeSdlLastPayloadRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSdlLastPayloadRequest) SetSensitiveCategory(v string) *DescribeSdlLastPayloadRequest {
	s.SensitiveCategory = &v
	return s
}

func (s *DescribeSdlLastPayloadRequest) SetSrcIp(v string) *DescribeSdlLastPayloadRequest {
	s.SrcIp = &v
	return s
}

func (s *DescribeSdlLastPayloadRequest) SetStartTime(v int64) *DescribeSdlLastPayloadRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSdlLastPayloadRequest) Validate() error {
	return dara.Validate(s)
}
