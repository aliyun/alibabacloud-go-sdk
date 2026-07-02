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
	// The default configuration parameters of the instance. To view the full list of parameters, call the [DescribeParameters](https://help.aliyun.com/document_detail/473847.html) operation.
	//
	// example:
	//
	// {\\"EvictionPolicy\\":\\"volatile-lru\\",\\"hash-max-ziplist-entries\\":512,\\"zset-max-ziplist-entries\\":128,\\"list-max-ziplist-entries\\":512,\\"list-max-ziplist-value\\":64,\\"zset-max-ziplist-value\\":64,\\"set-max-intset-entries\\":512,\\"hash-max-ziplist-value\\":64}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Specifies whether to enable Sentinel compatibility mode. This parameter applies only to non-cluster instances. Valid values:
	//
	// - **no*	- (default): Disabled
	//
	// - **yes**: Enabled
	//
	// > For more information, see [Sentinel compatibility mode](https://help.aliyun.com/document_detail/178911.html).
	//
	// example:
	//
	// no
	ParamNoLooseSentinelEnabled *string `json:"ParamNoLooseSentinelEnabled,omitempty" xml:"ParamNoLooseSentinelEnabled,omitempty"`
	// Specifies whether to allow password-free execution of Sentinel commands when Sentinel compatibility mode is enabled. Valid values:
	//
	// - **no*	- (default): Disabled.
	//
	// - **yes**: Enabled. Allows you to run Sentinel commands on any connection without a password and use the `SENTINEL` command to subscribe to the `+switch-master` channel.
	//
	// example:
	//
	// no
	ParamNoLooseSentinelPasswordFreeAccess *string `json:"ParamNoLooseSentinelPasswordFreeAccess,omitempty" xml:"ParamNoLooseSentinelPasswordFreeAccess,omitempty"`
	// Additional commands that can be run without a password. This parameter is valid only when Sentinel compatibility mode is enabled and `ParamNoLooseSentinelPasswordFreeAccess` is set to `yes`. By default, this parameter is empty.
	//
	// example:
	//
	// ****
	ParamNoLooseSentinelPasswordFreeCommands *string `json:"ParamNoLooseSentinelPasswordFreeCommands,omitempty" xml:"ParamNoLooseSentinelPasswordFreeCommands,omitempty"`
	// The replication mode. Valid values:
	//
	// - **async*	- (default): asynchronous mode
	//
	// - **semisync**: semi-synchronous mode
	//
	// example:
	//
	// async
	ParamReplMode *string `json:"ParamReplMode,omitempty" xml:"ParamReplMode,omitempty"`
	// The degradation threshold for the semi-synchronous mode. This parameter is valid only in semi-synchronous mode. Unit: milliseconds. Valid values: 10 to 60000. Default value: 500.
	//
	// > If replication latency exceeds this threshold, the replication mode degrades to asynchronous mode. When the replication latency returns to normal, the mode reverts to semi-synchronous mode.
	//
	// example:
	//
	// 500
	ParamReplTimeout *string `json:"ParamReplTimeout,omitempty" xml:"ParamReplTimeout,omitempty"`
	// Specifies whether to enable Sentinel compatibility mode. This parameter applies to instances that use the cluster architecture with proxy connection mode or the read/write splitting architecture. Valid values:
	//
	// - **0*	- (default): Disabled
	//
	// - **1**: Enabled
	//
	// > For more information, see [Sentinel compatibility mode](https://help.aliyun.com/document_detail/178911.html).
	//
	// example:
	//
	// 0
	ParamSentinelCompatEnable *string `json:"ParamSentinelCompatEnable,omitempty" xml:"ParamSentinelCompatEnable,omitempty"`
	// The request ID.
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
