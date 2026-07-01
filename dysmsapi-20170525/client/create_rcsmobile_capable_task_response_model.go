// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCSMobileCapableTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRCSMobileCapableTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRCSMobileCapableTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateRCSMobileCapableTaskResponseBody) *CreateRCSMobileCapableTaskResponse
	GetBody() *CreateRCSMobileCapableTaskResponseBody
}

type CreateRCSMobileCapableTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRCSMobileCapableTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRCSMobileCapableTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRCSMobileCapableTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRCSMobileCapableTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRCSMobileCapableTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRCSMobileCapableTaskResponse) GetBody() *CreateRCSMobileCapableTaskResponseBody {
	return s.Body
}

func (s *CreateRCSMobileCapableTaskResponse) SetHeaders(v map[string]*string) *CreateRCSMobileCapableTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRCSMobileCapableTaskResponse) SetStatusCode(v int32) *CreateRCSMobileCapableTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRCSMobileCapableTaskResponse) SetBody(v *CreateRCSMobileCapableTaskResponseBody) *CreateRCSMobileCapableTaskResponse {
	s.Body = v
	return s
}

func (s *CreateRCSMobileCapableTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
