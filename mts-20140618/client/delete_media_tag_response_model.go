// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMediaTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMediaTagResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMediaTagResponseBody) *DeleteMediaTagResponse
	GetBody() *DeleteMediaTagResponseBody
}

type DeleteMediaTagResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMediaTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMediaTagResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaTagResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediaTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMediaTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMediaTagResponse) GetBody() *DeleteMediaTagResponseBody {
	return s.Body
}

func (s *DeleteMediaTagResponse) SetHeaders(v map[string]*string) *DeleteMediaTagResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediaTagResponse) SetStatusCode(v int32) *DeleteMediaTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMediaTagResponse) SetBody(v *DeleteMediaTagResponseBody) *DeleteMediaTagResponse {
	s.Body = v
	return s
}

func (s *DeleteMediaTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
