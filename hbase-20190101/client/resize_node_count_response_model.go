// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeNodeCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResizeNodeCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResizeNodeCountResponse
	GetStatusCode() *int32
	SetBody(v *ResizeNodeCountResponseBody) *ResizeNodeCountResponse
	GetBody() *ResizeNodeCountResponseBody
}

type ResizeNodeCountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeNodeCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResizeNodeCountResponse) String() string {
	return dara.Prettify(s)
}

func (s ResizeNodeCountResponse) GoString() string {
	return s.String()
}

func (s *ResizeNodeCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResizeNodeCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResizeNodeCountResponse) GetBody() *ResizeNodeCountResponseBody {
	return s.Body
}

func (s *ResizeNodeCountResponse) SetHeaders(v map[string]*string) *ResizeNodeCountResponse {
	s.Headers = v
	return s
}

func (s *ResizeNodeCountResponse) SetStatusCode(v int32) *ResizeNodeCountResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeNodeCountResponse) SetBody(v *ResizeNodeCountResponseBody) *ResizeNodeCountResponse {
	s.Body = v
	return s
}

func (s *ResizeNodeCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
