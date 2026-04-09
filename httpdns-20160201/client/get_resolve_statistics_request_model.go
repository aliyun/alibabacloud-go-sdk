// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResolveStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *GetResolveStatisticsRequest
	GetDomainName() *string
	SetGranularity(v string) *GetResolveStatisticsRequest
	GetGranularity() *string
	SetProtocolName(v string) *GetResolveStatisticsRequest
	GetProtocolName() *string
	SetTimeSpan(v int32) *GetResolveStatisticsRequest
	GetTimeSpan() *int32
}

type GetResolveStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// month
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// example:
	//
	// https
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	TimeSpan *int32 `json:"TimeSpan,omitempty" xml:"TimeSpan,omitempty"`
}

func (s GetResolveStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResolveStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetResolveStatisticsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *GetResolveStatisticsRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *GetResolveStatisticsRequest) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *GetResolveStatisticsRequest) GetTimeSpan() *int32 {
	return s.TimeSpan
}

func (s *GetResolveStatisticsRequest) SetDomainName(v string) *GetResolveStatisticsRequest {
	s.DomainName = &v
	return s
}

func (s *GetResolveStatisticsRequest) SetGranularity(v string) *GetResolveStatisticsRequest {
	s.Granularity = &v
	return s
}

func (s *GetResolveStatisticsRequest) SetProtocolName(v string) *GetResolveStatisticsRequest {
	s.ProtocolName = &v
	return s
}

func (s *GetResolveStatisticsRequest) SetTimeSpan(v int32) *GetResolveStatisticsRequest {
	s.TimeSpan = &v
	return s
}

func (s *GetResolveStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
