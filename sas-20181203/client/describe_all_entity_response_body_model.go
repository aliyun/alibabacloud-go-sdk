// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntityList(v []*DescribeAllEntityResponseBodyEntityList) *DescribeAllEntityResponseBody
	GetEntityList() []*DescribeAllEntityResponseBodyEntityList
	SetRequestId(v string) *DescribeAllEntityResponseBody
	GetRequestId() *string
}

type DescribeAllEntityResponseBody struct {
	// An array that consists of servers.
	EntityList []*DescribeAllEntityResponseBodyEntityList `json:"EntityList,omitempty" xml:"EntityList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAllEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllEntityResponseBody) GetEntityList() []*DescribeAllEntityResponseBodyEntityList {
	return s.EntityList
}

func (s *DescribeAllEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAllEntityResponseBody) SetEntityList(v []*DescribeAllEntityResponseBodyEntityList) *DescribeAllEntityResponseBody {
	s.EntityList = v
	return s
}

func (s *DescribeAllEntityResponseBody) SetRequestId(v string) *DescribeAllEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllEntityResponseBody) Validate() error {
	if s.EntityList != nil {
		for _, item := range s.EntityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAllEntityResponseBodyEntityList struct {
	// The ID of the asset group.
	//
	// example:
	//
	// 281801
	GroupId *int32 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// abc
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 172.19.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 100.104.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The IP address of the server.
	//
	// example:
	//
	// 101.132.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The operating system of the server. Valid values:
	//
	// 	- **linux**
	//
	// 	- **windows**
	//
	// example:
	//
	// windows
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 4fe8e1cd-3c37-4851-b9de-124da32c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeAllEntityResponseBodyEntityList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllEntityResponseBodyEntityList) GoString() string {
	return s.String()
}

func (s *DescribeAllEntityResponseBodyEntityList) GetGroupId() *int32 {
	return s.GroupId
}

func (s *DescribeAllEntityResponseBodyEntityList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAllEntityResponseBodyEntityList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeAllEntityResponseBodyEntityList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeAllEntityResponseBodyEntityList) GetIp() *string {
	return s.Ip
}

func (s *DescribeAllEntityResponseBodyEntityList) GetOs() *string {
	return s.Os
}

func (s *DescribeAllEntityResponseBodyEntityList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeAllEntityResponseBodyEntityList) SetGroupId(v int32) *DescribeAllEntityResponseBodyEntityList {
	s.GroupId = &v
	return s
}

func (s *DescribeAllEntityResponseBodyEntityList) SetInstanceName(v string) *DescribeAllEntityResponseBodyEntityList {
	s.InstanceName = &v
	return s
}

func (s *DescribeAllEntityResponseBodyEntityList) SetInternetIp(v string) *DescribeAllEntityResponseBodyEntityList {
	s.InternetIp = &v
	return s
}

func (s *DescribeAllEntityResponseBodyEntityList) SetIntranetIp(v string) *DescribeAllEntityResponseBodyEntityList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeAllEntityResponseBodyEntityList) SetIp(v string) *DescribeAllEntityResponseBodyEntityList {
	s.Ip = &v
	return s
}

func (s *DescribeAllEntityResponseBodyEntityList) SetOs(v string) *DescribeAllEntityResponseBodyEntityList {
	s.Os = &v
	return s
}

func (s *DescribeAllEntityResponseBodyEntityList) SetUuid(v string) *DescribeAllEntityResponseBodyEntityList {
	s.Uuid = &v
	return s
}

func (s *DescribeAllEntityResponseBodyEntityList) Validate() error {
	return dara.Validate(s)
}
