// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePortalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePortalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePortalResponse
	GetStatusCode() *int32
	SetBody(v *DeletePortalResponseBody) *DeletePortalResponse
	GetBody() *DeletePortalResponseBody
}

type DeletePortalResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePortalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePortalResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePortalResponse) GoString() string {
	return s.String()
}

func (s *DeletePortalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePortalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePortalResponse) GetBody() *DeletePortalResponseBody {
	return s.Body
}

func (s *DeletePortalResponse) SetHeaders(v map[string]*string) *DeletePortalResponse {
	s.Headers = v
	return s
}

func (s *DeletePortalResponse) SetStatusCode(v int32) *DeletePortalResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePortalResponse) SetBody(v *DeletePortalResponseBody) *DeletePortalResponse {
	s.Body = v
	return s
}

func (s *DeletePortalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
