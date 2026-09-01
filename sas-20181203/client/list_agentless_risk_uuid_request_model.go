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
	// The page number of the current page when using paging.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the asset instance.
	//
	// example:
	//
	// s-bp1g6wxdwps7s9dz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the asset instance.
	//
	// example:
	//
	// ca_cpm_****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the asset to query.
	//
	// example:
	//
	// 1.1.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the asset to query.
	//
	// example:
	//
	// 172.26.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The instance name.
	//
	// example:
	//
	// oracle-win-001****
	MachineName *string `json:"MachineName,omitempty" xml:"MachineName,omitempty"`
	// The maximum number of entries to return per page when using paging.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether risks exist. Valid values:
	//
	// - **true**: Risks exist.
	//
	// - **false**: Risks do not exist.
	//
	// example:
	//
	// true
	Risk *bool `json:"Risk,omitempty" xml:"Risk,omitempty"`
	// The name of the detection target.
	//
	// example:
	//
	// source-test-obj-0****
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The object type of the detection target. Valid values:
	//
	// - **1**: host snapshot
	//
	// - **2**: host image
	//
	// - **3**: user snapshot
	//
	// - **4**: user custom image
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
