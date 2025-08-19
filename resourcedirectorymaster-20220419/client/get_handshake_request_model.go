// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHandshakeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHandshakeId(v string) *GetHandshakeRequest
	GetHandshakeId() *string
}

type GetHandshakeRequest struct {
	// The ID of the invitation.
	//
	// This parameter is required.
	//
	// example:
	//
	// h-ycm4rp****
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
}

func (s GetHandshakeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHandshakeRequest) GoString() string {
	return s.String()
}

func (s *GetHandshakeRequest) GetHandshakeId() *string {
	return s.HandshakeId
}

func (s *GetHandshakeRequest) SetHandshakeId(v string) *GetHandshakeRequest {
	s.HandshakeId = &v
	return s
}

func (s *GetHandshakeRequest) Validate() error {
	return dara.Validate(s)
}
