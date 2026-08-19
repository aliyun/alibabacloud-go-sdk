// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineEnvironmentVariablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *ListRoutineEnvironmentVariablesRequest
	GetEnv() *string
	SetKeyWord(v string) *ListRoutineEnvironmentVariablesRequest
	GetKeyWord() *string
	SetName(v string) *ListRoutineEnvironmentVariablesRequest
	GetName() *string
	SetPageNumber(v int64) *ListRoutineEnvironmentVariablesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRoutineEnvironmentVariablesRequest
	GetPageSize() *int64
}

type ListRoutineEnvironmentVariablesRequest struct {
	// The environment name.
	//
	// Valid values:
	//
	// - `production`: production environment
	//
	// - `staging`: staging environment
	//
	// This parameter is required.
	//
	// example:
	//
	// production
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The keyword used to perform a case-insensitive fuzzy search on environment variable keys.
	//
	// example:
	//
	// LOG
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The function name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_routine
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRoutineEnvironmentVariablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineEnvironmentVariablesRequest) GoString() string {
	return s.String()
}

func (s *ListRoutineEnvironmentVariablesRequest) GetEnv() *string {
	return s.Env
}

func (s *ListRoutineEnvironmentVariablesRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *ListRoutineEnvironmentVariablesRequest) GetName() *string {
	return s.Name
}

func (s *ListRoutineEnvironmentVariablesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRoutineEnvironmentVariablesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRoutineEnvironmentVariablesRequest) SetEnv(v string) *ListRoutineEnvironmentVariablesRequest {
	s.Env = &v
	return s
}

func (s *ListRoutineEnvironmentVariablesRequest) SetKeyWord(v string) *ListRoutineEnvironmentVariablesRequest {
	s.KeyWord = &v
	return s
}

func (s *ListRoutineEnvironmentVariablesRequest) SetName(v string) *ListRoutineEnvironmentVariablesRequest {
	s.Name = &v
	return s
}

func (s *ListRoutineEnvironmentVariablesRequest) SetPageNumber(v int64) *ListRoutineEnvironmentVariablesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRoutineEnvironmentVariablesRequest) SetPageSize(v int64) *ListRoutineEnvironmentVariablesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRoutineEnvironmentVariablesRequest) Validate() error {
	return dara.Validate(s)
}
