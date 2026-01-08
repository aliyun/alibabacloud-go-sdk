// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessengerPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMessengerPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMessengerPageResponse
	GetStatusCode() *int32
	SetBody(v *CreateMessengerPageResponseBody) *CreateMessengerPageResponse
	GetBody() *CreateMessengerPageResponseBody
}

type CreateMessengerPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMessengerPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMessengerPageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMessengerPageResponse) GoString() string {
	return s.String()
}

func (s *CreateMessengerPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMessengerPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMessengerPageResponse) GetBody() *CreateMessengerPageResponseBody {
	return s.Body
}

func (s *CreateMessengerPageResponse) SetHeaders(v map[string]*string) *CreateMessengerPageResponse {
	s.Headers = v
	return s
}

func (s *CreateMessengerPageResponse) SetStatusCode(v int32) *CreateMessengerPageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMessengerPageResponse) SetBody(v *CreateMessengerPageResponseBody) *CreateMessengerPageResponse {
	s.Body = v
	return s
}

func (s *CreateMessengerPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
