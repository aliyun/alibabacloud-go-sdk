// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationApiInvokeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableApplicationApiInvokeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableApplicationApiInvokeResponse
	GetStatusCode() *int32
	SetBody(v *DisableApplicationApiInvokeResponseBody) *DisableApplicationApiInvokeResponse
	GetBody() *DisableApplicationApiInvokeResponseBody
}

type DisableApplicationApiInvokeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableApplicationApiInvokeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableApplicationApiInvokeResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationApiInvokeResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationApiInvokeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableApplicationApiInvokeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableApplicationApiInvokeResponse) GetBody() *DisableApplicationApiInvokeResponseBody {
	return s.Body
}

func (s *DisableApplicationApiInvokeResponse) SetHeaders(v map[string]*string) *DisableApplicationApiInvokeResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationApiInvokeResponse) SetStatusCode(v int32) *DisableApplicationApiInvokeResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationApiInvokeResponse) SetBody(v *DisableApplicationApiInvokeResponseBody) *DisableApplicationApiInvokeResponse {
	s.Body = v
	return s
}

func (s *DisableApplicationApiInvokeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
