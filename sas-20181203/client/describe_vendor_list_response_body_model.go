// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVendorListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVendorListResponseBody
	GetRequestId() *string
	SetVendorNameList(v []*string) *DescribeVendorListResponseBody
	GetVendorNameList() []*string
}

type DescribeVendorListResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 340D7FC4-D575-1661-8ACD-CFA7BE57B795
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the service providers.
	VendorNameList []*string `json:"VendorNameList,omitempty" xml:"VendorNameList,omitempty" type:"Repeated"`
}

func (s DescribeVendorListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVendorListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVendorListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVendorListResponseBody) GetVendorNameList() []*string {
	return s.VendorNameList
}

func (s *DescribeVendorListResponseBody) SetRequestId(v string) *DescribeVendorListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVendorListResponseBody) SetVendorNameList(v []*string) *DescribeVendorListResponseBody {
	s.VendorNameList = v
	return s
}

func (s *DescribeVendorListResponseBody) Validate() error {
	return dara.Validate(s)
}
