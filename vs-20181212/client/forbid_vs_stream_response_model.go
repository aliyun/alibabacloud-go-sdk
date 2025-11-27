// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForbidVsStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ForbidVsStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ForbidVsStreamResponse
	GetStatusCode() *int32
	SetBody(v *ForbidVsStreamResponseBody) *ForbidVsStreamResponse
	GetBody() *ForbidVsStreamResponseBody
}

type ForbidVsStreamResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ForbidVsStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ForbidVsStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s ForbidVsStreamResponse) GoString() string {
	return s.String()
}

func (s *ForbidVsStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ForbidVsStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ForbidVsStreamResponse) GetBody() *ForbidVsStreamResponseBody {
	return s.Body
}

func (s *ForbidVsStreamResponse) SetHeaders(v map[string]*string) *ForbidVsStreamResponse {
	s.Headers = v
	return s
}

func (s *ForbidVsStreamResponse) SetStatusCode(v int32) *ForbidVsStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *ForbidVsStreamResponse) SetBody(v *ForbidVsStreamResponseBody) *ForbidVsStreamResponse {
	s.Body = v
	return s
}

func (s *ForbidVsStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
