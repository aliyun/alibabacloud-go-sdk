// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccessApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccessApplicationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAccessApplicationResponseBody) *DeleteAccessApplicationResponse
	GetBody() *DeleteAccessApplicationResponseBody
}

type DeleteAccessApplicationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccessApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccessApplicationResponse) GetBody() *DeleteAccessApplicationResponseBody {
	return s.Body
}

func (s *DeleteAccessApplicationResponse) SetHeaders(v map[string]*string) *DeleteAccessApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessApplicationResponse) SetStatusCode(v int32) *DeleteAccessApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessApplicationResponse) SetBody(v *DeleteAccessApplicationResponseBody) *DeleteAccessApplicationResponse {
	s.Body = v
	return s
}

func (s *DeleteAccessApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
