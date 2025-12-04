// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVccResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVccResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVccResponse
	GetStatusCode() *int32
	SetBody(v *CreateVccResponseBody) *CreateVccResponse
	GetBody() *CreateVccResponseBody
}

type CreateVccResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVccResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVccResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVccResponse) GoString() string {
	return s.String()
}

func (s *CreateVccResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVccResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVccResponse) GetBody() *CreateVccResponseBody {
	return s.Body
}

func (s *CreateVccResponse) SetHeaders(v map[string]*string) *CreateVccResponse {
	s.Headers = v
	return s
}

func (s *CreateVccResponse) SetStatusCode(v int32) *CreateVccResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVccResponse) SetBody(v *CreateVccResponseBody) *CreateVccResponse {
	s.Body = v
	return s
}

func (s *CreateVccResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
