// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLPlanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSQLPlanTaskRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeSQLPlanTaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSQLPlanTaskRequest
	GetOwnerId() *int64
	SetProcessId(v string) *DescribeSQLPlanTaskRequest
	GetProcessId() *string
	SetResourceOwnerAccount(v string) *DescribeSQLPlanTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSQLPlanTaskRequest
	GetResourceOwnerId() *int64
	SetStageId(v string) *DescribeSQLPlanTaskRequest
	GetStageId() *string
}

type DescribeSQLPlanTaskRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 201907241445301720211111570315107****
	ProcessId            *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The stage of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1785135913****
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
}

func (s DescribeSQLPlanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLPlanTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanTaskRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSQLPlanTaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSQLPlanTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSQLPlanTaskRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *DescribeSQLPlanTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSQLPlanTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSQLPlanTaskRequest) GetStageId() *string {
	return s.StageId
}

func (s *DescribeSQLPlanTaskRequest) SetDBClusterId(v string) *DescribeSQLPlanTaskRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetOwnerAccount(v string) *DescribeSQLPlanTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetOwnerId(v int64) *DescribeSQLPlanTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetProcessId(v string) *DescribeSQLPlanTaskRequest {
	s.ProcessId = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetResourceOwnerAccount(v string) *DescribeSQLPlanTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetResourceOwnerId(v int64) *DescribeSQLPlanTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) SetStageId(v string) *DescribeSQLPlanTaskRequest {
	s.StageId = &v
	return s
}

func (s *DescribeSQLPlanTaskRequest) Validate() error {
	return dara.Validate(s)
}
