// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStandardSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StandardSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StandardSearchResponse
	GetStatusCode() *int32
	SetBody(v *StandardSearchResponseBody) *StandardSearchResponse
	GetBody() *StandardSearchResponseBody
}

type StandardSearchResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StandardSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StandardSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s StandardSearchResponse) GoString() string {
	return s.String()
}

func (s *StandardSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StandardSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StandardSearchResponse) GetBody() *StandardSearchResponseBody {
	return s.Body
}

func (s *StandardSearchResponse) SetHeaders(v map[string]*string) *StandardSearchResponse {
	s.Headers = v
	return s
}

func (s *StandardSearchResponse) SetStatusCode(v int32) *StandardSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *StandardSearchResponse) SetBody(v *StandardSearchResponseBody) *StandardSearchResponse {
	s.Body = v
	return s
}

func (s *StandardSearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
