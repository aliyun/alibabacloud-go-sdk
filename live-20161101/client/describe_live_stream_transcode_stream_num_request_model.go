// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamTranscodeStreamNumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveStreamTranscodeStreamNumRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveStreamTranscodeStreamNumRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveStreamTranscodeStreamNumRequest
	GetRegionId() *string
	SetSplitType(v string) *DescribeLiveStreamTranscodeStreamNumRequest
	GetSplitType() *string
}

type DescribeLiveStreamTranscodeStreamNumRequest struct {
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The grouping method.
	//
	// 	- Domain name (default)
	//
	// 	- Template
	//
	// Valid values:
	//
	// 	- domain
	//
	// 	- template
	//
	// example:
	//
	// domain
	SplitType *string `json:"SplitType,omitempty" xml:"SplitType,omitempty"`
}

func (s DescribeLiveStreamTranscodeStreamNumRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamTranscodeStreamNumRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeStreamNumRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamTranscodeStreamNumRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamTranscodeStreamNumRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamTranscodeStreamNumRequest) GetSplitType() *string {
	return s.SplitType
}

func (s *DescribeLiveStreamTranscodeStreamNumRequest) SetDomainName(v string) *DescribeLiveStreamTranscodeStreamNumRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumRequest) SetOwnerId(v int64) *DescribeLiveStreamTranscodeStreamNumRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumRequest) SetRegionId(v string) *DescribeLiveStreamTranscodeStreamNumRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumRequest) SetSplitType(v string) *DescribeLiveStreamTranscodeStreamNumRequest {
	s.SplitType = &v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumRequest) Validate() error {
	return dara.Validate(s)
}
