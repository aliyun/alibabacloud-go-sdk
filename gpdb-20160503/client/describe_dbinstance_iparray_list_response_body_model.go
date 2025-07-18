// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceIPArrayListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDBInstanceIPArrayListResponseBodyItems) *DescribeDBInstanceIPArrayListResponseBody
	GetItems() *DescribeDBInstanceIPArrayListResponseBodyItems
	SetRequestId(v string) *DescribeDBInstanceIPArrayListResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceIPArrayListResponseBody struct {
	// The queried IP address whitelists.
	Items *DescribeDBInstanceIPArrayListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CB7AA0BF-BE41-480E-A3DC-C97BF85A391B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceIPArrayListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListResponseBody) GetItems() *DescribeDBInstanceIPArrayListResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstanceIPArrayListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceIPArrayListResponseBody) SetItems(v *DescribeDBInstanceIPArrayListResponseBodyItems) *DescribeDBInstanceIPArrayListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponseBody) SetRequestId(v string) *DescribeDBInstanceIPArrayListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceIPArrayListResponseBodyItems struct {
	DBInstanceIPArray []*DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray `json:"DBInstanceIPArray,omitempty" xml:"DBInstanceIPArray,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceIPArrayListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItems) GetDBInstanceIPArray() []*DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray {
	return s.DBInstanceIPArray
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItems) SetDBInstanceIPArray(v []*DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) *DescribeDBInstanceIPArrayListResponseBodyItems {
	s.DBInstanceIPArray = v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray struct {
	// The attribute of the IP address whitelist. By default, this parameter is empty. A whitelist with the `hidden` attribute is not displayed in the console.
	//
	// example:
	//
	// hidden
	DBInstanceIPArrayAttribute *string `json:"DBInstanceIPArrayAttribute,omitempty" xml:"DBInstanceIPArrayAttribute,omitempty"`
	// The name of the IP address whitelist.
	//
	// example:
	//
	// default
	DBInstanceIPArrayName *string `json:"DBInstanceIPArrayName,omitempty" xml:"DBInstanceIPArrayName,omitempty"`
	// The IP addresses listed in the whitelist. Up to 1,000 IP addresses are contained in a whitelist and separated by commas (,). The IP addresses must use one of the following formats:
	//
	// 	- 0.0.0.0/0
	//
	// 	- 10.23.12.24. This is a standard IP address.
	//
	// 	- 10.23.12.24/24. This is a CIDR block. The value `/24` indicates that the prefix of the CIDR block is 24-bit long. You can replace 24 with a value in the range of `1 to 32`.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) GetDBInstanceIPArrayAttribute() *string {
	return s.DBInstanceIPArrayAttribute
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) GetDBInstanceIPArrayName() *string {
	return s.DBInstanceIPArrayName
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) SetDBInstanceIPArrayAttribute(v string) *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray {
	s.DBInstanceIPArrayAttribute = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) SetDBInstanceIPArrayName(v string) *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray {
	s.DBInstanceIPArrayName = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) SetSecurityIPList(v string) *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray) Validate() error {
	return dara.Validate(s)
}
