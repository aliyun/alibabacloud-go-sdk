// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForNetworkZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListApplicationsForNetworkZoneRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListApplicationsForNetworkZoneRequest
	GetMaxResults() *int32
	SetNetworkZoneId(v string) *ListApplicationsForNetworkZoneRequest
	GetNetworkZoneId() *string
	SetNextToken(v string) *ListApplicationsForNetworkZoneRequest
	GetNextToken() *string
	SetPreviousToken(v string) *ListApplicationsForNetworkZoneRequest
	GetPreviousToken() *string
}

type ListApplicationsForNetworkZoneRequest struct {
	// The ID of the instance.
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
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The ID of the network domain associated with the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// network_11111
	NetworkZoneId *string `json:"NetworkZoneId,omitempty" xml:"NetworkZoneId,omitempty"`
	// The token used for the next query.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The token used to query the previous page.
	//
	// example:
	//
	// PTxxxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
}

func (s ListApplicationsForNetworkZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForNetworkZoneRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsForNetworkZoneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationsForNetworkZoneRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationsForNetworkZoneRequest) GetNetworkZoneId() *string {
	return s.NetworkZoneId
}

func (s *ListApplicationsForNetworkZoneRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationsForNetworkZoneRequest) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListApplicationsForNetworkZoneRequest) SetInstanceId(v string) *ListApplicationsForNetworkZoneRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationsForNetworkZoneRequest) SetMaxResults(v int32) *ListApplicationsForNetworkZoneRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationsForNetworkZoneRequest) SetNetworkZoneId(v string) *ListApplicationsForNetworkZoneRequest {
	s.NetworkZoneId = &v
	return s
}

func (s *ListApplicationsForNetworkZoneRequest) SetNextToken(v string) *ListApplicationsForNetworkZoneRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsForNetworkZoneRequest) SetPreviousToken(v string) *ListApplicationsForNetworkZoneRequest {
	s.PreviousToken = &v
	return s
}

func (s *ListApplicationsForNetworkZoneRequest) Validate() error {
	return dara.Validate(s)
}
