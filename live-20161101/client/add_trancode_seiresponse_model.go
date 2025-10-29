// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTrancodeSEIResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTrancodeSEIResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTrancodeSEIResponse
	GetStatusCode() *int32
	SetBody(v *AddTrancodeSEIResponseBody) *AddTrancodeSEIResponse
	GetBody() *AddTrancodeSEIResponseBody
}

type AddTrancodeSEIResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTrancodeSEIResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTrancodeSEIResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTrancodeSEIResponse) GoString() string {
	return s.String()
}

func (s *AddTrancodeSEIResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTrancodeSEIResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTrancodeSEIResponse) GetBody() *AddTrancodeSEIResponseBody {
	return s.Body
}

func (s *AddTrancodeSEIResponse) SetHeaders(v map[string]*string) *AddTrancodeSEIResponse {
	s.Headers = v
	return s
}

func (s *AddTrancodeSEIResponse) SetStatusCode(v int32) *AddTrancodeSEIResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTrancodeSEIResponse) SetBody(v *AddTrancodeSEIResponseBody) *AddTrancodeSEIResponse {
	s.Body = v
	return s
}

func (s *AddTrancodeSEIResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
