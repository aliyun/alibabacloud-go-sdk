// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStageModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStageModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStageModelsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStageModelsResponseBody) *DescribeStageModelsResponse
	GetBody() *DescribeStageModelsResponseBody
}

type DescribeStageModelsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStageModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStageModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStageModelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeStageModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStageModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStageModelsResponse) GetBody() *DescribeStageModelsResponseBody {
	return s.Body
}

func (s *DescribeStageModelsResponse) SetHeaders(v map[string]*string) *DescribeStageModelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeStageModelsResponse) SetStatusCode(v int32) *DescribeStageModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStageModelsResponse) SetBody(v *DescribeStageModelsResponseBody) *DescribeStageModelsResponse {
	s.Body = v
	return s
}

func (s *DescribeStageModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
