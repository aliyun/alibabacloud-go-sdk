// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQosPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeQosPoliciesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeQosPoliciesResponseBody
	GetPageSize() *int32
	SetQosPolicies(v *DescribeQosPoliciesResponseBodyQosPolicies) *DescribeQosPoliciesResponseBody
	GetQosPolicies() *DescribeQosPoliciesResponseBodyQosPolicies
	SetRequestId(v string) *DescribeQosPoliciesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeQosPoliciesResponseBody
	GetTotalCount() *int32
}

type DescribeQosPoliciesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QosPolicies *DescribeQosPoliciesResponseBodyQosPolicies `json:"QosPolicies,omitempty" xml:"QosPolicies,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 97862812-2C7E-4D25-B0D5-B26DAC7FA293
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeQosPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQosPoliciesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeQosPoliciesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeQosPoliciesResponseBody) GetQosPolicies() *DescribeQosPoliciesResponseBodyQosPolicies {
	return s.QosPolicies
}

func (s *DescribeQosPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQosPoliciesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeQosPoliciesResponseBody) SetPageNumber(v int32) *DescribeQosPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeQosPoliciesResponseBody) SetPageSize(v int32) *DescribeQosPoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeQosPoliciesResponseBody) SetQosPolicies(v *DescribeQosPoliciesResponseBodyQosPolicies) *DescribeQosPoliciesResponseBody {
	s.QosPolicies = v
	return s
}

func (s *DescribeQosPoliciesResponseBody) SetRequestId(v string) *DescribeQosPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQosPoliciesResponseBody) SetTotalCount(v int32) *DescribeQosPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeQosPoliciesResponseBody) Validate() error {
	if s.QosPolicies != nil {
		if err := s.QosPolicies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeQosPoliciesResponseBodyQosPolicies struct {
	QosPolicy []*DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy `json:"QosPolicy,omitempty" xml:"QosPolicy,omitempty" type:"Repeated"`
}

func (s DescribeQosPoliciesResponseBodyQosPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosPoliciesResponseBodyQosPolicies) GoString() string {
	return s.String()
}

func (s *DescribeQosPoliciesResponseBodyQosPolicies) GetQosPolicy() []*DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy {
	return s.QosPolicy
}

func (s *DescribeQosPoliciesResponseBodyQosPolicies) SetQosPolicy(v []*DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) *DescribeQosPoliciesResponseBodyQosPolicies {
	s.QosPolicy = v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPolicies) Validate() error {
	if s.QosPolicy != nil {
		for _, item := range s.QosPolicy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy struct {
	Description     *string                                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	DestCidr        *string                                                             `json:"DestCidr,omitempty" xml:"DestCidr,omitempty"`
	DestPortRange   *string                                                             `json:"DestPortRange,omitempty" xml:"DestPortRange,omitempty"`
	DpiGroupIds     *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiGroupIds     `json:"DpiGroupIds,omitempty" xml:"DpiGroupIds,omitempty" type:"Struct"`
	DpiSignatureIds *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiSignatureIds `json:"DpiSignatureIds,omitempty" xml:"DpiSignatureIds,omitempty" type:"Struct"`
	EndTime         *string                                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IpProtocol      *string                                                             `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Name            *string                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Priority        *int32                                                              `json:"Priority,omitempty" xml:"Priority,omitempty"`
	QosId           *string                                                             `json:"QosId,omitempty" xml:"QosId,omitempty"`
	QosPolicyId     *string                                                             `json:"QosPolicyId,omitempty" xml:"QosPolicyId,omitempty"`
	SourceCidr      *string                                                             `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	SourcePortRange *string                                                             `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	StartTime       *string                                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) GoString() string {
	return s.String()
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) GetDescription() *string {
	return s.Description
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) GetDestCidr() *string {
	return s.DestCidr
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) GetDestPortRange() *string {
	return s.DestPortRange
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) GetDpiGroupIds() *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiGroupIds {
	return s.DpiGroupIds
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) GetDpiSignatureIds() *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiSignatureIds {
	return s.DpiSignatureIds
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) GetName() *string {
	return s.Name
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) GetQosId() *string {
	return s.QosId
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) GetQosPolicyId() *string {
	return s.QosPolicyId
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) SetDescription(v string) *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy {
	s.Description = &v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) SetDestCidr(v string) *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy {
	s.DestCidr = &v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) SetDestPortRange(v string) *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy {
	s.DestPortRange = &v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) SetDpiGroupIds(v *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiGroupIds) *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy {
	s.DpiGroupIds = v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) SetDpiSignatureIds(v *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiSignatureIds) *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy {
	s.DpiSignatureIds = v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) SetEndTime(v string) *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy {
	s.EndTime = &v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) SetIpProtocol(v string) *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy {
	s.IpProtocol = &v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) SetName(v string) *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy {
	s.Name = &v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) SetPriority(v int32) *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy {
	s.Priority = &v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) SetQosId(v string) *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy {
	s.QosId = &v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) SetQosPolicyId(v string) *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy {
	s.QosPolicyId = &v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) SetSourceCidr(v string) *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy {
	s.SourceCidr = &v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) SetSourcePortRange(v string) *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy {
	s.SourcePortRange = &v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) SetStartTime(v string) *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy {
	s.StartTime = &v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicy) Validate() error {
	if s.DpiGroupIds != nil {
		if err := s.DpiGroupIds.Validate(); err != nil {
			return err
		}
	}
	if s.DpiSignatureIds != nil {
		if err := s.DpiSignatureIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiGroupIds struct {
	DpiGroupId []*string `json:"DpiGroupId,omitempty" xml:"DpiGroupId,omitempty" type:"Repeated"`
}

func (s DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiGroupIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiGroupIds) GetDpiGroupId() []*string {
	return s.DpiGroupId
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiGroupIds) SetDpiGroupId(v []*string) *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiGroupIds {
	s.DpiGroupId = v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiGroupIds) Validate() error {
	return dara.Validate(s)
}

type DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiSignatureIds struct {
	DpiSignatureId []*string `json:"DpiSignatureId,omitempty" xml:"DpiSignatureId,omitempty" type:"Repeated"`
}

func (s DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiSignatureIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiSignatureIds) GoString() string {
	return s.String()
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiSignatureIds) GetDpiSignatureId() []*string {
	return s.DpiSignatureId
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiSignatureIds) SetDpiSignatureId(v []*string) *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiSignatureIds {
	s.DpiSignatureId = v
	return s
}

func (s *DescribeQosPoliciesResponseBodyQosPoliciesQosPolicyDpiSignatureIds) Validate() error {
	return dara.Validate(s)
}
