// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataCorrectOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataCorrectOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataCorrectOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataCorrectOrderResponseBody) *CreateDataCorrectOrderResponse
	GetBody() *CreateDataCorrectOrderResponseBody
}

type CreateDataCorrectOrderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataCorrectOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataCorrectOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataCorrectOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateDataCorrectOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataCorrectOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataCorrectOrderResponse) GetBody() *CreateDataCorrectOrderResponseBody {
	return s.Body
}

func (s *CreateDataCorrectOrderResponse) SetHeaders(v map[string]*string) *CreateDataCorrectOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateDataCorrectOrderResponse) SetStatusCode(v int32) *CreateDataCorrectOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataCorrectOrderResponse) SetBody(v *CreateDataCorrectOrderResponseBody) *CreateDataCorrectOrderResponse {
	s.Body = v
	return s
}

func (s *CreateDataCorrectOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
