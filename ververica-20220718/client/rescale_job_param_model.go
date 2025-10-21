// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRescaleJobParam interface {
	dara.Model
	String() string
	GoString() string
	SetJobParallelism(v int64) *RescaleJobParam
	GetJobParallelism() *int64
}

type RescaleJobParam struct {
	JobParallelism *int64 `json:"jobParallelism,omitempty" xml:"jobParallelism,omitempty"`
}

func (s RescaleJobParam) String() string {
	return dara.Prettify(s)
}

func (s RescaleJobParam) GoString() string {
	return s.String()
}

func (s *RescaleJobParam) GetJobParallelism() *int64 {
	return s.JobParallelism
}

func (s *RescaleJobParam) SetJobParallelism(v int64) *RescaleJobParam {
	s.JobParallelism = &v
	return s
}

func (s *RescaleJobParam) Validate() error {
	return dara.Validate(s)
}
