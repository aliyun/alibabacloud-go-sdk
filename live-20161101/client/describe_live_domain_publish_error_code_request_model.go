// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainPublishErrorCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveDomainPublishErrorCodeRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveDomainPublishErrorCodeRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainPublishErrorCodeRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveDomainPublishErrorCodeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainPublishErrorCodeRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveDomainPublishErrorCodeRequest
	GetStartTime() *string
}

type DescribeLiveDomainPublishErrorCodeRequest struct {
	// The application name. The data is aggregated based on the application. If you specify this parameter, the DomainName parameter is required.
	//
	// example:
	//
	// AppName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ingest domain. If you want to specify multiple ingest domains, separate them with commas (,).
	//
	// >  This parameter is required.
	//
	// example:
	//
	// example.com,example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. Specify the time in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  If you do not configure StartTime, the data within the previous hour is queried.
	//
	// example:
	//
	// 2016-06-29T09:10:00Z
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. Specify the time in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  If you do not configure StartTime, the data within the previous hour is queried.
	//
	// example:
	//
	// 2016-06-29T09:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainPublishErrorCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPublishErrorCodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPublishErrorCodeRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveDomainPublishErrorCodeRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainPublishErrorCodeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainPublishErrorCodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainPublishErrorCodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainPublishErrorCodeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainPublishErrorCodeRequest) SetAppName(v string) *DescribeLiveDomainPublishErrorCodeRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeRequest) SetDomainName(v string) *DescribeLiveDomainPublishErrorCodeRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeRequest) SetEndTime(v string) *DescribeLiveDomainPublishErrorCodeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeRequest) SetOwnerId(v int64) *DescribeLiveDomainPublishErrorCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeRequest) SetRegionId(v string) *DescribeLiveDomainPublishErrorCodeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeRequest) SetStartTime(v string) *DescribeLiveDomainPublishErrorCodeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainPublishErrorCodeRequest) Validate() error {
	return dara.Validate(s)
}
