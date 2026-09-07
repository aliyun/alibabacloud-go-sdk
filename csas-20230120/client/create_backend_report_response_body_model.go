// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackendReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedCount(v int32) *CreateBackendReportResponseBody
	GetFailedCount() *int32
	SetItems(v []*CreateBackendReportResponseBodyItems) *CreateBackendReportResponseBody
	GetItems() []*CreateBackendReportResponseBodyItems
	SetObjectCount(v int32) *CreateBackendReportResponseBody
	GetObjectCount() *int32
	SetRequestId(v string) *CreateBackendReportResponseBody
	GetRequestId() *string
	SetSuccessCount(v int32) *CreateBackendReportResponseBody
	GetSuccessCount() *int32
	SetTargetCount(v int32) *CreateBackendReportResponseBody
	GetTargetCount() *int32
	SetTotalCount(v int32) *CreateBackendReportResponseBody
	GetTotalCount() *int32
}

type CreateBackendReportResponseBody struct {
	// The number of user-object combinations that failed to be created.
	//
	// example:
	//
	// 0
	FailedCount *int32 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The processing results for each user-object combination. If some combinations fail, the operation still returns results for all combinations.
	Items []*CreateBackendReportResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of deduplicated filing objects.
	//
	// example:
	//
	// 1
	ObjectCount *int32 `json:"ObjectCount,omitempty" xml:"ObjectCount,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of user-object combinations that are created.
	//
	// example:
	//
	// 1
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The number of deduplicated filing users.
	//
	// example:
	//
	// 1
	TargetCount *int32 `json:"TargetCount,omitempty" xml:"TargetCount,omitempty"`
	// The total number of expanded user-object combinations.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CreateBackendReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBackendReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackendReportResponseBody) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *CreateBackendReportResponseBody) GetItems() []*CreateBackendReportResponseBodyItems {
	return s.Items
}

func (s *CreateBackendReportResponseBody) GetObjectCount() *int32 {
	return s.ObjectCount
}

func (s *CreateBackendReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBackendReportResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *CreateBackendReportResponseBody) GetTargetCount() *int32 {
	return s.TargetCount
}

func (s *CreateBackendReportResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *CreateBackendReportResponseBody) SetFailedCount(v int32) *CreateBackendReportResponseBody {
	s.FailedCount = &v
	return s
}

func (s *CreateBackendReportResponseBody) SetItems(v []*CreateBackendReportResponseBodyItems) *CreateBackendReportResponseBody {
	s.Items = v
	return s
}

func (s *CreateBackendReportResponseBody) SetObjectCount(v int32) *CreateBackendReportResponseBody {
	s.ObjectCount = &v
	return s
}

func (s *CreateBackendReportResponseBody) SetRequestId(v string) *CreateBackendReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackendReportResponseBody) SetSuccessCount(v int32) *CreateBackendReportResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *CreateBackendReportResponseBody) SetTargetCount(v int32) *CreateBackendReportResponseBody {
	s.TargetCount = &v
	return s
}

func (s *CreateBackendReportResponseBody) SetTotalCount(v int32) *CreateBackendReportResponseBody {
	s.TotalCount = &v
	return s
}

func (s *CreateBackendReportResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateBackendReportResponseBodyItems struct {
	// The approval instance ID generated after a successful creation. This parameter is not returned if the creation fails.
	//
	// example:
	//
	// approval-6b5188a28634****
	ApprovalId *string `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	// The error code returned when the current combination fails to be created. This parameter is not returned if the creation succeeds.
	//
	// example:
	//
	// TargetNotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The filing effective status. Enabled is returned when the creation succeeds, which indicates that the filing is valid.
	//
	// example:
	//
	// Enabled
	EffectStatus *string `json:"EffectStatus,omitempty" xml:"EffectStatus,omitempty"`
	// The error message returned when the current combination fails to be created. This parameter is not returned if the creation succeeds.
	//
	// example:
	//
	// target user is not found
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The filing object corresponding to the current combination. The fields vary based on the PolicyType value.
	//
	// example:
	//
	// [{"ApplicationId":"pa-application-92b60359213a****"}]
	ReportObject interface{} `json:"ReportObject,omitempty" xml:"ReportObject,omitempty"`
	// The filing type. BackendReport is always returned when the creation succeeds, which indicates a backend filing.
	//
	// example:
	//
	// BackendReport
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The approval status. Approved is returned when the creation succeeds, which indicates that the filing is approved.
	//
	// example:
	//
	// Approved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the current combination is created.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The filing user corresponding to the current combination.
	Target *CreateBackendReportResponseBodyItemsTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
}

func (s CreateBackendReportResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s CreateBackendReportResponseBodyItems) GoString() string {
	return s.String()
}

func (s *CreateBackendReportResponseBodyItems) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *CreateBackendReportResponseBodyItems) GetCode() *string {
	return s.Code
}

func (s *CreateBackendReportResponseBodyItems) GetEffectStatus() *string {
	return s.EffectStatus
}

func (s *CreateBackendReportResponseBodyItems) GetMessage() *string {
	return s.Message
}

func (s *CreateBackendReportResponseBodyItems) GetReportObject() interface{} {
	return s.ReportObject
}

func (s *CreateBackendReportResponseBodyItems) GetReportType() *string {
	return s.ReportType
}

func (s *CreateBackendReportResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *CreateBackendReportResponseBodyItems) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBackendReportResponseBodyItems) GetTarget() *CreateBackendReportResponseBodyItemsTarget {
	return s.Target
}

func (s *CreateBackendReportResponseBodyItems) SetApprovalId(v string) *CreateBackendReportResponseBodyItems {
	s.ApprovalId = &v
	return s
}

func (s *CreateBackendReportResponseBodyItems) SetCode(v string) *CreateBackendReportResponseBodyItems {
	s.Code = &v
	return s
}

func (s *CreateBackendReportResponseBodyItems) SetEffectStatus(v string) *CreateBackendReportResponseBodyItems {
	s.EffectStatus = &v
	return s
}

func (s *CreateBackendReportResponseBodyItems) SetMessage(v string) *CreateBackendReportResponseBodyItems {
	s.Message = &v
	return s
}

func (s *CreateBackendReportResponseBodyItems) SetReportObject(v interface{}) *CreateBackendReportResponseBodyItems {
	s.ReportObject = v
	return s
}

func (s *CreateBackendReportResponseBodyItems) SetReportType(v string) *CreateBackendReportResponseBodyItems {
	s.ReportType = &v
	return s
}

func (s *CreateBackendReportResponseBodyItems) SetStatus(v string) *CreateBackendReportResponseBodyItems {
	s.Status = &v
	return s
}

func (s *CreateBackendReportResponseBodyItems) SetSuccess(v bool) *CreateBackendReportResponseBodyItems {
	s.Success = &v
	return s
}

func (s *CreateBackendReportResponseBodyItems) SetTarget(v *CreateBackendReportResponseBodyItemsTarget) *CreateBackendReportResponseBodyItems {
	s.Target = v
	return s
}

func (s *CreateBackendReportResponseBodyItems) Validate() error {
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBackendReportResponseBodyItemsTarget struct {
	// The SASE user ID.
	//
	// example:
	//
	// su_1b91e674235a25e4117faf5c36a8ad4e69a14303247fb7d9f2046ce8b372****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateBackendReportResponseBodyItemsTarget) String() string {
	return dara.Prettify(s)
}

func (s CreateBackendReportResponseBodyItemsTarget) GoString() string {
	return s.String()
}

func (s *CreateBackendReportResponseBodyItemsTarget) GetUserId() *string {
	return s.UserId
}

func (s *CreateBackendReportResponseBodyItemsTarget) SetUserId(v string) *CreateBackendReportResponseBodyItemsTarget {
	s.UserId = &v
	return s
}

func (s *CreateBackendReportResponseBodyItemsTarget) Validate() error {
	return dara.Validate(s)
}
