// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveAudioAuditNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveAudioAuditNotifyConfigList(v *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigList) *DescribeLiveAudioAuditNotifyConfigResponseBody
	GetLiveAudioAuditNotifyConfigList() *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigList
	SetRequestId(v string) *DescribeLiveAudioAuditNotifyConfigResponseBody
	GetRequestId() *string
}

type DescribeLiveAudioAuditNotifyConfigResponseBody struct {
	// The configuration of callbacks for audio moderation results.
	LiveAudioAuditNotifyConfigList *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigList `json:"LiveAudioAuditNotifyConfigList,omitempty" xml:"LiveAudioAuditNotifyConfigList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B908FF89-B03C-4831-B55B-48D2A******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveAudioAuditNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAudioAuditNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseBody) GetLiveAudioAuditNotifyConfigList() *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigList {
	return s.LiveAudioAuditNotifyConfigList
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseBody) SetLiveAudioAuditNotifyConfigList(v *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigList) *DescribeLiveAudioAuditNotifyConfigResponseBody {
	s.LiveAudioAuditNotifyConfigList = v
	return s
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseBody) SetRequestId(v string) *DescribeLiveAudioAuditNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigList struct {
	LiveAudioAuditNotifyConfig []*DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig `json:"LiveAudioAuditNotifyConfig,omitempty" xml:"LiveAudioAuditNotifyConfig,omitempty" type:"Repeated"`
}

func (s DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigList) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigList) GetLiveAudioAuditNotifyConfig() []*DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig {
	return s.LiveAudioAuditNotifyConfig
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigList) SetLiveAudioAuditNotifyConfig(v []*DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig) *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigList {
	s.LiveAudioAuditNotifyConfig = v
	return s
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigList) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig struct {
	// The callback URL.
	//
	// example:
	//
	// http://guide.aliyundoc.com/callback
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// The callback template. The following fields are configured:
	//
	// 	- **{DomainName}**: the streaming domain.
	//
	// 	- **{AppName}**: the name of the application to which the live stream belongs.
	//
	// 	- **{StreamName}**: the name of the live stream.
	//
	// 	- **{Timestamp}**: the time when the callback is returned. The value of this field is a UNIX timestamp. Unit: seconds.
	//
	// 	- **{Result}**: the moderation results.
	//
	// example:
	//
	// {\\"domain\\":{DomainName},\\"app\\":{AppName},\\"stream\\":{StreamName},\\"timestamp\\":{Timestamp},\\"result\\":{Result}}
	CallbackTemplate *string `json:"CallbackTemplate,omitempty" xml:"CallbackTemplate,omitempty"`
	// The main streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig) GetCallback() *string {
	return s.Callback
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig) GetCallbackTemplate() *string {
	return s.CallbackTemplate
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig) SetCallback(v string) *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig {
	s.Callback = &v
	return s
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig) SetCallbackTemplate(v string) *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig {
	s.CallbackTemplate = &v
	return s
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig) SetDomainName(v string) *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveAudioAuditNotifyConfigResponseBodyLiveAudioAuditNotifyConfigListLiveAudioAuditNotifyConfig) Validate() error {
	return dara.Validate(s)
}
