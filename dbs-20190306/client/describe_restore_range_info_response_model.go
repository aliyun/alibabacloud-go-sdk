// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreRangeInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRestoreRangeInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRestoreRangeInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRestoreRangeInfoResponseBody) *DescribeRestoreRangeInfoResponse
	GetBody() *DescribeRestoreRangeInfoResponseBody
}

type DescribeRestoreRangeInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreRangeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreRangeInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreRangeInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreRangeInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRestoreRangeInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRestoreRangeInfoResponse) GetBody() *DescribeRestoreRangeInfoResponseBody {
	return s.Body
}

func (s *DescribeRestoreRangeInfoResponse) SetHeaders(v map[string]*string) *DescribeRestoreRangeInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreRangeInfoResponse) SetStatusCode(v int32) *DescribeRestoreRangeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreRangeInfoResponse) SetBody(v *DescribeRestoreRangeInfoResponseBody) *DescribeRestoreRangeInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeRestoreRangeInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
