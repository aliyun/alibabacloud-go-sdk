// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteServicesInCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeRouteServicesInCenResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRouteServicesInCenResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRouteServicesInCenResponseBody
	GetRequestId() *string
	SetRouteServiceEntries(v *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) *DescribeRouteServicesInCenResponseBody
	GetRouteServiceEntries() *DescribeRouteServicesInCenResponseBodyRouteServiceEntries
	SetTotalCount(v int32) *DescribeRouteServicesInCenResponseBody
	GetTotalCount() *int32
}

type DescribeRouteServicesInCenResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 196C99CA-6997-5951-9721-AE89720DF856
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the cloud services.
	RouteServiceEntries *DescribeRouteServicesInCenResponseBodyRouteServiceEntries `json:"RouteServiceEntries,omitempty" xml:"RouteServiceEntries,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRouteServicesInCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteServicesInCenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouteServicesInCenResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRouteServicesInCenResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRouteServicesInCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRouteServicesInCenResponseBody) GetRouteServiceEntries() *DescribeRouteServicesInCenResponseBodyRouteServiceEntries {
	return s.RouteServiceEntries
}

func (s *DescribeRouteServicesInCenResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRouteServicesInCenResponseBody) SetPageNumber(v int32) *DescribeRouteServicesInCenResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBody) SetPageSize(v int32) *DescribeRouteServicesInCenResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBody) SetRequestId(v string) *DescribeRouteServicesInCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBody) SetRouteServiceEntries(v *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) *DescribeRouteServicesInCenResponseBody {
	s.RouteServiceEntries = v
	return s
}

func (s *DescribeRouteServicesInCenResponseBody) SetTotalCount(v int32) *DescribeRouteServicesInCenResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBody) Validate() error {
	if s.RouteServiceEntries != nil {
		if err := s.RouteServiceEntries.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouteServicesInCenResponseBodyRouteServiceEntries struct {
	RouteServiceEntry []*DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry `json:"RouteServiceEntry,omitempty" xml:"RouteServiceEntry,omitempty" type:"Repeated"`
}

func (s DescribeRouteServicesInCenResponseBodyRouteServiceEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteServicesInCenResponseBodyRouteServiceEntries) GoString() string {
	return s.String()
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) GetRouteServiceEntry() []*DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	return s.RouteServiceEntry
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) SetRouteServiceEntry(v []*DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) *DescribeRouteServicesInCenResponseBodyRouteServiceEntries {
	s.RouteServiceEntry = v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntries) Validate() error {
	if s.RouteServiceEntry != nil {
		for _, item := range s.RouteServiceEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry struct {
	// The ID of the region where the cloud service is accessed.
	//
	// example:
	//
	// cn-hangzhou
	AccessRegionId *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-pfa6ugf3xl0qsd****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The service addresses of the cloud service.
	Cidrs *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs `json:"Cidrs,omitempty" xml:"Cidrs,omitempty" type:"Struct"`
	// The description of the cloud service.
	//
	// example:
	//
	// descname
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The service address of the cloud service.
	//
	// example:
	//
	// 100.118.28.0/24
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The region ID of the cloud service.
	//
	// example:
	//
	// cn-hangzhou
	HostRegionId *string `json:"HostRegionId,omitempty" xml:"HostRegionId,omitempty"`
	// The ID of the VPC associated with the cloud service.
	//
	// example:
	//
	// vpc-bp1h8vbrbcgohcju5****
	HostVpcId *string `json:"HostVpcId,omitempty" xml:"HostVpcId,omitempty"`
	// The status of the cloud service. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Active**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) GoString() string {
	return s.String()
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) GetAccessRegionId() *string {
	return s.AccessRegionId
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) GetCenId() *string {
	return s.CenId
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) GetCidrs() *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs {
	return s.Cidrs
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) GetDescription() *string {
	return s.Description
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) GetHost() *string {
	return s.Host
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) GetHostRegionId() *string {
	return s.HostRegionId
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) GetHostVpcId() *string {
	return s.HostVpcId
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) GetStatus() *string {
	return s.Status
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetAccessRegionId(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.AccessRegionId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetCenId(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.CenId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetCidrs(v *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.Cidrs = v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetDescription(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.Description = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetHost(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.Host = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetHostRegionId(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.HostRegionId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetHostVpcId(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.HostVpcId = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) SetStatus(v string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry {
	s.Status = &v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntry) Validate() error {
	if s.Cidrs != nil {
		if err := s.Cidrs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs struct {
	Cidr []*string `json:"Cidr,omitempty" xml:"Cidr,omitempty" type:"Repeated"`
}

func (s DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs) GoString() string {
	return s.String()
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs) GetCidr() []*string {
	return s.Cidr
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs) SetCidr(v []*string) *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs {
	s.Cidr = v
	return s
}

func (s *DescribeRouteServicesInCenResponseBodyRouteServiceEntriesRouteServiceEntryCidrs) Validate() error {
	return dara.Validate(s)
}
