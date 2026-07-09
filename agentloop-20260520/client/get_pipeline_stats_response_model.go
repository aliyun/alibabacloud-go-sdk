// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPipelineStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPipelineStatsResponse
	GetStatusCode() *int32
	SetBody(v *GetPipelineStatsResponseBody) *GetPipelineStatsResponse
	GetBody() *GetPipelineStatsResponseBody
}

type GetPipelineStatsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPipelineStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPipelineStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineStatsResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPipelineStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPipelineStatsResponse) GetBody() *GetPipelineStatsResponseBody {
	return s.Body
}

func (s *GetPipelineStatsResponse) SetHeaders(v map[string]*string) *GetPipelineStatsResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineStatsResponse) SetStatusCode(v int32) *GetPipelineStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineStatsResponse) SetBody(v *GetPipelineStatsResponseBody) *GetPipelineStatsResponse {
	s.Body = v
	return s
}

func (s *GetPipelineStatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
