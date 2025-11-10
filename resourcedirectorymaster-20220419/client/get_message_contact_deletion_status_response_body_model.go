// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageContactDeletionStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactDeletionStatus(v *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) *GetMessageContactDeletionStatusResponseBody
	GetContactDeletionStatus() *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus
	SetRequestId(v string) *GetMessageContactDeletionStatusResponseBody
	GetRequestId() *string
}

type GetMessageContactDeletionStatusResponseBody struct {
	// The deletion information of the contact.
	ContactDeletionStatus *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus `json:"ContactDeletionStatus,omitempty" xml:"ContactDeletionStatus,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 95060F1D-6990-4645-8920-A81D1BBFE992
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMessageContactDeletionStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMessageContactDeletionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageContactDeletionStatusResponseBody) GetContactDeletionStatus() *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus {
	return s.ContactDeletionStatus
}

func (s *GetMessageContactDeletionStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMessageContactDeletionStatusResponseBody) SetContactDeletionStatus(v *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) *GetMessageContactDeletionStatusResponseBody {
	s.ContactDeletionStatus = v
	return s
}

func (s *GetMessageContactDeletionStatusResponseBody) SetRequestId(v string) *GetMessageContactDeletionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMessageContactDeletionStatusResponseBody) Validate() error {
	if s.ContactDeletionStatus != nil {
		if err := s.ContactDeletionStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMessageContactDeletionStatusResponseBodyContactDeletionStatus struct {
	// The ID of the contact.
	//
	// example:
	//
	// c-qL4HqKONzOM7****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The types of messages received by the contact.
	FailReasonList []*GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList `json:"FailReasonList,omitempty" xml:"FailReasonList,omitempty" type:"Repeated"`
	// The deletion status of the contact. Valid values:
	//
	// 	- Deleting
	//
	// 	- Failed
	//
	// example:
	//
	// Deleting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) String() string {
	return dara.Prettify(s)
}

func (s GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) GoString() string {
	return s.String()
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) GetContactId() *string {
	return s.ContactId
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) GetFailReasonList() []*GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList {
	return s.FailReasonList
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) GetStatus() *string {
	return s.Status
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) SetContactId(v string) *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus {
	s.ContactId = &v
	return s
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) SetFailReasonList(v []*GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList) *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus {
	s.FailReasonList = v
	return s
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) SetStatus(v string) *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus {
	s.Status = &v
	return s
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatus) Validate() error {
	if s.FailReasonList != nil {
		for _, item := range s.FailReasonList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 199796839435****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The types of messages received by the contact.
	MessageTypes []*string `json:"MessageTypes,omitempty" xml:"MessageTypes,omitempty" type:"Repeated"`
}

func (s GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList) String() string {
	return dara.Prettify(s)
}

func (s GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList) GoString() string {
	return s.String()
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList) GetAccountId() *string {
	return s.AccountId
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList) GetMessageTypes() []*string {
	return s.MessageTypes
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList) SetAccountId(v string) *GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList {
	s.AccountId = &v
	return s
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList) SetMessageTypes(v []*string) *GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList {
	s.MessageTypes = v
	return s
}

func (s *GetMessageContactDeletionStatusResponseBodyContactDeletionStatusFailReasonList) Validate() error {
	return dara.Validate(s)
}
