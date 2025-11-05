// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReplicaEdgeSupportedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListReplicaEdgeSupportedResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListReplicaEdgeSupportedResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListReplicaEdgeSupportedResponseBody
	GetRequestId() *string
	SetSupportedRegions(v []*ListReplicaEdgeSupportedResponseBodySupportedRegions) *ListReplicaEdgeSupportedResponseBody
	GetSupportedRegions() []*ListReplicaEdgeSupportedResponseBodySupportedRegions
}

type ListReplicaEdgeSupportedResponseBody struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// e71d8a535bd9c****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 17EE62D8-064E-5404-8B0D-72122478****
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupportedRegions []*ListReplicaEdgeSupportedResponseBodySupportedRegions `json:"SupportedRegions,omitempty" xml:"SupportedRegions,omitempty" type:"Repeated"`
}

func (s ListReplicaEdgeSupportedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListReplicaEdgeSupportedResponseBody) GoString() string {
	return s.String()
}

func (s *ListReplicaEdgeSupportedResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListReplicaEdgeSupportedResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListReplicaEdgeSupportedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListReplicaEdgeSupportedResponseBody) GetSupportedRegions() []*ListReplicaEdgeSupportedResponseBodySupportedRegions {
	return s.SupportedRegions
}

func (s *ListReplicaEdgeSupportedResponseBody) SetMaxResults(v int32) *ListReplicaEdgeSupportedResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListReplicaEdgeSupportedResponseBody) SetNextToken(v string) *ListReplicaEdgeSupportedResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListReplicaEdgeSupportedResponseBody) SetRequestId(v string) *ListReplicaEdgeSupportedResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListReplicaEdgeSupportedResponseBody) SetSupportedRegions(v []*ListReplicaEdgeSupportedResponseBodySupportedRegions) *ListReplicaEdgeSupportedResponseBody {
	s.SupportedRegions = v
	return s
}

func (s *ListReplicaEdgeSupportedResponseBody) Validate() error {
	if s.SupportedRegions != nil {
		for _, item := range s.SupportedRegions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListReplicaEdgeSupportedResponseBodySupportedRegions struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string                                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Zones    []*ListReplicaEdgeSupportedResponseBodySupportedRegionsZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListReplicaEdgeSupportedResponseBodySupportedRegions) String() string {
	return dara.Prettify(s)
}

func (s ListReplicaEdgeSupportedResponseBodySupportedRegions) GoString() string {
	return s.String()
}

func (s *ListReplicaEdgeSupportedResponseBodySupportedRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *ListReplicaEdgeSupportedResponseBodySupportedRegions) GetZones() []*ListReplicaEdgeSupportedResponseBodySupportedRegionsZones {
	return s.Zones
}

func (s *ListReplicaEdgeSupportedResponseBodySupportedRegions) SetRegionId(v string) *ListReplicaEdgeSupportedResponseBodySupportedRegions {
	s.RegionId = &v
	return s
}

func (s *ListReplicaEdgeSupportedResponseBodySupportedRegions) SetZones(v []*ListReplicaEdgeSupportedResponseBodySupportedRegionsZones) *ListReplicaEdgeSupportedResponseBodySupportedRegions {
	s.Zones = v
	return s
}

func (s *ListReplicaEdgeSupportedResponseBodySupportedRegions) Validate() error {
	if s.Zones != nil {
		for _, item := range s.Zones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListReplicaEdgeSupportedResponseBodySupportedRegionsZones struct {
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListReplicaEdgeSupportedResponseBodySupportedRegionsZones) String() string {
	return dara.Prettify(s)
}

func (s ListReplicaEdgeSupportedResponseBodySupportedRegionsZones) GoString() string {
	return s.String()
}

func (s *ListReplicaEdgeSupportedResponseBodySupportedRegionsZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListReplicaEdgeSupportedResponseBodySupportedRegionsZones) SetZoneId(v string) *ListReplicaEdgeSupportedResponseBodySupportedRegionsZones {
	s.ZoneId = &v
	return s
}

func (s *ListReplicaEdgeSupportedResponseBodySupportedRegionsZones) Validate() error {
	return dara.Validate(s)
}
