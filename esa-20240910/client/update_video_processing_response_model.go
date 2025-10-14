// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoProcessingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVideoProcessingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVideoProcessingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVideoProcessingResponseBody) *UpdateVideoProcessingResponse
	GetBody() *UpdateVideoProcessingResponseBody
}

type UpdateVideoProcessingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVideoProcessingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVideoProcessingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoProcessingResponse) GoString() string {
	return s.String()
}

func (s *UpdateVideoProcessingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVideoProcessingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVideoProcessingResponse) GetBody() *UpdateVideoProcessingResponseBody {
	return s.Body
}

func (s *UpdateVideoProcessingResponse) SetHeaders(v map[string]*string) *UpdateVideoProcessingResponse {
	s.Headers = v
	return s
}

func (s *UpdateVideoProcessingResponse) SetStatusCode(v int32) *UpdateVideoProcessingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVideoProcessingResponse) SetBody(v *UpdateVideoProcessingResponseBody) *UpdateVideoProcessingResponse {
	s.Body = v
	return s
}

func (s *UpdateVideoProcessingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
