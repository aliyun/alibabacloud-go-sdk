// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeylessServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKeylessServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKeylessServerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKeylessServerResponseBody) *DeleteKeylessServerResponse
	GetBody() *DeleteKeylessServerResponseBody
}

type DeleteKeylessServerResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKeylessServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKeylessServerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeylessServerResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeylessServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKeylessServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKeylessServerResponse) GetBody() *DeleteKeylessServerResponseBody {
	return s.Body
}

func (s *DeleteKeylessServerResponse) SetHeaders(v map[string]*string) *DeleteKeylessServerResponse {
	s.Headers = v
	return s
}

func (s *DeleteKeylessServerResponse) SetStatusCode(v int32) *DeleteKeylessServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKeylessServerResponse) SetBody(v *DeleteKeylessServerResponseBody) *DeleteKeylessServerResponse {
	s.Body = v
	return s
}

func (s *DeleteKeylessServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
