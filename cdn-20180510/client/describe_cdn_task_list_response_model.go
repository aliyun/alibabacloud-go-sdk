// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnTaskListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnTaskListResponseBody) *DescribeCdnTaskListResponse
	GetBody() *DescribeCdnTaskListResponseBody
}

type DescribeCdnTaskListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnTaskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnTaskListResponse) GetBody() *DescribeCdnTaskListResponseBody {
	return s.Body
}

func (s *DescribeCdnTaskListResponse) SetHeaders(v map[string]*string) *DescribeCdnTaskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnTaskListResponse) SetStatusCode(v int32) *DescribeCdnTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnTaskListResponse) SetBody(v *DescribeCdnTaskListResponseBody) *DescribeCdnTaskListResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
