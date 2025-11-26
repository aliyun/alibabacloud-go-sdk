// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCollationTimeZone interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentTimeOffset(v string) *CollationTimeZone
	GetCurrentTimeOffset() *string
	SetTimeZone(v string) *CollationTimeZone
	GetTimeZone() *string
}

type CollationTimeZone struct {
	// example:
	//
	// UTC+08:00
	CurrentTimeOffset *string `json:"CurrentTimeOffset,omitempty" xml:"CurrentTimeOffset,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s CollationTimeZone) String() string {
	return dara.Prettify(s)
}

func (s CollationTimeZone) GoString() string {
	return s.String()
}

func (s *CollationTimeZone) GetCurrentTimeOffset() *string {
	return s.CurrentTimeOffset
}

func (s *CollationTimeZone) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CollationTimeZone) SetCurrentTimeOffset(v string) *CollationTimeZone {
	s.CurrentTimeOffset = &v
	return s
}

func (s *CollationTimeZone) SetTimeZone(v string) *CollationTimeZone {
	s.TimeZone = &v
	return s
}

func (s *CollationTimeZone) Validate() error {
	return dara.Validate(s)
}
