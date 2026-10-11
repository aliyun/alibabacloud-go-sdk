// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEgressProxyConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyConfigurationId(v string) *CreateEgressProxyConfigurationResponseBody
	GetEgressProxyConfigurationId() *string
	SetRequestId(v string) *CreateEgressProxyConfigurationResponseBody
	GetRequestId() *string
}

type CreateEgressProxyConfigurationResponseBody struct {
	// The ID of the egress proxy configuration.
	//
	// example:
	//
	// epc-uf609a6a8fdf9f4d****
	EgressProxyConfigurationId *string `json:"EgressProxyConfigurationId,omitempty" xml:"EgressProxyConfigurationId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C79118A9-CA6D-5359-B53E-111BE8A03DCF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEgressProxyConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyConfigurationResponseBody) GetEgressProxyConfigurationId() *string {
	return s.EgressProxyConfigurationId
}

func (s *CreateEgressProxyConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEgressProxyConfigurationResponseBody) SetEgressProxyConfigurationId(v string) *CreateEgressProxyConfigurationResponseBody {
	s.EgressProxyConfigurationId = &v
	return s
}

func (s *CreateEgressProxyConfigurationResponseBody) SetRequestId(v string) *CreateEgressProxyConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEgressProxyConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
