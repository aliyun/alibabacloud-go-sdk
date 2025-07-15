// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveLazyPullStreamInfoConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteLiveLazyPullStreamInfoConfigRequest
	GetAppName() *string
	SetDomainName(v string) *DeleteLiveLazyPullStreamInfoConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLiveLazyPullStreamInfoConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveLazyPullStreamInfoConfigRequest
	GetRegionId() *string
}

type DeleteLiveLazyPullStreamInfoConfigRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// >  If you want to delete configurations of triggered stream pulling for all applications, set the value to **ali_all_app**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ali_all_app
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

func (s DeleteLiveLazyPullStreamInfoConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveLazyPullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveLazyPullStreamInfoConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteLiveLazyPullStreamInfoConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveLazyPullStreamInfoConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveLazyPullStreamInfoConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveLazyPullStreamInfoConfigRequest) SetAppName(v string) *DeleteLiveLazyPullStreamInfoConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLiveLazyPullStreamInfoConfigRequest) SetDomainName(v string) *DeleteLiveLazyPullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveLazyPullStreamInfoConfigRequest) SetOwnerId(v int64) *DeleteLiveLazyPullStreamInfoConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveLazyPullStreamInfoConfigRequest) SetRegionId(v string) *DeleteLiveLazyPullStreamInfoConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveLazyPullStreamInfoConfigRequest) Validate() error {
	return dara.Validate(s)
}
