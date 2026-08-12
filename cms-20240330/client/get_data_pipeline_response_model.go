// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataPipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataPipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataPipelineResponse
	GetStatusCode() *int32
	SetBody(v *GetDataPipelineResponseBody) *GetDataPipelineResponse
	GetBody() *GetDataPipelineResponseBody
}

type GetDataPipelineResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataPipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineResponse) GoString() string {
	return s.String()
}

func (s *GetDataPipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataPipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataPipelineResponse) GetBody() *GetDataPipelineResponseBody {
	return s.Body
}

func (s *GetDataPipelineResponse) SetHeaders(v map[string]*string) *GetDataPipelineResponse {
	s.Headers = v
	return s
}

func (s *GetDataPipelineResponse) SetStatusCode(v int32) *GetDataPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataPipelineResponse) SetBody(v *GetDataPipelineResponseBody) *GetDataPipelineResponse {
	s.Body = v
	return s
}

func (s *GetDataPipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
