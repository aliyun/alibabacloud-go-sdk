// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderStatusCheckJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderStatusCheckJob(v *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) *GetIdentityProviderStatusCheckJobResponseBody
	GetIdentityProviderStatusCheckJob() *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob
	SetRequestId(v string) *GetIdentityProviderStatusCheckJobResponseBody
	GetRequestId() *string
}

type GetIdentityProviderStatusCheckJobResponseBody struct {
	// The information about the IdP status check job.
	IdentityProviderStatusCheckJob *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob `json:"IdentityProviderStatusCheckJob,omitempty" xml:"IdentityProviderStatusCheckJob,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIdentityProviderStatusCheckJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderStatusCheckJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderStatusCheckJobResponseBody) GetIdentityProviderStatusCheckJob() *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob {
	return s.IdentityProviderStatusCheckJob
}

func (s *GetIdentityProviderStatusCheckJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentityProviderStatusCheckJobResponseBody) SetIdentityProviderStatusCheckJob(v *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) *GetIdentityProviderStatusCheckJobResponseBody {
	s.IdentityProviderStatusCheckJob = v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBody) SetRequestId(v string) *GetIdentityProviderStatusCheckJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBody) Validate() error {
	if s.IdentityProviderStatusCheckJob != nil {
		if err := s.IdentityProviderStatusCheckJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob struct {
	// The result of the check task.
	//
	// example:
	//
	// success
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The end time.
	//
	// example:
	//
	// 1763776265757
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the IdP.
	//
	// example:
	//
	// idp_ncehkms65fiefobrvwy2blrxxx
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// The ID of the IdP status check job.
	//
	// example:
	//
	// async_000036tfu8cgngmakngrr2rk75qgf87pf3rxxx
	IdentityProviderStatusCheckJobId *string `json:"IdentityProviderStatusCheckJobId,omitempty" xml:"IdentityProviderStatusCheckJobId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ki6hd7ihir4ybawogqk6kqsfxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The results of the status check subtasks.
	JobCheckItems []*GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems `json:"JobCheckItems,omitempty" xml:"JobCheckItems,omitempty" type:"Repeated"`
	// The start time.
	//
	// example:
	//
	// 1763776265757
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the check task.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) GetCheckResult() *string {
	return s.CheckResult
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) GetIdentityProviderStatusCheckJobId() *string {
	return s.IdentityProviderStatusCheckJobId
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) GetJobCheckItems() []*GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems {
	return s.JobCheckItems
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) GetStatus() *string {
	return s.Status
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) SetCheckResult(v string) *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob {
	s.CheckResult = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) SetEndTime(v int64) *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob {
	s.EndTime = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) SetIdentityProviderId(v string) *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob {
	s.IdentityProviderId = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) SetIdentityProviderStatusCheckJobId(v string) *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob {
	s.IdentityProviderStatusCheckJobId = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) SetInstanceId(v string) *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob {
	s.InstanceId = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) SetJobCheckItems(v []*GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems) *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob {
	s.JobCheckItems = v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) SetStartTime(v int64) *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob {
	s.StartTime = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) SetStatus(v string) *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob {
	s.Status = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJob) Validate() error {
	if s.JobCheckItems != nil {
		for _, item := range s.JobCheckItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems struct {
	// The reason for the error.
	ErrorReason *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItemsErrorReason `json:"ErrorReason,omitempty" xml:"ErrorReason,omitempty" type:"Struct"`
	// The major check item.
	//
	// example:
	//
	// network_check
	MajorCheckItem *string `json:"MajorCheckItem,omitempty" xml:"MajorCheckItem,omitempty"`
	// The minor check item.
	//
	// example:
	//
	// network_access_status
	MinorCheckItem *string `json:"MinorCheckItem,omitempty" xml:"MinorCheckItem,omitempty"`
	// The result.
	//
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems) GetErrorReason() *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItemsErrorReason {
	return s.ErrorReason
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems) GetMajorCheckItem() *string {
	return s.MajorCheckItem
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems) GetMinorCheckItem() *string {
	return s.MinorCheckItem
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems) GetResult() *string {
	return s.Result
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems) SetErrorReason(v *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItemsErrorReason) *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems {
	s.ErrorReason = v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems) SetMajorCheckItem(v string) *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems {
	s.MajorCheckItem = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems) SetMinorCheckItem(v string) *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems {
	s.MinorCheckItem = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems) SetResult(v string) *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems {
	s.Result = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItems) Validate() error {
	if s.ErrorReason != nil {
		if err := s.ErrorReason.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItemsErrorReason struct {
	// The error code.
	//
	// example:
	//
	// NetworkAccessPointWarning.SingleNetworkAccessPath
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error level.
	//
	// example:
	//
	// high
	ErrorLevel *string `json:"ErrorLevel,omitempty" xml:"ErrorLevel,omitempty"`
	// The error message.
	//
	// example:
	//
	// There is only one path in the current network access endpoint.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItemsErrorReason) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItemsErrorReason) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItemsErrorReason) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItemsErrorReason) GetErrorLevel() *string {
	return s.ErrorLevel
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItemsErrorReason) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItemsErrorReason) SetErrorCode(v string) *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItemsErrorReason {
	s.ErrorCode = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItemsErrorReason) SetErrorLevel(v string) *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItemsErrorReason {
	s.ErrorLevel = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItemsErrorReason) SetErrorMessage(v string) *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItemsErrorReason {
	s.ErrorMessage = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponseBodyIdentityProviderStatusCheckJobJobCheckItemsErrorReason) Validate() error {
	return dara.Validate(s)
}
