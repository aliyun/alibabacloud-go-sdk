// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSparkAppAttemptsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSparkAppAttemptsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSparkAppAttemptsResponse
	GetStatusCode() *int32
	SetBody(v *ListSparkAppAttemptsResponseBody) *ListSparkAppAttemptsResponse
	GetBody() *ListSparkAppAttemptsResponseBody
}

type ListSparkAppAttemptsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSparkAppAttemptsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSparkAppAttemptsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSparkAppAttemptsResponse) GoString() string {
	return s.String()
}

func (s *ListSparkAppAttemptsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSparkAppAttemptsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSparkAppAttemptsResponse) GetBody() *ListSparkAppAttemptsResponseBody {
	return s.Body
}

func (s *ListSparkAppAttemptsResponse) SetHeaders(v map[string]*string) *ListSparkAppAttemptsResponse {
	s.Headers = v
	return s
}

func (s *ListSparkAppAttemptsResponse) SetStatusCode(v int32) *ListSparkAppAttemptsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSparkAppAttemptsResponse) SetBody(v *ListSparkAppAttemptsResponseBody) *ListSparkAppAttemptsResponse {
	s.Body = v
	return s
}

func (s *ListSparkAppAttemptsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
