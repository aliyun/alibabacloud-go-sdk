// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmNoConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmNoConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmNoConnectionResponse
	GetStatusCode() *int32
	SetBody(v *ConfirmNoConnectionResponseBody) *ConfirmNoConnectionResponse
	GetBody() *ConfirmNoConnectionResponseBody
}

type ConfirmNoConnectionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmNoConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmNoConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmNoConnectionResponse) GoString() string {
	return s.String()
}

func (s *ConfirmNoConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmNoConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmNoConnectionResponse) GetBody() *ConfirmNoConnectionResponseBody {
	return s.Body
}

func (s *ConfirmNoConnectionResponse) SetHeaders(v map[string]*string) *ConfirmNoConnectionResponse {
	s.Headers = v
	return s
}

func (s *ConfirmNoConnectionResponse) SetStatusCode(v int32) *ConfirmNoConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmNoConnectionResponse) SetBody(v *ConfirmNoConnectionResponseBody) *ConfirmNoConnectionResponse {
	s.Body = v
	return s
}

func (s *ConfirmNoConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
