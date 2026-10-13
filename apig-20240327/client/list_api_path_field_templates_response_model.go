// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiPathFieldTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApiPathFieldTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApiPathFieldTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListApiPathFieldTemplatesResponseBody) *ListApiPathFieldTemplatesResponse
	GetBody() *ListApiPathFieldTemplatesResponseBody
}

type ListApiPathFieldTemplatesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApiPathFieldTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApiPathFieldTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApiPathFieldTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListApiPathFieldTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApiPathFieldTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApiPathFieldTemplatesResponse) GetBody() *ListApiPathFieldTemplatesResponseBody {
	return s.Body
}

func (s *ListApiPathFieldTemplatesResponse) SetHeaders(v map[string]*string) *ListApiPathFieldTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListApiPathFieldTemplatesResponse) SetStatusCode(v int32) *ListApiPathFieldTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApiPathFieldTemplatesResponse) SetBody(v *ListApiPathFieldTemplatesResponseBody) *ListApiPathFieldTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListApiPathFieldTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
