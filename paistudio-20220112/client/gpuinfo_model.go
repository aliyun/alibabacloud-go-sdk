// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGPUInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *GPUInfo
	GetCount() *int64
	SetType(v string) *GPUInfo
	GetType() *string
}

type GPUInfo struct {
	Count *int64  `json:"count,omitempty" xml:"count,omitempty"`
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GPUInfo) String() string {
	return dara.Prettify(s)
}

func (s GPUInfo) GoString() string {
	return s.String()
}

func (s *GPUInfo) GetCount() *int64 {
	return s.Count
}

func (s *GPUInfo) GetType() *string {
	return s.Type
}

func (s *GPUInfo) SetCount(v int64) *GPUInfo {
	s.Count = &v
	return s
}

func (s *GPUInfo) SetType(v string) *GPUInfo {
	s.Type = &v
	return s
}

func (s *GPUInfo) Validate() error {
	return dara.Validate(s)
}
