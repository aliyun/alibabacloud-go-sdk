// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointServicesByEndUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListVpcEndpointServicesByEndUserResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListVpcEndpointServicesByEndUserResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVpcEndpointServicesByEndUserResponseBody
	GetRequestId() *string
	SetServices(v []*ListVpcEndpointServicesByEndUserResponseBodyServices) *ListVpcEndpointServicesByEndUserResponseBody
	GetServices() []*ListVpcEndpointServicesByEndUserResponseBodyServices
	SetTotalCount(v string) *ListVpcEndpointServicesByEndUserResponseBody
	GetTotalCount() *string
}

type ListVpcEndpointServicesByEndUserResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next queries are sent.
	//
	// 	- If **NextToken*	- is returned, the value is the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0AB1129F-32C1-5E4D-9E22-E4A859CA46EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of entries returned.
	Services []*ListVpcEndpointServicesByEndUserResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVpcEndpointServicesByEndUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServicesByEndUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) GetServices() []*ListVpcEndpointServicesByEndUserResponseBodyServices {
	return s.Services
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) SetMaxResults(v int64) *ListVpcEndpointServicesByEndUserResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) SetNextToken(v string) *ListVpcEndpointServicesByEndUserResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) SetRequestId(v string) *ListVpcEndpointServicesByEndUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) SetServices(v []*ListVpcEndpointServicesByEndUserResponseBodyServices) *ListVpcEndpointServicesByEndUserResponseBody {
	s.Services = v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) SetTotalCount(v string) *ListVpcEndpointServicesByEndUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVpcEndpointServicesByEndUserResponseBodyServices struct {
	// The default access policy.
	//
	// example:
	//
	// {   \\"Version\\" : \\"1\\",   \\"Statement\\" : [ {     \\"Effect\\" : \\"Allow\\",     \\"Action\\" : \\"*\\",     \\"Principal\\" : \\"*\\",     \\"Resource\\" : \\"*\\"   } ] }
	DefaultPolicyDocument *string `json:"DefaultPolicyDocument,omitempty" xml:"DefaultPolicyDocument,omitempty"`
	// The domain name of the cloud service to which the endpoint service belongs.
	//
	// example:
	//
	// oss-admin.aliyuncs.com
	ServiceDomain *string `json:"ServiceDomain,omitempty" xml:"ServiceDomain,omitempty"`
	// The ID of the endpoint service.
	//
	// example:
	//
	// vpces-m5enwdmilo210aibo9****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service.
	//
	// example:
	//
	// com.aliyun.cn-hangzhou.oss
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// Indicate whether the endpoint service supports the access policy. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	SupportPolicy *bool `json:"SupportPolicy,omitempty" xml:"SupportPolicy,omitempty"`
}

func (s ListVpcEndpointServicesByEndUserResponseBodyServices) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServicesByEndUserResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetDefaultPolicyDocument() *string {
	return s.DefaultPolicyDocument
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetServiceDomain() *string {
	return s.ServiceDomain
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) GetSupportPolicy() *bool {
	return s.SupportPolicy
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetDefaultPolicyDocument(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.DefaultPolicyDocument = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetServiceDomain(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ServiceDomain = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetServiceId(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetServiceName(v string) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.ServiceName = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) SetSupportPolicy(v bool) *ListVpcEndpointServicesByEndUserResponseBodyServices {
	s.SupportPolicy = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponseBodyServices) Validate() error {
	return dara.Validate(s)
}
