// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProduceAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListProduceAuthorizationRequest
	GetBizId() *string
	SetBizType(v string) *ListProduceAuthorizationRequest
	GetBizType() *string
	SetPageNum(v int32) *ListProduceAuthorizationRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListProduceAuthorizationRequest
	GetPageSize() *int32
}

type ListProduceAuthorizationRequest struct {
	// example:
	//
	// P20210709190452000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// esp.bookkeeping_ai
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListProduceAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProduceAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *ListProduceAuthorizationRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListProduceAuthorizationRequest) GetBizType() *string {
	return s.BizType
}

func (s *ListProduceAuthorizationRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListProduceAuthorizationRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProduceAuthorizationRequest) SetBizId(v string) *ListProduceAuthorizationRequest {
	s.BizId = &v
	return s
}

func (s *ListProduceAuthorizationRequest) SetBizType(v string) *ListProduceAuthorizationRequest {
	s.BizType = &v
	return s
}

func (s *ListProduceAuthorizationRequest) SetPageNum(v int32) *ListProduceAuthorizationRequest {
	s.PageNum = &v
	return s
}

func (s *ListProduceAuthorizationRequest) SetPageSize(v int32) *ListProduceAuthorizationRequest {
	s.PageSize = &v
	return s
}

func (s *ListProduceAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
