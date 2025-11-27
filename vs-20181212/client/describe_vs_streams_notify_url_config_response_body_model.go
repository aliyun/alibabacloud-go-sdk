// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsStreamsNotifyUrlConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveStreamsNotifyConfig(v *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) *DescribeVsStreamsNotifyUrlConfigResponseBody
	GetLiveStreamsNotifyConfig() *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig
	SetRequestId(v string) *DescribeVsStreamsNotifyUrlConfigResponseBody
	GetRequestId() *string
}

type DescribeVsStreamsNotifyUrlConfigResponseBody struct {
	LiveStreamsNotifyConfig *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig `json:"LiveStreamsNotifyConfig,omitempty" xml:"LiveStreamsNotifyConfig,omitempty" type:"Struct"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsStreamsNotifyUrlConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsStreamsNotifyUrlConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBody) GetLiveStreamsNotifyConfig() *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig {
	return s.LiveStreamsNotifyConfig
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBody) SetLiveStreamsNotifyConfig(v *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) *DescribeVsStreamsNotifyUrlConfigResponseBody {
	s.LiveStreamsNotifyConfig = v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBody) SetRequestId(v string) *DescribeVsStreamsNotifyUrlConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBody) Validate() error {
	if s.LiveStreamsNotifyConfig != nil {
		if err := s.LiveStreamsNotifyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig struct {
	AuthKey  *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	NotifyUrl  *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
}

func (s DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) GetAuthKey() *string {
	return s.AuthKey
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) SetAuthKey(v string) *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig {
	s.AuthKey = &v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) SetAuthType(v string) *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig {
	s.AuthType = &v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) SetDomainName(v string) *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) SetNotifyUrl(v string) *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig {
	s.NotifyUrl = &v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigResponseBodyLiveStreamsNotifyConfig) Validate() error {
	return dara.Validate(s)
}
