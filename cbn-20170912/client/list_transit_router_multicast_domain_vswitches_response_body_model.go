// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterMulticastDomainVSwitchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterMulticastDomainVSwitchesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterMulticastDomainVSwitchesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTransitRouterMulticastDomainVSwitchesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransitRouterMulticastDomainVSwitchesResponseBody
	GetTotalCount() *int32
	SetVSwitchIds(v []*string) *ListTransitRouterMulticastDomainVSwitchesResponseBody
	GetVSwitchIds() []*string
}

type ListTransitRouterMulticastDomainVSwitchesResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// B0E7E43C-979A-5130-AA0D-B3ADA69E0827
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s ListTransitRouterMulticastDomainVSwitchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastDomainVSwitchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponseBody) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponseBody) SetMaxResults(v int32) *ListTransitRouterMulticastDomainVSwitchesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponseBody) SetNextToken(v string) *ListTransitRouterMulticastDomainVSwitchesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponseBody) SetRequestId(v string) *ListTransitRouterMulticastDomainVSwitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponseBody) SetTotalCount(v int32) *ListTransitRouterMulticastDomainVSwitchesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponseBody) SetVSwitchIds(v []*string) *ListTransitRouterMulticastDomainVSwitchesResponseBody {
	s.VSwitchIds = v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesResponseBody) Validate() error {
	return dara.Validate(s)
}
