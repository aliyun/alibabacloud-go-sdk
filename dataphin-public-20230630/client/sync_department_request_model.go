// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDepartmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SyncDepartmentRequest
	GetOpTenantId() *int64
	SetSyncDepartmentCommand(v *SyncDepartmentRequestSyncDepartmentCommand) *SyncDepartmentRequest
	GetSyncDepartmentCommand() *SyncDepartmentRequestSyncDepartmentCommand
}

type SyncDepartmentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	SyncDepartmentCommand *SyncDepartmentRequestSyncDepartmentCommand `json:"SyncDepartmentCommand,omitempty" xml:"SyncDepartmentCommand,omitempty" type:"Struct"`
}

func (s SyncDepartmentRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncDepartmentRequest) GoString() string {
	return s.String()
}

func (s *SyncDepartmentRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SyncDepartmentRequest) GetSyncDepartmentCommand() *SyncDepartmentRequestSyncDepartmentCommand {
	return s.SyncDepartmentCommand
}

func (s *SyncDepartmentRequest) SetOpTenantId(v int64) *SyncDepartmentRequest {
	s.OpTenantId = &v
	return s
}

func (s *SyncDepartmentRequest) SetSyncDepartmentCommand(v *SyncDepartmentRequestSyncDepartmentCommand) *SyncDepartmentRequest {
	s.SyncDepartmentCommand = v
	return s
}

func (s *SyncDepartmentRequest) Validate() error {
	if s.SyncDepartmentCommand != nil {
		if err := s.SyncDepartmentCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SyncDepartmentRequestSyncDepartmentCommand struct {
	// This parameter is required.
	DepartmentList []*SyncDepartmentRequestSyncDepartmentCommandDepartmentList `json:"DepartmentList,omitempty" xml:"DepartmentList,omitempty" type:"Repeated"`
}

func (s SyncDepartmentRequestSyncDepartmentCommand) String() string {
	return dara.Prettify(s)
}

func (s SyncDepartmentRequestSyncDepartmentCommand) GoString() string {
	return s.String()
}

func (s *SyncDepartmentRequestSyncDepartmentCommand) GetDepartmentList() []*SyncDepartmentRequestSyncDepartmentCommandDepartmentList {
	return s.DepartmentList
}

func (s *SyncDepartmentRequestSyncDepartmentCommand) SetDepartmentList(v []*SyncDepartmentRequestSyncDepartmentCommandDepartmentList) *SyncDepartmentRequestSyncDepartmentCommand {
	s.DepartmentList = v
	return s
}

func (s *SyncDepartmentRequestSyncDepartmentCommand) Validate() error {
	if s.DepartmentList != nil {
		for _, item := range s.DepartmentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SyncDepartmentRequestSyncDepartmentCommandDepartmentList struct {
	// This parameter is required.
	//
	// example:
	//
	// 10001
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 研发中心
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// example:
	//
	// 10001
	ParentDepartmentId *string `json:"ParentDepartmentId,omitempty" xml:"ParentDepartmentId,omitempty"`
}

func (s SyncDepartmentRequestSyncDepartmentCommandDepartmentList) String() string {
	return dara.Prettify(s)
}

func (s SyncDepartmentRequestSyncDepartmentCommandDepartmentList) GoString() string {
	return s.String()
}

func (s *SyncDepartmentRequestSyncDepartmentCommandDepartmentList) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *SyncDepartmentRequestSyncDepartmentCommandDepartmentList) GetDepartmentName() *string {
	return s.DepartmentName
}

func (s *SyncDepartmentRequestSyncDepartmentCommandDepartmentList) GetParentDepartmentId() *string {
	return s.ParentDepartmentId
}

func (s *SyncDepartmentRequestSyncDepartmentCommandDepartmentList) SetDepartmentId(v string) *SyncDepartmentRequestSyncDepartmentCommandDepartmentList {
	s.DepartmentId = &v
	return s
}

func (s *SyncDepartmentRequestSyncDepartmentCommandDepartmentList) SetDepartmentName(v string) *SyncDepartmentRequestSyncDepartmentCommandDepartmentList {
	s.DepartmentName = &v
	return s
}

func (s *SyncDepartmentRequestSyncDepartmentCommandDepartmentList) SetParentDepartmentId(v string) *SyncDepartmentRequestSyncDepartmentCommandDepartmentList {
	s.ParentDepartmentId = &v
	return s
}

func (s *SyncDepartmentRequestSyncDepartmentCommandDepartmentList) Validate() error {
	return dara.Validate(s)
}
