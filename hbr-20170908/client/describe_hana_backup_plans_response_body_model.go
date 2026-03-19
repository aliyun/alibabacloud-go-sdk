// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaBackupPlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHanaBackupPlansResponseBody
	GetCode() *string
	SetHanaBackupPlans(v *DescribeHanaBackupPlansResponseBodyHanaBackupPlans) *DescribeHanaBackupPlansResponseBody
	GetHanaBackupPlans() *DescribeHanaBackupPlansResponseBodyHanaBackupPlans
	SetMessage(v string) *DescribeHanaBackupPlansResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeHanaBackupPlansResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHanaBackupPlansResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHanaBackupPlansResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeHanaBackupPlansResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeHanaBackupPlansResponseBody
	GetTotalCount() *int64
}

type DescribeHanaBackupPlansResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code            *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	HanaBackupPlans *DescribeHanaBackupPlansResponseBodyHanaBackupPlans `json:"HanaBackupPlans,omitempty" xml:"HanaBackupPlans,omitempty" type:"Struct"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F029C1C7-26B6-5ADD-A73E-D85CCD7C73A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 6
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHanaBackupPlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaBackupPlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupPlansResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHanaBackupPlansResponseBody) GetHanaBackupPlans() *DescribeHanaBackupPlansResponseBodyHanaBackupPlans {
	return s.HanaBackupPlans
}

func (s *DescribeHanaBackupPlansResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHanaBackupPlansResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHanaBackupPlansResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHanaBackupPlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHanaBackupPlansResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeHanaBackupPlansResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeHanaBackupPlansResponseBody) SetCode(v string) *DescribeHanaBackupPlansResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetHanaBackupPlans(v *DescribeHanaBackupPlansResponseBodyHanaBackupPlans) *DescribeHanaBackupPlansResponseBody {
	s.HanaBackupPlans = v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetMessage(v string) *DescribeHanaBackupPlansResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetPageNumber(v int32) *DescribeHanaBackupPlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetPageSize(v int32) *DescribeHanaBackupPlansResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetRequestId(v string) *DescribeHanaBackupPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetSuccess(v bool) *DescribeHanaBackupPlansResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) SetTotalCount(v int64) *DescribeHanaBackupPlansResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBody) Validate() error {
	if s.HanaBackupPlans != nil {
		if err := s.HanaBackupPlans.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHanaBackupPlansResponseBodyHanaBackupPlans struct {
	HanaBackupPlan []*DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan `json:"HanaBackupPlan,omitempty" xml:"HanaBackupPlan,omitempty" type:"Repeated"`
}

func (s DescribeHanaBackupPlansResponseBodyHanaBackupPlans) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaBackupPlansResponseBodyHanaBackupPlans) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlans) GetHanaBackupPlan() []*DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	return s.HanaBackupPlan
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlans) SetHanaBackupPlan(v []*DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) *DescribeHanaBackupPlansResponseBodyHanaBackupPlans {
	s.HanaBackupPlan = v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlans) Validate() error {
	if s.HanaBackupPlan != nil {
		for _, item := range s.HanaBackupPlan {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan struct {
	BackupPrefix   *string `json:"BackupPrefix,omitempty" xml:"BackupPrefix,omitempty"`
	BackupType     *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DatabaseName   *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Disabled       *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	PlanId         *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PlanName       *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	Schedule       *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	VaultId        *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetBackupPrefix() *string {
	return s.BackupPrefix
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetPlanId() *string {
	return s.PlanId
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetPlanName() *string {
	return s.PlanName
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetSchedule() *string {
	return s.Schedule
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetBackupPrefix(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.BackupPrefix = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetBackupType(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.BackupType = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetBusinessStatus(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetClusterId(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetDatabaseName(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetDisabled(v bool) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.Disabled = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetPlanId(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.PlanId = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetPlanName(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.PlanName = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetSchedule(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.Schedule = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) SetVaultId(v string) *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan {
	s.VaultId = &v
	return s
}

func (s *DescribeHanaBackupPlansResponseBodyHanaBackupPlansHanaBackupPlan) Validate() error {
	return dara.Validate(s)
}
