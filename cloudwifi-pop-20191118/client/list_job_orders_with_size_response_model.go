// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobOrdersWithSizeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobOrdersWithSizeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobOrdersWithSizeResponse
	GetStatusCode() *int32
	SetBody(v *ListJobOrdersWithSizeResponseBody) *ListJobOrdersWithSizeResponse
	GetBody() *ListJobOrdersWithSizeResponseBody
}

type ListJobOrdersWithSizeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobOrdersWithSizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobOrdersWithSizeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobOrdersWithSizeResponse) GoString() string {
	return s.String()
}

func (s *ListJobOrdersWithSizeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobOrdersWithSizeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobOrdersWithSizeResponse) GetBody() *ListJobOrdersWithSizeResponseBody {
	return s.Body
}

func (s *ListJobOrdersWithSizeResponse) SetHeaders(v map[string]*string) *ListJobOrdersWithSizeResponse {
	s.Headers = v
	return s
}

func (s *ListJobOrdersWithSizeResponse) SetStatusCode(v int32) *ListJobOrdersWithSizeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobOrdersWithSizeResponse) SetBody(v *ListJobOrdersWithSizeResponseBody) *ListJobOrdersWithSizeResponse {
	s.Body = v
	return s
}

func (s *ListJobOrdersWithSizeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
