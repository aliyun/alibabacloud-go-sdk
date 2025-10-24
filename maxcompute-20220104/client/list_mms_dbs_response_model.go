// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsDbsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMmsDbsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMmsDbsResponse
	GetStatusCode() *int32
	SetBody(v *ListMmsDbsResponseBody) *ListMmsDbsResponse
	GetBody() *ListMmsDbsResponseBody
}

type ListMmsDbsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsDbsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsDbsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDbsResponse) GoString() string {
	return s.String()
}

func (s *ListMmsDbsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMmsDbsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMmsDbsResponse) GetBody() *ListMmsDbsResponseBody {
	return s.Body
}

func (s *ListMmsDbsResponse) SetHeaders(v map[string]*string) *ListMmsDbsResponse {
	s.Headers = v
	return s
}

func (s *ListMmsDbsResponse) SetStatusCode(v int32) *ListMmsDbsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsDbsResponse) SetBody(v *ListMmsDbsResponseBody) *ListMmsDbsResponse {
	s.Body = v
	return s
}

func (s *ListMmsDbsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
