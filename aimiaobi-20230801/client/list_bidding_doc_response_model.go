// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBiddingDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBiddingDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBiddingDocResponse
	GetStatusCode() *int32
	SetBody(v *ListBiddingDocResponseBody) *ListBiddingDocResponse
	GetBody() *ListBiddingDocResponseBody
}

type ListBiddingDocResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBiddingDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBiddingDocResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBiddingDocResponse) GoString() string {
	return s.String()
}

func (s *ListBiddingDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBiddingDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBiddingDocResponse) GetBody() *ListBiddingDocResponseBody {
	return s.Body
}

func (s *ListBiddingDocResponse) SetHeaders(v map[string]*string) *ListBiddingDocResponse {
	s.Headers = v
	return s
}

func (s *ListBiddingDocResponse) SetStatusCode(v int32) *ListBiddingDocResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBiddingDocResponse) SetBody(v *ListBiddingDocResponseBody) *ListBiddingDocResponse {
	s.Body = v
	return s
}

func (s *ListBiddingDocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
