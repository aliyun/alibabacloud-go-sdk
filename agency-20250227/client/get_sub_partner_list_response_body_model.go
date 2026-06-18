// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubPartnerListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *GetSubPartnerListResponseBody
	GetMessage() *string
	SetPageNo(v string) *GetSubPartnerListResponseBody
	GetPageNo() *string
	SetPageSize(v string) *GetSubPartnerListResponseBody
	GetPageSize() *string
	SetRequestId(v string) *GetSubPartnerListResponseBody
	GetRequestId() *string
	SetSubPartnerList(v []*GetSubPartnerListResponseBodySubPartnerList) *GetSubPartnerListResponseBody
	GetSubPartnerList() []*GetSubPartnerListResponseBodySubPartnerList
	SetSuccess(v bool) *GetSubPartnerListResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *GetSubPartnerListResponseBody
	GetTotal() *int32
}

type GetSubPartnerListResponseBody struct {
	// The message returned.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of secondary distributors returned per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of secondary distributors.
	SubPartnerList []*GetSubPartnerListResponseBodySubPartnerList `json:"SubPartnerList,omitempty" xml:"SubPartnerList,omitempty" type:"Repeated"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries that match the query conditions.
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSubPartnerListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubPartnerListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubPartnerListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSubPartnerListResponseBody) GetPageNo() *string {
	return s.PageNo
}

func (s *GetSubPartnerListResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *GetSubPartnerListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSubPartnerListResponseBody) GetSubPartnerList() []*GetSubPartnerListResponseBodySubPartnerList {
	return s.SubPartnerList
}

func (s *GetSubPartnerListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSubPartnerListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetSubPartnerListResponseBody) SetMessage(v string) *GetSubPartnerListResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubPartnerListResponseBody) SetPageNo(v string) *GetSubPartnerListResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetSubPartnerListResponseBody) SetPageSize(v string) *GetSubPartnerListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetSubPartnerListResponseBody) SetRequestId(v string) *GetSubPartnerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubPartnerListResponseBody) SetSubPartnerList(v []*GetSubPartnerListResponseBodySubPartnerList) *GetSubPartnerListResponseBody {
	s.SubPartnerList = v
	return s
}

func (s *GetSubPartnerListResponseBody) SetSuccess(v bool) *GetSubPartnerListResponseBody {
	s.Success = &v
	return s
}

func (s *GetSubPartnerListResponseBody) SetTotal(v int32) *GetSubPartnerListResponseBody {
	s.Total = &v
	return s
}

func (s *GetSubPartnerListResponseBody) Validate() error {
	if s.SubPartnerList != nil {
		for _, item := range s.SubPartnerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSubPartnerListResponseBodySubPartnerList struct {
	// The detailed registered address.
	//
	// example:
	//
	// xxx街道xxx号
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The agreement status code.
	//
	// example:
	//
	// ACTIVE_PERIOD
	AgreementStatus *string `json:"AgreementStatus,omitempty" xml:"AgreementStatus,omitempty"`
	// The description of the agreement status.
	//
	// example:
	//
	// 生效期
	AgreementStatusDesc *string `json:"AgreementStatusDesc,omitempty" xml:"AgreementStatusDesc,omitempty"`
	// The city of the registered address.
	//
	// example:
	//
	// 杭州市
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The name of the secondary distributor.
	//
	// example:
	//
	// xxx有限公司
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// The name of the contact.
	//
	// example:
	//
	// 张三
	Contact *string `json:"Contact,omitempty" xml:"Contact,omitempty"`
	// The district or county of the registered address.
	//
	// example:
	//
	// 西湖区
	District *string `json:"District,omitempty" xml:"District,omitempty"`
	// The time when the secondary distributor first joined.
	//
	// example:
	//
	// 2000-01-01
	JoinTime *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	// The name of the primary account of the secondary distributor.
	//
	// example:
	//
	// xxxxxx有限公司主账号
	MasterAccount *string `json:"MasterAccount,omitempty" xml:"MasterAccount,omitempty"`
	// The UID of the primary account of the secondary distributor.
	//
	// example:
	//
	// 123456
	MasterUid *string `json:"MasterUid,omitempty" xml:"MasterUid,omitempty"`
	// The PID of the secondary distributor.
	//
	// example:
	//
	// P123456
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The province of the registered address.
	//
	// example:
	//
	// 浙江省
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s GetSubPartnerListResponseBodySubPartnerList) String() string {
	return dara.Prettify(s)
}

func (s GetSubPartnerListResponseBodySubPartnerList) GoString() string {
	return s.String()
}

func (s *GetSubPartnerListResponseBodySubPartnerList) GetAddress() *string {
	return s.Address
}

func (s *GetSubPartnerListResponseBodySubPartnerList) GetAgreementStatus() *string {
	return s.AgreementStatus
}

func (s *GetSubPartnerListResponseBodySubPartnerList) GetAgreementStatusDesc() *string {
	return s.AgreementStatusDesc
}

func (s *GetSubPartnerListResponseBodySubPartnerList) GetCity() *string {
	return s.City
}

func (s *GetSubPartnerListResponseBodySubPartnerList) GetCompanyName() *string {
	return s.CompanyName
}

func (s *GetSubPartnerListResponseBodySubPartnerList) GetContact() *string {
	return s.Contact
}

func (s *GetSubPartnerListResponseBodySubPartnerList) GetDistrict() *string {
	return s.District
}

func (s *GetSubPartnerListResponseBodySubPartnerList) GetJoinTime() *string {
	return s.JoinTime
}

func (s *GetSubPartnerListResponseBodySubPartnerList) GetMasterAccount() *string {
	return s.MasterAccount
}

func (s *GetSubPartnerListResponseBodySubPartnerList) GetMasterUid() *string {
	return s.MasterUid
}

func (s *GetSubPartnerListResponseBodySubPartnerList) GetPid() *string {
	return s.Pid
}

func (s *GetSubPartnerListResponseBodySubPartnerList) GetProvince() *string {
	return s.Province
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetAddress(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.Address = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetAgreementStatus(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.AgreementStatus = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetAgreementStatusDesc(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.AgreementStatusDesc = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetCity(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.City = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetCompanyName(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.CompanyName = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetContact(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.Contact = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetDistrict(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.District = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetJoinTime(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.JoinTime = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetMasterAccount(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.MasterAccount = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetMasterUid(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.MasterUid = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetPid(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.Pid = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) SetProvince(v string) *GetSubPartnerListResponseBodySubPartnerList {
	s.Province = &v
	return s
}

func (s *GetSubPartnerListResponseBodySubPartnerList) Validate() error {
	return dara.Validate(s)
}
