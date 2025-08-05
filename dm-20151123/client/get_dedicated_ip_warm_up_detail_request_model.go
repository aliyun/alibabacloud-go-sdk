// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDedicatedIpWarmUpDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedIp(v string) *GetDedicatedIpWarmUpDetailRequest
	GetDedicatedIp() *string
	SetEndDayMark(v int64) *GetDedicatedIpWarmUpDetailRequest
	GetEndDayMark() *int64
	SetEsp(v string) *GetDedicatedIpWarmUpDetailRequest
	GetEsp() *string
	SetStartDayMark(v int64) *GetDedicatedIpWarmUpDetailRequest
	GetStartDayMark() *int64
}

type GetDedicatedIpWarmUpDetailRequest struct {
	DedicatedIp  *string `json:"DedicatedIp,omitempty" xml:"DedicatedIp,omitempty"`
	EndDayMark   *int64  `json:"EndDayMark,omitempty" xml:"EndDayMark,omitempty"`
	Esp          *string `json:"Esp,omitempty" xml:"Esp,omitempty"`
	StartDayMark *int64  `json:"StartDayMark,omitempty" xml:"StartDayMark,omitempty"`
}

func (s GetDedicatedIpWarmUpDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDedicatedIpWarmUpDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDedicatedIpWarmUpDetailRequest) GetDedicatedIp() *string {
	return s.DedicatedIp
}

func (s *GetDedicatedIpWarmUpDetailRequest) GetEndDayMark() *int64 {
	return s.EndDayMark
}

func (s *GetDedicatedIpWarmUpDetailRequest) GetEsp() *string {
	return s.Esp
}

func (s *GetDedicatedIpWarmUpDetailRequest) GetStartDayMark() *int64 {
	return s.StartDayMark
}

func (s *GetDedicatedIpWarmUpDetailRequest) SetDedicatedIp(v string) *GetDedicatedIpWarmUpDetailRequest {
	s.DedicatedIp = &v
	return s
}

func (s *GetDedicatedIpWarmUpDetailRequest) SetEndDayMark(v int64) *GetDedicatedIpWarmUpDetailRequest {
	s.EndDayMark = &v
	return s
}

func (s *GetDedicatedIpWarmUpDetailRequest) SetEsp(v string) *GetDedicatedIpWarmUpDetailRequest {
	s.Esp = &v
	return s
}

func (s *GetDedicatedIpWarmUpDetailRequest) SetStartDayMark(v int64) *GetDedicatedIpWarmUpDetailRequest {
	s.StartDayMark = &v
	return s
}

func (s *GetDedicatedIpWarmUpDetailRequest) Validate() error {
	return dara.Validate(s)
}
