// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeliveryChannelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListDeliveryChannelsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDeliveryChannelsRequest
	GetNextToken() *string
}

type ListDeliveryChannelsRequest struct {
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the MaxResults parameter, the entries are truncated. In this case, you can use the token to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListDeliveryChannelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryChannelsRequest) GoString() string {
	return s.String()
}

func (s *ListDeliveryChannelsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDeliveryChannelsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDeliveryChannelsRequest) SetMaxResults(v int32) *ListDeliveryChannelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDeliveryChannelsRequest) SetNextToken(v string) *ListDeliveryChannelsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDeliveryChannelsRequest) Validate() error {
	return dara.Validate(s)
}
