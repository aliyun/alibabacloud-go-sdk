// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogSyncToSLSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetLogSyncToSLSRequest
	GetDBClusterId() *string
	SetLogType(v string) *GetLogSyncToSLSRequest
	GetLogType() *string
	SetRegionId(v string) *GetLogSyncToSLSRequest
	GetRegionId() *string
}

type GetLogSyncToSLSRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf6g8w25jacm7****
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
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetLogSyncToSLSRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLogSyncToSLSRequest) GoString() string {
	return s.String()
}

func (s *GetLogSyncToSLSRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetLogSyncToSLSRequest) GetLogType() *string {
	return s.LogType
}

func (s *GetLogSyncToSLSRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLogSyncToSLSRequest) SetDBClusterId(v string) *GetLogSyncToSLSRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetLogSyncToSLSRequest) SetLogType(v string) *GetLogSyncToSLSRequest {
	s.LogType = &v
	return s
}

func (s *GetLogSyncToSLSRequest) SetRegionId(v string) *GetLogSyncToSLSRequest {
	s.RegionId = &v
	return s
}

func (s *GetLogSyncToSLSRequest) Validate() error {
	return dara.Validate(s)
}
