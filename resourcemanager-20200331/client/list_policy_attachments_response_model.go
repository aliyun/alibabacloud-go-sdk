// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyAttachmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPolicyAttachmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPolicyAttachmentsResponse
	GetStatusCode() *int32
	SetBody(v *ListPolicyAttachmentsResponseBody) *ListPolicyAttachmentsResponse
	GetBody() *ListPolicyAttachmentsResponseBody
}

type ListPolicyAttachmentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicyAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicyAttachmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPolicyAttachmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPolicyAttachmentsResponse) GetBody() *ListPolicyAttachmentsResponseBody {
	return s.Body
}

func (s *ListPolicyAttachmentsResponse) SetHeaders(v map[string]*string) *ListPolicyAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *ListPolicyAttachmentsResponse) SetStatusCode(v int32) *ListPolicyAttachmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicyAttachmentsResponse) SetBody(v *ListPolicyAttachmentsResponseBody) *ListPolicyAttachmentsResponse {
	s.Body = v
	return s
}

func (s *ListPolicyAttachmentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
