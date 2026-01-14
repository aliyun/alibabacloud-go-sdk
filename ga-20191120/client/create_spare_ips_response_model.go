// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSpareIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSpareIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSpareIpsResponse
	GetStatusCode() *int32
	SetBody(v *CreateSpareIpsResponseBody) *CreateSpareIpsResponse
	GetBody() *CreateSpareIpsResponseBody
}

type CreateSpareIpsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSpareIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSpareIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSpareIpsResponse) GoString() string {
	return s.String()
}

func (s *CreateSpareIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSpareIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSpareIpsResponse) GetBody() *CreateSpareIpsResponseBody {
	return s.Body
}

func (s *CreateSpareIpsResponse) SetHeaders(v map[string]*string) *CreateSpareIpsResponse {
	s.Headers = v
	return s
}

func (s *CreateSpareIpsResponse) SetStatusCode(v int32) *CreateSpareIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSpareIpsResponse) SetBody(v *CreateSpareIpsResponseBody) *CreateSpareIpsResponse {
	s.Body = v
	return s
}

func (s *CreateSpareIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
