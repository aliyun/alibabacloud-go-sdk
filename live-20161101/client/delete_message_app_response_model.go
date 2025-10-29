// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMessageAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMessageAppResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMessageAppResponseBody) *DeleteMessageAppResponse
	GetBody() *DeleteMessageAppResponseBody
}

type DeleteMessageAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMessageAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMessageAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteMessageAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMessageAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMessageAppResponse) GetBody() *DeleteMessageAppResponseBody {
	return s.Body
}

func (s *DeleteMessageAppResponse) SetHeaders(v map[string]*string) *DeleteMessageAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteMessageAppResponse) SetStatusCode(v int32) *DeleteMessageAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMessageAppResponse) SetBody(v *DeleteMessageAppResponseBody) *DeleteMessageAppResponse {
	s.Body = v
	return s
}

func (s *DeleteMessageAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
