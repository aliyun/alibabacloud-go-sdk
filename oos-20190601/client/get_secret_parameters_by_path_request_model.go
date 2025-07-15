// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretParametersByPathRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *GetSecretParametersByPathRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetSecretParametersByPathRequest
	GetNextToken() *string
	SetPath(v string) *GetSecretParametersByPathRequest
	GetPath() *string
	SetRecursive(v bool) *GetSecretParametersByPathRequest
	GetRecursive() *bool
	SetRegionId(v string) *GetSecretParametersByPathRequest
	GetRegionId() *string
	SetWithDecryption(v bool) *GetSecretParametersByPathRequest
	GetWithDecryption() *bool
}

type GetSecretParametersByPathRequest struct {
	// The number of entries per page. Valid values: 1 to 10. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The path of the encryption parameter. The path must be 1 to 200 characters in length. For example, if the name of an encryption parameter is /secretParameter/mySecretParameter, the path of the encryption parameter is /secretParameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// /secretParameter
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Specifies whether to recursively query encryption parameters from all levels of directories in the specified path. Valid values: true and false. For example, if you want to query the /secretParameter/mySecretParameter and /secretParameter/secretParameter 1/mySecretParameter parameters, the valid values specify the parameters to be returned.
	//
	// 	- true: returns both of the /secretParameter/mySecretParameter and /secretParameter/secretParameter1/mySecretParameter parameters.
	//
	// 	- false: returns only the /secretParameter/mySecretParameter parameter.
	//
	// example:
	//
	// false
	Recursive *bool `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to decrypt the parameter value. Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	WithDecryption *bool `json:"WithDecryption,omitempty" xml:"WithDecryption,omitempty"`
}

func (s GetSecretParametersByPathRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecretParametersByPathRequest) GoString() string {
	return s.String()
}

func (s *GetSecretParametersByPathRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetSecretParametersByPathRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetSecretParametersByPathRequest) GetPath() *string {
	return s.Path
}

func (s *GetSecretParametersByPathRequest) GetRecursive() *bool {
	return s.Recursive
}

func (s *GetSecretParametersByPathRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSecretParametersByPathRequest) GetWithDecryption() *bool {
	return s.WithDecryption
}

func (s *GetSecretParametersByPathRequest) SetMaxResults(v int32) *GetSecretParametersByPathRequest {
	s.MaxResults = &v
	return s
}

func (s *GetSecretParametersByPathRequest) SetNextToken(v string) *GetSecretParametersByPathRequest {
	s.NextToken = &v
	return s
}

func (s *GetSecretParametersByPathRequest) SetPath(v string) *GetSecretParametersByPathRequest {
	s.Path = &v
	return s
}

func (s *GetSecretParametersByPathRequest) SetRecursive(v bool) *GetSecretParametersByPathRequest {
	s.Recursive = &v
	return s
}

func (s *GetSecretParametersByPathRequest) SetRegionId(v string) *GetSecretParametersByPathRequest {
	s.RegionId = &v
	return s
}

func (s *GetSecretParametersByPathRequest) SetWithDecryption(v bool) *GetSecretParametersByPathRequest {
	s.WithDecryption = &v
	return s
}

func (s *GetSecretParametersByPathRequest) Validate() error {
	return dara.Validate(s)
}
