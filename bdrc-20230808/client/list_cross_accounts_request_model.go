// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrossAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccountOwnerId(v int64) *ListCrossAccountsRequest
	GetCrossAccountOwnerId() *int64
	SetManagementMode(v string) *ListCrossAccountsRequest
	GetManagementMode() *string
	SetMaxResults(v int32) *ListCrossAccountsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCrossAccountsRequest
	GetNextToken() *string
	SetTargetId(v string) *ListCrossAccountsRequest
	GetTargetId() *string
	SetTargetType(v string) *ListCrossAccountsRequest
	GetTargetType() *string
}

type ListCrossAccountsRequest struct {
	// example:
	//
	// 123***7890
	CrossAccountOwnerId *int64 `json:"CrossAccountOwnerId,omitempty" xml:"CrossAccountOwnerId,omitempty"`
	// example:
	//
	// MANUAL
	ManagementMode *string `json:"ManagementMode,omitempty" xml:"ManagementMode,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// cae**********699
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 123***7890
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// example:
	//
	// ACCOUNT
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListCrossAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCrossAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListCrossAccountsRequest) GetCrossAccountOwnerId() *int64 {
	return s.CrossAccountOwnerId
}

func (s *ListCrossAccountsRequest) GetManagementMode() *string {
	return s.ManagementMode
}

func (s *ListCrossAccountsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCrossAccountsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCrossAccountsRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *ListCrossAccountsRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ListCrossAccountsRequest) SetCrossAccountOwnerId(v int64) *ListCrossAccountsRequest {
	s.CrossAccountOwnerId = &v
	return s
}

func (s *ListCrossAccountsRequest) SetManagementMode(v string) *ListCrossAccountsRequest {
	s.ManagementMode = &v
	return s
}

func (s *ListCrossAccountsRequest) SetMaxResults(v int32) *ListCrossAccountsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCrossAccountsRequest) SetNextToken(v string) *ListCrossAccountsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCrossAccountsRequest) SetTargetId(v string) *ListCrossAccountsRequest {
	s.TargetId = &v
	return s
}

func (s *ListCrossAccountsRequest) SetTargetType(v string) *ListCrossAccountsRequest {
	s.TargetType = &v
	return s
}

func (s *ListCrossAccountsRequest) Validate() error {
	return dara.Validate(s)
}
