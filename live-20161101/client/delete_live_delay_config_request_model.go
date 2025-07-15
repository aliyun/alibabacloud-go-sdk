// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveDelayConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *DeleteLiveDelayConfigRequest
	GetApp() *string
	SetDomain(v string) *DeleteLiveDelayConfigRequest
	GetDomain() *string
	SetOwnerId(v int64) *DeleteLiveDelayConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveDelayConfigRequest
	GetRegionId() *string
	SetStream(v string) *DeleteLiveDelayConfigRequest
	GetStream() *string
}

type DeleteLiveDelayConfigRequest struct {
	// The name of the application to which the live stream belongs. You can specify an asterisk (\\*) as the value to match all applications that belong to the domain name.
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
	// The name of the live stream. You can use the wildcard (\\*) to specify all streams of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s DeleteLiveDelayConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveDelayConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveDelayConfigRequest) GetApp() *string {
	return s.App
}

func (s *DeleteLiveDelayConfigRequest) GetDomain() *string {
	return s.Domain
}

func (s *DeleteLiveDelayConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveDelayConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveDelayConfigRequest) GetStream() *string {
	return s.Stream
}

func (s *DeleteLiveDelayConfigRequest) SetApp(v string) *DeleteLiveDelayConfigRequest {
	s.App = &v
	return s
}

func (s *DeleteLiveDelayConfigRequest) SetDomain(v string) *DeleteLiveDelayConfigRequest {
	s.Domain = &v
	return s
}

func (s *DeleteLiveDelayConfigRequest) SetOwnerId(v int64) *DeleteLiveDelayConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveDelayConfigRequest) SetRegionId(v string) *DeleteLiveDelayConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveDelayConfigRequest) SetStream(v string) *DeleteLiveDelayConfigRequest {
	s.Stream = &v
	return s
}

func (s *DeleteLiveDelayConfigRequest) Validate() error {
	return dara.Validate(s)
}
