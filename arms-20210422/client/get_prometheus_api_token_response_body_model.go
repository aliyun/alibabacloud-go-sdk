// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusApiTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetPrometheusApiTokenResponseBody
	GetRequestId() *string
	SetToken(v string) *GetPrometheusApiTokenResponseBody
	GetToken() *string
}

type GetPrometheusApiTokenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetPrometheusApiTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusApiTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrometheusApiTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPrometheusApiTokenResponseBody) GetToken() *string {
	return s.Token
}

func (s *GetPrometheusApiTokenResponseBody) SetRequestId(v string) *GetPrometheusApiTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPrometheusApiTokenResponseBody) SetToken(v string) *GetPrometheusApiTokenResponseBody {
	s.Token = &v
	return s
}

func (s *GetPrometheusApiTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
