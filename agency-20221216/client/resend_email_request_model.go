// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInviteId(v int64) *ResendEmailRequest
	GetInviteId() *int64
}

type ResendEmailRequest struct {
	// Invitation ID, from interface InviteSubAccount </br>
	//
	// Note: This field type is Long, which may result in precision loss in serialization/deserialization process. Please ensure the value does not exceed 9007199254740991.
	//
	// This parameter is required.
	//
	// example:
	//
	// 176
	InviteId *int64 `json:"InviteId,omitempty" xml:"InviteId,omitempty"`
}

func (s ResendEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s ResendEmailRequest) GoString() string {
	return s.String()
}

func (s *ResendEmailRequest) GetInviteId() *int64 {
	return s.InviteId
}

func (s *ResendEmailRequest) SetInviteId(v int64) *ResendEmailRequest {
	s.InviteId = &v
	return s
}

func (s *ResendEmailRequest) Validate() error {
	return dara.Validate(s)
}
