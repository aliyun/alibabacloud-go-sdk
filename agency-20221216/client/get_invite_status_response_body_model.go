// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInviteStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInviteStatusResponseBody
	GetCode() *string
	SetData(v *GetInviteStatusResponseBodyData) *GetInviteStatusResponseBody
	GetData() *GetInviteStatusResponseBodyData
	SetMessage(v string) *GetInviteStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInviteStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInviteStatusResponseBody
	GetSuccess() *bool
}

type GetInviteStatusResponseBody struct {
	// Status Code. Error Code:
	//
	// - 3057 InviteId is empty
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetInviteStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, Alibaba Cloud will track errors with this.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInviteStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInviteStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetInviteStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInviteStatusResponseBody) GetData() *GetInviteStatusResponseBodyData {
	return s.Data
}

func (s *GetInviteStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInviteStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInviteStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInviteStatusResponseBody) SetCode(v string) *GetInviteStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetInviteStatusResponseBody) SetData(v *GetInviteStatusResponseBodyData) *GetInviteStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetInviteStatusResponseBody) SetMessage(v string) *GetInviteStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetInviteStatusResponseBody) SetRequestId(v string) *GetInviteStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInviteStatusResponseBody) SetSuccess(v bool) *GetInviteStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetInviteStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInviteStatusResponseBodyData struct {
	InviteStatus []*GetInviteStatusResponseBodyDataInviteStatus `json:"InviteStatus,omitempty" xml:"InviteStatus,omitempty" type:"Repeated"`
}

func (s GetInviteStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInviteStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInviteStatusResponseBodyData) GetInviteStatus() []*GetInviteStatusResponseBodyDataInviteStatus {
	return s.InviteStatus
}

func (s *GetInviteStatusResponseBodyData) SetInviteStatus(v []*GetInviteStatusResponseBodyDataInviteStatus) *GetInviteStatusResponseBodyData {
	s.InviteStatus = v
	return s
}

func (s *GetInviteStatusResponseBodyData) Validate() error {
	if s.InviteStatus != nil {
		for _, item := range s.InviteStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInviteStatusResponseBodyDataInviteStatus struct {
	// Result Code. Value Range:
	//
	// 	- 200 OK
	//
	// 	- 1109 system error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of Invitation Status result
	InviteStatusList *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList `json:"InviteStatusList,omitempty" xml:"InviteStatusList,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call itself is successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInviteStatusResponseBodyDataInviteStatus) String() string {
	return dara.Prettify(s)
}

func (s GetInviteStatusResponseBodyDataInviteStatus) GoString() string {
	return s.String()
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) GetCode() *string {
	return s.Code
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) GetInviteStatusList() *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	return s.InviteStatusList
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) GetMessage() *string {
	return s.Message
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) GetSuccess() *bool {
	return s.Success
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) SetCode(v string) *GetInviteStatusResponseBodyDataInviteStatus {
	s.Code = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) SetInviteStatusList(v *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) *GetInviteStatusResponseBodyDataInviteStatus {
	s.InviteStatusList = v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) SetMessage(v string) *GetInviteStatusResponseBodyDataInviteStatus {
	s.Message = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) SetSuccess(v bool) *GetInviteStatusResponseBodyDataInviteStatus {
	s.Success = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatus) Validate() error {
	if s.InviteStatusList != nil {
		if err := s.InviteStatusList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInviteStatusResponseBodyDataInviteStatusInviteStatusList struct {
	// The time that Distribution Customer successfully associated with Distributor.</br>
	//
	// This value will be empty if there is no existing association.
	//
	// example:
	//
	// 2018-02-13
	AssociationSuccessTime *string `json:"AssociationSuccessTime,omitempty" xml:"AssociationSuccessTime,omitempty"`
	// Distribution Customer\\"s CID
	//
	// example:
	//
	// 1234567890123
	Cid *int64 `json:"Cid,omitempty" xml:"Cid,omitempty"`
	// The time of email been sent out.
	//
	// example:
	//
	// 2018-02-12
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The parent organization ID.
	//
	// example:
	//
	// 1093238769140523
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
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
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Account Type:
	//
	// - 1 Agency\\"s End User
	//
	// - 2 Reseller\\"s End User
	//
	// - 5 T2 Reseller Partner
	//
	// example:
	//
	// 1
	SubAccountType *string `json:"SubAccountType,omitempty" xml:"SubAccountType,omitempty"`
	// Distribution Customer\\"s UID
	//
	// example:
	//
	// 1234567890123
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) String() string {
	return dara.Prettify(s)
}

func (s GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) GoString() string {
	return s.String()
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) GetAssociationSuccessTime() *string {
	return s.AssociationSuccessTime
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) GetCid() *int64 {
	return s.Cid
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) GetParentId() *string {
	return s.ParentId
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) GetStatus() *int32 {
	return s.Status
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) GetSubAccountType() *string {
	return s.SubAccountType
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) GetUid() *int64 {
	return s.Uid
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetAssociationSuccessTime(v string) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.AssociationSuccessTime = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetCid(v int64) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.Cid = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetGmtCreate(v string) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.GmtCreate = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetParentId(v string) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.ParentId = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetStatus(v int32) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.Status = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetSubAccountType(v string) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.SubAccountType = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) SetUid(v int64) *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList {
	s.Uid = &v
	return s
}

func (s *GetInviteStatusResponseBodyDataInviteStatusInviteStatusList) Validate() error {
	return dara.Validate(s)
}
