// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPartnerIntentionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *QueryPartnerIntentionListRequest
	GetBizId() *string
	SetBizType(v string) *QueryPartnerIntentionListRequest
	GetBizType() *string
	SetPageNum(v int64) *QueryPartnerIntentionListRequest
	GetPageNum() *int64
	SetPageSize(v int64) *QueryPartnerIntentionListRequest
	GetPageSize() *int64
}

type QueryPartnerIntentionListRequest struct {
	// example:
	//
	// I20211117092704000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// esp.wangwen
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 10
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryPartnerIntentionListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPartnerIntentionListRequest) GoString() string {
	return s.String()
}

func (s *QueryPartnerIntentionListRequest) GetBizId() *string {
	return s.BizId
}

func (s *QueryPartnerIntentionListRequest) GetBizType() *string {
	return s.BizType
}

func (s *QueryPartnerIntentionListRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *QueryPartnerIntentionListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryPartnerIntentionListRequest) SetBizId(v string) *QueryPartnerIntentionListRequest {
	s.BizId = &v
	return s
}

func (s *QueryPartnerIntentionListRequest) SetBizType(v string) *QueryPartnerIntentionListRequest {
	s.BizType = &v
	return s
}

func (s *QueryPartnerIntentionListRequest) SetPageNum(v int64) *QueryPartnerIntentionListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryPartnerIntentionListRequest) SetPageSize(v int64) *QueryPartnerIntentionListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPartnerIntentionListRequest) Validate() error {
	return dara.Validate(s)
}
