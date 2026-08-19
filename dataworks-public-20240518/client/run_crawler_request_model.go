// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCrawlerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *RunCrawlerRequest
	GetId() *int64
}

type RunCrawlerRequest struct {
	// The ID of the metadata crawler. You can call ListCrawlers to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s RunCrawlerRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCrawlerRequest) GoString() string {
	return s.String()
}

func (s *RunCrawlerRequest) GetId() *int64 {
	return s.Id
}

func (s *RunCrawlerRequest) SetId(v int64) *RunCrawlerRequest {
	s.Id = &v
	return s
}

func (s *RunCrawlerRequest) Validate() error {
	return dara.Validate(s)
}
