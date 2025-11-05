// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEsExecuteDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEsExecuteDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEsExecuteDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEsExecuteDataResponseBody) *DescribeEsExecuteDataResponse
	GetBody() *DescribeEsExecuteDataResponseBody
}

type DescribeEsExecuteDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEsExecuteDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEsExecuteDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEsExecuteDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeEsExecuteDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEsExecuteDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEsExecuteDataResponse) GetBody() *DescribeEsExecuteDataResponseBody {
	return s.Body
}

func (s *DescribeEsExecuteDataResponse) SetHeaders(v map[string]*string) *DescribeEsExecuteDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeEsExecuteDataResponse) SetStatusCode(v int32) *DescribeEsExecuteDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEsExecuteDataResponse) SetBody(v *DescribeEsExecuteDataResponseBody) *DescribeEsExecuteDataResponse {
	s.Body = v
	return s
}

func (s *DescribeEsExecuteDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
