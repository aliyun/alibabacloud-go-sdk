// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutomateResponseConfigFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutomateResponseConfigFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutomateResponseConfigFeatureResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutomateResponseConfigFeatureResponseBody) *DescribeAutomateResponseConfigFeatureResponse
	GetBody() *DescribeAutomateResponseConfigFeatureResponseBody
}

type DescribeAutomateResponseConfigFeatureResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutomateResponseConfigFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutomateResponseConfigFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutomateResponseConfigFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutomateResponseConfigFeatureResponse) GetBody() *DescribeAutomateResponseConfigFeatureResponseBody {
	return s.Body
}

func (s *DescribeAutomateResponseConfigFeatureResponse) SetHeaders(v map[string]*string) *DescribeAutomateResponseConfigFeatureResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponse) SetStatusCode(v int32) *DescribeAutomateResponseConfigFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponse) SetBody(v *DescribeAutomateResponseConfigFeatureResponseBody) *DescribeAutomateResponseConfigFeatureResponse {
	s.Body = v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
