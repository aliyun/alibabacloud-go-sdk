// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMemoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMemoriesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMemoriesResponseBody) *DeleteMemoriesResponse
	GetBody() *DeleteMemoriesResponseBody
}

type DeleteMemoriesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMemoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMemoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoriesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMemoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMemoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMemoriesResponse) GetBody() *DeleteMemoriesResponseBody {
	return s.Body
}

func (s *DeleteMemoriesResponse) SetHeaders(v map[string]*string) *DeleteMemoriesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMemoriesResponse) SetStatusCode(v int32) *DeleteMemoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMemoriesResponse) SetBody(v *DeleteMemoriesResponseBody) *DeleteMemoriesResponse {
	s.Body = v
	return s
}

func (s *DeleteMemoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
