// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceAdbAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceAdbAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceAdbAttributesResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceAdbAttributesResponseBody) *ListInstanceAdbAttributesResponse
	GetBody() *ListInstanceAdbAttributesResponseBody
}

type ListInstanceAdbAttributesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceAdbAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceAdbAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAdbAttributesResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceAdbAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceAdbAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceAdbAttributesResponse) GetBody() *ListInstanceAdbAttributesResponseBody {
	return s.Body
}

func (s *ListInstanceAdbAttributesResponse) SetHeaders(v map[string]*string) *ListInstanceAdbAttributesResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceAdbAttributesResponse) SetStatusCode(v int32) *ListInstanceAdbAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceAdbAttributesResponse) SetBody(v *ListInstanceAdbAttributesResponseBody) *ListInstanceAdbAttributesResponse {
	s.Body = v
	return s
}

func (s *ListInstanceAdbAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
