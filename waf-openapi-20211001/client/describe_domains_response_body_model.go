// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*DescribeDomainsResponseBodyDomains) *DescribeDomainsResponseBody
	GetDomains() []*DescribeDomainsResponseBodyDomains
	SetRequestId(v string) *DescribeDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDomainsResponseBody
	GetTotalCount() *int64
}

type DescribeDomainsResponseBody struct {
	// The domain names that are added to WAF in CNAME record mode.
	Domains []*DescribeDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// ABF68BB3-2C48-5FA4-9750-D5FE55700E36
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 146
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBody) GetDomains() []*DescribeDomainsResponseBodyDomains {
	return s.Domains
}

func (s *DescribeDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDomainsResponseBody) SetDomains(v []*DescribeDomainsResponseBodyDomains) *DescribeDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDomainsResponseBody) SetRequestId(v string) *DescribeDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetTotalCount(v int64) *DescribeDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsResponseBodyDomains struct {
	// The back-to-origin settings.
	Backeds *DescribeDomainsResponseBodyDomainsBackeds `json:"Backeds,omitempty" xml:"Backeds,omitempty" type:"Struct"`
	// The CNAME assigned by WAF to the domain name.
	//
	// example:
	//
	// xxxxxcvdaf.****.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The domain name that is added to WAF in CNAME record mode.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The configurations of the listeners.
	ListenPorts *DescribeDomainsResponseBodyDomainsListenPorts `json:"ListenPorts,omitempty" xml:"ListenPorts,omitempty" type:"Struct"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmvtc5z52****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The status of the domain name. Valid values:
	//
	// 	- **1:*	- The domain name is in a normal state.
	//
	// 	- **2:*	- The domain name is being created.
	//
	// 	- **3:*	- The domain name is being modified.
	//
	// 	- **4:*	- The domain name is being released.
	//
	// 	- **5:*	- WAF no longer forwards traffic that is sent to the domain name.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomains) GetBackeds() *DescribeDomainsResponseBodyDomainsBackeds {
	return s.Backeds
}

func (s *DescribeDomainsResponseBodyDomains) GetCname() *string {
	return s.Cname
}

func (s *DescribeDomainsResponseBodyDomains) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainsResponseBodyDomains) GetDomainId() *string {
	return s.DomainId
}

func (s *DescribeDomainsResponseBodyDomains) GetListenPorts() *DescribeDomainsResponseBodyDomainsListenPorts {
	return s.ListenPorts
}

func (s *DescribeDomainsResponseBodyDomains) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDomainsResponseBodyDomains) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeDomainsResponseBodyDomains) SetBackeds(v *DescribeDomainsResponseBodyDomainsBackeds) *DescribeDomainsResponseBodyDomains {
	s.Backeds = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCname(v string) *DescribeDomainsResponseBodyDomains {
	s.Cname = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetDomain(v string) *DescribeDomainsResponseBodyDomains {
	s.Domain = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetDomainId(v string) *DescribeDomainsResponseBodyDomains {
	s.DomainId = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetListenPorts(v *DescribeDomainsResponseBodyDomainsListenPorts) *DescribeDomainsResponseBodyDomains {
	s.ListenPorts = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetResourceManagerResourceGroupId(v string) *DescribeDomainsResponseBodyDomains {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetStatus(v int32) *DescribeDomainsResponseBodyDomains {
	s.Status = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsResponseBodyDomainsBackeds struct {
	// The HTTP addresses of the origin server.
	Http []*DescribeDomainsResponseBodyDomainsBackedsHttp `json:"Http,omitempty" xml:"Http,omitempty" type:"Repeated"`
	// The HTTPS addresses of the origin server.
	Https []*DescribeDomainsResponseBodyDomainsBackedsHttps `json:"Https,omitempty" xml:"Https,omitempty" type:"Repeated"`
}

func (s DescribeDomainsResponseBodyDomainsBackeds) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsBackeds) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsBackeds) GetHttp() []*DescribeDomainsResponseBodyDomainsBackedsHttp {
	return s.Http
}

func (s *DescribeDomainsResponseBodyDomainsBackeds) GetHttps() []*DescribeDomainsResponseBodyDomainsBackedsHttps {
	return s.Https
}

func (s *DescribeDomainsResponseBodyDomainsBackeds) SetHttp(v []*DescribeDomainsResponseBodyDomainsBackedsHttp) *DescribeDomainsResponseBodyDomainsBackeds {
	s.Http = v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsBackeds) SetHttps(v []*DescribeDomainsResponseBodyDomainsBackedsHttps) *DescribeDomainsResponseBodyDomainsBackeds {
	s.Https = v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsBackeds) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsResponseBodyDomainsBackedsHttp struct {
	// The HTTP address of the origin server.
	//
	// example:
	//
	// 1.1.XX.XX
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
}

func (s DescribeDomainsResponseBodyDomainsBackedsHttp) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsBackedsHttp) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsBackedsHttp) GetBackend() *string {
	return s.Backend
}

func (s *DescribeDomainsResponseBodyDomainsBackedsHttp) SetBackend(v string) *DescribeDomainsResponseBodyDomainsBackedsHttp {
	s.Backend = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsBackedsHttp) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsResponseBodyDomainsBackedsHttps struct {
	// The HTTPS address of the origin server.
	//
	// example:
	//
	// 1.1.XX.XX
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
}

func (s DescribeDomainsResponseBodyDomainsBackedsHttps) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsBackedsHttps) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsBackedsHttps) GetBackend() *string {
	return s.Backend
}

func (s *DescribeDomainsResponseBodyDomainsBackedsHttps) SetBackend(v string) *DescribeDomainsResponseBodyDomainsBackedsHttps {
	s.Backend = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsBackedsHttps) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsResponseBodyDomainsListenPorts struct {
	// The HTTP listener ports.
	Http []*int64 `json:"Http,omitempty" xml:"Http,omitempty" type:"Repeated"`
	// The HTTPS listener ports.
	Https []*int64 `json:"Https,omitempty" xml:"Https,omitempty" type:"Repeated"`
}

func (s DescribeDomainsResponseBodyDomainsListenPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsListenPorts) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsListenPorts) GetHttp() []*int64 {
	return s.Http
}

func (s *DescribeDomainsResponseBodyDomainsListenPorts) GetHttps() []*int64 {
	return s.Https
}

func (s *DescribeDomainsResponseBodyDomainsListenPorts) SetHttp(v []*int64) *DescribeDomainsResponseBodyDomainsListenPorts {
	s.Http = v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsListenPorts) SetHttps(v []*int64) *DescribeDomainsResponseBodyDomainsListenPorts {
	s.Https = v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsListenPorts) Validate() error {
	return dara.Validate(s)
}
