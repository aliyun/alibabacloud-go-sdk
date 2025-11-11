// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncWritingBiddingDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AsyncWritingBiddingDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AsyncWritingBiddingDocResponse
	GetStatusCode() *int32
	SetBody(v *AsyncWritingBiddingDocResponseBody) *AsyncWritingBiddingDocResponse
	GetBody() *AsyncWritingBiddingDocResponseBody
}

type AsyncWritingBiddingDocResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsyncWritingBiddingDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsyncWritingBiddingDocResponse) String() string {
	return dara.Prettify(s)
}

func (s AsyncWritingBiddingDocResponse) GoString() string {
	return s.String()
}

func (s *AsyncWritingBiddingDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AsyncWritingBiddingDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AsyncWritingBiddingDocResponse) GetBody() *AsyncWritingBiddingDocResponseBody {
	return s.Body
}

func (s *AsyncWritingBiddingDocResponse) SetHeaders(v map[string]*string) *AsyncWritingBiddingDocResponse {
	s.Headers = v
	return s
}

func (s *AsyncWritingBiddingDocResponse) SetStatusCode(v int32) *AsyncWritingBiddingDocResponse {
	s.StatusCode = &v
	return s
}

func (s *AsyncWritingBiddingDocResponse) SetBody(v *AsyncWritingBiddingDocResponseBody) *AsyncWritingBiddingDocResponse {
	s.Body = v
	return s
}

func (s *AsyncWritingBiddingDocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
