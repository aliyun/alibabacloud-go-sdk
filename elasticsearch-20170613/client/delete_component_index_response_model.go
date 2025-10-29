// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComponentIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteComponentIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteComponentIndexResponse
	GetStatusCode() *int32
	SetBody(v *DeleteComponentIndexResponseBody) *DeleteComponentIndexResponse
	GetBody() *DeleteComponentIndexResponseBody
}

type DeleteComponentIndexResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteComponentIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteComponentIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteComponentIndexResponse) GoString() string {
	return s.String()
}

func (s *DeleteComponentIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteComponentIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteComponentIndexResponse) GetBody() *DeleteComponentIndexResponseBody {
	return s.Body
}

func (s *DeleteComponentIndexResponse) SetHeaders(v map[string]*string) *DeleteComponentIndexResponse {
	s.Headers = v
	return s
}

func (s *DeleteComponentIndexResponse) SetStatusCode(v int32) *DeleteComponentIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteComponentIndexResponse) SetBody(v *DeleteComponentIndexResponseBody) *DeleteComponentIndexResponse {
	s.Body = v
	return s
}

func (s *DeleteComponentIndexResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
