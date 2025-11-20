// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActionExecutor(v []*GetInstanceByIdResponseBodyActionExecutor) *GetInstanceByIdResponseBody
	GetActionExecutor() []*GetInstanceByIdResponseBodyActionExecutor
	SetApprovedResult(v string) *GetInstanceByIdResponseBody
	GetApprovedResult() *string
	SetCreateTimeGMT(v string) *GetInstanceByIdResponseBody
	GetCreateTimeGMT() *string
	SetData(v map[string]interface{}) *GetInstanceByIdResponseBody
	GetData() map[string]interface{}
	SetFormUuid(v string) *GetInstanceByIdResponseBody
	GetFormUuid() *string
	SetInstanceStatus(v string) *GetInstanceByIdResponseBody
	GetInstanceStatus() *string
	SetModifiedTimeGMT(v string) *GetInstanceByIdResponseBody
	GetModifiedTimeGMT() *string
	SetOriginator(v *GetInstanceByIdResponseBodyOriginator) *GetInstanceByIdResponseBody
	GetOriginator() *GetInstanceByIdResponseBodyOriginator
	SetProcessCode(v string) *GetInstanceByIdResponseBody
	GetProcessCode() *string
	SetProcessInstanceId(v string) *GetInstanceByIdResponseBody
	GetProcessInstanceId() *string
	SetRequestId(v string) *GetInstanceByIdResponseBody
	GetRequestId() *string
	SetTitle(v string) *GetInstanceByIdResponseBody
	GetTitle() *string
	SetVendorRequestId(v string) *GetInstanceByIdResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetInstanceByIdResponseBody
	GetVendorType() *string
	SetVersion(v int64) *GetInstanceByIdResponseBody
	GetVersion() *int64
}

type GetInstanceByIdResponseBody struct {
	ActionExecutor []*GetInstanceByIdResponseBodyActionExecutor `json:"actionExecutor,omitempty" xml:"actionExecutor,omitempty" type:"Repeated"`
	// example:
	//
	// agree
	ApprovedResult *string `json:"approvedResult,omitempty" xml:"approvedResult,omitempty"`
	// example:
	//
	// 2021-05-01
	CreateTimeGMT *string                `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	Data          map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// FORM-EF6Y4xxx
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	// example:
	//
	// 2021-05-01
	ModifiedTimeGMT *string                                `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	Originator      *GetInstanceByIdResponseBodyOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// example:
	//
	// TPROC--X1Gxxx
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// example:
	//
	// f30233fb-72e1-4xxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 李四发起的请购单
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetInstanceByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBody) GetActionExecutor() []*GetInstanceByIdResponseBodyActionExecutor {
	return s.ActionExecutor
}

func (s *GetInstanceByIdResponseBody) GetApprovedResult() *string {
	return s.ApprovedResult
}

func (s *GetInstanceByIdResponseBody) GetCreateTimeGMT() *string {
	return s.CreateTimeGMT
}

func (s *GetInstanceByIdResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetInstanceByIdResponseBody) GetFormUuid() *string {
	return s.FormUuid
}

func (s *GetInstanceByIdResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *GetInstanceByIdResponseBody) GetModifiedTimeGMT() *string {
	return s.ModifiedTimeGMT
}

func (s *GetInstanceByIdResponseBody) GetOriginator() *GetInstanceByIdResponseBodyOriginator {
	return s.Originator
}

func (s *GetInstanceByIdResponseBody) GetProcessCode() *string {
	return s.ProcessCode
}

func (s *GetInstanceByIdResponseBody) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetInstanceByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceByIdResponseBody) GetTitle() *string {
	return s.Title
}

func (s *GetInstanceByIdResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetInstanceByIdResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetInstanceByIdResponseBody) GetVersion() *int64 {
	return s.Version
}

func (s *GetInstanceByIdResponseBody) SetActionExecutor(v []*GetInstanceByIdResponseBodyActionExecutor) *GetInstanceByIdResponseBody {
	s.ActionExecutor = v
	return s
}

func (s *GetInstanceByIdResponseBody) SetApprovedResult(v string) *GetInstanceByIdResponseBody {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetCreateTimeGMT(v string) *GetInstanceByIdResponseBody {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetData(v map[string]interface{}) *GetInstanceByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceByIdResponseBody) SetFormUuid(v string) *GetInstanceByIdResponseBody {
	s.FormUuid = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetInstanceStatus(v string) *GetInstanceByIdResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetModifiedTimeGMT(v string) *GetInstanceByIdResponseBody {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetOriginator(v *GetInstanceByIdResponseBodyOriginator) *GetInstanceByIdResponseBody {
	s.Originator = v
	return s
}

func (s *GetInstanceByIdResponseBody) SetProcessCode(v string) *GetInstanceByIdResponseBody {
	s.ProcessCode = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetProcessInstanceId(v string) *GetInstanceByIdResponseBody {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetRequestId(v string) *GetInstanceByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetTitle(v string) *GetInstanceByIdResponseBody {
	s.Title = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetVendorRequestId(v string) *GetInstanceByIdResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetVendorType(v string) *GetInstanceByIdResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetVersion(v int64) *GetInstanceByIdResponseBody {
	s.Version = &v
	return s
}

func (s *GetInstanceByIdResponseBody) Validate() error {
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

type GetInstanceByIdResponseBodyActionExecutor struct {
	// example:
	//
	// 开发部
	DeptName *string `json:"DeptName,omitempty" xml:"DeptName,omitempty"`
	// example:
	//
	// abc@alimail.com
	Email *string                                        `json:"Email,omitempty" xml:"Email,omitempty"`
	Name  *GetInstanceByIdResponseBodyActionExecutorName `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	// example:
	//
	// manager123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetInstanceByIdResponseBodyActionExecutor) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceByIdResponseBodyActionExecutor) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBodyActionExecutor) GetDeptName() *string {
	return s.DeptName
}

func (s *GetInstanceByIdResponseBodyActionExecutor) GetEmail() *string {
	return s.Email
}

func (s *GetInstanceByIdResponseBodyActionExecutor) GetName() *GetInstanceByIdResponseBodyActionExecutorName {
	return s.Name
}

func (s *GetInstanceByIdResponseBodyActionExecutor) GetUserId() *string {
	return s.UserId
}

func (s *GetInstanceByIdResponseBodyActionExecutor) SetDeptName(v string) *GetInstanceByIdResponseBodyActionExecutor {
	s.DeptName = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutor) SetEmail(v string) *GetInstanceByIdResponseBodyActionExecutor {
	s.Email = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutor) SetName(v *GetInstanceByIdResponseBodyActionExecutorName) *GetInstanceByIdResponseBodyActionExecutor {
	s.Name = v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutor) SetUserId(v string) *GetInstanceByIdResponseBodyActionExecutor {
	s.UserId = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutor) Validate() error {
	if s.Name != nil {
		if err := s.Name.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceByIdResponseBodyActionExecutorName struct {
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

func (s GetInstanceByIdResponseBodyActionExecutorName) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceByIdResponseBodyActionExecutorName) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBodyActionExecutorName) GetNameInChinese() *string {
	return s.NameInChinese
}

func (s *GetInstanceByIdResponseBodyActionExecutorName) GetNameInEnglish() *string {
	return s.NameInEnglish
}

func (s *GetInstanceByIdResponseBodyActionExecutorName) GetType() *string {
	return s.Type
}

func (s *GetInstanceByIdResponseBodyActionExecutorName) SetNameInChinese(v string) *GetInstanceByIdResponseBodyActionExecutorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutorName) SetNameInEnglish(v string) *GetInstanceByIdResponseBodyActionExecutorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutorName) SetType(v string) *GetInstanceByIdResponseBodyActionExecutorName {
	s.Type = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutorName) Validate() error {
	return dara.Validate(s)
}

type GetInstanceByIdResponseBodyOriginator struct {
	// example:
	//
	// 开发部
	DeptName *string `json:"DeptName,omitempty" xml:"DeptName,omitempty"`
	// example:
	//
	// abc@alimail.com
	Email *string                                    `json:"Email,omitempty" xml:"Email,omitempty"`
	Name  *GetInstanceByIdResponseBodyOriginatorName `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	// example:
	//
	// manager123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetInstanceByIdResponseBodyOriginator) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceByIdResponseBodyOriginator) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBodyOriginator) GetDeptName() *string {
	return s.DeptName
}

func (s *GetInstanceByIdResponseBodyOriginator) GetEmail() *string {
	return s.Email
}

func (s *GetInstanceByIdResponseBodyOriginator) GetName() *GetInstanceByIdResponseBodyOriginatorName {
	return s.Name
}

func (s *GetInstanceByIdResponseBodyOriginator) GetUserId() *string {
	return s.UserId
}

func (s *GetInstanceByIdResponseBodyOriginator) SetDeptName(v string) *GetInstanceByIdResponseBodyOriginator {
	s.DeptName = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginator) SetEmail(v string) *GetInstanceByIdResponseBodyOriginator {
	s.Email = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginator) SetName(v *GetInstanceByIdResponseBodyOriginatorName) *GetInstanceByIdResponseBodyOriginator {
	s.Name = v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginator) SetUserId(v string) *GetInstanceByIdResponseBodyOriginator {
	s.UserId = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginator) Validate() error {
	if s.Name != nil {
		if err := s.Name.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceByIdResponseBodyOriginatorName struct {
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

func (s GetInstanceByIdResponseBodyOriginatorName) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceByIdResponseBodyOriginatorName) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBodyOriginatorName) GetNameInChinese() *string {
	return s.NameInChinese
}

func (s *GetInstanceByIdResponseBodyOriginatorName) GetNameInEnglish() *string {
	return s.NameInEnglish
}

func (s *GetInstanceByIdResponseBodyOriginatorName) GetType() *string {
	return s.Type
}

func (s *GetInstanceByIdResponseBodyOriginatorName) SetNameInChinese(v string) *GetInstanceByIdResponseBodyOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginatorName) SetNameInEnglish(v string) *GetInstanceByIdResponseBodyOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginatorName) SetType(v string) *GetInstanceByIdResponseBodyOriginatorName {
	s.Type = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginatorName) Validate() error {
	return dara.Validate(s)
}
