// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationAccessPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *CreateApplicationAccessPointResponseBody
	GetArn() *string
	SetAuthenticationMethod(v string) *CreateApplicationAccessPointResponseBody
	GetAuthenticationMethod() *string
	SetDescription(v string) *CreateApplicationAccessPointResponseBody
	GetDescription() *string
	SetName(v string) *CreateApplicationAccessPointResponseBody
	GetName() *string
	SetPolicies(v string) *CreateApplicationAccessPointResponseBody
	GetPolicies() *string
	SetRequestId(v string) *CreateApplicationAccessPointResponseBody
	GetRequestId() *string
}

type CreateApplicationAccessPointResponseBody struct {
	// The Alibaba Cloud Resource Name (ARN) of the AAP.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:119285303511****:applicationaccesspoint/aap_test
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The authentication method.
	//
	// example:
	//
	// ClientKey
	AuthenticationMethod *string `json:"AuthenticationMethod,omitempty" xml:"AuthenticationMethod,omitempty"`
	// The description of the AAP.
	//
	// example:
	//
	// aap description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the AAP.
	//
	// example:
	//
	// aap_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The permission policy.
	//
	// example:
	//
	// ["kst-hzz62ee817bvyyr5x****.efkd","kst-hzz62ee817bvyyr5x****.eyyp"]
	Policies *string `json:"Policies,omitempty" xml:"Policies,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// bcfefe15-46f0-44a3-bd96-3d422474b71a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationAccessPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationAccessPointResponseBody) GetArn() *string {
	return s.Arn
}

func (s *CreateApplicationAccessPointResponseBody) GetAuthenticationMethod() *string {
	return s.AuthenticationMethod
}

func (s *CreateApplicationAccessPointResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationAccessPointResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateApplicationAccessPointResponseBody) GetPolicies() *string {
	return s.Policies
}

func (s *CreateApplicationAccessPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationAccessPointResponseBody) SetArn(v string) *CreateApplicationAccessPointResponseBody {
	s.Arn = &v
	return s
}

func (s *CreateApplicationAccessPointResponseBody) SetAuthenticationMethod(v string) *CreateApplicationAccessPointResponseBody {
	s.AuthenticationMethod = &v
	return s
}

func (s *CreateApplicationAccessPointResponseBody) SetDescription(v string) *CreateApplicationAccessPointResponseBody {
	s.Description = &v
	return s
}

func (s *CreateApplicationAccessPointResponseBody) SetName(v string) *CreateApplicationAccessPointResponseBody {
	s.Name = &v
	return s
}

func (s *CreateApplicationAccessPointResponseBody) SetPolicies(v string) *CreateApplicationAccessPointResponseBody {
	s.Policies = &v
	return s
}

func (s *CreateApplicationAccessPointResponseBody) SetRequestId(v string) *CreateApplicationAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationAccessPointResponseBody) Validate() error {
	return dara.Validate(s)
}
