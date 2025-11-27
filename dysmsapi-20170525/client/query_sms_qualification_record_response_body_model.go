// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsQualificationRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuerySmsQualificationRecordResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QuerySmsQualificationRecordResponseBody
	GetCode() *string
	SetData(v *QuerySmsQualificationRecordResponseBodyData) *QuerySmsQualificationRecordResponseBody
	GetData() *QuerySmsQualificationRecordResponseBodyData
	SetMessage(v string) *QuerySmsQualificationRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySmsQualificationRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySmsQualificationRecordResponseBody
	GetSuccess() *bool
}

type QuerySmsQualificationRecordResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QuerySmsQualificationRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 25D5AFDE-8EBC-132E-8909-1FDC071DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySmsQualificationRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsQualificationRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsQualificationRecordResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuerySmsQualificationRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySmsQualificationRecordResponseBody) GetData() *QuerySmsQualificationRecordResponseBodyData {
	return s.Data
}

func (s *QuerySmsQualificationRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySmsQualificationRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySmsQualificationRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySmsQualificationRecordResponseBody) SetAccessDeniedDetail(v string) *QuerySmsQualificationRecordResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBody) SetCode(v string) *QuerySmsQualificationRecordResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBody) SetData(v *QuerySmsQualificationRecordResponseBodyData) *QuerySmsQualificationRecordResponseBody {
	s.Data = v
	return s
}

func (s *QuerySmsQualificationRecordResponseBody) SetMessage(v string) *QuerySmsQualificationRecordResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBody) SetRequestId(v string) *QuerySmsQualificationRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBody) SetSuccess(v bool) *QuerySmsQualificationRecordResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySmsQualificationRecordResponseBodyData struct {
	List []*QuerySmsQualificationRecordResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 25
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QuerySmsQualificationRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsQualificationRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySmsQualificationRecordResponseBodyData) GetList() []*QuerySmsQualificationRecordResponseBodyDataList {
	return s.List
}

func (s *QuerySmsQualificationRecordResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QuerySmsQualificationRecordResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QuerySmsQualificationRecordResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *QuerySmsQualificationRecordResponseBodyData) SetList(v []*QuerySmsQualificationRecordResponseBodyDataList) *QuerySmsQualificationRecordResponseBodyData {
	s.List = v
	return s
}

func (s *QuerySmsQualificationRecordResponseBodyData) SetPageNo(v int64) *QuerySmsQualificationRecordResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBodyData) SetPageSize(v int64) *QuerySmsQualificationRecordResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBodyData) SetTotal(v int64) *QuerySmsQualificationRecordResponseBodyData {
	s.Total = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySmsQualificationRecordResponseBodyDataList struct {
	// 审核备注
	//
	// example:
	//
	// 示例值示例值
	AuditRemark *string `json:"AuditRemark,omitempty" xml:"AuditRemark,omitempty"`
	// 审核时间
	//
	// example:
	//
	// 2024-12-26 17:29:04
	AuditTime *string `json:"AuditTime,omitempty" xml:"AuditTime,omitempty"`
	// 公司名称或实人认证姓名
	//
	// example:
	//
	// 示例值示例值示例值
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 2025-02-20 11:59:30
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// 资质组ID
	//
	// example:
	//
	// 10000****
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 法人名称
	//
	// example:
	//
	// 示例值示例值
	LegalPersonName *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
	// 资质组名称
	//
	// example:
	//
	// 示例值示例值示例值
	QualificationGroupName *string `json:"QualificationGroupName,omitempty" xml:"QualificationGroupName,omitempty"`
	// 审核状态名
	//
	// example:
	//
	// INT
	StateName *string `json:"StateName,omitempty" xml:"StateName,omitempty"`
	// 是否自用
	//
	// example:
	//
	// true
	UseBySelf *string `json:"UseBySelf,omitempty" xml:"UseBySelf,omitempty"`
	// 工单ID
	//
	// example:
	//
	// 2001****
	WorkOrderId *int64 `json:"WorkOrderId,omitempty" xml:"WorkOrderId,omitempty"`
}

func (s QuerySmsQualificationRecordResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsQualificationRecordResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) GetAuditRemark() *string {
	return s.AuditRemark
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) GetAuditTime() *string {
	return s.AuditTime
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) GetCompanyName() *string {
	return s.CompanyName
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) GetCreateDate() *string {
	return s.CreateDate
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) GetGroupId() *int64 {
	return s.GroupId
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) GetLegalPersonName() *string {
	return s.LegalPersonName
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) GetQualificationGroupName() *string {
	return s.QualificationGroupName
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) GetStateName() *string {
	return s.StateName
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) GetUseBySelf() *string {
	return s.UseBySelf
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) GetWorkOrderId() *int64 {
	return s.WorkOrderId
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) SetAuditRemark(v string) *QuerySmsQualificationRecordResponseBodyDataList {
	s.AuditRemark = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) SetAuditTime(v string) *QuerySmsQualificationRecordResponseBodyDataList {
	s.AuditTime = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) SetCompanyName(v string) *QuerySmsQualificationRecordResponseBodyDataList {
	s.CompanyName = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) SetCreateDate(v string) *QuerySmsQualificationRecordResponseBodyDataList {
	s.CreateDate = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) SetGroupId(v int64) *QuerySmsQualificationRecordResponseBodyDataList {
	s.GroupId = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) SetLegalPersonName(v string) *QuerySmsQualificationRecordResponseBodyDataList {
	s.LegalPersonName = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) SetQualificationGroupName(v string) *QuerySmsQualificationRecordResponseBodyDataList {
	s.QualificationGroupName = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) SetStateName(v string) *QuerySmsQualificationRecordResponseBodyDataList {
	s.StateName = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) SetUseBySelf(v string) *QuerySmsQualificationRecordResponseBodyDataList {
	s.UseBySelf = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) SetWorkOrderId(v int64) *QuerySmsQualificationRecordResponseBodyDataList {
	s.WorkOrderId = &v
	return s
}

func (s *QuerySmsQualificationRecordResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
