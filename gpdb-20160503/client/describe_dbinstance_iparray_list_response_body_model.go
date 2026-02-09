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
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.DBInstanceIPArray != nil {
		for _, item := range s.DBInstanceIPArray {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceIPArrayListResponseBodyItemsDBInstanceIPArray struct {
	DBInstanceIPArrayAttribute *string `json:"DBInstanceIPArrayAttribute,omitempty" xml:"DBInstanceIPArrayAttribute,omitempty"`
	DBInstanceIPArrayName      *string `json:"DBInstanceIPArrayName,omitempty" xml:"DBInstanceIPArrayName,omitempty"`
	SecurityIPList             *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
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
