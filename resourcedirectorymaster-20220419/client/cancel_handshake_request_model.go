// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelHandshakeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHandshakeId(v string) *CancelHandshakeRequest
	GetHandshakeId() *string
}

type CancelHandshakeRequest struct {
	// The ID of the invitation.
	//
	// This parameter is required.
	//
	// example:
	//
	// h-ycm4rp****
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
}

func (s CancelHandshakeRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelHandshakeRequest) GoString() string {
	return s.String()
}

func (s *CancelHandshakeRequest) GetHandshakeId() *string {
	return s.HandshakeId
}

func (s *CancelHandshakeRequest) SetHandshakeId(v string) *CancelHandshakeRequest {
	s.HandshakeId = &v
	return s
}

func (s *CancelHandshakeRequest) Validate() error {
	return dara.Validate(s)
}
