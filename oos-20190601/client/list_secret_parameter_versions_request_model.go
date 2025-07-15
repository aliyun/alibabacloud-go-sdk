// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretParameterVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSecretParameterVersionsRequest
	GetMaxResults() *int32
	SetName(v string) *ListSecretParameterVersionsRequest
	GetName() *string
	SetNextToken(v string) *ListSecretParameterVersionsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListSecretParameterVersionsRequest
	GetRegionId() *string
	SetShareType(v string) *ListSecretParameterVersionsRequest
	GetShareType() *string
	SetWithDecryption(v bool) *ListSecretParameterVersionsRequest
	GetWithDecryption() *bool
}

type ListSecretParameterVersionsRequest struct {
	// The number of entries per page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the encryption parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The share type of the encryption parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// Specifies whether to decrypt the parameter value. The decrypted parameter value is returned only if this parameter is set to true. Otherwise, null is returned.
	//
	// example:
	//
	// false
	WithDecryption *bool `json:"WithDecryption,omitempty" xml:"WithDecryption,omitempty"`
}

func (s ListSecretParameterVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecretParameterVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListSecretParameterVersionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSecretParameterVersionsRequest) GetName() *string {
	return s.Name
}

func (s *ListSecretParameterVersionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSecretParameterVersionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSecretParameterVersionsRequest) GetShareType() *string {
	return s.ShareType
}

func (s *ListSecretParameterVersionsRequest) GetWithDecryption() *bool {
	return s.WithDecryption
}

func (s *ListSecretParameterVersionsRequest) SetMaxResults(v int32) *ListSecretParameterVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSecretParameterVersionsRequest) SetName(v string) *ListSecretParameterVersionsRequest {
	s.Name = &v
	return s
}

func (s *ListSecretParameterVersionsRequest) SetNextToken(v string) *ListSecretParameterVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSecretParameterVersionsRequest) SetRegionId(v string) *ListSecretParameterVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSecretParameterVersionsRequest) SetShareType(v string) *ListSecretParameterVersionsRequest {
	s.ShareType = &v
	return s
}

func (s *ListSecretParameterVersionsRequest) SetWithDecryption(v bool) *ListSecretParameterVersionsRequest {
	s.WithDecryption = &v
	return s
}

func (s *ListSecretParameterVersionsRequest) Validate() error {
	return dara.Validate(s)
}
