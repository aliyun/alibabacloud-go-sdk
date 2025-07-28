// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZoneInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeZoneInfoRequest
	GetLang() *string
	SetZoneId(v string) *DescribeZoneInfoRequest
	GetZoneId() *string
}

type DescribeZoneInfoRequest struct {
	// The language of the response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English.
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The zone ID. This ID uniquely identifies the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// df2d03865266bd9842306db586d3****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZoneInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeZoneInfoRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeZoneInfoRequest) SetLang(v string) *DescribeZoneInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeZoneInfoRequest) SetZoneId(v string) *DescribeZoneInfoRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeZoneInfoRequest) Validate() error {
	return dara.Validate(s)
}
