// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTotalSensitiveInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTotalSensitiveInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTotalSensitiveInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListTotalSensitiveInfoResponseBody) *ListTotalSensitiveInfoResponse
	GetBody() *ListTotalSensitiveInfoResponseBody
}

type ListTotalSensitiveInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTotalSensitiveInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTotalSensitiveInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTotalSensitiveInfoResponse) GoString() string {
	return s.String()
}

func (s *ListTotalSensitiveInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTotalSensitiveInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTotalSensitiveInfoResponse) GetBody() *ListTotalSensitiveInfoResponseBody {
	return s.Body
}

func (s *ListTotalSensitiveInfoResponse) SetHeaders(v map[string]*string) *ListTotalSensitiveInfoResponse {
	s.Headers = v
	return s
}

func (s *ListTotalSensitiveInfoResponse) SetStatusCode(v int32) *ListTotalSensitiveInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTotalSensitiveInfoResponse) SetBody(v *ListTotalSensitiveInfoResponseBody) *ListTotalSensitiveInfoResponse {
	s.Body = v
	return s
}

func (s *ListTotalSensitiveInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
