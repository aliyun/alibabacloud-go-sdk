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
	// The parameter settings of the instance. For more information, see [Parameter overview and configuration guide](https://help.aliyun.com/document_detail/43885.html).
	//
	// example:
	//
	// {\\"EvictionPolicy\\":\\"volatile-lru\\",\\"hash-max-ziplist-entries\\":512,\\"zset-max-ziplist-entries\\":128,\\"list-max-ziplist-entries\\":512,\\"list-max-ziplist-value\\":64,\\"zset-max-ziplist-value\\":64,\\"set-max-intset-entries\\":512,\\"hash-max-ziplist-value\\":64}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The Sentinel-compatible mode, which is applicable to non-cluster instances. For more information about the parameter, see the relevant documentation.
	//
	// example:
	//
	// ****
	ParamNoLooseSentinelEnabled *string `json:"ParamNoLooseSentinelEnabled,omitempty" xml:"ParamNoLooseSentinelEnabled,omitempty"`
	// Indicates whether Sentinel commands can be run without requiring a password when the Sentinel mode is enabled. Valid values: Valid values: yes and no. Default value: no. After you set this parameter to yes, you can run Sentinel commands in a virtual private cloud (VPC) without the need to enable the password-free access feature.
	//
	// example:
	//
	// no
	ParamNoLooseSentinelPasswordFreeAccess *string `json:"ParamNoLooseSentinelPasswordFreeAccess,omitempty" xml:"ParamNoLooseSentinelPasswordFreeAccess,omitempty"`
	// After you enable the Sentinel mode and set the ParamNoLooseSentinelPasswordFreeAccess parameter to yes, you can use this parameter to specify an additional list of commands that can be run without requiring a password. By default, this parameter is empty. After you configure this parameter, you can run the specified commands without a password on any connection. Proceed with caution. The commands must be written in lowercase letters. Multiple commands are separated by commas (,).
	//
	// example:
	//
	// ****
	ParamNoLooseSentinelPasswordFreeCommands *string `json:"ParamNoLooseSentinelPasswordFreeCommands,omitempty" xml:"ParamNoLooseSentinelPasswordFreeCommands,omitempty"`
	// The synchronization mode.
	//
	// 	- **semisync**
	//
	// 	- **async**
	//
	// example:
	//
	// async
	ParamReplMode *string `json:"ParamReplMode,omitempty" xml:"ParamReplMode,omitempty"`
	// The degradation threshold time of the semi-synchronous replication mode. This parameter is required only when semi-synchronous replication is enabled. Unit: milliseconds. Valid values: 10 to 60000.
	//
	// example:
	//
	// ****
	ParamReplTimeout *string `json:"ParamReplTimeout,omitempty" xml:"ParamReplTimeout,omitempty"`
	// The Sentinel-compatible mode, which is applicable to cluster instances in proxy mode or read/write splitting instances. For more information about the parameter, see the relevant documentation.
	//
	// example:
	//
	// 1
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
