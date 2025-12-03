// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEaisEiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEaisEiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEaisEiResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEaisEiResponseBody) *DeleteEaisEiResponse
	GetBody() *DeleteEaisEiResponseBody
}

type DeleteEaisEiResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEaisEiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEaisEiResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEaisEiResponse) GoString() string {
	return s.String()
}

func (s *DeleteEaisEiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEaisEiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEaisEiResponse) GetBody() *DeleteEaisEiResponseBody {
	return s.Body
}

func (s *DeleteEaisEiResponse) SetHeaders(v map[string]*string) *DeleteEaisEiResponse {
	s.Headers = v
	return s
}

func (s *DeleteEaisEiResponse) SetStatusCode(v int32) *DeleteEaisEiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEaisEiResponse) SetBody(v *DeleteEaisEiResponseBody) *DeleteEaisEiResponse {
	s.Body = v
	return s
}

func (s *DeleteEaisEiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
