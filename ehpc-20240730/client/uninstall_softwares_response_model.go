// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallSoftwaresResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallSoftwaresResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallSoftwaresResponse
	GetStatusCode() *int32
	SetBody(v *UninstallSoftwaresResponseBody) *UninstallSoftwaresResponse
	GetBody() *UninstallSoftwaresResponseBody
}

type UninstallSoftwaresResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallSoftwaresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallSoftwaresResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallSoftwaresResponse) GoString() string {
	return s.String()
}

func (s *UninstallSoftwaresResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallSoftwaresResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallSoftwaresResponse) GetBody() *UninstallSoftwaresResponseBody {
	return s.Body
}

func (s *UninstallSoftwaresResponse) SetHeaders(v map[string]*string) *UninstallSoftwaresResponse {
	s.Headers = v
	return s
}

func (s *UninstallSoftwaresResponse) SetStatusCode(v int32) *UninstallSoftwaresResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallSoftwaresResponse) SetBody(v *UninstallSoftwaresResponseBody) *UninstallSoftwaresResponse {
	s.Body = v
	return s
}

func (s *UninstallSoftwaresResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
