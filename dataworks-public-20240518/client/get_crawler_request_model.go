// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCrawlerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetCrawlerRequest
	GetId() *int64
}

type GetCrawlerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetCrawlerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCrawlerRequest) GoString() string {
	return s.String()
}

func (s *GetCrawlerRequest) GetId() *int64 {
	return s.Id
}

func (s *GetCrawlerRequest) SetId(v int64) *GetCrawlerRequest {
	s.Id = &v
	return s
}

func (s *GetCrawlerRequest) Validate() error {
	return dara.Validate(s)
}
