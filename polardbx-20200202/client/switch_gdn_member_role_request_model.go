// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchGdnMemberRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *SwitchGdnMemberRoleRequest
	GetDBInstanceName() *string
	SetDstMainConnectString(v string) *SwitchGdnMemberRoleRequest
	GetDstMainConnectString() *string
	SetDstMainPort(v string) *SwitchGdnMemberRoleRequest
	GetDstMainPort() *string
	SetIsModifyEndpoint(v string) *SwitchGdnMemberRoleRequest
	GetIsModifyEndpoint() *string
	SetRegionId(v string) *SwitchGdnMemberRoleRequest
	GetRegionId() *string
	SetSrcMainConnectString(v string) *SwitchGdnMemberRoleRequest
	GetSrcMainConnectString() *string
	SetSrcMainPort(v string) *SwitchGdnMemberRoleRequest
	GetSrcMainPort() *string
	SetSwitchMode(v string) *SwitchGdnMemberRoleRequest
	GetSwitchMode() *string
	SetTaskTimeout(v int64) *SwitchGdnMemberRoleRequest
	GetTaskTimeout() *int64
}

type SwitchGdnMemberRoleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName       *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DstMainConnectString *string `json:"DstMainConnectString,omitempty" xml:"DstMainConnectString,omitempty"`
	DstMainPort          *string `json:"DstMainPort,omitempty" xml:"DstMainPort,omitempty"`
	IsModifyEndpoint     *string `json:"IsModifyEndpoint,omitempty" xml:"IsModifyEndpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SrcMainConnectString *string `json:"SrcMainConnectString,omitempty" xml:"SrcMainConnectString,omitempty"`
	SrcMainPort          *string `json:"SrcMainPort,omitempty" xml:"SrcMainPort,omitempty"`
	// This parameter is required.
	SwitchMode  *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
	TaskTimeout *int64  `json:"TaskTimeout,omitempty" xml:"TaskTimeout,omitempty"`
}

func (s SwitchGdnMemberRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchGdnMemberRoleRequest) GoString() string {
	return s.String()
}

func (s *SwitchGdnMemberRoleRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *SwitchGdnMemberRoleRequest) GetDstMainConnectString() *string {
	return s.DstMainConnectString
}

func (s *SwitchGdnMemberRoleRequest) GetDstMainPort() *string {
	return s.DstMainPort
}

func (s *SwitchGdnMemberRoleRequest) GetIsModifyEndpoint() *string {
	return s.IsModifyEndpoint
}

func (s *SwitchGdnMemberRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SwitchGdnMemberRoleRequest) GetSrcMainConnectString() *string {
	return s.SrcMainConnectString
}

func (s *SwitchGdnMemberRoleRequest) GetSrcMainPort() *string {
	return s.SrcMainPort
}

func (s *SwitchGdnMemberRoleRequest) GetSwitchMode() *string {
	return s.SwitchMode
}

func (s *SwitchGdnMemberRoleRequest) GetTaskTimeout() *int64 {
	return s.TaskTimeout
}

func (s *SwitchGdnMemberRoleRequest) SetDBInstanceName(v string) *SwitchGdnMemberRoleRequest {
	s.DBInstanceName = &v
	return s
}

func (s *SwitchGdnMemberRoleRequest) SetDstMainConnectString(v string) *SwitchGdnMemberRoleRequest {
	s.DstMainConnectString = &v
	return s
}

func (s *SwitchGdnMemberRoleRequest) SetDstMainPort(v string) *SwitchGdnMemberRoleRequest {
	s.DstMainPort = &v
	return s
}

func (s *SwitchGdnMemberRoleRequest) SetIsModifyEndpoint(v string) *SwitchGdnMemberRoleRequest {
	s.IsModifyEndpoint = &v
	return s
}

func (s *SwitchGdnMemberRoleRequest) SetRegionId(v string) *SwitchGdnMemberRoleRequest {
	s.RegionId = &v
	return s
}

func (s *SwitchGdnMemberRoleRequest) SetSrcMainConnectString(v string) *SwitchGdnMemberRoleRequest {
	s.SrcMainConnectString = &v
	return s
}

func (s *SwitchGdnMemberRoleRequest) SetSrcMainPort(v string) *SwitchGdnMemberRoleRequest {
	s.SrcMainPort = &v
	return s
}

func (s *SwitchGdnMemberRoleRequest) SetSwitchMode(v string) *SwitchGdnMemberRoleRequest {
	s.SwitchMode = &v
	return s
}

func (s *SwitchGdnMemberRoleRequest) SetTaskTimeout(v int64) *SwitchGdnMemberRoleRequest {
	s.TaskTimeout = &v
	return s
}

func (s *SwitchGdnMemberRoleRequest) Validate() error {
	return dara.Validate(s)
}
