// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamTranscodeInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamTranscodeInfoRequest
	GetAppName() *string
	SetDomainTranscodeName(v string) *DescribeLiveStreamTranscodeInfoRequest
	GetDomainTranscodeName() *string
	SetOwnerId(v int64) *DescribeLiveStreamTranscodeInfoRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveStreamTranscodeInfoRequest
	GetRegionId() *string
}

type DescribeLiveStreamTranscodeInfoRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// myapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainTranscodeName *string `json:"DomainTranscodeName,omitempty" xml:"DomainTranscodeName,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveStreamTranscodeInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamTranscodeInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeInfoRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamTranscodeInfoRequest) GetDomainTranscodeName() *string {
	return s.DomainTranscodeName
}

func (s *DescribeLiveStreamTranscodeInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamTranscodeInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamTranscodeInfoRequest) SetAppName(v string) *DescribeLiveStreamTranscodeInfoRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoRequest) SetDomainTranscodeName(v string) *DescribeLiveStreamTranscodeInfoRequest {
	s.DomainTranscodeName = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoRequest) SetOwnerId(v int64) *DescribeLiveStreamTranscodeInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoRequest) SetRegionId(v string) *DescribeLiveStreamTranscodeInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoRequest) Validate() error {
	return dara.Validate(s)
}
