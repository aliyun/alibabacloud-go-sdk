// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchInfoResponse
	GetStatusCode() *int32
	SetBody(v *SearchInfoResponseBody) *SearchInfoResponse
	GetBody() *SearchInfoResponseBody
}

type SearchInfoResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchInfoResponse) GoString() string {
	return s.String()
}

func (s *SearchInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchInfoResponse) GetBody() *SearchInfoResponseBody {
	return s.Body
}

func (s *SearchInfoResponse) SetHeaders(v map[string]*string) *SearchInfoResponse {
	s.Headers = v
	return s
}

func (s *SearchInfoResponse) SetStatusCode(v int32) *SearchInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchInfoResponse) SetBody(v *SearchInfoResponseBody) *SearchInfoResponse {
	s.Body = v
	return s
}

func (s *SearchInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
