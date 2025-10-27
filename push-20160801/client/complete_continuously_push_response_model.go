// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteContinuouslyPushResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompleteContinuouslyPushResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompleteContinuouslyPushResponse
	GetStatusCode() *int32
	SetBody(v *CompleteContinuouslyPushResponseBody) *CompleteContinuouslyPushResponse
	GetBody() *CompleteContinuouslyPushResponseBody
}

type CompleteContinuouslyPushResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompleteContinuouslyPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompleteContinuouslyPushResponse) String() string {
	return dara.Prettify(s)
}

func (s CompleteContinuouslyPushResponse) GoString() string {
	return s.String()
}

func (s *CompleteContinuouslyPushResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompleteContinuouslyPushResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompleteContinuouslyPushResponse) GetBody() *CompleteContinuouslyPushResponseBody {
	return s.Body
}

func (s *CompleteContinuouslyPushResponse) SetHeaders(v map[string]*string) *CompleteContinuouslyPushResponse {
	s.Headers = v
	return s
}

func (s *CompleteContinuouslyPushResponse) SetStatusCode(v int32) *CompleteContinuouslyPushResponse {
	s.StatusCode = &v
	return s
}

func (s *CompleteContinuouslyPushResponse) SetBody(v *CompleteContinuouslyPushResponseBody) *CompleteContinuouslyPushResponse {
	s.Body = v
	return s
}

func (s *CompleteContinuouslyPushResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
