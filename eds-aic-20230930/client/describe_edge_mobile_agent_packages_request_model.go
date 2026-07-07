// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeMobileAgentPackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceClass(v string) *DescribeEdgeMobileAgentPackagesRequest
	GetDeviceClass() *string
	SetLicenseKeys(v string) *DescribeEdgeMobileAgentPackagesRequest
	GetLicenseKeys() *string
	SetMaxResults(v int32) *DescribeEdgeMobileAgentPackagesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeEdgeMobileAgentPackagesRequest
	GetNextToken() *string
	SetPackageIds(v string) *DescribeEdgeMobileAgentPackagesRequest
	GetPackageIds() *string
	SetStatus(v string) *DescribeEdgeMobileAgentPackagesRequest
	GetStatus() *string
}

type DescribeEdgeMobileAgentPackagesRequest struct {
	// The device type filter. Valid values: BOX, PHONE, PAD, and OTHER.
	//
	// example:
	//
	// BOX
	DeviceClass *string `json:"DeviceClass,omitempty" xml:"DeviceClass,omitempty"`
	// The list of license keys. Separate multiple keys with commas.
	//
	// example:
	//
	// lic-ez197xvdf0j5eo0*****
	LicenseKeys *string `json:"LicenseKeys,omitempty" xml:"LicenseKeys,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Leave this parameter empty for the first query. For subsequent queries, use the value returned in the previous response.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of package IDs. Separate multiple IDs with commas.
	//
	// example:
	//
	// cmag-0c1g77wjljl9hc****
	PackageIds *string `json:"PackageIds,omitempty" xml:"PackageIds,omitempty"`
	// The package status filter.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEdgeMobileAgentPackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeMobileAgentPackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEdgeMobileAgentPackagesRequest) GetDeviceClass() *string {
	return s.DeviceClass
}

func (s *DescribeEdgeMobileAgentPackagesRequest) GetLicenseKeys() *string {
	return s.LicenseKeys
}

func (s *DescribeEdgeMobileAgentPackagesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeEdgeMobileAgentPackagesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeEdgeMobileAgentPackagesRequest) GetPackageIds() *string {
	return s.PackageIds
}

func (s *DescribeEdgeMobileAgentPackagesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeEdgeMobileAgentPackagesRequest) SetDeviceClass(v string) *DescribeEdgeMobileAgentPackagesRequest {
	s.DeviceClass = &v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesRequest) SetLicenseKeys(v string) *DescribeEdgeMobileAgentPackagesRequest {
	s.LicenseKeys = &v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesRequest) SetMaxResults(v int32) *DescribeEdgeMobileAgentPackagesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesRequest) SetNextToken(v string) *DescribeEdgeMobileAgentPackagesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesRequest) SetPackageIds(v string) *DescribeEdgeMobileAgentPackagesRequest {
	s.PackageIds = &v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesRequest) SetStatus(v string) *DescribeEdgeMobileAgentPackagesRequest {
	s.Status = &v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesRequest) Validate() error {
	return dara.Validate(s)
}
