// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWehookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWehookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWehookResponse
	GetStatusCode() *int32
	SetBody(v *CreateWehookResponseBody) *CreateWehookResponse
	GetBody() *CreateWehookResponseBody
}

type CreateWehookResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWehookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWehookResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWehookResponse) GoString() string {
	return s.String()
}

func (s *CreateWehookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWehookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWehookResponse) GetBody() *CreateWehookResponseBody {
	return s.Body
}

func (s *CreateWehookResponse) SetHeaders(v map[string]*string) *CreateWehookResponse {
	s.Headers = v
	return s
}

func (s *CreateWehookResponse) SetStatusCode(v int32) *CreateWehookResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWehookResponse) SetBody(v *CreateWehookResponseBody) *CreateWehookResponse {
	s.Body = v
	return s
}

func (s *CreateWehookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
