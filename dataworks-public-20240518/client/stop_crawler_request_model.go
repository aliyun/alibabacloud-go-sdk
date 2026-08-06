// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCrawlerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *StopCrawlerRequest
	GetId() *int64
}

type StopCrawlerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s StopCrawlerRequest) String() string {
	return dara.Prettify(s)
}

func (s StopCrawlerRequest) GoString() string {
	return s.String()
}

func (s *StopCrawlerRequest) GetId() *int64 {
	return s.Id
}

func (s *StopCrawlerRequest) SetId(v int64) *StopCrawlerRequest {
	s.Id = &v
	return s
}

func (s *StopCrawlerRequest) Validate() error {
	return dara.Validate(s)
}
