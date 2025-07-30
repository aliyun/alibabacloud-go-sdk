// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTokenResponseBody
	GetRequestId() *string
	SetToken(v string) *GetTokenResponseBody
	GetToken() *string
}

type GetTokenResponseBody struct {
	// The request ID, which is used to troubleshoot issues.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-xxxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sharing token, used to view the information about the shared job.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9*****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTokenResponseBody) GetToken() *string {
	return s.Token
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenResponseBody) SetToken(v string) *GetTokenResponseBody {
	s.Token = &v
	return s
}

func (s *GetTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
