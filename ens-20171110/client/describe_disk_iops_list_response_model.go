// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskIopsListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiskIopsListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiskIopsListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiskIopsListResponseBody) *DescribeDiskIopsListResponse
	GetBody() *DescribeDiskIopsListResponseBody
}

type DescribeDiskIopsListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskIopsListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskIopsListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskIopsListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskIopsListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiskIopsListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiskIopsListResponse) GetBody() *DescribeDiskIopsListResponseBody {
	return s.Body
}

func (s *DescribeDiskIopsListResponse) SetHeaders(v map[string]*string) *DescribeDiskIopsListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskIopsListResponse) SetStatusCode(v int32) *DescribeDiskIopsListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskIopsListResponse) SetBody(v *DescribeDiskIopsListResponseBody) *DescribeDiskIopsListResponse {
	s.Body = v
	return s
}

func (s *DescribeDiskIopsListResponse) Validate() error {
	return dara.Validate(s)
}
