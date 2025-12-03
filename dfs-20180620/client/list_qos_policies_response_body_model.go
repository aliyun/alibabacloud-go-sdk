// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQosPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListQosPoliciesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListQosPoliciesResponseBody
	GetNextToken() *string
	SetQosPolicies(v []*ListQosPoliciesResponseBodyQosPolicies) *ListQosPoliciesResponseBody
	GetQosPolicies() []*ListQosPoliciesResponseBodyQosPolicies
	SetRequestId(v string) *ListQosPoliciesResponseBody
	GetRequestId() *string
}

type ListQosPoliciesResponseBody struct {
	MaxResults  *int32                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	QosPolicies []*ListQosPoliciesResponseBodyQosPolicies `json:"QosPolicies,omitempty" xml:"QosPolicies,omitempty" type:"Repeated"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListQosPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQosPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQosPoliciesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQosPoliciesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQosPoliciesResponseBody) GetQosPolicies() []*ListQosPoliciesResponseBodyQosPolicies {
	return s.QosPolicies
}

func (s *ListQosPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQosPoliciesResponseBody) SetMaxResults(v int32) *ListQosPoliciesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListQosPoliciesResponseBody) SetNextToken(v string) *ListQosPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQosPoliciesResponseBody) SetQosPolicies(v []*ListQosPoliciesResponseBodyQosPolicies) *ListQosPoliciesResponseBody {
	s.QosPolicies = v
	return s
}

func (s *ListQosPoliciesResponseBody) SetRequestId(v string) *ListQosPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQosPoliciesResponseBody) Validate() error {
	if s.QosPolicies != nil {
		for _, item := range s.QosPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQosPoliciesResponseBodyQosPolicies struct {
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FederationId   *string `json:"FederationId,omitempty" xml:"FederationId,omitempty"`
	FileSystemId   *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FlowIds        *string `json:"FlowIds,omitempty" xml:"FlowIds,omitempty"`
	MaxIOBandWidth *int64  `json:"MaxIOBandWidth,omitempty" xml:"MaxIOBandWidth,omitempty"`
	MaxIOps        *int64  `json:"MaxIOps,omitempty" xml:"MaxIOps,omitempty"`
	MaxMetaQps     *int64  `json:"MaxMetaQps,omitempty" xml:"MaxMetaQps,omitempty"`
	QosPolicyId    *string `json:"QosPolicyId,omitempty" xml:"QosPolicyId,omitempty"`
	ReqTags        *string `json:"ReqTags,omitempty" xml:"ReqTags,omitempty"`
	ZoneIds        *string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty"`
}

func (s ListQosPoliciesResponseBodyQosPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListQosPoliciesResponseBodyQosPolicies) GoString() string {
	return s.String()
}

func (s *ListQosPoliciesResponseBodyQosPolicies) GetDescription() *string {
	return s.Description
}

func (s *ListQosPoliciesResponseBodyQosPolicies) GetFederationId() *string {
	return s.FederationId
}

func (s *ListQosPoliciesResponseBodyQosPolicies) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListQosPoliciesResponseBodyQosPolicies) GetFlowIds() *string {
	return s.FlowIds
}

func (s *ListQosPoliciesResponseBodyQosPolicies) GetMaxIOBandWidth() *int64 {
	return s.MaxIOBandWidth
}

func (s *ListQosPoliciesResponseBodyQosPolicies) GetMaxIOps() *int64 {
	return s.MaxIOps
}

func (s *ListQosPoliciesResponseBodyQosPolicies) GetMaxMetaQps() *int64 {
	return s.MaxMetaQps
}

func (s *ListQosPoliciesResponseBodyQosPolicies) GetQosPolicyId() *string {
	return s.QosPolicyId
}

func (s *ListQosPoliciesResponseBodyQosPolicies) GetReqTags() *string {
	return s.ReqTags
}

func (s *ListQosPoliciesResponseBodyQosPolicies) GetZoneIds() *string {
	return s.ZoneIds
}

func (s *ListQosPoliciesResponseBodyQosPolicies) SetDescription(v string) *ListQosPoliciesResponseBodyQosPolicies {
	s.Description = &v
	return s
}

func (s *ListQosPoliciesResponseBodyQosPolicies) SetFederationId(v string) *ListQosPoliciesResponseBodyQosPolicies {
	s.FederationId = &v
	return s
}

func (s *ListQosPoliciesResponseBodyQosPolicies) SetFileSystemId(v string) *ListQosPoliciesResponseBodyQosPolicies {
	s.FileSystemId = &v
	return s
}

func (s *ListQosPoliciesResponseBodyQosPolicies) SetFlowIds(v string) *ListQosPoliciesResponseBodyQosPolicies {
	s.FlowIds = &v
	return s
}

func (s *ListQosPoliciesResponseBodyQosPolicies) SetMaxIOBandWidth(v int64) *ListQosPoliciesResponseBodyQosPolicies {
	s.MaxIOBandWidth = &v
	return s
}

func (s *ListQosPoliciesResponseBodyQosPolicies) SetMaxIOps(v int64) *ListQosPoliciesResponseBodyQosPolicies {
	s.MaxIOps = &v
	return s
}

func (s *ListQosPoliciesResponseBodyQosPolicies) SetMaxMetaQps(v int64) *ListQosPoliciesResponseBodyQosPolicies {
	s.MaxMetaQps = &v
	return s
}

func (s *ListQosPoliciesResponseBodyQosPolicies) SetQosPolicyId(v string) *ListQosPoliciesResponseBodyQosPolicies {
	s.QosPolicyId = &v
	return s
}

func (s *ListQosPoliciesResponseBodyQosPolicies) SetReqTags(v string) *ListQosPoliciesResponseBodyQosPolicies {
	s.ReqTags = &v
	return s
}

func (s *ListQosPoliciesResponseBodyQosPolicies) SetZoneIds(v string) *ListQosPoliciesResponseBodyQosPolicies {
	s.ZoneIds = &v
	return s
}

func (s *ListQosPoliciesResponseBodyQosPolicies) Validate() error {
	return dara.Validate(s)
}
