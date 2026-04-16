// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrendPointDTO interface {
	dara.Model
	String() string
	GoString() string
	SetTimestamp(v int64) *TrendPointDTO
	GetTimestamp() *int64
	SetValues(v string) *TrendPointDTO
	GetValues() *string
}

type TrendPointDTO struct {
	// example:
	//
	// 1700000000
	Timestamp *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Values    *string `json:"values,omitempty" xml:"values,omitempty"`
}

func (s TrendPointDTO) String() string {
	return dara.Prettify(s)
}

func (s TrendPointDTO) GoString() string {
	return s.String()
}

func (s *TrendPointDTO) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *TrendPointDTO) GetValues() *string {
	return s.Values
}

func (s *TrendPointDTO) SetTimestamp(v int64) *TrendPointDTO {
	s.Timestamp = &v
	return s
}

func (s *TrendPointDTO) SetValues(v string) *TrendPointDTO {
	s.Values = &v
	return s
}

func (s *TrendPointDTO) Validate() error {
	return dara.Validate(s)
}
