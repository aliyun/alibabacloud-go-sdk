// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataEventSelectorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataEventSelectorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataEventSelectorsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataEventSelectorsResponseBody) *ListDataEventSelectorsResponse
	GetBody() *ListDataEventSelectorsResponseBody
}

type ListDataEventSelectorsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataEventSelectorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataEventSelectorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataEventSelectorsResponse) GoString() string {
	return s.String()
}

func (s *ListDataEventSelectorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataEventSelectorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataEventSelectorsResponse) GetBody() *ListDataEventSelectorsResponseBody {
	return s.Body
}

func (s *ListDataEventSelectorsResponse) SetHeaders(v map[string]*string) *ListDataEventSelectorsResponse {
	s.Headers = v
	return s
}

func (s *ListDataEventSelectorsResponse) SetStatusCode(v int32) *ListDataEventSelectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataEventSelectorsResponse) SetBody(v *ListDataEventSelectorsResponseBody) *ListDataEventSelectorsResponse {
	s.Body = v
	return s
}

func (s *ListDataEventSelectorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
