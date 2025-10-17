// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExtensionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExtensionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExtensionsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExtensionsResponseBody) *DeleteExtensionsResponse
	GetBody() *DeleteExtensionsResponseBody
}

type DeleteExtensionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExtensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExtensionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExtensionsResponse) GoString() string {
	return s.String()
}

func (s *DeleteExtensionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExtensionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExtensionsResponse) GetBody() *DeleteExtensionsResponseBody {
	return s.Body
}

func (s *DeleteExtensionsResponse) SetHeaders(v map[string]*string) *DeleteExtensionsResponse {
	s.Headers = v
	return s
}

func (s *DeleteExtensionsResponse) SetStatusCode(v int32) *DeleteExtensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExtensionsResponse) SetBody(v *DeleteExtensionsResponseBody) *DeleteExtensionsResponse {
	s.Body = v
	return s
}

func (s *DeleteExtensionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
