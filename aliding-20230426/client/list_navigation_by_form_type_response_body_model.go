// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNavigationByFormTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListNavigationByFormTypeResponseBody
	GetRequestId() *string
	SetResult(v []*ListNavigationByFormTypeResponseBodyResult) *ListNavigationByFormTypeResponseBody
	GetResult() []*ListNavigationByFormTypeResponseBodyResult
	SetVendorRequestId(v string) *ListNavigationByFormTypeResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *ListNavigationByFormTypeResponseBody
	GetVendorType() *string
}

type ListNavigationByFormTypeResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListNavigationByFormTypeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ListNavigationByFormTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNavigationByFormTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListNavigationByFormTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNavigationByFormTypeResponseBody) GetResult() []*ListNavigationByFormTypeResponseBodyResult {
	return s.Result
}

func (s *ListNavigationByFormTypeResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *ListNavigationByFormTypeResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *ListNavigationByFormTypeResponseBody) SetRequestId(v string) *ListNavigationByFormTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNavigationByFormTypeResponseBody) SetResult(v []*ListNavigationByFormTypeResponseBodyResult) *ListNavigationByFormTypeResponseBody {
	s.Result = v
	return s
}

func (s *ListNavigationByFormTypeResponseBody) SetVendorRequestId(v string) *ListNavigationByFormTypeResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *ListNavigationByFormTypeResponseBody) SetVendorType(v string) *ListNavigationByFormTypeResponseBody {
	s.VendorType = &v
	return s
}

func (s *ListNavigationByFormTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNavigationByFormTypeResponseBodyResult struct {
	// example:
	//
	// FORM-EF6Yxxx
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// TPROC--X1Gxxx
	ProcessCode *string                                          `json:"ProcessCode,omitempty" xml:"ProcessCode,omitempty"`
	Title       *ListNavigationByFormTypeResponseBodyResultTitle `json:"Title,omitempty" xml:"Title,omitempty" type:"Struct"`
}

func (s ListNavigationByFormTypeResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListNavigationByFormTypeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListNavigationByFormTypeResponseBodyResult) GetFormUuid() *string {
	return s.FormUuid
}

func (s *ListNavigationByFormTypeResponseBodyResult) GetProcessCode() *string {
	return s.ProcessCode
}

func (s *ListNavigationByFormTypeResponseBodyResult) GetTitle() *ListNavigationByFormTypeResponseBodyResultTitle {
	return s.Title
}

func (s *ListNavigationByFormTypeResponseBodyResult) SetFormUuid(v string) *ListNavigationByFormTypeResponseBodyResult {
	s.FormUuid = &v
	return s
}

func (s *ListNavigationByFormTypeResponseBodyResult) SetProcessCode(v string) *ListNavigationByFormTypeResponseBodyResult {
	s.ProcessCode = &v
	return s
}

func (s *ListNavigationByFormTypeResponseBodyResult) SetTitle(v *ListNavigationByFormTypeResponseBodyResultTitle) *ListNavigationByFormTypeResponseBodyResult {
	s.Title = v
	return s
}

func (s *ListNavigationByFormTypeResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListNavigationByFormTypeResponseBodyResultTitle struct {
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
	// 未知
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNavigationByFormTypeResponseBodyResultTitle) String() string {
	return dara.Prettify(s)
}

func (s ListNavigationByFormTypeResponseBodyResultTitle) GoString() string {
	return s.String()
}

func (s *ListNavigationByFormTypeResponseBodyResultTitle) GetNameInChinese() *string {
	return s.NameInChinese
}

func (s *ListNavigationByFormTypeResponseBodyResultTitle) GetNameInEnglish() *string {
	return s.NameInEnglish
}

func (s *ListNavigationByFormTypeResponseBodyResultTitle) GetType() *string {
	return s.Type
}

func (s *ListNavigationByFormTypeResponseBodyResultTitle) SetNameInChinese(v string) *ListNavigationByFormTypeResponseBodyResultTitle {
	s.NameInChinese = &v
	return s
}

func (s *ListNavigationByFormTypeResponseBodyResultTitle) SetNameInEnglish(v string) *ListNavigationByFormTypeResponseBodyResultTitle {
	s.NameInEnglish = &v
	return s
}

func (s *ListNavigationByFormTypeResponseBodyResultTitle) SetType(v string) *ListNavigationByFormTypeResponseBodyResultTitle {
	s.Type = &v
	return s
}

func (s *ListNavigationByFormTypeResponseBodyResultTitle) Validate() error {
	return dara.Validate(s)
}
