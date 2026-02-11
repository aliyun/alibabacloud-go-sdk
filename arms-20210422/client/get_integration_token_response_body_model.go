// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntegrationTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetIntegrationTokenResponseBody
	GetRequestId() *string
	SetToken(v string) *GetIntegrationTokenResponseBody
	GetToken() *string
}

type GetIntegrationTokenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetIntegrationTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntegrationTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIntegrationTokenResponseBody) GetToken() *string {
	return s.Token
}

func (s *GetIntegrationTokenResponseBody) SetRequestId(v string) *GetIntegrationTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIntegrationTokenResponseBody) SetToken(v string) *GetIntegrationTokenResponseBody {
	s.Token = &v
	return s
}

func (s *GetIntegrationTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
