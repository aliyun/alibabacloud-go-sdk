// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSwitchDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *StartSwitchDatabaseRequest
	GetDBInstanceName() *string
	SetDstMainConnectString(v string) *StartSwitchDatabaseRequest
	GetDstMainConnectString() *string
	SetDstMainPort(v string) *StartSwitchDatabaseRequest
	GetDstMainPort() *string
	SetIsModifyEndpoint(v string) *StartSwitchDatabaseRequest
	GetIsModifyEndpoint() *string
	SetRegionId(v string) *StartSwitchDatabaseRequest
	GetRegionId() *string
	SetSlinkTaskId(v string) *StartSwitchDatabaseRequest
	GetSlinkTaskId() *string
	SetSrcMainConnectString(v string) *StartSwitchDatabaseRequest
	GetSrcMainConnectString() *string
	SetSrcMainPort(v string) *StartSwitchDatabaseRequest
	GetSrcMainPort() *string
}

type StartSwitchDatabaseRequest struct {
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// pxc-*********.polarx.rds.aliyuncs.com
	DstMainConnectString *string `json:"DstMainConnectString,omitempty" xml:"DstMainConnectString,omitempty"`
	// example:
	//
	// 3313
	DstMainPort *string `json:"DstMainPort,omitempty" xml:"DstMainPort,omitempty"`
	// example:
	//
	// true
	IsModifyEndpoint *string `json:"IsModifyEndpoint,omitempty" xml:"IsModifyEndpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
	// example:
	//
	// rm-*********.mysql.rds.aliyuncs.com
	SrcMainConnectString *string `json:"SrcMainConnectString,omitempty" xml:"SrcMainConnectString,omitempty"`
	// example:
	//
	// 3308
	SrcMainPort *string `json:"SrcMainPort,omitempty" xml:"SrcMainPort,omitempty"`
}

func (s StartSwitchDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s StartSwitchDatabaseRequest) GoString() string {
	return s.String()
}

func (s *StartSwitchDatabaseRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *StartSwitchDatabaseRequest) GetDstMainConnectString() *string {
	return s.DstMainConnectString
}

func (s *StartSwitchDatabaseRequest) GetDstMainPort() *string {
	return s.DstMainPort
}

func (s *StartSwitchDatabaseRequest) GetIsModifyEndpoint() *string {
	return s.IsModifyEndpoint
}

func (s *StartSwitchDatabaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartSwitchDatabaseRequest) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *StartSwitchDatabaseRequest) GetSrcMainConnectString() *string {
	return s.SrcMainConnectString
}

func (s *StartSwitchDatabaseRequest) GetSrcMainPort() *string {
	return s.SrcMainPort
}

func (s *StartSwitchDatabaseRequest) SetDBInstanceName(v string) *StartSwitchDatabaseRequest {
	s.DBInstanceName = &v
	return s
}

func (s *StartSwitchDatabaseRequest) SetDstMainConnectString(v string) *StartSwitchDatabaseRequest {
	s.DstMainConnectString = &v
	return s
}

func (s *StartSwitchDatabaseRequest) SetDstMainPort(v string) *StartSwitchDatabaseRequest {
	s.DstMainPort = &v
	return s
}

func (s *StartSwitchDatabaseRequest) SetIsModifyEndpoint(v string) *StartSwitchDatabaseRequest {
	s.IsModifyEndpoint = &v
	return s
}

func (s *StartSwitchDatabaseRequest) SetRegionId(v string) *StartSwitchDatabaseRequest {
	s.RegionId = &v
	return s
}

func (s *StartSwitchDatabaseRequest) SetSlinkTaskId(v string) *StartSwitchDatabaseRequest {
	s.SlinkTaskId = &v
	return s
}

func (s *StartSwitchDatabaseRequest) SetSrcMainConnectString(v string) *StartSwitchDatabaseRequest {
	s.SrcMainConnectString = &v
	return s
}

func (s *StartSwitchDatabaseRequest) SetSrcMainPort(v string) *StartSwitchDatabaseRequest {
	s.SrcMainPort = &v
	return s
}

func (s *StartSwitchDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
