// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventPayloadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRiskEventPayloadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRiskEventPayloadResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRiskEventPayloadResponseBody) *DescribeRiskEventPayloadResponse
	GetBody() *DescribeRiskEventPayloadResponseBody
}

type DescribeRiskEventPayloadResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskEventPayloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskEventPayloadResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventPayloadResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventPayloadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRiskEventPayloadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRiskEventPayloadResponse) GetBody() *DescribeRiskEventPayloadResponseBody {
	return s.Body
}

func (s *DescribeRiskEventPayloadResponse) SetHeaders(v map[string]*string) *DescribeRiskEventPayloadResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskEventPayloadResponse) SetStatusCode(v int32) *DescribeRiskEventPayloadResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskEventPayloadResponse) SetBody(v *DescribeRiskEventPayloadResponseBody) *DescribeRiskEventPayloadResponse {
	s.Body = v
	return s
}

func (s *DescribeRiskEventPayloadResponse) Validate() error {
	return dara.Validate(s)
}
