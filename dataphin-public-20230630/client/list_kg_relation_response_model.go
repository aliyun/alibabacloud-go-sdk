// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKgRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKgRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKgRelationResponse
	GetStatusCode() *int32
	SetBody(v *ListKgRelationResponseBody) *ListKgRelationResponse
	GetBody() *ListKgRelationResponseBody
}

type ListKgRelationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKgRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKgRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKgRelationResponse) GoString() string {
	return s.String()
}

func (s *ListKgRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKgRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKgRelationResponse) GetBody() *ListKgRelationResponseBody {
	return s.Body
}

func (s *ListKgRelationResponse) SetHeaders(v map[string]*string) *ListKgRelationResponse {
	s.Headers = v
	return s
}

func (s *ListKgRelationResponse) SetStatusCode(v int32) *ListKgRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKgRelationResponse) SetBody(v *ListKgRelationResponseBody) *ListKgRelationResponse {
	s.Body = v
	return s
}

func (s *ListKgRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
