// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContactTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactTemplates(v *QueryContactTemplateResponseBodyContactTemplates) *QueryContactTemplateResponseBody
	GetContactTemplates() *QueryContactTemplateResponseBodyContactTemplates
	SetCurrentPageNum(v int32) *QueryContactTemplateResponseBody
	GetCurrentPageNum() *int32
	SetNextPage(v bool) *QueryContactTemplateResponseBody
	GetNextPage() *bool
	SetPageSize(v int32) *QueryContactTemplateResponseBody
	GetPageSize() *int32
	SetPrePage(v bool) *QueryContactTemplateResponseBody
	GetPrePage() *bool
	SetRequestId(v string) *QueryContactTemplateResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *QueryContactTemplateResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *QueryContactTemplateResponseBody
	GetTotalPageNum() *int32
}

type QueryContactTemplateResponseBody struct {
	ContactTemplates *QueryContactTemplateResponseBodyContactTemplates `json:"ContactTemplates,omitempty" xml:"ContactTemplates,omitempty" type:"Struct"`
	CurrentPageNum   *int32                                            `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	NextPage         *bool                                             `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize         *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage          *bool                                             `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	RequestId        *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum     *int32                                            `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum     *int32                                            `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryContactTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryContactTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryContactTemplateResponseBody) GetContactTemplates() *QueryContactTemplateResponseBodyContactTemplates {
	return s.ContactTemplates
}

func (s *QueryContactTemplateResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryContactTemplateResponseBody) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryContactTemplateResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryContactTemplateResponseBody) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryContactTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryContactTemplateResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryContactTemplateResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryContactTemplateResponseBody) SetContactTemplates(v *QueryContactTemplateResponseBodyContactTemplates) *QueryContactTemplateResponseBody {
	s.ContactTemplates = v
	return s
}

func (s *QueryContactTemplateResponseBody) SetCurrentPageNum(v int32) *QueryContactTemplateResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryContactTemplateResponseBody) SetNextPage(v bool) *QueryContactTemplateResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryContactTemplateResponseBody) SetPageSize(v int32) *QueryContactTemplateResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryContactTemplateResponseBody) SetPrePage(v bool) *QueryContactTemplateResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryContactTemplateResponseBody) SetRequestId(v string) *QueryContactTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryContactTemplateResponseBody) SetTotalItemNum(v int32) *QueryContactTemplateResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryContactTemplateResponseBody) SetTotalPageNum(v int32) *QueryContactTemplateResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryContactTemplateResponseBody) Validate() error {
	if s.ContactTemplates != nil {
		if err := s.ContactTemplates.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryContactTemplateResponseBodyContactTemplates struct {
	ContactTemplate []*QueryContactTemplateResponseBodyContactTemplatesContactTemplate `json:"ContactTemplate,omitempty" xml:"ContactTemplate,omitempty" type:"Repeated"`
}

func (s QueryContactTemplateResponseBodyContactTemplates) String() string {
	return dara.Prettify(s)
}

func (s QueryContactTemplateResponseBodyContactTemplates) GoString() string {
	return s.String()
}

func (s *QueryContactTemplateResponseBodyContactTemplates) GetContactTemplate() []*QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	return s.ContactTemplate
}

func (s *QueryContactTemplateResponseBodyContactTemplates) SetContactTemplate(v []*QueryContactTemplateResponseBodyContactTemplatesContactTemplate) *QueryContactTemplateResponseBodyContactTemplates {
	s.ContactTemplate = v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplates) Validate() error {
	if s.ContactTemplate != nil {
		for _, item := range s.ContactTemplate {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryContactTemplateResponseBodyContactTemplatesContactTemplate struct {
	AuditStatus             *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	CCity                   *string `json:"CCity,omitempty" xml:"CCity,omitempty"`
	CCompany                *string `json:"CCompany,omitempty" xml:"CCompany,omitempty"`
	CCountry                *string `json:"CCountry,omitempty" xml:"CCountry,omitempty"`
	CName                   *string `json:"CName,omitempty" xml:"CName,omitempty"`
	CProvince               *string `json:"CProvince,omitempty" xml:"CProvince,omitempty"`
	CVenu                   *string `json:"CVenu,omitempty" xml:"CVenu,omitempty"`
	ContactTemplateId       *int64  `json:"ContactTemplateId,omitempty" xml:"ContactTemplateId,omitempty"`
	CreateTime              *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DefaultTemplate         *bool   `json:"DefaultTemplate,omitempty" xml:"DefaultTemplate,omitempty"`
	ECity                   *string `json:"ECity,omitempty" xml:"ECity,omitempty"`
	ECompany                *string `json:"ECompany,omitempty" xml:"ECompany,omitempty"`
	EName                   *string `json:"EName,omitempty" xml:"EName,omitempty"`
	EProvince               *string `json:"EProvince,omitempty" xml:"EProvince,omitempty"`
	EVenu                   *string `json:"EVenu,omitempty" xml:"EVenu,omitempty"`
	Email                   *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EmailVerificationStatus *int32  `json:"EmailVerificationStatus,omitempty" xml:"EmailVerificationStatus,omitempty"`
	PostalCode              *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	RegType                 *string `json:"RegType,omitempty" xml:"RegType,omitempty"`
	TelArea                 *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt                  *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	TelMain                 *string `json:"TelMain,omitempty" xml:"TelMain,omitempty"`
	UpdateTime              *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId                  *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryContactTemplateResponseBodyContactTemplatesContactTemplate) String() string {
	return dara.Prettify(s)
}

func (s QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GoString() string {
	return s.String()
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetCCity() *string {
	return s.CCity
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetCCompany() *string {
	return s.CCompany
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetCCountry() *string {
	return s.CCountry
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetCName() *string {
	return s.CName
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetCProvince() *string {
	return s.CProvince
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetCVenu() *string {
	return s.CVenu
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetContactTemplateId() *int64 {
	return s.ContactTemplateId
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetDefaultTemplate() *bool {
	return s.DefaultTemplate
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetECity() *string {
	return s.ECity
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetECompany() *string {
	return s.ECompany
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetEName() *string {
	return s.EName
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetEProvince() *string {
	return s.EProvince
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetEVenu() *string {
	return s.EVenu
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetEmail() *string {
	return s.Email
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetEmailVerificationStatus() *int32 {
	return s.EmailVerificationStatus
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetPostalCode() *string {
	return s.PostalCode
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetRegType() *string {
	return s.RegType
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetTelArea() *string {
	return s.TelArea
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetTelExt() *string {
	return s.TelExt
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetTelMain() *string {
	return s.TelMain
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) GetUserId() *string {
	return s.UserId
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetAuditStatus(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.AuditStatus = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetCCity(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.CCity = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetCCompany(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.CCompany = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetCCountry(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.CCountry = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetCName(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.CName = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetCProvince(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.CProvince = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetCVenu(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.CVenu = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetContactTemplateId(v int64) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.ContactTemplateId = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetCreateTime(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.CreateTime = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetDefaultTemplate(v bool) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.DefaultTemplate = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetECity(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.ECity = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetECompany(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.ECompany = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetEName(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.EName = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetEProvince(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.EProvince = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetEVenu(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.EVenu = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetEmail(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.Email = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetEmailVerificationStatus(v int32) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.EmailVerificationStatus = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetPostalCode(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.PostalCode = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetRegType(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.RegType = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetTelArea(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.TelArea = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetTelExt(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.TelExt = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetTelMain(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.TelMain = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetUpdateTime(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.UpdateTime = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) SetUserId(v string) *QueryContactTemplateResponseBodyContactTemplatesContactTemplate {
	s.UserId = &v
	return s
}

func (s *QueryContactTemplateResponseBodyContactTemplatesContactTemplate) Validate() error {
	return dara.Validate(s)
}
