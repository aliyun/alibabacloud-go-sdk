// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveStreamMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *UpdateLiveStreamMonitorRequest
	GetApp() *string
	SetCallbackUrl(v string) *UpdateLiveStreamMonitorRequest
	GetCallbackUrl() *string
	SetDingTalkWebHookUrl(v string) *UpdateLiveStreamMonitorRequest
	GetDingTalkWebHookUrl() *string
	SetDomain(v string) *UpdateLiveStreamMonitorRequest
	GetDomain() *string
	SetInputList(v string) *UpdateLiveStreamMonitorRequest
	GetInputList() *string
	SetMonitorConfig(v string) *UpdateLiveStreamMonitorRequest
	GetMonitorConfig() *string
	SetMonitorId(v string) *UpdateLiveStreamMonitorRequest
	GetMonitorId() *string
	SetMonitorName(v string) *UpdateLiveStreamMonitorRequest
	GetMonitorName() *string
	SetOutputTemplate(v string) *UpdateLiveStreamMonitorRequest
	GetOutputTemplate() *string
	SetOwnerId(v int64) *UpdateLiveStreamMonitorRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateLiveStreamMonitorRequest
	GetRegionId() *string
	SetStream(v string) *UpdateLiveStreamMonitorRequest
	GetStream() *string
}

type UpdateLiveStreamMonitorRequest struct {
	// The application name for the output stream of the monitoring session. You can specify a custom name. If you do not specify this parameter, **monitor*	- is used as the AppName.
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
	// The webhook URL of the DingTalk chatbot. Monitoring alerts are sent to a DingTalk group using a chatbot. Set up a chatbot and enter its webhook URL, which must be an HTTP or HTTPS address. For more information, see [Custom robot access](https://open.dingtalk.com/document/robots/custom-robot-access).
	//
	// > Set the custom keyword of the DingTalk chatbot to "Alerting". Otherwise, messages cannot be received.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=7a7d404056eee1f2fd944ace9bcfc361dc6448583e1d3d3baa****
	DingTalkWebHookUrl *string `json:"DingTalkWebHookUrl,omitempty" xml:"DingTalkWebHookUrl,omitempty"`
	// The output domain name for the monitoring session.
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
	// The settings for alert thresholds. The value is a JSON string. For more information, see the MonitorConfig table below.
	//
	// example:
	//
	// "{\\"fpsLowThres\\": 0.6,\\"brLowThres\\": 1.1,\\"eofDurationThresSec\\": 10}"
	MonitorConfig *string `json:"MonitorConfig,omitempty" xml:"MonitorConfig,omitempty"`
	// The ID of the monitoring session.
	//
	// > Obtain the MonitorId value from the response parameters of the [CreateLiveStreamMonitor](https://help.aliyun.com/document_detail/2848129.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	MonitorId *string `json:"MonitorId,omitempty" xml:"MonitorId,omitempty"`
	// The name of the monitoring session.
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
	// The name of the output stream for the monitoring session.
	//
	// example:
	//
	// monitorStream****
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s UpdateLiveStreamMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveStreamMonitorRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveStreamMonitorRequest) GetApp() *string {
	return s.App
}

func (s *UpdateLiveStreamMonitorRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *UpdateLiveStreamMonitorRequest) GetDingTalkWebHookUrl() *string {
	return s.DingTalkWebHookUrl
}

func (s *UpdateLiveStreamMonitorRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateLiveStreamMonitorRequest) GetInputList() *string {
	return s.InputList
}

func (s *UpdateLiveStreamMonitorRequest) GetMonitorConfig() *string {
	return s.MonitorConfig
}

func (s *UpdateLiveStreamMonitorRequest) GetMonitorId() *string {
	return s.MonitorId
}

func (s *UpdateLiveStreamMonitorRequest) GetMonitorName() *string {
	return s.MonitorName
}

func (s *UpdateLiveStreamMonitorRequest) GetOutputTemplate() *string {
	return s.OutputTemplate
}

func (s *UpdateLiveStreamMonitorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveStreamMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLiveStreamMonitorRequest) GetStream() *string {
	return s.Stream
}

func (s *UpdateLiveStreamMonitorRequest) SetApp(v string) *UpdateLiveStreamMonitorRequest {
	s.App = &v
	return s
}

func (s *UpdateLiveStreamMonitorRequest) SetCallbackUrl(v string) *UpdateLiveStreamMonitorRequest {
	s.CallbackUrl = &v
	return s
}

func (s *UpdateLiveStreamMonitorRequest) SetDingTalkWebHookUrl(v string) *UpdateLiveStreamMonitorRequest {
	s.DingTalkWebHookUrl = &v
	return s
}

func (s *UpdateLiveStreamMonitorRequest) SetDomain(v string) *UpdateLiveStreamMonitorRequest {
	s.Domain = &v
	return s
}

func (s *UpdateLiveStreamMonitorRequest) SetInputList(v string) *UpdateLiveStreamMonitorRequest {
	s.InputList = &v
	return s
}

func (s *UpdateLiveStreamMonitorRequest) SetMonitorConfig(v string) *UpdateLiveStreamMonitorRequest {
	s.MonitorConfig = &v
	return s
}

func (s *UpdateLiveStreamMonitorRequest) SetMonitorId(v string) *UpdateLiveStreamMonitorRequest {
	s.MonitorId = &v
	return s
}

func (s *UpdateLiveStreamMonitorRequest) SetMonitorName(v string) *UpdateLiveStreamMonitorRequest {
	s.MonitorName = &v
	return s
}

func (s *UpdateLiveStreamMonitorRequest) SetOutputTemplate(v string) *UpdateLiveStreamMonitorRequest {
	s.OutputTemplate = &v
	return s
}

func (s *UpdateLiveStreamMonitorRequest) SetOwnerId(v int64) *UpdateLiveStreamMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveStreamMonitorRequest) SetRegionId(v string) *UpdateLiveStreamMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLiveStreamMonitorRequest) SetStream(v string) *UpdateLiveStreamMonitorRequest {
	s.Stream = &v
	return s
}

func (s *UpdateLiveStreamMonitorRequest) Validate() error {
	return dara.Validate(s)
}
