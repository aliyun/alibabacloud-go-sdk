// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountDeletionStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRdAccountDeletionStatus(v *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) *GetAccountDeletionStatusResponseBody
	GetRdAccountDeletionStatus() *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus
	SetRequestId(v string) *GetAccountDeletionStatusResponseBody
	GetRequestId() *string
}

type GetAccountDeletionStatusResponseBody struct {
	// The deletion status of the member account.
	RdAccountDeletionStatus *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus `json:"RdAccountDeletionStatus,omitempty" xml:"RdAccountDeletionStatus,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8AA43293-7C8F-5730-8F2D-7F864EC092C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountDeletionStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountDeletionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionStatusResponseBody) GetRdAccountDeletionStatus() *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	return s.RdAccountDeletionStatus
}

func (s *GetAccountDeletionStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountDeletionStatusResponseBody) SetRdAccountDeletionStatus(v *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) *GetAccountDeletionStatusResponseBody {
	s.RdAccountDeletionStatus = v
	return s
}

func (s *GetAccountDeletionStatusResponseBody) SetRequestId(v string) *GetAccountDeletionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBody) Validate() error {
	if s.RdAccountDeletionStatus != nil {
		if err := s.RdAccountDeletionStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus struct {
	// The ID of the member account.
	//
	// example:
	//
	// 169946124551****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the deletion started, in RFC 3339 format with a UTC offset of +08:00.
	//
	// example:
	//
	// 2022-08-23T17:05:30+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the deletion ended, in RFC 3339 format with a UTC offset of +08:00.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-08-23T17:06:01+08:00
	DeletionTime *string `json:"DeletionTime,omitempty" xml:"DeletionTime,omitempty"`
	// The deletion type. Valid values:
	//
	// - 0: direct deletion. If the member account has no pay-as-you-go resources in the last 30 days, the system directly deletes the member account.
	//
	// - 1: cool-down period deletion. If the member account has pay-as-you-go resources in the last 30 days, the member account enters a cool-down period. The system starts to delete the member account after the cool-down period ends. For more information about the cool-down period, see [What is the cool-down period for member account deletion](https://help.aliyun.com/document_detail/446079.html).
	//
	// example:
	//
	// 0
	DeletionType *string `json:"DeletionType,omitempty" xml:"DeletionType,omitempty"`
	// The reasons for the deletion failure.
	FailReasonList []*GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList `json:"FailReasonList,omitempty" xml:"FailReasonList,omitempty" type:"Repeated"`
	// The status. Valid values:
	//
	// - Success: The member account is deleted.
	//
	// - Checking: The check is in progress.
	//
	// - Deleting: The deletion is in progress.
	//
	// - CheckFailed: The check failed.
	//
	// - DeleteFailed: The deletion failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) String() string {
	return dara.Prettify(s)
}

func (s GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) GetAccountId() *string {
	return s.AccountId
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) GetDeletionTime() *string {
	return s.DeletionTime
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) GetDeletionType() *string {
	return s.DeletionType
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) GetFailReasonList() []*GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList {
	return s.FailReasonList
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) GetStatus() *string {
	return s.Status
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetAccountId(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.AccountId = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetCreateTime(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.CreateTime = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetDeletionTime(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.DeletionTime = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetDeletionType(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.DeletionType = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetFailReasonList(v []*GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.FailReasonList = v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) SetStatus(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus {
	s.Status = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus) Validate() error {
	if s.FailReasonList != nil {
		for _, item := range s.FailReasonList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList struct {
	// The description of the check item.
	//
	// example:
	//
	// This account has a payer account. Please release the financial relationship of this account first.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The service name of the cloud service to which the check item belongs.
	//
	// example:
	//
	// Others
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) String() string {
	return dara.Prettify(s)
}

func (s GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) GetDescription() *string {
	return s.Description
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) GetName() *string {
	return s.Name
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) SetDescription(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList {
	s.Description = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) SetName(v string) *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList {
	s.Name = &v
	return s
}

func (s *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList) Validate() error {
	return dara.Validate(s)
}
