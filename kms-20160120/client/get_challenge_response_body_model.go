// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChallengeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChallengeToken(v string) *GetChallengeResponseBody
	GetChallengeToken() *string
	SetNonce(v string) *GetChallengeResponseBody
	GetNonce() *string
	SetRequestId(v string) *GetChallengeResponseBody
	GetRequestId() *string
}

type GetChallengeResponseBody struct {
	// example:
	//
	// eyJhbGciOiJSUzM4NCIsInR5cCI6IkpXVCJ9.eyJub25jZSI6Im1OWnpNVENTc3JVT1JTd1d1WFNneDlTNG80MW1Mc3FPS21xd0d4Tzk******E3NTU5NzIzMDB9.signature-part...=
	ChallengeToken *string `json:"ChallengeToken,omitempty" xml:"ChallengeToken,omitempty"`
	// example:
	//
	// mNZzMTCQ******4o1mLsqOKmqwGxO94i9c=
	Nonce *string `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
	// example:
	//
	// c337a6ee-27d1-465e-acb2-dddef7c3c589
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChallengeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChallengeResponseBody) GoString() string {
	return s.String()
}

func (s *GetChallengeResponseBody) GetChallengeToken() *string {
	return s.ChallengeToken
}

func (s *GetChallengeResponseBody) GetNonce() *string {
	return s.Nonce
}

func (s *GetChallengeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChallengeResponseBody) SetChallengeToken(v string) *GetChallengeResponseBody {
	s.ChallengeToken = &v
	return s
}

func (s *GetChallengeResponseBody) SetNonce(v string) *GetChallengeResponseBody {
	s.Nonce = &v
	return s
}

func (s *GetChallengeResponseBody) SetRequestId(v string) *GetChallengeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChallengeResponseBody) Validate() error {
	return dara.Validate(s)
}
