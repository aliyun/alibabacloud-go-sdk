// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationPlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationPlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationPlansResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationPlansResponseBody) *ListOperationPlansResponse
	GetBody() *ListOperationPlansResponseBody
}

type ListOperationPlansResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationPlansResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationPlansResponse) GoString() string {
	return s.String()
}

func (s *ListOperationPlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationPlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationPlansResponse) GetBody() *ListOperationPlansResponseBody {
	return s.Body
}

func (s *ListOperationPlansResponse) SetHeaders(v map[string]*string) *ListOperationPlansResponse {
	s.Headers = v
	return s
}

func (s *ListOperationPlansResponse) SetStatusCode(v int32) *ListOperationPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationPlansResponse) SetBody(v *ListOperationPlansResponseBody) *ListOperationPlansResponse {
	s.Body = v
	return s
}

func (s *ListOperationPlansResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
