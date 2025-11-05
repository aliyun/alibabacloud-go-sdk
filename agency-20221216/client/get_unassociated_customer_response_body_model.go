// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnassociatedCustomerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUnassociatedCustomerResponseBody
	GetCode() *string
	SetInviteInfoList(v *GetUnassociatedCustomerResponseBodyInviteInfoList) *GetUnassociatedCustomerResponseBody
	GetInviteInfoList() *GetUnassociatedCustomerResponseBodyInviteInfoList
	SetMessage(v string) *GetUnassociatedCustomerResponseBody
	GetMessage() *string
	SetPageInfo(v *GetUnassociatedCustomerResponseBodyPageInfo) *GetUnassociatedCustomerResponseBody
	GetPageInfo() *GetUnassociatedCustomerResponseBodyPageInfo
	SetRequestId(v string) *GetUnassociatedCustomerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUnassociatedCustomerResponseBody
	GetSuccess() *bool
}

type GetUnassociatedCustomerResponseBody struct {
	// Error Code, Candidate Value：
	//
	// 	- 200: OK
	//
	// 	- 1109: System error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of Invitation Information
	InviteInfoList *GetUnassociatedCustomerResponseBodyInviteInfoList `json:"InviteInfoList,omitempty" xml:"InviteInfoList,omitempty" type:"Struct"`
	// Message information
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Pagination Information
	PageInfo *GetUnassociatedCustomerResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Request ID, Alibaba Cloud will track errors with this.
	//
	// example:
	//
	// 23309219-4A34-589D-A3E0-9B2A3BFFD24F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUnassociatedCustomerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUnassociatedCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUnassociatedCustomerResponseBody) GetInviteInfoList() *GetUnassociatedCustomerResponseBodyInviteInfoList {
	return s.InviteInfoList
}

func (s *GetUnassociatedCustomerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUnassociatedCustomerResponseBody) GetPageInfo() *GetUnassociatedCustomerResponseBodyPageInfo {
	return s.PageInfo
}

func (s *GetUnassociatedCustomerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUnassociatedCustomerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUnassociatedCustomerResponseBody) SetCode(v string) *GetUnassociatedCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBody) SetInviteInfoList(v *GetUnassociatedCustomerResponseBodyInviteInfoList) *GetUnassociatedCustomerResponseBody {
	s.InviteInfoList = v
	return s
}

func (s *GetUnassociatedCustomerResponseBody) SetMessage(v string) *GetUnassociatedCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBody) SetPageInfo(v *GetUnassociatedCustomerResponseBodyPageInfo) *GetUnassociatedCustomerResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetUnassociatedCustomerResponseBody) SetRequestId(v string) *GetUnassociatedCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBody) SetSuccess(v bool) *GetUnassociatedCustomerResponseBody {
	s.Success = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBody) Validate() error {
	if s.InviteInfoList != nil {
		if err := s.InviteInfoList.Validate(); err != nil {
			return err
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUnassociatedCustomerResponseBodyInviteInfoList struct {
	InviteInfo []*GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo `json:"InviteInfo,omitempty" xml:"InviteInfo,omitempty" type:"Repeated"`
}

func (s GetUnassociatedCustomerResponseBodyInviteInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetUnassociatedCustomerResponseBodyInviteInfoList) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoList) GetInviteInfo() []*GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo {
	return s.InviteInfo
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoList) SetInviteInfo(v []*GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) *GetUnassociatedCustomerResponseBodyInviteInfoList {
	s.InviteInfo = v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoList) Validate() error {
	if s.InviteInfo != nil {
		for _, item := range s.InviteInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo struct {
	// The name of Customer who are to be invited.
	//
	// example:
	//
	// My Client
	AccountNickname *string `json:"AccountNickname,omitempty" xml:"AccountNickname,omitempty"`
	// The Email of Customer who are to be invited.
	//
	// example:
	//
	// 12345@qq.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The time of email been sent out.
	//
	// example:
	//
	// 2023-05-10
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Invitation ID
	//
	// example:
	//
	// 190
	InviteId *int64 `json:"InviteId,omitempty" xml:"InviteId,omitempty"`
	// Invitation Status:
	//
	// 	- 0 No visit on registration URL
	//
	// 	- 1 Successful Registration
	//
	// 	- 2 Unsuccessful Registration
	//
	// 	- 3 Registration URL have been visited, but no submitted sheet/ticket.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) String() string {
	return dara.Prettify(s)
}

func (s GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) GetAccountNickname() *string {
	return s.AccountNickname
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) GetEmail() *string {
	return s.Email
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) GetInviteId() *int64 {
	return s.InviteId
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) GetStatus() *int32 {
	return s.Status
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) SetAccountNickname(v string) *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo {
	s.AccountNickname = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) SetEmail(v string) *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo {
	s.Email = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) SetGmtCreate(v string) *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo {
	s.GmtCreate = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) SetInviteId(v int64) *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo {
	s.InviteId = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) SetStatus(v int32) *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo {
	s.Status = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyInviteInfoListInviteInfo) Validate() error {
	return dara.Validate(s)
}

type GetUnassociatedCustomerResponseBodyPageInfo struct {
	// Pagination, current page.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// Pagination, record number on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Pagination, page volume in total.
	//
	// example:
	//
	// 12
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetUnassociatedCustomerResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetUnassociatedCustomerResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetUnassociatedCustomerResponseBodyPageInfo) GetPage() *int32 {
	return s.Page
}

func (s *GetUnassociatedCustomerResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetUnassociatedCustomerResponseBodyPageInfo) GetTotal() *int32 {
	return s.Total
}

func (s *GetUnassociatedCustomerResponseBodyPageInfo) SetPage(v int32) *GetUnassociatedCustomerResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyPageInfo) SetPageSize(v int32) *GetUnassociatedCustomerResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyPageInfo) SetTotal(v int32) *GetUnassociatedCustomerResponseBodyPageInfo {
	s.Total = &v
	return s
}

func (s *GetUnassociatedCustomerResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
