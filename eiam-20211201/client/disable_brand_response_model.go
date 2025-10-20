// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableBrandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableBrandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableBrandResponse
	GetStatusCode() *int32
	SetBody(v *DisableBrandResponseBody) *DisableBrandResponse
	GetBody() *DisableBrandResponseBody
}

type DisableBrandResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableBrandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableBrandResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableBrandResponse) GoString() string {
	return s.String()
}

func (s *DisableBrandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableBrandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableBrandResponse) GetBody() *DisableBrandResponseBody {
	return s.Body
}

func (s *DisableBrandResponse) SetHeaders(v map[string]*string) *DisableBrandResponse {
	s.Headers = v
	return s
}

func (s *DisableBrandResponse) SetStatusCode(v int32) *DisableBrandResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableBrandResponse) SetBody(v *DisableBrandResponseBody) *DisableBrandResponse {
	s.Body = v
	return s
}

func (s *DisableBrandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
