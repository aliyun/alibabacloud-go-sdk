// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsNetLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnsNetLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnsNetLevelResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnsNetLevelResponseBody) *DescribeEnsNetLevelResponse
	GetBody() *DescribeEnsNetLevelResponseBody
}

type DescribeEnsNetLevelResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnsNetLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnsNetLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsNetLevelResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsNetLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnsNetLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnsNetLevelResponse) GetBody() *DescribeEnsNetLevelResponseBody {
	return s.Body
}

func (s *DescribeEnsNetLevelResponse) SetHeaders(v map[string]*string) *DescribeEnsNetLevelResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsNetLevelResponse) SetStatusCode(v int32) *DescribeEnsNetLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnsNetLevelResponse) SetBody(v *DescribeEnsNetLevelResponseBody) *DescribeEnsNetLevelResponse {
	s.Body = v
	return s
}

func (s *DescribeEnsNetLevelResponse) Validate() error {
	return dara.Validate(s)
}
