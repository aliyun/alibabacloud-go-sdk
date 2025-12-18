// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemediationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRemediationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRemediationsResponse
	GetStatusCode() *int32
	SetBody(v *ListRemediationsResponseBody) *ListRemediationsResponse
	GetBody() *ListRemediationsResponseBody
}

type ListRemediationsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRemediationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRemediationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRemediationsResponse) GoString() string {
	return s.String()
}

func (s *ListRemediationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRemediationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRemediationsResponse) GetBody() *ListRemediationsResponseBody {
	return s.Body
}

func (s *ListRemediationsResponse) SetHeaders(v map[string]*string) *ListRemediationsResponse {
	s.Headers = v
	return s
}

func (s *ListRemediationsResponse) SetStatusCode(v int32) *ListRemediationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRemediationsResponse) SetBody(v *ListRemediationsResponseBody) *ListRemediationsResponse {
	s.Body = v
	return s
}

func (s *ListRemediationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
