// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVaultReplicationRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeVaultReplicationRegionsResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeVaultReplicationRegionsResponseBody
	GetMessage() *string
	SetRegions(v *DescribeVaultReplicationRegionsResponseBodyRegions) *DescribeVaultReplicationRegionsResponseBody
	GetRegions() *DescribeVaultReplicationRegionsResponseBodyRegions
	SetRequestId(v string) *DescribeVaultReplicationRegionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeVaultReplicationRegionsResponseBody
	GetSuccess() *bool
}

type DescribeVaultReplicationRegionsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The regions that support cross-region replication.
	Regions *DescribeVaultReplicationRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F4EEB401-DD21-588D-AE3B-1E835C7655E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeVaultReplicationRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultReplicationRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVaultReplicationRegionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeVaultReplicationRegionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeVaultReplicationRegionsResponseBody) GetRegions() *DescribeVaultReplicationRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeVaultReplicationRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVaultReplicationRegionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeVaultReplicationRegionsResponseBody) SetCode(v string) *DescribeVaultReplicationRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeVaultReplicationRegionsResponseBody) SetMessage(v string) *DescribeVaultReplicationRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeVaultReplicationRegionsResponseBody) SetRegions(v *DescribeVaultReplicationRegionsResponseBodyRegions) *DescribeVaultReplicationRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeVaultReplicationRegionsResponseBody) SetRequestId(v string) *DescribeVaultReplicationRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVaultReplicationRegionsResponseBody) SetSuccess(v bool) *DescribeVaultReplicationRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeVaultReplicationRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVaultReplicationRegionsResponseBodyRegions struct {
	RegionId []*string `json:"RegionId,omitempty" xml:"RegionId,omitempty" type:"Repeated"`
}

func (s DescribeVaultReplicationRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultReplicationRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeVaultReplicationRegionsResponseBodyRegions) GetRegionId() []*string {
	return s.RegionId
}

func (s *DescribeVaultReplicationRegionsResponseBodyRegions) SetRegionId(v []*string) *DescribeVaultReplicationRegionsResponseBodyRegions {
	s.RegionId = v
	return s
}

func (s *DescribeVaultReplicationRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
