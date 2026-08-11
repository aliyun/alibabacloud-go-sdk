// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataTotalStatValue interface {
	dara.Model
	String() string
	GoString() string
	SetTotal(v int64) *DataTotalStatValue
	GetTotal() *int64
	SetShare(v string) *DataTotalStatValue
	GetShare() *string
}

type DataTotalStatValue struct {
	// The total count.
	//
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The proportion.
	//
	// example:
	//
	// 12.50%
	Share *string `json:"Share,omitempty" xml:"Share,omitempty"`
}

func (s DataTotalStatValue) String() string {
	return dara.Prettify(s)
}

func (s DataTotalStatValue) GoString() string {
	return s.String()
}

func (s *DataTotalStatValue) GetTotal() *int64 {
	return s.Total
}

func (s *DataTotalStatValue) GetShare() *string {
	return s.Share
}

func (s *DataTotalStatValue) SetTotal(v int64) *DataTotalStatValue {
	s.Total = &v
	return s
}

func (s *DataTotalStatValue) SetShare(v string) *DataTotalStatValue {
	s.Share = &v
	return s
}

func (s *DataTotalStatValue) Validate() error {
	return dara.Validate(s)
}
