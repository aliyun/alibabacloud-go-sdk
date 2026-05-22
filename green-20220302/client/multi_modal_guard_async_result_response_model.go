// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalGuardAsyncResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MultiModalGuardAsyncResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MultiModalGuardAsyncResultResponse
	GetStatusCode() *int32
	SetBody(v *MultiModalGuardAsyncResultResponseBody) *MultiModalGuardAsyncResultResponse
	GetBody() *MultiModalGuardAsyncResultResponseBody
}

type MultiModalGuardAsyncResultResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MultiModalGuardAsyncResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MultiModalGuardAsyncResultResponse) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncResultResponse) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MultiModalGuardAsyncResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MultiModalGuardAsyncResultResponse) GetBody() *MultiModalGuardAsyncResultResponseBody {
	return s.Body
}

func (s *MultiModalGuardAsyncResultResponse) SetHeaders(v map[string]*string) *MultiModalGuardAsyncResultResponse {
	s.Headers = v
	return s
}

func (s *MultiModalGuardAsyncResultResponse) SetStatusCode(v int32) *MultiModalGuardAsyncResultResponse {
	s.StatusCode = &v
	return s
}

func (s *MultiModalGuardAsyncResultResponse) SetBody(v *MultiModalGuardAsyncResultResponseBody) *MultiModalGuardAsyncResultResponse {
	s.Body = v
	return s
}

func (s *MultiModalGuardAsyncResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
