// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOnceTaskLeafRecordPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOnceTaskLeafRecordPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOnceTaskLeafRecordPageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOnceTaskLeafRecordPageResponseBody) *DescribeOnceTaskLeafRecordPageResponse
	GetBody() *DescribeOnceTaskLeafRecordPageResponseBody
}

type DescribeOnceTaskLeafRecordPageResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOnceTaskLeafRecordPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOnceTaskLeafRecordPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnceTaskLeafRecordPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeOnceTaskLeafRecordPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOnceTaskLeafRecordPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOnceTaskLeafRecordPageResponse) GetBody() *DescribeOnceTaskLeafRecordPageResponseBody {
	return s.Body
}

func (s *DescribeOnceTaskLeafRecordPageResponse) SetHeaders(v map[string]*string) *DescribeOnceTaskLeafRecordPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponse) SetStatusCode(v int32) *DescribeOnceTaskLeafRecordPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponse) SetBody(v *DescribeOnceTaskLeafRecordPageResponseBody) *DescribeOnceTaskLeafRecordPageResponse {
	s.Body = v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
