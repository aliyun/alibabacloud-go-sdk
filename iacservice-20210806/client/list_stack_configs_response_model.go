// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStackConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStackConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListStackConfigsResponseBody) *ListStackConfigsResponse
	GetBody() *ListStackConfigsResponseBody
}

type ListStackConfigsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStackConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStackConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStackConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListStackConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStackConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStackConfigsResponse) GetBody() *ListStackConfigsResponseBody {
	return s.Body
}

func (s *ListStackConfigsResponse) SetHeaders(v map[string]*string) *ListStackConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListStackConfigsResponse) SetStatusCode(v int32) *ListStackConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackConfigsResponse) SetBody(v *ListStackConfigsResponseBody) *ListStackConfigsResponse {
	s.Body = v
	return s
}

func (s *ListStackConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
