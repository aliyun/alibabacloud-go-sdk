// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskWorkforceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskWorkforceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskWorkforceResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskWorkforceResponseBody) *GetTaskWorkforceResponse
	GetBody() *GetTaskWorkforceResponseBody
}

type GetTaskWorkforceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskWorkforceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskWorkforceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskWorkforceResponse) GoString() string {
	return s.String()
}

func (s *GetTaskWorkforceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskWorkforceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskWorkforceResponse) GetBody() *GetTaskWorkforceResponseBody {
	return s.Body
}

func (s *GetTaskWorkforceResponse) SetHeaders(v map[string]*string) *GetTaskWorkforceResponse {
	s.Headers = v
	return s
}

func (s *GetTaskWorkforceResponse) SetStatusCode(v int32) *GetTaskWorkforceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskWorkforceResponse) SetBody(v *GetTaskWorkforceResponseBody) *GetTaskWorkforceResponse {
	s.Body = v
	return s
}

func (s *GetTaskWorkforceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
