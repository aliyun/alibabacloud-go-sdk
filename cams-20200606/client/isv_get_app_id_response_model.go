// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsvGetAppIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IsvGetAppIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IsvGetAppIdResponse
	GetStatusCode() *int32
	SetBody(v *IsvGetAppIdResponseBody) *IsvGetAppIdResponse
	GetBody() *IsvGetAppIdResponseBody
}

type IsvGetAppIdResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IsvGetAppIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IsvGetAppIdResponse) String() string {
	return dara.Prettify(s)
}

func (s IsvGetAppIdResponse) GoString() string {
	return s.String()
}

func (s *IsvGetAppIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IsvGetAppIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IsvGetAppIdResponse) GetBody() *IsvGetAppIdResponseBody {
	return s.Body
}

func (s *IsvGetAppIdResponse) SetHeaders(v map[string]*string) *IsvGetAppIdResponse {
	s.Headers = v
	return s
}

func (s *IsvGetAppIdResponse) SetStatusCode(v int32) *IsvGetAppIdResponse {
	s.StatusCode = &v
	return s
}

func (s *IsvGetAppIdResponse) SetBody(v *IsvGetAppIdResponseBody) *IsvGetAppIdResponse {
	s.Body = v
	return s
}

func (s *IsvGetAppIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
