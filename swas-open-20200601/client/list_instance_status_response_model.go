// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceStatusResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceStatusResponseBody) *ListInstanceStatusResponse
	GetBody() *ListInstanceStatusResponseBody
}

type ListInstanceStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceStatusResponse) GetBody() *ListInstanceStatusResponseBody {
	return s.Body
}

func (s *ListInstanceStatusResponse) SetHeaders(v map[string]*string) *ListInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceStatusResponse) SetStatusCode(v int32) *ListInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceStatusResponse) SetBody(v *ListInstanceStatusResponseBody) *ListInstanceStatusResponse {
	s.Body = v
	return s
}

func (s *ListInstanceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
