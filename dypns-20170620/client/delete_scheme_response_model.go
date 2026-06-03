// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSchemeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSchemeResponseBody) *DeleteSchemeResponse
	GetBody() *DeleteSchemeResponseBody
}

type DeleteSchemeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchemeResponse) GoString() string {
	return s.String()
}

func (s *DeleteSchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSchemeResponse) GetBody() *DeleteSchemeResponseBody {
	return s.Body
}

func (s *DeleteSchemeResponse) SetHeaders(v map[string]*string) *DeleteSchemeResponse {
	s.Headers = v
	return s
}

func (s *DeleteSchemeResponse) SetStatusCode(v int32) *DeleteSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSchemeResponse) SetBody(v *DeleteSchemeResponseBody) *DeleteSchemeResponse {
	s.Body = v
	return s
}

func (s *DeleteSchemeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
