// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TerminateInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TerminateInstanceResponse
	GetStatusCode() *int32
	SetBody(v *TerminateInstanceResponseBody) *TerminateInstanceResponse
	GetBody() *TerminateInstanceResponseBody
}

type TerminateInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminateInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s TerminateInstanceResponse) GoString() string {
	return s.String()
}

func (s *TerminateInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TerminateInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TerminateInstanceResponse) GetBody() *TerminateInstanceResponseBody {
	return s.Body
}

func (s *TerminateInstanceResponse) SetHeaders(v map[string]*string) *TerminateInstanceResponse {
	s.Headers = v
	return s
}

func (s *TerminateInstanceResponse) SetStatusCode(v int32) *TerminateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminateInstanceResponse) SetBody(v *TerminateInstanceResponseBody) *TerminateInstanceResponse {
	s.Body = v
	return s
}

func (s *TerminateInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
