// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomAttributesResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomAttributesResponseBody) *ListCustomAttributesResponse
	GetBody() *ListCustomAttributesResponseBody
}

type ListCustomAttributesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAttributesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomAttributesResponse) GetBody() *ListCustomAttributesResponseBody {
	return s.Body
}

func (s *ListCustomAttributesResponse) SetHeaders(v map[string]*string) *ListCustomAttributesResponse {
	s.Headers = v
	return s
}

func (s *ListCustomAttributesResponse) SetStatusCode(v int32) *ListCustomAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomAttributesResponse) SetBody(v *ListCustomAttributesResponseBody) *ListCustomAttributesResponse {
	s.Body = v
	return s
}

func (s *ListCustomAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
