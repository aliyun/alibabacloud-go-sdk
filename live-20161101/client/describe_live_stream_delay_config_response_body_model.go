// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamDelayConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveStreamFlvDelayConfig(v *DescribeLiveStreamDelayConfigResponseBodyLiveStreamFlvDelayConfig) *DescribeLiveStreamDelayConfigResponseBody
	GetLiveStreamFlvDelayConfig() *DescribeLiveStreamDelayConfigResponseBodyLiveStreamFlvDelayConfig
	SetLiveStreamHlsDelayConfig(v *DescribeLiveStreamDelayConfigResponseBodyLiveStreamHlsDelayConfig) *DescribeLiveStreamDelayConfigResponseBody
	GetLiveStreamHlsDelayConfig() *DescribeLiveStreamDelayConfigResponseBodyLiveStreamHlsDelayConfig
	SetLiveStreamRtmpDelayConfig(v *DescribeLiveStreamDelayConfigResponseBodyLiveStreamRtmpDelayConfig) *DescribeLiveStreamDelayConfigResponseBody
	GetLiveStreamRtmpDelayConfig() *DescribeLiveStreamDelayConfigResponseBodyLiveStreamRtmpDelayConfig
	SetRequestId(v string) *DescribeLiveStreamDelayConfigResponseBody
	GetRequestId() *string
}

type DescribeLiveStreamDelayConfigResponseBody struct {
	// The latency of FLV-based playback.
	LiveStreamFlvDelayConfig *DescribeLiveStreamDelayConfigResponseBodyLiveStreamFlvDelayConfig `json:"LiveStreamFlvDelayConfig,omitempty" xml:"LiveStreamFlvDelayConfig,omitempty" type:"Struct"`
	// The latency of HLS-based playback.
	LiveStreamHlsDelayConfig *DescribeLiveStreamDelayConfigResponseBodyLiveStreamHlsDelayConfig `json:"LiveStreamHlsDelayConfig,omitempty" xml:"LiveStreamHlsDelayConfig,omitempty" type:"Struct"`
	// The latency of RTMP-based playback.
	LiveStreamRtmpDelayConfig *DescribeLiveStreamDelayConfigResponseBodyLiveStreamRtmpDelayConfig `json:"LiveStreamRtmpDelayConfig,omitempty" xml:"LiveStreamRtmpDelayConfig,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 785C9CB0-EB8A-4912-BBF2-966BEFD32DC3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveStreamDelayConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamDelayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamDelayConfigResponseBody) GetLiveStreamFlvDelayConfig() *DescribeLiveStreamDelayConfigResponseBodyLiveStreamFlvDelayConfig {
	return s.LiveStreamFlvDelayConfig
}

func (s *DescribeLiveStreamDelayConfigResponseBody) GetLiveStreamHlsDelayConfig() *DescribeLiveStreamDelayConfigResponseBodyLiveStreamHlsDelayConfig {
	return s.LiveStreamHlsDelayConfig
}

func (s *DescribeLiveStreamDelayConfigResponseBody) GetLiveStreamRtmpDelayConfig() *DescribeLiveStreamDelayConfigResponseBodyLiveStreamRtmpDelayConfig {
	return s.LiveStreamRtmpDelayConfig
}

func (s *DescribeLiveStreamDelayConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamDelayConfigResponseBody) SetLiveStreamFlvDelayConfig(v *DescribeLiveStreamDelayConfigResponseBodyLiveStreamFlvDelayConfig) *DescribeLiveStreamDelayConfigResponseBody {
	s.LiveStreamFlvDelayConfig = v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponseBody) SetLiveStreamHlsDelayConfig(v *DescribeLiveStreamDelayConfigResponseBodyLiveStreamHlsDelayConfig) *DescribeLiveStreamDelayConfigResponseBody {
	s.LiveStreamHlsDelayConfig = v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponseBody) SetLiveStreamRtmpDelayConfig(v *DescribeLiveStreamDelayConfigResponseBodyLiveStreamRtmpDelayConfig) *DescribeLiveStreamDelayConfigResponseBody {
	s.LiveStreamRtmpDelayConfig = v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponseBody) SetRequestId(v string) *DescribeLiveStreamDelayConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamDelayConfigResponseBodyLiveStreamFlvDelayConfig struct {
	// The playback latency. Unit: seconds.
	//
	// example:
	//
	// 5
	Delay *int32 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The latency level. Valid values:
	//
	// 	- **short**: The latency is less than or equal to 4 seconds.
	//
	// 	- **medium**: The latency is greater than 4 seconds, and less than or equal to 8 seconds.
	//
	// 	- **long**: The latency is greater than 8 seconds.
	//
	// example:
	//
	// medium
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeLiveStreamDelayConfigResponseBodyLiveStreamFlvDelayConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamDelayConfigResponseBodyLiveStreamFlvDelayConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamDelayConfigResponseBodyLiveStreamFlvDelayConfig) GetDelay() *int32 {
	return s.Delay
}

func (s *DescribeLiveStreamDelayConfigResponseBodyLiveStreamFlvDelayConfig) GetLevel() *string {
	return s.Level
}

func (s *DescribeLiveStreamDelayConfigResponseBodyLiveStreamFlvDelayConfig) SetDelay(v int32) *DescribeLiveStreamDelayConfigResponseBodyLiveStreamFlvDelayConfig {
	s.Delay = &v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponseBodyLiveStreamFlvDelayConfig) SetLevel(v string) *DescribeLiveStreamDelayConfigResponseBodyLiveStreamFlvDelayConfig {
	s.Level = &v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponseBodyLiveStreamFlvDelayConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamDelayConfigResponseBodyLiveStreamHlsDelayConfig struct {
	// The playback latency. Unit: seconds.
	//
	// example:
	//
	// 3
	Delay *int32 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The latency level. Valid values:
	//
	// 	- **short**: The latency is less than or equal to 4 seconds.
	//
	// 	- **medium**: The latency is greater than 4 seconds, and less than or equal to 8 seconds.
	//
	// 	- **long**: The latency is greater than 8 seconds.
	//
	// example:
	//
	// short
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeLiveStreamDelayConfigResponseBodyLiveStreamHlsDelayConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamDelayConfigResponseBodyLiveStreamHlsDelayConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamDelayConfigResponseBodyLiveStreamHlsDelayConfig) GetDelay() *int32 {
	return s.Delay
}

func (s *DescribeLiveStreamDelayConfigResponseBodyLiveStreamHlsDelayConfig) GetLevel() *string {
	return s.Level
}

func (s *DescribeLiveStreamDelayConfigResponseBodyLiveStreamHlsDelayConfig) SetDelay(v int32) *DescribeLiveStreamDelayConfigResponseBodyLiveStreamHlsDelayConfig {
	s.Delay = &v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponseBodyLiveStreamHlsDelayConfig) SetLevel(v string) *DescribeLiveStreamDelayConfigResponseBodyLiveStreamHlsDelayConfig {
	s.Level = &v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponseBodyLiveStreamHlsDelayConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamDelayConfigResponseBodyLiveStreamRtmpDelayConfig struct {
	// The playback latency. Unit: seconds.
	//
	// example:
	//
	// 4
	Delay *int32 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The latency level. Valid values:
	//
	// 	- **short**: The latency is less than or equal to 4 seconds.
	//
	// 	- **medium**: The latency is greater than 4 seconds, and less than or equal to 8 seconds.
	//
	// 	- **long**: The latency is greater than 8 seconds.
	//
	// example:
	//
	// short
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeLiveStreamDelayConfigResponseBodyLiveStreamRtmpDelayConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamDelayConfigResponseBodyLiveStreamRtmpDelayConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamDelayConfigResponseBodyLiveStreamRtmpDelayConfig) GetDelay() *int32 {
	return s.Delay
}

func (s *DescribeLiveStreamDelayConfigResponseBodyLiveStreamRtmpDelayConfig) GetLevel() *string {
	return s.Level
}

func (s *DescribeLiveStreamDelayConfigResponseBodyLiveStreamRtmpDelayConfig) SetDelay(v int32) *DescribeLiveStreamDelayConfigResponseBodyLiveStreamRtmpDelayConfig {
	s.Delay = &v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponseBodyLiveStreamRtmpDelayConfig) SetLevel(v string) *DescribeLiveStreamDelayConfigResponseBodyLiveStreamRtmpDelayConfig {
	s.Level = &v
	return s
}

func (s *DescribeLiveStreamDelayConfigResponseBodyLiveStreamRtmpDelayConfig) Validate() error {
	return dara.Validate(s)
}
