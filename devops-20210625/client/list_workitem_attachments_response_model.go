// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkitemAttachmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkitemAttachmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkitemAttachmentsResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkitemAttachmentsResponseBody) *ListWorkitemAttachmentsResponse
	GetBody() *ListWorkitemAttachmentsResponseBody
}

type ListWorkitemAttachmentsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkitemAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkitemAttachmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkitemAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkitemAttachmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkitemAttachmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkitemAttachmentsResponse) GetBody() *ListWorkitemAttachmentsResponseBody {
	return s.Body
}

func (s *ListWorkitemAttachmentsResponse) SetHeaders(v map[string]*string) *ListWorkitemAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkitemAttachmentsResponse) SetStatusCode(v int32) *ListWorkitemAttachmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkitemAttachmentsResponse) SetBody(v *ListWorkitemAttachmentsResponseBody) *ListWorkitemAttachmentsResponse {
	s.Body = v
	return s
}

func (s *ListWorkitemAttachmentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
