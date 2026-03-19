// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContainerResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContainerResourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContainerResourceResponseBody) *DescribeContainerResourceResponse
	GetBody() *DescribeContainerResourceResponseBody
}

type DescribeContainerResourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContainerResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContainerResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContainerResourceResponse) GetBody() *DescribeContainerResourceResponseBody {
	return s.Body
}

func (s *DescribeContainerResourceResponse) SetHeaders(v map[string]*string) *DescribeContainerResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerResourceResponse) SetStatusCode(v int32) *DescribeContainerResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerResourceResponse) SetBody(v *DescribeContainerResourceResponseBody) *DescribeContainerResourceResponse {
	s.Body = v
	return s
}

func (s *DescribeContainerResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
