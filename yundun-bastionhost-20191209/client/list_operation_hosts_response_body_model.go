// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationHostsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHosts(v []*ListOperationHostsResponseBodyHosts) *ListOperationHostsResponseBody
	GetHosts() []*ListOperationHostsResponseBodyHosts
	SetRequestId(v string) *ListOperationHostsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListOperationHostsResponseBody
	GetTotalCount() *int64
}

type ListOperationHostsResponseBody struct {
	// The hosts returned.
	Hosts []*ListOperationHostsResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4F6C075F-FC86-476E-943B-097BD4E12948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of hosts returned.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOperationHostsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOperationHostsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationHostsResponseBody) GetHosts() []*ListOperationHostsResponseBodyHosts {
	return s.Hosts
}

func (s *ListOperationHostsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOperationHostsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOperationHostsResponseBody) SetHosts(v []*ListOperationHostsResponseBodyHosts) *ListOperationHostsResponseBody {
	s.Hosts = v
	return s
}

func (s *ListOperationHostsResponseBody) SetRequestId(v string) *ListOperationHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationHostsResponseBody) SetTotalCount(v int64) *ListOperationHostsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOperationHostsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOperationHostsResponseBodyHosts struct {
	// The address type of the host. Valid values:
	//
	// 	- **Public**
	//
	// 	- **Private**
	//
	// example:
	//
	// Public
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The remarks of the host.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The host ID.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The host name.
	//
	// example:
	//
	// host1
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The private IP address of the host.
	//
	// example:
	//
	// 192.168.XX.XX
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	// The public IP address of the host.
	//
	// example:
	//
	// 10.158.XX.XX
	HostPublicAddress *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	// The host OS.
	//
	// 	- **Linux**
	//
	// 	- **Windows**
	//
	// example:
	//
	// Linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The host type. Valid values:
	//
	// 	- **Local**: on-premises host.
	//
	// 	- **Ecs**: Elastic Compute Service (ECS) instance.
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ECS instance ID.
	//
	// example:
	//
	// i-bp19ienyt0yax748****
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The host status. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **Release**
	//
	// example:
	//
	// Normal
	SourceInstanceState *string `json:"SourceInstanceState,omitempty" xml:"SourceInstanceState,omitempty"`
}

func (s ListOperationHostsResponseBodyHosts) String() string {
	return dara.Prettify(s)
}

func (s ListOperationHostsResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *ListOperationHostsResponseBodyHosts) GetActiveAddressType() *string {
	return s.ActiveAddressType
}

func (s *ListOperationHostsResponseBodyHosts) GetComment() *string {
	return s.Comment
}

func (s *ListOperationHostsResponseBodyHosts) GetHostId() *string {
	return s.HostId
}

func (s *ListOperationHostsResponseBodyHosts) GetHostName() *string {
	return s.HostName
}

func (s *ListOperationHostsResponseBodyHosts) GetHostPrivateAddress() *string {
	return s.HostPrivateAddress
}

func (s *ListOperationHostsResponseBodyHosts) GetHostPublicAddress() *string {
	return s.HostPublicAddress
}

func (s *ListOperationHostsResponseBodyHosts) GetOSType() *string {
	return s.OSType
}

func (s *ListOperationHostsResponseBodyHosts) GetSource() *string {
	return s.Source
}

func (s *ListOperationHostsResponseBodyHosts) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *ListOperationHostsResponseBodyHosts) GetSourceInstanceState() *string {
	return s.SourceInstanceState
}

func (s *ListOperationHostsResponseBodyHosts) SetActiveAddressType(v string) *ListOperationHostsResponseBodyHosts {
	s.ActiveAddressType = &v
	return s
}

func (s *ListOperationHostsResponseBodyHosts) SetComment(v string) *ListOperationHostsResponseBodyHosts {
	s.Comment = &v
	return s
}

func (s *ListOperationHostsResponseBodyHosts) SetHostId(v string) *ListOperationHostsResponseBodyHosts {
	s.HostId = &v
	return s
}

func (s *ListOperationHostsResponseBodyHosts) SetHostName(v string) *ListOperationHostsResponseBodyHosts {
	s.HostName = &v
	return s
}

func (s *ListOperationHostsResponseBodyHosts) SetHostPrivateAddress(v string) *ListOperationHostsResponseBodyHosts {
	s.HostPrivateAddress = &v
	return s
}

func (s *ListOperationHostsResponseBodyHosts) SetHostPublicAddress(v string) *ListOperationHostsResponseBodyHosts {
	s.HostPublicAddress = &v
	return s
}

func (s *ListOperationHostsResponseBodyHosts) SetOSType(v string) *ListOperationHostsResponseBodyHosts {
	s.OSType = &v
	return s
}

func (s *ListOperationHostsResponseBodyHosts) SetSource(v string) *ListOperationHostsResponseBodyHosts {
	s.Source = &v
	return s
}

func (s *ListOperationHostsResponseBodyHosts) SetSourceInstanceId(v string) *ListOperationHostsResponseBodyHosts {
	s.SourceInstanceId = &v
	return s
}

func (s *ListOperationHostsResponseBodyHosts) SetSourceInstanceState(v string) *ListOperationHostsResponseBodyHosts {
	s.SourceInstanceState = &v
	return s
}

func (s *ListOperationHostsResponseBodyHosts) Validate() error {
	return dara.Validate(s)
}
