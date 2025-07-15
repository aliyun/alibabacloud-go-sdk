// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsNotifyUrlConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveStreamsNotifyConfig(v *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) *DescribeLiveStreamsNotifyUrlConfigResponseBody
	GetLiveStreamsNotifyConfig() *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig
	SetRequestId(v string) *DescribeLiveStreamsNotifyUrlConfigResponseBody
	GetRequestId() *string
}

type DescribeLiveStreamsNotifyUrlConfigResponseBody struct {
	// The callback configuration.
	LiveStreamsNotifyConfig *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig `json:"LiveStreamsNotifyConfig,omitempty" xml:"LiveStreamsNotifyConfig,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 40A4F36D-A7CC-473A-88E7-154F92242566
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveStreamsNotifyUrlConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsNotifyUrlConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseBody) GetLiveStreamsNotifyConfig() *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig {
	return s.LiveStreamsNotifyConfig
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseBody) SetLiveStreamsNotifyConfig(v *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) *DescribeLiveStreamsNotifyUrlConfigResponseBody {
	s.LiveStreamsNotifyConfig = v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseBody) SetRequestId(v string) *DescribeLiveStreamsNotifyUrlConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig struct {
	// The ingest domain.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Exception event callback URL.
	//
	// example:
	//
	// https://4a7e5f08.r37.cpolar.top/live/Record/call-back/streamException
	ExceptionNotifyUrl *string `json:"ExceptionNotifyUrl,omitempty" xml:"ExceptionNotifyUrl,omitempty"`
	// The authentication key.
	//
	// example:
	//
	// 123456
	NotifyAuthKey *string `json:"NotifyAuthKey,omitempty" xml:"NotifyAuthKey,omitempty"`
	// Indicates whether callback authentication is enabled. Valid values:
	//
	// 	- yes
	//
	// 	- no
	//
	// example:
	//
	// yes
	NotifyReqAuth *string `json:"NotifyReqAuth,omitempty" xml:"NotifyReqAuth,omitempty"`
	// The callback URL.
	//
	// example:
	//
	// http://guide.aliyundoc.com/notify
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
}

func (s DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) GetExceptionNotifyUrl() *string {
	return s.ExceptionNotifyUrl
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) GetNotifyAuthKey() *string {
	return s.NotifyAuthKey
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) GetNotifyReqAuth() *string {
	return s.NotifyReqAuth
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) SetDomainName(v string) *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) SetExceptionNotifyUrl(v string) *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig {
	s.ExceptionNotifyUrl = &v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) SetNotifyAuthKey(v string) *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig {
	s.NotifyAuthKey = &v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) SetNotifyReqAuth(v string) *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig {
	s.NotifyReqAuth = &v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) SetNotifyUrl(v string) *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig {
	s.NotifyUrl = &v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) Validate() error {
	return dara.Validate(s)
}
