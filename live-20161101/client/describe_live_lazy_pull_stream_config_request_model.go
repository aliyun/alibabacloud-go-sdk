// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveLazyPullStreamConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveLazyPullStreamConfigRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveLazyPullStreamConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveLazyPullStreamConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveLazyPullStreamConfigRequest
	GetRegionId() *string
}

type DescribeLiveLazyPullStreamConfigRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveLazyPullStreamConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveLazyPullStreamConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveLazyPullStreamConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveLazyPullStreamConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveLazyPullStreamConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveLazyPullStreamConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveLazyPullStreamConfigRequest) SetAppName(v string) *DescribeLiveLazyPullStreamConfigRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigRequest) SetDomainName(v string) *DescribeLiveLazyPullStreamConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigRequest) SetOwnerId(v int64) *DescribeLiveLazyPullStreamConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigRequest) SetRegionId(v string) *DescribeLiveLazyPullStreamConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigRequest) Validate() error {
	return dara.Validate(s)
}
