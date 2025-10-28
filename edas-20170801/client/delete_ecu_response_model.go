// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEcuResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEcuResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEcuResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEcuResponseBody) *DeleteEcuResponse
	GetBody() *DeleteEcuResponseBody
}

type DeleteEcuResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEcuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEcuResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEcuResponse) GoString() string {
	return s.String()
}

func (s *DeleteEcuResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEcuResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEcuResponse) GetBody() *DeleteEcuResponseBody {
	return s.Body
}

func (s *DeleteEcuResponse) SetHeaders(v map[string]*string) *DeleteEcuResponse {
	s.Headers = v
	return s
}

func (s *DeleteEcuResponse) SetStatusCode(v int32) *DeleteEcuResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEcuResponse) SetBody(v *DeleteEcuResponseBody) *DeleteEcuResponse {
	s.Body = v
	return s
}

func (s *DeleteEcuResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
