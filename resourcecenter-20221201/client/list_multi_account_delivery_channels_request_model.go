// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountDeliveryChannelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMultiAccountDeliveryChannelsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMultiAccountDeliveryChannelsRequest
	GetNextToken() *string
}

type ListMultiAccountDeliveryChannelsRequest struct {
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
	// AAAAARfZmVDe9NvRXloR5+8CK9nNJufMdRA7W1miLC1P****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListMultiAccountDeliveryChannelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountDeliveryChannelsRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountDeliveryChannelsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMultiAccountDeliveryChannelsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiAccountDeliveryChannelsRequest) SetMaxResults(v int32) *ListMultiAccountDeliveryChannelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsRequest) SetNextToken(v string) *ListMultiAccountDeliveryChannelsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountDeliveryChannelsRequest) Validate() error {
	return dara.Validate(s)
}
