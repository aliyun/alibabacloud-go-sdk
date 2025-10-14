// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdcVersionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdcVersionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdcVersionListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdcVersionListResponseBody) *DescribeCdcVersionListResponse
	GetBody() *DescribeCdcVersionListResponseBody
}

type DescribeCdcVersionListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdcVersionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdcVersionListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdcVersionListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdcVersionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdcVersionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdcVersionListResponse) GetBody() *DescribeCdcVersionListResponseBody {
	return s.Body
}

func (s *DescribeCdcVersionListResponse) SetHeaders(v map[string]*string) *DescribeCdcVersionListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdcVersionListResponse) SetStatusCode(v int32) *DescribeCdcVersionListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdcVersionListResponse) SetBody(v *DescribeCdcVersionListResponseBody) *DescribeCdcVersionListResponse {
	s.Body = v
	return s
}

func (s *DescribeCdcVersionListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
