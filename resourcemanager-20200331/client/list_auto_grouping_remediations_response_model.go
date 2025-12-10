// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoGroupingRemediationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAutoGroupingRemediationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAutoGroupingRemediationsResponse
	GetStatusCode() *int32
	SetBody(v *ListAutoGroupingRemediationsResponseBody) *ListAutoGroupingRemediationsResponse
	GetBody() *ListAutoGroupingRemediationsResponseBody
}

type ListAutoGroupingRemediationsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAutoGroupingRemediationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAutoGroupingRemediationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAutoGroupingRemediationsResponse) GoString() string {
	return s.String()
}

func (s *ListAutoGroupingRemediationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAutoGroupingRemediationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAutoGroupingRemediationsResponse) GetBody() *ListAutoGroupingRemediationsResponseBody {
	return s.Body
}

func (s *ListAutoGroupingRemediationsResponse) SetHeaders(v map[string]*string) *ListAutoGroupingRemediationsResponse {
	s.Headers = v
	return s
}

func (s *ListAutoGroupingRemediationsResponse) SetStatusCode(v int32) *ListAutoGroupingRemediationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutoGroupingRemediationsResponse) SetBody(v *ListAutoGroupingRemediationsResponseBody) *ListAutoGroupingRemediationsResponse {
	s.Body = v
	return s
}

func (s *ListAutoGroupingRemediationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
