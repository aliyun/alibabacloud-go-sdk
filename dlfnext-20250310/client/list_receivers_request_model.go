// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReceiversRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListReceiversRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListReceiversRequest
	GetPageToken() *string
	SetReceiverName(v string) *ListReceiversRequest
	GetReceiverName() *string
}

type ListReceiversRequest struct {
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	// example:
	//
	// receiver_name
	ReceiverName *string `json:"receiverName,omitempty" xml:"receiverName,omitempty"`
}

func (s ListReceiversRequest) String() string {
	return dara.Prettify(s)
}

func (s ListReceiversRequest) GoString() string {
	return s.String()
}

func (s *ListReceiversRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListReceiversRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListReceiversRequest) GetReceiverName() *string {
	return s.ReceiverName
}

func (s *ListReceiversRequest) SetMaxResults(v int32) *ListReceiversRequest {
	s.MaxResults = &v
	return s
}

func (s *ListReceiversRequest) SetPageToken(v string) *ListReceiversRequest {
	s.PageToken = &v
	return s
}

func (s *ListReceiversRequest) SetReceiverName(v string) *ListReceiversRequest {
	s.ReceiverName = &v
	return s
}

func (s *ListReceiversRequest) Validate() error {
	return dara.Validate(s)
}
