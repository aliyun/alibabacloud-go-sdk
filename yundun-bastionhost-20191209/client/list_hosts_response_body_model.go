// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHosts(v []*ListHostsResponseBodyHosts) *ListHostsResponseBody
	GetHosts() []*ListHostsResponseBodyHosts
	SetRequestId(v string) *ListHostsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHostsResponseBody
	GetTotalCount() *int32
}

type ListHostsResponseBody struct {
	// An array that consists of the hosts returned.
	Hosts []*ListHostsResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of hosts returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHostsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHostsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostsResponseBody) GetHosts() []*ListHostsResponseBodyHosts {
	return s.Hosts
}

func (s *ListHostsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHostsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHostsResponseBody) SetHosts(v []*ListHostsResponseBodyHosts) *ListHostsResponseBody {
	s.Hosts = v
	return s
}

func (s *ListHostsResponseBody) SetRequestId(v string) *ListHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostsResponseBody) SetTotalCount(v int32) *ListHostsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHostsResponseBodyHosts struct {
	// The address type of the host. Valid values:
	//
	// 	- **Public**: a public address
	//
	// 	- **Private**: a private address
	//
	// example:
	//
	// Public
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The description of the host.
	//
	// example:
	//
	// host
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The number of host accounts.
	//
	// example:
	//
	// 1
	HostAccountCount *int32 `json:"HostAccountCount,omitempty" xml:"HostAccountCount,omitempty"`
	// The ID of the host.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The name of the host.
	//
	// example:
	//
	// name
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The private address of the host. The value is a domain name or an IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	// The public address of the host. The value is a domain name or an IP address.
	//
	// example:
	//
	// 1.1.XX.XX
	HostPublicAddress *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	// The operating system of the host. Valid values:
	//
	// 	- **Linux**
	//
	// 	- **Windows**
	//
	// example:
	//
	// Linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The source of the host. Valid values:
	//
	// 	- **Local**: a host in a data center
	//
	// 	- **Ecs**: an ECS instance
	//
	// 	- **Rds**: a host in an ApsaraDB MyBase dedicated cluster
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the ECS instance or the host in an ApsaraDB MyBase dedicated cluster.
	//
	// > No value is returned for this parameter if the **Source*	- parameter is set to **Local**.
	//
	// example:
	//
	// i-bp19ienyt0yax748****
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The status of the host. Valid values:
	//
	// 	- **Normal**: normal
	//
	// 	- **Release**: released
	//
	// example:
	//
	// Normal
	SourceInstanceState *string `json:"SourceInstanceState,omitempty" xml:"SourceInstanceState,omitempty"`
}

func (s ListHostsResponseBodyHosts) String() string {
	return dara.Prettify(s)
}

func (s ListHostsResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *ListHostsResponseBodyHosts) GetActiveAddressType() *string {
	return s.ActiveAddressType
}

func (s *ListHostsResponseBodyHosts) GetComment() *string {
	return s.Comment
}

func (s *ListHostsResponseBodyHosts) GetHostAccountCount() *int32 {
	return s.HostAccountCount
}

func (s *ListHostsResponseBodyHosts) GetHostId() *string {
	return s.HostId
}

func (s *ListHostsResponseBodyHosts) GetHostName() *string {
	return s.HostName
}

func (s *ListHostsResponseBodyHosts) GetHostPrivateAddress() *string {
	return s.HostPrivateAddress
}

func (s *ListHostsResponseBodyHosts) GetHostPublicAddress() *string {
	return s.HostPublicAddress
}

func (s *ListHostsResponseBodyHosts) GetOSType() *string {
	return s.OSType
}

func (s *ListHostsResponseBodyHosts) GetSource() *string {
	return s.Source
}

func (s *ListHostsResponseBodyHosts) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *ListHostsResponseBodyHosts) GetSourceInstanceState() *string {
	return s.SourceInstanceState
}

func (s *ListHostsResponseBodyHosts) SetActiveAddressType(v string) *ListHostsResponseBodyHosts {
	s.ActiveAddressType = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetComment(v string) *ListHostsResponseBodyHosts {
	s.Comment = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetHostAccountCount(v int32) *ListHostsResponseBodyHosts {
	s.HostAccountCount = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetHostId(v string) *ListHostsResponseBodyHosts {
	s.HostId = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetHostName(v string) *ListHostsResponseBodyHosts {
	s.HostName = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetHostPrivateAddress(v string) *ListHostsResponseBodyHosts {
	s.HostPrivateAddress = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetHostPublicAddress(v string) *ListHostsResponseBodyHosts {
	s.HostPublicAddress = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetOSType(v string) *ListHostsResponseBodyHosts {
	s.OSType = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetSource(v string) *ListHostsResponseBodyHosts {
	s.Source = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetSourceInstanceId(v string) *ListHostsResponseBodyHosts {
	s.SourceInstanceId = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetSourceInstanceState(v string) *ListHostsResponseBodyHosts {
	s.SourceInstanceState = &v
	return s
}

func (s *ListHostsResponseBodyHosts) Validate() error {
	return dara.Validate(s)
}
