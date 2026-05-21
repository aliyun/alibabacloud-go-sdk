// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPostpaidRatePlanInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPostpaidRatePlanInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPostpaidRatePlanInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListPostpaidRatePlanInstancesResponseBody) *ListPostpaidRatePlanInstancesResponse
	GetBody() *ListPostpaidRatePlanInstancesResponseBody
}

type ListPostpaidRatePlanInstancesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPostpaidRatePlanInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPostpaidRatePlanInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPostpaidRatePlanInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListPostpaidRatePlanInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPostpaidRatePlanInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPostpaidRatePlanInstancesResponse) GetBody() *ListPostpaidRatePlanInstancesResponseBody {
	return s.Body
}

func (s *ListPostpaidRatePlanInstancesResponse) SetHeaders(v map[string]*string) *ListPostpaidRatePlanInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponse) SetStatusCode(v int32) *ListPostpaidRatePlanInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponse) SetBody(v *ListPostpaidRatePlanInstancesResponseBody) *ListPostpaidRatePlanInstancesResponse {
	s.Body = v
	return s
}

func (s *ListPostpaidRatePlanInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
