// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomImageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomImageResponseBody) *DeleteCustomImageResponse
	GetBody() *DeleteCustomImageResponseBody
}

type DeleteCustomImageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomImageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomImageResponse) GetBody() *DeleteCustomImageResponseBody {
	return s.Body
}

func (s *DeleteCustomImageResponse) SetHeaders(v map[string]*string) *DeleteCustomImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomImageResponse) SetStatusCode(v int32) *DeleteCustomImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomImageResponse) SetBody(v *DeleteCustomImageResponseBody) *DeleteCustomImageResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
