// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoScalingActivitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAutoScalingActivitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAutoScalingActivitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListAutoScalingActivitiesResponseBody) *ListAutoScalingActivitiesResponse
	GetBody() *ListAutoScalingActivitiesResponseBody
}

type ListAutoScalingActivitiesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAutoScalingActivitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAutoScalingActivitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingActivitiesResponse) GoString() string {
	return s.String()
}

func (s *ListAutoScalingActivitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAutoScalingActivitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAutoScalingActivitiesResponse) GetBody() *ListAutoScalingActivitiesResponseBody {
	return s.Body
}

func (s *ListAutoScalingActivitiesResponse) SetHeaders(v map[string]*string) *ListAutoScalingActivitiesResponse {
	s.Headers = v
	return s
}

func (s *ListAutoScalingActivitiesResponse) SetStatusCode(v int32) *ListAutoScalingActivitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutoScalingActivitiesResponse) SetBody(v *ListAutoScalingActivitiesResponseBody) *ListAutoScalingActivitiesResponse {
	s.Body = v
	return s
}

func (s *ListAutoScalingActivitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
