// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormDataByIDResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFormData(v map[string]interface{}) *GetFormDataByIDResponseBody
	GetFormData() map[string]interface{}
	SetFormInstId(v string) *GetFormDataByIDResponseBody
	GetFormInstId() *string
	SetModifiedTimeGMT(v string) *GetFormDataByIDResponseBody
	GetModifiedTimeGMT() *string
	SetOriginator(v *GetFormDataByIDResponseBodyOriginator) *GetFormDataByIDResponseBody
	GetOriginator() *GetFormDataByIDResponseBodyOriginator
	SetRequestId(v string) *GetFormDataByIDResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GetFormDataByIDResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetFormDataByIDResponseBody
	GetVendorType() *string
}

type GetFormDataByIDResponseBody struct {
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
	// example:
	//
	// FORM_INST_12345
	FormInstId *string `json:"formInstId,omitempty" xml:"formInstId,omitempty"`
	// example:
	//
	// 2021-05-01
	ModifiedTimeGMT *string                                `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	Originator      *GetFormDataByIDResponseBodyOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetFormDataByIDResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFormDataByIDResponseBody) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBody) GetFormData() map[string]interface{} {
	return s.FormData
}

func (s *GetFormDataByIDResponseBody) GetFormInstId() *string {
	return s.FormInstId
}

func (s *GetFormDataByIDResponseBody) GetModifiedTimeGMT() *string {
	return s.ModifiedTimeGMT
}

func (s *GetFormDataByIDResponseBody) GetOriginator() *GetFormDataByIDResponseBodyOriginator {
	return s.Originator
}

func (s *GetFormDataByIDResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFormDataByIDResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetFormDataByIDResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetFormDataByIDResponseBody) SetFormData(v map[string]interface{}) *GetFormDataByIDResponseBody {
	s.FormData = v
	return s
}

func (s *GetFormDataByIDResponseBody) SetFormInstId(v string) *GetFormDataByIDResponseBody {
	s.FormInstId = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetModifiedTimeGMT(v string) *GetFormDataByIDResponseBody {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetOriginator(v *GetFormDataByIDResponseBodyOriginator) *GetFormDataByIDResponseBody {
	s.Originator = v
	return s
}

func (s *GetFormDataByIDResponseBody) SetRequestId(v string) *GetFormDataByIDResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetVendorRequestId(v string) *GetFormDataByIDResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetVendorType(v string) *GetFormDataByIDResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetFormDataByIDResponseBody) Validate() error {
	if s.Originator != nil {
		if err := s.Originator.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFormDataByIDResponseBodyOriginator struct {
	// example:
	//
	// 开发部
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// example:
	//
	// abc@alimail.com
	Email *string                                    `json:"Email,omitempty" xml:"Email,omitempty"`
	Name  *GetFormDataByIDResponseBodyOriginatorName `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetFormDataByIDResponseBodyOriginator) String() string {
	return dara.Prettify(s)
}

func (s GetFormDataByIDResponseBodyOriginator) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBodyOriginator) GetDepartmentName() *string {
	return s.DepartmentName
}

func (s *GetFormDataByIDResponseBodyOriginator) GetEmail() *string {
	return s.Email
}

func (s *GetFormDataByIDResponseBodyOriginator) GetName() *GetFormDataByIDResponseBodyOriginatorName {
	return s.Name
}

func (s *GetFormDataByIDResponseBodyOriginator) GetUserId() *string {
	return s.UserId
}

func (s *GetFormDataByIDResponseBodyOriginator) SetDepartmentName(v string) *GetFormDataByIDResponseBodyOriginator {
	s.DepartmentName = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetEmail(v string) *GetFormDataByIDResponseBodyOriginator {
	s.Email = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetName(v *GetFormDataByIDResponseBodyOriginatorName) *GetFormDataByIDResponseBodyOriginator {
	s.Name = v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetUserId(v string) *GetFormDataByIDResponseBodyOriginator {
	s.UserId = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) Validate() error {
	if s.Name != nil {
		if err := s.Name.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFormDataByIDResponseBodyOriginatorName struct {
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

func (s GetFormDataByIDResponseBodyOriginatorName) String() string {
	return dara.Prettify(s)
}

func (s GetFormDataByIDResponseBodyOriginatorName) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBodyOriginatorName) GetNameInChinese() *string {
	return s.NameInChinese
}

func (s *GetFormDataByIDResponseBodyOriginatorName) GetNameInEnglish() *string {
	return s.NameInEnglish
}

func (s *GetFormDataByIDResponseBodyOriginatorName) GetType() *string {
	return s.Type
}

func (s *GetFormDataByIDResponseBodyOriginatorName) SetNameInChinese(v string) *GetFormDataByIDResponseBodyOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginatorName) SetNameInEnglish(v string) *GetFormDataByIDResponseBodyOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginatorName) SetType(v string) *GetFormDataByIDResponseBodyOriginatorName {
	s.Type = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginatorName) Validate() error {
	return dara.Validate(s)
}
