// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityIPListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeSecurityIPListResponseBodyData) *DescribeSecurityIPListResponseBody
	GetData() *DescribeSecurityIPListResponseBodyData
	SetRequestId(v string) *DescribeSecurityIPListResponseBody
	GetRequestId() *string
}

type DescribeSecurityIPListResponseBody struct {
	// The data returned.
	Data *DescribeSecurityIPListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSecurityIPListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPListResponseBody) GetData() *DescribeSecurityIPListResponseBodyData {
	return s.Data
}

func (s *DescribeSecurityIPListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityIPListResponseBody) SetData(v *DescribeSecurityIPListResponseBodyData) *DescribeSecurityIPListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSecurityIPListResponseBody) SetRequestId(v string) *DescribeSecurityIPListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityIPListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityIPListResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// TestCluster
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The details about the whitelists.
	GroupItems []*DescribeSecurityIPListResponseBodyDataGroupItems `json:"GroupItems,omitempty" xml:"GroupItems,omitempty" type:"Repeated"`
}

func (s DescribeSecurityIPListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPListResponseBodyData) GetDBInstanceID() *int32 {
	return s.DBInstanceID
}

func (s *DescribeSecurityIPListResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSecurityIPListResponseBodyData) GetGroupItems() []*DescribeSecurityIPListResponseBodyDataGroupItems {
	return s.GroupItems
}

func (s *DescribeSecurityIPListResponseBodyData) SetDBInstanceID(v int32) *DescribeSecurityIPListResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyData) SetDBInstanceName(v string) *DescribeSecurityIPListResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyData) SetGroupItems(v []*DescribeSecurityIPListResponseBodyDataGroupItems) *DescribeSecurityIPListResponseBodyData {
	s.GroupItems = v
	return s
}

func (s *DescribeSecurityIPListResponseBodyData) Validate() error {
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

type DescribeSecurityIPListResponseBodyDataGroupItems struct {
	// The name of the whitelist.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The tag of the whitelist.
	//
	// example:
	//
	// test
	GroupTag *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	// The IP addresses and CIDR blocks in the whitelist.
	//
	// example:
	//
	// 127.0.XX.XX
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The IP address type.
	//
	// example:
	//
	// ipv4
	SecurityIPType *string `json:"SecurityIPType,omitempty" xml:"SecurityIPType,omitempty"`
	// The network type of the whitelist.
	//
	// example:
	//
	// mix
	WhitelistNetType *string `json:"WhitelistNetType,omitempty" xml:"WhitelistNetType,omitempty"`
}

func (s DescribeSecurityIPListResponseBodyDataGroupItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPListResponseBodyDataGroupItems) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPListResponseBodyDataGroupItems) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeSecurityIPListResponseBodyDataGroupItems) GetGroupTag() *string {
	return s.GroupTag
}

func (s *DescribeSecurityIPListResponseBodyDataGroupItems) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeSecurityIPListResponseBodyDataGroupItems) GetSecurityIPType() *string {
	return s.SecurityIPType
}

func (s *DescribeSecurityIPListResponseBodyDataGroupItems) GetWhitelistNetType() *string {
	return s.WhitelistNetType
}

func (s *DescribeSecurityIPListResponseBodyDataGroupItems) SetGroupName(v string) *DescribeSecurityIPListResponseBodyDataGroupItems {
	s.GroupName = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyDataGroupItems) SetGroupTag(v string) *DescribeSecurityIPListResponseBodyDataGroupItems {
	s.GroupTag = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyDataGroupItems) SetSecurityIPList(v string) *DescribeSecurityIPListResponseBodyDataGroupItems {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyDataGroupItems) SetSecurityIPType(v string) *DescribeSecurityIPListResponseBodyDataGroupItems {
	s.SecurityIPType = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyDataGroupItems) SetWhitelistNetType(v string) *DescribeSecurityIPListResponseBodyDataGroupItems {
	s.WhitelistNetType = &v
	return s
}

func (s *DescribeSecurityIPListResponseBodyDataGroupItems) Validate() error {
	return dara.Validate(s)
}
