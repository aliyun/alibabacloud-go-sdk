// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebStaticsInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCpuUsage(v int64) *WebStaticsInfo
	GetCpuUsage() *int64
	SetInternetTrafficOut(v int64) *WebStaticsInfo
	GetInternetTrafficOut() *int64
	SetInvocations(v int64) *WebStaticsInfo
	GetInvocations() *int64
	SetMemoryUsage(v int64) *WebStaticsInfo
	GetMemoryUsage() *int64
}

type WebStaticsInfo struct {
	CpuUsage           *int64 `json:"CpuUsage,omitempty" xml:"CpuUsage,omitempty"`
	InternetTrafficOut *int64 `json:"InternetTrafficOut,omitempty" xml:"InternetTrafficOut,omitempty"`
	Invocations        *int64 `json:"Invocations,omitempty" xml:"Invocations,omitempty"`
	MemoryUsage        *int64 `json:"MemoryUsage,omitempty" xml:"MemoryUsage,omitempty"`
}

func (s WebStaticsInfo) String() string {
	return dara.Prettify(s)
}

func (s WebStaticsInfo) GoString() string {
	return s.String()
}

func (s *WebStaticsInfo) GetCpuUsage() *int64 {
	return s.CpuUsage
}

func (s *WebStaticsInfo) GetInternetTrafficOut() *int64 {
	return s.InternetTrafficOut
}

func (s *WebStaticsInfo) GetInvocations() *int64 {
	return s.Invocations
}

func (s *WebStaticsInfo) GetMemoryUsage() *int64 {
	return s.MemoryUsage
}

func (s *WebStaticsInfo) SetCpuUsage(v int64) *WebStaticsInfo {
	s.CpuUsage = &v
	return s
}

func (s *WebStaticsInfo) SetInternetTrafficOut(v int64) *WebStaticsInfo {
	s.InternetTrafficOut = &v
	return s
}

func (s *WebStaticsInfo) SetInvocations(v int64) *WebStaticsInfo {
	s.Invocations = &v
	return s
}

func (s *WebStaticsInfo) SetMemoryUsage(v int64) *WebStaticsInfo {
	s.MemoryUsage = &v
	return s
}

func (s *WebStaticsInfo) Validate() error {
	return dara.Validate(s)
}
