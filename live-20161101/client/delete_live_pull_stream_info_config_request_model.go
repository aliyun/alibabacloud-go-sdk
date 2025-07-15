// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePullStreamInfoConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteLivePullStreamInfoConfigRequest
	GetAppName() *string
	SetDomainName(v string) *DeleteLivePullStreamInfoConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLivePullStreamInfoConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLivePullStreamInfoConfigRequest
	GetRegionId() *string
	SetStreamName(v string) *DeleteLivePullStreamInfoConfigRequest
	GetStreamName() *string
}

type DeleteLivePullStreamInfoConfigRequest struct {
	// The name of the application to which the live stream belongs. You can call [DescribeLivePullStreamConfig](https://help.aliyun.com/document_detail/2847818.htmll) operation to query the application name.
	//
	// This parameter is required.
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
	// The name of the live stream. You can call [DescribeLivePullStreamConfig](https://help.aliyun.com/document_detail/2847818.htmll) operation to query the stream name.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DeleteLivePullStreamInfoConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLivePullStreamInfoConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteLivePullStreamInfoConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLivePullStreamInfoConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLivePullStreamInfoConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLivePullStreamInfoConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DeleteLivePullStreamInfoConfigRequest) SetAppName(v string) *DeleteLivePullStreamInfoConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLivePullStreamInfoConfigRequest) SetDomainName(v string) *DeleteLivePullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLivePullStreamInfoConfigRequest) SetOwnerId(v int64) *DeleteLivePullStreamInfoConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLivePullStreamInfoConfigRequest) SetRegionId(v string) *DeleteLivePullStreamInfoConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLivePullStreamInfoConfigRequest) SetStreamName(v string) *DeleteLivePullStreamInfoConfigRequest {
	s.StreamName = &v
	return s
}

func (s *DeleteLivePullStreamInfoConfigRequest) Validate() error {
	return dara.Validate(s)
}
