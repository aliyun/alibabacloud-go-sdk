// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVSwitchCidrReservationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListVSwitchCidrReservationsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListVSwitchCidrReservationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVSwitchCidrReservationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListVSwitchCidrReservationsResponseBody
	GetTotalCount() *int64
	SetVSwitchCidrReservations(v []*ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) *ListVSwitchCidrReservationsResponseBody
	GetVSwitchCidrReservations() []*ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations
}

type ListVSwitchCidrReservationsResponseBody struct {
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Valid values:
	//
	// - If **NextToken*	- is empty, no subsequent query exists.
	//
	// - If **NextToken*	- is returned, the value indicates the token for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of reserved CIDR blocks for a vSwitch.
	VSwitchCidrReservations []*ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations `json:"VSwitchCidrReservations,omitempty" xml:"VSwitchCidrReservations,omitempty" type:"Repeated"`
}

func (s ListVSwitchCidrReservationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVSwitchCidrReservationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVSwitchCidrReservationsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListVSwitchCidrReservationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVSwitchCidrReservationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVSwitchCidrReservationsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListVSwitchCidrReservationsResponseBody) GetVSwitchCidrReservations() []*ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations {
	return s.VSwitchCidrReservations
}

func (s *ListVSwitchCidrReservationsResponseBody) SetMaxResults(v int64) *ListVSwitchCidrReservationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBody) SetNextToken(v string) *ListVSwitchCidrReservationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBody) SetRequestId(v string) *ListVSwitchCidrReservationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBody) SetTotalCount(v int64) *ListVSwitchCidrReservationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBody) SetVSwitchCidrReservations(v []*ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) *ListVSwitchCidrReservationsResponseBody {
	s.VSwitchCidrReservations = v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBody) Validate() error {
	if s.VSwitchCidrReservations != nil {
		for _, item := range s.VSwitchCidrReservations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations struct {
	// The number of used prefix CIDR blocks in the reserved CIDR block for a vSwitch.
	//
	// example:
	//
	// 6
	AssignedCidrCount *int32 `json:"AssignedCidrCount,omitempty" xml:"AssignedCidrCount,omitempty"`
	// The number of active prefix CIDR blocks in the reserved CIDR block for a vSwitch.
	//
	// example:
	//
	// 10
	AvailableCidrCount *int32 `json:"AvailableCidrCount,omitempty" xml:"AvailableCidrCount,omitempty"`
	// The time when the reserved CIDR block was created.
	//
	// example:
	//
	// 2023-03-14T10:02:37Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The IP version of the reserved CIDR block for a vSwitch. Valid values:
	//
	// - **IPv4*	- (default): IPv4.
	//
	// - **IPv6**: IPv6.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The status of the reserved CIDR block. Valid values:
	//
	// - **Assigning**: being allocated.
	//
	// - **Assigned**: allocated.
	//
	// - **Releasing**: being released.
	//
	// - **Released**: released.
	//
	// example:
	//
	// Assigned
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservationsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the reserved CIDR block for a vSwitch. Valid values: **prefix*	- (default), which indicates that addresses are allocated by CIDR block.
	//
	// example:
	//
	// prefix
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The reserved CIDR block for a vSwitch.
	//
	// example:
	//
	// 192.168.1.64/28
	VSwitchCidrReservationCidr *string `json:"VSwitchCidrReservationCidr,omitempty" xml:"VSwitchCidrReservationCidr,omitempty"`
	// The description of the reserved CIDR block for a vSwitch.
	//
	// example:
	//
	// ReservationDescription
	VSwitchCidrReservationDescription *string `json:"VSwitchCidrReservationDescription,omitempty" xml:"VSwitchCidrReservationDescription,omitempty"`
	// The instance ID of the reserved CIDR block for a vSwitch.
	//
	// example:
	//
	// vcr-bp1m12saqteraw3rp****
	VSwitchCidrReservationId *string `json:"VSwitchCidrReservationId,omitempty" xml:"VSwitchCidrReservationId,omitempty"`
	// The name of the reserved CIDR block for a vSwitch.
	//
	// example:
	//
	// ReservationName
	VSwitchCidrReservationName *string `json:"VSwitchCidrReservationName,omitempty" xml:"VSwitchCidrReservationName,omitempty"`
	// The ID of the vSwitch to which the reserved CIDR block for a vSwitch belongs.
	//
	// example:
	//
	// vsw-25navfgbue4g****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC to which the reserved CIDR block for a vSwitch belongs.
	//
	// example:
	//
	// vpc-bp1wdz2pdhgurz1od****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) String() string {
	return dara.Prettify(s)
}

func (s ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) GoString() string {
	return s.String()
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) GetAssignedCidrCount() *int32 {
	return s.AssignedCidrCount
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) GetAvailableCidrCount() *int32 {
	return s.AvailableCidrCount
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) GetStatus() *string {
	return s.Status
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) GetTags() []*ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservationsTags {
	return s.Tags
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) GetType() *string {
	return s.Type
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) GetVSwitchCidrReservationCidr() *string {
	return s.VSwitchCidrReservationCidr
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) GetVSwitchCidrReservationDescription() *string {
	return s.VSwitchCidrReservationDescription
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) GetVSwitchCidrReservationId() *string {
	return s.VSwitchCidrReservationId
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) GetVSwitchCidrReservationName() *string {
	return s.VSwitchCidrReservationName
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) SetAssignedCidrCount(v int32) *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations {
	s.AssignedCidrCount = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) SetAvailableCidrCount(v int32) *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations {
	s.AvailableCidrCount = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) SetCreationTime(v string) *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations {
	s.CreationTime = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) SetIpVersion(v string) *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations {
	s.IpVersion = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) SetStatus(v string) *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations {
	s.Status = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) SetTags(v []*ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservationsTags) *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations {
	s.Tags = v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) SetType(v string) *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations {
	s.Type = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) SetVSwitchCidrReservationCidr(v string) *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations {
	s.VSwitchCidrReservationCidr = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) SetVSwitchCidrReservationDescription(v string) *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations {
	s.VSwitchCidrReservationDescription = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) SetVSwitchCidrReservationId(v string) *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations {
	s.VSwitchCidrReservationId = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) SetVSwitchCidrReservationName(v string) *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations {
	s.VSwitchCidrReservationName = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) SetVSwitchId(v string) *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations {
	s.VSwitchId = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) SetVpcId(v string) *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations {
	s.VpcId = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservations) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservationsTags struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservationsTags) String() string {
	return dara.Prettify(s)
}

func (s ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservationsTags) GoString() string {
	return s.String()
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservationsTags) GetKey() *string {
	return s.Key
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservationsTags) GetValue() *string {
	return s.Value
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservationsTags) SetKey(v string) *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservationsTags {
	s.Key = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservationsTags) SetValue(v string) *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservationsTags {
	s.Value = &v
	return s
}

func (s *ListVSwitchCidrReservationsResponseBodyVSwitchCidrReservationsTags) Validate() error {
	return dara.Validate(s)
}
