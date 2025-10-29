// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveLazyPullStreamConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveLazyPullConfigList(v *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigList) *DescribeLiveLazyPullStreamConfigResponseBody
	GetLiveLazyPullConfigList() *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigList
	SetRequestId(v string) *DescribeLiveLazyPullStreamConfigResponseBody
	GetRequestId() *string
}

type DescribeLiveLazyPullStreamConfigResponseBody struct {
	// The configurations of triggered stream pulling.
	LiveLazyPullConfigList *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigList `json:"LiveLazyPullConfigList,omitempty" xml:"LiveLazyPullConfigList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// A3136B58-5876-4168-83CA-B562781981A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveLazyPullStreamConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveLazyPullStreamConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveLazyPullStreamConfigResponseBody) GetLiveLazyPullConfigList() *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigList {
	return s.LiveLazyPullConfigList
}

func (s *DescribeLiveLazyPullStreamConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveLazyPullStreamConfigResponseBody) SetLiveLazyPullConfigList(v *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigList) *DescribeLiveLazyPullStreamConfigResponseBody {
	s.LiveLazyPullConfigList = v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseBody) SetRequestId(v string) *DescribeLiveLazyPullStreamConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseBody) Validate() error {
	if s.LiveLazyPullConfigList != nil {
		if err := s.LiveLazyPullConfigList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigList struct {
	LiveLazyPullConfig []*DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig `json:"LiveLazyPullConfig,omitempty" xml:"LiveLazyPullConfig,omitempty" type:"Repeated"`
}

func (s DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigList) GoString() string {
	return s.String()
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigList) GetLiveLazyPullConfig() []*DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig {
	return s.LiveLazyPullConfig
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigList) SetLiveLazyPullConfig(v []*DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigList {
	s.LiveLazyPullConfig = v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigList) Validate() error {
	if s.LiveLazyPullConfig != nil {
		for _, item := range s.LiveLazyPullConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig struct {
	// The name of the application.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the application for back-to-origin stream pulling. If the application specified in the streaming URL is used, this parameter is left empty.
	//
	// example:
	//
	// livePullApp****
	PullAppName *string `json:"PullAppName,omitempty" xml:"PullAppName,omitempty"`
	// The parameters of back-to-origin stream pulling.
	//
	// example:
	//
	// k=v
	PullArgs *string `json:"PullArgs,omitempty" xml:"PullArgs,omitempty"`
	// The domain name for back-to-origin stream pulling.
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
	// example:
	//
	// RTMP
	PullProtocol *string `json:"PullProtocol,omitempty" xml:"PullProtocol,omitempty"`
	// Indicates whether stream pulling is triggered when the transcoded stream is played. Default value: **no**. Valid values:
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

func (s DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) GetPullAppName() *string {
	return s.PullAppName
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) GetPullArgs() *string {
	return s.PullArgs
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) GetPullDomainName() *string {
	return s.PullDomainName
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) GetPullProtocol() *string {
	return s.PullProtocol
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) GetTranscodeLazy() *string {
	return s.TranscodeLazy
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) SetAppName(v string) *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig {
	s.AppName = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) SetDomainName(v string) *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) SetPullAppName(v string) *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig {
	s.PullAppName = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) SetPullArgs(v string) *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig {
	s.PullArgs = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) SetPullDomainName(v string) *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig {
	s.PullDomainName = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) SetPullProtocol(v string) *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig {
	s.PullProtocol = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) SetTranscodeLazy(v string) *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig {
	s.TranscodeLazy = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponseBodyLiveLazyPullConfigListLiveLazyPullConfig) Validate() error {
	return dara.Validate(s)
}
