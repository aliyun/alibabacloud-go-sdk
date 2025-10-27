// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerCriteriaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContainerCriteriaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContainerCriteriaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContainerCriteriaResponseBody) *DescribeContainerCriteriaResponse
	GetBody() *DescribeContainerCriteriaResponseBody
}

type DescribeContainerCriteriaResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerCriteriaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContainerCriteriaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerCriteriaResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerCriteriaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContainerCriteriaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContainerCriteriaResponse) GetBody() *DescribeContainerCriteriaResponseBody {
	return s.Body
}

func (s *DescribeContainerCriteriaResponse) SetHeaders(v map[string]*string) *DescribeContainerCriteriaResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerCriteriaResponse) SetStatusCode(v int32) *DescribeContainerCriteriaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerCriteriaResponse) SetBody(v *DescribeContainerCriteriaResponseBody) *DescribeContainerCriteriaResponse {
	s.Body = v
	return s
}

func (s *DescribeContainerCriteriaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
