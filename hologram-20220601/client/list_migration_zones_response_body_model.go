// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMigrationZonesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListMigrationZonesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMigrationZonesResponseBody
	GetRequestId() *string
	SetZoneList(v []*ListMigrationZonesResponseBodyZoneList) *ListMigrationZonesResponseBody
	GetZoneList() []*ListMigrationZonesResponseBodyZoneList
}

type ListMigrationZonesResponseBody struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 11
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 819A7F0F-2951-540F-BD94-6A41ECF0281F
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ZoneList  []*ListMigrationZonesResponseBodyZoneList `json:"zoneList,omitempty" xml:"zoneList,omitempty" type:"Repeated"`
}

func (s ListMigrationZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMigrationZonesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMigrationZonesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMigrationZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMigrationZonesResponseBody) GetZoneList() []*ListMigrationZonesResponseBodyZoneList {
	return s.ZoneList
}

func (s *ListMigrationZonesResponseBody) SetMaxResults(v int32) *ListMigrationZonesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMigrationZonesResponseBody) SetNextToken(v string) *ListMigrationZonesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMigrationZonesResponseBody) SetRequestId(v string) *ListMigrationZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMigrationZonesResponseBody) SetZoneList(v []*ListMigrationZonesResponseBodyZoneList) *ListMigrationZonesResponseBody {
	s.ZoneList = v
	return s
}

func (s *ListMigrationZonesResponseBody) Validate() error {
	if s.ZoneList != nil {
		for _, item := range s.ZoneList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMigrationZonesResponseBodyZoneList struct {
	// example:
	//
	// true
	Available *bool `json:"available,omitempty" xml:"available,omitempty"`
	// example:
	//
	// cn-beijing-i
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ListMigrationZonesResponseBodyZoneList) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationZonesResponseBodyZoneList) GoString() string {
	return s.String()
}

func (s *ListMigrationZonesResponseBodyZoneList) GetAvailable() *bool {
	return s.Available
}

func (s *ListMigrationZonesResponseBodyZoneList) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListMigrationZonesResponseBodyZoneList) SetAvailable(v bool) *ListMigrationZonesResponseBodyZoneList {
	s.Available = &v
	return s
}

func (s *ListMigrationZonesResponseBodyZoneList) SetZoneId(v string) *ListMigrationZonesResponseBodyZoneList {
	s.ZoneId = &v
	return s
}

func (s *ListMigrationZonesResponseBodyZoneList) Validate() error {
	return dara.Validate(s)
}
