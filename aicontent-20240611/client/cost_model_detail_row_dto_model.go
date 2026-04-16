// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostModelDetailRowDTO interface {
	dara.Model
	String() string
	GoString() string
	SetTimestamp(v int64) *CostModelDetailRowDTO
	GetTimestamp() *int64
	SetValues(v string) *CostModelDetailRowDTO
	GetValues() *string
}

type CostModelDetailRowDTO struct {
	// example:
	//
	// 1700000000
	Timestamp *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Values    *string `json:"values,omitempty" xml:"values,omitempty"`
}

func (s CostModelDetailRowDTO) String() string {
	return dara.Prettify(s)
}

func (s CostModelDetailRowDTO) GoString() string {
	return s.String()
}

func (s *CostModelDetailRowDTO) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *CostModelDetailRowDTO) GetValues() *string {
	return s.Values
}

func (s *CostModelDetailRowDTO) SetTimestamp(v int64) *CostModelDetailRowDTO {
	s.Timestamp = &v
	return s
}

func (s *CostModelDetailRowDTO) SetValues(v string) *CostModelDetailRowDTO {
	s.Values = &v
	return s
}

func (s *CostModelDetailRowDTO) Validate() error {
	return dara.Validate(s)
}
