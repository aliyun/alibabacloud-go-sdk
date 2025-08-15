// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeliveryHistoryJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeliveryHistoryJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeliveryHistoryJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListDeliveryHistoryJobsResponseBody) *ListDeliveryHistoryJobsResponse
	GetBody() *ListDeliveryHistoryJobsResponseBody
}

type ListDeliveryHistoryJobsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeliveryHistoryJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeliveryHistoryJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryHistoryJobsResponse) GoString() string {
	return s.String()
}

func (s *ListDeliveryHistoryJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeliveryHistoryJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeliveryHistoryJobsResponse) GetBody() *ListDeliveryHistoryJobsResponseBody {
	return s.Body
}

func (s *ListDeliveryHistoryJobsResponse) SetHeaders(v map[string]*string) *ListDeliveryHistoryJobsResponse {
	s.Headers = v
	return s
}

func (s *ListDeliveryHistoryJobsResponse) SetStatusCode(v int32) *ListDeliveryHistoryJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponse) SetBody(v *ListDeliveryHistoryJobsResponseBody) *ListDeliveryHistoryJobsResponse {
	s.Body = v
	return s
}

func (s *ListDeliveryHistoryJobsResponse) Validate() error {
	return dara.Validate(s)
}
