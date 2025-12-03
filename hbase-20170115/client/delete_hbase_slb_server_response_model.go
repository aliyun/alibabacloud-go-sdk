// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHbaseSlbServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHbaseSlbServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHbaseSlbServerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHbaseSlbServerResponseBody) *DeleteHbaseSlbServerResponse
	GetBody() *DeleteHbaseSlbServerResponseBody
}

type DeleteHbaseSlbServerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHbaseSlbServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHbaseSlbServerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHbaseSlbServerResponse) GoString() string {
	return s.String()
}

func (s *DeleteHbaseSlbServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHbaseSlbServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHbaseSlbServerResponse) GetBody() *DeleteHbaseSlbServerResponseBody {
	return s.Body
}

func (s *DeleteHbaseSlbServerResponse) SetHeaders(v map[string]*string) *DeleteHbaseSlbServerResponse {
	s.Headers = v
	return s
}

func (s *DeleteHbaseSlbServerResponse) SetStatusCode(v int32) *DeleteHbaseSlbServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHbaseSlbServerResponse) SetBody(v *DeleteHbaseSlbServerResponseBody) *DeleteHbaseSlbServerResponse {
	s.Body = v
	return s
}

func (s *DeleteHbaseSlbServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
