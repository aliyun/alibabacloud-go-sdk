// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudImportTaskTelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBridgeVoicePath(v string) *CloudImportTaskTelShrinkRequest
	GetBridgeVoicePath() *string
	SetBridgeVoiceType(v int64) *CloudImportTaskTelShrinkRequest
	GetBridgeVoiceType() *int64
	SetEnterpriseId(v int64) *CloudImportTaskTelShrinkRequest
	GetEnterpriseId() *int64
	SetFileId(v int64) *CloudImportTaskTelShrinkRequest
	GetFileId() *int64
	SetImportTelAutoStart(v int64) *CloudImportTaskTelShrinkRequest
	GetImportTelAutoStart() *int64
	SetIsRepeat(v int64) *CloudImportTaskTelShrinkRequest
	GetIsRepeat() *int64
	SetName(v string) *CloudImportTaskTelShrinkRequest
	GetName() *string
	SetOwnerId(v int64) *CloudImportTaskTelShrinkRequest
	GetOwnerId() *int64
	SetPriority(v int64) *CloudImportTaskTelShrinkRequest
	GetPriority() *int64
	SetResourceOwnerAccount(v string) *CloudImportTaskTelShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudImportTaskTelShrinkRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *CloudImportTaskTelShrinkRequest
	GetTaskId() *int64
	SetTaskTelListShrink(v string) *CloudImportTaskTelShrinkRequest
	GetTaskTelListShrink() *string
}

type CloudImportTaskTelShrinkRequest struct {
	// 座席接听时自动在双侧播放开场白语音，指定语音变量值；企业语音库里的语音变量值
	//
	// example:
	//
	// 示例值示例值
	BridgeVoicePath *string `json:"BridgeVoicePath,omitempty" xml:"BridgeVoicePath,omitempty"`
	// 座席接听时自动在双侧播放开场白语音类型；1. 公共语音库；2. 企业语音库，静态话术； 3. 企业语音库，动态话术（座席号），传bridgeVoicePath后生效，默认为3
	//
	// example:
	//
	// 3
	BridgeVoiceType *int64 `json:"BridgeVoiceType,omitempty" xml:"BridgeVoiceType,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 17
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 批次Id；传此值表示在批次中增加号码
	//
	// example:
	//
	// 666
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// 是否自动启动任务；0:不自动启动 1:自动启动，默认为0
	//
	// example:
	//
	// 0
	ImportTelAutoStart *int64 `json:"ImportTelAutoStart,omitempty" xml:"ImportTelAutoStart,omitempty"`
	// 是否排重；0.不排重 1.任务内排重 2.导入号码排重 3.批次内排重，默认为1。注：任务内排重与批次内排重不能同时支持，如果中途切换，则从本次切换开始进行排重。
	//
	// example:
	//
	// 1
	IsRepeat *int64 `json:"IsRepeat,omitempty" xml:"IsRepeat,omitempty"`
	// 批次名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 优先级；默认0，值越大越优先，最大999999
	//
	// example:
	//
	// 5
	Priority             *int64  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 任务Id
	//
	// This parameter is required.
	//
	// example:
	//
	// 25
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 任务号码列表；CtiLinkTaskTel中只有Tel是必选字段，长度大小不超过8MB 注：获取导入失败明细，需配置[事件推送](../字段定义/推送变量和值/预测外呼导入号码失败推送变量.md)
	//
	// This parameter is required.
	TaskTelListShrink *string `json:"TaskTelList,omitempty" xml:"TaskTelList,omitempty"`
}

func (s CloudImportTaskTelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudImportTaskTelShrinkRequest) GoString() string {
	return s.String()
}

func (s *CloudImportTaskTelShrinkRequest) GetBridgeVoicePath() *string {
	return s.BridgeVoicePath
}

func (s *CloudImportTaskTelShrinkRequest) GetBridgeVoiceType() *int64 {
	return s.BridgeVoiceType
}

func (s *CloudImportTaskTelShrinkRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudImportTaskTelShrinkRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *CloudImportTaskTelShrinkRequest) GetImportTelAutoStart() *int64 {
	return s.ImportTelAutoStart
}

func (s *CloudImportTaskTelShrinkRequest) GetIsRepeat() *int64 {
	return s.IsRepeat
}

func (s *CloudImportTaskTelShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CloudImportTaskTelShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudImportTaskTelShrinkRequest) GetPriority() *int64 {
	return s.Priority
}

func (s *CloudImportTaskTelShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudImportTaskTelShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudImportTaskTelShrinkRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CloudImportTaskTelShrinkRequest) GetTaskTelListShrink() *string {
	return s.TaskTelListShrink
}

func (s *CloudImportTaskTelShrinkRequest) SetBridgeVoicePath(v string) *CloudImportTaskTelShrinkRequest {
	s.BridgeVoicePath = &v
	return s
}

func (s *CloudImportTaskTelShrinkRequest) SetBridgeVoiceType(v int64) *CloudImportTaskTelShrinkRequest {
	s.BridgeVoiceType = &v
	return s
}

func (s *CloudImportTaskTelShrinkRequest) SetEnterpriseId(v int64) *CloudImportTaskTelShrinkRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudImportTaskTelShrinkRequest) SetFileId(v int64) *CloudImportTaskTelShrinkRequest {
	s.FileId = &v
	return s
}

func (s *CloudImportTaskTelShrinkRequest) SetImportTelAutoStart(v int64) *CloudImportTaskTelShrinkRequest {
	s.ImportTelAutoStart = &v
	return s
}

func (s *CloudImportTaskTelShrinkRequest) SetIsRepeat(v int64) *CloudImportTaskTelShrinkRequest {
	s.IsRepeat = &v
	return s
}

func (s *CloudImportTaskTelShrinkRequest) SetName(v string) *CloudImportTaskTelShrinkRequest {
	s.Name = &v
	return s
}

func (s *CloudImportTaskTelShrinkRequest) SetOwnerId(v int64) *CloudImportTaskTelShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudImportTaskTelShrinkRequest) SetPriority(v int64) *CloudImportTaskTelShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CloudImportTaskTelShrinkRequest) SetResourceOwnerAccount(v string) *CloudImportTaskTelShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudImportTaskTelShrinkRequest) SetResourceOwnerId(v int64) *CloudImportTaskTelShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudImportTaskTelShrinkRequest) SetTaskId(v int64) *CloudImportTaskTelShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *CloudImportTaskTelShrinkRequest) SetTaskTelListShrink(v string) *CloudImportTaskTelShrinkRequest {
	s.TaskTelListShrink = &v
	return s
}

func (s *CloudImportTaskTelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
