// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConversationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTimeLeftRange(v int64) *QueryConversationsRequest
	GetBeginTimeLeftRange() *int64
	SetBeginTimeRightRange(v int64) *QueryConversationsRequest
	GetBeginTimeRightRange() *int64
	SetCallingNumber(v string) *QueryConversationsRequest
	GetCallingNumber() *string
	SetInstanceId(v string) *QueryConversationsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *QueryConversationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryConversationsRequest
	GetPageSize() *int32
}

type QueryConversationsRequest struct {
	// example:
	//
	// 1582183381000
	BeginTimeLeftRange *int64 `json:"BeginTimeLeftRange,omitempty" xml:"BeginTimeLeftRange,omitempty"`
	// example:
	//
	// 1582356181000
	BeginTimeRightRange *int64 `json:"BeginTimeRightRange,omitempty" xml:"BeginTimeRightRange,omitempty"`
	// example:
	//
	// 02811111111
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12f407b22cbe4890ac595f09985848d5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryConversationsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConversationsRequest) GoString() string {
	return s.String()
}

func (s *QueryConversationsRequest) GetBeginTimeLeftRange() *int64 {
	return s.BeginTimeLeftRange
}

func (s *QueryConversationsRequest) GetBeginTimeRightRange() *int64 {
	return s.BeginTimeRightRange
}

func (s *QueryConversationsRequest) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *QueryConversationsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryConversationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryConversationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryConversationsRequest) SetBeginTimeLeftRange(v int64) *QueryConversationsRequest {
	s.BeginTimeLeftRange = &v
	return s
}

func (s *QueryConversationsRequest) SetBeginTimeRightRange(v int64) *QueryConversationsRequest {
	s.BeginTimeRightRange = &v
	return s
}

func (s *QueryConversationsRequest) SetCallingNumber(v string) *QueryConversationsRequest {
	s.CallingNumber = &v
	return s
}

func (s *QueryConversationsRequest) SetInstanceId(v string) *QueryConversationsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryConversationsRequest) SetPageNumber(v int32) *QueryConversationsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryConversationsRequest) SetPageSize(v int32) *QueryConversationsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryConversationsRequest) Validate() error {
	return dara.Validate(s)
}
