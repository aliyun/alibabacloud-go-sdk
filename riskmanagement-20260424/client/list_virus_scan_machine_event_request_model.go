// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanMachineEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListVirusScanMachineEventRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListVirusScanMachineEventRequest
	GetLang() *string
	SetOperateTaskId(v string) *ListVirusScanMachineEventRequest
	GetOperateTaskId() *string
	SetPageSize(v int32) *ListVirusScanMachineEventRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListVirusScanMachineEventRequest
	GetRegionId() *string
	SetUuid(v string) *ListVirusScanMachineEventRequest
	GetUuid() *string
}

type ListVirusScanMachineEventRequest struct {
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OperateTaskId *string `json:"OperateTaskId,omitempty" xml:"OperateTaskId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Uuid          *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListVirusScanMachineEventRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineEventRequest) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineEventRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListVirusScanMachineEventRequest) GetLang() *string {
	return s.Lang
}

func (s *ListVirusScanMachineEventRequest) GetOperateTaskId() *string {
	return s.OperateTaskId
}

func (s *ListVirusScanMachineEventRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVirusScanMachineEventRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVirusScanMachineEventRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListVirusScanMachineEventRequest) SetCurrentPage(v int32) *ListVirusScanMachineEventRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListVirusScanMachineEventRequest) SetLang(v string) *ListVirusScanMachineEventRequest {
	s.Lang = &v
	return s
}

func (s *ListVirusScanMachineEventRequest) SetOperateTaskId(v string) *ListVirusScanMachineEventRequest {
	s.OperateTaskId = &v
	return s
}

func (s *ListVirusScanMachineEventRequest) SetPageSize(v int32) *ListVirusScanMachineEventRequest {
	s.PageSize = &v
	return s
}

func (s *ListVirusScanMachineEventRequest) SetRegionId(v string) *ListVirusScanMachineEventRequest {
	s.RegionId = &v
	return s
}

func (s *ListVirusScanMachineEventRequest) SetUuid(v string) *ListVirusScanMachineEventRequest {
	s.Uuid = &v
	return s
}

func (s *ListVirusScanMachineEventRequest) Validate() error {
	return dara.Validate(s)
}
