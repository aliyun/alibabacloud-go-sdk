// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceDownStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceDownStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceDownStreamResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceDownStreamResponseBody) *GetInstanceDownStreamResponse
	GetBody() *GetInstanceDownStreamResponseBody
}

type GetInstanceDownStreamResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceDownStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceDownStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDownStreamResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceDownStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceDownStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceDownStreamResponse) GetBody() *GetInstanceDownStreamResponseBody {
	return s.Body
}

func (s *GetInstanceDownStreamResponse) SetHeaders(v map[string]*string) *GetInstanceDownStreamResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceDownStreamResponse) SetStatusCode(v int32) *GetInstanceDownStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceDownStreamResponse) SetBody(v *GetInstanceDownStreamResponseBody) *GetInstanceDownStreamResponse {
	s.Body = v
	return s
}

func (s *GetInstanceDownStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
