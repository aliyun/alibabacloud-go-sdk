// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFaceDbsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int64) *ListFaceDbsRequest
	GetLimit() *int64
	SetOffset(v int64) *ListFaceDbsRequest
	GetOffset() *int64
}

type ListFaceDbsRequest struct {
	// example:
	//
	// 50
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// 1
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
}

func (s ListFaceDbsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFaceDbsRequest) GoString() string {
	return s.String()
}

func (s *ListFaceDbsRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ListFaceDbsRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *ListFaceDbsRequest) SetLimit(v int64) *ListFaceDbsRequest {
	s.Limit = &v
	return s
}

func (s *ListFaceDbsRequest) SetOffset(v int64) *ListFaceDbsRequest {
	s.Offset = &v
	return s
}

func (s *ListFaceDbsRequest) Validate() error {
	return dara.Validate(s)
}
