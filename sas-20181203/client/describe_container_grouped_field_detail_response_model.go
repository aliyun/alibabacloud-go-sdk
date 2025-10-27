// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerGroupedFieldDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContainerGroupedFieldDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContainerGroupedFieldDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContainerGroupedFieldDetailResponseBody) *DescribeContainerGroupedFieldDetailResponse
	GetBody() *DescribeContainerGroupedFieldDetailResponseBody
}

type DescribeContainerGroupedFieldDetailResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerGroupedFieldDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContainerGroupedFieldDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerGroupedFieldDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupedFieldDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContainerGroupedFieldDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContainerGroupedFieldDetailResponse) GetBody() *DescribeContainerGroupedFieldDetailResponseBody {
	return s.Body
}

func (s *DescribeContainerGroupedFieldDetailResponse) SetHeaders(v map[string]*string) *DescribeContainerGroupedFieldDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponse) SetStatusCode(v int32) *DescribeContainerGroupedFieldDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponse) SetBody(v *DescribeContainerGroupedFieldDetailResponseBody) *DescribeContainerGroupedFieldDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeContainerGroupedFieldDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
