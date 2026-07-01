// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsQualificationRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyName(v string) *QuerySmsQualificationRecordRequest
	GetCompanyName() *string
	SetLegalPersonName(v string) *QuerySmsQualificationRecordRequest
	GetLegalPersonName() *string
	SetOwnerId(v int64) *QuerySmsQualificationRecordRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *QuerySmsQualificationRecordRequest
	GetPageNo() *int64
	SetPageSize(v int64) *QuerySmsQualificationRecordRequest
	GetPageSize() *int64
	SetQualificationGroupName(v string) *QuerySmsQualificationRecordRequest
	GetQualificationGroupName() *string
	SetResourceOwnerAccount(v string) *QuerySmsQualificationRecordRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySmsQualificationRecordRequest
	GetResourceOwnerId() *int64
	SetState(v string) *QuerySmsQualificationRecordRequest
	GetState() *string
	SetUseBySelf(v bool) *QuerySmsQualificationRecordRequest
	GetUseBySelf() *bool
	SetWorkOrderId(v int64) *QuerySmsQualificationRecordRequest
	GetWorkOrderId() *int64
}

type QuerySmsQualificationRecordRequest struct {
	// The company name.
	//
	// example:
	//
	// 阿里云云通信有限公司
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// The name of the legal representative.
	//
	// example:
	//
	// 李华
	LegalPersonName *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: **1 to 50**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The qualification name.
	//
	// example:
	//
	// 阿里云云通信有限公司李华
	QualificationGroupName *string `json:"QualificationGroupName,omitempty" xml:"QualificationGroupName,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The review status. Valid values:
	//
	// - INIT: Under review.
	//
	// - NOT_PASS: Review rejected.
	//
	// - PASS: Review approved.
	//
	// - NOT_FINISH: Additional information required.
	//
	// - CANCEL: Withdrawn.
	//
	// example:
	//
	// PASS
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The purpose of the qualification application. Valid values:
	//
	// - **true**: For self-use.
	//
	// - **false**: For use by others.
	//
	// example:
	//
	// true
	UseBySelf *bool `json:"UseBySelf,omitempty" xml:"UseBySelf,omitempty"`
	// The review ticket ID.
	//
	// example:
	//
	// 2001****
	WorkOrderId *int64 `json:"WorkOrderId,omitempty" xml:"WorkOrderId,omitempty"`
}

func (s QuerySmsQualificationRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsQualificationRecordRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsQualificationRecordRequest) GetCompanyName() *string {
	return s.CompanyName
}

func (s *QuerySmsQualificationRecordRequest) GetLegalPersonName() *string {
	return s.LegalPersonName
}

func (s *QuerySmsQualificationRecordRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySmsQualificationRecordRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QuerySmsQualificationRecordRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QuerySmsQualificationRecordRequest) GetQualificationGroupName() *string {
	return s.QualificationGroupName
}

func (s *QuerySmsQualificationRecordRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySmsQualificationRecordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySmsQualificationRecordRequest) GetState() *string {
	return s.State
}

func (s *QuerySmsQualificationRecordRequest) GetUseBySelf() *bool {
	return s.UseBySelf
}

func (s *QuerySmsQualificationRecordRequest) GetWorkOrderId() *int64 {
	return s.WorkOrderId
}

func (s *QuerySmsQualificationRecordRequest) SetCompanyName(v string) *QuerySmsQualificationRecordRequest {
	s.CompanyName = &v
	return s
}

func (s *QuerySmsQualificationRecordRequest) SetLegalPersonName(v string) *QuerySmsQualificationRecordRequest {
	s.LegalPersonName = &v
	return s
}

func (s *QuerySmsQualificationRecordRequest) SetOwnerId(v int64) *QuerySmsQualificationRecordRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsQualificationRecordRequest) SetPageNo(v int64) *QuerySmsQualificationRecordRequest {
	s.PageNo = &v
	return s
}

func (s *QuerySmsQualificationRecordRequest) SetPageSize(v int64) *QuerySmsQualificationRecordRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySmsQualificationRecordRequest) SetQualificationGroupName(v string) *QuerySmsQualificationRecordRequest {
	s.QualificationGroupName = &v
	return s
}

func (s *QuerySmsQualificationRecordRequest) SetResourceOwnerAccount(v string) *QuerySmsQualificationRecordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsQualificationRecordRequest) SetResourceOwnerId(v int64) *QuerySmsQualificationRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmsQualificationRecordRequest) SetState(v string) *QuerySmsQualificationRecordRequest {
	s.State = &v
	return s
}

func (s *QuerySmsQualificationRecordRequest) SetUseBySelf(v bool) *QuerySmsQualificationRecordRequest {
	s.UseBySelf = &v
	return s
}

func (s *QuerySmsQualificationRecordRequest) SetWorkOrderId(v int64) *QuerySmsQualificationRecordRequest {
	s.WorkOrderId = &v
	return s
}

func (s *QuerySmsQualificationRecordRequest) Validate() error {
	return dara.Validate(s)
}
