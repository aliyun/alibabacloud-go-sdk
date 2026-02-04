// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRoutersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRoutersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRoutersResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRoutersResponseBody) *ListTransitRoutersResponse
	GetBody() *ListTransitRoutersResponseBody
}

type ListTransitRoutersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRoutersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRoutersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRoutersResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRoutersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRoutersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRoutersResponse) GetBody() *ListTransitRoutersResponseBody {
	return s.Body
}

func (s *ListTransitRoutersResponse) SetHeaders(v map[string]*string) *ListTransitRoutersResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRoutersResponse) SetStatusCode(v int32) *ListTransitRoutersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRoutersResponse) SetBody(v *ListTransitRoutersResponseBody) *ListTransitRoutersResponse {
	s.Body = v
	return s
}

func (s *ListTransitRoutersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
