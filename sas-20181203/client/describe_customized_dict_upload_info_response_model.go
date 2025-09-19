// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizedDictUploadInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomizedDictUploadInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomizedDictUploadInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomizedDictUploadInfoResponseBody) *DescribeCustomizedDictUploadInfoResponse
	GetBody() *DescribeCustomizedDictUploadInfoResponseBody
}

type DescribeCustomizedDictUploadInfoResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomizedDictUploadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomizedDictUploadInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizedDictUploadInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedDictUploadInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomizedDictUploadInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomizedDictUploadInfoResponse) GetBody() *DescribeCustomizedDictUploadInfoResponseBody {
	return s.Body
}

func (s *DescribeCustomizedDictUploadInfoResponse) SetHeaders(v map[string]*string) *DescribeCustomizedDictUploadInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizedDictUploadInfoResponse) SetStatusCode(v int32) *DescribeCustomizedDictUploadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizedDictUploadInfoResponse) SetBody(v *DescribeCustomizedDictUploadInfoResponseBody) *DescribeCustomizedDictUploadInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomizedDictUploadInfoResponse) Validate() error {
	return dara.Validate(s)
}
