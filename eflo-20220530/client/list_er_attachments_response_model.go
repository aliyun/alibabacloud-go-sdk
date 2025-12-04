// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListErAttachmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListErAttachmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListErAttachmentsResponse
	GetStatusCode() *int32
	SetBody(v *ListErAttachmentsResponseBody) *ListErAttachmentsResponse
	GetBody() *ListErAttachmentsResponseBody
}

type ListErAttachmentsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListErAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListErAttachmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListErAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *ListErAttachmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListErAttachmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListErAttachmentsResponse) GetBody() *ListErAttachmentsResponseBody {
	return s.Body
}

func (s *ListErAttachmentsResponse) SetHeaders(v map[string]*string) *ListErAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *ListErAttachmentsResponse) SetStatusCode(v int32) *ListErAttachmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListErAttachmentsResponse) SetBody(v *ListErAttachmentsResponseBody) *ListErAttachmentsResponse {
	s.Body = v
	return s
}

func (s *ListErAttachmentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
