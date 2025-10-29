// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataStreamResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataStreamResponseBody) *DeleteDataStreamResponse
	GetBody() *DeleteDataStreamResponseBody
}

type DeleteDataStreamResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataStreamResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataStreamResponse) GetBody() *DeleteDataStreamResponseBody {
	return s.Body
}

func (s *DeleteDataStreamResponse) SetHeaders(v map[string]*string) *DeleteDataStreamResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataStreamResponse) SetStatusCode(v int32) *DeleteDataStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataStreamResponse) SetBody(v *DeleteDataStreamResponseBody) *DeleteDataStreamResponse {
	s.Body = v
	return s
}

func (s *DeleteDataStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
