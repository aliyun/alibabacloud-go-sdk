// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisWithStageNameIntegratedByAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisWithStageNameIntegratedByAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisWithStageNameIntegratedByAppResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisWithStageNameIntegratedByAppResponseBody) *DescribeApisWithStageNameIntegratedByAppResponse
	GetBody() *DescribeApisWithStageNameIntegratedByAppResponseBody
}

type DescribeApisWithStageNameIntegratedByAppResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisWithStageNameIntegratedByAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisWithStageNameIntegratedByAppResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisWithStageNameIntegratedByAppResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisWithStageNameIntegratedByAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisWithStageNameIntegratedByAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisWithStageNameIntegratedByAppResponse) GetBody() *DescribeApisWithStageNameIntegratedByAppResponseBody {
	return s.Body
}

func (s *DescribeApisWithStageNameIntegratedByAppResponse) SetHeaders(v map[string]*string) *DescribeApisWithStageNameIntegratedByAppResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponse) SetStatusCode(v int32) *DescribeApisWithStageNameIntegratedByAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponse) SetBody(v *DescribeApisWithStageNameIntegratedByAppResponseBody) *DescribeApisWithStageNameIntegratedByAppResponse {
	s.Body = v
	return s
}

func (s *DescribeApisWithStageNameIntegratedByAppResponse) Validate() error {
	return dara.Validate(s)
}
