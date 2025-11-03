// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVSwitchCidrReservationUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCidrReservationUsages(v []*GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) *GetVSwitchCidrReservationUsageResponseBody
	GetCidrReservationUsages() []*GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages
	SetMaxResults(v int64) *GetVSwitchCidrReservationUsageResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *GetVSwitchCidrReservationUsageResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetVSwitchCidrReservationUsageResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetVSwitchCidrReservationUsageResponseBody
	GetTotalCount() *int64
}

type GetVSwitchCidrReservationUsageResponseBody struct {
	// A list of reserved CIDR blocks that are in use.
	CidrReservationUsages []*GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages `json:"CidrReservationUsages,omitempty" xml:"CidrReservationUsages,omitempty" type:"Repeated"`
	// The number of entries to return per page.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
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
}

func (s GetVSwitchCidrReservationUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVSwitchCidrReservationUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetVSwitchCidrReservationUsageResponseBody) GetCidrReservationUsages() []*GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages {
	return s.CidrReservationUsages
}

func (s *GetVSwitchCidrReservationUsageResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *GetVSwitchCidrReservationUsageResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetVSwitchCidrReservationUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVSwitchCidrReservationUsageResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetVSwitchCidrReservationUsageResponseBody) SetCidrReservationUsages(v []*GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) *GetVSwitchCidrReservationUsageResponseBody {
	s.CidrReservationUsages = v
	return s
}

func (s *GetVSwitchCidrReservationUsageResponseBody) SetMaxResults(v int64) *GetVSwitchCidrReservationUsageResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageResponseBody) SetNextToken(v string) *GetVSwitchCidrReservationUsageResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageResponseBody) SetRequestId(v string) *GetVSwitchCidrReservationUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageResponseBody) SetTotalCount(v int64) *GetVSwitchCidrReservationUsageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageResponseBody) Validate() error {
	if s.CidrReservationUsages != nil {
		for _, item := range s.CidrReservationUsages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages struct {
	// The CIDR block allocated to the ENI from the reserved CIDR block.
	//
	// example:
	//
	// 192.168.1.64/28
	IpPrefixCidr *string `json:"IpPrefixCidr,omitempty" xml:"IpPrefixCidr,omitempty"`
	// The ID of the reserved CIDR block.
	//
	// example:
	//
	// vcr-bp1m12saqteraw3rp****
	IpPrefixId *string `json:"IpPrefixId,omitempty" xml:"IpPrefixId,omitempty"`
	// The ID of the elastic network interface (ENI) whose CIDR block is allocated from the reserved CIDR block.
	//
	// example:
	//
	// eni-bp14v2sdd3v8htln****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource to which a CIDR block is allocated from the reserved CIDR block. Only **NetworkInterface*	- may be returned, which indicates an ENI.
	//
	// example:
	//
	// NetworkInterface
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the reserved CIDR block.
	//
	// example:
	//
	// vcr-bp1m12saqteraw3rp****
	VSwitchCidrReservationId *string `json:"VSwitchCidrReservationId,omitempty" xml:"VSwitchCidrReservationId,omitempty"`
	// The ID of the vSwitch to which the reserved CIDR block belongs.
	//
	// example:
	//
	// vsw-25navfgbue4g****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) String() string {
	return dara.Prettify(s)
}

func (s GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) GoString() string {
	return s.String()
}

func (s *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) GetIpPrefixCidr() *string {
	return s.IpPrefixCidr
}

func (s *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) GetIpPrefixId() *string {
	return s.IpPrefixId
}

func (s *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) GetVSwitchCidrReservationId() *string {
	return s.VSwitchCidrReservationId
}

func (s *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) SetIpPrefixCidr(v string) *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages {
	s.IpPrefixCidr = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) SetIpPrefixId(v string) *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages {
	s.IpPrefixId = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) SetResourceId(v string) *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages {
	s.ResourceId = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) SetResourceType(v string) *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages {
	s.ResourceType = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) SetVSwitchCidrReservationId(v string) *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages {
	s.VSwitchCidrReservationId = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) SetVSwitchId(v string) *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages {
	s.VSwitchId = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageResponseBodyCidrReservationUsages) Validate() error {
	return dara.Validate(s)
}
