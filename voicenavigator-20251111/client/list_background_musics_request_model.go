// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackgroundMusicsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListBackgroundMusicsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListBackgroundMusicsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBackgroundMusicsRequest
	GetPageSize() *int32
}

type ListBackgroundMusicsRequest struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListBackgroundMusicsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBackgroundMusicsRequest) GoString() string {
	return s.String()
}

func (s *ListBackgroundMusicsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBackgroundMusicsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBackgroundMusicsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBackgroundMusicsRequest) SetInstanceId(v string) *ListBackgroundMusicsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBackgroundMusicsRequest) SetPageNumber(v int32) *ListBackgroundMusicsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBackgroundMusicsRequest) SetPageSize(v int32) *ListBackgroundMusicsRequest {
	s.PageSize = &v
	return s
}

func (s *ListBackgroundMusicsRequest) Validate() error {
	return dara.Validate(s)
}
