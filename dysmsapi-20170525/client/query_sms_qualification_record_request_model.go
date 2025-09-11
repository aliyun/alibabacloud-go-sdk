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
	// 公司名
	//
	// example:
	//
	// 示例值
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// 法人姓名
	//
	// example:
	//
	// 示例值示例值
	LegalPersonName *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 资质组名称
	//
	// example:
	//
	// 示例值示例值
	QualificationGroupName *string `json:"QualificationGroupName,omitempty" xml:"QualificationGroupName,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 审核状态。INT:审核中FAILED:审核失败,PASSED:审核通过,NOT_FINISH:资料待补充,CANCELED:已撤回
	//
	// example:
	//
	// PASSED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// 是否自用
	//
	// example:
	//
	// true
	UseBySelf *bool `json:"UseBySelf,omitempty" xml:"UseBySelf,omitempty"`
	// 工单ID
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
