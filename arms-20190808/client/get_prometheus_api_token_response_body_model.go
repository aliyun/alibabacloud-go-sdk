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
	// The request ID.
	//
	// example:
	//
	// 1A9C645C-C83F-4C9D-8CCB-29BEC9E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The token required for integrating Prometheus Service.
	//
	// example:
	//
	// 6dcbb77ef4ba6ef5466b5debf9e2****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
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
