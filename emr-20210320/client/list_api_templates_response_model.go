// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApiTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApiTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListApiTemplatesResponseBody) *ListApiTemplatesResponse
	GetBody() *ListApiTemplatesResponseBody
}

type ListApiTemplatesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApiTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApiTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApiTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListApiTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApiTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApiTemplatesResponse) GetBody() *ListApiTemplatesResponseBody {
	return s.Body
}

func (s *ListApiTemplatesResponse) SetHeaders(v map[string]*string) *ListApiTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListApiTemplatesResponse) SetStatusCode(v int32) *ListApiTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApiTemplatesResponse) SetBody(v *ListApiTemplatesResponseBody) *ListApiTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListApiTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
