// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMPULayoutInfoListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMPULayoutInfoListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMPULayoutInfoListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMPULayoutInfoListResponseBody) *DescribeMPULayoutInfoListResponse
	GetBody() *DescribeMPULayoutInfoListResponseBody
}

type DescribeMPULayoutInfoListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMPULayoutInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMPULayoutInfoListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMPULayoutInfoListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMPULayoutInfoListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMPULayoutInfoListResponse) GetBody() *DescribeMPULayoutInfoListResponseBody {
	return s.Body
}

func (s *DescribeMPULayoutInfoListResponse) SetHeaders(v map[string]*string) *DescribeMPULayoutInfoListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMPULayoutInfoListResponse) SetStatusCode(v int32) *DescribeMPULayoutInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponse) SetBody(v *DescribeMPULayoutInfoListResponseBody) *DescribeMPULayoutInfoListResponse {
	s.Body = v
	return s
}

func (s *DescribeMPULayoutInfoListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
