// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuota interface {
	dara.Model
	String() string
	GoString() string
	SetCpuCores(v int32) *Quota
	GetCpuCores() *int32
	SetMemoryGB(v int32) *Quota
	GetMemoryGB() *int32
	SetTagValue(v string) *Quota
	GetTagValue() *string
}

type Quota struct {
	CpuCores *int32  `json:"cpuCores,omitempty" xml:"cpuCores,omitempty"`
	MemoryGB *int32  `json:"memoryGB,omitempty" xml:"memoryGB,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s Quota) String() string {
	return dara.Prettify(s)
}

func (s Quota) GoString() string {
	return s.String()
}

func (s *Quota) GetCpuCores() *int32 {
	return s.CpuCores
}

func (s *Quota) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *Quota) GetTagValue() *string {
	return s.TagValue
}

func (s *Quota) SetCpuCores(v int32) *Quota {
	s.CpuCores = &v
	return s
}

func (s *Quota) SetMemoryGB(v int32) *Quota {
	s.MemoryGB = &v
	return s
}

func (s *Quota) SetTagValue(v string) *Quota {
	s.TagValue = &v
	return s
}

func (s *Quota) Validate() error {
	return dara.Validate(s)
}
