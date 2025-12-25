// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineAccessTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRoutineAccessTokenResponseBody
	GetRequestId() *string
	SetToken(v string) *GetRoutineAccessTokenResponseBody
	GetToken() *string
}

type GetRoutineAccessTokenResponseBody struct {
	// example:
	//
	// EB775B32-1148-1963-9ADD-74CC90C16459
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb3V0aW5lX2lkIjoidHktbWV0YXEtdGVzdC4xNzEzMTU1ODk3ODg1Njg2IiwiZXhwIjoxNzY0OTQ0NTU3fQ.g3gFr-6GQR8vcg6b_vy1qBZ1LDYOiDP-Sih0wtu3d64
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetRoutineAccessTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoutineAccessTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRoutineAccessTokenResponseBody) GetToken() *string {
	return s.Token
}

func (s *GetRoutineAccessTokenResponseBody) SetRequestId(v string) *GetRoutineAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoutineAccessTokenResponseBody) SetToken(v string) *GetRoutineAccessTokenResponseBody {
	s.Token = &v
	return s
}

func (s *GetRoutineAccessTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
