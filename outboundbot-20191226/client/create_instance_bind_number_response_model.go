// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceBindNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceBindNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceBindNumberResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceBindNumberResponseBody) *CreateInstanceBindNumberResponse
	GetBody() *CreateInstanceBindNumberResponseBody
}

type CreateInstanceBindNumberResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceBindNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceBindNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceBindNumberResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceBindNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceBindNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceBindNumberResponse) GetBody() *CreateInstanceBindNumberResponseBody {
	return s.Body
}

func (s *CreateInstanceBindNumberResponse) SetHeaders(v map[string]*string) *CreateInstanceBindNumberResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceBindNumberResponse) SetStatusCode(v int32) *CreateInstanceBindNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceBindNumberResponse) SetBody(v *CreateInstanceBindNumberResponseBody) *CreateInstanceBindNumberResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceBindNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
