// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserLatestPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountHandleStatus(v int32) *GetUserLatestPlanResponseBody
	GetAccountHandleStatus() *int32
	SetAccountHandleTime(v string) *GetUserLatestPlanResponseBody
	GetAccountHandleTime() *string
	SetAccountType(v int32) *GetUserLatestPlanResponseBody
	GetAccountType() *int32
	SetAgreementFirstSignTime(v string) *GetUserLatestPlanResponseBody
	GetAgreementFirstSignTime() *string
	SetAgreementLastSignTime(v string) *GetUserLatestPlanResponseBody
	GetAgreementLastSignTime() *string
	SetAgreementStatus(v int32) *GetUserLatestPlanResponseBody
	GetAgreementStatus() *int32
	SetDataHandleEndTime(v string) *GetUserLatestPlanResponseBody
	GetDataHandleEndTime() *string
	SetDataHandleStartTime(v string) *GetUserLatestPlanResponseBody
	GetDataHandleStartTime() *string
	SetDataHandleStatus(v int32) *GetUserLatestPlanResponseBody
	GetDataHandleStatus() *int32
	SetExclusivePlan(v int32) *GetUserLatestPlanResponseBody
	GetExclusivePlan() *int32
	SetNewAccountUid(v int64) *GetUserLatestPlanResponseBody
	GetNewAccountUid() *int64
	SetRequestId(v string) *GetUserLatestPlanResponseBody
	GetRequestId() *string
	SetStatus(v int32) *GetUserLatestPlanResponseBody
	GetStatus() *int32
	SetVendorRequestId(v string) *GetUserLatestPlanResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetUserLatestPlanResponseBody
	GetVendorType() *string
}

type GetUserLatestPlanResponseBody struct {
	// example:
	//
	// 1
	AccountHandleStatus *int32 `json:"accountHandleStatus,omitempty" xml:"accountHandleStatus,omitempty"`
	// example:
	//
	// 2022-02-02
	AccountHandleTime *string `json:"accountHandleTime,omitempty" xml:"accountHandleTime,omitempty"`
	// example:
	//
	// 0
	AccountType *int32 `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// example:
	//
	// 2022-02-05
	AgreementFirstSignTime *string `json:"agreementFirstSignTime,omitempty" xml:"agreementFirstSignTime,omitempty"`
	// example:
	//
	// 2022-02-06
	AgreementLastSignTime *string `json:"agreementLastSignTime,omitempty" xml:"agreementLastSignTime,omitempty"`
	// example:
	//
	// 1
	AgreementStatus *int32 `json:"agreementStatus,omitempty" xml:"agreementStatus,omitempty"`
	// example:
	//
	// 2022-02-04
	DataHandleEndTime *string `json:"dataHandleEndTime,omitempty" xml:"dataHandleEndTime,omitempty"`
	// example:
	//
	// 2022-02-03
	DataHandleStartTime *string `json:"dataHandleStartTime,omitempty" xml:"dataHandleStartTime,omitempty"`
	// example:
	//
	// 1
	DataHandleStatus *int32 `json:"dataHandleStatus,omitempty" xml:"dataHandleStatus,omitempty"`
	// example:
	//
	// 1
	ExclusivePlan *int32 `json:"exclusivePlan,omitempty" xml:"exclusivePlan,omitempty"`
	// example:
	//
	// 1
	NewAccountUid *int64 `json:"newAccountUid,omitempty" xml:"newAccountUid,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetUserLatestPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserLatestPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserLatestPlanResponseBody) GetAccountHandleStatus() *int32 {
	return s.AccountHandleStatus
}

func (s *GetUserLatestPlanResponseBody) GetAccountHandleTime() *string {
	return s.AccountHandleTime
}

func (s *GetUserLatestPlanResponseBody) GetAccountType() *int32 {
	return s.AccountType
}

func (s *GetUserLatestPlanResponseBody) GetAgreementFirstSignTime() *string {
	return s.AgreementFirstSignTime
}

func (s *GetUserLatestPlanResponseBody) GetAgreementLastSignTime() *string {
	return s.AgreementLastSignTime
}

func (s *GetUserLatestPlanResponseBody) GetAgreementStatus() *int32 {
	return s.AgreementStatus
}

func (s *GetUserLatestPlanResponseBody) GetDataHandleEndTime() *string {
	return s.DataHandleEndTime
}

func (s *GetUserLatestPlanResponseBody) GetDataHandleStartTime() *string {
	return s.DataHandleStartTime
}

func (s *GetUserLatestPlanResponseBody) GetDataHandleStatus() *int32 {
	return s.DataHandleStatus
}

func (s *GetUserLatestPlanResponseBody) GetExclusivePlan() *int32 {
	return s.ExclusivePlan
}

func (s *GetUserLatestPlanResponseBody) GetNewAccountUid() *int64 {
	return s.NewAccountUid
}

func (s *GetUserLatestPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserLatestPlanResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *GetUserLatestPlanResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetUserLatestPlanResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetUserLatestPlanResponseBody) SetAccountHandleStatus(v int32) *GetUserLatestPlanResponseBody {
	s.AccountHandleStatus = &v
	return s
}

func (s *GetUserLatestPlanResponseBody) SetAccountHandleTime(v string) *GetUserLatestPlanResponseBody {
	s.AccountHandleTime = &v
	return s
}

func (s *GetUserLatestPlanResponseBody) SetAccountType(v int32) *GetUserLatestPlanResponseBody {
	s.AccountType = &v
	return s
}

func (s *GetUserLatestPlanResponseBody) SetAgreementFirstSignTime(v string) *GetUserLatestPlanResponseBody {
	s.AgreementFirstSignTime = &v
	return s
}

func (s *GetUserLatestPlanResponseBody) SetAgreementLastSignTime(v string) *GetUserLatestPlanResponseBody {
	s.AgreementLastSignTime = &v
	return s
}

func (s *GetUserLatestPlanResponseBody) SetAgreementStatus(v int32) *GetUserLatestPlanResponseBody {
	s.AgreementStatus = &v
	return s
}

func (s *GetUserLatestPlanResponseBody) SetDataHandleEndTime(v string) *GetUserLatestPlanResponseBody {
	s.DataHandleEndTime = &v
	return s
}

func (s *GetUserLatestPlanResponseBody) SetDataHandleStartTime(v string) *GetUserLatestPlanResponseBody {
	s.DataHandleStartTime = &v
	return s
}

func (s *GetUserLatestPlanResponseBody) SetDataHandleStatus(v int32) *GetUserLatestPlanResponseBody {
	s.DataHandleStatus = &v
	return s
}

func (s *GetUserLatestPlanResponseBody) SetExclusivePlan(v int32) *GetUserLatestPlanResponseBody {
	s.ExclusivePlan = &v
	return s
}

func (s *GetUserLatestPlanResponseBody) SetNewAccountUid(v int64) *GetUserLatestPlanResponseBody {
	s.NewAccountUid = &v
	return s
}

func (s *GetUserLatestPlanResponseBody) SetRequestId(v string) *GetUserLatestPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserLatestPlanResponseBody) SetStatus(v int32) *GetUserLatestPlanResponseBody {
	s.Status = &v
	return s
}

func (s *GetUserLatestPlanResponseBody) SetVendorRequestId(v string) *GetUserLatestPlanResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetUserLatestPlanResponseBody) SetVendorType(v string) *GetUserLatestPlanResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetUserLatestPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
