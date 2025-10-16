// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeInstanceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodeInstanceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodeInstanceTypeResponse
	GetStatusCode() *int32
	SetBody(v *ListNodeInstanceTypeResponseBody) *ListNodeInstanceTypeResponse
	GetBody() *ListNodeInstanceTypeResponseBody
}

type ListNodeInstanceTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeInstanceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodeInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *ListNodeInstanceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodeInstanceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodeInstanceTypeResponse) GetBody() *ListNodeInstanceTypeResponseBody {
	return s.Body
}

func (s *ListNodeInstanceTypeResponse) SetHeaders(v map[string]*string) *ListNodeInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *ListNodeInstanceTypeResponse) SetStatusCode(v int32) *ListNodeInstanceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeInstanceTypeResponse) SetBody(v *ListNodeInstanceTypeResponseBody) *ListNodeInstanceTypeResponse {
	s.Body = v
	return s
}

func (s *ListNodeInstanceTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
