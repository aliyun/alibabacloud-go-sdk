// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStoryResponseBody) *DeleteStoryResponse
	GetBody() *DeleteStoryResponseBody
}

type DeleteStoryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteStoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStoryResponse) GetBody() *DeleteStoryResponseBody {
	return s.Body
}

func (s *DeleteStoryResponse) SetHeaders(v map[string]*string) *DeleteStoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteStoryResponse) SetStatusCode(v int32) *DeleteStoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStoryResponse) SetBody(v *DeleteStoryResponseBody) *DeleteStoryResponse {
	s.Body = v
	return s
}

func (s *DeleteStoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
