// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListNetworkZonesRequest
	GetInstanceId() *string
	SetMaxResults(v int64) *ListNetworkZonesRequest
	GetMaxResults() *int64
	SetNetworkZoneIds(v []*string) *ListNetworkZonesRequest
	GetNetworkZoneIds() []*string
	SetNextToken(v string) *ListNetworkZonesRequest
	GetNextToken() *string
	SetPreviousToken(v string) *ListNetworkZonesRequest
	GetPreviousToken() *string
}

type ListNetworkZonesRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A collection of network IDs.
	NetworkZoneIds []*string `json:"NetworkZoneIds,omitempty" xml:"NetworkZoneIds,omitempty" type:"Repeated"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The token that is used to retrieve the previous page of results.
	//
	// example:
	//
	// PTxxxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
}

func (s ListNetworkZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkZonesRequest) GoString() string {
	return s.String()
}

func (s *ListNetworkZonesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNetworkZonesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListNetworkZonesRequest) GetNetworkZoneIds() []*string {
	return s.NetworkZoneIds
}

func (s *ListNetworkZonesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNetworkZonesRequest) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListNetworkZonesRequest) SetInstanceId(v string) *ListNetworkZonesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNetworkZonesRequest) SetMaxResults(v int64) *ListNetworkZonesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNetworkZonesRequest) SetNetworkZoneIds(v []*string) *ListNetworkZonesRequest {
	s.NetworkZoneIds = v
	return s
}

func (s *ListNetworkZonesRequest) SetNextToken(v string) *ListNetworkZonesRequest {
	s.NextToken = &v
	return s
}

func (s *ListNetworkZonesRequest) SetPreviousToken(v string) *ListNetworkZonesRequest {
	s.PreviousToken = &v
	return s
}

func (s *ListNetworkZonesRequest) Validate() error {
	return dara.Validate(s)
}
