// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationCheckResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationCheckResponseBody) *ListOperationCheckResponse
	GetBody() *ListOperationCheckResponseBody
}

type ListOperationCheckResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationCheckResponse) GoString() string {
	return s.String()
}

func (s *ListOperationCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationCheckResponse) GetBody() *ListOperationCheckResponseBody {
	return s.Body
}

func (s *ListOperationCheckResponse) SetHeaders(v map[string]*string) *ListOperationCheckResponse {
	s.Headers = v
	return s
}

func (s *ListOperationCheckResponse) SetStatusCode(v int32) *ListOperationCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationCheckResponse) SetBody(v *ListOperationCheckResponseBody) *ListOperationCheckResponse {
	s.Body = v
	return s
}

func (s *ListOperationCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
