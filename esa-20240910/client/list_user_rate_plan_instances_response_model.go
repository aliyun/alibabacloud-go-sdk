// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserRatePlanInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserRatePlanInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserRatePlanInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserRatePlanInstancesResponseBody) *ListUserRatePlanInstancesResponse
	GetBody() *ListUserRatePlanInstancesResponseBody
}

type ListUserRatePlanInstancesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserRatePlanInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserRatePlanInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserRatePlanInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListUserRatePlanInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserRatePlanInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserRatePlanInstancesResponse) GetBody() *ListUserRatePlanInstancesResponseBody {
	return s.Body
}

func (s *ListUserRatePlanInstancesResponse) SetHeaders(v map[string]*string) *ListUserRatePlanInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListUserRatePlanInstancesResponse) SetStatusCode(v int32) *ListUserRatePlanInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserRatePlanInstancesResponse) SetBody(v *ListUserRatePlanInstancesResponseBody) *ListUserRatePlanInstancesResponse {
	s.Body = v
	return s
}

func (s *ListUserRatePlanInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
