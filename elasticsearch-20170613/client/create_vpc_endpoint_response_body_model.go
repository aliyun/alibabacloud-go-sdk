// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVpcEndpointResponseBody
	GetRequestId() *string
	SetResult(v *CreateVpcEndpointResponseBodyResult) *CreateVpcEndpointResponseBody
	GetResult() *CreateVpcEndpointResponseBodyResult
}

type CreateVpcEndpointResponseBody struct {
	// The endpoint domain name, which is used to configure the connection.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC47D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the endpoint on the service VPC side.
	Result *CreateVpcEndpointResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateVpcEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpcEndpointResponseBody) GetResult() *CreateVpcEndpointResponseBodyResult {
	return s.Result
}

func (s *CreateVpcEndpointResponseBody) SetRequestId(v string) *CreateVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetResult(v *CreateVpcEndpointResponseBodyResult) *CreateVpcEndpointResponseBody {
	s.Result = v
	return s
}

func (s *CreateVpcEndpointResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVpcEndpointResponseBodyResult struct {
	// example:
	//
	// ep-bp1tah7zbrwmkjef****.epsrv-bp1w0p3jdirbfmt6****.cn-hangzhou.privatelink.aliyuncs.com
	EndpointDomain *string `json:"endpointDomain,omitempty" xml:"endpointDomain,omitempty"`
	// example:
	//
	// ep-bp1tah7zbrwmkjef****
	EndpointId *string `json:"endpointId,omitempty" xml:"endpointId,omitempty"`
	// example:
	//
	// vpcElasticSearchABC
	EndpointName *string `json:"endpointName,omitempty" xml:"endpointName,omitempty"`
	// The name of the service VPC-side endpoint.
	//
	// example:
	//
	// epsrv-bp1w0p3jdirbfmt6****
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s CreateVpcEndpointResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcEndpointResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointResponseBodyResult) GetEndpointDomain() *string {
	return s.EndpointDomain
}

func (s *CreateVpcEndpointResponseBodyResult) GetEndpointId() *string {
	return s.EndpointId
}

func (s *CreateVpcEndpointResponseBodyResult) GetEndpointName() *string {
	return s.EndpointName
}

func (s *CreateVpcEndpointResponseBodyResult) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateVpcEndpointResponseBodyResult) SetEndpointDomain(v string) *CreateVpcEndpointResponseBodyResult {
	s.EndpointDomain = &v
	return s
}

func (s *CreateVpcEndpointResponseBodyResult) SetEndpointId(v string) *CreateVpcEndpointResponseBodyResult {
	s.EndpointId = &v
	return s
}

func (s *CreateVpcEndpointResponseBodyResult) SetEndpointName(v string) *CreateVpcEndpointResponseBodyResult {
	s.EndpointName = &v
	return s
}

func (s *CreateVpcEndpointResponseBodyResult) SetServiceId(v string) *CreateVpcEndpointResponseBodyResult {
	s.ServiceId = &v
	return s
}

func (s *CreateVpcEndpointResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
