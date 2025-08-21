// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEpnInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEPNInstanceId(v string) *DescribeEpnInstancesRequest
	GetEPNInstanceId() *string
	SetEPNInstanceName(v string) *DescribeEpnInstancesRequest
	GetEPNInstanceName() *string
	SetPageNumber(v int32) *DescribeEpnInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEpnInstancesRequest
	GetPageSize() *int32
}

type DescribeEpnInstancesRequest struct {
	// The version number.
	//
	// example:
	//
	// 2017-11-10
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	// The name of the EPN instance.
	//
	// example:
	//
	// testEPNInstanceName
	EPNInstanceName *string `json:"EPNInstanceName,omitempty" xml:"EPNInstanceName,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeEpnInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEpnInstancesRequest) GetEPNInstanceId() *string {
	return s.EPNInstanceId
}

func (s *DescribeEpnInstancesRequest) GetEPNInstanceName() *string {
	return s.EPNInstanceName
}

func (s *DescribeEpnInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEpnInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEpnInstancesRequest) SetEPNInstanceId(v string) *DescribeEpnInstancesRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *DescribeEpnInstancesRequest) SetEPNInstanceName(v string) *DescribeEpnInstancesRequest {
	s.EPNInstanceName = &v
	return s
}

func (s *DescribeEpnInstancesRequest) SetPageNumber(v int32) *DescribeEpnInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEpnInstancesRequest) SetPageSize(v int32) *DescribeEpnInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEpnInstancesRequest) Validate() error {
	return dara.Validate(s)
}
