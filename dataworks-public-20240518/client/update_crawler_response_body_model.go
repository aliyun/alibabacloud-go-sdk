// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCrawlerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpdateCrawlerResponseBody
	GetId() *int64
	SetRequestId(v string) *UpdateCrawlerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCrawlerResponseBody
	GetSuccess() *bool
}

type UpdateCrawlerResponseBody struct {
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

func (s UpdateCrawlerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrawlerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCrawlerResponseBody) GetId() *int64 {
	return s.Id
}

func (s *UpdateCrawlerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCrawlerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCrawlerResponseBody) SetId(v int64) *UpdateCrawlerResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateCrawlerResponseBody) SetRequestId(v string) *UpdateCrawlerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCrawlerResponseBody) SetSuccess(v bool) *UpdateCrawlerResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCrawlerResponseBody) Validate() error {
	return dara.Validate(s)
}
