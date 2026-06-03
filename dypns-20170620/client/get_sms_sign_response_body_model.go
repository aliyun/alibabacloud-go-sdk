// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSmsSignResponseBody
	GetCode() *string
	SetData(v []*GetSmsSignResponseBodyData) *GetSmsSignResponseBody
	GetData() []*GetSmsSignResponseBodyData
	SetMessage(v string) *GetSmsSignResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSmsSignResponseBody
	GetRequestId() *string
}

type GetSmsSignResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetSmsSignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSmsSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmsSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSmsSignResponseBody) GetData() []*GetSmsSignResponseBodyData {
	return s.Data
}

func (s *GetSmsSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSmsSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSmsSignResponseBody) SetCode(v string) *GetSmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *GetSmsSignResponseBody) SetData(v []*GetSmsSignResponseBodyData) *GetSmsSignResponseBody {
	s.Data = v
	return s
}

func (s *GetSmsSignResponseBody) SetMessage(v string) *GetSmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *GetSmsSignResponseBody) SetRequestId(v string) *GetSmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmsSignResponseBody) Validate() error {
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

type GetSmsSignResponseBodyData struct {
	AppUrl                 *string `json:"AppUrl,omitempty" xml:"AppUrl,omitempty"`
	AuditResult            *string `json:"AuditResult,omitempty" xml:"AuditResult,omitempty"`
	BusinessLicenseCert    *string `json:"BusinessLicenseCert,omitempty" xml:"BusinessLicenseCert,omitempty"`
	BusinessLicenseCertId  *string `json:"BusinessLicenseCertId,omitempty" xml:"BusinessLicenseCertId,omitempty"`
	CertType               *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CreateDate             *int64  `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DefaultFlag            *bool   `json:"DefaultFlag,omitempty" xml:"DefaultFlag,omitempty"`
	OrderId                *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrganizationCodeCert   *string `json:"OrganizationCodeCert,omitempty" xml:"OrganizationCodeCert,omitempty"`
	OrganizationCodeCertId *string `json:"OrganizationCodeCertId,omitempty" xml:"OrganizationCodeCertId,omitempty"`
	SmsSignName            *string `json:"SmsSignName,omitempty" xml:"SmsSignName,omitempty"`
	SmsSignRemark          *string `json:"SmsSignRemark,omitempty" xml:"SmsSignRemark,omitempty"`
	SmsSignSource          *string `json:"SmsSignSource,omitempty" xml:"SmsSignSource,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaxRegistrationCert    *string `json:"TaxRegistrationCert,omitempty" xml:"TaxRegistrationCert,omitempty"`
	TaxRegistrationCertId  *string `json:"TaxRegistrationCertId,omitempty" xml:"TaxRegistrationCertId,omitempty"`
	TestFlag               *bool   `json:"TestFlag,omitempty" xml:"TestFlag,omitempty"`
}

func (s GetSmsSignResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSmsSignResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSmsSignResponseBodyData) GetAppUrl() *string {
	return s.AppUrl
}

func (s *GetSmsSignResponseBodyData) GetAuditResult() *string {
	return s.AuditResult
}

func (s *GetSmsSignResponseBodyData) GetBusinessLicenseCert() *string {
	return s.BusinessLicenseCert
}

func (s *GetSmsSignResponseBodyData) GetBusinessLicenseCertId() *string {
	return s.BusinessLicenseCertId
}

func (s *GetSmsSignResponseBodyData) GetCertType() *string {
	return s.CertType
}

func (s *GetSmsSignResponseBodyData) GetCreateDate() *int64 {
	return s.CreateDate
}

func (s *GetSmsSignResponseBodyData) GetDefaultFlag() *bool {
	return s.DefaultFlag
}

func (s *GetSmsSignResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *GetSmsSignResponseBodyData) GetOrganizationCodeCert() *string {
	return s.OrganizationCodeCert
}

func (s *GetSmsSignResponseBodyData) GetOrganizationCodeCertId() *string {
	return s.OrganizationCodeCertId
}

func (s *GetSmsSignResponseBodyData) GetSmsSignName() *string {
	return s.SmsSignName
}

func (s *GetSmsSignResponseBodyData) GetSmsSignRemark() *string {
	return s.SmsSignRemark
}

func (s *GetSmsSignResponseBodyData) GetSmsSignSource() *string {
	return s.SmsSignSource
}

func (s *GetSmsSignResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetSmsSignResponseBodyData) GetTaxRegistrationCert() *string {
	return s.TaxRegistrationCert
}

func (s *GetSmsSignResponseBodyData) GetTaxRegistrationCertId() *string {
	return s.TaxRegistrationCertId
}

func (s *GetSmsSignResponseBodyData) GetTestFlag() *bool {
	return s.TestFlag
}

func (s *GetSmsSignResponseBodyData) SetAppUrl(v string) *GetSmsSignResponseBodyData {
	s.AppUrl = &v
	return s
}

func (s *GetSmsSignResponseBodyData) SetAuditResult(v string) *GetSmsSignResponseBodyData {
	s.AuditResult = &v
	return s
}

func (s *GetSmsSignResponseBodyData) SetBusinessLicenseCert(v string) *GetSmsSignResponseBodyData {
	s.BusinessLicenseCert = &v
	return s
}

func (s *GetSmsSignResponseBodyData) SetBusinessLicenseCertId(v string) *GetSmsSignResponseBodyData {
	s.BusinessLicenseCertId = &v
	return s
}

func (s *GetSmsSignResponseBodyData) SetCertType(v string) *GetSmsSignResponseBodyData {
	s.CertType = &v
	return s
}

func (s *GetSmsSignResponseBodyData) SetCreateDate(v int64) *GetSmsSignResponseBodyData {
	s.CreateDate = &v
	return s
}

func (s *GetSmsSignResponseBodyData) SetDefaultFlag(v bool) *GetSmsSignResponseBodyData {
	s.DefaultFlag = &v
	return s
}

func (s *GetSmsSignResponseBodyData) SetOrderId(v string) *GetSmsSignResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *GetSmsSignResponseBodyData) SetOrganizationCodeCert(v string) *GetSmsSignResponseBodyData {
	s.OrganizationCodeCert = &v
	return s
}

func (s *GetSmsSignResponseBodyData) SetOrganizationCodeCertId(v string) *GetSmsSignResponseBodyData {
	s.OrganizationCodeCertId = &v
	return s
}

func (s *GetSmsSignResponseBodyData) SetSmsSignName(v string) *GetSmsSignResponseBodyData {
	s.SmsSignName = &v
	return s
}

func (s *GetSmsSignResponseBodyData) SetSmsSignRemark(v string) *GetSmsSignResponseBodyData {
	s.SmsSignRemark = &v
	return s
}

func (s *GetSmsSignResponseBodyData) SetSmsSignSource(v string) *GetSmsSignResponseBodyData {
	s.SmsSignSource = &v
	return s
}

func (s *GetSmsSignResponseBodyData) SetStatus(v string) *GetSmsSignResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSmsSignResponseBodyData) SetTaxRegistrationCert(v string) *GetSmsSignResponseBodyData {
	s.TaxRegistrationCert = &v
	return s
}

func (s *GetSmsSignResponseBodyData) SetTaxRegistrationCertId(v string) *GetSmsSignResponseBodyData {
	s.TaxRegistrationCertId = &v
	return s
}

func (s *GetSmsSignResponseBodyData) SetTestFlag(v bool) *GetSmsSignResponseBodyData {
	s.TestFlag = &v
	return s
}

func (s *GetSmsSignResponseBodyData) Validate() error {
	return dara.Validate(s)
}
