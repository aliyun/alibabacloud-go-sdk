// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostsForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHosts(v []*ListHostsForUserResponseBodyHosts) *ListHostsForUserResponseBody
	GetHosts() []*ListHostsForUserResponseBodyHosts
	SetRequestId(v string) *ListHostsForUserResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHostsForUserResponseBody
	GetTotalCount() *int32
}

type ListHostsForUserResponseBody struct {
	// The hosts returned.
	Hosts []*ListHostsForUserResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The request ID.
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

func (s ListHostsForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHostsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostsForUserResponseBody) GetHosts() []*ListHostsForUserResponseBodyHosts {
	return s.Hosts
}

func (s *ListHostsForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHostsForUserResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHostsForUserResponseBody) SetHosts(v []*ListHostsForUserResponseBodyHosts) *ListHostsForUserResponseBody {
	s.Hosts = v
	return s
}

func (s *ListHostsForUserResponseBody) SetRequestId(v string) *ListHostsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostsForUserResponseBody) SetTotalCount(v int32) *ListHostsForUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostsForUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHostsForUserResponseBodyHosts struct {
	// The endpoint type of the host. Valid values:
	//
	// 	- **Public**: public endpoint
	//
	// 	- **Private**: internal endpoint
	//
	// example:
	//
	// Public
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The description of the host.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
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
	// host01
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The internal endpoint of the host. The value is a domain name or an IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	// The public endpoint of the host. The value is a domain name or an IP address.
	//
	// example:
	//
	// 10.158.XX.XX
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
}

func (s ListHostsForUserResponseBodyHosts) String() string {
	return dara.Prettify(s)
}

func (s ListHostsForUserResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *ListHostsForUserResponseBodyHosts) GetActiveAddressType() *string {
	return s.ActiveAddressType
}

func (s *ListHostsForUserResponseBodyHosts) GetComment() *string {
	return s.Comment
}

func (s *ListHostsForUserResponseBodyHosts) GetHostId() *string {
	return s.HostId
}

func (s *ListHostsForUserResponseBodyHosts) GetHostName() *string {
	return s.HostName
}

func (s *ListHostsForUserResponseBodyHosts) GetHostPrivateAddress() *string {
	return s.HostPrivateAddress
}

func (s *ListHostsForUserResponseBodyHosts) GetHostPublicAddress() *string {
	return s.HostPublicAddress
}

func (s *ListHostsForUserResponseBodyHosts) GetOSType() *string {
	return s.OSType
}

func (s *ListHostsForUserResponseBodyHosts) SetActiveAddressType(v string) *ListHostsForUserResponseBodyHosts {
	s.ActiveAddressType = &v
	return s
}

func (s *ListHostsForUserResponseBodyHosts) SetComment(v string) *ListHostsForUserResponseBodyHosts {
	s.Comment = &v
	return s
}

func (s *ListHostsForUserResponseBodyHosts) SetHostId(v string) *ListHostsForUserResponseBodyHosts {
	s.HostId = &v
	return s
}

func (s *ListHostsForUserResponseBodyHosts) SetHostName(v string) *ListHostsForUserResponseBodyHosts {
	s.HostName = &v
	return s
}

func (s *ListHostsForUserResponseBodyHosts) SetHostPrivateAddress(v string) *ListHostsForUserResponseBodyHosts {
	s.HostPrivateAddress = &v
	return s
}

func (s *ListHostsForUserResponseBodyHosts) SetHostPublicAddress(v string) *ListHostsForUserResponseBodyHosts {
	s.HostPublicAddress = &v
	return s
}

func (s *ListHostsForUserResponseBodyHosts) SetOSType(v string) *ListHostsForUserResponseBodyHosts {
	s.OSType = &v
	return s
}

func (s *ListHostsForUserResponseBodyHosts) Validate() error {
	return dara.Validate(s)
}
