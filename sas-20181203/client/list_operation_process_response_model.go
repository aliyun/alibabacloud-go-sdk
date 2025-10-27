// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationProcessResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationProcessResponseBody) *ListOperationProcessResponse
	GetBody() *ListOperationProcessResponseBody
}

type ListOperationProcessResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationProcessResponse) GoString() string {
	return s.String()
}

func (s *ListOperationProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationProcessResponse) GetBody() *ListOperationProcessResponseBody {
	return s.Body
}

func (s *ListOperationProcessResponse) SetHeaders(v map[string]*string) *ListOperationProcessResponse {
	s.Headers = v
	return s
}

func (s *ListOperationProcessResponse) SetStatusCode(v int32) *ListOperationProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationProcessResponse) SetBody(v *ListOperationProcessResponseBody) *ListOperationProcessResponse {
	s.Body = v
	return s
}

func (s *ListOperationProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
