// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceRelationsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceRelationsResponseBody) *ListResourceRelationsResponse
	GetBody() *ListResourceRelationsResponseBody
}

type ListResourceRelationsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceRelationsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceRelationsResponse) GetBody() *ListResourceRelationsResponseBody {
	return s.Body
}

func (s *ListResourceRelationsResponse) SetHeaders(v map[string]*string) *ListResourceRelationsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceRelationsResponse) SetStatusCode(v int32) *ListResourceRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceRelationsResponse) SetBody(v *ListResourceRelationsResponseBody) *ListResourceRelationsResponse {
	s.Body = v
	return s
}

func (s *ListResourceRelationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
