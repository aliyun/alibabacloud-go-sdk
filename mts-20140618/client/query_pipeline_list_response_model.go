// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPipelineListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPipelineListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPipelineListResponse
	GetStatusCode() *int32
	SetBody(v *QueryPipelineListResponseBody) *QueryPipelineListResponse
	GetBody() *QueryPipelineListResponseBody
}

type QueryPipelineListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPipelineListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPipelineListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPipelineListResponse) GoString() string {
	return s.String()
}

func (s *QueryPipelineListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPipelineListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPipelineListResponse) GetBody() *QueryPipelineListResponseBody {
	return s.Body
}

func (s *QueryPipelineListResponse) SetHeaders(v map[string]*string) *QueryPipelineListResponse {
	s.Headers = v
	return s
}

func (s *QueryPipelineListResponse) SetStatusCode(v int32) *QueryPipelineListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPipelineListResponse) SetBody(v *QueryPipelineListResponseBody) *QueryPipelineListResponse {
	s.Body = v
	return s
}

func (s *QueryPipelineListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
