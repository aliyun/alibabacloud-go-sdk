// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceTypesResponseBody) *ListInstanceTypesResponse
	GetBody() *ListInstanceTypesResponseBody
}

type ListInstanceTypesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceTypesResponse) GetBody() *ListInstanceTypesResponseBody {
	return s.Body
}

func (s *ListInstanceTypesResponse) SetHeaders(v map[string]*string) *ListInstanceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceTypesResponse) SetStatusCode(v int32) *ListInstanceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceTypesResponse) SetBody(v *ListInstanceTypesResponseBody) *ListInstanceTypesResponse {
	s.Body = v
	return s
}

func (s *ListInstanceTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
