// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationAccessPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *DescribeApplicationAccessPointResponseBody
	GetArn() *string
	SetAuthenticationMethod(v string) *DescribeApplicationAccessPointResponseBody
	GetAuthenticationMethod() *string
	SetDescription(v string) *DescribeApplicationAccessPointResponseBody
	GetDescription() *string
	SetName(v string) *DescribeApplicationAccessPointResponseBody
	GetName() *string
	SetPolicies(v string) *DescribeApplicationAccessPointResponseBody
	GetPolicies() *string
	SetRequestId(v string) *DescribeApplicationAccessPointResponseBody
	GetRequestId() *string
}

type DescribeApplicationAccessPointResponseBody struct {
	// The ARN of the AAP.
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
	// The description.
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
	// The permission policy that is bound to the AAP.
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

func (s DescribeApplicationAccessPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAccessPointResponseBody) GetArn() *string {
	return s.Arn
}

func (s *DescribeApplicationAccessPointResponseBody) GetAuthenticationMethod() *string {
	return s.AuthenticationMethod
}

func (s *DescribeApplicationAccessPointResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeApplicationAccessPointResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAccessPointResponseBody) GetPolicies() *string {
	return s.Policies
}

func (s *DescribeApplicationAccessPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationAccessPointResponseBody) SetArn(v string) *DescribeApplicationAccessPointResponseBody {
	s.Arn = &v
	return s
}

func (s *DescribeApplicationAccessPointResponseBody) SetAuthenticationMethod(v string) *DescribeApplicationAccessPointResponseBody {
	s.AuthenticationMethod = &v
	return s
}

func (s *DescribeApplicationAccessPointResponseBody) SetDescription(v string) *DescribeApplicationAccessPointResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeApplicationAccessPointResponseBody) SetName(v string) *DescribeApplicationAccessPointResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAccessPointResponseBody) SetPolicies(v string) *DescribeApplicationAccessPointResponseBody {
	s.Policies = &v
	return s
}

func (s *DescribeApplicationAccessPointResponseBody) SetRequestId(v string) *DescribeApplicationAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationAccessPointResponseBody) Validate() error {
	return dara.Validate(s)
}
