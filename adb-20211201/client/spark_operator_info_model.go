// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSparkOperatorInfo interface {
	dara.Model
	String() string
	GoString() string
	SetMetricValue(v int64) *SparkOperatorInfo
	GetMetricValue() *int64
	SetOperatorName(v []byte) *SparkOperatorInfo
	GetOperatorName() []byte
}

type SparkOperatorInfo struct {
	MetricValue  *int64 `json:"MetricValue,omitempty" xml:"MetricValue,omitempty"`
	OperatorName []byte `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
}

func (s SparkOperatorInfo) String() string {
	return dara.Prettify(s)
}

func (s SparkOperatorInfo) GoString() string {
	return s.String()
}

func (s *SparkOperatorInfo) GetMetricValue() *int64 {
	return s.MetricValue
}

func (s *SparkOperatorInfo) GetOperatorName() []byte {
	return s.OperatorName
}

func (s *SparkOperatorInfo) SetMetricValue(v int64) *SparkOperatorInfo {
	s.MetricValue = &v
	return s
}

func (s *SparkOperatorInfo) SetOperatorName(v []byte) *SparkOperatorInfo {
	s.OperatorName = v
	return s
}

func (s *SparkOperatorInfo) Validate() error {
	return dara.Validate(s)
}
