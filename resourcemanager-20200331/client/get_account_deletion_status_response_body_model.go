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
	// The deletion status of the member.
	RdAccountDeletionStatus *GetAccountDeletionStatusResponseBodyRdAccountDeletionStatus `json:"RdAccountDeletionStatus,omitempty" xml:"RdAccountDeletionStatus,omitempty" type:"Struct"`
	// The ID of the request.
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
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 169946124551****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The start time of the deletion.
	//
	// example:
	//
	// 2022-08-23T17:05:30+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The end time of the deletion.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-08-23T17:06:01+08:00
	DeletionTime *string `json:"DeletionTime,omitempty" xml:"DeletionTime,omitempty"`
	// The type of the deletion. Valid values:
	//
	// 	- 0: direct deletion. If the member does not have pay-as-you-go resources that are purchased within the previous 30 days, the system directly deletes the member.
	//
	// 	- 1: deletion with a silence period. If the member has pay-as-you-go resources that are purchased within the previous 30 days, the member enters a silence period. The system starts to delete the member until the silence period ends. For more information about the silence period, see [What is the silence period for member deletion?](https://help.aliyun.com/document_detail/446079.html)
	//
	// example:
	//
	// 0
	DeletionType *string `json:"DeletionType,omitempty" xml:"DeletionType,omitempty"`
	// The reasons why the member fails to be deleted.
	FailReasonList []*GetAccountDeletionStatusResponseBodyRdAccountDeletionStatusFailReasonList `json:"FailReasonList,omitempty" xml:"FailReasonList,omitempty" type:"Repeated"`
	// The status. Valid values:
	//
	// 	- Success: The member is deleted.
	//
	// 	- Checking: A deletion check is being performed for the member.
	//
	// 	- Deleting: The member is being deleted.
	//
	// 	- CheckFailed: The deletion check for the member fails.
	//
	// 	- DeleteFailed: The member fails to be deleted.
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
	// The name of the cloud service to which the check item belongs.
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
