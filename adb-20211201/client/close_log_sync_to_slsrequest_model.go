// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseLogSyncToSLSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CloseLogSyncToSLSRequest
	GetDBClusterId() *string
	SetLogType(v string) *CloseLogSyncToSLSRequest
	GetLogType() *string
	SetRegionId(v string) *CloseLogSyncToSLSRequest
	GetRegionId() *string
}

type CloseLogSyncToSLSRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp198m028ih55****
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

func (s CloseLogSyncToSLSRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseLogSyncToSLSRequest) GoString() string {
	return s.String()
}

func (s *CloseLogSyncToSLSRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CloseLogSyncToSLSRequest) GetLogType() *string {
	return s.LogType
}

func (s *CloseLogSyncToSLSRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CloseLogSyncToSLSRequest) SetDBClusterId(v string) *CloseLogSyncToSLSRequest {
	s.DBClusterId = &v
	return s
}

func (s *CloseLogSyncToSLSRequest) SetLogType(v string) *CloseLogSyncToSLSRequest {
	s.LogType = &v
	return s
}

func (s *CloseLogSyncToSLSRequest) SetRegionId(v string) *CloseLogSyncToSLSRequest {
	s.RegionId = &v
	return s
}

func (s *CloseLogSyncToSLSRequest) Validate() error {
	return dara.Validate(s)
}
