// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindMessengerPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindMessengerPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindMessengerPageResponse
	GetStatusCode() *int32
	SetBody(v *BindMessengerPageResponseBody) *BindMessengerPageResponse
	GetBody() *BindMessengerPageResponseBody
}

type BindMessengerPageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindMessengerPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindMessengerPageResponse) String() string {
	return dara.Prettify(s)
}

func (s BindMessengerPageResponse) GoString() string {
	return s.String()
}

func (s *BindMessengerPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindMessengerPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindMessengerPageResponse) GetBody() *BindMessengerPageResponseBody {
	return s.Body
}

func (s *BindMessengerPageResponse) SetHeaders(v map[string]*string) *BindMessengerPageResponse {
	s.Headers = v
	return s
}

func (s *BindMessengerPageResponse) SetStatusCode(v int32) *BindMessengerPageResponse {
	s.StatusCode = &v
	return s
}

func (s *BindMessengerPageResponse) SetBody(v *BindMessengerPageResponseBody) *BindMessengerPageResponse {
	s.Body = v
	return s
}

func (s *BindMessengerPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
