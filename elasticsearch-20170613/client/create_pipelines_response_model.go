// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePipelinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePipelinesResponse
	GetStatusCode() *int32
	SetBody(v *CreatePipelinesResponseBody) *CreatePipelinesResponse
	GetBody() *CreatePipelinesResponseBody
}

type CreatePipelinesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePipelinesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelinesResponse) GoString() string {
	return s.String()
}

func (s *CreatePipelinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePipelinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePipelinesResponse) GetBody() *CreatePipelinesResponseBody {
	return s.Body
}

func (s *CreatePipelinesResponse) SetHeaders(v map[string]*string) *CreatePipelinesResponse {
	s.Headers = v
	return s
}

func (s *CreatePipelinesResponse) SetStatusCode(v int32) *CreatePipelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePipelinesResponse) SetBody(v *CreatePipelinesResponseBody) *CreatePipelinesResponse {
	s.Body = v
	return s
}

func (s *CreatePipelinesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
