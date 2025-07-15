// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamAuthCheckingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveStreamAuthCheckingRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveStreamAuthCheckingRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveStreamAuthCheckingRequest
	GetRegionId() *string
	SetUrl(v string) *DescribeLiveStreamAuthCheckingRequest
	GetUrl() *string
}

type DescribeLiveStreamAuthCheckingRequest struct {
	// The ingest domain or streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The complete ingest URL or streaming URL. You can use the [URL generator](https://help.aliyun.com/document_detail/197400.html) to generate a URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example.com/live/test.flv?auth_key=1664248******
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeLiveStreamAuthCheckingRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamAuthCheckingRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamAuthCheckingRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamAuthCheckingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamAuthCheckingRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamAuthCheckingRequest) GetUrl() *string {
	return s.Url
}

func (s *DescribeLiveStreamAuthCheckingRequest) SetDomainName(v string) *DescribeLiveStreamAuthCheckingRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamAuthCheckingRequest) SetOwnerId(v int64) *DescribeLiveStreamAuthCheckingRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamAuthCheckingRequest) SetRegionId(v string) *DescribeLiveStreamAuthCheckingRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamAuthCheckingRequest) SetUrl(v string) *DescribeLiveStreamAuthCheckingRequest {
	s.Url = &v
	return s
}

func (s *DescribeLiveStreamAuthCheckingRequest) Validate() error {
	return dara.Validate(s)
}
