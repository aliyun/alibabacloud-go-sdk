// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeUpDownStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeUpDownStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeUpDownStreamResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeUpDownStreamResponseBody) *GetNodeUpDownStreamResponse
	GetBody() *GetNodeUpDownStreamResponseBody
}

type GetNodeUpDownStreamResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeUpDownStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeUpDownStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeUpDownStreamResponse) GoString() string {
	return s.String()
}

func (s *GetNodeUpDownStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeUpDownStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeUpDownStreamResponse) GetBody() *GetNodeUpDownStreamResponseBody {
	return s.Body
}

func (s *GetNodeUpDownStreamResponse) SetHeaders(v map[string]*string) *GetNodeUpDownStreamResponse {
	s.Headers = v
	return s
}

func (s *GetNodeUpDownStreamResponse) SetStatusCode(v int32) *GetNodeUpDownStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeUpDownStreamResponse) SetBody(v *GetNodeUpDownStreamResponseBody) *GetNodeUpDownStreamResponse {
	s.Body = v
	return s
}

func (s *GetNodeUpDownStreamResponse) Validate() error {
	return dara.Validate(s)
}
