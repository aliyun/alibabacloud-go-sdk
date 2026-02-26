// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGoodsShippingNoticeCreateResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GoodsShippingNoticeCreateResult
	GetRequestId() *string
	SetResult(v string) *GoodsShippingNoticeCreateResult
	GetResult() *string
}

type GoodsShippingNoticeCreateResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GoodsShippingNoticeCreateResult) String() string {
	return dara.Prettify(s)
}

func (s GoodsShippingNoticeCreateResult) GoString() string {
	return s.String()
}

func (s *GoodsShippingNoticeCreateResult) GetRequestId() *string {
	return s.RequestId
}

func (s *GoodsShippingNoticeCreateResult) GetResult() *string {
	return s.Result
}

func (s *GoodsShippingNoticeCreateResult) SetRequestId(v string) *GoodsShippingNoticeCreateResult {
	s.RequestId = &v
	return s
}

func (s *GoodsShippingNoticeCreateResult) SetResult(v string) *GoodsShippingNoticeCreateResult {
	s.Result = &v
	return s
}

func (s *GoodsShippingNoticeCreateResult) Validate() error {
	return dara.Validate(s)
}
