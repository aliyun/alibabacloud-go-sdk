// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancePlansModificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstancePlansModificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstancePlansModificationResponse
	GetStatusCode() *int32
	SetBody(v *ListInstancePlansModificationResponseBody) *ListInstancePlansModificationResponse
	GetBody() *ListInstancePlansModificationResponseBody
}

type ListInstancePlansModificationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancePlansModificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancePlansModificationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstancePlansModificationResponse) GoString() string {
	return s.String()
}

func (s *ListInstancePlansModificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstancePlansModificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstancePlansModificationResponse) GetBody() *ListInstancePlansModificationResponseBody {
	return s.Body
}

func (s *ListInstancePlansModificationResponse) SetHeaders(v map[string]*string) *ListInstancePlansModificationResponse {
	s.Headers = v
	return s
}

func (s *ListInstancePlansModificationResponse) SetStatusCode(v int32) *ListInstancePlansModificationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancePlansModificationResponse) SetBody(v *ListInstancePlansModificationResponseBody) *ListInstancePlansModificationResponse {
	s.Body = v
	return s
}

func (s *ListInstancePlansModificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
