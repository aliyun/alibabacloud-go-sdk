// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAdviceServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableAdviceServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableAdviceServiceResponse
	GetStatusCode() *int32
	SetBody(v *DisableAdviceServiceResponseBody) *DisableAdviceServiceResponse
	GetBody() *DisableAdviceServiceResponseBody
}

type DisableAdviceServiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAdviceServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableAdviceServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableAdviceServiceResponse) GoString() string {
	return s.String()
}

func (s *DisableAdviceServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableAdviceServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableAdviceServiceResponse) GetBody() *DisableAdviceServiceResponseBody {
	return s.Body
}

func (s *DisableAdviceServiceResponse) SetHeaders(v map[string]*string) *DisableAdviceServiceResponse {
	s.Headers = v
	return s
}

func (s *DisableAdviceServiceResponse) SetStatusCode(v int32) *DisableAdviceServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAdviceServiceResponse) SetBody(v *DisableAdviceServiceResponseBody) *DisableAdviceServiceResponse {
	s.Body = v
	return s
}

func (s *DisableAdviceServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
