// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetaCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMetaCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMetaCategoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMetaCategoryResponseBody) *DeleteMetaCategoryResponse
	GetBody() *DeleteMetaCategoryResponseBody
}

type DeleteMetaCategoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMetaCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMetaCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetaCategoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetaCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMetaCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMetaCategoryResponse) GetBody() *DeleteMetaCategoryResponseBody {
	return s.Body
}

func (s *DeleteMetaCategoryResponse) SetHeaders(v map[string]*string) *DeleteMetaCategoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetaCategoryResponse) SetStatusCode(v int32) *DeleteMetaCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMetaCategoryResponse) SetBody(v *DeleteMetaCategoryResponseBody) *DeleteMetaCategoryResponse {
	s.Body = v
	return s
}

func (s *DeleteMetaCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
