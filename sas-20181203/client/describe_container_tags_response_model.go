// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContainerTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContainerTagsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContainerTagsResponseBody) *DescribeContainerTagsResponse
	GetBody() *DescribeContainerTagsResponseBody
}

type DescribeContainerTagsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContainerTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContainerTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContainerTagsResponse) GetBody() *DescribeContainerTagsResponseBody {
	return s.Body
}

func (s *DescribeContainerTagsResponse) SetHeaders(v map[string]*string) *DescribeContainerTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerTagsResponse) SetStatusCode(v int32) *DescribeContainerTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerTagsResponse) SetBody(v *DescribeContainerTagsResponseBody) *DescribeContainerTagsResponse {
	s.Body = v
	return s
}

func (s *DescribeContainerTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
