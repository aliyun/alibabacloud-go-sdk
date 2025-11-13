// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityIPListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeSecurityIPListResponseBody
	GetDBInstanceName() *string
	SetGroupItems(v []*DescribeSecurityIPListResponseBodyGroupItems) *DescribeSecurityIPListResponseBody
	GetGroupItems() []*DescribeSecurityIPListResponseBodyGroupItems
	SetRequestId(v string) *DescribeSecurityIPListResponseBody
	GetRequestId() *string
}

type DescribeSecurityIPListResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The details about each IP address whitelist returned.
	GroupItems []*DescribeSecurityIPListResponseBodyGroupItems `json:"GroupItems,omitempty" xml:"GroupItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5CBE044D-4594-525D-AC65-E988553D853E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSecurityIPListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPListResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSecurityIPListResponseBody) GetGroupItems() []*DescribeSecurityIPListResponseBodyGroupItems {
	return s.GroupItems
}

func (s *DescribeSecurityIPListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityIPListResponseBody) SetDBInstanceName(v string) *DescribeSecurityIPListResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSecurityIPListResponseBody) SetGroupItems(v []*DescribeSecurityIPListResponseBodyGroupItems) *DescribeSecurityIPListResponseBody {
	s.GroupItems = v
	return s
}

func (s *DescribeSecurityIPListResponseBody) SetRequestId(v string) *DescribeSecurityIPListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityIPListResponseBody) Validate() error {
	if s.GroupItems != nil {
		for _, item := range s.GroupItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityIPListResponseBodyGroupItems struct {
	// The IP address type. Valid values:
	//
	// 	- ipv4
	//
	// 	- ipv6 (not supported)
	//
	// example:
	//
	// ipv4
	AecurityIPType *string `json:"AecurityIPType,omitempty" xml:"AecurityIPType,omitempty"`
	// The name of the whitelist.
	//
	// example:
	//
	// group1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The tag of the whitelist.
	//
	// example:
	//
	// ""
	GroupTag *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	// The IP addresses in the whitelist. Multiple IP addresses are separated by commas (,).
	//
	// example:
	//
	// 127.0.XX.XX
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The network type of the whitelist.
	//
	// example:
	//
	// mix
	WhitelistNetType *string `json:"WhitelistNetType,omitempty" xml:"WhitelistNetType,omitempty"`
}

func (s DescribeSecurityIPListResponseBodyGroupItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPListResponseBodyGroupItems) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPListResponseBodyGroupItems) GetAecurityIPType() *string {
	return s.AecurityIPType
}

func (s *DescribeSecurityIPListResponseBodyGroupItems) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeSecurityIPListResponseBodyGroupItems) GetGroupTag() *string {
	return s.GroupTag
}

func (s *DescribeSecurityIPListResponseBodyGroupItems) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeSecurityIPListResponseBodyGroupItems) GetWhitelistNetType() *string {
	return s.WhitelistNetType
}

func (s *DescribeSecurityIPListResponseBodyGroupItems) SetAecurityIPType(v string) *DescribeSecurityIPListResponseBodyGroupItems {
	s.AecurityIPType = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyGroupItems) SetGroupName(v string) *DescribeSecurityIPListResponseBodyGroupItems {
	s.GroupName = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyGroupItems) SetGroupTag(v string) *DescribeSecurityIPListResponseBodyGroupItems {
	s.GroupTag = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyGroupItems) SetSecurityIPList(v string) *DescribeSecurityIPListResponseBodyGroupItems {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyGroupItems) SetWhitelistNetType(v string) *DescribeSecurityIPListResponseBodyGroupItems {
	s.WhitelistNetType = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyGroupItems) Validate() error {
	return dara.Validate(s)
}
