// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackGroupOperationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStackGroupOperationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStackGroupOperationsResponse
	GetStatusCode() *int32
	SetBody(v *ListStackGroupOperationsResponseBody) *ListStackGroupOperationsResponse
	GetBody() *ListStackGroupOperationsResponseBody
}

type ListStackGroupOperationsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStackGroupOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStackGroupOperationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStackGroupOperationsResponse) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStackGroupOperationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStackGroupOperationsResponse) GetBody() *ListStackGroupOperationsResponseBody {
	return s.Body
}

func (s *ListStackGroupOperationsResponse) SetHeaders(v map[string]*string) *ListStackGroupOperationsResponse {
	s.Headers = v
	return s
}

func (s *ListStackGroupOperationsResponse) SetStatusCode(v int32) *ListStackGroupOperationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackGroupOperationsResponse) SetBody(v *ListStackGroupOperationsResponseBody) *ListStackGroupOperationsResponse {
	s.Body = v
	return s
}

func (s *ListStackGroupOperationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
