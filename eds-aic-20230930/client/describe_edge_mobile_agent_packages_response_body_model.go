// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeMobileAgentPackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeEdgeMobileAgentPackagesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeEdgeMobileAgentPackagesResponseBody
	GetNextToken() *string
	SetPackages(v []*DescribeEdgeMobileAgentPackagesResponseBodyPackages) *DescribeEdgeMobileAgentPackagesResponseBody
	GetPackages() []*DescribeEdgeMobileAgentPackagesResponseBodyPackages
	SetRequestId(v string) *DescribeEdgeMobileAgentPackagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeEdgeMobileAgentPackagesResponseBody
	GetTotalCount() *int32
}

type DescribeEdgeMobileAgentPackagesResponseBody struct {
	// The actual number of entries returned on the current page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. An empty value indicates that no more data exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of packages.
	Packages []*DescribeEdgeMobileAgentPackagesResponseBodyPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5C5CEF0A-D6E1-58D3-8750-67DB4F82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEdgeMobileAgentPackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeMobileAgentPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEdgeMobileAgentPackagesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeEdgeMobileAgentPackagesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeEdgeMobileAgentPackagesResponseBody) GetPackages() []*DescribeEdgeMobileAgentPackagesResponseBodyPackages {
	return s.Packages
}

func (s *DescribeEdgeMobileAgentPackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEdgeMobileAgentPackagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEdgeMobileAgentPackagesResponseBody) SetMaxResults(v int32) *DescribeEdgeMobileAgentPackagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesResponseBody) SetNextToken(v string) *DescribeEdgeMobileAgentPackagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesResponseBody) SetPackages(v []*DescribeEdgeMobileAgentPackagesResponseBodyPackages) *DescribeEdgeMobileAgentPackagesResponseBody {
	s.Packages = v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesResponseBody) SetRequestId(v string) *DescribeEdgeMobileAgentPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesResponseBody) SetTotalCount(v int32) *DescribeEdgeMobileAgentPackagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesResponseBody) Validate() error {
	if s.Packages != nil {
		for _, item := range s.Packages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEdgeMobileAgentPackagesResponseBodyPackages struct {
	// The device type.
	//
	// example:
	//
	// BOX
	DeviceClass *string `json:"DeviceClass,omitempty" xml:"DeviceClass,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2026-01-05 10:04:07
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The list of license keys.
	LicenseKeys []*string `json:"LicenseKeys,omitempty" xml:"LicenseKeys,omitempty" type:"Repeated"`
	// The package ID.
	//
	// example:
	//
	// cmag-0c1g77wjljl9hc****
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// The package specification. Currently, only hardware is supported.
	//
	// example:
	//
	// hardware
	PackageSpec *string `json:"PackageSpec,omitempty" xml:"PackageSpec,omitempty"`
	// The package status.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEdgeMobileAgentPackagesResponseBodyPackages) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeMobileAgentPackagesResponseBodyPackages) GoString() string {
	return s.String()
}

func (s *DescribeEdgeMobileAgentPackagesResponseBodyPackages) GetDeviceClass() *string {
	return s.DeviceClass
}

func (s *DescribeEdgeMobileAgentPackagesResponseBodyPackages) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *DescribeEdgeMobileAgentPackagesResponseBodyPackages) GetLicenseKeys() []*string {
	return s.LicenseKeys
}

func (s *DescribeEdgeMobileAgentPackagesResponseBodyPackages) GetPackageId() *string {
	return s.PackageId
}

func (s *DescribeEdgeMobileAgentPackagesResponseBodyPackages) GetPackageSpec() *string {
	return s.PackageSpec
}

func (s *DescribeEdgeMobileAgentPackagesResponseBodyPackages) GetStatus() *string {
	return s.Status
}

func (s *DescribeEdgeMobileAgentPackagesResponseBodyPackages) SetDeviceClass(v string) *DescribeEdgeMobileAgentPackagesResponseBodyPackages {
	s.DeviceClass = &v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesResponseBodyPackages) SetExpireDate(v string) *DescribeEdgeMobileAgentPackagesResponseBodyPackages {
	s.ExpireDate = &v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesResponseBodyPackages) SetLicenseKeys(v []*string) *DescribeEdgeMobileAgentPackagesResponseBodyPackages {
	s.LicenseKeys = v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesResponseBodyPackages) SetPackageId(v string) *DescribeEdgeMobileAgentPackagesResponseBodyPackages {
	s.PackageId = &v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesResponseBodyPackages) SetPackageSpec(v string) *DescribeEdgeMobileAgentPackagesResponseBodyPackages {
	s.PackageSpec = &v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesResponseBodyPackages) SetStatus(v string) *DescribeEdgeMobileAgentPackagesResponseBodyPackages {
	s.Status = &v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesResponseBodyPackages) Validate() error {
	return dara.Validate(s)
}
