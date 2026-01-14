// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSpareIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSpareIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSpareIpsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSpareIpsResponseBody) *DeleteSpareIpsResponse
	GetBody() *DeleteSpareIpsResponseBody
}

type DeleteSpareIpsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSpareIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSpareIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSpareIpsResponse) GoString() string {
	return s.String()
}

func (s *DeleteSpareIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSpareIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSpareIpsResponse) GetBody() *DeleteSpareIpsResponseBody {
	return s.Body
}

func (s *DeleteSpareIpsResponse) SetHeaders(v map[string]*string) *DeleteSpareIpsResponse {
	s.Headers = v
	return s
}

func (s *DeleteSpareIpsResponse) SetStatusCode(v int32) *DeleteSpareIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSpareIpsResponse) SetBody(v *DeleteSpareIpsResponseBody) *DeleteSpareIpsResponse {
	s.Body = v
	return s
}

func (s *DeleteSpareIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
