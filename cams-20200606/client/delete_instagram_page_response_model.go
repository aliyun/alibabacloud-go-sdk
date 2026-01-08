// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstagramPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstagramPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstagramPageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstagramPageResponseBody) *DeleteInstagramPageResponse
	GetBody() *DeleteInstagramPageResponseBody
}

type DeleteInstagramPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstagramPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstagramPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstagramPageResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstagramPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstagramPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstagramPageResponse) GetBody() *DeleteInstagramPageResponseBody {
	return s.Body
}

func (s *DeleteInstagramPageResponse) SetHeaders(v map[string]*string) *DeleteInstagramPageResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstagramPageResponse) SetStatusCode(v int32) *DeleteInstagramPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstagramPageResponse) SetBody(v *DeleteInstagramPageResponseBody) *DeleteInstagramPageResponse {
	s.Body = v
	return s
}

func (s *DeleteInstagramPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
