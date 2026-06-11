// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDynamicTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDynamicTagResponse
	GetStatusCode() *int32
	SetBody(v *ListDynamicTagResponseBody) *ListDynamicTagResponse
	GetBody() *ListDynamicTagResponseBody
}

type ListDynamicTagResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDynamicTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDynamicTagResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicTagResponse) GoString() string {
	return s.String()
}

func (s *ListDynamicTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDynamicTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDynamicTagResponse) GetBody() *ListDynamicTagResponseBody {
	return s.Body
}

func (s *ListDynamicTagResponse) SetHeaders(v map[string]*string) *ListDynamicTagResponse {
	s.Headers = v
	return s
}

func (s *ListDynamicTagResponse) SetStatusCode(v int32) *ListDynamicTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDynamicTagResponse) SetBody(v *ListDynamicTagResponseBody) *ListDynamicTagResponse {
	s.Body = v
	return s
}

func (s *ListDynamicTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
