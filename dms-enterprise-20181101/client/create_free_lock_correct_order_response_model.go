// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFreeLockCorrectOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFreeLockCorrectOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFreeLockCorrectOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateFreeLockCorrectOrderResponseBody) *CreateFreeLockCorrectOrderResponse
	GetBody() *CreateFreeLockCorrectOrderResponseBody
}

type CreateFreeLockCorrectOrderResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFreeLockCorrectOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFreeLockCorrectOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFreeLockCorrectOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateFreeLockCorrectOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFreeLockCorrectOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFreeLockCorrectOrderResponse) GetBody() *CreateFreeLockCorrectOrderResponseBody {
	return s.Body
}

func (s *CreateFreeLockCorrectOrderResponse) SetHeaders(v map[string]*string) *CreateFreeLockCorrectOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateFreeLockCorrectOrderResponse) SetStatusCode(v int32) *CreateFreeLockCorrectOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFreeLockCorrectOrderResponse) SetBody(v *CreateFreeLockCorrectOrderResponseBody) *CreateFreeLockCorrectOrderResponse {
	s.Body = v
	return s
}

func (s *CreateFreeLockCorrectOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
