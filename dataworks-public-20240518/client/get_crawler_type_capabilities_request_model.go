// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCrawlerTypeCapabilitiesRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetCrawlerTypeCapabilitiesRequest struct {
}

func (s GetCrawlerTypeCapabilitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCrawlerTypeCapabilitiesRequest) GoString() string {
	return s.String()
}

func (s *GetCrawlerTypeCapabilitiesRequest) Validate() error {
	return dara.Validate(s)
}
