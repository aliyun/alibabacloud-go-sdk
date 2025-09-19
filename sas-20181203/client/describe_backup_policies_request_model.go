// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeBackupPoliciesRequest
	GetCurrentPage() *int32
	SetMachineRemark(v string) *DescribeBackupPoliciesRequest
	GetMachineRemark() *string
	SetName(v string) *DescribeBackupPoliciesRequest
	GetName() *string
	SetPageSize(v int32) *DescribeBackupPoliciesRequest
	GetPageSize() *int32
	SetStatus(v string) *DescribeBackupPoliciesRequest
	GetStatus() *string
}

type DescribeBackupPoliciesRequest struct {
	// The number of the page to return. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The information that you want to use to identify the servers protected by the anti-ransomware policy. You can enter the IP address or ID of a server.
	//
	// example:
	//
	// 1.1.XX.XX
	MachineRemark *string `json:"MachineRemark,omitempty" xml:"MachineRemark,omitempty"`
	// The name of the anti-ransomware policy that you want to query.
	//
	// example:
	//
	// SecurityStrategy-20200303
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the anti-ransomware policy. Valid values:
	//
	// 	- **enabled**: The anti-ransomware policy is manually enabled.
	//
	// 	- **disabled**: The anti-ransomware policy is manually disabled. After an anti-ransomware policy is disabled, the data backup task that is running based on the policy stops.
	//
	// 	- **closed**: The anti-ransomware policy automatically stops because the anti-ransomware capacity is insufficient.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBackupPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPoliciesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeBackupPoliciesRequest) GetMachineRemark() *string {
	return s.MachineRemark
}

func (s *DescribeBackupPoliciesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeBackupPoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupPoliciesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeBackupPoliciesRequest) SetCurrentPage(v int32) *DescribeBackupPoliciesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeBackupPoliciesRequest) SetMachineRemark(v string) *DescribeBackupPoliciesRequest {
	s.MachineRemark = &v
	return s
}

func (s *DescribeBackupPoliciesRequest) SetName(v string) *DescribeBackupPoliciesRequest {
	s.Name = &v
	return s
}

func (s *DescribeBackupPoliciesRequest) SetPageSize(v int32) *DescribeBackupPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupPoliciesRequest) SetStatus(v string) *DescribeBackupPoliciesRequest {
	s.Status = &v
	return s
}

func (s *DescribeBackupPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
