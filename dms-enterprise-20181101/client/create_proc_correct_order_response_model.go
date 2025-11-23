// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProcCorrectOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProcCorrectOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProcCorrectOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateProcCorrectOrderResponseBody) *CreateProcCorrectOrderResponse
	GetBody() *CreateProcCorrectOrderResponseBody
}

type CreateProcCorrectOrderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProcCorrectOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProcCorrectOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProcCorrectOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateProcCorrectOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProcCorrectOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProcCorrectOrderResponse) GetBody() *CreateProcCorrectOrderResponseBody {
	return s.Body
}

func (s *CreateProcCorrectOrderResponse) SetHeaders(v map[string]*string) *CreateProcCorrectOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateProcCorrectOrderResponse) SetStatusCode(v int32) *CreateProcCorrectOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProcCorrectOrderResponse) SetBody(v *CreateProcCorrectOrderResponseBody) *CreateProcCorrectOrderResponse {
	s.Body = v
	return s
}

func (s *CreateProcCorrectOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
