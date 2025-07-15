// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConversationDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *GetConversationDetailRequest
	GetContactId() *string
	SetInstanceId(v string) *GetConversationDetailRequest
	GetInstanceId() *string
}

type GetConversationDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// job-25884193037652****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetConversationDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConversationDetailRequest) GoString() string {
	return s.String()
}

func (s *GetConversationDetailRequest) GetContactId() *string {
	return s.ContactId
}

func (s *GetConversationDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetConversationDetailRequest) SetContactId(v string) *GetConversationDetailRequest {
	s.ContactId = &v
	return s
}

func (s *GetConversationDetailRequest) SetInstanceId(v string) *GetConversationDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConversationDetailRequest) Validate() error {
	return dara.Validate(s)
}
