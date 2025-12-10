// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptHandshakeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHandshakeId(v string) *AcceptHandshakeRequest
	GetHandshakeId() *string
}

type AcceptHandshakeRequest struct {
	// The ID of the invitation.
	//
	// You can call the [ListHandshakesForAccount](https://help.aliyun.com/document_detail/160006.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// h-Ih8IuPfvV0t0****
	HandshakeId *string `json:"HandshakeId,omitempty" xml:"HandshakeId,omitempty"`
}

func (s AcceptHandshakeRequest) String() string {
	return dara.Prettify(s)
}

func (s AcceptHandshakeRequest) GoString() string {
	return s.String()
}

func (s *AcceptHandshakeRequest) GetHandshakeId() *string {
	return s.HandshakeId
}

func (s *AcceptHandshakeRequest) SetHandshakeId(v string) *AcceptHandshakeRequest {
	s.HandshakeId = &v
	return s
}

func (s *AcceptHandshakeRequest) Validate() error {
	return dara.Validate(s)
}
