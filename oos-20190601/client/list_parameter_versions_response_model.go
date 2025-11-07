// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParameterVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListParameterVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListParameterVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListParameterVersionsResponseBody) *ListParameterVersionsResponse
	GetBody() *ListParameterVersionsResponseBody
}

type ListParameterVersionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListParameterVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListParameterVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListParameterVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListParameterVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListParameterVersionsResponse) GetBody() *ListParameterVersionsResponseBody {
	return s.Body
}

func (s *ListParameterVersionsResponse) SetHeaders(v map[string]*string) *ListParameterVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListParameterVersionsResponse) SetStatusCode(v int32) *ListParameterVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListParameterVersionsResponse) SetBody(v *ListParameterVersionsResponseBody) *ListParameterVersionsResponse {
	s.Body = v
	return s
}

func (s *ListParameterVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
