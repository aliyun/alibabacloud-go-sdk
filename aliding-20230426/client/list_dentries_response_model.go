// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDentriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDentriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDentriesResponse
	GetStatusCode() *int32
	SetBody(v *ListDentriesResponseBody) *ListDentriesResponse
	GetBody() *ListDentriesResponseBody
}

type ListDentriesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDentriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDentriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDentriesResponse) GoString() string {
	return s.String()
}

func (s *ListDentriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDentriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDentriesResponse) GetBody() *ListDentriesResponseBody {
	return s.Body
}

func (s *ListDentriesResponse) SetHeaders(v map[string]*string) *ListDentriesResponse {
	s.Headers = v
	return s
}

func (s *ListDentriesResponse) SetStatusCode(v int32) *ListDentriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDentriesResponse) SetBody(v *ListDentriesResponseBody) *ListDentriesResponse {
	s.Body = v
	return s
}

func (s *ListDentriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
