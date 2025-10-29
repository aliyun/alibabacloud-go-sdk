// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadPageScrapeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ReadPageScrapeBody) *ReadPageScrapeRequest
	GetBody() *ReadPageScrapeBody
}

type ReadPageScrapeRequest struct {
	// post body
	Body *ReadPageScrapeBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadPageScrapeRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadPageScrapeRequest) GoString() string {
	return s.String()
}

func (s *ReadPageScrapeRequest) GetBody() *ReadPageScrapeBody {
	return s.Body
}

func (s *ReadPageScrapeRequest) SetBody(v *ReadPageScrapeBody) *ReadPageScrapeRequest {
	s.Body = v
	return s
}

func (s *ReadPageScrapeRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
