// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGdnInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGdnInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGdnInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGdnInstanceResponseBody) *DeleteGdnInstanceResponse
	GetBody() *DeleteGdnInstanceResponseBody
}

type DeleteGdnInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGdnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGdnInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGdnInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGdnInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGdnInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGdnInstanceResponse) GetBody() *DeleteGdnInstanceResponseBody {
	return s.Body
}

func (s *DeleteGdnInstanceResponse) SetHeaders(v map[string]*string) *DeleteGdnInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGdnInstanceResponse) SetStatusCode(v int32) *DeleteGdnInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGdnInstanceResponse) SetBody(v *DeleteGdnInstanceResponseBody) *DeleteGdnInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteGdnInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
