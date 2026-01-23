// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStandardsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStandardsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStandardsResponse
	GetStatusCode() *int32
	SetBody(v *ListStandardsResponseBody) *ListStandardsResponse
	GetBody() *ListStandardsResponseBody
}

type ListStandardsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStandardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStandardsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponse) GoString() string {
	return s.String()
}

func (s *ListStandardsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStandardsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStandardsResponse) GetBody() *ListStandardsResponseBody {
	return s.Body
}

func (s *ListStandardsResponse) SetHeaders(v map[string]*string) *ListStandardsResponse {
	s.Headers = v
	return s
}

func (s *ListStandardsResponse) SetStatusCode(v int32) *ListStandardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStandardsResponse) SetBody(v *ListStandardsResponseBody) *ListStandardsResponse {
	s.Body = v
	return s
}

func (s *ListStandardsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
