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
	// The instance parameters to modify, in JSON format. The new values overwrite the existing values. For example, if you want to set only the **maxmemory-policy*	- parameter to **noeviction**, pass in `{"maxmemory-policy":"noeviction"}`.
	//
	// > For more information about each parameter, see [Metric description](https://help.aliyun.com/document_detail/259681.html).
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
	// The Sentinel compatibility mode. This parameter applies to non-cluster instances. Valid values:
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
	// 	- **yes**: enabled. After this parameter is enabled, you can run Sentinel commands without a password on any connection and use the SENTINEL command to listen on the +switch-master channel.
	//
	// example:
	//
	// ****
	ParamNoLooseSentinelPasswordFreeAccess *string `json:"ParamNoLooseSentinelPasswordFreeAccess,omitempty" xml:"ParamNoLooseSentinelPasswordFreeAccess,omitempty"`
	// After Sentinel mode and the ParamNoLooseSentinelPasswordFreeAccess parameter are enabled, use this parameter to add additional password-free commands (empty by default).
	//
	// > 	- After this parameter is set, the specified commands can be run without a password on any connection. Proceed with caution.
	//
	// > 	- Commands must be in lowercase letters. Separate multiple commands with commas (,).
	//
	// example:
	//
	// ****
	ParamNoLooseSentinelPasswordFreeCommands *string `json:"ParamNoLooseSentinelPasswordFreeCommands,omitempty" xml:"ParamNoLooseSentinelPasswordFreeCommands,omitempty"`
	// The synchronization pattern. Valid values:
	//
	// 	- **async*	- (default): asynchronous
	//
	// 	- **semisync**: semi-synchronous.
	//
	// example:
	//
	// async
	ParamReplMode *string `json:"ParamReplMode,omitempty" xml:"ParamReplMode,omitempty"`
	// The degradation threshold for semi-synchronous mode. This parameter is supported only in semi-synchronous mode. Unit: ms. Valid values: 10 to 60000. Default value: 500.
	//
	// > 	- When the synchronization latency exceeds this threshold, the synchronous mode automatically transforms to asynchronous. When the latency is eliminated, the synchronous mode automatically transforms back to semi-synchronous.
	//
	// > 	- This parameter is supported only by Tair Enterprise instances. This feature is in public preview.
	//
	// example:
	//
	// 500
	ParamSemisyncReplTimeout *string `json:"ParamSemisyncReplTimeout,omitempty" xml:"ParamSemisyncReplTimeout,omitempty"`
	// The Sentinel compatibility mode. This parameter applies to instances that use the proxy connection mode in cluster architecture or instances that use the read/write splitting architecture. Valid values:
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
