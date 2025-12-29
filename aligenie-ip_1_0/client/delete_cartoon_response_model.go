// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCartoonResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCartoonResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCartoonResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCartoonResponseBody) *DeleteCartoonResponse
	GetBody() *DeleteCartoonResponseBody
}

type DeleteCartoonResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCartoonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCartoonResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCartoonResponse) GoString() string {
	return s.String()
}

func (s *DeleteCartoonResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCartoonResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCartoonResponse) GetBody() *DeleteCartoonResponseBody {
	return s.Body
}

func (s *DeleteCartoonResponse) SetHeaders(v map[string]*string) *DeleteCartoonResponse {
	s.Headers = v
	return s
}

func (s *DeleteCartoonResponse) SetStatusCode(v int32) *DeleteCartoonResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCartoonResponse) SetBody(v *DeleteCartoonResponseBody) *DeleteCartoonResponse {
	s.Body = v
	return s
}

func (s *DeleteCartoonResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
