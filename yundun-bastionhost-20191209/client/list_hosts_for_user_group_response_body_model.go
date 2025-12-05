// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostsForUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHosts(v []*ListHostsForUserGroupResponseBodyHosts) *ListHostsForUserGroupResponseBody
	GetHosts() []*ListHostsForUserGroupResponseBodyHosts
	SetRequestId(v string) *ListHostsForUserGroupResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHostsForUserGroupResponseBody
	GetTotalCount() *int32
}

type ListHostsForUserGroupResponseBody struct {
	// The hosts returned.
	Hosts []*ListHostsForUserGroupResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
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
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHostsForUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHostsForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostsForUserGroupResponseBody) GetHosts() []*ListHostsForUserGroupResponseBodyHosts {
	return s.Hosts
}

func (s *ListHostsForUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHostsForUserGroupResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHostsForUserGroupResponseBody) SetHosts(v []*ListHostsForUserGroupResponseBodyHosts) *ListHostsForUserGroupResponseBody {
	s.Hosts = v
	return s
}

func (s *ListHostsForUserGroupResponseBody) SetRequestId(v string) *ListHostsForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostsForUserGroupResponseBody) SetTotalCount(v int32) *ListHostsForUserGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostsForUserGroupResponseBody) Validate() error {
	if s.Hosts != nil {
		for _, item := range s.Hosts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHostsForUserGroupResponseBodyHosts struct {
	// The address type of the host. Valid values:
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
	// host1
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

func (s ListHostsForUserGroupResponseBodyHosts) String() string {
	return dara.Prettify(s)
}

func (s ListHostsForUserGroupResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *ListHostsForUserGroupResponseBodyHosts) GetActiveAddressType() *string {
	return s.ActiveAddressType
}

func (s *ListHostsForUserGroupResponseBodyHosts) GetComment() *string {
	return s.Comment
}

func (s *ListHostsForUserGroupResponseBodyHosts) GetHostId() *string {
	return s.HostId
}

func (s *ListHostsForUserGroupResponseBodyHosts) GetHostName() *string {
	return s.HostName
}

func (s *ListHostsForUserGroupResponseBodyHosts) GetHostPrivateAddress() *string {
	return s.HostPrivateAddress
}

func (s *ListHostsForUserGroupResponseBodyHosts) GetHostPublicAddress() *string {
	return s.HostPublicAddress
}

func (s *ListHostsForUserGroupResponseBodyHosts) GetOSType() *string {
	return s.OSType
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetActiveAddressType(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.ActiveAddressType = &v
	return s
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetComment(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.Comment = &v
	return s
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetHostId(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.HostId = &v
	return s
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetHostName(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.HostName = &v
	return s
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetHostPrivateAddress(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.HostPrivateAddress = &v
	return s
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetHostPublicAddress(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.HostPublicAddress = &v
	return s
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetOSType(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.OSType = &v
	return s
}

func (s *ListHostsForUserGroupResponseBodyHosts) Validate() error {
	return dara.Validate(s)
}
