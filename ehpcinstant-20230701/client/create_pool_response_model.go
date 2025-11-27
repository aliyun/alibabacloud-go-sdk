// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePoolResponse
	GetStatusCode() *int32
	SetBody(v *CreatePoolResponseBody) *CreatePoolResponse
	GetBody() *CreatePoolResponseBody
}

type CreatePoolResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePoolResponse) GoString() string {
	return s.String()
}

func (s *CreatePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePoolResponse) GetBody() *CreatePoolResponseBody {
	return s.Body
}

func (s *CreatePoolResponse) SetHeaders(v map[string]*string) *CreatePoolResponse {
	s.Headers = v
	return s
}

func (s *CreatePoolResponse) SetStatusCode(v int32) *CreatePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePoolResponse) SetBody(v *CreatePoolResponseBody) *CreatePoolResponse {
	s.Body = v
	return s
}

func (s *CreatePoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
