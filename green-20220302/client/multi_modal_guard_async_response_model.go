// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalGuardAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MultiModalGuardAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MultiModalGuardAsyncResponse
	GetStatusCode() *int32
	SetBody(v *MultiModalGuardAsyncResponseBody) *MultiModalGuardAsyncResponse
	GetBody() *MultiModalGuardAsyncResponseBody
}

type MultiModalGuardAsyncResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MultiModalGuardAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MultiModalGuardAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncResponse) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MultiModalGuardAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MultiModalGuardAsyncResponse) GetBody() *MultiModalGuardAsyncResponseBody {
	return s.Body
}

func (s *MultiModalGuardAsyncResponse) SetHeaders(v map[string]*string) *MultiModalGuardAsyncResponse {
	s.Headers = v
	return s
}

func (s *MultiModalGuardAsyncResponse) SetStatusCode(v int32) *MultiModalGuardAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *MultiModalGuardAsyncResponse) SetBody(v *MultiModalGuardAsyncResponseBody) *MultiModalGuardAsyncResponse {
	s.Body = v
	return s
}

func (s *MultiModalGuardAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
