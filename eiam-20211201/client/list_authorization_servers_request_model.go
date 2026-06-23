// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListAuthorizationServersRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListAuthorizationServersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAuthorizationServersRequest
	GetNextToken() *string
}

type ListAuthorizationServersRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page for a paged query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAuthorizationServersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationServersRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizationServersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizationServersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAuthorizationServersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuthorizationServersRequest) SetInstanceId(v string) *ListAuthorizationServersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizationServersRequest) SetMaxResults(v int32) *ListAuthorizationServersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizationServersRequest) SetNextToken(v string) *ListAuthorizationServersRequest {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizationServersRequest) Validate() error {
	return dara.Validate(s)
}
