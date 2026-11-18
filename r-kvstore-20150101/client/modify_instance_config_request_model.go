// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyInstanceConfigRequest
	GetConfig() *string
	SetInstanceId(v string) *ModifyInstanceConfigRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyInstanceConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceConfigRequest
	GetOwnerId() *int64
	SetParamNoLooseSentinelEnabled(v string) *ModifyInstanceConfigRequest
	GetParamNoLooseSentinelEnabled() *string
	SetParamNoLooseSentinelPasswordFreeAccess(v string) *ModifyInstanceConfigRequest
	GetParamNoLooseSentinelPasswordFreeAccess() *string
	SetParamNoLooseSentinelPasswordFreeCommands(v string) *ModifyInstanceConfigRequest
	GetParamNoLooseSentinelPasswordFreeCommands() *string
	SetParamReplMode(v string) *ModifyInstanceConfigRequest
	GetParamReplMode() *string
	SetParamSemisyncReplTimeout(v string) *ModifyInstanceConfigRequest
	GetParamSemisyncReplTimeout() *string
	SetParamSentinelCompatEnable(v string) *ModifyInstanceConfigRequest
	GetParamSentinelCompatEnable() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceConfigRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyInstanceConfigRequest
	GetSecurityToken() *string
}

type ModifyInstanceConfigRequest struct {
	// The instance parameters to be modified in JSON format. The modified values overwrite the original values. For example, to modify only the **maxmemory-policy*	- parameter to **noeviction**, you can pass `{"maxmemory-policy":"noeviction"}`.
	//
	// > For more information about each parameter, see [Parameter description](https://help.aliyun.com/document_detail/259681.html).
	//
	// example:
	//
	// {"maxmemory-policy":"volatile-lru","zset-max-ziplist-entries":128,"zset-max-ziplist-value":64,"hash-max-ziplist-entries":512,"set-max-intset-entries":512}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to enable Sentinel compatibility mode for non-cluster instances. Valid values:
	//
	// 	- **no*	- (default): disabled.
	//
	// 	- **yes**: enabled.
	//
	// > For more information, see [Sentinel compatibility mode](https://help.aliyun.com/document_detail/178911.html).
	//
	// example:
	//
	// yes
	ParamNoLooseSentinelEnabled *string `json:"ParamNoLooseSentinelEnabled,omitempty" xml:"ParamNoLooseSentinelEnabled,omitempty"`
	// Specifies whether to allow password-free execution of Sentinel-related commands when Sentinel mode is enabled. Valid values:
	//
	// 	- **no*	- (default): disabled.
	//
	// 	- **yes**: enabled. After this feature is enabled, you can execute Sentinel commands without a password on any connection and use the SENTINEL command to monitor the +switch-master channel.
	//
	// example:
	//
	// no
	ParamNoLooseSentinelPasswordFreeAccess *string `json:"ParamNoLooseSentinelPasswordFreeAccess,omitempty" xml:"ParamNoLooseSentinelPasswordFreeAccess,omitempty"`
	// The list of additional password-free commands. This parameter takes effect only after Sentinel mode and the ParamNoLooseSentinelPasswordFreeAccess parameter are enabled. Default value: empty.
	//
	// > 	- After this parameter is set, the specified commands can be executed without a password on any connection. Proceed with caution.
	//
	// > 	- Commands must be in lowercase letters. Separate multiple commands with commas (,).
	//
	// example:
	//
	// ****
	ParamNoLooseSentinelPasswordFreeCommands *string `json:"ParamNoLooseSentinelPasswordFreeCommands,omitempty" xml:"ParamNoLooseSentinelPasswordFreeCommands,omitempty"`
	// The synchronization mode. Valid values:
	//
	// 	- **async*	- (default): asynchronous.
	//
	// 	- **semisync**: semi-synchronous.
	//
	// example:
	//
	// async
	ParamReplMode *string `json:"ParamReplMode,omitempty" xml:"ParamReplMode,omitempty"`
	// The degradation threshold for the semi-synchronous mode. This parameter is supported only in the semi-synchronous mode. Unit: milliseconds. Valid values: 10 to 60000. Default value: 500.
	//
	// > 	- When the synchronization latency exceeds this threshold, the synchronization mode automatically switches to asynchronous. When the synchronization latency is eliminated, the synchronization mode automatically switches back to semi-synchronous.
	//
	// > 	- This feature is supported only by Tair Enterprise instances and is in public preview.
	//
	// example:
	//
	// 500
	ParamSemisyncReplTimeout *string `json:"ParamSemisyncReplTimeout,omitempty" xml:"ParamSemisyncReplTimeout,omitempty"`
	// Specifies whether to enable Sentinel compatibility mode for instances in cluster architecture proxy connection mode or read/write splitting architecture. Valid values:
	//
	// 	- **0*	- (default): disabled.
	//
	// 	- **1**: enabled.
	//
	// > For more information, see [Sentinel compatibility mode](https://help.aliyun.com/document_detail/178911.html).
	//
	// example:
	//
	// 1
	ParamSentinelCompatEnable *string `json:"ParamSentinelCompatEnable,omitempty" xml:"ParamSentinelCompatEnable,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken             *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyInstanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyInstanceConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceConfigRequest) GetParamNoLooseSentinelEnabled() *string {
	return s.ParamNoLooseSentinelEnabled
}

func (s *ModifyInstanceConfigRequest) GetParamNoLooseSentinelPasswordFreeAccess() *string {
	return s.ParamNoLooseSentinelPasswordFreeAccess
}

func (s *ModifyInstanceConfigRequest) GetParamNoLooseSentinelPasswordFreeCommands() *string {
	return s.ParamNoLooseSentinelPasswordFreeCommands
}

func (s *ModifyInstanceConfigRequest) GetParamReplMode() *string {
	return s.ParamReplMode
}

func (s *ModifyInstanceConfigRequest) GetParamSemisyncReplTimeout() *string {
	return s.ParamSemisyncReplTimeout
}

func (s *ModifyInstanceConfigRequest) GetParamSentinelCompatEnable() *string {
	return s.ParamSentinelCompatEnable
}

func (s *ModifyInstanceConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyInstanceConfigRequest) SetConfig(v string) *ModifyInstanceConfigRequest {
	s.Config = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetInstanceId(v string) *ModifyInstanceConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetOwnerAccount(v string) *ModifyInstanceConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetOwnerId(v int64) *ModifyInstanceConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetParamNoLooseSentinelEnabled(v string) *ModifyInstanceConfigRequest {
	s.ParamNoLooseSentinelEnabled = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetParamNoLooseSentinelPasswordFreeAccess(v string) *ModifyInstanceConfigRequest {
	s.ParamNoLooseSentinelPasswordFreeAccess = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetParamNoLooseSentinelPasswordFreeCommands(v string) *ModifyInstanceConfigRequest {
	s.ParamNoLooseSentinelPasswordFreeCommands = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetParamReplMode(v string) *ModifyInstanceConfigRequest {
	s.ParamReplMode = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetParamSemisyncReplTimeout(v string) *ModifyInstanceConfigRequest {
	s.ParamSemisyncReplTimeout = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetParamSentinelCompatEnable(v string) *ModifyInstanceConfigRequest {
	s.ParamSentinelCompatEnable = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetResourceOwnerAccount(v string) *ModifyInstanceConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetResourceOwnerId(v int64) *ModifyInstanceConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetSecurityToken(v string) *ModifyInstanceConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
