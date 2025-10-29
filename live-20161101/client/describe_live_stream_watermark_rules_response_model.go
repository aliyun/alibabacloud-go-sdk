// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamWatermarkRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamWatermarkRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamWatermarkRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamWatermarkRulesResponseBody) *DescribeLiveStreamWatermarkRulesResponse
	GetBody() *DescribeLiveStreamWatermarkRulesResponseBody
}

type DescribeLiveStreamWatermarkRulesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamWatermarkRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamWatermarkRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamWatermarkRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamWatermarkRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamWatermarkRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamWatermarkRulesResponse) GetBody() *DescribeLiveStreamWatermarkRulesResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamWatermarkRulesResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamWatermarkRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponse) SetStatusCode(v int32) *DescribeLiveStreamWatermarkRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponse) SetBody(v *DescribeLiveStreamWatermarkRulesResponseBody) *DescribeLiveStreamWatermarkRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamWatermarkRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
