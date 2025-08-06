// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPartnerProduceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *QueryPartnerProduceListRequest
	GetBizId() *string
	SetBizType(v string) *QueryPartnerProduceListRequest
	GetBizType() *string
	SetPageNum(v int64) *QueryPartnerProduceListRequest
	GetPageNum() *int64
	SetPageSize(v int64) *QueryPartnerProduceListRequest
	GetPageSize() *int64
}

type QueryPartnerProduceListRequest struct {
	// example:
	//
	// P20211216204717000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// esp.cdn
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryPartnerProduceListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPartnerProduceListRequest) GoString() string {
	return s.String()
}

func (s *QueryPartnerProduceListRequest) GetBizId() *string {
	return s.BizId
}

func (s *QueryPartnerProduceListRequest) GetBizType() *string {
	return s.BizType
}

func (s *QueryPartnerProduceListRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *QueryPartnerProduceListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryPartnerProduceListRequest) SetBizId(v string) *QueryPartnerProduceListRequest {
	s.BizId = &v
	return s
}

func (s *QueryPartnerProduceListRequest) SetBizType(v string) *QueryPartnerProduceListRequest {
	s.BizType = &v
	return s
}

func (s *QueryPartnerProduceListRequest) SetPageNum(v int64) *QueryPartnerProduceListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryPartnerProduceListRequest) SetPageSize(v int64) *QueryPartnerProduceListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPartnerProduceListRequest) Validate() error {
	return dara.Validate(s)
}
