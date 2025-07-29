// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcesDeleteProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourcesDeleteProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourcesDeleteProtectionResponse
	GetStatusCode() *int32
	SetBody(v []*DescribeResourcesDeleteProtectionResponseBody) *DescribeResourcesDeleteProtectionResponse
	GetBody() []*DescribeResourcesDeleteProtectionResponseBody
}

type DescribeResourcesDeleteProtectionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*DescribeResourcesDeleteProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s DescribeResourcesDeleteProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesDeleteProtectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcesDeleteProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourcesDeleteProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourcesDeleteProtectionResponse) GetBody() []*DescribeResourcesDeleteProtectionResponseBody {
	return s.Body
}

func (s *DescribeResourcesDeleteProtectionResponse) SetHeaders(v map[string]*string) *DescribeResourcesDeleteProtectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourcesDeleteProtectionResponse) SetStatusCode(v int32) *DescribeResourcesDeleteProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourcesDeleteProtectionResponse) SetBody(v []*DescribeResourcesDeleteProtectionResponseBody) *DescribeResourcesDeleteProtectionResponse {
	s.Body = v
	return s
}

func (s *DescribeResourcesDeleteProtectionResponse) Validate() error {
	return dara.Validate(s)
}

type DescribeResourcesDeleteProtectionResponseBody struct {
	// The resource name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The namespace to which the resource belongs.
	//
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// namespaces
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
	// Indicates whether deletion protection is enabled.
	//
	// 	- true: deletion protection is enabled.
	//
	// 	- false: deletion protection is disabled.
	//
	// example:
	//
	// false
	Protection *bool `json:"protection,omitempty" xml:"protection,omitempty"`
}

func (s DescribeResourcesDeleteProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesDeleteProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourcesDeleteProtectionResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeResourcesDeleteProtectionResponseBody) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeResourcesDeleteProtectionResponseBody) GetResource() *string {
	return s.Resource
}

func (s *DescribeResourcesDeleteProtectionResponseBody) GetProtection() *bool {
	return s.Protection
}

func (s *DescribeResourcesDeleteProtectionResponseBody) SetName(v string) *DescribeResourcesDeleteProtectionResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeResourcesDeleteProtectionResponseBody) SetNamespace(v string) *DescribeResourcesDeleteProtectionResponseBody {
	s.Namespace = &v
	return s
}

func (s *DescribeResourcesDeleteProtectionResponseBody) SetResource(v string) *DescribeResourcesDeleteProtectionResponseBody {
	s.Resource = &v
	return s
}

func (s *DescribeResourcesDeleteProtectionResponseBody) SetProtection(v bool) *DescribeResourcesDeleteProtectionResponseBody {
	s.Protection = &v
	return s
}

func (s *DescribeResourcesDeleteProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
