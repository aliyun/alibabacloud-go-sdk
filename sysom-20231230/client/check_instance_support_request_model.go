// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceSupportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *CheckInstanceSupportRequest
	GetXDebugId() *string
	SetInstances(v []*string) *CheckInstanceSupportRequest
	GetInstances() []*string
	SetRegion(v string) *CheckInstanceSupportRequest
	GetRegion() *string
	SetXSysomInvokeSource(v string) *CheckInstanceSupportRequest
	GetXSysomInvokeSource() *string
}

type CheckInstanceSupportRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The list of instance IDs to check.
	//
	// example:
	//
	// ["i-2zxxxxxx"]
	Instances []*string `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// The region to which the instances belong. All instance IDs specified in instances must belong to the same region.
	//
	// example:
	//
	// cn-hangzhou
	Region             *string `json:"region,omitempty" xml:"region,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s CheckInstanceSupportRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceSupportRequest) GoString() string {
	return s.String()
}

func (s *CheckInstanceSupportRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *CheckInstanceSupportRequest) GetInstances() []*string {
	return s.Instances
}

func (s *CheckInstanceSupportRequest) GetRegion() *string {
	return s.Region
}

func (s *CheckInstanceSupportRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *CheckInstanceSupportRequest) SetXDebugId(v string) *CheckInstanceSupportRequest {
	s.XDebugId = &v
	return s
}

func (s *CheckInstanceSupportRequest) SetInstances(v []*string) *CheckInstanceSupportRequest {
	s.Instances = v
	return s
}

func (s *CheckInstanceSupportRequest) SetRegion(v string) *CheckInstanceSupportRequest {
	s.Region = &v
	return s
}

func (s *CheckInstanceSupportRequest) SetXSysomInvokeSource(v string) *CheckInstanceSupportRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *CheckInstanceSupportRequest) Validate() error {
	return dara.Validate(s)
}
