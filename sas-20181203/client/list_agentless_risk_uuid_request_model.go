// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessRiskUuidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAgentlessRiskUuidRequest
	GetCurrentPage() *int32
	SetInstanceId(v string) *ListAgentlessRiskUuidRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ListAgentlessRiskUuidRequest
	GetInstanceName() *string
	SetInternetIp(v string) *ListAgentlessRiskUuidRequest
	GetInternetIp() *string
	SetIntranetIp(v string) *ListAgentlessRiskUuidRequest
	GetIntranetIp() *string
	SetMachineName(v string) *ListAgentlessRiskUuidRequest
	GetMachineName() *string
	SetPageSize(v int32) *ListAgentlessRiskUuidRequest
	GetPageSize() *int32
	SetRisk(v bool) *ListAgentlessRiskUuidRequest
	GetRisk() *bool
	SetTargetName(v string) *ListAgentlessRiskUuidRequest
	GetTargetName() *string
	SetTargetType(v int32) *ListAgentlessRiskUuidRequest
	GetTargetType() *int32
}

type ListAgentlessRiskUuidRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The instance ID of the asset.
	//
	// example:
	//
	// s-bp1g6wxdwps7s9dz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name of the asset.
	//
	// example:
	//
	// ca_cpm_****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the asset that you want to query.
	//
	// example:
	//
	// 1.1.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the asset that you want to query.
	//
	// example:
	//
	// 172.26.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// oracle-win-001****
	MachineName *string `json:"MachineName,omitempty" xml:"MachineName,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether risks exist. Valid values:
	//
	// 	- **true**: Risks exist.
	//
	// 	- **false**: Risks do not exist.
	//
	// example:
	//
	// true
	Risk *bool `json:"Risk,omitempty" xml:"Risk,omitempty"`
	// The name of the detection object.
	//
	// example:
	//
	// source-test-obj-0****
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// Specifies the type of the object being inspected. Valid values:
	//
	// 	- **1**: Host Snapshot.
	//
	// 	- **2**: Host Image.
	//
	// 	- **3**: User Snapshot.
	//
	// 	- **4**: User Image.
	//
	// example:
	//
	// 3
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListAgentlessRiskUuidRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessRiskUuidRequest) GoString() string {
	return s.String()
}

func (s *ListAgentlessRiskUuidRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAgentlessRiskUuidRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAgentlessRiskUuidRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListAgentlessRiskUuidRequest) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListAgentlessRiskUuidRequest) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListAgentlessRiskUuidRequest) GetMachineName() *string {
	return s.MachineName
}

func (s *ListAgentlessRiskUuidRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentlessRiskUuidRequest) GetRisk() *bool {
	return s.Risk
}

func (s *ListAgentlessRiskUuidRequest) GetTargetName() *string {
	return s.TargetName
}

func (s *ListAgentlessRiskUuidRequest) GetTargetType() *int32 {
	return s.TargetType
}

func (s *ListAgentlessRiskUuidRequest) SetCurrentPage(v int32) *ListAgentlessRiskUuidRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAgentlessRiskUuidRequest) SetInstanceId(v string) *ListAgentlessRiskUuidRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAgentlessRiskUuidRequest) SetInstanceName(v string) *ListAgentlessRiskUuidRequest {
	s.InstanceName = &v
	return s
}

func (s *ListAgentlessRiskUuidRequest) SetInternetIp(v string) *ListAgentlessRiskUuidRequest {
	s.InternetIp = &v
	return s
}

func (s *ListAgentlessRiskUuidRequest) SetIntranetIp(v string) *ListAgentlessRiskUuidRequest {
	s.IntranetIp = &v
	return s
}

func (s *ListAgentlessRiskUuidRequest) SetMachineName(v string) *ListAgentlessRiskUuidRequest {
	s.MachineName = &v
	return s
}

func (s *ListAgentlessRiskUuidRequest) SetPageSize(v int32) *ListAgentlessRiskUuidRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentlessRiskUuidRequest) SetRisk(v bool) *ListAgentlessRiskUuidRequest {
	s.Risk = &v
	return s
}

func (s *ListAgentlessRiskUuidRequest) SetTargetName(v string) *ListAgentlessRiskUuidRequest {
	s.TargetName = &v
	return s
}

func (s *ListAgentlessRiskUuidRequest) SetTargetType(v int32) *ListAgentlessRiskUuidRequest {
	s.TargetType = &v
	return s
}

func (s *ListAgentlessRiskUuidRequest) Validate() error {
	return dara.Validate(s)
}
