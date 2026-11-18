// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *DescribeInstanceConfigResponseBody
	GetConfig() *string
	SetParamNoLooseSentinelEnabled(v string) *DescribeInstanceConfigResponseBody
	GetParamNoLooseSentinelEnabled() *string
	SetParamNoLooseSentinelPasswordFreeAccess(v string) *DescribeInstanceConfigResponseBody
	GetParamNoLooseSentinelPasswordFreeAccess() *string
	SetParamNoLooseSentinelPasswordFreeCommands(v string) *DescribeInstanceConfigResponseBody
	GetParamNoLooseSentinelPasswordFreeCommands() *string
	SetParamReplMode(v string) *DescribeInstanceConfigResponseBody
	GetParamReplMode() *string
	SetParamReplTimeout(v string) *DescribeInstanceConfigResponseBody
	GetParamReplTimeout() *string
	SetParamSentinelCompatEnable(v string) *DescribeInstanceConfigResponseBody
	GetParamSentinelCompatEnable() *string
	SetRequestId(v string) *DescribeInstanceConfigResponseBody
	GetRequestId() *string
}

type DescribeInstanceConfigResponseBody struct {
	// The default configuration parameters of the instance. For more comprehensive parameter configurations, call the [Describeparamters](https://help.aliyun.com/document_detail/473847.html) operation.
	//
	// example:
	//
	// {\\"EvictionPolicy\\":\\"volatile-lru\\",\\"hash-max-ziplist-entries\\":512,\\"zset-max-ziplist-entries\\":128,\\"list-max-ziplist-entries\\":512,\\"list-max-ziplist-value\\":64,\\"zset-max-ziplist-value\\":64,\\"set-max-intset-entries\\":512,\\"hash-max-ziplist-value\\":64}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The Sentinel compatibility mode, which applies to non-cluster instances. Valid values:
	//
	// 	- **no**: disabled. This is the default value.
	//
	// 	- **yes**: enabled.
	//
	// > For more information, see [Sentinel compatibility mode](https://help.aliyun.com/document_detail/178911.html).
	//
	// example:
	//
	// no
	ParamNoLooseSentinelEnabled *string `json:"ParamNoLooseSentinelEnabled,omitempty" xml:"ParamNoLooseSentinelEnabled,omitempty"`
	// Specifies whether to allow password-free execution of Sentinel-related commands when the Sentinel mode is enabled. Valid values:
	//
	// 	- **no**: disabled. This is the default value.
	//
	// 	- **yes**: enabled. After this feature is enabled, you can execute Sentinel commands without a password on any connection and use the SENTINEL command to monitor the +switch-master channel.
	//
	// example:
	//
	// no
	ParamNoLooseSentinelPasswordFreeAccess *string `json:"ParamNoLooseSentinelPasswordFreeAccess,omitempty" xml:"ParamNoLooseSentinelPasswordFreeAccess,omitempty"`
	// The list of additional password-free commands. After the Sentinel mode and the ParamNoLooseSentinelPasswordFreeAccess parameter are enabled, you can use this parameter to add extra password-free commands. Default value: empty.
	//
	// example:
	//
	// ****
	ParamNoLooseSentinelPasswordFreeCommands *string `json:"ParamNoLooseSentinelPasswordFreeCommands,omitempty" xml:"ParamNoLooseSentinelPasswordFreeCommands,omitempty"`
	// The synchronization mode. Valid values:
	//
	// 	- **async**: asynchronous mode. This is the default value.
	//
	// 	- **semisync**: semi-synchronous mode.
	//
	// example:
	//
	// async
	ParamReplMode *string `json:"ParamReplMode,omitempty" xml:"ParamReplMode,omitempty"`
	// The degradation threshold for the semi-synchronous mode. This parameter can be configured only when the semi-synchronous mode is used. Unit: milliseconds. Valid values: 10 to 60000. Default value: 500.
	//
	// > When the synchronization latency exceeds this threshold, the synchronization mode automatically degrades to the asynchronous mode. After the synchronization latency is eliminated, the synchronization mode automatically switches back to the semi-synchronous mode.
	//
	// example:
	//
	// 500
	ParamReplTimeout *string `json:"ParamReplTimeout,omitempty" xml:"ParamReplTimeout,omitempty"`
	// The Sentinel compatibility mode, which applies to instances in the cluster architecture proxy mode or read/write splitting architecture. Valid values:
	//
	// 	- **0**: disabled. This is the default value.
	//
	// 	- **1**: enabled.
	//
	// > For more information, see [Sentinel compatibility mode](https://help.aliyun.com/document_detail/178911.html).
	//
	// example:
	//
	// 0
	ParamSentinelCompatEnable *string `json:"ParamSentinelCompatEnable,omitempty" xml:"ParamSentinelCompatEnable,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4E2C08F6-2D11-4ECD-9A4C-27EF2D3D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceConfigResponseBody) GetConfig() *string {
	return s.Config
}

func (s *DescribeInstanceConfigResponseBody) GetParamNoLooseSentinelEnabled() *string {
	return s.ParamNoLooseSentinelEnabled
}

func (s *DescribeInstanceConfigResponseBody) GetParamNoLooseSentinelPasswordFreeAccess() *string {
	return s.ParamNoLooseSentinelPasswordFreeAccess
}

func (s *DescribeInstanceConfigResponseBody) GetParamNoLooseSentinelPasswordFreeCommands() *string {
	return s.ParamNoLooseSentinelPasswordFreeCommands
}

func (s *DescribeInstanceConfigResponseBody) GetParamReplMode() *string {
	return s.ParamReplMode
}

func (s *DescribeInstanceConfigResponseBody) GetParamReplTimeout() *string {
	return s.ParamReplTimeout
}

func (s *DescribeInstanceConfigResponseBody) GetParamSentinelCompatEnable() *string {
	return s.ParamSentinelCompatEnable
}

func (s *DescribeInstanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceConfigResponseBody) SetConfig(v string) *DescribeInstanceConfigResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeInstanceConfigResponseBody) SetParamNoLooseSentinelEnabled(v string) *DescribeInstanceConfigResponseBody {
	s.ParamNoLooseSentinelEnabled = &v
	return s
}

func (s *DescribeInstanceConfigResponseBody) SetParamNoLooseSentinelPasswordFreeAccess(v string) *DescribeInstanceConfigResponseBody {
	s.ParamNoLooseSentinelPasswordFreeAccess = &v
	return s
}

func (s *DescribeInstanceConfigResponseBody) SetParamNoLooseSentinelPasswordFreeCommands(v string) *DescribeInstanceConfigResponseBody {
	s.ParamNoLooseSentinelPasswordFreeCommands = &v
	return s
}

func (s *DescribeInstanceConfigResponseBody) SetParamReplMode(v string) *DescribeInstanceConfigResponseBody {
	s.ParamReplMode = &v
	return s
}

func (s *DescribeInstanceConfigResponseBody) SetParamReplTimeout(v string) *DescribeInstanceConfigResponseBody {
	s.ParamReplTimeout = &v
	return s
}

func (s *DescribeInstanceConfigResponseBody) SetParamSentinelCompatEnable(v string) *DescribeInstanceConfigResponseBody {
	s.ParamSentinelCompatEnable = &v
	return s
}

func (s *DescribeInstanceConfigResponseBody) SetRequestId(v string) *DescribeInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
