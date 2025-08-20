// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackOperationRisksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStackOperationRisksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStackOperationRisksResponse
	GetStatusCode() *int32
	SetBody(v *ListStackOperationRisksResponseBody) *ListStackOperationRisksResponse
	GetBody() *ListStackOperationRisksResponseBody
}

type ListStackOperationRisksResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStackOperationRisksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStackOperationRisksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStackOperationRisksResponse) GoString() string {
	return s.String()
}

func (s *ListStackOperationRisksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStackOperationRisksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStackOperationRisksResponse) GetBody() *ListStackOperationRisksResponseBody {
	return s.Body
}

func (s *ListStackOperationRisksResponse) SetHeaders(v map[string]*string) *ListStackOperationRisksResponse {
	s.Headers = v
	return s
}

func (s *ListStackOperationRisksResponse) SetStatusCode(v int32) *ListStackOperationRisksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackOperationRisksResponse) SetBody(v *ListStackOperationRisksResponseBody) *ListStackOperationRisksResponse {
	s.Body = v
	return s
}

func (s *ListStackOperationRisksResponse) Validate() error {
	return dara.Validate(s)
}
