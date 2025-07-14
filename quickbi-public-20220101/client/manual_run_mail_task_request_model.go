// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManualRunMailTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMailId(v string) *ManualRunMailTaskRequest
	GetMailId() *string
}

type ManualRunMailTaskRequest struct {
	// The ID of the email task in the subscription management interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3423423sdfa****sdadw
	MailId *string `json:"MailId,omitempty" xml:"MailId,omitempty"`
}

func (s ManualRunMailTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ManualRunMailTaskRequest) GoString() string {
	return s.String()
}

func (s *ManualRunMailTaskRequest) GetMailId() *string {
	return s.MailId
}

func (s *ManualRunMailTaskRequest) SetMailId(v string) *ManualRunMailTaskRequest {
	s.MailId = &v
	return s
}

func (s *ManualRunMailTaskRequest) Validate() error {
	return dara.Validate(s)
}
