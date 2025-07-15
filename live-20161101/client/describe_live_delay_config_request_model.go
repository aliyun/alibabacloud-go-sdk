// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDelayConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *DescribeLiveDelayConfigRequest
	GetApp() *string
	SetDomain(v string) *DescribeLiveDelayConfigRequest
	GetDomain() *string
	SetOwnerId(v int64) *DescribeLiveDelayConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDelayConfigRequest
	GetRegionId() *string
	SetStream(v string) *DescribeLiveDelayConfigRequest
	GetStream() *string
}

type DescribeLiveDelayConfigRequest struct {
	// The name of the application to which the live stream belongs. You can specify an asterisk (\\*) as the value to match all applications under the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream. You can specify an asterisk (\\*) as the value to match all streams in the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s DescribeLiveDelayConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDelayConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDelayConfigRequest) GetApp() *string {
	return s.App
}

func (s *DescribeLiveDelayConfigRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLiveDelayConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDelayConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDelayConfigRequest) GetStream() *string {
	return s.Stream
}

func (s *DescribeLiveDelayConfigRequest) SetApp(v string) *DescribeLiveDelayConfigRequest {
	s.App = &v
	return s
}

func (s *DescribeLiveDelayConfigRequest) SetDomain(v string) *DescribeLiveDelayConfigRequest {
	s.Domain = &v
	return s
}

func (s *DescribeLiveDelayConfigRequest) SetOwnerId(v int64) *DescribeLiveDelayConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDelayConfigRequest) SetRegionId(v string) *DescribeLiveDelayConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDelayConfigRequest) SetStream(v string) *DescribeLiveDelayConfigRequest {
	s.Stream = &v
	return s
}

func (s *DescribeLiveDelayConfigRequest) Validate() error {
	return dara.Validate(s)
}
