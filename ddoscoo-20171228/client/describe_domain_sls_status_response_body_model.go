// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSlsStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDomainSlsStatusResponseBody
	GetRequestId() *string
	SetSlsLogstore(v string) *DescribeDomainSlsStatusResponseBody
	GetSlsLogstore() *string
	SetSlsProject(v string) *DescribeDomainSlsStatusResponseBody
	GetSlsProject() *string
	SetSlsStatus(v bool) *DescribeDomainSlsStatusResponseBody
	GetSlsStatus() *bool
}

type DescribeDomainSlsStatusResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ddoscoo-logstore
	SlsLogstore *string `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
	// example:
	//
	// ddoscoo-project-xxxx-cn-hangzhou
	SlsProject *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	// example:
	//
	// true
	SlsStatus *bool `json:"SlsStatus,omitempty" xml:"SlsStatus,omitempty"`
}

func (s DescribeDomainSlsStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSlsStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSlsStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainSlsStatusResponseBody) GetSlsLogstore() *string {
	return s.SlsLogstore
}

func (s *DescribeDomainSlsStatusResponseBody) GetSlsProject() *string {
	return s.SlsProject
}

func (s *DescribeDomainSlsStatusResponseBody) GetSlsStatus() *bool {
	return s.SlsStatus
}

func (s *DescribeDomainSlsStatusResponseBody) SetRequestId(v string) *DescribeDomainSlsStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSlsStatusResponseBody) SetSlsLogstore(v string) *DescribeDomainSlsStatusResponseBody {
	s.SlsLogstore = &v
	return s
}

func (s *DescribeDomainSlsStatusResponseBody) SetSlsProject(v string) *DescribeDomainSlsStatusResponseBody {
	s.SlsProject = &v
	return s
}

func (s *DescribeDomainSlsStatusResponseBody) SetSlsStatus(v bool) *DescribeDomainSlsStatusResponseBody {
	s.SlsStatus = &v
	return s
}

func (s *DescribeDomainSlsStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
