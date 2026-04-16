// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudImportTaskTelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBridgeVoicePath(v string) *CloudImportTaskTelRequest
	GetBridgeVoicePath() *string
	SetBridgeVoiceType(v int64) *CloudImportTaskTelRequest
	GetBridgeVoiceType() *int64
	SetEnterpriseId(v int64) *CloudImportTaskTelRequest
	GetEnterpriseId() *int64
	SetFileId(v int64) *CloudImportTaskTelRequest
	GetFileId() *int64
	SetImportTelAutoStart(v int64) *CloudImportTaskTelRequest
	GetImportTelAutoStart() *int64
	SetIsRepeat(v int64) *CloudImportTaskTelRequest
	GetIsRepeat() *int64
	SetName(v string) *CloudImportTaskTelRequest
	GetName() *string
	SetOwnerId(v int64) *CloudImportTaskTelRequest
	GetOwnerId() *int64
	SetPriority(v int64) *CloudImportTaskTelRequest
	GetPriority() *int64
	SetResourceOwnerAccount(v string) *CloudImportTaskTelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudImportTaskTelRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *CloudImportTaskTelRequest
	GetTaskId() *int64
	SetTaskTelList(v []*CloudImportTaskTelRequestTaskTelList) *CloudImportTaskTelRequest
	GetTaskTelList() []*CloudImportTaskTelRequestTaskTelList
}

type CloudImportTaskTelRequest struct {
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
	TaskTelList []*CloudImportTaskTelRequestTaskTelList `json:"TaskTelList,omitempty" xml:"TaskTelList,omitempty" type:"Repeated"`
}

func (s CloudImportTaskTelRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudImportTaskTelRequest) GoString() string {
	return s.String()
}

func (s *CloudImportTaskTelRequest) GetBridgeVoicePath() *string {
	return s.BridgeVoicePath
}

func (s *CloudImportTaskTelRequest) GetBridgeVoiceType() *int64 {
	return s.BridgeVoiceType
}

func (s *CloudImportTaskTelRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudImportTaskTelRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *CloudImportTaskTelRequest) GetImportTelAutoStart() *int64 {
	return s.ImportTelAutoStart
}

func (s *CloudImportTaskTelRequest) GetIsRepeat() *int64 {
	return s.IsRepeat
}

func (s *CloudImportTaskTelRequest) GetName() *string {
	return s.Name
}

func (s *CloudImportTaskTelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudImportTaskTelRequest) GetPriority() *int64 {
	return s.Priority
}

func (s *CloudImportTaskTelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudImportTaskTelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudImportTaskTelRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CloudImportTaskTelRequest) GetTaskTelList() []*CloudImportTaskTelRequestTaskTelList {
	return s.TaskTelList
}

func (s *CloudImportTaskTelRequest) SetBridgeVoicePath(v string) *CloudImportTaskTelRequest {
	s.BridgeVoicePath = &v
	return s
}

func (s *CloudImportTaskTelRequest) SetBridgeVoiceType(v int64) *CloudImportTaskTelRequest {
	s.BridgeVoiceType = &v
	return s
}

func (s *CloudImportTaskTelRequest) SetEnterpriseId(v int64) *CloudImportTaskTelRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudImportTaskTelRequest) SetFileId(v int64) *CloudImportTaskTelRequest {
	s.FileId = &v
	return s
}

func (s *CloudImportTaskTelRequest) SetImportTelAutoStart(v int64) *CloudImportTaskTelRequest {
	s.ImportTelAutoStart = &v
	return s
}

func (s *CloudImportTaskTelRequest) SetIsRepeat(v int64) *CloudImportTaskTelRequest {
	s.IsRepeat = &v
	return s
}

func (s *CloudImportTaskTelRequest) SetName(v string) *CloudImportTaskTelRequest {
	s.Name = &v
	return s
}

func (s *CloudImportTaskTelRequest) SetOwnerId(v int64) *CloudImportTaskTelRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudImportTaskTelRequest) SetPriority(v int64) *CloudImportTaskTelRequest {
	s.Priority = &v
	return s
}

func (s *CloudImportTaskTelRequest) SetResourceOwnerAccount(v string) *CloudImportTaskTelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudImportTaskTelRequest) SetResourceOwnerId(v int64) *CloudImportTaskTelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudImportTaskTelRequest) SetTaskId(v int64) *CloudImportTaskTelRequest {
	s.TaskId = &v
	return s
}

func (s *CloudImportTaskTelRequest) SetTaskTelList(v []*CloudImportTaskTelRequestTaskTelList) *CloudImportTaskTelRequest {
	s.TaskTelList = v
	return s
}

func (s *CloudImportTaskTelRequest) Validate() error {
	if s.TaskTelList != nil {
		for _, item := range s.TaskTelList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudImportTaskTelRequestTaskTelList struct {
	// 备选号码，tel呼叫不通时，呼叫备选号码最多支持8个，号码之间用英文逗号","分隔
	//
	// example:
	//
	// 7455332
	BackupTels *string `json:"BackupTels,omitempty" xml:"BackupTels,omitempty"`
	// 电话号对应的外显号码
	//
	// example:
	//
	// 7766551
	Clid *string `json:"Clid,omitempty" xml:"Clid,omitempty"`
	// 使用clidGroup需要账号支持按标识路由，使用此参数是clid参数无效
	//
	// example:
	//
	// 示例值
	ClidGroup *string `json:"ClidGroup,omitempty" xml:"ClidGroup,omitempty"`
	// 优先级，默认为0，值越大优先级越高，最大999999
	//
	// example:
	//
	// 38
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// 属性，json格式
	//
	// example:
	//
	// {}
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// 电话号
	//
	// example:
	//
	// 7455441
	Tel *string `json:"Tel,omitempty" xml:"Tel,omitempty"`
}

func (s CloudImportTaskTelRequestTaskTelList) String() string {
	return dara.Prettify(s)
}

func (s CloudImportTaskTelRequestTaskTelList) GoString() string {
	return s.String()
}

func (s *CloudImportTaskTelRequestTaskTelList) GetBackupTels() *string {
	return s.BackupTels
}

func (s *CloudImportTaskTelRequestTaskTelList) GetClid() *string {
	return s.Clid
}

func (s *CloudImportTaskTelRequestTaskTelList) GetClidGroup() *string {
	return s.ClidGroup
}

func (s *CloudImportTaskTelRequestTaskTelList) GetPriority() *int64 {
	return s.Priority
}

func (s *CloudImportTaskTelRequestTaskTelList) GetProperty() *string {
	return s.Property
}

func (s *CloudImportTaskTelRequestTaskTelList) GetTel() *string {
	return s.Tel
}

func (s *CloudImportTaskTelRequestTaskTelList) SetBackupTels(v string) *CloudImportTaskTelRequestTaskTelList {
	s.BackupTels = &v
	return s
}

func (s *CloudImportTaskTelRequestTaskTelList) SetClid(v string) *CloudImportTaskTelRequestTaskTelList {
	s.Clid = &v
	return s
}

func (s *CloudImportTaskTelRequestTaskTelList) SetClidGroup(v string) *CloudImportTaskTelRequestTaskTelList {
	s.ClidGroup = &v
	return s
}

func (s *CloudImportTaskTelRequestTaskTelList) SetPriority(v int64) *CloudImportTaskTelRequestTaskTelList {
	s.Priority = &v
	return s
}

func (s *CloudImportTaskTelRequestTaskTelList) SetProperty(v string) *CloudImportTaskTelRequestTaskTelList {
	s.Property = &v
	return s
}

func (s *CloudImportTaskTelRequestTaskTelList) SetTel(v string) *CloudImportTaskTelRequestTaskTelList {
	s.Tel = &v
	return s
}

func (s *CloudImportTaskTelRequestTaskTelList) Validate() error {
	return dara.Validate(s)
}
