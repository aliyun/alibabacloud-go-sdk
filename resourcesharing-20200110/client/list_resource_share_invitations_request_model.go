// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceShareInvitationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResourceShareInvitationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceShareInvitationsRequest
	GetNextToken() *string
	SetResourceShareIds(v []*string) *ListResourceShareInvitationsRequest
	GetResourceShareIds() []*string
	SetResourceShareInvitationIds(v []*string) *ListResourceShareInvitationsRequest
	GetResourceShareInvitationIds() []*string
}

type ListResourceShareInvitationsRequest struct {
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The IDs of the resource shares.
	ResourceShareIds []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	// The IDs of the resource sharing invitations.
	ResourceShareInvitationIds []*string `json:"ResourceShareInvitationIds,omitempty" xml:"ResourceShareInvitationIds,omitempty" type:"Repeated"`
}

func (s ListResourceShareInvitationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceShareInvitationsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceShareInvitationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceShareInvitationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceShareInvitationsRequest) GetResourceShareIds() []*string {
	return s.ResourceShareIds
}

func (s *ListResourceShareInvitationsRequest) GetResourceShareInvitationIds() []*string {
	return s.ResourceShareInvitationIds
}

func (s *ListResourceShareInvitationsRequest) SetMaxResults(v int32) *ListResourceShareInvitationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceShareInvitationsRequest) SetNextToken(v string) *ListResourceShareInvitationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceShareInvitationsRequest) SetResourceShareIds(v []*string) *ListResourceShareInvitationsRequest {
	s.ResourceShareIds = v
	return s
}

func (s *ListResourceShareInvitationsRequest) SetResourceShareInvitationIds(v []*string) *ListResourceShareInvitationsRequest {
	s.ResourceShareInvitationIds = v
	return s
}

func (s *ListResourceShareInvitationsRequest) Validate() error {
	return dara.Validate(s)
}
