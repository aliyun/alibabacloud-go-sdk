// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertAggregationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlertAggregationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlertAggregationsResponse
	GetStatusCode() *int32
	SetBody(v *ListAlertAggregationsResponseBody) *ListAlertAggregationsResponse
	GetBody() *ListAlertAggregationsResponseBody
}

type ListAlertAggregationsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertAggregationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertAggregationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlertAggregationsResponse) GoString() string {
	return s.String()
}

func (s *ListAlertAggregationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertAggregationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlertAggregationsResponse) GetBody() *ListAlertAggregationsResponseBody {
	return s.Body
}

func (s *ListAlertAggregationsResponse) SetHeaders(v map[string]*string) *ListAlertAggregationsResponse {
	s.Headers = v
	return s
}

func (s *ListAlertAggregationsResponse) SetStatusCode(v int32) *ListAlertAggregationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertAggregationsResponse) SetBody(v *ListAlertAggregationsResponseBody) *ListAlertAggregationsResponse {
	s.Body = v
	return s
}

func (s *ListAlertAggregationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
