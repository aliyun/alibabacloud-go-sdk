// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReceivedFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListReceivedFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListReceivedFileResponse
	GetStatusCode() *int32
	SetBody(v *ListReceivedFileResponseBody) *ListReceivedFileResponse
	GetBody() *ListReceivedFileResponseBody
}

type ListReceivedFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListReceivedFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListReceivedFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ListReceivedFileResponse) GoString() string {
	return s.String()
}

func (s *ListReceivedFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListReceivedFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListReceivedFileResponse) GetBody() *ListReceivedFileResponseBody {
	return s.Body
}

func (s *ListReceivedFileResponse) SetHeaders(v map[string]*string) *ListReceivedFileResponse {
	s.Headers = v
	return s
}

func (s *ListReceivedFileResponse) SetStatusCode(v int32) *ListReceivedFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ListReceivedFileResponse) SetBody(v *ListReceivedFileResponseBody) *ListReceivedFileResponse {
	s.Body = v
	return s
}

func (s *ListReceivedFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
