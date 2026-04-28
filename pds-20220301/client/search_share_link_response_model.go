// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchShareLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchShareLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchShareLinkResponse
	GetStatusCode() *int32
	SetBody(v *SearchShareLinkResponseBody) *SearchShareLinkResponse
	GetBody() *SearchShareLinkResponseBody
}

type SearchShareLinkResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchShareLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchShareLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchShareLinkResponse) GoString() string {
	return s.String()
}

func (s *SearchShareLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchShareLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchShareLinkResponse) GetBody() *SearchShareLinkResponseBody {
	return s.Body
}

func (s *SearchShareLinkResponse) SetHeaders(v map[string]*string) *SearchShareLinkResponse {
	s.Headers = v
	return s
}

func (s *SearchShareLinkResponse) SetStatusCode(v int32) *SearchShareLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchShareLinkResponse) SetBody(v *SearchShareLinkResponseBody) *SearchShareLinkResponse {
	s.Body = v
	return s
}

func (s *SearchShareLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
