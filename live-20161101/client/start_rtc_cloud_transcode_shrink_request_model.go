// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRtcCloudTranscodeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartRtcCloudTranscodeShrinkRequest
	GetAppId() *string
	SetChannelId(v string) *StartRtcCloudTranscodeShrinkRequest
	GetChannelId() *string
	SetInputParamShrink(v string) *StartRtcCloudTranscodeShrinkRequest
	GetInputParamShrink() *string
	SetMaxIdleTime(v int64) *StartRtcCloudTranscodeShrinkRequest
	GetMaxIdleTime() *int64
	SetOutputParamsShrink(v string) *StartRtcCloudTranscodeShrinkRequest
	GetOutputParamsShrink() *string
}

type StartRtcCloudTranscodeShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ********-7074-****-9ef5-85c19a4*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// myChannel
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	InputParamShrink *string `json:"InputParam,omitempty" xml:"InputParam,omitempty"`
	// example:
	//
	// 600
	MaxIdleTime *int64 `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	// This parameter is required.
	OutputParamsShrink *string `json:"OutputParams,omitempty" xml:"OutputParams,omitempty"`
}

func (s StartRtcCloudTranscodeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudTranscodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartRtcCloudTranscodeShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartRtcCloudTranscodeShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartRtcCloudTranscodeShrinkRequest) GetInputParamShrink() *string {
	return s.InputParamShrink
}

func (s *StartRtcCloudTranscodeShrinkRequest) GetMaxIdleTime() *int64 {
	return s.MaxIdleTime
}

func (s *StartRtcCloudTranscodeShrinkRequest) GetOutputParamsShrink() *string {
	return s.OutputParamsShrink
}

func (s *StartRtcCloudTranscodeShrinkRequest) SetAppId(v string) *StartRtcCloudTranscodeShrinkRequest {
	s.AppId = &v
	return s
}

func (s *StartRtcCloudTranscodeShrinkRequest) SetChannelId(v string) *StartRtcCloudTranscodeShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *StartRtcCloudTranscodeShrinkRequest) SetInputParamShrink(v string) *StartRtcCloudTranscodeShrinkRequest {
	s.InputParamShrink = &v
	return s
}

func (s *StartRtcCloudTranscodeShrinkRequest) SetMaxIdleTime(v int64) *StartRtcCloudTranscodeShrinkRequest {
	s.MaxIdleTime = &v
	return s
}

func (s *StartRtcCloudTranscodeShrinkRequest) SetOutputParamsShrink(v string) *StartRtcCloudTranscodeShrinkRequest {
	s.OutputParamsShrink = &v
	return s
}

func (s *StartRtcCloudTranscodeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
