// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeTask2Request interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v int64) *ListComputeTask2Request
	GetClientId() *int64
	SetPageNum(v int32) *ListComputeTask2Request
	GetPageNum() *int32
	SetPageSize(v int32) *ListComputeTask2Request
	GetPageSize() *int32
}

type ListComputeTask2Request struct {
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	PageNum  *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListComputeTask2Request) String() string {
	return dara.Prettify(s)
}

func (s ListComputeTask2Request) GoString() string {
	return s.String()
}

func (s *ListComputeTask2Request) GetClientId() *int64 {
	return s.ClientId
}

func (s *ListComputeTask2Request) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListComputeTask2Request) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComputeTask2Request) SetClientId(v int64) *ListComputeTask2Request {
	s.ClientId = &v
	return s
}

func (s *ListComputeTask2Request) SetPageNum(v int32) *ListComputeTask2Request {
	s.PageNum = &v
	return s
}

func (s *ListComputeTask2Request) SetPageSize(v int32) *ListComputeTask2Request {
	s.PageSize = &v
	return s
}

func (s *ListComputeTask2Request) Validate() error {
	return dara.Validate(s)
}
