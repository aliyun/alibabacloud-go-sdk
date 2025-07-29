// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenAckServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenAckServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenAckServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenAckServiceResponseBody) *OpenAckServiceResponse
	GetBody() *OpenAckServiceResponseBody
}

type OpenAckServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenAckServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenAckServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenAckServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenAckServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenAckServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenAckServiceResponse) GetBody() *OpenAckServiceResponseBody {
	return s.Body
}

func (s *OpenAckServiceResponse) SetHeaders(v map[string]*string) *OpenAckServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenAckServiceResponse) SetStatusCode(v int32) *OpenAckServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenAckServiceResponse) SetBody(v *OpenAckServiceResponseBody) *OpenAckServiceResponse {
	s.Body = v
	return s
}

func (s *OpenAckServiceResponse) Validate() error {
	return dara.Validate(s)
}
