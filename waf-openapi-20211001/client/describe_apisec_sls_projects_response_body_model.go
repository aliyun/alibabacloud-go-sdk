// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecSlsProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjects(v []*string) *DescribeApisecSlsProjectsResponseBody
	GetProjects() []*string
	SetRequestId(v string) *DescribeApisecSlsProjectsResponseBody
	GetRequestId() *string
}

type DescribeApisecSlsProjectsResponseBody struct {
	// The names of the projects in Simple Log Service.
	Projects []*string `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApisecSlsProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecSlsProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecSlsProjectsResponseBody) GetProjects() []*string {
	return s.Projects
}

func (s *DescribeApisecSlsProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecSlsProjectsResponseBody) SetProjects(v []*string) *DescribeApisecSlsProjectsResponseBody {
	s.Projects = v
	return s
}

func (s *DescribeApisecSlsProjectsResponseBody) SetRequestId(v string) *DescribeApisecSlsProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecSlsProjectsResponseBody) Validate() error {
	return dara.Validate(s)
}
