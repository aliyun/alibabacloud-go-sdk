// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenLogSyncToSLSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *OpenLogSyncToSLSRequest
	GetDBClusterId() *string
	SetLogType(v string) *OpenLogSyncToSLSRequest
	GetLogType() *string
	SetRegionId(v string) *OpenLogSyncToSLSRequest
	GetRegionId() *string
	SetTargetLogStore(v string) *OpenLogSyncToSLSRequest
	GetTargetLogStore() *string
	SetTargetProject(v string) *OpenLogSyncToSLSRequest
	GetTargetProject() *string
}

type OpenLogSyncToSLSRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1ub9grke1****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The log type. Valid values:
	//
	// - **ADBMYSQL_AUDIT_LOG**
	//
	// - **ADBMYSQL_INSERT_LOG**
	//
	// Default value: `ADBMYSQL_AUDIT_LOG`.
	//
	// example:
	//
	// ADBMYSQL_AUDIT_LOG
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Simple Log Service LogStore.
	//
	// This parameter is required.
	//
	// example:
	//
	// adbmysql-audit-log
	TargetLogStore *string `json:"TargetLogStore,omitempty" xml:"TargetLogStore,omitempty"`
	// The Simple Log Service project.
	//
	// This parameter is required.
	//
	// example:
	//
	// log-service-****-cn-shenzhen
	TargetProject *string `json:"TargetProject,omitempty" xml:"TargetProject,omitempty"`
}

func (s OpenLogSyncToSLSRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenLogSyncToSLSRequest) GoString() string {
	return s.String()
}

func (s *OpenLogSyncToSLSRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *OpenLogSyncToSLSRequest) GetLogType() *string {
	return s.LogType
}

func (s *OpenLogSyncToSLSRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenLogSyncToSLSRequest) GetTargetLogStore() *string {
	return s.TargetLogStore
}

func (s *OpenLogSyncToSLSRequest) GetTargetProject() *string {
	return s.TargetProject
}

func (s *OpenLogSyncToSLSRequest) SetDBClusterId(v string) *OpenLogSyncToSLSRequest {
	s.DBClusterId = &v
	return s
}

func (s *OpenLogSyncToSLSRequest) SetLogType(v string) *OpenLogSyncToSLSRequest {
	s.LogType = &v
	return s
}

func (s *OpenLogSyncToSLSRequest) SetRegionId(v string) *OpenLogSyncToSLSRequest {
	s.RegionId = &v
	return s
}

func (s *OpenLogSyncToSLSRequest) SetTargetLogStore(v string) *OpenLogSyncToSLSRequest {
	s.TargetLogStore = &v
	return s
}

func (s *OpenLogSyncToSLSRequest) SetTargetProject(v string) *OpenLogSyncToSLSRequest {
	s.TargetProject = &v
	return s
}

func (s *OpenLogSyncToSLSRequest) Validate() error {
	return dara.Validate(s)
}
