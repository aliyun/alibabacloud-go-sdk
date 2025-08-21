// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRolloverDataStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RolloverDataStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RolloverDataStreamResponse
	GetStatusCode() *int32
	SetBody(v *RolloverDataStreamResponseBody) *RolloverDataStreamResponse
	GetBody() *RolloverDataStreamResponseBody
}

type RolloverDataStreamResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RolloverDataStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RolloverDataStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s RolloverDataStreamResponse) GoString() string {
	return s.String()
}

func (s *RolloverDataStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RolloverDataStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RolloverDataStreamResponse) GetBody() *RolloverDataStreamResponseBody {
	return s.Body
}

func (s *RolloverDataStreamResponse) SetHeaders(v map[string]*string) *RolloverDataStreamResponse {
	s.Headers = v
	return s
}

func (s *RolloverDataStreamResponse) SetStatusCode(v int32) *RolloverDataStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *RolloverDataStreamResponse) SetBody(v *RolloverDataStreamResponseBody) *RolloverDataStreamResponse {
	s.Body = v
	return s
}

func (s *RolloverDataStreamResponse) Validate() error {
	return dara.Validate(s)
}
