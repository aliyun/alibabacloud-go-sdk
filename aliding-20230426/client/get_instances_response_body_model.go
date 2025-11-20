// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetInstancesResponseBodyData) *GetInstancesResponseBody
	GetData() []*GetInstancesResponseBodyData
	SetPageNumber(v int64) *GetInstancesResponseBody
	GetPageNumber() *int64
	SetRequestId(v string) *GetInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetInstancesResponseBody
	GetTotalCount() *int64
	SetVendorRequestId(v string) *GetInstancesResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetInstancesResponseBody
	GetVendorType() *string
}

type GetInstancesResponseBody struct {
	Data []*GetInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBody) GetData() []*GetInstancesResponseBodyData {
	return s.Data
}

func (s *GetInstancesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetInstancesResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetInstancesResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetInstancesResponseBody) SetData(v []*GetInstancesResponseBodyData) *GetInstancesResponseBody {
	s.Data = v
	return s
}

func (s *GetInstancesResponseBody) SetPageNumber(v int64) *GetInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetInstancesResponseBody) SetRequestId(v string) *GetInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstancesResponseBody) SetTotalCount(v int64) *GetInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetInstancesResponseBody) SetVendorRequestId(v string) *GetInstancesResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetInstancesResponseBody) SetVendorType(v string) *GetInstancesResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetInstancesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstancesResponseBodyData struct {
	ActionExecutor []*GetInstancesResponseBodyDataActionExecutor `json:"ActionExecutor,omitempty" xml:"ActionExecutor,omitempty" type:"Repeated"`
	// example:
	//
	// agree
	ApprovedResult *string `json:"ApprovedResult,omitempty" xml:"ApprovedResult,omitempty"`
	// example:
	//
	// 2021-05-01
	CreateTimeGMT *string                `json:"CreateTimeGMT,omitempty" xml:"CreateTimeGMT,omitempty"`
	Data          map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// FORM-EF6Y4G8WO2FN0SUB43TDQ3CGC3FMFQ1G9400RCJ3
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// example:
	//
	// 2021-05-01
	ModifiedTimeGMT *string                                 `json:"ModifiedTimeGMT,omitempty" xml:"ModifiedTimeGMT,omitempty"`
	Originator      *GetInstancesResponseBodyDataOriginator `json:"Originator,omitempty" xml:"Originator,omitempty" type:"Struct"`
	// example:
	//
	// TPROC--X1Gxxx
	ProcessCode *string `json:"ProcessCode,omitempty" xml:"ProcessCode,omitempty"`
	// example:
	//
	// f30233fb-72e1-4af4-8cb8-c7e0ea9ee530
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	// example:
	//
	// 小红发起的请购单
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 1.0
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyData) GetActionExecutor() []*GetInstancesResponseBodyDataActionExecutor {
	return s.ActionExecutor
}

func (s *GetInstancesResponseBodyData) GetApprovedResult() *string {
	return s.ApprovedResult
}

func (s *GetInstancesResponseBodyData) GetCreateTimeGMT() *string {
	return s.CreateTimeGMT
}

func (s *GetInstancesResponseBodyData) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetInstancesResponseBodyData) GetFormUuid() *string {
	return s.FormUuid
}

func (s *GetInstancesResponseBodyData) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *GetInstancesResponseBodyData) GetModifiedTimeGMT() *string {
	return s.ModifiedTimeGMT
}

func (s *GetInstancesResponseBodyData) GetOriginator() *GetInstancesResponseBodyDataOriginator {
	return s.Originator
}

func (s *GetInstancesResponseBodyData) GetProcessCode() *string {
	return s.ProcessCode
}

func (s *GetInstancesResponseBodyData) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetInstancesResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetInstancesResponseBodyData) GetVersion() *int64 {
	return s.Version
}

func (s *GetInstancesResponseBodyData) SetActionExecutor(v []*GetInstancesResponseBodyDataActionExecutor) *GetInstancesResponseBodyData {
	s.ActionExecutor = v
	return s
}

func (s *GetInstancesResponseBodyData) SetApprovedResult(v string) *GetInstancesResponseBodyData {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetCreateTimeGMT(v string) *GetInstancesResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetData(v map[string]interface{}) *GetInstancesResponseBodyData {
	s.Data = v
	return s
}

func (s *GetInstancesResponseBodyData) SetFormUuid(v string) *GetInstancesResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetInstanceStatus(v string) *GetInstancesResponseBodyData {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetModifiedTimeGMT(v string) *GetInstancesResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetOriginator(v *GetInstancesResponseBodyDataOriginator) *GetInstancesResponseBodyData {
	s.Originator = v
	return s
}

func (s *GetInstancesResponseBodyData) SetProcessCode(v string) *GetInstancesResponseBodyData {
	s.ProcessCode = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetProcessInstanceId(v string) *GetInstancesResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetTitle(v string) *GetInstancesResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetVersion(v int64) *GetInstancesResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetInstancesResponseBodyData) Validate() error {
	if s.ActionExecutor != nil {
		for _, item := range s.ActionExecutor {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Originator != nil {
		if err := s.Originator.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstancesResponseBodyDataActionExecutor struct {
	// example:
	//
	// 开发部
	DeptName *string `json:"DeptName,omitempty" xml:"DeptName,omitempty"`
	// example:
	//
	// abc@alimail.com
	Email *string                                         `json:"Email,omitempty" xml:"Email,omitempty"`
	Name  *GetInstancesResponseBodyDataActionExecutorName `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	// example:
	//
	// manager123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetInstancesResponseBodyDataActionExecutor) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesResponseBodyDataActionExecutor) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyDataActionExecutor) GetDeptName() *string {
	return s.DeptName
}

func (s *GetInstancesResponseBodyDataActionExecutor) GetEmail() *string {
	return s.Email
}

func (s *GetInstancesResponseBodyDataActionExecutor) GetName() *GetInstancesResponseBodyDataActionExecutorName {
	return s.Name
}

func (s *GetInstancesResponseBodyDataActionExecutor) GetUserId() *string {
	return s.UserId
}

func (s *GetInstancesResponseBodyDataActionExecutor) SetDeptName(v string) *GetInstancesResponseBodyDataActionExecutor {
	s.DeptName = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutor) SetEmail(v string) *GetInstancesResponseBodyDataActionExecutor {
	s.Email = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutor) SetName(v *GetInstancesResponseBodyDataActionExecutorName) *GetInstancesResponseBodyDataActionExecutor {
	s.Name = v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutor) SetUserId(v string) *GetInstancesResponseBodyDataActionExecutor {
	s.UserId = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutor) Validate() error {
	if s.Name != nil {
		if err := s.Name.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstancesResponseBodyDataActionExecutorName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"NameInChinese,omitempty" xml:"NameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"NameInEnglish,omitempty" xml:"NameInEnglish,omitempty"`
	// example:
	//
	// i18n
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstancesResponseBodyDataActionExecutorName) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesResponseBodyDataActionExecutorName) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyDataActionExecutorName) GetNameInChinese() *string {
	return s.NameInChinese
}

func (s *GetInstancesResponseBodyDataActionExecutorName) GetNameInEnglish() *string {
	return s.NameInEnglish
}

func (s *GetInstancesResponseBodyDataActionExecutorName) GetType() *string {
	return s.Type
}

func (s *GetInstancesResponseBodyDataActionExecutorName) SetNameInChinese(v string) *GetInstancesResponseBodyDataActionExecutorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutorName) SetNameInEnglish(v string) *GetInstancesResponseBodyDataActionExecutorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutorName) SetType(v string) *GetInstancesResponseBodyDataActionExecutorName {
	s.Type = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutorName) Validate() error {
	return dara.Validate(s)
}

type GetInstancesResponseBodyDataOriginator struct {
	// example:
	//
	// 开发部
	DeptName *string `json:"DeptName,omitempty" xml:"DeptName,omitempty"`
	// example:
	//
	// abc@alimail.com
	Email *string                                     `json:"Email,omitempty" xml:"Email,omitempty"`
	Name  *GetInstancesResponseBodyDataOriginatorName `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	// example:
	//
	// manager123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetInstancesResponseBodyDataOriginator) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyDataOriginator) GetDeptName() *string {
	return s.DeptName
}

func (s *GetInstancesResponseBodyDataOriginator) GetEmail() *string {
	return s.Email
}

func (s *GetInstancesResponseBodyDataOriginator) GetName() *GetInstancesResponseBodyDataOriginatorName {
	return s.Name
}

func (s *GetInstancesResponseBodyDataOriginator) GetUserId() *string {
	return s.UserId
}

func (s *GetInstancesResponseBodyDataOriginator) SetDeptName(v string) *GetInstancesResponseBodyDataOriginator {
	s.DeptName = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginator) SetEmail(v string) *GetInstancesResponseBodyDataOriginator {
	s.Email = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginator) SetName(v *GetInstancesResponseBodyDataOriginatorName) *GetInstancesResponseBodyDataOriginator {
	s.Name = v
	return s
}

func (s *GetInstancesResponseBodyDataOriginator) SetUserId(v string) *GetInstancesResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginator) Validate() error {
	if s.Name != nil {
		if err := s.Name.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstancesResponseBodyDataOriginatorName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"NameInChinese,omitempty" xml:"NameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"NameInEnglish,omitempty" xml:"NameInEnglish,omitempty"`
	// example:
	//
	// i18n
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstancesResponseBodyDataOriginatorName) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesResponseBodyDataOriginatorName) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyDataOriginatorName) GetNameInChinese() *string {
	return s.NameInChinese
}

func (s *GetInstancesResponseBodyDataOriginatorName) GetNameInEnglish() *string {
	return s.NameInEnglish
}

func (s *GetInstancesResponseBodyDataOriginatorName) GetType() *string {
	return s.Type
}

func (s *GetInstancesResponseBodyDataOriginatorName) SetNameInChinese(v string) *GetInstancesResponseBodyDataOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginatorName) SetNameInEnglish(v string) *GetInstancesResponseBodyDataOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginatorName) SetType(v string) *GetInstancesResponseBodyDataOriginatorName {
	s.Type = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginatorName) Validate() error {
	return dara.Validate(s)
}
