// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveAudioAuditConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveAudioAuditConfigList(v *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigList) *DescribeLiveAudioAuditConfigResponseBody
	GetLiveAudioAuditConfigList() *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigList
	SetRequestId(v string) *DescribeLiveAudioAuditConfigResponseBody
	GetRequestId() *string
}

type DescribeLiveAudioAuditConfigResponseBody struct {
	// The list of audio moderation configurations.
	LiveAudioAuditConfigList *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigList `json:"LiveAudioAuditConfigList,omitempty" xml:"LiveAudioAuditConfigList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7BF95F2A-3B24-4CDE-9346-7F6FA86697A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveAudioAuditConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAudioAuditConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditConfigResponseBody) GetLiveAudioAuditConfigList() *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigList {
	return s.LiveAudioAuditConfigList
}

func (s *DescribeLiveAudioAuditConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveAudioAuditConfigResponseBody) SetLiveAudioAuditConfigList(v *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigList) *DescribeLiveAudioAuditConfigResponseBody {
	s.LiveAudioAuditConfigList = v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponseBody) SetRequestId(v string) *DescribeLiveAudioAuditConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigList struct {
	LiveAudioAuditConfig []*DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig `json:"LiveAudioAuditConfig,omitempty" xml:"LiveAudioAuditConfig,omitempty" type:"Repeated"`
}

func (s DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigList) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigList) GetLiveAudioAuditConfig() []*DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig {
	return s.LiveAudioAuditConfig
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigList) SetLiveAudioAuditConfig(v []*DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig) *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigList {
	s.LiveAudioAuditConfig = v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigList) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The business type. You can specify a model. The default value is the domain name.
	//
	// example:
	//
	// example.edu
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The main streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The moderation scenarios.
	Scenes *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfigScenes `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Struct"`
	// The name of the live stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig) GetBizType() *string {
	return s.BizType
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig) GetScenes() *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfigScenes {
	return s.Scenes
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig) SetAppName(v string) *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig {
	s.AppName = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig) SetBizType(v string) *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig {
	s.BizType = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig) SetDomainName(v string) *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig) SetScenes(v *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfigScenes) *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig {
	s.Scenes = v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig) SetStreamName(v string) *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfigScenes struct {
	Scene []*string `json:"scene,omitempty" xml:"scene,omitempty" type:"Repeated"`
}

func (s DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfigScenes) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfigScenes) GoString() string {
	return s.String()
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfigScenes) GetScene() []*string {
	return s.Scene
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfigScenes) SetScene(v []*string) *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfigScenes {
	s.Scene = v
	return s
}

func (s *DescribeLiveAudioAuditConfigResponseBodyLiveAudioAuditConfigListLiveAudioAuditConfigScenes) Validate() error {
	return dara.Validate(s)
}
