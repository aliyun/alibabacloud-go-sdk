// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveTagResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveTagResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveTagResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveTagResourcesResponseBody) *DescribeLiveTagResourcesResponse
	GetBody() *DescribeLiveTagResourcesResponseBody
}

type DescribeLiveTagResourcesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveTagResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveTagResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveTagResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveTagResourcesResponse) GetBody() *DescribeLiveTagResourcesResponseBody {
	return s.Body
}

func (s *DescribeLiveTagResourcesResponse) SetHeaders(v map[string]*string) *DescribeLiveTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveTagResourcesResponse) SetStatusCode(v int32) *DescribeLiveTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveTagResourcesResponse) SetBody(v *DescribeLiveTagResourcesResponseBody) *DescribeLiveTagResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveTagResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
