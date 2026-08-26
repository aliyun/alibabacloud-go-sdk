// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveStreamMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *CreateLiveStreamMonitorRequest
	GetApp() *string
	SetCallbackUrl(v string) *CreateLiveStreamMonitorRequest
	GetCallbackUrl() *string
	SetDingTalkWebHookUrl(v string) *CreateLiveStreamMonitorRequest
	GetDingTalkWebHookUrl() *string
	SetDomain(v string) *CreateLiveStreamMonitorRequest
	GetDomain() *string
	SetInputList(v string) *CreateLiveStreamMonitorRequest
	GetInputList() *string
	SetMonitorConfig(v string) *CreateLiveStreamMonitorRequest
	GetMonitorConfig() *string
	SetMonitorName(v string) *CreateLiveStreamMonitorRequest
	GetMonitorName() *string
	SetOutputTemplate(v string) *CreateLiveStreamMonitorRequest
	GetOutputTemplate() *string
	SetOwnerId(v int64) *CreateLiveStreamMonitorRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateLiveStreamMonitorRequest
	GetRegionId() *string
	SetStream(v string) *CreateLiveStreamMonitorRequest
	GetStream() *string
}

type CreateLiveStreamMonitorRequest struct {
	// The application name for the output stream of the monitoring session. You can specify a custom name. If you do not specify this parameter, **monitor*	- is used as the application name.
	//
	// example:
	//
	// monitor****
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The webhook address. HTTP and HTTPS are supported.
	//
	// example:
	//
	// http://guide.aliyundoc.com/notify
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The webhook URL of the DingTalk chatbot. To receive alert notifications, configure a DingTalk chatbot and enter its webhook URL, which can be in HTTP or HTTPS format. For more information, see [Custom robot access](https://open.dingtalk.com/document/robots/custom-robot-access).
	//
	// > Set the custom keyword for the DingTalk chatbot to "Alerting". Otherwise, you cannot receive messages.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=7a7d404056eee1f2fd944ace9bcfc361dc6448583e1d3d3baa****
	DingTalkWebHookUrl *string `json:"DingTalkWebHookUrl,omitempty" xml:"DingTalkWebHookUrl,omitempty"`
	// The domain name to monitor.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The list of input streams to monitor. For more information, see the **InputConfig*	- table below.
	//
	// This parameter is required.
	//
	// example:
	//
	// InputConfig
	InputList *string `json:"InputList,omitempty" xml:"InputList,omitempty"`
	// The alert threshold settings for monitoring. The value must be a JSON string. For more information, see the MonitorConfig table below.
	//
	// example:
	//
	// "{\\"fpsLowThres\\": 0.6,\\"brLowThres\\": 1.1,\\"eofDurationThresSec\\": 10}"
	MonitorConfig *string `json:"MonitorConfig,omitempty" xml:"MonitorConfig,omitempty"`
	// The name of the monitoring session.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveMonitor****
	MonitorName *string `json:"MonitorName,omitempty" xml:"MonitorName,omitempty"`
	// The output template for the monitoring session. Valid values:
	//
	// - **lp_ld**: low definition.
	//
	// - **lp_sd**: standard definition.
	//
	// - **lp_hd**: high definition.
	//
	// - **lp_ud**: ultra-high definition.
	//
	// This parameter is required.
	//
	// example:
	//
	// lp_ud
	OutputTemplate *string `json:"OutputTemplate,omitempty" xml:"OutputTemplate,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the output stream for the monitoring session. If you do not specify this parameter, the system generates a random name.
	//
	// example:
	//
	// monitorStream****
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s CreateLiveStreamMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveStreamMonitorRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveStreamMonitorRequest) GetApp() *string {
	return s.App
}

func (s *CreateLiveStreamMonitorRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *CreateLiveStreamMonitorRequest) GetDingTalkWebHookUrl() *string {
	return s.DingTalkWebHookUrl
}

func (s *CreateLiveStreamMonitorRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateLiveStreamMonitorRequest) GetInputList() *string {
	return s.InputList
}

func (s *CreateLiveStreamMonitorRequest) GetMonitorConfig() *string {
	return s.MonitorConfig
}

func (s *CreateLiveStreamMonitorRequest) GetMonitorName() *string {
	return s.MonitorName
}

func (s *CreateLiveStreamMonitorRequest) GetOutputTemplate() *string {
	return s.OutputTemplate
}

func (s *CreateLiveStreamMonitorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLiveStreamMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLiveStreamMonitorRequest) GetStream() *string {
	return s.Stream
}

func (s *CreateLiveStreamMonitorRequest) SetApp(v string) *CreateLiveStreamMonitorRequest {
	s.App = &v
	return s
}

func (s *CreateLiveStreamMonitorRequest) SetCallbackUrl(v string) *CreateLiveStreamMonitorRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateLiveStreamMonitorRequest) SetDingTalkWebHookUrl(v string) *CreateLiveStreamMonitorRequest {
	s.DingTalkWebHookUrl = &v
	return s
}

func (s *CreateLiveStreamMonitorRequest) SetDomain(v string) *CreateLiveStreamMonitorRequest {
	s.Domain = &v
	return s
}

func (s *CreateLiveStreamMonitorRequest) SetInputList(v string) *CreateLiveStreamMonitorRequest {
	s.InputList = &v
	return s
}

func (s *CreateLiveStreamMonitorRequest) SetMonitorConfig(v string) *CreateLiveStreamMonitorRequest {
	s.MonitorConfig = &v
	return s
}

func (s *CreateLiveStreamMonitorRequest) SetMonitorName(v string) *CreateLiveStreamMonitorRequest {
	s.MonitorName = &v
	return s
}

func (s *CreateLiveStreamMonitorRequest) SetOutputTemplate(v string) *CreateLiveStreamMonitorRequest {
	s.OutputTemplate = &v
	return s
}

func (s *CreateLiveStreamMonitorRequest) SetOwnerId(v int64) *CreateLiveStreamMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLiveStreamMonitorRequest) SetRegionId(v string) *CreateLiveStreamMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLiveStreamMonitorRequest) SetStream(v string) *CreateLiveStreamMonitorRequest {
	s.Stream = &v
	return s
}

func (s *CreateLiveStreamMonitorRequest) Validate() error {
	return dara.Validate(s)
}
