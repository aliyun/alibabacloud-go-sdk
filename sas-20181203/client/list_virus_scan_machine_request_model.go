// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanMachineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListVirusScanMachineRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListVirusScanMachineRequest
	GetPageSize() *int32
	SetRemark(v string) *ListVirusScanMachineRequest
	GetRemark() *string
	SetUuid(v string) *ListVirusScanMachineRequest
	GetUuid() *string
}

type ListVirusScanMachineRequest struct {
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about the server that you want to query. The value can be the name or the IP address of the server.
	//
	// example:
	//
	// 192.168.1****
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 7cc91747-2845-40d4-bb69-c077597f****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListVirusScanMachineRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineRequest) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListVirusScanMachineRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVirusScanMachineRequest) GetRemark() *string {
	return s.Remark
}

func (s *ListVirusScanMachineRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListVirusScanMachineRequest) SetCurrentPage(v int32) *ListVirusScanMachineRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListVirusScanMachineRequest) SetPageSize(v int32) *ListVirusScanMachineRequest {
	s.PageSize = &v
	return s
}

func (s *ListVirusScanMachineRequest) SetRemark(v string) *ListVirusScanMachineRequest {
	s.Remark = &v
	return s
}

func (s *ListVirusScanMachineRequest) SetUuid(v string) *ListVirusScanMachineRequest {
	s.Uuid = &v
	return s
}

func (s *ListVirusScanMachineRequest) Validate() error {
	return dara.Validate(s)
}
