// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataSourceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataSourceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ListDataSourceAttributeResponseBody) *ListDataSourceAttributeResponse
	GetBody() *ListDataSourceAttributeResponseBody
}

type ListDataSourceAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataSourceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataSourceAttributeResponse) GetBody() *ListDataSourceAttributeResponseBody {
	return s.Body
}

func (s *ListDataSourceAttributeResponse) SetHeaders(v map[string]*string) *ListDataSourceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceAttributeResponse) SetStatusCode(v int32) *ListDataSourceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceAttributeResponse) SetBody(v *ListDataSourceAttributeResponseBody) *ListDataSourceAttributeResponse {
	s.Body = v
	return s
}

func (s *ListDataSourceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
