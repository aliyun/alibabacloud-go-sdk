// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosThresholdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDdosThresholdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDdosThresholdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDdosThresholdResponseBody) *DescribeDdosThresholdResponse
	GetBody() *DescribeDdosThresholdResponseBody
}

type DescribeDdosThresholdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDdosThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDdosThresholdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosThresholdResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosThresholdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDdosThresholdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDdosThresholdResponse) GetBody() *DescribeDdosThresholdResponseBody {
	return s.Body
}

func (s *DescribeDdosThresholdResponse) SetHeaders(v map[string]*string) *DescribeDdosThresholdResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosThresholdResponse) SetStatusCode(v int32) *DescribeDdosThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosThresholdResponse) SetBody(v *DescribeDdosThresholdResponseBody) *DescribeDdosThresholdResponse {
	s.Body = v
	return s
}

func (s *DescribeDdosThresholdResponse) Validate() error {
	return dara.Validate(s)
}
