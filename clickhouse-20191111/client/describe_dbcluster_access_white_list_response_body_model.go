// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAccessWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterAccessWhiteList(v *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList) *DescribeDBClusterAccessWhiteListResponseBody
	GetDBClusterAccessWhiteList() *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList
	SetRequestId(v string) *DescribeDBClusterAccessWhiteListResponseBody
	GetRequestId() *string
}

type DescribeDBClusterAccessWhiteListResponseBody struct {
	// The details about the IP address whitelist.
	DBClusterAccessWhiteList *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList `json:"DBClusterAccessWhiteList,omitempty" xml:"DBClusterAccessWhiteList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 905F13A4-5097-4897-A84D-527EC75FFF4F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterAccessWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) GetDBClusterAccessWhiteList() *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList {
	return s.DBClusterAccessWhiteList
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) SetDBClusterAccessWhiteList(v *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList) *DescribeDBClusterAccessWhiteListResponseBody {
	s.DBClusterAccessWhiteList = v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) SetRequestId(v string) *DescribeDBClusterAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) Validate() error {
	if s.DBClusterAccessWhiteList != nil {
		if err := s.DBClusterAccessWhiteList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList struct {
	IPArray []*DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray `json:"IPArray,omitempty" xml:"IPArray,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList) GetIPArray() []*DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray {
	return s.IPArray
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList) SetIPArray(v []*DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray) *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList {
	s.IPArray = v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteList) Validate() error {
	if s.IPArray != nil {
		for _, item := range s.IPArray {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray struct {
	// The attribute of the IP address whitelist.
	//
	// example:
	//
	// default
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	// The name of the IP address whitelist.
	//
	// example:
	//
	// default
	DBClusterIPArrayName *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	// The IP addresses in the IP address whitelist.
	//
	// example:
	//
	// 192.168.xx.xx,192.168.xx.xx
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray) GetDBClusterIPArrayAttribute() *string {
	return s.DBClusterIPArrayAttribute
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray) GetDBClusterIPArrayName() *string {
	return s.DBClusterIPArrayName
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray) SetDBClusterIPArrayAttribute(v string) *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray) SetDBClusterIPArrayName(v string) *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray) SetSecurityIPList(v string) *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyDBClusterAccessWhiteListIPArray) Validate() error {
	return dara.Validate(s)
}
