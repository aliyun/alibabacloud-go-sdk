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
	// 需修改的实例参数，格式为JSON，修改后的值会覆盖原来的值。例如您只希望修改**maxmemory-policy**参数为**noeviction**，您可以传入`{"maxmemory-policy":"noeviction"}`。
	//
	// > 关于各参数的详细说明，请参见[参数说明](https://help.aliyun.com/document_detail/259681.html)。
	//
	// example:
	//
	// {"maxmemory-policy":"volatile-lru","zset-max-ziplist-entries":128,"zset-max-ziplist-value":64,"hash-max-ziplist-entries":512,"set-max-intset-entries":512}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// 实例ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 哨兵兼容模式，适用于非集群实例。取值说明：
	//
	// 	- **no**（默认）：未开启
	//
	// 	- **yes**：开启
	//
	// > 更多信息请参见[Sentinel兼容模式](https://help.aliyun.com/document_detail/178911.html)。
	//
	// example:
	//
	// yes
	ParamNoLooseSentinelEnabled *string `json:"ParamNoLooseSentinelEnabled,omitempty" xml:"ParamNoLooseSentinelEnabled,omitempty"`
	// 开启哨兵模式时，是否允许免密执行Sentinel相关命令，取值说明：
	//
	// 	- **no**（默认）：关闭。
	//
	// 	- **yes**：开启。开启后，可以在任意连接上免密执行Sentinel命令以及使用SENTINEL命令监听+switch-master通道。
	//
	// example:
	//
	// ****
	ParamNoLooseSentinelPasswordFreeAccess *string `json:"ParamNoLooseSentinelPasswordFreeAccess,omitempty" xml:"ParamNoLooseSentinelPasswordFreeAccess,omitempty"`
	// 启用哨兵模式及ParamNoLooseSentinelPasswordFreeAccess参数后，可通过本参数添加额外的免密命令列表（默认为空）。
	//
	// > 	- 设置后可在任意连接上无需密码执行对应命令，请谨慎操作。
	//
	// > 	- 命令需使用小写字母，多个命令以英文逗号(,)分隔。
	//
	// example:
	//
	// ****
	ParamNoLooseSentinelPasswordFreeCommands *string `json:"ParamNoLooseSentinelPasswordFreeCommands,omitempty" xml:"ParamNoLooseSentinelPasswordFreeCommands,omitempty"`
	// 同步模式：
	//
	// 	- **async**（默认）：异步
	//
	// 	- **semisync**：半同步
	//
	// example:
	//
	// async
	ParamReplMode *string `json:"ParamReplMode,omitempty" xml:"ParamReplMode,omitempty"`
	// 半同步模式的降级阈值。仅半同步支持配置该参数，单位为ms，取值范围为10~60000，默认为500。
	//
	// > 	- 当同步延迟超出该阈值时，同步模式会自动转为异步，当同步延迟消除后，同步模式会自动转换为半同步。
	//
	// > 	- 仅Tair企业版实例支持，该功能公测中。
	//
	// example:
	//
	// 500
	ParamSemisyncReplTimeout *string `json:"ParamSemisyncReplTimeout,omitempty" xml:"ParamSemisyncReplTimeout,omitempty"`
	// 哨兵兼容模式，适用于集群架构代理连接模式或读写分离架构的实例，取值说明：
	//
	// 	- **0**（默认）：未开启
	//
	// 	- **1**：开启
	//
	// > 更多信息请参见[Sentinel兼容模式](https://help.aliyun.com/document_detail/178911.html)。
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
