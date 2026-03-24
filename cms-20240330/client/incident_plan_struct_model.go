// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentPlanStruct interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRecoverSeconds(v int32) *IncidentPlanStruct
	GetAutoRecoverSeconds() *int32
	SetCloseExpire(v int64) *IncidentPlanStruct
	GetCloseExpire() *int64
	SetCorporation(v []*IncidentPlanCorporationStruct) *IncidentPlanStruct
	GetCorporation() []*IncidentPlanCorporationStruct
	SetDescription(v string) *IncidentPlanStruct
	GetDescription() *string
	SetEscalationId(v []*string) *IncidentPlanStruct
	GetEscalationId() []*string
	SetGmtCreate(v int64) *IncidentPlanStruct
	GetGmtCreate() *int64
	SetGmtModified(v int64) *IncidentPlanStruct
	GetGmtModified() *int64
	SetGroupBy(v []*IncidentPlanFieldPath) *IncidentPlanStruct
	GetGroupBy() []*IncidentPlanFieldPath
	SetIncidentPlanId(v string) *IncidentPlanStruct
	GetIncidentPlanId() *string
	SetName(v string) *IncidentPlanStruct
	GetName() *string
	SetResourceFiled(v []*IncidentPlanFieldPath) *IncidentPlanStruct
	GetResourceFiled() []*IncidentPlanFieldPath
	SetStatus(v string) *IncidentPlanStruct
	GetStatus() *string
	SetUserId(v int64) *IncidentPlanStruct
	GetUserId() *int64
	SetWorkspace(v string) *IncidentPlanStruct
	GetWorkspace() *string
}

type IncidentPlanStruct struct {
	// 自动恢复等待时间。
	//
	// example:
	//
	// 3600
	AutoRecoverSeconds *int32 `json:"autoRecoverSeconds,omitempty" xml:"autoRecoverSeconds,omitempty"`
	// 事件关闭超时时间。
	//
	// example:
	//
	// 86400000
	CloseExpire *int64 `json:"closeExpire,omitempty" xml:"closeExpire,omitempty"`
	// 参与协作的团队或角色列表。
	Corporation []*IncidentPlanCorporationStruct `json:"corporation,omitempty" xml:"corporation,omitempty" type:"Repeated"`
	// 预案描述。
	//
	// example:
	//
	// 针对数据库连接数过高的应急处理方案
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 升级策略 ID的列表。
	EscalationId []*string `json:"escalationId,omitempty" xml:"escalationId,omitempty" type:"Repeated"`
	// 创建时间。
	//
	// example:
	//
	// 1741234567890
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 最后修改时间
	//
	// example:
	//
	// 1741234567890
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 分组字段路径
	GroupBy []*IncidentPlanFieldPath `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
	// 事件预案 ID。
	//
	// example:
	//
	// plan-001
	IncidentPlanId *string `json:"incidentPlanId,omitempty" xml:"incidentPlanId,omitempty"`
	// 预案名称。
	//
	// example:
	//
	// 数据库连接数告警预案
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 资源字段映射路径的列表。
	ResourceFiled []*IncidentPlanFieldPath `json:"resourceFiled,omitempty" xml:"resourceFiled,omitempty" type:"Repeated"`
	// 状态。
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 用户 ID。
	//
	// example:
	//
	// uesr-12345
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
	// 工作空间名称
	//
	// example:
	//
	// workspace-test
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s IncidentPlanStruct) String() string {
	return dara.Prettify(s)
}

func (s IncidentPlanStruct) GoString() string {
	return s.String()
}

func (s *IncidentPlanStruct) GetAutoRecoverSeconds() *int32 {
	return s.AutoRecoverSeconds
}

func (s *IncidentPlanStruct) GetCloseExpire() *int64 {
	return s.CloseExpire
}

func (s *IncidentPlanStruct) GetCorporation() []*IncidentPlanCorporationStruct {
	return s.Corporation
}

func (s *IncidentPlanStruct) GetDescription() *string {
	return s.Description
}

func (s *IncidentPlanStruct) GetEscalationId() []*string {
	return s.EscalationId
}

func (s *IncidentPlanStruct) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *IncidentPlanStruct) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *IncidentPlanStruct) GetGroupBy() []*IncidentPlanFieldPath {
	return s.GroupBy
}

func (s *IncidentPlanStruct) GetIncidentPlanId() *string {
	return s.IncidentPlanId
}

func (s *IncidentPlanStruct) GetName() *string {
	return s.Name
}

func (s *IncidentPlanStruct) GetResourceFiled() []*IncidentPlanFieldPath {
	return s.ResourceFiled
}

func (s *IncidentPlanStruct) GetStatus() *string {
	return s.Status
}

func (s *IncidentPlanStruct) GetUserId() *int64 {
	return s.UserId
}

func (s *IncidentPlanStruct) GetWorkspace() *string {
	return s.Workspace
}

func (s *IncidentPlanStruct) SetAutoRecoverSeconds(v int32) *IncidentPlanStruct {
	s.AutoRecoverSeconds = &v
	return s
}

func (s *IncidentPlanStruct) SetCloseExpire(v int64) *IncidentPlanStruct {
	s.CloseExpire = &v
	return s
}

func (s *IncidentPlanStruct) SetCorporation(v []*IncidentPlanCorporationStruct) *IncidentPlanStruct {
	s.Corporation = v
	return s
}

func (s *IncidentPlanStruct) SetDescription(v string) *IncidentPlanStruct {
	s.Description = &v
	return s
}

func (s *IncidentPlanStruct) SetEscalationId(v []*string) *IncidentPlanStruct {
	s.EscalationId = v
	return s
}

func (s *IncidentPlanStruct) SetGmtCreate(v int64) *IncidentPlanStruct {
	s.GmtCreate = &v
	return s
}

func (s *IncidentPlanStruct) SetGmtModified(v int64) *IncidentPlanStruct {
	s.GmtModified = &v
	return s
}

func (s *IncidentPlanStruct) SetGroupBy(v []*IncidentPlanFieldPath) *IncidentPlanStruct {
	s.GroupBy = v
	return s
}

func (s *IncidentPlanStruct) SetIncidentPlanId(v string) *IncidentPlanStruct {
	s.IncidentPlanId = &v
	return s
}

func (s *IncidentPlanStruct) SetName(v string) *IncidentPlanStruct {
	s.Name = &v
	return s
}

func (s *IncidentPlanStruct) SetResourceFiled(v []*IncidentPlanFieldPath) *IncidentPlanStruct {
	s.ResourceFiled = v
	return s
}

func (s *IncidentPlanStruct) SetStatus(v string) *IncidentPlanStruct {
	s.Status = &v
	return s
}

func (s *IncidentPlanStruct) SetUserId(v int64) *IncidentPlanStruct {
	s.UserId = &v
	return s
}

func (s *IncidentPlanStruct) SetWorkspace(v string) *IncidentPlanStruct {
	s.Workspace = &v
	return s
}

func (s *IncidentPlanStruct) Validate() error {
	if s.Corporation != nil {
		for _, item := range s.Corporation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.GroupBy != nil {
		for _, item := range s.GroupBy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceFiled != nil {
		for _, item := range s.ResourceFiled {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
