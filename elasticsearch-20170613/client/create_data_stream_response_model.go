// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataStreamResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataStreamResponseBody) *CreateDataStreamResponse
	GetBody() *CreateDataStreamResponseBody
}

type CreateDataStreamResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataStreamResponse) GoString() string {
	return s.String()
}

func (s *CreateDataStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataStreamResponse) GetBody() *CreateDataStreamResponseBody {
	return s.Body
}

func (s *CreateDataStreamResponse) SetHeaders(v map[string]*string) *CreateDataStreamResponse {
	s.Headers = v
	return s
}

func (s *CreateDataStreamResponse) SetStatusCode(v int32) *CreateDataStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataStreamResponse) SetBody(v *CreateDataStreamResponseBody) *CreateDataStreamResponse {
	s.Body = v
	return s
}

func (s *CreateDataStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
