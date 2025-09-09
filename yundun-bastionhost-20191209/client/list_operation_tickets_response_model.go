// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationTicketsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationTicketsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationTicketsResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationTicketsResponseBody) *ListOperationTicketsResponse
	GetBody() *ListOperationTicketsResponseBody
}

type ListOperationTicketsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationTicketsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationTicketsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationTicketsResponse) GoString() string {
	return s.String()
}

func (s *ListOperationTicketsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationTicketsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationTicketsResponse) GetBody() *ListOperationTicketsResponseBody {
	return s.Body
}

func (s *ListOperationTicketsResponse) SetHeaders(v map[string]*string) *ListOperationTicketsResponse {
	s.Headers = v
	return s
}

func (s *ListOperationTicketsResponse) SetStatusCode(v int32) *ListOperationTicketsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationTicketsResponse) SetBody(v *ListOperationTicketsResponseBody) *ListOperationTicketsResponse {
	s.Body = v
	return s
}

func (s *ListOperationTicketsResponse) Validate() error {
	return dara.Validate(s)
}
