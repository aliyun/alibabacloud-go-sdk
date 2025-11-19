// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDynamicImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDynamicImageResponse
	GetStatusCode() *int32
	SetBody(v *ListDynamicImageResponseBody) *ListDynamicImageResponse
	GetBody() *ListDynamicImageResponseBody
}

type ListDynamicImageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDynamicImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDynamicImageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicImageResponse) GoString() string {
	return s.String()
}

func (s *ListDynamicImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDynamicImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDynamicImageResponse) GetBody() *ListDynamicImageResponseBody {
	return s.Body
}

func (s *ListDynamicImageResponse) SetHeaders(v map[string]*string) *ListDynamicImageResponse {
	s.Headers = v
	return s
}

func (s *ListDynamicImageResponse) SetStatusCode(v int32) *ListDynamicImageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDynamicImageResponse) SetBody(v *ListDynamicImageResponseBody) *ListDynamicImageResponse {
	s.Body = v
	return s
}

func (s *ListDynamicImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
