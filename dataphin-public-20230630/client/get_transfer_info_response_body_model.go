// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTransferInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTransferInfoResponseBody
	GetCode() *string
	SetData(v *GetTransferInfoResponseBodyData) *GetTransferInfoResponseBody
	GetData() *GetTransferInfoResponseBodyData
	SetHttpStatusCode(v int32) *GetTransferInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTransferInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTransferInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTransferInfoResponseBody
	GetSuccess() *bool
}

type GetTransferInfoResponseBody struct {
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTransferInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTransferInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTransferInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetTransferInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTransferInfoResponseBody) GetData() *GetTransferInfoResponseBodyData {
	return s.Data
}

func (s *GetTransferInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTransferInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTransferInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTransferInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTransferInfoResponseBody) SetCode(v string) *GetTransferInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetTransferInfoResponseBody) SetData(v *GetTransferInfoResponseBodyData) *GetTransferInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetTransferInfoResponseBody) SetHttpStatusCode(v int32) *GetTransferInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTransferInfoResponseBody) SetMessage(v string) *GetTransferInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetTransferInfoResponseBody) SetRequestId(v string) *GetTransferInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTransferInfoResponseBody) SetSuccess(v bool) *GetTransferInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetTransferInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTransferInfoResponseBodyData struct {
	Creator *GetTransferInfoResponseBodyDataCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// example:
	//
	// 1753669315426
	FlowId *int64 `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// example:
	//
	// 1632036495973809
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1632036495973809
	GmtModified  *string                                      `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	LastModifier *GetTransferInfoResponseBodyDataLastModifier `json:"LastModifier,omitempty" xml:"LastModifier,omitempty" type:"Struct"`
	NewOwner     *GetTransferInfoResponseBodyDataNewOwner     `json:"NewOwner,omitempty" xml:"NewOwner,omitempty" type:"Struct"`
	OldOwner     *GetTransferInfoResponseBodyDataOldOwner     `json:"OldOwner,omitempty" xml:"OldOwner,omitempty" type:"Struct"`
	// example:
	//
	// ONE_STOP
	PrivilegeTransferMode          *string                                                          `json:"PrivilegeTransferMode,omitempty" xml:"PrivilegeTransferMode,omitempty"`
	PrivilegeTransferResultEntries []*GetTransferInfoResponseBodyDataPrivilegeTransferResultEntries `json:"PrivilegeTransferResultEntries,omitempty" xml:"PrivilegeTransferResultEntries,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	ProposalId *int64 `json:"ProposalId,omitempty" xml:"ProposalId,omitempty"`
	// example:
	//
	// transefer title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// comment
	TransferComment *string `json:"TransferComment,omitempty" xml:"TransferComment,omitempty"`
	// example:
	//
	// APPROVED
	TransferStatus *string `json:"TransferStatus,omitempty" xml:"TransferStatus,omitempty"`
}

func (s GetTransferInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTransferInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTransferInfoResponseBodyData) GetCreator() *GetTransferInfoResponseBodyDataCreator {
	return s.Creator
}

func (s *GetTransferInfoResponseBodyData) GetFlowId() *int64 {
	return s.FlowId
}

func (s *GetTransferInfoResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetTransferInfoResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetTransferInfoResponseBodyData) GetLastModifier() *GetTransferInfoResponseBodyDataLastModifier {
	return s.LastModifier
}

func (s *GetTransferInfoResponseBodyData) GetNewOwner() *GetTransferInfoResponseBodyDataNewOwner {
	return s.NewOwner
}

func (s *GetTransferInfoResponseBodyData) GetOldOwner() *GetTransferInfoResponseBodyDataOldOwner {
	return s.OldOwner
}

func (s *GetTransferInfoResponseBodyData) GetPrivilegeTransferMode() *string {
	return s.PrivilegeTransferMode
}

func (s *GetTransferInfoResponseBodyData) GetPrivilegeTransferResultEntries() []*GetTransferInfoResponseBodyDataPrivilegeTransferResultEntries {
	return s.PrivilegeTransferResultEntries
}

func (s *GetTransferInfoResponseBodyData) GetProposalId() *int64 {
	return s.ProposalId
}

func (s *GetTransferInfoResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetTransferInfoResponseBodyData) GetTransferComment() *string {
	return s.TransferComment
}

func (s *GetTransferInfoResponseBodyData) GetTransferStatus() *string {
	return s.TransferStatus
}

func (s *GetTransferInfoResponseBodyData) SetCreator(v *GetTransferInfoResponseBodyDataCreator) *GetTransferInfoResponseBodyData {
	s.Creator = v
	return s
}

func (s *GetTransferInfoResponseBodyData) SetFlowId(v int64) *GetTransferInfoResponseBodyData {
	s.FlowId = &v
	return s
}

func (s *GetTransferInfoResponseBodyData) SetGmtCreate(v string) *GetTransferInfoResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetTransferInfoResponseBodyData) SetGmtModified(v string) *GetTransferInfoResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetTransferInfoResponseBodyData) SetLastModifier(v *GetTransferInfoResponseBodyDataLastModifier) *GetTransferInfoResponseBodyData {
	s.LastModifier = v
	return s
}

func (s *GetTransferInfoResponseBodyData) SetNewOwner(v *GetTransferInfoResponseBodyDataNewOwner) *GetTransferInfoResponseBodyData {
	s.NewOwner = v
	return s
}

func (s *GetTransferInfoResponseBodyData) SetOldOwner(v *GetTransferInfoResponseBodyDataOldOwner) *GetTransferInfoResponseBodyData {
	s.OldOwner = v
	return s
}

func (s *GetTransferInfoResponseBodyData) SetPrivilegeTransferMode(v string) *GetTransferInfoResponseBodyData {
	s.PrivilegeTransferMode = &v
	return s
}

func (s *GetTransferInfoResponseBodyData) SetPrivilegeTransferResultEntries(v []*GetTransferInfoResponseBodyDataPrivilegeTransferResultEntries) *GetTransferInfoResponseBodyData {
	s.PrivilegeTransferResultEntries = v
	return s
}

func (s *GetTransferInfoResponseBodyData) SetProposalId(v int64) *GetTransferInfoResponseBodyData {
	s.ProposalId = &v
	return s
}

func (s *GetTransferInfoResponseBodyData) SetTitle(v string) *GetTransferInfoResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetTransferInfoResponseBodyData) SetTransferComment(v string) *GetTransferInfoResponseBodyData {
	s.TransferComment = &v
	return s
}

func (s *GetTransferInfoResponseBodyData) SetTransferStatus(v string) *GetTransferInfoResponseBodyData {
	s.TransferStatus = &v
	return s
}

func (s *GetTransferInfoResponseBodyData) Validate() error {
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.LastModifier != nil {
		if err := s.LastModifier.Validate(); err != nil {
			return err
		}
	}
	if s.NewOwner != nil {
		if err := s.NewOwner.Validate(); err != nil {
			return err
		}
	}
	if s.OldOwner != nil {
		if err := s.OldOwner.Validate(); err != nil {
			return err
		}
	}
	if s.PrivilegeTransferResultEntries != nil {
		for _, item := range s.PrivilegeTransferResultEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTransferInfoResponseBodyDataCreator struct {
	// example:
	//
	// zhangsan
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 30000001
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetTransferInfoResponseBodyDataCreator) String() string {
	return dara.Prettify(s)
}

func (s GetTransferInfoResponseBodyDataCreator) GoString() string {
	return s.String()
}

func (s *GetTransferInfoResponseBodyDataCreator) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetTransferInfoResponseBodyDataCreator) GetUserId() *string {
	return s.UserId
}

func (s *GetTransferInfoResponseBodyDataCreator) SetDisplayName(v string) *GetTransferInfoResponseBodyDataCreator {
	s.DisplayName = &v
	return s
}

func (s *GetTransferInfoResponseBodyDataCreator) SetUserId(v string) *GetTransferInfoResponseBodyDataCreator {
	s.UserId = &v
	return s
}

func (s *GetTransferInfoResponseBodyDataCreator) Validate() error {
	return dara.Validate(s)
}

type GetTransferInfoResponseBodyDataLastModifier struct {
	// example:
	//
	// zhangsan
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 30000001
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetTransferInfoResponseBodyDataLastModifier) String() string {
	return dara.Prettify(s)
}

func (s GetTransferInfoResponseBodyDataLastModifier) GoString() string {
	return s.String()
}

func (s *GetTransferInfoResponseBodyDataLastModifier) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetTransferInfoResponseBodyDataLastModifier) GetUserId() *string {
	return s.UserId
}

func (s *GetTransferInfoResponseBodyDataLastModifier) SetDisplayName(v string) *GetTransferInfoResponseBodyDataLastModifier {
	s.DisplayName = &v
	return s
}

func (s *GetTransferInfoResponseBodyDataLastModifier) SetUserId(v string) *GetTransferInfoResponseBodyDataLastModifier {
	s.UserId = &v
	return s
}

func (s *GetTransferInfoResponseBodyDataLastModifier) Validate() error {
	return dara.Validate(s)
}

type GetTransferInfoResponseBodyDataNewOwner struct {
	// example:
	//
	// zhangsan
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 30000001
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetTransferInfoResponseBodyDataNewOwner) String() string {
	return dara.Prettify(s)
}

func (s GetTransferInfoResponseBodyDataNewOwner) GoString() string {
	return s.String()
}

func (s *GetTransferInfoResponseBodyDataNewOwner) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetTransferInfoResponseBodyDataNewOwner) GetUserId() *string {
	return s.UserId
}

func (s *GetTransferInfoResponseBodyDataNewOwner) SetDisplayName(v string) *GetTransferInfoResponseBodyDataNewOwner {
	s.DisplayName = &v
	return s
}

func (s *GetTransferInfoResponseBodyDataNewOwner) SetUserId(v string) *GetTransferInfoResponseBodyDataNewOwner {
	s.UserId = &v
	return s
}

func (s *GetTransferInfoResponseBodyDataNewOwner) Validate() error {
	return dara.Validate(s)
}

type GetTransferInfoResponseBodyDataOldOwner struct {
	// example:
	//
	// zhangsan
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 30000001
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetTransferInfoResponseBodyDataOldOwner) String() string {
	return dara.Prettify(s)
}

func (s GetTransferInfoResponseBodyDataOldOwner) GoString() string {
	return s.String()
}

func (s *GetTransferInfoResponseBodyDataOldOwner) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetTransferInfoResponseBodyDataOldOwner) GetUserId() *string {
	return s.UserId
}

func (s *GetTransferInfoResponseBodyDataOldOwner) SetDisplayName(v string) *GetTransferInfoResponseBodyDataOldOwner {
	s.DisplayName = &v
	return s
}

func (s *GetTransferInfoResponseBodyDataOldOwner) SetUserId(v string) *GetTransferInfoResponseBodyDataOldOwner {
	s.UserId = &v
	return s
}

func (s *GetTransferInfoResponseBodyDataOldOwner) Validate() error {
	return dara.Validate(s)
}

type GetTransferInfoResponseBodyDataPrivilegeTransferResultEntries struct {
	// example:
	//
	// userId is error
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// example:
	//
	// table owner
	PrivilegeDisplayName *string `json:"PrivilegeDisplayName,omitempty" xml:"PrivilegeDisplayName,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTransferInfoResponseBodyDataPrivilegeTransferResultEntries) String() string {
	return dara.Prettify(s)
}

func (s GetTransferInfoResponseBodyDataPrivilegeTransferResultEntries) GoString() string {
	return s.String()
}

func (s *GetTransferInfoResponseBodyDataPrivilegeTransferResultEntries) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *GetTransferInfoResponseBodyDataPrivilegeTransferResultEntries) GetPrivilegeDisplayName() *string {
	return s.PrivilegeDisplayName
}

func (s *GetTransferInfoResponseBodyDataPrivilegeTransferResultEntries) GetStatus() *string {
	return s.Status
}

func (s *GetTransferInfoResponseBodyDataPrivilegeTransferResultEntries) SetErrMsg(v string) *GetTransferInfoResponseBodyDataPrivilegeTransferResultEntries {
	s.ErrMsg = &v
	return s
}

func (s *GetTransferInfoResponseBodyDataPrivilegeTransferResultEntries) SetPrivilegeDisplayName(v string) *GetTransferInfoResponseBodyDataPrivilegeTransferResultEntries {
	s.PrivilegeDisplayName = &v
	return s
}

func (s *GetTransferInfoResponseBodyDataPrivilegeTransferResultEntries) SetStatus(v string) *GetTransferInfoResponseBodyDataPrivilegeTransferResultEntries {
	s.Status = &v
	return s
}

func (s *GetTransferInfoResponseBodyDataPrivilegeTransferResultEntries) Validate() error {
	return dara.Validate(s)
}
