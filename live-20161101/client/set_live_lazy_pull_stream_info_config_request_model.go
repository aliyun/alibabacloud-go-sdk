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
	// The name of the application.
	//
	// > To trigger origin fetch for all applications, set this parameter to **ali_all_app**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ali_all_app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The streaming domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The source application name.
	//
	// > Leave this parameter empty to use the application name from the playback URL of the source stream.
	//
	// example:
	//
	// livePullApp****
	PullAppName *string `json:"PullAppName,omitempty" xml:"PullAppName,omitempty"`
	// The origin server that hosts the live stream. To specify multiple origin servers, separate them with semicolons (;).
	//
	// This parameter is required.
	//
	// example:
	//
	// guide.aliyundoc.com
	PullDomainName *string `json:"PullDomainName,omitempty" xml:"PullDomainName,omitempty"`
	// The protocol to use for pulling the stream from the source. Valid values:
	//
	// - **rtmp**
	//
	// - **httpflv**
	//
	// - **hls**
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp
	PullProtocol *string `json:"PullProtocol,omitempty" xml:"PullProtocol,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to trigger stream pulling when a request for a transcoded stream is made. Default value: **no**. Valid values:
	//
	// - **yes**
	//
	// - **no**
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
