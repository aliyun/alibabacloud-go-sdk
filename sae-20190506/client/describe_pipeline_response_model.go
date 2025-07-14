// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePipelineResponse
	GetStatusCode() *int32
	SetBody(v *DescribePipelineResponseBody) *DescribePipelineResponse
	GetBody() *DescribePipelineResponseBody
}

type DescribePipelineResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePipelineResponse) GoString() string {
	return s.String()
}

func (s *DescribePipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePipelineResponse) GetBody() *DescribePipelineResponseBody {
	return s.Body
}

func (s *DescribePipelineResponse) SetHeaders(v map[string]*string) *DescribePipelineResponse {
	s.Headers = v
	return s
}

func (s *DescribePipelineResponse) SetStatusCode(v int32) *DescribePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePipelineResponse) SetBody(v *DescribePipelineResponseBody) *DescribePipelineResponse {
	s.Body = v
	return s
}

func (s *DescribePipelineResponse) Validate() error {
	return dara.Validate(s)
}
