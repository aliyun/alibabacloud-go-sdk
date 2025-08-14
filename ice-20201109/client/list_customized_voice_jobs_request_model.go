// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomizedVoiceJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListCustomizedVoiceJobsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListCustomizedVoiceJobsRequest
	GetPageSize() *int32
	SetType(v string) *ListCustomizedVoiceJobsRequest
	GetType() *string
}

type ListCustomizedVoiceJobsRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the human voice cloning job. Valid values:
	//
	// 	- Basic
	//
	// 	- Standard
	//
	// > : If you do not specify this parameter, the default value Basic is used.
	//
	// example:
	//
	// Standard
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCustomizedVoiceJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomizedVoiceJobsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomizedVoiceJobsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListCustomizedVoiceJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomizedVoiceJobsRequest) GetType() *string {
	return s.Type
}

func (s *ListCustomizedVoiceJobsRequest) SetPageNo(v int32) *ListCustomizedVoiceJobsRequest {
	s.PageNo = &v
	return s
}

func (s *ListCustomizedVoiceJobsRequest) SetPageSize(v int32) *ListCustomizedVoiceJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomizedVoiceJobsRequest) SetType(v string) *ListCustomizedVoiceJobsRequest {
	s.Type = &v
	return s
}

func (s *ListCustomizedVoiceJobsRequest) Validate() error {
	return dara.Validate(s)
}
