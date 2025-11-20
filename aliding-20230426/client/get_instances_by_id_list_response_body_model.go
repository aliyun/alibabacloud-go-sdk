// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstancesByIdListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetInstancesByIdListResponseBody
	GetRequestId() *string
	SetResult(v []*GetInstancesByIdListResponseBodyResult) *GetInstancesByIdListResponseBody
	GetResult() []*GetInstancesByIdListResponseBodyResult
	SetVendorRequestId(v string) *GetInstancesByIdListResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetInstancesByIdListResponseBody
	GetVendorType() *string
}

type GetInstancesByIdListResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*GetInstancesByIdListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetInstancesByIdListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesByIdListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstancesByIdListResponseBody) GetResult() []*GetInstancesByIdListResponseBodyResult {
	return s.Result
}

func (s *GetInstancesByIdListResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetInstancesByIdListResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetInstancesByIdListResponseBody) SetRequestId(v string) *GetInstancesByIdListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstancesByIdListResponseBody) SetResult(v []*GetInstancesByIdListResponseBodyResult) *GetInstancesByIdListResponseBody {
	s.Result = v
	return s
}

func (s *GetInstancesByIdListResponseBody) SetVendorRequestId(v string) *GetInstancesByIdListResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetInstancesByIdListResponseBody) SetVendorType(v string) *GetInstancesByIdListResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetInstancesByIdListResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstancesByIdListResponseBodyResult struct {
	ActionExecutor []*GetInstancesByIdListResponseBodyResultActionExecutor `json:"ActionExecutor,omitempty" xml:"ActionExecutor,omitempty" type:"Repeated"`
	// example:
	//
	// agree
	ApprovedResult *string                `json:"ApprovedResult,omitempty" xml:"ApprovedResult,omitempty"`
	Data           map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// FORM-EF6Y4G8WO2FN0SUB43TDQ3CGC3FMFQ1G9400RCJ3
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// RUNNING
	InstanceStatus *string                                           `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	Originator     *GetInstancesByIdListResponseBodyResultOriginator `json:"Originator,omitempty" xml:"Originator,omitempty" type:"Struct"`
	// example:
	//
	// TPROC--X1Gxxx
	ProcessCode *string `json:"ProcessCode,omitempty" xml:"ProcessCode,omitempty"`
	// example:
	//
	// f30233fb-72e1-4xxx
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	// example:
	//
	// 李四发起的请购单
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetInstancesByIdListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesByIdListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListResponseBodyResult) GetActionExecutor() []*GetInstancesByIdListResponseBodyResultActionExecutor {
	return s.ActionExecutor
}

func (s *GetInstancesByIdListResponseBodyResult) GetApprovedResult() *string {
	return s.ApprovedResult
}

func (s *GetInstancesByIdListResponseBodyResult) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetInstancesByIdListResponseBodyResult) GetFormUuid() *string {
	return s.FormUuid
}

func (s *GetInstancesByIdListResponseBodyResult) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *GetInstancesByIdListResponseBodyResult) GetOriginator() *GetInstancesByIdListResponseBodyResultOriginator {
	return s.Originator
}

func (s *GetInstancesByIdListResponseBodyResult) GetProcessCode() *string {
	return s.ProcessCode
}

func (s *GetInstancesByIdListResponseBodyResult) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetInstancesByIdListResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *GetInstancesByIdListResponseBodyResult) SetActionExecutor(v []*GetInstancesByIdListResponseBodyResultActionExecutor) *GetInstancesByIdListResponseBodyResult {
	s.ActionExecutor = v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) SetApprovedResult(v string) *GetInstancesByIdListResponseBodyResult {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) SetData(v map[string]interface{}) *GetInstancesByIdListResponseBodyResult {
	s.Data = v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) SetFormUuid(v string) *GetInstancesByIdListResponseBodyResult {
	s.FormUuid = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) SetInstanceStatus(v string) *GetInstancesByIdListResponseBodyResult {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) SetOriginator(v *GetInstancesByIdListResponseBodyResultOriginator) *GetInstancesByIdListResponseBodyResult {
	s.Originator = v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) SetProcessCode(v string) *GetInstancesByIdListResponseBodyResult {
	s.ProcessCode = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) SetProcessInstanceId(v string) *GetInstancesByIdListResponseBodyResult {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) SetTitle(v string) *GetInstancesByIdListResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) Validate() error {
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

type GetInstancesByIdListResponseBodyResultActionExecutor struct {
	// example:
	//
	// 开发部
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// example:
	//
	// abc@alimail.com
	Email *string                                                   `json:"Email,omitempty" xml:"Email,omitempty"`
	Name  *GetInstancesByIdListResponseBodyResultActionExecutorName `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	// example:
	//
	// manager123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetInstancesByIdListResponseBodyResultActionExecutor) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesByIdListResponseBodyResultActionExecutor) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutor) GetDepartmentName() *string {
	return s.DepartmentName
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutor) GetEmail() *string {
	return s.Email
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutor) GetName() *GetInstancesByIdListResponseBodyResultActionExecutorName {
	return s.Name
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutor) GetUserId() *string {
	return s.UserId
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutor) SetDepartmentName(v string) *GetInstancesByIdListResponseBodyResultActionExecutor {
	s.DepartmentName = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutor) SetEmail(v string) *GetInstancesByIdListResponseBodyResultActionExecutor {
	s.Email = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutor) SetName(v *GetInstancesByIdListResponseBodyResultActionExecutorName) *GetInstancesByIdListResponseBodyResultActionExecutor {
	s.Name = v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutor) SetUserId(v string) *GetInstancesByIdListResponseBodyResultActionExecutor {
	s.UserId = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutor) Validate() error {
	if s.Name != nil {
		if err := s.Name.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstancesByIdListResponseBodyResultActionExecutorName struct {
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

func (s GetInstancesByIdListResponseBodyResultActionExecutorName) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesByIdListResponseBodyResultActionExecutorName) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutorName) GetNameInChinese() *string {
	return s.NameInChinese
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutorName) GetNameInEnglish() *string {
	return s.NameInEnglish
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutorName) GetType() *string {
	return s.Type
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutorName) SetNameInChinese(v string) *GetInstancesByIdListResponseBodyResultActionExecutorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutorName) SetNameInEnglish(v string) *GetInstancesByIdListResponseBodyResultActionExecutorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutorName) SetType(v string) *GetInstancesByIdListResponseBodyResultActionExecutorName {
	s.Type = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutorName) Validate() error {
	return dara.Validate(s)
}

type GetInstancesByIdListResponseBodyResultOriginator struct {
	// example:
	//
	// 开发部
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// example:
	//
	// abc@alimail.com
	Email *string                                               `json:"Email,omitempty" xml:"Email,omitempty"`
	Name  *GetInstancesByIdListResponseBodyResultOriginatorName `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	// example:
	//
	// manager123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetInstancesByIdListResponseBodyResultOriginator) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesByIdListResponseBodyResultOriginator) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListResponseBodyResultOriginator) GetDepartmentName() *string {
	return s.DepartmentName
}

func (s *GetInstancesByIdListResponseBodyResultOriginator) GetEmail() *string {
	return s.Email
}

func (s *GetInstancesByIdListResponseBodyResultOriginator) GetName() *GetInstancesByIdListResponseBodyResultOriginatorName {
	return s.Name
}

func (s *GetInstancesByIdListResponseBodyResultOriginator) GetUserId() *string {
	return s.UserId
}

func (s *GetInstancesByIdListResponseBodyResultOriginator) SetDepartmentName(v string) *GetInstancesByIdListResponseBodyResultOriginator {
	s.DepartmentName = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultOriginator) SetEmail(v string) *GetInstancesByIdListResponseBodyResultOriginator {
	s.Email = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultOriginator) SetName(v *GetInstancesByIdListResponseBodyResultOriginatorName) *GetInstancesByIdListResponseBodyResultOriginator {
	s.Name = v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultOriginator) SetUserId(v string) *GetInstancesByIdListResponseBodyResultOriginator {
	s.UserId = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultOriginator) Validate() error {
	if s.Name != nil {
		if err := s.Name.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstancesByIdListResponseBodyResultOriginatorName struct {
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

func (s GetInstancesByIdListResponseBodyResultOriginatorName) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesByIdListResponseBodyResultOriginatorName) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListResponseBodyResultOriginatorName) GetNameInChinese() *string {
	return s.NameInChinese
}

func (s *GetInstancesByIdListResponseBodyResultOriginatorName) GetNameInEnglish() *string {
	return s.NameInEnglish
}

func (s *GetInstancesByIdListResponseBodyResultOriginatorName) GetType() *string {
	return s.Type
}

func (s *GetInstancesByIdListResponseBodyResultOriginatorName) SetNameInChinese(v string) *GetInstancesByIdListResponseBodyResultOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultOriginatorName) SetNameInEnglish(v string) *GetInstancesByIdListResponseBodyResultOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultOriginatorName) SetType(v string) *GetInstancesByIdListResponseBodyResultOriginatorName {
	s.Type = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultOriginatorName) Validate() error {
	return dara.Validate(s)
}
