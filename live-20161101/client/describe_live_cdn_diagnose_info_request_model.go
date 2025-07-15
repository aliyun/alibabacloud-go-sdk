// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveCdnDiagnoseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityToken(v string) *DescribeLiveCdnDiagnoseInfoRequest
	GetSecurityToken() *string
	SetApp(v string) *DescribeLiveCdnDiagnoseInfoRequest
	GetApp() *string
	SetDomain(v string) *DescribeLiveCdnDiagnoseInfoRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeLiveCdnDiagnoseInfoRequest
	GetEndTime() *int64
	SetIntervalType(v string) *DescribeLiveCdnDiagnoseInfoRequest
	GetIntervalType() *string
	SetPhase(v int32) *DescribeLiveCdnDiagnoseInfoRequest
	GetPhase() *int32
	SetRequestType(v string) *DescribeLiveCdnDiagnoseInfoRequest
	GetRequestType() *string
	SetStartTime(v int64) *DescribeLiveCdnDiagnoseInfoRequest
	GetStartTime() *int64
	SetStreamName(v string) *DescribeLiveCdnDiagnoseInfoRequest
	GetStreamName() *string
	SetStreamSuffix(v string) *DescribeLiveCdnDiagnoseInfoRequest
	GetStreamSuffix() *string
}

type DescribeLiveCdnDiagnoseInfoRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	App *string `json:"app,omitempty" xml:"app,omitempty"`
	// This parameter is required.
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// This parameter is required.
	EndTime      *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	IntervalType *string `json:"intervalType,omitempty" xml:"intervalType,omitempty"`
	// This parameter is required.
	Phase       *int32  `json:"phase,omitempty" xml:"phase,omitempty"`
	RequestType *string `json:"requestType,omitempty" xml:"requestType,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	StreamName *string `json:"streamName,omitempty" xml:"streamName,omitempty"`
	// This parameter is required.
	StreamSuffix *string `json:"streamSuffix,omitempty" xml:"streamSuffix,omitempty"`
}

func (s DescribeLiveCdnDiagnoseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCdnDiagnoseInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) GetApp() *string {
	return s.App
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) GetIntervalType() *string {
	return s.IntervalType
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) GetPhase() *int32 {
	return s.Phase
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) GetRequestType() *string {
	return s.RequestType
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) GetStreamSuffix() *string {
	return s.StreamSuffix
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) SetSecurityToken(v string) *DescribeLiveCdnDiagnoseInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) SetApp(v string) *DescribeLiveCdnDiagnoseInfoRequest {
	s.App = &v
	return s
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) SetDomain(v string) *DescribeLiveCdnDiagnoseInfoRequest {
	s.Domain = &v
	return s
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) SetEndTime(v int64) *DescribeLiveCdnDiagnoseInfoRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) SetIntervalType(v string) *DescribeLiveCdnDiagnoseInfoRequest {
	s.IntervalType = &v
	return s
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) SetPhase(v int32) *DescribeLiveCdnDiagnoseInfoRequest {
	s.Phase = &v
	return s
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) SetRequestType(v string) *DescribeLiveCdnDiagnoseInfoRequest {
	s.RequestType = &v
	return s
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) SetStartTime(v int64) *DescribeLiveCdnDiagnoseInfoRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) SetStreamName(v string) *DescribeLiveCdnDiagnoseInfoRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) SetStreamSuffix(v string) *DescribeLiveCdnDiagnoseInfoRequest {
	s.StreamSuffix = &v
	return s
}

func (s *DescribeLiveCdnDiagnoseInfoRequest) Validate() error {
	return dara.Validate(s)
}
