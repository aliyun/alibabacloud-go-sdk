// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCrawlerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteCrawlerRequest
	GetId() *int64
}

type DeleteCrawlerRequest struct {
	// The ID of the metadata crawler. You can call ListCrawlers to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteCrawlerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCrawlerRequest) GoString() string {
	return s.String()
}

func (s *DeleteCrawlerRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteCrawlerRequest) SetId(v int64) *DeleteCrawlerRequest {
	s.Id = &v
	return s
}

func (s *DeleteCrawlerRequest) Validate() error {
	return dara.Validate(s)
}
