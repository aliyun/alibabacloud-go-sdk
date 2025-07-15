// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParametersByPathRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *GetParametersByPathRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetParametersByPathRequest
	GetNextToken() *string
	SetPath(v string) *GetParametersByPathRequest
	GetPath() *string
	SetRecursive(v bool) *GetParametersByPathRequest
	GetRecursive() *bool
	SetRegionId(v string) *GetParametersByPathRequest
	GetRegionId() *string
}

type GetParametersByPathRequest struct {
	// The number of entries per page.
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
	// The path of the parameter. For example, if the name of a parameter is /path/path1/Myparameter, the path of the parameter is /path/path1/.
	//
	// This parameter is required.
	//
	// example:
	//
	// /parameter
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
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetParametersByPathRequest) String() string {
	return dara.Prettify(s)
}

func (s GetParametersByPathRequest) GoString() string {
	return s.String()
}

func (s *GetParametersByPathRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetParametersByPathRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetParametersByPathRequest) GetPath() *string {
	return s.Path
}

func (s *GetParametersByPathRequest) GetRecursive() *bool {
	return s.Recursive
}

func (s *GetParametersByPathRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetParametersByPathRequest) SetMaxResults(v int32) *GetParametersByPathRequest {
	s.MaxResults = &v
	return s
}

func (s *GetParametersByPathRequest) SetNextToken(v string) *GetParametersByPathRequest {
	s.NextToken = &v
	return s
}

func (s *GetParametersByPathRequest) SetPath(v string) *GetParametersByPathRequest {
	s.Path = &v
	return s
}

func (s *GetParametersByPathRequest) SetRecursive(v bool) *GetParametersByPathRequest {
	s.Recursive = &v
	return s
}

func (s *GetParametersByPathRequest) SetRegionId(v string) *GetParametersByPathRequest {
	s.RegionId = &v
	return s
}

func (s *GetParametersByPathRequest) Validate() error {
	return dara.Validate(s)
}
