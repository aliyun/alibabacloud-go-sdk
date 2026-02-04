// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEventTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEventTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListEventTypesResponseBody) *ListEventTypesResponse
	GetBody() *ListEventTypesResponseBody
}

type ListEventTypesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEventTypesResponse) GoString() string {
	return s.String()
}

func (s *ListEventTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEventTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEventTypesResponse) GetBody() *ListEventTypesResponseBody {
	return s.Body
}

func (s *ListEventTypesResponse) SetHeaders(v map[string]*string) *ListEventTypesResponse {
	s.Headers = v
	return s
}

func (s *ListEventTypesResponse) SetStatusCode(v int32) *ListEventTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventTypesResponse) SetBody(v *ListEventTypesResponseBody) *ListEventTypesResponse {
	s.Body = v
	return s
}

func (s *ListEventTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
