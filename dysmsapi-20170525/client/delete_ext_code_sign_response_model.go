// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExtCodeSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExtCodeSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExtCodeSignResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExtCodeSignResponseBody) *DeleteExtCodeSignResponse
	GetBody() *DeleteExtCodeSignResponseBody
}

type DeleteExtCodeSignResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExtCodeSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExtCodeSignResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExtCodeSignResponse) GoString() string {
	return s.String()
}

func (s *DeleteExtCodeSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExtCodeSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExtCodeSignResponse) GetBody() *DeleteExtCodeSignResponseBody {
	return s.Body
}

func (s *DeleteExtCodeSignResponse) SetHeaders(v map[string]*string) *DeleteExtCodeSignResponse {
	s.Headers = v
	return s
}

func (s *DeleteExtCodeSignResponse) SetStatusCode(v int32) *DeleteExtCodeSignResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExtCodeSignResponse) SetBody(v *DeleteExtCodeSignResponseBody) *DeleteExtCodeSignResponse {
	s.Body = v
	return s
}

func (s *DeleteExtCodeSignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
