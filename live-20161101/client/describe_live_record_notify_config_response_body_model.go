// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRecordNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveRecordNotifyConfig(v *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) *DescribeLiveRecordNotifyConfigResponseBody
	GetLiveRecordNotifyConfig() *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig
	SetRequestId(v string) *DescribeLiveRecordNotifyConfigResponseBody
	GetRequestId() *string
}

type DescribeLiveRecordNotifyConfigResponseBody struct {
	// The configuration of callbacks for live stream recording.
	LiveRecordNotifyConfig *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig `json:"LiveRecordNotifyConfig,omitempty" xml:"LiveRecordNotifyConfig,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5056369B-D337-499E-B8B7-B761BD37B08A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveRecordNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordNotifyConfigResponseBody) GetLiveRecordNotifyConfig() *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig {
	return s.LiveRecordNotifyConfig
}

func (s *DescribeLiveRecordNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveRecordNotifyConfigResponseBody) SetLiveRecordNotifyConfig(v *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) *DescribeLiveRecordNotifyConfigResponseBody {
	s.LiveRecordNotifyConfig = v
	return s
}

func (s *DescribeLiveRecordNotifyConfigResponseBody) SetRequestId(v string) *DescribeLiveRecordNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveRecordNotifyConfigResponseBody) Validate() error {
	if s.LiveRecordNotifyConfig != nil {
		if err := s.LiveRecordNotifyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig struct {
	// The main streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Indicates whether recording status callbacks are enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	NeedStatusNotify *bool   `json:"NeedStatusNotify,omitempty" xml:"NeedStatusNotify,omitempty"`
	NotifyAuthKey    *string `json:"NotifyAuthKey,omitempty" xml:"NotifyAuthKey,omitempty"`
	NotifyReqAuth    *bool   `json:"NotifyReqAuth,omitempty" xml:"NotifyReqAuth,omitempty"`
	// The recording callback URL.
	//
	// example:
	//
	// http://learn.aliyundoc.com/examplecallback.action
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The callback URL for on-demand recording.
	//
	// example:
	//
	// http://guide.aliyundoc.com/ondemandcallback.action
	OnDemandUrl *string `json:"OnDemandUrl,omitempty" xml:"OnDemandUrl,omitempty"`
}

func (s DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) GetNeedStatusNotify() *bool {
	return s.NeedStatusNotify
}

func (s *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) GetNotifyAuthKey() *string {
	return s.NotifyAuthKey
}

func (s *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) GetNotifyReqAuth() *bool {
	return s.NotifyReqAuth
}

func (s *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) GetOnDemandUrl() *string {
	return s.OnDemandUrl
}

func (s *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) SetDomainName(v string) *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) SetNeedStatusNotify(v bool) *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig {
	s.NeedStatusNotify = &v
	return s
}

func (s *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) SetNotifyAuthKey(v string) *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig {
	s.NotifyAuthKey = &v
	return s
}

func (s *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) SetNotifyReqAuth(v bool) *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig {
	s.NotifyReqAuth = &v
	return s
}

func (s *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) SetNotifyUrl(v string) *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig {
	s.NotifyUrl = &v
	return s
}

func (s *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) SetOnDemandUrl(v string) *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig {
	s.OnDemandUrl = &v
	return s
}

func (s *DescribeLiveRecordNotifyConfigResponseBodyLiveRecordNotifyConfig) Validate() error {
	return dara.Validate(s)
}
