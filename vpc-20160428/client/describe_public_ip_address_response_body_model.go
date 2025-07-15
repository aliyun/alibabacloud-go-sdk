// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePublicIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribePublicIpAddressResponseBody
	GetCode() *string
	SetMessage(v string) *DescribePublicIpAddressResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribePublicIpAddressResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePublicIpAddressResponseBody
	GetPageSize() *int32
	SetPublicIpAddress(v []*string) *DescribePublicIpAddressResponseBody
	GetPublicIpAddress() []*string
	SetRegionId(v string) *DescribePublicIpAddressResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribePublicIpAddressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribePublicIpAddressResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribePublicIpAddressResponseBody
	GetTotalCount() *int32
}

type DescribePublicIpAddressResponseBody struct {
	// The HTTP status codes returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response messages.
	//
	// example:
	//
	// test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The range of the public IP addresses of the VPC in the region.
	//
	// example:
	//
	// 110.11.1.0/24
	PublicIpAddress []*string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty" type:"Repeated"`
	// The ID of the region to which the public IP addresses belong.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePublicIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePublicIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePublicIpAddressResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePublicIpAddressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePublicIpAddressResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePublicIpAddressResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePublicIpAddressResponseBody) GetPublicIpAddress() []*string {
	return s.PublicIpAddress
}

func (s *DescribePublicIpAddressResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePublicIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePublicIpAddressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribePublicIpAddressResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePublicIpAddressResponseBody) SetCode(v string) *DescribePublicIpAddressResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePublicIpAddressResponseBody) SetMessage(v string) *DescribePublicIpAddressResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePublicIpAddressResponseBody) SetPageNumber(v int32) *DescribePublicIpAddressResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePublicIpAddressResponseBody) SetPageSize(v int32) *DescribePublicIpAddressResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePublicIpAddressResponseBody) SetPublicIpAddress(v []*string) *DescribePublicIpAddressResponseBody {
	s.PublicIpAddress = v
	return s
}

func (s *DescribePublicIpAddressResponseBody) SetRegionId(v string) *DescribePublicIpAddressResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribePublicIpAddressResponseBody) SetRequestId(v string) *DescribePublicIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePublicIpAddressResponseBody) SetSuccess(v bool) *DescribePublicIpAddressResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePublicIpAddressResponseBody) SetTotalCount(v int32) *DescribePublicIpAddressResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePublicIpAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
