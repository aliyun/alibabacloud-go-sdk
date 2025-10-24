// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchPipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchPipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchPipelineResponse
	GetStatusCode() *int32
	SetBody(v *SearchPipelineResponseBody) *SearchPipelineResponse
	GetBody() *SearchPipelineResponseBody
}

type SearchPipelineResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchPipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchPipelineResponse) GoString() string {
	return s.String()
}

func (s *SearchPipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchPipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchPipelineResponse) GetBody() *SearchPipelineResponseBody {
	return s.Body
}

func (s *SearchPipelineResponse) SetHeaders(v map[string]*string) *SearchPipelineResponse {
	s.Headers = v
	return s
}

func (s *SearchPipelineResponse) SetStatusCode(v int32) *SearchPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchPipelineResponse) SetBody(v *SearchPipelineResponseBody) *SearchPipelineResponse {
	s.Body = v
	return s
}

func (s *SearchPipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
