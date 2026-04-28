// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAggregationsGroup interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *AggregationsGroup
	GetCount() *int64
	SetValue(v []byte) *AggregationsGroup
	GetValue() []byte
}

type AggregationsGroup struct {
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	Value []byte `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AggregationsGroup) String() string {
	return dara.Prettify(s)
}

func (s AggregationsGroup) GoString() string {
	return s.String()
}

func (s *AggregationsGroup) GetCount() *int64 {
	return s.Count
}

func (s *AggregationsGroup) GetValue() []byte {
	return s.Value
}

func (s *AggregationsGroup) SetCount(v int64) *AggregationsGroup {
	s.Count = &v
	return s
}

func (s *AggregationsGroup) SetValue(v []byte) *AggregationsGroup {
	s.Value = v
	return s
}

func (s *AggregationsGroup) Validate() error {
	return dara.Validate(s)
}
