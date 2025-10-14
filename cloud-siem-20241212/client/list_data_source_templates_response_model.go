// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataSourceTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataSourceTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataSourceTemplatesResponseBody) *ListDataSourceTemplatesResponse
	GetBody() *ListDataSourceTemplatesResponseBody
}

type ListDataSourceTemplatesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataSourceTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataSourceTemplatesResponse) GetBody() *ListDataSourceTemplatesResponseBody {
	return s.Body
}

func (s *ListDataSourceTemplatesResponse) SetHeaders(v map[string]*string) *ListDataSourceTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceTemplatesResponse) SetStatusCode(v int32) *ListDataSourceTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceTemplatesResponse) SetBody(v *ListDataSourceTemplatesResponseBody) *ListDataSourceTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListDataSourceTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
