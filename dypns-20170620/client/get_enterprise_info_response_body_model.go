// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEnterpriseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEnterpriseInfoResponseBody
	GetCode() *string
	SetData(v *GetEnterpriseInfoResponseBodyData) *GetEnterpriseInfoResponseBody
	GetData() *GetEnterpriseInfoResponseBodyData
	SetMessage(v string) *GetEnterpriseInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEnterpriseInfoResponseBody
	GetRequestId() *string
}

type GetEnterpriseInfoResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetEnterpriseInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEnterpriseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnterpriseInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEnterpriseInfoResponseBody) GetData() *GetEnterpriseInfoResponseBodyData {
	return s.Data
}

func (s *GetEnterpriseInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEnterpriseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEnterpriseInfoResponseBody) SetCode(v string) *GetEnterpriseInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetEnterpriseInfoResponseBody) SetData(v *GetEnterpriseInfoResponseBodyData) *GetEnterpriseInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetEnterpriseInfoResponseBody) SetMessage(v string) *GetEnterpriseInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetEnterpriseInfoResponseBody) SetRequestId(v string) *GetEnterpriseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEnterpriseInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEnterpriseInfoResponseBodyData struct {
	AuditDesc                      *string `json:"AuditDesc,omitempty" xml:"AuditDesc,omitempty"`
	AuditTime                      *string `json:"AuditTime,omitempty" xml:"AuditTime,omitempty"`
	BusinessLicenseAddress         *string `json:"BusinessLicenseAddress,omitempty" xml:"BusinessLicenseAddress,omitempty"`
	BusinessLicensePicture         *string `json:"BusinessLicensePicture,omitempty" xml:"BusinessLicensePicture,omitempty"`
	EnterpriseId                   *int64  `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	EnterpriseName                 *string `json:"EnterpriseName,omitempty" xml:"EnterpriseName,omitempty"`
	LegalPersonCertNumber          *string `json:"LegalPersonCertNumber,omitempty" xml:"LegalPersonCertNumber,omitempty"`
	LegalPersonCertPicture         *string `json:"LegalPersonCertPicture,omitempty" xml:"LegalPersonCertPicture,omitempty"`
	LegalPersonName                *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
	ManagerCertNumber              *string `json:"ManagerCertNumber,omitempty" xml:"ManagerCertNumber,omitempty"`
	ManagerCertPicture             *string `json:"ManagerCertPicture,omitempty" xml:"ManagerCertPicture,omitempty"`
	ManagerContactNumber           *string `json:"ManagerContactNumber,omitempty" xml:"ManagerContactNumber,omitempty"`
	ManagerName                    *string `json:"ManagerName,omitempty" xml:"ManagerName,omitempty"`
	NumberApplicationLetterPicture *string `json:"NumberApplicationLetterPicture,omitempty" xml:"NumberApplicationLetterPicture,omitempty"`
	OrderId                        *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrganizationCode               *string `json:"OrganizationCode,omitempty" xml:"OrganizationCode,omitempty"`
	Remark                         *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status                         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UndertakingPicture             *string `json:"UndertakingPicture,omitempty" xml:"UndertakingPicture,omitempty"`
}

func (s GetEnterpriseInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnterpriseInfoResponseBodyData) GetAuditDesc() *string {
	return s.AuditDesc
}

func (s *GetEnterpriseInfoResponseBodyData) GetAuditTime() *string {
	return s.AuditTime
}

func (s *GetEnterpriseInfoResponseBodyData) GetBusinessLicenseAddress() *string {
	return s.BusinessLicenseAddress
}

func (s *GetEnterpriseInfoResponseBodyData) GetBusinessLicensePicture() *string {
	return s.BusinessLicensePicture
}

func (s *GetEnterpriseInfoResponseBodyData) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *GetEnterpriseInfoResponseBodyData) GetEnterpriseName() *string {
	return s.EnterpriseName
}

func (s *GetEnterpriseInfoResponseBodyData) GetLegalPersonCertNumber() *string {
	return s.LegalPersonCertNumber
}

func (s *GetEnterpriseInfoResponseBodyData) GetLegalPersonCertPicture() *string {
	return s.LegalPersonCertPicture
}

func (s *GetEnterpriseInfoResponseBodyData) GetLegalPersonName() *string {
	return s.LegalPersonName
}

func (s *GetEnterpriseInfoResponseBodyData) GetManagerCertNumber() *string {
	return s.ManagerCertNumber
}

func (s *GetEnterpriseInfoResponseBodyData) GetManagerCertPicture() *string {
	return s.ManagerCertPicture
}

func (s *GetEnterpriseInfoResponseBodyData) GetManagerContactNumber() *string {
	return s.ManagerContactNumber
}

func (s *GetEnterpriseInfoResponseBodyData) GetManagerName() *string {
	return s.ManagerName
}

func (s *GetEnterpriseInfoResponseBodyData) GetNumberApplicationLetterPicture() *string {
	return s.NumberApplicationLetterPicture
}

func (s *GetEnterpriseInfoResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetEnterpriseInfoResponseBodyData) GetOrganizationCode() *string {
	return s.OrganizationCode
}

func (s *GetEnterpriseInfoResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *GetEnterpriseInfoResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetEnterpriseInfoResponseBodyData) GetUndertakingPicture() *string {
	return s.UndertakingPicture
}

func (s *GetEnterpriseInfoResponseBodyData) SetAuditDesc(v string) *GetEnterpriseInfoResponseBodyData {
	s.AuditDesc = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetAuditTime(v string) *GetEnterpriseInfoResponseBodyData {
	s.AuditTime = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetBusinessLicenseAddress(v string) *GetEnterpriseInfoResponseBodyData {
	s.BusinessLicenseAddress = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetBusinessLicensePicture(v string) *GetEnterpriseInfoResponseBodyData {
	s.BusinessLicensePicture = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetEnterpriseId(v int64) *GetEnterpriseInfoResponseBodyData {
	s.EnterpriseId = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetEnterpriseName(v string) *GetEnterpriseInfoResponseBodyData {
	s.EnterpriseName = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetLegalPersonCertNumber(v string) *GetEnterpriseInfoResponseBodyData {
	s.LegalPersonCertNumber = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetLegalPersonCertPicture(v string) *GetEnterpriseInfoResponseBodyData {
	s.LegalPersonCertPicture = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetLegalPersonName(v string) *GetEnterpriseInfoResponseBodyData {
	s.LegalPersonName = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetManagerCertNumber(v string) *GetEnterpriseInfoResponseBodyData {
	s.ManagerCertNumber = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetManagerCertPicture(v string) *GetEnterpriseInfoResponseBodyData {
	s.ManagerCertPicture = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetManagerContactNumber(v string) *GetEnterpriseInfoResponseBodyData {
	s.ManagerContactNumber = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetManagerName(v string) *GetEnterpriseInfoResponseBodyData {
	s.ManagerName = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetNumberApplicationLetterPicture(v string) *GetEnterpriseInfoResponseBodyData {
	s.NumberApplicationLetterPicture = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetOrderId(v int64) *GetEnterpriseInfoResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetOrganizationCode(v string) *GetEnterpriseInfoResponseBodyData {
	s.OrganizationCode = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetRemark(v string) *GetEnterpriseInfoResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetStatus(v int32) *GetEnterpriseInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) SetUndertakingPicture(v string) *GetEnterpriseInfoResponseBodyData {
	s.UndertakingPicture = &v
	return s
}

func (s *GetEnterpriseInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
