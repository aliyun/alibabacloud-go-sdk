// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportObjectSuffixResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSupportObjectSuffixResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSupportObjectSuffixResponse
	GetStatusCode() *int32
	SetBody(v *ListSupportObjectSuffixResponseBody) *ListSupportObjectSuffixResponse
	GetBody() *ListSupportObjectSuffixResponseBody
}

type ListSupportObjectSuffixResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSupportObjectSuffixResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSupportObjectSuffixResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSupportObjectSuffixResponse) GoString() string {
	return s.String()
}

func (s *ListSupportObjectSuffixResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSupportObjectSuffixResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSupportObjectSuffixResponse) GetBody() *ListSupportObjectSuffixResponseBody {
	return s.Body
}

func (s *ListSupportObjectSuffixResponse) SetHeaders(v map[string]*string) *ListSupportObjectSuffixResponse {
	s.Headers = v
	return s
}

func (s *ListSupportObjectSuffixResponse) SetStatusCode(v int32) *ListSupportObjectSuffixResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSupportObjectSuffixResponse) SetBody(v *ListSupportObjectSuffixResponseBody) *ListSupportObjectSuffixResponse {
	s.Body = v
	return s
}

func (s *ListSupportObjectSuffixResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
