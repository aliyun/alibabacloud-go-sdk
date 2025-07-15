// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonTicketFieldsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCommonTicketFieldsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCommonTicketFieldsResponse
	GetStatusCode() *int32
	SetBody(v *ListCommonTicketFieldsResponseBody) *ListCommonTicketFieldsResponse
	GetBody() *ListCommonTicketFieldsResponseBody
}

type ListCommonTicketFieldsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCommonTicketFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCommonTicketFieldsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCommonTicketFieldsResponse) GoString() string {
	return s.String()
}

func (s *ListCommonTicketFieldsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCommonTicketFieldsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCommonTicketFieldsResponse) GetBody() *ListCommonTicketFieldsResponseBody {
	return s.Body
}

func (s *ListCommonTicketFieldsResponse) SetHeaders(v map[string]*string) *ListCommonTicketFieldsResponse {
	s.Headers = v
	return s
}

func (s *ListCommonTicketFieldsResponse) SetStatusCode(v int32) *ListCommonTicketFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommonTicketFieldsResponse) SetBody(v *ListCommonTicketFieldsResponseBody) *ListCommonTicketFieldsResponse {
	s.Body = v
	return s
}

func (s *ListCommonTicketFieldsResponse) Validate() error {
	return dara.Validate(s)
}
