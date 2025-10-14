// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteListResponse
	GetStatusCode() *int32
	SetBody(v *DeleteListResponseBody) *DeleteListResponse
	GetBody() *DeleteListResponseBody
}

type DeleteListResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteListResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteListResponse) GoString() string {
	return s.String()
}

func (s *DeleteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteListResponse) GetBody() *DeleteListResponseBody {
	return s.Body
}

func (s *DeleteListResponse) SetHeaders(v map[string]*string) *DeleteListResponse {
	s.Headers = v
	return s
}

func (s *DeleteListResponse) SetStatusCode(v int32) *DeleteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteListResponse) SetBody(v *DeleteListResponseBody) *DeleteListResponse {
	s.Body = v
	return s
}

func (s *DeleteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
