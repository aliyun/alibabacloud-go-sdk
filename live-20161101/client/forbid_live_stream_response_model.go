// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForbidLiveStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ForbidLiveStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ForbidLiveStreamResponse
	GetStatusCode() *int32
	SetBody(v *ForbidLiveStreamResponseBody) *ForbidLiveStreamResponse
	GetBody() *ForbidLiveStreamResponseBody
}

type ForbidLiveStreamResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ForbidLiveStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ForbidLiveStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s ForbidLiveStreamResponse) GoString() string {
	return s.String()
}

func (s *ForbidLiveStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ForbidLiveStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ForbidLiveStreamResponse) GetBody() *ForbidLiveStreamResponseBody {
	return s.Body
}

func (s *ForbidLiveStreamResponse) SetHeaders(v map[string]*string) *ForbidLiveStreamResponse {
	s.Headers = v
	return s
}

func (s *ForbidLiveStreamResponse) SetStatusCode(v int32) *ForbidLiveStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *ForbidLiveStreamResponse) SetBody(v *ForbidLiveStreamResponseBody) *ForbidLiveStreamResponse {
	s.Body = v
	return s
}

func (s *ForbidLiveStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
