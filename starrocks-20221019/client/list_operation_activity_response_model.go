// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationActivityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationActivityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationActivityResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationActivityResponseBody) *ListOperationActivityResponse
	GetBody() *ListOperationActivityResponseBody
}

type ListOperationActivityResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationActivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationActivityResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationActivityResponse) GoString() string {
	return s.String()
}

func (s *ListOperationActivityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationActivityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationActivityResponse) GetBody() *ListOperationActivityResponseBody {
	return s.Body
}

func (s *ListOperationActivityResponse) SetHeaders(v map[string]*string) *ListOperationActivityResponse {
	s.Headers = v
	return s
}

func (s *ListOperationActivityResponse) SetStatusCode(v int32) *ListOperationActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationActivityResponse) SetBody(v *ListOperationActivityResponseBody) *ListOperationActivityResponse {
	s.Body = v
	return s
}

func (s *ListOperationActivityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
