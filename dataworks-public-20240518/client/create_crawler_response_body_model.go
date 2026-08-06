// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCrawlerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateCrawlerResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateCrawlerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCrawlerResponseBody
	GetSuccess() *bool
}

type CreateCrawlerResponseBody struct {
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

func (s CreateCrawlerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCrawlerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCrawlerResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateCrawlerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCrawlerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCrawlerResponseBody) SetId(v int64) *CreateCrawlerResponseBody {
	s.Id = &v
	return s
}

func (s *CreateCrawlerResponseBody) SetRequestId(v string) *CreateCrawlerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCrawlerResponseBody) SetSuccess(v bool) *CreateCrawlerResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCrawlerResponseBody) Validate() error {
	return dara.Validate(s)
}
