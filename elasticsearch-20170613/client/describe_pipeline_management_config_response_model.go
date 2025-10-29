// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePipelineManagementConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePipelineManagementConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePipelineManagementConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribePipelineManagementConfigResponseBody) *DescribePipelineManagementConfigResponse
	GetBody() *DescribePipelineManagementConfigResponseBody
}

type DescribePipelineManagementConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePipelineManagementConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePipelineManagementConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePipelineManagementConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribePipelineManagementConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePipelineManagementConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePipelineManagementConfigResponse) GetBody() *DescribePipelineManagementConfigResponseBody {
	return s.Body
}

func (s *DescribePipelineManagementConfigResponse) SetHeaders(v map[string]*string) *DescribePipelineManagementConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribePipelineManagementConfigResponse) SetStatusCode(v int32) *DescribePipelineManagementConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePipelineManagementConfigResponse) SetBody(v *DescribePipelineManagementConfigResponseBody) *DescribePipelineManagementConfigResponse {
	s.Body = v
	return s
}

func (s *DescribePipelineManagementConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
