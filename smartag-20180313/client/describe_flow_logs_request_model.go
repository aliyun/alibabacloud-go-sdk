// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeFlowLogsRequest
	GetDescription() *string
	SetFlowLogId(v string) *DescribeFlowLogsRequest
	GetFlowLogId() *string
	SetFlowLogName(v string) *DescribeFlowLogsRequest
	GetFlowLogName() *string
	SetOutputType(v string) *DescribeFlowLogsRequest
	GetOutputType() *string
	SetOwnerAccount(v string) *DescribeFlowLogsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeFlowLogsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeFlowLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFlowLogsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeFlowLogsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeFlowLogsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeFlowLogsRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *DescribeFlowLogsRequest
	GetStatus() *string
}

type DescribeFlowLogsRequest struct {
	// The description of the flow log instance.
	//
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the flow log instance.
	//
	// example:
	//
	// fl-7a56mar1kfw9vj****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// The name of the flow log instance.
	//
	// example:
	//
	// DDE
	FlowLogName *string `json:"FlowLogName,omitempty" xml:"FlowLogName,omitempty"`
	// The output type of the flow log. Valid values:
	//
	// - **sls**: Flow logs are stored in Simple Log Service.
	//
	// - **netflow**: Flow logs are stored on the NetFlow server that you configured.
	//
	// - **all**: Flow logs are stored in both Simple Log Service and the NetFlow server that you configured.
	//
	// example:
	//
	// all
	OutputType   *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page in a paging query. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the flow log instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the flow log instance. Valid values:
	//
	// - **Active**: Started.
	//
	// - **Inactive**: Not started.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeFlowLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogsRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeFlowLogsRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *DescribeFlowLogsRequest) GetFlowLogName() *string {
	return s.FlowLogName
}

func (s *DescribeFlowLogsRequest) GetOutputType() *string {
	return s.OutputType
}

func (s *DescribeFlowLogsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeFlowLogsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeFlowLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFlowLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFlowLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFlowLogsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeFlowLogsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeFlowLogsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeFlowLogsRequest) SetDescription(v string) *DescribeFlowLogsRequest {
	s.Description = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetFlowLogId(v string) *DescribeFlowLogsRequest {
	s.FlowLogId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetFlowLogName(v string) *DescribeFlowLogsRequest {
	s.FlowLogName = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetOutputType(v string) *DescribeFlowLogsRequest {
	s.OutputType = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetOwnerAccount(v string) *DescribeFlowLogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetOwnerId(v int64) *DescribeFlowLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetPageNumber(v int32) *DescribeFlowLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetPageSize(v int32) *DescribeFlowLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetRegionId(v string) *DescribeFlowLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetResourceOwnerAccount(v string) *DescribeFlowLogsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetResourceOwnerId(v int64) *DescribeFlowLogsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeFlowLogsRequest) SetStatus(v string) *DescribeFlowLogsRequest {
	s.Status = &v
	return s
}

func (s *DescribeFlowLogsRequest) Validate() error {
	return dara.Validate(s)
}
