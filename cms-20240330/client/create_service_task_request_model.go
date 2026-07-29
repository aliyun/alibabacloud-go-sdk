// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIp(v string) *CreateServiceTaskRequest
	GetIp() *string
	SetTaskConfig(v string) *CreateServiceTaskRequest
	GetTaskConfig() *string
	SetType(v string) *CreateServiceTaskRequest
	GetType() *string
}

type CreateServiceTaskRequest struct {
	// The IP address of the target instance. This parameter is optional. If not specified, some tasks can match instances by scope (such as instanceIds). This parameter is typically required for heap dump scenarios.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.1
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
	// The task configuration. The value is a JSON string with a maximum length of 65536 characters. This parameter is required for LiveDebug task types. Use a flat JSON structure and pass a single command or probe object directly. Do not wrap it in a commands or probes array. Probe example (dynamic log): {"probeType":"LOG","language":"java","target":{"typeName":"com.example.UserService","methodName":"getUser","location":"exit","instanceIds":["*"]},"action":{"type":"LOG","template":"userId=${args[0]}","templateSegments":[{"type":"TEXT","value":"userId="},{"type":"EXPRESSION","value":"args[0]"]},"ttl":"1h","captureCount":100}. Command example (OGNL): {"commandType":"EVALUATE_EXPRESSION","language":"java","params":{"expression":"@java.lang.System@getProperty(\\"java.home\\")"},"instanceIds":["*"]}. Note: The Command type must include instanceIds at the top level. For Probe types, instanceIds is placed inside the target object. The action.metricType for METRIC probes can be set to COUNTER, GAUGE, HISTOGRAM, or SUMMARY. The Java Agent supports only COUNTER and GAUGE.
	//
	// example:
	//
	// {"probeType":"LOG","language":"java","target":{"typeName":"com.example.service.UserServiceImpl","methodName":"findById","location":"exit","instanceIds":["*"]},"action":{"type":"LOG","template":"userId=${args[0]} cost=${durationMs}ms","templateSegments":[{"type":"TEXT","value":"userId="},{"type":"EXPRESSION","value":"args[0]"},{"type":"TEXT","value":" cost="},{"type":"EXPRESSION","value":"durationMs"},{"type":"TEXT","value":"ms"}]},"ttl":"1h","captureCount":100}
	TaskConfig *string `json:"taskConfig,omitempty" xml:"taskConfig,omitempty"`
	// The task type. This parameter is required. Valid values: heapdump (heap dump). LiveDebug Probe: live_debug_log_probe, live_debug_snapshot_probe, live_debug_metric_probe, live_debug_span_probe, live_debug_span_tag_probe. LiveDebug Command: live_debug_inspect_object, live_debug_search_type, live_debug_search_method, live_debug_decompile, live_debug_get_thread_info, live_debug_get_runtime_info, live_debug_get_memory_info, live_debug_evaluate_expression, live_debug_modify_logger_level. LiveDebug Code Replace: live_debug_code_replace.
	//
	// example:
	//
	// live_debug_log_probe
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateServiceTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceTaskRequest) GetIp() *string {
	return s.Ip
}

func (s *CreateServiceTaskRequest) GetTaskConfig() *string {
	return s.TaskConfig
}

func (s *CreateServiceTaskRequest) GetType() *string {
	return s.Type
}

func (s *CreateServiceTaskRequest) SetIp(v string) *CreateServiceTaskRequest {
	s.Ip = &v
	return s
}

func (s *CreateServiceTaskRequest) SetTaskConfig(v string) *CreateServiceTaskRequest {
	s.TaskConfig = &v
	return s
}

func (s *CreateServiceTaskRequest) SetType(v string) *CreateServiceTaskRequest {
	s.Type = &v
	return s
}

func (s *CreateServiceTaskRequest) Validate() error {
	return dara.Validate(s)
}
