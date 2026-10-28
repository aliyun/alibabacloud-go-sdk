// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGadOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGadOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGadOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateGadOrderResponseBody) *CreateGadOrderResponse
	GetBody() *CreateGadOrderResponseBody
}

type CreateGadOrderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGadOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGadOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGadOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateGadOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGadOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGadOrderResponse) GetBody() *CreateGadOrderResponseBody {
	return s.Body
}

func (s *CreateGadOrderResponse) SetHeaders(v map[string]*string) *CreateGadOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateGadOrderResponse) SetStatusCode(v int32) *CreateGadOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGadOrderResponse) SetBody(v *CreateGadOrderResponseBody) *CreateGadOrderResponse {
	s.Body = v
	return s
}

func (s *CreateGadOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
