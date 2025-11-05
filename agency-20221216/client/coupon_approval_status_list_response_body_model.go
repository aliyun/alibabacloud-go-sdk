// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCouponApprovalStatusListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CouponApprovalStatusListResponseBody
	GetCode() *string
	SetData(v []*CouponApprovalStatusListResponseBodyData) *CouponApprovalStatusListResponseBody
	GetData() []*CouponApprovalStatusListResponseBodyData
	SetMessage(v string) *CouponApprovalStatusListResponseBody
	GetMessage() *string
	SetPageNo(v int32) *CouponApprovalStatusListResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *CouponApprovalStatusListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *CouponApprovalStatusListResponseBody
	GetRequestId() *string
	SetTotal(v int32) *CouponApprovalStatusListResponseBody
	GetTotal() *int32
}

type CouponApprovalStatusListResponseBody struct {
	// example:
	//
	// 200
	Code    *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    []*CouponApprovalStatusListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s CouponApprovalStatusListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CouponApprovalStatusListResponseBody) GoString() string {
	return s.String()
}

func (s *CouponApprovalStatusListResponseBody) GetCode() *string {
	return s.Code
}

func (s *CouponApprovalStatusListResponseBody) GetData() []*CouponApprovalStatusListResponseBodyData {
	return s.Data
}

func (s *CouponApprovalStatusListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CouponApprovalStatusListResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *CouponApprovalStatusListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CouponApprovalStatusListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CouponApprovalStatusListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *CouponApprovalStatusListResponseBody) SetCode(v string) *CouponApprovalStatusListResponseBody {
	s.Code = &v
	return s
}

func (s *CouponApprovalStatusListResponseBody) SetData(v []*CouponApprovalStatusListResponseBodyData) *CouponApprovalStatusListResponseBody {
	s.Data = v
	return s
}

func (s *CouponApprovalStatusListResponseBody) SetMessage(v string) *CouponApprovalStatusListResponseBody {
	s.Message = &v
	return s
}

func (s *CouponApprovalStatusListResponseBody) SetPageNo(v int32) *CouponApprovalStatusListResponseBody {
	s.PageNo = &v
	return s
}

func (s *CouponApprovalStatusListResponseBody) SetPageSize(v int32) *CouponApprovalStatusListResponseBody {
	s.PageSize = &v
	return s
}

func (s *CouponApprovalStatusListResponseBody) SetRequestId(v string) *CouponApprovalStatusListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CouponApprovalStatusListResponseBody) SetTotal(v int32) *CouponApprovalStatusListResponseBody {
	s.Total = &v
	return s
}

func (s *CouponApprovalStatusListResponseBody) Validate() error {
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

type CouponApprovalStatusListResponseBodyData struct {
	// example:
	//
	// test@test.aliyunid.com
	IssuerAccount *string `json:"IssuerAccount,omitempty" xml:"IssuerAccount,omitempty"`
	// example:
	//
	// 5432738203821334
	IssuerUid *string `json:"IssuerUid,omitempty" xml:"IssuerUid,omitempty"`
	Note      *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// example:
	//
	// S00000101-100040
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 2
	TemplateStatus *int64 `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// example:
	//
	// 2024-02-02 09:46:59
	TimeOfRequest *string `json:"TimeOfRequest,omitempty" xml:"TimeOfRequest,omitempty"`
}

func (s CouponApprovalStatusListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CouponApprovalStatusListResponseBodyData) GoString() string {
	return s.String()
}

func (s *CouponApprovalStatusListResponseBodyData) GetIssuerAccount() *string {
	return s.IssuerAccount
}

func (s *CouponApprovalStatusListResponseBodyData) GetIssuerUid() *string {
	return s.IssuerUid
}

func (s *CouponApprovalStatusListResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *CouponApprovalStatusListResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CouponApprovalStatusListResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CouponApprovalStatusListResponseBodyData) GetTemplateStatus() *int64 {
	return s.TemplateStatus
}

func (s *CouponApprovalStatusListResponseBodyData) GetTimeOfRequest() *string {
	return s.TimeOfRequest
}

func (s *CouponApprovalStatusListResponseBodyData) SetIssuerAccount(v string) *CouponApprovalStatusListResponseBodyData {
	s.IssuerAccount = &v
	return s
}

func (s *CouponApprovalStatusListResponseBodyData) SetIssuerUid(v string) *CouponApprovalStatusListResponseBodyData {
	s.IssuerUid = &v
	return s
}

func (s *CouponApprovalStatusListResponseBodyData) SetNote(v string) *CouponApprovalStatusListResponseBodyData {
	s.Note = &v
	return s
}

func (s *CouponApprovalStatusListResponseBodyData) SetTemplateId(v string) *CouponApprovalStatusListResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *CouponApprovalStatusListResponseBodyData) SetTemplateName(v string) *CouponApprovalStatusListResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *CouponApprovalStatusListResponseBodyData) SetTemplateStatus(v int64) *CouponApprovalStatusListResponseBodyData {
	s.TemplateStatus = &v
	return s
}

func (s *CouponApprovalStatusListResponseBodyData) SetTimeOfRequest(v string) *CouponApprovalStatusListResponseBodyData {
	s.TimeOfRequest = &v
	return s
}

func (s *CouponApprovalStatusListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
