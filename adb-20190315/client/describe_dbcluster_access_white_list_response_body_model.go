// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAccessWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDBClusterAccessWhiteListResponseBodyItems) *DescribeDBClusterAccessWhiteListResponseBody
	GetItems() *DescribeDBClusterAccessWhiteListResponseBodyItems
	SetRequestId(v string) *DescribeDBClusterAccessWhiteListResponseBody
	GetRequestId() *string
}

type DescribeDBClusterAccessWhiteListResponseBody struct {
	Items *DescribeDBClusterAccessWhiteListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterAccessWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) GetItems() *DescribeDBClusterAccessWhiteListResponseBodyItems {
	return s.Items
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) SetItems(v *DescribeDBClusterAccessWhiteListResponseBodyItems) *DescribeDBClusterAccessWhiteListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) SetRequestId(v string) *DescribeDBClusterAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterAccessWhiteListResponseBodyItems struct {
	IPArray []*DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray `json:"IPArray,omitempty" xml:"IPArray,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAccessWhiteListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItems) GetIPArray() []*DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray {
	return s.IPArray
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItems) SetIPArray(v []*DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) *DescribeDBClusterAccessWhiteListResponseBodyItems {
	s.IPArray = v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItems) Validate() error {
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

type DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray struct {
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	DBClusterIPArrayName      *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	SecurityIPList            *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) GetDBClusterIPArrayAttribute() *string {
	return s.DBClusterIPArrayAttribute
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) GetDBClusterIPArrayName() *string {
	return s.DBClusterIPArrayName
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) SetDBClusterIPArrayAttribute(v string) *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) SetDBClusterIPArrayName(v string) *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) SetSecurityIPList(v string) *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeDBClusterAccessWhiteListResponseBodyItemsIPArray) Validate() error {
	return dara.Validate(s)
}
