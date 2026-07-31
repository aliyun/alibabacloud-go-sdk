// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAccessWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeClusterAccessWhiteListResponseBodyItems) *DescribeClusterAccessWhiteListResponseBody
	GetItems() *DescribeClusterAccessWhiteListResponseBodyItems
	SetRequestId(v string) *DescribeClusterAccessWhiteListResponseBody
	GetRequestId() *string
}

type DescribeClusterAccessWhiteListResponseBody struct {
	Items *DescribeClusterAccessWhiteListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 370D09FD-442A-5225-AAD3-7362CAE39177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterAccessWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAccessWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterAccessWhiteListResponseBody) GetItems() *DescribeClusterAccessWhiteListResponseBodyItems {
	return s.Items
}

func (s *DescribeClusterAccessWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterAccessWhiteListResponseBody) SetItems(v *DescribeClusterAccessWhiteListResponseBodyItems) *DescribeClusterAccessWhiteListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeClusterAccessWhiteListResponseBody) SetRequestId(v string) *DescribeClusterAccessWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterAccessWhiteListResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterAccessWhiteListResponseBodyItems struct {
	IPArray []*DescribeClusterAccessWhiteListResponseBodyItemsIPArray `json:"IPArray,omitempty" xml:"IPArray,omitempty" type:"Repeated"`
}

func (s DescribeClusterAccessWhiteListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAccessWhiteListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeClusterAccessWhiteListResponseBodyItems) GetIPArray() []*DescribeClusterAccessWhiteListResponseBodyItemsIPArray {
	return s.IPArray
}

func (s *DescribeClusterAccessWhiteListResponseBodyItems) SetIPArray(v []*DescribeClusterAccessWhiteListResponseBodyItemsIPArray) *DescribeClusterAccessWhiteListResponseBodyItems {
	s.IPArray = v
	return s
}

func (s *DescribeClusterAccessWhiteListResponseBodyItems) Validate() error {
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

type DescribeClusterAccessWhiteListResponseBodyItemsIPArray struct {
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	DBClusterIPArrayName      *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	SecurityIPList            *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s DescribeClusterAccessWhiteListResponseBodyItemsIPArray) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAccessWhiteListResponseBodyItemsIPArray) GoString() string {
	return s.String()
}

func (s *DescribeClusterAccessWhiteListResponseBodyItemsIPArray) GetDBClusterIPArrayAttribute() *string {
	return s.DBClusterIPArrayAttribute
}

func (s *DescribeClusterAccessWhiteListResponseBodyItemsIPArray) GetDBClusterIPArrayName() *string {
	return s.DBClusterIPArrayName
}

func (s *DescribeClusterAccessWhiteListResponseBodyItemsIPArray) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeClusterAccessWhiteListResponseBodyItemsIPArray) SetDBClusterIPArrayAttribute(v string) *DescribeClusterAccessWhiteListResponseBodyItemsIPArray {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *DescribeClusterAccessWhiteListResponseBodyItemsIPArray) SetDBClusterIPArrayName(v string) *DescribeClusterAccessWhiteListResponseBodyItemsIPArray {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *DescribeClusterAccessWhiteListResponseBodyItemsIPArray) SetSecurityIPList(v string) *DescribeClusterAccessWhiteListResponseBodyItemsIPArray {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeClusterAccessWhiteListResponseBodyItemsIPArray) Validate() error {
	return dara.Validate(s)
}
