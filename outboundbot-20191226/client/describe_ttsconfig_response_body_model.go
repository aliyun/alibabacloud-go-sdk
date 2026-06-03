// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTTSConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeTTSConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeTTSConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeTTSConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeTTSConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeTTSConfigResponseBody
	GetSuccess() *bool
	SetTTSConfig(v *DescribeTTSConfigResponseBodyTTSConfig) *DescribeTTSConfigResponseBody
	GetTTSConfig() *DescribeTTSConfigResponseBodyTTSConfig
}

type DescribeTTSConfigResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// f765d3ee-ec03-4765-b235-6877501d99d1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TTSConfig *DescribeTTSConfigResponseBodyTTSConfig `json:"TTSConfig,omitempty" xml:"TTSConfig,omitempty" type:"Struct"`
}

func (s DescribeTTSConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTTSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTTSConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeTTSConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeTTSConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTTSConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTTSConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTTSConfigResponseBody) GetTTSConfig() *DescribeTTSConfigResponseBodyTTSConfig {
	return s.TTSConfig
}

func (s *DescribeTTSConfigResponseBody) SetCode(v string) *DescribeTTSConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetHttpStatusCode(v int32) *DescribeTTSConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetMessage(v string) *DescribeTTSConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetRequestId(v string) *DescribeTTSConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetSuccess(v bool) *DescribeTTSConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTTSConfigResponseBody) SetTTSConfig(v *DescribeTTSConfigResponseBodyTTSConfig) *DescribeTTSConfigResponseBody {
	s.TTSConfig = v
	return s
}

func (s *DescribeTTSConfigResponseBody) Validate() error {
	if s.TTSConfig != nil {
		if err := s.TTSConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTTSConfigResponseBodyTTSConfig struct {
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 947e0875-b5d4-4b33-b18c-7b2cf85bcb4f
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Managed
	NlsServiceType *string `json:"NlsServiceType,omitempty" xml:"NlsServiceType,omitempty"`
	// example:
	//
	// -150
	SpeechRate *string `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// example:
	//
	// xiaoyun
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// example:
	//
	// 100
	Volume *string `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s DescribeTTSConfigResponseBodyTTSConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeTTSConfigResponseBodyTTSConfig) GoString() string {
	return s.String()
}

func (s *DescribeTTSConfigResponseBodyTTSConfig) GetAppKey() *string {
	return s.AppKey
}

func (s *DescribeTTSConfigResponseBodyTTSConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTTSConfigResponseBodyTTSConfig) GetNlsServiceType() *string {
	return s.NlsServiceType
}

func (s *DescribeTTSConfigResponseBodyTTSConfig) GetSpeechRate() *string {
	return s.SpeechRate
}

func (s *DescribeTTSConfigResponseBodyTTSConfig) GetVoice() *string {
	return s.Voice
}

func (s *DescribeTTSConfigResponseBodyTTSConfig) GetVolume() *string {
	return s.Volume
}

func (s *DescribeTTSConfigResponseBodyTTSConfig) SetAppKey(v string) *DescribeTTSConfigResponseBodyTTSConfig {
	s.AppKey = &v
	return s
}

func (s *DescribeTTSConfigResponseBodyTTSConfig) SetInstanceId(v string) *DescribeTTSConfigResponseBodyTTSConfig {
	s.InstanceId = &v
	return s
}

func (s *DescribeTTSConfigResponseBodyTTSConfig) SetNlsServiceType(v string) *DescribeTTSConfigResponseBodyTTSConfig {
	s.NlsServiceType = &v
	return s
}

func (s *DescribeTTSConfigResponseBodyTTSConfig) SetSpeechRate(v string) *DescribeTTSConfigResponseBodyTTSConfig {
	s.SpeechRate = &v
	return s
}

func (s *DescribeTTSConfigResponseBodyTTSConfig) SetVoice(v string) *DescribeTTSConfigResponseBodyTTSConfig {
	s.Voice = &v
	return s
}

func (s *DescribeTTSConfigResponseBodyTTSConfig) SetVolume(v string) *DescribeTTSConfigResponseBodyTTSConfig {
	s.Volume = &v
	return s
}

func (s *DescribeTTSConfigResponseBodyTTSConfig) Validate() error {
	return dara.Validate(s)
}
