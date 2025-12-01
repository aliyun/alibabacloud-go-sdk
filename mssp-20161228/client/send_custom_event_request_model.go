// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCustomEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerId(v string) *SendCustomEventRequest
	GetCustomerId() *string
	SetDataSource(v string) *SendCustomEventRequest
	GetDataSource() *string
	SetEventDescription(v string) *SendCustomEventRequest
	GetEventDescription() *string
	SetEventDetails(v string) *SendCustomEventRequest
	GetEventDetails() *string
	SetEventMarkdown(v string) *SendCustomEventRequest
	GetEventMarkdown() *string
	SetEventName(v string) *SendCustomEventRequest
	GetEventName() *string
	SetEventType(v string) *SendCustomEventRequest
	GetEventType() *string
	SetFindTime(v int64) *SendCustomEventRequest
	GetFindTime() *int64
	SetInstanceId(v string) *SendCustomEventRequest
	GetInstanceId() *string
	SetInstanceName(v string) *SendCustomEventRequest
	GetInstanceName() *string
	SetIsSend(v string) *SendCustomEventRequest
	GetIsSend() *string
	SetLevel(v string) *SendCustomEventRequest
	GetLevel() *string
	SetOccurrenceTime(v int64) *SendCustomEventRequest
	GetOccurrenceTime() *int64
	SetOwnerId(v string) *SendCustomEventRequest
	GetOwnerId() *string
	SetProduct(v string) *SendCustomEventRequest
	GetProduct() *string
	SetUniqueId(v string) *SendCustomEventRequest
	GetUniqueId() *string
	SetUuid(v string) *SendCustomEventRequest
	GetUuid() *string
}

type SendCustomEventRequest struct {
	// User ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1214484929940219
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// Data source.
	//
	// example:
	//
	// aegis_suspicious_event
	DataSource *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// Event details.
	//
	// example:
	//
	// 疑似病毒木马启动运行。
	EventDescription *string `json:"EventDescription,omitempty" xml:"EventDescription,omitempty"`
	// Alert event details.
	//
	// example:
	//
	// [{"name":"提示","type":"text","value":"在您的系统上发现可疑进程启动行为，通常与病毒木马或入侵事件相关"},{"name":"ATT&CK攻击阶段","type":"text","value":"代码执行"},{"name":"恶意行为","type":"text","value":"可疑的漏洞利用代码执行"},{"name":"规则类型","type":"text","value":"进程启动"},{"name":"规则引擎","type":"text","value":"精准攻击识别引擎"},{"name":"处置动作","type":"text","value":"阻断行为"},{"name":"进程路径","type":"text","value":"/usr/bin/python3.9"},{"name":"命令行","type":"text","value":"python3 /root/poc/python/cve-2018-15473.py --username root --port 22"},{"name":"父进程路径","type":"text","value":"/bin/gunkit"},{"name":"父进程命令行","type":"text","value":"gunkit serve-grpc --addr unix:///data/gunkit-grpc.sock"},{"name":"进程ID","type":"text","value":"22714"},{"name":"父进程ID","type":"text","value":"2986"},{"name":"描述","type":"text","value":"主动防御检测到可疑进程启动行为，这类可疑进程通常存在于特殊的系统目录，或通过后缀伪装成文档/音频/图片等文件诱导用户运行，该异常行为已被成功拦截"},{"name":"处置建议","type":"text","value":"请您及时排查是否是正常的业务操作，如果您觉得此次拦截是非预期的，那您可以在主动防御 - 恶意行为防御页面中，关闭“可疑进程启动“规则集或者将受影响机器从管理主机中移除"},{"name":"父进程关系","type":"processChain","value":"1:::/usr/lib/systemd/systemd --switched-root --system --deserialize 22 &&& 2939:::/usr/local/bin/containerd-shim-runc-v2 -namespace moby -id 270f164903b47d4e219b410b8d11d9079a7ad1bac8133aea604598300d3b03d5 -address /run/containerd/containerd.sock &&& 2962:::/usr/bin/python3 /usr/bin/supervisord -n &&& 2986:::gunkit serve-grpc --addr unix:///data/gunkit-grpc.sock"}]
	EventDetails *string `json:"EventDetails,omitempty" xml:"EventDetails,omitempty"`
	// Details of markdown format
	//
	// example:
	//
	// None
	EventMarkdown *string `json:"EventMarkdown,omitempty" xml:"EventMarkdown,omitempty"`
	// Event name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 客户端离线 Client offline
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// Event type.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUSP_CUSTOM_CFW
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// Alert discovery time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-23 14:47:34
	FindTime *int64 `json:"FindTime,omitempty" xml:"FindTime,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// i-uf60h3ns25bzq9eyf8ps
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Instance name.
	//
	// example:
	//
	// 猫吉-售卖-MDR扫描器集群1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Whether to send.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	IsSend *string `json:"IsSend,omitempty" xml:"IsSend,omitempty"`
	// Event level.
	//
	// This parameter is required.
	//
	// example:
	//
	// serious
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The first occurrence time of the alert event.
	//
	// example:
	//
	// 1724956996000
	OccurrenceTime *int64  `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Product name.
	//
	// example:
	//
	// CloudSecCenter
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// Unique ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 68888f02-98f2-492b-a2b2-5b13295755b7
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	// UUID.
	//
	// example:
	//
	// 93B6CDAB-7D2E-33D2-9EBA-25D561A2E95F
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s SendCustomEventRequest) String() string {
	return dara.Prettify(s)
}

func (s SendCustomEventRequest) GoString() string {
	return s.String()
}

func (s *SendCustomEventRequest) GetCustomerId() *string {
	return s.CustomerId
}

func (s *SendCustomEventRequest) GetDataSource() *string {
	return s.DataSource
}

func (s *SendCustomEventRequest) GetEventDescription() *string {
	return s.EventDescription
}

func (s *SendCustomEventRequest) GetEventDetails() *string {
	return s.EventDetails
}

func (s *SendCustomEventRequest) GetEventMarkdown() *string {
	return s.EventMarkdown
}

func (s *SendCustomEventRequest) GetEventName() *string {
	return s.EventName
}

func (s *SendCustomEventRequest) GetEventType() *string {
	return s.EventType
}

func (s *SendCustomEventRequest) GetFindTime() *int64 {
	return s.FindTime
}

func (s *SendCustomEventRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SendCustomEventRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *SendCustomEventRequest) GetIsSend() *string {
	return s.IsSend
}

func (s *SendCustomEventRequest) GetLevel() *string {
	return s.Level
}

func (s *SendCustomEventRequest) GetOccurrenceTime() *int64 {
	return s.OccurrenceTime
}

func (s *SendCustomEventRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SendCustomEventRequest) GetProduct() *string {
	return s.Product
}

func (s *SendCustomEventRequest) GetUniqueId() *string {
	return s.UniqueId
}

func (s *SendCustomEventRequest) GetUuid() *string {
	return s.Uuid
}

func (s *SendCustomEventRequest) SetCustomerId(v string) *SendCustomEventRequest {
	s.CustomerId = &v
	return s
}

func (s *SendCustomEventRequest) SetDataSource(v string) *SendCustomEventRequest {
	s.DataSource = &v
	return s
}

func (s *SendCustomEventRequest) SetEventDescription(v string) *SendCustomEventRequest {
	s.EventDescription = &v
	return s
}

func (s *SendCustomEventRequest) SetEventDetails(v string) *SendCustomEventRequest {
	s.EventDetails = &v
	return s
}

func (s *SendCustomEventRequest) SetEventMarkdown(v string) *SendCustomEventRequest {
	s.EventMarkdown = &v
	return s
}

func (s *SendCustomEventRequest) SetEventName(v string) *SendCustomEventRequest {
	s.EventName = &v
	return s
}

func (s *SendCustomEventRequest) SetEventType(v string) *SendCustomEventRequest {
	s.EventType = &v
	return s
}

func (s *SendCustomEventRequest) SetFindTime(v int64) *SendCustomEventRequest {
	s.FindTime = &v
	return s
}

func (s *SendCustomEventRequest) SetInstanceId(v string) *SendCustomEventRequest {
	s.InstanceId = &v
	return s
}

func (s *SendCustomEventRequest) SetInstanceName(v string) *SendCustomEventRequest {
	s.InstanceName = &v
	return s
}

func (s *SendCustomEventRequest) SetIsSend(v string) *SendCustomEventRequest {
	s.IsSend = &v
	return s
}

func (s *SendCustomEventRequest) SetLevel(v string) *SendCustomEventRequest {
	s.Level = &v
	return s
}

func (s *SendCustomEventRequest) SetOccurrenceTime(v int64) *SendCustomEventRequest {
	s.OccurrenceTime = &v
	return s
}

func (s *SendCustomEventRequest) SetOwnerId(v string) *SendCustomEventRequest {
	s.OwnerId = &v
	return s
}

func (s *SendCustomEventRequest) SetProduct(v string) *SendCustomEventRequest {
	s.Product = &v
	return s
}

func (s *SendCustomEventRequest) SetUniqueId(v string) *SendCustomEventRequest {
	s.UniqueId = &v
	return s
}

func (s *SendCustomEventRequest) SetUuid(v string) *SendCustomEventRequest {
	s.Uuid = &v
	return s
}

func (s *SendCustomEventRequest) Validate() error {
	return dara.Validate(s)
}
