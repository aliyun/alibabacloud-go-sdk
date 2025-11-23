// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEffectiveOrdersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEffectiveOrdersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEffectiveOrdersResponse
	GetStatusCode() *int32
	SetBody(v *ListEffectiveOrdersResponseBody) *ListEffectiveOrdersResponse
	GetBody() *ListEffectiveOrdersResponseBody
}

type ListEffectiveOrdersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEffectiveOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEffectiveOrdersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEffectiveOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListEffectiveOrdersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEffectiveOrdersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEffectiveOrdersResponse) GetBody() *ListEffectiveOrdersResponseBody {
	return s.Body
}

func (s *ListEffectiveOrdersResponse) SetHeaders(v map[string]*string) *ListEffectiveOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListEffectiveOrdersResponse) SetStatusCode(v int32) *ListEffectiveOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEffectiveOrdersResponse) SetBody(v *ListEffectiveOrdersResponseBody) *ListEffectiveOrdersResponse {
	s.Body = v
	return s
}

func (s *ListEffectiveOrdersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
