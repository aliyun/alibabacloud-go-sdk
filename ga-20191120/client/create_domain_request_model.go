// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorIds(v []*string) *CreateDomainRequest
	GetAcceleratorIds() []*string
	SetDomain(v string) *CreateDomainRequest
	GetDomain() *string
	SetDryRun(v bool) *CreateDomainRequest
	GetDryRun() *bool
	SetRegionId(v string) *CreateDomainRequest
	GetRegionId() *string
}

type CreateDomainRequest struct {
	// The ID of the GA instance.
	//
	// You can enter up to 50 IDs.
	//
	// This parameter is required.
	AcceleratorIds []*string `json:"AcceleratorIds,omitempty" xml:"AcceleratorIds,omitempty" type:"Repeated"`
	// The accelerated domain name to be added.
	//
	// Wildcard domain names are supported. A wildcard domain name must start with `*.`, such as `*.example.com`.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true:*	- performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, a 2xx HTTP status code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) GetAcceleratorIds() []*string {
	return s.AcceleratorIds
}

func (s *CreateDomainRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateDomainRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDomainRequest) SetAcceleratorIds(v []*string) *CreateDomainRequest {
	s.AcceleratorIds = v
	return s
}

func (s *CreateDomainRequest) SetDomain(v string) *CreateDomainRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainRequest) SetDryRun(v bool) *CreateDomainRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDomainRequest) SetRegionId(v string) *CreateDomainRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDomainRequest) Validate() error {
	return dara.Validate(s)
}
