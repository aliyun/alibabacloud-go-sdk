// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoProcessingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVideoProcessingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVideoProcessingResponse
	GetStatusCode() *int32
	SetBody(v *CreateVideoProcessingResponseBody) *CreateVideoProcessingResponse
	GetBody() *CreateVideoProcessingResponseBody
}

type CreateVideoProcessingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVideoProcessingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVideoProcessingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoProcessingResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoProcessingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVideoProcessingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVideoProcessingResponse) GetBody() *CreateVideoProcessingResponseBody {
	return s.Body
}

func (s *CreateVideoProcessingResponse) SetHeaders(v map[string]*string) *CreateVideoProcessingResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoProcessingResponse) SetStatusCode(v int32) *CreateVideoProcessingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVideoProcessingResponse) SetBody(v *CreateVideoProcessingResponseBody) *CreateVideoProcessingResponse {
	s.Body = v
	return s
}

func (s *CreateVideoProcessingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
