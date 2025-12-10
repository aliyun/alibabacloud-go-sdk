// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountDeletionCheckResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDeletionCheckResultInfo(v *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) *GetAccountDeletionCheckResultResponseBody
	GetAccountDeletionCheckResultInfo() *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo
	SetRequestId(v string) *GetAccountDeletionCheckResultResponseBody
	GetRequestId() *string
}

type GetAccountDeletionCheckResultResponseBody struct {
	// The result of the deletion check for the member.
	AccountDeletionCheckResultInfo *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo `json:"AccountDeletionCheckResultInfo,omitempty" xml:"AccountDeletionCheckResultInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 54AC391D-4F7F-5F08-B8D3-0AECDE6EC5BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountDeletionCheckResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountDeletionCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultResponseBody) GetAccountDeletionCheckResultInfo() *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo {
	return s.AccountDeletionCheckResultInfo
}

func (s *GetAccountDeletionCheckResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountDeletionCheckResultResponseBody) SetAccountDeletionCheckResultInfo(v *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) *GetAccountDeletionCheckResultResponseBody {
	s.AccountDeletionCheckResultInfo = v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBody) SetRequestId(v string) *GetAccountDeletionCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBody) Validate() error {
	if s.AccountDeletionCheckResultInfo != nil {
		if err := s.AccountDeletionCheckResultInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo struct {
	// The check items that you can choose to ignore for the member deletion.
	//
	// >  This parameter may be returned if the value of AllowDelete is true.
	AbandonableChecks []*GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks `json:"AbandonableChecks,omitempty" xml:"AbandonableChecks,omitempty" type:"Repeated"`
	// Indicates whether the member can be deleted. Valid values:
	//
	// 	- true: The member can be deleted.
	//
	// 	- false: The member cannot be deleted.
	//
	// example:
	//
	// false
	AllowDelete *string `json:"AllowDelete,omitempty" xml:"AllowDelete,omitempty"`
	// The reasons why the member cannot be deleted.
	//
	// >  This parameter is returned only if the value of AllowDelete is false.
	NotAllowReason []*GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason `json:"NotAllowReason,omitempty" xml:"NotAllowReason,omitempty" type:"Repeated"`
	// The status of the check. Valid values:
	//
	// 	- PreCheckComplete: The check is complete.
	//
	// 	- PreChecking: The check is in progress.
	//
	// example:
	//
	// PreCheckComplete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) GetAbandonableChecks() []*GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks {
	return s.AbandonableChecks
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) GetAllowDelete() *string {
	return s.AllowDelete
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) GetNotAllowReason() []*GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason {
	return s.NotAllowReason
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) GetStatus() *string {
	return s.Status
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) SetAbandonableChecks(v []*GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo {
	s.AbandonableChecks = v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) SetAllowDelete(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo {
	s.AllowDelete = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) SetNotAllowReason(v []*GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo {
	s.NotAllowReason = v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) SetStatus(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo {
	s.Status = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfo) Validate() error {
	if s.AbandonableChecks != nil {
		for _, item := range s.AbandonableChecks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NotAllowReason != nil {
		for _, item := range s.NotAllowReason {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks struct {
	// The ID of the check item.
	//
	// example:
	//
	// NON_SP_cs
	CheckId *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The name of the cloud service to which the check item belongs.
	//
	// example:
	//
	// Container Service for Kubernetes
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// The description of the check item.
	//
	// example:
	//
	// An instance of a cloud service is running within the member. Submit a ticket to contact Alibaba Cloud technical support.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) String() string {
	return dara.Prettify(s)
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) GetCheckId() *string {
	return s.CheckId
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) GetCheckName() *string {
	return s.CheckName
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) GetDescription() *string {
	return s.Description
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) SetCheckId(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks {
	s.CheckId = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) SetCheckName(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks {
	s.CheckName = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) SetDescription(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks {
	s.Description = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoAbandonableChecks) Validate() error {
	return dara.Validate(s)
}

type GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason struct {
	// The ID of the check item.
	//
	// example:
	//
	// NON_SP_efc
	CheckId *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The name of the cloud service to which the check item belongs.
	//
	// example:
	//
	// Enterprise finance
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// The description of the check item.
	//
	// example:
	//
	// This account is an Enterprise Finance associated account. Please remove the financial association of this account before deleting it.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) String() string {
	return dara.Prettify(s)
}

func (s GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) GoString() string {
	return s.String()
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) GetCheckId() *string {
	return s.CheckId
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) GetCheckName() *string {
	return s.CheckName
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) GetDescription() *string {
	return s.Description
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) SetCheckId(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason {
	s.CheckId = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) SetCheckName(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason {
	s.CheckName = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) SetDescription(v string) *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason {
	s.Description = &v
	return s
}

func (s *GetAccountDeletionCheckResultResponseBodyAccountDeletionCheckResultInfoNotAllowReason) Validate() error {
	return dara.Validate(s)
}
