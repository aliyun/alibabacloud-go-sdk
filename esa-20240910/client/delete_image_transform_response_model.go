// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageTransformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteImageTransformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteImageTransformResponse
	GetStatusCode() *int32
	SetBody(v *DeleteImageTransformResponseBody) *DeleteImageTransformResponse
	GetBody() *DeleteImageTransformResponseBody
}

type DeleteImageTransformResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteImageTransformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteImageTransformResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageTransformResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageTransformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteImageTransformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteImageTransformResponse) GetBody() *DeleteImageTransformResponseBody {
	return s.Body
}

func (s *DeleteImageTransformResponse) SetHeaders(v map[string]*string) *DeleteImageTransformResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageTransformResponse) SetStatusCode(v int32) *DeleteImageTransformResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImageTransformResponse) SetBody(v *DeleteImageTransformResponseBody) *DeleteImageTransformResponse {
	s.Body = v
	return s
}

func (s *DeleteImageTransformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
