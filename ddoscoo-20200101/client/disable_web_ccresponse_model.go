// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableWebCCResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableWebCCResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableWebCCResponse
	GetStatusCode() *int32
	SetBody(v *DisableWebCCResponseBody) *DisableWebCCResponse
	GetBody() *DisableWebCCResponseBody
}

type DisableWebCCResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableWebCCResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableWebCCResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableWebCCResponse) GoString() string {
	return s.String()
}

func (s *DisableWebCCResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableWebCCResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableWebCCResponse) GetBody() *DisableWebCCResponseBody {
	return s.Body
}

func (s *DisableWebCCResponse) SetHeaders(v map[string]*string) *DisableWebCCResponse {
	s.Headers = v
	return s
}

func (s *DisableWebCCResponse) SetStatusCode(v int32) *DisableWebCCResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableWebCCResponse) SetBody(v *DisableWebCCResponseBody) *DisableWebCCResponse {
	s.Body = v
	return s
}

func (s *DisableWebCCResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
