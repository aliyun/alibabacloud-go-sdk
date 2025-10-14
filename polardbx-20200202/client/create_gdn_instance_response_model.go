// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGdnInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGdnInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGdnInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateGdnInstanceResponseBody) *CreateGdnInstanceResponse
	GetBody() *CreateGdnInstanceResponseBody
}

type CreateGdnInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGdnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGdnInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGdnInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateGdnInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGdnInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGdnInstanceResponse) GetBody() *CreateGdnInstanceResponseBody {
	return s.Body
}

func (s *CreateGdnInstanceResponse) SetHeaders(v map[string]*string) *CreateGdnInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateGdnInstanceResponse) SetStatusCode(v int32) *CreateGdnInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGdnInstanceResponse) SetBody(v *CreateGdnInstanceResponseBody) *CreateGdnInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateGdnInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
