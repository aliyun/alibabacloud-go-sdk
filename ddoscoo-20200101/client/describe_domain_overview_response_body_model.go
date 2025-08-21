// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxHttp(v int64) *DescribeDomainOverviewResponseBody
	GetMaxHttp() *int64
	SetMaxHttps(v int64) *DescribeDomainOverviewResponseBody
	GetMaxHttps() *int64
	SetRequestId(v string) *DescribeDomainOverviewResponseBody
	GetRequestId() *string
}

type DescribeDomainOverviewResponseBody struct {
	// The peak queries per second (QPS) during HTTP traffic scrubbing. Unit: QPS.
	//
	// example:
	//
	// 41652
	MaxHttp *int64 `json:"MaxHttp,omitempty" xml:"MaxHttp,omitempty"`
	// The peak QPS during HTTPS traffic scrubbing. Unit: QPS.
	//
	// example:
	//
	// 0
	MaxHttps *int64 `json:"MaxHttps,omitempty" xml:"MaxHttps,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainOverviewResponseBody) GetMaxHttp() *int64 {
	return s.MaxHttp
}

func (s *DescribeDomainOverviewResponseBody) GetMaxHttps() *int64 {
	return s.MaxHttps
}

func (s *DescribeDomainOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainOverviewResponseBody) SetMaxHttp(v int64) *DescribeDomainOverviewResponseBody {
	s.MaxHttp = &v
	return s
}

func (s *DescribeDomainOverviewResponseBody) SetMaxHttps(v int64) *DescribeDomainOverviewResponseBody {
	s.MaxHttps = &v
	return s
}

func (s *DescribeDomainOverviewResponseBody) SetRequestId(v string) *DescribeDomainOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainOverviewResponseBody) Validate() error {
	return dara.Validate(s)
}
