// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeclineHandshakeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHandshakeId(v string) *DeclineHandshakeRequest
	GetHandshakeId() *string
}

type DeclineHandshakeRequest struct {
	// The ID of the invitation.
	//
	// This parameter is required.
	//
	// example:
	//
	// h-ycm4rp****
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
}

func (s DeclineHandshakeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeclineHandshakeRequest) GoString() string {
	return s.String()
}

func (s *DeclineHandshakeRequest) GetHandshakeId() *string {
	return s.HandshakeId
}

func (s *DeclineHandshakeRequest) SetHandshakeId(v string) *DeclineHandshakeRequest {
	s.HandshakeId = &v
	return s
}

func (s *DeclineHandshakeRequest) Validate() error {
	return dara.Validate(s)
}
