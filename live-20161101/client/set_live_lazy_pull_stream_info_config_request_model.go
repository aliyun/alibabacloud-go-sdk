// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveLazyPullStreamInfoConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *SetLiveLazyPullStreamInfoConfigRequest
	GetAppName() *string
	SetDomainName(v string) *SetLiveLazyPullStreamInfoConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *SetLiveLazyPullStreamInfoConfigRequest
	GetOwnerId() *int64
	SetPullAppName(v string) *SetLiveLazyPullStreamInfoConfigRequest
	GetPullAppName() *string
	SetPullDomainName(v string) *SetLiveLazyPullStreamInfoConfigRequest
	GetPullDomainName() *string
	SetPullProtocol(v string) *SetLiveLazyPullStreamInfoConfigRequest
	GetPullProtocol() *string
	SetRegionId(v string) *SetLiveLazyPullStreamInfoConfigRequest
	GetRegionId() *string
	SetTranscodeLazy(v string) *SetLiveLazyPullStreamInfoConfigRequest
	GetTranscodeLazy() *string
}

type SetLiveLazyPullStreamInfoConfigRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// >  If you want to configure triggered stream pulling for all applications, set the value to **ali_all_app**.
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
	// The name of the application for back-to-origin stream pulling.
	//
	// >  If you want to use the application specified in the streaming URL, leave this parameter empty.
	//
	// example:
	//
	// livePullApp****
	PullAppName *string `json:"PullAppName,omitempty" xml:"PullAppName,omitempty"`
	// The origin server address of the live stream. Separate multiple addresses with semicolons (;).
	//
	// This parameter is required.
	//
	// example:
	//
	// guide.aliyundoc.com
	PullDomainName *string `json:"PullDomainName,omitempty" xml:"PullDomainName,omitempty"`
	// The protocol for back-to-origin stream pulling. Valid values:
	//
	// 	- **rtmp**
	//
	// 	- **httpflv**
	//
	// 	- **hls**
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp
	PullProtocol *string `json:"PullProtocol,omitempty" xml:"PullProtocol,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to trigger stream pulling when the transcoded stream is played. The default value is **no**. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no**
	//
	// example:
	//
	// no
	TranscodeLazy *string `json:"TranscodeLazy,omitempty" xml:"TranscodeLazy,omitempty"`
}

func (s SetLiveLazyPullStreamInfoConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLiveLazyPullStreamInfoConfigRequest) GoString() string {
	return s.String()
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) GetPullAppName() *string {
	return s.PullAppName
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) GetPullDomainName() *string {
	return s.PullDomainName
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) GetPullProtocol() *string {
	return s.PullProtocol
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) GetTranscodeLazy() *string {
	return s.TranscodeLazy
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetAppName(v string) *SetLiveLazyPullStreamInfoConfigRequest {
	s.AppName = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetDomainName(v string) *SetLiveLazyPullStreamInfoConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetOwnerId(v int64) *SetLiveLazyPullStreamInfoConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetPullAppName(v string) *SetLiveLazyPullStreamInfoConfigRequest {
	s.PullAppName = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetPullDomainName(v string) *SetLiveLazyPullStreamInfoConfigRequest {
	s.PullDomainName = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetPullProtocol(v string) *SetLiveLazyPullStreamInfoConfigRequest {
	s.PullProtocol = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetRegionId(v string) *SetLiveLazyPullStreamInfoConfigRequest {
	s.RegionId = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) SetTranscodeLazy(v string) *SetLiveLazyPullStreamInfoConfigRequest {
	s.TranscodeLazy = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigRequest) Validate() error {
	return dara.Validate(s)
}
