// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCrawlerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteCrawlerResponseBody
	GetId() *int64
	SetRequestId(v string) *DeleteCrawlerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCrawlerResponseBody
	GetSuccess() *bool
}

type DeleteCrawlerResponseBody struct {
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 9252F32F-D855-549E-8898-61CF5A733050
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCrawlerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCrawlerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCrawlerResponseBody) GetId() *int64 {
	return s.Id
}

func (s *DeleteCrawlerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCrawlerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCrawlerResponseBody) SetId(v int64) *DeleteCrawlerResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteCrawlerResponseBody) SetRequestId(v string) *DeleteCrawlerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCrawlerResponseBody) SetSuccess(v bool) *DeleteCrawlerResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCrawlerResponseBody) Validate() error {
	return dara.Validate(s)
}
