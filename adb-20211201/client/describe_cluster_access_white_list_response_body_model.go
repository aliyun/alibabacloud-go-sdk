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
	// The queried IP address whitelists.
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
	// The attribute of the IP address whitelist.
	//
	// >  The IP address whitelists that have the **hidden*	- attribute are not displayed in the console. These IP address whitelists are used to access services such as Data Transmission Service (DTS) and PolarDB.
	//
	// example:
	//
	// hidden
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	// The name of the IP address whitelist.
	//
	// Each cluster supports up to 50 IP address whitelists.
	//
	// example:
	//
	// test
	DBClusterIPArrayName *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	// The IP addresses in the IP address whitelist. Up to 500 IP addresses can be returned. Multiple IP addresses are separated by commas (,).
	//
	// example:
	//
	// 127.0.xx.xx
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
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
