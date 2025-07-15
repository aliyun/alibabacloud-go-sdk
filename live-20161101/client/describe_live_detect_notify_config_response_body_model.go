// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDetectNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveDetectNotifyConfig(v *DescribeLiveDetectNotifyConfigResponseBodyLiveDetectNotifyConfig) *DescribeLiveDetectNotifyConfigResponseBody
	GetLiveDetectNotifyConfig() *DescribeLiveDetectNotifyConfigResponseBodyLiveDetectNotifyConfig
	SetRequestId(v string) *DescribeLiveDetectNotifyConfigResponseBody
	GetRequestId() *string
}

type DescribeLiveDetectNotifyConfigResponseBody struct {
	// The callback configuration.
	LiveDetectNotifyConfig *DescribeLiveDetectNotifyConfigResponseBodyLiveDetectNotifyConfig `json:"LiveDetectNotifyConfig,omitempty" xml:"LiveDetectNotifyConfig,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveDetectNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDetectNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDetectNotifyConfigResponseBody) GetLiveDetectNotifyConfig() *DescribeLiveDetectNotifyConfigResponseBodyLiveDetectNotifyConfig {
	return s.LiveDetectNotifyConfig
}

func (s *DescribeLiveDetectNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDetectNotifyConfigResponseBody) SetLiveDetectNotifyConfig(v *DescribeLiveDetectNotifyConfigResponseBodyLiveDetectNotifyConfig) *DescribeLiveDetectNotifyConfigResponseBody {
	s.LiveDetectNotifyConfig = v
	return s
}

func (s *DescribeLiveDetectNotifyConfigResponseBody) SetRequestId(v string) *DescribeLiveDetectNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDetectNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDetectNotifyConfigResponseBodyLiveDetectNotifyConfig struct {
	// The main streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The callback URL.
	//
	// example:
	//
	// http://aliyundoc.com
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
}

func (s DescribeLiveDetectNotifyConfigResponseBodyLiveDetectNotifyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDetectNotifyConfigResponseBodyLiveDetectNotifyConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveDetectNotifyConfigResponseBodyLiveDetectNotifyConfig) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDetectNotifyConfigResponseBodyLiveDetectNotifyConfig) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *DescribeLiveDetectNotifyConfigResponseBodyLiveDetectNotifyConfig) SetDomainName(v string) *DescribeLiveDetectNotifyConfigResponseBodyLiveDetectNotifyConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDetectNotifyConfigResponseBodyLiveDetectNotifyConfig) SetNotifyUrl(v string) *DescribeLiveDetectNotifyConfigResponseBodyLiveDetectNotifyConfig {
	s.NotifyUrl = &v
	return s
}

func (s *DescribeLiveDetectNotifyConfigResponseBodyLiveDetectNotifyConfig) Validate() error {
	return dara.Validate(s)
}
