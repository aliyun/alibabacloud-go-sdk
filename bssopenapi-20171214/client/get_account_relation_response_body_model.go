// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAccountRelationResponseBody
	GetCode() *string
	SetData(v *GetAccountRelationResponseBodyData) *GetAccountRelationResponseBody
	GetData() *GetAccountRelationResponseBodyData
	SetMessage(v string) *GetAccountRelationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAccountRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAccountRelationResponseBody
	GetSuccess() *bool
}

type GetAccountRelationResponseBody struct {
	// The request status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data *GetAccountRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message.
	//
	// example:
	//
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAccountRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountRelationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAccountRelationResponseBody) GetData() *GetAccountRelationResponseBodyData {
	return s.Data
}

func (s *GetAccountRelationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAccountRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAccountRelationResponseBody) SetCode(v string) *GetAccountRelationResponseBody {
	s.Code = &v
	return s
}

func (s *GetAccountRelationResponseBody) SetData(v *GetAccountRelationResponseBodyData) *GetAccountRelationResponseBody {
	s.Data = v
	return s
}

func (s *GetAccountRelationResponseBody) SetMessage(v string) *GetAccountRelationResponseBody {
	s.Message = &v
	return s
}

func (s *GetAccountRelationResponseBody) SetRequestId(v string) *GetAccountRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountRelationResponseBody) SetSuccess(v bool) *GetAccountRelationResponseBody {
	s.Success = &v
	return s
}

func (s *GetAccountRelationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccountRelationResponseBodyData struct {
	// The UID of the child account.
	//
	// example:
	//
	// Alibaba Cloud account UID of the child account
	ChildUserId *int64 `json:"ChildUserId,omitempty" xml:"ChildUserId,omitempty"`
	// The end time of the account linking, in the YYYY-MM-DD format.
	//
	// example:
	//
	// 2021-12-01
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The update time of the account linking.
	//
	// example:
	//
	// 2021-12-01
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The account linking ID.
	//
	// example:
	//
	// Account linking ID
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The UID of the parent account.
	//
	// example:
	//
	// Alibaba Cloud UID of the main account
	ParentUserId *int64 `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	// The start time of the account linking, in the YYYY-MM-DD format.
	//
	// example:
	//
	// 2021-11-01
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The account linking status. Valid values:
	//
	// - RELATED: linked.
	//
	// - CONFIRMING: pending confirmation from the other party.
	//
	// - REJECTED: rejected by the other party.
	//
	// - CONNECTION_CANCELED: application canceled by the child account.
	//
	// - CONNECTION_MASTER_CANCEL: invitation canceled by the main account.
	//
	// - CHANGE_CONFIRMING: pending confirmation for relationship change.
	//
	// - INITIAL: initial status of a newly created linking.
	//
	// example:
	//
	// RELATED: linked
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The account linking type.
	//
	// example:
	//
	// enterprise_group
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAccountRelationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAccountRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAccountRelationResponseBodyData) GetChildUserId() *int64 {
	return s.ChildUserId
}

func (s *GetAccountRelationResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetAccountRelationResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetAccountRelationResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetAccountRelationResponseBodyData) GetParentUserId() *int64 {
	return s.ParentUserId
}

func (s *GetAccountRelationResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetAccountRelationResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetAccountRelationResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetAccountRelationResponseBodyData) SetChildUserId(v int64) *GetAccountRelationResponseBodyData {
	s.ChildUserId = &v
	return s
}

func (s *GetAccountRelationResponseBodyData) SetEndTime(v int64) *GetAccountRelationResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetAccountRelationResponseBodyData) SetGmtModified(v int64) *GetAccountRelationResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetAccountRelationResponseBodyData) SetId(v int64) *GetAccountRelationResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetAccountRelationResponseBodyData) SetParentUserId(v int64) *GetAccountRelationResponseBodyData {
	s.ParentUserId = &v
	return s
}

func (s *GetAccountRelationResponseBodyData) SetStartTime(v int64) *GetAccountRelationResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetAccountRelationResponseBodyData) SetStatus(v string) *GetAccountRelationResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAccountRelationResponseBodyData) SetType(v string) *GetAccountRelationResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetAccountRelationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
