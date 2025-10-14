// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdcClassListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdcClassListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdcClassListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdcClassListResponseBody) *DescribeCdcClassListResponse
	GetBody() *DescribeCdcClassListResponseBody
}

type DescribeCdcClassListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdcClassListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdcClassListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdcClassListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdcClassListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdcClassListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdcClassListResponse) GetBody() *DescribeCdcClassListResponseBody {
	return s.Body
}

func (s *DescribeCdcClassListResponse) SetHeaders(v map[string]*string) *DescribeCdcClassListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdcClassListResponse) SetStatusCode(v int32) *DescribeCdcClassListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdcClassListResponse) SetBody(v *DescribeCdcClassListResponseBody) *DescribeCdcClassListResponse {
	s.Body = v
	return s
}

func (s *DescribeCdcClassListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
