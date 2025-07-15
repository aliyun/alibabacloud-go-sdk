// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePrivateLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerationType(v string) *DeleteLivePrivateLineRequest
	GetAccelerationType() *string
	SetAppName(v string) *DeleteLivePrivateLineRequest
	GetAppName() *string
	SetDomainName(v string) *DeleteLivePrivateLineRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLivePrivateLineRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLivePrivateLineRequest
	GetRegionId() *string
	SetStreamName(v string) *DeleteLivePrivateLineRequest
	GetStreamName() *string
}

type DeleteLivePrivateLineRequest struct {
	// The acceleration type. Valid values:
	//
	// 	- play: streaming acceleration
	//
	// 	- publish: stream ingest acceleration
	//
	// This parameter is required.
	//
	// example:
	//
	// play
	AccelerationType *string `json:"AccelerationType,omitempty" xml:"AccelerationType,omitempty"`
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// live
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// testStream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DeleteLivePrivateLineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePrivateLineRequest) GoString() string {
	return s.String()
}

func (s *DeleteLivePrivateLineRequest) GetAccelerationType() *string {
	return s.AccelerationType
}

func (s *DeleteLivePrivateLineRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteLivePrivateLineRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLivePrivateLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLivePrivateLineRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLivePrivateLineRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DeleteLivePrivateLineRequest) SetAccelerationType(v string) *DeleteLivePrivateLineRequest {
	s.AccelerationType = &v
	return s
}

func (s *DeleteLivePrivateLineRequest) SetAppName(v string) *DeleteLivePrivateLineRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLivePrivateLineRequest) SetDomainName(v string) *DeleteLivePrivateLineRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLivePrivateLineRequest) SetOwnerId(v int64) *DeleteLivePrivateLineRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLivePrivateLineRequest) SetRegionId(v string) *DeleteLivePrivateLineRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLivePrivateLineRequest) SetStreamName(v string) *DeleteLivePrivateLineRequest {
	s.StreamName = &v
	return s
}

func (s *DeleteLivePrivateLineRequest) Validate() error {
	return dara.Validate(s)
}
